# Disks

Other languages:

- [Deutsch](/wiki/Handbook:Alpha/Installation/Disks/de "Handbuch:Alpha/Installation/Festplatten (100% translated)")
- [English](/wiki/Handbook:Alpha/Installation/Disks "Handbook:Alpha/Installation/Disks (100% translated)")
- [Türkçe](/wiki/Handbook:Alpha/Installation/Disks/tr "Handbook:Alpha/Installation/Disks/tr (0% translated)")
- [español](/wiki/Handbook:Alpha/Installation/Disks/es "Manual de Gentoo: Alpha/Instalación/Discos (100% translated)")
- [français](/wiki/Handbook:Alpha/Installation/Disks/fr "Manuel:Alpha/Installation/Disques (100% translated)")
- [italiano](/wiki/Handbook:Alpha/Installation/Disks/it "Handbook:Alpha/Installation/Disks/it (100% translated)")
- [magyar](/wiki/Handbook:Alpha/Installation/Disks/hu "Handbook:Alpha/Installation/Disks/hu (100% translated)")
- [polski](/wiki/Handbook:Alpha/Installation/Disks/pl "Handbook:Alpha/Installation/Disks (100% translated)")
- [português do Brasil](/wiki/Handbook:Alpha/Installation/Disks/pt-br "Manual:Alpha/Instalação/Discos (100% translated)")
- čeština
- [русский](/wiki/Handbook:Alpha/Installation/Disks/ru "Handbook:Alpha/Installation/Disks (100% translated)")
- [தமிழ்](/wiki/Handbook:Alpha/Installation/Disks/ta "கையேடு:Alpha/நிறுவல்/வட்டுக்கள் (100% translated)")
- [中文（中国大陆）‎](/wiki/Handbook:Alpha/Installation/Disks/zh-cn "手册：Alpha/安装/准备磁盘 (100% translated)")
- [日本語](/wiki/Handbook:Alpha/Installation/Disks/ja "ハンドブック:Alpha/インストール/ディスク (100% translated)")
- [한국어](/wiki/Handbook:Alpha/Installation/Disks/ko "Handbook:Alpha/Installation/Disks/ko (100% translated)")

