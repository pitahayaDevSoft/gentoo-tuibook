# Tools

Other languages:

- [Deutsch](/wiki/Handbook:PPC64/Installation/Tools/de "Handbuch:PPC64/Installation/Tools (100% translated)")
- [English](/wiki/Handbook:PPC64/Installation/Tools "Handbook:PPC64/Installation/Tools (100% translated)")
- [español](/wiki/Handbook:PPC64/Installation/Tools/es "Manual de Gentoo: PPC64/Instalación/Herramientas (100% translated)")
- français
- [italiano](/wiki/Handbook:PPC64/Installation/Tools/it "Handbook:PPC64/Installation/Tools/it (100% translated)")
- [magyar](/wiki/Handbook:PPC64/Installation/Tools/hu "Handbook:PPC64/Installation/Tools/hu (100% translated)")
- [polski](/wiki/Handbook:PPC64/Installation/Tools/pl "Handbook:PPC64/Installation/Tools (100% translated)")
- [português do Brasil](/wiki/Handbook:PPC64/Installation/Tools/pt-br "Handbook:PPC64/Installation/Tools/pt-br (50% translated)")
- [русский](/wiki/Handbook:PPC64/Installation/Tools/ru "Handbook:PPC64/Installation/Tools (100% translated)")
- [தமிழ்](/wiki/Handbook:PPC64/Installation/Tools/ta "கையேடு:PPC64/நிறுவல்/கருவிகள் (100% translated)")
- [中文（中国大陆）‎](/wiki/Handbook:PPC64/Installation/Tools/zh-cn "手册：PPC64/安装/安装系统工具 (100% translated)")
- [日本語](/wiki/Handbook:PPC64/Installation/Tools/ja "ハンドブック:PPC64/インストール/ツール (100% translated)")
- [한국어](/wiki/Handbook:PPC64/Installation/Tools/ko "Handbook:PPC64/Installation/Tools/ko (100% translated)")

