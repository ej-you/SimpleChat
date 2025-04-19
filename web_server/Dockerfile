# --------
# frontend
# --------

FROM node:current-alpine3.21 AS frontend

WORKDIR /app/frontend

# install requirements
COPY ./frontend/package.json .
RUN npm install

# build
COPY ./frontend .
RUN npm run build

# -----
# nginx
# -----

FROM nginx:stable-alpine3.20-perl

WORKDIR /app/frontend

RUN chown nginx:nginx /var/cache/nginx/

# main nginx conf
COPY ./web_server/nginx.conf /etc/nginx/nginx.conf
# included nginx confs
COPY ./web_server/https_site.conf /etc/nginx/conf.d/https_site.conf

# copy built frontend files
COPY --from=frontend /app/frontend/public ./public
COPY --from=frontend /app/frontend/dist ./dist
