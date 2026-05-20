# Bootloader

Other languages:

- Deutsch
- [English](/wiki/Handbook:Alpha/Installation/Bootloader "Handbook:Alpha/Installation/Bootloader (100% translated)")
- [español](/wiki/Handbook:Alpha/Installation/Bootloader/es "Manual de Gentoo: Alpha/Instalación/Arranque (100% translated)")
- [français](/wiki/Handbook:Alpha/Installation/Bootloader/fr "Manuel:Alpha/Installation/Système d'amorçage (100% translated)")
- [italiano](/wiki/Handbook:Alpha/Installation/Bootloader/it "Handbook:Alpha/Installation/Bootloader/it (100% translated)")
- [magyar](/wiki/Handbook:Alpha/Installation/Bootloader/hu "Handbook:Alpha/Installation/Bootloader/hu (100% translated)")
- [polski](/wiki/Handbook:Alpha/Installation/Bootloader/pl "Handbook:Alpha/Installation/Bootloader (100% translated)")
- [português do Brasil](/wiki/Handbook:Alpha/Installation/Bootloader/pt-br "Manual:Alpha/Instalação/Gerenciador de boot (100% translated)")
- [čeština](/wiki/Handbook:Alpha/Installation/Bootloader/cs "Handbook:Alpha/Installation/Bootloader/cs (50% translated)")
- [русский](/wiki/Handbook:Alpha/Installation/Bootloader/ru "Handbook:Alpha/Installation/Bootloader (100% translated)")
- [தமிழ்](/wiki/Handbook:Alpha/Installation/Bootloader/ta "கையேடு:Alpha/நிறுவல்/துவக்கஏற்றி (100% translated)")
- [中文（中国大陆）‎](/wiki/Handbook:Alpha/Installation/Bootloader/zh-cn "手册：Alpha/安装/配置系统引导程序Bootloader (100% translated)")
- [日本語](/wiki/Handbook:Alpha/Installation/Bootloader/ja "ハンドブック:Alpha/インストール/ブートローダー (100% translated)")
- [한국어](/wiki/Handbook:Alpha/Installation/Bootloader/ko "Handbook:Alpha/Installation/Bootloader/ko (100% translated)")

[Alpha Handbuch](/wiki/Handbook:Alpha/de "Handbook:Alpha/de")[Installation](/wiki/Handbook:Alpha/Full/Installation/de "Handbook:Alpha/Full/Installation/de")[Über die Installation](/wiki/Handbook:Alpha/Installation/About/de "Handbook:Alpha/Installation/About/de")[Auswahl des Mediums](/wiki/Handbook:Alpha/Installation/Media/de "Handbook:Alpha/Installation/Media/de")[Konfiguration des Netzwerks](/wiki/Handbook:Alpha/Installation/Networking/de "Handbook:Alpha/Installation/Networking/de")[Vorbereiten der Festplatte(n)](/wiki/Handbook:Alpha/Installation/Disks/de "Handbook:Alpha/Installation/Disks/de")[Installation des Stage Archivs](/wiki/Handbook:Alpha/Installation/Stage/de "Handbook:Alpha/Installation/Stage/de")[Installation des Basissystems](/wiki/Handbook:Alpha/Installation/Base/de "Handbook:Alpha/Installation/Base/de")[Konfiguration des Kernels](/wiki/Handbook:Alpha/Installation/Kernel/de "Handbook:Alpha/Installation/Kernel/de")[Konfiguration des Systems](/wiki/Handbook:Alpha/Installation/System/de "Handbook:Alpha/Installation/System/de")[Installation der Tools](/wiki/Handbook:Alpha/Installation/Tools/de "Handbook:Alpha/Installation/Tools/de")Konfiguration des Bootloaders[Abschluss](/wiki/Handbook:Alpha/Installation/Finalizing/de "Handbook:Alpha/Installation/Finalizing/de")[Arbeiten mit Gentoo](/wiki/Handbook:Alpha/Full/Working/de "Handbook:Alpha/Full/Working/de")[Portage-Einführung](/wiki/Handbook:Alpha/Working/Portage/de "Handbook:Alpha/Working/Portage/de")[USE-Flags](/wiki/Handbook:Alpha/Working/USE/de "Handbook:Alpha/Working/USE/de")[Portage-Features](/wiki/Handbook:Alpha/Working/Features/de "Handbook:Alpha/Working/Features/de")[Initskript-System](/wiki/Handbook:Alpha/Working/Initscripts/de "Handbook:Alpha/Working/Initscripts/de")[Umgebungsvariablen](/wiki/Handbook:Alpha/Working/EnvVar/de "Handbook:Alpha/Working/EnvVar/de")[Arbeiten mit Portage](/wiki/Handbook:Alpha/Full/Portage/de "Handbook:Alpha/Full/Portage/de")[Dateien und Verzeichnisse](/wiki/Handbook:Alpha/Portage/Files/de "Handbook:Alpha/Portage/Files/de")[Variablen](/wiki/Handbook:Alpha/Portage/Variables/de "Handbook:Alpha/Portage/Variables/de")[Mischen von Softwarezweigen](/wiki/Handbook:Alpha/Portage/Branches/de "Handbook:Alpha/Portage/Branches/de")[Zusätzliche Tools](/wiki/Handbook:Alpha/Portage/Tools/de "Handbook:Alpha/Portage/Tools/de")[Eigener Portage-Tree](/wiki/Handbook:Alpha/Portage/CustomTree/de "Handbook:Alpha/Portage/CustomTree/de")[Erweiterte Portage-Features](/wiki/Handbook:Alpha/Portage/Advanced/de "Handbook:Alpha/Portage/Advanced/de")[Netzwerk-Konfiguration](/wiki/Handbook:Alpha/Full/Networking/de "Handbook:Alpha/Full/Networking/de")[Zu Beginn](/wiki/Handbook:Alpha/Networking/Introduction/de "Handbook:Alpha/Networking/Introduction/de")[Fortgeschrittene Konfiguration](/wiki/Handbook:Alpha/Networking/Advanced/de "Handbook:Alpha/Networking/Advanced/de")[Modulare Vernetzung](/wiki/Handbook:Alpha/Networking/Modular/de "Handbook:Alpha/Networking/Modular/de")[Drahtlose Netzwerke](/wiki/Handbook:Alpha/Networking/Wireless/de "Handbook:Alpha/Networking/Wireless/de")[Funktionalität hinzufügen](/wiki/Handbook:Alpha/Networking/Extending/de "Handbook:Alpha/Networking/Extending/de")[Dynamisches Management](/wiki/Handbook:Alpha/Networking/Dynamic/de "Handbook:Alpha/Networking/Dynamic/de")

