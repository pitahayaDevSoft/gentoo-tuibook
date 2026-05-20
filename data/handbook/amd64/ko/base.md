# Base

Other languages:

- [Deutsch](/wiki/Handbook:AMD64/Installation/Base/de "Handbook:AMD64/Installation/Base/de (100% translated)")
- [English](/wiki/Handbook:AMD64/Installation/Base "Handbook:AMD64/Installation/Base (100% translated)")
- [español](/wiki/Handbook:AMD64/Installation/Base/es "Handbook:AMD64/Installation/Base/es (100% translated)")
- [français](/wiki/Handbook:AMD64/Installation/Base/fr "Handbook:AMD64/Installation/Base/fr (100% translated)")
- [italiano](/wiki/Handbook:AMD64/Installation/Base/it "Handbook:AMD64/Installation/Base/it (100% translated)")
- [magyar](/wiki/Handbook:AMD64/Installation/Base/hu "Handbook:AMD64/Installation/Base/hu (100% translated)")
- [polski](/wiki/Handbook:AMD64/Installation/Base/pl "Handbook:AMD64/Installation/Base/pl (100% translated)")
- [português do Brasil](/wiki/Handbook:AMD64/Installation/Base/pt-br "Handbook:AMD64/Installation/Base/pt-br (100% translated)")
- [čeština](/wiki/Handbook:AMD64/Installation/Base/cs "Handbook:AMD64/Installation/Base/cs (100% translated)")
- [русский](/wiki/Handbook:AMD64/Installation/Base/ru "Handbook:AMD64/Installation/Base/ru (100% translated)")
- [العربية](/wiki/Handbook:AMD64/Installation/Base/ar "Handbook:AMD64/Installation/Base/ar (100% translated)")
- [தமிழ்](/wiki/Handbook:AMD64/Installation/Base/ta "Handbook:AMD64/Installation/Base/ta (100% translated)")
- [中文](/wiki/Handbook:AMD64/Installation/Base/zh "Handbook:AMD64/Installation/Base/zh (100% translated)")
- [中文（中国大陆）‎](/wiki/Handbook:AMD64/Installation/Base/zh-cn "Handbook:AMD64/Installation/Base/zh-cn (100% translated)")
- [日本語](/wiki/Handbook:AMD64/Installation/Base/ja "Handbook:AMD64/Installation/Base/ja (100% translated)")
- 한국어

[AMD64 핸드북](/wiki/Handbook:AMD64/ko "Handbook:AMD64/ko")[설치](/wiki/Handbook:AMD64/Full/Installation/ko "Handbook:AMD64/Full/Installation/ko")[설치 정보](/wiki/Handbook:AMD64/Installation/About/ko "Handbook:AMD64/Installation/About/ko")[매체 선택](/wiki/Handbook:AMD64/Installation/Media/ko "Handbook:AMD64/Installation/Media/ko")[네트워크 설정](/wiki/Handbook:AMD64/Installation/Networking/ko "Handbook:AMD64/Installation/Networking/ko")[디스크 준비](/wiki/Handbook:AMD64/Installation/Disks/ko "Handbook:AMD64/Installation/Disks/ko")[스테이지 3 설치](/wiki/Handbook:AMD64/Installation/Stage/ko "Handbook:AMD64/Installation/Stage/ko")베이스 시스템 설치[커널 설정](/wiki/Handbook:AMD64/Installation/Kernel/ko "Handbook:AMD64/Installation/Kernel/ko")[시스템 설정](/wiki/Handbook:AMD64/Installation/System/ko "Handbook:AMD64/Installation/System/ko")[도구 설치](/wiki/Handbook:AMD64/Installation/Tools/ko "Handbook:AMD64/Installation/Tools/ko")[부트로더 설정](/wiki/Handbook:AMD64/Installation/Bootloader/ko "Handbook:AMD64/Installation/Bootloader/ko")[마무리](/wiki/Handbook:AMD64/Installation/Finalizing/ko "Handbook:AMD64/Installation/Finalizing/ko")[젠투 활용](/wiki/Handbook:AMD64/Full/Working/ko "Handbook:AMD64/Full/Working/ko")[포티지 소개](/wiki/Handbook:AMD64/Working/Portage/ko "Handbook:AMD64/Working/Portage/ko")[USE 플래그](/wiki/Handbook:AMD64/Working/USE/ko "Handbook:AMD64/Working/USE/ko")[포티지 기능](/wiki/Handbook:AMD64/Working/Features/ko "Handbook:AMD64/Working/Features/ko")[초기화 스크립트 시스템](/wiki/Handbook:AMD64/Working/Initscripts/ko "Handbook:AMD64/Working/Initscripts/ko")[환경 변수](/wiki/Handbook:AMD64/Working/EnvVar/ko "Handbook:AMD64/Working/EnvVar/ko")[포티지 활용](/wiki/Handbook:AMD64/Full/Portage/ko "Handbook:AMD64/Full/Portage/ko")[파일 및 디렉터리](/wiki/Handbook:AMD64/Portage/Files/ko "Handbook:AMD64/Portage/Files/ko")[변수](/wiki/Handbook:AMD64/Portage/Variables/ko "Handbook:AMD64/Portage/Variables/ko")[소프트웨어 브랜치 함께 사용하기](/wiki/Handbook:AMD64/Portage/Branches/ko "Handbook:AMD64/Portage/Branches/ko")[추가 도구](/wiki/Handbook:AMD64/Portage/Tools/ko "Handbook:AMD64/Portage/Tools/ko")[꾸러미 저장소 개별 설정](/wiki/Handbook:AMD64/Portage/CustomTree/ko "Handbook:AMD64/Portage/CustomTree/ko")[고급 기능](/wiki/Handbook:AMD64/Portage/Advanced/ko "Handbook:AMD64/Portage/Advanced/ko")[네트워크 설정](/wiki/Handbook:AMD64/Full/Networking/ko "Handbook:AMD64/Full/Networking/ko")[시작하기](/wiki/Handbook:AMD64/Networking/Introduction/ko "Handbook:AMD64/Networking/Introduction/ko")[고급 설정](/wiki/Handbook:AMD64/Networking/Advanced/ko "Handbook:AMD64/Networking/Advanced/ko")[모듈러 네트워크](/wiki/Handbook:AMD64/Networking/Modular/ko "Handbook:AMD64/Networking/Modular/ko")[무선 네트워크](/wiki/Handbook:AMD64/Networking/Wireless/ko "Handbook:AMD64/Networking/Wireless/ko")[기능 추가](/wiki/Handbook:AMD64/Networking/Extending/ko "Handbook:AMD64/Networking/Extending/ko")[동적 관리](/wiki/Handbook:AMD64/Networking/Dynamic/ko "Handbook:AMD64/Networking/Dynamic/ko")

## Contents

