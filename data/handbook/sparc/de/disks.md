# Disks

Other languages:

- Deutsch
- [English](/wiki/Handbook:SPARC/Installation/Disks "Handbook:SPARC/Installation/Disks (100% translated)")
- [español](/wiki/Handbook:SPARC/Installation/Disks/es "Manual de Gentoo: SPARC/Instalación/Discos (100% translated)")
- [français](/wiki/Handbook:SPARC/Installation/Disks/fr "Handbook:SPARC/Installation/Disks/fr (100% translated)")
- [italiano](/wiki/Handbook:SPARC/Installation/Disks/it "Handbook:SPARC/Installation/Disks/it (50% translated)")
- [magyar](/wiki/Handbook:SPARC/Installation/Disks/hu "Handbook:SPARC/Installation/Disks/hu (100% translated)")
- [polski](/wiki/Handbook:SPARC/Installation/Disks/pl "Handbook:SPARC/Installation/Disks (100% translated)")
- [português do Brasil](/wiki/Handbook:SPARC/Installation/Disks/pt-br "Handbook:SPARC/Installation/Disks/pt-br (50% translated)")
- [русский](/wiki/Handbook:SPARC/Installation/Disks/ru "Handbook:SPARC/Installation/Disks (100% translated)")
- [தமிழ்](/wiki/Handbook:SPARC/Installation/Disks/ta "கையேடு:SPARC/நிறுவல்/வட்டுக்கள் (100% translated)")
- [中文（中国大陆）‎](/wiki/Handbook:SPARC/Installation/Disks/zh-cn "手册：SPARC/安装/准备磁盘 (100% translated)")
- [日本語](/wiki/Handbook:SPARC/Installation/Disks/ja "ハンドブック:SPARC/インストール/ディスク (100% translated)")
- [한국어](/wiki/Handbook:SPARC/Installation/Disks/ko "Handbook:SPARC/Installation/Disks/ko (100% translated)")

