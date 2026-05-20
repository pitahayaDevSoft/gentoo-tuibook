# Kernel

Other languages:

- [Deutsch](/wiki/Handbook:MIPS/Installation/Kernel/de "Handbuch:MIPS/Installation/Kernel (100% translated)")
- [English](/wiki/Handbook:MIPS/Installation/Kernel "Handbook:MIPS/Installation/Kernel (100% translated)")
- [español](/wiki/Handbook:MIPS/Installation/Kernel/es "Manual de Gentoo: MIPS/Instalación/Núcleo (100% translated)")
- [français](/wiki/Handbook:MIPS/Installation/Kernel/fr "Handbook:MIPS/Installation/Kernel/fr (100% translated)")
- [italiano](/wiki/Handbook:MIPS/Installation/Kernel/it "Handbook:MIPS/Installation/Kernel/it (100% translated)")
- [magyar](/wiki/Handbook:MIPS/Installation/Kernel/hu "Handbook:MIPS/Installation/Kernel/hu (100% translated)")
- polski
- [português do Brasil](/wiki/Handbook:MIPS/Installation/Kernel/pt-br "Manual:MIPS/Instalação/Kernel (100% translated)")
- [русский](/wiki/Handbook:MIPS/Installation/Kernel/ru "Handbook:MIPS/Installation/Kernel (100% translated)")
- [தமிழ்](/wiki/Handbook:MIPS/Installation/Kernel/ta "கையேடு:MIPS/நிறுவல்/கர்னல் (100% translated)")
- [中文（中国大陆）‎](/wiki/Handbook:MIPS/Installation/Kernel/zh-cn "手册：MIPS/安装/配置Linux内核 (100% translated)")
- [日本語](/wiki/Handbook:MIPS/Installation/Kernel/ja "ハンドブック:MIPS/インストール/カーネル (100% translated)")
- [한국어](/wiki/Handbook:MIPS/Installation/Kernel/ko "Handbook:MIPS/Installation/Kernel/ko (100% translated)")

