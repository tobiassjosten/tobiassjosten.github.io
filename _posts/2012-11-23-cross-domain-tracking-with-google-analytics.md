---
layout: post
title: Cross domain tracking with Google Analytics
category: google-analytics
tags: [google-analytics]
summary: Yesterday when I was mentoring at Hyper Island I got a lot of questions on cross domain tracking with Google Analytics. Since I have not implemented this myself and because I dislike not knowing, I took some time reading up on the subject today.
---
Yesterday when I was [mentoring at Hyper Island](/events/mentoring-at-hyper-island/) I got a lot of questions on *cross domain tracking* with [Google Analytics](/google-analytics/). Since I have not implemented this myself and because I dislike not knowing, I took some time reading up on the subject today.

The [official documentation](https://developers.google.com/analytics/devguides/collection/gajs/gaTrackingSite) is good but I feel there is room for another explanation on how to implement the tracking. Here we go.

## Basic Analytics tracking

Let us start with this site, [`vvv.tobiassjosten.net`](http://vvv.tobiassjosten.net/). My basic tracking code could look something like this:

    var _gaq = _gaq || [];
    _gaq.push(['_setAccount', 'UA-1234567-8']);
    _gaq.push(['_trackPageview']);

    (function() {
        var ga = document.createElement('script');
        // etc…
    })();

The last bit, with the anonymous function, should always be present when implementing Analytics tracking and so I will skip it in the rest of my examples. The `_gaq` variable is what matters.

## Subdomain tracking

Now let us say I want to add a subdomain for a site where I can post pictures of cats; `cats.tobiassjosten.net`. Then I would have to add the following tracking code to *both `vvv.` and `cats.`*. Mind the initial dot in the domain name!

    var _gaq = _gaq || [];
    _gaq.push(['_setAccount', 'UA-1234567-8']);
    _gaq.push(['_setDomainName', '.tobiassjosten.net']);
    _gaq.push(['_trackPageview']);

If I was only tracking subdomains on the same domain, that would really be all I had to do. Done!

## Tracking separate top level domains

So now that my obsession with cat pictures has gotten the better of me, I want to balance it out with pictures of dogs and for that I plan to use `www.awesomedogpictures.com`.

I keep the tracking code on both `vvv.` and `cats.` and then add the following to my new dog site. Notice the changed domain name.

    var _gaq = _gaq || [];
    _gaq.push(['_setAccount', 'UA-1234567-8']);
    _gaq.push(['_setDomainName', '.awesomedogpictures.com']);
    _gaq.push(['_trackPageview']);

Once that is implemented I have to link the sites together. Because they are entirely different domains this requires both enabling incoming and outgoing linking.

For the incoming I add `_setAllowLinker` to all three sites, *before the call to `_trackPageview`*.

    _gaq.push(['_setAllowLinker', true]);

For the clickable links between my sites I need to modify the anchor tags and add an `onclick` attribute. Here is an example from `vvv.tobiassjosten.net` telling the world about my new dog site.

    <a href="http://www.awesomedogpictures.com/"
        onclick="_gaq.push(['_link', 'http://www.awesomedogpictures.com/']); return false;"
    >Here are some awesome dogs!</a>

You do not need this linking between sites sharing the same base domain, like `vvv.` and `cats.`, but only when the domains are completely different.
