# Disks

Other languages:

- [Deutsch](/wiki/Handbook:PPC/Installation/Disks/de "Handbuch:PPC/Installation/Festplatten (100% translated)")
- [English](/wiki/Handbook:PPC/Installation/Disks "Handbook:PPC/Installation/Disks (100% translated)")
- [español](/wiki/Handbook:PPC/Installation/Disks/es "Manual de Gentoo: PPC/Instalación/Discos (100% translated)")
- [français](/wiki/Handbook:PPC/Installation/Disks/fr "Handbook:PPC/Installation/Disks/fr (100% translated)")
- [italiano](/wiki/Handbook:PPC/Installation/Disks/it "Handbook:PPC/Installation/Disks/it (100% translated)")
- [magyar](/wiki/Handbook:PPC/Installation/Disks/hu "Handbook:PPC/Installation/Disks/hu (100% translated)")
- [polski](/wiki/Handbook:PPC/Installation/Disks/pl "Handbook:PPC/Installation/Disks (100% translated)")
- [português do Brasil](/wiki/Handbook:PPC/Installation/Disks/pt-br "Handbook:PPC/Installation/Disks/pt-br (50% translated)")
- [русский](/wiki/Handbook:PPC/Installation/Disks/ru "Handbook:PPC/Installation/Disks (100% translated)")
- [தமிழ்](/wiki/Handbook:PPC/Installation/Disks/ta "கையேடு:PPC/நிறுவல்/வட்டுக்கள் (100% translated)")
- [中文（中国大陆）‎](/wiki/Handbook:PPC/Installation/Disks/zh-cn "手册：PPC/安装/准备磁盘 (100% translated)")
- [日本語](/wiki/Handbook:PPC/Installation/Disks/ja "ハンドブック:PPC/インストール/ディスク (100% translated)")
- 한국어

[PPC 핸드북](/wiki/Handbook:PPC/ko "Handbook:PPC/ko")[설치](/wiki/Handbook:PPC/Full/Installation/ko "Handbook:PPC/Full/Installation/ko")[설치 정보](/wiki/Handbook:PPC/Installation/About/ko "Handbook:PPC/Installation/About/ko")[매체 선택](/wiki/Handbook:PPC/Installation/Media/ko "Handbook:PPC/Installation/Media/ko")[네트워크 설정](/wiki/Handbook:PPC/Installation/Networking/ko "Handbook:PPC/Installation/Networking/ko")디스크 준비[스테이지 3 설치](/wiki/Handbook:PPC/Installation/Stage/ko "Handbook:PPC/Installation/Stage/ko")[베이스 시스템 설치](/wiki/Handbook:PPC/Installation/Base/ko "Handbook:PPC/Installation/Base/ko")[커널 설정](/wiki/Handbook:PPC/Installation/Kernel/ko "Handbook:PPC/Installation/Kernel/ko")[시스템 설정](/wiki/Handbook:PPC/Installation/System/ko "Handbook:PPC/Installation/System/ko")[도구 설치](/wiki/Handbook:PPC/Installation/Tools/ko "Handbook:PPC/Installation/Tools/ko")[부트로더 설정](/wiki/Handbook:PPC/Installation/Bootloader/ko "Handbook:PPC/Installation/Bootloader/ko")[마무리](/wiki/Handbook:PPC/Installation/Finalizing/ko "Handbook:PPC/Installation/Finalizing/ko")[젠투 활용](/wiki/Handbook:PPC/Full/Working/ko "Handbook:PPC/Full/Working/ko")[포티지 소개](/wiki/Handbook:PPC/Working/Portage/ko "Handbook:PPC/Working/Portage/ko")[USE 플래그](/wiki/Handbook:PPC/Working/USE/ko "Handbook:PPC/Working/USE/ko")[포티지 기능](/wiki/Handbook:PPC/Working/Features/ko "Handbook:PPC/Working/Features/ko")[초기화 스크립트 시스템](/wiki/Handbook:PPC/Working/Initscripts/ko "Handbook:PPC/Working/Initscripts/ko")[환경 변수](/wiki/Handbook:PPC/Working/EnvVar/ko "Handbook:PPC/Working/EnvVar/ko")[포티지 활용](/wiki/Handbook:PPC/Full/Portage/ko "Handbook:PPC/Full/Portage/ko")[파일 및 디렉터리](/wiki/Handbook:PPC/Portage/Files/ko "Handbook:PPC/Portage/Files/ko")[변수](/wiki/Handbook:PPC/Portage/Variables/ko "Handbook:PPC/Portage/Variables/ko")[소프트웨어 브랜치 함께 사용하기](/wiki/Handbook:PPC/Portage/Branches/ko "Handbook:PPC/Portage/Branches/ko")[추가 도구](/wiki/Handbook:PPC/Portage/Tools/ko "Handbook:PPC/Portage/Tools/ko")[꾸러미 저장소 개별 설정](/wiki/Handbook:PPC/Portage/CustomTree/ko "Handbook:PPC/Portage/CustomTree/ko")[고급 기능](/wiki/Handbook:PPC/Portage/Advanced/ko "Handbook:PPC/Portage/Advanced/ko")[네트워크 설정](/wiki/Handbook:PPC/Full/Networking/ko "Handbook:PPC/Full/Networking/ko")[시작하기](/wiki/Handbook:PPC/Networking/Introduction/ko "Handbook:PPC/Networking/Introduction/ko")[고급 설정](/wiki/Handbook:PPC/Networking/Advanced/ko "Handbook:PPC/Networking/Advanced/ko")[모듈러 네트워크](/wiki/Handbook:PPC/Networking/Modular/ko "Handbook:PPC/Networking/Modular/ko")[무선 네트워크](/wiki/Handbook:PPC/Networking/Wireless/ko "Handbook:PPC/Networking/Wireless/ko")[기능 추가](/wiki/Handbook:PPC/Networking/Extending/ko "Handbook:PPC/Networking/Extending/ko")[동적 관리](/wiki/Handbook:PPC/Networking/Dynamic/ko "Handbook:PPC/Networking/Dynamic/ko")

