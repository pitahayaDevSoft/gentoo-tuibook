# Disks

Other languages:

- Deutsch
- [English](/wiki/Handbook:MIPS/Installation/Disks "Handbook:MIPS/Installation/Disks (100% translated)")
- [español](/wiki/Handbook:MIPS/Installation/Disks/es "Manual de Gentoo: MIPS/Instalación/Discos (100% translated)")
- [français](/wiki/Handbook:MIPS/Installation/Disks/fr "Handbook:MIPS/Installation/Disks/fr (100% translated)")
- [italiano](/wiki/Handbook:MIPS/Installation/Disks/it "Handbook:MIPS/Installation/Disks/it (100% translated)")
- [magyar](/wiki/Handbook:MIPS/Installation/Disks/hu "Handbook:MIPS/Installation/Disks/hu (100% translated)")
- [polski](/wiki/Handbook:MIPS/Installation/Disks/pl "Handbook:MIPS/Installation/Disks (100% translated)")
- [português do Brasil](/wiki/Handbook:MIPS/Installation/Disks/pt-br "Manual:MIPS/Instalação/Discos (100% translated)")
- [русский](/wiki/Handbook:MIPS/Installation/Disks/ru "Handbook:MIPS/Installation/Disks (100% translated)")
- [தமிழ்](/wiki/Handbook:MIPS/Installation/Disks/ta "கையேடு:MIPS/நிறுவல்/வட்டுக்கள் (100% translated)")
- [中文（中国大陆）‎](/wiki/Handbook:MIPS/Installation/Disks/zh-cn "手册：MIPS/安装/准备磁盘 (100% translated)")
- [日本語](/wiki/Handbook:MIPS/Installation/Disks/ja "ハンドブック:MIPS/インストール/ディスク (100% translated)")
- [한국어](/wiki/Handbook:MIPS/Installation/Disks/ko "Handbook:MIPS/Installation/Disks/ko (100% translated)")

