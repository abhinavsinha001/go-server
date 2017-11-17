FROM golang:latest 
RUN mkdir /app 
ADD . /app/ 
WORKDIR /app 
RUN go get github.com/prometheus/client_golang/prometheus/promhttp github.com/prometheus/client_golang/prometheus
RUN go build -o main . 
CMD ["/app/main"]
