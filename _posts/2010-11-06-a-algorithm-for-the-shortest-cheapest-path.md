---
layout: post
title: A* algorithm for the shortest/cheapest path
category: php
tags: [php, ai]
summary: Yesterday I submitted my contribution to that PHP contest I have been writing about. I think this kind of event is a great opportunity to improve ones programming skills and since the deadline is over I thought I would share my code and thoughts on it.
---
Yesterday I submitted my contribution to [that PHP contest](http://www.phpportalen.net/viewtopic.php?t=113904) I [have been](/php/benchmarking-php-array-concatenation) [writing about](/php/performance-of-php-functions). I think this kind of event is a great opportunity to improve ones programming skills and since the deadline is over I thought I would share my code and thoughts on it.

The whole contribution is [published at GitHub](https://gist.github.com/665458) under [the GPL license](http://en.wikipedia.org/wiki/GNU_General_Public_License).

## The A* algorithm

I have previously wrestled with the problem of finding the shortest path for my [Protea project](https://github.com/tobiassjosten/protea). Back then I experimented with a couple of renowned solutions like [Djikstra's algorithm](http://en.wikipedia.org/wiki/Dijkstra's_algorithm), [Bi-directional search](http://en.wikipedia.org/wiki/Bidirectional_search) and of course the [A* search algorithm](http://en.wikipedia.org/wiki/A*_search_algorithm). I picked A* then and since it had really proven itself I decided to implement it for this contest as well.

A* is a heuristic algorithm. That means basically three things; that it is iterative, that it does not guarantee 100% optimization and that it is very fast. Just the edge I wanted for this speed contest!

The way A* works iteratively is that it is based on a list of open nodes. The first node in this list is the starting position and it has all its adjacent nodes added to the list, before it is moved to the closed list. Then the next node is picked from the open list and its adjacent nodes are added. You do this until you have found the end node, at which point you have the best path.

It sounds almost too simple but the magic lies in how A* works its open list. You always pick the lowest valued node from the list and the value is calculated by adding the actual cost of the node to the approximate proximity to the end node. That way you can be certain that once you hit the end, you will have the cheapest path.

Because for this contest there was no cost associated with movement, only that of the nodes, I did not consider proximity to the end node at all. In games you often do this though, because it will take longer to move a hundred meters than only fifty.

## Chasing speed

I did a lot of experimentation to achieve as fast a execution as possible. Great material for my articles on [array concatenation](/php/benchmarking-php-array-concatenation) and [PHP functions](/php/performance-of-php-functions)! I think the biggest optimization was throwing all internal function calls out the window. If you have a look at [the code](https://gist.github.com/665458) you will see that I have commented sections like so:

    /* START QUEUE ADJACENTS */

This is where the last of my function calls were made and I decided to leave that in so I would not get lost in my own code. Because removing those functions was the final nail in the coffin for readability.

There was one thing I never got around to optimize though – the order adjacent nodes are added. They are currently added top, right, bottom and left (like CSS does it) but I had wanted each node to consider its position in relation to the end node. It might have increased the execution time, because of the extra computation, or it could have decreased it. I am not sure but it would have been fun to measure that.

Now I will just have to wait and see if my improvements got me far enough.
