---
layout: post
title: Correct IP with reverse proxy and Drupal
category: drupal
tags: [drupal, -x-drupal-planet]
summary: If you are running a reverse proxy, like Varnish or Squid, in front of your webserver, then it will report the proxy's IP address instead of the visitors'. This is a problem if you, for example, allow anonymous users to vote with VotingAPI.
---
If you are running a reverse proxy, like [Varnish](http://www.varnish-cache.org/) or [Squid](http://www.squid-cache.org/), in front of your webserver, then it will report the proxy's IP address instead of the visitors'. This is a problem if you, for example, allow anonymous users to vote with [VotingAPI](http://drupal.org/project/votingapi).

It can be solved by instead checking the [X-Forwarded-For](http://en.wikipedia.org/wiki/X-Forwarded-For) header. Luckily, in [Drupal](http://drupal.org/), this is very easy to do.

You configure the *settings.php* file, which is most likely located in the *sites/default* directory. It already contains all the documentation you need but I prefer to keep my *settings.php* lean and strip that out. Instead I add the following lines to the bottom of my file:

    $conf['reverse_proxy'] = TRUE;
    $conf['reverse_proxy_addresses'] = array('1.2.3.4');

If you have more than one webfront, simply add its IP to the array. This tells Drupal which reverse proxies you trust and helps protect against users spoofing their IP with a fake *X-Forwarded-For* header.
