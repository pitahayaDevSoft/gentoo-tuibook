# Disks

Other languages:

- [Deutsch](/wiki/Handbook:PPC64/Installation/Disks/de "Handbuch:PPC64/Installation/Festplatten (100% translated)")
- [English](/wiki/Handbook:PPC64/Installation/Disks "Handbook:PPC64/Installation/Disks (100% translated)")
- [español](/wiki/Handbook:PPC64/Installation/Disks/es "Manual de Gentoo: PPC64/Instalación/Discos (100% translated)")
- français
- [italiano](/wiki/Handbook:PPC64/Installation/Disks/it "Handbook:PPC64/Installation/Disks/it (100% translated)")
- [magyar](/wiki/Handbook:PPC64/Installation/Disks/hu "Handbook:PPC64/Installation/Disks/hu (100% translated)")
- [polski](/wiki/Handbook:PPC64/Installation/Disks/pl "Handbook:PPC64/Installation/Disks (100% translated)")
- [português do Brasil](/wiki/Handbook:PPC64/Installation/Disks/pt-br "Handbook:PPC64/Installation/Disks/pt-br (50% translated)")
- [русский](/wiki/Handbook:PPC64/Installation/Disks/ru "Handbook:PPC64/Installation/Disks (100% translated)")
- [தமிழ்](/wiki/Handbook:PPC64/Installation/Disks/ta "கையேடு:PPC64/நிறுவல்/வட்டுக்கள் (100% translated)")
- [中文（中国大陆）‎](/wiki/Handbook:PPC64/Installation/Disks/zh-cn "手册：PPC64/安装/准备磁盘 (100% translated)")
- [日本語](/wiki/Handbook:PPC64/Installation/Disks/ja "ハンドブック:PPC64/インストール/ディスク (100% translated)")
- [한국어](/wiki/Handbook:PPC64/Installation/Disks/ko "Handbook:PPC64/Installation/Disks/ko (100% translated)")

