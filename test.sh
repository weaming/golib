#!/bin/bash
# File Name : test.sh
# Author    : weaming
# Mail      : garden.yuen@gmail.com
# Created   : 2018-04-29 16:23:02
set -e

for d in $(find . -type d  | grep -v '^\./\.git' | grep -v "^.$"); do
    # ignore blank dir
    if [ "$(ls -A $d)" ]; then
        echo "==== test $d ===="
    else
        continue
    fi

    pushd "$d" > /dev/null
    go test 1>/dev/null
    popd > /dev/null
done
