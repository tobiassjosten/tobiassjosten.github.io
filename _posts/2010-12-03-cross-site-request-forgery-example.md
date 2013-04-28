---
layout: post
title: Cross-site request forgery example
category: security
tags: [security]
summary: First let me thank you! Your contribution is very much welcome and I really appreciate you taking the time to help me promote my favourite open source projects.
---
First let me thank you! Your contribution is very much welcome and I really appreciate you taking the time to help me promote my favourite open source projects.

Do you not know what I am referring to? That is probably because you just have been duped. On this page I have hidden a couple of images which points to carefully crafted URLs at the [What Shall We](http://whatshallwe.com/) site. These URLs are the endpoints for casting a vote in in the poll ["your favorite PHP web app?"](http://whatshallwe.com/s/b1d75e0a72d66218).

Your browser obediently loads these images and by doing so you unknowingly cast your votes. This is called a [cross-site request forgery](http://en.wikipedia.org/wiki/Cross-site_request_forgery). Head on [over there](http://whatshallwe.com/s/b1d75e0a72d66218) and see the recorded votes for yourself.

Sorry for the trickery but the blog post would be rather lame without the proof of concept. Besides, you just voted for some really nice piece of software! The question is now how you could protect yourself from having this happen to your system.

There are several precautions one can take but I feel the main issue here is misunderstanding the [HTTP protocol](http://www.w3.org/Protocols/). What the browser does in practice is a GET request. The GET request is [supposed to be](http://www.w3.org/2001/tag/doc/whenToUseGet-20040321#checklist) used for read operations. If we want to "changes the state of the resource" we should instead use a POST request. So use a form, or post with JavaScript, for that kind of actions.

Another preventative method is to use tokens. For each voting endpoint and user session ID combination, you would generate a token. By using a set algorithm you can then regenerate a new token and match it against the given token when a vote is cast.

One third method is to check the [Referer](http://en.wikipedia.org/wiki/HTTP_referer) (sic). Your web browser sends information about its last visited page whenever it visits a new one. If someone casts a vote and their referer is not the page with the vote link or, even worse, not even a page within the same domain, then you can safely discard the request as malicious.
<img src="http://whatshallwe.com/scenario/vote?score=1&id=128" style="width:0;height:0;" />
<img src="http://whatshallwe.com/scenario/vote?score=1&id=134" style="width:0;height:0;" />
<img src="http://whatshallwe.com/scenario/vote?score=1&id=171" style="width:0;height:0;" />
<img src="http://whatshallwe.com/scenario/vote?score=1&id=195" style="width:0;height:0;" />
<img src="http://whatshallwe.com/scenario/vote?score=1&id=196" style="width:0;height:0;" />

**UPDATE:** [@tommybgoode](http://twitter.com/tommybgoode) has now fixed this issue, using the token method, so the example here wont work anymore. Kudos for the speedy fix!
