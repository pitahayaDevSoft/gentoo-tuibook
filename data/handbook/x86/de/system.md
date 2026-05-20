# System

Other languages:

- Deutsch
- [English](/wiki/Handbook:X86/Installation/System "Handbook:X86/Installation/System (100% translated)")
- [español](/wiki/Handbook:X86/Installation/System/es "Manual de Gentoo: X86/Instalación/Sistema (100% translated)")
- [français](/wiki/Handbook:X86/Installation/System/fr "Handbook:X86/Installation/System/fr (100% translated)")
- [italiano](/wiki/Handbook:X86/Installation/System/it "Handbook:X86/Installation/System (100% translated)")
- [magyar](/wiki/Handbook:X86/Installation/System/hu "Handbook:X86/Installation/System/hu (100% translated)")
- [polski](/wiki/Handbook:X86/Installation/System/pl "Handbook:X86/Installation/System (100% translated)")
- [português do Brasil](/wiki/Handbook:X86/Installation/System/pt-br "Manual:X86/Instalação/Sistema (100% translated)")
- [čeština](/wiki/Handbook:X86/Installation/System/cs "Handbook:X86/Installation/System/cs (50% translated)")
- [русский](/wiki/Handbook:X86/Installation/System/ru "Handbook:X86/Installation/System (100% translated)")
- [தமிழ்](/wiki/Handbook:X86/Installation/System/ta "கையேடு:X86/நிறுவல்/முறைமை (100% translated)")
- [中文（中国大陆）‎](/wiki/Handbook:X86/Installation/System/zh-cn "手册：X86/安装/配置系统 (100% translated)")
- [日本語](/wiki/Handbook:X86/Installation/System/ja "ハンドブック:X86/インストール/システムの設定 (100% translated)")
- [한국어](/wiki/Handbook:X86/Installation/System/ko "Handbook:X86/Installation/System/ko (100% translated)")

[X86 Handbuch](/wiki/Handbook:X86/de "Handbook:X86/de")[Installation](/wiki/Handbook:X86/Full/Installation/de "Handbook:X86/Full/Installation/de")[Über die Installation](/wiki/Handbook:X86/Installation/About/de "Handbook:X86/Installation/About/de")[Auswahl des Mediums](/wiki/Handbook:X86/Installation/Media/de "Handbook:X86/Installation/Media/de")[Konfiguration des Netzwerks](/wiki/Handbook:X86/Installation/Networking/de "Handbook:X86/Installation/Networking/de")[Vorbereiten der Festplatte(n)](/wiki/Handbook:X86/Installation/Disks/de "Handbook:X86/Installation/Disks/de")[Installation des Stage Archivs](/wiki/Handbook:X86/Installation/Stage/de "Handbook:X86/Installation/Stage/de")[Installation des Basissystems](/wiki/Handbook:X86/Installation/Base/de "Handbook:X86/Installation/Base/de")[Konfiguration des Kernels](/wiki/Handbook:X86/Installation/Kernel/de "Handbook:X86/Installation/Kernel/de")Konfiguration des Systems[Installation der Tools](/wiki/Handbook:X86/Installation/Tools/de "Handbook:X86/Installation/Tools/de")[Konfiguration des Bootloaders](/wiki/Handbook:X86/Installation/Bootloader/de "Handbook:X86/Installation/Bootloader/de")[Abschluss](/wiki/Handbook:X86/Installation/Finalizing/de "Handbook:X86/Installation/Finalizing/de")[Arbeiten mit Gentoo](/wiki/Handbook:X86/Full/Working/de "Handbook:X86/Full/Working/de")[Portage-Einführung](/wiki/Handbook:X86/Working/Portage/de "Handbook:X86/Working/Portage/de")[USE-Flags](/wiki/Handbook:X86/Working/USE/de "Handbook:X86/Working/USE/de")[Portage-Features](/wiki/Handbook:X86/Working/Features/de "Handbook:X86/Working/Features/de")[Initskript-System](/wiki/Handbook:X86/Working/Initscripts/de "Handbook:X86/Working/Initscripts/de")[Umgebungsvariablen](/wiki/Handbook:X86/Working/EnvVar/de "Handbook:X86/Working/EnvVar/de")[Arbeiten mit Portage](/wiki/Handbook:X86/Full/Portage/de "Handbook:X86/Full/Portage/de")[Dateien und Verzeichnisse](/wiki/Handbook:X86/Portage/Files/de "Handbook:X86/Portage/Files/de")[Variablen](/wiki/Handbook:X86/Portage/Variables/de "Handbook:X86/Portage/Variables/de")[Mischen von Softwarezweigen](/wiki/Handbook:X86/Portage/Branches/de "Handbook:X86/Portage/Branches/de")[Zusätzliche Tools](/wiki/Handbook:X86/Portage/Tools/de "Handbook:X86/Portage/Tools/de")[Eigener Portage-Tree](/wiki/Handbook:X86/Portage/CustomTree/de "Handbook:X86/Portage/CustomTree/de")[Erweiterte Portage-Features](/wiki/Handbook:X86/Portage/Advanced/de "Handbook:X86/Portage/Advanced/de")[Netzwerk-Konfiguration](/wiki/Handbook:X86/Full/Networking/de "Handbook:X86/Full/Networking/de")[Zu Beginn](/wiki/Handbook:X86/Networking/Introduction/de "Handbook:X86/Networking/Introduction/de")[Fortgeschrittene Konfiguration](/wiki/Handbook:X86/Networking/Advanced/de "Handbook:X86/Networking/Advanced/de")[Modulare Vernetzung](/wiki/Handbook:X86/Networking/Modular/de "Handbook:X86/Networking/Modular/de")[Drahtlose Netzwerke](/wiki/Handbook:X86/Networking/Wireless/de "Handbook:X86/Networking/Wireless/de")[Funktionalität hinzufügen](/wiki/Handbook:X86/Networking/Extending/de "Handbook:X86/Networking/Extending/de")[Dynamisches Management](/wiki/Handbook:X86/Networking/Dynamic/de "Handbook:X86/Networking/Dynamic/de")

