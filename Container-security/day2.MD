2. Container security fundamantals

	improve linux hardening process, we can able drop some of the capbilties of linux to improve the security of containers.
	
practicle...

	so lets spawn one conatiner and we can try to limit some of the capbilties
	
#docker run -it raesene/alpine-containertools /bin/bash

root@local:~# docker run -it raesene/alpine-containertools /bin/bash
	Unable to find image 'raesene/alpine-containertools:latest' locally
	latest: Pulling from raesene/alpine-containertools
	0a9a5dfd008f: Pulling fs layer
	b5613c270156: Pulling fs layer
	b5613c270156: Downloading [=============>                                     ]  4.678MB/17.15MB
	d904dd0c87f6: Waiting


so with the below command we can check what kind of capablilty that we go by default by the conatainer OR in docker in genral.

#amicontained
	6505b5eb0900:/# amicontained
	Container Runtime: docker
	Has Namespaces:
			pid: true
			user: false
	AppArmor Profile: unconfined
	Capabilities:
			BOUNDING -> chown dac_override fowner fsetid kill setgid setuid setpcap net_bind_service net_raw sys_chroot mknod audit_write setfcap
	Seccomp: filtering
	Blocked Syscalls (54):
			MSGRCV SYSLOG SETSID USELIB USTAT SYSFS VHANGUP PIVOT_ROOT _SYSCTL ACCT SETTIMEOFDAY MOUNT UMOUNT2 SWAPON SWAPOFF REBOOT SETHOSTNAME SETDOMAINNAME IOPL IOPERM CREATE_MODULE INIT_MODULE DELETE_MODULE GET_KERNEL_SYMS QUERY_MODULE QUOTACTL NFSSERVCTL GETPMSG PUTPMSG AFS_SYSCALL TUXCALL SECURITY LOOKUP_DCOOKIE CLOCK_SETTIME VSERVER MBIND SET_MEMPOLICY GET_MEMPOLICY KEXEC_LOAD ADD_KEY REQUEST_KEY KEYCTL MIGRATE_PAGES UNSHARE MOVE_PAGES PERF_EVENT_OPEN FANOTIFY_INIT OPEN_BY_HANDLE_AT SETNS KCMP FINIT_MODULE KEXEC_FILE_LOAD BPF USERFAULTFD
	Looking for Docker.sock
	6505b5eb0900:/#



so in the "capablilties" section we can see there are some capbilities that we got by default.

so when you run a container might you want to drop all the privilages, lets suppose if we have an application that treditionally ran as non-root user before you use containers, might you dont need that capablilties at all

example.

#docker run --cap-drop=all raesene/alpine-containertools /bin/bash

	with the --cap-drop  we have removed all the capablilties frrom the container, so lets check.
	
		root@local:~# docker run -it --cap-drop=all raesene/alpine-containertools /bin/bash
		38067276db49:/# amicontained
		Container Runtime: docker
		Has Namespaces:
				pid: true
				user: false
		AppArmor Profile: unconfined
		Capabilities:
		Seccomp: filtering
		Blocked Syscalls (56):
				MSGRCV SYSLOG SETSID SETGROUPS USELIB USTAT SYSFS VHANGUP PIVOT_ROOT _SYSCTL CHROOT ACCT SETTIMEOFDAY MOUNT UMOUNT2 SWAPON SWAPOFF REBOOT SETHOSTNAME SETDOMAINNAME IOPL IOPERM CREATE_MODULE INIT_MODULE DELETE_MODULE GET_KERNEL_SYMS QUERY_MODULE QUOTACTL NFSSERVCTL GETPMSG PUTPMSG AFS_SYSCALL TUXCALL SECURITY LOOKUP_DCOOKIE CLOCK_SETTIME VSERVER MBIND SET_MEMPOLICY GET_MEMPOLICY KEXEC_LOAD ADD_KEY REQUEST_KEY KEYCTL MIGRATE_PAGES UNSHARE MOVE_PAGES PERF_EVENT_OPEN FANOTIFY_INIT OPEN_BY_HANDLE_AT SETNS KCMP FINIT_MODULE KEXEC_FILE_LOAD BPF USERFAULTFD
		Looking for Docker.sock
		38067276db49:/#


so as we can see there is cap-drop worked, we can  able to perform all kind of operations that normal user can able to do 
but we couldent not able to do root lavel operations 


so we should do that when we want to run our appplication inside container.

	due to changes in way how the linux works over years, we actually have capabilities granted by default to every conatiner
	that we can get rid of completely and safely.

example.

