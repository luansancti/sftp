#!/bin/bash
chown -R admin:admin /data/public
chmod -R 775 /data/public
/etc/init.d/ssh start
while true; do sleep 30; done;