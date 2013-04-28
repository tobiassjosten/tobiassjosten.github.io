---
layout: post
title: Using www for your domain
category: internet
tags: [internet]
summary: One of the questions I was trying to sort out while creating tobiassjosten.net was whether to use the www subdomain or not. There are pros and cons both for using it and not using it and deciding on it is not always easy. In this article I will be trying to break it down and hopefully you and me both will be the wiser towards the end.
---
One of the questions I was trying to sort out while creating tobiassjosten.net was whether to use the www subdomain or not. There are pros and cons both for using it and not using it and deciding on it is not always easy. In this article I will be trying to break it down and hopefully you and me both will be the wiser towards the end.

Whatever alternative you do go with, you should keep in mind that today's web users are ignorant of the most basic concepts of the Internet. They can not tell HTTP from FTP, or even a subdomain from a TLD. Not that most people should have reason to know this but we need to consider the fact. With that in mind you need to make sure your website works no matter if the user visits it at http://example.com/ or at http://www.example.com/.

However I wont go into how you could set that up in this article. Instead we will focus on the arguments for and against. Let us go through those.

Semantics dictates that the www part of the domain is superfluous and should go. Since we are telling the browser (or having it default for us) to use HTTP/HTTPS, we have already decided to use port 80/443 and that we will be using the web service (HTTP protocol) of the domain.

DNS resolves one domain to exactly one IP address, no matter which service you will be using. Contrary, you might want to host select services on different machines. Be it for security, convenience or load balancing reasons. One way to split that up is using subdomains, like www.example.com for web browsing and ftp.example.com for file sharing.

Most sites' search engines are subpar. I often resort to using an external search engine to look up pages within a certain domain. Additionally, a lot of sites have more than just one focus, such as a product pages, development blog, support section, etc. Should these be divided into neatly arranged subdomains I can more easily narrow my search down with a query like "domain www site:blog.tobiassjosten.net".

Teaching users to use the www subdomain promotes requests to hosts like www.blog.example.com. And that's just retarded.

Say "tobiassjosten.com" to a random person and they could be reading this article themselves in a couple of seconds. For less common TLD's, foreign ones for example, this is not always as true. Adding www to the domain makes it as unambigious as .com's.

Cookies set for example.com (actually .example.com to be correct) are also sent along with requests to www.example.com. When setting up a CDN however, you do NOT need the extra overhead from cookies. By using www as your canonical subdomain and cookies host you wont be sending those extra kilobytes for each request to cdn.example.com.

In summary I would say it is more theoretically correct to ditch the www subdomain. Sadly I find the practical pros of using it far outweights the cleanliness of not using it. This is very much up to your specific site however and you might come to a whole other conclusion than me, depending on your perspective.

Myself, I picked something inbetween. With vvv.tobiassjosten.net I get the pros of a subdomain while not being burdened by the lame www one. Using www for recognition as an URL is not an issue with my target audience and I can easily redirect everything to vvv, so there is really no problem for me. Plus it sticks out some, does it not?
