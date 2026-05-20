# Tools

Other languages:

- [Deutsch](/wiki/Handbook:Alpha/Installation/Tools/de "Handbuch:Alpha/Installation/Tools (100% translated)")
- [English](/wiki/Handbook:Alpha/Installation/Tools "Handbook:Alpha/Installation/Tools (100% translated)")
- [español](/wiki/Handbook:Alpha/Installation/Tools/es "Manual de Gentoo: Alpha/Instalación/Herramientas (100% translated)")
- [français](/wiki/Handbook:Alpha/Installation/Tools/fr "Handbook:Alpha/Installation/Tools/fr (100% translated)")
- [italiano](/wiki/Handbook:Alpha/Installation/Tools/it "Handbook:Alpha/Installation/Tools/it (100% translated)")
- [magyar](/wiki/Handbook:Alpha/Installation/Tools/hu "Handbook:Alpha/Installation/Tools/hu (100% translated)")
- polski
- [português do Brasil](/wiki/Handbook:Alpha/Installation/Tools/pt-br "Manual:Alpha/Instalação/Ferramentas (100% translated)")
- [čeština](/wiki/Handbook:Alpha/Installation/Tools/cs "Handbook:Alpha/Installation/Tools/cs (50% translated)")
- [русский](/wiki/Handbook:Alpha/Installation/Tools/ru "Handbook:Alpha/Installation/Tools (100% translated)")
- [தமிழ்](/wiki/Handbook:Alpha/Installation/Tools/ta "கையேடு:Alpha/நிறுவல்/கருவிகள் (100% translated)")
- [中文（中国大陆）‎](/wiki/Handbook:Alpha/Installation/Tools/zh-cn "手册：Alpha/安装/安装系统工具 (100% translated)")
- [日本語](/wiki/Handbook:Alpha/Installation/Tools/ja "ハンドブック:Alpha/インストール/ツール (100% translated)")
- [한국어](/wiki/Handbook:Alpha/Installation/Tools/ko "Handbook:Alpha/Installation/Tools/ko (100% translated)")

[Alpha Handbook](/wiki/Handbook:Alpha/pl "Handbook:Alpha/pl")[Instalacja](/wiki/Handbook:Alpha/Full/Installation/pl "Handbook:Alpha/Full/Installation/pl")[O instalacji](/wiki/Handbook:Alpha/Installation/About/pl "Handbook:Alpha/Installation/About/pl")[Wybór medium instalacyjnego](/wiki/Handbook:Alpha/Installation/Media/pl "Handbook:Alpha/Installation/Media/pl")[Konfiguracja sieci](/wiki/Handbook:Alpha/Installation/Networking/pl "Handbook:Alpha/Installation/Networking/pl")[Przygotowanie dysków](/wiki/Handbook:Alpha/Installation/Disks/pl "Handbook:Alpha/Installation/Disks/pl")[Instalacja etapu 3](/wiki/Handbook:Alpha/Installation/Stage/pl "Handbook:Alpha/Installation/Stage/pl")[Instalacja systemu podstawowego](/wiki/Handbook:Alpha/Installation/Base/pl "Handbook:Alpha/Installation/Base/pl")[Konfiguracja jądra](/wiki/Handbook:Alpha/Installation/Kernel/pl "Handbook:Alpha/Installation/Kernel/pl")[Konfiguracja systemu](/wiki/Handbook:Alpha/Installation/System/pl "Handbook:Alpha/Installation/System/pl")Instalacja narzędzi[Instalacja systemu rozruchowego](/wiki/Handbook:Alpha/Installation/Bootloader/pl "Handbook:Alpha/Installation/Bootloader/pl")[Finalizacja](/wiki/Handbook:Alpha/Installation/Finalizing/pl "Handbook:Alpha/Installation/Finalizing/pl")[Praca z Gentoo](/wiki/Handbook:Alpha/Full/Working/pl "Handbook:Alpha/Full/Working/pl")[Wstęp do Portage](/wiki/Handbook:Alpha/Working/Portage/pl "Handbook:Alpha/Working/Portage/pl")[Flagi USE](/wiki/Handbook:Alpha/Working/USE/pl "Handbook:Alpha/Working/USE/pl")[Funkcje portage](/wiki/Handbook:Alpha/Working/Features/pl "Handbook:Alpha/Working/Features/pl")[System initscript](/wiki/Handbook:Alpha/Working/Initscripts/pl "Handbook:Alpha/Working/Initscripts/pl")[Zmienne środowiskowe](/wiki/Handbook:Alpha/Working/EnvVar/pl "Handbook:Alpha/Working/EnvVar/pl")[Praca z Portage](/wiki/Handbook:Alpha/Full/Portage/pl "Handbook:Alpha/Full/Portage/pl")[Pliki i katalogi](/wiki/Handbook:Alpha/Portage/Files/pl "Handbook:Alpha/Portage/Files/pl")[Zmienne](/wiki/Handbook:Alpha/Portage/Variables/pl "Handbook:Alpha/Portage/Variables/pl")[Mieszanie działów oprogramowania](/wiki/Handbook:Alpha/Portage/Branches/pl "Handbook:Alpha/Portage/Branches/pl")[Dodatkowe narzędzia](/wiki/Handbook:Alpha/Portage/Tools/pl "Handbook:Alpha/Portage/Tools/pl")[Custom package repository](/wiki/Handbook:Alpha/Portage/CustomTree/pl "Handbook:Alpha/Portage/CustomTree/pl")[Funkcje zaawansowane](/wiki/Handbook:Alpha/Portage/Advanced/pl "Handbook:Alpha/Portage/Advanced/pl")[Konfiguracja sieci](/wiki/Handbook:Alpha/Full/Networking/pl "Handbook:Alpha/Full/Networking/pl")[Zaczynamy](/wiki/Handbook:Alpha/Networking/Introduction/pl "Handbook:Alpha/Networking/Introduction/pl")[Zaawansowana konfiguracja](/wiki/Handbook:Alpha/Networking/Advanced/pl "Handbook:Alpha/Networking/Advanced/pl")[Sieć modularna](/wiki/Handbook:Alpha/Networking/Modular/pl "Handbook:Alpha/Networking/Modular/pl")[Sieć bezprzewodowa](/wiki/Handbook:Alpha/Networking/Wireless/pl "Handbook:Alpha/Networking/Wireless/pl")[Dodawanie funkcjonalności](/wiki/Handbook:Alpha/Networking/Extending/pl "Handbook:Alpha/Networking/Extending/pl")[Dynamiczne zarządzanie](/wiki/Handbook:Alpha/Networking/Dynamic/pl "Handbook:Alpha/Networking/Dynamic/pl")

