package dynamodb

import (
  "github.com/mattrayner/monzo-roundup/types/services/dynamodb"
  "github.com/aws/aws-sdk-go/aws"
  "fmt"
  "errors"
  _dynamodb "github.com/aws/aws-sdk-go/service/dynamodb"
  "github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
  "log"
)

func (s *service) GetUser(input *dynamodb.GetUserInput)    (*dynamodb.GetUserOutput,    error) {
  getItemInput := &_dynamodb.GetItemInput{
    TableName: aws.String("monzo-roundup"),
    Key: map[string]*_dynamodb.AttributeValue{
      "id": {
        S: aws.String(input.AccountID),
      },
    },
  }
  result, err := s.srv.GetItem(getItemInput); if err != nil {
    errorMessage := fmt.Sprintf("Error getting item from DyanamoDB (id: %v): %v", input.AccountID, err)

    return &dynamodb.GetUserOutput{}, errors.New(errorMessage)
  }

  // Were we unable to get an item for the userID?
  if len(result.Item) == 0 {
    errorMessage := fmt.Sprintf("Unable to find a DynamoDB item for (id: %v)", input.AccountID)

    return &dynamodb.GetUserOutput{}, errors.New(errorMessage)
  }

  // Try an unmarshal our item into a UserData object
  var output = dynamodb.User{}
  err = dynamodbattribute.UnmarshalMap(result.Item, &output); if err != nil {
    errorMessage := fmt.Sprintf("Failed to unmarshal DynamoDB item (id: %v), %v", input.AccountID, err)

    return &dynamodb.GetUserOutput{}, errors.New(errorMessage)
  }

  return &dynamodb.GetUserOutput{ User: &output }, nil
}

func (s *service) UpdateUser(input *dynamodb.UpdateUserInput) (*dynamodb.UpdateUserOutput, error) {
  type UserDataUpdate struct {
    AuthToken    string `json:":a"`
    RefreshToken string `json:":r"`
  }

  // Create an update object
  update, updateError := dynamodbattribute.MarshalMap(UserDataUpdate{
    AuthToken:    input.User.AuthKey,
    RefreshToken: input.User.RefreshKey,
  })

  if updateError != nil {
    log.Println(updateError.Error())
    return &dynamodb.UpdateUserOutput{}, updateError
  }

  // Update the object
  log.Println("Updating record")
  updateItemInput := &_dynamodb.UpdateItemInput{
    TableName: aws.String("monzo-roundup"),
    Key: map[string]*_dynamodb.AttributeValue{
      "id": {
        S: aws.String(input.User.AccountID),
      },
    },
    UpdateExpression:          aws.String("set auth_token=:a, refresh_token=:r"),
    ExpressionAttributeValues: update,
    ReturnValues:              aws.String("UPDATED_NEW"),
  }

  result, err := s.srv.UpdateItem(updateItemInput); if err != nil {
    errorMessage := fmt.Sprintf("Error updating item in DynamoDB: %v", err)

    return &dynamodb.UpdateUserOutput{}, errors.New(errorMessage)
  }

  if len(result.Attributes) == 0 {
    errorMessage := fmt.Sprintf("Update to DynamoDB did nothing (id: %v)", input.User.AccountID)

    return &dynamodb.UpdateUserOutput{}, errors.New(errorMessage)
  }

  // Try an unmarshal our item into a UserData object
  var output = dynamodb.User{}
  err = dynamodbattribute.UnmarshalMap(result.Attributes, &output); if err != nil {
    errorMessage := fmt.Sprintf("Failed to unmarshal DynamoDB item (id: %v), %v", input.User.AccountID, err)

    return &dynamodb.UpdateUserOutput{}, errors.New(errorMessage)
  }

  return &dynamodb.UpdateUserOutput{ User: &output }, nil
}