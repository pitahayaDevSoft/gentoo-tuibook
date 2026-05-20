# Disks

Other languages:

- [Deutsch](/wiki/Handbook:Alpha/Installation/Disks/de "Handbuch:Alpha/Installation/Festplatten (100% translated)")
- [English](/wiki/Handbook:Alpha/Installation/Disks "Handbook:Alpha/Installation/Disks (100% translated)")
- [Türkçe](/wiki/Handbook:Alpha/Installation/Disks/tr "Handbook:Alpha/Installation/Disks/tr (0% translated)")
- [español](/wiki/Handbook:Alpha/Installation/Disks/es "Manual de Gentoo: Alpha/Instalación/Discos (100% translated)")
- [français](/wiki/Handbook:Alpha/Installation/Disks/fr "Manuel:Alpha/Installation/Disques (100% translated)")
- [italiano](/wiki/Handbook:Alpha/Installation/Disks/it "Handbook:Alpha/Installation/Disks/it (100% translated)")
- [magyar](/wiki/Handbook:Alpha/Installation/Disks/hu "Handbook:Alpha/Installation/Disks/hu (100% translated)")
- [polski](/wiki/Handbook:Alpha/Installation/Disks/pl "Handbook:Alpha/Installation/Disks (100% translated)")
- [português do Brasil](/wiki/Handbook:Alpha/Installation/Disks/pt-br "Manual:Alpha/Instalação/Discos (100% translated)")
- [čeština](/wiki/Handbook:Alpha/Installation/Disks/cs "Handbook:Alpha/Installation/Disks/cs (50% translated)")
- [русский](/wiki/Handbook:Alpha/Installation/Disks/ru "Handbook:Alpha/Installation/Disks (100% translated)")
- [தமிழ்](/wiki/Handbook:Alpha/Installation/Disks/ta "கையேடு:Alpha/நிறுவல்/வட்டுக்கள் (100% translated)")
- [中文（中国大陆）‎](/wiki/Handbook:Alpha/Installation/Disks/zh-cn "手册：Alpha/安装/准备磁盘 (100% translated)")
- [日本語](/wiki/Handbook:Alpha/Installation/Disks/ja "ハンドブック:Alpha/インストール/ディスク (100% translated)")
- 한국어

