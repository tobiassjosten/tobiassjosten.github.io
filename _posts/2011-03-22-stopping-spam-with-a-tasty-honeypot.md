---
layout: post
title: Stopping spam with a tasty honeypot
category: internet
tags: [internet, spam]
summary: Yesterday I implemented spam protection that would resolve MX records in order to validate the comment. It seems to be working but some spammers still got through. So I needed to up my game and head for round two.
notice: This is part of an *[anti-spam](/spam) series*, including [Resolving MX pointers to fight spam](http://local.tobiassjosten.net:1234/php/resolving-mx-pointers-to-fight-spam) and [Stopping spam with Symfony forms](/symfony/stopping-spam-with-symfony-forms).
---
Yesterday I implemented spam protection that would [resolve MX records](http://vvv.tobiassjosten.net/php/resolving-mx-pointers-to-fight-spam) in order to validate the comment. It seems to be working but some spammers still got through. So I needed to up my game and head for round two.

The choice was either to integrate with a service like Akismet or Mollom, or to implement a CAPTCHA. I am not very pleased with the former – they do not always work very well and I would end up with a magic box I had no insight into or control over. Implementing a CAPTCHA would decrease the real experience for real commenters and I have too few of those already.

Between reading up on the Akismet API and googling for other innovative techniques, I stumbled upon a blog post by [Ignacio Segura](http://www.isegura.es/), whom explained how to set up a [honeypot trap for Drupal](http://www.isegura.es/blog/stop-spam-your-site-being-invisible-honeytrap-drupal-comments-form). That was just what I was looking for!

A [honeypot trap](http://en.wikipedia.org/wiki/Honeypot_(computing) (or honeytrap) is a way to trick unauthorized attempts to reveal themselves. In computer security this is usually what appears to be a vulnerable computer or network. Attempts to attack a honeypot can safely be disregarded, because no legitimate request would ever go there.

In my case this consists of a hidden field for phone numbers. It is hidden from human visitors using CSS but bots looking at the HTML form only have no means of differentiating this element from a real one. Once it is filled out I know you are a bot and so I can safely throw away the comment.

Again I am crossing my fingers and hoping to be rid of this problem from now on.
