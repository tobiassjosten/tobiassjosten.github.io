---
layout: post
title: Functional primitives for PHP
category: php
tags: [php]
summary: I believe your imperative code can benefit from a more functional approach and to help with this; there is Functional primitives for PHP.
---
Having coded [functional Haskell](/haskell/) for a while now, I feel my imperative coding style is changing as a result. This primarily means [PHP](/php/) nowadays. While I really do enjoy a lot of aspects of this language, I will be the first to admit it lacks certain elegance.

There is no [list comprehension](http://vvv.tobiassjosten.net/haskell/haskell-list-comprehension-to-color-states/), partial function application is bothersome, you can not overload methods, etc. PHP just is not designed for functional programming. Still, I believe your code can benefit from taking on a more functional approach.

To help with this endeavor; there is *Functional primitives for PHP* — a PHP library for coding in a functional and consistent manner.

- It does not only work with arrays, but anything implementing the `Traversable` interface.

- All its functions consistently takes a collection as a first parameter and a callback as a second parameter.

- Callbacks are always given the current value, its index and the entire collection.

- Any callable is a valid callback, whether that is a string function, an array with an object method or a closure.

- There is both a userland (Composer managed!) PHP library and an extension in C for performance.

Go check it out at [its GitHub repository](https://github.com/lstrojny/functional-php)! Kudos to [Igor Wiedler](https://twitter.com/igorwesome) for [recommending it](https://twitter.com/igorwesome/status/301717986916577280).
