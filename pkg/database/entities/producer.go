package entities

import "time"

type Producer struct {
	Id           int32  `pg:",pk,unique"`
	Uuid         string `pg:"type:uuid,pk,unique,default:uuid_generate_v4()"`
	ProducerName string
	CanDeliver   bool
	CreatedAt    time.Time `pg:"default:now()"`
	// relationships
	Products []*Product `pg:"rel:has-many"`
}
