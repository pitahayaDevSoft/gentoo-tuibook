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
- čeština
- [русский](/wiki/Handbook:AMD64/Installation/Tools/ru "Handbook:AMD64/Installation/Tools/ru (100% translated)")
- [தமிழ்](/wiki/Handbook:AMD64/Installation/Tools/ta "Handbook:AMD64/Installation/Tools/ta (100% translated)")
- [中文](/wiki/Handbook:AMD64/Installation/Tools/zh "Handbook:AMD64/Installation/Tools/zh (100% translated)")
- [中文（中国大陆）‎](/wiki/Handbook:AMD64/Installation/Tools/zh-cn "Handbook:AMD64/Installation/Tools/zh-cn (100% translated)")
- [日本語](/wiki/Handbook:AMD64/Installation/Tools/ja "Handbook:AMD64/Installation/Tools/ja (100% translated)")
- [한국어](/wiki/Handbook:AMD64/Installation/Tools/ko "Handbook:AMD64/Installation/Tools/ko (100% translated)")

[AMD64 Handbook](/wiki/Handbook:AMD64/cs "Handbook:AMD64/cs")[Instalace](/wiki/Handbook:AMD64/Full/Installation/cs "Handbook:AMD64/Full/Installation/cs")[O instalaci](/wiki/Handbook:AMD64/Installation/About/cs "Handbook:AMD64/Installation/About/cs")[Výběr média](/wiki/Handbook:AMD64/Installation/Media/cs "Handbook:AMD64/Installation/Media/cs")[Konfigurace sítě](/wiki/Handbook:AMD64/Installation/Networking/cs "Handbook:AMD64/Installation/Networking/cs")[Příprava disků](/wiki/Handbook:AMD64/Installation/Disks/cs "Handbook:AMD64/Installation/Disks/cs")[Instalace stage3](/wiki/Handbook:AMD64/Installation/Stage/cs "Handbook:AMD64/Installation/Stage/cs")[Instalace základního systému](/wiki/Handbook:AMD64/Installation/Base/cs "Handbook:AMD64/Installation/Base/cs")[Konfigurace jádra](/wiki/Handbook:AMD64/Installation/Kernel/cs "Handbook:AMD64/Installation/Kernel/cs")[Konfigurace systému](/wiki/Handbook:AMD64/Installation/System/cs "Handbook:AMD64/Installation/System/cs")Instalace nástrojů[Konfigurace zavaděče](/wiki/Handbook:AMD64/Installation/Bootloader/cs "Handbook:AMD64/Installation/Bootloader/cs")[Dokončení](/wiki/Handbook:AMD64/Installation/Finalizing/cs "Handbook:AMD64/Installation/Finalizing/cs")[Práce s Gentoo](/wiki/Handbook:AMD64/Full/Working/cs "Handbook:AMD64/Full/Working/cs")[Úvod do Portage](/wiki/Handbook:AMD64/Working/Portage/cs "Handbook:AMD64/Working/Portage/cs")[Přepínače USE](/wiki/Handbook:AMD64/Working/USE/cs "Handbook:AMD64/Working/USE/cs")[Funkce portage](/wiki/Handbook:AMD64/Working/Features/cs "Handbook:AMD64/Working/Features/cs")[Systém initskriptů](/wiki/Handbook:AMD64/Working/Initscripts/cs "Handbook:AMD64/Working/Initscripts/cs")[Proměnné prostředí](/wiki/Handbook:AMD64/Working/EnvVar/cs "Handbook:AMD64/Working/EnvVar/cs")[Práce s Portage](/wiki/Handbook:AMD64/Full/Portage/cs "Handbook:AMD64/Full/Portage/cs")[Soubory a adresáře](/wiki/Handbook:AMD64/Portage/Files/cs "Handbook:AMD64/Portage/Files/cs")[Proměnné](/wiki/Handbook:AMD64/Portage/Variables/cs "Handbook:AMD64/Portage/Variables/cs")[Mísení softwarových větví](/wiki/Handbook:AMD64/Portage/Branches/cs "Handbook:AMD64/Portage/Branches/cs")[Doplňkové nástroje](/wiki/Handbook:AMD64/Portage/Tools/cs "Handbook:AMD64/Portage/Tools/cs")[Vlastní strom Portage](/wiki/Handbook:AMD64/Portage/CustomTree/cs "Handbook:AMD64/Portage/CustomTree/cs")[Pokročilé funkce](/wiki/Handbook:AMD64/Portage/Advanced/cs "Handbook:AMD64/Portage/Advanced/cs")[Konfigurace sítě](/wiki/Handbook:AMD64/Full/Networking/cs "Handbook:AMD64/Full/Networking/cs")[Začínáme](/wiki/Handbook:AMD64/Networking/Introduction/cs "Handbook:AMD64/Networking/Introduction/cs")[Pokročilá konfigurace](/wiki/Handbook:AMD64/Networking/Advanced/cs "Handbook:AMD64/Networking/Advanced/cs")[Modulární síťování](/wiki/Handbook:AMD64/Networking/Modular/cs "Handbook:AMD64/Networking/Modular/cs")[Bezdrátové sítě](/wiki/Handbook:AMD64/Networking/Wireless/cs "Handbook:AMD64/Networking/Wireless/cs")[Přidání funkcí](/wiki/Handbook:AMD64/Networking/Extending/cs "Handbook:AMD64/Networking/Extending/cs")[Dynamická správa](/wiki/Handbook:AMD64/Networking/Dynamic/cs "Handbook:AMD64/Networking/Dynamic/cs")

