
build:
	docker image build -t alpine/tzutc -f Dockerfile.tzutc .
	docker image build -t alpine/tzja -f Dockerfile.tzja .

run: build
	docker container run -it --rm alpine/tzutc /bin/sh -c "date"
	docker container run -it --rm alpine/tzja /bin/sh -c "date"
