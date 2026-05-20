# Kernel

Other languages:

- [Deutsch](/wiki/Handbook:AMD64/Installation/Kernel/de "Handbook:AMD64/Installation/Kernel/de (100% translated)")
- [English](/wiki/Handbook:AMD64/Installation/Kernel "Handbook:AMD64/Installation/Kernel (100% translated)")
- [español](/wiki/Handbook:AMD64/Installation/Kernel/es "Handbook:AMD64/Installation/Kernel/es (100% translated)")
- [français](/wiki/Handbook:AMD64/Installation/Kernel/fr "Handbook:AMD64/Installation/Kernel/fr (100% translated)")
- [italiano](/wiki/Handbook:AMD64/Installation/Kernel/it "Handbook:AMD64/Installation/Kernel/it (100% translated)")
- [magyar](/wiki/Handbook:AMD64/Installation/Kernel/hu "Handbook:AMD64/Installation/Kernel/hu (100% translated)")
- [polski](/wiki/Handbook:AMD64/Installation/Kernel/pl "Handbook:AMD64/Installation/Kernel/pl (100% translated)")
- [português do Brasil](/wiki/Handbook:AMD64/Installation/Kernel/pt-br "Handbook:AMD64/Installation/Kernel/pt-br (100% translated)")
- [čeština](/wiki/Handbook:AMD64/Installation/Kernel/cs "Handbook:AMD64/Installation/Kernel/cs (100% translated)")
- русский
- [தமிழ்](/wiki/Handbook:AMD64/Installation/Kernel/ta "Handbook:AMD64/Installation/Kernel/ta (100% translated)")
- [中文（中国大陆）‎](/wiki/Handbook:AMD64/Installation/Kernel/zh-cn "Handbook:AMD64/Installation/Kernel/zh-cn (100% translated)")
- [日本語](/wiki/Handbook:AMD64/Installation/Kernel/ja "Handbook:AMD64/Installation/Kernel/ja (100% translated)")
- [한국어](/wiki/Handbook:AMD64/Installation/Kernel/ko "Handbook:AMD64/Installation/Kernel/ko (100% translated)")

[AMD64 Handbook](/wiki/Handbook:AMD64/ru "Handbook:AMD64/ru")[Установка](/wiki/Handbook:AMD64/Full/Installation/ru "Handbook:AMD64/Full/Installation/ru")[Об установке](/wiki/Handbook:AMD64/Installation/About/ru "Handbook:AMD64/Installation/About/ru")[Выбор подходящего источника для установки](/wiki/Handbook:AMD64/Installation/Media/ru "Handbook:AMD64/Installation/Media/ru")[Настройка сети](/wiki/Handbook:AMD64/Installation/Networking/ru "Handbook:AMD64/Installation/Networking/ru")[Подготовка дисков](/wiki/Handbook:AMD64/Installation/Disks/ru "Handbook:AMD64/Installation/Disks/ru")[Установка файла stage](/wiki/Handbook:AMD64/Installation/Stage/ru "Handbook:AMD64/Installation/Stage/ru")[Установка базовой системы](/wiki/Handbook:AMD64/Installation/Base/ru "Handbook:AMD64/Installation/Base/ru")Настройка ядра[Настройка системы](/wiki/Handbook:AMD64/Installation/System/ru "Handbook:AMD64/Installation/System/ru")[Установка системных утилит](/wiki/Handbook:AMD64/Installation/Tools/ru "Handbook:AMD64/Installation/Tools/ru")[Настройка загрузчика](/wiki/Handbook:AMD64/Installation/Bootloader/ru "Handbook:AMD64/Installation/Bootloader/ru")[Завершение](/wiki/Handbook:AMD64/Installation/Finalizing/ru "Handbook:AMD64/Installation/Finalizing/ru")[Работа с Gentoo](/wiki/Handbook:AMD64/Full/Working/ru "Handbook:AMD64/Full/Working/ru")[Введение в Portage](/wiki/Handbook:AMD64/Working/Portage/ru "Handbook:AMD64/Working/Portage/ru")[USE-флаги](/wiki/Handbook:AMD64/Working/USE/ru "Handbook:AMD64/Working/USE/ru")[Возможности Portage](/wiki/Handbook:AMD64/Working/Features/ru "Handbook:AMD64/Working/Features/ru")[Система сценариев инициализации](/wiki/Handbook:AMD64/Working/Initscripts/ru "Handbook:AMD64/Working/Initscripts/ru")[Переменные окружения](/wiki/Handbook:AMD64/Working/EnvVar/ru "Handbook:AMD64/Working/EnvVar/ru")[Работа с Portage](/wiki/Handbook:AMD64/Full/Portage/ru "Handbook:AMD64/Full/Portage/ru")[Файлы и каталоги](/wiki/Handbook:AMD64/Portage/Files/ru "Handbook:AMD64/Portage/Files/ru")[Переменные](/wiki/Handbook:AMD64/Portage/Variables/ru "Handbook:AMD64/Portage/Variables/ru")[Смешение ветвей программного обеспечения](/wiki/Handbook:AMD64/Portage/Branches/ru "Handbook:AMD64/Portage/Branches/ru")[Дополнительные утилиты](/wiki/Handbook:AMD64/Portage/Tools/ru "Handbook:AMD64/Portage/Tools/ru")[Дополнительные репозитории пакетов](/wiki/Handbook:AMD64/Portage/CustomTree/ru "Handbook:AMD64/Portage/CustomTree/ru")[Расширенные возможности](/wiki/Handbook:AMD64/Portage/Advanced/ru "Handbook:AMD64/Portage/Advanced/ru")[Настройка сети OpenRC](/wiki/Handbook:AMD64/Full/Networking/ru "Handbook:AMD64/Full/Networking/ru")[Начальная настройка](/wiki/Handbook:AMD64/Networking/Introduction/ru "Handbook:AMD64/Networking/Introduction/ru")[Расширенная настройка](/wiki/Handbook:AMD64/Networking/Advanced/ru "Handbook:AMD64/Networking/Advanced/ru")[Модульное построение сети](/wiki/Handbook:AMD64/Networking/Modular/ru "Handbook:AMD64/Networking/Modular/ru")[Беспроводная сеть](/wiki/Handbook:AMD64/Networking/Wireless/ru "Handbook:AMD64/Networking/Wireless/ru")[Добавляем функциональность](/wiki/Handbook:AMD64/Networking/Extending/ru "Handbook:AMD64/Networking/Extending/ru")[Динамическое управление](/wiki/Handbook:AMD64/Networking/Dynamic/ru "Handbook:AMD64/Networking/Dynamic/ru")

