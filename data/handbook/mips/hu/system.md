# System

Other languages:

- [Deutsch](/wiki/Handbook:MIPS/Installation/System/de "Handbuch:MIPS/Installation/System (100% translated)")
- [English](/wiki/Handbook:MIPS/Installation/System "Handbook:MIPS/Installation/System (100% translated)")
- [español](/wiki/Handbook:MIPS/Installation/System/es "Manual de Gentoo: MIPS/Instalación/Sistema (100% translated)")
- [français](/wiki/Handbook:MIPS/Installation/System/fr "Handbook:MIPS/Installation/System/fr (100% translated)")
- [italiano](/wiki/Handbook:MIPS/Installation/System/it "Handbook:MIPS/Installation/System/it (100% translated)")
- magyar
- [polski](/wiki/Handbook:MIPS/Installation/System/pl "Handbook:MIPS/Installation/System (100% translated)")
- [português do Brasil](/wiki/Handbook:MIPS/Installation/System/pt-br "Manual:MIPS/Instalação/Sistema (100% translated)")
- [русский](/wiki/Handbook:MIPS/Installation/System/ru "Handbook:MIPS/Installation/System (100% translated)")
- [தமிழ்](/wiki/Handbook:MIPS/Installation/System/ta "கையேடு:MIPS/நிறுவல்/முறைமை (100% translated)")
- [中文（中国大陆）‎](/wiki/Handbook:MIPS/Installation/System/zh-cn "手册：MIPS/安装/配置系统 (100% translated)")
- [日本語](/wiki/Handbook:MIPS/Installation/System/ja "ハンドブック:MIPS/インストール/システム (100% translated)")
- [한국어](/wiki/Handbook:MIPS/Installation/System/ko "Handbook:MIPS/Installation/System/ko (100% translated)")

