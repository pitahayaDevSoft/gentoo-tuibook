# Kernel

Other languages:

- [Deutsch](/wiki/Handbook:PPC/Installation/Kernel/de "Handbuch:PPC/Installation/Kernel (100% translated)")
- [English](/wiki/Handbook:PPC/Installation/Kernel "Handbook:PPC/Installation/Kernel (100% translated)")
- [español](/wiki/Handbook:PPC/Installation/Kernel/es "Manual de Gentoo: PPC/Instalación/Núcleo (100% translated)")
- [français](/wiki/Handbook:PPC/Installation/Kernel/fr "Handbook:PPC/Installation/Kernel/fr (100% translated)")
- [italiano](/wiki/Handbook:PPC/Installation/Kernel/it "Handbook:PPC/Installation/Kernel/it (100% translated)")
- [magyar](/wiki/Handbook:PPC/Installation/Kernel/hu "Handbook:PPC/Installation/Kernel/hu (100% translated)")
- [polski](/wiki/Handbook:PPC/Installation/Kernel/pl "Handbook:PPC/Installation/Kernel (100% translated)")
- [português do Brasil](/wiki/Handbook:PPC/Installation/Kernel/pt-br "Handbook:PPC/Installation/Kernel/pt-br (50% translated)")
- [русский](/wiki/Handbook:PPC/Installation/Kernel/ru "Handbook:PPC/Installation/Kernel (100% translated)")
- [தமிழ்](/wiki/Handbook:PPC/Installation/Kernel/ta "கையேடு:PPC/நிறுவல்/கர்னல் (100% translated)")
- [中文（中国大陆）‎](/wiki/Handbook:PPC/Installation/Kernel/zh-cn "手册：PPC/安装/配置Linux内核 (100% translated)")
- [日本語](/wiki/Handbook:PPC/Installation/Kernel/ja "ハンドブック:PPC/インストール/カーネル (100% translated)")
- 한국어

