# Bootloader

Other languages:

- [Deutsch](/wiki/Handbook:PPC/Installation/Bootloader/de "Handbuch:PPC/Installation/Bootloader (100% translated)")
- English
- [español](/wiki/Handbook:PPC/Installation/Bootloader/es "Manual de Gentoo: PPC/Instalación/Arranque (100% translated)")
- [français](/wiki/Handbook:PPC/Installation/Bootloader/fr "Handbook:PPC/Installation/Bootloader/fr (100% translated)")
- [italiano](/wiki/Handbook:PPC/Installation/Bootloader/it "Handbook:PPC/Installation/Bootloader/it (100% translated)")
- [magyar](/wiki/Handbook:PPC/Installation/Bootloader/hu "Handbook:PPC/Installation/Bootloader/hu (100% translated)")
- [polski](/wiki/Handbook:PPC/Installation/Bootloader/pl "Handbook:PPC/Installation/Bootloader (100% translated)")
- [português do Brasil](/wiki/Handbook:PPC/Installation/Bootloader/pt-br "Handbook:PPC/Installation/Bootloader/pt-br (50% translated)")
- [русский](/wiki/Handbook:PPC/Installation/Bootloader/ru "Handbook:PPC/Installation/Bootloader (100% translated)")
- [தமிழ்](/wiki/Handbook:PPC/Installation/Bootloader/ta "கையேடு:PPC/நிறுவல்/துவக்கஏற்றி (100% translated)")
- [中文（中国大陆）‎](/wiki/Handbook:PPC/Installation/Bootloader/zh-cn "手册：PPC/安装/配置系统引导程序Bootloader (100% translated)")
- [日本語](/wiki/Handbook:PPC/Installation/Bootloader/ja "ハンドブック:PPC/インストール/ブートローダー (100% translated)")
- [한국어](/wiki/Handbook:PPC/Installation/Bootloader/ko "Handbook:PPC/Installation/Bootloader/ko (100% translated)")

[PPC Handbook](/wiki/Handbook:PPC "Handbook:PPC")[Installation](/wiki/Handbook:PPC/Full/Installation "Handbook:PPC/Full/Installation")[About the installation](/wiki/Handbook:PPC/Installation/About "Handbook:PPC/Installation/About")[Choosing the media](/wiki/Handbook:PPC/Installation/Media "Handbook:PPC/Installation/Media")[Configuring the network](/wiki/Handbook:PPC/Installation/Networking "Handbook:PPC/Installation/Networking")[Preparing the disks](/wiki/Handbook:PPC/Installation/Disks "Handbook:PPC/Installation/Disks")[The stage file](/wiki/Handbook:PPC/Installation/Stage "Handbook:PPC/Installation/Stage")[Installing base system](/wiki/Handbook:PPC/Installation/Base "Handbook:PPC/Installation/Base")[Configuring the kernel](/wiki/Handbook:PPC/Installation/Kernel "Handbook:PPC/Installation/Kernel")[Configuring the system](/wiki/Handbook:PPC/Installation/System "Handbook:PPC/Installation/System")[Installing tools](/wiki/Handbook:PPC/Installation/Tools "Handbook:PPC/Installation/Tools")Configuring the bootloader[Finalizing](/wiki/Handbook:PPC/Installation/Finalizing "Handbook:PPC/Installation/Finalizing")[Working with Gentoo](/wiki/Handbook:PPC/Full/Working "Handbook:PPC/Full/Working")[Portage introduction](/wiki/Handbook:PPC/Working/Portage "Handbook:PPC/Working/Portage")[USE flags](/wiki/Handbook:PPC/Working/USE "Handbook:PPC/Working/USE")[Portage features](/wiki/Handbook:PPC/Working/Features "Handbook:PPC/Working/Features")[Initscript system](/wiki/Handbook:PPC/Working/Initscripts "Handbook:PPC/Working/Initscripts")[Environment variables](/wiki/Handbook:PPC/Working/EnvVar "Handbook:PPC/Working/EnvVar")[Working with Portage](/wiki/Handbook:PPC/Full/Portage "Handbook:PPC/Full/Portage")[Files and directories](/wiki/Handbook:PPC/Portage/Files "Handbook:PPC/Portage/Files")[Variables](/wiki/Handbook:PPC/Portage/Variables "Handbook:PPC/Portage/Variables")[Mixing software branches](/wiki/Handbook:PPC/Portage/Branches "Handbook:PPC/Portage/Branches")[Additional tools](/wiki/Handbook:PPC/Portage/Tools "Handbook:PPC/Portage/Tools")[Custom package repository](/wiki/Handbook:PPC/Portage/CustomTree "Handbook:PPC/Portage/CustomTree")[Advanced features](/wiki/Handbook:PPC/Portage/Advanced "Handbook:PPC/Portage/Advanced")[OpenRC network configuration](/wiki/Handbook:PPC/Full/Networking "Handbook:PPC/Full/Networking")[Getting started](/wiki/Handbook:PPC/Networking/Introduction "Handbook:PPC/Networking/Introduction")[Advanced configuration](/wiki/Handbook:PPC/Networking/Advanced "Handbook:PPC/Networking/Advanced")[Modular networking](/wiki/Handbook:PPC/Networking/Modular "Handbook:PPC/Networking/Modular")[Wireless](/wiki/Handbook:PPC/Networking/Wireless "Handbook:PPC/Networking/Wireless")[Adding functionality](/wiki/Handbook:PPC/Networking/Extending "Handbook:PPC/Networking/Extending")[Dynamic management](/wiki/Handbook:PPC/Networking/Dynamic "Handbook:PPC/Networking/Dynamic")

