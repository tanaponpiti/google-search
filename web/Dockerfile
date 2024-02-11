FROM node:20 AS base
WORKDIR /app

#Download dependencies stage
FROM base AS dependencies
COPY package.json ./
RUN yarn install

# Build stage
FROM dependencies AS build-stage
COPY . .
RUN yarn build

# Serve stage
FROM --platform=linux/amd64 nginx:stable AS serve-stage
COPY --from=build-stage /app/dist /usr/share/nginx/html
COPY envsubst.sh /docker-entrypoint.d/envsubst.sh
RUN chmod +x /docker-entrypoint.d/envsubst.sh

# Replace default Nginx configuration to handle SPA routing
RUN rm /etc/nginx/conf.d/default.conf
COPY nginx.conf /etc/nginx/conf.d

EXPOSE 80
CMD ["nginx", "-g", "daemon off;"]
