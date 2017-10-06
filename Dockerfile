FROM golang:latest
RUN mkdir /app
ADD . /app/
WORKDIR /app
RUN go build -o delay-server .
CMD ["/app/delay-server"]