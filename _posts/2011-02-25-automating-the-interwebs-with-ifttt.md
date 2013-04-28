---
layout: post
title: Automating the Interwebs with IFTTT
category: internet
tags: [internet]
summary: February 17th I was enrolled in the beta preview of IFTTT, short for *If This Then That*. Now that I have been trying it out some I want to share my experiences.
---
February 17th I was enrolled in the beta preview of [IFTTT](http://ifttt.com/), short for *If This Then That*. Now that I have been trying it out some I want to share my experiences.

Basically IFTTT is a flexible event based system with a slick and stylishly oversized web interface on top of it. You can create up to ten *tasks*, which in turn consists of one *trigger channel* and one *action channel*.

Your triggers looks for certain conditions, such as new items in a RSS feed, being tagged on Facebook or even temperature changes. When such an event occurs, your trigger executes its associated action. This can invoke things such as sending an email, saving URLs to [Pinboard](http://pinboard.in/) or uploading an image to [Flickr](http://www.flickr.com/).

There are currently 204 possible trigger/action combinations and your imagination is really the only limit to their implementation. I have been meaning to write a Twitter integration for my blog, so I can easily tweet about new posts when I publish them. Instead of spending one/two hours programming I managed this with a couple of clicks in IFTTT.

You can probably tell I am very satisfied with this service. Then there is also the icing on this cake – their support. When I wanted to implement my tweet-on-new-post feature I also wanted to throw in tracking parameters for [Google Analytics](http://www.google.com/analytics/). That leads to [humongously](http://www.worldwidewords.org/weirdwords/ww-hum2.htm) long URLs and this is incompatible with Twitter's 140 character limit. So they need to be [shortened](http://en.wikipedia.org/wiki/URL_shortening) before tweeting.

I sent an email to their support describing my use case and immediately got a response saying they would look into it. They got it slightly wrong first, with URL encoded ampersands in all URLs. Then in take two, just a few days after my initial mail, they not only fixed it but also implemented personal bit.ly shortening. That way your shortened links will be added to your own list of URLs and "owned" by yourself. They gave it that extra attention to go above and beyond!

Have you tried it out yet? Do let me know what implementations you have been (or would be) setting up.
