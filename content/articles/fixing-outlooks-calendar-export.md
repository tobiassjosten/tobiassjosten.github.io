---
title: "Fixing Outlook's calendar export"
date: "2025-03-03"
---

I use a variety of calendars to organize my life, my family, and the many projects and businesses I'm involved with. Most of them use Google Workspace, but at work we use, for some unfathomable reason, Office 365.

<!--more-->

There's no way I could shoulder the myriad of responsibilities I'm charged with or coordinate the countless projects, businesses, and initiatives I'm involved with, unless I tracked it all in calendars.

It wouldn't be practical, however, if I had to jump around a dozen different calendars when planning my weeks, or just to be able to RSVP to an invitation. I have to be able to get an overview of all my various calendars, in one place.

Combining them has mostly been straightforward (whether that's in Google Calendar, Spark, or any [other tool I use](https://www.seastone.io/uses/)). I've been able to export and import all calendars, except for my work one, in Outlook. No matter where I've imported the Microsoft calendar, the events have randomly appeared at the wrong time. Some an hour late, others an hour early.

Suboptimal for a calendar, whose core function is to show events at specific times…

Not all its events have had the wrong time, though. Some were shown at the correct time, and some were not, but that only adds to the problem. How can you trust a calendar that acts more like a random number generator than an actual time ledger?

## Understanding iCalendar

Outlook exports its calendars, like most others, in [the iCalendar format](https://en.wikipedia.org/wiki/ICalendar?ref=seastone.io). Its full name is Internet Calendaring and Scheduling Core Object Specification and it looks like this:

```iCalendar
BEGIN:VCALENDAR
METHOD:PUBLISH
PRODID:Microsoft Exchange Server 2010
VERSION:2.0
BEGIN:VEVENT
DESCRIPTION:Let's align some synergies!
RRULE:FREQ=WEEKLY;UNTIL=20240610T080000Z;INTERVAL=4;BYDAY=MO;WKST=MO
UID:040000008200E00074C5B7101A82E00800000000D00254522634DA01000000000000000
 010000000A670AFB915FBCA4086DA8E3B2C6DC017
SUMMARY:No I in team
DTSTART;TZID=W. Europe Standard Time:20240219T100000
DTEND;TZID=W. Europe Standard Time:20240219T110000
CLASS:PUBLIC
PRIORITY:5
DTSTAMP:20250303T085411Z
TRANSP:OPAQUE
STATUS:CONFIRMED
SEQUENCE:0
LOCATION:Microsoft Teams-möte
END:VEVENT
END:VCALENDAR
```

The iCalendar standard has been around since 2009 and is described in [RFC 5545](https://datatracker.ietf.org/doc/html/rfc5545?ref=seastone.io) (with extension in [RFC 5546](https://datatracker.ietf.org/doc/html/rfc5546?ref=seastone.io), [6868](https://datatracker.ietf.org/doc/html/rfc6868?ref=seastone.io), [7529](https://datatracker.ietf.org/doc/html/rfc7529?ref=seastone.io), and [7986](https://datatracker.ietf.org/doc/html/rfc7986?ref=seastone.io)).

RFCs are great! They leave very little room for interpretation and enable a high degree of interoperability. While they might look daunting at first glance, reading them isn't that bad, and doing so carefully is essential if you're implementing the standard in question.

_(Spoiler: If Microsoft had done so, you wouldn't be reading this article…)_

I live in Stockholm, Sweden, in the CET or Europe/Stockholm time zone. We use daylight saving time, so during winter, we're UTC+1, and during summer, we're UTC+2. Thus, time zone problems were highly suspicious with my events being an hour early or late.

Skimming through RFC 5545, focusing on time zones, one thing stood out: It kept referring to IANA, in contrast to "non-standard values". Here, my tech intuition smelled Microsoft!

This suspicion had me switch over to the calendar export from Outlook, where I soon noticed something weird. Among the many different time zones used (which was weird in itself), I saw for the first time one called "Romance Standard Time".

![Outlook export using "Romance Standard Time" time zone](https://www.seastone.io/content/images/2025/03/Screenshot-2025-03-02-at-14.41.23.png)
*'Nothing "Standard" about that!'*

This was not a time zone I'd ever seen before and I've been working with time for decades, specifically in this part of the world. So I began suspecting it must be one of those "non-standard" ones referenced in the RFC.

What are those anyway?

## Time zone standards

Referenced in the RFC, IANA (Internet Assigned Numbers Authority) is an organization that maintains a list of time zones called [tzdata](https://www.iana.org/time-zones?ref=seastone.io). With its five hundred entries, it's the de facto standard for time zones, used by most operating systems, programming languages, and other major systems.

"Romance Standard Time" is not one of those five hundred. So where did it come from? It couldn't be that Microsoft made up their own time zones? Right?

Wrong. That's [exactly what they did](https://learn.microsoft.com/en-us/windows-hardware/manufacture/desktop/default-time-zones?view=windows-11&ref=seastone.io#time-zones). 🤦‍♂️

I found an excellent write-up about these time zones, called "[Understanding Time Zones](https://medium.com/@gonzaloohk/understanding-time-zones-part-1-3469a6327905?ref=seastone.io)", by Gonzalo Osco Hernández. He explains the Microsoft custom database:

> Windows time zones are maintained by Microsoft and are updated much less frequently than IANA time zones. There are currently 107 Microsoft time zones.  
>   
> IANA time zone names follow a standard convention, while Microsoft time zones id seems to have no conventions.  
>   
> IANA implements its time zones based on reported information from what is actually in effect and used by people in the region. Microsoft implements only official time zones that are set by legislation or other official government policy.  
>   
> Microsoft time zones can only implement two DST transitions in a given year, so they have problems representing some real-world scenarios like Egypt and Morocco DST in 2010 and 2013.

So not only did they deviate from the established standards with their own variant but they did so with a worse alternative. (Not unlike many other Microsoft products, I'll admit.)

Calendar exports from Outlook work well with other Microsoft products but nothing else. It's effectively an artificial vendor lock-in, forcing you to stay within their ecosystem.

If I were conspiratorially inclined, I'd suspect that their departure from both the standard time zones and the RFC standard was anything but accidental. But I'm a techie, more interested in solving problems and minimizing wrestling whatever comes out of Redmond.

## How Microsoft could solve it

RFC 5545 does allow for arbitrary time zones, using any name the publisher fancies, as long as they are declared. Your events could be in the time zone "Blue", "Donald Duck", or even "Roman Standard Time" but they must have a definition.

From the RFC:

> Parameter Name: TZID  
> Description: \[…\]This property parameter specifies a text value that uniquely identifies the "VTIMEZONE" calendar component to be used when evaluating the time portion of the property. The value of the "TZID" property parameter will be equal to the value of the "TZID" property for the matching time zone definition. An individual "VTIMEZONE" calendar component MUST be specified for each unique "TZID" parameter value specified in the iCalendar object.  
> [RFC 5545 3.2.19](https://www.rfc-editor.org/rfc/rfc5545?ref=seastone.io#section-3.2.19)

> Component Name: VTIMEZONE  
> An individual "VTIMEZONE" calendar component MUST be specified for each unique "TZID" parameter value specified in the iCalendar object. \[…\]  
> [RFC 5545 3.6.5](https://www.rfc-editor.org/rfc/rfc5545?ref=seastone.io#section-3.6.5)

If Microsoft wanted to follow the standard and make their exports interoperable with other systems, they could simply add blocks like the following to their exports:

```iCalendar
BEGIN:VTIMEZONE
TZID:Romance Standard Time
BEGIN:DAYLIGHT
TZOFFSETFROM:+0100
TZOFFSETTO:+0200
DTSTART:19810329T020000
RRULE:FREQ=YEARLY;BYMONTH=3;BYDAY=-1SU
TZNAME:CEST
END:DAYLIGHT
BEGIN:STANDARD
TZOFFSETFROM:+0200
TZOFFSETTO:+0100
DTSTART:19961027T030000
RRULE:FREQ=YEARLY;BYMONTH=10;BYDAY=-1SU
TZNAME:CET
END:STANDARD
END:VTIMEZONE
```

Sample VTIMEZONE block

This would allow others to interpret these non-standard time zones correctly.

Curiously, Microsoft does add a few VTIMEZONE blocks, so I'm not sure why they don't do so for all referenced TZIDs. (There's one for UTC, which the RFC explicitly disallows.)

Thankfully, there are workarounds and ways to make sense of the exports. [The Unicode Consortium](https://home.unicode.org/?ref=seastone.io) runs an initiative called [Unicode CLDR Project](https://cldr.unicode.org/?ref=seastone.io), which publishes [mapping data](https://github.com/unicode-org/cldr-json/blob/main/cldr-json/cldr-core/supplemental/windowsZones.json?ref=seastone.io), translating Microsoft time zones to standard ones. This allows you to parse an iCalendar export and replace the non-standard time zones, so that it can be used in other systems.

## Introducing Calendzo

I couldn't wait for an official solution from Microsoft (nor do I really expect one). Instead I created a tiny service to do this time zone rewriting, so I could fix the problem and finally gather all my calendars in one place.

Say hello to [Calendzo](https://www.calendzo.com/?ref=seastone.io)!

It's basically a proxy between Outlook and your client, rewriting any iCalendar stream to replace Microsoft time zones with standard ones. As simple as it is effective.

![Flowchart showing how Calendzo intercepts and fixes crappy Outlook exports](https://www.seastone.io/content/images/2025/03/calendzo-overview-1.png)
*Outlook exports crap, Calendzo fixes it, systems can interoperate.*

It's 100% free and if it helps someone else, I'd love to hear about it!
