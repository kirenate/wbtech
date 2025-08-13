# Start by building the application.
FROM golang:1.24-bullseye AS build

WORKDIR /src

# Copy only dependency files first
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the rest of the source code
COPY . .

RUN go build -o=bootstrap main.go

## Now copy it into our base image.
FROM debian:trixie

RUN rm -rf /var/lib/apt/lists/* \
    && apt-get update \
    && apt-get install -y --no-install-recommends ca-certificates curl
RUN rm -rf /var/lib/apt/lists/*

COPY --from=build /src/bootstrap /

CMD ["/bootstrap"]
