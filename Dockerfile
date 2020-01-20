# Docker Definition for Preflight

FROM busybox:ubuntu-14.04
MAINTAINER Christian R. Vozar <cvozar@xumak.com>

ADD cmd/preflight/preflight /preflight
RUN chmod +x /preflight

CMD ["/preflight"]
