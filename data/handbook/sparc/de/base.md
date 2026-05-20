# Base

Other languages:

- Deutsch
- [English](/wiki/Handbook:SPARC/Installation/Base "Handbook:SPARC/Installation/Base (100% translated)")
- [español](/wiki/Handbook:SPARC/Installation/Base/es "Manual de Gentoo: SPARC/Instalación/Base (100% translated)")
- [français](/wiki/Handbook:SPARC/Installation/Base/fr "Handbook:SPARC/Installation/Base/fr (100% translated)")
- [italiano](/wiki/Handbook:SPARC/Installation/Base/it "Handbook:SPARC/Installation/Base/it (50% translated)")
- [magyar](/wiki/Handbook:SPARC/Installation/Base/hu "Handbook:SPARC/Installation/Base/hu (100% translated)")
- [polski](/wiki/Handbook:SPARC/Installation/Base/pl "Handbook:SPARC/Installation/Base (100% translated)")
- [português do Brasil](/wiki/Handbook:SPARC/Installation/Base/pt-br "Handbook:SPARC/Installation/Base/pt-br (50% translated)")
- [русский](/wiki/Handbook:SPARC/Installation/Base/ru "Handbook:SPARC/Installation/Base (100% translated)")
- [தமிழ்](/wiki/Handbook:SPARC/Installation/Base/ta "கையேடு:SPARC/நிறுவல்/அடிப்படை (100% translated)")
- [中文（中国大陆）‎](/wiki/Handbook:SPARC/Installation/Base/zh-cn "手册：SPARC/安装/安装基本系统 (100% translated)")
- [日本語](/wiki/Handbook:SPARC/Installation/Base/ja "ハンドブック:SPARC/インストール/ベース (100% translated)")
- [한국어](/wiki/Handbook:SPARC/Installation/Base/ko "Handbook:SPARC/Installation/Base/ko (100% translated)")

[SPARC Handbuch](/wiki/Handbook:SPARC/de "Handbook:SPARC/de")[Installation](/wiki/Handbook:SPARC/Full/Installation/de "Handbook:SPARC/Full/Installation/de")[Über die Installation](/wiki/Handbook:SPARC/Installation/About/de "Handbook:SPARC/Installation/About/de")[Auswahl des Mediums](/wiki/Handbook:SPARC/Installation/Media/de "Handbook:SPARC/Installation/Media/de")[Konfiguration des Netzwerks](/wiki/Handbook:SPARC/Installation/Networking/de "Handbook:SPARC/Installation/Networking/de")[Vorbereiten der Festplatte(n)](/wiki/Handbook:SPARC/Installation/Disks/de "Handbook:SPARC/Installation/Disks/de")[Installation des Stage Archivs](/wiki/Handbook:SPARC/Installation/Stage/de "Handbook:SPARC/Installation/Stage/de")Installation des Basissystems[Konfiguration des Kernels](/wiki/Handbook:SPARC/Installation/Kernel/de "Handbook:SPARC/Installation/Kernel/de")[Konfiguration des Systems](/wiki/Handbook:SPARC/Installation/System/de "Handbook:SPARC/Installation/System/de")[Installation der Tools](/wiki/Handbook:SPARC/Installation/Tools/de "Handbook:SPARC/Installation/Tools/de")[Konfiguration des Bootloaders](/wiki/Handbook:SPARC/Installation/Bootloader/de "Handbook:SPARC/Installation/Bootloader/de")[Abschluss](/wiki/Handbook:SPARC/Installation/Finalizing/de "Handbook:SPARC/Installation/Finalizing/de")[Arbeiten mit Gentoo](/wiki/Handbook:SPARC/Full/Working/de "Handbook:SPARC/Full/Working/de")[Portage-Einführung](/wiki/Handbook:SPARC/Working/Portage/de "Handbook:SPARC/Working/Portage/de")[USE-Flags](/wiki/Handbook:SPARC/Working/USE/de "Handbook:SPARC/Working/USE/de")[Portage-Features](/wiki/Handbook:SPARC/Working/Features/de "Handbook:SPARC/Working/Features/de")[Initskript-System](/wiki/Handbook:SPARC/Working/Initscripts/de "Handbook:SPARC/Working/Initscripts/de")[Umgebungsvariablen](/wiki/Handbook:SPARC/Working/EnvVar/de "Handbook:SPARC/Working/EnvVar/de")[Arbeiten mit Portage](/wiki/Handbook:SPARC/Full/Portage/de "Handbook:SPARC/Full/Portage/de")[Dateien und Verzeichnisse](/wiki/Handbook:SPARC/Portage/Files/de "Handbook:SPARC/Portage/Files/de")[Variablen](/wiki/Handbook:SPARC/Portage/Variables/de "Handbook:SPARC/Portage/Variables/de")[Mischen von Softwarezweigen](/wiki/Handbook:SPARC/Portage/Branches/de "Handbook:SPARC/Portage/Branches/de")[Zusätzliche Tools](/wiki/Handbook:SPARC/Portage/Tools/de "Handbook:SPARC/Portage/Tools/de")[Eigener Portage-Tree](/wiki/Handbook:SPARC/Portage/CustomTree/de "Handbook:SPARC/Portage/CustomTree/de")[Erweiterte Portage-Features](/wiki/Handbook:SPARC/Portage/Advanced/de "Handbook:SPARC/Portage/Advanced/de")[Netzwerk-Konfiguration](/wiki/Handbook:SPARC/Full/Networking/de "Handbook:SPARC/Full/Networking/de")[Zu Beginn](/wiki/Handbook:SPARC/Networking/Introduction/de "Handbook:SPARC/Networking/Introduction/de")[Fortgeschrittene Konfiguration](/wiki/Handbook:SPARC/Networking/Advanced/de "Handbook:SPARC/Networking/Advanced/de")[Modulare Vernetzung](/wiki/Handbook:SPARC/Networking/Modular/de "Handbook:SPARC/Networking/Modular/de")[Drahtlose Netzwerke](/wiki/Handbook:SPARC/Networking/Wireless/de "Handbook:SPARC/Networking/Wireless/de")[Funktionalität hinzufügen](/wiki/Handbook:SPARC/Networking/Extending/de "Handbook:SPARC/Networking/Extending/de")[Dynamisches Management](/wiki/Handbook:SPARC/Networking/Dynamic/de "Handbook:SPARC/Networking/Dynamic/de")

