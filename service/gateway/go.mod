module github.com/dr-livesey-team/service/service/gateway

go 1.19

replace github.com/dr-livesey-team/service/service/address_registry => ../address_registry

replace github.com/dr-livesey-team/service/service/request_registry => ../request_registry

require github.com/dr-livesey-team/service/service/request_registry v0.0.0-00010101000000-000000000000

require (
	github.com/dr-livesey-team/service/service/address_registry v0.0.0-00010101000000-000000000000 // indirect
	github.com/golang/protobuf v1.3.1 // indirect
	github.com/mattn/go-pointer v0.0.1 // indirect
	github.com/spacemonkeygo/spacelog v0.0.0-20180420211403-2296661a0572 // indirect
	github.com/tarantool/go-openssl v0.0.8-0.20220711094538-d93c1eff4f49 // indirect
	github.com/tarantool/go-tarantool v1.9.0 // indirect
	github.com/vmihailenco/msgpack/v5 v5.3.5 // indirect
	github.com/vmihailenco/tagparser/v2 v2.0.0 // indirect
	golang.org/x/net v0.0.0-20201021035429-f5854403a974 // indirect
	golang.org/x/sys v0.0.0-20210119212857-b64e53b001e4 // indirect
	google.golang.org/appengine v1.6.7 // indirect
	gopkg.in/vmihailenco/msgpack.v2 v2.9.2 // indirect
)
