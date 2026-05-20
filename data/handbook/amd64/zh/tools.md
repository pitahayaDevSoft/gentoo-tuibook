# Tools

Other languages:

- [Deutsch](/wiki/Handbook:AMD64/Installation/Tools/de "Handbook:AMD64/Installation/Tools/de (100% translated)")
- [English](/wiki/Handbook:AMD64/Installation/Tools "Handbook:AMD64/Installation/Tools (100% translated)")
- [español](/wiki/Handbook:AMD64/Installation/Tools/es "Handbook:AMD64/Installation/Tools/es (100% translated)")
- [français](/wiki/Handbook:AMD64/Installation/Tools/fr "Handbook:AMD64/Installation/Tools/fr (100% translated)")
- [italiano](/wiki/Handbook:AMD64/Installation/Tools/it "Handbook:AMD64/Installation/Tools/it (100% translated)")
- [magyar](/wiki/Handbook:AMD64/Installation/Tools/hu "Handbook:AMD64/Installation/Tools/hu (100% translated)")
- [polski](/wiki/Handbook:AMD64/Installation/Tools/pl "Handbook:AMD64/Installation/Tools/pl (100% translated)")
- [português do Brasil](/wiki/Handbook:AMD64/Installation/Tools/pt-br "Handbook:AMD64/Installation/Tools/pt-br (100% translated)")
- [čeština](/wiki/Handbook:AMD64/Installation/Tools/cs "Handbook:AMD64/Installation/Tools/cs (100% translated)")
- [русский](/wiki/Handbook:AMD64/Installation/Tools/ru "Handbook:AMD64/Installation/Tools/ru (100% translated)")
- [தமிழ்](/wiki/Handbook:AMD64/Installation/Tools/ta "Handbook:AMD64/Installation/Tools/ta (100% translated)")
- 中文
- [中文（中国大陆）‎](/wiki/Handbook:AMD64/Installation/Tools/zh-cn "Handbook:AMD64/Installation/Tools/zh-cn (100% translated)")
- [日本語](/wiki/Handbook:AMD64/Installation/Tools/ja "Handbook:AMD64/Installation/Tools/ja (100% translated)")
- [한국어](/wiki/Handbook:AMD64/Installation/Tools/ko "Handbook:AMD64/Installation/Tools/ko (100% translated)")

