# Kernel

Other languages:

- [Deutsch](/wiki/Handbook:PPC64/Installation/Kernel/de "Handbuch:PPC64/Installation/Kernel (100% translated)")
- [English](/wiki/Handbook:PPC64/Installation/Kernel "Handbook:PPC64/Installation/Kernel (100% translated)")
- [español](/wiki/Handbook:PPC64/Installation/Kernel/es "Manual de Gentoo: PPC64/Instalación/Núcleo (100% translated)")
- [français](/wiki/Handbook:PPC64/Installation/Kernel/fr "Handbook:PPC64/Installation/Kernel/fr (100% translated)")
- [italiano](/wiki/Handbook:PPC64/Installation/Kernel/it "Handbook:PPC64/Installation/Kernel/it (100% translated)")
- [magyar](/wiki/Handbook:PPC64/Installation/Kernel/hu "Handbook:PPC64/Installation/Kernel/hu (100% translated)")
- polski
- [português do Brasil](/wiki/Handbook:PPC64/Installation/Kernel/pt-br "Handbook:PPC64/Installation/Kernel/pt-br (50% translated)")
- [русский](/wiki/Handbook:PPC64/Installation/Kernel/ru "Handbook:PPC64/Installation/Kernel (100% translated)")
- [தமிழ்](/wiki/Handbook:PPC64/Installation/Kernel/ta "கையேடு:PPC64/நிறுவல்/கர்னல் (100% translated)")
- [中文（中国大陆）‎](/wiki/Handbook:PPC64/Installation/Kernel/zh-cn "手册：PPC64/安装/配置Linux内核 (100% translated)")
- [日本語](/wiki/Handbook:PPC64/Installation/Kernel/ja "ハンドブック:PPC64/インストール/カーネル (100% translated)")
- [한국어](/wiki/Handbook:PPC64/Installation/Kernel/ko "Handbook:PPC64/Installation/Kernel/ko (100% translated)")

