# Disks

Other languages:

- [Deutsch](/wiki/Handbook:HPPA/Installation/Disks/de "Handbuch:HPPA/Installation/Festplatten (100% translated)")
- English
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

[HPPA Handbook](/wiki/Handbook:HPPA "Handbook:HPPA")[Installation](/wiki/Handbook:HPPA/Full/Installation "Handbook:HPPA/Full/Installation")[About the installation](/wiki/Handbook:HPPA/Installation/About "Handbook:HPPA/Installation/About")[Choosing the media](/wiki/Handbook:HPPA/Installation/Media "Handbook:HPPA/Installation/Media")[Configuring the network](/wiki/Handbook:HPPA/Installation/Networking "Handbook:HPPA/Installation/Networking")Preparing the disks[The stage file](/wiki/Handbook:HPPA/Installation/Stage "Handbook:HPPA/Installation/Stage")[Installing base system](/wiki/Handbook:HPPA/Installation/Base "Handbook:HPPA/Installation/Base")[Configuring the kernel](/wiki/Handbook:HPPA/Installation/Kernel "Handbook:HPPA/Installation/Kernel")[Configuring the system](/wiki/Handbook:HPPA/Installation/System "Handbook:HPPA/Installation/System")[Installing tools](/wiki/Handbook:HPPA/Installation/Tools "Handbook:HPPA/Installation/Tools")[Configuring the bootloader](/wiki/Handbook:HPPA/Installation/Bootloader "Handbook:HPPA/Installation/Bootloader")[Finalizing](/wiki/Handbook:HPPA/Installation/Finalizing "Handbook:HPPA/Installation/Finalizing")[Working with Gentoo](/wiki/Handbook:HPPA/Full/Working "Handbook:HPPA/Full/Working")[Portage introduction](/wiki/Handbook:HPPA/Working/Portage "Handbook:HPPA/Working/Portage")[USE flags](/wiki/Handbook:HPPA/Working/USE "Handbook:HPPA/Working/USE")[Portage features](/wiki/Handbook:HPPA/Working/Features "Handbook:HPPA/Working/Features")[Initscript system](/wiki/Handbook:HPPA/Working/Initscripts "Handbook:HPPA/Working/Initscripts")[Environment variables](/wiki/Handbook:HPPA/Working/EnvVar "Handbook:HPPA/Working/EnvVar")[Working with Portage](/wiki/Handbook:HPPA/Full/Portage "Handbook:HPPA/Full/Portage")[Files and directories](/wiki/Handbook:HPPA/Portage/Files "Handbook:HPPA/Portage/Files")[Variables](/wiki/Handbook:HPPA/Portage/Variables "Handbook:HPPA/Portage/Variables")[Mixing software branches](/wiki/Handbook:HPPA/Portage/Branches "Handbook:HPPA/Portage/Branches")[Additional tools](/wiki/Handbook:HPPA/Portage/Tools "Handbook:HPPA/Portage/Tools")[Custom package repository](/wiki/Handbook:HPPA/Portage/CustomTree "Handbook:HPPA/Portage/CustomTree")[Advanced features](/wiki/Handbook:HPPA/Portage/Advanced "Handbook:HPPA/Portage/Advanced")[OpenRC network configuration](/wiki/Handbook:HPPA/Full/Networking "Handbook:HPPA/Full/Networking")[Getting started](/wiki/Handbook:HPPA/Networking/Introduction "Handbook:HPPA/Networking/Introduction")[Advanced configuration](/wiki/Handbook:HPPA/Networking/Advanced "Handbook:HPPA/Networking/Advanced")[Modular networking](/wiki/Handbook:HPPA/Networking/Modular "Handbook:HPPA/Networking/Modular")[Wireless](/wiki/Handbook:HPPA/Networking/Wireless "Handbook:HPPA/Networking/Wireless")[Adding functionality](/wiki/Handbook:HPPA/Networking/Extending "Handbook:HPPA/Networking/Extending")[Dynamic management](/wiki/Handbook:HPPA/Networking/Dynamic "Handbook:HPPA/Networking/Dynamic")

