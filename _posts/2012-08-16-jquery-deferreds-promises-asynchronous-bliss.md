---
layout: post
title: jQuery Deferreds promises asynchronous bliss
category: javascript
tags: [javascript, jquery]
summary: When running an asynchronous function, only the function itself knows when it is done. Thus if we want to continue after the function has completed we need to pass it a callback.
---
When running an asynchronous function, only the function itself knows when it is done. Thus if we want to continue after the function has completed we need to pass it a callback.

    fetch('/stuff', function (stuff) {
        save(stuff, db, function () {
            animate(element, {opacity: .5}, function () {
                alert('Cool, huh?');
            });
        });
    });

That works, though it can evidently become quite messy. However to solve that nesting hell you can just formalize the anonymous functions as named ones.

But what if we want to supply a callback for when two simultaneous, asynchronous functions both have completed? What if we want to `alert()` when we have fetched both from `/stuff` and `/things`?

We could have them share state.

    var stuff = things = false;
    fetch('/stuff', function (stuff) {
        if (things) { doTheAlert(); }
        stuff = true;
    }
    fetch('/things', function (things) {
        if (stuff) { doTheAlert(); }
        things = true;
    }

It works but, eh. We can do better. We can start making [Promises](http://wiki.commonjs.org/wiki/Promises/A).

## Promises

A *promise* is a construct used for synchronization in concurrent languages, like [JavaScript](/javascript). The *Promises/A* definition from [CommonJS](http://www.commonjs.org/) provides us with the spec we need and [jQuery since 1.5](http://blog.jquery.com/2011/01/31/jquery-15-released/) brings [an implementation](http://api.jquery.com/category/deferred-object/) we can use.

In jQuery this [ancient concept](http://en.wikipedia.org/wiki/Futures_and_promises) is accessed through `$.Deferred()`, giving us a `Deferred` object. This object can be [resolved](http://api.jquery.com/deferred.resolve/) or [rejected](http://api.jquery.com/deferred.reject/) once its associated operation finishes.

Code speaks clearer than words, so let us dive in and look at an example.

    // Declare that we will clean our room.
    var WillCleanMyRoom = $.Deferred();

    // Set up the callback for when we have finished the task.
    WillCleanMyRoom.done(function () {
        giveWeeklyAllowance();
    });

    // Now clean the room … after a little nap.
    setTimeout(function () { WillCleanMyRoom.resolve(); }, 5000);

After five seconds the timeout will resolve our deferred object, which in turn will give us our weekly allowance. Pretty nice, huh?

Sadly, it turns out our parents have more chores for us before we can have the allowance. Not only do we have to clean our room but we must also take out the trash.

    var WillCleanMyRoom = $.Deferred();
    setTimeout(function () { WillCleanMyRoom.resolve(); }, 5000);

    var WillTakeOutTheTrash = $.Deferred();
    WillTakeOutTheTrash.resolve();

    var WillDoChores = $.when(WillCleanMyRoom, WillTakeOutTheTrash);
    WillDoChores.done(function () {
        giveWeeklyAllowance();
    });

So we create two deferred objects, each to be resolved separately, and then we chain them together to a new deferred object; `WillDoChores`. Once chained we can set up our completion callback for both the chores.

## Deferreds?

So what is this talk about *promises* when all we used was deferreds? They are actually two sides of the same coin. The deferred can be resolved or rejected. It can also create a promise object, which is what we bind our callbacks to.

By keeping this separation we can encapsulate the deferred object in one scope, to have it be resolved or rejected by the responsible agent, while being able to pass around the promise object for subscribers. We do not want just anyone to be able to accidently trigger the bound callbacks.

When creating asynchronous functions yourself it is a good idea to have it return a promise.

    function countDown(time) {
        var timer = $.Deferred();
        setTimeout(function () { timer.resolve(); }, time);
        return timer.promise();
    }

    var WillCountDown = countDown(5000);
    WillCountDown.done(function () {
        alert('Time is up!');
    });

This way anyone using your function can easily bind as many callbacks as they please, without you having to accept them as callback parameters.

In fact, this is just what [`$.ajax()`](http://api.jquery.com/jQuery.ajax/) does; it returns a promise which you can add your callbacks to.

    $.ajax('/stuff')
        .done(function () { alert('We haz stuff!'); });

## Broken promises

Life is not all sunshine and rainbows however. Your asynchronous operations can fail and we need a way to handle that. Previously we have only been using `.resolve()`, now let us meet `.reject()`.

You can bind callbacks to a promise in three ways; with [`.done()`](http://api.jquery.com/deferred.done/) for when its deferred is resolved, with [`.fail()`](http://api.jquery.com/deferred.fail/) for when it is rejected and with [.then()](http://api.jquery.com/deferred.then/) for when it is either resolved or rejected.

    var WillCleanMyRoom = $.Deferred();
    WillCleanMyRoom
        .done(function () {
            giveWeeklyAllowance();
        })
        .fail(function () {
            groundedForAWeek();
        })
        .then(function () {
            giveNewChores();
        });

Even our ever loving parents have their limits. If we fail to carry out our chores we are not only withheld the allowance, we will also be grounded for a week. Equally sucky, whether we succeed or not, there will be a limitless supply of new chores…

Yours right now should be to start making more Promises!
