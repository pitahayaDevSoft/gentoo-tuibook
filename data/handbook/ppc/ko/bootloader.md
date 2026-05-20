# Bootloader

Other languages:

- [Deutsch](/wiki/Handbook:PPC/Installation/Bootloader/de "Handbuch:PPC/Installation/Bootloader (100% translated)")
- [English](/wiki/Handbook:PPC/Installation/Bootloader "Handbook:PPC/Installation/Bootloader (100% translated)")
- [español](/wiki/Handbook:PPC/Installation/Bootloader/es "Manual de Gentoo: PPC/Instalación/Arranque (100% translated)")
- [français](/wiki/Handbook:PPC/Installation/Bootloader/fr "Handbook:PPC/Installation/Bootloader/fr (100% translated)")
- [italiano](/wiki/Handbook:PPC/Installation/Bootloader/it "Handbook:PPC/Installation/Bootloader/it (100% translated)")
- [magyar](/wiki/Handbook:PPC/Installation/Bootloader/hu "Handbook:PPC/Installation/Bootloader/hu (100% translated)")
- [polski](/wiki/Handbook:PPC/Installation/Bootloader/pl "Handbook:PPC/Installation/Bootloader (100% translated)")
- [português do Brasil](/wiki/Handbook:PPC/Installation/Bootloader/pt-br "Handbook:PPC/Installation/Bootloader/pt-br (50% translated)")
- [русский](/wiki/Handbook:PPC/Installation/Bootloader/ru "Handbook:PPC/Installation/Bootloader (100% translated)")
- [தமிழ்](/wiki/Handbook:PPC/Installation/Bootloader/ta "கையேடு:PPC/நிறுவல்/துவக்கஏற்றி (100% translated)")
- [中文（中国大陆）‎](/wiki/Handbook:PPC/Installation/Bootloader/zh-cn "手册：PPC/安装/配置系统引导程序Bootloader (100% translated)")
- [日本語](/wiki/Handbook:PPC/Installation/Bootloader/ja "ハンドブック:PPC/インストール/ブートローダー (100% translated)")
- 한국어

