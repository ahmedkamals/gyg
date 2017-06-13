#!/usr/bin/env bash

which php

if [[ ${HOST_OS} == 'Linux' ]]; then
   HOST_IP=`/sbin/ip route|awk '/default/ { print $3 }'`
fi

cd /www/src/gyg && php -S localhost:8090 &&

# end up in shell
tail -f /dev/null
