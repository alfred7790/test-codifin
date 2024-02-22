APP_NAME = go_codifin
PKG_LIST = `go list ./... | grep -v /vendor`

.PHONY: dep
dep:
	@echo "Download dependencies..."
	@echo "GOPATH is:" $(GOPATH)
	go get -u github.com/swaggo/swag/cmd/swag

.PHONY: docs
docs:
	@echo "Building documentation"
	$(GOPATH)/bin/swag init -g ./cmd/codifin/main.go

.PHONY: build
build:
	@echo "Building binary"
	go build -o build/bin/$(APP_NAME) cmd/codifin/main.go

.PHONY: run
run: build
	./build/bin/$(APP_NAME)

.PHONY: docker-build
docker-build:
	docker build -t $(APP_NAME)_img .

.PHONY: docker-run
docker-run: docker-build
	docker-compose up -d go_codifin

.PHONY: db
db:
	@echo "Running db_products Data Base"
	docker-compose up -d db_products
	# need to make sure our db is up and available
	# 2 seconds
	sleep 2
	docker-compose ps


.PHONY: clean
clean:
	@echo "Cleaning up db, docs"
	docker-compose down
	rm -rf ./db
	rm -rf ./docs
	@echo "Cleaning up complete"

.PHONY: deploy
deploy: | build db docker-run
	sleep 2
	docker-compose ps