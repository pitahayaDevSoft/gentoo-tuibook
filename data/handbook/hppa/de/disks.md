# Disks

Other languages:

- Deutsch
- [English](/wiki/Handbook:HPPA/Installation/Disks "Handbook:HPPA/Installation/Disks (100% translated)")
- [español](/wiki/Handbook:HPPA/Installation/Disks/es "Manual de Gentoo: HPPA/Instalación/Discos (100% translated)")
- [français](/wiki/Handbook:HPPA/Installation/Disks/fr "Handbook:HPPA/Installation/Disks/fr (100% translated)")
- [italiano](/wiki/Handbook:HPPA/Installation/Disks/it "Handbook:HPPA/Installation/Disks/it (100% translated)")
- [magyar](/wiki/Handbook:HPPA/Installation/Disks/hu "Handbook:HPPA/Installation/Disks/hu (100% translated)")
- [polski](/wiki/Handbook:HPPA/Installation/Disks/pl "Handbook:HPPA/Installation/Disks (100% translated)")
- [português do Brasil](/wiki/Handbook:HPPA/Installation/Disks/pt-br "Manual:HPPA/Instalação/Discos (100% translated)")
- [русский](/wiki/Handbook:HPPA/Installation/Disks/ru "Handbook:HPPA/Installation/Disks (100% translated)")
- [தமிழ்](/wiki/Handbook:HPPA/Installation/Disks/ta "கையேடு:HPPA/நிறுவல்/வட்டுக்கள் (100% translated)")
- [中文（中国大陆）‎](/wiki/Handbook:HPPA/Installation/Disks/zh-cn "手册：HPPA/安装/准备磁盘 (100% translated)")
- [日本語](/wiki/Handbook:HPPA/Installation/Disks/ja "ハンドブック:HPPA/インストール/ディスク (100% translated)")
- [한국어](/wiki/Handbook:HPPA/Installation/Disks/ko "Handbook:HPPA/Installation/Disks/ko (100% translated)")

