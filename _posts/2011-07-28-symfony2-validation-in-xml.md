---
layout: post
title: Symfony2 validation in XML
category: symfony
tags: [symfony]
summary: I have made it through half of the Symfony2 bible and am now reading up on validation. I am using XML for my bundle configuration, so I switched to that format in the examples and used the following to constraint my Log entity.
---
I have made it through half of [the Symfony2 bible](http://symfony.com/doc/current/book/index.html) and am now reading up on [validation](http://symfony.com/doc/current/book/validation.html). My bundles uses XML configuration, so I switched to that format in the examples and used the following to constraint my Log entity.

    <class name="Nogfx\LogBundle\Entity\Log">
        <property name="title">
            <constraint name="NotBlank" />
        </property>
    </class>

Sadly it crashed with an error message:

    [ERROR 1845] Element 'class': No matching global declaration available for the validation root. (in /app/dir/src/Nogfx/LogBundle/Resources/config/validation.xml - line 1, column 0)

    500 Internal Server Error - MappingException

So the XML was obviously malformed. But how then should it be structured?

After some digging around I managed to find [an example](https://github.com/FriendsOfSymfony/FOSUserBundle/blob/master/Resources/config/validation.xml) and that cleared it out. Given my previous example, this is how it should have been structured instead.

    <?xml version="1.0" encoding="UTF-8" ?>
    <constraint-mapping
        xmlns="http://symfony.com/schema/dic/constraint-mapping"
        xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"
        xsi:schemaLocation="http://symfony.com/schema/dic/constraint-mapping
            http://symfony.com/schema/dic/services/constraint-mapping-1.0.xsd">

        <class name="Nogfx\LogBundle\Entity\Log">
            <property name="title">
                <constraint name="NotBlank" />
            </property>
        </class>

    <constraint-mapping>

Just like with [the routing](http://symfony.com/doc/current/book/routing.html#routing-in-action), we need to define a root element with a proper namespace and DTD. Adding that solved my problem.

This should of course be fixed in the documentation and I will do so myself if no one else beats me to it. Still, I hope this post can help someone with the same problem before the fix is in.

*Update:* [Ryan Weaver](http://www.thatsquality.com/ryan) accepted [my pull request](https://github.com/symfony/symfony-docs/pull/588) and this problem should be fixed on next deployment.
