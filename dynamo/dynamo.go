package dynamo

import (
	"strconv"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

const (
	StringKeyType                   = "S"
	NumberKeyType                   = "N"
	ConditionalCheckFailedException = "ConditionalCheckFailedException"
	ValidationException             = "ValidationException"
)

type Dynamo struct {
	TableName    string
	KeyName      string
	RangeKeyName string
	RangeKeyType string
	Limit        int64
}

func (dynamo *Dynamo) PutItemParams(values map[string]*dynamodb.AttributeValue) *dynamodb.PutItemInput {
	if len(dynamo.RangeKeyName) == 0 {
		return &dynamodb.PutItemInput{
			Item:      values,
			TableName: aws.String(dynamo.TableName),
		}
	} else {
		return &dynamodb.PutItemInput{
			Item:                values,
			TableName:           aws.String(dynamo.TableName), // Required
			ConditionExpression: aws.String("attribute_not_exists(" + dynamo.RangeKeyName + ")"),
		}
	}
}

func (dynamo *Dynamo) UpdateItemParams(values map[string]*dynamodb.AttributeValue) *dynamodb.PutItemInput {
	return &dynamodb.PutItemInput{
		Item:                values,
		TableName:           aws.String(dynamo.TableName), // Required
		ConditionExpression: aws.String("attribute_exists(" + dynamo.RangeKeyName + ")"),
	}
}

func (dynamo *Dynamo) GetItemParams(key string, rangeKey interface{}) *dynamodb.GetItemInput {
	return &dynamodb.GetItemInput{
		Key:            dynamo.GetKeyAttribute(key, rangeKey),
		TableName:      aws.String(dynamo.TableName),
		ConsistentRead: aws.Bool(true),
	}
}

func (dynamo *Dynamo) DeleteParams(key string, rangeKey interface{}) *dynamodb.DeleteItemInput {
	return &dynamodb.DeleteItemInput{
		Key:          dynamo.GetKeyAttribute(key, rangeKey),
		TableName:    aws.String(dynamo.TableName),
		ReturnValues: aws.String("ALL_OLD"),
	}
}

func (dynamo *Dynamo) GetKeyAttribute(key string, rangeKey interface{}) map[string]*dynamodb.AttributeValue {
	//TODO: inproper interface conversion will throw nil exceptions avoid them
	if dynamo.RangeKeyType == NumberKeyType {
		return map[string]*dynamodb.AttributeValue{
			dynamo.KeyName: {
				S: aws.String(key),
			},
			dynamo.RangeKeyName: {
				N: aws.String(strconv.Itoa(rangeKey.(int))),
			},
		}
	} else if dynamo.RangeKeyType == StringKeyType {
		return map[string]*dynamodb.AttributeValue{
			dynamo.KeyName: {
				S: aws.String(key),
			},
			dynamo.RangeKeyName: {
				S: aws.String(rangeKey.(string)),
			},
		}
	} else {
		return map[string]*dynamodb.AttributeValue{
			dynamo.KeyName: {
				S: aws.String(key),
			},
		}
	}
}

func (dynamo *Dynamo) QueryParams(key string) *dynamodb.QueryInput {

	params := &dynamodb.QueryInput{
		TableName:        aws.String(dynamo.TableName), // Required
		ConsistentRead:   aws.Bool(true),
		Limit:            aws.Int64(dynamo.Limit),
		ScanIndexForward: aws.Bool(false),
		KeyConditions: map[string]*dynamodb.Condition{
			dynamo.KeyName: { // Required
				ComparisonOperator: aws.String(dynamodb.ComparisonOperatorEq), // Required
				AttributeValueList: []*dynamodb.AttributeValue{
					{
						S: aws.String(key),
					},
				},
			},
		},
	}
	return params
}

func (dynamo *Dynamo) QueryParamsWithIndexName(indexName string, key string) *dynamodb.QueryInput {

	params := &dynamodb.QueryInput{
		TableName:        aws.String(dynamo.TableName), // Required
		IndexName:        aws.String(indexName),
		Limit:            aws.Int64(dynamo.Limit),
		ScanIndexForward: aws.Bool(false),
		KeyConditions: map[string]*dynamodb.Condition{
			dynamo.KeyName: { // Required
				ComparisonOperator: aws.String(dynamodb.ComparisonOperatorEq), // Required
				AttributeValueList: []*dynamodb.AttributeValue{
					{
						S: aws.String(key),
					},
				},
			},
		},
	}
	return params
}

func (dynamo *Dynamo) UpdateStringParams(key string, rangeKey interface{}, columnName string, columnValue string) *dynamodb.UpdateItemInput {
	updateKey := ":updateKey"
	return &dynamodb.UpdateItemInput{
		Key:          dynamo.GetKeyAttribute(key, rangeKey),
		TableName:    aws.String(dynamo.TableName),
		ReturnValues: aws.String("ALL_NEW"),
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			updateKey: {
				S: aws.String(columnValue),
			},
		},
		UpdateExpression:    aws.String("SET " + columnName + " = " + updateKey),
		ConditionExpression: aws.String("attribute_exists(" + dynamo.RangeKeyName + ")"),
	}
}

func (dynamo *Dynamo) UpdateParams(key string, rangeKey interface{}, columnName string, columnValue *dynamodb.AttributeValue) *dynamodb.UpdateItemInput {
	updateKey := ":updateKey"
	return &dynamodb.UpdateItemInput{
		Key:          dynamo.GetKeyAttribute(key, rangeKey),
		TableName:    aws.String(dynamo.TableName),
		ReturnValues: aws.String("ALL_NEW"),
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			updateKey: columnValue,
		},
		UpdateExpression:    aws.String("SET " + columnName + " = " + updateKey),
		ConditionExpression: aws.String("attribute_exists(" + dynamo.RangeKeyName + ")"),
	}
}
