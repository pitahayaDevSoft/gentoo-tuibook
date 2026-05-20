# Bootloader

Other languages:

- [Deutsch](/wiki/Handbook:MIPS/Installation/Bootloader/de "Handbuch:MIPS/Installation/Bootloader (100% translated)")
- [English](/wiki/Handbook:MIPS/Installation/Bootloader "Handbook:MIPS/Installation/Bootloader (100% translated)")
- [español](/wiki/Handbook:MIPS/Installation/Bootloader/es "Manual de Gentoo: MIPS/Instalación/Arranque (100% translated)")
- [français](/wiki/Handbook:MIPS/Installation/Bootloader/fr "Handbook:MIPS/Installation/Bootloader/fr (100% translated)")
- [italiano](/wiki/Handbook:MIPS/Installation/Bootloader/it "Handbook:MIPS/Installation/Bootloader/it (100% translated)")
- magyar
- [polski](/wiki/Handbook:MIPS/Installation/Bootloader/pl "Handbook:MIPS/Installation/Bootloader (100% translated)")
- [português do Brasil](/wiki/Handbook:MIPS/Installation/Bootloader/pt-br "Manual:MIPS/Instalação/Gerenciador de Boot (100% translated)")
- [русский](/wiki/Handbook:MIPS/Installation/Bootloader/ru "Handbook:MIPS/Installation/Bootloader (100% translated)")
- [தமிழ்](/wiki/Handbook:MIPS/Installation/Bootloader/ta "கையேடு:MIPS/நிறுவல்/துவக்கஏற்றி (100% translated)")
- [中文（中国大陆）‎](/wiki/Handbook:MIPS/Installation/Bootloader/zh-cn "手册：MIPS/安装/配置系统引导程序Bootloader (100% translated)")
- [日本語](/wiki/Handbook:MIPS/Installation/Bootloader/ja "ハンドブック:MIPS/インストール/ブートローダー (100% translated)")
- [한국어](/wiki/Handbook:MIPS/Installation/Bootloader/ko "Handbook:MIPS/Installation/Bootloader/ko (100% translated)")

