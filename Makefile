.PHONY: up-server-client
up-server-client:
	docker compose up -d --build

.PHONY:send-request
send-request:
	docker compose exec client client/main

# ===== packat capture =====
.PHONY: start-server-packet-capture
start-server-packet-capture:
	$(eval DATE:=$(shell TZ=JST-9 date "+%Y%m%d_%H%M%S"))
	docker compose exec -d server timeout 180  tcpdump -i any -w "/captured/server/${DATE}.pcap"

# ===== cpu benchmark =====
.PHONY: run-server-benchmark-cpu
run-server-benchmark-cpu:
	$(eval DATE:=$(shell TZ=JST-9 date "+%Y%m%d_%H%M%S"))
	$(eval OUT_FILE_PATH:=benchmark/server/${DATE}.txt)
	docker compose exec -d server sh -c "touch ${OUT_FILE_PATH} && sysbench cpu --time=10 --threads=1 run >> "${OUT_FILE_PATH}""

.PHONY: run-server-benchmark-cpu-5-times
run-server-benchmark-cpu-5-times:
	@for i in {1..5}; do \
		echo "==>Running benchmark $$i/5..." ; \
		make run-server-benchmark-cpu ; \
		sleep 6 ; \
	done

.PHONY: run-client-benchmark-cpu
run-client-benchmark-cpu:
	$(eval DATE:=$(shell TZ=JST-9 date "+%Y%m%d_%H%M%S"))
	$(eval OUT_FILE_PATH:=benchmark/client/${DATE}.txt)
	docker compose exec -d client sh -c "touch ${OUT_FILE_PATH} && sysbench cpu --time=10 --threads=1 run >> "${OUT_FILE_PATH}""

.PHONY: run-client-benchmark-cpu-5-times
run-client-benchmark-cpu-5-times:
	@for i in {1..5}; do \
		echo "==>Running benchmark $$i/5..." ; \
		make run-client-benchmark-cpu ; \
		sleep 6 ; \
	done

# ===== memory benchmark =====
.PHONY: run-server-benchmark-memory
run-server-benchmark-memory:
	$(eval DATE:=$(shell TZ=JST-9 date "+%Y%m%d_%H%M%S"))
	$(eval OUT_FILE_PATH:=benchmark/server/${DATE}.txt)
	docker compose exec -d server sh -c "touch ${OUT_FILE_PATH} && sysbench memory --time=10 --threads=1 run >> "${OUT_FILE_PATH}""

.PHONY: run-server-benchmark-memory-5-times
run-server-benchmark-memory-5-times:
	@for i in {1..5}; do \
		echo "==>Running benchmark $$i/5..." ; \
		make run-server-benchmark-memory ; \
		sleep 6 ; \
	done

.PHONY: run-client-benchmark-memory
run-client-benchmark-memory:
	$(eval DATE:=$(shell TZ=JST-9 date "+%Y%m%d_%H%M%S"))
	$(eval OUT_FILE_PATH:=benchmark/client/${DATE}.txt)
	docker compose exec -d client sh -c "touch ${OUT_FILE_PATH} && sysbench memory --time=10 --threads=1 run >> "${OUT_FILE_PATH}""

.PHONY: run-client-benchmark-memory-5-times
run-client-benchmark-memory-5-times:
	@for i in {1..5}; do \
		echo "==>Running benchmark $$i/5..." ; \
		make run-client-benchmark-memory ; \
		sleep 6 ; \
	done
