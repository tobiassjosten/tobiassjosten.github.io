---
layout: post
title: Partial Doctrine entities
category: doctrine
tags: [doctrine, symfony]
summary: The way to have Doctrine fetch only parts of your entities, and still have them properly hydrated, is to use the partial keyword.
---
While porting an ooold [MUD community](http://nogfx.org/) site of mine to [Symfony](/symfony/), I wanted to (admittedly prematurely) optimize its database queries. Here is how I made [Doctrine](/doctrine/) comply.

One view on the new site iterates over a bunch of `Log` entities and prints their `title` fields. These entities also have a `body` field, which can be HUGE. Since the `body` is not used on this particular view it means a lot of overhead fetching all that data.

## Selecting fields

This could be solved by `SELECT`ing just the fields you want.

    SELECT l.id, l.title FROM Log l

It gives you the data you want but Doctrine does not know exactly what you are trying to achieve, so it will play it safe and return an array with the data. I do not subscribe to [Array Oriented Programming](http://epixa.com/2012/04/array-oriented-programming.html) and so I want properly hydrate entities.

## Partial objects

It turns out the solution is really simple. We just need to tell Doctrine of our intentions to fetch just part of the entity, by using the `partial` DQL keyword.

    SELECT partial l{id, title} FROM Log l

Word of warning though: Doctrine will not lazily fetch the omitted fields, like it does for associations. If you skip a field in the DQL query you will have to live without it in this result set.
