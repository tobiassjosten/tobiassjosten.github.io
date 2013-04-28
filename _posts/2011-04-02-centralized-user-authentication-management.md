---
layout: post
title: Centralized user authentication/management
category: security
tags: [security]
summary: I am looking into how you could offer authentication and user management in a centralized manner, to a multitude of satellite sites. The most common way of doing this nowadays, especially for third party satellites, is using *the OAuth protocol* and while it is a very sturdy mechanism it does not fit exactly with my requirements.
---
I am looking into how you could offer authentication and user management in a centralized manner, to a multitude of satellite sites. The most common way of doing this nowadays, especially for third party satellites, is using *the OAuth protocol* and while it is a very sturdy mechanism it does not fit exactly with my requirements.

## The requirements

What I need is a repository to store user data; their names, email addresses, phone numbers, etc and what satellite services they are subscribing to. This data needs to be available to a multitude of different sites. It is also highly important that I can offer [single sign-on capabilities](http://en.wikipedia.org/wiki/Single_sign-on) so the user only needs to login to one sites and then be logged in to all of them.

Because the sites are all served from different domains I [can not have them share cookies](http://en.wikipedia.org/wiki/HTTP_cookie#Domain_and_Path), or it would have been a slim problem. So what I need is something more advanced.

There are plenty of systems that offer similar functionality but they are all missing a key ingredient for me – partial protection. Some pages should be protected as a whole but most often only certain parts of the page needs to require authentication. Think of it as a distributed pay wall where unauthorized users can only see the summary while paying users gets the full article.

The way that most of the common authentication processes fails is their assumption that a GET request could trigger the protection and send the user to an authentication page. I want to be able to determine if a user has already logged in to another satellite site and then give them one page if so and another variant of the page if not. All without interferring with the user's browsing.

## Existing systems

I have looked into a lot of technologies; [LDAP](http://en.wikipedia.org/wiki/LDAP), [OAuth](http://oauth.net/), [OpenID](http://openid.net/), [CAS](http://www.jasig.org/cas), [CoSign](http://www.jasig.org/cas), [JOSSO](http://www.josso.org/), [OpenAM](http://www.forgerock.com/openam.html) and [Pubcookie](http://www.pubcookie.org/). They all cover different aspects of the authentication, single sign-on and data storage requirements. *LDAP* has long been the de facto standard for storing user credentials. *OpenID* provides easy authentication. As does *OAuth*, while also expanding it with third-party, remote control. *CAS*, *CoSign*, *JOSSO*, *OpenAM* and *Pubcookie* are all authentication middleware with single sign-on functionality but with different approaches to it.

CoSign is especially interesting, as it has a ready made [Drupal module](http://drupal.org/project/cosign) and [a PHP library](http://www.fit.vutbr.cz/~lampa/cosign-php/) – enough to cover the needs of these specific satellite sites. It acts as a standalone CGI script that handles all aspects of session management. CoSign uses *factors* to execute the authentication and there is a ready made one for LDAP but you can write your own, in any language you want.

## Loose end

The seamless authentication requirement remains however. I am thinking it could be handled by browser [JavaScript](/javascript). It would poll the master authentication server for a session token to check if the user is logged in. If so the JavaScript would pass the token on to the backend, which could verify the token with the master server and then log the user in without her ever noticing anything.

That final piece of the puzzle would work fine, would it not? Then I think I have enough to wrap this together and start writing a specification.
