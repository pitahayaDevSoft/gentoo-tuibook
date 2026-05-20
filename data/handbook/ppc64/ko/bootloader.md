# Bootloader

Other languages:

- [Deutsch](/wiki/Handbook:PPC64/Installation/Bootloader/de "Handbuch:PPC64/Installation/Bootloader (100% translated)")
- [English](/wiki/Handbook:PPC64/Installation/Bootloader "Handbook:PPC64/Installation/Bootloader (100% translated)")
- [español](/wiki/Handbook:PPC64/Installation/Bootloader/es "Manual de Gentoo: PPC64/Instalación/Arranque (100% translated)")
- [français](/wiki/Handbook:PPC64/Installation/Bootloader/fr "Handbook:PPC64/Installation/Bootloader/fr (100% translated)")
- [italiano](/wiki/Handbook:PPC64/Installation/Bootloader/it "Handbook:PPC64/Installation/Bootloader/it (100% translated)")
- [magyar](/wiki/Handbook:PPC64/Installation/Bootloader/hu "Handbook:PPC64/Installation/Bootloader/hu (100% translated)")
- [polski](/wiki/Handbook:PPC64/Installation/Bootloader/pl "Handbook:PPC64/Installation/Bootloader (100% translated)")
- [português do Brasil](/wiki/Handbook:PPC64/Installation/Bootloader/pt-br "Handbook:PPC64/Installation/Bootloader/pt-br (50% translated)")
- [русский](/wiki/Handbook:PPC64/Installation/Bootloader/ru "Handbook:PPC64/Installation/Bootloader (100% translated)")
- [தமிழ்](/wiki/Handbook:PPC64/Installation/Bootloader/ta "கையேடு:PPC64/நிறுவல்/துவக்கஏற்றி (100% translated)")
- [中文（中国大陆）‎](/wiki/Handbook:PPC64/Installation/Bootloader/zh-cn "手册：PPC64/安装/配置系统引导程序Bootloader (100% translated)")
- [日本語](/wiki/Handbook:PPC64/Installation/Bootloader/ja "ハンドブック:PPC64/インストール/ブートローダー (100% translated)")
- 한국어

[PPC64 핸드북](/wiki/Handbook:PPC64/ko "Handbook:PPC64/ko")[설치](/wiki/Handbook:PPC64/Full/Installation/ko "Handbook:PPC64/Full/Installation/ko")[설치 정보](/wiki/Handbook:PPC64/Installation/About/ko "Handbook:PPC64/Installation/About/ko")[매체 선택](/wiki/Handbook:PPC64/Installation/Media/ko "Handbook:PPC64/Installation/Media/ko")[네트워크 설정](/wiki/Handbook:PPC64/Installation/Networking/ko "Handbook:PPC64/Installation/Networking/ko")[디스크 준비](/wiki/Handbook:PPC64/Installation/Disks/ko "Handbook:PPC64/Installation/Disks/ko")[스테이지 3 설치](/wiki/Handbook:PPC64/Installation/Stage/ko "Handbook:PPC64/Installation/Stage/ko")[베이스 시스템 설치](/wiki/Handbook:PPC64/Installation/Base/ko "Handbook:PPC64/Installation/Base/ko")[커널 설정](/wiki/Handbook:PPC64/Installation/Kernel/ko "Handbook:PPC64/Installation/Kernel/ko")[시스템 설정](/wiki/Handbook:PPC64/Installation/System/ko "Handbook:PPC64/Installation/System/ko")[도구 설치](/wiki/Handbook:PPC64/Installation/Tools/ko "Handbook:PPC64/Installation/Tools/ko")부트로더 설정[마무리](/wiki/Handbook:PPC64/Installation/Finalizing/ko "Handbook:PPC64/Installation/Finalizing/ko")[젠투 활용](/wiki/Handbook:PPC64/Full/Working/ko "Handbook:PPC64/Full/Working/ko")[포티지 소개](/wiki/Handbook:PPC64/Working/Portage/ko "Handbook:PPC64/Working/Portage/ko")[USE 플래그](/wiki/Handbook:PPC64/Working/USE/ko "Handbook:PPC64/Working/USE/ko")[포티지 기능](/wiki/Handbook:PPC64/Working/Features/ko "Handbook:PPC64/Working/Features/ko")[초기화 스크립트 시스템](/wiki/Handbook:PPC64/Working/Initscripts/ko "Handbook:PPC64/Working/Initscripts/ko")[환경 변수](/wiki/Handbook:PPC64/Working/EnvVar/ko "Handbook:PPC64/Working/EnvVar/ko")[포티지 활용](/wiki/Handbook:PPC64/Full/Portage/ko "Handbook:PPC64/Full/Portage/ko")[파일 및 디렉터리](/wiki/Handbook:PPC64/Portage/Files/ko "Handbook:PPC64/Portage/Files/ko")[변수](/wiki/Handbook:PPC64/Portage/Variables/ko "Handbook:PPC64/Portage/Variables/ko")[소프트웨어 브랜치 함께 사용하기](/wiki/Handbook:PPC64/Portage/Branches/ko "Handbook:PPC64/Portage/Branches/ko")[추가 도구](/wiki/Handbook:PPC64/Portage/Tools/ko "Handbook:PPC64/Portage/Tools/ko")[꾸러미 저장소 개별 설정](/wiki/Handbook:PPC64/Portage/CustomTree/ko "Handbook:PPC64/Portage/CustomTree/ko")[고급 기능](/wiki/Handbook:PPC64/Portage/Advanced/ko "Handbook:PPC64/Portage/Advanced/ko")[네트워크 설정](/wiki/Handbook:PPC64/Full/Networking/ko "Handbook:PPC64/Full/Networking/ko")[시작하기](/wiki/Handbook:PPC64/Networking/Introduction/ko "Handbook:PPC64/Networking/Introduction/ko")[고급 설정](/wiki/Handbook:PPC64/Networking/Advanced/ko "Handbook:PPC64/Networking/Advanced/ko")[모듈러 네트워크](/wiki/Handbook:PPC64/Networking/Modular/ko "Handbook:PPC64/Networking/Modular/ko")[무선 네트워크](/wiki/Handbook:PPC64/Networking/Wireless/ko "Handbook:PPC64/Networking/Wireless/ko")[기능 추가](/wiki/Handbook:PPC64/Networking/Extending/ko "Handbook:PPC64/Networking/Extending/ko")[동적 관리](/wiki/Handbook:PPC64/Networking/Dynamic/ko "Handbook:PPC64/Networking/Dynamic/ko")

