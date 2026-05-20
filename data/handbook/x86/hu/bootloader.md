# Bootloader

Other languages:

- [Deutsch](/wiki/Handbook:X86/Installation/Bootloader/de "Handbuch:X86/Installation/Bootloader (100% translated)")
- [English](/wiki/Handbook:X86/Installation/Bootloader "Handbook:X86/Installation/Bootloader (100% translated)")
- [español](/wiki/Handbook:X86/Installation/Bootloader/es "Manual de Gentoo: X86/Instalación/Arranque (100% translated)")
- [français](/wiki/Handbook:X86/Installation/Bootloader/fr "Manuel:X86/Installation/Système d'amorçage (100% translated)")
- [italiano](/wiki/Handbook:X86/Installation/Bootloader/it "Handbook:X86/Installation/Bootloader (100% translated)")
- magyar
- [polski](/wiki/Handbook:X86/Installation/Bootloader/pl "Handbook:X86/Installation/Bootloader (100% translated)")
- [português do Brasil](/wiki/Handbook:X86/Installation/Bootloader/pt-br "Manual:X86/Instalação/Gerenciador de Boot (100% translated)")
- [čeština](/wiki/Handbook:X86/Installation/Bootloader/cs "Handbook:X86/Installation/Bootloader/cs (50% translated)")
- [русский](/wiki/Handbook:X86/Installation/Bootloader/ru "Handbook:X86/Installation/Bootloader (100% translated)")
- [தமிழ்](/wiki/Handbook:X86/Installation/Bootloader/ta "கையேடு:X86/நிறுவல்/துவக்கஏற்றி (100% translated)")
- [中文（中国大陆）‎](/wiki/Handbook:X86/Installation/Bootloader/zh-cn "手册：X86/安装/配置系统引导程序Bootloader (100% translated)")
- [日本語](/wiki/Handbook:X86/Installation/Bootloader/ja "ハンドブック:X86/インストール/ブートローダの設定 (100% translated)")
- [한국어](/wiki/Handbook:X86/Installation/Bootloader/ko "Handbook:X86/Installation/Bootloader/ko (100% translated)")

