# Disks

Other languages:

- [Deutsch](/wiki/Handbook:PPC/Installation/Disks/de "Handbuch:PPC/Installation/Festplatten (100% translated)")
- [English](/wiki/Handbook:PPC/Installation/Disks "Handbook:PPC/Installation/Disks (100% translated)")
- [español](/wiki/Handbook:PPC/Installation/Disks/es "Manual de Gentoo: PPC/Instalación/Discos (100% translated)")
- [français](/wiki/Handbook:PPC/Installation/Disks/fr "Handbook:PPC/Installation/Disks/fr (100% translated)")
- [italiano](/wiki/Handbook:PPC/Installation/Disks/it "Handbook:PPC/Installation/Disks/it (100% translated)")
- magyar
- [polski](/wiki/Handbook:PPC/Installation/Disks/pl "Handbook:PPC/Installation/Disks (100% translated)")
- [português do Brasil](/wiki/Handbook:PPC/Installation/Disks/pt-br "Handbook:PPC/Installation/Disks/pt-br (50% translated)")
- [русский](/wiki/Handbook:PPC/Installation/Disks/ru "Handbook:PPC/Installation/Disks (100% translated)")
- [தமிழ்](/wiki/Handbook:PPC/Installation/Disks/ta "கையேடு:PPC/நிறுவல்/வட்டுக்கள் (100% translated)")
- [中文（中国大陆）‎](/wiki/Handbook:PPC/Installation/Disks/zh-cn "手册：PPC/安装/准备磁盘 (100% translated)")
- [日本語](/wiki/Handbook:PPC/Installation/Disks/ja "ハンドブック:PPC/インストール/ディスク (100% translated)")
- [한국어](/wiki/Handbook:PPC/Installation/Disks/ko "Handbook:PPC/Installation/Disks/ko (100% translated)")

