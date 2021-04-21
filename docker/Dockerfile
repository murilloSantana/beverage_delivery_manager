FROM golang:1.15.10-alpine
WORKDIR /usr/app
COPY . .
RUN cp cmd/main.go main.go
RUN go install
EXPOSE 5000
ENTRYPOINT beverage_delivery_manager
