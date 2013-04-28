---
layout: post
title: Monitor and debug with Charles proxy
category: development
tags: [development]
summary: A while ago I was introduced to Charles - a web debugging proxy application. It acts as a middle man between your browser and your web application, where it can do a multitude of helpful services.
---
A while ago I was introduced to [Charles](http://www.charlesproxy.com/) - *a web debugging proxy application*. It acts as a middle man between your browser and your web application, where it can do a multitude of helpful services.

You can use it for debugging, where it shows you exactly what traffic is sent and received. I used to use [Live HTTP Headers](https://addons.mozilla.org/sv-se/firefox/addon/live-http-headers/) for Firefox before but with Charles you are effectively browser agnostic.

One other use case is having Charles rewrite URLs. This is especially useful when you are working with [CDNs](http://en.wikipedia.org/wiki/Content_delivery_network) and want to serve local variations of files instead of the ones on the network.

The big drawback, in my opinion, is that Charles is not open source. I am very much surprised there are no competing [FOSS software](http://www.gnu.org/philosophy/free-sw.html) for this niche. But maybe I have just not found it? Let me know if you have any other recommendation!
