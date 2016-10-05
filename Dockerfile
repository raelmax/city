FROM alpine:latest

MAINTAINER Edward Muller <edward@heroku.com>

WORKDIR "/opt"

ADD .docker_build/city /opt/bin/city
ADD ./testdata /opt/testdata

CMD ["/opt/bin/city"]

