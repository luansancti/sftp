#!/bin/bash
chown -R admin:admin /data/public
chmod -R 775 /data/public
service ssh restart
while true; do sleep 30; done;