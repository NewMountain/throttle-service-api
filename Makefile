run:
	go run server.go

test:
	go test -v --cover

report:
	go test -coverprofile ./test-reports/cover.out && go tool cover -html=./test-reports/cover.out -o ./test-reports/cover.html && open ./test-reports/cover.html

dockerize:
	go build -o ./binaries/app && docker build -t throttle-service:latest .

dockerize-mac:
	env GOOS=linux GOARCH=amd64 go build -o ./binaries/app && docker build -t throttle-service:latest .

docker-run:
	docker run -p 1323:1323 -e ENV=PROD -it throttle-service:latest