[PPC 핸드북](/wiki/Handbook:PPC/ko "Handbook:PPC/ko")[설치](/wiki/Handbook:PPC/Full/Installation/ko "Handbook:PPC/Full/Installation/ko")[설치 정보](/wiki/Handbook:PPC/Installation/About/ko "Handbook:PPC/Installation/About/ko")[매체 선택](/wiki/Handbook:PPC/Installation/Media/ko "Handbook:PPC/Installation/Media/ko")[네트워크 설정](/wiki/Handbook:PPC/Installation/Networking/ko "Handbook:PPC/Installation/Networking/ko")[디스크 준비](/wiki/Handbook:PPC/Installation/Disks/ko "Handbook:PPC/Installation/Disks/ko")[스테이지 3 설치](/wiki/Handbook:PPC/Installation/Stage/ko "Handbook:PPC/Installation/Stage/ko")[베이스 시스템 설치](/wiki/Handbook:PPC/Installation/Base/ko "Handbook:PPC/Installation/Base/ko")커널 설정[시스템 설정](/wiki/Handbook:PPC/Installation/System/ko "Handbook:PPC/Installation/System/ko")[도구 설치](/wiki/Handbook:PPC/Installation/Tools/ko "Handbook:PPC/Installation/Tools/ko")[부트로더 설정](/wiki/Handbook:PPC/Installation/Bootloader/ko "Handbook:PPC/Installation/Bootloader/ko")[마무리](/wiki/Handbook:PPC/Installation/Finalizing/ko "Handbook:PPC/Installation/Finalizing/ko")[젠투 활용](/wiki/Handbook:PPC/Full/Working/ko "Handbook:PPC/Full/Working/ko")[포티지 소개](/wiki/Handbook:PPC/Working/Portage/ko "Handbook:PPC/Working/Portage/ko")[USE 플래그](/wiki/Handbook:PPC/Working/USE/ko "Handbook:PPC/Working/USE/ko")[포티지 기능](/wiki/Handbook:PPC/Working/Features/ko "Handbook:PPC/Working/Features/ko")[초기화 스크립트 시스템](/wiki/Handbook:PPC/Working/Initscripts/ko "Handbook:PPC/Working/Initscripts/ko")[환경 변수](/wiki/Handbook:PPC/Working/EnvVar/ko "Handbook:PPC/Working/EnvVar/ko")[포티지 활용](/wiki/Handbook:PPC/Full/Portage/ko "Handbook:PPC/Full/Portage/ko")[파일 및 디렉터리](/wiki/Handbook:PPC/Portage/Files/ko "Handbook:PPC/Portage/Files/ko")[변수](/wiki/Handbook:PPC/Portage/Variables/ko "Handbook:PPC/Portage/Variables/ko")[소프트웨어 브랜치 함께 사용하기](/wiki/Handbook:PPC/Portage/Branches/ko "Handbook:PPC/Portage/Branches/ko")[추가 도구](/wiki/Handbook:PPC/Portage/Tools/ko "Handbook:PPC/Portage/Tools/ko")[꾸러미 저장소 개별 설정](/wiki/Handbook:PPC/Portage/CustomTree/ko "Handbook:PPC/Portage/CustomTree/ko")[고급 기능](/wiki/Handbook:PPC/Portage/Advanced/ko "Handbook:PPC/Portage/Advanced/ko")[네트워크 설정](/wiki/Handbook:PPC/Full/Networking/ko "Handbook:PPC/Full/Networking/ko")[시작하기](/wiki/Handbook:PPC/Networking/Introduction/ko "Handbook:PPC/Networking/Introduction/ko")[고급 설정](/wiki/Handbook:PPC/Networking/Advanced/ko "Handbook:PPC/Networking/Advanced/ko")[모듈러 네트워크](/wiki/Handbook:PPC/Networking/Modular/ko "Handbook:PPC/Networking/Modular/ko")[무선 네트워크](/wiki/Handbook:PPC/Networking/Wireless/ko "Handbook:PPC/Networking/Wireless/ko")[기능 추가](/wiki/Handbook:PPC/Networking/Extending/ko "Handbook:PPC/Networking/Extending/ko")[동적 관리](/wiki/Handbook:PPC/Networking/Dynamic/ko "Handbook:PPC/Networking/Dynamic/ko")

## Contents

