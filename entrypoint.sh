#!/bin/sh
cd /usr/share/nginx/html

sed -i "s/{{CONTACT}}/${CONTACT:=mailto:hello@example.org}/" index.html 
sed -i "s/{{AUTHOR}}/${AUTHOR:=example.org}/" index.html
