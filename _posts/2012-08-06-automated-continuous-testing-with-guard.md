---
layout: post
title: Automated, continuous testing with Guard
category: development
tags: [development, ruby, haskell]
summary: Everything you are interacting with has an interface; whether it is a laptop touchpad, buttons on a touchscreen or even the handle of a hammer. The quality of this interface determines the pleasure of the interaction, and that it why I believe the most immersive way of gaming is on a MUD.
---
When doing *test driven development* I really like having continuous feedback. Does the code compile, is all the functionality still intact or what errors are there?

I could of course jump between terminal windows and manually execute the tests but that is not a very good use of my time. It needs to happen automatically and so far I have solved this with *inotifywait*.

    $ inotifywait -mre modify src | while read LINE; do runghc $LINE; done

That command watches the `src/` directory for changes and executes `runghc` on the changed files.

What I would do is open up a [tmux](/tmux) pane, next to my [Vim](/vim) pane, where I would run inotifywait. That way, whenever I saved a file in Vim the script would run and I could immediately see if the code was any good.

## Enter Guard

It turns out there is a better way than to muck about with [bash scripts](/bash). A project called [Guard](https://github.com/guard/guard/) does exactly this, but in a more controlled and reusable manner.

As with most new, cool software, *Guard* is written in [Ruby](/ruby) and distributed as a gem. Thus installation is very easy; `gem install guard` should get you everything you need.

Once installed you need to set up a `Guardfile` to hold your configuration. In this file you describe what to watch and what actions to take when files are changed. The many *Guard plugins* will help with this but most of them are geared toward Ruby projects in general and [Rails](/ruby-on-rails) one specifically.

Converting my above [Haskell](/haskell) tester above, using the *shell plugin*, would look something like this.

    guard :shell do
      watch(%r{^src/.+}) do |m|
        `runghc #{m[0]}`
      end
    end

If you are using [Symfony](/symfony) or any other [PHP](/php) framework, there is [a PHPUnit Guard plugin](https://github.com/Maher4Ever/guard-phpunit) you can use to automatically run your *PHPUnit test suites*.

## More than testing

Guard's usefulness does not stop at automating tests however. I am currently using it to write this blog post, where I have set it up to regenerate my Jekyll site whenever a file is changed.

You could also compile SASS/LESS code into CSS, CoffeeScript into JavaScript or whatever other automation you can think of.

Guard is a really nice general-purpose development tool. Do yourself a favor and [check it out](https://github.com/guard/guard/) today!
