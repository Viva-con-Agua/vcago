package vmod

// IDParam data struct for handling '/:id'.
// ID needs to be a uuid.
type IDParam struct {
	ID string `param:"id" validate:"uuid"`
}

// DeletedResponse used for handling response in deleted case.
type DeletedResponse struct {
	ID string `json:"id"`
}

// NewDeletedResponse creates a new DeletedResponse type.
func NewDeletedResponse(id string) *DeletedResponse {
	return &DeletedResponse{ID: id}
}

// Count is used for counting results in db aggregations
type Count struct {
	Total int32 `json:"total" bson:"total"`
}
