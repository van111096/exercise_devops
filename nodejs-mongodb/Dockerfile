FROM node:12-alpine
WORKDIR /app
COPY package.json .
RUN npm install
COPY ./src .
CMD ["node", "app.js"]