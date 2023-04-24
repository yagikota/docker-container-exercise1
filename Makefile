.PHONY: run-server
run-server:
	docker compose up -d --build server

.PHONY: run-client
run-client:
	docker compose up -d --build client

.PHONY: run-client-server
run-client-server: run-server run-client

# This command executes tcpdump and saves the captured packets in the "/captured" directory. It also runs the packet capture for 180 seconds and then automatically stops.
.PHONY: start-packet-capture
start-packet-capture:
	$(eval DATE:=$(shell TZ=JST-9 date "+%Y%m%d_%H%M%S"))
	docker compose exec -d server timeout 180  tcpdump -i eth0 -w "/captured/${DATE}.pcap"

# ===== for ubuntu =====

.PHONY: run-server-ubuntu
run-server-ubuntu:
	sudo docker compose up -d --build server

.PHONY: run-client-ubuntu
run-client-ubuntu:
	sudo docker compose up -d --build client

.PHONY: run-client-server-ubuntu
run-client-server-ubuntu: run-server run-client

# This command executes tcpdump and saves the captured packets in the "/captured" directory. It also runs the packet capture for 180 seconds and then automatically stops.
.PHONY: start-packet-capture-ubuntu
start-packet-capture-ubuntu:
	$(eval DATE:=$(shell TZ=JST-9 date "+%Y%m%d_%H%M%S"))
	sudo docker compose exec -d server timeout 180 tcpdump -i eth0 -w "/captured/${DATE}.pcap"