## Contents

- [1Chrooten](#Chrooten)
  - [1.1DNS-Info kopieren](#DNS-Info_kopieren)
  - [1.2Notwendige Dateisysteme einhängen](#Notwendige_Dateisysteme_einh.C3.A4ngen)
  - [1.3Betreten der neuen Umgebung](#Betreten_der_neuen_Umgebung)
  - [1.4Preparing for a bootloader](#Preparing_for_a_bootloader)
    - [1.4.1DOS/Legacy BIOS systems](#DOS.2FLegacy_BIOS_systems)
- [2Portage konfigurieren](#Portage_konfigurieren)
  - [2.1Ein Gentoo-Ebuild-Repositorium Snapshot aus dem Web installieren](#Ein_Gentoo-Ebuild-Repositorium_Snapshot_aus_dem_Web_installieren)
  - [2.2Optional: Spiegelserver wählen](#Optional:_Spiegelserver_w.C3.A4hlen)
    - [2.2.1Optional: rsync mirrors](#Optional:_rsync_mirrors)
  - [2.3Optional: Gentoo-Ebuild-Repositorium aktualisieren](#Optional:_Gentoo-Ebuild-Repositorium_aktualisieren)
  - [2.4News Items lesen](#News_Items_lesen)
  - [2.5Auswahl des richtigen Profils](#Auswahl_des_richtigen_Profils)
  - [2.6Optional: Adding a binary package host](#Optional:_Adding_a_binary_package_host)
    - [2.6.1Repository configuration](#Repository_configuration)
    - [2.6.2Installing binary packages](#Installing_binary_packages)
  - [2.7USE Variable konfigurieren](#USE_Variable_konfigurieren)
    - [2.7.1CPU\_FLAGS\_\*](#CPU_FLAGS_.2A)
    - [2.7.2VIDEO\_CARDS](#VIDEO_CARDS)
  - [2.8Optional: Die ACCEPT\_LICENSE Variable konfigurieren](#Optional:_Die_ACCEPT_LICENSE_Variable_konfigurieren)
  - [2.9@world set updaten](#.40world_set_updaten)
    - [2.9.1Removing obsolete packages](#Removing_obsolete_packages)
- [3Zeitzone](#Zeitzone)
- [4Locales konfigurieren](#Locales_konfigurieren)
  - [4.1Locales erzeugen](#Locales_erzeugen)
  - [4.2Locale auswählen](#Locale_ausw.C3.A4hlen)
- [5References](#References)

## Chrooten

### DNS-Info kopieren

Eine Sache bleibt noch zu tun, bevor Sie die neue Umgebung betreten - und das ist das Kopieren der DNS-Informationen in der Datei /etc/resolv.conf. Dies ist notwendig um sicherzustellen, dass Netz-Verbindungen auch nach dem Betreten der neuen Umgebung noch funktionieren. /etc/resolv.conf enthält u.a. die IP-Adressen der Namensserver.

Zum Kopieren dieser Information ist es empfehlenswert, beim Befehl cp die Option `--dereference` zu verwenden. Wenn /etc/resolv.conf ein symbolischer Link ist stellt dies sicher, dass die Zieldatei anstelle des symbolischen Links selbst kopiert wird. Andernfalls würde der symbolische Link auf eine nicht existierende Datei zeigen (weil das Link-Ziel höchstwahrscheinlich in der neuen Umgebung nicht verfügbar ist).

`root #` `cp --dereference /etc/resolv.conf /mnt/gentoo/etc/`

### Notwendige Dateisysteme einhängen

**Tipp**

If using Gentoo's install media, this step can be replaced with simply: arch-chroot /mnt/gentoo.

In wenigen Augenblicken wird der Linux-Root (/) auf den neuen Ort geändert werden.

Die Dateisysteme, die verfügbar gemacht werden müssen, sind:

- /proc/ ist ein ein Pseudo-Dateisystem. Es sieht aus wie normale Dateien, wird aber vom Linux-Kernel on-the-fly erzeugt
- /sys/ ist ein Pseudo-Dateisystem, genauso wie /proc/. Einst war es dafür gedacht, /proc/ zu ersetzen. Es ist besser strukturiert als dieses.
- /dev/ ist ein gewöhnliches Dateisystem, das alle Gerätedateien enthält. Es wird teilweise vom Linux Device Manager verwaltet (normalerweise udev).
- /run/ ist ein temporäres Dateisystem für Dateien, die im laufenden Betrieb benötigt werden, die aber einen Reboot nicht überleben müssen. Beispiele sind PID-Dateien oder Locks.

/proc/ wird an /mnt/gentoo/proc/ eingehängt, wohingegen die anderen Dateisysteme über Bind-Mounts eingehängt werden. Letzteres bedeutet, dass beispielsweise /mnt/gentoo/sys/ in Wirklichkeit /sys/ _ist_ (es ist lediglich ein zweiter Einstiegspunkt zum selben Dateisystem), wohingegen /mnt/gentoo/proc/ eine neue Einbindung (sozusagen eine neue Instanz) des Dateisystems ist.

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

`root #` `mount --bind /run /mnt/gentoo/run
`

`root #` `mount --make-slave /mnt/gentoo/run
`

**Hinweis**

Die `--make-rslave` Operationen werden für die spätere systemd Unterstützung bei der Installation benötigt.

**Warnung**

Bei der Verwendung von Nicht-Gentoo Installationsmedien ist dies möglicherweise nicht ausreichend. Bei einigen Distributionen ist /dev/shm ein symbolischer Link zu /run/shm/, der nach einem chroot ungültig wird. Dies kann behoben werden, indem Sie /dev/shm/ im Voraus zu einem entsprechenden tmpfs mount machen:

`root #` `test -L /dev/shm && rm /dev/shm && mkdir /dev/shm
`

`root #` `mount --types tmpfs --options nosuid,nodev,noexec shm /dev/shm
`

Stellen Sie zudem sicher, dass Mode 1777 gesetzt ist:

`root #` `chmod 1777 /dev/shm /run/shm`

### Betreten der neuen Umgebung

Nun, da alle Partitionen initialisiert sind und die Basis-Umgebung installiert ist, wird es Zeit, die neue Installationsumgebung durch chroot zu betreten. Das bedeutet, dass die Sitzung ihr Wurzelverzeichnis (/) von der aktuellen Installationsumgebung (Installations-CD oder anderes Installationsmedium) zum Installationssystem (nämlich die initialisierten Partitionen) ändert. Daher der Name _change root_ oder _chroot_.

Dieses Chrooten erfolgt in drei Schritten:

1. Das Wurzelverzeichnis wird mit Hilfe von chroot von / (auf dem Installationsmedium) auf /mnt/gentoo/ (auf den Partitionen) geändert.
2. Einige Einstellungen (jene in /etc/profile) werden über den Befehl source neu in den Speicher geladen.
3. Die primäre Eingabeaufforderung wird geändert, damit wir nicht vergessen, dass diese Sitzung innerhalb einer chroot-Umgebung läuft.

`root #` `chroot /mnt/gentoo /bin/bash
`

`root #` `source /etc/profile
`

`root #` `export PS1="(chroot) ${PS1}"`

Von diesem Punkt an werden alle Aktionen direkt auf der neuen Gentoo Linux Umgebung ausgeführt.

**Tipp**

Wenn die Installation bei einem der ab hier folgenden Schritte unterbrochen werden sollte, _sollte_ es möglich sein, ab dieser hier Stelle weiterzuarbeiten. Es ist nicht nötig, die Partitionen erneut zu erstellen! [Mounten Sie die Root-Partition](/wiki/Handbook:SPARC/Installation/Disks/de#Mounting_the_root_partition "Handbook:SPARC/Installation/Disks/de") und führen Sie die oben beschriebenen Schritte ab [DNS-Info kopieren](/wiki/Handbook:SPARC/Installation/Base/de#Copy_DNS_info "Handbook:SPARC/Installation/Base/de") erneut aus, um wieder in die neue Gentoo Linux Umgebung zu gelangen. Dieses Vorgehen ist ebenfalls sinnvoll, um Bootloader-Probleme zu beheben. Weitere Informationen erhalten Sie im [chroot](/wiki/Chroot/de "Chroot/de") Artikel.

### Preparing for a bootloader

Now that the new environment has been entered, it is necessary to prepare the new environment for the bootloader. It will be important to have the correct partition mounted when it is time to install the bootloader.

#### DOS/Legacy BIOS systems

For DOS/Legacy BIOS systems, the bootloader will be installed into the /boot directory, therefore mount as follows:

`root #` `mount  /boot`

## Portage konfigurieren

### Ein Gentoo-Ebuild-Repositorium Snapshot aus dem Web installieren

Der nächste Schritt besteht darin, einen Snapshot des Gentoo ebuild Repositorys zu installieren. Dieser Snapshot enthält eine Sammlung von Dateien, die Portage informiert über verfügbare Software-Titel (für die Installation), welche Profile der Administrator auswählen kann, Paket- oder Profil-spezifische News-Items, usw.

Die Verwendung von emerge-webrsync wird empfohlen für diejenigen, die hinter einer restriktiven Firewall sitzen (das Programm lädt den Snapshot über die Protokolle HTTP/FTP herunter) und für diejenigen, die Netzwerk-Bandbreite sparen wollen. Leser, die keine Einschränkungen durch Firewalls oder von der Netzwerk-Bandbreite haben, können zum nächsten Abschnitt springen.

Der folgende Befehl holt den neuesten Portage-Snapshot (den Gentoo tagesaktuell veröffentlicht) von einem der Gentoo-Spiegel und installiert ihn auf dem System.

`root #` `emerge-webrsync`

**Hinweis**

Während dieser Operation könnte sich emerge-webrsync über das Fehlen von /var/db/repos/gentoo/ beschweren. Dies ist zu erwarten und kein Grund zur Sorge - das Tool wird das Verzeichnis anlegen.

Von diesem Punkt an könnte Portage erwähnen, dass bestimmte Updates empfehlenswert sind. Dies ist deshalb so, weil es möglicherweise neuere Versionen von Paketen gibt, die durch das Stage Tar-Archiv installiert wurden. Nach der Installation des Repository Snapshots weiß Portage nun von diesen neueren Versionen. Paket-Updates können im Augenblick bedenkenlos ignoriert werden. Die Updates können verzögert werden, bis die Gentoo Installation abgeschlossen ist.

### Optional: Spiegelserver wählen

Um den Quellcode zügig herunterzuladen, wird empfohlen, einen schnellen Spiegel auszuwählen. Portage schaut in der Datei make.conf nach der Variable GENTOO\_MIRRORS und verwendet darin aufgelistete Spiegel. Es ist möglich, zur Gentoo Mirror-Liste zu surfen und nach einem Spiegel (oder mehreren Spiegeln) zu suchen, die nahe dem Systemstandort liegen (da diese meistens die schnellsten sind). Allerdings bieten wir ein nettes Tool namens mirrorselect, das den Benutzern ein schönes Interface zur Auswahl der benötigten Spiegel bietet. Gehen Sie einfach zu den Spiegeln der Wahl und drücken Sie die `Leertaste` um einen oder mehrere Spiegel auszuwählen.

A tool called mirrorselect provides a pretty text interface to more quickly query and select suitable mirrors. Just navigate to the mirrors of choice and press `Spacebar` to select one or more mirrors.

`root #` `emerge --ask --verbose --oneshot app-portage/mirrorselect`

`root #` `mirrorselect -i -o >> /etc/portage/make.conf`

Alternatively, a list of active mirrors are [available online](https://www.gentoo.org/downloads/mirrors/).

#### Optional: rsync mirrors

Gentoo also has many rsync mirrors which can speed up downloading the available package list using `emerge --sync` (explained in more detail later) by selecting a mirror closer that is geographically closer to the user.

`root #` `mkdir /etc/portage/repos.conf
`

`root #` `cp /usr/share/portage/config/repos.conf /etc/portage/repos.conf/gentoo.conf
`

Select a mirror close to the system's location from [https://www.gentoo.org/support/rsync-mirrors/](https://www.gentoo.org/support/rsync-mirrors/) and replace the sync-uri default mirror in /etc/portage/repos.conf/gentoo.conf with the desired mirror location.

DATEI **`/etc/portage/repos.conf/gentoo.conf`** **UK-based rsync mirror example**

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

### Optional: Gentoo-Ebuild-Repositorium aktualisieren

Es ist möglich, das Gentoo ebuild Repository mit emerge auf die neueste Version zu aktualisieren. Wenn Sie mit dem vorhergehenden Befehl emerge-webrsync einen aktuellen Snapshot installiert haben (Snapshots sind in der Regel nicht älter als 24 Stunden), ist dieser Schritt optional.

Angenommen Sie benötigen die neuesten Paket-Updates (bis zu 1 Stunde), dann benutzen Sie emerge --sync. Dieser Befehl nutzt das rsync Protokoll zur Aktualisierung des Gentoo ebuild Repository (welcher zuvor durch emerge-webrsync bezogen wurde) auf den aktuellsten Stand.

`root #` `emerge --sync`

Auf langsamen Terminals, wie einigen Framebuffer- oder seriellen Konsolen, ist es empfehlenswert, die Option `--quiet` zu nutzen, um den Vorgang zu beschleunigen:

`root #` `emerge --sync --quiet`

### News Items lesen

Wenn das Gentoo ebuild Repository auf das System synchronisiert wird, könnte Portage Hinweistexte wie im folgenden Beispiel ausgeben:

` * IMPORTANT: 2 news items need reading for repository 'gentoo'.
`

` * Use eselect news to read news items.
`

News Items wurden als Kommunikationsmedium geschaffen, um den Benutzern wichtige Mitteilungen über das Gentoo ebuild Repository zukommen lassen zu können. Zur Verwaltung der Mitteilungen verwenden Sie eselect news. Die Anwendung eselect ist ein Gentoo-spezifisches Programm, das eine gemeinsame Verwaltungsschnittstelle für verschiedene System-Administrations-Aufgaben bietet. In diesem Fall wird eselect aufgefordert, das Modul `news` zu verwenden.

Im Modul `news` werden drei Operationen am meisten genutzt:

- Mit `list` wird eine Übersicht der verfügbaren News-Einträge angezeigt.
- Mit `read` können die News-Einträge gelesen werden.
- Mit `purge` lassen sich News-Einträge löschen, sobald sie gelesen wurden. Ein erneutes Einlesen erfolgt nicht.

`root #` `eselect news list
`

`root #` `eselect news read`

Mehr Informationen zum News Reader sind über seine Manpage verfügbar:

`root #` `man news.eselect`

### Auswahl des richtigen Profils

**Tipp**

Desktop-Profile sind nicht ausschließlich für _Desktop-Umgebungen_ gedacht. Sie sind auch für minimale Fenstermanager wie i3 oder sway geeignet.

Ein Profil ( _profile_) ist wichtiger Baustein für jedes Gentoo System. Es definiert nicht nur Standardwerte für USE, CFLAGS und andere wichtige Variablen, sondern legt das System auch auf einen bestimmten Bereich von Paketversionen fest. Diese Einstellungen werden von den Gentoo-Entwicklern gepflegt.

Um zu sehen, welches Profil das System momentan verwendet, können Sie eselect mit dem `profile` Modul ausführen:

**Tipp**

On an install media without a scroll-able terminal, `eselect profile list | less` can provide an easy way to list all available profiles while providing the ability to scroll.

`root #` `eselect profile list`

```
Available profile symlink targets:
  [1]   default/linux/sparc/23.0 *
  [2]   default/linux/sparc/23.0/desktop
  [3]   default/linux/sparc/23.0/desktop/gnome
  [4]   default/linux/sparc/23.0/desktop/kde

```

**Hinweis**

Die Ausgabe des Befehls ist nur ein Beispiel und kann sich im Laufe der Zeit ändern.

**Hinweis**

Um **systemd** zu verwenden, wählen Sie ein Profil, dass "systemd" im Namen hat und umgekehrt, falls nicht

Für einige Architekturen gibt es auch Desktop-Unterprofile.

**Warnung**

Profil-Upgrades sind nicht einfach. Wenn Sie das initiale Profil auswählen, verwenden Sie das, dass die **selbe Versionsnummer** hat, wie das Profil, das vom Stage3 Tar-Archiv installiert wurde (beispielsweise 23.0). Neue Profil-Versionen werden über News Items angekündigt, die detaillierte Migrationsanleitungen enthalten. Folgen Sie diesen Migrationsanleitungen bevor Sie auf ein neues Profil wechseln.

Nach dem Betrachten der verfügbaren Profile für die sparc Architektur kann der Benutzer ein anderes Profil für das System wählen:

`root #` `eselect profile set 2`

**Hinweis**

Das `developer` Unterprofil ist eigens für die Gentoo Linux Entwicklung und nicht für die Nutzung durch gewöhnliche Benutzer gedacht.

### Optional: Adding a binary package host

Since December 2023, Gentoo's [Release Engineering team](/wiki/Project:RelEng "Project:RelEng") has offered an [official binary package host](/wiki/Binary_package_quickstart "Binary package quickstart") (colloquially shorted to just "binhost") for use by the general community to retrieve and install binary packages (binpkgs).[\[1\]](#cite_note-1)

Adding a binary package host allows Portage to install cryptographically signed, compiled packages. In many cases, adding a binary package host will _greatly_ decrease the mean time to package installation and adds much benefit when running Gentoo on older, slower, or low power systems.

#### Repository configuration

The repository configuration for a binhost is found in Portage's /etc/portage/binrepos.conf/ directory, which functions similarly to the configuration mentioned in the [Gentoo ebuild repository](/wiki/Handbook:SPARC/Installation/Base/de#Gentoo_ebuild_repository "Handbook:SPARC/Installation/Base/de") section.

When defining a binary host, there are two important aspects to consider:

1. The architecture and profile targets within the `sync-uri` value _do_ matter and should align to the respective computer architecture ( **sparc** in this case) and system profile selected in the [Choosing the right profile](/wiki/Handbook:SPARC/Installation/Base/de#Choosing_the_right_profile "Handbook:SPARC/Installation/Base/de") section.
2. Selecting a fast, geographically close mirror will generally shorten retrieval time. Review the mirrorselect tool mentioned in the [Optional: Selecting mirrors](/wiki/Handbook:SPARC/Installation/Base/de#Gentoo_ebuild_repository "Handbook:SPARC/Installation/Base/de") section or review the [online list of mirrors](https://www.gentoo.org/downloads/mirrors/) where URL values can be discovered.


DATEI **`/etc/portage/binrepos.conf/gentoo.conf`** **CDN-based binary package host example**

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

#### Installing binary packages

Portage will compile packages from code source by default. It can be instructed to use binary packages in the following ways:

1. The `--getbinpkg` option can be passed when invoking the emerge command. This method of binary package installation is useful to install only a particular binary package.
2. Changing the system's default via Portage's FEATURES variable, which is exposed through the /etc/portage/make.conf file. Applying this configuration change will cause Portage to query the binary package host for the package(s) to be requested and fall back to compiling locally when no results are found.

For example, to have Portage always install available binary packages:

DATEI **`/etc/portage/make.conf`** **Configure Portage to use binary packages by default**

```
# Appending getbinpkg to the list of values within the FEATURES variable
FEATURES="${FEATURES} getbinpkg"
# Require signatures
FEATURES="${FEATURES} binpkg-request-signature"

```

Please also run getuto for Portage to set up the necessary keyring for verification:

`root #` `getuto`

Additional Portage features will be discussed in the [the next chapter](/wiki/Handbook:SPARC/Working/Features/de#Portage_features "Handbook:SPARC/Working/Features/de") of the handbook.

### USE Variable konfigurieren

USE ist eine der mächtigsten Variablen, die Gentoo seinen Benutzern bietet. Viele Programme können mit oder ohne optionale Unterstützung für bestimmte Dinge kompiliert werden. Beispielsweise können einige Programme mit GTK- oder Qt-Unterstützung kompiliert werden. Andere können mit oder ohne SSL-Unterstützung kompiliert werden. Einige Programme können sogar mit Framebuffer-Unterstützung (svgalib) anstelle von X11-Unterstützung (X-Server) kompiliert werden.

Die meisten Distributionen kompilieren ihre Pakete mit Unterstützung für möglichst viel. Dies erhöht die Größe der Programme und verlängert die Programmstartzeit, nicht zu erwähnen die enorme Menge von Abhängigkeiten. Mit Gentoo können die Benutzer definieren, mit welchen Optionen ein Paket kompiliert werden soll. Hier kommt USE ins Spiel.

In der Variablen USE definieren die Benutzer Schlüsselwörter, die auf Optionen beim Kompilieren abgebildet werden. Beispielsweise kompiliert `ssl` SSL Unterstützung in die Programme, die das unterstützen. `-X` entfernt X-Server-Unterstützung (beachten Sie das Minuszeichen am Anfang). `gnome gtk -kde -qt5` kompiliert Programme mit GNOME und GTK+ Unterstützung, aber nicht mit KDE und Qt5 Unterstützung. Das führt zu einem System, das komplett für GNOME optimiert ist (vorausgesetzt die Architektur unterstützt es).

Die Standard-USE-Einstellungen befinden sich in den make.defaults Dateien des Gentoo-Profils, das das System verwendet. Gentoo benutzt ein (komplexes) Vererbungssystem für seine Profile, in das wir in dieser Phase nicht eintauchen wollen. Der einfachste Weg, die momentan aktiven USE Einstellungen zu überprüfen, ist emerge --info auszuführen und die Zeile auszuwählen, die mit USE beginnt:

`root #` `emerge --info | grep ^USE`

```
USE="X acl alsa amd64 berkdb bindist bzip2 cli cracklib crypt cxx dri ..."

```

**Hinweis**

Das obige Beispiel ist verkürzt, die tatsächliche Liste der USE-Werte ist viel viel länger.

Eine vollständige Beschreibung der verfügbaren USE-Flags finden Sie auf dem System in der Datei /var/db/repos/gentoo/profiles/use.desc.

`root #` `less /var/db/repos/gentoo/profiles/use.desc`

Innerhalb des Befehls less können Sie mit Hilfe der Tasten `↑` und `↓` scrollen. Zum Beenden drücken Sie `q`.

Als Beispiel zeigen wir die USE Einstellung für ein KDE-basiertes System mit DVD, ALSA und CD-Aufnahme Unterstützung:

`root #` `nano /etc/portage/make.conf`

DATEI **`/etc/portage/make.conf`** **USE-Flags für ein KDE/Plasma-basiertes System mit Unterstützung für DVD, ALSA und CD-Aufnahme**

```
USE="-gtk -gnome qt5 kde dvd alsa cdr"

```

Wenn ein USE Wert in /etc/portage/make.conf definiert wird, wird er zu der USE-Flag Liste des Systems _hinzugefügt_. USE-Flags auf dieser Liste können _entfernt_ werden, indem ein `-` Minuszeichen vor den Wert gesetzt wird. Um beispielsweise die Unterstützung für X11 zu deaktivieren, kann `-X` definiert werden:

DATEI **`/etc/portage/make.conf`** **Standard USE-Flags ignorieren**

```
USE="-X acl alsa"

```

**Warnung**

Obwohl es möglich ist, mit `-*` alle USE-Flags zu deaktivieren (mit Ausnahme der in make.conf definierten USE-Flags), raten wir _mit Nachdruck_ davon ab. Die Entwickler von Ebuilds wählen bestimmte Default-USE-Flags in ebuilds um Konflikte zu vermeiden, die Sicherheit zu gewährleisten, Fehler zu vermeiden und aus vielen weiteren Gründen. Die Deaktivierung _aller_ USE-Flags beeinträchtigt das Standard-Verhalten und kann zu schwerwiegenden Problemen führen

#### CPU\_FLAGS\_\*

Einige Architekturen (einschließlich AMD64/X86, ARM, PPC) haben eine [USE\_EXPAND](/wiki/USE_EXPAND "USE EXPAND")-Variable namens [CPU\_FLAGS\_<ARCH>](/wiki/CPU_FLAGS_*/de "CPU FLAGS */de"), wobei <ARCH> durch den Namen der jeweiligen Systemarchitektur ersetzt wird.

**Wichtig**

Nicht verwirren lassen! **AMD64**\- und **X86**-Systeme haben eine gemeinsame Architektur, daher ist der richtige Variablenname für **AMD64**-Systeme CPU\_FLAGS\_X86.

Dies wird verwendet um den Build so zu konfigurieren, dass er in einem bestimmten Assemblercode oder anderen Intrinsics kompiliert, normalerweise handgeschrieben oder anderweitig extra, und ist _nicht_ dasselbe wie die Aufforderung an den Compiler, optimierten Code für eine bestimmte CPU-Funktion auszugeben (z.B. `-march=`).

Benutzer sollten diese Variable zusätzlich zur Konfiguration ihrer COMMON\_FLAGS wie gewünscht setzen.

Zur Konfiguration sind folgende Schritte erforderlich:

`root #` `emerge --ask app-portage/cpuid2cpuflags`

Wenn Sie neugierig sind, können Sie das Programm starten und sich die Ausgabe ansehen:

`root #` `cpuid2cpuflags`

Kopieren Sie die Ausgabe des Programms nach package.use:

`root #` `echo "*/* $(cpuid2cpuflags)" > /etc/portage/package.use/00cpu-flags`

#### VIDEO\_CARDS

By default a profile already sets a few video cards. For many reasons this is not ideal, but a useful safety net.

To configure the system correctly the user needs to first unset the preset video cards with `VIDEO_CARDS: -*` then set the correct card for that system.

DATEI **`/etc/portage/package.use/00video_cards`** **Example**

```
*/* VIDEO_CARDS: -* <GPU DRIVER NAME>

```

Below is a table that can be used to easily compare the different video card types to their respective `VIDEO_CARDS` value.

GPU
VIDEO\_CARDS
(Official) Nvidia Maxwell and newer`nvidia`Nouveau Nvidia [Supported list](/wiki/NVIDIA "NVIDIA")`nouveau`AMD since Sea Islands`amdgpu radeonsi`ATI and older AMDSee [radeon#Feature support](/wiki/Radeon#Feature_support "Radeon")Intel Nehalem and newer`intel`Intel Gen 2 and 3 [Supported list](/wiki/Intel#.23Feature_support.2Fde "Intel")`intel i915`QEMU/KVM`virgl`WSL`d3d12`

Below is a few examples of a properly set package.use for _VIDEO\_CARDS_:

DATEI **`/etc/portage/package.use/00video_cards`** **AMD example**

```
*/* VIDEO_CARDS: -* amdgpu radeonsi

```

DATEI **`/etc/portage/package.use/00video_cards`** **Intel example**

```
*/* VIDEO_CARDS: -* intel

```

DATEI **`/etc/portage/package.use/00video_cards`** **Nvidia example**

```
*/* VIDEO_CARDS: -* nvidia

```

Details for various GPU(s) can be found at the [AMDGPU](/wiki/AMDGPU/de "AMDGPU/de"), [Intel](/wiki/Intel "Intel"), [Nouveau (Open Source)](/wiki/Nouveau "Nouveau"), or [NVIDIA (Proprietary)](/wiki/NVIDIA "NVIDIA") articles.

### Optional: Die ACCEPT\_LICENSE Variable konfigurieren

Starting with Gentoo Linux Enhancement Proposal 23 (GLEP 23), a mechanism was created to allow system administrators the ability to "regulate the software they install with regards to licenses... Some want a system free of any software that is not OSI-approved; others are simply curious as to what licenses they are implicitly accepting."[\[2\]](#cite_note-2) With a motivation to have more granular control over the type of software running on a Gentoo system, the ACCEPT\_LICENSE variable was born.

Portage überprüft anhand von ACCEPT\_LICENSE, welche Pakete installiert werden dürfen.
Den derzeit systemweit gültigen Wert können Sie anzeigen mit:

`user $` `portageq envvar ACCEPT_LICENSE`

```
@FREE
```

Die folgende Tabelle zeigt die im Gentoo Repository definierten Lizenz-Gruppen. Die Lizenz-Gruppen werden vom [Gentoo Licenses Projekt](/wiki/Project:Licenses "Project:Licenses") verwaltet.

Gruppen-NameBeschreibung
@GPL-COMPATIBLEGPL compatible licenses approved by the Free Software Foundation [\[a\_license 1\]](#cite_note-3)@FSF-APPROVEDFree software licenses approved by the FSF (includes @GPL-COMPATIBLE)
@OSI-APPROVEDLicenses approved by the Open Source Initiative [\[a\_license 2\]](#cite_note-4)@MISC-FREEMisc licenses that are probably free software, i.e. follow the Free Software Definition [\[a\_license 3\]](#cite_note-5) but are not approved by either FSF or OSI
@FREE-SOFTWARECombines @FSF-APPROVED, @OSI-APPROVED and @MISC-FREE
@FSF-APPROVED-OTHERFSF-approved licenses for "free documentation" and "works of practical use besides software and documentation" (including fonts)
@MISC-FREE-DOCSMisc licenses for free documents and other works (including fonts) that follow the free definition [\[a\_license 4\]](#cite_note-6) but are NOT listed in @FSF-APPROVED-OTHER
@FREE-DOCUMENTSCombines @FSF-APPROVED-OTHER and @MISC-FREE-DOCS
@FREEMetaset of all licenses with the freedom to use, share, modify and share modifications. Combines @FREE-SOFTWARE and @FREE-DOCUMENTS
@BINARY-REDISTRIBUTABLELicenses that at least permit free redistribution of the software in binary form. Includes @FREE
@EULALicense agreements that try to take away your rights. These are more restrictive than "all-rights-reserved" or require explicit approval

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

- Systemweit im ausgewählten Profil.
- Systemweit in der Datei /etc/portage/make.conf.
- Pro Paket in der Datei /etc/portage/package.license.
- Pro Paket in einem /etc/portage/package.license/ _Verzeichnis_ von Dateien.

Optional kann die in den Profilen gesetzte systemweite Voreinstellung überschrieben werden in der Datei /etc/portage/make.conf:

DATEI **`/etc/portage/make.conf`** **Beispiel, wie erlaubte Lizenzen systemweit über ACCEPT\_LICENSE definiert werden können**

```
ACCEPT_LICENSE="-* @FREE @BINARY-REDISTRIBUTABLE"

```

Optional kann man auch akzeptierte Lizenzen pro Paket definieren, wie im folgenden Beispiel für ein Dateiverzeichnis gezeigt. Beachten Sie, dass das Verzeichnis package.license erstellt werden muss, wenn es noch nicht existiert:

`root #` `mkdir /etc/portage/package.license`

Software license details for an individual Gentoo package are stored within the LICENSE variable of the associated ebuild. One package may have one or many software licenses, therefore it be necessary to specify multiple acceptable licenses for a single package.

DATEI **`/etc/portage/package.license/kernel`** **Annahme von Lizenzen pro Paket**

```
app-arch/unrar unRAR
sys-kernel/linux-firmware linux-fw-redistributable
sys-firmware/intel-microcode intel-ucode

```

**Wichtig**

Die LICENSE Variable in einem Ebuild ist nur eine Richtlinie für Gentoo Entwickler und Benutzer. Sie ist keine rechtliche Aussage und es gibt keine Garantie, dass sie die Realität widerspiegelt. Verlassen Sie sich also nicht darauf, sondern überprüfen Sie das Paket selbst eingehend, einschließlich aller Dateien, die auf dem System installiert wurden.

### @world set updaten

**Tipp**

Wenn Sie ein Profil für eine vollständige Desktop-Umgebung (wie KDE oder GNOME) gewählt haben, kann der obige emerge-Befehl recht lange dauern. Wer unter Zeitdruck steht, kann folgende 'Daumenregel' verwenden: je kürzer der Profil-Name, desto weniger umfangreich ist das [@world set](/wiki/World_set_(Portage) "World set (Portage)") des Systems, desto weniger Pakete müssen installiert werden. Mit anderen Worten:

- die Wahl von `default/linux/amd64/23.0` führt zu wenigen Paketen, die upgedatet werden müssen, während
- bei der Wahl von `default/linux/amd64/23.0/desktop/gnome/systemd` viele Pakete installiert werden müssen, weil ein Wechsel von OpenRC zu systemd stattfindet und die GNOME Desktop-Umgebung installiert werden muss.

Der folgende Schritt ist _erforderlich_, damit Ihre Änderungen am Profil wirksam werden. Weiterhin werden alle Pakete aktualisiert, für die es nach dem Erstellungszeitpunkt des von Ihnen installierten Stage Tar-Archivs Updates oder Änderungen an den USE Flags gibt.

1. A profile target _different_ from the stage file has been selected.
2. Additional USE flags have been set for installed packages.

Readers who are performing an 'install Gentoo speed run' may safely skip @world set updates until _after_ their system has rebooted into the new Gentoo environment.

Readers who are performing a slow run can have Portage perform updates for package, profile, and/or USE flag changes at the present time:

`root #` `emerge --ask --verbose --update --deep --changed-use @world`

Readers who added a binary host [above](/wiki/Handbook:SPARC/Installation/Base/de#Optional:_Adding_a_binary_package_host "Handbook:SPARC/Installation/Base/de") can add --getbinpkg (or -g) in order to fetch packages from the binary host instead of compiling them:

`root #` `emerge --ask --verbose --update --deep --newuse --getbinpkg @world`

#### Removing obsolete packages

It is important to always _depclean_ after system upgrades to remove obsolete packages. Review the output carefully with emerge --depclean --pretend to see if any of the to-be-cleaned packages should be kept if personally using them. To keep a package which would otherwise be depcleaned, use emerge --noreplace foo.

`root #` `emerge --ask --pretend --depclean`

If happy, then proceed with a real depclean:

`root #` `emerge --ask --depclean`

## Zeitzone

**Hinweis**

Dieser Schritt gilt nicht für Benutzer der musl libc. Benutzer die nicht wissen, was das bedeutet, sollten diesen Schritt ausführen.

Bitte vermeiden Sie die /usr/share/zoneinfo/Etc/GMT\* Zeitzonen, da deren Namen nicht die erwarteten Zonen anzeigen. Beispielsweise ist GMT-8 in der Tat GMT+8.

Wählen Sie die Zeitzone für das System. Schauen Sie nach den verfügbaren Zeitzonen in /usr/share/zoneinfo/.

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

Nehmen wir an, die gewünschte Zeitzone ist _Europe/Brussels_.

`root #` `ln -sf ../usr/share/zoneinfo/Europe/Brussels /etc/localtime`

**Tipp**

The target path with `../` at the start is _relative to the link location_, not the current directory.

**Hinweis**

An absolute path can be used for the symlink, but a relative link is also created by systemd's timedatectl and is more compatible with alternate _ROOT_ s.

## Locales konfigurieren

**Hinweis**

Dieser Abschnitt ist nicht erforderlich für Nutzer der "musl" libc. Anwender, die nicht wissen, was das bedeutet, sollten diesen Abschnitt durcharbeiten.

### Locales erzeugen

A default installation of Gentoo Linux provides an archive that contains all supported locales, numbering 500 or more. However, it is typical for an administrator to require only one or two of these. In that case, the /etc/locale.gen configuration file may be populated with a list of the required locales. By default, locale-gen shall read this file and compile only the locales that are specified, saving both time and space in the longer term.

Locales geben nicht nur die Sprache an, mit der der Anwender mit dem System interagieren soll. Sie definieren auch die Regeln zum Sortieren von Zeichenketten, zur Anzeige von Datum und Zeit, usw. Locales sind 'case-sensitive' und müssen genau so geschrieben werden wie vorgegeben. Die Datei /usr/share/i18n/SUPPORTED enthält eine vollständige Liste aller verfügbaren Locales.

`root #` `nano /etc/locale.gen`

Das folgende Beispiel zeigt Locales, um sowohl Englisch (Vereinigte Staaten), als auch Deutsch (Deutschland) mit den Zeichenkodierungen Latin-1 und UTF-8 zu erhalten:

DATEI **`/etc/locale.gen`** **Konfiguration von Locale und Zeichenkodierungen**

```
en_US.UTF-8 UTF-8
de_DE.UTF-8 UTF-8

```

**Tipp**

Anwender aus Österreich oder aus der Schweiz können anstelle von (oder zusätzlich zu) "de\_DE" wählen: "de\_AT" oder "de\_CH".

1. Non UTF-8 locales are discouraged and should only be used in special circumstances.
2. en\_US ISO-8859-1
3. de\_DE ISO-8859-1

}}

**Warnung**

Viele Anwendungen erfordern mindestens ein UTF-8 Locale, um korrekt gebaut zu werden.

Der nächste Schritt ist, locale-gen auszuführen. Dieses Programm erzeugt alle Locales, die in der Datei /etc/locale.gen angegeben sind.

`root #` `locale-gen`

Um zu überprüfen, ob die ausgewählten Locales jetzt verfügbar sind, führen Sie locale -a aus. Dieser Befehl muss die konfigurierten Locales anzeigen - ansonsten hat locale-gen nicht funktioniert.

Auf systemd-Installationen kann localectl verwendet werden, z.B. localectl set-locale ... oder localectl list-locales.

### Locale auswählen

Nachdem die Locales generiert wurden, ist es Zeit, die systemweiten Locale-Einstellungen zu setzen. Wir verwenden dafür wieder eselect, diesmal mit dem Modul `locale`.

Mit eselect locale list werden die verfügbaren Locales angezeigt:

`root #` `eselect locale list`

```
Verfügbare Locales für die LANG Variable:
  [1]  C
  [2]  C.UTF-8
  [3]  POSIX
  [4]  de_DE.UTF-8
  [5]  en_US.UTF-8
  [ ]  (free form)

```

Mit eselect locale set <WERT> kann die gewünschte Locale ausgewählt werden:

`root #` `eselect locale set 2`

Manuell kann dies auch durch Anpassen der Datei /etc/env.d/02locale und für Systemd durch Anpassen der Datei /etc/locale.conf erreicht werden:

DATEI **`/etc/env.d/02locale`** **Manuelles Setzen der System-Locale**

```
LANG="de_DE.UTF-8"
LC_COLLATE="C.UTF-8"

```

Das Setzen der System-Locale verhindert Warnungen und Fehlermeldungen beim Kompilieren des Kernels und von Software in den folgenden Installationsschritten.

Laden Sie jetzt die Umgebung erneut, damit die Änderung der Locale-Einstellung in Ihrer Shell wirksam wird:

`root #` `env-update && source /etc/profile && export PS1="(chroot) ${PS1}"`

Eine weiterführende Anleitung durch den Lokalisierungs-Prozess finden Sie im [Lokalisierungs-Leitfaden](/wiki/Localization/Guide/de "Localization/Guide/de") und in der [UTF-8](/wiki/UTF-8 "UTF-8")-Anleitung.

## References

1. [↑](#cite_ref-1)[https://www.gentoo.org/news/2023/12/29/Gentoo-binary.html](https://www.gentoo.org/news/2023/12/29/Gentoo-binary.html)
2. [↑](#cite_ref-2)[https://www.gentoo.org/glep/glep-0023.html#id7](https://www.gentoo.org/glep/glep-0023.html#id7)

[← Installation des Stage Archivs](/wiki/Handbook:SPARC/Installation/Stage/de "Previous part") [Anfang](/wiki/Handbook:SPARC/de "Handbook:SPARC/de") [Konfiguration des Kernels →](/wiki/Handbook:SPARC/Installation/Kernel/de "Next part")