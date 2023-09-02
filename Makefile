.PHONY: init

init:
	go install go.uber.org/mock/mockgen@bb5901fe6e45c7c5035afb29a274b9e970c8e348
	go install github.com/google/wire/cmd/wire@0ac845078ca01a1755571c53d7a8e7995b96e40d
	go install github.com/deepmap/oapi-codegen/cmd/oapi-codegen@7aa85bb88223ee606c2aaeb3e536aa0ed93d4054
	go install github.com/spf13/cobra-cli@74762ac083f2c4deffef229c887ffc15beb6ce0d

.PHONY: openapi_http
openapi_http:
	oapi-codegen -generate types -o "interface/rest_api/openapi_types.gen.go" -package "api" "asset/swagger/swagger.yaml"
	oapi-codegen -generate server -o "interface/rest_api/openapi_api.gen.go" -package "api" "asset/swagger/swagger.yaml"

.PHONY: di
di:
	wire gen

.PHONY: generate
generate: openapi_http di

