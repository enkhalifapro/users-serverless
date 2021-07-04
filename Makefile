.PHONY: build clean deploy gomodgen

build: gomodgen
	export GO111MODULE=on
	env GOOS=linux go build -ldflags="-s -w" -o bin/hello hello/main.go
	env GOOS=linux go build -ldflags="-s -w" -o bin/world world/main.go
	env GOOS=linux go build -ldflags="-s -w" -o bin/users internal/users/lambda/main.go
	env GOOS=linux go build -gcflags="all=-N -l" -o bin/users-debug internal/users/lambda/main.go

clean:
	rm -rf ./bin ./vendor ./api go.sum

deploy: build
	sls deploy --verbose

deploy-local: build
	serverless deploy --stage local --verbose

sls-debug:
	SLS_DEBUG=* sls offline --useDocker

gomodgen:
	chmod u+x gomod.sh
	./gomod.sh

generate: api

api: clean
	goa gen github.com/enkhalifapro/users-serverless/internal/users/api/design -o ./api/users