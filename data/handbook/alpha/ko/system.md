# System

Other languages:

- [Deutsch](/wiki/Handbook:Alpha/Installation/System/de "Handbuch:Alpha/Installation/System (100% translated)")
- [English](/wiki/Handbook:Alpha/Installation/System "Handbook:Alpha/Installation/System (100% translated)")
- [español](/wiki/Handbook:Alpha/Installation/System/es "Manual de Gentoo: Alpha/Instalación/Sistema (100% translated)")
- [français](/wiki/Handbook:Alpha/Installation/System/fr "Handbook:Alpha/Installation/System/fr (100% translated)")
- [italiano](/wiki/Handbook:Alpha/Installation/System/it "Handbook:Alpha/Installation/System/it (100% translated)")
- [magyar](/wiki/Handbook:Alpha/Installation/System/hu "Handbook:Alpha/Installation/System/hu (100% translated)")
- [polski](/wiki/Handbook:Alpha/Installation/System/pl "Handbook:Alpha/Installation/System (100% translated)")
- [português do Brasil](/wiki/Handbook:Alpha/Installation/System/pt-br "Manual:Alpha/Instalação/Sistema (100% translated)")
- [čeština](/wiki/Handbook:Alpha/Installation/System/cs "Handbook:Alpha/Installation/System/cs (50% translated)")
- [русский](/wiki/Handbook:Alpha/Installation/System/ru "Handbook:Alpha/Installation/System (100% translated)")
- [தமிழ்](/wiki/Handbook:Alpha/Installation/System/ta "கையேடு:Alpha/நிறுவல்/முறைமை (100% translated)")
- [中文（中国大陆）‎](/wiki/Handbook:Alpha/Installation/System/zh-cn "手册：Alpha/安装/配置系统 (100% translated)")
- [日本語](/wiki/Handbook:Alpha/Installation/System/ja "ハンドブック:Alpha/インストール/システム (100% translated)")
- 한국어

[Alpha 핸드북](/wiki/Handbook:Alpha/ko "Handbook:Alpha/ko")[설치](/wiki/Handbook:Alpha/Full/Installation/ko "Handbook:Alpha/Full/Installation/ko")[설치 정보](/wiki/Handbook:Alpha/Installation/About/ko "Handbook:Alpha/Installation/About/ko")[매체 선택](/wiki/Handbook:Alpha/Installation/Media/ko "Handbook:Alpha/Installation/Media/ko")[네트워크 설정](/wiki/Handbook:Alpha/Installation/Networking/ko "Handbook:Alpha/Installation/Networking/ko")[디스크 준비](/wiki/Handbook:Alpha/Installation/Disks/ko "Handbook:Alpha/Installation/Disks/ko")[스테이지 3 설치](/wiki/Handbook:Alpha/Installation/Stage/ko "Handbook:Alpha/Installation/Stage/ko")[베이스 시스템 설치](/wiki/Handbook:Alpha/Installation/Base/ko "Handbook:Alpha/Installation/Base/ko")[커널 설정](/wiki/Handbook:Alpha/Installation/Kernel/ko "Handbook:Alpha/Installation/Kernel/ko")시스템 설정[도구 설치](/wiki/Handbook:Alpha/Installation/Tools/ko "Handbook:Alpha/Installation/Tools/ko")[부트로더 설정](/wiki/Handbook:Alpha/Installation/Bootloader/ko "Handbook:Alpha/Installation/Bootloader/ko")[마무리](/wiki/Handbook:Alpha/Installation/Finalizing/ko "Handbook:Alpha/Installation/Finalizing/ko")[젠투 활용](/wiki/Handbook:Alpha/Full/Working/ko "Handbook:Alpha/Full/Working/ko")[포티지 소개](/wiki/Handbook:Alpha/Working/Portage/ko "Handbook:Alpha/Working/Portage/ko")[USE 플래그](/wiki/Handbook:Alpha/Working/USE/ko "Handbook:Alpha/Working/USE/ko")[포티지 기능](/wiki/Handbook:Alpha/Working/Features/ko "Handbook:Alpha/Working/Features/ko")[초기화 스크립트 시스템](/wiki/Handbook:Alpha/Working/Initscripts/ko "Handbook:Alpha/Working/Initscripts/ko")[환경 변수](/wiki/Handbook:Alpha/Working/EnvVar/ko "Handbook:Alpha/Working/EnvVar/ko")[포티지 활용](/wiki/Handbook:Alpha/Full/Portage/ko "Handbook:Alpha/Full/Portage/ko")[파일 및 디렉터리](/wiki/Handbook:Alpha/Portage/Files/ko "Handbook:Alpha/Portage/Files/ko")[변수](/wiki/Handbook:Alpha/Portage/Variables/ko "Handbook:Alpha/Portage/Variables/ko")[소프트웨어 브랜치 함께 사용하기](/wiki/Handbook:Alpha/Portage/Branches/ko "Handbook:Alpha/Portage/Branches/ko")[추가 도구](/wiki/Handbook:Alpha/Portage/Tools/ko "Handbook:Alpha/Portage/Tools/ko")[꾸러미 저장소 개별 설정](/wiki/Handbook:Alpha/Portage/CustomTree/ko "Handbook:Alpha/Portage/CustomTree/ko")[고급 기능](/wiki/Handbook:Alpha/Portage/Advanced/ko "Handbook:Alpha/Portage/Advanced/ko")[네트워크 설정](/wiki/Handbook:Alpha/Full/Networking/ko "Handbook:Alpha/Full/Networking/ko")[시작하기](/wiki/Handbook:Alpha/Networking/Introduction/ko "Handbook:Alpha/Networking/Introduction/ko")[고급 설정](/wiki/Handbook:Alpha/Networking/Advanced/ko "Handbook:Alpha/Networking/Advanced/ko")[모듈러 네트워크](/wiki/Handbook:Alpha/Networking/Modular/ko "Handbook:Alpha/Networking/Modular/ko")[무선 네트워크](/wiki/Handbook:Alpha/Networking/Wireless/ko "Handbook:Alpha/Networking/Wireless/ko")[기능 추가](/wiki/Handbook:Alpha/Networking/Extending/ko "Handbook:Alpha/Networking/Extending/ko")[동적 관리](/wiki/Handbook:Alpha/Networking/Dynamic/ko "Handbook:Alpha/Networking/Dynamic/ko")