## Contents

- [1Bootloader options](#Bootloader_options)
- [2NewWorld Macs](#NewWorld_Macs)
  - [2.1GRUB](#GRUB)
  - [2.2Installation](#Installation)
  - [2.3Setup bootstrap partition](#Setup_bootstrap_partition)
  - [2.4Setup GRUB](#Setup_GRUB)
- [3OldWorld Macs](#OldWorld_Macs)
  - [3.1BootX](#BootX)
- [4Pegasos](#Pegasos)
  - [4.1BootCreator](#BootCreator)
- [5Rebooting the system](#Rebooting_the_system)

## Bootloader options\[ [edit](/index.php?title=Handbook:PPC/Blocks/Bootloader&action=edit&section=T-1 "Edit section: ")\]

Now that the kernel is configured and compiled and the necessary system configuration files are filled in correctly, it is time to install a program that will fire up the kernel when the system is started. Such a program is called a boot loader.

The boot loader to use depends upon the type of PPC machine.

For a NewWorld Apple or IBM machine, grub needs to be selected. OldWorld Apple machines havs one option: BootX. The Pegasos does not require a boot loader, but it is necessary to emerge bootcreator to create SmartFirmware boot menus.

## NewWorld Macs\[ [edit](/index.php?title=Handbook:PPC/Blocks/Bootloader&action=edit&section=T-2 "Edit section: ")\]

### GRUB\[ [edit](/index.php?title=Handbook:PPC/Blocks/Bootloader&action=edit&section=T-3 "Edit section: ")\]

### Installation\[ [edit](/index.php?title=Handbook:PPC/Blocks/Bootloader&action=edit&section=T-4 "Edit section: ")\]

`root #` `emerge --ask sys-boot/grub`

### Setup bootstrap partition\[ [edit](/index.php?title=Handbook:PPC/Blocks/Bootloader&action=edit&section=T-5 "Edit section: ")\]

First, prepare the bootstrap partition that was created created during the [preparing the disk](/wiki/Handbook:PPC64/Installation/Disks "Handbook:PPC64/Installation/Disks") step. Following the example, this partition should be /dev/sda2. Optionally, confirm this by using parted:

Replace /dev/sda with the correct device if required.

`root #` `parted /dev/sda print`

```
Model: ATA Patriot Burst El (scsi)

<!--T:85-->
Disk /dev/sda: 120GB

<!--T:86-->
Sector size (logical/physical): 512B/512B

<!--T:87-->
Partition Table: mac

<!--T:88-->
Disk Flags:

<!--T:89-->
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

### Setup GRUB\[ [edit](/index.php?title=Handbook:PPC/Blocks/Bootloader&action=edit&section=T-6 "Edit section: ")\]

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

## OldWorld Macs\[ [edit](/index.php?title=Handbook:PPC/Blocks/Bootloader&action=edit&section=T-7 "Edit section: ")\]

### BootX\[ [edit](/index.php?title=Handbook:PPC/Blocks/Bootloader&action=edit&section=T-8 "Edit section: ")\]

**Important**

BootX can only be used on OldWorld Apple systems with MacOS classic 7 to 9. For machines under 32MB of RAM use MacOS classic 7.

BootX can be downloaded at [https://github.com/immolo/BootX/archive/refs/tags/1.2.2.tar.gz](https://github.com/immolo/BootX/archive/refs/tags/1.2.2.tar.gz)

Since BootX boots Linux from within MacOS, the kernel will need to be copied from the Linux Partition to the MacOS partition. First, mount the MacOS partition from outside of the chroot. Use mac-fdisk -l to find the MacOS partition number, sda6 is used as an example here. Once the partition is mounted, we'll copy the kernel to the system folder so BootX can find it.

`root #` `exit`

`cdimage ~#` `mkdir /mnt/mac
`

`cdimage ~#` `mount /dev/sda6 /mnt/mac -t hfs
`

`cdimage ~#` `cp /mnt/gentoo/usr/src/linux/vmlinux "/mnt/mac/System Folder/Linux Kernels/kernel-6.18.8-gentoo"`

Now that the kernel is copied over, we'll need to reboot to set up BootX.

`cdimage ~#` `cd /
`

`cdimage ~#` `umount -l /mnt/gentoo/dev{/pts,/shm,}
`

`cdimage ~#` `umount -l /mnt/gentoo{/proc,/sys,}
`

`cdimage ~#` `umount -l /mnt/mac
`

`cdimage ~#` `reboot`

Of course, don't forget to remove the bootable CD, otherwise the CD will be booted again instead of MacOS.

Once the machine has booted into MacOS, open the BootX control panel. When not using genkernel, select Options and uncheck Use specified RAM disk. If genkernel is used, ensure that the genkernel initrd is selected instead of the Installation CD initrd. If not using genkernel, there is now an option to specify the machine's Linux root disk and partition. Fill these in with the appropriate values. Depending upon the kernel configuration, additional boot arguments may need to be applied.

BootX can be configured to start Linux upon boot. If this is done, then the machine will boot into MacOS first and, during startup, BootX will load and start Linux. See the BootX home page for more information.

**Important**

Make sure to include support for the HFS and HFS+ filesystems in the kernel, otherwise upgrades or changes to the kernel on the MacOS partition will not be possible.

## Pegasos\[ [edit](/index.php?title=Handbook:PPC/Blocks/Bootloader&action=edit&section=T-9 "Edit section: ")\]

**Note**

Pegasos also has Grub support but this is currently undocumented in Gentoo. Please add this to the main wiki and notify on this page's discussion once ready to migrated here.

### BootCreator\[ [edit](/index.php?title=Handbook:PPC/Blocks/Bootloader&action=edit&section=T-10 "Edit section: ")\]

**Important**

BootCreator will build a nice SmartFirmware bootmenu written in Forth for the Pegasos.

First make sure to have bootcreator installed on the system:

`root #` `emerge --ask sys-boot/bootcreator`

Now copy the file /etc/bootmenu.example into /etc/bootmenu/ and edit it to suit personal needs:

`root #` `cp /etc/bootmenu.example /etc/bootmenu
`

`root #` `nano -w /etc/bootmenu`

Below is a complete /etc/bootmenu config file. vmlinux and initrd should be replaced by the kernel and initrd image names.

FILE **`/etc/bootmenu`** **Example bootcreator configuration file**

```
#
# Example description file for bootcreator 1.1
#

[VERSION]
1

[TITLE]
Boot Menu

[SETTINGS]
AbortOnKey = false
Timeout    = 9
Default    = 1

[SECTION]
Local HD -> Morphos      (Normal)
ide:0 boot2.img ramdebug edebugflags="logkprintf"

[SECTION]
Local HD -> Linux (Normal)
ide:0 kernel-6.18.8-gentoo video=radeonfb:1024x768@70 root=/dev/sda3

[SECTION]
Local HD -> Genkernel (Normal)
ide:0 kernel-genkernel-ppc-6.18.8-gentoo root=/dev/ram0
root=/dev/sda3 initrd=initramfs-genkernel-ppc-6.18.8-gentoo

```

Finally the bootmenu must be transferred into Forth and copied to the boot partition, so that the SmartFirmware can read it. Therefore it is necessar to call bootcreator:

`root #` `bootcreator /etc/bootmenu /boot/menu`

**Note**

Be sure to have a look into the SmartFirmware's settings when rebooting, that menu is the file that will be loaded by default.

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

Once rebooted in the fresh Gentoo environment, continue to the [Finalizing the Gentoo installation](/wiki/Handbook:PPC/Installation/Finalizing "Handbook:PPC/Installation/Finalizing") for important information on setting up the new installation, and finally some tips on how to start a productive Gentoo Linux journey.

[← Installing tools](/wiki/Handbook:PPC/Installation/Tools "Previous part") [Home](/wiki/Handbook:PPC "Handbook:PPC") [Finalizing →](/wiki/Handbook:PPC/Installation/Finalizing "Next part")