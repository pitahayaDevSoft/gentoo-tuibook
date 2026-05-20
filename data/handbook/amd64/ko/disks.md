# Disks

Other languages:

- [Deutsch](/wiki/Handbook:AMD64/Installation/Disks/de "Handbook:AMD64/Installation/Disks/de (100% translated)")
- [English](/wiki/Handbook:AMD64/Installation/Disks "Handbook:AMD64/Installation/Disks (100% translated)")
- [Türkçe](/wiki/Handbook:AMD64/Installation/Disks/tr "Handbook:AMD64/Installation/Disks/tr (0% translated)")
- [español](/wiki/Handbook:AMD64/Installation/Disks/es "Handbook:AMD64/Installation/Disks/es (100% translated)")
- [français](/wiki/Handbook:AMD64/Installation/Disks/fr "Handbook:AMD64/Installation/Disks/fr (100% translated)")
- [italiano](/wiki/Handbook:AMD64/Installation/Disks/it "Handbook:AMD64/Installation/Disks/it (100% translated)")
- [magyar](/wiki/Handbook:AMD64/Installation/Disks/hu "Handbook:AMD64/Installation/Disks/hu (100% translated)")
- [polski](/wiki/Handbook:AMD64/Installation/Disks/pl "Handbook:AMD64/Installation/Disks/pl (100% translated)")
- [português do Brasil](/wiki/Handbook:AMD64/Installation/Disks/pt-br "Handbook:AMD64/Installation/Disks/pt-br (100% translated)")
- [čeština](/wiki/Handbook:AMD64/Installation/Disks/cs "Handbook:AMD64/Installation/Disks/cs (100% translated)")
- [русский](/wiki/Handbook:AMD64/Installation/Disks/ru "Handbook:AMD64/Installation/Disks/ru (100% translated)")
- [العربية](/wiki/Handbook:AMD64/Installation/Disks/ar "Handbook:AMD64/Installation/Disks/ar (100% translated)")
- [தமிழ்](/wiki/Handbook:AMD64/Installation/Disks/ta "Handbook:AMD64/Installation/Disks/ta (100% translated)")
- [中文（中国大陆）‎](/wiki/Handbook:AMD64/Installation/Disks/zh-cn "Handbook:AMD64/Installation/Disks/zh-cn (100% translated)")
- [日本語](/wiki/Handbook:AMD64/Installation/Disks/ja "Handbook:AMD64/Installation/Disks/ja (100% translated)")
- 한국어

[AMD64 핸드북](/wiki/Handbook:AMD64/ko "Handbook:AMD64/ko")[설치](/wiki/Handbook:AMD64/Full/Installation/ko "Handbook:AMD64/Full/Installation/ko")[설치 정보](/wiki/Handbook:AMD64/Installation/About/ko "Handbook:AMD64/Installation/About/ko")[매체 선택](/wiki/Handbook:AMD64/Installation/Media/ko "Handbook:AMD64/Installation/Media/ko")[네트워크 설정](/wiki/Handbook:AMD64/Installation/Networking/ko "Handbook:AMD64/Installation/Networking/ko")디스크 준비[스테이지 3 설치](/wiki/Handbook:AMD64/Installation/Stage/ko "Handbook:AMD64/Installation/Stage/ko")[베이스 시스템 설치](/wiki/Handbook:AMD64/Installation/Base/ko "Handbook:AMD64/Installation/Base/ko")[커널 설정](/wiki/Handbook:AMD64/Installation/Kernel/ko "Handbook:AMD64/Installation/Kernel/ko")[시스템 설정](/wiki/Handbook:AMD64/Installation/System/ko "Handbook:AMD64/Installation/System/ko")[도구 설치](/wiki/Handbook:AMD64/Installation/Tools/ko "Handbook:AMD64/Installation/Tools/ko")[부트로더 설정](/wiki/Handbook:AMD64/Installation/Bootloader/ko "Handbook:AMD64/Installation/Bootloader/ko")[마무리](/wiki/Handbook:AMD64/Installation/Finalizing/ko "Handbook:AMD64/Installation/Finalizing/ko")[젠투 활용](/wiki/Handbook:AMD64/Full/Working/ko "Handbook:AMD64/Full/Working/ko")[포티지 소개](/wiki/Handbook:AMD64/Working/Portage/ko "Handbook:AMD64/Working/Portage/ko")[USE 플래그](/wiki/Handbook:AMD64/Working/USE/ko "Handbook:AMD64/Working/USE/ko")[포티지 기능](/wiki/Handbook:AMD64/Working/Features/ko "Handbook:AMD64/Working/Features/ko")[초기화 스크립트 시스템](/wiki/Handbook:AMD64/Working/Initscripts/ko "Handbook:AMD64/Working/Initscripts/ko")[환경 변수](/wiki/Handbook:AMD64/Working/EnvVar/ko "Handbook:AMD64/Working/EnvVar/ko")[포티지 활용](/wiki/Handbook:AMD64/Full/Portage/ko "Handbook:AMD64/Full/Portage/ko")[파일 및 디렉터리](/wiki/Handbook:AMD64/Portage/Files/ko "Handbook:AMD64/Portage/Files/ko")[변수](/wiki/Handbook:AMD64/Portage/Variables/ko "Handbook:AMD64/Portage/Variables/ko")[소프트웨어 브랜치 함께 사용하기](/wiki/Handbook:AMD64/Portage/Branches/ko "Handbook:AMD64/Portage/Branches/ko")[추가 도구](/wiki/Handbook:AMD64/Portage/Tools/ko "Handbook:AMD64/Portage/Tools/ko")[꾸러미 저장소 개별 설정](/wiki/Handbook:AMD64/Portage/CustomTree/ko "Handbook:AMD64/Portage/CustomTree/ko")[고급 기능](/wiki/Handbook:AMD64/Portage/Advanced/ko "Handbook:AMD64/Portage/Advanced/ko")[네트워크 설정](/wiki/Handbook:AMD64/Full/Networking/ko "Handbook:AMD64/Full/Networking/ko")[시작하기](/wiki/Handbook:AMD64/Networking/Introduction/ko "Handbook:AMD64/Networking/Introduction/ko")[고급 설정](/wiki/Handbook:AMD64/Networking/Advanced/ko "Handbook:AMD64/Networking/Advanced/ko")[모듈러 네트워크](/wiki/Handbook:AMD64/Networking/Modular/ko "Handbook:AMD64/Networking/Modular/ko")[무선 네트워크](/wiki/Handbook:AMD64/Networking/Wireless/ko "Handbook:AMD64/Networking/Wireless/ko")[기능 추가](/wiki/Handbook:AMD64/Networking/Extending/ko "Handbook:AMD64/Networking/Extending/ko")[동적 관리](/wiki/Handbook:AMD64/Networking/Dynamic/ko "Handbook:AMD64/Networking/Dynamic/ko")

