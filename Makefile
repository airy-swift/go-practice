init:
	docker build --no-cache -t go_test .

start:
	docker run --rm --publish 8080:8080 --name test go_test

del:
	docker rmi go_test
