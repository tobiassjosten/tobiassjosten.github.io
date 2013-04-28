---
layout: post
title: Skip THOSE comments in your code
category: development
tags: [development]
summary: You know the kind. Let us break free of them.
---
You know the kind, which adds more noise than they help clarify. It is a plague.

    def display(things):
        # Loop through the things.
        for v in things:
            # Print the thing.
            print v

Comments are supposed to tell something the code fails to communicate. Think about that for a second.

Can you instead change your code around to better express the solution?

    def loop_and_print(blogposts):
        for blogpost in blogposts:
            print blogpost

For example. Perhaps a very lame example but I am sure you get my point. ;)
