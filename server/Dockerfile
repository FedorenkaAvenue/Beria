FROM node:14.15.3-alpine3.10

WORKDIR /app

COPY package.json ./
COPY yarn.lock ./
COPY tsconfig.json ./

RUN yarn

CMD [ "yarn", "dev" ]
