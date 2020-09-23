FROM golang:1.13

LABEL maintainer="Shubham Dhanera<tshubham19@gmail.com>"
LABEL tag="go-adventure-land-api"

WORKDIR /src
# Copy over the app files
COPY . /src

RUN go build -o main

EXPOSE 8080

CMD ["./main"]
