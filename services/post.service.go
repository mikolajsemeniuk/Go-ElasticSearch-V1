package services

import (
	"github.com/mikolajsemeniuk/go-react-elasticsearch/domain"
	"github.com/mikolajsemeniuk/go-react-elasticsearch/inputs"
	"github.com/mikolajsemeniuk/go-react-elasticsearch/repositories"
)

var (
	PostService IPostService = postService{}
)

type IPostService interface {
	All() ([]domain.Post, error)
	Add(input inputs.Post) error
}

type postService struct{}

func (postService) All() ([]domain.Post, error) {
	posts, err := repositories.PostRepository.All()
	if err != nil {
		return nil, err
	}
	return posts, nil
}

func (postService) Add(input inputs.Post) error {
	err := repositories.PostRepository.Add(input)
	if err != nil {
		return err
	}
	return nil
}
