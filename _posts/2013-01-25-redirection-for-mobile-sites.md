---
layout: post
title: Redirection for mobile sites
category: internet
tags: [internet]
summary: If you did not go mobile first when building your site, you will need to create a whole separate, mobile one. Here are some pitfalls to think about before doing so.
---
If you did not go *mobile first* when you initially built your website, chances are you will need to create a whole separate one if you want to cater to the mobile audience. Which you do.

So you develop a second, mobile website and it works great on your shiny new smartphone. But now you need to bring your mobile visitors over from the main site, to your new mobile site. Setting up redirection is easy but there are two common pitfalls you need to watch out for.

**Stateless redirection**. When I click a link to an interesting article, I obviously want to read it. Way too often redirection instead sends me to the homepage of the mobile site and this is both infuriating for the user and incredibly simple to avoid.

**Ambivalent detection**. Should I be sent to the mobile site when I am using a 10" tablet? How about a 11" laptop? Or an old Nokia 3210? There are [lists of mobile user agents](http://www.zytrax.com/tech/web/mobile_ids.html) but you still need to draw a line yourself.

These two are both known problems, solved by most self-respecting (yet non-responsive) websites of today. However, there is one pet peeve of mine that pretty much everyone still fails at.

## Bidirectional redirection

I click a link in my smartphone and your brilliant redirection mechanism takes me to the mobile site. Great! But I do not have time to read it now, so I save it in my read-it-later-list.

Later on I boot up my laptop, hit up my read-it-later-list and click that link again. This takes me to the mobile site and 99 times out of 100 it looks like shit on my computer. In some cases there is a "take me to the normal site" link and sometimes I can reconstruct the URL to try and find the proper page. But the damage is already done.

What you need is to make your redirection bidirectional. Just as you *redirect mobile clients to the mobile site*, you need to *redirect normal clients to the normal site*. Makes a lot of sense, does it not?

All these headaches really shows how efficient *responsive design* is.
