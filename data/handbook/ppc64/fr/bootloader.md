# Bootloader

Other languages:

- [Deutsch](/wiki/Handbook:PPC64/Installation/Bootloader/de "Handbuch:PPC64/Installation/Bootloader (100% translated)")
- [English](/wiki/Handbook:PPC64/Installation/Bootloader "Handbook:PPC64/Installation/Bootloader (100% translated)")
- [español](/wiki/Handbook:PPC64/Installation/Bootloader/es "Manual de Gentoo: PPC64/Instalación/Arranque (100% translated)")
- français
- [italiano](/wiki/Handbook:PPC64/Installation/Bootloader/it "Handbook:PPC64/Installation/Bootloader/it (100% translated)")
- [magyar](/wiki/Handbook:PPC64/Installation/Bootloader/hu "Handbook:PPC64/Installation/Bootloader/hu (100% translated)")
- [polski](/wiki/Handbook:PPC64/Installation/Bootloader/pl "Handbook:PPC64/Installation/Bootloader (100% translated)")
- [português do Brasil](/wiki/Handbook:PPC64/Installation/Bootloader/pt-br "Handbook:PPC64/Installation/Bootloader/pt-br (50% translated)")
- [русский](/wiki/Handbook:PPC64/Installation/Bootloader/ru "Handbook:PPC64/Installation/Bootloader (100% translated)")
- [தமிழ்](/wiki/Handbook:PPC64/Installation/Bootloader/ta "கையேடு:PPC64/நிறுவல்/துவக்கஏற்றி (100% translated)")
- [中文（中国大陆）‎](/wiki/Handbook:PPC64/Installation/Bootloader/zh-cn "手册：PPC64/安装/配置系统引导程序Bootloader (100% translated)")
- [日本語](/wiki/Handbook:PPC64/Installation/Bootloader/ja "ハンドブック:PPC64/インストール/ブートローダー (100% translated)")
- [한국어](/wiki/Handbook:PPC64/Installation/Bootloader/ko "Handbook:PPC64/Installation/Bootloader/ko (100% translated)")

[Sommaire du manuel](/wiki/Handbook:PPC64 "Handbook:PPC64")[Installation](/wiki/Handbook:PPC64/Full/Installation/fr "Handbook:PPC64/Full/Installation/fr")[À propos de l’installation](/wiki/Handbook:PPC64/Installation/About/fr "Handbook:PPC64/Installation/About/fr")[Choix du support](/wiki/Handbook:PPC64/Installation/Media/fr "Handbook:PPC64/Installation/Media/fr")[Configurer le réseau](/wiki/Handbook:PPC64/Installation/Networking/fr "Handbook:PPC64/Installation/Networking/fr")[Préparer les disques](/wiki/Handbook:PPC64/Installation/Disks/fr "Handbook:PPC64/Installation/Disks/fr")[Installer l’archive _stage3_](/wiki/Handbook:PPC64/Installation/Stage/fr "Handbook:PPC64/Installation/Stage/fr")[Installer le système de base](/wiki/Handbook:PPC64/Installation/Base/fr "Handbook:PPC64/Installation/Base/fr")[Configurer le noyau](/wiki/Handbook:PPC64/Installation/Kernel/fr "Handbook:PPC64/Installation/Kernel/fr")[Configurer le système](/wiki/Handbook:PPC64/Installation/System/fr "Handbook:PPC64/Installation/System/fr")[Installer les outils](/wiki/Handbook:PPC64/Installation/Tools/fr "Handbook:PPC64/Installation/Tools/fr")Configurer le système d’amorçage[Finaliser](/wiki/Handbook:PPC64/Installation/Finalizing/fr "Handbook:PPC64/Installation/Finalizing/fr")[Utiliser Gentoo](/wiki/Handbook:PPC64/Full/Working/fr "Handbook:PPC64/Full/Working/fr")[Introduction à Portage](/wiki/Handbook:PPC64/Working/Portage/fr "Handbook:PPC64/Working/Portage/fr")[Les options de la variable USE](/wiki/Handbook:PPC64/Working/USE/fr "Handbook:PPC64/Working/USE/fr")[Les fonctionnalités de Portage](/wiki/Handbook:PPC64/Working/Features/fr "Handbook:PPC64/Working/Features/fr")[Scripts d’initialisation systèmes](/wiki/Handbook:PPC64/Working/Initscripts/fr "Handbook:PPC64/Working/Initscripts/fr")[Variables d’environnement](/wiki/Handbook:PPC64/Working/EnvVar/fr "Handbook:PPC64/Working/EnvVar/fr")[Utiliser Portage](/wiki/Handbook:PPC64/Full/Portage/fr "Handbook:PPC64/Full/Portage/fr")[Fichiers et répertoires](/wiki/Handbook:PPC64/Portage/Files/fr "Handbook:PPC64/Portage/Files/fr")[Les variables](/wiki/Handbook:PPC64/Portage/Variables/fr "Handbook:PPC64/Portage/Variables/fr")[Mélanger plusieurs branches logicielles](/wiki/Handbook:PPC64/Portage/Branches/fr "Handbook:PPC64/Portage/Branches/fr")[Outils supplémentaires](/wiki/Handbook:PPC64/Portage/Tools/fr "Handbook:PPC64/Portage/Tools/fr")[Dépôt personnalisé](/wiki/Handbook:PPC64/Portage/CustomTree/fr "Handbook:PPC64/Portage/CustomTree/fr")[Fonctionnalités avancées](/wiki/Handbook:PPC64/Portage/Advanced/fr "Handbook:PPC64/Portage/Advanced/fr")[Configuration du réseau avec OpenRC](/wiki/Handbook:PPC64/Full/Networking/fr "Handbook:PPC64/Full/Networking/fr")[Bien démarrer](/wiki/Handbook:PPC64/Networking/Introduction/fr "Handbook:PPC64/Networking/Introduction/fr")[Configuration avancée](/wiki/Handbook:PPC64/Networking/Advanced/fr "Handbook:PPC64/Networking/Advanced/fr")[Les modules réseau](/wiki/Handbook:PPC64/Networking/Modular/fr "Handbook:PPC64/Networking/Modular/fr")[Sans fil](/wiki/Handbook:PPC64/Networking/Wireless/fr "Handbook:PPC64/Networking/Wireless/fr")[Ajouter des fonctionnalités](/wiki/Handbook:PPC64/Networking/Extending/fr "Handbook:PPC64/Networking/Extending/fr")[Gestion dynamique](/wiki/Handbook:PPC64/Networking/Dynamic/fr "Handbook:PPC64/Networking/Dynamic/fr")

[Handbook:PPC64/Blocks/Bootloader/fr](/index.php?title=Handbook:PPC64/Blocks/Bootloader/fr&action=edit&redlink=1 "Handbook:PPC64/Blocks/Bootloader/fr (page does not exist)")

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

Une fois redémarré dans le nouvel environnement Gentoo, il est avisé de finir avec [Finalisation de l’installation de Gentoo](/wiki/Handbook:PPC64/Installation/Finalizing/fr "Handbook:PPC64/Installation/Finalizing/fr").

[← Installing tools](/wiki/Handbook:PPC64/Installation/Tools/fr "Previous part") [Accueil](/wiki/Handbook:PPC64/fr "Handbook:PPC64/fr") [Finalisation →](/wiki/Handbook:PPC64/Installation/Finalizing/fr "Next part")