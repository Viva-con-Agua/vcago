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
