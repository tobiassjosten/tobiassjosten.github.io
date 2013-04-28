---
layout: post
title: Factorials in Haskell
category: haskell
tags: [haskell]
summary: Lorna Jane posted an example of recursive programming earlier today, using factorials as her example of it. Because factorials is a good example for beginner progammers and since I have just begun programming Haskell myself, I thought it might be fitting to give an example of how to do the same thing she does in PHP, in Haskell.
---
[Lorna Jane](https://twitter.com/lornajane) posted [an example](http://www.lornajane.net/posts/2012/php-recursive-function-example-factorial-numbers) of *recursive programming* earlier today, using *factorials* as her example of it. Because factorials is a good example for beginner progammers and since I have [just begun programming Haskell](/haskell/haskell-list-comprehension-to-color-states) myself, I thought it might be fitting to give an example of how to do the same thing she does in [PHP](/php), in [Haskell](/haskell).

Her example is this.

    function factorial($number) {
        if ($number < 2) {
            return 1;
        } else {
            return ($number * factorial($number-1));
        }
    }

If you wanted to translate it directly to Haskell, it could look something like this.

    factorial n = if n < 2 then 1 else n * factorial (n-1)

## Pattern matching

You could also use pattern matching. This lets you define different behaviors for the same function, depending on what values you are feeding it.

    factorial 0 = 1
    factorial n = n * factorial (n - 1)

## Iteratively

But there are ways to make this even more efficient. As one of Lorna's readers points out in the comments, iteration could be a more elegant approach to solving the problem.

In the following example we use the `foldl` function, which takes three arguments. A [partial function](http://www.haskell.org/haskellwiki/Partial_application) (the multiplication one), a starting number and then list. It applies the function to the starting number and the numbers in the list, one by one.

We build the list by using the range notation, to get one starting at 1 and going up to whatever number is inputted in the `factorial` function.

    factorial n = foldl (*) 1 [1..n]

## Simplicity

But of course the best approach is often the simplest approach. And for **factorials in Haskell** this means using the *`product` function*.

We build the list just like above and then we feed that to `product`, having it multiply each number by the one before it throughout the list.

    factorial n = product [1..n]

Simplicity is divine, is it not?!
