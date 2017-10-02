TARGET  := toubun
DEBUG_FLAG = $(if $(DEBUG),-debug)
VERSION ?= $(shell git describe --tags)
SRCS := *.go core/*.go cmd/*.go runner/*.go
has_glide := $(shell command -v glide 2> /dev/null)


$(TARGET): $(SRCS)
	go build -ldflags "-X github.com/t-ashula/toubun/core.Version=$(VERSION)" -o $(TARGET) .

all: $(TARGET)

install-deps:
ifeq ('', $(has_glide))
	go get -u -d github.com/Masterminds/glide
endif
	glide -q install
	glide -q up

test:
	go test -cover $$(glide nv)

test-cover:
	echo 'mode: atomic' > cover-all.out
	$(foreach pkg, $(shell go list ./...), \
		go test -coverprofile=cover.out -covermode=atomic $(pkg); \
		tail -n +2 cover.out >> cover-all.out; \
	)
	go tool cover -func=cover-all.out

clean:
	rm -f $(TARGET) cover-all.out cover.out

force: clean all

.PHONY: force clean test install-deps all
