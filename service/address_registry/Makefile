srcs = \
	package/srv/client.go \
	package/srv/config.go \
	package/srv/handle.go \
	package/srv/request.go \
	package/srv/response.go \
	package/srv/service.go \
	package/util/log.go \
	package/main.go

all: main
main: $(srcs)
	go build -a package/main.go

.PHONY: init
init:
	mkdir -p log

.PHONY: clean
clean:
	rm -f main
	rm -rf log
	rm -f utils/launcher

.PHONY:utils
utils: utils/launcher.c
	clang -std=c11 utils/launcher.c -o utils/launcher