run:
	go run ./cmd/api

docker-build:
	docker build -t admin-api .

docker-run:
	docker run --network=host -it -p 8080:8080 admin-api
