---
layout: post
title: Compass LoadError in Assetic and Symfony2
category: symfony
tags: [symfony, ruby, php]
summary: I love using Sass to manage my stylesheets. What I do not love, however, is how the frail Ruby stack can mess with your flow.
---
I love using Sass/Compass to manage my stylesheets because it gives so much for free. Like the [px to em calculation](/css/px-to-em-with-sass/) I wrote about a while ago. What I do not love, however, is how the [frail Ruby stack](/ruby/) (RVM, rbenv, etc) can mess with your flow.

Every now and then I bump into the following error when compiling my Symfony2 assets with Assetic.

>ruby: no Ruby script found in input (LoadError) compass

Assetic is kind enough to give you the full command it was trying to run. In my case just now:

    '/home/tobias/.rbenv/shims/ruby' '/home/tobias/.rbenv/shims/compass' 'compile' '/tmp' '--config' '/tmp/assetic_compassdxfSxj' '--sass-dir' '' '--css-dir' '' '/tmp/assetic_compassxJejrM.scss'

## How Assetic uses Compass

The way this should work is the first item is your ruby executable, the second one your Compass Ruby script and the rest is configuration parameters.

Using `which ruby` and `which compass` will tell you where those are installed and this is what Assetic uses itself. So in my case `compass` was actually mapped to `/home/tobias/.rbenv/shims/compass` and `ruby` was mapped to its rbenv shim equivalent.

The problem, it turns out, is well visible in the error message. For me, using rbenv, my `compass` looked like this:

    #!/usr/bin/env bash
    set -e
    export RBENV_ROOT="/home/tobias/.rbenv"
    exec rbenv exec "${0##*/}" "$@"

The error *no Ruby script found in input* was clearly telling me it was expecting a Ruby script but got fed this Bash script. Looking around in `~/.rbenv` I found the actual Compass Ruby script in a different location.

## Configure Assetic and Compass

I had to configure Assetic to use another Compass file. I am showing the Ruby bin configuration as well, just in case you need to tweak that too.

    assetic:
        ruby: /home/tobias/.rbenv/versions/1.9.3-p392/bin/ruby
        filters:
            compass:
                bin: /home/tobias/.rbenv/versions/1.9.3-p392/bin/compass

Obviously you will want to move these to parameters instead, because the path will most likely vary between machines. But that works!
