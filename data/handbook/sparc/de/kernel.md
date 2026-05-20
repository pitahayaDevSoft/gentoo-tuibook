# Kernel

Other languages:

- Deutsch
- [English](/wiki/Handbook:SPARC/Installation/Kernel "Handbook:SPARC/Installation/Kernel (100% translated)")
- [español](/wiki/Handbook:SPARC/Installation/Kernel/es "Manual de Gentoo: SPARC/Instalación/Núcleo (100% translated)")
- [français](/wiki/Handbook:SPARC/Installation/Kernel/fr "Handbook:SPARC/Installation/Kernel/fr (100% translated)")
- [italiano](/wiki/Handbook:SPARC/Installation/Kernel/it "Handbook:SPARC/Installation/Kernel/it (50% translated)")
- [magyar](/wiki/Handbook:SPARC/Installation/Kernel/hu "Handbook:SPARC/Installation/Kernel/hu (100% translated)")
- [polski](/wiki/Handbook:SPARC/Installation/Kernel/pl "Handbook:SPARC/Installation/Kernel (100% translated)")
- [português do Brasil](/wiki/Handbook:SPARC/Installation/Kernel/pt-br "Handbook:SPARC/Installation/Kernel/pt-br (50% translated)")
- [русский](/wiki/Handbook:SPARC/Installation/Kernel/ru "Handbook:SPARC/Installation/Kernel (100% translated)")
- [தமிழ்](/wiki/Handbook:SPARC/Installation/Kernel/ta "கையேடு:SPARC/நிறுவல்/கர்னல் (100% translated)")
- [中文（中国大陆）‎](/wiki/Handbook:SPARC/Installation/Kernel/zh-cn "手册：SPARC/安装/配置Linux内核 (100% translated)")
- [日本語](/wiki/Handbook:SPARC/Installation/Kernel/ja "ハンドブック:SPARC/インストール/カーネル (100% translated)")
- [한국어](/wiki/Handbook:SPARC/Installation/Kernel/ko "Handbook:SPARC/Installation/Kernel/ko (100% translated)")

