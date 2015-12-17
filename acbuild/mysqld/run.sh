#!/usr/bin/env bash

set -e

grant () {
	sleep 2
	IP=`ifconfig eth0 | grep inet[[:space:]] | cut -d: -f2 | awk '{print $1}' | cut -d"." -f1,2,3`
	echo "GRANT ALL PRIVILEGES ON *.* TO 'root'@'$IP.%'"
	echo "GRANT ALL PRIVILEGES ON *.* TO 'root'@'$IP.%'" | mysql -u root
}

grant &

/usr/bin/mysql_install_db --user=mysql

/usr/bin/mysqld_safe --datadir='/var/lib/mysql'