[PPC64 Handbook](/wiki/Handbook:PPC64/pl "Handbook:PPC64/pl")[Instalacja](/wiki/Handbook:PPC64/Full/Installation/pl "Handbook:PPC64/Full/Installation/pl")[O instalacji](/wiki/Handbook:PPC64/Installation/About/pl "Handbook:PPC64/Installation/About/pl")[Wybór medium instalacyjnego](/wiki/Handbook:PPC64/Installation/Media/pl "Handbook:PPC64/Installation/Media/pl")[Konfiguracja sieci](/wiki/Handbook:PPC64/Installation/Networking/pl "Handbook:PPC64/Installation/Networking/pl")[Przygotowanie dysków](/wiki/Handbook:PPC64/Installation/Disks/pl "Handbook:PPC64/Installation/Disks/pl")[Instalacja etapu 3](/wiki/Handbook:PPC64/Installation/Stage/pl "Handbook:PPC64/Installation/Stage/pl")[Instalacja systemu podstawowego](/wiki/Handbook:PPC64/Installation/Base/pl "Handbook:PPC64/Installation/Base/pl")Konfiguracja jądra[Konfiguracja systemu](/wiki/Handbook:PPC64/Installation/System/pl "Handbook:PPC64/Installation/System/pl")[Instalacja narzędzi](/wiki/Handbook:PPC64/Installation/Tools/pl "Handbook:PPC64/Installation/Tools/pl")[Instalacja systemu rozruchowego](/wiki/Handbook:PPC64/Installation/Bootloader/pl "Handbook:PPC64/Installation/Bootloader/pl")[Finalizacja](/wiki/Handbook:PPC64/Installation/Finalizing/pl "Handbook:PPC64/Installation/Finalizing/pl")[Praca z Gentoo](/wiki/Handbook:PPC64/Full/Working/pl "Handbook:PPC64/Full/Working/pl")[Wstęp do Portage](/wiki/Handbook:PPC64/Working/Portage/pl "Handbook:PPC64/Working/Portage/pl")[Flagi USE](/wiki/Handbook:PPC64/Working/USE/pl "Handbook:PPC64/Working/USE/pl")[Funkcje portage](/wiki/Handbook:PPC64/Working/Features/pl "Handbook:PPC64/Working/Features/pl")[System initscript](/wiki/Handbook:PPC64/Working/Initscripts/pl "Handbook:PPC64/Working/Initscripts/pl")[Zmienne środowiskowe](/wiki/Handbook:PPC64/Working/EnvVar/pl "Handbook:PPC64/Working/EnvVar/pl")[Praca z Portage](/wiki/Handbook:PPC64/Full/Portage/pl "Handbook:PPC64/Full/Portage/pl")[Pliki i katalogi](/wiki/Handbook:PPC64/Portage/Files/pl "Handbook:PPC64/Portage/Files/pl")[Zmienne](/wiki/Handbook:PPC64/Portage/Variables/pl "Handbook:PPC64/Portage/Variables/pl")[Mieszanie działów oprogramowania](/wiki/Handbook:PPC64/Portage/Branches/pl "Handbook:PPC64/Portage/Branches/pl")[Dodatkowe narzędzia](/wiki/Handbook:PPC64/Portage/Tools/pl "Handbook:PPC64/Portage/Tools/pl")[Custom package repository](/wiki/Handbook:PPC64/Portage/CustomTree/pl "Handbook:PPC64/Portage/CustomTree/pl")[Funkcje zaawansowane](/wiki/Handbook:PPC64/Portage/Advanced/pl "Handbook:PPC64/Portage/Advanced/pl")[Konfiguracja sieci](/wiki/Handbook:PPC64/Full/Networking/pl "Handbook:PPC64/Full/Networking/pl")[Zaczynamy](/wiki/Handbook:PPC64/Networking/Introduction/pl "Handbook:PPC64/Networking/Introduction/pl")[Zaawansowana konfiguracja](/wiki/Handbook:PPC64/Networking/Advanced/pl "Handbook:PPC64/Networking/Advanced/pl")[Sieć modularna](/wiki/Handbook:PPC64/Networking/Modular/pl "Handbook:PPC64/Networking/Modular/pl")[Sieć bezprzewodowa](/wiki/Handbook:PPC64/Networking/Wireless/pl "Handbook:PPC64/Networking/Wireless/pl")[Dodawanie funkcjonalności](/wiki/Handbook:PPC64/Networking/Extending/pl "Handbook:PPC64/Networking/Extending/pl")[Dynamiczne zarządzanie](/wiki/Handbook:PPC64/Networking/Dynamic/pl "Handbook:PPC64/Networking/Dynamic/pl")

## Contents