[SPARC Handbuch](/wiki/Handbook:SPARC/de "Handbook:SPARC/de")[Installation](/wiki/Handbook:SPARC/Full/Installation/de "Handbook:SPARC/Full/Installation/de")[Über die Installation](/wiki/Handbook:SPARC/Installation/About/de "Handbook:SPARC/Installation/About/de")[Auswahl des Mediums](/wiki/Handbook:SPARC/Installation/Media/de "Handbook:SPARC/Installation/Media/de")[Konfiguration des Netzwerks](/wiki/Handbook:SPARC/Installation/Networking/de "Handbook:SPARC/Installation/Networking/de")Vorbereiten der Festplatte(n)[Installation des Stage Archivs](/wiki/Handbook:SPARC/Installation/Stage/de "Handbook:SPARC/Installation/Stage/de")[Installation des Basissystems](/wiki/Handbook:SPARC/Installation/Base/de "Handbook:SPARC/Installation/Base/de")[Konfiguration des Kernels](/wiki/Handbook:SPARC/Installation/Kernel/de "Handbook:SPARC/Installation/Kernel/de")[Konfiguration des Systems](/wiki/Handbook:SPARC/Installation/System/de "Handbook:SPARC/Installation/System/de")[Installation der Tools](/wiki/Handbook:SPARC/Installation/Tools/de "Handbook:SPARC/Installation/Tools/de")[Konfiguration des Bootloaders](/wiki/Handbook:SPARC/Installation/Bootloader/de "Handbook:SPARC/Installation/Bootloader/de")[Abschluss](/wiki/Handbook:SPARC/Installation/Finalizing/de "Handbook:SPARC/Installation/Finalizing/de")[Arbeiten mit Gentoo](/wiki/Handbook:SPARC/Full/Working/de "Handbook:SPARC/Full/Working/de")[Portage-Einführung](/wiki/Handbook:SPARC/Working/Portage/de "Handbook:SPARC/Working/Portage/de")[USE-Flags](/wiki/Handbook:SPARC/Working/USE/de "Handbook:SPARC/Working/USE/de")[Portage-Features](/wiki/Handbook:SPARC/Working/Features/de "Handbook:SPARC/Working/Features/de")[Initskript-System](/wiki/Handbook:SPARC/Working/Initscripts/de "Handbook:SPARC/Working/Initscripts/de")[Umgebungsvariablen](/wiki/Handbook:SPARC/Working/EnvVar/de "Handbook:SPARC/Working/EnvVar/de")[Arbeiten mit Portage](/wiki/Handbook:SPARC/Full/Portage/de "Handbook:SPARC/Full/Portage/de")[Dateien und Verzeichnisse](/wiki/Handbook:SPARC/Portage/Files/de "Handbook:SPARC/Portage/Files/de")[Variablen](/wiki/Handbook:SPARC/Portage/Variables/de "Handbook:SPARC/Portage/Variables/de")[Mischen von Softwarezweigen](/wiki/Handbook:SPARC/Portage/Branches/de "Handbook:SPARC/Portage/Branches/de")[Zusätzliche Tools](/wiki/Handbook:SPARC/Portage/Tools/de "Handbook:SPARC/Portage/Tools/de")[Eigener Portage-Tree](/wiki/Handbook:SPARC/Portage/CustomTree/de "Handbook:SPARC/Portage/CustomTree/de")[Erweiterte Portage-Features](/wiki/Handbook:SPARC/Portage/Advanced/de "Handbook:SPARC/Portage/Advanced/de")[Netzwerk-Konfiguration](/wiki/Handbook:SPARC/Full/Networking/de "Handbook:SPARC/Full/Networking/de")[Zu Beginn](/wiki/Handbook:SPARC/Networking/Introduction/de "Handbook:SPARC/Networking/Introduction/de")[Fortgeschrittene Konfiguration](/wiki/Handbook:SPARC/Networking/Advanced/de "Handbook:SPARC/Networking/Advanced/de")[Modulare Vernetzung](/wiki/Handbook:SPARC/Networking/Modular/de "Handbook:SPARC/Networking/Modular/de")[Drahtlose Netzwerke](/wiki/Handbook:SPARC/Networking/Wireless/de "Handbook:SPARC/Networking/Wireless/de")[Funktionalität hinzufügen](/wiki/Handbook:SPARC/Networking/Extending/de "Handbook:SPARC/Networking/Extending/de")[Dynamisches Management](/wiki/Handbook:SPARC/Networking/Dynamic/de "Handbook:SPARC/Networking/Dynamic/de")

## Contents

