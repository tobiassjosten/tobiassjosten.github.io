---
layout: post
title: PHP REPL for Symfony and Drupal
category: php
tags: [php, drupal, symfony, _drupalplanet]
summary: One thing I have always missed in PHP is a useful and robust REPL. There is the built-in interactive mode (`php -a`) but it does not qualify as neither very useful nor robust.
---
One thing I have always missed in [PHP](/php/) is a useful and robust REPL. A tool you could fire up, throw code at and have it evaluated in real time. There is the built-in interactive mode (`php -a`) but it does not qualify as neither very useful nor robust.

[Boris](https://github.com/d11wtq/boris), from [Chris Corbyn](https://twitter.com/d11wtq), solves this. He describes it as *a tiny little, but robust REPL for PHP* and this is exactly what I have been missing all these years. Thanks, Chris!

There are a couple of ways to [install it](https://github.com/d11wtq/boris#installation) but I prefer using [Composer](/composer/):

    $ composer global require 'd11wtq/boris=*'0

Now you should be able to run `boris` and fire away your code at it.

    [1] boris> 1 + 2 * 3;
     → int(7)

It properly catches errors and lets you carry on without having to restart.

    [1] boris> "one two $three";
    PHP Notice:  Undefined variable: three in phar:///usr/local/bin/boris/lib/Boris/EvalWorker.php(122) : eval()'d code on line 1
    PHP Stack trace:
    PHP   1. {main}() /usr/local/bin/boris:0
    PHP   2. require() /usr/local/bin/boris:10
    PHP   3. Boris\Boris->start() phar:///usr/local/bin/boris/bin/boris:15
    PHP   4. Boris\EvalWorker->start() phar:///usr/local/bin/boris/lib/Boris/Boris.php:139
    PHP   5. eval() phar:///usr/local/bin/boris/lib/Boris/EvalWorker.php:122
     → string(8) "one two "
    [2] boris> "one two three";
     → string(13) "one two three"

## Specialized PHP REPL

This is cool enough but I wanted to take Boris further, so I created a small library to load in project resources. Currently there is support for Symfony (Standard Edition), Drupal (7 and 8), eZ Publish and Composer based apps.

Check out [tobiassjosten/boris-loader](https://github.com/tobiassjosten/boris-loader).

Installation is super simple. Just clone the repository and then add to your `~/.borisrc` configuration file the following content:

    <?php
    require 'path/to/cloned/boris-loader.php';
    \Boris\Loader\Loader::load($boris);

And that it is; your Boris instance should now be able to recognize your standard Composer based projects. For more specialized projects you will need to load their providers explicitly.

## Symfony REPL

Tell boris-loader to use the Symfony provider by tweaking your `.borisrc` file.

    \Boris\Loader\Loader::load($boris, [new \Boris\Loader\Provider\Symfony2()]);

Now, starting Boris within a Symfony project will take care of autoloading and build a Symfony kernel for you. Both the kernel and its container will be available in Boris and you can tell it has loaded by the `symfony>` prompt.

    [1] symfony> $container->get('doctrine');
     → class Doctrine\Bundle\DoctrineBundle\Registry#332 (7) {
      protected $container =>
      class appDevDebugProjectContainer#337 (10) {
        protected $parameterBag =>
        NULL
    // …

For Symfony I am assuming you are using the Standard Edition, with `app/bootstrap.php.cache` and `app/AppKernel.php` files.

## Drupal REPL

Have boris-loader use the Drupal providers adding them to your `.borisrc` file.

    \Boris\Loader\Loader::load($boris, [
        new \Boris\Loader\Provider\Drupal7(),
        new \Boris\Loader\Provider\Drupal8()
    ]);

If you are standing in a Drupal project when firing up Boris with those providers, you will land in a fully bootstrapped Drupal instance. As with the Symfony provider you have access to Drupal's `$kernel` and `$container` and of course also the entire Drupal API.

    [1] drupal> node_load(1);
     → class Drupal\Core\Entity\EntityBCDecorator#1158 (2) {
      protected $decorated =>
      class Drupal\node\Plugin\Core\Entity\Node#1156 (9) {
        protected $bundle =>
        string(7) "article"
        protected $values =>
        array(21) {
    // …

## REPL for other PHP projects

With Symfony, Drupal, eZ Publish and generic Composer support, I am covering my own needs but not likely all of yours. The loader is open source though, so just send me a pull request for your favorite project!

Check [the project out at GitHub](https://github.com/tobiassjosten/boris-loader).
