# Disks

Other languages:

- [Deutsch](/wiki/Handbook:MIPS/Installation/Disks/de "Handbuch:MIPS/Installation/Festplatten (100% translated)")
- [English](/wiki/Handbook:MIPS/Installation/Disks "Handbook:MIPS/Installation/Disks (100% translated)")
- [español](/wiki/Handbook:MIPS/Installation/Disks/es "Manual de Gentoo: MIPS/Instalación/Discos (100% translated)")
- [français](/wiki/Handbook:MIPS/Installation/Disks/fr "Handbook:MIPS/Installation/Disks/fr (100% translated)")
- [italiano](/wiki/Handbook:MIPS/Installation/Disks/it "Handbook:MIPS/Installation/Disks/it (100% translated)")
- magyar
- [polski](/wiki/Handbook:MIPS/Installation/Disks/pl "Handbook:MIPS/Installation/Disks (100% translated)")
- [português do Brasil](/wiki/Handbook:MIPS/Installation/Disks/pt-br "Manual:MIPS/Instalação/Discos (100% translated)")
- [русский](/wiki/Handbook:MIPS/Installation/Disks/ru "Handbook:MIPS/Installation/Disks (100% translated)")
- [தமிழ்](/wiki/Handbook:MIPS/Installation/Disks/ta "கையேடு:MIPS/நிறுவல்/வட்டுக்கள் (100% translated)")
- [中文（中国大陆）‎](/wiki/Handbook:MIPS/Installation/Disks/zh-cn "手册：MIPS/安装/准备磁盘 (100% translated)")
- [日本語](/wiki/Handbook:MIPS/Installation/Disks/ja "ハンドブック:MIPS/インストール/ディスク (100% translated)")
- [한국어](/wiki/Handbook:MIPS/Installation/Disks/ko "Handbook:MIPS/Installation/Disks/ko (100% translated)")