- [1Einführung in blockorientierte Geräte](#Einf.C3.BChrung_in_blockorientierte_Ger.C3.A4te)
  - [1.1Blockorientierte Geräte](#Blockorientierte_Ger.C3.A4te)
  - [1.2Partition tables](#Partition_tables)
    - [1.2.1GUID Partition Table (GPT)](#GUID_Partition_Table_.28GPT.29)
    - [1.2.2Sun partition table](#Sun_partition_table)
  - [1.3Default partitioning scheme](#Default_partitioning_scheme)
    - [1.3.1GPT partition scheme](#GPT_partition_scheme)
    - [1.3.2Sun formatted partition scheme](#Sun_formatted_partition_scheme)
- [2Partitioning the disk with GPT](#Partitioning_the_disk_with_GPT)
  - [2.1Viewing the current partition layout](#Viewing_the_current_partition_layout)
  - [2.2Creating a new disklabel and removing all existing partitions](#Creating_a_new_disklabel_and_removing_all_existing_partitions)
  - [2.3Creating the BIOS boot partition](#Creating_the_BIOS_boot_partition)
  - [2.4Creating the swap partition](#Creating_the_swap_partition)
  - [2.5Creating the root partition](#Creating_the_root_partition)
  - [2.6Saving the partition layout](#Saving_the_partition_layout)
- [3Partitioning the disk with a Sun partition table](#Partitioning_the_disk_with_a_Sun_partition_table)
  - [3.1Viewing the current partition layout](#Viewing_the_current_partition_layout_2)
  - [3.2Creating a new disklabel / removing all partitions](#Creating_a_new_disklabel_.2F_removing_all_partitions)
  - [3.3Creating the whole disk partition](#Creating_the_whole_disk_partition)
  - [3.4Creating the root partition](#Creating_the_root_partition_2)
  - [3.5Creating the swap partition](#Creating_the_swap_partition_2)
  - [3.6Saving the partition layout](#Saving_the_partition_layout_2)
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

### Partition tables

Although it is theoretically possible to use a raw, unpartitioned disk to house a Linux system (when creating a btrfs RAID for example), this is almost never done in practice. Instead, disk block devices are split up into smaller, more manageable block devices. On **sparc** systems, these are called partitions. There are currently two standard partitioning technologies in use: Sun and GPT; the latter is supported only on more recent systems with a sufficiently recent firmware.

#### GUID Partition Table (GPT)

The _GUID Partition Table (GPT)_ setup (also called GPT disklabel) uses 64-bit identifiers for the partitions. The location in which it stores the partition information is much bigger than the 512 bytes of the MBR partition table (DOS disklabel), which means there is practically no limit on the amount of partitions for a GPT disk. Also the _size_ of a partition is bounded by a much greater limit (almost 8 ZiB - yes, zebibytes).

GPT also takes advantage of checksumming and redundancy. It carries CRC32 checksums to detect errors in the header and partition tables and has a backup GPT at the end of the disk. This backup table can be used to recover damage of the primary GPT near the beginning of the disk.

GPT is only supported on Oracle SPARC machines of the T4 generation or newer. Additionally, only certain more recent firmware includes GPT support. There are several methods to check whether GPT support is available.

From the OBP prompt, execute:

`{0} ok` `cd /packages/disk-label`

`{0} ok` `.properties`

gpt

supported-labels gpt

```
                   sun
                   mbr

```

name disk-label

If gpt is included in the output, then GPT support is available. Alternatively, this can be determined from the installation media without entering OBP. Use the prtconf command from [sys-apps/sparc-utils](https://packages.gentoo.org/packages/sys-apps/sparc-utils) to access this information from userspace:

`root #` `prtconf -pv | grep -c gpt`

Or, check if the file /sys/firmware/devicetree/base/packages/disk-label/gpt exists. If none of these methods succeeds, then a firmware update is required in order to support GPT.

#### Sun partition table

Systems not manufactured by Oracle, T3 or earlier systems, or systems running an earlier firmware must use the Sun partition table type.

The third partition on Sun systems is set aside as a special "whole disk" slice. This partition must **not** contain a file system.

Users who are used to the DOS partitioning scheme should note that Sun partition tables do not have "primary" and "extended" partitions. Instead, up to eight partitions are available per drive, with the third of these being reserved.

The Handbook authors suggest using [GPT](#GPT) whenever possible for Gentoo installations.

### Default partitioning scheme

Due to the differences in required partition layout between GPT and Sun partition tables, a single partitioning scheme is not sufficient to support all possible system requirements. Some example schemes are provided below.

#### GPT partition scheme

The following partitioning scheme will be used as an example for GPT-formatted disks:

Partition
Filesystem
Size
Mount Point
Description
/dev/sda1(none)
2M
none
[BIOS boot partition](/wiki/Handbook:X86/Blocks/Disks#What_is_the_BIOS_boot_partition.3F.2Fde "Handbook:X86/Blocks/Disks")/dev/sda2(swap)
RAM size \* 2
none
Swap partition
/dev/sda3ext4
Rest of the disk
/Root partition

#### Sun formatted partition scheme

The following partitioning scheme will be used as an example for Sun-formatted disks:

Partition
Filesystem
Size
Mount Point
Description
/dev/sda1ext4
Disk size minus swap
/Root partition
/dev/sda2(swap)
RAM size \* 2
none
Swap partition
/dev/sda3(none)
Whole disk
none
Whole disk partition.

Required on disks using the Sun partition table.

**Wichtig**

SPARC systems using OBP version 3 or older have additional restrictions on their partitioning scheme. The root partition must be the first partition on the disk, and it may be no larger than 2 GiB. For this reason, such systems will require additional sufficiently-sized partitions for top-level directories, such as /usr, /var, /home, and other directories which would likely cause the root partition to exceed this limit. These systems are also likely to require the Sun partition table type, so do not forget to include the whole disk partition.

## Partitioning the disk with GPT

The following parts explain how to create the example partition layout for a GPT installation using fdisk. The example partition layout was mentioned earlier:

Partition
Description
/dev/sda1Boot partition
/dev/sda2Swap partition
/dev/sda3Root partition

Change the partition layout according the system's needs.

### Viewing the current partition layout

fdisk is a popular and powerful tool to split a disk into partitions. Fire up fdisk against the disk (in the example, /dev/sda is used):

`root #` `fdisk /dev/sda`

Use the `p` key to display the disk's current partition configuration:

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

### Creating a new disklabel and removing all existing partitions

Type `g` to create a new GPT disklabel on the disk; this will remove all existing partitions.

`Command (m for help):` `g`

```
Created a new GPT disklabel (GUID: 9850A2C2-76C4-FC47-9F0B-DA60449D2413).

```

For an existing GPT disklabel (see the output of `p` above), alternatively consider removing the existing partitions one by one from the disk. Type `d` to delete a partition. For instance, to delete an existing /dev/sda1:

`Command (m for help):` `d`

```
Selected partition 1
Partition 1 has been deleted.

```

The partition has now been scheduled for deletion. It will no longer show up when printing the list of partitions ( `p`, but it will not be erased until the changes have been saved. This allows users to abort the operation if a mistake was made - in that case, type `q` immediately and hit `Enter` and the partition will not be deleted.

Repeatedly type `p` to print out a partition listing and then type `d` and the number of the partition to delete it. Eventually, the partition table will be empty:

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

Now that the in-memory partition table is empty, we're ready to create the partitions.

### Creating the BIOS boot partition

First, create the BIOS boot partition. Type `n` to create a new partition, followed by `1` to select the first partition. When prompted for the first sector, make sure it starts from 2048 (which may be needed for the boot loader) and hit `Enter`. When prompted for the last sector, type +2M to create a partition 2 Mbyte in size:

`Command (m for help):` `n`

```
Partition number (1-128, default 1):
First sector (2048-30548062, default 2048):
Last sector, +/-sectors or +/-size{K,M,G,T,P} (2048-30548062, default 30547967): +2M

Created a new partition 1 of type 'Linux filesystem' and of size 2 MiB.

```

Mark the partition as a BIOS boot partition:

`Command (m for help):` `t`

```
Selected partition 1
Partition type or alias (type L to list all): 4
Changed type of partition 'Linux filesystem' to 'BIOS boot'.

```

### Creating the swap partition

Next, to create the swap partition, type `n` to create a new partition, then type `2` to create the second partition, /dev/sda2. When prompted for the first sector, hit `Enter`. When prompted for the last sector, type +4G (or any other size needed for the swap space) to create a partition 4 GiB in size.

`Command (m for help):` `n`

```
Partition number (2-128, default 2):
First sector (6144-30548062, default 6144):
Last sector, +/-sectors or +/-size{K,M,G,T,P} (6144-30548062, default 30547967): +4G

Created a new partition 2 of type 'Linux filesystem' and of size 4 GiB.

```

After all this is done, type `t` to set the partition type, `2` to select the partition just created and then type in _19_ to set the partition type to "Linux Swap".

`Command (m for help):` `t`

```
Partition number (1,2, default 2): 2
Partition type (type L to list all types): 19

Changed type of partition 'Linux filesystem' to 'Linux swap'.

```

### Creating the root partition

Finally, to create the root partition, type `n` to create a new partition. Then type `3` to create the third partition, /dev/sda3. When prompted for the first sector, hit `Enter`. When prompted for the last sector, hit `Enter` to create a partition that takes up the rest of the remaining space on the disk. After completing these steps, typing `p` should display a partition table that looks similar to this:

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

### Saving the partition layout

To save the partition layout and exit fdisk, type `w`.

`Command (m for help):` `w`

With the partitions created, it is now time to put filesystems on them.

## Partitioning the disk with a Sun partition table

The following parts explain how to create the example partition layout for a Sun partition table installation using fdisk. The example partition layout was mentioned earlier:

Partition
Description
/dev/sda1Root partition
/dev/sda2Swap partition
/dev/sda3Whole disk partition

Change the partition layout according to personal preference. If partitioning for a system using OBP version 3 or earlier, ensure that the root partition is less than 2G in size, and additionally create partitions /dev/sda4 and onward for additional filesystems.

### Viewing the current partition layout

fdisk is a popular and powerful tool to split a disk into partitions. Fire up fdisk against the disk (in our example, we use /dev/sda):

`root #` `fdisk /dev/sda`

Use the `p` key to display the disk's current partition configuration:

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

### Creating a new disklabel / removing all partitions

Type `s` to create a new Sun disklabel on the disk; this will remove all existing partitions.

`Command (m for help):` `s`

```
Created a new partition 1 of type 'Linux native' and of size 14.5 GiB.
Created a new partition 2 of type 'Linux swap' and of size 50 MiB.
Created a new partition 3 of type 'Whole disk' and of size 14.6 GiB.
Created a new Sun disklabel.

```

For an existing Sun disklabel (see the output of `p` above), alternatively consider removing the existing partitions one by one from the disk. Type `d` to delete a partition. For instance, to delete an existing /dev/sda1:

`Command (m for help):` `d`

```
Partition number (1-3, default 3): 1

Partition 1 has been deleted.

```

The partition has now been scheduled for deletion. It will no longer show up when printing the list of partitions ( `p`, but it will not be erased until the changes have been saved. This allows users to abort the operation if a mistake was made - in that case, type `q` immediately and hit `Enter` and the partition will not be deleted.

Repeatedly type `p` to print out a partition listing and then type `d` and the number of the partition to delete it. Eventually, the partition table will be empty:

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

Now that the in-memory partition table is empty, we're ready to create the partitions.

### Creating the whole disk partition

First, create the whole disk partition. Type `n` to create a new partition, followed by `3` to select the third partition. When prompted for the first sector, hit `Enter`. When prompted for the last sector, hit `Enter` to create a partition that takes up all of the space on the disk.

`Command (m for help):` `n`

```
Partition number (1-8, default 1): 3

It is highly recommended that the third partition covers the whole disk and is of type `Whole disk'
First sector (0-30547968, default 0):
Last sector or +/-sectors or +/-size{K,M,G,T,P} (0-30547968, default 30547968):

Created a new partition 3 of type 'Whole disk' and of size 14.6 GiB.

```

fdisk will automatically set the type of such a partition to 'Whole disk', so there is no need to explicitly set the type.

### Creating the root partition

Next, to create the root partition, type `n` to create a new partition. Then type `1` to create the third partition, /dev/sda1. When prompted for the first sector, hit `Enter`. When prompted for the last sector, type -4G (or whatever space is required for non-root partitions). After completing these steps, typing `p` should display a partition table that looks similar to this:

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

### Creating the swap partition

Finally, to create the swap partition, type `n` to create a new partition, then type `2` to create the second partition, /dev/sda2. When prompted for the first sector, hit `Enter`. When prompted for the last sector, hit `Enter` to take up the remaining space on the disk.

`Command (m for help):` `n`

```
Partition number (2,4-8, default 2):
First sector (22159360-30547968, default 22159360):
Last sector or +/-sectors or +/-size{K,M,G,T,P} (22159360-30547968, default 30547968):

Created a new partition 2 of type 'Linux native' and of size 4 GiB.

```

After all this is done, type `t` to set the partition type, `2` to select the partition just created and then type in _82_ to set the partition type to "Linux Swap".

`Command (m for help):` `t`

```
Partition number (1-3, default 3): 2
Hex code (type L to list all codes): 82

Changed type of partition 'Linux native' to 'Linux swap'.

```

After completing these steps, typing `p` should display a partition table that looks similar to this:

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

### Saving the partition layout

To save the partition layout and exit fdisk, type `w`.

`Command (m for help):` `w`

With the partitions created, it is now time to put filesystems on them.

## Erstellen von Dateisystemen

**Warnung**

Wenn Sie ein SSD- oder NVMe-Laufwerk verwenden, prüfen Sie bitte, ob es ein Firmware-Upgrade benötigt. Insbesondere einige Intel-SSDs (600p und 6000p) benötigen ein Firmware-Upgrade für [kritische Fehlerbehebungen](https://bugzilla.redhat.com/show_bug.cgi?id=1402533), um Datenbeschädigungen zu vermeiden, die durch XFS-I/O-Nutzungsmuster verursacht werden (allerdings nicht durch einen Fehler des Dateisystems). smartctl kann helfen, das Modell und die Firmware-Version zu überprüfen.

### Einleitung

Nachdem die Partitionen angelegt wurden, ist es an der Zeit, Dateisysteme darauf anzulegen. Im nächsten Abschnitt werden die unterschiedlichen Dateisysteme beschrieben, die Linux unterstützt. Leser, die bereits wissen, welches Dateisystem sie verwenden wollen, können bei [Dateisystem auf einer Partition anlegen](/wiki/Handbook:SPARC/Installation/Networking/de#Applying_a_filesystem_to_a_partition "Handbook:SPARC/Installation/Networking/de") fortfahren. Alle anderen sollten weiterlesen, um mehr über die verfügbaren Dateisysteme zu erfahren ...

### Dateisysteme

Linux unterstützt mehrere Dutzend Dateisysteme, wobei allerdings viele davon für ganz spezielle Anwendungszwecke optimiert sind. Nur einige Dateisysteme gelten als stabil auf der sparc Architektur. Es ist ratsam, sich über Dateisysteme und deren Unterstützungsgrad zu informieren, damit Sie nicht für wichtige Partitionen ein eher experimentelles Dateisystem wählen. **XFS ist das empfohlene all-round Dateisystem für alle Plattformen.** Nachfolgend eine nicht-vollständige Auswahl von verfügbaren Dateisystemen.

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

Um beispielsweise die root-Partition (/dev/sda1) als xfs zu formatieren (wie in dem Beispiel-Partitionsschema), würde man folgende Befehle verwenden:

`root #` `mkfs.xfs /dev/sda1`

{{#ifeq: 0\|1\|

#### EFI system partition filesystem

The EFI system partition () must be formatted as FAT32:

#### Legacy BIOS boot partition filesystem

Systems booting via legacy BIOS with a MBR/DOS disklabel can use any filesystem format supported by the bootloader.

For example, to format with XFS:

`root #` `mkfs.xfs -c options=/usr/share/xfsprogs/mkfs/lts_6.18.conf`

#### Small ext4 partitions

Bei der Verwendung von ext4 auf kleinen Partitionen (kleiner als 8 GiB), sollte das Dateisystem mit den passenden Optionen erstellt werden, um genügend Inodes zu reservieren. Dies kann mit einer der folgenden Anweisungen erfolgen:

`root #` `mkfs.ext4 -T small /dev/<device>`

Dies vervierfacht die Zahl der Inodes für ein angegebenes Dateisystem in der Regel, da es dessen "bytes-per-inode" (Bytes pro Inode) von 16 kB auf 4 kB pro Inode reduziert.

### Aktivieren der Swap-Partition

mkswap ist der Befehl der verwendet wird um Swap-Partitionen zu initialisieren:

`root #` `mkswap /dev/sda2`

**Hinweis**

Installations which were previously started, but did not finish the installation process can resume the installation from this point in the handbook. Use this link as the permalink: [Resumed installations start here](/wiki/Handbook:SPARC/Installation/Disks/de#Resumed_installations_start_here "Handbook:SPARC/Installation/Disks/de").

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

`root #` `mount /dev/sda1 /mnt/gentoo`

Continue mounting additional (custom) partitions as necessary using the mount command.

**Hinweis**

Wenn sich /tmp/ auf einer separaten Partition befinden muss, ändern Sie die Berechtigungen nach dem Einhängen:

`root #` `chmod 1777 /mnt/gentoo/tmp`

Dies gilt ebenfalls für /var/tmp.

In der Anleitung wird später das Dateisystem proc (eine virtuelle Schnittstelle zum Kernel) zusammen mit anderen Kernel Pseudo-Dateisystemen eingehängt. Zunächst installieren wir jedoch die [Gentoo Installationsdateien](/wiki/Handbook:SPARC/Installation/Stage/de "Handbook:SPARC/Installation/Stage/de").

[← Konfiguration des Netzwerks](/wiki/Handbook:SPARC/Installation/Networking/de "Previous part") [Anfang](/wiki/Handbook:SPARC/de "Handbook:SPARC/de") [Installation des Stage Archivs →](/wiki/Handbook:SPARC/Installation/Stage/de "Next part")