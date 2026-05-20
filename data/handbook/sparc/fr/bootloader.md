# Bootloader

Other languages:

- [Deutsch](/wiki/Handbook:SPARC/Installation/Bootloader/de "Handbuch:SPARC/Installation/Bootloader (100% translated)")
- [English](/wiki/Handbook:SPARC/Installation/Bootloader "Handbook:SPARC/Installation/Bootloader (100% translated)")
- [español](/wiki/Handbook:SPARC/Installation/Bootloader/es "Manual de Gentoo: SPARC/Instalación/Arranque (100% translated)")
- français
- [italiano](/wiki/Handbook:SPARC/Installation/Bootloader/it "Handbook:SPARC/Installation/Bootloader/it (50% translated)")
- [magyar](/wiki/Handbook:SPARC/Installation/Bootloader/hu "Handbook:SPARC/Installation/Bootloader/hu (100% translated)")
- [polski](/wiki/Handbook:SPARC/Installation/Bootloader/pl "Handbook:SPARC/Installation/Bootloader (100% translated)")
- [português do Brasil](/wiki/Handbook:SPARC/Installation/Bootloader/pt-br "Handbook:SPARC/Installation/Bootloader/pt-br (50% translated)")
- [русский](/wiki/Handbook:SPARC/Installation/Bootloader/ru "Handbook:SPARC/Installation/Bootloader (100% translated)")
- [தமிழ்](/wiki/Handbook:SPARC/Installation/Bootloader/ta "கையேடு:SPARC/நிறுவல்/துவக்கஏற்றி (100% translated)")
- [中文（中国大陆）‎](/wiki/Handbook:SPARC/Installation/Bootloader/zh-cn "手册：SPARC/安装/配置系统引导程序Bootloader (100% translated)")
- [日本語](/wiki/Handbook:SPARC/Installation/Bootloader/ja "ハンドブック:SPARC/インストール/ブートローダー (100% translated)")
- [한국어](/wiki/Handbook:SPARC/Installation/Bootloader/ko "Handbook:SPARC/Installation/Bootloader/ko (100% translated)")

[Sommaire du manuel](/wiki/Handbook:SPARC "Handbook:SPARC")[Installation](/wiki/Handbook:SPARC/Full/Installation/fr "Handbook:SPARC/Full/Installation/fr")[À propos de l’installation](/wiki/Handbook:SPARC/Installation/About/fr "Handbook:SPARC/Installation/About/fr")[Choix du support](/wiki/Handbook:SPARC/Installation/Media/fr "Handbook:SPARC/Installation/Media/fr")[Configurer le réseau](/wiki/Handbook:SPARC/Installation/Networking/fr "Handbook:SPARC/Installation/Networking/fr")[Préparer les disques](/wiki/Handbook:SPARC/Installation/Disks/fr "Handbook:SPARC/Installation/Disks/fr")[Installer l’archive _stage3_](/wiki/Handbook:SPARC/Installation/Stage/fr "Handbook:SPARC/Installation/Stage/fr")[Installer le système de base](/wiki/Handbook:SPARC/Installation/Base/fr "Handbook:SPARC/Installation/Base/fr")[Configurer le noyau](/wiki/Handbook:SPARC/Installation/Kernel/fr "Handbook:SPARC/Installation/Kernel/fr")[Configurer le système](/wiki/Handbook:SPARC/Installation/System/fr "Handbook:SPARC/Installation/System/fr")[Installer les outils](/wiki/Handbook:SPARC/Installation/Tools/fr "Handbook:SPARC/Installation/Tools/fr")Configurer le système d’amorçage[Finaliser](/wiki/Handbook:SPARC/Installation/Finalizing/fr "Handbook:SPARC/Installation/Finalizing/fr")[Utiliser Gentoo](/wiki/Handbook:SPARC/Full/Working/fr "Handbook:SPARC/Full/Working/fr")[Introduction à Portage](/wiki/Handbook:SPARC/Working/Portage/fr "Handbook:SPARC/Working/Portage/fr")[Les options de la variable USE](/wiki/Handbook:SPARC/Working/USE/fr "Handbook:SPARC/Working/USE/fr")[Les fonctionnalités de Portage](/wiki/Handbook:SPARC/Working/Features/fr "Handbook:SPARC/Working/Features/fr")[Scripts d’initialisation systèmes](/wiki/Handbook:SPARC/Working/Initscripts/fr "Handbook:SPARC/Working/Initscripts/fr")[Variables d’environnement](/wiki/Handbook:SPARC/Working/EnvVar/fr "Handbook:SPARC/Working/EnvVar/fr")[Utiliser Portage](/wiki/Handbook:SPARC/Full/Portage/fr "Handbook:SPARC/Full/Portage/fr")[Fichiers et répertoires](/wiki/Handbook:SPARC/Portage/Files/fr "Handbook:SPARC/Portage/Files/fr")[Les variables](/wiki/Handbook:SPARC/Portage/Variables/fr "Handbook:SPARC/Portage/Variables/fr")[Mélanger plusieurs branches logicielles](/wiki/Handbook:SPARC/Portage/Branches/fr "Handbook:SPARC/Portage/Branches/fr")[Outils supplémentaires](/wiki/Handbook:SPARC/Portage/Tools/fr "Handbook:SPARC/Portage/Tools/fr")[Dépôt personnalisé](/wiki/Handbook:SPARC/Portage/CustomTree/fr "Handbook:SPARC/Portage/CustomTree/fr")[Fonctionnalités avancées](/wiki/Handbook:SPARC/Portage/Advanced/fr "Handbook:SPARC/Portage/Advanced/fr")[Configuration du réseau avec OpenRC](/wiki/Handbook:SPARC/Full/Networking/fr "Handbook:SPARC/Full/Networking/fr")[Bien démarrer](/wiki/Handbook:SPARC/Networking/Introduction/fr "Handbook:SPARC/Networking/Introduction/fr")[Configuration avancée](/wiki/Handbook:SPARC/Networking/Advanced/fr "Handbook:SPARC/Networking/Advanced/fr")[Les modules réseau](/wiki/Handbook:SPARC/Networking/Modular/fr "Handbook:SPARC/Networking/Modular/fr")[Sans fil](/wiki/Handbook:SPARC/Networking/Wireless/fr "Handbook:SPARC/Networking/Wireless/fr")[Ajouter des fonctionnalités](/wiki/Handbook:SPARC/Networking/Extending/fr "Handbook:SPARC/Networking/Extending/fr")[Gestion dynamique](/wiki/Handbook:SPARC/Networking/Dynamic/fr "Handbook:SPARC/Networking/Dynamic/fr")

[Handbook:SPARC/Blocks/Bootloader/fr](/index.php?title=Handbook:SPARC/Blocks/Bootloader/fr&action=edit&redlink=1 "Handbook:SPARC/Blocks/Bootloader/fr (page does not exist)")

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

Une fois redémarré dans le nouvel environnement Gentoo, il est avisé de finir avec [Finalisation de l’installation de Gentoo](/wiki/Handbook:SPARC/Installation/Finalizing/fr "Handbook:SPARC/Installation/Finalizing/fr").

[← Installing tools](/wiki/Handbook:SPARC/Installation/Tools/fr "Previous part") [Accueil](/wiki/Handbook:SPARC/fr "Handbook:SPARC/fr") [Finalisation →](/wiki/Handbook:SPARC/Installation/Finalizing/fr "Next part")