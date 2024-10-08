package vmod

import (
	"mime/multipart"

	"go.mongodb.org/mongo-driver/bson"
)

// IDParam data struct for handling '/:id'.
// ID needs to be a uuid.
type IDParam struct {
	ID string `param:"id" validate:"uuid"`
}

func (i *IDParam) Filter() bson.D {
	return bson.D{{Key: "_id", Value: i.ID}}
}

// DeletedResponse used for handling response in deleted case.
type DeletedResponse struct {
	ID string `json:"id"`
}

type Content struct {
	Fields map[string]interface{} `json:"content" bson:"content"`
}

// NewDeletedResponse creates a new DeletedResponse type.
func NewDeletedResponse(id string) *DeletedResponse {
	return &DeletedResponse{ID: id}
}

// Count is used for counting results in db aggregations
type Count struct {
	Total int32 `json:"total" bson:"total"`
}

type File struct {
	File   multipart.File
	Header *multipart.FileHeader
}
