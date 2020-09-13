#!/bin/bash

/etc/init.d/ssh start
/etc/init.d/rsyslog start
while true; do sleep 30; done;