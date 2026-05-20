# Disks

Other languages:

- [Deutsch](/wiki/Handbook:PPC64/Installation/Disks/de "Handbuch:PPC64/Installation/Festplatten (100% translated)")
- [English](/wiki/Handbook:PPC64/Installation/Disks "Handbook:PPC64/Installation/Disks (100% translated)")
- [español](/wiki/Handbook:PPC64/Installation/Disks/es "Manual de Gentoo: PPC64/Instalación/Discos (100% translated)")
- [français](/wiki/Handbook:PPC64/Installation/Disks/fr "Handbook:PPC64/Installation/Disks/fr (100% translated)")
- [italiano](/wiki/Handbook:PPC64/Installation/Disks/it "Handbook:PPC64/Installation/Disks/it (100% translated)")
- [magyar](/wiki/Handbook:PPC64/Installation/Disks/hu "Handbook:PPC64/Installation/Disks/hu (100% translated)")
- [polski](/wiki/Handbook:PPC64/Installation/Disks/pl "Handbook:PPC64/Installation/Disks (100% translated)")
- [português do Brasil](/wiki/Handbook:PPC64/Installation/Disks/pt-br "Handbook:PPC64/Installation/Disks/pt-br (50% translated)")
- [русский](/wiki/Handbook:PPC64/Installation/Disks/ru "Handbook:PPC64/Installation/Disks (100% translated)")
- [தமிழ்](/wiki/Handbook:PPC64/Installation/Disks/ta "கையேடு:PPC64/நிறுவல்/வட்டுக்கள் (100% translated)")
- [中文（中国大陆）‎](/wiki/Handbook:PPC64/Installation/Disks/zh-cn "手册：PPC64/安装/准备磁盘 (100% translated)")
- [日本語](/wiki/Handbook:PPC64/Installation/Disks/ja "ハンドブック:PPC64/インストール/ディスク (100% translated)")
- 한국어

[PPC64 핸드북](/wiki/Handbook:PPC64/ko "Handbook:PPC64/ko")[설치](/wiki/Handbook:PPC64/Full/Installation/ko "Handbook:PPC64/Full/Installation/ko")[설치 정보](/wiki/Handbook:PPC64/Installation/About/ko "Handbook:PPC64/Installation/About/ko")[매체 선택](/wiki/Handbook:PPC64/Installation/Media/ko "Handbook:PPC64/Installation/Media/ko")[네트워크 설정](/wiki/Handbook:PPC64/Installation/Networking/ko "Handbook:PPC64/Installation/Networking/ko")디스크 준비[스테이지 3 설치](/wiki/Handbook:PPC64/Installation/Stage/ko "Handbook:PPC64/Installation/Stage/ko")[베이스 시스템 설치](/wiki/Handbook:PPC64/Installation/Base/ko "Handbook:PPC64/Installation/Base/ko")[커널 설정](/wiki/Handbook:PPC64/Installation/Kernel/ko "Handbook:PPC64/Installation/Kernel/ko")[시스템 설정](/wiki/Handbook:PPC64/Installation/System/ko "Handbook:PPC64/Installation/System/ko")[도구 설치](/wiki/Handbook:PPC64/Installation/Tools/ko "Handbook:PPC64/Installation/Tools/ko")[부트로더 설정](/wiki/Handbook:PPC64/Installation/Bootloader/ko "Handbook:PPC64/Installation/Bootloader/ko")[마무리](/wiki/Handbook:PPC64/Installation/Finalizing/ko "Handbook:PPC64/Installation/Finalizing/ko")[젠투 활용](/wiki/Handbook:PPC64/Full/Working/ko "Handbook:PPC64/Full/Working/ko")[포티지 소개](/wiki/Handbook:PPC64/Working/Portage/ko "Handbook:PPC64/Working/Portage/ko")[USE 플래그](/wiki/Handbook:PPC64/Working/USE/ko "Handbook:PPC64/Working/USE/ko")[포티지 기능](/wiki/Handbook:PPC64/Working/Features/ko "Handbook:PPC64/Working/Features/ko")[초기화 스크립트 시스템](/wiki/Handbook:PPC64/Working/Initscripts/ko "Handbook:PPC64/Working/Initscripts/ko")[환경 변수](/wiki/Handbook:PPC64/Working/EnvVar/ko "Handbook:PPC64/Working/EnvVar/ko")[포티지 활용](/wiki/Handbook:PPC64/Full/Portage/ko "Handbook:PPC64/Full/Portage/ko")[파일 및 디렉터리](/wiki/Handbook:PPC64/Portage/Files/ko "Handbook:PPC64/Portage/Files/ko")[변수](/wiki/Handbook:PPC64/Portage/Variables/ko "Handbook:PPC64/Portage/Variables/ko")[소프트웨어 브랜치 함께 사용하기](/wiki/Handbook:PPC64/Portage/Branches/ko "Handbook:PPC64/Portage/Branches/ko")[추가 도구](/wiki/Handbook:PPC64/Portage/Tools/ko "Handbook:PPC64/Portage/Tools/ko")[꾸러미 저장소 개별 설정](/wiki/Handbook:PPC64/Portage/CustomTree/ko "Handbook:PPC64/Portage/CustomTree/ko")[고급 기능](/wiki/Handbook:PPC64/Portage/Advanced/ko "Handbook:PPC64/Portage/Advanced/ko")[네트워크 설정](/wiki/Handbook:PPC64/Full/Networking/ko "Handbook:PPC64/Full/Networking/ko")[시작하기](/wiki/Handbook:PPC64/Networking/Introduction/ko "Handbook:PPC64/Networking/Introduction/ko")[고급 설정](/wiki/Handbook:PPC64/Networking/Advanced/ko "Handbook:PPC64/Networking/Advanced/ko")[모듈러 네트워크](/wiki/Handbook:PPC64/Networking/Modular/ko "Handbook:PPC64/Networking/Modular/ko")[무선 네트워크](/wiki/Handbook:PPC64/Networking/Wireless/ko "Handbook:PPC64/Networking/Wireless/ko")[기능 추가](/wiki/Handbook:PPC64/Networking/Extending/ko "Handbook:PPC64/Networking/Extending/ko")[동적 관리](/wiki/Handbook:PPC64/Networking/Dynamic/ko "Handbook:PPC64/Networking/Dynamic/ko")

