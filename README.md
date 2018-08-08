# go-numbers-api
This is a sample web api service which given a number of other services which return a list of numbers collects numbers from the services , merge, sort and present them as json.

## Build & Run Instructions
* Clone this repo
* Run it `go run main.go intset.go controllers.go service.go`
* Open your browser and test the service at http://localhost:8080/numbers?u=http://localhost:8090/primes&u=http://localhost:8090/fibo
