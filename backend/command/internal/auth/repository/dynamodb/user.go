package dynamodb

import (
	"context"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"

	"github.com/oka311119/l4-app/backend/command/internal/domain/entity"
)

type User struct {
	ID       string `dynamodbav:"ID"`
	Username string `dynamodbav:"Username"`
	Password string `dynamodbav:"Password"`
}

type UserRepository struct {
	db *dynamodb.DynamoDB
	tableName string
}

func NewUserRepository(ddb *dynamodb.DynamoDB, tableName string) *UserRepository {
	return &UserRepository{
		db: ddb,
		tableName: tableName,
	}
}

func (r UserRepository) CreateUser(ctx context.Context, user *entity.User) error {
	av, err := dynamodbattribute.MarshalMap(user)
	if err != nil {
		return err
	}

	input := &dynamodb.PutItemInput{
		Item:      av,
		TableName: aws.String(r.tableName),
	}

	_, err = r.db.PutItemWithContext(ctx, input)

	return err
}

func (r UserRepository) GetUser(ctx context.Context, username, password string) (*entity.User, error) {
	input := &dynamodb.GetItemInput{
		TableName: aws.String(r.tableName),
		Key: map[string]*dynamodb.AttributeValue{
			"Username": {
				S: aws.String(username),
			},
			"Password": {
				S: aws.String(password),
			},
		},
	}

	result, err := r.db.GetItemWithContext(ctx, input)
	if err != nil {
		return nil, err
	}

	user := new(User)
	err = dynamodbattribute.UnmarshalMap(result.Item, user)
	if err != nil {
		return nil, err
	}

	// TODO: 確認
	entityUser := &entity.User{
		Username: user.Username,
		Password: user.Password,
	}
	return entityUser, nil
}