package entity

type Entity1 struct {
	Id   int    `db:"id_entity_1" goqu:"skipinsert" json:"id"`
	Name string `db:"name" json:"name"`
}
