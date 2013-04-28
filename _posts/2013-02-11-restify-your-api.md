---
layout: post
title: RESTify your API
category: development
tags: [development, rest, api]
summary: If you are interested in reading up on RESTful architecture; here are some really good resources.
---
Last night I was discussing [RESTful architecture](/rest/) on Twitter with [Emil Stenström](https://twitter.com/EmilStenstrom). Specifically how yet another API said it was RESTful, when it really was not.

You see these all the time now and while some will claim the term has lost its original meaning, I think REST is way too nice to have diluted because a bunch of [jimmies](http://www.codinghorror.com/blog/2012/07/new-programming-jargon.html) could not be arsed to RTFM.

So I wanted to whip up a blog post to show how to move this API from its current RPC like style, to being truly RESTful. Five hours in I realized it had grown out of control and would take me the entire day to complete… Since I do not have that kind of luxury and because others have already explained it so well, instead I will curate their efforts.

## RESTful reading

**Discoverability.** [Jeremy H](http://www.blogger.com/profile/06146930248780549129) goes through [how to consume a RESTful API](http://thereisnorightway.blogspot.se/2012/05/api-example-using-rest.html) in a way that concretizes the discovery aspect of REST. Media types, like `application/vnd.example.coolapp.apiIndex-v1+xml` are defined by the API and these should be the major focus of the documentation.

**Cachability.** [Mark Nottingham](https://twitter.com/mnot) has an excellent [caching tutorial](http://www.mnot.net/cache_docs/), which can help shed some light on the important cachable property of a RESTful API.

**Web basics.** Because REST tunes in well with the HTTP protocol, [RFC 2616](http://www.ietf.org/rfc/rfc2616.txt) and [RFC 3986](http://www.ietf.org/rfc/rfc3986.txt) are well worth at least skimming through.

**RESTful criteria.** Yours truly goes through [the RESTful criteria](/development/your-api-is-not-restful/).

**Workflow and links.** [InfoQ](http://www.infoq.com/about) shows [how to order a cup of coffee](http://www.infoq.com/articles/webber-rest-workflow) by means of a RESTful API.

**HATEOAS.** [Matt Cottingham](https://twitter.com/mattrcottingham) explains clearly [what hypermedia controls is](http://blueprintforge.com/blog/2012/01/01/a-short-explanation-of-hypermedia-controls-in-restful-services/) and why HATEOAS is important.

**Maturity model.** [Tobias Nyholm](http://www.tnyholm.se/om-tobias/) commented with a link to an very good read on [Richardson Maturity Model](http://martinfowler.com/articles/richardsonMaturityModel.html), by [Martin Fowler](http://martinfowler.com/aboutMe.html).
