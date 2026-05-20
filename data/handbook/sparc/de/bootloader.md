# Bootloader

Other languages:

- Deutsch
- [English](/wiki/Handbook:SPARC/Installation/Bootloader "Handbook:SPARC/Installation/Bootloader (100% translated)")
- [español](/wiki/Handbook:SPARC/Installation/Bootloader/es "Manual de Gentoo: SPARC/Instalación/Arranque (100% translated)")
- [français](/wiki/Handbook:SPARC/Installation/Bootloader/fr "Handbook:SPARC/Installation/Bootloader/fr (100% translated)")
- [italiano](/wiki/Handbook:SPARC/Installation/Bootloader/it "Handbook:SPARC/Installation/Bootloader/it (50% translated)")
- [magyar](/wiki/Handbook:SPARC/Installation/Bootloader/hu "Handbook:SPARC/Installation/Bootloader/hu (100% translated)")
- [polski](/wiki/Handbook:SPARC/Installation/Bootloader/pl "Handbook:SPARC/Installation/Bootloader (100% translated)")
- [português do Brasil](/wiki/Handbook:SPARC/Installation/Bootloader/pt-br "Handbook:SPARC/Installation/Bootloader/pt-br (50% translated)")
- [русский](/wiki/Handbook:SPARC/Installation/Bootloader/ru "Handbook:SPARC/Installation/Bootloader (100% translated)")
- [தமிழ்](/wiki/Handbook:SPARC/Installation/Bootloader/ta "கையேடு:SPARC/நிறுவல்/துவக்கஏற்றி (100% translated)")
- [中文（中国大陆）‎](/wiki/Handbook:SPARC/Installation/Bootloader/zh-cn "手册：SPARC/安装/配置系统引导程序Bootloader (100% translated)")
- [日本語](/wiki/Handbook:SPARC/Installation/Bootloader/ja "ハンドブック:SPARC/インストール/ブートローダー (100% translated)")
- [한국어](/wiki/Handbook:SPARC/Installation/Bootloader/ko "Handbook:SPARC/Installation/Bootloader/ko (100% translated)")

[SPARC Handbuch](/wiki/Handbook:SPARC/de "Handbook:SPARC/de")[Installation](/wiki/Handbook:SPARC/Full/Installation/de "Handbook:SPARC/Full/Installation/de")[Über die Installation](/wiki/Handbook:SPARC/Installation/About/de "Handbook:SPARC/Installation/About/de")[Auswahl des Mediums](/wiki/Handbook:SPARC/Installation/Media/de "Handbook:SPARC/Installation/Media/de")[Konfiguration des Netzwerks](/wiki/Handbook:SPARC/Installation/Networking/de "Handbook:SPARC/Installation/Networking/de")[Vorbereiten der Festplatte(n)](/wiki/Handbook:SPARC/Installation/Disks/de "Handbook:SPARC/Installation/Disks/de")[Installation des Stage Archivs](/wiki/Handbook:SPARC/Installation/Stage/de "Handbook:SPARC/Installation/Stage/de")[Installation des Basissystems](/wiki/Handbook:SPARC/Installation/Base/de "Handbook:SPARC/Installation/Base/de")[Konfiguration des Kernels](/wiki/Handbook:SPARC/Installation/Kernel/de "Handbook:SPARC/Installation/Kernel/de")[Konfiguration des Systems](/wiki/Handbook:SPARC/Installation/System/de "Handbook:SPARC/Installation/System/de")[Installation der Tools](/wiki/Handbook:SPARC/Installation/Tools/de "Handbook:SPARC/Installation/Tools/de")Konfiguration des Bootloaders[Abschluss](/wiki/Handbook:SPARC/Installation/Finalizing/de "Handbook:SPARC/Installation/Finalizing/de")[Arbeiten mit Gentoo](/wiki/Handbook:SPARC/Full/Working/de "Handbook:SPARC/Full/Working/de")[Portage-Einführung](/wiki/Handbook:SPARC/Working/Portage/de "Handbook:SPARC/Working/Portage/de")[USE-Flags](/wiki/Handbook:SPARC/Working/USE/de "Handbook:SPARC/Working/USE/de")[Portage-Features](/wiki/Handbook:SPARC/Working/Features/de "Handbook:SPARC/Working/Features/de")[Initskript-System](/wiki/Handbook:SPARC/Working/Initscripts/de "Handbook:SPARC/Working/Initscripts/de")[Umgebungsvariablen](/wiki/Handbook:SPARC/Working/EnvVar/de "Handbook:SPARC/Working/EnvVar/de")[Arbeiten mit Portage](/wiki/Handbook:SPARC/Full/Portage/de "Handbook:SPARC/Full/Portage/de")[Dateien und Verzeichnisse](/wiki/Handbook:SPARC/Portage/Files/de "Handbook:SPARC/Portage/Files/de")[Variablen](/wiki/Handbook:SPARC/Portage/Variables/de "Handbook:SPARC/Portage/Variables/de")[Mischen von Softwarezweigen](/wiki/Handbook:SPARC/Portage/Branches/de "Handbook:SPARC/Portage/Branches/de")[Zusätzliche Tools](/wiki/Handbook:SPARC/Portage/Tools/de "Handbook:SPARC/Portage/Tools/de")[Eigener Portage-Tree](/wiki/Handbook:SPARC/Portage/CustomTree/de "Handbook:SPARC/Portage/CustomTree/de")[Erweiterte Portage-Features](/wiki/Handbook:SPARC/Portage/Advanced/de "Handbook:SPARC/Portage/Advanced/de")[Netzwerk-Konfiguration](/wiki/Handbook:SPARC/Full/Networking/de "Handbook:SPARC/Full/Networking/de")[Zu Beginn](/wiki/Handbook:SPARC/Networking/Introduction/de "Handbook:SPARC/Networking/Introduction/de")[Fortgeschrittene Konfiguration](/wiki/Handbook:SPARC/Networking/Advanced/de "Handbook:SPARC/Networking/Advanced/de")[Modulare Vernetzung](/wiki/Handbook:SPARC/Networking/Modular/de "Handbook:SPARC/Networking/Modular/de")[Drahtlose Netzwerke](/wiki/Handbook:SPARC/Networking/Wireless/de "Handbook:SPARC/Networking/Wireless/de")[Funktionalität hinzufügen](/wiki/Handbook:SPARC/Networking/Extending/de "Handbook:SPARC/Networking/Extending/de")[Dynamisches Management](/wiki/Handbook:SPARC/Networking/Dynamic/de "Handbook:SPARC/Networking/Dynamic/de")

