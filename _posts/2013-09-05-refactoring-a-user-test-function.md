---
layout: post
title: Refactoring a user test function
category: development
tags: [development, php]
summary: Follow along as I pull out a function from a legacy system and, step by step, refactor it to meet higher quality standards.
---
Part of working as a consultant is wrestling legacy code. There is no limit to the WTF you can find in these systems and I often find myself wondering "what were they thinking?".

Building a big system with a lot of unknowns is no simple task however. If you are programming something of significant size you will write one or two lines of WTF code yourself. What is important is that you identify these and keep on working with your code, refactoring as you learn.

To that end I pulled out a function I recently stumbled over. It determines if a user is eligible for new features before they are released to all users.

    function customer_is_test($user_id)
    {
      $test_users = array(1234, 5678);
      
      if (!variable_get('production', false))
      {
        return true;
      }
      else if (in_array($user_id, $test_users))
      {
        return true;
      }
      else
      {
        return false;
      }
    }

## Simplify if logic

The very first sign of trouble is the long chain of if statements. Because it always ends with returning a boolean value (`false` or `true`) and the if conditions are nothing but boolean evaluators, we can simplify this a lot.

    function customer_is_test($user_id)
    {
      $test_users = array(1234, 5678);
      
      return !variable_get('production', false)
        || in_array($user_id, $test_users);
    }

We have the exact same result but a much better overview of what happens. There is no mistaking this function returns a boolean depending on the given conditions.

## Abstract opportunity

We trust this function to handle all the logic about whether the user is a tester or not. The whole idea here is to abstract that away so we do not have to bother with it ourselves. Yet we are giving it a user id.

This implies we know something about how the function implements its logic (it is based on user id) and that is a missed opportunity for abstraction. It also means we have a forwards compatibility problem, as if we ever wanted to switch this logic around we would need to also change all calls to this function.

Instead let us require the entire user object.

    function customer_is_test(UserInterface $user)
    {
      $test_users = array(1234, 5678);
      
      return !variable_get('production', false)
        || in_array($user->getId(), $test_users);
    }

I added in a typehint for `UserInterface` as well, so that we are 100% certain we are getting a proper user. That way we know for sure we can use the `getId()` method.

## Model domain logic

The tester attribute is entirely contained within the user domain. There are no outside dependencies for determining whether a certain user is a tester, so why even have an external function for it?

Let us instead make this a method directly on the `User` object.

    class User
    {
        public function isTest()
        {
            $test_users = array(1234, 5678);
            
            return !variable_get('production', false)
                || in_array($this->getId(), $test_users);
        }
    }

Now when we have a `$user` we do not have to pass it into a separate function but everthing is nicely packed together in the user itself.

## Environment configuration

Currently the method checks for a certain `production` variable, to see whether the application is in production, development or any other environment. The idea is that all users are test users outside of production, in which only a select few are.

It is a good idea but the implementation can certainly be improved. Instead of having our code rely on environment we should abstract this to configuration, which in turn is set by the environment.

This is a subtle but important change because it enables us to easily add in new environments (testing, qa, performance, etc) and our business logic becomes more flexible as we do not have to change the code for new practices but only update the configuration.

In our example we can achieve this by only relying on the configured list of test users.

    class User
    {
        public static $testUsers = array(1234, 5678);
        
        public function isTest()
        {
            return in_array($this->getId(), $this::$testUsers);
        }
    }

… and then configuring that list by environment.

    // In development, test, etc.
    User::$testUsers = range(1, 999999);

Preferrably you would also not hard code the list of testers like above, but you get the gist.

## Refactoring summary

- We went from twelve lines of code to one.

- It is way more readible and clear in what it does.

- The implementation is flexible enough to allow for major policy changes in the future.

- You no longer need a separate function call but evertyhing is neatly contained within the user.

- It is much more flexible with its shiny new configuration approach.

All in all, I am happy with this refactoring! Would you do something differently?
