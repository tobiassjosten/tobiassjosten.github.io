---
layout: post
title: Logging Doctrine SQL queries in Symfony2
category: symfony
tags: [symfony, doctrine]
summary: Doctrine is a great ORM and DBAL library to use with your Symfony2 application. But occasionally it can feel like a black box, where magic just happens. In those cases it would be good to have a look at what actually happens.
---
Doctrine is a great ORM and DBAL library to use with your Symfony2 application. But occasionally it can feel like a black box, where magic just happens. In those cases it would be good to have a look at what actually happens.

Luckily, Doctrine can easily be configured to log its doings. From a container aware object, such as your controllers or commands, you can enable logging like this.

    $this
        ->get('doctrine')
        ->getConnection()
        ->getConfiguration()
        ->setSQLLogger(new \Doctrine\DBAL\Logging\EchoSQLLogger());

This will set up [EchoSQLLogger](http://www.doctrine-project.org/api/dbal/2.2/class-Doctrine.DBAL.Logging.EchoSQLLogger.html) to do your logging for you, which results in every query and its parameters being printed with `echo` and `var_dump()`. Quick, easy and not too dirty!

## Supressed logging

You could also use the [DebugStack logger](http://www.doctrine-project.org/api/dbal/2.2/class-Doctrine.DBAL.Logging.DebugStack.html), which will record the queries without printing them. Just set it up similarly to as previously shown.

    $logger = new \Doctrine\DBAL\Logging\DebugStack();
    $container
        ->get('doctrine')
        ->getConnection()
        ->getConfiguration()
        ->setSQLLogger($logger);

And after some Doctrine usage, access your logged activity in the `queries` property.

    var_dump($logger->queries);

## Custom logging

If that is not enough, you can always create your own logger with whatever crazy functionality you need. As long as it implements the [SQLLogger interface](http://www.doctrine-project.org/api/dbal/2.2/class-Doctrine.DBAL.Logging.SQLLogger.html) it can be injected with the `setSQLLogger` method.
