# Disks

Other languages:

- [Deutsch](/wiki/Handbook:Alpha/Installation/Disks/de "Handbuch:Alpha/Installation/Festplatten (100% translated)")
- [English](/wiki/Handbook:Alpha/Installation/Disks "Handbook:Alpha/Installation/Disks (100% translated)")
- [Türkçe](/wiki/Handbook:Alpha/Installation/Disks/tr "Handbook:Alpha/Installation/Disks/tr (0% translated)")
- [español](/wiki/Handbook:Alpha/Installation/Disks/es "Manual de Gentoo: Alpha/Instalación/Discos (100% translated)")
- [français](/wiki/Handbook:Alpha/Installation/Disks/fr "Manuel:Alpha/Installation/Disques (100% translated)")
- [italiano](/wiki/Handbook:Alpha/Installation/Disks/it "Handbook:Alpha/Installation/Disks/it (100% translated)")
- magyar
- [polski](/wiki/Handbook:Alpha/Installation/Disks/pl "Handbook:Alpha/Installation/Disks (100% translated)")
- [português do Brasil](/wiki/Handbook:Alpha/Installation/Disks/pt-br "Manual:Alpha/Instalação/Discos (100% translated)")
- [čeština](/wiki/Handbook:Alpha/Installation/Disks/cs "Handbook:Alpha/Installation/Disks/cs (50% translated)")
- [русский](/wiki/Handbook:Alpha/Installation/Disks/ru "Handbook:Alpha/Installation/Disks (100% translated)")
- [தமிழ்](/wiki/Handbook:Alpha/Installation/Disks/ta "கையேடு:Alpha/நிறுவல்/வட்டுக்கள் (100% translated)")
- [中文（中国大陆）‎](/wiki/Handbook:Alpha/Installation/Disks/zh-cn "手册：Alpha/安装/准备磁盘 (100% translated)")
- [日本語](/wiki/Handbook:Alpha/Installation/Disks/ja "ハンドブック:Alpha/インストール/ディスク (100% translated)")
- [한국어](/wiki/Handbook:Alpha/Installation/Disks/ko "Handbook:Alpha/Installation/Disks/ko (100% translated)")

