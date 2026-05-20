# Bootloader

Other languages:

- [Deutsch](/wiki/Handbook:X86/Installation/Bootloader/de "Handbuch:X86/Installation/Bootloader (100% translated)")
- [English](/wiki/Handbook:X86/Installation/Bootloader "Handbook:X86/Installation/Bootloader (100% translated)")
- [español](/wiki/Handbook:X86/Installation/Bootloader/es "Manual de Gentoo: X86/Instalación/Arranque (100% translated)")
- [français](/wiki/Handbook:X86/Installation/Bootloader/fr "Manuel:X86/Installation/Système d'amorçage (100% translated)")
- [italiano](/wiki/Handbook:X86/Installation/Bootloader/it "Handbook:X86/Installation/Bootloader (100% translated)")
- [magyar](/wiki/Handbook:X86/Installation/Bootloader/hu "Handbook:X86/Installation/Bootloader/hu (100% translated)")
- [polski](/wiki/Handbook:X86/Installation/Bootloader/pl "Handbook:X86/Installation/Bootloader (100% translated)")
- [português do Brasil](/wiki/Handbook:X86/Installation/Bootloader/pt-br "Manual:X86/Instalação/Gerenciador de Boot (100% translated)")
- [čeština](/wiki/Handbook:X86/Installation/Bootloader/cs "Handbook:X86/Installation/Bootloader/cs (50% translated)")
- русский
- [தமிழ்](/wiki/Handbook:X86/Installation/Bootloader/ta "கையேடு:X86/நிறுவல்/துவக்கஏற்றி (100% translated)")
- [中文（中国大陆）‎](/wiki/Handbook:X86/Installation/Bootloader/zh-cn "手册：X86/安装/配置系统引导程序Bootloader (100% translated)")
- [日本語](/wiki/Handbook:X86/Installation/Bootloader/ja "ハンドブック:X86/インストール/ブートローダの設定 (100% translated)")
- [한국어](/wiki/Handbook:X86/Installation/Bootloader/ko "Handbook:X86/Installation/Bootloader/ko (100% translated)")

[X86 Handbook](/wiki/Handbook:X86/ru "Handbook:X86/ru")[Установка](/wiki/Handbook:X86/Full/Installation/ru "Handbook:X86/Full/Installation/ru")[Об установке](/wiki/Handbook:X86/Installation/About/ru "Handbook:X86/Installation/About/ru")[Выбор подходящего источника для установки](/wiki/Handbook:X86/Installation/Media/ru "Handbook:X86/Installation/Media/ru")[Настройка сети](/wiki/Handbook:X86/Installation/Networking/ru "Handbook:X86/Installation/Networking/ru")[Подготовка дисков](/wiki/Handbook:X86/Installation/Disks/ru "Handbook:X86/Installation/Disks/ru")[Установка файла stage](/wiki/Handbook:X86/Installation/Stage/ru "Handbook:X86/Installation/Stage/ru")[Установка базовой системы](/wiki/Handbook:X86/Installation/Base/ru "Handbook:X86/Installation/Base/ru")[Настройка ядра](/wiki/Handbook:X86/Installation/Kernel/ru "Handbook:X86/Installation/Kernel/ru")[Настройка системы](/wiki/Handbook:X86/Installation/System/ru "Handbook:X86/Installation/System/ru")[Установка системных утилит](/wiki/Handbook:X86/Installation/Tools/ru "Handbook:X86/Installation/Tools/ru")Настройка загрузчика[Завершение](/wiki/Handbook:X86/Installation/Finalizing/ru "Handbook:X86/Installation/Finalizing/ru")[Работа с Gentoo](/wiki/Handbook:X86/Full/Working/ru "Handbook:X86/Full/Working/ru")[Введение в Portage](/wiki/Handbook:X86/Working/Portage/ru "Handbook:X86/Working/Portage/ru")[USE-флаги](/wiki/Handbook:X86/Working/USE/ru "Handbook:X86/Working/USE/ru")[Возможности Portage](/wiki/Handbook:X86/Working/Features/ru "Handbook:X86/Working/Features/ru")[Система сценариев инициализации](/wiki/Handbook:X86/Working/Initscripts/ru "Handbook:X86/Working/Initscripts/ru")[Переменные окружения](/wiki/Handbook:X86/Working/EnvVar/ru "Handbook:X86/Working/EnvVar/ru")[Работа с Portage](/wiki/Handbook:X86/Full/Portage/ru "Handbook:X86/Full/Portage/ru")[Файлы и каталоги](/wiki/Handbook:X86/Portage/Files/ru "Handbook:X86/Portage/Files/ru")[Переменные](/wiki/Handbook:X86/Portage/Variables/ru "Handbook:X86/Portage/Variables/ru")[Смешение ветвей программного обеспечения](/wiki/Handbook:X86/Portage/Branches/ru "Handbook:X86/Portage/Branches/ru")[Дополнительные утилиты](/wiki/Handbook:X86/Portage/Tools/ru "Handbook:X86/Portage/Tools/ru")[Дополнительные репозитории пакетов](/wiki/Handbook:X86/Portage/CustomTree/ru "Handbook:X86/Portage/CustomTree/ru")[Расширенные возможности](/wiki/Handbook:X86/Portage/Advanced/ru "Handbook:X86/Portage/Advanced/ru")[Настройка сети OpenRC](/wiki/Handbook:X86/Full/Networking/ru "Handbook:X86/Full/Networking/ru")[Начальная настройка](/wiki/Handbook:X86/Networking/Introduction/ru "Handbook:X86/Networking/Introduction/ru")[Расширенная настройка](/wiki/Handbook:X86/Networking/Advanced/ru "Handbook:X86/Networking/Advanced/ru")[Модульное построение сети](/wiki/Handbook:X86/Networking/Modular/ru "Handbook:X86/Networking/Modular/ru")[Беспроводная сеть](/wiki/Handbook:X86/Networking/Wireless/ru "Handbook:X86/Networking/Wireless/ru")[Добавляем функциональность](/wiki/Handbook:X86/Networking/Extending/ru "Handbook:X86/Networking/Extending/ru")[Динамическое управление](/wiki/Handbook:X86/Networking/Dynamic/ru "Handbook:X86/Networking/Dynamic/ru")

