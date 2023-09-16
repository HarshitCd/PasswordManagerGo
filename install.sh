#! /bin/bash

mkdir -p ~/.config/passswdmanagergo
touch ~/.config/passswdmanagergo/password_details.yml
mkdir -p ~/mybin/passswdmanagergo/bin

cd cli
go build -o ~/mybin/passswdmanagergo/bin/prun
echo "export PATH=\"/Users/harshitcd/mybin/passswdmanagergo/bin/:${PATH}\"" >> ~/.zprofile
export PATH="/Users/harshitcd/mybin/passswdmanagergo/bin/:${PATH}"


