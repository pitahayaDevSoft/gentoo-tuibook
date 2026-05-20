# Bootloader

Other languages:

- Deutsch
- [English](/wiki/Handbook:HPPA/Installation/Bootloader "Handbook:HPPA/Installation/Bootloader (100% translated)")
- [español](/wiki/Handbook:HPPA/Installation/Bootloader/es "Manual de Gentoo: HPPA/Instalación/Arranque (100% translated)")
- [français](/wiki/Handbook:HPPA/Installation/Bootloader/fr "Handbook:HPPA/Installation/Bootloader/fr (100% translated)")
- [italiano](/wiki/Handbook:HPPA/Installation/Bootloader/it "Handbook:HPPA/Installation/Bootloader/it (100% translated)")
- [magyar](/wiki/Handbook:HPPA/Installation/Bootloader/hu "Handbook:HPPA/Installation/Bootloader/hu (100% translated)")
- [polski](/wiki/Handbook:HPPA/Installation/Bootloader/pl "Handbook:HPPA/Installation/Bootloader (100% translated)")
- [português do Brasil](/wiki/Handbook:HPPA/Installation/Bootloader/pt-br "Manual:HPPA/Instalação/Gerenciador de Boot (100% translated)")
- [русский](/wiki/Handbook:HPPA/Installation/Bootloader/ru "Handbook:HPPA/Installation/Bootloader (100% translated)")
- [தமிழ்](/wiki/Handbook:HPPA/Installation/Bootloader/ta "கையேடு:HPPA/நிறுவல்/துவக்கஏற்றி (100% translated)")
- [中文（中国大陆）‎](/wiki/Handbook:HPPA/Installation/Bootloader/zh-cn "手册：HPPA/安装/配置系统引导程序Bootloader (100% translated)")
- [日本語](/wiki/Handbook:HPPA/Installation/Bootloader/ja "ハンドブック:HPPA/インストール/ブートローダー (100% translated)")
- [한국어](/wiki/Handbook:HPPA/Installation/Bootloader/ko "Handbook:HPPA/Installation/Bootloader/ko (100% translated)")

[HPPA Handbuch](/wiki/Handbook:HPPA/de "Handbook:HPPA/de")[Installation](/wiki/Handbook:HPPA/Full/Installation/de "Handbook:HPPA/Full/Installation/de")[Über die Installation](/wiki/Handbook:HPPA/Installation/About/de "Handbook:HPPA/Installation/About/de")[Auswahl des Mediums](/wiki/Handbook:HPPA/Installation/Media/de "Handbook:HPPA/Installation/Media/de")[Konfiguration des Netzwerks](/wiki/Handbook:HPPA/Installation/Networking/de "Handbook:HPPA/Installation/Networking/de")[Vorbereiten der Festplatte(n)](/wiki/Handbook:HPPA/Installation/Disks/de "Handbook:HPPA/Installation/Disks/de")[Installation des Stage Archivs](/wiki/Handbook:HPPA/Installation/Stage/de "Handbook:HPPA/Installation/Stage/de")[Installation des Basissystems](/wiki/Handbook:HPPA/Installation/Base/de "Handbook:HPPA/Installation/Base/de")[Konfiguration des Kernels](/wiki/Handbook:HPPA/Installation/Kernel/de "Handbook:HPPA/Installation/Kernel/de")[Konfiguration des Systems](/wiki/Handbook:HPPA/Installation/System/de "Handbook:HPPA/Installation/System/de")[Installation der Tools](/wiki/Handbook:HPPA/Installation/Tools/de "Handbook:HPPA/Installation/Tools/de")Konfiguration des Bootloaders[Abschluss](/wiki/Handbook:HPPA/Installation/Finalizing/de "Handbook:HPPA/Installation/Finalizing/de")[Arbeiten mit Gentoo](/wiki/Handbook:HPPA/Full/Working/de "Handbook:HPPA/Full/Working/de")[Portage-Einführung](/wiki/Handbook:HPPA/Working/Portage/de "Handbook:HPPA/Working/Portage/de")[USE-Flags](/wiki/Handbook:HPPA/Working/USE/de "Handbook:HPPA/Working/USE/de")[Portage-Features](/wiki/Handbook:HPPA/Working/Features/de "Handbook:HPPA/Working/Features/de")[Initskript-System](/wiki/Handbook:HPPA/Working/Initscripts/de "Handbook:HPPA/Working/Initscripts/de")[Umgebungsvariablen](/wiki/Handbook:HPPA/Working/EnvVar/de "Handbook:HPPA/Working/EnvVar/de")[Arbeiten mit Portage](/wiki/Handbook:HPPA/Full/Portage/de "Handbook:HPPA/Full/Portage/de")[Dateien und Verzeichnisse](/wiki/Handbook:HPPA/Portage/Files/de "Handbook:HPPA/Portage/Files/de")[Variablen](/wiki/Handbook:HPPA/Portage/Variables/de "Handbook:HPPA/Portage/Variables/de")[Mischen von Softwarezweigen](/wiki/Handbook:HPPA/Portage/Branches/de "Handbook:HPPA/Portage/Branches/de")[Zusätzliche Tools](/wiki/Handbook:HPPA/Portage/Tools/de "Handbook:HPPA/Portage/Tools/de")[Eigener Portage-Tree](/wiki/Handbook:HPPA/Portage/CustomTree/de "Handbook:HPPA/Portage/CustomTree/de")[Erweiterte Portage-Features](/wiki/Handbook:HPPA/Portage/Advanced/de "Handbook:HPPA/Portage/Advanced/de")[Netzwerk-Konfiguration](/wiki/Handbook:HPPA/Full/Networking/de "Handbook:HPPA/Full/Networking/de")[Zu Beginn](/wiki/Handbook:HPPA/Networking/Introduction/de "Handbook:HPPA/Networking/Introduction/de")[Fortgeschrittene Konfiguration](/wiki/Handbook:HPPA/Networking/Advanced/de "Handbook:HPPA/Networking/Advanced/de")[Modulare Vernetzung](/wiki/Handbook:HPPA/Networking/Modular/de "Handbook:HPPA/Networking/Modular/de")[Drahtlose Netzwerke](/wiki/Handbook:HPPA/Networking/Wireless/de "Handbook:HPPA/Networking/Wireless/de")[Funktionalität hinzufügen](/wiki/Handbook:HPPA/Networking/Extending/de "Handbook:HPPA/Networking/Extending/de")[Dynamisches Management](/wiki/Handbook:HPPA/Networking/Dynamic/de "Handbook:HPPA/Networking/Dynamic/de")

