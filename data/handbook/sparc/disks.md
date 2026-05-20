# Disks

Other languages:

- [Deutsch](/wiki/Handbook:SPARC/Installation/Disks/de "Handbuch:SPARC/Installation/Festplatten (100% translated)")
- English
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

[SPARC Handbook](/wiki/Handbook:SPARC "Handbook:SPARC")[Installation](/wiki/Handbook:SPARC/Full/Installation "Handbook:SPARC/Full/Installation")[About the installation](/wiki/Handbook:SPARC/Installation/About "Handbook:SPARC/Installation/About")[Choosing the media](/wiki/Handbook:SPARC/Installation/Media "Handbook:SPARC/Installation/Media")[Configuring the network](/wiki/Handbook:SPARC/Installation/Networking "Handbook:SPARC/Installation/Networking")Preparing the disks[The stage file](/wiki/Handbook:SPARC/Installation/Stage "Handbook:SPARC/Installation/Stage")[Installing base system](/wiki/Handbook:SPARC/Installation/Base "Handbook:SPARC/Installation/Base")[Configuring the kernel](/wiki/Handbook:SPARC/Installation/Kernel "Handbook:SPARC/Installation/Kernel")[Configuring the system](/wiki/Handbook:SPARC/Installation/System "Handbook:SPARC/Installation/System")[Installing tools](/wiki/Handbook:SPARC/Installation/Tools "Handbook:SPARC/Installation/Tools")[Configuring the bootloader](/wiki/Handbook:SPARC/Installation/Bootloader "Handbook:SPARC/Installation/Bootloader")[Finalizing](/wiki/Handbook:SPARC/Installation/Finalizing "Handbook:SPARC/Installation/Finalizing")[Working with Gentoo](/wiki/Handbook:SPARC/Full/Working "Handbook:SPARC/Full/Working")[Portage introduction](/wiki/Handbook:SPARC/Working/Portage "Handbook:SPARC/Working/Portage")[USE flags](/wiki/Handbook:SPARC/Working/USE "Handbook:SPARC/Working/USE")[Portage features](/wiki/Handbook:SPARC/Working/Features "Handbook:SPARC/Working/Features")[Initscript system](/wiki/Handbook:SPARC/Working/Initscripts "Handbook:SPARC/Working/Initscripts")[Environment variables](/wiki/Handbook:SPARC/Working/EnvVar "Handbook:SPARC/Working/EnvVar")[Working with Portage](/wiki/Handbook:SPARC/Full/Portage "Handbook:SPARC/Full/Portage")[Files and directories](/wiki/Handbook:SPARC/Portage/Files "Handbook:SPARC/Portage/Files")[Variables](/wiki/Handbook:SPARC/Portage/Variables "Handbook:SPARC/Portage/Variables")[Mixing software branches](/wiki/Handbook:SPARC/Portage/Branches "Handbook:SPARC/Portage/Branches")[Additional tools](/wiki/Handbook:SPARC/Portage/Tools "Handbook:SPARC/Portage/Tools")[Custom package repository](/wiki/Handbook:SPARC/Portage/CustomTree "Handbook:SPARC/Portage/CustomTree")[Advanced features](/wiki/Handbook:SPARC/Portage/Advanced "Handbook:SPARC/Portage/Advanced")[OpenRC network configuration](/wiki/Handbook:SPARC/Full/Networking "Handbook:SPARC/Full/Networking")[Getting started](/wiki/Handbook:SPARC/Networking/Introduction "Handbook:SPARC/Networking/Introduction")[Advanced configuration](/wiki/Handbook:SPARC/Networking/Advanced "Handbook:SPARC/Networking/Advanced")[Modular networking](/wiki/Handbook:SPARC/Networking/Modular "Handbook:SPARC/Networking/Modular")[Wireless](/wiki/Handbook:SPARC/Networking/Wireless "Handbook:SPARC/Networking/Wireless")[Adding functionality](/wiki/Handbook:SPARC/Networking/Extending "Handbook:SPARC/Networking/Extending")[Dynamic management](/wiki/Handbook:SPARC/Networking/Dynamic "Handbook:SPARC/Networking/Dynamic")