## Contents

- [1System logger](#System_logger)
  - [1.1OpenRC](#OpenRC)
  - [1.2systemd](#systemd)
- [2Cron daemon](#Cron_daemon)
  - [2.1OpenRC](#OpenRC_2)
    - [2.1.1Default: cronie](#Default:_cronie)
    - [2.1.2Alternative: dcron](#Alternative:_dcron)
    - [2.1.3Alternative: fcron](#Alternative:_fcron)
    - [2.1.4Alternative: bcron](#Alternative:_bcron)
    - [2.1.5modprobed-db users](#modprobed-db_users)
  - [2.2systemd](#systemd_2)
    - [2.2.1modprobed-db users](#modprobed-db_users_2)
- [3Optional: File indexing](#Optional:_File_indexing)
- [4Optional: Remote shell access](#Optional:_Remote_shell_access)
  - [4.1OpenRC](#OpenRC_3)
  - [4.2systemd](#systemd_3)
- [5Optional: Shell completion](#Optional:_Shell_completion)
  - [5.1Bash](#Bash)
- [6Suggested: Time synchronization](#Suggested:_Time_synchronization)
  - [6.1chrony](#chrony)
    - [6.1.1OpenRC](#OpenRC_4)
    - [6.1.2systemd](#systemd_4)
  - [6.2systemd-timesyncd](#systemd-timesyncd)
- [7Filesystem tools](#Filesystem_tools)

## System logger

### OpenRC

Some tools are missing from the stage3 archive because several packages provide the same functionality. It is now up to the user to choose which ones to install.

The first tool decision is a logging mechanism for the system. Unix and Linux have an excellent history of logging capabilities - if needed, everything that happens on the system can be logged in a log file.

Gentoo offers several system logger utilities. A few of these include:

- [sysklogd](/wiki/Sysklogd "Sysklogd") ([app-admin/sysklogd](https://packages.gentoo.org/packages/app-admin/sysklogd)) \- Offers the traditional set of system logging daemons. The default logging configuration works well out of the box which makes this package a good option for beginners.
- [syslog-ng](/wiki/Syslog-ng "Syslog-ng") ([app-admin/syslog-ng](https://packages.gentoo.org/packages/app-admin/syslog-ng)) \- An advanced system logger. Requires additional configuration for anything beyond logging to one big file. More advanced users may choose this package based on its logging potential; be aware additional configuration is a necessity for any kind of smart logging.
- [metalog](/wiki/Metalog "Metalog") ([app-admin/metalog](https://packages.gentoo.org/packages/app-admin/metalog)) \- A highly-configurable system logger.

There may be other system logging utilities available through the Gentoo ebuild repository as well, since the number of available packages increases on a daily basis.

**Wskazówka**

If syslog-ng is going to be used, it is recommended to install and configure [logrotate](/wiki/Logrotate "Logrotate"). syslog-ng does not provide any rotation mechanism for the log files. Newer versions (>= 2.0) of sysklogd however handle their own log rotation.

To install the system logger of choice, emerge it. On OpenRC, add it to the default runlevel using rc-update. The following example installs and activates [app-admin/sysklogd](https://packages.gentoo.org/packages/app-admin/sysklogd) as the system's syslog utility:

`root #` `emerge --ask app-admin/sysklogd`

`root #` `rc-update add sysklogd default`

### systemd

While a selection of logging mechanisms are presented for OpenRC-based systems, systemd includes a built-in logger called the **systemd-journald** service. The systemd-journald service is capable of handling most of the logging functionality outlined in the previous system logger section. That is to say, the majority of installations that will run systemd as the system and service manager can _safely skip_ adding a additional syslog utilities.

See man journalctl for more details on using journalctl to query and review the systems logs.

For a number of reasons, such as the case of forwarding logs to a central host, it may be important to include _redundant_ system logging mechanisms on a systemd-based system. This is a irregular occurrence for the handbook's typical audience and considered an advanced use case. It is therefore not covered by the handbook.

## Cron daemon

### OpenRC

Although it is optional and not required for every system, it is wise to install a cron daemon.

A cron daemon executes commands on scheduled intervals. Intervals could be daily, weekly, or monthly, once every Tuesday, once every other week, etc. A wise system administrator will leverage the cron daemon to automate routine system maintenance tasks.

All cron daemons support high levels of granularity for scheduled tasks, and generally include the ability to send an email or other form of notification if a scheduled task does not complete as expected.

Gentoo offers several possible cron daemons, including:

- [cronie](/wiki/Cron#cronie "Cron") ([sys-process/cronie](https://packages.gentoo.org/packages/sys-process/cronie)) \- cronie is based on the original cron and has security and configuration enhancements like the ability to use PAM and SELinux.
- [dcron](/wiki/Cron#dcron_.28Dillon.27s_Cron.29 "Cron") ([sys-process/dcron](https://packages.gentoo.org/packages/sys-process/dcron)) \- This lightweight cron daemon aims to be simple and secure, with just enough features to stay useful.
- [fcron](/wiki/Cron#fcron "Cron") ([sys-process/fcron](https://packages.gentoo.org/packages/sys-process/fcron)) \- A command scheduler with extended capabilities over cron and anacron.
- [bcron](/wiki/Cron#bcron "Cron") ([sys-process/bcron](https://packages.gentoo.org/packages/sys-process/bcron)) \- A younger cron system designed with secure operations in mind. To do this, the system is divided into several separate programs, each responsible for a separate task, with strictly controlled communications between parts.

#### Default: cronie

The following example uses [sys-process/cronie](https://packages.gentoo.org/packages/sys-process/cronie) (not to be confused with the [NTP](/wiki/Network_Time_Protocol "Network Time Protocol") daemon named [chrony](/wiki/Chrony "Chrony")):

`root #` `emerge --ask sys-process/cronie`

Add cronie to the default system runlevel, which will automatically start it on power up:

`root #` `rc-update add cronie default`

#### Alternative: dcron

`root #` `emerge --ask sys-process/dcron`

If dcron is the go forward cron agent, an additional initialization command needs to be executed:

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

FILE **`/etc/crontab`** **Run modprobed-db once every 6 hours**

```
0 */6 * * *     /usr/bin/modprobed-db store &> /dev/null

```

At a later date of at least a week, please visit the kernel build section of the [modprobed-db](/wiki/Modprobed-db#Building_kernels.2Fpl "Modprobed-db") article to complete the setup.

### systemd

Similar to system logging, systemd-based systems include support for scheduled tasks out-of-the-box in the form of _timers_. systemd timers can run at a system-level or a user-level and include the same functionality that a traditional cron daemon would provide. Unless redundant capabilities are necessary, installing an additional task scheduler such as a cron daemon is generally unnecessary and can be safely skipped.

#### modprobed-db users

If using [sys-kernel/modprobed-db](https://packages.gentoo.org/packages/sys-kernel/modprobed-db) was picked as an option for manually compiling the kernel.

Now is the time to setup a systemd timer to periodically scan the system for hardware used.

`root #` `systemctl --user enable modprobed-db`

At a later date of at least a week, please visit the kernel build section of the [modprobed-db](/wiki/Modprobed-db#Building_kernels.2Fpl "Modprobed-db") article to complete the setup.

## Optional: File indexing

In order to index the file system to provide faster file location capabilities, install [mlocate](/wiki/Mlocate "Mlocate"):

`root #` `emerge --ask sys-apps/mlocate`

## Optional: Remote shell access

**Wskazówka**

opensshd's default configuration does not allow root to login as a remote user. Please [create a non-root user](/wiki/FAQ/pl#How_do_I_add_a_normal_user.3F "FAQ/pl") and configure it appropriately to allow access post-installation if required, or adjust /etc/ssh/sshd\_config to allow root.

To be able to access the system remotely after installation, sshd must be configured to start on boot.

For more in-depth details on the configuration of SSH, refer to the [SSH](/wiki/SSH "SSH") article.

### OpenRC

To add the sshd init script to the default runlevel on OpenRC:

`root #` `rc-update add sshd default`

If serial console access is needed (which is possible in case of remote servers), agetty must be configured.

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

Post installation, bash completion for specific commands can managed through eselect. See the [Shell completion integrations section](/wiki/Bash#Shell_completion_integrations.2Fpl "Bash") of the bash article for more details.

## Suggested: Time synchronization

**Ważne**

Systems without a functioning [Real-Time Clock (RTC)](/wiki/System_time#Software_clock_vs_Hardware_clock "System time") must set the [system time](/wiki/System_time "System time") at every system start, and on regular intervals thereafter.

It is important to use some method of synchronizing the system clock with Internet time servers. This is often mandatory for systems without an RTC, but is also beneficial for systems with a RTC, as the battery could fail, and clock skew can accumulate.

The clock is usually synchronized via the [Network Time Protocol](/wiki/Network_Time_Protocol "Network Time Protocol"), with an implementation such as [chrony](/wiki/Chrony "Chrony").

### chrony

Install chrony (not to be confused with the cron daemon named [cronie](/wiki/Cron#cronie "Cron")):

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

## Filesystem tools

Depending on the filesystems used, it may be necessary to install the required file system utilities (for checking the filesystem integrity, (re)formatting file systems, etc.). Note that ext4 user space tools ([sys-fs/e2fsprogs](https://packages.gentoo.org/packages/sys-fs/e2fsprogs)) are already installed as a part of the [@system set](/wiki/System_set_(Portage) "System set (Portage)").

The following table lists the tools to install if a certain filesystem tools will be needed in the installed environment.

Filesystem
Package
[XFS](/wiki/XFS "XFS")[sys-fs/xfsprogs](https://packages.gentoo.org/packages/sys-fs/xfsprogs)[ext4](/wiki/Ext4 "Ext4")[sys-fs/e2fsprogs](https://packages.gentoo.org/packages/sys-fs/e2fsprogs)[VFAT](/wiki/VFAT "VFAT") (FAT32, ...)
[sys-fs/dosfstools](https://packages.gentoo.org/packages/sys-fs/dosfstools)[Btrfs](/wiki/Btrfs/pl "Btrfs/pl")[sys-fs/btrfs-progs](https://packages.gentoo.org/packages/sys-fs/btrfs-progs)[F2FS](/wiki/F2FS "F2FS")[sys-fs/f2fs-tools](https://packages.gentoo.org/packages/sys-fs/f2fs-tools)[NTFS](/wiki/NTFS "NTFS")[sys-fs/ntfs3g](https://packages.gentoo.org/packages/sys-fs/ntfs3g)[ZFS](/wiki/ZFS "ZFS")[sys-fs/zfs](https://packages.gentoo.org/packages/sys-fs/zfs)

It's recommended that [sys-block/io-scheduler-udev-rules](https://packages.gentoo.org/packages/sys-block/io-scheduler-udev-rules) be installed for the correct scheduler behavior with e.g. nvme devices:

`root #` `emerge --ask sys-block/io-scheduler-udev-rules`

**Wskazówka**

For more information on filesystems in Gentoo see the [filesystem article](/wiki/Filesystem "Filesystem").

Now continue with [Configuring the bootloader](/wiki/Handbook:Alpha/Installation/Bootloader/pl "Handbook:Alpha/Installation/Bootloader/pl").

[← Configuring the system](/wiki/Handbook:Alpha/Installation/System "Previous part") [Home](/wiki/Handbook:Alpha "Handbook:Alpha") [Configuring the bootloader →](/wiki/Handbook:Alpha/Installation/Bootloader "Next part")