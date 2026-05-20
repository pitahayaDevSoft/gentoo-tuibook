# Disks

Other languages:

- [Deutsch](/wiki/Handbook:X86/Installation/Disks/de "Handbuch:X86/Installation/Festplatten (100% translated)")
- [English](/wiki/Handbook:X86/Installation/Disks "Handbook:X86/Installation/Disks (100% translated)")
- [Türkçe](/wiki/Handbook:X86/Installation/Disks/tr "Handbook:X86/Installation/Disks/tr (0% translated)")
- [español](/wiki/Handbook:X86/Installation/Disks/es "Manual de Gentoo: X86/Instalación/Discos (100% translated)")
- [français](/wiki/Handbook:X86/Installation/Disks/fr "Handbook:X86/Installation/Disks/fr (100% translated)")
- [italiano](/wiki/Handbook:X86/Installation/Disks/it "Handbook:X86/Installation/Disks (100% translated)")
- [magyar](/wiki/Handbook:X86/Installation/Disks/hu "Handbook:X86/Installation/Disks/hu (100% translated)")
- [polski](/wiki/Handbook:X86/Installation/Disks/pl "Handbook:X86/Installation/Disks (100% translated)")
- [português do Brasil](/wiki/Handbook:X86/Installation/Disks/pt-br "Manual:X86/Instalação/Discos (100% translated)")
- čeština
- [русский](/wiki/Handbook:X86/Installation/Disks/ru "Handbook:X86/Installation/Disks (100% translated)")
- [தமிழ்](/wiki/Handbook:X86/Installation/Disks/ta "கையேடு:X86/நிறுவல்/வட்டுக்கள் (100% translated)")
- [中文（中国大陆）‎](/wiki/Handbook:X86/Installation/Disks/zh-cn "手册：X86/安装/准备磁盘 (100% translated)")
- [日本語](/wiki/Handbook:X86/Installation/Disks/ja "ハンドブック:X86/インストール/ディスク (100% translated)")
- [한국어](/wiki/Handbook:X86/Installation/Disks/ko "Handbook:X86/Installation/Disks/ko (100% translated)")

[X86 Handbook](/wiki/Handbook:X86/cs "Handbook:X86/cs")[Instalace](/wiki/Handbook:X86/Full/Installation/cs "Handbook:X86/Full/Installation/cs")[O instalaci](/wiki/Handbook:X86/Installation/About/cs "Handbook:X86/Installation/About/cs")[Výběr média](/wiki/Handbook:X86/Installation/Media/cs "Handbook:X86/Installation/Media/cs")[Konfigurace sítě](/wiki/Handbook:X86/Installation/Networking/cs "Handbook:X86/Installation/Networking/cs")Příprava disků[Instalace stage3](/wiki/Handbook:X86/Installation/Stage/cs "Handbook:X86/Installation/Stage/cs")[Instalace základního systému](/wiki/Handbook:X86/Installation/Base/cs "Handbook:X86/Installation/Base/cs")[Konfigurace jádra](/wiki/Handbook:X86/Installation/Kernel/cs "Handbook:X86/Installation/Kernel/cs")[Konfigurace systému](/wiki/Handbook:X86/Installation/System/cs "Handbook:X86/Installation/System/cs")[Instalace nástrojů](/wiki/Handbook:X86/Installation/Tools/cs "Handbook:X86/Installation/Tools/cs")[Konfigurace zavaděče](/wiki/Handbook:X86/Installation/Bootloader/cs "Handbook:X86/Installation/Bootloader/cs")[Dokončení](/wiki/Handbook:X86/Installation/Finalizing/cs "Handbook:X86/Installation/Finalizing/cs")[Práce s Gentoo](/wiki/Handbook:X86/Full/Working/cs "Handbook:X86/Full/Working/cs")[Úvod do Portage](/wiki/Handbook:X86/Working/Portage/cs "Handbook:X86/Working/Portage/cs")[Přepínače USE](/wiki/Handbook:X86/Working/USE/cs "Handbook:X86/Working/USE/cs")[Funkce portage](/wiki/Handbook:X86/Working/Features/cs "Handbook:X86/Working/Features/cs")[Systém initskriptů](/wiki/Handbook:X86/Working/Initscripts/cs "Handbook:X86/Working/Initscripts/cs")[Proměnné prostředí](/wiki/Handbook:X86/Working/EnvVar/cs "Handbook:X86/Working/EnvVar/cs")[Práce s Portage](/wiki/Handbook:X86/Full/Portage/cs "Handbook:X86/Full/Portage/cs")[Soubory a adresáře](/wiki/Handbook:X86/Portage/Files/cs "Handbook:X86/Portage/Files/cs")[Proměnné](/wiki/Handbook:X86/Portage/Variables/cs "Handbook:X86/Portage/Variables/cs")[Mísení softwarových větví](/wiki/Handbook:X86/Portage/Branches/cs "Handbook:X86/Portage/Branches/cs")[Doplňkové nástroje](/wiki/Handbook:X86/Portage/Tools/cs "Handbook:X86/Portage/Tools/cs")[Vlastní strom Portage](/wiki/Handbook:X86/Portage/CustomTree/cs "Handbook:X86/Portage/CustomTree/cs")[Pokročilé funkce](/wiki/Handbook:X86/Portage/Advanced/cs "Handbook:X86/Portage/Advanced/cs")[Konfigurace sítě](/wiki/Handbook:X86/Full/Networking/cs "Handbook:X86/Full/Networking/cs")[Začínáme](/wiki/Handbook:X86/Networking/Introduction/cs "Handbook:X86/Networking/Introduction/cs")[Pokročilá konfigurace](/wiki/Handbook:X86/Networking/Advanced/cs "Handbook:X86/Networking/Advanced/cs")[Modulární síťování](/wiki/Handbook:X86/Networking/Modular/cs "Handbook:X86/Networking/Modular/cs")[Bezdrátové sítě](/wiki/Handbook:X86/Networking/Wireless/cs "Handbook:X86/Networking/Wireless/cs")[Přidání funkcí](/wiki/Handbook:X86/Networking/Extending/cs "Handbook:X86/Networking/Extending/cs")[Dynamická správa](/wiki/Handbook:X86/Networking/Dynamic/cs "Handbook:X86/Networking/Dynamic/cs")

