# Disks

Other languages:

- [Deutsch](/wiki/Handbook:HPPA/Installation/Disks/de "Handbuch:HPPA/Installation/Festplatten (100% translated)")
- [English](/wiki/Handbook:HPPA/Installation/Disks "Handbook:HPPA/Installation/Disks (100% translated)")
- [español](/wiki/Handbook:HPPA/Installation/Disks/es "Manual de Gentoo: HPPA/Instalación/Discos (100% translated)")
- [français](/wiki/Handbook:HPPA/Installation/Disks/fr "Handbook:HPPA/Installation/Disks/fr (100% translated)")
- [italiano](/wiki/Handbook:HPPA/Installation/Disks/it "Handbook:HPPA/Installation/Disks/it (100% translated)")
- magyar
- [polski](/wiki/Handbook:HPPA/Installation/Disks/pl "Handbook:HPPA/Installation/Disks (100% translated)")
- [português do Brasil](/wiki/Handbook:HPPA/Installation/Disks/pt-br "Manual:HPPA/Instalação/Discos (100% translated)")
- [русский](/wiki/Handbook:HPPA/Installation/Disks/ru "Handbook:HPPA/Installation/Disks (100% translated)")
- [தமிழ்](/wiki/Handbook:HPPA/Installation/Disks/ta "கையேடு:HPPA/நிறுவல்/வட்டுக்கள் (100% translated)")
- [中文（中国大陆）‎](/wiki/Handbook:HPPA/Installation/Disks/zh-cn "手册：HPPA/安装/准备磁盘 (100% translated)")
- [日本語](/wiki/Handbook:HPPA/Installation/Disks/ja "ハンドブック:HPPA/インストール/ディスク (100% translated)")
- [한국어](/wiki/Handbook:HPPA/Installation/Disks/ko "Handbook:HPPA/Installation/Disks/ko (100% translated)")

[HPPA kézikönyv](/wiki/Handbook:HPPA/hu "Handbook:HPPA/hu")[A Gentoo Linux telepítése](/wiki/Handbook:HPPA/Full/Installation/hu "Handbook:HPPA/Full/Installation/hu")[A telepítésről](/wiki/Handbook:HPPA/Installation/About/hu "Handbook:HPPA/Installation/About/hu")[Telepítőképfájl kiválasztása](/wiki/Handbook:HPPA/Installation/Media/hu "Handbook:HPPA/Installation/Media/hu")[Hálózat beállítása](/wiki/Handbook:HPPA/Installation/Networking/hu "Handbook:HPPA/Installation/Networking/hu")Adathordozók előkészítése[Stage fájl](/wiki/Handbook:HPPA/Installation/Stage/hu "Handbook:HPPA/Installation/Stage/hu")[Alaprendszer telepítése](/wiki/Handbook:HPPA/Installation/Base/hu "Handbook:HPPA/Installation/Base/hu")[Kernel beállítása](/wiki/Handbook:HPPA/Installation/Kernel/hu "Handbook:HPPA/Installation/Kernel/hu")[Rendszer beállítása](/wiki/Handbook:HPPA/Installation/System/hu "Handbook:HPPA/Installation/System/hu")[Eszközök telepítése](/wiki/Handbook:HPPA/Installation/Tools/hu "Handbook:HPPA/Installation/Tools/hu")[Bootloader beállítása](/wiki/Handbook:HPPA/Installation/Bootloader/hu "Handbook:HPPA/Installation/Bootloader/hu")[Telepítés véglegesítése](/wiki/Handbook:HPPA/Installation/Finalizing/hu "Handbook:HPPA/Installation/Finalizing/hu")[Munka a Gentoo rendszerrel](/wiki/Handbook:HPPA/Full/Working/hu "Handbook:HPPA/Full/Working/hu")[Portage bemutatása](/wiki/Handbook:HPPA/Working/Portage/hu "Handbook:HPPA/Working/Portage/hu")[USE jelölőzászlók](/wiki/Handbook:HPPA/Working/USE/hu "Handbook:HPPA/Working/USE/hu")[Portage jellemzői](/wiki/Handbook:HPPA/Working/Features/hu "Handbook:HPPA/Working/Features/hu")[Init-szkript rendszer](/wiki/Handbook:HPPA/Working/Initscripts/hu "Handbook:HPPA/Working/Initscripts/hu")[Környezeti változók](/wiki/Handbook:HPPA/Working/EnvVar/hu "Handbook:HPPA/Working/EnvVar/hu")[Munka a Portage szoftvercsomag-kezelővel](/wiki/Handbook:HPPA/Full/Portage/hu "Handbook:HPPA/Full/Portage/hu")[Fájlok és könyvtárak](/wiki/Handbook:HPPA/Portage/Files/hu "Handbook:HPPA/Portage/Files/hu")[Változók](/wiki/Handbook:HPPA/Portage/Variables/hu "Handbook:HPPA/Portage/Variables/hu")[Szoftverágak keverése](/wiki/Handbook:HPPA/Portage/Branches/hu "Handbook:HPPA/Portage/Branches/hu")[További eszközök](/wiki/Handbook:HPPA/Portage/Tools/hu "Handbook:HPPA/Portage/Tools/hu")[Egyéni szoftvercsomag-tárolók](/wiki/Handbook:HPPA/Portage/CustomTree/hu "Handbook:HPPA/Portage/CustomTree/hu")[Fejlett funkciók](/wiki/Handbook:HPPA/Portage/Advanced/hu "Handbook:HPPA/Portage/Advanced/hu")[Hálózat beállítása OpenRC init-rendszeren](/wiki/Handbook:HPPA/Full/Networking/hu "Handbook:HPPA/Full/Networking/hu")[Munka elkezdése](/wiki/Handbook:HPPA/Networking/Introduction/hu "Handbook:HPPA/Networking/Introduction/hu")[Fejlett beállítások](/wiki/Handbook:HPPA/Networking/Advanced/hu "Handbook:HPPA/Networking/Advanced/hu")[Moduláris hálózat](/wiki/Handbook:HPPA/Networking/Modular/hu "Handbook:HPPA/Networking/Modular/hu")[Vezeték nélküli (Wi-Fi)](/wiki/Handbook:HPPA/Networking/Wireless/hu "Handbook:HPPA/Networking/Wireless/hu")[Funkcionalitás hozzáadása](/wiki/Handbook:HPPA/Networking/Extending/hu "Handbook:HPPA/Networking/Extending/hu")[Dinamikus menedzsment](/wiki/Handbook:HPPA/Networking/Dynamic/hu "Handbook:HPPA/Networking/Dynamic/hu")

