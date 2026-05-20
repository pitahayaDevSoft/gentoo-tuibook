# Disks

Other languages:

- [Deutsch](/wiki/Handbook:AMD64/Installation/Disks/de "Handbook:AMD64/Installation/Disks/de (100% translated)")
- [English](/wiki/Handbook:AMD64/Installation/Disks "Handbook:AMD64/Installation/Disks (100% translated)")
- [Türkçe](/wiki/Handbook:AMD64/Installation/Disks/tr "Handbook:AMD64/Installation/Disks/tr (0% translated)")
- [español](/wiki/Handbook:AMD64/Installation/Disks/es "Handbook:AMD64/Installation/Disks/es (100% translated)")
- [français](/wiki/Handbook:AMD64/Installation/Disks/fr "Handbook:AMD64/Installation/Disks/fr (100% translated)")
- [italiano](/wiki/Handbook:AMD64/Installation/Disks/it "Handbook:AMD64/Installation/Disks/it (100% translated)")
- [magyar](/wiki/Handbook:AMD64/Installation/Disks/hu "Handbook:AMD64/Installation/Disks/hu (100% translated)")
- [polski](/wiki/Handbook:AMD64/Installation/Disks/pl "Handbook:AMD64/Installation/Disks/pl (100% translated)")
- [português do Brasil](/wiki/Handbook:AMD64/Installation/Disks/pt-br "Handbook:AMD64/Installation/Disks/pt-br (100% translated)")
- [čeština](/wiki/Handbook:AMD64/Installation/Disks/cs "Handbook:AMD64/Installation/Disks/cs (100% translated)")
- [русский](/wiki/Handbook:AMD64/Installation/Disks/ru "Handbook:AMD64/Installation/Disks/ru (100% translated)")
- العربية
- [தமிழ்](/wiki/Handbook:AMD64/Installation/Disks/ta "Handbook:AMD64/Installation/Disks/ta (100% translated)")
- [中文（中国大陆）‎](/wiki/Handbook:AMD64/Installation/Disks/zh-cn "Handbook:AMD64/Installation/Disks/zh-cn (100% translated)")
- [日本語](/wiki/Handbook:AMD64/Installation/Disks/ja "Handbook:AMD64/Installation/Disks/ja (100% translated)")
- [한국어](/wiki/Handbook:AMD64/Installation/Disks/ko "Handbook:AMD64/Installation/Disks/ko (100% translated)")

[دليل AMD64](/wiki/Handbook:AMD64 "Handbook:AMD64")[التنصيب](/wiki/Handbook:AMD64/Full/Installation "Handbook:AMD64/Full/Installation")[عن التنصيب](/wiki/Handbook:AMD64/Installation/About/ar "Handbook:AMD64/Installation/About/ar")[اختيار وسيط التنصيب](/wiki/Handbook:AMD64/Installation/Media/ar "Handbook:AMD64/Installation/Media/ar")[إعداد الشبكة](/wiki/Handbook:AMD64/Installation/Networking/ar "Handbook:AMD64/Installation/Networking/ar")تحضير الأقراص[ملف المرحلة](/wiki/Handbook:AMD64/Installation/Stage "Handbook:AMD64/Installation/Stage")[تنصيب النظام الأساسي](/wiki/Handbook:AMD64/Installation/Base/ar "Handbook:AMD64/Installation/Base/ar")[إعداد النواة](/wiki/Handbook:AMD64/Installation/Kernel "Handbook:AMD64/Installation/Kernel")[إعداد النظام](/wiki/Handbook:AMD64/Installation/System "Handbook:AMD64/Installation/System")[تنصيب الأدوات](/wiki/Handbook:AMD64/Installation/Tools "Handbook:AMD64/Installation/Tools")[إعداد محمل الإقلاع](/wiki/Handbook:AMD64/Installation/Bootloader "Handbook:AMD64/Installation/Bootloader")[إنهاء التنصيب](/wiki/Handbook:AMD64/Installation/Finalizing "Handbook:AMD64/Installation/Finalizing")[العمل مع جنتو](/wiki/Handbook:AMD64/Full/Working "Handbook:AMD64/Full/Working")[مقدمة إلى بورتج](/wiki/Handbook:AMD64/Working/Portage "Handbook:AMD64/Working/Portage")[وسوم الاستخدام (USE)](/wiki/Handbook:AMD64/Working/USE "Handbook:AMD64/Working/USE")[ميزات بورتج](/wiki/Handbook:AMD64/Working/Features "Handbook:AMD64/Working/Features")[نظام سكربتات البدء](/wiki/Handbook:AMD64/Working/Initscripts "Handbook:AMD64/Working/Initscripts")[متغيرات البيئة](/wiki/Handbook:AMD64/Working/EnvVar "Handbook:AMD64/Working/EnvVar")[العمل مع بورتج](/wiki/Handbook:AMD64/Full/Portage "Handbook:AMD64/Full/Portage")[الملفات والدلائل](/wiki/Handbook:AMD64/Portage/Files "Handbook:AMD64/Portage/Files")[المتغيرات](/wiki/Handbook:AMD64/Portage/Variables "Handbook:AMD64/Portage/Variables")[خلط فروع البرمجيات](/wiki/Handbook:AMD64/Portage/Branches "Handbook:AMD64/Portage/Branches")[أدوات إضافية](/wiki/Handbook:AMD64/Portage/Tools "Handbook:AMD64/Portage/Tools")[مستودع حزم مخصص](/wiki/Handbook:AMD64/Portage/CustomTree "Handbook:AMD64/Portage/CustomTree")[ميزات متقدمة](/wiki/Handbook:AMD64/Portage/Advanced "Handbook:AMD64/Portage/Advanced")[إعداد الشبكة في OpenRC](/wiki/Handbook:AMD64/Full/Networking "Handbook:AMD64/Full/Networking")[البداية](/wiki/Handbook:AMD64/Networking/Introduction "Handbook:AMD64/Networking/Introduction")[الإعداد المتقدم](/wiki/Handbook:AMD64/Networking/Advanced "Handbook:AMD64/Networking/Advanced")[الشبكات التركيبية](/wiki/Handbook:AMD64/Networking/Modular "Handbook:AMD64/Networking/Modular")[اللاسلكي](/wiki/Handbook:AMD64/Networking/Wireless "Handbook:AMD64/Networking/Wireless")[إضافة الوظائف](/wiki/Handbook:AMD64/Networking/Extending "Handbook:AMD64/Networking/Extending")[الإدارة الديناميكية](/wiki/Handbook:AMD64/Networking/Dynamic "Handbook:AMD64/Networking/Dynamic")

## Contents

