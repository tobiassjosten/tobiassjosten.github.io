---
layout: post
title: Readline in Ruby with rbenv
category: ruby
tags: [ruby]
summary: Ruby with rbenv does not come with readline support out of the box, so you will need to recompile your version for it to work. Here is how.
---
I just switched to using rbenv for handling my Ruby versions. At face value they are almost identical but I like the rbenv philosophy more, with its lighter footprint on my system.

One problem I just bumped into was when booting up my [continous tests with Guard](/development/automated-continuous-testing-with-guard/), where Ruby threw me the following error.

>You're running a version of ruby with no Readline support  
>Please gem install rb-readline or recompile ruby --with-readline.

Obviously installing the `rb-readline` gem did not work, or this would have been a rather lame blog post. I had a similar issue with [readline in Rails console](/ruby-on-rails/fixing-readline-for-the-ruby-on-rails-console/) but at the time I was using RVM, so now I had to look into a solution for rbenv.

Because it would take a recompilation of Ruby, to include readline, I first had to make sure the dev packages were in place. In Ubuntu this means installing the `libreadline-dev` package.

    $ sudo aptitude install libreadline-dev

After that I set out to recompile Ruby. Using rbenv this is done with `rbenv install <version>` and by assigning the readline flag to `CONFIGURE_OPTS` you get the expected result.

    $ CONFIGURE_OPTS="--with-readline-dir=/usr/include/readline" rbenv install 1.9.3-p392

Check the [Ruby website](http://www.ruby-lang.org/en/) for the latest version or use `rbenv versions` to see what you already have compiled and installed.

Once it has been compiled, be sure to use the readline enabled Ruby version and update the rbenv shims.

    $ rbenv global 1.9.3-p392
    $ rbenv rehash

And there you go, readline support for your Ruby version with rbenv.
