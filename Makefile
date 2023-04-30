.PHONY: up-server
up-server:
	docker compose up -d --build server

.PHONY: build-client
build-client:
	docker compose build client

.PHONY:start-client
start-client:
	docker compose start client

.PHONY: start-server-packet-capture
start-server-packet-capture:
	$(eval DATE:=$(shell TZ=JST-9 date "+%Y%m%d_%H%M%S"))
	docker compose exec -d server timeout 180  tcpdump -i eth0 -w "/captured/server/${DATE}.pcap"

.PHONY: start-client-packet-capture
start-client-packet-capture:
	$(eval DATE:=$(shell TZ=JST-9 date "+%Y%m%d_%H%M%S"))
	docker compose exec -d client timeout 180  tcpdump -i eth0 -w "/captured/client/${DATE}.pcap"

# ===== for ubuntu =====
