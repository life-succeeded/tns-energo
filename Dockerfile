FROM golang:1.24.1
WORKDIR /app
EXPOSE 8080

COPY . .
RUN chmod +x ./start.sh

RUN go install -mod vendor

ENTRYPOINT ["./start.sh"]