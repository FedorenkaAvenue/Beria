prep_build:
	@echo "\nPreparing installation environment...."
	go get github.com/cespare/reflex

dep_install:
	@echo "\nInstalling dependencies...."
	go mod download

run:
	@echo "\nRunning...."
	go run ./src/index.go

build:
	@echo "\nBuilding...."
	go build -o server ./src/index.go

watch:
	ulimit -n 1000
	bin/reflex --start-service -r '\.go$$' make run

install:
	$(MAKE) prep_build
	# $(MAKE) dep_install
