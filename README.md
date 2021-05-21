# Flint <img src="https://static.wikia.nocookie.net/minecraft/images/6/67/FlintNew.png" alt="flint" width="36"/>

Flint is a key/value store, where the key is the http request path, and the value is the http request body.

## Install

If you have go installed, run:

```bash
go install github.com/hyqe/flint@latest
```

to get help run:

```
flint -h
```

## run

run the server in verbose mode:
```
flint --verbose
```

### Examples

**PUT**

```
curl -X PUT 'http://localhost:2000/foo' \
-H 'Content-Type: application/json' \
-d'{
    "a": 1,
    "b": 2
}'
```

**GET**

```
curl 'http://localhost:2000/foo'
```

returns:
```
Content-Type: application/json
```
```
{
    "a": 1,
    "b": 2
}
```

**DELETE**

```
curl -X DELETE 'http://localhost:2000/foo'
```


## Docker

```
docker pull hyqe/flint:latest
```

run 

```bash
docker run \
    -p 2000:2000 \
    --name flint \
    hyqe/flint
```


run forever in the background
```bash
docker run \
    -d \
    -p 2000:2000 \
    --restart unless-stopped \
    --name flint \
    hyqe/flint
```