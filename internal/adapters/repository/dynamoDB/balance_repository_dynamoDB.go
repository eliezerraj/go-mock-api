package dynamoDB

import (
	"context"
	//"fmt"
	//	"encoding/json"
	//	"go.uber.org/zap"

	"github.com/go-mock-api/internal/utils/constants"
	"github.com/go-mock-api/internal/utils/loggers"
	"github.com/go-mock-api/internal/core/model"
	"github.com/go-mock-api/internal/exceptions"
	"github.com/go-mock-api/internal/adapters/repository"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

type BalanceRepositoryDynamoDB interface {
	FindById(ctx context.Context, id string) (model.Balance, error)
	List(ctx context.Context) ([]model.Balance, error)
	Save(ctx context.Context, balance model.Balance) (model.Balance, error)
}

type BalanceRepositoryDynamoDBImpl struct {
	client *dynamodb.DynamoDB
	table_name  *string
}

func NewBalanceRepositoryDynamoDB(table_name *string) (repository.BalanceRepository, error) {
	
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("aws_region"),
		Credentials: credentials.NewStaticCredentials( "aws_access_id" ,"aws_access_secret" , ""),},
	)
	if err != nil {
		return BalanceRepositoryDynamoDBImpl{}, exceptions.Throw(err, exceptions.ErrOpenDataBase)
	}
	
	return BalanceRepositoryDynamoDBImpl{
		client: dynamodb.New(sess),
		table_name: table_name,
	}, nil
}

func (b BalanceRepositoryDynamoDBImpl) Save(ctx context.Context, balance model.Balance) (model.Balance, error) {
	loggers.GetLogger().Named(constants.Database).Info("Save-----") 
	
	item, err := dynamodbattribute.MarshalMap(balance)
	if err != nil {
		return model.Balance{}, exceptions.Throw(err, exceptions.ErrSaveDatabase)
	}
	loggers.GetLogger().Named(constants.Database).Info("Save-----1") 
	transactItems := []*dynamodb.TransactWriteItem{}
	transactItems = append(transactItems, &dynamodb.TransactWriteItem{Put: &dynamodb.Put{
		TableName: b.table_name,
		Item:      item,
	}})
	loggers.GetLogger().Named(constants.Database).Info("Save-----2") 
	transaction := &dynamodb.TransactWriteItemsInput{TransactItems: transactItems}
	if err := transaction.Validate(); err != nil {
		return model.Balance{}, exceptions.Throw(err, exceptions.ErrSaveDatabase)
	}
	loggers.GetLogger().Named(constants.Database).Info("Save-----3") 
	_, err = b.client.TransactWriteItemsWithContext(ctx, transaction)
	loggers.GetLogger().Named(constants.Database).Info("Save-----4") 
	return model.Balance{} , nil
}

func (b BalanceRepositoryDynamoDBImpl) List(ctx context.Context) ([]model.Balance, error) {
	loggers.GetLogger().Named(constants.Database).Info("List") 
	var result []model.Balance
	return result, nil
}

func (b BalanceRepositoryDynamoDBImpl) FindById(ctx context.Context, id string) (model.Balance, error) {
	loggers.GetLogger().Named(constants.Database).Info("FindById") 

	return model.Balance{}, exceptions.Throw( exceptions.ErrNoDataFound, exceptions.ErrNoDataFound)
}
