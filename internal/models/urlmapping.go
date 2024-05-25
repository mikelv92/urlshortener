package models

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type URLMappingModelInterface interface {
	Exists(shortURL string) (bool, error)
	Insert(shortURL, longURL string) error
	Find(shortURL string) (URLMapping, error)
}

type URLMapping struct {
	LongURL  string
	ShortURL string
}

type URLMappingModel struct {
	MongoClient *mongo.Client
}

func (m *URLMappingModel) Exists(shortURL string) (bool, error) {
	_, err := m.Find(shortURL)
	if err != nil {
		if err == ErrNoRecord {
			return false, nil
		} else {
			return false, err
		}
	}

	return true, nil
}

func (m *URLMappingModel) Insert(shortURL, longURL string) error {
	collection := m.MongoClient.Database("urls").Collection("urlmappings")
	doc := URLMapping{ShortURL: shortURL, LongURL: longURL}
	_, err := collection.InsertOne(context.TODO(), doc)
	if err != nil {
		return err
	}
	return nil
}

func (m *URLMappingModel) Find(shortURL string) (URLMapping, error) {
	collection := m.MongoClient.Database("urls").Collection("urlmappings")

	var result URLMapping
	err := collection.FindOne(context.TODO(), bson.D{{Key: "shorturl", Value: shortURL}}).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return URLMapping{}, ErrNoRecord
		} else {
			return URLMapping{}, err
		}
	}

	return result, nil
}