## Contents

- [1Úvod k blokovým zařízením](#.C3.9Avod_k_blokov.C3.BDm_za.C5.99.C3.ADzen.C3.ADm)
  - [1.1Bloková zařízení](#Blokov.C3.A1_za.C5.99.C3.ADzen.C3.AD)
  - [1.2Tabulky oddílů](#Tabulky_odd.C3.ADl.C5.AF)
    - [1.2.1GPT](#GPT)
    - [1.2.2MBR](#MBR)
  - [1.3Pokročilá úložiště](#Pokro.C4.8Dil.C3.A1_.C3.BAlo.C5.BEi.C5.A1t.C4.9B)
  - [1.4Výchozí schéma oddílů](#V.C3.BDchoz.C3.AD_sch.C3.A9ma_odd.C3.ADl.C5.AF)
- [2Návrh rozdělení oddílů](#N.C3.A1vrh_rozd.C4.9Blen.C3.AD_odd.C3.ADl.C5.AF)
  - [2.1Kolik oddílů a jak velkých?](#Kolik_odd.C3.ADl.C5.AF_a_jak_velk.C3.BDch.3F)
  - [2.2A co swap?](#A_co_swap.3F)
    - [2.2.1Použití UEFI](#Pou.C5.BEit.C3.AD_UEFI)
  - [2.3Co je to bootovací oddíl BIOSu?](#Co_je_to_bootovac.C3.AD_odd.C3.ADl_BIOSu.3F)
- [3Alternativa: Použití fdisku k rozdělení disku](#Alternativa:_Pou.C5.BEit.C3.AD_fdisku_k_rozd.C4.9Blen.C3.AD_disku)
  - [3.1Prohlížení současného rozdělení oddílů](#Prohl.C3.AD.C5.BEen.C3.AD_sou.C4.8Dasn.C3.A9ho_rozd.C4.9Blen.C3.AD_odd.C3.ADl.C5.AF)
  - [3.2Odstranění všech oddílů s fdiskem](#Odstran.C4.9Bn.C3.AD_v.C5.A1ech_odd.C3.ADl.C5.AF_s_fdiskem)
  - [3.3Tvorba bootovacího oddílu BIOSu](#Tvorba_bootovac.C3.ADho_odd.C3.ADlu_BIOSu)
  - [3.4Vytvoření swap oddílu](#Vytvo.C5.99en.C3.AD_swap_odd.C3.ADlu)
  - [3.5Tvorba bootovacího oddílu](#Tvorba_bootovac.C3.ADho_odd.C3.ADlu)
  - [3.6Vytvoření kořenového oddílu](#Vytvo.C5.99en.C3.AD_ko.C5.99enov.C3.A9ho_odd.C3.ADlu)
  - [3.7Uložení rozložení oddílů](#Ulo.C5.BEen.C3.AD_rozlo.C5.BEen.C3.AD_odd.C3.ADl.C5.AF)
- [4Partitioning the disk with MBR for BIOS / legacy boot](#Partitioning_the_disk_with_MBR_for_BIOS_.2F_legacy_boot)
  - [4.1Viewing the current partition layout](#Viewing_the_current_partition_layout)
  - [4.2Creating a new disklabel / removing all partitions](#Creating_a_new_disklabel_.2F_removing_all_partitions)
  - [4.3Creating the boot partition](#Creating_the_boot_partition)
  - [4.4Creating the swap partition](#Creating_the_swap_partition)
  - [4.5Creating the root partition](#Creating_the_root_partition)
  - [4.6Saving the partition layout](#Saving_the_partition_layout)
- [5Vytvoření systému souborů](#Vytvo.C5.99en.C3.AD_syst.C3.A9mu_soubor.C5.AF)
  - [5.1Úvod](#.C3.9Avod)
  - [5.2Systémy souborů](#Syst.C3.A9my_soubor.C5.AF)
  - [5.3Aplikace systému souborů na diskový oddíl](#Aplikace_syst.C3.A9mu_soubor.C5.AF_na_diskov.C3.BD_odd.C3.ADl)
    - [5.3.1EFI system partition filesystem](#EFI_system_partition_filesystem)
    - [5.3.2Legacy BIOS boot partition filesystem](#Legacy_BIOS_boot_partition_filesystem)
    - [5.3.3Small ext4 partitions](#Small_ext4_partitions)
  - [5.4Aktivace swap oddílu](#Aktivace_swap_odd.C3.ADlu)
- [6Připojení kořenového oddílu](#P.C5.99ipojen.C3.AD_ko.C5.99enov.C3.A9ho_odd.C3.ADlu)

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

### Tabulky oddílů

Ačkoli je teoreticky možné použít surový, nerozdělený disk k umístění Linuxového systému (např. při tvorbě btrfs RAID), téměř nikdy se tak v praxi neděje. Namísto toho se disky rozdělují na menší, lépe spravovatelná bloková zařízení. V případě systémů **x86** se jim říká oddíly. V současnosti se používají dvě standardní technologie pro dělení oddílů: MBR a GPT.

#### GPT

"GTP (GUID Partition Table)" používá 64bitové identifikátory oddílů. Místo, kde se ukládají informace o oddílech, je také o mnoho větší, než 512 bytů u MBR, což znamená, že zde není prakticky žádné omezení počtu oddílů u GPT disků. Velikost oddílů je pak vázána mnohem vyšším limitem (téměř 8 ZiB - ano, zettabytů).

Pokud je systémovým softwarovým rozhraním mezi operačním systémem a firmwarem UEFI (namísto BIOSu), GPT je téměř nezbytné, jinak se objeví problémy kompatibility s MBR.

GPT také využívá výhod kontrolních součtů a redundance. Obsahuje CRC32 kontrolní součty k odhalování chyb v hlavičce a tabulce oddílů a na konci disku má zálohu GPT. Tato záložní tabulka může být využita k opravě poškozeného primárního GPT poblíž začátku disku.

**Důležité**

There are a few caveats regarding GPT:

- Using GPT on a BIOS-based computer works, but the user won't be able to dual-boot with a Microsoft Windows operating system, since Microsoft Windows refuses to boot from a GPT partition when in BIOS mode.
- Some buggy (old) motherboard firmware configured to boot in BIOS/CSM/legacy mode might also have problems with booting from GPT labeled disks.

#### MBR

"MBR (Master Boot Record") používá 32bitové identifikátory startovního sektoru a délky oddílů a podporuje tři typy oddílů: primární, rozšířený a logický. Informace o primárního oddílech jsou uloženy v master boot record samotném - velmi malé (obvykle 512 bytů) oblasti na začátku disku. V důsledku tohoto malého prostoru jsou podporovány pouze čtyři primární oddíly (například /dev/sda1 až /dev/sda4).

Pro podporu více oddílů je nutné označit jeden z primárních oddílů jak rozšířený oddíl. Tento oddíl pak může obsahovat logické oddíly (oddíly v oddílech).

**Důležité**

Ačkoli je většina výrobců základních desek dosud podporuje, tabulky oddílů jsou považovány za zastaralé. Pokud nepoužíváte hardware s původem před rokem 2010, je lepší rozdělit disk pomocí [GUID Partition Table](https://en.wikipedia.org/wiki/GUID_Partition_Table "wikipedia:GUID Partition Table"). Čtenáři, kteří jsou nuceni použít MBR, by měli vzít v potaz následující informace:

- Většina základních desek z obodobí po 2010 považuje MBR za zastaralý (podporvaný, ale ne ideální) mód zavádění.
- Kvůli použití 32bitových identifikátorů nelze tabulky oddílů master boot record použít pro disky s velikostí větší než 2 TiB.
- Pokud není použit rozšířený oddíl, MBR podporuje pouze 4 diskové oddíly.
- MBR neposkytuje žádnou zálohu, pokud tedy aplikace nebo uživatel přepíše MBR, veškeré informace o oddílech se ztratí.

Autoři příručky doporučují použití [GTP](#GTP) pro všechny instalace Gentoo, kde je to možné.

### Pokročilá úložiště

Instalační CD **x86** poskytuje podporu Správce logických oddílů (LVM). LVM zvyšuje flexibilitu při nastavování rozdělení disku. Níže uvedené instalační instrukce se budou soustředit na "běžné" oddíly, nicméně je dobré vědět, že LVM je podporováno, pokud se toužíte vydat touto cestou. Přečtěte si článek o [LVM](/wiki/LVM "LVM") ohledně detailů. Nováčci dejte si pozor: ačkoli je LVM plně podporováno, je mimo záběr tohoto průvodce.

Although usage is not covered in the handbook, below is a list helpful guides to get the system running:

- [Btrfs](/wiki/Btrfs/Native_System_Root_Guide "Btrfs/Native System Root Guide")
- [ZFS](/wiki/ZFS/rootfs "ZFS/rootfs")
- [LVM (Logical Volume Manager)](/wiki/LVM "LVM")

Disk encryption is also handled in the same manner:

- [Rootfs encryption](/wiki/Rootfs_encryption "Rootfs encryption")

### Výchozí schéma oddílů

Throughout the remainder of the handbook, we will discuss and explain two cases:

1. UEFI firmware with GUID Partition Table (GPT) disk.
2. MBR DOS/legacy BIOS firmware with a MBR partition table disk.

While it is possible to mix and match boot types with certain motherboard firmware, mixing goes beyond the intention of the handbook. As previously stated, it is strongly recommended for installations on modern hardware to use UEFI boot with a GPT disklabel disk.

Po celý zbytek této příručky bude využíváno následující schéma oddílů jako jednoduchý příklad rozdělení:

**Důležité**

The first row of the following table contains exclusive information for _**either**_ a GPT disklabel _**or**_ a MBR DOS/legacy BIOS disklabel. When in doubt, proceed with GPT, since **x86** machines manufactured after the year 2010 generally support UEFI firmware and GPT boot sector.

Oddíl
Systém souborů
Velikost
Popis
/dev/sda1(zavaděč)
2M
bootovací oddíl BIOsu
/dev/sda1ext2 (nebo fat32, pokud používáte UEFI)
128M
Zaváděcí oddíl/Systémový oddíl EFI
/dev/sda3(swap)
512M nebo více
swap oddíl
/dev/sda4ext4
Zbytek disku
kořenový oddíl

Pokud je to dostatečné a čtenář se vydá cestou GPT, může okamžitě přejít na [Výchozí: Použití parted k rozdělení disku](#Vychozi:_Pouziti_parted_k_rozdeleni_disku). Ti, kteří se stále zajímají o MBR (hej - stává se to!) a používají výchozí rozdělení, mohou přeskočit na [Alternativa: Použití fdisku k rozdělení disku](#Alternativa:_pouziti_fdisku_k_rozdeleni_disku).

fdisk a parted jsou obě utility k rozdělení disku. fdisk je dobře známá, stabilní, doporučená k rozdělení podle MBR, zatímco parted byla jedna z prvních utilit pro správu oddílů podporujících GTP oddíly. Ti, kterým se líbí rozhraní programu fdisk mohou použít gdisk (GPT disk) jako alternativu k parted.

Před uvedením instrukcí k tvorbě se první části budou věnovat detailnějšímu popisu toho, jak mohou být oddíly tvořeny a zmíní nejčastější problémy.

## Návrh rozdělení oddílů

### Kolik oddílů a jak velkých?

Návrh rozdělení diskových oddílů je značně závislý na požadavcích na systém. V případě velkého množství uživatelů je radno mít adresář /home/ na samostatném oddílu, což zvýší bezpečnost a usnadní zálohování a další údržbu. Pokud je Gentoo instalováno jako e-mailový server, pak by měl být adresář /var na samostatném oddíle, jelikož e-maily jsou uloženy uvnitř adresáře /var/. Správná volba systému souborů pak může maximalizovat výkon. Herní servery budou mít samostatný oddíl /opt/, jelikož většina herního software se instaluje tam. Důvod je stejný jako u adresáře /home/: bezpečnost a zálohy.

Ve většině případů by měly být na Gentoo /usr a /var co velikosti relativně rozsáhlé. /usr obsahuje většinu aplikací v systému a zdrojové kódy jádra Linux v adresáři /usr/src. Ve výchozím stavu se v adresáři /var nachází repozitář ebuildů (umístěný v /var/db/repos/gentoo), který v závislosti na souborovém systému obvykle zabírá okolo 650 MB protoru na disku. Tento odhad nezahrnuje adresáře /var/cache/distfiles a /var/cache/binpkgs, které se budou postupně zaplňovat zdrojovými kódy a (volitelně) binárními balíčky, tak jak budou postupně přidávány do systému.

Volba počtu oddílů a jejich velikosti závisí nejvíce na úvaze o výhodách a nevýhodách daného řešení a výběru nejlepší varianty podle okolností. Oddělené oddíly nebo svazky mají následující výhody:

- Vyberete nejlepší systém souborů pro každý oddíl nebo svazek
- Systému nemůže dojít volné místo na disku, pokud bude jediný nefunkční nástroj neustále zapisovat soubory na oddíl nebo svazek
- Kontrola systému souborů je v případě potřeby časově zkrácena, jelikož kontroly lze činit paralelně (ačkoli tato výhoda se týká spíše více disků než více oddílů).
- Bezpečnost může být zvýšena připojením některých oddílů v režimu pouze pro čtení, s volbou `nosuid` (setuid bity jsou ignorovány), `noexec` (spouštěcí bity jsou ignorovány) atd.

Nicméně velké množství oddílů má také své nevýhody:

- Při špatné konfiguraci může mít systém mnoho volného místo na jednom oddíle a málo volného místa na jiném.
- Další nepříjemností je, že je oddělené oddíly - zvláště důležité přípojné body jako /usr/ a /var/ \- často vyžadují aby správce při spuštění používal initramfs k připojení oddílů, než se spustí ostatní startovací skripty. Nemusí tomu tak však být vždy, tudíž se výsledky mohou lišit.
- Dále existuje u SATA a SCSI limit 15 oddílů, pokud disk nepoužívá GPT popisovače.

**Poznámka**

Installations that intend to use systemd as the service and init system must have the /usr directory available at boot, either as part of the root filesystem or mounted via an initramfs.

### A co swap?

Recommendations for swap space size
RAM sizeSuspend support?Hibernation support?
2 GB or less4GB4GB
2 to 8 GBRAM amount2 \* RAM
8 to 64 GB8 GB minimum, 16 maximum1.5 \* RAM
64 GB or greater8 GB minimumHibernation not recommended! Hibernation is _not_ recommended for systems with very large amounts of memory. While possible, the entire contents of memory must be written to disk in order to successfully hibernate. Writing tens of gigabytes (or worse!) out to disk can can take a considerable amount of time, especially when rotational disks are used. It is best to suspend in this scenario.

Neexistuje perfektní hodnota pro swap oddíl. Účelem swapu je poskytnout úložiště na disku v případě, že interní paměť jádra je pod tlakem. Swap umožňuje jádru přesunout stránky paměti, které nejspíše nebudou v krátkém časovém okamžiku potřebné na disk (swap nebo page-out) a paměť tak uvolnit. Samozřejmě pokud je tato paměť opět potřeba, je nutné tyto stránky opět vložit zpět do paměti (page-in), což zabere nějaký čas (jelikož disky jsou oproti vnitřní paměti pomalé).

Pokud systém nebude spouštět aplikace náročné na paměť nebo pokud má k dispozici hodně paměti, pak pravděpodobně nepotřebuje mnoho prostoru pro swap. Swap je však také používán pro uložení veškeré paměti v případě uspání. Pokud systém bude potřebovat uspávat, pak bude potřeba většího prostoru pro swap, obvykle nejméně v množství paměti instalované v systému.

#### Použití UEFI

Pokud instalujete Gentoo na systém, který používá k bootování operačního systému UEFI (namísto BIOSu), je důležité vytvořit systémový oddíl EFI (ESP). Níže uvedené instrukce pro parted obsahují vodítka nezbytná ke zdárnému provedení této operace.

ESP dále musí být variantou oddílu FAT (někdy zobrazovaného na Linuxových systémech jako _vfat_). Oficiální [specifikace UEFI](http://www.uefi.org/sites/default/files/resources/UEFI%202_5.pdf) zmiňuje že firmware rozpozná systém souborů FAT12, 16 nebo 32, nicméně FAT32 je u ESP doporučený. Pokračujte naformátováním ESP na FAT32:

`root #` `mkfs.vfat -F 32 /dev/sda1`

**Důležité**

Pokud pro oddíl boot nepoužijete variantu FAT, pak nelze garantovat že firmware UEFI vašeho sytému bude schopen najít zavaděč (nebo jádro Linux) a nejspíše nebude schopný spustit systém!

### Co je to bootovací oddíl BIOSu?

Bootvací oddíl BIOSu je velmi malý (1 až 2 MB) oddíl, do něhož mohou zavaděče jako GRUB2 umístit dodatečná data, která se nevejdou do přiděleného úložiště (několik stovek bytů v případě MBR) a která nemohou být umístěna jinam.

## Alternativa: Použití fdisku k rozdělení disku

Následující části vysvětlují jak vytvořit příkladmé rozdělení disku pomocí příkazu fdisk. Příkladmé rozdělení oddílů bylo zmíněno dříve:

Rozložení oddílů si změňte podle osobních preferencí.

Oddíl
Popis
/dev/sda1zaváděcí oddíl BIOSu
/dev/sda1zaváděcí oddíl
/dev/sda3oddíl swap
/dev/sda4kořenový oddíl

### Prohlížení současného rozdělení oddílů

fdisk je populární a silný nástroj k rozdělení disku na oddíly. Spusťte jej nad diskem (v našem případě použijeme /dev/sda):

`root #` `fdisk /dev/sda`

Použijte `p` k zobrazení aktuální konfigurace rozdělení disku:

`Command (m for help):` `p`

```
Disk /dev/sda: 240 heads, 63 sectors, 2184 cylinders
Units = cylinders of 15120 * 512 bytes

   Device Boot    Start       End    Blocks   Id  System
/dev/sda1   *         1        14    105808+  83  Linux
/dev/sda2            15        49    264600   82  Linux swap
/dev/sda3            50        70    158760   83  Linux
/dev/sda4            71      2184  15981840    5  Extended
/dev/sda5            71       209   1050808+  83  Linux
/dev/sda6           210       348   1050808+  83  Linux
/dev/sda7           349       626   2101648+  83  Linux
/dev/sda8           627       904   2101648+  83  Linux
/dev/sda9           905      2184   9676768+  83  Linux

```

Device Start End Sectors Size Type
/dev/sda1 2048 2099199 2097152 1G EFI System
/dev/sda2 2099200 10487807 8388608 4G Linux swap
/dev/sda3 10487808 1953523711 1943035904 926.5G Linux root (x86-64)

}}

Tento konkrétní disk byl nastaven tak, aby obsahoval sedm linuxových souborů systémů (každý z odpovídajícím oddílem označeným jako "Linux") a oddíl swap (označený jako "Linux swap").

### Odstranění všech oddílů s fdiskem

Pressing the `g` key will instantly remove all existing disk partitions and create a new GPT disklabel:

`Command (m for help):` `g`

```
Created a new GPT disklabel (GUID: 3E56EE74-0571-462B-A992-9872E3855D75).

```

Nejprve odstraníme všechny oddíly z disku. Pro vymazání oddílu stiskněte `d`. Vymazání existujícího oddílu /dev/sda1:

`Command (m for help):` `d`

```
Partition number (1-4): 1

```

Oddíl byl nyní připraven k vymazání. Nadále se nebude zobrazovat ve výpisu oddílů ( `p`), ale nebude odstraněn do doby, než budou změny uloženy. To umožňuje uživatelům ukončit operaci, pokud udělali chybu - v takovém případě stiskněte ihned `Enter` a enter a oddíl nebude vymazán.

Opakovaně stiskněte `p` pro zobrazení výpisu oddílů a stiskněte `d` a číslo oddílu k vymazání. Nakonec bude tabulka oddílů prázdná:

`Command (m for help):` `p`

```
Disk /dev/sda: 30.0 GB, 30005821440 bytes
240 heads, 63 sectors/track, 3876 cylinders
Units = cylinders of 15120 * 512 = 7741440 bytes

Device Boot    Start       End    Blocks   Id  System

```

Nyní, když je tabulka oddílů v paměti prázdná, můžeme vytvářet oddíly.

### Tvorba bootovacího oddílu BIOSu

**Poznámka**

A smaller ESP is possible but not recommended, especially given it may be shared with other OSes.

Nejprve vytvořte malý bootovací oddíl BIOSu. Pro vytvoření nového oddílu stiskněte `n`, potom `p` pro vytvoření primárního oddílu, následně `1` ke zvolení prvního primárního oddílu. Na dotaz ohledně prvního sektoru se ujistěte, že bude začínat na 2048 (což je třeba pro zavaděč) a stiskněte `Enter`. Na dotaz ohledně posledního sektoru napište +2M k vytvoření oddílu velkého 2 MB.

`Command (m for help):` `n`

```
Command action
  e   extended
  p   primary partition (1-4)
p
Partition number (1-4): 1
First sector (64-10486533532, default 64): 2048
Last sector, +sectors +size{M,K,G} (4096-10486533532, default 10486533532): +2M

```

Do you want to remove the signature? \[Y\]es/\[N\]o: Y
The signature will be removed by a write command.

}}

Označte oddíl pro účely UEFI:

`Command (m for help):` `t`

```
Selected partition 1
Hex code (type L to list codes): 4
Changed system type of partition 1 to 4 (BIOS boot)

```

### Vytvoření swap oddílu

Pro vytvoření swap oddílu stiskněte `n` pro vytvoření nového oddílu, poté `p` pro vytvoření primárního oddílu. Po té stiskněte `3` pro vytvoření třetího primárního oddílu, /dev/sda3. Na dotaz ohledně prvního sektoru stiskněte `Enter`. Na dotaz ohledně posledního sektoru napište +512M (nebo jinou hodnotu dle potřebného prostoru pro swap) k vytvoření oddílu velkého 512 MB.

### Tvorba bootovacího oddílu

Po tomto všem stiskněte `t` k nastavení typu oddílu, stisknutím `3` vyberete právě vytvořený oddíl a po té vepiště "82" pro nastavení typu na "Linux Swap".

`Command (m for help):` `t`

```
Partition number (1,2, default 2): 2
Partition type or alias (type L to list all): 19

Changed type of partition 'Linux filesystem' to 'Linux swap'.

```

### Vytvoření kořenového oddílu

Nakonec vytvoříme kořenový oddíl, k vytvoření nového oddílu stiskněte `n`, po té `p` pro vytvoření primárního oddílu. Po té stiskněte `4` k vytvoření čtvrtého primárního oddílu /dev/sda4. Na dotaz ohledně prvního sektoru stiskněte `Enter`. Na dotaze ohledně posledního sektoru, stiskněte `Enter`, čímž vytvoříte zaplníte zbývající prostor na disku. Po dokončení těchto kroků by mělo stisknutí `p` zobrazit tabulku oddílu, která bude vypadat podobně jako takto:

`Command (m for help):` `n`

```
Partition number (3-128, default 3): 3
First sector (10487808-1953525134, default 10487808):
Last sector, +/-sectors or +/-size{K,M,G,T,P} (10487808-1953525134, default 1953523711):
</div>

<div lang="en" dir="ltr" class="mw-content-ltr">
Created a new partition 3 of type 'Linux filesystem' and of size 926.5 GiB..

```

**Poznámka**

Setting the root partition's type to "Linux root (x86-64)" is not required and the system will function normally if it is set to the "Linux filesystem" type. This filesystem type is only necessary for cases where a bootloader that supports it (i.e. systemd-boot) is used and a fstab file is not wanted.

After creating the root partition, press `t` to set the partition type, `3` to select the partition just created, and then type in _23_ to set the partition type to "Linux Root (x86-64)".

`Command(m for help):` `t`

```
Partition number (1-3, default 3): 3
Partition type or alias (type L to list all): 23
</div>

<div lang="en" dir="ltr" class="mw-content-ltr">
Changed type of partition 'Linux filesystem' to 'Linux root (x86-64)'

```

After completing these steps, pressing `p` should display a partition table that looks similar to the following:

`Command (m for help):` `p`

```
Disk /dev/sda: 30.0 GB, 30005821440 bytes
240 heads, 63 sectors/track, 3876 cylinders
Units = cylinders of 15120 * 512 = 7741440 bytes

   Device Boot    Start       End    Blocks   Id  System
/dev/sda1             1         3      5198+  ef  EFI (FAT-12/16/32)
/dev/sda2   *         3        14    105808+  83  Linux
/dev/sda3            15        81    506520   82  Linux swap
/dev/sda4            82      3876  28690200   83  Linux

```

Device Start End Sectors Size Type
/dev/sda1 2048 2099199 2097152 1G EFI System
/dev/sda2 2099200 10487807 8388608 4G Linux swap
/dev/sda3 10487808 1953523711 1943035904 926.5G Linux root (x86-64)

Filesystem/RAID signature on partition 1 will be wiped.

}}

### Uložení rozložení oddílů

K uložení rozdělení oddílů disků a ukončení programu fdisk stiskněte `w`.

`Command (m for help):` `w`

Po té, co jsme vytvořili oddíly, je čas umístit na ně systém souborů.

## Partitioning the disk with MBR for BIOS / legacy boot

The following table provides a recommended partition layout for a trivial MBR DOS / legacy BIOS boot installation. Additional partitions can be added according to personal preference or system design goals.

Device path (sysfs)
Mount point
File system
DPS UUID (PARTUUID)
Description
/dev/sda1/bootext4
N/A
MBR DOS / legacy BIOS boot partition details.
/dev/sda2N/A. Swap is not mounted to the filesystem like a device file.0657fd6d-a4ab-43c4-84e5-0933c84b4f4f
Swap partition details.
/dev/sda3/xfs
44479540-f297-41b2-9af7-d131d5f0458a
Root partition details.

Change the partition layout according to personal preference.

### Viewing the current partition layout

Fire up fdisk against the disk (in our example, we use /dev/sda):

`root #` `fdisk /dev/sda`

Use the `p` key to display the disk's current partition configuration:

`Command (m for help):` `p`

```
Disk /dev/sda: 931.51 GiB, 1000204886016 bytes, 1953525168 sectors
Disk model: HGST HTS721010A9
Units: sectors of 1 * 512 = 512 bytes
Sector size (logical/physical): 512 bytes / 4096 bytes
I/O size (minimum/optimal): 4096 bytes / 4096 bytes
Disklabel type: dos
Disk identifier: 0xf163b576
</div>

<div lang="en" dir="ltr" class="mw-content-ltr">
Device     Boot    Start        End    Sectors   Size Id Type
/dev/sda1  *        2048    2099199    2097152     1G 83 Linux
/dev/sda2        2099200   10487807    8388608     4G 82 Linux swap / Solaris
/dev/sda3       10487808 1953525167 1943037360 926.5G 83 Linux

```

This particular disk was until now configured to house two Linux filesystems (each with a corresponding partition listed as "Linux") as well as a swap partition (listed as "Linux swap"), using a GPT table.

### Creating a new disklabel / removing all partitions

Pressing `o` will instantly remove all existing disk partitions and create a new MBR disklabel (also named DOS disklabel):

`Command (m for help):` `o`

```
Created a new DOS disklabel with disk identifier 0xf163b576.
The device contains 'gpt' signature and it will be removed by a write command. See fdisk(8) man page and --wipe option for more details.

```

Alternatively, to keep an existing DOS disklabel (see the output of `p` above), consider removing the existing partitions one by one from the disk. Press `d` to delete a partition. For instance, to delete an existing /dev/sda1:

`Command (m for help):` `d`

```
Partition number (1-4): 1

```

The partition has now been scheduled for deletion. It will no longer show up when printing the list of partitions ( `p`, but it will not be erased until the changes have been saved. This allows users to abort the operation if a mistake was made - in that case, type `q` immediately and hit `Enter` and the partition will not be deleted.

Repeatedly press `p` to print out a partition listing and then press `d` and the number of the partition to delete it. Eventually, the partition table will be empty:

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

The disk is now ready to create new partitions.

### Creating the boot partition

First, create a small partition which will be mounted as /boot. Press `n` to create a new partition, followed by `p` for a _primary_ partition and `1` to select the first primary partition. When prompted for the first sector, make sure it starts from 2048 (which may be needed for the boot loader) and press `Enter`. When prompted for the last sector, type +1G to create a partition 1 GB in size:

`Command (m for help):` `n`

```
Partition type
   p   primary (0 primary, 0 extended, 4 free)
   e   extended (container for logical partitions)
Select (default p): p
Partition number (1-4, default 1): 1
First sector (2048-1953525167, default 2048):
Last sector, +/-sectors or +/-size{K,M,G,T,P} (2048-1953525167, default 1953525167): +1G
</div>

<div lang="en" dir="ltr" class="mw-content-ltr">
Created a new partition 1 of type 'Linux' and of size 1 GiB.

```

Mark the partition as bootable by pressing the `a` key and pressing `Enter`:

`Command (m for help):` `a`

```
Selected partition 1
The bootable flag on partition 1 is enabled now.

```

Note: if more than one partition is available on the disk, then the partition to be flagged as bootable will have to be selected.

### Creating the swap partition

Next, to create the swap partition, press `n` to create a new partition, then `p`, then type `2` to create the second primary partition, /dev/sda2. When prompted for the first sector, press `Enter`. When prompted for the last sector, type +4G (or any other size needed for the swap space) to create a partition 4GB in size.

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

After all this is done, press `t` to set the partition type, `2` to select the partition just created and then type in _82_ to set the partition type to "Linux Swap".

`Command (m for help):` `t`

```
Partition number (1,2, default 2): 2
Hex code (type L to list all codes): 82
</div>

<div lang="en" dir="ltr" class="mw-content-ltr">
Changed type of partition 'Linux' to 'Linux swap / Solaris'.

```

### Creating the root partition

Finally, to create the root partition, press `n` to create a new partition. Then press `p` and `3` to create the third primary partition, /dev/sda3. When prompted for the first sector, hit `Enter`. When prompted for the last sector, hit `Enter` to create a partition that takes up the rest of the remaining space on the disk:

`Command (m for help):` `n`

```
Partition type
   p   primary (2 primary, 0 extended, 2 free)
   e   extended (container for logical partitions)
Select (default p): p
Partition number (3,4, default 3): 3
First sector (10487808-1953525167, default 10487808):
Last sector, +/-sectors or +/-size{K,M,G,T,P} (10487808-1953525167, default 1953525167):
</div>

<div lang="en" dir="ltr" class="mw-content-ltr">
Created a new partition 3 of type 'Linux' and of size 926.5 GiB.

```

After completing these steps, pressing `p` should display a partition table that looks similar to this:

`Command (m for help):` `p`

```
Disk /dev/sda: 931.51 GiB, 1000204886016 bytes, 1953525168 sectors
Disk model: HGST HTS721010A9
Units: sectors of 1 * 512 = 512 bytes
Sector size (logical/physical): 512 bytes / 4096 bytes
I/O size (minimum/optimal): 4096 bytes / 4096 bytes
Disklabel type: dos
Disk identifier: 0xf163b576
</div>

<div lang="en" dir="ltr" class="mw-content-ltr">
Device     Boot    Start        End    Sectors   Size Id Type
/dev/sda1  *        2048    2099199    2097152     1G 83 Linux
/dev/sda2        2099200   10487807    8388608     4G 82 Linux swap / Solaris
/dev/sda3       10487808 1953525167 1943037360 926.5G 83 Linux

```

### Saving the partition layout

Press `w` to write the partition layout and exit fdisk:

`Command (m for help):` `w`

```
The partition table has been altered.
Calling ioctl() to re-read partition table.
Syncing disks.

```

Now it is time to put filesystems on the partitions.

## Vytvoření systému souborů

**Upozornění**

When using an SSD or NVMe drive, it is wise to check for firmware upgrades. Some Intel SSDs in particular (600p and 6000p) require a firmware upgrade for [possible data corruption](https://bugzilla.redhat.com/show_bug.cgi?id=1402533) induced by XFS I/O usage patterns. The problem is at the firmware level and not any fault of the XFS filesystem. The smartctl utility can help check the device model and firmware version.

### Úvod

Po vytvoření oddílů na disku je načase vytvořit na nich systém souborů. V následující sekci si popíšeme různé systémy souborů, které Linux podporuje. Čtenáři, kteří už vědí, jaký systém souborů použít, mohou pokračovat na [aplikaci systému souborů na diskový oddíl](#Aplikace_syst.C3.A9mu_soubor.C5.AF_na_diskov.C3.BD_odd.C3.ADl). Ostatní by si měli přečíst, jaké systémy souborů jsou pro ně k dispozici.

### Systémy souborů

Linux podporuje tucty souborových systémů, ačkoli většinu z nich je radno používat pro specifické účely. Pouze některé souborové systémy mohou být považovány na x86 architečktuře za stabilní. Radíme proto přečist si informace o systémech souborů a stavu jejich podpory před tím, než zvolíte některý experimentálnější z nich pro důležitý diskový oddíl.

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

Installations which were previously started, but did not finish the installation process can resume the installation from this point in the handbook. Use this link as the permalink: [Resumed installations start here](/wiki/Handbook:X86/Installation/Disks/cs#Resumed_installations_start_here "Handbook:X86/Installation/Disks/cs").

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

Později dojde k připojení systému souborů proc (virtuální rozhraní jádra) stejně jako dalších pseudo systémů souborů. Nejprve však [nainstalujeme instalační soubory Gentoo](/wiki/Handbook:X86/Installation/Stage/cs "Handbook:X86/Installation/Stage/cs").

[← Konfigurace sítě](/wiki/Handbook:X86/Installation/Networking/cs "Previous part") [Home](/wiki/Handbook:X86 "Handbook:X86") [Instalace stage3 →](/wiki/Handbook:X86/Installation/Stage/cs "Next part")