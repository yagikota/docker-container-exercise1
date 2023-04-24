.PHONY: run-server
run-server:
	docker-compose up -d --build server

.PHONY: run-client
run-client:
	docker-compose up -d --build client

.PHONY: run-client-server
run-client-server:
	run-server
	run-client

.PHONY: send-client-request
send-client-request:
	docker-compose run --rm client client/main
