DEVNAME=clicker-back-dev
NAME=clicker-back
VERSION=1.0

dev-image:
	DOCKER_BUILDKIT=1 docker build -f ./build/Dockerfile --target develop -t $(DEVNAME):$(VERSION) .

dev-run:
	docker run -itd -p 80:80 --name $(DEVNAME) $(DEVNAME):$(VERSION)

dev-stop:
	docker rm -f $(DEVNAME)

dev-logs:
	docker logs $(DEVNAME)

main-image:
	DOCKER_BUILDKIT=1 docker build --secret id=credential,src=./build/.gitconfig --target main -f ./build/Dockerfile -t $(NAME):$(VERSION) .

main-start:
	docker run -itd -p 80:80 --name $(NAME) $(NAME):$(VERSION)

main-stop:
	docker rm -f $(NAME)

main-logs:
	docker logs $(NAME)