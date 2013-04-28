---
layout: post
title: RSS/Atom feed to Facebook page and Twitter
category: internet
tags: [internet]
summary: Here is how I automatically post content from my Atom feed, to my Facebook page and to Twitter.
---
I am trying to gain some traction for [a vegetarian workout blog](http://vvv.vegout.se/) of mine and part of my efforts goes into social networks. Here is how I *automatically post content* from my Atom feed, to my *Facebook page* and to *Twitter*.

This assumes your website already has a [RSS feed](/rss/).

## Posting to Twitter

Automatically posting to Twitter is dead simple with the awesome service [IFTTT](https://ifttt.com/). It lets you set up triggers for various conditions and which fires actions.

In my case the trigger is "*New feed item*", where I enter the feed URL, and the action is "*Post a tweet*". Once saved, IFTTT will check the feed every 15 minutes and send a tweet as soon as it sees a new entry.

If you have multiple Twitter accounts you might have to create an IFTTT account for each of them, because IFTTT can only handle a one-to-one relation. Sucks, but you probably will only set this up once and then be done with it.

## Posting to Facebook

You could use IFTTT for Facebook as well, but they do not support posting to Facebook pages. Unfortunately, this is a requirement for my use case.

What you often see when searching for a solution to automatically posting to Facebook pages is [RSS Grafitti](http://www.rssgraffiti.com/). It is definitely a nice service but I really do not like the way its shared posts are displayed.

Instead I would recommend [dlvr.it](http://dlvr.it/). This little gem detects new posts from a given feed and automatically shares them to any Facebook page, app or wall you want.

And it does so exactly the way I want them to be displayed; which is as if I would have posted them mnully myself. This respects my `og` tags as well, so the correct image, title and description text is used.

## Automatic posting

Implementing automatic posting to Twitter and Facebook is easy. But when should you use it?

For [tobiassjosten.net](http://vvv.tobiassjosten.net/) I do not use automatic posting, because I like being in control of what titles I use for the various channels. A post like "[PHP Mentoring](/php/php-mentoring/)" will be "#PHP Mentoring" in Twitter, but "+PHP Mentoring" on Google+.

This blog also caters to a more tech savvy audience, whom can subscribe to a feed with their syndication tool of choice. Hence I do not feel I would do readers a big favor by having my tech posts show up in their Facebook feed.

This is different for [Vegout.se](http://vvv.vegout.se/), so more channels is warrented and so automatic posting makes sense. Or else posting manually could easily become cumbersome.
