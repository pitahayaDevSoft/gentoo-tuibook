# Kernel

Other languages:

- [Deutsch](/wiki/Handbook:X86/Installation/Kernel/de "Handbuch:X86/Installation/Kernel (100% translated)")
- English
- [español](/wiki/Handbook:X86/Installation/Kernel/es "Manual de Gentoo: X86/Instalación/Núcleo (100% translated)")
- [français](/wiki/Handbook:X86/Installation/Kernel/fr "Handbook:X86/Installation/Kernel/fr (100% translated)")
- [italiano](/wiki/Handbook:X86/Installation/Kernel/it "Handbook:X86/Installation/Kernel (100% translated)")
- [magyar](/wiki/Handbook:X86/Installation/Kernel/hu "Handbook:X86/Installation/Kernel/hu (100% translated)")
- [polski](/wiki/Handbook:X86/Installation/Kernel/pl "Handbook:X86/Installation/Kernel (100% translated)")
- [português do Brasil](/wiki/Handbook:X86/Installation/Kernel/pt-br "Manual:X86/Instalação/Kernel (100% translated)")
- [čeština](/wiki/Handbook:X86/Installation/Kernel/cs "Handbook:X86/Installation/Kernel/cs (50% translated)")
- [русский](/wiki/Handbook:X86/Installation/Kernel/ru "Handbook:X86/Installation/Kernel (100% translated)")
- [தமிழ்](/wiki/Handbook:X86/Installation/Kernel/ta "கையேடு:X86/நிறுவல்/கர்னல் (100% translated)")
- [中文（中国大陆）‎](/wiki/Handbook:X86/Installation/Kernel/zh-cn "手册：X86/安装/配置Linux内核 (100% translated)")
- [日本語](/wiki/Handbook:X86/Installation/Kernel/ja "ハンドブック:X86/インストール/カーネルの設定 (100% translated)")
- [한국어](/wiki/Handbook:X86/Installation/Kernel/ko "Handbook:X86/Installation/Kernel/ko (100% translated)")

[X86 Handbook](/wiki/Handbook:X86 "Handbook:X86")[Installation](/wiki/Handbook:X86/Full/Installation "Handbook:X86/Full/Installation")[About the installation](/wiki/Handbook:X86/Installation/About "Handbook:X86/Installation/About")[Choosing the media](/wiki/Handbook:X86/Installation/Media "Handbook:X86/Installation/Media")[Configuring the network](/wiki/Handbook:X86/Installation/Networking "Handbook:X86/Installation/Networking")[Preparing the disks](/wiki/Handbook:X86/Installation/Disks "Handbook:X86/Installation/Disks")[The stage file](/wiki/Handbook:X86/Installation/Stage "Handbook:X86/Installation/Stage")[Installing base system](/wiki/Handbook:X86/Installation/Base "Handbook:X86/Installation/Base")Configuring the kernel[Configuring the system](/wiki/Handbook:X86/Installation/System "Handbook:X86/Installation/System")[Installing tools](/wiki/Handbook:X86/Installation/Tools "Handbook:X86/Installation/Tools")[Configuring the bootloader](/wiki/Handbook:X86/Installation/Bootloader "Handbook:X86/Installation/Bootloader")[Finalizing](/wiki/Handbook:X86/Installation/Finalizing "Handbook:X86/Installation/Finalizing")[Working with Gentoo](/wiki/Handbook:X86/Full/Working "Handbook:X86/Full/Working")[Portage introduction](/wiki/Handbook:X86/Working/Portage "Handbook:X86/Working/Portage")[USE flags](/wiki/Handbook:X86/Working/USE "Handbook:X86/Working/USE")[Portage features](/wiki/Handbook:X86/Working/Features "Handbook:X86/Working/Features")[Initscript system](/wiki/Handbook:X86/Working/Initscripts "Handbook:X86/Working/Initscripts")[Environment variables](/wiki/Handbook:X86/Working/EnvVar "Handbook:X86/Working/EnvVar")[Working with Portage](/wiki/Handbook:X86/Full/Portage "Handbook:X86/Full/Portage")[Files and directories](/wiki/Handbook:X86/Portage/Files "Handbook:X86/Portage/Files")[Variables](/wiki/Handbook:X86/Portage/Variables "Handbook:X86/Portage/Variables")[Mixing software branches](/wiki/Handbook:X86/Portage/Branches "Handbook:X86/Portage/Branches")[Additional tools](/wiki/Handbook:X86/Portage/Tools "Handbook:X86/Portage/Tools")[Custom package repository](/wiki/Handbook:X86/Portage/CustomTree "Handbook:X86/Portage/CustomTree")[Advanced features](/wiki/Handbook:X86/Portage/Advanced "Handbook:X86/Portage/Advanced")[OpenRC network configuration](/wiki/Handbook:X86/Full/Networking "Handbook:X86/Full/Networking")[Getting started](/wiki/Handbook:X86/Networking/Introduction "Handbook:X86/Networking/Introduction")[Advanced configuration](/wiki/Handbook:X86/Networking/Advanced "Handbook:X86/Networking/Advanced")[Modular networking](/wiki/Handbook:X86/Networking/Modular "Handbook:X86/Networking/Modular")[Wireless](/wiki/Handbook:X86/Networking/Wireless "Handbook:X86/Networking/Wireless")[Adding functionality](/wiki/Handbook:X86/Networking/Extending "Handbook:X86/Networking/Extending")[Dynamic management](/wiki/Handbook:X86/Networking/Dynamic "Handbook:X86/Networking/Dynamic")