[Sommaire du manuel](/wiki/Handbook:PPC64 "Handbook:PPC64")[Installation](/wiki/Handbook:PPC64/Full/Installation/fr "Handbook:PPC64/Full/Installation/fr")[À propos de l’installation](/wiki/Handbook:PPC64/Installation/About/fr "Handbook:PPC64/Installation/About/fr")[Choix du support](/wiki/Handbook:PPC64/Installation/Media/fr "Handbook:PPC64/Installation/Media/fr")[Configurer le réseau](/wiki/Handbook:PPC64/Installation/Networking/fr "Handbook:PPC64/Installation/Networking/fr")[Préparer les disques](/wiki/Handbook:PPC64/Installation/Disks/fr "Handbook:PPC64/Installation/Disks/fr")[Installer l’archive _stage3_](/wiki/Handbook:PPC64/Installation/Stage/fr "Handbook:PPC64/Installation/Stage/fr")[Installer le système de base](/wiki/Handbook:PPC64/Installation/Base/fr "Handbook:PPC64/Installation/Base/fr")[Configurer le noyau](/wiki/Handbook:PPC64/Installation/Kernel/fr "Handbook:PPC64/Installation/Kernel/fr")[Configurer le système](/wiki/Handbook:PPC64/Installation/System/fr "Handbook:PPC64/Installation/System/fr")Installer les outils[Configurer le système d’amorçage](/wiki/Handbook:PPC64/Installation/Bootloader/fr "Handbook:PPC64/Installation/Bootloader/fr")[Finaliser](/wiki/Handbook:PPC64/Installation/Finalizing/fr "Handbook:PPC64/Installation/Finalizing/fr")[Utiliser Gentoo](/wiki/Handbook:PPC64/Full/Working/fr "Handbook:PPC64/Full/Working/fr")[Introduction à Portage](/wiki/Handbook:PPC64/Working/Portage/fr "Handbook:PPC64/Working/Portage/fr")[Les options de la variable USE](/wiki/Handbook:PPC64/Working/USE/fr "Handbook:PPC64/Working/USE/fr")[Les fonctionnalités de Portage](/wiki/Handbook:PPC64/Working/Features/fr "Handbook:PPC64/Working/Features/fr")[Scripts d’initialisation systèmes](/wiki/Handbook:PPC64/Working/Initscripts/fr "Handbook:PPC64/Working/Initscripts/fr")[Variables d’environnement](/wiki/Handbook:PPC64/Working/EnvVar/fr "Handbook:PPC64/Working/EnvVar/fr")[Utiliser Portage](/wiki/Handbook:PPC64/Full/Portage/fr "Handbook:PPC64/Full/Portage/fr")[Fichiers et répertoires](/wiki/Handbook:PPC64/Portage/Files/fr "Handbook:PPC64/Portage/Files/fr")[Les variables](/wiki/Handbook:PPC64/Portage/Variables/fr "Handbook:PPC64/Portage/Variables/fr")[Mélanger plusieurs branches logicielles](/wiki/Handbook:PPC64/Portage/Branches/fr "Handbook:PPC64/Portage/Branches/fr")[Outils supplémentaires](/wiki/Handbook:PPC64/Portage/Tools/fr "Handbook:PPC64/Portage/Tools/fr")[Dépôt personnalisé](/wiki/Handbook:PPC64/Portage/CustomTree/fr "Handbook:PPC64/Portage/CustomTree/fr")[Fonctionnalités avancées](/wiki/Handbook:PPC64/Portage/Advanced/fr "Handbook:PPC64/Portage/Advanced/fr")[Configuration du réseau avec OpenRC](/wiki/Handbook:PPC64/Full/Networking/fr "Handbook:PPC64/Full/Networking/fr")[Bien démarrer](/wiki/Handbook:PPC64/Networking/Introduction/fr "Handbook:PPC64/Networking/Introduction/fr")[Configuration avancée](/wiki/Handbook:PPC64/Networking/Advanced/fr "Handbook:PPC64/Networking/Advanced/fr")[Les modules réseau](/wiki/Handbook:PPC64/Networking/Modular/fr "Handbook:PPC64/Networking/Modular/fr")[Sans fil](/wiki/Handbook:PPC64/Networking/Wireless/fr "Handbook:PPC64/Networking/Wireless/fr")[Ajouter des fonctionnalités](/wiki/Handbook:PPC64/Networking/Extending/fr "Handbook:PPC64/Networking/Extending/fr")[Gestion dynamique](/wiki/Handbook:PPC64/Networking/Dynamic/fr "Handbook:PPC64/Networking/Dynamic/fr")

## Contents

