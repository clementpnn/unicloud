FROM node:20-alpine as builder

RUN npm install -g pnpm

WORKDIR /app

COPY ./apps/frontend/pnpm-lock.yaml ./apps/frontend/package.json ./

RUN pnpm install

COPY ./apps/frontend ./

RUN pnpm build

FROM nginx:alpine

COPY --from=builder /app/dist /usr/share/nginx/html

EXPOSE 8080

CMD ["nginx", "-g", "daemon off;"]