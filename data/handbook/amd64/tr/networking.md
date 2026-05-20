# Networking

Other languages:

- [Deutsch](/wiki/Handbook:AMD64/Installation/Networking/de "Handbook:AMD64/Installation/Networking/de (100% translated)")
- [English](/wiki/Handbook:AMD64/Installation/Networking "Handbook:AMD64/Installation/Networking (100% translated)")
- Türkçe
- [español](/wiki/Handbook:AMD64/Installation/Networking/es "Handbook:AMD64/Installation/Networking/es (100% translated)")
- [français](/wiki/Handbook:AMD64/Installation/Networking/fr "Handbook:AMD64/Installation/Networking/fr (100% translated)")
- [italiano](/wiki/Handbook:AMD64/Installation/Networking/it "Handbook:AMD64/Installation/Networking/it (100% translated)")
- [magyar](/wiki/Handbook:AMD64/Installation/Networking/hu "Handbook:AMD64/Installation/Networking/hu (100% translated)")
- [polski](/wiki/Handbook:AMD64/Installation/Networking/pl "Handbook:AMD64/Installation/Networking/pl (100% translated)")
- [português do Brasil](/wiki/Handbook:AMD64/Installation/Networking/pt-br "Handbook:AMD64/Installation/Networking/pt-br (100% translated)")
- [čeština](/wiki/Handbook:AMD64/Installation/Networking/cs "Handbook:AMD64/Installation/Networking/cs (100% translated)")
- [русский](/wiki/Handbook:AMD64/Installation/Networking/ru "Handbook:AMD64/Installation/Networking/ru (100% translated)")
- [العربية](/wiki/Handbook:AMD64/Installation/Networking/ar "Handbook:AMD64/Installation/Networking/ar (100% translated)")
- [தமிழ்](/wiki/Handbook:AMD64/Installation/Networking/ta "Handbook:AMD64/Installation/Networking/ta (100% translated)")
- [中文（中国大陆）‎](/wiki/Handbook:AMD64/Installation/Networking/zh-cn "Handbook:AMD64/Installation/Networking/zh-cn (100% translated)")
- [日本語](/wiki/Handbook:AMD64/Installation/Networking/ja "Handbook:AMD64/Installation/Networking/ja (100% translated)")
- [한국어](/wiki/Handbook:AMD64/Installation/Networking/ko "Handbook:AMD64/Installation/Networking/ko (100% translated)")