[MIPS kézikönyv](/wiki/Handbook:MIPS/hu "Handbook:MIPS/hu")[A Gentoo Linux telepítése](/wiki/Handbook:MIPS/Full/Installation/hu "Handbook:MIPS/Full/Installation/hu")[A telepítésről](/wiki/Handbook:MIPS/Installation/About/hu "Handbook:MIPS/Installation/About/hu")[Telepítőképfájl kiválasztása](/wiki/Handbook:MIPS/Installation/Media/hu "Handbook:MIPS/Installation/Media/hu")[Hálózat beállítása](/wiki/Handbook:MIPS/Installation/Networking/hu "Handbook:MIPS/Installation/Networking/hu")[Adathordozók előkészítése](/wiki/Handbook:MIPS/Installation/Disks/hu "Handbook:MIPS/Installation/Disks/hu")[Stage fájl](/wiki/Handbook:MIPS/Installation/Stage/hu "Handbook:MIPS/Installation/Stage/hu")[Alaprendszer telepítése](/wiki/Handbook:MIPS/Installation/Base/hu "Handbook:MIPS/Installation/Base/hu")[Kernel beállítása](/wiki/Handbook:MIPS/Installation/Kernel/hu "Handbook:MIPS/Installation/Kernel/hu")Rendszer beállítása[Eszközök telepítése](/wiki/Handbook:MIPS/Installation/Tools/hu "Handbook:MIPS/Installation/Tools/hu")[Bootloader beállítása](/wiki/Handbook:MIPS/Installation/Bootloader/hu "Handbook:MIPS/Installation/Bootloader/hu")[Telepítés véglegesítése](/wiki/Handbook:MIPS/Installation/Finalizing/hu "Handbook:MIPS/Installation/Finalizing/hu")[Munka a Gentoo rendszerrel](/wiki/Handbook:MIPS/Full/Working/hu "Handbook:MIPS/Full/Working/hu")[Portage bemutatása](/wiki/Handbook:MIPS/Working/Portage/hu "Handbook:MIPS/Working/Portage/hu")[USE jelölőzászlók](/wiki/Handbook:MIPS/Working/USE/hu "Handbook:MIPS/Working/USE/hu")[Portage jellemzői](/wiki/Handbook:MIPS/Working/Features/hu "Handbook:MIPS/Working/Features/hu")[Init-szkript rendszer](/wiki/Handbook:MIPS/Working/Initscripts/hu "Handbook:MIPS/Working/Initscripts/hu")[Környezeti változók](/wiki/Handbook:MIPS/Working/EnvVar/hu "Handbook:MIPS/Working/EnvVar/hu")[Munka a Portage szoftvercsomag-kezelővel](/wiki/Handbook:MIPS/Full/Portage/hu "Handbook:MIPS/Full/Portage/hu")[Fájlok és könyvtárak](/wiki/Handbook:MIPS/Portage/Files/hu "Handbook:MIPS/Portage/Files/hu")[Változók](/wiki/Handbook:MIPS/Portage/Variables/hu "Handbook:MIPS/Portage/Variables/hu")[Szoftverágak keverése](/wiki/Handbook:MIPS/Portage/Branches/hu "Handbook:MIPS/Portage/Branches/hu")[További eszközök](/wiki/Handbook:MIPS/Portage/Tools/hu "Handbook:MIPS/Portage/Tools/hu")[Egyéni szoftvercsomag-tárolók](/wiki/Handbook:MIPS/Portage/CustomTree/hu "Handbook:MIPS/Portage/CustomTree/hu")[Fejlett funkciók](/wiki/Handbook:MIPS/Portage/Advanced/hu "Handbook:MIPS/Portage/Advanced/hu")[Hálózat beállítása OpenRC init-rendszeren](/wiki/Handbook:MIPS/Full/Networking/hu "Handbook:MIPS/Full/Networking/hu")[Munka elkezdése](/wiki/Handbook:MIPS/Networking/Introduction/hu "Handbook:MIPS/Networking/Introduction/hu")[Fejlett beállítások](/wiki/Handbook:MIPS/Networking/Advanced/hu "Handbook:MIPS/Networking/Advanced/hu")[Moduláris hálózat](/wiki/Handbook:MIPS/Networking/Modular/hu "Handbook:MIPS/Networking/Modular/hu")[Vezeték nélküli (Wi-Fi)](/wiki/Handbook:MIPS/Networking/Wireless/hu "Handbook:MIPS/Networking/Wireless/hu")[Funkcionalitás hozzáadása](/wiki/Handbook:MIPS/Networking/Extending/hu "Handbook:MIPS/Networking/Extending/hu")[Dinamikus menedzsment](/wiki/Handbook:MIPS/Networking/Dynamic/hu "Handbook:MIPS/Networking/Dynamic/hu")

## Contents

