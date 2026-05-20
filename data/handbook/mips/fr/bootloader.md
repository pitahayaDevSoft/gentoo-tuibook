# Bootloader

Other languages:

- [Deutsch](/wiki/Handbook:MIPS/Installation/Bootloader/de "Handbuch:MIPS/Installation/Bootloader (100% translated)")
- [English](/wiki/Handbook:MIPS/Installation/Bootloader "Handbook:MIPS/Installation/Bootloader (100% translated)")
- [español](/wiki/Handbook:MIPS/Installation/Bootloader/es "Manual de Gentoo: MIPS/Instalación/Arranque (100% translated)")
- français
- [italiano](/wiki/Handbook:MIPS/Installation/Bootloader/it "Handbook:MIPS/Installation/Bootloader/it (100% translated)")
- [magyar](/wiki/Handbook:MIPS/Installation/Bootloader/hu "Handbook:MIPS/Installation/Bootloader/hu (100% translated)")
- [polski](/wiki/Handbook:MIPS/Installation/Bootloader/pl "Handbook:MIPS/Installation/Bootloader (100% translated)")
- [português do Brasil](/wiki/Handbook:MIPS/Installation/Bootloader/pt-br "Manual:MIPS/Instalação/Gerenciador de Boot (100% translated)")
- [русский](/wiki/Handbook:MIPS/Installation/Bootloader/ru "Handbook:MIPS/Installation/Bootloader (100% translated)")
- [தமிழ்](/wiki/Handbook:MIPS/Installation/Bootloader/ta "கையேடு:MIPS/நிறுவல்/துவக்கஏற்றி (100% translated)")
- [中文（中国大陆）‎](/wiki/Handbook:MIPS/Installation/Bootloader/zh-cn "手册：MIPS/安装/配置系统引导程序Bootloader (100% translated)")
- [日本語](/wiki/Handbook:MIPS/Installation/Bootloader/ja "ハンドブック:MIPS/インストール/ブートローダー (100% translated)")
- [한국어](/wiki/Handbook:MIPS/Installation/Bootloader/ko "Handbook:MIPS/Installation/Bootloader/ko (100% translated)")

[Sommaire du manuel](/wiki/Handbook:MIPS "Handbook:MIPS")[Installation](/wiki/Handbook:MIPS/Full/Installation/fr "Handbook:MIPS/Full/Installation/fr")[À propos de l’installation](/wiki/Handbook:MIPS/Installation/About/fr "Handbook:MIPS/Installation/About/fr")[Choix du support](/wiki/Handbook:MIPS/Installation/Media/fr "Handbook:MIPS/Installation/Media/fr")[Configurer le réseau](/wiki/Handbook:MIPS/Installation/Networking/fr "Handbook:MIPS/Installation/Networking/fr")[Préparer les disques](/wiki/Handbook:MIPS/Installation/Disks/fr "Handbook:MIPS/Installation/Disks/fr")[Installer l’archive _stage3_](/wiki/Handbook:MIPS/Installation/Stage/fr "Handbook:MIPS/Installation/Stage/fr")[Installer le système de base](/wiki/Handbook:MIPS/Installation/Base/fr "Handbook:MIPS/Installation/Base/fr")[Configurer le noyau](/wiki/Handbook:MIPS/Installation/Kernel/fr "Handbook:MIPS/Installation/Kernel/fr")[Configurer le système](/wiki/Handbook:MIPS/Installation/System/fr "Handbook:MIPS/Installation/System/fr")[Installer les outils](/wiki/Handbook:MIPS/Installation/Tools/fr "Handbook:MIPS/Installation/Tools/fr")Configurer le système d’amorçage[Finaliser](/wiki/Handbook:MIPS/Installation/Finalizing/fr "Handbook:MIPS/Installation/Finalizing/fr")[Utiliser Gentoo](/wiki/Handbook:MIPS/Full/Working/fr "Handbook:MIPS/Full/Working/fr")[Introduction à Portage](/wiki/Handbook:MIPS/Working/Portage/fr "Handbook:MIPS/Working/Portage/fr")[Les options de la variable USE](/wiki/Handbook:MIPS/Working/USE/fr "Handbook:MIPS/Working/USE/fr")[Les fonctionnalités de Portage](/wiki/Handbook:MIPS/Working/Features/fr "Handbook:MIPS/Working/Features/fr")[Scripts d’initialisation systèmes](/wiki/Handbook:MIPS/Working/Initscripts/fr "Handbook:MIPS/Working/Initscripts/fr")[Variables d’environnement](/wiki/Handbook:MIPS/Working/EnvVar/fr "Handbook:MIPS/Working/EnvVar/fr")[Utiliser Portage](/wiki/Handbook:MIPS/Full/Portage/fr "Handbook:MIPS/Full/Portage/fr")[Fichiers et répertoires](/wiki/Handbook:MIPS/Portage/Files/fr "Handbook:MIPS/Portage/Files/fr")[Les variables](/wiki/Handbook:MIPS/Portage/Variables/fr "Handbook:MIPS/Portage/Variables/fr")[Mélanger plusieurs branches logicielles](/wiki/Handbook:MIPS/Portage/Branches/fr "Handbook:MIPS/Portage/Branches/fr")[Outils supplémentaires](/wiki/Handbook:MIPS/Portage/Tools/fr "Handbook:MIPS/Portage/Tools/fr")[Dépôt personnalisé](/wiki/Handbook:MIPS/Portage/CustomTree/fr "Handbook:MIPS/Portage/CustomTree/fr")[Fonctionnalités avancées](/wiki/Handbook:MIPS/Portage/Advanced/fr "Handbook:MIPS/Portage/Advanced/fr")[Configuration du réseau avec OpenRC](/wiki/Handbook:MIPS/Full/Networking/fr "Handbook:MIPS/Full/Networking/fr")[Bien démarrer](/wiki/Handbook:MIPS/Networking/Introduction/fr "Handbook:MIPS/Networking/Introduction/fr")[Configuration avancée](/wiki/Handbook:MIPS/Networking/Advanced/fr "Handbook:MIPS/Networking/Advanced/fr")[Les modules réseau](/wiki/Handbook:MIPS/Networking/Modular/fr "Handbook:MIPS/Networking/Modular/fr")[Sans fil](/wiki/Handbook:MIPS/Networking/Wireless/fr "Handbook:MIPS/Networking/Wireless/fr")[Ajouter des fonctionnalités](/wiki/Handbook:MIPS/Networking/Extending/fr "Handbook:MIPS/Networking/Extending/fr")[Gestion dynamique](/wiki/Handbook:MIPS/Networking/Dynamic/fr "Handbook:MIPS/Networking/Dynamic/fr")

[Handbook:MIPS/Blocks/Bootloader/fr](/index.php?title=Handbook:MIPS/Blocks/Bootloader/fr&action=edit&redlink=1 "Handbook:MIPS/Blocks/Bootloader/fr (page does not exist)")

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

Une fois redémarré dans le nouvel environnement Gentoo, il est avisé de finir avec [Finalisation de l’installation de Gentoo](/wiki/Handbook:MIPS/Installation/Finalizing/fr "Handbook:MIPS/Installation/Finalizing/fr").

[← Installing tools](/wiki/Handbook:MIPS/Installation/Tools/fr "Previous part") [Accueil](/wiki/Handbook:MIPS/fr "Handbook:MIPS/fr") [Finalisation →](/wiki/Handbook:MIPS/Installation/Finalizing/fr "Next part")