[HPPA Handbuch](/wiki/Handbook:HPPA/de "Handbook:HPPA/de")[Installation](/wiki/Handbook:HPPA/Full/Installation/de "Handbook:HPPA/Full/Installation/de")[Über die Installation](/wiki/Handbook:HPPA/Installation/About/de "Handbook:HPPA/Installation/About/de")[Auswahl des Mediums](/wiki/Handbook:HPPA/Installation/Media/de "Handbook:HPPA/Installation/Media/de")[Konfiguration des Netzwerks](/wiki/Handbook:HPPA/Installation/Networking/de "Handbook:HPPA/Installation/Networking/de")Vorbereiten der Festplatte(n)[Installation des Stage Archivs](/wiki/Handbook:HPPA/Installation/Stage/de "Handbook:HPPA/Installation/Stage/de")[Installation des Basissystems](/wiki/Handbook:HPPA/Installation/Base/de "Handbook:HPPA/Installation/Base/de")[Konfiguration des Kernels](/wiki/Handbook:HPPA/Installation/Kernel/de "Handbook:HPPA/Installation/Kernel/de")[Konfiguration des Systems](/wiki/Handbook:HPPA/Installation/System/de "Handbook:HPPA/Installation/System/de")[Installation der Tools](/wiki/Handbook:HPPA/Installation/Tools/de "Handbook:HPPA/Installation/Tools/de")[Konfiguration des Bootloaders](/wiki/Handbook:HPPA/Installation/Bootloader/de "Handbook:HPPA/Installation/Bootloader/de")[Abschluss](/wiki/Handbook:HPPA/Installation/Finalizing/de "Handbook:HPPA/Installation/Finalizing/de")[Arbeiten mit Gentoo](/wiki/Handbook:HPPA/Full/Working/de "Handbook:HPPA/Full/Working/de")[Portage-Einführung](/wiki/Handbook:HPPA/Working/Portage/de "Handbook:HPPA/Working/Portage/de")[USE-Flags](/wiki/Handbook:HPPA/Working/USE/de "Handbook:HPPA/Working/USE/de")[Portage-Features](/wiki/Handbook:HPPA/Working/Features/de "Handbook:HPPA/Working/Features/de")[Initskript-System](/wiki/Handbook:HPPA/Working/Initscripts/de "Handbook:HPPA/Working/Initscripts/de")[Umgebungsvariablen](/wiki/Handbook:HPPA/Working/EnvVar/de "Handbook:HPPA/Working/EnvVar/de")[Arbeiten mit Portage](/wiki/Handbook:HPPA/Full/Portage/de "Handbook:HPPA/Full/Portage/de")[Dateien und Verzeichnisse](/wiki/Handbook:HPPA/Portage/Files/de "Handbook:HPPA/Portage/Files/de")[Variablen](/wiki/Handbook:HPPA/Portage/Variables/de "Handbook:HPPA/Portage/Variables/de")[Mischen von Softwarezweigen](/wiki/Handbook:HPPA/Portage/Branches/de "Handbook:HPPA/Portage/Branches/de")[Zusätzliche Tools](/wiki/Handbook:HPPA/Portage/Tools/de "Handbook:HPPA/Portage/Tools/de")[Eigener Portage-Tree](/wiki/Handbook:HPPA/Portage/CustomTree/de "Handbook:HPPA/Portage/CustomTree/de")[Erweiterte Portage-Features](/wiki/Handbook:HPPA/Portage/Advanced/de "Handbook:HPPA/Portage/Advanced/de")[Netzwerk-Konfiguration](/wiki/Handbook:HPPA/Full/Networking/de "Handbook:HPPA/Full/Networking/de")[Zu Beginn](/wiki/Handbook:HPPA/Networking/Introduction/de "Handbook:HPPA/Networking/Introduction/de")[Fortgeschrittene Konfiguration](/wiki/Handbook:HPPA/Networking/Advanced/de "Handbook:HPPA/Networking/Advanced/de")[Modulare Vernetzung](/wiki/Handbook:HPPA/Networking/Modular/de "Handbook:HPPA/Networking/Modular/de")[Drahtlose Netzwerke](/wiki/Handbook:HPPA/Networking/Wireless/de "Handbook:HPPA/Networking/Wireless/de")[Funktionalität hinzufügen](/wiki/Handbook:HPPA/Networking/Extending/de "Handbook:HPPA/Networking/Extending/de")[Dynamisches Management](/wiki/Handbook:HPPA/Networking/Dynamic/de "Handbook:HPPA/Networking/Dynamic/de")

## Contents

