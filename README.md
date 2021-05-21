# Flint

At its core, Flint is just a key/value store, where the key is the http request path, and the value is the http request body. Flint will also preserve the content type of the PUT request so that subsequent GET requests return the same value. 


## CMD
```
NAME:
   Flint

USAGE:
   main [global options] command [command options] [arguments...]

DESCRIPTION:
   A simple http key/value store

COMMANDS:
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --port value, -p value  (default: 2000)
   --verbose, -v           (default: false)
   --help, -h              show help (default: false)
```


## HTTP API

flint has only three end-points.

| Method | Path | Body |
| ------ | ---- | ---- |
| PUT    | *    | *    |
| GET    | *    | null |
| DELETE | *    | null |


### Examples

**PUT**

```
curl -X PUT 'http://localhost:2000/foo' \
-H 'Content-Type: application/json' \
-d'{
    "a": "1",
    "b": "2"
}'
```

**GET**

```
curl 'http://localhost:2000/foo'
```

**DELETE**

```
curl -X DELETE 'http://localhost:2000/foo'
```


## Docker

build

```bash
docker build . -t flint
```

run 

```bash
docker run \
    -p 2000:2000 \
    --name flint \
    flint
```


run forever
```bash
docker run \
    -d \
    -p 2000:2000 \
    --restart unless-stopped \
    --name flint \
    flint
```