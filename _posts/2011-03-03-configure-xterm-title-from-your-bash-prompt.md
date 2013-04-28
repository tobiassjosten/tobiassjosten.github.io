---
layout: post
title: Configure xterm title from your Bash prompt
category: bash
tags: [bash]
summary: Many developers, me included, live out most of their lives in the terminal. But what about the rest of the time? Most likely your terminal stays minimized or otherwised hidden in an activity bar. If you are anything like me you will have lots and lots of such tabs in the activity bar. So how do we find our way back to the right terminal?
notice: This is a follow-up to my articles on adding [Git branch to your prompt](/git/add-current-git-branch-to-your-bash-prompt) and creating a [dynamic bash prompt](/bash/dynamic-prompt-with-git-and-ansi-colors).
---
Many developers, myself included, live out most of their lives in the terminal. But what about the rest of the time? Most likely your terminal stays minimized or otherwised hidden in an activity bar. If you are anything like me you will have lots and lots of such tabs in the activity bar. So how do we find our way back to the right terminal?

## Xterm control code

If you are on a *nix based system, you are most likely using xterm. And if you are using xterm you can easily modify your terminals title.

Like most terminal based negotiation configuration, there is a command code for starting and stopping this title sequence. It goes like this.

    PS1="\e]2;My title\a>"

Put that into your `~/.bashrc` and you will be given a > for prompt and a lovely "My title" for title. Adding dynamic data to the title is done just like how you do it for the prompt itself.

    PS1="\e]2;I am \u\a>"

## Final bash prompt

Let us merge this with the [magnificent bash prompt](http://vvv.tobiassjosten.net/bash/dynamic-prompt-with-git-and-ansi-colors) we built before. Do not forget to enclose the non-printing characters!

    PROMPT_COMMAND='PS1="\[\e]2;\u@\h:\w\a\]${c_user}\u${c_reset}@${c_user}\h${c_reset}:${c_path}\w${c_reset}$(git_prompt)\$ "'

Be sure to check out IMB's excellent [introduction to prompt customization](http://www.ibm.com/developerworks/linux/library/l-tip-prompt/)!
