FROM scratch
MAINTAINER Nicole Macfarlane

COPY postgres-go-makedb /postgres-go-makedb
ENTRYPOINT ["/postgres-go-makedb"]