[Alpha Handbook](/wiki/Handbook:Alpha/cs "Handbook:Alpha/cs")[Instalace](/wiki/Handbook:Alpha/Full/Installation/cs "Handbook:Alpha/Full/Installation/cs")[O instalaci](/wiki/Handbook:Alpha/Installation/About/cs "Handbook:Alpha/Installation/About/cs")[Výběr média](/wiki/Handbook:Alpha/Installation/Media/cs "Handbook:Alpha/Installation/Media/cs")[Konfigurace sítě](/wiki/Handbook:Alpha/Installation/Networking/cs "Handbook:Alpha/Installation/Networking/cs")Příprava disků[Instalace stage3](/wiki/Handbook:Alpha/Installation/Stage/cs "Handbook:Alpha/Installation/Stage/cs")[Instalace základního systému](/wiki/Handbook:Alpha/Installation/Base/cs "Handbook:Alpha/Installation/Base/cs")[Konfigurace jádra](/wiki/Handbook:Alpha/Installation/Kernel/cs "Handbook:Alpha/Installation/Kernel/cs")[Konfigurace systému](/wiki/Handbook:Alpha/Installation/System/cs "Handbook:Alpha/Installation/System/cs")[Instalace nástrojů](/wiki/Handbook:Alpha/Installation/Tools/cs "Handbook:Alpha/Installation/Tools/cs")[Konfigurace zavaděče](/wiki/Handbook:Alpha/Installation/Bootloader/cs "Handbook:Alpha/Installation/Bootloader/cs")[Dokončení](/wiki/Handbook:Alpha/Installation/Finalizing/cs "Handbook:Alpha/Installation/Finalizing/cs")[Práce s Gentoo](/wiki/Handbook:Alpha/Full/Working/cs "Handbook:Alpha/Full/Working/cs")[Úvod do Portage](/wiki/Handbook:Alpha/Working/Portage/cs "Handbook:Alpha/Working/Portage/cs")[Přepínače USE](/wiki/Handbook:Alpha/Working/USE/cs "Handbook:Alpha/Working/USE/cs")[Funkce portage](/wiki/Handbook:Alpha/Working/Features/cs "Handbook:Alpha/Working/Features/cs")[Systém initskriptů](/wiki/Handbook:Alpha/Working/Initscripts/cs "Handbook:Alpha/Working/Initscripts/cs")[Proměnné prostředí](/wiki/Handbook:Alpha/Working/EnvVar/cs "Handbook:Alpha/Working/EnvVar/cs")[Práce s Portage](/wiki/Handbook:Alpha/Full/Portage/cs "Handbook:Alpha/Full/Portage/cs")[Soubory a adresáře](/wiki/Handbook:Alpha/Portage/Files/cs "Handbook:Alpha/Portage/Files/cs")[Proměnné](/wiki/Handbook:Alpha/Portage/Variables/cs "Handbook:Alpha/Portage/Variables/cs")[Mísení softwarových větví](/wiki/Handbook:Alpha/Portage/Branches/cs "Handbook:Alpha/Portage/Branches/cs")[Doplňkové nástroje](/wiki/Handbook:Alpha/Portage/Tools/cs "Handbook:Alpha/Portage/Tools/cs")[Vlastní strom Portage](/wiki/Handbook:Alpha/Portage/CustomTree/cs "Handbook:Alpha/Portage/CustomTree/cs")[Pokročilé funkce](/wiki/Handbook:Alpha/Portage/Advanced/cs "Handbook:Alpha/Portage/Advanced/cs")[Konfigurace sítě](/wiki/Handbook:Alpha/Full/Networking/cs "Handbook:Alpha/Full/Networking/cs")[Začínáme](/wiki/Handbook:Alpha/Networking/Introduction/cs "Handbook:Alpha/Networking/Introduction/cs")[Pokročilá konfigurace](/wiki/Handbook:Alpha/Networking/Advanced/cs "Handbook:Alpha/Networking/Advanced/cs")[Modulární síťování](/wiki/Handbook:Alpha/Networking/Modular/cs "Handbook:Alpha/Networking/Modular/cs")[Bezdrátové sítě](/wiki/Handbook:Alpha/Networking/Wireless/cs "Handbook:Alpha/Networking/Wireless/cs")[Přidání funkcí](/wiki/Handbook:Alpha/Networking/Extending/cs "Handbook:Alpha/Networking/Extending/cs")[Dynamická správa](/wiki/Handbook:Alpha/Networking/Dynamic/cs "Handbook:Alpha/Networking/Dynamic/cs")

## Contents