## Contents

- [1블록 장치 소개](#.EB.B8.94.EB.A1.9D_.EC.9E.A5.EC.B9.98_.EC.86.8C.EA.B0.9C)
  - [1.1블록 장치](#.EB.B8.94.EB.A1.9D_.EC.9E.A5.EC.B9.98)
  - [1.2분할 영역 테이블](#.EB.B6.84.ED.95.A0_.EC.98.81.EC.97.AD_.ED.85.8C.EC.9D.B4.EB.B8.94)
    - [1.2.1GPT](#GPT)
    - [1.2.2MBR](#MBR)
  - [1.3고급 저장장치](#.EA.B3.A0.EA.B8.89_.EC.A0.80.EC.9E.A5.EC.9E.A5.EC.B9.98)
  - [1.4기본 분할 형태](#.EA.B8.B0.EB.B3.B8_.EB.B6.84.ED.95.A0_.ED.98.95.ED.83.9C)
- [2분할 배치 설계](#.EB.B6.84.ED.95.A0_.EB.B0.B0.EC.B9.98_.EC.84.A4.EA.B3.84)
  - [2.1분할 영역을 얼마나 많이, 크게 할까요?](#.EB.B6.84.ED.95.A0_.EC.98.81.EC.97.AD.EC.9D.84_.EC.96.BC.EB.A7.88.EB.82.98_.EB.A7.8E.EC.9D.B4.2C_.ED.81.AC.EA.B2.8C_.ED.95.A0.EA.B9.8C.EC.9A.94.3F)
  - [2.2스왑 공간이 무엇인가요?](#.EC.8A.A4.EC.99.91_.EA.B3.B5.EA.B0.84.EC.9D.B4_.EB.AC.B4.EC.97.87.EC.9D.B8.EA.B0.80.EC.9A.94.3F)
    - [2.2.1UEFI 사용](#UEFI_.EC.82.AC.EC.9A.A9)
  - [2.3BIOS 부트 분할 영역이란?](#BIOS_.EB.B6.80.ED.8A.B8_.EB.B6.84.ED.95.A0_.EC.98.81.EC.97.AD.EC.9D.B4.EB.9E.80.3F)
- [3대안: 디스크 분할시 fdisk 사용](#.EB.8C.80.EC.95.88:_.EB.94.94.EC.8A.A4.ED.81.AC_.EB.B6.84.ED.95.A0.EC.8B.9C_fdisk_.EC.82.AC.EC.9A.A9)
  - [3.1현재 분할 영역 배치 보기](#.ED.98.84.EC.9E.AC_.EB.B6.84.ED.95.A0_.EC.98.81.EC.97.AD_.EB.B0.B0.EC.B9.98_.EB.B3.B4.EA.B8.B0)
  - [3.2fdisk에서 모든 분할 영역 제거](#fdisk.EC.97.90.EC.84.9C_.EB.AA.A8.EB.93.A0_.EB.B6.84.ED.95.A0_.EC.98.81.EC.97.AD_.EC.A0.9C.EA.B1.B0)
  - [3.3BIOS 부팅 분할 영역 만들기](#BIOS_.EB.B6.80.ED.8C.85_.EB.B6.84.ED.95.A0_.EC.98.81.EC.97.AD_.EB.A7.8C.EB.93.A4.EA.B8.B0)
  - [3.4스왑 분할 영역 만들기](#.EC.8A.A4.EC.99.91_.EB.B6.84.ED.95.A0_.EC.98.81.EC.97.AD_.EB.A7.8C.EB.93.A4.EA.B8.B0)
  - [3.5부팅 분할 영역 만들기](#.EB.B6.80.ED.8C.85_.EB.B6.84.ED.95.A0_.EC.98.81.EC.97.AD_.EB.A7.8C.EB.93.A4.EA.B8.B0)
  - [3.6루트 분할 영역 만들기](#.EB.A3.A8.ED.8A.B8_.EB.B6.84.ED.95.A0_.EC.98.81.EC.97.AD_.EB.A7.8C.EB.93.A4.EA.B8.B0)
  - [3.7분할 영역 배치 저장하기](#.EB.B6.84.ED.95.A0_.EC.98.81.EC.97.AD_.EB.B0.B0.EC.B9.98_.EC.A0.80.EC.9E.A5.ED.95.98.EA.B8.B0)
- [4Partitioning the disk with MBR for BIOS / legacy boot](#Partitioning_the_disk_with_MBR_for_BIOS_.2F_legacy_boot)
  - [4.1Viewing the current partition layout](#Viewing_the_current_partition_layout)
  - [4.2Creating a new disklabel / removing all partitions](#Creating_a_new_disklabel_.2F_removing_all_partitions)
  - [4.3Creating the boot partition](#Creating_the_boot_partition)
  - [4.4Creating the swap partition](#Creating_the_swap_partition)
  - [4.5Creating the root partition](#Creating_the_root_partition)
  - [4.6Saving the partition layout](#Saving_the_partition_layout)
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

### 분할 영역 테이블

비록 (예를 들어서 btrfs RAID를 만든다 치면) 이론적으로 리눅스 시스템을 저장할 디스크 전체를 분할하지 않고도 사용할 수 있다고는 하지만, 실제로는 거의 불가능합니다. 대신, 전체 디스크 블록 장치를 더 작게 나누고, 더 관리하기 쉬운 블록 장치로 만들 수 있습니다. **amd64** 시스템에서는, 분할 영역(파티션)이라고 합니다. 현재 MBR과 GPT 두가지 분할 표준 기술이 있습니다.

#### GPT

_GPT (GUID 분할 영역 테이블)_ 설정에선 분할 영역에 64비트 식별자를 사용합니다. 분할 영역 정보를 저장하는 위치는 512 바이트의 MBR보다 훨씬 크며, GPT 디스크 분할 영역의 수는 실제로 제한이 없습니다. 또한 분할 영역 크기도 훨씬 큰 제한 용량 값 범위에 들어갑니다(거의 8 ZiB - 예, 제타바이트입니다).

운영 체제와 펌웨어의 시스템 프로그램 인터페이스가 (BIOS 대신) UEFI일 때, GPT는 MBR로 인한 호환성 문제가 일어나는 현 상황에서 단연 필수라 할 수 있습니다.

GPT는 또한 검사합과 중복을 활용하는 이점이 있습니다. CRC32 검사합 값을 활용하여 헤더와 분할 영역 테이블의 오류를 찾으며 디스크의 마지막 부분에 GPT 정보를 백업합니다. 백업 데이블은 디스크 시작 부분의 주 GPT가 손상됐을 경우 복원할 때 사용할 수 있습니다.

**중요**

There are a few caveats regarding GPT:

- Using GPT on a BIOS-based computer works, but the user won't be able to dual-boot with a Microsoft Windows operating system, since Microsoft Windows refuses to boot from a GPT partition when in BIOS mode.
- Some buggy (old) motherboard firmware configured to boot in BIOS/CSM/legacy mode might also have problems with booting from GPT labeled disks.

#### MBR

_MBR(주 부트 레코드)_ 설정은 시작 섹터 및 분할 영역의 길이, 그리고 다음의 분할 영역 형식을 지원하는 32비트 식별자를 사용합니다: 주, 확장, 논리. 주 분할 영역은 마스터 부트 레코드 자체에 저장한 정보를 지니고 있습니다. 주 부트 레코드는 매우 작으며(보통 512 바이트) 디스크의 맨 처음에 위치합니다. 공간이 작기 때문에 오직 네 개의 주 분할 영역만을 지원합니다(예를 들자면, /dev/sda1부터 /dev/sda4까지).

더 많은 분할 영역을 지원하려면 주 분할 영역 중 하나를 확장 분할 영역으로 표시할 수 있습니다. 이 분할 영역은 논리 분할 영역(분할 영역 안의 분할 영역)을 보유할 수 있습니다.

**중요**

대부분 마더보드 제조사에서 지원하지만 분할 테이블을 오래된 기술로 간주합니다. 2010년 이전 생산 하드웨어에서 동작하지 않는다면, [GUID 분할 테이블](https://en.wikipedia.org/wiki/GUID_Partition_Table "wikipedia:GUID Partition Table") 방식으로 디스크를 분할하시는 방법이 최선입니다. MBR 방식으로 진행하려는 독자 여러분은 다음 내용을 참고바랍니다:

- 대부분의 2010년 이후 생산 마더보드에서는 MBR 분할 방식을 오래된 (지원하지만 이상적은 아님) 부팅 설정으로 간주합니다.
- 32비트 식별자는 2TiB 이상의 디스크를 MBR 방식으로 분할 처리하지 못합니다.
- 확장 분할 영역을 사용하지 않는 한, MBR 방식은 4개의 최대 분할 영역을 지원합니다.
- MBR 설정은 MBR 설정을 대체하는 어떤 기술도 지원하지 않습니다. 따라서 어떤 프로그램이나 사용자가 MBR을 덮어쓸 경우 모든 분할 영역 정보를 분실합니다.

핸드북 저자는 젠투를 설치할 경우라면 [GPT](#GPT) 사용을 추천합니다.

### 고급 저장장치

**amd64** 설치 CD에서는 논리 볼륨 관리자(LVM)을 지원합니다. LVM에서는 분할 영역 설정을 통해 유연성을 제공합니다. 아래 설치 과정은 _일반_ 분할 영역을 중점적으로 다루지만, 이런 방식을 원할 경우를 대비하여 LVM을 지원한다는 사실을 알아두시는 게 유익합니다. 자세한 내용은 [LVM](/wiki/LVM "LVM") 게시글을 살펴보십시오. 초보자 여러분은 LVM의 완전한 설명은 이 안내서의 내용 범위에 벗어난다는 사실을 알아두셨으면 합니다.

Although usage is not covered in the handbook, below is a list helpful guides to get the system running:

- [Btrfs](/wiki/Btrfs/Native_System_Root_Guide "Btrfs/Native System Root Guide")
- [ZFS](/wiki/ZFS/rootfs "ZFS/rootfs")
- [LVM (Logical Volume Manager)](/wiki/LVM/ko "LVM/ko")

Disk encryption is also handled in the same manner:

- [Rootfs encryption](/wiki/Rootfs_encryption "Rootfs encryption")

### 기본 분할 형태

Throughout the remainder of the handbook, we will discuss and explain two cases:

1. UEFI firmware with GUID Partition Table (GPT) disk.
2. MBR DOS/legacy BIOS firmware with a MBR partition table disk.

While it is possible to mix and match boot types with certain motherboard firmware, mixing goes beyond the intention of the handbook. As previously stated, it is strongly recommended for installations on modern hardware to use UEFI boot with a GPT disklabel disk.

이 핸드북의 나머지 부분에서는 다음 공간 분할 형태를 간단한 배치 예제로 활용합니다:

**중요**

The first row of the following table contains exclusive information for _**either**_ a GPT disklabel _**or**_ a MBR DOS/legacy BIOS disklabel. When in doubt, proceed with GPT, since **amd64** machines manufactured after the year 2010 generally support UEFI firmware and GPT boot sector.

분할 영역
파일 시스템
크기
설명
/dev/sda1(부트로더)
2M
BIOS 부트 분할 영역
/dev/sda1ext2 (UEFI를 사용하는 경우 fat32)
128M
부트/EFI 시스템 분할 영역
/dev/sda3(스왑)
512M 이상
스왑 분할 영역
/dev/sda4ext4
디스크 나머지 부분
루트 분할 영역

이 내용으로 충분하고 독자 여러분이 GPT 방식으로 진행하겠다면 [기본: 디스크 분할시 parted 사용](#.EA.B8.B0.EB.B3.B8:_.EB.94.94.EC.8A.A4.ED.81.AC_.EB.B6.84.ED.95.A0.EC.8B.9C_parted_.EC.82.AC.EC.9A.A9) 편으로 넘어갈 수 있습니다. 여전히 MBR 방식(이보쇼?!)을 선호하고 예제 배치를 따라가겠다면 [대안: 디스크 분할시 fdisk 사용](#.EB.8C.80.EC.95.88:_.EB.94.94.EC.8A.A4.ED.81.AC_.EB.B6.84.ED.95.A0.EC.8B.9C_fdisk_.EC.82.AC.EC.9A.A9) 편으로 넘어갈 수 있습니다.

fdisk와 parted는 디스크 공간 분할 유틸리티입니다. fdisk는 잘 알려진, 안정적인 프로그램이며 MBR 공간 분할 방식에 추천하지만 parted는 GPT 공간 분할 방식을 지원하는 리눅스용 첫 블록 장치 관리 유틸리티 중 하나입니다. fdisk 인터페이스 방식을 선호하는 사용자라면 parted 대신 gdisk(GPT fdisk)를 사용할 수 있습니다.

만들기 절차를 진행하기 전에 해당 절의 첫번째 부분에서는 분할 영역 형태를 어떻게 만드는지와 일반적인 문제를 언급하는 내용을 다루겠습니다.

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

#### UEFI 사용

운영체제를 부팅할 때 (BIOS 대신) UEFI를 사용하는 시스템에 젠투를 설치할 경우, EFI 시스템 분할 영역(ESP)을 만드는 과정이 중요합니다. 아래에 설명할 parted의 절차에는 이 동작을 올바르게 처리하는데 필요한 포인터가 있습니다.

ESP는 (때로는 리눅스 시스템에서 _vfat_ 으로 나타나는) FAT 종류 중 하나여야 합니다. 공식 [UEFI 명세](http://www.uefi.org/sites/default/files/resources/UEFI%202_5.pdf) en에서는 UEFI 펌웨어에서, ESP를 활용할 때 FAT32를 추천하지만, FAT 12, 16, 32 파일 시스템을 인식한다고 명시되어 있습니다. ESP를 FAT32로 포맷하기로 합니다:

`root #` `mkfs.fat -F 32 /dev/sda1`

**중요**

FAT 변종을 ESP에 사용하지 않으면 시스템의 UEFI 펌웨어를 부트로더(또는 리눅스 커널)에서 찾을 수 있음을 장담할 수 없으며 대부분의 경우 시스템 부팅이 불가능합니다!

### BIOS 부트 분할 영역이란?

BIOS 부팅 분할 영역은 매우 작(1~2MB)으며 GRUB2와 같은 부트로더가 할당 저장 공간(MBR의 경우 몇 백 메가바이트정도 저장함)에 맞지 않는 추가 데이터를 저장할 수 있으며, 어떤 위치에든 둘 수 있는건 아닙니다.

## 대안: 디스크 분할시 fdisk 사용

다음 부분에서는 fdisk를 사용하여 예제 분할 영역 배치를 만드는 방법을 설명합니다. 예제 분할 영역 배치는 앞서 언급했습니다:

개인 취향에 따라 분할 영역 배치를 바꾸십시오.

분할 영역
설명
/dev/sda1BIOS 부팅 분할 영역
/dev/sda1부팅 분할 영역
/dev/sda3스왑 분할 영역
/dev/sda4루트 분할 영역

### 현재 분할 영역 배치 보기

fdisk는 디스크를 분할 영역으로 나누는 유명하고 강력한 도구입니다. fdisk를 디스크에서 다시 실행해보십시오(예제에서는 /dev/sda를 사용):

`root #` `fdisk /dev/sda`

`p` 키를 사용하여 현재 분할 영역 설정을 표시하십시오:

`Command (m for help):` `p`

```
Disk /dev/sda: 240 heads, 63 sectors, 2184 cylinders
Units = cylinders of 15120 * 512 bytes

   Device Boot    Start       End    Blocks   Id  System
/dev/sda1   *         1        14    105808+  83  Linux
/dev/sda2            15        49    264600   82  Linux swap
/dev/sda3            50        70    158760   83  Linux
/dev/sda4            71      2184  15981840    5  Extended
/dev/sda5            71       209   1050808+  83  Linux
/dev/sda6           210       348   1050808+  83  Linux
/dev/sda7           349       626   2101648+  83  Linux
/dev/sda8           627       904   2101648+  83  Linux
/dev/sda9           905      2184   9676768+  83  Linux

```

Device Start End Sectors Size Type
/dev/sda1 2048 2099199 2097152 1G EFI System
/dev/sda2 2099200 10487807 8388608 4G Linux swap
/dev/sda3 10487808 1953523711 1943035904 926.5G Linux root (x86-64)

}}

각각의 디스크에는 하나의 스왑 분할 영역과("Linux Swap"으로 나타남) 7개의 리눅스 시스템(각각의 분할 영역이 "Linux"라고 나타남)을 넣도록 설정했습니다.

### fdisk에서 모든 분할 영역 제거

Pressing the `g` key will instantly remove all existing disk partitions and create a new GPT disklabel:

`Command (m for help):` `g`

```
Created a new GPT disklabel (GUID: 3E56EE74-0571-462B-A992-9872E3855D75).

```

디스크에서 모든 분할 영역을 제거하십시오. `d` 를 입력하여 분할 영역을 삭제하십시오. 기존의 /dev/sda1 분할 영역을 삭제하려면:

`Command (m for help):` `d`

```
Partition number (1-4): 1

```

분할 영역을 삭제할 예정입니다. 분할 영역 목록에는 더이상 나타나지 않겠지만( `p`) 바뀐 내용을 저장하기 전에는 지워지지 않습니다. 사용자가 실수했을 경우 진행을 멈출 수 있습니다. 이 경우 `q` 를 바로 입력하고 `Enter` 를 치면 분할 영역을 삭제하지않습니다.

`p` 를 입력하여 분할 영역 목록을 출력하고 `d` 와 삭제할 분할 영역 번호를 입력하는 과정을 반복하십시오. 최종적으로, 분할 영역 테이블을 비웠습니다:

`Command (m for help):` `p`

```
Disk /dev/sda: 30.0 GB, 30005821440 bytes
240 heads, 63 sectors/track, 3876 cylinders
Units = cylinders of 15120 * 512 = 7741440 bytes

Device Boot    Start       End    Blocks   Id  System

```

이제 메모리에 있는 분할 영역 테이블을 비웠고, 분할 영역을 만들 준비를 끝냈습니다.

### BIOS 부팅 분할 영역 만들기

**참고**

A smaller ESP is possible but not recommended, especially given it may be shared with other OSes.

먼저 매우 작은 BIOS 부팅 분할 영역을 만들겠습니다. 새 분할 영역을 만드는 명령 `n` 을 입력하시고, 주 분할 영역을 선택할 `p` 를 입력한 후, `1` 을 입력하여 첫번째 주 분할 영역을 선택하십시오. 첫번째 섹터를 물어보면 2048부터 시작하는지(부트로더에 필요함)확인한 후 `Enter` 를 치십시오. 마지막 섹터를 물어보면 +2M을 입력하여 2MB 크기의 분할 영역을 만드십시오:

`Command (m for help):` `n`

```
Command action
  e   extended
  p   primary partition (1-4)
p
Partition number (1-4): 1
First sector (64-10486533532, default 64): 2048
Last sector, +sectors +size{M,K,G} (4096-10486533532, default 10486533532): +2M

```

Do you want to remove the signature? \[Y\]es/\[N\]o: Y
The signature will be removed by a write command.

}}

UEFI 용도로 분할 영역을 표시하십시오:

`Command (m for help):` `t`

```
Selected partition 1
Hex code (type L to list codes): 4
Changed system type of partition 1 to 4 (BIOS boot)

```

### 스왑 분할 영역 만들기

마지막으로, 루트 분할 영역을 만들려면, 새 분할 영역을 만드는 명령 `n` 을 입력하시고, `p` 를 입력하여 `fdisk` 에게 주 분할 영역 만들기를 지시한 후, `3` 을 입력하여 세번째 주 분할 영역 /dev/sda3을 만드십시오. 첫번째 섹터를 물어보면 `Enter` 를 치십시오. 마지막 섹터를 물어보면 +512M(또는 필요한 스왑 공간만큼)을 입력하여 512MB 크기의 분할 영역을 만드십시오.

### 부팅 분할 영역 만들기

이 과정이 끝난 후 분할 영역 형식을 설정하는 명령 `t` 를 입력하고, `3` 을 입력하여 방금 만든 분할 영역을 입력한 후, _82_ 를 입력하여 "Linux Swap" 형식으로 바꾸십시오.

`Command (m for help):` `t`

```
Partition number (1,2, default 2): 2
Partition type or alias (type L to list all): 19

Changed type of partition 'Linux filesystem' to 'Linux swap'.

```

### 루트 분할 영역 만들기

마지막으로, 루트 분할 영역을 만들려면, 새 분할 영역을 만드는 명령 `n` 을 입력하시고, `p` 를 입력하여 fdisk에게 주 분할 영역 만들기를 지시한 후, `4` 를 입력하여 네번째 주 분할 영역 /dev/sda4를 만드십시오. 첫번째 섹터를 물어보면 `Enter` 를 치십시오. 마지막 섹터를 물어보면 `Enter` 를 쳐서 디스크의 나머지 공간을 취하는 분할 영역을 만드십시오. 이 과정이 끝나면 `p` 를 입력하였을 때 다음과 같은 분할 영역 테이블 모습이 나타나야합니다:

`Command (m for help):` `n`

```
Partition number (3-128, default 3): 3
First sector (10487808-1953525134, default 10487808):
Last sector, +/-sectors or +/-size{K,M,G,T,P} (10487808-1953525134, default 1953523711):
</div>

<div lang="en" dir="ltr" class="mw-content-ltr">
Created a new partition 3 of type 'Linux filesystem' and of size 926.5 GiB..

```

**참고**

Setting the root partition's type to "Linux root (x86-64)" is not required and the system will function normally if it is set to the "Linux filesystem" type. This filesystem type is only necessary for cases where a bootloader that supports it (i.e. systemd-boot) is used and a fstab file is not wanted.

After creating the root partition, press `t` to set the partition type, `3` to select the partition just created, and then type in _23_ to set the partition type to "Linux Root (x86-64)".

`Command(m for help):` `t`

```
Partition number (1-3, default 3): 3
Partition type or alias (type L to list all): 23
</div>

<div lang="en" dir="ltr" class="mw-content-ltr">
Changed type of partition 'Linux filesystem' to 'Linux root (x86-64)'

```

After completing these steps, pressing `p` should display a partition table that looks similar to the following:

`Command (m for help):` `p`

```
Disk /dev/sda: 30.0 GB, 30005821440 bytes
240 heads, 63 sectors/track, 3876 cylinders
Units = cylinders of 15120 * 512 = 7741440 bytes

   Device Boot    Start       End    Blocks   Id  System
/dev/sda1             1         3      5198+  ef  EFI (FAT-12/16/32)
/dev/sda2   *         3        14    105808+  83  Linux
/dev/sda3            15        81    506520   82  Linux swap
/dev/sda4            82      3876  28690200   83  Linux

```

Device Start End Sectors Size Type
/dev/sda1 2048 2099199 2097152 1G EFI System
/dev/sda2 2099200 10487807 8388608 4G Linux swap
/dev/sda3 10487808 1953523711 1943035904 926.5G Linux root (x86-64)

Filesystem/RAID signature on partition 1 will be wiped.

}}

### 분할 영역 배치 저장하기

분할 영역 배치를 저장하고 fdisk를 빠져나가려면 `w` 를 입력하십시오.

`Command (m for help):` `w`

분할 영역을 만들었다면 이제 파일 시스템을 올려놓을 차례입니다.

## Partitioning the disk with MBR for BIOS / legacy boot

The following table provides a recommended partition layout for a trivial MBR DOS / legacy BIOS boot installation. Additional partitions can be added according to personal preference or system design goals.

Device path (sysfs)
Mount point
File system
DPS UUID (PARTUUID)
Description
/dev/sda1/bootxfs
N/A
MBR DOS / legacy BIOS boot partition details.
/dev/sda2N/A. Swap is not mounted to the filesystem like a device file.swap
0657fd6d-a4ab-43c4-84e5-0933c84b4f4f
Swap partition details.
/dev/sda3/xfs
4f68bce3-e8cd-4db1-96e7-fbcaf984b709
Root partition details.

Change the partition layout according to personal preference.

### Viewing the current partition layout

Fire up fdisk against the disk (in our example, we use /dev/sda):

`root #` `fdisk /dev/sda`

Use the `p` key to display the disk's current partition configuration:

`Command (m for help):` `p`

```
Disk /dev/sda: 931.51 GiB, 1000204886016 bytes, 1953525168 sectors
Disk model: HGST HTS721010A9
Units: sectors of 1 * 512 = 512 bytes
Sector size (logical/physical): 512 bytes / 4096 bytes
I/O size (minimum/optimal): 4096 bytes / 4096 bytes
Disklabel type: dos
Disk identifier: 0xf163b576
</div>

<div lang="en" dir="ltr" class="mw-content-ltr">
Device     Boot    Start        End    Sectors   Size Id Type
/dev/sda1  *        2048    2099199    2097152     1G 83 Linux
/dev/sda2        2099200   10487807    8388608     4G 82 Linux swap / Solaris
/dev/sda3       10487808 1953525167 1943037360 926.5G 83 Linux

```

This particular disk was until now configured to house two Linux filesystems (each with a corresponding partition listed as "Linux") as well as a swap partition (listed as "Linux swap"), using a GPT table.

### Creating a new disklabel / removing all partitions

Pressing `o` will instantly remove all existing disk partitions and create a new MBR disklabel (also named DOS disklabel):

`Command (m for help):` `o`

```
Created a new DOS disklabel with disk identifier 0xf163b576.
The device contains 'gpt' signature and it will be removed by a write command. See fdisk(8) man page and --wipe option for more details.

```

Alternatively, to keep an existing DOS disklabel (see the output of `p` above), consider removing the existing partitions one by one from the disk. Press `d` to delete a partition. For instance, to delete an existing /dev/sda1:

`Command (m for help):` `d`

```
Partition number (1-4): 1

```

The partition has now been scheduled for deletion. It will no longer show up when printing the list of partitions ( `p`, but it will not be erased until the changes have been saved. This allows users to abort the operation if a mistake was made - in that case, type `q` immediately and hit `Enter` and the partition will not be deleted.

Repeatedly press `p` to print out a partition listing and then press `d` and the number of the partition to delete it. Eventually, the partition table will be empty:

`Command (m for help):` `p`

```
Disk /dev/sda: 931.51 GiB, 1000204886016 bytes, 1953525168 sectors
Disk model: HGST HTS721010A9
Units: sectors of 1 * 512 = 512 bytes
Sector size (logical/physical): 512 bytes / 4096 bytes
I/O size (minimum/optimal): 4096 bytes / 4096 bytes
Disklabel type: dos
Disk identifier: 0xf163b576

```

The disk is now ready to create new partitions.

### Creating the boot partition

First, create a small partition which will be mounted as /boot. Press `n` to create a new partition, followed by `p` for a _primary_ partition and `1` to select the first primary partition. When prompted for the first sector, make sure it starts from 2048 (which may be needed for the boot loader) and press `Enter`. When prompted for the last sector, type +1G to create a partition 1 GB in size:

`Command (m for help):` `n`

```
Partition type
   p   primary (0 primary, 0 extended, 4 free)
   e   extended (container for logical partitions)
Select (default p): p
Partition number (1-4, default 1): 1
First sector (2048-1953525167, default 2048):
Last sector, +/-sectors or +/-size{K,M,G,T,P} (2048-1953525167, default 1953525167): +1G
</div>

<div lang="en" dir="ltr" class="mw-content-ltr">
Created a new partition 1 of type 'Linux' and of size 1 GiB.

```

Mark the partition as bootable by pressing the `a` key and pressing `Enter`:

`Command (m for help):` `a`

```
Selected partition 1
The bootable flag on partition 1 is enabled now.

```

Note: if more than one partition is available on the disk, then the partition to be flagged as bootable will have to be selected.

### Creating the swap partition

Next, to create the swap partition, press `n` to create a new partition, then `p`, then type `2` to create the second primary partition, /dev/sda2. When prompted for the first sector, press `Enter`. When prompted for the last sector, type +4G (or any other size needed for the swap space) to create a partition 4GB in size.

`Command (m for help):` `n`

```
Partition type
   p   primary (1 primary, 0 extended, 3 free)
   e   extended (container for logical partitions)
Select (default p): p
Partition number (2-4, default 2): 2
First sector (2099200-1953525167, default 2099200):
Last sector, +/-sectors or +/-size{K,M,G,T,P} (2099200-1953525167, default 1953525167): +4G

Created a new partition 2 of type 'Linux' and of size 4 GiB.

```

After all this is done, press `t` to set the partition type, `2` to select the partition just created and then type in _82_ to set the partition type to "Linux Swap".

`Command (m for help):` `t`

```
Partition number (1,2, default 2): 2
Hex code (type L to list all codes): 82
</div>

<div lang="en" dir="ltr" class="mw-content-ltr">
Changed type of partition 'Linux' to 'Linux swap / Solaris'.

```

### Creating the root partition

Finally, to create the root partition, press `n` to create a new partition. Then press `p` and `3` to create the third primary partition, /dev/sda3. When prompted for the first sector, hit `Enter`. When prompted for the last sector, hit `Enter` to create a partition that takes up the rest of the remaining space on the disk:

`Command (m for help):` `n`

```
Partition type
   p   primary (2 primary, 0 extended, 2 free)
   e   extended (container for logical partitions)
Select (default p): p
Partition number (3,4, default 3): 3
First sector (10487808-1953525167, default 10487808):
Last sector, +/-sectors or +/-size{K,M,G,T,P} (10487808-1953525167, default 1953525167):
</div>

<div lang="en" dir="ltr" class="mw-content-ltr">
Created a new partition 3 of type 'Linux' and of size 926.5 GiB.

```

After completing these steps, pressing `p` should display a partition table that looks similar to this:

`Command (m for help):` `p`

```
Disk /dev/sda: 931.51 GiB, 1000204886016 bytes, 1953525168 sectors
Disk model: HGST HTS721010A9
Units: sectors of 1 * 512 = 512 bytes
Sector size (logical/physical): 512 bytes / 4096 bytes
I/O size (minimum/optimal): 4096 bytes / 4096 bytes
Disklabel type: dos
Disk identifier: 0xf163b576
</div>

<div lang="en" dir="ltr" class="mw-content-ltr">
Device     Boot    Start        End    Sectors   Size Id Type
/dev/sda1  *        2048    2099199    2097152     1G 83 Linux
/dev/sda2        2099200   10487807    8388608     4G 82 Linux swap / Solaris
/dev/sda3       10487808 1953525167 1943037360 926.5G 83 Linux

```

### Saving the partition layout

Press `w` to write the partition layout and exit fdisk:

`Command (m for help):` `w`

```
The partition table has been altered.
Calling ioctl() to re-read partition table.
Syncing disks.

```

Now it is time to put filesystems on the partitions.

## 파일 시스템 만들기

**경고**

When using an SSD or NVMe drive, it is wise to check for firmware upgrades. Some Intel SSDs in particular (600p and 6000p) require a firmware upgrade for [possible data corruption](https://bugzilla.redhat.com/show_bug.cgi?id=1402533) induced by XFS I/O usage patterns. The problem is at the firmware level and not any fault of the XFS filesystem. The smartctl utility can help check the device model and firmware version.

### 도입부

이제 분할 영역을 만들었고, 파일 시스템을 제 위치에 얹어놓을 차례입니다. 다음 절에서는 리눅스에서 지원하는 다양한 파일 시스템을 설명합니다. 어떤 파일 시스템을 사용할 지 이미 알고 있는 독자라면 [파티션에 파일 시스템 반영하기](#.ED.8C.8C.ED.8B.B0.EC.85.98.EC.97.90_.ED.8C.8C.EC.9D.BC_.EC.8B.9C.EC.8A.A4.ED.85.9C_.EB.B0.98.EC.98.81.ED.95.98.EA.B8.B0) 로 계속 진행할 수 있습니다. 그렇지 않으면 계속 읽어 내려가면서 쓸 수 있는 파일시스템이 어떤 종류가 있는지 알아보십시오.

### 파일 시스템

다양한 파일 시스템이 있습니다. 일부는 amd64 아키텍처에서 안정적입니다 - 중요한 분할 영역을 위해서라면 좀 더 시험적인 분할 영역을 선택하기 전에 파일 시스템과 지원 상태에 대한 내용을 좀 더 읽어보시는 것이 좋겠습니다.

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

{{#ifeq: 1\|1\|

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

Installations which were previously started, but did not finish the installation process can resume the installation from this point in the handbook. Use this link as the permalink: [Resumed installations start here](/wiki/Handbook:AMD64/Installation/Disks/ko#Resumed_installations_start_here "Handbook:AMD64/Installation/Disks/ko").

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

For EFI installs only, the ESP should be mounted under the root partition location:

`root #` `mkdir --parents /mnt/gentoo/efi`

Continue mounting additional (custom) partitions as necessary using the mount command.

**참고**

/tmp/를 따로 나눈 분할 영역에 두어야 한다면, 마운트하기 전에 퍼미션을 바꾸었는지 확인하십시오:

`root #` `chmod 1777 /mnt/gentoo/tmp`

이 설정은 /var/tmp에도 적용 유지합니다.

지침을 따르고 나면 proc 파일 시스템(커널 가상 인터페이스)와 다른 커널 의사 파일 시스템을 마운트합니다. 그러나 우선 [젠투 설치 파일을 설치](/wiki/Handbook:AMD64/Installation/Stage/ko "Handbook:AMD64/Installation/Stage/ko") 하겠습니다.

[← 네트워크 설정](/wiki/Handbook:AMD64/Installation/Networking/ko "이전 내용") [처음](/wiki/Handbook:AMD64/ko "Handbook:AMD64/ko") [젠투 설치 파일 설치 →](/wiki/Handbook:AMD64/Installation/Stage/ko "다음 내용")