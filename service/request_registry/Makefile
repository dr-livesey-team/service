srcs = \
	package/rtr/anomaly_id.go \
	package/rtr/anomaly_insert_info.go \
	package/rtr/anomaly_select_info.go \
	package/rtr/client.go \
	package/rtr/config.go \
	package/rtr/filter.go \
	package/rtr/handler.go \
	package/rtr/request.go \
	package/rtr/request_insert_info.go \
	package/rtr/request_select_info.go \
	package/rtr/request_registry.go \
	package/rtr/response.go \
	package/rtr/server.go \
	package/rtr/space_desc.go \
	package/rtr/statistic.go \
	package/util/io.go \
	package/util/log.go \
	package/main.go

all: main
main: $(srcs)
	go build -a package/main.go

.PHONY: install
install:
	mkdir -p log

.PHONY: clean
clean:
	rm -f main
	rm -rf log
	rm -f utils/launcher

.PHONY:utils
utils: utils/launcher.c
	clang -std=c11 utils/launcher.c -o utils/launcher