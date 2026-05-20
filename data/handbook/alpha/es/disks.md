# Disks

Other languages:

- [Deutsch](/wiki/Handbook:Alpha/Installation/Disks/de "Handbuch:Alpha/Installation/Festplatten (100% translated)")
- [English](/wiki/Handbook:Alpha/Installation/Disks "Handbook:Alpha/Installation/Disks (100% translated)")
- [Türkçe](/wiki/Handbook:Alpha/Installation/Disks/tr "Handbook:Alpha/Installation/Disks/tr (0% translated)")
- español
- [français](/wiki/Handbook:Alpha/Installation/Disks/fr "Manuel:Alpha/Installation/Disques (100% translated)")
- [italiano](/wiki/Handbook:Alpha/Installation/Disks/it "Handbook:Alpha/Installation/Disks/it (100% translated)")
- [magyar](/wiki/Handbook:Alpha/Installation/Disks/hu "Handbook:Alpha/Installation/Disks/hu (100% translated)")
- [polski](/wiki/Handbook:Alpha/Installation/Disks/pl "Handbook:Alpha/Installation/Disks (100% translated)")
- [português do Brasil](/wiki/Handbook:Alpha/Installation/Disks/pt-br "Manual:Alpha/Instalação/Discos (100% translated)")
- [čeština](/wiki/Handbook:Alpha/Installation/Disks/cs "Handbook:Alpha/Installation/Disks/cs (50% translated)")
- [русский](/wiki/Handbook:Alpha/Installation/Disks/ru "Handbook:Alpha/Installation/Disks (100% translated)")
- [தமிழ்](/wiki/Handbook:Alpha/Installation/Disks/ta "கையேடு:Alpha/நிறுவல்/வட்டுக்கள் (100% translated)")
- [中文（中国大陆）‎](/wiki/Handbook:Alpha/Installation/Disks/zh-cn "手册：Alpha/安装/准备磁盘 (100% translated)")
- [日本語](/wiki/Handbook:Alpha/Installation/Disks/ja "ハンドブック:Alpha/インストール/ディスク (100% translated)")
- [한국어](/wiki/Handbook:Alpha/Installation/Disks/ko "Handbook:Alpha/Installation/Disks/ko (100% translated)")