[Alpha 핸드북](/wiki/Handbook:Alpha/ko "Handbook:Alpha/ko")[설치](/wiki/Handbook:Alpha/Full/Installation/ko "Handbook:Alpha/Full/Installation/ko")[설치 정보](/wiki/Handbook:Alpha/Installation/About/ko "Handbook:Alpha/Installation/About/ko")[매체 선택](/wiki/Handbook:Alpha/Installation/Media/ko "Handbook:Alpha/Installation/Media/ko")[네트워크 설정](/wiki/Handbook:Alpha/Installation/Networking/ko "Handbook:Alpha/Installation/Networking/ko")디스크 준비[스테이지 3 설치](/wiki/Handbook:Alpha/Installation/Stage/ko "Handbook:Alpha/Installation/Stage/ko")[베이스 시스템 설치](/wiki/Handbook:Alpha/Installation/Base/ko "Handbook:Alpha/Installation/Base/ko")[커널 설정](/wiki/Handbook:Alpha/Installation/Kernel/ko "Handbook:Alpha/Installation/Kernel/ko")[시스템 설정](/wiki/Handbook:Alpha/Installation/System/ko "Handbook:Alpha/Installation/System/ko")[도구 설치](/wiki/Handbook:Alpha/Installation/Tools/ko "Handbook:Alpha/Installation/Tools/ko")[부트로더 설정](/wiki/Handbook:Alpha/Installation/Bootloader/ko "Handbook:Alpha/Installation/Bootloader/ko")[마무리](/wiki/Handbook:Alpha/Installation/Finalizing/ko "Handbook:Alpha/Installation/Finalizing/ko")[젠투 활용](/wiki/Handbook:Alpha/Full/Working/ko "Handbook:Alpha/Full/Working/ko")[포티지 소개](/wiki/Handbook:Alpha/Working/Portage/ko "Handbook:Alpha/Working/Portage/ko")[USE 플래그](/wiki/Handbook:Alpha/Working/USE/ko "Handbook:Alpha/Working/USE/ko")[포티지 기능](/wiki/Handbook:Alpha/Working/Features/ko "Handbook:Alpha/Working/Features/ko")[초기화 스크립트 시스템](/wiki/Handbook:Alpha/Working/Initscripts/ko "Handbook:Alpha/Working/Initscripts/ko")[환경 변수](/wiki/Handbook:Alpha/Working/EnvVar/ko "Handbook:Alpha/Working/EnvVar/ko")[포티지 활용](/wiki/Handbook:Alpha/Full/Portage/ko "Handbook:Alpha/Full/Portage/ko")[파일 및 디렉터리](/wiki/Handbook:Alpha/Portage/Files/ko "Handbook:Alpha/Portage/Files/ko")[변수](/wiki/Handbook:Alpha/Portage/Variables/ko "Handbook:Alpha/Portage/Variables/ko")[소프트웨어 브랜치 함께 사용하기](/wiki/Handbook:Alpha/Portage/Branches/ko "Handbook:Alpha/Portage/Branches/ko")[추가 도구](/wiki/Handbook:Alpha/Portage/Tools/ko "Handbook:Alpha/Portage/Tools/ko")[꾸러미 저장소 개별 설정](/wiki/Handbook:Alpha/Portage/CustomTree/ko "Handbook:Alpha/Portage/CustomTree/ko")[고급 기능](/wiki/Handbook:Alpha/Portage/Advanced/ko "Handbook:Alpha/Portage/Advanced/ko")[네트워크 설정](/wiki/Handbook:Alpha/Full/Networking/ko "Handbook:Alpha/Full/Networking/ko")[시작하기](/wiki/Handbook:Alpha/Networking/Introduction/ko "Handbook:Alpha/Networking/Introduction/ko")[고급 설정](/wiki/Handbook:Alpha/Networking/Advanced/ko "Handbook:Alpha/Networking/Advanced/ko")[모듈러 네트워크](/wiki/Handbook:Alpha/Networking/Modular/ko "Handbook:Alpha/Networking/Modular/ko")[무선 네트워크](/wiki/Handbook:Alpha/Networking/Wireless/ko "Handbook:Alpha/Networking/Wireless/ko")[기능 추가](/wiki/Handbook:Alpha/Networking/Extending/ko "Handbook:Alpha/Networking/Extending/ko")[동적 관리](/wiki/Handbook:Alpha/Networking/Dynamic/ko "Handbook:Alpha/Networking/Dynamic/ko")

## Contents

