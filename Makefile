.PHONY: all docs
all: docs speakeasy

speakeasy:
	speakeasy generate sdk --lang terraform -o . -s https://docs.api.epilot.io/product.yaml

docs:
	go generate ./...

