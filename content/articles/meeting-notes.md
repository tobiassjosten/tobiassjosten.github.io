---
title: "Meeting Notes"
date: "2024-09-16"
---

I do a lot of meetings. Some more useful than others. A sure way to make them needlessly useless, however, is to disregard my notes for them; either the ones I write in order to prepare, or the ones I write during the meeting.

<!--more-->

I have previously shared about [how I write to learn](https://www.seastone.io/writing-learning-growing/) and the same principle holds true for meetings — the act of taking notes creates a framework for thinking, which makes for effective meetings.

## My method

There are countless methods for taking notes and I've tried most of them. None has really stuck, though. Either they have been too complicated, so I've inevitably abandoned them with time, or they have relied on pen and paper and I'm not a Luddite.

Taking notes has to be quick and simple. The notes also have to compress a lot of information, so as to enable me to jot down a few words in summary and stay present in the meeting.

In response to these requirements, I have "accidentally" developed my own notations for taking quick meeting notes. In hindsight, I'm sure they are influenced by [The Bullet Journal Method](https://www.seastone.io/the-bullet-journal-method/) (which I liked but no longer use).

_Practically, I write and keep my notes with_ [_Bear_](https://bear.app/)_, where I also store most content of my productivity system. That enables me to refer both to the meeting from other notes and from the meeting notes to specific resources, tying everything together._

## My notations

I use a few different signs to make some notes stand out.

### The bullet — agenda items

```
- This is something I want to bring up or a question I want to ask.
```

In most Markdown editor, these bullet items stand out and makes it easy for me to remember everything I wanted to talk about.

I add these both during and before the meeting. It's how I build my agenda, by creating the meeting note as soon as I book the meeting, so that I can add to it as thoughts pop up.

### The exclamation marks — action items

It's important to distinguish between what was said in general and what we decided needs to happen next, which is why I use exclamation marks for action items.

```
! This is something I have to do.
```

I make a difference between action points for myself and others by using one or two exclamation marks, and the responsible party's name if needed. It doesn't make sense during one-on-ones, for example.

```
!! Frank: This is something that my colleague has to do.
```

```
!! Look into that thing until next check-in.
```

Together, these two notations make it very easy to summarize the meeting and agree on who does what next.

A huge update to my productivity system has been to not only track what I have to do but what I expect others to do as well. Trust but verify, as it goes.

Earlier, I have naively trusted that others either do what we agreed they would or get back to me in time, if they can't complete their task for some reason. There's a difference between culpability and responsibility and the sooner you can catch problems, the easier they are to solve.

### The question mark — curiosities

```
? This is a take-away I want to look up or learn more about.
```

It could be a concept someone mentions, a new technology, a movie recommendation, or anything I want to follow up on for my own sake.

Separating `?` from `!` makes it easier to summarize everyone's action items and for me to distinguish between what I have to do and what I would like to do.

### The parentheses — unrelated thoughts

Sometimes, my mind wanders or I think of something unrelated to the meeting. To avoid cluttering the notes, and to maintain the ability to easily summarize my notes, I put these random thoughts in parentheses.

```
) It's been a while since I bought flowers for my wife.
```

I make sure to quickly clear these out when processing my notes later, so I tend to skip adding too much context here. If I want to make sure I remember to do something, I make it an unrealted action item instead.

```
!) Buy some flowers for wifey on the way home.
```

## Processing my notes

An often overlooked, but hugely important, part of taking notes is processing them. Unless you review them soon after the meeting, you might as well not waste the time taking them to begin with.

I try to leave myself about five minutes after each meeting, just for this reason; to sit down and process my notes. Worst case, I do it at the end of the day. Any longer than that and I find myself wondering what half of the notes mean.

This quick turnaround helps me write even shorter notes and stay present in the meeting.

One crucial part of processing notes is to throw away anything that's irrelevant. Sometimes, some things feels more important in the moment when you talk about them, than when you consider them with some distance.

I view my notes as something I prepare for my future self and so I remove anything that's not needed or risk confusing myself in some months. Often, this actually means I will delete the entire meeting document (once I've transferred the action items to [Todoist](https://todoist.com/)).

Notes are meant to be used, not merely to document.

## Example

This is what a note can look like before a meeting:

```
- Load testing. Are we confident?  

- Resources. More developers?
```

This is my agenda, with short notes just to remember what I want to bring up.

Here is how the note transforms during the meeting:

```
load testing done, 10x traffic, autoscaling enabled  

short on frontend  
! check budget, consultant  

- architecture
```

As you can see, I take notes all in lowercase. This makes it easier to know what still needs to be processed.

Finally, here's the note once processed:

> Team has carried out load tests and made sure that the application handles ten times the traffic we currently serve. Further, autoscaling is configured to allow us to handle spikes in traffic.  
>   
> ! Check with platform to make sure we have a limit on our cloud spendings.  
>   
> PM is worried we might not be able to deliver on the next milestone in time, since our frontend capacity is a bottleneck at the moment.  
> ! Check with X to see if we can hire a consultant to help the team or adjust the roadmap accordingly.  
>   
> They have implemented a pretty neat event architecture, which our other teams might benefit quite a bit from.  
> ! Help X prepare a presentation on the architecture for the other devs.

A little more fleshed out, so that I can refer back to this in some months or years, along with concrete action items (transferred to Todoist but saved for posterity).
