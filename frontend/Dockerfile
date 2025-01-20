FROM node:lts-slim

WORKDIR /app

COPY package.json ./
COPY package-lock.json ./

RUN npm install

EXPOSE 3000

COPY . .

CMD ["npm", "run", "dev"]