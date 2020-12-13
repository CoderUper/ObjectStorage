#!/bin/bash


#curl "localhost:9200/metadata" -XDELETE

curl -H "Content-Type: application/json"  "localhost:9200/metadata"  -XPUT -d'{"mappings":{"objects":{"properties":{"name":{"type":"text","index":"true","fielddata": "true"},"version":{"type":"integer"},"size":{"type":"integer"},"hash":{"type":"text"}}}}}'

