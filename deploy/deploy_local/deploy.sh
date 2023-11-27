#!/bin/bash

# 检查权限
echo "please user root or sudo"

WORK_DOCKER_PATH="/opt/block/monitor"
NOW_WORK_PATH=`pwd`
PRE_ROOT_PATH="${NOW_WORK_PATH}/"
NODE_EXPORTER_PATH="${PRE_ROOT_PATH}/node_exporter/"
DOCKER_PATH="${PRE_ROOT_PATH}/docker/"

# install docker docker-compose
docker_cnt=`which docker|wc -l`
if [ ${docker_cnt} -le 0 ];then
   apt install docker.io  docker-compose -y
   #下载docker-compose文件
   # curl -L "https://get.daocloud.io/docker/compose/releases/download/1.29.2/docker-compose-$(uname -s)-$(uname -m)" -o /usr/local/bin/docker-compose
   # chmod +x /usr/local/bin/docker-compose
fi

# install node_exporter binary
cp -f "${NODE_EXPORTER_PATH}/node_exporter" /usr/local/bin/
cp -f "${NODE_EXPORTER_PATH}/node_exporter.service" /etc/systemd/system/
systemctl daemon-reload
systemctl start node_exporter.service
systemctl enable node_exporter.service

ret=`ps -elf|grep node_exporter|grep -v grep|wc -l`
if [ ${ret} -lt 1 ];then
   echo "please reinstall node_exporter"
   return -1
fi

# install grafana prometheus mysql
rm -rf ${WORK_DOCKER_PATH}
# init
mkdir -pv ${WORK_DOCKER_PATH}
cp -rf grafana_data ${WORK_DOCKER_PATH}
mkdir -pv ${WORK_DOCKER_PATH}/prometheus_data
cd ${WORK_DOCKER_PATH}
chmod 777 grafana_data/grafana.db grafana_data prometheus_data
cp -f ${DOCKER_PATH}/docker-compose-monitor.yaml . 
cp -f ${DOCKER_PATH}/prometheus.yml . 
docker-compose -f docker-compose-monitor.yaml down
yes 'y'|docker volume prune
docker-compose -f docker-compose-monitor.yaml up -d
docker update --restart=always $(docker ps -aq)


# install ansible
# apt install -y ansible
