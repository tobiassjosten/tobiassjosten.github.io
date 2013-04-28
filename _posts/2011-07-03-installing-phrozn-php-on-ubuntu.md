---
layout: post
title: Installing Phrozn PHP on Ubuntu
category: php
tags: [php, ubuntu]
summary: I was recently inspired by an article written by Christian Schaefer, on the Phrozn project. A very impressive PHP based static site generator!
---
I was recently inspired by [an article](http://test.ical.ly/2011/07/01/phrozn-a-static-site-generator-written-in-php-with-the-help-of-zend-framework-2-symfony2-and-twig/) written by [Christian Schaefer](https://twitter.com/caefer), on the [Phrozn project](http://www.phrozn.info/). A very impressive PHP based static site generator!

The project mentions a few sources of inspiration, like [Jekyll](http://jekyllrb.com/) (that powers [GitHub pages](http://pages.github.com/), [Hyde](http://ringce.com/hyde) and a [bunch of others](http://www.phrozn.info/en/#alternatives). I wont go into details about them or even Phrozn but I do recommend that you [check it out](http://www.phrozn.info/en/documentation/articles/getting-started/) and read through the [many blog posts](http://www.google.com/search?q=phrozn&tbm=blg) on it.

Instead I will assume you are convinced it is an interesting system and want to install it on your Ubuntu machine.

The recommended way is using Phrozn's PEAR channel:

    $ pear channel-discover pear.phrozn.info
    $ pear install phrozn/Phrozn-beta

Because your PEAR setup can differ from mine, you need to figure out where Phrozn was installed to. That is done by checking out `bin_dir` and the `php_dir`.

    $ pear config-show|grep bin_dir
    PEAR executables directory     bin_dir          /home/tobias/pear/bin
    $ ls /home/tobias/pear/bin/
    dbunit  pear  peardev  pecl  phpcov  phptok  phpunit  phr  phrozn

Of course it is way too troublesome having to execute `/home/tobias/pear/bin/phr` command every time you want to invoke Phrozn. To be able to run `phr` straight from the prompt requires that it is in your $PATH. This might have been handled for you when you installed PEAR but if not, simply add the following to your .bashrc:

    # Include PEAR executables in PATH
    if [ -x pear ]; then
        PATH="$PATH:`pear config-show|grep bin_dir|awk '{ print $5 }'`"
    fi

Actually running Phrozn is where I bumped into problem.

    $ phr
    PHP Fatal error:  Class 'Console_CommandLine' not found in /home/tobias/projects/phrtest/Phrozn/Runner/CommandLine/Parser.php on line 39
    PHP Stack trace:
    PHP   1. {main}() /home/tobias/pear/bin/phr:0
    PHP   2. Phrozn\Runner\CommandLine->run() /home/tobias/pear/bin/phr:15
    PHP   3. Zend\Loader\StandardAutoloader->autoload() /home/tobias/projects/phrtest/Phrozn/Vendor/Zend/Loader/StandardAutoloader.php:0
    PHP   4. Zend\Loader\StandardAutoloader->loadClass() /home/tobias/projects/phrtest/Phrozn/Vendor/Zend/Loader/StandardAutoloader.php:224
    PHP   5. include() /home/tobias/projects/phrtest/Phrozn/Vendor/Zend/Loader/StandardAutoloader.php:304

The error message I got said that it tried but failed to include [Console_CommandLine](http://pear.php.net/package/Console_CommandLine). So I decided to try and install it:

    $ pear install Console_CommandLine
    pear/Console_CommandLine is already installed and is the same as the released version 1.1.3
    install failed

So it was installed but could not be loaded – this was obviously another problem with paths not being set up. Again looking at `pear config-show` I gathered that my `php_dir` was configured to `/home/tobias/pear/share/pear`. After checking my PHP's `include_path` I added an override:

    $ echo -e '[PHP]\ninclude_path = ".:/usr/share/php:/usr/share/pear:/home/tobias/pear/share/pear"' | sudo tee /etc/php5/conf.d/pear_include.ini

Then I tried again:

    $ phr -v
    /usr/bin/phrozn version 0.2.

Great success! Everything is in place and code can start being written.
