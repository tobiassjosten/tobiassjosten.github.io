---
layout: post
title: Music Hack Day Reykjavik 2012
category: events
tags: [events, javascript, php, symfony]
summary: Together with Viktor Miranda I have flewn to Iceland and Reykjavik to participate in Music Hack Day Reykjavik 2012. This is an event where hackers, designers and musicians travel from all around the world to meet up and create cool stuff somehow related to music.
---
Together with [Viktor Miranda](http://www.viktormiranda.com/) I have flewn to Iceland and Reykjavik to attend [Music Hack Day Reykjavik 2012](http://reykjavik.musichackday.org/2012/). This is an event where hackers, designers and musicians travel from all around the world to meet up and create cool stuff somehow related to music.

Me and Viktor are currently heavily involved in building a [tv-series startup](http://www.smartburk.se/) but felt it would be a good time to take a short break and do something else for a few days. Participating in Music Hack Day and visiting the beautiful Iceland seems like a great pause.

Our project for *Music Hack Day* is called [Songbrew](http://www.songbrew.com/) and our aim is to solve the problem of getting stuck listening to the same music for months and years. We are scratching our own itch but hope we can also help you discover new music.

## The tech stack

We want Songbrew to exist in a wide array of channels; like the web, smartphones, music players, etc. For that reason we are building the core in [a HTTP API](/api/), which will then be consumed by our different clients.

This API will be built on [Symfony2](/symfony/), hosted by [Pagoda Box](https://pagodabox.com/) and will use [Redis](/redis/) for persistance.

We will start with a web application client, built as a single-page-app with heavy [JavaScript](/) focus. Me and Viktor have both been eyeing [Backbone.js](http://backbonejs.org/) and so that is a given foundation, along with [jQuery](/jquery/) for utility. For templating we will be using the excellent [Handlebars.js](http://handlebarsjs.com/) library and everything will be made modular and nice using [RequireJS](http://requirejs.org/).

This web application client will also live on Pagoda Box but with [CloudFlare](https://www.cloudflare.com/) for edge-side delivery. We are thinking all requests will result in the same static HTML, JavaScript and CSS files being served, with routing and logic then kicking in on the client side.

## Ready, set, go

We bought the domain, drew the wireframes, specced the API, created the repositories, deployed the empty sites and we are now ready to start hacking!

You like music? Then [go sign up for Songbrew](http://signup.songbrew.com/)!
