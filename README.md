# Flint



## Docker

build

```bash
docker build . -t flint
```

run 

```bash
docker run \
    -p 1389:1389 \
    --name flint \
    flint
```


run forever
```bash
docker run \
    -d \
    -p 1389:1389 \
    --restart unless-stopped \
    --name flint \
    flint
```