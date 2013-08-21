---
layout: post
title: Installing Redis on Ubuntu with APT
category: linux
tags: [linux, redis, ubuntu]
summary: Deciding to use APT to maintain your software is really a no-brainer if you are Debian based. I prefer Ubuntu myself and its repositories has a lot of toys in them but for stability reasons they often do not hold the very latest versions.
---
Deciding to use [APT](http://en.wikipedia.org/wiki/Advanced_Packaging_Tool) to maintain your software is really a no-brainer if you are Debian based. I prefer *Ubuntu* myself and its repositories has a lot of toys in them but for stability reasons they often do not hold the very latest versions.

Earlier today I was looking for an alternative repository for [Redis](http://redis.io/) and luckily I ran into [Dotdeb](http://www.dotdeb.org/). They have a bunch of nice packages, always up to date with the *very latest version*.

## Installation

In order to get in on this goodness you must first add the _Dotdeb_ repositories to your APT sources. Create a new list file in `/etc/apt/sources.list.d/` and fill it with the following content.

    # /etc/apt/sources.list.d/dotdeb.org.list
    deb http://packages.dotdeb.org squeeze all
    deb-src http://packages.dotdeb.org squeeze all

Then you need to authenticate these repositories using their public key.

    wget -q -O - http://www.dotdeb.org/dotdeb.gpg | sudo apt-key add -

And finally, update your APT cache and *install Redis*.

    $ sudo apt-get update
    $ sudo apt-get install redis-server

Happy key-value-storing!
