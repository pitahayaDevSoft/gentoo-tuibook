# Disks

Other languages:

- Deutsch
- [English](/wiki/Handbook:PPC/Installation/Disks "Handbook:PPC/Installation/Disks (100% translated)")
- [español](/wiki/Handbook:PPC/Installation/Disks/es "Manual de Gentoo: PPC/Instalación/Discos (100% translated)")
- [français](/wiki/Handbook:PPC/Installation/Disks/fr "Handbook:PPC/Installation/Disks/fr (100% translated)")
- [italiano](/wiki/Handbook:PPC/Installation/Disks/it "Handbook:PPC/Installation/Disks/it (100% translated)")
- [magyar](/wiki/Handbook:PPC/Installation/Disks/hu "Handbook:PPC/Installation/Disks/hu (100% translated)")
- [polski](/wiki/Handbook:PPC/Installation/Disks/pl "Handbook:PPC/Installation/Disks (100% translated)")
- [português do Brasil](/wiki/Handbook:PPC/Installation/Disks/pt-br "Handbook:PPC/Installation/Disks/pt-br (50% translated)")
- [русский](/wiki/Handbook:PPC/Installation/Disks/ru "Handbook:PPC/Installation/Disks (100% translated)")
- [தமிழ்](/wiki/Handbook:PPC/Installation/Disks/ta "கையேடு:PPC/நிறுவல்/வட்டுக்கள் (100% translated)")
- [中文（中国大陆）‎](/wiki/Handbook:PPC/Installation/Disks/zh-cn "手册：PPC/安装/准备磁盘 (100% translated)")
- [日本語](/wiki/Handbook:PPC/Installation/Disks/ja "ハンドブック:PPC/インストール/ディスク (100% translated)")
- [한국어](/wiki/Handbook:PPC/Installation/Disks/ko "Handbook:PPC/Installation/Disks/ko (100% translated)")

[PPC Handbuch](/wiki/Handbook:PPC/de "Handbook:PPC/de")[Installation](/wiki/Handbook:PPC/Full/Installation/de "Handbook:PPC/Full/Installation/de")[Über die Installation](/wiki/Handbook:PPC/Installation/About/de "Handbook:PPC/Installation/About/de")[Auswahl des Mediums](/wiki/Handbook:PPC/Installation/Media/de "Handbook:PPC/Installation/Media/de")[Konfiguration des Netzwerks](/wiki/Handbook:PPC/Installation/Networking/de "Handbook:PPC/Installation/Networking/de")Vorbereiten der Festplatte(n)[Installation des Stage Archivs](/wiki/Handbook:PPC/Installation/Stage/de "Handbook:PPC/Installation/Stage/de")[Installation des Basissystems](/wiki/Handbook:PPC/Installation/Base/de "Handbook:PPC/Installation/Base/de")[Konfiguration des Kernels](/wiki/Handbook:PPC/Installation/Kernel/de "Handbook:PPC/Installation/Kernel/de")[Konfiguration des Systems](/wiki/Handbook:PPC/Installation/System/de "Handbook:PPC/Installation/System/de")[Installation der Tools](/wiki/Handbook:PPC/Installation/Tools/de "Handbook:PPC/Installation/Tools/de")[Konfiguration des Bootloaders](/wiki/Handbook:PPC/Installation/Bootloader/de "Handbook:PPC/Installation/Bootloader/de")[Abschluss](/wiki/Handbook:PPC/Installation/Finalizing/de "Handbook:PPC/Installation/Finalizing/de")[Arbeiten mit Gentoo](/wiki/Handbook:PPC/Full/Working/de "Handbook:PPC/Full/Working/de")[Portage-Einführung](/wiki/Handbook:PPC/Working/Portage/de "Handbook:PPC/Working/Portage/de")[USE-Flags](/wiki/Handbook:PPC/Working/USE/de "Handbook:PPC/Working/USE/de")[Portage-Features](/wiki/Handbook:PPC/Working/Features/de "Handbook:PPC/Working/Features/de")[Initskript-System](/wiki/Handbook:PPC/Working/Initscripts/de "Handbook:PPC/Working/Initscripts/de")[Umgebungsvariablen](/wiki/Handbook:PPC/Working/EnvVar/de "Handbook:PPC/Working/EnvVar/de")[Arbeiten mit Portage](/wiki/Handbook:PPC/Full/Portage/de "Handbook:PPC/Full/Portage/de")[Dateien und Verzeichnisse](/wiki/Handbook:PPC/Portage/Files/de "Handbook:PPC/Portage/Files/de")[Variablen](/wiki/Handbook:PPC/Portage/Variables/de "Handbook:PPC/Portage/Variables/de")[Mischen von Softwarezweigen](/wiki/Handbook:PPC/Portage/Branches/de "Handbook:PPC/Portage/Branches/de")[Zusätzliche Tools](/wiki/Handbook:PPC/Portage/Tools/de "Handbook:PPC/Portage/Tools/de")[Eigener Portage-Tree](/wiki/Handbook:PPC/Portage/CustomTree/de "Handbook:PPC/Portage/CustomTree/de")[Erweiterte Portage-Features](/wiki/Handbook:PPC/Portage/Advanced/de "Handbook:PPC/Portage/Advanced/de")[Netzwerk-Konfiguration](/wiki/Handbook:PPC/Full/Networking/de "Handbook:PPC/Full/Networking/de")[Zu Beginn](/wiki/Handbook:PPC/Networking/Introduction/de "Handbook:PPC/Networking/Introduction/de")[Fortgeschrittene Konfiguration](/wiki/Handbook:PPC/Networking/Advanced/de "Handbook:PPC/Networking/Advanced/de")[Modulare Vernetzung](/wiki/Handbook:PPC/Networking/Modular/de "Handbook:PPC/Networking/Modular/de")[Drahtlose Netzwerke](/wiki/Handbook:PPC/Networking/Wireless/de "Handbook:PPC/Networking/Wireless/de")[Funktionalität hinzufügen](/wiki/Handbook:PPC/Networking/Extending/de "Handbook:PPC/Networking/Extending/de")[Dynamisches Management](/wiki/Handbook:PPC/Networking/Dynamic/de "Handbook:PPC/Networking/Dynamic/de")