[MIPS Handbuch](/wiki/Handbook:MIPS/de "Handbook:MIPS/de")[Installation](/wiki/Handbook:MIPS/Full/Installation/de "Handbook:MIPS/Full/Installation/de")[Über die Installation](/wiki/Handbook:MIPS/Installation/About/de "Handbook:MIPS/Installation/About/de")[Auswahl des Mediums](/wiki/Handbook:MIPS/Installation/Media/de "Handbook:MIPS/Installation/Media/de")[Konfiguration des Netzwerks](/wiki/Handbook:MIPS/Installation/Networking/de "Handbook:MIPS/Installation/Networking/de")Vorbereiten der Festplatte(n)[Installation des Stage Archivs](/wiki/Handbook:MIPS/Installation/Stage/de "Handbook:MIPS/Installation/Stage/de")[Installation des Basissystems](/wiki/Handbook:MIPS/Installation/Base/de "Handbook:MIPS/Installation/Base/de")[Konfiguration des Kernels](/wiki/Handbook:MIPS/Installation/Kernel/de "Handbook:MIPS/Installation/Kernel/de")[Konfiguration des Systems](/wiki/Handbook:MIPS/Installation/System/de "Handbook:MIPS/Installation/System/de")[Installation der Tools](/wiki/Handbook:MIPS/Installation/Tools/de "Handbook:MIPS/Installation/Tools/de")[Konfiguration des Bootloaders](/wiki/Handbook:MIPS/Installation/Bootloader/de "Handbook:MIPS/Installation/Bootloader/de")[Abschluss](/wiki/Handbook:MIPS/Installation/Finalizing/de "Handbook:MIPS/Installation/Finalizing/de")[Arbeiten mit Gentoo](/wiki/Handbook:MIPS/Full/Working/de "Handbook:MIPS/Full/Working/de")[Portage-Einführung](/wiki/Handbook:MIPS/Working/Portage/de "Handbook:MIPS/Working/Portage/de")[USE-Flags](/wiki/Handbook:MIPS/Working/USE/de "Handbook:MIPS/Working/USE/de")[Portage-Features](/wiki/Handbook:MIPS/Working/Features/de "Handbook:MIPS/Working/Features/de")[Initskript-System](/wiki/Handbook:MIPS/Working/Initscripts/de "Handbook:MIPS/Working/Initscripts/de")[Umgebungsvariablen](/wiki/Handbook:MIPS/Working/EnvVar/de "Handbook:MIPS/Working/EnvVar/de")[Arbeiten mit Portage](/wiki/Handbook:MIPS/Full/Portage/de "Handbook:MIPS/Full/Portage/de")[Dateien und Verzeichnisse](/wiki/Handbook:MIPS/Portage/Files/de "Handbook:MIPS/Portage/Files/de")[Variablen](/wiki/Handbook:MIPS/Portage/Variables/de "Handbook:MIPS/Portage/Variables/de")[Mischen von Softwarezweigen](/wiki/Handbook:MIPS/Portage/Branches/de "Handbook:MIPS/Portage/Branches/de")[Zusätzliche Tools](/wiki/Handbook:MIPS/Portage/Tools/de "Handbook:MIPS/Portage/Tools/de")[Eigener Portage-Tree](/wiki/Handbook:MIPS/Portage/CustomTree/de "Handbook:MIPS/Portage/CustomTree/de")[Erweiterte Portage-Features](/wiki/Handbook:MIPS/Portage/Advanced/de "Handbook:MIPS/Portage/Advanced/de")[Netzwerk-Konfiguration](/wiki/Handbook:MIPS/Full/Networking/de "Handbook:MIPS/Full/Networking/de")[Zu Beginn](/wiki/Handbook:MIPS/Networking/Introduction/de "Handbook:MIPS/Networking/Introduction/de")[Fortgeschrittene Konfiguration](/wiki/Handbook:MIPS/Networking/Advanced/de "Handbook:MIPS/Networking/Advanced/de")[Modulare Vernetzung](/wiki/Handbook:MIPS/Networking/Modular/de "Handbook:MIPS/Networking/Modular/de")[Drahtlose Netzwerke](/wiki/Handbook:MIPS/Networking/Wireless/de "Handbook:MIPS/Networking/Wireless/de")[Funktionalität hinzufügen](/wiki/Handbook:MIPS/Networking/Extending/de "Handbook:MIPS/Networking/Extending/de")[Dynamisches Management](/wiki/Handbook:MIPS/Networking/Dynamic/de "Handbook:MIPS/Networking/Dynamic/de")

## Contents

