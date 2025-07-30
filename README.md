# Flint <img src="https://static.wikia.nocookie.net/minecraft/images/6/67/FlintNew.png" alt="flint" width="36"/>

Flint is a key/value store, where the key is the http request path, and the value is the http request body.

Flint has three http end-points. Flint will preserve the Content-Type 
of the PUT request. The path is the key. The body is the value.

| Method | Path | Body | Description               |
| ------ | ---- | ---- | ------------------------- |
| PUT    | *    | *    | create a key/value        |
| GET    | *    | none | get a value by its key    |
| DELETE | *    | none | delete a value by its key |

## Quick Start

```
docker run -p 2000:2000 hyqe/flint --verbose
```

put a value

```bash
curl -X PUT 'http://localhost:2000/foo' -H 'Content-Type: application/json' -d'{"foo": "bar"}' -i
```

get the value back

```bash
curl 'http://localhost:2000/foo' -i
```

```
Content-Type: application/json

{"foo": "bar"}
```

delete the value

```bash
curl -X DELETE 'http://localhost:2000/foo' -i
```


## Docker

```bash
docker pull hyqe/flint:latest
```

Run temporarily in the terminal. 

```bash
docker run -p 2000:2000 hyqe/flint --verbose
```


Run forever in the background with storage that persists after reboots.

```bash
mkdir -p ~/.flint

docker run \
    -d \
    -p 2000:2000 \
    -v ~/.flint:/app/cache \
    --restart unless-stopped \
    --name flint \
    hyqe/flint:latest \
        --storage cache \
        --verbose
```