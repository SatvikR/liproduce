package entities

import "time"

type Product struct {
	Id          int32     `pg:",pk,unique,notnull"`
	Uuid        string    `pg:"type:uuid,pk,unique,notnull,default:uuid_generate_v4()"`
	ProductName string    `pg:",notnull"`
	OwnerId     int32     `pg:",notnull"`
	CreatedAt   time.Time `pg:"default:now()"`
	// relationships
	Owner *Producer `pg:"has-one"`
}
