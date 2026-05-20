# Tools

Other languages:

- Deutsch
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
- [한국어](/wiki/Handbook:HPPA/Installation/Tools/ko "Handbook:HPPA/Installation/Tools/ko (100% translated)")

[HPPA Handbuch](/wiki/Handbook:HPPA/de "Handbook:HPPA/de")[Installation](/wiki/Handbook:HPPA/Full/Installation/de "Handbook:HPPA/Full/Installation/de")[Über die Installation](/wiki/Handbook:HPPA/Installation/About/de "Handbook:HPPA/Installation/About/de")[Auswahl des Mediums](/wiki/Handbook:HPPA/Installation/Media/de "Handbook:HPPA/Installation/Media/de")[Konfiguration des Netzwerks](/wiki/Handbook:HPPA/Installation/Networking/de "Handbook:HPPA/Installation/Networking/de")[Vorbereiten der Festplatte(n)](/wiki/Handbook:HPPA/Installation/Disks/de "Handbook:HPPA/Installation/Disks/de")[Installation des Stage Archivs](/wiki/Handbook:HPPA/Installation/Stage/de "Handbook:HPPA/Installation/Stage/de")[Installation des Basissystems](/wiki/Handbook:HPPA/Installation/Base/de "Handbook:HPPA/Installation/Base/de")[Konfiguration des Kernels](/wiki/Handbook:HPPA/Installation/Kernel/de "Handbook:HPPA/Installation/Kernel/de")[Konfiguration des Systems](/wiki/Handbook:HPPA/Installation/System/de "Handbook:HPPA/Installation/System/de")Installation der Tools[Konfiguration des Bootloaders](/wiki/Handbook:HPPA/Installation/Bootloader/de "Handbook:HPPA/Installation/Bootloader/de")[Abschluss](/wiki/Handbook:HPPA/Installation/Finalizing/de "Handbook:HPPA/Installation/Finalizing/de")[Arbeiten mit Gentoo](/wiki/Handbook:HPPA/Full/Working/de "Handbook:HPPA/Full/Working/de")[Portage-Einführung](/wiki/Handbook:HPPA/Working/Portage/de "Handbook:HPPA/Working/Portage/de")[USE-Flags](/wiki/Handbook:HPPA/Working/USE/de "Handbook:HPPA/Working/USE/de")[Portage-Features](/wiki/Handbook:HPPA/Working/Features/de "Handbook:HPPA/Working/Features/de")[Initskript-System](/wiki/Handbook:HPPA/Working/Initscripts/de "Handbook:HPPA/Working/Initscripts/de")[Umgebungsvariablen](/wiki/Handbook:HPPA/Working/EnvVar/de "Handbook:HPPA/Working/EnvVar/de")[Arbeiten mit Portage](/wiki/Handbook:HPPA/Full/Portage/de "Handbook:HPPA/Full/Portage/de")[Dateien und Verzeichnisse](/wiki/Handbook:HPPA/Portage/Files/de "Handbook:HPPA/Portage/Files/de")[Variablen](/wiki/Handbook:HPPA/Portage/Variables/de "Handbook:HPPA/Portage/Variables/de")[Mischen von Softwarezweigen](/wiki/Handbook:HPPA/Portage/Branches/de "Handbook:HPPA/Portage/Branches/de")[Zusätzliche Tools](/wiki/Handbook:HPPA/Portage/Tools/de "Handbook:HPPA/Portage/Tools/de")[Eigener Portage-Tree](/wiki/Handbook:HPPA/Portage/CustomTree/de "Handbook:HPPA/Portage/CustomTree/de")[Erweiterte Portage-Features](/wiki/Handbook:HPPA/Portage/Advanced/de "Handbook:HPPA/Portage/Advanced/de")[Netzwerk-Konfiguration](/wiki/Handbook:HPPA/Full/Networking/de "Handbook:HPPA/Full/Networking/de")[Zu Beginn](/wiki/Handbook:HPPA/Networking/Introduction/de "Handbook:HPPA/Networking/Introduction/de")[Fortgeschrittene Konfiguration](/wiki/Handbook:HPPA/Networking/Advanced/de "Handbook:HPPA/Networking/Advanced/de")[Modulare Vernetzung](/wiki/Handbook:HPPA/Networking/Modular/de "Handbook:HPPA/Networking/Modular/de")[Drahtlose Netzwerke](/wiki/Handbook:HPPA/Networking/Wireless/de "Handbook:HPPA/Networking/Wireless/de")[Funktionalität hinzufügen](/wiki/Handbook:HPPA/Networking/Extending/de "Handbook:HPPA/Networking/Extending/de")[Dynamisches Management](/wiki/Handbook:HPPA/Networking/Dynamic/de "Handbook:HPPA/Networking/Dynamic/de")

