# delay-server

[![Deploy](https://www.herokucdn.com/deploy/button.svg)](https://heroku.com/deploy?template=https://github.com/nexus-uw/delay-server)

## purpose
super simple http server that will allow you 'mock' delays in HTTP requires (ie: dont know what the request will be, just that you need to have a little delay and a real connection)

## running (only works on amd64 docker arch)
```
docker run -p 3000:3000 nexusuw/delay-server:latest
```
then vist [localhost:3000](http://localhost:3000)
