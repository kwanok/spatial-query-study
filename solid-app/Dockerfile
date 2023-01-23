FROM nginx:1.23.3-alpine

RUN mkdir /app
WORKDIR /app

RUN mkdir ./dist

COPY ./dist /usr/share/nginx/html

EXPOSE 80

CMD ["nginx", "-g", "daemon off;"]