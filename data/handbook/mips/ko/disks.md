# Disks

Other languages:

- [Deutsch](/wiki/Handbook:MIPS/Installation/Disks/de "Handbuch:MIPS/Installation/Festplatten (100% translated)")
- [English](/wiki/Handbook:MIPS/Installation/Disks "Handbook:MIPS/Installation/Disks (100% translated)")
- [español](/wiki/Handbook:MIPS/Installation/Disks/es "Manual de Gentoo: MIPS/Instalación/Discos (100% translated)")
- [français](/wiki/Handbook:MIPS/Installation/Disks/fr "Handbook:MIPS/Installation/Disks/fr (100% translated)")
- [italiano](/wiki/Handbook:MIPS/Installation/Disks/it "Handbook:MIPS/Installation/Disks/it (100% translated)")
- [magyar](/wiki/Handbook:MIPS/Installation/Disks/hu "Handbook:MIPS/Installation/Disks/hu (100% translated)")
- [polski](/wiki/Handbook:MIPS/Installation/Disks/pl "Handbook:MIPS/Installation/Disks (100% translated)")
- [português do Brasil](/wiki/Handbook:MIPS/Installation/Disks/pt-br "Manual:MIPS/Instalação/Discos (100% translated)")
- [русский](/wiki/Handbook:MIPS/Installation/Disks/ru "Handbook:MIPS/Installation/Disks (100% translated)")
- [தமிழ்](/wiki/Handbook:MIPS/Installation/Disks/ta "கையேடு:MIPS/நிறுவல்/வட்டுக்கள் (100% translated)")
- [中文（中国大陆）‎](/wiki/Handbook:MIPS/Installation/Disks/zh-cn "手册：MIPS/安装/准备磁盘 (100% translated)")
- [日本語](/wiki/Handbook:MIPS/Installation/Disks/ja "ハンドブック:MIPS/インストール/ディスク (100% translated)")
- 한국어

[MIPS 핸드북](/wiki/Handbook:MIPS/ko "Handbook:MIPS/ko")[설치](/wiki/Handbook:MIPS/Full/Installation/ko "Handbook:MIPS/Full/Installation/ko")[설치 정보](/wiki/Handbook:MIPS/Installation/About/ko "Handbook:MIPS/Installation/About/ko")[매체 선택](/wiki/Handbook:MIPS/Installation/Media/ko "Handbook:MIPS/Installation/Media/ko")[네트워크 설정](/wiki/Handbook:MIPS/Installation/Networking/ko "Handbook:MIPS/Installation/Networking/ko")디스크 준비[스테이지 3 설치](/wiki/Handbook:MIPS/Installation/Stage/ko "Handbook:MIPS/Installation/Stage/ko")[베이스 시스템 설치](/wiki/Handbook:MIPS/Installation/Base/ko "Handbook:MIPS/Installation/Base/ko")[커널 설정](/wiki/Handbook:MIPS/Installation/Kernel/ko "Handbook:MIPS/Installation/Kernel/ko")[시스템 설정](/wiki/Handbook:MIPS/Installation/System/ko "Handbook:MIPS/Installation/System/ko")[도구 설치](/wiki/Handbook:MIPS/Installation/Tools/ko "Handbook:MIPS/Installation/Tools/ko")[부트로더 설정](/wiki/Handbook:MIPS/Installation/Bootloader/ko "Handbook:MIPS/Installation/Bootloader/ko")[마무리](/wiki/Handbook:MIPS/Installation/Finalizing/ko "Handbook:MIPS/Installation/Finalizing/ko")[젠투 활용](/wiki/Handbook:MIPS/Full/Working/ko "Handbook:MIPS/Full/Working/ko")[포티지 소개](/wiki/Handbook:MIPS/Working/Portage/ko "Handbook:MIPS/Working/Portage/ko")[USE 플래그](/wiki/Handbook:MIPS/Working/USE/ko "Handbook:MIPS/Working/USE/ko")[포티지 기능](/wiki/Handbook:MIPS/Working/Features/ko "Handbook:MIPS/Working/Features/ko")[초기화 스크립트 시스템](/wiki/Handbook:MIPS/Working/Initscripts/ko "Handbook:MIPS/Working/Initscripts/ko")[환경 변수](/wiki/Handbook:MIPS/Working/EnvVar/ko "Handbook:MIPS/Working/EnvVar/ko")[포티지 활용](/wiki/Handbook:MIPS/Full/Portage/ko "Handbook:MIPS/Full/Portage/ko")[파일 및 디렉터리](/wiki/Handbook:MIPS/Portage/Files/ko "Handbook:MIPS/Portage/Files/ko")[변수](/wiki/Handbook:MIPS/Portage/Variables/ko "Handbook:MIPS/Portage/Variables/ko")[소프트웨어 브랜치 함께 사용하기](/wiki/Handbook:MIPS/Portage/Branches/ko "Handbook:MIPS/Portage/Branches/ko")[추가 도구](/wiki/Handbook:MIPS/Portage/Tools/ko "Handbook:MIPS/Portage/Tools/ko")[꾸러미 저장소 개별 설정](/wiki/Handbook:MIPS/Portage/CustomTree/ko "Handbook:MIPS/Portage/CustomTree/ko")[고급 기능](/wiki/Handbook:MIPS/Portage/Advanced/ko "Handbook:MIPS/Portage/Advanced/ko")[네트워크 설정](/wiki/Handbook:MIPS/Full/Networking/ko "Handbook:MIPS/Full/Networking/ko")[시작하기](/wiki/Handbook:MIPS/Networking/Introduction/ko "Handbook:MIPS/Networking/Introduction/ko")[고급 설정](/wiki/Handbook:MIPS/Networking/Advanced/ko "Handbook:MIPS/Networking/Advanced/ko")[모듈러 네트워크](/wiki/Handbook:MIPS/Networking/Modular/ko "Handbook:MIPS/Networking/Modular/ko")[무선 네트워크](/wiki/Handbook:MIPS/Networking/Wireless/ko "Handbook:MIPS/Networking/Wireless/ko")[기능 추가](/wiki/Handbook:MIPS/Networking/Extending/ko "Handbook:MIPS/Networking/Extending/ko")[동적 관리](/wiki/Handbook:MIPS/Networking/Dynamic/ko "Handbook:MIPS/Networking/Dynamic/ko")

