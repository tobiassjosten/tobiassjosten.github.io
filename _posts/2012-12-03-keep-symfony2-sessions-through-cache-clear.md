---
layout: post
title: Keep Symfony2 sessions through cache:clear
category: symfony
tags: [symfony]
summary: Per default Symfony2 will store session information in its cache directory. Clearing the cache thus logs everyone out. A simple solution to this is moving where Symfony saves its sessions.
---
Per default [Symfony2](/symfony/) will store session information in its cache directory. Clearing the cache thus logs everyone out.

A simple solution to this is moving where Symfony saves its sessions. We can do this by configuring the framework bundle via our `app/config/config.yml`.

    framework:
        session:
            save_path: %kernel.root_dir%/var/sessions

That one line `save_path` setting will redirect session data to `app/var/sessions`. Problem solved!

If you are also using [Capifony](http://capifony.org/) (and you should!) then remember to share the sessions between releases, in your `deploy.rb`.

    set :shared_children, [app_path + "/var"]
