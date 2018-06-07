# skyclad
skyclad observes the docker container.  
Notify you when there are old containers in time series.

# docker pull
```
$ docker pull ieee0824/skyclad
```

# docker build
```
$ docker build -t skyclad .
```

# run
```
# output stdout
$ docker run --rm --name skyclad -v /var/run/docker.sock:/var/run/docker.sock skyclad
```
or
```
# output slack
$ docker run --rm --name skyclad -v /var/run/docker.sock:/var/run/docker.sock skyclad -slack-api ${SLACK_KEY} -n slack
```

# help
```
$ skyclad -h
```

# License
MIT