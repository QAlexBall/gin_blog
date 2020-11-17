#!/bin/bash
if [ `lsof -i:80 | wc -l` -gt "0" ];then
    pg_ctl -D /usr/local/var/postgres start
else
    echo "posgresql is running..."
fi

air -c .air.conf
