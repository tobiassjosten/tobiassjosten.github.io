---
layout: post
title: The new era of PaaS to host and deploy PHP
category: php
tags: [php, paas]
summary: As a PHP developer I am expected to lament my fate of inconsistent naming and envy the slick hipster world of Ruby (on Rails). Be that as it may, I must honestly say I have missed something like the excellent Heroku platform for hosting and deploying my applications.
---
As a PHP developer I am expected to lament my fate of inconsistent naming and envy the slick hipster world of Ruby (on Rails). Be that as it may, I must honestly say I have missed something like the excellent [Heroku](http://www.heroku.com/) platform for hosting and deploying my applications.

Times are changing for the good though and it definitely feels like, in the words of [Henrik Sjökvist](http://henriksjokvist.net/), that [this is the summer of PHP PaaS](http://twitter.com/henrrrik/status/71602733173587968).

There is a growing number of contenders for the PaaS throne. I have tried two of them out so far; [PHP Fog](https://phpfog.com/) some time ago and [Pagoda](http://www.pagodabox.com/) just a few days ago. And I am amazed by both of them.

## PHP Fog and Pagoda

PHP Fog stands out with pre-configured packages for [Drupal](http://drupal.org/), [WordPress](http://wordpress.org/), [CodeIgniter](http://codeigniter.com/), etc. One click of a button and you have a Git repository to clone and start working from. It is silly easy.

They are very personable as well and you will often receive an answer the very same day you send a question. During beta their CEO, Lucas Carlsson, took the time to have lengthy conversations with me about their architecture and technology.

Pagoda hooks right into [GitHub](https://github.com/), for good and bad. If you are not a Git fan or do not want to use GitHub for some reason then you are out of luck. I love Git and try to use GitHub as much as possible so for me this is a match made in heaven.

There is a web interface to handle maintenance of your app, like deployment (you can pick a specific branch and commit from GitHub), monitoring statistics, moderate databases, etc. Beyond that there is also the [.box file](http://guides.pagodabox.com/getting-started/understanding-the-box-file) with which you can configure the precise PHP version to use, what mods to enable, which shared directories should be avilable and much, much more.

Both PHP Fog and Pagoda are growing and adding new features every day. To this day, however, I think Pagoda have solved the important issue of assets in the most beautiful way. It works with you defining shared directories which are available to write to from all your app instances. These are stored in a network drive which you can SSH into or rsync files from.

Still I recommend you check them both out. They are some very excellent pieces of technology and I am certain they can help you a lot in your day to day life. Unless you are a spoiled [Ruby on Rails](http://rubyonrails.org/) or [Node.js](http://nodejs.org/) developer of course. ;)

## Others

There are a couple of other platforms as well, which I have not checked out yet. If you have taken them for a spin then please let me know what you think!

[Orchestra](http://orchestra.io/) by the industrious team of [Echolibre](http://echolibre.com/).

[Dotcloud](http://www.dotcloud.com/) seems to want to cover *all* relevant web platforms.

[Cloud Control](http://cloudcontrol.com/) has a Heroku inspired Add-on provider program.
