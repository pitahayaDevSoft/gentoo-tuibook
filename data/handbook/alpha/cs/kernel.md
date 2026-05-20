# Kernel

Other languages:

- [Deutsch](/wiki/Handbook:Alpha/Installation/Kernel/de "Handbuch:Alpha/Installation/Kernel (100% translated)")
- [English](/wiki/Handbook:Alpha/Installation/Kernel "Handbook:Alpha/Installation/Kernel (100% translated)")
- [español](/wiki/Handbook:Alpha/Installation/Kernel/es "Manual de Gentoo: Alpha/Instalación/Núcleo (100% translated)")
- [français](/wiki/Handbook:Alpha/Installation/Kernel/fr "Manuel:Alpha/Installation/Noyau (100% translated)")
- [italiano](/wiki/Handbook:Alpha/Installation/Kernel/it "Handbook:Alpha/Installation/Kernel/it (100% translated)")
- [magyar](/wiki/Handbook:Alpha/Installation/Kernel/hu "Handbook:Alpha/Installation/Kernel/hu (100% translated)")
- [polski](/wiki/Handbook:Alpha/Installation/Kernel/pl "Handbook:Alpha/Installation/Kernel (100% translated)")
- [português do Brasil](/wiki/Handbook:Alpha/Installation/Kernel/pt-br "Manual:Alpha/Instalação/Kernel (100% translated)")
- čeština
- [русский](/wiki/Handbook:Alpha/Installation/Kernel/ru "Handbook:Alpha/Installation/Kernel (100% translated)")
- [தமிழ்](/wiki/Handbook:Alpha/Installation/Kernel/ta "கையேடு:Alpha/நிறுவல்/கர்னல் (100% translated)")
- [中文（中国大陆）‎](/wiki/Handbook:Alpha/Installation/Kernel/zh-cn "手册：Alpha/安装/配置Linux内核 (100% translated)")
- [日本語](/wiki/Handbook:Alpha/Installation/Kernel/ja "ハンドブック:Alpha/インストール/カーネル (100% translated)")
- [한국어](/wiki/Handbook:Alpha/Installation/Kernel/ko "Handbook:Alpha/Installation/Kernel/ko (100% translated)")

