---
layout: post
title: Booting Jekyll with Teamocil
category: jekyll
tags: [jekyll]
summary: Teamocil is a cool utility for us tmux users. It helps manage your sessions, windows and panes in a way that makes it really easy to boot up preconfigured environments.
---
[Teamocil](https://github.com/remiprev/teamocil) is a cool utility for us [tmux](/tmux/) users. It helps manage your sessions, windows and panes in a way that makes it really easy to boot up preconfigured environments.

You could use *teamocil* to quickly start [Vim](/vim/) in a split window with your continous tests running in an adjacent pane. Or you could SSH into all your servers in different panes, boot up `htop` and use teamocil as a wicked performance dashboard. The possibilities are endless.

I am using teamocil right now to write this blog post, having first booted up [my Jekyll server](/jekyll/) and [Guard instance](/development/automated-continuous-testing-with-guard/) in a separate window.

## Installing teamocil

Teamocil is a [Ruby gem](/ruby/) and as such is dead simple to install.

    $ gem install teamocil
    $ mkdir ~/.teamocil

## Configuring Teamocil for Jekyll

With Teamocil you can have multiple settings, each with a unique name in its separate YAML configuration file. Starting or editing one is done with the same command.

    $ teamocil --edit jekyll

This will launch the `jekyll` configuration file in your `$EDITOR` of choice. For my Jekyll configuration, I use the following.

    windows:
      - layout: even-horizontal
        clear: true
        panes:
          - cmd: "jekyll serve"
          - cmd: "guard"

What it does is it creates a new window with two panes, one running `jekyll serve` to start the Jekyll server and another running `guard` to initialize my Guard rules.

Save and then try starting it.

    $ teamocil jekyll
