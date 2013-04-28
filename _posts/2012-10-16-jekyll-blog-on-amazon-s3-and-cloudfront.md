---
layout: post
title: Jekyll blog on Amazon S3 and CloudFront
category: development
tags: [development, jekyll, amazon, aws]
summary: Last weekend I decided to migrate this blog away from my VPS and while doing so I wanted to see how much performance I could squeeze out of it.
---
Last weekend I decided to migrate this blog away from my VPS and while doing so I wanted to see how much performance I could squeeze out of it.

I have long been interested in checking out the [Amazon Simple Storage Service](http://aws.amazon.com/s3/), as it seems like a really nice way to host a static site.

Sure, I could have used [GitHub Pages](http://pages.github.com/), with built-in Jekyll support, but it has two disadvantages which made it a non-option for me:

- You can not use custom Jekyll plugins with GitHub Pages. Your site is generated with `--safe`.

- There is no possibility to redirect requests. This is vital if you do not want to lose any link juice when fixing a typo in your URL.

## Jekyll on Amazon S3

In order to serve your site from S3 you need to first create a bucket to hold it. Using the [S3 web console](https://console.aws.amazon.com/s3/home). you click to edit the properties of your new bucket, go to the Website tab and check *Enabled*. Note the *Endpoint* of your site and remember to save!

I chose `index.html` for my index document and `404.html` for my error document but your setup might vary. Amazon uses this configuration to map requests for `/whatever/` to `/whatever/index.html`.

You could click *Upload* next and push your site up on S3. But being a command line junkie I wanted to improve this process by finding a way to deploy in my terminal.

It seems [s3cmd](http://s3tools.org/s3cmd) is the answer to that. I really recommend pulling down the [code from GitHub](https://github.com/s3tools/s3cmd) and installing s3cmd from the sources, or you will miss a few crucial features.

Once installed you first configure it:

    s3cmd --configure

And then deploy:

    s3cmd sync _site/ s3://vvv.tobiassjosten.net/

Now you should be able to visit your site using the endpoint you noted in the web console. Congratulations, you are hosted on S3!

## CloudFront for über speed

My site was migrated and now it was time for the performance tweaking.

Because it is an entirely static site, each request is basically served a predetermined response. I generate the site once with Jekyll and the result can then be served as-is. This enables me to leverage the power of *content distribution networks*.

Normally you use CDNs for static assets like external JavaScript, CSS and images. The idea is that you put copies of these files on servers all around the world and through some networking magic you have each request routed to its closest server.

One such CDN is a cousin of S3, namely Amazon CloudFront.

Setting it up to serve your S3 hosted Jekyll site is dead simple. You create a new distribution channel and enter the S3 endpoint as its origin. After a couple of minutes your distribution channel will be ready and you can visit you speedy new site at the domain name listed for it.

One pitfall I happened into myself was that I picked the S3 bucket as my origin. You should use the *S3 endpoint*, as noted above, or else you will see an error like this:

    <Error>
      <Code>AccessDenied</Code>
      <Message>Access Denied</Message>
      <RequestId>47E209BD83AC0CD2</RequestId>
      <HostId>
        aMtPYlc+1n9Loa0/fjfBmOYOZQbdDL8xa+S8m+lXw5XqIIUMz80HMU3f4rxhipTd
      </HostId>
    </Error>

Once you are ready, set a CNAME pointer for your subdomain to point at the CloudFront endpoint. Since that will not work for bare domains you can use [WWWizer](http://wwwizer.com/) — just have an A pointer directed at `174.129.25.170` and your example.com will be redirected to www.example.com, [as should be](/internet/using-www-for-your-domain/).

## Deployment and invalidation

There are [more options for s3cmd](http://manpages.ubuntu.com/manpages/precise/en/man1/s3cmd.1.html) that you might want to use when deploying your site. For example you will probably want to delete all the files on S3 that you no longer have locally:

    s3cmd --delete-removed _site/ s3://vvv.tobiassjosten.net/

You can invalidate the CloudFront cache for the files you upload:

    s3cmd --cf-invalidate _site/ s3://vvv.tobiassjosten.net/

Create a new behavior for your CloudFront distribution channel, matching .js and .css files, and have it forward query strings. Then you can add a version to your JavaScript and CSS files to force clients to download new ones, so that you can slap an *Expires* header on those files:

    s3cmd --add-header='Expires: Sat, 20 Nov 2286 18:46:39 GMT' \
        _site/ s3://vvv.tobiassjosten.net/ \
        --exclude '*.*' --include '*.js' --include '*.css'

If you compress your files with *gzip* before uploading them, you can have them served as such:

    s3cmd --add-header='Content-Encoding: gzip' \
        _site/ s3://vvv.tobiassjosten.net/

Small steps towards making your administration easier and your site faster.

## My blog is the 1%

So how what was the results of my performance work? Running the [Full Page Test from Pingdom](http://tools.pingdom.com/fpt/#!/LgevdxNw5/http://vvv.tobiassjosten.net/) I learned that:

>Your website is faster than 99% of all tested websites

The initial HTML page takes an insane 9ms to load! Nine thousandths of a second! My perspective on speed will never be the same again.
