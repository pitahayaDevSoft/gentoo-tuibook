# Networking

Other languages:

- [Deutsch](/wiki/Handbook:PPC64/Installation/Networking/de "Handbuch:PPC64/Installation/Netzwerk (100% translated)")
- [English](/wiki/Handbook:PPC64/Installation/Networking "Handbook:PPC64/Installation/Networking (100% translated)")
- [español](/wiki/Handbook:PPC64/Installation/Networking/es "Manual de Gentoo: PPC64/Instalación/Redes (100% translated)")
- [français](/wiki/Handbook:PPC64/Installation/Networking/fr "Handbook:PPC64/Installation/Networking/fr (100% translated)")
- [italiano](/wiki/Handbook:PPC64/Installation/Networking/it "Handbook:PPC64/Installation/Networking/it (100% translated)")
- [magyar](/wiki/Handbook:PPC64/Installation/Networking/hu "Handbook:PPC64/Installation/Networking/hu (100% translated)")
- [polski](/wiki/Handbook:PPC64/Installation/Networking/pl "Handbook:PPC64/Installation/Networking (100% translated)")
- [português do Brasil](/wiki/Handbook:PPC64/Installation/Networking/pt-br "Handbook:PPC64/Installation/Networking/pt-br (50% translated)")
- [русский](/wiki/Handbook:PPC64/Installation/Networking/ru "Handbook:PPC64/Installation/Networking (100% translated)")
- [தமிழ்](/wiki/Handbook:PPC64/Installation/Networking/ta "கையேடு:PPC64/நிறுவல்/வலையமைத்தல் (100% translated)")
- [中文（中国大陆）‎](/wiki/Handbook:PPC64/Installation/Networking/zh-cn "手册：PPC64/安装/配置网络 (100% translated)")
- [日本語](/wiki/Handbook:PPC64/Installation/Networking/ja "ハンドブック:PPC64/インストール/ネットワーク (100% translated)")
- 한국어

