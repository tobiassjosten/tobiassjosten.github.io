---
layout: post
title: Update associated entities in Doctrine
category: symfony
tags: [symfony, doctrine]
summary: I am building a website for tv-series (Swedish) using Symfony2. Recently I ran into a problem in how to manage my data with Doctrine, which solution took some reading up on the Unit of Work part of the ORM.
---
I am building a [website for tv-series](http://www.smartburk.se/) (Swedish) using [Symfony2](/symfony) and recently I ran into a problem in how to manage my data with [Doctrine](/doctrine). The solution took some reading up on the *Unit of Work* part of the ORM and so I wanted to share it.

My entities are Series, which has many Seasons, which have many Episodes. When an episode is created or updated I want to also update its related Series so that it always has correct firstYear and lastYear data for when the series was broadcasted.

## Doctrine lifecycle callbacks

I started out using [lifecycle events](http://docs.doctrine-project.org/en/latest/reference/events.html#lifecycle-events).

    /**
     * @ORM\HasLifecycleCallbacks
     */
    class Episode
    {
        /**
         * @ORM\PrePersist()
         * @ORM\PreUpdate()
         * @ORM\PostPersist()
         * @ORM\PostUpdate()
         */
        public function updateYearSpan()
        {
            $series = $this->getSeason()->getSeries();
            $series->setFirstYear($this->getAirDate());
            $series->setLastYear($this->getAirDate());
        }
    }

It was triggered but the changes made to the series were not persisted to the database. Since you do not have access to the entity manager in these callbacks, there is furthermore no way to force a persist/flush from here.

Back to the drawing board.

## Doctrine event listeners

For my next attempt I read up more carefully on [how events are handled](https://doctrine-orm.readthedocs.org/en/latest/reference/events.html) within Doctrine. What I was looking for seemed to be a full-fledged event listener for the *onFlush* event.

I opened up my bundle's *services.yml* and added a new listener. This example is using YAML, so have a look [in the manual](http://symfony.com/doc/current/book/service_container.html#creating-configuring-services-in-the-container) for examples in other formats that you might us.

    smartburk_main.episode_listener:
        class: Smartburk\Bundle\MainBundle\Listener\EpisodeListener
        tags:
            - { name: doctrine.event_listener, event: onFlush }

Next I created the actual listener.

    <?php

    // src/Smartburk/Bundle/MainBundle/Listener/EpisodeListener.php

    namespace Smartburk\Bundle\MainBundle\Listener;

    use Doctrine\ORM\Event\OnFlushEventArgs;

    class EpisodeListener
    {
        public function onFlush(OnFlushEventArgs $args)
        {
        }
    }

Now what happens is that when *EntityManager#flush()* is executed, my listener callback is run and I can modify the changeset before it is being sent to the database.

## Updating the entities changeset

Now that we are hooked into the right place, we can make our changes. First we pick the changing entities we are interested in and start iterating over them.

    $entities = array_merge(
        $uow->getScheduledEntityInsertions(),
        $uow->getScheduledEntityUpdates()
    );

    foreach ($entities as $entity) {
        if (!($entity instanceof Episode)) {
            continue;
        }

        // Make changes.
    }

To make the actual changes we need to trigger a re-computation of the changeset of the affected entities. We use the *Unit of Work* to modify this changeset, through [recomputeSingleEntityChangeSet()](https://github.com/doctrine/doctrine2/blob/master/lib/Doctrine/ORM/UnitOfWork.php#L903).

    $em = $args->getEntityManager();
    $uow = $em->getUnitOfWork();

    $series = $entity->getSeason()->getSeries();
    $series->setFirstYear($entity->getAirDate());
    $series->setLastYear($entity->getAirDate());

    $em->persist($series);
    $md = $em->getClassMetadata('Smartburk\Bundle\MainBundle\Entity\Series');
    $uow->recomputeSingleEntityChangeSet($md, $series);

Here is the complete example of how to use the Unit of Work to modify the changeset before Doctrine persists the data.

    public function onFlush(OnFlushEventArgs $args)
    {
        $em = $args->getEntityManager();
        $uow = $em->getUnitOfWork();

        $entities = array_merge(
            $uow->getScheduledEntityInsertions(),
            $uow->getScheduledEntityUpdates()
        );

        foreach ($entities as $entity) {
            if (!($entity instanceof Episode)) {
                continue;
            }

            $series = $entity->getSeason()->getSeries();
            $series->setFirstYear($entity->getAirDate());
            $series->setLastYear($entity->getAirDate());

            $em->persist($series);
            $md = $em->getClassMetadata('Smartburk\Bundle\MainBundle\Entity\Series');
            $uow->recomputeSingleEntityChangeSet($md, $series);
        }
    }

I am only looking through the entities that have been changed or updated. If you are after something else, [see the documentation](https://doctrine-orm.readthedocs.org/en/latest/reference/events.html#onflush) for what else is accessible.
