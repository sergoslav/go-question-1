# go-question-1

What the proper way to fix error:

```
cmd/main.go:10:23: cannot use r (variable of type *repository.Repository) as controller.repository value in argument to controller.Init: *repository.Repository does not implement controller.repository (wrong type for method GetProductRepo)
		have GetProductRepo() repository.ProductRepo
		want GetProductRepo() controller.productRepo
```
