---
layout: post
title: Dispatching Drupal hooks
category: drupal
tags: [drupal, _drupalplanet]
summary: The word *function* implies its intended scope. A function should strive to leverage exactly one functionality. If you are cramming more than that into your function, then you probably need to do some refactoring. One situation where you are often forced to go against this is when implementing hook_block() in Drupal.
---
The word *function* implies its intended scope. A function should strive to leverage exactly one functionality. If you are cramming more than that into your function, then you probably need to do some refactoring. I think [Linus Torvalds](http://en.wikipedia.org/wiki/Linus_Torvalds) [hinted towards this](http://www.kernel.org/doc/Documentation/CodingStyle) quite eloquent.

> If you need more than 3 levels of indentation, you're screwed anyway, and should fix your program.

One situation where I seldom can obey that rule of thumb is when implementing [`hook_block()`](http://api.drupal.org/api/drupal/developer--hooks--core.php/function/hook_block/6) in a [Drupal](http://drupal.org/) module. That hook is first called once to fetch a list of your blocks. Then it is called again to fetch a configuration form for your block. And again to save that configuration. And then yet a fourth time to fetch the actual block content. Imagine growing more than two blocks or starting to use block configuration for your module. It is all downhill from there.

So how do we rectify this situation, besides whining about it and not submitting a patch? We dispatch the function!

## Enter function callbacks

PHP has this feature named [function callbacks](http://www.php.net/manual/en/language.pseudo-types.php#language.types.callback) which I really recommend reading up on if you are not familiar with it. It will enable us to elegantly split up our function into multiple, smaller functions.

Let us start with the hook implementation.

    /**
     * Implementation of hook_block().
     */
    function mymodule_block($op = 'list', $delta = '', $edit = array()) {
      if ($op == 'list') {
        $function = 'mymodule_block_list';
      }
      else {
        $function = sprintf('mymodule_block_%s_%s', $op, $delta);
      }

      if (function_exists($function)) {
        return $function($edit);
      }
    }

So what goes on here? First we set the `$function` variable to either a function name for the block listing or a function name including the operation (configure, save or view) and the specific block delta (ID) being requested. When we have our function name, we check to see that it exists before calling it.

The next step is to implement our list and view callback. In our example we will define just one, but one very special, block.

    /**
     * Implementation of hook_block():list.
     */
    function mymodule_block_list() {
      return array(
        'myblock' => array(
          'info' => t('My own, very special block'),
        ),
      );
    }

    /**
     * Implementation of hook_block():view:myblock.
     */
    function mymodule_block_view_myblock() {
      return array(
        'subject' => t('Special block'),
        'content' => t('Obviously, this is a very special block.'),
      );
    }

The `mymodule_block_list()` function will return our list of blocks when `mymodule_block('list')` is called, and block.module will then learn about our special block. When it is time to render the block, `mymodule_block('view', 'myblock')` will be called and this will in turn call our new `mymodule_block_view_myblock()` function.

So there we go – three slim functions instead of a big, bloated one. That is a win in my book!
