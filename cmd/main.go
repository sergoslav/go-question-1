package main

import (
	"github.com/sergoslav/go-question-1/internal/controller"
	"github.com/sergoslav/go-question-1/internal/repository"
)

func main() {
	r := repository.Init()
	c := controller.Init(r)
	c.CallRepoFunc()
}
