# Stage: Depedency download
###########################
FROM golang:1.23 AS dependecy-stage

# Copy go.mod and go sum
WORKDIR /app
COPY /source/go.mod ./
COPY /source/go.sum ./
RUN go mod download

# Stage: Depedency download
###########################
FROM dependecy-stage AS build-stage

# Copy all source code
WORKDIR /app
COPY /source ./

RUN CGO_ENABLED=0 GOOS=linux go build -o /server

# Stage: Runner application
###########################
FROM alpine:3.18 AS runner-stage

# Work directory
WORKDIR /
COPY /source/ /
COPY --from=build-stage /server /server
RUN chown nobody.nobody server
RUN apk --no-cache add tzdata

EXPOSE 8080
USER nobody

ENTRYPOINT ["/server"]
