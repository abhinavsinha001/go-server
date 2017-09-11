FROM golang:latest 
RUN mkdir /app 
RUN mkdir -p /etc/custom-metrics 
ADD . /app/ 
ADD ./definition.json /etc/custom-metrics/
WORKDIR /app 
RUN go get github.com/paulbellamy/ratecounter
RUN go build -o main . 
CMD ["/app/main"]
