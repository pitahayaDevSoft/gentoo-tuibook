# Bootloader

Other languages:

- Deutsch
- [English](/wiki/Handbook:PPC/Installation/Bootloader "Handbook:PPC/Installation/Bootloader (100% translated)")
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

[PPC Handbuch](/wiki/Handbook:PPC/de "Handbook:PPC/de")[Installation](/wiki/Handbook:PPC/Full/Installation/de "Handbook:PPC/Full/Installation/de")[Über die Installation](/wiki/Handbook:PPC/Installation/About/de "Handbook:PPC/Installation/About/de")[Auswahl des Mediums](/wiki/Handbook:PPC/Installation/Media/de "Handbook:PPC/Installation/Media/de")[Konfiguration des Netzwerks](/wiki/Handbook:PPC/Installation/Networking/de "Handbook:PPC/Installation/Networking/de")[Vorbereiten der Festplatte(n)](/wiki/Handbook:PPC/Installation/Disks/de "Handbook:PPC/Installation/Disks/de")[Installation des Stage Archivs](/wiki/Handbook:PPC/Installation/Stage/de "Handbook:PPC/Installation/Stage/de")[Installation des Basissystems](/wiki/Handbook:PPC/Installation/Base/de "Handbook:PPC/Installation/Base/de")[Konfiguration des Kernels](/wiki/Handbook:PPC/Installation/Kernel/de "Handbook:PPC/Installation/Kernel/de")[Konfiguration des Systems](/wiki/Handbook:PPC/Installation/System/de "Handbook:PPC/Installation/System/de")[Installation der Tools](/wiki/Handbook:PPC/Installation/Tools/de "Handbook:PPC/Installation/Tools/de")Konfiguration des Bootloaders[Abschluss](/wiki/Handbook:PPC/Installation/Finalizing/de "Handbook:PPC/Installation/Finalizing/de")[Arbeiten mit Gentoo](/wiki/Handbook:PPC/Full/Working/de "Handbook:PPC/Full/Working/de")[Portage-Einführung](/wiki/Handbook:PPC/Working/Portage/de "Handbook:PPC/Working/Portage/de")[USE-Flags](/wiki/Handbook:PPC/Working/USE/de "Handbook:PPC/Working/USE/de")[Portage-Features](/wiki/Handbook:PPC/Working/Features/de "Handbook:PPC/Working/Features/de")[Initskript-System](/wiki/Handbook:PPC/Working/Initscripts/de "Handbook:PPC/Working/Initscripts/de")[Umgebungsvariablen](/wiki/Handbook:PPC/Working/EnvVar/de "Handbook:PPC/Working/EnvVar/de")[Arbeiten mit Portage](/wiki/Handbook:PPC/Full/Portage/de "Handbook:PPC/Full/Portage/de")[Dateien und Verzeichnisse](/wiki/Handbook:PPC/Portage/Files/de "Handbook:PPC/Portage/Files/de")[Variablen](/wiki/Handbook:PPC/Portage/Variables/de "Handbook:PPC/Portage/Variables/de")[Mischen von Softwarezweigen](/wiki/Handbook:PPC/Portage/Branches/de "Handbook:PPC/Portage/Branches/de")[Zusätzliche Tools](/wiki/Handbook:PPC/Portage/Tools/de "Handbook:PPC/Portage/Tools/de")[Eigener Portage-Tree](/wiki/Handbook:PPC/Portage/CustomTree/de "Handbook:PPC/Portage/CustomTree/de")[Erweiterte Portage-Features](/wiki/Handbook:PPC/Portage/Advanced/de "Handbook:PPC/Portage/Advanced/de")[Netzwerk-Konfiguration](/wiki/Handbook:PPC/Full/Networking/de "Handbook:PPC/Full/Networking/de")[Zu Beginn](/wiki/Handbook:PPC/Networking/Introduction/de "Handbook:PPC/Networking/Introduction/de")[Fortgeschrittene Konfiguration](/wiki/Handbook:PPC/Networking/Advanced/de "Handbook:PPC/Networking/Advanced/de")[Modulare Vernetzung](/wiki/Handbook:PPC/Networking/Modular/de "Handbook:PPC/Networking/Modular/de")[Drahtlose Netzwerke](/wiki/Handbook:PPC/Networking/Wireless/de "Handbook:PPC/Networking/Wireless/de")[Funktionalität hinzufügen](/wiki/Handbook:PPC/Networking/Extending/de "Handbook:PPC/Networking/Extending/de")[Dynamisches Management](/wiki/Handbook:PPC/Networking/Dynamic/de "Handbook:PPC/Networking/Dynamic/de")

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
- [5Neustart des Systems](#Neustart_des_Systems)

## Bootloader options

Now that the kernel is configured and compiled and the necessary system configuration files are filled in correctly, it is time to install a program that will fire up the kernel when the system is started. Such a program is called a boot loader.

The boot loader to use depends upon the type of PPC machine.

For a NewWorld Apple or IBM machine, grub needs to be selected. OldWorld Apple machines havs one option: BootX. The Pegasos does not require a boot loader, but it is necessary to emerge bootcreator to create SmartFirmware boot menus.

## NewWorld Macs

### GRUB

### Installation

`root #` `emerge --ask sys-boot/grub`

### Setup bootstrap partition

First, prepare the bootstrap partition that was created created during the [preparing the disk](/wiki/Handbook:PPC64/Installation/Disks "Handbook:PPC64/Installation/Disks") step. Following the example, this partition should be /dev/sda2. Optionally, confirm this by using parted:

Replace /dev/sda with the correct device if required.

`root #` `parted /dev/sda print`

```
Model: ATA Patriot Burst El (scsi)
</div>

<div lang="en" dir="ltr" class="mw-content-ltr">
Disk /dev/sda: 120GB
</div>

<div lang="en" dir="ltr" class="mw-content-ltr">
Sector size (logical/physical): 512B/512B
</div>

<div lang="en" dir="ltr" class="mw-content-ltr">
Partition Table: mac
</div>

<div lang="en" dir="ltr" class="mw-content-ltr">
Disk Flags:
</div>

<div lang="en" dir="ltr" class="mw-content-ltr">
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

### Setup GRUB

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

## OldWorld Macs

### BootX

**Wichtig**

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

**Wichtig**

Make sure to include support for the HFS and HFS+ filesystems in the kernel, otherwise upgrades or changes to the kernel on the MacOS partition will not be possible.

## Pegasos

**Hinweis**

Pegasos also has Grub support but this is currently undocumented in Gentoo. Please add this to the main wiki and notify on this page's discussion once ready to migrated here.

### BootCreator

**Wichtig**

BootCreator will build a nice SmartFirmware bootmenu written in Forth for the Pegasos.

First make sure to have bootcreator installed on the system:

`root #` `emerge --ask sys-boot/bootcreator`

Now copy the file /etc/bootmenu.example into /etc/bootmenu/ and edit it to suit personal needs:

`root #` `cp /etc/bootmenu.example /etc/bootmenu
`

`root #` `nano -w /etc/bootmenu`

Below is a complete /etc/bootmenu config file. vmlinux and initrd should be replaced by the kernel and initrd image names.

DATEI **`/etc/bootmenu`** **Example bootcreator configuration file**

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

**Hinweis**

Be sure to have a look into the SmartFirmware's settings when rebooting, that menu is the file that will be loaded by default.

## Neustart des Systems

Verlassen Sie die chroot-Umgebung und hängen Sie alle gemounteten Partitionen aus. Geben Sie dann den magischen Befehl ein, der den alles entscheidenden Test einleitet - reboot.

`root #` `exit`

`livecd~#` `cd
`

`livecd~#` `umount -l /mnt/gentoo/dev{/shm,/pts,}
`

`livecd~#` `umount -R /mnt/gentoo
`

`livecd~#` `reboot`

Vergessen Sie nicht, das Installations-Medium zu entfernen. Andernfalls könnte erneut das Installations-Medium anstelle des neuen Gentoo Systems gebootet werden.

Nach dem Neustart in die neu installierte Gentoo Umgebung können Sie Ihre Installation mit dem Kapitel [Abschluss der Gentoo Installation](/wiki/Handbook:PPC/Installation/Finalizing/de "Handbook:PPC/Installation/Finalizing/de") fertigstellen.

[← Installation der Tools](/wiki/Handbook:PPC/Installation/Tools/de "Previous part") [Anfang](/wiki/Handbook:PPC/de "Handbook:PPC/de") [Abschluss der Installation →](/wiki/Handbook:PPC/Installation/Finalizing/de "Next part")