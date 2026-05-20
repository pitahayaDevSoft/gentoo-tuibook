# Bootloader

Other languages:

- [Deutsch](/wiki/Handbook:PPC/Installation/Bootloader/de "Handbuch:PPC/Installation/Bootloader (100% translated)")
- [English](/wiki/Handbook:PPC/Installation/Bootloader "Handbook:PPC/Installation/Bootloader (100% translated)")
- [español](/wiki/Handbook:PPC/Installation/Bootloader/es "Manual de Gentoo: PPC/Instalación/Arranque (100% translated)")
- [français](/wiki/Handbook:PPC/Installation/Bootloader/fr "Handbook:PPC/Installation/Bootloader/fr (100% translated)")
- [italiano](/wiki/Handbook:PPC/Installation/Bootloader/it "Handbook:PPC/Installation/Bootloader/it (100% translated)")
- [magyar](/wiki/Handbook:PPC/Installation/Bootloader/hu "Handbook:PPC/Installation/Bootloader/hu (100% translated)")
- [polski](/wiki/Handbook:PPC/Installation/Bootloader/pl "Handbook:PPC/Installation/Bootloader (100% translated)")
- [português do Brasil](/wiki/Handbook:PPC/Installation/Bootloader/pt-br "Handbook:PPC/Installation/Bootloader/pt-br (50% translated)")
- русский
- [தமிழ்](/wiki/Handbook:PPC/Installation/Bootloader/ta "கையேடு:PPC/நிறுவல்/துவக்கஏற்றி (100% translated)")
- [中文（中国大陆）‎](/wiki/Handbook:PPC/Installation/Bootloader/zh-cn "手册：PPC/安装/配置系统引导程序Bootloader (100% translated)")
- [日本語](/wiki/Handbook:PPC/Installation/Bootloader/ja "ハンドブック:PPC/インストール/ブートローダー (100% translated)")
- [한국어](/wiki/Handbook:PPC/Installation/Bootloader/ko "Handbook:PPC/Installation/Bootloader/ko (100% translated)")

[PPC Handbook](/wiki/Handbook:PPC/ru "Handbook:PPC/ru")[Установка](/wiki/Handbook:PPC/Full/Installation/ru "Handbook:PPC/Full/Installation/ru")[Об установке](/wiki/Handbook:PPC/Installation/About/ru "Handbook:PPC/Installation/About/ru")[Выбор подходящего источника для установки](/wiki/Handbook:PPC/Installation/Media/ru "Handbook:PPC/Installation/Media/ru")[Настройка сети](/wiki/Handbook:PPC/Installation/Networking/ru "Handbook:PPC/Installation/Networking/ru")[Подготовка дисков](/wiki/Handbook:PPC/Installation/Disks/ru "Handbook:PPC/Installation/Disks/ru")[Установка файла stage](/wiki/Handbook:PPC/Installation/Stage/ru "Handbook:PPC/Installation/Stage/ru")[Установка базовой системы](/wiki/Handbook:PPC/Installation/Base/ru "Handbook:PPC/Installation/Base/ru")[Настройка ядра](/wiki/Handbook:PPC/Installation/Kernel/ru "Handbook:PPC/Installation/Kernel/ru")[Настройка системы](/wiki/Handbook:PPC/Installation/System/ru "Handbook:PPC/Installation/System/ru")[Установка системных утилит](/wiki/Handbook:PPC/Installation/Tools/ru "Handbook:PPC/Installation/Tools/ru")Настройка загрузчика[Завершение](/wiki/Handbook:PPC/Installation/Finalizing/ru "Handbook:PPC/Installation/Finalizing/ru")[Работа с Gentoo](/wiki/Handbook:PPC/Full/Working/ru "Handbook:PPC/Full/Working/ru")[Введение в Portage](/wiki/Handbook:PPC/Working/Portage/ru "Handbook:PPC/Working/Portage/ru")[USE-флаги](/wiki/Handbook:PPC/Working/USE/ru "Handbook:PPC/Working/USE/ru")[Возможности Portage](/wiki/Handbook:PPC/Working/Features/ru "Handbook:PPC/Working/Features/ru")[Система сценариев инициализации](/wiki/Handbook:PPC/Working/Initscripts/ru "Handbook:PPC/Working/Initscripts/ru")[Переменные окружения](/wiki/Handbook:PPC/Working/EnvVar/ru "Handbook:PPC/Working/EnvVar/ru")[Работа с Portage](/wiki/Handbook:PPC/Full/Portage/ru "Handbook:PPC/Full/Portage/ru")[Файлы и каталоги](/wiki/Handbook:PPC/Portage/Files/ru "Handbook:PPC/Portage/Files/ru")[Переменные](/wiki/Handbook:PPC/Portage/Variables/ru "Handbook:PPC/Portage/Variables/ru")[Смешение ветвей программного обеспечения](/wiki/Handbook:PPC/Portage/Branches/ru "Handbook:PPC/Portage/Branches/ru")[Дополнительные утилиты](/wiki/Handbook:PPC/Portage/Tools/ru "Handbook:PPC/Portage/Tools/ru")[Дополнительные репозитории пакетов](/wiki/Handbook:PPC/Portage/CustomTree/ru "Handbook:PPC/Portage/CustomTree/ru")[Расширенные возможности](/wiki/Handbook:PPC/Portage/Advanced/ru "Handbook:PPC/Portage/Advanced/ru")[Настройка сети OpenRC](/wiki/Handbook:PPC/Full/Networking/ru "Handbook:PPC/Full/Networking/ru")[Начальная настройка](/wiki/Handbook:PPC/Networking/Introduction/ru "Handbook:PPC/Networking/Introduction/ru")[Расширенная настройка](/wiki/Handbook:PPC/Networking/Advanced/ru "Handbook:PPC/Networking/Advanced/ru")[Модульное построение сети](/wiki/Handbook:PPC/Networking/Modular/ru "Handbook:PPC/Networking/Modular/ru")[Беспроводная сеть](/wiki/Handbook:PPC/Networking/Wireless/ru "Handbook:PPC/Networking/Wireless/ru")[Добавляем функциональность](/wiki/Handbook:PPC/Networking/Extending/ru "Handbook:PPC/Networking/Extending/ru")[Динамическое управление](/wiki/Handbook:PPC/Networking/Dynamic/ru "Handbook:PPC/Networking/Dynamic/ru")

