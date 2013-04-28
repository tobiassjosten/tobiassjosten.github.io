---
layout: post
title: Introducing CSSNakedDay.js
category: javascript
tags: [javascript, css-naked-day]
summary: I have participated in a proud web tradition for two years straight now; CSS Naked Day 2011 and 2012. As I have recently moved from Symfony to Jekyll as my website engine, that means I can not do any server-side processing and so I had to rethink the solution.
---
I have participated in a proud web tradition for two years straight now; *CSS Naked Day* [2011](/internet/css-naked-day-2011) and [2012](/internet/i-look-good-naked).

It can be explained as removing the styling veil from the "source code" of the web, exposing the truth beneath. While no longer being eye-candy, the website should still look and work just fine. Unless you have used superficial styling to cover up some fundamental mistakes…

The next occasion will be different though, as I have recently moved from [Symfony](/symfony) to Jekyll as my website engine. That means I can not do any server-side processing outside of the few times I rebuild the entire site.

So I had to rethink the solution.

## Enter CSSNakedDay.js

The result is [**CSSNakedDay.js**](https://github.com/tobiassjosten/CSSNakedDay.js); the *client-side style stripper*!

Basically, it is a piece of [JavaScript](/javascript) which attaches itself to the `onload` event and, when triggered, removes all CSS from the page. It does so by first *disabling all external stylesheets* and then *removing any inline styles*. You do not even get to keep your undies.

Just add [the JavaScript](https://raw.github.com/tobiassjosten/CSSNakedDay.js/master/CSSNakedDay.js) to your page and it will handle the rest.

## Benefits of JavaScript

While working on a solution, I thought of a couple of benefits to using client-side processing with JavaScript, rather than placing the logic server-side.

* Your site should be naked for the user on her April 9th, not depending on what timezone your server is in. Some current solutions works around this by nuding for a full 48 hours, just to be sure. With a client based solution we solve this the right way.

* Doing these computations for every page load, all year around, is unnecessary work for your server. Offload it to the clients instead!

* Stripping away external CSS files is easy but I have yet to see a server side solution remove inline styles, as CSSNakedDay.js does.

* Opposed to other attempts, CSSNakedDay.js have no external dependencies. It is pure JavaScript and as lightweight as it gets.

* Using JavaScript means the HTML will look the same even if it is CSS Naked Day, which in turn means you can cache more to speed up your website.

* If you are on Jekyll or any other [static side generator](/php/installing-phrozn-php-on-ubuntu) then you really do not have an option.

## Get it while it is hot

So [get your copy today](https://github.com/tobiassjosten/CSSNakedDay.js) and [give me feedback](https://twitter.com/tobiassjosten) as you try it out!
