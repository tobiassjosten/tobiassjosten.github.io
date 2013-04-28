---
layout: post
title: Static class variable for all instances
category: php
tags: [php]
summary: I just got a question on Aardvark asking if you could somehow share a variable value between multiple instances of the same class. Since the value is an array you could not do this with a class constant.
---
I just got [a question on Aardvark](http://vark.com/t/hmDkna) asking if you could somehow share a variable value between multiple instances of the same class. Since the value is an array you could not do this with a class constant.

This could instead be solved by declaring a static variable in the class. That would make it available in all instances of the class. You could then also control encapsulation by declaring it public, protected or private.

Because I took the time to assemble a proof of concept, I felt like sharing it here for future reference.

    class MyClass
    {
        public static $x = 'x';
    }

    $a = new MyClass;
    $b = new MyClass;

    echo sprintf("%s %s\n", $a::$x, $b::$x);
    // prints "x x"

    $a::$x = 'y';

    echo sprintf("%s / %s\n", $a::$x, $b::$x);
    // prints "y y"