[AMD64 手册](/wiki/Handbook:AMD64 "Handbook:AMD64")[安装](/wiki/Handbook:AMD64/Full/Installation "Handbook:AMD64/Full/Installation")[关于安装](/wiki/Handbook:AMD64/Installation/About/zh "Handbook:AMD64/Installation/About/zh")[选择介质](/wiki/Handbook:AMD64/Installation/Media "Handbook:AMD64/Installation/Media")[配置网络](/wiki/Handbook:AMD64/Installation/Networking "Handbook:AMD64/Installation/Networking")[准备磁盘](/wiki/Handbook:AMD64/Installation/Disks "Handbook:AMD64/Installation/Disks")[stage 文件](/wiki/Handbook:AMD64/Installation/Stage "Handbook:AMD64/Installation/Stage")[安装基础系统](/wiki/Handbook:AMD64/Installation/Base/zh "Handbook:AMD64/Installation/Base/zh")[配置内核](/wiki/Handbook:AMD64/Installation/Kernel "Handbook:AMD64/Installation/Kernel")[配置系统](/wiki/Handbook:AMD64/Installation/System "Handbook:AMD64/Installation/System")安装工具[配置引导加载程序](/wiki/Handbook:AMD64/Installation/Bootloader "Handbook:AMD64/Installation/Bootloader")[完成安装](/wiki/Handbook:AMD64/Installation/Finalizing "Handbook:AMD64/Installation/Finalizing")[使用 Gentoo](/wiki/Handbook:AMD64/Full/Working "Handbook:AMD64/Full/Working")[Portage 简介](/wiki/Handbook:AMD64/Working/Portage "Handbook:AMD64/Working/Portage")[USE 标志](/wiki/Handbook:AMD64/Working/USE "Handbook:AMD64/Working/USE")[Portage 功能](/wiki/Handbook:AMD64/Working/Features "Handbook:AMD64/Working/Features")[初始化脚本系统](/wiki/Handbook:AMD64/Working/Initscripts "Handbook:AMD64/Working/Initscripts")[环境变量](/wiki/Handbook:AMD64/Working/EnvVar "Handbook:AMD64/Working/EnvVar")[使用 Portage](/wiki/Handbook:AMD64/Full/Portage "Handbook:AMD64/Full/Portage")[文件和目录](/wiki/Handbook:AMD64/Portage/Files "Handbook:AMD64/Portage/Files")[变量](/wiki/Handbook:AMD64/Portage/Variables "Handbook:AMD64/Portage/Variables")[混合使用软件分支](/wiki/Handbook:AMD64/Portage/Branches "Handbook:AMD64/Portage/Branches")[其他工具](/wiki/Handbook:AMD64/Portage/Tools "Handbook:AMD64/Portage/Tools")[自定义软件包仓库](/wiki/Handbook:AMD64/Portage/CustomTree "Handbook:AMD64/Portage/CustomTree")[高级功能](/wiki/Handbook:AMD64/Portage/Advanced "Handbook:AMD64/Portage/Advanced")[OpenRC 网络配置](/wiki/Handbook:AMD64/Full/Networking "Handbook:AMD64/Full/Networking")[入门](/wiki/Handbook:AMD64/Networking/Introduction "Handbook:AMD64/Networking/Introduction")[高级配置](/wiki/Handbook:AMD64/Networking/Advanced "Handbook:AMD64/Networking/Advanced")[模块化网络](/wiki/Handbook:AMD64/Networking/Modular "Handbook:AMD64/Networking/Modular")[无线网络](/wiki/Handbook:AMD64/Networking/Wireless "Handbook:AMD64/Networking/Wireless")[添加功能](/wiki/Handbook:AMD64/Networking/Extending "Handbook:AMD64/Networking/Extending")[动态管理](/wiki/Handbook:AMD64/Networking/Dynamic "Handbook:AMD64/Networking/Dynamic")

## Contents

