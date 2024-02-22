package repository

import (
	"errors"
	"sync"
	"url-shortener/model"
)

type URLRepository interface {
	Save(urlData model.URLData) error
	FindByShortURL(shortURL string) (*model.URLData, error)
	Exists(shortURL string) bool
}

type InMemoryURLStore struct {
	urlMap sync.Map
}

func NewInMemoryURLStore() *InMemoryURLStore {
	return &InMemoryURLStore{}
}

func (store *InMemoryURLStore) Save(urlData model.URLData) error {
	store.urlMap.Store(urlData.ShortenedURL, urlData)
	return nil
}

func (store *InMemoryURLStore) FindByShortURL(shortURL string) (*model.URLData, error) {
	if val, ok := store.urlMap.Load(shortURL); ok {
		urlData := val.(model.URLData)
		return &urlData, nil
	}
	return nil, errors.New("shortened URL not found")
}

func (store *InMemoryURLStore) Exists(shortURL string) bool {
	_, ok := store.urlMap.Load(shortURL)
	return ok
}
