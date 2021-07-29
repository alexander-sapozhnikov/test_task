package dao

import (
	"context"
	"database/sql"
	"github.com/doug-martin/goqu/v9"
	"github.com/jmoiron/sqlx"
	"test_task/internal/entity"
	"test_task/internal/logger"
)

type Entity2Dao struct {
	connect *sqlx.DB
}

const entity2Table = "entity2"

func GetEntity2Dao(db *sqlx.DB) *Entity2Dao {
	return &Entity2Dao{
		connect: db,
	}
}


func (u *Entity2Dao) GetByID(entityId int) (*entity.Entity2, error) {
	ctx := context.Background()
	con, err := u.connect.Connx(ctx)
	defer con.Close()

	if err != nil {
		logger.Logger.Warn(err)
		return nil, err
	}

	user := &entity.Entity2{}
	sqlString, _, _ := goqu.From(entity2Table).Where(goqu.C("id_entity_2").Eq(entityId)).ToSQL()
	err = con.GetContext(ctx, user, sqlString)

	if err != nil {
		if err != sql.ErrNoRows {
			logger.Logger.Warn(err)
		}
		return nil, err
	}

	return user, nil

}

func (u *Entity2Dao) GetAll( ) ([]*entity.Entity2, error) {
	ctx := context.Background()
	con, err := u.connect.Connx(ctx)
	defer con.Close()

	if err != nil {
		logger.Logger.Warn(err)
		return nil, err
	}

	var entities []*entity.Entity2
	sqlString, _, _ := goqu.From(entity2Table).ToSQL()
	err = con.SelectContext(ctx, &entities, sqlString)

	if err != nil {
		if err != sql.ErrNoRows {
			logger.Logger.Warn(err)
		}
		return nil, err
	}

	return entities, nil

}

func (u *Entity2Dao) Update(entity2 *entity.Entity2)  error {
	ctx := context.Background()
	con, err := u.connect.Connx(ctx)
	defer con.Close()

	if err != nil {
		logger.Logger.Warn(err)
		return err
	}

	sql, _, _ := goqu.Update(entity2Table).Set(
		goqu.Record{
			"name": entity2.Name,
		}).Where(goqu.C("id_entity_2").Eq(entity2.Id)).ToSQL()
	_, err = con.ExecContext(ctx, sql)

	if err != nil {
		logger.Logger.Warn(err)
		return err
	}

	return nil
}

func (u *Entity2Dao) Save(entity2 *entity.Entity2) error {
	ctx := context.Background()
	con, err := u.connect.Connx(ctx)
	defer con.Close()

	if err != nil {
		logger.Logger.Warn(err)
		return err
	}

	sql, _, _ := goqu.Insert(entity2Table).Rows(entity2).ToSQL()
	_, err = con.ExecContext(ctx, sql)

	if err != nil {
		logger.Logger.Warn(err)
		return err
	}

	return nil
}