- [1系统日志工具](#.E7.B3.BB.E7.BB.9F.E6.97.A5.E5.BF.97.E5.B7.A5.E5.85.B7)
  - [1.1OpenRC](#OpenRC)
  - [1.2systemd](#systemd)
- [2可选：Cron守护进程](#.E5.8F.AF.E9.80.89.EF.BC.9ACron.E5.AE.88.E6.8A.A4.E8.BF.9B.E7.A8.8B)
  - [2.1OpenRC](#OpenRC_2)
    - [2.1.1cronie](#cronie)
    - [2.1.2可选：dcron](#.E5.8F.AF.E9.80.89.EF.BC.9Adcron)
    - [2.1.3可选：fcron](#.E5.8F.AF.E9.80.89.EF.BC.9Afcron)
    - [2.1.4可选：bcron](#.E5.8F.AF.E9.80.89.EF.BC.9Abcron)
    - [2.1.5modprobed-db users](#modprobed-db_users)
  - [2.2systemd](#systemd_2)
    - [2.2.1modprobed-db users](#modprobed-db_users_2)
- [3可选：文件索引](#.E5.8F.AF.E9.80.89.EF.BC.9A.E6.96.87.E4.BB.B6.E7.B4.A2.E5.BC.95)
- [4可选：远程 shell 访问](#.E5.8F.AF.E9.80.89.EF.BC.9A.E8.BF.9C.E7.A8.8B_shell_.E8.AE.BF.E9.97.AE)
  - [4.1OpenRC](#OpenRC_3)
  - [4.2systemd](#systemd_3)
- [5可选：Shell 补全](#.E5.8F.AF.E9.80.89.EF.BC.9AShell_.E8.A1.A5.E5.85.A8)
  - [5.1Bash](#Bash)
- [6时间同步](#.E6.97.B6.E9.97.B4.E5.90.8C.E6.AD.A5)
  - [6.1chrony](#chrony)
  - [6.2OpenRC](#OpenRC_4)
  - [6.3systemd](#systemd_4)
  - [6.4systemd-timesyncd](#systemd-timesyncd)
- [7文件系统工具](#.E6.96.87.E4.BB.B6.E7.B3.BB.E7.BB.9F.E5.B7.A5.E5.85.B7)

## 系统日志工具

### OpenRC

因为有一些工具提供给用户的功能比较类似，它们就没有包含在stage3当中。现在就是你选择安装哪一个的时候了。

首先需要决定的工具就是系统日志机制。Unix 和 Linux 在日志记录功能方面有良好的传统——如果愿意的话，可以把系统发生的所有事件都记录到日志文件中。

Gentoo提供了多种系统日志工具可供选择。包括：

- [app-admin/sysklogd](https://packages.gentoo.org/packages/app-admin/sysklogd) -提供传统的系统日志记录守护程序。默认日志配置容易学习，这个包是初学者的好选择。
- [app-admin/syslog-ng](https://packages.gentoo.org/packages/app-admin/syslog-ng) -高级系统记录器。 需要额外配置很多东西， 更高级的用户可以根据它的日志潜力选择这个包; 注意额外的配置是任何种类的智能日志记录的必要条件。
- [app-admin/metalog](https://packages.gentoo.org/packages/app-admin/metalog) -一个可以灵活配置的系统日志工具。

Gentoo ebuild 仓库内或许还有其他系统日志工具，因为可用软件包数量每天在增加。

**Tip**

如果打算使用 syslog-ng ，建议安装并且配置 [logrotate](/wiki/Logrotate "Logrotate")。syslog-ng 并没有提供系统日志文件的滚动功能。新版本（>= 2.0）的 sysklogd 会自己处理日志滚动。

要安装你所选择的系统日志工具，你可以用 emerge 命令安装它。在 OpenRC 中，使用 rc-update 将它加入默认运行级别。以下就是一个安装 [app-admin/sysklogd](https://packages.gentoo.org/packages/app-admin/sysklogd) 并作为系统 syslog 工具的例子：

`root #` `emerge --ask app-admin/sysklogd`

`root #` `rc-update add sysklogd default`

### systemd

日志记录工具一般只有使用 OpenRC 的系统需要选择，systemd 已经自带了一个叫 **systemd-journald** 的服务作为日志记录工具。systemd-journald 服务通常能够支持并处理大部分上面提到的日志记录工具的功能。换而言之，使用 systemd 运行系统作为服务管理器可以 _安全跳过_ 添加额外的系统日志记录工具。

更多关于使用 journalctl 审查系统的细节，请查看 man journalctl。

对于因为某种原因需要将日志转发到中心主机的情况，在使用 systemd 的系统上支持 _冗余_ 的系统记录机制可能很重要。但对于本手册的受众和目的来说，这是一个非典型常见且高阶的情况。因此本手册没有涵盖它。

## 可选：Cron守护进程

### OpenRC

尽管这是可选的并且不是系统所必须的，但是最好能够安装一个 cron 守护进程。

cron守护程序执行计划中的命令。 如果某些命令需要定期执行（例如每天，每周或每月），这是非常方便的。

All cron daemons support high levels of granularity for scheduled tasks, and generally include the ability to send an email or other form of notification if a scheduled task does not complete as expected.

Gentoo 提供了三个可选的 cron 守护进程，包括：

- [cronie](/wiki/Cron#cronie "Cron") ([sys-process/cronie](https://packages.gentoo.org/packages/sys-process/cronie)) \- cronie is based on the original cron and has security and configuration enhancements like the ability to use PAM and SELinux.
- [dcron](/wiki/Cron#dcron_.28Dillon.27s_Cron.29 "Cron") ([sys-process/dcron](https://packages.gentoo.org/packages/sys-process/dcron)) \- This lightweight cron daemon aims to be simple and secure, with just enough features to stay useful.
- [fcron](/wiki/Cron#fcron "Cron") ([sys-process/fcron](https://packages.gentoo.org/packages/sys-process/fcron)) \- A command scheduler with extended capabilities over cron and anacron.
- [bcron](/wiki/Cron#bcron "Cron") ([sys-process/bcron](https://packages.gentoo.org/packages/sys-process/bcron)) \- A younger cron system designed with secure operations in mind. To do this, the system is divided into several separate programs, each responsible for a separate task, with strictly controlled communications between parts.

#### cronie

下面的示例使用 [sys-process/cronie](https://packages.gentoo.org/packages/sys-process/cronie)：

`root #` `emerge --ask sys-process/cronie`

添加 cronie 作为默认系统运行级别后，将会开机自启。

`root #` `rc-update add cronie default`

#### 可选：dcron

`root #` `emerge --ask sys-process/dcron`

如果 dcron 是前置 cron agent，则需要执行额外的初始化命令：

`root #` `crontab /etc/crontab`

#### 可选：fcron

`root #` `emerge --ask sys-process/fcron`

如果选择 fcron 作为任务调度器，则需要额外的 emerge 步骤：

`root #` `emerge --config sys-process/fcron`

#### 可选：bcron

bcron 是一款新的内建特权分离的 cron agent。

`root #` `emerge --ask sys-process/bcron`

#### modprobed-db users

If using [sys-kernel/modprobed-db](https://packages.gentoo.org/packages/sys-kernel/modprobed-db) was picked as an option for manually compiling the kernel.

Now is the time to setup a crontab to periodically scan the system for hardware used.

FILE **`/etc/crontab`** **Run modprobed-db once every 6 hours**

```
0 */6 * * *     /usr/bin/modprobed-db store &> /dev/null

```

At a later date of at least a week, please visit the kernel build section of the [modprobed-db](/wiki/Modprobed-db#Building_kernels.2Fzh "Modprobed-db") article to complete the setup.

### systemd

**Tip**

systemd 不 _需要_ cron 守护进程，因为它有 _timers_，但仍然可以运行 cron 守护程序。

#### modprobed-db users

If using [sys-kernel/modprobed-db](https://packages.gentoo.org/packages/sys-kernel/modprobed-db) was picked as an option for manually compiling the kernel.

Now is the time to setup a systemd timer to periodically scan the system for hardware used.

`root #` `systemctl --user enable modprobed-db`

At a later date of at least a week, please visit the kernel build section of the [modprobed-db](/wiki/Modprobed-db#Building_kernels.2Fzh "Modprobed-db") article to complete the setup.

## 可选：文件索引

如果你想索引你的系统文件使得你能够使用locate工具很快定位它们，你需要安装[sys-apps/mlocate](https://packages.gentoo.org/packages/sys-apps/mlocate)。

`root #` `emerge --ask sys-apps/mlocate`

## 可选：远程 shell 访问

**Tip**

opensshd 默认配置不允许远程登录 root 用户。如果有需要的话，请在安装后 [创建一个非 root 用户](/wiki/FAQ#How_do_I_add_a_normal_user.3F "FAQ")，并且配置合适的权限，或者调整 /etc/ssh/sshd\_config，使其允许 root 远程登录。

要在安装后远程访问系统，必须配置为在启动时运行sshd。

For more in-depth details on the configuration of SSH, refer to the [SSH](/wiki/SSH "SSH") article.

### OpenRC

在 OpenRC 上将 sshd init 脚本添加到默认运行级别：

`root #` `rc-update add sshd default`

`root #` `rc-update add sshd default`

如果需要访问串行控制台（在远程服务器的情况下这是可能的），必须配置 agetty。

在 /etc/inittab 中取消串行控制台部分的注释：

`root #` `nano -w /etc/inittab`

```
# SERIAL CONSOLES
s0:12345:respawn:/sbin/agetty 9600 ttyS0 vt100
s1:12345:respawn:/sbin/agetty 9600 ttyS1 vt100

```

### systemd

运行以下命令来启用 SSH 服务端：

`root #` `systemctl enable sshd`

`root #` `systemctl enable sshd`

运行以下命令来启用串行控制台支持：

`root #` `systemctl enable getty@tty1.service`

`root #` `systemctl enable getty@tty1.service`

## 可选：Shell 补全

### Bash

Bash 是Gentoo系统里默认的shell，所以安装补全建议插件支持可以更有效率，更便捷。[app-shells/bash-completion](https://packages.gentoo.org/packages/app-shells/bash-completion) 将会提供Gentooo系统特有的和其他常用命令和工具的可用的建议补全

`root #` `emerge --ask app-shells/bash-completion`

完成安装后，可以使用eselect命令管理bash的补全支持的项目 。 查阅关于bash补全的文章获取更多详情： [Shell completion integrations section](/wiki/Bash#Shell_completion_integrations.2Fzh "Bash")

## 时间同步

**Important**

Systems without a functioning [Real-Time Clock (RTC)](/wiki/System_time#Software_clock_vs_Hardware_clock "System time") must set the [system time](/wiki/System_time "System time") at every system start, and on regular intervals thereafter.

使用一些同步系统时钟的方法很重要。通常使用 [NTP](/wiki/NTP "NTP") 协议和软件。有一些 NTP 协议的其他实现，例如 [Chrony](/wiki/Chrony "Chrony")。

The clock is usually synchronized via the [Network Time Protocol](/wiki/Network_Time_Protocol "Network Time Protocol"), with an implementation such as [chrony](/wiki/Chrony "Chrony").

### chrony

设置 Chrony，例如

`root #` `emerge --ask net-misc/chrony`

`root #` `emerge --ask net-misc/chrony`

See the [chrony](/wiki/Chrony "Chrony") article for further information, for example if more advanced configurations are required.

### OpenRC

在 OpenRC 上，运行：

`root #` `rc-update add chronyd default`

`root #` `rc-update add chronyd default`

### systemd

在 systemd 上，运行：

`root #` `systemctl enable chronyd.service`

`root #` `systemctl enable chronyd.service`

### systemd-timesyncd

另外，systemd 用户可能希望使用默认安装的更简单的 systemd-timesyncd SNTP 客户端。

`root #` `systemctl enable systemd-timesyncd.service`

`root #` `systemctl enable systemd-timesyncd.service`

## 文件系统工具

根据你所使用的文件系统的不同，可能需要安装需要的文件系统工具（用于检查文件系统完整性、（重新）格式化文件系统等）。请注意，作为 [@system 集](/wiki/System_set_(Portage) "System set (Portage)") 一部分的，已经安装了 ext4 用户空间的工具 ([sys-fs/e2fsprogs](https://packages.gentoo.org/packages/sys-fs/e2fsprogs))。

以下的表格列出了特定文件系统所需要安装的工具。

文件系统
软件包
XFS
[sys-fs/xfsprogs](https://packages.gentoo.org/packages/sys-fs/xfsprogs)ext4
[sys-fs/e2fsprogs](https://packages.gentoo.org/packages/sys-fs/e2fsprogs)VFAT (FAT32, ...)
[sys-fs/dosfstools](https://packages.gentoo.org/packages/sys-fs/dosfstools)Btrfs
[sys-fs/btrfs-progs](https://packages.gentoo.org/packages/sys-fs/btrfs-progs)ZFS
[sys-fs/zfs](https://packages.gentoo.org/packages/sys-fs/zfs)JFS
[sys-fs/jfsutils](https://packages.gentoo.org/packages/sys-fs/jfsutils)

It's recommended that [sys-block/io-scheduler-udev-rules](https://packages.gentoo.org/packages/sys-block/io-scheduler-udev-rules) be installed for the correct scheduler behavior with e.g. nvme devices:

`root #` `emerge --ask sys-block/io-scheduler-udev-rules`

**Tip**

获取更多关于Gentoo上文件系统的信息请看 [文件系统文章](/wiki/Filesystem "Filesystem")。

现在继续 [配置引导启动程序](/wiki/Handbook:AMD64/Installation/Bootloader "Handbook:AMD64/Installation/Bootloader")。

[← 配置系统](/wiki/Handbook:AMD64/Installation/System/zh-cn "Previous part") [Home](/wiki/Handbook:AMD64/zh-cn "Handbook:AMD64/zh-cn") [配置引导加载程序 →](/wiki/Handbook:AMD64/Installation/Bootloader/zh-cn "Next part")