# Bootloader

Other languages:

- [Deutsch](/wiki/Handbook:HPPA/Installation/Bootloader/de "Handbuch:HPPA/Installation/Bootloader (100% translated)")
- [English](/wiki/Handbook:HPPA/Installation/Bootloader "Handbook:HPPA/Installation/Bootloader (100% translated)")
- [español](/wiki/Handbook:HPPA/Installation/Bootloader/es "Manual de Gentoo: HPPA/Instalación/Arranque (100% translated)")
- français
- [italiano](/wiki/Handbook:HPPA/Installation/Bootloader/it "Handbook:HPPA/Installation/Bootloader/it (100% translated)")
- [magyar](/wiki/Handbook:HPPA/Installation/Bootloader/hu "Handbook:HPPA/Installation/Bootloader/hu (100% translated)")
- [polski](/wiki/Handbook:HPPA/Installation/Bootloader/pl "Handbook:HPPA/Installation/Bootloader (100% translated)")
- [português do Brasil](/wiki/Handbook:HPPA/Installation/Bootloader/pt-br "Manual:HPPA/Instalação/Gerenciador de Boot (100% translated)")
- [русский](/wiki/Handbook:HPPA/Installation/Bootloader/ru "Handbook:HPPA/Installation/Bootloader (100% translated)")
- [தமிழ்](/wiki/Handbook:HPPA/Installation/Bootloader/ta "கையேடு:HPPA/நிறுவல்/துவக்கஏற்றி (100% translated)")
- [中文（中国大陆）‎](/wiki/Handbook:HPPA/Installation/Bootloader/zh-cn "手册：HPPA/安装/配置系统引导程序Bootloader (100% translated)")
- [日本語](/wiki/Handbook:HPPA/Installation/Bootloader/ja "ハンドブック:HPPA/インストール/ブートローダー (100% translated)")
- [한국어](/wiki/Handbook:HPPA/Installation/Bootloader/ko "Handbook:HPPA/Installation/Bootloader/ko (100% translated)")

[Sommaire du manuel](/wiki/Handbook:HPPA "Handbook:HPPA")[Installation](/wiki/Handbook:HPPA/Full/Installation/fr "Handbook:HPPA/Full/Installation/fr")[À propos de l’installation](/wiki/Handbook:HPPA/Installation/About/fr "Handbook:HPPA/Installation/About/fr")[Choix du support](/wiki/Handbook:HPPA/Installation/Media/fr "Handbook:HPPA/Installation/Media/fr")[Configurer le réseau](/wiki/Handbook:HPPA/Installation/Networking/fr "Handbook:HPPA/Installation/Networking/fr")[Préparer les disques](/wiki/Handbook:HPPA/Installation/Disks/fr "Handbook:HPPA/Installation/Disks/fr")[Installer l’archive _stage3_](/wiki/Handbook:HPPA/Installation/Stage/fr "Handbook:HPPA/Installation/Stage/fr")[Installer le système de base](/wiki/Handbook:HPPA/Installation/Base/fr "Handbook:HPPA/Installation/Base/fr")[Configurer le noyau](/wiki/Handbook:HPPA/Installation/Kernel/fr "Handbook:HPPA/Installation/Kernel/fr")[Configurer le système](/wiki/Handbook:HPPA/Installation/System/fr "Handbook:HPPA/Installation/System/fr")[Installer les outils](/wiki/Handbook:HPPA/Installation/Tools/fr "Handbook:HPPA/Installation/Tools/fr")Configurer le système d’amorçage[Finaliser](/wiki/Handbook:HPPA/Installation/Finalizing/fr "Handbook:HPPA/Installation/Finalizing/fr")[Utiliser Gentoo](/wiki/Handbook:HPPA/Full/Working/fr "Handbook:HPPA/Full/Working/fr")[Introduction à Portage](/wiki/Handbook:HPPA/Working/Portage/fr "Handbook:HPPA/Working/Portage/fr")[Les options de la variable USE](/wiki/Handbook:HPPA/Working/USE/fr "Handbook:HPPA/Working/USE/fr")[Les fonctionnalités de Portage](/wiki/Handbook:HPPA/Working/Features/fr "Handbook:HPPA/Working/Features/fr")[Scripts d’initialisation systèmes](/wiki/Handbook:HPPA/Working/Initscripts/fr "Handbook:HPPA/Working/Initscripts/fr")[Variables d’environnement](/wiki/Handbook:HPPA/Working/EnvVar/fr "Handbook:HPPA/Working/EnvVar/fr")[Utiliser Portage](/wiki/Handbook:HPPA/Full/Portage/fr "Handbook:HPPA/Full/Portage/fr")[Fichiers et répertoires](/wiki/Handbook:HPPA/Portage/Files/fr "Handbook:HPPA/Portage/Files/fr")[Les variables](/wiki/Handbook:HPPA/Portage/Variables/fr "Handbook:HPPA/Portage/Variables/fr")[Mélanger plusieurs branches logicielles](/wiki/Handbook:HPPA/Portage/Branches/fr "Handbook:HPPA/Portage/Branches/fr")[Outils supplémentaires](/wiki/Handbook:HPPA/Portage/Tools/fr "Handbook:HPPA/Portage/Tools/fr")[Dépôt personnalisé](/wiki/Handbook:HPPA/Portage/CustomTree/fr "Handbook:HPPA/Portage/CustomTree/fr")[Fonctionnalités avancées](/wiki/Handbook:HPPA/Portage/Advanced/fr "Handbook:HPPA/Portage/Advanced/fr")[Configuration du réseau avec OpenRC](/wiki/Handbook:HPPA/Full/Networking/fr "Handbook:HPPA/Full/Networking/fr")[Bien démarrer](/wiki/Handbook:HPPA/Networking/Introduction/fr "Handbook:HPPA/Networking/Introduction/fr")[Configuration avancée](/wiki/Handbook:HPPA/Networking/Advanced/fr "Handbook:HPPA/Networking/Advanced/fr")[Les modules réseau](/wiki/Handbook:HPPA/Networking/Modular/fr "Handbook:HPPA/Networking/Modular/fr")[Sans fil](/wiki/Handbook:HPPA/Networking/Wireless/fr "Handbook:HPPA/Networking/Wireless/fr")[Ajouter des fonctionnalités](/wiki/Handbook:HPPA/Networking/Extending/fr "Handbook:HPPA/Networking/Extending/fr")[Gestion dynamique](/wiki/Handbook:HPPA/Networking/Dynamic/fr "Handbook:HPPA/Networking/Dynamic/fr")

[Handbook:HPPA/Blocks/Bootloader/fr](/index.php?title=Handbook:HPPA/Blocks/Bootloader/fr&action=edit&redlink=1 "Handbook:HPPA/Blocks/Bootloader/fr (page does not exist)")

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

Une fois redémarré dans le nouvel environnement Gentoo, il est avisé de finir avec [Finalisation de l’installation de Gentoo](/wiki/Handbook:HPPA/Installation/Finalizing/fr "Handbook:HPPA/Installation/Finalizing/fr").

[← Installing tools](/wiki/Handbook:HPPA/Installation/Tools/fr "Previous part") [Accueil](/wiki/Handbook:HPPA/fr "Handbook:HPPA/fr") [Finalisation →](/wiki/Handbook:HPPA/Installation/Finalizing/fr "Next part")