## Contents

- [1Systémový záznamník](#Syst.C3.A9mov.C3.BD_z.C3.A1znamn.C3.ADk)
  - [1.1OpenRC](#OpenRC)
  - [1.2systemd](#systemd)
- [2Volitelné: Démon cron](#Voliteln.C3.A9:_D.C3.A9mon_cron)
  - [2.1OpenRC](#OpenRC_2)
    - [2.1.1Default: cronie](#Default:_cronie)
    - [2.1.2Alternative: dcron](#Alternative:_dcron)
    - [2.1.3Alternative: fcron](#Alternative:_fcron)
    - [2.1.4Alternative: bcron](#Alternative:_bcron)
    - [2.1.5modprobed-db users](#modprobed-db_users)
  - [2.2systemd](#systemd_2)
    - [2.2.1modprobed-db users](#modprobed-db_users_2)
- [3Volitelné: Indexování souborů](#Voliteln.C3.A9:_Indexov.C3.A1n.C3.AD_soubor.C5.AF)
- [4Volitelné: Vzdálený přístup](#Voliteln.C3.A9:_Vzd.C3.A1len.C3.BD_p.C5.99.C3.ADstup)
  - [4.1OpenRC](#OpenRC_3)
  - [4.2systemd](#systemd_3)
- [5Optional: Shell completion](#Optional:_Shell_completion)
  - [5.1Bash](#Bash)
- [6Suggested: Time synchronization](#Suggested:_Time_synchronization)
  - [6.1chrony](#chrony)
    - [6.1.1OpenRC](#OpenRC_4)
    - [6.1.2systemd](#systemd_4)
  - [6.2systemd-timesyncd](#systemd-timesyncd)
- [7Nástroje souborového systému](#N.C3.A1stroje_souborov.C3.A9ho_syst.C3.A9mu)

## Systémový záznamník

### OpenRC

Některé nástroje v archivu stage3 chybí, protože několik balíčků poskytuje tu samou funkcionalitu. Je na uživateli, aby vybral, který z nich nainstaluje.

První rozhodnutí se týká nástroje, který poskytuje záznamové funkce. Unix a Linux se vyznačuje v oblasti záznamů skvělými možnostmi - je-li to třeba, pak vše, co se v systémů děje, lze zapsat do logu. To se děje prostřednictvím systémového záznamníku.

Gentoo poskytuje několik utilit systémových záznamníků. Ty zahrnují:

- [app-admin/sysklogd](https://packages.gentoo.org/packages/app-admin/sysklogd) \- Poskytuje tradiční sadu démonů systémových záznamů. Výchozí konfigurace záznamů funguje přímo po nainstalování, což z něj dělá dobrou volbu pro začátečníky.
- [app-admin/syslog-ng](https://packages.gentoo.org/packages/app-admin/syslog-ng) \- Pokročilý systémový záznamník. Cokoli jiného, než záznam do jednoho velkého souboru, vyžaduje dodatečnou konfiguraci. Pokročilejší uživatelé mohou využít tuto aplikaci kvůli potenciálu; Vemte na vědomí, že dodatečná konfigurace je pro jakýkoli druh chytrého záznamenávání nezbytná.
- [app-admin/metalog](https://packages.gentoo.org/packages/app-admin/metalog) \- Systémový záznamník s širokými možnostmi nastavení.

Další programy jsou v nabídce Portage také - množství dostupných balíčků se denně navyšuje.

**Tip**

Pokud použijete sysklogd nebo syslog-ng, doporučuje se po nich nainstalovat a nakonfigurovat balíček [logrotate](/wiki/Logrotate "Logrotate"), jelikož tyto systémové záznamníky neposkytují mechanismus pro rotaci systémových souborů.

Zvolený systémový záznamník nainstalujte pomocí emerge a přidejte jej do výchozí úrovně běhu s pomocí rc-update. Následující příklad nainstaluje [app-admin/sysklogd](https://packages.gentoo.org/packages/app-admin/sysklogd):

`root #` `emerge --ask app-admin/sysklogd`

`root #` `rc-update add sysklogd default`

### systemd

While a selection of logging mechanisms are presented for OpenRC-based systems, systemd includes a built-in logger called the **systemd-journald** service. The systemd-journald service is capable of handling most of the logging functionality outlined in the previous system logger section. That is to say, the majority of installations that will run systemd as the system and service manager can _safely skip_ adding a additional syslog utilities.

See man journalctl for more details on using journalctl to query and review the systems logs.

For a number of reasons, such as the case of forwarding logs to a central host, it may be important to include _redundant_ system logging mechanisms on a systemd-based system. This is a irregular occurrence for the handbook's typical audience and considered an advanced use case. It is therefore not covered by the handbook.

## Volitelné: Démon cron

### OpenRC

Další na řadě je démon cron.

Démon cron spouští naplánované příkazy. To se hodí, pokud je potřeba spouštět některý příkaz v pravidelných insternvalech (například denně, týdně nebo měsíčně).

All cron daemons support high levels of granularity for scheduled tasks, and generally include the ability to send an email or other form of notification if a scheduled task does not complete as expected.

Gentoo poskytuje několik možných démonů cronu, včetně [sys-proces/bcron](https://packages.gentoo.org/packages/sys-proces/bcron), [sys-proces/dcron](https://packages.gentoo.org/packages/sys-proces/dcron), [sys-process/fcron](https://packages.gentoo.org/packages/sys-process/fcron) a [sys-process/cronie](https://packages.gentoo.org/packages/sys-process/cronie). Nainstalování jednoho z nich je podobné, jako v případě systémového záznamníku. Následující příklad používá [sys-process/cronie](https://packages.gentoo.org/packages/sys-process/cronie):

- [cronie](/wiki/Cron#cronie "Cron") ([sys-process/cronie](https://packages.gentoo.org/packages/sys-process/cronie)) \- cronie is based on the original cron and has security and configuration enhancements like the ability to use PAM and SELinux.
- [dcron](/wiki/Cron#dcron_.28Dillon.27s_Cron.29 "Cron") ([sys-process/dcron](https://packages.gentoo.org/packages/sys-process/dcron)) \- This lightweight cron daemon aims to be simple and secure, with just enough features to stay useful.
- [fcron](/wiki/Cron#fcron "Cron") ([sys-process/fcron](https://packages.gentoo.org/packages/sys-process/fcron)) \- A command scheduler with extended capabilities over cron and anacron.
- [bcron](/wiki/Cron#bcron "Cron") ([sys-process/bcron](https://packages.gentoo.org/packages/sys-process/bcron)) \- A younger cron system designed with secure operations in mind. To do this, the system is divided into several separate programs, each responsible for a separate task, with strictly controlled communications between parts.

#### Default: cronie

The following example uses [sys-process/cronie](https://packages.gentoo.org/packages/sys-process/cronie) (not to be confused with the [NTP](/wiki/Network_Time_Protocol "Network Time Protocol") daemon named [chrony](/wiki/Chrony "Chrony")):

`root #` `emerge --ask sys-process/cronie`

`root #` `rc-update add cronie default`

`root #` `rc-update add cronie default`

#### Alternative: dcron

`root #` `emerge --ask sys-process/dcron`

Pokud je použit dcron nebo fcron, je třeba spustit dodatečný inicializační příkaz:

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

At a later date of at least a week, please visit the kernel build section of the [modprobed-db](/wiki/Modprobed-db#Building_kernels.2Fcs "Modprobed-db") article to complete the setup.

### systemd

Similar to system logging, systemd-based systems include support for scheduled tasks out-of-the-box in the form of _timers_. systemd timers can run at a system-level or a user-level and include the same functionality that a traditional cron daemon would provide. Unless redundant capabilities are necessary, installing an additional task scheduler such as a cron daemon is generally unnecessary and can be safely skipped.

#### modprobed-db users

If using [sys-kernel/modprobed-db](https://packages.gentoo.org/packages/sys-kernel/modprobed-db) was picked as an option for manually compiling the kernel.

Now is the time to setup a systemd timer to periodically scan the system for hardware used.

`root #` `systemctl --user enable modprobed-db`

At a later date of at least a week, please visit the kernel build section of the [modprobed-db](/wiki/Modprobed-db#Building_kernels.2Fcs "Modprobed-db") article to complete the setup.

## Volitelné: Indexování souborů

Za účelem indexace souborového systému a tím rychlejší schopnosti vyhledávání, nainstalujte [sys-apps/mlocate](https://packages.gentoo.org/packages/sys-apps/mlocate).

`root #` `emerge --ask sys-apps/mlocate`

## Volitelné: Vzdálený přístup

**Tip**

opensshd's default configuration does not allow root to login as a remote user. Please [create a non-root user](/wiki/FAQ#How_do_I_add_a_normal_user.3F "FAQ") and configure it appropriately to allow access post-installation if required, or adjust /etc/ssh/sshd\_config to allow root.

Abyste byli schopní přistupovat k systému po instalaci vzdáleným způsobem, přidejte do výchozí úrovně běhu init skript sshd:

For more in-depth details on the configuration of SSH, refer to the [SSH](/wiki/SSH "SSH") article.

### OpenRC

`root #` `rc-update add sshd default`

`root #` `rc-update add sshd default`

Pokud je potřeba přístup přes sériovou konzoli (což je možné v případě vzdálených serverů), odkomentujte sekci sériové konzole v souboru /etc/inittab:

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

Post installation, bash completion for specific commands can managed through eselect. See the [Shell completion integrations section](/wiki/Bash#Shell_completion_integrations.2Fcs "Bash") of the bash article for more details.

## Suggested: Time synchronization

**Důležité**

Systems without a functioning [Real-Time Clock (RTC)](/wiki/System_time/cs#Software_clock_vs_Hardware_clock "System time/cs") must set the [system time](/wiki/System_time/cs "System time/cs") at every system start, and on regular intervals thereafter.

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

## Nástroje souborového systému

V závislosti na použitých souborových systémech, je nezbytné nainstalovat utility souborových systémů (pro kontrolu integrity souborového systému, vytváření dalších souborových systémů atd.). Vezměte na vědomí, nástroje pro správu souborových systémů ext2, ext3 nebo ext4 ([sys-fs/e2fsprogs](https://packages.gentoo.org/packages/sys-fs/e2fsprogs)) jsou již nainstalovány jakou součást [setu @system](/wiki/System_set_(Portage) "System set (Portage)").

Následující tabulka obsahuje nástroje, které nainstalujte v případě, že je použit určitý souborový systém:

Souborový systém
Balíček
Ext2, 3 a 4
[sys-fs/e2fsprogs](https://packages.gentoo.org/packages/sys-fs/e2fsprogs)XFS
[sys-fs/xfsprogs](https://packages.gentoo.org/packages/sys-fs/xfsprogs)ReiserFS
[sys-fs/reiserfsprogs](https://packages.gentoo.org/packages/sys-fs/reiserfsprogs)JFS
[sys-fs/jfsutils](https://packages.gentoo.org/packages/sys-fs/jfsutils)VFAT (FAT32, ...)
[sys-fs/dosfstools](https://packages.gentoo.org/packages/sys-fs/dosfstools)Btrfs
[sys-fs/btrfs-progs](https://packages.gentoo.org/packages/sys-fs/btrfs-progs)

It's recommended that [sys-block/io-scheduler-udev-rules](https://packages.gentoo.org/packages/sys-block/io-scheduler-udev-rules) be installed for the correct scheduler behavior with e.g. nvme devices:

`root #` `emerge --ask sys-block/io-scheduler-udev-rules`

**Tip**

Více informací o souborových systémech v Gentoo najdete v [článku o souborových systémech](/wiki/Filesystem "Filesystem").

Nyní pokračujte na [Nastavení zavaděče](/wiki/Handbook:AMD64/Installation/Bootloader/cs "Handbook:AMD64/Installation/Bootloader/cs").

[← Konfirauce systému](/wiki/Handbook:AMD64/Installation/System/cs "Previous part") [Home](/wiki/Handbook:AMD64 "Handbook:AMD64") [Konfigurace zavaděče →](/wiki/Handbook:AMD64/Installation/Bootloader/cs "Next part")