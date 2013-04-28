---
layout: post
title: Stopping spam with Symfony forms
category: symfony
tags: [symfony, spam]
summary: I have previously been writing about my efforts to stop spam by resolving MX records for email addresses and by using a honeypot to trick spammer bots into revealing themselves. But still spammers are bypassing my simple protections, so I needed to add another layer and test out one more idea.
notice: This is part of an *[anti-spam](/spam) series*, including [Resolving MX pointers to fight spam](http://local.tobiassjosten.net:1234/php/resolving-mx-pointers-to-fight-spam) and [Stopping spam with a tasty honeypot](/internet/stopping-spam-with-a-tasty-honeypot).
---
I have previously been writing about my efforts to stop spam by [resolving MX records](http://vvv.tobiassjosten.net/php/resolving-mx-pointers-to-fight-spam) for email addresses and by [using a honeypot](http://vvv.tobiassjosten.net/internet/stopping-spam-with-a-tasty-honeypot) to trick spammer bots into revealing themselves.

These two techniques have taken me from ~80 spam comments per day to three in over a week. Quite efficient! But still spammers are bypassing my simple protections, so I needed to add another layer and test out one more idea.

## Timing form submissions

The idea is that automated bots are efficient for spammers because they are fast. They load an HTML page and submit a form in it faster than any human could. And that is how we can catch them.

I started by adding a new element to my comment form.

    $this->widgetSchema['asdf'] = new sfWidgetFormInputHidden(
      array(),
      array('value' => base64_encode(time()))
    );
    $this->validatorSchema['asdf'] = new sfValidatorTimer;

It defaults to always have the current time as its value. It does not matter if this time differs from the client because all validation will take place on the server.

I also obfuscated the element by naming it *asdf* and [base64 encoding](http://en.wikipedia.org/wiki/Base64) its value – two more easy steps that could help throw off the lesser bots.

Next I wrote the validator that I assigned to the element. What follows is its validation method.

    protected function doClean($value)
    {
      $time = base64_decode($value);

      if (!$time)
      {
        throw new sfValidatorError($this, 'tampered');
      }

      if (!is_numeric($time))
      {
        throw new sfValidatorError($this, 'nan');
      }

      $time_ago = time() - $time;

      if ($time_ago > 84600)
      {
        throw new sfValidatorError($this, 'max_time');
      }

      if ($time_ago < 7)
      {
        throw new sfValidatorError($this, 'min_time');
      }

      return time();
    }

This checks to see if the given time, which defaulted to currently, is more than seven seconds ago and less than a day ago. If that is not the case, or if the value has been tampered with, it throws an error. This will again reset the value to the current time and you must wait another seven seconds before posting.

Now let's see if that will keep the remaining spammers away from here.
