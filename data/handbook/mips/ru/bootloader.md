# Bootloader

Other languages:

- [Deutsch](/wiki/Handbook:MIPS/Installation/Bootloader/de "Handbuch:MIPS/Installation/Bootloader (100% translated)")
- [English](/wiki/Handbook:MIPS/Installation/Bootloader "Handbook:MIPS/Installation/Bootloader (100% translated)")
- [español](/wiki/Handbook:MIPS/Installation/Bootloader/es "Manual de Gentoo: MIPS/Instalación/Arranque (100% translated)")
- [français](/wiki/Handbook:MIPS/Installation/Bootloader/fr "Handbook:MIPS/Installation/Bootloader/fr (100% translated)")
- [italiano](/wiki/Handbook:MIPS/Installation/Bootloader/it "Handbook:MIPS/Installation/Bootloader/it (100% translated)")
- [magyar](/wiki/Handbook:MIPS/Installation/Bootloader/hu "Handbook:MIPS/Installation/Bootloader/hu (100% translated)")
- [polski](/wiki/Handbook:MIPS/Installation/Bootloader/pl "Handbook:MIPS/Installation/Bootloader (100% translated)")
- [português do Brasil](/wiki/Handbook:MIPS/Installation/Bootloader/pt-br "Manual:MIPS/Instalação/Gerenciador de Boot (100% translated)")
- русский
- [தமிழ்](/wiki/Handbook:MIPS/Installation/Bootloader/ta "கையேடு:MIPS/நிறுவல்/துவக்கஏற்றி (100% translated)")
- [中文（中国大陆）‎](/wiki/Handbook:MIPS/Installation/Bootloader/zh-cn "手册：MIPS/安装/配置系统引导程序Bootloader (100% translated)")
- [日本語](/wiki/Handbook:MIPS/Installation/Bootloader/ja "ハンドブック:MIPS/インストール/ブートローダー (100% translated)")
- [한국어](/wiki/Handbook:MIPS/Installation/Bootloader/ko "Handbook:MIPS/Installation/Bootloader/ko (100% translated)")

[MIPS Handbook](/wiki/Handbook:MIPS/ru "Handbook:MIPS/ru")[Установка](/wiki/Handbook:MIPS/Full/Installation/ru "Handbook:MIPS/Full/Installation/ru")[Об установке](/wiki/Handbook:MIPS/Installation/About/ru "Handbook:MIPS/Installation/About/ru")[Выбор подходящего источника для установки](/wiki/Handbook:MIPS/Installation/Media/ru "Handbook:MIPS/Installation/Media/ru")[Настройка сети](/wiki/Handbook:MIPS/Installation/Networking/ru "Handbook:MIPS/Installation/Networking/ru")[Подготовка дисков](/wiki/Handbook:MIPS/Installation/Disks/ru "Handbook:MIPS/Installation/Disks/ru")[Установка файла stage](/wiki/Handbook:MIPS/Installation/Stage/ru "Handbook:MIPS/Installation/Stage/ru")[Установка базовой системы](/wiki/Handbook:MIPS/Installation/Base/ru "Handbook:MIPS/Installation/Base/ru")[Настройка ядра](/wiki/Handbook:MIPS/Installation/Kernel/ru "Handbook:MIPS/Installation/Kernel/ru")[Настройка системы](/wiki/Handbook:MIPS/Installation/System/ru "Handbook:MIPS/Installation/System/ru")[Установка системных утилит](/wiki/Handbook:MIPS/Installation/Tools/ru "Handbook:MIPS/Installation/Tools/ru")Настройка загрузчика[Завершение](/wiki/Handbook:MIPS/Installation/Finalizing/ru "Handbook:MIPS/Installation/Finalizing/ru")[Работа с Gentoo](/wiki/Handbook:MIPS/Full/Working/ru "Handbook:MIPS/Full/Working/ru")[Введение в Portage](/wiki/Handbook:MIPS/Working/Portage/ru "Handbook:MIPS/Working/Portage/ru")[USE-флаги](/wiki/Handbook:MIPS/Working/USE/ru "Handbook:MIPS/Working/USE/ru")[Возможности Portage](/wiki/Handbook:MIPS/Working/Features/ru "Handbook:MIPS/Working/Features/ru")[Система сценариев инициализации](/wiki/Handbook:MIPS/Working/Initscripts/ru "Handbook:MIPS/Working/Initscripts/ru")[Переменные окружения](/wiki/Handbook:MIPS/Working/EnvVar/ru "Handbook:MIPS/Working/EnvVar/ru")[Работа с Portage](/wiki/Handbook:MIPS/Full/Portage/ru "Handbook:MIPS/Full/Portage/ru")[Файлы и каталоги](/wiki/Handbook:MIPS/Portage/Files/ru "Handbook:MIPS/Portage/Files/ru")[Переменные](/wiki/Handbook:MIPS/Portage/Variables/ru "Handbook:MIPS/Portage/Variables/ru")[Смешение ветвей программного обеспечения](/wiki/Handbook:MIPS/Portage/Branches/ru "Handbook:MIPS/Portage/Branches/ru")[Дополнительные утилиты](/wiki/Handbook:MIPS/Portage/Tools/ru "Handbook:MIPS/Portage/Tools/ru")[Дополнительные репозитории пакетов](/wiki/Handbook:MIPS/Portage/CustomTree/ru "Handbook:MIPS/Portage/CustomTree/ru")[Расширенные возможности](/wiki/Handbook:MIPS/Portage/Advanced/ru "Handbook:MIPS/Portage/Advanced/ru")[Настройка сети OpenRC](/wiki/Handbook:MIPS/Full/Networking/ru "Handbook:MIPS/Full/Networking/ru")[Начальная настройка](/wiki/Handbook:MIPS/Networking/Introduction/ru "Handbook:MIPS/Networking/Introduction/ru")[Расширенная настройка](/wiki/Handbook:MIPS/Networking/Advanced/ru "Handbook:MIPS/Networking/Advanced/ru")[Модульное построение сети](/wiki/Handbook:MIPS/Networking/Modular/ru "Handbook:MIPS/Networking/Modular/ru")[Беспроводная сеть](/wiki/Handbook:MIPS/Networking/Wireless/ru "Handbook:MIPS/Networking/Wireless/ru")[Добавляем функциональность](/wiki/Handbook:MIPS/Networking/Extending/ru "Handbook:MIPS/Networking/Extending/ru")[Динамическое управление](/wiki/Handbook:MIPS/Networking/Dynamic/ru "Handbook:MIPS/Networking/Dynamic/ru")

