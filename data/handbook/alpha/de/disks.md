# Disks

Other languages:

- Deutsch
- [English](/wiki/Handbook:Alpha/Installation/Disks "Handbook:Alpha/Installation/Disks (100% translated)")
- [Türkçe](/wiki/Handbook:Alpha/Installation/Disks/tr "Handbook:Alpha/Installation/Disks/tr (0% translated)")
- [español](/wiki/Handbook:Alpha/Installation/Disks/es "Manual de Gentoo: Alpha/Instalación/Discos (100% translated)")
- [français](/wiki/Handbook:Alpha/Installation/Disks/fr "Manuel:Alpha/Installation/Disques (100% translated)")
- [italiano](/wiki/Handbook:Alpha/Installation/Disks/it "Handbook:Alpha/Installation/Disks/it (100% translated)")
- [magyar](/wiki/Handbook:Alpha/Installation/Disks/hu "Handbook:Alpha/Installation/Disks/hu (100% translated)")
- [polski](/wiki/Handbook:Alpha/Installation/Disks/pl "Handbook:Alpha/Installation/Disks (100% translated)")
- [português do Brasil](/wiki/Handbook:Alpha/Installation/Disks/pt-br "Manual:Alpha/Instalação/Discos (100% translated)")
- [čeština](/wiki/Handbook:Alpha/Installation/Disks/cs "Handbook:Alpha/Installation/Disks/cs (50% translated)")
- [русский](/wiki/Handbook:Alpha/Installation/Disks/ru "Handbook:Alpha/Installation/Disks (100% translated)")
- [தமிழ்](/wiki/Handbook:Alpha/Installation/Disks/ta "கையேடு:Alpha/நிறுவல்/வட்டுக்கள் (100% translated)")
- [中文（中国大陆）‎](/wiki/Handbook:Alpha/Installation/Disks/zh-cn "手册：Alpha/安装/准备磁盘 (100% translated)")
- [日本語](/wiki/Handbook:Alpha/Installation/Disks/ja "ハンドブック:Alpha/インストール/ディスク (100% translated)")
- [한국어](/wiki/Handbook:Alpha/Installation/Disks/ko "Handbook:Alpha/Installation/Disks/ko (100% translated)")

[Alpha Handbuch](/wiki/Handbook:Alpha/de "Handbook:Alpha/de")[Installation](/wiki/Handbook:Alpha/Full/Installation/de "Handbook:Alpha/Full/Installation/de")[Über die Installation](/wiki/Handbook:Alpha/Installation/About/de "Handbook:Alpha/Installation/About/de")[Auswahl des Mediums](/wiki/Handbook:Alpha/Installation/Media/de "Handbook:Alpha/Installation/Media/de")[Konfiguration des Netzwerks](/wiki/Handbook:Alpha/Installation/Networking/de "Handbook:Alpha/Installation/Networking/de")Vorbereiten der Festplatte(n)[Installation des Stage Archivs](/wiki/Handbook:Alpha/Installation/Stage/de "Handbook:Alpha/Installation/Stage/de")[Installation des Basissystems](/wiki/Handbook:Alpha/Installation/Base/de "Handbook:Alpha/Installation/Base/de")[Konfiguration des Kernels](/wiki/Handbook:Alpha/Installation/Kernel/de "Handbook:Alpha/Installation/Kernel/de")[Konfiguration des Systems](/wiki/Handbook:Alpha/Installation/System/de "Handbook:Alpha/Installation/System/de")[Installation der Tools](/wiki/Handbook:Alpha/Installation/Tools/de "Handbook:Alpha/Installation/Tools/de")[Konfiguration des Bootloaders](/wiki/Handbook:Alpha/Installation/Bootloader/de "Handbook:Alpha/Installation/Bootloader/de")[Abschluss](/wiki/Handbook:Alpha/Installation/Finalizing/de "Handbook:Alpha/Installation/Finalizing/de")[Arbeiten mit Gentoo](/wiki/Handbook:Alpha/Full/Working/de "Handbook:Alpha/Full/Working/de")[Portage-Einführung](/wiki/Handbook:Alpha/Working/Portage/de "Handbook:Alpha/Working/Portage/de")[USE-Flags](/wiki/Handbook:Alpha/Working/USE/de "Handbook:Alpha/Working/USE/de")[Portage-Features](/wiki/Handbook:Alpha/Working/Features/de "Handbook:Alpha/Working/Features/de")[Initskript-System](/wiki/Handbook:Alpha/Working/Initscripts/de "Handbook:Alpha/Working/Initscripts/de")[Umgebungsvariablen](/wiki/Handbook:Alpha/Working/EnvVar/de "Handbook:Alpha/Working/EnvVar/de")[Arbeiten mit Portage](/wiki/Handbook:Alpha/Full/Portage/de "Handbook:Alpha/Full/Portage/de")[Dateien und Verzeichnisse](/wiki/Handbook:Alpha/Portage/Files/de "Handbook:Alpha/Portage/Files/de")[Variablen](/wiki/Handbook:Alpha/Portage/Variables/de "Handbook:Alpha/Portage/Variables/de")[Mischen von Softwarezweigen](/wiki/Handbook:Alpha/Portage/Branches/de "Handbook:Alpha/Portage/Branches/de")[Zusätzliche Tools](/wiki/Handbook:Alpha/Portage/Tools/de "Handbook:Alpha/Portage/Tools/de")[Eigener Portage-Tree](/wiki/Handbook:Alpha/Portage/CustomTree/de "Handbook:Alpha/Portage/CustomTree/de")[Erweiterte Portage-Features](/wiki/Handbook:Alpha/Portage/Advanced/de "Handbook:Alpha/Portage/Advanced/de")[Netzwerk-Konfiguration](/wiki/Handbook:Alpha/Full/Networking/de "Handbook:Alpha/Full/Networking/de")[Zu Beginn](/wiki/Handbook:Alpha/Networking/Introduction/de "Handbook:Alpha/Networking/Introduction/de")[Fortgeschrittene Konfiguration](/wiki/Handbook:Alpha/Networking/Advanced/de "Handbook:Alpha/Networking/Advanced/de")[Modulare Vernetzung](/wiki/Handbook:Alpha/Networking/Modular/de "Handbook:Alpha/Networking/Modular/de")[Drahtlose Netzwerke](/wiki/Handbook:Alpha/Networking/Wireless/de "Handbook:Alpha/Networking/Wireless/de")[Funktionalität hinzufügen](/wiki/Handbook:Alpha/Networking/Extending/de "Handbook:Alpha/Networking/Extending/de")[Dynamisches Management](/wiki/Handbook:Alpha/Networking/Dynamic/de "Handbook:Alpha/Networking/Dynamic/de")

