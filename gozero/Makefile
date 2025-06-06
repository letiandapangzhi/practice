update_rbac_api:
	goctl api go -api ./app/rbac/api/desc/*.api -dir ./app/rbac/api/

update_rbac_rpc:
	goctl rpc protoc ./app/rbac/rpc/desc/*.proto --go_out=./app/rbac/rpc/ --go-grpc_out=./app/rbac/rpc/ --zrpc_out=./app/rbac/rpc/ --client=true