## Contents

- [1블록 장치 소개](#.EB.B8.94.EB.A1.9D_.EC.9E.A5.EC.B9.98_.EC.86.8C.EA.B0.9C)
  - [1.1블록 장치](#.EB.B8.94.EB.A1.9D_.EC.9E.A5.EC.B9.98)
  - [1.2공간 분할](#.EA.B3.B5.EA.B0.84_.EB.B6.84.ED.95.A0)
- [2분할 배치 설계](#.EB.B6.84.ED.95.A0_.EB.B0.B0.EC.B9.98_.EC.84.A4.EA.B3.84)
  - [2.1분할 영역을 얼마나 많이, 크게 할까요?](#.EB.B6.84.ED.95.A0_.EC.98.81.EC.97.AD.EC.9D.84_.EC.96.BC.EB.A7.88.EB.82.98_.EB.A7.8E.EC.9D.B4.2C_.ED.81.AC.EA.B2.8C_.ED.95.A0.EA.B9.8C.EC.9A.94.3F)
  - [2.2스왑 공간이 무엇인가요?](#.EC.8A.A4.EC.99.91_.EA.B3.B5.EA.B0.84.EC.9D.B4_.EB.AC.B4.EC.97.87.EC.9D.B8.EA.B0.80.EC.9A.94.3F)
- [3fdisk 사용](#fdisk_.EC.82.AC.EC.9A.A9)
  - [3.1SGI 머신: SGI 디스크 레이블 만들기](#SGI_.EB.A8.B8.EC.8B.A0:_SGI_.EB.94.94.EC.8A.A4.ED.81.AC_.EB.A0.88.EC.9D.B4.EB.B8.94_.EB.A7.8C.EB.93.A4.EA.B8.B0)
  - [3.2SGI 볼륨 헤더 크기 조절](#SGI_.EB.B3.BC.EB.A5.A8_.ED.97.A4.EB.8D.94_.ED.81.AC.EA.B8.B0_.EC.A1.B0.EC.A0.88)
  - [3.3Cobalt 드라이브 공간 분할](#Cobalt_.EB.93.9C.EB.9D.BC.EC.9D.B4.EB.B8.8C_.EA.B3.B5.EA.B0.84_.EB.B6.84.ED.95.A0)
- [4파일 시스템 만들기](#.ED.8C.8C.EC.9D.BC_.EC.8B.9C.EC.8A.A4.ED.85.9C_.EB.A7.8C.EB.93.A4.EA.B8.B0)
  - [4.1도입부](#.EB.8F.84.EC.9E.85.EB.B6.80)
  - [4.2파일 시스템](#.ED.8C.8C.EC.9D.BC_.EC.8B.9C.EC.8A.A4.ED.85.9C)
  - [4.3분할 영역에 파일 시스템 반영하기](#.EB.B6.84.ED.95.A0_.EC.98.81.EC.97.AD.EC.97.90_.ED.8C.8C.EC.9D.BC_.EC.8B.9C.EC.8A.A4.ED.85.9C_.EB.B0.98.EC.98.81.ED.95.98.EA.B8.B0)
    - [4.3.1EFI system partition filesystem](#EFI_system_partition_filesystem)
    - [4.3.2Legacy BIOS boot partition filesystem](#Legacy_BIOS_boot_partition_filesystem)
    - [4.3.3Small ext4 partitions](#Small_ext4_partitions)
  - [4.4스왑 분할 영역 활성화](#.EC.8A.A4.EC.99.91_.EB.B6.84.ED.95.A0_.EC.98.81.EC.97.AD_.ED.99.9C.EC.84.B1.ED.99.94)
- [5루트 분할 영역 마운트](#.EB.A3.A8.ED.8A.B8_.EB.B6.84.ED.95.A0_.EC.98.81.EC.97.AD_.EB.A7.88.EC.9A.B4.ED.8A.B8)

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

### 공간 분할

이론적으로는 리눅스 시스템을 전체 디스크에 넣을 수 있지만, 실제론 거의 불가능합니다. 대신 전체 블록 장치를 작게 나누어 더욱 관리하기 쉬운 블록 장치를 만들 수 있습니다. 이를 분할 영역(파티션)이라고 부릅니다.

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

## fdisk 사용

### SGI 머신: SGI 디스크 레이블 만들기

SGI 시스템의 모든 디스크는 SGI 디스크 레이블이 필요한데 썬과 MS-DOS의 디스크 레이블의 기능과 비슷합니다. 여기에는 디스크 분할 영역 정보를 저장합니다. 새 SGI 디스크 레이블을 만들면 디스크에 다음 두가지 특별한 분할 영역을 만듭니다.

- SGI 볼륨 헤더(9번째): 이 분할 영역은 중요합니다. 부트로더가 위치할 곳이며 어떤 경우는 커널 이미지도 들어갑니다.
- SGI 볼륨(11번째): 이 분할 영역은 썬 디스크 레이블의 세번째 "전체 디스크" 분할 영역과 유사한 목적으로 씁니다. 이 분할 영역은 전체 디스크를 차지하며, 건드려선 안됩니다. 문서상에 남기지 않은 방식(또는 어떤 방식에 있어서는 IRIX에서 사용)으로 PROM의 동작을 보조하는 다른 분할 영역과는 달리 특별한 목적으로 쓰지 않습니다.

**경고**

SGI 볼륨 헤더는 0번 실린더부터 시작해야합니다. 이게 실패하면 디스크 부팅이 실패한다는 의미입니다.

다음은 fdisk 세션에서 발췌한 예제입니다. 개인 취향에 맞춰 읽어내려가면서 설정을 적당히 바꾸십시오...

`root #` `fdisk /dev/sda`

전문가 모드로 전환하십시오:

`Command (m for help):` `x`

`m` 을 눌러 전체 메뉴 옵션을 표시하십시오:

`Expert command (m for help):` `m`

```
Command action
   b   move beginning of data in a partition
   c   change number of cylinders
   d   print the raw data in the partition table
   e   list extended partitions
   f   fix partition order
   g   create an IRIX (SGI) partition table
   h   change number of heads
   m   print this menu
   p   print the partition table
   q   quit without saving changes
   r   return to main menu
   s   change number of sectors/track
   v   verify the partition table
   w   write table to disk and exit

```

SGI 디스크 레이블을 만드십시오:

`Expert command (m for help):` `g`

```
Building a new SGI disklabel. Changes will remain in memory only,
until you decide to write them. After that, of course, the previous
content will be irrecoverably lost.

```

메인 메뉴로 돌아가십시오:

`Expert command (m for help):` `r`

현재 분할 영역 배치를 살펴보십시오:

`Command (m for help):` `p`

```
Disk /dev/sda (SGI disk label): 64 heads, 32 sectors, 17482 cylinders
Units = cylinders of 2048 * 512 bytes

----- partitions -----
Pt#     Device  Info     Start       End   Sectors  Id  System
 9:  /dev/sda1               0         4     10240   0  SGI volhdr
11:  /dev/sda2               0     17481  35803136   6  SGI volume
----- Bootinfo -----
Bootfile: /unix
----- Directory Entries -----

```

**참고**

디스크에 기존의 SGI 디스크 레이블이 남아있다면 fdisk에서 새 레이블을 만들 수 없습니다. 이런 문제를 회피할 수 있는 방법이 두가지가 있는데 하나는 썬 또는 MS-DOS 디스크 레이블을 만들고, 디스크에 바뀐 내용을 기록한 후, fdisk를 다시 시작하는 방법입니다. 다른 방법은 다음 명령으로 분할 영역 테이블에 널 데이터를 엎어쓰는 방법입니다: `dd if=/dev/zero of=/dev/sda bs=512 count=1`

### SGI 볼륨 헤더 크기 조절

**중요**

이 단계는 fdisk 버그 때문에 자주 필요합니다. 이런 문제 때문에 볼륨 헤더를 제대로 만들지 못하며, 시작/끝 실린더 값이 0이 됩니다. 이 동작은 다중 분할 영역을 만들지 못하게 합니다. 이 문제를 벗어나려면 ... 계속 읽어내려가십시오.

이제 SGI 디스크 레이블을 만들었고 분할 영역을 지정할 차례입니다. 위 예제에서 이미 두 분할 영역을 지정했습니다. 이 분할 영역은 위에서 언급한 대로 특별하며 일반 분할 영역으로 바꿀 수 없습니다. 그러나 젠투를 설치할 때는 부트로더를 불러와야 하고 가능하다면 여러가지 커널 이미지(시스템 형식에 따름)를 볼륨 헤더에 바로 넣어야합니다. 볼륨 헤더 자체는 어떤 크기로든 8개 문자 이름까지만 허용하는 여덟 개의 이미지를 가질 수 있습니다.

볼륨 헤더를 더 크게 만드는 과정은 쉽지 않습니다. 약간 요령을 피워야합니다. 간단하게 지울 수 없으며 fdisk 기능이 모자란 상황이라 볼륨 헤더를 다시 추가할 수 없습니다. 아래 보여드릴 예제에서 50MB /boot 분할 영역과 결합한 50MB 볼륨 헤더를 만들겠습니다. 실제 디스크의 배치는 다양하지만, 여기선 예제 목적입니다.

새 분할 영역을 만드십시오:

`Command (m for help):` `n`

```
Partition number (1-16): 1
First cylinder (5-8682, default 5): 51
 Last cylinder (51-8682, default 8682): 101

```

어떻게 fdisk에서 1번 분할 영역을 최소한 실린더 5로 시작하는 분할 영역으로 다시 만들었을까요? SGI 볼륨 헤더를 이 방식으로 삭제하고 다시 만들면 우리가 인식하는 동일한 문제가 반복됩니다. 이 예제에서는 /boot/를 50MB로 설정하려 했으므로 실린더 51(볼륨 헤더는 실린더 0으로 시작한다고 했죠? 아마?)에서 /boot/를 시작하고, 마지막 실린더를 101로 설정하여 거의 50MB(1-5MB 내외 차이)가 되도록 설정하겠습니다.

분할 영역을 삭제하십시오:

`Command (m for help):` `d`

```
Partition number (1-16): 9

```

이제 다시 만드십시오:

`Command (m for help):` `n`

```
Partition number (1-16): 9
First cylinder (0-50, default 0): 0
 Last cylinder (0-50, default 50): 50

```

fdisk를 쓰는 방법을 모르겠다면 Cobalt에서 하드 디스크 공간 분할하는 절차를 계속 읽어내려가십시오. 개념은 완전히 동일합니다. 단지 볼륨 헤더와 전체 디스크 분할 영역을 남겨놓는걸 기억하시면 됩니다.

이 과정을 끝내면 필요한만큼 나머지 분할 영역을 만듭니다. 나머지 모든 분할 영역을 배치하고 나면 스왑 분할 영역 ID를 리눅스 스왑 82번으로 설정했는지 확인하십시오. 기본값은, 리눅스 네이티브 83번입니다.

### Cobalt 드라이브 공간 분할

Cobalt 머신에서 BOOTROM은 MS-DOS MBR을 찾으므로 드라이브 공간 분할은 상대적으로 쉽습니다. 사실 인텔 x86 머신에서 진행하는 동일한 방법으로 끝냅니다. 그러나 알아두어야 할 몇가지가 있습니다.

- 앞으로 준비할 /dev/sda1 분할 영역의 Cobalt 펌웨어는 ext2 리비전 0로 포맷한 리눅스 분할 영역입니다. ext2 리비전 1 분할 영역으로는 동작하지 않습니다!(Cobalt BOOTROM은 ext2r0만 인식합니다)
- 위에 언급한 분할 영역에는 커널로서 읽어들이려 gzip으로 압축한 ELF 이미지 vmlinux.gz를 분할 영역의 루트에 넣어야합니다.

이런 이유로 CoLo와 커널을 설치할 분할 영역 /boot/은 EXT2r0로 포맷하고 20MB 미만의 공간으로 만드는 것이 좋습니다. 이를 통해 사용자는 루트 파일 시스템을 최신 파일 시스템(EXT3 또는 ReiserFS)를 다룰 수 있습니다.

예제를 통해 나중에 /boot/ 분할 영역으로 마운트할 /dev/sda1를 만들었다고 가정하겠습니다. / 를 만들려면 PROM의 기대 요구 값을 기억해두십시오.

따라서 계속 진행하겠습니다. 프롬프트에서 fdisk /dev/sda를 입력하여 분할 영역을 만드십시오. 주로 사용하는 명령은 다음과 같습니다:

코드 **중요한 fdisk 명령 목록**

```
'"`UNIQ--pre-00000007-QINU`"'

```

`root #` `fdisk /dev/sda`

```
The number of cylinders for this disk is set to 19870.
There is nothing wrong with that, but this is larger than 1024,
and could in certain setups cause problems with:
(1) software that runs at boot time (e.g., old versions of LILO)
(2) booting and partitioning software from other OSs
   (e.g., DOS FDISK, OS/2 FDISK)

```

기존 분할 영역을 지우는걸로 시작합니다:

`Command (m for help):` `o`

```
Building a new DOS disklabel. Changes will remain in memory only,
until you decide to write them. After that, of course, the previous
content won't be recoverable.


The number of cylinders for this disk is set to 19870.
There is nothing wrong with that, but this is larger than 1024,
and could in certain setups cause problems with:
(1) software that runs at boot time (e.g., old versions of LILO)
(2) booting and partitioning software from other OSs
   (e.g., DOS FDISK, OS/2 FDISK)
Warning: invalid flag 0x0000 of partition table 4 will be corrected by w(rite)

```

`p` 명령으로 분할 영역 테이블을 비웠는지 확인하십시오:

`Command (m for help):` `p`

```
Disk /dev/sda: 10.2 GB, 10254827520 bytes
16 heads, 63 sectors/track, 19870 cylinders
Units = cylinders of 1008 * 512 = 516096 bytes

   Device Boot      Start         End      Blocks   Id  System

```

/boot 분할 영역을 만드십시오:

`Command (m for help):` `n`

```
Command action
   e   extended
   p   primary partition (1-4)
p
Partition number (1-4): 1
First cylinder (1-19870, default 1):
Last cylinder or +size or +sizeM or +sizeK (1-19870, default 19870): +20M

```

분할 영역 내용을 표시할 때, 새로 만든 분할 영역을 확인하십시오:

`Command (m for help):` `p`

```
Disk /dev/sda: 10.2 GB, 10254827520 bytes
16 heads, 63 sectors/track, 19870 cylinders
Units = cylinders of 1008 * 512 = 516096 bytes

   Device Boot      Start         End      Blocks   Id  System
/dev/sda1               1          40       20128+  83  Linux

```

이제 나머지 디스크 영역을 차지하는 확장 분할 영역을 만들겠습니다. 확장 분할 영역에서 나머지 분할 영역(논리 분할 영역)을 만들겠습니다:

`Command (m for help):` `n`

```
Command action
   e   extended
   p   primary partition (1-4)
e
Partition number (1-4): 2
First cylinder (41-19870, default 41):
Using default value 41
Last cylinder or +size or +sizeM or +sizeK (41-19870, default 19870):
Using default value 19870

```

이제 /, /usr, /var 등의 분할 영역을 만들겠습니다.

`Command (m for help):` `n`

```
Command action
   l   logical (5 or over)
   p   primary partition (1-4)
l
First cylinder (41-19870, default 41):<Press ENTER>
Using default value 41
Last cylinder or +size or +sizeM or +sizeK (41-19870, default 19870): +500M

```

필요한 만큼 반복하십시오.

마지막일지 모르겠지만 이걸로 끝이 아닌 스왑 영역입니다. 최소한 250MB의 스왑 공간이 필요하며 1GB 정도면 충분합니다:

`Command (m for help):` `n`

```
Command action
   l   logical (5 or over)
   p   primary partition (1-4)
l
First cylinder (17294-19870, default 17294): <Press ENTER>
Using default value 17294
Last cylinder or +size or +sizeM or +sizeK (1011-19870, default 19870): <Press ENTER>
Using default value 19870

```

이미 확인했지만서도, 분할 영역 테이블을 다시 확인하면 모든 준비가 끝난 상태입니다.

`Command (m for help):` `p`

```
Disk /dev/sda: 10.2 GB, 10254827520 bytes
16 heads, 63 sectors/track, 19870 cylinders
Units = cylinders of 1008 * 512 = 516096 bytes

Device Boot      Start         End      Blocks      ID  System
/dev/sda1               1          21       10552+  83  Linux
/dev/sda2              22       19870    10003896    5  Extended
/dev/sda5              22        1037      512032+  83  Linux
/dev/sda6            1038        5101     2048224+  83  Linux
/dev/sda7            5102        9165     2048224+  83  Linux
/dev/sda8            9166       13229     2048224+  83  Linux
/dev/sda9           13230       17293     2048224+  83  Linux
/dev/sda10          17294       19870     1298776+  83  Linux

```

이제 #10 스왑 분할 영역이 여전히 83번 형식으로 되어 있지요? 적당한 형식으로 바꾸겠습니다.

`Command (m for help):` `t`

```
Partition number (1-10): 10
Hex code (type L to list codes): 82
Changed system type of partition 10 to 82 (Linux swap)

```

이제 확인해보겠습니다:

`Command (m for help):` `p`

```
Disk /dev/sda: 10.2 GB, 10254827520 bytes
16 heads, 63 sectors/track, 19870 cylinders
Units = cylinders of 1008 * 512 = 516096 bytes

Device Boot      Start         End      Blocks      ID  System
/dev/sda1               1          21       10552+  83  Linux
/dev/sda2              22       19870    10003896    5  Extended
/dev/sda5              22        1037      512032+  83  Linux
/dev/sda6            1038        5101     2048224+  83  Linux
/dev/sda7            5102        9165     2048224+  83  Linux
/dev/sda8            9166       13229     2048224+  83  Linux
/dev/sda9           13230       17293     2048224+  83  Linux
/dev/sda10          17294       19870     1298776+  82  Linux Swap

```

새 분할 영역 테이블을 기록하겠습니다:

`Command (m for help):` `w`

```
The partition table has been altered!

Calling ioctl() to re-read partition table.
Syncing disks.

```

## 파일 시스템 만들기

**경고**

When using an SSD or NVMe drive, it is wise to check for firmware upgrades. Some Intel SSDs in particular (600p and 6000p) require a firmware upgrade for [possible data corruption](https://bugzilla.redhat.com/show_bug.cgi?id=1402533) induced by XFS I/O usage patterns. The problem is at the firmware level and not any fault of the XFS filesystem. The smartctl utility can help check the device model and firmware version.

### 도입부

이제 분할 영역을 만들었고, 파일 시스템을 제 위치에 얹어놓을 차례입니다. 다음 절에서는 리눅스에서 지원하는 다양한 파일 시스템을 설명합니다. 어떤 파일 시스템을 사용할 지 이미 알고 있는 독자라면 [파티션에 파일 시스템 반영하기](#.ED.8C.8C.ED.8B.B0.EC.85.98.EC.97.90_.ED.8C.8C.EC.9D.BC_.EC.8B.9C.EC.8A.A4.ED.85.9C_.EB.B0.98.EC.98.81.ED.95.98.EA.B8.B0) 로 계속 진행할 수 있습니다. 그렇지 않으면 계속 읽어 내려가면서 쓸 수 있는 파일시스템이 어떤 종류가 있는지 알아보십시오.

### 파일 시스템

다양한 파일 시스템이 있습니다. 일부는 mips 아키텍처에서 안정적입니다 - 중요한 분할 영역을 위해서라면 좀 더 시험적인 분할 영역을 선택하기 전에 파일 시스템과 지원 상태에 대한 내용을 좀 더 읽어보시는 것이 좋겠습니다.

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

예를 들어, 예제 분할 영역 구조와 같이 ext2 형식의 부팅 분할 영역 (/dev/sda1) 과 ext4 형식의 루트 분할 영역 (/dev/sda5)을 취하려면, 다음 명령을 사용할 수 있습니다:

`root #` `mkfs.ext4 /dev/sda5`

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

`root #` `mkswap /dev/sda10`

**참고**

Installations which were previously started, but did not finish the installation process can resume the installation from this point in the handbook. Use this link as the permalink: [Resumed installations start here](/wiki/Handbook:MIPS/Installation/Disks/ko#Resumed_installations_start_here "Handbook:MIPS/Installation/Disks/ko").

스왑 분할 영역을 활성화하려면, swapon 명령을 사용하십시오:

`root #` `swapon /dev/sda10`

This 'activation' step is only necessary because the swap partition is newly created within the live environment. Once the system has been rebooted, as long as the swap partition is properly defined within fstab or other mount mechanism, swap space will activate automatically.

## 루트 분할 영역 마운트

Certain live environments may be missing the suggested mount point for Gentoo's root partition (/mnt/gentoo), or mount points for additional partitions created in the partitioning section:

`root #` `mkdir --parents /mnt/gentoo`

Continue creating additional mount points necessary for any additional (custom) partition(s) created during previous steps by using the mkdir command.

이제 분할 영역을 초기화했고 파일 시스템을 넣었으므로 분할 영역을 마운트할 차례입니다. mount 명령을 사용하지만 만들어놓은 모든 분할 영역에 대해 마운트 디렉터리를 만들 필요는 없다는 사실을 잊지 마십시오. 예제를 통해 우리는 루트 분할 영역을 마운트하겠습니다:

Mount the root partition:

`root #` `mount /dev/sda5 /mnt/gentoo`

Continue mounting additional (custom) partitions as necessary using the mount command.

**참고**

/tmp/를 따로 나눈 분할 영역에 두어야 한다면, 마운트하기 전에 퍼미션을 바꾸었는지 확인하십시오:

`root #` `chmod 1777 /mnt/gentoo/tmp`

이 설정은 /var/tmp에도 적용 유지합니다.

지침을 따르고 나면 proc 파일 시스템(커널 가상 인터페이스)와 다른 커널 의사 파일 시스템을 마운트합니다. 그러나 우선 [젠투 설치 파일을 설치](/wiki/Handbook:MIPS/Installation/Stage/ko "Handbook:MIPS/Installation/Stage/ko") 하겠습니다.

[← 네트워크 설정](/wiki/Handbook:MIPS/Installation/Networking/ko "이전 내용") [처음](/wiki/Handbook:MIPS/ko "Handbook:MIPS/ko") [젠투 설치 파일 설치 →](/wiki/Handbook:MIPS/Installation/Stage/ko "다음 내용")