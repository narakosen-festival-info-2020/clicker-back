DEVNAME=clicker-back-dev
VERSION=1.0

dev-image:
	docker build -f ./build/Dockerfile --target develop -t $(DEVNAME):$(VERSION) .

dev-run:
	docker run -itd -p 80:80 --name $(DEVNAME) $(DEVNAME):$(VERSION)

dev-stop:
	docker rm -f $(DEVNAME)

dev-logs:
	docker logs $(DEVNAME)