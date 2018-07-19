package networking

import (
  "log"
  "fmt"
  "errors"
  "io/ioutil"
  "net/http"
)

type HttpClient interface {
  Do(*http.Request) (*http.Response, error)
}

func MakeRequest(client HttpClient, httpRequest *http.Request, httpError error, uri string, includeAuth bool, authToken string) ([]byte, error) {
  if httpError != nil {
    errorMessage := fmt.Sprintf("error creating request object for: %s", uri)

    log.Println(errorMessage)

    return nil, errors.New(errorMessage)
  }

  if includeAuth {
    if authToken == "" {
      return nil, errors.New("argument error - authToken is required if includeAuth is true")
    }
    log.Println("Adding auth header")
    authHeaderValue := fmt.Sprintf("Bearer %s", authToken)
    httpRequest.Header.Add("Authorization", authHeaderValue)
  }

  log.Println("Making request")
  resp, err := client.Do(httpRequest)
  if err != nil {
    errorMessage := fmt.Sprintf("error making request to: %s", uri)

    log.Println(errorMessage)

    defer resp.Body.Close()

    return nil, errors.New(errorMessage)
  }

  log.Printf("Recieved status code: %v\n", resp.StatusCode)

  log.Println("Reading body")
  body, err := ioutil.ReadAll(resp.Body); if err != nil {
    errorMessage := fmt.Sprintf("Error reading body: %v", err)

    log.Println(errorMessage)

    return nil, errors.New(errorMessage)
  }

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
  client := &http.Client{}

  return MakeRequest(client, req, err, uri, includeAuth, authToken)
}