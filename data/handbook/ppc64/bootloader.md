# Bootloader

Other languages:

- [Deutsch](/wiki/Handbook:PPC64/Installation/Bootloader/de "Handbuch:PPC64/Installation/Bootloader (100% translated)")
- English
- [español](/wiki/Handbook:PPC64/Installation/Bootloader/es "Manual de Gentoo: PPC64/Instalación/Arranque (100% translated)")
- [français](/wiki/Handbook:PPC64/Installation/Bootloader/fr "Handbook:PPC64/Installation/Bootloader/fr (100% translated)")
- [italiano](/wiki/Handbook:PPC64/Installation/Bootloader/it "Handbook:PPC64/Installation/Bootloader/it (100% translated)")
- [magyar](/wiki/Handbook:PPC64/Installation/Bootloader/hu "Handbook:PPC64/Installation/Bootloader/hu (100% translated)")
- [polski](/wiki/Handbook:PPC64/Installation/Bootloader/pl "Handbook:PPC64/Installation/Bootloader (100% translated)")
- [português do Brasil](/wiki/Handbook:PPC64/Installation/Bootloader/pt-br "Handbook:PPC64/Installation/Bootloader/pt-br (50% translated)")
- [русский](/wiki/Handbook:PPC64/Installation/Bootloader/ru "Handbook:PPC64/Installation/Bootloader (100% translated)")
- [தமிழ்](/wiki/Handbook:PPC64/Installation/Bootloader/ta "கையேடு:PPC64/நிறுவல்/துவக்கஏற்றி (100% translated)")
- [中文（中国大陆）‎](/wiki/Handbook:PPC64/Installation/Bootloader/zh-cn "手册：PPC64/安装/配置系统引导程序Bootloader (100% translated)")
- [日本語](/wiki/Handbook:PPC64/Installation/Bootloader/ja "ハンドブック:PPC64/インストール/ブートローダー (100% translated)")
- [한국어](/wiki/Handbook:PPC64/Installation/Bootloader/ko "Handbook:PPC64/Installation/Bootloader/ko (100% translated)")

