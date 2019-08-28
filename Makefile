USERNAME=mahendrakariya
IMAGE=mainsvc
VERSION=0.0.9

build:
	go build -o mainsvc

docker-build: build
	docker build -t $(USERNAME)/$(IMAGE):$(VERSION) .

docker-run: docker-build
	docker run -p 3005:8000 --network host $(USERNAME)/$(IMAGE):$(VERSION)

docker-push: docker-build
	docker push $(USERNAME)/$(IMAGE):$(VERSION)
