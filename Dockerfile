FROM scratch
MAINTAINER Brian Hechinger <wonko@4amlunch.net>

ADD esi-srv-linux-amd64 esi-srv
VOLUME /etc/chremoas

ENTRYPOINT ["/esi-srv", "--configuration_file", "/etc/chremoas/chremoas.yaml"]
