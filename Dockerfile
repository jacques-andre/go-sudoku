FROM golang:latest

# create /app dir, everything is housed here
RUN mkdir /app

# copy pwd files into /app dir
ADD . /app

# specify commands should be run from /app
WORKDIR /app

RUN make
CMD ["./bin/main"]
