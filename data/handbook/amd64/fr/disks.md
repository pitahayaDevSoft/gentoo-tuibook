# Disks

Other languages:

- [Deutsch](/wiki/Handbook:AMD64/Installation/Disks/de "Handbook:AMD64/Installation/Disks/de (100% translated)")
- [English](/wiki/Handbook:AMD64/Installation/Disks "Handbook:AMD64/Installation/Disks (100% translated)")
- [Türkçe](/wiki/Handbook:AMD64/Installation/Disks/tr "Handbook:AMD64/Installation/Disks/tr (0% translated)")
- [español](/wiki/Handbook:AMD64/Installation/Disks/es "Handbook:AMD64/Installation/Disks/es (100% translated)")
- français
- [italiano](/wiki/Handbook:AMD64/Installation/Disks/it "Handbook:AMD64/Installation/Disks/it (100% translated)")
- [magyar](/wiki/Handbook:AMD64/Installation/Disks/hu "Handbook:AMD64/Installation/Disks/hu (100% translated)")
- [polski](/wiki/Handbook:AMD64/Installation/Disks/pl "Handbook:AMD64/Installation/Disks/pl (100% translated)")
- [português do Brasil](/wiki/Handbook:AMD64/Installation/Disks/pt-br "Handbook:AMD64/Installation/Disks/pt-br (100% translated)")
- [čeština](/wiki/Handbook:AMD64/Installation/Disks/cs "Handbook:AMD64/Installation/Disks/cs (100% translated)")
- [русский](/wiki/Handbook:AMD64/Installation/Disks/ru "Handbook:AMD64/Installation/Disks/ru (100% translated)")
- [العربية](/wiki/Handbook:AMD64/Installation/Disks/ar "Handbook:AMD64/Installation/Disks/ar (100% translated)")
- [தமிழ்](/wiki/Handbook:AMD64/Installation/Disks/ta "Handbook:AMD64/Installation/Disks/ta (100% translated)")
- [中文（中国大陆）‎](/wiki/Handbook:AMD64/Installation/Disks/zh-cn "Handbook:AMD64/Installation/Disks/zh-cn (100% translated)")
- [日本語](/wiki/Handbook:AMD64/Installation/Disks/ja "Handbook:AMD64/Installation/Disks/ja (100% translated)")
- [한국어](/wiki/Handbook:AMD64/Installation/Disks/ko "Handbook:AMD64/Installation/Disks/ko (100% translated)")

[Sommaire du manuel](/wiki/Handbook:AMD64 "Handbook:AMD64")[Installation](/wiki/Handbook:AMD64/Full/Installation/fr "Handbook:AMD64/Full/Installation/fr")[À propos de l’installation](/wiki/Handbook:AMD64/Installation/About/fr "Handbook:AMD64/Installation/About/fr")[Choix du support](/wiki/Handbook:AMD64/Installation/Media/fr "Handbook:AMD64/Installation/Media/fr")[Configurer le réseau](/wiki/Handbook:AMD64/Installation/Networking/fr "Handbook:AMD64/Installation/Networking/fr")Préparer les disques[Installer l’archive _stage3_](/wiki/Handbook:AMD64/Installation/Stage/fr "Handbook:AMD64/Installation/Stage/fr")[Installer le système de base](/wiki/Handbook:AMD64/Installation/Base/fr "Handbook:AMD64/Installation/Base/fr")[Configurer le noyau](/wiki/Handbook:AMD64/Installation/Kernel/fr "Handbook:AMD64/Installation/Kernel/fr")[Configurer le système](/wiki/Handbook:AMD64/Installation/System/fr "Handbook:AMD64/Installation/System/fr")[Installer les outils](/wiki/Handbook:AMD64/Installation/Tools/fr "Handbook:AMD64/Installation/Tools/fr")[Configurer le système d’amorçage](/wiki/Handbook:AMD64/Installation/Bootloader/fr "Handbook:AMD64/Installation/Bootloader/fr")[Finaliser](/wiki/Handbook:AMD64/Installation/Finalizing/fr "Handbook:AMD64/Installation/Finalizing/fr")[Utiliser Gentoo](/wiki/Handbook:AMD64/Full/Working/fr "Handbook:AMD64/Full/Working/fr")[Introduction à Portage](/wiki/Handbook:AMD64/Working/Portage/fr "Handbook:AMD64/Working/Portage/fr")[Les options de la variable USE](/wiki/Handbook:AMD64/Working/USE/fr "Handbook:AMD64/Working/USE/fr")[Les fonctionnalités de Portage](/wiki/Handbook:AMD64/Working/Features/fr "Handbook:AMD64/Working/Features/fr")[Scripts d’initialisation systèmes](/wiki/Handbook:AMD64/Working/Initscripts/fr "Handbook:AMD64/Working/Initscripts/fr")[Variables d’environnement](/wiki/Handbook:AMD64/Working/EnvVar/fr "Handbook:AMD64/Working/EnvVar/fr")[Utiliser Portage](/wiki/Handbook:AMD64/Full/Portage/fr "Handbook:AMD64/Full/Portage/fr")[Fichiers et répertoires](/wiki/Handbook:AMD64/Portage/Files/fr "Handbook:AMD64/Portage/Files/fr")[Les variables](/wiki/Handbook:AMD64/Portage/Variables/fr "Handbook:AMD64/Portage/Variables/fr")[Mélanger plusieurs branches logicielles](/wiki/Handbook:AMD64/Portage/Branches/fr "Handbook:AMD64/Portage/Branches/fr")[Outils supplémentaires](/wiki/Handbook:AMD64/Portage/Tools/fr "Handbook:AMD64/Portage/Tools/fr")[Dépôt personnalisé](/wiki/Handbook:AMD64/Portage/CustomTree/fr "Handbook:AMD64/Portage/CustomTree/fr")[Fonctionnalités avancées](/wiki/Handbook:AMD64/Portage/Advanced/fr "Handbook:AMD64/Portage/Advanced/fr")[Configuration du réseau avec OpenRC](/wiki/Handbook:AMD64/Full/Networking/fr "Handbook:AMD64/Full/Networking/fr")[Bien démarrer](/wiki/Handbook:AMD64/Networking/Introduction/fr "Handbook:AMD64/Networking/Introduction/fr")[Configuration avancée](/wiki/Handbook:AMD64/Networking/Advanced/fr "Handbook:AMD64/Networking/Advanced/fr")[Les modules réseau](/wiki/Handbook:AMD64/Networking/Modular/fr "Handbook:AMD64/Networking/Modular/fr")[Sans fil](/wiki/Handbook:AMD64/Networking/Wireless/fr "Handbook:AMD64/Networking/Wireless/fr")[Ajouter des fonctionnalités](/wiki/Handbook:AMD64/Networking/Extending/fr "Handbook:AMD64/Networking/Extending/fr")[Gestion dynamique](/wiki/Handbook:AMD64/Networking/Dynamic/fr "Handbook:AMD64/Networking/Dynamic/fr")

