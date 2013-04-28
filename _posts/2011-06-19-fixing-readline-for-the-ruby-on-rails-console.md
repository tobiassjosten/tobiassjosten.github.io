---
layout: post
title: Fixing readline for the Ruby on Rails console
category: ruby-on-rails
tags: [ruby-on-rails, ruby, cli]
summary: I have begun trying out the Ruby on Rails framework. While doing so I got stuck at a very early stage in the documentation.
---
I have begun trying out [the Ruby on Rails framework](http://rubyonrails.org/). While doing so I got stuck at a [very early stage](http://guides.rubyonrails.org/getting_started.html#using-the-console) in the [documentation](http://guides.rubyonrails.org/).

    $ rails console
    /home/tobias/.rvm/rubies/ruby-1.9.2-p180/lib/ruby/1.9.1/irb/completion.rb:9:in `require': no such file to load -- readline (LoadError)
    from /home/tobias/.rvm/rubies/ruby-1.9.2-p18080/lib/ruby/1.9.1/irb/completion.rb:9:in `<top (required)>'
    from /home/tobias/.rvm/gems/ruby-1.9.2-p180/gems/railties-3.0.9/lib/rails/commands/console.rb:3:in `require'
    from /home/tobias/.rvm/gems/ruby-1.9.2-p180/gemsems/railties-3.0.9/lib/rails/commands/console.rb:3:in `<top (required)>'
    from /home/tobias/.rvm/gems/ruby-1.9.2-p180/gems/railties-3.0.9/lib/rails/commands.commandsrb:20:in `require'
    from /home/tobias/.rvm/gems/ruby-1.9.2-p180/gems/railsilties-3.0.9/lib/rails/commands.rb:20:in `<top (required)>'
    from scriptt/rails:6:in `require'
    from script/rails:6:in `<main>'

The problem is that Ruby is somehow not cofigured to load the readline library and thus cannot run its console. This is a downside of using a language specific package manager, seperate from your operating system's package manager.

However it is easy to fix and after some searching I found [a blog post](http://dirk.net/2009/04/05/no-such-file-to-load-readline-loaderror-when-running-scriptconsole/) detailing exactly how to solve it.

First you need to install `libreadline5` and `libncurses5`. I can't see why `ncurses` would be needed, but I am too lazy to uninstall it and check if it still works without the library. Using Apt on Ubuntu/Debian:

    $ sudo aptitude install libreadline5-dev libncurses5-dev

Next you want to go configure and rebuild Ruby's readline extension. I am using [RVM](https://rvm.beginrescueend.com/) and so Ruby exists in `~/.rvm/src/ruby-1.9.2-p180/` for me. This can differ for you but running *which ruby* should tell you where yours is.

    cd ~/.rvm/src/ruby-1.9.2-p180/ext/readline
    ruby extconf.rb
    make
    make install

Depending on your setup you might need to install it by running `sudo make install` instead.

That is it! Now Ruby on Rails' console should load as intended.
