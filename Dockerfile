FROM golang:1.17-alpine
ADD bin/lccncamp /lccncamp
ENTRYPOINT ["/lccncamp"]