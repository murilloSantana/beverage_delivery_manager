package server

//go:generate mockery --name Runner --case=underscore --output ../../mocks

type Runner interface {
	Run() error
}
