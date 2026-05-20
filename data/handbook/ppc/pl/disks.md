# Disks

Other languages:

- [Deutsch](/wiki/Handbook:PPC/Installation/Disks/de "Handbuch:PPC/Installation/Festplatten (100% translated)")
- [English](/wiki/Handbook:PPC/Installation/Disks "Handbook:PPC/Installation/Disks (100% translated)")
- [español](/wiki/Handbook:PPC/Installation/Disks/es "Manual de Gentoo: PPC/Instalación/Discos (100% translated)")
- [français](/wiki/Handbook:PPC/Installation/Disks/fr "Handbook:PPC/Installation/Disks/fr (100% translated)")
- [italiano](/wiki/Handbook:PPC/Installation/Disks/it "Handbook:PPC/Installation/Disks/it (100% translated)")
- [magyar](/wiki/Handbook:PPC/Installation/Disks/hu "Handbook:PPC/Installation/Disks/hu (100% translated)")
- polski
- [português do Brasil](/wiki/Handbook:PPC/Installation/Disks/pt-br "Handbook:PPC/Installation/Disks/pt-br (50% translated)")
- [русский](/wiki/Handbook:PPC/Installation/Disks/ru "Handbook:PPC/Installation/Disks (100% translated)")
- [தமிழ்](/wiki/Handbook:PPC/Installation/Disks/ta "கையேடு:PPC/நிறுவல்/வட்டுக்கள் (100% translated)")
- [中文（中国大陆）‎](/wiki/Handbook:PPC/Installation/Disks/zh-cn "手册：PPC/安装/准备磁盘 (100% translated)")
- [日本語](/wiki/Handbook:PPC/Installation/Disks/ja "ハンドブック:PPC/インストール/ディスク (100% translated)")
- [한국어](/wiki/Handbook:PPC/Installation/Disks/ko "Handbook:PPC/Installation/Disks/ko (100% translated)")

[PPC Handbook](/wiki/Handbook:PPC/pl "Handbook:PPC/pl")[Instalacja](/wiki/Handbook:PPC/Full/Installation/pl "Handbook:PPC/Full/Installation/pl")[O instalacji](/wiki/Handbook:PPC/Installation/About/pl "Handbook:PPC/Installation/About/pl")[Wybór medium instalacyjnego](/wiki/Handbook:PPC/Installation/Media/pl "Handbook:PPC/Installation/Media/pl")[Konfiguracja sieci](/wiki/Handbook:PPC/Installation/Networking/pl "Handbook:PPC/Installation/Networking/pl")Przygotowanie dysków[Instalacja etapu 3](/wiki/Handbook:PPC/Installation/Stage/pl "Handbook:PPC/Installation/Stage/pl")[Instalacja systemu podstawowego](/wiki/Handbook:PPC/Installation/Base/pl "Handbook:PPC/Installation/Base/pl")[Konfiguracja jądra](/wiki/Handbook:PPC/Installation/Kernel/pl "Handbook:PPC/Installation/Kernel/pl")[Konfiguracja systemu](/wiki/Handbook:PPC/Installation/System/pl "Handbook:PPC/Installation/System/pl")[Instalacja narzędzi](/wiki/Handbook:PPC/Installation/Tools/pl "Handbook:PPC/Installation/Tools/pl")[Instalacja systemu rozruchowego](/wiki/Handbook:PPC/Installation/Bootloader/pl "Handbook:PPC/Installation/Bootloader/pl")[Finalizacja](/wiki/Handbook:PPC/Installation/Finalizing/pl "Handbook:PPC/Installation/Finalizing/pl")[Praca z Gentoo](/wiki/Handbook:PPC/Full/Working/pl "Handbook:PPC/Full/Working/pl")[Wstęp do Portage](/wiki/Handbook:PPC/Working/Portage/pl "Handbook:PPC/Working/Portage/pl")[Flagi USE](/wiki/Handbook:PPC/Working/USE/pl "Handbook:PPC/Working/USE/pl")[Funkcje portage](/wiki/Handbook:PPC/Working/Features/pl "Handbook:PPC/Working/Features/pl")[System initscript](/wiki/Handbook:PPC/Working/Initscripts/pl "Handbook:PPC/Working/Initscripts/pl")[Zmienne środowiskowe](/wiki/Handbook:PPC/Working/EnvVar/pl "Handbook:PPC/Working/EnvVar/pl")[Praca z Portage](/wiki/Handbook:PPC/Full/Portage/pl "Handbook:PPC/Full/Portage/pl")[Pliki i katalogi](/wiki/Handbook:PPC/Portage/Files/pl "Handbook:PPC/Portage/Files/pl")[Zmienne](/wiki/Handbook:PPC/Portage/Variables/pl "Handbook:PPC/Portage/Variables/pl")[Mieszanie działów oprogramowania](/wiki/Handbook:PPC/Portage/Branches/pl "Handbook:PPC/Portage/Branches/pl")[Dodatkowe narzędzia](/wiki/Handbook:PPC/Portage/Tools/pl "Handbook:PPC/Portage/Tools/pl")[Custom package repository](/wiki/Handbook:PPC/Portage/CustomTree/pl "Handbook:PPC/Portage/CustomTree/pl")[Funkcje zaawansowane](/wiki/Handbook:PPC/Portage/Advanced/pl "Handbook:PPC/Portage/Advanced/pl")[Konfiguracja sieci](/wiki/Handbook:PPC/Full/Networking/pl "Handbook:PPC/Full/Networking/pl")[Zaczynamy](/wiki/Handbook:PPC/Networking/Introduction/pl "Handbook:PPC/Networking/Introduction/pl")[Zaawansowana konfiguracja](/wiki/Handbook:PPC/Networking/Advanced/pl "Handbook:PPC/Networking/Advanced/pl")[Sieć modularna](/wiki/Handbook:PPC/Networking/Modular/pl "Handbook:PPC/Networking/Modular/pl")[Sieć bezprzewodowa](/wiki/Handbook:PPC/Networking/Wireless/pl "Handbook:PPC/Networking/Wireless/pl")[Dodawanie funkcjonalności](/wiki/Handbook:PPC/Networking/Extending/pl "Handbook:PPC/Networking/Extending/pl")[Dynamiczne zarządzanie](/wiki/Handbook:PPC/Networking/Dynamic/pl "Handbook:PPC/Networking/Dynamic/pl")

