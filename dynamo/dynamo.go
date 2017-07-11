package dynamo

import (
	"strconv"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

const (
	UserId                              = "userId"
	ItemId                              = "itemId"
	Email                               = "email"
	ItemDoesnotExistConditionExpression = "attribute_not_exists(" + ItemId + ")"
	ItemExistsConditionExpression       = "attribute_exists(" + ItemId + ")"
	ConditionalCheckFailedException     = "ConditionalCheckFailedException"
	ValidationException                 = "ValidationException"
)

func PutItemParams(tableName string, values map[string]*dynamodb.AttributeValue) *dynamodb.PutItemInput {
	params := &dynamodb.PutItemInput{
		Item:                values,
		TableName:           aws.String(tableName), // Required
		ConditionExpression: aws.String(ItemDoesnotExistConditionExpression),
	}

	return params
}

//This doesnot have ConditionExpression , so this update the item if it already exists
func PutOrUpdateItemParams(tableName string, values map[string]*dynamodb.AttributeValue) *dynamodb.PutItemInput {
	params := &dynamodb.PutItemInput{
		Item:      values,
		TableName: aws.String(tableName), // Required
	}

	return params
}

func GetItemParamsWithUserIdAndNumberItemId(tableName string, userId string, itemId int64) *dynamodb.GetItemInput {

	params := &dynamodb.GetItemInput{
		Key: map[string]*dynamodb.AttributeValue{
			UserId: {
				S: aws.String(userId),
			},
			ItemId: {
				N: aws.String(strconv.FormatInt(itemId, 10)),
			},
		},
		TableName:      aws.String(tableName),
		ConsistentRead: aws.Bool(true),
	}

	return params
}

func GetItemParamsWithUserIdAndStringItemId(tableName string, userId string, itemId string) *dynamodb.GetItemInput {

	params := &dynamodb.GetItemInput{
		Key: map[string]*dynamodb.AttributeValue{
			UserId: {
				S: aws.String(userId),
			},
			ItemId: {
				S: aws.String(itemId),
			},
		},
		TableName:      aws.String(tableName),
		ConsistentRead: aws.Bool(true),
	}

	return params
}

func GetItemParamsForEmailPrimarykey(tableName string, email string) *dynamodb.GetItemInput {
	params := &dynamodb.GetItemInput{
		Key: map[string]*dynamodb.AttributeValue{
			Email: {
				S: aws.String(email),
			},
		},
		TableName:      aws.String(tableName),
		ConsistentRead: aws.Bool(true),
	}

	return params
}

func GetItemParamsForUserIdPrimarykey(tableName string, userId string) *dynamodb.GetItemInput {
	params := &dynamodb.GetItemInput{
		Key: map[string]*dynamodb.AttributeValue{
			UserId: {
				S: aws.String(userId),
			},
		},
		TableName:      aws.String(tableName),
		ConsistentRead: aws.Bool(true),
	}

	return params
}

func QueryParamsUserIdWithLimit(tableName string, userId string, limit int64) *dynamodb.QueryInput {

	params := &dynamodb.QueryInput{
		TableName:        aws.String(tableName), // Required
		ConsistentRead:   aws.Bool(true),
		Limit:            aws.Int64(limit),
		ScanIndexForward: aws.Bool(false),
		KeyConditions: map[string]*dynamodb.Condition{
			UserId: { // Required
				ComparisonOperator: aws.String(dynamodb.ComparisonOperatorEq), // Required
				AttributeValueList: []*dynamodb.AttributeValue{
					{
						S: aws.String(userId),
					},
				},
			},
		},
	}
	return params
}

func QueryParamsForUserIdIndexName(tableName string, indexName string, userId string) *dynamodb.QueryInput {

	params := &dynamodb.QueryInput{
		TableName:        aws.String(tableName), // Required
		IndexName:        aws.String(indexName),
		Limit:            aws.Int64(1),
		ScanIndexForward: aws.Bool(false),
		KeyConditions: map[string]*dynamodb.Condition{
			UserId: { // Required
				ComparisonOperator: aws.String(dynamodb.ComparisonOperatorEq), // Required
				AttributeValueList: []*dynamodb.AttributeValue{
					{
						S: aws.String(userId),
					},
				},
			},
		},
	}
	return params
}

func QueryParamsForUserIdStartKeyAndLimit(tableName string, userId string, itemId int64, limit int64) *dynamodb.QueryInput {
	//FTODO: test this how it works, and is it efficient check cosumed Capacity also
	params := &dynamodb.QueryInput{
		TableName:        aws.String(tableName), // Required
		ConsistentRead:   aws.Bool(true),
		Limit:            aws.Int64(limit),
		ScanIndexForward: aws.Bool(false),
		KeyConditions: map[string]*dynamodb.Condition{
			UserId: { // Required
				ComparisonOperator: aws.String(dynamodb.ComparisonOperatorEq), // Required
				AttributeValueList: []*dynamodb.AttributeValue{
					{
						S: aws.String(userId),
					},
				},
			},
		},
		ExclusiveStartKey: map[string]*dynamodb.AttributeValue{
			UserId: {
				S: aws.String(userId),
			},
			ItemId: {
				N: aws.String(strconv.FormatInt(itemId, 10)),
			},
		},
	}
	return params
}

func QueryParamsUserIdStartKeyLimitAndIndexName(tableName string, indexName string, userId string, columnName string, itemId int64, limit int64) *dynamodb.QueryInput {

	params := &dynamodb.QueryInput{
		TableName:        aws.String(tableName), // Required
		IndexName:        aws.String(indexName),
		Limit:            aws.Int64(limit),
		ScanIndexForward: aws.Bool(false),
		KeyConditions: map[string]*dynamodb.Condition{
			UserId: { // Required
				ComparisonOperator: aws.String(dynamodb.ComparisonOperatorEq), // Required
				AttributeValueList: []*dynamodb.AttributeValue{
					{
						S: aws.String(userId),
					},
				},
			},
			columnName: {
				ComparisonOperator: aws.String(dynamodb.ComparisonOperatorGe), // Required
				AttributeValueList: []*dynamodb.AttributeValue{
					{
						N: aws.String(strconv.FormatInt(itemId, 10)),
					},
				},
			},
		},
	}
	return params
}
func QueryParamsUserLimitAndIndexName(tableName string, indexName string, userId string, limit int64) *dynamodb.QueryInput {

	params := &dynamodb.QueryInput{
		TableName:        aws.String(tableName), // Required
		IndexName:        aws.String(indexName),
		Limit:            aws.Int64(limit),
		ScanIndexForward: aws.Bool(false),
		KeyConditions: map[string]*dynamodb.Condition{
			UserId: { // Required
				ComparisonOperator: aws.String(dynamodb.ComparisonOperatorEq), // Required
				AttributeValueList: []*dynamodb.AttributeValue{
					{
						S: aws.String(userId),
					},
				},
			},
		},
	}
	return params
}

func DeleteParamsForUserIdAndNumberItemId(tableName string, userId string, itemId int64) *dynamodb.DeleteItemInput {
	params := &dynamodb.DeleteItemInput{
		Key: map[string]*dynamodb.AttributeValue{
			UserId: {
				S: aws.String(userId),
			},
			ItemId: {
				N: aws.String(strconv.FormatInt(itemId, 10)),
			},
		},
		TableName:    aws.String(tableName),
		ReturnValues: aws.String("ALL_OLD"),
	}

	return params
}

func UpdateStringParamsUserIdAndStringItemId(tableName string, userId string, itemId string, columnName string, columnValue string) *dynamodb.UpdateItemInput {
	updateKey := ":updateKey"
	params := &dynamodb.UpdateItemInput{
		Key: map[string]*dynamodb.AttributeValue{
			UserId: {
				S: aws.String(userId),
			},
			ItemId: {
				S: aws.String(itemId),
			},
		},
		TableName:    aws.String(tableName),
		ReturnValues: aws.String("ALL_NEW"),
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			updateKey: {
				S: aws.String(columnValue),
			},
		},
		UpdateExpression:    aws.String("SET " + columnName + " = " + updateKey),
		ConditionExpression: aws.String(ItemExistsConditionExpression),
	}
	return params
}

