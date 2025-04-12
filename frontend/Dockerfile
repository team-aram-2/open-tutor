FROM node:23 AS base

WORKDIR /app
COPY . /app

RUN npm ci
RUN npm run build
CMD ["npm", "run", "preview"]