FROM golang:1.15.7-alpine3.13

WORKDIR /app

COPY . .

RUN apk add --no-cache make && \
    make prep_build && \
    make dep_install

CMD [ "make", "watch" ]
