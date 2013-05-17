---
layout: post
title: Symfony Icu requires lib-icu
category: symfony
tags: [symfony]
summary: Are you getting errors about missing icu libs when installing Symfony 2.3? It is easily fixed.
---
Are you trying to install [Symfony](/symfony/) version 2.3 but get the following error message?

>symfony/icu v1.2.0-RC1 requires lib-icu >=4.4 -> the requested linked library icu has the wrong version installed or is missing from your system, make sure to have the extension providing it.

This comes from a requirement in [the new Icu component](https://github.com/symfony/Icu) but it is easily fixed. Just install the `intl` extension for PHP.

    $ sudo aptitude install php5-intl

And then you are good to go!