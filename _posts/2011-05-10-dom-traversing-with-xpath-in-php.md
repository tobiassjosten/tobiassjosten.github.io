---
layout: post
title: DOM traversing with XPath in PHP
category: php
tags: [php]
summary: You can extract data from an XML or (X)HTML document by using regexp. However as in many cases with techniques and systems – just because you can does not mean you should.
---
You can extract data from an XML or (X)HTML document by using regexp. However as in many cases with techniques and systems – just because you can does not mean you should.

Regexp exists to work with plain text and that is what it does best. For structured data, as is the case with XML and (X)HTML, we have more suitable tools at our disposal.

The libraries you will see most often used in PHP are [DOM](http://se2.php.net/manual/en/book.dom.php) and [SimpleXML](http://se2.php.net/simplexml). There [are others](http://se2.php.net/manual/en/refs.xml.php) built in or existing in PECL and there is two popular external ones named [PHP Simple HTML DOM Parser](http://simplehtmldom.sourceforge.net/) and [QueryPath](http://querypath.org/).

For this demonstration we will use DOM. It starts off by instantiating the DOM object and loading it with data.

    $dom = new DOMdocument();
    $dom->loadHTML($document);

Next is the matter of picking out our wanted nodes. In our use case we are looking for the *src* attribute of an image with the class *this*. The image can have multiple classes and so we need to be careful not to match against a static 'this' string.

One simple method to look up our data is to fetch all *img* elements and inspect them one by one until we find the desired class.

    foreach ($dom->getElementsByTagName('img') as $node) {
        if (!$node->hasAttribute('class') || !$node->hasAttribute('src')) {
            continue;
        }

        if (preg_match('~\bthis\b~', $node->getAttribute('class'))) {
            return $node->getAttribute('src');
        }
    }

It is simple to understand and we will find what we are looking for. However it is a performance hog; both because we are traversing an unknown number of elements in PHP and because we are still reverting to regexp for matching. Of course we could match the string more efficiently but then there is a tradeoff in readability.

There are better ways to skin this cat though. In XML we have a technique called [XPath](http://www.w3.org/TR/xpath/) which is used to traverse and lookup data in structured DOMs. Think of it as [jQuery's Sizzle](http://sizzlejs.com/) on speed.

    $xpath = new DOMXpath($dom);
    $nodes = $xpath->query('*/img[contains(@class, "this")]');

    if ($nodes->length) {
        return $nodes->item(0)->getAttribute('src');
    }

This searches for an img tag which class attribute contains *this*. Now we are letting the C implementation do the heavy lifting and we have a much more performant solution. Though I must confess I have not benchmarked the difference.

Our solution uses the *contains* operator and this will match both *this* and *notthis*. Depending on our actual use case this can be a problem. I found a much more [elegant solution](http://westhoffswelt.de/blog/0036_xpath_to_select_html_by_class.html) to this by [Jakob Westhoff](http://westhoffswelt.de/).

    $xpath = new DOMXpath($dom);
    $nodes = $xpath->query('*/img[
        contains(normalize-space(@class), " this ")
        or substring(normalize-space(@class), 1, string-length("this") + 1) = "this "
        or substring(normalize-space(@class), string-length(@class) - string-length("this")) = " this"
        or @class = "this"
    ]');

    if ($nodes->length) {
        return $nodes->item(0)->getAttribute('src');
    }

There we go – now it will only match the exact *this* class. Happy traversing!