- [1Einführung in blockorientierte Geräte](#Einf.C3.BChrung_in_blockorientierte_Ger.C3.A4te)
  - [1.1Blockorientierte Geräte](#Blockorientierte_Ger.C3.A4te)
  - [1.2Partitionen und Slices](#Partitionen_und_Slices)
- [2Ein Partitionsschema entwerfen](#Ein_Partitionsschema_entwerfen)
  - [2.1Wie viele Partitionen und wie groß?](#Wie_viele_Partitionen_und_wie_gro.C3.9F.3F)
  - [2.2Was ist mit dem Swap-Speicher?](#Was_ist_mit_dem_Swap-Speicher.3F)
- [3fdisk auf HPPA benutzen](#fdisk_auf_HPPA_benutzen)
- [4Erstellen von Dateisystemen](#Erstellen_von_Dateisystemen)
  - [4.1Einleitung](#Einleitung)
  - [4.2Dateisysteme](#Dateisysteme)
  - [4.3Dateisystem auf einer Partition anlegen](#Dateisystem_auf_einer_Partition_anlegen)
    - [4.3.1EFI system partition filesystem](#EFI_system_partition_filesystem)
    - [4.3.2Legacy BIOS boot partition filesystem](#Legacy_BIOS_boot_partition_filesystem)
    - [4.3.3Small ext4 partitions](#Small_ext4_partitions)
  - [4.4Aktivieren der Swap-Partition](#Aktivieren_der_Swap-Partition)
- [5Einhängen der Root-Partition](#Einh.C3.A4ngen_der_Root-Partition)

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

### Partitionen und Slices

Obwohl es theoretisch möglich wäre die gesamte Festplatte für die Unterbringung eines Linux Systems zu nutzen, wird das in der Praxis selten gemacht. Statt dessen teilt man das gesamte Festplatten Block-Device in kleinere, besser verwaltbare Block Devices auf. Auf den meisten Systemen nennt man diese Partitionen. Andere Architekturen verwenden eine ähnliche Technik, genannt "Slices".

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

## fdisk auf HPPA benutzen

Verwenden Sie fdisk um die benötigten Partitionen zu erzeugen:

`root #` `fdisk /dev/sda`

HPPA Maschinen verwenden PC Standard DOS Partitionstabellen. Drücken Sie einfach die `o` Taste, um eine neue DOS Partitionstabelle zu erstellen.

`Command (m for help):` `o`

```
Building a new DOS disklabel.

```

PALO (der HPPA Bootloader) benötigt eine besondere Partition, damit er funktioniert. Es muss eine mindestens 16 MB große Partition am Anfang der Festplatte für ihn erzeugt werden. Der Typ der Partition muss _f0 (Linux/PA-RISC boot)_ sein. Es ist auch möglich, die PALO-Partition als /boot zu verwenden.

**Wichtig**

Wenn dies vergessen wird und die Installation ohne eine spezielle PALO Partition fortgesetzt wird, kann das System möglicherweise nicht neugestartet werden. Falls die Festplatte größer als 2 GB ist stellen Sie bitte sicher, dass die Bootpartition innerhalb der ersten 2 GB der Festplatte liegt. PALO ist nicht in der Lage einen Kernel zu lesen, der außerhalb dieser 2 GB Grenze liegt.

DATEI **`/etc/fstab`** **Einfaches Standard-Partitionsschema**

```
/dev/sda2    /boot   ext2    noauto,noatime   1 1
/dev/sda3    none    swap    sw               0 0
/dev/sda4    /       ext4    noatime          0 0

```

In fdisk sieht so ein Partitionslayout folgendermaßen aus:

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

## Erstellen von Dateisystemen

**Warnung**

Wenn Sie ein SSD- oder NVMe-Laufwerk verwenden, prüfen Sie bitte, ob es ein Firmware-Upgrade benötigt. Insbesondere einige Intel-SSDs (600p und 6000p) benötigen ein Firmware-Upgrade für [kritische Fehlerbehebungen](https://bugzilla.redhat.com/show_bug.cgi?id=1402533), um Datenbeschädigungen zu vermeiden, die durch XFS-I/O-Nutzungsmuster verursacht werden (allerdings nicht durch einen Fehler des Dateisystems). smartctl kann helfen, das Modell und die Firmware-Version zu überprüfen.

### Einleitung

Nachdem die Partitionen angelegt wurden, ist es an der Zeit, Dateisysteme darauf anzulegen. Im nächsten Abschnitt werden die unterschiedlichen Dateisysteme beschrieben, die Linux unterstützt. Leser, die bereits wissen, welches Dateisystem sie verwenden wollen, können bei [Dateisystem auf einer Partition anlegen](/wiki/Handbook:HPPA/Installation/Networking/de#Applying_a_filesystem_to_a_partition "Handbook:HPPA/Installation/Networking/de") fortfahren. Alle anderen sollten weiterlesen, um mehr über die verfügbaren Dateisysteme zu erfahren ...

### Dateisysteme

Linux unterstützt mehrere Dutzend Dateisysteme, wobei allerdings viele davon für ganz spezielle Anwendungszwecke optimiert sind. Nur einige Dateisysteme gelten als stabil auf der hppa Architektur. Es ist ratsam, sich über Dateisysteme und deren Unterstützungsgrad zu informieren, damit Sie nicht für wichtige Partitionen ein eher experimentelles Dateisystem wählen. **XFS ist das empfohlene all-round Dateisystem für alle Plattformen.** Nachfolgend eine nicht-vollständige Auswahl von verfügbaren Dateisystemen.

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

Um beispielsweise die EFI System-Partition (/dev/sda2) als FAT32 und die root-Partition (/dev/sda4) als xfs zu formatieren (wie in dem Beispiel-Partitionsschema), würde man folgende Befehle verwenden:

`root #` `mkfs.xfs /dev/sda4`

{{#ifeq: 0\|1\|

#### EFI system partition filesystem

The EFI system partition (/dev/sda2) must be formatted as FAT32:

`root #` `mkfs.vfat -F 32 /dev/sda2`

#### Legacy BIOS boot partition filesystem

Systems booting via legacy BIOS with a MBR/DOS disklabel can use any filesystem format supported by the bootloader.

For example, to format with XFS:

`root #` `mkfs.xfs -c options=/usr/share/xfsprogs/mkfs/lts_6.18.conf /dev/sda2`

#### Small ext4 partitions

Bei der Verwendung von ext4 auf kleinen Partitionen (kleiner als 8 GiB), sollte das Dateisystem mit den passenden Optionen erstellt werden, um genügend Inodes zu reservieren. Dies kann mit einer der folgenden Anweisungen erfolgen:

`root #` `mkfs.ext4 -T small /dev/<device>`

Dies vervierfacht die Zahl der Inodes für ein angegebenes Dateisystem in der Regel, da es dessen "bytes-per-inode" (Bytes pro Inode) von 16 kB auf 4 kB pro Inode reduziert.

### Aktivieren der Swap-Partition

mkswap ist der Befehl der verwendet wird um Swap-Partitionen zu initialisieren:

`root #` `mkswap /dev/sda3`

**Hinweis**

Installations which were previously started, but did not finish the installation process can resume the installation from this point in the handbook. Use this link as the permalink: [Resumed installations start here](/wiki/Handbook:HPPA/Installation/Disks/de#Resumed_installations_start_here "Handbook:HPPA/Installation/Disks/de").

Zur Aktivierung der Swap-Partition verwenden Sie swapon:

`root #` `swapon /dev/sda3`

This 'activation' step is only necessary because the swap partition is newly created within the live environment. Once the system has been rebooted, as long as the swap partition is properly defined within fstab or other mount mechanism, swap space will activate automatically.

## Einhängen der Root-Partition

**Tipp**

Anwender, die ein Nicht-Gentoo Installationsmedium verwenden, müssen mit folgendem Befehl einen Mount-Point erzeugen:

`root #` `mkdir --parents /mnt/gentoo`

`root #` `mkdir --parents /mnt/gentoo`

Continue creating additional mount points necessary for any additional (custom) partition(s) created during previous steps by using the mkdir command.

Nachdem die Partitionen initialisiert wurden und ein Dateisystem beinhalten, ist es an der Zeit, diese einzuhängen. Verwenden Sie den Befehl mount, aber vergessen Sie nicht die notwendigen Einhänge-Verzeichnisse für jede Partition zu erzeugen. Als Beispiel hängen wir die Root-Partition ein:

Mount the root partition:

`root #` `mount /dev/sda4 /mnt/gentoo`

Continue mounting additional (custom) partitions as necessary using the mount command.

**Hinweis**

Wenn sich /tmp/ auf einer separaten Partition befinden muss, ändern Sie die Berechtigungen nach dem Einhängen:

`root #` `chmod 1777 /mnt/gentoo/tmp`

Dies gilt ebenfalls für /var/tmp.

In der Anleitung wird später das Dateisystem proc (eine virtuelle Schnittstelle zum Kernel) zusammen mit anderen Kernel Pseudo-Dateisystemen eingehängt. Zunächst installieren wir jedoch die [Gentoo Installationsdateien](/wiki/Handbook:HPPA/Installation/Stage/de "Handbook:HPPA/Installation/Stage/de").

[← Konfiguration des Netzwerks](/wiki/Handbook:HPPA/Installation/Networking/de "Previous part") [Anfang](/wiki/Handbook:HPPA/de "Handbook:HPPA/de") [Installation des Stage Archivs →](/wiki/Handbook:HPPA/Installation/Stage/de "Next part")