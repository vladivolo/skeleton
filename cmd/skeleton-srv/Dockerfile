# syntax=docker/dockerfile:1

FROM golang:1.22-alpine AS build

# Install required system packages
RUN apk add --no-cache --update \
git \
gcc \
musl-dev

WORKDIR /src

# Copy project sources
COPY . .

# Build binary
RUN echo "replace github.com/vladivolo/skeleton => /src" >> go.mod
RUN go mod download

WORKDIR /src/cmd/skeleton-srv
RUN echo "replace github.com/vladivolo/skeleton => /src" >> go.mod
RUN go mod download
RUN go build

# Build target image
FROM golang:1.22-alpine AS release

WORKDIR /app

# Copy binary and config to target image
COPY --from=build /src/cmd/skeleton-srv/skeleton-srv /app/
COPY --from=build /src/cmd/skeleton-srv/conf/config.yaml /app/

#EXPOSE 11223
#CMD ["/app/skeleton-srv", "-stderr"]