- [1Outil de journalisation du système](#Outil_de_journalisation_du_syst.C3.A8me)
  - [1.1_OpenRC_](#OpenRC)
  - [1.2systemd](#systemd)
- [2Facultatif : démon _cron_](#Facultatif_:_d.C3.A9mon_cron)
  - [2.1OpenRC](#OpenRC_2)
    - [2.1.1Défaut : cronie](#D.C3.A9faut_:_cronie)
    - [2.1.2Alternative : dcron](#Alternative_:_dcron)
    - [2.1.3Alternative : fcron](#Alternative_:_fcron)
    - [2.1.4Alternative : bcron](#Alternative_:_bcron)
    - [2.1.5modprobed-db users](#modprobed-db_users)
  - [2.2systemd](#systemd_2)
    - [2.2.1modprobed-db users](#modprobed-db_users_2)
- [3Facultatif : Accès distant](#Facultatif_:_Acc.C3.A8s_distant)
  - [3.1OpenRC](#OpenRC_3)
  - [3.2systemd](#systemd_3)
- [4Facultatif : Autocomplétion _shell_](#Facultatif_:_Autocompl.C3.A9tion_shell)
  - [4.1Bash](#Bash)
- [5Suggéré : Synchronisation du temps](#Sugg.C3.A9r.C3.A9_:_Synchronisation_du_temps)
  - [5.1chrony](#chrony)
  - [5.2OpenRC](#OpenRC_4)
  - [5.3systemd](#systemd_4)
  - [5.4systemd-timesyncd](#systemd-timesyncd)
- [6Outils de système de fichiers](#Outils_de_syst.C3.A8me_de_fichiers)

## Outil de journalisation du système

### _OpenRC_

Quelques outils sont absents de l’archive _stage3_ car plusieurs paquets fournissent la même fonctionnalité. Le choix est laissé à l’utilisateur de savoir quels paquets utiliser.

Le premier outil sur lequel un choix doit se faire est un outil de journalisation pour le système. Unix et Linux ont un historique excellent de capacités de journalisations – si besoin – tout ce qui se passe sur le système peut être enregistré dans des fichiers journaux.

Gentoo offre plusieurs utilitaires de journalisation. Cela inclut notamment :

- [sysklogd](/wiki/Sysklogd "Sysklogd") ([app-admin/sysklogd](https://packages.gentoo.org/packages/app-admin/sysklogd)) \- Offre l’ensemble traditionnel des démons de journalisation système. La configuration par défaut fonctionne correctement ce qui fait de ce paquet une bonne option pour les débutants.
- [syslog-ng](/wiki/Syslog-ng "Syslog-ng") ([app-admin/syslog-ng](https://packages.gentoo.org/packages/app-admin/syslog-ng)) \- Un système de journalisation avancé. Nécessite une configuration supplémentaire pour fonctionner au delà de la journalisation dans un seul gros fichier. Les utilisateurs avancés peuvent choisir ce système de journalisation du fait de son potentiel mais attention, un configuration est nécessaire pour une journalisation intelligente.
- [metalog](/wiki/Metalog "Metalog") ([app-admin/metalog](https://packages.gentoo.org/packages/app-admin/metalog)) \- Un système de journalisation hautement configurable.

Il pourrait y avoir d’autres systèmes de journalisation disponibles depuis les dépôts Gentoo, comme le nombre de paquets disponibles augmente tous les jours.

**Conseil**

Si syslog-ng est utilisé, il est recommandé d’installer et de configurer [logrotate](/wiki/Logrotate "Logrotate"). syslog-ng ne fournit aucun mécanisme de rotation pour les fichiers du journal. Les nouvelles versions (>= 2.0) de sysklogd intègrent par contre leur propre mécanisme de rotation.

Pour installer l’outil de journalisation désiré, installez-le. Sur OpenRC, ajoutez-le au niveau d’exécution par défaut en utilisant rc-update. L’exemple suivant installe et active [app-admin/sysklogd](https://packages.gentoo.org/packages/app-admin/sysklogd) comme système de journalisation :

`root #` `emerge --ask app-admin/sysklogd`

`root #` `rc-update add sysklogd default`

### systemd

Les systèmes de journalisation sont présentés pour les systèmes OpenRC, systemd inclut un service intégré appelé « _systemd-journald_ ». _systemd-journald_ est supporte la majorité des fonctionnalités mise à avant précédemment. Ceci étant écrit, la plupart des installations avec systemd peuvent ignorer cette section pour ajouter un utilitaire supplémentaire.

Consultez man journalctl pour plus de détails sur l’utilisation de journalctl pour interroger et lire les journaux systèmes.

Pour de multiples raisons, comme la centralisation des journaux sur un hôte dédié, il peut être important d’inclure un système redondant de journalisation sur un système basé sur systemd. Ceci est un usage peu courant pour le lecteur lambda du manuel, et est considéré comme un usage avancé. Il ne sera pas couvert dans le manuel.

## Facultatif : démon _cron_

### OpenRC

Ensuite viens le démon _cron_. Bien que cela soit facultatif et pas nécessaire pour tous les systèmes, il est judicieux d’en installer un.

Un démon _cron_ exécute des commandes dans des intervalles planifiés. Ils peuvent être journalier, à la semaine ou mensuel, tous les mardis, la semaine de 4 jeudis, etc. Un administrateur système sage utilisera le démon _cron_ pour automatiser des routines systèmes de maintenance.

Tous les démons _cron_ supportent un haut niveau de granularité pour planifier les tâches, et généralement incluent la possibilité d’envoyer un courriel ou une autre forme de notification si une tâche planifiée ne se termine pas comme prévu.

Gentoo offre plusieurs démons _cron_ possibles, incluant :

- [cronie](/wiki/Cron/fr#cronie "Cron/fr") ([sys-process/cronie](https://packages.gentoo.org/packages/sys-process/cronie)) \- cronie est basé sur le _cron_ originel et a des améliorations de sécurité et de configuration ; comme la possibilité d’utiliser PEM et SELinux.
- [dcron](/wiki/Cron/fr#dcron "Cron/fr") ([sys-process/dcron](https://packages.gentoo.org/packages/sys-process/dcron)) \- Ce démon _cron_ léger essaye d’être simple et sécurisé, avec juste assez de fonctionnalités utiles.
- [fcron](/wiki/Cron/fr#fcron "Cron/fr") ([sys-process/fcron](https://packages.gentoo.org/packages/sys-process/fcron)) \- Un planificateur de commandes avec des capacités étendues par rapport à _cron_ et _anacron_.
- [bcron](/wiki/Cron/fr#bcron "Cron/fr") ([sys-process/bcron](https://packages.gentoo.org/packages/sys-process/bcron)) \- Un jeune système de _cron_ conçu avec la sécurité en tête. Pour atteindre ce but, il est découpé en plusieurs modules, chacun responsable d’une tâche avec une communication strictement vérifiées entre chaque.

#### Défaut : cronie

L’exemple suivant utilise [sys-process/cronie](https://packages.gentoo.org/packages/sys-process/cronie) :

`root #` `emerge --ask sys-process/cronie`

Ajouter cronie au démarrage système par défaut, pour le démarrer automatiquement lors de l’allumage :

`root #` `rc-update add cronie default`

#### Alternative : dcron

`root #` `emerge --ask sys-process/dcron`

Si dcron est le démon _cron_ souhaité, une commande d’initialisation supplémentaire doit être exécutée :

`root #` `crontab /etc/crontab`

#### Alternative : fcron

`root #` `emerge --ask sys-process/fcron`

Si fcron est sélectionné comme planification de tâches, une étape supplémentaire est nécessaire :

`root #` `emerge --config sys-process/fcron`

#### Alternative : bcron

bcron est un _cron_ récent conçu pour une séparation des privilèges.

`root #` `emerge --ask sys-process/bcron`

#### modprobed-db users

If using [sys-kernel/modprobed-db](https://packages.gentoo.org/packages/sys-kernel/modprobed-db) was picked as an option for manually compiling the kernel.

Now is the time to setup a crontab to periodically scan the system for hardware used.

FILE **`/etc/crontab`** **Run modprobed-db once every 6 hours**

```
0 */6 * * *     /usr/bin/modprobed-db store &> /dev/null

```

At a later date of at least a week, please visit the kernel build section of the [modprobed-db](/wiki/Modprobed-db#Building_kernels.2Ffr "Modprobed-db") article to complete the setup.

### systemd

Similairement au système de journalisation, systemd inclus le support des tâches planifiées sous la forme de _timers_. Les _timers_ systemd peuvent être lancés au niveau système ou utilisateur et fournissent les mêmes fonctionnalités qu’un démon _cron_ traditionnel. Sauf si des capacités supplémentaires sont nécessaires, installer un planificateur additionnel n’est généralement par nécessaire et peut être passé.

#### modprobed-db users

If using [sys-kernel/modprobed-db](https://packages.gentoo.org/packages/sys-kernel/modprobed-db) was picked as an option for manually compiling the kernel.

Now is the time to setup a systemd timer to periodically scan the system for hardware used.

`root #` `systemctl --user enable modprobed-db`

At a later date of at least a week, please visit the kernel build section of the [modprobed-db](/wiki/Modprobed-db#Building_kernels.2Ffr "Modprobed-db") article to complete the setup.

== Facultatif : Indexation des fichiers ==

Pour indexer le système de fichiers afin de fournir des fonctions de recherche plus rapides, installez [mlocate](/wiki/Mlocate "Mlocate").

`root #` `emerge --ask sys-apps/mlocate`

## Facultatif : Accès distant

**Conseil**

La configuration par défaut opensshd n’autorise pas l’authentification _root_ pour un usage distant. [Créeez un utilisateur non- _root_](/wiki/FAQ/fr#How_do_I_add_a_normal_user.3F "FAQ/fr") et paramétrez-le pour permettre post-connexion ou modifiez /etc/ssh/sshd\_config pour autoriser _root_ (déconseillé).

Pour pouvoir accéder au système à distance après l’installation, sshd doit être configuré pour être lancé au démarrage.

Pour une plongée en profondeur dans les détails de la configuration SSH, consulter l’article [SSH](/wiki/SSH/fr "SSH/fr").

### OpenRC

Pour ajouter le script sshd au niveau d’exécution par défaut, sur OpenRC :

`root #` `rc-update add sshd default`

`root #` `rc-update add sshd default`

Si l’accès à la console série est nécessaire (ce qui est possible dans le cas de serveurs distants), agetty doit être configuré.

Décommenter la section sur la console série dans /etc/inittab :

`root #` `nano -w /etc/inittab`

```
# SERIAL CONSOLES
s0:12345:respawn:/sbin/agetty 9600 ttyS0 vt100
s1:12345:respawn:/sbin/agetty 9600 ttyS1 vt100

```

### systemd

Pour activer le serveur SSH, lancer :

`root #` `systemctl enable sshd`

`root #` `systemctl enable sshd`

Pour activer la console série, lancer :

`root #` `systemctl enable getty@tty1.service`

`root #` `systemctl enable getty@tty1.service`

## Facultatif : Autocomplétion _shell_

### Bash

Bash est le _shell_ par défaut sur les systèmes Gentoo aussi, installer les extensions d’autocomplétion peuvent apporter un confort et une efficacité pour gérer le système. Le paquet [app-shells/bash-completion](https://packages.gentoo.org/packages/app-shells/bash-completion) installera l’autocomplétion pour les commandes spécifiques Gentoo, ainsi que beaucoup d’autres commandes et utilitaires :

`root #` `emerge --ask app-shells/bash-completion`

La post-installation des commandes spécifiques d’autocomplétion bash peuvent être gérés avec eselect. Voyez la [section sur l’intégration de l’autocomplétion](/wiki/Bash#Shell_completion_integrations.2Ffr "Bash") dans l’article bash pour plus de détails.

## Suggéré : Synchronisation du temps

**Important**

Systems without a functioning [Real-Time Clock (RTC)](/wiki/System_time#Software_clock_vs_Hardware_clock "System time") must set the [system time](/wiki/System_time "System time") at every system start, and on regular intervals thereafter.

Il est important d’avoir un mécanisme de synchronisation de l’horloge système. C’est usuellement accompli avec le protocole [NTP](/wiki/NTP "NTP") et ce logiciel. D’autres implémentations de ce protocole NTP existent, comme [Chrony](/wiki/Chrony "Chrony").

The clock is usually synchronized via the [Network Time Protocol](/wiki/Network_Time_Protocol "Network Time Protocol"), with an implementation such as [chrony](/wiki/Chrony "Chrony").

### chrony

Pour installer Chrony, par exemple :

`root #` `emerge --ask net-misc/chrony`

`root #` `emerge --ask net-misc/chrony`

See the [chrony](/wiki/Chrony "Chrony") article for further information, for example if more advanced configurations are required.

### OpenRC

Sur OpenRC, lancez :

`root #` `rc-update add chronyd default`

`root #` `rc-update add chronyd default`

### systemd

Sur systemd, lancez :

`root #` `systemctl enable chronyd.service`

`root #` `systemctl enable chronyd.service`

### systemd-timesyncd

Alternativement, les utilisateurs systemd pourraient vouloir utiliser le simple client SNTP systemd-timesyncd qui est installé par défaut.

`root #` `systemctl enable systemd-timesyncd.service`

`root #` `systemctl enable systemd-timesyncd.service`

## Outils de système de fichiers

En fonction des systèmes de fichiers utilisés, il peut être nécessaire d’installer les utilitaires de systèmes de fichiers requis (pour vérifier l'intégrité su système de fichiers, reformatter, etc.). Noter que les outils pour gérer les système de fichiers ext4 ([sys-fs/e2fsprogs](https://packages.gentoo.org/packages/sys-fs/e2fsprogs)) sont déjà installés dans le cadre de [l’ensemble @system](/wiki/System_set_(Portage) "System set (Portage)").

Le tableau suivant liste les outils à installer si un certain système de fichiers est utilisé :

Filesystem
Package
[XFS](/wiki/XFS/fr "XFS/fr")[sys-fs/xfsprogs](https://packages.gentoo.org/packages/sys-fs/xfsprogs)[ext4](/wiki/Ext4/fr "Ext4/fr")[sys-fs/e2fsprogs](https://packages.gentoo.org/packages/sys-fs/e2fsprogs)[VFAT](/wiki/VFAT "VFAT") (FAT32, ...)
[sys-fs/dosfstools](https://packages.gentoo.org/packages/sys-fs/dosfstools)[Btrfs](/wiki/Btrfs/fr "Btrfs/fr")[sys-fs/btrfs-progs](https://packages.gentoo.org/packages/sys-fs/btrfs-progs)[F2FS](/wiki/F2FS/fr "F2FS/fr")[sys-fs/f2fs-tools](https://packages.gentoo.org/packages/sys-fs/f2fs-tools)[NTFS](/wiki/NTFS "NTFS")[sys-fs/ntfs3g](https://packages.gentoo.org/packages/sys-fs/ntfs3g)[ZFS](/wiki/ZFS "ZFS")[sys-fs/zfs](https://packages.gentoo.org/packages/sys-fs/zfs)[bcachefs](/wiki/Bcachefs "Bcachefs")[sys-fs/bcachefs-tools](https://packages.gentoo.org/packages/sys-fs/bcachefs-tools)

Il est recommandé que [sys-block/io-scheduler-udev-rules](https://packages.gentoo.org/packages/sys-block/io-scheduler-udev-rules) soit installé pour un comportement correct du planificateur avec les périphériques nvme :

`root #` `emerge --ask sys-block/io-scheduler-udev-rules`

**Conseil**

Pour plus d’nformations sur les systèmes de fichiers dans Gentoo, se référer à l’article sur les [systèmes de fichiers](/wiki/Filesystem/fr "Filesystem/fr").

Maintenant, continuez avec [Configuration du système d’amorçage](/wiki/Handbook:PPC64/Installation/Bootloader/fr "Handbook:PPC64/Installation/Bootloader/fr").

[← Configurer le Système](/wiki/Handbook:PPC64/Installation/System/fr "Previous part") [Accueil](/wiki/Handbook:PPC64/fr "Handbook:PPC64/fr") [Configurer le chargeur d’amorçage →](/wiki/Handbook:PPC64/Installation/Bootloader/fr "Next part")