커널을 설정했고 컴파일했으며 필요한 시스템 설정 파일의 내용을 올바르게 채워넣었으니, 시스템을 부팅할 때 커널을 실행할 프로그램을 설치할 차례입니다. 이 프로그램을 부트로더라고합니다.

**참고**

Currently using Petitboot on Talos systems is undocumented in Gentoo. Please add the steps to [TalosII#Bootloader](/wiki/TalosII#Bootloader.2Fko "TalosII") and notify on this Discussion page when ready to merge into the Handbook.

## Contents

- [1Using GRUB](#Using_GRUB)
  - [1.1Installation](#Installation)
  - [1.2Mac hardware (G5)](#Mac_hardware_.28G5.29)
    - [1.2.1Setup bootstrap partition](#Setup_bootstrap_partition)
    - [1.2.2Setup GRUB](#Setup_GRUB)
  - [1.3IBM hardware](#IBM_hardware)
    - [1.3.1Setup GRUB](#Setup_GRUB_2)
    - [1.3.2Grub config](#Grub_config)
- [2시스템 다시 부팅](#.EC.8B.9C.EC.8A.A4.ED.85.9C_.EB.8B.A4.EC.8B.9C_.EB.B6.80.ED.8C.85)

## Using GRUB

Linux/PPC64에서는 yaBoot 를 부트로더로 사용합니다.

### Installation

`root #` `emerge --ask sys-boot/grub`

### Mac hardware (G5)

#### Setup bootstrap partition

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

#### Setup GRUB

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

### IBM hardware

Setting up Grub on IBM hardware is as simple as:

#### Setup GRUB

`root #` `grub-install /dev/sda1`

**참고**

/dev/sda1 is the PReP boot partition made in the partitioning stage

#### Grub config

Finally. build the grub.cfg file:

`root #` `grub-mkconfig -o /boot/grub/grub.cfg`

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

새로 설치한 젠투 환경으로 다시 부팅하고 나면, [젠투 설치 마무리](/wiki/Handbook:PPC64/Installation/Finalizing/ko "Handbook:PPC64/Installation/Finalizing/ko") 로 끝내십시오.

[← 시스템 도구 설치](/wiki/Handbook:PPC64/Installation/Tools/ko "이전 내용") [처음](/wiki/Handbook:PPC64/ko "Handbook:PPC64/ko") [마무리 →](/wiki/Handbook:PPC64/Installation/Finalizing/ko "다음 내용")