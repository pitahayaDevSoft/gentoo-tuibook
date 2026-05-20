# Networking

Other languages:

- [Deutsch](/wiki/Handbook:Alpha/Installation/Networking/de "Handbuch:Alpha/Installation/Netzwerk (100% translated)")
- [English](/wiki/Handbook:Alpha/Installation/Networking "Handbook:Alpha/Installation/Networking (100% translated)")
- [Türkçe](/wiki/Handbook:Alpha/Installation/Networking/tr "Handbook:Alpha/Installation/Networking/tr (0% translated)")
- [español](/wiki/Handbook:Alpha/Installation/Networking/es "Manual de Gentoo: Alpha/Instalación/Redes (100% translated)")
- [français](/wiki/Handbook:Alpha/Installation/Networking/fr "Handbook:Alpha/Installation/Networking/fr (100% translated)")
- [italiano](/wiki/Handbook:Alpha/Installation/Networking/it "Handbook:Alpha/Installation/Networking/it (100% translated)")
- magyar
- [polski](/wiki/Handbook:Alpha/Installation/Networking/pl "Handbook:Alpha/Installation/Networking (100% translated)")
- [português do Brasil](/wiki/Handbook:Alpha/Installation/Networking/pt-br "Manual:Alpha/Instalação/Rede (100% translated)")
- [čeština](/wiki/Handbook:Alpha/Installation/Networking/cs "Handbook:Alpha/Installation/Networking/cs (50% translated)")
- [русский](/wiki/Handbook:Alpha/Installation/Networking/ru "Handbook:Alpha/Installation/Networking (100% translated)")
- [தமிழ்](/wiki/Handbook:Alpha/Installation/Networking/ta "கையேடு:Alpha/நிறுவல்/வலையமைத்தல் (100% translated)")
- [中文（中国大陆）‎](/wiki/Handbook:Alpha/Installation/Networking/zh-cn "手册：Alpha/安装/配置网络 (100% translated)")
- [日本語](/wiki/Handbook:Alpha/Installation/Networking/ja "ハンドブック:Alpha/インストール/ネットワーク (100% translated)")
- [한국어](/wiki/Handbook:Alpha/Installation/Networking/ko "Handbook:Alpha/Installation/Networking/ko (100% translated)")

