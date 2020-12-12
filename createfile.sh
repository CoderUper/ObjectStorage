#!/bin/bash

for i in `seq 1 6`
do
    mkdir -p /tmp/$i/objects
done

sudo ifconfig ens33:1 10.29.1.1/16
sudo ifconfig ens33:2 10.29.1.2/16
sudo ifconfig ens33:3 10.29.1.3/16
sudo ifconfig ens33:4 10.29.1.4/16
sudo ifconfig ens33:5 10.29.1.5/16
sudo ifconfig ens33:6 10.29.1.6/16
sudo ifconfig ens33:7 10.29.2.1/16
sudo ifconfig ens33:8 10.29.2.2/16