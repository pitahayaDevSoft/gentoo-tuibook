# Base

Other languages:

- [Deutsch](/wiki/Handbook:PPC64/Installation/Base/de "Handbuch:PPC64/Installation/Basis (100% translated)")
- [English](/wiki/Handbook:PPC64/Installation/Base "Handbook:PPC64/Installation/Base (100% translated)")
- [español](/wiki/Handbook:PPC64/Installation/Base/es "Manual de Gentoo: PPC64/Instalación/Base (100% translated)")
- français
- [italiano](/wiki/Handbook:PPC64/Installation/Base/it "Handbook:PPC64/Installation/Base/it (100% translated)")
- [magyar](/wiki/Handbook:PPC64/Installation/Base/hu "Handbook:PPC64/Installation/Base/hu (100% translated)")
- [polski](/wiki/Handbook:PPC64/Installation/Base/pl "Handbook:PPC64/Installation/Base (100% translated)")
- [português do Brasil](/wiki/Handbook:PPC64/Installation/Base/pt-br "Handbook:PPC64/Installation/Base/pt-br (50% translated)")
- [русский](/wiki/Handbook:PPC64/Installation/Base/ru "Handbook:PPC64/Installation/Base (100% translated)")
- [தமிழ்](/wiki/Handbook:PPC64/Installation/Base/ta "கையேடு:PPC64/நிறுவல்/அடிப்படை (100% translated)")
- [中文（中国大陆）‎](/wiki/Handbook:PPC64/Installation/Base/zh-cn "手册：PPC64/安装/安装基本系统 (100% translated)")
- [日本語](/wiki/Handbook:PPC64/Installation/Base/ja "ハンドブック:PPC64/インストール/ベース (100% translated)")
- [한국어](/wiki/Handbook:PPC64/Installation/Base/ko "Handbook:PPC64/Installation/Base/ko (100% translated)")

[Sommaire du manuel](/wiki/Handbook:PPC64 "Handbook:PPC64")[Installation](/wiki/Handbook:PPC64/Full/Installation/fr "Handbook:PPC64/Full/Installation/fr")[À propos de l’installation](/wiki/Handbook:PPC64/Installation/About/fr "Handbook:PPC64/Installation/About/fr")[Choix du support](/wiki/Handbook:PPC64/Installation/Media/fr "Handbook:PPC64/Installation/Media/fr")[Configurer le réseau](/wiki/Handbook:PPC64/Installation/Networking/fr "Handbook:PPC64/Installation/Networking/fr")[Préparer les disques](/wiki/Handbook:PPC64/Installation/Disks/fr "Handbook:PPC64/Installation/Disks/fr")[Installer l’archive _stage3_](/wiki/Handbook:PPC64/Installation/Stage/fr "Handbook:PPC64/Installation/Stage/fr")Installer le système de base[Configurer le noyau](/wiki/Handbook:PPC64/Installation/Kernel/fr "Handbook:PPC64/Installation/Kernel/fr")[Configurer le système](/wiki/Handbook:PPC64/Installation/System/fr "Handbook:PPC64/Installation/System/fr")[Installer les outils](/wiki/Handbook:PPC64/Installation/Tools/fr "Handbook:PPC64/Installation/Tools/fr")[Configurer le système d’amorçage](/wiki/Handbook:PPC64/Installation/Bootloader/fr "Handbook:PPC64/Installation/Bootloader/fr")[Finaliser](/wiki/Handbook:PPC64/Installation/Finalizing/fr "Handbook:PPC64/Installation/Finalizing/fr")[Utiliser Gentoo](/wiki/Handbook:PPC64/Full/Working/fr "Handbook:PPC64/Full/Working/fr")[Introduction à Portage](/wiki/Handbook:PPC64/Working/Portage/fr "Handbook:PPC64/Working/Portage/fr")[Les options de la variable USE](/wiki/Handbook:PPC64/Working/USE/fr "Handbook:PPC64/Working/USE/fr")[Les fonctionnalités de Portage](/wiki/Handbook:PPC64/Working/Features/fr "Handbook:PPC64/Working/Features/fr")[Scripts d’initialisation systèmes](/wiki/Handbook:PPC64/Working/Initscripts/fr "Handbook:PPC64/Working/Initscripts/fr")[Variables d’environnement](/wiki/Handbook:PPC64/Working/EnvVar/fr "Handbook:PPC64/Working/EnvVar/fr")[Utiliser Portage](/wiki/Handbook:PPC64/Full/Portage/fr "Handbook:PPC64/Full/Portage/fr")[Fichiers et répertoires](/wiki/Handbook:PPC64/Portage/Files/fr "Handbook:PPC64/Portage/Files/fr")[Les variables](/wiki/Handbook:PPC64/Portage/Variables/fr "Handbook:PPC64/Portage/Variables/fr")[Mélanger plusieurs branches logicielles](/wiki/Handbook:PPC64/Portage/Branches/fr "Handbook:PPC64/Portage/Branches/fr")[Outils supplémentaires](/wiki/Handbook:PPC64/Portage/Tools/fr "Handbook:PPC64/Portage/Tools/fr")[Dépôt personnalisé](/wiki/Handbook:PPC64/Portage/CustomTree/fr "Handbook:PPC64/Portage/CustomTree/fr")[Fonctionnalités avancées](/wiki/Handbook:PPC64/Portage/Advanced/fr "Handbook:PPC64/Portage/Advanced/fr")[Configuration du réseau avec OpenRC](/wiki/Handbook:PPC64/Full/Networking/fr "Handbook:PPC64/Full/Networking/fr")[Bien démarrer](/wiki/Handbook:PPC64/Networking/Introduction/fr "Handbook:PPC64/Networking/Introduction/fr")[Configuration avancée](/wiki/Handbook:PPC64/Networking/Advanced/fr "Handbook:PPC64/Networking/Advanced/fr")[Les modules réseau](/wiki/Handbook:PPC64/Networking/Modular/fr "Handbook:PPC64/Networking/Modular/fr")[Sans fil](/wiki/Handbook:PPC64/Networking/Wireless/fr "Handbook:PPC64/Networking/Wireless/fr")[Ajouter des fonctionnalités](/wiki/Handbook:PPC64/Networking/Extending/fr "Handbook:PPC64/Networking/Extending/fr")[Gestion dynamique](/wiki/Handbook:PPC64/Networking/Dynamic/fr "Handbook:PPC64/Networking/Dynamic/fr")

## Contents