[PPC kézikönyv](/wiki/Handbook:PPC/hu "Handbook:PPC/hu")[A Gentoo Linux telepítése](/wiki/Handbook:PPC/Full/Installation/hu "Handbook:PPC/Full/Installation/hu")[A telepítésről](/wiki/Handbook:PPC/Installation/About/hu "Handbook:PPC/Installation/About/hu")[Telepítőképfájl kiválasztása](/wiki/Handbook:PPC/Installation/Media/hu "Handbook:PPC/Installation/Media/hu")[Hálózat beállítása](/wiki/Handbook:PPC/Installation/Networking/hu "Handbook:PPC/Installation/Networking/hu")Adathordozók előkészítése[Stage fájl](/wiki/Handbook:PPC/Installation/Stage/hu "Handbook:PPC/Installation/Stage/hu")[Alaprendszer telepítése](/wiki/Handbook:PPC/Installation/Base/hu "Handbook:PPC/Installation/Base/hu")[Kernel beállítása](/wiki/Handbook:PPC/Installation/Kernel/hu "Handbook:PPC/Installation/Kernel/hu")[Rendszer beállítása](/wiki/Handbook:PPC/Installation/System/hu "Handbook:PPC/Installation/System/hu")[Eszközök telepítése](/wiki/Handbook:PPC/Installation/Tools/hu "Handbook:PPC/Installation/Tools/hu")[Bootloader beállítása](/wiki/Handbook:PPC/Installation/Bootloader/hu "Handbook:PPC/Installation/Bootloader/hu")[Telepítés véglegesítése](/wiki/Handbook:PPC/Installation/Finalizing/hu "Handbook:PPC/Installation/Finalizing/hu")[Munka a Gentoo rendszerrel](/wiki/Handbook:PPC/Full/Working/hu "Handbook:PPC/Full/Working/hu")[Portage bemutatása](/wiki/Handbook:PPC/Working/Portage/hu "Handbook:PPC/Working/Portage/hu")[USE jelölőzászlók](/wiki/Handbook:PPC/Working/USE/hu "Handbook:PPC/Working/USE/hu")[Portage jellemzői](/wiki/Handbook:PPC/Working/Features/hu "Handbook:PPC/Working/Features/hu")[Init-szkript rendszer](/wiki/Handbook:PPC/Working/Initscripts/hu "Handbook:PPC/Working/Initscripts/hu")[Környezeti változók](/wiki/Handbook:PPC/Working/EnvVar/hu "Handbook:PPC/Working/EnvVar/hu")[Munka a Portage szoftvercsomag-kezelővel](/wiki/Handbook:PPC/Full/Portage/hu "Handbook:PPC/Full/Portage/hu")[Fájlok és könyvtárak](/wiki/Handbook:PPC/Portage/Files/hu "Handbook:PPC/Portage/Files/hu")[Változók](/wiki/Handbook:PPC/Portage/Variables/hu "Handbook:PPC/Portage/Variables/hu")[Szoftverágak keverése](/wiki/Handbook:PPC/Portage/Branches/hu "Handbook:PPC/Portage/Branches/hu")[További eszközök](/wiki/Handbook:PPC/Portage/Tools/hu "Handbook:PPC/Portage/Tools/hu")[Egyéni szoftvercsomag-tárolók](/wiki/Handbook:PPC/Portage/CustomTree/hu "Handbook:PPC/Portage/CustomTree/hu")[Fejlett funkciók](/wiki/Handbook:PPC/Portage/Advanced/hu "Handbook:PPC/Portage/Advanced/hu")[Hálózat beállítása OpenRC init-rendszeren](/wiki/Handbook:PPC/Full/Networking/hu "Handbook:PPC/Full/Networking/hu")[Munka elkezdése](/wiki/Handbook:PPC/Networking/Introduction/hu "Handbook:PPC/Networking/Introduction/hu")[Fejlett beállítások](/wiki/Handbook:PPC/Networking/Advanced/hu "Handbook:PPC/Networking/Advanced/hu")[Moduláris hálózat](/wiki/Handbook:PPC/Networking/Modular/hu "Handbook:PPC/Networking/Modular/hu")[Vezeték nélküli (Wi-Fi)](/wiki/Handbook:PPC/Networking/Wireless/hu "Handbook:PPC/Networking/Wireless/hu")[Funkcionalitás hozzáadása](/wiki/Handbook:PPC/Networking/Extending/hu "Handbook:PPC/Networking/Extending/hu")[Dinamikus menedzsment](/wiki/Handbook:PPC/Networking/Dynamic/hu "Handbook:PPC/Networking/Dynamic/hu")

## Contents

