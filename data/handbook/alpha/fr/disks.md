# Disks

Other languages:

- [Deutsch](/wiki/Handbook:Alpha/Installation/Disks/de "Handbuch:Alpha/Installation/Festplatten (100% translated)")
- [English](/wiki/Handbook:Alpha/Installation/Disks "Handbook:Alpha/Installation/Disks (100% translated)")
- [Türkçe](/wiki/Handbook:Alpha/Installation/Disks/tr "Handbook:Alpha/Installation/Disks/tr (0% translated)")
- [español](/wiki/Handbook:Alpha/Installation/Disks/es "Manual de Gentoo: Alpha/Instalación/Discos (100% translated)")
- français
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

[Sommaire du manuel](/wiki/Handbook:Alpha "Handbook:Alpha")[Installation](/wiki/Handbook:Alpha/Full/Installation/fr "Handbook:Alpha/Full/Installation/fr")[À propos de l’installation](/wiki/Handbook:Alpha/Installation/About/fr "Handbook:Alpha/Installation/About/fr")[Choix du support](/wiki/Handbook:Alpha/Installation/Media/fr "Handbook:Alpha/Installation/Media/fr")[Configurer le réseau](/wiki/Handbook:Alpha/Installation/Networking/fr "Handbook:Alpha/Installation/Networking/fr")Préparer les disques[Installer l’archive _stage3_](/wiki/Handbook:Alpha/Installation/Stage/fr "Handbook:Alpha/Installation/Stage/fr")[Installer le système de base](/wiki/Handbook:Alpha/Installation/Base/fr "Handbook:Alpha/Installation/Base/fr")[Configurer le noyau](/wiki/Handbook:Alpha/Installation/Kernel/fr "Handbook:Alpha/Installation/Kernel/fr")[Configurer le système](/wiki/Handbook:Alpha/Installation/System/fr "Handbook:Alpha/Installation/System/fr")[Installer les outils](/wiki/Handbook:Alpha/Installation/Tools/fr "Handbook:Alpha/Installation/Tools/fr")[Configurer le système d’amorçage](/wiki/Handbook:Alpha/Installation/Bootloader/fr "Handbook:Alpha/Installation/Bootloader/fr")[Finaliser](/wiki/Handbook:Alpha/Installation/Finalizing/fr "Handbook:Alpha/Installation/Finalizing/fr")[Utiliser Gentoo](/wiki/Handbook:Alpha/Full/Working/fr "Handbook:Alpha/Full/Working/fr")[Introduction à Portage](/wiki/Handbook:Alpha/Working/Portage/fr "Handbook:Alpha/Working/Portage/fr")[Les options de la variable USE](/wiki/Handbook:Alpha/Working/USE/fr "Handbook:Alpha/Working/USE/fr")[Les fonctionnalités de Portage](/wiki/Handbook:Alpha/Working/Features/fr "Handbook:Alpha/Working/Features/fr")[Scripts d’initialisation systèmes](/wiki/Handbook:Alpha/Working/Initscripts/fr "Handbook:Alpha/Working/Initscripts/fr")[Variables d’environnement](/wiki/Handbook:Alpha/Working/EnvVar/fr "Handbook:Alpha/Working/EnvVar/fr")[Utiliser Portage](/wiki/Handbook:Alpha/Full/Portage/fr "Handbook:Alpha/Full/Portage/fr")[Fichiers et répertoires](/wiki/Handbook:Alpha/Portage/Files/fr "Handbook:Alpha/Portage/Files/fr")[Les variables](/wiki/Handbook:Alpha/Portage/Variables/fr "Handbook:Alpha/Portage/Variables/fr")[Mélanger plusieurs branches logicielles](/wiki/Handbook:Alpha/Portage/Branches/fr "Handbook:Alpha/Portage/Branches/fr")[Outils supplémentaires](/wiki/Handbook:Alpha/Portage/Tools/fr "Handbook:Alpha/Portage/Tools/fr")[Dépôt personnalisé](/wiki/Handbook:Alpha/Portage/CustomTree/fr "Handbook:Alpha/Portage/CustomTree/fr")[Fonctionnalités avancées](/wiki/Handbook:Alpha/Portage/Advanced/fr "Handbook:Alpha/Portage/Advanced/fr")[Configuration du réseau avec OpenRC](/wiki/Handbook:Alpha/Full/Networking/fr "Handbook:Alpha/Full/Networking/fr")[Bien démarrer](/wiki/Handbook:Alpha/Networking/Introduction/fr "Handbook:Alpha/Networking/Introduction/fr")[Configuration avancée](/wiki/Handbook:Alpha/Networking/Advanced/fr "Handbook:Alpha/Networking/Advanced/fr")[Les modules réseau](/wiki/Handbook:Alpha/Networking/Modular/fr "Handbook:Alpha/Networking/Modular/fr")[Sans fil](/wiki/Handbook:Alpha/Networking/Wireless/fr "Handbook:Alpha/Networking/Wireless/fr")[Ajouter des fonctionnalités](/wiki/Handbook:Alpha/Networking/Extending/fr "Handbook:Alpha/Networking/Extending/fr")[Gestion dynamique](/wiki/Handbook:Alpha/Networking/Dynamic/fr "Handbook:Alpha/Networking/Dynamic/fr")

