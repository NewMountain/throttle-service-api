# throttle-service-api

Scaffold for a throttling microservice in Go

## Creature Comforts

Some things to make your life easier:

1. `make run`: Go run this project
2. `make test`: Tests verbose with coverage
3. `make report`: Creates a nice little code cov report in your browser
4. `make dockerize`: Builds the project and docker builds it for you as `throttle-service:latest`
5. `make dockerize-mac`: For those who dev on mac and host on Lin
6. `make docker-run`: Confirm it's not vaporware on Locahost.

## Questions

Why Centos?

My firm loves all things RHEL, and I like stuff that works at my shop.

## How to deploy?

If you're on a Macbook:

1. `make dockerize-mac`
2. `make docker-compose`