## Contents

- [1Introduction aux périphériques de type bloc](#Introduction_aux_p.C3.A9riph.C3.A9riques_de_type_bloc)
  - [1.1Les périphériques de type bloc](#Les_p.C3.A9riph.C3.A9riques_de_type_bloc)
  - [1.2Tables de partition](#Tables_de_partition)
    - [1.2.1Table de partitions _GUID_ ( _GPT_)](#Table_de_partitions_GUID_.28GPT.29)
    - [1.2.2_Master boot record_ (MBR) or secteur de démarrage _DOS_](#Master_boot_record_.28MBR.29_or_secteur_de_d.C3.A9marrage_DOS)
  - [1.3Stockage avancé](#Stockage_avanc.C3.A9)
  - [1.4Schéma de partitionnement par défaut](#Sch.C3.A9ma_de_partitionnement_par_d.C3.A9faut)
- [2Concevoir un schéma de partitionnement](#Concevoir_un_sch.C3.A9ma_de_partitionnement)
  - [2.1Combien de partition et de quelle taille ?](#Combien_de_partition_et_de_quelle_taille_.3F)
  - [2.2Qu’en est-il de l’espace d’échange ( _swap_) ?](#Qu.E2.80.99en_est-il_de_l.E2.80.99espace_d.E2.80.99.C3.A9change_.28swap.29_.3F)
  - [2.3Qu’est-ce que l’ _EFI System Partition_ ( _ESP_, système de partition _EFI_) ?](#Qu.E2.80.99est-ce_que_l.E2.80.99EFI_System_Partition_.28ESP.2C_syst.C3.A8me_de_partition_EFI.29_.3F)
  - [2.4Qu’est-ce qu’une partition _BIOS_ ?](#Qu.E2.80.99est-ce_qu.E2.80.99une_partition_BIOS_.3F)
- [3Partitionner le disque avec _GPT_ pour _UEFI_](#Partitionner_le_disque_avec_GPT_pour_UEFI)
  - [3.1Visualiser le plan de partitionnement actuel](#Visualiser_le_plan_de_partitionnement_actuel)
  - [3.2Créer un nouveau libellé disque / supprimer toutes les partitions](#Cr.C3.A9er_un_nouveau_libell.C3.A9_disque_.2F_supprimer_toutes_les_partitions)
  - [3.3Créer une partition système _EFI_ (ESP)](#Cr.C3.A9er_une_partition_syst.C3.A8me_EFI_.28ESP.29)
  - [3.4Créer la partition d’échange](#Cr.C3.A9er_la_partition_d.E2.80.99.C3.A9change)
  - [3.5Créer la partition racine](#Cr.C3.A9er_la_partition_racine)
  - [3.6Sauvegarder la table de partitions](#Sauvegarder_la_table_de_partitions)
- [4Partitionner un disque avec _MBR/BIOS_](#Partitionner_un_disque_avec_MBR.2FBIOS)
  - [4.1Visualiser la table de partitionnements actuelle](#Visualiser_la_table_de_partitionnements_actuelle)
  - [4.2Créer un nouveau libellé disque / supprimer toutes les partitions](#Cr.C3.A9er_un_nouveau_libell.C3.A9_disque_.2F_supprimer_toutes_les_partitions_2)
  - [4.3Créer une partition d’amorçage](#Cr.C3.A9er_une_partition_d.E2.80.99amor.C3.A7age)
- [5Créer des systèmes de fichiers](#Cr.C3.A9er_des_syst.C3.A8mes_de_fichiers)
- [6Introduction](#Introduction)
  - [6.1Les systèmes de fichiers](#Les_syst.C3.A8mes_de_fichiers)
  - [6.2Mettre en œuvre un système de fichiers sur une partition](#Mettre_en_.C5.93uvre_un_syst.C3.A8me_de_fichiers_sur_une_partition)
    - [6.2.1Système de fichiers pour partition EFI](#Syst.C3.A8me_de_fichiers_pour_partition_EFI)
    - [6.2.2Système de fichiers pour un BIOS déprécié](#Syst.C3.A8me_de_fichiers_pour_un_BIOS_d.C3.A9pr.C3.A9ci.C3.A9)
    - [6.2.3Petite partition ext4](#Petite_partition_ext4)
  - [6.3Activer la partition d’échange](#Activer_la_partition_d.E2.80.99.C3.A9change)
- [7Monter la partition racine](#Monter_la_partition_racine)

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

### Tables de partition

Bien qu’il soit théoriquement possible d’utiliser un disque brut et non partitionné pour héberger un système Linux (lors de la création d’un _RAID btrfs_ par exemple), cela n’est réellement jamais fait. Les périphériques de bloc de disque sont scindés en blocs plus petits, plus faciles à gérer. Pour l’architecture **amd64**, on appelle ces blocs des partitions. Deux technologies de partitionnement standard sont actuellement disponibles : _MBR_ (parfois appelé label de disque _DOS_) et _GPT_ ; les deux sont liés au processus d’amorçage déprécié _BIOS_ ou _UEFI_.

#### Table de partitions _GUID_ ( _GPT_)

La table de partitions _GUI_ ( _GPT_) utilise des identifiants 64 bits pour les partitions. L’emplacement dans lequel elle stocke les informations de partition est beaucoup plus grand que les 512 octets d’une partition _MBR_ (libellé _DOS_), ce qui signifie qu’il n’y a pratiquement aucune limite sur le nombre de partitions d’un disque GPT. De plus, la taille maximale d’une partition est plus grande (presque 8 Zio – oui, zébioctets).

Lorsque l’interface logicielle entre le système d’exploitation et le micrologiciel est _UEFI_ (au lieu du _BIOS_), _GPT_ est presque obligatoire car des problèmes de compatibilité surviennent avec les anciens libellés _DOS_.

_GPT_ profite également de l’utilisation de la somme de contrôle et de la redondance. Il utilise les sommes de contrôle _CRC32_ pour détecter les erreurs dans les tables d’en-tête et de partition et dispose d’une sauvegarde _GPT_ en fin de disque. Cette table de sauvegarde peut être utilisée pour réparer les dommages subis par le _GPT_ principal situé au début du disque.

**Important**

Il y a quelques mises en garde à propos de _GPT_ :

- utiliser _GPT_ sur un ordinateur basé sur _BIOS_ fonctionne, mais l’utilisateur ne pourra pas faire un double amorçage avec un système d’exploitation Microsoft Windows, car Microsoft Windows refuse de démarrer avec une partition _GPT_ en mode _BIOS_ ;
- certaines vieilles cartes mères configurées pour démarrer en mode _BIOS/CSM/legacy_ peuvent aussi rencontrer des problèmes pour démarrer avec des disques _GPT_.

#### _Master boot record_ (MBR) or secteur de démarrage _DOS_

Le secteur de démarrage _[Master boot record](https://en.wikipedia.org/wiki/Master_boot_record "wikipedia:Master boot record")_ (aussi nommé secteur de démarrage _DOS_, libellé disque _DOS_ et plus récemment, par contraste avec _GPT/UEFI_, amorçage déprécié _BIOS_) a été introduit en 1983 avec _PC DOS 2.x MBR_ utilise des identifiants 32 bits pour le secteur de démarrage et la longueur des partitions et prend en charge trois types de partitions : primaire, étendue et logique. Les partitions primaires stockent leurs informations directement dans le MBR – un très petit emplacement (généralement 512 octets) au tout début d’un disque. En raison de cet espace restreint, seules quatre partitions primaires sont prises en charge (par exemple, /dev/sda1 à /dev/sda4).

Pour prendre en charge davantage de partitions, l’une des partitions primaires du _MBR_ peut être marquée en tant que partition « étendue ». Cette partition peut alors contenir des partitions logiques supplémentaires (des partitions dans une partition).

**Important**

Bien que toujours prises en charge par la plupart des fabricants de cartes mères, les tables de partitions _MBR_ sont considérées comme anciennes. Si vous ne travaillez pas avec du matériel antérieur à 2010, il est préférable de partitionner un disque à l’aide d’une [Table de partition GUID](https://en.wikipedia.org/wiki/GUID_Partition_Table "wikipedia:GUID Partition Table"). Les lecteurs qui veulent poursuivre avec _MBR_ doivent connaître les informations suivantes :

- La plupart des cartes mères sorties après 2010 considèrent le _MBR_ comme un mode de démarrage déprécié (pris en charge, mais pas idéal).
- En raison de l’utilisation d'identificateurs 32 bits, le _MBR_ ne peut pas gérer les disques dont la taille est supérieure à 2 Tio.
- À moins de créer une partition étendue, le _MBR_ prend en charge un maximum de quatre partitions.
- La configuration du _MBR_ ne fournit aucun _MBR_ de sauvegarde. Par conséquent, si une application ou un utilisateur écrase le MBR, toutes les informations sur la partition sont perdues.

Ceci étant dit, _MBR_ et les amorçages _BIOS_ sont encore fréquemment utilisés dans les environnements virtualisés dans le nuage, comme AWS.

Les auteurs de ce manuel suggèrent d’utiliser [GPT](#GPT) autant que possible pour les installations de Gentoo.

### Stockage avancé

Le média officiel Gentoo fournit le support des gestionnaires de volume logique [_Logical Volume Manager (LVM)_](/wiki/LVM/fr "LVM/fr"). _LVM_ peut combiner des volumes physiques ( _physical volume_) comme des partitions ou des disques dans des groupes de volume ( _volume groups_). Les groupes de volume sont plus flexibles que des partitions et peuvent être utilisés pour définir des groupes _RAID_ ou de cache dans des disques _SSD_ rapides pour des disques lents. Bien que cet usage ne soit pas couvert par le manuel, _LVM_ est totalement supporté par Gentoo.

Although usage is not covered in the handbook, below is a list helpful guides to get the system running:

- [Btrfs](/wiki/Btrfs/Native_System_Root_Guide "Btrfs/Native System Root Guide")
- [ZFS](/wiki/ZFS/rootfs "ZFS/rootfs")
- [LVM (Logical Volume Manager)](/wiki/LVM/fr "LVM/fr")

Disk encryption is also handled in the same manner:

- [Rootfs encryption](/wiki/Rootfs_encryption/fr "Rootfs encryption/fr")

### Schéma de partitionnement par défaut

Tout au long de ce manuel, nous allons expliquer les deux cas :

1. _UEFI_ avec disque _GPT_.
2. _MBR DOS/BIOS_ déprécié avec table de partitions _MBR_.

Bien qu’il soit possible de mixer les types d’amorçage avec certaines cartes mères, ceci dépasse les intentions du manuel. Comme indiqué précédemment, il est fortement recommandé avec du matériel moderne d’utiliser un démarrage _UEFI_ avec un libellé disque _GPT_.

Le schéma de partitionnement suivant sera utilisé comme exemple simple :

**Important**

La première ligne du tableau suivant contient des informations _**soit**_ pour un libellé _GPT_ _**ou**_ un libellé _MBR/BIOS_. En cas de doute, suivez _GPT_ car les machines **amd64** après 2010 supportent généralement _UEFI_ et l’amorçage _GPT_.

Partition
Système de fichiers
Taille
Description
/dev/sda1fat32 Le système de fichiers nécessaire pour un système _EFI_, toujours associé à un libellé _GPT_.1 Gio
Système de partition _EFI_. _Applicable pour les systèmes supportant une implémentation UEFI. Typiquement le cas des systèmes conçus à partir des années 2010._xfs Système de fichiers recommandé pour amorcer avec une table de partition _MBR_, utilisé en complément d’étiquettes _DOS/BIOS_ dépréciées.Détails _MBR DOS/BIOS_. _Applicable pour les machines_ BIOS _dépréciées. Les systèmes de ce type sont typiquement fabriqués <u>avant</u> les années 2010 et ne sont plus en production._/dev/sda2linux-swap
RAM × 2
Détails de partition d’échange
/dev/sda3xfs
Reste du disque **Le _profil_ sélectionné, les _partitions_ supplémentaires (facultatif) et le _but_ du système ajoute de la complexité pour dimensionner de manière appropriée la partition racine, aussi les auteurs du manuel ne peuvent proposer une règle correspondant à tous les usages</br></br>Lorsque Gentoo est le seul système à utiliser le disque, utiliser tout l’espace restant disponible est le choix le plus simple et rapide.**Détails de la partition racine.

Si cela est suffisant, le lecteur avancé peut immédiatement passer au partitionnement.

fdisk et parted sont des utilitaires de partitionnement inclus dans les images Gentoo officielles. fdisk est le plus connu, stable et support autant _MBR_ que _GPT_. parted était l’un des premiers outils à gérer les périphériques de stockage à supporter _GPT_. Il peut être utiliser à la place de fdisk si le lecteur préfère, cependant le manuel ne donnera que les commandes pour fdisk car c’est le plus courant dans les environnements Linux.

Avant de donner les instructions de création, la première section va décrire en détail comment créer un schéma de partition et les erreurs les plus courantes.

## Concevoir un schéma de partitionnement

### Combien de partition et de quelle taille ?

Le nombre de partitions dépend fortement des besoins systèmes et des systèmes de fichiers utilisés pour les périphériques. S’il y a beaucoup d’utilisateurs, il est alors conseillé de séparer /home/ afin d’augmenter la sécurité et de faciliter les sauvegardes et la maintenance. Si Gentoo est installé en tant que serveur de courriel, il est préférable de séparer /var/ car c’est là que sont stockés les courriels. Les serveurs de jeux devraient séparer /opt/, car c’est là que la plupart de ces logiciels sont installés. Les raisons de ces recommandations sont similaires à /home : sécurité, sauvegarde et maintenance.

Dans la majorité des situations /usr et /var doivent avoir une grande taille : /usr contient la majorité des applications du système et les sources du noyau (dans /usr/src). Par défaut, /var héberge le dépôt _ebuild_ de Gentoo (par défaut situé à /var/db/repos/gentoo), qui prend déjà environ 650 Mio d’espace disque. Cette estimation exclut les répertoires /var/cache/distfiles et /var/cache/binpkgs qui se remplissent au fur et à mesure avec le code source et les paquets binaires optionnels lorsqu’ils sont ajoutés au système.

Combien de partitions et quelles tailles dépend de compromis et le meilleur choix des circonstances. Séparer les partitions ou volumes ont les avantages suivants :

- choix du système de fichiers le plus performant pour chaque partition ou volume ;
- le système entier ne peut pas se retrouver à cours d’espace libre si un outil écrit continuellement des fichiers sur une partition ou un volume ;
- si nécessaire, les contrôles du système de fichiers sont plus rapides car plusieurs contrôles peuvent être effectués en parallèle (bien que cet avantage soit plus vrai avec plusieurs disques qu’avec plusieurs partitions) ;
- la sécurité peut être améliorée en montant certaines partitions ou volumes en lecture seule, `nosuid` (les bits _setuid_ sont ignorés), `noexec` (les bits exécutables sont ignorés), etc.

Cependant, avoir plusieurs partitions présente également des inconvénients :

- si elles ne sont pas configurées proprement, le système peut se retrouver avec beaucoup d’espace libre sur une partition et aucun sur une autre ;
- une partition séparée pour /usr/ nécessitent souvent que l’administrateur amorce le système avec un fichier _initramfs_ afin de monter la partition avant que d’autres scripts de démarrage ; comme la création et la maintenance d’un _initramfs_ sort du périmètre de ce manuel, **il est recommandé aux nouveaux utilisateurs de ne pas utiliser une partition /usr/ séparée** ;
- il existe également une limite de 15 partitions pour _SCSI_ et _SATA_ à moins que le disque n’utilise les étiquettes _GPT_.

**Remarque**

Les installations avec _systemd_ doivent avoir /usr disponible lors de l’amorçage, comme partition racine ou monté via _initramfs_.

### Qu’en est-il de l’espace d’échange ( _swap_) ?

Recommandations pour l’espace d’échange
Taille RAMSupport de la suspension ?Support de l’hibernation ?
2 Gio ou moins2 × RAM3 × RAM
2 à 8 GioRAM2 × RAM
8 à 64 Gio8 à 16 Gio1,5 × RAM
64 Gio ou plus8 Gio minimumHibernation non recommandés ! L’hibernation **n’est pas** recommandée pour les systèmes qui ont une très grande quantité de mémoire vive. Même s’il est possible d’écrire l’ensemble du contenu de la mémoire sur le disque pour accomplir l’hibernation, écrire des dizaines de gibioctets (ou plus !) sur disque prend un temps considérable, particulièrement sur des disques à rotation. Il est mieux de seulement suspendre le système dans ce scénario.

Il n’y a pas de valeur parfaite pour la partition d’échange. Le but de l’espace d’échange est de fournir un espace de stockage sur disque au noyau quand la mémoire vive interne ( _RAM_) est sous pression. Cet espace permet au noyau de déplacer des pages de mémoire qui sont susceptibles de na pas être utilisées prochainement sur le disque ( _swap_ ou _page-out_), libérant ainsi de la mémoire vive _RAM_ pour les tâches actives. Bien sûr, si ces pages déplacées sont soudainement nécessaires, elles doivent être remisent en mémoire vive ( _page-in_) ce qui prendra un temps considérablement plus long que lire depuis la mémoire _RAM_ (car les disque sont très lents en comparaison).

Si le système n’est pas prévu pour exécuter des applications utilisant beaucoup de mémoire ou si le système a beaucoup de mémoire disponible, il n’est alors probablement pas nécessaire d’avoir beaucoup d’espace d’échange. Cependant, l’espace d’échange est aussi utilisé pour stocker toute la mémoire en cas d’hibernation (plutôt un ordinateur de bureau ou portable qu’un serveur). Si le système doit supporter l’hibernation, alors un espace d’échange supérieur à au moins la quantité de mémoire installée sur le système est nécessaire.

La règle générale est que pour une mémoire _RAM_ inférieure à 4 Gio, l’espace d’échange devrait être le double. Pour les systèmes avec plusieurs disques durs, il peut être judicieux d’en créer par disque, ainsi les opérations de lecture et écriture pourront être parallélisées. Au plus les disques peuvent écrire dans la _swap_, au plus le système sera rapide dans l’accès de ce type de données. Entre un disque dur rotatif ( _HDD_) et un disque flash ( _SSD_), il est mieux de position la partition d’échange sur un disque flash.

Il est à noter que des **fichiers** d’échange peuvent être utilisés au lieu des **partitions** ; cela peut être utile sur des systèmes avec un espace disque limité.

### Qu’est-ce que l’ _EFI System Partition_ ( _ESP_, système de partition _EFI_) ?

Lors de l’installation de Gentoo sur un système utilisant _UEFI_ pour démarrer le système d’exploitation (au lieu de BIOS), il est important de créer une partition système _EFI_ ( _ESP_). Les instructions pour parted ci-dessous contiennent les indications nécessaires à la bonne réalisation de cette opération. **La partition système _EFI_ n’est pas nécessaire en mode _BIOS_ déprécié.**

L’ESP _doit_ être une variante _FAT_ (parfois indiquée par _vfat_ sur les systèmes Linux). La [spécification UEFI (EN)](http://www.uefi.org/sites/default/files/resources/UEFI%202_5.pdf) officielle indique que les systèmes de fichiers FAT12, 16 ou 32 seront reconnus par le microprogramme _UEFI_, bien que _FAT32_ soit recommandé pour l’ _ESP_. Après la création de la partition, formater l’ _ESP_ avec :

`root #` `mkfs.fat -F 32 /dev/sda1`

**Important**

Si une variante FAT n’est pas utilisée pour l’ESP, le micrologiciel UEFI du système n’est pas sûr de trouver le chargeur d’amorçage (ou le noyau Linux) et ne sera probablement pas en mesure de démarrer le système !

### Qu’est-ce qu’une partition _BIOS_ ?

Une partition d’amorçage _BIOS_ est très petite (1 à 2 Mio), les programmes d’amorçage comme _GRUB2_ mettent de la données additionnelles qui ne rentrent par dans l’espace d’allocation. C’est seulement nécessaire pour un disque formaté avec un libellé _GPT_ où le système va être démarré avec _GRUB2_ dans un mode déprécié _MBR/BIOS DOS_. **Il n’est pas requis pour démarrer en mode _EFI/UEFI_ et également avec un libellé _MBR/DOS_ déprécié.**
Une partition d’amorçage _BIOS_ ne sera pas utilisé dans ce guide.

## Partitionner le disque avec _GPT_ pour _UEFI_

Cette partie explique comment créer une partition pour un simple disque _GPT_ conformément aux spécifications _UEFI_ et _DPS_ ( _Discoverable Partitions Specification_, spécification de partitions découvrables). _DPS_ est une spécification fournie par l’espace utilisateur ( _userspace_) Linux ( _UAPI_) et est recommandé, mais entièrement facultatif. Les spécifications sont implémentés dans l’utilitaire fdisk, lequel fait partie du paquet [sys-apps/util-linux](https://packages.gentoo.org/packages/sys-apps/util-linux).

Le tableau fourni des valeurs par défaut recommandées pour une installation triviale Gentoo. Des partitions supplémentaires peuvent être ajoutées en fonction de préférence personnelles ou de besoins systèmes.

Chemin de périphérique
Point de montage
Système de fichiers
_DPS UUID_ (_Type-UUID_)
Description
/dev/sda1/efivfat
c12a7328-f81f-11d2-ba4b-00a0c93ec93b
Détails de la partition système _EFI_ (ESP).
/dev/sda2N/A La partition d’échange ( _swap_) n’est pas montée dans le système de fichiers comme un périphérique.swap
0657fd6d-a4ab-43c4-84e5-0933c84b4f4f
Détails de la partition d’échange.
/dev/sda3/xfs
4f68bce3-e8cd-4db1-96e7-fbcaf984b709
Détails de la partition racine.

### Visualiser le plan de partitionnement actuel

fdisk est un outil populaire et puissant pour créer des partitions sur un disque. Lancez fdisk sur un disque (dans notre exemple, nous utilisons /dev/sda) :

`root #` `fdisk /dev/sda`

Utilisez la touche `p` pour afficher les partitions actuelles du disque :

`Command (m for help):` `p`

```
Disk /dev/sda: 931.51 GiB, 1000204886016 bytes, 1953525168 sectors
Disk model: HGST HTS721010A9
Units: sectors of 1 * 512 = 512 bytes
Sector size (logical/physical): 512 bytes / 4096 bytes
I/O size (minimum/optimal): 4096 bytes / 4096 bytes
Disklabel type: gpt
Disk identifier: 3E56EE74-0571-462B-A992-9872E3855D75

Device        Start        End    Sectors   Size Type
/dev/sda1      2048    2099199    2097152     1G EFI System
/dev/sda2   2099200   10487807    8388608     4G Linux swap
/dev/sda3  10487808 1953523711 1943035904 926.5G Linux root (x86-64)

```

Ce disque en particulier est configuré pour héberger deux systèmes de fichiers Linux (chacun correspondant à une partition « Linux ») et une partition d’échange (correspondant à « Linux swap »).

### Créer un nouveau libellé disque / supprimer toutes les partitions

La touche `g` va supprimer instantanément toutes les partitions du disque et créer un libellé _GPT_ :

`Command (m for help):` `g`

```
Created a new GPT disklabel (GUID: 3E56EE74-0571-462B-A992-9872E3855D75).

```

Alternativement, pour conserver les libellés _GPT_ existants (voir la sortie de `p` ci-dessus), envisagez de supprimer les partitions existantes une par une du disque. Utilisez `d` pour effacer une partition. Par exemple, pour supprimer /dev/sda1 :

`Command (m for help):` `d`

```
Partition number (1-4): 1

```

La partition est maintenant planifiée pour une suppression. Elle ne sera plus visible dans la liste des partitions ( `p`), mais ne sera pas effacée tant que les changements n’auront pas été sauvegardés. Cela permet à l’utilisateur d’annuler l’opération si une erreur est commise – dans ce cas, tapez `q` immédiatement et `Entrée` pour ne pas supprimer la partition.

Presser de manière répétée `p` pour afficher la liste des partitions et `d` avec le numéro pour la supprimer. Éventuellement, la table de partition peut-être vide :

`Command (m for help):` `p`

```
Disk /dev/sda: 931.51 GiB, 1000204886016 bytes, 1953525168 sectors
Disk model: HGST HTS721010A9
Units: sectors of 1 * 512 = 512 bytes
Sector size (logical/physical): 512 bytes / 4096 bytes
I/O size (minimum/optimal): 4096 bytes / 4096 bytes
Disklabel type: gpt
Disk identifier: 3E56EE74-0571-462B-A992-9872E3855D75

```

Maintenant que la table de partition en mémoire est vide, il est possible de créer des partitions.

### Créer une partition système _EFI_ (ESP)

**Remarque**

Un _ESP_ plus petit est possible mais pas recommandé car il peut être partagé avec d’autres systèmes d’exploitation.

Créez une petite partition _EFI_, laquelle sera montée dans /efi. Tapez `n` pour créer la nouvelle partition, suivi de `1` pour choisir la première partition. Lors de la demande du premier secteur, assurez-vous de commencer de 2048 (ce qui peut être nécessaire pour l’amorçage) et tapez `Enter`. Pour le dernier secteur, choisissez « +1G » pour créer une partition d’un gibioctet :

`Command (m for help):` `n`

```
Partition number (1-128, default 1): 1
First sector (2048-1953525134, default 2048):
Last sector, +/-sectors or +/-size{K,M,G,T,P} (2048-1953525134, default 1953523711): +1G

Created a new partition 1 of type 'Linux filesystem' and of size 1 GiB.
Partition #1 contains a vfat signature.

Do you want to remove the signature? [Y]es/[N]o: Y
The signature will be removed by a write command.

```

Marquez la partition comme _EFI_ :

`Command (m for help):` `t`

```
Selected partition 1
Partition type or alias (type L to list all): 1
Changed type of partition 'Linux filesystem' to 'EFI System'.

```

### Créer la partition d’échange

Ensuite, pour créer la partition d’échange, pressez `n`, puis `2` pour créer la deuxième partition, /dev/sda2. Lorsque demandé le premier secteur, tapez `Entrée`. Lorsque demandé le dernier secteur, saisissez « +4G » (ou toute taille voulue pour la partition d’échange) pour créer une partition de 4 gibioctets.

`Command (m for help):` `n`

```
Partition number (2-128, default 2):
First sector (2099200-1953525134, default 2099200):
Last sector, +/-sectors or +/-size{K,M,G,T,P} (2099200-1953525134, default 1953523711): +4G

Created a new partition 2 of type 'Linux filesystem' and of size 4 GiB.

```

Après cela, tapez `t` pour définir le type de partition, `2` pour choisir la partition tout juste créée et tapez « 19 » pour définir le type à « Linux Swap ».

`Command (m for help):` `t`

```
Partition number (1,2, default 2): 2
Partition type or alias (type L to list all): 19

Changed type of partition 'Linux filesystem' to 'Linux swap'.

```

### Créer la partition racine

Finalement, pour créer la partition racine, pressez `n`, puis `3` pour créer une troisième partition : /dev/sda3. Lorsque interrogé sur le premier secteur, pressez `Entrée`. Lorsque interrogé sur le dernier secteur, tapez `Entrée` pour créer une partition qui prend l’espace disque restant disponible sur le disque.

`Command (m for help):` `n`

```
Partition number (3-128, default 3): 3
First sector (10487808-1953525134, default 10487808):
Last sector, +/-sectors or +/-size{K,M,G,T,P} (10487808-1953525134, default 1953523711):

Created a new partition 3 of type 'Linux filesystem' and of size 926.5 GiB..

```

**Remarque**

Définir le type de partition racine à « Linux root (x86-64) » n’est pas nécessaire et le système fonctionnera normalement s’il est paramétré sur « Linux filesystem ». Le type de système de fichiers est seulement nécessaire si une application d’amorçage le supportant est utilisée (comme systemd-boot) et qu’un fichier fstab n’est pas souhaité.

Après avoir créer une partition racine, pressez `t` pour définir le type de partition, `3` pour sélectionner celle tout juste créée et tapez « 23 » pour définir le type de partition à « Linux Root (x86-64) ».

`Command(m for help):` `t`

```
Partition number (1-3, default 3): 3
Partition type or alias (type L to list all): 23

Changed type of partition 'Linux filesystem' to 'Linux root (x86-64)'

```

Après avoir accompli toutes ces étapes, pressez `p` pour afficher une table de partition qui ressemblera à :

`Command (m for help):` `p`

```
Disk /dev/sda: 931.51 GiB, 1000204886016 bytes, 1953525168 sectors
Disk model: HGST HTS721010A9
Units: sectors of 1 * 512 = 512 bytes
Sector size (logical/physical): 512 bytes / 4096 bytes
I/O size (minimum/optimal): 4096 bytes / 4096 bytes
Disklabel type: gpt
Disk identifier: 3E56EE74-0571-462B-A992-9872E3855D75

Device        Start        End    Sectors   Size Type
/dev/sda1      2048    2099199    2097152     1G EFI System
/dev/sda2   2099200   10487807    8388608     4G Linux swap
/dev/sda3  10487808 1953523711 1943035904 926.5G Linux root (x86-64)

Filesystem/RAID signature on partition 1 will be wiped.

```

### Sauvegarder la table de partitions

Pressez `w` pour enregistrer la table de partitions et quitter l’utilitaire fdisk :

`Command (m for help):` `w`

```
The partition table has been altered.
Calling ioctl() to re-read partition table.
Syncing disks.

```

Maintenant que les partitions sont disponibles, la suite de l’installation est de remplir le système de fichiers.

## Partitionner un disque avec _MBR/BIOS_

Le tableau suivant propose une table de partition recommandée pour une installation classique _MBR DOS/BIOS_. Les partitions additionnelles peuvent être utilisées pour des préférences personnelles ou des besoins systèmes.

Chemin de périphérique
Point de montage
Système de fichiers
_DPS UUID_ (_PARTUUID_)
Description
/dev/sda1/bootxfs
N/A
Détails de la partition _MBR DOS/BIOS_ dépréciée.
/dev/sda2N/A La partition d’échange ( _swap_) n’est pas montée dans le système de fichiers comme un périphérique.swap
0657fd6d-a4ab-43c4-84e5-0933c84b4f4f
Détails de la partition d’échange.
/dev/sda3/xfs
4f68bce3-e8cd-4db1-96e7-fbcaf984b709
Détails de la partition racine.

Adaptez la table de partitions à votre convenance personnelle.

### Visualiser la table de partitionnements actuelle

Lancez fdisk pour le disque concerné (dans notre exemple /dev/sda) :

`root #` `fdisk /dev/sda`

Utilisez la touche `p` pour afficher les partitions actuelles du disque :

`Command (m for help):` `p`

```
Disk /dev/sda: 931.51 GiB, 1000204886016 bytes, 1953525168 sectors
Disk model: HGST HTS721010A9
Units: sectors of 1 * 512 = 512 bytes
Sector size (logical/physical): 512 bytes / 4096 bytes
I/O size (minimum/optimal): 4096 bytes / 4096 bytes
Disklabel type: dos
Disk identifier: 0xf163b576

Device     Boot    Start        End    Sectors   Size Id Type
/dev/sda1  *        2048    2099199    2097152     1G 83 Linux
/dev/sda2        2099200   10487807    8388608     4G 82 Linux swap / Solaris
/dev/sda3       10487808 1953525167 1943037360 926.5G 83 Linux

```

Ce disque en particulier est configuré pour héberger deux systèmes de fichiers Linux (chacun correspondant à une partition « Linux ») et une partition d’échange (correspondant à « Linux swap ») en utilisant une table _GPT_.

### Créer un nouveau libellé disque / supprimer toutes les partitions

La touche `o` va supprimer instantanément toutes les partitions du disque et créer un libellé _MBR_ (aussi appelé libellé _DOS_) :

`Command (m for help):` `o`

```
Created a new DOS disklabel with disk identifier 0xf163b576.
The device contains 'gpt' signature and it will be removed by a write command. See fdisk(8) man page and --wipe option for more details.

```

Alternativement, pour conserver les libellés _DOS_ existants (voir la sortie de `p` ci-dessus), envisagez de supprimer les partitions existantes une par une du disque. Utilisez `d` pour effacer une partition. Par exemple, pour supprimer /dev/sda1 :

`Command (m for help):` `d`

```
Partition number (1-4): 1

```

La partition est maintenant planifiée pour une suppression. Elle ne sera plus visible dans la liste des partitions ( `p`), mais ne sera pas effacée tant que les changements n’auront pas été sauvegardés. Cela permet à l’utilisateur d’annuler l’opération si une erreur est commise – dans ce cas, tapez `q` immédiatement et `Entrée` pour ne pas supprimer la partition.

Presser de manière répétée `p` pour afficher la liste des partitions et `d` avec le numéro pour la supprimer. Éventuellement, la table de partition peut-être vide :

`Command (m for help):` `p`

```
Disk /dev/sda: 931.51 GiB, 1000204886016 bytes, 1953525168 sectors
Disk model: HGST HTS721010A9
Units: sectors of 1 * 512 = 512 bytes
Sector size (logical/physical): 512 bytes / 4096 bytes
I/O size (minimum/optimal): 4096 bytes / 4096 bytes
Disklabel type: dos
Disk identifier: 0xf163b576

```

Maintenant que la table de partition en mémoire est vide, il est possible de créer des partitions.

### Créer une partition d’amorçage

Créez une petite partition laquelle sera montée dans /boot. Tapez `n` pour créer la nouvelle partition, suivi de `p` pour une partition « primaire » et `1` pour choisir la première partition. Lors de la demande du premier secteur, assurez-vous de commencer de 2048 (ce qui peut être nécessaire pour l’amorçage) et tapez `Enter`. Pour le dernier secteur, choisissez « +1G » pour créer une partition d’un gibioctet :

`Command (m for help):` `n`

```
Partition type
   p   primary (0 primary, 0 extended, 4 free)
   e   extended (container for logical partitions)
Select (default p): p
Partition number (1-4, default 1): 1
First sector (2048-1953525167, default 2048):
Last sector, +/-sectors or +/-size{K,M,G,T,P} (2048-1953525167, default 1953525167): +1G

Created a new partition 1 of type 'Linux' and of size 1 GiB.

```

Marquez la partition comme amorçable en utilisant `a` et `Entrée` :

`Command (m for help):` `a`

```
Selected partition 1
The bootable flag on partition 1 is enabled now.

```

Note : si plus d’une partition est disponible sur le disque, alors la partition qui doit avoir le drapeau « amorçable » doit être choisie.

Créer la partition d’échange

Ensuite, pour créer la partition d’échange, pressez `n`, puis `2` pour créer la deuxième partition, /dev/sda2. Lorsque demandé le premier secteur, tapez `Entrée`. Lorsque demandé le dernier secteur, saisissez « +4G » (ou toute taille voulue pour la partition d’échange) pour créer une partition de 4 gibioctets.

`Command (m for help):` `n`

```
Partition type
   p   primary (1 primary, 0 extended, 3 free)
   e   extended (container for logical partitions)
Select (default p): p
Partition number (2-4, default 2): 2
First sector (2099200-1953525167, default 2099200):
Last sector, +/-sectors or +/-size{K,M,G,T,P} (2099200-1953525167, default 1953525167): +4G

Created a new partition 2 of type 'Linux' and of size 4 GiB.

```

Après cela, tapez `t` pour définir le type de partition, `2` pour choisir la partition tout juste créée et tapez « 82 » pour définir le type à « Linux Swap ».

`Command (m for help):` `t`

```
Partition number (1,2, default 2): 2
Hex code (type L to list all codes): 82

Changed type of partition 'Linux' to 'Linux swap / Solaris'.

```

Créer la partition racine

Finalement, pour créer la partition racine, pressez `n`, puis `3` pour créer une troisième partition : /dev/sda3. Lorsque interrogé sur le premier secteur, pressez `Entrée`. Lorsque interrogé sur le dernier secteur, tapez `Entrée` pour créer une partition qui prend l’espace disque restant disponible sur le disque.

`Command (m for help):` `n`

```
Partition type
   p   primary (2 primary, 0 extended, 2 free)
   e   extended (container for logical partitions)
Select (default p): p
Partition number (3,4, default 3): 3
First sector (10487808-1953525167, default 10487808):
Last sector, +/-sectors or +/-size{K,M,G,T,P} (10487808-1953525167, default 1953525167):

Created a new partition 3 of type 'Linux' and of size 926.5 GiB.

```

Après avoir accompli toutes ces étapes, pressez `p` pour afficher une table de partition qui ressemblera à :

`Command (m for help):` `p`

```
Disk /dev/sda: 931.51 GiB, 1000204886016 bytes, 1953525168 sectors
Disk model: HGST HTS721010A9
Units: sectors of 1 * 512 = 512 bytes
Sector size (logical/physical): 512 bytes / 4096 bytes
I/O size (minimum/optimal): 4096 bytes / 4096 bytes
Disklabel type: dos
Disk identifier: 0xf163b576

Device     Boot    Start        End    Sectors   Size Id Type
/dev/sda1  *        2048    2099199    2097152     1G 83 Linux
/dev/sda2        2099200   10487807    8388608     4G 82 Linux swap / Solaris
/dev/sda3       10487808 1953525167 1943037360 926.5G 83 Linux

```

Sauvegarder la table de partitions

Pressez `w` pour enregistrer la table de partitions et quitter l’utilitaire fdisk :

`Command (m for help):` `w`

```
The partition table has been altered.
Calling ioctl() to re-read partition table.
Syncing disks.

```

Maintenant il est temps de créer des systèmes de fichiers sur les partitions.

## Créer des systèmes de fichiers

**Attention !**

Il est sage de vérifier les mises à jour du micrologiciel des disques SSD ou MVNe. Certains SSD Intel (en particulier 600p et 6000p) nécessaire une mise à jour pour [éviter une corruption de données](https://bugzilla.redhat.com/show_bug.cgi?id=1402533) sur des types d’usage avec XFS. Le problème est au niveau du micrologiciel et pas XFS. L’outil smartctl peut aider à identifier le modèle et la version du micrologiciel.

## Introduction

Maintenant que les partitions ont été créées, il est temps d’y placer un système de fichiers. Dans la section suivante les différents systèmes de fichiers que Linux prend en charge seront décris. Les lecteurs qui connaissent déjà quel système de fichiers utiliser peuvent continuer avec [Appliquer un système de fichiers à une partition](/wiki/Handbook:AMD64/Installation/Disks/fr#Appliquer_un_syst.C3.A8me_de_fichiers_.C3.A0_une_partition "Handbook:AMD64/Installation/Disks/fr"). Les autres devraient continuer à lire pour en apprendre plus sur les systèmes de fichiers disponibles.

### Les systèmes de fichiers

Linux supporte des douzaines de systèmes de fichiers, bien que la plupart ne devrait être utilisés que dans des cas spécifiques. Seuls certains seront dans l’architecture amd64 – il est conseillé de se renseigner sur les systèmes de fichiers et leur prise en charge avant d’en choisir un plus expérimental pour les partitions importantes. **XFS est celui recommandé pour tous les usages et toutes les plates-formes.** Ci-dessous une liste non exhaustive :

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

#### Système de fichiers pour partition EFI

Les partitions systèmes EFI (/dev/sda1) doivent être formatés en FAT32 :

`root #` `mkfs.vfat -F 32 /dev/sda1`

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

Les installations débutées précédemment, mais non terminée peuvent reprendre ici dans le manuel. Utilisez ce lien comme lien permanent : [La reprise d’installation reprend ici](/wiki/Handbook:AMD64/Installation/Disks/fr#Resumed_installations_start_here "Handbook:AMD64/Installation/Disks/fr").

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

Pour l’installation EFI, la partition ESP devrait être monté dans la partition racine :

`root #` `mkdir --parents /mnt/gentoo/efi`

Continuez à monter des partitions additionnelles avec la commande mount.

**Remarque**

Si /tmp/ doit se trouver sur une partition séparée, pensez à changer ses droits d'accès après le montage :

`root #` `chmod 1777 /mnt/gentoo/tmp`

Cela vaut également pour /var/tmp.

Plus loin dans les instructions, le système de fichiers _proc_ (une interface virtuelle avec le noyau) ainsi que d’autres pseudos systèmes de fichiers du noyau seront montés. Mais d’abord, nous devons [copier les fichiers d’installation de Gentoo](/wiki/Handbook:AMD64/Installation/Stage/fr "Handbook:AMD64/Installation/Stage/fr").

[← Configurer le réseau](/wiki/Handbook:AMD64/Installation/Networking/fr "Previous part") [Accueil](/wiki/Handbook:AMD64/fr "Handbook:AMD64/fr") [Installer l’archive _stage3_ →](/wiki/Handbook:AMD64/Installation/Stage/fr "Next part")