[Alpha Handbook](/wiki/Handbook:Alpha/cs "Handbook:Alpha/cs")[Instalace](/wiki/Handbook:Alpha/Full/Installation/cs "Handbook:Alpha/Full/Installation/cs")[O instalaci](/wiki/Handbook:Alpha/Installation/About/cs "Handbook:Alpha/Installation/About/cs")[Výběr média](/wiki/Handbook:Alpha/Installation/Media/cs "Handbook:Alpha/Installation/Media/cs")[Konfigurace sítě](/wiki/Handbook:Alpha/Installation/Networking/cs "Handbook:Alpha/Installation/Networking/cs")[Příprava disků](/wiki/Handbook:Alpha/Installation/Disks/cs "Handbook:Alpha/Installation/Disks/cs")[Instalace stage3](/wiki/Handbook:Alpha/Installation/Stage/cs "Handbook:Alpha/Installation/Stage/cs")[Instalace základního systému](/wiki/Handbook:Alpha/Installation/Base/cs "Handbook:Alpha/Installation/Base/cs")Konfigurace jádra[Konfigurace systému](/wiki/Handbook:Alpha/Installation/System/cs "Handbook:Alpha/Installation/System/cs")[Instalace nástrojů](/wiki/Handbook:Alpha/Installation/Tools/cs "Handbook:Alpha/Installation/Tools/cs")[Konfigurace zavaděče](/wiki/Handbook:Alpha/Installation/Bootloader/cs "Handbook:Alpha/Installation/Bootloader/cs")[Dokončení](/wiki/Handbook:Alpha/Installation/Finalizing/cs "Handbook:Alpha/Installation/Finalizing/cs")[Práce s Gentoo](/wiki/Handbook:Alpha/Full/Working/cs "Handbook:Alpha/Full/Working/cs")[Úvod do Portage](/wiki/Handbook:Alpha/Working/Portage/cs "Handbook:Alpha/Working/Portage/cs")[Přepínače USE](/wiki/Handbook:Alpha/Working/USE/cs "Handbook:Alpha/Working/USE/cs")[Funkce portage](/wiki/Handbook:Alpha/Working/Features/cs "Handbook:Alpha/Working/Features/cs")[Systém initskriptů](/wiki/Handbook:Alpha/Working/Initscripts/cs "Handbook:Alpha/Working/Initscripts/cs")[Proměnné prostředí](/wiki/Handbook:Alpha/Working/EnvVar/cs "Handbook:Alpha/Working/EnvVar/cs")[Práce s Portage](/wiki/Handbook:Alpha/Full/Portage/cs "Handbook:Alpha/Full/Portage/cs")[Soubory a adresáře](/wiki/Handbook:Alpha/Portage/Files/cs "Handbook:Alpha/Portage/Files/cs")[Proměnné](/wiki/Handbook:Alpha/Portage/Variables/cs "Handbook:Alpha/Portage/Variables/cs")[Mísení softwarových větví](/wiki/Handbook:Alpha/Portage/Branches/cs "Handbook:Alpha/Portage/Branches/cs")[Doplňkové nástroje](/wiki/Handbook:Alpha/Portage/Tools/cs "Handbook:Alpha/Portage/Tools/cs")[Vlastní strom Portage](/wiki/Handbook:Alpha/Portage/CustomTree/cs "Handbook:Alpha/Portage/CustomTree/cs")[Pokročilé funkce](/wiki/Handbook:Alpha/Portage/Advanced/cs "Handbook:Alpha/Portage/Advanced/cs")[Konfigurace sítě](/wiki/Handbook:Alpha/Full/Networking/cs "Handbook:Alpha/Full/Networking/cs")[Začínáme](/wiki/Handbook:Alpha/Networking/Introduction/cs "Handbook:Alpha/Networking/Introduction/cs")[Pokročilá konfigurace](/wiki/Handbook:Alpha/Networking/Advanced/cs "Handbook:Alpha/Networking/Advanced/cs")[Modulární síťování](/wiki/Handbook:Alpha/Networking/Modular/cs "Handbook:Alpha/Networking/Modular/cs")[Bezdrátové sítě](/wiki/Handbook:Alpha/Networking/Wireless/cs "Handbook:Alpha/Networking/Wireless/cs")[Přidání funkcí](/wiki/Handbook:Alpha/Networking/Extending/cs "Handbook:Alpha/Networking/Extending/cs")[Dynamická správa](/wiki/Handbook:Alpha/Networking/Dynamic/cs "Handbook:Alpha/Networking/Dynamic/cs")

## Contents

