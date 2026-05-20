# Bootloader

Other languages:

- [Deutsch](/wiki/Handbook:SPARC/Installation/Bootloader/de "Handbuch:SPARC/Installation/Bootloader (100% translated)")
- [English](/wiki/Handbook:SPARC/Installation/Bootloader "Handbook:SPARC/Installation/Bootloader (100% translated)")
- [español](/wiki/Handbook:SPARC/Installation/Bootloader/es "Manual de Gentoo: SPARC/Instalación/Arranque (100% translated)")
- [français](/wiki/Handbook:SPARC/Installation/Bootloader/fr "Handbook:SPARC/Installation/Bootloader/fr (100% translated)")
- [italiano](/wiki/Handbook:SPARC/Installation/Bootloader/it "Handbook:SPARC/Installation/Bootloader/it (50% translated)")
- [magyar](/wiki/Handbook:SPARC/Installation/Bootloader/hu "Handbook:SPARC/Installation/Bootloader/hu (100% translated)")
- [polski](/wiki/Handbook:SPARC/Installation/Bootloader/pl "Handbook:SPARC/Installation/Bootloader (100% translated)")
- [português do Brasil](/wiki/Handbook:SPARC/Installation/Bootloader/pt-br "Handbook:SPARC/Installation/Bootloader/pt-br (50% translated)")
- русский
- [தமிழ்](/wiki/Handbook:SPARC/Installation/Bootloader/ta "கையேடு:SPARC/நிறுவல்/துவக்கஏற்றி (100% translated)")
- [中文（中国大陆）‎](/wiki/Handbook:SPARC/Installation/Bootloader/zh-cn "手册：SPARC/安装/配置系统引导程序Bootloader (100% translated)")
- [日本語](/wiki/Handbook:SPARC/Installation/Bootloader/ja "ハンドブック:SPARC/インストール/ブートローダー (100% translated)")
- [한국어](/wiki/Handbook:SPARC/Installation/Bootloader/ko "Handbook:SPARC/Installation/Bootloader/ko (100% translated)")

[SPARC Handbook](/wiki/Handbook:SPARC/ru "Handbook:SPARC/ru")[Установка](/wiki/Handbook:SPARC/Full/Installation/ru "Handbook:SPARC/Full/Installation/ru")[Об установке](/wiki/Handbook:SPARC/Installation/About/ru "Handbook:SPARC/Installation/About/ru")[Выбор подходящего источника для установки](/wiki/Handbook:SPARC/Installation/Media/ru "Handbook:SPARC/Installation/Media/ru")[Настройка сети](/wiki/Handbook:SPARC/Installation/Networking/ru "Handbook:SPARC/Installation/Networking/ru")[Подготовка дисков](/wiki/Handbook:SPARC/Installation/Disks/ru "Handbook:SPARC/Installation/Disks/ru")[Установка файла stage](/wiki/Handbook:SPARC/Installation/Stage/ru "Handbook:SPARC/Installation/Stage/ru")[Установка базовой системы](/wiki/Handbook:SPARC/Installation/Base/ru "Handbook:SPARC/Installation/Base/ru")[Настройка ядра](/wiki/Handbook:SPARC/Installation/Kernel/ru "Handbook:SPARC/Installation/Kernel/ru")[Настройка системы](/wiki/Handbook:SPARC/Installation/System/ru "Handbook:SPARC/Installation/System/ru")[Установка системных утилит](/wiki/Handbook:SPARC/Installation/Tools/ru "Handbook:SPARC/Installation/Tools/ru")Настройка загрузчика[Завершение](/wiki/Handbook:SPARC/Installation/Finalizing/ru "Handbook:SPARC/Installation/Finalizing/ru")[Работа с Gentoo](/wiki/Handbook:SPARC/Full/Working/ru "Handbook:SPARC/Full/Working/ru")[Введение в Portage](/wiki/Handbook:SPARC/Working/Portage/ru "Handbook:SPARC/Working/Portage/ru")[USE-флаги](/wiki/Handbook:SPARC/Working/USE/ru "Handbook:SPARC/Working/USE/ru")[Возможности Portage](/wiki/Handbook:SPARC/Working/Features/ru "Handbook:SPARC/Working/Features/ru")[Система сценариев инициализации](/wiki/Handbook:SPARC/Working/Initscripts/ru "Handbook:SPARC/Working/Initscripts/ru")[Переменные окружения](/wiki/Handbook:SPARC/Working/EnvVar/ru "Handbook:SPARC/Working/EnvVar/ru")[Работа с Portage](/wiki/Handbook:SPARC/Full/Portage/ru "Handbook:SPARC/Full/Portage/ru")[Файлы и каталоги](/wiki/Handbook:SPARC/Portage/Files/ru "Handbook:SPARC/Portage/Files/ru")[Переменные](/wiki/Handbook:SPARC/Portage/Variables/ru "Handbook:SPARC/Portage/Variables/ru")[Смешение ветвей программного обеспечения](/wiki/Handbook:SPARC/Portage/Branches/ru "Handbook:SPARC/Portage/Branches/ru")[Дополнительные утилиты](/wiki/Handbook:SPARC/Portage/Tools/ru "Handbook:SPARC/Portage/Tools/ru")[Дополнительные репозитории пакетов](/wiki/Handbook:SPARC/Portage/CustomTree/ru "Handbook:SPARC/Portage/CustomTree/ru")[Расширенные возможности](/wiki/Handbook:SPARC/Portage/Advanced/ru "Handbook:SPARC/Portage/Advanced/ru")[Настройка сети OpenRC](/wiki/Handbook:SPARC/Full/Networking/ru "Handbook:SPARC/Full/Networking/ru")[Начальная настройка](/wiki/Handbook:SPARC/Networking/Introduction/ru "Handbook:SPARC/Networking/Introduction/ru")[Расширенная настройка](/wiki/Handbook:SPARC/Networking/Advanced/ru "Handbook:SPARC/Networking/Advanced/ru")[Модульное построение сети](/wiki/Handbook:SPARC/Networking/Modular/ru "Handbook:SPARC/Networking/Modular/ru")[Беспроводная сеть](/wiki/Handbook:SPARC/Networking/Wireless/ru "Handbook:SPARC/Networking/Wireless/ru")[Добавляем функциональность](/wiki/Handbook:SPARC/Networking/Extending/ru "Handbook:SPARC/Networking/Extending/ru")[Динамическое управление](/wiki/Handbook:SPARC/Networking/Dynamic/ru "Handbook:SPARC/Networking/Dynamic/ru")

