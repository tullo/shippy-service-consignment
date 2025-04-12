SHELL = /bin/bash -o pipefail

.PHONY: init
init:
	@go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	@go install github.com/micro/micro/v5/cmd/protoc-gen-micro@latest

docker-image: export DOCKER_BUILDKIT = 1
docker-image:
	@docker build -f dockerfile -t shippy-service-consignment .

docker-run:
	@docker run -p 50051:50051 shippy-service-consignment

# https://github.com/go-micro/generator
.PHONY: grcp
grcp:
	@echo '~ protoc OK'
	@protoc --proto_path=. --go_out=. --micro_out=. proto/consignment.proto

.PHONY: api
api:
	@echo '~ protoc OK'
	@protoc --proto_path=. --openapi_out=. proto/consignment.proto

micro-log:
	@micro logs -f shippy-service-consignment

micro-kill:
	@micro kill shippy-service-consignment
	@micro status

micro-run:
	@micro run .

micro-update:
	@micro update .
	@micro logs -f shippy-service-consignment

micro-status:
	@micro status

tidy:
	@go mod tidy

vendor: tidy
	@go mod vendor

# =========================================================================== #
# ================= CONSIGNMENT OPS ========================================= #
# =========================================================================== #

micro-help:
	@micro shippy.service.consignment --help

micro-consignment-create:
	@micro shippy.service.consignment shippingService createConsignment \
	--description 'CONSIGNMENT' --weight 33333 --vessel_id "123456789"

micro-consignment-getAll:
	@micro shippy.service.consignment shippingService getConsignments

go-mod-edit:
	go mod edit -replace github.com/tullo/shippy-service-user=../shippy-service-user
	go mod edit -replace github.com/tullo/shippy-service-vessel=../shippy-service-vessel

curl-consignment-create:
	curl http://localhost:8080/shippy.service.consignment/shippingService/createConsignment \
	-H "Content-Type:application/json" \
	-d @consignment.json 

curl-consignment-getAll:
	curl http://localhost:8080/shippy.service.consignment/shippingService/getConsignments \
	-H "Accept:application/json" \