- [1Volitelné: Instalace firmwaru](#Voliteln.C3.A9:_Instalace_firmwaru)
- [2Firmware](#Firmware)
  - [2.1Suggested: Linux Firmware](#Suggested:_Linux_Firmware)
    - [2.1.1Firmware Loading](#Firmware_Loading)
- [3sys-kernel/installkernel](#sys-kernel.2Finstallkernel)
  - [3.1Bootloader](#Bootloader)
    - [3.1.1GRUB](#GRUB)
    - [3.1.2Traditional layout, other bootloaders (e.g. (e)lilo, syslinux, etc.)](#Traditional_layout.2C_other_bootloaders_.28e.g._.28e.29lilo.2C_syslinux.2C_etc..29)
  - [3.2Initramfs](#Initramfs)
    - [3.2.1Chroot detection](#Chroot_detection)
- [4Kernel selection](#Kernel_selection)
- [5Výchozí: Ruční konfigurace](#V.C3.BDchoz.C3.AD:_Ru.C4.8Dn.C3.AD_konfigurace)
- [6Instalace zdrojových kódů](#Instalace_zdrojov.C3.BDch_k.C3.B3d.C5.AF)
  - [6.1Option 2 - Assisted manual process](#Option_2_-_Assisted_manual_process)
  - [6.2Option 3 - Configuring by hand](#Option_3_-_Configuring_by_hand)
  - [6.3Optional: Signed kernel modules](#Optional:_Signed_kernel_modules)

### Volitelné: Instalace firmwaru

### Firmware

#### Suggested: Linux Firmware

On many systems, non-FOSS firmware is required for certain hardware to function. The [sys-kernel/linux-firmware](https://packages.gentoo.org/packages/sys-kernel/linux-firmware) package contains firmware for many, but not all, devices.

**Tip**

Most wireless cards and GPUs require firmware to function.

`root #` `emerge --ask sys-kernel/linux-firmware`

**Poznámka**

Installing certain firmware packages often requires accepting the associated firmware licenses. If necessary, visit the [license handling section](/wiki/Handbook:Alpha/Working/Portage/cs#Licenses "Handbook:Alpha/Working/Portage/cs") of the Handbook for help on accepting licenses.

##### Firmware Loading

Firmware files are typically loaded when the associated kernel module is loaded. This means the firmware must be built into the kernel using **CONFIG\_EXTRA\_FIRMWARE** if the kernel module is set to _Y_ instead of _M_. In most cases, building-in a module which required firmware can complicate or break loading.

## sys-kernel/installkernel

[Installkernel](/wiki/Installkernel "Installkernel") may be used to automate the kernel installation, [initramfs](/wiki/Initramfs "Initramfs") generation, [unified kernel image](/wiki/Unified_kernel_image "Unified kernel image") generation and bootloader configuration, among other things. [sys-kernel/installkernel](https://packages.gentoo.org/packages/sys-kernel/installkernel) implements two paths of achieving this: the traditional installkernel originating from Debian and [systemd](/wiki/Systemd "Systemd")'s kernel-install. Which one to choose depends, among other things, on the system's bootloader. By default, systemd's kernel-install is used on systemd profiles, while the traditional installkernel is the default for other profiles.

### Bootloader

Now is the time to think about which bootloader the user wants for the system.

**Důležité**

Only one selection is required in the following subsection, if unsure of which to use then go with the first listed for now. It's always possible to switch at a later date if required.

#### GRUB

Users of GRUB can use either systemd's kernel-install or the traditional Debian installkernel. The [systemd](https://packages.gentoo.org/useflags/systemd) [USE flag/cs (page does not exist)](/index.php?title=USE_flag/cs&action=edit&redlink=1 "USE flag/cs (page does not exist)") USE flag switches between these implementations. To automatically run grub-mkconfig when installing the kernel, enable the [grub](https://packages.gentoo.org/useflags/grub) [USE flag/cs (page does not exist)](/index.php?title=USE_flag/cs&action=edit&redlink=1 "USE flag/cs (page does not exist)") [USE flag](/wiki/USE_flag "USE flag").

**Poznámka**

GRUB requires kernels to be installed to /boot.

FILE **`/etc/portage/package.use/installkernel`**

```
sys-kernel/installkernel grub

```

`root #` `emerge --ask sys-kernel/installkernel`

**Poznámka**

systemd-boot requires kernels to be installed to /efi.

**Poznámka**

When [app-emulation/virt-firmware](https://packages.gentoo.org/packages/app-emulation/virt-firmware) is used to configure the UEFI ensure that the kernel-bootcfg-boot-successful service is enabled before attempting to install the kernel. This implementation of EFIstub booting is the default for systemd systems.

`root #` `systemctl enable kernel-bootcfg-boot-successful.service`

**Poznámka**

EFIstub requires kernels to be installed to /efi.

#### Traditional layout, other bootloaders (e.g. (e)lilo, syslinux, etc.)

The traditional /boot layout (for e.g. (e)LILO, syslinux, etc.) is used by default if the [grub](https://packages.gentoo.org/useflags/grub) [USE flag/cs (page does not exist)](/index.php?title=USE_flag/cs&action=edit&redlink=1 "USE flag/cs (page does not exist)"), [systemd-boot](https://packages.gentoo.org/useflags/systemd-boot) [USE flag/cs (page does not exist)](/index.php?title=USE_flag/cs&action=edit&redlink=1 "USE flag/cs (page does not exist)"), [efistub](https://packages.gentoo.org/useflags/efistub) [USE flag/cs (page does not exist)](/index.php?title=USE_flag/cs&action=edit&redlink=1 "USE flag/cs (page does not exist)") and [uki](https://packages.gentoo.org/useflags/uki) [USE flag/cs (page does not exist)](/index.php?title=USE_flag/cs&action=edit&redlink=1 "USE flag/cs (page does not exist)") USE flags are **not** enabled. No further action is required.

### Initramfs

An **init** ial **ram**-based **f** ile **s** ystem, or [initramfs](/wiki/Initramfs "Initramfs"), may be required for a system to boot. A wide of variety of cases may necessitate one, but common cases include:

- Kernels where storage/filesystem drivers are modules.
- Layouts with /usr/ or /var/ on separate partitions.
- Encrypted root filesystems.

**Tip**

[Distribution kernels](/wiki/Project:Distribution_Kernel "Project:Distribution Kernel") are designed to be used with an initramfs, as many storage and filesystem drivers are built as modules.

In addition to mounting the root filesystem, an initramfs may also perform other tasks such as:

- Running **f** ile **s** ystem **c** onsistency chec **k** fsck, a tool to check and repair consistency of a file system in such events of uncleanly shutdown a system.
- Providing a recovery environment in the event of late-boot failures.

[Installkernel](/wiki/Installkernel "Installkernel") can automatically generate an initramfs when installing the kernel if the [dracut](https://packages.gentoo.org/useflags/dracut) [USE flag/cs (page does not exist)](/index.php?title=USE_flag/cs&action=edit&redlink=1 "USE flag/cs (page does not exist)") or [ugrd](https://packages.gentoo.org/useflags/ugrd) [USE flag/cs (page does not exist)](/index.php?title=USE_flag/cs&action=edit&redlink=1 "USE flag/cs (page does not exist)") USE flag is enabled:

FILE **`/etc/portage/package.use/installkernel`**

```
sys-kernel/installkernel dracut

```

#### Chroot detection

Bootloaders such as [systemd-boot](/wiki/Systemd/systemd-boot "Systemd/systemd-boot") and [EFI stub](/wiki/EFI_stub "EFI stub") use the kernel arguments of the running system as set in /proc/cmdline by default. Because of the wide range of ways Gentoo can be installed users will randomly get tripped up by this.

To help solve any issues this may cause, [sys-kernel/installkernel](https://packages.gentoo.org/packages/sys-kernel/installkernel) will check if it is running in a chroot and fail if the kernel command line is not explicitly configured. Otherwise the bootloader would use the install media's boot arguments which would lead to boot failure.

One way to satisfy [sys-kernel/installkernel](https://packages.gentoo.org/packages/sys-kernel/installkernel) is by using Dracut's config file to set the root partition UUID as shown below, or alternatively for more information on what this check helps with and different ways to configure it, see [Installkernel#Install\_chroot\_check](/wiki/Installkernel#Install_chroot_check.2Fcs "Installkernel").

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

**Důležité**

Only one selection is required in the following subsection, if unsure of which to use then go with the first listed for now. It's always possible to switch at a later date if required.

Ranked from least involved to most involved:

1. Jádro se konfiguruje a sestaví ručně.
2. K automatickému sestavení a instalaci jádra Linux se použije nástroj zvaný genkernel.

Základ, kolem něhož jsou vybudovány všechny distribuce, je jádro Linux. Je to vrstva mezi uživatelskými programy a systémovým hardwarem. Gentoo nabízí svým uživatelům několik možný zdrojů jádra. Celý seznam s popisem je dostupný na [stránce Přehled jader](/wiki/Kernel/Overview "Kernel/Overview").

**Tip**

Kernel installation tasks such as copying the kernel image to /boot or the [EFI System Partition](/wiki/EFI_System_Partition "EFI System Partition"), generating an [initramfs](/wiki/Initramfs "Initramfs") and/or [Unified Kernel Image](/wiki/Unified_Kernel_Image "Unified Kernel Image"), updating bootloader configuration, can be automated with [installkernel](/wiki/Installkernel "Installkernel"). Users may wish to configure and install [sys-kernel/installkernel](https://packages.gentoo.org/packages/sys-kernel/installkernel) before proceeding. See the [Kernel installation section below](/wiki/Handbook:Alpha/Installation/Kernel#Kernel_installation.2Fcs "Handbook:Alpha/Installation/Kernel") for more more information.

## Výchozí: Ruční konfigurace

## Instalace zdrojových kódů

When manually compiling the kernel for alpha-based systems, Gentoo recommends the [sys-kernel/gentoo-sources](https://packages.gentoo.org/packages/sys-kernel/gentoo-sources) package.

Choose an appropriate kernel source and install it using emerge:

`root #` `emerge --ask sys-kernel/gentoo-sources`

Tím nainstalujete zdrojové kódy jádra do /usr/src, v němž bude symbolický odkaz linux ukazovat na nainstalované zdrojové kódy jádra:

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

**Důležité**

Only one selection is required in the following subsection, if unsure of which to use then go with the first listed for now. It's always possible to switch at a later date if required.

##### Option 2 - Assisted manual process

This method allows a user to have full control of how their kernel is built with as minimal help from outside tools as they wish. Some could consider this as making it hard for the sake of it.

Nicméně jedno je pravdou: při manuální konfiguraci je nutné znát systém. Většinu informací získáte nainstalováním balíčku [sys-apps/pciutils](https://packages.gentoo.org/packages/sys-apps/pciutils), který obsahuje příkaz lspci:

`root #` `emerge --ask sys-apps/pciutils`

**Poznámka**

Uvnitř chrootu můžete bezpečně ignorovat varování pcilib (jako je "pcilib cannot open /sys/bus/pci/devices"), které může lspci vypisovat.

Dalším zdrojem informací o systému je spuštění lsmod, po němž uvidíte jaké moduly jádra používá instalační CD, čímž můžete získat indicii o tom, co povolit.

Nyní přejděte do adresáře se zdrojovým kódem jádra a spusťte příkaz make menuconfig. Po něm na vás vyskočí obrazovka s konfiguračním menu.

`root #` `cd /usr/src/linux
`

`root #` `make menuconfig`

**Tip**

To view the full list of make arguments available for the kernel, run `make help`.

The kernel has a method of autodetecting the modules currently being used on the installcd which will give a great starting point to allow a user to configure their own. This can be called by using:

`root #` `make localmodconfig`

Manually configuring should not be needed at this point, but if a user wish to check:

`root #` `make nconfig`

Now it's time to decide if modules signing is required in the steps listed in [here](/wiki/Handbook:Alpha/Installation/Kernel#Optional:_Signed_kernel_modules_2.2Fcs "Handbook:Alpha/Installation/Kernel")

If not, proceed to building described [here](/wiki/Handbook:Alpha/Installation/Kernel#Compiling_and_installing.2Fcs "Handbook:Alpha/Installation/Kernel")

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

To also sign external kernel modules installed by other packages via `linux-mod-r1.eclass`, enable the [modules-sign](https://packages.gentoo.org/useflags/modules-sign) [USE flag/cs (page does not exist)](/index.php?title=USE_flag/cs&action=edit&redlink=1 "USE flag/cs (page does not exist)") USE flag globally:

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

**Poznámka**

MODULES\_SIGN\_KEY and MODULES\_SIGN\_CERT may point to different files. For this example, the pem file generated by OpenSSL includes both the key and the accompanying certificate, and thus both variables are set to the same value.

[Handbook:Alpha/Blocks/Kernel/cs](/index.php?title=Handbook:Alpha/Blocks/Kernel/cs&action=edit&redlink=1 "Handbook:Alpha/Blocks/Kernel/cs (page does not exist)")

V instalaci pokračujte [Konfigurací systému](/wiki/Handbook:Alpha/Installation/System/cs "Handbook:Alpha/Installation/System/cs").

[← Instalace základního systému](/wiki/Handbook:Alpha/Installation/Base/cs "Previous part") [Home](/wiki/Handbook:Alpha "Handbook:Alpha") [Konfirauce systému →](/wiki/Handbook:Alpha/Installation/System/cs "Next part")