# Disks

Other languages:

- [Deutsch](/wiki/Handbook:MIPS/Installation/Disks/de "Handbuch:MIPS/Installation/Festplatten (100% translated)")
- [English](/wiki/Handbook:MIPS/Installation/Disks "Handbook:MIPS/Installation/Disks (100% translated)")
- [español](/wiki/Handbook:MIPS/Installation/Disks/es "Manual de Gentoo: MIPS/Instalación/Discos (100% translated)")
- [français](/wiki/Handbook:MIPS/Installation/Disks/fr "Handbook:MIPS/Installation/Disks/fr (100% translated)")
- [italiano](/wiki/Handbook:MIPS/Installation/Disks/it "Handbook:MIPS/Installation/Disks/it (100% translated)")
- [magyar](/wiki/Handbook:MIPS/Installation/Disks/hu "Handbook:MIPS/Installation/Disks/hu (100% translated)")
- polski
- [português do Brasil](/wiki/Handbook:MIPS/Installation/Disks/pt-br "Manual:MIPS/Instalação/Discos (100% translated)")
- [русский](/wiki/Handbook:MIPS/Installation/Disks/ru "Handbook:MIPS/Installation/Disks (100% translated)")
- [தமிழ்](/wiki/Handbook:MIPS/Installation/Disks/ta "கையேடு:MIPS/நிறுவல்/வட்டுக்கள் (100% translated)")
- [中文（中国大陆）‎](/wiki/Handbook:MIPS/Installation/Disks/zh-cn "手册：MIPS/安装/准备磁盘 (100% translated)")
- [日本語](/wiki/Handbook:MIPS/Installation/Disks/ja "ハンドブック:MIPS/インストール/ディスク (100% translated)")
- [한국어](/wiki/Handbook:MIPS/Installation/Disks/ko "Handbook:MIPS/Installation/Disks/ko (100% translated)")

[MIPS Handbook](/wiki/Handbook:MIPS/pl "Handbook:MIPS/pl")[Instalacja](/wiki/Handbook:MIPS/Full/Installation/pl "Handbook:MIPS/Full/Installation/pl")[O instalacji](/wiki/Handbook:MIPS/Installation/About/pl "Handbook:MIPS/Installation/About/pl")[Wybór medium instalacyjnego](/wiki/Handbook:MIPS/Installation/Media/pl "Handbook:MIPS/Installation/Media/pl")[Konfiguracja sieci](/wiki/Handbook:MIPS/Installation/Networking/pl "Handbook:MIPS/Installation/Networking/pl")Przygotowanie dysków[Instalacja etapu 3](/wiki/Handbook:MIPS/Installation/Stage/pl "Handbook:MIPS/Installation/Stage/pl")[Instalacja systemu podstawowego](/wiki/Handbook:MIPS/Installation/Base/pl "Handbook:MIPS/Installation/Base/pl")[Konfiguracja jądra](/wiki/Handbook:MIPS/Installation/Kernel/pl "Handbook:MIPS/Installation/Kernel/pl")[Konfiguracja systemu](/wiki/Handbook:MIPS/Installation/System/pl "Handbook:MIPS/Installation/System/pl")[Instalacja narzędzi](/wiki/Handbook:MIPS/Installation/Tools/pl "Handbook:MIPS/Installation/Tools/pl")[Instalacja systemu rozruchowego](/wiki/Handbook:MIPS/Installation/Bootloader/pl "Handbook:MIPS/Installation/Bootloader/pl")[Finalizacja](/wiki/Handbook:MIPS/Installation/Finalizing/pl "Handbook:MIPS/Installation/Finalizing/pl")[Praca z Gentoo](/wiki/Handbook:MIPS/Full/Working/pl "Handbook:MIPS/Full/Working/pl")[Wstęp do Portage](/wiki/Handbook:MIPS/Working/Portage/pl "Handbook:MIPS/Working/Portage/pl")[Flagi USE](/wiki/Handbook:MIPS/Working/USE/pl "Handbook:MIPS/Working/USE/pl")[Funkcje portage](/wiki/Handbook:MIPS/Working/Features/pl "Handbook:MIPS/Working/Features/pl")[System initscript](/wiki/Handbook:MIPS/Working/Initscripts/pl "Handbook:MIPS/Working/Initscripts/pl")[Zmienne środowiskowe](/wiki/Handbook:MIPS/Working/EnvVar/pl "Handbook:MIPS/Working/EnvVar/pl")[Praca z Portage](/wiki/Handbook:MIPS/Full/Portage/pl "Handbook:MIPS/Full/Portage/pl")[Pliki i katalogi](/wiki/Handbook:MIPS/Portage/Files/pl "Handbook:MIPS/Portage/Files/pl")[Zmienne](/wiki/Handbook:MIPS/Portage/Variables/pl "Handbook:MIPS/Portage/Variables/pl")[Mieszanie działów oprogramowania](/wiki/Handbook:MIPS/Portage/Branches/pl "Handbook:MIPS/Portage/Branches/pl")[Dodatkowe narzędzia](/wiki/Handbook:MIPS/Portage/Tools/pl "Handbook:MIPS/Portage/Tools/pl")[Custom package repository](/wiki/Handbook:MIPS/Portage/CustomTree/pl "Handbook:MIPS/Portage/CustomTree/pl")[Funkcje zaawansowane](/wiki/Handbook:MIPS/Portage/Advanced/pl "Handbook:MIPS/Portage/Advanced/pl")[Konfiguracja sieci](/wiki/Handbook:MIPS/Full/Networking/pl "Handbook:MIPS/Full/Networking/pl")[Zaczynamy](/wiki/Handbook:MIPS/Networking/Introduction/pl "Handbook:MIPS/Networking/Introduction/pl")[Zaawansowana konfiguracja](/wiki/Handbook:MIPS/Networking/Advanced/pl "Handbook:MIPS/Networking/Advanced/pl")[Sieć modularna](/wiki/Handbook:MIPS/Networking/Modular/pl "Handbook:MIPS/Networking/Modular/pl")[Sieć bezprzewodowa](/wiki/Handbook:MIPS/Networking/Wireless/pl "Handbook:MIPS/Networking/Wireless/pl")[Dodawanie funkcjonalności](/wiki/Handbook:MIPS/Networking/Extending/pl "Handbook:MIPS/Networking/Extending/pl")[Dynamiczne zarządzanie](/wiki/Handbook:MIPS/Networking/Dynamic/pl "Handbook:MIPS/Networking/Dynamic/pl")

