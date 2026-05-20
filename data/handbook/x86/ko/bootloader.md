# Bootloader

Other languages:

- [Deutsch](/wiki/Handbook:X86/Installation/Bootloader/de "Handbuch:X86/Installation/Bootloader (100% translated)")
- [English](/wiki/Handbook:X86/Installation/Bootloader "Handbook:X86/Installation/Bootloader (100% translated)")
- [español](/wiki/Handbook:X86/Installation/Bootloader/es "Manual de Gentoo: X86/Instalación/Arranque (100% translated)")
- [français](/wiki/Handbook:X86/Installation/Bootloader/fr "Manuel:X86/Installation/Système d'amorçage (100% translated)")
- [italiano](/wiki/Handbook:X86/Installation/Bootloader/it "Handbook:X86/Installation/Bootloader (100% translated)")
- [magyar](/wiki/Handbook:X86/Installation/Bootloader/hu "Handbook:X86/Installation/Bootloader/hu (100% translated)")
- [polski](/wiki/Handbook:X86/Installation/Bootloader/pl "Handbook:X86/Installation/Bootloader (100% translated)")
- [português do Brasil](/wiki/Handbook:X86/Installation/Bootloader/pt-br "Manual:X86/Instalação/Gerenciador de Boot (100% translated)")
- [čeština](/wiki/Handbook:X86/Installation/Bootloader/cs "Handbook:X86/Installation/Bootloader/cs (50% translated)")
- [русский](/wiki/Handbook:X86/Installation/Bootloader/ru "Handbook:X86/Installation/Bootloader (100% translated)")
- [தமிழ்](/wiki/Handbook:X86/Installation/Bootloader/ta "கையேடு:X86/நிறுவல்/துவக்கஏற்றி (100% translated)")
- [中文（中国大陆）‎](/wiki/Handbook:X86/Installation/Bootloader/zh-cn "手册：X86/安装/配置系统引导程序Bootloader (100% translated)")
- [日本語](/wiki/Handbook:X86/Installation/Bootloader/ja "ハンドブック:X86/インストール/ブートローダの設定 (100% translated)")
- 한국어

[X86 핸드북](/wiki/Handbook:X86/ko "Handbook:X86/ko")[설치](/wiki/Handbook:X86/Full/Installation/ko "Handbook:X86/Full/Installation/ko")[설치 정보](/wiki/Handbook:X86/Installation/About/ko "Handbook:X86/Installation/About/ko")[매체 선택](/wiki/Handbook:X86/Installation/Media/ko "Handbook:X86/Installation/Media/ko")[네트워크 설정](/wiki/Handbook:X86/Installation/Networking/ko "Handbook:X86/Installation/Networking/ko")[디스크 준비](/wiki/Handbook:X86/Installation/Disks/ko "Handbook:X86/Installation/Disks/ko")[스테이지 3 설치](/wiki/Handbook:X86/Installation/Stage/ko "Handbook:X86/Installation/Stage/ko")[베이스 시스템 설치](/wiki/Handbook:X86/Installation/Base/ko "Handbook:X86/Installation/Base/ko")[커널 설정](/wiki/Handbook:X86/Installation/Kernel/ko "Handbook:X86/Installation/Kernel/ko")[시스템 설정](/wiki/Handbook:X86/Installation/System/ko "Handbook:X86/Installation/System/ko")[도구 설치](/wiki/Handbook:X86/Installation/Tools/ko "Handbook:X86/Installation/Tools/ko")부트로더 설정[마무리](/wiki/Handbook:X86/Installation/Finalizing/ko "Handbook:X86/Installation/Finalizing/ko")[젠투 활용](/wiki/Handbook:X86/Full/Working/ko "Handbook:X86/Full/Working/ko")[포티지 소개](/wiki/Handbook:X86/Working/Portage/ko "Handbook:X86/Working/Portage/ko")[USE 플래그](/wiki/Handbook:X86/Working/USE/ko "Handbook:X86/Working/USE/ko")[포티지 기능](/wiki/Handbook:X86/Working/Features/ko "Handbook:X86/Working/Features/ko")[초기화 스크립트 시스템](/wiki/Handbook:X86/Working/Initscripts/ko "Handbook:X86/Working/Initscripts/ko")[환경 변수](/wiki/Handbook:X86/Working/EnvVar/ko "Handbook:X86/Working/EnvVar/ko")[포티지 활용](/wiki/Handbook:X86/Full/Portage/ko "Handbook:X86/Full/Portage/ko")[파일 및 디렉터리](/wiki/Handbook:X86/Portage/Files/ko "Handbook:X86/Portage/Files/ko")[변수](/wiki/Handbook:X86/Portage/Variables/ko "Handbook:X86/Portage/Variables/ko")[소프트웨어 브랜치 함께 사용하기](/wiki/Handbook:X86/Portage/Branches/ko "Handbook:X86/Portage/Branches/ko")[추가 도구](/wiki/Handbook:X86/Portage/Tools/ko "Handbook:X86/Portage/Tools/ko")[꾸러미 저장소 개별 설정](/wiki/Handbook:X86/Portage/CustomTree/ko "Handbook:X86/Portage/CustomTree/ko")[고급 기능](/wiki/Handbook:X86/Portage/Advanced/ko "Handbook:X86/Portage/Advanced/ko")[네트워크 설정](/wiki/Handbook:X86/Full/Networking/ko "Handbook:X86/Full/Networking/ko")[시작하기](/wiki/Handbook:X86/Networking/Introduction/ko "Handbook:X86/Networking/Introduction/ko")[고급 설정](/wiki/Handbook:X86/Networking/Advanced/ko "Handbook:X86/Networking/Advanced/ko")[모듈러 네트워크](/wiki/Handbook:X86/Networking/Modular/ko "Handbook:X86/Networking/Modular/ko")[무선 네트워크](/wiki/Handbook:X86/Networking/Wireless/ko "Handbook:X86/Networking/Wireless/ko")[기능 추가](/wiki/Handbook:X86/Networking/Extending/ko "Handbook:X86/Networking/Extending/ko")[동적 관리](/wiki/Handbook:X86/Networking/Dynamic/ko "Handbook:X86/Networking/Dynamic/ko")

