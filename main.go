package main

import (
  "errors"
  "os"
  "github.com/aws/aws-sdk-go/aws/session"
  "github.com/aws/aws-sdk-go/aws"
  "github.com/aws/aws-sdk-go/service/dynamodb"
  "fmt"
  "github.com/aws/aws-lambda-go/lambda"
  "net/http"
  "net/url"
  "strings"
  "encoding/json"
  "github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
  "strconv"
  "io/ioutil"
)

var (
  CLIENT_ID     = os.Getenv("MONZO_CLIENT_ID")
  CLIENT_SECRET = os.Getenv("MONZO_CLIENT_SECRET")
  AWS_REGION    = os.Getenv("AWS_REGION")
  PRODUCTION    = os.Getenv("GO_ENV") == "production"
)

type Request struct {
  AccountID     string `json:"account_id"`
  TransactionID string `json:"transaction_id"`
}

type UserData struct {
  ID           string `json:"id"`
  AuthToken    string `json:"auth_token"`
  RefreshToken string `json:"refresh_token"`
}

type UserDataUpdate struct {
  AuthToken    string `json:":a"`
  RefreshToken string `json:":r"`
}

type RefreshAPIResponse struct {
  AccessToken  string `json:"access_token"`
  ClientID     string `json:"client_id"`
  ExpiresIn    int    `json:"expires_in"`
  RefreshToken string `json:"refresh_token"`
  TokenType    string `json:"token_type"`
  UserID       string `json:"user_id"`
}

type TransactionAPIResponse struct {
  Transaction Transaction `json:"transaction"`
}

type Transaction struct {
  ID          string `json:"id"`
  Created     string `json:"created"`
  Description string `json:"description"`
  Amount      int    `json:"amount"`
  LocalAmount int    `json:"local_amount"`
  Originator  bool   `json:"originator"`
}

type PotsAPIResponse struct {
  Pots []Pot `json:"pots"`
}

type Pot struct {
  ID       string `json:"id"`
  Name     string `json:"name"`
  Style    string `json:"style"`
  Balance  int    `json:"balance"`
  Currency string `json:"currency"`
  Created  string `json:"created"`
  Updated  string `json:"updated"`
  Deleted  bool   `json:"deleted"`
}

// AWS_REGION=eu-west-1 aws dynamodb create-table --table-name monzo-roundup --attribute-definitions AttributeName=id,AttributeType=S --token-schema AttributeName=id,TokenType=HASH --provisioned-throughput ReadCapacityUnits=5,WriteCapacityUnits=5 --endpoint-url http://localhost:8000

func DynamoDBService() (*dynamodb.DynamoDB, error) {
  config := &aws.Config{Region: aws.String(AWS_REGION)}

  if !PRODUCTION {
    // config.Endpoint = aws.String("http://localhost:8000")
  }

  sess, err := session.NewSession(config)
  if err != nil {
    return nil, errors.New("error creating new session object")
  }

  return dynamodb.New(sess), nil
}

func GetUser(userId string, svc *dynamodb.DynamoDB) (UserData, error) {
  userData := UserData{}

  result, err := svc.GetItem(&dynamodb.GetItemInput{
    TableName: aws.String("monzo-roundup"),
    Key: map[string]*dynamodb.AttributeValue{
      "id": {
        S: aws.String(userId),
      },
    },
  })

  if err != nil {
    fmt.Println(err.Error())
    return userData, err
  }

  err = dynamodbattribute.UnmarshalMap(result.Item, &userData)

  if err != nil {
    fmt.Printf("Failed to unmarshal Record, %v", err)
    return UserData{}, err
  }

  return userData, nil
}

func RefreshToken(refreshToken string, accountID string) (UserData, error) {
  uri := "https://api.monzo.com/oauth2/token"

  data := url.Values{}
  data.Set("grant_type", "refresh_token")
  data.Add("client_id", CLIENT_ID)
  data.Add("client_secret", CLIENT_SECRET)
  data.Add("refresh_token", refreshToken)

  // Construct a new Request
  fmt.Println("Creating request object")
  req, err := http.NewRequest("POST", uri, strings.NewReader(data.Encode()))
  req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
  // req.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))

  fmt.Println("Calling MakeRequest")
  refreshData, err := MakeRequest(req, err, uri, false, "")
  fmt.Println("Called MakeRequest")

  if err != nil {
    errorMessage := fmt.Sprintf("Error making request: %v", err)

    fmt.Print(errorMessage)

    return UserData{}, errors.New(errorMessage)
  }

  fmt.Println("Decoding JSON")
  var refreshResponse RefreshAPIResponse
  if err := json.Unmarshal(refreshData, &refreshResponse); err != nil {
    errorMessage := fmt.Sprintf("Error parsing JSON: %v", err)

    fmt.Print(errorMessage)

    return UserData{}, errors.New(errorMessage)
  }

  userData := UserData{
    ID:           accountID,
    AuthToken:    refreshResponse.AccessToken,
    RefreshToken: refreshResponse.RefreshToken,
  }

  return userData, nil
}

