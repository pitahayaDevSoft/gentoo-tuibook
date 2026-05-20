# Kernel

Other languages:

- [Deutsch](/wiki/Handbook:HPPA/Installation/Kernel/de "Handbuch:HPPA/Installation/Kernel (100% translated)")
- [English](/wiki/Handbook:HPPA/Installation/Kernel "Handbook:HPPA/Installation/Kernel (100% translated)")
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
- 한국어

[HPPA 핸드북](/wiki/Handbook:HPPA/ko "Handbook:HPPA/ko")[설치](/wiki/Handbook:HPPA/Full/Installation/ko "Handbook:HPPA/Full/Installation/ko")[설치 정보](/wiki/Handbook:HPPA/Installation/About/ko "Handbook:HPPA/Installation/About/ko")[매체 선택](/wiki/Handbook:HPPA/Installation/Media/ko "Handbook:HPPA/Installation/Media/ko")[네트워크 설정](/wiki/Handbook:HPPA/Installation/Networking/ko "Handbook:HPPA/Installation/Networking/ko")[디스크 준비](/wiki/Handbook:HPPA/Installation/Disks/ko "Handbook:HPPA/Installation/Disks/ko")[스테이지 3 설치](/wiki/Handbook:HPPA/Installation/Stage/ko "Handbook:HPPA/Installation/Stage/ko")[베이스 시스템 설치](/wiki/Handbook:HPPA/Installation/Base/ko "Handbook:HPPA/Installation/Base/ko")커널 설정[시스템 설정](/wiki/Handbook:HPPA/Installation/System/ko "Handbook:HPPA/Installation/System/ko")[도구 설치](/wiki/Handbook:HPPA/Installation/Tools/ko "Handbook:HPPA/Installation/Tools/ko")[부트로더 설정](/wiki/Handbook:HPPA/Installation/Bootloader/ko "Handbook:HPPA/Installation/Bootloader/ko")[마무리](/wiki/Handbook:HPPA/Installation/Finalizing/ko "Handbook:HPPA/Installation/Finalizing/ko")[젠투 활용](/wiki/Handbook:HPPA/Full/Working/ko "Handbook:HPPA/Full/Working/ko")[포티지 소개](/wiki/Handbook:HPPA/Working/Portage/ko "Handbook:HPPA/Working/Portage/ko")[USE 플래그](/wiki/Handbook:HPPA/Working/USE/ko "Handbook:HPPA/Working/USE/ko")[포티지 기능](/wiki/Handbook:HPPA/Working/Features/ko "Handbook:HPPA/Working/Features/ko")[초기화 스크립트 시스템](/wiki/Handbook:HPPA/Working/Initscripts/ko "Handbook:HPPA/Working/Initscripts/ko")[환경 변수](/wiki/Handbook:HPPA/Working/EnvVar/ko "Handbook:HPPA/Working/EnvVar/ko")[포티지 활용](/wiki/Handbook:HPPA/Full/Portage/ko "Handbook:HPPA/Full/Portage/ko")[파일 및 디렉터리](/wiki/Handbook:HPPA/Portage/Files/ko "Handbook:HPPA/Portage/Files/ko")[변수](/wiki/Handbook:HPPA/Portage/Variables/ko "Handbook:HPPA/Portage/Variables/ko")[소프트웨어 브랜치 함께 사용하기](/wiki/Handbook:HPPA/Portage/Branches/ko "Handbook:HPPA/Portage/Branches/ko")[추가 도구](/wiki/Handbook:HPPA/Portage/Tools/ko "Handbook:HPPA/Portage/Tools/ko")[꾸러미 저장소 개별 설정](/wiki/Handbook:HPPA/Portage/CustomTree/ko "Handbook:HPPA/Portage/CustomTree/ko")[고급 기능](/wiki/Handbook:HPPA/Portage/Advanced/ko "Handbook:HPPA/Portage/Advanced/ko")[네트워크 설정](/wiki/Handbook:HPPA/Full/Networking/ko "Handbook:HPPA/Full/Networking/ko")[시작하기](/wiki/Handbook:HPPA/Networking/Introduction/ko "Handbook:HPPA/Networking/Introduction/ko")[고급 설정](/wiki/Handbook:HPPA/Networking/Advanced/ko "Handbook:HPPA/Networking/Advanced/ko")[모듈러 네트워크](/wiki/Handbook:HPPA/Networking/Modular/ko "Handbook:HPPA/Networking/Modular/ko")[무선 네트워크](/wiki/Handbook:HPPA/Networking/Wireless/ko "Handbook:HPPA/Networking/Wireless/ko")[기능 추가](/wiki/Handbook:HPPA/Networking/Extending/ko "Handbook:HPPA/Networking/Extending/ko")[동적 관리](/wiki/Handbook:HPPA/Networking/Dynamic/ko "Handbook:HPPA/Networking/Dynamic/ko")

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
- [5기본: 직접 설정](#.EA.B8.B0.EB.B3.B8:_.EC.A7.81.EC.A0.91_.EC.84.A4.EC.A0.95)
- [6소스 코드 설치](#.EC.86.8C.EC.8A.A4_.EC.BD.94.EB.93.9C_.EC.84.A4.EC.B9.98)
  - [6.1Option 2 - Assisted manual process](#Option_2_-_Assisted_manual_process)
  - [6.2Option 3 - Configuring by hand](#Option_3_-_Configuring_by_hand)
  - [6.3Optional: Signed kernel modules](#Optional:_Signed_kernel_modules)
  - [6.4컴파일 및 설치](#.EC.BB.B4.ED.8C.8C.EC.9D.BC_.EB.B0.8F_.EC.84.A4.EC.B9.98)

### 선택: 펌웨어 설치

### Firmware

#### Suggested: Linux Firmware

On many systems, non-FOSS firmware is required for certain hardware to function. The [sys-kernel/linux-firmware](https://packages.gentoo.org/packages/sys-kernel/linux-firmware) package contains firmware for many, but not all, devices.

**요령**

Most wireless cards and GPUs require firmware to function.

`root #` `emerge --ask sys-kernel/linux-firmware`

**참고**

Installing certain firmware packages often requires accepting the associated firmware licenses. If necessary, visit the [license handling section](/wiki/Handbook:HPPA/Working/Portage/ko#Licenses "Handbook:HPPA/Working/Portage/ko") of the Handbook for help on accepting licenses.

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

Kernel installation tasks such as copying the kernel image to /boot or the [EFI System Partition](/wiki/EFI_System_Partition "EFI System Partition"), generating an [initramfs](/wiki/Initramfs "Initramfs") and/or [Unified Kernel Image](/wiki/Unified_Kernel_Image "Unified Kernel Image"), updating bootloader configuration, can be automated with [installkernel](/wiki/Installkernel "Installkernel"). Users may wish to configure and install [sys-kernel/installkernel](https://packages.gentoo.org/packages/sys-kernel/installkernel) before proceeding. See the [Kernel installation section below](/wiki/Handbook:HPPA/Installation/Kernel#Kernel_installation.2Fko "Handbook:HPPA/Installation/Kernel") for more more information.

## 기본: 직접 설정

## 소스 코드 설치

When manually compiling the kernel for hppa-based systems, Gentoo recommends the [sys-kernel/gentoo-sources](https://packages.gentoo.org/packages/sys-kernel/gentoo-sources) package.

Choose an appropriate kernel source and install it using emerge:

`root #` `emerge --ask sys-kernel/gentoo-sources`

/usr/src를 들여다보면 설치한 커널 소스를 가리키는 linux 심볼릭 링크를 볼 수 있습니다:

It is conventional for a /usr/src/linux symlink to be maintained, such that it refers to whichever sources correspond with the currently running kernel. However, this symbolic link will not be created by default. An easy way to create the symbolic link is to utilize eselect's kernel module.

For further information regarding the purpose of the symlink, and how to manage it, please refer to [Kernel/Upgrade](/wiki/Kernel/Upgrade/ko "Kernel/Upgrade/ko").

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

**중요**

Only one selection is required in the following subsection, if unsure of which to use then go with the first listed for now. It's always possible to switch at a later date if required.

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

Now it's time to decide if modules signing is required in the steps listed in [here](/wiki/Handbook:HPPA/Installation/Kernel#Optional:_Signed_kernel_modules_2.2Fko "Handbook:HPPA/Installation/Kernel")

If not, proceed to building described [here](/wiki/Handbook:HPPA/Installation/Kernel#Compiling_and_installing.2Fko "Handbook:HPPA/Installation/Kernel")

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

**중요**

64비트 커널을 컴파일 하려면, 우선 kgcc64를 이머지하십시오. 허나, 64비트 커널은 권장하지 않습니다. 시스템에 4GB 이상의 RAM을 장착했거나 A500 같은 서버에서 필요한 경우에만 64비트 커널을 실행하십시오.

이제 커널을 설정했고 컴파일 하고 설치할 차례입니다. 설정을 빠져나간 후 컴파일 과정을 시작하십시오:

`root #` `make && make modules_install`

If building a 64-bit kernel, do this instead (it's necessary even for native builds, see [here](https://www.spinics.net/lists/kernel/msg3987729.html)):

`root #` `CROSS_COMPILE=hppa64-unknown-linux-gnu- make && CROSS_COMPILE=hppa64-unknown-linux-gnu- make modules_install`

**참고**

`make -jX` 명령을 사용하고 X에 실행 가능토록 허용할 빌드 프로세스 갯수를 넣어 병렬 빌드를 활성화 할 수 있습니다. 이는 앞서 언급한 /etc/portage/make.conf의 `MAKEOPTS` 변수와 비슷합니다.

커널 컴파일이 끝나면, 커널 이미지를 /boot/에 복사하십시오. 어떤 커널을 선택하고 적당한 이름을 사용하든지간에 부트로더 설정에 필요하므로 기억해두십시오. 설치한 커널의 kernel-6.19.1-gentoo 이름과 버전을 바꾸어야함을 기억하십시오.

`root #` `cp vmlinux /boot/kernel-6.19.1-gentoo`

[시스템 설정](/wiki/Handbook:HPPA/Installation/System/ko "Handbook:HPPA/Installation/System/ko") 으로 설치 과정을 계속 진행하십시오.

[← 젠투 베이스 시스템 설치](/wiki/Handbook:HPPA/Installation/Base/ko "이전 내용") [처음](/wiki/Handbook:HPPA/ko "Handbook:HPPA/ko") [시스템 설정 →](/wiki/Handbook:HPPA/Installation/System/ko "다음 내용")