[X86 kézikönyv](/wiki/Handbook:X86/hu "Handbook:X86/hu")[A Gentoo Linux telepítése](/wiki/Handbook:X86/Full/Installation/hu "Handbook:X86/Full/Installation/hu")[A telepítésről](/wiki/Handbook:X86/Installation/About/hu "Handbook:X86/Installation/About/hu")[Telepítőképfájl kiválasztása](/wiki/Handbook:X86/Installation/Media/hu "Handbook:X86/Installation/Media/hu")[Hálózat beállítása](/wiki/Handbook:X86/Installation/Networking/hu "Handbook:X86/Installation/Networking/hu")[Adathordozók előkészítése](/wiki/Handbook:X86/Installation/Disks/hu "Handbook:X86/Installation/Disks/hu")[Stage fájl](/wiki/Handbook:X86/Installation/Stage/hu "Handbook:X86/Installation/Stage/hu")[Alaprendszer telepítése](/wiki/Handbook:X86/Installation/Base/hu "Handbook:X86/Installation/Base/hu")[Kernel beállítása](/wiki/Handbook:X86/Installation/Kernel/hu "Handbook:X86/Installation/Kernel/hu")[Rendszer beállítása](/wiki/Handbook:X86/Installation/System/hu "Handbook:X86/Installation/System/hu")[Eszközök telepítése](/wiki/Handbook:X86/Installation/Tools/hu "Handbook:X86/Installation/Tools/hu")Bootloader beállítása[Telepítés véglegesítése](/wiki/Handbook:X86/Installation/Finalizing/hu "Handbook:X86/Installation/Finalizing/hu")[Munka a Gentoo rendszerrel](/wiki/Handbook:X86/Full/Working/hu "Handbook:X86/Full/Working/hu")[Portage bemutatása](/wiki/Handbook:X86/Working/Portage/hu "Handbook:X86/Working/Portage/hu")[USE jelölőzászlók](/wiki/Handbook:X86/Working/USE/hu "Handbook:X86/Working/USE/hu")[Portage jellemzői](/wiki/Handbook:X86/Working/Features/hu "Handbook:X86/Working/Features/hu")[Init-szkript rendszer](/wiki/Handbook:X86/Working/Initscripts/hu "Handbook:X86/Working/Initscripts/hu")[Környezeti változók](/wiki/Handbook:X86/Working/EnvVar/hu "Handbook:X86/Working/EnvVar/hu")[Munka a Portage szoftvercsomag-kezelővel](/wiki/Handbook:X86/Full/Portage/hu "Handbook:X86/Full/Portage/hu")[Fájlok és könyvtárak](/wiki/Handbook:X86/Portage/Files/hu "Handbook:X86/Portage/Files/hu")[Változók](/wiki/Handbook:X86/Portage/Variables/hu "Handbook:X86/Portage/Variables/hu")[Szoftverágak keverése](/wiki/Handbook:X86/Portage/Branches/hu "Handbook:X86/Portage/Branches/hu")[További eszközök](/wiki/Handbook:X86/Portage/Tools/hu "Handbook:X86/Portage/Tools/hu")[Egyéni szoftvercsomag-tárolók](/wiki/Handbook:X86/Portage/CustomTree/hu "Handbook:X86/Portage/CustomTree/hu")[Fejlett funkciók](/wiki/Handbook:X86/Portage/Advanced/hu "Handbook:X86/Portage/Advanced/hu")[Hálózat beállítása OpenRC init-rendszeren](/wiki/Handbook:X86/Full/Networking/hu "Handbook:X86/Full/Networking/hu")[Munka elkezdése](/wiki/Handbook:X86/Networking/Introduction/hu "Handbook:X86/Networking/Introduction/hu")[Fejlett beállítások](/wiki/Handbook:X86/Networking/Advanced/hu "Handbook:X86/Networking/Advanced/hu")[Moduláris hálózat](/wiki/Handbook:X86/Networking/Modular/hu "Handbook:X86/Networking/Modular/hu")[Vezeték nélküli (Wi-Fi)](/wiki/Handbook:X86/Networking/Wireless/hu "Handbook:X86/Networking/Wireless/hu")[Funkcionalitás hozzáadása](/wiki/Handbook:X86/Networking/Extending/hu "Handbook:X86/Networking/Extending/hu")[Dinamikus menedzsment](/wiki/Handbook:X86/Networking/Dynamic/hu "Handbook:X86/Networking/Dynamic/hu")

Noha 32 bites CPU-hoz telepít, szinte az összes **x86** alaplap (nagyjából 2006-2007-től napjainkig), amely UEFI-támogatással készült, _64 bites UEFI firmware_-rel rendelkezik. Néhány felhasználó észreveheti a "64" számot a beállítások és fájlok nevében az alábbi szakaszokban. Ez _szinte minden esetben_ normális.

