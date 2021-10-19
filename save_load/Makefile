default: run

clean:
	docker image rm -f try-docker/save_load:latest

build:
	docker image build -f Dockerfile -t try-docker/save_load:latest .

run: build
	docker image save -o saveload001.tar try-docker/save_load:latest
	@echo '---------------------------------------------------------'
	ls -lh
	@echo '---------------------------------------------------------'
	docker image rm -f try-docker/save_load:latest
	docker image list
	@echo '---------------------------------------------------------'
	docker image load --input saveload001.tar
	docker image list
