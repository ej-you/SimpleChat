FROM nginx:1.25.3-alpine

# -----
# nginx
# -----

RUN chown nginx:nginx /var/cache/nginx/
RUN apk add npm

# main nginx conf
COPY ./web_server/nginx.conf /etc/nginx/nginx.conf
# included nginx confs
COPY ./web_server/http_site.conf /etc/nginx/conf.d/http_site.conf

# --------
# frontend
# --------

#WORKDIR /app/frontend

# install requirements
#COPY ./frontend/package.json .
#RUN npm install

# build
#COPY ./frontend .
#RUN npm run build
