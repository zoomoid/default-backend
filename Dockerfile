FROM nginx:alpine

COPY entrypoint.sh /docker-entrypoint.d/template-index-html.sh
COPY src/ /usr/share/nginx/html/

COPY nginx.conf /etc/nginx/templates/backend.conf.template
