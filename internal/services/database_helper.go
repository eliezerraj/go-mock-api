package services

import (
	"context"
	"time"
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"

	"github.com/go-mock-api/internal/utils/constants"
	"github.com/go-mock-api/internal/utils/loggers"
	"github.com/go-mock-api/internal/core/model"

)

type DatabaseHelper interface {
	GetConnection(ctx context.Context) (*sql.DB, error)
	CloseConnection()
}

type DatabaseHelperImpl struct {
	client   	*sql.DB
	timeout		time.Duration
}

func NewDatabaseHelper(databaseRDS model.DatabaseRDS) (DatabaseHelper, error) {
	loggers.GetLogger().Named(constants.Service).Info("DatabaseHelper") 

	_ , cancel := context.WithTimeout(context.Background(), time.Duration(databaseRDS.Db_timeout)*time.Second)
	defer cancel()

	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", databaseRDS.User, databaseRDS.Password, databaseRDS.Host, databaseRDS.Port, databaseRDS.DatabaseName) 
	
	fmt.Println("==========>", databaseRDS.Postgres_Driver, connStr)

	client, err := sql.Open(databaseRDS.Postgres_Driver, connStr)
	if err != nil {
		return DatabaseHelperImpl{}, err
	}
	err = client.Ping()
	if err != nil {
		return DatabaseHelperImpl{}, err
	}

	return DatabaseHelperImpl{
		client: client,
		timeout:  time.Duration(databaseRDS.Db_timeout) * time.Second,
	}, nil
}

func (d DatabaseHelperImpl) GetConnection(ctx context.Context) (*sql.DB, error) {
	loggers.GetLogger().Named(constants.Service).Info("GetConnection") 
	return d.client, nil
}

func (d DatabaseHelperImpl) CloseConnection()  {
	loggers.GetLogger().Named(constants.Service).Info("CloseConnection !!!!") 
	defer d.client.Close()
}