## Contents

- [1Introduction aux périphériques de type bloc](#Introduction_aux_p.C3.A9riph.C3.A9riques_de_type_bloc)
  - [1.1Les périphériques de type bloc](#Les_p.C3.A9riph.C3.A9riques_de_type_bloc)
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
- [5Créer des systèmes de fichiers](#Cr.C3.A9er_des_syst.C3.A8mes_de_fichiers)
- [6Introduction](#Introduction)
  - [6.1Les systèmes de fichiers](#Les_syst.C3.A8mes_de_fichiers)
  - [6.2Mettre en œuvre un système de fichiers sur une partition](#Mettre_en_.C5.93uvre_un_syst.C3.A8me_de_fichiers_sur_une_partition)
    - [6.2.1Système de fichiers pour un BIOS déprécié](#Syst.C3.A8me_de_fichiers_pour_un_BIOS_d.C3.A9pr.C3.A9ci.C3.A9)
    - [6.2.2Petite partition ext4](#Petite_partition_ext4)
  - [6.3Activer la partition d’échange](#Activer_la_partition_d.E2.80.99.C3.A9change)
- [7Monter la partition racine](#Monter_la_partition_racine)

## Introduction aux périphériques de type bloc

### Les périphériques de type bloc

Étudions en détail les aspects de Gentoo et de Linux en général qui concernent les disques, incluant les périphériques de type bloc, les partitions, et les systèmes de fichiers de Linux. Une fois que les tenants et les aboutissants des disques seront compris, il sera possible d’établir les partitions et les systèmes de fichiers pour l’installation.

Pour commencer, intéressons-nous aux périphériques de type bloc. Les disques SCSI et Serial ATA sont tous les deux étiquetés avec des noms de périphérique tels que : /dev/sda, /dev/sdb, /dev/sdc, etc. Sur des machines plus modernes, les disques SSD NVMe basés sur PCI Express ont des noms de périphérique tels que : /dev/nvme0n1, /dev/nvme0n2, etc.

Le tableau suivant aidera les lecteurs à déterminer où trouver certaines catégories de périphériques de type bloc sur le système :

Type de périphériqueChemin par défautCommentaires
SATA, SAS, SCSI ou USB/dev/sdaSe trouve dans le matériel à partir de 2007 jusqu’à présent, ce sont les plus courants. Ils se connectent via un bus [_SATA_](https://en.wikipedia.org/wiki/Serial_ATA "wikipedia:Serial ATA"), [_SCSI_](https://en.wikipedia.org/wiki/SCSI "wikipedia:SCSI") ou [_USB_](https://en.wikipedia.org/wiki/USB "wikipedia:USB") comme un bloc de stockage. Par exemple, la 1re partition d’un périphérique SATA sera /dev/sda1.
NVM Express (NVMe)/dev/nvme0n1La dernière technologie pour les disques _SSD_. Les disques [NVMe](https://en.wikipedia.org/wiki/NVM_Express "wikipedia:NVM Express") sont connectés directement au bus PCI Express et ont la plus rapide vitesse de transfert sur le marché. Le matériel est sorti à partir de 2014, le support est maintenant fréquent. La 1re partition d’un périphérique NVMe sera /dev/nvme0n1p1.
MMC, eMMC et SD/dev/mmcblk0Les cartes [MMC](https://en.wikipedia.org/wiki/MultiMediaCard#eMMC "wikipedia:MultiMediaCard"), SD et les autres types de mémoire peuvent être utiles pour le stockage. Ceci étant dit, peu de systèmes permettent une amorce à partir de ce type de périphérique. Il est suggéré de ne **pas** utiliser cela pour une installation Linux ; mais plutôt de les cantonner au transfert de fichiers, leur but initial. Alternativement, ce type de stockage peut être pratique pour des sauvegardes à court terme.

Les périphériques de type bloc ci-dessus représentent une interface abstraite pour le disque. Les programmes utilisateurs peuvent utiliser ces périphériques de type bloc pour interagir avec le disque sans se soucier de savoir s’il est SATA, SCSI ou quelque chose d’autre. Le programme peut simplement adresser le stockage sur le disque comme un groupe de blocs contigus de 4096 octets (4 Kio), accessibles aléatoirement.

### Slices

Although it is theoretically possible to use a full disk to house a Linux system, this is almost never done in practice. Instead, full disk block devices are split up in smaller, more manageable block devices. On Alpha systems, these are called _slices_.

**Remarque**

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

**Remarque**

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

**Conseil**

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
First cylinder (1-5290, default 1): 1004
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
  b:     1004      5289      4286       ext2
  c:        1      5290*     5289*    unused        0     0

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

## Créer des systèmes de fichiers

**Attention !**

Il est sage de vérifier les mises à jour du micrologiciel des disques SSD ou MVNe. Certains SSD Intel (en particulier 600p et 6000p) nécessaire une mise à jour pour [éviter une corruption de données](https://bugzilla.redhat.com/show_bug.cgi?id=1402533) sur des types d’usage avec XFS. Le problème est au niveau du micrologiciel et pas XFS. L’outil smartctl peut aider à identifier le modèle et la version du micrologiciel.

## Introduction

Maintenant que les partitions ont été créées, il est temps d’y placer un système de fichiers. Dans la section suivante les différents systèmes de fichiers que Linux prend en charge seront décris. Les lecteurs qui connaissent déjà quel système de fichiers utiliser peuvent continuer avec [Appliquer un système de fichiers à une partition](/wiki/Handbook:Alpha/Installation/Disks/fr#Appliquer_un_syst.C3.A8me_de_fichiers_.C3.A0_une_partition "Handbook:Alpha/Installation/Disks/fr"). Les autres devraient continuer à lire pour en apprendre plus sur les systèmes de fichiers disponibles.

### Les systèmes de fichiers

Linux supporte des douzaines de systèmes de fichiers, bien que la plupart ne devrait être utilisés que dans des cas spécifiques. Seuls certains seront dans l’architecture alpha – il est conseillé de se renseigner sur les systèmes de fichiers et leur prise en charge avant d’en choisir un plus expérimental pour les partitions importantes. **XFS est celui recommandé pour tous les usages et toutes les plates-formes.** Ci-dessous une liste non exhaustive :

[XFS](/wiki/XFS/fr "XFS/fr")Un système de fichiers avec journalisation des métadonnées, doté d’un ensemble de fonctionnalités robuste et optimisé pour la mise à l’échelle. Il a été continuellement mis à jour avec de nouvelles fonctionnalités. Le seul inconvénient est que les partitions XFS ne peuvent pas encore être réduites, bien que ce soit en cours de développement. XFS supporte _reflink_ (un lien vers un bloc de données ; ie. : si le fichier original est modifié, les blocs sont copiés pour que le _reflink_ continue de comporter les mêmes valeurs) et CoW ( _Copy on Write_ ; ie. : si plusieurs processus ouvrent le même fichier et qu’un le modifie, cela crée une copie transparent pour que le changement ne soit pas visible par les autres processus) qui sont particulièrement utilie dans un système Gentoo par rapport au nombre de compilation qu’un utilisateur fait. Nécessaire d’avoir une partition d’au moins 300 Mio.[ext4](/wiki/Ext4/fr "Ext4/fr")Ext4 est un système de fichiers fiable, tous usages, toutes plates-forme tout, bien qu’il manque des fonctionnalités modernes comme les _reflinks_.[VFAT](/wiki/VFAT "VFAT")Également connu sous le nom FAT32, ce format est supporté par Linux mais ne prend pas en charge les paramètres d’autorisation UNIX standards. Il est principalement utilisé pour l’interopérabilité avec d’autres systèmes d’exploitation (Microsoft Windows ou Apple macOS) mais est également une nécessité pour certains micrologiciels systèmes (comme UEFI). Les utilisateurs d’un système UEFI devront avoir une [Partition système EFI](/index.php?title=Partition_syst%C3%A8me_EFI&action=edit&redlink=1 "Partition système EFI (page does not exist)") formattée avec VFAT pour démarrer.[btrfs](/wiki/Btrfs/fr "Btrfs/fr")Un système de fichiers de nouvelle génération offrant de nombreuses fonctionnalités avancées telles que les sauvegardes instantanés, l’auto-guérison via des sommes de contrôle, la compression transparente, les sous-volumes et le RAID intégré. Les noyaux avant 5.4.x ne sont pas sûrs pour un usage de Btrfs en production car certains correctifs importants sont seulement présent dans les dernières versions. Les quotas par groupe et le RAID 5/6 ne sont pas sûrs d’usage quelque soit la version.[F2FS](/wiki/F2FS/fr "F2FS/fr")Le système de fichiers ami des SSD (disque flash) a été créé par Samsung pour l’utilisation avec la mémoire flash NAND. C’est un choix décent lors de l’installation de Gentoo sur des cartes microSD, des clés USB ou autres périphériques de stockage flash.[NTFS](/wiki/NTFS "NTFS")Ce système de fichiers « nouvelle technologie » est le système de fichiers phare de Microsoft Windows depuis NT 3.1. Comme VFAT ci-dessus, il ne stocke pas les paramètres d’autorisation UNIX ni les attributs étendus nécessaires au bon fonctionnement de BSD ou de Linux. Il ne peut donc pas être utilisé comme système de fichiers racine (aussi appelé _root_) dans la majorité des cas. Il devrait _seulement_ être utilisé pour l’interopérabilité avec les systèmes Microsoft Windows (noter l’emphase sur _seulement_).[ZFS](/wiki/ZFS "ZFS")**Important:** Les partitions ZFS peuvent seulement être crées depuis un LiveGUI. Pour plus d’informations, se référer à [ZFS/rootfs](/wiki/ZFS/rootfs "ZFS/rootfs")Système de fichiers nouvelle génération créé par Matthew Ahrens et Jeff Bonwick. Il a été conçu autour de quelques idées clés : l’administration du stockage doit être simple, la redondance être géré par le système de fichiers, le système ne soit pas être arrêté pour une réparation, des simulations des pires scénarios avant de livrer les versions et l’intégrité des données est primordial.[bcachefs](/wiki/Bcachefs "Bcachefs")**Important:** Bcachefs est encore noté expérimental dans le noyau ; aussi vérifiez que les données soient bien sauvegardées régulièrement dans un support sans Bcachefs.Bcachefs est un système de fichiers arbre B ( _B-tree_) basé sur Bcache. Il comporte des fonctionnalités comme CoW ( _Copy on Write_), la compression et le chiffrement. Bcachefs est comparable à Btrfs et ZFS. Une différence notable est le stockage multi-niveaux natif ; permettant d’utiliser un disque plus rapide (comme un disque SSD ou NVMe) pour faire office de cache pour un disque plus lent dans une grappe gérant de manière transparente l’activité sur les fichiers.

Des informations plus approfondies sur les systèmes de fichiers peuvent être trouvées dans les documentations maintenues par la [communauté](/wiki/Filesystem/fr "Filesystem/fr").

### Mettre en œuvre un système de fichiers sur une partition

**Remarque**

Assurez-vous d’installer les utilitaires correspondant au système de fichiers choisi avant de redémarrer. Il y aura un rappel à la fin de cette procédure d’installation.

Pour créer un système de fichiers sur une partition ou un volume, des outils sont disponibles pour chaque système de fichiers. Cliquer sur le nom du système de fichiers dans le tableau ci-dessous pour plus d’informations sur chaque système de fichiers :

Système de fichiers
Commande pour la création
Sur le CD minimal ?
Paquet
[XFS](/wiki/XFS/fr "XFS/fr")mkfs.xfs Oui
[sys-fs/xfsprogs](https://packages.gentoo.org/packages/sys-fs/xfsprogs)[ext4](/wiki/Ext4/fr "Ext4/fr")mkfs.ext4 Oui
[sys-fs/e2fsprogs](https://packages.gentoo.org/packages/sys-fs/e2fsprogs)[VFAT](/wiki/VFAT "VFAT") (FAT32, ...)
mkfs.vfat Oui
[sys-fs/dosfstools](https://packages.gentoo.org/packages/sys-fs/dosfstools)[btrfs](/wiki/Btrfs/fr "Btrfs/fr")mkfs.btrfs Oui
[sys-fs/btrfs-progs](https://packages.gentoo.org/packages/sys-fs/btrfs-progs)[F2FS](/wiki/F2FS/fr "F2FS/fr")mkfs.f2fs Oui
[sys-fs/f2fs-tools](https://packages.gentoo.org/packages/sys-fs/f2fs-tools)[NTFS](/wiki/NTFS "NTFS")mkfs.ntfs Oui
[sys-fs/ntfs3g](https://packages.gentoo.org/packages/sys-fs/ntfs3g)[ZFS](/wiki/ZFS "ZFS")zpool create ... Non
[sys-fs/zfs](https://packages.gentoo.org/packages/sys-fs/zfs)

\|}

**Important**

Le manuel recommande une nouvelle partition dans le cadre d’une installation, mais il est important de noter que n’importe quelle commande mkfs va effacer toute donnée contenue dans ladite partition. Si nécessaire, assurez-vous que chaque donnée soit sauvegardée avant cela.

Par exemple, pour avoir une partition racine (/dev/sda3) en xfs, les commandes suivantes doivent être utilisées :

`root #` `mkfs.xfs /dev/sda3`

#### Système de fichiers pour un BIOS déprécié

Les systèmes démarrant via un BIOS déprécié et avec un disque MBR/DOS peuvent utiliser n’importe quelle système de fichiers supportés par l’application d’amorçage ( _bootloader_, comme GRUB).

Par exemple, pour formater en XFS :

`root #` `mkfs.xfs /dev/sda1`

#### Petite partition ext4

Lors de l'utilisation d’u système de fichiers ext4 sur une petite partition (moins de 8 Gio), le système de fichiers devraient être créé avec les options appropriées pour réserver suffisamment de nœuds d’index ( _inodes_). Cela peut se faire avec l’option `-T small` :

`root #` `mkfs.ext4 -T small /dev/<device>`

Faire cela quadruple le nombre de nœuds d’index pour un système de fichiers étant donné que son paramètre _bytes-per-inode_ (octets par nœud d’index) passe de un tous les 16 Kio à un tous les 4 Kio.

### Activer la partition d’échange

mkswap est la commande à utiliser pour initialiser les partitions d’échange :

`root #` `mkswap /dev/sda2`

**Remarque**

Les installations débutées précédemment, mais non terminée peuvent reprendre ici dans le manuel. Utilisez ce lien comme lien permanent : [La reprise d’installation reprend ici](/wiki/Handbook:Alpha/Installation/Disks/fr#Resumed_installations_start_here "Handbook:Alpha/Installation/Disks/fr").

Pour activer la partition d’échange, utilisez la commande swapon :

`root #` `swapon /dev/sda2`

Cette étape d’« activation » est nécessaire car la partition d’échange est nouvellement créée dans un environnement démarré. Lors du redémarrage, si la partition est correctement déclarée dans fstab ou un autre mécanisme de montage, l’espace d’échange sera activé automatiquement.

## Monter la partition racine

Certains environnements peuvent ignorer le montage de la partition racine Gentoo (/mnt/gentoo) suggéré, ou monter des partitions additionnelles créées lors du partitionnement :

`root #` `mkdir --parents /mnt/gentoo`

Continuez à créer les points de montage additionnel pour chaque partition durant les étapes précédentes avec la commande mkdir.

Avec les points de montage créés, il est maintenant le moment de rendre les partitions accessibles via la commande mount.

Monter la partition racine :

`root #` `mount /dev/sda3 /mnt/gentoo`

Continuez à monter des partitions additionnelles avec la commande mount.

**Remarque**

Si /tmp/ doit se trouver sur une partition séparée, pensez à changer ses droits d'accès après le montage :

`root #` `chmod 1777 /mnt/gentoo/tmp`

Cela vaut également pour /var/tmp.

Plus loin dans les instructions, le système de fichiers _proc_ (une interface virtuelle avec le noyau) ainsi que d’autres pseudos systèmes de fichiers du noyau seront montés. Mais d’abord, nous devons [copier les fichiers d’installation de Gentoo](/wiki/Handbook:Alpha/Installation/Stage/fr "Handbook:Alpha/Installation/Stage/fr").

[← Configurer le réseau](/wiki/Handbook:Alpha/Installation/Networking/fr "Previous part") [Accueil](/wiki/Handbook:Alpha/fr "Handbook:Alpha/fr") [Installer l’archive _stage3_ →](/wiki/Handbook:Alpha/Installation/Stage/fr "Next part")