[MIPS kézikönyv](/wiki/Handbook:MIPS/hu "Handbook:MIPS/hu")[A Gentoo Linux telepítése](/wiki/Handbook:MIPS/Full/Installation/hu "Handbook:MIPS/Full/Installation/hu")[A telepítésről](/wiki/Handbook:MIPS/Installation/About/hu "Handbook:MIPS/Installation/About/hu")[Telepítőképfájl kiválasztása](/wiki/Handbook:MIPS/Installation/Media/hu "Handbook:MIPS/Installation/Media/hu")[Hálózat beállítása](/wiki/Handbook:MIPS/Installation/Networking/hu "Handbook:MIPS/Installation/Networking/hu")Adathordozók előkészítése[Stage fájl](/wiki/Handbook:MIPS/Installation/Stage/hu "Handbook:MIPS/Installation/Stage/hu")[Alaprendszer telepítése](/wiki/Handbook:MIPS/Installation/Base/hu "Handbook:MIPS/Installation/Base/hu")[Kernel beállítása](/wiki/Handbook:MIPS/Installation/Kernel/hu "Handbook:MIPS/Installation/Kernel/hu")[Rendszer beállítása](/wiki/Handbook:MIPS/Installation/System/hu "Handbook:MIPS/Installation/System/hu")[Eszközök telepítése](/wiki/Handbook:MIPS/Installation/Tools/hu "Handbook:MIPS/Installation/Tools/hu")[Bootloader beállítása](/wiki/Handbook:MIPS/Installation/Bootloader/hu "Handbook:MIPS/Installation/Bootloader/hu")[Telepítés véglegesítése](/wiki/Handbook:MIPS/Installation/Finalizing/hu "Handbook:MIPS/Installation/Finalizing/hu")[Munka a Gentoo rendszerrel](/wiki/Handbook:MIPS/Full/Working/hu "Handbook:MIPS/Full/Working/hu")[Portage bemutatása](/wiki/Handbook:MIPS/Working/Portage/hu "Handbook:MIPS/Working/Portage/hu")[USE jelölőzászlók](/wiki/Handbook:MIPS/Working/USE/hu "Handbook:MIPS/Working/USE/hu")[Portage jellemzői](/wiki/Handbook:MIPS/Working/Features/hu "Handbook:MIPS/Working/Features/hu")[Init-szkript rendszer](/wiki/Handbook:MIPS/Working/Initscripts/hu "Handbook:MIPS/Working/Initscripts/hu")[Környezeti változók](/wiki/Handbook:MIPS/Working/EnvVar/hu "Handbook:MIPS/Working/EnvVar/hu")[Munka a Portage szoftvercsomag-kezelővel](/wiki/Handbook:MIPS/Full/Portage/hu "Handbook:MIPS/Full/Portage/hu")[Fájlok és könyvtárak](/wiki/Handbook:MIPS/Portage/Files/hu "Handbook:MIPS/Portage/Files/hu")[Változók](/wiki/Handbook:MIPS/Portage/Variables/hu "Handbook:MIPS/Portage/Variables/hu")[Szoftverágak keverése](/wiki/Handbook:MIPS/Portage/Branches/hu "Handbook:MIPS/Portage/Branches/hu")[További eszközök](/wiki/Handbook:MIPS/Portage/Tools/hu "Handbook:MIPS/Portage/Tools/hu")[Egyéni szoftvercsomag-tárolók](/wiki/Handbook:MIPS/Portage/CustomTree/hu "Handbook:MIPS/Portage/CustomTree/hu")[Fejlett funkciók](/wiki/Handbook:MIPS/Portage/Advanced/hu "Handbook:MIPS/Portage/Advanced/hu")[Hálózat beállítása OpenRC init-rendszeren](/wiki/Handbook:MIPS/Full/Networking/hu "Handbook:MIPS/Full/Networking/hu")[Munka elkezdése](/wiki/Handbook:MIPS/Networking/Introduction/hu "Handbook:MIPS/Networking/Introduction/hu")[Fejlett beállítások](/wiki/Handbook:MIPS/Networking/Advanced/hu "Handbook:MIPS/Networking/Advanced/hu")[Moduláris hálózat](/wiki/Handbook:MIPS/Networking/Modular/hu "Handbook:MIPS/Networking/Modular/hu")[Vezeték nélküli (Wi-Fi)](/wiki/Handbook:MIPS/Networking/Wireless/hu "Handbook:MIPS/Networking/Wireless/hu")[Funkcionalitás hozzáadása](/wiki/Handbook:MIPS/Networking/Extending/hu "Handbook:MIPS/Networking/Extending/hu")[Dinamikus menedzsment](/wiki/Handbook:MIPS/Networking/Dynamic/hu "Handbook:MIPS/Networking/Dynamic/hu")

## Contents