## Contents

- [1Параметры загрузчика](#.D0.9F.D0.B0.D1.80.D0.B0.D0.BC.D0.B5.D1.82.D1.80.D1.8B_.D0.B7.D0.B0.D0.B3.D1.80.D1.83.D0.B7.D1.87.D0.B8.D0.BA.D0.B0)
- [2NewWorld Macs](#NewWorld_Macs)
  - [2.1GRUB](#GRUB)
  - [2.2Установка](#.D0.A3.D1.81.D1.82.D0.B0.D0.BD.D0.BE.D0.B2.D0.BA.D0.B0)
  - [2.3Setup bootstrap partition](#Setup_bootstrap_partition)
  - [2.4Установка GRUB](#.D0.A3.D1.81.D1.82.D0.B0.D0.BD.D0.BE.D0.B2.D0.BA.D0.B0_GRUB)
- [3OldWorld Macs](#OldWorld_Macs)
  - [3.1BootX](#BootX)
- [4Pegasos](#Pegasos)
  - [4.1BootCreator](#BootCreator)
- [5Перезагрузка системы](#.D0.9F.D0.B5.D1.80.D0.B5.D0.B7.D0.B0.D0.B3.D1.80.D1.83.D0.B7.D0.BA.D0.B0_.D1.81.D0.B8.D1.81.D1.82.D0.B5.D0.BC.D1.8B)

## Параметры загрузчика

Now that the kernel is configured and compiled and the necessary system configuration files are filled in correctly, it is time to install a program that will fire up the kernel when the system is started. Such a program is called a boot loader.

The boot loader to use depends upon the type of PPC machine.

For a NewWorld Apple or IBM machine, grub needs to be selected. OldWorld Apple machines havs one option: BootX. The Pegasos does not require a boot loader, but it is necessary to emerge bootcreator to create SmartFirmware boot menus.

## NewWorld Macs

### GRUB

### Установка

`root #` `emerge --ask sys-boot/grub`

### Setup bootstrap partition

First, prepare the bootstrap partition that was created created during the [preparing the disk](/wiki/Handbook:PPC64/Installation/Disks "Handbook:PPC64/Installation/Disks") step. Following the example, this partition should be /dev/sda2. Optionally, confirm this by using parted:

Replace /dev/sda with the correct device if required.

`root #` `parted /dev/sda print`

```
Model: ATA Patriot Burst El (scsi)

Disk /dev/sda: 120GB

Sector size (logical/physical): 512B/512B

Partition Table: mac

Disk Flags:

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

### Установка GRUB

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

**Важно**

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

**Важно**

Make sure to include support for the HFS and HFS+ filesystems in the kernel, otherwise upgrades or changes to the kernel on the MacOS partition will not be possible.

## Pegasos

**Заметка**

Pegasos also has Grub support but this is currently undocumented in Gentoo. Please add this to the main wiki and notify on this page's discussion once ready to migrated here.

### BootCreator

**Важно**

BootCreator will build a nice SmartFirmware bootmenu written in Forth for the Pegasos.

First make sure to have bootcreator installed on the system:

`root #` `emerge --ask sys-boot/bootcreator`

Now copy the file /etc/bootmenu.example into /etc/bootmenu/ and edit it to suit personal needs:

`root #` `cp /etc/bootmenu.example /etc/bootmenu
`

`root #` `nano -w /etc/bootmenu`

Below is a complete /etc/bootmenu config file. vmlinux and initrd should be replaced by the kernel and initrd image names.

ФАЙЛ **`/etc/bootmenu`** **Пример конфигурационного файла для bootcreator**

```
#
# Пример файла описаний для bootcreator 1.1
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

**Заметка**

Be sure to have a look into the SmartFirmware's settings when rebooting, that menu is the file that will be loaded by default.

## Перезагрузка системы

Выйдите из изолированной среды и размонтируйте все смонтированные разделы. Затем введите ту самую волшебную команду, которая запускает последний, настоящий тест: reboot.

`(chroot) livecd #` `exit`

`livecd~#` `cd
`

`livecd~#` `umount -l /mnt/gentoo/dev{/shm,/pts,}
`

`livecd~#` `umount -R /mnt/gentoo
`

`livecd~#` `reboot`

Не забудьте извлечь загрузочный компакт-диск, иначе он может загрузиться снова вместо новой системы Gentoo!

Перезагрузившись в новое окружение Gentoo, переходите к [завершению установки Gentoo](/wiki/Handbook:PPC/Installation/Finalizing/ru "Handbook:PPC/Installation/Finalizing/ru").

[← Установка системных средств](/wiki/Handbook:PPC/Installation/Tools/ru "Предыдущая часть") [К содержанию](/wiki/Handbook:PPC/ru "Handbook:PPC/ru") [Завершение →](/wiki/Handbook:PPC/Installation/Finalizing/ru "Следующая часть")