[MIPS Handbook](/wiki/Handbook:MIPS/pl "Handbook:MIPS/pl")[Instalacja](/wiki/Handbook:MIPS/Full/Installation/pl "Handbook:MIPS/Full/Installation/pl")[O instalacji](/wiki/Handbook:MIPS/Installation/About/pl "Handbook:MIPS/Installation/About/pl")[Wybór medium instalacyjnego](/wiki/Handbook:MIPS/Installation/Media/pl "Handbook:MIPS/Installation/Media/pl")[Konfiguracja sieci](/wiki/Handbook:MIPS/Installation/Networking/pl "Handbook:MIPS/Installation/Networking/pl")[Przygotowanie dysków](/wiki/Handbook:MIPS/Installation/Disks/pl "Handbook:MIPS/Installation/Disks/pl")[Instalacja etapu 3](/wiki/Handbook:MIPS/Installation/Stage/pl "Handbook:MIPS/Installation/Stage/pl")[Instalacja systemu podstawowego](/wiki/Handbook:MIPS/Installation/Base/pl "Handbook:MIPS/Installation/Base/pl")Konfiguracja jądra[Konfiguracja systemu](/wiki/Handbook:MIPS/Installation/System/pl "Handbook:MIPS/Installation/System/pl")[Instalacja narzędzi](/wiki/Handbook:MIPS/Installation/Tools/pl "Handbook:MIPS/Installation/Tools/pl")[Instalacja systemu rozruchowego](/wiki/Handbook:MIPS/Installation/Bootloader/pl "Handbook:MIPS/Installation/Bootloader/pl")[Finalizacja](/wiki/Handbook:MIPS/Installation/Finalizing/pl "Handbook:MIPS/Installation/Finalizing/pl")[Praca z Gentoo](/wiki/Handbook:MIPS/Full/Working/pl "Handbook:MIPS/Full/Working/pl")[Wstęp do Portage](/wiki/Handbook:MIPS/Working/Portage/pl "Handbook:MIPS/Working/Portage/pl")[Flagi USE](/wiki/Handbook:MIPS/Working/USE/pl "Handbook:MIPS/Working/USE/pl")[Funkcje portage](/wiki/Handbook:MIPS/Working/Features/pl "Handbook:MIPS/Working/Features/pl")[System initscript](/wiki/Handbook:MIPS/Working/Initscripts/pl "Handbook:MIPS/Working/Initscripts/pl")[Zmienne środowiskowe](/wiki/Handbook:MIPS/Working/EnvVar/pl "Handbook:MIPS/Working/EnvVar/pl")[Praca z Portage](/wiki/Handbook:MIPS/Full/Portage/pl "Handbook:MIPS/Full/Portage/pl")[Pliki i katalogi](/wiki/Handbook:MIPS/Portage/Files/pl "Handbook:MIPS/Portage/Files/pl")[Zmienne](/wiki/Handbook:MIPS/Portage/Variables/pl "Handbook:MIPS/Portage/Variables/pl")[Mieszanie działów oprogramowania](/wiki/Handbook:MIPS/Portage/Branches/pl "Handbook:MIPS/Portage/Branches/pl")[Dodatkowe narzędzia](/wiki/Handbook:MIPS/Portage/Tools/pl "Handbook:MIPS/Portage/Tools/pl")[Custom package repository](/wiki/Handbook:MIPS/Portage/CustomTree/pl "Handbook:MIPS/Portage/CustomTree/pl")[Funkcje zaawansowane](/wiki/Handbook:MIPS/Portage/Advanced/pl "Handbook:MIPS/Portage/Advanced/pl")[Konfiguracja sieci](/wiki/Handbook:MIPS/Full/Networking/pl "Handbook:MIPS/Full/Networking/pl")[Zaczynamy](/wiki/Handbook:MIPS/Networking/Introduction/pl "Handbook:MIPS/Networking/Introduction/pl")[Zaawansowana konfiguracja](/wiki/Handbook:MIPS/Networking/Advanced/pl "Handbook:MIPS/Networking/Advanced/pl")[Sieć modularna](/wiki/Handbook:MIPS/Networking/Modular/pl "Handbook:MIPS/Networking/Modular/pl")[Sieć bezprzewodowa](/wiki/Handbook:MIPS/Networking/Wireless/pl "Handbook:MIPS/Networking/Wireless/pl")[Dodawanie funkcjonalności](/wiki/Handbook:MIPS/Networking/Extending/pl "Handbook:MIPS/Networking/Extending/pl")[Dynamiczne zarządzanie](/wiki/Handbook:MIPS/Networking/Dynamic/pl "Handbook:MIPS/Networking/Dynamic/pl")

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
  - [3.1Alternative: Manual configuration](#Alternative:_Manual_configuration)
    - [3.1.1Installing and Configuring the kernel sources](#Installing_and_Configuring_the_kernel_sources)
      - [3.1.1.1Option 2 - Assisted manual process](#Option_2_-_Assisted_manual_process)
      - [3.1.1.2Option 3 - Configuring by hand](#Option_3_-_Configuring_by_hand)
    - [3.1.2Optional: Signed kernel modules](#Optional:_Signed_kernel_modules)
  - [3.2Preparing the configuration](#Preparing_the_configuration)
  - [3.3Customizing the configuration](#Customizing_the_configuration)
  - [3.4Compiling and installing](#Compiling_and_installing)

## Installing firmware and/or microcode

### Firmware

#### Suggested: Linux Firmware

On many systems, non-FOSS firmware is required for certain hardware to function. The [sys-kernel/linux-firmware](https://packages.gentoo.org/packages/sys-kernel/linux-firmware) package contains firmware for many, but not all, devices.

**Wskazówka**

Most wireless cards and GPUs require firmware to function.

`root #` `emerge --ask sys-kernel/linux-firmware`

**Informacja**

Installing certain firmware packages often requires accepting the associated firmware licenses. If necessary, visit the [license handling section](/wiki/Handbook:MIPS/Working/Portage/pl#Licenses "Handbook:MIPS/Working/Portage/pl") of the Handbook for help on accepting licenses.

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

[Manual approach](/wiki/Handbook:MIPS/Installation/Kernel/pl#Alternative:_Manual_configuration "Handbook:MIPS/Installation/Kernel/pl")New kernel sources are installed via the system package manager. The kernel is manually configured, built, and installed using the eselect kernel and a slew of make commands. Future kernel updates repeat the manual process of configuring, building, and installing the kernel files. This is the most involved process, but offers maximum control over the kernel update process.

The core around which all distributions are built is the Linux kernel. It is the layer between the user's programs and the system hardware. Although the handbook provides its users several possible kernel sources, a more comprehensive listing with more detailed descriptions is available at the [kernel packages page](/wiki/Kernel/Packages "Kernel/Packages").

**Wskazówka**

Kernel installation tasks such as copying the kernel image to /boot or the [EFI System Partition](/wiki/EFI_System_Partition "EFI System Partition"), generating an [initramfs](/wiki/Initramfs "Initramfs") and/or [Unified Kernel Image](/wiki/Unified_Kernel_Image "Unified Kernel Image"), updating bootloader configuration, can be automated with [installkernel](/wiki/Installkernel "Installkernel"). Users may wish to configure and install [sys-kernel/installkernel](https://packages.gentoo.org/packages/sys-kernel/installkernel) before proceeding. See the [Kernel installation section below](/wiki/Handbook:MIPS/Installation/Kernel#Kernel_installation.2Fpl "Handbook:MIPS/Installation/Kernel") for more more information.

### Alternative: Manual configuration

#### Installing and Configuring the kernel sources

When manually compiling the kernel for mips-based systems, Gentoo recommends the [sys-kernel/mips-sources](https://packages.gentoo.org/packages/sys-kernel/mips-sources) package.

Choose an appropriate kernel source and install it using emerge:

`root #` `emerge --ask sys-kernel/mips-sources`

This will install the Linux kernel sources in /usr/src/ using the specific kernel version in the path. It will not create a symbolic link by itself without the [symlink](https://packages.gentoo.org/useflags/symlink) [USE flag/pl](/wiki/USE_flag/pl "USE flag/pl") USE flag being enabled on the chosen kernel sources package.

It is conventional for a /usr/src/linux symlink to be maintained, such that it refers to whichever sources correspond with the currently running kernel. However, this symbolic link will not be created by default. An easy way to create the symbolic link is to utilize eselect's kernel module.

For further information regarding the purpose of the symlink, and how to manage it, please refer to [Kernel/Upgrade](/wiki/Kernel/Upgrade "Kernel/Upgrade").

First, list all installed kernels:

`root #` `eselect kernel list`

```
Available kernel symlink targets:
  [1]   linux-6.19.1-gentoo

```

In order to create a symbolic link called linux, use:

`root #` `eselect kernel set 1`

`root #` `ls -l /usr/src/linux`

```
lrwxrwxrwx    1 root   root    12 Oct 13 11:04 /usr/src/linux -> linux-6.19.1-gentoo

```

Manually configuring a kernel is commonly seen as one of the most difficult procedures a system administrator has to perform. Nothing is less true - after configuring a few kernels no one remembers that it was difficult! There are two ways for a Gentoo user to manage a manual kernel system, both of which are listed below:

**Ważne**

Only one selection is required in the following subsection, if unsure of which to use then go with the first listed for now. It's always possible to switch at a later date if required.

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

Now it's time to decide if modules signing is required in the steps listed in [here](/wiki/Handbook:MIPS/Installation/Kernel#Optional:_Signed_kernel_modules_2.2Fpl "Handbook:MIPS/Installation/Kernel")

If not, proceed to building described [here](/wiki/Handbook:MIPS/Installation/Kernel#Compiling_and_installing.2Fpl "Handbook:MIPS/Installation/Kernel")

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

### Preparing the configuration

**Ważne**

On the Origin 200/2000, Indigo2 Impact (R10000), Octane/Octane2 and O2, a 64-bit kernel is required to boot these systems. For these machines, emerge [sys-devel/kgcc64](https://packages.gentoo.org/packages/sys-devel/kgcc64) to create a cross-compiler for building 64-bit kernels.

Many of the systems supported have sample default .config files hiding in amongst the kernel source. Not all systems have configs distributed in this way. Those that do, can be configured using the commands mentioned in the table below.

System
Configure command
Cobalt Servers
make cobalt\_defconfigIndy, Indigo2 (R4k), Challenge S
make ip22\_defconfigOrigin 200/2000
make ip27\_defconfigIndigo2 Impact (R10k)
make ip28\_defconfigO2
make ip32\_defconfig

All of the Gentoo installation images provide a kernel config option as part of the image itself, accessible as /proc/config.gz. This may be used in many cases. It is best though if the kernel source matches closely the kernel that is currently running. To extract it, simply run it through zcat as shown below.

`root #` `zcat /proc/config.gz > .config`

**Ważne**

This kernel config is set up for a netboot image. That is, it will expect to find a root filesystem image somewhere nearby, either as a directory for initramfs, or a loopback device for initrd. When executing make menuconfig, don't forget to go into General Setup and disable the options for initramfs.

### Customizing the configuration

Once a configuration is found, download it into the kernel source directory, and rename it to .config. From there, run make oldconfig to bring everything up to date according to the instructions above, and customize the configuration before compiling.

`root #` `cd /usr/src/linux
`

`root #` `cp /path/to/example-config .config
`

`root #` `make oldconfig`

Just press the `ENTER` (or `Return`) key at each prompt to accept the defaults for now ...

`root #` `make menuconfig`

**Ważne**

In the Kernel Hacking section, there is an option named "Are You Using A Cross Compiler?". This tells the kernel Makefiles to prepend "mips-linux-" (or mipsel-linux ... etc) to gcc and as commands when compiling the kernel. This should be turned off, even if cross-compiling. Instead, if a cross-compiler needs to be called, specify the prefix using the CROSS\_COMPILE variable as shown in the next section.

**Ważne**

There is a known issue with JFS and ALSA on Octane systems where the ALSA fails to work. Given the experimental nature of JFS on MIPS, it is recommended that people avoid using JFS for the time being.

### Compiling and installing

Now that the kernel is configured, it is time to compile and install it. Exit the configuration and start the compilation process:

**Informacja**

On 64-bit machines, specify CROSS\_COMPILE=mips64-unknown-linux-gnu- (or mips64el-... if on a little-endian system) to use the 64-bit compiler.

To compile natively:

`root #` `make vmlinux modules modules_install`

Cross-compiling on target machine, adjust the mips64-unknown-linux-gnu- accordingly:

`root #` `make vmlinux modules modules_install CROSS_COMPILE=mips64-unknown-linux-gnu-`

When compiling on another machine, such as an x86 box, use the following commands to compile the kernel & install modules into a specific directory to be transferred to the target machine.

`root #` `make vmlinux modules CROSS_COMPILE=mips64-unknown-linux-gnu-
`

`root #` `make modules_install INSTALL_MOD_PATH=/somewhere`

**Ważne**

When compiling a 64-bit kernel for the Indy, Indigo2 (R4k), Challenge S and O2, use the vmlinux.32 target instead of vmlinux. Otherwise, the machine will not be able to boot. This is to work around the PROM not understanding the ELF64 format.

`root #` `make vmlinux.32`

**Informacja**

It is possible to enable parallel builds using make -jN with N being the number of parallel tasks that the build process is allowed to launch. This is similar to the instructions about /etc/portage/make.conf earlier, with the MAKEOPTS variable.

The above will create vmlinux.32, which is the final kernel.

When the kernel has finished compiling, copy the kernel image to /boot/.

**Informacja**

On Cobalt servers, the bootloader will expect to see a compressed kernel image. Remember to gzip -9 the file once it is in /boot/.

`root #` `cp vmlinux /boot/kernel-6.19.1-gentoo`

For Cobalt servers, compress the kernel image:

`root #` `gzip -9v /boot/kernel-6.19.1-gentoo`

Continue the installation with [Configuring the system](/wiki/Handbook:MIPS/Installation/System/pl "Handbook:MIPS/Installation/System/pl").

[← Installing base system](/wiki/Handbook:MIPS/Installation/Base "Previous part") [Home](/wiki/Handbook:MIPS "Handbook:MIPS") [Configuring the system →](/wiki/Handbook:MIPS/Installation/System "Next part")