0 0 * * *	root	mysqldump --defaults-extra-file=/var/www/smartburk/shared/mysql.cfn smartburk > /var/www/smartburk/shared/backup/`date +%A|tr '[:upper:]' '[:lower:]'`.sql

[client]
user="smartburk"
password="secretpassword"

chmod