## Contents

- [1Dateisystem-Information](#Dateisystem-Information)
  - [1.1Dateisystem Labels und UUIDs](#Dateisystem_Labels_und_UUIDs)
  - [1.2Partitionslabel und UUIDs](#Partitionslabel_und_UUIDs)
  - [1.3Über fstab](#.C3.9Cber_fstab)
  - [1.4Die fstab-Datei erstellen](#Die_fstab-Datei_erstellen)
    - [1.4.1DOS/Legacy BIOS systems](#DOS.2FLegacy_BIOS_systems)
    - [1.4.2DPS UEFI PARTUUID](#DPS_UEFI_PARTUUID)
- [2Netzwerk-Konfiguration](#Netzwerk-Konfiguration)
  - [2.1Hostname](#Hostname)
    - [2.1.1Setzen des Hostnamens (OpenRC oder systemd)](#Setzen_des_Hostnamens_.28OpenRC_oder_systemd.29)
    - [2.1.2systemd](#systemd)
  - [2.2Netzwerk](#Netzwerk)
    - [2.2.1DHCP mit dhcpcd (bei allen Init-Systemen)](#DHCP_mit_dhcpcd_.28bei_allen_Init-Systemen.29)
    - [2.2.2netifrc (OpenRC)](#netifrc_.28OpenRC.29)
      - [2.2.2.1Konfigurieren des Netzwerks](#Konfigurieren_des_Netzwerks)
      - [2.2.2.2Automatischer Start der Netzwerk-Interfaces beim Booten](#Automatischer_Start_der_Netzwerk-Interfaces_beim_Booten)
  - [2.3Die hosts Datei](#Die_hosts_Datei)
  - [2.4Optional: PCMCIA zum Laufen bekommen](#Optional:_PCMCIA_zum_Laufen_bekommen)
- [3System-Konfiguration](#System-Konfiguration)
  - [3.1Root Passwort](#Root_Passwort)
  - [3.2Init- and Boot-Konfiguration](#Init-_and_Boot-Konfiguration)
    - [3.2.1OpenRC](#OpenRC)
    - [3.2.2systemd](#systemd_2)

## Dateisystem-Information

#### Dateisystem Labels und UUIDs

Unabhängig davon, ob Sie MBR (BIOS) oder GPT Partitionstabellen verwenden, können Sie _Dateisystem_ Labels und _Dateisystem_ UUIDs nutzen. Diese Attribute können in /etc/fstab als Alternative zu den bisherigen Block-Gerätedateien (/dev/sd\*) angegeben werden. Das Kommando blkid zeigt Ihnen die LABELs und UUIDs der Dateisysteme auf Ihrem System an. In der Datei /etc/fstab geben Sie diese mit dem Prefix "LABEL=" bzw. "UUID=" an. Anführungszeichen werden - im Gegensatz zu der Ausgabe von blkid - nicht verwendet.

`root #` `blkid`

**Warnung**

Wenn ein Dateisystem innerhalb einer Partition neu erstellt oder gelöscht wird, ändern sich die Dateisystem Labels und UUIDs - oder sie verschwinden ganz.

Um die Eindeutigkeit zu gewährleisten, sollten Anwender, die eine MBR Partitionstabelle verwenden, besser Dateisystem UUIDs als Dateisystem Labels in /etc/fstab verwenden.

**Wichtig**

UUIDs des Dateisystems auf einem LVM volume und von LVM Snapshots von diesem sind identisch. Deshalb sollten keine UUIDs zum Mounten von LVM volumes verwendet werden.

#### Partitionslabel und UUIDs

Anwender, die eine GPT Partitionstabelle verwenden, haben eine 'robustere' Möglichkeit, um Partitionen in /etc/fstab anzugeben: _Partition_ Labels und _Partition_ UUIDs. Diese kennzeichnen Partitionen selbst, unabhängig von deren Inhalt oder dem Dateisystem, das dort angelegt ist. Sie ändern sich deshalb auch nicht, wenn der Inhalt der Partition gelöscht wird oder ein neues Dateisystem erstellt wird. Das Kommando blkid zeigt Ihnen die PARTLABELs und PARTUUIDs der Partitionen auf Ihrem System an. In der Datei /etc/fstab geben Sie diese mit dem Prefix "PARTLABEL=" bzw. "PARTUUID=" an. Anführungszeichen werden - im Gegensatz zu der Ausgabe von blkid - nicht verwendet.

Output for an **amd64** EFI system using the Discoverable Partition Specification UUIDs may like the following:

`root #` `blkid`

Der Name eines Block-Geräts hängt von mehreren Faktoren ab, u.a. von der Reihenfolge, in der der Kernel die Block-Geräte im frühen Boot-Prozess erkennt. Bei Systemen, die häufig gebootet werden und bei denen regelmäßig SATA Block-Geräte entfernt oder hinzugefügt werden, können sich die Namen der Block-Geräte nach einem Neustart ändern. Es ist deshalb riskant, die älteren Block-Gerätedateien (/dev/sd\*N) zur Angabe von Partitionen in fstab zu verwenden. Wenn Sie stattdessen Partition UUIDs verwenden, ist garantiert, dass Linux das gewünschte Dateisystem verwendet - selbst wenn das Dateisystem später geändert wird.

Nichtsdestotrotz ist die Verwendung der hergebrachten Block-Gerätedateien eine einfache, geradlinige und für die meisten Anwender sinnvolle Methode. Wenn Sie einen komplexen Server mit vielen Festplatten haben oder wenn Sie die Hardware ihres Systems häufiger ändern werden, sollten Sie über die Verwendung von Partition UUIDs nachdenken.

### Über fstab

Unter Linux müssen alle Partitionen, die im System genutzt werden, in [/etc/fstab](/wiki//etc/fstab/de "/etc/fstab/de") aufgelistet werden. Diese Datei beinhaltet die Mountpunkte ("Einhängepunkte") dieser Partitionen ( _wo_ sie im Dateisystem erscheinen), _wie_ sie eingehängt werden sollen und _mit welchen speziellen Optionen_ sie eingehängt werden sollen (automatisch einhängen oder nicht, dürfen Benutzer sie einhängen oder nicht, etc.)

### Die fstab-Datei erstellen

**Hinweis**

If the init system being used is systemd, the partition UUIDs conform to the Discoverable Partition Specification as given in [Preparing the disks](/wiki/Handbook:X86/Installation/Disks/de "Handbook:X86/Installation/Disks/de"), and the system uses UEFI, then creating an fstab can be skipped, since systemd auto-mounts partitions that follow the spec.

Die Datei /etc/fstab verwendet eine tabellenartige Syntax. Jede Zeile besteht aus sechs Feldern, die jeweils durch Leerräume (Leerzeichen, Tabulatoren oder beides gemischt) getrennt sind. Jedes Feld hat seine eigene Bedeutung:

1. Das erste Feld enthält eine Block-Gerätedatei oder ein Remote-Dateisystem, die/das eingehängt werden soll. Block-Gerätedateien können über mehrere verschiedene Arten angegeben werden: u.a. über den Dateinamen des Gerätedatei, über Dateisystem Labels und UUIDs oder über Partition Labels und UUIDs.
2. Das zweite Feld definiert den Einhängepunkt, an dem die Partition eingehängt werden soll.
3. Das dritte Feld enthält den Typ des Dateisystem (ext2, etxt3, ...)
4. Im vierten Feld stehen die Einhänge-Optionen, die von mount genutzt werden, wenn es die Partition eingehängt. Da jedes Dateisystem seine eigenen Optionen hat, empfiehlt sich ein Blick in die Manpage (man mount), wo sich eine vollständige Liste findet. Mehrere Einhänge-Optionen werden mit Kommata getrennt.
5. Das fünfte Feld wird von dump verwendet, um herauszufinden ob die Partition in einem Dump-Backup berücksichtigt werden soll. Dieser Eintrag kann üblicherweise auf `0` (null) belassen werden.
6. Das sechste Feld wird von fsck verwendet, um die Reihenfolge festzulegen, in der Dateisysteme nach einem unsauberen Neustart überprüft werden. Für das Wurzeldateisystem (/) sollte hier `1` stehen, für alle anderen Dateisysteme `2` (oder `0`, wenn eine Dateisystemprüfung nicht nötig ist.)

**Wichtig**

Die bei der Installation von Gentoo Linux installierte Datei /etc/fstab ist _keine_ gültige fstab-Datei, sondern eine Vorlage, die von Ihnen ausgefüllt werden muss.

`root #` `nano /etc/fstab`

#### DOS/Legacy BIOS systems

Schauen wir uns an, wie man die /boot/-Partition in fstab konfigurieren würde. Das folgende Beispiel sollte so angepasst werden, dass es zu der gewählten Partitionierung Ihres Systems passt.
Bei der in unserem x86 Handbuch gewählten Partitionierung ist /boot/ die /dev/sda1 Partition, mit einem xfs Dateisystem. Das Dateisystem soll beim Booten überprüft werden. Es ergibt sich:

DATEI **`/etc/fstab`** **Beispiel: eine /boot Zeile für eine /etc/fstab Datei**

```
# Bitte passen Sie Formatierungs-Unterschiede vom Schritt "Vorbereiten der Festplatten" an
/dev/sda1   /boot     ext4    defaults        0 2

```

Einige Anwender möchten aus Sicherheitsgründen nicht, dass ihre /boot/ Partition automatisch eingehängt wird. Diese Anwender können _defaults_ durch _noauto_ ersetzen. Dies bedeutet aber auch, dass diese Anwender die Partition jedes Mal von Hand einhängen müssen, wenn Sie sie nutzen wollen.

Fügen Sie weitere Zeilen hinzu, so dass alle gewünschten Dateisysteme eingehängt werden. Wenn Sie ein CD-ROM Laufwerk haben, fügen Sie auch eine Regel hierfür hinzu.

Hier ist ein Beispiel für eine vollständige /etc/fstab Datei:

DATEI **`/etc/fstab`** **Beispiel: eine vollständige /etc/fstab Datei**

```
Bitte passen Sie Formatierungs-Unterschiede und zusätzliche Partitionen vom Schritt "Vorbereiten der Festplatten" an
/dev/sda1   /boot        ext4    defaults    0 2
/dev/sda2   none         swap    sw                   0 0
/dev/sda3   /            xfs    defaults,noatime              0 1

/dev/cdrom  /mnt/cdrom   auto    noauto,user          0 0

```

Bitte passen Sie Formatierungs-Unterschiede vom Schritt "Vorbereiten der Festplatten" an

/dev/cdrom /mnt/cdrom auto noauto,user 0 0
}}

#### DPS UEFI PARTUUID

Below is an example of an /etc/fstab file for a disk formatted with a GPT disklabel and Discoverable Partition Specification (DPS) UUIDs set for UEFI firmware:

DATEI **`/etc/fstab`** **GPT disklabel DPS PARTUUID fstab example**

```
# Adjust any formatting difference and additional partitions created from the "Preparing the disks" step.
# This example shows a GPT disklabel with Discoverable Partition Specification (DSP) UUID set:
PARTUUID=c12a7328-f81f-11d2-ba4b-00a0c93ec93b   /efi        vfat    umask=0077,tz=UTC            0 2
PARTUUID=0657fd6d-a4ab-43c4-84e5-0933c84b4f4f   none            sw                           0 0
PARTUUID=44479540-f297-41b2-9af7-d131d5f0458a   /           xfs    defaults,noatime              0 1

```

When `auto` im dritten Feld verwendet wird, "rät" mount den Typ des Dateisystems beim Einhängen. Dies wird empfohlen für Wechselmedien, da sie unterschiedliche Typen von Dateisystemen haben können. Die Option `user` im vierten Feld ermöglicht es nicht-root Usern, CDs einzuhängen.

To improve performance, most users would want to add the `noatime` mount option, which results in a faster system since access times are not registered (those are not needed generally anyway). This is also recommended for systems with solid state drives (SSDs). Users may wish to consider `lazytime` instead.

**Tipp**

Es wird nicht empfohlen, die Mount-Option `discard` in /etc/fstab zu verwenden - weil es die Performance verringern kann. Stattdessen ist es besser, regelmäßig Block-Discards auszuführen. Dies kann über einen Job Scheduler wie cron oder mit einem Timer (systemd) erfolgen. Weitere Informationen finden Sie in dem Abschnitt [Periodische fstrim Jobs](/wiki/SSD#Periodic_fstrim_jobs "SSD").

Überprüfen Sie die Datei /etc/fstab noch einmal, speichern Sie sie und verlassen Sie den Editor.

## Netzwerk-Konfiguration

Es ist wichtig zu beachten, dass die folgenden Abschnitte dem Leser helfen sollen, sein System schnell für die Teilnahme an einem lokalen Netzwerk einzurichten.

Für Systeme, auf denen OpenRC läuft, ist eine detailliertere Referenz für die Netzwerkeinrichtung im Abschnitt [Erweiterte Netzwerkkonfiguration](/wiki/Handbook:X86/Networking/Introduction/de "Handbook:X86/Networking/Introduction/de") verfügbar, der gegen Ende des Handbuchs behandelt wird. Systeme mit spezielleren Netzwerkanforderungen müssen eventuell den Abschnitt überspringen und dann hierher zurückkehren, um mit dem Rest der Installation fortzufahren.

Für spezifischere Systemd-Netzwerkeinstellungen lesen Sie bitte den [Netzwerkteil](/wiki/Systemd/de#Network "Systemd/de") des [systemd](/wiki/Systemd/de "Systemd/de")-Artikels.

### Hostname

Eine der Entscheidungen, die System-Administratoren treffen müssen, ist der Name ihres PCs. Auf den ersten Blick scheint dies einfach zu sein, aber viele Benutzer haben Schwierigkeiten, einen passenden Namen für den hostname zu finden. Um diesen Prozess zu beschleunigen, sei darauf hingewiesen, dass der Name später wieder geändert werden kann. In den folgenden Beispielen wird der Hostname _tux_ verwendet.

#### Setzen des Hostnamens (OpenRC oder systemd)

`root #` `echo tux > /etc/hostname`

#### systemd

Um den System-Hostnamen für ein System, auf dem systemd läuft, zu setzen, kann das Dienstprogramm hostnamectl verwendet werden.

Um den Hostname auf "tux" zu setzen, würde man ausführen:

`root #` `hostnamectl hostname tux`

Hilfe erhält man durch Ausführung von hostnamectl --help oder man 1 hostnamectl.

### Netzwerk

Es gibt _viele_ verschiedene Alternativen, mit denen das Netzwerk konfiguriert werden kann. Dieser Abschnitt behandelt nur ein paar davon. Wählen Sie die Methode, die am besten zu Ihren Anforderungen passt.

#### DHCP mit dhcpcd (bei allen Init-Systemen)

In den meisten LAN Netzen läuft ein DHCP Server. Wenn dies der Fall ist, wird empfohlen, das Programm dhcpcd zu verwenden, um eine IP-Adresse zu erhalten.

Zur Installation:

`root #` `emerge --ask net-misc/dhcpcd`

Um den Service auf OpenRC Systemen zu aktivieren und danach zu starten:

`root #` `rc-update add dhcpcd default
`

`root #` `rc-service dhcpcd start
`

Um den Service auf systemd Systemen zu aktivieren und danach zu starten:

`root #` `systemctl enable --now dhcpcd`

Nachdem diese Schritte durchgeführt wurden, sollte dhcpcd beim nächsten Systemstart eine IP Adresse vom DHCP Server erhalten.
Weitere Details finden Sie im [Dhcpcd](/wiki/Dhcpcd "Dhcpcd") Artikel.

#### netifrc (OpenRC)

**Tipp**

Dieser Abschnitt beschreibt eine Netzwerk-Konfiguration mit [Netifrc](/wiki/Netifrc "Netifrc") unter OpenRC. Für einfache Anschaltungen kann auch [Dhcpcd](/wiki/Dhcpcd "Dhcpcd") verwendet werden.

##### Konfigurieren des Netzwerks

Bereits zu Beginn der Installation von Gentoo Linux wurde das Netzwerk konfiguriert. Diese Konfiguration betraf jedoch das Netzwerk der Live-Umgebung - und nicht das Netzwerk des neu installierten Systems. Wir werden jetzt die Netzwerk-Konfiguration für Ihr neu installiertes Linux-System erstellen.

**Hinweis**

Weitere Informationen über Netzwerke, insbesondere auch über fortgeschrittene Themen wie Bonding, Bridges, 802.1Q VLANs oder WLAN, finden Sie in dem Abschnitt [Fortgeschrittene Netzwerk-Konfiguration](/wiki/Handbook:X86/Networking/Introduction/de "Handbook:X86/Networking/Introduction/de").

Die Netzwerk-Konfiguration wird gespeichert in /etc/conf.d/net. Die Syntax ist unkompliziert, aber vielleicht etwas un-intuitiv. Aber keine Angst - wir werden alles erklären. Ein gut dokumentiertes Beispiel mit mehreren verschiedenen Konfigurationen finden Sie unter /usr/share/doc/netifrc-\*/net.example.bz2.

Installieren Sie zuerst das Paket [net-misc/netifrc](https://packages.gentoo.org/packages/net-misc/netifrc):

`root #` `emerge --ask --noreplace net-misc/netifrc`

Standardmäßig wird DHCP verwendet. Damit DHCP funktioniert, muss ein DHCP-Client installiert werden. Dies wird später im Abschnitt "Installation von erforderlichen System-Tools" beschrieben.

Wenn Sie kein DHCP verwenden wollen (statische IP-Adressen) oder wenn Sie spezielle DHCP-Optionen benötigen, dann editieren Sie jetzt die Datei /etc/conf.d/net:

`root #` `nano /etc/conf.d/net`

Definieren Sie IP-Adresse und Routing in den beiden Variablen config\_eth0 und routes\_eth0.

**Hinweis**

Wir gehen davon aus, dass das Netzwerk-Interface "eth0" heißt. Bei Ihnen heißt das Netzwerk-Interface möglicherweise anders. Verwenden Sie im Folgenden statt "eth0" immer den Namen Ihres Netzwerk-Interfaces. Der Name des Netzwerk-Interfaces hängt ab vom gebooteten System und den Kernel-Optionen. Da Sie noch nicht von dem neu installierten System gebootet haben, können Sie den Namen des Netzwerk-Interfaces noch nicht wissen. Wenn das Installationsmedium aktuell ist, ist die Wahrscheinlichkeit hoch, dass das Netzwerk-Interface genauso heißen wird, wie in Ihrem jetzigen System, das von dem Installationsmedium gebootet wurde. Weitere Informationen finden Sie im Abschnitt [Benennung von Netzwerkschnittstellen](/wiki/Handbook:X86/Networking/Advanced/de#Network_interface_naming "Handbook:X86/Networking/Advanced/de").

DATEI **`/etc/conf.d/net`** **Konfiguration einer statischen IP-Adresse**

```
config_eth0="192.168.0.2 netmask 255.255.255.0 brd 192.168.0.255"
routes_eth0="default via 192.168.0.1"

```

Um DHCP zu verwenden, setzen Sie config\_eth0:

DATEI **`/etc/conf.d/net`** **Konfiguration von DHCP**

```
config_eth0="dhcp"

```

Eine Liste zusätzlicher Konfigurations-Optionen finden Sie in /usr/share/doc/netifrc-\*/net.example.bz2. Bitte lesen Sie auch die DHCP-Client man page, wenn besondere DHCP-Optionen gesetzt werden müssen.

Wenn das System mehrere Netzwerk-Interfaces hat, wiederholen Sie bitte die oben beschriebenen Schritte für alle Netzwerk-Interfaces (config\_eth1, config\_eth2 usw.) - falls diese Interfaces beim Booten initialisiert und aktiviert werden sollen.

Speichern Sie die Konfigurations-Datei und verlassen Sie den Editor.

##### Automatischer Start der Netzwerk-Interfaces beim Booten

Damit die Netzwerk-Interfaces beim Booten konfiguriert und aktiviert werden, müssen sie zum Runlevel 'default' hinzugefügt werden.

`root #` `cd /etc/init.d
`

`root #` `ln -s net.lo net.eth0
`

`root #` `rc-update add net.eth0 default`

Wenn Ihr System mehrere Netzwerk-Interfaces hat, muss der vorherige Schritt für alle Netzwerk-Interfaces, die beim Booten konfiguriert und aktiviert werden sollen, wiederholt werden.

Wenn sich nach dem Booten herausstellen sollte, dass der gewählte Name für das Netzwerk-Interface verkehrt ist, führen Sie die folgenden Schritte aus, um das Problem zu beheben:

1. Editieren Sie die Datei /etc/conf.d/net und ersetzen Sie den verkehrten Interface-Namen durch den korrekten Namen (beispielsweise `enp3s0` oder `enp5s0` statt `eth0`).
2. Erstellen Sie den korrekten symbolischen Link (beispielsweise /etc/init.d/net.enp3s0).
3. Entfernen Sie den alten (fehlerhaften) Link (rm /etc/init.d/net.eth0).
4. Fügen Sie das neue Interface zum Runlevel 'default' hinzu
5. Entfernen Sie das alte Interface vom Runlevel 'default': rc-update del net.eth0 default.

### Die hosts Datei

Bitte editieren Sie die Datei /etc/hosts. Diese Datei muss auf jeden Fall einen korrekten Eintrag zu localhost enthalten. Weiterhin können Sie IP-Adressen und Hostnamen von wichtigen Hosts in ihrem _eigenen_ Netzwerk eintragen. Letzteres ist jedoch nur notwendig, wenn Ihr Nameserver diese Informationen nicht liefern kann, wenn Sie gar keinen Nameserver verwenden, oder wenn Sie eine Namensauflösung für die Zeiten benötigen, in denen der Nameserver nicht verfügbar ist (z.B. beim Booten oder bei Netzstörungen).

`root #` `nano /etc/hosts`

DATEI **`/etc/hosts`** **Zuordnung zwischen IP-Adressen und Hostnamen für lokale Namensauflösung**

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

Speichern Sie die Datei und beenden Sie den Editor.

### Optional: PCMCIA zum Laufen bekommen

**x86** Systeme, die PCMCIA-Unterstützung benötigen, sollten nun das Paket [sys-apps/pcmciautils](https://packages.gentoo.org/packages/sys-apps/pcmciautils) installieren.

`root #` `emerge --ask sys-apps/pcmciautils`

## System-Konfiguration

### Root Passwort

Setzen Sie das root-Passwort mit dem passwd Kommando:

`root #` `passwd`

Später werden wir einen gewöhnlichen User mit eingeschränkten Rechten anlegen, unter dem Sie alle normalen täglichen Arbeiten verrichten können.

### Init- and Boot-Konfiguration

#### OpenRC

Wenn Gentoo mit OpenRC verwendet wird, werden die Dienste, die beim Booten oder Herunterfahren des Systems gestartet bzw. gestoppt werden, in der Datei /etc/rc.conf konfiguriert. Öffnen Sie diese Datei mit einem Editor, und erfreuen Sie sich an den vielen Kommentaren in der Datei. Überprüfen Sie alle Einstellungen und ändern Sie sie, falls gewünscht oder erforderlich.

`root #` `nano /etc/rc.conf`

Editieren Sie als nächstes die Datei /etc/conf.d/keymaps und konfigurieren Sie Ihre Tastatur.

`root #` `nano /etc/conf.d/keymaps`

Seien Sie vorsichtig bei der keymap Variable. Wenn Sie die falsche Tastatur konfigurieren, erhalten Sie merkwürdige Ergebnisse, wenn sie Texte auf der Tastatur tippen.

Editieren Sie zum Schluss die Datei /etc/conf.d/hwclock und konfigurieren Sie Ihre Hardware-Uhr.

`root #` `nano /etc/conf.d/hwclock`

Wenn die Hardware-Uhr nicht unter der Zeitzone UTC laufen soll, sollten Sie `clock="local"` in die Datei schreiben. Ansonsten kann es zu Zeitfehlern oder -sprüngen kommen.

#### systemd

Zunächst wird empfohlen, systemd-firstboot auszuführen, um verschiedene Komponenten des Systems für den ersten Start in die neue systemd-Umgebung korrekt einzustellen. Durch die Übergabe der folgenden Optionen wird der Benutzer aufgefordert, eine Locale, eine Zeitzone, einen Hostnamen, ein Root-Passwort und Root-Shell-Werte festzulegen. Außerdem wird der Installation eine zufällige Maschinen-ID zugewiesen:

`root #` `systemd-firstboot --prompt --setup-machine-id`

Als Nächstes sollten Benutzer systemctl ausführen, um alle installierten Unit-Dateien auf die voreingestellten Richtwerte zurückzusetzen:

**Hinweis**

It is possible that a failure warning will print out after running the following command. This is normal, as Gentoo's LiveCD uses OpenRC.

`root #` `systemctl preset-all --preset-mode=enable-only`

Es ist möglich, die vollständigen Voreinstellungsänderungen auszuführen, aber dies kann alle Dienste zurücksetzen, die während des Prozesses bereits konfiguriert wurden:

`root #` `systemctl preset-all`

Diese beiden Schritte sorgen für einen reibungslosen Übergang von der Live-Umgebung zum ersten Start der Installation.

[← Konfiguration des Kernels](/wiki/Handbook:X86/Installation/Kernel/de "Previous part") [Anfang](/wiki/Handbook:X86/de "Handbook:X86/de") [Installation der Tools →](/wiki/Handbook:X86/Installation/Tools/de "Next part")