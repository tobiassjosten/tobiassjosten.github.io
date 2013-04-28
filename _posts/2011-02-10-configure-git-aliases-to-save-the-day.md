---
layout: post
title: Configure Git aliases to save the day
category: git
tags: [git, bash]
summary: My morning started out great today. The sun came out solely to greet me, birds were chirping my praise and the coffee tasted especially well. You know the scenario. And then this...
---
My morning started out great today. The sun came out solely to greet me, birds were chirping my praise and the coffee tasted especially well. You know the scenario.

And then this.

    $ git statsu
    git: 'statsu' is not a git command. See 'git --help'.

    Did you mean this?
    statsu

It could ruin the best of days. So I am making sure my good spirits are forever kept high so my fingers can keep soaring the keyboard. And you should too.

Let us start by fixing this particular problem. It is done by one simple line.

    $ git config --global alias.statsu status

This will add an *statsu* alias to the global Git configuration. When invoked it will simply just run *status* instead. If you have a look at `~/.gitconfig` you will see all your global configurations.

But this is not enough. We will save even more precious keystrokes by defining some shorter Git aliases. I found an [excellent list](http://stevehodgkiss.com/posts/speed-up-your-git-workflow-with-bash-aliases) from [Steve Hodgkiss](http://stevehodgkiss.com/) and that is what I will use. Copy and paste it into your `~/.bash_aliases` (or equivalent).

    alias gl='git pull'
    alias gp='git push'
    alias gd='git diff'
    alias gc='git commit'
    alias gca='git commit -a'
    alias gco='git checkout'
    alias gb='git branch'
    alias gs='git status'
    alias grm="git status | grep deleted | awk '{print \$3}' | xargs git rm"

The vigilant one will notice that I stripped out his *git pull* and *git push* aliases. That is because you do not need them if your repository is properly configured. And if you do have them, you will not be able to configure another default remote branch.

Make certain you have these two INI sections your project's `.git/config` file.

    [remote "origin"]
    url = git@github.com:tobiassjosten/tobiassjosten.githubit
    fetch = +refs/heads/*:refs/remotes/origin/*
    [branch "master"]
    remotete = origin
    merge = refs/heads/master

That summarizes our Git tweaking for today. Maybe you have other tips to share though?
