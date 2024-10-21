#!/bin/bash
set -x 

#variables
URL_rpm=https://isrcportal.magna.global/repo/CrowdStrike/Linux/latest/falcon-sensor-7.18.0-17106.el6.x86_64.rpm
URL_deb=https://isrcportal.magna.global/repo/CrowdStrike/Linux/latest/falcon-sensor_7.18.0-17106_amd64.deb
CID=$1

cd /tmp
os-release=awk -F= '/^ID=/ {gsub(/"/, "", $2); print $2}' /etc/os-release
if [ "$os-release" == "redhat" -o "$os-release" == "centos"]; then
wget "$URL_rpm" --no-check-certificate

chmod 777 falcon-sensor-7.11.0-16405.el7.x86_64.rpm
yum localinstall falcon-sensor-7.11.0-16405.el7.x86_64.rpm -y
sudo /opt/CrowdStrike/falconctl -s --cid="$CID"

systemctl enable falcon-sensor
systemctl start falcon-sensor
systemctl status falcon-sensor
else 

cd /tmp
wget "$URL_rpm" --no-check-certificate
chmod 777 falcon-sensor_7.11.0-16405_amd64.deb
dpkg -i falcon-sensor_7.11.0-16405_amd64.deb
sudo /opt/CrowdStrike/falconctl -s --cid="$CID"

systemctl enable falcon-sensor
systemctl start falcon-sensor
systemctl status falcon-sensor
fi
