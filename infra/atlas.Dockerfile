FROM alpine:3.20.2

COPY --from=arigaio/atlas:0.27.0-alpine /atlas atlas

COPY migrations/ migrations
COPY atlas.hcl atlas.hcl

CMD [ "./atlas", "migrate", "apply", "--env", "docker" ]