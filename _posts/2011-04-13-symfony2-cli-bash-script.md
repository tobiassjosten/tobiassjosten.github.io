---
layout: post
title: Symfony2 CLI bash script
category: symfony
tags: [symfony, bash, cli]
summary: Since I have been trying out Symfony2 lately, I felt it was time to revise my old bash script for Symfony CLI and have it work with the new version of this lovely framework.
---
Since I have been trying out *Symfony2* lately, I felt it was time to revise my old [bash script for Symfony CLI](/symfony/symfony-cli-helper-script) and have it work with the new version of this lovely framework.

I have [a working copy](https://gist.github.com/275690) up and it turned out a little sexier than my previous script. This time around I am using glob expansion to look for Symfony's files; _symfony_ on Symfony 1.x and _app/console_ on Symfony2. It taught me using the [nullglob shell option](http://www.faqs.org/docs/bashman/bashref_34.html) to have [bash](/bash) keep looking for the other file variation.

If you are using Symfony, whatever version, and are occasionally running commands in your *terminal*, then I recommend you install and check it out! If you are running bash, installation is as simple as:

    $ sudo wget http://bit.ly/osAzEk -O /usr/local/bin/sf
    $ sudo chmod +x /usr/local/bin/sf
