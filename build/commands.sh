#!/bin/sh

sleep 15 # wait for postgres to be ready for connections

/usr/bin/migrations init
/usr/bin/migrations up
/usr/bin/liproduce