- [١مقدمة عن الأجهزة الكتلية](#.D9.85.D9.82.D8.AF.D9.85.D8.A9_.D8.B9.D9.86_.D8.A7.D9.84.D8.A3.D8.AC.D9.87.D8.B2.D8.A9_.D8.A7.D9.84.D9.83.D8.AA.D9.84.D9.8A.D8.A9)
  - [١.١الأجهزة الكتلية](#.D8.A7.D9.84.D8.A3.D8.AC.D9.87.D8.B2.D8.A9_.D8.A7.D9.84.D9.83.D8.AA.D9.84.D9.8A.D8.A9)
  - [١.٢Partition tables](#Partition_tables)
    - [١.٢.١GUID Partition Table (GPT)](#GUID_Partition_Table_.28GPT.29)
    - [١.٢.٢Master boot record (MBR) or DOS boot sector](#Master_boot_record_.28MBR.29_or_DOS_boot_sector)
  - [١.٣Advanced storage](#Advanced_storage)
  - [١.٤Default partitioning scheme](#Default_partitioning_scheme)
- [٢Designing a partition scheme](#Designing_a_partition_scheme)
  - [٢.١How many partitions and how big?](#How_many_partitions_and_how_big.3F)
  - [٢.٢What about swap space?](#What_about_swap_space.3F)
  - [٢.٣What is the EFI System Partition (ESP)?](#What_is_the_EFI_System_Partition_.28ESP.29.3F)
  - [٢.٤What is the BIOS boot partition?](#What_is_the_BIOS_boot_partition.3F)
- [٣Partitioning the disk with GPT for UEFI](#Partitioning_the_disk_with_GPT_for_UEFI)
  - [٣.١Viewing the current partition layout](#Viewing_the_current_partition_layout)
  - [٣.٢Creating a new disklabel / removing all partitions](#Creating_a_new_disklabel_.2F_removing_all_partitions)
  - [٣.٣Creating the EFI System Partition (ESP)](#Creating_the_EFI_System_Partition_.28ESP.29)
  - [٣.٤Creating the swap partition](#Creating_the_swap_partition)
  - [٣.٥Creating the root partition](#Creating_the_root_partition)
  - [٣.٦Saving the partition layout](#Saving_the_partition_layout)
- [٤Partitioning the disk with MBR for BIOS / legacy boot](#Partitioning_the_disk_with_MBR_for_BIOS_.2F_legacy_boot)
  - [٤.١Viewing the current partition layout](#Viewing_the_current_partition_layout_2)
  - [٤.٢Creating a new disklabel / removing all partitions](#Creating_a_new_disklabel_.2F_removing_all_partitions_2)
  - [٤.٣Creating the boot partition](#Creating_the_boot_partition)
  - [٤.٤Creating the swap partition](#Creating_the_swap_partition_2)
  - [٤.٥Creating the root partition](#Creating_the_root_partition_2)
  - [٤.٦Saving the partition layout](#Saving_the_partition_layout_2)
- [٥إنشاء أنظمة الملفات](#.D8.A5.D9.86.D8.B4.D8.A7.D8.A1_.D8.A3.D9.86.D8.B8.D9.85.D8.A9_.D8.A7.D9.84.D9.85.D9.84.D9.81.D8.A7.D8.AA)
  - [٥.١مقدمة](#.D9.85.D9.82.D8.AF.D9.85.D8.A9)
  - [٥.٢نظم الملفات](#.D9.86.D8.B8.D9.85_.D8.A7.D9.84.D9.85.D9.84.D9.81.D8.A7.D8.AA)
  - [٥.٣تطبيق نظام ملفات على قسم](#.D8.AA.D8.B7.D8.A8.D9.8A.D9.82_.D9.86.D8.B8.D8.A7.D9.85_.D9.85.D9.84.D9.81.D8.A7.D8.AA_.D8.B9.D9.84.D9.89_.D9.82.D8.B3.D9.85)
    - [٥.٣.١نظام ملفات قسم نظام EFI](#.D9.86.D8.B8.D8.A7.D9.85_.D9.85.D9.84.D9.81.D8.A7.D8.AA_.D9.82.D8.B3.D9.85_.D9.86.D8.B8.D8.A7.D9.85_EFI)
    - [٥.٣.٢نظام ملفات قسم إقلاع BIOS العتيق](#.D9.86.D8.B8.D8.A7.D9.85_.D9.85.D9.84.D9.81.D8.A7.D8.AA_.D9.82.D8.B3.D9.85_.D8.A5.D9.82.D9.84.D8.A7.D8.B9_BIOS_.D8.A7.D9.84.D8.B9.D8.AA.D9.8A.D9.82)
    - [٥.٣.٣أقسام ext4 الصغيرة](#.D8.A3.D9.82.D8.B3.D8.A7.D9.85_ext4_.D8.A7.D9.84.D8.B5.D8.BA.D9.8A.D8.B1.D8.A9)
  - [٥.٤تفعيل قسم الإبدال](#.D8.AA.D9.81.D8.B9.D9.8A.D9.84_.D9.82.D8.B3.D9.85_.D8.A7.D9.84.D8.A5.D8.A8.D8.AF.D8.A7.D9.84)
- [٦وصل قسم الجذر](#.D9.88.D8.B5.D9.84_.D9.82.D8.B3.D9.85_.D8.A7.D9.84.D8.AC.D8.B0.D8.B1)

## مقدمة عن الأجهزة الكتلية

### الأجهزة الكتلية

فلنلقِ نظرة فاحصة على الجوانب المتعلقة بالأقراص في جنتو لينكس ولينكس عمومًا، بما في ذلك الأجهزة الكتلية، والأقسام، وأنظمة ملفات لينكس. بمجرد استيعاب تفاصيل الأقراص وخباياها، يمكن حينئذٍ إنشاء الأقسام وأنظمة الملفات اللازمة للتنصيب.

بدايةً، لنلقِ نظرة على الأجهزة الكتلية. تُصنف محركات الأقراص من نوع SCSI وSerial ATA تحت معرفات أجهزة مثل: /dev/sda و/dev/sdb و/dev/sdc، وما إلى ذلك. أما في الأجهزة الأحدث، فإن أقراص الحالة الصلبة NVMe القائمة على ناقل PCI Express تحمل معرفات أجهزة مثل /dev/nvme0n1 و/dev/nvme0n2، وهكذا.

سيساعد الجدول التالي القراء في تحديد مكان العثور على نوع معين من الأجهزة الكتلية في النظام:

نوع الجهازمعرف الجهاز المبدئيملاحظات واعتبارات تحريرية
IDE، أو SATA، أو SAS، أو SCSI، أو ذاكرة USB وميضية/dev/sdaيوجد هذا المعرف في العتاد المنتج منذ عام 2007 تقريبًا وحتى الوقت الحاضر، وهو ربما المعرف الأكثر استخدامًا في لينكس. يمكن توصيل هذه الأنواع من الأجهزة عبر [ناقل SATA](https://en.wikipedia.org/wiki/ar:%D8%B3%D8%A7%D8%AA%D8%A7 "wikipedia:ar:ساتا")، أو [SCSI](https://en.wikipedia.org/wiki/ar:%D9%88%D8%A7%D8%AC%D9%87%D8%A9_%D8%A7%D9%84%D9%86%D8%B8%D8%A7%D9%85_%D8%A7%D9%84%D8%AD%D8%A7%D8%B3%D9%88%D8%A8%D9%8A_%D8%A7%D9%84%D8%B5%D8%BA%D9%8A%D8%B1 "wikipedia:ar:واجهة النظام الحاسوبي الصغير")، أو ناقل [USB](https://en.wikipedia.org/wiki/ar:%D8%A7%D9%84%D9%85%D8%B3%D8%B1%D9%89_%D8%A7%D9%84%D8%AA%D8%B3%D9%84%D8%B3%D9%84%D9%8A_%D8%A7%D9%84%D8%B9%D8%A7%D9%85 "wikipedia:ar:المسرى التسلسلي العام") كوحدات تخزين كتلية. على سبيل المثال، يسمى القسم الأول على أول جهاز SATA ب‍ /dev/sda1.
NVM Express (NVMe)/dev/nvme0n1تمثل محركات [NVMe](https://en.wikipedia.org/wiki/ar:%D8%A5%D9%86_%D9%81%D9%8A_%D8%A5%D9%85_%D8%A5%D9%83%D8%B3%D8%A8%D8%B1%D8%B3 "wikipedia:ar:إن في إم إكسبرس") أحدث تقنيات الحالة الصلبة، وتتصل بناقل PCI Express وتتمتع بأسرع سرعات نقل كتلية في السوق. الأنظمة المنتجة منذ عام 2014 والأحدث قد تدعم عتاد NVMe. يسمى القسم الأول على أول جهاز NVMe ب‍ /dev/nvme0n1p1.
MMC، وeMMC، وSD/dev/mmcblk0يمكن أن تكون أجهزة [eMMC المدمجة](https://en.wikipedia.org/wiki/MultiMediaCard#eMMC "wikipedia:MultiMediaCard") وبطاقات SD وأنواع بطاقات الذاكرة الأخرى مفيدة لتخزين البيانات. ومع ذلك، قد لا تسمح العديد من الأنظمة بالإقلاع من هذه الأنواع من الأجهزة. ويُنصح ب **عدم** استخدام هذه الأجهزة لتنصيبات لينكس النشطة، بل يفضل استخدامها لنقل الملفات، وهو الغرض الأساسي من تصميمها. وبدلاً من ذلك، قد يكون هذا النوع من التخزين مفيدًا للنسخ الاحتياطي قصير المدى أو اللقطات (snapshots).
VirtIO/dev/vdaلا توجد الأجهزة [الافتراضية](/wiki/Virtualization "Virtualization") إلا داخل محاكي [QEMU](/wiki/QEMU "QEMU") الافتراضي. يوفر المحرك `virtio_blk` إمكانية الوصول إلى مساحة القرص المخصصة من المضيف لهذا الجهاز الافتراضي. ومع ذلك، فهي وسيلة رائعة لتجربة جنتو داخل جهاز افتراضي.

تمثل الأجهزة الكتلية المذكورة أعلاه واجهة تجريدية للقرص. ويمكن لبرامج المستخدم استخدام هذه الأجهزة الكتلية للتفاعل مع القرص دون القلق بشأن ما إذا كانت المحركات من نوع SATA أو SCSI أو غير ذلك. إذ يمكن للبرنامج ببساطة التعامل مع التخزين على القرص كمجموعة من الكتل المتجاورة بحجم 4096 بايت (4K) والتي يمكن الوصول إليها عشوائيًا.

### Partition tables

Although it is theoretically possible to use a raw, unpartitioned disk to house a Linux system (when creating a btrfs RAID for example), this is almost never done in practice. Instead, disk block devices are split up into smaller, more manageable block devices. On **amd64** systems, these are called partitions. There are currently two standard partitioning technologies in use: MBR (sometimes also called DOS disklabel) and GPT; these are tied to the two boot process types: legacy BIOS boot and UEFI.

#### GUID Partition Table (GPT)

The _GUID Partition Table (GPT)_ setup (also called GPT disklabel) uses 64-bit identifiers for the partitions. The location in which it stores the partition information is much bigger than the 512 bytes of the MBR partition table (DOS disklabel), which means there is practically no limit on the number of partitions for a GPT disk. Also, the maximum partition size is much larger (almost 8 ZiB -- yes, zebibytes).

When a system's software interface between the operating system and firmware is UEFI (instead of BIOS), GPT is almost mandatory as compatibility issues will arise with DOS disklabel.

GPT also takes advantage of checksumming and redundancy. It carries CRC32 checksums to detect errors in the header and partition tables and has a backup GPT at the end of the disk. This backup table can be used to recover damage of the primary GPT near the beginning of the disk.

**Important**

There are a few caveats regarding GPT:

- Using GPT on a BIOS-based computer works, but the user won't be able to dual-boot with a Microsoft Windows operating system, since Microsoft Windows refuses to boot from a GPT partition when in BIOS mode.
- Some buggy (old) motherboard firmware configured to boot in BIOS/CSM/legacy mode might also have problems with booting from GPT labeled disks.

#### Master boot record (MBR) or DOS boot sector

The _[Master boot record](https://en.wikipedia.org/wiki/Master_boot_record "wikipedia:Master boot record")_ boot sector (also called DOS boot sector, DOS disklabel, and - more recently, in contrast to GPT/UEFI setups - legacy BIOS boot) was first introduced in 1983 with PC DOS 2.x. MBR uses 32-bit identifiers for the start sector and length of the partitions, and supports three partition types: primary, extended, and logical. Primary partitions have their information stored in the master boot record itself - a very small (usually 512 bytes) location at the very beginning of a disk. Due to this small space, only four primary partitions are supported (for instance, /dev/sda1 to /dev/sda4).

In order to support more partitions, one of the primary partitions in the MBR can be marked as an _extended_ partition. This partition can then contain additional logical partitions (partitions within a partition).

**Important**

Although still supported by most motherboard manufacturers, MBR boot sectors and their associated partitioning limitations are considered legacy. Unless working with hardware that is pre-2010, it best to partition a disk with [GUID Partition Table](https://en.wikipedia.org/wiki/GUID_Partition_Table "wikipedia:GUID Partition Table"). Readers who must proceed with setup type should knowingly acknowledge the following information:

- Most post-2010 motherboards consider using MBR boot sectors a legacy (supported, but not ideal) boot mode.
- Due to using 32-bit identifiers, partition tables in the MBR cannot address storage space that is larger than 2 TiBs in size.
- Unless an extended partition is created, MBR supports a maximum of four partitions.
- This setup does not provide a backup boot sector, so if something overwrites the partition table, all partition information will be lost.

That said, MBR and legacy BIOS boot may still used in virtualized cloud environments such as AWS.

The Handbook authors suggest using [GPT](#GPT) whenever possible for Gentoo installations.

### Advanced storage

The official Gentoo boot media provides support for many advanced filesystem and tool setups, which offer more flexible changes, snapshots and some cases more caching abitles

Although usage is not covered in the handbook, below is a list helpful guides to get the system running:

- [Btrfs](/wiki/Btrfs/Native_System_Root_Guide "Btrfs/Native System Root Guide")
- [ZFS](/wiki/ZFS/rootfs "ZFS/rootfs")
- [LVM (Logical Volume Manager)](/wiki/LVM "LVM")

Disk encryption is also handled in the same manner:

- [Rootfs encryption](/wiki/Rootfs_encryption "Rootfs encryption")

### Default partitioning scheme

Throughout the remainder of the handbook, we will discuss and explain two cases:

1. UEFI firmware with GUID Partition Table (GPT) disk.
2. MBR DOS/legacy BIOS firmware with a MBR partition table disk.

While it is possible to mix and match boot types with certain motherboard firmware, mixing goes beyond the intention of the handbook. As previously stated, it is strongly recommended for installations on modern hardware to use UEFI boot with a GPT disklabel disk.

The following partitioning scheme will be used as a simple example layout.

**Important**

The first row of the following table contains exclusive information for _**either**_ a GPT disklabel _**or**_ a MBR DOS/legacy BIOS disklabel. When in doubt, proceed with GPT, since **amd64** machines manufactured after the year 2010 generally support UEFI firmware and GPT boot sector.

Partition
Filesystem
Size
Description
/dev/sda1fat32 File system required for the EFI System Partition, which is always associated with a GPT disklabel.1 GiB
EFI System Partition details. _Applicable to system firmware supporting an UEFI implementation. This is typically the case for systems manufactured around the year 2010 to the present._xfs Recommended file system for the boot partition of a MBR partition table, which is used in conjunction with older firmware limited to the DOS/legacy BIOS disklabel.MBR DOS/legacy BIOS boot partition details. _Applicable to legacy BIOS machine firmware. Systems of this type were typically manufactured <u>before</u> the year 2010 and have generally phased out of production._/dev/sda2linux-swap
RAM size \* 2
Swap partition details.
/dev/sda3xfs
Remainder of the disk The selected _**profile**_, additional _**partitions**_ (optional), and system _**purpose**_ add complexities for appropriately sizing the rootfs, therefore the Handbook authors cannot offer a 'one-size-fits-all' suggestion for the rootfs partition.</br></br> When Gentoo is the only operating system using the disk, selecting the remainder of the disk is the safest and suggested choice.Root partition details.

If this suffices as information, the advanced reader can directly skip ahead to the actual partitioning.

Both fdisk and parted are partitioning utilities included within the official Gentoo live image environments. fdisk is well known, stable, and handles both MBR and GPT disks. parted was one of the first Linux block device management utilities to support GPT partitions. It can be used as an alternative to fdisk if the reader prefers, however the handbook will only provide instructions for fdisk, since it is commonly available on most Linux environments.

Before going to the creation instructions, the first set of sections will describe in more detail how partitioning schemes can be created and mention some common pitfalls.

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

**Note**

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

**Tip**

With the rise in popularity with [zram](/wiki/Zram "Zram"), it should be noted that Gentoo is more likely to suffer with performance hits, as compiling uses large amounts of RAM, and programs can't be directly compiled in swap. It's usually better to create a smaller zram as a cache swap area with higher priority over the disk based swap of around 512MB, if a user really wants to use zram to benefit their desktop needs and not harm compile times.

### What is the EFI System Partition (ESP)?

When installing Gentoo on a system that uses UEFI to boot the operating system (instead of BIOS) it is essential that an EFI System Partition (ESP) is created. The instructions below contain the necessary pointers to correctly handle this operation. **The EFI system partition is not required when booting in BIOS/Legacy mode.**

The ESP _must_ be a FAT variant (sometimes shown as _vfat_ on Linux systems). The official [UEFI specification](http://www.uefi.org/sites/default/files/resources/UEFI%202_5.pdf) denotes FAT12, 16, or 32 filesystems will be recognized by the UEFI firmware, although FAT32 is recommended for the ESP. After partitioning, format the ESP accordingly:

`root #` `mkfs.fat -F 32 /dev/sda1`

**Important**

If the ESP is not formatted with a FAT variant, the system's UEFI firmware will not find the bootloader (or Linux kernel) and will most likely be unable to boot the system!

### What is the BIOS boot partition?

A _BIOS boot partition_ is a very small (1 to 2 MB) partition in which boot loaders like GRUB2 can put additional data that doesn't fit in the allocated storage. It is only needed when a disk is formatted with a GPT disklabel, but the system's firmware will be booting via GRUB2 in legacy BIOS/MBR DOS boot mode. **It is _not required_ when booting in EFI/UEFI mode, and also _not required_ when using a MBR/Legacy DOS disklabel.** A _BIOS boot partition_ will not be used in this guide.

## Partitioning the disk with GPT for UEFI

The following parts explain how to create an example partition layout for a single GPT disk device which will conform to the UEFI Specification and Discoverable Partitions Specification (DPS). DPS is a specification provided as part of the Linux Userspace API (UAPI) Group Specification and is recommended, but entirely optional. The specifications are implemented using the fdisk utility, which is part of the [sys-apps/util-linux](https://packages.gentoo.org/packages/sys-apps/util-linux) package.

The table provides a recommended default for a trivial Gentoo installation. Additional partitions can be added according to personal preference or system design goals.

Device path (sysfs)
Mount point
File system
DPS UUID (Type-UUID)
Description
/dev/sda1/efivfat
c12a7328-f81f-11d2-ba4b-00a0c93ec93b
EFI system partition (ESP) details.
/dev/sda2N/A. Swap is not mounted to the filesystem like a device file.swap
0657fd6d-a4ab-43c4-84e5-0933c84b4f4f
Swap partition details.
/dev/sda3/xfs
4f68bce3-e8cd-4db1-96e7-fbcaf984b709
Root partition details.

### Viewing the current partition layout

fdisk is a popular and powerful tool to split a disk into partitions. Fire up fdisk against the disk (in our example, we use /dev/sda):

`root #` `fdisk /dev/sda`

Use the `p` key to display the disk's current partition configuration:

`Command (m for help):` `p`

```
Disk /dev/sda: 931.51 GiB, 1000204886016 bytes, 1953525168 sectors
Disk model: HGST HTS721010A9
Units: sectors of 1 * 512 = 512 bytes
Sector size (logical/physical): 512 bytes / 4096 bytes
I/O size (minimum/optimal): 4096 bytes / 4096 bytes
Disklabel type: gpt
Disk identifier: 3E56EE74-0571-462B-A992-9872E3855D75

<!--T:236-->
Device        Start        End    Sectors   Size Type
/dev/sda1      2048    2099199    2097152     1G EFI System
/dev/sda2   2099200   10487807    8388608     4G Linux swap
/dev/sda3  10487808 1953523711 1943035904 926.5G Linux root (x86-64)

```

This particular disk was configured to house two Linux filesystems (each with a corresponding partition listed as "Linux") as well as a swap partition (listed as "Linux swap").

### Creating a new disklabel / removing all partitions

Pressing the `g` key will instantly remove all existing disk partitions and create a new GPT disklabel:

`Command (m for help):` `g`

```
Created a new GPT disklabel (GUID: 3E56EE74-0571-462B-A992-9872E3855D75).

```

Alternatively, to keep an existing GPT disklabel (see the output of `p` above), consider removing the existing partitions one by one from the disk. Press `d` to delete a partition. For instance, to delete an existing /dev/sda1:

`Command (m for help):` `d`

```
Partition number (1-4): 1

```

The partition has now been scheduled for deletion. It will no longer show up when printing the list of partitions ( `p`, but it will not be erased until the changes have been saved. This allows users to abort the operation if a mistake was made - in that case, press `q` immediately and hit `Enter` and the partition will not be deleted.

Repeatedly press `p` to print out a partition listing and then press `d` and the number of the partition to delete it. Eventually, the partition table will be empty:

`Command (m for help):` `p`

```
Disk /dev/sda: 931.51 GiB, 1000204886016 bytes, 1953525168 sectors
Disk model: HGST HTS721010A9
Units: sectors of 1 * 512 = 512 bytes
Sector size (logical/physical): 512 bytes / 4096 bytes
I/O size (minimum/optimal): 4096 bytes / 4096 bytes
Disklabel type: gpt
Disk identifier: 3E56EE74-0571-462B-A992-9872E3855D75

```

Now that the in-memory partition table is empty, we're ready to create the partitions.

### Creating the EFI System Partition (ESP)

**Note**

A smaller ESP is possible but not recommended, especially given it may be shared with other OSes.

First create a small EFI system partition, which will also be mounted as /efi. Type `n` to create a new partition, followed by `1` to select the first partition. When prompted for the first sector, make sure it starts from 2048 (which may be needed for the boot loader) and hit `Enter`. When prompted for the last sector, type +1G to create a partition 1 gibibyte in size:

`Command (m for help):` `n`

```
Partition number (1-128, default 1): 1
First sector (2048-1953525134, default 2048):
Last sector, +/-sectors or +/-size{K,M,G,T,P} (2048-1953525134, default 1953523711): +1G

Created a new partition 1 of type 'Linux filesystem' and of size 1 GiB.
Partition #1 contains a vfat signature.

<!--T:247-->
Do you want to remove the signature? [Y]es/[N]o: Y
The signature will be removed by a write command.

```

Mark the partition as an EFI system partition:

`Command (m for help):` `t`

```
Selected partition 1
Partition type or alias (type L to list all): 1
Changed type of partition 'Linux filesystem' to 'EFI System'.

```

### Creating the swap partition

Next, to create the swap partition, press `n` to create a new partition, then press `2` to create the second partition, /dev/sda2. When prompted for the first sector, hit `Enter`. When prompted for the last sector, type +4G (or any other size needed for the swap space) to create a partition 4 GiB in size.

`Command (m for help):` `n`

```
Partition number (2-128, default 2):
First sector (2099200-1953525134, default 2099200):
Last sector, +/-sectors or +/-size{K,M,G,T,P} (2099200-1953525134, default 1953523711): +4G

Created a new partition 2 of type 'Linux filesystem' and of size 4 GiB.

```

After this, press `t` to set the partition type, `2` to select the partition just created and then type in _19_ to set the partition type to "Linux Swap".

`Command (m for help):` `t`

```
Partition number (1,2, default 2): 2
Partition type or alias (type L to list all): 19

Changed type of partition 'Linux filesystem' to 'Linux swap'.

```

### Creating the root partition

Finally, to create the root partition, press `n` to create a new partition, and then press `3` to create the third partition: /dev/sda3. When prompted for the first sector, press `Enter`. When prompted for the last sector, hit `Enter` to create a partition that takes up the rest of the remaining space on the disk.

`Command (m for help):` `n`

```
Partition number (3-128, default 3): 3
First sector (10487808-1953525134, default 10487808):
Last sector, +/-sectors or +/-size{K,M,G,T,P} (10487808-1953525134, default 1953523711):

<!--T:238-->
Created a new partition 3 of type 'Linux filesystem' and of size 926.5 GiB..

```

**Note**

Setting the root partition's type to "Linux root (x86-64)" is not required and the system will function normally if it is set to the "Linux filesystem" type. This filesystem type is only necessary for cases where a bootloader that supports it (i.e. systemd-boot) is used and a fstab file is not wanted.

After creating the root partition, press `t` to set the partition type, `3` to select the partition just created, and then type in _23_ to set the partition type to "Linux Root (x86-64)".

`Command(m for help):` `t`

```
Partition number (1-3, default 3): 3
Partition type or alias (type L to list all): 23

<!--T:239-->
Changed type of partition 'Linux filesystem' to 'Linux root (x86-64)'

```

After completing these steps, pressing `p` should display a partition table that looks similar to the following:

`Command (m for help):` `p`

```
Disk /dev/sda: 931.51 GiB, 1000204886016 bytes, 1953525168 sectors
Disk model: HGST HTS721010A9
Units: sectors of 1 * 512 = 512 bytes
Sector size (logical/physical): 512 bytes / 4096 bytes
I/O size (minimum/optimal): 4096 bytes / 4096 bytes
Disklabel type: gpt
Disk identifier: 3E56EE74-0571-462B-A992-9872E3855D75

<!--T:240-->
Device        Start        End    Sectors   Size Type
/dev/sda1      2048    2099199    2097152     1G EFI System
/dev/sda2   2099200   10487807    8388608     4G Linux swap
/dev/sda3  10487808 1953523711 1943035904 926.5G Linux root (x86-64)

<!--T:241-->
Filesystem/RAID signature on partition 1 will be wiped.

```

### Saving the partition layout

Press `w` to save the partition layout and exit the fdisk utility:

`Command (m for help):` `w`

```
The partition table has been altered.
Calling ioctl() to re-read partition table.
Syncing disks.

```

With partitions now available, the next installation step is to fill them with filesystems.

## Partitioning the disk with MBR for BIOS / legacy boot

The following table provides a recommended partition layout for a trivial MBR DOS / legacy BIOS boot installation. Additional partitions can be added according to personal preference or system design goals.

Device path (sysfs)
Mount point
File system
DPS UUID (PARTUUID)
Description
/dev/sda1/bootxfs
N/A
MBR DOS / legacy BIOS boot partition details.
/dev/sda2N/A. Swap is not mounted to the filesystem like a device file.swap
0657fd6d-a4ab-43c4-84e5-0933c84b4f4f
Swap partition details.
/dev/sda3/xfs
4f68bce3-e8cd-4db1-96e7-fbcaf984b709
Root partition details.

Change the partition layout according to personal preference.

### Viewing the current partition layout

Fire up fdisk against the disk (in our example, we use /dev/sda):

`root #` `fdisk /dev/sda`

Use the `p` key to display the disk's current partition configuration:

`Command (m for help):` `p`

```
Disk /dev/sda: 931.51 GiB, 1000204886016 bytes, 1953525168 sectors
Disk model: HGST HTS721010A9
Units: sectors of 1 * 512 = 512 bytes
Sector size (logical/physical): 512 bytes / 4096 bytes
I/O size (minimum/optimal): 4096 bytes / 4096 bytes
Disklabel type: dos
Disk identifier: 0xf163b576

<!--T:242-->
Device     Boot    Start        End    Sectors   Size Id Type
/dev/sda1  *        2048    2099199    2097152     1G 83 Linux
/dev/sda2        2099200   10487807    8388608     4G 82 Linux swap / Solaris
/dev/sda3       10487808 1953525167 1943037360 926.5G 83 Linux

```

This particular disk was until now configured to house two Linux filesystems (each with a corresponding partition listed as "Linux") as well as a swap partition (listed as "Linux swap"), using a GPT table.

### Creating a new disklabel / removing all partitions

Pressing `o` will instantly remove all existing disk partitions and create a new MBR disklabel (also named DOS disklabel):

`Command (m for help):` `o`

```
Created a new DOS disklabel with disk identifier 0xf163b576.
The device contains 'gpt' signature and it will be removed by a write command. See fdisk(8) man page and --wipe option for more details.

```

Alternatively, to keep an existing DOS disklabel (see the output of `p` above), consider removing the existing partitions one by one from the disk. Press `d` to delete a partition. For instance, to delete an existing /dev/sda1:

`Command (m for help):` `d`

```
Partition number (1-4): 1

```

The partition has now been scheduled for deletion. It will no longer show up when printing the list of partitions ( `p`, but it will not be erased until the changes have been saved. This allows users to abort the operation if a mistake was made - in that case, type `q` immediately and hit `Enter` and the partition will not be deleted.

Repeatedly press `p` to print out a partition listing and then press `d` and the number of the partition to delete it. Eventually, the partition table will be empty:

`Command (m for help):` `p`

```
Disk /dev/sda: 931.51 GiB, 1000204886016 bytes, 1953525168 sectors
Disk model: HGST HTS721010A9
Units: sectors of 1 * 512 = 512 bytes
Sector size (logical/physical): 512 bytes / 4096 bytes
I/O size (minimum/optimal): 4096 bytes / 4096 bytes
Disklabel type: dos
Disk identifier: 0xf163b576

```

The disk is now ready to create new partitions.

### Creating the boot partition

First, create a small partition which will be mounted as /boot. Press `n` to create a new partition, followed by `p` for a _primary_ partition and `1` to select the first primary partition. When prompted for the first sector, make sure it starts from 2048 (which may be needed for the boot loader) and press `Enter`. When prompted for the last sector, type +1G to create a partition 1 GB in size:

`Command (m for help):` `n`

```
Partition type
   p   primary (0 primary, 0 extended, 4 free)
   e   extended (container for logical partitions)
Select (default p): p
Partition number (1-4, default 1): 1
First sector (2048-1953525167, default 2048):
Last sector, +/-sectors or +/-size{K,M,G,T,P} (2048-1953525167, default 1953525167): +1G

<!--T:243-->
Created a new partition 1 of type 'Linux' and of size 1 GiB.

```

Mark the partition as bootable by pressing the `a` key and pressing `Enter`:

`Command (m for help):` `a`

```
Selected partition 1
The bootable flag on partition 1 is enabled now.

```

Note: if more than one partition is available on the disk, then the partition to be flagged as bootable will have to be selected.

### Creating the swap partition

Next, to create the swap partition, press `n` to create a new partition, then `p`, then type `2` to create the second primary partition, /dev/sda2. When prompted for the first sector, press `Enter`. When prompted for the last sector, type +4G (or any other size needed for the swap space) to create a partition 4GB in size.

`Command (m for help):` `n`

```
Partition type
   p   primary (1 primary, 0 extended, 3 free)
   e   extended (container for logical partitions)
Select (default p): p
Partition number (2-4, default 2): 2
First sector (2099200-1953525167, default 2099200):
Last sector, +/-sectors or +/-size{K,M,G,T,P} (2099200-1953525167, default 1953525167): +4G

Created a new partition 2 of type 'Linux' and of size 4 GiB.

```

After all this is done, press `t` to set the partition type, `2` to select the partition just created and then type in _82_ to set the partition type to "Linux Swap".

`Command (m for help):` `t`

```
Partition number (1,2, default 2): 2
Hex code (type L to list all codes): 82

<!--T:244-->
Changed type of partition 'Linux' to 'Linux swap / Solaris'.

```

### Creating the root partition

Finally, to create the root partition, press `n` to create a new partition. Then press `p` and `3` to create the third primary partition, /dev/sda3. When prompted for the first sector, hit `Enter`. When prompted for the last sector, hit `Enter` to create a partition that takes up the rest of the remaining space on the disk:

`Command (m for help):` `n`

```
Partition type
   p   primary (2 primary, 0 extended, 2 free)
   e   extended (container for logical partitions)
Select (default p): p
Partition number (3,4, default 3): 3
First sector (10487808-1953525167, default 10487808):
Last sector, +/-sectors or +/-size{K,M,G,T,P} (10487808-1953525167, default 1953525167):

<!--T:245-->
Created a new partition 3 of type 'Linux' and of size 926.5 GiB.

```

After completing these steps, pressing `p` should display a partition table that looks similar to this:

`Command (m for help):` `p`

```
Disk /dev/sda: 931.51 GiB, 1000204886016 bytes, 1953525168 sectors
Disk model: HGST HTS721010A9
Units: sectors of 1 * 512 = 512 bytes
Sector size (logical/physical): 512 bytes / 4096 bytes
I/O size (minimum/optimal): 4096 bytes / 4096 bytes
Disklabel type: dos
Disk identifier: 0xf163b576

<!--T:246-->
Device     Boot    Start        End    Sectors   Size Id Type
/dev/sda1  *        2048    2099199    2097152     1G 83 Linux
/dev/sda2        2099200   10487807    8388608     4G 82 Linux swap / Solaris
/dev/sda3       10487808 1953525167 1943037360 926.5G 83 Linux

```

### Saving the partition layout

Press `w` to write the partition layout and exit fdisk:

`Command (m for help):` `w`

```
The partition table has been altered.
Calling ioctl() to re-read partition table.
Syncing disks.

```

Now it is time to put filesystems on the partitions.

## إنشاء أنظمة الملفات

**Warning**

عند استخدام قرص SSD أو NVMe، من الحكمة التحقق من تحديثات البرمجيات الثابتة. تتطلب بعض أقراص Intel SSD على وجه الخصوص (طرازي 600p و6000p) تحديثًا للبرمجيات الثابتة لتجنب [احتمال فساد البيانات](https://bugzilla.redhat.com/show_bug.cgi?id=1402533) الناتج عن أنماط استخدام الإدخال والإخراج في XFS. المشكلة تكمن في مستوى البرمجيات الثابتة وليس خطأً في نظام ملفات XFS. يمكن لأداة smartctl المساعدة في التحقق من طراز الجهاز وإصدار البرمجيات الثابتة.

### مقدمة

بعد أن أنشئت الأقسام، حان الوقت لوضع نظام ملفات عليها. يصف القسم التالي مختلف أنظمة الملفات التي يدعمها لينكس. يمكن للقراء الذين يعرفون بالفعل نظام الملفات الذي سيستخدمونه الاستمرار في قسم [تطبيق نظام ملفات على قسم](/wiki/Handbook:AMD64/Installation/Disks/ar#Applying_a_filesystem_to_a_partition "Handbook:AMD64/Installation/Disks/ar"). أما الآخرون فينبغي لهم مواصلة القراءة للتعرف على نظم الملفات المتاحة...

### نظم الملفات

يدعم لينكس العشرات من نظم الملفات، على الرغم من أن العديد منها لا يُنصح باستخدامها إلا لأغراض محددة. قد تتوفر أنظمة ملفات معينة فقط بشكل مستقر على بنية amd64 - لذا يُنصح بالقراءة عن أنظمة الملفات وحالة دعمها قبل اختيار نظام تجريبي للأقسام المهمة. **يُعد XFS هو نظام الملفات الموصى به لجميع الأغراض والمنصات.** وفيما يلي قائمة غير شاملة:

[XFS](/wiki/XFS "XFS")نظام ملفات مع تدوين سجلات البيانات الوصفية، ويأتي مع مجموعة ميزات قوية ومحسن للقابلية للتوسع. لقد حُدِّث باستمرار ليشمل الميزات الحديثة. العيب الوحيد هو أن أقسام XFS لا يمكن تقليص حجمها بعد، رغم أن العمل جارٍ على ذلك. يدعم XFS بشكل ملحوظ الروابط المرجعية والنسخ عند الكتابة، وهو أمر مفيد بشكل خاص في أنظمة جنتو نظرًا لكثرة عمليات الترجمة التي يقوم بها المستخدمون. XFS هو نظام الملفات الحديث الموصى به لجميع الأغراض والمنصات. يتطلب أن يكون حجم القسم 300 ميجابايت على الأقل.[ext4](/wiki/Ext4 "Ext4")يعد Ext4 نظام ملفات موثوقًا لجميع الأغراض والمنصات، على الرغم من افتقاره للميزات الحديثة مثل الروابط المرجعية.[VFAT](/wiki/VFAT "VFAT")يُعرف أيضًا باسم FAT32، وهو مدعوم في لينكس ولكنه لا يدعم إعدادات أذونات يونكس القياسية. يستخدم في الغالب للتوافق والتبادل مع نظم التشغيل الأخرى (مايكروسوفت ويندوز أو أبل ماك أو إس) ولكنه ضروري أيضًا لبعض البرمجيات الثابتة لمحمل إقلاع النظام (مثل UEFI). سيحتاج مستخدمو أنظمة UEFI إلى [قسم نظام EFI](/wiki/EFI_System_Partition "EFI System Partition") مهيأ بصيغة VFAT لكي يتمكنوا من الإقلاع.[btrfs](/wiki/Btrfs "Btrfs")نظام ملفات من الجيل الجديد. يوفر ميزات متقدمة مثل اللقطات، والترميم الذاتي من خلال الاختبارات الجدولية (checksums)، والضغط الشفاف، والمجلدات الفرعية، ومصفوفة الأقراص (RAID) المدمجة. لا تضمن الأنوية الأقدم من 5.4.y السلامة عند استخدام btrfs في بيئات الإنتاج لأن إصلاحات المشاكل الخطيرة لا تتوفر إلا في الإصدارات الأحدث من فروع النواة طويلة الأمد (LTS). كما أن RAID 5/6 ومجموعات الحصص غير آمنة في جميع إصدارات btrfs.[F2FS](/wiki/F2FS "F2FS")«نظام الملفات الصديق للذاكرة الوميضية» أنشأته شركة سامسونج في الأصل للاستخدام مع ذاكرة NAND الوميضية. ويعد خيارًا جيدًا عند تثبيت جنتو على بطاقات microSD، أو محركات USB، أو غيرها من أجهزة التخزين القائمة على الذاكرة الوميضية.[NTFS](/wiki/NTFS "NTFS")نظام ملفات «التقنية الجديدة» هو نظام الملفات الرائد في مايكروسوفت ويندوز منذ الإصدار Windows NT 3.1. وبشكل مشابه ل‍ VFAT، فإنه لا يخزن إعدادات أذونات يونكس أو السمات الموسعة الضرورية لعمل أنظمة بي إس دي أو لينكس بشكل صحيح، لذا لا ينبغي استخدامه كنظام ملفات للجذر في معظم الحالات. يجب أن يستخدم _فقط_ للتوافق أو تبادل البيانات مع أنظمة مايكروسوفت ويندوز (لاحظ التشديد على كلمة _فقط_).[ZFS](/wiki/ZFS "ZFS")**Important:** يمكن إنشاء مجمعات ZFS على قرص التنصيب الأدنى، ولمزيد من المعلومات، راجع [ZFS/rootfs](/wiki/ZFS/rootfs "ZFS/rootfs")نظام ملفات من الجيل القادم صممه ماثيو أرينز وجيف بونويك. صُمم حول بضع أفكار رئيسية: يجب أن تكون إدارة التخزين بسيطة، ويجب التعامل مع التكرارية من خلال نظام الملفات، ويجب ألا توقف أنظمة الملفات عن العمل للإصلاح أبدًا، وأهمية المحاكاة التلقائية لأسوأ السيناريوهات قبل طرح المصدر البرمجي، وكون سلامة البيانات هي الأولوية القصوى.

يمكن العثور على معلومات أكثر استفاضة حول نظم الملفات في مقال [نظام الملفات](/wiki/Filesystem "Filesystem") الذي يديره المجتمع.

### تطبيق نظام ملفات على قسم

**Note**

تأكد من تنصيب (emerge) حزمة أدوات مساحة المستخدم ذات الصلة لنظام الملفات المختار قبل إعادة التشغيل. ستكون هناك ملاحظة تذكيرية للقيام بذلك في نهاية عملية التنصيب.

لإنشاء نظام ملفات على قسم أو مجلد، تتوفر أدوات في مساحة المستخدم لكل نظام ملفات محتمل. انقر على اسم نظام الملفات في الجدول أدناه للحصول على معلومات إضافية حول كل منها:

نظام الملفات
أمر الإنشاء
متوفر في بيئة العمل المباشرة؟
الحزمة
[XFS](/wiki/XFS "XFS")mkfs.xfs نعم
[sys-fs/xfsprogs](https://packages.gentoo.org/packages/sys-fs/xfsprogs)[ext4](/wiki/Ext4 "Ext4")mkfs.ext4 نعم
[sys-fs/e2fsprogs](https://packages.gentoo.org/packages/sys-fs/e2fsprogs)[VFAT](/wiki/VFAT "VFAT") (FAT32, ...)
mkfs.vfat نعم
[sys-fs/dosfstools](https://packages.gentoo.org/packages/sys-fs/dosfstools)[btrfs](/wiki/Btrfs "Btrfs")mkfs.btrfs نعم
[sys-fs/btrfs-progs](https://packages.gentoo.org/packages/sys-fs/btrfs-progs)[F2FS](/wiki/F2FS "F2FS")mkfs.f2fs نعم
[sys-fs/f2fs-tools](https://packages.gentoo.org/packages/sys-fs/f2fs-tools)[NTFS](/wiki/NTFS "NTFS")mkfs.ntfs نعم
[sys-fs/ntfs3g](https://packages.gentoo.org/packages/sys-fs/ntfs3g)[ZFS](/wiki/ZFS "ZFS")zpool create ... نعم
[sys-fs/zfs](https://packages.gentoo.org/packages/sys-fs/zfs)

**Important**

يوصي الدليل بإنشاء أقسام جديدة كجزء من عملية التنصيب، ولكن من المهم ملاحظة أن تشغيل أي أمر mkfs سيؤدي إلى محو أي بيانات موجودة داخل القسم. عند الضرورة، تأكد من نسخ أي بيانات موجودة احتياطيًا بشكل مناسب قبل إنشاء نظام ملفات جديد.

على سبيل المثال، لتهيئة قسم الجذر (/dev/sda3) بصيغة xfs كما هو مستخدم في هيكل الأقسام التجريبي، ستسخدم الأوامر التالية:

`root #` `mkfs.xfs -c options=/usr/share/xfsprogs/mkfs/lts_6.6.conf /dev/sda3`

#### نظام ملفات قسم نظام EFI

يجب تهيئة قسم نظام EFI (/dev/sda1) بصيغة FAT32:

`root #` `mkfs.vfat -F 32 /dev/sda1`

#### نظام ملفات قسم إقلاع BIOS العتيق

يمكن للأنظمة التي تقلع عبر BIOS العتيق (Legacy BIOS) مع جدول أقسام MBR/DOS استخدام أي صيغة نظام ملفات يدعمها محمل الإقلاع.

على سبيل المثال، للتهيئة بصيغة XFS:

`root #` `mkfs.xfs -c options=/usr/share/xfsprogs/mkfs/lts_6.6.conf /dev/sda1`

#### أقسام ext4 الصغيرة

عند استخدام نظام ملفات ext4 على قسم صغير (أقل من 8 جيبي بايت)، يجب إنشاء نظام الملفات بالخيارات المناسبة لحجز فهارس (inodes) كافية. يمكن تحديد ذلك باستخدام الخيار `-T small`:

`root #` `mkfs.ext4 -T small /dev/<device>`

سيؤدي القيام بذلك إلى مضاعفة عدد الفهارس بمقدار أربع مرات لنظام ملفات معين، حيث تنخفض قيمة «البايت لكل فهرس» من فهرس واحد لكل 16 كيلوبايت إلى فهرس واحد لكل 4 كيلوبايت.

### تفعيل قسم الإبدال

أمر mkswap هو الذي يُستخدم لتهيئة أقسام الإبدال:

`root #` `mkswap /dev/sda2`

**Note**

يمكن للتنصيبات التي بدأت سابقًا ولكن لم تكتمل، استئناف عملية التنصيب من هذه النقطة في الدليل. استخدم هذا المسار كرابط دائم: [استئناف التنصيب من هنا](/wiki/Handbook:AMD64/Installation/Disks/ar#Resumed_installations_start_here "Handbook:AMD64/Installation/Disks/ar").

لتفعيل قسم الإبدال، استخدم أمر swapon:

`root #` `swapon /dev/sda2`

خطوة 'التفعيل' هذه ضرورية فقط لأن قسم الإبدال قد أنشئ حديثًا داخل بيئة العمل المباشرة. وبمجرد إعادة تشغيل النظام، وطالما عُرِّف قسم الإبدال بشكل صحيح داخل ملف fstab أو أي آلية وصل أخرى، فستفعل مساحة التبديل تلقائيًا.

## وصل قسم الجذر

قد تفتقر بعض بيئات العمل المباشرة إلى نقطة الوصل المقترحة لقسم جذر جنتو (/mnt/gentoo)، أو نقاط الوصل للأقسام الإضافية التي أنشئت في قسم تقسيم الأقراص:

`root #` `mkdir --parents /mnt/gentoo`

استمر في إنشاء نقاط الوصل الإضافية اللازمة لأي أقسام (مخصصة) إضافية أنشئت خلال الخطوات السابقة باستخدام أمر mkdir.

بعد إنشاء نقاط الوصل، حان الوقت لجعل الأقسام متاحة عبر أمر mount.

وصل قسم الجذر:

`root #` `mount /dev/sda3 /mnt/gentoo`

لتنصيبات EFI فقط، يجب وصل قسم ESP تحت موقع قسم الجذر:

`root #` `mkdir --parents /mnt/gentoo/efi`

استمر في وصل الأقسام الإضافية (المخصصة) حسب الضرورة باستخدام أمر mount.

**Note**

إذا كان المجلد /tmp/ يحتاج للوجود في قسم منفصل، فتأكد من تغيير أذوناته بعد الوصل:

`root #` `chmod 1777 /mnt/gentoo/tmp`

ينطبق هذا أيضًا على المجلد /var/tmp.

لاحقًا في الإرشادات، سيوصل نظام ملفات proc (واجهة افتراضية مع النواة) بالإضافة إلى أنظمة الملفات الوهمية الأخرى للنواة. ولكن أولاً يجب استخراج [ملف مرحلة جنتو](/wiki/Handbook:AMD64/Installation/Stage "Handbook:AMD64/Installation/Stage").

[← Configuring the network](/wiki/Handbook:AMD64/Installation/Networking "Previous part") [Home](/wiki/Handbook:AMD64 "Handbook:AMD64") [Installing the stage file →](/wiki/Handbook:AMD64/Installation/Stage "Next part")