# Tools

Other languages:

- [Deutsch](/wiki/Handbook:HPPA/Installation/Tools/de "Handbuch:HPPA/Installation/Tools (100% translated)")
- [English](/wiki/Handbook:HPPA/Installation/Tools "Handbook:HPPA/Installation/Tools (100% translated)")
- [español](/wiki/Handbook:HPPA/Installation/Tools/es "Manual de Gentoo: HPPA/Instalación/Herramientas (100% translated)")
- [français](/wiki/Handbook:HPPA/Installation/Tools/fr "Handbook:HPPA/Installation/Tools/fr (100% translated)")
- [italiano](/wiki/Handbook:HPPA/Installation/Tools/it "Handbook:HPPA/Installation/Tools/it (100% translated)")
- [magyar](/wiki/Handbook:HPPA/Installation/Tools/hu "Handbook:HPPA/Installation/Tools/hu (100% translated)")
- [polski](/wiki/Handbook:HPPA/Installation/Tools/pl "Handbook:HPPA/Installation/Tools (100% translated)")
- [português do Brasil](/wiki/Handbook:HPPA/Installation/Tools/pt-br "Manual:HPPA/Instalação/Ferramentas (100% translated)")
- [русский](/wiki/Handbook:HPPA/Installation/Tools/ru "Handbook:HPPA/Installation/Tools (100% translated)")
- [தமிழ்](/wiki/Handbook:HPPA/Installation/Tools/ta "கையேடு:HPPA/நிறுவல்/கருவிகள் (100% translated)")
- [中文（中国大陆）‎](/wiki/Handbook:HPPA/Installation/Tools/zh-cn "手册：HPPA/安装/安装系统工具 (100% translated)")
- [日本語](/wiki/Handbook:HPPA/Installation/Tools/ja "ハンドブック:HPPA/インストール/ツール (100% translated)")
- 한국어

[HPPA 핸드북](/wiki/Handbook:HPPA/ko "Handbook:HPPA/ko")[설치](/wiki/Handbook:HPPA/Full/Installation/ko "Handbook:HPPA/Full/Installation/ko")[설치 정보](/wiki/Handbook:HPPA/Installation/About/ko "Handbook:HPPA/Installation/About/ko")[매체 선택](/wiki/Handbook:HPPA/Installation/Media/ko "Handbook:HPPA/Installation/Media/ko")[네트워크 설정](/wiki/Handbook:HPPA/Installation/Networking/ko "Handbook:HPPA/Installation/Networking/ko")[디스크 준비](/wiki/Handbook:HPPA/Installation/Disks/ko "Handbook:HPPA/Installation/Disks/ko")[스테이지 3 설치](/wiki/Handbook:HPPA/Installation/Stage/ko "Handbook:HPPA/Installation/Stage/ko")[베이스 시스템 설치](/wiki/Handbook:HPPA/Installation/Base/ko "Handbook:HPPA/Installation/Base/ko")[커널 설정](/wiki/Handbook:HPPA/Installation/Kernel/ko "Handbook:HPPA/Installation/Kernel/ko")[시스템 설정](/wiki/Handbook:HPPA/Installation/System/ko "Handbook:HPPA/Installation/System/ko")도구 설치[부트로더 설정](/wiki/Handbook:HPPA/Installation/Bootloader/ko "Handbook:HPPA/Installation/Bootloader/ko")[마무리](/wiki/Handbook:HPPA/Installation/Finalizing/ko "Handbook:HPPA/Installation/Finalizing/ko")[젠투 활용](/wiki/Handbook:HPPA/Full/Working/ko "Handbook:HPPA/Full/Working/ko")[포티지 소개](/wiki/Handbook:HPPA/Working/Portage/ko "Handbook:HPPA/Working/Portage/ko")[USE 플래그](/wiki/Handbook:HPPA/Working/USE/ko "Handbook:HPPA/Working/USE/ko")[포티지 기능](/wiki/Handbook:HPPA/Working/Features/ko "Handbook:HPPA/Working/Features/ko")[초기화 스크립트 시스템](/wiki/Handbook:HPPA/Working/Initscripts/ko "Handbook:HPPA/Working/Initscripts/ko")[환경 변수](/wiki/Handbook:HPPA/Working/EnvVar/ko "Handbook:HPPA/Working/EnvVar/ko")[포티지 활용](/wiki/Handbook:HPPA/Full/Portage/ko "Handbook:HPPA/Full/Portage/ko")[파일 및 디렉터리](/wiki/Handbook:HPPA/Portage/Files/ko "Handbook:HPPA/Portage/Files/ko")[변수](/wiki/Handbook:HPPA/Portage/Variables/ko "Handbook:HPPA/Portage/Variables/ko")[소프트웨어 브랜치 함께 사용하기](/wiki/Handbook:HPPA/Portage/Branches/ko "Handbook:HPPA/Portage/Branches/ko")[추가 도구](/wiki/Handbook:HPPA/Portage/Tools/ko "Handbook:HPPA/Portage/Tools/ko")[꾸러미 저장소 개별 설정](/wiki/Handbook:HPPA/Portage/CustomTree/ko "Handbook:HPPA/Portage/CustomTree/ko")[고급 기능](/wiki/Handbook:HPPA/Portage/Advanced/ko "Handbook:HPPA/Portage/Advanced/ko")[네트워크 설정](/wiki/Handbook:HPPA/Full/Networking/ko "Handbook:HPPA/Full/Networking/ko")[시작하기](/wiki/Handbook:HPPA/Networking/Introduction/ko "Handbook:HPPA/Networking/Introduction/ko")[고급 설정](/wiki/Handbook:HPPA/Networking/Advanced/ko "Handbook:HPPA/Networking/Advanced/ko")[모듈러 네트워크](/wiki/Handbook:HPPA/Networking/Modular/ko "Handbook:HPPA/Networking/Modular/ko")[무선 네트워크](/wiki/Handbook:HPPA/Networking/Wireless/ko "Handbook:HPPA/Networking/Wireless/ko")[기능 추가](/wiki/Handbook:HPPA/Networking/Extending/ko "Handbook:HPPA/Networking/Extending/ko")[동적 관리](/wiki/Handbook:HPPA/Networking/Dynamic/ko "Handbook:HPPA/Networking/Dynamic/ko")

