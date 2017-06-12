#!/bin/bash

SERVER_SOURCE=./backend/
WORK_DIR=/www/gyg
EXPOSED_PORT=8000
IDE_KEY="PHPSTORM"
XDEBUG_PORT=9999
CONFIG_FILE_PATH="config/main.json"

while getopts "s:w:p:i:x:" name
do
    case $name in
    s)  SERVER_SOURCE="$OPTARG";;
    w)  WORK_DIR="$OPTARG";;
    p)  EXPOSED_PORT="$OPTARG";;
    i)  IDE_KEY="$OPTARG";;
    x)  XDEBUG_PORT="$OPTARG";;
    ?)  printf "Usage: %s: [-s server source dir]\n" $0
          exit 2;;
    esac
done

export SERVER_SOURCE="$SERVER_SOURCE"
export WORK_DIR="$WORK_DIR"
export EXPOSED_PORT="$EXPOSED_PORT"
export IDE_KEY="$IDE_KEY"
export XDEBUG_PORT="$XDEBUG_PORT"
export CONFIG_FILE_PATH="$CONFIG_FILE_PATH"

echo "Server source dir: $SERVER_SOURCE"
echo "Server working dir: $WORK_DIR"
echo "Server exposed port: $EXPOSED_PORT"
echo "Server xdebug ide key: $IDE_KEY"
echo "Server xdebug port: $XDEBUG_PORT"

export IGNORE_FILES="
ignore = Path var/log
"

# Gets the current platform.
getPlatform() {
  echo $(uname -s)
}

# Get host ip
getHostIP() {
    echo /sbin/ip route|awk '/default/ { print $3 }'
}

# Should an ip alias used based on the platform.
shouldUseIpAlias() {
  local platform=$1

  if [[ $platform == 'linux' ]]; then
     return 1
  elif [[ $platform == 'Darwin' ]]; then
     return 0
  fi

  return 1
}

export HOST_OS=$(getPlatform)
HOST_IP=$(getHostIP)

if ( shouldUseIpAlias $HOST_OS )
then
  HOST_IP=10.200.10.1
  echo "Aliasing host ip: $HOST_IP"
  sudo ifconfig lo0 alias $HOST_IP/24
fi

export HOST_IP="$HOST_IP"

docker-compose --verbose build --pull
docker-compose --verbose up -d

cd backend/src/ && php -s localhost:8080
