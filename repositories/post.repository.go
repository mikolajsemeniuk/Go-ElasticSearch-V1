package repositories

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/elastic/go-elasticsearch/esapi"
	"github.com/google/uuid"
	"github.com/jinzhu/copier"
	"github.com/mikolajsemeniuk/go-react-elasticsearch/data"
	"github.com/mikolajsemeniuk/go-react-elasticsearch/entities"
	"github.com/mikolajsemeniuk/go-react-elasticsearch/inputs"
	"github.com/mikolajsemeniuk/go-react-elasticsearch/payloads"
	"github.com/mitchellh/mapstructure"
)

const index = "posts"

var (
	PostRepository IPostRepository = postRepository{}
)

type IPostRepository interface {
	All() ([]payloads.Post, error)
	Add(inputs.Post) error
}

type postRepository struct{}

func (postRepository) All() ([]payloads.Post, error) {
	var buffer bytes.Buffer
	query := map[string]interface{}{
		"query": map[string]interface{}{
			"match_all": map[string]interface{}{},
		},
	}

	if err := json.NewEncoder(&buffer).Encode(query); err != nil {
		return nil, err
	}

	response, err := data.ElasticSearchClient.Search(
		data.ElasticSearchClient.Search.WithContext(context.Background()),
		data.ElasticSearchClient.Search.WithIndex(index),
		data.ElasticSearchClient.Search.WithBody(&buffer),
	)

	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	var body map[string]interface{}
	if response.IsError() {
		err = json.NewDecoder(response.Body).Decode(&body)
	}

	if response.IsError() && err != nil {
		return nil, fmt.Errorf("error parsing the response body: %s", err)
	}

	if response.IsError() && err == nil {
		return nil, errors.New(body["error"].(map[string]interface{})["reason"].(string))
	}

	if err := json.NewDecoder(response.Body).Decode(&body); err != nil {
		return nil, fmt.Errorf("error parsing the response body: %s", err)
	}

	posts := []payloads.Post{}
	for _, hit := range body["hits"].(map[string]interface{})["hits"].([]interface{}) {
		var post payloads.Post
		mapstructure.Decode(hit.(map[string]interface{})["_source"].(map[string]interface{}), &post)
		post.Id, _ = uuid.Parse(hit.(map[string]interface{})["_id"].(string))
		posts = append(posts, post)
	}

	return posts, nil
}

func (postRepository) Add(input inputs.Post) error {
	channel := make(chan error)
	go func() {
		post := entities.Post{
			Title: "dasdsa",
			// Created: time.Now().Format(time.RFC3339),
			Created: time.Now().Format(time.RFC3339),
		}

		copier.Copy(&post, &input)

		fmt.Println("post: ", post)

		body, err := json.Marshal(post)

		if err != nil {
			channel <- err
		}

		request := esapi.IndexRequest{
			Index:      index,
			DocumentID: uuid.New().String(),
			Body:       strings.NewReader(string(body)),
			Refresh:    "true",
		}

		res, err := request.Do(context.Background(), data.ElasticSearchClient)
		if err != nil {
			channel <- err
		}
		defer res.Body.Close()

		if res.IsError() {
			channel <- fmt.Errorf("[%s] Error indexing document", res.Status())
		}
		channel <- nil
	}()
	return <-channel
}
