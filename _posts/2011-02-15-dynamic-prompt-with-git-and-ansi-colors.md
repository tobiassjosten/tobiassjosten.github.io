---
layout: post
title: Dynamic prompt with Git and ANSI colors
category: bash
tags: [bash, bash]
summary: I recently tinkered (again) to put my current Git working branch in my prompt. It resulted in very useful addition to the tool I use a lot in my everyday work. However it turned out my implementation had a bug.
---
I recently tinkered (again) to put my current [Git working branch in my prompt](/git/add-current-git-branch-to-your-bash-prompt). It resulted in very useful addition to the tool I use a lot in my everyday work.

However it turned out my implementation had a bug. When my prompt plus command got near the window border it wrapped the text. It should wrap but in my case it did so too early and not on a new line but with the overflowing text replacing the beginning of my prompt.

I have had this problem in the past when tinkering with my prompt. This time around I was not going to throw in the towel however and so I went to read up on terminals.

The problem, I found out, was that I was not [enclosing my non-printing characters](http://tldp.org/HOWTO/Bash-Prompt-HOWTO/nonprintingchars.html). When I changed my code accordingly it handled the wrapping perfectly.

    c_reset='\[\e[0m\]'
    c_user='\[\e[1;33m\]'
    c_path='\[\e[0;33m\]'
    c_git_clean='\[\e[0;36m\]'
    c_git_dirty='\[\e[0;35m\]'

    PS1="${c_user}\u${c_reset}@${c_user}\h${c_reset}:${c_path}\w${c_reset}$(git_prompt)\$ "

Relieved having conquered this tiny area of Bash I got back to work. Then when next I committed my code and was expecting the dirty branch color to switch to the clean branch color, it did not. I fired up a new terminal and lo and behold – it had the correct color for my branch.

Some experimenting later it became apparent that I could either use double quotes ("like this") to have my string parsed and the ANSI colors properly applied, or I could use single quotes ('like this') to have the branch value updated dynamically.

A dive in the bash manual (RTFM, right?) showed me a solution to the problem.

    PROMPT_COMMAND
        If set, the value is executed as a command prior to issuing each primary prompt.

Instead of setting PS1 directly, I set up PROMPT_COMMAND to have it set PS1 in turn. Then I could both have the colors from double quotes parsing and the dynamic value from calling my function on each new prompt.

The final result?

    # Configure colors, if available.
    if [ -x /usr/bin/tput ] && tput setaf 1 >&/dev/null; then
      c_reset='\[\e[0m\]'
      c_user='\[\e[1;33m\]'
      c_path='\[\e[0;33m\]'
      c_git_cleancleann='\[\e[0;36m\]'
      c_git_dirty='\[\e[0;35m\]'
    else
      c_reset=
      c_user=
      c_git_cleancleann_path=
      c_git_clean=
      c_git_dirty=
    fi

    # Function to assemble the Git parsingart of our prompt.
    git_prompt ()
    {
      if ! git rev-parse --git-dir > /dev/null 2>&1; then
        return 0
      fi

      git_branch=$(git branch 2>/dev/null | sed -n '/^\*/s/^\* //p')

      if git diff --quiet 2>/dev/null >&2; then
        git_color="$c_git_clean"
      else
        git_color="$c_git_dirty"
      fi

      echo "[$git_color$git_branch${c_reset}]"
    }

    # Thy holy prompt.
    PROMPT_COMMAND='PS1="${c_user}\u${c_reset}@${c_user}\h${c_reset}:${c_path}\w${c_reset}$(git_prompt)\$ "'

The one thing I would like to change now is how I am detecting ANSI color support. Since tput went out the window, I am not sure that is a proper indicator anymore. Hoping to remedy this, I have put a [Gist for the prompt](https://gist.github.com/828432) up at GitHub. Feel free to fork and improve!
