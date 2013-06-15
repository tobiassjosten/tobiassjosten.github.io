---
layout: post
title: Automatic MySQL backups
category: operations
tags: [operations]
summary: Automate rolling MySQL backups to keep the data in you production databases safe and sound.
---
Quick notes how to *automate rolling MySQL backups* for production databases. I keep having to look this up but now I hope I can just copy and paste it for future projects.

First set up the credentials in a file with secure permissions.

    [client]
    user="myuser"
    password="mysecret"

Then create two cron jobs to backup once per day and once per month.

    # /etc/cron.d/mysql-backup
    59 23 * * *     root    mysqldump --defaults-extra-file=mysql.cfn dbname > "/var/backups/`date +\%A|tr '[:upper:]' '[:lower:]'`.sql"
    59 23 28 * *    root    mysqldump --defaults-extra-file=/var/mysql.cfn dbname > "/var/backups/`date +\%B|tr '[:upper:]' '[:lower:]'`.sql"

This will create files like `saturday.sql`, `february.sql`, etc.

For extra credit you could also set up a cron job to send these backups to an offsite storage, like Amazon S3 or just another VPS provider.