## Contents

- [1Introduction to block devices](#Introduction_to_block_devices)
  - [1.1Block devices](#Block_devices)
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

### Partition tables\[ [edit](/index.php?title=Handbook:SPARC/Blocks/Disks&action=edit&section=T-1 "Edit section: ")\]

Although it is theoretically possible to use a raw, unpartitioned disk to house a Linux system (when creating a btrfs RAID for example), this is almost never done in practice. Instead, disk block devices are split up into smaller, more manageable block devices. On **sparc** systems, these are called partitions. There are currently two standard partitioning technologies in use: Sun and GPT; the latter is supported only on more recent systems with a sufficiently recent firmware.

#### GUID Partition Table (GPT)\[ [edit](/index.php?title=Handbook:SPARC/Blocks/Disks&action=edit&section=T-2 "Edit section: ")\]

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

#### Sun partition table\[ [edit](/index.php?title=Handbook:SPARC/Blocks/Disks&action=edit&section=T-3 "Edit section: ")\]

Systems not manufactured by Oracle, T3 or earlier systems, or systems running an earlier firmware must use the Sun partition table type.

The third partition on Sun systems is set aside as a special "whole disk" slice. This partition must **not** contain a file system.

Users who are used to the DOS partitioning scheme should note that Sun partition tables do not have "primary" and "extended" partitions. Instead, up to eight partitions are available per drive, with the third of these being reserved.

The Handbook authors suggest using [GPT](#GPT) whenever possible for Gentoo installations.

### Default partitioning scheme\[ [edit](/index.php?title=Handbook:SPARC/Blocks/Disks&action=edit&section=T-4 "Edit section: ")\]

Due to the differences in required partition layout between GPT and Sun partition tables, a single partitioning scheme is not sufficient to support all possible system requirements. Some example schemes are provided below.

#### GPT partition scheme\[ [edit](/index.php?title=Handbook:SPARC/Blocks/Disks&action=edit&section=T-5 "Edit section: ")\]

The following partitioning scheme will be used as an example for GPT-formatted disks:

Partition
Filesystem
Size
Mount Point
Description
/dev/sda1(none)
2M
none
[BIOS boot partition](/wiki/Handbook:X86/Blocks/Disks#What_is_the_BIOS_boot_partition.3F "Handbook:X86/Blocks/Disks")/dev/sda2(swap)
RAM size \* 2
none
Swap partition
/dev/sda3ext4
Rest of the disk
/Root partition

#### Sun formatted partition scheme\[ [edit](/index.php?title=Handbook:SPARC/Blocks/Disks&action=edit&section=T-6 "Edit section: ")\]

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

**Important**

SPARC systems using OBP version 3 or older have additional restrictions on their partitioning scheme. The root partition must be the first partition on the disk, and it may be no larger than 2 GiB. For this reason, such systems will require additional sufficiently-sized partitions for top-level directories, such as /usr, /var, /home, and other directories which would likely cause the root partition to exceed this limit. These systems are also likely to require the Sun partition table type, so do not forget to include the whole disk partition.

## Partitioning the disk with GPT\[ [edit](/index.php?title=Handbook:SPARC/Blocks/Disks&action=edit&section=T-7 "Edit section: ")\]

The following parts explain how to create the example partition layout for a GPT installation using fdisk. The example partition layout was mentioned earlier:

Partition
Description
/dev/sda1Boot partition
/dev/sda2Swap partition
/dev/sda3Root partition

Change the partition layout according the system's needs.

### Viewing the current partition layout\[ [edit](/index.php?title=Handbook:SPARC/Blocks/Disks&action=edit&section=T-8 "Edit section: ")\]

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

<!--T:79-->
Device     Start      End  Sectors  Size Type
/dev/sda1   2048 30547967 30545920 14.6G Linux filesystem

```

### Creating a new disklabel and removing all existing partitions\[ [edit](/index.php?title=Handbook:SPARC/Blocks/Disks&action=edit&section=T-9 "Edit section: ")\]

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

### Creating the BIOS boot partition\[ [edit](/index.php?title=Handbook:SPARC/Blocks/Disks&action=edit&section=T-10 "Edit section: ")\]

First, create the BIOS boot partition. Type `n` to create a new partition, followed by `1` to select the first partition. When prompted for the first sector, make sure it starts from 2048 (which may be needed for the boot loader) and hit `Enter`. When prompted for the last sector, type +2M to create a partition 2 Mbyte in size:

`Command (m for help):` `n`

```
Partition number (1-128, default 1):
First sector (2048-30548062, default 2048):
Last sector, +/-sectors or +/-size{K,M,G,T,P} (2048-30548062, default 30547967): +2M

<!--T:92-->
Created a new partition 1 of type 'Linux filesystem' and of size 2 MiB.

```

Mark the partition as a BIOS boot partition:

`Command (m for help):` `t`

```
Selected partition 1
Partition type or alias (type L to list all): 4
Changed type of partition 'Linux filesystem' to 'BIOS boot'.

```

### Creating the swap partition\[ [edit](/index.php?title=Handbook:SPARC/Blocks/Disks&action=edit&section=T-11 "Edit section: ")\]

Next, to create the swap partition, type `n` to create a new partition, then type `2` to create the second partition, /dev/sda2. When prompted for the first sector, hit `Enter`. When prompted for the last sector, type +4G (or any other size needed for the swap space) to create a partition 4 GiB in size.

`Command (m for help):` `n`

```
Partition number (2-128, default 2):
First sector (6144-30548062, default 6144):
Last sector, +/-sectors or +/-size{K,M,G,T,P} (6144-30548062, default 30547967): +4G

<!--T:98-->
Created a new partition 2 of type 'Linux filesystem' and of size 4 GiB.

```

After all this is done, type `t` to set the partition type, `2` to select the partition just created and then type in _19_ to set the partition type to "Linux Swap".

`Command (m for help):` `t`

```
Partition number (1,2, default 2): 2
Partition type (type L to list all types): 19

Changed type of partition 'Linux filesystem' to 'Linux swap'.

```

### Creating the root partition\[ [edit](/index.php?title=Handbook:SPARC/Blocks/Disks&action=edit&section=T-12 "Edit section: ")\]

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

<!--T:104-->
Device       Start      End  Sectors  Size Type
/dev/sda1     2048     6143     4096    2M BIOS boot
/dev/sda2     6144  8394751  8388608    4G Linux swap
/dev/sda3  8394752 30547967 22153216 10.6G Linux filesystem

```

### Saving the partition layout\[ [edit](/index.php?title=Handbook:SPARC/Blocks/Disks&action=edit&section=T-13 "Edit section: ")\]

To save the partition layout and exit fdisk, type `w`.

`Command (m for help):` `w`

With the partitions created, it is now time to put filesystems on them.

## Partitioning the disk with a Sun partition table\[ [edit](/index.php?title=Handbook:SPARC/Blocks/Disks&action=edit&section=T-14 "Edit section: ")\]

The following parts explain how to create the example partition layout for a Sun partition table installation using fdisk. The example partition layout was mentioned earlier:

Partition
Description
/dev/sda1Root partition
/dev/sda2Swap partition
/dev/sda3Whole disk partition

Change the partition layout according to personal preference. If partitioning for a system using OBP version 3 or earlier, ensure that the root partition is less than 2G in size, and additionally create partitions /dev/sda4 and onward for additional filesystems.

### Viewing the current partition layout\[ [edit](/index.php?title=Handbook:SPARC/Blocks/Disks&action=edit&section=T-15 "Edit section: ")\]

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

<!--T:117-->
Device        Start      End  Sectors  Size Id Type         Flags
/dev/sda1         0 30445567 30445568 14.5G 83 Linux native
/dev/sda2  30445568 30547967   102400   50M 82 Linux swap      u
/dev/sda3         0 30547967 30547968 14.6G  5 Whole disk

```

### Creating a new disklabel / removing all partitions\[ [edit](/index.php?title=Handbook:SPARC/Blocks/Disks&action=edit&section=T-16 "Edit section: ")\]

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

<!--T:123-->
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

### Creating the whole disk partition\[ [edit](/index.php?title=Handbook:SPARC/Blocks/Disks&action=edit&section=T-17 "Edit section: ")\]

First, create the whole disk partition. Type `n` to create a new partition, followed by `3` to select the third partition. When prompted for the first sector, hit `Enter`. When prompted for the last sector, hit `Enter` to create a partition that takes up all of the space on the disk.

`Command (m for help):` `n`

```
Partition number (1-8, default 1): 3

<!--T:131-->
It is highly recommended that the third partition covers the whole disk and is of type `Whole disk'
First sector (0-30547968, default 0):
Last sector or +/-sectors or +/-size{K,M,G,T,P} (0-30547968, default 30547968):

<!--T:132-->
Created a new partition 3 of type 'Whole disk' and of size 14.6 GiB.

```

fdisk will automatically set the type of such a partition to 'Whole disk', so there is no need to explicitly set the type.

### Creating the root partition\[ [edit](/index.php?title=Handbook:SPARC/Blocks/Disks&action=edit&section=T-18 "Edit section: ")\]

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

<!--T:137-->
Device     Start      End  Sectors  Size Id Type         Flags
/dev/sda1      0 22159359 22159360 10.6G 83 Linux native
/dev/sda3      0 30547967 30547968 14.6G  5 Whole disk

<!--T:138-->

```

### Creating the swap partition\[ [edit](/index.php?title=Handbook:SPARC/Blocks/Disks&action=edit&section=T-19 "Edit section: ")\]

Finally, to create the swap partition, type `n` to create a new partition, then type `2` to create the second partition, /dev/sda2. When prompted for the first sector, hit `Enter`. When prompted for the last sector, hit `Enter` to take up the remaining space on the disk.

`Command (m for help):` `n`

```
Partition number (2,4-8, default 2):
First sector (22159360-30547968, default 22159360):
Last sector or +/-sectors or +/-size{K,M,G,T,P} (22159360-30547968, default 30547968):

<!--T:142-->
Created a new partition 2 of type 'Linux native' and of size 4 GiB.

```

After all this is done, type `t` to set the partition type, `2` to select the partition just created and then type in _82_ to set the partition type to "Linux Swap".

`Command (m for help):` `t`

```
Partition number (1-3, default 3): 2
Hex code (type L to list all codes): 82

<!--T:145-->
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

<!--T:148-->
Device        Start      End  Sectors  Size Id Type         Flags
/dev/sda1         0 22159359 22159360 10.6G 83 Linux native
/dev/sda2  22159360 30547967  8388608    4G 82 Linux swap      u
/dev/sda3         0 30547967 30547968 14.6G  5 Whole disk

```

### Saving the partition layout\[ [edit](/index.php?title=Handbook:SPARC/Blocks/Disks&action=edit&section=T-20 "Edit section: ")\]

To save the partition layout and exit fdisk, type `w`.

`Command (m for help):` `w`

With the partitions created, it is now time to put filesystems on them.

## Creating file systems\[ [edit](/index.php?title=Handbook:Parts/Installation/Disks&action=edit&section=T-3 "Edit section: ")\]

**Warning**

When using an SSD or NVMe drive, it is wise to check for firmware upgrades. Some Intel SSDs in particular (600p and 6000p) require a firmware upgrade for [possible data corruption](https://bugzilla.redhat.com/show_bug.cgi?id=1402533) induced by XFS I/O usage patterns. The problem is at the firmware level and not any fault of the XFS filesystem. The smartctl utility can help check the device model and firmware version.

### Introduction\[ [edit](/index.php?title=Handbook:Parts/Installation/Disks&action=edit&section=T-4 "Edit section: ")\]

Now that the partitions have been created, it is time to place a filesystem on them. In the next section the various file systems that Linux supports are described. Readers that already know which filesystem to use can continue with [Applying a filesystem to a partition](/wiki/Handbook:SPARC/Installation/Disks#Applying_a_filesystem_to_a_partition "Handbook:SPARC/Installation/Disks"). The others should read on to learn about the available filesystems...

### Filesystems\[ [edit](/index.php?title=Handbook:Parts/Installation/Disks&action=edit&section=T-5 "Edit section: ")\]

Linux supports several dozen filesystems, although many of them are only wise to deploy for specific purposes. Only certain filesystems may be found stable on the sparc architecture - it is advised to read up on the filesystems and their support state before selecting a more experimental one for important partitions. **XFS is the recommended all-purpose, all-platform filesystem.** The below is a non-exhaustive list:

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

For instance, to have the root partition (/dev/sda1) as xfs as used in the example partition structure, the following commands would be used:

`root #` `mkfs.xfs -c options=/usr/share/xfsprogs/mkfs/lts_6.18.conf /dev/sda1`

#### Legacy BIOS boot partition filesystem\[ [edit](/index.php?title=Handbook:Parts/Installation/Disks&action=edit&section=T-8 "Edit section: ")\]

Systems booting via legacy BIOS with a MBR/DOS disklabel can use any filesystem format supported by the bootloader.

For example, to format with XFS:

`root #` `mkfs.xfs -c options=/usr/share/xfsprogs/mkfs/lts_6.18.conf`

#### Small ext4 partitions\[ [edit](/index.php?title=Handbook:Parts/Installation/Disks&action=edit&section=T-9 "Edit section: ")\]

When using the ext4 filesystem on a small partition (less than 8 GiB), the filesystem should be created with the proper options to reserve enough inodes. This can specified using the `-T small` option:

`root #` `mkfs.ext4 -T small /dev/<device>`

Doing so will quadruple the number of inodes for a given filesystem, since its "bytes-per-inode" reduces from one every 16kB to one every 4kB.

### Activating the swap partition\[ [edit](/index.php?title=Handbook:Parts/Installation/Disks&action=edit&section=T-10 "Edit section: ")\]

mkswap is the command that is used to initialize swap partitions:

`root #` `mkswap /dev/sda2`

**Note**

Installations which were previously started, but did not finish the installation process can resume the installation from this point in the handbook. Use this link as the permalink: [Resumed installations start here](/wiki/Handbook:SPARC/Installation/Disks#Resumed_installations_start_here "Handbook:SPARC/Installation/Disks").

To activate the swap partition, use swapon:

`root #` `swapon /dev/sda2`

This 'activation' step is only necessary because the swap partition is newly created within the live environment. Once the system has been rebooted, as long as the swap partition is properly defined within fstab or other mount mechanism, swap space will activate automatically.

## Mounting the root partition\[ [edit](/index.php?title=Handbook:Parts/Installation/Disks&action=edit&section=T-11 "Edit section: ")\]

Certain live environments may be missing the suggested mount point for Gentoo's root partition (/mnt/gentoo), or mount points for additional partitions created in the partitioning section:

`root #` `mkdir --parents /mnt/gentoo`

Continue creating additional mount points necessary for any additional (custom) partition(s) created during previous steps by using the mkdir command.

With mount points created, it is time to make the partitions accessible via mount command.

Mount the root partition:

`root #` `mount /dev/sda1 /mnt/gentoo`

Continue mounting additional (custom) partitions as necessary using the mount command.

**Note**

If /tmp/ needs to reside on a separate partition, be sure to change its permissions after mounting:

`root #` `chmod 1777 /mnt/gentoo/tmp`

This also holds for /var/tmp.

Later in the instructions, the proc filesystem (a virtual interface with the kernel) as well as other kernel pseudo-filesystems will be mounted. But first [the Gentoo stage file](/wiki/Handbook:SPARC/Installation/Stage "Handbook:SPARC/Installation/Stage") must be extracted.

[← Configuring the network](/wiki/Handbook:SPARC/Installation/Networking "Previous part") [Home](/wiki/Handbook:SPARC "Handbook:SPARC") [Installing the stage file →](/wiki/Handbook:SPARC/Installation/Stage "Next part")