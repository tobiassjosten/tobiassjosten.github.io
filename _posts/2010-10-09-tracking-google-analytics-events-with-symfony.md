---
layout: post
title: Tracking Google Analytics events with Symfony
category: symfony
tags: [symfony, google-analytics]
summary: I just finished implementing a first whack at tracking events with Google Analytics. For this experiment I picked the "elsewhere links" in the footer. Now when you click one these, you will trigger an event which will be tracked by Google Analytics. The idea is that I can then easily see which one is more popular, how their popularity changes over time, how nationality affects popularity, etc.
---
I just finished implementing a first whack at [tracking events](http://code.google.com/apis/analytics/docs/tracking/eventTrackerGuide.html) with [Google Analytics](http://www.google.com/analytics/). For this experiment I picked the "elsewhere links" in the footer. Now when you click one these, you will trigger an event which will be tracked by Google Analytics. The idea is that I can then easily see which one is more popular, how their popularity changes over time, how nationality affects popularity, etc.

Of course I only have a handful of visitors per day and so the experiement is more a proof of concept than anything else. That said, now is a great time to pick one or more links and come find me on other places of the Internet!

## Google Analytics in Symfony

I built this site using [Symfony](http://www.symfony-project.org/). Why and how is a story better left for another blog post. What is interesting here is how I implemented Google Analytics.

Obviously I'm using a plugin for this. And when it comes to plugins you are bound to some day come across one made by [Kris Wallsmith](http://kriswallsmith.net/). I did so with [sfGoogleAnalyticsPlugin](http://github.com/kriswallsmith/sfGoogleAnalyticsPlugin). It gives you a lot of control over the generated javascript and I learned about some new ways of using Google Analytics, like [disabling data collection](http://github.com/kriswallsmith/sfGoogleAnalyticsPlugin/blob/master/lib/tracker/sfGoogleAnalyticsTrackerGoogle.class.php#L81) for example.

To do event tracking, however, you will need to implement the asynchronous javascript and that is something the plugin was missing. So I dug in and added [a new asynchronous tracker class](http://github.com/tobiassjosten/sfGoogleAnalyticsPlugin/commit/ff258e218a9eef25816b97c77740bf1788ca5095) based on Kris' work. It is still experimental and if you find anything to improve, please do.

Using the plugin was simple. First you enable it in your ProjectConfiguration.class.php:

    class ProjectConfiguration extends sfProjectConfiguration
    {
      public function setup()
      {
        $this->enablePlugins('sfGoogleAnalyticsPlugin');
      }
    }

Then you add it to your filters.yml:

    sf_google_analytics_plugin:
      class: sfGoogleAnalyticsFilter

And finally you configure it in your app.yml:

    sf_google_analytics_plugin:
      enabled:       on
      profile_id:    UA-XXXXXXXX-X
      tracker:       asynchronous
      insertion:     top
      params:
        domain_name:   false
        linker_policy: true

## Javascripting Google Analytics

For the business logic I decided to base it on jQuery. As a result the code is quite slim:

    $(function() {
      $('#block-about a').each(function() {
        var _gaq = _gaq || [];
        $(this)
          .bind('click', function() {
            _gaq.push(['_trackEvent', 'Elsewhere', this.rel]);
          })
          .attr('target', '_blank');
      });
    });

It iterates all my "elsewhere links" and binds a function to the 'click' event. When it triggers, the function calls the Google Analytics code to track the event. Simple as that!

As a precaution I also added the target attribute, with "_blank" as value, so that the XHR request wont be interrupted.

## Conclusion

One thing I would like to clean up is the hard coded GA object name in my javascript. sfGoogleAnalyticsPlugin uses a variable name, for some reason, and so '_gaq' is not set in stone. Perhaps base it off of a value set in the app.yml configuration? It depends on what the reason is for the variable naming I guess.

I have yet to test to see if this actually work (crash early, crash cheap, right?) but if it does then I will want to post a follow-up here.
