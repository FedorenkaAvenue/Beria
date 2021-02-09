FROM golang:1.15.7-alpine3.13

WORKDIR /app

COPY . .

RUN make install

CMD [ "make", "watch" ]
