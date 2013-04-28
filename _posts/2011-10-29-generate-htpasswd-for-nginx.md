---
layout: post
title: Generate htpasswd for Nginx
category: nginx
tags: [nginx]
summary: When I configure a new site on Nginx I always have to look up how to generate those htpasswd files. So here is a quick note for future reference.
---
When I configure a new site on [Nginx](/nginx) I always have to look up how to generate those htpasswd files. So here is a quick note for future reference.

In your vhost file:

    location / {
      auth_basic  "Some message to the user";
      auth_basic_user_file  /etc/nginx/htpasswd;
    }

Generate the hashed password (using [PHP](/php)):

    $ php -a
    php > echo crypt('asdf', base64_encode('asdf'));
    YXWM35gonN/VU

Put it in your htpasswd:

    $ echo 'tobias:YXWM35gonN/VU' >> /etc/nginx/htpasswd