[PPC64 Handbook](/wiki/Handbook:PPC64 "Handbook:PPC64")[Installation](/wiki/Handbook:PPC64/Full/Installation "Handbook:PPC64/Full/Installation")[About the installation](/wiki/Handbook:PPC64/Installation/About "Handbook:PPC64/Installation/About")[Choosing the media](/wiki/Handbook:PPC64/Installation/Media "Handbook:PPC64/Installation/Media")[Configuring the network](/wiki/Handbook:PPC64/Installation/Networking "Handbook:PPC64/Installation/Networking")[Preparing the disks](/wiki/Handbook:PPC64/Installation/Disks "Handbook:PPC64/Installation/Disks")[The stage file](/wiki/Handbook:PPC64/Installation/Stage "Handbook:PPC64/Installation/Stage")[Installing base system](/wiki/Handbook:PPC64/Installation/Base "Handbook:PPC64/Installation/Base")[Configuring the kernel](/wiki/Handbook:PPC64/Installation/Kernel "Handbook:PPC64/Installation/Kernel")[Configuring the system](/wiki/Handbook:PPC64/Installation/System "Handbook:PPC64/Installation/System")[Installing tools](/wiki/Handbook:PPC64/Installation/Tools "Handbook:PPC64/Installation/Tools")Configuring the bootloader[Finalizing](/wiki/Handbook:PPC64/Installation/Finalizing "Handbook:PPC64/Installation/Finalizing")[Working with Gentoo](/wiki/Handbook:PPC64/Full/Working "Handbook:PPC64/Full/Working")[Portage introduction](/wiki/Handbook:PPC64/Working/Portage "Handbook:PPC64/Working/Portage")[USE flags](/wiki/Handbook:PPC64/Working/USE "Handbook:PPC64/Working/USE")[Portage features](/wiki/Handbook:PPC64/Working/Features "Handbook:PPC64/Working/Features")[Initscript system](/wiki/Handbook:PPC64/Working/Initscripts "Handbook:PPC64/Working/Initscripts")[Environment variables](/wiki/Handbook:PPC64/Working/EnvVar "Handbook:PPC64/Working/EnvVar")[Working with Portage](/wiki/Handbook:PPC64/Full/Portage "Handbook:PPC64/Full/Portage")[Files and directories](/wiki/Handbook:PPC64/Portage/Files "Handbook:PPC64/Portage/Files")[Variables](/wiki/Handbook:PPC64/Portage/Variables "Handbook:PPC64/Portage/Variables")[Mixing software branches](/wiki/Handbook:PPC64/Portage/Branches "Handbook:PPC64/Portage/Branches")[Additional tools](/wiki/Handbook:PPC64/Portage/Tools "Handbook:PPC64/Portage/Tools")[Custom package repository](/wiki/Handbook:PPC64/Portage/CustomTree "Handbook:PPC64/Portage/CustomTree")[Advanced features](/wiki/Handbook:PPC64/Portage/Advanced "Handbook:PPC64/Portage/Advanced")[OpenRC network configuration](/wiki/Handbook:PPC64/Full/Networking "Handbook:PPC64/Full/Networking")[Getting started](/wiki/Handbook:PPC64/Networking/Introduction "Handbook:PPC64/Networking/Introduction")[Advanced configuration](/wiki/Handbook:PPC64/Networking/Advanced "Handbook:PPC64/Networking/Advanced")[Modular networking](/wiki/Handbook:PPC64/Networking/Modular "Handbook:PPC64/Networking/Modular")[Wireless](/wiki/Handbook:PPC64/Networking/Wireless "Handbook:PPC64/Networking/Wireless")[Adding functionality](/wiki/Handbook:PPC64/Networking/Extending "Handbook:PPC64/Networking/Extending")[Dynamic management](/wiki/Handbook:PPC64/Networking/Dynamic "Handbook:PPC64/Networking/Dynamic")

With the kernel configured and compiled and the necessary system configuration files filled in correctly, it is time to install a program that will fire up the kernel when the system boots. Such a program is called a boot loader.

**Note**

