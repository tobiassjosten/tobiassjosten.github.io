---
layout: post
title: Continuous integration for your Symfony2 app
category: symfony
tags: [symfony, development, ci]
summary: The hosted continuous integration service Travis CI has been around for a while now, amplifying the testing ambitions of our open source communities. It is really a great service and I encourage you to familiarize yourself with it if you have not already.
---
The hosted *continuous integration service* [Travis CI](http://travis-ci.org/) has been around for a while now, amplifying the testing ambitions of our open source communities. It is really a great service and I encourage you to familiarize yourself with it if you have not already.

For example usage, check out how [Symfony2](http://travis-ci.org/#!/symfony/symfony), [Ruby on Rails](http://travis-ci.org/#!/rails/rails) or even my own [Facebook service provider](http://travis-ci.org/#!/tobiassjosten/FacebookServiceProvider) is doing. Every code change is extensively tested with the project's built-in test suites.

Travis CI is free for open source projects but for your private repositories you have still had to set up Jenkins or similar, to build a CI environment of your own. Until now, when the fine folks at Travis announced [Travis Pro](http://travis-ci.com/) — the same service but for private repositories!

I was recently *invited to Travis PRO* and since I am currently building [a project](http://www.smartburk.se/) using Symfony2, I hooked that up for continuous bliss. Here is how you can too.

## Symfony2 in Travis CI

First you need to create a `.travis.yml` configuration file in the root of your project. Start simple with this bare-bones configuration:

    language: php
    php:
      - 5.3
      - 5.4

This tells Travis that you are running a PHP app and you want it tested for both version 5.3 and 5.4.

Having a PHP app implicates that you are using PHPUnit to test your project. By default this means running `phpunit` in the project root, but since in Symfony2 Standard Edition your PHPUnit configuration resides in `app/`, you will want to tweak how Travis initiates the tests.

    script: phpunit -c app/

Next up is your parameters configuration. Because of its nature it should not be checked in to Git, so what I do is I keep a skeleton file at `app/config/parameters.yml.dist`, which I have Travis copy to the correct location before testing the app.

    before_script:
      - cp app/config/parameters.yml.dist app/config/parameters.yml

Since you are using [Composer](http://getcomposer.org/) (right?!) you can easily have Travis install all your dependencies by adding it to the `before_script` list.

      - composer install

That should be all! Now commit this, go to Travis PRO and add your repository and then push the commit to it. Travis will then pick up your changes and run through your tests. You will get an email about its success or failure and with its very nice integration with GitHub you are able to see directly in Pull Requests the status of the tests.

## Maximum function nesting level

>"Fatal error: Maximum function nesting level of '100' reached, aborting!"

That was one very annoying error I bumped into when I started testing. It turns out this comes from me enabling Symfony2's built-in reverse proxy, where I use ESI tags to render partial bits and pieces of the website.

Fixing it requires setting `xdebug.max_nesting_level` to a high enough value. This is a PHP configuration so we need to do this in two steps; first identifying where to add our configuration and then setting it there.

    before_script:
      - export ADDITIONAL_PATH=`php -i | grep -F --color=never 'Scan this dir for additional .ini files'`
      - echo 'xdebug.max_nesting_level=9999' | sudo tee ${ADDITIONAL_PATH:42}/symfony2.ini

## More Travis

There are of course a lot more you can configure with `.travis.yml` and I have only touched on the basics. Check out [the documentation](http://about.travis-ci.org/docs/user/build-configuration/) for more in-depth options.

For example we have set up the notifications so that Travis sends a message to our Basecamp chat whenever a build fails or succeeds. That is a excellent trigger for us to go code review before merging and deploying.

But why are you still reading this? Go [sign up for **Travis PRO**](http://beta.travis-ci.com/?lrRef=ByOdv) already! :)

**Update:** Apparently [Travis has upped](https://github.com/travis-ci/travis-cookbooks/commit/bd0ea97a72582640b13e9bb03836058076bca173) their `max_nestling_level` to faciliate Symfony2 tests. Thanks, [Damien Alexandre](https://twitter.com/damienalexandre) for letting us know!

**Update:** [Lukas Smith](https://twitter.com/lsmith) explains in the comments that Travis has *built-in support for Composer*, so there is no need to download it first. I have updated the configuration examples as such.
