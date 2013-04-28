---
layout: post
title: Reproduce Jekyll on GitHub Pages
category: jekyll
tags: [jekyll, github]
summary: Hosting your Jekyll site on GitHub Pages can certainly save a lot of headache. However, being able to reproduce the GitHub environment can sometimes be the only way to solve a bug.
---
Hosting your *Jekyll site on GitHub Pages* can certainly save a lot of headache. Unless something goes wrong. Since your only interface to it is `git push`, being able to reproduce the GitHub environment can be the only way to solve a bug.

There are two gems you will need to align, *Jekyll version 0.11.0* and *Liquid version 2.2.2*. If you are uncertain about your current versions, `gem list` can show you.

    $ gem list|egrep 'liquid|jekyll'
    jekyll (0.11.0)
    liquid (2.2.2)

If they do not match, you will need to remove them before installing the correct versions.

    gem uninstall jekyll liquid
    gem install liquid -v 2.2.2
    gem install jekyll -v 0.11.0

After the installation has completed you should now have the same environment as on GitHub Pages. Just be sure to run Jekyll with the same arguments!

    $ jekyll --pygments --no-lsi --safe
