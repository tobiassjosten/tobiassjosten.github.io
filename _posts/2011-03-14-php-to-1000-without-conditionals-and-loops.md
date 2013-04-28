---
layout: post
title: PHP to 1000 without conditionals and loops
category: php
tags: [php]
summary: Earlier today I stumbled on a Stackoverflow challenge where the poster asked for C/C++ code that would count from 1 to 1000 without using conditionals or loops. Then just now I noticed how Justin Vincent at Pluggio had done the same in PHP, my interest sparked. A moment later I now have my own code to meet the challenge.
---
Earlier today I stumbled on [a Stackoverflow challenge](http://stackoverflow.com/questions/4568645/printing-1-to-1000-without-loop-or-conditionals/4583502) where the poster asked for C/C++ code that would count from 1 to 1000 without using conditionals or loops. There are some crazy examples in there and I wont presume to understand even half of them.

So when I just now noticed how [Justin Vincent](http://twitter.com/justinvincent) at [Pluggio](http://pluggio.com/) had done [the same in PHP](http://justinvincent.com/page/1381/1-to-1000-in-php-with-no-conditionals-or-loops), my interest sparked. A moment later I now have my own code to meet the challenge.

    ini_set('xdebug.max_nesting_level', 1002);
    function x($i) { echo "$i\n"; @x(++$i); }
    x(1);

It leverages the built-in protection for infinite loops, [max_nesting_level](http://xdebug.org/docs/basic#max_nesting_level). You have probably seen the error before some time, when going too wild on recursive functions.

>`PHP Fatal error:  Maximum function nesting level of '1002' reached, aborting!`

I set the max nesting level to just above what I need and then, using *error supression* (the `@` character), I make sure the fatal error does nothing else but silently kill the program.

So there you have it, counting from 1 to 1000 in PHP without using conditionals or loops!

**Update:** Stackoverflow now has [a post up](http://stackoverflow.com/questions/5305156/printing-1-to-1000-without-loop-or-conditionals-in-php) with this challenge.

**Update:** Be sure to check out [Andrew Curioso's](http://andrewcurioso.com/) excellent [solution to the problem](http://andrewcurioso.com/2011/03/counting-to-1000-in-php-without-loops-or-conditionals/).
