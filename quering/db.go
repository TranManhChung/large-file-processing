package quering

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/TranManhChung/large-file-processing/parser"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/sqlitedialect"
)

type IPriceRepo interface {
	GetByOffsetLimit(ctx context.Context, symbol string, offset, limit int) ([]parser.Price, error)
	GetByDuration(ctx context.Context, symbol string, start, end int64) ([]parser.Price, error)
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

func (o *PriceRepo) GetByOffsetLimit(ctx context.Context, symbol string, offset, limit int) ([]parser.Price, error) {
	var prices []parser.Price
	err := o.db.NewSelect().Model(&prices).ColumnExpr(symbol + ".*").ModelTableExpr(symbol).Limit(limit).Offset(offset).Scan(ctx)
	return prices, err
}

func (o *PriceRepo) GetByDuration(ctx context.Context, symbol string, start, end int64) ([]parser.Price, error) {
	var prices []parser.Price
	err := o.db.NewSelect().Model(&prices).ColumnExpr(symbol + ".*").ModelTableExpr(symbol).Where("id >= ? and id <= ?", start, end).Scan(ctx)
	return prices, err
}
