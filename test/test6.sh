#!/bin/bash

dd if=/dev/urandom of=/tmp/file bs=1000 count=100

openssl dgst -sha256 -binary /tmp/file | base64

dd if=/tmp/file of=/tmp/first bs=1000 count=50

dd if=/tmp/file of=/tmp/second bs=1000 skip=32 count=68




curl -v 192.168.246.131:12347/objects/test6 -XPOST -H "Digest: SHA-256=$1" -H "Size: 100000"


curl -I 192.168.246.131:12347/$1

curl -v -XPUT --data-binary @/tmp/first 192.168.246.131:12347/$1

curl -I 192.168.246.131:12347/$1

curl -v -XPUT --data-binary @/tmp/second -H "range: bytes=32000-" 192.168.246.131:12347/$1

curl -I 192.168.246.131:12347/$1

curl -v -XPUT --data-binary @/tmp/last -H "range: bytes=96000-" 192.168.246.131:12347/$1

curl 192.168.246.131:12347/objects/test6 > /tmp/output

diff -s /tmp/output /tmp/file

curl 192.168.246.131:12347/objects/test6 -H "range: bytes=32000-" > /tmp/output2

diff -s /tmp/output2 /tmp/second