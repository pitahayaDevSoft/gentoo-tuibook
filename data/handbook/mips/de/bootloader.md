# Bootloader

Other languages:

- Deutsch
- [English](/wiki/Handbook:MIPS/Installation/Bootloader "Handbook:MIPS/Installation/Bootloader (100% translated)")
- [español](/wiki/Handbook:MIPS/Installation/Bootloader/es "Manual de Gentoo: MIPS/Instalación/Arranque (100% translated)")
- [français](/wiki/Handbook:MIPS/Installation/Bootloader/fr "Handbook:MIPS/Installation/Bootloader/fr (100% translated)")
- [italiano](/wiki/Handbook:MIPS/Installation/Bootloader/it "Handbook:MIPS/Installation/Bootloader/it (100% translated)")
- [magyar](/wiki/Handbook:MIPS/Installation/Bootloader/hu "Handbook:MIPS/Installation/Bootloader/hu (100% translated)")
- [polski](/wiki/Handbook:MIPS/Installation/Bootloader/pl "Handbook:MIPS/Installation/Bootloader (100% translated)")
- [português do Brasil](/wiki/Handbook:MIPS/Installation/Bootloader/pt-br "Manual:MIPS/Instalação/Gerenciador de Boot (100% translated)")
- [русский](/wiki/Handbook:MIPS/Installation/Bootloader/ru "Handbook:MIPS/Installation/Bootloader (100% translated)")
- [தமிழ்](/wiki/Handbook:MIPS/Installation/Bootloader/ta "கையேடு:MIPS/நிறுவல்/துவக்கஏற்றி (100% translated)")
- [中文（中国大陆）‎](/wiki/Handbook:MIPS/Installation/Bootloader/zh-cn "手册：MIPS/安装/配置系统引导程序Bootloader (100% translated)")
- [日本語](/wiki/Handbook:MIPS/Installation/Bootloader/ja "ハンドブック:MIPS/インストール/ブートローダー (100% translated)")
- [한국어](/wiki/Handbook:MIPS/Installation/Bootloader/ko "Handbook:MIPS/Installation/Bootloader/ko (100% translated)")

[MIPS Handbuch](/wiki/Handbook:MIPS/de "Handbook:MIPS/de")[Installation](/wiki/Handbook:MIPS/Full/Installation/de "Handbook:MIPS/Full/Installation/de")[Über die Installation](/wiki/Handbook:MIPS/Installation/About/de "Handbook:MIPS/Installation/About/de")[Auswahl des Mediums](/wiki/Handbook:MIPS/Installation/Media/de "Handbook:MIPS/Installation/Media/de")[Konfiguration des Netzwerks](/wiki/Handbook:MIPS/Installation/Networking/de "Handbook:MIPS/Installation/Networking/de")[Vorbereiten der Festplatte(n)](/wiki/Handbook:MIPS/Installation/Disks/de "Handbook:MIPS/Installation/Disks/de")[Installation des Stage Archivs](/wiki/Handbook:MIPS/Installation/Stage/de "Handbook:MIPS/Installation/Stage/de")[Installation des Basissystems](/wiki/Handbook:MIPS/Installation/Base/de "Handbook:MIPS/Installation/Base/de")[Konfiguration des Kernels](/wiki/Handbook:MIPS/Installation/Kernel/de "Handbook:MIPS/Installation/Kernel/de")[Konfiguration des Systems](/wiki/Handbook:MIPS/Installation/System/de "Handbook:MIPS/Installation/System/de")[Installation der Tools](/wiki/Handbook:MIPS/Installation/Tools/de "Handbook:MIPS/Installation/Tools/de")Konfiguration des Bootloaders[Abschluss](/wiki/Handbook:MIPS/Installation/Finalizing/de "Handbook:MIPS/Installation/Finalizing/de")[Arbeiten mit Gentoo](/wiki/Handbook:MIPS/Full/Working/de "Handbook:MIPS/Full/Working/de")[Portage-Einführung](/wiki/Handbook:MIPS/Working/Portage/de "Handbook:MIPS/Working/Portage/de")[USE-Flags](/wiki/Handbook:MIPS/Working/USE/de "Handbook:MIPS/Working/USE/de")[Portage-Features](/wiki/Handbook:MIPS/Working/Features/de "Handbook:MIPS/Working/Features/de")[Initskript-System](/wiki/Handbook:MIPS/Working/Initscripts/de "Handbook:MIPS/Working/Initscripts/de")[Umgebungsvariablen](/wiki/Handbook:MIPS/Working/EnvVar/de "Handbook:MIPS/Working/EnvVar/de")[Arbeiten mit Portage](/wiki/Handbook:MIPS/Full/Portage/de "Handbook:MIPS/Full/Portage/de")[Dateien und Verzeichnisse](/wiki/Handbook:MIPS/Portage/Files/de "Handbook:MIPS/Portage/Files/de")[Variablen](/wiki/Handbook:MIPS/Portage/Variables/de "Handbook:MIPS/Portage/Variables/de")[Mischen von Softwarezweigen](/wiki/Handbook:MIPS/Portage/Branches/de "Handbook:MIPS/Portage/Branches/de")[Zusätzliche Tools](/wiki/Handbook:MIPS/Portage/Tools/de "Handbook:MIPS/Portage/Tools/de")[Eigener Portage-Tree](/wiki/Handbook:MIPS/Portage/CustomTree/de "Handbook:MIPS/Portage/CustomTree/de")[Erweiterte Portage-Features](/wiki/Handbook:MIPS/Portage/Advanced/de "Handbook:MIPS/Portage/Advanced/de")[Netzwerk-Konfiguration](/wiki/Handbook:MIPS/Full/Networking/de "Handbook:MIPS/Full/Networking/de")[Zu Beginn](/wiki/Handbook:MIPS/Networking/Introduction/de "Handbook:MIPS/Networking/Introduction/de")[Fortgeschrittene Konfiguration](/wiki/Handbook:MIPS/Networking/Advanced/de "Handbook:MIPS/Networking/Advanced/de")[Modulare Vernetzung](/wiki/Handbook:MIPS/Networking/Modular/de "Handbook:MIPS/Networking/Modular/de")[Drahtlose Netzwerke](/wiki/Handbook:MIPS/Networking/Wireless/de "Handbook:MIPS/Networking/Wireless/de")[Funktionalität hinzufügen](/wiki/Handbook:MIPS/Networking/Extending/de "Handbook:MIPS/Networking/Extending/de")[Dynamisches Management](/wiki/Handbook:MIPS/Networking/Dynamic/de "Handbook:MIPS/Networking/Dynamic/de")

