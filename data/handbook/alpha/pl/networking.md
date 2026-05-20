# Networking

Other languages:

- [Deutsch](/wiki/Handbook:Alpha/Installation/Networking/de "Handbuch:Alpha/Installation/Netzwerk (100% translated)")
- [English](/wiki/Handbook:Alpha/Installation/Networking "Handbook:Alpha/Installation/Networking (100% translated)")
- [Türkçe](/wiki/Handbook:Alpha/Installation/Networking/tr "Handbook:Alpha/Installation/Networking/tr (0% translated)")
- [español](/wiki/Handbook:Alpha/Installation/Networking/es "Manual de Gentoo: Alpha/Instalación/Redes (100% translated)")
- [français](/wiki/Handbook:Alpha/Installation/Networking/fr "Handbook:Alpha/Installation/Networking/fr (100% translated)")
- [italiano](/wiki/Handbook:Alpha/Installation/Networking/it "Handbook:Alpha/Installation/Networking/it (100% translated)")
- [magyar](/wiki/Handbook:Alpha/Installation/Networking/hu "Handbook:Alpha/Installation/Networking/hu (100% translated)")
- polski
- [português do Brasil](/wiki/Handbook:Alpha/Installation/Networking/pt-br "Manual:Alpha/Instalação/Rede (100% translated)")
- [čeština](/wiki/Handbook:Alpha/Installation/Networking/cs "Handbook:Alpha/Installation/Networking/cs (50% translated)")
- [русский](/wiki/Handbook:Alpha/Installation/Networking/ru "Handbook:Alpha/Installation/Networking (100% translated)")
- [தமிழ்](/wiki/Handbook:Alpha/Installation/Networking/ta "கையேடு:Alpha/நிறுவல்/வலையமைத்தல் (100% translated)")
- [中文（中国大陆）‎](/wiki/Handbook:Alpha/Installation/Networking/zh-cn "手册：Alpha/安装/配置网络 (100% translated)")
- [日本語](/wiki/Handbook:Alpha/Installation/Networking/ja "ハンドブック:Alpha/インストール/ネットワーク (100% translated)")
- [한국어](/wiki/Handbook:Alpha/Installation/Networking/ko "Handbook:Alpha/Installation/Networking/ko (100% translated)")