[Alpha kézikönyv](/wiki/Handbook:Alpha/hu "Handbook:Alpha/hu")[A Gentoo Linux telepítése](/wiki/Handbook:Alpha/Full/Installation/hu "Handbook:Alpha/Full/Installation/hu")[A telepítésről](/wiki/Handbook:Alpha/Installation/About/hu "Handbook:Alpha/Installation/About/hu")[Telepítőképfájl kiválasztása](/wiki/Handbook:Alpha/Installation/Media/hu "Handbook:Alpha/Installation/Media/hu")[Hálózat beállítása](/wiki/Handbook:Alpha/Installation/Networking/hu "Handbook:Alpha/Installation/Networking/hu")Adathordozók előkészítése[Stage fájl](/wiki/Handbook:Alpha/Installation/Stage/hu "Handbook:Alpha/Installation/Stage/hu")[Alaprendszer telepítése](/wiki/Handbook:Alpha/Installation/Base/hu "Handbook:Alpha/Installation/Base/hu")[Kernel beállítása](/wiki/Handbook:Alpha/Installation/Kernel/hu "Handbook:Alpha/Installation/Kernel/hu")[Rendszer beállítása](/wiki/Handbook:Alpha/Installation/System/hu "Handbook:Alpha/Installation/System/hu")[Eszközök telepítése](/wiki/Handbook:Alpha/Installation/Tools/hu "Handbook:Alpha/Installation/Tools/hu")[Bootloader beállítása](/wiki/Handbook:Alpha/Installation/Bootloader/hu "Handbook:Alpha/Installation/Bootloader/hu")[Telepítés véglegesítése](/wiki/Handbook:Alpha/Installation/Finalizing/hu "Handbook:Alpha/Installation/Finalizing/hu")[Munka a Gentoo rendszerrel](/wiki/Handbook:Alpha/Full/Working/hu "Handbook:Alpha/Full/Working/hu")[Portage bemutatása](/wiki/Handbook:Alpha/Working/Portage/hu "Handbook:Alpha/Working/Portage/hu")[USE jelölőzászlók](/wiki/Handbook:Alpha/Working/USE/hu "Handbook:Alpha/Working/USE/hu")[Portage jellemzői](/wiki/Handbook:Alpha/Working/Features/hu "Handbook:Alpha/Working/Features/hu")[Init-szkript rendszer](/wiki/Handbook:Alpha/Working/Initscripts/hu "Handbook:Alpha/Working/Initscripts/hu")[Környezeti változók](/wiki/Handbook:Alpha/Working/EnvVar/hu "Handbook:Alpha/Working/EnvVar/hu")[Munka a Portage szoftvercsomag-kezelővel](/wiki/Handbook:Alpha/Full/Portage/hu "Handbook:Alpha/Full/Portage/hu")[Fájlok és könyvtárak](/wiki/Handbook:Alpha/Portage/Files/hu "Handbook:Alpha/Portage/Files/hu")[Változók](/wiki/Handbook:Alpha/Portage/Variables/hu "Handbook:Alpha/Portage/Variables/hu")[Szoftverágak keverése](/wiki/Handbook:Alpha/Portage/Branches/hu "Handbook:Alpha/Portage/Branches/hu")[További eszközök](/wiki/Handbook:Alpha/Portage/Tools/hu "Handbook:Alpha/Portage/Tools/hu")[Egyéni szoftvercsomag-tárolók](/wiki/Handbook:Alpha/Portage/CustomTree/hu "Handbook:Alpha/Portage/CustomTree/hu")[Fejlett funkciók](/wiki/Handbook:Alpha/Portage/Advanced/hu "Handbook:Alpha/Portage/Advanced/hu")[Hálózat beállítása OpenRC init-rendszeren](/wiki/Handbook:Alpha/Full/Networking/hu "Handbook:Alpha/Full/Networking/hu")[Munka elkezdése](/wiki/Handbook:Alpha/Networking/Introduction/hu "Handbook:Alpha/Networking/Introduction/hu")[Fejlett beállítások](/wiki/Handbook:Alpha/Networking/Advanced/hu "Handbook:Alpha/Networking/Advanced/hu")[Moduláris hálózat](/wiki/Handbook:Alpha/Networking/Modular/hu "Handbook:Alpha/Networking/Modular/hu")[Vezeték nélküli (Wi-Fi)](/wiki/Handbook:Alpha/Networking/Wireless/hu "Handbook:Alpha/Networking/Wireless/hu")[Funkcionalitás hozzáadása](/wiki/Handbook:Alpha/Networking/Extending/hu "Handbook:Alpha/Networking/Extending/hu")[Dinamikus menedzsment](/wiki/Handbook:Alpha/Networking/Dynamic/hu "Handbook:Alpha/Networking/Dynamic/hu")

## Contents