## Contents

- [1Необязательно: Установка файлов прошивки и/или микрокода](#.D0.9D.D0.B5.D0.BE.D0.B1.D1.8F.D0.B7.D0.B0.D1.82.D0.B5.D0.BB.D1.8C.D0.BD.D0.BE:_.D0.A3.D1.81.D1.82.D0.B0.D0.BD.D0.BE.D0.B2.D0.BA.D0.B0_.D1.84.D0.B0.D0.B9.D0.BB.D0.BE.D0.B2_.D0.BF.D1.80.D0.BE.D1.88.D0.B8.D0.B2.D0.BA.D0.B8_.D0.B8.2F.D0.B8.D0.BB.D0.B8_.D0.BC.D0.B8.D0.BA.D1.80.D0.BE.D0.BA.D0.BE.D0.B4.D0.B0)
  - [1.1Файлы прошивки](#.D0.A4.D0.B0.D0.B9.D0.BB.D1.8B_.D0.BF.D1.80.D0.BE.D1.88.D0.B8.D0.B2.D0.BA.D0.B8)
    - [1.1.1Рекомендуется: Linux Firmware](#.D0.A0.D0.B5.D0.BA.D0.BE.D0.BC.D0.B5.D0.BD.D0.B4.D1.83.D0.B5.D1.82.D1.81.D1.8F:_Linux_Firmware)
      - [1.1.1.1Загрузка прошивки](#.D0.97.D0.B0.D0.B3.D1.80.D1.83.D0.B7.D0.BA.D0.B0_.D0.BF.D1.80.D0.BE.D1.88.D0.B8.D0.B2.D0.BA.D0.B8)
    - [1.1.2SOF Firmware](#SOF_Firmware)
  - [1.2Микрокод](#.D0.9C.D0.B8.D0.BA.D1.80.D0.BE.D0.BA.D0.BE.D0.B4)
- [2sys-kernel/installkernel](#sys-kernel.2Finstallkernel)
  - [2.1Начальный загрузчик](#.D0.9D.D0.B0.D1.87.D0.B0.D0.BB.D1.8C.D0.BD.D1.8B.D0.B9_.D0.B7.D0.B0.D0.B3.D1.80.D1.83.D0.B7.D1.87.D0.B8.D0.BA)
    - [2.1.1GRUB](#GRUB)
    - [2.1.2systemd-boot](#systemd-boot)
    - [2.1.3EFI stub](#EFI_stub)
    - [2.1.4Traditional layout, other bootloaders (e.g. (e)lilo, syslinux, etc.)](#Traditional_layout.2C_other_bootloaders_.28e.g._.28e.29lilo.2C_syslinux.2C_etc..29)
  - [2.2Initramfs](#Initramfs)
    - [2.2.1Chroot detection](#Chroot_detection)
  - [2.3Дополнительно: Unified Kernel Image](#.D0.94.D0.BE.D0.BF.D0.BE.D0.BB.D0.BD.D0.B8.D1.82.D0.B5.D0.BB.D1.8C.D0.BD.D0.BE:_Unified_Kernel_Image)
    - [2.3.1Общий Unified Kernel Image (только для systemd)](#.D0.9E.D0.B1.D1.89.D0.B8.D0.B9_Unified_Kernel_Image_.28.D1.82.D0.BE.D0.BB.D1.8C.D0.BA.D0.BE_.D0.B4.D0.BB.D1.8F_systemd.29)
    - [2.3.2Secure Boot](#Secure_Boot)
- [3Конфигурация и компиляция ядра](#.D0.9A.D0.BE.D0.BD.D1.84.D0.B8.D0.B3.D1.83.D1.80.D0.B0.D1.86.D0.B8.D1.8F_.D0.B8_.D0.BA.D0.BE.D0.BC.D0.BF.D0.B8.D0.BB.D1.8F.D1.86.D0.B8.D1.8F_.D1.8F.D0.B4.D1.80.D0.B0)
  - [3.1Distribution-ядра](#Distribution-.D1.8F.D0.B4.D1.80.D0.B0)
    - [3.1.1Optional: Signed kernel modules](#Optional:_Signed_kernel_modules)
    - [3.1.2Optional: Signing the kernel image (Secure Boot)](#Optional:_Signing_the_kernel_image_.28Secure_Boot.29)
    - [3.1.3Установка distribution-ядра](#.D0.A3.D1.81.D1.82.D0.B0.D0.BD.D0.BE.D0.B2.D0.BA.D0.B0_distribution-.D1.8F.D0.B4.D1.80.D0.B0)
    - [3.1.4Обновление и очистка](#.D0.9E.D0.B1.D0.BD.D0.BE.D0.B2.D0.BB.D0.B5.D0.BD.D0.B8.D0.B5_.D0.B8_.D0.BE.D1.87.D0.B8.D1.81.D1.82.D0.BA.D0.B0)
    - [3.1.5Задачи после установки/обновления](#.D0.97.D0.B0.D0.B4.D0.B0.D1.87.D0.B8_.D0.BF.D0.BE.D1.81.D0.BB.D0.B5_.D1.83.D1.81.D1.82.D0.B0.D0.BD.D0.BE.D0.B2.D0.BA.D0.B8.2F.D0.BE.D0.B1.D0.BD.D0.BE.D0.B2.D0.BB.D0.B5.D0.BD.D0.B8.D1.8F)
      - [3.1.5.1Ручная пересборка initramfs или Unified Kernel Image](#.D0.A0.D1.83.D1.87.D0.BD.D0.B0.D1.8F_.D0.BF.D0.B5.D1.80.D0.B5.D1.81.D0.B1.D0.BE.D1.80.D0.BA.D0.B0_initramfs_.D0.B8.D0.BB.D0.B8_Unified_Kernel_Image)
  - [3.2Альтернатива: Ручная настройка](#.D0.90.D0.BB.D1.8C.D1.82.D0.B5.D1.80.D0.BD.D0.B0.D1.82.D0.B8.D0.B2.D0.B0:_.D0.A0.D1.83.D1.87.D0.BD.D0.B0.D1.8F_.D0.BD.D0.B0.D1.81.D1.82.D1.80.D0.BE.D0.B9.D0.BA.D0.B0)
  - [3.3Установка исходного кода ядра](#.D0.A3.D1.81.D1.82.D0.B0.D0.BD.D0.BE.D0.B2.D0.BA.D0.B0_.D0.B8.D1.81.D1.85.D0.BE.D0.B4.D0.BD.D0.BE.D0.B3.D0.BE_.D0.BA.D0.BE.D0.B4.D0.B0_.D1.8F.D0.B4.D1.80.D0.B0)
    - [3.3.1Процесс modprobed-db](#.D0.9F.D1.80.D0.BE.D1.86.D0.B5.D1.81.D1.81_modprobed-db)
      - [3.3.1.1Option 2 - Assisted manual process](#Option_2_-_Assisted_manual_process)
      - [3.3.1.2Option 3 - Configuring by hand](#Option_3_-_Configuring_by_hand)
    - [3.3.2Optional: Signed kernel modules](#Optional:_Signed_kernel_modules_2)
    - [3.3.3Optional: Signing the kernel image (Secure Boot)](#Optional:_Signing_the_kernel_image_.28Secure_Boot.29_2)
    - [3.3.4Компиляция и установка](#.D0.9A.D0.BE.D0.BC.D0.BF.D0.B8.D0.BB.D1.8F.D1.86.D0.B8.D1.8F_.D0.B8_.D1.83.D1.81.D1.82.D0.B0.D0.BD.D0.BE.D0.B2.D0.BA.D0.B0)

## Необязательно: Установка файлов прошивки и/или микрокода

### Файлы прошивки

#### Рекомендуется: Linux Firmware

On many systems, non-FOSS firmware is required for certain hardware to function. The [sys-kernel/linux-firmware](https://packages.gentoo.org/packages/sys-kernel/linux-firmware) package contains firmware for many, but not all, devices.

**Совет**

Most wireless cards and GPUs require firmware to function.

`root #` `emerge --ask sys-kernel/linux-firmware`

**Заметка**

Установка определённых пакетов прошивок часто требует принятия соответствующих лицензий на прошивку. При необходимости посетите раздел руководства [о принятии лицензии](/wiki/Handbook:AMD64/Working/Portage/ru#Licenses "Handbook:AMD64/Working/Portage/ru") для получения помощи.

##### Загрузка прошивки

Firmware files are typically loaded when the associated kernel module is loaded. This means the firmware must be built into the kernel using **CONFIG\_EXTRA\_FIRMWARE** if the kernel module is set to _Y_ instead of _M_. In most cases, building-in a module which required firmware can complicate or break loading.

#### SOF Firmware

Sound Open Firmware (SOF) is a new open source audio driver meant to replace the proprietary Smart Sound Technology (SST) audio driver from Intel. 10th gen+ and Apollo Lake (Atom E3900, Celeron N3350, and Pentium N4200) Intel CPUs require this firmware for certain features and certain AMD APUs also have support for this firmware. SOF's supported platforms matrix can be found [here](https://thesofproject.github.io/latest/platforms/index.html) for more information.

`root #` `emerge --ask sys-firmware/sof-firmware`

### Микрокод

Вдобавок к сетевому оборудованию и видеокартам, процессоры также могут требовать обновления прошивки. Обычно подобный вид прошивок называется _микрокодом_. Обновления микрокода иногда нужны, чтобы исправить нестабильность, улучшить безопасность, или исправить прочие разнообразные баги в процессоре.

Обновления микрокода для процессоров AMD распространяются в вышеупомянутом пакете [sys-kernel/linux-firmware](https://packages.gentoo.org/packages/sys-kernel/linux-firmware). Обновления микрокода для процессоров Intel находятся в пакете [sys-firmware/intel-microcode](https://packages.gentoo.org/packages/sys-firmware/intel-microcode), который необходимо установить отдельно. См. [статью Микрокод](/wiki/Microcode "Microcode") для получения дополнительной информации о том, как применять обновления микрокода.

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

#### systemd-boot

When using [systemd-boot](/wiki/Systemd/systemd-boot "Systemd/systemd-boot") (formerly gummiboot) as the bootloader, systemd's kernel-install must be used. Therefore ensure the [systemd](https://packages.gentoo.org/useflags/systemd) [USE flag/ru](/wiki/USE_flag/ru "USE flag/ru") and the [systemd-boot](https://packages.gentoo.org/useflags/systemd-boot) [USE flag/ru](/wiki/USE_flag/ru "USE flag/ru") USE flags are enabled on [sys-kernel/installkernel](https://packages.gentoo.org/packages/sys-kernel/installkernel), and then install the relevant package for systemd-boot.

ФАЙЛ **`/etc/portage/package.use/systemd-boot`**

```
sys-apps/systemd-utils boot kernel-install
sys-kernel/installkernel systemd systemd-boot

```

The kernel command line to use for new kernels should be specified in /etc/kernel/cmdline and /etc/cmdline. As they both need to be updated together, it make sense to symlink the two files now.

ФАЙЛ **`/etc/kernel/cmdline`**

```
quiet splash

```

`root #` `mkdir -p /etc/cmdline.d`

`root #` `ln -s /etc/kernel/cmdline /etc/cmdline.d/00-installkernel.conf`

`root #` `emerge --ask sys-kernel/installkernel`

}}

**Заметка**

systemd-boot requires kernels to be installed to /efi.

#### EFI stub

UEFI-based computer systems technically do not need secondary bootloaders in order to boot kernels. Secondary bootloaders exist to _extend_ the functionality of UEFI firmware during the boot process. That being said, using a secondary bootloader is typically easier and more robust because it offers a more flexible approach for quickly modifying kernel parameters at boot time. Note also that UEFI implementations strongly differ between vendors and between models and there is no guarantee that a given firmware follows the UEFI specification. Therefore, EFI Stub booting is not guaranteed to work on every UEFI-based system.

`root #` `emerge --ask sys-kernel/installkernel`

`root #` `mkdir -p /efi/EFI/Gentoo`

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

### Дополнительно: Unified Kernel Image

A [Unified Kernel Image](/wiki/Unified_Kernel_Image "Unified Kernel Image") (UKI) combines, among other things, the kernel, the initramfs and the kernel command line into a single executable. Since the kernel command line is embedded into the unified kernel image, it should be specified before generating the unified kernel image (see below). Note that any kernel command line arguments supplied by the bootloader or firmware at boot are ignored when booting with secure boot enabled.

A unified kernel image requires a stub loader. Currently, the only one available is systemd-stub. To enable it:

For systemd systems:

ФАЙЛ **`/etc/portage/package.use/uki`**

```
sys-apps/systemd boot

```

`root #` `emerge --ask sys-apps/systemd`

For OpenRC systems:

ФАЙЛ **`/etc/portage/package.use/uki`**

```
sys-apps/systemd-utils boot kernel-install

```

`root #` `emerge --ask sys-apps/systemd-utils`

[Installkernel](/wiki/Installkernel "Installkernel") can automatically generate a unified kernel image using either [dracut](/wiki/Unified_kernel_image#dracut.2Fru "Unified kernel image") or [ukify](/wiki/Unified_kernel_image#ukify.2Fru "Unified kernel image") by enabling the respective flag and the [uki](https://packages.gentoo.org/useflags/uki) [USE flag/ru](/wiki/USE_flag/ru "USE flag/ru") USE flag.

For dracut:

ФАЙЛ **`/etc/portage/package.use/uki`**

```
sys-kernel/installkernel dracut uki

```

ФАЙЛ **`/etc/dracut.conf.d/uki.conf`**

```
uefi="yes"
kernel_cmdline="some-kernel-command-line-arguments"

```

`root #` `emerge --ask sys-kernel/installkernel`

For ukify:

ФАЙЛ **`/etc/portage/package.use/uki`**

```
sys-apps/systemd boot ukify                         # Для систем systemd
sys-apps/systemd-utils kernel-install boot ukify    # Для систем OpenRC
sys-kernel/installkernel dracut ukify uki

```

ФАЙЛ **`/etc/kernel/cmdline`**

```
some-kernel-command-line-arguments

```

`root #` `emerge --ask sys-kernel/installkernel`

Note that while dracut can generate both an initramfs and a unified kernel image, ukify can only generate the latter and therefore the initramfs must be generated separately with dracut.

**Важно**

In the above configuration examples (for both Dracut and ukify) it is important to specify at least an appropriate root= parameter for the kernel command line to ensure that the Unified Kernel Image can find the root partition. This is not required for systemd based systems following the Discoverable Partitions Specification (DPS), in that case the embedded initramfs will be able to dynamically find the root partition.

#### Общий Unified Kernel Image (только для systemd)

The prebuilt [sys-kernel/gentoo-kernel-bin](https://packages.gentoo.org/packages/sys-kernel/gentoo-kernel-bin) can optionally install a prebuilt generic unified kernel image containing a generic initramfs that is able to boot most systemd based systems. It can be installed by enabling the [generic-uki](https://packages.gentoo.org/useflags/generic-uki) [USE flag/ru](/wiki/USE_flag/ru "USE flag/ru") USE flag, and configuring [installkernel](/wiki/Installkernel "Installkernel") to not generate a custom initramfs or unified kernel image:

ФАЙЛ **`/etc/portage/package.use/uki`**

```
sys-kernel/gentoo-kernel-bin generic-uki
sys-kernel/installkernel -dracut -ukify -ugrd uki

```

#### Secure Boot

**Предупреждение**

If following this section and manually compiling your own kernel, then make sure to follow the steps outlined in [Signing the kernel](/wiki/Kernel#Optional:_Signing_the_kernel_image_.28Secure_Boot.29.2Fru "Kernel")

The generic Unified Kernel Image optionally distributed by [sys-kernel/gentoo-kernel-bin](https://packages.gentoo.org/packages/sys-kernel/gentoo-kernel-bin) is already pre-signed. How to sign a locally generated unified kernel image depends on whether dracut or ukify is used. Note that the location of the key and certificate should be the same as the SECUREBOOT\_SIGN\_KEY and SECUREBOOT\_SIGN\_CERT as specified in /etc/portage/make.conf.

Для dracut:

ФАЙЛ **`/etc/dracut.conf.d/uki.conf`**

```
uefi="yes"
kernel_cmdline="some-kernel-command-line-arguments"
uefi_secureboot_key="/path/to/kernel_key.pem"
uefi_secureboot_cert="/path/to/kernel_key.pem"

```

Для ukify:

ФАЙЛ **`/etc/kernel/uki.conf`**

```
[UKI]
SecureBootPrivateKey=/path/to/kernel_key.pem
SecureBootCertificate=/path/to/kernel_key.pem

```

## Конфигурация и компиляция ядра

It can be a wise move to use the dist-kernel on the first boot as it provides a very simple method to rule out system issues and kernel config issues. Always having a known working kernel to fallback on can speed up debugging and alleviate anxiety when updating that your system will no longer boot.

A common misconception is that a manually compiled kernel will use a lot less RAM than a pre configured distribution kernel. Due to the modular nature of the Linux kernel, only what is needed by the system is loaded and in most cases resulting in similar memory usage.

**Важно**

Only one selection is required in the following subsection, if unsure of which to use then go with the first listed for now. It's always possible to switch at a later date if required.

От наименьшего вмешательства к наибольшему:

[Полностью автоматический подход: Distribution-ядра](/wiki/Handbook:AMD64/Installation/Kernel/ru#Distribution_kernels "Handbook:AMD64/Installation/Kernel/ru")[Проект Distribution Kernel](/wiki/Project:Distribution_Kernel "Project:Distribution Kernel") используется для конфигурации, автоматической сборки и установки ядра Linux, связанных с ним модулей и (опционально, но по умолчанию включено) файла initramfs. Новые обновления ядра полностью автоматизированы, поскольку они обрабатываются через менеджер пакетов, как и любой другой системный пакет. В случае необходимости [можно предоставить пользовательский конфигурационный файл ядра](/wiki/Project:Distribution_Kernel#Modifying_kernel_configuration "Project:Distribution Kernel"). Это наименее сложный процесс и идеально подходит для новых пользователей Gentoo, так как работает "из коробки" и требует минимального участия системного администратора.[Гибридный подход: Genkernel](/wiki/Handbook:AMD64/Installation/Kernel/ru#Alternative:_Genkernel "Handbook:AMD64/Installation/Kernel/ru")Новые обновления ядра устанавливаются через системный менеджер пакетов. Системные администраторы могут использовать инструмент Gentoo genkernel для общей конфигурации, автоматической сборки и установки ядра Linux, связанных с ним модулей и (опционально, но _**не**_ включено по умолчанию) файла initramfs. Можно предоставить пользовательский файл конфигурации ядра, если необходима кастомизация. Будущая конфигурация, сборка и установка ядра требуют участия системного администратора в виде выполнения eselect kernel, genkernel и, возможно, других команд для каждого обновления.[Полностью ручная настройка](/wiki/Handbook:AMD64/Installation/Kernel/ru#Alternative:_Manual_configuration "Handbook:AMD64/Installation/Kernel/ru")Новые исходные тексты ядра устанавливаются с помощью системного менеджера пакетов. Ядро конфигурируется, собирается и устанавливается вручную с помощью команды eselect kernel и множества команд make. С новыми обновлениями ядра повторяется ручной процесс конфигурирования, сборки и установки файлов ядра. Это самый сложный процесс, но он обеспечивает максимальный контроль над процессом обновления ядра.

Основой, вокруг которой строятся все дистрибутивы, является ядро Linux. Оно является прослойкой между пользовательскими программами и аппаратным обеспечением системы. Хотя руководство предоставляет своим пользователям несколько возможных источников ядра, более подробная информация с более детальным описанием доступна на странице [Пакеты ядра](/wiki/Kernel/Packages/ru "Kernel/Packages/ru").

**Совет**

Kernel installation tasks such as copying the kernel image to /boot or the [EFI System Partition](/wiki/EFI_System_Partition/ru "EFI System Partition/ru"), generating an [initramfs](/wiki/Initramfs "Initramfs") and/or [Unified Kernel Image](/wiki/Unified_Kernel_Image "Unified Kernel Image"), updating bootloader configuration, can be automated with [installkernel](/wiki/Installkernel "Installkernel"). Users may wish to configure and install [sys-kernel/installkernel](https://packages.gentoo.org/packages/sys-kernel/installkernel) before proceeding. See the [Kernel installation section below](/wiki/Handbook:AMD64/Installation/Kernel#Kernel_installation.2Fru "Handbook:AMD64/Installation/Kernel") for more more information.

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

##### Optional: Signing the kernel image (Secure Boot)

The kernel image in the prebuilt distribution kernel ([sys-kernel/gentoo-kernel-bin](https://packages.gentoo.org/packages/sys-kernel/gentoo-kernel-bin)) is already signed for use with [Secure Boot](/wiki/Secure_Boot "Secure Boot"). To sign the kernel image of kernels built from source enable the [secureboot](https://packages.gentoo.org/useflags/secureboot) [USE flag/ru](/wiki/USE_flag/ru "USE flag/ru") USE flag, and optionally specify which key to use for signing in [/etc/portage/make.conf](/wiki//etc/portage/make.conf/ru "/etc/portage/make.conf/ru"). Note that signing the kernel image for use with secureboot requires that the kernel modules are also signed, the same key may be used to sign both the kernel image and the kernel modules:

ФАЙЛ **`/etc/portage/make.conf`** **Enable custom signing keys**

```
USE="modules-sign secureboot"
</div>

<div lang="en" dir="ltr" class="mw-content-ltr">
# Optionally, to use custom signing keys.
MODULES_SIGN_KEY="/path/to/kernel_key.pem"
MODULES_SIGN_CERT="/path/to/kernel_key.pem" # Only required if the MODULES_SIGN_KEY does not also contain the certificate.
MODULES_SIGN_HASH="sha512" # Defaults to sha512.
</div>

<div lang="en" dir="ltr" class="mw-content-ltr">
# Optionally, to boot with secureboot enabled, may be the same or different signing key.
SECUREBOOT_SIGN_KEY="/path/to/kernel_key.pem"
SECUREBOOT_SIGN_CERT="/path/to/kernel_key.pem"

```

**Заметка**

The SECUREBOOT\_SIGN\_KEY and SECUREBOOT\_SIGN\_CERT may be different files. For this example the pem file generated by OpenSSL includes both the key and the accompanying certificate, and thus both variables are set to the same value.

**Заметка**

For this example the same key that was generated to sign the modules is used to sign the kernel image. It is also possible to generate and use a second separate key for signing the kernel image. The same OpenSSL command as in the previous section may be used again.

See the above section for instructions on generating a new key, the steps may be repeated if a separate key should be used to sign the kernel image.

To successfully boot with Secure Boot enabled, the used bootloader must also be signed and the certificate must be accepted by the [UEFI](/wiki/UEFI "UEFI") firmware or [Shim](/wiki/Shim "Shim"). This will be explained later in the handbook.

#### Установка distribution-ядра

Чтобы собрать ядро из исходного кода с патчами Gentoo, введите:

`root #` `emerge --ask sys-kernel/gentoo-kernel`

Администраторы систем, которые хотят избежать сборки ядра из исходных текстов на компьютере, могут вместо этого использовать предварительно скомпилированные образы ядра:

`root #` `emerge --ask sys-kernel/gentoo-kernel-bin`

**Важно**

_[Distribution Kernels](/wiki/Project:Distribution_Kernel "Project:Distribution Kernel")_, such as [sys-kernel/gentoo-kernel](https://packages.gentoo.org/packages/sys-kernel/gentoo-kernel) and [sys-kernel/gentoo-kernel-bin](https://packages.gentoo.org/packages/sys-kernel/gentoo-kernel-bin), by default, expect to be installed alongside an [initramfs](/wiki/Handbook:AMD64/Installation/Kernel#Initramfs.2Fru "Handbook:AMD64/Installation/Kernel"). Before running emerge to install the kernel users should ensure that [sys-kernel/installkernel](https://packages.gentoo.org/packages/sys-kernel/installkernel) has been configured to utilize an initramfs generator (for example [Dracut](/wiki/Dracut "Dracut")) as described in the [installkernel section](/wiki/Handbook:AMD64/Installation/Kernel#Initramfs.2Fru "Handbook:AMD64/Installation/Kernel").

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

After installing the Distribution Kernel successfully, it is now time to proceed to the next section: [Configuring the system](/wiki/Handbook:AMD64/Installation/System/ru "Handbook:AMD64/Installation/System/ru").

### Альтернатива: Ручная настройка

### Установка исходного кода ядра

При установке и компиляции ядра для систем на базе amd64 Gentoo рекомендует использовать пакет [sys-kernel/gentoo-sources](https://packages.gentoo.org/packages/sys-kernel/gentoo-sources).

Выберите подходящий исходный код ядра и установите его с помощью emerge:

`root #` `emerge --ask sys-kernel/gentoo-sources`

Данная команда установит исходный код ядра Linux в /usr/src/, используя в названии версию ядра. Эта команда не установит автоматически символьную ссылку, пока вы не укажете USE-флаг [symlink](https://packages.gentoo.org/useflags/symlink) [USE flag/ru](/wiki/USE_flag/ru "USE flag/ru") для выбранного исходного кода ядра.

Обычно, символьная ссылка /usr/src/linux указывает на исходный код текущего работающего ядра. Однако, эта символьная ссылка не создаётся по умолчанию. Создать её поможет kernel модуль для eselect.

Чтобы подробнее узнать, зачем нужна эта символьная ссылка и как ею управлять, смотрите [Kernel/Upgrade](/wiki/Kernel/Upgrade/ru "Kernel/Upgrade/ru").

Для начала, просмотрите список установленных ядер (в виде исходного кода):

`root #` `eselect kernel list`

```
Available kernel symlink targets:
  [1]   linux-6.18.8-gentoo

```

Для того, чтобы создать символьную ссылку linux, используйте:

`root #` `eselect kernel set 1`

`root #` `ls -l /usr/src/linux`

```
lrwxrwxrwx    1 root   root    20 мар  3 22:44 /usr/src/linux -> linux-6.18.8-gentoo

```

Manually configuring a kernel is commonly seen as one of the most difficult procedures a system administrator has to perform. Nothing is less true - after configuring a few kernels no one remembers that it was difficult! There are two ways for a Gentoo user to manage a manual kernel system, both of which are listed below:

**Важно**

Only one selection is required in the following subsection, if unsure of which to use then go with the first listed for now. It's always possible to switch at a later date if required.

#### Процесс modprobed-db

A very easy way to manage the kernel is to first install [sys-kernel/gentoo-kernel-bin](https://packages.gentoo.org/packages/sys-kernel/gentoo-kernel-bin) and use the [sys-kernel/modprobed-db](https://packages.gentoo.org/packages/sys-kernel/modprobed-db) to collect information about what the system requires. modprobed-db is a tool which monitors the system via crontab to add all modules of all devices over the system's life to make sure it everything a user needs is supported. For example, if an Xbox controller is added after installation, then modprobed-db will add the modules to be built next time the kernel is rebuilt.

For now please follow installing [sys-kernel/gentoo-kernel-bin](https://packages.gentoo.org/packages/sys-kernel/gentoo-kernel-bin) via [Distribution\_kernels](/wiki/Handbook:AMD64/Installation/Kernel#Distribution_kernels.2Fru "Handbook:AMD64/Installation/Kernel").

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

Now it's time to decide if modules signing is required in the steps listed in [here](/wiki/Handbook:AMD64/Installation/Kernel#Optional:_Signed_kernel_modules_2.2Fru "Handbook:AMD64/Installation/Kernel")

If not, proceed to building described [here](/wiki/Handbook:AMD64/Installation/Kernel#Compiling_and_installing.2Fru "Handbook:AMD64/Installation/Kernel")

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

#### Optional: Signing the kernel image (Secure Boot)

When signing the kernel image (for use on systems with [Secure Boot](/wiki/Secure_Boot "Secure Boot") enabled) it is recommended to set the following kernel config options:

ЯДРО **Lockdown for secureboot**

```
General setup  --->
  Kexec and crash features  --->
    [*] Enable kexec system call
    [*] Enable kexec file based system call
    [*]   Verify kernel signature during kexec_file_load() syscall
    [*]     Require a valid signature in kexec_file_load() syscall
    [*]     Enable ""image"" signature verification support
</div>

[*] Enable loadable module support
  -*-   Module signature verification
    [*]     Require modules to be validly signed
    [*]     Automatically sign all modules
    Which hash algorithm should modules be signed with? (Sign modules with SHA-512) --->

Security options  --->
[*] Integrity subsystem
  [*] Basic module for enforcing kernel lockdown
  [*]   Enable lockdown LSM early in init
        Kernel default lockdown mode (Integrity)  --->

  [*]   Digital signature verification using multiple keyrings
  [*]     Enable asymmetric keys support
  -*-       Require all keys on the integrity keyrings be signed
  [*]       Provide keyring for platform/firmware trusted keys
  [*]       Provide a keyring to which Machine Owner Keys may be added
  [ ]         Enforce Machine Keyring CA Restrictions

```

Where ""image"" is a placeholder for the architecture specific image name. These options, from the top to the bottom: enforces that the kernel image in a kexec call must be signed (kexec allows replacing the kernel in-place), enforces that kernel modules are signed, enables lockdown integrity mode (prevents modifying the kernel at runtime), and enables various keychains.

On arches that do not natively support decompressing the kernel (e.g. **arm64** and **riscv**), the kernel must be built with its own decompressor (zboot):

ЯДРО **zboot CONFIG\_EFI\_ZBOOT**

```
Device Drivers --->
  Firmware Drivers --->
    EFI (Extensible Firmware Interface) Support --->
      [*] Enable the generic EFI decompressor

```

After compilation of the kernel, as explained in the next section, the kernel image must be signed. First install [app-crypt/sbsigntools](https://packages.gentoo.org/packages/app-crypt/sbsigntools) and then sign the kernel image:

`root #` `emerge --ask app-crypt/sbsigntools`

`root #` `sbsign /usr/src/linux-x.y.z/path/to/kernel-image --cert /path/to/kernel_key.pem --key /path/to/kernel_key.pem --output /usr/src/linux-x.y.z/path/to/kernel-image`

**Заметка**

For this example, the same key that was generated to sign the modules is used to sign the kernel image. It is also possible to generate and use a second separate key for signing the kernel image. The same OpenSSL command as in the previous section may be used again.

Then proceed with the installation.

To automatically sign EFI executables installed by other packages, enable the [secureboot](https://packages.gentoo.org/useflags/secureboot) [USE flag/ru](/wiki/USE_flag/ru "USE flag/ru") USE flag globally:

ФАЙЛ **`/etc/portage/make.conf`** **Включение Secure Boot**

```
USE="modules-sign secureboot"

<div lang="en" dir="ltr" class="mw-content-ltr">
# Optionally, to use custom signing keys.
MODULES_SIGN_KEY="/path/to/kernel_key.pem"
MODULES_SIGN_CERT="/path/to/kernel_key.pem" # Only required if the MODULES_SIGN_KEY does not also contain the certificate.
MODULES_SIGN_HASH="sha512" # Defaults to sha512
</div>

<div lang="en" dir="ltr" class="mw-content-ltr">
# Optionally, to boot with secureboot enabled, may be the same or different signing key.
SECUREBOOT_SIGN_KEY="/path/to/kernel_key.pem"
SECUREBOOT_SIGN_CERT="/path/to/kernel_key.pem"

```

**Заметка**

SECUREBOOT\_SIGN\_KEY and SECUREBOOT\_SIGN\_CERT may point to different files. For this example, the pem file generated by OpenSSL includes both the key and the accompanying certificate, and thus both variables are set to the same value.

**Заметка**

When generating an [Unified Kernel Image](/wiki/Unified_Kernel_Image "Unified Kernel Image") with systemd's `ukify` the kernel image will be signed automatically before inclusion in the unified kernel image and it is not necessary to sign it manually.

#### Компиляция и установка

Когда настройка закончена, настало время скомпилировать и установить ядро. Выйдите из настройки и запустите процесс компиляции:

`root #` `make -j$(nproc) && make modules_install`

**Заметка**

Можно уменьшить количество параллельных сборок, используя make -jX, где `X` — это число параллельных задач, которые может запустить процесс сборки. Это похоже на инструкции, которые были даны ранее относительно файла /etc/portage/make.conf в части переменной MAKEOPTS.

По завершении компиляции, скопируйте образ ядра в каталог /boot/. Это делается командой make install:

`root #` `make install`

Данная команда скопирует образ ядра в каталог /boot/. Если установлен пакет [sys-kernel/installkernel](https://packages.gentoo.org/packages/sys-kernel/installkernel), то он вместо этого вызовет /sbin/installkernel, которому предоставит установку ядра. Вместо простого копирования ядра в /boot [Installkernel](/wiki/Installkernel "Installkernel") устанавливает каждое ядро, содержащее версию ядра в своём имени файла. Дополнительно installkernel реализует инфраструктуру автоматического выполнения различных задач, связанных с установкой ядра, например, создание [initramfs](/wiki/Initramfs "Initramfs"), сборку [Unified Kernel Image](/wiki/Unified_Kernel_Image "Unified Kernel Image") и обновление конфигурации [начального загрузчика](/wiki/Bootloader/ru "Bootloader/ru").

Продолжите установку с раздела [Настройка системы](/wiki/Handbook:AMD64/Installation/System/ru "Handbook:AMD64/Installation/System/ru").

[← Установка базовой системы Gentoo](/wiki/Handbook:AMD64/Installation/Base/ru "Предыдущая часть") [К содержанию](/wiki/Handbook:AMD64/ru "Handbook:AMD64/ru") [Настройка системы →](/wiki/Handbook:AMD64/Installation/System/ru "Следующая часть")