run:
	ENV=DEV go run server.go handlers.go

test:
	ENV=DEV o test -v --cover

report:
	ENV=DEV go test -coverprofile ./test-reports/cover.out && go tool cover -html=./test-reports/cover.out -o ./test-reports/cover.html && open ./test-reports/cover.html

dockerize:
	rm -rf /binaries && go build -o ./binaries/app && docker build --no-cache -t throttle-service:latest .

dockerize-mac:
	rm -rf /binaries && env GOOS=linux GOARCH=amd64 go build -o ./binaries/app && docker build -t throttle-service:latest .

docker-run:
	docker run -p 1323:1323 -e ENV=PROD -it throttle-service:latest

docker-compose:
	docker-compose rm -f && docker-compose pull && docker-compose build --no-cache && docker-compose up
