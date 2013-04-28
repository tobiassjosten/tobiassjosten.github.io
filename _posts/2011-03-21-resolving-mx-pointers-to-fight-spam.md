---
layout: post
title: Resolving MX pointers to fight spam
category: php
tags: [php, spam]
summary: As some of you might have noticed (and some even pointed out) I have had some trouble with spam lately. After dismissing the notion to disable comments altogether, I set out to find a solution to fight the spam.
notice: This is part of an *[anti-spam](/spam) series*, including [Stopping spam with a tasty honeypot](/internet/stopping-spam-with-a-tasty-honeypot) and [Stopping spam with Symfony forms](/symfony/stopping-spam-with-symfony-forms).
---
As some of you might have noticed (and some even [pointed out](http://twitter.com/masviken/statuses/49809006071529472)) I have had some trouble with spam lately. After dismissing the notion to disable comments altogether, I set out to find a solution to fight the spam.

First I checked out the [Mollom](http://mollom.com/) and [Akismet](http://akismet.com/) *anti-spam services*. They are both great at what they do and if you are in the same position I recommend looking them up. It was more hassle than I was willing to deal with however and so I started thinking about a simpler solution.

I currently require three fields to be filled out when commenting. Name, email and the actual message. Mollom and Akismet both target the message part, by analyzing its content. The name alone is not that much data to make a decision on, so we are left with the email address.

What makes an email address valid? The obvious answer is that it has a recipient in the other end. However, since we are looking for an easy solution we do not want to go about connecting to the SMTP server to see if our given user is valid there. Especially since not all SMTP servers will give up this information freely.

So what does a lazy developer do? We use cheap tricks!

## Resolving the hostname

I took a first stab at this anti-spam approach by using [`gethostbyname()`](http://php.net/gethostbyname). It does exactly what I wanted – resolves the hostname to an IP if there are proper A pointers configured. If not, it returns the unresolved hostname.

After a couple of minutes hacking I had a validator for the email widget. It tried resolving the hostname and raised an error if that did not yield an IP address. I tried a couple of the email addresses that had previously spammed me and it worked like a charm.

Then I deployed my solution to the production server and went about testing it there as well, just to make sure before I called it a victory. But it did not work. Whatever I threw at it passed through.

I picked a domain from a spam email address and tried running it from PHP CLI.

    $ php -r 'echo gethostbyname("qeilfcyw.com")."\n";'
    109.74.195.184
    $ php -r 'echo gethostbyname("qeilfcyw.com")."\n";'
    97.107.142.101

It looked up different IP addresses each time! Very odd. So I tried running dig to see what pointers it had configured.

    $ dig qeilfcyw.com

    ; <<>> DiG 9.7.0-P1 <<>> qeilfcyw.com
    ;; global options: +cmd
    ;; Got answer:
    ;; ->>HEADER<<- opcode: QUERY, status: NXDOMAIN, id: 12590
    ;; flags: qr rd ra; QUERY: 1, ANSWER: 0, AUTHORITY: 1, ADDITIONAL: 0

    ;; QUESTION SECTION:
    ;qeilfcyw.com.                  IN      A

    ;; AUTHORITY SECTION:
    com.                    411     IN      SOA     a.gtld-servers.net. nstld.verisign-grs.com.
    1300285512 1800 900 604800 86400

    ;; Query time: 0 msec
    ;; SERVER: 79.99.4.100#53(79.99.4.100)
    ;; WHEN: Wed Mar 16 14:34:06 2011
    ;; MSG SIZE  rcvd: 103

Nothing. Could it be that my resolv.conf had a search directive?

    $ cat /etc/resolv.conf 
    nameserver 79.99.4.100
    nameserver 79.99.4.101

Nope, no luck there neither. Maybe a traceroute could help find out where the traffic goes (and why)?

    $ traceroute qeilfcyw.com
    traceroute to qeilfcyw.com (97.107.142.101), 30 hops max, 60 byte packets
     1  109-74-0-112-static.serverhotell.net (109.74.0.112)  0.053 ms
    0.032 ms  0.029 ms
     2  79-99-4-2.serverhotell.net (79.99.4.2)  1.362 ms  1.336 ms  1.332 ms
     3  195-20-206-2.serverhotell.net (195.20.206.2)  9.789 ms  9.769 ms  9.746 ms
     4  te7-2.228.ccr01.sto01.atlas.cogentco.com (149.6.168.109)  10.953
    ms  10.991 ms  11.143 ms
     5  te0-3-0-5.ccr21.ham01.atlas.cogentco.com (154.54.38.29)  29.288 ms
     29.416 ms te0-2-0-5.ccr21.ham01.atlas.cogentco.com (154.54.38.25)
    29.129 ms
     6  te0-0-0-3.mpd21.fra03.atlas.cogentco.com (130.117.49.237)  34.932
    ms  34.392 ms  34.533 ms
     7  tiscali.fra03.atlas.cogentco.com (130.117.14.86)  41.831 ms
    xe-4-3-0.fra23.ip4.tinet.net (77.67.74.41)  41.690 ms  37.861 ms
     8  xe-7-2-0.nyc20.ip4.tinet.net (89.149.183.26)  123.139 ms  127.031
    ms  126.938 ms
     9  netaccess-gw.ip4.tinet.net (213.200.73.122)  136.085 ms  135.549
    ms  138.901 ms
    10  0.e1-4.tbr2.mmu.nac.net (209.123.10.77)  136.521 ms  141.461 ms  138.447 ms
    11  vlan804.esd2.mmu.nac.net (209.123.10.14)  145.733 ms
    vlan805.esd1.mmu.nac.net (209.123.10.34)  142.128 ms  142.070 ms
    12  207.99.53.46 (207.99.53.46)  142.012 ms 207.99.53.42
    (207.99.53.42)  143.203 ms 207.99.53.46 (207.99.53.46)  136.314 ms
    13  search9.infoweb.net (97.107.142.101)  140.275 ms  140.552 ms  140.109 ms

Traces left off at search9.infoweb.net. After some more testing it seemed every fake domain would resolve to an IP that would reverse resolve to searchX.infoweb.net. I shot an email to [my server provider](http://glesys.se/) but they were as clueless as me. When they wanted to charge me for starting to look into the problem I digressed.

This was obviously not the lazy solution I was looking for.

## Resolving the relevant

Then today I got my breakthrough. It was when I went through the code in the [Email Verification module](http://drupal.org/project/email_verify) that I stumbled on a PHP function I did not know of (life of a PHP dev) – [`checkdnsrr()`](http://php.net/checkdnsrr).

What it does is that it checks for a specified [DNS pointer type](http://en.wikipedia.org/wiki/Resource_record#DNS_resource_records) only. I had previously been trying to resolve A pointers for the domains but when I thought about it that was completely irrelevant. What I really wanted to know was whether the domain had [MX records](http://en.wikipedia.org/wiki/MX_record) or not. And the `checkdnsrr()` function defaults to MX. Bingo!

The code has been adjusted, deployed and tested. It looks like it is working and I am crossing my thumbs that it does. Apologies to anyone coming here looking for cheap viagra.

I hope this helps someone else! Please let me know if it does. Just be sure to provide a real email address if you leave a comment. ;)
