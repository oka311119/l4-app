package dynamodb

import (
	"github.com/oka311119/l4-app/backend/command/internal/model"
	"github.com/oka311119/l4-app/backend/command/internal/storage"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

type ItemRepository struct {
	db *dynamodb.DynamoDB
}

func NewItemRepository() storage.ItemRepository {
	sess := session.Must(session.NewSession(&aws.Config{
		Region: aws.String("us-west-2"),
	}))
	db := dynamodb.New(sess)
	return &ItemRepository{db: db}
}

func (r *ItemRepository) CreateItem(item *model.Item) error {
	av, err := dynamodbattribute.MarshalMap(item)
	if err != nil {
		return err
	}

	input := &dynamodb.PutItemInput{
		Item:      av,
		TableName: aws.String("Items"),
	}

	_, err = r.db.PutItem(input)
	return err
}

func (r *ItemRepository) UpdateItem(item *model.Item) error {
	av, err := dynamodbattribute.MarshalMap(item)
	if err != nil {
		return err
	}

	input := &dynamodb.UpdateItemInput{
		ExpressionAttributeValues: av,
		TableName: aws.String("Items"),
		Key: map[string]*dynamodb.AttributeValue{
			"id": {
				S: aws.String(item.ID.String()),
			},
		},
		ReturnValues:     aws.String("UPDATED_NEW"),
		UpdateExpression: aws.String("set name = :n"),
	}

	_, err = r.db.UpdateItem(input)
	return err
}