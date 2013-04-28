---
layout: post
title: Keeping Google juice with HTTP 503
category: internet
tags: [internet, seo]
summary: My Drupal consultance firm recently re-launched our website. Our time schedule was pressed, to say the least, so we decided to not migrate the blog just yet. In order to not lose out on our "Google juice" we decided to set up some HTTP magic to keep the crawlers at bay.
---
My [Drupal consultance firm](http://www.kollegorna.se/) recently re-launched our website. We did this now because we are running a minor campaign for a select customer segment. Since we hate throwing away work, and because we want to do more with the site later on, we decided to rebuild the site on Drupal 7.

Our time schedule was pressed, to say the least. That is why we excluded our blog section from the new site. We will re-implement it later but for now we just wanted to launch.

But we also did not want to lose the *Google juice* for our blog posts, so their URLs needed handling. We did so using the *HTTP 503* (temporarily unavailble) response. Two Apache rewrite lines later and we were done.

    RewriteCond %{REQUEST_URI} ^/blogg(/.*)? [NC]
    RewriteRule .* pleasecomeagain [L,R=503]

If you want to do the same but have more than one URL segment you want to put on hold, here is how you could use the OR operator.

    RewriteCond %{REQUEST_URI} ^/thingy1(/.*)? [NC,OR]
    RewriteCond %{REQUEST_URI} ^/thingy2(/.*)? [NC]
    RewriteRule .* pleasecomeagain [L,R=503]

Presto! When Google receieves the 503 response it knows that the requested page is still alive and kicking, but it can not be fetched right now. So it tries again later and thus our page remains indexed.

I am not sure for how long this is a viable reason but that will hopefully only spur us to complete the blog section.
