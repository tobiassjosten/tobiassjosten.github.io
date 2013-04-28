---
layout: post
title: Performance of PHP functions
category: php
tags: [php, benchmark]
summary: I was recently watching a JSConf.eu session by Thomas Fuchs, in which he goes through some known and some not so known techniques to speed up your JavaScript code. One of the speed impediments was function calls and because it had such a major impact on the performance, I decided to benchmark how it performed in PHP.
---
I was recently watching a [JSConf.eu](http://jsconf.eu/) session called [Extreme JavaScript Performance](http://blip.tv/file/2999333), by [Thomas Fuchs](http://mir.aculo.us/). In it Thomas goes through some known and some not so known techniques to speed up your JavaScript code. I definitely recommend checking it out if you are into JavaScript!

One of the speed impediments, Thomas explained, was functions calls. I had no idea they are such performance killers! Since I am participating in a small [PHP contest](http://www.phpportalen.net/viewtopic.php?t=113904) (swedish), where speed is key, this got me thinking how function calls affects PHP in general and my implementation specifically.

## Benchmarking PHP functions

Following up on my article on [benchmarking array concatenation in PHP](/php/benchmarking-php-array-concatenation), here goes another benchmark post. Let us start with the contenders.

    // Plain variable assignment.
    $i = 1000000;
    while ($i--)
    {
    $x = 'a';
    }

    //php Variable assignment by function return value.
    function a() { return 'a'; }
    $i = 1000000;
    while ($i--)
    {
    $x = a();
    }

To compare the two, I gave  them a couple of runs with [the time tool](http://manpages.ubuntu.com/manpages/maverick/man1/time.1.html). Then I picked their fastest execution for comparison.

Variable assignment was pretty fast at *0m0.170s* while the function call came quite a bit after at *0m2.090s*. I was taken aback. Adding that function call increased the execution time by over twelve times!

## "Real world" implementation

Now I just had to apply this insight to my contest contribution. I went to work and a couple of minutes later I ran my original implementation against the new, functionless one.

Truth be told I was still thinking my function calls would not represent a major part of the execution time. But lo and behold - I managed to shave off a whopping 22%!

The code honestly looks like shit this way and I am probably violating every programming rule and best practice there is. But darn that code flies.