[SPARC Handbuch](/wiki/Handbook:SPARC/de "Handbook:SPARC/de")[Installation](/wiki/Handbook:SPARC/Full/Installation/de "Handbook:SPARC/Full/Installation/de")[Über die Installation](/wiki/Handbook:SPARC/Installation/About/de "Handbook:SPARC/Installation/About/de")[Auswahl des Mediums](/wiki/Handbook:SPARC/Installation/Media/de "Handbook:SPARC/Installation/Media/de")[Konfiguration des Netzwerks](/wiki/Handbook:SPARC/Installation/Networking/de "Handbook:SPARC/Installation/Networking/de")[Vorbereiten der Festplatte(n)](/wiki/Handbook:SPARC/Installation/Disks/de "Handbook:SPARC/Installation/Disks/de")[Installation des Stage Archivs](/wiki/Handbook:SPARC/Installation/Stage/de "Handbook:SPARC/Installation/Stage/de")[Installation des Basissystems](/wiki/Handbook:SPARC/Installation/Base/de "Handbook:SPARC/Installation/Base/de")Konfiguration des Kernels[Konfiguration des Systems](/wiki/Handbook:SPARC/Installation/System/de "Handbook:SPARC/Installation/System/de")[Installation der Tools](/wiki/Handbook:SPARC/Installation/Tools/de "Handbook:SPARC/Installation/Tools/de")[Konfiguration des Bootloaders](/wiki/Handbook:SPARC/Installation/Bootloader/de "Handbook:SPARC/Installation/Bootloader/de")[Abschluss](/wiki/Handbook:SPARC/Installation/Finalizing/de "Handbook:SPARC/Installation/Finalizing/de")[Arbeiten mit Gentoo](/wiki/Handbook:SPARC/Full/Working/de "Handbook:SPARC/Full/Working/de")[Portage-Einführung](/wiki/Handbook:SPARC/Working/Portage/de "Handbook:SPARC/Working/Portage/de")[USE-Flags](/wiki/Handbook:SPARC/Working/USE/de "Handbook:SPARC/Working/USE/de")[Portage-Features](/wiki/Handbook:SPARC/Working/Features/de "Handbook:SPARC/Working/Features/de")[Initskript-System](/wiki/Handbook:SPARC/Working/Initscripts/de "Handbook:SPARC/Working/Initscripts/de")[Umgebungsvariablen](/wiki/Handbook:SPARC/Working/EnvVar/de "Handbook:SPARC/Working/EnvVar/de")[Arbeiten mit Portage](/wiki/Handbook:SPARC/Full/Portage/de "Handbook:SPARC/Full/Portage/de")[Dateien und Verzeichnisse](/wiki/Handbook:SPARC/Portage/Files/de "Handbook:SPARC/Portage/Files/de")[Variablen](/wiki/Handbook:SPARC/Portage/Variables/de "Handbook:SPARC/Portage/Variables/de")[Mischen von Softwarezweigen](/wiki/Handbook:SPARC/Portage/Branches/de "Handbook:SPARC/Portage/Branches/de")[Zusätzliche Tools](/wiki/Handbook:SPARC/Portage/Tools/de "Handbook:SPARC/Portage/Tools/de")[Eigener Portage-Tree](/wiki/Handbook:SPARC/Portage/CustomTree/de "Handbook:SPARC/Portage/CustomTree/de")[Erweiterte Portage-Features](/wiki/Handbook:SPARC/Portage/Advanced/de "Handbook:SPARC/Portage/Advanced/de")[Netzwerk-Konfiguration](/wiki/Handbook:SPARC/Full/Networking/de "Handbook:SPARC/Full/Networking/de")[Zu Beginn](/wiki/Handbook:SPARC/Networking/Introduction/de "Handbook:SPARC/Networking/Introduction/de")[Fortgeschrittene Konfiguration](/wiki/Handbook:SPARC/Networking/Advanced/de "Handbook:SPARC/Networking/Advanced/de")[Modulare Vernetzung](/wiki/Handbook:SPARC/Networking/Modular/de "Handbook:SPARC/Networking/Modular/de")[Drahtlose Netzwerke](/wiki/Handbook:SPARC/Networking/Wireless/de "Handbook:SPARC/Networking/Wireless/de")[Funktionalität hinzufügen](/wiki/Handbook:SPARC/Networking/Extending/de "Handbook:SPARC/Networking/Extending/de")[Dynamisches Management](/wiki/Handbook:SPARC/Networking/Dynamic/de "Handbook:SPARC/Networking/Dynamic/de")

## Contents

