#!/bin/bash
systemctl stop katy.service 
systemctl disable katy.service 
rm /lib/systemd/system/katy.service
rm /etc/default/katy
rm /usr/local/sbin/katy
rm -rf /var/lib/katy
