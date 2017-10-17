NAME   := toubun
VERSION  := v0.1.0
REVISION := $(shell git rev-parse --short HEAD)

SRCS     := *.go core/*.go cmd/*.go runner/*.go runner/**/*.go
LDFLAGS  := "-X github.com/t-ashula/toubun/core.Version=$(VERSION)"

glide:
ifeq ($(shell command -v glide 2> /dev/null),)
	go get -u github.com/Masterminds/glide
endif

deps: glide
	glide install

$(NAME): $(SRCS)
	go build -ldflags $(LDFLAGS) -o $(NAME)

all: $(NAME)

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
	rm -rf $(NAME) cover-all.out cover.out bin vendor

force: clean all

install:
	go install

.PHONY: force clean test-cover test all deps glide
