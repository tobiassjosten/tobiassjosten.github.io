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
