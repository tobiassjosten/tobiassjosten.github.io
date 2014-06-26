---
layout: post
title: Add Drupal nodes to the front of nodequeues
category: drupal
tags: [drupal, _drupalplanet]
summary: The default behaviour of Nodequeues is to append new nodes last in its queues. Recently for a client we needed to change this and it turned out to be easy, using Nodequeue's rich API. Still, someone else mightalso have this need so I wanted to share the recipe.
---
The default behaviour of [Nodequeue](http://drupal.org/project/nodequeue) is to append new nodes last in its queues. Recently for a client we needed to change this and it turned out to be easy, using Nodequeue's rich API. Still, someone else might also have this need so I wanted to share the recipe.

    /**
     * Implementation of hook_nodequeue_add().
     *
     * Move all added entries to the top of their nodequeue.
     */
    function mymodule_nodequeue_add($sqid, $nid) {
      $subqueue = subqueue_load($sqid);
      nodequeue_queue_front($subqueue, $subqueue->count);
    }

This hook is invoked when a node is added to a nodequeue. Then we use the API function to move the last node to the front. You could optionally check the $sqid parameter if you want to limit this behaviour to a select few queues.
