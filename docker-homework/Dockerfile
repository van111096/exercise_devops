FROM golang:alpine as bookApi
LABEL maintainer="Hoang ANH VAN <hoanganhvan1996@gmail.com>"

#set the current working directory inside the container
WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN go mod download
# copy source code from the local to container
COPY . .

#build the book api
RUN go build -o main .

#-----------------

# Start a new stage from scratch
FROM alpine:latest
WORKDIR /app
# Copy the Pre-build binary from the previous stage. Observer e also copy .evn file
COPY --from=bookApi /app/main .
COPY ./config/app/.env .
COPY ./config/app/app.conf .
CMD ["./main"]