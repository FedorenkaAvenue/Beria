prep_build:
	@echo "Preparing installation environment...."
	go get -u github.com/cespare/reflex

dep_install:
	@echo "Installing dependencies...."
	go mod download

run:
	@echo "Running...."
	go run ./src/index.go

build:
	@echo "Building...."
	go build -o server ./src/index.go

watch:
	@echo "Restarting..."
	ulimit -n 1000
	${GOPATH}/bin/reflex --start-service -r '\.go$$' make run
