.PHONY: all run up down mkdirs clean

SHELL:=/bin/bash
PROJECT=log-arch
INTERVAL=5

all: setup up run

run:
	cd app && \
		go build -o $(PROJECT) main.go && \
		./$(PROJECT) $(INTERVAL) >> logs/app.log 2>&1

setup:
	sudo sysctl -w vm.max_map_count=262144
	mkdir -p infra/data/{redis,filebeat,elasticsearch,logstash}
	sudo chown -R 1000:1000 infra/data/elasticsearch

up:
	docker-compose -p $(PROJECT) -f infra/docker-compose.yml up -d

down:
	docker-compose -p $(PROJECT) -f infra/docker-compose.yml down --remove-orphans

logs:
	docker-compose -p $(PROJECT) -f infra/docker-compose.yml logs -t -f

clean: down
	sudo rm -rf infra/data/{redis,filebeat,elasticsearch,logstash}/*