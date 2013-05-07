---
layout: post
title: Responsible Silex controllers
category: development
tags: [development, silex, php]
summary: 
---
Micro-framework [Silex](/silex/) is an excellent wrapper around the [Symfony components](/symfony/). It abstracts away a lot of the complexities, so you can get right to the nitty gritty details.

    $app = new Application();
    $app->get('/foo', function() {
        return new Response('Bar');
    });

You can also pull in `TwigServiceProvider` for an excellent rendering engine to craft your beautiful websites. However, what happens if what you want to show is not a website but formatted data? What if you are building an API?

## API responses

I recently bumped into this problem where I first wanted to assemble the data in my controller and then have it formatted to JSON or XML, depending on what was requested.

    $app->get('/foo', function(Request $request) {
        $data = assembleData();
        if (/* check if $request supports JSON */) {
            return new Response(
                json_encode('Bar'),
                200,
                ['Content-Type' => 'application/json']
            );
        } elseif (/* check if $request supports XML */) {
            return new Response(
                /* format $data in XML */,
                200,
                ['Content-Type' => 'application/xml']
            );
        }
    });

It works, but copying that across every controller makes for some really shitty code. I could of course chuck it into `$app['format']` and use that in my controllers but I wanted to take it one step further to free my controllers doing anything other than fetching and assembling data.

The result is [`ResponsibleServiceProvider`](https://github.com/tobiassjosten/ResponsibleServiceProvider)!

## Automatic response formatting

With `ResponsibleServiceProvider` all you need to do is return an array. An event listener will then pick that up, check what format has been requested and do the appropriate formatting.

    $app->get('/foo', function() {
        return ['Bar'];
    });

When a client requests the JSON representation of your data, it will get it.

    $ curl -I -H 'Accept: application/json' http://example.com/foo
    HTTP/1.1 200 OK
    Date: Tue, 07 May 2013 08:30:58 GMT
    Server: Apache/2.2.22 (Ubuntu)
    X-Powered-By: PHP/5.4.9-4ubuntu2
    Cache-Control: no-cache
    Transfer-Encoding: chunked
    Content-Type: application/json
    
    ["Bar"]

And the same goes for XML.

    $ curl -I -H 'Accept: application/xml' http://example.com/foo
    HTTP/1.1 200 OK
    Date: Tue, 07 May 2013 08:30:58 GMT
    Server: Apache/2.2.22 (Ubuntu)
    X-Powered-By: PHP/5.4.9-4ubuntu2
    Cache-Control: no-cache
    Transfer-Encoding: chunked
    Content-Type: application/xml
    
    <?xml version="1.0"?>
    <response><item key="0">Bar</item></response>

[`ResponsibleServiceProvider`](https://github.com/tobiassjosten/ResponsibleServiceProvider) is released under the MIT license and all you need to start using it is to fetch the [Packagist package](https://packagist.org/packages/tobiassjosten/responsible-service-provider).