- [1Bevezetés a blokktípusú eszközökbe](#Bevezet.C3.A9s_a_blokkt.C3.ADpus.C3.BA_eszk.C3.B6z.C3.B6kbe)
  - [1.1Blokkeszközök](#Blokkeszk.C3.B6z.C3.B6k)
  - [1.2Partíciók](#Part.C3.ADci.C3.B3k)
- [2Partíciós séma megtervezése](#Part.C3.ADci.C3.B3s_s.C3.A9ma_megtervez.C3.A9se)
  - [2.1Hány partíció és mekkora méretű?](#H.C3.A1ny_part.C3.ADci.C3.B3_.C3.A9s_mekkora_m.C3.A9ret.C5.B1.3F)
  - [2.2Mi a helyzet a swap területtel?](#Mi_a_helyzet_a_swap_ter.C3.BClettel.3F)
- [3Az fdisk használata](#Az_fdisk_haszn.C3.A1lata)
  - [3.1SGI számítógépek: SGI adathordozó táblázat (címke) létrehozása](#SGI_sz.C3.A1m.C3.ADt.C3.B3g.C3.A9pek:_SGI_adathordoz.C3.B3_t.C3.A1bl.C3.A1zat_.28c.C3.ADmke.29_l.C3.A9trehoz.C3.A1sa)
  - [3.2SGI Volume Header átméretezése](#SGI_Volume_Header_.C3.A1tm.C3.A9retez.C3.A9se)
  - [3.3Cobalt adathordozók particionálása](#Cobalt_adathordoz.C3.B3k_particion.C3.A1l.C3.A1sa)
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

### Partíciók

Bár elméletileg lehetséges egy teljes adathordozót felhasználni a Linux rendszer tárolására, gyakorlatban ez szinte soha nem történik meg. Ehelyett a teljes adathordozó blokkeszközei kisebb, kezelhetőbb blokkeszközökre vannak felosztva. Ezeket a blokkeszközöket partícióknak nevezzük.

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

## Az fdisk használata

### SGI számítógépek: SGI adathordozó táblázat (címke) létrehozása

Minden adathordozó egy SGI rendszerben igényel egy SGI adathordozó táblázatot (címkét), amely hasonló funkciót lát el, mint a Sun és MS-DOS adathordozó táblázatok (címkék). Tárolja az adathordozó partícióinak az adatait. Egy új SGI adathordozó táblázat (címke) létrehozása két speciális partíciót hoz létre az adathordozón:

**Warning**

Az SGI Volume Header-nek a 0. cilinderrel kell kezdődnie. Ennek elmulasztása azt eredményezi, hogy az adathordozóról való bootolás sikertelen lesz.

Az alábbiakban egy példát láthat egy fdisk munkamenet kivonatára. Olvassa el figyelmesen, és igazítsa azt az Ön személyes preferenciáihoz...

`root #` `fdisk /dev/sda`

Váltson szakértői módba:

`Command (m for help):` `x`

Az `m` billentyűgomb megnyomásával megjelenik az összes lehetőséget tartalmazó teljes menü:

`Expert command (m for help):` `m`

```
Command action
   b   move beginning of data in a partition
   c   change number of cylinders
   d   print the raw data in the partition table
   e   list extended partitions
   f   fix partition order
   g   create an IRIX (SGI) partition table
   h   change number of heads
   m   print this menu
   p   print the partition table
   q   quit without saving changes
   r   return to main menu
   s   change number of sectors/track
   v   verify the partition table
   w   write table to disk and exit

```

Készítsen SGI adathordozó táblázatot (címkét):

`Expert command (m for help):` `g`

```
Building a new SGI disklabel. Changes will remain in memory only,
until you decide to write them. After that, of course, the previous
content will be irrecoverably lost.

```

Térjen vissza a főmenübe:

`Expert command (m for help):` `r`

Tekintse meg a jelenlegi partícióelrendezést:

`Command (m for help):` `p`

```
Disk /dev/sda (SGI disk label): 64 heads, 32 sectors, 17482 cylinders
Units = cylinders of 2048 * 512 bytes

----- partitions -----
Pt#     Device  Info     Start       End   Sectors  Id  System
 9:  /dev/sda1               0         4     10240   0  SGI volhdr
11:  /dev/sda2               0     17481  35803136   6  SGI volume
----- Bootinfo -----
Bootfile: /unix
----- Directory Entries -----

```

**Note**

Ha az adathordozó már rendelkezik meglévő SGI adathordozó táblázattal (címkével), akkor az fdisk nem engedi az új adathordozó táblázat (címke) létrehozását. Erre két megoldás van. Az egyik, hogy létrehozzon egy Sun vagy MS-DOS adathordozó táblázatot (címkét), mentse el a változtatásokat az adathordozóra, majd indítsa újra az fdisk particionáló segédprogramot. A második, hogy a partíciós táblázatot null adatokkal felülírja a következő parancs segítségével: `dd if=/dev/zero of=/dev/sda bs=512 count=1`

### SGI Volume Header átméretezése

**Important**

Ez a lépés gyakran szükséges egy hiba miatt az fdisk particionáló segéprogramban. Valamilyen oknál fogva a volume header nem megfelelően jön létre, ennek eredményeképpen a 0. cilinderen kezdődik és végződik. Ez megakadályozza több partíció létrehozását. A probléma megoldásához olvasson tovább.

Most, hogy létrehozott egy SGI adathordozó táblázatot (címkét), a partíciók meghatározhatók. Az előző példában már két partíció van meghatározva. Ezek azok a speciális partíciók, amelyeket fent említettünk, és amelyeket általában nem szabad módosítani. Azonban a Gentoo telepítéséhez be kell töltenünk egy bootloadert, és esetleg több bináris kernelképfájlt is (a rendszer típusától függően) közvetlenül a volume header-be. Maga a volume header akár nyolc képfájlt is el tud tárolni bármilyen méretben, és minden képfájlnak nyolc karakter hosszúságú nevet lehet adni.

A volume header méretének növelése nem teljesen egyszerű, van benne némi trükk. Az fdisk furcsa viselkedése miatt nem lehet egyszerűen törölni és újra hozzáadni a volume header-t. Az alábbi példában egy 50 MB mértetű volume header-t hozunk létre egy 50 MB méretű /boot/ partícióval együtt. Az adathordozó tényleges elrendezése változhat, de ez csak szemléltetőcélokat szolgál.

Hozzon létre egy új partíciót:

`Command (m for help):` `n`

```
Partition number (1-16): 1
First cylinder (5-8682, default 5): 51
 Last cylinder (51-8682, default 8682): 101

```

Figyelte meg, hogy vajon az fdisk csak azt engedi-e, hogy az 1. partíciót újra létrehozzák, és minimum az 5. cilindernél kezdődjön? Ha megpróbálnánk törölni és újra létrehozni az SGI Volume Header-t ezen a módon, akkor ugyanezzel a problémával találkoznánk. Példánkban a /boot/ partíciót 50 MB-osra szeretnénk, ezért azt a 51. cilindernél kezdjük (ne feledje, hogy a Volume Header-nek a 0. cilindernél kell kezdődnie), és a végső cilinderét a 101-re állítjuk, amely körülbelül 50 MB méretű lesz (+/- 1-5 MB).

Törölje a partíciót:

`Command (m for help):` `d`

```
Partition number (1-16): 9

```

Most hozza létre ismét:

`Command (m for help):` `n`

```
Partition number (1-16): 9
First cylinder (0-50, default 0): 0
 Last cylinder (0-50, default 50): 50

```

Ha nem biztos abban, hogy miként kell használni az fdisk particionáló segédprogramot, akkor tekintse meg lejjebb a Cobaltok particionálására vonatkozó utasításokat. A koncepciók pontosan ugyanazok – csak ne feledje, hogy a volume header-t és a teljes adathordozó partíciókat hagyja érintetlenül.

Miután ez megtörtént, hozzon létre minden szükséges partíciót. Miután az összes partíció elrendezése befejeződött, győződjön meg arról, hogy a swap partíció azonosítóját 82-re állította, amely a Linux Swap partíciót jelenti. Alapértelmezés szerint ez 83 lesz, azaz Linux Native.

### Cobalt adathordozók particionálása

Cobalt gépeken a BOOTROM egy MS-DOS MBR-t vár el, ezért a meghajtó particionálása viszonylag egyszerű – valójában ugyanúgy történik, mint egy Intel x86 gépen. Azonban van néhány dolog, amit figyelembe kell venni.

- A Cobalt firmware /dev/sda1 Linux partíciót vár, amely EXT2 Revision 0 formátumú. Az EXT2 Revision 1 partíciók NEM MŰKÖDNEK! (A Cobalt BOOTROM csak az EXT2r0-t érti meg).
- Az említett partíciónak tartalmaznia kell egy tömörített ELF bináris képfájlt, a vmlinux.gz képfájlt, a partíció gyökerében, amelyet a kernelként tölt be.

Ezért javasolt egy körülbelül 20 MB méretű /boot/ partíciót létrehozni, amely EXT2r0 formátumú, és amelyre telepíthető a CoLo és a kernel. Ez lehetővé teszi, hogy a felhasználó egy modern fájlrendszert (EXT3 vagy ReiserFS) használjon a gyökérfájlrendszerhez.

A példában feltételezzük, hogy a /dev/sda1 partíció később a /boot/ partícióként lesz felcsatolva. Ha ezt gyökérfájlrendszerré kívánja tenni, akkor tartsa szem előtt a PROM elvárásait.

Tehát folytatva... A partíciók létrehozásához írja be a fdisk /dev/sda parancsot a promptnál. Az alapvető parancsok, amelyeket ismernie kell, ezek:

CODE **Az fdisk legfontosabb parancsainak a listája**

```
o: Wipe out old partition table, starting with an empty MS-DOS partition table
n: New Partition
t: Change Partition Type
    Use type 82 for Linux Swap, 83 for Linux FS
d: Delete a partition
p: Display (print) Partition Table
q: Quit -- leaving old partition table as is.
w: Quit -- writing partition table in the process.

```

`root #` `fdisk /dev/sda`

```
The number of cylinders for this disk is set to 19870.
There is nothing wrong with that, but this is larger than 1024,
and could in certain setups cause problems with:
1) software that runs at boot time (e.g., old versions of LILO)
2) booting and partitioning software from other OSs
   (e.g., DOS FDISK, OS/2 FDISK)

```

Kezdje a meglévő partíciók törlésével:

`Command (m for help):` `o`

```
Building a new DOS disklabel. Changes will remain in memory only,
until you decide to write them. After that, of course, the previous
content won't be recoverable.


The number of cylinders for this disk is set to 19870.
There is nothing wrong with that, but this is larger than 1024,
and could in certain setups cause problems with:
1) software that runs at boot time (e.g., old versions of LILO)
2) booting and partitioning software from other OSs
   (e.g., DOS FDISK, OS/2 FDISK)
Warning: invalid flag 0x0000 of partition table 4 will be corrected by w(rite)

```

Most ellenőrizze, hogy a partíciós táblázat üres-e, a `p` parancs használatával:

`Command (m for help):` `p`

```
Disk /dev/sda: 10.2 GB, 10254827520 bytes
16 heads, 63 sectors/track, 19870 cylinders
Units = cylinders of 1008 * 512 = 516096 bytes

   Device Boot      Start         End      Blocks   Id  System

```

Hozza létre a /boot partíciót:

`Command (m for help):` `n`

```
Command action
   e   extended
   p   primary partition (1-4)
p
Partition number (1-4): 1
First cylinder (1-19870, default 1):
Last cylinder or +size or +sizeM or +sizeK (1-19870, default 19870): +20M

```

A partíciók kilistázásakor figyelje meg az újonnan létrehozott partíciót:

`Command (m for help):` `p`

```
Disk /dev/sda: 10.2 GB, 10254827520 bytes
16 heads, 63 sectors/track, 19870 cylinders
Units = cylinders of 1008 * 512 = 516096 bytes

   Device Boot      Start         End      Blocks   Id  System
/dev/sda1               1          40       20128+  83  Linux

```

Hozzunk létre egy kiterjesztett partíciót, amely lefedi az adathordozó fennmaradó még üres részét. Ebben a kiterjesztett partícióban létrehozzuk a többi (logikai partíciókat):

`Command (m for help):` `n`

```
Command action
   e   extended
   p   primary partition (1-4)
e
Partition number (1-4): 2
First cylinder (41-19870, default 41):
Using default value 41
Last cylinder or +size or +sizeM or +sizeK (41-19870, default 19870):
Using default value 19870

```

Most hozzuk létre a / partíciót, valamint a /usr, /var és egyéb szükséges partíciókat.

`Command (m for help):` `n`

```
Command action
   l   logical (5 or over)
   p   primary partition (1-4)
l
First cylinder (41-19870, default 41):<Press ENTER>
Using default value 41
Last cylinder or +size or +sizeM or +sizeK (41-19870, default 19870): +500M

```

Ismételje meg ezt szükség szerint.

Végül, de nem utolsósorban, a swap terület. Ajánlott, hogy legalább 250 MB swap terület legyen, előnyben részesítve az 1 GB méretet:

`Command (m for help):` `n`

```
Command action
   l   logical (5 or over)
   p   primary partition (1-4)
l
First cylinder (17294-19870, default 17294): <Press ENTER>
Using default value 17294
Last cylinder or +size or +sizeM or +sizeK (1011-19870, default 19870): <Press ENTER>
Using default value 19870

```

Amikor ellenőrzi a partíciós táblázatot, mindennek készen kell lennie - egyetlen dolgot leszámítva.

`Command (m for help):` `p`

```
Disk /dev/sda: 10.2 GB, 10254827520 bytes
16 heads, 63 sectors/track, 19870 cylinders
Units = cylinders of 1008 * 512 = 516096 bytes

Device Boot      Start         End      Blocks      ID  System
/dev/sda1               1          21       10552+  83  Linux
/dev/sda2              22       19870    10003896    5  Extended
/dev/sda5              22        1037      512032+  83  Linux
/dev/sda6            1038        5101     2048224+  83  Linux
/dev/sda7            5102        9165     2048224+  83  Linux
/dev/sda8            9166       13229     2048224+  83  Linux
/dev/sda9           13230       17293     2048224+  83  Linux
/dev/sda10          17294       19870     1298776+  83  Linux

```

Vegye észre, hogy a #10, a swap partíció, még mindig 83-as típusú. Változtassuk meg a megfelelő típusra:

`Command (m for help):` `t`

```
Partition number (1-10): 10
Hex code (type L to list codes): 82
Changed system type of partition 10 to 82 (Linux swap)

```

Most ellenőrizze le:

`Command (m for help):` `p`

```
Disk /dev/sda: 10.2 GB, 10254827520 bytes
16 heads, 63 sectors/track, 19870 cylinders
Units = cylinders of 1008 * 512 = 516096 bytes

Device Boot      Start         End      Blocks      ID  System
/dev/sda1               1          21       10552+  83  Linux
/dev/sda2              22       19870    10003896    5  Extended
/dev/sda5              22        1037      512032+  83  Linux
/dev/sda6            1038        5101     2048224+  83  Linux
/dev/sda7            5102        9165     2048224+  83  Linux
/dev/sda8            9166       13229     2048224+  83  Linux
/dev/sda9           13230       17293     2048224+  83  Linux
/dev/sda10          17294       19870     1298776+  82  Linux Swap

```

Kiírjuk az új partíciós táblázatot (címkét):

`Command (m for help):` `w`

```
The partition table has been altered!

Calling ioctl() to re-read partition table.
Syncing disks.

```

## Fájlrendszerek létrehozása

**Warning**

SSD vagy NVMe adathordozó használatakor bölcs dolog ellenőrizni a firmware-frissítéseket. Különösen egyes Intel SSD adathordozók (600p és 6000p) firmware-frissítést igényelnek az XFS I/O használati minták által okozott [lehetséges adatsérülések miatt](https://bugzilla.redhat.com/show_bug.cgi?id=1402533). A probléma a firmware szintjén van, és nem az XFS fájlrendszer hibája. A smartctl segédprogram segíthet az adathordozzó eszköz modelljének és firmware-verziójának ellenőrzésében.

### Bevezetés

Most, hogy a partíciók elkészültek, ideje fájlrendszert helyezni rájuk. A következő részben a Linux által támogatott különféle fájlrendszereket ismertetjük. Azok az olvasók, akik már tudják, hogy melyik fájlrendszert fogják használni, folytathatják a [Fájlrendszer rárakása egy partícióra](/wiki/Handbook:MIPS/Installation/Disks/hu#Applying_a_filesystem_to_a_partition "Handbook:MIPS/Installation/Disks/hu") című bekezdéssel. A többi felhasználónak érdemes továbbolvasniuk, hogy megismerjék az alkalmazható fájlrendszereket...

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

Például ahhoz, hogy a gyökérpartíció (tehát a root partíció) (/dev/sda5) fájlrendszertípusa xfs legyen, ahogy a partíciókészítés példa szerkezetében is szerepel, Önnek a következő parancsokat kell futtatnia:

`root #` `mkfs.xfs /dev/sda5`

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

`root #` `mkswap /dev/sda10`

**Note**

Innentől folytatható az a rendszertelepítés, amely korábban el lett kezdve, de a telepítési folyamat nem let végig befejezve. Használja ezt a hivatkozást állandó hivatkozásként: [A telepítés folytatása itt kezdődik](/wiki/Handbook:MIPS/Installation/Disks/hu#Resumed_installations_start_here "Handbook:MIPS/Installation/Disks/hu").

A swap partíciót aktiválni is kell. Használja a swapon parancsot:

`root #` `swapon /dev/sda10`

Ez az 'aktiválás' azért szükséges most, mert a swap partíciót újonnan hozzuk létre a Live ISO telepítőkörnyezetben. A rendszer újraindítása után mindaddig, amíg a swap partíció megfelelően van definiálva az fstab fájlban vagy más csatolási mechanizmusban, a swap terület automatikusan fog aktiválódni.

## Gyökérpartíció (root partíció) felcsatolása

Előfordulhat, hogy bizonyos Live ISO telepítőkörnyezetekből hiányzik a javasolt csatolási pont a Gentoo gyökérpartíciójához (/mnt/gentoo), vagy hiányzik a particionálási szakaszban létrehozott további partíciók csatolási pontja:

`root #` `mkdir --parents /mnt/gentoo`

Az mkdir paranccsal folytassa az előző lépések során létrehozott (egyéni) partíció(k)hoz szükséges további felcsatolási pontok létrehozását.

A felcsatolási pontok létrehozását követően ideje elérhetővé tenni a partíciókat a mount paranccsal.

Csatolja fel a gyökérpartíciót (a root partíciót):

`root #` `mount /dev/sda5 /mnt/gentoo`

Szükség szerint folytassa a további (egyéni) partíciók felcsatolását a fájlrendszerbe a mount paranccsal.

**Note**

Ha a /tmp/ könyvtárnak külön partíción kell lennie, akkor a felcsatolás után mindenképpen módosítsa a hozzá tartozó jogosultságokat:

`root #` `chmod 1777 /mnt/gentoo/tmp`

Ugyanez érvényes a /var/tmp könyvtárra is.

Később az utasításokban a proc fájlrendszer (a kernellel kapcsolatban álló virtuális interfész) és a többi kernel pszeudofájlrendszer lesz felcsatolva. Először viszont még a [Gentoo-stage fájlt](/wiki/Handbook:MIPS/Installation/Stage/hu "Handbook:MIPS/Installation/Stage/hu") ki kell csomagolnunk.

[← Hálózat beállítása](/wiki/Handbook:MIPS/Installation/Networking/hu "Previous part") [Kezdőlap](/wiki/Handbook:MIPS/hu "Handbook:MIPS/hu") [Stage fájl telepítése →](/wiki/Handbook:MIPS/Installation/Stage/hu "Next part")