## Contents

- [1Introduction to block devices](#Introduction_to_block_devices)
  - [1.1Block devices](#Block_devices)
  - [1.2Partitions and slices](#Partitions_and_slices)
- [2Designing a partition scheme](#Designing_a_partition_scheme)
  - [2.1How many partitions and how big?](#How_many_partitions_and_how_big.3F)
  - [2.2What about swap space?](#What_about_swap_space.3F)
- [3Using fdisk on HPPA](#Using_fdisk_on_HPPA)
- [4Creating file systems](#Creating_file_systems)
  - [4.1Introduction](#Introduction)
  - [4.2Filesystems](#Filesystems)
  - [4.3Applying a filesystem to a partition](#Applying_a_filesystem_to_a_partition)
    - [4.3.1Legacy BIOS boot partition filesystem](#Legacy_BIOS_boot_partition_filesystem)
    - [4.3.2Small ext4 partitions](#Small_ext4_partitions)
  - [4.4Activating the swap partition](#Activating_the_swap_partition)
- [5Mounting the root partition](#Mounting_the_root_partition)

## Introduction to block devices\[ [edit](/index.php?title=Handbook:Parts/Installation/Disks&action=edit&section=T-1 "Edit section: ")\]

### Block devices\[ [edit](/index.php?title=Handbook:Parts/Installation/Disks&action=edit&section=T-2 "Edit section: ")\]

Let's take a good look at disk-oriented aspects of Gentoo Linux and Linux in general, including block devices, partitions, and Linux filesystems. Once the ins and outs of disks are understood, partitions and filesystems can be established for installation.

To begin, let's look at block devices. SCSI and Serial ATA drives are both labeled under device handles such as: /dev/sda, /dev/sdb, /dev/sdc, etc. On more modern machines, PCI Express based NVMe solid state disks have device handles such as /dev/nvme0n1, /dev/nvme0n2, etc.

The following table will help readers determine where to find a certain type of block device on the system:

Type of deviceDefault device handleEditorial notes and considerations
IDE, SATA, SAS, SCSI, or USB flash/dev/sdaFound on hardware from roughly 2007 until the present, this device handle is perhaps the most commonly used in Linux. These types of devices can be connected via the [SATA bus](https://en.wikipedia.org/wiki/Serial_ATA "wikipedia:Serial ATA"), [SCSI](https://en.wikipedia.org/wiki/SCSI "wikipedia:SCSI"), [USB](https://en.wikipedia.org/wiki/USB "wikipedia:USB") bus as block storage. As example, the first partition on the first SATA device is called /dev/sda1.
NVM Express (NVMe)/dev/nvme0n1The latest in solid state technology, [NVMe](https://en.wikipedia.org/wiki/NVM_Express "wikipedia:NVM Express") drives are connected to the PCI Express bus and have the fastest transfer block speeds on the market. Systems from around 2014 and newer may have support for NVMe hardware. The first partition on the first NVMe device is called /dev/nvme0n1p1.
MMC, eMMC, and SD/dev/mmcblk0[embedded MMC](https://en.wikipedia.org/wiki/MultiMediaCard#eMMC "wikipedia:MultiMediaCard") devices, SD cards, and other types of memory cards can be useful for data storage. That said, many systems may not permit booting from these types of devices. It is suggested to **not** use these devices for active Linux installations; rather consider using them to transfer files, which is their typical design intention. Alternatively this storage type could be useful for short-term file backups or snapshots.
VirtIO/dev/vda[Virtualized](/wiki/Virtualization "Virtualization") devices are found only within a [QEMU](/wiki/QEMU "QEMU") virtual emulator. `virtio_blk` driver provides this access to host's allocated drive space for this virtual machine. That said, it is a great way to try out Gentoo inside a virtual machine.

The block devices above represent an abstract interface to the disk. User programs can use these block devices to interact with the disk without worrying about whether the drives are SATA, SCSI, or something else. The program can simply address the storage on the disk as a bunch of contiguous, randomly-accessible 4096-byte (4K) blocks.

### Partitions and slices\[ [edit](/index.php?title=Handbook:HPPA/Blocks/Disks&action=edit&section=T-1 "Edit section: ")\]

Although it is theoretically possible to use a full disk to house the Linux system, this is almost never done in practice. Instead, full disk block devices are split up in smaller, more manageable block devices. On most systems, these are called partitions. Other architectures use a similar technique, called _slices_.

## Designing a partition scheme\[ [edit](/index.php?title=Handbook:Parts/Blocks/DesigningPartitionScheme&action=edit&section=T-1 "Edit section: ")\]

### How many partitions and how big?\[ [edit](/index.php?title=Handbook:Parts/Blocks/DesigningPartitionScheme&action=edit&section=T-2 "Edit section: ")\]

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

**Note**

Installations that intend to use systemd as the service and init system must have the /usr directory available at boot, either as part of the root filesystem or mounted via an initramfs.

### What about swap space?\[ [edit](/index.php?title=Handbook:Parts/Blocks/DesigningPartitionScheme&action=edit&section=T-3 "Edit section: ")\]

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

**Tip**

With the rise in popularity with [zram](/wiki/Zram "Zram"), it should be noted that Gentoo is more likely to suffer with performance hits, as compiling uses large amounts of RAM, and programs can't be directly compiled in swap. It's usually better to create a smaller zram as a cache swap area with higher priority over the disk based swap of around 512MB, if a user really wants to use zram to benefit their desktop needs and not harm compile times.

## Using fdisk on HPPA\[ [edit](/index.php?title=Handbook:HPPA/Blocks/Disks&action=edit&section=T-2 "Edit section: ")\]

Use fdisk to create the partitions needed:

`root #` `fdisk /dev/sda`

HPPA machines use the PC standard DOS partition tables. To create a new DOS partition table press the `o` key.

`Command (m for help):` `o`

```
Building a new DOS disklabel.

```

PALO (the HPPA bootloader) needs a special partition to work. A partition of at least 16 MB at the beginning of the disk needs to be created for it. The partition type must be of type _f0 (Linux/PA-RISC boot)_. It is also possible to use the PALO partition as /boot.

**Important**

If this is forgotten and the installation continues without a special PALO partition, then eventually the system will fail to restart. Also, if the disk is larger than 2 GB, make sure that the boot partition is in the first 2 GB of the disk. PALO is unable to read a kernel after the 2 GB limit.

FILE **`/etc/fstab`** **A simple default partition scheme**

```
/dev/sda2    /boot   ext2    noauto,noatime   1 1
/dev/sda3    none    swap    sw               0 0
/dev/sda4    /       ext4    noatime          0 0

```

In fdisk, such a partition layout looks like so:

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

## Creating file systems\[ [edit](/index.php?title=Handbook:Parts/Installation/Disks&action=edit&section=T-3 "Edit section: ")\]

**Warning**

When using an SSD or NVMe drive, it is wise to check for firmware upgrades. Some Intel SSDs in particular (600p and 6000p) require a firmware upgrade for [possible data corruption](https://bugzilla.redhat.com/show_bug.cgi?id=1402533) induced by XFS I/O usage patterns. The problem is at the firmware level and not any fault of the XFS filesystem. The smartctl utility can help check the device model and firmware version.

### Introduction\[ [edit](/index.php?title=Handbook:Parts/Installation/Disks&action=edit&section=T-4 "Edit section: ")\]

Now that the partitions have been created, it is time to place a filesystem on them. In the next section the various file systems that Linux supports are described. Readers that already know which filesystem to use can continue with [Applying a filesystem to a partition](/wiki/Handbook:HPPA/Installation/Disks#Applying_a_filesystem_to_a_partition "Handbook:HPPA/Installation/Disks"). The others should read on to learn about the available filesystems...

### Filesystems\[ [edit](/index.php?title=Handbook:Parts/Installation/Disks&action=edit&section=T-5 "Edit section: ")\]

Linux supports several dozen filesystems, although many of them are only wise to deploy for specific purposes. Only certain filesystems may be found stable on the hppa architecture - it is advised to read up on the filesystems and their support state before selecting a more experimental one for important partitions. **XFS is the recommended all-purpose, all-platform filesystem.** The below is a non-exhaustive list:

[XFS](/wiki/XFS "XFS")Filesystem with metadata journaling which comes with a robust feature-set and is optimized for scalability. It has been continuously upgraded to include modern features. The only downside is that XFS partitions cannot yet be shrunk, although this is being worked on. XFS notably supports reflinks and Copy on Write (CoW) which is particularly helpful on Gentoo systems because of the amount of compiles users complete. XFS is the recommended modern all-purpose all-platform filesystem. Requires a partition to be at least 300MB.[ext4](/wiki/Ext4 "Ext4")Ext4 is a reliable, all-purpose all-platform filesystem, although it lacks modern features like reflinks.[VFAT](/wiki/VFAT "VFAT")Also known as FAT32, is supported by Linux but does not support standard UNIX permission settings. It is mostly used for interoperability/interchange with other operating systems (Microsoft Windows or Apple's macOS) but is also a necessity for some system bootloader firmware (like UEFI). Users of UEFI systems will need an [EFI System Partition](/wiki/EFI_System_Partition "EFI System Partition") formatted with VFAT in order to boot.[btrfs](/wiki/Btrfs "Btrfs")Newer generation filesystem. Provides advanced features like snapshotting, self-healing through checksums, transparent compression, subvolumes, and integrated RAID. Kernels prior to 5.4.y are not guaranteed to be safe to use with btrfs in production because fixes for serious issues are only present in the more recent releases of the LTS kernel branches. RAID 5/6 and quota groups unsafe on all versions of btrfs.[F2FS](/wiki/F2FS "F2FS")The Flash-Friendly File System was originally created by Samsung for the use with NAND flash memory. It is a decent choice when installing Gentoo to microSD cards, USB drives, or other flash-based storage devices.[NTFS](/wiki/NTFS "NTFS")This "New Technology" filesystem is the flagship filesystem of Microsoft Windows since Windows NT 3.1. Similarly to VFAT, it does not store UNIX permission settings or extended attributes necessary for BSD or Linux to function properly, therefore it should not be used as a root filesystem for most cases. It should _only_ be used for interoperability or data interchange with Microsoft Windows systems (note the emphasis on _only_).[ZFS](/wiki/ZFS "ZFS")**Important:** ZFS pools can be created on the minimal installation CD, for further information, refer to [ZFS/rootfs](/wiki/ZFS/rootfs "ZFS/rootfs")Next generation file system created by Matthew Ahrens and Jeff Bonwick. It was designed around a few key ideas: Administration of storage should be simple, redundancy should be handled by the filesystem, file systems should never be taken offline for repair, automated simulations of worst case scenarios before shipping code is important, and data integrity is paramount.

More extensive information on filesystems can be found in the community maintained [Filesystem article](/wiki/Filesystem "Filesystem").

### Applying a filesystem to a partition\[ [edit](/index.php?title=Handbook:Parts/Installation/Disks&action=edit&section=T-6 "Edit section: ")\]

**Note**

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
[sys-fs/dosfstools](https://packages.gentoo.org/packages/sys-fs/dosfstools)[btrfs](/wiki/Btrfs "Btrfs")mkfs.btrfs Yes
[sys-fs/btrfs-progs](https://packages.gentoo.org/packages/sys-fs/btrfs-progs)[F2FS](/wiki/F2FS "F2FS")mkfs.f2fs Yes
[sys-fs/f2fs-tools](https://packages.gentoo.org/packages/sys-fs/f2fs-tools)[NTFS](/wiki/NTFS "NTFS")mkfs.ntfs Yes
[sys-fs/ntfs3g](https://packages.gentoo.org/packages/sys-fs/ntfs3g)[ZFS](/wiki/ZFS "ZFS")zpool create ... Yes
[sys-fs/zfs](https://packages.gentoo.org/packages/sys-fs/zfs)

**Important**

The handbook recommends new partitions as part of the installation process, but it is important to note running any mkfs command will erase any data contained within the partition. When necessary, ensure any data that exists within is appropriately backed up before creating a new filesystem.

For instance, to have the root partition (/dev/sda4) as xfs as used in the example partition structure, the following commands would be used:

`root #` `mkfs.xfs -c options=/usr/share/xfsprogs/mkfs/lts_6.18.conf /dev/sda4`

#### Legacy BIOS boot partition filesystem\[ [edit](/index.php?title=Handbook:Parts/Installation/Disks&action=edit&section=T-8 "Edit section: ")\]

Systems booting via legacy BIOS with a MBR/DOS disklabel can use any filesystem format supported by the bootloader.

For example, to format with XFS:

`root #` `mkfs.xfs -c options=/usr/share/xfsprogs/mkfs/lts_6.18.conf /dev/sda2`

#### Small ext4 partitions\[ [edit](/index.php?title=Handbook:Parts/Installation/Disks&action=edit&section=T-9 "Edit section: ")\]

When using the ext4 filesystem on a small partition (less than 8 GiB), the filesystem should be created with the proper options to reserve enough inodes. This can specified using the `-T small` option:

`root #` `mkfs.ext4 -T small /dev/<device>`

Doing so will quadruple the number of inodes for a given filesystem, since its "bytes-per-inode" reduces from one every 16kB to one every 4kB.

### Activating the swap partition\[ [edit](/index.php?title=Handbook:Parts/Installation/Disks&action=edit&section=T-10 "Edit section: ")\]

mkswap is the command that is used to initialize swap partitions:

`root #` `mkswap /dev/sda3`

**Note**

Installations which were previously started, but did not finish the installation process can resume the installation from this point in the handbook. Use this link as the permalink: [Resumed installations start here](/wiki/Handbook:HPPA/Installation/Disks#Resumed_installations_start_here "Handbook:HPPA/Installation/Disks").

To activate the swap partition, use swapon:

`root #` `swapon /dev/sda3`

This 'activation' step is only necessary because the swap partition is newly created within the live environment. Once the system has been rebooted, as long as the swap partition is properly defined within fstab or other mount mechanism, swap space will activate automatically.

## Mounting the root partition\[ [edit](/index.php?title=Handbook:Parts/Installation/Disks&action=edit&section=T-11 "Edit section: ")\]

Certain live environments may be missing the suggested mount point for Gentoo's root partition (/mnt/gentoo), or mount points for additional partitions created in the partitioning section:

`root #` `mkdir --parents /mnt/gentoo`

Continue creating additional mount points necessary for any additional (custom) partition(s) created during previous steps by using the mkdir command.

With mount points created, it is time to make the partitions accessible via mount command.

Mount the root partition:

`root #` `mount /dev/sda4 /mnt/gentoo`

Continue mounting additional (custom) partitions as necessary using the mount command.

**Note**

If /tmp/ needs to reside on a separate partition, be sure to change its permissions after mounting:

`root #` `chmod 1777 /mnt/gentoo/tmp`

This also holds for /var/tmp.

Later in the instructions, the proc filesystem (a virtual interface with the kernel) as well as other kernel pseudo-filesystems will be mounted. But first [the Gentoo stage file](/wiki/Handbook:HPPA/Installation/Stage "Handbook:HPPA/Installation/Stage") must be extracted.

[← Configuring the network](/wiki/Handbook:HPPA/Installation/Networking "Previous part") [Home](/wiki/Handbook:HPPA "Handbook:HPPA") [Installing the stage file →](/wiki/Handbook:HPPA/Installation/Stage "Next part")