#!/bin/bash

dd if=/dev/urandom of=/tmp/file bs=1000 count=100

base = $(base64  <<< openssl dgst -sha256 -binary <<< ($cat /tmp/file))

echo $base
curl -v 192.168.246.131:12347/objects/test6 -XPOST -H "Digest: SHA-256=${hash}" -H "Size: 100000"