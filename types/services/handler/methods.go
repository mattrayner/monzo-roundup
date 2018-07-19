package handler

type Methods interface {
  Handle(input *HandleInput) (HandleOutput, error)
}