[AMD64 Handbook](/wiki/Handbook:AMD64 "Handbook:AMD64")[Installation](/wiki/Handbook:AMD64/Full/Installation "Handbook:AMD64/Full/Installation")[About the installation](/wiki/Handbook:AMD64/Installation/About/tr "Handbook:AMD64/Installation/About/tr")[Choosing the media](/wiki/Handbook:AMD64/Installation/Media/tr "Handbook:AMD64/Installation/Media/tr")Configuring the network[Preparing the disks](/wiki/Handbook:AMD64/Installation/Disks/tr "Handbook:AMD64/Installation/Disks/tr")[The stage file](/wiki/Handbook:AMD64/Installation/Stage/tr "Handbook:AMD64/Installation/Stage/tr")[Installing base system](/wiki/Handbook:AMD64/Installation/Base "Handbook:AMD64/Installation/Base")[Configuring the kernel](/wiki/Handbook:AMD64/Installation/Kernel "Handbook:AMD64/Installation/Kernel")[Configuring the system](/wiki/Handbook:AMD64/Installation/System "Handbook:AMD64/Installation/System")[Installing tools](/wiki/Handbook:AMD64/Installation/Tools "Handbook:AMD64/Installation/Tools")[Configuring the bootloader](/wiki/Handbook:AMD64/Installation/Bootloader "Handbook:AMD64/Installation/Bootloader")[Finalizing](/wiki/Handbook:AMD64/Installation/Finalizing "Handbook:AMD64/Installation/Finalizing")[Working with Gentoo](/wiki/Handbook:AMD64/Full/Working "Handbook:AMD64/Full/Working")[Portage introduction](/wiki/Handbook:AMD64/Working/Portage "Handbook:AMD64/Working/Portage")[USE flags](/wiki/Handbook:AMD64/Working/USE "Handbook:AMD64/Working/USE")[Portage features](/wiki/Handbook:AMD64/Working/Features "Handbook:AMD64/Working/Features")[Initscript system](/wiki/Handbook:AMD64/Working/Initscripts "Handbook:AMD64/Working/Initscripts")[Environment variables](/wiki/Handbook:AMD64/Working/EnvVar "Handbook:AMD64/Working/EnvVar")[Working with Portage](/wiki/Handbook:AMD64/Full/Portage "Handbook:AMD64/Full/Portage")[Files and directories](/wiki/Handbook:AMD64/Portage/Files "Handbook:AMD64/Portage/Files")[Variables](/wiki/Handbook:AMD64/Portage/Variables "Handbook:AMD64/Portage/Variables")[Mixing software branches](/wiki/Handbook:AMD64/Portage/Branches "Handbook:AMD64/Portage/Branches")[Additional tools](/wiki/Handbook:AMD64/Portage/Tools "Handbook:AMD64/Portage/Tools")[Custom package repository](/wiki/Handbook:AMD64/Portage/CustomTree "Handbook:AMD64/Portage/CustomTree")[Advanced features](/wiki/Handbook:AMD64/Portage/Advanced "Handbook:AMD64/Portage/Advanced")[OpenRC network configuration](/wiki/Handbook:AMD64/Full/Networking "Handbook:AMD64/Full/Networking")[Getting started](/wiki/Handbook:AMD64/Networking/Introduction "Handbook:AMD64/Networking/Introduction")[Advanced configuration](/wiki/Handbook:AMD64/Networking/Advanced "Handbook:AMD64/Networking/Advanced")[Modular networking](/wiki/Handbook:AMD64/Networking/Modular "Handbook:AMD64/Networking/Modular")[Wireless](/wiki/Handbook:AMD64/Networking/Wireless "Handbook:AMD64/Networking/Wireless")[Adding functionality](/wiki/Handbook:AMD64/Networking/Extending "Handbook:AMD64/Networking/Extending")[Dynamic management](/wiki/Handbook:AMD64/Networking/Dynamic "Handbook:AMD64/Networking/Dynamic")

## Contents

