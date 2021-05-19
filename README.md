# wrk-connector

A Goland published.

## Requirements

```text
  - Yarn
  - dotenv-cli (by local stage)
  - Go ~>1.13
  - Docker ~>18
  - Docker compose ~>1.28.0 (Optional)
```

# Running in local stage, default port: 3000

```shell
yarn run start:dev
```

### with dotenv-cli

```shell
dotenv -e .env yarn run start:dev
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
