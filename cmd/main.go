package main

import (
	"github.com/sergoslav/go-question-1/pkg/controller"
	"github.com/sergoslav/go-question-1/pkg/repository"
)

func main() {
	r := repository.Init()
	c := controller.Init(r)
	c.CallRepoFunc()
}
