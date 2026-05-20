# Disks

Other languages:

- [Deutsch](/wiki/Handbook:SPARC/Installation/Disks/de "Handbuch:SPARC/Installation/Festplatten (100% translated)")
- [English](/wiki/Handbook:SPARC/Installation/Disks "Handbook:SPARC/Installation/Disks (100% translated)")
- [español](/wiki/Handbook:SPARC/Installation/Disks/es "Manual de Gentoo: SPARC/Instalación/Discos (100% translated)")
- français
- [italiano](/wiki/Handbook:SPARC/Installation/Disks/it "Handbook:SPARC/Installation/Disks/it (50% translated)")
- [magyar](/wiki/Handbook:SPARC/Installation/Disks/hu "Handbook:SPARC/Installation/Disks/hu (100% translated)")
- [polski](/wiki/Handbook:SPARC/Installation/Disks/pl "Handbook:SPARC/Installation/Disks (100% translated)")
- [português do Brasil](/wiki/Handbook:SPARC/Installation/Disks/pt-br "Handbook:SPARC/Installation/Disks/pt-br (50% translated)")
- [русский](/wiki/Handbook:SPARC/Installation/Disks/ru "Handbook:SPARC/Installation/Disks (100% translated)")
- [தமிழ்](/wiki/Handbook:SPARC/Installation/Disks/ta "கையேடு:SPARC/நிறுவல்/வட்டுக்கள் (100% translated)")
- [中文（中国大陆）‎](/wiki/Handbook:SPARC/Installation/Disks/zh-cn "手册：SPARC/安装/准备磁盘 (100% translated)")
- [日本語](/wiki/Handbook:SPARC/Installation/Disks/ja "ハンドブック:SPARC/インストール/ディスク (100% translated)")
- [한국어](/wiki/Handbook:SPARC/Installation/Disks/ko "Handbook:SPARC/Installation/Disks/ko (100% translated)")

[Sommaire du manuel](/wiki/Handbook:SPARC "Handbook:SPARC")[Installation](/wiki/Handbook:SPARC/Full/Installation/fr "Handbook:SPARC/Full/Installation/fr")[À propos de l’installation](/wiki/Handbook:SPARC/Installation/About/fr "Handbook:SPARC/Installation/About/fr")[Choix du support](/wiki/Handbook:SPARC/Installation/Media/fr "Handbook:SPARC/Installation/Media/fr")[Configurer le réseau](/wiki/Handbook:SPARC/Installation/Networking/fr "Handbook:SPARC/Installation/Networking/fr")Préparer les disques[Installer l’archive _stage3_](/wiki/Handbook:SPARC/Installation/Stage/fr "Handbook:SPARC/Installation/Stage/fr")[Installer le système de base](/wiki/Handbook:SPARC/Installation/Base/fr "Handbook:SPARC/Installation/Base/fr")[Configurer le noyau](/wiki/Handbook:SPARC/Installation/Kernel/fr "Handbook:SPARC/Installation/Kernel/fr")[Configurer le système](/wiki/Handbook:SPARC/Installation/System/fr "Handbook:SPARC/Installation/System/fr")[Installer les outils](/wiki/Handbook:SPARC/Installation/Tools/fr "Handbook:SPARC/Installation/Tools/fr")[Configurer le système d’amorçage](/wiki/Handbook:SPARC/Installation/Bootloader/fr "Handbook:SPARC/Installation/Bootloader/fr")[Finaliser](/wiki/Handbook:SPARC/Installation/Finalizing/fr "Handbook:SPARC/Installation/Finalizing/fr")[Utiliser Gentoo](/wiki/Handbook:SPARC/Full/Working/fr "Handbook:SPARC/Full/Working/fr")[Introduction à Portage](/wiki/Handbook:SPARC/Working/Portage/fr "Handbook:SPARC/Working/Portage/fr")[Les options de la variable USE](/wiki/Handbook:SPARC/Working/USE/fr "Handbook:SPARC/Working/USE/fr")[Les fonctionnalités de Portage](/wiki/Handbook:SPARC/Working/Features/fr "Handbook:SPARC/Working/Features/fr")[Scripts d’initialisation systèmes](/wiki/Handbook:SPARC/Working/Initscripts/fr "Handbook:SPARC/Working/Initscripts/fr")[Variables d’environnement](/wiki/Handbook:SPARC/Working/EnvVar/fr "Handbook:SPARC/Working/EnvVar/fr")[Utiliser Portage](/wiki/Handbook:SPARC/Full/Portage/fr "Handbook:SPARC/Full/Portage/fr")[Fichiers et répertoires](/wiki/Handbook:SPARC/Portage/Files/fr "Handbook:SPARC/Portage/Files/fr")[Les variables](/wiki/Handbook:SPARC/Portage/Variables/fr "Handbook:SPARC/Portage/Variables/fr")[Mélanger plusieurs branches logicielles](/wiki/Handbook:SPARC/Portage/Branches/fr "Handbook:SPARC/Portage/Branches/fr")[Outils supplémentaires](/wiki/Handbook:SPARC/Portage/Tools/fr "Handbook:SPARC/Portage/Tools/fr")[Dépôt personnalisé](/wiki/Handbook:SPARC/Portage/CustomTree/fr "Handbook:SPARC/Portage/CustomTree/fr")[Fonctionnalités avancées](/wiki/Handbook:SPARC/Portage/Advanced/fr "Handbook:SPARC/Portage/Advanced/fr")[Configuration du réseau avec OpenRC](/wiki/Handbook:SPARC/Full/Networking/fr "Handbook:SPARC/Full/Networking/fr")[Bien démarrer](/wiki/Handbook:SPARC/Networking/Introduction/fr "Handbook:SPARC/Networking/Introduction/fr")[Configuration avancée](/wiki/Handbook:SPARC/Networking/Advanced/fr "Handbook:SPARC/Networking/Advanced/fr")[Les modules réseau](/wiki/Handbook:SPARC/Networking/Modular/fr "Handbook:SPARC/Networking/Modular/fr")[Sans fil](/wiki/Handbook:SPARC/Networking/Wireless/fr "Handbook:SPARC/Networking/Wireless/fr")[Ajouter des fonctionnalités](/wiki/Handbook:SPARC/Networking/Extending/fr "Handbook:SPARC/Networking/Extending/fr")[Gestion dynamique](/wiki/Handbook:SPARC/Networking/Dynamic/fr "Handbook:SPARC/Networking/Dynamic/fr")

