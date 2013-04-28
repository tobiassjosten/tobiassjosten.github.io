---
layout: post
title: Scrutinize your PHP
category: php
tags: [php]
summary: Scritinizer is a service for automated code reviews. PHP Analyzer, Code Sniffer, Mess Detector and more will make sure your PHP code is thoroughly scrutinized for sure.
---
Earlier today I was looking at [a pull request](https://github.com/FriendsOfSymfony/FOSUserBundle/pull/1047) I had sent to an open source project on GitHub. It turns out the patch is likely rejected but something interesting popped out. Just below my PR there was a system comment.

>Good to merge — Scrutinizer: No Comments — Travis: Passed

Now I know about [continous integration with Travis](http://vvv.tobiassjosten.net/symfony/continuous-integration-for-your-symfony2-app/) from before but what was that *Scrutinizer* business?

## Automated code reviews

Some Google searches later I found the [source](https://scrutinizer-ci.com/). Apparently it is the brain child of code machine [Johannes Schmitt](https://twitter.com/schmittjoh); known and greatly appreciated for all the awesome PHP libraries he has churned out over the years. The guy is a machine and it seems he keeps going.

Scritinizer is a service for automated code reviews. [A live example](https://scrutinizer-ci.com/g/sonata-project/SonataCacheBundle/inspections/d05d7a27-2d0b-4706-a6fa-b056595a05aa) shows us that this tool can take a code base like [SonataCacheBundle](https://github.com/sonata-project/SonataCacheBundle), find within it a code snippet like this:

    public function set(array $keys, $data, $ttl = 84600, array $contextualKeys = array())
    {
        $cacheElement = new CacheElement($keys, $data, $ttl);

        $result = apc_store(
            $this->computeCacheKeys($keys),
            $cacheElement,
            $cacheElement->getTtl()
        );

        return $cacheElement;
    }

… and automatically suggest how to improve it …

>The assignment to `$result` is dead and can be removed.

Quite cool, huh? But it gets better.

This dead code check is just one out of fourteen checks in the [PHP Analyzer](https://scrutinizer-ci.com/docs/tools/php/php-analyzer/) tool. Then there is [PHP Code Sniffer](https://scrutinizer-ci.com/docs/tools/php/code-sniffer/), [PHP Mess Detector](https://scrutinizer-ci.com/docs/tools/php/mess-detector/), [SensioLabs Security Checker](https://scrutinizer-ci.com/docs/tools/php/security-advisory-checker/), etc, etc. Your PHP code will be thoroughly scrutinized for sure.

## Scrutinize now

The service is entirely open and while I could not see any information about it; it seems it is still in public testing. Private repositories are supported and even they are free to use, thouhg I have no idea for how much longer.

So what are you waiting for? Go [take it for a ride](https://scrutinizer-ci.com/login)!