[Alpha Handbook](/wiki/Handbook:Alpha "Handbook:Alpha")[Installation](/wiki/Handbook:Alpha/Full/Installation "Handbook:Alpha/Full/Installation")[About the installation](/wiki/Handbook:Alpha/Installation/About/es "Handbook:Alpha/Installation/About/es")[Choosing the media](/wiki/Handbook:Alpha/Installation/Media/es "Handbook:Alpha/Installation/Media/es")[Configuring the network](/wiki/Handbook:Alpha/Installation/Networking/es "Handbook:Alpha/Installation/Networking/es")Preparing the disks[The stage file](/wiki/Handbook:Alpha/Installation/Stage/es "Handbook:Alpha/Installation/Stage/es")[Installing base system](/wiki/Handbook:Alpha/Installation/Base/es "Handbook:Alpha/Installation/Base/es")[Configuring the kernel](/wiki/Handbook:Alpha/Installation/Kernel/es "Handbook:Alpha/Installation/Kernel/es")[Configuring the system](/wiki/Handbook:Alpha/Installation/System/es "Handbook:Alpha/Installation/System/es")[Installing tools](/wiki/Handbook:Alpha/Installation/Tools/es "Handbook:Alpha/Installation/Tools/es")[Configuring the bootloader](/wiki/Handbook:Alpha/Installation/Bootloader/es "Handbook:Alpha/Installation/Bootloader/es")[Finalizing](/wiki/Handbook:Alpha/Installation/Finalizing/es "Handbook:Alpha/Installation/Finalizing/es")[Working with Gentoo](/wiki/Handbook:Alpha/Full/Working "Handbook:Alpha/Full/Working")[Portage introduction](/wiki/Handbook:Alpha/Working/Portage/es "Handbook:Alpha/Working/Portage/es")[USE flags](/wiki/Handbook:Alpha/Working/USE/es "Handbook:Alpha/Working/USE/es")[Portage features](/wiki/Handbook:Alpha/Working/Features/es "Handbook:Alpha/Working/Features/es")[Initscript system](/wiki/Handbook:Alpha/Working/Initscripts/es "Handbook:Alpha/Working/Initscripts/es")[Environment variables](/wiki/Handbook:Alpha/Working/EnvVar/es "Handbook:Alpha/Working/EnvVar/es")[Working with Portage](/wiki/Handbook:Alpha/Full/Portage "Handbook:Alpha/Full/Portage")[Files and directories](/wiki/Handbook:Alpha/Portage/Files/es "Handbook:Alpha/Portage/Files/es")[Variables](/wiki/Handbook:Alpha/Portage/Variables/es "Handbook:Alpha/Portage/Variables/es")[Mixing software branches](/wiki/Handbook:Alpha/Portage/Branches/es "Handbook:Alpha/Portage/Branches/es")[Additional tools](/wiki/Handbook:Alpha/Portage/Tools/es "Handbook:Alpha/Portage/Tools/es")[Custom package repository](/wiki/Handbook:Alpha/Portage/CustomTree/es "Handbook:Alpha/Portage/CustomTree/es")[Advanced features](/wiki/Handbook:Alpha/Portage/Advanced/es "Handbook:Alpha/Portage/Advanced/es")[OpenRC network configuration](/wiki/Handbook:Alpha/Full/Networking "Handbook:Alpha/Full/Networking")[Getting started](/wiki/Handbook:Alpha/Networking/Introduction/es "Handbook:Alpha/Networking/Introduction/es")[Advanced configuration](/wiki/Handbook:Alpha/Networking/Advanced/es "Handbook:Alpha/Networking/Advanced/es")[Modular networking](/wiki/Handbook:Alpha/Networking/Modular/es "Handbook:Alpha/Networking/Modular/es")[Wireless](/wiki/Handbook:Alpha/Networking/Wireless/es "Handbook:Alpha/Networking/Wireless/es")[Adding functionality](/wiki/Handbook:Alpha/Networking/Extending/es "Handbook:Alpha/Networking/Extending/es")[Dynamic management](/wiki/Handbook:Alpha/Networking/Dynamic/es "Handbook:Alpha/Networking/Dynamic/es")

## Contents

