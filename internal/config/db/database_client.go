package db

import "context"

type ConfigDatabaseClient interface {
	GetDocFromCollection(ctx context.Context, colName string, docName string) (interface{}, error)
	Close() error
}