[Alpha kézikönyv](/wiki/Handbook:Alpha/hu "Handbook:Alpha/hu")[A Gentoo Linux telepítése](/wiki/Handbook:Alpha/Full/Installation/hu "Handbook:Alpha/Full/Installation/hu")[A telepítésről](/wiki/Handbook:Alpha/Installation/About/hu "Handbook:Alpha/Installation/About/hu")[Telepítőképfájl kiválasztása](/wiki/Handbook:Alpha/Installation/Media/hu "Handbook:Alpha/Installation/Media/hu")Hálózat beállítása[Adathordozók előkészítése](/wiki/Handbook:Alpha/Installation/Disks/hu "Handbook:Alpha/Installation/Disks/hu")[Stage fájl](/wiki/Handbook:Alpha/Installation/Stage/hu "Handbook:Alpha/Installation/Stage/hu")[Alaprendszer telepítése](/wiki/Handbook:Alpha/Installation/Base/hu "Handbook:Alpha/Installation/Base/hu")[Kernel beállítása](/wiki/Handbook:Alpha/Installation/Kernel/hu "Handbook:Alpha/Installation/Kernel/hu")[Rendszer beállítása](/wiki/Handbook:Alpha/Installation/System/hu "Handbook:Alpha/Installation/System/hu")[Eszközök telepítése](/wiki/Handbook:Alpha/Installation/Tools/hu "Handbook:Alpha/Installation/Tools/hu")[Bootloader beállítása](/wiki/Handbook:Alpha/Installation/Bootloader/hu "Handbook:Alpha/Installation/Bootloader/hu")[Telepítés véglegesítése](/wiki/Handbook:Alpha/Installation/Finalizing/hu "Handbook:Alpha/Installation/Finalizing/hu")[Munka a Gentoo rendszerrel](/wiki/Handbook:Alpha/Full/Working/hu "Handbook:Alpha/Full/Working/hu")[Portage bemutatása](/wiki/Handbook:Alpha/Working/Portage/hu "Handbook:Alpha/Working/Portage/hu")[USE jelölőzászlók](/wiki/Handbook:Alpha/Working/USE/hu "Handbook:Alpha/Working/USE/hu")[Portage jellemzői](/wiki/Handbook:Alpha/Working/Features/hu "Handbook:Alpha/Working/Features/hu")[Init-szkript rendszer](/wiki/Handbook:Alpha/Working/Initscripts/hu "Handbook:Alpha/Working/Initscripts/hu")[Környezeti változók](/wiki/Handbook:Alpha/Working/EnvVar/hu "Handbook:Alpha/Working/EnvVar/hu")[Munka a Portage szoftvercsomag-kezelővel](/wiki/Handbook:Alpha/Full/Portage/hu "Handbook:Alpha/Full/Portage/hu")[Fájlok és könyvtárak](/wiki/Handbook:Alpha/Portage/Files/hu "Handbook:Alpha/Portage/Files/hu")[Változók](/wiki/Handbook:Alpha/Portage/Variables/hu "Handbook:Alpha/Portage/Variables/hu")[Szoftverágak keverése](/wiki/Handbook:Alpha/Portage/Branches/hu "Handbook:Alpha/Portage/Branches/hu")[További eszközök](/wiki/Handbook:Alpha/Portage/Tools/hu "Handbook:Alpha/Portage/Tools/hu")[Egyéni szoftvercsomag-tárolók](/wiki/Handbook:Alpha/Portage/CustomTree/hu "Handbook:Alpha/Portage/CustomTree/hu")[Fejlett funkciók](/wiki/Handbook:Alpha/Portage/Advanced/hu "Handbook:Alpha/Portage/Advanced/hu")[Hálózat beállítása OpenRC init-rendszeren](/wiki/Handbook:Alpha/Full/Networking/hu "Handbook:Alpha/Full/Networking/hu")[Munka elkezdése](/wiki/Handbook:Alpha/Networking/Introduction/hu "Handbook:Alpha/Networking/Introduction/hu")[Fejlett beállítások](/wiki/Handbook:Alpha/Networking/Advanced/hu "Handbook:Alpha/Networking/Advanced/hu")[Moduláris hálózat](/wiki/Handbook:Alpha/Networking/Modular/hu "Handbook:Alpha/Networking/Modular/hu")[Vezeték nélküli (Wi-Fi)](/wiki/Handbook:Alpha/Networking/Wireless/hu "Handbook:Alpha/Networking/Wireless/hu")[Funkcionalitás hozzáadása](/wiki/Handbook:Alpha/Networking/Extending/hu "Handbook:Alpha/Networking/Extending/hu")[Dinamikus menedzsment](/wiki/Handbook:Alpha/Networking/Dynamic/hu "Handbook:Alpha/Networking/Dynamic/hu")

## Contents

