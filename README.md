# wrk-connector

A Goland published.

## Requirements

```text
  - GNU Make ~>4.3
  - dotenv-cli (by local stage)
  - Go ~>1.13
  - Docker ~>18
  - Docker compose ~>1.28.0 (Optional)
```

### **Configuration**

```shell
  make config project-name=<name>
```

Note: This creates the initializer configuration, and a .env file with basic variables

# Running in local stage, default port: 3000

```shell
make start
```

### with dotenv-cli

```shell
dotenv -e .env make start
```

# Running with Docker, default port: 4000

### With environment file

```shell
docker-compose -f docker-compose.yml up -d --build
```


## A new container every time and then log output (debug):

```shell
docker-compose -f docker-compose.yml up -d --build --force-recreate; docker-compose -f docker-compose.yml logs -f
```

## Reference Links

+ [docker-compose Documentation](https://docs.docker.com/compose/)
