---
layout: post
title: Doctrine migrations gotchas in Symfony
category: symfony
tags: [symfony, doctrine]
summary: After having wrestled with Doctrine migrations I now have working code and my data is safe. But before I go on with my life I wanted to take a minute and post my findings, with hopes of helping future migraters.
---
After having [wrestled](http://twitter.com/tobiassjosten/status/33172676159864832) with Doctrine migrations I now have working code and my data is safe. But before I go on with my life I wanted to take a minute and post my findings, with hopes of helping future migraters.

If you are a first time user of *Doctrine migrations*, I really recommend reading [Dennis Benkert](http://www.denderello.com/)'s excellent [writeup about it](https://github.com/denderello/guide-to-doctrine-migrations). It really goes through all the basics you need to know and so I will only post a few gotchas here.

## Bash script for speed

It's helpful to setup a reset bash script, so you can quickly test your changes to the generated migration scripts.

    # REPLICATE PRODUCTION
    php symfony doctrine:drop-db --no-confirmation
    php symfony doctrine:build-db
    mysql -ua -pb c<dbdump.sql

    # RESET MIGRATIONS
    mysql -ua -pb c -e 'DROP TABLE IF EXISTS `migration_version`'
    mysql -ua -pb c -e 'CREATE TABLE `migration_version` (`version` int(11) DEFAULT NULL) ENGINE=InnoDB DEFAULT CHARSET=latin1'
    mysql -ua -pb c -e 'INSERT INTO `migration_version` VALUES (0)'

Replace *a* with the username, *b* with the password and *c* with the database name. Save it as *pre-migrate.sh*, run *chmod +x pre-migrate.sh* to make it executable and then run it with *./pre-migrate.sh*.

## Order is everything

If you read through Denderello's blog post you will notice that he promotes a certain order of doing things. The importance of this order can not be overstated!

1) First you enable migrations by creating the above *migration_version* table, starting with 0 for its one *version* value. This is only done the first time you start using Doctrine migrations.

2) Make your changes to `schema.yml`. If you are overriding or extending the model definitions of plugins, be sure to delete their classes from *lib/model/doctrine/sfWhateverPlugin* and the equivalent form and filter directories.

3) Generate the migration code with `php symfony doctrine:generate-migrations-diff`.

4) Run `php symfony doctrine:migrate` and your data is migrated! Then run `php symfony doctrine:build --all-classes` to also bring the classes up to date.

## Table race conditions

    - SQLSTATE[23000]: Integrity constraint violation: 1217 Cannot delete or update a parent row: a foreign key constraint fails. Failing Query: "DROP TABLE group_table"

If you have done everything correctly and this still shows up, you probably have a `race condition` in your migration code. First check your foreign key constraints with this mysqldump trick:

    mysqldump -ua -pb c | grep -i foreign

Then go and edit your migrations. Look for `dropTable()` calls and see if they are in the correct logical order. Most likely they are not. Fix it and your your speedy migration bash script again.

## Foreign key constraints

    - SQLSTATE[HY000]: General error: 1005 Can't create table 'project.#sql-56d_619' (errno: 121). Failing Query: "ALTER TABLE sf_guard_user_challenge ADD CONSTRAINT sf_guard_user_challenge_sf_guard_user_id_sf_guard_user_id FOREIGN KEY (sf_guard_user_id) REFERENCES sf_guard_user(id)"

Be sure to check that all model relations uses the same datatypes for their matching columns. If category.id is a BIGINT then blogpost.category_id must also be a BIGINT.

My particular issue was that [sfDoctrineGuardPlugin](http://www.symfony-project.org/plugins/sfDoctrineGuardPlugin) had changed to BIGINT for all its IDs. This is a good change but not paying attention to it cost me some time.

## Moving data

My main objective was to throw out [sfFacebookConnectPlugin](http://www.symfony-project.org/plugins/sfFacebookConnectPlugin), in favor of [sfMelodyPlugin](http://www.symfony-project.org/plugins/sfMelodyPlugin). The former promoted putting your Facebook IDs in a secondary profile table but now I wanted them directly in my sfGuardUser objects.

For some reason Doctrine generates the destructive `dropTable()` and `addColumn()` calls before the constructive `createTable()` and `removeColumn()` ones. Since I needed data from the profile table, before dropping it, I switched order for the methods calls.

However this was not enough. It seems Doctrine somehow delays the execution of your database changes. When I tried to put data into my newly created facebook_id column Doctrine said it didn't exist. When I moved the code to another, subsequent migration method, it worked as intended.

The actual code to move my data:

    Doctrine_Manager::getInstance()
      ->getCurrentConnection()
      ->execute('UPDATE sf_guard_user u, sf_guard_user_profile p SET u.facebook_id = p.facebook_uid WHERE u.id = p.user_id');

So there you have it, a few gotchas when using Doctrine migrations!