- [1Installing firmware and/or microcode](#Installing_firmware_and.2For_microcode)
  - [1.1Firmware](#Firmware)
    - [1.1.1Suggested: Linux Firmware](#Suggested:_Linux_Firmware)
      - [1.1.1.1Firmware Loading](#Firmware_Loading)
- [2sys-kernel/installkernel](#sys-kernel.2Finstallkernel)
  - [2.1Bootloader](#Bootloader)
    - [2.1.1GRUB](#GRUB)
    - [2.1.2Traditional layout, other bootloaders (e.g. (e)lilo, syslinux, etc.)](#Traditional_layout.2C_other_bootloaders_.28e.g._.28e.29lilo.2C_syslinux.2C_etc..29)
  - [2.2Initramfs](#Initramfs)
    - [2.2.1Chroot detection](#Chroot_detection)
- [3Kernel selection](#Kernel_selection)
  - [3.1Distribution kernels](#Distribution_kernels)
    - [3.1.1Optional: Signed kernel modules](#Optional:_Signed_kernel_modules)
    - [3.1.2Installing a distribution kernel](#Installing_a_distribution_kernel)
    - [3.1.3Upgrading and cleaning up](#Upgrading_and_cleaning_up)
    - [3.1.4Post-install/upgrade tasks](#Post-install.2Fupgrade_tasks)
      - [3.1.4.1Manually rebuilding the initramfs or Unified Kernel Image](#Manually_rebuilding_the_initramfs_or_Unified_Kernel_Image)
  - [3.2Alternative: Manual configuration](#Alternative:_Manual_configuration)
    - [3.2.1Installing and Configuring the kernel sources](#Installing_and_Configuring_the_kernel_sources)
      - [3.2.1.1Option 1 - Modprobed-db process](#Option_1_-_Modprobed-db_process)
      - [3.2.1.2Option 2 - Assisted manual process](#Option_2_-_Assisted_manual_process)
      - [3.2.1.3Option 3 - Configuring by hand](#Option_3_-_Configuring_by_hand)
    - [3.2.2Optional: Signed kernel modules](#Optional:_Signed_kernel_modules_2)
    - [3.2.3Compiling and installing](#Compiling_and_installing)

## Installing firmware and/or microcode

### Firmware

#### Suggested: Linux Firmware

On many systems, non-FOSS firmware is required for certain hardware to function. The [sys-kernel/linux-firmware](https://packages.gentoo.org/packages/sys-kernel/linux-firmware) package contains firmware for many, but not all, devices.

**Wskazówka**

Most wireless cards and GPUs require firmware to function.

`root #` `emerge --ask sys-kernel/linux-firmware`

**Informacja**

Installing certain firmware packages often requires accepting the associated firmware licenses. If necessary, visit the [license handling section](/wiki/Handbook:PPC64/Working/Portage/pl#Licenses "Handbook:PPC64/Working/Portage/pl") of the Handbook for help on accepting licenses.

##### Firmware Loading

Firmware files are typically loaded when the associated kernel module is loaded. This means the firmware must be built into the kernel using **CONFIG\_EXTRA\_FIRMWARE** if the kernel module is set to _Y_ instead of _M_. In most cases, building-in a module which required firmware can complicate or break loading.

## sys-kernel/installkernel

[Installkernel](/wiki/Installkernel "Installkernel") may be used to automate the kernel installation, [initramfs](/wiki/Initramfs "Initramfs") generation, [unified kernel image](/wiki/Unified_kernel_image "Unified kernel image") generation and bootloader configuration, among other things. [sys-kernel/installkernel](https://packages.gentoo.org/packages/sys-kernel/installkernel) implements two paths of achieving this: the traditional installkernel originating from Debian and [systemd](/wiki/Systemd "Systemd")'s kernel-install. Which one to choose depends, among other things, on the system's bootloader. By default, systemd's kernel-install is used on systemd profiles, while the traditional installkernel is the default for other profiles.

### Bootloader

Now is the time to think about which bootloader the user wants for the system.

**Ważne**

Only one selection is required in the following subsection, if unsure of which to use then go with the first listed for now. It's always possible to switch at a later date if required.

#### GRUB

Users of GRUB can use either systemd's kernel-install or the traditional Debian installkernel. The [systemd](https://packages.gentoo.org/useflags/systemd) [USE flag/pl](/wiki/USE_flag/pl "USE flag/pl") USE flag switches between these implementations. To automatically run grub-mkconfig when installing the kernel, enable the [grub](https://packages.gentoo.org/useflags/grub) [USE flag/pl](/wiki/USE_flag/pl "USE flag/pl") [USE flag](/wiki/USE_flag/pl "USE flag/pl").

**Informacja**

GRUB requires kernels to be installed to /boot.

FILE **`/etc/portage/package.use/installkernel`**

```
sys-kernel/installkernel grub

```

`root #` `emerge --ask sys-kernel/installkernel`

**Informacja**

systemd-boot requires kernels to be installed to /efi.

**Informacja**

When [app-emulation/virt-firmware](https://packages.gentoo.org/packages/app-emulation/virt-firmware) is used to configure the UEFI ensure that the kernel-bootcfg-boot-successful service is enabled before attempting to install the kernel. This implementation of EFIstub booting is the default for systemd systems.

`root #` `systemctl enable kernel-bootcfg-boot-successful.service`

**Informacja**

EFIstub requires kernels to be installed to /efi.

#### Traditional layout, other bootloaders (e.g. (e)lilo, syslinux, etc.)

The traditional /boot layout (for e.g. (e)LILO, syslinux, etc.) is used by default if the [grub](https://packages.gentoo.org/useflags/grub) [USE flag/pl](/wiki/USE_flag/pl "USE flag/pl"), [systemd-boot](https://packages.gentoo.org/useflags/systemd-boot) [USE flag/pl](/wiki/USE_flag/pl "USE flag/pl"), [efistub](https://packages.gentoo.org/useflags/efistub) [USE flag/pl](/wiki/USE_flag/pl "USE flag/pl") and [uki](https://packages.gentoo.org/useflags/uki) [USE flag/pl](/wiki/USE_flag/pl "USE flag/pl") USE flags are **not** enabled. No further action is required.

### Initramfs

An **init** ial **ram**-based **f** ile **s** ystem, or [initramfs](/wiki/Initramfs "Initramfs"), may be required for a system to boot. A wide of variety of cases may necessitate one, but common cases include:

- Kernels where storage/filesystem drivers are modules.
- Layouts with /usr/ or /var/ on separate partitions.
- Encrypted root filesystems.

**Wskazówka**

[Distribution kernels](/wiki/Project:Distribution_Kernel "Project:Distribution Kernel") are designed to be used with an initramfs, as many storage and filesystem drivers are built as modules.

In addition to mounting the root filesystem, an initramfs may also perform other tasks such as:

- Running **f** ile **s** ystem **c** onsistency chec **k** fsck, a tool to check and repair consistency of a file system in such events of uncleanly shutdown a system.
- Providing a recovery environment in the event of late-boot failures.

[Installkernel](/wiki/Installkernel "Installkernel") can automatically generate an initramfs when installing the kernel if the [dracut](https://packages.gentoo.org/useflags/dracut) [USE flag/pl](/wiki/USE_flag/pl "USE flag/pl") or [ugrd](https://packages.gentoo.org/useflags/ugrd) [USE flag/pl](/wiki/USE_flag/pl "USE flag/pl") USE flag is enabled:

FILE **`/etc/portage/package.use/installkernel`**

```
sys-kernel/installkernel dracut

```

#### Chroot detection

Bootloaders such as [systemd-boot](/wiki/Systemd/systemd-boot "Systemd/systemd-boot") and [EFI stub](/wiki/EFI_stub "EFI stub") use the kernel arguments of the running system as set in /proc/cmdline by default. Because of the wide range of ways Gentoo can be installed users will randomly get tripped up by this.

To help solve any issues this may cause, [sys-kernel/installkernel](https://packages.gentoo.org/packages/sys-kernel/installkernel) will check if it is running in a chroot and fail if the kernel command line is not explicitly configured. Otherwise the bootloader would use the install media's boot arguments which would lead to boot failure.

One way to satisfy [sys-kernel/installkernel](https://packages.gentoo.org/packages/sys-kernel/installkernel) is by using Dracut's config file to set the root partition UUID as shown below, or alternatively for more information on what this check helps with and different ways to configure it, see [Installkernel#Install\_chroot\_check](/wiki/Installkernel#Install_chroot_check.2Fpl "Installkernel").

`root #` `mkdir /etc/dracut.conf.d`

`root #` `blkid`

```
/dev/sda3: UUID="2122cd72-94d7-4dcc-821e-3705926deecc"
```

In the above example, the root partition is /dev/sda3 and the UUID is 2122cd72-94d7-4dcc-821e-3705926deecc.
The dracut config file would then look like:

FILE **`/etc/dracut.conf.d/00-installkernel.conf`**

```
kernel_cmdline=" root=UUID=2122cd72-94d7-4dcc-821e-3705926deecc " # Note leading and trailing spaces

```

`root #` `emerge --ask sys-kernel/installkernel`

## Kernel selection

It can be a wise move to use the dist-kernel on the first boot as it provides a very simple method to rule out system issues and kernel config issues. Always having a known working kernel to fallback on can speed up debugging and alleviate anxiety when updating that your system will no longer boot.

A common misconception is that a manually compiled kernel will use a lot less RAM than a pre configured distribution kernel. Due to the modular nature of the Linux kernel, only what is needed by the system is loaded and in most cases resulting in similar memory usage.

**Ważne**

Only one selection is required in the following subsection, if unsure of which to use then go with the first listed for now. It's always possible to switch at a later date if required.

Ranked from least involved to most involved:

[Full automation approach: Distribution kernels](/wiki/Handbook:PPC64/Installation/Kernel/pl#Distribution_kernels "Handbook:PPC64/Installation/Kernel/pl")A [Distribution Kernel](/wiki/Project:Distribution_Kernel "Project:Distribution Kernel") is used to configure, automatically build, and install the Linux kernel, its associated modules, and (optionally, but enabled by default) an initramfs file. Future kernel updates are fully automated since they are handled through the package manager, just like any other system package. It is possible [provide a custom kernel configuration file](/wiki/Project:Distribution_Kernel#Modifying_kernel_configuration "Project:Distribution Kernel") if customization is necessary. This is the least involved process and is perfect for new Gentoo users due to it working out-of-the-box and offering minimal involvement from the system administrator.[Manual approach](/wiki/Handbook:PPC64/Installation/Kernel/pl#Alternative:_Manual_configuration "Handbook:PPC64/Installation/Kernel/pl")New kernel sources are installed via the system package manager. The kernel is manually configured, built, and installed using the eselect kernel and a slew of make commands. Future kernel updates repeat the manual process of configuring, building, and installing the kernel files. This is the most involved process, but offers maximum control over the kernel update process.

The core around which all distributions are built is the Linux kernel. It is the layer between the user's programs and the system hardware. Although the handbook provides its users several possible kernel sources, a more comprehensive listing with more detailed descriptions is available at the [kernel packages page](/wiki/Kernel/Packages "Kernel/Packages").

**Wskazówka**

Kernel installation tasks such as copying the kernel image to /boot or the [EFI System Partition](/wiki/EFI_System_Partition "EFI System Partition"), generating an [initramfs](/wiki/Initramfs "Initramfs") and/or [Unified Kernel Image](/wiki/Unified_Kernel_Image "Unified Kernel Image"), updating bootloader configuration, can be automated with [installkernel](/wiki/Installkernel "Installkernel"). Users may wish to configure and install [sys-kernel/installkernel](https://packages.gentoo.org/packages/sys-kernel/installkernel) before proceeding. See the [Kernel installation section below](/wiki/Handbook:PPC64/Installation/Kernel#Kernel_installation.2Fpl "Handbook:PPC64/Installation/Kernel") for more more information.

### Distribution kernels

_[Distribution Kernels](/wiki/Project:Distribution_Kernel "Project:Distribution Kernel")_ are ebuilds that cover the complete process of unpacking, configuring, compiling, and installing the kernel. The primary advantage of this method is that the kernels are updated to new versions by the package manager as part of @world upgrade. This requires no more involvement than running an emerge command. Distribution kernels default to a configuration supporting the majority of hardware, however two mechanisms are offered for customization: savedconfig and config snippets. See the project page for [more details on configuration.](/wiki/Project:Distribution_Kernel#Modifying_kernel_configuration "Project:Distribution Kernel")

##### Optional: Signed kernel modules

The kernel modules in the prebuilt distribution kernel ([sys-kernel/gentoo-kernel-bin](https://packages.gentoo.org/packages/sys-kernel/gentoo-kernel-bin)) are already signed. To sign the modules of kernels built from source enable the [modules-sign](https://packages.gentoo.org/useflags/modules-sign) [USE flag/pl](/wiki/USE_flag/pl "USE flag/pl") USE flag, and optionally specify which key to use for signing in [/etc/portage/make.conf](/wiki//etc/portage/make.conf/pl "/etc/portage/make.conf/pl"):

FILE **`/etc/portage/make.conf`** **Enable module signing**

```
USE="modules-sign"

# Optionally, to use custom signing keys.
MODULES_SIGN_KEY="/path/to/kernel_key.pem"
MODULES_SIGN_CERT="/path/to/kernel_key.pem" # Only required if the MODULES_SIGN_KEY does not also contain the certificate.
MODULES_SIGN_HASH="sha512" # Defaults to sha512.

```

If MODULES\_SIGN\_KEY is not specified the kernel build system will generate a key, it will be stored in /usr/src/linux-x.y.z/certs. It is recommended to manually generate a key to ensure that it will be the same for each kernel release. A key may be generated with:

`root #` `openssl req -new -noenc -utf8 -sha256 -x509 -outform PEM -out kernel_key.pem -keyout kernel_key.pem`

**Informacja**

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

#### Installing a distribution kernel

To build a kernel with Gentoo patches from source, type:

`root #` `emerge --ask sys-kernel/gentoo-kernel`

System administrators who want to avoid compiling the kernel sources locally can instead use precompiled kernel images:

`root #` `emerge --ask sys-kernel/gentoo-kernel-bin`

**Ważne**

_[Distribution Kernels](/wiki/Project:Distribution_Kernel "Project:Distribution Kernel")_, such as [sys-kernel/gentoo-kernel](https://packages.gentoo.org/packages/sys-kernel/gentoo-kernel) and [sys-kernel/gentoo-kernel-bin](https://packages.gentoo.org/packages/sys-kernel/gentoo-kernel-bin), by default, expect to be installed alongside an [initramfs](/wiki/Handbook:PPC64/Installation/Kernel#Initramfs.2Fpl "Handbook:PPC64/Installation/Kernel"). Before running emerge to install the kernel users should ensure that [sys-kernel/installkernel](https://packages.gentoo.org/packages/sys-kernel/installkernel) has been configured to utilize an initramfs generator (for example [Dracut](/wiki/Dracut "Dracut")) as described in the [installkernel section](/wiki/Handbook:PPC64/Installation/Kernel#Initramfs.2Fpl "Handbook:PPC64/Installation/Kernel").

#### Upgrading and cleaning up

Once the kernel is installed, the package manager will automatically update it to newer versions. The previous versions will be kept until the package manager is requested to clean up stale packages. To reclaim disk space, stale packages can be trimmed by periodically running emerge with the `--depclean` option:

`root #` `emerge --depclean`

Alternatively, to specifically clean up old kernel versions:

`root #` `emerge --prune sys-kernel/gentoo-kernel sys-kernel/gentoo-kernel-bin`

**Wskazówka**

By design, emerge only removes the kernel build directory. It does not actually remove the kernel modules, nor the installed kernel image. To completely clean-up old kernels, the [app-admin/eclean-kernel](https://packages.gentoo.org/packages/app-admin/eclean-kernel) tool may be used.

#### Post-install/upgrade tasks

An upgrade of a distribution kernel is capable of triggering an automatic rebuild for external kernel modules installed by other packages (for example: [sys-fs/zfs-kmod](https://packages.gentoo.org/packages/sys-fs/zfs-kmod) or [x11-drivers/nvidia-drivers](https://packages.gentoo.org/packages/x11-drivers/nvidia-drivers)). This automated behaviour is enabled by enabling the [dist-kernel](https://packages.gentoo.org/useflags/dist-kernel) [USE flag/pl](/wiki/USE_flag/pl "USE flag/pl") USE flag. When required, this same flag will also trigger re-generation of the [initramfs](/wiki/Initramfs "Initramfs").

It is highly recommended to enable this flag globally via /etc/portage/make.conf when using a distribution kernel:

FILE **`/etc/portage/make.conf`** **Enabling USE=dist-kernel**

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

After installing the Distribution Kernel successfully, it is now time to proceed to the next section: [Configuring the system](/wiki/Handbook:PPC64/Installation/System/pl "Handbook:PPC64/Installation/System/pl").

### Alternative: Manual configuration

#### Installing and Configuring the kernel sources

When manually compiling the kernel for ppc64-based systems, Gentoo recommends the [sys-kernel/gentoo-sources](https://packages.gentoo.org/packages/sys-kernel/gentoo-sources) package.

Choose an appropriate kernel source and install it using emerge:

`root #` `emerge --ask sys-kernel/gentoo-sources`

This will install the Linux kernel sources in /usr/src/ using the specific kernel version in the path. It will not create a symbolic link by itself without the [symlink](https://packages.gentoo.org/useflags/symlink) [USE flag/pl](/wiki/USE_flag/pl "USE flag/pl") USE flag being enabled on the chosen kernel sources package.

It is conventional for a /usr/src/linux symlink to be maintained, such that it refers to whichever sources correspond with the currently running kernel. However, this symbolic link will not be created by default. An easy way to create the symbolic link is to utilize eselect's kernel module.

For further information regarding the purpose of the symlink, and how to manage it, please refer to [Kernel/Upgrade](/wiki/Kernel/Upgrade "Kernel/Upgrade").

First, list all installed kernels:

`root #` `eselect kernel list`

```
Available kernel symlink targets:
  [1]   linux-6.18.8-gentoo

```

In order to create a symbolic link called linux, use:

`root #` `eselect kernel set 1`

`root #` `ls -l /usr/src/linux`

```
lrwxrwxrwx    1 root   root    12 Oct 13 11:04 /usr/src/linux -> linux-6.18.8-gentoo

```

Manually configuring a kernel is commonly seen as one of the most difficult procedures a system administrator has to perform. Nothing is less true - after configuring a few kernels no one remembers that it was difficult! There are two ways for a Gentoo user to manage a manual kernel system, both of which are listed below:

**Ważne**

Only one selection is required in the following subsection, if unsure of which to use then go with the first listed for now. It's always possible to switch at a later date if required.

##### Option 1 - Modprobed-db process

A very easy way to manage the kernel is to first install [sys-kernel/gentoo-kernel-bin](https://packages.gentoo.org/packages/sys-kernel/gentoo-kernel-bin) and use the [sys-kernel/modprobed-db](https://packages.gentoo.org/packages/sys-kernel/modprobed-db) to collect information about what the system requires. modprobed-db is a tool which monitors the system via crontab to add all modules of all devices over the system's life to make sure it everything a user needs is supported. For example, if an Xbox controller is added after installation, then modprobed-db will add the modules to be built next time the kernel is rebuilt.

For now please follow installing [sys-kernel/gentoo-kernel-bin](https://packages.gentoo.org/packages/sys-kernel/gentoo-kernel-bin) via [Distribution\_kernels](/wiki/Handbook:PPC64/Installation/Kernel#Distribution_kernels.2Fpl "Handbook:PPC64/Installation/Kernel").

Next, install [sys-kernel/modprobed-db](https://packages.gentoo.org/packages/sys-kernel/modprobed-db):

`root #` `emerge --ask sys-kernel/modprobed-db`

Please watch out for further steps related to modprobed-db in the Handbook.

More on this topic can be found in the [Modprobed-db](/wiki/Modprobed-db "Modprobed-db") article.

##### Option 2 - Assisted manual process

This method allows a user to have full control of how their kernel is built with as minimal help from outside tools as they wish. Some could consider this as making it hard for the sake of it.

However, with this choice one thing is true: it is vital to know the system when a kernel is configured manually. Most information can be gathered by emerging [sys-apps/pciutils](https://packages.gentoo.org/packages/sys-apps/pciutils) which contains the lspci command:

`root #` `emerge --ask sys-apps/pciutils`

**Informacja**

Inside the chroot, it is safe to ignore any pcilib warnings (like _pcilib: cannot open /sys/bus/pci/devices_) that lspci might throw out.

Another source of system information is to run lsmod to see what kernel modules the installation CD uses as it might provide a nice hint on what to enable.

Now go to the kernel source directory.

`root #` `cd /usr/src/linux
`

**Wskazówka**

To view the full list of make arguments available for the kernel, run `make help`.

The kernel has a method of autodetecting the modules currently being used on the installcd which will give a great starting point to allow a user to configure their own. This can be called by using:

`root #` `make localmodconfig`

Manually configuring should not be needed at this point, but if a user wish to check:

`root #` `make nconfig`

Now it's time to decide if modules signing is required in the steps listed in [here](/wiki/Handbook:PPC64/Installation/Kernel#Optional:_Signed_kernel_modules_2.2Fpl "Handbook:PPC64/Installation/Kernel")

If not, proceed to building described [here](/wiki/Handbook:PPC64/Installation/Kernel#Compiling_and_installing.2Fpl "Handbook:PPC64/Installation/Kernel")

##### Option 3 - Configuring by hand

The Linux kernel configuration has many, many sections and as configuring a kernel by hand is rarely needed thanks to the two tools listed above. The steps to do it by hand are now included at [Kernel/Gentoo\_Kernel\_Configuration\_Guide](/wiki/Kernel/Gentoo_Kernel_Configuration_Guide "Kernel/Gentoo Kernel Configuration Guide")

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

To also sign external kernel modules installed by other packages via `linux-mod-r1.eclass`, enable the [modules-sign](https://packages.gentoo.org/useflags/modules-sign) [USE flag/pl](/wiki/USE_flag/pl "USE flag/pl") USE flag globally:

FILE **`/etc/portage/make.conf`** **Enable module signing**

```
USE="modules-sign"
</div>

<div lang="en" dir="ltr" class="mw-content-ltr">
# Optionally, when using custom signing keys.
MODULES_SIGN_KEY="/path/to/kernel_key.pem"
MODULES_SIGN_CERT="/path/to/kernel_key.pem" # Only required if the MODULES_SIGN_KEY does not also contain the certificate
MODULES_SIGN_HASH="sha512" # Defaults to sha512

```

**Informacja**

MODULES\_SIGN\_KEY and MODULES\_SIGN\_CERT may point to different files. For this example, the pem file generated by OpenSSL includes both the key and the accompanying certificate, and thus both variables are set to the same value.

#### Compiling and installing

Now that the kernel is configured, it is time to compile and install it. Exit the configuration and start the compilation process:

`root #` `make && make modules_install`

**Informacja**

It is possible to enable parallel builds using make -j N with N being the integer _number_ of parallel tasks that the build process is allowed to launch. This is similar to the instructions about /etc/portage/make.conf earlier, with the MAKEOPTS variable.

When the kernel has finished compiling, copy the kernel image to /boot/. This is handled by the make install command:

`root #` `make install`

This command will copy the kernel image to /boot. If [sys-kernel/installkernel](https://packages.gentoo.org/packages/sys-kernel/installkernel) is installed it will call /sbin/installkernel instead and delegate the kernel installation. Instead of simply copying the kernel to /boot, [Installkernel](/wiki/Installkernel "Installkernel") installs each kernel with its version number in the file name. Additionally, installkernel provides a framework for automatically accomplishing various tasks relating to kernel installation, such as: generating an [initramfs](/wiki/Initramfs "Initramfs"), building an [Unified Kernel Image](/wiki/Unified_Kernel_Image "Unified Kernel Image"), and updating the [bootloader](/wiki/Bootloader "Bootloader") configuration.

Continue the installation with [Configuring the system](/wiki/Handbook:PPC64/Installation/System/pl "Handbook:PPC64/Installation/System/pl").

[← Installing base system](/wiki/Handbook:PPC64/Installation/Base "Previous part") [Home](/wiki/Handbook:PPC64 "Handbook:PPC64") [Configuring the system →](/wiki/Handbook:PPC64/Installation/System "Next part")