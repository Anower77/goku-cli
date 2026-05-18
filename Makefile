APP_NAME=goku

build:
	go build -o $(APP_NAME) .

run:
	go run . -i test.json -o test.yaml

clean:
	del /f /q $(APP_NAME).exe 2>nul || true
	del /f /q test.yaml 2>nul || true

deps:
	go mod tidy

install:
	go get gopkg.in/yaml.v3

help:
	@echo Available commands:
	@echo make build    - Build the project
	@echo make run      - Run the converter
	@echo make clean    - Remove generated files
	@echo make deps     - Clean dependencies
	@echo make install  - Install YAML package