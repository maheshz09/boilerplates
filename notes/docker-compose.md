version: 3.8
services:
  samba:
    image: dperson/samba
    restart: always
    container_name: samba_server
    ports:
      - "137:137/udp"
      - "138:138/udp"
      - "139:139/tcp"
      - "445:445/tcp"
    volumes:
      - /root/data/shared:/mount
    environment:
      - USER=mahesh;password  # Format: username;password
      - SHARE=public;/mount;yes;no;yes;all;mahesh
    command: >
      -p  # persist user accounts
      -s "public;/mount;yes;no;yes;all;mahesh"