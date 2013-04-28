---
layout: post
title: tree - structured files and directories listing
category: linux
tags: [linux, cli]
summary: The tree command is one of those tools that makes our CLI supreme to its GUI interface. Its use is to list files and directories in a structured manner. I find it gives me an excellent overview of the directory structure and I use it a lot to familiarize myself with new projects.
---
The `tree` command is one of those tools that makes the CLI supreme to its GUI interface. Its use is to list files and directories *in a structured manner*. I find it gives me an excellent overview of the directory structure and I use it a lot to familiarize myself with new projects.

The base `tree` command will *list all files and directories*, below your current working directory.

    $ tree
    .
    ├── httpare.cabal
    ├── LICENSE
    ├── README.md
    ├── Setup.hs
    └── src
        ├── Httpare
        │   └── Class.hs
        └── Main.hs
    
    2 directories, 6 files

If that is too many to see, have `tree` just show the directory structure.

    $ tree -d
    .
    └── src
        └── Httpare
    
    2 directories

There is a useful level parameter you can use to have `tree` now descend too many levels. I employ that a lot in projects with massive amounts of files, like web sites.

    $ tree -L 2
    .
    ├── httpare.cabal
    ├── LICENSE
    ├── README.md
    ├── Setup.hs
    └── src
        ├── Httpare
        └── Main.hs
    
    2 directories, 5 files

You can also tell `tree` to exclude certain files by giving it an ignore parameter.

    $ tree -I '*.hs'
    .
    ├── httpare.cabal
    ├── LICENSE
    ├── README.md
    └── src
        └── Httpare
    
    2 directories, 3 files

As with most *nix tools, `tree` is very flexible. You can for example have it simulate `find` by throwing in parameters to skip indentation, show all files, print the full paths, skip colors and the summary report.

    $ tree -afin --noreport
    .
    ./.gitignore
    ./httpare.cabal
    ./LICENSE
    ./README.md
    ./Setup.hs
    ./src
    ./src/Httpare
    ./src/Httpare/Class.hs
    ./src/Main.hs
    ./.travis.yml

## Installation

In Ubuntu you can install `tree` with `aptitude`.

    $ sudo aptitude install tree

If you are on Mac OSX you should use `brew` to install your *nix software.

    $ sudo brew install tree