[MIPS kézikönyv](/wiki/Handbook:MIPS/hu "Handbook:MIPS/hu")[A Gentoo Linux telepítése](/wiki/Handbook:MIPS/Full/Installation/hu "Handbook:MIPS/Full/Installation/hu")[A telepítésről](/wiki/Handbook:MIPS/Installation/About/hu "Handbook:MIPS/Installation/About/hu")[Telepítőképfájl kiválasztása](/wiki/Handbook:MIPS/Installation/Media/hu "Handbook:MIPS/Installation/Media/hu")[Hálózat beállítása](/wiki/Handbook:MIPS/Installation/Networking/hu "Handbook:MIPS/Installation/Networking/hu")[Adathordozók előkészítése](/wiki/Handbook:MIPS/Installation/Disks/hu "Handbook:MIPS/Installation/Disks/hu")[Stage fájl](/wiki/Handbook:MIPS/Installation/Stage/hu "Handbook:MIPS/Installation/Stage/hu")[Alaprendszer telepítése](/wiki/Handbook:MIPS/Installation/Base/hu "Handbook:MIPS/Installation/Base/hu")[Kernel beállítása](/wiki/Handbook:MIPS/Installation/Kernel/hu "Handbook:MIPS/Installation/Kernel/hu")[Rendszer beállítása](/wiki/Handbook:MIPS/Installation/System/hu "Handbook:MIPS/Installation/System/hu")[Eszközök telepítése](/wiki/Handbook:MIPS/Installation/Tools/hu "Handbook:MIPS/Installation/Tools/hu")Bootloader beállítása[Telepítés véglegesítése](/wiki/Handbook:MIPS/Installation/Finalizing/hu "Handbook:MIPS/Installation/Finalizing/hu")[Munka a Gentoo rendszerrel](/wiki/Handbook:MIPS/Full/Working/hu "Handbook:MIPS/Full/Working/hu")[Portage bemutatása](/wiki/Handbook:MIPS/Working/Portage/hu "Handbook:MIPS/Working/Portage/hu")[USE jelölőzászlók](/wiki/Handbook:MIPS/Working/USE/hu "Handbook:MIPS/Working/USE/hu")[Portage jellemzői](/wiki/Handbook:MIPS/Working/Features/hu "Handbook:MIPS/Working/Features/hu")[Init-szkript rendszer](/wiki/Handbook:MIPS/Working/Initscripts/hu "Handbook:MIPS/Working/Initscripts/hu")[Környezeti változók](/wiki/Handbook:MIPS/Working/EnvVar/hu "Handbook:MIPS/Working/EnvVar/hu")[Munka a Portage szoftvercsomag-kezelővel](/wiki/Handbook:MIPS/Full/Portage/hu "Handbook:MIPS/Full/Portage/hu")[Fájlok és könyvtárak](/wiki/Handbook:MIPS/Portage/Files/hu "Handbook:MIPS/Portage/Files/hu")[Változók](/wiki/Handbook:MIPS/Portage/Variables/hu "Handbook:MIPS/Portage/Variables/hu")[Szoftverágak keverése](/wiki/Handbook:MIPS/Portage/Branches/hu "Handbook:MIPS/Portage/Branches/hu")[További eszközök](/wiki/Handbook:MIPS/Portage/Tools/hu "Handbook:MIPS/Portage/Tools/hu")[Egyéni szoftvercsomag-tárolók](/wiki/Handbook:MIPS/Portage/CustomTree/hu "Handbook:MIPS/Portage/CustomTree/hu")[Fejlett funkciók](/wiki/Handbook:MIPS/Portage/Advanced/hu "Handbook:MIPS/Portage/Advanced/hu")[Hálózat beállítása OpenRC init-rendszeren](/wiki/Handbook:MIPS/Full/Networking/hu "Handbook:MIPS/Full/Networking/hu")[Munka elkezdése](/wiki/Handbook:MIPS/Networking/Introduction/hu "Handbook:MIPS/Networking/Introduction/hu")[Fejlett beállítások](/wiki/Handbook:MIPS/Networking/Advanced/hu "Handbook:MIPS/Networking/Advanced/hu")[Moduláris hálózat](/wiki/Handbook:MIPS/Networking/Modular/hu "Handbook:MIPS/Networking/Modular/hu")[Vezeték nélküli (Wi-Fi)](/wiki/Handbook:MIPS/Networking/Wireless/hu "Handbook:MIPS/Networking/Wireless/hu")[Funkcionalitás hozzáadása](/wiki/Handbook:MIPS/Networking/Extending/hu "Handbook:MIPS/Networking/Extending/hu")[Dinamikus menedzsment](/wiki/Handbook:MIPS/Networking/Dynamic/hu "Handbook:MIPS/Networking/Dynamic/hu")

## Contents

