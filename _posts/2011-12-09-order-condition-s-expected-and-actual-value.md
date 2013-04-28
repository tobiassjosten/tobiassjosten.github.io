---
layout: post
title: Order condition's expected and actual value
category: symfony
tags: [symfony, development]
summary: Following Symfony's Coding Standards you should place the expected value before the actual value in an if statement.
---
Following [Symfony's Coding Standards](http://symfony.com/doc/2.0/contributing/code/standards.html) you should place the expected value before the actual value in an if statement.

    if (true === $dummy) {
        return;
    }
    if ('string' === $dummy) {
        $dummy = substr($dummy, 0, 5);
    }

It felt weird to me when I first looked at it. I was used to the other way around and when trying to familiarize myself with the code by reading it out in my mind it went something like: "if blue is the sky".

What the fuck? It made no sense.

But a standard is a standard and you do not want to be one of [*those*](http://xkcd.com/927/), do you? Also it did resemble how I unit tested my code, with the expected value first and the actual value second. So I followed suit and now, some months later, I feel much more comfortable with it.

But I was still curious as to why [Symfony](http://symfony.com/) deviates from most other PHP code I am familiar with. So I sent out [a tweet](https://twitter.com/tobiassjosten/status/145059400632115200) asking for input. Boy do I have some really awesome followers!

## Yoda Conditions

It turns out this is called [Yoda Conditions](http://stackoverflow.com/questions/2349378/new-programming-jargon-you-coined/2430307#2430307). "If blue is the skye" - I am sure you see the reference.

The reason for this idiom is to prevent errors like this.

    if ($dummy = true) {
        return;
    }

In [PHP](/php) this will assign $dummy the value true and the condition evaluated will be $dummy. So that piece of code will always return. Probably not what you wanted.

The scary thing is that the code will run as if nothing is wrong. With some bad luck that can cause quite some problems for you.

If you instead make it a Yoda Condition.

    if (true = $dummy) {
        return;
    }

Then your code will instead crash with a syntax error, saying there is an unexpected '=' and give you the exact file and line.

In both cases your program will not function correctly. I do not know about you but I would much rather know about it.

Thanks [@kallepersson](http://twitter.com/kallepersson), [@freakphp](http://twitter.com/freakphp), [@Mikael0hlsson](http://twitter.com/Mikael0hlsson), [@drrotmos](http://twitter.com/drrotmos), [@igorwesome](http://twitter.com/igorwesome), [@rickard2](http://twitter.com/rickard2) and [@benjick](http://twitter.com/benjick) for helping me figure this out. May with you the force be!