## Contents

- [1블록 장치 소개](#.EB.B8.94.EB.A1.9D_.EC.9E.A5.EC.B9.98_.EC.86.8C.EA.B0.9C)
  - [1.1블록 장치](#.EB.B8.94.EB.A1.9D_.EC.9E.A5.EC.B9.98)
    - [1.1.1공간 분할](#.EA.B3.B5.EA.B0.84_.EB.B6.84.ED.95.A0)
- [2분할 배치 설계](#.EB.B6.84.ED.95.A0_.EB.B0.B0.EC.B9.98_.EC.84.A4.EA.B3.84)
  - [2.1분할 영역을 얼마나 많이, 크게 할까요?](#.EB.B6.84.ED.95.A0_.EC.98.81.EC.97.AD.EC.9D.84_.EC.96.BC.EB.A7.88.EB.82.98_.EB.A7.8E.EC.9D.B4.2C_.ED.81.AC.EA.B2.8C_.ED.95.A0.EA.B9.8C.EC.9A.94.3F)
  - [2.2스왑 공간이 무엇인가요?](#.EC.8A.A4.EC.99.91_.EA.B3.B5.EA.B0.84.EC.9D.B4_.EB.AC.B4.EC.97.87.EC.9D.B8.EA.B0.80.EC.9A.94.3F)
  - [2.3애플 New World](#.EC.95.A0.ED.94.8C_New_World)
  - [2.4애플 Old World](#.EC.95.A0.ED.94.8C_Old_World)
  - [2.5페가소스](#.ED.8E.98.EA.B0.80.EC.86.8C.EC.8A.A4)
  - [2.6IBM PReP (RS/6000)](#IBM_PReP_.28RS.2F6000.29)
- [3mac-fdisk(Apple) 사용](#mac-fdisk.28Apple.29_.EC.82.AC.EC.9A.A9)
- [4parted 사용(페가소스와 RS/6000)](#parted_.EC.82.AC.EC.9A.A9.28.ED.8E.98.EA.B0.80.EC.86.8C.EC.8A.A4.EC.99.80_RS.2F6000.29)
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

#### 공간 분할

이론적으로는 리눅스 시스템을 전체 디스크에 넣을 수 있지만, 실제론 거의 불가능합니다. 대신 전체 블록 장치를 작게 나누어 더욱 관리하기 쉬운 블록 장치를 만들 수 있습니다. 대부분의 시스템에서는 파티션이라고 부릅니다.

**참고**

설치 과정의 나머지 부분에서는 페가소스 예제 파티션 배치를 활용하겠습니다. 개인의 취향에 맞게 조절하십시오.

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

### 애플 New World

애플 뉴 월드 머신은 거의 설정하기 쉽습니다. 첫번째 파티션은 항상 애플 파티션 맵입니다. 이 파티션은 디스크 배치 정보를 저장합니다. 이 파티션은 제거할 수 없습니다. 다음 파티션은 항상 부트스트랩 파티션이 와야합니다. 이 파티션에는 작은(800k) HFS 파일 시스템이 있으며 Yaboot 부트로더 사본과 설정 파일이 들어갑니다. 이 파티션은 다른 아키텍처의 /boot 파티션과 다릅니다. 부트 파티션 다음에는, 보통의 리눅스 파티션이 들어가며 아래 형태를 따라갑니다. 스왑 파티션은 시스템의 물리 메모리가 부족할 때를 대비한 대용 임시 저장소입니다. 루트 파티션은 젠투를 설치할 파일 시스템입니다. 듀얼 부팅을 수행하려면, yaboot를 먼저 시작하는 부트스트랩 파티션의 위치가 보장되고 나서 OSX 파티션을 어디다가든 둘 수 있습니다.

**참고**

Apple\_Driver63, Apple\_Driver\_ATA, Apple\_FWDriver, Apple\_Driver\_IOKit, Apple\_Patche와 같은 "디스크 드라이버" 파티션이 있을지도 모릅니다. MacOS를 부팅할 때 사용하기 때문에 해당 파티션이 필요없다면 mac-fdisk의 `i` 옵션으로 디스크를 초기화하여 제거할 수 있습니다. 이 명령은 디스크를 완전히 지웁니다! 의심이 간다면, 그대로 두십시오.

**참고**

애플의 디스크 유틸리티로 디스크 공간을 분할했다면, 파티션 사이사이에 애플에서 "나중에 사용할" 용도로 예약해둔 128Mb 공간이 남습니다. 이 공간은 안전하게 제거할 수 있습니다.

분할 영역
크기
파일 시스템
설명
/dev/sda132k
없음
Apple Partition Map
/dev/sda2800k
HFS
Apple Bootstrap
/dev/sda3512 MB
Swap
Linux Swap
/dev/sda4Rest of disk
ext3, ext4, reiserfs, xfs, etc.
Linux Root

### 애플 Old World

애플 올드 월드 머신은 설정이 조금 복잡합니다. 첫번째 파티션은 항상 애플 파티션 맵입니다. 이 파티션에는 디스크의 배치 정보를 저장하며 제거할 수 없습니다. BootX를 사용할 때, 아래의 설정은 MacOS를 다른 디스크에 설정했음을 가정합니다. 이 경우가 아니라면, Apple\_Driver63, Apple\_Driver\_ATA, Apple\_FWDriver, Apple\_Driver\_IOKit, Apple\_Patches, MacOS 설치 와 같은 "애플 디스크 드라이버"용 추가 파티션이 될 수 있습니다. Quik을 사용한다면, 다른 애플 부팅 방식과는 달리, 커널을 담아둘 부트 파티션을 만들어야 합니다. 부트 파티션을 만들고 나면, 아래 형태를 따라 보통의 리눅스 파일 시스템이 올라갑니다. 스왑 파티션은 시스템의 물리 메모리가 부족할 때를 대비한 대용 임시 저장소입니다. 루트 파티션에는 젠투를 설치한 파일 시스템이 들어갑니다.

**참고**

올드월드 머신을 사용한다면, MacOS를 사용할 수 있게 두어야합니다. 여기서 사용하는 파티션 배치는 MacOS를 다른 드라이브에 설치했다고 가정합니다.

Example partition layout for an Old World machine:

Partition
Size
Filesystem
Description
/dev/sda132KiB
None.
Apple Partition Map (APM).
/dev/sda232MiB
ext2
Quik Boot Partition (quik only).
/dev/sda3512MiB
swap
Linux swap (type 0x82).
/dev/sda4Rest of the disk.
ext4, xfs, etc.
Linux root.

### 페가소스

페가소스 파티션 배치는 애플 배치에 비해 조금 간단합니다. 첫번째 파티션은 부팅할 커널이 들어간 부트 파티션이며, 부팅시 메뉴로 뜨는 오픈펌웨어 스크립트가 들어갑니다. 부트 파티션 다음에는 보통의 리눅스 파일 시스템이 들어가는데, 다음의 형태를 따릅니다. 스왑 파티션은 시스템의 물리 메모리가 부족할 때를 대비한 대용 임시 저장소입니다. 루트 파티션은 젠투를 설치할 파일 시스템이 들어가는 파티션입니다.

Example partition layout for Pegasos systems:

Partition
Size
Filesystem
Description
/dev/sda132MiB
affs1 or ext2
Boot partition.
/dev/sda2512MiB
swap
Linux swap (type 0x82).
/dev/sda3Rest of the disk.
ext4, xfs, etc.
Linux root.

### IBM PReP (RS/6000)

IBM PowerPC 레퍼런스 플랫폼(PReP)은 디스크 첫번째 파티션에 작은 PReP 부트 파티션이 필요하며, 그 다음에 스왑, 루트 순으로 배치합니다.

Example partition layout for the IBM PReP:

Partition
Size
Filesystem
Description
/dev/sda1800KiB
None
PReP boot partition (type 0x41).
/dev/sda2512MiB
swap
Linux swap (type 0x82).
/dev/sda3Rest of the disk
ext4, xfs, etc.
Linux root (type 0x83).

**경고**

parted는 HFS+ 등의 파티션의 크기를 조절할 수 있습니다. 불행하게도 HFS+ 저널링 파일 시스템 크기 조절에 문제가 있기 때문에, 최상의 결과를 얻으려면, 크기 조절을 진행하기 전에 Mac OS X의 저널링을 끄십시오. 크기 조절 처리 과정은 위험하므로, 위험을 감수하고 진행하십시오! 크기 조절을 진행하기 전에 항상 모든 데이터를 옮겼는지 확인하십시오!

## mac-fdisk(Apple) 사용

이 시점에서 mac-fdisk로 파티션을 만들겠습니다:

`root #` `mac-fdisk /dev/sda`

애플 디스크 유틸리티에서 리눅스에 쓸 공간을 남겨놓았다면, 우선 새로 설치할 공간을 위해 이전에 만들어둔 파티션을 삭제하십시오. mac-fdisk에서 `d` 명령으로 이 파티션을 삭제하십시오. 삭제할 파티션 번호를 물어봅니다. 보통 뉴월드 머신(Apple\_partition\_map)에서 첫번째 파티션은 삭제할 수 없습니다. 디스크 지우기를 시작하려면 간단하게 `i` 를 눌러 디스크를 초기화하십시오. 디스크 전체를 지우므로 주의하여 사용하십시오.

다음, `b` 키를 사용하여 Apple\_Bootstrap 파티션을 만드십시오. 시작 블록 위치를 물어봅니다. 처음 남은 파티션의 번호를 입력하고 `p` 를 입력하십시오. 이 경우, "2p"입니다.

**참고**

이 파티션은 /boot 파티션이 아닙니다. 모든 파티션을 리눅스가 쓰진 않습니다. 이 파티션에 어떤 파일 시스템도 올라갈 필요가 없으며 마운트조차도 해서는 안됩니다. PPC 사용자는 /boot에 대한 추가 파티션이 필요없습니다.

스왑 파티션을 만들려면 `c` 를 입력하십시오. 그 다음 이 파티션의 블록 시작 위치를 물어봅니다. 이전에 Apple\_Bootstrap 파티션을 만들때 2를 입력했던것처럼 선택할 때 "3p"라고 입력하십시오. 크기를 물어보면 512M(또는 필요한 용량만큼. 최소한 512MB가 필요하지만 보통 용납하는 크기는 물리 메모리의 두배입니다)을 입력하십시오. 이름을 물어보면 "swap"이라고 입력하십시오.

루트 파티션을 만들려면 `c` 를 입력하십시오. 그 다음 루트 파티션의 블록 시작 위치를 선택할 때 "4p"라고 입력하십시오. 크기를 물어보면 "4p"를 다시 입력하십시오. mac-fdisk에서는 이 입력을 "존재하는 전체 공간 사용"이란 뜻으로 해석합니다. 이름을 물어보면 "root"라고 입력하십시오.

끝내려면 `w` 로 파티션 정보를 디스크에 기록하고 `q` 로 mac-fdisk를 끝내십시오.

**참고**

모든 부분이 문제없는지 확인하려면, mac-fdisk -l 을 실행하여 모든 파티션이 제자리에 있는지 확인하십시오. 만들어놓은 모든 파티션이 나타나지 않거나, 바뀐 내용 일부가 보이지 않는다면 mac-fdisk에서 `i` 를 눌러 파티션을 다시 초기화하십시오. 참고로 이 명령은 파티션 맵을 다시 만들기 때문에 모든 파티션을 제거합니다.

## parted 사용(페가소스와 RS/6000)

파티션 편집기 parted는 이제 Mac OS와 Mac OS X에서 사용하는 HFS+ 파티션을 다룰 수 있습니다. 이 도구로 Mac 파티션의 크기를 조절할 수 있으며 리눅스 파티션 공간을 만들 수 있습니다. 어쨌거나, 아래 예제에서는 페가소스 머신 전용 디스크 공간 분할 방법을 설명합니다.

parted를 실행하며 시작해보겠습니다:

`root #` `parted /dev/sda`

드라이브의 공간을 분할하지 않았다면 `mklabel amiga` 를 실행하여 드라이브의 새 디스크레이블을 만드십시오.

`print` 를 입력하여 언제든지 parted에서 현재 파티션 테이블을 표시할 수 있습니다. parted 동작을 멈추려면 `Ctrl+C` 를 누르십시오.

리눅스 다음에 MorphOS를 설치했다면, affs1 파일 시스템을 드라이브 시작 지점에 만드십시오. MorphOS 커널을 저장하는데 32MB 정도 되어야 합니다. 페가소스 I 또는 ext2, ext3 이외의 파일 시스템을 리눅스에서 사용한다면 이 파티션에 리눅스 커널을 올려놓아야 합니다(페가소스 II에서는 ext2/ext3/affs1 파티션에서만 부팅할 수 있습니다). 파티션을 만들려면 `mkpart primary affs1 START END` 명령을 실행하고 START와 END 자리에는 메가바이드 범위 값(예: 0 32, 0MB에서 시작해서 32MB 크기의 파티션으로 만듬)으로 넣으십시오. ext2, ext3 파티션을 대신 만든다면, mkpart 명령의 affs1 대신 ext2 또는 ext3를 넣으십시오.

리눅스용 파티션을 두개 만드십시오. 하나는 루트 파일시스템이고 하나는 스왑 파티션입니다. 각각의 파티션을 만들려면 `mkpart primary START END` 를 실행하시고 START와 END 대신 메가바이트 단위 범위 값을 넣으십시오.

보통 스왑 파티션은 컴퓨터에 장착한 RAM 용량의 두배 크기로 만들 것을 추천합니다만 최소한 512Mb 정도를 추천합니다. 스왑 파티션을 만들려면 `mkpart primary linux-swap START END` 명령을 실행하시고 START와 END자리에는 파티션 영역 값을 입력하십시오.

parted 작업이 끝나면 간단하게 `quit` 를 입력하십시오.

## 파일 시스템 만들기

**경고**

When using an SSD or NVMe drive, it is wise to check for firmware upgrades. Some Intel SSDs in particular (600p and 6000p) require a firmware upgrade for [possible data corruption](https://bugzilla.redhat.com/show_bug.cgi?id=1402533) induced by XFS I/O usage patterns. The problem is at the firmware level and not any fault of the XFS filesystem. The smartctl utility can help check the device model and firmware version.

### 도입부

이제 분할 영역을 만들었고, 파일 시스템을 제 위치에 얹어놓을 차례입니다. 다음 절에서는 리눅스에서 지원하는 다양한 파일 시스템을 설명합니다. 어떤 파일 시스템을 사용할 지 이미 알고 있는 독자라면 [파티션에 파일 시스템 반영하기](#.ED.8C.8C.ED.8B.B0.EC.85.98.EC.97.90_.ED.8C.8C.EC.9D.BC_.EC.8B.9C.EC.8A.A4.ED.85.9C_.EB.B0.98.EC.98.81.ED.95.98.EA.B8.B0) 로 계속 진행할 수 있습니다. 그렇지 않으면 계속 읽어 내려가면서 쓸 수 있는 파일시스템이 어떤 종류가 있는지 알아보십시오.

### 파일 시스템

다양한 파일 시스템이 있습니다. 일부는 ppc 아키텍처에서 안정적입니다 - 중요한 분할 영역을 위해서라면 좀 더 시험적인 분할 영역을 선택하기 전에 파일 시스템과 지원 상태에 대한 내용을 좀 더 읽어보시는 것이 좋겠습니다.

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

Installations which were previously started, but did not finish the installation process can resume the installation from this point in the handbook. Use this link as the permalink: [Resumed installations start here](/wiki/Handbook:PPC/Installation/Disks/ko#Resumed_installations_start_here "Handbook:PPC/Installation/Disks/ko").

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

지침을 따르고 나면 proc 파일 시스템(커널 가상 인터페이스)와 다른 커널 의사 파일 시스템을 마운트합니다. 그러나 우선 [젠투 설치 파일을 설치](/wiki/Handbook:PPC/Installation/Stage/ko "Handbook:PPC/Installation/Stage/ko") 하겠습니다.

[← 네트워크 설정](/wiki/Handbook:PPC/Installation/Networking/ko "이전 내용") [처음](/wiki/Handbook:PPC/ko "Handbook:PPC/ko") [젠투 설치 파일 설치 →](/wiki/Handbook:PPC/Installation/Stage/ko "다음 내용")