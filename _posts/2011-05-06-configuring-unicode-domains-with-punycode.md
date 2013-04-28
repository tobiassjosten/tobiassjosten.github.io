---
layout: post
title: Configuring Unicode domains with Punycode
category: internet
summary: If you have a domain with Unicode characters, like tobiassjösten.se, there are some gotchas to consider.
---
If you have a domain with Unicode characters, like [tobiassjösten.se](http://tobiassjösten.se/), there are some gotchas to consider.

First you should know that these domains are not actually communicated in Unicode, but converted to [Punycode](http://en.wikipedia.org/wiki/Punycode) before requests are sent.

In my case this converts to xn--tobiassjsten-cjb.se. The *xn--* part is a prefix to let the server know we are using *Punycode*. The rest, tobiassjsten-cjb.se, is the actual transcoded domain name.

Secondly there is a difference in how servers handles these domains. With [Nginx](http://nginx.org/) you can simply type the Unicoded domain into the configuration and everything will work as expected. For [Apache](http://www.apache.org/), however, you need to use Punycode in the ServerName and ServerAlias configurations.

I use [the Verisign converter](http://mct.verisign-grs.com/) for this task.
