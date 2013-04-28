---
layout: post
title: PHP session garbage collection
category: php
tags: [php]
summary: I just bumped into a notice from PHP about a failed session_start(). Here is how I solved it.
---
I just bumped into a rather omnious notice from PHP, about `session_start()` not being able to complete.

    Notice: session_start() [function.session-start]: ps_files_cleanup_dir: opendir(/var/lib/php5) failed: Permission denied (13) in ../vendor/symfony/src/Symfony/Component/HttpFoundation/SessionStorage/NativeSessionStorage.php line 83

After digging around some it turned out this is PHP trying to purge session data files from the file system (`/var/lib/php5` in Ubuntu) without having the access to find what files to delete.

## Fixing the problem

You could change the group owner of the directory to `www-data` and give the group read access with `chmod g+r`. On Ubuntu there is already a fix in place however, in `/etc/cron.d/php5`. This cronjob clears out old sessions so PHP does not have to bother.

So if you are on Ubuntu this misconfiguration is easily remedied by disabling `session.gc_probability` in your `php.ini`.

    session.gc_probability = 0

Done deal.

## PHP probability configuration

There are a couple of parameters in play here. First there is [`session.save_handler`](http://www.php.net/manual/en/session.configuration.php#ini.session.save-handler) which decides how to save session data. It defaults to *files* but can also be used to offload the filesystem using [memcached](http://memcached.org/).

Second there is [`session.save_path`](http://www.php.net/manual/en/session.configuration.php#ini.session.save-path). When saving data to files this controls what directory to write to.

For the actual garbage collection we have [`session.gc_probability`](http://www.php.net/manual/en/session.configuration.php#ini.session.gc-probability) (defaults to 1), [`session.gc_divisor`](http://www.php.net/manual/en/session.configuration.php#ini.session.gc-divisor) (defaults to 100) and [`session.gc_maxtime`](http://www.php.net/manual/en/session.configuration.php#ini.session.gc-maxlifetime) (defaults to 1440). The first two decides the probability for garbage collection to happen automatically on [`session_start()`](http://www.php.net/manual/en/function.session-start.php), using a `gc_probability`/`gc_divisor` chance of proccing.

When garbage collection happens it removes all data older than `gc_maxlife` seconds. This means PHP per default has a 1% chance of remove all sessions older than 24 minutes. An interesting feature but I would rather not add on to the page load times any more than absolutely necessary.