Currently using Petitboot on Talos systems is undocumented in Gentoo. Please add the steps to [TalosII#Bootloader](/wiki/TalosII#Bootloader "TalosII") and notify on this Discussion page when ready to merge into the Handbook.

## Contents

- [1Using GRUB](#Using_GRUB)
  - [1.1Installation](#Installation)
  - [1.2Mac hardware (G5)](#Mac_hardware_.28G5.29)
    - [1.2.1Setup bootstrap partition](#Setup_bootstrap_partition)
    - [1.2.2Setup GRUB](#Setup_GRUB)
  - [1.3IBM hardware](#IBM_hardware)
    - [1.3.1Setup GRUB](#Setup_GRUB_2)
    - [1.3.2Grub config](#Grub_config)
- [2Rebooting the system](#Rebooting_the_system)

## Using GRUB\[ [edit](/index.php?title=Handbook:PPC64/Blocks/Bootloader&action=edit&section=T-1 "Edit section: ")\]

GRUB is a bootloader for PPC64 powered Linux machines.

### Installation\[ [edit](/index.php?title=Handbook:PPC64/Blocks/Bootloader&action=edit&section=T-2 "Edit section: ")\]

`root #` `emerge --ask sys-boot/grub`

### Mac hardware (G5)\[ [edit](/index.php?title=Handbook:PPC64/Blocks/Bootloader&action=edit&section=T-3 "Edit section: ")\]

#### Setup bootstrap partition\[ [edit](/index.php?title=Handbook:PPC64/Blocks/Bootloader&action=edit&section=T-4 "Edit section: ")\]

First, prepare the bootstrap partition that was created created during the [preparing the disk](/wiki/Handbook:PPC64/Installation/Disks "Handbook:PPC64/Installation/Disks") step. Following the example, this partition should be /dev/sda2. Optionally, confirm this by using parted:

Replace /dev/sda with the correct device if required.

`root #` `parted /dev/sda print`

```
Model: ATA Patriot Burst El (scsi)

<!--T:56-->
Disk /dev/sda: 120GB

<!--T:57-->
Sector size (logical/physical): 512B/512B

<!--T:58-->
Partition Table: mac

<!--T:59-->
Disk Flags:

<!--T:60-->
Number  Start   End     Size    File system  Name       Flags
 1      512B    32.8kB  32.3kB               Apple
 2      32.8kB  852kB   819kB   hfs          bootstrap  boot
 3      852kB   538MB   537MB   ext4         Boot
 4      538MB   54.2GB  53.7GB  ext4         Gentoo

```

In this output, partition 2 has the bootstrap information so /dev/sda2 is the correct path to use.

Format this partition as [HFS](/wiki/HFS "HFS") using the hformat command which is part of the [sys-fs/hfsutils](https://packages.gentoo.org/packages/sys-fs/hfsutils) package:

`root #` `dd if=/dev/zero of=/dev/sda2 bs=512`

`root #` `hformat -l bootstrap /dev/sda2`

Create a directory to mount the bootstrap partition and then mount it:

`root #` `mkdir /boot/NWBB
`

`root #` `mount --types hfs /dev/sda2 /boot/NWBB`

#### Setup GRUB\[ [edit](/index.php?title=Handbook:PPC64/Blocks/Bootloader&action=edit&section=T-5 "Edit section: ")\]

`root #` `grub-install --macppc-directory=/boot/NWBB /dev/sda2`

If it installs without errors, unmount the bootstrap:

`root #` `umount /boot/NWBB`

Next, bless the partition so it will boot:

`root #` `hmount /dev/sda2
`

`root #` `hattrib -t tbxi -c UNIX :System:Library:CoreServices:BootX
`

`root #` `hattrib -b :System:Library:CoreServices
`

`root #` `humount`

Finally, build the grub.cfg file:

`root #` `grub-mkconfig -o /boot/grub/grub.cfg`

### IBM hardware\[ [edit](/index.php?title=Handbook:PPC64/Blocks/Bootloader&action=edit&section=T-6 "Edit section: ")\]

Setting up Grub on IBM hardware is as simple as:

#### Setup GRUB\[ [edit](/index.php?title=Handbook:PPC64/Blocks/Bootloader&action=edit&section=T-7 "Edit section: ")\]

`root #` `grub-install /dev/sda1`

**Note**

/dev/sda1 is the PReP boot partition made in the partitioning stage

#### Grub config\[ [edit](/index.php?title=Handbook:PPC64/Blocks/Bootloader&action=edit&section=T-8 "Edit section: ")\]

Finally. build the grub.cfg file:

`root #` `grub-mkconfig -o /boot/grub/grub.cfg`

## Rebooting the system\[ [edit](/index.php?title=Handbook:Parts/Installation/Bootloader&action=edit&section=T-1 "Edit section: ")\]

Exit the chrooted environment and unmount all mounted partitions. Then type in that one magical command that initiates the final, true test: reboot.

`(chroot) livecd #` `exit`

`livecd~#` `cd
`

`livecd~#` `umount -l /mnt/gentoo/dev{/shm,/pts,}
`

`livecd~#` `umount -R /mnt/gentoo
`

`livecd~#` `reboot`

Do not forget to remove the live image, otherwise it may be targeted again instead of the newly installed Gentoo system!

Once rebooted in the fresh Gentoo environment, continue to the [Finalizing the Gentoo installation](/wiki/Handbook:PPC64/Installation/Finalizing "Handbook:PPC64/Installation/Finalizing") for important information on setting up the new installation, and finally some tips on how to start a productive Gentoo Linux journey.

[← Installing tools](/wiki/Handbook:PPC64/Installation/Tools "Previous part") [Home](/wiki/Handbook:PPC64 "Handbook:PPC64") [Finalizing →](/wiki/Handbook:PPC64/Installation/Finalizing "Next part")