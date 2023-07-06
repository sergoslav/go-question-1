package controller

import "log"

type Controller struct {
	repo repository
}

type repository interface {
	GetProductRepo() productRepo
}

type productRepo interface {
	GetProduct() string
}

func Init(repo repository) *Controller {
	return &Controller{repo: repo}
}

func (c *Controller) CallRepoFunc() {
	h := c.repo.GetProductRepo().GetProduct()
	log.Printf(h)
}