- [1Optional: Firmware und/oder Microcode installieren](#Optional:_Firmware_und.2Foder_Microcode_installieren)
- [2Firmware](#Firmware)
  - [2.1Suggested: Linux Firmware](#Suggested:_Linux_Firmware)
    - [2.1.1Firmware Loading](#Firmware_Loading)
- [3Microcode](#Microcode)
- [4sys-kernel/installkernel](#sys-kernel.2Finstallkernel)
  - [4.1Bootloader](#Bootloader)
    - [4.1.1GRUB](#GRUB)
    - [4.1.2Traditional layout, other bootloaders (e.g. (e)lilo, syslinux, etc.)](#Traditional_layout.2C_other_bootloaders_.28e.g._.28e.29lilo.2C_syslinux.2C_etc..29)
  - [4.2Initramfs](#Initramfs)
    - [4.2.1Chroot detection](#Chroot_detection)
- [5Kernel konfigurieren und Kompilieren](#Kernel_konfigurieren_und_Kompilieren)
  - [5.1Alternative: Manuelle Konfiguration](#Alternative:_Manuelle_Konfiguration)
  - [5.2Installieren der Kernel-Quellen](#Installieren_der_Kernel-Quellen)
    - [5.2.1Option 2 - Assisted manual process](#Option_2_-_Assisted_manual_process)
    - [5.2.2Option 3 - Configuring by hand](#Option_3_-_Configuring_by_hand)
    - [5.2.3Optional: Signed kernel modules](#Optional:_Signed_kernel_modules)
  - [5.3Kompilieren und installieren](#Kompilieren_und_installieren)

### Optional: Firmware und/oder Microcode installieren

### Firmware

#### Suggested: Linux Firmware

On many systems, non-FOSS firmware is required for certain hardware to function. The [sys-kernel/linux-firmware](https://packages.gentoo.org/packages/sys-kernel/linux-firmware) package contains firmware for many, but not all, devices.

**Tipp**

Most wireless cards and GPUs require firmware to function.

`root #` `emerge --ask sys-kernel/linux-firmware`

**Hinweis**

Die Installation bestimmter Firmware-Pakete erfordert oft das Akzeptieren der zugehörigen Firmware-Lizenzen. Falls erforderlich, finden Sie im [Lizenzhandhabungsabschnitt](/wiki/Handbook:SPARC/Working/Portage/de#Licenses "Handbook:SPARC/Working/Portage/de") des Handbuchs Hilfe zum Akzeptieren von Lizenzen.

##### Firmware Loading

Firmware files are typically loaded when the associated kernel module is loaded. This means the firmware must be built into the kernel using **CONFIG\_EXTRA\_FIRMWARE** if the kernel module is set to _Y_ instead of _M_. In most cases, building-in a module which required firmware can complicate or break loading.

{{#switch: sparc \| amd64 \| x86 =

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

**Hinweis**

systemd-boot requires kernels to be installed to /efi.

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

## Kernel konfigurieren und Kompilieren

It can be a wise move to use the dist-kernel on the first boot as it provides a very simple method to rule out system issues and kernel config issues. Always having a known working kernel to fallback on can speed up debugging and alleviate anxiety when updating that your system will no longer boot.

A common misconception is that a manually compiled kernel will use a lot less RAM than a pre configured distribution kernel. Due to the modular nature of the Linux kernel, only what is needed by the system is loaded and in most cases resulting in similar memory usage.

**Wichtig**

Only one selection is required in the following subsection, if unsure of which to use then go with the first listed for now. It's always possible to switch at a later date if required.

In der Reihenfolge vom geringsten bis zum größten Aufwand:

[Full automation approach: Distribution kernels](/wiki/Handbook:SPARC/Installation/Kernel/de#Distribution_kernels "Handbook:SPARC/Installation/Kernel/de")A [Distribution Kernel](/wiki/Project:Distribution_Kernel "Project:Distribution Kernel") is used to configure, automatically build, and install the Linux kernel, its associated modules, and (optionally, but enabled by default) an initramfs file. Future kernel updates are fully automated since they are handled through the package manager, just like any other system package. It is possible [provide a custom kernel configuration file](/wiki/Project:Distribution_Kernel#Modifying_kernel_configuration "Project:Distribution Kernel") if customization is necessary. This is the least involved process and is perfect for new Gentoo users due to it working out-of-the-box and offering minimal involvement from the system administrator.[Hybrid approach: Genkernel](/wiki/Handbook:SPARC/Installation/Kernel/de#Alternative:_Genkernel "Handbook:SPARC/Installation/Kernel/de")New kernel sources are installed via the system package manager. System administrators use Gentoo's genkernel tool to generically configure, automatically build and install the Linux kernel, its associated modules, and (optionally, but _**not**_ enabled by default) an initramfs file. It is possible provide a custom kernel configuration file if customization is necessary. Future kernel configuration, compilation, and installation require the system administrator's involvement in the form of running eselect kernel, genkernel, and potentially other commands for each update.[Full manual approach](/wiki/Handbook:SPARC/Installation/Kernel/de#Alternative:_Manual_configuration "Handbook:SPARC/Installation/Kernel/de")New kernel sources are installed via the system package manager. The kernel is manually configured, built, and installed using the eselect kernel and a slew of make commands. Future kernel updates repeat the manual process of configuring, building, and installing the kernel files. This is the most involved process, but offers maximum control over the kernel update process.[Vollständiger Automatisierungsansatz: Distributionskernel](/wiki/Handbook:SPARC/Installation/Kernel/de#Distribution_kernels "Handbook:SPARC/Installation/Kernel/de")Ein [Distributionskernel](/wiki/Project:Distribution_Kernel "Project:Distribution Kernel") wird verwendet, um den Linux-Kernel, seine dazugehörigen Module und (optional, aber standardmäßig aktiviert) eine initramfs-Datei zu konfigurieren, automatisch zu bauen und zu installieren. Zukünftige Kernel-Aktualisierungen sind vollständig automatisiert, da sie über den Paketmanager abgewickelt werden, genauso wie jedes andere Systempaket. Es ist möglich [eine eigene Kernelkonfigurationsdatei bereitzustellen](/wiki/Project:Distribution_Kernel#Modifying_kernel_configuration "Project:Distribution Kernel"), wenn Anpassungen notwendig sind. Dies ist der am wenigsten aufwendige Prozess und eignet sich perfekt für neue Gentoo-Benutzer, da er sofort funktioniert und nur minimale Eingriffe des Systemadministrators erfordert.[Hybridansatz: Genkernel](/wiki/Handbook:SPARC/Installation/Kernel/de#Alternative:_Genkernel "Handbook:SPARC/Installation/Kernel/de")Neue Kernelquellen werden über den Systempaketmanager installiert. Systemadministratoren benutzen Gentoos genkernel Werkzeug, um den Linux-Kernel, seine dazugehörigen Module und (optional aber _nicht_ standardmäßig aktiviert) eine initramfs-Datei zu konfigurieren, automatisch zu bauen und zu installieren. Es ist möglich, eine benutzerdefinierte Kernelkonfigurationsdatei bereitzustellen, wenn eine Anpassung erforderlich ist. Zukünftige Kernelkonfigurationen, -kompilierungen und -installationen erfordern die Beteiligung des Systemadministrators in Form der Ausführung von eselect kernel, genkernel und möglicherweise andere Befehle für jede Aktualisierung.[Vollständige manuelle Vorgehensweise](/wiki/Handbook:SPARC/Installation/Kernel/de#Alternative:_Manual_configuration "Handbook:SPARC/Installation/Kernel/de")Neue Kernelquellen werden über den Systempaketmanager installiert. Der Kernel wird manuell konfiguriert, gebaut und installiert unter Verwendung des eselect kernel und einer Reihe von make Befehlen. Künftige Kernel-Updates wiederholen den manuellen Prozess des Konfigurierens, Erstellens und Installieren der Kernel-Dateien. Dies ist der aufwändigste Prozess, bietet aber maximale Kontrolle über den Kernel-Aktualisierungsprozess.

Der Kern, um den alle Distributionen gebaut sind, ist der Linux Kernel. Er ist die Schicht zwischen den Benutzerprogrammen und der Systemhardware. Obwohl das Handbuch seinen Benutzern verschiedene Kernel-Quellen anbietet, ist eine umfassendere Auflistung mit detaillierter Beschreibung auf der [Kernel Übersichtsseite](/wiki/Kernel/Overview "Kernel/Overview") verfügbar.

**Tipp**

Kernel installation tasks such as copying the kernel image to /boot or the [EFI System Partition](/wiki/EFI_System_Partition/de "EFI System Partition/de"), generating an [initramfs](/wiki/Initramfs "Initramfs") and/or [Unified Kernel Image](/wiki/Unified_Kernel_Image "Unified Kernel Image"), updating bootloader configuration, can be automated with [installkernel](/wiki/Installkernel "Installkernel"). Users may wish to configure and install [sys-kernel/installkernel](https://packages.gentoo.org/packages/sys-kernel/installkernel) before proceeding. See the [Kernel installation section below](/wiki/Handbook:SPARC/Installation/Kernel#Kernel_installation.2Fde "Handbook:SPARC/Installation/Kernel") for more more information.

### Alternative: Manuelle Konfiguration

### Installieren der Kernel-Quellen

Für die Installation und Kompilierung des Kernels für sparc-basierte Systeme empfiehlt Gentoo das Paket [sys-kernel/gentoo-sources](https://packages.gentoo.org/packages/sys-kernel/gentoo-sources).

Wählen Sie eine geeignete Kernelquelle und installieren Sie sie mit emerge:

`root #` `emerge --ask sys-kernel/gentoo-sources`

Dieser Befehl installiert die Quellen des Linux-Kernels in dem Verzeichnis /usr/src/. Im Namen des angelegten Unterverzeichnisses steht die Version des Linux-Kernels (beispielsweise /usr/src/linux-4.9.16-gentoo). Der emerge-Befehl erzeugt einen symbolischen Link /usr/src/linux auf dieses Verzeichnis, wenn das USE-Flag `symlink` bei dem Paket gesetzt ist.

Es ist gute Praxis, einen Symlink /usr/src/linux auf das Verzeichnis mit dem Source-Code der Kernel-Version zeigen zu lassen, die auf dem System läuft. Dieser Symlink wird nicht automatisch erzeugt. Er kann jedoch einfach mit Hilfe des eselect Kernel Modules erzeugt werden.

Weitere Informationen über Zweck und Handhabung des Symlinks finden Sie unter [Kernel/Upgrade](/wiki/Kernel/Upgrade/de "Kernel/Upgrade/de").

Lassen Sie eine Liste der installierten Kernel anzeigen:

`root #` `eselect kernel list`

```
Available kernel symlink targets:
  [1]   linux-6.19.3-gentoo

```

Mit folgendem Befehl können Sie einen symbolischen Link namens linux anlegen:

`root #` `eselect kernel set 1`

`root #` `ls -l /usr/src/linux`

```
lrwxrwxrwx    1 root   root    12 Oct 13 11:04 /usr/src/linux -> linux-6.19.3-gentoo

```

Manually configuring a kernel is commonly seen as one of the most difficult procedures a system administrator has to perform. Nothing is less true - after configuring a few kernels no one remembers that it was difficult! There are two ways for a Gentoo user to manage a manual kernel system, both of which are listed below:

**Wichtig**

Only one selection is required in the following subsection, if unsure of which to use then go with the first listed for now. It's always possible to switch at a later date if required.

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

Now it's time to decide if modules signing is required in the steps listed in [here](/wiki/Handbook:SPARC/Installation/Kernel#Optional:_Signed_kernel_modules_2.2Fde "Handbook:SPARC/Installation/Kernel")

If not, proceed to building described [here](/wiki/Handbook:SPARC/Installation/Kernel#Compiling_and_installing.2Fde "Handbook:SPARC/Installation/Kernel")

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

### Kompilieren und installieren

Mit beendeter Konfiguration ist es an der Zeit den Kernel zu kompilieren und zu installieren. Schließen Sie die Konfiguration und starten Sie den Kompiliervorgang:

`root #` `make && make modules_install`

**Hinweis**

Es ist möglich parallele Builds durch `make -jX` zu aktivieren. Wobei X die Anzahl der Tasks ist, die der Build-Prozess parallel starten darf. Dies ist ähnlich der Anleitung zu /etc/portage/make.conf, mit der Variable `MAKEOPTS`.

Wenn der Kernel mit dem Kompilieren fertig ist, überprüfen Sie die Größe der resultierenden Datei:

`root #` `ls -lh arch/sparc/boot/image`

```
-rw-r--r--    1 root     root         2.4M Oct 25 14:38 image

```

Wenn die (unkomprimierte) Größe einen Wert größer als 7,5 MB hat, konfigurieren Sie den Kernel erneut bis diese Grenze nicht mehr überschritten wird. Ein Weg dies zu erreichen ist die meisten Kernel-Treiber als Module zu kompilieren. Falls Sie dies ignorieren, kann es zu einem nicht bootenden Kernel führen.

Falls der Kernel nur ein bisschen zu groß ist, können Sie versuchen ihn mit dem Befehl `strip` ein bisschen zu verkleinern.

`root #` `strip -R .comment -R .note arch/sparc/boot/image`

Kopieren Sie das Kernelabbild am Ende in das Verzeichnis /boot/.

`root #` `cp arch/sparc/boot/image /boot/kernel-6.19.3-gentoo`

Setzen Sie die Installation mit der [Konfiguration des Systems](/wiki/Handbook:SPARC/Installation/System/de "Handbook:SPARC/Installation/System/de") fort.

[← Installation des Basissystems](/wiki/Handbook:SPARC/Installation/Base/de "Previous part") [Anfang](/wiki/Handbook:SPARC/de "Handbook:SPARC/de") [Konfiguration des Systems →](/wiki/Handbook:SPARC/Installation/System/de "Next part")