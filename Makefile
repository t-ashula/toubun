TARGET   := toubun
VERSION  := v0.1.0
REVISION := $(shell git rev-parse --short HEAD)

SRCS     := *.go core/*.go cmd/*.go runner/*.go runner/**/*.go
LDFLAGS  := "-X github.com/t-ashula/toubun/core.Version=$(VERSION)"

glide:
ifeq ($(shell command -v glide 2> /dev/null),)
	go get -u -d github.com/Masterminds/glide
endif

deps: glide
	glide install

$(TARGET): $(SRCS)
	go build -ldflags $(LDFLAGS) -o $(TARGET)

all: $(TARGET)

test:
	go test -cover $$(glide nv)

test-cover:
	echo 'mode: atomic' > cover-all.out
	$(foreach pkg, $(shell glide nv), \
		go test -coverprofile=cover.out -covermode=atomic -v $(pkg); \
		tail -n +2 cover.out >> cover-all.out; \
	)
	go tool cover -func=cover-all.out

clean:
	rm -rf $(TARGET) cover-all.out cover.out bin vendor

force: clean all

.PHONY: force clean test deps glide
