#!/bin/bash
RUN_PATH=`pwd`
TARGET_SERVICE="meepo.service"
NODE_EXPORTER_PATH="${RUN_PATH}/node_exporter/"

# compression
# tar -czvf - openethereum | split -b 50m - openethereum.tar.gz.part.
# decompression
# cat openethereum.tar.gz.part* | tar -xzvf -

# install Meepo
meepo_exe_cmd="${RUN_PATH}/openethereum --config ${RUN_PATH}/node.toml --base-path ${RUN_PATH}/nodedata --password ${RUN_PATH}/node.pwds"
echo '#!/bin/bash' > run.sh
echo ${meepo_exe_cmd} >> run.sh
cp -f "${RUN_PATH}/meepotmp.service" ${TARGET_SERVICE}
sed -i "s#WORK#${RUN_PATH}#g" ${TARGET_SERVICE}
sed -i "s#EXEC#${meepo_exe_cmd}#g" ${TARGET_SERVICE}
mv ${TARGET_SERVICE} /etc/systemd/system/
systemctl daemon-reload
systemctl restart ${TARGET_SERVICE}
systemctl enable ${TARGET_SERVICE}

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
