---
layout: post
title: Facebook canvas with HTTP problems
category: internet
tags: [internet, facebook]
summary: A few weeks ago I built a simple Facebook app that, among other things, would provide a canvas page. It worked perfectly when prototyping but a few coding hours later it suddenly stopped working.
---
A few weeks ago I built a simple Facebook app that, among other things, would provide a canvas page. It worked perfectly when prototyping but a few coding hours later it suddenly stopped working.

The canvas page was served just fine when I visited it directly but it blanked out in Facebook. Whatever different approach I tried I could not really find anything wrong with my implementation. So I started looking at the Facebook end if it.

What finally cracked the nut was Chrome's Developer tool and its Network tab specifically. It showed me Facebook sent a POST request to the canvas page. [Symfony](http://symfony.com/), on which I built the app, has excellent [HTTP](http://en.wikipedia.org/wiki/Representational_State_Transfer) routing and I was using that to split requests up between different controllers, depending on [HTTP method](http://www.w3.org/Protocols/rfc2616/rfc2616-sec9.html). GET requests got to see the page while POST requests was assumed to be sending in a form.

This is a violation of the very basis of the HTTP protocol. So why are Facebook doing it? My first guess was that they wanted to prevent AJAX requests from caching. But why would they not use a randomly generated URL parameter for this?

I am holding the Facebook web developers to a much higher standard than that. There must be a better reason for them to do this. I just can not think of any. What is your guess?
