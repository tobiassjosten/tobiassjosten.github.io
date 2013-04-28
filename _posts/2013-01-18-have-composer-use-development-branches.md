---
layout: post
title: Have Composer use development branches
category: php
tags: [php, composer]
summary: Using Composer for our PHP projects is a given. But what about when you stumble over a bug in one of those libraries?
---
Using Composer for our [PHP projects](/php/) is a given. It helps immensely in keeping track of our dependencies and we have easy access to [several thousand useful libraries](https://packagist.org/) just by adding them to our `composer.json`.

But what about when you stumble over a bug in one of those libraries? Being [a good open source citizen](/open-source/) you of course patch it up, push to GitHub and send the project your pull request.

That solves the bug but it could take some time before the maintainer accepts your patch, during while we need to be able to move on with the code we know solves our problem. Enter [*custom Composer repositories*](http://getcomposer.org/doc/05-repositories.md#loading-a-package-from-a-vcs-repository).

By defining your own repositories explicitly, Composer will read their branches and use them first, before falling back on the Packagist repository.

## Overriding a package

Say I fix a bug in the `Someone\SomeAwesomeBundle` library on GitHub. I fork the project to my own repository `tobiassjosten\SomeAwesomeBundle`, to which I commit the patch in a `bugfixes` branch.

Next I add my own fork to my `composer.json`.

    "repositories": [
        {
            "type": "vcs",
            "url": "https://github.com/tobiassjosten/SomeAwesomeBundle"
        }
    ]

This will make Composer first look in the given repository for a `composer.json` file, to see which package it supplies. If any is found then that package will take precedence over Packagist's default ones.

Finally, tell Composer to use your custom branch instead of whatever you were using before. Custom branches needs to be prefixed by `dev-`.

    "require" : {
        "php": ">=5.3.3",
        "some/awesome-bundle": "dev-bugfixes"
    }

Now just run `composer update` and Composer will fetch your patch and update `composer.lock` for the rest of your team to use the same.

Just be sure to roll back your `composer.json` once the maintainer merges your fixes.

## Inline aliases

**Update:** [Igor Wiedler](https://github.com/igorw) pointed out that I should propably mention [*inline aliases*](http://getcomposer.org/doc/articles/aliases.md#require-inline-alias); a mechanism that lets you override a package with your custom bugfixes, while keeping the dependency tree intact.

Say for example that the above `SomeAwesomeBundle` is a requirement for another library you are using; the `SomeSuckyBundle`. This sucky piece of code needs at least version 2.0.0 of its awesome dependency. Then if you make it use version `dev-bugfixes` instead, Composer would not be able to compare the version to see if the dependency is satisfied.

That is when you use an inline alias, to make Composer use your custom development branch but treat it as a canonical version.

    "require" : {
        "php": ">=5.3.3",
        "some/awesome-bundle": "dev-bugfixes as 2.0.0"
    }