- [1선택: 펌웨어 설치](#.EC.84.A0.ED.83.9D:_.ED.8E.8C.EC.9B.A8.EC.96.B4_.EC.84.A4.EC.B9.98)
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
  - [4.1Distribution kernels](#Distribution_kernels)
    - [4.1.1Optional: Signed kernel modules](#Optional:_Signed_kernel_modules)
    - [4.1.2Installing a distribution kernel](#Installing_a_distribution_kernel)
    - [4.1.3Upgrading and cleaning up](#Upgrading_and_cleaning_up)
    - [4.1.4Post-install/upgrade tasks](#Post-install.2Fupgrade_tasks)
      - [4.1.4.1Manually rebuilding the initramfs or Unified Kernel Image](#Manually_rebuilding_the_initramfs_or_Unified_Kernel_Image)
- [5기본: 직접 설정](#.EA.B8.B0.EB.B3.B8:_.EC.A7.81.EC.A0.91_.EC.84.A4.EC.A0.95)
- [6소스 코드 설치](#.EC.86.8C.EC.8A.A4_.EC.BD.94.EB.93.9C_.EC.84.A4.EC.B9.98)
  - [6.1Option 1 - Modprobed-db process](#Option_1_-_Modprobed-db_process)
  - [6.2Option 2 - Assisted manual process](#Option_2_-_Assisted_manual_process)
  - [6.3Option 3 - Configuring by hand](#Option_3_-_Configuring_by_hand)
  - [6.4Optional: Signed kernel modules](#Optional:_Signed_kernel_modules_2)
  - [6.5컴파일 및 설치](#.EC.BB.B4.ED.8C.8C.EC.9D.BC_.EB.B0.8F_.EC.84.A4.EC.B9.98)

### 선택: 펌웨어 설치

### Firmware

#### Suggested: Linux Firmware

On many systems, non-FOSS firmware is required for certain hardware to function. The [sys-kernel/linux-firmware](https://packages.gentoo.org/packages/sys-kernel/linux-firmware) package contains firmware for many, but not all, devices.

**요령**

Most wireless cards and GPUs require firmware to function.

`root #` `emerge --ask sys-kernel/linux-firmware`

**참고**

Installing certain firmware packages often requires accepting the associated firmware licenses. If necessary, visit the [license handling section](/wiki/Handbook:PPC/Working/Portage/ko#Licenses "Handbook:PPC/Working/Portage/ko") of the Handbook for help on accepting licenses.

##### Firmware Loading

Firmware files are typically loaded when the associated kernel module is loaded. This means the firmware must be built into the kernel using **CONFIG\_EXTRA\_FIRMWARE** if the kernel module is set to _Y_ instead of _M_. In most cases, building-in a module which required firmware can complicate or break loading.

## sys-kernel/installkernel

[Installkernel](/wiki/Installkernel "Installkernel") may be used to automate the kernel installation, [initramfs](/wiki/Initramfs "Initramfs") generation, [unified kernel image](/wiki/Unified_kernel_image "Unified kernel image") generation and bootloader configuration, among other things. [sys-kernel/installkernel](https://packages.gentoo.org/packages/sys-kernel/installkernel) implements two paths of achieving this: the traditional installkernel originating from Debian and [systemd](/wiki/Systemd/ko "Systemd/ko")'s kernel-install. Which one to choose depends, among other things, on the system's bootloader. By default, systemd's kernel-install is used on systemd profiles, while the traditional installkernel is the default for other profiles.

### Bootloader

Now is the time to think about which bootloader the user wants for the system.

**중요**

Only one selection is required in the following subsection, if unsure of which to use then go with the first listed for now. It's always possible to switch at a later date if required.

#### GRUB

Users of GRUB can use either systemd's kernel-install or the traditional Debian installkernel. The [systemd](https://packages.gentoo.org/useflags/systemd) [USE flag/ko (page does not exist)](/index.php?title=USE_flag/ko&action=edit&redlink=1 "USE flag/ko (page does not exist)") USE flag switches between these implementations. To automatically run grub-mkconfig when installing the kernel, enable the [grub](https://packages.gentoo.org/useflags/grub) [USE flag/ko (page does not exist)](/index.php?title=USE_flag/ko&action=edit&redlink=1 "USE flag/ko (page does not exist)") [USE flag](/wiki/USE_flag "USE flag").

**참고**

GRUB requires kernels to be installed to /boot.

파일 **`/etc/portage/package.use/installkernel`**

```
sys-kernel/installkernel grub

```

`root #` `emerge --ask sys-kernel/installkernel`

**참고**

systemd-boot requires kernels to be installed to /efi.

**참고**

When [app-emulation/virt-firmware](https://packages.gentoo.org/packages/app-emulation/virt-firmware) is used to configure the UEFI ensure that the kernel-bootcfg-boot-successful service is enabled before attempting to install the kernel. This implementation of EFIstub booting is the default for systemd systems.

`root #` `systemctl enable kernel-bootcfg-boot-successful.service`

**참고**

EFIstub requires kernels to be installed to /efi.

#### Traditional layout, other bootloaders (e.g. (e)lilo, syslinux, etc.)

The traditional /boot layout (for e.g. (e)LILO, syslinux, etc.) is used by default if the [grub](https://packages.gentoo.org/useflags/grub) [USE flag/ko (page does not exist)](/index.php?title=USE_flag/ko&action=edit&redlink=1 "USE flag/ko (page does not exist)"), [systemd-boot](https://packages.gentoo.org/useflags/systemd-boot) [USE flag/ko (page does not exist)](/index.php?title=USE_flag/ko&action=edit&redlink=1 "USE flag/ko (page does not exist)"), [efistub](https://packages.gentoo.org/useflags/efistub) [USE flag/ko (page does not exist)](/index.php?title=USE_flag/ko&action=edit&redlink=1 "USE flag/ko (page does not exist)") and [uki](https://packages.gentoo.org/useflags/uki) [USE flag/ko (page does not exist)](/index.php?title=USE_flag/ko&action=edit&redlink=1 "USE flag/ko (page does not exist)") USE flags are **not** enabled. No further action is required.

### Initramfs

An **init** ial **ram**-based **f** ile **s** ystem, or [initramfs](/wiki/Initramfs "Initramfs"), may be required for a system to boot. A wide of variety of cases may necessitate one, but common cases include:

- Kernels where storage/filesystem drivers are modules.
- Layouts with /usr/ or /var/ on separate partitions.
- Encrypted root filesystems.

**요령**

[Distribution kernels](/wiki/Project:Distribution_Kernel "Project:Distribution Kernel") are designed to be used with an initramfs, as many storage and filesystem drivers are built as modules.

In addition to mounting the root filesystem, an initramfs may also perform other tasks such as:

- Running **f** ile **s** ystem **c** onsistency chec **k** fsck, a tool to check and repair consistency of a file system in such events of uncleanly shutdown a system.
- Providing a recovery environment in the event of late-boot failures.

[Installkernel](/wiki/Installkernel "Installkernel") can automatically generate an initramfs when installing the kernel if the [dracut](https://packages.gentoo.org/useflags/dracut) [USE flag/ko (page does not exist)](/index.php?title=USE_flag/ko&action=edit&redlink=1 "USE flag/ko (page does not exist)") or [ugrd](https://packages.gentoo.org/useflags/ugrd) [USE flag/ko (page does not exist)](/index.php?title=USE_flag/ko&action=edit&redlink=1 "USE flag/ko (page does not exist)") USE flag is enabled:

파일 **`/etc/portage/package.use/installkernel`**

```
sys-kernel/installkernel dracut

```

#### Chroot detection

Bootloaders such as [systemd-boot](/wiki/Systemd/systemd-boot "Systemd/systemd-boot") and [EFI stub](/wiki/EFI_stub "EFI stub") use the kernel arguments of the running system as set in /proc/cmdline by default. Because of the wide range of ways Gentoo can be installed users will randomly get tripped up by this.

To help solve any issues this may cause, [sys-kernel/installkernel](https://packages.gentoo.org/packages/sys-kernel/installkernel) will check if it is running in a chroot and fail if the kernel command line is not explicitly configured. Otherwise the bootloader would use the install media's boot arguments which would lead to boot failure.

One way to satisfy [sys-kernel/installkernel](https://packages.gentoo.org/packages/sys-kernel/installkernel) is by using Dracut's config file to set the root partition UUID as shown below, or alternatively for more information on what this check helps with and different ways to configure it, see [Installkernel#Install\_chroot\_check](/wiki/Installkernel#Install_chroot_check.2Fko "Installkernel").

`root #` `mkdir /etc/dracut.conf.d`

`root #` `blkid`

```
/dev/sda3: UUID="2122cd72-94d7-4dcc-821e-3705926deecc"
```

In the above example, the root partition is /dev/sda3 and the UUID is 2122cd72-94d7-4dcc-821e-3705926deecc.
The dracut config file would then look like:

파일 **`/etc/dracut.conf.d/00-installkernel.conf`**

```
kernel_cmdline=" root=UUID=2122cd72-94d7-4dcc-821e-3705926deecc " # Note leading and trailing spaces

```

`root #` `emerge --ask sys-kernel/installkernel`

## Kernel selection

It can be a wise move to use the dist-kernel on the first boot as it provides a very simple method to rule out system issues and kernel config issues. Always having a known working kernel to fallback on can speed up debugging and alleviate anxiety when updating that your system will no longer boot.

A common misconception is that a manually compiled kernel will use a lot less RAM than a pre configured distribution kernel. Due to the modular nature of the Linux kernel, only what is needed by the system is loaded and in most cases resulting in similar memory usage.

**중요**

Only one selection is required in the following subsection, if unsure of which to use then go with the first listed for now. It's always possible to switch at a later date if required.

Ranked from least involved to most involved:

1. 직접 설정하고 빌드하는 방법, 또는
2. genkernel 도구를 사용하여 자동으로 리눅스 커널을 빌드하고 설치하는 방법

주변에 빌드한 모든 배포판의 핵심은 리눅스 커널입니다. 이는 사용자 프로그램과 여러분의 시스템 하드웨어 사이에 있는 계층입니다. 젠투는 사용자에게 최대한 다양한 커널 소스코드를 제공합니다. 설명을 포함한 전체 목록은 [커널 개요 페이지](/index.php?title=Kernel/Overview/ko&action=edit&redlink=1 "Kernel/Overview/ko (page does not exist)") 에 있습니다.

**요령**

Kernel installation tasks such as copying the kernel image to /boot or the [EFI System Partition](/wiki/EFI_System_Partition "EFI System Partition"), generating an [initramfs](/wiki/Initramfs "Initramfs") and/or [Unified Kernel Image](/wiki/Unified_Kernel_Image "Unified Kernel Image"), updating bootloader configuration, can be automated with [installkernel](/wiki/Installkernel "Installkernel"). Users may wish to configure and install [sys-kernel/installkernel](https://packages.gentoo.org/packages/sys-kernel/installkernel) before proceeding. See the [Kernel installation section below](/wiki/Handbook:PPC/Installation/Kernel#Kernel_installation.2Fko "Handbook:PPC/Installation/Kernel") for more more information.

### Distribution kernels

_[Distribution Kernels](/wiki/Project:Distribution_Kernel "Project:Distribution Kernel")_ are ebuilds that cover the complete process of unpacking, configuring, compiling, and installing the kernel. The primary advantage of this method is that the kernels are updated to new versions by the package manager as part of @world upgrade. This requires no more involvement than running an emerge command. Distribution kernels default to a configuration supporting the majority of hardware, however two mechanisms are offered for customization: savedconfig and config snippets. See the project page for [more details on configuration.](/wiki/Project:Distribution_Kernel#Modifying_kernel_configuration "Project:Distribution Kernel")

##### Optional: Signed kernel modules

The kernel modules in the prebuilt distribution kernel ([sys-kernel/gentoo-kernel-bin](https://packages.gentoo.org/packages/sys-kernel/gentoo-kernel-bin)) are already signed. To sign the modules of kernels built from source enable the [modules-sign](https://packages.gentoo.org/useflags/modules-sign) [USE flag/ko (page does not exist)](/index.php?title=USE_flag/ko&action=edit&redlink=1 "USE flag/ko (page does not exist)") USE flag, and optionally specify which key to use for signing in [/etc/portage/make.conf](/wiki//etc/portage/make.conf "/etc/portage/make.conf"):

파일 **`/etc/portage/make.conf`** **Enable module signing**

```
USE="modules-sign"

# Optionally, to use custom signing keys.
MODULES_SIGN_KEY="/path/to/kernel_key.pem"
MODULES_SIGN_CERT="/path/to/kernel_key.pem" # Only required if the MODULES_SIGN_KEY does not also contain the certificate.
MODULES_SIGN_HASH="sha512" # Defaults to sha512.

```

If MODULES\_SIGN\_KEY is not specified the kernel build system will generate a key, it will be stored in /usr/src/linux-x.y.z/certs. It is recommended to manually generate a key to ensure that it will be the same for each kernel release. A key may be generated with:

`root #` `openssl req -new -noenc -utf8 -sha256 -x509 -outform PEM -out kernel_key.pem -keyout kernel_key.pem`

**참고**

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

**중요**

_[Distribution Kernels](/wiki/Project:Distribution_Kernel "Project:Distribution Kernel")_, such as [sys-kernel/gentoo-kernel](https://packages.gentoo.org/packages/sys-kernel/gentoo-kernel) and [sys-kernel/gentoo-kernel-bin](https://packages.gentoo.org/packages/sys-kernel/gentoo-kernel-bin), by default, expect to be installed alongside an [initramfs](/wiki/Handbook:PPC/Installation/Kernel#Initramfs.2Fko "Handbook:PPC/Installation/Kernel"). Before running emerge to install the kernel users should ensure that [sys-kernel/installkernel](https://packages.gentoo.org/packages/sys-kernel/installkernel) has been configured to utilize an initramfs generator (for example [Dracut](/wiki/Dracut "Dracut")) as described in the [installkernel section](/wiki/Handbook:PPC/Installation/Kernel#Initramfs.2Fko "Handbook:PPC/Installation/Kernel").

#### Upgrading and cleaning up

Once the kernel is installed, the package manager will automatically update it to newer versions. The previous versions will be kept until the package manager is requested to clean up stale packages. To reclaim disk space, stale packages can be trimmed by periodically running emerge with the `--depclean` option:

`root #` `emerge --depclean`

Alternatively, to specifically clean up old kernel versions:

`root #` `emerge --prune sys-kernel/gentoo-kernel sys-kernel/gentoo-kernel-bin`

**요령**

By design, emerge only removes the kernel build directory. It does not actually remove the kernel modules, nor the installed kernel image. To completely clean-up old kernels, the [app-admin/eclean-kernel](https://packages.gentoo.org/packages/app-admin/eclean-kernel) tool may be used.

#### Post-install/upgrade tasks

An upgrade of a distribution kernel is capable of triggering an automatic rebuild for external kernel modules installed by other packages (for example: [sys-fs/zfs-kmod](https://packages.gentoo.org/packages/sys-fs/zfs-kmod) or [x11-drivers/nvidia-drivers](https://packages.gentoo.org/packages/x11-drivers/nvidia-drivers)). This automated behaviour is enabled by enabling the [dist-kernel](https://packages.gentoo.org/useflags/dist-kernel) [USE flag/ko (page does not exist)](/index.php?title=USE_flag/ko&action=edit&redlink=1 "USE flag/ko (page does not exist)") USE flag. When required, this same flag will also trigger re-generation of the [initramfs](/wiki/Initramfs "Initramfs").

It is highly recommended to enable this flag globally via /etc/portage/make.conf when using a distribution kernel:

파일 **`/etc/portage/make.conf`** **Enabling USE=dist-kernel**

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

After installing the Distribution Kernel successfully, it is now time to proceed to the next section: [Configuring the system](/wiki/Handbook:PPC/Installation/System/ko "Handbook:PPC/Installation/System/ko").

## 기본: 직접 설정

## 소스 코드 설치

When manually compiling the kernel for ppc-based systems, Gentoo recommends the [sys-kernel/gentoo-sources](https://packages.gentoo.org/packages/sys-kernel/gentoo-sources) package.

Choose an appropriate kernel source and install it using emerge:

`root #` `emerge --ask sys-kernel/gentoo-sources`

/usr/src를 들여다보면 설치한 커널 소스를 가리키는 linux 심볼릭 링크를 볼 수 있습니다:

It is conventional for a /usr/src/linux symlink to be maintained, such that it refers to whichever sources correspond with the currently running kernel. However, this symbolic link will not be created by default. An easy way to create the symbolic link is to utilize eselect's kernel module.

For further information regarding the purpose of the symlink, and how to manage it, please refer to [Kernel/Upgrade](/wiki/Kernel/Upgrade/ko "Kernel/Upgrade/ko").

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

**중요**

Only one selection is required in the following subsection, if unsure of which to use then go with the first listed for now. It's always possible to switch at a later date if required.

##### Option 1 - Modprobed-db process

A very easy way to manage the kernel is to first install [sys-kernel/gentoo-kernel-bin](https://packages.gentoo.org/packages/sys-kernel/gentoo-kernel-bin) and use the [sys-kernel/modprobed-db](https://packages.gentoo.org/packages/sys-kernel/modprobed-db) to collect information about what the system requires. modprobed-db is a tool which monitors the system via crontab to add all modules of all devices over the system's life to make sure it everything a user needs is supported. For example, if an Xbox controller is added after installation, then modprobed-db will add the modules to be built next time the kernel is rebuilt.

For now please follow installing [sys-kernel/gentoo-kernel-bin](https://packages.gentoo.org/packages/sys-kernel/gentoo-kernel-bin) via [Distribution\_kernels](/wiki/Handbook:PPC/Installation/Kernel#Distribution_kernels.2Fko "Handbook:PPC/Installation/Kernel").

Next, install [sys-kernel/modprobed-db](https://packages.gentoo.org/packages/sys-kernel/modprobed-db):

`root #` `emerge --ask sys-kernel/modprobed-db`

Please watch out for further steps related to modprobed-db in the Handbook.

More on this topic can be found in the [Modprobed-db](/wiki/Modprobed-db "Modprobed-db") article.

##### Option 2 - Assisted manual process

This method allows a user to have full control of how their kernel is built with as minimal help from outside tools as they wish. Some could consider this as making it hard for the sake of it.

그러나 맞는 이야기이기도 합니다. 커널을 직접 설정했을 때 시스템을 알아둘 필요가 있습니다. 대부분의 정보는 lspci 명령이 들어있는 [sys-apps/pciutils](https://packages.gentoo.org/packages/sys-apps/pciutils)를 이머지하여 수집할 수 있습니다:

`root #` `emerge --ask sys-apps/pciutils`

**참고**

chroot를 하고 나면, lspci가 출력하는 ( _pcilib: cannot open /sys/bus/pci/devices_ 와 같은) pcilib 경고를 무시하는게 안전합니다.

시스템 정보를 알아볼 수 있는 또 다른 부분은 설치 CD에서 사용하는 커널 모듈이 무엇인지 보여주는 lsmod를 실행했을 때 나타나는 활성화 할 모듈에 대한 바람직한 실마리입니다.

이제 커널 소스 디렉터리로 이동하여 make menuconfig를 실행하십시오. 메뉴 기반 설정 화면을 실행합니다.

`root #` `cd /usr/src/linux
`

`root #` `make menuconfig`

**요령**

To view the full list of make arguments available for the kernel, run `make help`.

The kernel has a method of autodetecting the modules currently being used on the installcd which will give a great starting point to allow a user to configure their own. This can be called by using:

`root #` `make localmodconfig`

Manually configuring should not be needed at this point, but if a user wish to check:

`root #` `make nconfig`

Now it's time to decide if modules signing is required in the steps listed in [here](/wiki/Handbook:PPC/Installation/Kernel#Optional:_Signed_kernel_modules_2.2Fko "Handbook:PPC/Installation/Kernel")

If not, proceed to building described [here](/wiki/Handbook:PPC/Installation/Kernel#Compiling_and_installing.2Fko "Handbook:PPC/Installation/Kernel")

##### Option 3 - Configuring by hand

The Linux kernel configuration has many, many sections and as configuring a kernel by hand is rarely needed thanks to the two tools listed above. The steps to do it by hand are now included at [Kernel/Gentoo\_Kernel\_Configuration\_Guide](/wiki/Kernel/Gentoo_Kernel_Configuration_Guide/ko "Kernel/Gentoo Kernel Configuration Guide/ko")

#### Optional: Signed kernel modules

To automatically sign the kernel modules enable CONFIG\_MODULE\_SIG\_ALL:

커널 **Sign kernel modules CONFIG\_MODULE\_SIG\_ALL**

```
[*] Enable loadable module support
  -*-   Module signature verification
    [*]     Automatically sign all modules
    Which hash algorithm should modules be signed with? (Sign modules with SHA-512) --->

```

Optionally change the hash algorithm if desired.

To enforce that all modules are signed with a valid signature, enable CONFIG\_MODULE\_SIG\_FORCE as well:

커널 **Enforce signed kernel modules CONFIG\_MODULE\_SIG\_FORCE**

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

커널 **Specify signing key CONFIG\_MODULE\_SIG\_KEY**

```
-*- Cryptographic API  --->
  Certificates for signature checking  --->
    (/path/to/kernel_key.pem) File name or PKCS#11 URI of module signing key

```

To also sign external kernel modules installed by other packages via `linux-mod-r1.eclass`, enable the [modules-sign](https://packages.gentoo.org/useflags/modules-sign) [USE flag/ko (page does not exist)](/index.php?title=USE_flag/ko&action=edit&redlink=1 "USE flag/ko (page does not exist)") USE flag globally:

파일 **`/etc/portage/make.conf`** **Enable module signing**

```
USE="modules-sign"
</div>

<div lang="en" dir="ltr" class="mw-content-ltr">
# Optionally, when using custom signing keys.
MODULES_SIGN_KEY="/path/to/kernel_key.pem"
MODULES_SIGN_CERT="/path/to/kernel_key.pem" # Only required if the MODULES_SIGN_KEY does not also contain the certificate
MODULES_SIGN_HASH="sha512" # Defaults to sha512

```

**참고**

MODULES\_SIGN\_KEY and MODULES\_SIGN\_CERT may point to different files. For this example, the pem file generated by OpenSSL includes both the key and the accompanying certificate, and thus both variables are set to the same value.

### 컴파일 및 설치

커널을 설정했다면, 컴파일하고 설치할 차례입니다. 설정 메뉴를 빠져나가고 다음 명령을 실행하십시오:

`root #` `make && make modules_install`

**참고**

`make -jX` 명령을 사용하고 X에 실행 가능토록 허용할 빌드 프로세스 갯수를 넣어 병렬 빌드를 활성화 할 수 있습니다. 이는 앞서 언급한 /etc/portage/make.conf의 `MAKEOPTS` 변수와 비슷합니다.

커널 컴파일이 끝나면 아래와 같이 /boot/ 에 커널 이미지를 복사하십시오. 부트 파티션이 나누어져 있다면, 페가소스 컴퓨터에서는 제대로 마운트했는지 확인하십시오. BootX를 부팅할 때 사용한다면, 커널을 나중에 복사하겠습니다.

`root #` `make install`

This command will copy the kernel image to /boot. If [sys-kernel/installkernel](https://packages.gentoo.org/packages/sys-kernel/installkernel) is installed it will call /sbin/installkernel instead and delegate the kernel installation. Instead of simply copying the kernel to /boot, [Installkernel](/wiki/Installkernel "Installkernel") installs each kernel with its version number in the file name. Additionally, installkernel provides a framework for automatically accomplishing various tasks relating to kernel installation, such as: generating an [initramfs](/wiki/Initramfs "Initramfs"), building an [Unified Kernel Image](/wiki/Unified_Kernel_Image "Unified Kernel Image"), and updating the [bootloader](/wiki/Bootloader "Bootloader") configuration.

[시스템 설정](/wiki/Handbook:PPC/Installation/System/ko "Handbook:PPC/Installation/System/ko") 으로 설치 과정을 계속 진행하십시오.

[← 젠투 베이스 시스템 설치](/wiki/Handbook:PPC/Installation/Base/ko "이전 내용") [처음](/wiki/Handbook:PPC/ko "Handbook:PPC/ko") [시스템 설정 →](/wiki/Handbook:PPC/Installation/System/ko "다음 내용")