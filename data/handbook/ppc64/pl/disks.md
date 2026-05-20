# Disks

Other languages:

- [Deutsch](/wiki/Handbook:PPC64/Installation/Disks/de "Handbuch:PPC64/Installation/Festplatten (100% translated)")
- [English](/wiki/Handbook:PPC64/Installation/Disks "Handbook:PPC64/Installation/Disks (100% translated)")
- [español](/wiki/Handbook:PPC64/Installation/Disks/es "Manual de Gentoo: PPC64/Instalación/Discos (100% translated)")
- [français](/wiki/Handbook:PPC64/Installation/Disks/fr "Handbook:PPC64/Installation/Disks/fr (100% translated)")
- [italiano](/wiki/Handbook:PPC64/Installation/Disks/it "Handbook:PPC64/Installation/Disks/it (100% translated)")
- [magyar](/wiki/Handbook:PPC64/Installation/Disks/hu "Handbook:PPC64/Installation/Disks/hu (100% translated)")
- polski
- [português do Brasil](/wiki/Handbook:PPC64/Installation/Disks/pt-br "Handbook:PPC64/Installation/Disks/pt-br (50% translated)")
- [русский](/wiki/Handbook:PPC64/Installation/Disks/ru "Handbook:PPC64/Installation/Disks (100% translated)")
- [தமிழ்](/wiki/Handbook:PPC64/Installation/Disks/ta "கையேடு:PPC64/நிறுவல்/வட்டுக்கள் (100% translated)")
- [中文（中国大陆）‎](/wiki/Handbook:PPC64/Installation/Disks/zh-cn "手册：PPC64/安装/准备磁盘 (100% translated)")
- [日本語](/wiki/Handbook:PPC64/Installation/Disks/ja "ハンドブック:PPC64/インストール/ディスク (100% translated)")
- [한국어](/wiki/Handbook:PPC64/Installation/Disks/ko "Handbook:PPC64/Installation/Disks/ko (100% translated)")

[PPC64 Handbook](/wiki/Handbook:PPC64/pl "Handbook:PPC64/pl")[Instalacja](/wiki/Handbook:PPC64/Full/Installation/pl "Handbook:PPC64/Full/Installation/pl")[O instalacji](/wiki/Handbook:PPC64/Installation/About/pl "Handbook:PPC64/Installation/About/pl")[Wybór medium instalacyjnego](/wiki/Handbook:PPC64/Installation/Media/pl "Handbook:PPC64/Installation/Media/pl")[Konfiguracja sieci](/wiki/Handbook:PPC64/Installation/Networking/pl "Handbook:PPC64/Installation/Networking/pl")Przygotowanie dysków[Instalacja etapu 3](/wiki/Handbook:PPC64/Installation/Stage/pl "Handbook:PPC64/Installation/Stage/pl")[Instalacja systemu podstawowego](/wiki/Handbook:PPC64/Installation/Base/pl "Handbook:PPC64/Installation/Base/pl")[Konfiguracja jądra](/wiki/Handbook:PPC64/Installation/Kernel/pl "Handbook:PPC64/Installation/Kernel/pl")[Konfiguracja systemu](/wiki/Handbook:PPC64/Installation/System/pl "Handbook:PPC64/Installation/System/pl")[Instalacja narzędzi](/wiki/Handbook:PPC64/Installation/Tools/pl "Handbook:PPC64/Installation/Tools/pl")[Instalacja systemu rozruchowego](/wiki/Handbook:PPC64/Installation/Bootloader/pl "Handbook:PPC64/Installation/Bootloader/pl")[Finalizacja](/wiki/Handbook:PPC64/Installation/Finalizing/pl "Handbook:PPC64/Installation/Finalizing/pl")[Praca z Gentoo](/wiki/Handbook:PPC64/Full/Working/pl "Handbook:PPC64/Full/Working/pl")[Wstęp do Portage](/wiki/Handbook:PPC64/Working/Portage/pl "Handbook:PPC64/Working/Portage/pl")[Flagi USE](/wiki/Handbook:PPC64/Working/USE/pl "Handbook:PPC64/Working/USE/pl")[Funkcje portage](/wiki/Handbook:PPC64/Working/Features/pl "Handbook:PPC64/Working/Features/pl")[System initscript](/wiki/Handbook:PPC64/Working/Initscripts/pl "Handbook:PPC64/Working/Initscripts/pl")[Zmienne środowiskowe](/wiki/Handbook:PPC64/Working/EnvVar/pl "Handbook:PPC64/Working/EnvVar/pl")[Praca z Portage](/wiki/Handbook:PPC64/Full/Portage/pl "Handbook:PPC64/Full/Portage/pl")[Pliki i katalogi](/wiki/Handbook:PPC64/Portage/Files/pl "Handbook:PPC64/Portage/Files/pl")[Zmienne](/wiki/Handbook:PPC64/Portage/Variables/pl "Handbook:PPC64/Portage/Variables/pl")[Mieszanie działów oprogramowania](/wiki/Handbook:PPC64/Portage/Branches/pl "Handbook:PPC64/Portage/Branches/pl")[Dodatkowe narzędzia](/wiki/Handbook:PPC64/Portage/Tools/pl "Handbook:PPC64/Portage/Tools/pl")[Custom package repository](/wiki/Handbook:PPC64/Portage/CustomTree/pl "Handbook:PPC64/Portage/CustomTree/pl")[Funkcje zaawansowane](/wiki/Handbook:PPC64/Portage/Advanced/pl "Handbook:PPC64/Portage/Advanced/pl")[Konfiguracja sieci](/wiki/Handbook:PPC64/Full/Networking/pl "Handbook:PPC64/Full/Networking/pl")[Zaczynamy](/wiki/Handbook:PPC64/Networking/Introduction/pl "Handbook:PPC64/Networking/Introduction/pl")[Zaawansowana konfiguracja](/wiki/Handbook:PPC64/Networking/Advanced/pl "Handbook:PPC64/Networking/Advanced/pl")[Sieć modularna](/wiki/Handbook:PPC64/Networking/Modular/pl "Handbook:PPC64/Networking/Modular/pl")[Sieć bezprzewodowa](/wiki/Handbook:PPC64/Networking/Wireless/pl "Handbook:PPC64/Networking/Wireless/pl")[Dodawanie funkcjonalności](/wiki/Handbook:PPC64/Networking/Extending/pl "Handbook:PPC64/Networking/Extending/pl")[Dynamiczne zarządzanie](/wiki/Handbook:PPC64/Networking/Dynamic/pl "Handbook:PPC64/Networking/Dynamic/pl")

