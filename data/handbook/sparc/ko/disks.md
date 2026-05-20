# Disks

Other languages:

- [Deutsch](/wiki/Handbook:SPARC/Installation/Disks/de "Handbuch:SPARC/Installation/Festplatten (100% translated)")
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
- 한국어

[SPARC 핸드북](/wiki/Handbook:SPARC/ko "Handbook:SPARC/ko")[설치](/wiki/Handbook:SPARC/Full/Installation/ko "Handbook:SPARC/Full/Installation/ko")[설치 정보](/wiki/Handbook:SPARC/Installation/About/ko "Handbook:SPARC/Installation/About/ko")[매체 선택](/wiki/Handbook:SPARC/Installation/Media/ko "Handbook:SPARC/Installation/Media/ko")[네트워크 설정](/wiki/Handbook:SPARC/Installation/Networking/ko "Handbook:SPARC/Installation/Networking/ko")디스크 준비[스테이지 3 설치](/wiki/Handbook:SPARC/Installation/Stage/ko "Handbook:SPARC/Installation/Stage/ko")[베이스 시스템 설치](/wiki/Handbook:SPARC/Installation/Base/ko "Handbook:SPARC/Installation/Base/ko")[커널 설정](/wiki/Handbook:SPARC/Installation/Kernel/ko "Handbook:SPARC/Installation/Kernel/ko")[시스템 설정](/wiki/Handbook:SPARC/Installation/System/ko "Handbook:SPARC/Installation/System/ko")[도구 설치](/wiki/Handbook:SPARC/Installation/Tools/ko "Handbook:SPARC/Installation/Tools/ko")[부트로더 설정](/wiki/Handbook:SPARC/Installation/Bootloader/ko "Handbook:SPARC/Installation/Bootloader/ko")[마무리](/wiki/Handbook:SPARC/Installation/Finalizing/ko "Handbook:SPARC/Installation/Finalizing/ko")[젠투 활용](/wiki/Handbook:SPARC/Full/Working/ko "Handbook:SPARC/Full/Working/ko")[포티지 소개](/wiki/Handbook:SPARC/Working/Portage/ko "Handbook:SPARC/Working/Portage/ko")[USE 플래그](/wiki/Handbook:SPARC/Working/USE/ko "Handbook:SPARC/Working/USE/ko")[포티지 기능](/wiki/Handbook:SPARC/Working/Features/ko "Handbook:SPARC/Working/Features/ko")[초기화 스크립트 시스템](/wiki/Handbook:SPARC/Working/Initscripts/ko "Handbook:SPARC/Working/Initscripts/ko")[환경 변수](/wiki/Handbook:SPARC/Working/EnvVar/ko "Handbook:SPARC/Working/EnvVar/ko")[포티지 활용](/wiki/Handbook:SPARC/Full/Portage/ko "Handbook:SPARC/Full/Portage/ko")[파일 및 디렉터리](/wiki/Handbook:SPARC/Portage/Files/ko "Handbook:SPARC/Portage/Files/ko")[변수](/wiki/Handbook:SPARC/Portage/Variables/ko "Handbook:SPARC/Portage/Variables/ko")[소프트웨어 브랜치 함께 사용하기](/wiki/Handbook:SPARC/Portage/Branches/ko "Handbook:SPARC/Portage/Branches/ko")[추가 도구](/wiki/Handbook:SPARC/Portage/Tools/ko "Handbook:SPARC/Portage/Tools/ko")[꾸러미 저장소 개별 설정](/wiki/Handbook:SPARC/Portage/CustomTree/ko "Handbook:SPARC/Portage/CustomTree/ko")[고급 기능](/wiki/Handbook:SPARC/Portage/Advanced/ko "Handbook:SPARC/Portage/Advanced/ko")[네트워크 설정](/wiki/Handbook:SPARC/Full/Networking/ko "Handbook:SPARC/Full/Networking/ko")[시작하기](/wiki/Handbook:SPARC/Networking/Introduction/ko "Handbook:SPARC/Networking/Introduction/ko")[고급 설정](/wiki/Handbook:SPARC/Networking/Advanced/ko "Handbook:SPARC/Networking/Advanced/ko")[모듈러 네트워크](/wiki/Handbook:SPARC/Networking/Modular/ko "Handbook:SPARC/Networking/Modular/ko")[무선 네트워크](/wiki/Handbook:SPARC/Networking/Wireless/ko "Handbook:SPARC/Networking/Wireless/ko")[기능 추가](/wiki/Handbook:SPARC/Networking/Extending/ko "Handbook:SPARC/Networking/Extending/ko")[동적 관리](/wiki/Handbook:SPARC/Networking/Dynamic/ko "Handbook:SPARC/Networking/Dynamic/ko")

