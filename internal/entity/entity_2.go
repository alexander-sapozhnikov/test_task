package entity

type Entity2 struct {
	Id   int    `db:"id_entity_2" goqu:"skipinsert" json:"id"`
	Name string `db:"name" json:"name"`
}