[Alpha Handbook](/wiki/Handbook:Alpha/pl "Handbook:Alpha/pl")[Instalacja](/wiki/Handbook:Alpha/Full/Installation/pl "Handbook:Alpha/Full/Installation/pl")[O instalacji](/wiki/Handbook:Alpha/Installation/About/pl "Handbook:Alpha/Installation/About/pl")[Wybór medium instalacyjnego](/wiki/Handbook:Alpha/Installation/Media/pl "Handbook:Alpha/Installation/Media/pl")Konfiguracja sieci[Przygotowanie dysków](/wiki/Handbook:Alpha/Installation/Disks/pl "Handbook:Alpha/Installation/Disks/pl")[Instalacja etapu 3](/wiki/Handbook:Alpha/Installation/Stage/pl "Handbook:Alpha/Installation/Stage/pl")[Instalacja systemu podstawowego](/wiki/Handbook:Alpha/Installation/Base/pl "Handbook:Alpha/Installation/Base/pl")[Konfiguracja jądra](/wiki/Handbook:Alpha/Installation/Kernel/pl "Handbook:Alpha/Installation/Kernel/pl")[Konfiguracja systemu](/wiki/Handbook:Alpha/Installation/System/pl "Handbook:Alpha/Installation/System/pl")[Instalacja narzędzi](/wiki/Handbook:Alpha/Installation/Tools/pl "Handbook:Alpha/Installation/Tools/pl")[Instalacja systemu rozruchowego](/wiki/Handbook:Alpha/Installation/Bootloader/pl "Handbook:Alpha/Installation/Bootloader/pl")[Finalizacja](/wiki/Handbook:Alpha/Installation/Finalizing/pl "Handbook:Alpha/Installation/Finalizing/pl")[Praca z Gentoo](/wiki/Handbook:Alpha/Full/Working/pl "Handbook:Alpha/Full/Working/pl")[Wstęp do Portage](/wiki/Handbook:Alpha/Working/Portage/pl "Handbook:Alpha/Working/Portage/pl")[Flagi USE](/wiki/Handbook:Alpha/Working/USE/pl "Handbook:Alpha/Working/USE/pl")[Funkcje portage](/wiki/Handbook:Alpha/Working/Features/pl "Handbook:Alpha/Working/Features/pl")[System initscript](/wiki/Handbook:Alpha/Working/Initscripts/pl "Handbook:Alpha/Working/Initscripts/pl")[Zmienne środowiskowe](/wiki/Handbook:Alpha/Working/EnvVar/pl "Handbook:Alpha/Working/EnvVar/pl")[Praca z Portage](/wiki/Handbook:Alpha/Full/Portage/pl "Handbook:Alpha/Full/Portage/pl")[Pliki i katalogi](/wiki/Handbook:Alpha/Portage/Files/pl "Handbook:Alpha/Portage/Files/pl")[Zmienne](/wiki/Handbook:Alpha/Portage/Variables/pl "Handbook:Alpha/Portage/Variables/pl")[Mieszanie działów oprogramowania](/wiki/Handbook:Alpha/Portage/Branches/pl "Handbook:Alpha/Portage/Branches/pl")[Dodatkowe narzędzia](/wiki/Handbook:Alpha/Portage/Tools/pl "Handbook:Alpha/Portage/Tools/pl")[Custom package repository](/wiki/Handbook:Alpha/Portage/CustomTree/pl "Handbook:Alpha/Portage/CustomTree/pl")[Funkcje zaawansowane](/wiki/Handbook:Alpha/Portage/Advanced/pl "Handbook:Alpha/Portage/Advanced/pl")[Konfiguracja sieci](/wiki/Handbook:Alpha/Full/Networking/pl "Handbook:Alpha/Full/Networking/pl")[Zaczynamy](/wiki/Handbook:Alpha/Networking/Introduction/pl "Handbook:Alpha/Networking/Introduction/pl")[Zaawansowana konfiguracja](/wiki/Handbook:Alpha/Networking/Advanced/pl "Handbook:Alpha/Networking/Advanced/pl")[Sieć modularna](/wiki/Handbook:Alpha/Networking/Modular/pl "Handbook:Alpha/Networking/Modular/pl")[Sieć bezprzewodowa](/wiki/Handbook:Alpha/Networking/Wireless/pl "Handbook:Alpha/Networking/Wireless/pl")[Dodawanie funkcjonalności](/wiki/Handbook:Alpha/Networking/Extending/pl "Handbook:Alpha/Networking/Extending/pl")[Dynamiczne zarządzanie](/wiki/Handbook:Alpha/Networking/Dynamic/pl "Handbook:Alpha/Networking/Dynamic/pl")

## Contents

