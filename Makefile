.PHONY: run-server
run-server:
	docker-compose up -d --build server

.PHONY: run-client
run-client:
	docker-compose up -d --build client

.PHONY: run-client-server
run-client-server: run-server run-client

.PHONY: send-client-request
send-client-request:
	docker-compose run --rm client client/main

# This command executes tcpdump and saves the captured packets in the "/captured" directory. It also runs the packet capture for 180 seconds and then automatically stops.
.PHONY: start-packet-capture
start-packet-capture:
	docker-compose exec -d server timeout 180 tcpdump -i eth0 -w "/captured/$(date +'%Y%m%d_%H%M%S').pcap"