## PALO installieren

Auf der PA-RISC Plattform heißt der Bootloader palo. Installieren Sie zuerst die Software auf dem System:

`root #` `emerge --ask sys-boot/palo`

Die Konfiguration ist unter /etc/palo.conf abgelegt. Es folgt eine Beispielkonfiguration:

**Warnung**

This configuration **must** be changed after running palo for the first time! See below for after first setup.

DATEI **`/etc/palo.conf`** **Einfaches PALO Konfigurationsbeispiel**

```
--commandline=2/kernel-6.19.1-gentoo root=/dev/sda4
--recoverykernel=/vmlinux.old
--init-partitioned=/dev/sda

```

Die erste Zeile gibt palo die Lage des Kernels und die zu verwendenden Bootparameter an. Die Zeichenfolge `2/kernel-6.19.1-gentoo` bedeutet, dass der Kernel mit dem Namen /kernel-6.19.1-gentoo auf der zweiten Partition liegt. Vorsicht, der Pfad zum Kernel ist relativ zur Boot Partition und nicht zur Root Partition.

Die zweite Zeile gibt an, welcher Notfall-Kernel (recovery kernel) zu verwenden ist. Wenn es die erste Installation ist und es (noch) keinen Notfall-Kernel gibt, kommentieren Sie dies bitte aus. Die dritte Zeile gibt an, auf welcher Festplatte sich palo befinden wird.

To format the disk, palo must be run with certain arguments. This example uses _ext4_ for the first partition:

`root #` `palo --format-as=4 --init-partitioned=/dev/sda`

Wenn die Konfiguration abgeschlossen ist, führen Sie palo einfach aus.

`root #` `palo`

The configuration must then be updated for post-first-install use:

DATEI **`/etc/palo.conf`** **Simple PALO configuration example**

```
--commandline=2/kernel-6.19.1-gentoo root=/dev/sda4
--recoverykernel=/vmlinux.old
# Don't throw away the old partition, just update the existing one on `palo` runs.
--update-partitioned=/dev/sda
# --format-as has two meanings depending on whether --init-partitioned or --update-partitioned is used. Keep this line.
--format-as=4

```

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

Nach dem Neustart in die neu installierte Gentoo Umgebung können Sie Ihre Installation mit dem Kapitel [Abschluss der Gentoo Installation](/wiki/Handbook:HPPA/Installation/Finalizing/de "Handbook:HPPA/Installation/Finalizing/de") fertigstellen.

[← Installation der Tools](/wiki/Handbook:HPPA/Installation/Tools/de "Previous part") [Anfang](/wiki/Handbook:HPPA/de "Handbook:HPPA/de") [Abschluss der Installation →](/wiki/Handbook:HPPA/Installation/Finalizing/de "Next part")