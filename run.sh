#!/bin/bash
set -e

timestamp() {
  date +"%T"
}

cleanup() {
    echo "$(timestamp) [RUN] Docker clean up started."
    set +e
    docker stop $(docker ps -f name=nginx-test -q) 1>/dev/null 2>&1
    docker rm $(docker ps -a -f name=nginx-test -q) 1>/dev/null 2>&1
    docker stop $(docker ps -f name=opa-test -q) 1>/dev/null 2>&1
    docker rm $(docker ps -a -f name=opa-test -q) 1>/dev/null 2>&1
    docker network rm opaauthznginx 1>/dev/null 2>&1
    set -e
    echo "$(timestamp) [RUN] Docker clean up finished."
}

cleanup

echo "$(timestamp) [RUN] Creating docker network"
docker network create --driver bridge opaauthznginx 1>/dev/null
echo "$(timestamp) [RUN] Docker network created"

echo "$(timestamp) [RUN] Starting opa"
docker run --network opaauthznginx -p 8181:8181 -v "$(pwd)"/opa:/mnt:ro --name opa-test -d openpolicyagent/opa:0.29.4-debug run --log-level info --server --addr :8181 /mnt 1>/dev/null 2>&1
echo "$(timestamp) [RUN] Started opa"

echo "$(timestamp) [RUN] Starting nginx"
docker run --network opaauthznginx -p 8080:8080 -v "$(pwd)"/nginx/nginx.conf:/etc/nginx/nginx.conf:ro --name nginx-test -d nginx:1.20.1 1>/dev/null 2>&1
echo "$(timestamp) [RUN] Started nginx"

echo "$(timestamp) [RUN] Press enter to stop."
read

cleanup