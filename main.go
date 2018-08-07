package monzo_roundup

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
  "log"
  "time"
)

var (
  CLIENT_ID     = os.Getenv("MONZO_CLIENT_ID")
  CLIENT_SECRET = os.Getenv("MONZO_CLIENT_SECRET")
  AWS_REGION    = os.Getenv("AWS_REGION")
  PRODUCTION    = os.Getenv("GO_ENV") == "production"
)

type Request struct {
  AccountID     string `json:"account_id"`
  Amount        string `json:"amount"`
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

  getItemInput := &dynamodb.GetItemInput{
    TableName: aws.String("monzo-roundup"),
    Key: map[string]*dynamodb.AttributeValue{
      "id": {
        S: aws.String(userId),
      },
    },
  }
  result, err := svc.GetItem(getItemInput); if err != nil {
    errorMessage := fmt.Sprintf("Error getting item from DyanamoDB (id: %v): %v", userId, err)

    return userData, errors.New(errorMessage)
  }

  // Were we unable to get an item for the userID?
  if len(result.Item) == 0 {
    errorMessage := fmt.Sprintf("Unable to find a DynamoDB item for (id: %v)", userId)

    return userData, errors.New(errorMessage)
  }

  // Try an unmarshal our item into a UserData object
  err = dynamodbattribute.UnmarshalMap(result.Item, &userData); if err != nil {
    errorMessage := fmt.Sprintf("Failed to unmarshal DynamoDB item (id: %v), %v", userId, err)

    return UserData{}, errors.New(errorMessage)
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
  log.Println("Creating request object")
  req, err := http.NewRequest("POST", uri, strings.NewReader(data.Encode()))
  req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
  // req.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))

  log.Println("Calling MakeRequest")
  refreshData, err := MakeRequest(req, err, uri, false, "")
  log.Println("Called MakeRequest")

  if err != nil {
    errorMessage := fmt.Sprintf("Error refreshing Monzo token: %v", err)

    log.Println(errorMessage)

    return UserData{}, errors.New(errorMessage)
  }

  log.Println("Decoding JSON")
  var refreshResponse RefreshAPIResponse
  if err := json.Unmarshal(refreshData, &refreshResponse); err != nil {
    errorMessage := fmt.Sprintf("Error parsing JSON: %v", err)

    log.Println(errorMessage)

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

    log.Print(errorMessage)

    return nil, errors.New(errorMessage)
  }

  if includeAuth {
    log.Println("Adding auth header")
    authHeaderValue := fmt.Sprintf("Bearer %s", authToken)
    httpRequest.Header.Add("Authorization", authHeaderValue)
  }

  log.Println("Making request")
  resp, err := client.Do(httpRequest)
  if err != nil {
    errorMessage := fmt.Sprintf("error making request to: %s\n", uri)

    fmt.Print(errorMessage)

    defer resp.Body.Close()

    return nil, errors.New(errorMessage)
  }

  log.Printf("Recieved status code: %v\n", resp.StatusCode)

  log.Println("Reading body")
  body, err := ioutil.ReadAll(resp.Body)

  defer resp.Body.Close()

  if resp.StatusCode != 200 {
    errorMessage := fmt.Sprintf("Non-200 Status code. Status code: (%v), body: %s", resp.StatusCode, body)

    log.Println(errorMessage)

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
    log.Println(updateError.Error())
    return updateError
  }

  // Update the object
  log.Println("Updating record")
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

  result, err := svc.UpdateItem(input); if err != nil {
    errorMessage := fmt.Sprintf("Error updating item in DynamoDB: %v", err)

    return errors.New(errorMessage)
  }

  if len(result.Attributes) == 0 {
    errorMessage := fmt.Sprintf("Update to DynamoDB did nothing (id: %v)", userData.ID)

    return errors.New(errorMessage)
  }

  return nil
}

func GetTransaction(transactionId string, requestToken string) (TransactionAPIResponse, error) {
  log.Printf("Requesting transaction: %v\n", transactionId)

  uri := fmt.Sprintf("https://api.monzo.com/transactions/%s", transactionId)

  transactionData, _ := GetRequest(uri, true, requestToken)

  var data TransactionAPIResponse
  if err := json.Unmarshal(transactionData, &data); err != nil {
    errorMessage := fmt.Sprintf("error parsing JSON: %s", err)

    log.Print(errorMessage)

    return TransactionAPIResponse{}, errors.New(errorMessage)
  }

  return data, nil
}

func Transfer(diff int, userData UserData) (string, error) {
  log.Println("Getting Coin Jar Pot")
  pot, err := GetCoinJarPot(userData)
  if err != nil {
    errorMessage := fmt.Sprintf("error getting Coin Jar: %v", err)

    log.Println(errorMessage)

    return "", errors.New(errorMessage)
  }
  log.Println("Got Coin Jar Pot")

  log.Printf("Depositing into pot (%v)\n", pot.ID)
  balance, err := DepositIntoPot(diff, pot, userData)
  if err != nil {
    errorMessage := fmt.Sprintf("error depositing into Pot: %v", err)

    log.Println(errorMessage)

    return "", errors.New(errorMessage)
  }

  var message string
  if pot.Balance == balance {
    message = fmt.Sprintf("Transfer did not happen - pot (%v) balance unmodified: %v", pot.ID, balance)
  } else {
    message = fmt.Sprintf("Updated pot (%v) old balance: %v, new balance is: %v", pot.ID, pot.Balance, balance)
  }

  log.Println(message)

  return message, nil
}

func GetCoinJarPot(userData UserData) (Pot, error) {
  potsData, err := GetRequest("https://api.monzo.com/pots", true, userData.AuthToken); if err != nil {
    errorMessage := fmt.Sprintf("error getting pots from Monzo: %v", err)

    return Pot{}, errors.New(errorMessage)
  }

  var data PotsAPIResponse
  err = json.Unmarshal(potsData, &data); if err != nil {
    errorMessage := fmt.Sprintf("error parsing JSON: %s", err)

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

    log.Println(errorMessage)

    return coinJar, errors.New(errorMessage)
  }

  return coinJar, nil
}

func DepositIntoPot(amount int, pot Pot, userData UserData) (int, error) {
  uri := fmt.Sprintf("https://api.monzo.com/pots/%v/deposit", pot.ID)

  dedupeID := fmt.Sprintf("%v_%v", userData.ID, time.Now().Unix())

  data := url.Values{}
  data.Set("source_account_id", userData.ID)
  data.Add("amount", strconv.Itoa(amount))
  data.Add("dedupe_id", dedupeID)

  log.Printf("Attempting deposit with dedupe_id: %v\n", dedupeID)

  // Construct a new Request
  req, err := http.NewRequest("PUT", uri, strings.NewReader(data.Encode()))
  req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
  req.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))

  potData, err := MakeRequest(req, err, uri, true, userData.AuthToken)
  if err != nil {
    log.Println(err.Error())

    return -1, err
  }

  var potResponse Pot
  if err := json.Unmarshal(potData, &potResponse); err != nil {
    errorMessage := fmt.Sprintf("Error parsing JSON: %v", err)

    log.Print(errorMessage)

    return -1, errors.New(errorMessage)
  }

  return potResponse.Balance, nil
}

