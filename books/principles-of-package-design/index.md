---
layout: book
title: Principles of Package Design
category: development
tags: [development, php, symfony]
amazonURL: /books/principles-of-package-design/buy/
image: principles-of-package-design.jpg
summary: I'll start at the end and say that this is one of the most useful programming books I have ever read and it should be required reading for all professional programmers.
grade: 5
---
I'll start at the end and say that this is one of the most useful programming books I have ever read. [*Principles of Package Design*](/books/principles-of-package-design/buy/) should be required reading for anyone writing code for a living, independent of one's programming language.

[Matthias Noback](https://twitter.com/matthiasnoback) is a well known figure ([and blogger](http://php-and-symfony.matthiasnoback.nl/)) in the [Symfony world](/symfony/), where he's been pioneering the concept of [hexagonal architecture](http://hexagonal-symfony.eu/) and developing [some very cool libraries](http://simplebus.github.io/MessageBus/). So I have of course been looking forward to reading his book – **Principles of Package Design**.

I was not disappointed, whatsoever.

The book starts off with a thorough introduction to the SOLID principles of object oriented programming. Best practices, really, for how to structure your code in a way that makes it as flexible as possible. A lot of thought goes into keeping the code simple, so as to mitigate as many future bugs as possible.

Single responsibility, Open-closed, Liskov substitution, Interface segregation and Dependency inversion – the book goes through all the SOLID principles and explains why they exist, how they are used, what are signs of them being violated and how you can fix it.

I'm hit by how hands-on and pedagogic it is. Where some would focus on showing off their prowess, Matthias assumes nothing about your skill level and seems to genuinely want to teach the ins and outs of this subject. Anyone could pick this up and improve their OOP game a great deal.

>And still every day there's something new to learn about class design, some old habit to drop, some new principle to apply.  
>- **Matthias Noback, Principles of Package Design**

## Package Design

38% in, the book turns its focus to the meat of the book – packaging.

The package design principles were originally conceived by [Uncle Bob](https://twitter.com/unclebobmartin) ([Robert C. martin](http://en.wikipedia.org/wiki/Robert_Cecil_Martin)) and are divided into groups of three:

*Cohesion principles*:

- Release/Reuse Equivalency Principle
- Common Closure Principle
- Common Reuse Principle

*Coupling principles*:

- Acyclic Dependencies Principle
- Stable Dependencies Principle
- Stable Abstractions Principle

While having well designed classes is obviously a great foundation for well designed packages, it now becomes clear why major focus was on the SOLID principles. Matthias draws some really interesting parallels between class and package structures and shows how in some cases the exact same logic applies, just on a different level.

Again, I'm struck by the attention to details, as the book goes from birds eye view of package architecture to details about how to structure your README file and [format the CHANGELOG](http://keepachangelog.com/).

To exemplify the packaging principles, Matthias shows some real world projects and how they could be improved. Being use cases we are all very familiar with, no matter if you have used these exact libraries, this makes it very simple to understand how these small tweaks can make your overall architecture more robust.

In another example we are shown two interdependent classes and taught how to untangle them from each other. As the classes are transformed and fixed, we jump between code and a more high level dependency graph. This approach made both the problem and solution abundantly clear.

## Summary

I learned a lot reading this book and am already planning on re-reading it, as I'm sure there are more things to learn from it.

If you are considering getting Principles of Package Design – stop right now and do it. I can't recommend it enough.
