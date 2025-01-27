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

func (dynamo *Dynamo) PutOrReplaceItemParams(values map[string]*dynamodb.AttributeValue) *dynamodb.PutItemInput {
	return &dynamodb.PutItemInput{
		Item:      values,
		TableName: aws.String(dynamo.TableName),
	}
}

func (dynamo *Dynamo) UpdateItemParams(values map[string]*dynamodb.AttributeValue) *dynamodb.PutItemInput {
	return &dynamodb.PutItemInput{
		Item:                values,
		TableName:           aws.String(dynamo.TableName), // Required
		ConditionExpression: aws.String("attribute_exists(" + dynamo.RangeKeyName + ")"),
	}
}

func (dynamo *Dynamo) UpdateItemParamsWithOutCondition(values map[string]*dynamodb.AttributeValue) *dynamodb.PutItemInput {
	return &dynamodb.PutItemInput{
		Item:      values,
		TableName: aws.String(dynamo.TableName), // Required\
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
	//FTODO: inproper interface conversion will throw nil exceptions avoid them
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

func (dynamo *Dynamo) QueryParamsWithIndexName(indexName string, keyName string, keyValue string) *dynamodb.QueryInput {

	params := &dynamodb.QueryInput{
		TableName:        aws.String(dynamo.TableName), // Required
		IndexName:        aws.String(indexName),
		Limit:            aws.Int64(dynamo.Limit),
		ScanIndexForward: aws.Bool(false),
		KeyConditions: map[string]*dynamodb.Condition{
			keyName: { // Required
				ComparisonOperator: aws.String(dynamodb.ComparisonOperatorEq), // Required
				AttributeValueList: []*dynamodb.AttributeValue{
					{
						S: aws.String(keyValue),
					},
				},
			},
		},
	}
	return params
}

func (dynamo *Dynamo) QueryParamsWithIndexNameNumber(indexName string, keyName string, keyValue int) *dynamodb.QueryInput {

	params := &dynamodb.QueryInput{
		TableName:        aws.String(dynamo.TableName), // Required
		IndexName:        aws.String(indexName),
		Limit:            aws.Int64(dynamo.Limit),
		ScanIndexForward: aws.Bool(false),
		KeyConditions: map[string]*dynamodb.Condition{
			keyName: { // Required
				ComparisonOperator: aws.String(dynamodb.ComparisonOperatorEq), // Required
				AttributeValueList: []*dynamodb.AttributeValue{
					{
						N: aws.String(strconv.Itoa(keyValue)),
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
		ConditionExpression: aws.String("attribute_exists(" + dynamo.KeyName + ")"),
	}
}

func (dynamo *Dynamo) UpdateNumberParams(key string, rangeKey interface{}, columnName string, columnValue int) *dynamodb.UpdateItemInput {
	updateKey := ":updateKey"
	return &dynamodb.UpdateItemInput{
		Key:          dynamo.GetKeyAttribute(key, rangeKey),
		TableName:    aws.String(dynamo.TableName),
		ReturnValues: aws.String("ALL_NEW"),
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			updateKey: {
				N: aws.String(strconv.Itoa(columnValue)),
			},
		},
		UpdateExpression:    aws.String("SET " + columnName + " = " + updateKey),
		ConditionExpression: aws.String("attribute_exists(" + dynamo.KeyName + ")"),
	}
}

func (dynamo *Dynamo) UpdateBoolParams(key string, rangeKey interface{}, columnName string, columnValue bool) *dynamodb.UpdateItemInput {
	updateKey := ":updateKey"
	return &dynamodb.UpdateItemInput{
		Key:          dynamo.GetKeyAttribute(key, rangeKey),
		TableName:    aws.String(dynamo.TableName),
		ReturnValues: aws.String("ALL_NEW"),
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			updateKey: {
				BOOL: aws.Bool(columnValue),
			},
		},
		UpdateExpression:    aws.String("SET " + columnName + " = " + updateKey),
		ConditionExpression: aws.String("attribute_exists(" + dynamo.KeyName + ")"),
	}
}

func (dynamo *Dynamo) IncrementColumnValue(key string, rangeKey interface{}, columnName string) *dynamodb.UpdateItemInput {
	updateKey := ":updateKey"
	return &dynamodb.UpdateItemInput{
		Key:          dynamo.GetKeyAttribute(key, rangeKey),
		TableName:    aws.String(dynamo.TableName),
		ReturnValues: aws.String("ALL_NEW"),
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			updateKey: {
				N: aws.String(strconv.Itoa(1)),
			},
		},
		UpdateExpression:    aws.String("ADD " + columnName + " " + updateKey),
		ConditionExpression: aws.String("attribute_exists(" + dynamo.RangeKeyName + ")"),
	}
}

func (dynamo *Dynamo) IncrementColumnValueWithoutRangeKey(key string, columnName string) *dynamodb.UpdateItemInput {
	updateKey := ":updateKey"
	return &dynamodb.UpdateItemInput{
		Key:          dynamo.GetKeyAttribute(key, nil),
		TableName:    aws.String(dynamo.TableName),
		ReturnValues: aws.String("ALL_NEW"),
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			updateKey: {
				N: aws.String(strconv.Itoa(1)),
			},
		},
		UpdateExpression: aws.String("ADD " + columnName + " " + updateKey),
	}
}

func (dynamo *Dynamo) DecrementColumnValue(key string, rangeKey interface{}, columnName string) *dynamodb.UpdateItemInput {
	updateKey := ":updateKey"
	return &dynamodb.UpdateItemInput{
		Key:          dynamo.GetKeyAttribute(key, rangeKey),
		TableName:    aws.String(dynamo.TableName),
		ReturnValues: aws.String("ALL_NEW"),
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			updateKey: {
				N: aws.String(strconv.Itoa(-1)),
			},
		},
		UpdateExpression:    aws.String("ADD " + columnName + " " + updateKey),
		ConditionExpression: aws.String("attribute_exists(" + dynamo.RangeKeyName + ")"),
	}
}

func (dynamo *Dynamo) UpdateInterfaceParams(key string, rangeKey interface{}, columnName string, columnValue interface{}) *dynamodb.UpdateItemInput {
	switch columnValue.(type) {
	case int:
		return dynamo.UpdateNumberParams(key, rangeKey, columnName, columnValue.(int))
	case string:
		return dynamo.UpdateStringParams(key, rangeKey, columnName, columnValue.(string))
	case bool:
		return dynamo.UpdateBoolParams(key, rangeKey, columnName, columnValue.(bool))
	default:
		return nil
	}
}

//this is update params object
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

func (dynamo *Dynamo) UpdateParamsWithoutRangeKey(key string, columnName string, columnValue *dynamodb.AttributeValue) *dynamodb.UpdateItemInput {
	updateKey := ":updateKey"
	return &dynamodb.UpdateItemInput{
		Key: map[string]*dynamodb.AttributeValue{
			dynamo.KeyName: {
				S: aws.String(key),
			},
		},
		TableName:    aws.String(dynamo.TableName),
		ReturnValues: aws.String("ALL_NEW"),
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			updateKey: columnValue,
		},
		UpdateExpression: aws.String("SET " + columnName + " = " + updateKey),
	}
}

func (dynamo *Dynamo) ScanParams() *dynamodb.ScanInput {
	return &dynamodb.ScanInput{
		TableName: aws.String(dynamo.TableName),
	}

}