## Contents

- [1시스템 로거](#.EC.8B.9C.EC.8A.A4.ED.85.9C_.EB.A1.9C.EA.B1.B0)
  - [1.1OpenRC](#OpenRC)
  - [1.2systemd](#systemd)
- [2선택: 크론 데몬](#.EC.84.A0.ED.83.9D:_.ED.81.AC.EB.A1.A0_.EB.8D.B0.EB.AA.AC)
  - [2.1OpenRC](#OpenRC_2)
    - [2.1.1Default: cronie](#Default:_cronie)
    - [2.1.2Alternative: dcron](#Alternative:_dcron)
    - [2.1.3Alternative: fcron](#Alternative:_fcron)
    - [2.1.4Alternative: bcron](#Alternative:_bcron)
    - [2.1.5modprobed-db users](#modprobed-db_users)
  - [2.2systemd](#systemd_2)
    - [2.2.1modprobed-db users](#modprobed-db_users_2)
- [3선택: 파일 색인](#.EC.84.A0.ED.83.9D:_.ED.8C.8C.EC.9D.BC_.EC.83.89.EC.9D.B8)
- [4선택: 원격 접근](#.EC.84.A0.ED.83.9D:_.EC.9B.90.EA.B2.A9_.EC.A0.91.EA.B7.BC)
  - [4.1OpenRC](#OpenRC_3)
  - [4.2systemd](#systemd_3)
- [5Optional: Shell completion](#Optional:_Shell_completion)
  - [5.1Bash](#Bash)
- [6Suggested: Time synchronization](#Suggested:_Time_synchronization)
  - [6.1chrony](#chrony)
    - [6.1.1OpenRC](#OpenRC_4)
    - [6.1.2systemd](#systemd_4)
  - [6.2systemd-timesyncd](#systemd-timesyncd)
- [7파일 시스템 도구](#.ED.8C.8C.EC.9D.BC_.EC.8B.9C.EC.8A.A4.ED.85.9C_.EB.8F.84.EA.B5.AC)

## 시스템 로거

### OpenRC

스테이지 3 아카이브에서 몇가지 도구가 빠졌는데 대부분 꾸러미가 동일한 기능을 지니고 있기 때문입니다. 이제 설치할 도구를 선택하는건 사용자의 몫입니다.

첫번째로 결정해야 할 도구는 시스템 로깅 수단을 제공합니다. 유닉스 및 리눅스는 로깅 능력에 있어 멋진 역사를 지니고 있습니다 \- 필요하다면, 로그 파일에 시스템에 일어나는 모든 일을 기록할 수 있습니다. 이 일은 시스템 로거가 처리합니다.

젠투는 선택할 다양한 시스템 로거를 제공합니다. 그 중 몇가지가 있다면:

- [app-admin/sysklogd](https://packages.gentoo.org/packages/app-admin/sysklogd) \- 시스템 로깅 데몬의 기존 모음입니다. 초보자를 배려하여 기본 로깅 설정으로도 그 자체로 특별하게 잘 동작합니다.
- [app-admin/syslog-ng](https://packages.gentoo.org/packages/app-admin/syslog-ng) \- 최근의 시스템 로거입니다. 하나의 큰 파일이 아닌 다른 방식으로 로깅하려면 추가 설정이 필요합니다. 좀 더 능력있는 사용자라면 로깅 잠재 기능때문에 이 꾸러미를 사용합니다. 지능 로깅 동작은 추가 설정이 필요합니다.
- [app-admin/metalog](https://packages.gentoo.org/packages/app-admin/metalog) \- 매우 설정하기 쉬운 시스템 로거

마찬가지로 포티지에 다른 다양한 로거가 존재합니다. 수많은 꾸러미는 매일 늘어납니다.

**요령**

sysklogd 또는 syslog-ng를 사용하려 한다면, 시스템 로거가 로그 파일에 대한 순환 매커니즘을 제공하지 않으므로, 이들 꾸러미를 설치한 후 [logrotate](/wiki/Logrotate "Logrotate") 를 설치 및 설정하는 것이 좋습니다.

선택한 시스템 로거를 설치하려면, emerge후, rc-update를 사용하여 기본 런레벨에 추가해야합니다. 다음 예제에서는 [app-admin/sysklogd](https://packages.gentoo.org/packages/app-admin/sysklogd)를 설치합니다:

`root #` `emerge --ask app-admin/sysklogd`

`root #` `rc-update add sysklogd default`

### systemd

While a selection of logging mechanisms are presented for OpenRC-based systems, systemd includes a built-in logger called the **systemd-journald** service. The systemd-journald service is capable of handling most of the logging functionality outlined in the previous system logger section. That is to say, the majority of installations that will run systemd as the system and service manager can _safely skip_ adding a additional syslog utilities.

See man journalctl for more details on using journalctl to query and review the systems logs.

For a number of reasons, such as the case of forwarding logs to a central host, it may be important to include _redundant_ system logging mechanisms on a systemd-based system. This is a irregular occurrence for the handbook's typical audience and considered an advanced use case. It is therefore not covered by the handbook.

## 선택: 크론 데몬

### OpenRC

다음은 크론 데몬입니다. 설치를 해도 안해도 그만이며, 모든 시스템에서 설치할 필요는 없지만, 설치하는게 현명합니다.

크론 데몬은 일정별로 계획한 명령을 실행합니다. 규칙적으로(예를 들어 매일, 주별, 월별) 실행할 필요가 있는 명령에 대해 매우 간편합니다.

All cron daemons support high levels of granularity for scheduled tasks, and generally include the ability to send an email or other form of notification if a scheduled task does not complete as expected.

젠투는 [sys-process/bcron](https://packages.gentoo.org/packages/sys-process/bcron), [sys-process/dcron](https://packages.gentoo.org/packages/sys-process/dcron), [sys-process/fcron](https://packages.gentoo.org/packages/sys-process/fcron), [sys-process/cronie](https://packages.gentoo.org/packages/sys-process/cronie) 등의 다양한 크론 데몬을 제공합니다. 이들 중 하나를 설치하는 건 시스템 로거를 설치할 때와 마찬가지입니다. 다음 예제에서는 [sys-process/cronie](https://packages.gentoo.org/packages/sys-process/cronie)를 설치합니다:

- [cronie](/wiki/Cron/ko#cronie "Cron/ko") ([sys-process/cronie](https://packages.gentoo.org/packages/sys-process/cronie)) \- cronie is based on the original cron and has security and configuration enhancements like the ability to use PAM and SELinux.
- [dcron](/wiki/Cron/ko#dcron_.28Dillon.27s_Cron.29 "Cron/ko") ([sys-process/dcron](https://packages.gentoo.org/packages/sys-process/dcron)) \- This lightweight cron daemon aims to be simple and secure, with just enough features to stay useful.
- [fcron](/wiki/Cron/ko#fcron "Cron/ko") ([sys-process/fcron](https://packages.gentoo.org/packages/sys-process/fcron)) \- A command scheduler with extended capabilities over cron and anacron.
- [bcron](/wiki/Cron/ko#bcron "Cron/ko") ([sys-process/bcron](https://packages.gentoo.org/packages/sys-process/bcron)) \- A younger cron system designed with secure operations in mind. To do this, the system is divided into several separate programs, each responsible for a separate task, with strictly controlled communications between parts.

#### Default: cronie

The following example uses [sys-process/cronie](https://packages.gentoo.org/packages/sys-process/cronie) (not to be confused with the [NTP](/wiki/Network_Time_Protocol "Network Time Protocol") daemon named [chrony](/wiki/Chrony "Chrony")):

`root #` `emerge --ask sys-process/cronie`

`root #` `rc-update add cronie default`

`root #` `rc-update add cronie default`

#### Alternative: dcron

`root #` `emerge --ask sys-process/dcron`

dcron또는 fcron을 사용한다면, 추가 초기화 명령을 실행해야합니다:

`root #` `crontab /etc/crontab`

#### Alternative: fcron

`root #` `emerge --ask sys-process/fcron`

If fcron is the selected scheduled task handler, an additional emerge step is required:

`root #` `emerge --config sys-process/fcron`

#### Alternative: bcron

bcron is a younger cron agent with built-in privilege separation.

`root #` `emerge --ask sys-process/bcron`

#### modprobed-db users

If using [sys-kernel/modprobed-db](https://packages.gentoo.org/packages/sys-kernel/modprobed-db) was picked as an option for manually compiling the kernel.

Now is the time to setup a crontab to periodically scan the system for hardware used.

파일 **`/etc/crontab`** **Run modprobed-db once every 6 hours**

```
0 */6 * * *     /usr/bin/modprobed-db store &> /dev/null

```

At a later date of at least a week, please visit the kernel build section of the [modprobed-db](/wiki/Modprobed-db#Building_kernels.2Fko "Modprobed-db") article to complete the setup.

### systemd

Similar to system logging, systemd-based systems include support for scheduled tasks out-of-the-box in the form of _timers_. systemd timers can run at a system-level or a user-level and include the same functionality that a traditional cron daemon would provide. Unless redundant capabilities are necessary, installing an additional task scheduler such as a cron daemon is generally unnecessary and can be safely skipped.

#### modprobed-db users

If using [sys-kernel/modprobed-db](https://packages.gentoo.org/packages/sys-kernel/modprobed-db) was picked as an option for manually compiling the kernel.

Now is the time to setup a systemd timer to periodically scan the system for hardware used.

`root #` `systemctl --user enable modprobed-db`

At a later date of at least a week, please visit the kernel build section of the [modprobed-db](/wiki/Modprobed-db#Building_kernels.2Fko "Modprobed-db") article to complete the setup.

## 선택: 파일 색인

파일 시스템을 색인 처리하여 파일 탐색을 더 빠르게 하려면 [sys-apps/mlocate](https://packages.gentoo.org/packages/sys-apps/mlocate)를 설치하십시오.

`root #` `emerge --ask sys-apps/mlocate`

## 선택: 원격 접근

**요령**

opensshd's default configuration does not allow root to login as a remote user. Please [create a non-root user](/wiki/FAQ/ko#How_do_I_add_a_normal_user.3F "FAQ/ko") and configure it appropriately to allow access post-installation if required, or adjust /etc/ssh/sshd\_config to allow root.

설치 후 시스템을 원격으로 접근하려면, 기본 런레벨에 sshd 초기화 스크립트를 추가하십시오:

For more in-depth details on the configuration of SSH, refer to the [SSH](/wiki/SSH/ko "SSH/ko") article.

### OpenRC

`root #` `rc-update add sshd default`

`root #` `rc-update add sshd default`

직렬 콘솔 접근이 필요하다면 (원격 서버의 경우 가능) /etc/inittab에서 직렬 콘솔 섹션의 주석 표시를 빼십시오:

Uncomment the serial console section in /etc/inittab:

`root #` `nano -w /etc/inittab`

```
# SERIAL CONSOLES
s0:12345:respawn:/sbin/agetty 9600 ttyS0 vt100
s1:12345:respawn:/sbin/agetty 9600 ttyS1 vt100

```

### systemd

To enable the SSH server, run:

`root #` `systemctl enable sshd`

To enable serial console support, run:

`root #` `systemctl enable getty@tty1.service`

## Optional: Shell completion

### Bash

Bash is the default shell for Gentoo systems, and therefore installing completion extensions can aid in efficiency and convenience to managing the system. The [app-shells/bash-completion](https://packages.gentoo.org/packages/app-shells/bash-completion) package will install completions available for Gentoo specific commands, as well as many other common commands and utilities:

`root #` `emerge --ask app-shells/bash-completion`

Post installation, bash completion for specific commands can managed through eselect. See the [Shell completion integrations section](/wiki/Bash#Shell_completion_integrations.2Fko "Bash") of the bash article for more details.

## Suggested: Time synchronization

**중요**

Systems without a functioning [Real-Time Clock (RTC)](/wiki/System_time/ko#Software_clock_vs_Hardware_clock "System time/ko") must set the [system time](/wiki/System_time/ko "System time/ko") at every system start, and on regular intervals thereafter.

It is important to use some method of synchronizing the system clock with Internet time servers. This is often mandatory for systems without an RTC, but is also beneficial for systems with a RTC, as the battery could fail, and clock skew can accumulate.

The clock is usually synchronized via the [Network Time Protocol](/wiki/Network_Time_Protocol "Network Time Protocol"), with an implementation such as [chrony](/wiki/Chrony "Chrony").

### chrony

Install chrony (not to be confused with the cron daemon named [cronie](/wiki/Cron/ko#cronie "Cron/ko")):

`root #` `emerge --ask net-misc/chrony`

See the [chrony](/wiki/Chrony "Chrony") article for further information, for example if more advanced configurations are required.

#### OpenRC

Add chronyd to the default runlevel to start it at boot, so time will be synchronized automatically:

`root #` `rc-update add chronyd default`

#### systemd

Start chronyd at boot, to have the time synchronized automatically:

`root #` `systemctl enable chronyd.service`

### systemd-timesyncd

Alternatively, systemd users may wish to use the simpler systemd-timesyncd SNTP client which is installed by default and enabled with:

`root #` `systemctl enable systemd-timesyncd.service`

## 파일 시스템 도구

사용하는 파일 시스템에 따라 필요한 파일 시스템 유틸리티를 설치해야 합니다(파일 시스템 무결성 검사, 추가 파일 시스템 만들기 등). ext2, ext3, ext4 파일 시스템([sys-fs/e2fsprogs](https://packages.gentoo.org/packages/sys-fs/e2fsprogs))을 관리하는 도구는 이미 [@system 세트](/wiki/System_set_(Portage) "System set (Portage)") 의 일부로 설치했음을 참고하십시오.

다음 테이블 목록에서는 각각의 파일 시스템을 사용할 경우 설치할 도구를 보여줍니다:

파일 시스템
꾸러미
Ext2, 3, and 4
[sys-fs/e2fsprogs](https://packages.gentoo.org/packages/sys-fs/e2fsprogs)XFS
[sys-fs/xfsprogs](https://packages.gentoo.org/packages/sys-fs/xfsprogs)ReiserFS
[sys-fs/reiserfsprogs](https://packages.gentoo.org/packages/sys-fs/reiserfsprogs)JFS
[sys-fs/jfsutils](https://packages.gentoo.org/packages/sys-fs/jfsutils)VFAT (FAT32, ...)
[sys-fs/dosfstools](https://packages.gentoo.org/packages/sys-fs/dosfstools)Btrfs
[sys-fs/btrfs-progs](https://packages.gentoo.org/packages/sys-fs/btrfs-progs)

It's recommended that [sys-block/io-scheduler-udev-rules](https://packages.gentoo.org/packages/sys-block/io-scheduler-udev-rules) be installed for the correct scheduler behavior with e.g. nvme devices:

`root #` `emerge --ask sys-block/io-scheduler-udev-rules`

**요령**

자세한 젠투 파일 시스템 정보를 보려면 [파일 시스템 게시글](/wiki/Filesystem/ko "Filesystem/ko") 을 살펴보십시오.

이제 [부트로더 설정](/wiki/Handbook:HPPA/Installation/Bootloader/ko "Handbook:HPPA/Installation/Bootloader/ko") 으로 계속 진행하십시오.

[← 시스템 설정](/wiki/Handbook:HPPA/Installation/System/ko "이전 내용") [처음](/wiki/Handbook:HPPA/ko "Handbook:HPPA/ko") [부트로더 설정 →](/wiki/Handbook:HPPA/Installation/Bootloader/ko "다음 내용")