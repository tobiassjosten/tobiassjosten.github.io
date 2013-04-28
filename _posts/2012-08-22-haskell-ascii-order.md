---
layout: post
title: Haskell ASCII order
category: haskell
tags: [haskell]
summary: This weekend I will be running a 5x5 km relay race and in preparation our team must decide on the order of when each runner will have their turn. Being a geek I naturally turned to computer science to help solve the problem.
---
This weekend I will be running a 5x5 km relay race and in preparation our team must decide on the order of when each runner will have their turn. Being a geek I naturally turned to computer science to help solve the problem.

Ordering by the initial letter in your first or last name is boring. Instead I decided on a more interesting approach, where I calculated the *average ASCII value of our names*.

For example *a* has the ASCII value 97, *b* 98 and so on. By adding that up and dividing the sum by the length of the name you have a full name representation and can use that for ordering. Coincidentally, this makes for an excellent opportunity to sharpen my [Haskell skills](/haskell)!

The result is [up at GitHub](https://github.com/tobiassjosten/hs-ascii-order). Feel free to comment on it — I would love some feedback to improve my chops!

## Team order

What of the race? Well, it looks like I got that third position I wanted!

    $ ./AsciiOrder tobias viktor mariya tomas elin
    106 elin
    107 mariya
    107 tobias
    110 tomas
    112 viktor

Now I just have to convince the team about the merits of this highly scientific method.
