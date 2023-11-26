#!/bin/bash

RUN_PATH=`pwd`
NODE_EXPORTER_PATH="${RUN_PATH}/"

# install node_exporter binary
cp -f "${NODE_EXPORTER_PATH}/node_exporter" /usr/local/bin/
cp -f "${NODE_EXPORTER_PATH}/node_exporter.service" /etc/systemd/system/
systemctl daemon-reload
systemctl restart node_exporter.service
systemctl enable node_exporter.service

ret=`ps -elf|grep node_exporter|grep -v grep|wc -l`
if [ ${ret} -lt 1 ];then
    echo "please reinstall node_exporter"
    return -1
fi
