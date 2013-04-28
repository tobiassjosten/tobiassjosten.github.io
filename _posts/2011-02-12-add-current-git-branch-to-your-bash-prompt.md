---
layout: post
title: Add current Git branch to your Bash prompt
category: git
tags: [git, bash]
summary: Following up on my post about Git aliases, I want to share another convenience tool for using Git. This time we will talk prompt enhancement – the terminal alternative to plastic surgery.
---
Following up on my post about [Git aliases](/git/configure-git-aliases-to-save-the-day), I want to share another convenience tool for using Git. This time we will talk prompt enhancement – the terminal alternative to plastic surgery.

By then end of this post, you too can have a beautiful prompt like this!

    tobias@laptop:~/projects/smartburk [graphapi]$

I am going to assume that you are either using [Bash](http://www.gnu.org/software/bash/) or have enough knowledge to adjust these instructions to whatever shell you are on.

The first thing we need to add to our `~/.bashrc` is color configuration. We will do so using the `tput` program, which lets us fetch dynamic (environment specific) terminal data.

    if [ -x /usr/bin/tput ] && tput setaf 1 >&/dev/null; then
      c_reset=`tput sgr0`
      c_user=`tput setaf 2; tput bold`
      c_path=`tput setaf 4; tput bold`
      c_git_clean=`tput setaf 2`
      c_git_dirty=`tput setaf 1`
    else
      c_reset=
      c_user=
      c_path=
      c_git_cleanclean=
      c_git_dirty=
    fi

So now we have five variables with different colors, if our terminal supports it. Next up is defining the function that will poll Git for our current branch and color red or green it depending on whether our working directory is clean or not.

    git_prompt ()
    {
      if ! git rev-parse --git-dir > /dev/null 2>&1; then
        return 0
      fi

      git_branch=$(git branch 2>/dev/null| sed -n '/^\*/s/^\* //p')

      if git diff --quiet 2>/dev/null >&2; then
        git_color="${c_git_clean}"
      else
        git_color=${c_git_cleanit_dirty}
      fi

      echo " [$git_color$git_branch${c_reset}]"
    }

This Furthernction first checks to see whether we are in a git repository or not. The `git rev-parse` command terminates with an error exit code if not and we use that to crash and die early if there are no Git data.

Then it calls `git branch` and singles out the one you are at using the `sed` tool. This also works even if you are not on a branch, because then you are given "* (no branch)" in the branch list.

Finally, it runs `git diff` and works off of its exit code to determine if we have a working directory or not. Then we know what color to use for our branch.

By using exit codes like this, we speed up our data harvesting a lot. Running `git status` or `git diff` normally can take some time and we can not have that in our prompt.

Now that we have all our data we can put together our prompt.

    PS1='${c_user}\u${c_reset}@${c_user}\h${c_reset}:${c_path}\w${c_reset}$(git_prompt)\$ '

There we go. Aint that a beauty?

Maybe you have something better or just completely different? Please share!

**Update:** There is a problem with wrapping text in the terminal using this prompt. See my post on [Dynamic prompt with Git and ANSI colors](/bash/dynamic-prompt-with-git-and-ansi-colors) an explanation and solution to this problem.
