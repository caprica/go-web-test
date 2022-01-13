package repository

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	"go-web/pkg/api"
)

type Storage interface {
	FindAlbums(albumArtistName string) ([]api.Album, error)
}

type storage struct {
	client *mongo.Client
}

func NewStorage(client *mongo.Client) Storage {
	return &storage{client: client}
}

func (s *storage) FindAlbums(albumArtistName string) ([]api.Album, error) {
	coll := s.client.Database("choonio").Collection("album")
	filter := bson.M{"mediaId.albumArtistName": albumArtistName}
	cursor, err := coll.Find(context.TODO(), filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())

	var albums []api.Album
	for cursor.Next(context.TODO()) {
		var album api.Album
		cursor.Decode(&album)
		albums = append(albums, album)
	}

	return albums, nil
}
