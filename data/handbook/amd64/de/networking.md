# Networking

Other languages:

- Deutsch
- [English](/wiki/Handbook:AMD64/Installation/Networking "Handbook:AMD64/Installation/Networking (100% translated)")
- [Türkçe](/wiki/Handbook:AMD64/Installation/Networking/tr "Handbook:AMD64/Installation/Networking/tr (0% translated)")
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

[AMD64 Handbuch](/wiki/Handbook:AMD64/de "Handbook:AMD64/de")[Installation](/wiki/Handbook:AMD64/Full/Installation/de "Handbook:AMD64/Full/Installation/de")[Über die Installation](/wiki/Handbook:AMD64/Installation/About/de "Handbook:AMD64/Installation/About/de")[Auswahl des Mediums](/wiki/Handbook:AMD64/Installation/Media/de "Handbook:AMD64/Installation/Media/de")Konfiguration des Netzwerks[Vorbereiten der Festplatte(n)](/wiki/Handbook:AMD64/Installation/Disks/de "Handbook:AMD64/Installation/Disks/de")[Installation des Stage Archivs](/wiki/Handbook:AMD64/Installation/Stage/de "Handbook:AMD64/Installation/Stage/de")[Installation des Basissystems](/wiki/Handbook:AMD64/Installation/Base/de "Handbook:AMD64/Installation/Base/de")[Konfiguration des Kernels](/wiki/Handbook:AMD64/Installation/Kernel/de "Handbook:AMD64/Installation/Kernel/de")[Konfiguration des Systems](/wiki/Handbook:AMD64/Installation/System/de "Handbook:AMD64/Installation/System/de")[Installation der Tools](/wiki/Handbook:AMD64/Installation/Tools/de "Handbook:AMD64/Installation/Tools/de")[Konfiguration des Bootloaders](/wiki/Handbook:AMD64/Installation/Bootloader/de "Handbook:AMD64/Installation/Bootloader/de")[Abschluss](/wiki/Handbook:AMD64/Installation/Finalizing/de "Handbook:AMD64/Installation/Finalizing/de")[Arbeiten mit Gentoo](/wiki/Handbook:AMD64/Full/Working/de "Handbook:AMD64/Full/Working/de")[Portage-Einführung](/wiki/Handbook:AMD64/Working/Portage/de "Handbook:AMD64/Working/Portage/de")[USE-Flags](/wiki/Handbook:AMD64/Working/USE/de "Handbook:AMD64/Working/USE/de")[Portage-Features](/wiki/Handbook:AMD64/Working/Features/de "Handbook:AMD64/Working/Features/de")[Initskript-System](/wiki/Handbook:AMD64/Working/Initscripts/de "Handbook:AMD64/Working/Initscripts/de")[Umgebungsvariablen](/wiki/Handbook:AMD64/Working/EnvVar/de "Handbook:AMD64/Working/EnvVar/de")[Arbeiten mit Portage](/wiki/Handbook:AMD64/Full/Portage/de "Handbook:AMD64/Full/Portage/de")[Dateien und Verzeichnisse](/wiki/Handbook:AMD64/Portage/Files/de "Handbook:AMD64/Portage/Files/de")[Variablen](/wiki/Handbook:AMD64/Portage/Variables/de "Handbook:AMD64/Portage/Variables/de")[Mischen von Softwarezweigen](/wiki/Handbook:AMD64/Portage/Branches/de "Handbook:AMD64/Portage/Branches/de")[Zusätzliche Tools](/wiki/Handbook:AMD64/Portage/Tools/de "Handbook:AMD64/Portage/Tools/de")[Eigener Portage-Tree](/wiki/Handbook:AMD64/Portage/CustomTree/de "Handbook:AMD64/Portage/CustomTree/de")[Erweiterte Portage-Features](/wiki/Handbook:AMD64/Portage/Advanced/de "Handbook:AMD64/Portage/Advanced/de")[Netzwerk-Konfiguration](/wiki/Handbook:AMD64/Full/Networking/de "Handbook:AMD64/Full/Networking/de")[Zu Beginn](/wiki/Handbook:AMD64/Networking/Introduction/de "Handbook:AMD64/Networking/Introduction/de")[Fortgeschrittene Konfiguration](/wiki/Handbook:AMD64/Networking/Advanced/de "Handbook:AMD64/Networking/Advanced/de")[Modulare Vernetzung](/wiki/Handbook:AMD64/Networking/Modular/de "Handbook:AMD64/Networking/Modular/de")[Drahtlose Netzwerke](/wiki/Handbook:AMD64/Networking/Wireless/de "Handbook:AMD64/Networking/Wireless/de")[Funktionalität hinzufügen](/wiki/Handbook:AMD64/Networking/Extending/de "Handbook:AMD64/Networking/Extending/de")[Dynamisches Management](/wiki/Handbook:AMD64/Networking/Dynamic/de "Handbook:AMD64/Networking/Dynamic/de")

