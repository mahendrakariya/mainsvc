build:
	go build -o mainsvc
docker-build:
	docker build -t mainsvc .
docker-run:
	docker run -p 3005:8000 mainsvc