## Contents

- [1arcload für Silicon Graphics Maschinen](#arcload_f.C3.BCr_Silicon_Graphics_Maschinen)
- [2CoLo für Cobalt MicroServer](#CoLo_f.C3.BCr_Cobalt_MicroServer)
  - [2.1CoLo installieren](#CoLo_installieren)
  - [2.2CoLo konfigurieren](#CoLo_konfigurieren)
- [3Serielle Konsole einrichten](#Serielle_Konsole_einrichten)
- [4SGI PROM Anpassung](#SGI_PROM_Anpassung)
  - [4.1Allgemeine PROM Einstellungen](#Allgemeine_PROM_Einstellungen)
  - [4.2Einstellungen für direktes booten vom Volume-Header](#Einstellungen_f.C3.BCr_direktes_booten_vom_Volume-Header)
  - [4.3arcload Einstellungen](#arcload_Einstellungen)
- [5Neustart des Systems](#Neustart_des_Systems)

## arcload für Silicon Graphics Maschinen

arcload wurde für Maschinen geschrieben, die 64-Bit-Kernel benötigen und daher arcboot nicht verwenden können (das nicht ohne weiteres als 64-Bit-Binärdatei kompiliert werden kann). Es umgeht auch die Besonderheiten, die beim Laden von Kernel direkt aus dem Volume-Header auftreten. Um mit der Installation fortzufahren, beginnen Sie mit:

`root #` `emerge arcload dvhtool`

Wenn das abgeschlossen ist, finden Sie die `arcload` Binärdatei in /usr/lib/arcload/. Es gibt zwei Dateien:

- sashARCS: Die 32-Bit Binärdatei für Indy, Indigo2 (R4k), Challenge S und O2 Systeme
- sash64: Die 64-Bit Binärdatei für Octane/Octane2, Origin 200/2000 und Indigo2 Impact Systeme

- sashARCS: The 32-bit binary for Indy, Indigo2 (R4k), Challenge S and O2 systems
- sash64: The 64-bit binary for Octane/Octane2, Origin 200/2000 and Indigo2 Impact systems

Verwenden Sie `dvhtool` um die für das System geeignete Binärdatei in den Volume-Header zu installieren:

Für Indy/Indigo2/Challenge S/O2 Benutzer:

`root #` `dvhtool --unix-to-vh /usr/lib/arcload/sashARCS sashARCS`

Für Indigo2 Impact/Octane/Octane2/Origin 200/Origin 2000 Benutzer:

`root #` `dvhtool --unix-to-vh /usr/lib/arcload/sash64 sash64`

**Hinweis**

Der Name sashARCS oder sash64 muss nicht unbedingt verwendet werden, es sei denn Sie installieren auf den Volume-Header einer bootfähigen CD. Zum normalen Booten von der Festplatte, kann es beliebig benannt werden.

Verwenden Sie jetzt einfach `dvhtool` um zu prüfen, dass sie sich im Volume-Header befindet:

`root #` `dvhtool --print-volume-directory`

```
----- directory entries -----
Entry #0, name "sash64", start 4, bytes 55859

```

Die Datei arc.cf hat eine C-ähnliche Syntax. Für vollständige Details wie man sie konfiguriert, schauen Sie sich bitte die arcload Seite im Linux/MIPS Wiki an. Die Kurzfassung: Definieren Sie mit Hilfe der `OSLoadFilename` Variable eine Reihe von Optionen die beim Booten aktiviert und deaktiviert werden.

DATEI **`arc.cf`** **Beispiel arc.cf**

```
# ARCLoad Konfiguration

# Einige Standardeinstellungen...
append  "root=/dev/sda5";
append  "ro";
append  "console=ttyS0,9600";

# Die Hauptdefinition. ip28 kann geändert werden, wenn Sie dies wünschen
ip28 {
        # Definition für einen "funktionierenden" Kernel
        # Wählen Sie dies durch die Einstellung von OSLoadFilename="ip28(working)"
        working {
                description     "SGI Indigo2 Impact R10000\n\r";
                image system    "/working";
        }

        # Definition für einen "neuen" Kernel
        # Wählen Sie dies durch die Einstellung von OSLoadFilename="ip28(new)"
        new {
                description     "SGI Indigo2 Impact R10000 - Testing Kernel\n\r";
                image system    "/new";
        }

        # Zur Fehlersuche in einem Kernel
        # Wählen Sie dies durch die Einstellung von OSLoadFilename="ip28(working,debug)"
        # oder OSLoadFilename="ip28(new,debug)"
        debug {
                description     "Debug console";
                append          "init=/bin/bash";
        }
}

```

Beginnend mit arcload-0.5 können arc.cf und Kernel entweder im Volume-Header oder auf einer Partition gespeichert sein. Um diese neue Eigenschaft zu nutzen, können Sie die Dateien in der /boot/ Partition ablegen (oder / wenn es keine extra Boot Partition gibt). arcload verwendet den Dateisystemtreiber-Code des beliebten grub Bootloader und unterstützt somit den gleichen Umfang an Dateisystemen.

`root #` `dvhtool --unix-to-vh arc.cf arc.cf
`

`root #` `dvhtool --unix-to-vh /usr/src/linux/vmlinux new`

## CoLo für Cobalt MicroServer

### CoLo installieren

Cobalt Server haben eine viel weniger leistungsfähige Firmware auf dem Chip installiert. Das Cobalt BOOTROM ist im Vergleich zum SGI PROM primitiv und weist eine Reihe erheblicher Limitierungen auf.

- Es besteht eine (ungefähr) 675 KB Größenlimitierung des Kernels. Die aktuelle Größe von Linux 2.4 macht es nahezu unmöglich einen Kernel dieser Größe zu erzeugen. Linux 2.6 und 3.x stehen vollkommen außer Frage.
- 64-Bit Kernel werden von der Original-Firmware nicht unterstützt (obwohl diese momentan sehr experimentell auf Cobalt Maschinen sind)
- Die Shell ist im günstigsten Fall rudimentär

Um diese Limitierungen zu überwinden wurde eine alternative Firmware genannt CoLo (Cobalt Loader) entwickelt. Dies ist ein BOOTROM Abbild das entweder auf den Chip im Cobalt Server geflasht, oder von der existierenden Firmware geladen werden kann.

**Hinweis**

Dieses Handbuch wird durch die Einrichtung von CoLo führen, damit es durch die Original-Firmware geladen wird. Dies ist der einzig wirklich sichere und empfohlene Weg um CoLo einzurichten.

**Warnung**

Wenn Sie möchten, können diese in den Server geflasht werden, um die Original-Firmware vollständig zu ersetzen - allerdings sind Sie bei diesem Unterfangen völlig auf sich allein gestellt. Sollte etwas schief gehen, entfernen Sie das BOOTROM und programmieren Sie es mit der Standard-Firmware neu. Wenn sich das beängstigend anhört - dann flashen Sie die Maschine NICHT. Gentoo übernimmt keine Verantwortung für das, was passiert, wenn dieser Ratschlag ignoriert wird.

Nun zur Installation von CoLo. Beginnen Sie damit, das Paket zu emergen:

`root #` `emerge --ask sys-boot/colo`

Nach der Installation werfen Sie einen Blick in das Verzeichnis /usr/lib/colo/. Hier liegen die folgenden zwei Dateien:

- colo-chain.elf (der "Kernel" der von der Original-Firmware geladen wird)
- colo-rom-image.bin (ein ROM Abbild zum Flashen auf das BOOTROM)

- colo-chain.elf \- the "kernel" for the stock firmware to load.
- colo-rom-image.bin \- a ROM image for flashing into the BOOTROM.

Beginnen Sie mit dem Einhängen von /boot/ und dem Ablegen einer komprimierten Kopie von colo-chain.elf in /boot/, wo das System es erwartet.

`root #` `gzip -9vc /usr/lib/colo/colo-chain.elf > /boot/vmlinux.gz`

### CoLo konfigurieren

Wenn das System bootet wird es CoLo laden, welches ein Menü auf dem LCD ausspuckt. Die erste Option (und die Voreinstellung, die nach ca. 5 Sekunden übernommen wird) ist das Booten von der Festplatte. Das System versucht dann die erste Linux Partition die es findet zu mounten und anschließend das Script default.colo zu starten. Die Syntax ist vollständig in der CoLo Dokumentation beschrieben (werfen Sie einen Blick auf /usr/share/doc/colo-X.YY/README.shell.gz \-\- wobei X.YY die installierte Versionsnummer ist) und sehr einfach.

**Hinweis**

Nur ein Tipp: Bei der Installation eines Kernels ist es empfehlenswert zwei Kernel-Abbilder zu erzeugen: kernel.gz.working \-\- ein bekanntermaßen funktionierender Kernel und kernel.gz.new \-\- der Kernel der gerade kompiliert wurde. Es ist möglich symbolische Links zu nutzen, um auf die aktuellen "...new" und "...working" Kernel zu verweisen. Oder benennen Sie die Kernel-Abbilder einfach entsprechend um.

DATEI **`default.colo`** **CoLo Beispielkonfiguration**

```
#:CoLo:#
mount hda1
load /kernel.gz.working
execute root=/dev/sda5 ro console=ttyS0,115200

```

**Hinweis**

CoLo verweigert das Laden eines Skriptes, das nicht mit der Zeile `#:CoLo:#` beginnt. Sie können es als Äquivalent von `#!/bin/sh` in Shell Skripten ansehen.

Es ist ebenfalls möglich eine Farge nach dem zu bootenden Kernel und der Konfiguration mit einem Standard-Timeout zu stellen. Die nachfolgende Konfiguration tut genau dies: Die fragt den Benutzer welchen Kernel er verwenden möchte und führt das ausgewählte Abbild aus. vmlinux.gz.new und vmlinux.gz.working können entweder die eigentlichen Kernel-Abbilder, oder lediglich symbolische Links die auf Kernel-Abbilder auf dieser Festplatte verweisen sein. Das Argument `50` von `select` gibt an, dass mit der ersten Option ("Working") nach 50/10 Sekunden fortgefahren werden soll.

DATEI **`default.colo`** **Menübasierte Konfiguration**

```
#:CoLo:#
lcd "Mounting hda1"
mount hda1
select "Which Kernel?" 50 Working New

goto {menu-option}
var image-name vmlinux.gz.working
goto 3f
@var image-name vmlinux.gz.working
goto 2f
@var image-name vmlinux.gz.new

@lcd "Loading Linux" {image-name}
load /{image-name}
lcd "Booting..."
execute root=/dev/sda5 ro console=ttyS0,115200
boot

```

Für mehr Informationen sehen Sie sich bitte die Dokumentation unter /usr/share/doc/colo-VERSION an.

## Serielle Konsole einrichten

Die Linux Installation so wie sie jetzt ist würde booten aber voraussetzen, dass der Benutzer an einem physischen Terminal eingeloggt ist. Auf Cobalt Maschinen ist das besonders schlecht -- es gibt so etwas wie ein physisches Terminal nicht.

**Hinweis**

Diejenigen die den Luxus eines unterstützten Video-Chipsatz haben, können wenn sie möchten diesen Abschnitt überspringen.

Öffnen Sie als erstes die Datei /etc/inittab mit einem Editor. Etwas weiter unten in der Datei betrachten Sie folgendes:

DATEI **`/etc/inittab`** **inittab Ausschnitt**

```
# SERIAL CONSOLE
#c0:12345:respawn:/sbin/agetty 9600 ttyS0 vt102

# TERMINALS
c1:12345:respawn:/sbin/agetty 38400 tty1 linux
c2:12345:respawn:/sbin/agetty 38400 tty2 linux
c3:12345:respawn:/sbin/agetty 38400 tty3 linux
c4:12345:respawn:/sbin/agetty 38400 tty4 linux
c5:12345:respawn:/sbin/agetty 38400 tty5 linux
c6:12345:respawn:/sbin/agetty 38400 tty6 linux

# What to do at the "Three Finger Salute".
ca:12345:ctrlaltdel:/sbin/shutdown -r now

```

Entkommentieren Sie zunächst die c0 Zeile. Standardmäßig ist sie auf eine Terminal-Baudrate von 9600 bps eingestellt. Auf Cobalt-Servern kann dies auf 115200 geändert werden, um der vom BOOTROM festgelegten Baudrate zu entsprechen. Dieser Abschnitt sieht dann folgendermaßen aus. Auf einer Headless-Maschine (z.B. Cobalt-Servern) ist es auch empfehlenswert die lokalen Terminalzeilen (c1 bis c6) auszukommentieren, da diese die Angewohnheit haben, sich falsch zu verhalten, wenn sie /dev/ttyX nicht öffnen können.

DATEI **`/etc/inittab`** **inittab Beispielausschnitt**

```
# SERIAL CONSOLE
c0:12345:respawn:/sbin/agetty 115200 ttyS0 vt102

# TERMINALS -- These are useless on a headless qube
#c1:12345:respawn:/sbin/agetty 38400 tty1 linux
#c2:12345:respawn:/sbin/agetty 38400 tty2 linux
#c3:12345:respawn:/sbin/agetty 38400 tty3 linux
#c4:12345:respawn:/sbin/agetty 38400 tty4 linux
#c5:12345:respawn:/sbin/agetty 38400 tty5 linux
#c6:12345:respawn:/sbin/agetty 38400 tty6 linux

```

Als letztes... teilen Sie dem System mit, dass der lokalen seriellen Schnittstelle als sicheres Terminal vertraut werden kann. Die zu ändernde Datei ist /etc/securetty. Sie enthält eine Liste von Terminals, denen das System vertraut. Fügen Sie einfach zwei weitere Zeilen ein, die es erlauben, die serielle Leitung für Root-Anmeldungen zu verwenden.

`root #` `echo 'ttyS0' >> /etc/securetty`

Neuerdings nennt Linux dies auch /dev/tts/0 -- fügen Sie also auch dies hinzu:

`root #` `echo 'tts/0' >> /etc/securetty`

## SGI PROM Anpassung

### Allgemeine PROM Einstellungen

Wenn der Bootloader installiert ist, gehen Sie nach dem Neustart (der in Kürze erfolgen wird) zum Systemwartungsmenü und wählen Sie `Enter` auf _Command Monitor (5)_, ähnlich wie beim ersten Neustart des Systems.

CODE **Menü nach dem Booten**

```
1&#41; Start System
2&#41; Install System Software
3&#41; Run Diagnostics
4&#41; Recover System
5&#41; Enter Command Monitor

```

Geben Sie den Speicherort des Volume-Header an:

`>>` `setenv SystemPartition scsi(0)disk(1)rdisk(0)partition(8)`

Automatisch Gentoo booten:

`>>` `setenv AutoLoad Yes`

Die Zeitzone setzen:

`>>` `setenv TimeZone EST5EDT`

Verwenden Sie die serielle Konsole - Benutzer von Grafikkarten sollten "g" anstelle von "d1" (eins) angeben:

`>>` `setenv console d1`

Stellen Sie die Baudrate der seriellen Konsole ein. Dies ist optional. Die Standardeinstellung ist 9600, wenngleich man Datenraten bis zu 38400 verwenden kann, falls gewünscht:

`>>` `setenv dbaud 9600`

Die nächsten Einstellungen hängen davon ab, wie das System gebootet wird.

### Einstellungen für direktes booten vom Volume-Header

**Hinweis**

Dies wird hier der Vollständigkeit halber abgedeckt. Wir empfehlen stattdessen, dass sich der Benutzer die Installation von arcload ansieht.

**Hinweis**

Dies funktioniert nur auf Indy, Indigo2 (R4k) und Challenge S.

Ersetzen Sie `<root device>` durch die Root Partition von Gentoo, wie beispielsweise /dev/sda3:

`>>` `setenv OSLoadPartition <root device>`

Um die verfügbaren Kernel aufzulisten geben sie "ls" ein.

`>>` `setenv OSLoader <kernel name>
`

`>>` `setenv OSLoadFilename <kernel name>`

Legen die zu übergebenden Kernel-Parameter fest:

`>>` `setenv OSLoadOptions <kernel parameters>`

Um einen Kernel zu versuchen ohne mit den Kernel-Parametern herumzuhantieren, verwenden Sie den `boot -f` PROM Befehl:

`root #` `boot -f new root=/dev/sda5 ro`

### arcload Einstellungen

arcload verwendet die Option `OSLoadFilename` um festzulegen welche Optionen von arc.cf eingestellt werden. Die Konfigurationsdatei ist im Grunde ein Skript, das Boot-Abbilder auf der obersten Ebene für verschiedene Systeme definiert und innerhalb von diesen optionale Einstellungen. Somit zieht OSLoadFilename=mysys(serial) die Einstellungen für den mysys block herein und stellt weitere in serial überschriebene Optionen ein.

In der obigen Beispieldatei ist ein Systemblock definiert, ip28 mit den Optionen working, new und debug. Als nächstes werden die PROM-Variablen definiert:

Wählen Sie arcload als Bootloader:- `sash64` oder `sashARCS`:

`>>` `setenv OSLoader sash64`

Verwenden Sie das "working" Kernel-Abbild, das im "ip28" Abschnitt von arc.cf definiert ist:

`>>` `setenv OSLoadFilename ip28(working)`

Beginnend mit arcload-0.5 müssen Dateien nicht länger im Volume-Header plaziert sein -- sie können statt dessen in einer Partition liegen. Um arcload mitzuteilen wo es nach seiner Konfigurationsdatei und Kernels suchen soll, müssen Sie die PROM Variable `OSLoadPartition` setzen. Der hier genau anzugebende Wert hängt davon ab, wo die Festplatte auf dem SCSI Bus liegt. Verwenden Sie die PROM Variable `SystemPartition` als Vorlage -- nur die Partitionsnummer sollte geändert werden müssen.

**Hinweis**

Partitionen sind durchnummeriert beginnend mit 0, nicht 1 wie bei Linux.

Um vom Volume-Header zu laden, verwenden Sie Partition 8:

`>>` `setenv OSLoadPartition scsi(0)disk(1)rdisk(0)partition(8)`

Andernfalls geben Sie die Partition und den Dateisystemtyp an:

`>>` `setenv OSLoadPartition scsi(0)disk(1)rdisk(0)partition(0)[ext2]`

## Neustart des Systems

Verlassen Sie die chroot-Umgebung und hängen Sie alle gemounteten Partitionen aus. Geben Sie dann den magischen Befehl ein, der den alles entscheidenden Test einleitet - reboot.

`root #` `exit`

`livecd~#` `cd
`

`livecd~#` `umount -l /mnt/gentoo/dev{/shm,/pts,}
`

`livecd~#` `umount -R /mnt/gentoo
`

`livecd~#` `reboot`

Vergessen Sie nicht, das Installations-Medium zu entfernen. Andernfalls könnte erneut das Installations-Medium anstelle des neuen Gentoo Systems gebootet werden.

Nach dem Neustart in die neu installierte Gentoo Umgebung können Sie Ihre Installation mit dem Kapitel [Abschluss der Gentoo Installation](/wiki/Handbook:MIPS/Installation/Finalizing/de "Handbook:MIPS/Installation/Finalizing/de") fertigstellen.

[← Installation der Tools](/wiki/Handbook:MIPS/Installation/Tools/de "Previous part") [Anfang](/wiki/Handbook:MIPS/de "Handbook:MIPS/de") [Abschluss der Installation →](/wiki/Handbook:MIPS/Installation/Finalizing/de "Next part")