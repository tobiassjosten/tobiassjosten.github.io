---
layout: post
title: Expected identifier error in Internet Explorer
category: javascript
tags: [javascript]
summary: I recently bumped into the enigmatic error message Expected identifier in Internet Explorer. For the umpteenth time, I should say, and so I figured I might as well share and document it for future reference.
---
I recently bumped into the enigmatic error message *Expected identifier* in Internet Explorer. For the umpteenth time, I should say, and so I figured I might as well share and document it for future reference.

This expected identifier error message comes from trying to assign an unnamed variable or object property. Which sounds odd but consider this.

    {
      a: "a",
      b: "b",
    }

As a [PHP](/php) developer this looks fine, but it is invalid [JavaScript](/javascript) because there is a missing piece between `"b",` and `}`. That is an expected identifier.

Another cause for this could be the following.

    super = "asdf";

In JavaScript *super* is a [reserved word](https://developer.mozilla.org/en/JavaScript/Reference/Reserved_Words), so Explorer very kindly removes it and instead evaluates this.

    = "asdf";

Absolutely brilliant! So instead of telling us about the real problem, it says that it expects an identifier to assign that string to.

These two problems both will give you errors in other browsers as well, because they are actual errors in your code. But with a proper browser at least you will get a proper error message.

My last example, however, will not raise an error in neither Firefox nor Chrome. But their retarded cousin will of course come and bite you in the ass.

    object.super = "asdf"

Again, in Internet Explorer, this will be evaluated without the reserved word `class` and thus raise the expected identifier error message.

There is [an article on MSDN](http://msdn.microsoft.com/en-us/library/ie/k4eee20x%28v=vs.85%29.aspx) about this problem but it does not give you a lot of help figuring it out. Hopefully this blog post does!