- [1Automatic network configuration](#Automatic_network_configuration)
  - [1.1Using DHCP](#Using_DHCP)
  - [1.2Testing the network](#Testing_the_network)
- [2Obtaining interface info](#Obtaining_interface_info)
- [3Wireless Setup](#Wireless_Setup)
  - [3.1Optional: Checking wireless](#Optional:_Checking_wireless)
  - [3.2Using NetworkManager](#Using_NetworkManager)
  - [3.3Using net-setup](#Using_net-setup)
  - [3.4Manual setup](#Manual_setup)
    - [3.4.1Manually connecting to an open network](#Manually_connecting_to_an_open_network)
- [4Application specific configuration](#Application_specific_configuration)
  - [4.1Configure web proxies](#Configure_web_proxies)
  - [4.2Using pppoe-setup for ADSL](#Using_pppoe-setup_for_ADSL)
  - [4.3Using PPTP](#Using_PPTP)
- [5Internet and IP basics](#Internet_and_IP_basics)
  - [5.1Interfaces and addresses](#Interfaces_and_addresses)
  - [5.2Networks and CIDR](#Networks_and_CIDR)
  - [5.3The Internet](#The_Internet)
  - [5.4The Domain Name System](#The_Domain_Name_System)
- [6Manual and static IP network configuration](#Manual_and_static_IP_network_configuration)
  - [6.1Interface address configuration](#Interface_address_configuration)
  - [6.2Default route configuration](#Default_route_configuration)
  - [6.3DNS configuration](#DNS_configuration)

## Automatic network configuration

Maybe it just works?

If the system is connected to an Ethernet network with an IPv6 router or DHCP server, it's very likely that the system's network was configured automatically. If additional, advanced configuration is not required, [Internet connectivity can be tested](/wiki/Handbook:AMD64/Installation/Networking/tr#Testing_the_network "Handbook:AMD64/Installation/Networking/tr").

### Using DHCP

DHCP (Dynamic Host Configuration Protocol) assists in network configuration, and can automatically provide configuration for a variety of parameters including: IP address, network mask, routes, DNS servers, NTP servers, etc.

DHCP requires that a server be running on the same _Layer 2_ ( _Ethernet_) segment as the client requesting a _lease_. DHCP is often used on RFC1918 ( _private_) networks, but is also used to acquire public IP information from ISPs.

By default the Gentoo live media uses [NetworkManager](/wiki/NetworkManager "NetworkManager") which will automatically connect to via DHCP, if this is not the case then check the Ethernet cable for issues.

### Testing the network

A properly configured _default_ route is a critical component of Internet connectivity, route configuration can be checked with:

`root #` `ip route`

```
default via 192.168.0.1 dev enp1s0
```

If no _default_ route is defined, Internet connectivity is unavailable, and additional configuration is required.

Basic internet connectivity can be confirmed with a ping:

`root #` `ping -c 3 1.1.1.1`

**Tip**

It's helpful to start by pinging a known IP address instead of a hostname. This can isolate DNS issues from basic Internet connectivity issues.

Outbound HTTPS access and DNS resolution can be confirmed with:

`root #` `curl --location gentoo.org --output /dev/null`

Unless curl reports an error, or other tests fail, the installation process can be continued with [disk preparation](/wiki/Handbook:AMD64/Installation/Disks/tr "Handbook:AMD64/Installation/Disks/tr").

If curl reports an error, but Internet-bound pings work, [DNS may need configuration](/wiki/Handbook:AMD64/Installation/Networking/tr#DNS_configuration "Handbook:AMD64/Installation/Networking/tr").

If Internet connectivity has not been established, first [interface information should be verified](/wiki/Handbook:AMD64/Installation/Networking/tr#Obtaining_interface_info "Handbook:AMD64/Installation/Networking/tr"), then:

- [nmtui can be used](/wiki/Handbook:AMD64/Installation/Networking/tr#Using_NetworkManager "Handbook:AMD64/Installation/Networking/tr") to assist in network configuration.
- [Application specific configuration](/wiki/Handbook:AMD64/Installation/Networking/tr#Application_specific_configuration "Handbook:AMD64/Installation/Networking/tr") may be required.
- [Manual network configuration](/wiki/Handbook:AMD64/Installation/Networking/tr#Manual_network_configuration "Handbook:AMD64/Installation/Networking/tr") can be attempted.

## Obtaining interface info

If networking doesn't work out of the box, additional steps must be taken to enable Internet connectivity. Generally, the first step is to enumerate host network interfaces.

The ip command, part of the [sys-apps/iproute2](https://packages.gentoo.org/packages/sys-apps/iproute2) package, can be used to query and configure system networking.

The **link** argument can be used to display network interface links:

`root #` `ip link`

```
1: lo: <LOOPBACK,UP,LOWER_UP> mtu 65536 qdisc noqueue state UNKNOWN mode DEFAULT group default qlen 1000
    link/loopback 00:00:00:00:00:00 brd 00:00:00:00:00:00
4: enp1s0: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc pfifo_fast state UP mode DEFAULT group default qlen 1000
    link/ether e8:40:f2:ac:25:7a brd ff:ff:ff:ff:ff:ff
```

The **address** argument can be used to query device address information:

`root #` `ip address`

```
2: enp1s0: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc pfifo_fast state UP group default qlen 1000<pre>
    link/ether e8:40:f2:ac:25:7a brd ff:ff:ff:ff:ff:ff
    inet 10.0.20.77/22 brd 10.0.23.255 scope global enp1s0
       valid_lft forever preferred_lft forever
    inet6 fe80::ea40:f2ff:feac:257a/64 scope link
       valid_lft forever preferred_lft forever
```

The output of this command contains information for each network interface on the system. Entries begin with the device index, followed by the device name: enp1s0.

**Tip**

If no interfaces other than the **lo** ( _loopack_) are displayed, then the networking hardware is faulty, or the driver for the interface has not been loaded into the kernel. Both situations reach beyond the scope of this Handbook. Please ask for support in contact [#gentoo](ircs://irc.libera.chat/#gentoo) ([webchat](https://web.libera.chat/#gentoo)).

For consistency, the handbook will assume that the primary network interface is called enp1s0.

**Not**

As a result of the shift toward [predictable network interface names](https://www.freedesktop.org/wiki/Software/systemd/PredictableNetworkInterfaceNames/), the interface name on the system can be quite different than the old eth0 naming convention. Modern Gentoo boot media uses interface names with prefixes such as eno0, ens1, or enp5s0.

## Wireless Setup

### Optional: Checking wireless

iw may be used to display the current wireless settings on the card, this should also show that the kernel wireless stack is working (note that the iw command is only available on the following architectures: **amd64**, **x86**, **arm**, **arm64**, **ppc**, **ppc64**, and **riscv**):

`root #` `iw dev wlp9s0 info`

```
Interface wlp9s0
	ifindex 3
	wdev 0x1
	addr 00:00:00:00:00:00
	type managed
	wiphy 0
	channel 11 (2462 MHz), width: 20 MHz (no HT), center1: 2462 MHz
	txpower 30.00 dBm

```

**Not**

Some wireless cards may have a device name of wlan0 or ra0 instead of wlp9s0. Run ip link to determine the correct device name.

### Using NetworkManager

The fastest way to enable wireless connectivity in the Gentoo Live media is by using the nmtui command and following the on-screen instructions to setup the wireless network.

`root #` `nmtui`

LiveGUI users can click on the network icon in the bottom right of the taskbar to also configure WiFi.

### Using net-setup

On some architectures, e.g. **HPPA**, a user may be required to use the tool net-setup to configure a wireless connection if NetworkManager isn't available yet.

`root #` `killall dhcpcd`

`root #` `net-setup enp1s0`

net-setup will ask some questions about the network environment and will use that information to configure wpa\_supplicant or _static addressing_.

**Önemli**

Network status should be [tested](/wiki/Handbook:AMD64/Installation/Networking/tr#Testing_the_network "Handbook:AMD64/Installation/Networking/tr") after any configuration steps are taken. In the event that the configuration scripts do not work, [manual network configuration](/wiki/Handbook:AMD64/Installation/Networking/tr#Manual_network_configuration "Handbook:AMD64/Installation/Networking/tr") is required.

### Manual setup

Wireless networks may alternatively be set up with a daemon such as wpa\_supplicant or iwd. For more information on configuring wireless networking in Gentoo Linux, please read the [Wireless networking chapter](/wiki/Handbook:AMD64/Networking/Wireless "Handbook:AMD64/Networking/Wireless") in the Gentoo Handbook.}}

#### Manually connecting to an open network

For most users, there are only two settings needed to connect, the ESSID (aka wireless network name) and, optionally, the WEP key.

- First, ensure the interface is active:

`root #` `ip link set dev wlp9s0 up`

- To connect to an open network with the name _GentooNode_:

`root #` `iw dev wlp9s0 connect -w GentooNode`

## Application specific configuration

**Tip**

These steps are provided for users where using nmtui is not able to configure their network's needs.

The following methods are not generally required, but may be helpful in situations where additional configuration is required for Internet connectivity.

### Configure web proxies

If the internet is accessed through a web proxy, then it will be necessary to define proxy information to for Portage to properly access the proxy for each supported protocol. Portage observes the http\_proxy, ftp\_proxy, and RSYNC\_PROXY environment variables in order to download packages via its wget and rsync retrieval mechanisms.

Certain text-mode web browsers such as links can also make use of environment variables that define web proxy settings; in particular for the HTTPS access it also will require the https\_proxy environment variable to be defined. While Portage will be influenced without passing extra run time parameters during invocation, links will require proxy settings to be set.

In most cases, it is sufficient to define environment variables using the server hostname. In the following example, it is assumed the proxy server host is called proxy.gentoo.org and the port is 8080.

**Not**

The `#` symbol in the following commands is a comment. It has been added for clarity only and does _not_ need to be typed when entering the commands.

To define an HTTP proxy (for HTTP and HTTPS traffic):

`root #` `export http_proxy="http://proxy.gentoo.org:8080" # Applies to Portage and Links
`

`root #` `export https_proxy="http://proxy.gentoo.org:8080" # Only applies for Links
`

If the HTTP proxy requires authentication, set a username and password with the following syntax:

`root #` `export http_proxy="http://username:password@proxy.gentoo.org:8080" # Applies to Portage and Links
`

`root #` `export https_proxy="http://username:password@proxy.gentoo.org:8080" # Only applies for Links
`

Start links using the following parameters for proxy support:

`user $` `links -http-proxy ${http_proxy} -https-proxy ${https_proxy} `

To define an FTP proxy for Portage and/or links:

`root #` `export ftp_proxy="ftp://proxy.gentoo.org:8080" # Applies to Portage and Links`

Start links using the following parameter for a FTP proxy:

`user $` `links -ftp-proxy ${ftp_proxy} `

To define an RSYNC proxy for Portage:

`root #` `export RSYNC_PROXY="proxy.gentoo.org:8080" # Applies to Portage; Links does not support a rsync proxy`

### Using pppoe-setup for ADSL

If PPPoE is required for Internet access, the Gentoo _boot media_ includes the pppoe-setup script to simplify ppp configuration.

During setup, pppoe-setup will ask for:

- The name of the Ethernet _interface_ connected to the ADSL modem.
- The PPPoE username and password.
- DNS server IPs.
- Whether or not a firewall is needed.

`root #` `pppoe-setup
`

`root #` `pppoe-start`

In the event of failure, credentials in /etc/ppp/pap-secrets or /etc/ppp/chap-secrets should be verified. If credentials are correct, PPPoE Ethernet interface selection should be checked.

### Using PPTP

If PPTP support is needed, pptpclient can be used, but requires configuration prior to usage.

Edit /etc/ppp/pap-secrets or /etc/ppp/chap-secrets so it contains the correct username/password combination:

`root #` `nano /etc/ppp/chap-secrets`

Then adjust /etc/ppp/options.pptp if necessary:

`root #` `nano /etc/ppp/options.pptp`

Once configuration is complete, run pptp (along with the options that couldn't be set in options.pptp) to connect the server:

`root #` `pptp <server ipv4 address>`

## Internet and IP basics

If all of the above fails, the network must be configured manually. This is not particularly difficult, but should be done with consideration. This section serves to clarify terminology and introduce users to basic networking concepts pertaining to manually configuring an Internet connection.

**Tip**

Some **CPE** ( **Carrier Provided Equipment**) combines the functions of a _router_, _access point_, _modem_, _DHCP server_, and _DNS server_ into one unit. It's important to differentiate the functions of a device from the physical appliance.

### Interfaces and addresses

Network _interfaces_ are logical representations of network devices. An _interface_ needs an _address_ to communicate with other devices on the _network_. While only a single _address_ is required, multiple addresses can be assigned to a single _interface_. This is especially useful for dual stack (IPv4 + IPv6) configurations.

For consistency, this primer will assume the interface enp1s0 will be using the address 192.168.0.2.

**Önemli**

IP addresses can be set arbitrarily. As a result, it's possible for multiple devices to use the same IP address, resulting in an _address conflict_. Address conflicts should be avoided by using DHCP or SLAAC.

**Tip**

IPv6 typically uses **S** tate **L** ess **A** ddress **A** uto **C** onfiguration (SLAAC) for address configuration. In most cases, manually setting IPv6 addresses is a bad practice. If a specific address suffix is preferred, [interface identification tokens](/wiki/IPv6_Static_Addresses_using_Tokens "IPv6 Static Addresses using Tokens") can be used.

### Networks and CIDR

Once an address is chosen, how does a device know how to talk to other devices?

IP _addresses_ are associated with _networks_. IP _networks_ are contiguous logical ranges of addresses.

_Classless Inter-Domain Routing_ or _CIDR_ notation is used to distinguish network sizes.

- The _CIDR_ value, often notated starting with a **/**, represents the size of the network.

  - The formula _2 ^ (32 - CIDR)_ can be used to calculate network size.
  - Once network size is calculated, usable node count must be reduced by 2.
    - The first IP in a network is the _Network address_, and the last is typically the _Broadcast address_. These addresses are special and cannot be used by normal hosts.

**Tip**

The most common _CIDR_ values are **/24**, and **/32**, representing **254** nodes and a single node respectively.

A _CIDR_ of **/24** is the de-facto default network size. This corresponds to a subnet mask of _255.255.255.0_, where the last 8 bits are reserved for IP addresses for nodes on a network.

The notation: 192.168.0.2/24 can be interpreted as:

- The _address_ 192.168.0.2
- On the _network_ 192.168.0.0
- With a size of **254** ( _2 ^ (32 - 24) \- 2_)

  - Usable IPs are in the range 192.168.0.1 - 192.168.0.254
- With a _broadcast address_ of 192.168.0.255

  - In most cases, the last address on a network is used as the _broadcast address_, but this can be changed.

Using this configuration, a device should be able to communicate with any host on the same _network_ (192.168.0.0).

### The Internet

Once a device is on a network, how does it know how to talk to devices on the Internet?

To communicate with devices outside of local _networks_, _routing_ must be used. A _router_ is simply a network device that forwards traffic for other devices. The term _default route_ or _gateway_ typically refers to whatever device on the current network is used for external network access.

**Tip**

It's a standard practice to make the _gateway_ the first or last IP on a network.

If an Internet-connected router is available at 192.168.0.1, it can be used as the _default route_, granting Internet access.

To summarize:

- Interfaces must be configured with an _address_ and _network information_, such as the _CIDR_ value.
- Local network access is used to access a _router_ on the same network.
- The _default route_ is configured, so traffic destined for external networks is forwarded to the _gateway_, providing Internet access.

### The Domain Name System

Remembering IPs is hard. The _Domain Name System_ was created to allow mapping between _Domain Names_ and _IP addresses_.

Linux systems use /etc/resolv.conf to define _nameservers_ to be used for _DNS resolution_.

**Tip**

Many _routers_ can also function as a DNS server, and using a local DNS server can augment privacy and speed up queries through caching.

Many ISPs run a DNS server that is generally advertised to the _gateway_ over DHCP. Using a local DNS server tends to improve query latency, but most public DNS servers will return the same results, so server usage is largely based on preference.

## Manual and static IP network configuration

### Interface address configuration

**Önemli**

When manually configuring IP addresses, the local network topology must be considered. IP addresses can be set arbitrarily; conflicts may cause network disruption.

To configure enp1s0 with the address 192.168.0.2 and CIDR /24:

`root #` `ip address add 192.168.0.2/24 dev enp1s0`

**Tip**

The start of this command can be shortened to ip a.

### Default route configuration

Configuring address and network information for an interface will configure link routes, allowing communication with that network segment:

`root #` `ip route`

```
192.168.0.0/24 dev enp1s0 proto kernel scope link src 192.168.0.2
```

**Tip**

This command can be shortened to ip r.

The default route can be set to 192.168.0.1 with:

`root #` `ip route add default via 192.168.0.1`

### DNS configuration

Nameserver info is typically acquired using DHCP, but can be set manually by adding `nameserver` entries to /etc/resolv.conf.

**Uyarı**

If _dhcpcd_ is running, changes to /etc/resolv.conf will not persist. Status can be checked with `ps x | grep dhcpcd`.

nano is included in Gentoo _boot media_ and can be used to edit /etc/resolv.conf with:

`root #` `nano /etc/resolv.conf`

Lines containing the keyword `nameserver` followed by a DNS server IP address are queried in order of definition:

Dosya **`/etc/resolv.conf`** **Use Quad9 DNS.**

```
nameserver 9.9.9.9
nameserver 149.112.112.112

```

Dosya **`/etc/resolv.conf`** **Use Cloudflare DNS.**

```
nameserver 1.1.1.1
nameserver 1.0.0.1

```

DNS status can be checked by pinging a domain name:

`root #` `ping -c 3 gentoo.org`

Once connectivity [has been verified](/wiki/Handbook:AMD64/Installation/Networking/tr#Testing_the_network "Handbook:AMD64/Installation/Networking/tr"), continue with [Preparing the disks](/wiki/Handbook:AMD64/Installation/Disks/tr "Handbook:AMD64/Installation/Disks/tr").

[← Choosing the media](/wiki/Handbook:AMD64/Installation/Media "Previous part") [Home](/wiki/Handbook:AMD64 "Handbook:AMD64") [Preparing the disks →](/wiki/Handbook:AMD64/Installation/Disks "Next part")