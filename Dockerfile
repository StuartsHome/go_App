# sets the base image
FROM golang:1.16 AS base

ENV GOSUMDB="off"

WORKDIR /src/

# ==================
# Dev Container
FROM base as dev

# RUN echo "$PWD"
# Cache dependencies
COPY go.mod go.sum /src/
RUN go mod download -x


# ==================
# Build Container

From dev AS build

ENV CGO_ENABLED=0
ENV GOOS=linux

# Build executables
COPY . /src/
RUN go install -v ./...
