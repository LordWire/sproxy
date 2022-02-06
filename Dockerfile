FROM golang:1.17-bullseye AS build
WORKDIR /app
COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY *.go ./
RUN go build -o /proxy

FROM gcr.io/distroless/base-debian11
WORKDIR /app
COPY --from=build /proxy /app/proxy
EXPOSE 8080
CMD [ "/app/proxy" ]
