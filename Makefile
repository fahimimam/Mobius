FRONT_END_BINARY=frontApp
BROKER_BINARY=brokerApp
AUTH_BINARY=authApp

## up: starts all containers in the background without forcing build
up:
	@echo "Starting Docker images..."
	docker-compose up -d
	@echo "Docker images started!"

## up_build: stops docker-compose (if running), builds all projects and starts docker compose
up_build: build_broker build_auth
	@echo "Stopping docker images (if running...)"
	docker-compose down
	@echo "Building (when required) and starting docker images..."
	docker-compose up --build -d
	@echo "Docker images built and started!"

## down: stop docker compose
down:
	@echo "Stopping docker compose..."
	docker-compose down
	@echo "Done!"

## build_broker: builds the broker binary as a linux executable
build_broker:
	@echo "Building broker binary..."
	cd ./broker-service && env GOOS=linux CGO_ENABLED=0 go build -o ${BROKER_BINARY} ./cmd/api
	@echo "Done!"

## build_auth: builds the auth binary as a linux executable
build_auth:
	@echo "Building auth binary..."
	cd ./authentication-service && env GOOS=linux CGO_ENABLED=0 go build -o ${AUTH_BINARY} ./cmd/api
	@echo "Done!"

## build_front: builds the frone end binary
build_front:
	@echo "Building front end binary..."
	cd ./front-end && env CGO_ENABLED=0 go build -o ${FRONT_END_BINARY} .
	@echo "Done!"

## start: starts the front end
start: build_front
	@echo "Starting front end"
	cd ./front-end && ./${FRONT_END_BINARY} &

## stop: stop the front end
stop:
	@echo "Stopping front end..."
	@-pkill -SIGTERM -f "./${FRONT_END_BINARY}"
	@echo "Stopped front end!"

help:
	@echo "##################################################################################################################"
	@echo "#                                            Command Center                                                      #"
	@echo "##################################################################################################################"
	@echo "#      Command        |                                   Description                                            #"
	@echo "#---------------------|------------------------------------------------------------------------------------------#"
	@echo "# 👉 make up          | Fire up all containers and dive into action!     🔥                                      #"
	@echo "# 👉 make up_build    | Halt docker-compose (if active), craft all projects, and kickstart docker-compose! 🛠️   #"
	@echo "# 👉 make down        | Pull the plug on docker compose when it's time to take a breather. 🛑                    #"
	@echo "# 👉 make build_broker| Forge the broker binary into a formidable Linux executable! ⚙️                           #"
	@echo "# 👉 make build_auth  | Craft the authentication microservice, securing your system! 🔐️                         #"
	@echo "# 👉 make build_front | Assemble the front end binary with precision and finesse! 🎨                             #"
	@echo "# 👉 make start       | Launch the front end and let the journey begin!      🚀                                  #"
	@echo "# 👉 make stop        | Halt the front end gracefully, wrapping up today's adventures.  🛑                       #"
	@echo "# 👉 make help        | Seek guidance? Display these vibrant commands once more!     ℹ️                          #"
	@echo "##################################################################################################################"

.PHONY: up up_build down build_broker build_front start stop help