# Bootloader

Other languages:

- [Deutsch](/wiki/Handbook:PPC64/Installation/Bootloader/de "Handbuch:PPC64/Installation/Bootloader (100% translated)")
- [English](/wiki/Handbook:PPC64/Installation/Bootloader "Handbook:PPC64/Installation/Bootloader (100% translated)")
- [español](/wiki/Handbook:PPC64/Installation/Bootloader/es "Manual de Gentoo: PPC64/Instalación/Arranque (100% translated)")
- [français](/wiki/Handbook:PPC64/Installation/Bootloader/fr "Handbook:PPC64/Installation/Bootloader/fr (100% translated)")
- [italiano](/wiki/Handbook:PPC64/Installation/Bootloader/it "Handbook:PPC64/Installation/Bootloader/it (100% translated)")
- [magyar](/wiki/Handbook:PPC64/Installation/Bootloader/hu "Handbook:PPC64/Installation/Bootloader/hu (100% translated)")
- [polski](/wiki/Handbook:PPC64/Installation/Bootloader/pl "Handbook:PPC64/Installation/Bootloader (100% translated)")
- [português do Brasil](/wiki/Handbook:PPC64/Installation/Bootloader/pt-br "Handbook:PPC64/Installation/Bootloader/pt-br (50% translated)")
- русский
- [தமிழ்](/wiki/Handbook:PPC64/Installation/Bootloader/ta "கையேடு:PPC64/நிறுவல்/துவக்கஏற்றி (100% translated)")
- [中文（中国大陆）‎](/wiki/Handbook:PPC64/Installation/Bootloader/zh-cn "手册：PPC64/安装/配置系统引导程序Bootloader (100% translated)")
- [日本語](/wiki/Handbook:PPC64/Installation/Bootloader/ja "ハンドブック:PPC64/インストール/ブートローダー (100% translated)")
- [한국어](/wiki/Handbook:PPC64/Installation/Bootloader/ko "Handbook:PPC64/Installation/Bootloader/ko (100% translated)")

[PPC64 Handbook](/wiki/Handbook:PPC64/ru "Handbook:PPC64/ru")[Установка](/wiki/Handbook:PPC64/Full/Installation/ru "Handbook:PPC64/Full/Installation/ru")[Об установке](/wiki/Handbook:PPC64/Installation/About/ru "Handbook:PPC64/Installation/About/ru")[Выбор подходящего источника для установки](/wiki/Handbook:PPC64/Installation/Media/ru "Handbook:PPC64/Installation/Media/ru")[Настройка сети](/wiki/Handbook:PPC64/Installation/Networking/ru "Handbook:PPC64/Installation/Networking/ru")[Подготовка дисков](/wiki/Handbook:PPC64/Installation/Disks/ru "Handbook:PPC64/Installation/Disks/ru")[Установка файла stage](/wiki/Handbook:PPC64/Installation/Stage/ru "Handbook:PPC64/Installation/Stage/ru")[Установка базовой системы](/wiki/Handbook:PPC64/Installation/Base/ru "Handbook:PPC64/Installation/Base/ru")[Настройка ядра](/wiki/Handbook:PPC64/Installation/Kernel/ru "Handbook:PPC64/Installation/Kernel/ru")[Настройка системы](/wiki/Handbook:PPC64/Installation/System/ru "Handbook:PPC64/Installation/System/ru")[Установка системных утилит](/wiki/Handbook:PPC64/Installation/Tools/ru "Handbook:PPC64/Installation/Tools/ru")Настройка загрузчика[Завершение](/wiki/Handbook:PPC64/Installation/Finalizing/ru "Handbook:PPC64/Installation/Finalizing/ru")[Работа с Gentoo](/wiki/Handbook:PPC64/Full/Working/ru "Handbook:PPC64/Full/Working/ru")[Введение в Portage](/wiki/Handbook:PPC64/Working/Portage/ru "Handbook:PPC64/Working/Portage/ru")[USE-флаги](/wiki/Handbook:PPC64/Working/USE/ru "Handbook:PPC64/Working/USE/ru")[Возможности Portage](/wiki/Handbook:PPC64/Working/Features/ru "Handbook:PPC64/Working/Features/ru")[Система сценариев инициализации](/wiki/Handbook:PPC64/Working/Initscripts/ru "Handbook:PPC64/Working/Initscripts/ru")[Переменные окружения](/wiki/Handbook:PPC64/Working/EnvVar/ru "Handbook:PPC64/Working/EnvVar/ru")[Работа с Portage](/wiki/Handbook:PPC64/Full/Portage/ru "Handbook:PPC64/Full/Portage/ru")[Файлы и каталоги](/wiki/Handbook:PPC64/Portage/Files/ru "Handbook:PPC64/Portage/Files/ru")[Переменные](/wiki/Handbook:PPC64/Portage/Variables/ru "Handbook:PPC64/Portage/Variables/ru")[Смешение ветвей программного обеспечения](/wiki/Handbook:PPC64/Portage/Branches/ru "Handbook:PPC64/Portage/Branches/ru")[Дополнительные утилиты](/wiki/Handbook:PPC64/Portage/Tools/ru "Handbook:PPC64/Portage/Tools/ru")[Дополнительные репозитории пакетов](/wiki/Handbook:PPC64/Portage/CustomTree/ru "Handbook:PPC64/Portage/CustomTree/ru")[Расширенные возможности](/wiki/Handbook:PPC64/Portage/Advanced/ru "Handbook:PPC64/Portage/Advanced/ru")[Настройка сети OpenRC](/wiki/Handbook:PPC64/Full/Networking/ru "Handbook:PPC64/Full/Networking/ru")[Начальная настройка](/wiki/Handbook:PPC64/Networking/Introduction/ru "Handbook:PPC64/Networking/Introduction/ru")[Расширенная настройка](/wiki/Handbook:PPC64/Networking/Advanced/ru "Handbook:PPC64/Networking/Advanced/ru")[Модульное построение сети](/wiki/Handbook:PPC64/Networking/Modular/ru "Handbook:PPC64/Networking/Modular/ru")[Беспроводная сеть](/wiki/Handbook:PPC64/Networking/Wireless/ru "Handbook:PPC64/Networking/Wireless/ru")[Добавляем функциональность](/wiki/Handbook:PPC64/Networking/Extending/ru "Handbook:PPC64/Networking/Extending/ru")[Динамическое управление](/wiki/Handbook:PPC64/Networking/Dynamic/ru "Handbook:PPC64/Networking/Dynamic/ru")

