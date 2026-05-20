# System

Other languages:

- [Deutsch](/wiki/Handbook:HPPA/Installation/System/de "Handbuch:HPPA/Installation/System (100% translated)")
- [English](/wiki/Handbook:HPPA/Installation/System "Handbook:HPPA/Installation/System (100% translated)")
- [español](/wiki/Handbook:HPPA/Installation/System/es "Manual de Gentoo: HPPA/Instalación/Sistema (100% translated)")
- français
- [italiano](/wiki/Handbook:HPPA/Installation/System/it "Handbook:HPPA/Installation/System/it (100% translated)")
- [magyar](/wiki/Handbook:HPPA/Installation/System/hu "Handbook:HPPA/Installation/System/hu (100% translated)")
- [polski](/wiki/Handbook:HPPA/Installation/System/pl "Handbook:HPPA/Installation/System (100% translated)")
- [português do Brasil](/wiki/Handbook:HPPA/Installation/System/pt-br "Manual:HPPA/Instalação/Sistema (100% translated)")
- [русский](/wiki/Handbook:HPPA/Installation/System/ru "Handbook:HPPA/Installation/System (100% translated)")
- [தமிழ்](/wiki/Handbook:HPPA/Installation/System/ta "கையேடு:HPPA/நிறுவல்/முறைமை (100% translated)")
- [中文（中国大陆）‎](/wiki/Handbook:HPPA/Installation/System/zh-cn "手册：HPPA/安装/配置系统 (100% translated)")
- [日本語](/wiki/Handbook:HPPA/Installation/System/ja "ハンドブック:HPPA/インストール/システム (100% translated)")
- [한국어](/wiki/Handbook:HPPA/Installation/System/ko "Handbook:HPPA/Installation/System/ko (100% translated)")

[Sommaire du manuel](/wiki/Handbook:HPPA "Handbook:HPPA")[Installation](/wiki/Handbook:HPPA/Full/Installation/fr "Handbook:HPPA/Full/Installation/fr")[À propos de l’installation](/wiki/Handbook:HPPA/Installation/About/fr "Handbook:HPPA/Installation/About/fr")[Choix du support](/wiki/Handbook:HPPA/Installation/Media/fr "Handbook:HPPA/Installation/Media/fr")[Configurer le réseau](/wiki/Handbook:HPPA/Installation/Networking/fr "Handbook:HPPA/Installation/Networking/fr")[Préparer les disques](/wiki/Handbook:HPPA/Installation/Disks/fr "Handbook:HPPA/Installation/Disks/fr")[Installer l’archive _stage3_](/wiki/Handbook:HPPA/Installation/Stage/fr "Handbook:HPPA/Installation/Stage/fr")[Installer le système de base](/wiki/Handbook:HPPA/Installation/Base/fr "Handbook:HPPA/Installation/Base/fr")[Configurer le noyau](/wiki/Handbook:HPPA/Installation/Kernel/fr "Handbook:HPPA/Installation/Kernel/fr")Configurer le système[Installer les outils](/wiki/Handbook:HPPA/Installation/Tools/fr "Handbook:HPPA/Installation/Tools/fr")[Configurer le système d’amorçage](/wiki/Handbook:HPPA/Installation/Bootloader/fr "Handbook:HPPA/Installation/Bootloader/fr")[Finaliser](/wiki/Handbook:HPPA/Installation/Finalizing/fr "Handbook:HPPA/Installation/Finalizing/fr")[Utiliser Gentoo](/wiki/Handbook:HPPA/Full/Working/fr "Handbook:HPPA/Full/Working/fr")[Introduction à Portage](/wiki/Handbook:HPPA/Working/Portage/fr "Handbook:HPPA/Working/Portage/fr")[Les options de la variable USE](/wiki/Handbook:HPPA/Working/USE/fr "Handbook:HPPA/Working/USE/fr")[Les fonctionnalités de Portage](/wiki/Handbook:HPPA/Working/Features/fr "Handbook:HPPA/Working/Features/fr")[Scripts d’initialisation systèmes](/wiki/Handbook:HPPA/Working/Initscripts/fr "Handbook:HPPA/Working/Initscripts/fr")[Variables d’environnement](/wiki/Handbook:HPPA/Working/EnvVar/fr "Handbook:HPPA/Working/EnvVar/fr")[Utiliser Portage](/wiki/Handbook:HPPA/Full/Portage/fr "Handbook:HPPA/Full/Portage/fr")[Fichiers et répertoires](/wiki/Handbook:HPPA/Portage/Files/fr "Handbook:HPPA/Portage/Files/fr")[Les variables](/wiki/Handbook:HPPA/Portage/Variables/fr "Handbook:HPPA/Portage/Variables/fr")[Mélanger plusieurs branches logicielles](/wiki/Handbook:HPPA/Portage/Branches/fr "Handbook:HPPA/Portage/Branches/fr")[Outils supplémentaires](/wiki/Handbook:HPPA/Portage/Tools/fr "Handbook:HPPA/Portage/Tools/fr")[Dépôt personnalisé](/wiki/Handbook:HPPA/Portage/CustomTree/fr "Handbook:HPPA/Portage/CustomTree/fr")[Fonctionnalités avancées](/wiki/Handbook:HPPA/Portage/Advanced/fr "Handbook:HPPA/Portage/Advanced/fr")[Configuration du réseau avec OpenRC](/wiki/Handbook:HPPA/Full/Networking/fr "Handbook:HPPA/Full/Networking/fr")[Bien démarrer](/wiki/Handbook:HPPA/Networking/Introduction/fr "Handbook:HPPA/Networking/Introduction/fr")[Configuration avancée](/wiki/Handbook:HPPA/Networking/Advanced/fr "Handbook:HPPA/Networking/Advanced/fr")[Les modules réseau](/wiki/Handbook:HPPA/Networking/Modular/fr "Handbook:HPPA/Networking/Modular/fr")[Sans fil](/wiki/Handbook:HPPA/Networking/Wireless/fr "Handbook:HPPA/Networking/Wireless/fr")[Ajouter des fonctionnalités](/wiki/Handbook:HPPA/Networking/Extending/fr "Handbook:HPPA/Networking/Extending/fr")[Gestion dynamique](/wiki/Handbook:HPPA/Networking/Dynamic/fr "Handbook:HPPA/Networking/Dynamic/fr")