- [1Introduction to block devices](#Introduction_to_block_devices)
  - [1.1Block devices](#Block_devices)
  - [1.2Slices](#Slices)
- [2Designing a partition scheme](#Designing_a_partition_scheme)
  - [2.1How many partitions and how big?](#How_many_partitions_and_how_big.3F)
  - [2.2What about swap space?](#What_about_swap_space.3F)
- [3Using fdisk to partition a disk (SRM only)](#Using_fdisk_to_partition_a_disk_.28SRM_only.29)
  - [3.1Identifying available disks](#Identifying_available_disks)
  - [3.2Creating a BSD disklabel](#Creating_a_BSD_disklabel)
  - [3.3Deleting all slices](#Deleting_all_slices)
  - [3.4Creating the swap slice](#Creating_the_swap_slice)
  - [3.5Creating the boot slice](#Creating_the_boot_slice)
  - [3.6Creating the root slice](#Creating_the_root_slice)
  - [3.7Save the slice layout and exit](#Save_the_slice_layout_and_exit)
- [4Using fdisk to partition the disk (ARC/AlphaBIOS only)](#Using_fdisk_to_partition_the_disk_.28ARC.2FAlphaBIOS_only.29)
  - [4.1Identifying the available disks](#Identifying_the_available_disks)
  - [4.2Deleting all partitions](#Deleting_all_partitions)
  - [4.3Creating the boot partition](#Creating_the_boot_partition)
  - [4.4Creating the swap partition](#Creating_the_swap_partition)
  - [4.5Creating the root partition](#Creating_the_root_partition)
  - [4.6Save the partition layout and exit](#Save_the_partition_layout_and_exit)
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

### Slices

Although it is theoretically possible to use a full disk to house a Linux system, this is almost never done in practice. Instead, full disk block devices are split up in smaller, more manageable block devices. On Alpha systems, these are called _slices_.

**Nota**

In further sections, the installation instructions will use the example partitioning for the ARC/AlphaBIOS setup. Please adjust to personal preference!

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

**Nota**

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

**Consejo**

With the rise in popularity with [zram](/wiki/Zram "Zram"), it should be noted that Gentoo is more likely to suffer with performance hits, as compiling uses large amounts of RAM, and programs can't be directly compiled in swap. It's usually better to create a smaller zram as a cache swap area with higher priority over the disk based swap of around 512MB, if a user really wants to use zram to benefit their desktop needs and not harm compile times.

## Using fdisk to partition a disk (SRM only)

The following parts explain how to create the example slice layout for the SRM:

Slice
Description
/dev/sda1Swap slice
/dev/sda2Root slice
/dev/sda3Full disk (required)

Change the slice layout according to personal preference.

### Identifying available disks

To figure out what disks are running in the system, use the following commands:

For IDE disks:

`root #` `dmesg | grep 'drive$'`

For SCSI disks:

`root #` `dmesg | grep 'scsi'`

The output will show what disks were detected and their respective /dev/ entry. In the following parts we assume that the disk is a SCSI disk on /dev/sda.

### Creating a BSD disklabel

If the hard drive is completely blank, then first create a BSD disklabel. On alpha, this can't be done using fdisk, so using parted is the way forward here.

Now fire up parted:

`root #` `parted /dev/sda `

```
Using /dev/sda
Welcome to GNU Parted! Type 'help' to view a list of commands.

```

`(parted)` `mklabel bsd`

```
Warning: The existing disk label on /dev/sda will be destroyed and all data on this disk will be lost. Do you want to continue?
Yes/No? yes
(parted) quit
Information: You may need to update /etc/fstab.

```

Now that we have a bsd disklabel on our drive, continue creating slices. This can be done using parted or, as in the examples below, using fdisk:

`root #` `fdisk /dev/sda`

### Deleting all slices

If the hard drive is completely blank, then first create a BSD disklabel.

`Command (m for help):` `b`

```
/dev/sda contains no disklabel.
Do you want to create a disklabel? (y/n) y
A bunch of drive-specific info will show here
3 partitions:
#       start       end      size     fstype   [fsize bsize   cpg]
  c:        1      5290*     5289*    unused        0     0

```

We start with deleting all slices except the 'c'-slice (a requirement for using BSD disklabels). The following shows how to delete a slice (in the example we use 'a'). Repeat the process to delete all other slices (again, except the 'c'-slice).

Use `p` to view all existing slices. `d` is used to delete a slice.

`BSD disklabel command (m for help):` `p`

```
8 partitions:
#       start       end      size     fstype   [fsize bsize   cpg]
  a:        1       235*      234*    4.2BSD     1024  8192    16
  b:      235*      469*      234*      swap
  c:        1      5290*     5289*    unused        0     0
  d:      469*     2076*     1607*    unused        0     0
  e:     2076*     3683*     1607*    unused        0     0
  f:     3683*     5290*     1607*    unused        0     0
  g:      469*     1749*     1280     4.2BSD     1024  8192    16
  h:     1749*     5290*     3541*    unused        0     0

```

`BSD disklabel command (m for help):` `d`

```
Partition (a-h): a

```

After repeating this process for all slices, a listing should show something similar to this:

`BSD disklabel command (m for help):` `p`

```
3 partitions:
#       start       end      size     fstype   [fsize bsize   cpg]
  c:        1      5290*     5289*    unused        0     0

```

### Creating the swap slice

On Alpha based systems there is no need for a separate boot slice, but in this example we will create a separate /boot slice to hold the kernels and boot loader configuration files also, the first cylinder cannot be used as the _aboot_ image will be placed there.

We will create a swap slice starting at the third cylinder, with a total size of 1 GB. Use `n` to create a new slice. After creating the slice, we will change its type to `1` (one), meaning swap.

`BSD disklabel command (m for help):` `n`

```
Partition (a-p): a
First cylinder (1-5290, default 1): 3
Last cylinder or +size or +sizeM or +sizeK (3-5290, default 5290): +1024M

```

`BSD disklabel command (m for help):` `t`

```
Partition (a-c): a
Hex code (type L to list codes): 1

```

After these steps a layout similar to the following should be shown:

`BSD disklabel command (m for help):` `p`

```
3 partitions:
#       start       end      size     fstype   [fsize bsize   cpg]
  a:        3      1003      1001       swap
  c:        1      5290*     5289*    unused        0     0

```

### Creating the boot slice

Create a boot file system containing the kernels and boot loader configuration file (aboot.conf). aboot supports ext2 and ext3 filesystems only. Create the boot slice starting from the first cylinder after the swap slice. Use the `p` command to view where the swap slice ends. In our example, this is at 1003, making the boot slice start at 1004.

`BSD disklabel command (m for help):` `n`

```
Partition (a-p): b
First cylinder (1-5290, default 1): 1004
Last cylinder or +size or +sizeM or +sizeK (3-5290, default 5290): +1024M

```

`BSD disklabel command (m for help):` `t`

```
Partition (a-c): b
Hex code (type L to list codes): 08

```

After these steps a layout similar to the following should be shown:

`BSD disklabel command (m for help):` `p`

```
3 partitions:
#       start       end      size     fstype   [fsize bsize   cpg]
  a:        3      1003      1001       swap
  b:     1004      2005      1001       ext2
  c:        1      5290*     5289*    unused        0     0

```

### Creating the root slice

We will now create the root slice, starting from the first cylinder after the boot slice. Use the `p` command to view where the boot slice ends. In our example, this is at 2005, making the root slice start at 2006.

Another problem is that there is currently a bug in fdisk making it think the number of available cylinders is one above the real number of cylinders. In other words, when asked for the last cylinder, decrease the cylinder number (in this example: 5290) with one.

When the slice is created, we change the type to 8, for ext2.

`BSD disklabel command (m for help):` `n`

```
Partition (a-p): b
First cylinder (1-5290, default 1): 2006
Last cylinder or +size or +sizeM or +sizeK (1004-5290, default 5290): 5289

```

`BSD disklabel command (m for help):` `t`

```
Partition (a-c): b
Hex code (type L to list codes): 8

```

The resulting slice layout should now be similar to this:

`BSD disklabel command (m for help):` `p`

```
3 partitions:
#       start       end      size     fstype   [fsize bsize   cpg]
  a:        3      1003      1001       swap
  b:     1004      2005      4286       ext2
  c:        1      5290*     5289*    unused        0     0
  d:     2006      5289      3283       ext2

```

### Save the slice layout and exit

Exit the fdisk application by typing `w`. This will also save the slice layout.

`Command (m for help):` `w`

## Using fdisk to partition the disk (ARC/AlphaBIOS only)

The following parts explain how to create the example partition layout for ARC/AlphaBIOS:

Partition
Description
/dev/sda1Boot partition
/dev/sda2Swap partition
/dev/sda3Root partition

Change the partition layout according to personal preference.

### Identifying the available disks

To figure out what disks are running, use the following commands:

For IDE disks:

`root #` `dmesg | grep 'drive$'`

For SCSI disks:

`root #` `dmesg | grep 'scsi'`

From this output it should be easy to see what disks were detected and their respective /dev/ entry. In the following parts we assume that the disk is a SCSI disk on /dev/sda.

Now fire up fdisk:

`root #` `fdisk /dev/sda`

### Deleting all partitions

If the hard drive is completely blank, then first create a DOS disklabel.

`Command (m for help):` `o`

```
Building a new DOS disklabel.

```

We start with deleting all partitions. The following shows how to delete a partition (in the example we use '1'). Repeat the process to delete all other partitions.

Use `p` to view all existing partitions. `d` is used to delete a partition.

`command (m for help):` `p`

```
Disk /dev/sda: 9150 MB, 9150996480 bytes
64 heads, 32 sectors/track, 8727 cylinders
Units = cylinders of 2048 * 512 = 1048576 bytes

   Device Boot      Start         End      Blocks   Id  System
/dev/sda1               1         478      489456   83  Linux
/dev/sda2             479        8727     8446976    5  Extended
/dev/sda5             479        1433      977904   83  Linux Swap
/dev/sda6            1434        8727     7469040   83  Linux

```

`command (m for help):` `d`

```
Partition number (1-6): 1

```

### Creating the boot partition

On Alpha systems which use MILO to boot, we have to create a small vfat boot partition.

`Command (m for help):` `n`

```
Command action
  e   extended
  p   primary partition (1-4)
p
Partition number (1-4): 1
First cylinder (1-8727, default 1): 1
Last cylinder or +size or +sizeM or +sizeK (1-8727, default 8727): +16M

```

`Command (m for help):` `t`

```
Selected partition 1
Hex code (type L to list codes): 6
Changed system type of partition 1 to 6 (FAT16)

```

### Creating the swap partition

We will create a swap partition with a total size of 1 GB. Use `n` to create a new partition.

`Command (m for help):` `n`

```
Command action
  e   extended
  p   primary partition (1-4)
p
Partition number (1-4): 2
First cylinder (17-8727, default 17): 17
Last cylinder or +size or +sizeM or +sizeK (17-8727, default 8727): +1000M

```

`Command (m for help):` `t`

```
Partition number (1-4): 2
Hex code (type L to list codes): 82
Changed system type of partition 2 to 82 (Linux swap)

```

After these steps a layout similar to the following is shown:

`Command (m for help):` `p`

```
Disk /dev/sda: 9150 MB, 9150996480 bytes
64 heads, 32 sectors/track, 8727 cylinders
Units = cylinders of 2048 * 512 = 1048576 bytes

   Device Boot      Start         End      Blocks   Id  System
/dev/sda1               1          16       16368    6  FAT16
/dev/sda2              17         971      977920   82  Linux swap

```

### Creating the root partition

We will now create the root partition. Again, just use the `n` command.

`Command (m for help):` `n`

```
Command action
  e   extended
  p   primary partition (1-4)
p
Partition number (1-4): 3
First cylinder (972-8727, default 972): 972
Last cylinder or +size or +sizeM or +sizeK (972-8727, default 8727): 8727

```

After these steps a layout similar to the following should be shown:

`Command (m for help):` `p`

```
Disk /dev/sda: 9150 MB, 9150996480 bytes
64 heads, 32 sectors/track, 8727 cylinders
Units = cylinders of 2048 * 512 = 1048576 bytes

   Device Boot      Start         End      Blocks   Id  System
/dev/sda1               1          16       16368    6  FAT16
/dev/sda2              17         971      977920   82  Linux swap
/dev/sda3             972        8727     7942144   83  Linux

```

### Save the partition layout and exit

Save the changes made in fdisk by typing `w`.

`Command (m for help):` `w`

Now that the partitions are created, continue with [Creating filesystems](#Creating_filesystems).

## Creating file systems

**Advertencia**

When using an SSD or NVMe drive, it is wise to check for firmware upgrades. Some Intel SSDs in particular (600p and 6000p) require a firmware upgrade for [possible data corruption](https://bugzilla.redhat.com/show_bug.cgi?id=1402533) induced by XFS I/O usage patterns. The problem is at the firmware level and not any fault of the XFS filesystem. The smartctl utility can help check the device model and firmware version.

### Introduction

Now that the partitions have been created, it is time to place a filesystem on them. In the next section the various file systems that Linux supports are described. Readers that already know which filesystem to use can continue with [Applying a filesystem to a partition](/wiki/Handbook:Alpha/Installation/Disks/es#Applying_a_filesystem_to_a_partition "Handbook:Alpha/Installation/Disks/es"). The others should read on to learn about the available filesystems...

### Filesystems

Linux supports several dozen filesystems, although many of them are only wise to deploy for specific purposes. Only certain filesystems may be found stable on the alpha architecture - it is advised to read up on the filesystems and their support state before selecting a more experimental one for important partitions. **XFS is the recommended all-purpose, all-platform filesystem.** The below is a non-exhaustive list:

[XFS](/wiki/XFS "XFS")Filesystem with metadata journaling which comes with a robust feature-set and is optimized for scalability. It has been continuously upgraded to include modern features. The only downside is that XFS partitions cannot yet be shrunk, although this is being worked on. XFS notably supports reflinks and Copy on Write (CoW) which is particularly helpful on Gentoo systems because of the amount of compiles users complete. XFS is the recommended modern all-purpose all-platform filesystem. Requires a partition to be at least 300MB.[ext4](/wiki/Ext4 "Ext4")Ext4 is a reliable, all-purpose all-platform filesystem, although it lacks modern features like reflinks.[VFAT](/wiki/VFAT "VFAT")Also known as FAT32, is supported by Linux but does not support standard UNIX permission settings. It is mostly used for interoperability/interchange with other operating systems (Microsoft Windows or Apple's macOS) but is also a necessity for some system bootloader firmware (like UEFI). Users of UEFI systems will need an [EFI System Partition](/wiki/EFI_System_Partition "EFI System Partition") formatted with VFAT in order to boot.[btrfs](/wiki/Btrfs/es "Btrfs/es")Newer generation filesystem. Provides advanced features like snapshotting, self-healing through checksums, transparent compression, subvolumes, and integrated RAID. Kernels prior to 5.4.y are not guaranteed to be safe to use with btrfs in production because fixes for serious issues are only present in the more recent releases of the LTS kernel branches. RAID 5/6 and quota groups unsafe on all versions of btrfs.[F2FS](/wiki/F2FS "F2FS")The Flash-Friendly File System was originally created by Samsung for the use with NAND flash memory. It is a decent choice when installing Gentoo to microSD cards, USB drives, or other flash-based storage devices.[NTFS](/wiki/NTFS "NTFS")This "New Technology" filesystem is the flagship filesystem of Microsoft Windows since Windows NT 3.1. Similarly to VFAT, it does not store UNIX permission settings or extended attributes necessary for BSD or Linux to function properly, therefore it should not be used as a root filesystem for most cases. It should _only_ be used for interoperability or data interchange with Microsoft Windows systems (note the emphasis on _only_).[ZFS](/wiki/ZFS "ZFS")**Importante:** ZFS pools can be created on the minimal installation CD, for further information, refer to [ZFS/rootfs](/wiki/ZFS/rootfs "ZFS/rootfs")Next generation file system created by Matthew Ahrens and Jeff Bonwick. It was designed around a few key ideas: Administration of storage should be simple, redundancy should be handled by the filesystem, file systems should never be taken offline for repair, automated simulations of worst case scenarios before shipping code is important, and data integrity is paramount.

More extensive information on filesystems can be found in the community maintained [Filesystem article](/wiki/Filesystem/es "Filesystem/es").

### Applying a filesystem to a partition

**Nota**

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
[sys-fs/dosfstools](https://packages.gentoo.org/packages/sys-fs/dosfstools)[btrfs](/wiki/Btrfs/es "Btrfs/es")mkfs.btrfs Yes
[sys-fs/btrfs-progs](https://packages.gentoo.org/packages/sys-fs/btrfs-progs)[F2FS](/wiki/F2FS "F2FS")mkfs.f2fs Yes
[sys-fs/f2fs-tools](https://packages.gentoo.org/packages/sys-fs/f2fs-tools)[NTFS](/wiki/NTFS "NTFS")mkfs.ntfs Yes
[sys-fs/ntfs3g](https://packages.gentoo.org/packages/sys-fs/ntfs3g)[ZFS](/wiki/ZFS "ZFS")zpool create ... Yes
[sys-fs/zfs](https://packages.gentoo.org/packages/sys-fs/zfs)

**Importante**

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

**Nota**

Installations which were previously started, but did not finish the installation process can resume the installation from this point in the handbook. Use this link as the permalink: [Resumed installations start here](/wiki/Handbook:Alpha/Installation/Disks/es#Resumed_installations_start_here "Handbook:Alpha/Installation/Disks/es").

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

**Nota**

If /tmp/ needs to reside on a separate partition, be sure to change its permissions after mounting:

`root #` `chmod 1777 /mnt/gentoo/tmp`

This also holds for /var/tmp.

Later in the instructions, the proc filesystem (a virtual interface with the kernel) as well as other kernel pseudo-filesystems will be mounted. But first [the Gentoo stage file](/wiki/Handbook:Alpha/Installation/Stage/es "Handbook:Alpha/Installation/Stage/es") must be extracted.

[← Configuring the network](/wiki/Handbook:Alpha/Installation/Networking "Previous part") [Home](/wiki/Handbook:Alpha "Handbook:Alpha") [Installing the stage file →](/wiki/Handbook:Alpha/Installation/Stage "Next part")