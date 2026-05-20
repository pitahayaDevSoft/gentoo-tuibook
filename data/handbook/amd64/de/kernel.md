# Kernel

Other languages:

- Deutsch
- [English](/wiki/Handbook:AMD64/Installation/Kernel "Handbook:AMD64/Installation/Kernel (100% translated)")
- [español](/wiki/Handbook:AMD64/Installation/Kernel/es "Handbook:AMD64/Installation/Kernel/es (100% translated)")
- [français](/wiki/Handbook:AMD64/Installation/Kernel/fr "Handbook:AMD64/Installation/Kernel/fr (100% translated)")
- [italiano](/wiki/Handbook:AMD64/Installation/Kernel/it "Handbook:AMD64/Installation/Kernel/it (100% translated)")
- [magyar](/wiki/Handbook:AMD64/Installation/Kernel/hu "Handbook:AMD64/Installation/Kernel/hu (100% translated)")
- [polski](/wiki/Handbook:AMD64/Installation/Kernel/pl "Handbook:AMD64/Installation/Kernel/pl (100% translated)")
- [português do Brasil](/wiki/Handbook:AMD64/Installation/Kernel/pt-br "Handbook:AMD64/Installation/Kernel/pt-br (100% translated)")
- [čeština](/wiki/Handbook:AMD64/Installation/Kernel/cs "Handbook:AMD64/Installation/Kernel/cs (100% translated)")
- [русский](/wiki/Handbook:AMD64/Installation/Kernel/ru "Handbook:AMD64/Installation/Kernel/ru (100% translated)")
- [தமிழ்](/wiki/Handbook:AMD64/Installation/Kernel/ta "Handbook:AMD64/Installation/Kernel/ta (100% translated)")
- [中文（中国大陆）‎](/wiki/Handbook:AMD64/Installation/Kernel/zh-cn "Handbook:AMD64/Installation/Kernel/zh-cn (100% translated)")
- [日本語](/wiki/Handbook:AMD64/Installation/Kernel/ja "Handbook:AMD64/Installation/Kernel/ja (100% translated)")
- [한국어](/wiki/Handbook:AMD64/Installation/Kernel/ko "Handbook:AMD64/Installation/Kernel/ko (100% translated)")

[AMD64 Handbuch](/wiki/Handbook:AMD64/de "Handbook:AMD64/de")[Installation](/wiki/Handbook:AMD64/Full/Installation/de "Handbook:AMD64/Full/Installation/de")[Über die Installation](/wiki/Handbook:AMD64/Installation/About/de "Handbook:AMD64/Installation/About/de")[Auswahl des Mediums](/wiki/Handbook:AMD64/Installation/Media/de "Handbook:AMD64/Installation/Media/de")[Konfiguration des Netzwerks](/wiki/Handbook:AMD64/Installation/Networking/de "Handbook:AMD64/Installation/Networking/de")[Vorbereiten der Festplatte(n)](/wiki/Handbook:AMD64/Installation/Disks/de "Handbook:AMD64/Installation/Disks/de")[Installation des Stage Archivs](/wiki/Handbook:AMD64/Installation/Stage/de "Handbook:AMD64/Installation/Stage/de")[Installation des Basissystems](/wiki/Handbook:AMD64/Installation/Base/de "Handbook:AMD64/Installation/Base/de")Konfiguration des Kernels[Konfiguration des Systems](/wiki/Handbook:AMD64/Installation/System/de "Handbook:AMD64/Installation/System/de")[Installation der Tools](/wiki/Handbook:AMD64/Installation/Tools/de "Handbook:AMD64/Installation/Tools/de")[Konfiguration des Bootloaders](/wiki/Handbook:AMD64/Installation/Bootloader/de "Handbook:AMD64/Installation/Bootloader/de")[Abschluss](/wiki/Handbook:AMD64/Installation/Finalizing/de "Handbook:AMD64/Installation/Finalizing/de")[Arbeiten mit Gentoo](/wiki/Handbook:AMD64/Full/Working/de "Handbook:AMD64/Full/Working/de")[Portage-Einführung](/wiki/Handbook:AMD64/Working/Portage/de "Handbook:AMD64/Working/Portage/de")[USE-Flags](/wiki/Handbook:AMD64/Working/USE/de "Handbook:AMD64/Working/USE/de")[Portage-Features](/wiki/Handbook:AMD64/Working/Features/de "Handbook:AMD64/Working/Features/de")[Initskript-System](/wiki/Handbook:AMD64/Working/Initscripts/de "Handbook:AMD64/Working/Initscripts/de")[Umgebungsvariablen](/wiki/Handbook:AMD64/Working/EnvVar/de "Handbook:AMD64/Working/EnvVar/de")[Arbeiten mit Portage](/wiki/Handbook:AMD64/Full/Portage/de "Handbook:AMD64/Full/Portage/de")[Dateien und Verzeichnisse](/wiki/Handbook:AMD64/Portage/Files/de "Handbook:AMD64/Portage/Files/de")[Variablen](/wiki/Handbook:AMD64/Portage/Variables/de "Handbook:AMD64/Portage/Variables/de")[Mischen von Softwarezweigen](/wiki/Handbook:AMD64/Portage/Branches/de "Handbook:AMD64/Portage/Branches/de")[Zusätzliche Tools](/wiki/Handbook:AMD64/Portage/Tools/de "Handbook:AMD64/Portage/Tools/de")[Eigener Portage-Tree](/wiki/Handbook:AMD64/Portage/CustomTree/de "Handbook:AMD64/Portage/CustomTree/de")[Erweiterte Portage-Features](/wiki/Handbook:AMD64/Portage/Advanced/de "Handbook:AMD64/Portage/Advanced/de")[Netzwerk-Konfiguration](/wiki/Handbook:AMD64/Full/Networking/de "Handbook:AMD64/Full/Networking/de")[Zu Beginn](/wiki/Handbook:AMD64/Networking/Introduction/de "Handbook:AMD64/Networking/Introduction/de")[Fortgeschrittene Konfiguration](/wiki/Handbook:AMD64/Networking/Advanced/de "Handbook:AMD64/Networking/Advanced/de")[Modulare Vernetzung](/wiki/Handbook:AMD64/Networking/Modular/de "Handbook:AMD64/Networking/Modular/de")[Drahtlose Netzwerke](/wiki/Handbook:AMD64/Networking/Wireless/de "Handbook:AMD64/Networking/Wireless/de")[Funktionalität hinzufügen](/wiki/Handbook:AMD64/Networking/Extending/de "Handbook:AMD64/Networking/Extending/de")[Dynamisches Management](/wiki/Handbook:AMD64/Networking/Dynamic/de "Handbook:AMD64/Networking/Dynamic/de")

