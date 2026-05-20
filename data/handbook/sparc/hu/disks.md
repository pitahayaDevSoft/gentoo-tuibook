# Disks

Other languages:

- [Deutsch](/wiki/Handbook:SPARC/Installation/Disks/de "Handbuch:SPARC/Installation/Festplatten (100% translated)")
- [English](/wiki/Handbook:SPARC/Installation/Disks "Handbook:SPARC/Installation/Disks (100% translated)")
- [español](/wiki/Handbook:SPARC/Installation/Disks/es "Manual de Gentoo: SPARC/Instalación/Discos (100% translated)")
- [français](/wiki/Handbook:SPARC/Installation/Disks/fr "Handbook:SPARC/Installation/Disks/fr (100% translated)")
- [italiano](/wiki/Handbook:SPARC/Installation/Disks/it "Handbook:SPARC/Installation/Disks/it (50% translated)")
- magyar
- [polski](/wiki/Handbook:SPARC/Installation/Disks/pl "Handbook:SPARC/Installation/Disks (100% translated)")
- [português do Brasil](/wiki/Handbook:SPARC/Installation/Disks/pt-br "Handbook:SPARC/Installation/Disks/pt-br (50% translated)")
- [русский](/wiki/Handbook:SPARC/Installation/Disks/ru "Handbook:SPARC/Installation/Disks (100% translated)")
- [தமிழ்](/wiki/Handbook:SPARC/Installation/Disks/ta "கையேடு:SPARC/நிறுவல்/வட்டுக்கள் (100% translated)")
- [中文（中国大陆）‎](/wiki/Handbook:SPARC/Installation/Disks/zh-cn "手册：SPARC/安装/准备磁盘 (100% translated)")
- [日本語](/wiki/Handbook:SPARC/Installation/Disks/ja "ハンドブック:SPARC/インストール/ディスク (100% translated)")
- [한국어](/wiki/Handbook:SPARC/Installation/Disks/ko "Handbook:SPARC/Installation/Disks/ko (100% translated)")

[SPARC kézikönyv](/wiki/Handbook:SPARC/hu "Handbook:SPARC/hu")[A Gentoo Linux telepítése](/wiki/Handbook:SPARC/Full/Installation/hu "Handbook:SPARC/Full/Installation/hu")[A telepítésről](/wiki/Handbook:SPARC/Installation/About/hu "Handbook:SPARC/Installation/About/hu")[Telepítőképfájl kiválasztása](/wiki/Handbook:SPARC/Installation/Media/hu "Handbook:SPARC/Installation/Media/hu")[Hálózat beállítása](/wiki/Handbook:SPARC/Installation/Networking/hu "Handbook:SPARC/Installation/Networking/hu")Adathordozók előkészítése[Stage fájl](/wiki/Handbook:SPARC/Installation/Stage/hu "Handbook:SPARC/Installation/Stage/hu")[Alaprendszer telepítése](/wiki/Handbook:SPARC/Installation/Base/hu "Handbook:SPARC/Installation/Base/hu")[Kernel beállítása](/wiki/Handbook:SPARC/Installation/Kernel/hu "Handbook:SPARC/Installation/Kernel/hu")[Rendszer beállítása](/wiki/Handbook:SPARC/Installation/System/hu "Handbook:SPARC/Installation/System/hu")[Eszközök telepítése](/wiki/Handbook:SPARC/Installation/Tools/hu "Handbook:SPARC/Installation/Tools/hu")[Bootloader beállítása](/wiki/Handbook:SPARC/Installation/Bootloader/hu "Handbook:SPARC/Installation/Bootloader/hu")[Telepítés véglegesítése](/wiki/Handbook:SPARC/Installation/Finalizing/hu "Handbook:SPARC/Installation/Finalizing/hu")[Munka a Gentoo rendszerrel](/wiki/Handbook:SPARC/Full/Working/hu "Handbook:SPARC/Full/Working/hu")[Portage bemutatása](/wiki/Handbook:SPARC/Working/Portage/hu "Handbook:SPARC/Working/Portage/hu")[USE jelölőzászlók](/wiki/Handbook:SPARC/Working/USE/hu "Handbook:SPARC/Working/USE/hu")[Portage jellemzői](/wiki/Handbook:SPARC/Working/Features/hu "Handbook:SPARC/Working/Features/hu")[Init-szkript rendszer](/wiki/Handbook:SPARC/Working/Initscripts/hu "Handbook:SPARC/Working/Initscripts/hu")[Környezeti változók](/wiki/Handbook:SPARC/Working/EnvVar/hu "Handbook:SPARC/Working/EnvVar/hu")[Munka a Portage szoftvercsomag-kezelővel](/wiki/Handbook:SPARC/Full/Portage/hu "Handbook:SPARC/Full/Portage/hu")[Fájlok és könyvtárak](/wiki/Handbook:SPARC/Portage/Files/hu "Handbook:SPARC/Portage/Files/hu")[Változók](/wiki/Handbook:SPARC/Portage/Variables/hu "Handbook:SPARC/Portage/Variables/hu")[Szoftverágak keverése](/wiki/Handbook:SPARC/Portage/Branches/hu "Handbook:SPARC/Portage/Branches/hu")[További eszközök](/wiki/Handbook:SPARC/Portage/Tools/hu "Handbook:SPARC/Portage/Tools/hu")[Egyéni szoftvercsomag-tárolók](/wiki/Handbook:SPARC/Portage/CustomTree/hu "Handbook:SPARC/Portage/CustomTree/hu")[Fejlett funkciók](/wiki/Handbook:SPARC/Portage/Advanced/hu "Handbook:SPARC/Portage/Advanced/hu")[Hálózat beállítása OpenRC init-rendszeren](/wiki/Handbook:SPARC/Full/Networking/hu "Handbook:SPARC/Full/Networking/hu")[Munka elkezdése](/wiki/Handbook:SPARC/Networking/Introduction/hu "Handbook:SPARC/Networking/Introduction/hu")[Fejlett beállítások](/wiki/Handbook:SPARC/Networking/Advanced/hu "Handbook:SPARC/Networking/Advanced/hu")[Moduláris hálózat](/wiki/Handbook:SPARC/Networking/Modular/hu "Handbook:SPARC/Networking/Modular/hu")[Vezeték nélküli (Wi-Fi)](/wiki/Handbook:SPARC/Networking/Wireless/hu "Handbook:SPARC/Networking/Wireless/hu")[Funkcionalitás hozzáadása](/wiki/Handbook:SPARC/Networking/Extending/hu "Handbook:SPARC/Networking/Extending/hu")[Dinamikus menedzsment](/wiki/Handbook:SPARC/Networking/Dynamic/hu "Handbook:SPARC/Networking/Dynamic/hu")