## Contents

- [1Automatische Netzwerk-Erkennung](#Automatische_Netzwerk-Erkennung)
  - [1.1Verwendung von DHCP](#Verwendung_von_DHCP)
  - [1.2Das Netzwerk testen](#Das_Netzwerk_testen)
  - [1.3Ermitteln der Interface-Namen](#Ermitteln_der_Interface-Namen)
- [2Wireless Setup](#Wireless_Setup)
  - [2.1Optional: Checking wireless](#Optional:_Checking_wireless)
  - [2.2Using NetworkManager](#Using_NetworkManager)
- [3Automatische Netzwerk-Konfiguration](#Automatische_Netzwerk-Konfiguration)
  - [3.1Manual setup](#Manual_setup)
    - [3.1.1Manually connecting to an open network](#Manually_connecting_to_an_open_network)
- [4Application specific configuration](#Application_specific_configuration)
  - [4.1Alternativ: Verwendung von PPP](#Alternativ:_Verwendung_von_PPP)
  - [4.2Alternativ: Verwendung von PPTP](#Alternativ:_Verwendung_von_PPTP)
  - [4.3Verstehen der Netzwerk-Terminologie](#Verstehen_der_Netzwerk-Terminologie)
  - [4.4Interfaces and addresses](#Interfaces_and_addresses)
  - [4.5Networks and CIDR](#Networks_and_CIDR)
  - [4.6The Internet](#The_Internet)
  - [4.7The Domain Name System](#The_Domain_Name_System)
- [5Manuelle Netzwerk-Konfiguration](#Manuelle_Netzwerk-Konfiguration)
  - [5.1Interface address configuration](#Interface_address_configuration)
  - [5.2Default route configuration](#Default_route_configuration)
  - [5.3DNS configuration](#DNS_configuration)

## Automatische Netzwerk-Erkennung

Vielleicht funktioniert es einfach?

Wenn sich Ihr System in einem Ethernet-Netzwerk mit einem DHCP-Server befindet, ist es sehr wahrscheinlich, dass Ihr Netz bereits konfiguriert ist. Sie können nun die zahlreichen Netzwerktools auf dem Installationsmedium wie beispielsweise ssh, scp, ping, irssi, wget und links nutzen.

### Verwendung von DHCP

DHCP (Dynamic Host Configuration Protocol) ermöglicht es die gesamte Netzwerkkonfiguration (IP-Adresse, Netzwerkmaske, Broadcast-Adresse, Gateway, DNS-Server etc.) dynamisch von einem Server zu beziehen. Das funktioniert logischerweise nur, wenn Sie einen DHCP-Server in Ihrem LAN haben oder Ihr Provider einen solchen Dienst anbietet. Benutzen Sie dhcpcd:

DHCP requires that a server be running on the same _Layer 2_ ( _Ethernet_) segment as the client requesting a _lease_. DHCP is often used on RFC1918 ( _private_) networks, but is also used to acquire public IP information from ISPs.

By default the Gentoo live media uses [NetworkManager](/wiki/NetworkManager "NetworkManager") which will automatically connect to via DHCP, if this is not the case then check the Ethernet cable for issues.

### Das Netzwerk testen

A properly configured _default_ route is a critical component of Internet connectivity, route configuration can be checked with:

`root #` `ip route`

```
default via 192.168.0.1 dev enp1s0
```

If no _default_ route is defined, Internet connectivity is unavailable, and additional configuration is required.

Basic internet connectivity can be confirmed with a ping:

`root #` `ping -c 3 1.1.1.1`

**Tipp**

It's helpful to start by pinging a known IP address instead of a hostname. This can isolate DNS issues from basic Internet connectivity issues.

Outbound HTTPS access and DNS resolution can be confirmed with:

`root #` `curl --location gentoo.org --output /dev/null`

Wenn Sie nun in der Lage sind, Ihr Netzwerk zu verwenden, dann können Sie den Rest dieses Kapitels überspringen und mit dem [Vorbereiten der Festplatte(n)](/wiki/Handbook:AMD64/Installation/Disks/de "Handbook:AMD64/Installation/Disks/de") fortfahren.

If curl reports an error, but Internet-bound pings work, [DNS may need configuration](/wiki/Handbook:AMD64/Installation/Networking/de#DNS_configuration "Handbook:AMD64/Installation/Networking/de").

If Internet connectivity has not been established, first [interface information should be verified](/wiki/Handbook:AMD64/Installation/Networking/de#Obtaining_interface_info "Handbook:AMD64/Installation/Networking/de"), then:

- [nmtui can be used](/wiki/Handbook:AMD64/Installation/Networking/de#Using_NetworkManager "Handbook:AMD64/Installation/Networking/de") to assist in network configuration.
- [Application specific configuration](/wiki/Handbook:AMD64/Installation/Networking/de#Application_specific_configuration "Handbook:AMD64/Installation/Networking/de") may be required.
- [Manual network configuration](/wiki/Handbook:AMD64/Installation/Networking/de#Manual_network_configuration "Handbook:AMD64/Installation/Networking/de") can be attempted.

### Ermitteln der Interface-Namen

If networking doesn't work out of the box, additional steps must be taken to enable Internet connectivity. Generally, the first step is to enumerate host network interfaces.

Als Alternative zu ifconfig kann zur Anzeige von Interface-Namen das Kommando ip verwendet werden. Das folgende Beispiel zeigt die Ausgabe von ip addr. Die ausgegebenen Daten unterscheiden sich vom letzten Beispiel, weil das ip-Kommando auf einem anderen System eingegeben wurde:

The **link** argument can be used to display network interface links:

`root #` `ip link`

```
1: lo: <LOOPBACK,UP,LOWER_UP> mtu 65536 qdisc noqueue state UNKNOWN mode DEFAULT group default qlen 1000
    link/loopback 00:00:00:00:00:00 brd 00:00:00:00:00:00
4: enp1s0: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc pfifo_fast state UP mode DEFAULT group default qlen 1000
    link/ether e8:40:f2:ac:25:7a brd ff:ff:ff:ff:ff:ff
```

The **address** argument can be used to query device address information:

`root #` `ip addr`

```
2: eno1: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc pfifo_fast state UP group default qlen 1000
    link/ether e8:40:f2:ac:25:7a brd ff:ff:ff:ff:ff:ff
    inet 10.0.20.77/22 brd 10.0.23.255 scope global eno1
       valid_lft forever preferred_lft forever
    inet6 fe80::ea40:f2ff:feac:257a/64 scope link
       valid_lft forever preferred_lft forever

```

The output of this command contains information for each network interface on the system. Entries begin with the device index, followed by the device name: enp1s0.

**Tipp**

Wenn die Ausgabe von ifconfig keine Interfaces anzeigt, starten Sie das Kommando noch einmal mit der Option `-a`.
Mit dieser Option zeigt ifconfig alle vom System erkannten Interfaces, unabhängig davon, ob sie im Zustand "up" oder "down" sind. Wenn ifconfig -a keine Interfaces anzeigt, ist entweder die Hardware defekt oder der erforderliche Kernel-Treiber ist nicht geladen. Beide Fälle können nicht in diesem Handbuch besprochen werden. Bitte kontaktieren Sie [#gentoo](ircs://irc.libera.chat/#gentoo) ([webchat](https://web.libera.chat/#gentoo)) für Hilfe.

In dem Rest dieses Dokuments geht das Handbuch davon aus, dass das genutzte Netzwerk-Interface den Namen eth0 hat.

Als Folge des Wechsels zu [predictable network interface names](http://www.freedesktop.org/wiki/Software/systemd/PredictableNetworkInterfaceNames/), kann sich der Interface-Name deutlich von der alten
"eth0"-Namens-Konvention unterscheiden. Aktuelle Installations-Medien zeigen möglicherweise Namen an wie: eno0, ens1, oder enp5s0. Suchen Sie nach dem Interface in der Ausgabe von ifconfig, das eine IP-Adresse aus Ihrem lokalen Netwerk hat.

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

**Hinweis**

Some wireless cards may have a device name of wlan0 or ra0 instead of wlp9s0. Run ip link to determine the correct device name.

### Using NetworkManager

The fastest way to enable wireless connectivity in the Gentoo Live media is by using the nmtui command and following the on-screen instructions to setup the wireless network.

`root #` `nmtui`

LiveGUI users can click on the network icon in the bottom right of the taskbar to also configure WiFi.

## Automatische Netzwerk-Konfiguration

On some architectures, e.g. **HPPA**, a user may be required to use the tool net-setup to configure a wireless connection if NetworkManager isn't available yet.

`root #` `net-setup eth0`

net-setup wird Ihnen einige Fragen bezüglich Ihrer Netzwerkumgebung stellen. Haben Sie alle Fragen beantwortet, sollten Sie eine funktionsfähige Netzwerkverbindung haben. Testen Sie Ihr Netzwerk wieder, wie oben beschrieben. Sollten die Tests funktionieren, so haben Sie es geschafft; Sie können nun mit der Installation von Gentoo fortfahren. Überspringen Sie den Rest dieses Kapitels und fahren Sie mit der [Vorbereitung der Festplatte(n) fort](/wiki/Handbook:AMD64/Installation/Disks/de "Handbook:AMD64/Installation/Disks/de").

**Wichtig**

Network status should be [tested](/wiki/Handbook:AMD64/Installation/Networking/de#Testing_the_network "Handbook:AMD64/Installation/Networking/de") after any configuration steps are taken. In the event that the configuration scripts do not work, [manual network configuration](/wiki/Handbook:AMD64/Installation/Networking/de#Manual_network_configuration "Handbook:AMD64/Installation/Networking/de") is required.

### Manual setup

Wireless networks may alternatively be set up with a daemon such as wpa\_supplicant or iwd. For more information on configuring wireless networking in Gentoo Linux, please read the [Wireless networking chapter](/wiki/Handbook:AMD64/Networking/Wireless/de "Handbook:AMD64/Networking/Wireless/de") in the Gentoo Handbook.}}

#### Manually connecting to an open network

For most users, there are only two settings needed to connect, the ESSID (aka wireless network name) and, optionally, the WEP key.

- First, ensure the interface is active:

`root #` `ip link set dev wlp9s0 up`

- To connect to an open network with the name _GentooNode_:

`root #` `iw dev wlp9s0 connect -w GentooNode`

## Application specific configuration

**Tipp**

These steps are provided for users where using nmtui is not able to configure their network's needs.

The following methods are not generally required, but may be helpful in situations where additional configuration is required for Internet connectivity.

===  Optional: Konfiguration eines Web-Proxies ===

Wenn Sie auf das Internet nur über einen Proxy-Server zugreifen können, müssen Sie während der Installation das System für die Verwendung des Proxy-Servers vorbereiten. Das ist aber recht einfach. Sie müssen dazu lediglich eine Variable mit den Informationen über den Proxy-Server setzen.

Certain text-mode web browsers such as links can also make use of environment variables that define web proxy settings; in particular for the HTTPS access it also will require the https\_proxy environment variable to be defined. While Portage will be influenced without passing extra run time parameters during invocation, links will require proxy settings to be set.

In den meisten Fällen können Sie den Hostnamen des Proxy-Servers in die Variable schreiben. Nehmen wir an, der Server ist proxy.gentoo.org und der Port ist 8080.

**Hinweis**

The `#` symbol in the following commands is a comment. It has been added for clarity only and does _not_ need to be typed when entering the commands.

Zur Einrichtung eines HTTP-Proxies (für HTTP- und HTTPS-Traffic):

`root #` `export http_proxy="http://proxy.gentoo.org:8080"`

Wenn der Proxy-Server einen Benutzernamen und Passwort erfordert, sollten Sie die folgende Syntax in der Variable verwenden:

CODE **Adding username/password to the proxy variable**

```
http://username:password@proxy.gentoo.org:8080

```

Start links using the following parameters for proxy support:

`user $` `links -http-proxy ${http_proxy} -https-proxy ${https_proxy} `

Zur Einrichtung eines FTP-Proxies:

`root #` `export ftp_proxy="ftp://proxy.gentoo.org:8080"`

Start links using the following parameter for a FTP proxy:

`user $` `links -ftp-proxy ${ftp_proxy} `

Zur Einrichtung eines RSYNC-Proxies:

`root #` `export RSYNC_PROXY="proxy.gentoo.org:8080"`

### Alternativ: Verwendung von PPP

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

### Alternativ: Verwendung von PPTP

Wenn Sie PPTP-Unterstüzung benötigen, können Sie das Programm pptpclient, das Ihnen von der Installations-CD bereitgestellt wird, verwenden. Allerdings müssen Sie vorher sichergehen, dass Ihre Konfiguration korrekt ist. Dazu editieren Sie die Datei /etc/ppp/pap-secrets oder /etc/ppp/chap-secrets, so dass diese die korrekte Benutzername/Kennwort-Kombination beinhalten.

Edit /etc/ppp/pap-secrets or /etc/ppp/chap-secrets so it contains the correct username/password combination:

`root #` `nano -w /etc/ppp/chap-secrets`

Wenn nötig, sollten Sie nun noch /etc/ppp/options.pptp anpassen:

`root #` `nano -w /etc/ppp/options.pptp`

Nun geben Sie den Befehl pptp (mit den Optionen, die Sie in options.pptp setzen könnten) ein, um sich mit dem Server zu verbinden.

`root #` `pptp <server ipv4 address>`

### Verstehen der Netzwerk-Terminologie

If all of the above fails, the network must be configured manually. This is not particularly difficult, but should be done with consideration. This section serves to clarify terminology and introduce users to basic networking concepts pertaining to manually configuring an Internet connection.

**Tipp**

Some **CPE** ( **Carrier Provided Equipment**) combines the functions of a _router_, _access point_, _modem_, _DHCP server_, and _DNS server_ into one unit. It's important to differentiate the functions of a device from the physical appliance.

### Interfaces and addresses

Network _interfaces_ are logical representations of network devices. An _interface_ needs an _address_ to communicate with other devices on the _network_. While only a single _address_ is required, multiple addresses can be assigned to a single _interface_. This is especially useful for dual stack (IPv4 + IPv6) configurations.

For consistency, this primer will assume the interface enp1s0 will be using the address 192.168.0.2.

**Wichtig**

IP addresses can be set arbitrarily. As a result, it's possible for multiple devices to use the same IP address, resulting in an _address conflict_. Address conflicts should be avoided by using DHCP or SLAAC.

**Tipp**

IPv6 typically uses **S** tate **L** ess **A** ddress **A** uto **C** onfiguration (SLAAC) for address configuration. In most cases, manually setting IPv6 addresses is a bad practice. If a specific address suffix is preferred, [interface identification tokens](/wiki/IPv6_Static_Addresses_using_Tokens "IPv6 Static Addresses using Tokens") can be used.

### Networks and CIDR

Once an address is chosen, how does a device know how to talk to other devices?

IP _addresses_ are associated with _networks_. IP _networks_ are contiguous logical ranges of addresses.

_Classless Inter-Domain Routing_ or _CIDR_ notation is used to distinguish network sizes.

- The _CIDR_ value, often notated starting with a **/**, represents the size of the network.

  - The formula _2 ^ (32 - CIDR)_ can be used to calculate network size.
  - Once network size is calculated, usable node count must be reduced by 2.
    - The first IP in a network is the _Network address_, and the last is typically the _Broadcast address_. These addresses are special and cannot be used by normal hosts.

**Tipp**

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

**Tipp**

It's a standard practice to make the _gateway_ the first or last IP on a network.

If an Internet-connected router is available at 192.168.0.1, it can be used as the _default route_, granting Internet access.

To summarize:

- Interfaces must be configured with an _address_ and _network information_, such as the _CIDR_ value.
- Local network access is used to access a _router_ on the same network.
- The _default route_ is configured, so traffic destined for external networks is forwarded to the _gateway_, providing Internet access.

### The Domain Name System

Remembering IPs is hard. The _Domain Name System_ was created to allow mapping between _Domain Names_ and _IP addresses_.

Linux systems use /etc/resolv.conf to define _nameservers_ to be used for _DNS resolution_.

**Tipp**

Many _routers_ can also function as a DNS server, and using a local DNS server can augment privacy and speed up queries through caching.

Many ISPs run a DNS server that is generally advertised to the _gateway_ over DHCP. Using a local DNS server tends to improve query latency, but most public DNS servers will return the same results, so server usage is largely based on preference.

## Manuelle Netzwerk-Konfiguration

### Interface address configuration

**Wichtig**

When manually configuring IP addresses, the local network topology must be considered. IP addresses can be set arbitrarily; conflicts may cause network disruption.

To configure enp1s0 with the address 192.168.0.2 and CIDR /24:

`root #` `ip address add 192.168.0.2/24 dev enp1s0`

**Tipp**

The start of this command can be shortened to ip a.

### Default route configuration

Configuring address and network information for an interface will configure link routes, allowing communication with that network segment:

`root #` `ip route`

```
192.168.0.0/24 dev enp1s0 proto kernel scope link src 192.168.0.2
```

**Tipp**

This command can be shortened to ip r.

The default route can be set to 192.168.0.1 with:

`root #` `ip route add default via 192.168.0.1`

### DNS configuration

Nameserver info is typically acquired using DHCP, but can be set manually by adding `nameserver` entries to /etc/resolv.conf.

**Warnung**

If _dhcpcd_ is running, changes to /etc/resolv.conf will not persist. Status can be checked with `ps x | grep dhcpcd`.

nano is included in Gentoo _boot media_ and can be used to edit /etc/resolv.conf with:

`root #` `nano -w /etc/resolv.conf`

Lines containing the keyword `nameserver` followed by a DNS server IP address are queried in order of definition:

DATEI **`/etc/resolv.conf`** **Use Quad9 DNS.**

```
nameserver 9.9.9.9
nameserver 149.112.112.112

```

DATEI **`/etc/resolv.conf`** **Use Cloudflare DNS.**

```
nameserver 1.1.1.1
nameserver 1.0.0.1

```

DNS status can be checked by pinging a domain name:

`root #` `ping -c 3 gentoo.org`

Once connectivity [has been verified](/wiki/Handbook:AMD64/Installation/Networking/de#Testing_the_network "Handbook:AMD64/Installation/Networking/de"), continue with [Preparing the disks](/wiki/Handbook:AMD64/Installation/Disks/de "Handbook:AMD64/Installation/Disks/de").

[← Auswahl des Mediums](/wiki/Handbook:AMD64/Installation/Media/de "Previous part") [Anfang](/wiki/Handbook:AMD64/de "Handbook:AMD64/de") [Vorbereiten der Festplatte(n) →](/wiki/Handbook:AMD64/Installation/Disks/de "Next part")