## Triff eine Wahl

Da der Kernel nun konfiguriert und kompiliert ist und die notwendigen Systemkonfigurationsdateien korrekt ausgefüllt sind, ist es an der Zeit ein Programm zu installieren, das den Kernel beim Einschalten des Systems startet. Solch ein Programm nennt man Bootloader.

Es gibt verschiedene Booloader für Linux/Alpha. Wählen Sie _einen_ der unterstützten Bootloader, nicht alle. Wir beschreiben [aBoot](#Der_Standard:_aBoot) und [MILO](#Die_Alternative:_MILO).

## Der Standard: aBoot

**Hinweis**

aboot unterstützt das Booten nur von ext2 und ext3 Partitionen.

Installieren Sie als erstes aboot auf dem System:

`root #` `emerge --ask sys-boot/aboot`

Der nächste Schritt ist die Bootplatte bootbar zu machen. Dies wird aboot beim Booten des Systems starten. Wir machen unsere Bootplatte bootbar indem wir den Bootloader aboot an den Anfang der Platte schreiben:

`root #` `swriteboot -f3 /dev/sda /boot/bootlx
`

`root #` `abootconf /dev/sda 2`

**Hinweis**

Falls Sie ein anderes Partitionierungsschema verwenden als das in diesem Kapitel verwendete, dann müssen Sie die Befehle entsprechend abändern. Lesen Sie bitte die entsprechenden Manpages (man 8 swriteboot und man 8 abootconf). Wenn das Root-Dateisystem unter dem JFS Dateisystem läuft gehen Sie bitte auch sicher, dass es zu Beginn als nur lesbar (read-only) durch Hinzufügen von ro als Kerneloption gemounted wird.

Obwohl aboot jetzt installiert ist, müssen wir noch eine Konfigurationsdatei dafür schreiben. aboot braucht nur eine Zeile für jede Konfiguration, deshalb können wir folgendes tun:

`root #` `echo '0:2/boot/vmlinux.gz root=/dev/sda3' > /etc/aboot.conf`

Wenn beim Bau des Linux-Kernel ebenfalls ein initramfs erzeugt wurde, dann ist es notwendig die Konfiguration durch Referenzierung auf diese initramfs-Datei zu ändern. Außerdem muss initramfs mitgeteilt werden, an welcher Stelle sich das echte Root-Laufwerk befindet:

`root #` `echo '0:2/boot/vmlinux.gz initrd=/boot/initramfs-genkernel-alpha-6.19.1-gentoo real_root=/dev/sda3' > /etc/aboot.conf`

Darüber hinaus ist es möglich, dass Gentoo durch das Setzten einiger SRM Variablen automatisch bootet. Versuchen Sie diese Variablen von Linux aus zu setzten, aber es könnte einfacher sein dies von der SRM Konsole selbst zu tun.

`root #` `cd /proc/srm_environment/named_variables
`

`root #` `echo -n 0 > boot_osflags
`

`root #` `echo -n '' > boot_file
`

`root #` `echo -n 'BOOT' > auto_action
`

`root #` `echo -n 'dkc100' > bootdef_dev`

Natürlich müssen Sie dkc100 durch das ersetzten, was auch immer das Boot-Laufwerk ist.

Um in der Zukunft erneut in die SRM Konsole zu kommen (um die Gentoo-Installation wiederherzustellen, mit einigen Variablen zu spielen, oder was auch immer), drücken Sie zum Abbruch des automatischen Ladeprozess einfach `Ctrl+C`.

Vergessen Sie bei der Installation unter der seriellen Konsole nicht das Boot-Flag der seriellen Konsole in aboot.conf einzubinden. Siehe /etc/aboot.conf.example für weitere Informationen.

aboot ist jetzt konfiguriert und betriebsbereit. Fahren Sie fort mit dem [Neustart des Systems](#Rebooting_the_system).

Fahren Sie nun mit dem [Neustart des Systems](#Rebooting_the_system) fort.

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

Nach dem Neustart in die neu installierte Gentoo Umgebung können Sie Ihre Installation mit dem Kapitel [Abschluss der Gentoo Installation](/wiki/Handbook:Alpha/Installation/Finalizing/de "Handbook:Alpha/Installation/Finalizing/de") fertigstellen.

[← Installation der Tools](/wiki/Handbook:Alpha/Installation/Tools/de "Previous part") [Anfang](/wiki/Handbook:Alpha/de "Handbook:Alpha/de") [Abschluss der Installation →](/wiki/Handbook:Alpha/Installation/Finalizing/de "Next part")