## Contents

- [1파일 시스템 정보](#.ED.8C.8C.EC.9D.BC_.EC.8B.9C.EC.8A.A4.ED.85.9C_.EC.A0.95.EB.B3.B4)
  - [1.1파일 시스템 레이블과 UUID](#.ED.8C.8C.EC.9D.BC_.EC.8B.9C.EC.8A.A4.ED.85.9C_.EB.A0.88.EC.9D.B4.EB.B8.94.EA.B3.BC_UUID)
  - [1.2분할 영역 레이블 및 UUID](#.EB.B6.84.ED.95.A0_.EC.98.81.EC.97.AD_.EB.A0.88.EC.9D.B4.EB.B8.94_.EB.B0.8F_UUID)
  - [1.3fstab 정보](#fstab_.EC.A0.95.EB.B3.B4)
  - [1.4fstab 파일 만들기](#fstab_.ED.8C.8C.EC.9D.BC_.EB.A7.8C.EB.93.A4.EA.B8.B0)
    - [1.4.1DOS/Legacy BIOS systems](#DOS.2FLegacy_BIOS_systems)
    - [1.4.2DPS UEFI PARTUUID](#DPS_UEFI_PARTUUID)
- [2네트워크 정보](#.EB.84.A4.ED.8A.B8.EC.9B.8C.ED.81.AC_.EC.A0.95.EB.B3.B4)
  - [2.1호스트, 도메인 정보](#.ED.98.B8.EC.8A.A4.ED.8A.B8.2C_.EB.8F.84.EB.A9.94.EC.9D.B8_.EC.A0.95.EB.B3.B4)
    - [2.1.1Set the hostname (OpenRC or systemd)](#Set_the_hostname_.28OpenRC_or_systemd.29)
    - [2.1.2systemd](#systemd)
  - [2.2Network](#Network)
    - [2.2.1DHCP via dhcpcd (any init system)](#DHCP_via_dhcpcd_.28any_init_system.29)
    - [2.2.2netifrc (OpenRC)](#netifrc_.28OpenRC.29)
  - [2.3네트워크 설정](#.EB.84.A4.ED.8A.B8.EC.9B.8C.ED.81.AC_.EC.84.A4.EC.A0.95)
  - [2.4부팅 과정에서 네트워크 자동으로 시작하기](#.EB.B6.80.ED.8C.85_.EA.B3.BC.EC.A0.95.EC.97.90.EC.84.9C_.EB.84.A4.ED.8A.B8.EC.9B.8C.ED.81.AC_.EC.9E.90.EB.8F.99.EC.9C.BC.EB.A1.9C_.EC.8B.9C.EC.9E.91.ED.95.98.EA.B8.B0)
  - [2.5hosts 파일](#hosts_.ED.8C.8C.EC.9D.BC)
- [3시스템 정보](#.EC.8B.9C.EC.8A.A4.ED.85.9C_.EC.A0.95.EB.B3.B4)
  - [3.1루트 암호](#.EB.A3.A8.ED.8A.B8_.EC.95.94.ED.98.B8)
  - [3.2Init와 부팅 설정](#Init.EC.99.80_.EB.B6.80.ED.8C.85_.EC.84.A4.EC.A0.95)
    - [3.2.1OpenRC](#OpenRC)
    - [3.2.2systemd](#systemd_2)

## 파일 시스템 정보

#### 파일 시스템 레이블과 UUID

MBR(BIOS)와 GPT 에는 _파일 시스템_ 레이블과 _파일 시스템_ UUID가 있습니다. 이 속성은 블록 장치를 찾아 마운드할 때 mount 명령에서 대신 사용하여 /etc/fstab에 지정할 수 있습니다. 파일 시스템 레이블과 UUID는 앞에 LABEL과 UUID를 붙여 식별 명칭을 부여하며, blkid 명령으로 확인할 수 있습니다:

`root #` `blkid`

**경고**

분할 영역의 파일 시스템을 날렸다면, 파일 시스템 레이블과 UUID 값도 바뀌거나 제거됩니다.

고유성을 확보하기 위해, MBR 방식 분할 영역 테이블을 사용하는 독자 여러분의 경우 /etc/fstab의 마운트 가능한 볼륨을 지정할 때는 레이블에 UUID를 사용하는 방식을 추천합니다.

**중요**

UUIDs of the filesystem on a LVM volume and its LVM snapshots are identical, therefore using UUIDs to mount LVM volumes should be avoided.

#### 분할 영역 레이블 및 UUID

GPT쪽으로 간 사용자는 /etc/fstab 에 분할 영역을 정의할 수 있는 '믿을 수 있는' 옵션을 더 넣습니다. 분할 영역 레이블과 분할 영역 UUID는 GPT 방식으로 포맷한 장치에 활용할 수 있는데 블록 장치의 개별 분할 영역을 분할 영역에 어떤 파일 시스템을 사용하든 상관 없이 유일하게 구별하려는 목적입니다. 분할 영역 레이블과 UUID는 PARTLABEL과 PARTUUID로 지정하며, 터미널에서 blkid 명령을 실행하면 간단하고 편리하게 깔끔한 모양새로 볼 수 있습니다:

Output for an **amd64** EFI system using the Discoverable Partition Specification UUIDs may like the following:

`root #` `blkid`

분할 영역 레이블에 대해서는 항상 그렇진 않지만, fstab에서 분할 영역을 UUID로 식별할 때는 개별 볼륨을 찾을 때, 심지어는 파일 시스템이 나중에 바뀌더라도 부트로더에서 햇갈리지 않게 합니다. fstab에 이전 방식으로 기본 블록 장치 파일(/dev/sd\*N)을 활용할 경우, 보통 SATA 블록 장치를 추가/제거할 경우 종종 시스템을 다시 시작하는데, 이때 위험 부담이 있습니다.

블록 장치 파일 작명은 시스템에 디스크가 어떻게 어떤 순서로 붙어있나 등의 요인에 따라 다릅니다. 또한 앞서 부팅 과정에서 커널이 어떤 장치를 먼저 찾았냐에 따라 다른 순서로 보여줄 수 있습니다. 디스크 장착 순서를 계속 다루지 않는 이상, 기존 상태로는 기본 블록 장치 파일을 활용하는게 간단하고 직관적인 접근 방식입니다.

### fstab 정보

리눅스 시스템에서 사용하는 모든 분할 영역 정보는 /etc/fstab 에 있습니다. 이 파일에는 분할 영역의 마운트 지점(파일 시스템 구조를 볼 수 있는 곳), 마운트해야 할 방법, 특수 옵션(자동인지 아닌지, 사용자가 마운트를 할 수 있는지 없는지 등)이 들어있습니다

### fstab 파일 만들기

**참고**

If the init system being used is systemd, the partition UUIDs conform to the Discoverable Partition Specification as given in [Preparing the disks](/wiki/Handbook:Alpha/Installation/Disks/ko "Handbook:Alpha/Installation/Disks/ko"), and the system uses UEFI, then creating an fstab can be skipped, since systemd auto-mounts partitions that follow the spec.

/etc/fstab 파일은 표와 비슷한 문법을 사용합니다. 각 줄은 6개의 내용으로 채워져있으며 공백(단일 공백, 탭 또는 혼합)문자로 나눕니다. 각각의 필드는 자체적인 의미를 지니고 있습니다:

1. 첫번째 내용은 마운트할 블록 특수 장치 또는 원격 파일 시스템을 나타냅니다. 대부분의 장치 식별자는 장치 파일 경로, 파일 시스템 레이블과 UUID, 분할 영역 레이블과 UUID와 같은 식으로 블록 특수 장치 노드에서 사용할 수 있습니다.
2. 두번째 내용은 분할 영역을 마운트할 마운트 지점을 나타냅니다
3. 세번째 내용은 분할 영역에서 사용하는 파일 시스템을 나타냅니다
4. 네번째 내용은 분할 영역을 마운트할 때 mount에서 사용하는 마운트 옵션을 나타냅니다. 모든 분할 영역에는 자체 마운트 옵션이 있기에 옵션 전체 목록을 알아보려면 mount 맨 페이지(man mount)를 읽어보시는게 좋겠습니다. 여러가지 마운트 옵션은 콤마로 구분합니다.
5. 다섯번째 내용은 분할 영역 덤프를 남겨둔지 여부를 나타냅니다. 보통 0(영)으로 남겨둡니다.
6. 여섯번째 내용은 시스템을 제대로 된 과정을 거쳐 끄지 못했을 때 fsck에서 파일 시스템을 점검할 순서를 나타냅니다. 루트 파일 시스템은 1이어야 하며 나머지는 2가 되어야 합니다(점검이 필요하지 않다면 0으로 남겨둡니다).

**중요**

젠투에서 제공하는 기본 /etc/fstab 파일은 올바른 fstab 파일이 아니지만, 양식 내용이 좀 더 자세하게 들어있습니다.

`root #` `nano -w /etc/fstab`

#### DOS/Legacy BIOS systems

/boot/ 분할 영역에 대한 옵션을 적어가는 방법을 살펴보도록 하겠습니다. 단지 예제일 뿐이며 설치 과정에서 결정한 공간 분할 형태에 따라 바꾸어야합니다.

alpha 분할 영역 예제에서, /boot/는 보통 ext2 파일 시스템을 쓰는 /dev/sda1 분할 영역입니다. 부팅 과정에 검사해야 하기 때문에 다음의 대용으로 적어내려가겠습니다:

파일 **`/etc/fstab`** **/etc/fstab의 /boot 줄 예제**

```
/dev/sda1 /boot ext2 defaults 0 2

```

어떤 사용자는 시스템의 보안을 개선하려는 이유로 /boot/ 분할 영역을 자동으로 마운트하려 하지 않습니다. 이러한 사용자는 _defaults_ 를 _noauto_ 로 바꾸어야합니다. 이 옵션은 해당 분할 영역을 사용하려고 할 때마다 직접 마운트해야 함을 의미합니다.

이전에 결정한 분할 영역 모양새에 따라 규칙을 추가하시고, CD-ROM 드라이브라든지, 물론, 다른 분할 영역과 드라이브도 사용한다면 해당 장치도 추가하십시오.

좀 더 내용을 추가한 /etc/fstab 파일 예제는 다음과 같습니다:

파일 **`/etc/fstab`** **/etc/fstab 전체 예제**

```
/dev/sda1   /boot        ext2    defaults,noatime     0 2
/dev/sda2   none         swap    sw                   0 0
/dev/sda3   /            ext4    noatime              0 1

/dev/cdrom  /mnt/cdrom   auto    noauto,user          0 0

```

/dev/cdrom /mnt/cdrom auto noauto,user 0 0
}}

#### DPS UEFI PARTUUID

Below is an example of an /etc/fstab file for a disk formatted with a GPT disklabel and Discoverable Partition Specification (DPS) UUIDs set for UEFI firmware:

파일 **`/etc/fstab`** **GPT disklabel DPS PARTUUID fstab example**

```
# Adjust any formatting difference and additional partitions created from the "Preparing the disks" step.
# This example shows a GPT disklabel with Discoverable Partition Specification (DSP) UUID set:
PARTUUID=c12a7328-f81f-11d2-ba4b-00a0c93ec93b                                  0 2
PARTUUID=0657fd6d-a4ab-43c4-84e5-0933c84b4f4f   none            sw                           0 0
PARTUUID=6523f8ae-3eb1-4e2a-a05a-18b695ae656f   /           xfs    defaults,noatime              0 1

```

세번째 내용에 _auto_ 를 사용하면 mount 명령에서 파일시스템에 처리해야 할 방식을 짐작하여 처리합니다. 여러 파일 시스템을 만들 수 있는 이동식 미디어에 추천합니다. 네번째에 `user` 를 넣으면 비-루트 사용자도 CD에 마운트할 수 있습니다.

To improve performance, most users would want to add the `noatime` mount option, which results in a faster system since access times are not registered (those are not needed generally anyway). This is also recommended for systems with solid state drives (SSDs). Users may wish to consider `lazytime` instead.

**요령**

Due to degradation in performance, defining the `discard` mount option in /etc/fstab is not recommended. It is generally better to schedule block discards on a periodic basis using a job scheduler such as cron or a timer (systemd). See [Periodic fstrim jobs](/wiki/SSD#Periodic_fstrim_jobs "SSD") for more information.

/etc/fstab 파일을 다시 확인하시고, 저장하고 빠져나가서 다음 단계를 계속 진행하십시오.

## 네트워크 정보

It is important to note the following sections are provided to help the reader quickly setup their system to partake in a local area network.

For systems running OpenRC, a more detailed reference for network setup is available in the [advanced network configuration](/wiki/Handbook:Alpha/Networking/Introduction/ko "Handbook:Alpha/Networking/Introduction/ko") section, which is covered near the end of the handbook. Systems with more specific network needs may need to skip ahead, then return here to continue with the rest of the installation.

For more specific systemd network setup, please review see the [networking portion](/wiki/Systemd/ko#Network "Systemd/ko") of the [systemd](/wiki/Systemd/ko "Systemd/ko") article.

### 호스트, 도메인 정보

사용자가 처리할 수 있는 선택중 하나는 PC의 이름을 부여하는 것입니다. 조금 쉬워보이긴 합니다만 리눅스 PC에 적당한 이름을 찾기란 대부분의 사용자에겐 어렵습니다. 빨리 넘어가기 위해 이 결정이 끝이 아님을 알아두십시오. 나중에 바꿀 수 있습니다. 아래 예제에서는 "homenetwork" 도메인에서 "tux" 호스트이름을 사용함을 보여줍니다.

#### Set the hostname (OpenRC or systemd)

`root #` `echo tux > /etc/hostname`

#### systemd

To set the system hostname for a system currently _running_ systemd, the hostnamectl utility may be used. During the installation process however, [systemd-firstboot](/wiki/Handbook:Alpha/Installation/System/ko#Init_and_boot_configuration_systemd "Handbook:Alpha/Installation/System/ko") command must be used instead (see later on in handbook).

For setting the hostname to "tux", one would run:

`root #` `hostnamectl hostname tux`

View help by running hostnamectl --help or man 1 hostnamectl.

### Network

There are _many_ options available for configuring network interfaces. This section covers a only a few methods. Choose the one which seems best suited to the setup needed.

#### DHCP via dhcpcd (any init system)

Most LAN networks operate a DHCP server. If this is the case, then using the dhcpcd program to obtain an IP address is recommended.

To install:

`root #` `emerge --ask net-misc/dhcpcd`

To enable and then start the service on OpenRC systems:

`root #` `rc-update add dhcpcd default
`

`root #` `rc-service dhcpcd start
`

To enable the service on systemd systems:

`root #` `systemctl enable dhcpcd`

With these steps completed, next time the system boots, dhcpcd should obtain an IP address from the DHCP server. See the [Dhcpcd](/wiki/Dhcpcd/ko "Dhcpcd/ko") article for more details.

#### netifrc (OpenRC)

**요령**

This is one particular way of setting up the network using [Netifrc](/wiki/Netifrc "Netifrc") on OpenRC. Other methods exist for simpler setups like [Dhcpcd](/wiki/Dhcpcd/ko "Dhcpcd/ko").

### 네트워크 설정

젠투 리눅스 설치 과정에서 네트워크를 거의 설정했습니다만 설치 CD 자체에서의 설정이었으며 설치한 환경에 대한 설정은 아니었습니다. 이제 네트워크 설정을 설치한 젠투 리눅스 시스템에 만들겠습니다.

**참고**

본딩, 브릿징, 802.1Q VLAN, 무선 네트워크를 다루는 네트워크에 대한 자세한 내용은 젠투 네트워크 설정 절에서 다룹니다.

모든 네트워크 정보를 /etc/conf.d/net에서 가져왔습니다. 별로 직관적이지 않은것 같은 간단한 문법을 사용합니다. 그러나 두려워 하실 일이 없습니다. 모든 내용은 아래에 설명해드립니다. 여러가지 설정을 다루는 자세한 설명 예제는 /usr/share/doc/netifrc-\*/net.example.bz2에 있습니다.

먼저 [net-misc/netifrc](https://packages.gentoo.org/packages/net-misc/netifrc)를 설치하십시오:

`root #` `emerge --ask --noreplace net-misc/netifrc`

기본적으로 DHCP를 사용합니다. DHCP를 동작하게 하려면 DHCP 클라이언트를 설치해야합니다. 필요한 시스템 도구 설치 편에서 설명하겠습니다.

DHCP 별개 옵션 때문이거나, DHCP를 모든 컴퓨터에서 쓸 수 있는게 아니어서 네트워크 연결을 설정해야한다면 /etc/conf.d/net 파일을 여십시오:

`root #` `nano -w /etc/conf.d/net`

config\_eth0 와 routes\_eth0 변수에 IP 주소 정보와 라우팅 정보를 입력하여 설정하십시오:

**참고**

여기서는 eth0 네트워크 인터페이스로 가정합니다. 그러나 이 이름이 시스템에 따라 다릅니다. 설치 매체가 최근의 것이라면 설치 매체로 부팅했을 때 나타나는 인터페이스 이름과 동일하다고 가정함을 추천합니다.

파일 **`/etc/conf.d/net`** **고정 IP 정의**

```
config_eth0="192.168.0.2 netmask 255.255.255.0 brd 192.168.0.255"
routes_eth0="default via 192.168.0.1"

```

DHCP를 사용하려면, config\_eth0를 정의하십시오:

파일 **`/etc/conf.d/net`** **DHCP 정의**

```
config_eth0="dhcp"

```

존재하는 모든 옵션을 보려면 /usr/share/doc/netifrc-\*/net.example.bz2를 읽어보십시오. DHCP 옵션을 설정해야 한다면 DHCP 클라이언트 맨 페이지도 읽어보십시오.

시스템에 여러가지 네트워크 인터페이스를 달고 있다면, config\_eth1, config\_eth2 등에 대해 위 과정을 반복하십시오.

이제 설정을 저장하고 빠져나간 후 다음 과정으로 계속 진행하십시오.

### 부팅 과정에서 네트워크 자동으로 시작하기

부팅 과정에서 네트워크 인터페이스를 활성화하려면, 기본 실행 레벨에 추가해야합니다.

`root #` `cd /etc/init.d
`

`root #` `ln -s net.lo net.eth0
`

`root #` `rc-update add net.eth0 default`

시스템에 여러가지 네트워크 인터페이스가 있다면 net.eth0 처럼 적당한 net.\* 파일을 만들어야합니다.

다음에 시스템을 부팅하면 네트워크 인터페이스 이름(현재 문서에 남긴 이름은 `eth0`)의 가정이 틀렸음을 알아챌 것입니다. 이 문제를 바로잡으려면 다음 단계를 따라 처리하십시오:

1. ( `eth0` 대신 `enp3s0` 같이) 올바른 인터페이스 이름으로 /etc/conf.d/net 파일을 업데이트하십시오
2. 새 심볼릭 링크를 만드십시오(/etc/init.d/net.enp3s0)
3. 이전 심볼릭 링크를 제거하십시오(rm /etc/init.d/net.eth0)
4. 새로 만든 심볼릭 링크를 기본 실행 레벨에 추가하십시오
5. 이전 심볼릭 링크를 rc-update del net.eth0 default 명령으로 제거하십시오

### hosts 파일

다음은 네트워크 환경을 리눅스에 알려야 합니다. /etc/hosts 에서 정의하며, 이름 서버에서 해석할 수 없는 호스트에서 호스트의 이름을 IP 주소로 바꾸는 과정을 돕습니다.

`root #` `nano -w /etc/hosts`

파일 **`/etc/hosts`** **네트워크 정보 채우기**

```
# This defines the current system and must be set
127.0.0.1     tux.homenetwork tux localhost

# Optional definition of extra systems on the network
192.168.0.5   jenny.homenetwork jenny
192.168.0.6   benny.homenetwork benny

```

1. Optional definition of extra systems on the network

192.168.0.5 jenny.homenetwork jenny
192.168.0.6 benny.homenetwork benny
}}

편집기에서 저장하고 빠져나가서 다음 과정을 계속 진행하십시오.

## 시스템 정보

### 루트 암호

passwd 명령으로 루트 암호를 설정하십시오.

`root #` `passwd`

루트 리눅스 계정은 가장 강력한 계정이므로 강력한 암호를 선택해야 합니다. 나중에 매일 사용할 일반 사용자 계정을 추가로 만듭니다.

### Init와 부팅 설정

#### OpenRC

(최소한 OpenRC를 쓸 때)젠투에서 서비스와 시스템의 시작과 마침 과정을 설정할 때 /etc/rc.conf 파일을 활용합니다. /etc/rc.conf 파일을 열고 파일의 모든 주석을 맘대로 제거하십시오. 필요한 부분의 설정을 다시 살펴보고 바꾸십시오.

`root #` `nano -w /etc/rc.conf`

그 다음 /etc/conf.d/keymaps 파일을 열어 키보드 설정을 처리하십시오. 해당 파일을 편집하여 올바른 키보드를 선택하고 설정하십시오.

`root #` `nano -w /etc/conf.d/keymaps`

keymap 변수는 특히 조심스럽게 다루십시오. 잘못된 키맵을 선택하면 키보드로 입력할 때, 이상한 결과가 나타납니다.

마지막으로 시계 옵션을 설정하려 /etc/conf.d/hwclock 파일을 편집하겠습니다. 개인 취향에 맞춰 편집하십시오.

`root #` `nano -w /etc/conf.d/hwclock`

하드웨어 클록에서 UTC 방식을 사용하지 않는다면, 파일에 `clock="local"` 를 설정해야 합니다. 그렇지 않으면 시스템의 시계 동작이 꼬이는 일이 생깁니다.

#### systemd

First, it is recommended to run systemd-machine-id-setup and then systemd-firstboot which will prepare various components of the system are set correctly for the first boot into the new systemd environment. The passing the following options will include a prompt for the user to set a locale, timezone, hostname, root password, and root shell values. It will also assign a random machine ID to the installation:

`root #` `systemd-machine-id-setup`

`root #` `systemd-firstboot --prompt`

Next users should run systemctl to reset all installed unit files to the preset policy values:

**참고**

It is possible that a failure warning will print out after running the following command. This is normal, as Gentoo's LiveCD uses OpenRC.

`root #` `systemctl preset-all --preset-mode=enable-only`

It's possible to run the full preset changes but this may reset any services which were already configured during the process:

`root #` `systemctl preset-all`

These two steps will help ensure a smooth transition from the live environment to the installation's first boot.

[← 리눅스 커널 설정](/wiki/Handbook:Alpha/Installation/Kernel/ko "이전 내용") [처음](/wiki/Handbook:Alpha/ko "Handbook:Alpha/ko") [시스템 도구 설치 →](/wiki/Handbook:Alpha/Installation/Tools/ko "다음 내용")