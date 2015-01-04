---
layout: post
title: Assign by reference with PHP's ternary operator
category: development
tags: [development, php]
summary: If you want to assign a variable with a reference in a ternary operator you run into problems. Syntax errors to be precise.
---
Setting a PHP variable to one of two values, depending on some third value, is easily achieved with a ternary expression.

    $name = isAuthor() ? $authorName : 'John Doe';

But if you want to assign the variable with a reference you run into problems. Syntax errors to be precise.

    $name = isAuthor() ? &$authorName : 'John Doe';

>PHP Parse error:  syntax error, unexpected '&' in …

## Solutions

There are two ways to solve this. Either you assign the variable in the ternary expression.

    $name = isAuthor() ? $name = &$authorName : 'John Doe';

Or pick from one of two references using a dynamic variable name.

    $name = &${isAuthor() ? 'authorName' : 'otherName'};

## Video

The people over at [Webucator](https://www.webucator.com/) have put together [a short video](https://www.youtube.com/watch?v=DojE8Cz9znc) where they go over this solution. They asked if I could link to their [PHP training course](https://www.webucator.com/webdev/php.cfm) and I feel this is good enough SEO work that I'll bite. :)

<iframe width="600" height="337" src="//www.youtube.com/embed/DojE8Cz9znc" frameborder="0" allowfullscreen></iframe>