func UpdateBooleanParamsUserIdAndStringItemId(tableName string, userId string, itemId string, columnName string, columnValue bool) *dynamodb.UpdateItemInput {
	updateKey := ":updateKey"
	params := &dynamodb.UpdateItemInput{
		Key: map[string]*dynamodb.AttributeValue{
			UserId: {
				S: aws.String(userId),
			},
			ItemId: {
				S: aws.String(itemId),
			},
		},
		TableName:    aws.String(tableName),
		ReturnValues: aws.String("ALL_NEW"),
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			updateKey: {
				BOOL: aws.Bool(columnValue),
			},
		},
		UpdateExpression:    aws.String("SET " + columnName + " = " + updateKey),
		ConditionExpression: aws.String(ItemExistsConditionExpression),
	}
	return params
}

func DeleteParamsForUserIdAndStringItemId(tableName string, userId string, itemId string) *dynamodb.DeleteItemInput {
	params := &dynamodb.DeleteItemInput{
		Key: map[string]*dynamodb.AttributeValue{
			UserId: {
				S: aws.String(userId),
			},
			ItemId: {
				S: aws.String(itemId),
			},
		},
		TableName:    aws.String(tableName),
		ReturnValues: aws.String("ALL_OLD"),
	}

	return params
}
func DeleteParamsForUserIdPrimaryKey(tableName string, primaryKey string) *dynamodb.DeleteItemInput {
	params := &dynamodb.DeleteItemInput{
		Key: map[string]*dynamodb.AttributeValue{
			UserId: {
				S: aws.String(primaryKey),
			},
		},
		TableName:    aws.String(tableName),
		ReturnValues: aws.String("ALL_OLD"),
	}

	return params
}

func DeleteParamsForEmailPrimaryKey(tableName string, primaryKey string) *dynamodb.DeleteItemInput {
	params := &dynamodb.DeleteItemInput{
		Key: map[string]*dynamodb.AttributeValue{
			Email: {
				S: aws.String(primaryKey),
			},
		},
		TableName:    aws.String(tableName),
		ReturnValues: aws.String("ALL_OLD"),
	}

	return params
}

//session_id
func DeleteParamsForStringPrimaryKey(tableName string, primaryKeyName string, primaryKeyValue string) *dynamodb.DeleteItemInput {
	params := &dynamodb.DeleteItemInput{
		Key: map[string]*dynamodb.AttributeValue{
			primaryKeyName: {
				S: aws.String(primaryKeyValue),
			},
		},
		TableName:    aws.String(tableName),
		ReturnValues: aws.String("ALL_OLD"),
	}

	return params
}

func GetItemParamsForStringPrimaryKey(tableName string, primaryKeyName string, primaryKeyValue string) *dynamodb.GetItemInput {

	params := &dynamodb.GetItemInput{
		Key: map[string]*dynamodb.AttributeValue{
			primaryKeyName: {
				S: aws.String(primaryKeyValue),
			},
		},
		TableName:      aws.String(tableName),
		ConsistentRead: aws.Bool(true),
	}

	return params
}
