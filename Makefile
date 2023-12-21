BINARY_NAME=scan

build:
	GOARCH=amd64 GOOS=darwin go build -o ./bin/${BINARY_NAME}-darwin ./cmd/${BINARY_NAME}/main.go
	GOARCH=amd64 GOOS=linux go build -o ./bin/${BINARY_NAME}-linux ./cmd/${BINARY_NAME}/main.go

docker-run-mysql:
	docker run --rm --name chall-mysql-temp --net=host -e MYSQL_ROOT_PASSWORD=challenge -d mysql:8.2.0

docker-build:
	docker build -f ./deploy/Dockerfile -t chall-scan-temp .

docker-run:
	docker run --rm --name=chall-scan-temp --net=host chall-scan-temp

test: docker-run