## Contents

- [1Einführung in blockorientierte Geräte](#Einf.C3.BChrung_in_blockorientierte_Ger.C3.A4te)
  - [1.1Blockorientierte Geräte](#Blockorientierte_Ger.C3.A4te)
  - [1.2Partitionen](#Partitionen)
- [2Designing a partition scheme](#Designing_a_partition_scheme)
  - [2.1How many partitions and how big?](#How_many_partitions_and_how_big.3F)
  - [2.2What about swap space?](#What_about_swap_space.3F)
  - [2.3Apple New World](#Apple_New_World)
  - [2.4Apple Old World](#Apple_Old_World)
  - [2.5Pegasos](#Pegasos)
  - [2.6IBM PReP (RS/6000)](#IBM_PReP_.28RS.2F6000.29)
- [3Using mac-fdisk (Apple)](#Using_mac-fdisk_.28Apple.29)
- [4Using parted (Pegasos and RS/6000)](#Using_parted_.28Pegasos_and_RS.2F6000.29)
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

### Partitionen

Although it is theoretically possible to use a full disk to house a Linux system, this is almost never done in practice. Instead, full disk block devices are split up in smaller, more manageable block devices. On most systems, these are called partitions.

**Hinweis**

In the remainder of the installation instructions, we will use the Pegasos example partition layout. Adjust to personal preference.

## Designing a partition scheme

### How many partitions and how big?

The design of disk partition layout is highly dependent on the demands of the system and the file system(s) applied to the device. If there are lots of users, then it is advised to have /home on a separate partition which will increase security and make backups and other types of maintenance easier. If Gentoo is being installed to perform as a mail server, then /var should be a separate partition as all mails are stored inside the /var directory. Game servers may have a separate /opt partition since most gaming server software is installed therein. The reason for these recommendations is similar to the /home directory: security, backups, and maintenance.

In most situations on Gentoo, /usr and /var should be kept relatively large in size. /usr hosts the majority of applications available on the system and the Linux kernel sources (under /usr/src). By default, /var hosts the Gentoo ebuild repository (located at /var/db/repos/gentoo) which, depending on the file system, generally consumes around 650 MiB of disk space. This space estimate _excludes_ the /var/cache/distfiles and /var/cache/binpkgs directories, which will gradually fill with source files and (optionally) binary packages respectively as they are added to the system.

How many partitions and how big very much depends on considering the trade-offs and choosing the best option for the circumstance. Separate partitions or volumes have the following advantages:

- Choose the best performing filesystem for each partition or volume.
- The entire system cannot run out of free space if one defunct tool is continuously writing files to a partition or volume.
- If necessary, file system checks are reduced in time, as multiple checks can be done in parallel (although this advantage is realized more with multiple disks than it is with multiple partitions).
- Security can be enhanced by mounting some partitions or volumes read-only, `nosuid` (setuid bits are ignored), `noexec` (executable bits are ignored), etc.

However, multiple partitions have certain disadvantages as well:

- If not configured properly, the system might have lots of free space on one partition and little free space on another.
- A separate partition for /usr/ may require the administrator to boot with an initramfs to mount the partition before other boot scripts start. Since the generation and maintenance of an initramfs is beyond the scope of this handbook, **we recommend that newcomers do not use a separate partition for /usr/**.
- There is also a 15-partition limit for SCSI and SATA unless the disk uses GPT labels.

**Hinweis**

Installations that intend to use systemd as the service and init system must have the /usr directory available at boot, either as part of the root filesystem or mounted via an initramfs.

### What about swap space?

Recommendations for swap space size
RAM sizeSuspend support?Hibernation support?
2 GB or less4GB4GB
2 to 8 GBRAM amount2 \* RAM
8 to 64 GB8 GB minimum, 16 maximum1.5 \* RAM
64 GB or greater8 GB minimumHibernation not recommended! Hibernation is _not_ recommended for systems with very large amounts of memory. While possible, the entire contents of memory must be written to disk in order to successfully hibernate. Writing tens of gigabytes (or worse!) out to disk can can take a considerable amount of time, especially when rotational disks are used. It is best to suspend in this scenario.

There is no perfect value for swap space size. The purpose of the space is to provide disk storage to the kernel when internal dynamic memory (RAM) is under pressure. A swap space allows for the kernel to move memory pages that are not likely to be accessed soon to disk (swap or page-out), which will free memory in RAM for the current task. Of course, if the pages swapped to disk are suddenly needed, they will need to be put back in memory (page-in) which will take considerably longer than reading from RAM (as disks are very slow compared to internal memory).

When a system is not going to run memory intensive applications or has lots of RAM available, then it probably does not need much swap space. However do note in case of hibernation that swap space is used to store _the entire contents of memory_ (likely on desktop and laptop systems rather than on server systems). If the system requires support for hibernation, then swap space larger than or equal to the amount of memory is necessary.

As a general rule for RAM amounts less than 4 GB, the swap space size is recommended to be twice the internal memory (RAM). For systems with multiple hard disks, it is wise to create one swap partition on each disk so that they can be utilized for parallel read/write operations. The faster a disk can swap, the faster the system will run when data in swap space must be accessed. When choosing between rotational and solid state disks, it is better for performance to put swap on the solid state hardware.

It is worth noting that swap _files_ can be used as an alternative to swap _partitions_; this is mostly helpful for systems with very limited disk space.

**Tipp**

With the rise in popularity with [zram](/wiki/Zram "Zram"), it should be noted that Gentoo is more likely to suffer with performance hits, as compiling uses large amounts of RAM, and programs can't be directly compiled in swap. It's usually better to create a smaller zram as a cache swap area with higher priority over the disk based swap of around 512MB, if a user really wants to use zram to benefit their desktop needs and not harm compile times.

### Apple New World

Apple New World machines are fairly straightforward to configure. The first partition is always an Apple Partition Map (APM). This partition keeps track of the layout of the disk. It is not possible to remove this partition. The next partition should always be a bootstrap partition. This partition contains a small (800KiB) HFS filesystem that holds a copy of the bootloader Yaboot and its configuration file. This partition is not the same as a /boot partition as found on other architectures. After the boot partition, the usual Linux filesystems are placed, according to the scheme below. The swap partition is a temporary storage place for when the system runs out of physical memory. The root partition will contain the filesystem that Gentoo is installed on. To dual boot, the OSX partition can go anywhere after the bootstrap partition to insure that yaboot starts first.

**Hinweis**

There may be "Disk Driver" partitions on the disk such as Apple\_Driver63, Apple\_Driver\_ATA, Apple\_FWDriver, Apple\_Driver\_IOKit, and Apple\_Patches. These are used to boot MacOS, so if there is no need for this, they can be removed by initializing the disk with mac-fdisk's `i` option. Be careful, this will completely erase the disk! If in doubt do not remove them.

**Hinweis**

If the disk is partitioned with Apple's Disk Utility, there may be 128 MiB spaces between partitions which Apple reserves for "future use". These can be safely removed.

Partition
Size
Filesystem
Description
/dev/sda132KiB
None.
Apple Partition Map (APM).
/dev/sda2800KiB
HFS
Apple bootstrap.
/dev/sda3512 MiB
swap
Linux swap (type 0x82).
/dev/sda4Rest of the disk.
ext4, xfs, etc.
Linux root.

### Apple Old World

Apple Old World machines are a bit more complicated to configure. The first partition is always an Apple Partition Map (APM). This partition keeps track of the layout of the disk. It is not possible to remove this partition. When using BootX, the configuration below assumes that MacOS is installed on a separate disk. If this is not the case, there will be additional partitions for "Apple Disk Drivers" such as Apple\_Driver63, Apple\_Driver\_ATA, Apple\_FWDriver, Apple\_Driver\_IOKit, Apple\_Patches and the MacOS install. When using Quik, it is necessary to create a boot partition to hold the kernel, unlike other Apple boot methods. After the boot partition, the usual Linux filesystems are placed, according to the scheme below. The swap partition is a temporary storage place for when the system runs out of physical memory. The root partition will contain the filesystem that Gentoo is installed on.

**Hinweis**

When using an Old World machine, it is necessary to keep MacOS available. The layout here assumes MacOS is installed on a separate drive.

Example partition layout for an Old World machine:

Partition
Size
Filesystem
Description
/dev/sda132KiB
None.
Apple Partition Map (APM).
/dev/sda232MiB
ext2
Quik Boot Partition (quik only).
/dev/sda3512MiB
swap
Linux swap (type 0x82).
/dev/sda4Rest of the disk.
ext4, xfs, etc.
Linux root.

### Pegasos

The Pegasos partition layout is quite simple compared to the Apple layouts. The first partition is a boot partition, which contains kernels to be booted along with an Open Firmware script that presents a menu on boot. After the boot partition, the usual Linux filesystems are placed, according to the scheme below. The swap partition is a temporary storage place for when the system runs out of physical memory. The root partition will contain the filesystem that Gentoo is installed on.

Example partition layout for Pegasos systems:

Partition
Size
Filesystem
Description
/dev/sda132MiB
affs1 or ext2
Boot partition.
/dev/sda2512MiB
swap
Linux swap (type 0x82).
/dev/sda3Rest of the disk.
ext4, xfs, etc.
Linux root.

### IBM PReP (RS/6000)

The IBM PowerPC Reference Platform (PReP) requires a small PReP boot partition on the disk's first partition, followed by the swap and root partitions.

Example partition layout for the IBM PReP:

Partition
Size
Filesystem
Description
/dev/sda1800KiB
None
PReP boot partition (type 0x41).
/dev/sda2512MiB
swap
Linux swap (type 0x82).
/dev/sda3Rest of the disk
ext4, xfs, etc.
Linux root (type 0x83).

**Warnung**

parted is able to resize partitions including HFS+. Unfortunately there may be issues with resizing HFS+ journaled filesystems, so, for the best results, switch off journaling in Mac OS X before resizing. Remember that any resizing operation is dangerous, so attempt at own risk! Be sure to always have a backup of all data before resizing!

## Using mac-fdisk (Apple)

At this point, create the partitions using mac-fdisk:

`root #` `mac-fdisk /dev/sda`

If Apple's Disk Utility was used prior to leave space for Linux, first delete the partitions that might have been created previously to make room for the new install. Use `d` in mac-fdisk to delete those partition(s). It will ask for the partition number to delete. Usually the first partition on NewWorld machines (Apple\_partition\_map) cannot be deleted. To start with a clean disk, simply initialize the disk by pressing `i`. This will completely erase the disk, so use this with caution.

Second, create an Apple\_Bootstrap partition by using `b`. It will ask for what block to start. Enter the number of the first free partition, followed by a `p`. For instance this is _2p_.

**Hinweis**

This partition is not a /boot partition. It is not used by Linux at all; there is no need to place any filesystem on it and it should never be mounted. Apple users don't need an extra partition for /boot.

Now create a swap partition by pressing `c`. Again mac-fdisk will ask for what block to start this partition from. As we used 2 before to create the Apple\_Bootstrap partition, now enter _3p_. When sked for the size, enter 512M (or whatever size needed - a minimum of 512MiB is recommended, but 2 times the physical memory is the generally accepted size). When asked for a name, enter _swap_.

To create the root partition, enter `c`, followed by _4p_ to select from what block the root partition should start. When asked for the size, enter _4p_ again. mac-fdisk will interpret this as "Use all available space". When asked for the name, enter _root_.

To finish up, write the partition to the disk using `w` and `q` to quit mac-fdisk.

**Hinweis**

To make sure everything is okay, run mac-fdisk -l and check whether all the partitions are there. If not all partitions created previously are shown, or the changes made are not reflected in the output, reinitialize the partitions by pressing `i` in mac-fdisk. Note that this will recreate the partition map and thus remove all existing partitions.

## Using parted (Pegasos and RS/6000)

parted, the partition editor, can now handle HFS+ partitions used by Mac OS and Mac OS X. With this tool it is possible to resize the Mac partitions and create space for the Linux partitions. Nevertheless, the example below describes partitioning for Pegasos machines only.

To begin let's fire up parted:

`root #` `parted /dev/sda`

If the drive is unpartitioned, run mklabel amiga to create a new disklabel for the drive.

It is possible to type print at any time in parted to display the current partition table. To abort parted, press `Ctrl` + `c`.

If next to Linux, the system is also meant to have MorphOS installed, then create an affs1 filesystem at the start of the drive. 32MB should be more than enough to store the MorphOS kernel. With a Pegasos I, or when Linux will use any filesystem besides ext2 or ext3, then it is necessary to also store the Linux kernel on this partition (the Pegasos II can only boot from ext2/ext3 or affs1 partitions). To create the partition run `mkpart primary affs1 START END` where START and END should be replaced with the megabyte range (e.g. 0 32) which creates a 32 MB partition starting at 0MB and ending at 32MB. When creating an ext2 or ext3 partition instead, substitute ext2 or ext3 for affs1 in the mkpart command.

Create two partitions for Linux, one root filesystem and one swap partition. Run `mkpart primary START END` to create each partition, replacing START and END with the desired megabyte boundaries.

It is generally recommended to create a swap partition that is two times bigger than the amount of RAM in the computer, but at least 512MiB is recommended. To create the swap partition, run `mkpart primary linux-swap START END` with START and END again denoting the partition boundaries.

When done in parted simply type `quit`.

## Erstellen von Dateisystemen

**Warnung**

Wenn Sie ein SSD- oder NVMe-Laufwerk verwenden, prüfen Sie bitte, ob es ein Firmware-Upgrade benötigt. Insbesondere einige Intel-SSDs (600p und 6000p) benötigen ein Firmware-Upgrade für [kritische Fehlerbehebungen](https://bugzilla.redhat.com/show_bug.cgi?id=1402533), um Datenbeschädigungen zu vermeiden, die durch XFS-I/O-Nutzungsmuster verursacht werden (allerdings nicht durch einen Fehler des Dateisystems). smartctl kann helfen, das Modell und die Firmware-Version zu überprüfen.

### Einleitung

Nachdem die Partitionen angelegt wurden, ist es an der Zeit, Dateisysteme darauf anzulegen. Im nächsten Abschnitt werden die unterschiedlichen Dateisysteme beschrieben, die Linux unterstützt. Leser, die bereits wissen, welches Dateisystem sie verwenden wollen, können bei [Dateisystem auf einer Partition anlegen](/wiki/Handbook:PPC/Installation/Networking/de#Applying_a_filesystem_to_a_partition "Handbook:PPC/Installation/Networking/de") fortfahren. Alle anderen sollten weiterlesen, um mehr über die verfügbaren Dateisysteme zu erfahren ...

### Dateisysteme

Linux unterstützt mehrere Dutzend Dateisysteme, wobei allerdings viele davon für ganz spezielle Anwendungszwecke optimiert sind. Nur einige Dateisysteme gelten als stabil auf der ppc Architektur. Es ist ratsam, sich über Dateisysteme und deren Unterstützungsgrad zu informieren, damit Sie nicht für wichtige Partitionen ein eher experimentelles Dateisystem wählen. **XFS ist das empfohlene all-round Dateisystem für alle Plattformen.** Nachfolgend eine nicht-vollständige Auswahl von verfügbaren Dateisystemen.

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

Installations which were previously started, but did not finish the installation process can resume the installation from this point in the handbook. Use this link as the permalink: [Resumed installations start here](/wiki/Handbook:PPC/Installation/Disks/de#Resumed_installations_start_here "Handbook:PPC/Installation/Disks/de").

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

In der Anleitung wird später das Dateisystem proc (eine virtuelle Schnittstelle zum Kernel) zusammen mit anderen Kernel Pseudo-Dateisystemen eingehängt. Zunächst installieren wir jedoch die [Gentoo Installationsdateien](/wiki/Handbook:PPC/Installation/Stage/de "Handbook:PPC/Installation/Stage/de").

[← Konfiguration des Netzwerks](/wiki/Handbook:PPC/Installation/Networking/de "Previous part") [Anfang](/wiki/Handbook:PPC/de "Handbook:PPC/de") [Installation des Stage Archivs →](/wiki/Handbook:PPC/Installation/Stage/de "Next part")