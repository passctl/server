FROM golang:1.21.3

WORKDIR /app

COPY public ./public
COPY *.sum ./
COPY *.mod ./
COPY *.go ./

RUN go build . 
CMD ["./server"]
