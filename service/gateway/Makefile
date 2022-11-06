srcs = \
	package/anomalies/handler.go \
	package/auth/access.go \
	package/auth/auth_info.go \
	package/auth/handler.go \
	package/gtw/config.go \
	package/info/handler.go \
	package/statistic/handler.go \
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