## Contents

- [1Introduction aux périphériques de type bloc](#Introduction_aux_p.C3.A9riph.C3.A9riques_de_type_bloc)
  - [1.1Les périphériques de type bloc](#Les_p.C3.A9riph.C3.A9riques_de_type_bloc)
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
- [4Créer des systèmes de fichiers](#Cr.C3.A9er_des_syst.C3.A8mes_de_fichiers)
- [5Introduction](#Introduction)
  - [5.1Les systèmes de fichiers](#Les_syst.C3.A8mes_de_fichiers)
  - [5.2Mettre en œuvre un système de fichiers sur une partition](#Mettre_en_.C5.93uvre_un_syst.C3.A8me_de_fichiers_sur_une_partition)
    - [5.2.1Système de fichiers pour un BIOS déprécié](#Syst.C3.A8me_de_fichiers_pour_un_BIOS_d.C3.A9pr.C3.A9ci.C3.A9)
    - [5.2.2Petite partition ext4](#Petite_partition_ext4)
  - [5.3Activer la partition d’échange](#Activer_la_partition_d.E2.80.99.C3.A9change)
- [6Monter la partition racine](#Monter_la_partition_racine)

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
[BIOS boot partition](/wiki/Handbook:X86/Blocks/Disks#What_is_the_BIOS_boot_partition.3F.2Ffr "Handbook:X86/Blocks/Disks")/dev/sda2(swap)
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

**Important**

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

## Créer des systèmes de fichiers

**Attention !**

Il est sage de vérifier les mises à jour du micrologiciel des disques SSD ou MVNe. Certains SSD Intel (en particulier 600p et 6000p) nécessaire une mise à jour pour [éviter une corruption de données](https://bugzilla.redhat.com/show_bug.cgi?id=1402533) sur des types d’usage avec XFS. Le problème est au niveau du micrologiciel et pas XFS. L’outil smartctl peut aider à identifier le modèle et la version du micrologiciel.

## Introduction

Maintenant que les partitions ont été créées, il est temps d’y placer un système de fichiers. Dans la section suivante les différents systèmes de fichiers que Linux prend en charge seront décris. Les lecteurs qui connaissent déjà quel système de fichiers utiliser peuvent continuer avec [Appliquer un système de fichiers à une partition](/wiki/Handbook:SPARC/Installation/Disks/fr#Appliquer_un_syst.C3.A8me_de_fichiers_.C3.A0_une_partition "Handbook:SPARC/Installation/Disks/fr"). Les autres devraient continuer à lire pour en apprendre plus sur les systèmes de fichiers disponibles.

### Les systèmes de fichiers

Linux supporte des douzaines de systèmes de fichiers, bien que la plupart ne devrait être utilisés que dans des cas spécifiques. Seuls certains seront dans l’architecture sparc – il est conseillé de se renseigner sur les systèmes de fichiers et leur prise en charge avant d’en choisir un plus expérimental pour les partitions importantes. **XFS est celui recommandé pour tous les usages et toutes les plates-formes.** Ci-dessous une liste non exhaustive :

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

Par exemple, pour avoir une partition racine (/dev/sda1) en xfs, les commandes suivantes doivent être utilisées :

`root #` `mkfs.xfs /dev/sda1`

#### Système de fichiers pour un BIOS déprécié

Les systèmes démarrant via un BIOS déprécié et avec un disque MBR/DOS peuvent utiliser n’importe quelle système de fichiers supportés par l’application d’amorçage ( _bootloader_, comme GRUB).

Par exemple, pour formater en XFS :

`root #` `mkfs.xfs `

#### Petite partition ext4

Lors de l'utilisation d’u système de fichiers ext4 sur une petite partition (moins de 8 Gio), le système de fichiers devraient être créé avec les options appropriées pour réserver suffisamment de nœuds d’index ( _inodes_). Cela peut se faire avec l’option `-T small` :

`root #` `mkfs.ext4 -T small /dev/<device>`

Faire cela quadruple le nombre de nœuds d’index pour un système de fichiers étant donné que son paramètre _bytes-per-inode_ (octets par nœud d’index) passe de un tous les 16 Kio à un tous les 4 Kio.

### Activer la partition d’échange

mkswap est la commande à utiliser pour initialiser les partitions d’échange :

`root #` `mkswap /dev/sda2`

**Remarque**

Les installations débutées précédemment, mais non terminée peuvent reprendre ici dans le manuel. Utilisez ce lien comme lien permanent : [La reprise d’installation reprend ici](/wiki/Handbook:SPARC/Installation/Disks/fr#Resumed_installations_start_here "Handbook:SPARC/Installation/Disks/fr").

Pour activer la partition d’échange, utilisez la commande swapon :

`root #` `swapon /dev/sda2`

Cette étape d’« activation » est nécessaire car la partition d’échange est nouvellement créée dans un environnement démarré. Lors du redémarrage, si la partition est correctement déclarée dans fstab ou un autre mécanisme de montage, l’espace d’échange sera activé automatiquement.

## Monter la partition racine

Certains environnements peuvent ignorer le montage de la partition racine Gentoo (/mnt/gentoo) suggéré, ou monter des partitions additionnelles créées lors du partitionnement :

`root #` `mkdir --parents /mnt/gentoo`

Continuez à créer les points de montage additionnel pour chaque partition durant les étapes précédentes avec la commande mkdir.

Avec les points de montage créés, il est maintenant le moment de rendre les partitions accessibles via la commande mount.

Monter la partition racine :

`root #` `mount /dev/sda1 /mnt/gentoo`

Continuez à monter des partitions additionnelles avec la commande mount.

**Remarque**

Si /tmp/ doit se trouver sur une partition séparée, pensez à changer ses droits d'accès après le montage :

`root #` `chmod 1777 /mnt/gentoo/tmp`

Cela vaut également pour /var/tmp.

Plus loin dans les instructions, le système de fichiers _proc_ (une interface virtuelle avec le noyau) ainsi que d’autres pseudos systèmes de fichiers du noyau seront montés. Mais d’abord, nous devons [copier les fichiers d’installation de Gentoo](/wiki/Handbook:SPARC/Installation/Stage/fr "Handbook:SPARC/Installation/Stage/fr").

[← Configurer le réseau](/wiki/Handbook:SPARC/Installation/Networking/fr "Previous part") [Accueil](/wiki/Handbook:SPARC/fr "Handbook:SPARC/fr") [Installer l’archive _stage3_ →](/wiki/Handbook:SPARC/Installation/Stage/fr "Next part")