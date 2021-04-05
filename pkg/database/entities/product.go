package entities

import "time"

type Product struct {
	Id          int32  `pg:",pk,unique"`
	Uuid        string `pg:"type:uuid,unique,default:uuid_generate_v4()"`
	ProductName string
	OwnerId     int32
	CreatedAt   time.Time `pg:"default:now()"`
	// relationships
	Owner *Producer `pg:"rel:has-one,join_fk:owner_id"`
}
