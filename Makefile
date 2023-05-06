.PHONY: up-server-client
up-server-client:
	docker compose up -d --build

.PHONY:send-request
send-request:
	docker compose exec client client/main

.PHONY: start-server-packet-capture
start-server-packet-capture:
	$(eval DATE:=$(shell TZ=JST-9 date "+%Y%m%d_%H%M%S"))
	docker compose exec -d server timeout 180  tcpdump -i any -w "/captured/server/${DATE}.pcap"

# TODO: add commands for ubuntu
# ===== for ubuntu =====
