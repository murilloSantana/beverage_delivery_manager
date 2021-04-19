package server

//go:generate mockery --name Runner --case=underscore

type Runner interface {
	Run() error
}
