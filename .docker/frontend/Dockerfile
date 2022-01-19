FROM node:15.8-alpine3.11 as builder

ENV APPDIR /app

RUN mkdir -p ${APPDIR}
WORKDIR ${APPDIR}

COPY ./pkg/http/web/app/package.json .

RUN yarn install

COPY ./pkg/http/web/app/ .

RUN yarn build

FROM nginx:1.19-alpine
COPY --from=builder /app/dist /app
COPY ./.docker/frontend/nginx.conf /etc/nginx/nginx.conf

ENV API_HOST=http://localhost:8000

EXPOSE 80