- [1블록 장치 소개](#.EB.B8.94.EB.A1.9D_.EC.9E.A5.EC.B9.98_.EC.86.8C.EA.B0.9C)
  - [1.1블록 장치](#.EB.B8.94.EB.A1.9D_.EC.9E.A5.EC.B9.98)
  - [1.2슬라이스](#.EC.8A.AC.EB.9D.BC.EC.9D.B4.EC.8A.A4)
- [2분할 배치 설계](#.EB.B6.84.ED.95.A0_.EB.B0.B0.EC.B9.98_.EC.84.A4.EA.B3.84)
  - [2.1분할 영역을 얼마나 많이, 크게 할까요?](#.EB.B6.84.ED.95.A0_.EC.98.81.EC.97.AD.EC.9D.84_.EC.96.BC.EB.A7.88.EB.82.98_.EB.A7.8E.EC.9D.B4.2C_.ED.81.AC.EA.B2.8C_.ED.95.A0.EA.B9.8C.EC.9A.94.3F)
  - [2.2스왑 공간이 무엇인가요?](#.EC.8A.A4.EC.99.91_.EA.B3.B5.EA.B0.84.EC.9D.B4_.EB.AC.B4.EC.97.87.EC.9D.B8.EA.B0.80.EC.9A.94.3F)
- [3디스크 공간 분할에 fdisk 사용하기(SRM 전용)](#.EB.94.94.EC.8A.A4.ED.81.AC_.EA.B3.B5.EA.B0.84_.EB.B6.84.ED.95.A0.EC.97.90_fdisk_.EC.82.AC.EC.9A.A9.ED.95.98.EA.B8.B0.28SRM_.EC.A0.84.EC.9A.A9.29)
  - [3.1존재하는 디스크 확인하기](#.EC.A1.B4.EC.9E.AC.ED.95.98.EB.8A.94_.EB.94.94.EC.8A.A4.ED.81.AC_.ED.99.95.EC.9D.B8.ED.95.98.EA.B8.B0)
  - [3.2Creating a BSD disklabel](#Creating_a_BSD_disklabel)
  - [3.3모든 슬라이스 삭제](#.EB.AA.A8.EB.93.A0_.EC.8A.AC.EB.9D.BC.EC.9D.B4.EC.8A.A4_.EC.82.AD.EC.A0.9C)
  - [3.4스왑 슬라이스 만들기](#.EC.8A.A4.EC.99.91_.EC.8A.AC.EB.9D.BC.EC.9D.B4.EC.8A.A4_.EB.A7.8C.EB.93.A4.EA.B8.B0)
  - [3.5Creating the boot slice](#Creating_the_boot_slice)
  - [3.6루트 슬라이스 만들기](#.EB.A3.A8.ED.8A.B8_.EC.8A.AC.EB.9D.BC.EC.9D.B4.EC.8A.A4_.EB.A7.8C.EB.93.A4.EA.B8.B0)
  - [3.7슬라이스 배치 저장하고 나가기](#.EC.8A.AC.EB.9D.BC.EC.9D.B4.EC.8A.A4_.EB.B0.B0.EC.B9.98_.EC.A0.80.EC.9E.A5.ED.95.98.EA.B3.A0_.EB.82.98.EA.B0.80.EA.B8.B0)
- [4디스크 분할에 fdisk 사용하기(ARC/AlphaBIOS 전용)](#.EB.94.94.EC.8A.A4.ED.81.AC_.EB.B6.84.ED.95.A0.EC.97.90_fdisk_.EC.82.AC.EC.9A.A9.ED.95.98.EA.B8.B0.28ARC.2FAlphaBIOS_.EC.A0.84.EC.9A.A9.29)
  - [4.1존재하는 디스크 확인하기](#.EC.A1.B4.EC.9E.AC.ED.95.98.EB.8A.94_.EB.94.94.EC.8A.A4.ED.81.AC_.ED.99.95.EC.9D.B8.ED.95.98.EA.B8.B0_2)
  - [4.2모든 파티션 삭제](#.EB.AA.A8.EB.93.A0_.ED.8C.8C.ED.8B.B0.EC.85.98_.EC.82.AD.EC.A0.9C)
  - [4.3부트 파티션 만들기](#.EB.B6.80.ED.8A.B8_.ED.8C.8C.ED.8B.B0.EC.85.98_.EB.A7.8C.EB.93.A4.EA.B8.B0)
  - [4.4스왑 파티션 만들기](#.EC.8A.A4.EC.99.91_.ED.8C.8C.ED.8B.B0.EC.85.98_.EB.A7.8C.EB.93.A4.EA.B8.B0)
  - [4.5루트 파티션 만들기](#.EB.A3.A8.ED.8A.B8_.ED.8C.8C.ED.8B.B0.EC.85.98_.EB.A7.8C.EB.93.A4.EA.B8.B0)
  - [4.6파티션 배치를 저장하고 나가기](#.ED.8C.8C.ED.8B.B0.EC.85.98_.EB.B0.B0.EC.B9.98.EB.A5.BC_.EC.A0.80.EC.9E.A5.ED.95.98.EA.B3.A0_.EB.82.98.EA.B0.80.EA.B8.B0)
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

### 슬라이스

이론적으로는 리눅스 시스템을 전체 디스크에 넣을 수 있지만, 실제론 거의 불가능합니다. 대신 전체 블록 장치를 작게 나누어 더욱 관리하기 쉬운 블록 장치를 만들 수 있습니다. 알파 시스템에서는 _슬라이스_ 라고 부릅니다.

**참고**

다음 절에서는 ARC/AlphaBIOS 설정 용도의 예제 공간 분할 절차를 설치 절차로 활용합니다. 개인 취향에 따라 적절히 바꾸십시오!

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

## 디스크 공간 분할에 fdisk 사용하기(SRM 전용)

다음 부분은 SRM에서 사용하는 슬라이스 배치 예제를 만드는 방법을 설명합니다:

슬라이스
설명
/dev/sda1스왑 슬라이스
/dev/sda2루트 슬라이스
/dev/sda3전체 디스크(필요)

개인 취향에 따라 슬라이스 배치를 바꾸십시오.

### 존재하는 디스크 확인하기

어떤 디스크가 동작중인지 확인하려면, 다음 명령을 사용하십시오:

IDE 디스크에서는:

`root #` `dmesg | grep 'drive$'`

SCSI 디스크에서는:

`root #` `dmesg | grep 'scsi'`

이 출력 결과를 통해 어떤 디스크를 감지했으며 각각의 /dev/ 항목이 무엇인지를 보여줍니다. 다음 부분에서는 /dev/sda에 SCSI 디스크를 사용한다고 가정하겠습니다.

### Creating a BSD disklabel

If the hard drive is completely blank, then first create a BSD disklabel. On alpha, this can't be done using fdisk, so using parted is the way forward here.

Now fire up parted:

`root #` `parted /dev/sda `

```
Using /dev/sda
Welcome to GNU Parted! Type 'help' to view a list of commands.

```

`(parted)` `mklabel bsd`

```
Warning: The existing disk label on /dev/sda will be destroyed and all data on this disk will be lost. Do you want to continue?
Yes/No? yes
(parted) quit
Information: You may need to update /etc/fstab.

```

Now that we have a bsd disklabel on our drive, continue creating slices. This can be done using parted or, as in the examples below, using fdisk:

`root #` `fdisk /dev/sda`

### 모든 슬라이스 삭제

하드 드라이브를 완전히 비워두었다면, 먼저 BSD 디스크 레이블을 만드십시오.

`Command (m for help):` `b`

```
/dev/sda contains no disklabel.
Do you want to create a disklabel? (y/n) y
A bunch of drive-specific info will show here
3 partitions:
#       start       end      size     fstype   [fsize bsize   cpg]
  c:        1      5290*     5289*    unused        0     0

```

'c' 슬라이스를 제외한 나머지 모든 슬라이스를 삭제하겠습니다.(BSD 디스크 레이블을 사용하려면 필요합니다). 다음 과정은 슬라이스를 삭제하는 방법을 보여줍니다(예제에서는 'a'를 사용하겠습니다). 다시 말하지만 모든 다른 슬라이스를 삭제하는 과정입니다(자꾸 반복해서 말하지만 'c' 슬라이스는 제외입니다).

현존하는 파티션을 보려면 `p` 를 사용하십시오. `d` 는 파티션을 지울 때 사용합니다.

`BSD disklabel command (m for help):` `p`

```
8 partitions:
#       start       end      size     fstype   [fsize bsize   cpg]
  a:        1       235*      234*    4.2BSD     1024  8192    16
  b:      235*      469*      234*      swap
  c:        1      5290*     5289*    unused        0     0
  d:      469*     2076*     1607*    unused        0     0
  e:     2076*     3683*     1607*    unused        0     0
  f:     3683*     5290*     1607*    unused        0     0
  g:      469*     1749*     1280     4.2BSD     1024  8192    16
  h:     1749*     5290*     3541*    unused        0     0

```

`BSD disklabel command (m for help):` `d`

```
Partition (a-h): a

```

모든 슬라이스에 대해 이 단계를 반복한 후에는 다음과 같은 결과가 나타납니다:

`BSD disklabel command (m for help):` `p`

```
3 partitions:
#       start       end      size     fstype   [fsize bsize   cpg]
  c:        1      5290*     5289*    unused        0     0

```

### 스왑 슬라이스 만들기

알파 기반 시스템에서는 부팅 슬라이스를 쪼갤 필요가 없습니다. 그러나 첫번째 실린더는 _aboot_ 이미지를 넣을 용도로 활용할 수 없습니다.

세번째 실린더에서 시작하는 총 용량 1GB의 스왑 슬라이스를 만들겠습니다. `n` 으로 새 슬라이스를 만드십시오. 슬라이스를 만든 후, 스왑을 의미하는 `1`(일)을 입력하여 형식을 바꾸겠습니다.

`BSD disklabel command (m for help):` `n`

```
Partition (a-p): a
First cylinder (1-5290, default 1): 3
Last cylinder or +size or +sizeM or +sizeK (3-5290, default 5290): +1024M

```

`BSD disklabel command (m for help):` `t`

```
Partition (a-c): a
Hex code (type L to list codes): 1

```

이 단계가 끝나면 배치는 다음에 나타난 모습과 비슷합니다:

`BSD disklabel command (m for help):` `p`

```
3 partitions:
#       start       end      size     fstype   [fsize bsize   cpg]
  a:        3      1003      1001       swap
  c:        1      5290*     5289*    unused        0     0

```

### Creating the boot slice

Create a boot file system containing the kernels and boot loader configuration file (aboot.conf). aboot supports ext2 and ext3 filesystems only. Create the boot slice starting from the first cylinder after the swap slice. Use the `p` command to view where the swap slice ends. In our example, this is at 1003, making the boot slice start at 1004.

`BSD disklabel command (m for help):` `n`

```
Partition (a-p): b
First cylinder (1-5290, default 1): 1004
Last cylinder or +size or +sizeM or +sizeK (3-5290, default 5290): +1024M

```

`BSD disklabel command (m for help):` `t`

```
Partition (a-c): b
Hex code (type L to list codes): 08

```

After these steps a layout similar to the following should be shown:

`BSD disklabel command (m for help):` `p`

```
3 partitions:
#       start       end      size     fstype   [fsize bsize   cpg]
  a:        3      1003      1001       swap
  b:     1004      2005      1001       ext2
  c:        1      5290*     5289*    unused        0     0

```

### 루트 슬라이스 만들기

이제 스왑 슬라이드 다음의 첫번째 실린더에서 시작하는 루트 슬라이스를 만들겠습니다. `p` 명령을 사용하여 스왑 슬라이스 끝이 어디인지 살펴보십시오. 예제에서는 1003으로 끝나며, 앞으로 만들 루트 슬라이스는 1004에서 시작합니다.

다른 문제를 언급하자면, `fdisk` 에서, 존재하는 실린더의 갯수가 실제 실린더의 갯수보다 하나가 더 많은 것으로 간주하는 버그가 아직도 있습니다. 달리 말해 마지막 실린더를 요청하면 실린더 숫자를 1 감소(이 경우 5290)하여 보여줍니다.

슬라이스를 만들었다면, ext2 형식을 의미하는 형식 8번으로 바꾸겠습니다.

`BSD disklabel command (m for help):` `n`

```
Partition (a-p): b
First cylinder (1-5290, default 1): 1004
Last cylinder or +size or +sizeM or +sizeK (1004-5290, default 5290): 5289

```

`BSD disklabel command (m for help):` `t`

```
Partition (a-c): b
Hex code (type L to list codes): 8

```

결과 슬라이스 레이아웃은 다음과 같아야합니다:

`BSD disklabel command (m for help):` `p`

```
3 partitions:
#       start       end      size     fstype   [fsize bsize   cpg]
  a:        3      1003      1001       swap
  b:     1004      5289      4286       ext2
  c:        1      5290*     5289*    unused        0     0

```

### 슬라이스 배치 저장하고 나가기

`w` 를 입력하여 `fdisk` 프로그램을 빠져나가십시오. 이 명령은 슬라이스 배치를 저장하기도 합니다.

`Command (m for help):` `w`

## 디스크 분할에 fdisk 사용하기(ARC/AlphaBIOS 전용)

다음 부분은 ARC/AlphaBIOS에서 예제 슬라이스 배치를 만드는 방법을 설명합니다:

슬라이스
설명
/dev/sda1부트 파티션
/dev/sda2스왑 파티션
/dev/sda3루트 파티션

개인 취향에 따라 파티션 배치를 바꾸십시오.

### 존재하는 디스크 확인하기

어떤 디스크가 동작중인지 확인하려면, 다음 명령을 활용하십시오:

IDE 디스크에서는:

`root #` `dmesg | grep 'drive$'`

SCSI 디스크에서는:

`root #` `dmesg | grep 'scsi'`

이 출력 결과를 통해 어떤 디스크를 감지했으며 각각의 /dev/ 항목이 무엇인지 쉽게 찾아볼 수 있어야 합니다. 다음에서는 /dev/sda의 SCSI 디스크를 사용함을 가정하겠습니다.

이제 fdisk를 실행하십시오:

`root #` `fdisk /dev/sda`

### 모든 파티션 삭제

하드 드라이브를 완전히 비워두었다면, 먼저 DOS 디스크 레이블을 만드십시오.

`Command (m for help):` `o`

```
Building a new DOS disklabel.

```

모든 파티션에 대한 삭제를 시작하겠습니다. 다음은 파티션을 삭제하는 방법을 보여줍니다(예제에서는 '1'을 사용합니다). 모든 파티션을 삭제하려면 과정을 반복하십시오.

현존하는 파티션을 보려면 `p` 를 사용하십시오. `d` 는 파티션을 지울 때 사용합니다.

`command (m for help):` `p`

```
Disk /dev/sda: 9150 MB, 9150996480 bytes
64 heads, 32 sectors/track, 8727 cylinders
Units = cylinders of 2048 * 512 = 1048576 bytes

   Device Boot      Start         End      Blocks   Id  System
/dev/sda1               1         478      489456   83  Linux
/dev/sda2             479        8727     8446976    5  Extended
/dev/sda5             479        1433      977904   83  Linux Swap
/dev/sda6            1434        8727     7469040   83  Linux

```

`command (m for help):` `d`

```
Partition number (1-6): 1

```

### 부트 파티션 만들기

부팅시 MILO를 사용하는 알파 시스템에서는, 작은 vfat 부트 파티션을 만들어야 합니다.

`Command (m for help):` `n`

```
Command action
  e   extended
  p   primary partition (1-4)
p
Partition number (1-4): 1
First cylinder (1-8727, default 1): 1
Last cylinder or +size or +sizeM or +sizeK (1-8727, default 8727): +16M

```

`Command (m for help):` `t`

```
Selected partition 1
Hex code (type L to list codes): 6
Changed system type of partition 1 to 6 (FAT16)

```

### 스왑 파티션 만들기

총 용량 1GB의 스왑 파티션을 만들겠습니다. `n` 을 사용하여 새 파티션을 만드십시오.

`Command (m for help):` `n`

```
Command action
  e   extended
  p   primary partition (1-4)
p
Partition number (1-4): 2
First cylinder (17-8727, default 17): 17
Last cylinder or +size or +sizeM or +sizeK (17-8727, default 8727): +1000M

```

`Command (m for help):` `t`

```
Partition number (1-4): 2
Hex code (type L to list codes): 82
Changed system type of partition 2 to 82 (Linux swap)

```

이 단계가 끝나면 배치는 다음과 같습니다:

`Command (m for help):` `p`

```
Disk /dev/sda: 9150 MB, 9150996480 bytes
64 heads, 32 sectors/track, 8727 cylinders
Units = cylinders of 2048 * 512 = 1048576 bytes

   Device Boot      Start         End      Blocks   Id  System
/dev/sda1               1          16       16368    6  FAT16
/dev/sda2              17         971      977920   82  Linux swap

```

### 루트 파티션 만들기

이제 루트 파티션을 만들겠습니다. 다시 `n` 명령만을 사용하십시오.

`Command (m for help):` `n`

```
Command action
  e   extended
  p   primary partition (1-4)
p
Partition number (1-4): 3
First cylinder (972-8727, default 972): 972
Last cylinder or +size or +sizeM or +sizeK (972-8727, default 8727): 8727

```

이 단계가 끝나면 배치는 다음과 같이 나타나야합니다:

`Command (m for help):` `p`

```
Disk /dev/sda: 9150 MB, 9150996480 bytes
64 heads, 32 sectors/track, 8727 cylinders
Units = cylinders of 2048 * 512 = 1048576 bytes

   Device Boot      Start         End      Blocks   Id  System
/dev/sda1               1          16       16368    6  FAT16
/dev/sda2              17         971      977920   82  Linux swap
/dev/sda3             972        8727     7942144   83  Linux

```

### 파티션 배치를 저장하고 나가기

`fdisk` 에서 바꾼 내용을 `w` 를 입력하여 저장하십시오.

`Command (m for help):` `w`

이제 파티션을 만들었습니다. [파일 시스템 만들기](#.ED.8C.8C.EC.9D.BC_.EC.8B.9C.EC.8A.A4.ED.85.9C_.EB.A7.8C.EB.93.A4.EA.B8.B0) 로 계속 진행하십시오.

## 파일 시스템 만들기

**경고**

When using an SSD or NVMe drive, it is wise to check for firmware upgrades. Some Intel SSDs in particular (600p and 6000p) require a firmware upgrade for [possible data corruption](https://bugzilla.redhat.com/show_bug.cgi?id=1402533) induced by XFS I/O usage patterns. The problem is at the firmware level and not any fault of the XFS filesystem. The smartctl utility can help check the device model and firmware version.

### 도입부

이제 분할 영역을 만들었고, 파일 시스템을 제 위치에 얹어놓을 차례입니다. 다음 절에서는 리눅스에서 지원하는 다양한 파일 시스템을 설명합니다. 어떤 파일 시스템을 사용할 지 이미 알고 있는 독자라면 [파티션에 파일 시스템 반영하기](#.ED.8C.8C.ED.8B.B0.EC.85.98.EC.97.90_.ED.8C.8C.EC.9D.BC_.EC.8B.9C.EC.8A.A4.ED.85.9C_.EB.B0.98.EC.98.81.ED.95.98.EA.B8.B0) 로 계속 진행할 수 있습니다. 그렇지 않으면 계속 읽어 내려가면서 쓸 수 있는 파일시스템이 어떤 종류가 있는지 알아보십시오.

### 파일 시스템

다양한 파일 시스템이 있습니다. 일부는 alpha 아키텍처에서 안정적입니다 - 중요한 분할 영역을 위해서라면 좀 더 시험적인 분할 영역을 선택하기 전에 파일 시스템과 지원 상태에 대한 내용을 좀 더 읽어보시는 것이 좋겠습니다.

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

Installations which were previously started, but did not finish the installation process can resume the installation from this point in the handbook. Use this link as the permalink: [Resumed installations start here](/wiki/Handbook:Alpha/Installation/Disks/ko#Resumed_installations_start_here "Handbook:Alpha/Installation/Disks/ko").

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

지침을 따르고 나면 proc 파일 시스템(커널 가상 인터페이스)와 다른 커널 의사 파일 시스템을 마운트합니다. 그러나 우선 [젠투 설치 파일을 설치](/wiki/Handbook:Alpha/Installation/Stage/ko "Handbook:Alpha/Installation/Stage/ko") 하겠습니다.

[← 네트워크 설정](/wiki/Handbook:Alpha/Installation/Networking/ko "이전 내용") [처음](/wiki/Handbook:Alpha/ko "Handbook:Alpha/ko") [젠투 설치 파일 설치 →](/wiki/Handbook:Alpha/Installation/Stage/ko "다음 내용")