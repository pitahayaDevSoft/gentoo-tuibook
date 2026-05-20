# Media

Other languages:

- [Deutsch](/wiki/Handbook:MIPS/Installation/Media/de "Handbook:MIPS/Installation/Media/de (57% translated)")
- [English](/wiki/Handbook:MIPS/Installation/Media "Handbook:MIPS/Installation/Media (100% translated)")
- [español](/wiki/Handbook:MIPS/Installation/Media/es "Handbook:MIPS/Installation/Media/es (57% translated)")
- [français](/wiki/Handbook:MIPS/Installation/Media/fr "Handbook:MIPS/Installation/Media/fr (57% translated)")
- [italiano](/wiki/Handbook:MIPS/Installation/Media/it "Handbook:MIPS/Installation/Media/it (100% translated)")
- magyar
- [polski](/wiki/Handbook:MIPS/Installation/Media/pl "Handbook:MIPS/Installation/Media/pl (14% translated)")
- [português do Brasil](/wiki/Handbook:MIPS/Installation/Media/pt-br "Handbook:MIPS/Installation/Media/pt-br (57% translated)")
- [русский](/wiki/Handbook:MIPS/Installation/Media/ru "Handbook:MIPS/Installation/Media/ru (100% translated)")
- [தமிழ்](/wiki/Handbook:MIPS/Installation/Media/ta "Handbook:MIPS/Installation/Media/ta (57% translated)")
- [中文（中国大陆）‎](/wiki/Handbook:MIPS/Installation/Media/zh-cn "Handbook:MIPS/Installation/Media/zh-cn (57% translated)")
- [日本語](/wiki/Handbook:MIPS/Installation/Media/ja "Handbook:MIPS/Installation/Media/ja (100% translated)")
- [한국어](/wiki/Handbook:MIPS/Installation/Media/ko "Handbook:MIPS/Installation/Media/ko (57% translated)")

[MIPS kézikönyv](/wiki/Handbook:MIPS/hu "Handbook:MIPS/hu")[A Gentoo Linux telepítése](/wiki/Handbook:MIPS/Full/Installation/hu "Handbook:MIPS/Full/Installation/hu")[A telepítésről](/wiki/Handbook:MIPS/Installation/About/hu "Handbook:MIPS/Installation/About/hu")Telepítőképfájl kiválasztása[Hálózat beállítása](/wiki/Handbook:MIPS/Installation/Networking/hu "Handbook:MIPS/Installation/Networking/hu")[Adathordozók előkészítése](/wiki/Handbook:MIPS/Installation/Disks/hu "Handbook:MIPS/Installation/Disks/hu")[Stage fájl](/wiki/Handbook:MIPS/Installation/Stage/hu "Handbook:MIPS/Installation/Stage/hu")[Alaprendszer telepítése](/wiki/Handbook:MIPS/Installation/Base/hu "Handbook:MIPS/Installation/Base/hu")[Kernel beállítása](/wiki/Handbook:MIPS/Installation/Kernel/hu "Handbook:MIPS/Installation/Kernel/hu")[Rendszer beállítása](/wiki/Handbook:MIPS/Installation/System/hu "Handbook:MIPS/Installation/System/hu")[Eszközök telepítése](/wiki/Handbook:MIPS/Installation/Tools/hu "Handbook:MIPS/Installation/Tools/hu")[Bootloader beállítása](/wiki/Handbook:MIPS/Installation/Bootloader/hu "Handbook:MIPS/Installation/Bootloader/hu")[Telepítés véglegesítése](/wiki/Handbook:MIPS/Installation/Finalizing/hu "Handbook:MIPS/Installation/Finalizing/hu")[Munka a Gentoo rendszerrel](/wiki/Handbook:MIPS/Full/Working/hu "Handbook:MIPS/Full/Working/hu")[Portage bemutatása](/wiki/Handbook:MIPS/Working/Portage/hu "Handbook:MIPS/Working/Portage/hu")[USE jelölőzászlók](/wiki/Handbook:MIPS/Working/USE/hu "Handbook:MIPS/Working/USE/hu")[Portage jellemzői](/wiki/Handbook:MIPS/Working/Features/hu "Handbook:MIPS/Working/Features/hu")[Init-szkript rendszer](/wiki/Handbook:MIPS/Working/Initscripts/hu "Handbook:MIPS/Working/Initscripts/hu")[Környezeti változók](/wiki/Handbook:MIPS/Working/EnvVar/hu "Handbook:MIPS/Working/EnvVar/hu")[Munka a Portage szoftvercsomag-kezelővel](/wiki/Handbook:MIPS/Full/Portage/hu "Handbook:MIPS/Full/Portage/hu")[Fájlok és könyvtárak](/wiki/Handbook:MIPS/Portage/Files/hu "Handbook:MIPS/Portage/Files/hu")[Változók](/wiki/Handbook:MIPS/Portage/Variables/hu "Handbook:MIPS/Portage/Variables/hu")[Szoftverágak keverése](/wiki/Handbook:MIPS/Portage/Branches/hu "Handbook:MIPS/Portage/Branches/hu")[További eszközök](/wiki/Handbook:MIPS/Portage/Tools/hu "Handbook:MIPS/Portage/Tools/hu")[Egyéni szoftvercsomag-tárolók](/wiki/Handbook:MIPS/Portage/CustomTree/hu "Handbook:MIPS/Portage/CustomTree/hu")[Fejlett funkciók](/wiki/Handbook:MIPS/Portage/Advanced/hu "Handbook:MIPS/Portage/Advanced/hu")[Hálózat beállítása OpenRC init-rendszeren](/wiki/Handbook:MIPS/Full/Networking/hu "Handbook:MIPS/Full/Networking/hu")[Munka elkezdése](/wiki/Handbook:MIPS/Networking/Introduction/hu "Handbook:MIPS/Networking/Introduction/hu")[Fejlett beállítások](/wiki/Handbook:MIPS/Networking/Advanced/hu "Handbook:MIPS/Networking/Advanced/hu")[Moduláris hálózat](/wiki/Handbook:MIPS/Networking/Modular/hu "Handbook:MIPS/Networking/Modular/hu")[Vezeték nélküli (Wi-Fi)](/wiki/Handbook:MIPS/Networking/Wireless/hu "Handbook:MIPS/Networking/Wireless/hu")[Funkcionalitás hozzáadása](/wiki/Handbook:MIPS/Networking/Extending/hu "Handbook:MIPS/Networking/Extending/hu")[Dinamikus menedzsment](/wiki/Handbook:MIPS/Networking/Dynamic/hu "Handbook:MIPS/Networking/Dynamic/hu")