[PPC64 핸드북](/wiki/Handbook:PPC64/ko "Handbook:PPC64/ko")[설치](/wiki/Handbook:PPC64/Full/Installation/ko "Handbook:PPC64/Full/Installation/ko")[설치 정보](/wiki/Handbook:PPC64/Installation/About/ko "Handbook:PPC64/Installation/About/ko")[매체 선택](/wiki/Handbook:PPC64/Installation/Media/ko "Handbook:PPC64/Installation/Media/ko")네트워크 설정[디스크 준비](/wiki/Handbook:PPC64/Installation/Disks/ko "Handbook:PPC64/Installation/Disks/ko")[스테이지 3 설치](/wiki/Handbook:PPC64/Installation/Stage/ko "Handbook:PPC64/Installation/Stage/ko")[베이스 시스템 설치](/wiki/Handbook:PPC64/Installation/Base/ko "Handbook:PPC64/Installation/Base/ko")[커널 설정](/wiki/Handbook:PPC64/Installation/Kernel/ko "Handbook:PPC64/Installation/Kernel/ko")[시스템 설정](/wiki/Handbook:PPC64/Installation/System/ko "Handbook:PPC64/Installation/System/ko")[도구 설치](/wiki/Handbook:PPC64/Installation/Tools/ko "Handbook:PPC64/Installation/Tools/ko")[부트로더 설정](/wiki/Handbook:PPC64/Installation/Bootloader/ko "Handbook:PPC64/Installation/Bootloader/ko")[마무리](/wiki/Handbook:PPC64/Installation/Finalizing/ko "Handbook:PPC64/Installation/Finalizing/ko")[젠투 활용](/wiki/Handbook:PPC64/Full/Working/ko "Handbook:PPC64/Full/Working/ko")[포티지 소개](/wiki/Handbook:PPC64/Working/Portage/ko "Handbook:PPC64/Working/Portage/ko")[USE 플래그](/wiki/Handbook:PPC64/Working/USE/ko "Handbook:PPC64/Working/USE/ko")[포티지 기능](/wiki/Handbook:PPC64/Working/Features/ko "Handbook:PPC64/Working/Features/ko")[초기화 스크립트 시스템](/wiki/Handbook:PPC64/Working/Initscripts/ko "Handbook:PPC64/Working/Initscripts/ko")[환경 변수](/wiki/Handbook:PPC64/Working/EnvVar/ko "Handbook:PPC64/Working/EnvVar/ko")[포티지 활용](/wiki/Handbook:PPC64/Full/Portage/ko "Handbook:PPC64/Full/Portage/ko")[파일 및 디렉터리](/wiki/Handbook:PPC64/Portage/Files/ko "Handbook:PPC64/Portage/Files/ko")[변수](/wiki/Handbook:PPC64/Portage/Variables/ko "Handbook:PPC64/Portage/Variables/ko")[소프트웨어 브랜치 함께 사용하기](/wiki/Handbook:PPC64/Portage/Branches/ko "Handbook:PPC64/Portage/Branches/ko")[추가 도구](/wiki/Handbook:PPC64/Portage/Tools/ko "Handbook:PPC64/Portage/Tools/ko")[꾸러미 저장소 개별 설정](/wiki/Handbook:PPC64/Portage/CustomTree/ko "Handbook:PPC64/Portage/CustomTree/ko")[고급 기능](/wiki/Handbook:PPC64/Portage/Advanced/ko "Handbook:PPC64/Portage/Advanced/ko")[네트워크 설정](/wiki/Handbook:PPC64/Full/Networking/ko "Handbook:PPC64/Full/Networking/ko")[시작하기](/wiki/Handbook:PPC64/Networking/Introduction/ko "Handbook:PPC64/Networking/Introduction/ko")[고급 설정](/wiki/Handbook:PPC64/Networking/Advanced/ko "Handbook:PPC64/Networking/Advanced/ko")[모듈러 네트워크](/wiki/Handbook:PPC64/Networking/Modular/ko "Handbook:PPC64/Networking/Modular/ko")[무선 네트워크](/wiki/Handbook:PPC64/Networking/Wireless/ko "Handbook:PPC64/Networking/Wireless/ko")[기능 추가](/wiki/Handbook:PPC64/Networking/Extending/ko "Handbook:PPC64/Networking/Extending/ko")[동적 관리](/wiki/Handbook:PPC64/Networking/Dynamic/ko "Handbook:PPC64/Networking/Dynamic/ko")

## Contents

