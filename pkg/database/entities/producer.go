package entities

import "time"

type Producer struct {
	Id           int32     `pg:",pk,unique,notnull"`
	Uuid         string    `pg:"type:uuid,pk,unique,notnull,default:uuid_generate_v4()"`
	ProducerName string    `pg:",notnull"`
	CanDeliver   bool      `pg:"notnull"`
	CreatedAt    time.Time `pg:"default:now()"`
	// relationships
	Products []*Product `pg:"rel:has-many"`
}