[PPC 핸드북](/wiki/Handbook:PPC/ko "Handbook:PPC/ko")[설치](/wiki/Handbook:PPC/Full/Installation/ko "Handbook:PPC/Full/Installation/ko")[설치 정보](/wiki/Handbook:PPC/Installation/About/ko "Handbook:PPC/Installation/About/ko")[매체 선택](/wiki/Handbook:PPC/Installation/Media/ko "Handbook:PPC/Installation/Media/ko")[네트워크 설정](/wiki/Handbook:PPC/Installation/Networking/ko "Handbook:PPC/Installation/Networking/ko")[디스크 준비](/wiki/Handbook:PPC/Installation/Disks/ko "Handbook:PPC/Installation/Disks/ko")[스테이지 3 설치](/wiki/Handbook:PPC/Installation/Stage/ko "Handbook:PPC/Installation/Stage/ko")[베이스 시스템 설치](/wiki/Handbook:PPC/Installation/Base/ko "Handbook:PPC/Installation/Base/ko")[커널 설정](/wiki/Handbook:PPC/Installation/Kernel/ko "Handbook:PPC/Installation/Kernel/ko")[시스템 설정](/wiki/Handbook:PPC/Installation/System/ko "Handbook:PPC/Installation/System/ko")[도구 설치](/wiki/Handbook:PPC/Installation/Tools/ko "Handbook:PPC/Installation/Tools/ko")부트로더 설정[마무리](/wiki/Handbook:PPC/Installation/Finalizing/ko "Handbook:PPC/Installation/Finalizing/ko")[젠투 활용](/wiki/Handbook:PPC/Full/Working/ko "Handbook:PPC/Full/Working/ko")[포티지 소개](/wiki/Handbook:PPC/Working/Portage/ko "Handbook:PPC/Working/Portage/ko")[USE 플래그](/wiki/Handbook:PPC/Working/USE/ko "Handbook:PPC/Working/USE/ko")[포티지 기능](/wiki/Handbook:PPC/Working/Features/ko "Handbook:PPC/Working/Features/ko")[초기화 스크립트 시스템](/wiki/Handbook:PPC/Working/Initscripts/ko "Handbook:PPC/Working/Initscripts/ko")[환경 변수](/wiki/Handbook:PPC/Working/EnvVar/ko "Handbook:PPC/Working/EnvVar/ko")[포티지 활용](/wiki/Handbook:PPC/Full/Portage/ko "Handbook:PPC/Full/Portage/ko")[파일 및 디렉터리](/wiki/Handbook:PPC/Portage/Files/ko "Handbook:PPC/Portage/Files/ko")[변수](/wiki/Handbook:PPC/Portage/Variables/ko "Handbook:PPC/Portage/Variables/ko")[소프트웨어 브랜치 함께 사용하기](/wiki/Handbook:PPC/Portage/Branches/ko "Handbook:PPC/Portage/Branches/ko")[추가 도구](/wiki/Handbook:PPC/Portage/Tools/ko "Handbook:PPC/Portage/Tools/ko")[꾸러미 저장소 개별 설정](/wiki/Handbook:PPC/Portage/CustomTree/ko "Handbook:PPC/Portage/CustomTree/ko")[고급 기능](/wiki/Handbook:PPC/Portage/Advanced/ko "Handbook:PPC/Portage/Advanced/ko")[네트워크 설정](/wiki/Handbook:PPC/Full/Networking/ko "Handbook:PPC/Full/Networking/ko")[시작하기](/wiki/Handbook:PPC/Networking/Introduction/ko "Handbook:PPC/Networking/Introduction/ko")[고급 설정](/wiki/Handbook:PPC/Networking/Advanced/ko "Handbook:PPC/Networking/Advanced/ko")[모듈러 네트워크](/wiki/Handbook:PPC/Networking/Modular/ko "Handbook:PPC/Networking/Modular/ko")[무선 네트워크](/wiki/Handbook:PPC/Networking/Wireless/ko "Handbook:PPC/Networking/Wireless/ko")[기능 추가](/wiki/Handbook:PPC/Networking/Extending/ko "Handbook:PPC/Networking/Extending/ko")[동적 관리](/wiki/Handbook:PPC/Networking/Dynamic/ko "Handbook:PPC/Networking/Dynamic/ko")

## Contents