- [1arcload a Silicon Graphics számítógépekhez](#arcload_a_Silicon_Graphics_sz.C3.A1m.C3.ADt.C3.B3g.C3.A9pekhez)
- [2CoLo a Cobalt MicroServers számára](#CoLo_a_Cobalt_MicroServers_sz.C3.A1m.C3.A1ra)
  - [2.1CoLo telepítése](#CoLo_telep.C3.ADt.C3.A9se)
  - [2.2CoLo beállítása](#CoLo_be.C3.A1ll.C3.ADt.C3.A1sa)
- [3Soros konzol beállítása](#Soros_konzol_be.C3.A1ll.C3.ADt.C3.A1sa)
- [4SGI PROM finomhangolása](#SGI_PROM_finomhangol.C3.A1sa)
  - [4.1Általános PROM beállítások](#.C3.81ltal.C3.A1nos_PROM_be.C3.A1ll.C3.ADt.C3.A1sok)
  - [4.2Direct volume-header booting beállítása](#Direct_volume-header_booting_be.C3.A1ll.C3.ADt.C3.A1sa)
  - [4.3Az arcload beállítása](#Az_arcload_be.C3.A1ll.C3.ADt.C3.A1sa)
- [5Rendszer újraindítása](#Rendszer_.C3.BAjraind.C3.ADt.C3.A1sa)

## arcload a Silicon Graphics számítógépekhez

Az arcload olyan számítógépekhez készült, amelyek 64 bites kerneleket igényelnek, ezért nem használható az arcboot (mivel ezt nehéz 64 bites binárisként lefordítani). Emellett megoldást nyújt azokra a sajátosságokra is, amelyek akkor jelentkeznek, amikor a kerneleket közvetlenül a kötet fejléceiből töltik be. A telepítés folytatásához kezdje az alábbi lépéssel:

`root #` `emerge arcload dvhtool`

Miután ez befejeződött, keresse meg az arcload bináris fájlt a /usr/lib/arcload/ könyvtárban. Most két fájl létezik:

- sashARCS: A 32 bites bináris fájl az Indy, Indigo2 (R4k), Challenge S és O2 rendszerekhez.
- sash64: A 64 bites bináris fájl az Octane/Octane2, Origin 200/2000 és Indigo2 Impact rendszerekhez.

Használja a dvhtool parancsot a megfelelő bináris fájl telepítéséhez a rendszer számára a kötet fejlécébe:

Indy/Indigo2/Challenge S/O2 felhasználók számára:

`root #` `dvhtool --unix-to-vh /usr/lib/arcload/sashARCS sashARCS`

Indigo2 Impact/Octane/Octane2/Origin 200/Origin 2000 felhasználók számára:

`root #` `dvhtool --unix-to-vh /usr/lib/arcload/sash64 sash64`

**Note**

A sashARCS vagy sash64 név használata nem kötelező, kivéve, ha a művelet egy bootolható CD kötetfejlécébe történő telepítésre irányul. Normál adathordozókról történő indítás esetén bármi más név adható a fájlnak, amit a felhasználó kíván.

Most csak használja a dvhtool parancsot annak ellenőrzésére, hogy azok a kötetfejlécben találhatók-e.

`root #` `dvhtool --print-volume-directory`

```
----- directory entries -----
Entry #0, name "sash64", start 4, bytes 55859

```

A arc.cf fájl C-szerű szintaxissal rendelkezik. A teljes részletezéshez, hogy miként lehet beállítani, tekintse meg az arcload oldalt a Linux/MIPS wikiben. Röviden: definiáljon egy sor opciót, amelyek az OSLoadFilename változó használatával indításkor engedélyezhetők vagy letilthatók.

FILE **`arc.cf`** **Példa az arc.cf fájlra**

```
# ARCLoad Configuration

# Some default settings...
append  "root=/dev/sda5";
append  "ro";
append  "console=ttyS0,9600";

# The main definition. ip28 may be changed if desired.
ip28 {
        # Definition for a "working" kernel
        # Select this by setting OSLoadFilename="ip28(working)"
        working {
                description     "SGI Indigo2 Impact R10000\n\r";
                image system    "/working";
        }

        # Definition for a "new" kernel
        # Select this by setting OSLoadFilename="ip28(new)"
        new {
                description     "SGI Indigo2 Impact R10000 - Testing Kernel\n\r";
                image system    "/new";
        }

        # For debugging a kernel
        # Select this by setting OSLoadFilename="ip28(working,debug)"
        # or OSLoadFilename="ip28(new,debug)"
        debug {
                description     "Debug console";
                append          "init=/bin/bash";
        }
}

```

Az arcload-0.5 verziótól kezdve az arc.cf és a bináris kernelképfájlok vagy a kötetfejlécben, vagy egy partíción helyezhetőek el. Az újabb funkció használatához helyezze a fájlokat a /boot/ partícióba (vagy / ha a boot partíció nem különálló). Az arcload a népszerű grub bootloader fájlrendszer-illesztőprogramját használja, és ezért ugyanazokat a fájlrendszereket támogatja.

`root #` `dvhtool --unix-to-vh arc.cf arc.cf
`

`root #` `dvhtool --unix-to-vh /usr/src/linux/vmlinux new`

## CoLo a Cobalt MicroServers számára

### CoLo telepítése

Cobalt szervereken ezek a gépek sokkal kevésbé képes firmware-t tartalmaznak, amelyet a chipre telepítettek. A Cobalt BOOTROM primitív az SGI PROM-hoz képest, és számos komoly korlátozással rendelkezik.

- A kernelfájlok mérete körülbelül 675 kB méretre van korlátozva. A Linux 2.4 jelenlegi mérete szinte lehetetlenné teszi egy ilyen méretű kernel létrehozását. A Linux 2.6 és 3.x teljesen kizárt.
- A készleten lévő firmware nem támogatja a 64 bites kernelképfájlokat (bár ezek jelenleg rendkívül kísérletiek a Cobalt gépeken).
- A shell a legjobb esetben is alapvető funkciókat biztosít.

A korlátozások leküzdése érdekében egy alternatív firmware-t fejlesztettek ki, amelyet CoLo-nak (Cobalt Loader) neveztek el. Ez egy BOOTROM képfájl, amelyet vagy a Cobalt szerver belsejében található chipbe lehet flash-elni, vagy a meglévő firmware-ből betölthető.

**Note**

Ez az útmutató végigvezeti Önt a CoLo beállításának folyamatán, hogy az az alapértelmezett firmware által legyen betöltve. Ez az egyetlen valóban biztonságos és ajánlott módja a CoLo beállításának.

**Warning**

Ha akarja, akkor ezeket a fájlokat a szerverre lehet flashelni, hogy teljesen le legyen lecserélve az eredeti firmware, azonban ebben a vállalkozásban teljes mértékben egyedül van. Ha bármi rosszul sül el, akkor fizikailag távolítsa el a BOOTROM-ot, és programozza újra az alapértelmezett firmware-rel. Ha ez ijesztőnek hangzik akkor NE flaselje a számítógépet. A Gentoo nem vállal felelősséget azért, ami történik, amennyiben ezt a tanácsot figyelmen kívül hagyja.

Most telepítse a CoLo szoftvert. Kezdje a szoftvercsomag letöltésével:

`root #` `emerge --ask sys-boot/colo`

A telepítés után nézzen bele a /usr/lib/colo/ könyvtárba, hogy megtalálja a két fájlt:

- colo-chain.elf – A "kernelképfájl", amelyet az alapértelmezett firmware betöltésére használnak.
- colo-rom-image.bin – Egy ROM képfájl, amelyet a BOOTROM-ba lehet be flashelni.

Kezdje a /boot/ felcsatolásával, majd dobjon egy tömörített másolatot a colo-chain.elf fájlról a /boot/ könyvtárba, ahová a rendszer várja.

`root #` `gzip -9vc /usr/lib/colo/colo-chain.elf > /boot/vmlinux.gz`

### CoLo beállítása

Most, amikor a rendszer először bootol, betölti a CoLo fájlt, amely megjelenít egy menüt a hátsó LCD-n. Az első lehetőség (és az alapértelmezett, amely körülbelül 5 másodperc múlva automatikusan elfogadásra kerül) az adathordozóra történő bootolás. A rendszer ezután megpróbálja felcsatolni az első Linux-partíciót, amelyet talál, és futtatja a default.colo nevű szkriptet. A szintaxis teljes mértékben dokumentált a CoLo dokumentációjában (pillantson bele a /usr/share/doc/colo-X.YY/README.shell.gz fájlba – ahol X.YY a telepített verzió), és nagyon egyszerű.

**Note**

Egy tipp: Amikor kernelt telepít, ajánlott kettő bináris kernelképfájlt létrehozni, kernel.gz.working – egy ismert működőképes bináris kernelképfájlt, és kernel.gz.new – egy éppen most binárisra lefordított kernelképfájlt. Lehetséges szimbolikus linkeket használni, hogy az aktuális 'új' és 'működő' bináris kernelképfájlokra mutassanak, vagy egyszerűen átnevezheti a bináris kernelképfájlokat.

FILE **`default.colo`** **Példa a CoLo beállítására**

```
#:CoLo:#
mount hda1
load /kernel.gz.working
execute root=/dev/sda5 ro console=ttyS0,115200

```

**Note**

A CoLo nem fog betölteni olyan szkriptet, amely nem a `#:CoLo:#` sorral kezdődik. Gondoljon rá úgy, mint a shell szkriptekben a `#!/bin/sh` sor megfelelőjére.

Az is lehetséges, hogy kérdést tegyen fel, például melyik bináris kernelképfájlt és beállítást szeretné betölteni, alapértelmezett időkorláttal. Az alábbi beállítás pontosan ezt teszi: Megkérdezi a felhasználót, hogy melyik bináris kernelképfájlt kívánja használni, majd végrehajtja a kiválasztott bináris képfájlt. A vmlinux.gz.new és vmlinux.gz.working lehetnek tényleges bináris kernelképfájlok, vagy szimbolikus linkek, amelyek az adathordozón lévő bináris kernelképfájlokra mutatnak. A select parancs 50 argumentuma azt jelöli, hogy az első lehetőséggel ("Working") folytassa 50/10 másodperc után.

FILE **`default.colo`** **Menülapú beállítás**

```
#:CoLo:#
lcd "Mounting hda1"
mount hda1
select "Which Kernel?" 50 Working New

goto {menu-option}
var image-name vmlinux.gz.working
goto 3f
@var image-name vmlinux.gz.working
goto 2f
@var image-name vmlinux.gz.new

@lcd "Loading Linux" {image-name}
load /{image-name}
lcd "Booting..."
execute root=/dev/sda5 ro console=ttyS0,115200
boot

```

További részletekért tekintse meg a dokumentációt a /usr/share/doc/colo-VERSION könyvtárban.

## Soros konzol beállítása

Rendben, a jelenlegi Linux telepítés jól indulna, de feltételezi, hogy a felhasználó egy fizikai terminálon jelentkezik be. Cobalt számítógépeken ez különösen problémás, mivel nincs olyan, hogy fizikai terminál.

**Note**

Azok, akiknek van lehetőségük támogatott videó lapkakészlet használatára, kihagyhatják ezt a részt, ha szeretnék.

Először nyisson meg egy szerkesztőt, és kezdje el módosítani a /etc/inittab fájlt. Lentebb a fájlban vegye észre a következőket:

FILE **`/etc/inittab`** **Részlet az inittab fájlból**

```
# SERIAL CONSOLE
#c0:12345:respawn:/sbin/agetty 9600 ttyS0 vt102

# TERMINALS
c1:12345:respawn:/sbin/agetty 38400 tty1 linux
c2:12345:respawn:/sbin/agetty 38400 tty2 linux
c3:12345:respawn:/sbin/agetty 38400 tty3 linux
c4:12345:respawn:/sbin/agetty 38400 tty4 linux
c5:12345:respawn:/sbin/agetty 38400 tty5 linux
c6:12345:respawn:/sbin/agetty 38400 tty6 linux

# What to do at the "Three Finger Salute".
ca:12345:ctrlaltdel:/sbin/shutdown -r now

```

Először, törölje a c0 sor megjegyzést. Alapértelmezés szerint ez 9600 bps terminál baud sebesség használatára van állítva. Cobalt szervereken ezt meg lehet változtatni 115200 értékre, hogy megegyezzen a BOOT ROM által meghatározott baud sebességgel. Ezután ez a rész így fog kinézni. Fej nélküli számítógépen (például Cobalt szervereken) szintén ajánlott a helyi terminálsorok (c1-től c6-ig) megjegyzésének hozzáadása, mivel ezek hajlamosak helytelenül működni, amikor nem tudják megnyitni a /dev/ttyX fájlt.

FILE **`/etc/inittab`** **Részlet az inittab fájlból**

```
# SERIAL CONSOLE
c0:12345:respawn:/sbin/agetty 115200 ttyS0 vt102

# TERMINALS -- These are useless on a headless qube
#c1:12345:respawn:/sbin/agetty 38400 tty1 linux
#c2:12345:respawn:/sbin/agetty 38400 tty2 linux
#c3:12345:respawn:/sbin/agetty 38400 tty3 linux
#c4:12345:respawn:/sbin/agetty 38400 tty4 linux
#c5:12345:respawn:/sbin/agetty 38400 tty5 linux
#c6:12345:respawn:/sbin/agetty 38400 tty6 linux

```

Végül, mondja meg a rendszernek, hogy a helyi soros port megbízható biztonságos terminálként használható. Az /etc/securetty fájlt kell módosítani. Ez tartalmazza a terminálok listáját, amelyeket a rendszer megbízhatónak tekint. Egyszerűen adjon hozzá két sort, amely lehetővé teszi a soros vonal használatát root bejelentkezésekhez.

`root #` `echo 'ttyS0' >> /etc/securetty`

Mostanában a Linux ezt /dev/tts/0 néven is nevezi, ezért ezt is adja hozzá:

`root #` `echo 'tts/0' >> /etc/securetty`

## SGI PROM finomhangolása

### Általános PROM beállítások

A bootloader telepítése után, az számítógép újraindítását követően (amely hamarosan megtörténik), lépjen a System Maintenance Menu menübe, és válassza az `Enter` billentyűgombot a _Command Monitor (5)_ opcióra, hasonlóan a rendszer kezdeti hálózati indításához.

CODE **Menü a bootolást követően**

```
1) Start System
2) Install System Software
3) Run Diagnostics
4) Recover System
5) Enter Command Monitor

```

Adja meg a Volume Header helyét:

`>>` `setenv SystemPartition scsi(0)disk(1)rdisk(0)partition(8)`

Gentoo automatikus indítása:

`>>` `setenv AutoLoad Yes`

Állítsa be az időzónát:

`>>` `setenv TimeZone EST5EDT`

Használja a soros konzolt. Grafikus adaptert használó felhasználóknak az "d1" helyett "g" értéket kell megadniuk:

`>>` `setenv console d1`

Állítsa be a soros konzol baud sebességét. Ez választható, az alapértelmezett beállítás 9600, bár akár 38400-ig terjedő sebességeket is használhat, ha ezt kívánja.

`>>` `setenv dbaud 9600`

Most a következő beállítások attól függenek, hogy hogyan indul a rendszer.

### Direct volume-header booting beállítása

**Note**

Ez a téma teljesség kedvéért került ide. Ajánlott, hogy a felhasználók inkább az "arcload" telepítését vizsgálják meg.

**Note**

Ez csak az Indy, Indigo2 (R4k) és Challenge S eszközökön működik.

Állítsa be a root eszközt a Gentoo root partíciójára, például /dev/sda3:

`>>` `setenv OSLoadPartition <root device>`

Az elérhető kernelfájlok listázásához írja be az "ls" parancsot.

`>>` `setenv OSLoader <kernel name>
`

`>>` `setenv OSLoadFilename <kernel name>`

Határozza meg az átadandó kernelparamétereket:

`>>` `setenv OSLoadOptions <kernel parameters>`

Ha a kernelparaméterek módosítása nélkül szeretné kipróbálni egy kernelt, használja a boot -f PROM parancsot.

`root #` `boot -f new root=/dev/sda5 ro`

### Az arcload beállítása

Az arcload az OSLoadFilename opciót használja annak meghatározására, hogy mely beállításokat kell alkalmazni az arc.cf fájlból. A beállításfájl lényegében egy szkript, amelyben a legfelső szintű blokkok különböző rendszerekhez tartozó bináris boot képfájlokat határoznak meg, ezekben pedig opcionális beállítások találhatók. Így az OSLoadFilename=mysys(serial) beállítása betölti a mysys blokk beállításait, majd felülírja azokat a serial-ben meghatározott további opciókkal.

A példafájlban egy rendszertömb van meghatározva, az ip28, amelyhez működő, új és hibakeresési lehetőségek állnak rendelkezésre. Ezután határozza meg a PROM változókat:

Válassza az arcload fájlt bootloadernek: sash64 vagy sashARCS:

`>>` `setenv OSLoader sash64`

Használja a "working" bináris kernelképfájlt, amely az arc.cf fájl ip28 szakaszában van meghatározva.

`>>` `setenv OSLoadFilename ip28(working)`

Az arcload-0.5 verziótól kezdve a fájlokat már nem szükséges a kötetfejlécbe helyezni. Ezek partícióba is elhelyezhetők. Annak megadásához, hogy az arcload hol keresse a beállításfájlt és a bináris kernelképfájlokat, be kell állítani az OSLoadPartition PROM változót. Ennek pontos értéke attól függ, hogy az adathordozó hol található a SCSI buszon. Használja a SystemPartition PROM változót útmutatóként - csak a partíciószámot kell módosítani.

**Note**

Partitions are numbered starting at 0, not 1 as is the case in Linux.

A kötetfejlécről való betöltéshez használja a 8-as partíciót.

`>>` `setenv OSLoadPartition scsi(0)disk(1)rdisk(0)partition(8)`

Ellenkező esetben adja meg a partíciót és a fájlrendszer típusát:

`>>` `setenv OSLoadPartition scsi(0)disk(1)rdisk(0)partition(0)[ext2]`

## Rendszer újraindítása

Lépjen ki a chrootolt környezetből, és válassza le az összes felcsatolt partíciót. Ezt követően írja be azt az egyetlen mágikus parancsot, amely elindítja a végső, valódi tesztet: reboot.

`(chroot) livecd #` `exit`

`livecd~#` `cd
`

`livecd~#` `umount -l /mnt/gentoo/dev{/shm,/pts,}
`

`livecd~#` `umount -R /mnt/gentoo
`

`livecd~#` `reboot`

Ne feledje el eltávolítani az Live ISO telepítőt, különben ismét elindulhat a számítógépen az újonnan telepített Gentoo rendszer helyett!

Miután újraindította a számítógépet, és belépett a frissen feltelepített Gentoo környezetben, bölcs dolog [véglegesíteni a Gentoo telepítést](/wiki/Handbook:MIPS/Installation/Finalizing/hu "Handbook:MIPS/Installation/Finalizing/hu").

[← Eszközök telepítése](/wiki/Handbook:MIPS/Installation/Tools/hu "Previous part") [Kezdőlap](/wiki/Handbook:MIPS/hu "Handbook:MIPS/hu") [Telepítés véglegesítése →](/wiki/Handbook:MIPS/Installation/Finalizing/hu "Next part")