Несмотря на то, что установка производится для 32-битного процессора, почти все **x86** материнские платы (начиная примерно с 2006–2007 годов и до настоящего времени), которые были выпущены с поддержкой UEFI, имеют _64–битную прошивку UEFI_. Некоторые пользователи могут заметить "64" в названиях параметров конфигурации и файлов в последующих разделах. Это вполне ожидаемо почти во всех случаях.

Есть несколько очень небольших исключений из этого стандарта 64-битной прошивки UEFI, а именно: несколько ранних моделей [Apple Mac](http://www.everymac.com/systems/apple/mac_pro/specs/mac-pro-quad-2.66-specs.html) и некоторые [планшеты Dell](http://www.dell.com/en-us/shop/productdetails/dell-venue-8-pro) на базе Intel Atom имели поддержку 32-битной прошивки UEFI. Подавляющее большинство читателей _никогда_ не столкнется с 32–битной прошивкой UEFI в природе. По этой причине 32–битная прошивка UEFI не рассматривается в **x86** Руководстве.

## Contents

- [1Выбор загрузчика](#.D0.92.D1.8B.D0.B1.D0.BE.D1.80_.D0.B7.D0.B0.D0.B3.D1.80.D1.83.D0.B7.D1.87.D0.B8.D0.BA.D0.B0)
- [2По умолчанию: GRUB](#.D0.9F.D0.BE_.D1.83.D0.BC.D0.BE.D0.BB.D1.87.D0.B0.D0.BD.D0.B8.D1.8E:_GRUB)
  - [2.1Emerge](#Emerge)
  - [2.2Установка](#.D0.A3.D1.81.D1.82.D0.B0.D0.BD.D0.BE.D0.B2.D0.BA.D0.B0)
    - [2.2.1DOS/Устаревшие BIOS системы](#DOS.2F.D0.A3.D1.81.D1.82.D0.B0.D1.80.D0.B5.D0.B2.D1.88.D0.B8.D0.B5_BIOS_.D1.81.D0.B8.D1.81.D1.82.D0.B5.D0.BC.D1.8B)
    - [2.2.2При использовании UEFI](#.D0.9F.D1.80.D0.B8_.D0.B8.D1.81.D0.BF.D0.BE.D0.BB.D1.8C.D0.B7.D0.BE.D0.B2.D0.B0.D0.BD.D0.B8.D0.B8_UEFI)
      - [2.2.2.1Optional: Secure Boot](#Optional:_Secure_Boot)
      - [2.2.2.2Отладка GRUB](#.D0.9E.D1.82.D0.BB.D0.B0.D0.B4.D0.BA.D0.B0_GRUB)
  - [2.3Настройка](#.D0.9D.D0.B0.D1.81.D1.82.D1.80.D0.BE.D0.B9.D0.BA.D0.B0)
- [3Alternative 1: systemd-boot](#Alternative_1:_systemd-boot)
  - [3.1Emerge](#Emerge_2)
  - [3.2Installation](#Installation)
  - [3.3Optional: Secure Boot](#Optional:_Secure_Boot_2)
- [4Альтернатива 2: EFI Stub](#.D0.90.D0.BB.D1.8C.D1.82.D0.B5.D1.80.D0.BD.D0.B0.D1.82.D0.B8.D0.B2.D0.B0_2:_EFI_Stub)
  - [4.1Unified Kernel Image](#Unified_Kernel_Image)
- [5Other Alternatives](#Other_Alternatives)
- [6Перезагрузка системы](#.D0.9F.D0.B5.D1.80.D0.B5.D0.B7.D0.B0.D0.B3.D1.80.D1.83.D0.B7.D0.BA.D0.B0_.D1.81.D0.B8.D1.81.D1.82.D0.B5.D0.BC.D1.8B)

## Выбор загрузчика

Когда ядро Linux настроено, системные утилиты установлены и конфигурационные файлы отредактированы, настало время для установки последней важной части системы Linux: загрузчика.

Загрузчик отвечает за запуск ядра Linux во время загрузки — без него система не будет знать, как действовать после нажатия кнопки питания.

Для **x86**, мы описали настройку либо [GRUB](#.D0.9F.D0.BE_.D1.83.D0.BC.D0.BE.D0.BB.D1.87.D0.B0.D0.BD.D0.B8.D1.8E:_GRUB), либо [LILO](#.D0.90.D0.BB.D1.8C.D1.82.D0.B5.D1.80.D0.BD.D0.B0.D1.82.D0.B8.D0.B2.D0.B0_1:_LILO) для систем на базе DOS/Legacy BIOS, и [GRUB](#.D0.9F.D0.BE_.D1.83.D0.BC.D0.BE.D0.BB.D1.87.D0.B0.D0.BD.D0.B8.D1.8E:_GRUB) или [efibootmgr](#.D0.90.D0.BB.D1.8C.D1.82.D0.B5.D1.80.D0.BD.D0.B0.D1.82.D0.B8.D0.B2.D0.B0_2:_efibootmgr) для UEFI-систем.

В этом разделе Руководства было сделано разделение между _установкой пакета_ (emerge) загрузчика и _установкой загрузчика на системный диск_. Термин _установка пакета_ (emerge) будет использоваться для установки пакета в систему с помощью пакетного менеджера [Portage](/wiki/Portage "Portage"). Термин _установка загрузчика на системный диск_ будет означать копирование файлов загрузчика или физическое изменение соответствующих разделов диска для того, чтобы _активировать и подготовить к работе_ загрузчик для следующей перезагрузки.

## По умолчанию: GRUB

По умолчанию, большинство систем Gentoo на данный момент используют [GRUB](/wiki/GRUB/ru "GRUB/ru") (из пакета [sys-boot/grub](https://packages.gentoo.org/packages/sys-boot/grub)), который является прямым продолжателем [GRUB Legacy](/wiki/GRUB_Legacy "GRUB Legacy"). Без дополнительных настроек GRUB поддерживает старые системы BIOS ("pc"). С небольшими настройкам, которые нужно выполнить до сборки, GRUB может поддерживать более чем полутора десятка дополнительных платформ. Для получения более подробной информации смотрите раздел [Предварительные требования](/wiki/GRUB/ru#.D0.9F.D1.80.D0.B5.D0.B4.D0.B2.D0.B0.D1.80.D0.B8.D1.82.D0.B5.D0.BB.D1.8C.D0.BD.D1.8B.D0.B5_.D1.82.D1.80.D0.B5.D0.B1.D0.BE.D0.B2.D0.B0.D0.BD.D0.B8.D1.8F "GRUB/ru") статьи [GRUB](/wiki/GRUB/ru "GRUB/ru").

### Emerge

Если используется старая материнская плата, BIOS которой поддерживает только таблицу разделов MBR, для установки GRUB не нужно никаких дополнительных настроек:

`root #` `emerge --ask --verbose sys-boot/grub`

Заметка для пользователей UEFI: запущенная команда выведет включенные значения в переменной GRUB\_PLATFORMS, перед компиляцией. Если используется более новая UEFI-совместимая материнская плата, пользователям сперва нужно убедиться, что `GRUB_PLATFORMS="efi-64"` включено (обычно это уже сделано по умолчанию). Если это не так, добавьте `GRUB_PLATFORMS="efi-64"` в файл /etc/portage/make.conf _до_ компиляции GRUB, что позволит собрать пакет с поддержкой EFI:

`root #` `echo 'GRUB_PLATFORMS="efi-64"' >> /etc/portage/make.conf`

`root #` `emerge --ask sys-boot/grub`

Если GRUB был каким-то образом был установлен до включения `GRUB_PLATFORMS="efi-64"`, то добавьте строку (из примера выше) в make.conf, после чего повторно переопределите зависимости для [для набора пакетов world](/wiki/World_set_(Portage) "World set (Portage)") с помощью emerge `--update --newuse`:

`root #` `emerge --ask --update --newuse --verbose sys-boot/grub`

GRUB теперь установлен в системе, но еще не активирован.

### Установка

Далее установим необходимые для GRUB файлы в каталог /boot/grub/ с помощью команды grub-install. Если предположить, что первый диском (тот, с которого будет загружаться система) является /dev/sda, то одна из следующих команд сделает это:

#### DOS/Устаревшие BIOS системы

При использовании BIOS:

`root #` `grub-install /dev/sda`

#### При использовании UEFI

**Важно**

Убедитесь, что системный раздел EFI был смонтирован _перед_ запуском grub-install. grub-install может установить файл GRUB EFI (grubx64.efi) в неправильном каталоге **без каких-либо** сообщений, о том, что использовался неправильный каталог.

For UEFI systems:

`root #` `grub-install --efi-directory=/efi`

```
Installing for x86_64-efi platform.
Installation finished. No error reported.

```

Upon successful installation, the output should match the output of the previous command. If the output does not match exactly, then proceed to [Debugging GRUB](/wiki/Handbook:X86/Blocks/Bootloader/ru#Debugging_GRUB "Handbook:X86/Blocks/Bootloader/ru"), otherwise jump to the [Configure step](/wiki/Handbook:X86/Blocks/Bootloader/ru#GRUB_Configure "Handbook:X86/Blocks/Bootloader/ru").

##### Optional: Secure Boot

To successfully boot with secure boot enabled the signing certificate must either be accepted by the [UEFI](/wiki/UEFI "UEFI") firmware, or [shim](/wiki/Shim "Shim") must be used as a pre-loader. Shim is pre-signed with the third-party Microsoft Certificate, accepted by default by most UEFI motherboards.

How to configure the UEFI firmware to accept custom keys depends on the firmware vendor, which is beyond the scope of the handbook. Below is shown how to setup shim instead. Here it is assumed that the user has already followed the instructions in the previous sections to generate a signing key and to configure portage to use it. If this is not the case please return first to the [Kernel installation](/wiki/Handbook:X86/Installation/Kernel/ru "Handbook:X86/Installation/Kernel/ru") section.

The package [sys-boot/grub](https://packages.gentoo.org/packages/sys-boot/grub) installs a prebuilt and signed stand-alone EFI executable if the [secureboot](https://packages.gentoo.org/useflags/secureboot) [USE flag/ru](/wiki/USE_flag/ru "USE flag/ru") USE flag is enabled. Install the required packages and copy the stand-alone grub, Shim, and the MokManager to the same directory on the EFI System Partition. For example:

`root #` `emerge sys-boot/grub sys-boot/shim sys-boot/mokutil sys-boot/efibootmgr`

`root #` `cp /usr/share/shim/BOOTX64.EFI /efi/EFI/Gentoo/shimx64.efi
`

`root #` `cp /usr/share/shim/mmx64.efi /efi/EFI/Gentoo/mmx64.efi
`

`root #` `cp /usr/lib/grub/grub-x86_64.efi.signed /efi/EFI/Gentoo/grubx64.efi
`

Next register the signing key in shims MOKlist, this requires keys in the DER format, whereas sbsign and the kernel build system expect keys in the PEM format. In the previous sections of the handbook an example was shown to generate such a signing PEM key, this key must now be converted to the DER format:

`root #` `openssl x509 -in /path/to/kernel_key.pem -inform PEM -out /path/to/kernel_key.der -outform DER`

**Заметка**

The path used here must be the path to the pem file containing the certificate belonging to the generated key. In this example both key and certificate are in the same pem file.

Then the converted certificate can be imported into Shims MOKlist, this command will ask to set some password for the import request:

`root #` `mokutil --import /path/to/kernel_key.der`

**Заметка**

When the currently booted kernel already trusts the certificate being imported, the message "Already in kernel trusted keyring." will be returned here. If this happens, re-run the above command with the argument --ignore-keyring added.

Next, register Shim with the UEFI firmware. In the following command, `boot-disk` and `boot-partition-id` must be replaced with the disk and partition identifier of the EFI system partition:

`root #` `efibootmgr --create --disk /dev/boot-disk --part boot-partition-id --loader '\EFI\Gentoo\shimx64.efi' --label 'GRUB via Shim' --unicode`

Note that this prebuilt and signed stand-alone version of grub reads the grub.cfg from a different location then usual. Instead of the default /boot/grub/grub.cfg the config file should be in the same directory that the grub EFI executable is in, e.g. /efi/EFI/Gentoo/grub.cfg. When [sys-kernel/installkernel](https://packages.gentoo.org/packages/sys-kernel/installkernel) is used to install the kernel and update the grub configuration then the GRUB\_CFG environment variable may be used to override the usual location of the grub config file.

For example:

`root #` `grub-mkconfig -o /efi/EFI/Gentoo/grub.cfg`

Or, via [installkernel](/wiki/Installkernel "Installkernel"):

ФАЙЛ **`/etc/env.d/99grub`**

```
GRUB_CFG=/efi/EFI/Gentoo/grub.cfg

```

`root #` `env-update`

**Заметка**

The import process will not be completed until the system is rebooted. After completing all steps in the handbook, restart the system and Shim will load, it will find the import request registered by mokutil. The MokManager application will start and ask for the password that was set when creating the import request. Follow the on-screen instructions to complete the import of the certificate, then reboot the system into the UEFI menu and enable the Secure Boot setting.

##### Отладка GRUB

When debugging GRUB, there are a couple of quick fixes that may result in a bootable installation without having to reboot to a new live image environment.

In the event that "EFI variables are not supported on this system" is displayed somewhere in the output, it is likely the live image was not booted in EFI mode and is presently in Legacy BIOS boot mode. The solution is to try the [removable GRUB step](/wiki/Handbook:X86/Blocks/Bootloader/ru#GRUB_Install_EFI_systems_removable "Handbook:X86/Blocks/Bootloader/ru") mentioned below. This will overwrite the executable EFI file located at /EFI/BOOT/BOOTX64.EFI. Upon rebooting in EFI mode, the motherboard firmware may execute this default boot entry and execute GRUB.

**Важно**

Если grub-install возвращает ошибку наподобии `Could not prepare Boot variable: Read-only file system`, необходимо перемонтировать специальную точку монтирования efivars в режим чтения-записи:

`root #` `mount -o remount,rw,nosuid,nodev,noexec --types efivarfs efivarfs /sys/firmware/efi/efivars`

`root #` `mount -o remount,rw,nosuid,nodev,noexec --types efivarfs efivarfs /sys/firmware/efi/efivars`

This is caused by certain non-official Gentoo environments not mounting the special EFI filesystem by default. If the previous command does not run, then reboot using an official Gentoo live image environment in EFI mode.

Некоторые производители материнских плат поддерживают только каталог /efi/boot/ для расположения файла .EFI в системном разделе EFI (ESP). Установщик GRUB может выполнить эту операцию автоматически, если использовать параметр `--removable`. Убедитесь, что ESP смонтирован до запуска следующих команд. Предполагая, что ESP смонтирован в /boot (как было предложено ранее), выполните:

`root #` `grub-install --target=x86_64-efi --efi-directory=/efi --removable`

Это создает каталог по умолчанию, определенный спецификацией UEFI, а затем скопирует файл grubx64.efi в каталог EFI «по умолчанию», определенный той же спецификацией.

### Настройка

Далее, нужно сгенерировать конфигурационный файл GRUB на основе настроек пользователя, указанных в файле /etc/default/grub и сценариях /etc/grub.d. В большинстве случаев ничего не нужно настраивать, так как GRUB автоматически определяет, какое ядро есть для загрузки (самый высокий приоритет у /boot/) и какая файловая система у rootfs. Здесь также можно добавить параметры ядра в /etc/default/grub, используя переменную GRUB\_CMDLINE\_LINUX.

Для создания окончательной конфигурации GRUB, запустите команду grub-mkconfig:

`root #` `grub-mkconfig -o /boot/grub/grub.cfg`

```
Generating grub.cfg ...
Found linux image: /boot/vmlinuz-6.18.8-gentoo
Found initrd image: /boot/initramfs-genkernel-x86-6.18.8-gentoo
done

```

Вывод команды должен содержать по крайней мере один образ Linux, так как он необходим для загрузки системы. Если используется initramfs или ядро создавалось с помощью genkernel, также должен быть указан правильный образ initrd. Если это не так, перейдите в /boot/ и проверьте содержимое, используя команду ls. Если файлы действительно отсутствуют, вернитесь к инструкции по настройке и установке ядра.

**Совет**

Для обнаружения других операционных систем на других диска можно использовать утилиту os-prober. Она может обнаруживать Windows 7, 8.1, 10, и другие дистрибутивы Linux. Для таких систем с двойной загрузкой необходимо установить пакет [sys-boot/os-prober](https://packages.gentoo.org/packages/sys-boot/os-prober) и затем перезапустить команду grub-mkconfig (как было показано выше). Если появятся проблемы, полностью перечитайте статью [GRUB](/wiki/GRUB/ru "GRUB/ru") до того, как спрашивать у сообщества Gentoo поддержки.

## Alternative 1: systemd-boot

Another option is [systemd-boot](/wiki/Systemd/systemd-boot "Systemd/systemd-boot"), which works on both OpenRC and systemd machines. It is a thin chainloader and works well with secure boot.

### Emerge

To install systemd-boot, enable the [boot](https://packages.gentoo.org/useflags/boot) [USE flag/ru](/wiki/USE_flag/ru "USE flag/ru") USE flag and re-install [sys-apps/systemd](https://packages.gentoo.org/packages/sys-apps/systemd) (for systemd systems) or [sys-apps/systemd-utils](https://packages.gentoo.org/packages/sys-apps/systemd-utils) (for OpenRC systems):

ФАЙЛ **`/etc/portage/package.use/systemd-boot`**

```
sys-apps/systemd boot
sys-apps/systemd-utils boot

```

`root #` `emerge --ask sys-apps/systemd`

Or

`root #` `emerge --ask sys-apps/systemd-utils`

### Installation

Now, install the systemd-boot loader to the [EFI System Partition](/wiki/EFI_System_Partition/ru "EFI System Partition/ru"):

`root #` `bootctl install`

**Важно**

Make sure the EFI system partition has been mounted _before_ running bootctl install.

When using this bootloader, before rebooting, verify that a new bootable entry exists using:

`root #` `bootctl list`

**Предупреждение**

The kernel command line for new systemd-boot entries is read from /etc/kernel/cmdline or /usr/lib/kernel/cmdline. If neither file is present, then the kernel command line of the currently booted kernel is re-used (/proc/cmdline). On new installs it might therefore happen that the kernel command line of the live CD is accidentally used to boot the new kernel. The kernel command line for registered entries can be checked with:

`root #` `bootctl list`

If this does not show the desired kernel command line then create /etc/kernel/cmdline containing the correct kernel command line and re-install the kernel.

If no new entry exists, ensure the [sys-kernel/installkernel](https://packages.gentoo.org/packages/sys-kernel/installkernel) package has been installed with the [systemd](https://packages.gentoo.org/useflags/systemd) [USE flag/ru](/wiki/USE_flag/ru "USE flag/ru") and [systemd-boot](https://packages.gentoo.org/useflags/systemd-boot) [USE flag/ru](/wiki/USE_flag/ru "USE flag/ru") USE flags enabled, and re-run the kernel installation.

For the distribution kernels:

`root #` `emerge --ask --config sys-kernel/gentoo-kernel`

For a manually configured and compiled kernel:

`root #` `make install`

**Важно**

When installing kernels for systemd-boot, no root= kernel command line argument is added by default. On systemd systems that are using an initramfs users may rely instead on [systemd-gpt-auto-generator](/wiki/Systemd#Automatic_mounting_of_partitions_at_boot.2Fru "Systemd") to automatically find the root partition at boot. Otherwise users should manually specify the location of the root partition by setting root= in /etc/kernel/cmdline as well as any other kernel command line arguments that should be used. And then reinstalling the kernel as described above.

### Optional: Secure Boot

When the [secureboot](https://packages.gentoo.org/useflags/secureboot) [USE flag/ru](/wiki/USE_flag/ru "USE flag/ru") USE flag is enabled, the systemd-boot EFI executable will be signed by portage automatically. Furthermore, bootctl install will automatically install the signed version.

To successfully boot with secure boot enabled the used certificate must either be accepted by the [UEFI](/wiki/UEFI "UEFI") firmware, or [shim](/wiki/Shim "Shim") must be used as a pre-loader. Shim is pre-signed with the third-party Microsoft Certificate, accepted by default by most UEFI motherboards.

How to configure the UEFI firmware to accept custom keys depends on the firmware vendor, which is beyond the scope of the handbook. Below is shown how to setup shim instead. Here it is assumed that the user has already followed the instructions in the previous sections to generate a signing key and to configure portage to use it. If this is not the case please return first to the [Kernel installation](/wiki/Handbook:X86/Installation/Kernel/ru "Handbook:X86/Installation/Kernel/ru") section.

`root #` `emerge --ask sys-boot/shim sys-boot/mokutil sys-boot/efibootmgr`

`root #` `bootctl install --no-variables`

`root #` `cp /usr/share/shim/BOOTX64.EFI /efi/EFI/systemd/shimx64.efi`

`root #` `cp /usr/share/shim/mmx64.efi /efi/EFI/systemd/mmx64.efi`

Shims MOKlist requires keys in the DER format, whereas sbsign and the kernel build system expect keys in the PEM format. In the previous sections of the handbook an example was shown to generate such a signing PEM key, this key must now be converted to the DER format:

`root #` `openssl x509 -in /path/to/kernel_key.pem -inform PEM -out /path/to/kernel_key.der -outform DER`

**Заметка**

The path used here must be the path to the pem file containing the certificate belonging to the generated key. In this example both key and certificate are in the same pem file.

Then the converted certificate can be imported into Shims MOKlist:

`root #` `mokutil --import /path/to/kernel_key.der`

**Заметка**

When the currently booted kernel already trusts the certificate being imported, the message "Already in kernel trusted keyring." will be returned here. If this happens, re-run the above command with the argument --ignore-keyring added.

And finally we register Shim with the UEFI firmware. In the following command, `boot-disk` and `boot-partition-id` must be replaced with the disk and partition identifier of the EFI system partition:

`root #` `efibootmgr --create --disk /dev/boot-disk --part boot-partition-id --loader '\EFI\systemd\shimx64.efi'  --label 'Systemd-boot via Shim' --unicode '\EFI\systemd\systemd-bootx64.efi'`

**Заметка**

The import process will not be completed until the system is rebooted. After completing all steps in the handbook, restart the system and Shim will load, it will find the import request registered by mokutil. The MokManager application will start and ask for the password that was set when creating the import request. Follow the on-screen instructions to complete the import of the certificate, then reboot the system into the UEFI menu and enable the Secure Boot setting.

## Альтернатива 2: EFI Stub

На UEFI-системах, с прошивкой UEFI (другими словами, основным загрузчиком), можно напрямую управлять загрузочными записями UEFI. Таким системам не требуется дополнительный (также известный как вторичный) загрузчик, такой как GRUB, который помогает загрузить систему. Учитывая сказанное, использование дополнительного EFI-загрузчика типа GRUB имеет смысл лишь в том, чтобы _расширить_ функциональность UEFI во время загрузки. Использование efibootmgr подойдёт больше для тех, кто хочет получить больше минимализма (хотя это сложнее) при загрузке системы; использование GRUB проще для большинства пользователей, так как он предлагает более гибкий подход для загрузки UEFI-систем.

System administrators who desire to take a minimalist, although more rigid, approach to booting the system can avoid secondary bootloaders and boot the Linux kernel as an [EFI stub](/wiki/EFI_stub "EFI stub").

Помните, [sys-boot/efibootmgr](https://packages.gentoo.org/packages/sys-boot/efibootmgr) — это не загрузчик; это средство для взаимодействия с прошивкой UEFI и обновления её настроек, для того, чтобы можно было загрузить установленное ядро Linux с дополнительными параметрами (если необходимо) или организовать несколько загрузочных записей. Это взаимодействие осуществляется через переменные EFI (что требует поддержки переменных EFI со стороны ядра).

Перед тем, как продолжить, _обязательно_ прочитайте всю статью [EFI stub](/wiki/EFI_stub "EFI stub"). В ядре должны быть включены определенные параметр, чтобы ядро могло загрузится напрямую с системной прошивкой UEFI. Это может потребовать пересборки ядра. Также взгляните на статью [efibootmgr](/wiki/Efibootmgr/ru "Efibootmgr/ru").

It is also a good idea to take a look at the [efibootmgr](/wiki/Efibootmgr "Efibootmgr") article for additional information.

**Заметка**

Повторим еще раз, efibootmgr _не является_ обязательным требованием для загрузки из UEFI. Ядро Linux может загружаться сразу же, а дополнительные параметры ядра могут быть встроены в само ядра (в конфигурации ядра есть параметр CONFIG\_CMDLINE, который позволяет пользователю определить параметры загрузки). Даже initramfs может быть встроен в ядро.

Установите efibootmgr:

ФАЙЛ **`/etc/portage/package.accept_keywords/installkernel`**

```
sys-kernel/installkernel
sys-boot/uefi-mkconfig
app-emulation/virt-firmware

```

ФАЙЛ **`/etc/portage/package.use/installkernel`**

```
sys-kernel/installkernel efistub

```

Then reinstall [installkernel](/wiki/Installkernel "Installkernel"), create the /efi directory and reinstall the kernel:

`root #` `emerge --ask sys-kernel/installkernel`

`root #` `mkdir -p /efi`

For distribution kernels:

`root #` `emerge --ask --config sys-kernel/gentoo-kernel{,-bin}`

For manually managed kernels:

`root #` `make install`

Создайте каталог /efi и затем скопируйте в него ядро, назвав его bootx64.efi:

`root #` `mkdir -p /efi
`

`root #` `cp /boot/vmlinuz-* /efi/bootx64.efi
`

Install the efibootmgr package:

`root #` `emerge --ask sys-boot/efibootmgr`

Далее, сообщите прошивке UEFI создать загрузочную запись и называть её "Gentoo", в которой будет свежее ядро с EFI stub:

`root #` `efibootmgr --create --disk /dev/sda --part 1 --label "gentoo" --loader "\bootx64.efi"`

**Заметка**

Символ обратной косой черты ( `\`), применяемый для разделения каталогов, является обязательным при использовании определений UEFI.

Если используется файловая система инициализации, размещаемая в оперативной памяти (initramfs), добавьте соответствующий загрузочный параметр:

`root #` `efibootmgr --create --disk /dev/sda --part 1 --label "gentoo" --loader "\bootx64.efi" --unicode "initrd=\initramfs.img"`

**Совет**

Additional kernel command line options may be parsed by the firmware to the kernel by specifying them along with the initrd=... option as shown above.

После внесения изменений и перезагрузки системы появится загрузочная запись под именем «gentoo».

### Unified Kernel Image

If [installkernel](/wiki/Installkernel "Installkernel") was configured to build and install unified kernel images. The unified kernel image should already be installed to the EFI/Linux directory on the EFI system partition, if this is not the case ensure the directory exists and then run the kernel installation again as described earlier in the handbook.

To add a direct boot entry for the installed unified kernel image:

`root #` `efibootmgr --create --disk /dev/sda --part 1 --label "gentoo" --loader "\EFI\Linux\gentoo-x.y.z.efi"`

## Other Alternatives

For other options that are not covered in the Handbook, see the full list of available [bootloaders](/wiki/Bootloader/ru "Bootloader/ru").

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

Перезагрузившись в новое окружение Gentoo, переходите к [завершению установки Gentoo](/wiki/Handbook:X86/Installation/Finalizing/ru "Handbook:X86/Installation/Finalizing/ru").

[← Установка системных средств](/wiki/Handbook:X86/Installation/Tools/ru "Предыдущая часть") [К содержанию](/wiki/Handbook:X86/ru "Handbook:X86/ru") [Завершение →](/wiki/Handbook:X86/Installation/Finalizing/ru "Следующая часть")