## Contents

- [1Introduction to block devices](#Introduction_to_block_devices)
  - [1.1Block devices](#Block_devices)
  - [1.2Partitions](#Partitions)
- [2Designing a partition scheme](#Designing_a_partition_scheme)
  - [2.1How many partitions and how big?](#How_many_partitions_and_how_big.3F)
  - [2.2What about swap space?](#What_about_swap_space.3F)
- [3Using fdisk](#Using_fdisk)
  - [3.1SGI machines: Creating an SGI disk label](#SGI_machines:_Creating_an_SGI_disk_label)
  - [3.2Resizing the SGI volume header](#Resizing_the_SGI_volume_header)
  - [3.3Partitioning Cobalt drives](#Partitioning_Cobalt_drives)
- [4Creating file systems](#Creating_file_systems)
  - [4.1Introduction](#Introduction)
  - [4.2Filesystems](#Filesystems)
  - [4.3Applying a filesystem to a partition](#Applying_a_filesystem_to_a_partition)
    - [4.3.1Legacy BIOS boot partition filesystem](#Legacy_BIOS_boot_partition_filesystem)
    - [4.3.2Small ext4 partitions](#Small_ext4_partitions)
  - [4.4Activating the swap partition](#Activating_the_swap_partition)
- [5Mounting the root partition](#Mounting_the_root_partition)

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

Although it is theoretically possible to use a full disk to house your Linux system, this is almost never done in practice. Instead, full disk block devices are split up in smaller, more manageable block devices. These are called partitions.

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

## Using fdisk

### SGI machines: Creating an SGI disk label

All disks in an SGI System require an SGI Disk Label, which serves a similar function as Sun & MS-DOS disklabels -- It stores information about the disk partitions. Creating a new SGI Disk Label will create two special partitions on the disk:

- SGI Volume Header (9th partition): This partition is important. It is where the bootloader will reside, and in some cases, it will also contain the kernel images.
- SGI Volume (11th partition): This partition is similar in purpose to the Sun Disklabel's third partition of "Whole Disk". This partition spans the entire disk, and should be left untouched. It serves no special purpose other than to assist the PROM in some undocumented fashion (or it is used by IRIX in some way).

**Uwaga**

The SGI Volume Header must begin at cylinder 0. Failure to do so means a failure to boot from the disk.

The following is an example excerpt from an fdisk session. Read and tailor it to personal preference...

`root #` `fdisk /dev/sda`

Switch to expert mode:

`Command (m for help):` `x`

With `m` the full menu of options is displayed:

`Expert command (m for help):` `m`

```
Command action
   b   move beginning of data in a partition
   c   change number of cylinders
   d   print the raw data in the partition table
   e   list extended partitions
   f   fix partition order
   g   create an IRIX (SGI) partition table
   h   change number of heads
   m   print this menu
   p   print the partition table
   q   quit without saving changes
   r   return to main menu
   s   change number of sectors/track
   v   verify the partition table
   w   write table to disk and exit

```

Build an SGI disk label:

`Expert command (m for help):` `g`

```
Building a new SGI disklabel. Changes will remain in memory only,
until you decide to write them. After that, of course, the previous
content will be irrecoverably lost.

```

Return to the main menu:

`Expert command (m for help):` `r`

Take a look at the current partition layout:

`Command (m for help):` `p`

```
Disk /dev/sda (SGI disk label): 64 heads, 32 sectors, 17482 cylinders
Units = cylinders of 2048 * 512 bytes

----- partitions -----
Pt#     Device  Info     Start       End   Sectors  Id  System
 9:  /dev/sda1               0         4     10240   0  SGI volhdr
11:  /dev/sda2               0     17481  35803136   6  SGI volume
----- Bootinfo -----
Bootfile: /unix
----- Directory Entries -----

```

**Informacja**

If the disk already has an existing SGI Disklabel, then fdisk will not allow the creation of a new label. There are two ways around this. One is to create a Sun or MS-DOS disklabel, write the changes to disk, and restart fdisk. The second is to overwrite the partition table with null data via the following command: `dd if=/dev/zero of=/dev/sda bs=512 count=1`

### Resizing the SGI volume header

**Ważne**

This step is often needed, due to a bug in fdisk. For some reason, the volume header isn't created correctly, the end result being it starts and ends on cylinder 0. This prevents multiple partitions from being created. To get around this issue... read on.

Now that an SGI Disklabel is created, partitions may now be defined. In the above example, there are already two partitions defined. These are the special partitions mentioned above and should not normally be altered. However, for installing Gentoo, we'll need to load a bootloader, and possibly multiple kernel images (depending on system type) directly into the volume header. The volume header itself can hold up to eight images of any size, with each image allowed eight-character names.

The process of making the volume header larger isn't exactly straight-forward; there's a bit of a trick to it. One cannot simply delete and re-add the volume header due to odd fdisk behavior. In the example provided below, we'll create a 50MB Volume header in conjunction with a 50MB /boot/ partition. The actual layout of a disk may vary, but this is for illustrative purposes only.

Create a new partition:

`Command (m for help):` `n`

```
Partition number (1-16): 1
First cylinder (5-8682, default 5): 51
 Last cylinder (51-8682, default 8682): 101

```

Notice how fdisk only allows Partition #1 to be re-created starting at a minimum of cylinder 5? If we attempted to delete & re-create the SGI Volume Header this way, this is the same issue we would have encountered. In our example, we want /boot/ to be 50MB, so we start it at cylinder 51 (the Volume Header needs to start at cylinder 0, remember?), and set its ending cylinder to 101, which will roughly be 50MB (+/- 1-5MB).

Delete the partition:

`Command (m for help):` `d`

```
Partition number (1-16): 9

```

Now recreate it:

`Command (m for help):` `n`

```
Partition number (1-16): 9
First cylinder (0-50, default 0): 0
 Last cylinder (0-50, default 50): 50

```

If unsure how to use fdisk have a look down further at the instructions for partitioning on Cobalts. The concepts are exactly the same -- just remember to leave the volume header and whole disk partitions alone.

Once this is done, create the rest of your partitions as needed. After all the partitions are laid out, make sure to set the partition ID of the swap partition to 82, which is Linux Swap. By default, it will be 83, Linux Native.

### Partitioning Cobalt drives

On Cobalt machines, the BOOTROM expects to see a MS-DOS MBR, so partitioning the drive is relatively straightforward -- in fact, it's done the same way as done for an Intel x86 machine. However there are some things you need to bear in mind.

- Cobalt firmware will expect /dev/sda1 to be a Linux partition formatted EXT2 Revision 0. EXT2 Revision 1 partitions will NOT WORK! (The Cobalt BOOTROM only understands EXT2r0)
- The above said partition must contain a gzipped ELF image, vmlinux.gz in the root of that partition, which it loads as the kernel

For that reason, it is recommended to create a ~20MB /boot/ partition formatted EXT2r0 upon which to install CoLo & kernels. This allows the user to run a modern filesystem (EXT3 or ReiserFS) for the root filesystem.

In the example, it is assumed that /dev/sda1 is created to mount later as a /boot/ partition. To make this /, keep the PROM's expectations in mind.

So, continuing on... To create the partitions type fdisk /dev/sda at the prompt. The main commands to know are these:

CODE **List of important fdisk commands**

```
o: Wipe out old partition table, starting with an empty MS-DOS partition table
n: New Partition
t: Change Partition Type
    Use type 82 for Linux Swap, 83 for Linux FS
d: Delete a partition
p: Display (print) Partition Table
q: Quit -- leaving old partition table as is.
w: Quit -- writing partition table in the process.

```

`root #` `fdisk /dev/sda`

```
The number of cylinders for this disk is set to 19870.
There is nothing wrong with that, but this is larger than 1024,
and could in certain setups cause problems with:
1) software that runs at boot time (e.g., old versions of LILO)
2) booting and partitioning software from other OSs
   (e.g., DOS FDISK, OS/2 FDISK)

```

Start by clearing out any existing partitions:

`Command (m for help):` `o`

```
Building a new DOS disklabel. Changes will remain in memory only,
until you decide to write them. After that, of course, the previous
content won't be recoverable.


The number of cylinders for this disk is set to 19870.
There is nothing wrong with that, but this is larger than 1024,
and could in certain setups cause problems with:
1) software that runs at boot time (e.g., old versions of LILO)
2) booting and partitioning software from other OSs
   (e.g., DOS FDISK, OS/2 FDISK)
Warning: invalid flag 0x0000 of partition table 4 will be corrected by w(rite)

```

Now verify the partition table is empty using the `p` command:

`Command (m for help):` `p`

```
Disk /dev/sda: 10.2 GB, 10254827520 bytes
16 heads, 63 sectors/track, 19870 cylinders
Units = cylinders of 1008 * 512 = 516096 bytes

   Device Boot      Start         End      Blocks   Id  System

```

Create the /boot partition:

`Command (m for help):` `n`

```
Command action
   e   extended
   p   primary partition (1-4)
p
Partition number (1-4): 1
First cylinder (1-19870, default 1):
Last cylinder or +size or +sizeM or +sizeK (1-19870, default 19870): +20M

```

When printing the partitions, notice the newly created one:

`Command (m for help):` `p`

```
Disk /dev/sda: 10.2 GB, 10254827520 bytes
16 heads, 63 sectors/track, 19870 cylinders
Units = cylinders of 1008 * 512 = 516096 bytes

   Device Boot      Start         End      Blocks   Id  System
/dev/sda1               1          40       20128+  83  Linux

```

Let's now create an extended partition that covers the remainder of the disk. In that extended partition, we'll create the rest (logical partitions):

`Command (m for help):` `n`

```
Command action
   e   extended
   p   primary partition (1-4)
e
Partition number (1-4): 2
First cylinder (41-19870, default 41):
Using default value 41
Last cylinder or +size or +sizeM or +sizeK (41-19870, default 19870):
Using default value 19870

```

Now we create the / partition, /usr, /var, et.

`Command (m for help):` `n`

```
Command action
   l   logical (5 or over)
   p   primary partition (1-4)
l
First cylinder (41-19870, default 41):<Press ENTER>
Using default value 41
Last cylinder or +size or +sizeM or +sizeK (41-19870, default 19870): +500M

```

Repeat this as needed.

Last but not least, the swap space. It is recommended to have at least 250MB swap,
preferrably 1GB:

`Command (m for help):` `n`

```
Command action
   l   logical (5 or over)
   p   primary partition (1-4)
l
First cylinder (17294-19870, default 17294): <Press ENTER>
Using default value 17294
Last cylinder or +size or +sizeM or +sizeK (1011-19870, default 19870): <Press ENTER>
Using default value 19870

```

When checking the partition table, everything should be ready - one thing notwithstanding.

`Command (m for help):` `p`

```
Disk /dev/sda: 10.2 GB, 10254827520 bytes
16 heads, 63 sectors/track, 19870 cylinders
Units = cylinders of 1008 * 512 = 516096 bytes

Device Boot      Start         End      Blocks      ID  System
/dev/sda1               1          21       10552+  83  Linux
/dev/sda2              22       19870    10003896    5  Extended
/dev/sda5              22        1037      512032+  83  Linux
/dev/sda6            1038        5101     2048224+  83  Linux
/dev/sda7            5102        9165     2048224+  83  Linux
/dev/sda8            9166       13229     2048224+  83  Linux
/dev/sda9           13230       17293     2048224+  83  Linux
/dev/sda10          17294       19870     1298776+  83  Linux

```

Notice how #10, the swap partition is still type 83? Let's change that to the proper type:

`Command (m for help):` `t`

```
Partition number (1-10): 10
Hex code (type L to list codes): 82
Changed system type of partition 10 to 82 (Linux swap)

```

Now verify:

`Command (m for help):` `p`

```
Disk /dev/sda: 10.2 GB, 10254827520 bytes
16 heads, 63 sectors/track, 19870 cylinders
Units = cylinders of 1008 * 512 = 516096 bytes

Device Boot      Start         End      Blocks      ID  System
/dev/sda1               1          21       10552+  83  Linux
/dev/sda2              22       19870    10003896    5  Extended
/dev/sda5              22        1037      512032+  83  Linux
/dev/sda6            1038        5101     2048224+  83  Linux
/dev/sda7            5102        9165     2048224+  83  Linux
/dev/sda8            9166       13229     2048224+  83  Linux
/dev/sda9           13230       17293     2048224+  83  Linux
/dev/sda10          17294       19870     1298776+  82  Linux Swap

```

We write out the new partition table:

`Command (m for help):` `w`

```
The partition table has been altered!

Calling ioctl() to re-read partition table.
Syncing disks.

```

## Creating file systems

**Uwaga**

When using an SSD or NVMe drive, it is wise to check for firmware upgrades. Some Intel SSDs in particular (600p and 6000p) require a firmware upgrade for [possible data corruption](https://bugzilla.redhat.com/show_bug.cgi?id=1402533) induced by XFS I/O usage patterns. The problem is at the firmware level and not any fault of the XFS filesystem. The smartctl utility can help check the device model and firmware version.

### Introduction

Now that the partitions have been created, it is time to place a filesystem on them. In the next section the various file systems that Linux supports are described. Readers that already know which filesystem to use can continue with [Applying a filesystem to a partition](/wiki/Handbook:MIPS/Installation/Disks/pl#Applying_a_filesystem_to_a_partition "Handbook:MIPS/Installation/Disks/pl"). The others should read on to learn about the available filesystems...

### Filesystems

Linux supports several dozen filesystems, although many of them are only wise to deploy for specific purposes. Only certain filesystems may be found stable on the mips architecture - it is advised to read up on the filesystems and their support state before selecting a more experimental one for important partitions. **XFS is the recommended all-purpose, all-platform filesystem.** The below is a non-exhaustive list:

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

For instance, to have the root partition (/dev/sda5) as xfs as used in the example partition structure, the following commands would be used:

`root #` `mkfs.xfs -c options=/usr/share/xfsprogs/mkfs/lts_6.18.conf /dev/sda5`

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

`root #` `mkswap /dev/sda10`

**Informacja**

Installations which were previously started, but did not finish the installation process can resume the installation from this point in the handbook. Use this link as the permalink: [Resumed installations start here](/wiki/Handbook:MIPS/Installation/Disks/pl#Resumed_installations_start_here "Handbook:MIPS/Installation/Disks/pl").

To activate the swap partition, use swapon:

`root #` `swapon /dev/sda10`

This 'activation' step is only necessary because the swap partition is newly created within the live environment. Once the system has been rebooted, as long as the swap partition is properly defined within fstab or other mount mechanism, swap space will activate automatically.

## Mounting the root partition

Certain live environments may be missing the suggested mount point for Gentoo's root partition (/mnt/gentoo), or mount points for additional partitions created in the partitioning section:

`root #` `mkdir --parents /mnt/gentoo`

Continue creating additional mount points necessary for any additional (custom) partition(s) created during previous steps by using the mkdir command.

With mount points created, it is time to make the partitions accessible via mount command.

Mount the root partition:

`root #` `mount /dev/sda5 /mnt/gentoo`

Continue mounting additional (custom) partitions as necessary using the mount command.

**Informacja**

If /tmp/ needs to reside on a separate partition, be sure to change its permissions after mounting:

`root #` `chmod 1777 /mnt/gentoo/tmp`

This also holds for /var/tmp.

Later in the instructions, the proc filesystem (a virtual interface with the kernel) as well as other kernel pseudo-filesystems will be mounted. But first [the Gentoo stage file](/wiki/Handbook:MIPS/Installation/Stage/pl "Handbook:MIPS/Installation/Stage/pl") must be extracted.

[← Configuring the network](/wiki/Handbook:MIPS/Installation/Networking "Previous part") [Home](/wiki/Handbook:MIPS "Handbook:MIPS") [Installing the stage file →](/wiki/Handbook:MIPS/Installation/Stage "Next part")