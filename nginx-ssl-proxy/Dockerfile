FROM       nginx
MAINTAINER Jetstack <docker@jetstack.io>
RUN        rm /etc/nginx/conf.d/*.conf
WORKDIR    /usr/src
ADD        start.sh             /usr/src/
ADD        nginx/nginx.conf     /etc/nginx/
ADD        nginx/proxy-ssl.conf /etc/nginx/conf.d/
ENTRYPOINT ["./start.sh"]
