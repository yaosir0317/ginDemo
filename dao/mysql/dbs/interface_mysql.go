package dbs

import (
	"context"
	"database/sql"
	dals "ginDemo/dao/mysql"

	"github.com/gocraft/dbr/v2"
)

type Interface1 interface {
	GetCount(ctx context.Context) (uint64, error)
}

type Interface2 interface {
	GetRow(ctx context.Context)
}

type Interface3 interface {
	GetArticle(ctx context.Context, articleId string) (*PArticle, error)
	GetUser(ctx context.Context, userId, phone string) (*User, error)
	GetTitleMap(ctx context.Context, articleIds []string) (map[string]string, error)
	GetArticleCategoryMap(ctx context.Context, articleIds []string) (map[string]string, error)
	GetNewUsers(ctx context.Context, date string) ([]string, error)
}

type Interface interface {
	Interface1
	Interface2
	Interface3
}

type impl struct {
	dbr     *dbr.Connection
	dialect dbr.Dialect
}

var _ Interface1 = (*impl)(nil)

var _ Interface2 = (*impl)(nil)

var _ Interface3 = (*impl)(nil)

func NewMysqlInterface(db *sql.DB) Interface {
	connection := dbr.Connection{
		DB:            db,
		Dialect:       dals.MysqlDialect,
		EventReceiver: &dbr.NullEventReceiver{},
	}
	return &impl{
		dbr:     &connection,
		dialect: dals.MysqlDialect,
	}
}
