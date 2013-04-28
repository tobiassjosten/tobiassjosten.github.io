---
layout: post
title: Changing your MAC address
category: security
tags: [security]
summary: Say that you are, purely theoretically, sitting at an airport in France after a great Symfony Live 2012 Paris. This hypothetical airport would give you 15 min free WiFi and then require you to pay to keep using the Internet connection. If we wanted to continue this thought experiment we could try and hack the airport network security.
---
Say that you are, purely theoretically, sitting at an airport in France after a great [Symfony Live 2012 Paris](/symfony/symfony-live-2012-paris). This hypothetical airport would give you 15 min free WiFi and then require you to pay to keep using the Internet connection.

You do not have to give any information to start using this free time, which indicates that they are using *MAC addresses* to make sure you get the 15 min but no more than that. If we wanted to continue this thought experiment by trying to hack the security, we would first check our current MAC address by using the `ip addr` command.

    $ sudo ip addr
    ...
    3: wlan0: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc mq state UP qlen 1000
        link/ether 33:27:2f:11:1e:01 brd ff:ff:ff:ff:ff:ff
        inet 10.240.26.121/16 brd 10.240.255.255 scope global wlan0
        inet6 fe80::8a53:2eff:fe1b:c0be/64 scope link 
           valid_lft forever preferred_lft forever

Our WLAN card has *33:27:2f:11:1e:01* for its MAC address. We could try and change that to an arbitrary address.

    $ sudo ifconfig wlan0 down
    $ sudo ifconfig wlan0 hw ether 11:22:33:44:55:66
    SIOCSIFHWADDR: Cannot assign requested address

The network card does not want to play along. This is because of a limit imposed by the manifacturer. But in these cases they still allow you to change the least significant part of the address, the number furthest to the right.

So we can try to increase our address by 1.

    $ sudo ifconfig wlan0 down
    $ sudo ifconfig wlan0 hw ether 33:27:2f:11:1e:02

No complaints! In *nix that usually means it worked, but let us double check that.

    $ sudo ip addr
    ...
    3: wlan0: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc mq state UP qlen 1000
        link/ether 33:27:2f:11:1e:02 brd ff:ff:ff:ff:ff:ff
        inet 10.240.26.121/16 brd 10.240.255.255 scope global wlan0
        inet6 fe80::8a53:2eff:fe1b:c0be/64 scope link 
           valid_lft forever preferred_lft forever

Voila, c'est tres bien! On Ubuntu you would now automatically reconnect to the network and it would not recognize you. Had this not been purely food for thought, one could now be enjoying another fifteen minutes of free internet.