[Sommaire du manuel](/wiki/Handbook:PPC64 "Handbook:PPC64")[Installation](/wiki/Handbook:PPC64/Full/Installation/fr "Handbook:PPC64/Full/Installation/fr")[À propos de l’installation](/wiki/Handbook:PPC64/Installation/About/fr "Handbook:PPC64/Installation/About/fr")[Choix du support](/wiki/Handbook:PPC64/Installation/Media/fr "Handbook:PPC64/Installation/Media/fr")[Configurer le réseau](/wiki/Handbook:PPC64/Installation/Networking/fr "Handbook:PPC64/Installation/Networking/fr")Préparer les disques[Installer l’archive _stage3_](/wiki/Handbook:PPC64/Installation/Stage/fr "Handbook:PPC64/Installation/Stage/fr")[Installer le système de base](/wiki/Handbook:PPC64/Installation/Base/fr "Handbook:PPC64/Installation/Base/fr")[Configurer le noyau](/wiki/Handbook:PPC64/Installation/Kernel/fr "Handbook:PPC64/Installation/Kernel/fr")[Configurer le système](/wiki/Handbook:PPC64/Installation/System/fr "Handbook:PPC64/Installation/System/fr")[Installer les outils](/wiki/Handbook:PPC64/Installation/Tools/fr "Handbook:PPC64/Installation/Tools/fr")[Configurer le système d’amorçage](/wiki/Handbook:PPC64/Installation/Bootloader/fr "Handbook:PPC64/Installation/Bootloader/fr")[Finaliser](/wiki/Handbook:PPC64/Installation/Finalizing/fr "Handbook:PPC64/Installation/Finalizing/fr")[Utiliser Gentoo](/wiki/Handbook:PPC64/Full/Working/fr "Handbook:PPC64/Full/Working/fr")[Introduction à Portage](/wiki/Handbook:PPC64/Working/Portage/fr "Handbook:PPC64/Working/Portage/fr")[Les options de la variable USE](/wiki/Handbook:PPC64/Working/USE/fr "Handbook:PPC64/Working/USE/fr")[Les fonctionnalités de Portage](/wiki/Handbook:PPC64/Working/Features/fr "Handbook:PPC64/Working/Features/fr")[Scripts d’initialisation systèmes](/wiki/Handbook:PPC64/Working/Initscripts/fr "Handbook:PPC64/Working/Initscripts/fr")[Variables d’environnement](/wiki/Handbook:PPC64/Working/EnvVar/fr "Handbook:PPC64/Working/EnvVar/fr")[Utiliser Portage](/wiki/Handbook:PPC64/Full/Portage/fr "Handbook:PPC64/Full/Portage/fr")[Fichiers et répertoires](/wiki/Handbook:PPC64/Portage/Files/fr "Handbook:PPC64/Portage/Files/fr")[Les variables](/wiki/Handbook:PPC64/Portage/Variables/fr "Handbook:PPC64/Portage/Variables/fr")[Mélanger plusieurs branches logicielles](/wiki/Handbook:PPC64/Portage/Branches/fr "Handbook:PPC64/Portage/Branches/fr")[Outils supplémentaires](/wiki/Handbook:PPC64/Portage/Tools/fr "Handbook:PPC64/Portage/Tools/fr")[Dépôt personnalisé](/wiki/Handbook:PPC64/Portage/CustomTree/fr "Handbook:PPC64/Portage/CustomTree/fr")[Fonctionnalités avancées](/wiki/Handbook:PPC64/Portage/Advanced/fr "Handbook:PPC64/Portage/Advanced/fr")[Configuration du réseau avec OpenRC](/wiki/Handbook:PPC64/Full/Networking/fr "Handbook:PPC64/Full/Networking/fr")[Bien démarrer](/wiki/Handbook:PPC64/Networking/Introduction/fr "Handbook:PPC64/Networking/Introduction/fr")[Configuration avancée](/wiki/Handbook:PPC64/Networking/Advanced/fr "Handbook:PPC64/Networking/Advanced/fr")[Les modules réseau](/wiki/Handbook:PPC64/Networking/Modular/fr "Handbook:PPC64/Networking/Modular/fr")[Sans fil](/wiki/Handbook:PPC64/Networking/Wireless/fr "Handbook:PPC64/Networking/Wireless/fr")[Ajouter des fonctionnalités](/wiki/Handbook:PPC64/Networking/Extending/fr "Handbook:PPC64/Networking/Extending/fr")[Gestion dynamique](/wiki/Handbook:PPC64/Networking/Dynamic/fr "Handbook:PPC64/Networking/Dynamic/fr")

## Contents