- [1Bevezetés a blokktípusú eszközökbe](#Bevezet.C3.A9s_a_blokkt.C3.ADpus.C3.BA_eszk.C3.B6z.C3.B6kbe)
  - [1.1Blokkeszközök](#Blokkeszk.C3.B6z.C3.B6k)
  - [1.2Szeletek](#Szeletek)
- [2Partíciós séma megtervezése](#Part.C3.ADci.C3.B3s_s.C3.A9ma_megtervez.C3.A9se)
  - [2.1Hány partíció és mekkora méretű?](#H.C3.A1ny_part.C3.ADci.C3.B3_.C3.A9s_mekkora_m.C3.A9ret.C5.B1.3F)
  - [2.2Mi a helyzet a swap területtel?](#Mi_a_helyzet_a_swap_ter.C3.BClettel.3F)
- [3Az fdisk használata az fizikai adathordozó particionálásához (csak SRM esetében)](#Az_fdisk_haszn.C3.A1lata_az_fizikai_adathordoz.C3.B3_particion.C3.A1l.C3.A1s.C3.A1hoz_.28csak_SRM_eset.C3.A9ben.29)
  - [3.1Elérhető fizikai adathordozók azonosítása](#El.C3.A9rhet.C5.91_fizikai_adathordoz.C3.B3k_azonos.C3.ADt.C3.A1sa)
  - [3.2BSD fizikai adathordozó címke létrehozása](#BSD_fizikai_adathordoz.C3.B3_c.C3.ADmke_l.C3.A9trehoz.C3.A1sa)
  - [3.3Összes szelet törlése](#.C3.96sszes_szelet_t.C3.B6rl.C3.A9se)
  - [3.4Swap szelet létrehozása](#Swap_szelet_l.C3.A9trehoz.C3.A1sa)
  - [3.5Boot szelet létrehozása](#Boot_szelet_l.C3.A9trehoz.C3.A1sa)
  - [3.6Root szelet létrehozása](#Root_szelet_l.C3.A9trehoz.C3.A1sa)
  - [3.7Mentse el a szelet elrendezését és lépjen ki](#Mentse_el_a_szelet_elrendez.C3.A9s.C3.A9t_.C3.A9s_l.C3.A9pjen_ki)
- [4Az fdisk használata a fizikai adathordozó particionálásához (csak ARC/AlphaBIOS esetén)](#Az_fdisk_haszn.C3.A1lata_a_fizikai_adathordoz.C3.B3_particion.C3.A1l.C3.A1s.C3.A1hoz_.28csak_ARC.2FAlphaBIOS_eset.C3.A9n.29)
  - [4.1Rendelkezésre álló fizikai adathordozók azonosítása](#Rendelkez.C3.A9sre_.C3.A1ll.C3.B3_fizikai_adathordoz.C3.B3k_azonos.C3.ADt.C3.A1sa)
  - [4.2Összes partíció törlése](#.C3.96sszes_part.C3.ADci.C3.B3_t.C3.B6rl.C3.A9se)
  - [4.3Boot partíció létrehozása](#Boot_part.C3.ADci.C3.B3_l.C3.A9trehoz.C3.A1sa)
  - [4.4Cserepartíció (swap) létrehozása](#Cserepart.C3.ADci.C3.B3_.28swap.29_l.C3.A9trehoz.C3.A1sa)
  - [4.5Root partíció létrehozása](#Root_part.C3.ADci.C3.B3_l.C3.A9trehoz.C3.A1sa)
  - [4.6Mentse el a partíció elrendezését, majd lépjen ki](#Mentse_el_a_part.C3.ADci.C3.B3_elrendez.C3.A9s.C3.A9t.2C_majd_l.C3.A9pjen_ki)
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

### Szeletek

Bár elméletileg lehetséges egy teljes adathordozót felhasználni egy Linux rendszer tárolására, a gyakorlatban ez szinte soha nem történik meg. Ehelyett a teljes adathordozó blokkeszközöket kisebb, kezelhetőbb blokkeszközökre osztják fel. Az Alpha rendszereken ezeket _szeleteknek_ hívják.

**Note**

A későbbi szakaszokban a telepítési utasítások az ARC/AlphaBIOS beállításhoz tartozó példaparticionálást fogják használni. Kérjük, hogy igazítsa a példákban bemutatottakat a személyes preferenciához!

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

## Az fdisk használata az fizikai adathordozó particionálásához (csak SRM esetében)

A következő részek bemutatják, hogy miként hozható létre az SRM-hez tartozó példa szelet-elrendezés.

Szelet
Leírás
/dev/sda1Swap szelet
/dev/sda2Root szelet
/dev/sda3Teljes adathordozó (szükséges)

Módosítsa a szelet-elrendezést az Ön személyes preferenciáinak megfelelően.

### Elérhető fizikai adathordozók azonosítása

A rendszerben futó adathordozók meghatározásához használja a következő parancsokat:

IDE esetében:

`root #` `dmesg | grep 'drive$'`

SCSI fizikai adathordozók esetében:

`root #` `dmesg | grep 'scsi'`

A kimenet megmutatja, hogy milyen fizikai adathordozókat észlelt a rendszer, valamint azok megfelelő /dev/ bejegyzését. A következő részekben feltételezzük, hogy a fizikai adathordozó egy SCSI fizikai adathordozó amely a /dev/sda helyen található meg.

### BSD fizikai adathordozó címke létrehozása

Ha a fizikai adathordozó teljesen üres, akkor először hozzon létre egy BSD fizikai adathordozó címkét. Alphán ezt nem lehet az fdisk segítségével megtenni, ezért itt a parted használata jelenti a megoldást.

Most indítsa el a parted particionáló szoftvert:

`root #` `parted /dev/sda `

```
Using /dev/sda
Welcome to GNU Parted! Type 'help' to view a list of commands.

```

`(parted)` `mklabel bsd`

```
Warning: The existing disk label on /dev/sda will be destroyed and all data on this disk will be lost. Do you want to continue?
Yes/No? yes
(parted) quit
Information: You may need to update /etc/fstab.

```

Most, hogy létrehoztuk a BSD fizikai adathordozó címkét a fizikai adathordozónkon, folytassa a szeletek létrehozását. Ezt a parted segítségével lehet megtenni, vagy ahogyan az alábbi példákban, az fdisk használatával is megtehető:

`root #` `fdisk /dev/sda`

### Összes szelet törlése

Ha a fizikai adathordozó teljesen üres, akkor először hozzon létre egy BSD lemezcímkét.

`Command (m for help):` `b`

```
/dev/sda contains no disklabel.
Do you want to create a disklabel? (y/n) y
A bunch of drive-specific info will show here
3 partitions:
#       start       end      size     fstype   [fsize bsize   cpg]
  c:        1      5290*     5289*    unused        0     0

```

Kezdjük azzal, hogy törlünk minden szeletet, kivéve a 'c'-szeletet (ez a BSD lemezcímkék használatának előfeltétele). Az alábbiakban bemutatjuk, hogyan lehet törölni egy szeletet (a példában az 'a'-t használjuk). Ismételje meg a folyamatot, hogy törölje az összes többi szeletet (ismételten kivéve a 'c'-szeletet).

Használja a `p` gombot az összes meglévő szelet megtekintéséhez. A `d` gombot egy szelet törlésére használják.

`BSD lemezcímke parancs (m a súgóhoz):` `p`

```
8 partitions:
#       start       end      size     fstype   [fsize bsize   cpg]
  a:        1       235*      234*    4.2BSD     1024  8192    16
  b:      235*      469*      234*      swap
  c:        1      5290*     5289*    unused        0     0
  d:      469*     2076*     1607*    unused        0     0
  e:     2076*     3683*     1607*    unused        0     0
  f:     3683*     5290*     1607*    unused        0     0
  g:      469*     1749*     1280     4.2BSD     1024  8192    16
  h:     1749*     5290*     3541*    unused        0     0

```

`BSD disklabel command (m for help):` `d`

```
Partition (a-h): a

```

Miután ezt a folyamatot minden szelet esetében megismételte, a listázásnak valami hasonlót kell mutatnia:

`BSD lemezcímke parancs (m a súgóhoz):` `p`

```
3 partitions:
#       start       end      size     fstype   [fsize bsize   cpg]
  c:        1      5290*     5289*    unused        0     0

```

### Swap szelet létrehozása

Alpha alapú rendszereken nincs szükség külön boot szeletre, de ebben a példában létrehozunk egy külön /boot szeletet a kernelfájlok és a boot loader beállításfájlok számára. Az első cilinder azonban nem használható, mivel az _aboot_ képetfájlt oda helyezik.

Létrehozunk egy swap szeletet, amely a harmadik cilindernél kezdődik, és összesen 1 GB méretű lesz. Használja a `n` billentyűgombot egy új szelet létrehozásához. A szelet létrehozása után változtassa meg annak típusát `1` (egy) típusra, ami swap partíciót jelent.

`BSD lemezcímke parancs (m a súgóhoz):` `n`

```
Partition (a-p): a
First cylinder (1-5290, default 1): 3
Last cylinder or +size or +sizeM or +sizeK (3-5290, default 5290): +1024M

```

`BSD lemezcímke parancs (m a súgóhoz):` `t`

```
Partition (a-c): a
Hex code (type L to list codes): 1

```

Ezek után egy elrendezésnek, amely hasonló a következőhöz, kell megjelennie:

`BSD lemezcímke parancs (m a súgóhoz):` `p`

```
3 partitions:
#       start       end      size     fstype   [fsize bsize   cpg]
  a:        3      1003      1001       swap
  c:        1      5290*     5289*    unused        0     0

```

### Boot szelet létrehozása

Hozzon létre egy boot fájlrendszert, amely tartalmazza a kernel fájlokat és a boot loader beállításfájlt (aboot.conf). Az aboot csak ext2 és ext3 fájlrendszereket támogat. Hozza létre a boot szeletet a swap szelet utáni első cilindertől kezdve. Használja a `p` parancsot annak megtekintésére, hogy hol ér véget a swap szelet. Példánkban ez 1003-nál van, így a boot szelet 1004-nél kezdődik.

`BSD lemezcímke parancs (m a súgóhoz):` `n`

```
Partition (a-p): b
First cylinder (1-5290, default 1): 1004
Last cylinder or +size or +sizeM or +sizeK (3-5290, default 5290): +1024M

```

`BSD lemezcímke parancs (m a súgóhoz):` `t`

```
Partition (a-c): b
Hex code (type L to list codes): 08

```

Ezek után egy olyan elrendezésnek kell megjelennie, amely hasonló az alábbihoz:

`BSD lemezcímke parancs (m a súgóhoz):` `p`

```
3 partitions:
#       start       end      size     fstype   [fsize bsize   cpg]
  a:        3      1003      1001       swap
  b:     1004      2005      1001       ext2
  c:        1      5290*     5289*    unused        0     0

```

### Root szelet létrehozása

Most létrehozzuk a root szeletet, amely a boot szelet utáni első cilindernél kezdődik. Használja a `p` parancsot annak megtekintésére, hogy hol ér véget a boot szelet. Példánkban ez 2005-nél van, így a root szelet 2006-nál kezdődik.

Egy másik probléma, hogy jelenleg van egy hiba a fdisk programban, amely miatt azt gondolja, hogy a rendelkezésre álló cilinderek száma eggyel több, mint a valós szám. Más szavakkal, amikor az utolsó cilinderről kérdez, akkor csökkentse a cilinderszámot (ebben a példában: 5290) eggyel.

Amikor a szelet létrejön, változtassa meg a típusát 8-ra, ami az ext2 fájlrendszert jelöli.

`BSD lemezcímke parancs (m a súgóhoz):` `n`

```
Partition (a-p): b
First cylinder (1-5290, default 1): 2006
Last cylinder or +size or +sizeM or +sizeK (1004-5290, default 5290): 5289

```

`BSD lemezcímke parancs (m a súgóhoz):` `t`

```
Partition (a-c): b
Hex code (type L to list codes): 8

```

Az eredményül kapott szelet elrendezésnek most valami hasonlónak kell lennie:

`BSD lemezcímke parancs (m a súgóhoz):` `p`

```
3 partitions:
#       start       end      size     fstype   [fsize bsize   cpg]
  a:        3      1003      1001       swap
  b:     1004      2005      4286       ext2
  c:        1      5290*     5289*    unused        0     0
  d:     2006      5289      3283       ext2

```

### Mentse el a szelet elrendezését és lépjen ki

Lépjen ki a fdisk alkalmazásból a `w` megadásával. Ez a művelet elmenti a szelet elrendezést is.

`Command (m for help):` `w`

## Az fdisk használata a fizikai adathordozó particionálásához (csak ARC/AlphaBIOS esetén)

A következő részek bemutatják, hogy miként kell létrehozni a példában szereplő partíció elrendezést ARC/AlphaBIOS esetén:

Partíció
Leírás
/dev/sda1Boot partíció
/dev/sda2Swap partíció
/dev/sda3Root partíció

Módosítsa a partíció elrendezését személyes preferenciák szerint.

### Rendelkezésre álló fizikai adathordozók azonosítása

Az alábbi parancsok segítségével meghatározhatja, hogy mely fizikai adathordozók működnek:

IDE adathordozók számára:

`root #` `dmesg | grep 'drive$'`

SCSI fizikai adathordozók számára:

`root #` `dmesg | grep 'scsi'`

Ebből a kimenetből könnyen látható, hogy mely fizikai adathordozókat észlelte a rendszer, valamint azok megfelelő /dev/ bejegyzéseit. A következő részekben azt feltételezzük, hogy a fizikai adathordozó egy SCSI fizikai adathordozó, amelynek elérési útja az /dev/sda útvonalon található meg.

Most indítsa el az fdisk particionáló szoftvert:

`root #` `fdisk /dev/sda`

### Összes partíció törlése

Ha a fizikai adathordozó teljesen üres, akkor először hozzon létre egy DOS adathordozó táblázatot (lemezcímkét).

`Parancs (m a súgóhoz):` `o`

```
Building a new DOS disklabel.

```

Kezdjük az összes partíció törlésével. Az alábbiakban bemutatom, hogyan töröljön egy partíciót (a példában az '1'-est használjuk). Ismételje meg a folyamatot, hogy törölje a többi partíciót is.

Használja a `p` billentyűgombot az összes létező partíció megtekintéséhez. A `d` billentyűgomb a partíció törlésére szolgál.

`command (m for help):` `p`

```
Disk /dev/sda: 9150 MB, 9150996480 bytes
64 heads, 32 sectors/track, 8727 cylinders
Units = cylinders of 2048 * 512 = 1048576 bytes

   Device Boot      Start         End      Blocks   Id  System
/dev/sda1               1         478      489456   83  Linux
/dev/sda2             479        8727     8446976    5  Extended
/dev/sda5             479        1433      977904   83  Linux Swap
/dev/sda6            1434        8727     7469040   83  Linux

```

`command (m for help):` `d`

```
Partition number (1-6): 1

```

### Boot partíció létrehozása

Az Alpha rendszereken, amelyek MILO-t használnak a bootoláshoz, létre kell hozni egy kis vfat boot partíciót.

`Command (m for help):` `n`

```
Command action
  e   extended
  p   primary partition (1-4)
p
Partition number (1-4): 1
First cylinder (1-8727, default 1): 1
Last cylinder or +size or +sizeM or +sizeK (1-8727, default 8727): +16M

```

`Command (m for help):` `t`

```
Selected partition 1
Hex code (type L to list codes): 6
Changed system type of partition 1 to 6 (FAT16)

```

### Cserepartíció (swap) létrehozása

Létrehozunk egy cserepartíciót, amelynek teljes mérete 1 GB. Használja a `n` billentyűgombot egy új partíció létrehozásához.

`Command (m for help):` `n`

```
Command action
  e   extended
  p   primary partition (1-4)
p
Partition number (1-4): 2
First cylinder (17-8727, default 17): 17
Last cylinder or +size or +sizeM or +sizeK (17-8727, default 8727): +1000M

```

`Command (m for help):` `t`

```
Partition number (1-4): 2
Hex code (type L to list codes): 82
Changed system type of partition 2 to 82 (Linux swap)

```

A fenti lépések után a következőhöz hasonló elrendezés látható:

`Command (m for help):` `p`

```
Disk /dev/sda: 9150 MB, 9150996480 bytes
64 heads, 32 sectors/track, 8727 cylinders
Units = cylinders of 2048 * 512 = 1048576 bytes

   Device Boot      Start         End      Blocks   Id  System
/dev/sda1               1          16       16368    6  FAT16
/dev/sda2              17         971      977920   82  Linux swap

```

### Root partíció létrehozása

Most létrehozzuk a root partíciót. Ismét használja a `n` parancsot.

`Command (m for help):` `n`

```
Command action
  e   extended
  p   primary partition (1-4)
p
Partition number (1-4): 3
First cylinder (972-8727, default 972): 972
Last cylinder or +size or +sizeM or +sizeK (972-8727, default 8727): 8727

```

Ezek után egy a következőhöz hasonló elrendezésnek kell megjelennie:

`Command (m for help):` `p`

```
Disk /dev/sda: 9150 MB, 9150996480 bytes
64 heads, 32 sectors/track, 8727 cylinders
Units = cylinders of 2048 * 512 = 1048576 bytes

   Device Boot      Start         End      Blocks   Id  System
/dev/sda1               1          16       16368    6  FAT16
/dev/sda2              17         971      977920   82  Linux swap
/dev/sda3             972        8727     7942144   83  Linux

```

### Mentse el a partíció elrendezését, majd lépjen ki

Mentse a fdisk programban végrehajtott módosításokat a `w` billentyűgomb megnyomásával.

`Command (m for help):` `w`

Most, hogy a partíciók létrejöttek, folytassa a [Fájlrendszerek létrehozása](#F.C3.A1jlrendszerek_l.C3.A9trehoz.C3.A1sa) című rész követésével.

## Fájlrendszerek létrehozása

**Warning**

SSD vagy NVMe adathordozó használatakor bölcs dolog ellenőrizni a firmware-frissítéseket. Különösen egyes Intel SSD adathordozók (600p és 6000p) firmware-frissítést igényelnek az XFS I/O használati minták által okozott [lehetséges adatsérülések miatt](https://bugzilla.redhat.com/show_bug.cgi?id=1402533). A probléma a firmware szintjén van, és nem az XFS fájlrendszer hibája. A smartctl segédprogram segíthet az adathordozzó eszköz modelljének és firmware-verziójának ellenőrzésében.

### Bevezetés

Most, hogy a partíciók elkészültek, ideje fájlrendszert helyezni rájuk. A következő részben a Linux által támogatott különféle fájlrendszereket ismertetjük. Azok az olvasók, akik már tudják, hogy melyik fájlrendszert fogják használni, folytathatják a [Fájlrendszer rárakása egy partícióra](/wiki/Handbook:Alpha/Installation/Disks/hu#Applying_a_filesystem_to_a_partition "Handbook:Alpha/Installation/Disks/hu") című bekezdéssel. A többi felhasználónak érdemes továbbolvasniuk, hogy megismerjék az alkalmazható fájlrendszereket...

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

Innentől folytatható az a rendszertelepítés, amely korábban el lett kezdve, de a telepítési folyamat nem let végig befejezve. Használja ezt a hivatkozást állandó hivatkozásként: [A telepítés folytatása itt kezdődik](/wiki/Handbook:Alpha/Installation/Disks/hu#Resumed_installations_start_here "Handbook:Alpha/Installation/Disks/hu").

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

Később az utasításokban a proc fájlrendszer (a kernellel kapcsolatban álló virtuális interfész) és a többi kernel pszeudofájlrendszer lesz felcsatolva. Először viszont még a [Gentoo-stage fájlt](/wiki/Handbook:Alpha/Installation/Stage/hu "Handbook:Alpha/Installation/Stage/hu") ki kell csomagolnunk.

[← Hálózat beállítása](/wiki/Handbook:Alpha/Installation/Networking/hu "Previous part") [Kezdőlap](/wiki/Handbook:Alpha/hu "Handbook:Alpha/hu") [Stage fájl telepítése →](/wiki/Handbook:Alpha/Installation/Stage/hu "Next part")