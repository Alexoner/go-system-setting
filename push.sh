#########################################################################
# File Name: push.sh
# Author: hao
# mail: onerhao@gmail.com
# Created Time: Fri 08 Nov 2013 03:18:08 PM CST
#########################################################################
#!/bin/sh

git commit -am $@
git push

cp ./power/{org.deepin.daemon.power.service,power.go} ~/Documents/deepin/src/dde-daemon/power/
cp ./sound/{org.deepin.daemon.sound.service,pulse.go} ~/Documents/deepin/src/dde-daemon/sound/
cp ./display/{org.deepin.daemon.display.service,display.go} ~/Documents/deepin/src/dde-daemon/display/

cd ~/Documents/deepin/src/dde-daemon/
git commit -am $@
git push
