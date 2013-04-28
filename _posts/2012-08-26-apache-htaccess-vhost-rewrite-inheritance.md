---
layout: post
title: Apache .htaccess and vhost rewrite inheritance
category: drupal
tags: [drupal, -x-drupal-planet]
summary: Currently when you want to affect the way requests are routed into Drupal, you have two choices; either you hack the core `.htaccess` or you move all its content into your virtual host. We all know hacking core is bad but code duplication is also irky and we do not want to miss future updates to `.htaccess` when upgrading.
---
Currently when you want to affect the way requests are routed into [Drupal](/drupal), you have two choices; either you hack the core `.htaccess` or you move all its content into your virtual host. We all know [hacking core is bad](http://www.flickr.com/photos/hagengraf/2802915470/) but code duplication is also irky and we do not want to miss future updates to `.htaccess` when upgrading.

This dilemma is not an edge case. Just installing [robotstxt.module](http://drupal.org/project/robotstxt) requires this kind of functionality, not to mention if you have moved from another system to Drupal and want to handle old URLs.

So it looks like a choice between plague and cholera. But actually, there is quite an elegant solution to our problem — *rewrite inheritance*.

## Rewrite inheritance

Using rewrite inheritance is not a new concept to Drupal; it is even recommended in the [Multisite Install & Configuration guide](http://drupal.org/node/138889). However, the way it is applied only lets the vhost inherit configuration from the main server. Our `.htaccess` will still come in and ruin everything.

Let us have a look at how it currently works.

    $ grep Rewrite /etc/apache2/sites-available/rewritetest
    RewriteEngine on
    RewriteOptions inherit
    RewriteRule ^vhost /works [R=301,L]

    $ grep Rewrite .htaccess
    RewriteEngine on
    RewriteRule ^htaccess /works [R=301,L]

    $ curl -I http://rewritetest/vhost
    HTTP/1.1 404 Not Found
    Server: Apache/2.2.22 (Ubuntu)

    $ curl -I http://rewritetest/htaccess
    HTTP/1.1 301 Moved Permanently
    Server: Apache/2.2.22 (Ubuntu)
    Location: http://rewritetest/works

The rewrite rules in our vhost is completely disregarded while the `.htaccess` works just fine. So the recommended [`RewriteOptions`](http://httpd.apache.org/docs/current/mod/mod_rewrite.html#rewriteoptions) configuration does nothing.

Now let us see how to fix this, by adding the `RewriteOptions` configuration to the `.htaccess` to have it inherit our rewrite rules.

    $ grep Rewrite .htaccess
    RewriteEngine on
    RewriteOptions inherit # Play nice, .htaccess!
    RewriteRule ^htaccess /works [R=301,L]

    $ curl -I http://rewritetest/htaccess
    HTTP/1.1 301 Moved Permanently
    Server: Apache/2.2.22 (Ubuntu)
    Location: http://rewritetest/works

    $ curl -I http://rewritetest/vhost
    HTTP/1.1 301 Moved Permanently
    Server: Apache/2.2.22 (Ubuntu)
    Location: http://rewritetest/works

Bam! Both configurations now work perfectly, side by side. This means we can keep the `.htaccess` pristine while being able to add our custom rewrite rules.

## Where is the patch?

So, when can we expect this in Drupal proper? As soon as I am able to convince [chx](https://twitter.com/chx) that *this is an excellent idea*. Feel free to [pitch in](http://drupal.org/node/1707998) if you agree!