- [1Introduction aux périphériques de type bloc](#Introduction_aux_p.C3.A9riph.C3.A9riques_de_type_bloc)
  - [1.1Les périphériques de type bloc](#Les_p.C3.A9riph.C3.A9riques_de_type_bloc)
- [2Créer des systèmes de fichiers](#Cr.C3.A9er_des_syst.C3.A8mes_de_fichiers)
- [3Introduction](#Introduction)
  - [3.1Les systèmes de fichiers](#Les_syst.C3.A8mes_de_fichiers)
  - [3.2Mettre en œuvre un système de fichiers sur une partition](#Mettre_en_.C5.93uvre_un_syst.C3.A8me_de_fichiers_sur_une_partition)
    - [3.2.1Système de fichiers pour un BIOS déprécié](#Syst.C3.A8me_de_fichiers_pour_un_BIOS_d.C3.A9pr.C3.A9ci.C3.A9)
    - [3.2.2Petite partition ext4](#Petite_partition_ext4)
  - [3.3Activer la partition d’échange](#Activer_la_partition_d.E2.80.99.C3.A9change)
- [4Monter la partition racine](#Monter_la_partition_racine)

## Introduction aux périphériques de type bloc

### Les périphériques de type bloc

Étudions en détail les aspects de Gentoo et de Linux en général qui concernent les disques, incluant les périphériques de type bloc, les partitions, et les systèmes de fichiers de Linux. Une fois que les tenants et les aboutissants des disques seront compris, il sera possible d’établir les partitions et les systèmes de fichiers pour l’installation.

Pour commencer, intéressons-nous aux périphériques de type bloc. Les disques SCSI et Serial ATA sont tous les deux étiquetés avec des noms de périphérique tels que : /dev/sda, /dev/sdb, /dev/sdc, etc. Sur des machines plus modernes, les disques SSD NVMe basés sur PCI Express ont des noms de périphérique tels que : /dev/nvme0n1, /dev/nvme0n2, etc.

Le tableau suivant aidera les lecteurs à déterminer où trouver certaines catégories de périphériques de type bloc sur le système :

Type de périphériqueChemin par défautCommentaires
SATA, SAS, SCSI ou USB/dev/sdaSe trouve dans le matériel à partir de 2007 jusqu’à présent, ce sont les plus courants. Ils se connectent via un bus [_SATA_](https://en.wikipedia.org/wiki/Serial_ATA "wikipedia:Serial ATA"), [_SCSI_](https://en.wikipedia.org/wiki/SCSI "wikipedia:SCSI") ou [_USB_](https://en.wikipedia.org/wiki/USB "wikipedia:USB") comme un bloc de stockage. Par exemple, la 1re partition d’un périphérique SATA sera /dev/sda1.
NVM Express (NVMe)/dev/nvme0n1La dernière technologie pour les disques _SSD_. Les disques [NVMe](https://en.wikipedia.org/wiki/NVM_Express "wikipedia:NVM Express") sont connectés directement au bus PCI Express et ont la plus rapide vitesse de transfert sur le marché. Le matériel est sorti à partir de 2014, le support est maintenant fréquent. La 1re partition d’un périphérique NVMe sera /dev/nvme0n1p1.
MMC, eMMC et SD/dev/mmcblk0Les cartes [MMC](https://en.wikipedia.org/wiki/MultiMediaCard#eMMC "wikipedia:MultiMediaCard"), SD et les autres types de mémoire peuvent être utiles pour le stockage. Ceci étant dit, peu de systèmes permettent une amorce à partir de ce type de périphérique. Il est suggéré de ne **pas** utiliser cela pour une installation Linux ; mais plutôt de les cantonner au transfert de fichiers, leur but initial. Alternativement, ce type de stockage peut être pratique pour des sauvegardes à court terme.

Les périphériques de type bloc ci-dessus représentent une interface abstraite pour le disque. Les programmes utilisateurs peuvent utiliser ces périphériques de type bloc pour interagir avec le disque sans se soucier de savoir s’il est SATA, SCSI ou quelque chose d’autre. Le programme peut simplement adresser le stockage sur le disque comme un groupe de blocs contigus de 4096 octets (4 Kio), accessibles aléatoirement.

[Handbook:PPC64/Blocks/Disks/fr](/index.php?title=Handbook:PPC64/Blocks/Disks/fr&action=edit&redlink=1 "Handbook:PPC64/Blocks/Disks/fr (page does not exist)")

## Créer des systèmes de fichiers

**Attention !**

Il est sage de vérifier les mises à jour du micrologiciel des disques SSD ou MVNe. Certains SSD Intel (en particulier 600p et 6000p) nécessaire une mise à jour pour [éviter une corruption de données](https://bugzilla.redhat.com/show_bug.cgi?id=1402533) sur des types d’usage avec XFS. Le problème est au niveau du micrologiciel et pas XFS. L’outil smartctl peut aider à identifier le modèle et la version du micrologiciel.

## Introduction

Maintenant que les partitions ont été créées, il est temps d’y placer un système de fichiers. Dans la section suivante les différents systèmes de fichiers que Linux prend en charge seront décris. Les lecteurs qui connaissent déjà quel système de fichiers utiliser peuvent continuer avec [Appliquer un système de fichiers à une partition](/wiki/Handbook:PPC64/Installation/Disks/fr#Appliquer_un_syst.C3.A8me_de_fichiers_.C3.A0_une_partition "Handbook:PPC64/Installation/Disks/fr"). Les autres devraient continuer à lire pour en apprendre plus sur les systèmes de fichiers disponibles.

### Les systèmes de fichiers

Linux supporte des douzaines de systèmes de fichiers, bien que la plupart ne devrait être utilisés que dans des cas spécifiques. Seuls certains seront dans l’architecture ppc64 – il est conseillé de se renseigner sur les systèmes de fichiers et leur prise en charge avant d’en choisir un plus expérimental pour les partitions importantes. **XFS est celui recommandé pour tous les usages et toutes les plates-formes.** Ci-dessous une liste non exhaustive :

[XFS](/wiki/XFS/fr "XFS/fr")Un système de fichiers avec journalisation des métadonnées, doté d’un ensemble de fonctionnalités robuste et optimisé pour la mise à l’échelle. Il a été continuellement mis à jour avec de nouvelles fonctionnalités. Le seul inconvénient est que les partitions XFS ne peuvent pas encore être réduites, bien que ce soit en cours de développement. XFS supporte _reflink_ (un lien vers un bloc de données ; ie. : si le fichier original est modifié, les blocs sont copiés pour que le _reflink_ continue de comporter les mêmes valeurs) et CoW ( _Copy on Write_ ; ie. : si plusieurs processus ouvrent le même fichier et qu’un le modifie, cela crée une copie transparent pour que le changement ne soit pas visible par les autres processus) qui sont particulièrement utilie dans un système Gentoo par rapport au nombre de compilation qu’un utilisateur fait. Nécessaire d’avoir une partition d’au moins 300 Mio.[ext4](/wiki/Ext4/fr "Ext4/fr")Ext4 est un système de fichiers fiable, tous usages, toutes plates-forme tout, bien qu’il manque des fonctionnalités modernes comme les _reflinks_.[VFAT](/wiki/VFAT "VFAT")Également connu sous le nom FAT32, ce format est supporté par Linux mais ne prend pas en charge les paramètres d’autorisation UNIX standards. Il est principalement utilisé pour l’interopérabilité avec d’autres systèmes d’exploitation (Microsoft Windows ou Apple macOS) mais est également une nécessité pour certains micrologiciels systèmes (comme UEFI). Les utilisateurs d’un système UEFI devront avoir une [Partition système EFI](/index.php?title=Partition_syst%C3%A8me_EFI&action=edit&redlink=1 "Partition système EFI (page does not exist)") formattée avec VFAT pour démarrer.[btrfs](/wiki/Btrfs/fr "Btrfs/fr")Un système de fichiers de nouvelle génération offrant de nombreuses fonctionnalités avancées telles que les sauvegardes instantanés, l’auto-guérison via des sommes de contrôle, la compression transparente, les sous-volumes et le RAID intégré. Les noyaux avant 5.4.x ne sont pas sûrs pour un usage de Btrfs en production car certains correctifs importants sont seulement présent dans les dernières versions. Les quotas par groupe et le RAID 5/6 ne sont pas sûrs d’usage quelque soit la version.[F2FS](/wiki/F2FS/fr "F2FS/fr")Le système de fichiers ami des SSD (disque flash) a été créé par Samsung pour l’utilisation avec la mémoire flash NAND. C’est un choix décent lors de l’installation de Gentoo sur des cartes microSD, des clés USB ou autres périphériques de stockage flash.[NTFS](/wiki/NTFS "NTFS")Ce système de fichiers « nouvelle technologie » est le système de fichiers phare de Microsoft Windows depuis NT 3.1. Comme VFAT ci-dessus, il ne stocke pas les paramètres d’autorisation UNIX ni les attributs étendus nécessaires au bon fonctionnement de BSD ou de Linux. Il ne peut donc pas être utilisé comme système de fichiers racine (aussi appelé _root_) dans la majorité des cas. Il devrait _seulement_ être utilisé pour l’interopérabilité avec les systèmes Microsoft Windows (noter l’emphase sur _seulement_).[ZFS](/wiki/ZFS "ZFS")**Important:** Les partitions ZFS peuvent seulement être crées depuis un LiveGUI. Pour plus d’informations, se référer à [ZFS/rootfs](/wiki/ZFS/rootfs "ZFS/rootfs")Système de fichiers nouvelle génération créé par Matthew Ahrens et Jeff Bonwick. Il a été conçu autour de quelques idées clés : l’administration du stockage doit être simple, la redondance être géré par le système de fichiers, le système ne soit pas être arrêté pour une réparation, des simulations des pires scénarios avant de livrer les versions et l’intégrité des données est primordial.[bcachefs](/wiki/Bcachefs "Bcachefs")**Important:** Bcachefs est encore noté expérimental dans le noyau ; aussi vérifiez que les données soient bien sauvegardées régulièrement dans un support sans Bcachefs.Bcachefs est un système de fichiers arbre B ( _B-tree_) basé sur Bcache. Il comporte des fonctionnalités comme CoW ( _Copy on Write_), la compression et le chiffrement. Bcachefs est comparable à Btrfs et ZFS. Une différence notable est le stockage multi-niveaux natif ; permettant d’utiliser un disque plus rapide (comme un disque SSD ou NVMe) pour faire office de cache pour un disque plus lent dans une grappe gérant de manière transparente l’activité sur les fichiers.

Des informations plus approfondies sur les systèmes de fichiers peuvent être trouvées dans les documentations maintenues par la [communauté](/wiki/Filesystem/fr "Filesystem/fr").

### Mettre en œuvre un système de fichiers sur une partition

**Remarque**

Assurez-vous d’installer les utilitaires correspondant au système de fichiers choisi avant de redémarrer. Il y aura un rappel à la fin de cette procédure d’installation.

Pour créer un système de fichiers sur une partition ou un volume, des outils sont disponibles pour chaque système de fichiers. Cliquer sur le nom du système de fichiers dans le tableau ci-dessous pour plus d’informations sur chaque système de fichiers :

Système de fichiers
Commande pour la création
Sur le CD minimal ?
Paquet
[XFS](/wiki/XFS/fr "XFS/fr")mkfs.xfs Oui
[sys-fs/xfsprogs](https://packages.gentoo.org/packages/sys-fs/xfsprogs)[ext4](/wiki/Ext4/fr "Ext4/fr")mkfs.ext4 Oui
[sys-fs/e2fsprogs](https://packages.gentoo.org/packages/sys-fs/e2fsprogs)[VFAT](/wiki/VFAT "VFAT") (FAT32, ...)
mkfs.vfat Oui
[sys-fs/dosfstools](https://packages.gentoo.org/packages/sys-fs/dosfstools)[btrfs](/wiki/Btrfs/fr "Btrfs/fr")mkfs.btrfs Oui
[sys-fs/btrfs-progs](https://packages.gentoo.org/packages/sys-fs/btrfs-progs)[F2FS](/wiki/F2FS/fr "F2FS/fr")mkfs.f2fs Oui
[sys-fs/f2fs-tools](https://packages.gentoo.org/packages/sys-fs/f2fs-tools)[NTFS](/wiki/NTFS "NTFS")mkfs.ntfs Oui
[sys-fs/ntfs3g](https://packages.gentoo.org/packages/sys-fs/ntfs3g)[ZFS](/wiki/ZFS "ZFS")zpool create ... Non
[sys-fs/zfs](https://packages.gentoo.org/packages/sys-fs/zfs)

\|}

**Important**

Le manuel recommande une nouvelle partition dans le cadre d’une installation, mais il est important de noter que n’importe quelle commande mkfs va effacer toute donnée contenue dans ladite partition. Si nécessaire, assurez-vous que chaque donnée soit sauvegardée avant cela.

Par exemple, pour avoir une partition racine (/dev/sda3) en xfs, les commandes suivantes doivent être utilisées :

`root #` `mkfs.xfs /dev/sda3`

#### Système de fichiers pour un BIOS déprécié

Les systèmes démarrant via un BIOS déprécié et avec un disque MBR/DOS peuvent utiliser n’importe quelle système de fichiers supportés par l’application d’amorçage ( _bootloader_, comme GRUB).

Par exemple, pour formater en XFS :

`root #` `mkfs.xfs /dev/sda1`

#### Petite partition ext4

Lors de l'utilisation d’u système de fichiers ext4 sur une petite partition (moins de 8 Gio), le système de fichiers devraient être créé avec les options appropriées pour réserver suffisamment de nœuds d’index ( _inodes_). Cela peut se faire avec l’option `-T small` :

`root #` `mkfs.ext4 -T small /dev/<device>`

Faire cela quadruple le nombre de nœuds d’index pour un système de fichiers étant donné que son paramètre _bytes-per-inode_ (octets par nœud d’index) passe de un tous les 16 Kio à un tous les 4 Kio.

### Activer la partition d’échange

mkswap est la commande à utiliser pour initialiser les partitions d’échange :

`root #` `mkswap /dev/sda2`

**Remarque**

Les installations débutées précédemment, mais non terminée peuvent reprendre ici dans le manuel. Utilisez ce lien comme lien permanent : [La reprise d’installation reprend ici](/wiki/Handbook:PPC64/Installation/Disks/fr#Resumed_installations_start_here "Handbook:PPC64/Installation/Disks/fr").

Pour activer la partition d’échange, utilisez la commande swapon :

`root #` `swapon /dev/sda2`

Cette étape d’« activation » est nécessaire car la partition d’échange est nouvellement créée dans un environnement démarré. Lors du redémarrage, si la partition est correctement déclarée dans fstab ou un autre mécanisme de montage, l’espace d’échange sera activé automatiquement.

## Monter la partition racine

Certains environnements peuvent ignorer le montage de la partition racine Gentoo (/mnt/gentoo) suggéré, ou monter des partitions additionnelles créées lors du partitionnement :

`root #` `mkdir --parents /mnt/gentoo`

Continuez à créer les points de montage additionnel pour chaque partition durant les étapes précédentes avec la commande mkdir.

Avec les points de montage créés, il est maintenant le moment de rendre les partitions accessibles via la commande mount.

Monter la partition racine :

`root #` `mount /dev/sda3 /mnt/gentoo`

Continuez à monter des partitions additionnelles avec la commande mount.

**Remarque**

Si /tmp/ doit se trouver sur une partition séparée, pensez à changer ses droits d'accès après le montage :

`root #` `chmod 1777 /mnt/gentoo/tmp`

Cela vaut également pour /var/tmp.

Plus loin dans les instructions, le système de fichiers _proc_ (une interface virtuelle avec le noyau) ainsi que d’autres pseudos systèmes de fichiers du noyau seront montés. Mais d’abord, nous devons [copier les fichiers d’installation de Gentoo](/wiki/Handbook:PPC64/Installation/Stage/fr "Handbook:PPC64/Installation/Stage/fr").

[← Configuring the network](/wiki/Handbook:PPC64/Installation/Networking/fr "Previous part") [Accueil](/wiki/Handbook:PPC64/fr "Handbook:PPC64/fr") [Installer l’archive _stage3_ →](/wiki/Handbook:PPC64/Installation/Stage/fr "Next part")