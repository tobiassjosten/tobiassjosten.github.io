---
published: false
layout: post
title: Law of Demeter
category: development
tags: [development]
summary: 
---
Once an object directly interacts with another object, they are coupled. Changing the behavior of the interactee could then easily crash the interactor and for this reason we want to keep our coupling to a minimum.

One easy set of rules to achieve this is the *Principle of Least Knowledge* or better known as *Law of Demeter* (LoD). Learn it, consider it and your code has a very good chance to be nicer and cleaner.

The Law of Demeter can be summarized like so:

- Objects should have very limited knowledge about other objects and not rely at all on their internal structure.

- Objects should only talk to immediate friends and never to strangers.

Basically your objects should be very shy to whom they interact with and what they assume in their interactees.

## Limited knowledge

A line of code is worth a thousand words, right?

    class Controller
    {
        public function showAction($request)
        {
            $request->getCookies()->addCookie('LOLZ');
            // NO-NO! Our Controller does not have direct access to the
            // cookies object and thus
        }
    }
