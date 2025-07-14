.PHONY: default lint-errors lint

default: build

lint-errors:
	vacuum lint -d todo-openapi.yaml --no-banner --no-clip --ignore-file vacuum-ignore-file.yaml
	golangci-lint run

lint:
	vacuum lint -d todo-openapi.yaml --no-banner --no-clip --hard-mode --errors --fail-severity 'error' --ignore-file vacuum-ignore-file.yaml
	golangci-lint run

build: lint-errors
	go test ./...

install:
	# brew install daveshanley/vacuum/vacuum
	# brew install golangci-lint
	# curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/HEAD/install.sh | sh -s -- -b $(go env GOPATH)/bin v2.2.2
	echo "Nope"