- [1Chrootage](#Chrootage)
  - [1.1Copier les informations DNS](#Copier_les_informations_DNS)
  - [1.2Monter les partitions nécessaires](#Monter_les_partitions_n.C3.A9cessaires)
  - [1.3Entrer dans le nouvel environnement](#Entrer_dans_le_nouvel_environnement)
  - [1.4Préparer le programme d’amorçage ( _bootloader_)](#Pr.C3.A9parer_le_programme_d.E2.80.99amor.C3.A7age_.28bootloader.29)
    - [1.4.1Système DOS/BIOS dépréciés](#Syst.C3.A8me_DOS.2FBIOS_d.C3.A9pr.C3.A9ci.C3.A9s)
- [2Configurer Portage](#Configurer_Portage)
  - [2.1Installer un instantané du dépôt _ebuild_ Gentoo depuis Internet](#Installer_un_instantan.C3.A9_du_d.C3.A9p.C3.B4t_ebuild_Gentoo_depuis_Internet)
  - [2.2Facultatif : sélectionner les miroirs](#Facultatif_:_s.C3.A9lectionner_les_miroirs)
    - [2.2.1Optional: rsync mirrors](#Optional:_rsync_mirrors)
  - [2.3Facultatif : Mettre à jour le dépôt _ebuild_ de Gentoo](#Facultatif_:_Mettre_.C3.A0_jour_le_d.C3.A9p.C3.B4t_ebuild_de_Gentoo)
  - [2.4Lire les nouvelles](#Lire_les_nouvelles)
  - [2.5Choisir le bon profil](#Choisir_le_bon_profil)
  - [2.6Facultatif : Ajouter un hôte pour paquets binaires](#Facultatif_:_Ajouter_un_h.C3.B4te_pour_paquets_binaires)
    - [2.6.1Configuration du dépôt](#Configuration_du_d.C3.A9p.C3.B4t)
    - [2.6.2Installer des paquets binaires](#Installer_des_paquets_binaires)
  - [2.7Facultatif : Configuration de la variable USE](#Facultatif_:_Configuration_de_la_variable_USE)
    - [2.7.1CPU\_FLAGS\_\*](#CPU_FLAGS_.2A)
    - [2.7.2VIDEO\_CARDS](#VIDEO_CARDS)
  - [2.8Optionnel : Configurer la variable ACCEPT\_LICENSE](#Optionnel_:_Configurer_la_variable_ACCEPT_LICENSE)
  - [2.9Mettre à jour l’ensemble ( _set_) @world](#Mettre_.C3.A0_jour_l.E2.80.99ensemble_.28set.29_.40world)
    - [2.9.1Supprimer les paquets obsolètes](#Supprimer_les_paquets_obsol.C3.A8tes)
- [3Fuseau horaire](#Fuseau_horaire)
- [4Configurer les paramètres régionaux ( _locale_)](#Configurer_les_param.C3.A8tres_r.C3.A9gionaux_.28locale.29)
  - [4.1Génération des paramètres régionaux](#G.C3.A9n.C3.A9ration_des_param.C3.A8tres_r.C3.A9gionaux)
  - [4.2Sélection du paramètre régional](#S.C3.A9lection_du_param.C3.A8tre_r.C3.A9gional)
- [5Références](#R.C3.A9f.C3.A9rences)

## Chrootage

### Copier les informations DNS

Il reste une chose à faire avant d’entrer dans le nouvel environnement et c’est de copier les informations DNS dans /etc/resolv.conf. C'est nécessaire afin de s’assurer que le réseau fonctionne toujours même après être entré dans le nouvel environnement. /etc/resolv.conf contient les serveurs de nom pour le réseau.

Pour copier ces informations, il est recommandé de passer l’option `--dereference` à la commande cp. Cela permet de s’assurer que, si /etc/resolv.conf est un lien symbolique, la cible du lien est copiée à la place du lien lui-même. Le lien symbolique dans le nouvel environnement ponterait autrement vers un fichier non existant (vu que la cible du lien n’existe probablement pas dans le nouvel environnement).

`root #` `cp --dereference /etc/resolv.conf /mnt/gentoo/etc/`

### Monter les partitions nécessaires

**Conseil**

Depuis un support officiel Gentoo, cette étape peut être remplacée simplement par : arch-chroot /mnt/gentoo.

Dans quelques instants, la racine Linux sera modifiée vers le nouvel emplacement.

Les partitions qui doivent être rendues disponibles sont :

- /proc/ qui est un pseudo système de fichiers. Il ressemble à des fichiers normaux, mais est en fait généré à la volée par le noyau Linux;
- /sys/ qui est un pseudo système de fichiers, comme /proc/ qu’il était autrefois sensé remplacer, et il est plus structuré que /proc/ ;
- /dev/ est un système de fichiers régulier qui contient les périphériques. Il est partiellement géré par le gestionnaire de périphérique Linux (généralement udev) ;
- /run/ est un système de fichiers temporaire utilisé pour des fichiers générés à l’exécution, comme des fichiers _PID_ (comportement le numéro de processus d’un service) ou des verrous.

L’emplacement /proc/ sera monté sur /mnt/gentoo/proc/ alors que les autres seront remontés ailleurs. Ce dernier signifie que, par exemple, /mnt/gentoo/sys/ sera en fait /sys/ (c’est juste un deuxième point d’entrée sur le même système de fichiers) alors que /mnt/gentoo/proc/ est un nouveau montage (nouvelle instance pour ainsi dire) du système de fichiers.

`root #` `mount --types proc /proc /mnt/gentoo/proc
`

`root #` `mount --rbind /sys /mnt/gentoo/sys
`

`root #` `mount --make-rslave /mnt/gentoo/sys
`

`root #` `mount --rbind /dev /mnt/gentoo/dev
`

`root #` `mount --make-rslave /mnt/gentoo/dev
`

`root #` `mount --bind /run /mnt/gentoo/run
`

`root #` `mount --make-slave /mnt/gentoo/run
`

**Remarque**

Les options `--make-rslave` ne sont nécessaires que pour supporter systemd plus loin dans l’installation.

**Attention !**

Lorsqu’un support d’installation non officiel de Gentoo est utilisé, cela peut ne pas suffire. Certaines distributions font de /dev/shm un lien symbolique vers /run/shm/ qui, après le _chroot_, devient invalide. Faire de /dev/shm/ un montage tmpfs correct d’entrée permet de corriger ce problème :

`root #` `test -L /dev/shm && rm /dev/shm && mkdir /dev/shm
`

`root #` `mount --types tmpfs --options nosuid,nodev,noexec shm /dev/shm
`

Assurez-vous également que le mode 1777 est appliqué :

`root #` ` chmod 1777 /dev/shm /run/shm`

### Entrer dans le nouvel environnement

Maintenant que toutes les partitions sont initialisées et que l’environnement de base est installé, il est temps d'entrer dans le nouvel environnement d’installation en utilisant chroot. Cela signifie que la session changera de racine (emplacement de plus haut niveau pouvant être atteint) depuis l’environnement d’installation courant (cédérom d'’installation ou autre support) vers le système d’installation (à savoir les partitions précédemment initialisées). D’où le nom _change root_ ou _chroot_.

Ce processus de _chroot_ se déroule en trois étapes :

1. l’emplacement de la racine est changé de / (sur le support d’installation) à /mnt/gentoo/ (sur les partitions) en utilisant _chroot_ ou _arch-chroot_, si disponible ;
2. certains paramètres (situés dans /etc/profile) sont rafraichis en mémoire en utilisant la commande source ;
3. l’invite de commande principale est modifiée afin de se rappeler plus facilement que cette session se situe dans un environnement _chroot_.

`root #` `chroot /mnt/gentoo /bin/bash
`

`root #` `source /etc/profile
`

`root #` `export PS1="(chroot) ${PS1}"`

À partir de maintenant, toutes les actions réalisées le sont dans le nouvel environnement Gentoo.

**Conseil**

Si l’installation de Gentoo est interrompue n’importe où après ce point, il _devrait_ être possible de reprendre l’installation depuis cette étape. Il n'y a pas besoin de refaire le partitionnement des disques ! Il suffit simplement de [monter la partition racine](/wiki/Handbook:PPC64/Installation/Disks/fr#Mounting_the_root_partition "Handbook:PPC64/Installation/Disks/fr") et d’exécuter les commandes ci-dessus depuis la [copie des informations DNS](/wiki/Handbook:PPC64/Installation/Base/fr#Copy_DNS_info "Handbook:PPC64/Installation/Base/fr") pour réintégrer l’environnement de travail. Ceci est également utile pour résoudre les problèmes de chargeur d’amorçage. Plus d’informations peuvent être trouvées dans l’article sur [chroot](/wiki/Chroot "Chroot").

### Préparer le programme d’amorçage ( _bootloader_)

Maintenant que vous êtes entré dans le nouvel environnement, il est nécessaire de préparer une partition pour le programme d’amorçage. Cela est important d’avoir la bonne partition montée lors de l’installation ce programme.

#### Système DOS/BIOS dépréciés

Pour les systèmes DOS/BIOS dépréciés, le programme d’amorçage devra être installé dans le répertoire /boot, ensuite monté selon :

`root #` `mount /dev/sda1 /boot`

## Configurer Portage

### Installer un instantané du dépôt _ebuild_ Gentoo depuis Internet

L’étape suivant consiste à installer un instantané du dépôt _ebuild_ Gentoo. Cet instantané contient une collection de fichiers qui informent Portage des logiciels disponibles (pour installation), quels profils l’administrateur système peut sélectionner, des informations spécifiques aux paquets ou profils, etc.

L’utilisation de la commande emerge-webrsync est recommandée pour ceux situés derrière des pare-feu restrictifs (elle utilise les protocoles HTTP/FTP pour télécharger l’instantané) et économise de la bande passante. Les lecteurs n’ayant pas de restriction réseau ou de bande passante peuvent passer directement à la section suivante.

Ceci va récupérer le dernier instantané (qui est publié quotidiennement) depuis l’un des miroirs de Gentoo et l’installer sur le système :

`root #` `emerge-webrsync`

**Remarque**

Pendant cette opération, emerge-webrsync peut se plaindre d’un emplacement /var/db/repos/gentoo/ inexistant. Cela est à prévoir et n’est en rien inquiétant – l’outil se chargera lui-même de créer l’emplacement.

À partir de ce moment, Portage peut mentionner que l’exécution de certaines mises à jour soit recommandée. Cela s’explique par le fait que certains paquets du système puissent avoir des versions plus récentes disponibles ; Portage est dès maintenant au courant des nouvelles versions en raison de l’installation de l’instantané. Les mises à jour peuvent être ignorées en toute sécurité pour l’instant ; les mises à jour peuvent être effectuées une fois l’installation de Gentoo terminée.

### Facultatif : sélectionner les miroirs

Afin de télécharger le code source rapidement, il est recommandé de sélectionner un miroir rapide et géographiquement proche. Portage cherche dans le fichier make.conf la variable GENTOO\_MIRRORS et utilise les miroirs listés à l’intérieur. Il est possible de naviguer vers la liste des miroirs de Gentoo et de chercher celui (ou ceux) qui se situe le plus près de la position géographique du système (ce sont souvent les plus rapides).

Un outil appelé mirrorselect fournit une interface textuelle sympathique pour permettre de rechercher et sélectionner plus rapidement le meilleur miroir. Naviguez simplement sur les miroirs et presser `Espace` pour choisir un ou plusieurs miroirs.

`root #` `emerge --ask --verbose --oneshot app-portage/mirrorselect`

`root #` `mirrorselect -i -o >> /etc/portage/make.conf`

Alternativement, une liste des miroirs actifs se trouve [en ligne](https://www.gentoo.org/downloads/mirrors/).

#### Optional: rsync mirrors

Gentoo also has many rsync mirrors which can speed up downloading the available package list using `emerge --sync` (explained in more detail later) by selecting a mirror closer that is geographically closer to the user.

`root #` `mkdir /etc/portage/repos.conf
`

`root #` `cp /usr/share/portage/config/repos.conf /etc/portage/repos.conf/gentoo.conf
`

Select a mirror close to the system's location from [https://www.gentoo.org/support/rsync-mirrors/](https://www.gentoo.org/support/rsync-mirrors/) and replace the sync-uri default mirror in /etc/portage/repos.conf/gentoo.conf with the desired mirror location.

FILE **`/etc/portage/repos.conf/gentoo.conf`** **UK-based rsync mirror example**

```
[DEFAULT]
main-repo = gentoo
</div>

<div lang="en" dir="ltr" class="mw-content-ltr">
[gentoo]
location = /var/db/repos/gentoo
sync-type = rsync
sync-uri = rsync://rsync.uk.gentoo.org/gentoo-portage/
auto-sync = yes
sync-rsync-verify-jobs = 1
sync-rsync-verify-metamanifest = yes
sync-rsync-verify-max-age = 3
sync-openpgp-key-path = /usr/share/openpgp-keys/gentoo-release.asc
sync-openpgp-keyserver = hkps://keys.gentoo.org
sync-openpgp-key-refresh-retry-count = 40
sync-openpgp-key-refresh-retry-overall-timeout = 1200
sync-openpgp-key-refresh-retry-delay-exp-base = 2
sync-openpgp-key-refresh-retry-delay-max = 60
sync-openpgp-key-refresh-retry-delay-mult = 4
sync-webrsync-verify-signature = yes
sync-git-verify-commit-signature = true

```

### Facultatif : Mettre à jour le dépôt _ebuild_ de Gentoo

Il est possible de mettre à jour le dépôt ebuild de Gentoo vers la dernière version. La commande précédente emerge-webrsync aura installé un instantané récent (généralement moins de 24 h) donc cette étape reste optionnelle.

S’il est cependant nécessaire d’avoir la version la plus récente du dépôt (moins d’une heure), utiliser emerge --sync. Cette commande utiliser le protocole _rsync_ pour mettre à jour le dépôt _ebuild_ de Gentoo (qui fut récupéré plus tôt via emerge-webrsync) vers l’état le plus récent.

`root #` `emerge --sync`

Sur les terminaux lents, comme certains tampon de trame ( _framebuffers_) ou consoles, il est recommandé d’utiliser l’option `--quiet` pour accélérer le processus.

`root #` `emerge --sync --quiet`

### Lire les nouvelles

Quand le dépôt _ebuild_ de Gentoo est synchronisé sur le système, Portage peut afficher des messages informatifs similaires à ceux-ci :

` * IMPORTANT: 2 news items need reading for repository 'gentoo'.
`

` * Use eselect news to read news items.
`

Les nouvelles furent créées afin de fournir un moyen de communication permettant d’envoyer des messages critiques aux utilisateurs via le dépôt _ebuild_ de Gentoo. Pour les gérer, utiliser eselect news. L’application eselect est un utilitaire spécifique à Gentoo qui permet d’avoir une interface de gestion commune pour l’administration système. Ici, eselect est invité à utiliser son module de `news`.

Pour le module de `news`, trois opérations principales sont utilisées :

- `list` est un aperçu des nouvelles disponibles ;
- `read` permet de lire les nouvelles ;
- `purge` supprime les nouvelles une fois lues et n’ont plus à être relue.

`root #` `eselect news list
`

`root #` `eselect news read`

Plus d’information sur le lecteur de nouvelles est disponible via sa page de manuel :

`root #` `man news.eselect`

### Choisir le bon profil

**Conseil**

Les profils de bureau ne sont pas conçus seulement pour les « environnements de bureau ». Ils sont également indiqué pour les gestionnaires de fenêtres minimalistes comme i3 ou Sway.

Un « profil » est un élément de construction pour tout système Gentoo. Non seulement il spécifie des valeurs par défaut pour USE, CFLAGS, et autres variables importantes. Il limite aussi aussi le système à une certaine gamme de versions des paquets. Ces paramètres sont tous gérés par les développeurs Portage de Gentoo.

Pour voir quel profil le système utilise actuellement, lancer eselect avec le module `profile` :

**Conseil**

On an install media without a scroll-able terminal, `eselect profile list | less` can provide an easy way to list all available profiles while providing the ability to scroll.

`root #` `eselect profile list`

```
Available profile symlink targets:
  [1]   default/linux/ppc64/23.0 *
  [2]   default/linux/ppc64/23.0/desktop
  [3]   default/linux/ppc64/23.0/desktop/gnome
  [4]   default/linux/ppc64/23.0/desktop/kde

```

**Remarque**

Le résultat de la commande n’est qu'un exemple et évolue avec le temps.

**Remarque**

Pour utiliser _systemd_, sélectionnez un profil qui a « systemd » dans son nom et _vice versa_ sinon.

Il existe également des sous-profils pour certaines architectures qui incluent des paquets applicatifs couramment nécessaire pour une expérience utilisateur de bureau.

**Attention !**

Les mises à niveau de profil ne doivent pas être prises à la légère. Lors de la sélection du profil initial, veillez à utiliser le profil correspondant à **la même version** que celle initialement utilisée par _stage3_ (par exemple : 23.0). Chaque nouvelle version de profil est annoncée via une news contenant des instructions de migration. Assurez-vous de suivre attentivement ces instructions avant de passer à un nouveau profil.

Après avoir visionné les profils disponibles pour l’architecture ppc64, les utilisateurs peuvent sélectionner un profil différent pour le système :

`root #` `eselect profile set 2`

**Remarque**

Le sous-profil `developer` est spécifique au développement de Gentoo Linux et ne doit pas être utilisé par des utilisateurs normaux.

### Facultatif : Ajouter un hôte pour paquets binaires

Depuis décembre 2023,la [Release Engineering team](/wiki/Project:RelEng "Project:RelEng") de Gentoo offre un [hôte pour paquets binaires officiel](/wiki/Binary_package_quickstart "Binary package quickstart")
(communément appelé « _binhost_ ») pour permettre à la communauté de récupérer et installer des paquets binaires ( _binpkgs_)[\[1\]](#cite_note-1).

Ajouter un hôte de paquets binaires autorise Portage à installer des paquets compilés signés. Dans beaucoup de situations, ajouter un tel hôte diminue _fortement_ le temps moyen d’installation d’un paquet et offre des avantages pour faire fonctionner Gentoo sur des systèmes anciens, vieux ou économes.

#### Configuration du dépôt

La configuration pour un _binhost_ se trouve dans le répertoire Portage /etc/portage/binrepos.conf/, qui fonctionne comme mentionné dans la section [dépôt _ebuild_ Gentoo](/wiki/Handbook:PPC64/Installation/Base/fr#Gentoo_ebuild_repository "Handbook:PPC64/Installation/Base/fr").

Lors de la définition d’un hôte de paquets binaires, les deux aspects les plus importants à considérer sont :

1. La cible d’architecture et de profil avec `sync-uri` _compte_ et devrait être identique à l’architecture de l’ordinateur ( **ppc64**, dans ce cas) et le profil système sélectionné dans la section [Choisir le bon profil](/wiki/Handbook:PPC64/Installation/Base/fr#Choosing_the_right_profile "Handbook:PPC64/Installation/Base/fr").
2. Choisir un miroir rapide et proche géographiquement va en général diminuer le temps de téléchargement. Consultez l’outil mirrorselect mentionné dans la section [Facultatif : Choisir un mioir](/wiki/Handbook:PPC64/Installation/Base/fr#Gentoo_ebuild_repository "Handbook:PPC64/Installation/Base/fr") ou regardez la [liste des miroirs en ligne](https://www.gentoo.org/downloads/mirrors/) où les URL peuvent être retrouvées.


FILE **`/etc/portage/binrepos.conf/gentoobinhost.conf`** **Exemple d’hôte de paquets binaires**

```
[binhost]
priority = 9999
sync-uri = https://distfiles.gentoo.org/releases/<arch>/binpackages/<profile>/x86-64/

```

1. Introduced in portage-3.0.74 for per-repo verification choices

verify-signature = true

1. Default value with >=portage-3.0.77

location = /var/cache/binhost/gentoo
}}

#### Installer des paquets binaires

Portage compilera les paquets depuis le code source par défaut. Il peut lui être demandé d’utiliser un paquet binaire avec ce qui suit :

1. l’option `--getbinpkg` lors de l’appel à la commande emerge ; cette méthode est utile pour installer un paquet binaire en particulier ;
2. changer la valeur par défaut de la variable FEATURES de Portage, ce qui fait dans le fichier /etc/portage/make.conf ; appliquer ce paramètre fera que Portage essayera de télécharger le paquet binaire et, s’il n’en a pas, compilera le code source.

Par exemple, pour que Portage installe toujours les paquets binaires :

FILE **`/etc/portage/make.conf`** **Configurer Portage pour utiliser les paquets binaires par défaut**

```
# Ajout getbinpkg à la liste des paramètre de la variable FEATURES
FEATURES="${FEATURES} getbinpkg"
# Force la vérification de la signature
FEATURES="${FEATURES} binpkg-request-signature"

```

Lancez également getuto pour que Portage mette en place le nécessaire pour la vérification des clés :

`root #` `getuto`

Des fonctionnalités supplémentaires de Portage sera détaillées dans le [prochain chapitre](/wiki/Handbook:PPC64/Working/Features/fr#Portage_features "Handbook:PPC64/Working/Features/fr") du manuel.

### Facultatif : Configuration de la variable USE

USE est l’une des variables les plus puissantes que Gentoo propose à l’utilisateur. Plusieurs programmes peuvent être compilés avec ou sans support facultatif pour certaines options. Par exemple, certains programmes peuvent être compilés avec le support pour GTK+ ou le support pour Qt. D’autres peuvent être compilés avec ou sans le support pour SSL. Certains programmes peuvent être compilés avec le support pour _framebuffer_ ( _svgalib_) au lieu du support pour _X11_ ( _X-server_).

La plupart des distributions compilent leurs paquets avec autant de supports que possible, augmentant la taille des programmes et les temps de démarrage, sans oublier de mentionner un nombre énorme de dépendances. Avec Gentoo, les utilisateurs peuvent choisir avec quelles options un package doit être compilé. C'est là que la variable USE entre en jeu.

Dans la variable USE, les utilisateurs définissent des mots-clés qui correspondent à des options du compilateur. Par exemple, `ssl` ajoutera le support de SSL dans les programmes qui le supporte. `-X` supprimera le support du serveur X. `gnome gtk -kde -qt5` compilera les programmes avec le support de GNOME (et de GTK+), mais sans le support de KDE (et Qt), ce rend le système complètement adapté pour gnome (si l’architecture le permet).

Les paramètres par défaut de la variable USE sont placés dans le fichier make.defaults du profil Gentoo utilisé par le système. Gentoo utilise un système d’héritage complexe pour ses profils, qui ne sera pas expliqué en profondeur durant la processus d’installation. Le moyen le plus simple de vérifier les paramètres de la variable USE actuellement actifs est d’exécuter emerge --info et de sélectionner la ligne commençant par USE :

`root #` `emerge --info | grep ^USE`

```
USE="X acl alsa amd64 berkdb bindist bzip2 cli cracklib crypt cxx dri ..."

```

**Remarque**

L’exemple ci-dessus est tronqué, la liste réelle des valeurs de la variable USE est beaucoup, beaucoup plus longue.

Un description complète des options de la variable USE peut se trouver sur le système dans /var/db/repos/gentoo/profiles/use.desc.

`root #` `less /var/db/repos/gentoo/profiles/use.desc`

À l’intérieur de le commande less, le défilement peut s’effectuer à l'aide des touches ` ↑` et ` ↓`, et le programme fermé en appuyant sur `q`.

Par exemple, voici les paramètres de la variable USE pour un système basé sur KDE avec le support pour DVD, ALSA et l’enregistrement de CD :

`root #` `nano /etc/portage/make.conf`

FILE **`/etc/portage/make.conf`** **Activer les drapeaux USE pour un système basé sur KDE/Plasma avec le support pour DVD, ALSA et l’enregistrement de CD**

```
USE="-gtk -gnome qt5 kde dvd alsa cdr"

```

Quand la variable USE est définie dans /etc/portage/make.conf, elle s’ **ajoute** à la liste des drapeaux du système. Les drapeaux USE peuvent être **supprimés** totalement en ajoutant `-` devant une valeur de la liste. Par exemple, pour supprimer le support de l’environnement graphique X, `-X` peut être saisi :

FILE **`/etc/portage/make.conf`** **Ignorer les options par défaut de la variable USE**

```
USE="-X acl alsa"

```

**Attention !**

Bien que possible, utiliser `-*` (lequel va désactiver tous les drapeaux USE, sauf ceux définis dans make.conf) est **fortement** découragé et peu judicieux. Les développeurs d’ _ebuild_ choisissent certaines valeurs de drapeaux dans les _ebuilds_ de manière à prévenir les conflits, améliorer la sécurité, éviter des erreurs et plein d’autres raisons. Désactiver _tous_ les drapeaux USE va annuler le comportement par défaut et pourrait causer de gros problèmes.

#### CPU\_FLAGS\_\*

Certaines architectures (incluant AMD64/X86, ARM et PPC) ont une variable [USE\_EXPAND](/wiki/USE_EXPAND "USE EXPAND") appelée [CPU\_FLAGS\_<ARCH>](/wiki/CPU_FLAGS_X86 "CPU FLAGS X86"), où <ARCH> est remplacé par le nom de l’architecture.

**Important**

Attention à ne pas confondre ! **AMD64** et **X86** partagent en commun certaines architectures, aussi, la variable pour un système **AMD64** est CPU\_FLAGS\_X86.

Ceci est utilisé pour configuré la compilation dans un code assemble spécifique ou un autre langage machine, en général écrit à la main ou autre, et n’est **pas** la même chose que de demander au compilateur de sortir du code optimisé pour certaines fonctionnalités CPU (comme `-march=`).

Les utilisateurs devraient paramétrer cette variable en ajoutant les  COMMON\_FLAGS désirés.

Quelques actions sont nécessaires pour paramétrer cela :

`root #` `emerge --ask --oneshot app-portage/cpuid2cpuflags`

Vérifiez la sortie si vous êtes curieux :

`root #` `cpuid2cpuflags`

Puis, copier la sortie dans package.use :

`root #` `echo "*/* $(cpuid2cpuflags)" > /etc/portage/package.use/00cpu-flags`

#### VIDEO\_CARDS

By default a profile already sets a few video cards. For many reasons this is not ideal, but a useful safety net.

To configure the system correctly the user needs to first unset the preset video cards with `VIDEO_CARDS: -*` then set the correct card for that system.

FILE **`/etc/portage/package.use/00video_cards`** **Example**

```
*/* VIDEO_CARDS: -* <GPU DRIVER NAME>

```

Ci-dessous un tableau pouvant être utilisé pour comparer facilement les différents types de cartes graphiques avec leur valeur `VIDEO_CARDS`.

Machine
Famille de carte vidéo
VIDEO\_CARDS
Intel x86AucuneVoir [Support Intel](/wiki/Intel#Feature_support "Intel")x86/ARMNvidia`nvidia`ToutesNvidia sauf Maxwell, Pascal et Volta`nouveau`ToutesAMD à partir de Sea Islands`amdgpu radeonsi`ToutesATI et les anciennes AMDVoir [Support Radeon](/wiki/Radeon#Feature_support "Radeon")ToutesIntel`intel`Raspberry PiN/A`vc4`QEMU/KVMToutes`virgl`WSLToutes`d3d12`

Ci-dessous un exemple d’un package.use correctement défini pour « VIDEO\_CARDS ». Adaptez le nom du pilote à utiliser.

FILE **`/etc/portage/package.use/00video_cards`**

```
*/* VIDEO_CARDS: amdgpu radeonsi

```

FILE **`/etc/portage/package.use/00video_cards`** **Intel example**

```
*/* VIDEO_CARDS: -* intel

```

FILE **`/etc/portage/package.use/00video_cards`** **Nvidia example**

```
*/* VIDEO_CARDS: -* nvidia

```

Plus de détails pour les différents GPU peuvent être trouvés dans les articles [AMDGPU](/wiki/AMDGPU "AMDGPU"), [Intel](/wiki/Intel "Intel"), [Nouveau (code ouvert)](/wiki/Nouveau/fr "Nouveau/fr") ou [NVIDIA (propriétaire)](/wiki/NVIDIA "NVIDIA").

### Optionnel : Configurer la variable ACCEPT\_LICENSE

Depuis la GLEP 23 ( _Gentoo Linux Enhancement Proposal_, le processus interne de proposition d’amélioration), un mécanisme a été ajouté pour permettre aux administrateurs de réguler les logiciels installés en fonction des licences… Certains voudront un système libre de tous logiciels approuvés par l’OSI ( _Open Source Initiative_) ; d’autres seront curieux de connaître quelles licences ils acceptent implicitement[\[2\]](#cite_note-2). La motivation finale étant d’avoir un contrôle plus fin sur quel type de logiciel tourne sur un système Gentoo, la variable ACCEPT\_LICENSE était née.

Durant le processus d’installation, Portage prend en compte les valeurs paramétrées dans la variable ACCEPT\_LICENSE pour déterminer si le paquet demandé répond à la définition de l’administrateur d’une licence acceptable. Ceci créé un autre problème : le dépôt d’ _ebuilds_ Gentoo contient des milliers de paquets ce qui correspond à des [_centaines_ de licences logicielles distinctes](https://gitweb.gentoo.org/repo/gentoo.git/tree/licenses)… Est-ce que cela force un administrateur système à approuver manuellement chaque licence logicielle ? Heureusement non ; la GLEP 23 répond également à cette problématique en créant des groupes de licences.

Pour une administration pratique du système, les licences logicielles similaires ont été regroupées ensemble – chacune en fonction de sa finalité. Les définitions de groupe de licence sont [accessibles en lecture](https://gitweb.gentoo.org/repo/gentoo.git/tree/profiles/license_groups) et gérée par le [Gentoo Licenses project](/wiki/Project:Licenses "Project:Licenses"). Contrairement aux licences individuelles, les groupes de licences sont préfixées par le symbole `@`, ce qui permet de facilement les distinguer dans la variable ACCEPT\_LICENSE.

Quelques groupes de licences communs incluent :

Une liste de licences logicielles groupées selon leur finalité.
Nom du groupeDescription
`@GPL-COMPATIBLE`Licences compatibles GPL approuvées par la _Free Software Foundation_ (FSF) [\[a\_license 1\]](#cite_note-3)`@FSF-APPROVED`Licences de logiciel libre approuvées par la FSF (contient `@GPL-COMPATIBLE`)
`@OSI-APPROVED`Licences approuvées par l’ _Open Source Initiative_ (OSI) [\[a\_license 2\]](#cite_note-4)`@MISC-FREE`Licences diverses qui sont probablement des logiciels libres, c’est-à-dire qui suivent la définition du logiciel libre [\[a\_license 3\]](#cite_note-5) mais qui ne sont approuvées ni par la FSF ni par l’OSI
`@FREE-SOFTWARE`Combine `@FSF-APPROVED`, `@OSI-APPROVED` et `@MISC-FREE`.
`@FSF-APPROVED-OTHER`Licences approuvées par la FSF pour « documentation libre » et « œuvres à usage pratique autres que les logiciels et la documentation » (polices de caractères incluses).
`@MISC-FREE-DOCS`Licences diverses pour les documents libres et autres œuvres (polices de caractères incluses) qui suivent la définition libre [\[a\_license 4\]](#cite_note-6) mais qui **ne sont pas** listées dans `@FSF-APPROVED-OTHER``@FREE-DOCUMENTS`Combine `@FSF-APPROVED-OTHER` et `@MISC-FREE-DOCS``@FREE`Méta-ensemble de toutes les licences avec liberté d’utilisation, partage, modification et partage de modifications.

Combine `@FREE-SOFTWARE` et `@FREE-DOCUMENTS`.

`@BINARY-REDISTRIBUTABLE`Licences qui permettent au moins la libre redistribution du logiciel sous forme de binaire. Contient @FREE.
`@EULA`Contrats de licences qui essaient de vous retirer des droits. Elles sont plus restrictives que « tous-droits-reservés » ou demandent un accord explicite

1. [↑](#cite_ref-3)[https://www.gnu.org/licenses/license-list.html](https://www.gnu.org/licenses/license-list.html)
2. [↑](#cite_ref-4)[https://www.opensource.org/licenses](https://www.opensource.org/licenses)
3. [↑](#cite_ref-5)[https://www.gnu.org/philosophy/free-sw.html](https://www.gnu.org/philosophy/free-sw.html)
4. [↑](#cite_ref-6)[https://freedomdefined.org/](https://freedomdefined.org/)

Actuellement, les licences acceptées par le système peuvent être connues via :

`user $` `portageq envvar ACCEPT_LICENSE`

```
@FREE
```

Comme lisible dans la sortie, la valeur par défaut est d’autoriser seulement l’installation de licences logicielles qui sont dans le groupe `@FREE`.

Des licences ou groupes de licences spécifiques pour un système peuvent être définies dans les emplacements suivants :

- au niveau système à l’intérieur du profil sélectionné – ceci paramètre la valeur par défaut ;
- au niveau système dans le fichier /etc/portage/make.conf ; cela permet aux administrateurs de surcharger la valeur du profil par défaut ;
- par paquet dans le fichier /etc/portage/package.license ;
- par paquet dans une arborescence de « répertoires » et fichiers /etc/portage/package.license/.

Le paramétrage système par défaut du profil peut être surchargé avec /etc/portage/make.conf :

FILE **`/etc/portage/make.conf`** **Accepter des licences au niveau système avec ACCEPT\_LICENSE**

```
# Surcharge la valeur par défaut du profil d’ACCEPT_LICENSE
ACCEPT_LICENSE="-* @FREE @BINARY-REDISTRIBUTABLE"

```

Optionnellement, les administrateurs systèmes peuvent définir les licences acceptées par paquet comme montré dans l’exemple de répertoires et fichiers suivants. Notez que le « répertoire » package.license doit être créé s’il n’existe pas déjà :

`root #` `mkdir /etc/portage/package.license`

Le détail des licences de chaque paquet Gentoo est stocké dans la variable LICENSE de l’ _ebuild_. Un paquet peut avoir une ou plusieurs licences, aussi, il peut être nécessaire d’ajouter plusieurs licences pour un seul paquet.

FILE **`/etc/portage/package.license/kernel`** **Accepter des licences par paquet**

```
app-arch/unrar unRAR
sys-kernel/linux-firmware linux-fw-redistributable
sys-firmware/intel-microcode intel-ucode

```

**Important**

La variable LICENSE d’un _ebuild_ n’est qu’une directive pour les développeurs Gentoo et les utilisateurs. Ce n’est **pas** une déclaration légale, et il n’y a aucune garantie que cela reflétera la réalité. Il est recommandé de ne pas se reposer seulement sur l’interprétation du développeur de l’ _ebuild_ d’une licence logicielle ; vérifiez le paquet en profondeur, y compris tous les fichiers qui ont été installés dans le système.

### Mettre à jour l’ensemble ( _set_) @world

**Conseil**

Si le profil d’un environnement de bureau est choisi depuis une archive _stage_ non bureau, la mise à jour de @world pourrait augmenter considérablement le temps nécessaire à l’installation du système. Ceux en manque de temps peuvent utiliser cette « règle de base » : plus le nom du profil est court, moins [l’ensemble @world](/wiki/World_set_(Portage) "World set (Portage)") ; moins l’ensemble @world ne sera spécifique, moins de paquets ne seront requis par le système. Autrement dit :

- choisir `default/linux/amd64/23.0` ne nécessitera la mise à jour que de peu de paquets ;
- alors que choisir `default/linux/amd64/23.0/desktop/gnome/systemd` nécessitera l’installation de plusieurs paquets car le profil choisi a un ensemble [@system](/wiki/Package_sets#.40system "Package sets") et [@profile](/wiki/Package_sets#.40profile "Package sets") plus grand : les dépendances pour supporter un environnement de bureau GNOME.

Mettre à jour l’ensemble système [@world](/wiki/World_set_(Portage) "World set (Portage)") est **optionnel** et va conduire à peu de changements fonctionnels, sauf si une ou plusieurs des étapes facultatives suivantes est entreprises :

1. un profil différent de celui de l’archive _stage_ a été sélectionné ;
2. des drapeaux USE additionnels ont été configurés pour des paquets installés.

Les lecteurs qui effectuent une « installation Gentoo rapide » pourraient passer la mise à jour de l’ensemble @world jusqu’au redémarrage du système dans un nouvel environnement Gentoo.

Les lecteurs qui effectuent une installation « lente » peuvent mettre à jour un paquet, un profil ou un drapeau USE maintenant :

`root #` `emerge --ask --verbose --update --deep --changed-use @world`

Les lecteurs qui ont ajouté un hôte de paquets binaires [ci-dessus](/wiki/Handbook:PPC64/Installation/Base/fr#Optional:_Adding_a_binary_package_host "Handbook:PPC64/Installation/Base/fr") peuvent ajouter --getbinpkg (ou -g) dans le but de récupérer les paquets depuis un hôte binaire au lieu de les compiler :

`root #` `emerge --ask --verbose --update --deep --newuse --getbinpkg @world`

#### Supprimer les paquets obsolètes

Il est important de toujours lancer « depclean » après une mise à jour du système pour supprimer les paquets obsolètes. Vérifiez attentivement la sortie avec emerge --depclean --pretend pour contrôler si vous utilisez personnellement un paquet à supprimer qui devrait être conservé. Pour conserver un tel paquet, utilisez emerge --noreplace foo.

`root #` `emerge --ask --pretend --depclean`

Si vous êtes satisfait, lancez le vrai nettoyage :

`root #` `emerge --ask --depclean`

## Fuseau horaire

**Remarque**

Cette étape ne concerne pas les utilisateurs de _musl libc_. Les utilisateurs qui ne comprennent pas ce que cela signifie peuvent mettre en œuvre cette étape.

**Attention !**

Veuillez éviter les fuseaux horaires tels que /usr/share/zoneinfo/Etc/GMT\* car leurs noms n’indiquent pas les zones attendues. Par exemple, GMT-8 est en réalité GMT+8.

Pour Sélectionner le fuseau horaire pour le système. Recherchez les fuseaux horaires disponibles dans /usr/share/zoneinfo/ :

`root #` `ls -l /usr/share/zoneinfo`

```
total 352
drwxr-xr-x 2 root root   1120 Jan  7 17:41 Africa
drwxr-xr-x 6 root root   2960 Jan  7 17:41 America
drwxr-xr-x 2 root root    280 Jan  7 17:41 Antarctica
drwxr-xr-x 2 root root     60 Jan  7 17:41 Arctic
drwxr-xr-x 2 root root   2020 Jan  7 17:41 Asia
drwxr-xr-x 2 root root    280 Jan  7 17:41 Atlantic
drwxr-xr-x 2 root root    500 Jan  7 17:41 Australia
drwxr-xr-x 2 root root    120 Jan  7 17:41 Brazil
-rw-r--r-- 1 root root   2094 Dec  3 17:19 CET
-rw-r--r-- 1 root root   2310 Dec  3 17:19 CST6CDT
drwxr-xr-x 2 root root    200 Jan  7 17:41 Canada
drwxr-xr-x 2 root root     80 Jan  7 17:41 Chile
-rw-r--r-- 1 root root   2416 Dec  3 17:19 Cuba
-rw-r--r-- 1 root root   1908 Dec  3 17:19 EET
-rw-r--r-- 1 root root    114 Dec  3 17:19 EST
-rw-r--r-- 1 root root   2310 Dec  3 17:19 EST5EDT
-rw-r--r-- 1 root root   2399 Dec  3 17:19 Egypt
-rw-r--r-- 1 root root   3492 Dec  3 17:19 Eire
drwxr-xr-x 2 root root    740 Jan  7 17:41 Etc
drwxr-xr-x 2 root root   1320 Jan  7 17:41 Europe
...

```

`root #` `ls -l /usr/share/zoneinfo/Europe/`

```
total 256
-rw-r--r-- 1 root root 2933 Dec  3 17:19 Amsterdam
-rw-r--r-- 1 root root 1742 Dec  3 17:19 Andorra
-rw-r--r-- 1 root root 1151 Dec  3 17:19 Astrakhan
-rw-r--r-- 1 root root 2262 Dec  3 17:19 Athens
-rw-r--r-- 1 root root 3664 Dec  3 17:19 Belfast
-rw-r--r-- 1 root root 1920 Dec  3 17:19 Belgrade
-rw-r--r-- 1 root root 2298 Dec  3 17:19 Berlin
-rw-r--r-- 1 root root 2301 Dec  3 17:19 Bratislava
-rw-r--r-- 1 root root 2933 Dec  3 17:19 Brussels
...

```

Imaginons que le fuseau horaire choisi soit Europe/Brussels, pour paramétrer cette zone, un lien symbolique peut être créer de la description de la zone vers /etc/localtime :

`root #` `ln -sf ../usr/share/zoneinfo/Europe/Brussels /etc/localtime`

**Conseil**

La cible avec `../` au début signifie « chemin relatif à la location du lien », pas le répertoire courant.

**Remarque**

Un chemin absolu peut être utilisé pour le lien symbolique, mais un chemin relatif est créé par timedatectl de _systemd_ et est plus compatible avec les racines alternatives.

## Configurer les paramètres régionaux ( _locale_)

**Remarque**

Cette étape ne concerne pas les utilisateurs de _musl libc_. Les utilisateurs qui ne comprennent pas ce que cela signifie peuvent mettre en œuvre cette étape.

### Génération des paramètres régionaux

A default installation of Gentoo Linux provides an archive that contains all supported locales, numbering 500 or more. However, it is typical for an administrator to require only one or two of these. In that case, the /etc/locale.gen configuration file may be populated with a list of the required locales. By default, locale-gen shall read this file and compile only the locales that are specified, saving both time and space in the longer term.

Les paramètres régionaux spécifient non seulement la langue pour interagir avec le système, mais aussi les règles pour trier les chaînes de caractères, afficher les dates et les heures, etc. Les paramètres régionaux sont sensibles à la casse et doivent être représentés exactement comme prévus. Une liste complète des paramètres régionaux disponibles se trouve dans le fichier /usr/share/i18n/SUPPORTED.

`root #` `nano /etc/locale.gen`

Les paramètres régionaux suivant sont un exemple pour avoir et l’anglais (États-Unis) et le français (France) avec les encodages de caractères correspondants (comme UTF-8).

FILE **`/etc/locale.gen`** **Activer les paramètres régionaux en/US et fr/FR avec les encodages de caractères correspondants**

```
en_US ISO-8859-1
en_US.UTF-8 UTF-8
fr_FR ISO-8859-1
fr_FR.UTF-8 UTF-8

```

1. Non UTF-8 locales are discouraged and should only be used in special circumstances.
2. en\_US ISO-8859-1
3. de\_DE ISO-8859-1

}}

**Attention !**

Beaucoup de logiciels nécessitent au moins un paramètre régional UTF-8 pour fonctionner correctement.

L’étape suivante consiste à exécuter la commande locale-gen. Elle génère tous les paramètres régionaux spécifiés dans le fichier /etc/locale.gen.

`root #` `locale-gen`

Pour vérifier que les paramètres régionaux sélectionnés sont maintenant disponibles, exécutez locale -a.

Pour une installation _systemd_, localectl peut être utilisé ; c’est-à-dire localectl set-locale … ou localectl list-locales.

### Sélection du paramètre régional

Une fois terminé, il est maintenant temps de définir les paramètres régionaux du système. Encore une fois, eselect est utilisé, cette fois avec le module `locale`.

Avec eselect locale list, les choix disponibles sont affichés :

`root #` `eselect locale list`

```
Available targets for the LANG variable:
  [1] C
  [2] C.utf8
  [3] en_US
  [4] en_US.iso88591
  [5] en_US.utf8
  [6] fr_FR
  [7] fr_FR.iso88591
  [8] fr_FR.iso885915
  [9] fr_FR.utf8
  [10] POSIX
  [ ] (free form)

```

Avec eselect locale set <NOMBRE>, les paramètres régionaux corrects peuvent être sélectionnés :

`root #` `eselect locale set 9`

Manuellement, cela peut être réalisé via le fichier /etc/env.d/02locale et pour _systemd_ le fichier /etc/locale.conf :

FILE **`/etc/env.d/02locale`** **Configurer manuellement les paramètres régionaux du système**

```
LANG="fr_FR.UTF-8"
LC_COLLATE="C.UTF-8"

```

Configurer un paramètre régional évitera des avertissements et erreurs lors des compilations du noyau et d’autres logiciels plus tard dans l’installation.

Mettre maintenant l’environnement à jour :

`root #` `env-update && source /etc/profile && export PS1="(chroot) ${PS1}"`

Pour une aide additionnelle sur le processus de sélection des paramètres régionaux, vous pouvez lire aussi le [guide de localisation](/wiki/Localization/Guide/fr "Localization/Guide/fr") et celui [UTF-8](/wiki/UTF-8/fr "UTF-8/fr").

## Références

1. [↑](#cite_ref-1)[https://www.gentoo.org/news/2023/12/29/Gentoo-binary.html](https://www.gentoo.org/news/2023/12/29/Gentoo-binary.html)
2. [↑](#cite_ref-2)[https://www.gentoo.org/glep/glep-0023.html#id7](https://www.gentoo.org/glep/glep-0023.html#id7)

[← Installer l’archive _stage3_](/wiki/Handbook:PPC64/Installation/Stage/fr "Previous part") [Accueil](/wiki/Handbook:PPC64/fr "Handbook:PPC64/fr") [Configurer le noyau →](/wiki/Handbook:PPC64/Installation/Kernel/fr "Next part")