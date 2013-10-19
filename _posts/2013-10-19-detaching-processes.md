---
layout: post
title: Detaching processes
category: linux
tags: [linux, cli, tmux]
summary: What about when you forget to start up your terminal multiplexer before starting that long running process?
---
I run all my terminal sessions in a multiplexer, like [tmux](/tmux/). One feature I like with this tool, especially on remote servers, is the ability to easily detach long running processes, so you can log out without killing the process.

But what about when you forget to start up your multiplexer beforehand?

## Disown the process

Start by hitting `Ctrl-z` to send the SIGTSTP signal. This stops your process temporarily and gives you back your shell.

    $ sleep 100
    ^Z
    [1]  + 8398 suspended  sleep 100

Your shell can start multiple child processes this way. List them all using the `jobs` command. To target a specific one with any of the process commands, you can use `%<id>`.

    $ jobs
    [1]  + suspended  sleep 100
    $ kill %1
    [1]  + 8398 terminated  sleep 100

Next send it to the background using `bg`. This will resume the process from its temporary stop and let it continue working.

    $ bg
    [1]  + 8398 continued  sleep 100

Once in the background we can `disown` it. Problem solved!

## Reroute its output

Well, not entirely solved. The process still uses your TTY for its STDOUT, so any output of its will be printed to your terminal. This can be quite annoying but is definitely solvable!

You need to know its process id for this next step. In previous examples this was 8398 but you can use `ps` to find yours. Once you have the pid, attach to the process with `gdb` and reroute STDOUT/STDERR.

    $ sudo dgb -p 8398
    … spam from gdb …
    (gdb) p dup2(open("/dev/null",0),1)
    $1 = 1
    (gdb) p dup2(open("/dev/null",0),2)
    $2 = 2
    (gdb) detach
    Detaching from program: /bin/zsh5, process 8398
    (gdb) quit

There, the process now runs in the background, disowned from your terminal and with its output redirected to /dev/null. Now all you need to do is find a way not to forget starting your multiplexer first next time.
