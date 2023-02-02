# MATCHING-TIMESTAMPS

This service follows domain-driven design architecture. The main application can be found in cmd/main.go. In presentation, app and persistence folders there are the handlers, the services and the repositories respectively. Inside the first two there are fakes folders in which you can found fake implementations of methods as obtained by using counterfeiter package. Furthermore, unit tests have been implemented and can be found in there as well. The domain folder contains the definitions of structures and interfaces.  

## Environment Variables
`SERVE_ON_PORT`: The port that this service runs on. The default value is 8000.

## Unit Tests
- `Counterfeiter` package is used in order to generate fake implementations of the methods for both service and repository.
- Unit tests for both handlers and services have been implemented in presentation and app/services folders respectively. 
- Run `go test ./...` in command line to run all the tests at once.

## Docker
- Run `docker build -t matching-timestamps .` to build an image for this application.
- Run `docker run -p 8000:8000 matching-timestamps` to run the service on port 8000.

## How to Run
You can either run the main application in a GoLand environment or run the docker image.

