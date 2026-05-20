# Bootloader

Other languages:

- [Deutsch](/wiki/Handbook:Alpha/Installation/Bootloader/de "Handbuch:Alpha/Installation/Bootloader (100% translated)")
- [English](/wiki/Handbook:Alpha/Installation/Bootloader "Handbook:Alpha/Installation/Bootloader (100% translated)")
- [español](/wiki/Handbook:Alpha/Installation/Bootloader/es "Manual de Gentoo: Alpha/Instalación/Arranque (100% translated)")
- français
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

[Sommaire du manuel](/wiki/Handbook:Alpha "Handbook:Alpha")[Installation](/wiki/Handbook:Alpha/Full/Installation/fr "Handbook:Alpha/Full/Installation/fr")[À propos de l’installation](/wiki/Handbook:Alpha/Installation/About/fr "Handbook:Alpha/Installation/About/fr")[Choix du support](/wiki/Handbook:Alpha/Installation/Media/fr "Handbook:Alpha/Installation/Media/fr")[Configurer le réseau](/wiki/Handbook:Alpha/Installation/Networking/fr "Handbook:Alpha/Installation/Networking/fr")[Préparer les disques](/wiki/Handbook:Alpha/Installation/Disks/fr "Handbook:Alpha/Installation/Disks/fr")[Installer l’archive _stage3_](/wiki/Handbook:Alpha/Installation/Stage/fr "Handbook:Alpha/Installation/Stage/fr")[Installer le système de base](/wiki/Handbook:Alpha/Installation/Base/fr "Handbook:Alpha/Installation/Base/fr")[Configurer le noyau](/wiki/Handbook:Alpha/Installation/Kernel/fr "Handbook:Alpha/Installation/Kernel/fr")[Configurer le système](/wiki/Handbook:Alpha/Installation/System/fr "Handbook:Alpha/Installation/System/fr")[Installer les outils](/wiki/Handbook:Alpha/Installation/Tools/fr "Handbook:Alpha/Installation/Tools/fr")Configurer le système d’amorçage[Finaliser](/wiki/Handbook:Alpha/Installation/Finalizing/fr "Handbook:Alpha/Installation/Finalizing/fr")[Utiliser Gentoo](/wiki/Handbook:Alpha/Full/Working/fr "Handbook:Alpha/Full/Working/fr")[Introduction à Portage](/wiki/Handbook:Alpha/Working/Portage/fr "Handbook:Alpha/Working/Portage/fr")[Les options de la variable USE](/wiki/Handbook:Alpha/Working/USE/fr "Handbook:Alpha/Working/USE/fr")[Les fonctionnalités de Portage](/wiki/Handbook:Alpha/Working/Features/fr "Handbook:Alpha/Working/Features/fr")[Scripts d’initialisation systèmes](/wiki/Handbook:Alpha/Working/Initscripts/fr "Handbook:Alpha/Working/Initscripts/fr")[Variables d’environnement](/wiki/Handbook:Alpha/Working/EnvVar/fr "Handbook:Alpha/Working/EnvVar/fr")[Utiliser Portage](/wiki/Handbook:Alpha/Full/Portage/fr "Handbook:Alpha/Full/Portage/fr")[Fichiers et répertoires](/wiki/Handbook:Alpha/Portage/Files/fr "Handbook:Alpha/Portage/Files/fr")[Les variables](/wiki/Handbook:Alpha/Portage/Variables/fr "Handbook:Alpha/Portage/Variables/fr")[Mélanger plusieurs branches logicielles](/wiki/Handbook:Alpha/Portage/Branches/fr "Handbook:Alpha/Portage/Branches/fr")[Outils supplémentaires](/wiki/Handbook:Alpha/Portage/Tools/fr "Handbook:Alpha/Portage/Tools/fr")[Dépôt personnalisé](/wiki/Handbook:Alpha/Portage/CustomTree/fr "Handbook:Alpha/Portage/CustomTree/fr")[Fonctionnalités avancées](/wiki/Handbook:Alpha/Portage/Advanced/fr "Handbook:Alpha/Portage/Advanced/fr")[Configuration du réseau avec OpenRC](/wiki/Handbook:Alpha/Full/Networking/fr "Handbook:Alpha/Full/Networking/fr")[Bien démarrer](/wiki/Handbook:Alpha/Networking/Introduction/fr "Handbook:Alpha/Networking/Introduction/fr")[Configuration avancée](/wiki/Handbook:Alpha/Networking/Advanced/fr "Handbook:Alpha/Networking/Advanced/fr")[Les modules réseau](/wiki/Handbook:Alpha/Networking/Modular/fr "Handbook:Alpha/Networking/Modular/fr")[Sans fil](/wiki/Handbook:Alpha/Networking/Wireless/fr "Handbook:Alpha/Networking/Wireless/fr")[Ajouter des fonctionnalités](/wiki/Handbook:Alpha/Networking/Extending/fr "Handbook:Alpha/Networking/Extending/fr")[Gestion dynamique](/wiki/Handbook:Alpha/Networking/Dynamic/fr "Handbook:Alpha/Networking/Dynamic/fr")

