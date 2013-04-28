---
layout: post
title: jQuery ID selector performance
category: javascript
tags: [javascript, jquery, benchmark]
summary: I recently stumbled over a piece of JavaScript code looking something like this.
---
I recently stumbled over a piece of [JavaScript code](/javascript) looking something like this.

    $('#my-div').css('color', 'black');
    $('#my-div').addClass('asdf');
    $('#my-div').show();
    if ($('#my-div').hasClass('this')) {
      $('#my-div').addClass('that');
    }

A much  better approach would be to instantiate the jQuery object only once and then save it for re-use. While we are at it we could also have it use jQuery's chaining feature.

    var $mydiv = $('#my-div');
    $mydiv
      .css('color', 'black')
      .addClass('asdf')
      .show();
    if ($mydiv.hasClass('this')) {
      $mydiv.addClass('that');
    }

That is much more readable if you ask me.

Since [ID selectors](http://api.jquery.com/id-selector/) uses the native [getElementsById()](https://developer.mozilla.org/en/DOM/document.getElementById) method I was pretty confident just saving the object would not affect its performance much. But I did not know for a fact just how much and so I set out to conduct a little benchmark.

## jsPerf

The [jsPerf](http://jsperf.com/) tool, by [Mathias Bynens](http://mathiasbynens.be/), has lately become one of my more frequently used ones. It lets you set up JavaScript snippets that you run against each other to produce pretty graphs for which is faster. You can even run the tests in multiple browsers and easily find incompatibilities.

I hacked together a test and found that the difference between re-using an object and instantiating it over and over is HUGE.

After some feedback I have revised the test and thrown in two more approaches for reference. I would say it proves beyond any doubt that **you should not re-instantiate a jQuery object** if you can help it.

[http://jsperf.com/jquery-id-selector-performance](http://jsperf.com/jquery-id-selector-performance)

The real kicker, however, is how much faster instantiation is, given a proper DOM element. Almost three times as fast! I have no idea why but now I am definitely intrigued and wanting to dig deeper into this.
