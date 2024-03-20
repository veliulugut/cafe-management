run:
	go run ./cmd/api

swag:
	swag init -g ./cmd/api/main.go

test:
	go test ./...

generate:
	go run -mod=mod entgo.io/ent/cmd/ent generate --feature "sql/upsert","sql/execquery" ./ent/schema

docker:
	sudo docker container start 48b5501d1883