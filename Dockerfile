FROM golang:1.17
ADD bin/lccncamp /lccncamp
ENTRYPOINT ["/lccncamp"]