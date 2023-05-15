
`ip addr` result:

```
24: mac-1@wlp4s0: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc noqueue state UP group default qlen 1000
    link/ether 5a:79:e0:92:8b:84 brd ff:ff:ff:ff:ff:ff
    inet 192.168.1.2/24 scope global mac-1
       valid_lft forever preferred_lft forever
    inet6 fe80::5879:e0ff:fe92:8b84/64 scope link
       valid_lft forever preferred_lft forever
25: mac-2@wlp4s0: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc noqueue state UP group default qlen 1000
    link/ether aa:3f:0f:24:83:61 brd ff:ff:ff:ff:ff:ff
    inet 192.168.1.3/24 scope global mac-2
       valid_lft forever preferred_lft forever
    inet6 fe80::a83f:fff:fe24:8361/64 scope link
       valid_lft forever preferred_lft forever

```

Ping `ping -I 192.168.1.2 192.168.1.3`??

tcpdump result:
```
sudo tcpdump -vvvvvv -ni any -e icmp
tcpdump: listening on any, link-type LINUX_SLL (Linux cooked v1), capture size 262144 bytes
23:13:33.871661  In 00:00:00:00:00:00 ethertype IPv4 (0x0800), length 100: (tos 0x0, ttl 64, id 52231, offset 0, flags [DF], proto ICMP (1), length 84)
    192.168.1.2 > 192.168.1.3: ICMP echo request, id 9, seq 75, length 64
23:13:33.871674  In 00:00:00:00:00:00 ethertype IPv4 (0x0800), length 100: (tos 0x0, ttl 64, id 53273, offset 0, flags [none], proto ICMP (1), length 84)
    192.168.1.3 > 192.168.1.2: ICMP echo reply, id 9, seq 75, length 64
```