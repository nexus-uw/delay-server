FROM alpine:latest

MAINTAINER Edward Muller <edward@heroku.com>

WORKDIR "/opt"

ADD .docker_build/delay-server /opt/bin/delay-server


CMD ["/opt/bin/delay-server"]