## Contents

- [1GRUB](#GRUB)
  - [1.1Emerge](#Emerge)
  - [1.2Installation](#Installation)
    - [1.2.1GPT](#GPT)
    - [1.2.2Sun-Partitionstabelle](#Sun-Partitionstabelle)
  - [1.3Konfiguration](#Konfiguration)
- [2SILO der SPARC Bootloader](#SILO_der_SPARC_Bootloader)
- [3Neustart des Systems](#Neustart_des_Systems)

## GRUB

Wenn ein [64-Bit-Profils ausgewählt wurde](/wiki/Handbook:SPARC/Installation/Base/de#Choose_the_right_profile "Handbook:SPARC/Installation/Base/de"), dann ist [GRUB](/wiki/GRUB/de "GRUB/de") der einzige unterstützte Bootloader.

### Emerge

GRUB sollte auf der Grundlage des Profils automatisch korrekt für die Plattform konfiguriert werden. Um es jedoch explizit zu machen, geben Sie es an:

`root #` `echo 'GRUB_PLATFORMS="ieee1275"' >> /etc/portage/make.conf`

`root #` `emerge --ask --verbose sys-boot/grub`

Die GRUB Software wurde nun zu dem System hinzugefügt. Sie ist aber noch nicht als Bootloader installiert.

### Installation

#### GPT

Wenn die [Festplatte mit GPT partitioniert ist](/wiki/Handbook:SPARC/Installation/Disks/de#Using_fdisk_to_partition_the_disk "Handbook:SPARC/Installation/Disks/de") (die bevorzugte Methode), dann installieren Sie GRUB auf der BIOS-Bootpartition. Unter der Annahme, dass die erste Festplatte (diejenige, von der das System startet) ist, genügen die folgenden Befehle:

`root #` `grub-install --target=sparc64-ieee1275 --recheck /dev/sda`

**Tipp**

Um den Boot-Geräte-String zu finden, den Sie in die Firmware eingeben müssen, verwenden Sie das Werkzeug grub-ofpathname. Wenn die BIOS-Bootpartition die erste Partition auf der Festplatte ist, wählen Sie die gesamte Festplatte aus:

`root #` `grub-ofpathname /dev/sda`

Ansonsten wählen Sie explizit die BIOS-Bootpartition aus:

`root #` `grub-ofpathname /dev/sda2`

#### Sun-Partitionstabelle

Wenn die Festplatte stattdessen mit einer Sun-Partitionstabelle partitioniert ist, muss GRUB mit Hilfe von Blocklisten installiert werden. In diesem Modus geben Sie statt der physischen Festplatte als Argument den Pfad zur Partition an, auf der /boot/grub eingehängt ist.

`root #` `grub-install --target=sparc64-ieee1275 --recheck --force --skip-fs-probe /dev/sda1`

### Konfiguration

Als Nächstes wird die GRUB2-Konfiguration auf der Grundlage der in der Datei /etc/default/grub und den Skripten /etc/grub.d angegebenen Benutzerkonfiguration erstellt. In den meisten Fällen ist keine Konfiguration durch den Benutzer erforderlich, da GRUB2 automatisch erkennt, welcher Kernel zu booten ist (der höchste in /boot verfügbare) und welches das Root-Dateisystem ist. Es ist auch möglich, Kernel-Parameter in /etc/default/grub mithilfe der Variable GRUB\_CMDLINE\_LINUX hinzuzufügen.

Um die endgültige GRUB2-Konfiguration zu erstellen, führen Sie den Befehl grub-mkconfig aus:

`root #` `grub-mkconfig -o /boot/grub/grub.cfg`

```
Generating grub.cfg ...
Found linux image: /boot/vmlinuz-6.19.3-gentoo
Found initrd image: /boot/initramfs-genkernel-sparc-6.19.3-gentoo
done

```

Die Ausgabe des Befehls muss erwähnen, dass mindestens ein Linux Image gefunden wurde, da dieses zum Booten des Systems erforderlich sind. Wenn ein initramfs verwendet wird, oder der Kernel mit Hilfe von genkernel erzeugt wurde, sollte das korrekte initrd Image ebenfalls erkannt werden. Falls dies nicht der Fall ist, überprüfen Sie das Verzeichnis /boot/ mit dem Befehl ls auf dessen Inhalt. Wenn die Dateien in der Tat fehlen sollten, gehen Sie zurück zur Kernel-Konfiguration und der dortigen Installationsanleitung.

## SILO der SPARC Bootloader

Wenn [ein 32-Bit-Profil](/wiki/Handbook:SPARC/Installation/Base/de#Choosing_the_right_profile "Handbook:SPARC/Installation/Base/de") während der Installation ausgewählt wurde, dann ist [SILO](https://git.kernel.org/?p=linux/kernel/git/davem/silo.git;a=summary) (Sparc Improved boot LOader) der einzige unterstützte Bootloader.

`root #` `emerge --ask sys-boot/silo`

Als Nächstes erstellen Sie /etc/silo.conf:

`root #` `nano -w /etc/silo.conf`

Unterhalb wird eine silo.conf Beispieldatei gezeigt. Sie verwendet das Partitions-Schema, das wir im Rahmen dieses Buches verwenden, kernel-6.19.3-gentoo als Kernel Abbild und initramfs-genkernel-sparc64-6.19.3-gentoo als initramfs.

DATEI **`/etc/silo.conf`** **Beispiel Konfigurationsdatei**

```
partition = 1         # Boot Partition (= root Partition)
root = /dev/sda1      # Root Partition
timeout = 150         # Warte 15 Sekunden vor dem Booten des Standardabschnitts

image = /boot/kernel-6.19.3-gentoo
  label = linux
  append = "initrd=/boot/initramfs-genkernel-sparc64-6.19.3-gentoo real_root=/dev/sda1"

```

Falls Sie die silo.conf Beispieldatei verwenden wie sie von Portage ausgeliefert wird, stellen Sie sicher dass Sie alle Zeilen auskommentieren, die nicht benötigt werden.

Wenn die physische Festplatte auf der SILO (als Bootloader) installiert wird eine andere ist, als auf der /etc/silo.conf liegt, dann kopieren Sie zuerst /etc/silo.conf auf eine Partition dieser Festplatte. Wenn /boot/ eine separate Partition auf dieser Festplatte ist, kopieren Sie die Konfigurationsdatei in /boot/ und starten Sie /sbin/silo:

`root #` `cp /etc/silo.conf /boot
`

`root #` `/sbin/silo -C /boot/silo.conf`

```
/boot/silo.conf appears to be valid

```

Andernfalls starten Sie einfach /sbin/silo:

`root #` `/sbin/silo`

```
/etc/silo.conf appears to be valid

```

**Hinweis**

Starten Sie silo jedes mal erneut nach der Aktualisierung oder Installation des Paketes [sys-boot/silo](https://packages.gentoo.org/packages/sys-boot/silo) (falls notwendig mit Parametern).

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

Nach dem Neustart in die neu installierte Gentoo Umgebung können Sie Ihre Installation mit dem Kapitel [Abschluss der Gentoo Installation](/wiki/Handbook:SPARC/Installation/Finalizing/de "Handbook:SPARC/Installation/Finalizing/de") fertigstellen.

[← Installation der Tools](/wiki/Handbook:SPARC/Installation/Tools/de "Previous part") [Anfang](/wiki/Handbook:SPARC/de "Handbook:SPARC/de") [Abschluss der Installation →](/wiki/Handbook:SPARC/Installation/Finalizing/de "Next part")