## Contents

- [1GRUB](#GRUB)
  - [1.1Emerge](#Emerge)
  - [1.2Установка](#.D0.A3.D1.81.D1.82.D0.B0.D0.BD.D0.BE.D0.B2.D0.BA.D0.B0)
    - [1.2.1GPT](#GPT)
    - [1.2.2Sun partition table](#Sun_partition_table)
  - [1.3Настройка](#.D0.9D.D0.B0.D1.81.D1.82.D1.80.D0.BE.D0.B9.D0.BA.D0.B0)
- [2SILO, начальный загрузчик для SPARC](#SILO.2C_.D0.BD.D0.B0.D1.87.D0.B0.D0.BB.D1.8C.D0.BD.D1.8B.D0.B9_.D0.B7.D0.B0.D0.B3.D1.80.D1.83.D0.B7.D1.87.D0.B8.D0.BA_.D0.B4.D0.BB.D1.8F_SPARC)
- [3Перезагрузка системы](#.D0.9F.D0.B5.D1.80.D0.B5.D0.B7.D0.B0.D0.B3.D1.80.D1.83.D0.B7.D0.BA.D0.B0_.D1.81.D0.B8.D1.81.D1.82.D0.B5.D0.BC.D1.8B)

## GRUB

When [selecting a 64-bit profile](/wiki/Handbook:SPARC/Installation/Base/ru#Choosing_the_right_profile "Handbook:SPARC/Installation/Base/ru") during installation, then [GRUB](/wiki/GRUB/ru "GRUB/ru") is the only supported bootloader.

### Emerge

GRUB should be correctly configured for the platform automatically based on the profile. To make it explicit, however, specify it using:

`root #` `echo 'GRUB_PLATFORMS="ieee1275"' >> /etc/portage/make.conf`

`root #` `emerge --ask --verbose sys-boot/grub`

The GRUB software has now been merged to the system, but not yet installed as a bootloader.

### Установка

#### GPT

If the [disk is partitioned](/wiki/Handbook:SPARC/Installation/Disks/ru#Using_fdisk_to_partition_the_disk "Handbook:SPARC/Installation/Disks/ru") using GPT (the preferred method), then install GRUB to the BIOS boot partition. Presuming the first disk (the one where the system boots from) is /dev/sda, the following commands will do:

`root #` `grub-install --target=sparc64-ieee1275 --recheck /dev/sda`

**Совет**

Чтобы найти строку загрузочного устройства для ввода в прошивку, используйте утилиту grub-ofpathname. Если загрузочный раздел BIOS находится на первом разделе диска, выберите весь диск:

`root #` `grub-ofpathname /dev/sda`

Иначе явно укажите загрузочный раздел BIOS:

`root #` `grub-ofpathname /dev/sda2`

#### Sun partition table

If the disk is partitioned using a Sun partition table instead, GRUB must be installed using blocklists. In this mode, instead of providing the physical disk as an argument, provide the path to the partition on which /boot/grub is mounted.

`root #` `grub-install --target=sparc64-ieee1275 --recheck --force --skip-fs-probe /dev/sda1`

### Настройка

Next, generate the GRUB2 configuration based on the user configuration specified in the /etc/default/grub file and /etc/grub.d scripts. In most cases, no configuration is needed by users as GRUB2 will automatically detect which kernel to boot (the highest one available in /boot/) and what the root file system is. It is also possible to append kernel parameters in /etc/default/grub using the GRUB\_CMDLINE\_LINUX variable.

To generate the final GRUB2 configuration, run the grub-mkconfig command:

`root #` `grub-mkconfig -o /boot/grub/grub.cfg`

```
Generating grub.cfg ...
Found linux image: /boot/vmlinuz-6.19.3-gentoo
Found initrd image: /boot/initramfs-genkernel-sparc-6.19.3-gentoo
done

```

The output of the command must mention that at least one Linux image is found, as those are needed to boot the system. If an initramfs is used or genkernel was used to build the kernel, the correct initrd image should be detected as well. If this is not the case, go to /boot/ and check the contents using the ls command. If the files are indeed missing, go back to the kernel configuration and installation instructions.

## SILO, начальный загрузчик для SPARC

Если был выбран [32-битный профиль](/wiki/Handbook:SPARC/Installation/Base/ru#Choosing_the_right_profile "Handbook:SPARC/Installation/Base/ru") во время установки, то единственными поддерживаемым загрузчиком является только [SILO](https://git.kernel.org/?p=linux/kernel/git/davem/silo.git;a=summary) (Sparc Improved boot LOader).

`root #` `emerge --ask sys-boot/silo`

Далее создайте /etc/silo.conf:

`root #` `nano -w /etc/silo.conf`

Ниже приведен пример файла silo.conf. В нем используется схема разбиения, которую мы применяем во всей книге, а также файл kernel-6.19.3-gentoo в качестве образа ядра и initramfs-genkernel-sparc64-6.19.3-gentoo в качестве initramfs.

ФАЙЛ **`/etc/silo.conf`** **Пример файла настройки**

```
partition = 1         # Загрузочный раздел (= корневой раздел)
root = /dev/sda1      # Корневой раздел
timeout = 150         # Ждать 15 секунд до начала загрузки раздела по умолчанию

image = /boot/kernel-6.19.3-gentoo
  label = linux
  append = "initrd=/boot/initramfs-genkernel-sparc64-6.19.3-gentoo root=/dev/sda1"

```

Если используете файлом-примером silo.conf, установленным Portage, обязательно закомментируйте все строки, которые не нужны.

Если файл /etc/silo.conf не находится на том же физическом диске, на который вы собираетесь установить SILO (в качестве загрузчика), необходимо скопировать файл /etc/silo.conf на раздел того диска, на который устанавливается загрузчик. Например, если /boot/ является отдельным разделом на этом диске, скопируйте настроечный файл в /boot/ и запустите /sbin/silo:

`root #` `cp /etc/silo.conf /boot
`

`root #` `/sbin/silo -C /boot/silo.conf`

```
/boot/silo.conf appears to be valid

```

В противном случае, просто запустите /sbin/silo:

`root #` `/sbin/silo`

```
/etc/silo.conf appears to be valid

```

**Заметка**

Запускайте silo (с параметрами) каждый раз, когда обновляете или переустанавливаете пакет [sys-boot/silo](https://packages.gentoo.org/packages/sys-boot/silo).

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

Перезагрузившись в новое окружение Gentoo, переходите к [завершению установки Gentoo](/wiki/Handbook:SPARC/Installation/Finalizing/ru "Handbook:SPARC/Installation/Finalizing/ru").

[← Установка системных средств](/wiki/Handbook:SPARC/Installation/Tools/ru "Предыдущая часть") [К содержанию](/wiki/Handbook:SPARC/ru "Handbook:SPARC/ru") [Завершение →](/wiki/Handbook:SPARC/Installation/Finalizing/ru "Следующая часть")