## Contents

- [1Installing firmware and/or microcode](#Installing_firmware_and.2For_microcode)
  - [1.1Firmware](#Firmware)
    - [1.1.1Suggested: Linux Firmware](#Suggested:_Linux_Firmware)
      - [1.1.1.1Firmware Loading](#Firmware_Loading)
  - [1.2Suggested: Microcode](#Suggested:_Microcode)
- [2sys-kernel/installkernel](#sys-kernel.2Finstallkernel)
  - [2.1Bootloader](#Bootloader)
    - [2.1.1GRUB](#GRUB)
    - [2.1.2systemd-boot](#systemd-boot)
    - [2.1.3EFI stub](#EFI_stub)
    - [2.1.4Traditional layout, other bootloaders (e.g. (e)lilo, syslinux, etc.)](#Traditional_layout.2C_other_bootloaders_.28e.g._.28e.29lilo.2C_syslinux.2C_etc..29)
  - [2.2Initramfs](#Initramfs)
    - [2.2.1Chroot detection](#Chroot_detection)
  - [2.3Optional: Unified Kernel Image](#Optional:_Unified_Kernel_Image)
    - [2.3.1Generic Unified Kernel Image (systemd only)](#Generic_Unified_Kernel_Image_.28systemd_only.29)
    - [2.3.2Secure Boot](#Secure_Boot)
- [3Kernel selection](#Kernel_selection)
  - [3.1Distribution kernels](#Distribution_kernels)
    - [3.1.1Optional: Signed kernel modules](#Optional:_Signed_kernel_modules)
    - [3.1.2Optional: Signing the kernel image (Secure Boot)](#Optional:_Signing_the_kernel_image_.28Secure_Boot.29)
    - [3.1.3Installing a distribution kernel](#Installing_a_distribution_kernel)
    - [3.1.4Upgrading and cleaning up](#Upgrading_and_cleaning_up)
    - [3.1.5Post-install/upgrade tasks](#Post-install.2Fupgrade_tasks)
      - [3.1.5.1Manually rebuilding the initramfs or Unified Kernel Image](#Manually_rebuilding_the_initramfs_or_Unified_Kernel_Image)
  - [3.2Alternative: Manual configuration](#Alternative:_Manual_configuration)
    - [3.2.1Installing and Configuring the kernel sources](#Installing_and_Configuring_the_kernel_sources)
      - [3.2.1.1Option 1 - Modprobed-db process](#Option_1_-_Modprobed-db_process)
      - [3.2.1.2Option 2 - Assisted manual process](#Option_2_-_Assisted_manual_process)
      - [3.2.1.3Option 3 - Configuring by hand](#Option_3_-_Configuring_by_hand)
    - [3.2.2Optional: Signed kernel modules](#Optional:_Signed_kernel_modules_2)
    - [3.2.3Optional: Signing the kernel image (Secure Boot)](#Optional:_Signing_the_kernel_image_.28Secure_Boot.29_2)
    - [3.2.4Compiling and installing](#Compiling_and_installing)

## Installing firmware and/or microcode\[ [edit](/index.php?title=Handbook:Parts/Installation/Kernel&action=edit&section=T-1 "Edit section: ")\]

### Firmware\[ [edit](/index.php?title=Handbook:Parts/Installation/Kernel&action=edit&section=T-2 "Edit section: ")\]

#### Suggested: Linux Firmware\[ [edit](/index.php?title=Handbook:Parts/Installation/Kernel&action=edit&section=T-3 "Edit section: ")\]

On many systems, non-FOSS firmware is required for certain hardware to function. The [sys-kernel/linux-firmware](https://packages.gentoo.org/packages/sys-kernel/linux-firmware) package contains firmware for many, but not all, devices.

**Tip**

Most wireless cards and GPUs require firmware to function.

`root #` `emerge --ask sys-kernel/linux-firmware`

**Note**

Installing certain firmware packages often requires accepting the associated firmware licenses. If necessary, visit the [license handling section](/wiki/Handbook:X86/Working/Portage#Licenses "Handbook:X86/Working/Portage") of the Handbook for help on accepting licenses.

##### Firmware Loading\[ [edit](/index.php?title=Handbook:Parts/Installation/Kernel&action=edit&section=T-4 "Edit section: ")\]

Firmware files are typically loaded when the associated kernel module is loaded. This means the firmware must be built into the kernel using **CONFIG\_EXTRA\_FIRMWARE** if the kernel module is set to _Y_ instead of _M_. In most cases, building-in a module which required firmware can complicate or break loading.

### Suggested: Microcode

In addition to discrete graphics hardware and network interfaces, CPUs can also require firmware updates. Typically this kind of firmware is referred to as _microcode_. Newer revisions of microcode are sometimes necessary to patch instability, security concerns, or other miscellaneous bugs in CPU hardware.

Microcode updates for AMD CPUs are distributed within the aforementioned [sys-kernel/linux-firmware](https://packages.gentoo.org/packages/sys-kernel/linux-firmware) package. Microcode for Intel CPUs can be found within the [sys-firmware/intel-microcode](https://packages.gentoo.org/packages/sys-firmware/intel-microcode) package, which will need to be installed separately. See the [Microcode article](/wiki/Microcode "Microcode") for more information on how to apply microcode updates.

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

#### systemd-boot

When using [systemd-boot](/wiki/Systemd/systemd-boot "Systemd/systemd-boot") (formerly gummiboot) as the bootloader, systemd's kernel-install must be used. Therefore ensure the [systemd](https://packages.gentoo.org/useflags/systemd) [USE flag](/wiki/USE_flag "USE flag") and the [systemd-boot](https://packages.gentoo.org/useflags/systemd-boot) [USE flag](/wiki/USE_flag "USE flag") USE flags are enabled on [sys-kernel/installkernel](https://packages.gentoo.org/packages/sys-kernel/installkernel), and then install the relevant package for systemd-boot.

FILE **`/etc/portage/package.use/systemd-boot`**

```
sys-apps/systemd boot
sys-apps/systemd-utils boot kernel-install
sys-kernel/installkernel systemd systemd-boot

```

The kernel command line to use for new kernels should be specified in /etc/kernel/cmdline and /etc/cmdline. As they both need to be updated together, it make sense to symlink the two files now.

FILE **`/etc/kernel/cmdline`**

```
quiet splash

```

`root #` `mkdir -p /etc/cmdline.d`

`root #` `ln -s /etc/kernel/cmdline /etc/cmdline.d/00-installkernel.conf`

`root #` `emerge --ask sys-kernel/installkernel`

**Note**

systemd-boot requires kernels to be installed to /efi.

#### EFI stub

UEFI-based computer systems technically do not need secondary bootloaders in order to boot kernels. Secondary bootloaders exist to _extend_ the functionality of UEFI firmware during the boot process. That being said, using a secondary bootloader is typically easier and more robust because it offers a more flexible approach for quickly modifying kernel parameters at boot time. Note also that UEFI implementations strongly differ between vendors and between models and there is no guarantee that a given firmware follows the UEFI specification. Therefore, EFI Stub booting is not guaranteed to work on every UEFI-based system.

`root #` `emerge --ask sys-kernel/installkernel`

`root #` `mkdir -p /efi`

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

### Optional: Unified Kernel Image

A [Unified Kernel Image](/wiki/Unified_Kernel_Image "Unified Kernel Image") (UKI) combines, among other things, the kernel, the initramfs and the kernel command line into a single executable. Since the kernel command line is embedded into the unified kernel image, it should be specified before generating the unified kernel image (see below). Note that any kernel command line arguments supplied by the bootloader or firmware at boot are ignored when booting with secure boot enabled.

A unified kernel image requires a stub loader. Currently, the only one available is systemd-stub. To enable it:

For systemd systems:

FILE **`/etc/portage/package.use/uki`**

```
sys-apps/systemd boot

```

`root #` `emerge --ask sys-apps/systemd`

For OpenRC systems:

FILE **`/etc/portage/package.use/uki`**

```
sys-apps/systemd-utils boot kernel-install

```

`root #` `emerge --ask sys-apps/systemd-utils`

[Installkernel](/wiki/Installkernel "Installkernel") can automatically generate a unified kernel image using either [dracut](/wiki/Unified_kernel_image#dracut "Unified kernel image") or [ukify](/wiki/Unified_kernel_image#ukify "Unified kernel image") by enabling the respective flag and the [uki](https://packages.gentoo.org/useflags/uki) [USE flag](/wiki/USE_flag "USE flag") USE flag.

For dracut:

FILE **`/etc/portage/package.use/uki`**

```
sys-kernel/installkernel dracut uki

```

FILE **`/etc/dracut.conf.d/uki.conf`**

```
uefi="yes"
kernel_cmdline="some-kernel-command-line-arguments"

```

`root #` `emerge --ask sys-kernel/installkernel`

For ukify:

FILE **`/etc/portage/package.use/uki`**

```
sys-apps/systemd boot ukify                         # For systemd systems
sys-apps/systemd-utils kernel-install boot ukify    # For OpenRC systems
sys-kernel/installkernel dracut ukify uki

```

FILE **`/etc/kernel/cmdline`**

```
some-kernel-command-line-arguments

```

`root #` `emerge --ask sys-kernel/installkernel`

Note that while dracut can generate both an initramfs and a unified kernel image, ukify can only generate the latter and therefore the initramfs must be generated separately with dracut.

**Important**

In the above configuration examples (for both Dracut and ukify) it is important to specify at least an appropriate root= parameter for the kernel command line to ensure that the Unified Kernel Image can find the root partition. This is not required for systemd based systems following the Discoverable Partitions Specification (DPS), in that case the embedded initramfs will be able to dynamically find the root partition.

#### Generic Unified Kernel Image (systemd only)

The prebuilt [sys-kernel/gentoo-kernel-bin](https://packages.gentoo.org/packages/sys-kernel/gentoo-kernel-bin) can optionally install a prebuilt generic unified kernel image containing a generic initramfs that is able to boot most systemd based systems. It can be installed by enabling the [generic-uki](https://packages.gentoo.org/useflags/generic-uki) [USE flag](/wiki/USE_flag "USE flag") USE flag, and configuring [installkernel](/wiki/Installkernel "Installkernel") to not generate a custom initramfs or unified kernel image:

FILE **`/etc/portage/package.use/uki`**

```
sys-kernel/gentoo-kernel-bin generic-uki
sys-kernel/installkernel -dracut -ukify -ugrd uki

```

#### Secure Boot

**Warning**

If following this section and manually compiling your own kernel, then make sure to follow the steps outlined in [Signing the kernel](/wiki/Kernel#Optional:_Signing_the_kernel_image_.28Secure_Boot.29 "Kernel")

The generic Unified Kernel Image optionally distributed by [sys-kernel/gentoo-kernel-bin](https://packages.gentoo.org/packages/sys-kernel/gentoo-kernel-bin) is already pre-signed. How to sign a locally generated unified kernel image depends on whether dracut or ukify is used. Note that the location of the key and certificate should be the same as the SECUREBOOT\_SIGN\_KEY and SECUREBOOT\_SIGN\_CERT as specified in /etc/portage/make.conf.

For dracut:

FILE **`/etc/dracut.conf.d/uki.conf`**

```
uefi="yes"
kernel_cmdline="some-kernel-command-line-arguments"
uefi_secureboot_key="/path/to/kernel_key.pem"
uefi_secureboot_cert="/path/to/kernel_key.pem"

```

For ukify:

FILE **`/etc/kernel/uki.conf`**

```
[UKI]
SecureBootPrivateKey=/path/to/kernel_key.pem
SecureBootCertificate=/path/to/kernel_key.pem

```

## Kernel selection\[ [edit](/index.php?title=Handbook:Parts/Installation/Kernel&action=edit&section=T-17 "Edit section: ")\]

It can be a wise move to use the dist-kernel on the first boot as it provides a very simple method to rule out system issues and kernel config issues. Always having a known working kernel to fallback on can speed up debugging and alleviate anxiety when updating that your system will no longer boot.

A common misconception is that a manually compiled kernel will use a lot less RAM than a pre configured distribution kernel. Due to the modular nature of the Linux kernel, only what is needed by the system is loaded and in most cases resulting in similar memory usage.

**Important**

Only one selection is required in the following subsection, if unsure of which to use then go with the first listed for now. It's always possible to switch at a later date if required.

Ranked from least involved to most involved:

[Full automation approach: Distribution kernels](/wiki/Handbook:X86/Installation/Kernel#Distribution_kernels "Handbook:X86/Installation/Kernel")A [Distribution Kernel](/wiki/Project:Distribution_Kernel "Project:Distribution Kernel") is used to configure, automatically build, and install the Linux kernel, its associated modules, and (optionally, but enabled by default) an initramfs file. Future kernel updates are fully automated since they are handled through the package manager, just like any other system package. It is possible [provide a custom kernel configuration file](/wiki/Project:Distribution_Kernel#Modifying_kernel_configuration "Project:Distribution Kernel") if customization is necessary. This is the least involved process and is perfect for new Gentoo users due to it working out-of-the-box and offering minimal involvement from the system administrator.[Manual approach](/wiki/Handbook:X86/Installation/Kernel#Alternative:_Manual_configuration "Handbook:X86/Installation/Kernel")New kernel sources are installed via the system package manager. The kernel is manually configured, built, and installed using the eselect kernel and a slew of make commands. Future kernel updates repeat the manual process of configuring, building, and installing the kernel files. This is the most involved process, but offers maximum control over the kernel update process.

The core around which all distributions are built is the Linux kernel. It is the layer between the user's programs and the system hardware. Although the handbook provides its users several possible kernel sources, a more comprehensive listing with more detailed descriptions is available at the [kernel packages page](/wiki/Kernel/Packages "Kernel/Packages").

**Tip**

Kernel installation tasks such as copying the kernel image to /boot or the [EFI System Partition](/wiki/EFI_System_Partition "EFI System Partition"), generating an [initramfs](/wiki/Initramfs "Initramfs") and/or [Unified Kernel Image](/wiki/Unified_Kernel_Image "Unified Kernel Image"), updating bootloader configuration, can be automated with [installkernel](/wiki/Installkernel "Installkernel"). Users may wish to configure and install [sys-kernel/installkernel](https://packages.gentoo.org/packages/sys-kernel/installkernel) before proceeding. See the [Kernel installation section below](/wiki/Handbook:X86/Installation/Kernel#Kernel_installation "Handbook:X86/Installation/Kernel") for more more information.

### Distribution kernels\[ [edit](/index.php?title=Handbook:Parts/Installation/Kernel/Dist-Kernel&action=edit&section=T-1 "Edit section: ")\]

_[Distribution Kernels](/wiki/Project:Distribution_Kernel "Project:Distribution Kernel")_ are ebuilds that cover the complete process of unpacking, configuring, compiling, and installing the kernel. The primary advantage of this method is that the kernels are updated to new versions by the package manager as part of @world upgrade. This requires no more involvement than running an emerge command. Distribution kernels default to a configuration supporting the majority of hardware, however two mechanisms are offered for customization: savedconfig and config snippets. See the project page for [more details on configuration.](/wiki/Project:Distribution_Kernel#Modifying_kernel_configuration "Project:Distribution Kernel")

##### Optional: Signed kernel modules\[ [edit](/index.php?title=Handbook:Parts/Installation/Kernel/Dist-Kernel&action=edit&section=T-2 "Edit section: ")\]

The kernel modules in the prebuilt distribution kernel ([sys-kernel/gentoo-kernel-bin](https://packages.gentoo.org/packages/sys-kernel/gentoo-kernel-bin)) are already signed. To sign the modules of kernels built from source enable the [modules-sign](https://packages.gentoo.org/useflags/modules-sign) [USE flag](/wiki/USE_flag "USE flag") USE flag, and optionally specify which key to use for signing in [/etc/portage/make.conf](/wiki//etc/portage/make.conf "/etc/portage/make.conf"):

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

**Note**

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

The kernel image in the prebuilt distribution kernel ([sys-kernel/gentoo-kernel-bin](https://packages.gentoo.org/packages/sys-kernel/gentoo-kernel-bin)) is already signed for use with [Secure Boot](/wiki/Secure_Boot "Secure Boot"). To sign the kernel image of kernels built from source enable the [secureboot](https://packages.gentoo.org/useflags/secureboot) [USE flag](/wiki/USE_flag "USE flag") USE flag, and optionally specify which key to use for signing in [/etc/portage/make.conf](/wiki//etc/portage/make.conf "/etc/portage/make.conf"). Note that signing the kernel image for use with secureboot requires that the kernel modules are also signed, the same key may be used to sign both the kernel image and the kernel modules:

FILE **`/etc/portage/make.conf`** **Enable custom signing keys**

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

**Note**

The SECUREBOOT\_SIGN\_KEY and SECUREBOOT\_SIGN\_CERT may be different files. For this example the pem file generated by OpenSSL includes both the key and the accompanying certificate, and thus both variables are set to the same value.

**Note**

For this example the same key that was generated to sign the modules is used to sign the kernel image. It is also possible to generate and use a second separate key for signing the kernel image. The same OpenSSL command as in the previous section may be used again.

See the above section for instructions on generating a new key, the steps may be repeated if a separate key should be used to sign the kernel image.

To successfully boot with Secure Boot enabled, the used bootloader must also be signed and the certificate must be accepted by the [UEFI](/wiki/UEFI "UEFI") firmware or [Shim](/wiki/Shim "Shim"). This will be explained later in the handbook.

#### Installing a distribution kernel\[ [edit](/index.php?title=Handbook:Parts/Installation/Kernel/Dist-Kernel&action=edit&section=T-4 "Edit section: ")\]

To build a kernel with Gentoo patches from source, type:

`root #` `emerge --ask sys-kernel/gentoo-kernel`

System administrators who want to avoid compiling the kernel sources locally can instead use precompiled kernel images:

`root #` `emerge --ask sys-kernel/gentoo-kernel-bin`

**Important**

_[Distribution Kernels](/wiki/Project:Distribution_Kernel "Project:Distribution Kernel")_, such as [sys-kernel/gentoo-kernel](https://packages.gentoo.org/packages/sys-kernel/gentoo-kernel) and [sys-kernel/gentoo-kernel-bin](https://packages.gentoo.org/packages/sys-kernel/gentoo-kernel-bin), by default, expect to be installed alongside an [initramfs](/wiki/Handbook:X86/Installation/Kernel#Initramfs "Handbook:X86/Installation/Kernel"). Before running emerge to install the kernel users should ensure that [sys-kernel/installkernel](https://packages.gentoo.org/packages/sys-kernel/installkernel) has been configured to utilize an initramfs generator (for example [Dracut](/wiki/Dracut "Dracut")) as described in the [installkernel section](/wiki/Handbook:X86/Installation/Kernel#Initramfs "Handbook:X86/Installation/Kernel").

#### Upgrading and cleaning up\[ [edit](/index.php?title=Handbook:Parts/Installation/Kernel/Dist-Kernel&action=edit&section=T-5 "Edit section: ")\]

Once the kernel is installed, the package manager will automatically update it to newer versions. The previous versions will be kept until the package manager is requested to clean up stale packages. To reclaim disk space, stale packages can be trimmed by periodically running emerge with the `--depclean` option:

`root #` `emerge --depclean`

Alternatively, to specifically clean up old kernel versions:

`root #` `emerge --prune sys-kernel/gentoo-kernel sys-kernel/gentoo-kernel-bin`

**Tip**

By design, emerge only removes the kernel build directory. It does not actually remove the kernel modules, nor the installed kernel image. To completely clean-up old kernels, the [app-admin/eclean-kernel](https://packages.gentoo.org/packages/app-admin/eclean-kernel) tool may be used.

#### Post-install/upgrade tasks\[ [edit](/index.php?title=Handbook:Parts/Installation/Kernel/Dist-Kernel&action=edit&section=T-6 "Edit section: ")\]

An upgrade of a distribution kernel is capable of triggering an automatic rebuild for external kernel modules installed by other packages (for example: [sys-fs/zfs-kmod](https://packages.gentoo.org/packages/sys-fs/zfs-kmod) or [x11-drivers/nvidia-drivers](https://packages.gentoo.org/packages/x11-drivers/nvidia-drivers)). This automated behaviour is enabled by enabling the [dist-kernel](https://packages.gentoo.org/useflags/dist-kernel) [USE flag](/wiki/USE_flag "USE flag") USE flag. When required, this same flag will also trigger re-generation of the [initramfs](/wiki/Initramfs "Initramfs").

It is highly recommended to enable this flag globally via /etc/portage/make.conf when using a distribution kernel:

FILE **`/etc/portage/make.conf`** **Enabling USE=dist-kernel**

```
USE="dist-kernel"

```

##### Manually rebuilding the initramfs or Unified Kernel Image\[ [edit](/index.php?title=Handbook:Parts/Installation/Kernel/Dist-Kernel&action=edit&section=T-7 "Edit section: ")\]

If required, manually trigger such rebuilds by, after a kernel upgrade, executing:

`root #` `emerge --ask @module-rebuild`

If any kernel modules (e.g. ZFS) are needed at early boot, rebuild the initramfs afterward via:

`root #` `emerge --config sys-kernel/gentoo-kernel
`

`root #` `emerge --config sys-kernel/gentoo-kernel-bin
`

After installing the Distribution Kernel successfully, it is now time to proceed to the next section: [Configuring the system](/wiki/Handbook:X86/Installation/System "Handbook:X86/Installation/System").

### Alternative: Manual configuration\[ [edit](/index.php?title=Handbook:Parts/Installation/Kernel&action=edit&section=T-18 "Edit section: ")\]

#### Installing and Configuring the kernel sources\[ [edit](/index.php?title=Handbook:Parts/Installation/Kernel&action=edit&section=T-19 "Edit section: ")\]

When manually compiling the kernel for x86-based systems, Gentoo recommends the [sys-kernel/gentoo-sources](https://packages.gentoo.org/packages/sys-kernel/gentoo-sources) package.

Choose an appropriate kernel source and install it using emerge:

`root #` `emerge --ask sys-kernel/gentoo-sources`

This will install the Linux kernel sources in /usr/src/ using the specific kernel version in the path. It will not create a symbolic link by itself without the [symlink](https://packages.gentoo.org/useflags/symlink) [USE flag](/wiki/USE_flag "USE flag") USE flag being enabled on the chosen kernel sources package.

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

**Important**

Only one selection is required in the following subsection, if unsure of which to use then go with the first listed for now. It's always possible to switch at a later date if required.

##### Option 1 - Modprobed-db process

A very easy way to manage the kernel is to first install [sys-kernel/gentoo-kernel-bin](https://packages.gentoo.org/packages/sys-kernel/gentoo-kernel-bin) and use the [sys-kernel/modprobed-db](https://packages.gentoo.org/packages/sys-kernel/modprobed-db) to collect information about what the system requires. modprobed-db is a tool which monitors the system via crontab to add all modules of all devices over the system's life to make sure it everything a user needs is supported. For example, if an Xbox controller is added after installation, then modprobed-db will add the modules to be built next time the kernel is rebuilt.

For now please follow installing [sys-kernel/gentoo-kernel-bin](https://packages.gentoo.org/packages/sys-kernel/gentoo-kernel-bin) via [Distribution\_kernels](/wiki/Handbook:X86/Installation/Kernel#Distribution_kernels "Handbook:X86/Installation/Kernel").

Next, install [sys-kernel/modprobed-db](https://packages.gentoo.org/packages/sys-kernel/modprobed-db):

`root #` `emerge --ask sys-kernel/modprobed-db`

Please watch out for further steps related to modprobed-db in the Handbook.

More on this topic can be found in the [Modprobed-db](/wiki/Modprobed-db "Modprobed-db") article.

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

Now it's time to decide if modules signing is required in the steps listed in [here](/wiki/Handbook:X86/Installation/Kernel#Optional:_Signed_kernel_modules_2 "Handbook:X86/Installation/Kernel")

If not, proceed to building described [here](/wiki/Handbook:X86/Installation/Kernel#Compiling_and_installing "Handbook:X86/Installation/Kernel")

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

**Note**

For this example, the same key that was generated to sign the modules is used to sign the kernel image. It is also possible to generate and use a second separate key for signing the kernel image. The same OpenSSL command as in the previous section may be used again.

Then proceed with the installation.

To automatically sign EFI executables installed by other packages, enable the [secureboot](https://packages.gentoo.org/useflags/secureboot) [USE flag](/wiki/USE_flag "USE flag") USE flag globally:

FILE **`/etc/portage/make.conf`** **Enable Secure Boot**

```
USE="modules-sign secureboot"

# Optionally, to use custom signing keys.
MODULES_SIGN_KEY="/path/to/kernel_key.pem"
MODULES_SIGN_CERT="/path/to/kernel_key.pem" # Only required if the MODULES_SIGN_KEY does not also contain the certificate.
MODULES_SIGN_HASH="sha512" # Defaults to sha512

# Optionally, to boot with secureboot enabled, may be the same or different signing key.
SECUREBOOT_SIGN_KEY="/path/to/kernel_key.pem"
SECUREBOOT_SIGN_CERT="/path/to/kernel_key.pem"

```

**Note**

SECUREBOOT\_SIGN\_KEY and SECUREBOOT\_SIGN\_CERT may point to different files. For this example, the pem file generated by OpenSSL includes both the key and the accompanying certificate, and thus both variables are set to the same value.

**Note**

When generating an [Unified Kernel Image](/wiki/Unified_Kernel_Image "Unified Kernel Image") with systemd's `ukify` the kernel image will be signed automatically before inclusion in the unified kernel image and it is not necessary to sign it manually.

#### Compiling and installing\[ [edit](/index.php?title=Handbook:X86/Blocks/Kernel&action=edit&section=T-1 "Edit section: ")\]

With the configuration now done, it is time to compile and install the kernel. Exit the configuration and start the compilation process:

`root #` `make && make modules_install`

**Note**

It is possible to enable parallel builds using make -j N with N being an integer number of parallel tasks that the build process is allowed to launch. This is similar to the instructions about /etc/portage/make.conf earlier, with the MAKEOPTS variable.

When the kernel has finished compiling, copy the kernel image to /boot/. This is handled by the make install command:

`root #` `make install`

This will copy the kernel image into /boot/ together with the System.map file and the kernel configuration file.

Continue the installation with [Configuring the system](/wiki/Handbook:X86/Installation/System "Handbook:X86/Installation/System").

[← Installing base system](/wiki/Handbook:X86/Installation/Base "Previous part") [Home](/wiki/Handbook:X86 "Handbook:X86") [Configuring the system →](/wiki/Handbook:X86/Installation/System "Next part")