- [1자동 네트워크 감지](#.EC.9E.90.EB.8F.99_.EB.84.A4.ED.8A.B8.EC.9B.8C.ED.81.AC_.EA.B0.90.EC.A7.80)
  - [1.1DHCP 사용](#DHCP_.EC.82.AC.EC.9A.A9)
  - [1.2네트워크 시험](#.EB.84.A4.ED.8A.B8.EC.9B.8C.ED.81.AC_.EC.8B.9C.ED.97.98)
  - [1.3인터페이스 이름 결정](#.EC.9D.B8.ED.84.B0.ED.8E.98.EC.9D.B4.EC.8A.A4_.EC.9D.B4.EB.A6.84_.EA.B2.B0.EC.A0.95)
- [2Wireless Setup](#Wireless_Setup)
  - [2.1Optional: Checking wireless](#Optional:_Checking_wireless)
  - [2.2Using NetworkManager](#Using_NetworkManager)
- [3자동 네트워크 설정](#.EC.9E.90.EB.8F.99_.EB.84.A4.ED.8A.B8.EC.9B.8C.ED.81.AC_.EC.84.A4.EC.A0.95)
  - [3.1Manual setup](#Manual_setup)
    - [3.1.1Manually connecting to an open network](#Manually_connecting_to_an_open_network)
- [4Application specific configuration](#Application_specific_configuration)
  - [4.1선택: 프록시 설정](#.EC.84.A0.ED.83.9D:_.ED.94.84.EB.A1.9D.EC.8B.9C_.EC.84.A4.EC.A0.95)
  - [4.2대안: PPP 사용](#.EB.8C.80.EC.95.88:_PPP_.EC.82.AC.EC.9A.A9)
  - [4.3대안: PPTP 사용](#.EB.8C.80.EC.95.88:_PPTP_.EC.82.AC.EC.9A.A9)
  - [4.4네트워크 용어 이해](#.EB.84.A4.ED.8A.B8.EC.9B.8C.ED.81.AC_.EC.9A.A9.EC.96.B4_.EC.9D.B4.ED.95.B4)
  - [4.5Interfaces and addresses](#Interfaces_and_addresses)
  - [4.6Networks and CIDR](#Networks_and_CIDR)
  - [4.7The Internet](#The_Internet)
  - [4.8The Domain Name System](#The_Domain_Name_System)
- [5직접 네트워크 설정](#.EC.A7.81.EC.A0.91_.EB.84.A4.ED.8A.B8.EC.9B.8C.ED.81.AC_.EC.84.A4.EC.A0.95)
  - [5.1Interface address configuration](#Interface_address_configuration)
  - [5.2Default route configuration](#Default_route_configuration)
  - [5.3DNS configuration](#DNS_configuration)

## 자동 네트워크 감지

아마도 바로 동작하겠죠?

시스템을 DHCP 서버가 붙은 이더넷에 연결했다면, 네트워크 설정은 거의 자동으로 이루어집니다. ssh, scp, ping, irssi, wget, links 등, 설치 CD에 들어있는 대부분의 네트워크 관련 명령 역시 바로 동작합니다.

### DHCP 사용

DHCP(동적 호스트 설정 프로토콜)은 네트워크 정보(IP주소, 네트워크 마스크, 브로드캐스트 주소, 게이트웨이, 네임서버 등)를 자동으로 받을 수 있게 합니다. DHCP 서버가 네트워크에 있을 때(또는 ISP 서비스 업체에서 DHCP 서비스를 제공할 때)만 동작합니다. 네트워크 인터페이스가 이 정보를 자동으로 받게 하려면, dhcpcd를 사용하십시오:

DHCP requires that a server be running on the same _Layer 2_ ( _Ethernet_) segment as the client requesting a _lease_. DHCP is often used on RFC1918 ( _private_) networks, but is also used to acquire public IP information from ISPs.

By default the Gentoo live media uses [NetworkManager](/wiki/NetworkManager "NetworkManager") which will automatically connect to via DHCP, if this is not the case then check the Ethernet cable for issues.

### 네트워크 시험

A properly configured _default_ route is a critical component of Internet connectivity, route configuration can be checked with:

`root #` `ip route`

```
default via 192.168.0.1 dev enp1s0
```

If no _default_ route is defined, Internet connectivity is unavailable, and additional configuration is required.

Basic internet connectivity can be confirmed with a ping:

`root #` `ping -c 3 1.1.1.1`

**요령**

It's helpful to start by pinging a known IP address instead of a hostname. This can isolate DNS issues from basic Internet connectivity issues.

Outbound HTTPS access and DNS resolution can be confirmed with:

`root #` `curl --location gentoo.org --output /dev/null`

모든 기능이 제대로, 이 장의 나머지를 건너뛰고 바로 다음 단계 설치 과정 ( [디스크 준비](/wiki/Handbook:PPC64/Installation/Disks/ko "Handbook:PPC64/Installation/Disks/ko"))으로 진행할 수 있습니다.

If curl reports an error, but Internet-bound pings work, [DNS may need configuration](/wiki/Handbook:PPC64/Installation/Networking/ko#DNS_configuration "Handbook:PPC64/Installation/Networking/ko").

If Internet connectivity has not been established, first [interface information should be verified](/wiki/Handbook:PPC64/Installation/Networking/ko#Obtaining_interface_info "Handbook:PPC64/Installation/Networking/ko"), then:

- [nmtui can be used](/wiki/Handbook:PPC64/Installation/Networking/ko#Using_NetworkManager "Handbook:PPC64/Installation/Networking/ko") to assist in network configuration.
- [Application specific configuration](/wiki/Handbook:PPC64/Installation/Networking/ko#Application_specific_configuration "Handbook:PPC64/Installation/Networking/ko") may be required.
- [Manual network configuration](/wiki/Handbook:PPC64/Installation/Networking/ko#Manual_network_configuration "Handbook:PPC64/Installation/Networking/ko") can be attempted.

### 인터페이스 이름 결정

If networking doesn't work out of the box, additional steps must be taken to enable Internet connectivity. Generally, the first step is to enumerate host network interfaces.

ifconfig의 대안 수단으로 ip 명령을 인터페이스 이름으로 결정하요 사용할 수 있습니다. 다음 예제에서는 ip addr 출력 내용(은 다른 시스템의 내용이며 이전 내용과 조금 다름)을 보여줍니다:

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

**요령**

ifconfig 명령을 사용했을 때 인터페이스가 나타나지 않으면, 동일한 명령에 `-a` 옵션을 사용해보십시오. 이 옵션은 유틸리티에 시스템에서 발견한 모든 인터페이스의 가동 여부를 표시하도록 강제합니다. ifconfig -a 출력에 내용이 나타나지 않으면 하드웨어에 문제가 있거나 인터페이스 드라이버를 커널에 불러오지 않았음음을 의미합니다. 두 경우는 이 핸드북의 주제 범위를 벗어납니다. [#gentoo](ircs://irc.libera.chat/#gentoo) ([webchat](https://web.libera.chat/#gentoo))에 지원을 문의하십시오.

이 문서 나머지 과정에서, 핸드북은 동작하는 네트워크 인터페이스를 eth0라 하겠습니다.

[유추 가능 인터페이스 이름](http://www.freedesktop.org/wiki/Software/systemd/PredictableNetworkInterfaceNames/) 으로 추세가 이동함에 따라, 시스템에 있는 인터페이스 이름은 이전에 사용하던 eth0 이름 부여 방식과 약간 다를 수 있습니다. 최근 설치 미디어에서는 eno0, ens1, enp5s0와 같은 규칙적인 네트워크 인터페이스 이름을 표시합니다. ifconfig 출력에서 로컬 네트워크와 관련된 IP 주소와 함께 네트워크 인터페이스를 찾아 나타냄을 살펴보십시오.

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

**참고**

Some wireless cards may have a device name of wlan0 or ra0 instead of wlp9s0. Run ip link to determine the correct device name.

### Using NetworkManager

The fastest way to enable wireless connectivity in the Gentoo Live media is by using the nmtui command and following the on-screen instructions to setup the wireless network.

`root #` `nmtui`

LiveGUI users can click on the network icon in the bottom right of the taskbar to also configure WiFi.

## 자동 네트워크 설정

On some architectures, e.g. **HPPA**, a user may be required to use the tool net-setup to configure a wireless connection if NetworkManager isn't available yet.

`root #` `net-setup eth0`

net-setup에서는 네트워크 환경에 대한 일부 사항을 질문합니다. 모든 과정이 끝나면 네트워크 연결은 동작해야 합니다. 네트워크 연결 시험 방법은 앞서 언급했습니다. 시험 결과가 긍정적이라면 축하드립니다! 이 절의 나머지 부분을 건너뛰고 [디스크 준비](/wiki/Handbook:PPC64/Installation/Disks/ko "Handbook:PPC64/Installation/Disks/ko") 로 계속 진행하십시오.

**중요**

Network status should be [tested](/wiki/Handbook:PPC64/Installation/Networking/ko#Testing_the_network "Handbook:PPC64/Installation/Networking/ko") after any configuration steps are taken. In the event that the configuration scripts do not work, [manual network configuration](/wiki/Handbook:PPC64/Installation/Networking/ko#Manual_network_configuration "Handbook:PPC64/Installation/Networking/ko") is required.

### Manual setup

Wireless networks may alternatively be set up with a daemon such as wpa\_supplicant or iwd. For more information on configuring wireless networking in Gentoo Linux, please read the [Wireless networking chapter](/wiki/Handbook:PPC64/Networking/Wireless/ko "Handbook:PPC64/Networking/Wireless/ko") in the Gentoo Handbook.}}

#### Manually connecting to an open network

For most users, there are only two settings needed to connect, the ESSID (aka wireless network name) and, optionally, the WEP key.

- First, ensure the interface is active:

`root #` `ip link set dev wlp9s0 up`

- To connect to an open network with the name _GentooNode_:

`root #` `iw dev wlp9s0 connect -w GentooNode`

## Application specific configuration

**요령**

These steps are provided for users where using nmtui is not able to configure their network's needs.

The following methods are not generally required, but may be helpful in situations where additional configuration is required for Internet connectivity.

### 선택: 프록시 설정

인터넷을 프록시로 연결했다면 설치 과정에 프록시를 설정해야 합니다. 프록시 설정은 정말 쉽습니다. 프록시 서버 정보를 넣을 변수를 설정하기만 하면 됩니다.

Certain text-mode web browsers such as links can also make use of environment variables that define web proxy settings; in particular for the HTTPS access it also will require the https\_proxy environment variable to be defined. While Portage will be influenced without passing extra run time parameters during invocation, links will require proxy settings to be set.

대부분의 경우, 서버의 호스트 이름을 사용하여 변수를 정의하는 것으로 충분합니다. 예제에서는 proxy.gentoo.org라는 프록시 서버와 8080포트를 사용한다고 가정하겠습니다.

**참고**

The `#` symbol in the following commands is a comment. It has been added for clarity only and does _not_ need to be typed when entering the commands.

HTTP 프록시(HTTP와 HTTPS 트래픽용)를 설정하려면:

`root #` `export http_proxy="http://proxy.gentoo.org:8080"`

프록시에 사용자 이름과 암호가 필요하다면, 변수에 다음 문법을 사용하십시오:

코드 **프록시 변수에 사용자 이름/암호 추가**

```
http://username:password@proxy.gentoo.org:8080

```

Start links using the following parameters for proxy support:

`user $` `links -http-proxy ${http_proxy} -https-proxy ${https_proxy} `

FTP 프록시를 설정하려면:

`root #` `export ftp_proxy="ftp://proxy.gentoo.org:8080"`

Start links using the following parameter for a FTP proxy:

`user $` `links -ftp-proxy ${ftp_proxy} `

RSYNC 프록시를 설정하려면:

`root #` `export RSYNC_PROXY="proxy.gentoo.org:8080"`

### 대안: PPP 사용

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

### 대안: PPTP 사용

PPTP 지원이 필요하다면, 설치 CD에서 제공하는 pptpclient를 사용하십시오. 그러나 우선은 설정이 올바른지부터 확인하십시오. /etc/ppp/pap-secrets 또는 /etc/ppp/chap-secrets 파일을 편집하여 올바른 사용자 이름과 암호 조합이 들어가도록 하십시오:

Edit /etc/ppp/pap-secrets or /etc/ppp/chap-secrets so it contains the correct username/password combination:

`root #` `nano -w /etc/ppp/chap-secrets`

다음, 필요한 경우 /etc/ppp/options.pptp를 편집하십시오:

`root #` `nano -w /etc/ppp/options.pptp`

모든 조치가 완료되었다면, (options.pptp에 설정할 수 없던 옵션으로) pptp를 실행하여 서버에 연걸하십시오:

`root #` `pptp <server ip>`

### 네트워크 용어 이해

If all of the above fails, the network must be configured manually. This is not particularly difficult, but should be done with consideration. This section serves to clarify terminology and introduce users to basic networking concepts pertaining to manually configuring an Internet connection.

**요령**

Some **CPE** ( **Carrier Provided Equipment**) combines the functions of a _router_, _access point_, _modem_, _DHCP server_, and _DNS server_ into one unit. It's important to differentiate the functions of a device from the physical appliance.

### Interfaces and addresses

Network _interfaces_ are logical representations of network devices. An _interface_ needs an _address_ to communicate with other devices on the _network_. While only a single _address_ is required, multiple addresses can be assigned to a single _interface_. This is especially useful for dual stack (IPv4 + IPv6) configurations.

For consistency, this primer will assume the interface enp1s0 will be using the address 192.168.0.2.

**중요**

IP addresses can be set arbitrarily. As a result, it's possible for multiple devices to use the same IP address, resulting in an _address conflict_. Address conflicts should be avoided by using DHCP or SLAAC.

**요령**

IPv6 typically uses **S** tate **L** ess **A** ddress **A** uto **C** onfiguration (SLAAC) for address configuration. In most cases, manually setting IPv6 addresses is a bad practice. If a specific address suffix is preferred, [interface identification tokens](/wiki/IPv6_Static_Addresses_using_Tokens "IPv6 Static Addresses using Tokens") can be used.

### Networks and CIDR

Once an address is chosen, how does a device know how to talk to other devices?

IP _addresses_ are associated with _networks_. IP _networks_ are contiguous logical ranges of addresses.

_Classless Inter-Domain Routing_ or _CIDR_ notation is used to distinguish network sizes.

- The _CIDR_ value, often notated starting with a **/**, represents the size of the network.

  - The formula _2 ^ (32 - CIDR)_ can be used to calculate network size.
  - Once network size is calculated, usable node count must be reduced by 2.
    - The first IP in a network is the _Network address_, and the last is typically the _Broadcast address_. These addresses are special and cannot be used by normal hosts.

**요령**

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

**요령**

It's a standard practice to make the _gateway_ the first or last IP on a network.

If an Internet-connected router is available at 192.168.0.1, it can be used as the _default route_, granting Internet access.

To summarize:

- Interfaces must be configured with an _address_ and _network information_, such as the _CIDR_ value.
- Local network access is used to access a _router_ on the same network.
- The _default route_ is configured, so traffic destined for external networks is forwarded to the _gateway_, providing Internet access.

### The Domain Name System

Remembering IPs is hard. The _Domain Name System_ was created to allow mapping between _Domain Names_ and _IP addresses_.

Linux systems use /etc/resolv.conf to define _nameservers_ to be used for _DNS resolution_.

**요령**

Many _routers_ can also function as a DNS server, and using a local DNS server can augment privacy and speed up queries through caching.

Many ISPs run a DNS server that is generally advertised to the _gateway_ over DHCP. Using a local DNS server tends to improve query latency, but most public DNS servers will return the same results, so server usage is largely based on preference.

## 직접 네트워크 설정

### Interface address configuration

**중요**

When manually configuring IP addresses, the local network topology must be considered. IP addresses can be set arbitrarily; conflicts may cause network disruption.

To configure enp1s0 with the address 192.168.0.2 and CIDR /24:

`root #` `ip address add 192.168.0.2/24 dev enp1s0`

**요령**

The start of this command can be shortened to ip a.

### Default route configuration

Configuring address and network information for an interface will configure link routes, allowing communication with that network segment:

`root #` `ip route`

```
192.168.0.0/24 dev enp1s0 proto kernel scope link src 192.168.0.2
```

**요령**

This command can be shortened to ip r.

The default route can be set to 192.168.0.1 with:

`root #` `ip route add default via 192.168.0.1`

### DNS configuration

Nameserver info is typically acquired using DHCP, but can be set manually by adding `nameserver` entries to /etc/resolv.conf.

**경고**

If _dhcpcd_ is running, changes to /etc/resolv.conf will not persist. Status can be checked with `ps x | grep dhcpcd`.

nano is included in Gentoo _boot media_ and can be used to edit /etc/resolv.conf with:

`root #` `nano -w /etc/resolv.conf`

Lines containing the keyword `nameserver` followed by a DNS server IP address are queried in order of definition:

파일 **`/etc/resolv.conf`** **Use Quad9 DNS.**

```
nameserver 9.9.9.9
nameserver 149.112.112.112

```

파일 **`/etc/resolv.conf`** **Use Cloudflare DNS.**

```
nameserver 1.1.1.1
nameserver 1.0.0.1

```

DNS status can be checked by pinging a domain name:

`root #` `ping -c 3 gentoo.org`

Once connectivity [has been verified](/wiki/Handbook:PPC64/Installation/Networking/ko#Testing_the_network "Handbook:PPC64/Installation/Networking/ko"), continue with [Preparing the disks](/wiki/Handbook:PPC64/Installation/Disks/ko "Handbook:PPC64/Installation/Disks/ko").

[← 올바른 매체 선택](/wiki/Handbook:PPC64/Installation/Media/ko "이전 내용") [처음](/wiki/Handbook:PPC64/ko "Handbook:PPC64/ko") [디스크 준비 →](/wiki/Handbook:PPC64/Installation/Disks/ko "다음 내용")