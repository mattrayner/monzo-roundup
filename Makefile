build:
	GOOS=linux go build main.go
	zip main.zip main

deploy:
	aws lambda update-function-code \
		--function-name monzo-roundup \
		--zip-file fileb://main.zip