## Contents

- [1arcload для машин с Silicon Graphics](#arcload_.D0.B4.D0.BB.D1.8F_.D0.BC.D0.B0.D1.88.D0.B8.D0.BD_.D1.81_Silicon_Graphics)
- [2CoLo для Cobalt MicroServers](#CoLo_.D0.B4.D0.BB.D1.8F_Cobalt_MicroServers)
  - [2.1Установка CoLo](#.D0.A3.D1.81.D1.82.D0.B0.D0.BD.D0.BE.D0.B2.D0.BA.D0.B0_CoLo)
  - [2.2Настройка CoLo](#.D0.9D.D0.B0.D1.81.D1.82.D1.80.D0.BE.D0.B9.D0.BA.D0.B0_CoLo)
- [3Настройка последовательной консоли](#.D0.9D.D0.B0.D1.81.D1.82.D1.80.D0.BE.D0.B9.D0.BA.D0.B0_.D0.BF.D0.BE.D1.81.D0.BB.D0.B5.D0.B4.D0.BE.D0.B2.D0.B0.D1.82.D0.B5.D0.BB.D1.8C.D0.BD.D0.BE.D0.B9_.D0.BA.D0.BE.D0.BD.D1.81.D0.BE.D0.BB.D0.B8)
- [4Tweaking the SGI PROM](#Tweaking_the_SGI_PROM)
  - [4.1Setting generic PROM settings](#Setting_generic_PROM_settings)
  - [4.2Settings for direct volume-header booting](#Settings_for_direct_volume-header_booting)
  - [4.3Настройки для arcload](#.D0.9D.D0.B0.D1.81.D1.82.D1.80.D0.BE.D0.B9.D0.BA.D0.B8_.D0.B4.D0.BB.D1.8F_arcload)
- [5Перезагрузка системы](#.D0.9F.D0.B5.D1.80.D0.B5.D0.B7.D0.B0.D0.B3.D1.80.D1.83.D0.B7.D0.BA.D0.B0_.D1.81.D0.B8.D1.81.D1.82.D0.B5.D0.BC.D1.8B)

## arcload для машин с Silicon Graphics

arcload был написан для компьютеров, для которых требовалась загрузка 64-битных ядер и для которых невозможно использовать arcboot (который очень трудно скомпилировать в 64-битный исполняемый файл). Этот загрузчик также учитывает особенности загрузки ядер напрямую из заголовка тома. Чтобы начать установку, наберите:

`root #` `emerge arcload dvhtool`

После установки исполняемый файл arcload будет находиться в /usr/lib/arcload/. Также есть два файла:

- sashARCS: 32-битный двоичный файл для систем Indy, Indigo2 (R4k), Challenge S и O2.
- sash64: 64-битный двоичный файл для систем Octane/Octane2, Origin 200/2000 и Indigo2 Impact.

Используйте dvhtool для установки соответствующего файла в заголовок тома:

Для пользователей Indy/Indigo2/Challenge S/O2:

`root #` `dvhtool --unix-to-vh /usr/lib/arcload/sashARCS sashARCS`

Для пользователей Indigo2 Impact/Octane/Octane2/Origin 200/Origin 2000:

`root #` `dvhtool --unix-to-vh /usr/lib/arcload/sash64 sash64`

**Заметка**

Необязательно использовать имена sashARCS или sash64 (если только не выполняется установка в заголовок тома загрузочного компакт-диска). Для нормальной загрузки с жёсткого диска можно использовать любое имя.

Теперь запустите dvhtool, чтобы удостовериться, что заголовок тома записан:

`root #` `dvhtool --print-volume-directory`

```
----- directory entries -----
Entry #0, name "sash64", start 4, bytes 55859

```

Файл arc.cf имеет Си-подобный синтаксис. Для подробного описания о том, как его настраивать, обратитесь к странице arcload на Linux/MIPS wiki. Вкратце, необходимо определить несколько параметров, которые могут быть включены или выключены во время загрузки при помощи переменной OSLoadFilename.

ФАЙЛ **`arc.cf`** **Пример arc.cf**

```
# Конфигурация ARCLoad

# Значения по умолчанию...
append  "root=/dev/sda5";
append  "ro";
append  "console=ttyS0,9600";

# Главная запись. ip28 может быть изменено по собственному усмотрению.
ip28 {
        # Определение "рабочего" ядра
        # Выберите его, установив OSLoadFilename="ip28(working)"
        working {
                description     "SGI Indigo2 Impact R10000\n\r";
                image system    "/working";
        }

        # Определение "нового" ядра
        # Выберите его, установив OSLoadFilename="ip28(new)"
        new {
                description     "SGI Indigo2 Impact R10000 - Testing Kernel\n\r";
                image system    "/new";
        }

        # Отладка ядра
        # Выберите его, установив OSLoadFilename="ip28(working,debug)"
        # или OSLoadFilename="ip28(new,debug)"
        debug {
                description     "Debug console";
                append          "init=/bin/bash";
        }
}

```

Начиная с arcload-0.5, arc.cf и ядра могут находиться как в заголовке тома, так и в разделе. Чтобы воспользоваться этой новой возможностью, разместите файлы в разделе /boot/ (или /, если /boot не является отдельным разделом). arcload будет использовать код драйвера файловой системы от популярного загрузчика grub, благодаря чему он поддерживает тот же набор файловых систем.

`root #` `dvhtool --unix-to-vh arc.cf arc.cf
`

`root #` `dvhtool --unix-to-vh /usr/src/linux/vmlinux new`

## CoLo для Cobalt MicroServers

### Установка CoLo

На серверах Cobalt предустановленная прошивка более ограничена в возможностях. Возможности Cobalt BOOTROM примитивны по сравнению с SGI PROM и имеют серьёзные ограничения.

- Размер ядра ограничен 675 КБ (примерно). Размер Linux 2.4 делает практически невозможным создание ядра такого размера. Linux 2.6 и 3.x и вовсе можно не рассматривать.
- Предустановленная прошивка не поддерживает загрузку 64-битных ядер (хотя они экспериментальные для компьютеров Cobalt на данный момент).
- Поддержка командной оболочки в лучшем случае очень базовая.

Чтобы преодолеть эти ограничения, была разработана альтернативная прошивка под названием CoLo (Cobalt Loader). Она представляет собой образ BOOTROM, который можно либо прошить в микросхему сервера Cobalt, либо загрузить из существующей прошивки.

**Заметка**

В данном руководстве будет описан процесс настройки CoLo для его загрузки из существующей прошивки. Это самый безопасный и рекомендуемый способ использования CoLo.

**Предупреждение**

При желании, можно прошить CoLo на сервер, заменив оригинальную прошивку, однако это придётся сделать вам самостоятельно. Если что-то пойдёт не так, придётся физически извлечь BOOTROM и перепрограммировать его оригинальной прошивкой. Если это звучит страшно и непонятно — НЕ ПЕРЕПРОШИВАЙТЕ компьютер. Gentoo не отвечает ни за какие последствия, если этот совет будет проигнорирован.

Теперь установим CoLo. Соберите соответствующий пакет:

`root #` `emerge --ask sys-boot/colo`

После установке загляните внутрь каталога /usr/lib/colo/, там будет два файла:

- colo-chain.elf — «ядро» для загрузки оригинальной прошивкой.
- colo-rom-image.bin — образ ROM для прошивки в BOOTROM.

Смонтируем /boot/ и разместим там сжатую копию colo-chain.elf таким образом, каким ожидает увидеть её система.

`root #` `gzip -9vc /usr/lib/colo/colo-chain.elf > /boot/vmlinux.gz`

### Настройка CoLo

Now, when the system first boots up, it'll load CoLo which will spit up a menu on the back LCD. The first option (and default that is assumed after roughly 5 seconds) is to boot to the hard disk. The system would then attempt to mount the first Linux partition it finds, and run the script default.colo. The syntax is fully documented in the CoLo documentation (have a peek at /usr/share/doc/colo-X.YY/README.shell.gz \-\- where X.YY is the version installed), and is very simple.

**Заметка**

Just a tip: when installing kernels, it is recommended to create two kernel images, kernel.gz.working -- a known working kernel, and kernel.gz.new -- a kernel that's just been compiled. It is possible to use symlinks to point to the curent "new" and "working" kernels, or just rename the kernel images.

ФАЙЛ **`default.colo`** **An example CoLo configuration**

```
#:CoLo:#
mount hda1
load /kernel.gz.working
execute root=/dev/sda5 ro console=ttyS0,115200

```

**Заметка**

CoLo will refuse to load a script that does not begin with the `#:CoLo:#` line. Think of it as the equivalent of saying `#!/bin/sh` in shell scripts.

It is also possible to ask a question, such as which kernel & configuration to boot, with a default timeout. The following configuration does exactly this, asks the user which kernel they wish to use, and executes the chosen image. vmlinux.gz.new and vmlinux.gz.working may be actual kernel images, or just symlinks pointing to the kernel images on that disk. The 50 argument to select specifies that it should proceed with the first option ("Working") after 50/10 seconds.

ФАЙЛ **`default.colo`** **Menu-based configuration**

```
#:CoLo:#
lcd "Mounting hda1"
mount hda1
select "Which Kernel?" 50 Working New

goto {menu-option}
var image-name vmlinux.gz.working
goto 3f
@var image-name vmlinux.gz.working
goto 2f
@var image-name vmlinux.gz.new

@lcd "Loading Linux" {image-name}
load /{image-name}
lcd "Booting..."
execute root=/dev/sda5 ro console=ttyS0,115200
boot

```

See the documentation in /usr/share/doc/colo-VERSION for more details.

## Настройка последовательной консоли

Okay, the Linux installation as it stands now, would boot fine, but assumes the user will be logged in at a physical terminal. On Cobalt machines, this is particularly bad -- there's no such thing as a physical terminal.

**Заметка**

Those who do have the luxury of a supported video chipset may skip this section if they wish.

First, pull up an editor and hack away at /etc/inittab. Further down in the file, notice the following:

ФАЙЛ **`/etc/inittab`** **Выдержка из inittab**

```
# SERIAL CONSOLE
#c0:12345:respawn:/sbin/agetty 9600 ttyS0 vt102

# TERMINALS
c1:12345:respawn:/sbin/agetty 38400 tty1 linux
c2:12345:respawn:/sbin/agetty 38400 tty2 linux
c3:12345:respawn:/sbin/agetty 38400 tty3 linux
c4:12345:respawn:/sbin/agetty 38400 tty4 linux
c5:12345:respawn:/sbin/agetty 38400 tty5 linux
c6:12345:respawn:/sbin/agetty 38400 tty6 linux

# What to do at the "Three Finger Salute".
ca:12345:ctrlaltdel:/sbin/shutdown -r now

```

First, uncomment the c0 line. By default, it's set to use a terminal baud rate of 9600 bps. On Cobalt servers, this may be changed to 115200 to match the baud rate decided by the BOOT ROM. The following is how that section looks then. On a headless machine (e.g. Cobalt servers), it is also recommended to comment out the local terminal lines (c1 through to c6) as these have a habit of misbehaving when they can't open /dev/ttyX.

ФАЙЛ **`/etc/inittab`** **Выдержка из inittab**

```
# SERIAL CONSOLE
c0:12345:respawn:/sbin/agetty 115200 ttyS0 vt102

# TERMINALS -- These are useless on a headless qube
#c1:12345:respawn:/sbin/agetty 38400 tty1 linux
#c2:12345:respawn:/sbin/agetty 38400 tty2 linux
#c3:12345:respawn:/sbin/agetty 38400 tty3 linux
#c4:12345:respawn:/sbin/agetty 38400 tty4 linux
#c5:12345:respawn:/sbin/agetty 38400 tty5 linux
#c6:12345:respawn:/sbin/agetty 38400 tty6 linux

```

Now, lastly... tell the system, that the local serial port can be trusted as a secure terminal. The file to change at is /etc/securetty. It contains a list of terminals that the system trusts. Simply stick in two more lines, permitting the serial line to be used for root logins.

`root #` `echo 'ttyS0' >> /etc/securetty`

Lately, Linux also calls this /dev/tts/0 \-\- so add this too:

`root #` `echo 'tts/0' >> /etc/securetty`

## Tweaking the SGI PROM

### Setting generic PROM settings

With the bootloader installed, after rebooting (which will occur shortly), go to the System Maintenance Menu and select `Enter` on _Command Monitor (5)_ similar to the initial netbooting of the system.

КОД **Menu after boot**

```
1) Start System
2) Install System Software
3) Run Diagnostics
4) Recover System
5) Enter Command Monitor

```

Provide the location of the Volume Header:

`>>` `setenv SystemPartition scsi(0)disk(1)rdisk(0)partition(8)`

Автоматическая загрузка Gentoo:

`>>` `setenv AutoLoad Yes`

Установка часового пояса:

`>>` `setenv TimeZone EST5EDT`

Use the serial console - graphic adapter users should have "g" instead of "d1" (one):

`>>` `setenv console d1`

Set the serial console baud rate. This is optional, 9600 is the default setting, although one may use rates up to 38400 if that is desired:

`>>` `setenv dbaud 9600`

Now, the next settings depend on how the system is booted.

### Settings for direct volume-header booting

**Заметка**

This is covered here for completeness. It's recommended that users look into installing arcload instead.

**Заметка**

This only works on the Indy, Indigo2 (R4k) and Challenge S.

Set the root device to Gentoo's root partition, such as /dev/sda3:

`>>` `setenv OSLoadPartition <root device>`

To list the available kernels, type "ls".

`>>` `setenv OSLoader <kernel name>
`

`>>` `setenv OSLoadFilename <kernel name>`

Declare the kernel parameters to pass:

`>>` `setenv OSLoadOptions <kernel parameters>`

To try a kernel without messing with kernel parameters, use the boot -f PROM command:

`root #` `boot -f new root=/dev/sda5 ro`

### Настройки для arcload

arcload uses the OSLoadFilename option to specify which options to set from arc.cf. The configuration file is essentially a script, with the top-level blocks defining boot images for different systems, and inside that, optional settings. Thus, setting OSLoadFilename=mysys(serial) pulls in the settings for the mysys block, then sets further options overridden in serial.

In the example file above, one system block is defined, ip28 with working, new and debug options available. Next, define the PROM variables:

Select arcload as the bootloader:- sash64 or sashARCS:

`>>` `setenv OSLoader sash64`

Use the "working" kernel image, defined in "ip28" section of arc.cf:

`>>` `setenv OSLoadFilename ip28(working)`

Starting with arcload-0.5, files no longer need to be placed in the volume header -- they may be placed in a partition instead. To tell arcload where to look for its configuration file and kernels, one must set the OSLoadPartition PROM variable. The exact value here will depend on where the disk resides on the SCSI bus. Use the SystemPartition PROM variable as a guide -- only the partition number should need to change.

**Заметка**

Partitions are numbered starting at 0, not 1 as is the case in Linux.

To load from the volume header -- use partition 8:

`>>` `setenv OSLoadPartition scsi(0)disk(1)rdisk(0)partition(8)`

Otherwise, specify the partition and filesystem type:

`>>` `setenv OSLoadPartition scsi(0)disk(1)rdisk(0)partition(0)[ext2]`

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

Перезагрузившись в новое окружение Gentoo, переходите к [завершению установки Gentoo](/wiki/Handbook:MIPS/Installation/Finalizing/ru "Handbook:MIPS/Installation/Finalizing/ru").

[← Установка системных средств](/wiki/Handbook:MIPS/Installation/Tools/ru "Предыдущая часть") [К содержанию](/wiki/Handbook:MIPS/ru "Handbook:MIPS/ru") [Завершение →](/wiki/Handbook:MIPS/Installation/Finalizing/ru "Следующая часть")