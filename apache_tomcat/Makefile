default: run

clean:
	docker container stop tomcat001 httpd001
	docker container rm tomcat001 httpd001
	docker network rm net1

build:
	docker network create net1
	docker image build -f Dockerfile.apache -t try-docker/httpd:latest .
	docker image build -f Dockerfile.tomcat -t try-docker/tomcat:latest .

run: build
	docker container run -d --name tomcat001 --net net1 try-docker/tomcat:latest
	docker container run -d --name httpd001 --net net1 --publish 8085:80 try-docker/httpd:latest