func Handler(request Request) (string, error) {
  log.Printf("STARTING WITH:\n%v\n%v\n\n", request.AccountID, request.Amount)

  // Fetch the user's details from DynamoDB
  log.Println("Creating DynamoDB Service")
  svc, err := DynamoDBService()
  if err != nil {
    return "", err
  }
  log.Println("DynamoDB Service Created")

  // Get user from DynamoDB
  log.Println("Getting user from DynamoDB")
  userData, err := GetUser(request.AccountID, svc)
  if err != nil {
    return "Error getting user from DynamoDB", err
  }
  log.Println("Got user from DynamoDB")

  // Refresh the token from Monzo
  log.Println("Refreshing users Access Token")
  refreshedUserData, err := RefreshToken(userData.RefreshToken, request.AccountID)
  if err != nil {
    return "Error refreshing user token", err
  }
  log.Println("Refreshed users Access Token")

  // Save updated token to DynamoDB
  log.Println("Updating users entry in DynamoDB")
  dynamodbError := UpdateUser(refreshedUserData, svc)
  if dynamodbError != nil {
    return "Error updating user data in DynamoDB", dynamodbError
  }
  log.Println("Updated users entry in DynamoDB")

  // Work out if there is change
  log.Printf("Calclating Diff (100 - %v)\n", request.Amount)
  intAmount, err := strconv.ParseInt(request.Amount, 10, 32)
  if err != nil {
    return "Unable to convert string to int", err
  }
  diff := 100 + int(intAmount)
  log.Printf("Calculated Diff: %v\n", diff)

  log.Println("Do we need to round?")
  if diff <= 0 || diff >= 100 {
    log.Println("No")
    return "No need to round", nil
  }
  log.Println("Yes")

  // Move change into pot
  log.Println("Initialising transfer")
  newBalanceMessage, err := Transfer(diff, refreshedUserData)
  if err != nil {
    log.Println("Error transferring")
    return "Error transferring", err
  }
  log.Println("Transferred")

  return newBalanceMessage, nil
}

func main() {
  lambda.Start(Handler)
}