- [1Fájlrendszer információ](#F.C3.A1jlrendszer_inform.C3.A1ci.C3.B3)
  - [1.1Fájlrendszer-címkék és UUID azonosítók](#F.C3.A1jlrendszer-c.C3.ADmk.C3.A9k_.C3.A9s_UUID_azonos.C3.ADt.C3.B3k)
  - [1.2Partíció címkék és UUID-k](#Part.C3.ADci.C3.B3_c.C3.ADmk.C3.A9k_.C3.A9s_UUID-k)
  - [1.3Az fstab fájltól](#Az_fstab_f.C3.A1jlt.C3.B3l)
  - [1.4Az fstab fájl létrehozása](#Az_fstab_f.C3.A1jl_l.C3.A9trehoz.C3.A1sa)
    - [1.4.1DOS/Örökölt BIOS alapú operációs rendszerek](#DOS.2F.C3.96r.C3.B6k.C3.B6lt_BIOS_alap.C3.BA_oper.C3.A1ci.C3.B3s_rendszerek)
- [2Hálózati információk](#H.C3.A1l.C3.B3zati_inform.C3.A1ci.C3.B3k)
  - [2.1Számítógép neve (hostname)](#Sz.C3.A1m.C3.ADt.C3.B3g.C3.A9p_neve_.28hostname.29)
    - [2.1.1Számítógép nevének a beállítása (OpenRC vagy systemd)](#Sz.C3.A1m.C3.ADt.C3.B3g.C3.A9p_nev.C3.A9nek_a_be.C3.A1ll.C3.ADt.C3.A1sa_.28OpenRC_vagy_systemd.29)
    - [2.1.2systemd](#systemd)
  - [2.2Hálózat](#H.C3.A1l.C3.B3zat)
    - [2.2.1DHCP a dhcpcd szolgáltatás által (bármelyik init rendszer esetében)](#DHCP_a_dhcpcd_szolg.C3.A1ltat.C3.A1s_.C3.A1ltal_.28b.C3.A1rmelyik_init_rendszer_eset.C3.A9ben.29)
    - [2.2.2netifrc (OpenRC)](#netifrc_.28OpenRC.29)
      - [2.2.2.1Hálózat beállítása](#H.C3.A1l.C3.B3zat_be.C3.A1ll.C3.ADt.C3.A1sa)
      - [2.2.2.2Hálózat automatikus elindítása a bootoláskor](#H.C3.A1l.C3.B3zat_automatikus_elind.C3.ADt.C3.A1sa_a_bootol.C3.A1skor)
  - [2.3A hosts fájl](#A_hosts_f.C3.A1jl)
- [3Rendszerinformációk](#Rendszerinform.C3.A1ci.C3.B3k)
  - [3.1Root jelszó](#Root_jelsz.C3.B3)
  - [3.2Az init és a boot beállítása](#Az_init_.C3.A9s_a_boot_be.C3.A1ll.C3.ADt.C3.A1sa)
    - [3.2.1OpenRC](#OpenRC)
    - [3.2.2systemd](#systemd_2)

## Fájlrendszer információ

### Fájlrendszer-címkék és UUID azonosítók

Mind az MBR táblázatos séma (BIOS), mind a GPT táblázatos séma, támogatják a bennük létrehozható _fájlrendszer_ címkéket és _fájlrendszer_ UUID-kat. Ezek az attribútumok megadhatók a /etc/fstab fájlban, alternatívaként a mount parancshoz a blokkeszközök keresése és felcsatolása során. A fájlrendszer címkék és az UUID-k a LABEL és UUID előtagokkal azonosíthatók, és a blkid parancs segítségével tekinthetők meg.

`root #` `blkid`

**Warning**

Ha a partíción belüli fájlrendszer törlésre kerül, akkor a fájlrendszer címke és az UUID értékek ezt követően módosulnak vagy eltűnnek.

Az egyediség érdekében az MBR-stílusú partíciós táblázatokat használó olvasóknak azt javasoljuk, hogy a /etc/fstab fájlban a kötetek megadásához címkék helyett UUID azonosítókat használjanak.

**Important**

A fájlrendszer UUID-jai egy LVM kötetben és annak LVM pillanatképeiben azonosak, ezért az UUID-k használatát az LVM kötetek csatolásához kerülni kell.

### Partíció címkék és UUID-k

Azok a rendszerek, amelyek GPT táblázat támogatással rendelkeznek, további 'robosztus' lehetőségeket kínálnak a partíciók meghatározásához a /etc/fstab fájlban. A partíció címkék és a partíció UUID-k használhatók a blokkeszköz egyes partícióinak azonosítására, függetlenül attól, hogy milyen fájlrendszert választottak a partícióhoz. A partíció címkék és az UUID-k a PARTLABEL és/vagy PARTUUID előtagokkal azonosíthatók, és szépen megtekinthetők a parancssorban a blkid parancs futtatásával.

Az **amd64** EFI rendszer kimenete, amely a felfedezhető partíció specifikáció UUID-it használja, az alábbiakhoz hasonló lehet:

`root #` `blkid`

```
/dev/sr0: BLOCK_SIZE="2048" UUID="2023-08-28-03-54-40-00" LABEL="ISOIMAGE" TYPE="iso9660" PTTYPE="PMBR"
/dev/loop0: TYPE="squashfs"
/dev/sda2: PARTUUID="0657fd6d-a4ab-43c4-84e5-0933c84b4f4f"
/dev/sda3: PARTUUID="1cdf763a-5b4c-4dbf-99db-a056c504e8b2"
/dev/sda1: PARTUUID="c12a7328-f81f-11d2-ba4b-00a0c93ec93b"

```

Míg a partíció címkék esetében ez nem mindig igaz, a UUID használata a partíció azonosítására a fstab fájlban garantálja, hogy a bootloader nem fog összezavarodni egy adott kötet keresésekor, még akkor sem, ha a _fájlrendszer_ megváltozik vagy újraírásra kerül a jövőben. Az idősebb, alapértelmezett blokk eszköz fájlok használata (/dev/sd\*N) a partíciók meghatározásához a fstab fájlban kockázatos azoknál az operációs rendszereknél, ahol rendszeresen adnak hozzá vagy távolítanak el SATA blokk eszközöket.

A blokkeszközfájlok elnevezése számos tényezőtől függ, beleértve azt is, hogy hogyan és milyen sorrendben vannak a adathordozók csatlakoztatva a fájlrendszerhez. Előfordulhat, hogy más sorrendben jelennek meg, attól függően, hogy az indítási folyamat elején a kernel melyik eszközöket észleli először. Ezt figyelembe véve, hacsak a rendszergazda nem szándékozik állandóan az adathordozók sorrendjével babrálni, az alapértelmezett blokkeszközfájlok használata egyszerű és egyértelmű megközelítés.

### Az fstab fájltól

Linux alatt az operációs rendszer által használt összes partíciót fel kell tüntetni a [/etc/fstab](/wiki//etc/fstab/hu "/etc/fstab/hu") fájlban. Ez a fájl tartalmazza az adott partíciók csatolási pontjait (ahol a fájlrendszer szerkezetében láthatók), hogy miként kell őket felcsatolni és milyen speciális opciókkal kell azt megtenni (automatikusan/nem automatikusan, a felhasználók felcsatolhatják-e a partíciókat vagy sem, stb.).

### Az fstab fájl létrehozása

**Note**

Ha a használt init rendszer a systemd, akkor a partíció UUID-k megfelelnek a felfedezhető partíció specifikációnak (DPS), ahogy az a [Preparing the disks](/wiki/Handbook:MIPS/Installation/Disks/hu "Handbook:MIPS/Installation/Disks/hu") részben meg van adva, és a rendszer UEFI-t használ, akkor az fstab létrehozását kihagyhatjuk, mivel a systemd automatikusan felcsatolja azokat a partíciókat, amelyek követik a specifikációt.

A /etc/fstab fájl táblázatos szintaxist használ. Minden sor hat mezőből áll, amelyeket szóközök, tabulátorok vagy ezek keveréke választ el. Minden mezőnek megvan a saját jelentése:

1. Az első mező a csatolandó blokk-speciális eszközt vagy távoli fájlrendszert mutatja. Számos eszközazonosító típus érhető el a blokk-speciális eszköz csomópontokhoz, beleértve az eszközfájlok elérési útjait, a fájlrendszer címkéket és az UUID-kat, valamint a partíció címkéket és az UUID-kat.
2. A második mező a csatolási pontot mutatja, ahova a partíciót csatolni kell.
3. A harmadik mező a partíció által használt fájlrendszer típusát mutatja.
4. A negyedik mező a mount parancs által használt csatolási opciókat mutatja, amikor csatolni szeretné a partíciót. Mivel minden fájlrendszernek megvannak a saját csatolási opciói, a rendszergazdáknak ajánlott elolvasni a mount kézikönyv oldalát (man mount) a teljes lista megtekintéséhez. Több csatolási opció vesszővel van elválasztva.
5. Az ötödik mezőt a dump használja annak meghatározására, hogy a partíciót ki kell-e írni (dump) vagy sem. Ezt általában nullára (0) lehet hagyni.
6. A hatodik mezőt az fsck használja annak meghatározására, hogy milyen sorrendben kell ellenőrizni a fájlrendszereket, ha a rendszer nem lett megfelelően leállítva. A gyökérfájlrendszernek `1`-et kell kapnia, míg a többi `2`-t (vagy `0`-át, ha nincs szükség fájlrendszer ellenőrzésre).

**Important**

A Gentoo fokozat (stage) fájlokban található alapértelmezett /etc/fstab fájl _nem_ egy érvényes fstab fájl, hanem egy sablon, amely releváns értékek bevitelére használható.

`root #` `nano /etc/fstab`

#### DOS/Örökölt BIOS alapú operációs rendszerek

Nézzük meg, hogyan írjuk le az opciókat a /boot partícióhoz. Ez csak egy példa, és az előzőleg azonosított partíciós döntések alapján módosítani kell a telepítés során.
A **mips** partíciós példában a /boot általában a /dev/sda1 partíció, a fájlrendszerhez pedig a xfs ajánlott. Ezt ellenőrizni kell az indítás során, így a következőképpen írnánk le:

FILE **`/etc/fstab`** **Egy példa egy DOS/Örökölt BIOS boot sorra az /etc/fstab fájlban:**

```
Alkalmazkodjon bármilyen formázási különbséghez és/vagy az "Előkészítés a lemezekhez" lépés során létrehozott további partíciókhoz
/dev/sda1   /boot     ext2    defaults        0 2

```

Néhány rendszergazda szeretné, ha a /boot partíció nem kerülne automatikusan csatolásra a rendszer biztonságának javítása érdekében. Ezeknek az embereknek a `defaults` helyett a `noauto` szót kell használniuk. Ez azt jelenti, hogy azok a felhasználók manuálisan kell, hogy csatolják ezt a partíciót minden alkalommal, amikor használni szeretnék.

Adja hozzá azokat a szabályokat, amelyek megfelelnek a korábban meghatározott partíciós sémának, és adjon hozzá szabályokat az olyan eszközökkel kapcsolatban, mint a DVD-ROM meghajtó(k), és természetesen. Ha egyéb partíciókat vagy meghajtókat is használ, akkor azokat is adja hozzá szabályként a fájlhoz.

Alább látható egy részletesebb példa egy /etc/fstab fájlra:

FILE **`/etc/fstab`** **Egy teljes /etc/fstab példa egy DOS/Örökölt BIOS rendszerhez**

```
# Alkalmazkodjon bármilyen formázási különbséghez és/vagy az "Előkészítés a lemezekhez" lépés során létrehozott további partíciókhoz.
/dev/sda1   /boot        ext2    defaults    0 2
/dev/sda10   none         swap    sw                   0 0
/dev/sda5   /            xfs    defaults,noatime              0 1

/dev/cdrom  /mnt/cdrom   auto    noauto,user          0 0

```

Amikor az `auto` van használva a harmadik mezőben, a mount parancs kitalálja, hogy milyen fájlrendszer lesz az amit fel kell csatolni. Ez javasolt az eltávolítható adathordozókhoz, mivel sokféle fájlrendszerrel hozhatók létre. A negyedik mezőben lévő `user` opció lehetővé teszi, hogy nem-root felhasználók is csatolják a CD-t.

A teljesítmény javítása érdekében a legtöbb felhasználó szeretné hozzáadni a `noatime` csatolási opciót, amely gyorsabb rendszert eredményez, mivel a hozzáférési időket nem rögzítik (általában nincs is rájuk szükség). Ez szintén ajánlott szilárdtest-meghajtóval (SSD-vel) rendelkező rendszerek esetében. A felhasználók fontolóra vehetik a `lazytime` opció használatát is.

**Tip**

A teljesítmény romlása miatt nem ajánlott a `discard` csatolási opció megadása a /etc/fstab fájlban. Általában jobb időzített blokktörléseket ütemezni egy feladatütemező, például a cron vagy egy időzítő (systemd) segítségével. További információkért lásd: [Időzített fstrim feladatok](/wiki/SSD#Periodic_fstrim_jobs "SSD").

Ellenőrizze kétszer a /etc/fstab fájlt, majd mentse és lépjen ki a folytatáshoz.

## Hálózati információk

Fontos megjegyezni, hogy az alábbi szakaszok célja, hogy segítsenek az olvasónak gyorsan beállítani a rendszerét egy helyi hálózat használatához.

OpenRC-t futtató rendszerek esetében részletesebb hivatkozás a hálózati beállításokhoz a kézikönyv végén található [haladó hálózati beállítás](/wiki/Handbook:MIPS/Networking/Introduction/hu "Handbook:MIPS/Networking/Introduction/hu") szakaszban érhető el. Azok a rendszerek, amelyek specifikusabb hálózati igényekkel rendelkeznek, átugorhatják ezt a részt, majd visszatérhetnek ide, hogy folytassák a telepítést.

A rendszer-specifikusabb systemd hálózati beállításokért kérjük, hogy tekintse meg a [hálózat](/wiki/Systemd/hu#H.C3.A1l.C3.B3zat "Systemd/hu") részt a [systemd](/wiki/Systemd/hu "Systemd/hu") cikkben.

### Számítógép neve (hostname)

Az egyik döntés, amelyet a rendszergazdának meg kell hoznia, az, hogy nevet adjon a számítógépének. Ez elég könnyűnek tűnik, de sok felhasználó nehézségekbe ütközik a megfelelő számítógépnév megtalálásában. A folyamat felgyorsítása érdekében tudni kell, hogy a döntés nem végleges - később megváltoztatható. Az alábbi példákban a _tux_ számítógépnév szerepel.

#### Számítógép nevének a beállítása (OpenRC vagy systemd)

`root #` `echo tux > /etc/hostname`

#### systemd

A systemd init-rendszerrel _futó_ operációs rendszer esetében a számítógép nevének a beállítása a hostnamectl segédprogrammal végezhető el. A telepítési folyamat során azonban a [systemd-firstboot](/wiki/Handbook:MIPS/Installation/System/hu#Init_and_boot_configuration_systemd "Handbook:MIPS/Installation/System/hu") parancsot kell használni (további részletek később a kézikönyvben).

A "tux" számítógépnév beállításához a következő parancsot kell futtatni:

`root #` `hostnamectl hostname tux`

Futtassa a hostnamectl --help vagy man 1 hostnamectl parancsot a súgó megtekintéséhez.

### Hálózat

_Számos_ lehetőség áll rendelkezésre a hálózati interfészek beállításához. Ez a szakasz csak néhány módszert ismertet. Válassza ki azt, amelyik a legmegfelelőbbnek tűnik a szükséges beállításhoz.

#### DHCP a dhcpcd szolgáltatás által (bármelyik init rendszer esetében)

A legtöbb helyi hálózat (LAN) üzemeltet DHCP szervert. Ha ez a helyzet, akkor az IP-cím megszerzéséhez ajánlott a dhcpcd program használata.

A telepítéshez futtassa a következő parancsot:

`root #` `emerge --ask net-misc/dhcpcd`

Az OpenRC init-rendszert használó operációs rendszerek esetében, futtassa a következő parancsot a szolgáltatás bekapcsolásához és azonnali elindításához:

`root #` `rc-update add dhcpcd default
`

`root #` `rc-service dhcpcd start
`

A systemd init-rendszert használó operációs rendszerek esetében, futtassa a következő parancsot a szolgáltatás bekapcsolásának érdekében:

`root #` `systemctl enable dhcpcd`

Ezeknek a lépéseknek a befejezése után az operációs rendszer következő indításakor a dhcpcd IP-címet kell, hogy szerezzen a DHCP szervertől. További részletekért tekintse meg a [dhcpcd](/wiki/Dhcpcd/hu "Dhcpcd/hu") cikket.

#### netifrc (OpenRC)

**Tip**

Ez az egyik módja a hálózat beállításának [Netifrc](/wiki/Netifrc "Netifrc") használatával OpenRC rendszereken. Egyszerűbb beállításokhoz más módszerek is léteznek, például a [Dhcpcd](/wiki/Dhcpcd/hu "Dhcpcd/hu") használata.

##### Hálózat beállítása

A Gentoo Linux telepítése során a hálózat már be lett állítva. Azonban ez csak az Live élő bootolható telepítőnek a környezetére vonatkozott, és nem a számítógépre feltelepített környezetre. Most a számítógépre feltelepített Gentoo Linux rendszer számára készül el a hálózati beállítás.

**Note**

Részletesebb információk a hálózatról, beleértve az olyan haladó témákat, mint a bonding, bridging, 802.1Q VLAN-ok vagy a vezeték nélküli hálózatok, a [haladó hálózati beállítás](/wiki/Handbook:MIPS/Networking/Introduction/hu "Handbook:MIPS/Networking/Introduction/hu") szakaszban találhatók.

Minden hálózati információ a /etc/conf.d/net fájlban található. Egyszerű - bár nem feltétlenül intuitív - szintaxist használ. Ne aggódjon! Minden meg van magyarázva. Sok különböző beállítást lefedő, teljesen jól kommentezett példa található a /usr/share/doc/netifrc-\*/net.example.bz2 fájlban.

Először telepítse a [net-misc/netifrc](https://packages.gentoo.org/packages/net-misc/netifrc) szoftvercsomagot:

`root #` `emerge --ask --noreplace net-misc/netifrc`

Alapértelmezettként a DHCP van használva. Ahhoz, hogy a DHCP működjön, szükséges egy DHCP kliens telepítése. Ennek részleteit később, a Szükséges rendszereszközök telepítése részben találja.

Ha a hálózati kapcsolatot speciális DHCP opciók miatt vagy azért kell beállítani, mert a DHCP szolgáltatást egyáltalán nem szeretné használni, akkor nyissa meg a /etc/conf.d/net fájlt:

`root #` `nano /etc/conf.d/net`

Állítsa a config\_eth0 és a routes\_eth0 változót, hogy megadja az IP-cím információkat és az útvonal információkat:

**Note**

Ez feltételezi, hogy a hálózati interfész neve eth0 lesz. Ez azonban nagyon rendszerspecifikus. Ajánlott feltételezni, hogy az interfész neve megegyezik azzal az interfész névvel, amelyet a telepítési adathordozó indításakor látunk, ha a telepítési adathordozó elég friss. További információk a [Hálózati interfész nevek](/wiki/Handbook:MIPS/Networking/Advanced/hu#Network_interface_naming "Handbook:MIPS/Networking/Advanced/hu") szakaszban találhatók.

FILE **`/etc/conf.d/net`** **Statikus IP-cím megadása**

```
config_eth0="192.168.0.2 netmask 255.255.255.0 brd 192.168.0.255"
routes_eth0="default via 192.168.0.1"

```

A DHCP használatának az érdekében, definiálja a config\_eth0 változónak az értékét:

FILE **`/etc/conf.d/net`** **DHCP megadása**

```
config_eth0="dhcp"

```

Olvassa el a /usr/share/doc/netifrc-\*/net.example.bz2 fájlt a további beállítási lehetőségek listájáért. Ha speciális DHCP opciókat kell beállítani, akkor mindenképp olvassa el a DHCP kliens man súgóoldalát is.

Ha az operációs rendszernek több hálózati interfésze van, akkor ismételje meg a fenti lépéseket config\_eth1, config\_eth2 stb. esetén is.

Most mentse el a beállítást, és lépjen ki a folytatáshoz.

##### Hálózat automatikus elindítása a bootoláskor

Ahhoz, hogy a hálózati interfészek a számítógép indulásakor önmaguktól aktiválódjanak, hozzá kell adni őket az alapértelmezett futási szinthez.

`root #` `cd /etc/init.d
`

`root #` `ln -s net.lo net.eth0
`

`root #` `rc-update add net.eth0 default`

Ha az operációs rendszernek több hálózati interfésze van, akkor a megfelelő net.\* fájlokat létre kell hozni, ahogyan azt a net.eth0 esetében is tettük.

Ha a rendszer indítása után kiderül, hogy a hálózati interfész neve (ami jelenleg `eth0`-ként van dokumentálva) hibás volt, akkor hajtsa végre a következő lépéseket a helyesbítéshez:

1. Frissítse a /etc/conf.d/net fájlt a helyes interfész névvel (például `enp3s0` vagy `enp5s0`, a `eth0` helyett).
2. Hozzon létre új szimbolikus linket (például /etc/init.d/net.enp3s0).
3. Távolítsa el a régi szimbolikus linket (rm /etc/init.d/net.eth0).
4. Adja hozzá az újat az alapértelmezett futási szinthez.
5. Távolítsa el a régit a következő parancs használatával: rc-update del net.eth0 default.

### A hosts fájl

Fontos következő lépés lehet tájékoztatni az újonnan telepített operációs rendszert a hálózati környezetében lévő egyéb számítógépekről. A hálózaton lévő számítógépeket (hosztokat) a /etc/hosts fájlban lehet megadni. A hosztnevek hozzáadása itt lehetővé teszi a hosztnevek IP-címekre való feloldását azoknál a hosztoknál, amelyeket a névszerver nem old fel.

`root #` `nano /etc/hosts`

FILE **`/etc/hosts`** **Hálózati információk kitöltése**

```
# Ez határozza meg az aktuális rendszert, és be kell állítani.
127.0.0.1     tux.homenetwork tux localhost
::1           tux.homenetwork tux localhost

# Opcionális meghatározás további rendszerekhez a hálózaton.
192.168.0.5   jenny.homenetwork jenny
192.168.0.6   benny.homenetwork benny

```

Mentse el és lépjen ki a szövegszerkesztőből a folytatáshoz.

## Rendszerinformációk

### Root jelszó

Állítsa be a root jelszót a passwd parancs használatával.

`root #` `passwd`

Később létrehozunk egy további felhasználói fiókot a napi műveletekhez.

### Az init és a boot beállítása

#### OpenRC

Amikor az OpenRC init rendszer van használatban a Gentoo operációs rendszeren, akkor az OpenRC az /etc/rc.conf fájl segítségével állítja be a szolgáltatásokat, továbbá azzal állítja be az indítást és a leállítást az operációs rendszer számára. Nyissa meg az /etc/rc.conf fájlt, és élvezze az összes megjegyzést a fájlban. Tekintse át a beállításokat, és változtassa meg, ahol szükséges.

`root #` `nano /etc/rc.conf`

Ezután nyissa meg az /etc/conf.d/keymaps fájlt a billentyűzet beállításainak kezeléséhez. Szerkessze a fájlt a megfelelő billentyűzet kiválasztásának és a kiválasztott billentyűzet beállításának az érdekében.

`root #` `nano /etc/conf.d/keymaps`

Különös gonddal kezelje a keymap változót. Ha a rossz billentyűzetkiosztást választja ki, akkor furcsa eredmények jelenhetnek meg a billentyűzet használata során.

Végül szerkessze az /etc/conf.d/hwclock fájlt azért, hogy be legyen állítva az óra. Szerkessze az Ön személyes preferenciáinak megfelelően.

`root #` `nano /etc/conf.d/hwclock`

Ha a hardveróra nem az UTC időszámítást használja, akkor a fájlban be kell állítani a következőt: `clock="local"`. Ellenkező esetben az operációs rendszer időeltolódási viselkedést mutathat.

#### systemd

Először ajánlott futtatni a systemd-machine-id-setup és a systemd-firstboot parancsokat, amelyek előkészítik az operációs rendszerünk különböző komponenseit az első indításra az új systemd környezetben. A következő opciók átadása során fel lesz kérve a felhasználót, hogy állítson be nyelvi beállításokat, időzónát, hosztnevet, root jelszót és root shell értékeket. Ez a telepítéshez véletlenszerű számítógép ID-t (gépazonosítót) is hozzárendel:

`root #` `systemd-machine-id-setup`

`root #` `systemd-firstboot --prompt`

Ezután a felhasználóknak futtatniuk kell a systemctl parancsot, hogy az összes telepített unit fájlt visszaállítsák az előre beállított szabályértékekre:

**Note**

Lehetséges, hogy a következő parancs futtatása után hibaüzenet jelenik meg. Ez normális jelenség, mivel a Gentoo LiveCD az OpenRC init rendszert használja.

`root #` `systemctl preset-all --preset-mode=enable-only`

Lehetséges az összes előre beállított módosítás végrehajtása, de ez alaphelyzetbe állíthatja a már beállított szolgáltatásokat a folyamat során:

`root #` `systemctl preset-all`

Ez a két lépés segít biztosítani a zökkenőmentes átmenetet az Live ISO környezetről a telepítés első rendszerindítására.

[← Kernel beállítása](/wiki/Handbook:MIPS/Installation/Kernel/hu "Previous part") [Kezdőlap](/wiki/Handbook:MIPS/hu "Handbook:MIPS/hu") [Eszközök telepítése →](/wiki/Handbook:MIPS/Installation/Tools/hu "Next part")