With the kernel configured and compiled and the necessary system configuration files filled in correctly, it is time to install a program that will fire up the kernel when the system boots. Such a program is called a boot loader.

**Заметка**

Currently using Petitboot on Talos systems is undocumented in Gentoo. Please add the steps to [TalosII#Bootloader](/wiki/TalosII#Bootloader.2Fru "TalosII") and notify on this Discussion page when ready to merge into the Handbook.

## Contents

- [1Использование GRUB](#.D0.98.D1.81.D0.BF.D0.BE.D0.BB.D1.8C.D0.B7.D0.BE.D0.B2.D0.B0.D0.BD.D0.B8.D0.B5_GRUB)
  - [1.1Установка](#.D0.A3.D1.81.D1.82.D0.B0.D0.BD.D0.BE.D0.B2.D0.BA.D0.B0)
  - [1.2Mac hardware (G5)](#Mac_hardware_.28G5.29)
    - [1.2.1Setup bootstrap partition](#Setup_bootstrap_partition)
    - [1.2.2Установка GRUB](#.D0.A3.D1.81.D1.82.D0.B0.D0.BD.D0.BE.D0.B2.D0.BA.D0.B0_GRUB)
  - [1.3IBM hardware](#IBM_hardware)
    - [1.3.1Установка GRUB](#.D0.A3.D1.81.D1.82.D0.B0.D0.BD.D0.BE.D0.B2.D0.BA.D0.B0_GRUB_2)
    - [1.3.2Grub config](#Grub_config)
- [2Перезагрузка системы](#.D0.9F.D0.B5.D1.80.D0.B5.D0.B7.D0.B0.D0.B3.D1.80.D1.83.D0.B7.D0.BA.D0.B0_.D1.81.D0.B8.D1.81.D1.82.D0.B5.D0.BC.D1.8B)

## Использование GRUB

GRUB is a bootloader for PPC64 powered Linux machines.

### Установка

`root #` `emerge --ask sys-boot/grub`

### Mac hardware (G5)

#### Setup bootstrap partition

First, prepare the bootstrap partition that was created created during the [preparing the disk](/wiki/Handbook:PPC64/Installation/Disks "Handbook:PPC64/Installation/Disks") step. Following the example, this partition should be /dev/sda2. Optionally, confirm this by using parted:

Replace /dev/sda with the correct device if required.

`root #` `Вывод parted /dev/sda`

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

#### Установка GRUB

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

#### Установка GRUB

`root #` `grub-install /dev/sda1`

**Заметка**

/dev/sda1 is the PReP boot partition made in the partitioning stage

#### Grub config

Finally. build the grub.cfg file:

`root #` `grub-mkconfig -o /boot/grub/grub.cfg`

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

Перезагрузившись в новое окружение Gentoo, переходите к [завершению установки Gentoo](/wiki/Handbook:PPC64/Installation/Finalizing/ru "Handbook:PPC64/Installation/Finalizing/ru").

[← Установка системных средств](/wiki/Handbook:PPC64/Installation/Tools/ru "Предыдущая часть") [К содержанию](/wiki/Handbook:PPC64/ru "Handbook:PPC64/ru") [Завершение →](/wiki/Handbook:PPC64/Installation/Finalizing/ru "Следующая часть")