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

.PHONY: run-server-benchmark
run-server-benchmark:
	$(eval DATE:=$(shell TZ=JST-9 date "+%Y%m%d_%H%M%S"))
	$(eval OUT_FILE_PATH:=benchmark/server/${DATE}.txt)
	docker compose exec -d server sh -c "touch ${OUT_FILE_PATH} && sysbench cpu --time=10 --threads=1 run >> "${OUT_FILE_PATH}""

.PHONY: run-server-benchmark-5-times
run-server-benchmark-5-times:
	@for i in {1..5}; do \
		echo "==>Running benchmark $$i/5..." ; \
		make run-server-benchmark ; \
		sleep 6 ; \
	done

.PHONY: run-client-benchmark
run-client-benchmark:
	$(eval DATE:=$(shell TZ=JST-9 date "+%Y%m%d_%H%M%S"))
	$(eval OUT_FILE_PATH:=benchmark/client/${DATE}.txt)
	docker compose exec -d client sh -c "touch ${OUT_FILE_PATH} && sysbench cpu --time=10 --threads=1 run >> "${OUT_FILE_PATH}""

.PHONY: run-client-benchmark-5-times
run-client-benchmark-5-times:
	@for i in {1..5}; do \
		echo "==>Running benchmark $$i/5..." ; \
		make run-client-benchmark ; \
		sleep 6 ; \
	done

