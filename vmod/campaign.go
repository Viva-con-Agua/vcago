package vmod

type (

	//Campaign represents a default campaign model
	Campaign struct {
		ID          string   `bson:"_id" json:"id" validate:"required"`
		CiviID      int      `bson:"civi_id" json:"civi_id" validate:"required"`
		Name        string   `bson:"name" json:"name" validate:"required"`
		Description string   `bson:"name" json:"description" validate:"required"`
		Tags        []string `bson:"tags" json:"tags" validate:"required"`
		Type        string   `bson:"type" json:"type" validate:"required"`
		Modified    Modified `bson:"modified" json:"modified" validate:"required"`
	}
)
