---
layout: post
title: Hosting Silex apps on PHP PaaS Pagoda Box
category: php
tags: [php, paas, silex, pagodabox]
summary: The PHP microframework Silex is an outstanding piece of software and Pagoda is a PHP PaaS I have grown quite fond of. Obviously marrying these two would be ideal.
---
The PHP microframework [Silex](http://silex-project.org/) is an outstanding piece of software. I use it for a few of my personal projects and I love how easy it is to work with and yet how powerful and flexible it is.

Then there is [Pagoda](http://www.pagodabox.com/), a [PHP PaaS](/php/the-new-era-of-paas-to-host-and-deploy-php) I have grown quite fond of. Similar to Silex it lets me focus on the nuts and bolts of my applications.

Obviously marrying these two would be ideal.

The way Pagoda does configuration is via [a .box file](http://guides.pagodabox.com/getting-started/understanding-the-box-file) that you add to your project. They have [gathered examples](https://github.com/pagodabox) for a variety of PHP frameworks and are apparently looking to expand that list. Yesterday they asked me if I would like to help with a Silex box and of course I would!

So far I have taken a quick stab at it and uploaded [the results to GitHub](https://github.com/tobiassjosten/silex-example). I am by no means an expert on Silex and so I am hoping some who are could come and help improve it. Let us make it easy to get started with Silex on Pagoda!

I have also uploaded the example to Pagoda, where it is [publically visible](http://silex-example.pagodabox.com/).

If you can find something to improve or give any feedback what so ever, please feel free to either drop a comment here or send me a pull request on GitHub.
