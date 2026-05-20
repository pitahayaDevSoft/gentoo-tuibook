# Kernel

Other languages:

- [Deutsch](/wiki/Handbook:HPPA/Installation/Kernel/de "Handbuch:HPPA/Installation/Kernel (100% translated)")
- [English](/wiki/Handbook:HPPA/Installation/Kernel "Handbook:HPPA/Installation/Kernel (100% translated)")
- [español](/wiki/Handbook:HPPA/Installation/Kernel/es "Manual de Gentoo: HPPA/Instalación/Núcleo (100% translated)")
- [français](/wiki/Handbook:HPPA/Installation/Kernel/fr "Handbook:HPPA/Installation/Kernel/fr (100% translated)")
- [italiano](/wiki/Handbook:HPPA/Installation/Kernel/it "Handbook:HPPA/Installation/Kernel/it (100% translated)")
- [magyar](/wiki/Handbook:HPPA/Installation/Kernel/hu "Handbook:HPPA/Installation/Kernel/hu (100% translated)")
- [polski](/wiki/Handbook:HPPA/Installation/Kernel/pl "Handbook:HPPA/Installation/Kernel (100% translated)")
- [português do Brasil](/wiki/Handbook:HPPA/Installation/Kernel/pt-br "Manual:HPPA/Instalação/Kernel (100% translated)")
- русский
- [தமிழ்](/wiki/Handbook:HPPA/Installation/Kernel/ta "கையேடு:HPPA/நிறுவல்/கர்னல் (100% translated)")
- [中文（中国大陆）‎](/wiki/Handbook:HPPA/Installation/Kernel/zh-cn "手册：HPPA/安装/配置Linux内核 (100% translated)")
- [日本語](/wiki/Handbook:HPPA/Installation/Kernel/ja "ハンドブック:HPPA/インストール/カーネル (100% translated)")
- [한국어](/wiki/Handbook:HPPA/Installation/Kernel/ko "Handbook:HPPA/Installation/Kernel/ko (100% translated)")

[HPPA Handbook](/wiki/Handbook:HPPA/ru "Handbook:HPPA/ru")[Установка](/wiki/Handbook:HPPA/Full/Installation/ru "Handbook:HPPA/Full/Installation/ru")[Об установке](/wiki/Handbook:HPPA/Installation/About/ru "Handbook:HPPA/Installation/About/ru")[Выбор подходящего источника для установки](/wiki/Handbook:HPPA/Installation/Media/ru "Handbook:HPPA/Installation/Media/ru")[Настройка сети](/wiki/Handbook:HPPA/Installation/Networking/ru "Handbook:HPPA/Installation/Networking/ru")[Подготовка дисков](/wiki/Handbook:HPPA/Installation/Disks/ru "Handbook:HPPA/Installation/Disks/ru")[Установка файла stage](/wiki/Handbook:HPPA/Installation/Stage/ru "Handbook:HPPA/Installation/Stage/ru")[Установка базовой системы](/wiki/Handbook:HPPA/Installation/Base/ru "Handbook:HPPA/Installation/Base/ru")Настройка ядра[Настройка системы](/wiki/Handbook:HPPA/Installation/System/ru "Handbook:HPPA/Installation/System/ru")[Установка системных утилит](/wiki/Handbook:HPPA/Installation/Tools/ru "Handbook:HPPA/Installation/Tools/ru")[Настройка загрузчика](/wiki/Handbook:HPPA/Installation/Bootloader/ru "Handbook:HPPA/Installation/Bootloader/ru")[Завершение](/wiki/Handbook:HPPA/Installation/Finalizing/ru "Handbook:HPPA/Installation/Finalizing/ru")[Работа с Gentoo](/wiki/Handbook:HPPA/Full/Working/ru "Handbook:HPPA/Full/Working/ru")[Введение в Portage](/wiki/Handbook:HPPA/Working/Portage/ru "Handbook:HPPA/Working/Portage/ru")[USE-флаги](/wiki/Handbook:HPPA/Working/USE/ru "Handbook:HPPA/Working/USE/ru")[Возможности Portage](/wiki/Handbook:HPPA/Working/Features/ru "Handbook:HPPA/Working/Features/ru")[Система сценариев инициализации](/wiki/Handbook:HPPA/Working/Initscripts/ru "Handbook:HPPA/Working/Initscripts/ru")[Переменные окружения](/wiki/Handbook:HPPA/Working/EnvVar/ru "Handbook:HPPA/Working/EnvVar/ru")[Работа с Portage](/wiki/Handbook:HPPA/Full/Portage/ru "Handbook:HPPA/Full/Portage/ru")[Файлы и каталоги](/wiki/Handbook:HPPA/Portage/Files/ru "Handbook:HPPA/Portage/Files/ru")[Переменные](/wiki/Handbook:HPPA/Portage/Variables/ru "Handbook:HPPA/Portage/Variables/ru")[Смешение ветвей программного обеспечения](/wiki/Handbook:HPPA/Portage/Branches/ru "Handbook:HPPA/Portage/Branches/ru")[Дополнительные утилиты](/wiki/Handbook:HPPA/Portage/Tools/ru "Handbook:HPPA/Portage/Tools/ru")[Дополнительные репозитории пакетов](/wiki/Handbook:HPPA/Portage/CustomTree/ru "Handbook:HPPA/Portage/CustomTree/ru")[Расширенные возможности](/wiki/Handbook:HPPA/Portage/Advanced/ru "Handbook:HPPA/Portage/Advanced/ru")[Настройка сети OpenRC](/wiki/Handbook:HPPA/Full/Networking/ru "Handbook:HPPA/Full/Networking/ru")[Начальная настройка](/wiki/Handbook:HPPA/Networking/Introduction/ru "Handbook:HPPA/Networking/Introduction/ru")[Расширенная настройка](/wiki/Handbook:HPPA/Networking/Advanced/ru "Handbook:HPPA/Networking/Advanced/ru")[Модульное построение сети](/wiki/Handbook:HPPA/Networking/Modular/ru "Handbook:HPPA/Networking/Modular/ru")[Беспроводная сеть](/wiki/Handbook:HPPA/Networking/Wireless/ru "Handbook:HPPA/Networking/Wireless/ru")[Добавляем функциональность](/wiki/Handbook:HPPA/Networking/Extending/ru "Handbook:HPPA/Networking/Extending/ru")[Динамическое управление](/wiki/Handbook:HPPA/Networking/Dynamic/ru "Handbook:HPPA/Networking/Dynamic/ru")

