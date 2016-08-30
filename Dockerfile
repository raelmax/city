FROM alpine:latest

MAINTAINER Edward Muller <edward@heroku.com>

WORKDIR "/opt"

ADD .docker_build/city /opt/bin/city
ADD ./test_data /opt/test_data

CMD ["/opt/bin/city"]