- [1Úvod k blokovým zařízením](#.C3.9Avod_k_blokov.C3.BDm_za.C5.99.C3.ADzen.C3.ADm)
  - [1.1Bloková zařízení](#Blokov.C3.A1_za.C5.99.C3.ADzen.C3.AD)
- [2Vytvoření systému souborů](#Vytvo.C5.99en.C3.AD_syst.C3.A9mu_soubor.C5.AF)
  - [2.1Úvod](#.C3.9Avod)
  - [2.2Systémy souborů](#Syst.C3.A9my_soubor.C5.AF)
  - [2.3Aplikace systému souborů na diskový oddíl](#Aplikace_syst.C3.A9mu_soubor.C5.AF_na_diskov.C3.BD_odd.C3.ADl)
    - [2.3.1EFI system partition filesystem](#EFI_system_partition_filesystem)
    - [2.3.2Legacy BIOS boot partition filesystem](#Legacy_BIOS_boot_partition_filesystem)
    - [2.3.3Small ext4 partitions](#Small_ext4_partitions)
  - [2.4Aktivace swap oddílu](#Aktivace_swap_odd.C3.ADlu)
- [3Připojení kořenového oddílu](#P.C5.99ipojen.C3.AD_ko.C5.99enov.C3.A9ho_odd.C3.ADlu)

## Úvod k blokovým zařízením

### Bloková zařízení

Pojďme se podrobně podívat na problematiku Gentoo Linuxu a Linuxu obecně ohledně disků, včetně blokových zařízení, diskových oddílů a systémů souborů. Jakmile si ukážeme všechny souvislosti, můžeme vytvořit oddíly disku a souborové systémy k nainstalování Gentoo Linuxu.

Pro začátek se podívejme na bloková zařízení. SCSI a Serial ATA disky jsou oboje označovány jako zařízení /dev/sda, /dev/sdb, /dev/sdc atd. Na novějších strojích s PCI Express NVMe pevnými disky jsou to zařízení /dev/nvme0n1, /dev/nvme0n2 atd.

Následující tabulka pomůže čtenářům určit, kde hledat v systému vyjmenovaná bloková zařízení:

Typ zařízeníVýchozí soubor zařízeníPoznámky a postřehy autorů
NVM Express (NVMe)/dev/nvme0n1Nejnovější technologie v oblasti solid state, disky [NVMe](https://en.wikipedia.org/wiki/NVM_Express "wikipedia:NVM Express") jsou připojené ke sběrnice PCI Express a mají nejvyšší přenosovou rychlost v blokových transkacích. Podpora zařízení NVMe se dá najít v systémech od roku 2014 a dál.
SATA, SAS, SCSI nebo USB flash/dev/sdaNalézá se v hardwaru zhruba od roku 2007 až do současnosti, tato zařízení jsou v Linuxu zřejmě nejpoužívanější. Tato zařízení se připojují prostřednictvím sběrnic [SATA](https://en.wikipedia.org/wiki/Serial_ATA "wikipedia:Serial ATA"), [SCSI](https://en.wikipedia.org/wiki/SCSI "wikipedia:SCSI") nebo [USB](https://en.wikipedia.org/wiki/USB "wikipedia:USB") jako bloová úložiště.
MMC, eMMC a SD/dev/mmcblk0Zařízení [embedded MMC](https://en.wikipedia.org/wiki/MultiMediaCard#eMMC "wikipedia:MultiMediaCard"), SD karty a další typy paměťových karet jsou použitelné jako úložiště dat. Nicméně mnoho systémů nemusí umožňovat nastartování pomocí těchto zařízení. Doporučuje se **nepoužívat** tato zařízení pro aktivní instalace Linuxu; spíše zvažte jejich použití pro přenos souborů, pro tento účel byla navržena. Alternativně je lze využít pro krátkodobé zálohy.
IDE/PATA/dev/hdaStarší ovladače Linuxového jádra [IDE/Parallel ATA](https://en.wikipedia.org/wiki/Parallel_ATA "wikipedia:Parallel ATA") hardwaru zorbazovaly zařízení rotačních blokových zařízení připojených ke sběrnici IDE v tomto místě. Všeobecně lze říct, že tato zařízení jsou vyřazována z osobních počítačů od roku 2003, kdy se průmysl uchýlil ke standardu SATA. Většina systémů s jedním IDE řadičem může podporovat až čtyři zařízení (hda-hdd).

Alternativní pojmenování pro tato starší rozhraní zahrnují Extended IDE (EIDE) a Ultra ATA (UATA).

Shora uvedená bloková zařízení představují abstraktní rozhraní disků. Uživatelské programy využívají tato bloková zařízení k interakci s diskem bez toho, aby se musely starat o to, zda jde o disky IDE, SCSI nebo nějaké jiné. Program může přistupovat k úložnému prostoru na disku jako k hromadě souvislých, náhodně přístupných 4096bytových (4K) bloků.

[Handbook:Alpha/Blocks/Disks/cs](/index.php?title=Handbook:Alpha/Blocks/Disks/cs&action=edit&redlink=1 "Handbook:Alpha/Blocks/Disks/cs (page does not exist)")

## Vytvoření systému souborů

**Upozornění**

When using an SSD or NVMe drive, it is wise to check for firmware upgrades. Some Intel SSDs in particular (600p and 6000p) require a firmware upgrade for [possible data corruption](https://bugzilla.redhat.com/show_bug.cgi?id=1402533) induced by XFS I/O usage patterns. The problem is at the firmware level and not any fault of the XFS filesystem. The smartctl utility can help check the device model and firmware version.

### Úvod

Po vytvoření oddílů na disku je načase vytvořit na nich systém souborů. V následující sekci si popíšeme různé systémy souborů, které Linux podporuje. Čtenáři, kteří už vědí, jaký systém souborů použít, mohou pokračovat na [aplikaci systému souborů na diskový oddíl](#Aplikace_syst.C3.A9mu_soubor.C5.AF_na_diskov.C3.BD_odd.C3.ADl). Ostatní by si měli přečíst, jaké systémy souborů jsou pro ně k dispozici.

### Systémy souborů

Linux podporuje tucty souborových systémů, ačkoli většinu z nich je radno používat pro specifické účely. Pouze některé souborové systémy mohou být považovány na alpha architečktuře za stabilní. Radíme proto přečist si informace o systémech souborů a stavu jejich podpory před tím, než zvolíte některý experimentálnější z nich pro důležitý diskový oddíl.

[btrfs](/wiki/Btrfs "Btrfs")Systém souborů příští generace, který poskytuje mnoho pokročilých funkcí jako je tvorba snapshotů, samooprava prostřednictvím kontrolních součtů, transparentní komprese, pododdíly a integrovaný RAID. U verzí jádra nižších než je 5.4.y nelze garantovat bezpečné používání btrf v produkčních systémech, protože opravy závažných chyb jsou přítomné v novějších vydáních větví jádra LTS. Chyby systému souborů jsou na starších větvích jáder časté, přičemž jakákoli verze starší než je 4.4.y je obvzláště nebezpečná a náchylná k chybám. Pravěpodobnost chyb je vyšší na starších jádrech (než 5.4.y), když je zapnutá komprese. RAID 5/6 a skupiny kvót nejsou bezpečné v žádné verzi btrfs. Dále je třeba uvést, že btrfs může nečekaně selhávat při operacích souborového systému s chybou ENOSPC, ačkoli df hlásí volný prostor, v důsledku vnitřní fragmentace (volné místo rezervované pro DATA + SYSTEM kusy, ale potřebné pro METADATA kusy). Navíc, jediný 4K odkaz na 128M extent uvnitř btrfs může znamenat, že je tu volné místo, které však není k dispozici pro alokaci. To může také způsobit, že btrfs vrátí chybu ENOSPC, ačkoli df oznamuje volné místo. Instalace balíčku [sys-fs/btrfsmaintenance](https://packages.gentoo.org/packages/sys-fs/btrfsmaintenance) a nastavení opakoveného spouštění skriptů může napomoci snížit výskyt problémů s ENOSPC rebalancováním btrfs, ale zcela to riziku výskytu ENOSPC nezabrání, i když je na disku volné místo. Některé činnosti nevyvolají ENOSPC nikdy, zatímco jiné ano. Pokud je riziko chyby ENOSPC v produkci neakceptovatelné, měli byste použít něco jiného. Pokud použijte btrfs, určitě se vyhněte konfiguracím, o nichž je známo, že působí problémy. S výjimkou ENOSPC jsou informace o těchto problémech btrfs v posledních nejnovějších k dispozi na [wiki stránce statusu btrfs](https://btrfs.wiki.kernel.org/index.php/Status).ext2Toto je vyzkoušený a opravdový linuxový systém souborů, který ovšem nemá žurnál metadat, což znamená, že běžná kontrola systému souborů během startu může být značně časově náročná. V současnosti je na výběr celá řada systémů souborů nové generace, jejichž konzistence může být zkontrolována rychle a proto jsou obecně upřednostňovány oproti svým bezžurnálovým protějšků. Žurnálovací systémy souborů předchází dlouhým prodlevám při bootování systému, když je systém souborů v nekonzistentním stavu.ext3Žurnálem vybavená verze systému souborů ext2, která poskytuje žurnálování metadat pro rychlou obnovu s přídavkem dalších vylepšených módů žurnálu jako je full data nebo ordered data žurnálování. Používá HTree index, který umožňuje vysoký výkon téměř ve všech situacích. ve zlratce je ext3 velmi dobrý a spolehlivý systém souborů.[ext4](/wiki/Ext4 "Ext4")Ext4 byl od počátku oddělen z ext3 a přináší nové funkce, vylepšení výkonu a odstraňuje velikostní limity s menšími změna ve formátu dat na disku. Může pokrýt svazky až do velikosti 1 EB s maximální velikostí jednoho souboru 16TB. Na rozdíl od klasické alokace bloků pomocí bitmap z ext2/3, ext4 používá extenty, což vylepšuje výkon při práci s velkými soubory a snižuje fragmentaci. Ext4 přináší také sofistikovanější algoritmy pro alokaci bloků (zpožděná alokace a vícebloková alokace) poskytující ovladači systému souborů více cest k optimalizaci uložení dat na disku. Ext4 je doporučeným všestranným systémem souborů na všech platformách.[f2fs](/wiki/F2fs "F2fs")Flash-Friendly File System byl původně vytořen společností Samsung k použití v pamětech NAND flash. Ke 2. čtvrtletí roku 2016 je stále považován za nezralý, ale jedná se o dobrou volbu při instalaci Gentoo na microSD karty, USB disky a další úložná zařízení založená na technologii flash.JFSVysoce výkonný žurnálovací systém souborů společnosti IBM. JFS je lehký, rychlý a spolehlivý systém souborů založený na B+tree, s dobrým výkonem v různých podmínkách.[ReiserFS](/wiki/ReiserFS "ReiserFS")B+tree žurnálovací systém souborů, který má dobrý celkový výkon, zejména při práci s mnoha malými soubory, za cenu více spotřebovaných cyklů CPU. ReiserFS verze 3 je součástí hlavní větve Linuxového jádra, ale jeho použití pči instalaci Gentoo se nedporučuje. Existují také novější verze souborového systému ReiserFS, ale ty vyžadují využití patchů hlavní větve jádra.[XFS](/wiki/XFS "XFS")Systém souborů s žurnálováním metadat, který robustní sadou vlastností a je optimalizován ke škálování. XFS je náchylnější k různým problémům s hardwarem, ale je postupně doplňován o novější funkce.[VFAT](/wiki/VFAT "VFAT")Systém souborů známý také jako FAT32 je Linuxem podporován, avšak neobsahuje podporu pro standardní UNIXové nastavení práv. Nejvíce je používán z důvodů interoperability s ostatními operečními systémy (především Microsoft Windows nebo Apple OSX), ale je nepostradatelný také pro firmware systémového zavaděče (jako je UEFI).[NTFS](/wiki/NTFS "NTFS")Systém souborů "New Technology" je hlavním systémem souborů Microsoft Windows od Windows NT 3.1. Podobně jako výše uvedený vfat neukládá nastavení UNIXových práv nebo rozšířené atributy nezbytné pro řádné fungování BSD nebo Linuxu, protože by neměl být používán jako kořenový systém souborů. Měl by být použit _pouze_ pro interoperabilitu se systémy Microsoft Windows (povšimněte si zvýraznění slova _pouze_).

More extensive information on filesystems can be found in the community maintained [Filesystem article](/wiki/Filesystem "Filesystem").

### Aplikace systému souborů na diskový oddíl

**Poznámka**

Please make sure to emerge the relevant user space utilities package for the chosen filesystem before rebooting. There will be a reminder to do so near the end of the installation process.

Systém souborů vytvoříme pomocí uživatelských utilit, které jsou k dispozici pro každý z možných systémů souborů. Klikněte na název systému souborů v níže uvedené tabulce pro doplňující informace o každém z nich.

Systém souborů
Příkaz k vytvoření
Na minimálním CD?
Balíček
[btrfs](/wiki/Btrfs "Btrfs")mkfs.btrfs Ano
[sys-fs/btrfs-progs](https://packages.gentoo.org/packages/sys-fs/btrfs-progs)[ext2](/wiki/Ext2 "Ext2")mkfs.ext2 Ano
[sys-fs/e2fsprogs](https://packages.gentoo.org/packages/sys-fs/e2fsprogs)[ext3](/wiki/Ext3 "Ext3")mkfs.ext3 Ano
[sys-fs/e2fsprogs](https://packages.gentoo.org/packages/sys-fs/e2fsprogs)[ext4](/wiki/Ext4 "Ext4")mkfs.ext4 Ano
[sys-fs/e2fsprogs](https://packages.gentoo.org/packages/sys-fs/e2fsprogs)[f2fs](/wiki/F2fs "F2fs")mkfs.f2fs Ano
[sys-fs/f2fs-tools](https://packages.gentoo.org/packages/sys-fs/f2fs-tools)[jfs](/wiki/JFS "JFS")mkfs.jfs Ano
[sys-fs/jfsutils](https://packages.gentoo.org/packages/sys-fs/jfsutils)[reiserfs](/wiki/ReiserFS "ReiserFS")mkfs.reiserfs Ano
[sys-fs/reiserfsprogs](https://packages.gentoo.org/packages/sys-fs/reiserfsprogs)[xfs](/wiki/XFS "XFS")mkfs.xfs Ano
[sys-fs/xfsprogs](https://packages.gentoo.org/packages/sys-fs/xfsprogs)[vfat](/wiki/FAT "FAT")mkfs.vfat Ano
[sys-fs/dosfstools](https://packages.gentoo.org/packages/sys-fs/dosfstools)[NTFS](/wiki/NTFS "NTFS")mkfs.ntfs Ano
[sys-fs/ntfs3g](https://packages.gentoo.org/packages/sys-fs/ntfs3g)

\|}

**Důležité**

The handbook recommends new partitions as part of the installation process, but it is important to note running any mkfs command will erase any data contained within the partition. When necessary, ensure any data that exists within is appropriately backed up before creating a new filesystem.

Máme-li například zaváděcí oddíl (/dev/sda1) s ext2 a kořenový oddíl (/dev/sda3) s ext4 jak je uvedeno v příkladu strukutury oddílů, použijeme následující příkazy:

`root #` `mkfs.ext4 /dev/sda3`

{{#ifeq: 0\|1\|

#### EFI system partition filesystem

The EFI system partition (/dev/sda1) must be formatted as FAT32:

`root #` `mkfs.ext2 /dev/sda1`

#### Legacy BIOS boot partition filesystem

Systems booting via legacy BIOS with a MBR/DOS disklabel can use any filesystem format supported by the bootloader.

For example, to format with XFS:

`root #` `mkfs.xfs -c options=/usr/share/xfsprogs/mkfs/lts_6.18.conf /dev/sda1`

#### Small ext4 partitions

Pokud používáte ext2, ext3 nebo ext4 na malých oddílech (méně než 8 GiB), pak jej musíte vytvořit s pomocí správných voleb, aby byl rezervován dostatek uzlů. Aplikace mke2fs (mkfs.ext2) používá nastavení "bytes-per-inode" k výpočtu množství uzlů souborového systému. U menších oddílů je radno toto číslo zvýšit.

`root #` `mkfs.ext2 -T small /dev/<device>`

`root #` `mkfs.ext3 -T small /dev/<device>`

`root #` `mkfs.ext4 -T small /dev/<device>`

To obvykle zečtyřnásobí množství uzlů daného systému souborů, jelikož hodnota "bytes-per-inode" se zmenší z každých 16kB na každé 4kB. Dále lze tuto hodnotu nastavovat poskytnutím poměru:

### Aktivace swap oddílu

mkswap je příkaz použitý k inicializaci swap oddílů:

`root #` `mkswap /dev/sda2`

**Poznámka**

Installations which were previously started, but did not finish the installation process can resume the installation from this point in the handbook. Use this link as the permalink: [Resumed installations start here](/wiki/Handbook:Alpha/Installation/Disks/cs#Resumed_installations_start_here "Handbook:Alpha/Installation/Disks/cs").

K aktivaci oddílu swap použijte příkaz swapon:

`root #` `swapon /dev/sda2`

This 'activation' step is only necessary because the swap partition is newly created within the live environment. Once the system has been rebooted, as long as the swap partition is properly defined within fstab or other mount mechanism, swap space will activate automatically.

## Připojení kořenového oddílu

Certain live environments may be missing the suggested mount point for Gentoo's root partition (/mnt/gentoo), or mount points for additional partitions created in the partitioning section:

`root #` `mkdir --parents /mnt/gentoo`

Continue creating additional mount points necessary for any additional (custom) partition(s) created during previous steps by using the mkdir command.

Teď, když jsou oddíly inicializovány a obsahují systém souborů je načase je připojit. Použijte příkaz mount, ale nezapomeňte vytvořit potřebné adresáře pro každý vytvořený oddíl. Jako příklad připojíme oddíl root:

Mount the root partition:

`root #` `mount /dev/sda3 /mnt/gentoo`

Continue mounting additional (custom) partitions as necessary using the mount command.

**Poznámka**

Pokud je nutné, aby adresář /tmp/ byl na vlastním oddíle, ujistěte se, že oprávnění budou po připojení změněna:

`root #` `chmod 1777 /mnt/gentoo/tmp`

To stejné platí pro adresář /var/tmp.

Později dojde k připojení systému souborů proc (virtuální rozhraní jádra) stejně jako dalších pseudo systémů souborů. Nejprve však [nainstalujeme instalační soubory Gentoo](/wiki/Handbook:Alpha/Installation/Stage/cs "Handbook:Alpha/Installation/Stage/cs").

[← Konfigurace sítě](/wiki/Handbook:Alpha/Installation/Networking/cs "Previous part") [Home](/wiki/Handbook:Alpha "Handbook:Alpha") [Instalace stage3 →](/wiki/Handbook:Alpha/Installation/Stage/cs "Next part")