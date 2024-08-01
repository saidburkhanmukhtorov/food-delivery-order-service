gen-proto:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	export PATH="$PATH:$(go env GOPATH)/bin"

	protoc --go_out=./ \
    --go-grpc_out=./ \
	submodule-food-delivery/order/*.proto

mig-run:
	migrate create -ext sql -dir migrations -seq food_delivery 