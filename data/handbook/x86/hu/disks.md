# Disks

Other languages:

- [Deutsch](/wiki/Handbook:X86/Installation/Disks/de "Handbuch:X86/Installation/Festplatten (100% translated)")
- [English](/wiki/Handbook:X86/Installation/Disks "Handbook:X86/Installation/Disks (100% translated)")
- [Türkçe](/wiki/Handbook:X86/Installation/Disks/tr "Handbook:X86/Installation/Disks/tr (0% translated)")
- [español](/wiki/Handbook:X86/Installation/Disks/es "Manual de Gentoo: X86/Instalación/Discos (100% translated)")
- [français](/wiki/Handbook:X86/Installation/Disks/fr "Handbook:X86/Installation/Disks/fr (100% translated)")
- [italiano](/wiki/Handbook:X86/Installation/Disks/it "Handbook:X86/Installation/Disks (100% translated)")
- magyar
- [polski](/wiki/Handbook:X86/Installation/Disks/pl "Handbook:X86/Installation/Disks (100% translated)")
- [português do Brasil](/wiki/Handbook:X86/Installation/Disks/pt-br "Manual:X86/Instalação/Discos (100% translated)")
- [čeština](/wiki/Handbook:X86/Installation/Disks/cs "Handbook:X86/Installation/Disks/cs (50% translated)")
- [русский](/wiki/Handbook:X86/Installation/Disks/ru "Handbook:X86/Installation/Disks (100% translated)")
- [தமிழ்](/wiki/Handbook:X86/Installation/Disks/ta "கையேடு:X86/நிறுவல்/வட்டுக்கள் (100% translated)")
- [中文（中国大陆）‎](/wiki/Handbook:X86/Installation/Disks/zh-cn "手册：X86/安装/准备磁盘 (100% translated)")
- [日本語](/wiki/Handbook:X86/Installation/Disks/ja "ハンドブック:X86/インストール/ディスク (100% translated)")
- [한국어](/wiki/Handbook:X86/Installation/Disks/ko "Handbook:X86/Installation/Disks/ko (100% translated)")

[X86 kézikönyv](/wiki/Handbook:X86/hu "Handbook:X86/hu")[A Gentoo Linux telepítése](/wiki/Handbook:X86/Full/Installation/hu "Handbook:X86/Full/Installation/hu")[A telepítésről](/wiki/Handbook:X86/Installation/About/hu "Handbook:X86/Installation/About/hu")[Telepítőképfájl kiválasztása](/wiki/Handbook:X86/Installation/Media/hu "Handbook:X86/Installation/Media/hu")[Hálózat beállítása](/wiki/Handbook:X86/Installation/Networking/hu "Handbook:X86/Installation/Networking/hu")Adathordozók előkészítése[Stage fájl](/wiki/Handbook:X86/Installation/Stage/hu "Handbook:X86/Installation/Stage/hu")[Alaprendszer telepítése](/wiki/Handbook:X86/Installation/Base/hu "Handbook:X86/Installation/Base/hu")[Kernel beállítása](/wiki/Handbook:X86/Installation/Kernel/hu "Handbook:X86/Installation/Kernel/hu")[Rendszer beállítása](/wiki/Handbook:X86/Installation/System/hu "Handbook:X86/Installation/System/hu")[Eszközök telepítése](/wiki/Handbook:X86/Installation/Tools/hu "Handbook:X86/Installation/Tools/hu")[Bootloader beállítása](/wiki/Handbook:X86/Installation/Bootloader/hu "Handbook:X86/Installation/Bootloader/hu")[Telepítés véglegesítése](/wiki/Handbook:X86/Installation/Finalizing/hu "Handbook:X86/Installation/Finalizing/hu")[Munka a Gentoo rendszerrel](/wiki/Handbook:X86/Full/Working/hu "Handbook:X86/Full/Working/hu")[Portage bemutatása](/wiki/Handbook:X86/Working/Portage/hu "Handbook:X86/Working/Portage/hu")[USE jelölőzászlók](/wiki/Handbook:X86/Working/USE/hu "Handbook:X86/Working/USE/hu")[Portage jellemzői](/wiki/Handbook:X86/Working/Features/hu "Handbook:X86/Working/Features/hu")[Init-szkript rendszer](/wiki/Handbook:X86/Working/Initscripts/hu "Handbook:X86/Working/Initscripts/hu")[Környezeti változók](/wiki/Handbook:X86/Working/EnvVar/hu "Handbook:X86/Working/EnvVar/hu")[Munka a Portage szoftvercsomag-kezelővel](/wiki/Handbook:X86/Full/Portage/hu "Handbook:X86/Full/Portage/hu")[Fájlok és könyvtárak](/wiki/Handbook:X86/Portage/Files/hu "Handbook:X86/Portage/Files/hu")[Változók](/wiki/Handbook:X86/Portage/Variables/hu "Handbook:X86/Portage/Variables/hu")[Szoftverágak keverése](/wiki/Handbook:X86/Portage/Branches/hu "Handbook:X86/Portage/Branches/hu")[További eszközök](/wiki/Handbook:X86/Portage/Tools/hu "Handbook:X86/Portage/Tools/hu")[Egyéni szoftvercsomag-tárolók](/wiki/Handbook:X86/Portage/CustomTree/hu "Handbook:X86/Portage/CustomTree/hu")[Fejlett funkciók](/wiki/Handbook:X86/Portage/Advanced/hu "Handbook:X86/Portage/Advanced/hu")[Hálózat beállítása OpenRC init-rendszeren](/wiki/Handbook:X86/Full/Networking/hu "Handbook:X86/Full/Networking/hu")[Munka elkezdése](/wiki/Handbook:X86/Networking/Introduction/hu "Handbook:X86/Networking/Introduction/hu")[Fejlett beállítások](/wiki/Handbook:X86/Networking/Advanced/hu "Handbook:X86/Networking/Advanced/hu")[Moduláris hálózat](/wiki/Handbook:X86/Networking/Modular/hu "Handbook:X86/Networking/Modular/hu")[Vezeték nélküli (Wi-Fi)](/wiki/Handbook:X86/Networking/Wireless/hu "Handbook:X86/Networking/Wireless/hu")[Funkcionalitás hozzáadása](/wiki/Handbook:X86/Networking/Extending/hu "Handbook:X86/Networking/Extending/hu")[Dinamikus menedzsment](/wiki/Handbook:X86/Networking/Dynamic/hu "Handbook:X86/Networking/Dynamic/hu")