## Contents

- [1Optional: Firmware und/oder Microcode installieren](#Optional:_Firmware_und.2Foder_Microcode_installieren)
- [2Firmware](#Firmware)
  - [2.1Suggested: Linux Firmware](#Suggested:_Linux_Firmware)
    - [2.1.1Firmware Loading](#Firmware_Loading)
  - [2.2SOF Firmware](#SOF_Firmware)
- [3Microcode](#Microcode)
- [4sys-kernel/installkernel](#sys-kernel.2Finstallkernel)
  - [4.1Bootloader](#Bootloader)
    - [4.1.1GRUB](#GRUB)
    - [4.1.2systemd-boot](#systemd-boot)
    - [4.1.3EFI stub](#EFI_stub)
    - [4.1.4Traditional layout, other bootloaders (e.g. (e)lilo, syslinux, etc.)](#Traditional_layout.2C_other_bootloaders_.28e.g._.28e.29lilo.2C_syslinux.2C_etc..29)
  - [4.2Initramfs](#Initramfs)
    - [4.2.1Chroot detection](#Chroot_detection)
  - [4.3Optional: Unified Kernel Image](#Optional:_Unified_Kernel_Image)
    - [4.3.1Generic Unified Kernel Image (systemd only)](#Generic_Unified_Kernel_Image_.28systemd_only.29)
    - [4.3.2Secure Boot](#Secure_Boot)
- [5Kernel konfigurieren und Kompilieren](#Kernel_konfigurieren_und_Kompilieren)
  - [5.1Distribution kernels](#Distribution_kernels)
    - [5.1.1Optional: Signed kernel modules](#Optional:_Signed_kernel_modules)
    - [5.1.2Optional: Signing the kernel image (Secure Boot)](#Optional:_Signing_the_kernel_image_.28Secure_Boot.29)
    - [5.1.3Installing a distribution kernel](#Installing_a_distribution_kernel)
    - [5.1.4Upgrading and cleaning up](#Upgrading_and_cleaning_up)
    - [5.1.5Post-install/upgrade tasks](#Post-install.2Fupgrade_tasks)
      - [5.1.5.1Manually rebuilding the initramfs or Unified Kernel Image](#Manually_rebuilding_the_initramfs_or_Unified_Kernel_Image)
  - [5.2Alternative: Manuelle Konfiguration](#Alternative:_Manuelle_Konfiguration)
  - [5.3Installieren der Kernel-Quellen](#Installieren_der_Kernel-Quellen)
    - [5.3.1Option 1 - Modprobed-db process](#Option_1_-_Modprobed-db_process)
    - [5.3.2Option 2 - Assisted manual process](#Option_2_-_Assisted_manual_process)
    - [5.3.3Option 3 - Configuring by hand](#Option_3_-_Configuring_by_hand)
    - [5.3.4Optional: Signed kernel modules](#Optional:_Signed_kernel_modules_2)
    - [5.3.5Optional: Signing the kernel image (Secure Boot)](#Optional:_Signing_the_kernel_image_.28Secure_Boot.29_2)
    - [5.3.6Kompilieren und installieren](#Kompilieren_und_installieren)

### Optional: Firmware und/oder Microcode installieren

### Firmware

#### Suggested: Linux Firmware

On many systems, non-FOSS firmware is required for certain hardware to function. The [sys-kernel/linux-firmware](https://packages.gentoo.org/packages/sys-kernel/linux-firmware) package contains firmware for many, but not all, devices.

**Tipp**

Most wireless cards and GPUs require firmware to function.

`root #` `emerge --ask sys-kernel/linux-firmware`

**Hinweis**

Die Installation bestimmter Firmware-Pakete erfordert oft das Akzeptieren der zugehörigen Firmware-Lizenzen. Falls erforderlich, finden Sie im [Lizenzhandhabungsabschnitt](/wiki/Handbook:AMD64/Working/Portage/de#Licenses "Handbook:AMD64/Working/Portage/de") des Handbuchs Hilfe zum Akzeptieren von Lizenzen.

##### Firmware Loading

Firmware files are typically loaded when the associated kernel module is loaded. This means the firmware must be built into the kernel using **CONFIG\_EXTRA\_FIRMWARE** if the kernel module is set to _Y_ instead of _M_. In most cases, building-in a module which required firmware can complicate or break loading.

#### SOF Firmware

Sound Open Firmware (SOF) is a new open source audio driver meant to replace the proprietary Smart Sound Technology (SST) audio driver from Intel. 10th gen+ and Apollo Lake (Atom E3900, Celeron N3350, and Pentium N4200) Intel CPUs require this firmware for certain features and certain AMD APUs also have support for this firmware. SOF's supported platforms matrix can be found [here](https://thesofproject.github.io/latest/platforms/index.html) for more information.

`root #` `emerge --ask sys-firmware/sof-firmware`

{{#switch: amd64 \| amd64 \| x86 =

### Microcode

Neben Grafik- und Netzwerk-Hardware können auch CPUs Firmware-Updates benötigen. Dieser Typ von Firmware wird _Microcode_ genannt. Manchmal sind aktuelle Versionen von Microcode erforderlich, um Stabilitätsprobleme, Sicherheitsprobleme oder andere Bugs in CPU-Hardware zu patchen.

Microcode Updates für AMD CPUs sind in dem bereits genannten Paket [sys-kernel/linux-firmware](https://packages.gentoo.org/packages/sys-kernel/linux-firmware) enthalten. Microcode Updates für Intel CPUs sind in dem Paket [sys-firmware/intel-microcode](https://packages.gentoo.org/packages/sys-firmware/intel-microcode) enthalten. Dieses muss zusätzlich installiert werden. Weitere Informationen finden Sie im [Microcode Artikel](/wiki/Microcode "Microcode"), in dem auch beschrieben wird, wie Microcode Updates auf der CPU aktiviert werden.

## sys-kernel/installkernel

[Installkernel](/wiki/Installkernel "Installkernel") may be used to automate the kernel installation, [initramfs](/wiki/Initramfs "Initramfs") generation, [unified kernel image](/wiki/Unified_kernel_image "Unified kernel image") generation and bootloader configuration, among other things. [sys-kernel/installkernel](https://packages.gentoo.org/packages/sys-kernel/installkernel) implements two paths of achieving this: the traditional installkernel originating from Debian and [systemd](/wiki/Systemd/de "Systemd/de")'s kernel-install. Which one to choose depends, among other things, on the system's bootloader. By default, systemd's kernel-install is used on systemd profiles, while the traditional installkernel is the default for other profiles.

### Bootloader

Now is the time to think about which bootloader the user wants for the system.

**Wichtig**

Only one selection is required in the following subsection, if unsure of which to use then go with the first listed for now. It's always possible to switch at a later date if required.

#### GRUB

Users of GRUB can use either systemd's kernel-install or the traditional Debian installkernel. The [systemd](https://packages.gentoo.org/useflags/systemd) [USE flag/de](/wiki/USE_flag/de "USE flag/de") USE flag switches between these implementations. To automatically run grub-mkconfig when installing the kernel, enable the [grub](https://packages.gentoo.org/useflags/grub) [USE flag/de](/wiki/USE_flag/de "USE flag/de") [USE flag](/wiki/USE_flag/de "USE flag/de").

**Hinweis**

GRUB requires kernels to be installed to /boot.

DATEI **`/etc/portage/package.use/installkernel`**

```
sys-kernel/installkernel grub

```

`root #` `emerge --ask sys-kernel/installkernel`

#### systemd-boot

When using [systemd-boot](/wiki/Systemd/systemd-boot "Systemd/systemd-boot") (formerly gummiboot) as the bootloader, systemd's kernel-install must be used. Therefore ensure the [systemd](https://packages.gentoo.org/useflags/systemd) [USE flag/de](/wiki/USE_flag/de "USE flag/de") and the [systemd-boot](https://packages.gentoo.org/useflags/systemd-boot) [USE flag/de](/wiki/USE_flag/de "USE flag/de") USE flags are enabled on [sys-kernel/installkernel](https://packages.gentoo.org/packages/sys-kernel/installkernel), and then install the relevant package for systemd-boot.

DATEI **`/etc/portage/package.use/systemd-boot`**

```
sys-apps/systemd boot
sys-apps/systemd-utils boot kernel-install
sys-kernel/installkernel systemd systemd-boot

```

The kernel command line to use for new kernels should be specified in /etc/kernel/cmdline and /etc/cmdline. As they both need to be updated together, it make sense to symlink the two files now.

DATEI **`/etc/kernel/cmdline`**

```
quiet splash

```

`root #` `mkdir -p /etc/cmdline.d`

`root #` `ln -s /etc/kernel/cmdline /etc/cmdline.d/00-installkernel.conf`

`root #` `emerge --ask sys-kernel/installkernel`

**Hinweis**

systemd-boot requires kernels to be installed to /efi.

#### EFI stub

UEFI-based computer systems technically do not need secondary bootloaders in order to boot kernels. Secondary bootloaders exist to _extend_ the functionality of UEFI firmware during the boot process. That being said, using a secondary bootloader is typically easier and more robust because it offers a more flexible approach for quickly modifying kernel parameters at boot time. Note also that UEFI implementations strongly differ between vendors and between models and there is no guarantee that a given firmware follows the UEFI specification. Therefore, EFI Stub booting is not guaranteed to work on every UEFI-based system.

`root #` `emerge --ask sys-kernel/installkernel`

`root #` `mkdir -p /efi/EFI/Gentoo`

**Hinweis**

When [app-emulation/virt-firmware](https://packages.gentoo.org/packages/app-emulation/virt-firmware) is used to configure the UEFI ensure that the kernel-bootcfg-boot-successful service is enabled before attempting to install the kernel. This implementation of EFIstub booting is the default for systemd systems.

`root #` `systemctl enable kernel-bootcfg-boot-successful.service`

**Hinweis**

EFIstub requires kernels to be installed to /efi.

#### Traditional layout, other bootloaders (e.g. (e)lilo, syslinux, etc.)

The traditional /boot layout (for e.g. (e)LILO, syslinux, etc.) is used by default if the [grub](https://packages.gentoo.org/useflags/grub) [USE flag/de](/wiki/USE_flag/de "USE flag/de"), [systemd-boot](https://packages.gentoo.org/useflags/systemd-boot) [USE flag/de](/wiki/USE_flag/de "USE flag/de"), [efistub](https://packages.gentoo.org/useflags/efistub) [USE flag/de](/wiki/USE_flag/de "USE flag/de") and [uki](https://packages.gentoo.org/useflags/uki) [USE flag/de](/wiki/USE_flag/de "USE flag/de") USE flags are **not** enabled. No further action is required.

### Initramfs

An **init** ial **ram**-based **f** ile **s** ystem, or [initramfs](/wiki/Initramfs "Initramfs"), may be required for a system to boot. A wide of variety of cases may necessitate one, but common cases include:

- Kernels where storage/filesystem drivers are modules.
- Layouts with /usr/ or /var/ on separate partitions.
- Encrypted root filesystems.

**Tipp**

[Distribution kernels](/wiki/Project:Distribution_Kernel "Project:Distribution Kernel") are designed to be used with an initramfs, as many storage and filesystem drivers are built as modules.

In addition to mounting the root filesystem, an initramfs may also perform other tasks such as:

- Running **f** ile **s** ystem **c** onsistency chec **k** fsck, a tool to check and repair consistency of a file system in such events of uncleanly shutdown a system.
- Providing a recovery environment in the event of late-boot failures.

[Installkernel](/wiki/Installkernel "Installkernel") can automatically generate an initramfs when installing the kernel if the [dracut](https://packages.gentoo.org/useflags/dracut) [USE flag/de](/wiki/USE_flag/de "USE flag/de") or [ugrd](https://packages.gentoo.org/useflags/ugrd) [USE flag/de](/wiki/USE_flag/de "USE flag/de") USE flag is enabled:

DATEI **`/etc/portage/package.use/installkernel`**

```
sys-kernel/installkernel dracut

```

#### Chroot detection

Bootloaders such as [systemd-boot](/wiki/Systemd/systemd-boot "Systemd/systemd-boot") and [EFI stub](/wiki/EFI_stub "EFI stub") use the kernel arguments of the running system as set in /proc/cmdline by default. Because of the wide range of ways Gentoo can be installed users will randomly get tripped up by this.

To help solve any issues this may cause, [sys-kernel/installkernel](https://packages.gentoo.org/packages/sys-kernel/installkernel) will check if it is running in a chroot and fail if the kernel command line is not explicitly configured. Otherwise the bootloader would use the install media's boot arguments which would lead to boot failure.

One way to satisfy [sys-kernel/installkernel](https://packages.gentoo.org/packages/sys-kernel/installkernel) is by using Dracut's config file to set the root partition UUID as shown below, or alternatively for more information on what this check helps with and different ways to configure it, see [Installkernel#Install\_chroot\_check](/wiki/Installkernel#Install_chroot_check.2Fde "Installkernel").

`root #` `mkdir /etc/dracut.conf.d`

`root #` `blkid`

```
/dev/sda3: UUID="2122cd72-94d7-4dcc-821e-3705926deecc"
```

In the above example, the root partition is /dev/sda3 and the UUID is 2122cd72-94d7-4dcc-821e-3705926deecc.
The dracut config file would then look like:

DATEI **`/etc/dracut.conf.d/00-installkernel.conf`**

```
kernel_cmdline=" root=UUID=2122cd72-94d7-4dcc-821e-3705926deecc " # Note leading and trailing spaces

```

`root #` `emerge --ask sys-kernel/installkernel`

### Optional: Unified Kernel Image

A [Unified Kernel Image](/wiki/Unified_Kernel_Image "Unified Kernel Image") (UKI) combines, among other things, the kernel, the initramfs and the kernel command line into a single executable. Since the kernel command line is embedded into the unified kernel image, it should be specified before generating the unified kernel image (see below). Note that any kernel command line arguments supplied by the bootloader or firmware at boot are ignored when booting with secure boot enabled.

A unified kernel image requires a stub loader. Currently, the only one available is systemd-stub. To enable it:

For systemd systems:

DATEI **`/etc/portage/package.use/uki`**

```
sys-apps/systemd boot

```

`root #` `emerge --ask sys-apps/systemd`

For OpenRC systems:

DATEI **`/etc/portage/package.use/uki`**

```
sys-apps/systemd-utils boot kernel-install

```

`root #` `emerge --ask sys-apps/systemd-utils`

[Installkernel](/wiki/Installkernel "Installkernel") can automatically generate a unified kernel image using either [dracut](/wiki/Unified_kernel_image#dracut.2Fde "Unified kernel image") or [ukify](/wiki/Unified_kernel_image#ukify.2Fde "Unified kernel image") by enabling the respective flag and the [uki](https://packages.gentoo.org/useflags/uki) [USE flag/de](/wiki/USE_flag/de "USE flag/de") USE flag.

For dracut:

DATEI **`/etc/portage/package.use/uki`**

```
sys-kernel/installkernel dracut uki

```

DATEI **`/etc/dracut.conf.d/uki.conf`**

```
uefi="yes"
kernel_cmdline="some-kernel-command-line-arguments"

```

`root #` `emerge --ask sys-kernel/installkernel`

For ukify:

DATEI **`/etc/portage/package.use/uki`**

```
sys-apps/systemd boot ukify                         # For systemd systems
sys-apps/systemd-utils kernel-install boot ukify    # For OpenRC systems
sys-kernel/installkernel dracut ukify uki

```

DATEI **`/etc/kernel/cmdline`**

```
some-kernel-command-line-arguments

```

`root #` `emerge --ask sys-kernel/installkernel`

Note that while dracut can generate both an initramfs and a unified kernel image, ukify can only generate the latter and therefore the initramfs must be generated separately with dracut.

**Wichtig**

In the above configuration examples (for both Dracut and ukify) it is important to specify at least an appropriate root= parameter for the kernel command line to ensure that the Unified Kernel Image can find the root partition. This is not required for systemd based systems following the Discoverable Partitions Specification (DPS), in that case the embedded initramfs will be able to dynamically find the root partition.

#### Generic Unified Kernel Image (systemd only)

The prebuilt [sys-kernel/gentoo-kernel-bin](https://packages.gentoo.org/packages/sys-kernel/gentoo-kernel-bin) can optionally install a prebuilt generic unified kernel image containing a generic initramfs that is able to boot most systemd based systems. It can be installed by enabling the [generic-uki](https://packages.gentoo.org/useflags/generic-uki) [USE flag/de](/wiki/USE_flag/de "USE flag/de") USE flag, and configuring [installkernel](/wiki/Installkernel "Installkernel") to not generate a custom initramfs or unified kernel image:

DATEI **`/etc/portage/package.use/uki`**

```
sys-kernel/gentoo-kernel-bin generic-uki
sys-kernel/installkernel -dracut -ukify -ugrd uki

```

#### Secure Boot

**Warnung**

If following this section and manually compiling your own kernel, then make sure to follow the steps outlined in [Signing the kernel](/wiki/Kernel#Optional:_Signing_the_kernel_image_.28Secure_Boot.29.2Fde "Kernel")

The generic Unified Kernel Image optionally distributed by [sys-kernel/gentoo-kernel-bin](https://packages.gentoo.org/packages/sys-kernel/gentoo-kernel-bin) is already pre-signed. How to sign a locally generated unified kernel image depends on whether dracut or ukify is used. Note that the location of the key and certificate should be the same as the SECUREBOOT\_SIGN\_KEY and SECUREBOOT\_SIGN\_CERT as specified in /etc/portage/make.conf.

For dracut:

DATEI **`/etc/dracut.conf.d/uki.conf`**

```
uefi="yes"
kernel_cmdline="some-kernel-command-line-arguments"
uefi_secureboot_key="/path/to/kernel_key.pem"
uefi_secureboot_cert="/path/to/kernel_key.pem"

```

For ukify:

DATEI **`/etc/kernel/uki.conf`**

```
[UKI]
SecureBootPrivateKey=/path/to/kernel_key.pem
SecureBootCertificate=/path/to/kernel_key.pem

```

## Kernel konfigurieren und Kompilieren

It can be a wise move to use the dist-kernel on the first boot as it provides a very simple method to rule out system issues and kernel config issues. Always having a known working kernel to fallback on can speed up debugging and alleviate anxiety when updating that your system will no longer boot.

A common misconception is that a manually compiled kernel will use a lot less RAM than a pre configured distribution kernel. Due to the modular nature of the Linux kernel, only what is needed by the system is loaded and in most cases resulting in similar memory usage.

**Wichtig**

Only one selection is required in the following subsection, if unsure of which to use then go with the first listed for now. It's always possible to switch at a later date if required.

In der Reihenfolge vom geringsten bis zum größten Aufwand:

[Full automation approach: Distribution kernels](/wiki/Handbook:AMD64/Installation/Kernel/de#Distribution_kernels "Handbook:AMD64/Installation/Kernel/de")A [Distribution Kernel](/wiki/Project:Distribution_Kernel "Project:Distribution Kernel") is used to configure, automatically build, and install the Linux kernel, its associated modules, and (optionally, but enabled by default) an initramfs file. Future kernel updates are fully automated since they are handled through the package manager, just like any other system package. It is possible [provide a custom kernel configuration file](/wiki/Project:Distribution_Kernel#Modifying_kernel_configuration "Project:Distribution Kernel") if customization is necessary. This is the least involved process and is perfect for new Gentoo users due to it working out-of-the-box and offering minimal involvement from the system administrator.[Hybrid approach: Genkernel](/wiki/Handbook:AMD64/Installation/Kernel/de#Alternative:_Genkernel "Handbook:AMD64/Installation/Kernel/de")New kernel sources are installed via the system package manager. System administrators use Gentoo's genkernel tool to generically configure, automatically build and install the Linux kernel, its associated modules, and (optionally, but _**not**_ enabled by default) an initramfs file. It is possible provide a custom kernel configuration file if customization is necessary. Future kernel configuration, compilation, and installation require the system administrator's involvement in the form of running eselect kernel, genkernel, and potentially other commands for each update.[Full manual approach](/wiki/Handbook:AMD64/Installation/Kernel/de#Alternative:_Manual_configuration "Handbook:AMD64/Installation/Kernel/de")New kernel sources are installed via the system package manager. The kernel is manually configured, built, and installed using the eselect kernel and a slew of make commands. Future kernel updates repeat the manual process of configuring, building, and installing the kernel files. This is the most involved process, but offers maximum control over the kernel update process.[Vollständiger Automatisierungsansatz: Distributionskernel](/wiki/Handbook:AMD64/Installation/Kernel/de#Distribution_kernels "Handbook:AMD64/Installation/Kernel/de")Ein [Distributionskernel](/wiki/Project:Distribution_Kernel "Project:Distribution Kernel") wird verwendet, um den Linux-Kernel, seine dazugehörigen Module und (optional, aber standardmäßig aktiviert) eine initramfs-Datei zu konfigurieren, automatisch zu bauen und zu installieren. Zukünftige Kernel-Aktualisierungen sind vollständig automatisiert, da sie über den Paketmanager abgewickelt werden, genauso wie jedes andere Systempaket. Es ist möglich [eine eigene Kernelkonfigurationsdatei bereitzustellen](/wiki/Project:Distribution_Kernel#Modifying_kernel_configuration "Project:Distribution Kernel"), wenn Anpassungen notwendig sind. Dies ist der am wenigsten aufwendige Prozess und eignet sich perfekt für neue Gentoo-Benutzer, da er sofort funktioniert und nur minimale Eingriffe des Systemadministrators erfordert.[Hybridansatz: Genkernel](/wiki/Handbook:AMD64/Installation/Kernel/de#Alternative:_Genkernel "Handbook:AMD64/Installation/Kernel/de")Neue Kernelquellen werden über den Systempaketmanager installiert. Systemadministratoren benutzen Gentoos genkernel Werkzeug, um den Linux-Kernel, seine dazugehörigen Module und (optional aber _nicht_ standardmäßig aktiviert) eine initramfs-Datei zu konfigurieren, automatisch zu bauen und zu installieren. Es ist möglich, eine benutzerdefinierte Kernelkonfigurationsdatei bereitzustellen, wenn eine Anpassung erforderlich ist. Zukünftige Kernelkonfigurationen, -kompilierungen und -installationen erfordern die Beteiligung des Systemadministrators in Form der Ausführung von eselect kernel, genkernel und möglicherweise andere Befehle für jede Aktualisierung.[Vollständige manuelle Vorgehensweise](/wiki/Handbook:AMD64/Installation/Kernel/de#Alternative:_Manual_configuration "Handbook:AMD64/Installation/Kernel/de")Neue Kernelquellen werden über den Systempaketmanager installiert. Der Kernel wird manuell konfiguriert, gebaut und installiert unter Verwendung des eselect kernel und einer Reihe von make Befehlen. Künftige Kernel-Updates wiederholen den manuellen Prozess des Konfigurierens, Erstellens und Installieren der Kernel-Dateien. Dies ist der aufwändigste Prozess, bietet aber maximale Kontrolle über den Kernel-Aktualisierungsprozess.

Der Kern, um den alle Distributionen gebaut sind, ist der Linux Kernel. Er ist die Schicht zwischen den Benutzerprogrammen und der Systemhardware. Obwohl das Handbuch seinen Benutzern verschiedene Kernel-Quellen anbietet, ist eine umfassendere Auflistung mit detaillierter Beschreibung auf der [Kernel Übersichtsseite](/wiki/Kernel/Overview "Kernel/Overview") verfügbar.

**Tipp**

Kernel installation tasks such as copying the kernel image to /boot or the [EFI System Partition](/wiki/EFI_System_Partition/de "EFI System Partition/de"), generating an [initramfs](/wiki/Initramfs "Initramfs") and/or [Unified Kernel Image](/wiki/Unified_Kernel_Image "Unified Kernel Image"), updating bootloader configuration, can be automated with [installkernel](/wiki/Installkernel "Installkernel"). Users may wish to configure and install [sys-kernel/installkernel](https://packages.gentoo.org/packages/sys-kernel/installkernel) before proceeding. See the [Kernel installation section below](/wiki/Handbook:AMD64/Installation/Kernel#Kernel_installation.2Fde "Handbook:AMD64/Installation/Kernel") for more more information.

### Distribution kernels

_[Distribution Kernels](/wiki/Project:Distribution_Kernel "Project:Distribution Kernel")_ are ebuilds that cover the complete process of unpacking, configuring, compiling, and installing the kernel. The primary advantage of this method is that the kernels are updated to new versions by the package manager as part of @world upgrade. This requires no more involvement than running an emerge command. Distribution kernels default to a configuration supporting the majority of hardware, however two mechanisms are offered for customization: savedconfig and config snippets. See the project page for [more details on configuration.](/wiki/Project:Distribution_Kernel#Modifying_kernel_configuration "Project:Distribution Kernel")

##### Optional: Signed kernel modules

The kernel modules in the prebuilt distribution kernel ([sys-kernel/gentoo-kernel-bin](https://packages.gentoo.org/packages/sys-kernel/gentoo-kernel-bin)) are already signed. To sign the modules of kernels built from source enable the [modules-sign](https://packages.gentoo.org/useflags/modules-sign) [USE flag/de](/wiki/USE_flag/de "USE flag/de") USE flag, and optionally specify which key to use for signing in [/etc/portage/make.conf](/wiki//etc/portage/make.conf/de "/etc/portage/make.conf/de"):

DATEI **`/etc/portage/make.conf`** **Enable module signing**

```
USE="modules-sign"

# Optionally, to use custom signing keys.
MODULES_SIGN_KEY="/path/to/kernel_key.pem"
MODULES_SIGN_CERT="/path/to/kernel_key.pem" # Only required if the MODULES_SIGN_KEY does not also contain the certificate.
MODULES_SIGN_HASH="sha512" # Defaults to sha512.

```

If MODULES\_SIGN\_KEY is not specified the kernel build system will generate a key, it will be stored in /usr/src/linux-x.y.z/certs. It is recommended to manually generate a key to ensure that it will be the same for each kernel release. A key may be generated with:

`root #` `openssl req -new -noenc -utf8 -sha256 -x509 -outform PEM -out kernel_key.pem -keyout kernel_key.pem`

**Hinweis**

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

The kernel image in the prebuilt distribution kernel ([sys-kernel/gentoo-kernel-bin](https://packages.gentoo.org/packages/sys-kernel/gentoo-kernel-bin)) is already signed for use with [Secure Boot](/wiki/Secure_Boot "Secure Boot"). To sign the kernel image of kernels built from source enable the [secureboot](https://packages.gentoo.org/useflags/secureboot) [USE flag/de](/wiki/USE_flag/de "USE flag/de") USE flag, and optionally specify which key to use for signing in [/etc/portage/make.conf](/wiki//etc/portage/make.conf/de "/etc/portage/make.conf/de"). Note that signing the kernel image for use with secureboot requires that the kernel modules are also signed, the same key may be used to sign both the kernel image and the kernel modules:

DATEI **`/etc/portage/make.conf`** **Enable custom signing keys**

```
USE="modules-sign secureboot"

# Optionally, to use custom signing keys.
MODULES_SIGN_KEY="/path/to/kernel_key.pem"
MODULES_SIGN_CERT="/path/to/kernel_key.pem" # Only required if the MODULES_SIGN_KEY does not also contain the certificate.
MODULES_SIGN_HASH="sha512" # Defaults to sha512.

# Optionally, to boot with secureboot enabled, may be the same or different signing key.
SECUREBOOT_SIGN_KEY="/path/to/kernel_key.pem"
SECUREBOOT_SIGN_CERT="/path/to/kernel_key.pem"

```

**Hinweis**

The SECUREBOOT\_SIGN\_KEY and SECUREBOOT\_SIGN\_CERT may be different files. For this example the pem file generated by OpenSSL includes both the key and the accompanying certificate, and thus both variables are set to the same value.

**Hinweis**

For this example the same key that was generated to sign the modules is used to sign the kernel image. It is also possible to generate and use a second separate key for signing the kernel image. The same OpenSSL command as in the previous section may be used again.

See the above section for instructions on generating a new key, the steps may be repeated if a separate key should be used to sign the kernel image.

To successfully boot with Secure Boot enabled, the used bootloader must also be signed and the certificate must be accepted by the [UEFI](/wiki/UEFI "UEFI") firmware or [Shim](/wiki/Shim "Shim"). This will be explained later in the handbook.

#### Installing a distribution kernel

To build a kernel with Gentoo patches from source, type:

`root #` `emerge --ask sys-kernel/gentoo-kernel`

System administrators who want to avoid compiling the kernel sources locally can instead use precompiled kernel images:

`root #` `emerge --ask sys-kernel/gentoo-kernel-bin`

**Wichtig**

_[Distribution Kernels](/wiki/Project:Distribution_Kernel "Project:Distribution Kernel")_, such as [sys-kernel/gentoo-kernel](https://packages.gentoo.org/packages/sys-kernel/gentoo-kernel) and [sys-kernel/gentoo-kernel-bin](https://packages.gentoo.org/packages/sys-kernel/gentoo-kernel-bin), by default, expect to be installed alongside an [initramfs](/wiki/Handbook:AMD64/Installation/Kernel#Initramfs.2Fde "Handbook:AMD64/Installation/Kernel"). Before running emerge to install the kernel users should ensure that [sys-kernel/installkernel](https://packages.gentoo.org/packages/sys-kernel/installkernel) has been configured to utilize an initramfs generator (for example [Dracut](/wiki/Dracut "Dracut")) as described in the [installkernel section](/wiki/Handbook:AMD64/Installation/Kernel#Initramfs.2Fde "Handbook:AMD64/Installation/Kernel").

#### Upgrading and cleaning up

Once the kernel is installed, the package manager will automatically update it to newer versions. The previous versions will be kept until the package manager is requested to clean up stale packages. To reclaim disk space, stale packages can be trimmed by periodically running emerge with the `--depclean` option:

`root #` `emerge --depclean`

Alternatively, to specifically clean up old kernel versions:

`root #` `emerge --prune sys-kernel/gentoo-kernel sys-kernel/gentoo-kernel-bin`

**Tipp**

By design, emerge only removes the kernel build directory. It does not actually remove the kernel modules, nor the installed kernel image. To completely clean-up old kernels, the [app-admin/eclean-kernel](https://packages.gentoo.org/packages/app-admin/eclean-kernel) tool may be used.

#### Post-install/upgrade tasks

An upgrade of a distribution kernel is capable of triggering an automatic rebuild for external kernel modules installed by other packages (for example: [sys-fs/zfs-kmod](https://packages.gentoo.org/packages/sys-fs/zfs-kmod) or [x11-drivers/nvidia-drivers](https://packages.gentoo.org/packages/x11-drivers/nvidia-drivers)). This automated behaviour is enabled by enabling the [dist-kernel](https://packages.gentoo.org/useflags/dist-kernel) [USE flag/de](/wiki/USE_flag/de "USE flag/de") USE flag. When required, this same flag will also trigger re-generation of the [initramfs](/wiki/Initramfs "Initramfs").

It is highly recommended to enable this flag globally via /etc/portage/make.conf when using a distribution kernel:

DATEI **`/etc/portage/make.conf`** **Enabling USE=dist-kernel**

```
USE="dist-kernel"

```

##### Manually rebuilding the initramfs or Unified Kernel Image

If required, manually trigger such rebuilds by, after a kernel upgrade, executing:

`root #` `emerge --ask @module-rebuild`

If any kernel modules (e.g. ZFS) are needed at early boot, rebuild the initramfs afterward via:

`root #` `emerge --config sys-kernel/gentoo-kernel
`

`root #` `emerge --config sys-kernel/gentoo-kernel-bin
`

After installing the Distribution Kernel successfully, it is now time to proceed to the next section: [Configuring the system](/wiki/Handbook:AMD64/Installation/System/de "Handbook:AMD64/Installation/System/de").

### Alternative: Manuelle Konfiguration

### Installieren der Kernel-Quellen

Für die Installation und Kompilierung des Kernels für amd64-basierte Systeme empfiehlt Gentoo das Paket [sys-kernel/gentoo-sources](https://packages.gentoo.org/packages/sys-kernel/gentoo-sources).

Wählen Sie eine geeignete Kernelquelle und installieren Sie sie mit emerge:

`root #` `emerge --ask sys-kernel/gentoo-sources`

Dieser Befehl installiert die Quellen des Linux-Kernels in dem Verzeichnis /usr/src/. Im Namen des angelegten Unterverzeichnisses steht die Version des Linux-Kernels (beispielsweise /usr/src/linux-4.9.16-gentoo). Der emerge-Befehl erzeugt einen symbolischen Link /usr/src/linux auf dieses Verzeichnis, wenn das USE-Flag `symlink` bei dem Paket gesetzt ist.

Es ist gute Praxis, einen Symlink /usr/src/linux auf das Verzeichnis mit dem Source-Code der Kernel-Version zeigen zu lassen, die auf dem System läuft. Dieser Symlink wird nicht automatisch erzeugt. Er kann jedoch einfach mit Hilfe des eselect Kernel Modules erzeugt werden.

Weitere Informationen über Zweck und Handhabung des Symlinks finden Sie unter [Kernel/Upgrade](/wiki/Kernel/Upgrade/de "Kernel/Upgrade/de").

Lassen Sie eine Liste der installierten Kernel anzeigen:

`root #` `eselect kernel list`

```
Available kernel symlink targets:
  [1]   linux-6.18.8-gentoo

```

Mit folgendem Befehl können Sie einen symbolischen Link namens linux anlegen:

`root #` `eselect kernel set 1`

`root #` `ls -l /usr/src/linux`

```
lrwxrwxrwx    1 root   root    12 Oct 13 11:04 /usr/src/linux -> linux-6.18.8-gentoo

```

Manually configuring a kernel is commonly seen as one of the most difficult procedures a system administrator has to perform. Nothing is less true - after configuring a few kernels no one remembers that it was difficult! There are two ways for a Gentoo user to manage a manual kernel system, both of which are listed below:

**Wichtig**

Only one selection is required in the following subsection, if unsure of which to use then go with the first listed for now. It's always possible to switch at a later date if required.

##### Option 1 - Modprobed-db process

A very easy way to manage the kernel is to first install [sys-kernel/gentoo-kernel-bin](https://packages.gentoo.org/packages/sys-kernel/gentoo-kernel-bin) and use the [sys-kernel/modprobed-db](https://packages.gentoo.org/packages/sys-kernel/modprobed-db) to collect information about what the system requires. modprobed-db is a tool which monitors the system via crontab to add all modules of all devices over the system's life to make sure it everything a user needs is supported. For example, if an Xbox controller is added after installation, then modprobed-db will add the modules to be built next time the kernel is rebuilt.

For now please follow installing [sys-kernel/gentoo-kernel-bin](https://packages.gentoo.org/packages/sys-kernel/gentoo-kernel-bin) via [Distribution\_kernels](/wiki/Handbook:AMD64/Installation/Kernel#Distribution_kernels.2Fde "Handbook:AMD64/Installation/Kernel").

Next, install [sys-kernel/modprobed-db](https://packages.gentoo.org/packages/sys-kernel/modprobed-db):

`root #` `emerge --ask sys-kernel/modprobed-db`

Please watch out for further steps related to modprobed-db in the Handbook.

More on this topic can be found in the [Modprobed-db](/wiki/Modprobed-db "Modprobed-db") article.

##### Option 2 - Assisted manual process

This method allows a user to have full control of how their kernel is built with as minimal help from outside tools as they wish. Some could consider this as making it hard for the sake of it.

Eine Sache ist jedoch wahr: um einen Kernel manuell konfigurieren zu können, ist es wichtig, das System zu kennen. Die meisten Informationen erhalten Sie durch das Programm lspci, das im Paket [sys-apps/pciutils](https://packages.gentoo.org/packages/sys-apps/pciutils) enthalten ist.

`root #` `emerge --ask sys-apps/pciutils`

`root #` `lspci`

`root #` `lspci -v`

**Hinweis**

Innerhalb der chroot-Umgebung können Sie jegliche pcilib-Warnung (wie _pcilib: cannot open /sys/bus/pci/devices_) ignorieren, die lspci auswerfen könnte.

Eine weitere Quelle von Systeminformationen ist lsmod. Diese Anweisung zeigt Ihnen, welche Kernel-Module die Installations-CD verwendet. Dies liefert gute Hinweise darauf, was im Kernel aktiviert werden sollte.

Gehen Sie in das Kernel Quellverzeichnis und führen Sie make menuconfig aus. Dies wird eine menübasierte Konfigurationsmaske starten.

`root #` `cd /usr/src/linux
`

`root #` `make menuconfig
`

**Tipp**

To view the full list of make arguments available for the kernel, run `make help`.

The kernel has a method of autodetecting the modules currently being used on the installcd which will give a great starting point to allow a user to configure their own. This can be called by using:

`root #` `make localmodconfig`

Manually configuring should not be needed at this point, but if a user wish to check:

`root #` `make nconfig`

Now it's time to decide if modules signing is required in the steps listed in [here](/wiki/Handbook:AMD64/Installation/Kernel#Optional:_Signed_kernel_modules_2.2Fde "Handbook:AMD64/Installation/Kernel")

If not, proceed to building described [here](/wiki/Handbook:AMD64/Installation/Kernel#Compiling_and_installing.2Fde "Handbook:AMD64/Installation/Kernel")

##### Option 3 - Configuring by hand

The Linux kernel configuration has many, many sections and as configuring a kernel by hand is rarely needed thanks to the two tools listed above. The steps to do it by hand are now included at [Kernel/Gentoo\_Kernel\_Configuration\_Guide](/wiki/Kernel/Gentoo_Kernel_Configuration_Guide/de "Kernel/Gentoo Kernel Configuration Guide/de")

#### Optional: Signed kernel modules

To automatically sign the kernel modules enable CONFIG\_MODULE\_SIG\_ALL:

KERNEL **Sign kernel modules CONFIG\_MODULE\_SIG\_ALL**

```
[*] Enable loadable module support
  -*-   Module signature verification
    [*]     Automatically sign all modules
    Which hash algorithm should modules be signed with? (Sign modules with SHA-512) --->

```

Optionally change the hash algorithm if desired.

To enforce that all modules are signed with a valid signature, enable CONFIG\_MODULE\_SIG\_FORCE as well:

KERNEL **Enforce signed kernel modules CONFIG\_MODULE\_SIG\_FORCE**

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

KERNEL **Specify signing key CONFIG\_MODULE\_SIG\_KEY**

```
-*- Cryptographic API  --->
  Certificates for signature checking  --->
    (/path/to/kernel_key.pem) File name or PKCS#11 URI of module signing key

```

To also sign external kernel modules installed by other packages via `linux-mod-r1.eclass`, enable the [modules-sign](https://packages.gentoo.org/useflags/modules-sign) [USE flag/de](/wiki/USE_flag/de "USE flag/de") USE flag globally:

DATEI **`/etc/portage/make.conf`** **Enable module signing**

```
USE="modules-sign"
</div>

<div lang="en" dir="ltr" class="mw-content-ltr">
# Optionally, when using custom signing keys.
MODULES_SIGN_KEY="/path/to/kernel_key.pem"
MODULES_SIGN_CERT="/path/to/kernel_key.pem" # Only required if the MODULES_SIGN_KEY does not also contain the certificate
MODULES_SIGN_HASH="sha512" # Defaults to sha512

```

**Hinweis**

MODULES\_SIGN\_KEY and MODULES\_SIGN\_CERT may point to different files. For this example, the pem file generated by OpenSSL includes both the key and the accompanying certificate, and thus both variables are set to the same value.

#### Optional: Signing the kernel image (Secure Boot)

When signing the kernel image (for use on systems with [Secure Boot](/wiki/Secure_Boot "Secure Boot") enabled) it is recommended to set the following kernel config options:

KERNEL **Lockdown for secureboot**

```
General setup  --->
  Kexec and crash features  --->
    [*] Enable kexec system call
    [*] Enable kexec file based system call
    [*]   Verify kernel signature during kexec_file_load() syscall
    [*]     Require a valid signature in kexec_file_load() syscall
    [*]     Enable ""image"" signature verification support
</div>

<div lang="en" dir="ltr" class="mw-content-ltr">
[*] Enable loadable module support
  -*-   Module signature verification
    [*]     Require modules to be validly signed
    [*]     Automatically sign all modules
    Which hash algorithm should modules be signed with? (Sign modules with SHA-512) --->
</div>

<div lang="en" dir="ltr" class="mw-content-ltr">
Security options  --->
[*] Integrity subsystem
  [*] Basic module for enforcing kernel lockdown
  [*]   Enable lockdown LSM early in init
        Kernel default lockdown mode (Integrity)  --->
</div>

  <div lang="en" dir="ltr" class="mw-content-ltr">
[*]   Digital signature verification using multiple keyrings
  [*]     Enable asymmetric keys support
  -*-       Require all keys on the integrity keyrings be signed
  [*]       Provide keyring for platform/firmware trusted keys
  [*]       Provide a keyring to which Machine Owner Keys may be added
  [ ]         Enforce Machine Keyring CA Restrictions

```

Where ""image"" is a placeholder for the architecture specific image name. These options, from the top to the bottom: enforces that the kernel image in a kexec call must be signed (kexec allows replacing the kernel in-place), enforces that kernel modules are signed, enables lockdown integrity mode (prevents modifying the kernel at runtime), and enables various keychains.

On arches that do not natively support decompressing the kernel (e.g. **arm64** and **riscv**), the kernel must be built with its own decompressor (zboot):

KERNEL **zboot CONFIG\_EFI\_ZBOOT**

```
Device Drivers --->
  Firmware Drivers --->
    EFI (Extensible Firmware Interface) Support --->
      [*] Enable the generic EFI decompressor

```

After compilation of the kernel, as explained in the next section, the kernel image must be signed. First install [app-crypt/sbsigntools](https://packages.gentoo.org/packages/app-crypt/sbsigntools) and then sign the kernel image:

`root #` `emerge --ask app-crypt/sbsigntools`

`root #` `sbsign /usr/src/linux-x.y.z/path/to/kernel-image --cert /path/to/kernel_key.pem --key /path/to/kernel_key.pem --output /usr/src/linux-x.y.z/path/to/kernel-image`

**Hinweis**

For this example, the same key that was generated to sign the modules is used to sign the kernel image. It is also possible to generate and use a second separate key for signing the kernel image. The same OpenSSL command as in the previous section may be used again.

Then proceed with the installation.

To automatically sign EFI executables installed by other packages, enable the [secureboot](https://packages.gentoo.org/useflags/secureboot) [USE flag/de](/wiki/USE_flag/de "USE flag/de") USE flag globally:

DATEI **`/etc/portage/make.conf`** **Enable Secure Boot**

```
USE="modules-sign secureboot"
</div>

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

**Hinweis**

SECUREBOOT\_SIGN\_KEY and SECUREBOOT\_SIGN\_CERT may point to different files. For this example, the pem file generated by OpenSSL includes both the key and the accompanying certificate, and thus both variables are set to the same value.

**Hinweis**

When generating an [Unified Kernel Image](/wiki/Unified_Kernel_Image "Unified Kernel Image") with systemd's `ukify` the kernel image will be signed automatically before inclusion in the unified kernel image and it is not necessary to sign it manually.

#### Kompilieren und installieren

Mit beendeter Konfiguration ist es nun an der Zeit, den Kernel zu kompilieren und zu installieren. Schließen Sie die Konfiguration und starten Sie den Kompiliervorgang:

`root #` `make && make modules_install`

**Hinweis**

Es ist möglich, parallele Builds durch make -jX zu aktivieren, wobei `X` eine Ganzzahl ist und die Anzahl der Tasks angibt, die der Build-Prozess parallel starten darf. Dies ist ähnlich wie die Variable MAKEOPTS in der Anleitung zu /etc/portage/make.conf.

Wenn der Kernel fertig kompiliert ist, kopieren Sie das Kernel-Abbild nach /boot/. Dies erfolgt durch den Befehl make install:

`root #` `make install`

Dies kopiert das Kernel-Abbild zusammen mit der Datei System.map und der Kernel-Konfigurationsdatei nach /boot/.

Setzen Sie die Installation mit der [Konfiguration des Systems](/wiki/Handbook:AMD64/Installation/System/de "Handbook:AMD64/Installation/System/de") fort.

[← Installation des Basissystems](/wiki/Handbook:AMD64/Installation/Base/de "Previous part") [Anfang](/wiki/Handbook:AMD64/de "Handbook:AMD64/de") [Konfiguration des Systems →](/wiki/Handbook:AMD64/Installation/System/de "Next part")