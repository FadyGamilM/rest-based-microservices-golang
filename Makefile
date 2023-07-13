DSN="host=localhost port=1122 user=fady password=fady dbname=bankingdb sslmode=disable timezone=UTC connect_timeout=5"
PORT=5050

# the container for the database
DB_DOCKER_CONTAINER=banking_db_container

# when we need to deploy our api to for example to an EC2 instance, we will need to build a binary for example to 
# linux and run this binary, so we will name this binary from here 
BINARY_NAME=bankingapi

# to create a running container instance for postgres using the name of the container we specified above
postgres:
	docker run -d --name ${DB_DOCKER_CONTAINER} -p 1122:5432 -e POSTGRES_USER=fady -e POSTGRES_PASSWORD=fady postgres:14

# now lets create an actual datbase inside this container 
createdb:
	docker exec -it ${DB_DOCKER_CONTAINER} createdb --username=fady --owner=fady bankingdb

migrate_up:
	docker run -i -v "H:\1- freelancing path\Courses\golang stack\projects\rest-based-microservices-golang\migrations:/migrations" --network host migrate/migrate -path=/migrations/ -database "postgresql://fady:fady@127.0.0.1:1122/bankingdb?sslmode=disable" up 1

migrate_down:
	docker run -i -v "H:\1- freelancing path\Courses\golang stack\projects\rest-based-microservices-golang\migrations:/migrations" --network host migrate/migrate -path=/migrations/ -database "postgresql://fady:fady@127.0.0.1:1122/bankingdb?sslmode=disable" down 1


build:
	@echo " + Building the binary of the backend ... "
	go build -o ${BINARY_NAME} .
	@echo " + The binary are ready !"

start_docker:
	@echo " + starting the db docker container"
	docker start ${DB_DOCKER_CONTAINER}

run: build start_docker
	@echo " + starting the api"
	@env DSN=${DSN} ./${BINARY_NAME} &
	@echo " + api started"
	

# i have to run this with unix-based system 
stop:
	@echo " + stopping the running api service"
	@-pkill -SIGTERM -f "./${BINARY_NAME}"
	@echo " + api service is stopped successfully!"

restart: stop run



# this is how to handle the dirty schema_migrations table >    UPDATE schema_migrations SET dirty = false WHERE version = 1;



