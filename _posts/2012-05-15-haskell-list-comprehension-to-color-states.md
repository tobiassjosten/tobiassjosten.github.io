---
layout: post
title: Haskell list comprehension to color states
category: haskell
tags: [haskell]
summary: I am reading the book Seven Languages in Seven Weeks, by Bruce A. Tate, and my current chapter is on Haskell. This language is very different from what I am used to but I am definitely seeing its allure, with clean design and powerful expression.
---
I am reading the book *Seven Languages in Seven Weeks*, by Bruce A. Tate, and my current chapter is on Haskell. This language is very different from what I am used to but I am definitely seeing its allure, with clean design and powerful expression.

An especially interesting feature is *list comprehension*. I have touched on something similar in Python before but one example homework really drove home the point for me and I wanted to share this epiphany.

Basically it looks like this.

    [x * 2 | x <- [1, 2, 3, 4, 5, 6]]
    >> [2, 4, 6, 8, 10, 12]

That reads out something like: "Iterate through the list of 1, 2, 3, etc. For each value, put that into `x` and then evaluate the expression `x * 2`. Take each evaluation and build a list of that". The result being that list from 2 to 12.

You can add guards (conditions) to this comprehension as well. Let us try adding `even x` to make sure no odd number is being passed on to `x`.

    [x * 2 | x <- [1, 2, 3, 4, 5, 6], even x]
    >> [4, 8, 12]

## Coloring the states

The homework was to color five different states: *Tennessee*, *Mississippi*, *Alabama*, *Georgia* and *Florida*. They are all connected to at least two other states; Alabama being connected to all other states. However you only have three colors and two bordering states must not have the same colors.

List comprehension is really all we need to solve this problem. Have it build a list where all states are assigned a color and then add in guards to make sure the bordering states does not have the same colors.

    [(tennessee, mississippi, alabama, georgia, florida) |
        tennessee <- ["red", "blue", "green"],
        mississippi <- ["red", "blue", "green"],
        alabama <- ["red", "blue", "green"],
        georgia <- ["red", "blue", "green"],
        florida <- ["red", "blue", "green"],
        tennessee /= mississippi,
        tennessee /= alabama,
        tennessee /= georgia,
        mississippi /= alabama,
        alabama /= georgia,
        florida /= alabama,
        florida /= georgia]
    >> [("red","blue","green","blue","red"), ...]

And that is it! You just explain the problem to Haskell and it is nice enough to solve it for you.
