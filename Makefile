APP = trelldo
# WORKING_PATH = /go/src/$(APP)
# DOCKER_CMD = docker run --rm -it -e GOCACHE=/tmp --user $$(id -u):$$(id -g) -v $$PWD:$(WORKING_PATH) -v $$GOPATH/pkg:/go/pkg -v $$GOPATH/bin:/go/bin -w $(WORKING_PATH) golang:1.18-buster
WORKING_PATH = .
DOCKER_CMD =
VERSION ?= 0.0.1

$(APP): main.go
	$(DOCKER_CMD) go build -ldflags="-X 'main.version=$(VERSION)'" -o $(WORKING_PATH)/$(APP) main.go

build: $(APP)

run: build
	./$(APP) help

clean:
	rm -f $(APP)

fmt:
	$(DOCKER_CMD) gofmt -s -w $(WORKING_PATH)

release:
	goreleaser build --single-target --rm-dist --snapshot --skip-validate --skip-post-hooks -o $(WORKING_PATH)/$(APP)

watch:
	fd -e go | entr make --no-print-directory --always-make

install: build
	sudo cp $(APP) /usr/local/bin/$(APP)
	sudo chmod +x /usr/local/bin/$(APP)
