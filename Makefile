.PHONY: default lint-errors lint

default: lint

lint-errors:
	vacuum lint -d todo-openapi.yaml --no-clip --ignore-file vacuum-ignore-file.yaml

lint:
	vacuum lint -d todo-openapi.yaml --no-banner --no-clip --hard-mode --errors --fail-severity 'error' --ignore-file vacuum-ignore-file.yaml