## Contents

- [1Einführung in blockorientierte Geräte](#Einf.C3.BChrung_in_blockorientierte_Ger.C3.A4te)
  - [1.1Blockorientierte Geräte](#Blockorientierte_Ger.C3.A4te)
  - [1.2Slices](#Slices)
- [2Ein Partitionsschema entwerfen](#Ein_Partitionsschema_entwerfen)
  - [2.1Wie viele Partitionen und wie groß?](#Wie_viele_Partitionen_und_wie_gro.C3.9F.3F)
  - [2.2Was ist mit dem Swap-Speicher?](#Was_ist_mit_dem_Swap-Speicher.3F)
- [3Festplatte mit fdisk partitionieren (nur SRM)](#Festplatte_mit_fdisk_partitionieren_.28nur_SRM.29)
  - [3.1Verfügbare Festplatten identifizieren](#Verf.C3.BCgbare_Festplatten_identifizieren)
  - [3.2Creating a BSD disklabel](#Creating_a_BSD_disklabel)
  - [3.3Löschen aller Slices](#L.C3.B6schen_aller_Slices)
  - [3.4Swap Slice erstellen](#Swap_Slice_erstellen)
  - [3.5Creating the boot slice](#Creating_the_boot_slice)
  - [3.6Root Slice erstellen](#Root_Slice_erstellen)
  - [3.7Slice Layout speichern und fdisk beenden](#Slice_Layout_speichern_und_fdisk_beenden)
- [4Festplatte mit fdisk partitionieren (nur ARC/AlphaBIOS)](#Festplatte_mit_fdisk_partitionieren_.28nur_ARC.2FAlphaBIOS.29)
  - [4.1Verfügbare Festplatten identifizieren](#Verf.C3.BCgbare_Festplatten_identifizieren_2)
  - [4.2Löschen aller Partitionen](#L.C3.B6schen_aller_Partitionen)
  - [4.3Boot Partition erstellen](#Boot_Partition_erstellen)
  - [4.4Swap Partition erstellen](#Swap_Partition_erstellen)
  - [4.5Root Partition erstellen](#Root_Partition_erstellen)
  - [4.6Partitionslayout Speichern und fdisk beenden](#Partitionslayout_Speichern_und_fdisk_beenden)
- [5Erstellen von Dateisystemen](#Erstellen_von_Dateisystemen)
  - [5.1Einleitung](#Einleitung)
  - [5.2Dateisysteme](#Dateisysteme)
  - [5.3Dateisystem auf einer Partition anlegen](#Dateisystem_auf_einer_Partition_anlegen)
    - [5.3.1EFI system partition filesystem](#EFI_system_partition_filesystem)
    - [5.3.2Legacy BIOS boot partition filesystem](#Legacy_BIOS_boot_partition_filesystem)
    - [5.3.3Small ext4 partitions](#Small_ext4_partitions)
  - [5.4Aktivieren der Swap-Partition](#Aktivieren_der_Swap-Partition)
- [6Einhängen der Root-Partition](#Einh.C3.A4ngen_der_Root-Partition)

## Einführung in blockorientierte Geräte

### Blockorientierte Geräte

Schauen wir uns die Festplatten-spezifischen Aspekte von Gentoo Linux und Linux im Allgemeinen an - insbesondere blockorientierte Geräte (Block Devices), Partitionen und Linux Dateisysteme. Wenn Sie die Vor- und Nachteile von Festplatten verstanden haben, können Sie Partitionen und Dateisysteme für die Installation erstellen.

Zu Beginn schauen wir uns blockorientierte Geräte an. SCSI- und SATA-Laufwerke haben Device-Namen wie: /dev/sda, /dev/sdb, /dev/sdc usw. Modernere Rechner können PCI-Express basierte NVMe Solid-State-Disks haben, die Device-Namen haben wie: /dev/nvme0n1, /dev/nvme0n2 usw.

Die folgende Tabelle soll Lesern dabei helfen herauszufinden, wo bestimmte Arten von blockorientierten Geräten zu finden sind:

Device-TypStandard Device-NameAnmerkungen
IDE, SATA, SAS, SCSI, or USB flash/dev/sdaDiese Device-Typen werden auf Hardware ab 2007 verwendet - und sind vermutlich die am häufigsten genutzten Device-Namen unter Linux. Diese Geräte werden als blockorientierter Speicher angeschlossen über den [SATA bus](https://en.wikipedia.org/wiki/Serial_ATA "wikipedia:Serial ATA"), über [SCSI](https://en.wikipedia.org/wiki/SCSI "wikipedia:SCSI") und über [USB](https://en.wikipedia.org/wiki/USB "wikipedia:USB"). Beispielsweise wird die erste Partition des ersten SATA-Devices /dev/sda1 genannt.
NVM Express (NVMe)/dev/nvme0n1The latest in solid state technology, [NVMe](https://en.wikipedia.org/wiki/NVM_Express "wikipedia:NVM Express") drives are connected to the PCI Express bus and have the fastest transfer block speeds on the market. Systems from around 2014 and newer may have support for NVMe hardware. The first partition on the first NVMe device is called /dev/nvme0n1p1.
MMC, eMMC, and SD/dev/mmcblk0[embedded MMC](https://en.wikipedia.org/wiki/MultiMediaCard#eMMC "wikipedia:MultiMediaCard") devices, SD cards, and other types of memory cards can be useful for data storage. That said, many systems may not permit booting from these types of devices. It is suggested to **not** use these devices for active Linux installations; rather consider using them to transfer files, which is their design goal. Alternatively they could be useful for short-term backups.

Die oben genannten blockorientierten Geräte repräsentieren eine abstrakte Schnittstelle zur Festplatte. Benutzerprogramme können diese Block Devices nutzen, um mit der Festplatte zu interagieren, ohne sich darum sorgen zu müssen, ob die Festplatten über SATA, SCSI oder etwas anderem angebunden sind. Das Programm kann den Speicher auf der Festplatte einfach als eine Anhäufung zusammenhängender 4096-Byte (4k) Blöcke mit wahlfreiem Zugriff ansprechen.

### Slices

Obwohl es theoretisch möglich wäre eine vollständige Festplatte zu nutzen um ein Linux-System unterzubringen, kommt das in der Praxis fast nie vor. Stattdessen werden komplette Festplatten Block Devices in kleinere, besser handhabbare Block Devices unterteilt. Auf Alpha-Systemen werden diese _Slices_ genannt.

**Hinweis**

In den folgenden Sektionen verwenden die Anweisungen zur Installation die Beispiel-Partitionierung des ARC/AlphaBIOS Setup. Bitte passen Sie diese Ihren persönlichen Vorstellungen an!

## Ein Partitionsschema entwerfen

### Wie viele Partitionen und wie groß?

Bei dem Design des Partitionsschemas sollten die Anforderungen an das System und an die Dateisysteme berücksichtigt werden. Wenn es viele Nutzer gibt, ist eine eigene Partition /home/ ratsam, da diese die Sicherheit erhöht und Backups und andere Wartungsarbeiten vereinfacht. Wenn Gentoo installiert wird, um als Mailserver zu dienen, dann sollte es eine eigene Partition /var/ geben, weil alle Mails im Verzeichnis /var/ gespeichert werden. Spiele-Server werden eine eigene Partition /opt/ besitzen, da die meiste Spiele-Server-Software dort installiert wird. Der Grund für diese Empfehlungen ist ähnlich wie für das /home/ Verzeichnis: Sicherheit, Backups und Wartung.

Bei den meisten Gentoo-Installationen sollten /usr/ und /var/ relativ groß sein. In /usr werden die Mehrzahl der Anwendungen und auch der Linux Kernel Quellcode gespeichert (unter /usr/src). Standardmäßig enthält /var/ das Gentoo ebuild Repository (unter /var/db/repos/gentoo), das alleine schon rund 650 MiB Plattenplatz benötigt. Diese Größenabschätzung enthält noch nicht den benötigten Plattenplatz für die Verzeichnisse /var/cache/distfiles und /var/cache/binpkgs, die sich im Laufe der Zeit mit Source-Code Dateien und (optional) mit Binärpaketen füllen werden - je nachdem, wann und wie sie dem System hinzugefügt werden.

Die Anzahl und Größe der Partitionen hängt vom Abwägen der Vor- und Nachteile und der Auswahl der besten Lösung für einen gegebenen Anwendungsfall ab. Separate Partitionen oder Volumes haben folgende Vorteile:

- Sie können das performanteste Dateisystem für jede Partition oder jedes Volume wählen.
- Dem Gesamtsystem kann der freie Speicherplatz nicht ausgehen, wenn ein fehlerhaftes Tool kontinuierlich Dateien auf eine Partition oder ein Volume schreibt.
- Falls nötig, kann die Zeit für Dateisystemüberprüfungen reduziert werden, da mehrere Überprüfungen gleichzeitig durchgeführt werden können. (Dieser Vorteil kommt aber eher bei mehreren Festplatten, als bei mehreren Partitionen auf einer Festplatte zum Tragen.)
- Sie können die Sicherheit erhöhen, indem Sie einige Partitionen oder Volumes "read-only", `nosuid` (setuid Flags werden ignoriert), `noexec` (executable Flags werden ignoriert) etc. einbinden.

Viele separate Partitionen können aber auch Nachteile haben:

- Wenn diese schlecht an das System angepasst sind, kann es sein, dass eine Partition voll ist und auf einer anderen Partition noch viel freier Platz verfügbar ist.
- Eine separate Partition für /usr/ kann es erforderlich machen, dass beim Booten ein initramfs verwendet wird, welches diese Partitionen vor der Ausführung anderer Boot-Skripte mountet. Das Erzeugen und Betreiben eines initramsfs ist nicht Teil dieses Handbuchs. **Wir empfehlen Anfängern, für /usr/ keine eigene Partition zu verwenden.**
- Es gibt ein Limit von maximal 15 Partitionen für SCSI und SATA - es sei denn, der Datenträger nutzt GPT-Labels.

**Hinweis**

Installationen, die systemd als Dienst-und Init-System verwenden wollen, müssen /usr/ beim Booten verfügbar haben, entweder als Teil des Root-Dateisystems oder eingehängt über ein initramfs.

### Was ist mit dem Swap-Speicher?

Recommendations for swap space size
RAM sizeSuspend support?Hibernation support?
2 GB or less4GB4GB
2 to 8 GBRAM amount2 \* RAM
8 to 64 GB8 GB minimum, 16 maximum1.5 \* RAM
64 GB or greater8 GB minimumHibernation not recommended! Hibernation is _not_ recommended for systems with very large amounts of memory. While possible, the entire contents of memory must be written to disk in order to successfully hibernate. Writing tens of gigabytes (or worse!) out to disk can can take a considerable amount of time, especially when rotational disks are used. It is best to suspend in this scenario.

Es gibt keine perfekte Größe für den Swap-Speicher. Der Zweck von Swap-Speicher ist, Festplattenspeicherplatz für den Kernel bereitzuhalten, wenn der interne Speicher (RAM) knapp wird. Der Swap-Speicher erlaubt dem Kernel, Speicherseiten, auf die vermutlich nicht bald zugegriffen wird, auf die Platte auszulagern (Swap oder Page-Out). Dadurch kann Arbeitsspeicher im RAM für den aktuell laufenden Prozess freigemacht werden. Werden die auf die Festplatte ausgelagerten Speicherseiten (Pages) jedoch plötzlich benötigt, müssen diese Seiten wieder zurück in den Arbeitsspeicher geladen werden (Page-In). Dies dauert jedoch erheblich länger, als wenn die Daten direkt aus dem RAM gelesen werden könnten (da Festplatten verglichen mit Arbeitsspeicher sehr langsam sind).

Wenn auf einem System keine speicherintensiven Anwendungen ausgeführt werden oder das System viel RAM zur Verfügung hat, benötigt es vermutlich nicht viel Swap-Speicher. Wenn jedoch der Ruhezustand "Hibernation" verwendet werden soll, wird der Swap-Speicher verwendet, um den _gesamten Inhalt des Hauptspeichers (RAM)_ zu sichern (dieser Ruhezustand wird bei Desktop- und Laptop-Systemen häufiger verwendet, als bei Servern). Wenn das System den Ruhezustand "Hibernation" unterstützen soll, muss der Swap-Speicher so groß wie oder größer als der Hauptspeicher (RAM) sein.

Als generelle Regel gilt: der Swap-Speicher sollte zwei Mal so groß sein wie der Arbeitsspeicher (RAM). Auf Systemen mit mehreren (rotierenden) Festplatten ist es sinnvoll, eine Swap-Partition auf jeder Festplatte einzurichten, damit Schreib-/Lese-Operationen parallel ausgeführt werden können. Je schneller auf einen Festplatte zugegriffen werden kann, desto schneller wird das System arbeiten, wenn auf Swap-Speicher zugegriffen werden muss. Wenn zwischen rotierenden Festplatten und SSDs gewählt werden kann, ist es aus Performance-Sicht besser, den Swap-Speicher auf die SSD zu legen. Alternativ zu Swap-Partitionen können auch Swap-Dateien verwendet werden; dies ist hauptsächlich interessant bei Systemen mit sehr geringem Festplatten-Platz.

## Festplatte mit fdisk partitionieren (nur SRM)

Die folgenden Abschnitte erklären, wie das Slice-Layout Beispiel für SRM zu erstellen ist:

Slice
Beschreibung
/dev/sda1Swap Slice
/dev/sda2Root Slice
/dev/sda3Gesamte Festplatte (benötigt)

Ändern Sie das Slice Layout Ihren Vorstellungen entsprechend ab.

### Verfügbare Festplatten identifizieren

Um herauszufinden welche Festplatten im System laufen, verwenden sie die folgenden Befehle:

Für IDE Festplatten:

`root #` `dmesg | grep 'drive$'`

Für SCSI Festplatten:

`root #` `dmesg | grep 'scsi'`

Die Ausgabe zeigt an, welche Festplatten erkannt wurden und deren jeweiligen /dev/ Eintrag. In den folgenden Abschnitten gehen wir davon aus, dass es sich um eine SCSI Festplatte auf /dev/sda handelt.

### Creating a BSD disklabel

If the hard drive is completely blank, then first create a BSD disklabel. On alpha, this can't be done using fdisk, so using parted is the way forward here.

Now fire up parted:

`root #` `parted /dev/sda `

```
Using /dev/sda
Welcome to GNU Parted! Type 'help' to view a list of commands.

```

`(parted)` `mklabel bsd`

```
Warning: The existing disk label on /dev/sda will be destroyed and all data on this disk will be lost. Do you want to continue?
Yes/No? yes
(parted) quit
Information: You may need to update /etc/fstab.

```

Now that we have a bsd disklabel on our drive, continue creating slices. This can be done using parted or, as in the examples below, using fdisk:

`root #` `fdisk /dev/sda`

### Löschen aller Slices

Wenn die Festplatte komplett leer ist, dann erstellen Sie zunächst ein BDS Disklabel.

`Command (m for help):` `b`

```
/dev/sda contains no disklabel.
Do you want to create a disklabel? (y/n) y
A bunch of drive-specific info will show here
3 partitions:
#       start       end      size     fstype   [fsize bsize   cpg]
  c:        1      5290*     5289*    unused        0     0

```

Wir beginnen mit dem Löschen aller Slices mit Ausnahme der 'c'-Slice (eine Anforderung bei der Nutzung von BSD Disklabels). Im Nachfolgenden zeigen wir, wie eine Slice gelöscht wird (im Beispiel verwenden wir 'a'). Wiederholen Sie den Vorgang um alle anderen Slices zu löschen (wieder mit Ausnahme der 'c'-Slice).

Verwenden Sie `p` um alle existierenden Slices anzuzeigen. `d` wird zum Löschen einer Slice betätigt.

`BSD disklabel command (m for help):` `p`

```
8 partitions:
#       start       end      size     fstype   [fsize bsize   cpg]
  a:        1       235*      234*    4.2BSD     1024  8192    16
  b:      235*      469*      234*      swap
  c:        1      5290*     5289*    unused        0     0
  d:      469*     2076*     1607*    unused        0     0
  e:     2076*     3683*     1607*    unused        0     0
  f:     3683*     5290*     1607*    unused        0     0
  g:      469*     1749*     1280     4.2BSD     1024  8192    16
  h:     1749*     5290*     3541*    unused        0     0

```

`BSD disklabel command (m for help):` `d`

```
Partition (a-h): a

```

Nachdem Sie diesen Vorgang für alle Slices wiederholt haben, sollte die Auflistung in etwa wie folgt aussehen:

`BSD disklabel command (m for help):` `p`

```
3 partitions:
#       start       end      size     fstype   [fsize bsize   cpg]
  c:        1      5290*     5289*    unused        0     0

```

### Swap Slice erstellen

Auf Alpha basierten Systemen gibt es keine Notwendigkeit für eine separate Boot Slice. Wie auch immer, der erste Zylinder darf nicht verwendet werden, weil das "aboot" Abbild dort abgelegt wird.

Wir werden eine Swap Slice beginnend beim dritten Zylinder, mit einer Gesamtgröße von 1 GB erstellen. Verwenden Sie `n` um eine neue Slice anzulegen. Nach der Erzeugung der Slice ändern Sie den Typ auf `1` (eins), was Swap bedeutet.

`BSD disklabel command (m for help):` `n`

```
Partition (a-p): a
First cylinder (1-5290, default 1): 3
Last cylinder or +size or +sizeM or +sizeK (3-5290, default 5290): +1024M

```

`BSD disklabel command (m for help):` `t`

```
Partition (a-c): a
Hex code (type L to list codes): 1

```

Nach diesen Schritten sollte Ihr Layout ähnlich dem folgenden aussehen:

`BSD disklabel command (m for help):` `p`

```
3 partitions:
#       start       end      size     fstype   [fsize bsize   cpg]
  a:        3      1003      1001       swap
  c:        1      5290*     5289*    unused        0     0

```

### Creating the boot slice

Create a boot file system containing the kernels and boot loader configuration file (aboot.conf). aboot supports ext2 and ext3 filesystems only. Create the boot slice starting from the first cylinder after the swap slice. Use the `p` command to view where the swap slice ends. In our example, this is at 1003, making the boot slice start at 1004.

`BSD disklabel command (m for help):` `n`

```
Partition (a-p): b
First cylinder (1-5290, default 1): 1004
Last cylinder or +size or +sizeM or +sizeK (3-5290, default 5290): +1024M

```

`BSD disklabel command (m for help):` `t`

```
Partition (a-c): b
Hex code (type L to list codes): 08

```

After these steps a layout similar to the following should be shown:

`BSD disklabel command (m for help):` `p`

```
3 partitions:
#       start       end      size     fstype   [fsize bsize   cpg]
  a:        3      1003      1001       swap
  b:     1004      2005      1001       ext2
  c:        1      5290*     5289*    unused        0     0

```

### Root Slice erstellen

Wer werden nun die Root Slice beginnend mit dem ersten Zylinder nach der Swap Slice erstellen. Verwenden Sie den `p` Befehl um zu sehen, wo die Swap Slice endet. In unserem Beispiel ist dies bei 1003, so dass die Root Slice bei 1004 beginnt.

Ein weiteres Problem ist, dass es derzeit einen Bug in `fdisk` gibt. Dieser führt dazu, dass `fdisk` meint es würde ein Zylinder mehr zur Verfügung stehen als tatsächlich der Fall ist. In anderen Worten: Wenn Sie nach dem letzten Zylinder gefragt werden, verringern Sie die Zylindernummer (im Beispiel: 5290) um eins.

Wenn die Slice erzeugt wurde ändern Sie den Typ für ext2 auf 8.

`BSD disklabel command (m for help):` `n`

```
Partition (a-p): b
First cylinder (1-5290, default 1): 1004
Last cylinder or +size or +sizeM or +sizeK (1004-5290, default 5290): 5289

```

`BSD disklabel command (m for help):` `t`

```
Partition (a-c): b
Hex code (type L to list codes): 8

```

Das Slice Layout sollte nun ähnlich dem folgenden aussehen:

`BSD disklabel command (m for help):` `p`

```
3 partitions:
#       start       end      size     fstype   [fsize bsize   cpg]
  a:        3      1003      1001       swap
  b:     1004      5289      4286       ext2
  c:        1      5290*     5289*    unused        0     0

```

### Slice Layout speichern und fdisk beenden

Beenden Sie das Programm `fdisk` durch Drücken der Taste `w`. Dies wird gleichzeitig das Slice Layout speichern.

`Command (m for help):` `w`

## Festplatte mit fdisk partitionieren (nur ARC/AlphaBIOS)

Die folgenden Abschnitte beschreiben, wie das Slice-Layout Beispiel für ARC/AlphaBIOS erzeugt wird:

Slice
Beschreibugn
/dev/sda1Boot Partition
/dev/sda2Swap Partition
/dev/sda3Root Partition

Ändern Sie das Layout der Partitionen entsprechend Ihrer Vorstellungen ab.

### Verfügbare Festplatten identifizieren

Um herauszufinden welche Festplatten im System laufen, verwenden Sie die Folgenden Befehle:

Für IDE Festplatten:

`root #` `dmesg | grep 'drive$'`

Für SCSI Festplatten:

`root #` `dmesg | grep 'scsi'`

Abhängig von der Ausgabe sollte einfach zu erkennen sein, welche Festplatten gefunden wurden und deren zugehörige /dev/ Einträge. In den folgenden Abschnitten gehen wir davon aus, Dass es sich um eine SCSI Festplatte an /dev/sda handelt.

Starten Sie jetzt fdisk:

`root #` `fdisk /dev/sda`

### Löschen aller Partitionen

Wenn die Festplatte komplett leer ist, dann müssen Sie zuerst ein DOS Disklabel erstellen.

`Command (m for help):` `o`

```
Building a new DOS disklabel.

```

Wir beginnen mit dem Löschen aller Partitionen. Im folgenden wird gezeigt, wie man eine Partition löscht (im Beispiel verwenden wir '1'). Wiederholen Sie den Vorgang um auch alle anderen Partitionen zu löschen.

Verwenden Sie `p` um alle existierenden Partitionen anzeigen zu lassen. Betätigen Sie `d` zum Löschen einer Partition.

`command (m for help):` `p`

```
Disk /dev/sda: 9150 MB, 9150996480 bytes
64 heads, 32 sectors/track, 8727 cylinders
Units = cylinders of 2048 * 512 = 1048576 bytes

   Device Boot      Start         End      Blocks   Id  System
/dev/sda1               1         478      489456   83  Linux
/dev/sda2             479        8727     8446976    5  Extended
/dev/sda5             479        1433      977904   83  Linux Swap
/dev/sda6            1434        8727     7469040   83  Linux

```

`command (m for help):` `d`

```
Partition number (1-6): 1

```

### Boot Partition erstellen

Auf Alpha-Systemen die MILO zum Booten verwenden, müssen wir eine kleine vfat Bootpartition erstellen.

`Command (m for help):` `n`

```
Command action
  e   extended
  p   primary partition (1-4)
p
Partition number (1-4): 1
First cylinder (1-8727, default 1): 1
Last cylinder or +size or +sizeM or +sizeK (1-8727, default 8727): +16M

```

`Command (m for help):` `t`

```
Selected partition 1
Hex code (type L to list codes): 6
Changed system type of partition 1 to 6 (FAT16)

```

### Swap Partition erstellen

Wir werden eine Swap Partition mit einer Gesamtkapazität von 1 GB erzeugen. Drücken Sie `n` um eine neue Partition zu erstellen.

`Command (m for help):` `n`

```
Command action
  e   extended
  p   primary partition (1-4)
p
Partition number (1-4): 2
First cylinder (17-8727, default 17): 17
Last cylinder or +size or +sizeM or +sizeK (17-8727, default 8727): +1000M

```

`Command (m for help):` `t`

```
Partition number (1-4): 2
Hex code (type L to list codes): 82
Changed system type of partition 2 to 82 (Linux swap)

```

Nach diesen Schritten sollte Ihr Layout ähnlich dem folgenden aussehen:

`Command (m for help):` `p`

```
Disk /dev/sda: 9150 MB, 9150996480 bytes
64 heads, 32 sectors/track, 8727 cylinders
Units = cylinders of 2048 * 512 = 1048576 bytes

   Device Boot      Start         End      Blocks   Id  System
/dev/sda1               1          16       16368    6  FAT16
/dev/sda2              17         971      977920   82  Linux swap

```

### Root Partition erstellen

Wir werden nun die Roop Partition erzeugen. Verwenden Sie wieder einfach den Befehl `n`.

`Command (m for help):` `n`

```
Command action
  e   extended
  p   primary partition (1-4)
p
Partition number (1-4): 3
First cylinder (972-8727, default 972): 972
Last cylinder or +size or +sizeM or +sizeK (972-8727, default 8727): 8727

```

Nach diesen Schritten sollte Ihr Layout ähnlich dem folgenden aussehen:

`Command (m for help):` `p`

```
Disk /dev/sda: 9150 MB, 9150996480 bytes
64 heads, 32 sectors/track, 8727 cylinders
Units = cylinders of 2048 * 512 = 1048576 bytes

   Device Boot      Start         End      Blocks   Id  System
/dev/sda1               1          16       16368    6  FAT16
/dev/sda2              17         971      977920   82  Linux swap
/dev/sda3             972        8727     7942144   83  Linux

```

### Partitionslayout Speichern und fdisk beenden

Speichern Sie die Änderungen, die Sie in `fdisk` vorgenommen haben durch Eingabe von `w`. Durch diesen Befehl wird gleichzeitig auch das Programm beendet.

`Command (m for help):` `w`

Nachdem die Partitionen jetzt erstellt sind, fahren sie nun beim Abschnitt [Dateisystem erstellen](#Creating_filesystems) fort.

## Erstellen von Dateisystemen

**Warnung**

Wenn Sie ein SSD- oder NVMe-Laufwerk verwenden, prüfen Sie bitte, ob es ein Firmware-Upgrade benötigt. Insbesondere einige Intel-SSDs (600p und 6000p) benötigen ein Firmware-Upgrade für [kritische Fehlerbehebungen](https://bugzilla.redhat.com/show_bug.cgi?id=1402533), um Datenbeschädigungen zu vermeiden, die durch XFS-I/O-Nutzungsmuster verursacht werden (allerdings nicht durch einen Fehler des Dateisystems). smartctl kann helfen, das Modell und die Firmware-Version zu überprüfen.

### Einleitung

Nachdem die Partitionen angelegt wurden, ist es an der Zeit, Dateisysteme darauf anzulegen. Im nächsten Abschnitt werden die unterschiedlichen Dateisysteme beschrieben, die Linux unterstützt. Leser, die bereits wissen, welches Dateisystem sie verwenden wollen, können bei [Dateisystem auf einer Partition anlegen](/wiki/Handbook:Alpha/Installation/Networking/de#Applying_a_filesystem_to_a_partition "Handbook:Alpha/Installation/Networking/de") fortfahren. Alle anderen sollten weiterlesen, um mehr über die verfügbaren Dateisysteme zu erfahren ...

### Dateisysteme

Linux unterstützt mehrere Dutzend Dateisysteme, wobei allerdings viele davon für ganz spezielle Anwendungszwecke optimiert sind. Nur einige Dateisysteme gelten als stabil auf der alpha Architektur. Es ist ratsam, sich über Dateisysteme und deren Unterstützungsgrad zu informieren, damit Sie nicht für wichtige Partitionen ein eher experimentelles Dateisystem wählen. **XFS ist das empfohlene all-round Dateisystem für alle Plattformen.** Nachfolgend eine nicht-vollständige Auswahl von verfügbaren Dateisystemen.

[btrfs](/wiki/Btrfs/de "Btrfs/de")Dateisystem der neueren Generation.

Bietet erweiterte Funktionen wie Snapshotting, Selbstheilung durch Prüfsummen, transparente Kompression, Subvolumes und integriertes RAID. Kernel vor 5.4.y sind nicht garantiert sicher für die Verwendung mit btrfs in der Produktion, da Korrekturen für ernsthafte Probleme nur in den neueren Versionen der LTS-Kernelzweige vorhanden sind. RAID 5/6 und Quota-Gruppen sind bei allen Versionen von btrfs unsicher.

[ext4](/wiki/Ext4/de "Ext4/de")Ext4 ist ein zuverlässiges, universell einsetztbares Dateisystem für alle Plattformen, auch wenn ihm moderne Funktionen wie Reflinks fehlen.[f2fs](/wiki/F2fs "F2fs")Das Flash-Friendly File System wurde ursprünglich von Samsung für die Verwendung mit NAND-Flash-Speicher entwickelt. Es ist eine gute Wahl für die Installation von Gentoo auf microSD-Karten, USB-Laufwerken oder anderen Flash-basierten Speichergeräten.[XFS](/wiki/XFS/de "XFS/de")Dateisystem mit Metadaten-Journaling, das über einen robusten Funktionsumfang verfügt und für Skalierbarkeit optimiert ist. Es wurde kontinuierlich weiterentwickelt, um moderne Funktionen einzubeziehen. Der einzige Nachteil ist, dass XFS-Partitionen noch nicht verkleinert werden können, obwohl daran gearbeitet wird. XFS unterstützt vor allem Reflinks und Copy on Write (CoW), was besonders auf Gentoo-Systemenen hilfreich ist, da die Benutzer viele Kompilierungen durchführen müssen. XFS ist das empfohlene modernen Allzweck-Dateisystem für alle Plattformen. Erfordert, dass eine Partition mindestens 300 MB groß ist.[VFAT](/wiki/VFAT "VFAT")Auch bekannt als FAT32, wird von Linux unterstützt, unterstützt aber nicht die Standard-UNIX-Berechtigungseinstellungen. Es wird hauptsächlich für die Interoperabilität/den Austausch mit anderen Betriebssystemen (Microsoft Windows oder Apples MacOS) verwendet, ist aber auch eine Notwendigkeit für einige System-Bootloader-Firmware (wie UEFI). Benutzer von UEFI-Systemen benötigen eine [EFI System Partition](/wiki/EFI_System_Partition/de "EFI System Partition/de"), die mit VFAT formatiert ist, um booten zu können.[NTFS](/wiki/NTFS/de "NTFS/de")Dieses 'New Technology"-Dateisystem ist das Vorzeige-Dateisystem von Microsoft Windows seit Windows NT 3.1. Ähnlich wie VFAT speichert es keine UNIX-Berechtigungseinstellungen oder erweiterte Attribute, die für BSD oder Linux notwendig sind, um ordnungsgemäß zu funktionieren, daher sollte es in den meisten Fällen nicht als Root-Dateisystem verwendet werden. Es sollte _nur_ für die Interoperabilität oder den Datenaustausch mit Microsoft Windows-Systemen verwendet werden (beachten Sie die Betonung auf _nur_).

Ausführlichere Informationen über Dateisysteme finden Sie in dem von der Community gepflegten [Dateisystem](/wiki/Filesystem/de "Filesystem/de")-Artikel.

### Dateisystem auf einer Partition anlegen

**Hinweis**

Bitte stellen Sie sicher, dass Sie das entsprechende Paket für das gewählte Dateisystem später im Handbuch emergen, bevor Sie am Ende des Installationsprozesses neu booten.

Dateisysteme können mit Hilfe von Programmen auf einer Partition oder auf einem Datenträger angelegt werden. Die folgende Tabelle zeigt, welchen Befehl Sie für welches Dateisystem benötigen. Um weitere Informationen zu einem Dateisystem zu erhalten, können Sie auf den Namen des Dateisystems klicken.

Dateisystem
Befehl zum Anlegen
Teil der Minimal CD?
Gentoo Paket
[btrfs](/wiki/Btrfs/de "Btrfs/de")mkfs.btrfs Yes
[sys-fs/btrfs-progs](https://packages.gentoo.org/packages/sys-fs/btrfs-progs)[ext4](/wiki/Ext4/de "Ext4/de")mkfs.ext4 Yes
[sys-fs/e2fsprogs](https://packages.gentoo.org/packages/sys-fs/e2fsprogs)[f2fs](/wiki/F2FS "F2FS")mkfs.f2fs Yes
[sys-fs/f2fs-tools](https://packages.gentoo.org/packages/sys-fs/f2fs-tools)[xfs](/wiki/XFS/de "XFS/de")mkfs.xfs Yes
[sys-fs/xfsprogs](https://packages.gentoo.org/packages/sys-fs/xfsprogs)[vfat](/wiki/FAT "FAT")mkfs.vfat Yes
[sys-fs/dosfstools](https://packages.gentoo.org/packages/sys-fs/dosfstools)[NTFS](/wiki/NTFS/de "NTFS/de")mkfs.ntfs Yes
[sys-fs/ntfs3g](https://packages.gentoo.org/packages/sys-fs/ntfs3g)

\|}

**Wichtig**

The handbook recommends new partitions as part of the installation process, but it is important to note running any mkfs command will erase any data contained within the partition. When necessary, ensure any data that exists within is appropriately backed up before creating a new filesystem.

Um beispielsweise die EFI System-Partition (/dev/sda1) als FAT32 und die root-Partition (/dev/sda3) als xfs zu formatieren (wie in dem Beispiel-Partitionsschema), würde man folgende Befehle verwenden:

`root #` `mkfs.xfs /dev/sda3`

{{#ifeq: 0\|1\|

#### EFI system partition filesystem

The EFI system partition (/dev/sda1) must be formatted as FAT32:

`root #` `mkfs.vfat -F 32 /dev/sda1`

#### Legacy BIOS boot partition filesystem

Systems booting via legacy BIOS with a MBR/DOS disklabel can use any filesystem format supported by the bootloader.

For example, to format with XFS:

`root #` `mkfs.xfs -c options=/usr/share/xfsprogs/mkfs/lts_6.18.conf /dev/sda1`

#### Small ext4 partitions

Bei der Verwendung von ext4 auf kleinen Partitionen (kleiner als 8 GiB), sollte das Dateisystem mit den passenden Optionen erstellt werden, um genügend Inodes zu reservieren. Dies kann mit einer der folgenden Anweisungen erfolgen:

`root #` `mkfs.ext4 -T small /dev/<device>`

Dies vervierfacht die Zahl der Inodes für ein angegebenes Dateisystem in der Regel, da es dessen "bytes-per-inode" (Bytes pro Inode) von 16 kB auf 4 kB pro Inode reduziert.

### Aktivieren der Swap-Partition

mkswap ist der Befehl der verwendet wird um Swap-Partitionen zu initialisieren:

`root #` `mkswap /dev/sda2`

**Hinweis**

Installations which were previously started, but did not finish the installation process can resume the installation from this point in the handbook. Use this link as the permalink: [Resumed installations start here](/wiki/Handbook:Alpha/Installation/Disks/de#Resumed_installations_start_here "Handbook:Alpha/Installation/Disks/de").

Zur Aktivierung der Swap-Partition verwenden Sie swapon:

`root #` `swapon /dev/sda2`

This 'activation' step is only necessary because the swap partition is newly created within the live environment. Once the system has been rebooted, as long as the swap partition is properly defined within fstab or other mount mechanism, swap space will activate automatically.

## Einhängen der Root-Partition

**Tipp**

Anwender, die ein Nicht-Gentoo Installationsmedium verwenden, müssen mit folgendem Befehl einen Mount-Point erzeugen:

`root #` `mkdir --parents /mnt/gentoo`

`root #` `mkdir --parents /mnt/gentoo`

Continue creating additional mount points necessary for any additional (custom) partition(s) created during previous steps by using the mkdir command.

Nachdem die Partitionen initialisiert wurden und ein Dateisystem beinhalten, ist es an der Zeit, diese einzuhängen. Verwenden Sie den Befehl mount, aber vergessen Sie nicht die notwendigen Einhänge-Verzeichnisse für jede Partition zu erzeugen. Als Beispiel hängen wir die Root-Partition ein:

Mount the root partition:

`root #` `mount /dev/sda3 /mnt/gentoo`

Continue mounting additional (custom) partitions as necessary using the mount command.

**Hinweis**

Wenn sich /tmp/ auf einer separaten Partition befinden muss, ändern Sie die Berechtigungen nach dem Einhängen:

`root #` `chmod 1777 /mnt/gentoo/tmp`

Dies gilt ebenfalls für /var/tmp.

In der Anleitung wird später das Dateisystem proc (eine virtuelle Schnittstelle zum Kernel) zusammen mit anderen Kernel Pseudo-Dateisystemen eingehängt. Zunächst installieren wir jedoch die [Gentoo Installationsdateien](/wiki/Handbook:Alpha/Installation/Stage/de "Handbook:Alpha/Installation/Stage/de").

[← Konfiguration des Netzwerks](/wiki/Handbook:Alpha/Installation/Networking/de "Previous part") [Anfang](/wiki/Handbook:Alpha/de "Handbook:Alpha/de") [Installation des Stage Archivs →](/wiki/Handbook:Alpha/Installation/Stage/de "Next part")