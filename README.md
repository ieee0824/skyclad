# skyclad
skyclad observes the docker container.
Notify you when there are old containers in time series.

# build
```
$ docker build -t skyclad .
```

# run
```
$ docker run --rm --name skyclad -v /var/run/docker.sock:/var/run/docker.sock skyclad
```