## Contents

- [1블록 장치 소개](#.EB.B8.94.EB.A1.9D_.EC.9E.A5.EC.B9.98_.EC.86.8C.EA.B0.9C)
  - [1.1블록 장치](#.EB.B8.94.EB.A1.9D_.EC.9E.A5.EC.B9.98)
  - [1.2파티션과 슬라이스](#.ED.8C.8C.ED.8B.B0.EC.85.98.EA.B3.BC_.EC.8A.AC.EB.9D.BC.EC.9D.B4.EC.8A.A4)
- [2분할 배치 설계](#.EB.B6.84.ED.95.A0_.EB.B0.B0.EC.B9.98_.EC.84.A4.EA.B3.84)
  - [2.1분할 영역을 얼마나 많이, 크게 할까요?](#.EB.B6.84.ED.95.A0_.EC.98.81.EC.97.AD.EC.9D.84_.EC.96.BC.EB.A7.88.EB.82.98_.EB.A7.8E.EC.9D.B4.2C_.ED.81.AC.EA.B2.8C_.ED.95.A0.EA.B9.8C.EC.9A.94.3F)
  - [2.2스왑 공간이 무엇인가요?](#.EC.8A.A4.EC.99.91_.EA.B3.B5.EA.B0.84.EC.9D.B4_.EB.AC.B4.EC.97.87.EC.9D.B8.EA.B0.80.EC.9A.94.3F)
- [3기본: mac-fdisk 사용](#.EA.B8.B0.EB.B3.B8:_mac-fdisk_.EC.82.AC.EC.9A.A9)
- [4대안: fdisk 사용](#.EB.8C.80.EC.95.88:_fdisk_.EC.82.AC.EC.9A.A9)
  - [4.1현재 파티션 배치 보기](#.ED.98.84.EC.9E.AC_.ED.8C.8C.ED.8B.B0.EC.85.98_.EB.B0.B0.EC.B9.98_.EB.B3.B4.EA.B8.B0)
  - [4.2모든 파티션 제거](#.EB.AA.A8.EB.93.A0_.ED.8C.8C.ED.8B.B0.EC.85.98_.EC.A0.9C.EA.B1.B0)
  - [4.3PPC PReP 부트 파티션 만들기](#PPC_PReP_.EB.B6.80.ED.8A.B8_.ED.8C.8C.ED.8B.B0.EC.85.98_.EB.A7.8C.EB.93.A4.EA.B8.B0)
  - [4.4스왑 파티션 만들기](#.EC.8A.A4.EC.99.91_.ED.8C.8C.ED.8B.B0.EC.85.98_.EB.A7.8C.EB.93.A4.EA.B8.B0)
  - [4.5루트 파티션 만들기](#.EB.A3.A8.ED.8A.B8_.ED.8C.8C.ED.8B.B0.EC.85.98_.EB.A7.8C.EB.93.A4.EA.B8.B0)
  - [4.6파티션 배치 저장하기](#.ED.8C.8C.ED.8B.B0.EC.85.98_.EB.B0.B0.EC.B9.98_.EC.A0.80.EC.9E.A5.ED.95.98.EA.B8.B0)
- [5파일 시스템 만들기](#.ED.8C.8C.EC.9D.BC_.EC.8B.9C.EC.8A.A4.ED.85.9C_.EB.A7.8C.EB.93.A4.EA.B8.B0)
  - [5.1도입부](#.EB.8F.84.EC.9E.85.EB.B6.80)
  - [5.2파일 시스템](#.ED.8C.8C.EC.9D.BC_.EC.8B.9C.EC.8A.A4.ED.85.9C)
  - [5.3분할 영역에 파일 시스템 반영하기](#.EB.B6.84.ED.95.A0_.EC.98.81.EC.97.AD.EC.97.90_.ED.8C.8C.EC.9D.BC_.EC.8B.9C.EC.8A.A4.ED.85.9C_.EB.B0.98.EC.98.81.ED.95.98.EA.B8.B0)
    - [5.3.1EFI system partition filesystem](#EFI_system_partition_filesystem)
    - [5.3.2Legacy BIOS boot partition filesystem](#Legacy_BIOS_boot_partition_filesystem)
    - [5.3.3Small ext4 partitions](#Small_ext4_partitions)
  - [5.4스왑 분할 영역 활성화](#.EC.8A.A4.EC.99.91_.EB.B6.84.ED.95.A0_.EC.98.81.EC.97.AD_.ED.99.9C.EC.84.B1.ED.99.94)
- [6루트 분할 영역 마운트](#.EB.A3.A8.ED.8A.B8_.EB.B6.84.ED.95.A0_.EC.98.81.EC.97.AD_.EB.A7.88.EC.9A.B4.ED.8A.B8)

## 블록 장치 소개

### 블록 장치

리눅스 파일 시스템, 분할 영역, 블록 장치 등 젠투 리눅스 및 일반적인 리눅스 운영체제의 바람직한 디스크 측면의 양상을 살펴보도록 하겠습니다. 디스크와 파일 시스템의 입출력을 이해하고 나서, 젠투 리눅스 설치에 필요한 분할 영역과 파일 시스템을 설정하겠습니다.

시작에 앞서 블록 장치를 살펴보도록 하죠. 아마~도 리눅스 시스템에서 첫번째 드라이브로 표시하는 대부분 잘 알려진 블록 장치는 /dev/sda겠죠. SCSI와 직렬 ATA 드라이브 둘 다 /dev/sd\*와 같은 식으로 표시합니다. 게다가 커널의 libata 프레임워크에서는 IDE 드라이브도 마찬가지로 /dev/sd\*로 표시합니다. 이전 장치 프레임워크에서 첫번째 IDE 드라이브는 /dev/hda입니다.

The following table will help readers determine where to find a certain type of block device on the system:

Type of deviceDefault device handleEditorial notes and considerations
IDE, SATA, SAS, SCSI, or USB flash/dev/sdaFound on hardware from roughly 2007 until the present, this device handle is perhaps the most commonly used in Linux. These types of devices can be connected via the [SATA bus](https://en.wikipedia.org/wiki/Serial_ATA "wikipedia:Serial ATA"), [SCSI](https://en.wikipedia.org/wiki/SCSI "wikipedia:SCSI"), [USB](https://en.wikipedia.org/wiki/USB "wikipedia:USB") bus as block storage. As example, the first partition on the first SATA device is called /dev/sda1.
NVM Express (NVMe)/dev/nvme0n1The latest in solid state technology, [NVMe](https://en.wikipedia.org/wiki/NVM_Express "wikipedia:NVM Express") drives are connected to the PCI Express bus and have the fastest transfer block speeds on the market. Systems from around 2014 and newer may have support for NVMe hardware. The first partition on the first NVMe device is called /dev/nvme0n1p1.
MMC, eMMC, and SD/dev/mmcblk0[embedded MMC](https://en.wikipedia.org/wiki/MultiMediaCard#eMMC "wikipedia:MultiMediaCard") devices, SD cards, and other types of memory cards can be useful for data storage. That said, many systems may not permit booting from these types of devices. It is suggested to **not** use these devices for active Linux installations; rather consider using them to transfer files, which is their typical design intention. Alternatively this storage type could be useful for short-term file backups or snapshots.
VirtIO/dev/vda[Virtualized](/wiki/Virtualization "Virtualization") devices are found only within a [QEMU](/wiki/QEMU "QEMU") virtual emulator. `virtio_blk` driver provides this access to host's allocated drive space for this virtual machine. That said, it is a great way to try out Gentoo inside a virtual machine.

위에 나타낸 블록 장치는 디스크의 추상 인터페이스를 표현합니다. 사용자 프로그램은 블록 장치가 IDE가 됐든 SCSI가 됐든 뭐가 됐든지간에 신경쓰지 않고 디스크와 소통을 수행할 때 이 블록 장치를 사용할 수 있습니다. 프로그램에서는 디스크의 저장 공간에 대해, 연속적이며, 임의로 접근하는 512 바이트 블록의 모음으로 다룰 수 있습니다.

### 파티션과 슬라이스

비록 이론적으로 리눅스 시스템을 저장할 디스크 전체를 사용할 수 있다고는 하지만, 실제로는 거의 불가능합니다. 대신, 전체 디스크 블록 장치를 더 작게 나누고, 더 관리하기 쉬운 블록 장치로 만들 수 있습니다. 대부분의 시스템에서는 파티션이라고 합니다. 다른 아키텍처에서는 _슬라이스_ 라고 하는 비슷한 기술이 있습니다.

## 분할 배치 설계

### 분할 영역을 얼마나 많이, 크게 할까요?

분할 영역의 수는 환경에 따라 다릅니다. 예를 들어, 사용자가 많을 경우 보안성을 개선하고 백업을 쉽게 하기 위해 /home/을 나누는 것이 좋습니다. 젠투를 메일 서버로 설치한다면, /var/에 모든 메일을 저장하므로 /var/를 나누어야 합니다. 파일 시스템의 탁월한 선택은 성능을 극대화합니다. 게임 서버는 게임 서버를 설치할 /opt/를 따로 나눕니다. 이유는 /home/과 비슷합니다: 보안과 백업이죠. 대부분의 상황에서 /usr/는 거대한 상태고 남아있습니다. 주요 프로그램을 저장할 뿐만 아니라, (보통 /usr/portage에 기본으로 들어가는) 젠투 이빌드 저장소는 거의 650MB를 차지합니다. 이 디스크 공간은 보통 이빌드 저장소내에 저장하는 packages/와 distfiles/ 디렉터리는 _제외_ 하고 추산합니다.

In most situations on Gentoo, /usr and /var should be kept relatively large in size. /usr hosts the majority of applications available on the system and the Linux kernel sources (under /usr/src). By default, /var hosts the Gentoo ebuild repository (located at /var/db/repos/gentoo) which, depending on the file system, generally consumes around 650 MiB of disk space. This space estimate _excludes_ the /var/cache/distfiles and /var/cache/binpkgs directories, which will gradually fill with source files and (optionally) binary packages respectively as they are added to the system.

관리자 취향에 달려있습니다. 분할 영역 또는 볼륨을 나누면 다음과 같은 장점이 있습니다:

- 각 분할 영역 또는 볼륨에 대해 최상의 동작을 수행하는 파일 시스템을 선택합니다.
- 제 기능을 상실한 도구가 분할 영역 또는 볼륨에 계속 파일을 기록할 경우, 남아 있는 공간이 없어져 전체 시스템이 동작하지 않습니다.
- 필요한 경우, (이 장점은 여러 개의 분할 영역보다는 여러 대의 디스크에서 더 돋보이지만) 동시에 여러 분할 영역을 검사할 수 있어, 파일 시스템 검사 시간을 줄일 수 있습니다.
- 일부 분할 영역 또는 볼륨을 읽기 전용, `nosuid`(setuid 무시), `noexec`(실행 비트 무시) 등으로 마운트하여 보안성을 개선할 수 있습니다.

그러나, 마찬가지로 다중 분할 영역에는 단점도 존재합니다. 제대로 설정하지 않으면 어떤 분할 영역에는 공간이 상당히 남지만, 다른 분할 영역은 그렇지 않을 수 있습니다. 다른 골칫거리는 분할 영역이 나뉘어져 있는 상황입니다. /usr/ 또는 /var/와 같은 중요한 마운트 지점은 특히 그렇습니다. 다른 부팅 스크립트를 시작하기 전에 분할 영역을 마운트하려면 관리자가 종종 initramfs로 부팅해야합니다. 항상 있는 경우는 아니기 때문에 결과가 다양하게 나타납니다.

디스크에서 GPT 레이블을 사용하지 않으면 SCSI와 SATA에서는 분할 영역 갯수가 15개로 제한되어있습니다.

**참고**

Installations that intend to use systemd as the service and init system must have the /usr directory available at boot, either as part of the root filesystem or mounted via an initramfs.

### 스왑 공간이 무엇인가요?

Recommendations for swap space size
RAM sizeSuspend support?Hibernation support?
2 GB or less4GB4GB
2 to 8 GBRAM amount2 \* RAM
8 to 64 GB8 GB minimum, 16 maximum1.5 \* RAM
64 GB or greater8 GB minimumHibernation not recommended! Hibernation is _not_ recommended for systems with very large amounts of memory. While possible, the entire contents of memory must be written to disk in order to successfully hibernate. Writing tens of gigabytes (or worse!) out to disk can can take a considerable amount of time, especially when rotational disks are used. It is best to suspend in this scenario.

완벽한 스왑 분할 영역 값은 없습니다. 스왑 영역의 존재 목적은 내부 메모리(RAM)가 용량 고갈에 처해있을 때 커널에서 디스크 공간을 제공하려는 것입니다. 스왑 영역은 커널에서 곧 접근하지 않을 메모리 페이지를 디스크(스왑 또는 페이지-아웃)에 옮기고 메모리를 확보할 수 있도록 합니다. 물론 메모리가 갑자기 필요할 때도 이 페이지를 메모리에 되돌려놓습니다만(페이지-인), 시간이 오래걸립니다(내부 메모리에 비해 디스크는 비교적 매우 느립니다).

시스템이 메모리를 집중적으로 사용하는 프로그램을 실행하려 하지 않거나 시스템에 충분한 메모리가 있을 경우 많은 스왑 영역이 필요하지 않을지도 모릅니다. 그러나 스왑 영역은 최대 절전모드 기능을 사용할 경우 전체 메모리 공간을 사용하기도 합니다. 시스템을 최대 절전모드로 진입하려 한다면, 더 큰 스왑 영역이 필요하며, 최소한, 종종 시스템에 대용량의 메모리를 설치합니다.

## 기본: mac-fdisk 사용

**중요**

이 절차는 애플 G5 시스템용입니다.

Partition
Description
/dev/sda1Apple partition map, automatically created when the disk is formatted with a "mac" partition table.
/dev/sda2New World boot block
/dev/sda3Swap partition
/dev/sda4Root partition

mac-fdisk를 시작하십시오:

`root #` `mac-fdisk /dev/sda`

리눅스 파티션을 만들려면 우선 앞서 지운 파티션을 삭제하십시오. 이 파티션을 삭제하려면 mac-fdisk에서 `d` 키를 사용하십시오. 삭제할 파티션 번호를 물어봅니다.

그 다음, `b` 키를 사용하여 Apple\_Bootstrap 파티션을 만드십시오. 시작 블록 위치를 물어봅니다. 처음 남은 파티션의 번호를 입력하고 `p` 를 입력하십시오. 이 경우, "2p"입니다.

**참고**

이 파티션은 "boot" 파티션이 아닙니다. 모든 파티션이 리눅스가 쓰진 않습니다. 이 파티션에 어떤 파일 시스템도 올라갈 필요가 없으며 마운트조차도 해서는 안됩니다. PPC 사용자는 /boot에 대한 추가 파티션이 필요없습니다.

이제 `c` 를 눌러 스왑 파티션을 만드십시오. 다시 말해, mac-fdisk에서는 몇번째 블록에서 시작하는지 물어봅니다. Apple\_Bootstrap 파티션을 만들때 2번을 입력했던바와 같이, "3p"를 입력하십시오 크기를 물어보면 512M(또는 필요한 용량만큼)을 입력하십시오. 이름을 물어보면 "swap"(필수)이라고 입력하십시오.

루트 파티션을 만들려면 `c` 를 입력하십시오. 그 다음 루트 파티션의 블록 시작 위치를 선택할 때 "4p"라고 입력하십시오. 크기를 물어보면 "4p"를 다시 입력하십시오. mac-fdisk에서는 이 입력을 "존재하는 전체 공간 사용"이란 뜻으로 해석합니다. 이름을 물어보면 "root"(필수)라고 입력하십시오.

끝내려면 `w` 로 파티션 정보를 디스크에 기록하고 `q` 로 mac-fdisk를 끝내십시오.

**참고**

모든 부분이 문제없는지 확인하려면, mac-fdisk를 한번 이상 실행하고 모든 파티션이 제자리에 위치했는지 검사하십시오. 만들어놓은 모든 파티션이 나타나지 않거나, 바뀐 내용 일부가 보이지 않는다면 mac-fdisk에서 `i` 를 눌러 파티션을 다시 초기화하십시오. 참고로 이 명령은 파티션 맵을 다시 만들기 때문에 모든 파티션을 제거합니다.

## 대안: fdisk 사용

**중요**

이 절차는 IBM p 시리즈, i 시리즈, 오픈파워 시스템에 해당합니다.

**참고**

POWER5 기반 하드웨어에서 젠투 설치 과정에 RAID 디스크 어레이를 사용하려 한다면, 고급 기능 포맷 방식으로 iprconfig를 실행하여 디스크 어레이를 만드십시오. 설치가 끝나면 [sys-fs/iprutils](https://packages.gentoo.org/packages/sys-fs/iprutils)를 이머지하십시오.

ipr 기반의 SCSI 어댑터를 사용한다면 ipr 유틸리티를 시작하십시오.

`root #` `/etc/init.d/iprinit start`

다음 내용에서는 이전에 설명한 파티션 배치를 만드는 방법의 예제를 분명하게 설명합니다:

파티션
설명
/dev/sda1PPC PReP 부트 파티션
/dev/sda2스왑 파티션
/dev/sda3루트 파티션

개인 취향에 따라 파티션 배치를 바꾸십시오.

### 현재 파티션 배치 보기

[fdisk](/wiki/Fdisk "Fdisk")는 디스크를 파티션으로 나누는 저명하고 강력한 도구입니다. fdisk를 현재 디스크에서 실행해보십시오(예제에서는 /dev/sda를 사용):

`root #` `fdisk /dev/sda`

```
Command (m for help)

```

시스템에 AIX 파티션이 남아있다면 다음 오류 메시지가 나타납니다:

`root #` `fdisk /dev/sda`

```
  There is a valid AIX label on this disk.
  Unfortunately Linux cannot handle these
  disks at the moment.  Nevertheless some
  advice:
  1. fdisk will destroy its contents on write.
  2. Be sure that this disk is NOT a still vital
     part of a volume group. (Otherwise you may
     erase the other disks as well, if unmirrored.)
  3. Before deleting this physical volume be sure
     to remove the disk logically from your AIX
     machine.  (Otherwise you become an AIXpert).

```

걱정 마시고 `o` 를 누르면 비어있는 새 DOS 파티션 테이블을 만들 수 있습니다.

**경고**

이 과정은 설치한 AIX를 망가뜨립니다!

`p` 를 입력하여 현재 파티션 설정을 표시하십시오:

`Command (m for help):` `p`

```
Disk /dev/sda: 30.7 GB, 30750031872 bytes
141 heads, 63 sectors/track, 6761 cylinders
Units = cylinders of 8883 * 512 = 4548096 bytes

   Device Boot      Start         End      Blocks   Id  System
/dev/sda1               1          12       53266+  83  Linux
/dev/sda2              13         233      981571+  82  Linux swap
/dev/sda3             234         674     1958701+  83  Linux
/dev/sda4             675        6761    27035410+   5  Extended
/dev/sda5             675        2874     9771268+  83  Linux
/dev/sda6            2875        2919      199836   83  Linux
/dev/sda7            2920        3008      395262   83  Linux
/dev/sda8            3009        6761    16668918   83  Linux

```

각각의 디스크에는 하나의 스왑 파티션과("Linux Swap"으로 나타남) 6개의 리눅스 시스템(각각의 파티션이 "Linux"라고 나타남)을 넣도록 설정했습니다.

### 모든 파티션 제거

디스크에서 모든 파티션을 제거하십시오. `d` 를 입력하여 파티션을 삭제하십시오. 기존의 /dev/sda1 파티션을 삭제하려면:

`Command (m for help):` `d`

```
Partition number (1-4): 1

```

파티션은 삭제 예정입니다. `p` 를 누르면 더이상 나타나지 않습니다만, 저장하기 전에는 지워지지 않습니다. 실수해서 세션을 멈춰야한다면, `q` 를 입력하고 `Enter` 를 입력하여 어떤 파티션도 삭제하거나 수정하지 않게 하십시오.

이제 의도한 모든 파티션을 날려야 한다고 간주하고, `p` 를 입력하여 파티션 목록을 출력하고 `d` 를 입력한 후 삭제할 파티션의 번호를 입력하는 이 과정을 반복하십시오. 최종적으로, 파티션 테이블에는 어떤 파티션도 나타나지 않습니다:

`Command (m for help):` `p`

```
Disk /dev/sda: 30.7 GB, 30750031872 bytes
141 heads, 63 sectors/track, 6761 cylinders
Units = cylinders of 8883 * 512 = 4548096 bytes

Device Boot    Start       End    Blocks   Id  System

```

이제 파티션 테이블을 비웠으니 파티션을 만들어보겠습니다. 앞에서 이야기한 대로 기본 파티션 형태를 사용하겠습니다. 물론, 문자 그대로 따라오시지 말고 개인 취향에 따라 적절히 조절하세요.

### PPC PReP 부트 파티션 만들기

우선 작은 PReP 부트 파티션을 만들겠습니다. `n` 을 입력하여 새 파티션을 만들고, `p` 입력하여 주 파티션을 선택한 다음, `1` 을 입력하여 첫번째 주 파티션을 선택하십시오. 첫번째 실린더를 물어보면 `Enter` 를 치십시오. 마지막 실린더를 물어보면 +7M을 입력하여 7MB 크기의 파티션을 만드십시오. 여기가지 과정이 끝나면 `t` 를 입력하여 파티션 형식을 바꾸시고, 방금 만든 파티션 `1` 번을 선택하신 후, `41` 을 입력하여 "PPC PReP Boot"를 선택하십시오. 마지막으로, PReP 파티션을 부팅 가능하도록 표시하십시오.

**참고**

PReP 파티션 크기는 8MB 미만이어야 합니다!

`Command (m for help):` `p`

```
Disk /dev/sda: 30.7 GB, 30750031872 bytes
141 heads, 63 sectors/track, 6761 cylinders
Units = cylinders of 8883 * 512 = 4548096 bytes

   Device Boot      Start         End      Blocks   Id  System

```

`Command (m for help):` `n`

```
Command action
      e   extended
      p   primary partition (1-4)
p
Partition number (1-4): 1
First cylinder (1-6761, default 1):
Using default value 1
Last cylinder or +size or +sizeM or +sizeK (1-6761, default
6761): +8M

```

`Command (m for help):` `t`

```
Selected partition 1
Hex code (type L to list codes): 41
Changed system type of partition 1 to 41 (PPC PReP Boot)

```

`Command (m for help):` `a`

```
Partition number (1-4): 1
Command (m for help):

```

이제 ( `p` 명령으로) 파티션 테이블을 다시 살펴봤을 때, 다음 파티션 정보가 나타나야합니다:

`Command (m for help):` `p`

```
Disk /dev/sda: 30.7 GB, 30750031872 bytes
141 heads, 63 sectors/track, 6761 cylinders
Units = cylinders of 8883 * 512 = 4548096 bytes

   Device Boot      Start         End      Blocks   Id  System
/dev/sda1  *            1           3       13293   41  PPC PReP Boot

```

### 스왑 파티션 만들기

다음 스왑 파티션을 만들도록 하겠습니다. 스왑 파티션을 만들려면 `n` 키를 입력하여 새 파티션을 만들고, 우리 같은 경우는 /dev/sda2를 만들겠으니 `2` 를 입력하여 두번째 파티션을 만듭니다. 첫 실린더를 물어보면 그냥 `Enter` 키를 칩니다. 마지막 실린더를 물어보면 +512M을 입력하여 512MB 크기의 파티션을 만듭니다. 그 다음 파티션 형식을 설정하려면 `t` 를 입력하고, 방금 만든 파티션 `2` 를 선택한 후, 파티션 형식을 _Linux Swap_ 으로 바꾸기 위해 82를 입력합니다. 이 단계가 끝나면 `p` 를 입력하여 다음과 같은 파티션 표를 출력해야합니다:

`Command (m for help):` `p`

```
Disk /dev/sda: 30.7 GB, 30750031872 bytes
141 heads, 63 sectors/track, 6761 cylinders
Units = cylinders of 8883 * 512 = 4548096 bytes

   Device Boot      Start         End      Blocks   Id  System
/dev/sda1               1           3       13293   41  PPC PReP Boot
/dev/sda2               4         117      506331   82  Linux swap

```

### 루트 파티션 만들기

마지막으로, 루트 파티션을 만들겠습니다. 과정을 진행하려면, 새 파티션을 만드는 명령 `n` 을 입력하시고, 주 파티션을 만들겠다고 `fdisk` 에게 알리기 위해 `p` 를 입력한 후, 세번째 주 파티션 /dev/sda3를 만들기 위해 `3` 을 입력하십시오. 첫번째 섹터를 물어보면 `Enter` 를 치십시오. 마지막 섹터를 물어보면 `Enter` 를 쳐서 디스크의 나머지 공간을 취하는 파티션을 만드십시오. 이 과정이 끝나면 `p` 를 입력하였을 때 다음과 같은 파티션 테이블 모습이 나타나야합니다:

`Command (m for help):` `p`

```
Disk /dev/sda: 30.7 GB, 30750031872 bytes
141 heads, 63 sectors/track, 6761 cylinders
Units = cylinders of 8883 * 512 = 4548096 bytes

   Device Boot      Start         End      Blocks   Id  System
/dev/sda1               1           3       13293   41  PPC PReP Boot
/dev/sda2               4         117      506331   82  Linux swap
/dev/sda3             118        6761    29509326   83  Linux

```

### 파티션 배치 저장하기

파티션 배치를 저장하고 fdisk를 빠져나가려면 `w` 를 입력하십시오.

`Command (m for help):` `w`

## 파일 시스템 만들기

**경고**

When using an SSD or NVMe drive, it is wise to check for firmware upgrades. Some Intel SSDs in particular (600p and 6000p) require a firmware upgrade for [possible data corruption](https://bugzilla.redhat.com/show_bug.cgi?id=1402533) induced by XFS I/O usage patterns. The problem is at the firmware level and not any fault of the XFS filesystem. The smartctl utility can help check the device model and firmware version.

### 도입부

이제 분할 영역을 만들었고, 파일 시스템을 제 위치에 얹어놓을 차례입니다. 다음 절에서는 리눅스에서 지원하는 다양한 파일 시스템을 설명합니다. 어떤 파일 시스템을 사용할 지 이미 알고 있는 독자라면 [파티션에 파일 시스템 반영하기](#.ED.8C.8C.ED.8B.B0.EC.85.98.EC.97.90_.ED.8C.8C.EC.9D.BC_.EC.8B.9C.EC.8A.A4.ED.85.9C_.EB.B0.98.EC.98.81.ED.95.98.EA.B8.B0) 로 계속 진행할 수 있습니다. 그렇지 않으면 계속 읽어 내려가면서 쓸 수 있는 파일시스템이 어떤 종류가 있는지 알아보십시오.

### 파일 시스템

다양한 파일 시스템이 있습니다. 일부는 ppc64 아키텍처에서 안정적입니다 - 중요한 분할 영역을 위해서라면 좀 더 시험적인 분할 영역을 선택하기 전에 파일 시스템과 지원 상태에 대한 내용을 좀 더 읽어보시는 것이 좋겠습니다.

btrfs스냅샷, 검사합을 통한 자체복구, 투명 압축, 하위 볼륨, 통합 RAID 같은 고급 기능을 제공하는 차세대 파일 시스템입니다. 일부 배포판은 이미 특별한 옵션으로 탑재했지만 실무에서 쓰기엔 준비가 미흡합니다. 파일 시스템이 깨지는 경우가 다반사입니다. 개발자들은 이전 버전에 문제가 있기 때문에 안전을 위해 최신 커널 버전을 사용하라고 합니다. 몇년 동안 이래왔고 무엇인가 바뀐다고 하면 너무 일찍 언급합니다. 깨지는 문제를 고친다고 하면 가끔 이전 커널에 있던 대로 돌아갑니다. 파일 시스템을 쓰려 한다면 위험을 감수하십시오!ext2검증된 리눅스 파일시스템이지만 메타데이터 저널링기능이 없습니다. 이는 시작시간의 파일시스템 검사루틴에서 조금 더 많은 시간소모를 할 수 있다는 의미입니다. 이제 일관성 검사를 더욱 빠르게 할 수 있고 비 저널링의 대체 수단으로써 일반적으로 더욱 선호하는 차세대 저널링 파일시스템의 상당한 선택요소가 있습니다. 저널링 파일시스템은 시스템을 시동하고 파일시스템에 비일관 상태가 발생했을 때 긴 지연시간을 줄입니다.ext3빠른 복구 기능을 제공하는 메타데이터 저널링을 제공하며, 게다가 전체 데이터와 정렬된 데이터 저널링과 같은 강화 저널링 모드도 지원하는 ext2 파일시스템의 저널링 버전입니다. 대부분의 모든 상황에서 고성능 동작이 가능한 HTree 색인을 사용합니다. 간단히 말해 ext3는 아주 좋은 믿을 수 있는 파일시스템입니다. ext3을 모든 목적의 모든 플랫폼 파일시스템으로 추천합니다.ext4ext3으로부터 갈라져 나와 성능을 향상시키고 디스크상 형식에 대해 적절한 수정을 가하여 용량 제한을 없애는 새로운 기능을 포함하여 만든 파일시스템입니다. 볼륨 하나의 크기를 1EB까지 늘릴 수 있고, 파일 최대 크기는 16TB가 될 수 있습니다. 기존의 ext2/3 비트맵 블록 할당 대신에 ext4는 대용량 파일 성능을 끌어올리고 단편화를 줄인 extents를 사용합니다. ext4는 디스크의 데이터 배치에 대해 최적화 할 더 많은 방법을 파일시스템 드라이버에 제공하는 좀 더 세련된 블록 할당 알고리즘(지연할당 및 다중블록 할당)을 제공합니다. ext4는 모든 목적의 모든 플랫폼의 파일 시스템에 추천합니다.f2fs플래시 지향 파일 시스템은 처음에 낸드 플래시 메모리에서 활용할 목적으로 삼성에서 만들었습니다. 2016년 2/4분기 시점에, 이 파일 시스템은 여전히 미완의 상태지만 젠투를 마이크로SD 카드, USB 드라이브, 기타 플래시 기반 저장 장치에 설치할 경우 괜찮은 선택입니다.JFSIBM의 고성능 저널링 파일시스템입니다. JFS는 다양한 상황속에서도 좋은 성능을 내는, 가볍고 빠르며 믿을 수 있는 B+트리 기반 파일시스템입니다.ReiserFS전반적으로 좋은 성능을 내며 특히 용량이 작은 수많은 파일들을 다룰 때 더 많은 CPU 사이클을 소비하는 경우 좋은 성능이 나는 B+트리 기반 저널링 파일시스템입니다. ReiserFS는 다른 파일시스템보다 덜 관리중인 것으로 보입니다.XFS견고한 기능 모음을 지니고 있으며 확장성에 있어 최적화 된 메타데이터 저널링 파일시스템입니다. XFS는 다양한 하드웨어 문제에 대해 그다지 관대하진 않은 것 같습니다.vfatFAT32로 알려진 vfat은 리눅스에서 지원하지만 권한 설정은 지원하지 않습니다. 여러 운영 체제간 상호 운용성을 목적으로(주로 마이크로소프트 윈도우) 활용하지만 일부 시스템 펌웨어(UEFI)용으로도 필요합니다.NTFS"New Technology" 파일 시스템은 마이크로 소프트의 대표 파일 시스템입니다. 위의 vfat과 비슷하게 BSD 또는 리눅스에서 필요한 권한 설정 또는 확장 속성을 저장하지 않기에 루트 파일 시스템으로 활용할 수 없습니다. 오직 마이크로소프트 윈도우와 상호 연동할 _때만_ 활용해야합니다(오직 이 경우에만 역점을 둠을 참고하십시오).

More extensive information on filesystems can be found in the community maintained [Filesystem article](/wiki/Filesystem/ko "Filesystem/ko").

### 분할 영역에 파일 시스템 반영하기

**참고**

Please make sure to emerge the relevant user space utilities package for the chosen filesystem before rebooting. There will be a reminder to do so near the end of the installation process.

분할 영역 또는 볼륨에 파일 시스템을 만들 때, 각 파일 시스템에서 사용할 수 있는 도구가 있습니다. 각 파일 시스템의 추가 정보를 살펴보려면 하단 표의 파일 시스템 이름을 누르십시오:

파일시스템
구성 명령
최소 CD 포함?
꾸러미
[btrfs](/wiki/Btrfs "Btrfs")mkfs.btrfs Yes
[sys-fs/btrfs-progs](https://packages.gentoo.org/packages/sys-fs/btrfs-progs)[ext2](/wiki/Ext2 "Ext2")mkfs.ext2 Yes
[sys-fs/e2fsprogs](https://packages.gentoo.org/packages/sys-fs/e2fsprogs)[ext3](/wiki/Ext3 "Ext3")mkfs.ext3 Yes
[sys-fs/e2fsprogs](https://packages.gentoo.org/packages/sys-fs/e2fsprogs)[ext4](/wiki/Ext4 "Ext4")mkfs.ext4 Yes
[sys-fs/e2fsprogs](https://packages.gentoo.org/packages/sys-fs/e2fsprogs)[f2fs](/wiki/F2fs "F2fs")mkfs.f2fs Yes
[sys-fs/f2fs-tools](https://packages.gentoo.org/packages/sys-fs/f2fs-tools)[jfs](/wiki/JFS "JFS")mkfs.jfs Yes
[sys-fs/jfsutils](https://packages.gentoo.org/packages/sys-fs/jfsutils)[reiserfs](/wiki/ReiserFS "ReiserFS")mkfs.reiserfs Yes
[sys-fs/reiserfsprogs](https://packages.gentoo.org/packages/sys-fs/reiserfsprogs)[xfs](/wiki/XFS "XFS")mkfs.xfs Yes
[sys-fs/xfsprogs](https://packages.gentoo.org/packages/sys-fs/xfsprogs)[vfat](/wiki/FAT "FAT")mkfs.vfat Yes
[sys-fs/dosfstools](https://packages.gentoo.org/packages/sys-fs/dosfstools)[NTFS](/wiki/NTFS "NTFS")mkfs.ntfs Yes
[sys-fs/ntfs3g](https://packages.gentoo.org/packages/sys-fs/ntfs3g)

\|}

**중요**

The handbook recommends new partitions as part of the installation process, but it is important to note running any mkfs command will erase any data contained within the partition. When necessary, ensure any data that exists within is appropriately backed up before creating a new filesystem.

예를 들어, 예제 분할 영역 구조와 같이 ext2 형식의 부팅 분할 영역 (/dev/sda1) 과 ext4 형식의 루트 분할 영역 (/dev/sda3)을 취하려면, 다음 명령을 사용할 수 있습니다:

`root #` `mkfs.ext4 /dev/sda3`

{{#ifeq: 0\|1\|

#### EFI system partition filesystem

The EFI system partition (/dev/sda1) must be formatted as FAT32:

`root #` `mkfs.ext2 /dev/sda1`

#### Legacy BIOS boot partition filesystem

Systems booting via legacy BIOS with a MBR/DOS disklabel can use any filesystem format supported by the bootloader.

For example, to format with XFS:

`root #` `mkfs.xfs -c options=/usr/share/xfsprogs/mkfs/lts_6.18.conf /dev/sda1`

#### Small ext4 partitions

(8GB 이하의) 작은 분할 영역에서 ext2, ext3, ext4 를 사용한다면, 충분한 inode 갯수를 예약할 적당한 옵션으로 파일 시스템을 만들어야합니다. mke2fs(mkfs.ext2)에서는 "아이노드 당 바이트" 설정을 사용하여 파일 시스템에서 보유할 아이노드 갯수를 계산합니다. 작은 분할 영역일수록 아이노드 갯수를 늘리는 것이 좋습니다.

`root #` `mkfs.ext2 -T small /dev/<device>`

각 16kB 영역을 하나의 4kB 영역으로 줄이는 "아이노드 당 바이트"로 주어진 파일 시스템의 아이노드 갯수를 네 배로 뻥튀기(?)합니다. 비율값을 부여하여 속성을 조절할 수 있습니다:

### 스왑 분할 영역 활성화

mkswap은 스왑 분할 영역을 초기화하는 명령입니다:

`root #` `mkswap /dev/sda2`

**참고**

Installations which were previously started, but did not finish the installation process can resume the installation from this point in the handbook. Use this link as the permalink: [Resumed installations start here](/wiki/Handbook:PPC64/Installation/Disks/ko#Resumed_installations_start_here "Handbook:PPC64/Installation/Disks/ko").

스왑 분할 영역을 활성화하려면, swapon 명령을 사용하십시오:

`root #` `swapon /dev/sda2`

This 'activation' step is only necessary because the swap partition is newly created within the live environment. Once the system has been rebooted, as long as the swap partition is properly defined within fstab or other mount mechanism, swap space will activate automatically.

## 루트 분할 영역 마운트

Certain live environments may be missing the suggested mount point for Gentoo's root partition (/mnt/gentoo), or mount points for additional partitions created in the partitioning section:

`root #` `mkdir --parents /mnt/gentoo`

Continue creating additional mount points necessary for any additional (custom) partition(s) created during previous steps by using the mkdir command.

이제 분할 영역을 초기화했고 파일 시스템을 넣었으므로 분할 영역을 마운트할 차례입니다. mount 명령을 사용하지만 만들어놓은 모든 분할 영역에 대해 마운트 디렉터리를 만들 필요는 없다는 사실을 잊지 마십시오. 예제를 통해 우리는 루트 분할 영역을 마운트하겠습니다:

Mount the root partition:

`root #` `mount /dev/sda3 /mnt/gentoo`

Continue mounting additional (custom) partitions as necessary using the mount command.

**참고**

/tmp/를 따로 나눈 분할 영역에 두어야 한다면, 마운트하기 전에 퍼미션을 바꾸었는지 확인하십시오:

`root #` `chmod 1777 /mnt/gentoo/tmp`

이 설정은 /var/tmp에도 적용 유지합니다.

지침을 따르고 나면 proc 파일 시스템(커널 가상 인터페이스)와 다른 커널 의사 파일 시스템을 마운트합니다. 그러나 우선 [젠투 설치 파일을 설치](/wiki/Handbook:PPC64/Installation/Stage/ko "Handbook:PPC64/Installation/Stage/ko") 하겠습니다.

[← 네트워크 설정](/wiki/Handbook:PPC64/Installation/Networking/ko "이전 내용") [처음](/wiki/Handbook:PPC64/ko "Handbook:PPC64/ko") [젠투 설치 파일 설치 →](/wiki/Handbook:PPC64/Installation/Stage/ko "다음 내용")