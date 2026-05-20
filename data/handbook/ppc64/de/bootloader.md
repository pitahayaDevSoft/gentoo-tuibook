# Bootloader

Other languages:

- Deutsch
- [English](/wiki/Handbook:PPC64/Installation/Bootloader "Handbook:PPC64/Installation/Bootloader (100% translated)")
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

[PPC64 Handbuch](/wiki/Handbook:PPC64/de "Handbook:PPC64/de")[Installation](/wiki/Handbook:PPC64/Full/Installation/de "Handbook:PPC64/Full/Installation/de")[Über die Installation](/wiki/Handbook:PPC64/Installation/About/de "Handbook:PPC64/Installation/About/de")[Auswahl des Mediums](/wiki/Handbook:PPC64/Installation/Media/de "Handbook:PPC64/Installation/Media/de")[Konfiguration des Netzwerks](/wiki/Handbook:PPC64/Installation/Networking/de "Handbook:PPC64/Installation/Networking/de")[Vorbereiten der Festplatte(n)](/wiki/Handbook:PPC64/Installation/Disks/de "Handbook:PPC64/Installation/Disks/de")[Installation des Stage Archivs](/wiki/Handbook:PPC64/Installation/Stage/de "Handbook:PPC64/Installation/Stage/de")[Installation des Basissystems](/wiki/Handbook:PPC64/Installation/Base/de "Handbook:PPC64/Installation/Base/de")[Konfiguration des Kernels](/wiki/Handbook:PPC64/Installation/Kernel/de "Handbook:PPC64/Installation/Kernel/de")[Konfiguration des Systems](/wiki/Handbook:PPC64/Installation/System/de "Handbook:PPC64/Installation/System/de")[Installation der Tools](/wiki/Handbook:PPC64/Installation/Tools/de "Handbook:PPC64/Installation/Tools/de")Konfiguration des Bootloaders[Abschluss](/wiki/Handbook:PPC64/Installation/Finalizing/de "Handbook:PPC64/Installation/Finalizing/de")[Arbeiten mit Gentoo](/wiki/Handbook:PPC64/Full/Working/de "Handbook:PPC64/Full/Working/de")[Portage-Einführung](/wiki/Handbook:PPC64/Working/Portage/de "Handbook:PPC64/Working/Portage/de")[USE-Flags](/wiki/Handbook:PPC64/Working/USE/de "Handbook:PPC64/Working/USE/de")[Portage-Features](/wiki/Handbook:PPC64/Working/Features/de "Handbook:PPC64/Working/Features/de")[Initskript-System](/wiki/Handbook:PPC64/Working/Initscripts/de "Handbook:PPC64/Working/Initscripts/de")[Umgebungsvariablen](/wiki/Handbook:PPC64/Working/EnvVar/de "Handbook:PPC64/Working/EnvVar/de")[Arbeiten mit Portage](/wiki/Handbook:PPC64/Full/Portage/de "Handbook:PPC64/Full/Portage/de")[Dateien und Verzeichnisse](/wiki/Handbook:PPC64/Portage/Files/de "Handbook:PPC64/Portage/Files/de")[Variablen](/wiki/Handbook:PPC64/Portage/Variables/de "Handbook:PPC64/Portage/Variables/de")[Mischen von Softwarezweigen](/wiki/Handbook:PPC64/Portage/Branches/de "Handbook:PPC64/Portage/Branches/de")[Zusätzliche Tools](/wiki/Handbook:PPC64/Portage/Tools/de "Handbook:PPC64/Portage/Tools/de")[Eigener Portage-Tree](/wiki/Handbook:PPC64/Portage/CustomTree/de "Handbook:PPC64/Portage/CustomTree/de")[Erweiterte Portage-Features](/wiki/Handbook:PPC64/Portage/Advanced/de "Handbook:PPC64/Portage/Advanced/de")[Netzwerk-Konfiguration](/wiki/Handbook:PPC64/Full/Networking/de "Handbook:PPC64/Full/Networking/de")[Zu Beginn](/wiki/Handbook:PPC64/Networking/Introduction/de "Handbook:PPC64/Networking/Introduction/de")[Fortgeschrittene Konfiguration](/wiki/Handbook:PPC64/Networking/Advanced/de "Handbook:PPC64/Networking/Advanced/de")[Modulare Vernetzung](/wiki/Handbook:PPC64/Networking/Modular/de "Handbook:PPC64/Networking/Modular/de")[Drahtlose Netzwerke](/wiki/Handbook:PPC64/Networking/Wireless/de "Handbook:PPC64/Networking/Wireless/de")[Funktionalität hinzufügen](/wiki/Handbook:PPC64/Networking/Extending/de "Handbook:PPC64/Networking/Extending/de")[Dynamisches Management](/wiki/Handbook:PPC64/Networking/Dynamic/de "Handbook:PPC64/Networking/Dynamic/de")

With the kernel configured and compiled and the necessary system configuration files filled in correctly, it is time to install a program that will fire up the kernel when the system boots. Such a program is called a boot loader.

**Hinweis**

Currently using Petitboot on Talos systems is undocumented in Gentoo. Please add the steps to [TalosII#Bootloader](/wiki/TalosII#Bootloader.2Fde "TalosII") and notify on this Discussion page when ready to merge into the Handbook.

## Contents

- [1Using GRUB](#Using_GRUB)
  - [1.1Installation](#Installation)
  - [1.2Mac hardware (G5)](#Mac_hardware_.28G5.29)
    - [1.2.1Setup bootstrap partition](#Setup_bootstrap_partition)
    - [1.2.2Setup GRUB](#Setup_GRUB)
  - [1.3IBM hardware](#IBM_hardware)
    - [1.3.1Setup GRUB](#Setup_GRUB_2)
    - [1.3.2Grub config](#Grub_config)
- [2Neustart des Systems](#Neustart_des_Systems)

## Using GRUB

GRUB is a bootloader for PPC64 powered Linux machines.

### Installation

`root #` `emerge --ask sys-boot/grub`

### Mac hardware (G5)

#### Setup bootstrap partition

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

#### Setup GRUB

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

### IBM hardware

Setting up Grub on IBM hardware is as simple as:

#### Setup GRUB

`root #` `grub-install /dev/sda1`

**Hinweis**

/dev/sda1 is the PReP boot partition made in the partitioning stage

#### Grub config

Finally. build the grub.cfg file:

`root #` `grub-mkconfig -o /boot/grub/grub.cfg`

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

Nach dem Neustart in die neu installierte Gentoo Umgebung können Sie Ihre Installation mit dem Kapitel [Abschluss der Gentoo Installation](/wiki/Handbook:PPC64/Installation/Finalizing/de "Handbook:PPC64/Installation/Finalizing/de") fertigstellen.

[← Installation der Tools](/wiki/Handbook:PPC64/Installation/Tools/de "Previous part") [Anfang](/wiki/Handbook:PPC64/de "Handbook:PPC64/de") [Abschluss der Installation →](/wiki/Handbook:PPC64/Installation/Finalizing/de "Next part")