func MakeRequest(httpRequest *http.Request, httpError error, uri string, includeAuth bool, authToken string) ([]byte, error) {
  client := &http.Client{}

  if httpError != nil {
    errorMessage := fmt.Sprintf("error creating request object for: %s\n", uri)

    fmt.Print(errorMessage)

    return nil, errors.New(errorMessage)
  }

  if includeAuth {
    fmt.Println("Adding auth header")
    authHeaderValue := fmt.Sprintf("Bearer %s", authToken)
    httpRequest.Header.Add("Authorization", authHeaderValue)
  }

  fmt.Println("Making request")
  resp, err := client.Do(httpRequest)
  if err != nil {
    errorMessage := fmt.Sprintf("error making request to: %s\n", uri)

    fmt.Print(errorMessage)

    defer resp.Body.Close()

    return nil, errors.New(errorMessage)
  }

  fmt.Printf("Recieved status code: %v\n", resp.StatusCode)

  fmt.Println("Reading body")
  body, err := ioutil.ReadAll(resp.Body)

  defer resp.Body.Close()

  if resp.StatusCode != 200 {
    errorMessage := fmt.Sprintf("Non-200 Status code: %s\n", body)

    fmt.Print(errorMessage)

    return nil, errors.New(errorMessage)
  }

  return body, err
}

func GetRequest(uri string, includeAuth bool, authToken string) ([]byte, error) {
  // Construct a new Request
  req, err := http.NewRequest("GET", uri, nil)

  return MakeRequest(req, err, uri, includeAuth, authToken)
}

func UpdateUser(userData UserData, svc *dynamodb.DynamoDB) (error) {
  // Create an update object
  update, updateError := dynamodbattribute.MarshalMap(UserDataUpdate{
    AuthToken:    userData.AuthToken,
    RefreshToken: userData.RefreshToken,
  })
  if updateError != nil {
    fmt.Println(updateError.Error())
    return updateError
  }

  // Update the object
  fmt.Println("Updating record")
  input := &dynamodb.UpdateItemInput{
    TableName: aws.String("monzo-roundup"),
    Key: map[string]*dynamodb.AttributeValue{
      "id": {
        S: aws.String(userData.ID),
      },
    },
    UpdateExpression:          aws.String("set auth_token=:a, refresh_token=:r"),
    ExpressionAttributeValues: update,
    ReturnValues:              aws.String("UPDATED_NEW"),
  }

  _, err := svc.UpdateItem(input)

  if err != nil {
    fmt.Println(err.Error())
    return err
  }

  fmt.Println("Updated record")

  return nil
}

func GetTransaction(transactionId string, requestToken string) (TransactionAPIResponse, error) {
  fmt.Printf("Requesting transaction: %v\n", transactionId)

  uri := fmt.Sprintf("https://api.monzo.com/transactions/%s", transactionId)

  transactionData, _ := GetRequest(uri, true, requestToken)

  var data TransactionAPIResponse
  if err := json.Unmarshal(transactionData, &data); err != nil {
    errorMessage := fmt.Sprintf("error parsing JSON: %s", err)

    fmt.Print(errorMessage)

    return TransactionAPIResponse{}, errors.New(errorMessage)
  }

  return data, nil
}

func Transfer(diff int, userData UserData, transactionID string) (string, error) {
  fmt.Println("Getting Coin Jar Pot")
  pot, err := GetCoinJarPot(userData)
  if err != nil {
    errorMessage := fmt.Sprintf("error getting Coin Jar: %v", err)

    fmt.Println(errorMessage)

    return "", errors.New(errorMessage)
  }
  fmt.Println("Got Coin Jar Pot")

  fmt.Printf("Depositing into pot (%v)\n", pot.ID)
  balance, err := DepositIntoPot(diff, pot, userData, transactionID)
  if err != nil {
    errorMessage := fmt.Sprintf("error depositing into Pot: %v", err)

    fmt.Println(errorMessage)

    return "", errors.New(errorMessage)
  }

  var message string
  if pot.Balance == balance {
    message = fmt.Sprintf("Transfer did not happen - pot (%v) balance unmodified: %v", pot.ID, balance)
  } else {
    message = fmt.Sprintf("Updated pot (%v) old balance: %v, new balance is: %v", pot.ID, pot.Balance, balance)
  }

  fmt.Println(message)

  return message, nil
}

