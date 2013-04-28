---
layout: post
title: Language prefix and Google Analytics
category: business-intelligence
tags: [business-intelligence, google-analytics]
summary: With an out-of-the-box tracking script for Google Analytics you will record every page on your site as is. This means that the English version of your "About" page could be reported as /en/about and the Swedish one as /sv/about. That setup makes it difficult to analyze your traffic.
---
With an out-of-the-box tracking script for Google Analytics you will record every page on your site as is. This means that the English version of your "About" page could be reported as /en/about and the Swedish one as /sv/about.

That setup makes it difficult to analyze your traffic. Both when comparing languages for a specific page but also when comparing pages with eachother.

The solution is a two part one.

## Virtual paths

First you need to tell Google Analytics what path you want tracked for each page. This is done by using the `_pageUrl` parameter with the [`_trackPageView` method](http://code.google.com/apis/analytics/docs/gaJS/gaJSApiBasicConfiguration.html#_gat.GA_Tracker_._trackPageview).

    _gaq.push(['_setAccount', 'UA-12345-1']);
    _gaq.push(['_trackPageview', '/about']);

This rids you of the language prefix in your reports.

## Custom variables

Language should still be tracked however. Now that we have it extracted from our paths we can instead use a better suited channel for the data -  [custom variables](http://code.google.com/apis/analytics/docs/tracking/gaTrackingCustomVariables.html).

The `_setCustomVar` method takes up to four paramters; *index*, *name', *value* and *scope*. If this is the first time you are using *custom variables* then you I suggest you pick 1 for index, Language for name, your actual language for value and 3 for scope.

    _gaq.push(['_setAccount', 'UA-12345-1']);
    _gaq.push(['_setCustomVar', 1, 'Language', 'en', 3]);
    _gaq.push(['_trackPageview', '/about']);

Custom variables can then be used to create an advanced segment in Google Analytics. These are what you will use to compare languages.

## Tracking Events

If you are [tracking events with Google Analytics](/symfony/tracking-google-analytics-events-with-symfony) you might want to take care when combining it with the above techniques.

One potential pitfall is that setting a custom variable attaches it to the tracker object. If you are using the same object to track an event later on you might, or might not, want to remove the variable.

    _gaq.push(['_deleteCustomVar', 1]);

In the case of language tracking it could be good to keep gathering that data but it all depends on your use case.

One other problem is that event tracking uses the _current actual path_ as given by the browser. This means that events are not tracked, and cannot be tracked, for your virtual paths. If you ask me, this is a bug in Google Analytics but I have not found any information about it.

Just keep that in mind though and you could improve your analytics experience using these techniques. Happy tracking!
