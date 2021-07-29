package dao

import (
	"context"
	"database/sql"
	"github.com/doug-martin/goqu/v9"
	"github.com/jmoiron/sqlx"
	"test_task/internal/entity"
	"test_task/internal/logger"
)

type Entity1Dao struct {
	connect *sqlx.DB
}

const entity1Table = "entity1"

func GetEntity1Dao(db *sqlx.DB) *Entity1Dao {
	return &Entity1Dao{
		connect: db,
	}
}


func (u *Entity1Dao) GetByID(entityId int) (*entity.Entity1, error) {
	ctx := context.Background()
	con, err := u.connect.Connx(ctx)
	defer con.Close()

	if err != nil {
		logger.Logger.Warn(err)
		return nil, err
	}

	user := &entity.Entity1{}
	sqlString, _, _ := goqu.From(entity1Table).Where(goqu.C("id_entity_1").Eq(entityId)).ToSQL()
	err = con.GetContext(ctx, user, sqlString)

	if err != nil {
		if err != sql.ErrNoRows {
			logger.Logger.Warn(err)
		}
		return nil, err
	}

	return user, nil

}

func (u *Entity1Dao) GetAll( ) ([]*entity.Entity1, error) {
	ctx := context.Background()
	con, err := u.connect.Connx(ctx)
	defer con.Close()

	if err != nil {
		logger.Logger.Warn(err)
		return nil, err
	}

	var entities []*entity.Entity1
	sqlString, _, _ := goqu.From(entity1Table).ToSQL()
	err = con.SelectContext(ctx, &entities, sqlString)

	if err != nil {
		if err != sql.ErrNoRows {
			logger.Logger.Warn(err)
		}
		return nil, err
	}

	return entities, nil

}

func (u *Entity1Dao) Update(entity1 *entity.Entity1) error {
	ctx := context.Background()
	con, err := u.connect.Connx(ctx)
	defer con.Close()

	if err != nil {
		logger.Logger.Warn(err)
		return err
	}

	sql, _, _ := goqu.Update(entity1Table).Set(
		goqu.Record{
			"name": entity1.Name,
		}).Where(goqu.C("id_entity_1").Eq(entity1.Id)).ToSQL()
	_, err = con.ExecContext(ctx, sql)

	if err != nil {
		logger.Logger.Warn(err)
		return err
	}

	return nil
}

func (u *Entity1Dao) Save(entity1 *entity.Entity1) error {
	ctx := context.Background()
	con, err := u.connect.Connx(ctx)
	defer con.Close()

	if err != nil {
		logger.Logger.Warn(err)
		return err
	}

	sql, _, _ := goqu.Insert(entity1Table).Rows(entity1).ToSQL()
	_, err = con.ExecContext(ctx, sql)

	if err != nil {
		logger.Logger.Warn(err)
		return err
	}

	return nil
}