## Contents

- [1Bevezetés a blokktípusú eszközökbe](#Bevezet.C3.A9s_a_blokkt.C3.ADpus.C3.BA_eszk.C3.B6z.C3.B6kbe)
  - [1.1Blokkeszközök](#Blokkeszk.C3.B6z.C3.B6k)
  - [1.2Partíciós táblázatok](#Part.C3.ADci.C3.B3s_t.C3.A1bl.C3.A1zatok)
    - [1.2.1GUID Partíciós Táblázat (GPT)](#GUID_Part.C3.ADci.C3.B3s_T.C3.A1bl.C3.A1zat_.28GPT.29)
    - [1.2.2Master boot record (MBR) vagyis DOS boot sector](#Master_boot_record_.28MBR.29_vagyis_DOS_boot_sector)
  - [1.3Fejlett adathordozó](#Fejlett_adathordoz.C3.B3)
  - [1.4Alapértelmezett particionálási séma](#Alap.C3.A9rtelmezett_particion.C3.A1l.C3.A1si_s.C3.A9ma)
- [2Partíciós séma megtervezése](#Part.C3.ADci.C3.B3s_s.C3.A9ma_megtervez.C3.A9se)
  - [2.1Hány partíció és mekkora méretű?](#H.C3.A1ny_part.C3.ADci.C3.B3_.C3.A9s_mekkora_m.C3.A9ret.C5.B1.3F)
  - [2.2Mi a helyzet a swap területtel?](#Mi_a_helyzet_a_swap_ter.C3.BClettel.3F)
  - [2.3Mi az EFI rendszerpartíció (ESP)?](#Mi_az_EFI_rendszerpart.C3.ADci.C3.B3_.28ESP.29.3F)
  - [2.4Mi a BIOS boot partíció?](#Mi_a_BIOS_boot_part.C3.ADci.C3.B3.3F)
- [3Adathordozó particionálása GPT táblázattal az UEFI számára](#Adathordoz.C3.B3_particion.C3.A1l.C3.A1sa_GPT_t.C3.A1bl.C3.A1zattal_az_UEFI_sz.C3.A1m.C3.A1ra)
  - [3.1Aktuális partícióelrendezés megtekintése](#Aktu.C3.A1lis_part.C3.ADci.C3.B3elrendez.C3.A9s_megtekint.C3.A9se)
  - [3.2Új táblázat létrehozása az adathordozón / Összes partíció eltávolítása az adathordozóról](#.C3.9Aj_t.C3.A1bl.C3.A1zat_l.C3.A9trehoz.C3.A1sa_az_adathordoz.C3.B3n_.2F_.C3.96sszes_part.C3.ADci.C3.B3_elt.C3.A1vol.C3.ADt.C3.A1sa_az_adathordoz.C3.B3r.C3.B3l)
  - [3.3EFI rendszerpartíció (ESP - EFI System Partition) létrehozása](#EFI_rendszerpart.C3.ADci.C3.B3_.28ESP_-_EFI_System_Partition.29_l.C3.A9trehoz.C3.A1sa)
  - [3.4Cserepartíció (swap partíció) létrehozása](#Cserepart.C3.ADci.C3.B3_.28swap_part.C3.ADci.C3.B3.29_l.C3.A9trehoz.C3.A1sa)
  - [3.5Gyökérpartíció (root partíció) létrehozása](#Gy.C3.B6k.C3.A9rpart.C3.ADci.C3.B3_.28root_part.C3.ADci.C3.B3.29_l.C3.A9trehoz.C3.A1sa)
  - [3.6Partíciók elrendezésének az elmentése az adathordozón](#Part.C3.ADci.C3.B3k_elrendez.C3.A9s.C3.A9nek_az_elment.C3.A9se_az_adathordoz.C3.B3n)
- [4Adathordozó particionálása MBR táblázattal a(z) BIOS/Örökölt boot számára](#Adathordoz.C3.B3_particion.C3.A1l.C3.A1sa_MBR_t.C3.A1bl.C3.A1zattal_a.28z.29_BIOS.2F.C3.96r.C3.B6k.C3.B6lt_boot_sz.C3.A1m.C3.A1ra)
  - [4.1Aktuális partícióelrendezés megtekintése](#Aktu.C3.A1lis_part.C3.ADci.C3.B3elrendez.C3.A9s_megtekint.C3.A9se_2)
  - [4.2Új táblázat létrehozása az adathordozón / Összes partíció eltávolítása az adathordozóról](#.C3.9Aj_t.C3.A1bl.C3.A1zat_l.C3.A9trehoz.C3.A1sa_az_adathordoz.C3.B3n_.2F_.C3.96sszes_part.C3.ADci.C3.B3_elt.C3.A1vol.C3.ADt.C3.A1sa_az_adathordoz.C3.B3r.C3.B3l_2)
  - [4.3Boot partíció létrehozása](#Boot_part.C3.ADci.C3.B3_l.C3.A9trehoz.C3.A1sa)
  - [4.4Cserepartíció (swap partíció) létrehozása](#Cserepart.C3.ADci.C3.B3_.28swap_part.C3.ADci.C3.B3.29_l.C3.A9trehoz.C3.A1sa_2)
  - [4.5Gyökérpartíció (root partíció) létrehozása](#Gy.C3.B6k.C3.A9rpart.C3.ADci.C3.B3_.28root_part.C3.ADci.C3.B3.29_l.C3.A9trehoz.C3.A1sa_2)
  - [4.6Partíciók elrendezésének az elmentése az adathordozón](#Part.C3.ADci.C3.B3k_elrendez.C3.A9s.C3.A9nek_az_elment.C3.A9se_az_adathordoz.C3.B3n_2)
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

### Partíciós táblázatok

Bár elméletileg lehetséges nyers, particionálatlan adathordozó használata Linux rendszer elhelyezésére (például btrfs RAID létrehozásakor), ez a gyakorlatban szinte soha nem történik meg. Ehelyett az adathordozó blokkeszközöket kisebb, jobban kezelhető blokkeszközökre osztják fel. Az **x86** rendszereken ezeket partícióknak nevezzük. Jelenleg kettő szabványos particionálási technológia van használatban:

- MBR - Néha DOS adathordozó címkének is nevezik.
- GPT - Ez két rendszerindítási folyamattípushoz van kötve. Az Örökölt BIOS boot -hoz, és az UEFI boot -hoz.

#### GUID Partíciós Táblázat (GPT)

A _GUID Partíciós Táblázat (GPT)_ (más néven a GPT adathordozó címke) 64 bites azonosítókat használ a partíciók azonosítására. A partícióinformációkra fenntartott hely sokkal nagyobb, mint az MBR partíciós táblázat (DOS adathordozó címke) esetében, ami ott csak 512 bájt nagyságot foglal el az adathordozón. Ez azt jelenti, hogy gyakorlatilag nincs korlátozás a GPT táblázattal létrehozott adathordozó partícióinak a számában. Ezenkívül a maximális partícióméret sokkal nagyobb (majdnem 8 ZiB -- igen, zebibyte) méretű is lehet.

Ha a rendszer szoftveres interfésze az operációs rendszer és a firmware között UEFI (a BIOS helyett), akkor a GPT szinte kötelező, mivel kompatibilitási problémák merülnek fel a DOS adathordozó címkével kapcsolatban.

A GPT kihasználja az ellenőrzőösszeg-számítást és a redundanciát is. CRC32 ellenőrző összegeket hordoz, amelyek segítségével a fejlécben és a partíciós táblázatokban esetlegesen keletkező hibákat lehet észlelni, és van egy biztonsági mentésként létrehozott második GPT táblázat az adathordozó végén. Ez az biztonsági mentésként létrehozott második táblázat felhasználható az elsődleges GPT táblázat sérüléseinek a helyreállítására az adathordozó elején.

**Important**

Létezik néhány figyelmeztetés a GPT táblázattal kapcsolatban:

- A GPT használata BIOS-alapú számítógépen működik, de a felhasználó nem tud duál-boot opciót használni Microsoft Windows operációs rendszerrel, mivel a Microsoft Windows nem hajlandó GPT partícióról bootolni BIOS módban.
- Egyes, hibás (régi) alaplapi firmware-ek, amelyek BIOS/CSM/Örökölt módban történő rendszerindításra vannak beállítva, szintén problémákat okozhatnak a GPT címkével rendelkező adathordozókról történő rendszerindításkor.

#### Master boot record (MBR) vagyis DOS boot sector

A _[Master boot record](https://en.wikipedia.org/wiki/Master_boot_record "wikipedia:Master boot record")_ rendszerindító szektort (más néven DOS boot szektor, DOS adathordozó címke – újabban, a GPT/UEFI beállításokkal ellentétes – Örökölt BIOS rendszerindítás) 1983-ban mutatták be először a PC DOS 2.x rendszerrel. Az MBR 32 bites azonosítókat használ a kezdő szektorhoz és a partíciók hosszához, és három partíciótípust támogat: elsődleges, kiterjesztett és logikai. Az elsődleges partíciók információi magában a fő rendszerindító rekordban vannak eltárolva – egy nagyon kicsi (általában 512 bájt méretű) helyen az adathordozó legelején. Ennek a kis helynek köszönhetően mindösszesen csak legfeljebb négy elsődleges partíció támogatására van lehetőség (például /dev/sda1 partíciótól /dev/sda4 partícióig).

Négynél több partíció támogatása érdekében az MBR egyik elsődleges partíciója megjelölhető _bővített_ partícióként. Ez a bővített partíció ezután további logikai partíciókat (partíciókat egy partíción belül) tartalmazhat.

**Important**

A legtöbb alaplapgyártó továbbra is támogatja egyelőre, az MBR rendszerindító szektorokat és a hozzájuk kapcsolódó particionálási korlátozásokat örököltnek tekintik. Hacsak nem 2010 előtti hardverrel dolgozik, akkor a legjobb megoldás, ha az adathordozót a [GUID Partition Table](https://en.wikipedia.org/wiki/GUID_Partition_Table "wikipedia:GUID Partition Table") technológiával particionálja. Azoknak az olvasóknak, akiknek folytatniuk kell a beállítási típust, tudomásul kell venniük a következő információkat:

- A 2010 utáni alaplapok többsége az MBR rendszerindító szektorok használatát örökölt (támogatott, de nem ideális) rendszerindítási módnak tekinti.
- A 32 bites azonosítóhasználatnak köszönhetően az MBR partíciós táblázat nem tud 2 TiB-nél nagyobb tárterületet megcímezni.
- Hacsak nem jön létre kiterjesztett partíció, az MBR legfeljebb négy partíciót támogat.
- Ez a technológia nem biztosít tartalék rendszerindító szektort, így ha valami felülírja a partíciós táblázatot, akkor minden partíció információ el fog veszni.

Ennek ellenére az MBR és a régi BIOS rendszerindítás továbbra is használható virtualizált felhőkörnyezetekben, például az AWS-ben.

A kézikönyv szerzői a Gentoo telepítésekhez a [GPT](#GPT) használatát javasolják, mindig amikor ez lehetséges.

### Fejlett adathordozó

A hivatalos Gentoo bootolható képfájlja támogatja a [Logical Volume Manager (LVM)](/wiki/LVM/hu "LVM/hu") használatát. Az LVM képes a _fizikai köteteket_, például partíciókat vagy adathordozókat _kötetcsoportokba_ kombinálni. A kötetcsoportok rugalmasabbak, mint a partíciók, és RAID-csoportok vagy gyorsítótárak meghatározására használhatók gyors SSD adathordozókon a lassú HDD adathordozók számára. Bár a kézikönyv nem tárgyalja a használatát, az LVM-et a Gentoo teljes mértékben támogatja.

Although usage is not covered in the handbook, below is a list helpful guides to get the system running:

- [Btrfs](/wiki/Btrfs/Native_System_Root_Guide "Btrfs/Native System Root Guide")
- [ZFS](/wiki/ZFS/rootfs "ZFS/rootfs")
- [LVM (Logical Volume Manager)](/wiki/LVM/hu "LVM/hu")

Disk encryption is also handled in the same manner:

- [Rootfs encryption](/wiki/Rootfs_encryption/hu "Rootfs encryption/hu")

### Alapértelmezett particionálási séma

A kézikönyv hátralévő részében kettő esetet fogunk megvitatni és elmagyarázni:

1. UEFI firmware, amely GUID partíciós táblázatú (GPT) adathordozóval rendelkezik.
2. MBR DOS/Örökölt BIOS firmware, amely MBR partíciós táblázatú adathordozóval rendelkezik.

Bár lehetséges a rendszerindítási típusok keverése és párosítása bizonyos alaplapi firmware esetében, a keverés meghaladja a kézikönyv szándékát. Amint azt korábban említettük, a modern hardverre telepített Gentoo operációs rendszerek esetében erősen ajánlott az UEFI rendszerindítás használata GPT címkés adathordozóval.

A következő particionálási sémát használjuk egyszerű példa elrendezésként.

**Important**

A következő táblázat első sora kizárólagos információkat tartalmaz a GPT adathordozó címkéről _**vagy**_ az MBR DOS/régi BIOS adathordozó címkéről. Ha kétségei vannak, akkor folytassa a GPT táblázattal, mivel a 2010 után gyártott **x86** számítógépek általában támogatják az UEFI firmware-t és a GPT boot szektort.

Partíció
Fájlrendszer
Méret
Leírás
/dev/sda1fat32 Az EFI rendszerpartícióhoz szükséges fájlrendszer, amely mindig GPT adathordozó címkével van társítva.1 GiB
EFI rendszerpartíció leírása. UEFI implementációt támogató rendszer firmware -re vonatkozik. Ez jellemzően a 2010-től napjainkig gyártott rendszerekre vonatkozik.ext4 Az MBR partíciós táblázat rendszerindító partíciójához ajánlott fájlrendszer, amelyet a DOS/Örökölt BIOS adathordozó címkére korlátozódó régebbi firmware -el együtt használnak.MBR DOS/Örökölt BIOS boot partíció leírása. Alkalmazható a régi BIOS-al ellátott számítógép firmware -re. Az ilyen típusú rendszereket jellemzően 2010 <u>előtt</u> gyártották, és általában fokozatosan meg lett szüntetve a gyártásuk./dev/sda2linux-swap
(RAM mérete) \* 2
A cserepartíció (swap partíció) leírása.
/dev/sda3xfs
Az adathordozón lévő összes többi hely A kiválasztott _**profil**_, a további _**partíciók**_ (opcionális) és a rendszer _**célja**_ megnehezíti a gyökér (root) fájlrendszer megfelelő méretezését, ezért a kézikönyv szerzői nem tudnak _mindenre alkalmas_ javaslatot adni a gyökér (root) fájlrendszer partícióját illetően.</br></br> Ha a Gentoo az egyetlen operációs rendszer, amely az adathordozót használja, akkor az adathordozón lévő összes fennmaradó hely kiválasztása a legbiztonságosabb és javasolt választási lehetőség.A gyökér (root) partíció leírása.

Ha ez az információ elegendő, akkor a haladószintű olvasó közvetlenül a tényleges particionálásra ugorhat.

Az fdisk és a parted is particionáló segédprogramok, amelyek benne vannak a hivatalos Gentoo bootolható lemezképfájlban. Az fdisk jól ismert, stabil, és mind az MBR, mind a GPT adathordozókat kezelni tudja. A parted volt az egyik első Linux blokkeszközkezelő-segédprogram, amely támogatta a GPT-partíciókat. Használható az fdisk alternatívájaként, ha az olvasó úgy kívánja, azonban a kézikönyv csak az fdisk segédprogramhoz ad utasításokat, mivel a legtöbb Linux-környezetben általánosan elérhető.

Mielőtt rátérnénk a létrehozási utasításokra, az első szakaszok részletesebben leírják, hogy miként hozhatók létre particionálási sémák, és megemlítenek néhány gyakori buktatót.

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

### Mi az EFI rendszerpartíció (ESP)?

Ha a Gentoo operációs rendszert olyan számítógépre telepíti, amely UEFI-t használ az operációs rendszer indítására (a BIOS helyett), akkor elengedhetetlen, hogy létre legyen hozva egy EFI rendszerpartíció (ESP - EFI System Partition). Az alábbi utasítások tartalmazzák a létrehozás műveletéhes szükséges lépéseket. **A(z) BIOS/Örökölt módban történő rendszerindításkor nincs szükség az EFI rendszerpartícióra.**

Az EFI rendszerpartíciónak valamilyen FAT variánsnak _kell_ lennie (Linux rendszereken néha _vfat_ fájlrendszerként jelenik meg). A hivatalos [UEFI specifikáció](http://www.uefi.org/sites/default/files/resources/UEFI%202_5.pdf) azt írja, hogy a FAT12, a FAT16 és a FAT32 fájlrendszert is felismeri az UEFI firmware-je, mégis az EFI rendszerpartícióhoz csak a FAT32 az ajánlott fájlrendszer. A particionálás után formázza az EFI rendszerpartíciót ennek megfelelően:

`root #` `mkfs.fat -F 32 /dev/sda1`

**Important**

Ha az EFI rendszerpartíció nem a FAT valamelyik változatával van formázva, akkor a rendszer UEFI firmware-je nem találja meg a rendszerbetöltő bootloader-t (vagy Linux kernelt), és nagy valószínűséggel nem tudja elindítani a rendszert!

### Mi a BIOS boot partíció?

A _BIOS rendszerindító partíció_ (BIOS boot partition) egy nagyon kicsi (1–2 MB méretű) partíció, amelybe a rendszerbetöltők (boot loaders), például a GRUB2, további adatokat helyezhetnek el, amelyek nem férnek el a lefoglalt tárhelyen. Kizárólag akkor van rá szükség, ha az adathordozó GPT táblázattal van formázva, de a rendszer firmware-je a GRUB2-n keresztül fog elindulni régi BIOS/MBR DOS rendszerindítási módban. **_Nem szükséges_ a megléte az adathordozón az EFI/UEFI módban történő rendszerindításkor, és nem szükséges a megléte az adathordozón az MBR/Örökölt DOS adathordozó táblázat használatakor sem.** Ez az útmutató nem tartalmaz _BIOS rendszerindító partíciót_ (BIOS boot partition).

## Adathordozó particionálása GPT táblázattal az UEFI számára

A következő részek bemutatják, hogy miként hozhatunk létre egy partícióelrendezést egyetlen GPT adathordozóhoz, amely megfelel az UEFI-specifikációnak és a felfedezhető partíciók specifikációnak (DPS). A DPS a Linux Userspace API (UAPI) csoportspecifikáció részeként biztosított specifikáció, és ajánlott a használata, de teljesen opcionális választási lehetőség. A specifikációk az fdisk segédprogram segítségével valósulnak meg, amely a [sys-apps/util-linux](https://packages.gentoo.org/packages/sys-apps/util-linux) szoftvercsomag része.

A táblázat egy ajánlott alapértelmezést nyújt egy egyszerű Gentoo telepítés számára. További partíciók hozzáadhatóak az egyéni preferenciák vagy a rendszertervezési célok szerint.

Eszköz elérési útvonala (sysfs)
Csatlakoztatási pont
Fájlrendszer
DPS UUID (Type-UUID)
Leírás
/dev/sda1/efivfat
c12a7328-f81f-11d2-ba4b-00a0c93ec93b
EFI rendszerpartíció (ESP) leírása
/dev/sda2Nem alkalmazható A swap nem úgy van csatlakoztatva a fájlrendszerhez, mint egy eszközfájl.0657fd6d-a4ab-43c4-84e5-0933c84b4f4f
Swap partíció leírása
/dev/sda3/xfs
44479540-f297-41b2-9af7-d131d5f0458a
Root partíció leírása

### Aktuális partícióelrendezés megtekintése

Az fdisk egy népszerű és hatékony segédprogram az adathordozó partíciókra történő felosztására. Indítsa el az fdisk segédprogramot a particionálni kívánt adathordozó megadásával (ebben a példában a fájlrendszer /dev/sda elérési útvonalán megtalálható adathordozót adjuk meg neki):

`root #` `fdisk /dev/sda`

Használja a `p` billentyűgombot az adathordozón lévő partícióbeállításoknak a megjelenítése érdekében:

`Command (m for help):` `p`

```
Disk /dev/sda: 931.51 GiB, 1000204886016 bytes, 1953525168 sectors
Disk model: HGST HTS721010A9
Units: sectors of 1 * 512 = 512 bytes
Sector size (logical/physical): 512 bytes / 4096 bytes
I/O size (minimum/optimal): 4096 bytes / 4096 bytes
Disklabel type: gpt
Disk identifier: 3E56EE74-0571-462B-A992-9872E3855D75

Device        Start        End    Sectors   Size Type
/dev/sda1      2048    2099199    2097152     1G EFI System
/dev/sda2   2099200   10487807    8388608     4G Linux swap
/dev/sda3  10487808 1953523711 1943035904 926.5G Linux root (x86-64)

```

Ebben a példában szereplő adathordozó úgy lett beállítva, hogy két Linux fájlrendszert tartalmazzon (mindegyik "Linux" partícióként van felsorolva), amelyek közül az egyik ("Linux swap"-ként) cserepartícióként van megjelölve.

### Új táblázat létrehozása az adathordozón / Összes partíció eltávolítása az adathordozóról

A `g` billentyűgomb lenyomása azonnal eltávolítja GPT táblázatot, és értelemszerűen ezzel törlődik az összes meglévő partíció is az adathordozóról, majd létrehoz egy új GPT táblázatot (nevezik még GPT lemezcímkének is):

`Command (m for help):` `g`

```
Created a new GPT disklabel (GUID: 3E56EE74-0571-462B-A992-9872E3855D75).

```

Alternatív megoldásként, egy meglévő GPT adathordozó táblázat megtartásának az érdekében (tekintse meg fent a `p` kimenetét), fontolja meg a meglévő partíciók egyenként történő eltávolítását az adathordozóról (magából a meglévő GPT táblázatból fogjuk kitörölni a partíciókat, de a táblázat az marad). Nyomja meg a `d` billentyűgombot a partíció törléséhez. Például a meglévő /dev/sda1 partíció törléséhez:

`Command (m for help):` `d`

```
Partition number (1-4): 1

```

A partíció most ütemezve lett a törlésre. A partíciók listájának képernyőn történő megjelenítése során ( `p`) már nem jelenik meg, de a partíció nem törlődik a változtatások mentéséig. Ez a viselkedés lehetővé teszi a felhasználó számára, hogy megszakítsa a törlési műveletet, ha hibát vétett - Ilyen esetben nyomja meg a `q`, majd ezt követően az `Enter` billentyűgombot, melynek következménye, hogy a partíció nem fog törlődni, tehát a partíciótörlés ütemezése vissza lesz vonva.

Nyomja meg többször a `p` billentyűgombot a partíciólista képernyőn való magjelenítéséhez, majd nyomja meg a `d` billentyűgombot és adja meg a partíció számát a törléshez. Végül a partíciós táblázat ki fog ürülni:

`Command (m for help):` `p`

```
Disk /dev/sda: 931.51 GiB, 1000204886016 bytes, 1953525168 sectors
Disk model: HGST HTS721010A9
Units: sectors of 1 * 512 = 512 bytes
Sector size (logical/physical): 512 bytes / 4096 bytes
I/O size (minimum/optimal): 4096 bytes / 4096 bytes
Disklabel type: gpt
Disk identifier: 3E56EE74-0571-462B-A992-9872E3855D75

```

Most, hogy a partíciós táblázat üres, készen állunk az új partíciók létrehozására.

### EFI rendszerpartíció (ESP - EFI System Partition) létrehozása

**Note**

Kisebb méretű ESP rendszerpartíció létrehozására lehetőség van, de az nem ajánlott eljárás. Különösen abban az esetben nem ajánlott, ha az ESP meg lesz osztva másik operációs rendszerekkel is.

Először hozzon létre egy kis EFI rendszerpartíciót, amely szintén /efi könyvtárnéven lesz felcsatolva a fájlrendszerbe. Új partíció létrehozásához nyomja meg az `n` billentyűgombot, majd ezt követően az `1` billentyűgombot az első partíció kiválasztásához. Amikor a rendszer kéri az első szektort, győződjön meg arról, hogy az 2048 számozástól indul (mivel a 0-tól 2048-ig terjedő szabad területre később szüksége lehet a rendszertöltő bootloadernek), és nyomja meg az `Enter` billentyűgombot. Amikor a rendszer az utolsó szektort kéri, írja be a +1G karaktereket az 1 gibibyte nagyságú partíció létrehozásának az érdekében:

`Command (m for help):` `n`

```
Partition number (1-128, default 1): 1
First sector (2048-1953525134, default 2048):
Last sector, +/-sectors or +/-size{K,M,G,T,P} (2048-1953525134, default 1953523711): +1G

Created a new partition 1 of type 'Linux filesystem' and of size 1 GiB.
Partition #1 contains a vfat signature.

Do you want to remove the signature? [Y]es/[N]o: Y
The signature will be removed by a write command.

```

Jelölje meg a partíciót EFI rendszerpartícióként:

`Command (m for help):` `t`

```
Selected partition 1
Partition type or alias (type L to list all): 1
Changed type of partition 'Linux filesystem' to 'EFI System'.

```

### Cserepartíció (swap partíció) létrehozása

Ezután a cserepartíció létrehozásához nyomja meg az `n` billentyűgombot egy új partíció létrehozásához, majd nyomja meg a `2` billentyűgombot a második (/dev/sda2) partíció létrehozásának érdekében. Amikor kéri az első szektort, nyomja meg az `Enter` billentyűt. Amikor a rendszer az utolsó szektort kéri, írja be a +4G értéket (vagy bármely más, a cserepartíció területhez szükséges méretet) egy 4 GiB nagyságú partíció létrehozásához.

`Command (m for help):` `n`

```
Partition number (2-128, default 2):
First sector (2099200-1953525134, default 2099200):
Last sector, +/-sectors or +/-size{K,M,G,T,P} (2099200-1953525134, default 1953523711): +4G

Created a new partition 2 of type 'Linux filesystem' and of size 4 GiB.

```

Ezután nyomja meg a `t` billentyűgombot a partíció típusának beállításához. Nyomja meg a `2` billentyűgombot az imént létrehozott partíció kiválasztásához, majd írja be a _19_ számot a partíció típusának "Linux Swap" (tehát a "Linux cserepartíció") beállításához.

`Command (m for help):` `t`

```
Partition number (1,2, default 2): 2
Partition type or alias (type L to list all): 19

Changed type of partition 'Linux filesystem' to 'Linux swap'.

```

### Gyökérpartíció (root partíció) létrehozása

Végül a gyökérpartíció létrehozásához nyomja meg az `n` billentyűgombot egy új partíció létrehozásához, majd nyomja meg a `3` billentyűgombot a harmadik (/dev/sda3) partíció létrehozásához. Amikor a rendszer kéri az első szektort, nyomja meg az `Enter` billentyűt. Amikor a rendszer kéri az utolsó szektort, nyomja le ismét az `Enter` billentyűt, hogy létrehozzon egy partíciót, amely lefoglalja az adathordozón az összes még lefoglalatlan tárhelyet a harmadik partíció számára:

`Command (m for help):` `n`

```
Partition number (3-128, default 3): 3
First sector (10487808-1953525134, default 10487808):
Last sector, +/-sectors or +/-size{K,M,G,T,P} (10487808-1953525134, default 1953523711):

Created a new partition 3 of type 'Linux filesystem' and of size 926.5 GiB..

```

**Note**

A gyökérpartíció (root partíció) típusának beállítása "Linux root (x86-64)" típusra nem szükséges, a rendszer normálisan fog működni, amennyiben az a "Linux filesystem" típusra van beállítva. A "Linux root (x86-64)" fájlrendszer-típusra történő ráállítás olyan esetekben szükséges, amikor az azt támogató rendszerbetöltő bootloder (azaz a systemd-boot) használatakor nincs szükség fstab fájlra.

A gyökérpartíció (a root partíció) létrehozása után nyomja meg a `t` billentyűgombot a partíció típusának beállításához, a `3` billentyűgombot az imént létrehozott partíció kiválasztásához, majd írja be a _23_ számot a partíció típusának "Linux Root (x86-64)" beállításához.

`Command(m for help):` `t`

```
Partition number (1-3, default 3): 3
Partition type or alias (type L to list all): 23

Changed type of partition 'Linux filesystem' to 'Linux root (x86-64)'

```

A lépések végrehajtása után a `p` billentyűgomb megnyomásával egy partíciós táblázat jelenik meg, amely a következőhöz hasonló:

`Command (m for help):` `p`

```
Disk /dev/sda: 931.51 GiB, 1000204886016 bytes, 1953525168 sectors
Disk model: HGST HTS721010A9
Units: sectors of 1 * 512 = 512 bytes
Sector size (logical/physical): 512 bytes / 4096 bytes
I/O size (minimum/optimal): 4096 bytes / 4096 bytes
Disklabel type: gpt
Disk identifier: 3E56EE74-0571-462B-A992-9872E3855D75

Device        Start        End    Sectors   Size Type
/dev/sda1      2048    2099199    2097152     1G EFI System
/dev/sda2   2099200   10487807    8388608     4G Linux swap
/dev/sda3  10487808 1953523711 1943035904 926.5G Linux root (x86-64)

Filesystem/RAID signature on partition 1 will be wiped.

```

### Partíciók elrendezésének az elmentése az adathordozón

Nyomja meg a `w` billentyűgombot a partícióelrendezés mentéséhez. Az fdisk segédprogram elmenti az adathordozón a partíciók elrendezését, majd befejezi a működését:

`Command (m for help):` `w`

```
The partition table has been altered.
Calling ioctl() to re-read partition table.
Syncing disks.

```

Mivel innentől kezdve a partíciók léteznek az adathordozón, a következő lépés a telepítésben a partíciókra rárakni fájlrendszereket.

## Adathordozó particionálása MBR táblázattal a(z) BIOS/Örökölt boot számára

A következő táblázat egy ajánlott partícióelrendezést kínál triviális MBR DOS / Örökölt BIOS rendszerindítási adathordozókhoz. További partíciók hozzáadhatók a személyes preferenciák vagy a rendszertervezési célok szerint.

Device path (sysfs)
Csatolási pont
Fájlrendszer
DPS UUID (PARTUUID)
Leírás
/dev/sda1/bootext4
N/A
MBR DOS / Örökölt BIOS boot partíció részletei
/dev/sda2Nem alkalmazható A cserepartíció nem úgy van felcsatlakoztatva a fájlrendszerhez, mint egy eszközfájl.0657fd6d-a4ab-43c4-84e5-0933c84b4f4f
Cserepartíció részletei
/dev/sda3/xfs
44479540-f297-41b2-9af7-d131d5f0458a
A gyökérpartíció részletei

Módosítsa a partíció elrendezését az Ön személyes preferenciái szerint.

### Aktuális partícióelrendezés megtekintése

Indítsa el a fdisk segédprogramot az adathordozó megadásával (a példánkban a /dev/sda adathordozót használjuk):

`root #` `fdisk /dev/sda`

Használja a `p` billentyűgombot az adathordozón található aktuális partícióelrendezés megjelenítésének az érdekében:

`Command (m for help):` `p`

```
Disk /dev/sda: 931.51 GiB, 1000204886016 bytes, 1953525168 sectors
Disk model: HGST HTS721010A9
Units: sectors of 1 * 512 = 512 bytes
Sector size (logical/physical): 512 bytes / 4096 bytes
I/O size (minimum/optimal): 4096 bytes / 4096 bytes
Disklabel type: dos
Disk identifier: 0xf163b576

Device     Boot    Start        End    Sectors   Size Id Type
/dev/sda1  *        2048    2099199    2097152     1G 83 Linux
/dev/sda2        2099200   10487807    8388608     4G 82 Linux swap / Solaris
/dev/sda3       10487808 1953525167 1943037360 926.5G 83 Linux

```

Ez az adathordozó jelenleg úgy van beállítva, hogy két Linux fájlrendszerrel megformázott partíciót tartalmaz (mind a kettő partíció "Linux" néven szerepel), amely közül az egyik cserepartíció is egyben ("Linux swap" néven szerepel).

### Új táblázat létrehozása az adathordozón / Összes partíció eltávolítása az adathordozóról

Az `o` billentyűgomb lenyomása azonnal eltávolítja az összes meglévő partíciót az adathordozóról, és létrehoz azon egy új MBR táblázatot (nevezik még DOS lemezcímkének is):

`Command (m for help):` `o`

```
Created a new DOS disklabel with disk identifier 0xf163b576.
The device contains 'gpt' signature and it will be removed by a write command. See fdisk(8) man page and --wipe option for more details.

```

Alternatív megoldásként, egy meglévő DOS adathordozó táblázat megtartásának az érdekében (tekintse meg fent a `p` kimenetét), fontolja meg a meglévő partíciók egyenként történő eltávolítását az adathordozóról (magából a meglévő DOS táblázatból fogjuk kitörölni a partíciókat, de a táblázat az marad). Nyomja meg a `d` billentyűgombot a partíció törléséhez. Például a meglévő /dev/sda1 partíció törléséhez:

`Command (m for help):` `d`

```
Partition number (1-4): 1

```

A partíció most ütemezve lett a törlésre. A partíciók listájának képernyőn történő megjelenítése során ( `p`) már nem jelenik meg, de a partíció nem törlődik a változtatások mentéséig. Ez a viselkedés lehetővé teszi a felhasználó számára, hogy megszakítsa a törlési műveletet, ha hibát vétett - Ilyen esetben nyomja meg a `q`, majd ezt követően az `Enter` billentyűgombot, melynek következménye, hogy a partíció nem fog törlődni, tehát a partíciótörlés ütemezése vissza lesz vonva.

Nyomja meg többször a `p` billentyűgombot a partíciólista képernyőn való magjelenítéséhez, majd nyomja meg a `d` billentyűgombot és adja meg a partíció számát a törléshez. Végül a partíciós táblázat ki fog ürülni:

`Command (m for help):` `p`

```
Disk /dev/sda: 931.51 GiB, 1000204886016 bytes, 1953525168 sectors
Disk model: HGST HTS721010A9
Units: sectors of 1 * 512 = 512 bytes
Sector size (logical/physical): 512 bytes / 4096 bytes
I/O size (minimum/optimal): 4096 bytes / 4096 bytes
Disklabel type: dos
Disk identifier: 0xf163b576

```

Az adathordozó készen áll az új partíciók létrehozására.

### Boot partíció létrehozása

Először hozzon létre egy kismértékű partíciót, amely /boot könyvtárnéven lesz majd felcsatlakoztatva a fájlrendszerbe. Nyomja meg az `n` billentyűgombot egy új partíció létrehozásához. Ezt követően nyomja meg a `p` billentyűgombot az _elsődleges_ partícióhoz, és az `1` billentyűgombot az első elsődleges partíció kiválasztásához. Amikor a rendszer kéri az első szektort, győződjön meg arról, hogy az 2048 számtól indul (ez szükséges lehet a rendszertöltőhöz), és nyomja meg az `Enter` billentyűt. Amikor a rendszer az utolsó szektort kéri, írja be a +1G karaktereket egy 1 GiB méretű partíció létrehozásához:

`Command (m for help):` `n`

```
Partition type
   p   primary (0 primary, 0 extended, 4 free)
   e   extended (container for logical partitions)
Select (default p): p
Partition number (1-4, default 1): 1
First sector (2048-1953525167, default 2048):
Last sector, +/-sectors or +/-size{K,M,G,T,P} (2048-1953525167, default 1953525167): +1G

Created a new partition 1 of type 'Linux' and of size 1 GiB.

```

Jelölje meg a partíciót bootable-ként az `a` billentyűgomb, majd az `Enter` billentyűgomb lenyomásával (Ezáltal lehet majd róla operációs-rendszert indítani):

`Command (m for help):` `a`

```
Selected partition 1
The bootable flag on partition 1 is enabled now.

```

Megjegyzés: Ha egynél több partíció áll rendelkezésre az adathordozón, akkor azt a partíciót kell kiválasztani, amelyet rendszerindítóként (bootolóként) kell megjelölni.

### Cserepartíció (swap partíció) létrehozása

Ezután a cserepartíció létrehozásához nyomja meg az `n` billentyűgombot egy új partíció létrehozásához, majd a `p` billentyűgombot. Ezt követően nyomja meg a `2` billentyűgombot a második elsődleges partíció (/dev/sda2) létrehozásához. Amikor a rendszer kéri az első szektort, nyomja meg az Enter billentyűt. Amikor a rendszer az utolsó szektort kéri, írja be a +4G értéket (vagy bármely más, a csereterülethez szükséges méretet), hogy létrehozzon egy 4 GB-os partíciót.

`Command (m for help):` `n`

```
Partition type
   p   primary (1 primary, 0 extended, 3 free)
   e   extended (container for logical partitions)
Select (default p): p
Partition number (2-4, default 2): 2
First sector (2099200-1953525167, default 2099200):
Last sector, +/-sectors or +/-size{K,M,G,T,P} (2099200-1953525167, default 1953525167): +4G

Created a new partition 2 of type 'Linux' and of size 4 GiB.

```

Miután mindez megtörtént, nyomja meg a `t` billentyűgombot a partíció típusának beállításához. Nyomja meg a `2` billentyűgombot az imént létrehozott partíció kiválasztásához, majd írja be a _82_ számot a partíció "Linux Swap" típusúra történő beállításához.

`Command (m for help):` `t`

```
Partition number (1,2, default 2): 2
Hex code (type L to list all codes): 82

Changed type of partition 'Linux' to 'Linux swap / Solaris'.

```

### Gyökérpartíció (root partíció) létrehozása

Végül a gyökérpartíció létrehozásához nyomja meg az `n` billentyűgombot egy új partíció létrehozásához. Ezután nyomja meg a p és a `3` billentyűgombot a harmadik elsődleges partíció (/dev/sda3) létrehozásához. Amikor a segédprogram az első szektor megadását kéri, nyomja meg az `Enter` billentyűgombot. Amikor a rendszer kéri az utolsó szektort, nyomja le ismét az `Enter` billentyűt, hogy létrehozzon egy partíciót, amely lefoglalja az adathordozón az összes még lefoglalatlan tárhelyet a harmadik partíció számára:

`Command (m for help):` `n`

```
Partition type
   p   primary (2 primary, 0 extended, 2 free)
   e   extended (container for logical partitions)
Select (default p): p
Partition number (3,4, default 3): 3
First sector (10487808-1953525167, default 10487808):
Last sector, +/-sectors or +/-size{K,M,G,T,P} (10487808-1953525167, default 1953525167):

Created a new partition 3 of type 'Linux' and of size 926.5 GiB.

```

Ezen lépések elvégzése után a `p` billentyűgomb megnyomását követően egy ehhez hasonló partíciós táblázat fog megjelenni:

`Command (m for help):` `p`

```
Disk /dev/sda: 931.51 GiB, 1000204886016 bytes, 1953525168 sectors
Disk model: HGST HTS721010A9
Units: sectors of 1 * 512 = 512 bytes
Sector size (logical/physical): 512 bytes / 4096 bytes
I/O size (minimum/optimal): 4096 bytes / 4096 bytes
Disklabel type: dos
Disk identifier: 0xf163b576

Device     Boot    Start        End    Sectors   Size Id Type
/dev/sda1  *        2048    2099199    2097152     1G 83 Linux
/dev/sda2        2099200   10487807    8388608     4G 82 Linux swap / Solaris
/dev/sda3       10487808 1953525167 1943037360 926.5G 83 Linux

```

### Partíciók elrendezésének az elmentése az adathordozón

Nyomja meg a `w` billentyűgombot a partícióelrendezés adathordozóra történő ráírásához. Ezen művelet sikeres befejeztével a fdisk szoftver befejezi a futását:

`Command (m for help):` `w`

```
The partition table has been altered.
Calling ioctl() to re-read partition table.
Syncing disks.

```

Itt az ideje, hogy fájlrendszereket helyezzünk el a partíciókon.

## Fájlrendszerek létrehozása

**Warning**

SSD vagy NVMe adathordozó használatakor bölcs dolog ellenőrizni a firmware-frissítéseket. Különösen egyes Intel SSD adathordozók (600p és 6000p) firmware-frissítést igényelnek az XFS I/O használati minták által okozott [lehetséges adatsérülések miatt](https://bugzilla.redhat.com/show_bug.cgi?id=1402533). A probléma a firmware szintjén van, és nem az XFS fájlrendszer hibája. A smartctl segédprogram segíthet az adathordozzó eszköz modelljének és firmware-verziójának ellenőrzésében.

### Bevezetés

Most, hogy a partíciók elkészültek, ideje fájlrendszert helyezni rájuk. A következő részben a Linux által támogatott különféle fájlrendszereket ismertetjük. Azok az olvasók, akik már tudják, hogy melyik fájlrendszert fogják használni, folytathatják a [Fájlrendszer rárakása egy partícióra](/wiki/Handbook:X86/Installation/Disks/hu#Applying_a_filesystem_to_a_partition "Handbook:X86/Installation/Disks/hu") című bekezdéssel. A többi felhasználónak érdemes továbbolvasniuk, hogy megismerjék az alkalmazható fájlrendszereket...

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

Innentől folytatható az a rendszertelepítés, amely korábban el lett kezdve, de a telepítési folyamat nem let végig befejezve. Használja ezt a hivatkozást állandó hivatkozásként: [A telepítés folytatása itt kezdődik](/wiki/Handbook:X86/Installation/Disks/hu#Resumed_installations_start_here "Handbook:X86/Installation/Disks/hu").

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

Később az utasításokban a proc fájlrendszer (a kernellel kapcsolatban álló virtuális interfész) és a többi kernel pszeudofájlrendszer lesz felcsatolva. Először viszont még a [Gentoo-stage fájlt](/wiki/Handbook:X86/Installation/Stage/hu "Handbook:X86/Installation/Stage/hu") ki kell csomagolnunk.

[← Hálózat beállítása](/wiki/Handbook:X86/Installation/Networking/hu "Previous part") [Kezdőlap](/wiki/Handbook:X86/hu "Handbook:X86/hu") [Stage fájl telepítése →](/wiki/Handbook:X86/Installation/Stage/hu "Next part")