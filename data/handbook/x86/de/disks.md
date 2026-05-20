# Disks

Other languages:

- Deutsch
- [English](/wiki/Handbook:X86/Installation/Disks "Handbook:X86/Installation/Disks (100% translated)")
- [Türkçe](/wiki/Handbook:X86/Installation/Disks/tr "Handbook:X86/Installation/Disks/tr (0% translated)")
- [español](/wiki/Handbook:X86/Installation/Disks/es "Manual de Gentoo: X86/Instalación/Discos (100% translated)")
- [français](/wiki/Handbook:X86/Installation/Disks/fr "Handbook:X86/Installation/Disks/fr (100% translated)")
- [italiano](/wiki/Handbook:X86/Installation/Disks/it "Handbook:X86/Installation/Disks (100% translated)")
- [magyar](/wiki/Handbook:X86/Installation/Disks/hu "Handbook:X86/Installation/Disks/hu (100% translated)")
- [polski](/wiki/Handbook:X86/Installation/Disks/pl "Handbook:X86/Installation/Disks (100% translated)")
- [português do Brasil](/wiki/Handbook:X86/Installation/Disks/pt-br "Manual:X86/Instalação/Discos (100% translated)")
- [čeština](/wiki/Handbook:X86/Installation/Disks/cs "Handbook:X86/Installation/Disks/cs (50% translated)")
- [русский](/wiki/Handbook:X86/Installation/Disks/ru "Handbook:X86/Installation/Disks (100% translated)")
- [தமிழ்](/wiki/Handbook:X86/Installation/Disks/ta "கையேடு:X86/நிறுவல்/வட்டுக்கள் (100% translated)")
- [中文（中国大陆）‎](/wiki/Handbook:X86/Installation/Disks/zh-cn "手册：X86/安装/准备磁盘 (100% translated)")
- [日本語](/wiki/Handbook:X86/Installation/Disks/ja "ハンドブック:X86/インストール/ディスク (100% translated)")
- [한국어](/wiki/Handbook:X86/Installation/Disks/ko "Handbook:X86/Installation/Disks/ko (100% translated)")

[X86 Handbuch](/wiki/Handbook:X86/de "Handbook:X86/de")[Installation](/wiki/Handbook:X86/Full/Installation/de "Handbook:X86/Full/Installation/de")[Über die Installation](/wiki/Handbook:X86/Installation/About/de "Handbook:X86/Installation/About/de")[Auswahl des Mediums](/wiki/Handbook:X86/Installation/Media/de "Handbook:X86/Installation/Media/de")[Konfiguration des Netzwerks](/wiki/Handbook:X86/Installation/Networking/de "Handbook:X86/Installation/Networking/de")Vorbereiten der Festplatte(n)[Installation des Stage Archivs](/wiki/Handbook:X86/Installation/Stage/de "Handbook:X86/Installation/Stage/de")[Installation des Basissystems](/wiki/Handbook:X86/Installation/Base/de "Handbook:X86/Installation/Base/de")[Konfiguration des Kernels](/wiki/Handbook:X86/Installation/Kernel/de "Handbook:X86/Installation/Kernel/de")[Konfiguration des Systems](/wiki/Handbook:X86/Installation/System/de "Handbook:X86/Installation/System/de")[Installation der Tools](/wiki/Handbook:X86/Installation/Tools/de "Handbook:X86/Installation/Tools/de")[Konfiguration des Bootloaders](/wiki/Handbook:X86/Installation/Bootloader/de "Handbook:X86/Installation/Bootloader/de")[Abschluss](/wiki/Handbook:X86/Installation/Finalizing/de "Handbook:X86/Installation/Finalizing/de")[Arbeiten mit Gentoo](/wiki/Handbook:X86/Full/Working/de "Handbook:X86/Full/Working/de")[Portage-Einführung](/wiki/Handbook:X86/Working/Portage/de "Handbook:X86/Working/Portage/de")[USE-Flags](/wiki/Handbook:X86/Working/USE/de "Handbook:X86/Working/USE/de")[Portage-Features](/wiki/Handbook:X86/Working/Features/de "Handbook:X86/Working/Features/de")[Initskript-System](/wiki/Handbook:X86/Working/Initscripts/de "Handbook:X86/Working/Initscripts/de")[Umgebungsvariablen](/wiki/Handbook:X86/Working/EnvVar/de "Handbook:X86/Working/EnvVar/de")[Arbeiten mit Portage](/wiki/Handbook:X86/Full/Portage/de "Handbook:X86/Full/Portage/de")[Dateien und Verzeichnisse](/wiki/Handbook:X86/Portage/Files/de "Handbook:X86/Portage/Files/de")[Variablen](/wiki/Handbook:X86/Portage/Variables/de "Handbook:X86/Portage/Variables/de")[Mischen von Softwarezweigen](/wiki/Handbook:X86/Portage/Branches/de "Handbook:X86/Portage/Branches/de")[Zusätzliche Tools](/wiki/Handbook:X86/Portage/Tools/de "Handbook:X86/Portage/Tools/de")[Eigener Portage-Tree](/wiki/Handbook:X86/Portage/CustomTree/de "Handbook:X86/Portage/CustomTree/de")[Erweiterte Portage-Features](/wiki/Handbook:X86/Portage/Advanced/de "Handbook:X86/Portage/Advanced/de")[Netzwerk-Konfiguration](/wiki/Handbook:X86/Full/Networking/de "Handbook:X86/Full/Networking/de")[Zu Beginn](/wiki/Handbook:X86/Networking/Introduction/de "Handbook:X86/Networking/Introduction/de")[Fortgeschrittene Konfiguration](/wiki/Handbook:X86/Networking/Advanced/de "Handbook:X86/Networking/Advanced/de")[Modulare Vernetzung](/wiki/Handbook:X86/Networking/Modular/de "Handbook:X86/Networking/Modular/de")[Drahtlose Netzwerke](/wiki/Handbook:X86/Networking/Wireless/de "Handbook:X86/Networking/Wireless/de")[Funktionalität hinzufügen](/wiki/Handbook:X86/Networking/Extending/de "Handbook:X86/Networking/Extending/de")[Dynamisches Management](/wiki/Handbook:X86/Networking/Dynamic/de "Handbook:X86/Networking/Dynamic/de")

## Contents

