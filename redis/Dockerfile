FROM redis:alpine

WORKDIR /app

COPY redis.conf ./

CMD [ "redis-server", "/app/redis.conf" ]
