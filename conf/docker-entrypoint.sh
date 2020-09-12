#!/bin/bash
chown -R admin:admin /data/public
chmod -R 775 /data/public
chmod 755 /data/users
/etc/init.d/ssh start
/etc/init.d/rsyslog start
while true; do sleep 30; done;