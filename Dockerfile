FROM golang:latest 
RUN mkdir /app 
ADD . /app/ 
WORKDIR /app 
RUN go get github.com/paulbellamy/ratecounter
RUN go build -o main . 
CMD ["/app/main"]
