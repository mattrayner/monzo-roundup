.PHONY: build build-binary build-proto clean clean-binary clean-proto deploy test test-unit test-skip-stress

# Build our application
build: clean build-proto build-binary

build-binary:
	GOOS=linux go build main.go
	zip main.zip main

build-proto: clean-proto
	protowrap -I. --go_out `go env GOPATH`/src ./**/**/*.proto

# Clean up artifacts
clean: clean-binary clean-proto

clean-binary:
	rm main
	rm main.zip

clean-proto:
	rm  `find . -name "*.pb.go"`

# Deploy our built code to AWS
deploy:
	aws lambda update-function-code \
		--function-name monzo-roundup \
		--zip-file fileb://main.zip

# Test commands
test:
	ginkgo ./...

test-unit:
	go test -cover `go list ./... | grep -v /stress | grep -v /spec | grep -v /types/`

test-skip-stress:
	ginkgo --skip=Stress ./...

# DynamoDB test setup commands
db: db-start db-create

db-build:
	docker-compose -f docker-compose.test.yml build

db-start:
	docker-compose -f docker-compose.test.yml up -d

db-create:
	AWS_REGION=eu-west-1 aws dynamodb create-table --table-name monzo-roundup --attribute-definitions AttributeName=id,AttributeType=S --key-schema AttributeName=id,KeyType=HASH --provisioned-throughput ReadCapacityUnits=5,WriteCapacityUnits=5 --endpoint-url http://localhost:8000

db-stop:
	docker-compose -f docker-compose.test.yml stop