비록 32비트 CPU에 설치한다지만 대부분의 모든 **x86** 마더보드(2006-2007년경부터 현재까지)는 _64비트 UEFI 펌웨어_ 를 지원하는 모델로 출시합니다. 일부 사용자는 설정 이름과 파일에서 "64"라는 숫자를 눈여겨보게 될지도 모릅니다. 이는 거의 모든 경우에 _해당되리라_ 봅니다.

당연스럽게도 UEFI 펌웨어가 64비트라지만 매우 드문 예외가 있는데, 극히 일부 초기 [애플 맥](http://www.everymac.com/systems/apple/mac_pro/specs/mac-pro-quad-2.66-specs.html) 모델과 인텔 아톰 기반 [델 태블릿 PC](http://www.dell.com/en-us/shop/productdetails/dell-venue-8-pro) 에서는 32비트 UEFI 펌웨어를 지원합니다. 대부분의 독자는 _절대적으로_ 32비트 UEFI 펌웨어를 볼 일이 없습니다. 이 때문에 32비트 UEFI 펌웨어는 **x86** 핸드북에서 다루지 않습니다.

## Contents

- [1부트로더 선택](#.EB.B6.80.ED.8A.B8.EB.A1.9C.EB.8D.94_.EC.84.A0.ED.83.9D)
- [2기본: GRUB2](#.EA.B8.B0.EB.B3.B8:_GRUB2)
  - [2.1Emerge](#Emerge)
  - [2.2설치](#.EC.84.A4.EC.B9.98)
    - [2.2.1Optional: Secure Boot](#Optional:_Secure_Boot)
    - [2.2.2Debugging GRUB](#Debugging_GRUB)
  - [2.3설정](#.EC.84.A4.EC.A0.95)
- [3Alternative 1: systemd-boot](#Alternative_1:_systemd-boot)
  - [3.1Emerge](#Emerge_2)
  - [3.2Installation](#Installation)
  - [3.3Optional: Secure Boot](#Optional:_Secure_Boot_2)
- [4대안 2: efibootmgr](#.EB.8C.80.EC.95.88_2:_efibootmgr)
  - [4.1Unified Kernel Image](#Unified_Kernel_Image)
- [5Other Alternatives](#Other_Alternatives)
- [6시스템 다시 부팅](#.EC.8B.9C.EC.8A.A4.ED.85.9C_.EB.8B.A4.EC.8B.9C_.EB.B6.80.ED.8C.85)

## 부트로더 선택

리눅스 커널을 설정하고, 시스템 도구를 설치하고, 설정 파일을 편집하고 나면, 리눅스 설치의 일부로서 마지막으로 중요한 부분, 부트로더를 설치할 차례입니다.

부트로더는 부팅 과정에서 리눅스 커널의 실행을 담당합니다. 부트로더가 없으면 시스템에서는 전원 단추를 누른 후 그 다음 과정을 어떻게 처리해야 할지 모릅니다.

**x86** 에서는 BIOS 기반 시스템에서 [GRUB2](#.EA.B8.B0.EB.B3.B8:_GRUB2_.EC.82.AC.EC.9A.A9) 또는 [LILO](#.EB.8C.80.EC.95.88:_LILO_.EC.82.AC.EC.9A.A9) 를 설정하는 방법, UEFI 시스템에서 [GRUB2](#.EA.B8.B0.EB.B3.B8:_GRUB2_.EC.82.AC.EC.9A.A9) 또는 [efibootmgr](#.EB.8C.80.EC.95.88:_efibootmgr_.EC.82.AC.EC.9A.A9) 을 사용하는 방법을 문서에 남겨두었습니다.

여기서는 부트로더 꾸러미를 _이머징_ 하고 시스템 디스크에 부트로더를 _설치_ 하는 방법을 설명합니다. _이머징_ 은 시스템에서 활용할 수 있는 프로그램 꾸러미를 [Portage](/wiki/Portage "Portage") 에 만들려고 요청할 때 활용합니다. _설치_ 는 다음에 전원을 다시 켰을 때 부트로더를 _활성화한 후 동작 대기_ 상태로 두도록 시스템 디스크 드라이브의 적당한 부분에 파일을 복사하거나 물리적으로 수정함을 의미합니다.

## 기본: GRUB2

기본적으로, 젠투 시스템에서는 기존 GRUB을 그대로 이어가는 GRUB2([sys-boot/grub](https://packages.gentoo.org/packages/sys-boot/grub))로 의존 대세가 바뀌었습니다. 설정을 추가로 하지 않아도 이전 BIOS("pc") 시스템을 기꺼이 지원합니다. 빌드 이전에 필요한 일부 설정으로 여러 추가 플랫폼을 지원할 수 있습니다. 자세히 알아보시려면 [GRUB](/wiki/GRUB/ko "GRUB/ko") 게시글의 [선행 과정 부분](/wiki/GRUB/ko#.EC.84.A0.ED.96.89_.EA.B3.BC.EC.A0.95 "GRUB/ko") 을 참고하십시오.

### Emerge

MBR 분할 영역 테이블만 지원하는 이전 BIOS 시스템을 활용한다면 GRUB을 이머징할 때 추가로 설정할 필요가 없습니다:

`root #` `emerge --ask --verbose sys-boot/grub:2`

- UEFI 사용자 참고사항: 위 명령을 실행하면 이머징을 시작하기 전 GRUB\_PLATFORMS 활성 변수 값을 나타냅니다. UEFI 기능이 동작하는 시스템을 활용한다면, 해당 사용자는 `GRUB_PLATFORMS="efi-64"`(기본적인 고려 사항)값을 넣었는지 확인해야합니다. 설치 과정에 이 값이 나타나지 않은 경우라면 GRUB2를 이머징하기 _전_/etc/portage/make.conf 파일에 `GRUB_PLATFORMS="efi-64"` 값을 추가하여, GRUB2에 EFI 기능을 넣어 빌드할 수 있게 해야합니다.

`root #` `echo 'GRUB_PLATFORMS="efi-64"' >> /etc/portage/make.conf`

`root #` `emerge --ask sys-boot/grub:2`

(위에서와 같이) make.conf에 `GRUB_PLATFORMS="efi-64"` 를 활성화하지 않고 GRUB2를 이머지했다 하더라도 나중에 해당 줄을 추가할 수 있고, emerge에 `--update --newuse` 옵션을 전달하여 [전체 꾸러미 모음](/wiki/World_set_(Portage) "World set (Portage)") en의 의존성을 다시 처리할 수 있습니다.

`root #` `emerge --ask --update --newuse --verbose sys-boot/grub:2`

이제 GRUB2 프로그램을 이머징했지만, 설치한 상태는 아닙니다.

### 설치

다음, /boot/grub/에 필요한 GRUB2 파일을 grub-install 명령으로 설치할 차례입니다. 첫번째 디스크(부팅할 디스크)를 /dev/sda라고 가정하고, 다음 명령중 하나로 GRUB2 파일을 설치합니다:

- BIOS를 사용하면:

`root #` `grub-install /dev/sda`

For DOS/Legacy BIOS systems:

`root #` `grub-install /dev/sda`

- UEFI를 사용하면:

**중요**

grub-install을 실행하기 전 EFI 시스템 공간을 마운트 했는지 확인하십시오. grub-install 프로그램이 잘못된 디렉터리를 사용한 "어떤" 설정 내지는 표시 **없이** 잘못된 디렉터리의 GRUB EFI 파일(grubx64.efi)로 설치할 수 있습니다.

For UEFI systems:

`root #` `grub-install --target=x86_64-efi --efi-directory=/boot`

Upon successful installation, the output should match the output of the previous command. If the output does not match exactly, then proceed to [Debugging GRUB](/wiki/Handbook:X86/Blocks/Bootloader/ko#Debugging_GRUB "Handbook:X86/Blocks/Bootloader/ko"), otherwise jump to the [Configure step](/wiki/Handbook:X86/Blocks/Bootloader/ko#GRUB_Configure "Handbook:X86/Blocks/Bootloader/ko").

##### Optional: Secure Boot

To successfully boot with secure boot enabled the signing certificate must either be accepted by the [UEFI](/wiki/UEFI "UEFI") firmware, or [shim](/wiki/Shim "Shim") must be used as a pre-loader. Shim is pre-signed with the third-party Microsoft Certificate, accepted by default by most UEFI motherboards.

How to configure the UEFI firmware to accept custom keys depends on the firmware vendor, which is beyond the scope of the handbook. Below is shown how to setup shim instead. Here it is assumed that the user has already followed the instructions in the previous sections to generate a signing key and to configure portage to use it. If this is not the case please return first to the [Kernel installation](/wiki/Handbook:X86/Installation/Kernel/ko "Handbook:X86/Installation/Kernel/ko") section.

The package [sys-boot/grub](https://packages.gentoo.org/packages/sys-boot/grub) installs a prebuilt and signed stand-alone EFI executable if the [secureboot](https://packages.gentoo.org/useflags/secureboot) [USE flag/ko (page does not exist)](/index.php?title=USE_flag/ko&action=edit&redlink=1 "USE flag/ko (page does not exist)") USE flag is enabled. Install the required packages and copy the stand-alone grub, Shim, and the MokManager to the same directory on the EFI System Partition. For example:

`root #` `emerge sys-boot/grub sys-boot/shim sys-boot/mokutil sys-boot/efibootmgr`

`root #` `cp /usr/share/shim/BOOTX64.EFI /efi/EFI/Gentoo/shimx64.efi
`

`root #` `cp /usr/share/shim/mmx64.efi /efi/EFI/Gentoo/mmx64.efi
`

`root #` `cp /usr/lib/grub/grub-x86_64.efi.signed /efi/EFI/Gentoo/grubx64.efi
`

Next register the signing key in shims MOKlist, this requires keys in the DER format, whereas sbsign and the kernel build system expect keys in the PEM format. In the previous sections of the handbook an example was shown to generate such a signing PEM key, this key must now be converted to the DER format:

`root #` `openssl x509 -in /path/to/kernel_key.pem -inform PEM -out /path/to/kernel_key.der -outform DER`

**참고**

The path used here must be the path to the pem file containing the certificate belonging to the generated key. In this example both key and certificate are in the same pem file.

Then the converted certificate can be imported into Shims MOKlist, this command will ask to set some password for the import request:

`root #` `mokutil --import /path/to/kernel_key.der`

**참고**

When the currently booted kernel already trusts the certificate being imported, the message "Already in kernel trusted keyring." will be returned here. If this happens, re-run the above command with the argument --ignore-keyring added.

Next, register Shim with the UEFI firmware. In the following command, `boot-disk` and `boot-partition-id` must be replaced with the disk and partition identifier of the EFI system partition:

`root #` `efibootmgr --create --disk /dev/boot-disk --part boot-partition-id --loader '\EFI\Gentoo\shimx64.efi' --label 'GRUB via Shim' --unicode`

Note that this prebuilt and signed stand-alone version of grub reads the grub.cfg from a different location then usual. Instead of the default /boot/grub/grub.cfg the config file should be in the same directory that the grub EFI executable is in, e.g. /efi/EFI/Gentoo/grub.cfg. When [sys-kernel/installkernel](https://packages.gentoo.org/packages/sys-kernel/installkernel) is used to install the kernel and update the grub configuration then the GRUB\_CFG environment variable may be used to override the usual location of the grub config file.

For example:

`root #` `grub-mkconfig -o /efi/EFI/Gentoo/grub.cfg`

Or, via [installkernel](/wiki/Installkernel "Installkernel"):

파일 **`/etc/env.d/99grub`**

```
GRUB_CFG=/efi/EFI/Gentoo/grub.cfg

```

`root #` `env-update`

**참고**

The import process will not be completed until the system is rebooted. After completing all steps in the handbook, restart the system and Shim will load, it will find the import request registered by mokutil. The MokManager application will start and ask for the password that was set when creating the import request. Follow the on-screen instructions to complete the import of the certificate, then reboot the system into the UEFI menu and enable the Secure Boot setting.

##### Debugging GRUB

When debugging GRUB, there are a couple of quick fixes that may result in a bootable installation without having to reboot to a new live image environment.

In the event that "EFI variables are not supported on this system" is displayed somewhere in the output, it is likely the live image was not booted in EFI mode and is presently in Legacy BIOS boot mode. The solution is to try the [removable GRUB step](/wiki/Handbook:X86/Blocks/Bootloader/ko#GRUB_Install_EFI_systems_removable "Handbook:X86/Blocks/Bootloader/ko") mentioned below. This will overwrite the executable EFI file located at /EFI/BOOT/BOOTX64.EFI. Upon rebooting in EFI mode, the motherboard firmware may execute this default boot entry and execute GRUB.

**중요**

grub-install 에서 `Could not prepare Boot variable: Read-only file system` 와 같은 오류가 나타났을 경우, 과정을 제대로 넘어가기 위해 efivars 특수 마운트 지점을 다시 마운트 해야합니다:

`root #` `mount -o remount,rw /sys/firmware/efi/efivars`

`root #` `mount -o remount,rw,nosuid,nodev,noexec --types efivarfs efivarfs /sys/firmware/efi/efivars`

This is caused by certain non-official Gentoo environments not mounting the special EFI filesystem by default. If the previous command does not run, then reboot using an official Gentoo live image environment in EFI mode.

일부 마더보드 제조사에서는 .EFI 파일을 저장할 EFI 시스템 파티션(ESP)의 디렉터리 위치 경로를 /efi/boot/ 디렉터리만 지원하는 것 같습니다. GRUB 설치 관리자에서는 `--removable` 옵션을 주어 자동으로 처리할 수 있습니다. 다음 명령을 실행하기 전에 ESP를 마운트했는지 확인하십시오. (앞서 말한대로) /boot에 ESP를 마운트했다고 생각한다면, 다음 명령을 실행하십시오:

`root #` `grub-install --target=x86_64-efi --efi-directory=/boot --removable`

이 명령은 UEFI 명세로 정의한 기본 디렉터리를 만들고 grubx64.efi 파일을 동일한 명세로 정의한 '기본' EFI 파일 위치에 복사합니다.

### 설정

다음, /etc/default/grub 파일과 /etc/grub.d 스크립트를 기반으로 정의한 사용자 설정을 기반으로 하여 GRUB2 설정을 만드십시오. GRUB2에서 어떤 커널로 부팅할지 (/boot/에 존재하는 가장 최신의 버전), 루트 파일 시스템이 어디인지 자동으로 찾기 때문에 사용자가 직접 설정할 필요가 없습니다. /etc/default/grub 의 GRUB\_CMDLINE\_LINUX 변수에 커널 매개 변수를 넣을 수도 있습니다.

최종 GRUB2 설정을 만들려면 grub-mkconfig 명령을 실행하십시오:

`root #` `grub-mkconfig -o /boot/grub/grub.cfg`

```
Generating grub.cfg ...
Found linux image: /boot/vmlinuz-6.18.8-gentoo
Found initrd image: /boot/initramfs-genkernel-x86-6.18.8-gentoo
done

```

명령 출력은 최소한 시스템 부팅에 필요한 하나의 리눅스 이미지를 찾았다는 알림이 나와야 합니다. initramfs를 사용했거나, 커널 빌드시 genkernel를 사용했다면, 마찬가지로 올바른 initrd 이미지를 찾아야 합니다. 이런 경우가 아니라면 /boot/로 이동해서 ls 명령으로 내용을 확인해보십시오. 여전히 파일이 빠졌다면, 커널 설정으로 되돌아가서 설치 절차를 따르십시오.

**요령**

os-prober 유틸리티로 붙어있는 드라이브에서 다른 우영체제를 찾아 붙일 수 있습니다. 윈도우 7, 8.1, 10, 그리고 다른 리눅스 배포판을 찾을 수 있습니다. 듀얼 부팅 시스템을 원한다면 [sys-boot/os-prober](https://packages.gentoo.org/packages/sys-boot/os-prober) 꾸러미를 설치하고 grub-mkconfig 명령을 (위에서 보신 것처럼) 다시 실행하십시오. 운영체제를 찾는 과정에 문제가 있다면, 젠투 커뮤니티에 지원을 요청하기 _전_ [GRUB](/wiki/GRUB/ko "GRUB/ko") 게시글을 다시 확인하십시오.

## Alternative 1: systemd-boot

Another option is [systemd-boot](/wiki/Systemd/systemd-boot "Systemd/systemd-boot"), which works on both OpenRC and systemd machines. It is a thin chainloader and works well with secure boot.

### Emerge

To install systemd-boot, enable the [boot](https://packages.gentoo.org/useflags/boot) [USE flag/ko (page does not exist)](/index.php?title=USE_flag/ko&action=edit&redlink=1 "USE flag/ko (page does not exist)") USE flag and re-install [sys-apps/systemd](https://packages.gentoo.org/packages/sys-apps/systemd) (for systemd systems) or [sys-apps/systemd-utils](https://packages.gentoo.org/packages/sys-apps/systemd-utils) (for OpenRC systems):

파일 **`/etc/portage/package.use/systemd-boot`**

```
sys-apps/systemd boot
sys-apps/systemd-utils boot

```

`root #` `emerge --ask sys-apps/systemd`

Or

`root #` `emerge --ask sys-apps/systemd-utils`

### Installation

Now, install the systemd-boot loader to the [EFI System Partition](/wiki/EFI_System_Partition "EFI System Partition"):

`root #` `bootctl install`

**중요**

Make sure the EFI system partition has been mounted _before_ running bootctl install.

When using this bootloader, before rebooting, verify that a new bootable entry exists using:

`root #` `bootctl list`

**경고**

The kernel command line for new systemd-boot entries is read from /etc/kernel/cmdline or /usr/lib/kernel/cmdline. If neither file is present, then the kernel command line of the currently booted kernel is re-used (/proc/cmdline). On new installs it might therefore happen that the kernel command line of the live CD is accidentally used to boot the new kernel. The kernel command line for registered entries can be checked with:

`root #` `bootctl list`

If this does not show the desired kernel command line then create /etc/kernel/cmdline containing the correct kernel command line and re-install the kernel.

If no new entry exists, ensure the [sys-kernel/installkernel](https://packages.gentoo.org/packages/sys-kernel/installkernel) package has been installed with the [systemd](https://packages.gentoo.org/useflags/systemd) [USE flag/ko (page does not exist)](/index.php?title=USE_flag/ko&action=edit&redlink=1 "USE flag/ko (page does not exist)") and [systemd-boot](https://packages.gentoo.org/useflags/systemd-boot) [USE flag/ko (page does not exist)](/index.php?title=USE_flag/ko&action=edit&redlink=1 "USE flag/ko (page does not exist)") USE flags enabled, and re-run the kernel installation.

For the distribution kernels:

`root #` `emerge --ask --config sys-kernel/gentoo-kernel`

For a manually configured and compiled kernel:

`root #` `make install`

**중요**

When installing kernels for systemd-boot, no root= kernel command line argument is added by default. On systemd systems that are using an initramfs users may rely instead on [systemd-gpt-auto-generator](/wiki/Systemd#Automatic_mounting_of_partitions_at_boot.2Fko "Systemd") to automatically find the root partition at boot. Otherwise users should manually specify the location of the root partition by setting root= in /etc/kernel/cmdline as well as any other kernel command line arguments that should be used. And then reinstalling the kernel as described above.

### Optional: Secure Boot

When the [secureboot](https://packages.gentoo.org/useflags/secureboot) [USE flag/ko (page does not exist)](/index.php?title=USE_flag/ko&action=edit&redlink=1 "USE flag/ko (page does not exist)") USE flag is enabled, the systemd-boot EFI executable will be signed by portage automatically. Furthermore, bootctl install will automatically install the signed version.

To successfully boot with secure boot enabled the used certificate must either be accepted by the [UEFI](/wiki/UEFI "UEFI") firmware, or [shim](/wiki/Shim "Shim") must be used as a pre-loader. Shim is pre-signed with the third-party Microsoft Certificate, accepted by default by most UEFI motherboards.

How to configure the UEFI firmware to accept custom keys depends on the firmware vendor, which is beyond the scope of the handbook. Below is shown how to setup shim instead. Here it is assumed that the user has already followed the instructions in the previous sections to generate a signing key and to configure portage to use it. If this is not the case please return first to the [Kernel installation](/wiki/Handbook:X86/Installation/Kernel/ko "Handbook:X86/Installation/Kernel/ko") section.

`root #` `emerge --ask sys-boot/shim sys-boot/mokutil sys-boot/efibootmgr`

`root #` `bootctl install --no-variables`

`root #` `cp /usr/share/shim/BOOTX64.EFI /efi/EFI/systemd/shimx64.efi`

`root #` `cp /usr/share/shim/mmx64.efi /efi/EFI/systemd/mmx64.efi`

Shims MOKlist requires keys in the DER format, whereas sbsign and the kernel build system expect keys in the PEM format. In the previous sections of the handbook an example was shown to generate such a signing PEM key, this key must now be converted to the DER format:

`root #` `openssl x509 -in /path/to/kernel_key.pem -inform PEM -out /path/to/kernel_key.der -outform DER`

**참고**

The path used here must be the path to the pem file containing the certificate belonging to the generated key. In this example both key and certificate are in the same pem file.

Then the converted certificate can be imported into Shims MOKlist:

`root #` `mokutil --import /path/to/kernel_key.der`

**참고**

When the currently booted kernel already trusts the certificate being imported, the message "Already in kernel trusted keyring." will be returned here. If this happens, re-run the above command with the argument --ignore-keyring added.

And finally we register Shim with the UEFI firmware. In the following command, `boot-disk` and `boot-partition-id` must be replaced with the disk and partition identifier of the EFI system partition:

`root #` `efibootmgr --create --disk /dev/boot-disk --part boot-partition-id --loader '\EFI\systemd\shimx64.efi'  --label 'Systemd-boot via Shim' --unicode '\EFI\systemd\systemd-bootx64.efi'`

**참고**

The import process will not be completed until the system is rebooted. After completing all steps in the handbook, restart the system and Shim will load, it will find the import request registered by mokutil. The MokManager application will start and ask for the password that was set when creating the import request. Follow the on-screen instructions to complete the import of the certificate, then reboot the system into the UEFI menu and enable the Secure Boot setting.

## 대안 2: efibootmgr

UEFI 기반 시스템에서는, 시스템의 UEFI 펌웨어(기본 부트 로더)에서 UEFI 부팅 항목을 직접 다룰 수 있습니다. 일부 시스템에서는 부팅할 때 GRUB2와 같은 추가(또는 차선) 부트로더가 필요하지 않습니다. 이 말은 결국, GRUB2 같은 EFI 기반 부트로더의 존재 이유가 바로 부팅 과정에서 UEFI 시스템 기능을 "확장"하는데 있습니다. efibootmgr를 사용하는 것이야말로(융통성이 없긴 하지만) 시스템 부팅 관점 접근 상 간소화를 원하는 이들에게 안성맞춤입니다. GRUB2(상부 참고)를 사용하는게 주로 사용자에게 더 쉬운 방법인데 UEFI 시스템을 부팅할 때 유연한 접근 방식을 취하기 때문입니다.

System administrators who desire to take a minimalist, although more rigid, approach to booting the system can avoid secondary bootloaders and boot the Linux kernel as an [EFI stub](/wiki/EFI_stub "EFI stub").

[sys-boot/efibootmgr](https://packages.gentoo.org/packages/sys-boot/efibootmgr) 프로그램은 부트로더가 아님을 기억해두십시오. efibootmgr은 UEFI 펌웨어와 소통하고 UEFI 펌웨어 설정을 업데이트할 때 사용하는 도구이기 때문에 이전에 설치한 리눅스 커널을 (필요한 경우)추가 옵션을 통해 부팅할 수 있으며, 또는 다중 부팅 항목을 넣을 수 있습니다. 이러한 동작간 작용은 EFI 변수로 처리할 수 있습니다(따라서 이를 처리하기 전에 커널의 EFI 변수 지원이 필요합니다).

계속하기 _전_, [EFI stub kernel](/wiki/EFI_stub_kernel "EFI stub kernel") 게시글을 우선 읽어두십시오. 시스템의 UEFI 펌웨어로 바로 부팅할 수 있게 커널의 몇가지 지정 옵션을 활성화해야합니다. 커널을 다시 컴파일 해야 할 수도 있습니다. [efibootmgr](/wiki/Efibootmgr/ko "Efibootmgr/ko") 게시물을 살펴보시는 것도 좋습니다.

It is also a good idea to take a look at the [efibootmgr](/wiki/Efibootmgr "Efibootmgr") article for additional information.

**참고**

다시 반복하지만, efibootmgr은 UEFI 시스템에서 부팅하는데 필수 요건이 _아닙니다_. 리눅스 커널 자체에서 즉시 부팅할 수 있으며, 추가 커널 명령행 옵션은 리눅스 커널 자체의 일부로 내장할 수 있습니다(사용자가 명령행에서 처럼 부팅 매개 변수를 지원하도록 호출할 수 있는 커널 설정 옵션이 있습니다). 심지어 initramfs도 커널 '그 자체'로 빌드할 수 있습니다.

이런 방식으로 진행하기로 마음먹었다면 프로그램을 설치해야합니다:

파일 **`/etc/portage/package.accept_keywords/installkernel`**

```
sys-kernel/installkernel
sys-boot/uefi-mkconfig
app-emulation/virt-firmware

```

파일 **`/etc/portage/package.use/installkernel`**

```
sys-kernel/installkernel efistub

```

Then reinstall [installkernel](/wiki/Installkernel "Installkernel"), create the /efi directory and reinstall the kernel:

`root #` `emerge --ask sys-kernel/installkernel`

`root #` `mkdir -p /efi`

For distribution kernels:

`root #` `emerge --ask --config sys-kernel/gentoo-kernel{,-bin}`

For manually managed kernels:

`root #` `make install`

그 다음 /boot/efi/boot/ 디렉터리를 만들고, 커널을 bootx64.efi 위치에 복사하십시오:

`root #` `mkdir -p /boot/efi/boot
`

`root #` `cp /boot/vmlinuz-* /boot/efi/boot/bootx64.efi
`

Install the efibootmgr package:

`root #` `emerge --ask sys-boot/efibootmgr`

다음, 새로 컴파일한 EFI 커널로 부팅하는 "Gentoo" 항목을 UEFI 펌웨어에 적어주십시오:

`root #` `efibootmgr --create --disk /dev/sda --part 2 --label "Gentoo" --loader "\efi\boot\bootx64.efi"`

**참고**

UEFI 정의를 활용한다면 디렉터리 구분자로서 `\` 사용은 필수입니다.

초기 램 파일 시스템(initramfs)를 사용한다면, 적당한 부팅 옵션을 추가하십시오:

`root #` `efibootmgr -c -d /dev/sda -p 2 -L "Gentoo" -l "\efi\boot\bootx64.efi" initrd='\initramfs-genkernel-x86-6.18.8-gentoo'`

**요령**

Additional kernel command line options may be parsed by the firmware to the kernel by specifying them along with the initrd=... option as shown above.

설정을 다 바꿨다면, 시스템을 다시 부팅할 때, 부팅 항목에 "Gentoo"가 뜹니다.

### Unified Kernel Image

If [installkernel](/wiki/Installkernel "Installkernel") was configured to build and install unified kernel images. The unified kernel image should already be installed to the EFI/Linux directory on the EFI system partition, if this is not the case ensure the directory exists and then run the kernel installation again as described earlier in the handbook.

To add a direct boot entry for the installed unified kernel image:

`root #` `efibootmgr --create --disk /dev/sda --part 1 --label "gentoo" --loader "\EFI\Linux\gentoo-x.y.z.efi"`

## Other Alternatives

For other options that are not covered in the Handbook, see the full list of available [bootloaders](/wiki/Bootloader "Bootloader").

## 시스템 다시 부팅

chroot로 진입한 환경을 빠져나가고 모든 파티션의 마운트를 해제하십시오. 그 다음 대미를 장식할 마법의 명령을 입력하여, 실제로 시험해보십시오:
reboot.

`root #` `exit`

`cdimage ~#` `cd
`

`cdimage ~#` `umount -l /mnt/gentoo/dev{/shm,/pts,}
`

`cdimage ~#` `umount -R /mnt/gentoo
`

`cdimage ~#` `reboot`

물론 부팅 CD를 제거하는걸 잊지 않으면 새 젠투 시스템 대신 CD로 부팅합니다.

새로 설치한 젠투 환경으로 다시 부팅하고 나면, [젠투 설치 마무리](/wiki/Handbook:X86/Installation/Finalizing/ko "Handbook:X86/Installation/Finalizing/ko") 로 끝내십시오.

[← 시스템 도구 설치](/wiki/Handbook:X86/Installation/Tools/ko "이전 내용") [처음](/wiki/Handbook:X86/ko "Handbook:X86/ko") [마무리 →](/wiki/Handbook:X86/Installation/Finalizing/ko "다음 내용")