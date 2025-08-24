# Go REST API Starter Pack

* Golang 
* Gin Framework
* MVC
* Hot Reloading with Air
* Dependency Injection with Wire
* Dockerized
* MySQL
* Redis
* RabbitMQ

### Installation

`````bash
git clone
docker compose up --build

# without docker
cd internal/dependencies 
wire
cd cmd/app
go run main.go
`````

### Create a new CRUD 

* 1- Create a new dependency injection structure in internal/dependencies/wire.go
* 2- Run wire ./internal/dependencies command to generate wire_gen.go file
* 3- Create new route file in /internal/routes.go
* 4- Trigger your recently created route file in internal/routes/initialize.go
* 5- Add your model, handler, service, repository files.
* 6- If your APP_ENV value in .env file is "development", add new AutoMigrate in internal/config/database.go 