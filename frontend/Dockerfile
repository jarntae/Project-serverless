FROM node:18 as builder

WORKDIR /app
COPY frontend/package.json frontend/package-lock.json ./

RUN npm install

COPY . .

EXPOSE 5173

CMD ["npm","run","dev","--","--host","0.0.0.0"]