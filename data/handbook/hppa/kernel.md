# Kernel

Other languages:

- [Deutsch](/wiki/Handbook:HPPA/Installation/Kernel/de "Handbuch:HPPA/Installation/Kernel (100% translated)")
- English
- [español](/wiki/Handbook:HPPA/Installation/Kernel/es "Manual de Gentoo: HPPA/Instalación/Núcleo (100% translated)")
- [français](/wiki/Handbook:HPPA/Installation/Kernel/fr "Handbook:HPPA/Installation/Kernel/fr (100% translated)")
- [italiano](/wiki/Handbook:HPPA/Installation/Kernel/it "Handbook:HPPA/Installation/Kernel/it (100% translated)")
- [magyar](/wiki/Handbook:HPPA/Installation/Kernel/hu "Handbook:HPPA/Installation/Kernel/hu (100% translated)")
- [polski](/wiki/Handbook:HPPA/Installation/Kernel/pl "Handbook:HPPA/Installation/Kernel (100% translated)")
- [português do Brasil](/wiki/Handbook:HPPA/Installation/Kernel/pt-br "Manual:HPPA/Instalação/Kernel (100% translated)")
- [русский](/wiki/Handbook:HPPA/Installation/Kernel/ru "Handbook:HPPA/Installation/Kernel (100% translated)")
- [தமிழ்](/wiki/Handbook:HPPA/Installation/Kernel/ta "கையேடு:HPPA/நிறுவல்/கர்னல் (100% translated)")
- [中文（中国大陆）‎](/wiki/Handbook:HPPA/Installation/Kernel/zh-cn "手册：HPPA/安装/配置Linux内核 (100% translated)")
- [日本語](/wiki/Handbook:HPPA/Installation/Kernel/ja "ハンドブック:HPPA/インストール/カーネル (100% translated)")
- [한국어](/wiki/Handbook:HPPA/Installation/Kernel/ko "Handbook:HPPA/Installation/Kernel/ko (100% translated)")