- [1Automatikusan megtörténő hálózati beállítás](#Automatikusan_megt.C3.B6rt.C3.A9n.C5.91_h.C3.A1l.C3.B3zati_be.C3.A1ll.C3.ADt.C3.A1s)
  - [1.1DHCP használata](#DHCP_haszn.C3.A1lata)
  - [1.2Hálózat tesztelése](#H.C3.A1l.C3.B3zat_tesztel.C3.A9se)
- [2Interfész információinak a megszerzése](#Interf.C3.A9sz_inform.C3.A1ci.C3.B3inak_a_megszerz.C3.A9se)
- [3Wireless Setup](#Wireless_Setup)
  - [3.1Optional: Checking wireless](#Optional:_Checking_wireless)
  - [3.2Using NetworkManager](#Using_NetworkManager)
- [4A net-setup segédalkalmazás használata](#A_net-setup_seg.C3.A9dalkalmaz.C3.A1s_haszn.C3.A1lata)
  - [4.1Manual setup](#Manual_setup)
    - [4.1.1Manually connecting to an open network](#Manually_connecting_to_an_open_network)
- [5Opcionális: Alkalmazásspecifikus beállítás](#Opcion.C3.A1lis:_Alkalmaz.C3.A1sspecifikus_be.C3.A1ll.C3.ADt.C3.A1s)
  - [5.1Webproxyk beállítása](#Webproxyk_be.C3.A1ll.C3.ADt.C3.A1sa)
  - [5.2A pppoe-setup használata az ADSL internetkapcsolathoz](#A_pppoe-setup_haszn.C3.A1lata_az_ADSL_internetkapcsolathoz)
  - [5.3PPTP protokoll használata](#PPTP_protokoll_haszn.C3.A1lata)
- [6Internet és IP alapismeretek](#Internet_.C3.A9s_IP_alapismeretek)
  - [6.1Interfészek és címek](#Interf.C3.A9szek_.C3.A9s_c.C3.ADmek)
  - [6.2Hálózatok és a CIDR](#H.C3.A1l.C3.B3zatok_.C3.A9s_a_CIDR)
  - [6.3Internet](#Internet)
  - [6.4Domain Name System](#Domain_Name_System)
- [7Hálózat kézi úton történő beállítása](#H.C3.A1l.C3.B3zat_k.C3.A9zi_.C3.BAton_t.C3.B6rt.C3.A9n.C5.91_be.C3.A1ll.C3.ADt.C3.A1sa)
  - [7.1Hálózati interfész címének a beállítása](#H.C3.A1l.C3.B3zati_interf.C3.A9sz_c.C3.ADm.C3.A9nek_a_be.C3.A1ll.C3.ADt.C3.A1sa)
  - [7.2Alapértelmezett útvonal beállítása](#Alap.C3.A9rtelmezett_.C3.BAtvonal_be.C3.A1ll.C3.ADt.C3.A1sa)
  - [7.3DNS beállítása](#DNS_be.C3.A1ll.C3.ADt.C3.A1sa)

## Automatikusan megtörténő hálózati beállítás

Talán már ebben a pillanatban önmagától működik a hálózat?

Ha az Gentoo operációs rendszer IPv6-útválasztóval vagy DHCP-kiszolgálóval csatlakozik az Ethernet-hálózathoz, akkor nagyon valószínű, hogy a Gentoo operációs rendszer hálózata automatikusan be lett állítva. Ha nincs szükség további speciális beállításra, akkor innentől kezdve [tesztelhető az internetkapcsolat](/wiki/Handbook:Alpha/Installation/Networking/hu#Testing_the_network "Handbook:Alpha/Installation/Networking/hu").

### DHCP használata

A DHCP (Dynamic Host Configuration Protocol) segít a hálózat beállításában, és automatikusan beállíthat számos paramétert, beleértve az IP-címet, hálózati maszkot, útvonalakat, DNS-kiszolgálókat, NTP-kiszolgálókat, stb.

A DHCP protokoll a tervezéséből adódóan megköveteli, hogy a DHCP-szerver ugyanazon a _Layer 2_ ( _Ethernet_) szegmensen fusson, mint a _bérletet_ kérő DHCP-kliens. A DHCP protokollt gyakran használják RFC1918 (privát) hálózatokban, de arra is használják, hogy vele nyilvános IP-információkat kérjenek le az internetszolgáltatóktól.

By default the Gentoo live media uses [NetworkManager](/wiki/NetworkManager "NetworkManager") which will automatically connect to via DHCP, if this is not the case then check the Ethernet cable for issues.

### Hálózat tesztelése

A megfelelően beállított _alapértelmezett_ útvonal az internetkapcsolat kritikus összetevője. Az útvonal-beállítás a következőkkel ellenőrizhető:

`root #` `ip route`

```
default via 192.168.0.1 dev enp1s0
```

Ha nincs _alapértelmezett_ útvonal megadva, akkor az internetkapcsolat nem érhető el, és további beállításra van szükség.

Az alapvető internetkapcsolat megléte a ping paranccsal ellenőrizhető le:

`root #` `ping -c 3 1.1.1.1`

**Tip**

Hasznos, ha Ön először egy ismert IP-címet pingel a hostname helyett. Ez elkülönítheti a DNS-problémákat az alapvető internetkapcsolat problémáitól.

A kimenő HTTPS hozzáférés és a DNS-névfeloldás meglétéről a következőképpen győződhet meg:

`root #` `curl --location gentoo.org --output /dev/null`

Hacsak a curl hibát nem jelez, vagy más tesztek sikertelenek, akkor a telepítési folyamat folytatható a [lemez előkészítésével](/wiki/Handbook:Alpha/Installation/Disks/hu "Handbook:Alpha/Installation/Disks/hu").

Ha a curl hibát jelez, de az internethez kötött pingek működnek, akkor előfordulhat, hogy [be kell állítani a DNS-t](/wiki/Handbook:Alpha/Installation/Networking/hu#DNS_configuration "Handbook:Alpha/Installation/Networking/hu").

Ha az internetkapcsolat nem jött létre, akkor először [ellenőrizni kell az interfész információit](/wiki/Handbook:Alpha/Installation/Networking/hu#Obtaining_interface_info "Handbook:Alpha/Installation/Networking/hu"), majd:

- A [net-setup használata](/wiki/Handbook:Alpha/Installation/Networking/hu#Using_net-setup "Handbook:Alpha/Installation/Networking/hu") segíthet a hálózat beállításában.
- [Alkalmazásspecifikus beállításra](/wiki/Handbook:Alpha/Installation/Networking/hu#Application_specific_configuration "Handbook:Alpha/Installation/Networking/hu") lehet szükség.
- Megpróbálható a [hálózat manuális beállítása](/wiki/Handbook:Alpha/Installation/Networking/hu#Manual_network_configuration "Handbook:Alpha/Installation/Networking/hu").

## Interfész információinak a megszerzése

Ha a hálózat automatikusan nem működőképes, akkor további lépéseket kell tenni az internetkapcsolat engedélyezéséhez. Általában az első lépés a gazdagépen lévő hálózati interfészek felsorolása.

A [sys-apps/iproute2](https://packages.gentoo.org/packages/sys-apps/iproute2) szoftvercsomag részét képező ip parancs használható a rendszerünk hálózatának a lekérdezésére és beállítására.

A **link** argumentum használható hálózati-interfész linkjeinek a kilistázására:

`root #` `ip link`

```
1: lo: <LOOPBACK,UP,LOWER_UP> mtu 65536 qdisc noqueue state UNKNOWN mode DEFAULT group default qlen 1000
    link/loopback 00:00:00:00:00:00 brd 00:00:00:00:00:00
4: enp1s0: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc pfifo_fast state UP mode DEFAULT group default qlen 1000
    link/ether e8:40:f2:ac:25:7a brd ff:ff:ff:ff:ff:ff
```

A **address** argumentum használható az hálózati-eszköz címére vonatkozó információk kilistázására:

`root #` `ip address`

```
2: enp1s0: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc pfifo_fast state UP group default qlen 1000<pre>
    link/ether e8:40:f2:ac:25:7a brd ff:ff:ff:ff:ff:ff
    inet 10.0.20.77/22 brd 10.0.23.255 scope global enp1s0
       valid_lft forever preferred_lft forever
    inet6 fe80::ea40:f2ff:feac:257a/64 scope link
       valid_lft forever preferred_lft forever
```

Ennek a parancsnak a kimenete a rendszer minden hálózati-interfészére vonatkozó információját tartalmazza. A bejegyzések sorai az eszközindexszel kezdődnek, amiket az eszköznevek követnek: enp1s0.

**Tip**

Ha a **lo** (loopack) interfészen kívül más interfész nem jelenik meg, akkor a hálózati hardver (maga a hálózati-kártya ha kivehető, vagy hálózati-chip a számítógép alaplapján) hibás, vagy az interfész illesztőprogramja nincs betöltve a kernelbe. Mindkét helyzet túlmutat e kézikönyv keretein. Kérjen segítséget a [#gentoo](ircs://irc.libera.chat/#gentoo) ([webchat](https://web.libera.chat/#gentoo)) elérhetőségen.

A következetesség érdekében a kézikönyv feltételezi, hogy az elsődleges hálózati interfész neve enp1s0.

**Note**

A [kiszámítható hálózati interfésznevek felé való elmozdulás eredményeként](https://www.freedesktop.org/wiki/Software/systemd/PredictableNetworkInterfaceNames/) a rendszer interfészneve egészen más lehet, mint a régi eth0 elnevezési konvenció. A modern Gentoo rendszerindító adathordozók interfészneveket használnak előtagokkal, például eno0, ens1 vagy enp5s0.

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

**Note**

Some wireless cards may have a device name of wlan0 or ra0 instead of wlp9s0. Run ip link to determine the correct device name.

### Using NetworkManager

The fastest way to enable wireless connectivity in the Gentoo Live media is by using the nmtui command and following the on-screen instructions to setup the wireless network.

`root #` `nmtui`

LiveGUI users can click on the network icon in the bottom right of the taskbar to also configure WiFi.

## A net-setup segédalkalmazás használata

A Gentoo _bootolható adathordozója_ a hálózat beállításához Önnek szkripteket biztosít azokra az esetekre, amikor az automatikusan történő hálózati beállítás folyamata sikertelenül megy végbe. A net-setup nevű szkript a vezeték nélküli hálózati információk és a statikus IP-címek beállítására használható fel.

`root #` `net-setup enp1s0`

A net-setup szkript néhány kérdést tesz fel a hálózati környezettel kapcsolatban, és ezeket az információkat felhasználja a wpa\_supplicant vagy a _statikus címzés_ beállítására.

**Important**

A hálózat állapotát minden beállítási lépés után [tesztelni](/wiki/Handbook:Alpha/Installation/Networking/hu#Testing_the_network "Handbook:Alpha/Installation/Networking/hu") kell. Abban az esetben, ha a beállító szkriptek nem működnek mefelelően, akkor [kézi úton történő hálózati beállításra](/wiki/Handbook:Alpha/Installation/Networking/hu#Manual_network_configuration "Handbook:Alpha/Installation/Networking/hu") van szükség.

### Manual setup

Wireless networks may alternatively be set up with a daemon such as wpa\_supplicant or iwd. For more information on configuring wireless networking in Gentoo Linux, please read the [Wireless networking chapter](/wiki/Handbook:Alpha/Networking/Wireless/hu "Handbook:Alpha/Networking/Wireless/hu") in the Gentoo Handbook.}}

#### Manually connecting to an open network

For most users, there are only two settings needed to connect, the ESSID (aka wireless network name) and, optionally, the WEP key.

- First, ensure the interface is active:

`root #` `ip link set dev wlp9s0 up`

- To connect to an open network with the name _GentooNode_:

`root #` `iw dev wlp9s0 connect -w GentooNode`

## Opcionális: Alkalmazásspecifikus beállítás

**Tip**

These steps are provided for users where using nmtui is not able to configure their network's needs.

A következő módszerek általában nem szükségesek, de hasznosak lehetnek olyan helyzetekben, amikor további beállításokra van szükség az internetkapcsolathoz.

### Webproxyk beállítása

Ha az internetet webproxyn keresztül kell elérni, akkor meg kell adni a proxyinformációkat, hogy a Portage szoftvercsomag-kezelő a proxyn keresztül férjen hozzá mindegyik támogatott protokoll esetében. A Portage szoftvercsomag-kezelő figyeli a http\_proxy, ftp\_proxy, és a RSYNC\_PROXY környezeti változókat, hogy szoftvercsomagokat tölthessen le a wget és a rsync lekérési mechanizmusain keresztül.

Egyes parancssorban működő webböngészők, mint például a links, környezeti változókat is használhatnak, amelyek meghatározzák a webproxy beállításait. Különösen a HTTPS hozzáféréshez meg kell adni a https\_proxy környezeti változót is. Míg a Portage szoftvercsomag-kezelő a futtatása során további futásidőben történő paraméter átadás nélkül lesz befolyásolva, addig a links böngészőnek meg kell adni a proxy beállításokat.

A legtöbb esetben elegendő a környezeti változók meghatározása a szerver hostname -vel. A következő példában feltételezzük, hogy a proxyszerver host neve proxy.gentoo.org, a portja pedig 8080.

**Note**

A következő parancsok utáni `#` szimbólum csak egy megjegyzés. Az egyértelműség kedvéért került hozzáadásra, és _nem_ kell beírni a parancsok begépelésekor.

HTTP-proxy megadása (HTTP és HTTPS-forgalomhoz):

`root #` `export http_proxy="http://proxy.gentoo.org:8080" # A Portage szoftvercsomag-kezelőre és a Links webböngészőre vonatkozik.
`

`root #` `export https_proxy="http://proxy.gentoo.org:8080" # Kizárólag, csak a Links webböngészőre vonatkozik.
`

Ha a HTTP proxy hitelesítést igényel, akkor a következő szintaxissal állítson be egy felhasználónevet és állítsa be a hozzá tartozó jelszót:

`root #` `export http_proxy="http://username:password@proxy.gentoo.org:8080" # A Portage szoftvercsomag-kezelőre és a Links webböngészőre vonatkozik.
`

`root #` `export https_proxy="http://username:password@proxy.gentoo.org:8080" # Kizárólag, csak a Links webböngészőre vonatkozik.
`

A proxy támogatásának megléte érdekében, a következő paraméterekkel indítsa el a links webböngészőt:

`user $` `links -http-proxy ${http_proxy} -https-proxy ${https_proxy} `

FTP proxy megadása a Portage szoftvercsomag-kezelő és/vagy a links webböngésző számára:

`root #` `export ftp_proxy="ftp://proxy.gentoo.org:8080" # A Portage szoftvercsomag-kezelőre és a Links webböngészőre vonatkozik.`

Egy FTP proxyhoz a következő paraméterrel indítsa el a links webböngészőt:

`user $` `links -ftp-proxy ${ftp_proxy} `

Egy RSYNC proxy megadása a Portage szoftvercsomag-kezelő számára:

`root #` `export RSYNC_PROXY="proxy.gentoo.org:8080" # Kizárólag, csak a Portage szoftvercsomag-kezelőre vonatkozik. A Links webböngésző nem támogatja az rsync proxy lehetőséget.`

### A pppoe-setup használata az ADSL internetkapcsolathoz

Akkor sincs probléma ha az interneteléréshez szükséges a PPPoE, mivel a Gentoo _bootolható adathordozója_ tartalmazza a pppoe-setup futtatható szkriptet a ppp beállításának a leegyszerűsítése érdekében.

A internetbeállítás során a pppoe-setup szkript a következőket fogja kérni:

- ADSL modemhez csatlakoztatott Ethernet _interfész_ neve.
- PPPoE felhasználónév és jelszó.
- DNS szerver IP-címe(i).
- Kérdezni fogja, hogy szükség van-e tűzfalra vagy sem.

`root #` `pppoe-setup
`

`root #` `pppoe-start`

Hiba esetén ellenőrizni kell a /etc/ppp/pap-secrets fájlban vagy /etc/ppp/chap-secrets fájlban lévő hitelesítő adatokat. Ha a hitelesítési adatok helyesek, akkor ellenőrizni kell a PPPoE Ethernet interfésznek a megfelelő kiválasztását.

### PPTP protokoll használata

Ha PPTP protokoll támogatásra van szükség, akkor a pptpclient parancs használható erre a célra, de használat előtt azt még be kell állítani.

Szerkessze az /etc/ppp/pap-secrets vagy a /etc/ppp/chap-secrets fájlt úgy, hogy az tartalmazza a megfelelő felhasználónév/jelszó kombinációt:

`root #` `nano /etc/ppp/chap-secrets`

Majd ha szükséges, akkor állítsa be az /etc/ppp/options.pptp fájlt:

`root #` `nano /etc/ppp/options.pptp`

A beállítás elvégzése után futtassa a pptp parancsot (a beállításokkal együtt, amelyeket az options.pptp fájlban nem lehetett beállítani) a szerver csatlakoztatásához:

`root #` `pptp <szerver ipv4 címe>`

## Internet és IP alapismeretek

Ha a fentiek mindegyike sikertelen, akkor a hálózatot manuálisan kell beállítani. Ez nem nehéz különösebben, de megfontoltan kell azt elvégezni. Ez a rész a terminológia tisztázására és az internetkapcsolat kézi úton történő beállítására vonatkozó alapvető hálózati fogalmak megismertetésére szolgál.

**Tip**

Egyes **CPE** ( **Carrier Provided Equipment**) egy egységben egyesíti az _útválasztó_, a _hozzáférési pont_, a _modem_, a _DHCP-kiszolgáló_ és a _DNS-kiszolgáló_ funkcióit. Fontos megkülönböztetni az eszköz funkcióit magától a fizikai eszköztől.

### Interfészek és címek

A hálózati _interfészek_ a hálózati eszközök logikai reprezentációi. Az _interfésznek_ _címre_ van szüksége ahhoz, hogy a _hálózaton_ lévő többi eszközzel kommunikálhasson. Bár csak egyetlen _cím_ szükséges, egyetlen _interfészhez_ több cím is hozzárendelhető. Ez különösen hasznos a kettős stack (IPv4 + IPv6) beállítások esetén.

A következetesség érdekében ez a példa feltételezi, hogy az enp1s0 interfész a 192.168.0.2 címet fogja használni.

**Important**

Az IP-címeket tetszőleges módon lehet beállítani. Ennek eredményeként előfordulhat, hogy több eszköz ugyanazt az IP-címet használja, ami _címütközést_ eredményez. A DHCP vagy a SLAAC használatával elkerülhetőek a címütközések.

**Tip**

Az IPv6 általában a **S** tate **L** ess **A** ddress **A** uto **C** onfiguration (SLAAC) funkciót használja a címek beállításához. A legtöbb esetben az IPv6-címek kézi úton történő beállítása rossz gyakorlat. Ha egy adott címutótagot részesítünk előnyben, akkor [interfészazonosító tokenek](/wiki/IPv6_Static_Addresses_using_Tokens "IPv6 Static Addresses using Tokens") használhatók.

### Hálózatok és a CIDR

Ha az eszköz kiválasztott egyszer egy címet, akkor honnan tudja, hogy miként kommunikáljon a másik eszközökkel?

Az IP- _címek_ hálózatokhoz vannak társítva. Az IP- _hálózatok_ összefüggő logikai címtartományok.

Az _osztály nélküli tartományok közötti útválasztás_ ( _Classless Inter-Domain Routing_) vagyis a _CIDR_ a hálózatok méretének megkülönböztetésére szolgál.

- A _CIDR_ érték, amelyet gyakran **/** karakterrel kezdődően írnak le, a hálózat méretét jelzi.

  - A _2 ^ (32 - CIDR)_ képlet használható a hálózat méretének kiszámításához.
  - A hálózat méretének kiszámítása után a használható csomópontok számát 2-vel kell csökkenteni.
    - A hálózat első IP-címe a _Hálózati cím_ ( _Network address_), az utolsó pedig általában a _Broadcast cím_ ( _Broadcast address_). Ezek a címek speciálisak, és normál gazdagépek nem használhatják őket.

**Tip**

A leggyakoribb _CIDR_-értékek a **/24** és **/32**, amelyek **254** csomópontokat, illetve egyetlen csomópontot jelentenek.

A **/24** értékű _CIDR_ tulajdonképpen az alapértelmezett hálózatméret. Ez pontosan a _255.255.255.0_ alhálózati maszknak felel meg, ahol az utolsó 8 bit a hálózat csomópontjainak IP-címei számára van fenntartva.

A 192.168.0.2/24 jelölésmód így értelmezhető:

- A _cím_ ( _address_) 192.168.0.2 .
- Amely a 192.168.0.0 _hálózat_ ( _network_) részét képezi.
- A **254** ( _2 ^ (32 - 24) \- 2_) mérettel rendelkezik.

  - A használható IP-címek a 192.168.0.1 - 192.168.0.254 tartományban vannak benne.
- A 192.168.0.255 _szórási címmel_ ( _broadcast address_) rendelkezik.

  - A legtöbb esetben a gyakorlatban a hálózat utolsó címét használjuk _szórási címként_ ( _broadcast address_), de ez megváltoztatható.

Ezzel a beállítással az eszköznek képesnek kell lennie bármely gazdagéppel kommunikálni ugyanazon a _hálózaton_ (192.168.0.0).

### Internet

Ha egy eszköz a hálózaton van, akkor honnan tudja, hogy miként kommunikáljon az interneten lévő többi eszközzel?

A helyi _hálózatokon_ ( _networks_) kívüli eszközökkel való kommunikációhoz _útválasztást_ ( _routing_) kell használni. Az _útválasztó_ ( _router_) egyszerűen csak egy hálózati eszköz, amely forgalmat továbbít más eszközök számára. Az _alapértelmezett útvonal_ ( _default route_) vagy _átjáró_ ( _gateway_) kifejezés általában az aktuális hálózat bármely eszközére utal, amelyet külső hálózati hozzáféréshez használnak.

**Tip**

Bevett gyakorlati szokás, hogy az _átjárót_ ( _gateway_) a hálózat első vagy utolsó IP-címére teszik rá.

Ha elérhető egy internetre csatlakoztatott útválasztó a 192.168.0.1 címen, akkor az használható az _alapértelmezett útvonalként_ ( _default route_), ami internet-hozzáférést biztosít.

Összegezve:

- Az interfészeket _címmel_ ( _address_) és _hálózati információkkal_ ( _network information_), például a _CIDR_ értékkel kell beállítani.
- A helyi hálózati hozzáférés az ugyanazon a hálózaton lévő _útválasztó_ ( _router_) elérésére szolgál.
- Az _alapértelmezett útvonal_ ( _default route_) be van állítva, így a külső hálózatokra irányuló forgalom az _átjáróra_ ( _gateway_) kerül továbbításra, amely internet-hozzáférést biztosít.

### Domain Name System

Az embereknek fejből nehéz megjegyezni az IP-címeket a gépekhez képest. A _tartománynévrendszert_ ( _Domain Name System_) azért hozták létre, hogy lehetővé tegye a _tartománynevek_ ( _Domain Names_) és az IP-címek _(_ IP addresses _) közötti leképezést._

A Linux rendszerek az /etc/resolv.conf fájlt használják a DNS-névfeloldáshoz ( _DNS resolution_) használandó névszerverek ( _nameservers_) meghatározásához.

**Tip**

Sok _útválasztó_ ( _router_) DNS-szerverként is funkcionálhat, és egy helyi DNS-szerver használata növelheti a magánéletünk védelmét, valamint a gyorsítótárazás használatával felgyorsíthatja az internetes lekérdezéseinket.

Sok internetszolgáltató (ISP) olyan DNS szervert működtet az ő hálózatában, amelyet általában DHCP szolgáltatás segítségével hirdet meg a mi átjárónknak ( _gateway_). A helyi DNS szerver használata általában javítja a lekérdezési késleltetést, de a legtöbb nyilvános DNS szerver ugyanazokat az eredményeket adja vissza, így a szerverhasználat nagyrészt a hozzáértő felhasználók egyéni preferenciáin alapulnak.

## Hálózat kézi úton történő beállítása

### Hálózati interfész címének a beállítása

**Important**

Az IP-címek kézi úton történő beállításakor figyelembe kell venni a helyi hálózati topológiát. Az IP-címek tetszőlegesen beállíthatók, de tudni kell, hogy a konfliktusok, címütközések mind hálózati zavarokat okozhatnak.

Az enp1s0 beállítása a 192.168.0.2 címmel és a CIDR /24 alhálózattal:

`root #` `ip address add 192.168.0.2/24 dev enp1s0`

**Tip**

Ennek a parancsnak az eleje lerövidíthető az ip a parancsra.

### Alapértelmezett útvonal beállítása

Egy interfész címének és hálózati információinak megadása beállítja a link útvonalakat, lehetővé téve a kommunikációt az adott hálózati szegmenssel (network segment):

`root #` `ip route`

```
192.168.0.0/24 dev enp1s0 proto kernel scope link src 192.168.0.2
```

**Tip**

Ezt a parancsot le lehet rövidíteni a ip r parancsra.

A alapértelmezett (default) útvonal a következőképpen állítható be (192.168.0.1):

`root #` `ip route add default via 192.168.0.1`

### DNS beállítása

A névszerver adatait általában DHCP szolgáltatással szerezzük be, de sajátkezűleg is beállíthatóak az /etc/resolv.conf fájl `nameserver` bejegyzés hozzáadásával.

**Warning**

Ha a _dhcpcd_ fut, akkor az /etc/resolv.conf módosításai nem maradnak meg. Az állapot a `ps x | grep dhcpcd` segítségével ellenőrizhető le.

A nano parancssoros szövegszerkesztő benne van a Gentoo _bootolható adathordozóján_, használható az /etc/resolv.conf fájl szerkesztésére a következő parancs kiadásával:

`root #` `nano /etc/resolv.conf`

A `nameserver` kulcsszót és a DNS-kiszolgáló IP-címét tartalmazó sorok a definíció szerinti sorrendben kerülnek lekérdezésre:

FILE **`/etc/resolv.conf`** **Quad9 DNS használata**

```
nameserver 9.9.9.9
nameserver 149.112.112.112

```

FILE **`/etc/resolv.conf`** **Cloudflare DNS használata**

```
nameserver 1.1.1.1
nameserver 1.0.0.1

```

A DNS állapota ellenőrizhető egy domain név pingelésével:

`root #` `ping -c 3 gentoo.org`

Miután [ellenőrizte a kapcsolatot](/wiki/Handbook:Alpha/Installation/Networking/hu#Testing_the_network "Handbook:Alpha/Installation/Networking/hu"), folytassa az [adathordozók előkészítésével](/wiki/Handbook:Alpha/Installation/Disks/hu "Handbook:Alpha/Installation/Disks/hu").

[← Telepítőképfájl kiválasztása](/wiki/Handbook:Alpha/Installation/Media/hu "Previous part") [Kezdőlap](/wiki/Handbook:Alpha/hu "Handbook:Alpha/hu") [Adathordozók előkészítése →](/wiki/Handbook:Alpha/Installation/Disks/hu "Next part")