## Contents

- [1블록 장치 소개](#.EB.B8.94.EB.A1.9D_.EC.9E.A5.EC.B9.98_.EC.86.8C.EA.B0.9C)
  - [1.1블록 장치](#.EB.B8.94.EB.A1.9D_.EC.9E.A5.EC.B9.98)
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
- [4파일 시스템 만들기](#.ED.8C.8C.EC.9D.BC_.EC.8B.9C.EC.8A.A4.ED.85.9C_.EB.A7.8C.EB.93.A4.EA.B8.B0)
  - [4.1도입부](#.EB.8F.84.EC.9E.85.EB.B6.80)
  - [4.2파일 시스템](#.ED.8C.8C.EC.9D.BC_.EC.8B.9C.EC.8A.A4.ED.85.9C)
  - [4.3분할 영역에 파일 시스템 반영하기](#.EB.B6.84.ED.95.A0_.EC.98.81.EC.97.AD.EC.97.90_.ED.8C.8C.EC.9D.BC_.EC.8B.9C.EC.8A.A4.ED.85.9C_.EB.B0.98.EC.98.81.ED.95.98.EA.B8.B0)
    - [4.3.1EFI system partition filesystem](#EFI_system_partition_filesystem)
    - [4.3.2Legacy BIOS boot partition filesystem](#Legacy_BIOS_boot_partition_filesystem)
    - [4.3.3Small ext4 partitions](#Small_ext4_partitions)
  - [4.4스왑 분할 영역 활성화](#.EC.8A.A4.EC.99.91_.EB.B6.84.ED.95.A0_.EC.98.81.EC.97.AD_.ED.99.9C.EC.84.B1.ED.99.94)
- [5루트 분할 영역 마운트](#.EB.A3.A8.ED.8A.B8_.EB.B6.84.ED.95.A0_.EC.98.81.EC.97.AD_.EB.A7.88.EC.9A.B4.ED.8A.B8)

## 블록 장치 소개

### 블록 장치

리눅스 파일 시스템, 분할 영역, 블록 장치 등 젠투 리눅스 및 일반적인 리눅스 운영체제의 바람직한 디스크 측면의 양상을 살펴보도록 하겠습니다. 디스크와 파일 시스템의 입출력을 이해하고 나서, 젠투 리눅스 설치에 필요한 분할 영역과 파일 시스템을 설정하겠습니다.

시작에 앞서 블록 장치를 살펴보도록 하죠. 아마~도 리눅스 시스템에서 첫번째 드라이브로 표시하는 대부분 잘 알려진 블록 장치는 /dev/sda겠죠. SCSI와 직렬 ATA 드라이브 둘 다 /dev/sd\*와 같은 식으로 표시합니다. 게다가 커널의 libata 프레임워크에서는 IDE 드라이브도 마찬가지로 /dev/sd\*로 표시합니다. 이전 장치 프레임워크에서 첫번째 IDE 드라이브는 /dev/hda입니다.

The following table will help readers determine where to find a certain type of block device on the system:

Type of deviceDefault device handleEditorial notes and considerations
IDE, SATA, SAS, SCSI, or USB flash/dev/sdaFound on hardware from roughly 2007 until the present, this device handle is perhaps the most commonly used in Linux. These types of devices can be connected via the [SATA bus](https://en.wikipedia.org/wiki/Serial_ATA "wikipedia:Serial ATA"), [SCSI](https://en.wikipedia.org/wiki/SCSI "wikipedia:SCSI"), [USB](https://en.wikipedia.org/wiki/USB "wikipedia:USB") bus as block storage. As example, the first partition on the first SATA device is called /dev/sda1.
NVM Express (NVMe)/dev/nvme0n1The latest in solid state technology, [NVMe](https://en.wikipedia.org/wiki/NVM_Express "wikipedia:NVM Express") drives are connected to the PCI Express bus and have the fastest transfer block speeds on the market. Systems from around 2014 and newer may have support for NVMe hardware. The first partition on the first NVMe device is called /dev/nvme0n1p1.
MMC, eMMC, and SD/dev/mmcblk0[embedded MMC](https://en.wikipedia.org/wiki/MultiMediaCard#eMMC "wikipedia:MultiMediaCard") devices, SD cards, and other types of memory cards can be useful for data storage. That said, many systems may not permit booting from these types of devices. It is suggested to **not** use these devices for active Linux installations; rather consider using them to transfer files, which is their typical design intention. Alternatively this storage type could be useful for short-term file backups or snapshots.
VirtIO/dev/vda[Virtualized](/wiki/Virtualization "Virtualization") devices are found only within a [QEMU](/wiki/QEMU "QEMU") virtual emulator. `virtio_blk` driver provides this access to host's allocated drive space for this virtual machine. That said, it is a great way to try out Gentoo inside a virtual machine.

위에 나타낸 블록 장치는 디스크의 추상 인터페이스를 표현합니다. 사용자 프로그램은 블록 장치가 IDE가 됐든 SCSI가 됐든 뭐가 됐든지간에 신경쓰지 않고 디스크와 소통을 수행할 때 이 블록 장치를 사용할 수 있습니다. 프로그램에서는 디스크의 저장 공간에 대해, 연속적이며, 임의로 접근하는 512 바이트 블록의 모음으로 다룰 수 있습니다.

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
[BIOS boot partition](/wiki/Handbook:X86/Blocks/Disks#What_is_the_BIOS_boot_partition.3F.2Fko "Handbook:X86/Blocks/Disks")/dev/sda2(swap)
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

**중요**

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

## 파일 시스템 만들기

**경고**

When using an SSD or NVMe drive, it is wise to check for firmware upgrades. Some Intel SSDs in particular (600p and 6000p) require a firmware upgrade for [possible data corruption](https://bugzilla.redhat.com/show_bug.cgi?id=1402533) induced by XFS I/O usage patterns. The problem is at the firmware level and not any fault of the XFS filesystem. The smartctl utility can help check the device model and firmware version.

### 도입부

이제 분할 영역을 만들었고, 파일 시스템을 제 위치에 얹어놓을 차례입니다. 다음 절에서는 리눅스에서 지원하는 다양한 파일 시스템을 설명합니다. 어떤 파일 시스템을 사용할 지 이미 알고 있는 독자라면 [파티션에 파일 시스템 반영하기](#.ED.8C.8C.ED.8B.B0.EC.85.98.EC.97.90_.ED.8C.8C.EC.9D.BC_.EC.8B.9C.EC.8A.A4.ED.85.9C_.EB.B0.98.EC.98.81.ED.95.98.EA.B8.B0) 로 계속 진행할 수 있습니다. 그렇지 않으면 계속 읽어 내려가면서 쓸 수 있는 파일시스템이 어떤 종류가 있는지 알아보십시오.

### 파일 시스템

다양한 파일 시스템이 있습니다. 일부는 sparc 아키텍처에서 안정적입니다 - 중요한 분할 영역을 위해서라면 좀 더 시험적인 분할 영역을 선택하기 전에 파일 시스템과 지원 상태에 대한 내용을 좀 더 읽어보시는 것이 좋겠습니다.

btrfs스냅샷, 검사합을 통한 자체복구, 투명 압축, 하위 볼륨, 통합 RAID 같은 고급 기능을 제공하는 차세대 파일 시스템입니다. 일부 배포판은 이미 특별한 옵션으로 탑재했지만 실무에서 쓰기엔 준비가 미흡합니다. 파일 시스템이 깨지는 경우가 다반사입니다. 개발자들은 이전 버전에 문제가 있기 때문에 안전을 위해 최신 커널 버전을 사용하라고 합니다. 몇년 동안 이래왔고 무엇인가 바뀐다고 하면 너무 일찍 언급합니다. 깨지는 문제를 고친다고 하면 가끔 이전 커널에 있던 대로 돌아갑니다. 파일 시스템을 쓰려 한다면 위험을 감수하십시오!ext2검증된 리눅스 파일시스템이지만 메타데이터 저널링기능이 없습니다. 이는 시작시간의 파일시스템 검사루틴에서 조금 더 많은 시간소모를 할 수 있다는 의미입니다. 이제 일관성 검사를 더욱 빠르게 할 수 있고 비 저널링의 대체 수단으로써 일반적으로 더욱 선호하는 차세대 저널링 파일시스템의 상당한 선택요소가 있습니다. 저널링 파일시스템은 시스템을 시동하고 파일시스템에 비일관 상태가 발생했을 때 긴 지연시간을 줄입니다.ext3빠른 복구 기능을 제공하는 메타데이터 저널링을 제공하며, 게다가 전체 데이터와 정렬된 데이터 저널링과 같은 강화 저널링 모드도 지원하는 ext2 파일시스템의 저널링 버전입니다. 대부분의 모든 상황에서 고성능 동작이 가능한 HTree 색인을 사용합니다. 간단히 말해 ext3는 아주 좋은 믿을 수 있는 파일시스템입니다. ext3을 모든 목적의 모든 플랫폼 파일시스템으로 추천합니다.ext4ext3으로부터 갈라져 나와 성능을 향상시키고 디스크상 형식에 대해 적절한 수정을 가하여 용량 제한을 없애는 새로운 기능을 포함하여 만든 파일시스템입니다. 볼륨 하나의 크기를 1EB까지 늘릴 수 있고, 파일 최대 크기는 16TB가 될 수 있습니다. 기존의 ext2/3 비트맵 블록 할당 대신에 ext4는 대용량 파일 성능을 끌어올리고 단편화를 줄인 extents를 사용합니다. ext4는 디스크의 데이터 배치에 대해 최적화 할 더 많은 방법을 파일시스템 드라이버에 제공하는 좀 더 세련된 블록 할당 알고리즘(지연할당 및 다중블록 할당)을 제공합니다. ext4는 모든 목적의 모든 플랫폼의 파일 시스템에 추천합니다.f2fs플래시 지향 파일 시스템은 처음에 낸드 플래시 메모리에서 활용할 목적으로 삼성에서 만들었습니다. 2016년 2/4분기 시점에, 이 파일 시스템은 여전히 미완의 상태지만 젠투를 마이크로SD 카드, USB 드라이브, 기타 플래시 기반 저장 장치에 설치할 경우 괜찮은 선택입니다.JFSIBM의 고성능 저널링 파일시스템입니다. JFS는 다양한 상황속에서도 좋은 성능을 내는, 가볍고 빠르며 믿을 수 있는 B+트리 기반 파일시스템입니다.ReiserFS전반적으로 좋은 성능을 내며 특히 용량이 작은 수많은 파일들을 다룰 때 더 많은 CPU 사이클을 소비하는 경우 좋은 성능이 나는 B+트리 기반 저널링 파일시스템입니다. ReiserFS는 다른 파일시스템보다 덜 관리중인 것으로 보입니다.XFS견고한 기능 모음을 지니고 있으며 확장성에 있어 최적화 된 메타데이터 저널링 파일시스템입니다. XFS는 다양한 하드웨어 문제에 대해 그다지 관대하진 않은 것 같습니다.vfatFAT32로 알려진 vfat은 리눅스에서 지원하지만 권한 설정은 지원하지 않습니다. 여러 운영 체제간 상호 운용성을 목적으로(주로 마이크로소프트 윈도우) 활용하지만 일부 시스템 펌웨어(UEFI)용으로도 필요합니다.NTFS"New Technology" 파일 시스템은 마이크로 소프트의 대표 파일 시스템입니다. 위의 vfat과 비슷하게 BSD 또는 리눅스에서 필요한 권한 설정 또는 확장 속성을 저장하지 않기에 루트 파일 시스템으로 활용할 수 없습니다. 오직 마이크로소프트 윈도우와 상호 연동할 _때만_ 활용해야합니다(오직 이 경우에만 역점을 둠을 참고하십시오).

More extensive information on filesystems can be found in the community maintained [Filesystem article](/wiki/Filesystem/ko "Filesystem/ko").

### 분할 영역에 파일 시스템 반영하기

**참고**

Please make sure to emerge the relevant user space utilities package for the chosen filesystem before rebooting. There will be a reminder to do so near the end of the installation process.

분할 영역 또는 볼륨에 파일 시스템을 만들 때, 각 파일 시스템에서 사용할 수 있는 도구가 있습니다. 각 파일 시스템의 추가 정보를 살펴보려면 하단 표의 파일 시스템 이름을 누르십시오:

파일시스템
구성 명령
최소 CD 포함?
꾸러미
[btrfs](/wiki/Btrfs "Btrfs")mkfs.btrfs Yes
[sys-fs/btrfs-progs](https://packages.gentoo.org/packages/sys-fs/btrfs-progs)[ext2](/wiki/Ext2 "Ext2")mkfs.ext2 Yes
[sys-fs/e2fsprogs](https://packages.gentoo.org/packages/sys-fs/e2fsprogs)[ext3](/wiki/Ext3 "Ext3")mkfs.ext3 Yes
[sys-fs/e2fsprogs](https://packages.gentoo.org/packages/sys-fs/e2fsprogs)[ext4](/wiki/Ext4 "Ext4")mkfs.ext4 Yes
[sys-fs/e2fsprogs](https://packages.gentoo.org/packages/sys-fs/e2fsprogs)[f2fs](/wiki/F2fs "F2fs")mkfs.f2fs Yes
[sys-fs/f2fs-tools](https://packages.gentoo.org/packages/sys-fs/f2fs-tools)[jfs](/wiki/JFS "JFS")mkfs.jfs Yes
[sys-fs/jfsutils](https://packages.gentoo.org/packages/sys-fs/jfsutils)[reiserfs](/wiki/ReiserFS "ReiserFS")mkfs.reiserfs Yes
[sys-fs/reiserfsprogs](https://packages.gentoo.org/packages/sys-fs/reiserfsprogs)[xfs](/wiki/XFS "XFS")mkfs.xfs Yes
[sys-fs/xfsprogs](https://packages.gentoo.org/packages/sys-fs/xfsprogs)[vfat](/wiki/FAT "FAT")mkfs.vfat Yes
[sys-fs/dosfstools](https://packages.gentoo.org/packages/sys-fs/dosfstools)[NTFS](/wiki/NTFS "NTFS")mkfs.ntfs Yes
[sys-fs/ntfs3g](https://packages.gentoo.org/packages/sys-fs/ntfs3g)

\|}

**중요**

The handbook recommends new partitions as part of the installation process, but it is important to note running any mkfs command will erase any data contained within the partition. When necessary, ensure any data that exists within is appropriately backed up before creating a new filesystem.

예를 들어, 예제 분할 영역 구조와 같이 ext2 형식의 ext4 형식의 루트 분할 영역 (/dev/sda1)을 취하려면, 다음 명령을 사용할 수 있습니다:

`root #` `mkfs.ext4 /dev/sda1`

{{#ifeq: 0\|1\|

#### EFI system partition filesystem

The EFI system partition () must be formatted as FAT32:

#### Legacy BIOS boot partition filesystem

Systems booting via legacy BIOS with a MBR/DOS disklabel can use any filesystem format supported by the bootloader.

For example, to format with XFS:

`root #` `mkfs.xfs -c options=/usr/share/xfsprogs/mkfs/lts_6.18.conf`

#### Small ext4 partitions

(8GB 이하의) 작은 분할 영역에서 ext2, ext3, ext4 를 사용한다면, 충분한 inode 갯수를 예약할 적당한 옵션으로 파일 시스템을 만들어야합니다. mke2fs(mkfs.ext2)에서는 "아이노드 당 바이트" 설정을 사용하여 파일 시스템에서 보유할 아이노드 갯수를 계산합니다. 작은 분할 영역일수록 아이노드 갯수를 늘리는 것이 좋습니다.

`root #` `mkfs.ext2 -T small /dev/<device>`

각 16kB 영역을 하나의 4kB 영역으로 줄이는 "아이노드 당 바이트"로 주어진 파일 시스템의 아이노드 갯수를 네 배로 뻥튀기(?)합니다. 비율값을 부여하여 속성을 조절할 수 있습니다:

### 스왑 분할 영역 활성화

mkswap은 스왑 분할 영역을 초기화하는 명령입니다:

`root #` `mkswap /dev/sda2`

**참고**

Installations which were previously started, but did not finish the installation process can resume the installation from this point in the handbook. Use this link as the permalink: [Resumed installations start here](/wiki/Handbook:SPARC/Installation/Disks/ko#Resumed_installations_start_here "Handbook:SPARC/Installation/Disks/ko").

스왑 분할 영역을 활성화하려면, swapon 명령을 사용하십시오:

`root #` `swapon /dev/sda2`

This 'activation' step is only necessary because the swap partition is newly created within the live environment. Once the system has been rebooted, as long as the swap partition is properly defined within fstab or other mount mechanism, swap space will activate automatically.

## 루트 분할 영역 마운트

Certain live environments may be missing the suggested mount point for Gentoo's root partition (/mnt/gentoo), or mount points for additional partitions created in the partitioning section:

`root #` `mkdir --parents /mnt/gentoo`

Continue creating additional mount points necessary for any additional (custom) partition(s) created during previous steps by using the mkdir command.

이제 분할 영역을 초기화했고 파일 시스템을 넣었으므로 분할 영역을 마운트할 차례입니다. mount 명령을 사용하지만 만들어놓은 모든 분할 영역에 대해 마운트 디렉터리를 만들 필요는 없다는 사실을 잊지 마십시오. 예제를 통해 우리는 루트 분할 영역을 마운트하겠습니다:

Mount the root partition:

`root #` `mount /dev/sda1 /mnt/gentoo`

Continue mounting additional (custom) partitions as necessary using the mount command.

**참고**

/tmp/를 따로 나눈 분할 영역에 두어야 한다면, 마운트하기 전에 퍼미션을 바꾸었는지 확인하십시오:

`root #` `chmod 1777 /mnt/gentoo/tmp`

이 설정은 /var/tmp에도 적용 유지합니다.

지침을 따르고 나면 proc 파일 시스템(커널 가상 인터페이스)와 다른 커널 의사 파일 시스템을 마운트합니다. 그러나 우선 [젠투 설치 파일을 설치](/wiki/Handbook:SPARC/Installation/Stage/ko "Handbook:SPARC/Installation/Stage/ko") 하겠습니다.

[← 네트워크 설정](/wiki/Handbook:SPARC/Installation/Networking/ko "이전 내용") [처음](/wiki/Handbook:SPARC/ko "Handbook:SPARC/ko") [젠투 설치 파일 설치 →](/wiki/Handbook:SPARC/Installation/Stage/ko "다음 내용")