[HPPA Handbook](/wiki/Handbook:HPPA "Handbook:HPPA")[Installation](/wiki/Handbook:HPPA/Full/Installation "Handbook:HPPA/Full/Installation")[About the installation](/wiki/Handbook:HPPA/Installation/About "Handbook:HPPA/Installation/About")[Choosing the media](/wiki/Handbook:HPPA/Installation/Media "Handbook:HPPA/Installation/Media")[Configuring the network](/wiki/Handbook:HPPA/Installation/Networking "Handbook:HPPA/Installation/Networking")[Preparing the disks](/wiki/Handbook:HPPA/Installation/Disks "Handbook:HPPA/Installation/Disks")[The stage file](/wiki/Handbook:HPPA/Installation/Stage "Handbook:HPPA/Installation/Stage")[Installing base system](/wiki/Handbook:HPPA/Installation/Base "Handbook:HPPA/Installation/Base")Configuring the kernel[Configuring the system](/wiki/Handbook:HPPA/Installation/System "Handbook:HPPA/Installation/System")[Installing tools](/wiki/Handbook:HPPA/Installation/Tools "Handbook:HPPA/Installation/Tools")[Configuring the bootloader](/wiki/Handbook:HPPA/Installation/Bootloader "Handbook:HPPA/Installation/Bootloader")[Finalizing](/wiki/Handbook:HPPA/Installation/Finalizing "Handbook:HPPA/Installation/Finalizing")[Working with Gentoo](/wiki/Handbook:HPPA/Full/Working "Handbook:HPPA/Full/Working")[Portage introduction](/wiki/Handbook:HPPA/Working/Portage "Handbook:HPPA/Working/Portage")[USE flags](/wiki/Handbook:HPPA/Working/USE "Handbook:HPPA/Working/USE")[Portage features](/wiki/Handbook:HPPA/Working/Features "Handbook:HPPA/Working/Features")[Initscript system](/wiki/Handbook:HPPA/Working/Initscripts "Handbook:HPPA/Working/Initscripts")[Environment variables](/wiki/Handbook:HPPA/Working/EnvVar "Handbook:HPPA/Working/EnvVar")[Working with Portage](/wiki/Handbook:HPPA/Full/Portage "Handbook:HPPA/Full/Portage")[Files and directories](/wiki/Handbook:HPPA/Portage/Files "Handbook:HPPA/Portage/Files")[Variables](/wiki/Handbook:HPPA/Portage/Variables "Handbook:HPPA/Portage/Variables")[Mixing software branches](/wiki/Handbook:HPPA/Portage/Branches "Handbook:HPPA/Portage/Branches")[Additional tools](/wiki/Handbook:HPPA/Portage/Tools "Handbook:HPPA/Portage/Tools")[Custom package repository](/wiki/Handbook:HPPA/Portage/CustomTree "Handbook:HPPA/Portage/CustomTree")[Advanced features](/wiki/Handbook:HPPA/Portage/Advanced "Handbook:HPPA/Portage/Advanced")[OpenRC network configuration](/wiki/Handbook:HPPA/Full/Networking "Handbook:HPPA/Full/Networking")[Getting started](/wiki/Handbook:HPPA/Networking/Introduction "Handbook:HPPA/Networking/Introduction")[Advanced configuration](/wiki/Handbook:HPPA/Networking/Advanced "Handbook:HPPA/Networking/Advanced")[Modular networking](/wiki/Handbook:HPPA/Networking/Modular "Handbook:HPPA/Networking/Modular")[Wireless](/wiki/Handbook:HPPA/Networking/Wireless "Handbook:HPPA/Networking/Wireless")[Adding functionality](/wiki/Handbook:HPPA/Networking/Extending "Handbook:HPPA/Networking/Extending")[Dynamic management](/wiki/Handbook:HPPA/Networking/Dynamic "Handbook:HPPA/Networking/Dynamic")

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
  - [3.2Compiling and installing](#Compiling_and_installing)

## Installing firmware and/or microcode\[ [edit](/index.php?title=Handbook:Parts/Installation/Kernel&action=edit&section=T-1 "Edit section: ")\]

### Firmware\[ [edit](/index.php?title=Handbook:Parts/Installation/Kernel&action=edit&section=T-2 "Edit section: ")\]

#### Suggested: Linux Firmware\[ [edit](/index.php?title=Handbook:Parts/Installation/Kernel&action=edit&section=T-3 "Edit section: ")\]

On many systems, non-FOSS firmware is required for certain hardware to function. The [sys-kernel/linux-firmware](https://packages.gentoo.org/packages/sys-kernel/linux-firmware) package contains firmware for many, but not all, devices.

**Tip**

Most wireless cards and GPUs require firmware to function.

`root #` `emerge --ask sys-kernel/linux-firmware`

**Note**

Installing certain firmware packages often requires accepting the associated firmware licenses. If necessary, visit the [license handling section](/wiki/Handbook:HPPA/Working/Portage#Licenses "Handbook:HPPA/Working/Portage") of the Handbook for help on accepting licenses.

##### Firmware Loading\[ [edit](/index.php?title=Handbook:Parts/Installation/Kernel&action=edit&section=T-4 "Edit section: ")\]

Firmware files are typically loaded when the associated kernel module is loaded. This means the firmware must be built into the kernel using **CONFIG\_EXTRA\_FIRMWARE** if the kernel module is set to _Y_ instead of _M_. In most cases, building-in a module which required firmware can complicate or break loading.

## sys-kernel/installkernel\[ [edit](/index.php?title=Handbook:Parts/Installation/Kernel&action=edit&section=T-6 "Edit section: ")\]

[Installkernel](/wiki/Installkernel "Installkernel") may be used to automate the kernel installation, [initramfs](/wiki/Initramfs "Initramfs") generation, [unified kernel image](/wiki/Unified_kernel_image "Unified kernel image") generation and bootloader configuration, among other things. [sys-kernel/installkernel](https://packages.gentoo.org/packages/sys-kernel/installkernel) implements two paths of achieving this: the traditional installkernel originating from Debian and [systemd](/wiki/Systemd "Systemd")'s kernel-install. Which one to choose depends, among other things, on the system's bootloader. By default, systemd's kernel-install is used on systemd profiles, while the traditional installkernel is the default for other profiles.

### Bootloader\[ [edit](/index.php?title=Handbook:Parts/Installation/Kernel&action=edit&section=T-7 "Edit section: ")\]

Now is the time to think about which bootloader the user wants for the system.

**Important**

Only one selection is required in the following subsection, if unsure of which to use then go with the first listed for now. It's always possible to switch at a later date if required.

#### GRUB\[ [edit](/index.php?title=Handbook:Parts/Installation/Kernel&action=edit&section=T-8 "Edit section: ")\]

Users of GRUB can use either systemd's kernel-install or the traditional Debian installkernel. The [systemd](https://packages.gentoo.org/useflags/systemd) [USE flag](/wiki/USE_flag "USE flag") USE flag switches between these implementations. To automatically run grub-mkconfig when installing the kernel, enable the [grub](https://packages.gentoo.org/useflags/grub) [USE flag](/wiki/USE_flag "USE flag") [USE flag](/wiki/USE_flag "USE flag").

**Note**

GRUB requires kernels to be installed to /boot.

FILE **`/etc/portage/package.use/installkernel`**

```
sys-kernel/installkernel grub

```

`root #` `emerge --ask sys-kernel/installkernel`

**Note**

systemd-boot requires kernels to be installed to /efi.

**Note**

When [app-emulation/virt-firmware](https://packages.gentoo.org/packages/app-emulation/virt-firmware) is used to configure the UEFI ensure that the kernel-bootcfg-boot-successful service is enabled before attempting to install the kernel. This implementation of EFIstub booting is the default for systemd systems.

`root #` `systemctl enable kernel-bootcfg-boot-successful.service`

**Note**

EFIstub requires kernels to be installed to /efi.

#### Traditional layout, other bootloaders (e.g. (e)lilo, syslinux, etc.)\[ [edit](/index.php?title=Handbook:Parts/Installation/Kernel&action=edit&section=T-11 "Edit section: ")\]

The traditional /boot layout (for e.g. (e)LILO, syslinux, etc.) is used by default if the [grub](https://packages.gentoo.org/useflags/grub) [USE flag](/wiki/USE_flag "USE flag"), [systemd-boot](https://packages.gentoo.org/useflags/systemd-boot) [USE flag](/wiki/USE_flag "USE flag"), [efistub](https://packages.gentoo.org/useflags/efistub) [USE flag](/wiki/USE_flag "USE flag") and [uki](https://packages.gentoo.org/useflags/uki) [USE flag](/wiki/USE_flag "USE flag") USE flags are **not** enabled. No further action is required.

### Initramfs\[ [edit](/index.php?title=Handbook:Parts/Installation/Kernel&action=edit&section=T-12 "Edit section: ")\]

An **init** ial **ram**-based **f** ile **s** ystem, or [initramfs](/wiki/Initramfs "Initramfs"), may be required for a system to boot. A wide of variety of cases may necessitate one, but common cases include:

- Kernels where storage/filesystem drivers are modules.
- Layouts with /usr/ or /var/ on separate partitions.
- Encrypted root filesystems.

**Tip**

[Distribution kernels](/wiki/Project:Distribution_Kernel "Project:Distribution Kernel") are designed to be used with an initramfs, as many storage and filesystem drivers are built as modules.

In addition to mounting the root filesystem, an initramfs may also perform other tasks such as:

- Running **f** ile **s** ystem **c** onsistency chec **k** fsck, a tool to check and repair consistency of a file system in such events of uncleanly shutdown a system.
- Providing a recovery environment in the event of late-boot failures.

[Installkernel](/wiki/Installkernel "Installkernel") can automatically generate an initramfs when installing the kernel if the [dracut](https://packages.gentoo.org/useflags/dracut) [USE flag](/wiki/USE_flag "USE flag") or [ugrd](https://packages.gentoo.org/useflags/ugrd) [USE flag](/wiki/USE_flag "USE flag") USE flag is enabled:

FILE **`/etc/portage/package.use/installkernel`**

```
sys-kernel/installkernel dracut

```

#### Chroot detection\[ [edit](/index.php?title=Handbook:Parts/Installation/Kernel&action=edit&section=T-13 "Edit section: ")\]

Bootloaders such as [systemd-boot](/wiki/Systemd/systemd-boot "Systemd/systemd-boot") and [EFI stub](/wiki/EFI_stub "EFI stub") use the kernel arguments of the running system as set in /proc/cmdline by default. Because of the wide range of ways Gentoo can be installed users will randomly get tripped up by this.

To help solve any issues this may cause, [sys-kernel/installkernel](https://packages.gentoo.org/packages/sys-kernel/installkernel) will check if it is running in a chroot and fail if the kernel command line is not explicitly configured. Otherwise the bootloader would use the install media's boot arguments which would lead to boot failure.

One way to satisfy [sys-kernel/installkernel](https://packages.gentoo.org/packages/sys-kernel/installkernel) is by using Dracut's config file to set the root partition UUID as shown below, or alternatively for more information on what this check helps with and different ways to configure it, see [Installkernel#Install\_chroot\_check](/wiki/Installkernel#Install_chroot_check "Installkernel").

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

## Kernel selection\[ [edit](/index.php?title=Handbook:Parts/Installation/Kernel&action=edit&section=T-17 "Edit section: ")\]

It can be a wise move to use the dist-kernel on the first boot as it provides a very simple method to rule out system issues and kernel config issues. Always having a known working kernel to fallback on can speed up debugging and alleviate anxiety when updating that your system will no longer boot.

A common misconception is that a manually compiled kernel will use a lot less RAM than a pre configured distribution kernel. Due to the modular nature of the Linux kernel, only what is needed by the system is loaded and in most cases resulting in similar memory usage.

**Important**

Only one selection is required in the following subsection, if unsure of which to use then go with the first listed for now. It's always possible to switch at a later date if required.

Ranked from least involved to most involved:

[Manual approach](/wiki/Handbook:HPPA/Installation/Kernel#Alternative:_Manual_configuration "Handbook:HPPA/Installation/Kernel")New kernel sources are installed via the system package manager. The kernel is manually configured, built, and installed using the eselect kernel and a slew of make commands. Future kernel updates repeat the manual process of configuring, building, and installing the kernel files. This is the most involved process, but offers maximum control over the kernel update process.

The core around which all distributions are built is the Linux kernel. It is the layer between the user's programs and the system hardware. Although the handbook provides its users several possible kernel sources, a more comprehensive listing with more detailed descriptions is available at the [kernel packages page](/wiki/Kernel/Packages "Kernel/Packages").

**Tip**

Kernel installation tasks such as copying the kernel image to /boot or the [EFI System Partition](/wiki/EFI_System_Partition "EFI System Partition"), generating an [initramfs](/wiki/Initramfs "Initramfs") and/or [Unified Kernel Image](/wiki/Unified_Kernel_Image "Unified Kernel Image"), updating bootloader configuration, can be automated with [installkernel](/wiki/Installkernel "Installkernel"). Users may wish to configure and install [sys-kernel/installkernel](https://packages.gentoo.org/packages/sys-kernel/installkernel) before proceeding. See the [Kernel installation section below](/wiki/Handbook:HPPA/Installation/Kernel#Kernel_installation "Handbook:HPPA/Installation/Kernel") for more more information.

### Alternative: Manual configuration\[ [edit](/index.php?title=Handbook:Parts/Installation/Kernel&action=edit&section=T-18 "Edit section: ")\]

#### Installing and Configuring the kernel sources\[ [edit](/index.php?title=Handbook:Parts/Installation/Kernel&action=edit&section=T-19 "Edit section: ")\]

When manually compiling the kernel for hppa-based systems, Gentoo recommends the [sys-kernel/gentoo-sources](https://packages.gentoo.org/packages/sys-kernel/gentoo-sources) package.

Choose an appropriate kernel source and install it using emerge:

`root #` `emerge --ask sys-kernel/gentoo-sources`

This will install the Linux kernel sources in /usr/src/ using the specific kernel version in the path. It will not create a symbolic link by itself without the [symlink](https://packages.gentoo.org/useflags/symlink) [USE flag](/wiki/USE_flag "USE flag") USE flag being enabled on the chosen kernel sources package.

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

**Important**

Only one selection is required in the following subsection, if unsure of which to use then go with the first listed for now. It's always possible to switch at a later date if required.

##### Option 2 - Assisted manual process\[ [edit](/index.php?title=Handbook:Parts/Installation/Kernel&action=edit&section=T-21 "Edit section: ")\]

This method allows a user to have full control of how their kernel is built with as minimal help from outside tools as they wish. Some could consider this as making it hard for the sake of it.

However, with this choice one thing is true: it is vital to know the system when a kernel is configured manually. Most information can be gathered by emerging [sys-apps/pciutils](https://packages.gentoo.org/packages/sys-apps/pciutils) which contains the lspci command:

`root #` `emerge --ask sys-apps/pciutils`

**Note**

Inside the chroot, it is safe to ignore any pcilib warnings (like _pcilib: cannot open /sys/bus/pci/devices_) that lspci might throw out.

Another source of system information is to run lsmod to see what kernel modules the installation CD uses as it might provide a nice hint on what to enable.

Now go to the kernel source directory.

`root #` `cd /usr/src/linux
`

**Tip**

To view the full list of make arguments available for the kernel, run `make help`.

The kernel has a method of autodetecting the modules currently being used on the installcd which will give a great starting point to allow a user to configure their own. This can be called by using:

`root #` `make localmodconfig`

Manually configuring should not be needed at this point, but if a user wish to check:

`root #` `make nconfig`

Now it's time to decide if modules signing is required in the steps listed in [here](/wiki/Handbook:HPPA/Installation/Kernel#Optional:_Signed_kernel_modules_2 "Handbook:HPPA/Installation/Kernel")

If not, proceed to building described [here](/wiki/Handbook:HPPA/Installation/Kernel#Compiling_and_installing "Handbook:HPPA/Installation/Kernel")

##### Option 3 - Configuring by hand\[ [edit](/index.php?title=Handbook:Parts/Installation/Kernel&action=edit&section=T-22 "Edit section: ")\]

The Linux kernel configuration has many, many sections and as configuring a kernel by hand is rarely needed thanks to the two tools listed above. The steps to do it by hand are now included at [Kernel/Gentoo\_Kernel\_Configuration\_Guide](/wiki/Kernel/Gentoo_Kernel_Configuration_Guide "Kernel/Gentoo Kernel Configuration Guide")

#### Optional: Signed kernel modules\[ [edit](/index.php?title=Handbook:Parts/Installation/Kernel&action=edit&section=T-23 "Edit section: ")\]

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

To also sign external kernel modules installed by other packages via `linux-mod-r1.eclass`, enable the [modules-sign](https://packages.gentoo.org/useflags/modules-sign) [USE flag](/wiki/USE_flag "USE flag") USE flag globally:

FILE **`/etc/portage/make.conf`** **Enable module signing**

```
USE="modules-sign"

# Optionally, when using custom signing keys.
MODULES_SIGN_KEY="/path/to/kernel_key.pem"
MODULES_SIGN_CERT="/path/to/kernel_key.pem" # Only required if the MODULES_SIGN_KEY does not also contain the certificate
MODULES_SIGN_HASH="sha512" # Defaults to sha512

```

**Note**

MODULES\_SIGN\_KEY and MODULES\_SIGN\_CERT may point to different files. For this example, the pem file generated by OpenSSL includes both the key and the accompanying certificate, and thus both variables are set to the same value.

### Compiling and installing\[ [edit](/index.php?title=Handbook:HPPA/Blocks/Kernel&action=edit&section=T-1 "Edit section: ")\]

**Important**

To compile a 64-bit kernel, first emerge [sys-devel/kgcc64](https://packages.gentoo.org/packages/sys-devel/kgcc64). Running a 64-bit kernel used to be discouraged but should now be safe. If in doubt, only run a 64-bit kernel if the system has more than 4GB of RAM or if the server requires it, i.e. on the A500.

Now that the kernel is configured, it is time to compile and install it. Exit the configuration and start the compilation process:

`root #` `make && make modules_install`

If building a 64-bit kernel, do this instead (it's necessary even for native builds, see [here](https://www.spinics.net/lists/kernel/msg3987729.html)):

`root #` `CROSS_COMPILE=hppa64-unknown-linux-gnu- make && CROSS_COMPILE=hppa64-unknown-linux-gnu- make modules_install`

**Note**

It is possible to enable parallel builds using `make -jX` with X being the number of parallel tasks that the build process is allowed to launch. This is similar to the instructions about /etc/portage/make.conf earlier, with the `MAKEOPTS` variable.

When the kernel has finished compiling, copy the kernel image to /boot/. Use whatever name is appropriate for the kernel choice and remember it as it will be needed later on when configuring the bootloader. Remember to replace kernel-6.19.1-gentoo with the name and version of the installed kernel.

`root #` `cp vmlinux /boot/kernel-6.19.1-gentoo`

Continue the installation with [Configuring the system](/wiki/Handbook:HPPA/Installation/System "Handbook:HPPA/Installation/System").

[← Installing base system](/wiki/Handbook:HPPA/Installation/Base "Previous part") [Home](/wiki/Handbook:HPPA "Handbook:HPPA") [Configuring the system →](/wiki/Handbook:HPPA/Installation/System "Next part")