## Contents

- [1Introduction to block devices](#Introduction_to_block_devices)
  - [1.1Block devices](#Block_devices)
  - [1.2Partitions](#Partitions)
- [2Designing a partition scheme](#Designing_a_partition_scheme)
  - [2.1How many partitions and how big?](#How_many_partitions_and_how_big.3F)
  - [2.2What about swap space?](#What_about_swap_space.3F)
  - [2.3Apple New World](#Apple_New_World)
  - [2.4Apple Old World](#Apple_Old_World)
  - [2.5Pegasos](#Pegasos)
  - [2.6IBM PReP (RS/6000)](#IBM_PReP_.28RS.2F6000.29)
- [3Using mac-fdisk (Apple)](#Using_mac-fdisk_.28Apple.29)
- [4Using parted (Pegasos and RS/6000)](#Using_parted_.28Pegasos_and_RS.2F6000.29)
- [5Creating file systems](#Creating_file_systems)
  - [5.1Introduction](#Introduction)
  - [5.2Filesystems](#Filesystems)
  - [5.3Applying a filesystem to a partition](#Applying_a_filesystem_to_a_partition)
    - [5.3.1Legacy BIOS boot partition filesystem](#Legacy_BIOS_boot_partition_filesystem)
    - [5.3.2Small ext4 partitions](#Small_ext4_partitions)
  - [5.4Activating the swap partition](#Activating_the_swap_partition)
- [6Mounting the root partition](#Mounting_the_root_partition)

## Introduction to block devices

### Block devices

Let's take a good look at disk-oriented aspects of Gentoo Linux and Linux in general, including block devices, partitions, and Linux filesystems. Once the ins and outs of disks are understood, partitions and filesystems can be established for installation.

To begin, let's look at block devices. SCSI and Serial ATA drives are both labeled under device handles such as: /dev/sda, /dev/sdb, /dev/sdc, etc. On more modern machines, PCI Express based NVMe solid state disks have device handles such as /dev/nvme0n1, /dev/nvme0n2, etc.

The following table will help readers determine where to find a certain type of block device on the system:

Type of deviceDefault device handleEditorial notes and considerations
IDE, SATA, SAS, SCSI, or USB flash/dev/sdaFound on hardware from roughly 2007 until the present, this device handle is perhaps the most commonly used in Linux. These types of devices can be connected via the [SATA bus](https://en.wikipedia.org/wiki/Serial_ATA "wikipedia:Serial ATA"), [SCSI](https://en.wikipedia.org/wiki/SCSI "wikipedia:SCSI"), [USB](https://en.wikipedia.org/wiki/USB "wikipedia:USB") bus as block storage. As example, the first partition on the first SATA device is called /dev/sda1.
NVM Express (NVMe)/dev/nvme0n1The latest in solid state technology, [NVMe](https://en.wikipedia.org/wiki/NVM_Express "wikipedia:NVM Express") drives are connected to the PCI Express bus and have the fastest transfer block speeds on the market. Systems from around 2014 and newer may have support for NVMe hardware. The first partition on the first NVMe device is called /dev/nvme0n1p1.
MMC, eMMC, and SD/dev/mmcblk0[embedded MMC](https://en.wikipedia.org/wiki/MultiMediaCard#eMMC "wikipedia:MultiMediaCard") devices, SD cards, and other types of memory cards can be useful for data storage. That said, many systems may not permit booting from these types of devices. It is suggested to **not** use these devices for active Linux installations; rather consider using them to transfer files, which is their typical design intention. Alternatively this storage type could be useful for short-term file backups or snapshots.
VirtIO/dev/vda[Virtualized](/wiki/Virtualization "Virtualization") devices are found only within a [QEMU](/wiki/QEMU "QEMU") virtual emulator. `virtio_blk` driver provides this access to host's allocated drive space for this virtual machine. That said, it is a great way to try out Gentoo inside a virtual machine.

The block devices above represent an abstract interface to the disk. User programs can use these block devices to interact with the disk without worrying about whether the drives are SATA, SCSI, or something else. The program can simply address the storage on the disk as a bunch of contiguous, randomly-accessible 4096-byte (4K) blocks.

### Partitions

Although it is theoretically possible to use a full disk to house a Linux system, this is almost never done in practice. Instead, full disk block devices are split up in smaller, more manageable block devices. On most systems, these are called partitions.

**Informacja**

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

**Informacja**

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

**Wskazówka**

With the rise in popularity with [zram](/wiki/Zram "Zram"), it should be noted that Gentoo is more likely to suffer with performance hits, as compiling uses large amounts of RAM, and programs can't be directly compiled in swap. It's usually better to create a smaller zram as a cache swap area with higher priority over the disk based swap of around 512MB, if a user really wants to use zram to benefit their desktop needs and not harm compile times.

### Apple New World

Apple New World machines are fairly straightforward to configure. The first partition is always an Apple Partition Map (APM). This partition keeps track of the layout of the disk. It is not possible to remove this partition. The next partition should always be a bootstrap partition. This partition contains a small (800KiB) HFS filesystem that holds a copy of the bootloader Yaboot and its configuration file. This partition is not the same as a /boot partition as found on other architectures. After the boot partition, the usual Linux filesystems are placed, according to the scheme below. The swap partition is a temporary storage place for when the system runs out of physical memory. The root partition will contain the filesystem that Gentoo is installed on. To dual boot, the OSX partition can go anywhere after the bootstrap partition to insure that yaboot starts first.

**Informacja**

There may be "Disk Driver" partitions on the disk such as Apple\_Driver63, Apple\_Driver\_ATA, Apple\_FWDriver, Apple\_Driver\_IOKit, and Apple\_Patches. These are used to boot MacOS, so if there is no need for this, they can be removed by initializing the disk with mac-fdisk's `i` option. Be careful, this will completely erase the disk! If in doubt do not remove them.

**Informacja**

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

**Informacja**

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

**Uwaga**

parted is able to resize partitions including HFS+. Unfortunately there may be issues with resizing HFS+ journaled filesystems, so, for the best results, switch off journaling in Mac OS X before resizing. Remember that any resizing operation is dangerous, so attempt at own risk! Be sure to always have a backup of all data before resizing!

## Using mac-fdisk (Apple)

At this point, create the partitions using mac-fdisk:

`root #` `mac-fdisk /dev/sda`

If Apple's Disk Utility was used prior to leave space for Linux, first delete the partitions that might have been created previously to make room for the new install. Use `d` in mac-fdisk to delete those partition(s). It will ask for the partition number to delete. Usually the first partition on NewWorld machines (Apple\_partition\_map) cannot be deleted. To start with a clean disk, simply initialize the disk by pressing `i`. This will completely erase the disk, so use this with caution.

Second, create an Apple\_Bootstrap partition by using `b`. It will ask for what block to start. Enter the number of the first free partition, followed by a `p`. For instance this is _2p_.

**Informacja**

This partition is not a /boot partition. It is not used by Linux at all; there is no need to place any filesystem on it and it should never be mounted. Apple users don't need an extra partition for /boot.

Now create a swap partition by pressing `c`. Again mac-fdisk will ask for what block to start this partition from. As we used 2 before to create the Apple\_Bootstrap partition, now enter _3p_. When sked for the size, enter 512M (or whatever size needed - a minimum of 512MiB is recommended, but 2 times the physical memory is the generally accepted size). When asked for a name, enter _swap_.

To create the root partition, enter `c`, followed by _4p_ to select from what block the root partition should start. When asked for the size, enter _4p_ again. mac-fdisk will interpret this as "Use all available space". When asked for the name, enter _root_.

To finish up, write the partition to the disk using `w` and `q` to quit mac-fdisk.

**Informacja**

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

## Creating file systems

**Uwaga**

When using an SSD or NVMe drive, it is wise to check for firmware upgrades. Some Intel SSDs in particular (600p and 6000p) require a firmware upgrade for [possible data corruption](https://bugzilla.redhat.com/show_bug.cgi?id=1402533) induced by XFS I/O usage patterns. The problem is at the firmware level and not any fault of the XFS filesystem. The smartctl utility can help check the device model and firmware version.

### Introduction

Now that the partitions have been created, it is time to place a filesystem on them. In the next section the various file systems that Linux supports are described. Readers that already know which filesystem to use can continue with [Applying a filesystem to a partition](/wiki/Handbook:PPC/Installation/Disks/pl#Applying_a_filesystem_to_a_partition "Handbook:PPC/Installation/Disks/pl"). The others should read on to learn about the available filesystems...

### Filesystems

Linux supports several dozen filesystems, although many of them are only wise to deploy for specific purposes. Only certain filesystems may be found stable on the ppc architecture - it is advised to read up on the filesystems and their support state before selecting a more experimental one for important partitions. **XFS is the recommended all-purpose, all-platform filesystem.** The below is a non-exhaustive list:

[XFS](/wiki/XFS "XFS")Filesystem with metadata journaling which comes with a robust feature-set and is optimized for scalability. It has been continuously upgraded to include modern features. The only downside is that XFS partitions cannot yet be shrunk, although this is being worked on. XFS notably supports reflinks and Copy on Write (CoW) which is particularly helpful on Gentoo systems because of the amount of compiles users complete. XFS is the recommended modern all-purpose all-platform filesystem. Requires a partition to be at least 300MB.[ext4](/wiki/Ext4 "Ext4")Ext4 is a reliable, all-purpose all-platform filesystem, although it lacks modern features like reflinks.[VFAT](/wiki/VFAT "VFAT")Also known as FAT32, is supported by Linux but does not support standard UNIX permission settings. It is mostly used for interoperability/interchange with other operating systems (Microsoft Windows or Apple's macOS) but is also a necessity for some system bootloader firmware (like UEFI). Users of UEFI systems will need an [EFI System Partition](/wiki/EFI_System_Partition "EFI System Partition") formatted with VFAT in order to boot.[btrfs](/wiki/Btrfs/pl "Btrfs/pl")Newer generation filesystem. Provides advanced features like snapshotting, self-healing through checksums, transparent compression, subvolumes, and integrated RAID. Kernels prior to 5.4.y are not guaranteed to be safe to use with btrfs in production because fixes for serious issues are only present in the more recent releases of the LTS kernel branches. RAID 5/6 and quota groups unsafe on all versions of btrfs.[F2FS](/wiki/F2FS "F2FS")The Flash-Friendly File System was originally created by Samsung for the use with NAND flash memory. It is a decent choice when installing Gentoo to microSD cards, USB drives, or other flash-based storage devices.[NTFS](/wiki/NTFS "NTFS")This "New Technology" filesystem is the flagship filesystem of Microsoft Windows since Windows NT 3.1. Similarly to VFAT, it does not store UNIX permission settings or extended attributes necessary for BSD or Linux to function properly, therefore it should not be used as a root filesystem for most cases. It should _only_ be used for interoperability or data interchange with Microsoft Windows systems (note the emphasis on _only_).[ZFS](/wiki/ZFS "ZFS")**Ważne:** ZFS pools can be created on the minimal installation CD, for further information, refer to [ZFS/rootfs](/wiki/ZFS/rootfs "ZFS/rootfs")Next generation file system created by Matthew Ahrens and Jeff Bonwick. It was designed around a few key ideas: Administration of storage should be simple, redundancy should be handled by the filesystem, file systems should never be taken offline for repair, automated simulations of worst case scenarios before shipping code is important, and data integrity is paramount.

More extensive information on filesystems can be found in the community maintained [Filesystem article](/wiki/Filesystem "Filesystem").

### Applying a filesystem to a partition

**Informacja**

Please make sure to emerge the relevant user space utilities package for the chosen filesystem before rebooting. There will be a reminder to do so near the end of the installation process.

To create a filesystem on a partition or volume, there are user space utilities available for each possible filesystem. Click the filesystem's name in the table below for additional information on each filesystem:

Filesystem
Creation command
Within the live environment?
Package
[XFS](/wiki/XFS "XFS")mkfs.xfs Yes
[sys-fs/xfsprogs](https://packages.gentoo.org/packages/sys-fs/xfsprogs)[ext4](/wiki/Ext4 "Ext4")mkfs.ext4 Yes
[sys-fs/e2fsprogs](https://packages.gentoo.org/packages/sys-fs/e2fsprogs)[VFAT](/wiki/VFAT "VFAT") (FAT32, ...)
mkfs.vfat Yes
[sys-fs/dosfstools](https://packages.gentoo.org/packages/sys-fs/dosfstools)[btrfs](/wiki/Btrfs/pl "Btrfs/pl")mkfs.btrfs Yes
[sys-fs/btrfs-progs](https://packages.gentoo.org/packages/sys-fs/btrfs-progs)[F2FS](/wiki/F2FS "F2FS")mkfs.f2fs Yes
[sys-fs/f2fs-tools](https://packages.gentoo.org/packages/sys-fs/f2fs-tools)[NTFS](/wiki/NTFS "NTFS")mkfs.ntfs Yes
[sys-fs/ntfs3g](https://packages.gentoo.org/packages/sys-fs/ntfs3g)[ZFS](/wiki/ZFS "ZFS")zpool create ... Yes
[sys-fs/zfs](https://packages.gentoo.org/packages/sys-fs/zfs)

**Ważne**

The handbook recommends new partitions as part of the installation process, but it is important to note running any mkfs command will erase any data contained within the partition. When necessary, ensure any data that exists within is appropriately backed up before creating a new filesystem.

For instance, to have the root partition (/dev/sda3) as xfs as used in the example partition structure, the following commands would be used:

`root #` `mkfs.xfs -c options=/usr/share/xfsprogs/mkfs/lts_6.18.conf /dev/sda3`

#### Legacy BIOS boot partition filesystem

Systems booting via legacy BIOS with a MBR/DOS disklabel can use any filesystem format supported by the bootloader.

For example, to format with XFS:

`root #` `mkfs.xfs -c options=/usr/share/xfsprogs/mkfs/lts_6.18.conf /dev/sda1`

#### Small ext4 partitions

When using the ext4 filesystem on a small partition (less than 8 GiB), the filesystem should be created with the proper options to reserve enough inodes. This can specified using the `-T small` option:

`root #` `mkfs.ext4 -T small /dev/<device>`

Doing so will quadruple the number of inodes for a given filesystem, since its "bytes-per-inode" reduces from one every 16kB to one every 4kB.

### Activating the swap partition

mkswap is the command that is used to initialize swap partitions:

`root #` `mkswap /dev/sda2`

**Informacja**

Installations which were previously started, but did not finish the installation process can resume the installation from this point in the handbook. Use this link as the permalink: [Resumed installations start here](/wiki/Handbook:PPC/Installation/Disks/pl#Resumed_installations_start_here "Handbook:PPC/Installation/Disks/pl").

To activate the swap partition, use swapon:

`root #` `swapon /dev/sda2`

This 'activation' step is only necessary because the swap partition is newly created within the live environment. Once the system has been rebooted, as long as the swap partition is properly defined within fstab or other mount mechanism, swap space will activate automatically.

## Mounting the root partition

Certain live environments may be missing the suggested mount point for Gentoo's root partition (/mnt/gentoo), or mount points for additional partitions created in the partitioning section:

`root #` `mkdir --parents /mnt/gentoo`

Continue creating additional mount points necessary for any additional (custom) partition(s) created during previous steps by using the mkdir command.

With mount points created, it is time to make the partitions accessible via mount command.

Mount the root partition:

`root #` `mount /dev/sda3 /mnt/gentoo`

Continue mounting additional (custom) partitions as necessary using the mount command.

**Informacja**

If /tmp/ needs to reside on a separate partition, be sure to change its permissions after mounting:

`root #` `chmod 1777 /mnt/gentoo/tmp`

This also holds for /var/tmp.

Later in the instructions, the proc filesystem (a virtual interface with the kernel) as well as other kernel pseudo-filesystems will be mounted. But first [the Gentoo stage file](/wiki/Handbook:PPC/Installation/Stage/pl "Handbook:PPC/Installation/Stage/pl") must be extracted.

[← Configuring the network](/wiki/Handbook:PPC/Installation/Networking "Previous part") [Home](/wiki/Handbook:PPC "Handbook:PPC") [Installing the stage file →](/wiki/Handbook:PPC/Installation/Stage "Next part")