## Contents

- [1Необязательно: Установка файлов прошивки и/или микрокода](#.D0.9D.D0.B5.D0.BE.D0.B1.D1.8F.D0.B7.D0.B0.D1.82.D0.B5.D0.BB.D1.8C.D0.BD.D0.BE:_.D0.A3.D1.81.D1.82.D0.B0.D0.BD.D0.BE.D0.B2.D0.BA.D0.B0_.D1.84.D0.B0.D0.B9.D0.BB.D0.BE.D0.B2_.D0.BF.D1.80.D0.BE.D1.88.D0.B8.D0.B2.D0.BA.D0.B8_.D0.B8.2F.D0.B8.D0.BB.D0.B8_.D0.BC.D0.B8.D0.BA.D1.80.D0.BE.D0.BA.D0.BE.D0.B4.D0.B0)
  - [1.1Файлы прошивки](#.D0.A4.D0.B0.D0.B9.D0.BB.D1.8B_.D0.BF.D1.80.D0.BE.D1.88.D0.B8.D0.B2.D0.BA.D0.B8)
    - [1.1.1Рекомендуется: Linux Firmware](#.D0.A0.D0.B5.D0.BA.D0.BE.D0.BC.D0.B5.D0.BD.D0.B4.D1.83.D0.B5.D1.82.D1.81.D1.8F:_Linux_Firmware)
      - [1.1.1.1Загрузка прошивки](#.D0.97.D0.B0.D0.B3.D1.80.D1.83.D0.B7.D0.BA.D0.B0_.D0.BF.D1.80.D0.BE.D1.88.D0.B8.D0.B2.D0.BA.D0.B8)
- [2sys-kernel/installkernel](#sys-kernel.2Finstallkernel)
  - [2.1Начальный загрузчик](#.D0.9D.D0.B0.D1.87.D0.B0.D0.BB.D1.8C.D0.BD.D1.8B.D0.B9_.D0.B7.D0.B0.D0.B3.D1.80.D1.83.D0.B7.D1.87.D0.B8.D0.BA)
    - [2.1.1GRUB](#GRUB)
    - [2.1.2Traditional layout, other bootloaders (e.g. (e)lilo, syslinux, etc.)](#Traditional_layout.2C_other_bootloaders_.28e.g._.28e.29lilo.2C_syslinux.2C_etc..29)
  - [2.2Initramfs](#Initramfs)
    - [2.2.1Chroot detection](#Chroot_detection)
- [3Конфигурация и компиляция ядра](#.D0.9A.D0.BE.D0.BD.D1.84.D0.B8.D0.B3.D1.83.D1.80.D0.B0.D1.86.D0.B8.D1.8F_.D0.B8_.D0.BA.D0.BE.D0.BC.D0.BF.D0.B8.D0.BB.D1.8F.D1.86.D0.B8.D1.8F_.D1.8F.D0.B4.D1.80.D0.B0)
  - [3.1Distribution-ядра](#Distribution-.D1.8F.D0.B4.D1.80.D0.B0)
    - [3.1.1Optional: Signed kernel modules](#Optional:_Signed_kernel_modules)
    - [3.1.2Установка distribution-ядра](#.D0.A3.D1.81.D1.82.D0.B0.D0.BD.D0.BE.D0.B2.D0.BA.D0.B0_distribution-.D1.8F.D0.B4.D1.80.D0.B0)
    - [3.1.3Обновление и очистка](#.D0.9E.D0.B1.D0.BD.D0.BE.D0.B2.D0.BB.D0.B5.D0.BD.D0.B8.D0.B5_.D0.B8_.D0.BE.D1.87.D0.B8.D1.81.D1.82.D0.BA.D0.B0)
    - [3.1.4Задачи после установки/обновления](#.D0.97.D0.B0.D0.B4.D0.B0.D1.87.D0.B8_.D0.BF.D0.BE.D1.81.D0.BB.D0.B5_.D1.83.D1.81.D1.82.D0.B0.D0.BD.D0.BE.D0.B2.D0.BA.D0.B8.2F.D0.BE.D0.B1.D0.BD.D0.BE.D0.B2.D0.BB.D0.B5.D0.BD.D0.B8.D1.8F)
      - [3.1.4.1Ручная пересборка initramfs или Unified Kernel Image](#.D0.A0.D1.83.D1.87.D0.BD.D0.B0.D1.8F_.D0.BF.D0.B5.D1.80.D0.B5.D1.81.D0.B1.D0.BE.D1.80.D0.BA.D0.B0_initramfs_.D0.B8.D0.BB.D0.B8_Unified_Kernel_Image)
  - [3.2Альтернатива: Ручная настройка](#.D0.90.D0.BB.D1.8C.D1.82.D0.B5.D1.80.D0.BD.D0.B0.D1.82.D0.B8.D0.B2.D0.B0:_.D0.A0.D1.83.D1.87.D0.BD.D0.B0.D1.8F_.D0.BD.D0.B0.D1.81.D1.82.D1.80.D0.BE.D0.B9.D0.BA.D0.B0)
  - [3.3Установка исходного кода ядра](#.D0.A3.D1.81.D1.82.D0.B0.D0.BD.D0.BE.D0.B2.D0.BA.D0.B0_.D0.B8.D1.81.D1.85.D0.BE.D0.B4.D0.BD.D0.BE.D0.B3.D0.BE_.D0.BA.D0.BE.D0.B4.D0.B0_.D1.8F.D0.B4.D1.80.D0.B0)
    - [3.3.1Процесс modprobed-db](#.D0.9F.D1.80.D0.BE.D1.86.D0.B5.D1.81.D1.81_modprobed-db)
      - [3.3.1.1Option 2 - Assisted manual process](#Option_2_-_Assisted_manual_process)
      - [3.3.1.2Option 3 - Configuring by hand](#Option_3_-_Configuring_by_hand)
    - [3.3.2Optional: Signed kernel modules](#Optional:_Signed_kernel_modules_2)
  - [3.4Компиляция и установка](#.D0.9A.D0.BE.D0.BC.D0.BF.D0.B8.D0.BB.D1.8F.D1.86.D0.B8.D1.8F_.D0.B8_.D1.83.D1.81.D1.82.D0.B0.D0.BD.D0.BE.D0.B2.D0.BA.D0.B0)

## Необязательно: Установка файлов прошивки и/или микрокода

### Файлы прошивки

#### Рекомендуется: Linux Firmware

On many systems, non-FOSS firmware is required for certain hardware to function. The [sys-kernel/linux-firmware](https://packages.gentoo.org/packages/sys-kernel/linux-firmware) package contains firmware for many, but not all, devices.

**Совет**

Most wireless cards and GPUs require firmware to function.

`root #` `emerge --ask sys-kernel/linux-firmware`

**Заметка**

Установка определённых пакетов прошивок часто требует принятия соответствующих лицензий на прошивку. При необходимости посетите раздел руководства [о принятии лицензии](/wiki/Handbook:HPPA/Working/Portage/ru#Licenses "Handbook:HPPA/Working/Portage/ru") для получения помощи.

##### Загрузка прошивки

Firmware files are typically loaded when the associated kernel module is loaded. This means the firmware must be built into the kernel using **CONFIG\_EXTRA\_FIRMWARE** if the kernel module is set to _Y_ instead of _M_. In most cases, building-in a module which required firmware can complicate or break loading.

## sys-kernel/installkernel

[Installkernel](/wiki/Installkernel "Installkernel") may be used to automate the kernel installation, [initramfs](/wiki/Initramfs "Initramfs") generation, [unified kernel image](/wiki/Unified_kernel_image "Unified kernel image") generation and bootloader configuration, among other things. [sys-kernel/installkernel](https://packages.gentoo.org/packages/sys-kernel/installkernel) implements two paths of achieving this: the traditional installkernel originating from Debian and [systemd](/wiki/Systemd/ru "Systemd/ru")'s kernel-install. Which one to choose depends, among other things, on the system's bootloader. By default, systemd's kernel-install is used on systemd profiles, while the traditional installkernel is the default for other profiles.

### Начальный загрузчик

Now is the time to think about which bootloader the user wants for the system.

**Важно**

Only one selection is required in the following subsection, if unsure of which to use then go with the first listed for now. It's always possible to switch at a later date if required.

#### GRUB

Users of GRUB can use either systemd's kernel-install or the traditional Debian installkernel. The [systemd](https://packages.gentoo.org/useflags/systemd) [USE flag/ru](/wiki/USE_flag/ru "USE flag/ru") USE flag switches between these implementations. To automatically run grub-mkconfig when installing the kernel, enable the [grub](https://packages.gentoo.org/useflags/grub) [USE flag/ru](/wiki/USE_flag/ru "USE flag/ru") [USE flag](/wiki/USE_flag/ru "USE flag/ru").

**Заметка**

GRUB requires kernels to be installed to /boot.

ФАЙЛ **`/etc/portage/package.use/installkernel`**

```
sys-kernel/installkernel grub

```

`root #` `emerge --ask sys-kernel/installkernel`

`root #` `mkdir -p /etc/cmdline.d`

`root #` `ln -s /etc/kernel/cmdline /etc/cmdline.d/00-installkernel.conf`

`root #` `emerge --ask sys-kernel/installkernel`

}}

**Заметка**

systemd-boot requires kernels to be installed to /efi.

**Заметка**

When [app-emulation/virt-firmware](https://packages.gentoo.org/packages/app-emulation/virt-firmware) is used to configure the UEFI ensure that the kernel-bootcfg-boot-successful service is enabled before attempting to install the kernel. This implementation of EFIstub booting is the default for systemd systems.

`root #` `systemctl enable kernel-bootcfg-boot-successful.service`

**Заметка**

EFIstub requires kernels to be installed to /efi.

#### Traditional layout, other bootloaders (e.g. (e)lilo, syslinux, etc.)

The traditional /boot layout (for e.g. (e)LILO, syslinux, etc.) is used by default if the [grub](https://packages.gentoo.org/useflags/grub) [USE flag/ru](/wiki/USE_flag/ru "USE flag/ru"), [systemd-boot](https://packages.gentoo.org/useflags/systemd-boot) [USE flag/ru](/wiki/USE_flag/ru "USE flag/ru"), [efistub](https://packages.gentoo.org/useflags/efistub) [USE flag/ru](/wiki/USE_flag/ru "USE flag/ru") and [uki](https://packages.gentoo.org/useflags/uki) [USE flag/ru](/wiki/USE_flag/ru "USE flag/ru") USE flags are **not** enabled. No further action is required.

### Initramfs

An **init** ial **ram**-based **f** ile **s** ystem, or [initramfs](/wiki/Initramfs "Initramfs"), may be required for a system to boot. A wide of variety of cases may necessitate one, but common cases include:

- Kernels where storage/filesystem drivers are modules.
- Layouts with /usr/ or /var/ on separate partitions.
- Encrypted root filesystems.

**Совет**

[Distribution kernels](/wiki/Project:Distribution_Kernel "Project:Distribution Kernel") are designed to be used with an initramfs, as many storage and filesystem drivers are built as modules.

In addition to mounting the root filesystem, an initramfs may also perform other tasks such as:

- Running **f** ile **s** ystem **c** onsistency chec **k** fsck, a tool to check and repair consistency of a file system in such events of uncleanly shutdown a system.
- Providing a recovery environment in the event of late-boot failures.

[Installkernel](/wiki/Installkernel "Installkernel") can automatically generate an initramfs when installing the kernel if the [dracut](https://packages.gentoo.org/useflags/dracut) [USE flag/ru](/wiki/USE_flag/ru "USE flag/ru") or [ugrd](https://packages.gentoo.org/useflags/ugrd) [USE flag/ru](/wiki/USE_flag/ru "USE flag/ru") USE flag is enabled:

ФАЙЛ **`/etc/portage/package.use/installkernel`**

```
sys-kernel/installkernel dracut

```

#### Chroot detection

Bootloaders such as [systemd-boot](/wiki/Systemd/systemd-boot "Systemd/systemd-boot") and [EFI stub](/wiki/EFI_stub "EFI stub") use the kernel arguments of the running system as set in /proc/cmdline by default. Because of the wide range of ways Gentoo can be installed users will randomly get tripped up by this.

To help solve any issues this may cause, [sys-kernel/installkernel](https://packages.gentoo.org/packages/sys-kernel/installkernel) will check if it is running in a chroot and fail if the kernel command line is not explicitly configured. Otherwise the bootloader would use the install media's boot arguments which would lead to boot failure.

One way to satisfy [sys-kernel/installkernel](https://packages.gentoo.org/packages/sys-kernel/installkernel) is by using Dracut's config file to set the root partition UUID as shown below, or alternatively for more information on what this check helps with and different ways to configure it, see [Installkernel#Install\_chroot\_check](/wiki/Installkernel#Install_chroot_check.2Fru "Installkernel").

`root #` `mkdir /etc/dracut.conf.d`

`root #` `blkid`

```
/dev/sda3: UUID="2122cd72-94d7-4dcc-821e-3705926deecc"
```

In the above example, the root partition is /dev/sda3 and the UUID is 2122cd72-94d7-4dcc-821e-3705926deecc.
The dracut config file would then look like:

ФАЙЛ **`/etc/dracut.conf.d/00-installkernel.conf`**

```
kernel_cmdline=" root=UUID=2122cd72-94d7-4dcc-821e-3705926deecc " # Note leading and trailing spaces

```

`root #` `emerge --ask sys-kernel/installkernel`

## Конфигурация и компиляция ядра

It can be a wise move to use the dist-kernel on the first boot as it provides a very simple method to rule out system issues and kernel config issues. Always having a known working kernel to fallback on can speed up debugging and alleviate anxiety when updating that your system will no longer boot.

A common misconception is that a manually compiled kernel will use a lot less RAM than a pre configured distribution kernel. Due to the modular nature of the Linux kernel, only what is needed by the system is loaded and in most cases resulting in similar memory usage.

**Важно**

Only one selection is required in the following subsection, if unsure of which to use then go with the first listed for now. It's always possible to switch at a later date if required.

От наименьшего вмешательства к наибольшему:

[Полностью автоматический подход: Distribution-ядра](/wiki/Handbook:HPPA/Installation/Kernel/ru#Distribution_kernels "Handbook:HPPA/Installation/Kernel/ru")[Проект Distribution Kernel](/wiki/Project:Distribution_Kernel "Project:Distribution Kernel") используется для конфигурации, автоматической сборки и установки ядра Linux, связанных с ним модулей и (опционально, но по умолчанию включено) файла initramfs. Новые обновления ядра полностью автоматизированы, поскольку они обрабатываются через менеджер пакетов, как и любой другой системный пакет. В случае необходимости [можно предоставить пользовательский конфигурационный файл ядра](/wiki/Project:Distribution_Kernel#Modifying_kernel_configuration "Project:Distribution Kernel"). Это наименее сложный процесс и идеально подходит для новых пользователей Gentoo, так как работает "из коробки" и требует минимального участия системного администратора.[Гибридный подход: Genkernel](/wiki/Handbook:HPPA/Installation/Kernel/ru#Alternative:_Genkernel "Handbook:HPPA/Installation/Kernel/ru")Новые обновления ядра устанавливаются через системный менеджер пакетов. Системные администраторы могут использовать инструмент Gentoo genkernel для общей конфигурации, автоматической сборки и установки ядра Linux, связанных с ним модулей и (опционально, но _**не**_ включено по умолчанию) файла initramfs. Можно предоставить пользовательский файл конфигурации ядра, если необходима кастомизация. Будущая конфигурация, сборка и установка ядра требуют участия системного администратора в виде выполнения eselect kernel, genkernel и, возможно, других команд для каждого обновления.[Полностью ручная настройка](/wiki/Handbook:HPPA/Installation/Kernel/ru#Alternative:_Manual_configuration "Handbook:HPPA/Installation/Kernel/ru")Новые исходные тексты ядра устанавливаются с помощью системного менеджера пакетов. Ядро конфигурируется, собирается и устанавливается вручную с помощью команды eselect kernel и множества команд make. С новыми обновлениями ядра повторяется ручной процесс конфигурирования, сборки и установки файлов ядра. Это самый сложный процесс, но он обеспечивает максимальный контроль над процессом обновления ядра.

Основой, вокруг которой строятся все дистрибутивы, является ядро Linux. Оно является прослойкой между пользовательскими программами и аппаратным обеспечением системы. Хотя руководство предоставляет своим пользователям несколько возможных источников ядра, более подробная информация с более детальным описанием доступна на странице [Пакеты ядра](/wiki/Kernel/Packages/ru "Kernel/Packages/ru").

**Совет**

Kernel installation tasks such as copying the kernel image to /boot or the [EFI System Partition](/wiki/EFI_System_Partition/ru "EFI System Partition/ru"), generating an [initramfs](/wiki/Initramfs "Initramfs") and/or [Unified Kernel Image](/wiki/Unified_Kernel_Image "Unified Kernel Image"), updating bootloader configuration, can be automated with [installkernel](/wiki/Installkernel "Installkernel"). Users may wish to configure and install [sys-kernel/installkernel](https://packages.gentoo.org/packages/sys-kernel/installkernel) before proceeding. See the [Kernel installation section below](/wiki/Handbook:HPPA/Installation/Kernel#Kernel_installation.2Fru "Handbook:HPPA/Installation/Kernel") for more more information.

### Distribution-ядра

_[Distribution-ядра](/wiki/Project:Distribution_Kernel "Project:Distribution Kernel")_ — это ebuild-файлы, которые охватывают полный процесс распаковки, конфигурирования, компиляции и установки ядра. Основным преимуществом этого метода является то, что ядра обновляются до новых версий менеджером пакетов во время обновления @world. Для этого используется только команда emerge. Distribution-ядра по умолчанию сконфигурированы для поддержки большинства оборудования, для более тонкой настройки предлагаются два механизма: saveconfig и сниппеты конфигурации. Смотрите страницу проекта для [более подробной информации о конфигурации.](/wiki/Project:Distribution_Kernel#Modifying_kernel_configuration "Project:Distribution Kernel")

##### Optional: Signed kernel modules

The kernel modules in the prebuilt distribution kernel ([sys-kernel/gentoo-kernel-bin](https://packages.gentoo.org/packages/sys-kernel/gentoo-kernel-bin)) are already signed. To sign the modules of kernels built from source enable the [modules-sign](https://packages.gentoo.org/useflags/modules-sign) [USE flag/ru](/wiki/USE_flag/ru "USE flag/ru") USE flag, and optionally specify which key to use for signing in [/etc/portage/make.conf](/wiki//etc/portage/make.conf/ru "/etc/portage/make.conf/ru"):

ФАЙЛ **`/etc/portage/make.conf`** **Enable module signing**

```
USE="modules-sign"
</div>

<div lang="en" dir="ltr" class="mw-content-ltr">
# Optionally, to use custom signing keys.
MODULES_SIGN_KEY="/path/to/kernel_key.pem"
MODULES_SIGN_CERT="/path/to/kernel_key.pem" # Only required if the MODULES_SIGN_KEY does not also contain the certificate.
MODULES_SIGN_HASH="sha512" # Defaults to sha512.

```

If MODULES\_SIGN\_KEY is not specified the kernel build system will generate a key, it will be stored in /usr/src/linux-x.y.z/certs. It is recommended to manually generate a key to ensure that it will be the same for each kernel release. A key may be generated with:

`root #` `openssl req -new -noenc -utf8 -sha256 -x509 -outform PEM -out kernel_key.pem -keyout kernel_key.pem`

**Заметка**

The MODULES\_SIGN\_KEY and MODULES\_SIGN\_CERT may be different files. For this example the pem file generated by OpenSSL includes both the key and the accompanying certificate, and thus both variables are set to the same value.

OpenSSL will ask some questions about the user generating the key, it is recommended to fill in these questions as detailed as possible.

Store the key in a safe location, at the very least the key should be readable only by the root user. Verify this with:

`root #` `ls -l kernel_key.pem`

```
 -r-------- 1 root root 3164 Jan  4 10:38 kernel_key.pem
```

If this outputs anything other then the above, correct the permissions with:

`root #` `chown root:root kernel_key.pem`

`root #` `chmod 400 kernel_key.pem`

#### Установка distribution-ядра

Чтобы собрать ядро из исходного кода с патчами Gentoo, введите:

`root #` `emerge --ask sys-kernel/gentoo-kernel`

Администраторы систем, которые хотят избежать сборки ядра из исходных текстов на компьютере, могут вместо этого использовать предварительно скомпилированные образы ядра:

`root #` `emerge --ask sys-kernel/gentoo-kernel-bin`

**Важно**

_[Distribution Kernels](/wiki/Project:Distribution_Kernel "Project:Distribution Kernel")_, such as [sys-kernel/gentoo-kernel](https://packages.gentoo.org/packages/sys-kernel/gentoo-kernel) and [sys-kernel/gentoo-kernel-bin](https://packages.gentoo.org/packages/sys-kernel/gentoo-kernel-bin), by default, expect to be installed alongside an [initramfs](/wiki/Handbook:HPPA/Installation/Kernel#Initramfs.2Fru "Handbook:HPPA/Installation/Kernel"). Before running emerge to install the kernel users should ensure that [sys-kernel/installkernel](https://packages.gentoo.org/packages/sys-kernel/installkernel) has been configured to utilize an initramfs generator (for example [Dracut](/wiki/Dracut "Dracut")) as described in the [installkernel section](/wiki/Handbook:HPPA/Installation/Kernel#Initramfs.2Fru "Handbook:HPPA/Installation/Kernel").

#### Обновление и очистка

После установки ядра менеджер пакетов будет автоматически обновлять его до более новых версий. Предыдущие версии будут храниться до тех пор, пока менеджер пакетов не получит запрос на очистку устаревших пакетов. Чтобы освободить место на диске, устаревшие пакеты можно удалить, периодически запуская emerge с опцией `--depclean`:

`root #` `emerge --depclean`

Также можно удалить именно устаревшие ядра:

`root #` `emerge --prune sys-kernel/gentoo-kernel sys-kernel/gentoo-kernel-bin`

**Совет**

By design, emerge only removes the kernel build directory. It does not actually remove the kernel modules, nor the installed kernel image. To completely clean-up old kernels, the [app-admin/eclean-kernel](https://packages.gentoo.org/packages/app-admin/eclean-kernel) tool may be used.

#### Задачи после установки/обновления

An upgrade of a distribution kernel is capable of triggering an automatic rebuild for external kernel modules installed by other packages (for example: [sys-fs/zfs-kmod](https://packages.gentoo.org/packages/sys-fs/zfs-kmod) or [x11-drivers/nvidia-drivers](https://packages.gentoo.org/packages/x11-drivers/nvidia-drivers)). This automated behaviour is enabled by enabling the [dist-kernel](https://packages.gentoo.org/useflags/dist-kernel) [USE flag/ru](/wiki/USE_flag/ru "USE flag/ru") USE flag. When required, this same flag will also trigger re-generation of the [initramfs](/wiki/Initramfs "Initramfs").

Включение этого USE-флага для таких пакетов, как [sys-fs/zfs](https://packages.gentoo.org/packages/sys-fs/zfs) и [sys-fs/zfs-kmod](https://packages.gentoo.org/packages/sys-fs/zfs-kmod) позволит им автоматически пересобираться в соответствии с обновленным ядром и, в случае необходимости, пересобирать initramfs.

ФАЙЛ **`/etc/portage/make.conf`** **Enabling USE=dist-kernel**

```
USE="dist-kernel"

```

##### Ручная пересборка initramfs или Unified Kernel Image

Если понадобится, вручную запустите перестройку, выполнив после обновления ядра команду:

`root #` `emerge --ask @module-rebuild`

Если какой-то модуль ядра (например, ZFS) необходим при ранней загрузке, пересоберите initramfs при помощи:

`root #` `emerge --config sys-kernel/gentoo-kernel
`

`root #` `emerge --config sys-kernel/gentoo-kernel-bin
`

After installing the Distribution Kernel successfully, it is now time to proceed to the next section: [Configuring the system](/wiki/Handbook:HPPA/Installation/System/ru "Handbook:HPPA/Installation/System/ru").

### Альтернатива: Ручная настройка

### Установка исходного кода ядра

При установке и компиляции ядра для систем на базе hppa Gentoo рекомендует использовать пакет [sys-kernel/gentoo-sources](https://packages.gentoo.org/packages/sys-kernel/gentoo-sources).

Выберите подходящий исходный код ядра и установите его с помощью emerge:

`root #` `emerge --ask sys-kernel/gentoo-sources`

Данная команда установит исходный код ядра Linux в /usr/src/, используя в названии версию ядра. Эта команда не установит автоматически символьную ссылку, пока вы не укажете USE-флаг [symlink](https://packages.gentoo.org/useflags/symlink) [USE flag/ru](/wiki/USE_flag/ru "USE flag/ru") для выбранного исходного кода ядра.

Обычно, символьная ссылка /usr/src/linux указывает на исходный код текущего работающего ядра. Однако, эта символьная ссылка не создаётся по умолчанию. Создать её поможет kernel модуль для eselect.

Чтобы подробнее узнать, зачем нужна эта символьная ссылка и как ею управлять, смотрите [Kernel/Upgrade](/wiki/Kernel/Upgrade/ru "Kernel/Upgrade/ru").

Для начала, просмотрите список установленных ядер (в виде исходного кода):

`root #` `eselect kernel list`

```
Available kernel symlink targets:
  [1]   linux-6.19.1-gentoo

```

Для того, чтобы создать символьную ссылку linux, используйте:

`root #` `eselect kernel set 1`

`root #` `ls -l /usr/src/linux`

```
lrwxrwxrwx    1 root   root    20 мар  3 22:44 /usr/src/linux -> linux-6.19.1-gentoo

```

Manually configuring a kernel is commonly seen as one of the most difficult procedures a system administrator has to perform. Nothing is less true - after configuring a few kernels no one remembers that it was difficult! There are two ways for a Gentoo user to manage a manual kernel system, both of which are listed below:

**Важно**

Only one selection is required in the following subsection, if unsure of which to use then go with the first listed for now. It's always possible to switch at a later date if required.

#### Процесс modprobed-db

A very easy way to manage the kernel is to first install [sys-kernel/gentoo-kernel-bin](https://packages.gentoo.org/packages/sys-kernel/gentoo-kernel-bin) and use the [sys-kernel/modprobed-db](https://packages.gentoo.org/packages/sys-kernel/modprobed-db) to collect information about what the system requires. modprobed-db is a tool which monitors the system via crontab to add all modules of all devices over the system's life to make sure it everything a user needs is supported. For example, if an Xbox controller is added after installation, then modprobed-db will add the modules to be built next time the kernel is rebuilt.

For now please follow installing [sys-kernel/gentoo-kernel-bin](https://packages.gentoo.org/packages/sys-kernel/gentoo-kernel-bin) via [Distribution\_kernels](/wiki/Handbook:HPPA/Installation/Kernel#Distribution_kernels.2Fru "Handbook:HPPA/Installation/Kernel").

Next, install [sys-kernel/modprobed-db](https://packages.gentoo.org/packages/sys-kernel/modprobed-db):

`root #` `emerge --ask sys-kernel/modprobed-db`

Please watch out for further steps related to modprobed-db in the Handbook.

More on this topic can be found in the [Modprobed-db](/wiki/Modprobed-db "Modprobed-db") article.
}}

##### Option 2 - Assisted manual process

This method allows a user to have full control of how their kernel is built with as minimal help from outside tools as they wish. Some could consider this as making it hard for the sake of it.

Однако одна вещь является истиной: при ручной конфигурации ядра очень важно понимать свою систему. Большую часть сведений можно почерпнуть, установив пакет [sys-apps/pciutils](https://packages.gentoo.org/packages/sys-apps/pciutils), который содержит в команду lspci:

`root #` `emerge --ask sys-apps/pciutils`

**Заметка**

Находясь внутри изолированного окружения chroot, можно спокойно игнорировать любые предупреждения pcilib (например, _pcilib: cannot open /sys/bus/pci/devices_), которые могут появляться в выводе lspci.

Другим источником информации о системе может стать вывод команды lsmod, по которому можно понять, какие модули ядра использует установочный носитель, чтобы потом включить аналогичные настройки.

Теперь необходимо перейти в каталог с ядром.

`root #` `cd /usr/src/linux
`

**Совет**

To view the full list of make arguments available for the kernel, run `make help`.

The kernel has a method of autodetecting the modules currently being used on the installcd which will give a great starting point to allow a user to configure their own. This can be called by using:

`root #` `make localmodconfig`

Manually configuring should not be needed at this point, but if a user wish to check:

`root #` `make nconfig`

Now it's time to decide if modules signing is required in the steps listed in [here](/wiki/Handbook:HPPA/Installation/Kernel#Optional:_Signed_kernel_modules_2.2Fru "Handbook:HPPA/Installation/Kernel")

If not, proceed to building described [here](/wiki/Handbook:HPPA/Installation/Kernel#Compiling_and_installing.2Fru "Handbook:HPPA/Installation/Kernel")

##### Option 3 - Configuring by hand

The Linux kernel configuration has many, many sections and as configuring a kernel by hand is rarely needed thanks to the two tools listed above. The steps to do it by hand are now included at [Kernel/Gentoo\_Kernel\_Configuration\_Guide](/wiki/Kernel/Gentoo_Kernel_Configuration_Guide/ru "Kernel/Gentoo Kernel Configuration Guide/ru")

#### Optional: Signed kernel modules

To automatically sign the kernel modules enable CONFIG\_MODULE\_SIG\_ALL:

ЯДРО **Sign kernel modules CONFIG\_MODULE\_SIG\_ALL**

```
[*] Enable loadable module support
  -*-   Module signature verification
    [*]     Automatically sign all modules
    Which hash algorithm should modules be signed with? (Sign modules with SHA-512) --->

```

Optionally change the hash algorithm if desired.

To enforce that all modules are signed with a valid signature, enable CONFIG\_MODULE\_SIG\_FORCE as well:

ЯДРО **Enforce signed kernel modules CONFIG\_MODULE\_SIG\_FORCE**

```
[*] Enable loadable module support
  -*-   Module signature verification
    [*]     Require modules to be validly signed
    [*]     Automatically sign all modules
    Which hash algorithm should modules be signed with? (Sign modules with SHA-512) --->

```

To use a custom key, specify the location of this key in CONFIG\_MODULE\_SIG\_KEY. If unspecified, the kernel build system will generate a key. It is recommended to generate one manually instead. This can be done with:

`root #` `openssl req -new -nodes -utf8 -sha256 -x509 -outform PEM -out kernel_key.pem -keyout kernel_key.pem`

OpenSSL will ask some questions about the user generating the key, it is recommended to fill in these questions as detailed as possible.

Store the key in a safe location, at the very least the key should be readable only by the root user. Verify this with:

`root #` `ls -l kernel_key.pem`

```
 -r-------- 1 root root 3164 Jan  4 10:38 kernel_key.pem
```

If this outputs anything other then the above, correct the permissions with:

`root #` `chown root:root kernel_key.pem
`

`root #` `chmod 400 kernel_key.pem
`

ЯДРО **Specify signing key CONFIG\_MODULE\_SIG\_KEY**

```
-*- Cryptographic API  --->
  Certificates for signature checking  --->
    (/path/to/kernel_key.pem) File name or PKCS#11 URI of module signing key

```

To also sign external kernel modules installed by other packages via `linux-mod-r1.eclass`, enable the [modules-sign](https://packages.gentoo.org/useflags/modules-sign) [USE flag/ru](/wiki/USE_flag/ru "USE flag/ru") USE flag globally:

ФАЙЛ **`/etc/portage/make.conf`** **Включение подписи модулей**

```
USE="modules-sign"

<div lang="en" dir="ltr" class="mw-content-ltr">
# Optionally, when using custom signing keys.
MODULES_SIGN_KEY="/path/to/kernel_key.pem"
MODULES_SIGN_CERT="/path/to/kernel_key.pem" # Only required if the MODULES_SIGN_KEY does not also contain the certificate
MODULES_SIGN_HASH="sha512" # Defaults to sha512

```

**Заметка**

MODULES\_SIGN\_KEY and MODULES\_SIGN\_CERT may point to different files. For this example, the pem file generated by OpenSSL includes both the key and the accompanying certificate, and thus both variables are set to the same value.

### Компиляция и установка

**Важно**

Чтобы скомпилировать 64-битное ядро, необходимо сначала установить [sys-devel/kgcc64](https://packages.gentoo.org/packages/sys-devel/kgcc64). Ранее использование 64-битного ядра не рекомендовалось, но теперь это вполне безопасно. Если всё-таки есть сомнения, используйте 64-битное ядро, если в системе более 4 Гб ОЗУ, либо если вашему серверу это необходимо, например, на A500.

Когда настройка закончена, настало время скомпилировать и установить ядро. Выйдите из настройки и запустите процесс компиляции:

`root #` `make && make modules_install`

В случае сборки 64-битного ядра, вместо этого выполните следующее (это необходимо выполнить даже в случае нативных сборок, см. [здесь](https://www.spinics.net/lists/kernel/msg3987729.html)):

`root #` `CROSS_COMPILE=hppa64-unknown-linux-gnu- make && CROSS_COMPILE=hppa64-unknown-linux-gnu- make modules_install`

**Заметка**

Можно включить параллельную сборку, используя make -jX, где `X` — это число параллельных задач, которые может запустить процесс сборки. Это похоже на инструкции, которые были даны ранее относительно файла /etc/portage/make.conf в части переменной MAKEOPTS

Когда процесс компиляции будет завершён, скопируйте образ ядра в каталог /boot/. Используйте любое имя, кажущееся вам наиболее подходящим для ядра, но не забудьте его, так как оно нам потребуется позже при настройке загрузчика. Не забудьте заменить kernel-6.19.1-gentoo на имя и версию своего ядра.

`root #` `cp vmlinux /boot/kernel-6.19.1-gentoo`

Продолжите установку с раздела [Настройка системы](/wiki/Handbook:HPPA/Installation/System/ru "Handbook:HPPA/Installation/System/ru").

[← Установка базовой системы Gentoo](/wiki/Handbook:HPPA/Installation/Base/ru "Предыдущая часть") [К содержанию](/wiki/Handbook:HPPA/ru "Handbook:HPPA/ru") [Настройка системы →](/wiki/Handbook:HPPA/Installation/System/ru "Следующая часть")