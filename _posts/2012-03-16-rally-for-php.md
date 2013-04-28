---
layout: post
title: Rally for PHP
category: php
tags: [php]
summary: The PHP community has a reputation for being fragmented and unable to pool together. It is often seen as the language where wheel reinventing is a recognized sport. While this might have been true 5+ years ago, there is now a strong movement to improve.
---
The PHP community has a reputation for being fragmented and unable to pool together. It is often seen as the language where wheel reinventing is a recognized sport.

While this might have been true 5+ years ago, there is now a strong movement to improve.

With the common [PSR-0 standard](https://gist.github.com/1234504) for structuring PHP code and initiatives like [Composer](http://getcomposer.org/) for package management, we are slowly turning this beast to mutual beneficial mode. Add to that the increasingly common component approach from modern PHP frameworks, with adoption across widely different pools, and I am sure you will start seeing the trend as well.

The most recent example of this trend is [Lukas Kahwe Smith](http://pooteeweet.org/)'s pull request to the Facebook PHP SDK, where [he adds support](https://github.com/facebook/facebook-php-sdk/pull/12) for the composer library. Within the last 10+ hours it has already received the attention and encouragement from well over a hundred developers.

The technical benefit of this improvement is that you will be able to more easily use and include Facebook's SDK in your application. Very nice indeed. Especially for projects like [my FacebookServiceProvider](https://github.com/tobiassjosten/FacebookServiceProvider).

For the bigger picture however, this will help cement a very vital part of our new PHP world and send a message that PHP is still a strong contender in the web development space.

If Facebook recognizes the opportunity and merges it, that is. So let's rally together and [let them know we want this](https://github.com/facebook/facebook-php-sdk/pull/12)!

*Update:* [Scott MacVicar](https://github.com/scottmac) just merged the pull request, adding the [composer.js](https://github.com/facebook/facebook-php-sdk/blob/master/composer.json) definition file.

*Update:* I wrote a quick howto to help you get started with [Composer and the Facebook PHP SDK](/php/facebook-php-sdk-with-composer).
