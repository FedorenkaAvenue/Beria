FROM redis:alpine

WORKDIR /app

COPY redis.conf ./
COPY ./backup/dump.rdb /app

CMD [ "redis-server", "/app/redis.conf" ]