- [1Bevezetés a blokktípusú eszközökbe](#Bevezet.C3.A9s_a_blokkt.C3.ADpus.C3.BA_eszk.C3.B6z.C3.B6kbe)
  - [1.1Blokkeszközök](#Blokkeszk.C3.B6z.C3.B6k)
  - [1.2Partíciók](#Part.C3.ADci.C3.B3k)
- [2Partíciós séma megtervezése](#Part.C3.ADci.C3.B3s_s.C3.A9ma_megtervez.C3.A9se)
  - [2.1Hány partíció és mekkora méretű?](#H.C3.A1ny_part.C3.ADci.C3.B3_.C3.A9s_mekkora_m.C3.A9ret.C5.B1.3F)
  - [2.2Mi a helyzet a swap területtel?](#Mi_a_helyzet_a_swap_ter.C3.BClettel.3F)
  - [2.3Apple New World](#Apple_New_World)
  - [2.4Apple Old World](#Apple_Old_World)
  - [2.5Pegasos](#Pegasos)
  - [2.6IBM PReP (RS/6000)](#IBM_PReP_.28RS.2F6000.29)
- [3A mac-fdisk (Apple) használata](#A_mac-fdisk_.28Apple.29_haszn.C3.A1lata)
- [4Parted segédszoftvert használata (Pegasos és RS/6000)](#Parted_seg.C3.A9dszoftvert_haszn.C3.A1lata_.28Pegasos_.C3.A9s_RS.2F6000.29)
- [5Fájlrendszerek létrehozása](#F.C3.A1jlrendszerek_l.C3.A9trehoz.C3.A1sa)
  - [5.1Bevezetés](#Bevezet.C3.A9s)
  - [5.2Fájlrendszerek](#F.C3.A1jlrendszerek)
  - [5.3Fájlrendszer rárakása egy partícióra](#F.C3.A1jlrendszer_r.C3.A1rak.C3.A1sa_egy_part.C3.ADci.C3.B3ra)
    - [5.3.1Örökölt BIOS rendszerindító partíciónak a fájlrendszere](#.C3.96r.C3.B6k.C3.B6lt_BIOS_rendszerind.C3.ADt.C3.B3_part.C3.ADci.C3.B3nak_a_f.C3.A1jlrendszere)
    - [5.3.2Kicsi ext4 partíciók](#Kicsi_ext4_part.C3.ADci.C3.B3k)
  - [5.4A swap (lapozásra használt) partíció aktiválása](#A_swap_.28lapoz.C3.A1sra_haszn.C3.A1lt.29_part.C3.ADci.C3.B3_aktiv.C3.A1l.C3.A1sa)
- [6Gyökérpartíció (root partíció) felcsatolása](#Gy.C3.B6k.C3.A9rpart.C3.ADci.C3.B3_.28root_part.C3.ADci.C3.B3.29_felcsatol.C3.A1sa)

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

### Partíciók

Bár elméletileg lehetséges, hogy egy teljes adathordozót felhasználjunk egy Linux operációs rendszer számára, mégis a gyakorlatban ez szinte soha nem történik meg. Ehelyett a teljes adathordozó blokkeszközöket kisebb, könnyebben kezelhető blokkeszközökre osztják fel. A legtöbb rendszeren ezeket partícióknak nevezik.

**Note**

A telepítési utasítás hátralévő részében a Pegasos például szolgáló partíció-elrendezést fogjuk használni. Kérjük, Önt hogy igazítsa a példád a saját preferenciáihoz.

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

### Apple New World

Az Apple New World számítógépek beállítása viszonylag egyszerű. Az első partíció mindig egy Apple Partition Map (APM). Ez a partíció az adathordozó elrendezését tartja nyilván. Ezt a partíciót nem lehet eltávolítani. A következő partíciónak mindig egy bootstrap partíciónak kell lennie. Ez a partíció egy kis méretű (800 KiB) HFS fájlrendszert tartalmaz, amely magába foglalja a Yaboot bootloader egy példányát és annak beállításfájlját. Ez a partíció nem ugyanaz, mint egy /boot partíció, amely más architektúrákon található meg. A boot partíció után a szokásos Linux fájlrendszerek helyezkednek el, az alábbi séma szerint. A swap partíció ideiglenes tárolóhely a rendszer számára, amikor a fizikai memória elfogy. A root partíció tartalmazza a Gentoo telepítését szolgáló fájlrendszert. A dual bootolás érdekében az OSX partíció bárhol elhelyezhető a bootstrap partíció után, hogy biztosítsa a Yaboot indulását.

**Note**

Lehetnek "Disk Driver" partíciók az adathordozón, mint például Apple\_Driver63, Apple\_Driver\_ATA, Apple\_FWDriver, Apple\_Driver\_IOKit és Apple\_Patches. Ezeket a MacOS indításához használják, így ha nincs szükség rájuk, akkor eltávolíthatók a mac-fdisk `i` opciójával az adathordozó inicializálásakor. Legyen óvatos, ez teljesen törölni fogja az adathordozót! Ha kétségei vannak, akkor ne távolítsa el ezeket.

**Note**

Ha az adathordozó az Apple Disk Utility segítségével van particionálva, akkor előfordulhatnak 128 MiB méretű helyek a partíciók között, amelyeket az Apple "jövőbeni felhasználásra" tart fenn. Ezek mind biztonságosan eltávolíthatók.

Partíció
Méret
Fájlrendszer
Leírás
/dev/sda132KiB
Semmi.
Apple Partition Map (APM).
/dev/sda2800KiB
HFS
Apple bootstrap.
/dev/sda3512 MiB
swap
Linux swap (type 0x82).
/dev/sda4Az adathordozó fennmaradó része.
ext4, xfs, etc.
Linux root.

### Apple Old World

Az Apple Old World számítógépek beállítása valamivel bonyolultabb. Az első partíció mindig egy Apple Partition Map (APM). Ez a partíció az adathordozó elrendezését tartja nyilván. Ezt a partíciót nem lehet eltávolítani. BootX használatakor az alábbi beállítás azt feltételezi, hogy a MacOS egy külön adathordozón van telepítve. Ha ez nem így van, akkor további partíciók lesznek az "Apple Disk Drivers" számára, mint például Apple\_Driver63, Apple\_Driver\_ATA, Apple\_FWDriver, Apple\_Driver\_IOKit, Apple\_Patches és a MacOS telepítés. Quik használatakor szükséges létrehozni egy boot partíciót a kernel számára, ellentétben más Apple boot módszerekkel. A boot partíció után a szokásos Linux fájlrendszerek helyezkednek el, az alábbi séma szerint. A swap partíció ideiglenes tárolóhely a rendszer számára, amikor a fizikai memória elfogy. A root partíció tartalmazza a Gentoo operációs rendszert telepítését szolgáló fájlrendszert.

**Note**

Old World számítógép használatakor elengedhetetlen, hogy a MacOS elérhető legyen. Az itt bemutatott elrendezés azt feltételezi, hogy a MacOS egy külön meghajtóra van telepítve.

Példa partícióelrendezés egy Old World számítógéphez:

Partíció
Méret
Fájlrendszer
Leírás
/dev/sda132KiB
None.
Apple Partition Map (APM).
/dev/sda232MiB
ext2
Quik Boot Partition (kizárólag quik).
/dev/sda3512MiB
swap
Linux swap (type 0x82).
/dev/sda4Az adathordozó többi fennmaradó része.
ext4, xfs, etc.
Linux root.

### Pegasos

A Pegasos partícióelrendezés meglehetősen egyszerű az Apple elrendezésekhez képest. Az első partíció egy boot partíció, amely a betöltendő kerneleket és egy Open Firmware szkriptet tartalmaz, amely egy menüt jelenít meg indításkor. A boot partíció után a szokásos Linux fájlrendszerek helyezkednek el az alábbi séma szerint. A swap partíció ideiglenes tárolóhely a rendszer számára, amikor a fizikai memória elfogy. A root partíció tartalmazza a Gentoo telepítését szolgáló fájlrendszert.

Példa partícióelrendezés a Pegasos rendszerekhez:

Partíció
Méret
Fájlrendszer
Leírás
/dev/sda132MiB
affs1 or ext2
Boot partition.
/dev/sda2512MiB
swap
Linux swap (type 0x82).
/dev/sda3Az adathordozó többi fennmaradó része.
ext4, xfs, etc.
Linux root.

### IBM PReP (RS/6000)

Az IBM PowerPC Reference Platform (PReP) egy kis PReP boot partíciót igényel az adathordozó első partícióján, amelyet a swap és root partíciók követnek.

Példa partícióelrendezés az IBM PReP számára:

Partíció
Méret
Fájlrendszer
Leírás
/dev/sda1800KiB
None
PReP boot partition (type 0x41).
/dev/sda2512MiB
swap
Linux swap (type 0x82).
/dev/sda3Az adathordozó többi fennmaradó része.
ext4, xfs, etc.
Linux root (type 0x83).

**Warning**

A parted segédszoftver képes partíciók átméretezésére, beleértve a HFS+ -t is. Sajnos előfordulhatnak problémák a HFS+ naplózott fájlrendszerek átméretezésével, ezért a legjobb eredmények érdekében kapcsolja ki a naplózást Mac OS X-ben az átméretezés előtt. Ne feledje, hogy bármely átméretezési művelet veszélyes, ezért csak saját felelősségére végezze! Mindig legyen biztonsági másolata az összes adatnak az átméretezés előtt!

## A mac-fdisk (Apple) használata

Ezen a ponton hozza létre a partíciókat a mac-fdisk segítségével:

`root #` `mac-fdisk /dev/sda`

Ha az Apple Disk Utility-t korábban használták hely biztosítására Linux számára, akkor először törölje azokat a partíciókat, amelyek esetleg korábban létrejöttek, hogy helyet teremtsen az új telepítéshez. Használja a `d` opciót a mac-fdisk-ben ezeknek a partícióknak a törléséhez. Meg fogja kérdezni a törlendő partíció számát. Általában az első partíció a NewWorld számítógépeken (Apple\_partition\_map) nem törölhető. Ha tiszta adathordozóval szeretné kezdeni, akkor egyszerűen inicializálja az adathordozót a `i` megnyomásával. Ez teljesen törli az adathordozót, ezért óvatosan használja.

Másodszor, hozzon létre egy Apple\_Bootstrap partíciót a `b` használatával. Ez meg fogja kérdezni, hogy melyik blokktól kezdje. Adja meg az első szabad partíció számát, majd nyomja le a `p` billentyűgombot. Például ez _2p_ lesz.

**Note**

Ez a partíció nem egy /boot partíció. Egyáltalán nem használja a Linux, ezért nincs szükség fájlrendszer létrehozására rajta, és soha nem szabad felcsatolni. Az Apple felhasználóknak nincs szükségük egy külön partícióra a /boot számára.

Most hozzon létre egy swap partíciót a `c` megnyomásával. A mac-fdisk ismét meg fogja kérdezni, hogy melyik blokktól kezdje ezt a partíciót. Mivel a 2-t használtuk az Apple\_Bootstrap partíció létrehozásához, ezért most adja meg a _3p_ értéket. Amikor a méretet kérdezi, akkor írja be, hogy 512M (vagy amekkora szükséges – a minimum ajánlott méret 512MiB, de általánosan elfogadott a fizikai memória kétszerese). Amikor a név megadását kéri, írja be, hogy _swap_.

A root partíció létrehozásához nyomja meg a `c` billentyűgombot, majd a _4p_ értéket írja be, hogy kiválassza, melyik blokktól kezdődjön a root partíció. Amikor a méretet kérdezi, írja be ismét a _4p_ értéket. A mac-fdisk ezt úgy értelmezi, hogy "Használja az összes rendelkezésre álló helyet". Amikor a név megadását kéri, írja be, hogy _root_.

A befejezéshez írja a partíciókat az adathordozóra a `w` opció használatával, majd lépjen ki a mac-fdisk-ből a `q` megnyomásával.

**Note**

Annak érdekében, hogy minden rendben legyen, futtassa a mac-fdisk -l parancsot, és ellenőrizze, hogy az összes partíció ott van-e. Ha az előzőleg létrehozott partíciók nem jelennek meg, vagy a végrehajtott módosítások nem tükröződnek a kimenetben, akkor inicializálja újra a partíciókat a `i` billentyűgomb megnyomásával a mac-fdisk szoftverben. Vegye figyelembe, hogy ez újra létre fogja hozni a partíciótérképet, és ezzel eltávolítja az összes már meglévő partíciót.

## Parted segédszoftvert használata (Pegasos és RS/6000)

parted, a partíciószerkesztő, most már képes kezelni a Mac OS és Mac OS X által használt HFS+ partíciókat. Ezzel az eszközzel lehetséges a Mac partíciók átméretezése és hely létrehozása a Linux partíciók számára. Ennek ellenére az alábbi példa kizárólag Pegasos számítógépek particionálását írja le.

Kezdésképpen indítsuk el a parted szoftvert:

`root #` `parted /dev/sda`

Ha az adathordozó nincs particionálva, akkor futtassa a mklabel amiga parancsot, hogy új adathordozó táblázatot (lemezcímkét) hozzon létre az adathordozón.

A parted bármelyik pontján be lehet írni a print parancsot, hogy megjelenítse az aktuális partíciós táblázatot. A parted megszakításához nyomja meg a `Ctrl` \+ `c` billentyűgomb kombinációt.

Ha a Linux mellett a rendszernek MorphOS-t is telepíteni kell, akkor hozzon létre egy affs1 fájlrendszert az adathordozó elején. A 32 MiB bőven elegendő a MorphOS kernel tárolására. Ha Pegasos I -et használ, vagy ha Linux egyéb fájlrendszert használ, mint az ext2 vagy ext3, akkor szükséges a Linux kernel tárolása ezen a partíción (a Pegasos II csak ext2/ext3 vagy affs1 partíciókból tud bootolni). A partíció létrehozásához futtassa a `mkpart primary affs1 START END` parancsot, ahol a START és END helyére a megabyte tartományt kell beírni (például 0 32), amely egy 32 MiB-os partíciót hoz létre 0 MiB-tól 32 MiB-ig. Ha ext2 vagy ext3 partíciót szeretne létrehozni helyette, cserélje ki az affs1-et ext2-re vagy ext3-ra a mkpart parancsban.

Hozzon létre két partíciót a Linux számára: Egy root fájlrendszer partíciót és egy swap partíciót. Futtassa a `mkpart primary START END` parancsot minden partíció létrehozásához, a START és END értékeket pedig a kívánt megabyte határokkal helyettesítve.

Általánosan ajánlott olyan swap partíciót létrehozni, amely kétszer akkora, mint a számítógép RAM memóriája, de legalább 512MiB javasolt. A swap partíció létrehozásához futtassa a `mkpart primary linux-swap START END` parancsot, ahol a START és END ismét a partíció határait jelöli.

Ha végzett a parted használatával, akkor egyszerűen írja be a `quit` parancsot.

## Fájlrendszerek létrehozása

**Warning**

SSD vagy NVMe adathordozó használatakor bölcs dolog ellenőrizni a firmware-frissítéseket. Különösen egyes Intel SSD adathordozók (600p és 6000p) firmware-frissítést igényelnek az XFS I/O használati minták által okozott [lehetséges adatsérülések miatt](https://bugzilla.redhat.com/show_bug.cgi?id=1402533). A probléma a firmware szintjén van, és nem az XFS fájlrendszer hibája. A smartctl segédprogram segíthet az adathordozzó eszköz modelljének és firmware-verziójának ellenőrzésében.

### Bevezetés

Most, hogy a partíciók elkészültek, ideje fájlrendszert helyezni rájuk. A következő részben a Linux által támogatott különféle fájlrendszereket ismertetjük. Azok az olvasók, akik már tudják, hogy melyik fájlrendszert fogják használni, folytathatják a [Fájlrendszer rárakása egy partícióra](/wiki/Handbook:PPC/Installation/Disks/hu#Applying_a_filesystem_to_a_partition "Handbook:PPC/Installation/Disks/hu") című bekezdéssel. A többi felhasználónak érdemes továbbolvasniuk, hogy megismerjék az alkalmazható fájlrendszereket...

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

Például ahhoz, hogy a gyökérpartíció (tehát a root partíció) (/dev/sda3) fájlrendszertípusa xfs legyen, ahogy a partíciókészítés példa szerkezetében is szerepel, Önnek a következő parancsokat kell futtatnia:

`root #` `mkfs.xfs /dev/sda3`

#### Örökölt BIOS rendszerindító partíciónak a fájlrendszere

A régebbi, MBR/DOS adathordozó partíciós táblázattal ellátott BIOS-on keresztül induló rendszerek bármilyen, a rendszerbetöltő által támogatott fájlrendszert használhatnak.

Például XFS fájlrendszerrel történő formázáshoz futtassa a következő parancsot:

`root #` `mkfs.xfs /dev/sda1`

#### Kicsi ext4 partíciók

Ha Ön egy kicsi méretű partíción (kevesebb, mint 8 GiB) ext4 fájlrendszert szeretne használ, akkor a fájlrendszert a megfelelő beállításokkal kell létrehozni, hogy az elegendő inode-okat foglalhasson le. Ezt a `-T small` opcióval lehet megadni:

`root #` `mkfs.ext4 -T small /dev/<device>`

Ezzel megnégyszerezi az adott fájlrendszer inode-jainak a számát, mivel a "bytes-per-inode" 16 kB-onként 4 kB-ra csökken.

### A swap (lapozásra használt) partíció aktiválása

Az mkswap parancs szolgál a swap partíciók létrehozásához:

`root #` `mkswap /dev/sda2`

**Note**

Innentől folytatható az a rendszertelepítés, amely korábban el lett kezdve, de a telepítési folyamat nem let végig befejezve. Használja ezt a hivatkozást állandó hivatkozásként: [A telepítés folytatása itt kezdődik](/wiki/Handbook:PPC/Installation/Disks/hu#Resumed_installations_start_here "Handbook:PPC/Installation/Disks/hu").

A swap partíciót aktiválni is kell. Használja a swapon parancsot:

`root #` `swapon /dev/sda2`

Ez az 'aktiválás' azért szükséges most, mert a swap partíciót újonnan hozzuk létre a Live ISO telepítőkörnyezetben. A rendszer újraindítása után mindaddig, amíg a swap partíció megfelelően van definiálva az fstab fájlban vagy más csatolási mechanizmusban, a swap terület automatikusan fog aktiválódni.

## Gyökérpartíció (root partíció) felcsatolása

Előfordulhat, hogy bizonyos Live ISO telepítőkörnyezetekből hiányzik a javasolt csatolási pont a Gentoo gyökérpartíciójához (/mnt/gentoo), vagy hiányzik a particionálási szakaszban létrehozott további partíciók csatolási pontja:

`root #` `mkdir --parents /mnt/gentoo`

Az mkdir paranccsal folytassa az előző lépések során létrehozott (egyéni) partíció(k)hoz szükséges további felcsatolási pontok létrehozását.

A felcsatolási pontok létrehozását követően ideje elérhetővé tenni a partíciókat a mount paranccsal.

Csatolja fel a gyökérpartíciót (a root partíciót):

`root #` `mount /dev/sda3 /mnt/gentoo`

Szükség szerint folytassa a további (egyéni) partíciók felcsatolását a fájlrendszerbe a mount paranccsal.

**Note**

Ha a /tmp/ könyvtárnak külön partíción kell lennie, akkor a felcsatolás után mindenképpen módosítsa a hozzá tartozó jogosultságokat:

`root #` `chmod 1777 /mnt/gentoo/tmp`

Ugyanez érvényes a /var/tmp könyvtárra is.

Később az utasításokban a proc fájlrendszer (a kernellel kapcsolatban álló virtuális interfész) és a többi kernel pszeudofájlrendszer lesz felcsatolva. Először viszont még a [Gentoo-stage fájlt](/wiki/Handbook:PPC/Installation/Stage/hu "Handbook:PPC/Installation/Stage/hu") ki kell csomagolnunk.

[← Hálózat beállítása](/wiki/Handbook:PPC/Installation/Networking/hu "Previous part") [Kezdőlap](/wiki/Handbook:PPC/hu "Handbook:PPC/hu") [Stage fájl telepítése →](/wiki/Handbook:PPC/Installation/Stage/hu "Next part")