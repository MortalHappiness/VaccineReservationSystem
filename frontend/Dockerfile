FROM node:16-buster

WORKDIR /app

COPY ["package.json", "yarn.lock", "./"]

RUN yarn

RUN npm install serve -g

COPY . .
RUN yarn build

CMD ["serve", "-s", "dist", "-l", "3000"]

