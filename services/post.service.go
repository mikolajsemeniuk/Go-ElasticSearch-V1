package services

import (
	"github.com/google/uuid"
	"github.com/mikolajsemeniuk/go-react-elasticsearch/inputs"
	"github.com/mikolajsemeniuk/go-react-elasticsearch/payloads"
	"github.com/mikolajsemeniuk/go-react-elasticsearch/repositories"
)

var (
	PostService IPostService = postService{}
)

type IPostService interface {
	All() ([]payloads.Post, error)
	Add(input inputs.Post) error
	Remove(id uuid.UUID) error
	Update(id uuid.UUID, input inputs.Post) error
}

type postService struct{}

func (postService) All() ([]payloads.Post, error) {
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

func (postService) Remove(id uuid.UUID) error {
	err := repositories.PostRepository.Remove(id)
	if err != nil {
		return err
	}
	return nil
}

func (postService) Update(id uuid.UUID, input inputs.Post) error {
	err := repositories.PostRepository.Update(id, input)
	if err != nil {
		return err
	}
	return nil
}
