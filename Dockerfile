FROM golang:1.20.5

WORKDIR /app

COPY public ./public
COPY *.sum ./
COPY *.mod ./
COPY *.go ./

RUN go build . 
CMD ["./server"]
