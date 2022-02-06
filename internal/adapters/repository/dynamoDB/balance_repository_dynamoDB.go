package dynamoDB

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
	"github.com/aws/aws-sdk-go/service/dynamodb/expression"

	"github.com/go-mock-api/internal/utils/constants"
	"github.com/go-mock-api/internal/utils/loggers"
	"github.com/go-mock-api/internal/core/model"
	"github.com/go-mock-api/internal/exceptions"
	"github.com/go-mock-api/internal/adapters/repository"
	"github.com/go-mock-api/internal/viper"
)

type BalanceRepositoryDynamoDB interface {
	FindById(ctx context.Context, pk string) (model.Balance, error)
	List(ctx context.Context) ([]model.Balance, error)
	ListById(ctx context.Context, balance model.Balance) ([]model.Balance, error)
	Save(ctx context.Context, balance model.Balance) (model.Balance, error)
}

type BalanceRepositoryDynamoDBImpl struct {
	client dynamodbiface.DynamoDBAPI
	table_name  *string
}

func NewBalanceRepositoryDynamoDB(table_name string) (repository.BalanceRepository, error) {
	loggers.GetLogger().Named(constants.Database).Info("NewBalanceRepositoryDynamoDB") 
	
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(viper.Application.AwsEnv.Aws_region),
		Credentials: credentials.NewStaticCredentials(viper.Application.AwsEnv.Aws_access_id, viper.Application.AwsEnv.Aws_access_secret , ""),},
	)
	if err != nil {
		return BalanceRepositoryDynamoDBImpl{}, exceptions.Throw(exceptions.ErrOpenDataBase, err)
	}

	return BalanceRepositoryDynamoDBImpl{
		client: dynamodb.New(sess),
		table_name: aws.String(table_name),
	}, nil
}

func (b BalanceRepositoryDynamoDBImpl) Save(ctx context.Context, balance model.Balance) (model.Balance, error) {
	loggers.GetLogger().Named(constants.Database).Info("Save-----") 
	
	item, err := dynamodbattribute.MarshalMap(balance)
	if err != nil {
		return model.Balance{}, exceptions.Throw(exceptions.ErrSaveDatabase, err)
	}

	transactItems := []*dynamodb.TransactWriteItem{}
	transactItems = append(transactItems, &dynamodb.TransactWriteItem{Put: &dynamodb.Put{
		TableName: b.table_name,
		Item:      item,
	}})

	transaction := &dynamodb.TransactWriteItemsInput{TransactItems: transactItems}
	if err := transaction.Validate(); err != nil {
		return model.Balance{}, exceptions.Throw(err, exceptions.ErrSaveDatabase)
	}

	_, err = b.client.TransactWriteItemsWithContext(ctx, transaction)
	if err != nil {
		return model.Balance{}, exceptions.Throw( exceptions.ErrSaveDatabase, err)
	}

	return balance , nil
}

func (b BalanceRepositoryDynamoDBImpl) List(ctx context.Context) ([]model.Balance, error) {
	loggers.GetLogger().Named(constants.Database).Info("List") 
	return []model.Balance{}, nil
}

func (b BalanceRepositoryDynamoDBImpl) ListById(ctx context.Context, balance model.Balance) ([]model.Balance, error) {
	loggers.GetLogger().Named(constants.Database).Info("List") 

	balance_id := balance.Id
	account := balance.Account

	var keyCond expression.KeyConditionBuilder

	keyCond = expression.KeyAnd(
		expression.Key("balance_id").Equal(expression.Value(balance_id)),
		expression.Key("account").BeginsWith(account),
	)
	expr, err := expression.NewBuilder().
							WithKeyCondition(keyCond).
							Build()
	if err != nil {
		return []model.Balance{}, exceptions.Throw(exceptions.ErrNoDataFound, err)
	}

	key := &dynamodb.QueryInput{
		TableName:                 b.table_name,
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		KeyConditionExpression:    expr.KeyCondition(),
	}

	fmt.Println("key => ", key)

	result, err := b.client.QueryWithContext(ctx, key)
	if err != nil {
		return []model.Balance{}, exceptions.Throw(exceptions.ErrNoDataFound, err)
	}

	fmt.Println("result => ", result)

	balances := []model.Balance{}
	err = dynamodbattribute.UnmarshalListOfMaps(result.Items, &balances)
    if err != nil {
		fmt.Println("err => ", err)
		return []model.Balance{}, exceptions.Throw(exceptions.ErrJsonCode, err)
    }

	if len(balances) == 0 {
		return []model.Balance{}, exceptions.Throw(exceptions.ErrNoDataFound, exceptions.ErrNoDataFound)
	} else {
		return balances, nil
	}
}

func (b BalanceRepositoryDynamoDBImpl) FindById(ctx context.Context, pk string) (model.Balance, error) {
	loggers.GetLogger().Named(constants.Database).Info("FindById") 

	balance_id := pk

	var keyCond expression.KeyConditionBuilder
	keyCond = expression.Key("balance_id").Equal(expression.Value(balance_id))

	expr, err := expression.NewBuilder().
							WithKeyCondition(keyCond).
							Build()
	if err != nil {
		return model.Balance{}, exceptions.Throw(exceptions.ErrNoDataFound, err)
	}

	key := &dynamodb.QueryInput{
		TableName:                 b.table_name,
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		KeyConditionExpression:    expr.KeyCondition(),
	}

	fmt.Println("key => ", key)

	result, err := b.client.QueryWithContext(ctx, key)
	if err != nil {
		return model.Balance{}, exceptions.Throw(exceptions.ErrNoDataFound, err)
	}

	fmt.Println("result => ", result)

	balances := []model.Balance{}
	err = dynamodbattribute.UnmarshalListOfMaps(result.Items, &balances)
    if err != nil {
		fmt.Println("err => ", err)
		return model.Balance{}, exceptions.Throw(exceptions.ErrJsonCode, err)
    }

	if len(balances) == 0 {
		return model.Balance{}, exceptions.Throw(exceptions.ErrNoDataFound, exceptions.ErrNoDataFound)
	} else {
		return balances[0], nil
	}
}