1 -- netraw
	it use to send raw packets like process. and this kind of security standpoint seems dangrous becasue it allows you to do 
	packet spoofing attacks, so might attackers can create fake traffic for process for host. so this is the one we should try to disable. so the reson behind why conainer get that permissions by defalt they needed for "PING". ping treditionally uses this Capabilities in order to send and recive icmp packets. 
	
	eg. 
	# docker run -it --cap-drop=net_raw raesene/sec_test /bin/bash
		root@local:~# docker run -it --cap-drop=net_raw raesene/sec_test /bin/bash
						Unable to find image 'raesene/sec_test:latest' locally
						latest: Pulling from raesene/sec_test
						405f018f9d1d: Downloading [==============>                                    ]  8.714MB/30.42MB
						6b04427bfa7e: Downloading [=>  
						
		root@0e4fd02a74d2:/# ping 8.8.8.8
			PING 8.8.8.8 (8.8.8.8) 56(84) bytes of data.
			64 bytes from 8.8.8.8: icmp_seq=1 ttl=117 time=7.48 ms
			64 bytes from 8.8.8.8: icmp_seq=2 ttl=117 time=5.74 ms

		
		and it works fine, so reson is that linux exposes a setting specifically for "icmp echo" and it says who allows to send this whithout any privilages or Capabilities or whithout being root user on the host.
		
		root@0e4fd02a74d2:/#  sysctl net.ipv4.ping_group_range
			net.ipv4.ping_group_range = 0   2147483647
			root@0e4fd02a74d2:/#
			
		so as we can see its like a fullrange that means amy group can send this on the host. mean to anyone can send icmp echo messages. so we no longer need "net_raw". so only reson you hade it beacuse of ping so you can drop it off from cap-drop
		
thats one option to harden

2 -- bind privilaged ports in linux.

	so another common when pepople need capbilities to bind ports lower than 1024. this considerd as privilaged ports in linux and unix. so like port  80,443 are very common that might you need to bind them to your container. so eg. if i try to run below commnd witch essentally bind server to port 80
	
	# python3 -m http.server 80 (on host machine)
		kiosk@local:/root$ python3 -m http.server 80
		Traceback (most recent call last):
		  File "/usr/lib/python3.10/runpy.py", line 196, in _run_module_as_main
			return _run_code(code, main_globals, None,
		  File "/usr/lib/python3.10/runpy.py", line 86, in _run_code
			exec(code, run_globals)
		  File "/usr/lib/python3.10/http/server.py", line 1307, in <module>
			test(
		  File "/usr/lib/python3.10/http/server.py", line 1258, in test
			with ServerClass(addr, HandlerClass) as httpd:
		  File "/usr/lib/python3.10/socketserver.py", line 452, in __init__
			self.server_bind()
		  File "/usr/lib/python3.10/http/server.py", line 1301, in server_bind
			return super().server_bind()
		  File "/usr/lib/python3.10/http/server.py", line 137, in server_bind
			socketserver.TCPServer.server_bind(self)
		  File "/usr/lib/python3.10/socketserver.py", line 466, in server_bind
			self.socket.bind(self.server_address)
		PermissionError: [Errno 13] Permission denied
kiosk@local:/root$

so it essentally says you're not a root user. so basically mean you dont have net_bind service Capabilitity therefore i am not let you do this.

so however if try to run this same command inside container. it works just fine.

	# docker run -it --cap-drop=net_bind_service raesene/sec_test /bin/bash
		root@7bb4ac05dbae:/# python3 -m http.server 80
			Serving HTTP on 0.0.0.0 port 80 (http://0.0.0.0:80/) .
	
	
so again there is system control setting on the linux, so that allows you to change whare the privilaged port range. and explans the unprivilaged port range that anyone can can bind.

	so what dockerhub have done default there no specific contxt of privilaged port as an idea, so what they have done they set unprivilaged port range start at 0 so essentally means anyone can bind port's inside the container (conainer network namespace) so that if you need that Capabilities to bind you can drop that Capabilities.
	
	eg.
		root@7bb4ac05dbae:/# sysctl net.ipv4.ip_unprivileged_port_start
			net.ipv4.ip_unprivileged_port_start = 0
			root@7bb4ac05dbae:/#
			
	so on host you can see the setting.
		
		root@local:~# sysctl net.ipv4.ip_unprivileged_port_start
          net.ipv4.ip_unprivileged_port_start = 1024

so this treditionally setting present on the host becasue of witch we are not able to bind port 80 at that moment.


ref:
	link: https://www.youtube.com/watch?v=EFOA2nCZ0gg