func GetCoinJarPot(userData UserData) (Pot, error) {
  potsData, _ := GetRequest("https://api.monzo.com/pots", true, userData.AuthToken)

  var data PotsAPIResponse
  if err := json.Unmarshal(potsData, &data); err != nil {
    errorMessage := fmt.Sprintf("error parsing JSON: %s", err)

    fmt.Print(errorMessage)

    return Pot{}, errors.New(errorMessage)
  }

  var coinJar Pot
  for _, pot := range data.Pots {
    if pot.Name == "Coin Jar" && !pot.Deleted {
      coinJar = pot
      break
    }
  }

  if len(coinJar.ID) == 0 {
    errorMessage := "unable to find an active pot named 'Coin Jar'"

    fmt.Println(errorMessage)

    return coinJar, errors.New(errorMessage)
  }

  return coinJar, nil
}

func DepositIntoPot(amount int, pot Pot, userData UserData, transactionID string) (int, error) {
  uri := fmt.Sprintf("https://api.monzo.com/pots/%v/deposit", pot.ID)

  dedupeID := fmt.Sprintf("%v_%v", userData.ID, transactionID)

  data := url.Values{}
  data.Set("source_account_id", userData.ID)
  data.Add("amount", strconv.Itoa(amount))
  data.Add("dedupe_id", dedupeID)

  fmt.Printf("Attempting deposit with dedupe_id: %v\n", dedupeID)

  // Construct a new Request
  req, err := http.NewRequest("PUT", uri, strings.NewReader(data.Encode()))
  req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
  req.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))

  potData, err := MakeRequest(req, err, uri, true, userData.AuthToken)

  var potResponse Pot
  if err := json.Unmarshal(potData, &potResponse); err != nil {
    errorMessage := fmt.Sprintf("Error parsing JSON: %v", err)

    fmt.Print(errorMessage)

    return -1, errors.New(errorMessage)
  }

  return potResponse.Balance, nil
}

func Handler(request Request) (string, error) {
  fmt.Printf("STARTING WITH:\n%v\n%v\n\n", request.AccountID, request.TransactionID)

  // Fetch the user's details from DynamoDB
  fmt.Println("Creating DynamoDB Service")
  svc, err := DynamoDBService()
  if err != nil {
    return "", err
  }
  fmt.Println("DynamoDB Service Created")

  // Get user from DynamoDB
  fmt.Println("Getting user from DynamoDB")
  userData, err := GetUser(request.AccountID, svc)
  if err != nil {
    return "Error getting user from DynamoDB", err
  }
  fmt.Println("Got user from DynamoDB")

  // Refresh the token from Monzo
  fmt.Println("Refreshing users Access Token")
  refreshedUserData, err := RefreshToken(userData.RefreshToken, request.AccountID)
  if err != nil {
    return "Error refreshing user token", err
  }
  fmt.Println("Refreshed users Access Token")

  // Save updated token to DynamoDB
  fmt.Println("Updating users entry in DynamoDB")
  dynamodbError := UpdateUser(refreshedUserData, svc)
  if dynamodbError != nil {
    return "Error updating user data in DynamoDB", dynamodbError
  }
  fmt.Println("Updated users entry in DynamoDB")

  // Fetch transaction from Monzo
  fmt.Println("Getting transaction from Monzo")
  transactionData, err := GetTransaction(request.TransactionID, refreshedUserData.AuthToken)
  if err != nil {
    return "Error getting transaction", err
  }
  fmt.Println("Got transaction from Monzo")

  // Work out if there is change
  fmt.Printf("Calclating Diff (100 - %v)\n", transactionData.Transaction.LocalAmount)
  diff := 100 + transactionData.Transaction.LocalAmount
  fmt.Printf("Calculated Diff: %v\n", diff)

  fmt.Println("Do we need to round?")
  if diff <= 0 || diff >= 100 {
    fmt.Println("No")
    return "No need to round", nil
  }
  fmt.Println("Yes")

  // Move change into pot
  fmt.Println("Initialising transfer")
  newBalanceMessage, err := Transfer(diff, refreshedUserData, request.TransactionID)
  if err != nil {
    fmt.Println("Error transferring")
    return "Error transferring", err
  }
  fmt.Println("Transferred")

  return newBalanceMessage, nil
}

func main() {
  lambda.Start(Handler)
}