## Contents

- [1Bevezetés a blokktípusú eszközökbe](#Bevezet.C3.A9s_a_blokkt.C3.ADpus.C3.BA_eszk.C3.B6z.C3.B6kbe)
  - [1.1Blokkeszközök](#Blokkeszk.C3.B6z.C3.B6k)
  - [1.2Partíciós táblázatok](#Part.C3.ADci.C3.B3s_t.C3.A1bl.C3.A1zatok)
    - [1.2.1GUID Partition Table (GPT)](#GUID_Partition_Table_.28GPT.29)
    - [1.2.2Sun partíciós táblázat](#Sun_part.C3.ADci.C3.B3s_t.C3.A1bl.C3.A1zat)
  - [1.3Alapértelmezett partíciós séma](#Alap.C3.A9rtelmezett_part.C3.ADci.C3.B3s_s.C3.A9ma)
    - [1.3.1GPT partíciós séma](#GPT_part.C3.ADci.C3.B3s_s.C3.A9ma)
    - [1.3.2Sun formázott partíciós séma](#Sun_form.C3.A1zott_part.C3.ADci.C3.B3s_s.C3.A9ma)
- [2Adathordozó particionálása GPT táblázattal](#Adathordoz.C3.B3_particion.C3.A1l.C3.A1sa_GPT_t.C3.A1bl.C3.A1zattal)
  - [2.1Jelenlegi partíciós elrendezés megtekintése](#Jelenlegi_part.C3.ADci.C3.B3s_elrendez.C3.A9s_megtekint.C3.A9se)
  - [2.2Új partíciós tábla (lemezcímke) létrehozása és az összes meglévő partíció eltávolítása](#.C3.9Aj_part.C3.ADci.C3.B3s_t.C3.A1bla_.28lemezc.C3.ADmke.29_l.C3.A9trehoz.C3.A1sa_.C3.A9s_az_.C3.B6sszes_megl.C3.A9v.C5.91_part.C3.ADci.C3.B3_elt.C3.A1vol.C3.ADt.C3.A1sa)
  - [2.3BIOS boot partíció létrehozása](#BIOS_boot_part.C3.ADci.C3.B3_l.C3.A9trehoz.C3.A1sa)
  - [2.4Swap partíció létrehozása](#Swap_part.C3.ADci.C3.B3_l.C3.A9trehoz.C3.A1sa)
  - [2.5Gyökérpartíció létrehozása](#Gy.C3.B6k.C3.A9rpart.C3.ADci.C3.B3_l.C3.A9trehoz.C3.A1sa)
  - [2.6Partíciók elrendezésének a mentése](#Part.C3.ADci.C3.B3k_elrendez.C3.A9s.C3.A9nek_a_ment.C3.A9se)
- [3Adathordozó particionálása Sun partíciós táblázattal](#Adathordoz.C3.B3_particion.C3.A1l.C3.A1sa_Sun_part.C3.ADci.C3.B3s_t.C3.A1bl.C3.A1zattal)
  - [3.1Jelenlegi partíciós elrendezés megtekintése](#Jelenlegi_part.C3.ADci.C3.B3s_elrendez.C3.A9s_megtekint.C3.A9se_2)
  - [3.2Új adathordozó táblázat (lemezcímke) létrehozása / Az összes partíció eltávolítása](#.C3.9Aj_adathordoz.C3.B3_t.C3.A1bl.C3.A1zat_.28lemezc.C3.ADmke.29_l.C3.A9trehoz.C3.A1sa_.2F_Az_.C3.B6sszes_part.C3.ADci.C3.B3_elt.C3.A1vol.C3.ADt.C3.A1sa)
  - [3.3Egész adathordozó partíciójának a létrehozása](#Eg.C3.A9sz_adathordoz.C3.B3_part.C3.ADci.C3.B3j.C3.A1nak_a_l.C3.A9trehoz.C3.A1sa)
  - [3.4Gyökérpartíció létrehozása](#Gy.C3.B6k.C3.A9rpart.C3.ADci.C3.B3_l.C3.A9trehoz.C3.A1sa_2)
  - [3.5Swap partíció létrehozása](#Swap_part.C3.ADci.C3.B3_l.C3.A9trehoz.C3.A1sa_2)
  - [3.6Partíciók elrendezésének a mentése](#Part.C3.ADci.C3.B3k_elrendez.C3.A9s.C3.A9nek_a_ment.C3.A9se_2)
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

### Partíciós táblázatok

Bár elméletileg lehetséges egy nyers, particionálatlan adathordozót használni egy Linux rendszer tárolására (például btrfs RAID létrehozásakor), a gyakorlatban ez szinte soha nem történik meg. Ehelyett az adathordozó blokkeszközöket kisebb, kezelhetőbb blokkeszközökre osztják fel. **sparc** rendszereken ezeket partícióknak nevezik. Jelenleg két standard partíciótechnológia van használatban: Sun és GPT, utóbbit csak a közelmúltban megjelent firmware-rel rendelkező újabb rendszerek támogatják.

#### GUID Partition Table (GPT)

A _GUID Partition Table (GPT)_ beállítás (más néven GPT adathordozó táblázat -lemezcímke-) 64 bites azonosítókat használ a partíciókhoz. Az a hely az adathordozón, ahol a partíció információit tárolja. Sokkal nagyobb, mint az MBR partíciós táblázat (DOS lemezcímke) 512 bájtja, ami gyakorlatilag korlátlan számú partíciót tesz lehetővé egy GPT adathordozón. Emellett a partíció _mérete_ sokkal nagyobb határértékkel rendelkezik (közel 8 ZiB - igen, zebibájt).

A GPT előnyei közé tartozik a hibadetektálási és redundancia mechanizmusok használata. CRC32 ellenőrzőösszegeket tartalmaz, amelyek segítenek hibákat észlelni a fejlécben és a partíciós táblázatokban, továbbá egy biztonsági GPT-t helyez el az adathordozó végén. Ez a biztonsági táblázat használható a fő GPT helyreállítására, ha az az adathordozó eleje közelében sérül.

A GPT táblázatot csak az Oracle SPARC gépek T4 generációja vagy újabb verziói támogatják. Emellett csak bizonyos, újabb firmware-ek tartalmaznak GPT támogatást. Számos módszer létezik annak ellenőrzésére, hogy elérhető-e a GPT támogatás.

Az OBP promptból hajtsa végre a következőket:

`{0} ok` `cd /packages/disk-label`

`{0} ok` `.properties`

gpt

supported-labels gpt

```
                   sun
                   mbr

```

name disk-label

Ha a kimenet tartalmazza a gpt jelölést, akkor a GPT támogatás elérhető. Alternatívaként ez meghatározható a telepítési adathordozóról is, az OBP-be való belépés nélkül. Használja a prtconf parancsot a [sys-apps/sparc-utils](https://packages.gentoo.org/packages/sys-apps/sparc-utils) szoftvercsomagból, hogy ezt az információt a felhasználói térből elérje:

`root #` `prtconf -pv | grep -c gpt`

Vagy ellenőrizze, hogy létezik-e a /sys/firmware/devicetree/base/packages/disk-label/gpt fájl. Ha egyik módszer sem jár sikerrel, akkor firmware frissítés szükséges a GPT támogatásához.

#### Sun partíciós táblázat

Az Oracle által nem gyártott rendszerek, a T3 vagy korábbi rendszerek, illetve a korábbi firmware-rel működő rendszerek a Sun partíciós táblázat típust kell, hogy használják.

A harmadik partíció a Sun rendszereken egy speciális "teljes adathordozó" szeletként van fenntartva. Ez a partíció **nem** tartalmazhat fájlrendszert.

A DOS partíciós séma használatához szokott felhasználók vegyék figyelembe, hogy a Sun partíciós táblázatok nem rendelkeznek "elsődleges" és "kiterjesztett" partíciókkal. Ehelyett meghajtónként legfeljebb nyolc partíció érhető el, amelyek közül a harmadik fenntartott.

A kézikönyv szerzői azt javasolják, hogy a Gentoo telepítésekhez lehetőség szerint mindig [GPT](#GPT) partíciót használjanak.

### Alapértelmezett partíciós séma

A GPT és a Sun partíciós táblázatok közötti szükséges partíciós elrendezés különbségei miatt egyetlen partíciós séma nem elegendő minden lehetséges rendszerkövetelmény támogatásához. Az alábbiakban bemutatunk néhány példasémát.

#### GPT partíciós séma

A következő partíciós séma kerül bemutatásra példaként GPT-formázott adathordozók számára:

Partíció
Fájlrendszer
Méret
Csatolási pont
Leírás
/dev/sda1(Semmi)
2 MiB
Semmi
[BIOS boot partíció](/wiki/Handbook:X86/Blocks/Disks#What_is_the_BIOS_boot_partition.3F.2Fhu "Handbook:X86/Blocks/Disks")/dev/sda2(swap)
RAM mérete \* 2
Semmi
Swap partíció
/dev/sda3ext4
Adathordozó fennmaradó része
/Root partíció

#### Sun formázott partíciós séma

A következő particionálási séma lesz példaként a Sun által formázott adathordozók számára:

Partíció
Fájlrendszer
Méret
Csatolási pont
Leírás
/dev/sda1ext4
Adathordozó mérete mínusz a swap mérete
/Root partíció
/dev/sda2(swap)
RAM mérete \* 2
Semmi
Swap partíció
/dev/sda3(none)
Teljes adathordozó
Semmi
Teljes adathordozó partíció. Kötelező a Sun partíciós táblázatot használó adathordozók esetében.

**Important**

Az OBP 3-as vagy korábbi verzióját használó SPARC rendszereknek további korlátozásai vannak a partíciók elrendezésében. A gyökérpartíciónak az adathordozón az első partíciónak kell lennie, és nem lehet nagyobb, mint 2 GiB méretű. Emiatt az ilyen rendszereknek megfelelő méretű partíciókra lesz szükségük a felső szintű könyvtárak, például /usr, /var, /home, és más könyvtárak számára, amelyek valószínűleg meghaladnák a gyökérpartíció limitjét. Ezek a rendszerek valószínűleg a Sun partíciós táblázat típust igénylik, ezért ne felejtse el belefoglalni a teljes adathordozó partíciót.

## Adathordozó particionálása GPT táblázattal

A következő részben bemutatjuk, hogy miként hozható létre a példaként említett partíciós elrendezés egy GPT telepítéshez a fdisk használatával. Az említett partíciós elrendezést korábban ismertettük:

Partíció
Leírás
/dev/sda1Boot partíció
/dev/sda2Swap partíció
/dev/sda3Root partíció

Módosítsa a partíciós elrendezést a rendszer igényeinek megfelelően.

### Jelenlegi partíciós elrendezés megtekintése

Az fdisk egy népszerű és hatékony szoftver az adathordozók partíciókra történő felosztására. Indítsa el a fdisk parancsot, paraméterként megadva az adathordozót (ebben a példában most a /dev/sda van használva):

`root #` `fdisk /dev/sda`

Használja a `p` billentyűgombot az adathordozó jelenlegi partíciós beállításának a megjelenítéséhez:

`Command (m for help):` `p`

```
Disk /dev/sda: 14.57 GiB, 15640625152 bytes, 30548096 sectors
Disk model: USB Flash Disk
Units: sectors of 1 * 512 = 512 bytes
Sector size (logical/physical): 512 bytes / 512 bytes
I/O size (minimum/optimal): 512 bytes / 512 bytes
Disklabel type: gpt
Disk identifier: 9850A2C2-76C4-FC47-9F0B-DA60449D2413

Device     Start      End  Sectors  Size Type
/dev/sda1   2048 30547967 30545920 14.6G Linux filesystem

```

### Új partíciós tábla (lemezcímke) létrehozása és az összes meglévő partíció eltávolítása

Nyomja le a `g` billentyűgombot, hogy egy új GPT partíciós táblázatot hozzon létre az adathordozón. Ez az összes meglévő partíciót eltávolítja.

`Command (m for help):` `g`

```
Created a new GPT disklabel (GUID: 9850A2C2-76C4-FC47-9F0B-DA60449D2413).

```

Egy meglévő GPT partíciós táblázat esetén (részletek a fentebb megjelenített `p` kimenetnél), fontolja meg az egyes meglévő partíciók egyenkénti eltávolítását az adathordozóról. Nyomja meg a `d` billentyűgombot egy partíció törléséhez. Például egy meglévő /dev/sda1 törléséhez:

`Command (m for help):` `d`

```
Selected partition 1
Partition 1 has been deleted.

```

A partíció törlése már ütemezve lett. Nem fog megjelenni, amikor a partíciók listáját megjelenítjük ( `p`, de nem kerül törlésre, amíg a változtatásokat el nem mentik. Ez lehetővé teszi, hogy a felhasználók megszakítsák a műveletet, ha hibát követtek el. Ebben az esetben azonnal nyomja le a `q` billentyűgombot, majd nyomja meg a `Enter` billentyűgombot, és a partíció nem lesz törölve.

Ismételten nyomja le a `p` billentyűgombot, hogy megjelenítse a partíciók listáját, majd nyomja le a `d` billentyűgombot és a törölni kívánt partíció számát. Végül a partíciós táblázat üres lesz:

`Command (m for help):` `p`

```
Disk /dev/sda: 14.57 GiB, 15640625152 bytes, 30548096 sectors
Disk model: USB Flash Disk
Units: sectors of 1 * 512 = 512 bytes
Sector size (logical/physical): 512 bytes / 512 bytes
I/O size (minimum/optimal): 512 bytes / 512 bytes
Disklabel type: gpt
Disk identifier: 9850A2C2-76C4-FC47-9F0B-DA60449D2413

```

Most, hogy a RAM memóriában lévő partíciós táblázat üres, készen állunk a partíciók létrehozására.

### BIOS boot partíció létrehozása

Először hozza létre a BIOS boot partíciót. Nyomja le a `n` billentyűgombot egy új partíció létrehozásához, majd a `1` billentyűgombot nyomja le az első partíció kiválasztásához. Amikor az első szektorról kérdezi a rendszer, győződjön meg róla, hogy az 2048-tól kezdődik (ami szükséges lehet az boot loader számára), majd nyomja meg a `Enter` billentyűgombot. Amikor a rendszer az utolsó szektorról kérdez, írja be a +2M értéket, hogy egy 2 megabyte méretű partíciót hozzon létre.

`Command (m for help):` `n`

```
Partition number (1-128, default 1):
First sector (2048-30548062, default 2048):
Last sector, +/-sectors or +/-size{K,M,G,T,P} (2048-30548062, default 30547967): +2M

Created a new partition 1 of type 'Linux filesystem' and of size 2 MiB.

```

Jelölje meg a partíciót BIOS boot partícióként:

`Command (m for help):` `t`

```
Selected partition 1
Partition type or alias (type L to list all): 4
Changed type of partition 'Linux filesystem' to 'BIOS boot'.

```

### Swap partíció létrehozása

Ezután a swap partíció létrehozásához nyomja le a `n` billentyűt egy új partíció létrehozásához, majd nyomja meg a `2` billentyűgombot a második partíció létrehozásához, /dev/sda2. Amikor az első szektorról kérdezi a rendszer, nyomja meg a `Enter` billentyűgombot. Amikor az utolsó szektorról kérdez, írja be a +4G értéket (vagy bármilyen más szükséges méretet a swap terület számára), hogy egy 4 GiB méretű partíciót hozzon létre.

`Command (m for help):` `n`

```
Partition number (2-128, default 2):
First sector (6144-30548062, default 6144):
Last sector, +/-sectors or +/-size{K,M,G,T,P} (6144-30548062, default 30547967): +4G

Created a new partition 2 of type 'Linux filesystem' and of size 4 GiB.

```

Miután mindez megtörtént, nyomja meg a `t` billentyűgombot a partíció típusának beállításához, majd `2` a most létrehozott partíció kiválasztásához, és ezután írja be a _19_ értéket, hogy a partíció típusát "Linux Swap" típusúra állítsa be.

`Command (m for help):` `t`

```
Partition number (1,2, default 2): 2
Partition type (type L to list all types): 19

Changed type of partition 'Linux filesystem' to 'Linux swap'.

```

### Gyökérpartíció létrehozása

Végül a root partíció létrehozásához nyomja le a `n` billentyűgombot egy új partíció létrehozásához. Ezután nyomja le a `3` billentyűt a harmadik partíció létrehozásához, /dev/sda3. Amikor az első szektorról kérdezi a rendszer, nyomja meg a `Enter` billentyűgombot. Amikor az utolsó szektorról kérdez, nyomja meg a `Enter` billentyűgombot, hogy egy partíciót hozzon létre, amely az adathordozón fennmaradó még szabad helyet teljesen kitölti. Miután befejezte ezeket a lépéseket, a `p` parancs billentyűgomb megnyomásával egy partíciós táblázatot kell kapnia, amely hasonlóan kell, hogy kinézzen:

`Command (m for help):` `p`

```
Disk /dev/sda: 14.57 GiB, 15640625152 bytes, 30548096 sectors
Disk model: USB Flash Disk
Units: sectors of 1 * 512 = 512 bytes
Sector size (logical/physical): 512 bytes / 512 bytes
I/O size (minimum/optimal): 512 bytes / 512 bytes
Disklabel type: gpt
Disk identifier: 9850A2C2-76C4-FC47-9F0B-DA60449D2413

Device       Start      End  Sectors  Size Type
/dev/sda1     2048     6143     4096    2M BIOS boot
/dev/sda2     6144  8394751  8388608    4G Linux swap
/dev/sda3  8394752 30547967 22153216 10.6G Linux filesystem

```

### Partíciók elrendezésének a mentése

A partíciók elrendezésének a mentésé érdekében és egyúttal az fdisk programból való kilépéséhez nyomja meg a `w` billentyűgombot.

`Command (m for help):` `w`

A partíciók létrehozása után most már ideje a fájlrendszereket létrehozni a partíciókon.

## Adathordozó particionálása Sun partíciós táblázattal

A következő részek bemutatják, hogy miként lehet létrehozni a példaként megadott partíciós elrendezést egy Sun partíciós táblázattal történő telepítéshez az fdisk segítségével. (A példaként megadott partíciós elrendezés korábban volt említve). :

Partíció
Leírás
/dev/sda1Root partíció
/dev/sda2Swap partíció
/dev/sda3Teljes adathordozó partíció

Módosítsa a partíciós elrendezést az Ön személyes preferenciáinak megfelelően. Ha egy OBP 3-as vagy korábbi verzióját használó rendszer számára történik a particionálás, akkor győződjön meg arról, hogy a root partíció mérete kevesebb, mint 2GiB, és ezen felül hozzon létre további fájlrendszerek számára partíciókat /dev/sda4 partíciótól kezdődően.

### Jelenlegi partíciós elrendezés megtekintése

Az fdisk szoftver egy népszerű és hatékony eszköz az adathordozó partíciókra történő felosztására. Indítsa el a fdisk parancsot, paraméterként a particionálandó adathordozót megadva (ebben a példában most a /dev/sda van használva):

`root #` `fdisk /dev/sda`

Használja a `p` billentyűgombot az adathordozó jelenlegi partíciós elrendezésének a megjelenítéséhez:

`Command (m for help):` `p`

```
Disk model: USB Flash Disk
Geometry: 64 heads, 32 sectors/track, 14916 cylinders
Units: sectors of 1 * 512 = 512 bytes
Sector size (logical/physical): 512 bytes / 512 bytes
I/O size (minimum/optimal): 512 bytes / 512 bytes
Disklabel type: sun

Device        Start      End  Sectors  Size Id Type         Flags
/dev/sda1         0 30445567 30445568 14.5G 83 Linux native
/dev/sda2  30445568 30547967   102400   50M 82 Linux swap      u
/dev/sda3         0 30547967 30547968 14.6G  5 Whole disk

```

### Új adathordozó táblázat (lemezcímke) létrehozása / Az összes partíció eltávolítása

Nyomja le az `s` billentyűgombot, hogy egy új Sun adathordozó táblázatot hozzon létre az adathordozón. Ez a művelet egyben az összes meglévő partíciót eltávolítja.

`Command (m for help):` `s`

```
Created a new partition 1 of type 'Linux native' and of size 14.5 GiB.
Created a new partition 2 of type 'Linux swap' and of size 50 MiB.
Created a new partition 3 of type 'Whole disk' and of size 14.6 GiB.
Created a new Sun disklabel.

```

Egy meglévő Sun adathordozó táblázat (lemezcímke) esetén (részletek a fentebb megjelenített `p` kimeneten), fontolja meg az egyes meglévő partíciók egyenkénti eltávolítását az adathordozóról. Nyomja meg a `d` billentyűgombot egy partíció törléséhez. Például egy meglévő /dev/sda1 törléséhez:

`Command (m for help):` `d`

```
Partition number (1-3, default 3): 1

Partition 1 has been deleted.

```

A partíció törlése ütemezve lett. Innentől a partíció nem fog megjelenni, amikor a partíciók listáját megjelenítjük ( `p`), de valójában nem kerül törlésre, amíg a változtatásokat el nem mentjük. Ez lehetővé teszi, hogy a felhasználók megszakítsák a műveletet, ha hibát követtek el. Ebben az esetben azonnal nyomja le a `q` billentyűgombot, majd nyomja meg a `Enter` billentyűgombot, és a partíció nem lesz letörölve.

Ismételten nyomja meg a `p` billentyűgombot a partíciók listájának a megjelenítéséhez, majd nyomja meg a `d` billentyűgombot és a törölni kívánt partíciónak a számát. Végül a partíciós táblázat teljesen üressé válik.

`Command (m for help):` `p`

```
Disk /dev/sda: 14.57 GiB, 15640625152 bytes, 30548096 sectors
Disk model: USB Flash Disk
Geometry: 64 heads, 32 sectors/track, 14916 cylinders
Units: sectors of 1 * 512 = 512 bytes
Sector size (logical/physical): 512 bytes / 512 bytes
I/O size (minimum/optimal): 512 bytes / 512 bytes
Disklabel type: sun

```

Most, hogy a memóriában lévő partíciós tábla üres, készen állunk a partíciók létrehozására.

### Egész adathordozó partíciójának a létrehozása

Először hozza létre az egész adathordozó partícióját. Nyomja meg a `n` billentyűgombot egy új partíció létrehozásához, majd nyomja meg a `3` billentyűt a harmadik partíció kiválasztásához. Amikor az első szektorról kérdez a rendszer, akkor nyomja meg a `Enter` billentyűgombot. Amikor az utolsó szektorról kérdez, akkor nyomja meg a `Enter` billentyűgombot, hogy egy partíciót hozzon létre, amely az adathordozó teljes fennmaradó helyét elfoglalja.

`Command (m for help):` `n`

```
Partition number (1-8, default 1): 3

It is highly recommended that the third partition covers the whole disk and is of type `Whole disk'
First sector (0-30547968, default 0):
Last sector or +/-sectors or +/-size{K,M,G,T,P} (0-30547968, default 30547968):

Created a new partition 3 of type 'Whole disk' and of size 14.6 GiB.

```

A fdisk automatikusan beállítja az ilyen partíció típusát 'Whole disk' típusra, így nem szükséges azt kifejezetten megadni.

### Gyökérpartíció létrehozása

Ezután a gyökérpartíció (root partíció) létrehozásához nyomja meg a `n` billentyűgombot egy új partíció létrehozásához. Ezután nyomja meg a `1` billentyűgombot az első partíció létrehozásához, /dev/sda1. Amikor az első szektorról kérdez a rendszer, nyomja meg a `Enter` billentyűgombot. Amikor az utolsó szektorról kérdez, írja be a -4G értéket (vagy bármilyen más helyet, amely a nem root partíciókhoz szükséges). Miután befejezte ezeket a lépéseket, a `p` parancs beírásával egy partíciós táblázatot kell kapnia, amely hasonlóan néz ki:

`Command (m for help):` `p`

```
Disk /dev/sda: 14.57 GiB, 15640625152 bytes, 30548096 sectors
Disk model: USB Flash Disk
Geometry: 64 heads, 32 sectors/track, 14916 cylinders
Units: sectors of 1 * 512 = 512 bytes
Sector size (logical/physical): 512 bytes / 512 bytes
I/O size (minimum/optimal): 512 bytes / 512 bytes
Disklabel type: sun

Device     Start      End  Sectors  Size Id Type         Flags
/dev/sda1      0 22159359 22159360 10.6G 83 Linux native
/dev/sda3      0 30547967 30547968 14.6G  5 Whole disk

```

### Swap partíció létrehozása

Végül a swap partíció létrehozásához nyomja le az `n` billentyűgombot egy új partíció létrehozásához, majd nyomja le a `2` billentyűgombot a második partíció létrehozásához, /dev/sda2. Amikor az első szektorról kérdez a rendszer, nyomja meg a `Enter` billentyűgombot. Amikor az utolsó szektorról kérdez, nyomja meg a `Enter` billentyűgombot, hogy az összes fennmaradó helyet felhasználja az adathordozón.

`Command (m for help):` `n`

```
Partition number (2,4-8, default 2):
First sector (22159360-30547968, default 22159360):
Last sector or +/-sectors or +/-size{K,M,G,T,P} (22159360-30547968, default 30547968):

Created a new partition 2 of type 'Linux native' and of size 4 GiB.

```

Miután mindez megtörtént, nyomja le a `t` billentyűgombot a partíció típusának beállításához, majd nyomja le a `2` billentyűgombot az imént létrehozott partíció kiválasztásához. Ezután írja be a _82_ értéket, hogy a partíció típusát "Linux Swap" típusúra állítsa be.

`Command (m for help):` `t`

```
Partition number (1-3, default 3): 2
Hex code (type L to list all codes): 82

Changed type of partition 'Linux native' to 'Linux swap'.

```

Miután befejezte ezeket a lépéseket, a `p` parancs beírásával egy partíciós táblázatot kell kapnia, amely hasonlóan néz ki:

`Command (m for help):` `p`

```
Disk /dev/sda: 14.57 GiB, 15640625152 bytes, 30548096 sectors
Disk model: USB Flash Disk
Geometry: 64 heads, 32 sectors/track, 14916 cylinders
Units: sectors of 1 * 512 = 512 bytes
Sector size (logical/physical): 512 bytes / 512 bytes
I/O size (minimum/optimal): 512 bytes / 512 bytes
Disklabel type: sun

Device        Start      End  Sectors  Size Id Type         Flags
/dev/sda1         0 22159359 22159360 10.6G 83 Linux native
/dev/sda2  22159360 30547967  8388608    4G 82 Linux swap      u
/dev/sda3         0 30547967 30547968 14.6G  5 Whole disk

```

### Partíciók elrendezésének a mentése

A partíciók elrendezésének a mentéséhez és egyben a fdisk szoftverből való kilépéséhez nyomja meg a `w` billentyűgombot.

`Command (m for help):` `w`

A partíciók létrehozása után elérkezett az idő, hogy fájlrendszereket hozzunk létre rajtuk.

## Fájlrendszerek létrehozása

**Warning**

SSD vagy NVMe adathordozó használatakor bölcs dolog ellenőrizni a firmware-frissítéseket. Különösen egyes Intel SSD adathordozók (600p és 6000p) firmware-frissítést igényelnek az XFS I/O használati minták által okozott [lehetséges adatsérülések miatt](https://bugzilla.redhat.com/show_bug.cgi?id=1402533). A probléma a firmware szintjén van, és nem az XFS fájlrendszer hibája. A smartctl segédprogram segíthet az adathordozzó eszköz modelljének és firmware-verziójának ellenőrzésében.

### Bevezetés

Most, hogy a partíciók elkészültek, ideje fájlrendszert helyezni rájuk. A következő részben a Linux által támogatott különféle fájlrendszereket ismertetjük. Azok az olvasók, akik már tudják, hogy melyik fájlrendszert fogják használni, folytathatják a [Fájlrendszer rárakása egy partícióra](/wiki/Handbook:SPARC/Installation/Disks/hu#Applying_a_filesystem_to_a_partition "Handbook:SPARC/Installation/Disks/hu") című bekezdéssel. A többi felhasználónak érdemes továbbolvasniuk, hogy megismerjék az alkalmazható fájlrendszereket...

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

Például ahhoz, hogy a gyökérpartíció (tehát a root partíció) (/dev/sda1) fájlrendszertípusa xfs legyen, ahogy a partíciókészítés példa szerkezetében is szerepel, Önnek a következő parancsokat kell futtatnia:

`root #` `mkfs.xfs /dev/sda1`

#### Örökölt BIOS rendszerindító partíciónak a fájlrendszere

A régebbi, MBR/DOS adathordozó partíciós táblázattal ellátott BIOS-on keresztül induló rendszerek bármilyen, a rendszerbetöltő által támogatott fájlrendszert használhatnak.

Például XFS fájlrendszerrel történő formázáshoz futtassa a következő parancsot:

`root #` `mkfs.xfs `

#### Kicsi ext4 partíciók

Ha Ön egy kicsi méretű partíción (kevesebb, mint 8 GiB) ext4 fájlrendszert szeretne használ, akkor a fájlrendszert a megfelelő beállításokkal kell létrehozni, hogy az elegendő inode-okat foglalhasson le. Ezt a `-T small` opcióval lehet megadni:

`root #` `mkfs.ext4 -T small /dev/<device>`

Ezzel megnégyszerezi az adott fájlrendszer inode-jainak a számát, mivel a "bytes-per-inode" 16 kB-onként 4 kB-ra csökken.

### A swap (lapozásra használt) partíció aktiválása

Az mkswap parancs szolgál a swap partíciók létrehozásához:

`root #` `mkswap /dev/sda2`

**Note**

Innentől folytatható az a rendszertelepítés, amely korábban el lett kezdve, de a telepítési folyamat nem let végig befejezve. Használja ezt a hivatkozást állandó hivatkozásként: [A telepítés folytatása itt kezdődik](/wiki/Handbook:SPARC/Installation/Disks/hu#Resumed_installations_start_here "Handbook:SPARC/Installation/Disks/hu").

A swap partíciót aktiválni is kell. Használja a swapon parancsot:

`root #` `swapon /dev/sda2`

Ez az 'aktiválás' azért szükséges most, mert a swap partíciót újonnan hozzuk létre a Live ISO telepítőkörnyezetben. A rendszer újraindítása után mindaddig, amíg a swap partíció megfelelően van definiálva az fstab fájlban vagy más csatolási mechanizmusban, a swap terület automatikusan fog aktiválódni.

## Gyökérpartíció (root partíció) felcsatolása

Előfordulhat, hogy bizonyos Live ISO telepítőkörnyezetekből hiányzik a javasolt csatolási pont a Gentoo gyökérpartíciójához (/mnt/gentoo), vagy hiányzik a particionálási szakaszban létrehozott további partíciók csatolási pontja:

`root #` `mkdir --parents /mnt/gentoo`

Az mkdir paranccsal folytassa az előző lépések során létrehozott (egyéni) partíció(k)hoz szükséges további felcsatolási pontok létrehozását.

A felcsatolási pontok létrehozását követően ideje elérhetővé tenni a partíciókat a mount paranccsal.

Csatolja fel a gyökérpartíciót (a root partíciót):

`root #` `mount /dev/sda1 /mnt/gentoo`

Szükség szerint folytassa a további (egyéni) partíciók felcsatolását a fájlrendszerbe a mount paranccsal.

**Note**

Ha a /tmp/ könyvtárnak külön partíción kell lennie, akkor a felcsatolás után mindenképpen módosítsa a hozzá tartozó jogosultságokat:

`root #` `chmod 1777 /mnt/gentoo/tmp`

Ugyanez érvényes a /var/tmp könyvtárra is.

Később az utasításokban a proc fájlrendszer (a kernellel kapcsolatban álló virtuális interfész) és a többi kernel pszeudofájlrendszer lesz felcsatolva. Először viszont még a [Gentoo-stage fájlt](/wiki/Handbook:SPARC/Installation/Stage/hu "Handbook:SPARC/Installation/Stage/hu") ki kell csomagolnunk.

[← Hálózat beállítása](/wiki/Handbook:SPARC/Installation/Networking/hu "Previous part") [Kezdőlap](/wiki/Handbook:SPARC/hu "Handbook:SPARC/hu") [Stage fájl telepítése →](/wiki/Handbook:SPARC/Installation/Stage/hu "Next part")