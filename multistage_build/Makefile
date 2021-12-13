build: \
	copy-gomod \
	build-nostage \
	build-multistage \
	rm-gomod

copy-gomod:
	cp ../go.* .

rm-gomod:
	rm go.*

build-nostage:
	docker buildx build -t multistage-ex:nostage -f Dockerfile.nostage ${PWD}

build-multistage:
	docker buildx build -t multistage-ex:multistage -f Dockerfile.multistage ${PWD}

run:
	time -p docker container run --rm multistage-ex:nostage
	time -p docker container run --rm multistage-ex:multistage

clean:
	docker image rm multistage-ex:nostage multistage-ex:multistage
	docker image prune -f
