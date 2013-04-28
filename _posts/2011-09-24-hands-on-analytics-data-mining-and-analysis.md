---
layout: post
title: Hands-on Analytics data mining and analysis
category: business-intelligence
tags: [business-intelligence, google-analytics]
summary: I am a big advocate for data driven decisions and I believe that, as a business, you should try to gather as much data as you can. You never know in advance what discoveries might jump out and surprise you.
---
I am a big advocate for data driven decisions and I believe that, as a business, you should try to gather as much data as you can. You never know in advance what discoveries might jump out and surprise you.

So when I started building [Smartburk](http://www.smartburk.se/) (swedish) I wanted to measure the potential interest in the service. I did so by putting up a splash page explaining the benefits of using Smartburk, accompanied by a signup form. This gives me a list of interested users I can communicate with for feedback and that is absolutely invaluable.

For every signup I also trigger a [Google Analytics event](http://vvv.tobiassjosten.net/symfony/tracking-google-analytics-events-with-symfony). This allowed me to set up a goal within Analytics, to measure the conversion rate for all traffic. By cross referencing this with traffic sources it turns out that Twitter drives the best conversion rates by far and so I might already have some grounds for future campaigns.

My next step was to segment this interest to see what kind of series that would be most interesting for Smartburk to target first.

I used [IMDb's list](http://www.imdb.com/search/title?num_votes=5000,&sort=user_rating,desc&title_type=tv_series) for the best TV series, plus a few of my own personal favourites and dislikes, which I then pulled out the number of Likes from swedish users for each and every one of these shows. I now had a list of close to 300 series and their level of engagement in Sweden.

Then I split this list up in eight categories and picked a couple of top series in each category. These I created individual ads for on Facebook, targeting users who had liked the series. I used the same copy for all ads but changed the picture to something from the series. All links were tagged with [Analytics campaign](http://www.google.com/support/analytics/bin/answer.py?hl=en&answer=55540) parameters.

These ads have now been running for a week on Facebook, while I have been actually building Smartburk, and the results are in. Analytics lets you browse your campaigns and add a secondary dimension to the list ('content' for specific series). With the goal created above I can also see the conversion rate for each individual series.

With very little work on my part, except for quite some initial research on the series, I now know which series are the most important ones to support for Smartburk. I have quality data on which series drives the most traffic and which users are interested in Smartburk. Most of it are in line with what I could guess but there has already been some odd results I could never anticipate.

I will continue to measure and gather as much data as I can. But I am also interested in hearing - what are you measuring?
