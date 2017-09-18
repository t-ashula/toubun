TARGET  := toubun
DEBUG_FLAG = $(if $(DEBUG),-debug)
VERSION ?= $(shell git describe --tags)
SRCS := *.go core/*.go cmd/*.go
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
	go test -v $$(glide nv)

clean:
	rm -f $(TARGET)

force: clean all

.PHONY: force clean test install-deps all