## Contents

- [1Hardverkövetelmények](#Hardverk.C3.B6vetelm.C3.A9nyek)
- [2Telepítési megjegyzések](#Telep.C3.ADt.C3.A9si_megjegyz.C3.A9sek)
- [3Netbooting áttekintés](#Netbooting_.C3.A1ttekint.C3.A9s)
  - [3.1TFTP és DHCP beállítása](#TFTP_.C3.A9s_DHCP_be.C3.A1ll.C3.ADt.C3.A1sa)
- [4Hálózati indítás SGI munkaállomásokon](#H.C3.A1l.C3.B3zati_ind.C3.ADt.C3.A1s_SGI_munka.C3.A1llom.C3.A1sokon)
  - [4.1Netboot képfájl letöltése](#Netboot_k.C3.A9pf.C3.A1jl_let.C3.B6lt.C3.A9se)
  - [4.2DHCP beállítása az SGI kliens számára](#DHCP_be.C3.A1ll.C3.ADt.C3.A1sa_az_SGI_kliens_sz.C3.A1m.C3.A1ra)
  - [4.3Kernelopciók](#Kernelopci.C3.B3k)
  - [4.4Szolgáltatások elindítása](#Szolg.C3.A1ltat.C3.A1sok_elind.C3.ADt.C3.A1sa)
  - [4.5SGI munkaállomás hálózatról történő indítása (Netbooting)](#SGI_munka.C3.A1llom.C3.A1s_h.C3.A1l.C3.B3zatr.C3.B3l_t.C3.B6rt.C3.A9n.C5.91_ind.C3.ADt.C3.A1sa_.28Netbooting.29)
  - [4.6Hibaelhárítás](#Hibaelh.C3.A1r.C3.ADt.C3.A1s)
- [5Cobalt munkaállomások hálózati úton történő indítása (Netbooting)](#Cobalt_munka.C3.A1llom.C3.A1sok_h.C3.A1l.C3.B3zati_.C3.BAton_t.C3.B6rt.C3.A9n.C5.91_ind.C3.ADt.C3.A1sa_.28Netbooting.29)
  - [5.1Hálózati indítási eljárás áttekintése](#H.C3.A1l.C3.B3zati_ind.C3.ADt.C3.A1si_elj.C3.A1r.C3.A1s_.C3.A1ttekint.C3.A9se)
  - [5.2Cobalt hálózati indító képfájl (netboot image) letöltése](#Cobalt_h.C3.A1l.C3.B3zati_ind.C3.ADt.C3.B3_k.C3.A9pf.C3.A1jl_.28netboot_image.29_let.C3.B6lt.C3.A9se)
  - [5.3NFS szerver beállítása](#NFS_szerver_be.C3.A1ll.C3.ADt.C3.A1sa)
  - [5.4Cobalt számítógép DHCP beállítása](#Cobalt_sz.C3.A1m.C3.ADt.C3.B3g.C3.A9p_DHCP_be.C3.A1ll.C3.ADt.C3.A1sa)
  - [5.5Szolgáltatások elindítása](#Szolg.C3.A1ltat.C3.A1sok_elind.C3.ADt.C3.A1sa_2)
  - [5.6Cobalt számítógép hálózati úton történő indítása](#Cobalt_sz.C3.A1m.C3.ADt.C3.B3g.C3.A9p_h.C3.A1l.C3.B3zati_.C3.BAton_t.C3.B6rt.C3.A9n.C5.91_ind.C3.ADt.C3.A1sa)
  - [5.7Hibaelhárítás](#Hibaelh.C3.A1r.C3.ADt.C3.A1s_2)
- [6Telepítő CD használata](#Telep.C3.ADt.C5.91_CD_haszn.C3.A1lata)

## Hardverkövetelmények

CPU (Big Endian port)
MIPS3, MIPS4, MIPS5 or MIPS64-class CPU
CPU (Little Endian port)
MIPS4, MIPS5 or MIPS64-class CPU
Memória
128 MB
Adathordozó területe
3.0 GB (Swap területe nélkül)
Swap területe
Legalább 256 MB

További információért olvassa el a [MIPS hardverkövetelmények](/wiki/MIPS/Hardware_Requirements "MIPS/Hardware Requirements") dokumentumot.

## Telepítési megjegyzések

Sok architektúránál a processzor több generáción ment keresztül, mindegyik újabb generáció az előző alapjaira épült. A MIPS sem kivétel. A MIPS architektúra alá tartozóan több CPU-generáció létezik. Annak érdekében, hogy megfelelő netboot bináris képfájl tömörített állományt és CFLAGS értékeket válasszunk, szükséges tisztában lenni azzal, hogy a rendszer processzora melyik családhoz tartozik. Ezeket a családokat utasításkészlet architektúrának nevezik.

MIPS ISA
32/64-bit
Processzorokat fedi le
MIPS 1
32-bit
R2000, R3000
MIPS 2
32-bit
R6000
MIPS 3
64-bit
R4000, R4400, R4600, R4700
MIPS 4
64-bit
R5000, RM5000, RM7000 R8000, R9000, R10000, R12000, R14000, R16000
MIPS 5
4-bit
Egyik sem
MIPS32
32-bit
AMD Alchemy sorozat, 4kc, 4km, sok más... A MIPS32 ISA-ban van néhány revízió.
MIPS64
64-bit
Broadcom SiByte SB1, 5kc ... stb... A MIPS64 ISA-ban van néhány revízió.

**Note**

A MIPS5 ISA szintet a Silicon Graphics tervezte még 1994-ben, de soha nem használták ténylegesen egy valós CPU-ban. Része maradt a MIPS64 ISA-nak.

**Note**

A MIPS32 és MIPS64 ISA gyakran okoz zavart. A MIPS64 ISA szint valójában a MIPS5 ISA bővített változata, így tartalmazza az összes utasítást a MIPS5 és korábbi ISA szintekből. A MIPS32 a MIPS64 32-bites részkészlete, amely azért létezik, mert a legtöbb alkalmazásnak csak 32-bites feldolgozásra van szüksége.

Egy másik fontos fogalom, amelyet érdemes megérteni, az endianness fogalma. Az endianness arra utal, hogy a CPU hogyan olvassa be a szavakat a főmemóriából. Egy szó olvasható big endian formátumban (legjelentősebb bájt elöl) vagy little endian formátumban (legkisebb jelentőségű bájt elöl). Az Intel x86 számítógépek általában little endian módúak, míg az Apple és Sparc gépek big endian módúak. A MIPS esetében mindkettő lehetséges. Ezek megkülönböztetésére _el_ toldalékot adunk az architektúra nevéhez, hogy jelöljük a little endian formátumot.

Architektúra
32/64-bit
Endianness
Számítógépeket magába folgalja
mips
32-bit
Big Endian
Silicon Graphics
mipsel
32-bit
Little Endian
Cobalt Servers
mips64
64-bit
Big Endian
Silicon Graphics
mips64el
64-bit
Little Endian
Cobalt Servers

Azok számára, akik többet szeretnének megtudni az ISA-król, az alábbi weboldalak hasznosak lehetnek:

- [Linux/MIPS Website: MIPS ISA](http://www.linux-mips.org/wiki/Instruction_Set_Architecture)
- [Linux/MIPS Website: Endianness](http://www.linux-mips.org/wiki/Endianess)
- [Linux/MIPS Website: Processors](http://www.linux-mips.org/wiki/Processors)
- [Wikipedia: Instruction Set](https://en.wikipedia.org/wiki/Instruction_set)

## Netbooting áttekintés

Ebben a szakaszban áttekintjük, hogy mire van szükség egy Silicon Graphics munkaállomás vagy Cobalt Server eszköz sikeres hálózati indításához. Ez csak egy rövid útmutató, így nem célja, hogy átfogó legyen. További információért ajánlott elolvasni a [Lemez nélküli csomópontok](/wiki/Diskless_nodes/hu "Diskless nodes/hu") című cikket.

A számítógéptől függően bizonyos mennyiségű hardverre van szükség ahhoz, hogy a hálózati indítás és a Linux telepítése sikeres legyen.

- Általában:
  - DHCP/BOAMD Alchemy sorozat, 4kc, 4km, sok más... A MIPS32 ISA-ban van néhány revízió. OTP szerver (ajánlott az ISC DHCPd).
  - Türelem -- és sok belőle.
- Silicon Graphics munkaállomások esetében:
  - TFTP szerver (tftp-hpa az ajánlott).
  - Amikor a soros konzolt használni kell:
    - MiniDIN8 --> RS-232 soros kábel (csak IP22 és IP28 rendszerekhez szükséges).
    - Null-modem kábel.
    - VT100 vagy ANSI kompatibilis terminál, amely képes a 9600 baud sebességre.
- Cobalt szerverek esetében (NEM az eredeti Qube):
  - NFS szerver.
  - Null-modem kábel.
  - VT100 vagy ANSI kompatibilis terminál, amely képes az 115200 baud sebességre.

**Note**

Az SGI számítógépek MiniDIN 8 csatlakozót használnak a soros portokhoz. Nyilvánvalóan az Apple modemkábelek soros kábelként is jól működnek, de mivel az Apple számítógépeket USB porttal és belső modemekkel szerelik fel, ezek egyre nehezebben találhatók meg. Egy bekötési diagram elérhető a [Linux/MIPS Wiki](http://www.linux-mips.org/wiki/Serial_Cable) oldalon, és a legtöbb elektronikai üzletben kaphatók a szükséges csatlakozók.

**Note**

A terminál lehet egy valódi VT100/ANSI terminál, vagy lehet egy PC, amely terminálemulációs szoftvert futtat (például HyperTerminal, Minicom, seyon, Telex, xc, screen - amelyik az Ön preferenciája). Nem számít, hogy ez a számítógép milyen platformot futtat - csak az a lényeg, hogy rendelkezzen egy RS-232 soros porttal és megfelelő szoftverrel.

**Note**

Ez az útmutató NEM foglalkozik az eredeti Qube-val. Az eredeti Qube szervereszköz alapértelmezett beállításában nem rendelkezik soros porttal, ezért nem lehet Gentoo operációs rendszert telepíteni rá anélkül, hogy ne lenne szükség egy csavarhúzóra és egy helyettesítő számítógépre a telepítéshez.

### TFTP és DHCP beállítása

Ahogy korábban említettük, ez nem egy teljes értékű útmutató. Ez egy alapvető beállítás, amely csak az indulást biztosítja. Használja ezt akkor, amikor nulláról kezdi a beállítást, vagy használja az ajánlásokat egy meglévő beállítás módosításához, hogy támogassa a hálózati indítást.

Érdemes megjegyezni, hogy a használt szervereknek nem kell feltétlenül Gentoo Linuxot futtatniuk, akár FreeBSD-t vagy bármely Unix-szerű platformot is használhatnak. Ez az útmutató azonban Gentoo Linux használatát feltételezi. Ha szükséges, akkor az TFTP/NFS futtatása lehetséges egy külön számítógépen is, a DHCP szervertől függetlenül.

**Warning**

A Gentoo/MIPS csapat nem tud segítséget nyújtani más operációs rendszerek beállításához hálózati indítószerverként.

Első lépés a DHCP beállítása. Annak érdekében, hogy az ISC DHCP szolgáltatás reagáljon a BOOTP kérésekre (ahogy azt az SGI és Cobalt BOOTROM megköveteli), először engedélyezze a dinamikus BOOTP-t a használt címkörzeten, majd hozzon létre egy bejegyzést minden kliens számára, amely hivatkozásokat tartalmaz az boot képfájlra.

`root #` `emerge --ask net-misc/dhcp`

Miután telepítette, hozza létre a /etc/dhcp/dhcpd.conf fájlt. Íme egy egyszerű beállítás, amely segít az indulásban:

FILE **`/etc/dhcp/dhcpd.conf`** **Csupasz dhcpd.conf fájl**

```
# Tell dhcpd to disable dynamic DNS.
# dhcpd will refuse to start without this.
ddns-update-style none;

# Create a subnet:
subnet 192.168.10.0 netmask 255.255.255.0 {
  # Address pool for our booting clients. Don't forget the 'dynamic-bootp' bit!
  pool {
    range dynamic-bootp 192.168.10.1 192.168.10.254;
  }

  # DNS servers and default gateway -- substitute as appropriate
  option domain-name-servers 203.1.72.96, 202.47.56.17;
  option routers 192.168.10.1;

  # Tell the DHCP server it's authoritative for this subnet.
  authoritative;

  # Allow BOOTP to be used on this subnet.
  allow bootp;
}

```

A megadott beállítással tetszőleges számú klienst adhat hozzá az alhálózati záradékban. A későbbiekben ebben az útmutatóban tárgyaljuk, hogy mit kell hozzáadni.

Következő lépés a TFTP szerver beállítása. Ajánlott a tftp-hpa használata, mivel ez az egyetlen ismert TFTP szolgáltatás, amely helyesen működik. A telepítést az alábbiak szerint végezze el:

`root #` `emerge --ask net-ftp/tftp-hpa`

Ez létrehozza a /tftproot könyvtárat, amelybe a hálózati indító bináris képfájlt tárolhatja. Szükség esetén helyezze át máshova. Ennek az útmutatónak a céljaira feltételezzük, hogy az alapértelmezett helyen marad.

## Hálózati indítás SGI munkaállomásokon

### Netboot képfájl letöltése

A telepítés céljának megfelelően többféle letölthető képfájl érhető el. Ezek mind a rendszer típusának és a processzorának megfelelően vannak megjelölve, amelyre készültek. A számítógéptípusok az alábbiak:

Kódnév
Számítógépek
IP22
Indy, \*Indigo 2, Challenge S
IP26
\*Indigo 2 Power
IP27
Origin 200, Origin 2000
IP28
\*Indigo 2 Impact
IP30
Octane
IP32
O2

**Note**

_Indigo 2_ \- Gyakori hiba összekeverni az IRIS Indigo számítógépet (IP12 R3000 CPU-val vagy IP20 R4000 CPU-val, amelyek egyike sem futtat Linuxot), az Indigo 2 számítógépet (IP22, amely jól futtatja a Linuxot), az R8000-alapú Indigo 2 Power számítógépet (amely egyáltalán nem futtat Linuxot) és az R10000-alapú Indigo 2 Impact számítógépet (IP28, amely nagyon kísérleti). Kérjük, hogy tartsa szem előtt, hogy ezek mind különböző számítógépek.

Továbbá a fájlnévben az r4k az R4000-sorozatú processzorokra utal, az r5k az R5000-re, az rm5k az RM5200-ra, és az r10k az R10000-re. A képfájlok elérhetők a [Gentoo tükörszervereken](https://www.gentoo.org/downloads/mirrors/).

### DHCP beállítása az SGI kliens számára

A fájl letöltése után helyezze az kicsomagolt képfájlt a /tftproot/ könyvtárba. (A kicsomagoláshoz használja a bzip2 -d parancsot). Ezután szerkessze a /etc/dhcp/dhcpd.conf fájlt, és adja hozzá az SGI klienshez megfelelő bejegyzést.

FILE **`/etc/dhcp/dhcpd.conf`** **Részlet az SGI munkaállomás számára**

```
subnet xxx.xxx.xxx.xxx netmask xxx.xxx.xxx.xxx {
  # ... usual stuff here ...

  # SGI Workstation... change 'sgi' to your SGI machine's hostname.
  host sgi {

    # MAC Address of SGI Machine. Normally this is written on the back
    # or base of the machine.
    hardware ethernet 08:00:69:08:db:77;

    # TFTP Server to download from (by default, same as DHCP server)
    next-server 192.168.10.1;

    # IP address to give to the SGI machine
    fixed-address 192.168.10.3;

    # Filename for the PROM to download and boot
    filename "/gentoo-r4k.img";
  }
}

```

### Kernelopciók

Majdnem készen vagyunk, de van néhány apró módosítás, amit még el kell végezni. Nyisson meg egy konzolt root jogosultságokkal.

Tiltsa le a "Path Maximum Transfer Unit"-t, különben az SGI PROM nem találja meg a kernelt:

`root #` `echo 1 > /proc/sys/net/ipv4/ip_no_pmtu_disc`

Állítsa be az SGI PROM által használható porttartományt:

`root #` `echo "2048 32767" > /proc/sys/net/ipv4/ip_local_port_range`

Ez elegendőnek kell lennie ahhoz, hogy a Linux szerver megfelelően működjön együtt az SGI PROM-mal.

### Szolgáltatások elindítása

Ezen a ponton indítsa el a szolgáltatásokat.

`root #` `/etc/init.d/dhcp start
`

`root #` `/etc/init.d/in.tftpd start
`

Ha az utolsó lépésnél semmi probléma nem történt, akkor minden készen áll arra, hogy bekapcsolja a munkaállomást és folytassa az útmutatót. Ha a DHCP szerver valamilyen okból nem indul el, akkor próbálja meg futtatni a dhcpd parancsot a parancssorban, és nézze meg, mit mond - ha minden rendben, akkor egyszerűen háttérbe kerül, ellenkező esetben az alatta megjelenő panasz mellett az 'exiting.' üzenetet fogja megjeleníteni.

Egy egyszerű módja annak, hogy ellenőrizze, fut-e a tftp szolgáltatás, az, ha beírja a következő parancsot és megerősíti a kimenetet:

`root #` `netstat -al | grep ^udp`

```
udp        0      0 *:bootpc                *:*
udp        0      0 *:631                   *:*
udp        0      0 *:xdmcp                 *:*
udp        0      0 *:tftp                  *:* <-- (look for this line)

```

### SGI munkaállomás hálózatról történő indítása (Netbooting)

Rendben, minden készen áll, a DHCP és a TFTP fut. Most itt az ideje bekapcsolni az SGI számítógépet. Kapcsolja be az egységet. Amikor a "Running power-on diagnostics" megjelenik a képernyőn, akkor kattintson a "Stop For Maintenance" gombra, vagy nyomja meg a `Escape` billentyűgombot. Egy a következőhöz hasonló menü fog megjelenni.

`Running power-on diagnostics`

```
System Maintenance Menu

1) Start System
2) Install System Software
3) Run Diagnostics
4) Recover System
5) Enter Command Monitor
Option?

```

Nyomja meg a `5` billentyűgombot, hogy belépjen a parancs monitorba. A monitoron indítsa el a BootP folyamatot:

`>>` `bootp(): root=/dev/ram0`

Ettől a ponttól kezdve a számítógépnek el kell kezdenie letölteni a képfájlt, majd nagyjából 20 másodperccel később el kell kezdenie a Linux indítását. Ha minden rendben van, akkor egy busybox ash shell fog elindulni az alábbiak szerint, és a Gentoo Linux telepítése folytatódhat.

CODE **Amikor a dolgok jól mennek...**

```
init started:  BusyBox v1.00-pre10 (2004.04.27-02:55+0000) multi-call binary

Gentoo Linux; http://www.gentoo.org/
 Copyright 2001-2004 Gentoo Technologies, Inc.; Distributed under the GPL

 Gentoo/MIPS Netboot for Silicon Graphics Machines
 Build Date: April 26th, 2004

 * To configure networking, do the following:

 * For Static IP:
 * /bin/net-setup <IP Address> <Gateway Address> [telnet]

 * For Dynamic IP:
 * /bin/net-setup dhcp [telnet]

 * If you would like a telnetd daemon loaded as well, pass "telnet"
 * As the final argument to /bin/net-setup.

Please press Enter to activate this console.

```

### Hibaelhárítás

Ha a számítógép makacs és nem hajlandó letölteni a képfájlt, annak két oka lehet:

1. Az utasítások nem lettek helyesen követve, vagy
2. Egy kis gyengéd ösztönzésre van szükség. (Nem! Tegye le azt a kalapácsot!)

Itt van néhány ellenőrizendő dolog:

- A dhcpd biztosítja az SGI számítógépnek az IP-címet. A rendszer naplóiban kell lennie néhány üzenetnek egy BOOTP kérésről. A tcpdump itt is hasznos lehet.
- A tftp könyvtár (általában /tftproot/) jogosultságai megfelelően vannak beállítva (mindenki által olvashatónak kell lennie).
- Ellenőrizze a rendszer naplóit, hogy lássa, mit jelez a tftp szerver (esetleges hibák).

Ha ellenőrizett mindent a szerveren, és az SGI számítógépen időtúllépés vagy más hibák jelentkeznek, akkor próbálja meg beírni ezt a konzolba:

`>>` `resetenv
`

`>>` `unsetenv netaddr
`

`>>` `unsetenv dlserver
`

`>>` `init
`

`>>` `bootp(): root=/dev/ram0`

## Cobalt munkaállomások hálózati úton történő indítása (Netbooting)

### Hálózati indítási eljárás áttekintése

Ellentétben az SGI számítógépekkel, a Cobalt szerverek NFS-t használnak a kernel átviteléhez az indításhoz. Indítsa el a számítógépet úgy, hogy lenyomva tartja a bal és jobb nyíl gombokat, miközben bekapcsolja az egységet. A számítógép megpróbál egy IP-címet kapni BOOTP-n keresztül, felcsatolja az /nfsroot/ könyvtárat a szerverről NFS-en keresztül, majd megpróbálja letölteni és indítani a vmlinux\_raq-2800.gz fájlt (a modelltől függően), amelyet egy szabványos ELF bináris fájlnak tekint.

### Cobalt hálózati indító képfájl (netboot image) letöltése

A Cobalt számítógép beindításához szükséges indító képfájlok elérhetők a következő helyen: [http://distfiles.gentoo.org/experimental/mips/historical/netboot/cobalt/](http://distfiles.gentoo.org/experimental/mips/historical/netboot/cobalt/). A fájlok neve nfsroot-KERNEL-COLO-DATE-cobalt.tar formátumú lesz. Válassza ki a legfrissebbet, majd csomagolja ki a / könyvtárba az alábbiak szerint:

`root #` `tar -C / -xvf nfsroot-2.6.13.4-1.19-20051122-cobalt.tar`

### NFS szerver beállítása

Mivel ez a számítógép NFS protokollt használ a képfájl letöltéséhez, szükséges a /nfsroot/ könyvtár megosztása a szerveren. Telepítse a [net-fs/nfs-utils](https://packages.gentoo.org/packages/net-fs/nfs-utils) szoftvercsomagot:

`root #` `emerge --ask net-fs/nfs-utils`

Miután ez megtörtént, helyezze az alábbi bejegyzéseket a /etc/exports fájlba:

FILE **`/etc/exports`** **Exporting the /nfsroot directory**

```
/nfsroot      *(ro,sync)

```

Most, hogy ez megtörtént, indítsa el az NFS szervert:

`root #` `/etc/init.d/nfs start`

Ha az NFS szerver már futott, amikor a változtatásokat végrehajtotta, kérje meg, hogy nézze át újra az exports fájlt az exportfs parancs használatával.

`root #` `exportfs -av`

### Cobalt számítógép DHCP beállítása

Most biztosan egyszerűvé válik a DHCP beállítása. Adja hozzá az alábbi beállításokat a /etc/dhcp/dhcpd.conf fájlhoz:

FILE **`/etc/dhcp/dhcpd.conf`** **Részlet a Cobalt szerver számára**

```
subnet xxx.xxx.xxx.xxx netmask xxx.xxx.xxx.xxx {
  # ... usual stuff here ...

  # Configuration for a Cobalt Server
  # Set the hostname here:
  host qube {
    # Path to the nfsroot directory.
    # This is mainly for when using the TFTP boot option on CoLo
    # You shouldn't need to change this.
    option root-path "/nfsroot";

    # Cobalt server's ethernet MAC address
    hardware ethernet 00:10:e0:00:86:3d;

    # Server to download image from
    next-server 192.168.10.1;

    # IP address of Cobalt server
    fixed-address 192.168.10.2;

    # Location of the default.colo file relative to /nfsroot
    # You shouldn't need to change this.
    filename "default.colo";
  }
}

```

### Szolgáltatások elindítása

Most indítsa el a szolgáltatásokat. Írja be a következőket:

`root #` `/etc/init.d/dhcp start
`

`root #` `/etc/init.d/nfs start`

Ha az utolsó lépés során nem történt semmi probléma, akkor minden készen áll arra, hogy bekapcsolja a munkaállomást és folytassa az útmutatót. Ha a DHCP szerver valamilyen okból nem indul el, akkor próbálja meg futtatni a dhcpd parancsot a parancssorban, és nézze meg, mit mond - ha minden rendben van, akkor egyszerűen háttérbe kerül, ellenkező esetben az alatta megjelenő panasz mellett az 'exiting.' üzenetet fogja megjeleníteni.

### Cobalt számítógép hálózati úton történő indítása

Most itt az ideje elindítani a Cobalt számítógépet. Csatlakoztassa a null modem kábelt, és állítsa be a soros terminált 115200 baud sebességre, 8 bitre, paritás nélkül, 1 stop bitre, VT100 emulációval. Miután ez megtörtént, tartsa lenyomva a bal és jobb nyíl gombokat, miközben bekapcsolja az egységet.

A hátsó panelnek meg kell jelenítenie a "Net Booting" feliratot, és némi hálózati aktivitásnak láthatónak kell lennie, amit hamarosan követ CoLo működésbe lépése. A hátsó panelen görgessen le a menüben az "Network (NFS)" opcióig, majd nyomja meg a `Enter` billentyűgombot. Figyelje meg, hogy a számítógép elindul a soros konzolon.

`...`

```
elf: 80080000 <-- 00001000 6586368t + 192624t
elf: entry 80328040
net: interface down
CPU revision is: 000028a0
FPU revision is: 000028a0
Primary instruction cache 32kB, physically tagged, 2-way, linesize 32 bytes.
Primary data cache 32kB 2-way, linesize 32 bytes.
Linux version 2.4.26-mipscvs-20040415 (root@khazad-dum) (gcc version 3.3.3...
Determined physical RAM map:
 memory: 08000000 @ 00000000 (usable)
Initial ramdisk at: 0x80392000 (3366912 bytes)
On node 0 totalpages: 32768
zone(0): 32768 pages.
zone(1): 0 pages.
zone(2): 0 pages.
Kernel command line: console=ttyS0,115200 root=/dev/ram0
Calibrating delay loop... 249.85 BogoMIPS
Memory: 122512k/131072k available (2708k kernel code, 8560k reserved, 3424k dat)

```

Amint megjelenik a busybox ash shell, készen áll a Gentoo Linux telepítésének folytatására. Az alábbi lépések segítenek biztosítani, hogy zökkenőmentesen haladhasson:

CODE **Amikor minden a terv szerint halad...**

```
VFS: Mounted root (ext2 filesystem) readonly.
Freeing unused kernel memory: 280k freed
init started:  BusyBox v1.00-pre10 (2004.04.27-02:55+0000) multi-call binary

Gentoo Linux; http://www.gentoo.org/
 Copyright 2001-2004 Gentoo Technologies, Inc.; Distributed under the GPL

 Gentoo/MIPS Netboot for Cobalt Microserver Machines
 Build Date: April 26th, 2004

 * To configure networking, do the following:

 * For Static IP:
 * /bin/net-setup <IP Address> <Gateway Address> [telnet]

 * For Dynamic IP:
 * /bin/net-setup dhcp [telnet]

 * If you would like a telnetd daemon loaded as well, pass "telnet"
 * As the final argument to /bin/net-setup.

Please press Enter to activate this console.

```

### Hibaelhárítás

Ha a számítógép nem hajlandó letölteni a képfájlt, akkor két dolog állhat a háttérben:

1. Az utasítások nem lettek megfelelően követve, vagy
2. Egy kis gyengéd ösztönzésre van szükség. (Nem! Azonnal tegye le azt a kalapácsot!)

Itt van néhány dolog, amit érdemes ellenőrizni:

- A dhcpd a Cobalt számítógépnek IP-címet ad. Figyelje a rendszer naplóiban a BOOTP kérésekkel kapcsolatos üzeneteket. A tcpdump is hasznos lehet itt.
- A jogosultságok megfelelően vannak beállítva a /nfsroot/ könyvtárban (világosan olvashatónak kell lennie).
- Győződjön meg róla, hogy az NFS szerver működik, és exportálja a /nfsroot/ könyvtárat. Ezt ellenőrizheti a szerveren a exportfs -v paranccsal.

## Telepítő CD használata

Silicon Graphics gépeken lehetőség van CD-ről indítani az operációs rendszerek telepítését. (Például így telepítik az IRIX-et.) Nemrégiben lehetővé vált ilyen bootolható CD-k képeinek létrehozása Gentoo telepítéséhez. Ezek a CD-k hasonló módon működnek.

Jelenleg a Gentoo/MIPS Live CD csak az SGI Indy, Indigo 2 és O2 munkaállomásokkal működik, amelyek R4000 és R5000 sorozatú CPU-kkal vannak felszerelve, azonban a jövőben más platformok is lehetségesek lehetnek.

A Live CD képfájlok megtalálhatók a Gentoo tükörszerveren az experimental/mips/livecd/ könyvtárban.

**Warning**

Ezek a CD-k jelenleg erősen kísérleti jellegűek. Előfordulhat, hogy működnek, vagy nem működnek. Kérjük, jelezze a sikereket vagy kudarcokat a Bugzilla oldalon, a [ezen fórumbejegyzésben](https://forums.gentoo.org/viewtopic.php?t=242518), vagy a #gentoo-mips IRC csatornán.

[← A telepítésről](/wiki/Handbook:MIPS/Installation/About/hu "Previous part") [Home](/wiki/Handbook:MIPS "Handbook:MIPS") [Hálózat beállítása →](/wiki/Handbook:MIPS/Installation/Networking/hu "Next part")