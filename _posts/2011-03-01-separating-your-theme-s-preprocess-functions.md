---
layout: post
title: Separating your theme's preprocess functions
category: drupal
tags: [drupal, _drupalplanet]
summary: Preprocess functions are often a major part of any Drupal theme. It is a vital part of themes but more than often it leads to very big `template.php` files. Today I stumbled upon a simple but brilliant way to reduce its size by separating its logic into nicely shaped and separated files.
---
Preprocess functions are often a major part of any Drupal theme. It is a vital part of themes but more than often it leads to very big `template.php` files. Today I stumbled upon a simple but brilliant way to reduce its size by separating its logic into nicely shaped and separated files.

It was when working together with [Pelle Wessman](http://kodfabrik.se/) at [Good Old](http://goodold.se/) that I noticed something in his template.php (slightly modified):

    function mytheme_preprocess(&$vars, $hook) {
      $filename = sprintf(
        '%s/preprocesses/preprocess-%s.inc',
        drupal_get_path('theme', 'mytheme'),
        str_replace('_', '-', $hook)
      );

      if (is_file($filename)) {
        include($filename);
      }
    }

Then there was an accompanying preprocesses directory with files like `preprocess-block.inc`.

    $block = $vars['block'];

    $vars['attributes'] = array();
    $vars['attributes']['class'] = 'block';

    if (isset($block->view)) {
      $vars['attributes']['class'] .= ' block-view-' . views_css_safe($block->view->name);
    }

Because the code is executed from within the scope of the `hook_preprocess()` implementation it has access to all its variables.

Brilliant, is it not? [Apparently](http://twitter.com/voxpelli/statuses/42661571473842176) the idea originally comes from [the Studio base theme](http://drupal.org/project/studio). I just love finding these small improvements!