- [1루트 변경](#.EB.A3.A8.ED.8A.B8_.EB.B3.80.EA.B2.BD)
  - [1.1DNS 정보 복사](#DNS_.EC.A0.95.EB.B3.B4_.EB.B3.B5.EC.82.AC)
  - [1.2필요한 파일 시스템 마운트](#.ED.95.84.EC.9A.94.ED.95.9C_.ED.8C.8C.EC.9D.BC_.EC.8B.9C.EC.8A.A4.ED.85.9C_.EB.A7.88.EC.9A.B4.ED.8A.B8)
  - [1.3새 환경으로 진입](#.EC.83.88_.ED.99.98.EA.B2.BD.EC.9C.BC.EB.A1.9C_.EC.A7.84.EC.9E.85)
  - [1.4Preparing for a bootloader](#Preparing_for_a_bootloader)
    - [1.4.1UEFI systems](#UEFI_systems)
    - [1.4.2DOS/Legacy BIOS systems](#DOS.2FLegacy_BIOS_systems)
- [2포티지 설정](#.ED.8F.AC.ED.8B.B0.EC.A7.80_.EC.84.A4.EC.A0.95)
  - [2.1웹에서 이빌드 저장소 스냅샷 가져와서 설치하기](#.EC.9B.B9.EC.97.90.EC.84.9C_.EC.9D.B4.EB.B9.8C.EB.93.9C_.EC.A0.80.EC.9E.A5.EC.86.8C_.EC.8A.A4.EB.83.85.EC.83.B7_.EA.B0.80.EC.A0.B8.EC.99.80.EC.84.9C_.EC.84.A4.EC.B9.98.ED.95.98.EA.B8.B0)
  - [2.2선택: 미러 선택](#.EC.84.A0.ED.83.9D:_.EB.AF.B8.EB.9F.AC_.EC.84.A0.ED.83.9D)
    - [2.2.1Optional: rsync mirrors](#Optional:_rsync_mirrors)
  - [2.3선택: 젠투 이빌드 저장소 업데이트](#.EC.84.A0.ED.83.9D:_.EC.A0.A0.ED.88.AC_.EC.9D.B4.EB.B9.8C.EB.93.9C_.EC.A0.80.EC.9E.A5.EC.86.8C_.EC.97.85.EB.8D.B0.EC.9D.B4.ED.8A.B8)
  - [2.4뉴스 항목 보기](#.EB.89.B4.EC.8A.A4_.ED.95.AD.EB.AA.A9_.EB.B3.B4.EA.B8.B0)
  - [2.5적절한 프로파일 선택](#.EC.A0.81.EC.A0.88.ED.95.9C_.ED.94.84.EB.A1.9C.ED.8C.8C.EC.9D.BC_.EC.84.A0.ED.83.9D)
    - [2.5.1No-multilib](#No-multilib)
  - [2.6Optional: Adding a binary package host](#Optional:_Adding_a_binary_package_host)
    - [2.6.1Repository configuration](#Repository_configuration)
    - [2.6.2Installing binary packages](#Installing_binary_packages)
  - [2.7USE 변수 설정](#USE_.EB.B3.80.EC.88.98_.EC.84.A4.EC.A0.95)
    - [2.7.1CPU\_FLAGS\_\*](#CPU_FLAGS_.2A)
    - [2.7.2VIDEO\_CARDS](#VIDEO_CARDS)
  - [2.8선택: ACCEPT\_LICENSE 변수 설정](#.EC.84.A0.ED.83.9D:_ACCEPT_LICENSE_.EB.B3.80.EC.88.98_.EC.84.A4.EC.A0.95)
  - [2.9@world 세트 업데이트](#.40world_.EC.84.B8.ED.8A.B8_.EC.97.85.EB.8D.B0.EC.9D.B4.ED.8A.B8)
    - [2.9.1Removing obsolete packages](#Removing_obsolete_packages)
- [3시간대](#.EC.8B.9C.EA.B0.84.EB.8C.80)
- [4로캘 설정](#.EB.A1.9C.EC.BA.98_.EC.84.A4.EC.A0.95)
  - [4.1Locale generation](#Locale_generation)
  - [4.2Locale selection](#Locale_selection)
- [5References](#References)

## 루트 변경

### DNS 정보 복사

새 환경에 들어가기 전 아직 남은 하나는 /etc/resolv.conf의 DNS 정보를 복사하는 일입니다. 새 환경에 들어가고 나서 네트워크가 그대로 동작할 수 있게 하려면 꼭 필요합니다. /etc/resolv.conf 파일에는 네트워크를 사용할 때 활용하는 네임 서버 주소가 들어있습니다.

이 정보를 복사하려면 cp 명령에 `--dereference` 옵션을 전달하는게 좋습니다. /etc/resolv.conf 파일이 심볼릭 링크라면 심볼릭 링크가 아니라 링크의 대상 파일 그 자체를 찾아서 복사합니다. 그렇지 않으면 새 환경에서 심볼릭 링크로 남아있으며(링크 대상은 새 환경에 존재하지 않습니다), 실제 존재하지 않는 파일을 참조합니다.

`root #` `cp --dereference /etc/resolv.conf /mnt/gentoo/etc/`

### 필요한 파일 시스템 마운트

**요령**

If using Gentoo's install media, this step can be replaced with simply: arch-chroot /mnt/gentoo.

잠시 동안, 리눅스 루트는 새 위치로 바뀝니다. 새 환경이 제대로 동작하는지 보려면 각각의 파일 시스템을 활성화해야 합니다.

활성화해야 할 파일 시스템은 다음과 같습니다:

- 리눅스 커널에서 환경에 공개하려는 정보가 만든 의사 파일 시스템(일반 파일 같지만, 실제로는 동적으로 생성하는 파일 시스템) /proc/
- /proc/보다 구조가 잘 갖춰져있어 대체 용도로 쓸 수 있는 의사 파일 시스템 /sys/
- 리눅스 장치 관리자(보통 udev)가 일부 관리하는 일반 파일 시스템이며, 모든 장치 파일이 들어있는 /dev/

다른 두개의 파일 시스템은 바인드 마운트를 하는데 반해 /proc/ 위치는 /mnt/gentoo/proc/에 마운트합니다. 후자의 경우, /mnt/gentoo/proc/는 (말 그대로) 파일 시스템에 대한 새 마운트지만, /mnt/gentoo/sys/는 실제로 /sys/(동일한 파일 시스템에 대한 두번째 마운트 지점)이 된다는 의미입니다.

`root #` `mount --types proc /proc /mnt/gentoo/proc
`

`root #` `mount --rbind /sys /mnt/gentoo/sys
`

`root #` `mount --make-rslave /mnt/gentoo/sys
`

`root #` `mount --rbind /dev /mnt/gentoo/dev
`

`root #` `mount --make-rslave /mnt/gentoo/dev
`

**참고**

`--make-rslave` 동작은 설치 과정에서 나중에 systemd 지원 기능에 필요합니다.

**경고**

비 젠투 설치 매체를 사용할 경우 이 과정이 충분하지 않을 수도 있습니다. 일부 배포판은 /dev/shm 심볼릭 링크를 /run/shm/으로 만들지만 루트를 바꾸고 난 후에는 무효처리됩니다. /dev/shm/을 적당한 tmpfs로 마운트 하려면 다음 명령으로 문제를 처리할 수 있습니다:

`root #` `test -L /dev/shm && rm /dev/shm && mkdir /dev/shm
`

`root #` `mount --types tmpfs --options nosuid,nodev,noexec shm /dev/shm
`

또한 파일 모드를 1777로 설정했는지 확인하십시오:

`root #` ` chmod 1777 /dev/shm`

### 새 환경으로 진입

모든 파티션을 초기화 하고 기반 환경을 설치했으니, 새 설치 환경에 chroot로 들어갈 차례입니다. 현재 설치 환경의 세션의 루트(접근할 수 있는 최상위 환경)를 설치 시스템의 루트(초기화한 파티션)로 바꾼다는 의미입니다. 그래서 이름이 _change root_ 또는 _chroot_ 라고 합니다.

루트 위치 전환은 세 단계로 처리합니다:

1. chroot를 사용하여 루트 위치를 (설치 매체의)/에서 (파티션의) /mnt/gentoo/로 바꿉니다
2. 몇가지 설정(/etc/profile에 있음)을 `source` 명령으로 메모리에 다시 불러옵니다
3. chroot로 바꾼 환경임을 인지하기 위해 초기 프롬프트를 바꿉니다.

`root #` `chroot /mnt/gentoo /bin/bash
`

`root #` `source /etc/profile
`

`root #` `export PS1="(chroot) ${PS1}"`

이 때, 모든 동작을 새 젠투 리눅스 환경에서 바로 처리할 수 있습니다. 물론 끝나려면 한참 멀었지만, 설치 절이 여전히 아직도 남아있는 이유입니다!

**요령**

젠투 설치가 여기 어디에선가 멈췄다면 이 단계에서 설치를 '재개'할 수 있"어야"합니다. 디스크 영역을 다시 분할할 필요가 없습니다! 간단하게 [루트 분할 영역 마운트](/wiki/Handbook:AMD64/Installation/Disks/ko#Mounting_the_root_partition "Handbook:AMD64/Installation/Disks/ko") 를 진행하고, 작업 환경에 다시 들어오기 전 [#DNS 정보 복사](#DNS_.EC.A0.95.EB.B3.B4_.EB.B3.B5.EC.82.AC) 로 위 단계 진행을 시작하십시오. 이 방법은 부트로더 문제를 해결할 때도 쓸만합니다. 더 많은 내용은 [chroot](/wiki/Chroot "Chroot") 게시글에서 찾아보실 수 있습니다.

### Preparing for a bootloader

Now that the new environment has been entered, it is necessary to prepare the new environment for the bootloader. It will be important to have the correct partition mounted when it is time to install the bootloader.

#### UEFI systems

For UEFI systems, /dev/sda1 was formatted with the FAT32 filesystem and will be used as the EFI System Partition (ESP). Create a new /efi directory (if not yet created), and then mount ESP there:

`root #` `mount /dev/sda1 /efi
`

#### DOS/Legacy BIOS systems

For DOS/Legacy BIOS systems, the bootloader will be installed into the /boot directory, therefore mount as follows:

`root #` `mount /dev/sda1 /boot`

## 포티지 설정

### 웹에서 이빌드 저장소 스냅샷 가져와서 설치하기

다음은 이빌드 저장소 스냅샷을 설치할 차례입니다. 이 스냅샷에는 사용할 수 있는 프로그램 (설치) 이름, 시스템 관리자가 선택할 수 있는 프로파일, 꾸러미, 프로파일별 소식 항목 등이 들어있습니다.

제한적인 방화벽 환경에 있어 네트워크 대역폭 활용을 아낄 분들은 (스냅샷을 다운로드할 때 HTTP/FTP를 활용하므로) emerge-webrsync 사용을 권장합니다. 네트워크 또는 대역폭 제한이 없는 독자분들은 다음으로 신나게(!) 건너 뛸 수 있습니다.

이 과정을 통해 젠투 미러 중 한 곳에서 (매일 최신 내용으로 바뀌는) 최신 스냅샷을 가져와서 시스템에 설치합니다:

`root #` `emerge-webrsync`

**참고**

설치 과정 중 emerge-webrsync에서 /var/db/repos/gentoo/ 위치가 없는 문제를 보고합니다. 당연한 결과이며 걱정할 필요가 없습니다. 도구에서 해당 위치를 만듭니다.

이 시점부터는 포티지에서 각각의 추천 업데이트를 실행하라고 알려줍니다. 스테이지 파일을 통해 설치한 시스템 꾸러미는 새 버전이 존재하며, 저장소 스냅샷을 새로 설치하여 포티지가 새 꾸러미를 인식하기 때문입니다. 지금은 꾸러미 업데이트를 안전하게 무시할 수 있으며, 젠투 설치가 끝나기 전이라도 업데이트를 미룰 수 있습니다.

### 선택: 미러 선택

소스코드를 빨리 다운로드 하려면 빠른 미러를 선택하시는 것이 좋습니다. 포티지는 make.conf 파일의 GENTOO\_MIRRORS 변수에서 미러를 찾아보며 해당 변수에 들어간 미러를 활용합니다 젠투 미러 목록 및 시스템에서 물리적으로 가까운(대부분 이런 미러가 빠름) 미러(또는 복수의 미러) 를 검색할 수 있습니다. 그러나 우리에겐 필요한 미러를 선택할 때 멋진 인터페이스를 제공해 주는 mirrorselect 도구가 있습니다. 선택할 미러를 찾아보고 하나 이상의 미러를 `Spacebar` 키로 선택하면 됩니다.

A tool called mirrorselect provides a pretty text interface to more quickly query and select suitable mirrors. Just navigate to the mirrors of choice and press `Spacebar` to select one or more mirrors.

`root #` `mirrorselect -i -o >> /mnt/gentoo/etc/portage/make.conf`

Alternatively, a list of active mirrors are [available online](https://www.gentoo.org/downloads/mirrors/).

#### Optional: rsync mirrors

Gentoo also has many rsync mirrors which can speed up downloading the available package list using `emerge --sync` (explained in more detail later) by selecting a mirror closer that is geographically closer to the user.

`root #` `mkdir /etc/portage/repos.conf
`

`root #` `cp /usr/share/portage/config/repos.conf /etc/portage/repos.conf/gentoo.conf
`

Select a mirror close to the system's location from [https://www.gentoo.org/support/rsync-mirrors/](https://www.gentoo.org/support/rsync-mirrors/) and replace the sync-uri default mirror in /etc/portage/repos.conf/gentoo.conf with the desired mirror location.

파일 **`/etc/portage/repos.conf/gentoo.conf`** **UK-based rsync mirror example**

```
[DEFAULT]
main-repo = gentoo
</div>

<div lang="en" dir="ltr" class="mw-content-ltr">
[gentoo]
location = /var/db/repos/gentoo
sync-type = rsync
sync-uri = rsync://rsync.uk.gentoo.org/gentoo-portage/
auto-sync = yes
sync-rsync-verify-jobs = 1
sync-rsync-verify-metamanifest = yes
sync-rsync-verify-max-age = 3
sync-openpgp-key-path = /usr/share/openpgp-keys/gentoo-release.asc
sync-openpgp-keyserver = hkps://keys.gentoo.org
sync-openpgp-key-refresh-retry-count = 40
sync-openpgp-key-refresh-retry-overall-timeout = 1200
sync-openpgp-key-refresh-retry-delay-exp-base = 2
sync-openpgp-key-refresh-retry-delay-max = 60
sync-openpgp-key-refresh-retry-delay-mult = 4
sync-webrsync-verify-signature = yes
sync-git-verify-commit-signature = true

```

### 선택: 젠투 이빌드 저장소 업데이트

젠투 이빌드 저장소를 최신 버전으로 업데이트할 수 있습니다. 이전에  emerge-webrsync 명령은 상당히 최근의 포티지 스냅샷(보통 최근 24시간 까지)을 설치하기 때문에 분명히 말하자면 선택적인 동작입니다.

최근 꾸러미 업데이트(최대 한시간 동안)가 필요하다면, emerge --sync 명령을 사용하십시오. 이 명령은 젠투 이빌드 저장소(이전에 emerge-webrsync 명령으로 가져옴)를 최신 상태로 업데이트하는데 rsync 프로토콜을 사용합니다.

`root #` `emerge --sync`

몇가지 프레임 버퍼와 직렬 콘솔 같은 느린 터미널에서는, 처리 과정의 속도를 높이기 위해 `--quiet` 옵션을 사용하시는 것이 좋습니다:

`root #` `emerge --sync --quiet`

### 뉴스 항목 보기

젠투 이빌드 저장소를 시스템과 동기화 하면, 포티지에서 다음과 같은 메시지로 사용자에게 경고합니다:

` * IMPORTANT: 2 news items need reading for repository 'gentoo'.
`

` * Use eselect news to read news items.
`

뉴스 항목은 rsync 트리로 사용자에게 중요한 메시지를 강제로 전달하는 통신 매체를 제공하려 만들었습니다. 뉴스 항목을 관리하려면 eselect news를 사용하십시오. eselect 프로그램은 시스템에서 바뀐 항목 또는 시스템 전반 설정을 처리하는 일반 관리 인터페이스입니다. 이 경우 eselect는 `news` 모듈 사용을 요청합니다.

`news` 모듈에서 다음 동작을 주로 사용합니다:

- `list` 명령으로 표시할 뉴스 목록의 개요를 표시합니다
- `read` 명령으로 읽을 수 있는 뉴스 항목을 표시합니다
- `purge` 명령으로 이미 읽어서 더 이상 읽을 일이 없는 뉴스 항목을 제거할 수 있습니다

`root #` `eselect news list
`

`root #` `eselect news read`

뉴스 리더에서 사용할 수 있는 기능이 무엇인지 더 살펴보려면 설명서 페이지를 참고하십시오:

`root #` `man news.eselect`

### 적절한 프로파일 선택

**요령**

Desktop profiles are not exclusively for _desktop environments_. They are also suitable for minimal window managers like i3 or sway.

_프로파일_ 이란 젠투 시스템의 구성요소입니다. USE, CFLAGS 등 중요한 변수 값의 기본값만을 지정하는 것이 아니라 꾸러미 버전 범위를 시스템에 고정합니다. 이 설정 데이터는 젠투 포티지 개발자가 관리합니다.

현재 시스템에서 활용하는 프로파일을 eselect로 볼 수 있으며, 이제 `profile` 모듈을 사용해보면:

**요령**

On an install media without a scroll-able terminal, `eselect profile list | less` can provide an easy way to list all available profiles while providing the ability to scroll.

`root #` `eselect profile list`

```
Available profile symlink targets:
  [1]   default/linux/amd64/23.0 *
  [2]   default/linux/amd64/23.0/desktop
  [3]   default/linux/amd64/23.0/desktop/gnome
  [4]   default/linux/amd64/23.0/desktop/kde

```

**참고**

명령 출력 결과는 예제일 뿐이며, 언제든 바뀝니다.

보신 바와 같이, 몇가지 데스크톱에 대한 데스크톱 하위 프로파일이 있습니다.

**경고**

프로파일 업그레이드를 가벼이 여기면 안됩니다. 초기 프로파일을 선택할 때, 스테이지3에서 처음 사용하는 **동일한 버전**(이를테면 17.0)에 해당하는 프로파일을 활용하는지 확인하십시오. 각 새 프로파일 버전은 이전 절차 내용을 담은 뉴스 항목으로 공지합니다. 새 프로파일로 전환하기 전에 해당 내용을 반드시 확인하고 따르십시오.

amd64 아키텍처에서 존재하는 프로파일을 확인한 후 사용자는 시스템의 다른 프로파일을 선택할 수 있습니다:

`root #` `eselect profile set 2`

#### No-multilib

32비트 프로그램 또는 라이브러리를 뺀 순수한 64비트 환경을 선택하려면, no-multilib 프로파일을 활용하십시오:

`root #` `eselect profile list`

```
Available profile symlink targets:
  [1]   default/linux/amd64/13.0 *
  [2]   default/linux/amd64/13.0/desktop
  [3]   default/linux/amd64/13.0/desktop/gnome
  [4]   default/linux/amd64/13.0/desktop/kde
  [5]   default/linux/amd64/13.0/no-multilib

```

다음 _no-multilib_ 프로파일을 선택하십시오:

`root #` `eselect profile set 5
`

`root #` `eselect profile list`

```
Available profile symlink targets:
  [1]   default/linux/amd64/13.0
  [2]   default/linux/amd64/13.0/desktop
  [3]   default/linux/amd64/13.0/desktop/gnome
  [4]   default/linux/amd64/13.0/desktop/kde
  [5]   default/linux/amd64/13.0/no-multilib *

```

**참고**

`developer` 하위 프로파일은 젠투 리눅스 개발 용도로 사용하며, 일반 사용자가 사용한다는 의미가 아닙니다.

### Optional: Adding a binary package host

Since December 2023, Gentoo's [Release Engineering team](/wiki/Project:RelEng "Project:RelEng") has offered an [official binary package host](/wiki/Binary_package_quickstart "Binary package quickstart") (colloquially shorted to just "binhost") for use by the general community to retrieve and install binary packages (binpkgs).[\[1\]](#cite_note-1)

Adding a binary package host allows Portage to install cryptographically signed, compiled packages. In many cases, adding a binary package host will _greatly_ decrease the mean time to package installation and adds much benefit when running Gentoo on older, slower, or low power systems.

#### Repository configuration

The repository configuration for a binhost is found in Portage's /etc/portage/binrepos.conf/ directory, which functions similarly to the configuration mentioned in the [Gentoo ebuild repository](/wiki/Handbook:AMD64/Installation/Base/ko#Gentoo_ebuild_repository "Handbook:AMD64/Installation/Base/ko") section.

When defining a binary host, there are two important aspects to consider:

1. The architecture and profile targets within the `sync-uri` value _do_ matter and should align to the respective computer architecture ( **amd64** in this case) and system profile selected in the [Choosing the right profile](/wiki/Handbook:AMD64/Installation/Base/ko#Choosing_the_right_profile "Handbook:AMD64/Installation/Base/ko") section.
2. Selecting a fast, geographically close mirror will generally shorten retrieval time. Review the mirrorselect tool mentioned in the [Optional: Selecting mirrors](/wiki/Handbook:AMD64/Installation/Base/ko#Gentoo_ebuild_repository "Handbook:AMD64/Installation/Base/ko") section or review the [online list of mirrors](https://www.gentoo.org/downloads/mirrors/) where URL values can be discovered.


파일 **`/etc/portage/binrepos.conf/gentoo.conf`** **CDN-based binary package host example**

```
[gentoo]
priority = 9959
# NOTE: Must adjust <arch> and <variant> as appropriate!
sync-uri = https://distfiles.gentoo.org/releases/<arch>/binpackages/<variant>
# x86-64 example sync-uri
# sync-uri = https://distfiles.gentoo.org/releases/amd64/binpackages/23.0/x86-64/
</div>

<div lang="en" dir="ltr" class="mw-content-ltr">
# Introduced in portage-3.0.74 for per-repo verification choices
verify-signature = true
# Default value with >=portage-3.0.77
location = /var/cache/binhost/gentoo

```

If the CPU supports [x86-64-v3](https://en.wikipedia.org/wiki/X86-64#Microarchitecture_levels) then the binhost offers binpkgs which are compiled to support these features as well.

For ease of explaining all main line Intel, Haswell and newer support this and same for AMD's Ryzen range.
For other CPU lines please check or just use binhost example given above.

파일 **`/etc/portage/binrepos.conf/gentoo.conf`** **CDN-based binary package x86-64-v3 host example**

```
[gentoo-x86-64-v3]
priority = 9999
sync-uri = https://distfiles.gentoo.org/releases/amd64/binpackages/23.0/x86-64-v3/
# Introduced in portage-3.0.74 for per-repo verification choices
verify-signature = true
# Default value with >=portage-3.0.77
location = /var/cache/binhost/gentoo-x86-64-v3

```

**중요**

If using the x86-64-v3 binhost then is important to add both examples to /etc/portage/binrepos.conf/gentoo.conf as sometimes a package might not yet be available in a x86-64-v3 variant.

#### Installing binary packages

Portage will compile packages from code source by default. It can be instructed to use binary packages in the following ways:

1. The `--getbinpkg` option can be passed when invoking the emerge command. This method of binary package installation is useful to install only a particular binary package.
2. Changing the system's default via Portage's FEATURES variable, which is exposed through the /etc/portage/make.conf file. Applying this configuration change will cause Portage to query the binary package host for the package(s) to be requested and fall back to compiling locally when no results are found.

For example, to have Portage always install available binary packages:

파일 **`/etc/portage/make.conf`** **Configure Portage to use binary packages by default**

```
# Appending getbinpkg to the list of values within the FEATURES variable
FEATURES="${FEATURES} getbinpkg"
# Require signatures
FEATURES="${FEATURES} binpkg-request-signature"

```

Please also run getuto for Portage to set up the necessary keyring for verification:

`root #` `getuto`

Additional Portage features will be discussed in the [the next chapter](/wiki/Handbook:AMD64/Working/Features/ko#Portage_features "Handbook:AMD64/Working/Features/ko") of the handbook.

### USE 변수 설정

USE는 젠투가 사용자에게 제공하는 가장 강력한 변수중 하나입니다. 여러 프로그램 각 항목을 추가로 지원하든 안하든 컴파일할 수 있습니다. 예를 들어 어떤 프로그램은 GTK+ 지원 또는 Qt 지원을 넣고 컴파일할 수 있습니다. 다른 프로그램은 SSL 지원을 빼고 컴파일할 수 있습니다. 어떤 프로그램은 X11 지원(X-서버) 대신 프레임버퍼 지원(svgalib)을 빼고도 컴파일할 수 있습니다.

대부분의 배포판에서는 가능한한 최대한의 지원을 포함하여 꾸러미를 컴파일합니다. 상당한 양의 의존성에 상관 없이 프로그램의 크기와 시작 시간이 늘어납니다. 젠투 사용자는 컴파일할 때 어떤 옵션을 넣을지 지정할 수 있습니다. 이것이 바로 USE 변수가 동작하는 위치입니다.

USE 변수에는 컴파일 옵션에 매핑할 키워드가 들어있습니다. 예를 들어 `ssl` 은 SSL 지원을 프로그램에 넣어 프로그램에서 SSL 기능이 동작하도록 컴파일합니다. `-X` 는 X 서버 지원을 제거합닏(앞에 음수부호가 들어감에 주목). `gnome gtk -kde -qt4 -qt5` 는 시스템을 GNOME(아키텍처에서 지원한다면)에 완전히 맞추려 그놈(및 GTK+) 지원을 넣고 KDE(및 Qt) 지원을 뺍니다.

기본 USE 설정은 시스템에서 사용하는 젠투 프로파일의 make.defaults 파일에 있습니다. 젠투에서는 프로파일의 (복잡한) 계층 시스템을 사용하는데, 이 단계로는 깊이 들어가지 않겠습니다. 현재 활성화한 USE 설정을 확인하는 가장 쉬운 방법은 emerge --info를 실행하고 "USE"로 시작하는 줄을 선택해서 확인하는 방법입니다:

`root #` `emerge --info | grep ^USE`

```
USE="X acl alsa amd64 berkdb bindist bzip2 cli cracklib crypt cxx dri ..."

```

**참고**

위 예제는 잘렸으며, 실제 USE 값 설정 목록은 엄청 큽니다.

시스템에서 사용할 수 있는 USE 플래그의 전체 설명은 /var/db/repos/gentoo/profiles/use.desc에 있습니다.

`root #` `less /var/db/repos/gentoo/profiles/use.desc`

less 명령에서는 `↑`, `↓` 키로 스크롤할 수 있고, `q` 를 눌러 빠져나갈 수 있습니다.

예제를 통해 DVD, ALSA, CD 기록 기능을 지원하는 KDE 기반 시스템의 USE 플래그 설정을 보여드리겠습니다:

`root #` `nano -w /etc/portage/make.conf`

파일 **`/etc/portage/make.conf`** **KDE 기반 시스템에서 DVD, ALSA, CD 기록 기능을 포함하는 USE 플래그 활성화**

```
USE="-gtk -gnome qt4 qt5 kde dvd alsa cdr"

```

/etc/portage/make.conf에서 USE를 정의할 때, 기본 목록에서 _추가_ (또는 USE 플래그가 `-` 부호로 시작하는경우 _삭제_)됩니다. 기본 USE 설정을 무시하고 자체적으로 관리하려는 사용자는 make.conf의 USE 정의 앞부분에 `-*` 를 넣으십시오:

파일 **`/etc/portage/make.conf`** **기본 USE 플래그 무시**

```
USE="-* X acl alsa"

```

**경고**

(위 설정과 마찬가지로) `-*` 설정이 가능하다 하더라도, 기본 USE 플래그는 일부 이빌드의 설정 충돌을 막고 다른 오류가 일어나지 않게 심혈을 기울여 설정했으므로 권장하지 않습니다.

#### CPU\_FLAGS\_\*

Some architectures (including AMD64/X86, ARM, PPC) have a [USE\_EXPAND](/wiki/USE_EXPAND "USE EXPAND") variable called [CPU\_FLAGS\_<ARCH>](/wiki/CPU_FLAGS_X86 "CPU FLAGS X86"), where <ARCH> is replaced with the relevant system architecture name.

**중요**

Do not be confused! **AMD64** and **X86** systems share some common architecture, so the proper variable name for **AMD64** systems is CPU\_FLAGS\_X86.

This is used to configure the build to compile in specific assembly code or other intrinsics, usually hand-written or otherwise extra,
and is **not** the same as asking the compiler to output optimized code for a certain CPU feature (e.g. `-march=`).

Users should set this variable in addition to configuring their COMMON\_FLAGS as desired.

A few steps are needed to set this up:

`root #` `emerge --ask --oneshot app-portage/cpuid2cpuflags`

Inspect the output manually if curious:

`root #` `cpuid2cpuflags`

Then copy the output into package.use:

`root #` `echo "*/* $(cpuid2cpuflags)" > /etc/portage/package.use/00cpu-flags`

#### VIDEO\_CARDS

By default a profile already sets a few video cards. For many reasons this is not ideal, but a useful safety net.

To configure the system correctly the user needs to first unset the preset video cards with `VIDEO_CARDS: -*` then set the correct card for that system.

파일 **`/etc/portage/package.use/00video_cards`** **Example**

```
*/* VIDEO_CARDS: -* <GPU DRIVER NAME>

```

Below is a table that can be used to easily compare the different video card types to their respective `VIDEO_CARDS` value.

GPU
VIDEO\_CARDS
(Official) Nvidia Maxwell and newer`nvidia`Nouveau Nvidia [Supported list](/wiki/NVIDIA "NVIDIA")`nouveau`AMD since Sea Islands`amdgpu radeonsi`ATI and older AMDSee [radeon#Feature support](/wiki/Radeon#Feature_support "Radeon")Intel Nehalem and newer`intel`Intel Gen 2 and 3 [Supported list](/wiki/Intel#.23Feature_support.2Fko "Intel")`intel i915`QEMU/KVM`virgl`WSL`d3d12`

Below is a few examples of a properly set package.use for _VIDEO\_CARDS_:

파일 **`/etc/portage/package.use/00video_cards`** **AMD example**

```
*/* VIDEO_CARDS: -* amdgpu radeonsi

```

파일 **`/etc/portage/package.use/00video_cards`** **Intel example**

```
*/* VIDEO_CARDS: -* intel

```

파일 **`/etc/portage/package.use/00video_cards`** **Nvidia example**

```
*/* VIDEO_CARDS: -* nvidia

```

Details for various GPU(s) can be found at the [AMDGPU](/wiki/AMDGPU "AMDGPU"), [Intel](/wiki/Intel "Intel"), [Nouveau (Open Source)](/wiki/Nouveau/ko "Nouveau/ko"), or [NVIDIA (Proprietary)](/wiki/NVIDIA "NVIDIA") articles.

### 선택: ACCEPT\_LICENSE 변수 설정

Starting with Gentoo Linux Enhancement Proposal 23 (GLEP 23), a mechanism was created to allow system administrators the ability to "regulate the software they install with regards to licenses... Some want a system free of any software that is not OSI-approved; others are simply curious as to what licenses they are implicitly accepting."[\[2\]](#cite_note-2) With a motivation to have more granular control over the type of software running on a Gentoo system, the ACCEPT\_LICENSE variable was born.

젠투는 다음과 같이 프로파일에 기 지정 초기값을 보유하고 있습니다:

`user $` `portageq envvar ACCEPT_LICENSE`

```
@FREE
```

라이선스 그룹은 [젠투 라이선스 프로젝트](/wiki/Project:Licenses "Project:Licenses") 가 젠투 저장소를 관리하며 지정합니다:

그룹 이름설명
@GPL-COMPATIBLE자유 소프트웨어 재단에서 인가한 GPL 호환 라이선스 [\[a\_license 1\]](#cite_note-3)@FSF-APPROVED자유 소프트웨어 재단에서 인가한 라이선스 (@GPL-COMPATIBLE도 해당)
@OSI-APPROVED오픈 소스 이니셔티브(OSI)에서 인가한 라이선스 [\[a\_license 2\]](#cite_note-4)@MISC-FREE아마도 자유 소프트웨어 일지도 모르는 기타 라이선스, 예: 자유 소프트웨어 정의를 따름 [\[a\_license 3\]](#cite_note-5) 그러나 FSF 또는 OSI에서 인가하지 않음
@FREE-SOFTWARE@FSF-APPROVED, @OSI-APPROVED, @MISC-FREE를 합침
@FSF-APPROVED-OTHER"자유 문서"와 "프로그램 및 문서에서 활용하는 저작물"에 해당하는(글꼴 포함) 자유 소프트웨어 재단 인가 라이선스
@MISC-FREE-DOCS'자유' 정의를 따르는 자유 문서와 기타 저작물 (글꼴 포함) [\[a\_license 4\]](#cite_note-6) 에 대해 @FSF-APPROVED-OTHER 에 없는 라이선스
@FREE-DOCUMENTS@FSF-APPROVED-OTHER, @MISC-FREE-DOCS 를 합침
@FREE사용, 공유, 수정, 공유 수정물에 대한 모든 라이선스 모음. @FREE-SOFTWARE, @FREE-DOCUMENTS 를 합침
@BINARY-REDISTRIBUTABLE바이너리 형식 소프트웨어의 자유 재배포를 최소한 허용하는 라이선스. @FREE도 해당.
@EULA권리를 제한하는 라이선스 동의서. "all-rights-reserved" 보다 제한적이거나 분명한 허용이 필요함

Some common license groups include:

A list of software licenses grouped according to their kinds.
NameDescription
`@GPL-COMPATIBLE`GPL compatible licenses approved by the Free Software Foundation [\[a\_license 5\]](#cite_note-7)`@FSF-APPROVED`Free software licenses approved by the FSF (includes `@GPL-COMPATIBLE`)
`@OSI-APPROVED`Licenses approved by the Open Source Initiative [\[a\_license 6\]](#cite_note-8)`@MISC-FREE`Misc licenses that are probably free software, i.e. follow the Free Software Definition [\[a\_license 7\]](#cite_note-9) but are not approved by either FSF or OSI
`@FREE-SOFTWARE`Combines `@FSF-APPROVED`, `@OSI-APPROVED`, and `@MISC-FREE`.
`@FSF-APPROVED-OTHER`FSF-approved licenses for "free documentation" and "works of practical use besides software and documentation" (including fonts)
`@MISC-FREE-DOCS`Misc licenses for free documents and other works (including fonts) that follow the free definition [\[a\_license 8\]](#cite_note-10) but are NOT listed in `@FSF-APPROVED-OTHER`.
`@FREE-DOCUMENTS`Combines `@FSF-APPROVED-OTHER` and `@MISC-FREE-DOCS`.
`@FREE`Metaset of all licenses with the freedom to use, share, modify and share modifications. Combines `@FREE-SOFTWARE` and `@FREE-DOCUMENTS`.
`@BINARY-REDISTRIBUTABLE`Licenses that at least permit free redistribution of the software in binary form. Includes `@FREE`.
`@EULA`License agreements that try to take away your rights. These are more restrictive than "all-rights-reserved" or require explicit approval

1. [↑](#cite_ref-3)[https://www.gnu.org/licenses/license-list.html](https://www.gnu.org/licenses/license-list.html)
2. [↑](#cite_ref-4)[https://www.opensource.org/licenses](https://www.opensource.org/licenses)
3. [↑](#cite_ref-5)[https://www.gnu.org/philosophy/free-sw.html](https://www.gnu.org/philosophy/free-sw.html)
4. [↑](#cite_ref-6)[https://freedomdefined.org/](https://freedomdefined.org/)
5. [↑](#cite_ref-7)[https://www.gnu.org/licenses/license-list.html](https://www.gnu.org/licenses/license-list.html)
6. [↑](#cite_ref-8)[https://www.opensource.org/licenses](https://www.opensource.org/licenses)
7. [↑](#cite_ref-9)[https://www.gnu.org/philosophy/free-sw.html](https://www.gnu.org/philosophy/free-sw.html)
8. [↑](#cite_ref-10)[https://freedomdefined.org/](https://freedomdefined.org/)

Currently set system wide acceptable license values can be viewed via:

`user $` `portageq envvar ACCEPT_LICENSE`

```
@FREE
```

As visible in the output, the default value is to only allow software which has been grouped into the `@FREE` category to be installed.

Specific licenses or licenses groups for a system can be defined in the following locations:

- System wide within the selected profile - this sets the default value.
- System wide within the /etc/portage/make.conf file. System administrators override the profile's default value within this file.
- Per-package within a /etc/portage/package.license file.
- Per-package within a /etc/portage/package.license/ _directory_ of files.

/etc/portage/make.conf 파일을 수정하여 시스템 전체 영역 값을 조정할 수 있습니다. 기본 값은 자유 소프트웨어 재단, 오픈소스 이니셔티브에서 분명하게 인가했거나, 자유 소프트웨어 정의를 따르는 라이선스만을 따릅니다.:

파일 **`/etc/portage/make.conf`** **ACCEPT\_LICENSE 개별 설정**

```
ACCEPT_LICENSE="-* @FREE"

```

필요한 경우 다음과 같이 꾸러미 별로 설정할 수 있습니다:

파일 **`/etc/portage/package.license/kernel`** **라이선스 허용 예제**

```
app-arch/unrar unRAR
sys-kernel/linux-firmware @BINARY-REDISTRIBUTABLE
sys-firmware/intel-microcode intel-ucode

```

`root #` `mkdir /etc/portage/package.license`

Software license details for an individual Gentoo package are stored within the LICENSE variable of the associated ebuild. One package may have one or many software licenses, therefore it be necessary to specify multiple acceptable licenses for a single package.

파일 **`/etc/portage/package.license/kernel`** **Accepting licenses on a per-package basis**

```
app-arch/unrar unRAR
sys-kernel/linux-firmware linux-fw-redistributable
sys-firmware/intel-microcode intel-ucode

```

**중요**

ebuild의 LICENSE변수는 젠투 개발자와 사용자에게 필요한 지침일 뿐입니다. 법적 조항이 아니며, 실제로 어떻게 반영이 되는지는 보증하지 않습니다. 따라서 이 변수에 신경쓰실 일이 없으며, 여러분이 사용하려믄 모든 파일이 들어간 꾸러미 자체를 유심히 살펴보십시오.

### @world 세트 업데이트

**요령**

완전한 구성을 갖춘 데스크톱 환경 프로파일을 선택하면 설치 과정에 필요한 시간은 상당히 늘어날 수 있습니다. 진행 과정의 일은 '과정상 경험' 으로 처리할 수 있습니다. 짧은 프로파일 이름을 지닌, 드문 경우의 시스템 [@world 세트](/wiki/World_set_(Portage) "World set (Portage)") 가 있는데 이 프로파일은 시스템에 필요한 꾸러미 수가 적습니다. 다시 말해서:

- `default/linux/amd64/13.0` 을 선택하면 상당히 적은 꾸러미를 최신으로 유지합니다만,
- `default/linux/amd64/13.0/desktop/gnome/systemd` 는 OpenRC에서 Systemd로, 그놈 데스크톱 환경 프레임워크를 설치한 만큼 상당한 꾸러미를 설치해야합니다.

이 다음 단계는 stage3를 빌드한 후 어떤 프로파일을 선택하든지간에 시스템에서 최신 내용을 반영하도록 "필요한" 과정입니다:

1. A profile target _different_ from the stage file has been selected.
2. Additional USE flags have been set for installed packages.

Readers who are performing an 'install Gentoo speed run' may safely skip @world set updates until _after_ their system has rebooted into the new Gentoo environment.

Readers who are performing a slow run can have Portage perform updates for package, profile, and/or USE flag changes at the present time:

`root #` `emerge --ask --verbose --update --deep --newuse @world`

Readers who added a binary host [above](/wiki/Handbook:AMD64/Installation/Base/ko#Optional:_Adding_a_binary_package_host "Handbook:AMD64/Installation/Base/ko") can add --getbinpkg (or -g) in order to fetch packages from the binary host instead of compiling them:

`root #` `emerge --ask --verbose --update --deep --newuse --getbinpkg @world`

#### Removing obsolete packages

It is important to always _depclean_ after system upgrades to remove obsolete packages. Review the output carefully with emerge --depclean --pretend to see if any of the to-be-cleaned packages should be kept if personally using them. To keep a package which would otherwise be depcleaned, use emerge --noreplace foo.

`root #` `emerge --ask --pretend --depclean`

If happy, then proceed with a real depclean:

`root #` `emerge --ask --depclean`

## 시간대

**참고**

This step does not apply to users of the musl libc. Users who do not know what that means should perform this step.

기대한 대로의 시간대 영역을 나타내지 않는 /usr/share/zoneinfo/Etc/GMT\* 시간대 이름 사용을 피하십시오. GMT-8의 경우 실제로 GMT+8입니다.

시스템 시간대를 선택하십시오. 존재하는 시간대를 /usr/share/zoneinfo/에서 찾아보시고 /etc/timezone 파일에 작성하십시오.

`root #` `ls /usr/share/zoneinfo`

`root #` `ls -l /usr/share/zoneinfo/Europe/`

```
total 256
-rw-r--r-- 1 root root 2933 Dec  3 17:19 Amsterdam
-rw-r--r-- 1 root root 1742 Dec  3 17:19 Andorra
-rw-r--r-- 1 root root 1151 Dec  3 17:19 Astrakhan
-rw-r--r-- 1 root root 2262 Dec  3 17:19 Athens
-rw-r--r-- 1 root root 3664 Dec  3 17:19 Belfast
-rw-r--r-- 1 root root 1920 Dec  3 17:19 Belgrade
-rw-r--r-- 1 root root 2298 Dec  3 17:19 Berlin
-rw-r--r-- 1 root root 2301 Dec  3 17:19 Bratislava
-rw-r--r-- 1 root root 2933 Dec  3 17:19 Brussels
...

```

Suppose the timezone of choice is Europe/Brussels, to select this timezone, a symlink can be created from this zoneinfo file to /etc/localtime:

`root #` `ln -sf ../usr/share/zoneinfo/Europe/Brussels /etc/localtime`

**요령**

The target path with `../` at the start is _relative to the link location_, not the current directory.

**참고**

An absolute path can be used for the symlink, but a relative link is also created by systemd's timedatectl and is more compatible with alternate _ROOT_ s.

## 로캘 설정

**참고**

This step does not apply to users of the musl libc. Users who do not know what that means should perform this step.

### Locale generation

A default installation of Gentoo Linux provides an archive that contains all supported locales, numbering 500 or more. However, it is typical for an administrator to require only one or two of these. In that case, the /etc/locale.gen configuration file may be populated with a list of the required locales. By default, locale-gen shall read this file and compile only the locales that are specified, saving both time and space in the longer term.

로캘은 시스템과 대화할 때 사용자가 사용할 언어에 한정하지 않으며 정렬 문자열의 규칙, 날짜 및 시간의 표시 등의 항목도 포함합니다.

`root #` `nano -w /etc/locale.gen`

다음 로캘은 (UTF-8과 같은)문자 형식에 따라 영문(미국)과 독일어(독일)를 설정하는 예제입니다.

파일 **`/etc/locale.gen`** **적절한 문자 형식으로 US 및 DE 로캘 활성화**

```
en_US ISO-8859-1
en_US.UTF-8 UTF-8
de_DE ISO-8859-1
de_DE.UTF-8 UTF-8

```

1. Non UTF-8 locales are discouraged and should only be used in special circumstances.
2. en\_US ISO-8859-1
3. de\_DE ISO-8859-1

}}

**경고**

일부 프로그램에서 UTF-8이 필요하므로, 되도록이면 UTF-8로 설정하십시오.

다음 단계는 locale-gen을 실행할 차례입니다. /etc/locale.gen파일에 지정한 모든 로캘을 만듭니다.

`root #` `locale-gen`

선택한 로캘을 사용할 수 있는지 확인하려면 locale -a을 실행하십시오.

On systemd installs, localectl can be used, e.g. localectl set-locale ... or localectl list-locales.

### Locale selection

이 과정이 끝나면 시스템 범위 로캘을 설정할 차례입니다. 이제 eselect 명령에 `locale` 모듈을 사용하겠습니다.

eselect locale list 명령으로 존재 대상을 나타냈습니다:

`root #` `eselect locale list`

```
Available targets for the LANG variable:
  [1] C
  [2] POSIX
  [3] en_US
  [4] en_US.iso88591
  [5] en_US.utf8
  [6] de_DE
  [7] de_DE.iso88591
  [8] de_DE.iso885915
  [9] de_DE.utf8
  [ ] (free form)

```

eselect locale set VALUE 명령으로 올바른 로캘을 설정할 수 있습니다:

`root #` `eselect locale set 9`

직접 설정한다면 /etc/env.d/02locale 파일에서도 처리할 수 있습니다:

파일 **`/etc/env.d/02locale`** **시스템 로캘 정의 직접 설정**

```
LANG="de_DE.UTF-8"
LC_COLLATE="C"

```

로캘을 설정했는지 확인하십시오. 그렇지 않으면, 커널을 빌드할 때와 설치 과정에서 나중에 다른 프로그램을 배포할 때 시스템에서 경고와 오류를 출력합니다.

이제 환경을 다시 불러오십시오:

`root #` `env-update && source /etc/profile && export PS1="(chroot) ${PS1}"`

이 과정을 통해 사용자 안내를 도울 완전한 [지역화 안내서](/wiki/Localization/Guide/ko "Localization/Guide/ko") 를 만들었습니다. 시스템에서 UTF-8 문자 코드를 활성화 할 주제로만 작성한 [UTF-8](/wiki/UTF-8/ko "UTF-8/ko") 안내서도 흥미로운 읽을거리입니다.

## References

1. [↑](#cite_ref-1)[https://www.gentoo.org/news/2023/12/29/Gentoo-binary.html](https://www.gentoo.org/news/2023/12/29/Gentoo-binary.html)
2. [↑](#cite_ref-2)[https://www.gentoo.org/glep/glep-0023.html#id7](https://www.gentoo.org/glep/glep-0023.html#id7)

[← 젠투 설치 파일 설치](/wiki/Handbook:AMD64/Installation/Stage/ko "이전 내용") [처음](/wiki/Handbook:AMD64/ko "Handbook:AMD64/ko") [리눅스 커널 설정 →](/wiki/Handbook:AMD64/Installation/Kernel/ko "다음 내용")