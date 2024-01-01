FROM golang:1.20-alpine

WORKDIR /project

COPY . .
RUN go mod download
RUN mkdir /build
RUN cp .env* /build

RUN go build -o /build/app
WORKDIR /build

RUN rm -rf /project

CMD [ "./app" ]