## Contents

- [1Introduction to block devices](#Introduction_to_block_devices)
  - [1.1Block devices](#Block_devices)
  - [1.2Partitions and slices](#Partitions_and_slices)
- [2Designing a partition scheme](#Designing_a_partition_scheme)
  - [2.1How many partitions and how big?](#How_many_partitions_and_how_big.3F)
  - [2.2What about swap space?](#What_about_swap_space.3F)
- [3Default: Using mac-fdisk](#Default:_Using_mac-fdisk)
- [4Alternative: Using fdisk](#Alternative:_Using_fdisk)
  - [4.1Viewing current partition layout](#Viewing_current_partition_layout)
  - [4.2Removing all partitions](#Removing_all_partitions)
  - [4.3Creating the PPC PReP boot partition](#Creating_the_PPC_PReP_boot_partition)
  - [4.4Creating the swap partition](#Creating_the_swap_partition)
  - [4.5Creating the root partition](#Creating_the_root_partition)
  - [4.6Saving the partition layout](#Saving_the_partition_layout)
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

### Partitions and slices

Although it is theoretically possible to use a full disk to house a Linux system, this is almost never done in practice. Instead, full disk block devices are split up in smaller, more manageable block devices. On most systems, these are called partitions. Other architectures use a similar technique, called _slices_.

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

## Default: Using mac-fdisk

**Ważne**

These instructions are for the Apple G5 system.

Partition
Description
/dev/sda1Apple partition map, automatically created when the disk is formatted with a "mac" partition table.
/dev/sda2New World boot block
/dev/sda3Swap partition
/dev/sda4Root partition

Start mac-fdisk:

`root #` `mac-fdisk /dev/sda`

First delete the partitions that have been cleared previously to make room for Linux partitions. Use the `d` key in mac-fdisk to delete those partition(s). It will ask for the partition number to delete.

Second, create an Apple\_Bootstrap partition by pressing the `b` key. It will ask for a block from which to start. Enter the number of the first free partition, followed by entering a `p`. For instance this is _2p_.

**Informacja**

This partition is not a "boot" partition. It is not used by Linux at all; there is no need to place any filesystem on it and it should never be mounted. PPC users do _not_ need an extra partition for /boot.

Now create a swap partition by pressing the `c` key. Again mac-fdisk will ask what block to start from. As we used 2 before to create the Apple\_Bootstrap partition, enter _3p_. When asked for the size, enter **512M** (or whatever size needed). When asked for a name, enter _swap_ (mandatory).

To create the root partition, enter `c`, followed by _4p_ to select from what block the root partition should start. When asked for the size, enter _4p_ again. mac-fdisk will interpret this as "Use all available space". When asked for the name, enter _root_ (mandatory).

To finish up, write the partition to the disk using `w` and `q` to quit mac-fdisk.

**Informacja**

To make sure everything is okay, run mac-fdisk once more to verify all the partitions are present. If a partition is absent, or it is missing some of the changes that were made, then reinitialize the partitions by pressing `i` in mac-fdisk. Note that this will recreate the partition map and thus remove all the partitions.

## Alternative: Using fdisk

**Ważne**

The following instructions are for IBM pSeries, iSeries, and OpenPower systems.

**Informacja**

When planning to use a RAID disk array for the Gentoo installation on POWER5-based hardware, first run iprconfig to format the disks to Advanced Function format and create the disk array. Emerge [sys-fs/iprutils](https://packages.gentoo.org/packages/sys-fs/iprutils) after the installation is complete.

If the system has an ipr-based SCSI adapter, start the ipr utilities now.

`root #` `/etc/init.d/iprinit start`

The following parts explain how to create the example partition layout described previously, namely:

Partition
Description
/dev/sda1PPC PReP Boot partition
/dev/sda2Swap partition
/dev/sda3Root partition

Change or modify the partition layout according to personal preference.

### Viewing current partition layout

[fdisk](/wiki/Fdisk "Fdisk") is a popular and powerful tool to split a disk into partitions. Fire up fdisk on the current disk (in our example, we use /dev/sda):

`root #` `fdisk /dev/sda`

```
Command (m for help)

```

If there is still an AIX partition layout on the system, then the following error message will be displayed:

`root #` `fdisk /dev/sda`

```
  There is a valid AIX label on this disk.
  Unfortunately Linux cannot handle these
  disks at the moment.  Nevertheless some
  advice:
  1. fdisk will destroy its contents on write.
  2. Be sure that this disk is NOT a still vital
     part of a volume group. (Otherwise you may
     erase the other disks as well, if unmirrored.)
  3. Before deleting this physical volume be sure
     to remove the disk logically from your AIX
     machine.  (Otherwise you become an AIXpert).

```

Don't worry, new empty DOS partition table can be created by pressing `o`.

**Uwaga**

This will destroy any installed AIX version!

Type `p` to display the disk current partition configuration:

`Command (m for help):` `p`

```
Disk /dev/sda: 30.7 GB, 30750031872 bytes
141 heads, 63 sectors/track, 6761 cylinders
Units = cylinders of 8883 * 512 = 4548096 bytes

   Device Boot      Start         End      Blocks   Id  System
/dev/sda1               1          12       53266+  83  Linux
/dev/sda2              13         233      981571+  82  Linux swap
/dev/sda3             234         674     1958701+  83  Linux
/dev/sda4             675        6761    27035410+   5  Extended
/dev/sda5             675        2874     9771268+  83  Linux
/dev/sda6            2875        2919      199836   83  Linux
/dev/sda7            2920        3008      395262   83  Linux
/dev/sda8            3009        6761    16668918   83  Linux

```

This particular disk is configured to house six Linux filesystems (each with a corresponding partition listed as "Linux") as well as a swap partition (listed as "Linux swap").

### Removing all partitions

First remove all existing partitions from the disk. Type `d` to delete a partition. For instance, to delete an existing /dev/sda1:

`Command (m for help):` `d`

```
Partition number (1-4): 1

```

The partition has been scheduled for deletion. It will no longer show up when typing `p`, but it will not be erased until the changes have been saved. If a mistake was made and the session needs to be aborted, then type `q` immediately and hit `Enter` and none of the partitions will be deleted or modified.

Now, assuming that indeed all partitions need to be wiped out, repeatedly type `p` to print out a partition listing and then type `d` and the number of the partition to delete it. Eventually, the partition table will show no more partitions:

`Command (m for help):` `p`

```
Disk /dev/sda: 30.7 GB, 30750031872 bytes
141 heads, 63 sectors/track, 6761 cylinders
Units = cylinders of 8883 * 512 = 4548096 bytes

Device Boot    Start       End    Blocks   Id  System

```

Now that the in-memory partition table is empty, let's create the partitions. We will use a default partitioning scheme as discussed previously. Of course, don't follow these instructions to the letter but adjust to personal preference.

### Creating the PPC PReP boot partition

First create a small PReP boot partition. Type `n` to create a new partition, then `p` to select a primary partition, followed by `1` to select the first primary partition. When prompted for the first cylinder, hit `Enter`. When prompted for the last cylinder, type +7M to create a partition 7 MB in size. After this, type `t` to set the partition type, `1` to select the partition just created and then type in _41_ to set the partition type to "PPC PReP Boot". Finally, mark the PReP partition as bootable.

**Informacja**

The PReP partition has to be smaller than 8 MB!

`Command (m for help):` `p`

```
Disk /dev/sda: 30.7 GB, 30750031872 bytes
141 heads, 63 sectors/track, 6761 cylinders
Units = cylinders of 8883 * 512 = 4548096 bytes

   Device Boot      Start         End      Blocks   Id  System

```

`Command (m for help):` `n`

```
Command action
      e   extended
      p   primary partition (1-4)
p
Partition number (1-4): 1
First cylinder (1-6761, default 1):
Using default value 1
Last cylinder or +size or +sizeM or +sizeK (1-6761, default
6761): +8M

```

`Command (m for help):` `t`

```
Selected partition 1
Hex code (type L to list codes): 41
Changed system type of partition 1 to 41 (PPC PReP Boot)

```

`Command (m for help):` `a`

```
Partition number (1-4): 1
Command (m for help):

```

Now, when looking at the partition table again (through `p`), the following partition information should be shown:

`Command (m for help):` `p`

```
Disk /dev/sda: 30.7 GB, 30750031872 bytes
141 heads, 63 sectors/track, 6761 cylinders
Units = cylinders of 8883 * 512 = 4548096 bytes

   Device Boot      Start         End      Blocks   Id  System
/dev/sda1  *            1           3       13293   41  PPC PReP Boot

```

### Creating the swap partition

Now create the swap partition. To do this, type `n` to create a new partition, then `p` to tell fdisk to create a primary partition. Then type `2` to create the second primary partition, /dev/sda2 in our case. When prompted for the first cylinder, hit `Enter`. When prompted for the last cylinder, type +512M to create a partition 512MB in size. After this, type `t` to set the partition type, `2` to select the partition just created and then type in _82_ to set the partition type to "Linux Swap". After completing these steps, typing `p` should display a partition table that looks similar to this:

`Command (m for help):` `p`

```
Disk /dev/sda: 30.7 GB, 30750031872 bytes
141 heads, 63 sectors/track, 6761 cylinders
Units = cylinders of 8883 * 512 = 4548096 bytes

   Device Boot      Start         End      Blocks   Id  System
/dev/sda1               1           3       13293   41  PPC PReP Boot
/dev/sda2               4         117      506331   82  Linux swap

```

### Creating the root partition

Finally, create the root partition. To do this, type `n` to create a new partition, then `p` to tell fdisk to create a primary partition. Then type `3` to create the third primary partition, /dev/sda3 in our case. When prompted for the first cylinder, hit `Enter`. When prompted for the last cylinder, hit enter to create a partition that takes up the rest of the remaining space on the disk. After completing these steps, typing `p` should display a partition table that looks similar to this:

`Command (m for help):` `p`

```
Disk /dev/sda: 30.7 GB, 30750031872 bytes
141 heads, 63 sectors/track, 6761 cylinders
Units = cylinders of 8883 * 512 = 4548096 bytes

   Device Boot      Start         End      Blocks   Id  System
/dev/sda1               1           3       13293   41  PPC PReP Boot
/dev/sda2               4         117      506331   82  Linux swap
/dev/sda3             118        6761    29509326   83  Linux

```

### Saving the partition layout

To save the partition layout and exit fdisk, type `w`.

`Command (m for help):` `w`

## Creating file systems

**Uwaga**

When using an SSD or NVMe drive, it is wise to check for firmware upgrades. Some Intel SSDs in particular (600p and 6000p) require a firmware upgrade for [possible data corruption](https://bugzilla.redhat.com/show_bug.cgi?id=1402533) induced by XFS I/O usage patterns. The problem is at the firmware level and not any fault of the XFS filesystem. The smartctl utility can help check the device model and firmware version.

### Introduction

Now that the partitions have been created, it is time to place a filesystem on them. In the next section the various file systems that Linux supports are described. Readers that already know which filesystem to use can continue with [Applying a filesystem to a partition](/wiki/Handbook:PPC64/Installation/Disks/pl#Applying_a_filesystem_to_a_partition "Handbook:PPC64/Installation/Disks/pl"). The others should read on to learn about the available filesystems...

### Filesystems

Linux supports several dozen filesystems, although many of them are only wise to deploy for specific purposes. Only certain filesystems may be found stable on the ppc64 architecture - it is advised to read up on the filesystems and their support state before selecting a more experimental one for important partitions. **XFS is the recommended all-purpose, all-platform filesystem.** The below is a non-exhaustive list:

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

Installations which were previously started, but did not finish the installation process can resume the installation from this point in the handbook. Use this link as the permalink: [Resumed installations start here](/wiki/Handbook:PPC64/Installation/Disks/pl#Resumed_installations_start_here "Handbook:PPC64/Installation/Disks/pl").

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

Later in the instructions, the proc filesystem (a virtual interface with the kernel) as well as other kernel pseudo-filesystems will be mounted. But first [the Gentoo stage file](/wiki/Handbook:PPC64/Installation/Stage/pl "Handbook:PPC64/Installation/Stage/pl") must be extracted.

[← Configuring the network](/wiki/Handbook:PPC64/Installation/Networking "Previous part") [Home](/wiki/Handbook:PPC64 "Handbook:PPC64") [Installing the stage file →](/wiki/Handbook:PPC64/Installation/Stage "Next part")