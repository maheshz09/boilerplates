# Switching

to get the current interfaces in the machine
# ip link
    kiosk@local:~$ ip link                                                                                                  1: lo: <LOOPBACK,UP,LOWER_UP> mtu 65536 qdisc noqueue state UNKNOWN mode DEFAULT group default qlen 1000                    link/loopback 00:00:00:00:00:00 brd 00:00:00:00:00:00                                                               2: eth0: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc mq state UP mode DEFAULT group default qlen 1000                  link/ether 00:15:5d:62:8d:04 brd ff:ff:ff:ff:ff:ff                                                                  3: docker0: <NO-CARRIER,BROADCAST,MULTICAST,UP> mtu 1500 qdisc noqueue state DOWN mode DEFAULT group default                link/ether 02:42:84:16:99:fe brd ff:ff:ff:ff:ff:ff                                                                  kiosk@local:~$  

to add ip address to to interface suppose and eth0
# ip addr add 192.168.1.10/24 dev eth0

so to look into the exesting route tables & and the gateways defined in the machine we can see using
# route
    kiosk@local:~$ route                                                                                                    Kernel IP routing table                                                                                                 Destination     Gateway         Genmask         Flags Metric Ref    Use Iface                                           default         local.mshome.ne 0.0.0.0         UG    0      0        0 eth0                                            172.17.0.0      0.0.0.0         255.255.0.0     U     0      0        0 docker0                                         192.168.48.0    0.0.0.0         255.255.240.0   U     0      0        0 eth0                                            kiosk@local:~$                                                                                                                        
to configure a gateway to a diffrent network range like 192.168.10.0 then? 

# ip route add 192.168.10.0/24 via 192.168.48.0
    so basically we are routing OR reach the traffic of "10.0" network from and exesting gateway of "48.0"
    this will be system specific, we need to configure for every system that we need to communicate with "10.0" network.

like same eg.

# ip route add 192.168.1.0/24 via 192.168.2.0

suppose this systems needs access to internet then ??
    suppose they need access to google.com at 172.217.194.0
so before that we connect our router to internet and then add route in our machine.

# ip route add 172.217.194.0/24 via 192.168.2.1

there are so many site's on the diffrent network on the internet, insted of adding routing table entry for same routers ip address for each of those networks we can say, use this router as default gateway the network range you dont know

# ip route add default via 192.168.2.1
    this way any request to any network outside of our exesting network goes to perticular router.


EG. 

basically setting up manual routing tables to linux host.
suppose we have two networks, 

<MACHING-1>   ------> 192.168.1.0 -----> <MACHING-2><2-interfaces>         ----------> 192.168.2.0 <MACHING-3>
<192.168.2.5>                            <192.168.1.6> --> <192.168.2.6>                           <192.168.2.5>

so we have two network's as you can see 192.168.1.0 and 192.168.2.0 respectively
    so the machine-2 in our case can able to commuicate with both networks

so suppose and example, if we want to connect machine-1 can able to commuicate with machine-3 it not able to do because it dont have routes to connect to it.
    genrally it will say "Network is Unreachable"

host-1 has not idea how to reach host-3.
so in our case the gateway of one network to another is our machine-2 it can able to communicate with both the network's... 
we can add route to access

# ip route add 192.168.2.0/24 via 192.168.1.6 

tru this you can reach the machine-3 but if we want to get response back from machine-3 then what ?? we need to add another route for machine-3 to communicate back to machine one with same principal

# ip route add 192.168.1.0/24 via 192.168.2.6

when we try to communicate now we no longer getting warning of unable to reach network or anything. that means our routing entries are right
    but we still dont get response back, by default in linux packets are not forwared from one interface to other. this is implemnted due to security reson.
eg. your one interface connected to your private network(eth0) and another to a public network(eth01). we dont want to anyone from public network to easly send messages to our private network.
    unless we explicitly allow that, so in this case we wanted to expose the network so what we do. 
	
so now we can allow our machine-2 to send packetes from one interface to another.

# cat /proc/sys/net/ipv4/ip_forward
 by DEFAULT the value in this file is 0 that means no forward.

if we change it to 1 and you should see the ping will go thru.
# echo 1 > /proc/sys/net/ipv4/ip_forward

those change are no persistent, we should modify the same value in 

# /etc/sysctl.conf
	net.ipv4.ip_forward = 1
	
