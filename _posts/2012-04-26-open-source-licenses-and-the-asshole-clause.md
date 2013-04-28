---
layout: post
title: Open source licenses and the asshole clause
category: open-source
tags: [open-source]
summary: I was recently asked add a proper license to one of my open source projects. My immediate reaction was to slap GPL on it and call it a day. But then I figured I should really take some time to research the subject more thorough.
---
I was [recently asked](https://github.com/tobiassjosten/FacebookServiceProvider/pull/7) add a proper license to one of my open source projects. My immediate reaction was to slap GPL on it and call it a day. But then I figured I should really take some time to research the subject more thorough.

It quickly became apparent why I had put this off for so long though. Licensing is a jungle!

There are so many different licenses with so many varying attributes and obscure terminology. From the copyleft *GPL2*, *GPL3*, *LGPL* and *AGPL* to the permissive *MIT* and *BSD* as well as the public *CC0* and *WTFPL*. And that is only the most popular ones.

Instead of trying to compare them to eachother I realized a better approach was to first figure out what I wanted from a license. So I wrote down a shopping list.

## My license criteria

1) I want anyone to be able to download and use my software. The single most important reason for sharing my code is to empower others and help the common good.

2) If you feel like modifying my code you should have the right to do so. There are people out there whom are smarter than me and with other use cases than mine. Our mutual result could easily become greater than our sum.

3) Your improvements to my code should be shared like how I share "the original". With today's ease of sharing back upstream you would have to be an asshole not to do so. My interest in working gratis to help freeloading assholes is approximately zero.

4) My code is a tool and I want to free its use from my own judgements. You should feel safe to implement my code to whatever end you want.

5) I should not be responsible for what happens when you run my code. It is your own risk and I will give no warranties what so ever.

6) You should have the right to make money off of my software without having to pay me a dime. I really want to empower anyone to go ahead and use my code and not worry about any unknown future concequences of doing so.

7) While I would not want someone else taking credit for my work, I have no need to force anyone to credit me for it either. Again I am not doing open source to further my own cause.

8) I want to use an existing, recognizable license so that you do not have to read big chunks of legal mumbo-jumbo.

## A matching license

One big disappointment was learning how the viral nature of the copyleft licenses really works. If you use *GPL* licensed libraries in a project you are distributing, you must license that whole project with GPL. I have no interest what so ever in your project specific implementations, nor do I see the benefit in forcing you to share that.

This is further useless to me because the code I write often is just run on a web server and never actually distributed. So for me the *GPL* and *LGPL* licenses does not affect anything, unless I use the *AGPL* version which again just gives me worthless implementation details.

It could still make sense to use GPL but not for this library of mine. That would only serve to inhibit its use and for no gain what so ever.

The permissive licenses *MIT/X11* and *BSD* are both refreshingly short and to the point and they give exactly the kind of freedom I am looking for. However I do miss support for my asshole clause - where you need to share your improvements to my code.

Finally the public domain licenses *WTFPL* and *CC0* are even more open ended than the MIT and BSD ones. While giving that much freedom to potential users is nice, it does not protect this freedom from abuse.

So much research and no good match. I was just about to throw in the towel when I discovered the [Mozilla Public License](http://www.mozilla.org/MPL/2.0/) and realized I had struck licensing gold!

*MPL 2.0* lets you use my code together with other code, side by side. Just like with GPL you need to release your modifications to my code when distributing your work but unlike copyleft that does not include all code - just the MPL licensed parts.

As with most licenses (AGPL aside) this is veered towards compiled software that is distributed in executable binaries. Not exactly my use case and so the asshole clause is probably kind of moot. But it is as close as I can seem to get and at least it sends some kind of a message about my wishes.

This will suffice for now. I do not want to waste a single hour more reading legal texts when I could do some actual hacking instead!

Reference: [GPL2](http://www.gnu.org/licenses/old-licenses/gpl-2.0.html), [GPL3](http://www.gnu.org/licenses/gpl.html), [LGPL](http://www.gnu.org/licenses/lgpl.html), [AGPL](http://www.gnu.org/licenses/agpl.html), [BSD](http://www.xfree86.org/3.3.6/COPYRIGHT2.html#6), [MPL](http://www.mozilla.org/MPL/2.0/), [MIT](http://www.opensource.org/licenses/mit-license.php), [ISC](http://www.isc.org/software/license), [CDDL](http://hub.opensolaris.org/bin/download/Main/licensing/cddllicense.txt), [APL](http://www.opensource.org/licenses/apl1.0.php), [Apache 1.1](http://www.apache.org/licenses/LICENSE-1.1), [Apache 2.0](http://www.apache.org/licenses/LICENSE-2.0), [WTFPL](http://sam.zoy.org/wtfpl/), [CC0](http://creativecommons.org/publicdomain/zero/1.0/legalcode).