- [1Einführung in blockorientierte Geräte](#Einf.C3.BChrung_in_blockorientierte_Ger.C3.A4te)
  - [1.1Blockorientierte Geräte](#Blockorientierte_Ger.C3.A4te)
  - [1.2Partitionstabellen](#Partitionstabellen)
    - [1.2.1GUID Partitionstabelle (GPT)](#GUID_Partitionstabelle_.28GPT.29)
    - [1.2.2Master Boot Record (MBR) oder DOS Boot-Sector](#Master_Boot_Record_.28MBR.29_oder_DOS_Boot-Sector)
  - [1.3Fortgeschrittene Speicherlösungen](#Fortgeschrittene_Speicherl.C3.B6sungen)
  - [1.4Standard-Partitionsschema](#Standard-Partitionsschema)
- [2Ein Partitionsschema entwerfen](#Ein_Partitionsschema_entwerfen)
  - [2.1Wie viele Partitionen und wie groß?](#Wie_viele_Partitionen_und_wie_gro.C3.9F.3F)
  - [2.2Was ist mit dem Swap-Speicher?](#Was_ist_mit_dem_Swap-Speicher.3F)
  - [2.3Was ist die EFI System-Partition (ESP)?](#Was_ist_die_EFI_System-Partition_.28ESP.29.3F)
  - [2.4Was ist die BIOS Boot-Partition?](#Was_ist_die_BIOS_Boot-Partition.3F)
- [3Partitionieren des Laufwerks mit GPT für UEFI](#Partitionieren_des_Laufwerks_mit_GPT_f.C3.BCr_UEFI)
  - [3.1Anzeigen des Partitionslayouts](#Anzeigen_des_Partitionslayouts)
  - [3.2Erzeugen eines neuen Disklabels / Löschen aller Partitionen](#Erzeugen_eines_neuen_Disklabels_.2F_L.C3.B6schen_aller_Partitionen)
  - [3.3Erstellen der EFI System-Partition (ESP)](#Erstellen_der_EFI_System-Partition_.28ESP.29)
  - [3.4Erstellen der Swap-Partition](#Erstellen_der_Swap-Partition)
  - [3.5Erstellen der Root-Partition](#Erstellen_der_Root-Partition)
  - [3.6Speichern des Partitionslayouts](#Speichern_des_Partitionslayouts)
- [4Partitionieren des Laufwerks mit MBR für BIOS / Legacy-Boot](#Partitionieren_des_Laufwerks_mit_MBR_f.C3.BCr_BIOS_.2F_Legacy-Boot)
  - [4.1Anzeigen des aktuellen Partitionslayouts](#Anzeigen_des_aktuellen_Partitionslayouts)
  - [4.2Erzeugen eines neuen Disklabels / Löschen aller Partitionen](#Erzeugen_eines_neuen_Disklabels_.2F_L.C3.B6schen_aller_Partitionen_2)
  - [4.3Erstellen der Boot-Partition](#Erstellen_der_Boot-Partition)
  - [4.4Erstellen der Swap-Partition](#Erstellen_der_Swap-Partition_2)
  - [4.5Erstellen der Root-Partition](#Erstellen_der_Root-Partition_2)
  - [4.6Speichern des Partitionslayouts](#Speichern_des_Partitionslayouts_2)
- [5Erstellen von Dateisystemen](#Erstellen_von_Dateisystemen)
  - [5.1Einleitung](#Einleitung)
  - [5.2Dateisysteme](#Dateisysteme)
  - [5.3Dateisystem auf einer Partition anlegen](#Dateisystem_auf_einer_Partition_anlegen)
    - [5.3.1EFI system partition filesystem](#EFI_system_partition_filesystem)
    - [5.3.2Legacy BIOS boot partition filesystem](#Legacy_BIOS_boot_partition_filesystem)
    - [5.3.3Small ext4 partitions](#Small_ext4_partitions)
  - [5.4Aktivieren der Swap-Partition](#Aktivieren_der_Swap-Partition)
- [6Einhängen der Root-Partition](#Einh.C3.A4ngen_der_Root-Partition)

## Einführung in blockorientierte Geräte

### Blockorientierte Geräte

Schauen wir uns die Festplatten-spezifischen Aspekte von Gentoo Linux und Linux im Allgemeinen an - insbesondere blockorientierte Geräte (Block Devices), Partitionen und Linux Dateisysteme. Wenn Sie die Vor- und Nachteile von Festplatten verstanden haben, können Sie Partitionen und Dateisysteme für die Installation erstellen.

Zu Beginn schauen wir uns blockorientierte Geräte an. SCSI- und SATA-Laufwerke haben Device-Namen wie: /dev/sda, /dev/sdb, /dev/sdc usw. Modernere Rechner können PCI-Express basierte NVMe Solid-State-Disks haben, die Device-Namen haben wie: /dev/nvme0n1, /dev/nvme0n2 usw.

Die folgende Tabelle soll Lesern dabei helfen herauszufinden, wo bestimmte Arten von blockorientierten Geräten zu finden sind:

Device-TypStandard Device-NameAnmerkungen
IDE, SATA, SAS, SCSI, or USB flash/dev/sdaDiese Device-Typen werden auf Hardware ab 2007 verwendet - und sind vermutlich die am häufigsten genutzten Device-Namen unter Linux. Diese Geräte werden als blockorientierter Speicher angeschlossen über den [SATA bus](https://en.wikipedia.org/wiki/Serial_ATA "wikipedia:Serial ATA"), über [SCSI](https://en.wikipedia.org/wiki/SCSI "wikipedia:SCSI") und über [USB](https://en.wikipedia.org/wiki/USB "wikipedia:USB"). Beispielsweise wird die erste Partition des ersten SATA-Devices /dev/sda1 genannt.
NVM Express (NVMe)/dev/nvme0n1The latest in solid state technology, [NVMe](https://en.wikipedia.org/wiki/NVM_Express "wikipedia:NVM Express") drives are connected to the PCI Express bus and have the fastest transfer block speeds on the market. Systems from around 2014 and newer may have support for NVMe hardware. The first partition on the first NVMe device is called /dev/nvme0n1p1.
MMC, eMMC, and SD/dev/mmcblk0[embedded MMC](https://en.wikipedia.org/wiki/MultiMediaCard#eMMC "wikipedia:MultiMediaCard") devices, SD cards, and other types of memory cards can be useful for data storage. That said, many systems may not permit booting from these types of devices. It is suggested to **not** use these devices for active Linux installations; rather consider using them to transfer files, which is their design goal. Alternatively they could be useful for short-term backups.

Die oben genannten blockorientierten Geräte repräsentieren eine abstrakte Schnittstelle zur Festplatte. Benutzerprogramme können diese Block Devices nutzen, um mit der Festplatte zu interagieren, ohne sich darum sorgen zu müssen, ob die Festplatten über SATA, SCSI oder etwas anderem angebunden sind. Das Programm kann den Speicher auf der Festplatte einfach als eine Anhäufung zusammenhängender 4096-Byte (4k) Blöcke mit wahlfreiem Zugriff ansprechen.

### Partitionstabellen

Obwohl es theoretisch möglich wäre, eine vollständige und unpartitionierte Festplatte für die Unterbringung eines Linux Systems zu nutzen (beispielsweise um ein btrfs RAID zu erzeugen), wird das in der Praxis selten gemacht. Statt dessen teilt man das Festplatten Block Device in kleinere, besser verwaltbare Block Devices auf. Auf **x86** Systemen nennt man diese Partitionen. Derzeit gibt es zwei verschiedene Standard Partitionierungs-Technologien: MBR (manchmal aus "DOS Disklabel" genannt) und GPT. Diese hängen zusammen mit den beiden Boot-Typen: BIOS und UEFI.

#### GUID Partitionstabelle (GPT)

Das _GUID Partition Table (GPT)_ Setup (auch "GPT disklabel" genannt) verwendet 64-Bit Kennzeichner für die Partitionen. Der Ort, an dem die Partitions-Informationen gespeichert werden, ist außerdem viel größer, als die 512 Bytes der MBR Partitionstabelle ("DOS disklabel"). Dies bedeutet, dass es praktisch kein Limit für die Anzahl der Partitionen gibt. Darüber hinaus wird die Größe einer Partition durch ein viel größeres Limit begrenzt (fast 8 ZiB - ja, Zebibytes).

Wenn die Software-Schnittstelle zwischen dem Betriebssystem und der Firmware UEFI ist (statt BIOS), ist GPT fast schon zwingend erforderlich, weil mit MBR (DOS disklabel) Kompatibilitäts-Probleme entstehen werden.

GPT nutzt die Vorteile von Prüfsummen und Redundanz. Es verwendet CRC32 Prüfsummen, um Fehler in den Kopfdaten und in den Partitionseinträgen zu erkennen. Weiterhin gibt es eine Kopie der GPT am Ende der Festplatte. Diese Backup Partitionstabelle kann verwendet werden, um eine beschädigte Primär-GPT am Anfang der Festplatte wiederherzustellen.

**Wichtig**

Es gibt ein paar Warnhinweise zu GPT:

- GPT bei kann bei BIOS-basierten Computern verwendet werden. In diesem Fall kann aber keine Dual-Boot-Konfiguration mit Microsoft Windows Betriebssystemen eingerichtet werden. Der Grund ist, dass Microsoft Windows im UEFI-Modus bootet, wenn es eine GPT Partitionstabelle erkennt.
- Manche ältere Motherboards mit fehlerhafter Firmware können Probleme beim Booten haben, wenn sie für den BIOS/CSM/Legacy Boot-Modus konfiguriert sind, aber auf der Festplatte eine GPT Partitionstabelle finden.

#### Master Boot Record (MBR) oder DOS Boot-Sector

Der _[Master Boot Record](https://en.wikipedia.org/wiki/Master_boot_record "wikipedia:Master boot record")_ Boot-Sektor (auch "DOS Boot-Sektor" oder "DOS Disklabel" genannt) wurde im Jahr 1983 mit PC DOS 2.x eingeführt. MBR verwendet 32-Bit Kennzeichner für den Start der Sektoren und die Länge der Partitionen. Drei Partitions-Typen werden unterstützt: primär, erweitert und logisch. Primäre Partitionen speichern ihre Informationen im Master Boot Record selbst - ein sehr kleiner Bereich (meist 512 Bytes) ganz am Anfang der Festplatte. Aufgrund des geringen Platzes werden nur vier Primäre Partitionen unterstützt (beispielsweise /dev/sda1 bis /dev/sda4).

Um mehr als vier Partitionen zu unterstützen, kann eine der primären Partitionen in dem MBR als _erweitert_ markiert werden. Diese Partition kann dann zusätzliche logische Partitionen beinhalten (Partitionen in einer Partition).

**Wichtig**

Obwohl MBR Partitionstabellen von den meisten Mainboard Herstellern noch unterstützt werden, gelten sie und die damit verbundenen Einschränkungen als veraltet. Sofern Sie nicht mit Hardware arbeiten, die von 2010 oder früher ist, sollten Sie Ihre Festplatte mit einer [GUID Partititionstabelle](https://en.wikipedia.org/wiki/GUID_Partition_Table "wikipedia:GUID Partition Table") partitionieren. Leser, die MBR Partitionstabellen verwenden müssen, sollten Folgendes beachten:

- Die meisten Mainboards von 2010 oder später betrachten die Verwendung des MBR Boot-Sektors als veraltet (unterstützt, aber nicht optimal).
- Da MBR Partitionstabellen 32-Bit Zeiger verwenden, können mit ihnen keine Speicherbereiche adressiert werden, die größer als 2 TiB sind.
- Sofern keine "Erweiterte Partition" erstellt wird, unterstützen MBR Partitionstabellen maximal 4 Partitionen.
- Von MBR Partitionstabellen gibt es kein Backup auf der Festplatte. Wenn etwas die MBR Partitionstabelle (versehentlich) überschreibt, sind die Partitions-Informationen verloren.

Trotz all dieser Nachteile werden MBR und die BIOS Boot-Methode immer noch häufig in virtualisierten Cloud-Umgebungen wie AWS verwendet.

Die Autoren dieses Handbuchs empfehlen - wann immer möglich - [GPT](#GPT) für Gentoo Installationen zu verwenden.

### Fortgeschrittene Speicherlösungen

Die **x86** Installations-CDs bieten Unterstützung für den Logical Volume Manager (LVM). LVM erhöht die Flexibilität, die man mit "normalem" Partitionieren erreichen kann. So können Partitionen und Festplatten in "Volume Groups" zusammengefasst werden, "RAID Groups" können definiert werden, oder Caches auf schnellen SSDs können für langsame Festplatten eingerichtet werden. Während der folgenden Installationsanleitung konzentrieren wir uns auf "normale" Partitionen. Es ist dennoch gut zu wissen, dass auch LVM unterstützt wird. Weitere Informationen finden Sie in dem [LVM](/wiki/LVM "LVM") Artikel. Bitte beachten Sie: Gentoo Linux unterstützt LVM, aber der Umgang mit LVM würde den Rahmen dieses Handbuchs sprengen.

Although usage is not covered in the handbook, below is a list helpful guides to get the system running:

- [Btrfs](/wiki/Btrfs/Native_System_Root_Guide "Btrfs/Native System Root Guide")
- [ZFS](/wiki/ZFS/rootfs "ZFS/rootfs")
- [LVM (Logical Volume Manager)](/wiki/LVM/de "LVM/de")

Disk encryption is also handled in the same manner:

- [Rootfs encryption](/wiki/Rootfs_encryption "Rootfs encryption")

### Standard-Partitionsschema

In dem Rest dieses Handbuchs werden wir zwei Fälle besprechen und erklären: 1) GPT Partitionstabelle und UEFI-Boot und 2) MBR Partitionstabelle und BIOS-Boot (veraltet). Obwohl es möglich ist zu mischen, würde dies über den Rahmen dieses Handbuchs hinausgehen. Wie bereits geschrieben, sollten Installationen auf moderner Hardware eine GPT Partitionstabelle und UEFI-Boot verwenden. Als Ausnahme von dieser Regel gelten virtualisierte (Cloud-) Umgebungen, bei denen immer noch häufig eine MBR Partitionstabelle und BIOS-Boot verwendet werden.

1. UEFI firmware with GUID Partition Table (GPT) disk.
2. MBR DOS/legacy BIOS firmware with a MBR partition table disk.

While it is possible to mix and match boot types with certain motherboard firmware, mixing goes beyond the intention of the handbook. As previously stated, it is strongly recommended for installations on modern hardware to use UEFI boot with a GPT disklabel disk.

Wir werden das folgende Partitionsschema als einfaches Beispiel verwenden:

**Wichtig**

The first row of the following table contains exclusive information for _**either**_ a GPT disklabel _**or**_ a MBR DOS/legacy BIOS disklabel. When in doubt, proceed with GPT, since **x86** machines manufactured after the year 2010 generally support UEFI firmware and GPT boot sector.

Partition
Dateisystem
Größe
Beschreibung
/dev/sda1fat32 (UEFI) oder xfs (BIOS auch bekannt als Legacy Boot)
256M
Boot/EFI System Partition
/dev/sda2(swap)
RAM Größe \* 2
Swap Partition
/dev/sda3xfs
Rest der Festplatte
Root Partition

Fortgeschrittene, denen diese Informationen ausreichen, können die folgenden Abschnitte überspringen und direkt zur Partitionierung weitergehen.

Sowohl fdisk, als auch parted sind Partitionierungs-Tools. fdisk ist sehr bekannt, stabil und die erste Wahl für Arbeiten an MBR Partitionstabellen. parted war eines der ersten Block Device Management-Programme, das auch mit GPT Partitionstabellen umgehen konnte. Im Folgenden wird fdisk verwendet, weil es ein besseres textbasiertes Benutzer-Interface hat.

Bevor wir zu den Anweisungen zur Erstellung kommen, beschreiben die ersten Abschnitte im Detail, wie Partitionsschemas erstellt werden können und was die häufigsten Fallstricke sind.

## Ein Partitionsschema entwerfen

### Wie viele Partitionen und wie groß?

Bei dem Design des Partitionsschemas sollten die Anforderungen an das System und an die Dateisysteme berücksichtigt werden. Wenn es viele Nutzer gibt, ist eine eigene Partition /home/ ratsam, da diese die Sicherheit erhöht und Backups und andere Wartungsarbeiten vereinfacht. Wenn Gentoo installiert wird, um als Mailserver zu dienen, dann sollte es eine eigene Partition /var/ geben, weil alle Mails im Verzeichnis /var/ gespeichert werden. Spiele-Server werden eine eigene Partition /opt/ besitzen, da die meiste Spiele-Server-Software dort installiert wird. Der Grund für diese Empfehlungen ist ähnlich wie für das /home/ Verzeichnis: Sicherheit, Backups und Wartung.

Bei den meisten Gentoo-Installationen sollten /usr/ und /var/ relativ groß sein. In /usr werden die Mehrzahl der Anwendungen und auch der Linux Kernel Quellcode gespeichert (unter /usr/src). Standardmäßig enthält /var/ das Gentoo ebuild Repository (unter /var/db/repos/gentoo), das alleine schon rund 650 MiB Plattenplatz benötigt. Diese Größenabschätzung enthält noch nicht den benötigten Plattenplatz für die Verzeichnisse /var/cache/distfiles und /var/cache/binpkgs, die sich im Laufe der Zeit mit Source-Code Dateien und (optional) mit Binärpaketen füllen werden - je nachdem, wann und wie sie dem System hinzugefügt werden.

Die Anzahl und Größe der Partitionen hängt vom Abwägen der Vor- und Nachteile und der Auswahl der besten Lösung für einen gegebenen Anwendungsfall ab. Separate Partitionen oder Volumes haben folgende Vorteile:

- Sie können das performanteste Dateisystem für jede Partition oder jedes Volume wählen.
- Dem Gesamtsystem kann der freie Speicherplatz nicht ausgehen, wenn ein fehlerhaftes Tool kontinuierlich Dateien auf eine Partition oder ein Volume schreibt.
- Falls nötig, kann die Zeit für Dateisystemüberprüfungen reduziert werden, da mehrere Überprüfungen gleichzeitig durchgeführt werden können. (Dieser Vorteil kommt aber eher bei mehreren Festplatten, als bei mehreren Partitionen auf einer Festplatte zum Tragen.)
- Sie können die Sicherheit erhöhen, indem Sie einige Partitionen oder Volumes "read-only", `nosuid` (setuid Flags werden ignoriert), `noexec` (executable Flags werden ignoriert) etc. einbinden.

Viele separate Partitionen können aber auch Nachteile haben:

- Wenn diese schlecht an das System angepasst sind, kann es sein, dass eine Partition voll ist und auf einer anderen Partition noch viel freier Platz verfügbar ist.
- Eine separate Partition für /usr/ kann es erforderlich machen, dass beim Booten ein initramfs verwendet wird, welches diese Partitionen vor der Ausführung anderer Boot-Skripte mountet. Das Erzeugen und Betreiben eines initramsfs ist nicht Teil dieses Handbuchs. **Wir empfehlen Anfängern, für /usr/ keine eigene Partition zu verwenden.**
- Es gibt ein Limit von maximal 15 Partitionen für SCSI und SATA - es sei denn, der Datenträger nutzt GPT-Labels.

**Hinweis**

Installationen, die systemd als Dienst-und Init-System verwenden wollen, müssen /usr/ beim Booten verfügbar haben, entweder als Teil des Root-Dateisystems oder eingehängt über ein initramfs.

### Was ist mit dem Swap-Speicher?

Recommendations for swap space size
RAM sizeSuspend support?Hibernation support?
2 GB or less4GB4GB
2 to 8 GBRAM amount2 \* RAM
8 to 64 GB8 GB minimum, 16 maximum1.5 \* RAM
64 GB or greater8 GB minimumHibernation not recommended! Hibernation is _not_ recommended for systems with very large amounts of memory. While possible, the entire contents of memory must be written to disk in order to successfully hibernate. Writing tens of gigabytes (or worse!) out to disk can can take a considerable amount of time, especially when rotational disks are used. It is best to suspend in this scenario.

Es gibt keine perfekte Größe für den Swap-Speicher. Der Zweck von Swap-Speicher ist, Festplattenspeicherplatz für den Kernel bereitzuhalten, wenn der interne Speicher (RAM) knapp wird. Der Swap-Speicher erlaubt dem Kernel, Speicherseiten, auf die vermutlich nicht bald zugegriffen wird, auf die Platte auszulagern (Swap oder Page-Out). Dadurch kann Arbeitsspeicher im RAM für den aktuell laufenden Prozess freigemacht werden. Werden die auf die Festplatte ausgelagerten Speicherseiten (Pages) jedoch plötzlich benötigt, müssen diese Seiten wieder zurück in den Arbeitsspeicher geladen werden (Page-In). Dies dauert jedoch erheblich länger, als wenn die Daten direkt aus dem RAM gelesen werden könnten (da Festplatten verglichen mit Arbeitsspeicher sehr langsam sind).

Wenn auf einem System keine speicherintensiven Anwendungen ausgeführt werden oder das System viel RAM zur Verfügung hat, benötigt es vermutlich nicht viel Swap-Speicher. Wenn jedoch der Ruhezustand "Hibernation" verwendet werden soll, wird der Swap-Speicher verwendet, um den _gesamten Inhalt des Hauptspeichers (RAM)_ zu sichern (dieser Ruhezustand wird bei Desktop- und Laptop-Systemen häufiger verwendet, als bei Servern). Wenn das System den Ruhezustand "Hibernation" unterstützen soll, muss der Swap-Speicher so groß wie oder größer als der Hauptspeicher (RAM) sein.

Als generelle Regel gilt: der Swap-Speicher sollte zwei Mal so groß sein wie der Arbeitsspeicher (RAM). Auf Systemen mit mehreren (rotierenden) Festplatten ist es sinnvoll, eine Swap-Partition auf jeder Festplatte einzurichten, damit Schreib-/Lese-Operationen parallel ausgeführt werden können. Je schneller auf einen Festplatte zugegriffen werden kann, desto schneller wird das System arbeiten, wenn auf Swap-Speicher zugegriffen werden muss. Wenn zwischen rotierenden Festplatten und SSDs gewählt werden kann, ist es aus Performance-Sicht besser, den Swap-Speicher auf die SSD zu legen. Alternativ zu Swap-Partitionen können auch Swap-Dateien verwendet werden; dies ist hauptsächlich interessant bei Systemen mit sehr geringem Festplatten-Platz.

### Was ist die EFI System-Partition (ESP)?

Wenn Sie Gentoo auf einem System installieren, das UEFI zum Booten des Betriebssystems verwendet (statt des BIOS), ist es wichtig, dass eine EFI System-Partition (ESP) erzeugt wird. Die folgenden Anweisungen enthalten die erforderlichen Hinweise, um eine ESP zu erzeugen. **Eine EFI System-Partition ist nicht erforderlich, wenn im BIOS/Legacy-Modus gebootet werden soll.**

Die EFI System-Partition _muss_ eine FAT Variante sein (derartige Dateisysteme werden auf Linux Systemen oft als _vfat_ angezeigt). Die offizielle [UEFI Spezifikation](http://www.uefi.org/sites/default/files/resources/UEFI%202_5.pdf) deutet darauf hin, dass FAT12-, FAT16- oder FAT32-Dateisysteme von der UEFI Firmware akzeptiert werden, wobei für die ESP FAT32 empfohlen wird. Nach dem Partitionieren sollten Sie die ESP entsprechend formatieren:

`root #` `mkfs.fat -F 32 /dev/sda1`

**Wichtig**

Wenn die ESP nicht mit einer FAT-Variante formatiert wird, wird die UEFI Firmware Ihres Systems den Bootloader (oder den Linux Kernel) nicht finden und wahrscheinlich nicht in der Lage sein, das Betriebssystem zu booten!

### Was ist die BIOS Boot-Partition?

Eine BIOS Boot-Partition ist nur notwendig, wenn eine GPT Partitionstabelle mit GRUB2 im BIOS/Legacy Boot-Modus verwendet wird. **Sie ist nicht notwendig, wenn im EFI/UEFI-Modus gebootet werden soll. Sie ist ebenfalls nicht notwendig, wenn eine MBR Partitionstabelle verwendet wird.** Sie ist eine sehr kleine Partition (1 bis 2 MB), in der Bootloader wie GRUB2 zusätzliche Daten ablegen können, die nicht in den zugeordneten Speicher passen (einige hundert Bytes im Fall des MBR) und die nirgendwo anders gespeichert werden können. In diesem Handbuch werden wir keine BIOS Boot-Partition verwenden.

## Partitionieren des Laufwerks mit GPT für UEFI

Die folgenden Abschnitte erklären, wie das Beispiel Partitionslayout für eine GPT / UEFI-Boot Installation mit fdisk erstellt werden kann. Das Beispiel Partitionslayout wurde bereits früher erwähnt:

Passen Sie das Partitionslayout nach Ihren persönlichen Bedürfnissen an.

Partition
Beschreibung
/dev/sda1EFI System (und boot)
/dev/sda2Swap Partition
/dev/sda3Root Partition

### Anzeigen des Partitionslayouts

fdisk ist ein beliebtes und leistungsstarkes Tool zum Aufteilen eine Festplatte in Partitionen. Starten Sie fdisk mit der Festplatte (in unserem Beispiel verwenden wir /dev/sda):

`root #` `fdisk /dev/sda`

Verwenden Sie die `p`-Taste, um die aktuelle Konfiguration der Partitionen anzuzeigen:

`Command (m for help):` `p`

```
Disk /dev/sda: 28.89 GiB, 31001149440 bytes, 60549120 sectors
Disk model: DataTraveler 2.0
Units: sectors of 1 * 512 = 512 bytes
Sector size (logical/physical): 512 bytes / 512 bytes
I/O size (minimum/optimal): 512 bytes / 512 bytes
Disklabel type: gpt
Disk identifier: 21AAD8CF-DB67-0F43-9374-416C7A4E31EA

Device        Start      End  Sectors  Size Type
/dev/sda1      2048   526335   524288    1G EFI System
/dev/sda2    526336  2623487  2097152    1G Linux swap
/dev/sda3   2623488 19400703 16777216    8G Linux filesystem
/dev/sda4  19400704 60549086 41148383 19.6G Linux filesystem

```

Device Start End Sectors Size Type
/dev/sda1 2048 2099199 2097152 1G EFI System
/dev/sda2 2099200 10487807 8388608 4G Linux swap
/dev/sda3 10487808 1953523711 1943035904 926.5G Linux root (x86-64)

}}

Diese Festplatte beherbergt bisher zwei Linux-Dateisysteme (jedes mit einer dazugehörigen Partition gelistet als "Linux") und auch eine Swap-Partition (gelistet als "Linux swap").

### Erzeugen eines neuen Disklabels / Löschen aller Partitionen

Drücken Sie `g`, um eine neue GPT Partitionstabelle auf der Festplatte zu erstellen. Dies wird alle existierenden Partitionen löschen.

`Command (m for help):` `g`

```
Created a new GPT disklabel (GUID: 87EA4497-2722-DF43-A954-368E46AE5C5F).

```

Wenn es bereits eine GPT Partitionstabelle gibt (siehe Ausgabe der Taste `p` weiter oben), können Sie alternativ die einzelnen Partitionen der Reihe nach löschen.
Drücken Sie `d` um eine Partition zu löschen. Zum Löschen einer vorhandenen Partition /dev/sda1:

`Command (m for help):` `d`

```
Partition number (1-4): 1

```

Die Partition ist nun zum Löschen vorgemerkt. Sie taucht nicht länger auf, wenn die Liste der der Partitionen ausgegeben wird ( `p`). Sie wird jedoch nicht gelöscht, solange die Änderungen nicht gespeichert werden. Dies erlaubt dem Benutzer, die Operation abzubrechen, falls ein Fehler passiert ist - in diesem Fall drücken Sie umgehend `q` gefolgt von `Enter`. Die bisherige Partition wird dann nicht gelöscht.

Drücken Sie wiederholt `p`, um die Partitionsliste anzuzeigen, gefolgt von `d` und der Nummer der zu löschenden Partition. Schließlich wird die Partitionstabelle leer sein:

`Command (m for help):` `p`

```
Disk /dev/sda: 28.89 GiB, 31001149440 bytes, 60549120 sectors
Disk model: DataTraveler 2.0
Units: sectors of 1 * 512 = 512 bytes
Sector size (logical/physical): 512 bytes / 512 bytes
I/O size (minimum/optimal): 512 bytes / 512 bytes
Disklabel type: gpt
Disk identifier: 87EA4497-2722-DF43-A954-368E46AE5C5F

```

Jetzt, da die Partitionstabelle im Speicher leer ist, sind wir bereit die Partitionen zu erstellen.

### Erstellen der EFI System-Partition (ESP)

**Hinweis**

Eine kleinere ESP ist möglich, aber nicht empfehlenswert, zumal sie mit anderen Betriebssystemen gemeinsam genutzt werden kann.

Erstellen Sie zunächst eine kleine EFI System-Partition, die später unter /boot eingehängt werden wird. Drücken Sie `n` für neue Partition, gefolgt von `1`, um die erste Partition zu wählen. Wenn Sie aufgefordert werden den ersten Sektor einzugeben, stellen Sie sicher, dass die Partition bei 2048 beginnt (dies wird möglicherweise für den Bootloader benötigt) und drücken Sie `Enter`. Wenn Sie nach dem letzten Sektor gefragt werden geben Sie +1G ein, um eine 1 GByte große Partition zu erstellen:

`Command (m for help):` `n`

```
Partition number (1-128, default 1): 1
First sector (2048-60549086, default 2048):
Last sector, +/-sectors or +/-size{K,M,G,T,P} (2048-60549086, default 60549086): +1G

Created a new partition 1 of type 'Linux filesystem' and of size 1 GiB.

```

Do you want to remove the signature? \[Y\]es/\[N\]o: Y
The signature will be removed by a write command.

}}

Markieren Sie die Partition als EFI System-Partition:

`Command (m for help):` `t`

```
Selected partition 1
Partition type (type L to list all types): 1
Changed type of partition 'Linux filesystem' to 'EFI System'.

```

### Erstellen der Swap-Partition

Erstellen Sie als nächstes die Swap-Partition. Drücken Sie `n` für neue Partition, dann `2`, um die zweite Partition (/dev/sda2) zu erstellen. Wenn Sie aufgefordert werden den ersten Sektor einzugeben, bestätigen Sie die Voreinstellung durch `Enter`. Wenn Sie nach dem letzten Sektor gefragt werden geben Sie +4G ein (oder jede andere Größe, die Sie als Swap-Speicherplatz benötigen), um eine 4 GB große Partition zu erstellen.

`Command (m for help):` `n`

```
Partition number (2-128, default 2):
First sector (526336-60549086, default 526336):
Last sector, +/-sectors or +/-size{K,M,G,T,P} (526336-60549086, default 60549086): +4G

Created a new partition 2 of type 'Linux filesystem' and of size 4 GiB.

```

Nachdem dies erledigt ist, drücken Sie `t` um den Partitionstyp einzustellen, `2` um die gerade erzeugte Partition auszuwählen und geben Sie _19_ ein, um den Partitionstyp auf "Linux Swap" zu setzen.

`Command (m for help):` `t`

```
Partition number (1,2, default 2): 2
Partition type (type L to list all types): 19

Changed type of partition 'Linux filesystem' to 'Linux swap'.

```

### Erstellen der Root-Partition

Um schließlich die Root-Partition zu erstellen, drücken Sie ein weiteres Mal `n`, um eine neue Partition zu erstellen. Drücken Sie `3` um die dritte Partition /dev/sda3 zu erstellen. Wenn Sie aufgefordert werden den ersten Sektor einzugeben, bestätigen Sie die Voreinstellung durch `Enter`. Wenn Sie nach dem letzten Sektor gefragt werden bestätigen Sie nochmals die Voreinstellung durch `Enter`, um den bisher noch frei verbliebenen restlichen Festplattenanteil dafür zu verwenden. Nachdem Sie diese Schritte abgeschlossen haben, sollte die Eingabe von `p` eine Partitionstabelle ausgeben, die der folgenden ähnlich sehen sollte:

`Command (m for help):` `n`

```
Partition number (3-128, default 3): 3
First sector (10487808-1953525134, default 10487808):
Last sector, +/-sectors or +/-size{K,M,G,T,P} (10487808-1953525134, default 1953523711):
</div>

<div lang="en" dir="ltr" class="mw-content-ltr">
Created a new partition 3 of type 'Linux filesystem' and of size 926.5 GiB..

```

**Hinweis**

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
Disk /dev/sda: 28.89 GiB, 31001149440 bytes, 60549120 sectors
Disk model: DataTraveler 2.0
Units: sectors of 1 * 512 = 512 bytes
Sector size (logical/physical): 512 bytes / 512 bytes
I/O size (minimum/optimal): 512 bytes / 512 bytes
Disklabel type: gpt
Disk identifier: 87EA4497-2722-DF43-A954-368E46AE5C5F

Device       Start      End  Sectors  Size Type
/dev/sda1     2048   526335   524288    1G EFI System
/dev/sda2   526336  8914943  8388608    4G Linux swap
/dev/sda3  8914944 60549086 51634143 24.6G Linux filesystem

```

Device Start End Sectors Size Type
/dev/sda1 2048 2099199 2097152 1G EFI System
/dev/sda2 2099200 10487807 8388608 4G Linux swap
/dev/sda3 10487808 1953523711 1943035904 926.5G Linux root (x86-64)

Filesystem/RAID signature on partition 1 will be wiped.

}}

### Speichern des Partitionslayouts

Um die Partitionstabelle zu speichern und fdisk zu beenden, drücken Sie `w`.

`Command (m for help):` `w`

Nachdem die Partitionen erstellt wurden, ist es jetzt Zeit, Dateisysteme darauf anzulegen.

## Partitionieren des Laufwerks mit MBR für BIOS / Legacy-Boot

Die folgenden Abschnitte erklären, wie das Beispiel Partitionslayout für eine MBR / BIOS-Boot Installation erstellt werden kann. Das Beispiel Partitionslayout wurde bereits früher erwähnt:

Partition
Description
/dev/sda1Boot partition
/dev/sda2Swap partition
/dev/sda3Root partition

Passen Sie das Partitionslayout nach Ihren persönlichen Bedürfnissen an.

### Anzeigen des aktuellen Partitionslayouts

Starten Sie fdisk mit der Festplatte (in unserem Beispiel verwenden wir /dev/sda):

`root #` `fdisk /dev/sda`

Verwenden Sie die `p`-Taste, um die aktuelle Konfiguration der Partitionen anzuzeigen:

`Command (m for help):` `p`

```
Disk /dev/sda: 28.89 GiB, 31001149440 bytes, 60549120 sectors
Disk model: DataTraveler 2.0
Units: sectors of 1 * 512 = 512 bytes
Sector size (logical/physical): 512 bytes / 512 bytes
I/O size (minimum/optimal): 512 bytes / 512 bytes
Disklabel type: gpt
Disk identifier: 21AAD8CF-DB67-0F43-9374-416C7A4E31EA

Device        Start      End  Sectors  Size Type
/dev/sda1      2048   526335   524288    1G EFI System
/dev/sda2    526336  2623487  2097152    1G Linux swap
/dev/sda3   2623488 19400703 16777216    8G Linux filesystem
/dev/sda4  19400704 60549086 41148383 19.6G Linux filesystem

```

Device Boot Start End Sectors Size Id Type
/dev/sda1 \* 2048 2099199 2097152 1G 83 Linux
/dev/sda2 2099200 10487807 8388608 4G 82 Linux swap / Solaris
/dev/sda3 10487808 1953525167 1943037360 926.5G 83 Linux

}}

Diese Festplatte beherbergt bisher zwei Linux-Dateisysteme (jedes mit einer dazugehörigen Partition gelistet als "Linux") und auch eine Swap-Partition (gelistet als "Linux swap").

### Erzeugen eines neuen Disklabels / Löschen aller Partitionen

Drücken Sie `o`, um eine neue MBR Partitionstabelle (im Folgenden auch "DOS disklabel" genannt) auf der Festplatte zu erstellen. Dies wird alle existierenden Partitionen löschen.

`Command (m for help):` `o`

```
Created a new DOS disklabel with disk identifier 0xe04e67c4.
The device contains 'gpt' signature and it will be removed by a write command. See fdisk(8) man page and --wipe option for more details.

```

Wenn bereits eine MBR Partitionstabelle existiert (siehe Ausgabe der Taste `p` weiter oben), können Sie alternativ die einzelnen Partitionen der Reihe nach löschen.
Drücken Sie `d` um eine Partition zu löschen. Zum Löschen einer vorhandenen Partition /dev/sda1:

`Command (m for help):` `d`

```
Partition number (1-4): 1

```

Die Partition ist nun zum Löschen vorgemerkt. Sie taucht nicht länger auf, wenn die Liste der der Partitionen ausgegeben wird ( `p`). Sie wird jedoch nicht gelöscht, solange die Änderungen nicht gespeichert werden. Dies erlaubt dem Benutzer, die Operation abzubrechen, falls ein Fehler passiert ist - in diesem Fall drücken Sie umgehend `q` gefolgt von `Enter`. Die bisherige Partition wird dann nicht gelöscht.

Drücken Sie wiederholt `p`, um die Partitionsliste anzuzeigen, gefolgt von `d` und der Nummer der zu löschenden Partition. Schließlich wird die Partitionstabelle leer sein:

`Command (m for help):` `p`

```
Disk /dev/sda: 28.89 GiB, 31001149440 bytes, 60549120 sectors
Disk model: DataTraveler 2.0
Units: sectors of 1 * 512 = 512 bytes
Sector size (logical/physical): 512 bytes / 512 bytes
I/O size (minimum/optimal): 512 bytes / 512 bytes
Disklabel type: dos
Disk identifier: 0xe04e67c4

```

Jetzt sind wir bereit die Partitionen zu erstellen.

### Erstellen der Boot-Partition

Erstellen Sie zunächst eine kleine Partition, die später unter /boot eingehängt werden wird. Drücken Sie `n` für neue Partition, gefolgt von `p` für eine primäre Partition und `1`, um die erste primäre Partition zu wählen. Wenn Sie aufgefordert werden den ersten Sektor einzugeben, stellen Sie sicher, dass die Partition bei 2048 beginnt (dies wird möglicherweise für den Bootloader benötigt) und drücken Sie `Enter`. Wenn Sie nach dem letzten Sektor gefragt werden geben Sie +1G ein, um eine 1 GByte große Partition zu erstellen:

`Command (m for help):` `n`

```
Partition type
   p   primary (0 primary, 0 extended, 4 free)
   e   extended (container for logical partitions)
Select (default p): p
Partition number (1-4, default 1): 1
First sector (2048-60549119, default 2048):
Last sector, +/-sectors or +/-size{K,M,G,T,P} (2048-60549119, default 60549119): +1G

Created a new partition 1 of type 'Linux' and of size 1 GiB.

```

Created a new partition 1 of type 'Linux' and of size 1 GiB.

}}

Mark the partition as bootable by pressing the `a` key and pressing `Enter`:

`Command (m for help):` `a`

```
Selected partition 1
The bootable flag on partition 1 is enabled now.

```

Note: if more than one partition is available on the disk, then the partition to be flagged as bootable will have to be selected.

### Erstellen der Swap-Partition

Erstellen Sie als nächstes die Swap-Partition. Drücken Sie `n` für neue Partition, gefolgt von `p` für eine primäre Partition und dann `2`, um die zweite primäre Partition (/dev/sda2) zu erstellen. Wenn Sie aufgefordert werden den ersten Sektor einzugeben, bestätigen Sie die Voreinstellung durch `Enter`. Wenn Sie nach dem letzten Sektor gefragt werden geben Sie +4G ein (oder jede andere Größe, die Sie als Swap-Speicherplatz benötigen), um eine 4 GB große Partition zu erstellen.

`Command (m for help):` `n`

```
Partition type
   p   primary (1 primary, 0 extended, 3 free)
   e   extended (container for logical partitions)
Select (default p): p
Partition number (2-4, default 2): 2
First sector (526336-60549119, default 526336):
Last sector, +/-sectors or +/-size{K,M,G,T,P} (526336-60549119, default 60549119): +4G

Created a new partition 2 of type 'Linux' and of size 4 GiB.

```

Nachdem dies erledigt ist, drücken Sie `t` um den Partitionstyp einzustellen, `2` um die gerade erzeugte Partition auszuwählen und geben Sie _82_ ein, um den Partitionstyp auf "Linux Swap" zu setzen.

`Command (m for help):` `t`

```
Partition number (1,2, default 2): 2
Hex code (type L to list all codes): 82

<div lang="en" dir="ltr" class="mw-content-ltr">
Changed type of partition 'Linux' to 'Linux swap / Solaris'.

```

### Erstellen der Root-Partition

Um schließlich die Root-Partition zu erstellen, drücken Sie ein weiteres Mal `n`, um eine neue Partition zu erstellen. Drücken Sie `p` und `3`, um die dritte primäre Partition /dev/sda3 zu erstellen. Wenn Sie aufgefordert werden den ersten Sektor einzugeben, bestätigen Sie die Voreinstellung durch `Enter`. Wenn Sie nach dem letzten Sektor gefragt werden bestätigen Sie nochmals die Voreinstellung durch `Enter`, um den bisher noch frei verbliebenen restlichen Festplattenanteil dafür zu verwenden. Nachdem Sie diese Schritte abgeschlossen haben, sollte die Eingabe von `p` eine Partitionstabelle ausgeben, die der folgenden ähnlich sehen sollte:

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
Disk /dev/sda: 28.89 GiB, 31001149440 bytes, 60549120 sectors
Disk model: DataTraveler 2.0
Units: sectors of 1 * 512 = 512 bytes
Sector size (logical/physical): 512 bytes / 512 bytes
I/O size (minimum/optimal): 512 bytes / 512 bytes
Disklabel type: dos
Disk identifier: 0xe04e67c4

Device     Boot   Start      End  Sectors  Size Id Type
/dev/sda1          2048   526335   524288    1G 83 Linux
/dev/sda2        526336  8914943  8388608    4G 82 Linux swap / Solaris
/dev/sda3       8914944 60549119 51634176 24.6G 83 Linux

```

Device Boot Start End Sectors Size Id Type
/dev/sda1 \* 2048 2099199 2097152 1G 83 Linux
/dev/sda2 2099200 10487807 8388608 4G 82 Linux swap / Solaris
/dev/sda3 10487808 1953525167 1943037360 926.5G 83 Linux

}}

### Speichern des Partitionslayouts

Um die Partitionstabelle zu speichern und fdisk zu beenden, drücken Sie `w`.

`Command (m for help):` `w`

Nachdem die Partitionen erstellt wurden, ist es jetzt Zeit, Dateisysteme darauf anzulegen.

## Erstellen von Dateisystemen

**Warnung**

Wenn Sie ein SSD- oder NVMe-Laufwerk verwenden, prüfen Sie bitte, ob es ein Firmware-Upgrade benötigt. Insbesondere einige Intel-SSDs (600p und 6000p) benötigen ein Firmware-Upgrade für [kritische Fehlerbehebungen](https://bugzilla.redhat.com/show_bug.cgi?id=1402533), um Datenbeschädigungen zu vermeiden, die durch XFS-I/O-Nutzungsmuster verursacht werden (allerdings nicht durch einen Fehler des Dateisystems). smartctl kann helfen, das Modell und die Firmware-Version zu überprüfen.

### Einleitung

Nachdem die Partitionen angelegt wurden, ist es an der Zeit, Dateisysteme darauf anzulegen. Im nächsten Abschnitt werden die unterschiedlichen Dateisysteme beschrieben, die Linux unterstützt. Leser, die bereits wissen, welches Dateisystem sie verwenden wollen, können bei [Dateisystem auf einer Partition anlegen](/wiki/Handbook:X86/Installation/Networking/de#Applying_a_filesystem_to_a_partition "Handbook:X86/Installation/Networking/de") fortfahren. Alle anderen sollten weiterlesen, um mehr über die verfügbaren Dateisysteme zu erfahren ...

### Dateisysteme

Linux unterstützt mehrere Dutzend Dateisysteme, wobei allerdings viele davon für ganz spezielle Anwendungszwecke optimiert sind. Nur einige Dateisysteme gelten als stabil auf der x86 Architektur. Es ist ratsam, sich über Dateisysteme und deren Unterstützungsgrad zu informieren, damit Sie nicht für wichtige Partitionen ein eher experimentelles Dateisystem wählen. **XFS ist das empfohlene all-round Dateisystem für alle Plattformen.** Nachfolgend eine nicht-vollständige Auswahl von verfügbaren Dateisystemen.

[btrfs](/wiki/Btrfs/de "Btrfs/de")Dateisystem der neueren Generation.

Bietet erweiterte Funktionen wie Snapshotting, Selbstheilung durch Prüfsummen, transparente Kompression, Subvolumes und integriertes RAID. Kernel vor 5.4.y sind nicht garantiert sicher für die Verwendung mit btrfs in der Produktion, da Korrekturen für ernsthafte Probleme nur in den neueren Versionen der LTS-Kernelzweige vorhanden sind. RAID 5/6 und Quota-Gruppen sind bei allen Versionen von btrfs unsicher.

[ext4](/wiki/Ext4/de "Ext4/de")Ext4 ist ein zuverlässiges, universell einsetztbares Dateisystem für alle Plattformen, auch wenn ihm moderne Funktionen wie Reflinks fehlen.[f2fs](/wiki/F2fs "F2fs")Das Flash-Friendly File System wurde ursprünglich von Samsung für die Verwendung mit NAND-Flash-Speicher entwickelt. Es ist eine gute Wahl für die Installation von Gentoo auf microSD-Karten, USB-Laufwerken oder anderen Flash-basierten Speichergeräten.[XFS](/wiki/XFS/de "XFS/de")Dateisystem mit Metadaten-Journaling, das über einen robusten Funktionsumfang verfügt und für Skalierbarkeit optimiert ist. Es wurde kontinuierlich weiterentwickelt, um moderne Funktionen einzubeziehen. Der einzige Nachteil ist, dass XFS-Partitionen noch nicht verkleinert werden können, obwohl daran gearbeitet wird. XFS unterstützt vor allem Reflinks und Copy on Write (CoW), was besonders auf Gentoo-Systemenen hilfreich ist, da die Benutzer viele Kompilierungen durchführen müssen. XFS ist das empfohlene modernen Allzweck-Dateisystem für alle Plattformen. Erfordert, dass eine Partition mindestens 300 MB groß ist.[VFAT](/wiki/VFAT "VFAT")Auch bekannt als FAT32, wird von Linux unterstützt, unterstützt aber nicht die Standard-UNIX-Berechtigungseinstellungen. Es wird hauptsächlich für die Interoperabilität/den Austausch mit anderen Betriebssystemen (Microsoft Windows oder Apples MacOS) verwendet, ist aber auch eine Notwendigkeit für einige System-Bootloader-Firmware (wie UEFI). Benutzer von UEFI-Systemen benötigen eine [EFI System Partition](/wiki/EFI_System_Partition/de "EFI System Partition/de"), die mit VFAT formatiert ist, um booten zu können.[NTFS](/wiki/NTFS/de "NTFS/de")Dieses 'New Technology"-Dateisystem ist das Vorzeige-Dateisystem von Microsoft Windows seit Windows NT 3.1. Ähnlich wie VFAT speichert es keine UNIX-Berechtigungseinstellungen oder erweiterte Attribute, die für BSD oder Linux notwendig sind, um ordnungsgemäß zu funktionieren, daher sollte es in den meisten Fällen nicht als Root-Dateisystem verwendet werden. Es sollte _nur_ für die Interoperabilität oder den Datenaustausch mit Microsoft Windows-Systemen verwendet werden (beachten Sie die Betonung auf _nur_).

Ausführlichere Informationen über Dateisysteme finden Sie in dem von der Community gepflegten [Dateisystem](/wiki/Filesystem/de "Filesystem/de")-Artikel.

### Dateisystem auf einer Partition anlegen

**Hinweis**

Bitte stellen Sie sicher, dass Sie das entsprechende Paket für das gewählte Dateisystem später im Handbuch emergen, bevor Sie am Ende des Installationsprozesses neu booten.

Dateisysteme können mit Hilfe von Programmen auf einer Partition oder auf einem Datenträger angelegt werden. Die folgende Tabelle zeigt, welchen Befehl Sie für welches Dateisystem benötigen. Um weitere Informationen zu einem Dateisystem zu erhalten, können Sie auf den Namen des Dateisystems klicken.

Dateisystem
Befehl zum Anlegen
Teil der Minimal CD?
Gentoo Paket
[btrfs](/wiki/Btrfs/de "Btrfs/de")mkfs.btrfs Yes
[sys-fs/btrfs-progs](https://packages.gentoo.org/packages/sys-fs/btrfs-progs)[ext4](/wiki/Ext4/de "Ext4/de")mkfs.ext4 Yes
[sys-fs/e2fsprogs](https://packages.gentoo.org/packages/sys-fs/e2fsprogs)[f2fs](/wiki/F2FS "F2FS")mkfs.f2fs Yes
[sys-fs/f2fs-tools](https://packages.gentoo.org/packages/sys-fs/f2fs-tools)[xfs](/wiki/XFS/de "XFS/de")mkfs.xfs Yes
[sys-fs/xfsprogs](https://packages.gentoo.org/packages/sys-fs/xfsprogs)[vfat](/wiki/FAT "FAT")mkfs.vfat Yes
[sys-fs/dosfstools](https://packages.gentoo.org/packages/sys-fs/dosfstools)[NTFS](/wiki/NTFS/de "NTFS/de")mkfs.ntfs Yes
[sys-fs/ntfs3g](https://packages.gentoo.org/packages/sys-fs/ntfs3g)

\|}

**Wichtig**

The handbook recommends new partitions as part of the installation process, but it is important to note running any mkfs command will erase any data contained within the partition. When necessary, ensure any data that exists within is appropriately backed up before creating a new filesystem.

Um beispielsweise die EFI System-Partition (/dev/sda1) als FAT32 und die root-Partition (/dev/sda3) als xfs zu formatieren (wie in dem Beispiel-Partitionsschema), würde man folgende Befehle verwenden:

`root #` `mkfs.xfs /dev/sda3`

{{#ifeq: 0\|1\|

#### EFI system partition filesystem

The EFI system partition (/dev/sda1) must be formatted as FAT32:

`root #` `mkfs.vfat -F 32 /dev/sda1`

#### Legacy BIOS boot partition filesystem

Systems booting via legacy BIOS with a MBR/DOS disklabel can use any filesystem format supported by the bootloader.

For example, to format with XFS:

`root #` `mkfs.xfs -c options=/usr/share/xfsprogs/mkfs/lts_6.18.conf /dev/sda1`

#### Small ext4 partitions

Bei der Verwendung von ext4 auf kleinen Partitionen (kleiner als 8 GiB), sollte das Dateisystem mit den passenden Optionen erstellt werden, um genügend Inodes zu reservieren. Dies kann mit einer der folgenden Anweisungen erfolgen:

`root #` `mkfs.ext4 -T small /dev/<device>`

Dies vervierfacht die Zahl der Inodes für ein angegebenes Dateisystem in der Regel, da es dessen "bytes-per-inode" (Bytes pro Inode) von 16 kB auf 4 kB pro Inode reduziert.

### Aktivieren der Swap-Partition

mkswap ist der Befehl der verwendet wird um Swap-Partitionen zu initialisieren:

`root #` `mkswap /dev/sda2`

**Hinweis**

Installations which were previously started, but did not finish the installation process can resume the installation from this point in the handbook. Use this link as the permalink: [Resumed installations start here](/wiki/Handbook:X86/Installation/Disks/de#Resumed_installations_start_here "Handbook:X86/Installation/Disks/de").

Zur Aktivierung der Swap-Partition verwenden Sie swapon:

`root #` `swapon /dev/sda2`

This 'activation' step is only necessary because the swap partition is newly created within the live environment. Once the system has been rebooted, as long as the swap partition is properly defined within fstab or other mount mechanism, swap space will activate automatically.

## Einhängen der Root-Partition

**Tipp**

Anwender, die ein Nicht-Gentoo Installationsmedium verwenden, müssen mit folgendem Befehl einen Mount-Point erzeugen:

`root #` `mkdir --parents /mnt/gentoo`

`root #` `mkdir --parents /mnt/gentoo`

Continue creating additional mount points necessary for any additional (custom) partition(s) created during previous steps by using the mkdir command.

Nachdem die Partitionen initialisiert wurden und ein Dateisystem beinhalten, ist es an der Zeit, diese einzuhängen. Verwenden Sie den Befehl mount, aber vergessen Sie nicht die notwendigen Einhänge-Verzeichnisse für jede Partition zu erzeugen. Als Beispiel hängen wir die Root-Partition ein:

Mount the root partition:

`root #` `mount /dev/sda3 /mnt/gentoo`

Continue mounting additional (custom) partitions as necessary using the mount command.

**Hinweis**

Wenn sich /tmp/ auf einer separaten Partition befinden muss, ändern Sie die Berechtigungen nach dem Einhängen:

`root #` `chmod 1777 /mnt/gentoo/tmp`

Dies gilt ebenfalls für /var/tmp.

In der Anleitung wird später das Dateisystem proc (eine virtuelle Schnittstelle zum Kernel) zusammen mit anderen Kernel Pseudo-Dateisystemen eingehängt. Zunächst installieren wir jedoch die [Gentoo Installationsdateien](/wiki/Handbook:X86/Installation/Stage/de "Handbook:X86/Installation/Stage/de").

[← Konfiguration des Netzwerks](/wiki/Handbook:X86/Installation/Networking/de "Previous part") [Anfang](/wiki/Handbook:X86/de "Handbook:X86/de") [Installation des Stage Archivs →](/wiki/Handbook:X86/Installation/Stage/de "Next part")