## Contents

- [1Bevezetés a blokktípusú eszközökbe](#Bevezet.C3.A9s_a_blokkt.C3.ADpus.C3.BA_eszk.C3.B6z.C3.B6kbe)
  - [1.1Blokkeszközök](#Blokkeszk.C3.B6z.C3.B6k)
  - [1.2Partíciók és szeletek](#Part.C3.ADci.C3.B3k_.C3.A9s_szeletek)
- [2Partíciós séma megtervezése](#Part.C3.ADci.C3.B3s_s.C3.A9ma_megtervez.C3.A9se)
  - [2.1Hány partíció és mekkora méretű?](#H.C3.A1ny_part.C3.ADci.C3.B3_.C3.A9s_mekkora_m.C3.A9ret.C5.B1.3F)
  - [2.2Mi a helyzet a swap területtel?](#Mi_a_helyzet_a_swap_ter.C3.BClettel.3F)
- [3Az fdisk használata HPPA-n](#Az_fdisk_haszn.C3.A1lata_HPPA-n)
- [4Fájlrendszerek létrehozása](#F.C3.A1jlrendszerek_l.C3.A9trehoz.C3.A1sa)
  - [4.1Bevezetés](#Bevezet.C3.A9s)
  - [4.2Fájlrendszerek](#F.C3.A1jlrendszerek)
  - [4.3Fájlrendszer rárakása egy partícióra](#F.C3.A1jlrendszer_r.C3.A1rak.C3.A1sa_egy_part.C3.ADci.C3.B3ra)
    - [4.3.1Örökölt BIOS rendszerindító partíciónak a fájlrendszere](#.C3.96r.C3.B6k.C3.B6lt_BIOS_rendszerind.C3.ADt.C3.B3_part.C3.ADci.C3.B3nak_a_f.C3.A1jlrendszere)
    - [4.3.2Kicsi ext4 partíciók](#Kicsi_ext4_part.C3.ADci.C3.B3k)
  - [4.4A swap (lapozásra használt) partíció aktiválása](#A_swap_.28lapoz.C3.A1sra_haszn.C3.A1lt.29_part.C3.ADci.C3.B3_aktiv.C3.A1l.C3.A1sa)
- [5Gyökérpartíció (root partíció) felcsatolása](#Gy.C3.B6k.C3.A9rpart.C3.ADci.C3.B3_.28root_part.C3.ADci.C3.B3.29_felcsatol.C3.A1sa)

## Bevezetés a blokktípusú eszközökbe

### Blokkeszközök

Vessen egy pillantást a Gentoo Linux és általában a Linux adathordozó órientelt vonatkozásaira, beleértve a blokkeszközöket, partíciókat és Linux fájlrendszereket. Miután megértette a lemezek csínját-bínját, partíciókat és fájlrendszereket hozhat létre a telepítéshez.

Kezdésként nézzük meg a blokkeszközöket. Az SCSI és a Serial ATA meghajtók is az /dev könyvtár alatt vannak címkézve, mint például: /dev/sda, /dev/sdb, /dev/sdc stb. A modernebb számítógépeken a PCI Express alapú NVMe szilárdtestalapú lemezek esetében olyan könyvtárak találhatók, mint a /dev/nvme0n1, /dev/nvme0n2 stb.

A következő táblázat segít az olvasóknak meghatározni, hogy hol találnak egy bizonyos típusú blokkeszközt a rendszeren:

Készülék típusaKészülék alapértelmezett elérési útvonalaMegjegyzések és megfontolások
IDE, SATA, SAS, SCSI, vagy USB flash/dev/sdaA hardver nagyjából 2007-től kedve egészen a napjainkig létezik. Ez az elérési útvonal talán a leggyakrabban használt a Linux rendszerekben. Az ilyen típusú eszközök [SATA](https://en.wikipedia.org/wiki/Serial_ATA "wikipedia:Serial ATA"), [SCSI](https://en.wikipedia.org/wiki/SCSI "wikipedia:SCSI"), [USB](https://en.wikipedia.org/wiki/USB "wikipedia:USB") buszon keresztül csatolhatóak fel a rendszerünkbe blokktípusú adattároló formájában. Például a legelső SATA készüléken lévő első partíciónak a teljes elérési útvonala a /dev/sda1.
NVM Express (NVMe)/dev/nvme0n1A legújabb szilárdtestalapú technológiát képviselő [NVMe](https://en.wikipedia.org/wiki/NVM_Express "wikipedia:NVM Express") meghajtók a PCI Express buszhoz csatlakoznak, és jelenleg a piacon ezek a leggyorsabb átviteli blokksebességgel rendelkező készülékek. A 2014 körüli és újabb számítógépes rendszerek általában már támogathatják az NVMe hardvert. A legelső NVMe típusú készüléken lévő első partíció elérési útvonala a következő: /dev/nvme0n1p1.
MMC, eMMC, és SD/dev/mmcblk0A [beágyazott MMC-eszközök](https://en.wikipedia.org/wiki/MultiMediaCard#eMMC "wikipedia:MultiMediaCard"), SD-kártyák és más típusú memóriakártyák hasznosak lehetnek az adattároláshoz. Ennek ellenére előfordulhat, hogy sok számítógépes rendszer nem engedélyezi az ilyen típusú eszközökről történő rendszerindítást. Javasoljuk, hogy egyáltalán **ne** használja ezeket a készülékeket aktív Linux telepítő ISO-képfájlokhoz. Fontolja meg inkább ezeknek a kártyáknak a fájlátvitelre való felhasználását, ami a tipikus tervezési szándékuk is egyben. Alternatív megoldásként ez a típusú adattároló hasznos lehet rövidtávú fájlmentések vagy pillanatképek készítéséhez.

A fenti fizikai adathordozó blokkeszközök egy absztrakt interfészt jelenítenek meg a rendszerben látható adathordozó számára. A felhasználói programok ezeket a fizikai adathordozó blokkeszközöket a rendszerben látható adathordozóval való interakciókon keresztül használhatják, anélkül, hogy aggódniuk kellene amiatt, hogy a fizikai adathordozók SATA, SCSI vagy valami más típusú-e. A program egyszerűen meg tudja címezni a fizikai adathordozón lévő tárhelyet mint egy csomó összefüggő, véletlenszerűen elérhető, 4096 bájtos (4K) blokkok csoportja.

### Partíciók és szeletek

Bár elméletileg lehetséges egy teljes adathordozót használni a Linux rendszerhez, ez szinte soha nem történik meg a gyakorlatban. Ehelyett a teljes adathordozó blokkeszközeit kisebb, könnyebben kezelhető blokkeszközökre osztják szét. A legtöbb rendszeren ezeket partícióknak nevezik. Más architektúrák hasonló technikát alkalmaznak, amelyet _szeleteknek_ hívnak.

## Partíciós séma megtervezése

### Hány partíció és mekkora méretű?

Az adathordozón a partíciók elrendezésének a kialakítása nagymértékben függ a Gentoo operációs rendszer igényeitől és az adathordozón alkalmazott fájlrendszer(ek) igényeitől. Ha sok felhasználó lesz a rendszerben, akkor tanácsos a /home könyvtárat külön partícióra elhelyezni, ami növeli a biztonságot, és megkönnyíti a biztonsági mentéseket és más típusú karbantartásokat. Ha a Gentoo rendszert levelezőszerverként telepítik, akkor a /var könyvtárnak külön partíción kell lennie, mivel minden levél a /var könyvtárban lesz eltárolva. A játékszervereknek lehet külön /opt partíciója, mivel a legtöbb játékszerver-szoftver ebbe a könyvtárba van telepítve. Ezeknek az ajánlásoknak az oka hasonló a /home könyvtárhoz: biztonság, biztonsági mentések és karbantartás.

A legtöbb esetben a Gentoo rendszeren az /usr és a /var könyvtárak viszonylag nagy méretűek szoktak lenni. A /usr könyvtár tárolja a rendszeren elérhető alkalmazások többségét és a Linux kernel forráskódokat (a /usr/src alkönyvtárban). Alapértelmezés szerint a /var tárolja a Gentoo ebuild szoftvertárolót (a /var/db/repos/gentoo alkönyvtárban), amely a fájlrendszertől függően általában körülbelül 650 MiB területet foglal el az adathordozón. Ez a becsült terület _nem tartalmazza_ a /var/cache/distfiles és /var/cache/binpkgs könyvtárakat, amelyek fokozatosan megtelnek forráskódfájlokkal, illetve (opcionálisan) bináris szoftvercsomagokkal, ahogy a rendszergazdák hozzáadják azokat a rendszerhez.

Az, hogy hány partíció és mekkora méretű kell, nagymértékben függ a kompromisszumok mérlegelésétől és az adott körülményekhez képest a legjobb választástól. A különálló partícióknak vagy köteteknek a következő előnyei vannak:

- Kiválasztható a legjobban teljesítő fájlrendszer minden partícióhoz vagy kötethez.
- A teljes rendszer nem fogyhat ki a szabad területből, ha az egyik meghibásodott adathordozó elkezd folyamatosan fájlokat írni egy partícióra vagy kötetre.
- Ha szükséges, akkor a fájlrendszer-ellenőrzések időben lerövidülnek, mivel párhuzamosan több ellenőrzés is elvégezhető (bár ez az előny több adathordozó esetében jobban érvényesül, mint a több partíció esetében).
- A biztonság fokozható az egyes partíciók vagy kötetek írásvédett módban történő felcsatlakoztatása által, `nosuid` (a setuid biteket figyelmen kívül hagyva), `noexec` (a végrehajtható biteket figyelmen kívül hagyva) stb.

A több partíciónak azonban vannak bizonyos hátrányai is:

- Ha nincs megfelelően beállítva, akkor előfordulhat, hogy a rendszernek sok szabad területe lesz az egyik partíción, és kevés szabad területe lesz a másikon.
- Az /usr/ könyvtár külön partícióra történő rárakása megkövetelheti a rendszergazdától, hogy az initramfs segítségével indítsa el a rendszert a partíció felcsatlakoztatásának érdekében, még mielőtt más rendszerindító szkriptek elindulnának. Mivel az initramfs generálása és karbantartása túlmutat ennek a kézikönyvnek a hatókörén, **javasoljuk, hogy az újonnan érkező felhasználók ne használjanak külön partíciót az /usr/ könyvtárhoz**.
- Az SCSI és a SATA esetében létezik a 15 partíciós korlát, kivétel ha az adathordozó GPT típusú táblázatot használ.

**Note**

Azon Gentoo operációs rendszerek számára, amelyek a systemd-t szolgáltatásként és init rendszerként kívánják használni, az /usr könyvtárnak elérhetőnek kell lennie a rendszerindításkor, vagy a gyökér fájlrendszer részeként, vagy egy initramfs által felcsatlakoztatva.

### Mi a helyzet a swap területtel?

Ajánlások a swap méretére
RAM méreteFelfüggesztéstámogatás?Hibernációtámogatás?
2 GB vagy kevesebb2 \* RAM3 \* RAM
2 GB-tól 8 GB-igRAM mennyisége2 \* RAM
8 GB-tól 64 GB-ig8 GB minimum, 16 maximum1.5 \* RAM
64 GB vagy nagyobb8 GB minimumNem javasolt a hibernáció! A hibernálás _nem_ ajánlott nagyon nagy mennyiségű memóriával rendelkező rendszerek esetén, mivel a sikeres hibernáláshoz a memória teljes tartalmát a adathordozóra kell írni. Több tíz gigabájt (vagy még rosszabb!) adathordozóra történő kiírása sok időt vehet igénybe, különösen forgókorongos adathordozó lemezek használata esetén. Ha nagyon sok RAM van a rendszerben, akkor a legjobb döntés az, ha ki van kapcsolva a hibernáció.

Valójában, nincs egyáltalán előre kőbevésve, hogy pontosan mekkorának kell lennie az adathordozón a swap területnek. A területnek az a célja, hogy az adathordozón helyet biztosítson a kernel számára, amikor a RAM nagyon intenzív szintű használat alatt áll. A swap terület lehetővé teszi a RAM-ban futó kernel számára, hogy azok a RAM-ban található memórialapok ideiglenesen ki legyenek rakva az adathordozóra, amelyekre hamarosan valószínűleg ismét szüksége lesz a kernelnek a RAM-ban (ezt nevezik kiswapolásnak vagy kilapozásnak a memóriából). Ez a művelet felszabadítja a helyet a RAM-ban az éppen aktuális feladathoz. Természetesen, ha a kernelnek hirtelen ismét szüksége lesz az adathordozóra kiswapolt oldalakra, akkor azokat vissza kell tölteni a RAM-ba (lapozás művelete), ami jóval tovább tart, mint ha csak a RAM-ban zajlana az írás/olvasás munkafolyamata (mivel az RAM-on kívüli adathordozók, különösen a HDD-k, nagyon lassúak a RAM-hoz képest).

Ha egy rendszer nem fog memóriaigényes alkalmazásokat futtatni, vagy sok RAM áll rendelkezésére, akkor valószínűleg nincs szüksége sok swap területre. Hibernálás esetén azonban ne feledje, hogy a swap terület a _memória teljes tartalmának_ a tárolására szolgál (valószínűleg asztali számítógépeket és laptopokat érint, nem szerverkörnyezeteket). Ha a rendszernek szüksége van a hibernált állapot támogatására, akkor a memória mennyiségénél nagyobb vagy azzal megegyező swap területre van szükség.

Általános szabály, hogy 4 GB-nál kisebb RAM esetén a swap terület mérete a RAM kétszerese legyen. Több adathordozóval rendelkező operációs rendszerek esetén célszerű minden adathordozón egy swap partíciót létrehozni, hogy párhuzamos olvasási/írási műveletekhez használhatók legyenek. Minél gyorsabban tud "swap"-olni egy adathordozót, annál gyorsabban fog futni a rendszer, amikor a swap területen lévő adatokhoz kell hozzáférni. Amikor a fizikailag forgólemezes és a szilárdtestalapú adathordozók között választunk, akkor a teljesítmény szempontjából jobb, ha a swap-ot a szilárdtestalapú hardverre helyezzük.

Érdemes megjegyezni, hogy a swap _fájlok_ a swap _partíciók_ alternatívájaként használhatók. Ez leginkább a nagyon korlátozott adathordozó területtel rendelkező operációs rendszerek számára hasznos.

## Az fdisk használata HPPA-n

Használja a fdisk parancsot a szükséges partíciók létrehozásához:

`root #` `fdisk /dev/sda`

A HPPA számítógépek a PC szabványos DOS partíciós táblákat használják. Egy új DOS partíciós tábla létrehozásához nyomja meg az `o` billentyűgombot.

`Command (m for help):` `o`

```
Building a new DOS disklabel.

```

A PALO (a HPPA betöltő) működéséhez speciális partícióra van szükség. Legalább 16 MB méretű partíciót kell létrehozni az adathordozó elején. A partíció típusának _f0 (Linux/PA-RISC boot)_ típusúnak kell lennie. Az is lehetséges, hogy a PALO partíciót /boot könyvtárként használja.

**Important**

Ha ezt elfelejtik, és az telepítés külön PALO partíció nélkül folytatódik, akkor az operációs rendszer végül nem fog tudni újraindításkor elindulni ismét. Továbbá, ha az adathordozó nagyobb, mint 2 GiB, akkor ügyeljen arra, hogy a boot partíció az adathordozó első 2 GiB-jában legyen. A PALO nem tud olvasni egy kernelt a 2 GiB-os korlát után.

FILE **`/etc/fstab`** **Egyszerű alapértelmezett partíciós séma**

```
/dev/sda2    /boot   ext2    noauto,noatime   1 1
/dev/sda3    none    swap    sw               0 0
/dev/sda4    /       ext4    noatime          0 0

```

Az fdisk particionáló segédprogramban egy ilyen partíciós elrendezés így néz ki:

`Command (m for help):` `p`

```
Disk /dev/sda: 4294 MB, 4294816768 bytes
133 heads, 62 sectors/track, 1017 cylinders
Units = cylinders of 8246 * 512 = 4221952 bytes

   Device Boot      Start         End      Blocks   Id  System
/dev/sda1               1           8       32953   f0  Linux/PA-RISC boot
/dev/sda2               9          20       49476   83  Linux
/dev/sda3              21          70      206150   82  Linux swap
/dev/sda4              71        1017     3904481   83  Linux

```

## Fájlrendszerek létrehozása

**Warning**

SSD vagy NVMe adathordozó használatakor bölcs dolog ellenőrizni a firmware-frissítéseket. Különösen egyes Intel SSD adathordozók (600p és 6000p) firmware-frissítést igényelnek az XFS I/O használati minták által okozott [lehetséges adatsérülések miatt](https://bugzilla.redhat.com/show_bug.cgi?id=1402533). A probléma a firmware szintjén van, és nem az XFS fájlrendszer hibája. A smartctl segédprogram segíthet az adathordozzó eszköz modelljének és firmware-verziójának ellenőrzésében.

### Bevezetés

Most, hogy a partíciók elkészültek, ideje fájlrendszert helyezni rájuk. A következő részben a Linux által támogatott különféle fájlrendszereket ismertetjük. Azok az olvasók, akik már tudják, hogy melyik fájlrendszert fogják használni, folytathatják a [Fájlrendszer rárakása egy partícióra](/wiki/Handbook:HPPA/Installation/Disks/hu#Applying_a_filesystem_to_a_partition "Handbook:HPPA/Installation/Disks/hu") című bekezdéssel. A többi felhasználónak érdemes továbbolvasniuk, hogy megismerjék az alkalmazható fájlrendszereket...

### Fájlrendszerek

A Linux több tucat fájlrendszert támogat, bár ezek közül sokat csak meghatározott célokra érdemes telepíteni. Nem mindegyik fájlrendszer tekinthetők stabilnak az architektúrán. Javasoljuk, hogy tájékozódjon a fájlrendszerekről és azok támogatási állapotáról, még mielőtt egy kísérleti állapotban lévőt választana az Ön által fontosnak ítélt partíciókhoz. **Az XFS fájlrendszer univerzálisan ajánlott, mert minden platformra kiterjed.** Az alábbi egy nem teljes lista:

[XFS](/wiki/XFS/hu "XFS/hu")Fájlrendszer metaadat-naplózással, amely robusztus funkciókkal rendelkezik, és a méretezhetőségre van optimalizálva. Folyamatosan frissítik, hogy modern funkciókat is tartalmazzon. Az egyetlen hátrány, hogy az XFS-partíciók még nem zsugoríthatók, bár ezen dolgoznak. Az XFS különösen támogatja a reflinkeket és a Copy on Write (CoW) funkciót, ami különösen hasznos a Gentoo rendszereken a számos fordítás miatt, amit a felhasználók végeznek. Az XFS az ajánlott modern, minden célra használható, minden platformon elérhető fájlrendszer. Legalább 300 MB méretű partíciót igényel.[ext4](/wiki/Ext4/hu "Ext4/hu")Az Ext4 egy megbízható, általános célú, minden platformon használható fájlrendszer, bár hiányoznak belőle a modern funkciók, mint például a reflinkek.[VFAT](/wiki/VFAT "VFAT")Más néven FAT32, támogatott a Linux által, de nem támogatja a szabványos UNIX jogosultságbeállításokat. Főként más operációs rendszerekkel (például Microsoft Windows vagy Apple macOS) való együttműködésre/cserére használják, de szükségszerű bizonyos rendszerindító firmware-ekhez (például UEFI-hez) is. Az UEFI rendszerek felhasználóinak egy [EFI System Partition](/wiki/EFI_System_Partition/hu "EFI System Partition/hu") partíciót kell VFAT formátumban létrehozniuk a rendszerindításhoz.[btrfs](/wiki/Btrfs/hu "Btrfs/hu")Új generációs fájlrendszer. Fejlett funkciókat kínál, mint például pillanatképek készítése, ellenőrzőösszegek alapján történő öngyógyítás, átlátható tömörítés, alhálózatok és integrált RAID. Az 5.4.y előtti kernellel rendelkező verziók nem garantáltan biztonságosak a btrfs termelési környezetben való használatához, mivel a súlyos problémákra vonatkozó javítások csak az LTS kernelágak újabb kiadásaiban találhatók meg. A RAID 5/6 és a kvótacsoportok minden btrfs verziónál nem biztonságosak.[F2FS](/wiki/F2FS/hu "F2FS/hu")A Flash-Friendly File System-et eredetileg a Samsung hozta létre NAND flash memóriákhoz való használatra. Jó választás lehet, ha a Gentoo-t microSD kártyákra, USB meghajtókra vagy más flash-alapú tárolóeszközökre telepítik.[NTFS](/wiki/NTFS/hu "NTFS/hu")Ez a "New Technology" fájlrendszer a Microsoft Windows zászlóshajó fájlrendszere a Windows NT 3.1 óta. Hasonlóan a VFAT-hoz, nem tárol UNIX jogosultságbeállításokat vagy a BSD vagy Linux megfelelő működéséhez szükséges kiterjesztett attribútumokat, ezért a legtöbb esetben nem szabad gyökérfájlrendszerként használni. Csak és kizárólag interoperabilitásra vagy adatcserére használható a Microsoft Windows rendszerekkel (kiemelten a "csak" hangsúlyozásával).[ZFS](/wiki/ZFS "ZFS")**Important:** A ZFS tárak kizárólag az admincd és a LiveGUI ISO-ken hozhatók létre. További információért tekintse meg a [ZFS/rootfs](/wiki/ZFS/rootfs "ZFS/rootfs") oldalt.Következő generációs fájlrendszer, amelyet Matthew Ahrens és Jeff Bonwick hozott létre. Néhány kulcsfontosságú elképzelés alapján tervezték: a tárolás kezelése legyen egyszerű, a redundanciát a fájlrendszernek kell kezelnie, a fájlrendszereket soha ne kelljen javítás miatt offline állapotba helyezni, a legrosszabb forgatókönyvek automatizált szimulációja a kód kiadása előtt fontos, és az adatintegritás kiemelt jelentőségű.

A fájlrendszerekkel kapcsolatban bővebb információkat talál, ha elolvassa a közösség által karbantartott [Fájlrendszer](/wiki/Filesystem/hu "Filesystem/hu") nevű cikket.

### Fájlrendszer rárakása egy partícióra

**Note**

Kérjük, győződjön meg a számítógép újraindítása előtt, hogy az emerge segítségével valóban fel lett telepítve a kiválasztott fájlrendszerhez tartozó, (felhasználótérben működő, segédprogramokat tartalmazó) szoftvercsomag. A telepítési folyamat végén egy emlékeztető jelenik meg erre vonatkozólag.

Már előre elkészítve (minden lehetséges fájlrendszerhez) rendelkezésre állnak olyan felhasználói térben működő segédprogramok, amelyek segítségével egy partíción vagy egy köteten létre tudunk hozni fájlrendszert. Az egyes fájlrendszerekkel kapcsolatos további információkért kattintson a fájlrendszer nevére az alábbi táblázatban:

Filesystem
Létrehozási parancs
Elő környezeten belül?
Szoftvercsomag
[XFS](/wiki/XFS/hu "XFS/hu")mkfs.xfs Igen
[sys-fs/xfsprogs](https://packages.gentoo.org/packages/sys-fs/xfsprogs)[ext4](/wiki/Ext4/hu "Ext4/hu")mkfs.ext4 Igen
[sys-fs/e2fsprogs](https://packages.gentoo.org/packages/sys-fs/e2fsprogs)[VFAT](/wiki/VFAT "VFAT") (FAT32, ...)
mkfs.vfat Igen
[sys-fs/dosfstools](https://packages.gentoo.org/packages/sys-fs/dosfstools)[btrfs](/wiki/Btrfs/hu "Btrfs/hu")mkfs.btrfs Igen
[sys-fs/btrfs-progs](https://packages.gentoo.org/packages/sys-fs/btrfs-progs)[F2FS](/wiki/F2FS/hu "F2FS/hu")mkfs.f2fs Igen
[sys-fs/f2fs-tools](https://packages.gentoo.org/packages/sys-fs/f2fs-tools)[NTFS](/wiki/NTFS/hu "NTFS/hu")mkfs.ntfs Igen
[sys-fs/ntfs3g](https://packages.gentoo.org/packages/sys-fs/ntfs3g)[ZFS](/wiki/ZFS "ZFS")zpool create ... Nem
[sys-fs/zfs](https://packages.gentoo.org/packages/sys-fs/zfs)

\|}

**Important**

A kézikönyv az Ön fizikai adathordozóján új partíciók létrehozását javasolja a telepítési folyamat részeként. Fontos megjegyezni, hogy minden mkfs parancs futtatása törli a már meglévő partíciókon lévő (Önnek esetleg nagyon értékes) adatokat. Amennyiben szükséges, akkor Ön még az új fájlrendszerek létrehozása előtt győződjön meg arról, hogy a mostani fájlrendszereken lévő adatokról biztonsági másolat készült.

Például ahhoz, hogy a gyökérpartíció (tehát a root partíció) (/dev/sda4) fájlrendszertípusa xfs legyen, ahogy a partíciókészítés példa szerkezetében is szerepel, Önnek a következő parancsokat kell futtatnia:

`root #` `mkfs.xfs /dev/sda4`

#### Örökölt BIOS rendszerindító partíciónak a fájlrendszere

A régebbi, MBR/DOS adathordozó partíciós táblázattal ellátott BIOS-on keresztül induló rendszerek bármilyen, a rendszerbetöltő által támogatott fájlrendszert használhatnak.

Például XFS fájlrendszerrel történő formázáshoz futtassa a következő parancsot:

`root #` `mkfs.xfs /dev/sda2`

#### Kicsi ext4 partíciók

Ha Ön egy kicsi méretű partíción (kevesebb, mint 8 GiB) ext4 fájlrendszert szeretne használ, akkor a fájlrendszert a megfelelő beállításokkal kell létrehozni, hogy az elegendő inode-okat foglalhasson le. Ezt a `-T small` opcióval lehet megadni:

`root #` `mkfs.ext4 -T small /dev/<device>`

Ezzel megnégyszerezi az adott fájlrendszer inode-jainak a számát, mivel a "bytes-per-inode" 16 kB-onként 4 kB-ra csökken.

### A swap (lapozásra használt) partíció aktiválása

Az mkswap parancs szolgál a swap partíciók létrehozásához:

`root #` `mkswap /dev/sda3`

**Note**

Innentől folytatható az a rendszertelepítés, amely korábban el lett kezdve, de a telepítési folyamat nem let végig befejezve. Használja ezt a hivatkozást állandó hivatkozásként: [A telepítés folytatása itt kezdődik](/wiki/Handbook:HPPA/Installation/Disks/hu#Resumed_installations_start_here "Handbook:HPPA/Installation/Disks/hu").

A swap partíciót aktiválni is kell. Használja a swapon parancsot:

`root #` `swapon /dev/sda3`

Ez az 'aktiválás' azért szükséges most, mert a swap partíciót újonnan hozzuk létre a Live ISO telepítőkörnyezetben. A rendszer újraindítása után mindaddig, amíg a swap partíció megfelelően van definiálva az fstab fájlban vagy más csatolási mechanizmusban, a swap terület automatikusan fog aktiválódni.

## Gyökérpartíció (root partíció) felcsatolása

Előfordulhat, hogy bizonyos Live ISO telepítőkörnyezetekből hiányzik a javasolt csatolási pont a Gentoo gyökérpartíciójához (/mnt/gentoo), vagy hiányzik a particionálási szakaszban létrehozott további partíciók csatolási pontja:

`root #` `mkdir --parents /mnt/gentoo`

Az mkdir paranccsal folytassa az előző lépések során létrehozott (egyéni) partíció(k)hoz szükséges további felcsatolási pontok létrehozását.

A felcsatolási pontok létrehozását követően ideje elérhetővé tenni a partíciókat a mount paranccsal.

Csatolja fel a gyökérpartíciót (a root partíciót):

`root #` `mount /dev/sda4 /mnt/gentoo`

Szükség szerint folytassa a további (egyéni) partíciók felcsatolását a fájlrendszerbe a mount paranccsal.

**Note**

Ha a /tmp/ könyvtárnak külön partíción kell lennie, akkor a felcsatolás után mindenképpen módosítsa a hozzá tartozó jogosultságokat:

`root #` `chmod 1777 /mnt/gentoo/tmp`

Ugyanez érvényes a /var/tmp könyvtárra is.

Később az utasításokban a proc fájlrendszer (a kernellel kapcsolatban álló virtuális interfész) és a többi kernel pszeudofájlrendszer lesz felcsatolva. Először viszont még a [Gentoo-stage fájlt](/wiki/Handbook:HPPA/Installation/Stage/hu "Handbook:HPPA/Installation/Stage/hu") ki kell csomagolnunk.

[← Hálózat beállítása](/wiki/Handbook:HPPA/Installation/Networking/hu "Previous part") [Kezdőlap](/wiki/Handbook:HPPA/hu "Handbook:HPPA/hu") [Stage fájl telepítése →](/wiki/Handbook:HPPA/Installation/Stage/hu "Next part")