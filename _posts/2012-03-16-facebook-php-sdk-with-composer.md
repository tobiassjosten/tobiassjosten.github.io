---
layout: post
title: Facebook PHP SDK with Composer
category: php
tags: [php, facebook]
summary: Just a little while ago, Facebook merged a pull request which enables developers to use Composer to handle their dependencies to the Facebook PHP SDK. Since I wrote about this earlier, I thought it would be fitting to follow up with a howto now that it is in. Let us dive right into it.
---
Just a little while ago, [Facebook merged a pull request](https://github.com/facebook/facebook-php-sdk/pull/12#issuecomment-4547036) which enables developers to use [Composer](http://getcomposer.org/) to handle their dependencies to the Facebook PHP SDK.

Since I [wrote about this earlier](http://vvv.tobiassjosten.net/php/rally-for-php), I thought it would be fitting to follow up with a howto now that it is in. Let us dive right into it.

First you need to create your `composer.json` file, which will describe the dependencies for your application. Place this in the root of your project.

    {
        "require": {
            "facebook/php-sdk": "dev-master"
        }
    }

The package name `facebook/php-sdk` can be found on [Packagist](http://packagist.org/packages/facebook/php-sdk). We are using the *dev-master* version here, which is the only one available for this project.

Next you will have to actually download Composer - in the form of a file named `composer.phar`. This is easily done by executing an install script from the Composer site.

    $ curl -s http://getcomposer.org/installer | php

Now all that remains is to run Composer to have your dependencies installed.

    $ php composer.phar install

That should give you a file named `composer.lock` and a directory named `vendor`. The lock file sets the dependency to one specific revision, so that all your application's developers use the same code. In the vendor directory you will find the Facebook SDK.

Pretty nice, eh? But the goodies does not stop there. You can also use the generated `autoload.php` file to get automatic class autoload.

    <?php
    require 'vendor/.composer/autoload.php';
    $facebook = new Facebook(array(
        'appId' => '1234567890',
        'secret' => 'asdfghjkl',
    ));

Done!
