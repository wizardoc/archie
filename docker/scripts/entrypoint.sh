#!/bin/sh
ENTRY_POINT="../../build/archie"

/usr/local/bin/wait-for-it.sh postgres:5432 -- chmod 777 $ENTRY_POINT && $ENTRY_POINT