## Contents

- [1Syslog Daemon](#Syslog_Daemon)
  - [1.1OpenRC](#OpenRC)
  - [1.2systemd](#systemd)
- [2Optional: Cron Daemon](#Optional:_Cron_Daemon)
  - [2.1OpenRC](#OpenRC_2)
    - [2.1.1cronie](#cronie)
    - [2.1.2Alternative: dcron](#Alternative:_dcron)
    - [2.1.3Alternative: fcron](#Alternative:_fcron)
    - [2.1.4Alternative: bcron](#Alternative:_bcron)
    - [2.1.5modprobed-db users](#modprobed-db_users)
  - [2.2systemd](#systemd_2)
    - [2.2.1modprobed-db users](#modprobed-db_users_2)
- [3Optional: Datei-Index](#Optional:_Datei-Index)
- [4Optional: Remote Zugriff](#Optional:_Remote_Zugriff)
  - [4.1OpenRC](#OpenRC_3)
  - [4.2systemd](#systemd_3)
- [5Optional: Shell-Vervollständigung](#Optional:_Shell-Vervollst.C3.A4ndigung)
  - [5.1Bash](#Bash)
- [6Zeitsynchronisation](#Zeitsynchronisation)
  - [6.1chrony](#chrony)
  - [6.2OpenRC](#OpenRC_4)
  - [6.3systemd](#systemd_4)
  - [6.4systemd-timesyncd](#systemd-timesyncd)
- [7Dateisystemwerkzeuge](#Dateisystemwerkzeuge)

## Syslog Daemon

### OpenRC

Einige Tools fehlen in dem Stage Tar-Archiv, weil es mehrere Pakete gibt, die die gleiche Funktionalität bereitstellen. Der Anwender kann wählen, welches dieser Pakete er installieren möchte.

Das erste Werkzeug, bei dem eine Auswahl getroffen werden muss, ist der Logging-Mechanismus. UNIX und Linux bieten hervorragende Unterstützung für Logging. Falls notwendig, kann alles, was auf dem System passiert, in Log-Dateien protokolliert werden.

Gentoo bietet verschiedene Syslog Daemons, unter anderem:

- [app-admin/sysklogd](https://packages.gentoo.org/packages/app-admin/sysklogd) \- Das Paket beinhaltet das traditionelle Set von Syslog Diensten. Die mitgelieferte Standard-Konfiguration funktioniert ohne zusätzliche Konfigurationsarbeiten. Deshalb ist dieses Paket eine gute Wahl für Anfänger.
- [app-admin/syslog-ng](https://packages.gentoo.org/packages/app-admin/syslog-ng) \- Ein fortgeschrittener Syslog Daemon, der für fortgeschrittene Anwender gedacht ist, die das Logging feiner steuern und zusätzliche Funktionen nutzen wollen. Er benötigt zusätzliche Konfigurationsaufwand, wenn in mehr als eine Datei protokolliert werden soll.
- [app-admin/metalog](https://packages.gentoo.org/packages/app-admin/metalog) \- Ein hochgradig konfigurierbarer Syslog Daemon.

Es kann sein, dass im Gentoo-Ebuild-Repositorium auch andere Systemloggerwerkzeuge verfügbar sind, da die Anzahl der verfügbaren Pakete täglich steigt.

**Tipp**

Wenn syslog-ng verwendet werden soll, wird empfohlen, auch das Paket [logrotate](/wiki/Logrotate "Logrotate") zu installieren, weil syslog-ng keine Funktionen zum Rotieren und Löschen von Log-Dateien enthält. Neuere Versionen von sysklogd (>= 2.0) enthalten Logrotate-Funktionen.

Wenn Sie einen Syslog Daemon ausgewählt haben, installieren Sie ihn mit emerge. Wenn Sie OpenRC verwenden, fügen Sie ihn mit rc-update zum Runlevel "default" hinzu. Das folgende Beispiel installiert und aktiviert [app-admin/sysklogd](https://packages.gentoo.org/packages/app-admin/sysklogd) als System Syslog-Werkzeug:

`root #` `emerge --ask app-admin/sysklogd`

`root #` `rc-update add sysklogd default`

### systemd

Während eine Auswahl von Protokollierungsmechanismen für OpenRC-basierte Systeme vorgestellt wird, enthält systemd einen eingebauten Logger den **systemd-journald**-Dienst. Der systemd-journald-Dienst ist in der Lage, die meisten der im vorherigen Abschnitt über Systemlogger beschriebenen Logging-Funktionen auszuführen. Das heißt dass die meisten Installationen, die systemd als System-und Dienstmanager verwenden, das Hinzufügen eines zusätzlichen Syslog-Dienstes getrost überspringen können.

Siehe man journalctl für weitere Details zur Verwendung von journalctl zur Abfrage und Überprüfung der Systemprotokolle.

Aus einer Reihe von Gründen, wie z.B. der Weiterleitung von Protokollen an einen zentralen Host, kann es wichtig sein, "redundante" Systemprotokollierungsmechanismen in ein systemd-basiertes System einzubinden. Dies ist für die typische Zielgruppe des Handbuchs unüblich und gilt als fortgeschrittener Anwendungsfall. Er wird daher im Handbuch nicht behandelt.

## Optional: Cron Daemon

### OpenRC

Die Installation eines Cron-Daemons ist optional und wird nicht auf jedem System benötigt. Auf den meisten Systemen ist die Installation eines Cron-Daemons jedoch sinnvoll.

Ein Cron-Daemon führt Befehle in geplanten Intervallen aus. Das können tägliche, wöchentliche oder monatliche Intervalle sein, einmal jeden Dienstag, einmal jede zweite Woche, usw. Ein kluger Systemadministrator wird den Cron-Daemon nutzen, um routinemäßige Systemwartungsaufgaben zu automatisieren.

Alle Cron-Daemons unterstützen eine hohe Granularität für geplante Aufgaben und bieten im Allgemeinen die Möglichkeit, eine E-Mail oder eine andere Form der Benachrichtigung zu senden, wenn eine geplante Aufgabe nicht wie erwartet abgeschlossen wird.

Gentoo bietet verschiedene Cron Daemons an, unter anderem:

- [sys-process/cronie](https://packages.gentoo.org/packages/sys-process/cronie) \- cronie basiert auf dem ursprünglichen cron und verfügt über Sicherheits- und Konfigurationsverbesserungen wie die Möglichkeit, PAM und SELinux zu verwenden.
- [sys-process/dcron](https://packages.gentoo.org/packages/sys-process/dcron) \- Dieser leichgewichtige Cron-Daemon soll einfach und sicher sein, mit gerade genug Funktionen, um nützlich zu bleiben.
- [sys-process/fcron](https://packages.gentoo.org/packages/sys-process/fcron) \- Ein Befehl-Scheduler mit erweiterten Fähigkeiten gegenüber cron und anacron.
- [sys-process/bcron](https://packages.gentoo.org/packages/sys-process/bcron) \- Ein jüngeres Cron-System, das mit Blick auf sichere Abläufe entwickelt wurde. Zu diesem Zweck ist das System in mehrere separate Programme eingeteilt, von denen jedes für eine eigene Aufgabe zuständig ist, wobei die Kommunikation zwischen den Teilen streng kontrolliert wird.

#### cronie

Das folgende Beispiel verwendet [sys-process/cronie](https://packages.gentoo.org/packages/sys-process/cronie):

`root #` `emerge --ask sys-process/cronie`

Fügen Sie cronie zum Standard-Runlevel des Systems hinzu, so dass es beim Einschalten automatisch gestartet wird:

`root #` `rc-update add cronie default`

#### Alternative: dcron

`root #` `emerge --ask sys-process/dcron`

Wenn dcron der zukünftige Cron-Agent ist, muss ein zusätzlicher Initialisierungsbefehl ausgeführt werden:

`root #` `crontab /etc/crontab`

#### Alternative: fcron

`root #` `emerge --ask sys-process/fcron`

Wenn fcron der ausgewählte Task-Handler ist, ist ein zusätzlicher emerge-Schritt erforderlich:

`root #` `emerge --config sys-process/fcron`

#### Alternative: bcron

bcron ist ein jüngerer Cron-Agent mit eingebauter Privilegientrennung.

`root #` `emerge --ask sys-process/bcron`

#### modprobed-db users

If using [sys-kernel/modprobed-db](https://packages.gentoo.org/packages/sys-kernel/modprobed-db) was picked as an option for manually compiling the kernel.

Now is the time to setup a crontab to periodically scan the system for hardware used.

DATEI **`/etc/crontab`** **Run modprobed-db once every 6 hours**

```
0 */6 * * *     /usr/bin/modprobed-db store &> /dev/null

```

At a later date of at least a week, please visit the kernel build section of the [modprobed-db](/wiki/Modprobed-db#Building_kernels.2Fde "Modprobed-db") article to complete the setup.

### systemd

Ähnlich wie bei der Systemprotokollierung bieten systemd-basierte Systeme standardmäßig Unterstützung für geplante Aufgaben in Form von _Timern_. systemd-Timer können auf System- oder Benutzerebenen ausgeführt werden und bieten die gleiche Funktionalität wie ein herkömmlicher Cron-Daemon. Sofern keine redundanten Fähigkeiten erforderlich sind, ist die Installation eines zusätzlichen Aufgabenplaners wie eines Cron-Daemons im Allgemeinen unnötig und kann getrost übersprungen werden.

#### modprobed-db users

If using [sys-kernel/modprobed-db](https://packages.gentoo.org/packages/sys-kernel/modprobed-db) was picked as an option for manually compiling the kernel.

Now is the time to setup a systemd timer to periodically scan the system for hardware used.

`root #` `systemctl --user enable modprobed-db`

At a later date of at least a week, please visit the kernel build section of the [modprobed-db](/wiki/Modprobed-db#Building_kernels.2Fde "Modprobed-db") article to complete the setup.

## Optional: Datei-Index

Mit Hilfe des Pakets [sys-apps/mlocate](https://packages.gentoo.org/packages/sys-apps/mlocate) kann man einen Index des Dateisystems erstellen und schnell nach Dateien suchen.

`root #` `emerge --ask sys-apps/mlocate`

## Optional: Remote Zugriff

**Tipp**

Die Standardkonfiguration von opensshd erlaubt es root nicht, sich als remote-Benutzer anzumelden. Bitte [erstellen Sie einen Nicht-Root-Benutzer](/wiki/FAQ/de#How_do_I_add_a_normal_user.3F "FAQ/de") und konfigurieren Sie ihn entsprechend, um den Zugriff nach der Installation zu erlauben, falls erforderlich, oder passen Sie /etc/ssh/shhd\_config an, um root zu erlauben.

Wenn Sie sich von Remote Systemen über SSH bei Ihrem neu installierten System anmelden wollen, muss sshd so konfiguriert werden, dass es beim Booten startet.

For more in-depth details on the configuration of SSH, refer to the [SSH](/wiki/SSH/de "SSH/de") article.

### OpenRC

Um das sshd Init-Script unter OpenRC zum Runlevel "default" hinzufügen:

`root #` `rc-update add sshd default`

`root #` `rc-update add sshd default`

Wenn Sie sich über die serielle Schnittstelle bei Ihrem neu installierten System anmelden wollen, muss agetty konfiguriert werden.

Entfernen Sie das Kommentar-Zeichen bei den Einträgen zur seriellen Konsole in /etc/inittab:

`root #` `nano -w /etc/inittab`

```
# SERIAL CONSOLES
s0:12345:respawn:/sbin/agetty 9600 ttyS0 vt100
s1:12345:respawn:/sbin/agetty 9600 ttyS1 vt100

```

### systemd

Um den SSH-Server zu aktivieren, führen Sie aus:

`root #` `systemctl enable sshd`

`root #` `systemctl enable sshd`

Um die Unterstützung der seriellen Konsole zu aktivieren, führen Sie aus:

`root #` `systemctl enable getty@tty1.service`

`root #` `systemctl enable getty@tty1.service`

## Optional: Shell-Vervollständigung

### Bash

Bash ist die Standard-Shell für Gentoo-Systeme, und daher kann die Installation von Vervollständigungserweiterungen die Effizienz und den Komfort bei der Verwaltung des Systems erhöhen. Das Paket [app-shells/bash-completion](https://packages.gentoo.org/packages/app-shells/bash-completion) installiert Vervollständigungen für Gentoo-spezifische Befehle, sowie für viele andere gängige Befehle und Dienstprogramme:

`root #` `emerge --ask app-shells/bash-completion`

Nach der Installation kann die Bash-Vervollständigung für bestimmte Befehle über eselect verwaltet werden. Siehe den [Shell-Vervollständigungsintegrationsabschnitt](/wiki/Bash#Shell_completion_integrations "Bash") des Bash-Artikels für weitere Details.

## Zeitsynchronisation

**Wichtig**

Systems without a functioning [Real-Time Clock (RTC)](/wiki/System_time/de#Software_clock_vs_Hardware_clock "System time/de") must set the [system time](/wiki/System_time/de "System time/de") at every system start, and on regular intervals thereafter.

Es ist wichtig, die Systemuhr mit der aktuellen Zeit zu synchronisieren. Normalerweise wird dafür das [NTP](/wiki/NTP "NTP") Protokoll und [NTP](/wiki/NTP "NTP") Software verwendet. Es gibt andere Implementierungen des NTP Protokolls, beispielsweise [Chrony](/wiki/Chrony "Chrony").

The clock is usually synchronized via the [Network Time Protocol](/wiki/Network_Time_Protocol "Network Time Protocol"), with an implementation such as [chrony](/wiki/Chrony "Chrony").

### chrony

Um beispielsweise Chrony zu installieren:

`root #` `emerge --ask net-misc/chrony`

`root #` `emerge --ask net-misc/chrony`

See the [chrony](/wiki/Chrony "Chrony") article for further information, for example if more advanced configurations are required.

### OpenRC

Unter OpenRC, starten Sie:

`root #` `rc-update add chronyd default`

`root #` `rc-update add chronyd default`

### systemd

Unter systemd, starten Sie:

`root #` `systemctl enable chronyd.service`

`root #` `systemctl enable chronyd.service`

### systemd-timesyncd

Alternativ dazu können systemd-Benutzer auch den einfacheren systemd-timesyncd SNTP-Client verwenden, der standardmäßig installiert ist.

`root #` `systemctl enable systemd-timesyncd.service`

`root #` `systemctl enable systemd-timesyncd.service`

## Dateisystemwerkzeuge

Abhängig von den verwendeten Dateisystemen kann es notwendig sein, die erforderlichen Dateisystem-Dienstprogramme zu installieren (zur Überprüfung der Dateisystem-Integrität, zur (Neu-)Formatierung von Dateisystemen usw.). Beachten Sie, dass die ext4 Userspace-Werkzeuge [sys-fs/e2fsprogs](https://packages.gentoo.org/packages/sys-fs/e2fsprogs) bereits als Teil des [@system Sets](/wiki/System_set_(Portage) "System set (Portage)") installiert sind.

In der folgenden Tabelle sind die zu installierenden Werkzeuge aufgeführt, wenn bestimmte Dateisystem-Werkzeuge in der installierten Umgebung benötigt werden:

Dateisystem
Paket
XFS
[sys-fs/xfsprogs](https://packages.gentoo.org/packages/sys-fs/xfsprogs)ext4
[sys-fs/e2fsprogs](https://packages.gentoo.org/packages/sys-fs/e2fsprogs)VFAT (FAT32, ...)
[sys-fs/dosfstools](https://packages.gentoo.org/packages/sys-fs/dosfstools)Btrfs
[sys-fs/btrfs-progs](https://packages.gentoo.org/packages/sys-fs/btrfs-progs)ZFS
[sys-fs/zfs](https://packages.gentoo.org/packages/sys-fs/zfs)JFS
[sys-fs/jfsutils](https://packages.gentoo.org/packages/sys-fs/jfsutils)ReiserFS
[sys-fs/reiserfsprogs](https://packages.gentoo.org/packages/sys-fs/reiserfsprogs)

Es wird empfohlen, dass [sys-block/io-scheduler-udev-rules](https://packages.gentoo.org/packages/sys-block/io-scheduler-udev-rules) für das korrekte Scheduler-Verhalten mit z.B. nvme-Geräten installiert wird:

`root #` `emerge --ask sys-block/io-scheduler-udev-rules`

`root #` `emerge --ask sys-block/io-scheduler-udev-rules`

**Tipp**

Weitere Informationen zu Dateisystemen in Gentoo finden Sie im [Artikel zu Dateisystemen](/wiki/Filesystem/de "Filesystem/de").

Als nächstes folgt das Kapitel [Konfigurieren des Bootloaders](/wiki/Handbook:HPPA/Installation/Bootloader/de "Handbook:HPPA/Installation/Bootloader/de").

[← Konfiguration des Systems](/wiki/Handbook:HPPA/Installation/System/de "Previous part") [Anfang](/wiki/Handbook:HPPA/de "Handbook:HPPA/de") [Konfiguration des Bootloaders →](/wiki/Handbook:HPPA/Installation/Bootloader/de "Next part")