package api

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MediaId struct {
	Type            string
	AlbumArtistName string
	AlbumName       string
}

type Album struct {
	ID        primitive.ObjectID `json:"id" bson:"_id"`
	MediaId   MediaId
	Year      int16
	Genre     string
	Tracks    []Track
	Duration  int32
	Rgb       []int16
	Timestamp time.Time
	Location  string
}

type Track struct {
	MediaId    MediaId
	Number     int16
	ArtistName string
	Duration   int16
	Filename   string
}