## Contents

- [1Informations sur le système de fichiers](#Informations_sur_le_syst.C3.A8me_de_fichiers)
  - [1.1Étiquettes de systèmes de fichiers et UUID](#.C3.89tiquettes_de_syst.C3.A8mes_de_fichiers_et_UUID)
  - [1.2Étiquettes de partitions et UUID](#.C3.89tiquettes_de_partitions_et_UUID)
  - [1.3À propos de _fstab_](#.C3.80_propos_de_fstab)
  - [1.4Créer le fichier _fstab_](#Cr.C3.A9er_le_fichier_fstab)
    - [1.4.1Système DOS/BIOS déprécié](#Syst.C3.A8me_DOS.2FBIOS_d.C3.A9pr.C3.A9ci.C3.A9)
    - [1.4.2DPS UEFI PARTUUID](#DPS_UEFI_PARTUUID)
- [2Informations réseau](#Informations_r.C3.A9seau)
  - [2.1Nom de l’hôte](#Nom_de_l.E2.80.99h.C3.B4te)
    - [2.1.1Paramétrer le nom de l’hôte ( _OpenRC_ ou _systemd_)](#Param.C3.A9trer_le_nom_de_l.E2.80.99h.C3.B4te_.28OpenRC_ou_systemd.29)
    - [2.1.2_systemd_](#systemd)
  - [2.2Réseau](#R.C3.A9seau)
    - [2.2.1DHCP via _dhcpcd_ (tout système d’initialisation)](#DHCP_via_dhcpcd_.28tout_syst.C3.A8me_d.E2.80.99initialisation.29)
    - [2.2.2_netifrc_ ( _OpenRC_)](#netifrc_.28OpenRC.29)
      - [2.2.2.1Configurer le réseau](#Configurer_le_r.C3.A9seau)
      - [2.2.2.2Démarrer automatiquement la mise en réseau au démarrage](#D.C3.A9marrer_automatiquement_la_mise_en_r.C3.A9seau_au_d.C3.A9marrage)
  - [2.3Le fichier des hôtes](#Le_fichier_des_h.C3.B4tes)
- [3Informations système](#Informations_syst.C3.A8me)
  - [3.1Mot de passe _root_](#Mot_de_passe_root)
  - [3.2Configuration de l’initialisation et du démarrage](#Configuration_de_l.E2.80.99initialisation_et_du_d.C3.A9marrage)
    - [3.2.1_OpenRC_](#OpenRC)
    - [3.2.2_systemd_](#systemd_2)

## Informations sur le système de fichiers

#### Étiquettes de systèmes de fichiers et UUID

MBR (BIOS) et GPT incluent tous les deux le support pour les étiquettes de système de fichiers et les UUID ( _Universally Unique Identifier_) de système de fichiers. Ces attributs peuvent être définis dans /etc/fstab comme alternatives lors de l’utilisation de la commande mount pour détecter et monter les blocs de périphériques. Les étiquettes de système de fichiers et les UUIDs de système de fichiers sont identifiés par le préfixe LABEL et UUID et peuvent être visualisés grâce à la commande blkid :

`root #` `blkid`

**Attention !**

Si le système de fichiers à l’intérieur de la partition est supprimé, alors les valeurs d’étiquettes et d’UUID seront également modifiées ou supprimées.

Pour l’unicité, les lecteurs, utilisant des tables de partitions de type MBR, sont recommandés d’utiliser les UUID à la place des étiquettes pour définir les volumes montables dans /etc/fstab.

**Important**

Les UUID des systèmes de fichiers sur un volume LVM et ses _snapshots_ sont identiques aussi, l’utilisation d’UUID pour monter des volumes LVM devraient être évitées.

#### Étiquettes de partitions et UUID

Les systèmes avec des étiquettes disques GPT ont quelques options plus « robustes » disponibles à définir dans /etc/fstab. Les étiquettes de partition et les UUID de partition peuvent être utilisés pour identifier les partitions individuelles du bloc de périphérique, quel que soit le système de fichiers choisi pour la partition elle-même. Les étiquettes de partition et les UUIDs sont identifiés par les préfixes PARTLABEL et/ou PARTUUID respectivement et peuvent être consultés dans le terminal en exécutant la commande blkid :

La sortie d’un système EFI **amd64** utilisant _Discoverable Partition Specification UUID_ (découverte automatique de partitions) ressemblera à :

`root #` `blkid`

```
/dev/sr0: BLOCK_SIZE="2048" UUID="2023-08-28-03-54-40-00" LABEL="ISOIMAGE" TYPE="iso9660" PTTYPE="PMBR"
/dev/loop0: TYPE="squashfs"
/dev/sda2: PARTUUID="0657fd6d-a4ab-43c4-84e5-0933c84b4f4f"
/dev/sda3: PARTUUID="1cdf763a-5b4c-4dbf-99db-a056c504e8b2"
/dev/sda1: PARTUUID="c12a7328-f81f-11d2-ba4b-00a0c93ec93b"

```

Bien que pas toujours vrai pour les étiquettes de partition, utiliser un UUID pour identifier une partition dans fstab garantit que le chargeur de démarrage ne sera pas confus en cherchant un volume spécifique, même si le système de fichiers est changé ou réécrit à l’avenir. L’utilisation des anciens fichiers de bloc de périphériques par défaut (/dev/sd\*N) pour définir les partitions dans fstab est risquée pour les systèmes qui sont redémarrés régulièrement et qui possèdent des blocs de périphériques SATA régulièrement ajoutés ou supprimés.

La dénomination des fichiers de périphériques de bloc dépend d’un certain nombre de facteurs, y compris comment et dans quel ordre les disques sont attachés au système. Ils peuvent également apparaître dans un ordre différent en fonction duquel les périphériques sont détectés par le noyau au cours du processus de démarrage. Cela étant dit, à moins que l’on ait l’intention de constamment jouer avec la commande de disque, l’utilisation des fichiers de périphériques par défaut est une approche simple et directe.

### À propos de _fstab_

Sous Linux, toutes les partitions utilisées par le système doivent être listées dans [/etc/fstab](/wiki//etc/fstab/fr "/etc/fstab/fr"). Ce fichier contient les points de montage des partitions (où elles sont vues dans la structure du système de fichier), comment elles doivent être montées et avec quels paramètres (automatiquement ou non, si les utilisateurs peuvent les monter ou non, etc.)

### Créer le fichier _fstab_

**Remarque**

Si _systemd_ est utilisé comme système d’initialisation, l’UUID de partition conforme à _Discoverable Partition Specification_ est renseigné dans [préparer les disques](/wiki/Handbook:HPPA/Installation/Disks/fr "Handbook:HPPA/Installation/Disks/fr") ; et pour les systèmes UEFI, la création dans _fstab_ peut être passé puisque l’auto-montage 'systemd _respecte les spécifications._

Le fichier /etc/fstab utilise un format ressemblant à celui d’un tableau. Chaque ligne comporte six champs, séparés par des espaces blancs (espace, tabulation ou les deux). Chaque champ à sa propre signification :

1. Le premier champ indique le périphérique ou système de fichier distant à monter. Plusieurs types d’identificateurs sont disponibles pour les périphériques : chemin vers les fichiers du périphérique, étiquettes des systèmes de fichiers et UUID, étiquettes de partitions et UUID ( _Universally unique identifier_).
2. Le second champ indique le point de montage sur lequel la partition sera montée.
3. Le troisième champ indique le type de système de fichiers utilisé par la partition.

Le quatrième champ indique les options utilisées par mount lors du montage de la partition. Comme chaque système de fichiers à ses propres options de montage, les administrateurs système sont encouragés à lire la page de manuel de mount (man mount) pour une liste complète. Des options multiples sont séparées par une virgule.

1. Le cinquième champ est utilisé par dump pour déterminer si la partition doit être sauvegardée ou non. Ce champ peut généralement être laissé à `0`.
2. Le sixième champ est utilisé par fsck pour déterminer dans quel ordre les systèmes de fichiers doivent être vérifiés si le système n’a pas été terminé correctement. Le système de fichier _root_ devrait être à `1` et les autres à `2` (ou `0` si une vérification n’est pas nécessaire).

**Important**

Le fichier /etc/fstab par défaut fourni par Gentoo n’est pas un fichier /etc/fstab valide, mais plutôt une sorte de modèle.

`root #` `nano /etc/fstab`

#### Système DOS/BIOS déprécié

Regardons comment noter les options pour la partition /boot/. Ceci n’est qu'un exemple et doit être modifié en fonctions des décisions prises plus tôt dans l’installation. Dans notre exemple de partitionnement **hppa**, /boot/ est généralement la /dev/sda2 partition, avec xfs comme système de fichier. Cette partition nécessite d’être vérifiée lors du démarrage et nous écrirons donc :

FILE **`/etc/fstab`** **Une exemple de ligne BIOS/DOS déprécié pour /etc/fstab**

```
# Ajuster le formatage et les partitions additionnelles crées pendant l’étape « Préparer les disques »
/dev/sda2   /boot     ext2     defaults        0 2

```

Certains administrateurs ne veulent pas que leur partition /boot/ soit montée automatiquement afin d’augmenter la sécurité de leur système. Ces personnes doivent remplacer `defaults` par `noauto`. Cela signifie que ces utilisateurs devront monter manuellement cette partition à chaque fois qu’ils voudront l’utiliser.

Ajouter les règles qui correspondent au schéma de partitionnement décidé précédemment et ajouter des règles pour les périphériques tels que les lecteurs de CD-ROM, et bien sûr, si d’autres partitions ou lecteurs sont utilisés, pour ceux-là également.

Ci-dessous est un exemple plus élaboré de fichier /etc/fstab :

FILE **`/etc/fstab`** **Un exemple complet de /etc/fstab pour DOS/BIOS déprécié**

```
# Ajustez le formattage et ajoutez les partitions supplémentaires créées dans « Préparer les disques »
/dev/sda2   /boot        ext2    defaults    0 2
/dev/sda3   none         swap    sw                   0 0
/dev/sda4   /            xfs    defaults,noatime              0 1

/dev/cdrom  /mnt/cdrom   auto    noauto,user          0 0

```

#### DPS UEFI PARTUUID

Below is an example of an /etc/fstab file for a disk formatted with a GPT disklabel and Discoverable Partition Specification (DPS) UUIDs set for UEFI firmware:

FILE **`/etc/fstab`** **GPT disklabel DPS PARTUUID fstab example**

```
# Adjust any formatting difference and additional partitions created from the "Preparing the disks" step.
# This example shows a GPT disklabel with Discoverable Partition Specification (DSP) UUID set:
PARTUUID=c12a7328-f81f-11d2-ba4b-00a0c93ec93b   /efi        vfat                       0 2
PARTUUID=0657fd6d-a4ab-43c4-84e5-0933c84b4f4f   none            sw                           0 0
PARTUUID=1aacdb3b-5444-4138-bd9e-e5c2239b2346   /           xfs    defaults,noatime              0 1

```

Quand `auto` est utilisé en tant que troisième champ, cela fait deviner à la commande mount le système de fichiers utilisé. Cela est recommandé pour les supports amovibles car ils peuvent être créés avec des systèmes de fichiers différents. L’option `user` dans le quatrième champ rend possible pour les utilisateurs non _root_ de monter le CD.

Pour améliorer les performances, la plupart des utilisateurs ajouteront l’option `noatime` sur le point de montage ; cela accélère les lectures disques en évitant d’enregistrer les heures d’accès (généralement inutile). Cela est également conseillé pour les systèmes avec des _SSD_ ( _Solid State Drive_). Certains utilisateurs pourrait considérer l’usage de `lazytime` en remplacement.

**Conseil**

En raison de la dégradation des performances, utiliser l’option de montage `discard` dans /etc/fstab n’est pas recommandé. Il est généralement mieux de planifier l’effacement des blocs périodiquement avec cron ou un minuteur ( _systemd_). Lire [Traitement _fstrim_ périodique](/wiki/SSD#Periodic_fstrim_jobs "SSD") pour plus d’informations.

Bien vérifier le fichier /etc/fstab, puis sauvegarder avant de quitter pour continuer.

## Informations réseau

Il est important de noter que les sections suivantes sont proposées afin d’aider le lecteur à paramétrer rapidement son système pour une utilisation réseau local.

Pour les systèmes _OpenRC_, un guide plus détaillé pour le paramétrage du réseau est disponible dans la [configuration avancé du réseau](/wiki/Handbook:HPPA/Networking/Introduction/fr "Handbook:HPPA/Networking/Introduction/fr"), qui est situé à la fin du manuel. Les systèmes qui ont des besoins réseaux plus spécifiques devraient passer ce qui suit, et revenir après avoir configuré le réseau.

Pour un usage spécifique _systemd_, consultez la [partie réseau](/wiki/Systemd/fr#Network "Systemd/fr") de l’article [systemd](/wiki/Systemd/fr "Systemd/fr").

### Nom de l’hôte

L’un des choix qui incombe à l’administrateur système est de nommer son ordinateur. Cela peut sembler assez facile, mais la plupart des utilisateurs ont des difficultés à trouver un nom approprié pour le nom d’hôte. Pour aider, il est bon de savoir que cette décision n’est pas finale – cela peut être changé par la suite. Dans les exemples ci-dessous, le nom d’hôte « tux » est utilisé.

#### Paramétrer le nom de l’hôte ( _OpenRC_ ou _systemd_)

`root #` `echo tux > /etc/hostname`

#### _systemd_

Pour configurer le nom de l’hôte sur un système utilisant _systemd_, l’utilitaire hostnamectl peut être employé. Cependant, pendant la procédure d’installation, [1er démarrage de _systemd_](/wiki/Handbook:HPPA/Installation/System/fr#Init_and_boot_configuration_systemd "Handbook:HPPA/Installation/System/fr") doit être utilisé (voir plus loin dans le manuel).

Pour définir le nom d’hôte à « tux », lancer :

`root #` `hostnamectl hostname tux`

Pour afficher l’aide, exécuter hostnamectl --help ou man 1 hostnamectl.

### Réseau

Il y a _beaucoup_ de possibilités pour configurer une interface réseau. Cette section couvre seulement quelques méthodes. Choisissez celle qui semble le plus adapté à votre besoin.

#### DHCP via _dhcpcd_ (tout système d’initialisation)

La plupart des réseaux ont un serveur _DHCP_ ( _Dynamic Host Configuration Protocol_). Si cela est le cas, alors utilisez le programme `dhcpd` est recommandé pour obtenir une adresse _IP_ ( _Internet Protocol_).

Pour installer :

`root #` `emerge --ask net-misc/dhcpcd`

Pour activer puis lancer le service sur un système _OpenRC_ :

`root #` `rc-update add dhcpcd default
`

`root #` `rc-service dhcpcd start
`

Pour activer puis lancer le service sur un système _systemd_ :

`root #` `systemctl enable --now dhcpcd`

Une fois ces étapes terminées, lors du prochain démarrage système, _dhcpd_ devrait obtenir une adresse IP du serveur DHCP. Voir [Dhcpcd](/wiki/Dhcpcd "Dhcpcd") pour plus d’informations.

#### _netifrc_ ( _OpenRC_)

**Conseil**

C’est une façon particulière de configurer un réseau avec [Netifrc](/wiki/Netifrc "Netifrc") sur _OpenRC_. des méthodes plus simples existent comment [Dhcpcd](/wiki/Dhcpcd "Dhcpcd").

##### Configurer le réseau

Lors de l’installation de Gentoo Linux, la mise en réseau a déjà été configurée. Cependant, c’était pour l’environnement _live_ et non pour l’environnement installé. Ici, la configuration du réseau est réalisée pour le système Gentoo Linux installé.

**Remarque**

Des informations plus détaillées sur la mise en réseau, incluant des sujets avancés tels que _bonding_, _bridging_, _VLAN_ 802.1Q ou les réseaux sans fil se trouvent dans la section [configuration réseau avancée](/wiki/Handbook:HPPA/Networking/Introduction/fr "Handbook:HPPA/Networking/Introduction/fr").

Toutes les informations concernant la mise en réseau sont regroupées dans le fichier /etc/conf.d/net. Ce fichier utilise une syntaxe directe mais peu intuitive. Pas de panique, tout est expliqué plus bas. Un exemple complet et commenté couvrant plusieurs configurations possibles se trouve dans /usr/share/doc/netifrc-\*/net.example.bz2.

D’abord, installer [net-misc/netifrc](https://packages.gentoo.org/packages/net-misc/netifrc) :

`root #` `emerge --ask --noreplace net-misc/netifrc`

_DHCP_ est utilisé par défaut. Pour que _DHCP_ fonctionne, un client DHCP doit être installé. Cela est expliqué plus tard lors de l’installation des outils systèmes nécessaires.

Si la connexion au réseau doit être configurée à cause d’options _DHCP_ spécifiques or car _DHCP_ n’est pas du tout utilisé, alors ouvrir le fichier /etc/conf.d/net :

`root #` `nano /etc/conf.d/net`

Configurer les deux variables config\_eth0 et routes\_eth0 avec les informations d’adresse _IP_ et de routage appropriées :

**Remarque**

Il est supposé que l’interface réseau s’appellera eth0. Cela est, cependant, entièrement dépendant du système. Il est recommandé de partir du principe que l’interface portera la même nom que lors du démarrage depuis le support d’installation si ce dernier est suffisamment récent. Plus d’informations sont accessibles dans la section [Nommage de l’interface réseau](/wiki/Handbook:HPPA/Networking/Advanced#Network_interface_naming.2Ffr "Handbook:HPPA/Networking/Advanced").

FILE **`/etc/conf.d/net`** **Définition d'un adresse IP statique**

```
config_eth0="192.168.0.2 netmask 255.255.255.0 brd 192.168.0.255"
routes_eth0="default via 192.168.0.1"

```

Pour utiliser DHCP, définir la variable config\_eth0>/var>  :

FILE **`/etc/conf.d/net`** **Paramétrage DHCP**

```
config_eth0="dhcp"

```

Lire /usr/share/doc/netifrc-\*/net.example.bz2 pour une liste d’options de configuration supplémentaires. Lire également la page de manuel de DHCP si des options DHCP spécifiques doivent être utilisées.

Si le système possède plusieurs interfaces réseau, répéter les étapes précédentes pour config\_eth1, config\_eth2, etc.

Sauvegarder la configuration et quitter avant de continuer.

##### Démarrer automatiquement la mise en réseau au démarrage

Pour activer les interfaces réseau lors du démarrage, elles doivent être ajoutées au _runlevel_ par défaut.

`root #` `cd /etc/init.d
`

`root #` `ln -s net.lo net.eth0
`

`root #` `rc-update add net.eth0 default`

Si le système possède plusieurs interfaces réseau, les fichiers appropriés net.\* doivent être crées de la même manière que pour net.eth0.

Si, après le démarrage du système, il est découvert que le nom de l’interface réseau (qui est actuellement documentée comme `eth0`) était erronée, exécuter les étapes suivantes afin de rectifier le problème :

1. Mettre à jour le fichier /etc/conf.d/net avec le nom d’interface correct (comme `enp3s0` ou `enp5s0`, au lieu de `eth0`).
2. Créer un nouveau lien symbolique (comme /etc/init.d/net.enp3s0).
3. Supprimer l’ancien lien symbolique (rm /etc/init.d/net.eth0).
4. Ajouter le nouveau au _runlevel_ par défaut.
5. Supprimer l’ancien en utilisant la commande rc-update del net.eth0 default.

### Le fichier des hôtes

Une étape suivante importante est d’informer le nouveau système des autres hôtes de l’environnement réseau. Les autres hôtes du réseau peuvent être définis dans le fichier /etc/hosts. Ajouter un nom permet une résolution d’adresse IP pour un hôte inconnu du serveur de nom ( _DNS_).

`root #` `nano /etc/hosts`

FILE **`/etc/hosts`** **Remplir les informations réseau**

```
# Cela définie le système actuelle et doit être mis
127.0.0.1     tux.homenetwork tux localhost
::1           tux.homenetwork tux localhost

# Définitions optionnelles d’autres systèmes sur le réseau
192.168.0.5   jenny.homenetwork jenny
192.168.0.6   benny.homenetwork benny

```

Sauvegarder et quitter l’éditeur pour continuer.

## Informations système

### Mot de passe _root_

Configurer le mot de passe _root_ en utilisant la commande passwd.

`root #` `passwd`

Plus tard, un compte utilisateur sera créé pour les utilisations quotidiennes.

### Configuration de l’initialisation et du démarrage

#### _OpenRC_

Quand OpenRC est utilisé avec Gentoo, il utilise le fichier /etc/rc.conf pour configurer les services, le démarrage et l’arrêt d’un système. Ouvrir /etc/rc.conf et s’émerveiller devant tous les commentaires du fichier. Vérifier les configurations et les modifier si nécessaire.

`root #` `nano /etc/rc.conf`

Ensuite, ouvrir le fichier /etc/conf.d/keymaps afin de gérer la configuration du clavier. Le modifier pour configurer et sélectionner le bon clavier.

`root #` `nano /etc/conf.d/keymaps`

Prendre bien soin lors de la configuration de la variable keymap. Si le mauvais schéma de clavier est sélectionné, il se passera des choses inattendues lors de l’utilisation du clavier.

Finalement, modifier le fichier /etc/conf.d/hwclock afin de configurer les options d’horloge.

`root #` `nano /etc/conf.d/hwclock`

Si l’horloge matérielle n’utilise pas le temps universel coordonné ( _UTC_), il est nécessaire de configurer `clock="local"` dans le fichier. Sinon le système peut montrer des horloge faussées.

#### _systemd_

Tout d’abord, il est recommandé d’exécuter systemd-machine-id-setup et systemd-firstboot lesquels vont préparer plusieurs composants système pour un premier démarrage dans un environnement _systemd_. Passer les options suivantes permet d’inclure une invite de commandes pour paramétrer la locale, le fuseau horaire, le nom de l’hôte, le mot de passe _root_ et des valeurs _shell_. Un numéro aléatoire sera également assigné à la machine :

`root #` `systemd-machine-id-setup`

`root #` `systemd-firstboot --prompt`

Les utilisateurs devraient lancer systemctl pour réinitialiser tous les fichiers _unit_ avec les nouvelles valeurs paramétrées :

**Remarque**

Il est possible qu’une erreur ou un avertissement s’affiche après le lancement de la commande suivante. C’est normal, le _live_ Gentoo utiliser _OpenRC_.

`root #` `systemctl preset-all --preset-mode=enable-only`

Il est possible de lancer un changement complet de la configuration, mais cela peut également réinitialiser tous les services déjà paramétré pendant la procédure :

`root #` `systemctl preset-all`

Ces deux étapes assureront une transition douce de l’environnement _live_ au premier démarrage.

[← Configurer le noyau](/wiki/Handbook:HPPA/Installation/Kernel/fr "Previous part") [Accueil](/wiki/Handbook:HPPA/fr "Handbook:HPPA/fr") [Installing tools →](/wiki/Handbook:HPPA/Installation/Tools/fr "Next part")