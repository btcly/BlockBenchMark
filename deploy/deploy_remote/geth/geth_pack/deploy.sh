#!/bin/bash

RUN_PATH=`pwd`
NODE_EXPORTER_PATH="${RUN_PATH}/node_exporter/"

# 生成geth账户
# for ((i = 0; i < 50; i++))
# do
# 	${RUN_PATH}/geth --datadir nodedata account new --password=passwd
# done


${RUN_PATH}/geth --datadir ${RUN_PATH}/nodedata --networkid "123456" init "${RUN_PATH}/geth.json"

# ./geth --datadir nodedata account new #passwd=123456
echo "123456" > passwd
echo '#!/bin/bash' > attach.sh
echo "${RUN_PATH}/geth attach ipc:${RUN_PATH}/nodedata/geth.ipc" >> attach.sh


# install Geth
minerbase=`cat geth.miner`
geth_exe_cmd="${RUN_PATH}/geth  --nodiscover --ethash.cachedir '${RUN_PATH}/ethashdata' --ethash.dagdir '${RUN_PATH}/ethashdata' --allow-insecure-unlock --unlock=${minerbase}  --password  ${RUN_PATH}/passwd  --networkid  '123456'  --datadir  '${RUN_PATH}/nodedata'  --http --http.api 'admin,debug,web3,eth,txpool,personal,ethash,miner,net' --http.corsdomain='*' --http.port=8545 --http.addr='0.0.0.0'  --ws --ws.addr '0.0.0.0' --ws.port=8546 --ws.origins '*' --ws.api 'admin,debug,web3,eth,txpool,personal,ethash,miner,net' --syncmode full --nodiscover --nat=extip:IPADDR_REPLACE --mine --miner.threads=2 --miner.etherbase=${minerbase}"

echo '#!/bin/bash' > run.sh
echo ${geth_exe_cmd} >> run.sh
cp -f "${RUN_PATH}/gethtmp.service" geth.service
sed -i "s#WORK#${RUN_PATH}#g" geth.service
sed -i "s#EXEC#${geth_exe_cmd}#g" geth.service
mv geth.service /etc/systemd/system/
systemctl daemon-reload
systemctl restart geth.service
systemctl enable geth.service

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
