#!/usr/bin/env bash

# syslog
rsyslogd

which php

if [[ $HOST_OS == 'Linux' ]]; then
   HOST_IP=`/sbin/ip route|awk '/default/ { print $3 }'`
fi

while [ ! -f /www/gyg/src/test_full_stack.php ]
do
  echo "Waiting for gyg files to be synced."
  sleep 1
done

# end up in shell
tail -f /dev/null
