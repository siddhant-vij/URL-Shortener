package repository

import (
	"context"
	"errors"
	"sync"
	"url-shortener/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const mongoUri = "mongodb://localhost:27017"

type URLRepository interface {
	Save(urlData model.URLData) error
	FindByShortURL(shortURL string) (*model.URLData, error)
	Exists(shortURL string) bool
}

// In-Memory Storage: Utilizes an in-memory repository for fast retrieval and storage of URL mappings.
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

// Persistence Storage: Integrate a database (MongoDB) to persist URL mappings beyond the application lifecycle.
type MongoDBURLStore struct {
	collection *mongo.Collection
}

func NewMongoDBURLStore() (*MongoDBURLStore, error) {
	clientOptions := options.Client().ApplyURI(mongoUri)
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		return nil, err
	}

	collection := client.Database("url-shortener").Collection("urls")

	return &MongoDBURLStore{collection: collection}, nil
}

func (store *MongoDBURLStore) Save(urlData model.URLData) error {
	_, err := store.collection.InsertOne(context.Background(), urlData)
	if err != nil {
		return err
	}
	return nil
}

func (store *MongoDBURLStore) FindByShortURL(shortURL string) (*model.URLData, error) {
	filter := bson.M{"shortenedurl": shortURL}
	var result model.URLData
	err := store.collection.FindOne(context.Background(), filter).Decode(&result)
	if err != nil {
		return nil, errors.New("shortened URL not found")
	}
	return &result, nil
}

func (store *MongoDBURLStore) Exists(shortURL string) bool {
	filter := bson.M{"shortenedurl": shortURL}
	err := store.collection.FindOne(context.Background(), filter).Decode(&model.URLData{})
	return err == nil
}
