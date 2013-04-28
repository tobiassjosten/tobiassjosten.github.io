---
layout: post
title: The missing Facebook trio
category: internet
tags: [internet, facebook]
summary: I am building an app that, among other things, lists your Facebook friends. All of a sudden I noticed my images were mismatched. That is not my Facebook friends! So I scanned the list from the top and noticed a face that was familiar – The Zuck, founder of Facebook. Obviously I had ran into a bug.
---
I am building an app that, among other things, lists your Facebook friends. Their images are being displayed from Facebook using the *http://graph.facebook.com/ID/picture* resource syntax. I am doing this from a loop construction, printing the friend's name and id as I go along.

All of a sudden I noticed my images were mismatched. That is not my Facebook friends! So I scanned the list from the top and noticed a face that was familiar – [The Zuck](http://www.facebook.com/zuck), founder of Facebook. Obviously I had ran into a bug.

It turned out I had typecasted my object to an array and had thereby lost my friend's IDs. Instead I was printing IDs from 1 and upwards. Not that notable of a bug to introduce in your code, half past one a saturday night. No, the interesting part is that Mr Zuckerberg was listed fourth in my list.

![User 1](http://graph.facebook.com/1/picture)
![User 2](http://graph.facebook.com/2/picture)
![User 3](http://graph.facebook.com/3/picture)
![User 4](http://graph.facebook.com/4/picture)

Anyone with a little knowledge of databases knows that sequential IDs (as is the case with Facebook) starts at 1. So why wasn't Mark Zuckerberg listed first? Who were those three before him?

Let us start by checking out [Mark in the Graph API](http://graph.facebook.com/4).

    {
       "id": "4",
       "name": "Mark Zuckerberg",
       "first_name": "Mark",
       "last_name": "Zuckerberg",
       "link": "http://www.facebook.com/zuck",
       "gender": "male",
       "locale": "en_US"
    }

Next we turn our gaze to [Anonymous 1](http://graph.facebook.com/1), [Anonymous 2](http://graph.facebook.com/2) and [Anonymous 3](http://graph.facebook.com/3).

    false

    false

    false

Empty! Three user accounts that does not exist (anymore). Who were they?
