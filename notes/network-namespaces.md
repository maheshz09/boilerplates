a# Network Namespaces

	Network namespaces are used by linux to provide isolation to docker conatiners, eg. if we ran an process inside containtainer only that process will be visible inside the container. no other host process or anything should no be visible at container. it has its seprate namespace and the services withing that namespace it could see;
	
lets an example. if you have nginx process running on your docker conatiner, the process id's of same process looks diffrent inside container and in the host machine.(PID)

when it comes to networking, our hosts has its interfaces that connects to LAN.
	our host has its own route and ARP table, when we create a new container it has its own namespace, own virual networs(interfaces) and virtual route table and arp table.

to create new network namespace on linux host.
# ip netns add red
# ip netns add blue
 
to list namespace that we created.
# ip nets 

to list interfaces on host.
# ip link

how do we run same ip link command withing namespaces that we created.
	prefix the command
# ip netns exec red ip link
	ip link command has ran inside red namespace

another method is....

# ip link -n red

so successfully we have restricted red namespace it to on our host namespaces.
---------------------------------

same with the arp table 
	if we ran command on host 
# arp
	we can see multiple entries.
	
but if we ran inside containers we can see no entries.
# ip netns exec red arp 

\\\ exactly same with the routing table.

# route
# ip netns exec red route 


so of now those network namespace's have no network connectivity, they dont have Network interfaces even.
and they can't see underlaying host n/w

establishing connectivity between namespace using virual peer or cable. same like just we connect two machines together with cable.

virtual cable at two end's

# ip link add veth-red type veth peer name veth-blue

attach thease network interface to approporiate namespace.

# ip link set veth-red netns red
		simlierly attach blue interface to blue namespace
		
# ip link set veth-blue netns blue

attach ip addresses to each namespace
# ip -n red addr add 192.168.15.1 dev veth-red
# ip -n blue addr add 192.168.15.2 dev veth-blue

up the link that we have created 
# ip -n red link set veth-red up
# ip -n blue link set veth-blue up

try to ping red namespace to blue namespace.
# ip netns exec red ping 192.168.15.2 

if we look at the arp table we can see from blue namespace the nabour is red and vice versa
# ip netns exec red arp 
# ip netns exec blue arp

if we see arp table of the host host is unaware of this things. about the interfaces and namespaces


-----

now we have configured for two namespace what if we have multiple Namespaces.
	just like physical network to communicate with multiple namespaces to communicate with each other.
	
to create virtual switch within our host and connect namespaces to it, there are multiple solutions availble like native  linux bridge or ovs openswitch


linux bridge:
to create internal bridge option we add a new interface to the host using the add command with the type set to bridge. will name it v-net-0

# ip link add v-net-0 type bridge

as far is our host is concened its just onther interface, just like eth0 interface.

# ip link

to up the interface

# ip link set dev v-net-0 up 
	so this interface act's as a switch for us and we can connect our interface to it. (as host) now we can connect our namesapces to that switch.
so we will be connecting all networks to bridge network.

to delete exesting links that we have created directly for red and blue ns 
# ip -n red link del veth-red
	when we delete one end the other end will deleted automatically 


so now create new link to connect to bridge network.
# ip link add veth-red type veth peer name veth-red-br
	so we have created same link as before but this time but the other end named veth-red-br as its connect to bridge network. this naming convention will help us easily identify the interfaces that associate to red namespace. 

similerly create a cable to connect the blue namespace to bridge network.
# ip link add veth-blue type veth peer name veth-blue-br

so our cables are ready we can directly attach to the namesapces.

so one end of this interface, to red namesapce.
# ip link set veth-red nets red

to connect other end to bridge network
# ip link set veth-red-br master v-net-0


-----

same with blue namespace...

# ip link set veth-blue netns blue
# ip link set veth-blue-br master v-net-0


------

let'us now set ip addresses for these links and turn them up.

# ip -n red addr add 192.168.15.1 dev veth-red
# ip -n blue addr add 192.168.15.2 dev veth-blue

finally turn devices up, now container can reach each other over the network.

# ip -n red link set veth-red up
# ip -n blue link set veth-blue up

remaing two namespaces to same network. we now have all four namespace connected to our internal bridge network
and they can all communicate with each other. they have all ip addresses 192.168.15.1 --- 192.168.15.4. we assing our host 192.168.1.2 to my host what if i tried to reach one of those interfaces in these namespace?

will it work... no ?

what if we have to establish connectivity between my host and these namespaces ?


the bridge switch that we have created it is actually network interface that connected to our host.
	so actually we do have interface on host 192.168.15. network on our host.
since its just interface we can assign ipaddress to it.
# ip addr add 192.168.15.5/24 dev v-net-0

we now able ping red namespace from localhost.
	remember this connection still in private.

from within namespaces we couldent reach outside world also no one can reach from ouside world can reach internal ips or applications etc.

so we have to configure our bridge network to reach lan network tru ethernet port
	eg. lets suppose there is another host named 192.168.1.3
so how can we reach this host from our namespace
# ip netns exec blue ping 192.168.1.3
	from our current network 192.168.15. 
so lets see its routing table 
# ip netns exec blue route
	the route table does not have any visbility other than our network 

so for that we need to provide gateway, how do we connect to our local gateway.
so how do we find that gateway...
	 a door or a gateway is a system on local Network that connects to other network, but do have any system that connects to our local namespace also ouside lan aswell.
its the localhost that has all namespaces on...

so we can ping the namespaces..
	our localhost has an interface to attach the private network so you can ping the namesapces
so our localhost is the network that connects to two networks together..

we can now add row entry in the blue namespace to say route all traffic to 192.168.1. network  to gateway at 192.168.15.5 now remember our host has two ip addresses on on the bridge network(internal n/w) 15.5 and external network 1.2
# ip netns exec blue ip route add 192.168.1.0/24 via 192.168.15.5

can we use any in the route ?

no because the blue namespace can only reach the gateway in its local network at 192.168.15.5. the default gateway should be reachable from your namesapce when you add it to your route.

while when we try to ping now we no longer get the unreachable message.

# ip netns exec blue ping 192.168.1.3


