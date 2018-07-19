package dynamodb

type Methods interface {
  GetUser    (*GetUserInput)    (*GetUserOutput,    error)
  UpdateUser (*UpdateUserInput) (*UpdateUserOutput, error)
}
