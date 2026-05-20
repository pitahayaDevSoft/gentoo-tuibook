# Bootloader

Other languages:

- [Deutsch](/wiki/Handbook:PPC/Installation/Bootloader/de "Handbuch:PPC/Installation/Bootloader (100% translated)")
- [English](/wiki/Handbook:PPC/Installation/Bootloader "Handbook:PPC/Installation/Bootloader (100% translated)")
- [español](/wiki/Handbook:PPC/Installation/Bootloader/es "Manual de Gentoo: PPC/Instalación/Arranque (100% translated)")
- français
- [italiano](/wiki/Handbook:PPC/Installation/Bootloader/it "Handbook:PPC/Installation/Bootloader/it (100% translated)")
- [magyar](/wiki/Handbook:PPC/Installation/Bootloader/hu "Handbook:PPC/Installation/Bootloader/hu (100% translated)")
- [polski](/wiki/Handbook:PPC/Installation/Bootloader/pl "Handbook:PPC/Installation/Bootloader (100% translated)")
- [português do Brasil](/wiki/Handbook:PPC/Installation/Bootloader/pt-br "Handbook:PPC/Installation/Bootloader/pt-br (50% translated)")
- [русский](/wiki/Handbook:PPC/Installation/Bootloader/ru "Handbook:PPC/Installation/Bootloader (100% translated)")
- [தமிழ்](/wiki/Handbook:PPC/Installation/Bootloader/ta "கையேடு:PPC/நிறுவல்/துவக்கஏற்றி (100% translated)")
- [中文（中国大陆）‎](/wiki/Handbook:PPC/Installation/Bootloader/zh-cn "手册：PPC/安装/配置系统引导程序Bootloader (100% translated)")
- [日本語](/wiki/Handbook:PPC/Installation/Bootloader/ja "ハンドブック:PPC/インストール/ブートローダー (100% translated)")
- [한국어](/wiki/Handbook:PPC/Installation/Bootloader/ko "Handbook:PPC/Installation/Bootloader/ko (100% translated)")

[Sommaire du manuel](/wiki/Handbook:PPC "Handbook:PPC")[Installation](/wiki/Handbook:PPC/Full/Installation/fr "Handbook:PPC/Full/Installation/fr")[À propos de l’installation](/wiki/Handbook:PPC/Installation/About/fr "Handbook:PPC/Installation/About/fr")[Choix du support](/wiki/Handbook:PPC/Installation/Media/fr "Handbook:PPC/Installation/Media/fr")[Configurer le réseau](/wiki/Handbook:PPC/Installation/Networking/fr "Handbook:PPC/Installation/Networking/fr")[Préparer les disques](/wiki/Handbook:PPC/Installation/Disks/fr "Handbook:PPC/Installation/Disks/fr")[Installer l’archive _stage3_](/wiki/Handbook:PPC/Installation/Stage/fr "Handbook:PPC/Installation/Stage/fr")[Installer le système de base](/wiki/Handbook:PPC/Installation/Base/fr "Handbook:PPC/Installation/Base/fr")[Configurer le noyau](/wiki/Handbook:PPC/Installation/Kernel/fr "Handbook:PPC/Installation/Kernel/fr")[Configurer le système](/wiki/Handbook:PPC/Installation/System/fr "Handbook:PPC/Installation/System/fr")[Installer les outils](/wiki/Handbook:PPC/Installation/Tools/fr "Handbook:PPC/Installation/Tools/fr")Configurer le système d’amorçage[Finaliser](/wiki/Handbook:PPC/Installation/Finalizing/fr "Handbook:PPC/Installation/Finalizing/fr")[Utiliser Gentoo](/wiki/Handbook:PPC/Full/Working/fr "Handbook:PPC/Full/Working/fr")[Introduction à Portage](/wiki/Handbook:PPC/Working/Portage/fr "Handbook:PPC/Working/Portage/fr")[Les options de la variable USE](/wiki/Handbook:PPC/Working/USE/fr "Handbook:PPC/Working/USE/fr")[Les fonctionnalités de Portage](/wiki/Handbook:PPC/Working/Features/fr "Handbook:PPC/Working/Features/fr")[Scripts d’initialisation systèmes](/wiki/Handbook:PPC/Working/Initscripts/fr "Handbook:PPC/Working/Initscripts/fr")[Variables d’environnement](/wiki/Handbook:PPC/Working/EnvVar/fr "Handbook:PPC/Working/EnvVar/fr")[Utiliser Portage](/wiki/Handbook:PPC/Full/Portage/fr "Handbook:PPC/Full/Portage/fr")[Fichiers et répertoires](/wiki/Handbook:PPC/Portage/Files/fr "Handbook:PPC/Portage/Files/fr")[Les variables](/wiki/Handbook:PPC/Portage/Variables/fr "Handbook:PPC/Portage/Variables/fr")[Mélanger plusieurs branches logicielles](/wiki/Handbook:PPC/Portage/Branches/fr "Handbook:PPC/Portage/Branches/fr")[Outils supplémentaires](/wiki/Handbook:PPC/Portage/Tools/fr "Handbook:PPC/Portage/Tools/fr")[Dépôt personnalisé](/wiki/Handbook:PPC/Portage/CustomTree/fr "Handbook:PPC/Portage/CustomTree/fr")[Fonctionnalités avancées](/wiki/Handbook:PPC/Portage/Advanced/fr "Handbook:PPC/Portage/Advanced/fr")[Configuration du réseau avec OpenRC](/wiki/Handbook:PPC/Full/Networking/fr "Handbook:PPC/Full/Networking/fr")[Bien démarrer](/wiki/Handbook:PPC/Networking/Introduction/fr "Handbook:PPC/Networking/Introduction/fr")[Configuration avancée](/wiki/Handbook:PPC/Networking/Advanced/fr "Handbook:PPC/Networking/Advanced/fr")[Les modules réseau](/wiki/Handbook:PPC/Networking/Modular/fr "Handbook:PPC/Networking/Modular/fr")[Sans fil](/wiki/Handbook:PPC/Networking/Wireless/fr "Handbook:PPC/Networking/Wireless/fr")[Ajouter des fonctionnalités](/wiki/Handbook:PPC/Networking/Extending/fr "Handbook:PPC/Networking/Extending/fr")[Gestion dynamique](/wiki/Handbook:PPC/Networking/Dynamic/fr "Handbook:PPC/Networking/Dynamic/fr")

[Handbook:PPC/Blocks/Bootloader/fr](/index.php?title=Handbook:PPC/Blocks/Bootloader/fr&action=edit&redlink=1 "Handbook:PPC/Blocks/Bootloader/fr (page does not exist)")

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

Une fois redémarré dans le nouvel environnement Gentoo, il est avisé de finir avec [Finalisation de l’installation de Gentoo](/wiki/Handbook:PPC/Installation/Finalizing/fr "Handbook:PPC/Installation/Finalizing/fr").

[← Installing tools](/wiki/Handbook:PPC/Installation/Tools/fr "Previous part") [Accueil](/wiki/Handbook:PPC/fr "Handbook:PPC/fr") [Finalisation →](/wiki/Handbook:PPC/Installation/Finalizing/fr "Next part")