Van néhány nagyon apró kivétel a 64 bites UEFI firmware szabály alól, nevezetesen néhány korai [Apple Mac](http://www.everymac.com/systems/apple/mac_pro/specs/mac-pro-quad-2.66-specs.html) és néhány Intel Atom által hajtott [Dell tablet PC](http://www.dell.com/en-us/shop/productdetails/dell-venue-8-pro) támogatta a 32 bites UEFI firmware-t. Az olvasók túlnyomó többsége _soha_ nem fog találkozni 32 bites UEFI firmware-rel a valóságban. Emiatt a 32 bites UEFI firmware nincs tárgyalva a **x86** kézikönyvben.

## Contents

- [1Selecting a boot loader](#Selecting_a_boot_loader)
- [2Default: GRUB](#Default:_GRUB)
  - [2.1Emerge](#Emerge)
  - [2.2Install](#Install)
    - [2.2.1DOS/Legacy BIOS systems](#DOS.2FLegacy_BIOS_systems)
    - [2.2.2UEFI systems](#UEFI_systems)
      - [2.2.2.1Optional: Secure Boot](#Optional:_Secure_Boot)
      - [2.2.2.2Debugging GRUB](#Debugging_GRUB)
  - [2.3Configure](#Configure)
- [3Alternative 1: systemd-boot](#Alternative_1:_systemd-boot)
  - [3.1Emerge](#Emerge_2)
  - [3.2Installation](#Installation)
  - [3.3Optional: Secure Boot](#Optional:_Secure_Boot_2)
- [4Alternative 2: EFI Stub](#Alternative_2:_EFI_Stub)
  - [4.1Unified Kernel Image](#Unified_Kernel_Image)
- [5Other Alternatives](#Other_Alternatives)
- [6Rendszer újraindítása](#Rendszer_.C3.BAjraind.C3.ADt.C3.A1sa)

## Selecting a boot loader

With the Linux kernel configured, system tools installed and configuration files edited, it is time to install the last important piece of a Linux installation: the boot loader.

The boot loader is responsible for firing up the Linux kernel upon boot - without it, the system would not know how to proceed when the power button has been pressed.

For **x86**, we document how to configure [GRUB](/wiki/Handbook:X86/Blocks/Bootloader/hu#Default:_GRUB "Handbook:X86/Blocks/Bootloader/hu") for DOS/Legacy BIOS based systems, and [GRUB](#Default:_GRUB), [systemd-boot](#Alternative_1:_systemd-boot) or [EFI Stub](#Alternative_2:_efibootmgr) for UEFI systems.

In this section of the Handbook a delineation has been made between _emerging_ the boot loader's package and _installing_ a boot loader to a system disk. Here the term _emerge_ will be used to ask [Portage](/wiki/Portage "Portage") to make the software package available to the system. The term _install_ will signify the boot loader copying files or physically modifying appropriate sections of the system's disk drive in order to render the boot loader _activated and ready to operate_ on the next power cycle.

## Default: GRUB

By default, the majority of Gentoo systems now rely upon [GRUB](/wiki/GRUB "GRUB") (found in the [sys-boot/grub](https://packages.gentoo.org/packages/sys-boot/grub) package), which is the direct successor to [GRUB Legacy](/wiki/GRUB_Legacy "GRUB Legacy"). With no additional configuration, GRUB gladly supports older BIOS ("pc") systems. With a small amount of configuration, necessary before build time, GRUB can support more than a half a dozen additional platforms. For more information, consult the [Prerequisites section](/wiki/GRUB#Prerequisites "GRUB") of the [GRUB](/wiki/GRUB "GRUB") article.

### Emerge

When using an older BIOS system supporting only MBR partition tables, no additional configuration is needed in order to emerge GRUB:

`root #` `emerge --ask --verbose sys-boot/grub`

A note for UEFI users: running the above command will output the enabled GRUB\_PLATFORMS values before emerging. When using UEFI capable systems, users will need to ensure `GRUB_PLATFORMS="efi-64"` is enabled (as it is the case by default). If that is not the case for the setup, `GRUB_PLATFORMS="efi-64"` will need to be added to the /etc/portage/make.conf file _before_ emerging GRUB so that the package will be built with EFI functionality:

`root #` `echo 'GRUB_PLATFORMS="efi-64"' >> /etc/portage/make.conf`

`root #` `emerge --ask sys-boot/grub`

If GRUB was somehow emerged without enabling `GRUB_PLATFORMS="efi-64"`, the line (as shown above) can be added to make.conf and then dependencies for the [world package set](/wiki/World_set_(Portage) "World set (Portage)") can be re-calculated by passing the `--update --newuse` options to emerge:

`root #` `emerge --ask --update --newuse --verbose sys-boot/grub`

The GRUB software has now been merged onto the _system_, but it has not yet been installed as a secondary _bootloader_.

### Install

Next, install the necessary GRUB files to the /boot/grub/ directory via the grub-install command. Presuming the first disk (the one where the system boots from) is /dev/sda, one of the following commands will do:

#### DOS/Legacy BIOS systems

For DOS/Legacy BIOS systems:

`root #` `grub-install /dev/sda`

#### UEFI systems

**Important**

Make sure the EFI system partition has been mounted _before_ running grub-install. It is possible for grub-install to install the GRUB EFI file (grubx64.efi) into the wrong directory **without** providing _any_ indication the wrong directory was used.

For UEFI systems:

`root #` `grub-install --efi-directory=/efi`

```
Installing for x86_64-efi platform.
Installation finished. No error reported.

```

Upon successful installation, the output should match the output of the previous command. If the output does not match exactly, then proceed to [Debugging GRUB](/wiki/Handbook:X86/Blocks/Bootloader/hu#Debugging_GRUB "Handbook:X86/Blocks/Bootloader/hu"), otherwise jump to the [Configure step](/wiki/Handbook:X86/Blocks/Bootloader/hu#GRUB_Configure "Handbook:X86/Blocks/Bootloader/hu").

##### Optional: Secure Boot

To successfully boot with secure boot enabled the signing certificate must either be accepted by the [UEFI](/wiki/UEFI "UEFI") firmware, or [shim](/wiki/Shim "Shim") must be used as a pre-loader. Shim is pre-signed with the third-party Microsoft Certificate, accepted by default by most UEFI motherboards.

How to configure the UEFI firmware to accept custom keys depends on the firmware vendor, which is beyond the scope of the handbook. Below is shown how to setup shim instead. Here it is assumed that the user has already followed the instructions in the previous sections to generate a signing key and to configure portage to use it. If this is not the case please return first to the [Kernel installation](/wiki/Handbook:X86/Installation/Kernel/hu "Handbook:X86/Installation/Kernel/hu") section.

The package [sys-boot/grub](https://packages.gentoo.org/packages/sys-boot/grub) installs a prebuilt and signed stand-alone EFI executable if the [secureboot](https://packages.gentoo.org/useflags/secureboot) [USE flag/hu](/wiki/USE_flag/hu "USE flag/hu") USE flag is enabled. Install the required packages and copy the stand-alone grub, Shim, and the MokManager to the same directory on the EFI System Partition. For example:

`root #` `emerge sys-boot/grub sys-boot/shim sys-boot/mokutil sys-boot/efibootmgr`

`root #` `cp /usr/share/shim/BOOTX64.EFI /efi/EFI/Gentoo/shimx64.efi
`

`root #` `cp /usr/share/shim/mmx64.efi /efi/EFI/Gentoo/mmx64.efi
`

`root #` `cp /usr/lib/grub/grub-x86_64.efi.signed /efi/EFI/Gentoo/grubx64.efi
`

Next register the signing key in shims MOKlist, this requires keys in the DER format, whereas sbsign and the kernel build system expect keys in the PEM format. In the previous sections of the handbook an example was shown to generate such a signing PEM key, this key must now be converted to the DER format:

`root #` `openssl x509 -in /path/to/kernel_key.pem -inform PEM -out /path/to/kernel_key.der -outform DER`

**Note**

The path used here must be the path to the pem file containing the certificate belonging to the generated key. In this example both key and certificate are in the same pem file.

Then the converted certificate can be imported into Shims MOKlist, this command will ask to set some password for the import request:

`root #` `mokutil --import /path/to/kernel_key.der`

**Note**

When the currently booted kernel already trusts the certificate being imported, the message "Already in kernel trusted keyring." will be returned here. If this happens, re-run the above command with the argument --ignore-keyring added.

Next, register Shim with the UEFI firmware. In the following command, `boot-disk` and `boot-partition-id` must be replaced with the disk and partition identifier of the EFI system partition:

`root #` `efibootmgr --create --disk /dev/boot-disk --part boot-partition-id --loader '\EFI\Gentoo\shimx64.efi' --label 'GRUB via Shim' --unicode`

Note that this prebuilt and signed stand-alone version of grub reads the grub.cfg from a different location then usual. Instead of the default /boot/grub/grub.cfg the config file should be in the same directory that the grub EFI executable is in, e.g. /efi/EFI/Gentoo/grub.cfg. When [sys-kernel/installkernel](https://packages.gentoo.org/packages/sys-kernel/installkernel) is used to install the kernel and update the grub configuration then the GRUB\_CFG environment variable may be used to override the usual location of the grub config file.

For example:

`root #` `grub-mkconfig -o /efi/EFI/Gentoo/grub.cfg`

Or, via [installkernel](/wiki/Installkernel/hu "Installkernel/hu"):

FILE **`/etc/env.d/99grub`**

```
GRUB_CFG=/efi/EFI/Gentoo/grub.cfg

```

`root #` `env-update`

**Note**

The import process will not be completed until the system is rebooted. After completing all steps in the handbook, restart the system and Shim will load, it will find the import request registered by mokutil. The MokManager application will start and ask for the password that was set when creating the import request. Follow the on-screen instructions to complete the import of the certificate, then reboot the system into the UEFI menu and enable the Secure Boot setting.

##### Debugging GRUB

When debugging GRUB, there are a couple of quick fixes that may result in a bootable installation without having to reboot to a new live image environment.

In the event that "EFI variables are not supported on this system" is displayed somewhere in the output, it is likely the live image was not booted in EFI mode and is presently in Legacy BIOS boot mode. The solution is to try the [removable GRUB step](/wiki/Handbook:X86/Blocks/Bootloader/hu#GRUB_Install_EFI_systems_removable "Handbook:X86/Blocks/Bootloader/hu") mentioned below. This will overwrite the executable EFI file located at /EFI/BOOT/BOOTX64.EFI. Upon rebooting in EFI mode, the motherboard firmware may execute this default boot entry and execute GRUB.

If grub-install returns an error that says "Could not prepare Boot variable: Read-only file system", and the live environment was correctly booted in UEFI mode, then it should be possible to remount the efivars special mount as read-write and then re-run the [aforementioned grub-install command](/wiki/Handbook:X86/Blocks/Bootloader/hu#GRUB_Install_EFI_systems_command "Handbook:X86/Blocks/Bootloader/hu"):

`root #` `mount -o remount,rw,nosuid,nodev,noexec --types efivarfs efivarfs /sys/firmware/efi/efivars`

This is caused by certain non-official Gentoo environments not mounting the special EFI filesystem by default. If the previous command does not run, then reboot using an official Gentoo live image environment in EFI mode.

Some motherboard manufacturers with poor UEFI implementations seem to _only_ support the /EFI/BOOT directory location for the .EFI file in the EFI System Partition (ESP). The GRUB installer can create the .EFI file in this location automatically by appending the `--removable` option to the install command. Ensure the ESP has been mounted before running the following command; presuming it is mounted at /efi (as defined earlier), run:

`root #` `grub-install --target=x86_64-efi --efi-directory=/efi --removable`

This creates the 'default' directory defined by the UEFI specification, and then creates a file with the default name: bootx64.efi.

### Configure

Next, generate the GRUB configuration based on the user configuration specified in the /etc/default/grub file and /etc/grub.d scripts. In most cases, no configuration is needed by users as GRUB will automatically detect which kernel to boot (the highest one available in /boot/) and what the root file system is. It is also possible to append kernel parameters in /etc/default/grub using the GRUB\_CMDLINE\_LINUX variable.

To generate the final GRUB configuration, run the grub-mkconfig command:

`root #` `grub-mkconfig -o /boot/grub/grub.cfg`

```
Generating grub.cfg ...
Found linux image: /boot/vmlinuz-6.18.8-gentoo
Found initrd image: /boot/initramfs-genkernel-x86-6.18.8-gentoo
done

```

The output of the command must mention that at least one Linux image is found, as those are needed to boot the system. If an initramfs is used or genkernel was used to build the kernel, the correct initrd image should be detected as well. If this is not the case, go to /boot/ and check the contents using the ls command. If the files are indeed missing, go back to the kernel configuration and installation instructions.

**Tip**

The os-prober utility can be used in conjunction with GRUB to detect other operating systems from attached drives. Windows 7, 8.1, 10, and other distributions of Linux are detectable. Those desiring dual boot systems should emerge the [sys-boot/os-prober](https://packages.gentoo.org/packages/sys-boot/os-prober) package then re-run the grub-mkconfig command (as seen above). If detection problems are encountered be sure to read the [GRUB](/wiki/GRUB "GRUB") article in its entirety _before_ asking the Gentoo community for support.

## Alternative 1: systemd-boot

Another option is [systemd-boot](/wiki/Systemd/systemd-boot "Systemd/systemd-boot"), which works on both OpenRC and systemd machines. It is a thin chainloader and works well with secure boot.

### Emerge

To install systemd-boot, enable the [boot](https://packages.gentoo.org/useflags/boot) [USE flag/hu](/wiki/USE_flag/hu "USE flag/hu") USE flag and re-install [sys-apps/systemd](https://packages.gentoo.org/packages/sys-apps/systemd) (for systemd systems) or [sys-apps/systemd-utils](https://packages.gentoo.org/packages/sys-apps/systemd-utils) (for OpenRC systems):

FILE **`/etc/portage/package.use/systemd-boot`**

```
sys-apps/systemd boot
sys-apps/systemd-utils boot

```

`root #` `emerge --ask sys-apps/systemd`

Or

`root #` `emerge --ask sys-apps/systemd-utils`

### Installation

Now, install the systemd-boot loader to the [EFI System Partition](/wiki/EFI_System_Partition/hu "EFI System Partition/hu"):

`root #` `bootctl install`

**Important**

Make sure the EFI system partition has been mounted _before_ running bootctl install.

When using this bootloader, before rebooting, verify that a new bootable entry exists using:

`root #` `bootctl list`

**Warning**

The kernel command line for new systemd-boot entries is read from /etc/kernel/cmdline or /usr/lib/kernel/cmdline. If neither file is present, then the kernel command line of the currently booted kernel is re-used (/proc/cmdline). On new installs it might therefore happen that the kernel command line of the live CD is accidentally used to boot the new kernel. The kernel command line for registered entries can be checked with:

`root #` `bootctl list`

If this does not show the desired kernel command line then create /etc/kernel/cmdline containing the correct kernel command line and re-install the kernel.

If no new entry exists, ensure the [sys-kernel/installkernel](https://packages.gentoo.org/packages/sys-kernel/installkernel) package has been installed with the [systemd](https://packages.gentoo.org/useflags/systemd) [USE flag/hu](/wiki/USE_flag/hu "USE flag/hu") and [systemd-boot](https://packages.gentoo.org/useflags/systemd-boot) [USE flag/hu](/wiki/USE_flag/hu "USE flag/hu") USE flags enabled, and re-run the kernel installation.

For the distribution kernels:

`root #` `emerge --ask --config sys-kernel/gentoo-kernel`

For a manually configured and compiled kernel:

`root #` `make install`

**Important**

When installing kernels for systemd-boot, no root= kernel command line argument is added by default. On systemd systems that are using an initramfs users may rely instead on [systemd-gpt-auto-generator](/wiki/Systemd#Automatic_mounting_of_partitions_at_boot.2Fhu "Systemd") to automatically find the root partition at boot. Otherwise users should manually specify the location of the root partition by setting root= in /etc/kernel/cmdline as well as any other kernel command line arguments that should be used. And then reinstalling the kernel as described above.

### Optional: Secure Boot

When the [secureboot](https://packages.gentoo.org/useflags/secureboot) [USE flag/hu](/wiki/USE_flag/hu "USE flag/hu") USE flag is enabled, the systemd-boot EFI executable will be signed by portage automatically. Furthermore, bootctl install will automatically install the signed version.

To successfully boot with secure boot enabled the used certificate must either be accepted by the [UEFI](/wiki/UEFI "UEFI") firmware, or [shim](/wiki/Shim "Shim") must be used as a pre-loader. Shim is pre-signed with the third-party Microsoft Certificate, accepted by default by most UEFI motherboards.

How to configure the UEFI firmware to accept custom keys depends on the firmware vendor, which is beyond the scope of the handbook. Below is shown how to setup shim instead. Here it is assumed that the user has already followed the instructions in the previous sections to generate a signing key and to configure portage to use it. If this is not the case please return first to the [Kernel installation](/wiki/Handbook:X86/Installation/Kernel/hu "Handbook:X86/Installation/Kernel/hu") section.

`root #` `emerge --ask sys-boot/shim sys-boot/mokutil sys-boot/efibootmgr`

`root #` `bootctl install --no-variables`

`root #` `cp /usr/share/shim/BOOTX64.EFI /efi/EFI/systemd/shimx64.efi`

`root #` `cp /usr/share/shim/mmx64.efi /efi/EFI/systemd/mmx64.efi`

Shims MOKlist requires keys in the DER format, whereas sbsign and the kernel build system expect keys in the PEM format. In the previous sections of the handbook an example was shown to generate such a signing PEM key, this key must now be converted to the DER format:

`root #` `openssl x509 -in /path/to/kernel_key.pem -inform PEM -out /path/to/kernel_key.der -outform DER`

**Note**

The path used here must be the path to the pem file containing the certificate belonging to the generated key. In this example both key and certificate are in the same pem file.

Then the converted certificate can be imported into Shims MOKlist:

`root #` `mokutil --import /path/to/kernel_key.der`

**Note**

When the currently booted kernel already trusts the certificate being imported, the message "Already in kernel trusted keyring." will be returned here. If this happens, re-run the above command with the argument --ignore-keyring added.

And finally we register Shim with the UEFI firmware. In the following command, `boot-disk` and `boot-partition-id` must be replaced with the disk and partition identifier of the EFI system partition:

`root #` `efibootmgr --create --disk /dev/boot-disk --part boot-partition-id --loader '\EFI\systemd\shimx64.efi'  --label 'Systemd-boot via Shim' --unicode '\EFI\systemd\systemd-bootx64.efi'`

**Note**

The import process will not be completed until the system is rebooted. After completing all steps in the handbook, restart the system and Shim will load, it will find the import request registered by mokutil. The MokManager application will start and ask for the password that was set when creating the import request. Follow the on-screen instructions to complete the import of the certificate, then reboot the system into the UEFI menu and enable the Secure Boot setting.

## Alternative 2: EFI Stub

Computer systems with UEFI-based firmware technically do not need secondary bootloaders (e.g. GRUB) in order to boot kernels. Secondary bootloaders exist to _extend_ the functionality of UEFI firmware during the boot process. Using GRUB (see the prior section) is typically easier and more robust because it offers a more flexible approach for quickly modifying kernel parameters at boot time.

System administrators who desire to take a minimalist, although more rigid, approach to booting the system can avoid secondary bootloaders and boot the Linux kernel as an [EFI stub](/wiki/EFI_stub/hu "EFI stub/hu").

The [sys-boot/efibootmgr](https://packages.gentoo.org/packages/sys-boot/efibootmgr) application is a tool to used interact with UEFI firmware - the system's primary bootloader. Normally this looks like adding or removing boot entries to the firmware's list of bootable entries. It can also update firmware settings so that the Linux kernels that were previously added as bootable entries can be executed with additional options. These interactions are performed through special data structures called EFI variables (hence the need for kernel support of EFI vars).

Ensure the [EFI stub](/wiki/EFI_stub/hu "EFI stub/hu") kernel article has been reviewed _before_ continuing. The kernel must have specific options enabled to be directly bootable by the UEFI firmware. It may be necessary to recompile the kernel in order to build-in this support.

It is also a good idea to take a look at the [efibootmgr](/wiki/Efibootmgr "Efibootmgr") article for additional information.

**Note**

To reiterate, efibootmgr is _not_ a requirement to boot an UEFI system; it is merely necessary to add an entry for an EFI-stub kernel into the UEFI firmware. When built appropriately with EFI stub support, the Linux kernel itself can be booted _directly_. Additional kernel command-line options can be built-in to the Linux kernel (there is a kernel configuration option called CONFIG\_CMDLINE. Similarly, support for initramfs can be 'built-in' to the kernel as well.

To boot the kernel directly from the firmware, the kernel image must be present on the [EFI System Partition](/wiki/EFI_System_Partition/hu "EFI System Partition/hu"). This may be accomplished by enabling the [efistub](https://packages.gentoo.org/useflags/efistub) [USE flag/hu](/wiki/USE_flag/hu "USE flag/hu") USE flag on [sys-kernel/installkernel](https://packages.gentoo.org/packages/sys-kernel/installkernel). Given that EFI Stub booting is not guaranteed to work on every UEFI system, the USE flag is stable masked, testing keywords must be accepted for installkernel to use this feature.

FILE **`/etc/portage/package.accept_keywords/installkernel`**

```
sys-kernel/installkernel
sys-boot/uefi-mkconfig
app-emulation/virt-firmware

```

FILE **`/etc/portage/package.use/installkernel`**

```
sys-kernel/installkernel efistub

```

Then reinstall [installkernel](/wiki/Installkernel/hu "Installkernel/hu"), create the /efi directory and reinstall the kernel:

`root #` `emerge --ask sys-kernel/installkernel`

`root #` `mkdir -p /efi`

For distribution kernels:

`root #` `emerge --ask --config sys-kernel/gentoo-kernel{,-bin}`

For manually managed kernels:

`root #` `make install`

Alternatively, when [installkernel](/wiki/Installkernel/hu "Installkernel/hu") is not used, the kernel may be copied manually to the EFI System Partition, calling it bootx64.efi:

`root #` `mkdir -p /efi
`

`root #` `cp /boot/vmlinuz-* /efi/bootx64.efi
`

Install the efibootmgr package:

`root #` `emerge --ask sys-boot/efibootmgr`

Create boot entry called "Gentoo" for the freshly compiled EFI stub kernel within the UEFI firmware:

`root #` `efibootmgr --create --disk /dev/sda --part 1 --label "gentoo" --loader "\bootx64.efi"`

**Note**

The use of a backslash ( `\`) as directory path separator is mandatory when using UEFI definitions.

If an initial RAM file system (initramfs) is used, copy it to the EFI System Partition as well, then add the proper boot option to it:

`root #` `efibootmgr --create --disk /dev/sda --part 1 --label "gentoo" --loader "\bootx64.efi" --unicode "initrd=\initramfs.img"`

**Tip**

Additional kernel command line options may be parsed by the firmware to the kernel by specifying them along with the initrd=... option as shown above.

With these changes done, when the system reboots, a boot entry called "gentoo" will be available.

### Unified Kernel Image

If [installkernel](/wiki/Installkernel/hu "Installkernel/hu") was configured to build and install unified kernel images. The unified kernel image should already be installed to the EFI/Linux directory on the EFI system partition, if this is not the case ensure the directory exists and then run the kernel installation again as described earlier in the handbook.

To add a direct boot entry for the installed unified kernel image:

`root #` `efibootmgr --create --disk /dev/sda --part 1 --label "gentoo" --loader "\EFI\Linux\gentoo-x.y.z.efi"`

## Other Alternatives

For other options that are not covered in the Handbook, see the full list of available [bootloaders](/wiki/Bootloader/hu "Bootloader/hu").

## Rendszer újraindítása

Lépjen ki a chrootolt környezetből, és válassza le az összes felcsatolt partíciót. Ezt követően írja be azt az egyetlen mágikus parancsot, amely elindítja a végső, valódi tesztet: reboot.

`(chroot) livecd #` `exit`

`livecd~#` `cd
`

`livecd~#` `umount -l /mnt/gentoo/dev{/shm,/pts,}
`

`livecd~#` `umount -R /mnt/gentoo
`

`livecd~#` `reboot`

Ne feledje el eltávolítani az Live ISO telepítőt, különben ismét elindulhat a számítógépen az újonnan telepített Gentoo rendszer helyett!

Miután újraindította a számítógépet, és belépett a frissen feltelepített Gentoo környezetben, bölcs dolog [véglegesíteni a Gentoo telepítést](/wiki/Handbook:X86/Installation/Finalizing/hu "Handbook:X86/Installation/Finalizing/hu").

[← Eszközök telepítése](/wiki/Handbook:X86/Installation/Tools/hu "Previous part") [Kezdőlap](/wiki/Handbook:X86/hu "Handbook:X86/hu") [Telepítés véglegesítése →](/wiki/Handbook:X86/Installation/Finalizing/hu "Next part")