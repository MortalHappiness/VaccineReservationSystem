FROM node:16-buster

WORKDIR /app

COPY ["package.json", "yarn.lock", "./"]

RUN yarn

COPY . .

EXPOSE 3000

CMD ["yarn", "dev", "--host"]

