FROM golang:1.15.10-alpine
WORKDIR /usr/app
COPY . .
RUN go build -o bin/beverage_delivery_service
EXPOSE 5000
CMD [ "./bin/beverage_delivery_service" ]