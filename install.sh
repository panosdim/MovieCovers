#!/bin/bash

if [ "$(id -u)" != "0" ]; then
    echo "You must be root or have root priviledges in order to install this software."
    exit 1
fi

INSTALL_PATH=/opt/autotag

mkdir -p $INSTALL_PATH
cp ths $INSTALL_PATH
cp app.properties $INSTALL_PATH

read -p "Do you wish to install a systemd service that will run autotag every day at 03:00?" yn
case $yn in
    [Yy]* ) install_service; break;;
    [Nn]* ) exit;;
    * ) echo "Please answer yes or no.";;
esac

function install_service {
    read -p "Provide the path to movies files that AutoTag will scan?" path
    cp autotag.service tmp.service
    sed -i "s/<PATH>/-p $path/g" tmp.service
    read -p "Do you wish to provide a regex to exclude from scan?" yn
    case $yn in
        [Yy]* ) 
            read -p "Provide exclude regex?" regex; 
            sed -i "s/<REGEX>/-e $regex/g" tmp.service;;
            break;;
        [Nn]* ) sed -i 's/<REGEX>//g' tmp.service;;
        * ) echo "Please answer yes or no.";;
    esac
    mv tmp.service /usr/lib/systemd/system/autotag.service
    cp autotag.timer /usr/lib/systemd/system

    systemctl start autotag.service
    systemctl enable autotag.service
}

