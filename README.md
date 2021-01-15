# Logging Architecture

This project aims to simulate and explore a logging architecture for multiple applications, using mainly ELK for visualizing and treating logs.

# Getting started

## Prerequisites

- A Linux machine (tested on Ubuntu LTS 20.04);
- [Docker] and [post-install steps](docker-post-install);
- [Docker Compose](docker-compose);
- [Go] version 1.15+;
  - Recommended to install Go with [GVM];
- Linux `build-essential` package:
  ```bash
  sudo apt-get install -y build-essential
  ```

After installing all of these prerequisites, you can start the simulation by running the following command at the root of this project (where the `Makefile` is located):
```bash
make
```

Then you will have Redis, Filebeat, Logstash, ElasticSearch and Kibana running as services in Docker, and the app project will keep generating logs indefinitely at the folder `app/logs/`. After Kibana and ElasticSearch is up, and also after there is enough log generation, you can access the url `localhost:5601` to access Kibana management console and study the created indexes (which are named `app-dev-YYYY.MM.DD` and `app-entity-dev-YYYY.MM.DD`).


[redis]: https://redis.io/
[docker]: https://docs.docker.com/engine/install/ubuntu/
[docker-post-install]: https://docs.docker.com/engine/install/linux-postinstall/
[docker-compose]: https://docs.docker.com/compose/install/
[go]: https://golang.org/dl/
[gvm]: https://github.com/moovweb/gvm