FROM alpine:latest AS base
WORKDIR /app/

FROM golang:latest AS build
WORKDIR /src/
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 go build -o /build/app main.go

FROM base AS final
WORKDIR /app/
COPY --from=build /build/app /app/app
ENTRYPOINT [ "/app/app" ]