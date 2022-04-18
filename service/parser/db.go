package parser

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/sqlitedialect"
	"strings"
)

type IPriceRepo interface {
	Adds(ctx context.Context, prices ...Price) error
}

type PriceRepo struct {
	db *bun.DB
}

func NewBunDB(cfg Config) (*bun.DB, error) {
	dataSource := fmt.Sprintf(cfg.Postgresql.Format,
		cfg.Postgresql.Username,
		cfg.Postgresql.Password,
		cfg.Postgresql.Address,
		cfg.Postgresql.Database,
	)
	db, err := sql.Open(cfg.Postgresql.DriverName, dataSource)
	if err != nil {
		return nil, err
	}
	bunDB := bun.NewDB(db, sqlitedialect.New())

	//bunDB.AddQueryHook(bundebug.NewQueryHook(
	//	bundebug.WithVerbose(true),
	//	bundebug.FromEnv("BUNDEBUG"),
	//))

	return bunDB, nil
}

func NewPriceRepo(db *bun.DB) *PriceRepo {
	return &PriceRepo{
		db: db,
	}
}

func (o *PriceRepo) Adds(ctx context.Context, prices ...Price) error {
	tableName:= strings.ToLower(prices[0].Symbol)
	if _, err := o.db.NewCreateTable().IfNotExists().Model((*Price)(nil)).Table(tableName).ModelTableExpr(tableName).Exec(ctx); err != nil {
		return err
	}

	_, err := o.db.NewInsert().Model(&prices).ModelTableExpr(tableName).Exec(ctx)
	return err
}