- [1Einführung in blockorientierte Geräte](#Einf.C3.BChrung_in_blockorientierte_Ger.C3.A4te)
  - [1.1Blockorientierte Geräte](#Blockorientierte_Ger.C3.A4te)
  - [1.2Partitionen](#Partitionen)
- [2Ein Partitionsschema entwerfen](#Ein_Partitionsschema_entwerfen)
  - [2.1Wie viele Partitionen und wie groß?](#Wie_viele_Partitionen_und_wie_gro.C3.9F.3F)
  - [2.2Was ist mit dem Swap-Speicher?](#Was_ist_mit_dem_Swap-Speicher.3F)
- [3fdisk verwenden](#fdisk_verwenden)
  - [3.1SGI Maschinen: SGI Plattenlabel erstellen](#SGI_Maschinen:_SGI_Plattenlabel_erstellen)
  - [3.2SGI Volume-Header Größenänderung](#SGI_Volume-Header_Gr.C3.B6.C3.9Fen.C3.A4nderung)
  - [3.3Cobalt Festplatten partitionieren](#Cobalt_Festplatten_partitionieren)
- [4Erstellen von Dateisystemen](#Erstellen_von_Dateisystemen)
  - [4.1Einleitung](#Einleitung)
  - [4.2Dateisysteme](#Dateisysteme)
  - [4.3Dateisystem auf einer Partition anlegen](#Dateisystem_auf_einer_Partition_anlegen)
    - [4.3.1EFI system partition filesystem](#EFI_system_partition_filesystem)
    - [4.3.2Legacy BIOS boot partition filesystem](#Legacy_BIOS_boot_partition_filesystem)
    - [4.3.3Small ext4 partitions](#Small_ext4_partitions)
  - [4.4Aktivieren der Swap-Partition](#Aktivieren_der_Swap-Partition)
- [5Einhängen der Root-Partition](#Einh.C3.A4ngen_der_Root-Partition)

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

### Partitionen

Obwohl es theoretisch möglich wäre eine vollständige Festplatte zu nutzen um ein Linux-System unterzubringen, kommt das in der Praxis fast nie vor. Stattdessen werden komplette Festplatten Block Devices in kleinere, besser handhabbare Block Devices unterteilt. Diese werden _Partitionen_ genannt.

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

## fdisk verwenden

### SGI Maschinen: SGI Plattenlabel erstellen

Alle Festplatten in einem SGI System benötigen ein SGI Plattenlabel, welches eine ähnliche Funktionalität wie ein Sun oder MS-DOS Plattenlabel bietet -- es speichert Informationen über die Partitionen einer Festplatte. Die Erzeugung eines neuen SGI Plattenlabels erzeugt zwei spezielle Partitionen auf der Festplatte:

- SGI Volume Header (9. Partition): Diese Partition ist wichtig. Sie ist der Ort an dem sich der Bootloader befindet und in einigen Fällen enthält sie ebenfalls die Kernel-Abbilder.
- SGI Volume (11. Partition): Diese Partition ist ähnlich wichtig wie die dritte Partition des Sun Plattenlabels "Whole Disk". Diese Partition umschließt die gesamte Festplatte und solle unberührt bleiben. Sie dient keinem anderen speziellen Zweck außer das PROM in undokumentierter Weise zu unterstützen (oder es wird irgendwie von IRIX verwendet).

**Warnung**

Der SGI Volume Header muss bei Zylinder 0 beginnen. Ein Fehler hierbei bedeutet ein Scheitern beim Booten von der Platte.

Das Folgende ist ein Beispiel-Auszug einer fdisk Sitzung. Lesen und passen Sie es Ihren persönlichen Bedürfnissen an ...

`root #` `fdisk /dev/sda`

Wechseln Sie in den Expertenmodus:

`Command (m for help):` `x`

Mit `m` wird das vollständige Menü der Optionen angezeigt:

`Expert command (m for help):` `m`

```
Command action
   b   move beginning of data in a partition
   c   change number of cylinders
   d   print the raw data in the partition table
   e   list extended partitions
   f   fix partition order
   g   create an IRIX (SGI) partition table
   h   change number of heads
   m   print this menu
   p   print the partition table
   q   quit without saving changes
   r   return to main menu
   s   change number of sectors/track
   v   verify the partition table
   w   write table to disk and exit

```

Erzeugen Sie ein SGI Plattenlabel:

`Expert command (m for help):` `g`

```
Building a new SGI disklabel. Changes will remain in memory only,
until you decide to write them. After that, of course, the previous
content will be irrecoverably lost.

```

Kehren Sie zum Hauptmenü zurück:

`Expert command (m for help):` `r`

Werfen wir einen Blick auf das aktuelle Partitions-Layout:

`Command (m for help):` `p`

```
Disk /dev/sda (SGI disk label): 64 heads, 32 sectors, 17482 cylinders
Units = cylinders of 2048 * 512 bytes

----- partitions -----
Pt#     Device  Info     Start       End   Sectors  Id  System
 9:  /dev/sda1               0         4     10240   0  SGI volhdr
11:  /dev/sda2               0     17481  35803136   6  SGI volume
----- Bootinfo -----
Bootfile: /unix
----- Directory Entries -----

```

**Hinweis**

Wenn die Festplatte bereits ein bestehendes SGI Plattenlabel hat, wird fdisk die Erzeugung eines neuen Labels nicht gestatten. Es gibt zwei Möglichkeiten das zu umgehen. Die erste ist die Erstellung eines Sun oder MS-DOS Plattenlabels, die Änderungen auf die Festplatte zu schreiben und fdisk neu zu starten. Die zweite ist die Partitionstabelle mit Nullwerten durch den folgenden Befehl zu überschreiben: `dd if=/dev/zero of=/dev/sda bs=512 count=1`

### SGI Volume-Header Größenänderung

**Wichtig**

Diesen Schritt benötigt man aufgrund eines Bugs in fdisk oft. Aus irgendeinem Grund wird der Volume-Header nicht korrekt erstellt. Das Ergebnis ist, dass er auf Zylinder 0 startet und endet. Dies verhindert, dass mehreren Partitionen erstellt werden. Um dieses Problem zu umgehen ... lesen Sie weiter.

Da jetzt ein SGI Plattenlabel erstellt wurde, können nun Partitionen definiert werden. Im obigen Beispiel sind bereits zwei Partitionen definiert. Das sind wie erwähnt die besonderen Partitionen und sie sollten normalerweise nicht verändert werden. Wie auch immer, zur Installation von Gentoo müssen wir einen Bootloader und möglicherweise mehrere Kernel-Abbilder (abhängig vom Systemtyp) direkt in den Volume-Header laden. Der Volume-Header selbst kann bis zu acht Abbilder jeglicher Größe beinhalten mit jeweils einem acht Zeichen langen Namen.

Der Vorgang den Volume-Header größer zu machen ist etwas verworren und mit einem kleinen Trick verbunden. Man kann den Volume-Header wegen dem eigenartigen Verhalten von fdisk nicht einfach löschen und ihn dann wieder neu hinzufügen. Im Beispiel unten erzeugen wir einen 50 MB großen Volume-Header in Verbindung mit einer 50 MB großen /boot/ Partition. Das tatsächliche Plattenlayout kann sich unterscheiden, dies dient nur der Veranschaulichung.

Eine neue Partition erstellen:

`Command (m for help):` `n`

```
Partition number (1-16): 1
First cylinder (5-8682, default 5): 51
 Last cylinder (51-8682, default 8682): 101

```

Beachten Sie, dass fdisk zur Neuerstellung von Partition Nr. 1 als kleinsten Zylinder 5 gestattet. Wenn wir versucht hätten den SGI Volume-Header zu löschen und auf diese Weise wiederherzustellen, würden wir vor dem gleichen Problem stehen. In unserem Beispiel wollen wir dass /boot/ 50 MB groß ist, deshalb starten wir bei Zylinder 51. (- Der Volume-Header muss bei Zylinder 0 beginnen, erinnern Sie sich?) Den End-Zylinder setzten wir bei 101, das in etwa 50 MB entspricht (+/- 1..5 MB).

Die Partition löschen:

`Command (m for help):` `d`

```
Partition number (1-16): 9

```

Jetzt die Neuerstellung der Volume-Header Partition:

`Command (m for help):` `n`

```
Partition number (1-16): 9
First cylinder (0-50, default 0): 0
 Last cylinder (0-50, default 50): 50

```

Wenn Sie sich unsicher über die Verwendung fdisk sind, werfen sie weiter unten einen Blick auf die Anleitung zur Partitionierung auf Cobalt Systemen. Das Konzept ist genau das gleiche -- denken Sie nur daran die Volume-Header und die "Whole Disk" Partition in Ruhe zu lassen.

Sobald dies geschehen ist können Sie die übrigen Partitionen die Sie benötigen erzeugen. Nachdem Sie alle Partitionen angelegt haben, stellen Sie sicher die Partitions-ID der Swap Partition auf 82 zu stellen, "Linux Swap". Der Standard ist 83, "Linux Native".

### Cobalt Festplatten partitionieren

Auf Cobalt Maschinen erwartet das BOOTROM einen MS-DOS MBR, deshalb ist die Festplattenpartitionierung relativ geradlinig. -- In der Tat wird dies wie bei einer Intel x86 Maschine gemacht. Es gibt jedoch ein paar Dinge die Sie beachten sollten.

- Die Cobalt Firmware erwartet /dev/sda1 als Linux Partition im Format EXT2 Revision 0. EXT Revision 1 Partitionen funktionieren NICHT! (Das Cobalt BOOTROM versteht nur EXT2r0.)
- Die oben angesprochene Partition muss das gzip-komprimierte ELF Abbild vmlinux.gz in der Wurzel ("root") dieser Partition enthalten, das als Kernel geladen wird.

Aus diesem Grund wird eine ca. 20 MB große mit EXT2r0 formatierte /boot/ Partition empfohlen auf der CoLo und die Kernel installiert werden. Dies gestattet dem Benutzer ein modernes Dateisystem (EXT3 oder ReiserFS) auf der Root Partition zu betreiben.

Im Beispiel wird davon ausgegangen, dass /dev/sda1 erzeugt wird, um später als /boot/ Partition eingehängt zu werden. Falls Sie die Partition zu / machen wollen, denken Sie an die Erwartungen des PROMs.

Also weiter ... Um die Partition zu erstellen geben Sie an der Eingabeaufforderung fdisk /dev/sda ein. Die wichtigsten Befehle, die Sie wissen sollten sind diese:

CODE **Wichtige fdisk Befehle**

```
'"`UNIQ--pre-00000007-QINU`"'

```

`root #` `fdisk /dev/sda`

```
The number of cylinders for this disk is set to 19870.
There is nothing wrong with that, but this is larger than 1024,
and could in certain setups cause problems with:
1) software that runs at boot time (e.g., old versions of LILO)
2) booting and partitioning software from other OSs
   (e.g., DOS FDISK, OS/2 FDISK)

```

Fangen Sie damit an, alle vorhandenen Partitionen zu löschen:

`Command (m for help):` `o`

```
Building a new DOS disklabel. Changes will remain in memory only,
until you decide to write them. After that, of course, the previous
content won't be recoverable.


The number of cylinders for this disk is set to 19870.
There is nothing wrong with that, but this is larger than 1024,
and could in certain setups cause problems with:
1) software that runs at boot time (e.g., old versions of LILO)
2) booting and partitioning software from other OSs
   (e.g., DOS FDISK, OS/2 FDISK)
Warning: invalid flag 0x0000 of partition table 4 will be corrected by w(rite)

```

Überprüfen Sie nun durch Drücken der Befehlstaste `p`, dass die Partitionstabelle leer ist:

`Command (m for help):` `p`

```
Disk /dev/sda: 10.2 GB, 10254827520 bytes
16 heads, 63 sectors/track, 19870 cylinders
Units = cylinders of 1008 * 512 = 516096 bytes

   Device Boot      Start         End      Blocks   Id  System

```

Erstellen Sie die /boot Partition:

`Command (m for help):` `n`

```
Command action
   e   extended
   p   primary partition (1-4)
p
Partition number (1-4): 1
First cylinder (1-19870, default 1):
Last cylinder or +size or +sizeM or +sizeK (1-19870, default 19870): +20M

```

Wenn Sie die Partitionen ausgeben lassen, beachten Sie die neu erstellte:

`Command (m for help):` `p`

```
Disk /dev/sda: 10.2 GB, 10254827520 bytes
16 heads, 63 sectors/track, 19870 cylinders
Units = cylinders of 1008 * 512 = 516096 bytes

   Device Boot      Start         End      Blocks   Id  System
/dev/sda1               1          40       20128+  83  Linux

```

Lassen Sie uns nun eine erweiterte Partition erstellen, die den Rest der Festplatte umfasst. In dieser erweiterten Partition legen wir die übrigen Partitionen (logische Partitionen) an:

`Command (m for help):` `n`

```
Command action
   e   extended
   p   primary partition (1-4)
e
Partition number (1-4): 2
First cylinder (41-19870, default 41):
Using default value 41
Last cylinder or +size or +sizeM or +sizeK (41-19870, default 19870):
Using default value 19870

```

Jetzt erstellen wir die Partitionen /, /usr, /var usw.

`Command (m for help):` `n`

```
Command action
   l   logical (5 or over)
   p   primary partition (1-4)
l
First cylinder (41-19870, default 41):<Press ENTER>
Using default value 41
Last cylinder or +size or +sizeM or +sizeK (41-19870, default 19870): +500M

```

Wiederholen Sie dies wie benötigt.

Zum Schluss zur Swap Partition. Es wird empfohlen mindestens 250 MB, besser 1 GB Speicherplatz zu verwenden:

`Command (m for help):` `n`

```
Command action
   l   logical (5 or over)
   p   primary partition (1-4)
l
First cylinder (17294-19870, default 17294): <Press ENTER>
Using default value 17294
Last cylinder or +size or +sizeM or +sizeK (1011-19870, default 19870): <Press ENTER>
Using default value 19870

```

Wenn Sie die Partitionstabelle überprüfen, sollte alles bereit sein - bis auf eine Sache.

`Command (m for help):` `p`

```
Disk /dev/sda: 10.2 GB, 10254827520 bytes
16 heads, 63 sectors/track, 19870 cylinders
Units = cylinders of 1008 * 512 = 516096 bytes

Device Boot      Start         End      Blocks      ID  System
/dev/sda1               1          21       10552+  83  Linux
/dev/sda2              22       19870    10003896    5  Extended
/dev/sda5              22        1037      512032+  83  Linux
/dev/sda6            1038        5101     2048224+  83  Linux
/dev/sda7            5102        9165     2048224+  83  Linux
/dev/sda8            9166       13229     2048224+  83  Linux
/dev/sda9           13230       17293     2048224+  83  Linux
/dev/sda10          17294       19870     1298776+  83  Linux

```

Ist Ihnen aufgefallen, dass Partition 10 - die Swap Partition - immer noch vom Typ 83 ist? Lassen Sie uns das auf den richtigen Typ ändern:

`Command (m for help):` `t`

```
Partition number (1-10): 10
Hex code (type L to list codes): 82
Changed system type of partition 10 to 82 (Linux swap)

```

Nun zur Überprüfung:

`Command (m for help):` `p`

```
Disk /dev/sda: 10.2 GB, 10254827520 bytes
16 heads, 63 sectors/track, 19870 cylinders
Units = cylinders of 1008 * 512 = 516096 bytes

Device Boot      Start         End      Blocks      ID  System
/dev/sda1               1          21       10552+  83  Linux
/dev/sda2              22       19870    10003896    5  Extended
/dev/sda5              22        1037      512032+  83  Linux
/dev/sda6            1038        5101     2048224+  83  Linux
/dev/sda7            5102        9165     2048224+  83  Linux
/dev/sda8            9166       13229     2048224+  83  Linux
/dev/sda9           13230       17293     2048224+  83  Linux
/dev/sda10          17294       19870     1298776+  82  Linux Swap

```

Wir speichern die neue Partitionstabelle:

`Command (m for help):` `w`

```
The partition table has been altered!

Calling ioctl() to re-read partition table.
Syncing disks.

```

## Erstellen von Dateisystemen

**Warnung**

Wenn Sie ein SSD- oder NVMe-Laufwerk verwenden, prüfen Sie bitte, ob es ein Firmware-Upgrade benötigt. Insbesondere einige Intel-SSDs (600p und 6000p) benötigen ein Firmware-Upgrade für [kritische Fehlerbehebungen](https://bugzilla.redhat.com/show_bug.cgi?id=1402533), um Datenbeschädigungen zu vermeiden, die durch XFS-I/O-Nutzungsmuster verursacht werden (allerdings nicht durch einen Fehler des Dateisystems). smartctl kann helfen, das Modell und die Firmware-Version zu überprüfen.

### Einleitung

Nachdem die Partitionen angelegt wurden, ist es an der Zeit, Dateisysteme darauf anzulegen. Im nächsten Abschnitt werden die unterschiedlichen Dateisysteme beschrieben, die Linux unterstützt. Leser, die bereits wissen, welches Dateisystem sie verwenden wollen, können bei [Dateisystem auf einer Partition anlegen](/wiki/Handbook:MIPS/Installation/Networking/de#Applying_a_filesystem_to_a_partition "Handbook:MIPS/Installation/Networking/de") fortfahren. Alle anderen sollten weiterlesen, um mehr über die verfügbaren Dateisysteme zu erfahren ...

### Dateisysteme

Linux unterstützt mehrere Dutzend Dateisysteme, wobei allerdings viele davon für ganz spezielle Anwendungszwecke optimiert sind. Nur einige Dateisysteme gelten als stabil auf der mips Architektur. Es ist ratsam, sich über Dateisysteme und deren Unterstützungsgrad zu informieren, damit Sie nicht für wichtige Partitionen ein eher experimentelles Dateisystem wählen. **XFS ist das empfohlene all-round Dateisystem für alle Plattformen.** Nachfolgend eine nicht-vollständige Auswahl von verfügbaren Dateisystemen.

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

Um beispielsweise die EFI System-Partition (/dev/sda1) als FAT32 und die root-Partition (/dev/sda5) als xfs zu formatieren (wie in dem Beispiel-Partitionsschema), würde man folgende Befehle verwenden:

`root #` `mkfs.xfs /dev/sda5`

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

`root #` `mkswap /dev/sda10`

**Hinweis**

Installations which were previously started, but did not finish the installation process can resume the installation from this point in the handbook. Use this link as the permalink: [Resumed installations start here](/wiki/Handbook:MIPS/Installation/Disks/de#Resumed_installations_start_here "Handbook:MIPS/Installation/Disks/de").

Zur Aktivierung der Swap-Partition verwenden Sie swapon:

`root #` `swapon /dev/sda10`

This 'activation' step is only necessary because the swap partition is newly created within the live environment. Once the system has been rebooted, as long as the swap partition is properly defined within fstab or other mount mechanism, swap space will activate automatically.

## Einhängen der Root-Partition

**Tipp**

Anwender, die ein Nicht-Gentoo Installationsmedium verwenden, müssen mit folgendem Befehl einen Mount-Point erzeugen:

`root #` `mkdir --parents /mnt/gentoo`

`root #` `mkdir --parents /mnt/gentoo`

Continue creating additional mount points necessary for any additional (custom) partition(s) created during previous steps by using the mkdir command.

Nachdem die Partitionen initialisiert wurden und ein Dateisystem beinhalten, ist es an der Zeit, diese einzuhängen. Verwenden Sie den Befehl mount, aber vergessen Sie nicht die notwendigen Einhänge-Verzeichnisse für jede Partition zu erzeugen. Als Beispiel hängen wir die Root-Partition ein:

Mount the root partition:

`root #` `mount /dev/sda5 /mnt/gentoo`

Continue mounting additional (custom) partitions as necessary using the mount command.

**Hinweis**

Wenn sich /tmp/ auf einer separaten Partition befinden muss, ändern Sie die Berechtigungen nach dem Einhängen:

`root #` `chmod 1777 /mnt/gentoo/tmp`

Dies gilt ebenfalls für /var/tmp.

In der Anleitung wird später das Dateisystem proc (eine virtuelle Schnittstelle zum Kernel) zusammen mit anderen Kernel Pseudo-Dateisystemen eingehängt. Zunächst installieren wir jedoch die [Gentoo Installationsdateien](/wiki/Handbook:MIPS/Installation/Stage/de "Handbook:MIPS/Installation/Stage/de").

[← Konfiguration des Netzwerks](/wiki/Handbook:MIPS/Installation/Networking/de "Previous part") [Anfang](/wiki/Handbook:MIPS/de "Handbook:MIPS/de") [Installation des Stage Archivs →](/wiki/Handbook:MIPS/Installation/Stage/de "Next part")