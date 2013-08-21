---
layout: post
title: Symfony2 file permissions in Ubuntu
category: symfony
tags: [symfony, ubuntu]
summary: The Book has a section on setting up permissions for your project. It says to first try and use chmod because that is, it claims, how most systems controls their ACLs. This is only true for Mac OSX however. For Unix and Linux based ones, ACL is actually controlled using setfacl.
---
The installation section of [The Book](http://symfony.com/doc/current/book/) has a section on [setting up permissions](http://symfony.com/doc/current/book/installation.html). This walks through the process of setting up the `cache` and `log` directories so that both you and your webserver may modify their contents.

It says to first try and use [chmod](http://en.wikipedia.org/wiki/Chmod) because that is, it claims, how most systems controls their ACLs. This is only true for Mac OSX however. For Linux we use [setfacl](http://linuxcommand.org/man_pages/setfacl1.html) instead.

## Install and enable ACL

First you need to install ACL, preferrably through your package manager; `apt-get` on Ubuntu, `pacman` on Arch, `yum` on Fedora, etc:

    $ sudo apt-get install acl

Then you edit your `/etc/fstab` to enable ACL for your partition. Simply add `acl` to the list of options.

    /dev/sda1 / ext4 rw,auto,acl 0 1

Then lastly remount the partition to have the new options take effect.

    $ sudo mount -o remount /

## Linux file permissions

With ACL enabled for your partition, we can now solve our problem using three ingenious Linux tricks.

First we change the ownership of our directories, so that they are owned by our `www-data` group.

    $ sudo chown -R :www-data app/cache app/logs

Then we set a sticky guid on them. This ensures that new files and directories are automatically owned by the same group as their parent.

    $ sudo chmod g+s app/cache app/logs

Per default new files and directories are not writable by their group owner and so the last piece of our puzzle is to use the previously enabled ACL to change that.

    $ sudo setfacl -dR -m g::rwX app/cache app/logs

## Encrypted home directory

When I installed [Ubuntu](http://www.ubuntu.com/) I opted in to use the "encrypt home directory" feature. This makes use of eCryptFS and [it turns out](http://serverfault.com/questions/294158/enable-acl-for-ecryptfs-mounted-home-directory) this file system lacks support for ACL.

So what is a tin foil, Ubuntuist, Symfonian hatter to do?

The solution I used was to set up a `/symfony/project-name` directory, in which I created a `cache` and a `log` directory. Then I symlinked these to the project in my home directory.

    $ ln -s /symfony/nogfx/cache /home/tobias/projects/nogfx/app/cache
    $ ln -s /symfony/nogfx/log /home/tobias/projects/nogfx/app/log

Because they actually exist in my root partition I can easily enable ACL for them, while reaping the benefits of using [Symfony CLI](http://vvv.tobiassjosten.net/symfony/symfony2-cli-bash-script) on my encrypted home!
