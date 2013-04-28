---
layout: post
title: Symfony CLI helper script
category: symfony
tags: [symfony, cli, bash]
summary: I move around a lot when developing with Symfony. This is usually not a problem but with the Symfony CLI tool stuck in the root of projects, I felt it was a hassle. To solve this problem I wrote a little bash script which will give you a convenient way to access the CLI script from anywhere in your Symfony project.
---
I move around a lot when developing with Symfony. Well, not physically-physically but in my terminal I do. I move between the many directories when working with different aspects of my websites.

I am also using the Symfony CLI script a lot. Clearing the cache, generating stub classes, working with the database, etc. The problem is that the Symfony CLI script lives in the root of a Symfony project. Typing in '/home/tobias/projects/tobiassjosten/symfony cc' every time I want to clear the cache wont do. To help with this I wrote a little bash script which will give you a convenient way to access the CLI script from anywhere in your Symfony project.

The bash script works by traversing up in the filesystem tree, from where you are to the root (/), while looking for a 'symfony' executable. When it finds one it runs it for you with the parameters you specified.

This is a bash script rather than an .bashrc alias because I want access to it from all my users (tobias, root and www-data for example). With an alias you would either need to repeat yourself in multiple .bash_aliases or use a symlink strategy. And that is hardly adhering to [KISS](http://en.wikipedia.org/wiki/KISS_principle) or [DRY](http://en.wikipedia.org/wiki/Don't_repeat_yourself), is it?

## Awesome! How do I install it?

Installation is pretty straight forward. You [download the script file](http://bit.ly/aZ5xiW) to your $PATH and make it executable. That is it!

In terms of bash commands, on Linux, here is what you will need to do:

    $ sudo wget http://bit.ly/aZ5xiW -O /usr/local/bin/sf
    $ sudo chmod +x /usr/local/bin/sf

Now you can execute commands like 'sf cc' from anywhere in your Symfony project!

I also published the script as a [GitHub gist](http://gist.github.com/275690), so feel free to dig in if you think you can improve it.

**Update:** I how now extended the bash script to also work with [Symfony2 CLI](/symfony/symfony2-cli-bash-script).