- [1선택](#.EC.84.A0.ED.83.9D)
- [2NewWorld Macs](#NewWorld_Macs)
  - [2.1GRUB](#GRUB)
  - [2.2Installation](#Installation)
  - [2.3Setup bootstrap partition](#Setup_bootstrap_partition)
  - [2.4Setup GRUB](#Setup_GRUB)
- [3OldWorld Macs](#OldWorld_Macs)
- [4대안: BootX 사용](#.EB.8C.80.EC.95.88:_BootX_.EC.82.AC.EC.9A.A9)
- [5Pegasos](#Pegasos)
- [6대안: BootCreator 사용](#.EB.8C.80.EC.95.88:_BootCreator_.EC.82.AC.EC.9A.A9)
- [7시스템 다시 부팅](#.EC.8B.9C.EC.8A.A4.ED.85.9C_.EB.8B.A4.EC.8B.9C_.EB.B6.80.ED.8C.85)

## 선택

이제 커널을 설정했고 컴파일했으며 필요한 시스템 설정 파일의 내용을 올바르게 채워넣었으니, 시스템을 시작할 때 커널을 실행할 프로그램을 설치할 차례입니다. 이 프로그램을 부트로더라고합니다.

사용할 부트로터는 PPC 머신의 형식에 따라 다릅니다.

NewWorld 애플 또는 IBM 머신에서는, yaboot을 선택해야 합니다. OldWorld 애플 머신에서는 BootX(추천)과 quik 이라는 두가지 선택지가 있습니다. 페가소스는 부트로더가 필요하지 않지만 SmartFirmware 부팅 메뉴를 만들려면 bootcreator를 이머지해야 합니다.

## NewWorld Macs

### GRUB

### Installation

`root #` `emerge --ask sys-boot/grub`

### Setup bootstrap partition

First, prepare the bootstrap partition that was created created during the [preparing the disk](/wiki/Handbook:PPC64/Installation/Disks "Handbook:PPC64/Installation/Disks") step. Following the example, this partition should be /dev/sda2. Optionally, confirm this by using parted:

Replace /dev/sda with the correct device if required.

`root #` `parted /dev/sda print`

```
Model: ATA Patriot Burst El (scsi)
</div>

<div lang="en" dir="ltr" class="mw-content-ltr">
Disk /dev/sda: 120GB
</div>

<div lang="en" dir="ltr" class="mw-content-ltr">
Sector size (logical/physical): 512B/512B
</div>

<div lang="en" dir="ltr" class="mw-content-ltr">
Partition Table: mac
</div>

<div lang="en" dir="ltr" class="mw-content-ltr">
Disk Flags:
</div>

<div lang="en" dir="ltr" class="mw-content-ltr">
Number  Start   End     Size    File system  Name       Flags
 1      512B    32.8kB  32.3kB               Apple
 2      32.8kB  852kB   819kB   hfs          bootstrap  boot
 3      852kB   538MB   537MB   ext4         Boot
 4      538MB   54.2GB  53.7GB  ext4         Gentoo

```

In this output, partition 2 has the bootstrap information so /dev/sda2 is the correct path to use.

Format this partition as [HFS](/wiki/HFS "HFS") using the hformat command which is part of the [sys-fs/hfsutils](https://packages.gentoo.org/packages/sys-fs/hfsutils) package:

`root #` `dd if=/dev/zero of=/dev/sda2 bs=512`

`root #` `hformat -l bootstrap /dev/sda2`

Create a directory to mount the bootstrap partition and then mount it:

`root #` `mkdir /boot/NWBB
`

`root #` `mount --types hfs /dev/sda2 /boot/NWBB`

### Setup GRUB

`root #` `grub-install --macppc-directory=/boot/NWBB /dev/sda2`

If it installs without errors, unmount the bootstrap:

`root #` `umount /boot/NWBB`

Next, bless the partition so it will boot:

`root #` `hmount /dev/sda2
`

`root #` `hattrib -t tbxi -c UNIX :System:Library:CoreServices:BootX
`

`root #` `hattrib -b :System:Library:CoreServices
`

`root #` `humount`

Finally, build the grub.cfg file:

`root #` `grub-mkconfig -o /boot/grub/grub.cfg`

## OldWorld Macs

## 대안: BootX 사용

**중요**

BootX는 MacOS 9 이전의 OldWorld 애플 시스템에서만 사용할 수 있습니다!

BootX can be downloaded at [https://github.com/immolo/BootX/archive/refs/tags/1.2.2.tar.gz](https://github.com/immolo/BootX/archive/refs/tags/1.2.2.tar.gz)

BootX는 MacOS에서 리눅스를 부팅하기 때문에 커널을 리눅스 분할 영역에서 MacOS 분할 영역으로 복사해야합니다. 우선, chroot에서 빠져나와서 MacOS 분할 영역을 마운트해야합니다. mac-fdisk -l 명령을 활용하여 MacOS 분할 영역 번호를 확인하십시오. 여기 예제에서는 sda6를 사용합니다. 분할 영역을 마운트한 후 커널을 시스템 폴더에 복사하여 BootX에서 찾을 수 있게 하겠습니다.

`root #` `exit`

`cdimage ~#` `mkdir /mnt/mac
`

`cdimage ~#` `mount /dev/sda6 /mnt/mac -t hfs
`

`cdimage ~#` `cp /mnt/gentoo/usr/src/linux/vmlinux "/mnt/mac/System Folder/Linux Kernels/kernel-6.18.8-gentoo"`

이제 커널을 복사해서 넘겼으니, BootX를 설정하기 위해 다시 부팅해야 합니다.

`cdimage ~#` `cd /
`

`cdimage ~#` `umount -l /mnt/gentoo/dev{/pts,/shm,}
`

`cdimage ~#` `umount -l /mnt/gentoo{/proc,/sys,}
`

`cdimage ~#` `umount -l /mnt/mac
`

`cdimage ~#` `reboot`

물론 부팅 CD를 제거하는걸 잊지 않으면 MacOS 시스템 대신 CD로 부팅합니다.

MacOS로 머신을 부팅하고 나면 BootX 제어판을 여십시오. genkernel을 사용하지 않는다면 옵션을 선택하시고 지정 RAM 디스크 사용 항목을 끄십시오. genkernel을 사용한다면 설치 CD의 initrd 대신 genkernel initrd를 선택했는지 확인하십시오. genkernel을 사용하지 않는다면 머신의 리눅스 루트 디스크와 분할 영역을 여기서 지정합니다. 적당한 값을 채워넣으십시오. 커널 설정에 따라 추가 부팅 매개 변수가 필요할지도 모릅니다.

BootX는 부팅 과정에 리눅스를 시작하도록 설정할 수 있습니다. 이 과정이 끝나면 MacOS로 먼저 부팅하고 시작 과정에 BootX를 불러오며 리눅스를 시작합니다. 더 많은 정보는 BootX 홈페이지를 살펴보십시오.

**중요**

HFS와 HFS+ 파일 시스템 지원을 커널에 포함했는지 확인하지 않으면 MacOS 분할 영역의 커널 업그레이드 및 업데이트가 불가능합니다.

## Pegasos

**참고**

Pegasos also has Grub support but this is currently undocumented in Gentoo. Please add this to the main wiki and notify on this page's discussion once ready to migrated here.

## 대안: BootCreator 사용

**중요**

BootCreator는 페가소스용으로 적절하게 작성한 SmartFirmware 부팅 메뉴를 만듭니다

먼저 시스템에 bootcreator를 설치했는지 확인하십시오:

`root #` `emerge --ask sys-boot/bootcreator`

이제 /etc/bootmenu/에 /etc/bootmenu/파일을 복사하고 개인의 필요에 맞춰 편집하십시오:

`root #` `cp /etc/bootmenu.example /etc/bootmenu
`

`root #` `nano -w /etc/bootmenu`

아래는 완전한 /etc/bootmenu 설정 파일입니다. vmlinux와 initrd를 커널과 initrd 이미지 이름으로 바꿔야합니다.

파일 **`/etc/bootmenu`** **bootcreator 설정 파일 예제**

```
#
# Example description file for bootcreator 1.1
#

[VERSION]
1

[TITLE]
Boot Menu

[SETTINGS]
AbortOnKey = false
Timeout    = 9
Default    = 1

[SECTION]
Local HD -> Morphos      (Normal)
ide:0 boot2.img ramdebug edebugflags="logkprintf"

[SECTION]
Local HD -> Linux (Normal)
ide:0 kernel-6.18.8-gentoo video=radeonfb:1024x768@70 root=/dev/sda3

[SECTION]
Local HD -> Genkernel (Normal)
ide:0 kernel-genkernel-ppc-6.18.8-gentoo root=/dev/ram0
root=/dev/sda3 initrd=initramfs-genkernel-ppc-6.18.8-gentoo

```

마지막으로 bootmenu는 나중에 언급한 위치로 보내야 하며 부팅 분할 영역에 복사하여 SmartFirmware에서 읽을 수 있도록 해야 합니다. 따라서 bootcreator에서 호출하도록 해야 합니다:

`root #` `bootcreator /etc/bootmenu /boot/menu`

**참고**

다시 부팅할 때 파일로서의 메뉴를 기본값대로 불러오는지 SmartFirmware의 설정을 확인하십시오.

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

새로 설치한 젠투 환경으로 다시 부팅하고 나면, [젠투 설치 마무리](/wiki/Handbook:PPC/Installation/Finalizing/ko "Handbook:PPC/Installation/Finalizing/ko") 로 끝내십시오.

[← 시스템 도구 설치](/wiki/Handbook:PPC/Installation/Tools/ko "이전 내용") [처음](/wiki/Handbook:PPC/ko "Handbook:PPC/ko") [마무리 →](/wiki/Handbook:PPC/Installation/Finalizing/ko "다음 내용")