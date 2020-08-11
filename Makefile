run:
	docker-compose up --force-recreate --build

test:
	curl localhost:8888/ping
