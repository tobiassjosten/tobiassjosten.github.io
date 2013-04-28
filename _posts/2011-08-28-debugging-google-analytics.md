---
layout: post
title: Debugging Google Analytics
category: business-intelligence
tags: [business-intelligence, google-analytics]
summary: As much as I love working with Google Analytics, developing for it can sometimes be quite a hassle. Basically it consists of defining your data, calling the ga.js script and then hoping it takes care of business. But does it really and how can you tell?
---
As much as I love working with Google Analytics, developing for it can sometimes be quite a hassle. Basically it consists of defining your data, calling the ga.js script and then hoping it takes care of business. But does it really and how can you tell?

## Debug tracker

One easy way to see what the Analytics tracker is doing is switching to the debugging tracker. Simply replace `/ga.js` with `/u/ga_debug.js` in your code and you are done. This will make the tracker print its actions to your console, so you will need to bring that up before reloading the page.

If you are running Chrome you can install the [Google Analytics Debugger](https://chrome.google.com/webstore/detail/jnkmfdileelhofjcijamephohjechhna) extension to get the very same information without changing your code. Perfect if you want to debug in production.

## Network sniffer

The way tracking works with Google Analytics is it requests a [GIF](http://en.wikipedia.org/wiki/Graphics_Interchange_Format) image, `__utm.gif`, with your data in the request. This means we can easily spot each tracking action using a network sniffer.

I have previously written about [debugging with Charles proxy](/development/monitor-and-debug-with-charles-proxy) and this is also a great tool for working with Analytics. Fire it up, have it filter for `__utm.gif` and you will see every tracking request made.

One word of warning though. If you are working with a site using SSL, you will probably see the following error:

    "SSLHandshake: Received fatal alert: unknown_ca"

This means you must enable SSL in Charles and have your browser accept its certificate. Following [the official guide](http://www.charlesproxy.com/documentation/using-charles/ssl-certificates/) to do this is easy.

## Common pitfalls

Double check your types. If the documentation says to use an integer, do not send a float. And 'true' is not a boolean, it is a string. Should you have the type wrong then everything will look fine in the tracking, except the data wont actually get registered. If you are unsure about what data you have, [type cast](http://jibbering.com/faq/notes/type-conversion/) it to be certain.

Know which actions sends requests and which are only modifiers. [Custom variables](/business-intelligence/language-prefix-and-google-analytics), for example, only adds parameters to the next tracking request but you will still need to run [`_trackPageview`](http://code.google.com/apis/analytics/docs/gaJS/gaJSApiBasicConfiguration.html#_gat.GA_Tracker_._trackPageview) for the page.

It takes time for data to reach your reports. Most commonly this is a matter of a day before your page views will show up but for custom variables [it can take](https://www.google.com/support/forum/p/Google+Analytics/thread?tid=210c60def4c05e19) two days. First make sure your sending the right data when tracking and then verify later that it actually worked as intended.

The data Analytics reports is not 100% accurate. Do not think otherwise.

Move your tracking to the closing head tag of your site, even though Google [says otherwise](http://www.google.com/support/analytics/bin/answer.py?hl=en&answer=55574). Their reasons [are moot](http://stackoverflow.com/questions/3694356/does-there-exist-a-reason-to-put-google-analytics-in-head-and-not-in-body/3695553#3695553) with the asynchronous tracker. By moving it earlier in the page you will catch more visitors that might otherwise click away before being tracked.