- [1Automatyczne wykrywanie sieci](#Automatyczne_wykrywanie_sieci)
  - [1.1Używając DHCP](#U.C5.BCywaj.C4.85c_DHCP)
  - [1.2Testowanie sieci](#Testowanie_sieci)
  - [1.3Ustal nazwy interfejsów](#Ustal_nazwy_interfejs.C3.B3w)
- [2Wireless Setup](#Wireless_Setup)
  - [2.1Optional: Checking wireless](#Optional:_Checking_wireless)
  - [2.2Using NetworkManager](#Using_NetworkManager)
- [3Automatyczna konfiguracja sieci](#Automatyczna_konfiguracja_sieci)
  - [3.1Manual setup](#Manual_setup)
    - [3.1.1Manually connecting to an open network](#Manually_connecting_to_an_open_network)
- [4Application specific configuration](#Application_specific_configuration)
  - [4.1Opcjonalnie: Konfiguracja serwera proxy](#Opcjonalnie:_Konfiguracja_serwera_proxy)
  - [4.2Alternatywa: Używając PPP](#Alternatywa:_U.C5.BCywaj.C4.85c_PPP)
  - [4.3Alternatywa: Używając PPTP](#Alternatywa:_U.C5.BCywaj.C4.85c_PPTP)
  - [4.4Zrozumienie terminologii sieciowej](#Zrozumienie_terminologii_sieciowej)
  - [4.5Interfaces and addresses](#Interfaces_and_addresses)
  - [4.6Networks and CIDR](#Networks_and_CIDR)
  - [4.7The Internet](#The_Internet)
  - [4.8The Domain Name System](#The_Domain_Name_System)
- [5Ręczna konfiguracja sieci](#R.C4.99czna_konfiguracja_sieci)
  - [5.1Interface address configuration](#Interface_address_configuration)
  - [5.2Default route configuration](#Default_route_configuration)
  - [5.3DNS configuration](#DNS_configuration)

## Automatyczne wykrywanie sieci

Może to po prostu działa?

Jeśli system jest podłączony do sieci Ethernet z serwerem DHCP, jest bardzo prawdopodobne, że sieć została już skonfigurowana automatycznie. Jeśli tak, to wiele poleceń obsługujących sieć na płycie instalacyjnej, takich jak ssh, scp, ping, irssi, wget i links, będzie działać natychmiast.

### Używając DHCP

DHCP (Dynamic Host Configuration Protocol) umożliwia automatyczne otrzymywanie informacji sieciowych (adres IP, maska sieci, adres rozgłoszeniowy, brama, serwery nazw itp.). Działa to tylko wtedy, gdy w sieci znajduje się serwer DHCP (lub jeśli dostawca usług internetowych zapewnia usługę DHCP). Aby interfejs sieciowy odbierał te informacje automatycznie, użyj dhcpcd:

DHCP requires that a server be running on the same _Layer 2_ ( _Ethernet_) segment as the client requesting a _lease_. DHCP is often used on RFC1918 ( _private_) networks, but is also used to acquire public IP information from ISPs.

By default the Gentoo live media uses [NetworkManager](/wiki/NetworkManager "NetworkManager") which will automatically connect to via DHCP, if this is not the case then check the Ethernet cable for issues.

### Testowanie sieci

A properly configured _default_ route is a critical component of Internet connectivity, route configuration can be checked with:

`root #` `ip route`

```
default via 192.168.0.1 dev enp1s0
```

If no _default_ route is defined, Internet connectivity is unavailable, and additional configuration is required.

Basic internet connectivity can be confirmed with a ping:

`root #` `ping -c 3 1.1.1.1`

**Wskazówka**

It's helpful to start by pinging a known IP address instead of a hostname. This can isolate DNS issues from basic Internet connectivity issues.

Outbound HTTPS access and DNS resolution can be confirmed with:

`root #` `curl --location gentoo.org --output /dev/null`

Jeśli wszystko zadziała, to pozostałą część tego rozdziału można pominąć i przejść od razu do następnego kroku instrukcji instalacji ( [Przygotowanie dysków](/wiki/Handbook:Alpha/Installation/Disks/pl "Handbook:Alpha/Installation/Disks/pl")).

If curl reports an error, but Internet-bound pings work, [DNS may need configuration](/wiki/Handbook:Alpha/Installation/Networking/pl#DNS_configuration "Handbook:Alpha/Installation/Networking/pl").

If Internet connectivity has not been established, first [interface information should be verified](/wiki/Handbook:Alpha/Installation/Networking/pl#Obtaining_interface_info "Handbook:Alpha/Installation/Networking/pl"), then:

- [nmtui can be used](/wiki/Handbook:Alpha/Installation/Networking/pl#Using_NetworkManager "Handbook:Alpha/Installation/Networking/pl") to assist in network configuration.
- [Application specific configuration](/wiki/Handbook:Alpha/Installation/Networking/pl#Application_specific_configuration "Handbook:Alpha/Installation/Networking/pl") may be required.
- [Manual network configuration](/wiki/Handbook:Alpha/Installation/Networking/pl#Manual_network_configuration "Handbook:Alpha/Installation/Networking/pl") can be attempted.

### Ustal nazwy interfejsów

If networking doesn't work out of the box, additional steps must be taken to enable Internet connectivity. Generally, the first step is to enumerate host network interfaces.

Jako alternatywę dla ifconfig można użyć polecenia ip do określenia nazw interfejsów. Poniższy przykład przedstawia dane wyjściowe funkcji ip addr (z innego systemu, więc wyświetlane informacje różnią się od poprzedniego przykładu):

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

**Wskazówka**

Jeśli nie zostały wyświetlone żadne interfejsy, gdy użyto standardowego polecenia ifconfig, spróbuj użyć tego samego polecenia z opcją `-a`. Ta opcja wymusza na narzędziu wyświetlanie wszystkich interfejsów sieciowych wykrytych przez system, niezależnie od tego, czy są włączone, czy wyłączone. Jeśli ifconfig -a nie daje żadnych wyników, oznacza to, że sprzęt jest uszkodzony lub sterownik interfejsu nie został załadowany do jądra. Obie sytuacje wykraczają poza zakres niniejszego podręcznika. Aby uzyskać pomoc, skontaktuj się z [#gentoo](ircs://irc.libera.chat/#gentoo) ([webchat](https://web.libera.chat/#gentoo)).

W pozostałej części tego podręcznika zakładamy, że działający interfejs sieciowy nosi nazwę eth0.

W wyniku przejścia w kierunku [przewidywalnych nazw interfejsów sieciowych](http://www.freedesktop.org/wiki/Software/systemd/PredictableNetworkInterfaceNames/), nazwa interfejsu eth0 w systemie może się znacznie różnić od starej konwencji nazewnictwa. Najnowsze nośniki instalacyjne mogą pokazywać zwykłe nazwy interfejsów sieciowych, takie jak eno0, ens1 lub enp5s0. Poszukaj interfejsu w danych wyjściowych ifconfig, który ma adres IP powiązany z siecią lokalną.

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

**Informacja**

Some wireless cards may have a device name of wlan0 or ra0 instead of wlp9s0. Run ip link to determine the correct device name.

### Using NetworkManager

The fastest way to enable wireless connectivity in the Gentoo Live media is by using the nmtui command and following the on-screen instructions to setup the wireless network.

`root #` `nmtui`

LiveGUI users can click on the network icon in the bottom right of the taskbar to also configure WiFi.

## Automatyczna konfiguracja sieci

On some architectures, e.g. **HPPA**, a user may be required to use the tool net-setup to configure a wireless connection if NetworkManager isn't available yet.

`root #` `net-setup eth0`

net-setup zada kilka pytań dotyczących środowiska sieciowego. Kiedy wszystko zostanie zrobione, połączenie sieciowe powinno działać. Przetestuj połączenie sieciowe zgodnie z wcześniejszym opisem. Jeśli testy wypadną pozytywnie, gratulujemy! Pomiń resztę tej sekcji i przejdź do [Przygotowanie dysków](/wiki/Handbook:Alpha/Installation/Disks/pl "Handbook:Alpha/Installation/Disks/pl").

**Ważne**

Network status should be [tested](/wiki/Handbook:Alpha/Installation/Networking/pl#Testing_the_network "Handbook:Alpha/Installation/Networking/pl") after any configuration steps are taken. In the event that the configuration scripts do not work, [manual network configuration](/wiki/Handbook:Alpha/Installation/Networking/pl#Manual_network_configuration "Handbook:Alpha/Installation/Networking/pl") is required.

### Manual setup

Wireless networks may alternatively be set up with a daemon such as wpa\_supplicant or iwd. For more information on configuring wireless networking in Gentoo Linux, please read the [Wireless networking chapter](/wiki/Handbook:Alpha/Networking/Wireless/pl "Handbook:Alpha/Networking/Wireless/pl") in the Gentoo Handbook.}}

#### Manually connecting to an open network

For most users, there are only two settings needed to connect, the ESSID (aka wireless network name) and, optionally, the WEP key.

- First, ensure the interface is active:

`root #` `ip link set dev wlp9s0 up`

- To connect to an open network with the name _GentooNode_:

`root #` `iw dev wlp9s0 connect -w GentooNode`

## Application specific configuration

**Wskazówka**

These steps are provided for users where using nmtui is not able to configure their network's needs.

The following methods are not generally required, but may be helpful in situations where additional configuration is required for Internet connectivity.

### Opcjonalnie: Konfiguracja serwera proxy

Jeśli dostęp do Internetu odbywa się za pośrednictwem serwera proxy, podczas instalacji konieczne jest skonfigurowanie informacji o serwerze proxy. Definiowanie proxy jest bardzo łatwe: wystarczy zdefiniować zmienną, która zawiera informacje o serwerze proxy.

Certain text-mode web browsers such as links can also make use of environment variables that define web proxy settings; in particular for the HTTPS access it also will require the https\_proxy environment variable to be defined. While Portage will be influenced without passing extra run time parameters during invocation, links will require proxy settings to be set.

W większości przypadków wystarczy zdefiniować zmienne za pomocą nazwy hosta serwera. Na przykład zakładamy, że serwer proxy nazywa się proxy.gentoo.org, a port to 8080.

**Informacja**

The `#` symbol in the following commands is a comment. It has been added for clarity only and does _not_ need to be typed when entering the commands.

Aby skonfigurować HTTP proxy (dla ruchu HTTP i HTTPS):

`root #` `export http_proxy="http://proxy.gentoo.org:8080"`

Jeśli serwer proxy wymaga nazwy użytkownika i hasła, użyj następującej składni zmiennej:

CODE **Dodawanie nazwy użytkownika i hasła do zmiennej proxy**

```
http://nazwa_użytkownika:hasło@proxy.gentoo.org:8080

```

Start links using the following parameters for proxy support:

`user $` `links -http-proxy ${http_proxy} -https-proxy ${https_proxy} `

Aby skonfigurować FTP proxy:

`root #` `export ftp_proxy="ftp://proxy.gentoo.org:8080"`

Start links using the following parameter for a FTP proxy:

`user $` `links -ftp-proxy ${ftp_proxy} `

Aby skonfigurować RSYNC proxy:

`root #` `export RSYNC_PROXY="proxy.gentoo.org:8080"`

### Alternatywa: Używając PPP

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

### Alternatywa: Używając PPTP

Jeśli potrzebna jest obsługa PPTP, użyj pptpclient, który jest dostarczany na instalacyjnych dyskach CD. Jednak najpierw upewnij się, że konfiguracja jest prawidłowa. Edytuj /etc/ppp/pap-secrets lub /etc/ppp/chap-secrets, aby zawierały poprawną kombinację nazwy użytkownika i hasła:

Edit /etc/ppp/pap-secrets or /etc/ppp/chap-secrets so it contains the correct username/password combination:

`root #` `nano -w /etc/ppp/chap-secrets`

Następnie dostosuj /etc/ppp/options.pptp, jeśli jest to konieczne:

`root #` `nano -w /etc/ppp/options.pptp`

Po wykonaniu wszystkich czynności uruchom pptp (wraz z opcjami, których nie można ustawić w options.pptp), aby połączyć się z serwerem:

`root #` `pptp <adres ip serwera>`

### Zrozumienie terminologii sieciowej

If all of the above fails, the network must be configured manually. This is not particularly difficult, but should be done with consideration. This section serves to clarify terminology and introduce users to basic networking concepts pertaining to manually configuring an Internet connection.

**Wskazówka**

Some **CPE** ( **Carrier Provided Equipment**) combines the functions of a _router_, _access point_, _modem_, _DHCP server_, and _DNS server_ into one unit. It's important to differentiate the functions of a device from the physical appliance.

### Interfaces and addresses

Network _interfaces_ are logical representations of network devices. An _interface_ needs an _address_ to communicate with other devices on the _network_. While only a single _address_ is required, multiple addresses can be assigned to a single _interface_. This is especially useful for dual stack (IPv4 + IPv6) configurations.

For consistency, this primer will assume the interface enp1s0 will be using the address 192.168.0.2.

**Ważne**

IP addresses can be set arbitrarily. As a result, it's possible for multiple devices to use the same IP address, resulting in an _address conflict_. Address conflicts should be avoided by using DHCP or SLAAC.

**Wskazówka**

IPv6 typically uses **S** tate **L** ess **A** ddress **A** uto **C** onfiguration (SLAAC) for address configuration. In most cases, manually setting IPv6 addresses is a bad practice. If a specific address suffix is preferred, [interface identification tokens](/wiki/IPv6_Static_Addresses_using_Tokens "IPv6 Static Addresses using Tokens") can be used.

### Networks and CIDR

Once an address is chosen, how does a device know how to talk to other devices?

IP _addresses_ are associated with _networks_. IP _networks_ are contiguous logical ranges of addresses.

_Classless Inter-Domain Routing_ or _CIDR_ notation is used to distinguish network sizes.

- The _CIDR_ value, often notated starting with a **/**, represents the size of the network.

  - The formula _2 ^ (32 - CIDR)_ can be used to calculate network size.
  - Once network size is calculated, usable node count must be reduced by 2.
    - The first IP in a network is the _Network address_, and the last is typically the _Broadcast address_. These addresses are special and cannot be used by normal hosts.

**Wskazówka**

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

**Wskazówka**

It's a standard practice to make the _gateway_ the first or last IP on a network.

If an Internet-connected router is available at 192.168.0.1, it can be used as the _default route_, granting Internet access.

To summarize:

- Interfaces must be configured with an _address_ and _network information_, such as the _CIDR_ value.
- Local network access is used to access a _router_ on the same network.
- The _default route_ is configured, so traffic destined for external networks is forwarded to the _gateway_, providing Internet access.

### The Domain Name System

Remembering IPs is hard. The _Domain Name System_ was created to allow mapping between _Domain Names_ and _IP addresses_.

Linux systems use /etc/resolv.conf to define _nameservers_ to be used for _DNS resolution_.

**Wskazówka**

Many _routers_ can also function as a DNS server, and using a local DNS server can augment privacy and speed up queries through caching.

Many ISPs run a DNS server that is generally advertised to the _gateway_ over DHCP. Using a local DNS server tends to improve query latency, but most public DNS servers will return the same results, so server usage is largely based on preference.

## Ręczna konfiguracja sieci

### Interface address configuration

**Ważne**

When manually configuring IP addresses, the local network topology must be considered. IP addresses can be set arbitrarily; conflicts may cause network disruption.

To configure enp1s0 with the address 192.168.0.2 and CIDR /24:

`root #` `ip address add 192.168.0.2/24 dev enp1s0`

**Wskazówka**

The start of this command can be shortened to ip a.

### Default route configuration

Configuring address and network information for an interface will configure link routes, allowing communication with that network segment:

`root #` `ip route`

```
192.168.0.0/24 dev enp1s0 proto kernel scope link src 192.168.0.2
```

**Wskazówka**

This command can be shortened to ip r.

The default route can be set to 192.168.0.1 with:

`root #` `ip route add default via 192.168.0.1`

### DNS configuration

Nameserver info is typically acquired using DHCP, but can be set manually by adding `nameserver` entries to /etc/resolv.conf.

**Uwaga**

If _dhcpcd_ is running, changes to /etc/resolv.conf will not persist. Status can be checked with `ps x | grep dhcpcd`.

nano is included in Gentoo _boot media_ and can be used to edit /etc/resolv.conf with:

`root #` `nano /etc/resolv.conf`

Lines containing the keyword `nameserver` followed by a DNS server IP address are queried in order of definition:

FILE **`/etc/resolv.conf`** **Use Quad9 DNS.**

```
nameserver 9.9.9.9
nameserver 149.112.112.112

```

FILE **`/etc/resolv.conf`** **Use Cloudflare DNS.**

```
nameserver 1.1.1.1
nameserver 1.0.0.1

```

DNS status can be checked by pinging a domain name:

`root #` `ping -c 3 gentoo.org`

Once connectivity [has been verified](/wiki/Handbook:Alpha/Installation/Networking/pl#Testing_the_network "Handbook:Alpha/Installation/Networking/pl"), continue with [Preparing the disks](/wiki/Handbook:Alpha/Installation/Disks/pl "Handbook:Alpha/Installation/Disks/pl").

[← Choosing the media](/wiki/Handbook:Alpha/Installation/Media "Previous part") [Home](/wiki/Handbook:Alpha "Handbook:Alpha") [Preparing the disks →](/wiki/Handbook:Alpha/Installation/Disks "Next part")