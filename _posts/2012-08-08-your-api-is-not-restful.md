---
layout: post
title: Your API is not RESTful
category: development
tags: [development, rest, api]
summary: In software development it is always good to take a look at existing *design patterns* and *standards* before deciding on a path. Two such leading standards for API design used to be *XML-RPC* and *SOAP* but nowadays *RESTful API* is all the hype. It is clean, powerful and easily maintainable. But also completely misunderstood.
---

In software development it is always good to take a look at existing *design patterns* and *standards* before deciding on a path. Two such leading standards for API design used to be *XML-RPC* and *SOAP* but nowadays *RESTful APIs* is all the hype. It is clean, powerful and easily maintainable. But also completely misunderstood.

This is my gripe with "RESTful APIs" — companies are boasting openly about their shiny new API and how RESTful it is. And then it turns out they have implemented RPC over HTTP. Maybe they even throw in a link to [the Wikipedia article](https://en.wikipedia.org/wiki/REST). But somehow they fail to grasp the [guiding principles](https://en.wikipedia.org/wiki/REST#Guiding_principles_of_the_interface) in that article.

The very *essence of REST* is about four distinct criteria.

## RESTful criteria

**Identification of resources.** Opposed to RPC calls, REST is about working with resources instead of methods. If you are creating users by calling `/users/register?name=tobias` you are doing it wrong. Instead you probably want to `POST` data to the `/users` collection.

**Manipulation of resources through these representations.** All resources should be represented in a format like *JSON* and *XML*. More exotic resources like geometric shapes or geographical information can be represented in formats like *SVG* or *WKT*.

**Self-descriptive messages.** One response should be all that is needed to consume the data. Using HTTP for example, you should have a `Content-Type` header to describe the resource's representation. This should also force the communication to be agnostic to whatever cache layers, proxies, gateways, etc are between the endpoints.

**Hypermedia as the engine of application state.** Building on the above point, resource representation should describe how to manipulate themselves. You should not have to read documentation to know that you should `PUT` the user `{"name":"tobias"}` to resource `/users`. This should all be *discoverable* through the API.

Those are the REST criteria. You can not pick and choose.

Most "RESTful APIs" make it as far as to the last point but the majority [fail the hypermedia constraint](http://roy.gbiv.com/untangled/2008/rest-apis-must-be-hypertext-driven). However this is so innate to the architecture that is has its own abbreviation — [HATEOAS](https://en.wikipedia.org/wiki/HATEOAS). It is also what makes RESTful design stand out and be as powerful as it is.

All you should need to integrate with an API is knowledge about the resources and a starting point. Everything else should be discoverable through consumption of the resources.

So in conclusion; your API is probably **not RESTful**.

## … and that is fine

It does not matter that your API is not RESTful. It does not need to be.

Especially if you are developing an internal API, where you control both the client and the server, there is less to gain from the discoverable aspect of a *RESTful architecture*. Or if your audience of developers are less savvy it might be too big a hurdle for them.

There are still pros to REST but you need to figure out if they are relevant for you. I recommend informing yourself by reading the [network architecture design dissertation](http://www.ics.uci.edu/~fielding/pubs/dissertation/top.htm) by [Roy Fielding](http://roy.gbiv.com/). The [chapter on REST](http://www.ics.uci.edu/~fielding/pubs/dissertation/rest_arch_style.htm) in particular.

If you then decide against going RESTful that is totally cool. Just stop calling it something it is not.
