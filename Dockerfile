FROM golang:1.19 AS build
WORKDIR /build
COPY . ./
RUN CGO_ENABLED=0 GOOS=linux go build -o /arraytest .

FROM alpine:latest
COPY --from=build /arraytest .
CMD ["/arraytest"]