## Making a choice

Now that the kernel is configured and compiled and the necessary system configuration files are filled in correctly, it is time to install a program that will fire up the kernel when the system is started. Such a program is called a bootloader.

Several bootloaders exist for Linux/Alpha. Choose one of the supported bootloaders, not all. We document [aBoot](#Default:_Using_aboot) and [MILO](#Alternative:_Using_MILO).

## Par défaut : Utiliser aBoot

**Remarque**

aboot only supports booting from ext2 and ext3 partitions.

First install aboot on the system

`root #` `emerge --ask sys-boot/aboot`

The next step is to make the bootdisk bootable. This will start aboot when booting the system. We make our bootdisk bootable by writing the aboot bootloader to the start of the disk.

`root #` `swriteboot -f3 /dev/sda /boot/bootlx
`

`root #` `abootconf /dev/sda 2
`

**Remarque**

When using a different partitioning scheme than the one used throughout this chapter, the commands need to be changed accordingly. Please read the appropriate manual pages (man 8 swriteboot and man 8 abootconf). Also, if the root filesystem is ran using the JFS filesystem, make sure it gets mounted read-only at first by adding ro as a kernel option.

Although aboot is now installed, we still need to write a configuration file for it. aboot only requires one line for each configuration, so we can do this:

`root #` `echo '0:2/boot/vmlinux.gz root=/dev/sda3' > /etc/aboot.conf`

If, while building the Linux kernel, an initramfs was build as well to boot from, then it is necessary to change the configuration by referring to this initramfs file and telling the initramfs where the real root device is at:

`root #` `echo '0:2/boot/vmlinux.gz initrd=/boot/initramfs-genkernel-alpha-6.19.1-gentoo root=/dev/sda3' > /etc/aboot.conf`

Additionally, it is possible to make Gentoo boot automatically by setting up some SRM variables. Try setting these variables from Linux, but it may be easier to do so from the SRM console itself.

`root #` `cd /proc/srm_environment/named_variables
`

`root #` `echo -n 0 > boot_osflags
`

`root #` `echo -n '' > boot_file
`

`root #` `echo -n 'BOOT' > auto_action
`

`root #` `echo -n 'dkc100' > bootdef_dev`

Of course substitute dkc100 with whatever the boot device is.

To get in the SRM console again in the future (to recover the Gentoo install, play with some variables, or whatever), just hit `Ctrl+C` to abort the automatic loading process.

When installing using a serial console, don't forget to include the serial console boot flag in aboot.conf. See /etc/aboot.conf.example for some further information.

Aboot is now configured and ready to use. Continue with [Rebooting the system](#Rebooting_the_system).

Now continue with [Rebooting the system](#Rebooting_the_system).

## Redémarrer le système

Quittez l’environnement et démontez toutes les partitions montées. Ensuite, exécutez cette commande magique qui lance le vrai test final : reboot.

`(chroot) livecd #` `exit`

`livecd~#` `cd
`

`livecd~#` `umount -l /mnt/gentoo/dev{/shm,/pts,}
`

`livecd~#` `umount -R /mnt/gentoo
`

`livecd~#` `reboot`

N’oubliez pas de retirer l’image _live_, sinon il pourrait être redémarré à la place du nouveau système Gentoo.

Une fois redémarré dans le nouvel environnement Gentoo, il est avisé de finir avec [Finalisation de l’installation de Gentoo](/wiki/Handbook:Alpha/Installation/Finalizing/fr "Handbook:Alpha/Installation/Finalizing/fr").

[← Installing tools](/wiki/Handbook:Alpha/Installation/Tools/fr "Previous part") [Accueil](/wiki/Handbook:Alpha/fr "Handbook:Alpha/fr") [Finalisation →](/wiki/Handbook:Alpha/Installation/Finalizing/fr "Next part")