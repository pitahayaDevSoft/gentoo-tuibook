# Stage

Other languages:

- [Deutsch](/wiki/Handbook:AMD64/Installation/Stage/de "Handbook:AMD64/Installation/Stage/de (100% translated)")
- [English](/wiki/Handbook:AMD64/Installation/Stage "Handbook:AMD64/Installation/Stage (100% translated)")
- [Türkçe](/wiki/Handbook:AMD64/Installation/Stage/tr "Handbook:AMD64/Installation/Stage/tr (0% translated)")
- [español](/wiki/Handbook:AMD64/Installation/Stage/es "Handbook:AMD64/Installation/Stage/es (100% translated)")
- français
- [italiano](/wiki/Handbook:AMD64/Installation/Stage/it "Handbook:AMD64/Installation/Stage/it (100% translated)")
- [magyar](/wiki/Handbook:AMD64/Installation/Stage/hu "Handbook:AMD64/Installation/Stage/hu (100% translated)")
- [polski](/wiki/Handbook:AMD64/Installation/Stage/pl "Handbook:AMD64/Installation/Stage/pl (100% translated)")
- [português do Brasil](/wiki/Handbook:AMD64/Installation/Stage/pt-br "Handbook:AMD64/Installation/Stage/pt-br (100% translated)")
- [čeština](/wiki/Handbook:AMD64/Installation/Stage/cs "Handbook:AMD64/Installation/Stage/cs (100% translated)")
- [русский](/wiki/Handbook:AMD64/Installation/Stage/ru "Handbook:AMD64/Installation/Stage/ru (100% translated)")
- [தமிழ்](/wiki/Handbook:AMD64/Installation/Stage/ta "Handbook:AMD64/Installation/Stage/ta (100% translated)")
- [中文（中国大陆）‎](/wiki/Handbook:AMD64/Installation/Stage/zh-cn "Handbook:AMD64/Installation/Stage/zh-cn (100% translated)")
- [日本語](/wiki/Handbook:AMD64/Installation/Stage/ja "Handbook:AMD64/Installation/Stage/ja (100% translated)")
- [한국어](/wiki/Handbook:AMD64/Installation/Stage/ko "Handbook:AMD64/Installation/Stage/ko (100% translated)")

[Sommaire du manuel](/wiki/Handbook:AMD64 "Handbook:AMD64")[Installation](/wiki/Handbook:AMD64/Full/Installation/fr "Handbook:AMD64/Full/Installation/fr")[À propos de l’installation](/wiki/Handbook:AMD64/Installation/About/fr "Handbook:AMD64/Installation/About/fr")[Choix du support](/wiki/Handbook:AMD64/Installation/Media/fr "Handbook:AMD64/Installation/Media/fr")[Configurer le réseau](/wiki/Handbook:AMD64/Installation/Networking/fr "Handbook:AMD64/Installation/Networking/fr")[Préparer les disques](/wiki/Handbook:AMD64/Installation/Disks/fr "Handbook:AMD64/Installation/Disks/fr")Installer l’archive _stage3_[Installer le système de base](/wiki/Handbook:AMD64/Installation/Base/fr "Handbook:AMD64/Installation/Base/fr")[Configurer le noyau](/wiki/Handbook:AMD64/Installation/Kernel/fr "Handbook:AMD64/Installation/Kernel/fr")[Configurer le système](/wiki/Handbook:AMD64/Installation/System/fr "Handbook:AMD64/Installation/System/fr")[Installer les outils](/wiki/Handbook:AMD64/Installation/Tools/fr "Handbook:AMD64/Installation/Tools/fr")[Configurer le système d’amorçage](/wiki/Handbook:AMD64/Installation/Bootloader/fr "Handbook:AMD64/Installation/Bootloader/fr")[Finaliser](/wiki/Handbook:AMD64/Installation/Finalizing/fr "Handbook:AMD64/Installation/Finalizing/fr")[Utiliser Gentoo](/wiki/Handbook:AMD64/Full/Working/fr "Handbook:AMD64/Full/Working/fr")[Introduction à Portage](/wiki/Handbook:AMD64/Working/Portage/fr "Handbook:AMD64/Working/Portage/fr")[Les options de la variable USE](/wiki/Handbook:AMD64/Working/USE/fr "Handbook:AMD64/Working/USE/fr")[Les fonctionnalités de Portage](/wiki/Handbook:AMD64/Working/Features/fr "Handbook:AMD64/Working/Features/fr")[Scripts d’initialisation systèmes](/wiki/Handbook:AMD64/Working/Initscripts/fr "Handbook:AMD64/Working/Initscripts/fr")[Variables d’environnement](/wiki/Handbook:AMD64/Working/EnvVar/fr "Handbook:AMD64/Working/EnvVar/fr")[Utiliser Portage](/wiki/Handbook:AMD64/Full/Portage/fr "Handbook:AMD64/Full/Portage/fr")[Fichiers et répertoires](/wiki/Handbook:AMD64/Portage/Files/fr "Handbook:AMD64/Portage/Files/fr")[Les variables](/wiki/Handbook:AMD64/Portage/Variables/fr "Handbook:AMD64/Portage/Variables/fr")[Mélanger plusieurs branches logicielles](/wiki/Handbook:AMD64/Portage/Branches/fr "Handbook:AMD64/Portage/Branches/fr")[Outils supplémentaires](/wiki/Handbook:AMD64/Portage/Tools/fr "Handbook:AMD64/Portage/Tools/fr")[Dépôt personnalisé](/wiki/Handbook:AMD64/Portage/CustomTree/fr "Handbook:AMD64/Portage/CustomTree/fr")[Fonctionnalités avancées](/wiki/Handbook:AMD64/Portage/Advanced/fr "Handbook:AMD64/Portage/Advanced/fr")[Configuration du réseau avec OpenRC](/wiki/Handbook:AMD64/Full/Networking/fr "Handbook:AMD64/Full/Networking/fr")[Bien démarrer](/wiki/Handbook:AMD64/Networking/Introduction/fr "Handbook:AMD64/Networking/Introduction/fr")[Configuration avancée](/wiki/Handbook:AMD64/Networking/Advanced/fr "Handbook:AMD64/Networking/Advanced/fr")[Les modules réseau](/wiki/Handbook:AMD64/Networking/Modular/fr "Handbook:AMD64/Networking/Modular/fr")[Sans fil](/wiki/Handbook:AMD64/Networking/Wireless/fr "Handbook:AMD64/Networking/Wireless/fr")[Ajouter des fonctionnalités](/wiki/Handbook:AMD64/Networking/Extending/fr "Handbook:AMD64/Networking/Extending/fr")[Gestion dynamique](/wiki/Handbook:AMD64/Networking/Dynamic/fr "Handbook:AMD64/Networking/Dynamic/fr")

## Contents

- [1Choix d’une archive _stage_](#Choix_d.E2.80.99une_archive_stage)
  - [1.1OpenRC](#OpenRC)
  - [1.2systemd](#systemd)
  - [1.3Multi-librairie (32 et 64 bits, _multilib_)](#Multi-librairie_.2832_et_64_bits.2C_multilib.29)
  - [1.4Non multi-librairie (64 bits pur)](#Non_multi-librairie_.2864_bits_pur.29)
- [2Téléchargement de l’archive _stage_](#T.C3.A9l.C3.A9chargement_de_l.E2.80.99archive_stage)
- [3Réglage de la date et de l’heure](#R.C3.A9glage_de_la_date_et_de_l.E2.80.99heure)
  - [3.1Automatiquement](#Automatiquement)
  - [3.2Manuellement](#Manuellement)
  - [3.3Navigateurs graphiques](#Navigateurs_graphiques)
  - [3.4Navigateurs en ligne de commande](#Navigateurs_en_ligne_de_commande)
  - [3.5Vérifier et valider](#V.C3.A9rifier_et_valider)
- [4Installation d’une archive _stage_](#Installation_d.E2.80.99une_archive_stage)
- [5Configuration des options de compilation](#Configuration_des_options_de_compilation)
  - [5.1Introduction](#Introduction)
  - [5.2CFLAGS et CXXFLAGS](#CFLAGS_et_CXXFLAGS)
  - [5.3RUSTFLAGS (drapeaux Rust)](#RUSTFLAGS_.28drapeaux_Rust.29)
  - [5.4MAKEOPTS (options _make_)](#MAKEOPTS_.28options_make.29)
  - [5.5À vos marques, prêts, partez !](#.C3.80_vos_marques.2C_pr.C3.AAts.2C_partez_.21)
- [6Références](#R.C3.A9f.C3.A9rences)

### Choix d’une archive _stage_

**Conseil**

Sur les architectures supportées, il est recommandé pour les utilisateurs visant un système d’exploitation du bureau (espace graphique) d’utiliser une archive avec le mot « `desktop` » dans le nom de fichier. Ces fichiers incluent des paquets comme [sys-devel/llvm](https://packages.gentoo.org/packages/sys-devel/llvm), [dev-lang/rust-bin](https://packages.gentoo.org/packages/dev-lang/rust-bin) et des drapeaux USE qui améliore le temps d’installation.

Le [fichier _stage_](/wiki/Stage_file/fr "Stage file/fr") fonctionne comme une graine d’amorce pour une installation Gentoo. Elles sont générés via [Catalyst](/wiki/Catalyst "Catalyst") par le [Release Engineering Team](/wiki/Project:RelEng "Project:RelEng"). Les fichiers _stage_ sont spécifiquesà un [profil](/wiki/Profile_(Portage) "Profile (Portage)") et contiennent un système presque complet.

Lors du choix d’un fichier _stage_, il est important de choisir un profil qui correspond au système final désiré.

**Important**

Bien qu’il soit possible de changer de profil après l’installation, changer nécessite des efforts importants et une attention qui sort du périmètre de ce manuel d’installation. Changer de système d’initialisation est difficile, changer de `no-multilib` à `multilib` nécessite une connaissance importante de Gentoo et de la chaîne de compilation ( _toolchain_).

**Conseil**

La plupart des utilisateurs ne devrait pas avoir besoin d’utiliser les archives tar « avancées » ; elles existent pour des configurations logicielles et matérielles atypiques.

#### OpenRC

[OpenRC](/wiki/OpenRC "OpenRC") est un système d’initialisation (responsable de démarrer les services systèmes une fois le noyau démarré) prenant en compte les dépendances et qui maintient la rétrocompatibilité avec les systèmes d’initialisation des applications normalement situé dans /sbin/init. C’est le système d’initialisation natif et créé de Gentoo, il est également utilisé par quelques autres distributions Linux et systèmes BSD.

#### systemd

systemd est un équivalent moderne de SysV _init_ et _rc_ pour les systèmes Linux. Il est majoritairement utilisé comme système d’initialisation par les distributions GNU/Linux. systemd est pleinement supporté par Gentoo et fonctionne comme prévu. Si quelque chose semble manquer dans le manuel, regardez [l’article systemd](/wiki/Systemd/fr "Systemd/fr") avant de demander du support.

#### Multi-librairie (32 et 64 bits, _multilib_)

**Remarque**

Toutes les architectures n’ont pas d’option multi-librairie ( _multilib_). La plupart fonctionne seulement avec le code natif. C’est principalement **amd64** qui est concerné.

Le profil _multilib_ utilise des bibliothèques 64 bits lorsque cela est possible et se replie seulement sur les versions 32 bits pour régler des problèmes de compatibilité. C’est une option excellente pour la majorité des installations car elle permet une grande flexibilité de personnalisation par le futur.

**Conseil**

Utiliser `multilib` rend plus facile un changement de profil ultérieur comparé à `no-multilib`

#### Non multi-librairie (64 bits pur)

**Attention !**

Les utilisateurs débutant avec Gentoo ne devraient _pas_ choisir une archive tar _no-multilib_ à moins que cela ne soit absolument nécessaire. La migration dun système _no-multilib_ vers un système _multilib_ nécessite une connaissance avancée du fonctionnement de Gentoo et de la chaîne de compilation (cela peut même faire frémir nos [développeurs Toolchain](/wiki/Project:Toolchain "Project:Toolchain")). Ce n’est pas pour les cœurs fragiles et cela dépasse largement la portée de ce guide.

Choisir une archive tar _no-multilib_ en tant que base du système fournit un environnement de système d’exploitation 64 bits complet – libre de tout programme 32 bits. Cela rend la capacité à passer vers des profils _multilib_ lourde, bien que techniquement toujours possible.

### Téléchargement de l’archive _stage_

Avant de télécharger le fichier _stage_, le répertoire courant devrait être celui du montage utilisé pour l’installation :

`root #` `cd /mnt/gentoo`

### Réglage de la date et de l’heure

Les archives _stage_ sont généralement téléchargées avec une connexion HTTPS qui nécessite une heure système juste. Un décalage peut empêcher un téléchargement de fonctionner, et si l’heure est ajustée après l’installation, cela peut créer des erreurs imprévisibles.

La date et l’heure actuelle peuvent être vérifiées avec date :

`root #` `date`

```
Mon Oct  3 13:16:22 PDT 2021

```

Si la date affichée est décalée de plus de quelques minutes, elle devrait être mise à jour en utilisant l’une des méthodes.

#### Automatiquement

Utiliser [NTP](/wiki/NTP "NTP") pour régler l’horloge est plus simple et plus fiable plutôt que de faire le réglage manuellement.

chronyd, de [net-misc/chrony](https://packages.gentoo.org/packages/net-misc/chrony), peut être utilisé pour ajuster l’horloge système en UTC (temps universel) avec :

`root #` `chronyd -q`

**Important**

Les systèmes sans un RTC ( _Real-Time Clock_) fonctionnel devront synchroniser l’horloge à chaque démarrage, et à intervalles réguliers. C’est aussi utile pour des systèmes avec RTC, car la batterie peut être défaillante et un décalage d’horloge peut s’accumuler.

**Attention !**

Le trafic standard NTP n’est pas sécurisé, il est important de vérifier les données horaires obtenues depuis le réseau.

#### Manuellement

Lorsqu’un accès NTP n’est pas possible, date peut être utilisé pour configurer manuellement l’horloge système.

**Remarque**

L’heure UTC est recommandée pour tous les systèmes Linux. Plus tard, un fuseau horaire sera défini ; cela décalera l’heure affichée.

Le format des arguments pour paramétrer l’horloge est : `MMJJhhmmAAAA` ( **M** ois, **J** our, **h** eure, **m** inute et **A** nnée).

Par exemple, pour régler la date au 3 octobre 2021 à 13:16 :

`root #` `date 100313162021`

#### Navigateurs graphiques

Ceux utilisant un environnement avec des navigateurs Internet graphiques n’auront aucun problème à copier l’adresse d’une archive tar depuis la [section téléchargements](https://www.gentoo.org/downloads/#other-arches) du site principal. Sélectionnez simplement l’onglet approprié, cliquez droit sur le lien de l’archive _stage_, ensuite Copier l’adresse du lien pour copier le lien vers le presse-papiers, puis collez le lien à l’utilitaire wget en ligne de commande pour télécharger l’archive _stage_ :

`root #` `wget <URL_DE_L_ARCHIVE_COLLÉE>`

#### Navigateurs en ligne de commande

Les lecteurs plus traditionnels ou utilisateurs de Gentoo 'vieux jeu', travaillant exclusivement depuis la ligne de commande peuvent préférer l’utilisation de links ([www-client/links](https://packages.gentoo.org/packages/www-client/links)), un navigateur non-graphique et utilisable avec des menus. Pour télécharger une archive _stage_, naviguez vers la liste des miroirs Gentoo comme suit :

`root #` `links https://www.gentoo.org/downloads/mirrors/`

Pour utiliser un proxy HTTP avec links, passez l’URL avec l’option `http-proxy` :

`root #` `links -http-proxy proxy.server.com:8080 https://www.gentoo.org/downloads/mirrors/`

Outre links, il y a également le navigateur lynx ([www-client/lynx](https://packages.gentoo.org/packages/www-client/lynx)). Comme links c’est un navigateur non-graphique mais celui-là n’est pas utilisable avec des menus.

`root #` `lynx https://www.gentoo.org/downloads/mirrors/`

Si un proxy est nécessaire, exportez les variables http\_proxy et/ou ftp\_proxy :

`root #` `export http_proxy="http://proxy.server.com:port"
`

`root #` `export ftp_proxy="http://proxy.server.com:port"`

Sur la liste des miroirs, choisissez-en un à proximité. En général les miroirs HTTP suffisent, mais d’autres protocoles sont également disponibles. Naviguez vers le répertoire releases/amd64/autobuilds/. Ici, toutes les archives _stage_ disponibles sont affichées (elles peuvent être stockées dans des sous-répertoires nommés après les différents types d’architectures). Sélectionnez-en une et appuyez sur la touche `d` pour la télécharger.

Une fois le téléchargement de l'archive terminé, il es possible d’en vérifier l’intégrité et d’en valider son contenu. Les intéressés peuvent passer à la [section suivante](/wiki/Handbook:AMD64/Installation/Stage/fr#Verifying_and_validating "Handbook:AMD64/Installation/Stage/fr").

Ceux qui ne sont pas intéressés peuvent fermer le navigateur en ligne de commande en appuyant sur la touche `q` et peuvent aller directement à la section [Extraction de l’archive tar](/wiki/Handbook:AMD64/Installation/Stage/fr#Unpacking_the_stage_file "Handbook:AMD64/Installation/Stage/fr").

#### Vérifier et valider

Comme pour les CD d’installation, il est possible de vérifier et de valider l’archive _stage_ téléchargée. Bien que ces étapes peuvent être sautées, ces fichiers sont proposés pour les utilisateurs qui se soucient de l’intégrité des fichiers qu’ils viennent de télécharger. Les fichiers supplémentaires sont disponibles dans le répertoire racine du miroir. Naviguez dans le répertoire approprié à l’architecture et le profil, puis téléchargez les fichiers .CONTENTS.gz, .DIGESTS et .sha256 associés.

`root #` `wget https://distfiles.gentoo.org/releases/`

- .CONTENTS.gz contient la liste de tous les fichiers contenus dans l’archive _stage_.
- .DIGESTS contient les sommes de contrôle de l’archive _stage_ dans plusieurs algorithmes cryptographiques différents.
- .sha256 contient une signature SHA256 (aussi appelée somme de contrôle ou _checksum_) signée par PGP de l’archive _stage_. Ce fichier n’est pas toujours disponible pour toutes les archives _stage_.

Tout comme pour le fichier ISO, il est également possible de vérifier la signature cryptographique du fichier tar.xz en utilisant gpg afin de s’assurer que l’archive tar n’a pas été altérée.

Pour les images officielles Gentoo, le paquet [sec-keys/openpgp-keys-gentoo-release](https://packages.gentoo.org/packages/sec-keys/openpgp-keys-gentoo-release) fournit les clés PGP des sorties automatiques. Ces clés doit d’abord être importées dans la session de l’utilisateur de manière à pouvoir être utilisée pour les vérifications :

`root #` `gpg --import /usr/share/openpgp-keys/gentoo-release.asc`

Pour les images non officielles qui comprennent les outils gpg et wget, un lot contenant les clés de Gentoo peut être téléchargé et importé :

Pour vérifier la signature de l’archive tar et, éventuellement, les fichiers de somme de contrôle associés :

`root #` `gpg --verify stage3-amd64-<release>-<init>.tar.xz.asc
`

`root #` `gpg --verify stage3-amd64-<release>-<init>.tar.xz.DIGESTS
`

`root #` `gpg --verify stage3-amd64-<release>-<init>.tar.xz.sha256
`

Si la vérification réussit, « _Good signature from_ » sera affiché en sortie de la commande précédente.

Les empreintes des clés OpenPGP utilisés pour les signatures des fichiers peuvent être trouvée sur la [page des signature de média](https://www.gentoo.org/downloads/signatures/).

**Remarque**

La majorité des archives _stage_ ont maintenant un [suffixe](https://www.gentoo.org/news/2021/07/20/more-downloads.html) avec le type de système d’initialisation (openrc ou systemd) ; bien que certaines architectures ne l’ont pas encore pour le moment.

Un outil cryptographique comme openssl, sha256 ou sha512 peut être utilisé pour comparer les signatures avec celle indiquée dans le fichier .DIGESTS.

Pour vérifier la somme de contrôle SHA512 avec openssl :

`root #` `openssl dgst -r -sha512 stage3-amd64-<release>-<init>.tar.xz`

`dgst` indique à la commande openssl d’utiliser la sous-commande _Message Digest_, `-r` affiche le résultat selon le format _coreutils_ et `-sha512` choisi l’algorythme SHA526.

Pour vérifier la somme de contrôle BLAKE2B512 avec openssl :

`root #` `openssl dgst -r -blake2b512 stage3-amd64-<release>-<init>.tar.xz`

Comparez le résultat de ces commandes avec les signatures et les noms de fichiers contenus dans le .DIGESTS. Les paires de valeurs doivent être identiques à la sortie de la commande, sinon le fichier téléchargé est corrompu et doit être détruit et retéléchargé.

Pour vérifier une somme de contrôle SHA256 avec un fichier .sha256 associé, utilisez la commande sha256sum :

`root #` `sha256sum --check stage3-amd64-<release>-<init>.tar.xz.sha256`

L’option `--check` indique à sha256sum de lire une liste de fichiers et leur signature attendue, et d’afficher « OK » lorsque la somme de contrôle correspond ou « FAILED » dans le cas contraire.

## Installation d’une archive _stage_

Une fois le fichier _stage_ téléchargé et vérifié, il peut être extrait avec tar :

`root #` `tar xpvf stage3-*.tar.xz --xattrs-include='*.*' --numeric-owner -C /mnt/gentoo`

Avant d’extraire, vérifier les options :

- `x` e **x** traire, indique que tar va extraire le contenu de l’archive ;
- `p` **p** réserver les permissions ;
- `v` sortie **v** erbeuse ;
- `f` **f** ichier, indique à tar le nom du fichier d’entrée ;
- `--xattrs-include='*.*'` permet de conserver les attributs étendus contenus dans tous les espaces de noms de l’archive ;
- `--numeric-owner` assure que les identifiants de groupe et d’utilisateur des fichiers extraits depuis l’archive tar restent les mêmes que ceux voulus par l’équipe de Gentoo (même si certains utilisateurs aventureux n’utilisent pas les environnements Gentoo officiels) ;
- `-C /mnt/gentoo` extrait les fichiers dans la partition racine, peu importe le répertoire actuel.

Maintenant que l’archive est extraite, continuez avec la [Configuration des options de compilation](/wiki/Handbook:AMD64/Installation/Stage/fr#Configuring_compile_options "Handbook:AMD64/Installation/Stage/fr").

## Configuration des options de compilation

### Introduction

Pour optimiser le système, il est possible de configurer des variables qui influent sur le comportement de Portage, le gestionnaire de paquets officiel de Gentoo. Toutes ces variables peuvent être configurées en tant que variable d'environnement (en utilisant export), mais cela n’est pas permanent.

**Remarque**

Techniquement, les variables peuvent être exportées via le fichier profile ou rc de l’ [interpréteur de commandes ( _shell_)](/wiki/Shell "Shell"), toutefois, ce n’est pas une bonne pratique pour l’administration basique du système.

Portage lit le fichier [make.conf](/wiki/Make.conf "Make.conf") lorsqu’il s’exécute, ce qui permet de changer son comportement en fonction des valeurs sauvegardées dans ce fichier. make.conf peut être vu comme le fichier principal de configuration pour Portage, donc son contenu doit être rempli avec attention.

{{Tip\|Une liste commentée de toutes les variables possibles se trouve dans /mnt/gentoo/usr/share/portage/config/make.conf.example. De la documentation supplémentaire concernant make.conf est accessible en exécutant man 5 make.conf.

Pour réussir une installation de Gentoo, seules les variables mentionnées ci-dessous doivent être paramétrées :

Lancez un éditeur (dans ce guide nous utiliserons nano) pour modifier les variables d’optimisation décrites ci-dessous.

`root #` `nano /mnt/gentoo/etc/portage/make.conf`

En regardant dans le fichier make.conf.example, la manière dans laquelle le fichier doit être structuré est évidente : les lignes commentées démarrent par `#`, les autres lignes définissent des variables en utilisant la syntaxe `VARIABLE="valeur"`. Plusieurs de ces variables sont présentées dans la section suivante.

### CFLAGS et CXXFLAGS

Les variables CFLAGS et CXXFLAGS définissent les paramètres d’optimisation des compilateurs GCC C et C++, respectivement. Bien que ces variables soient généralement définies ici, il est possible, pour une performance maximale, d’optimiser ces paramètres pour chaque programme séparément. La raison pour cela est que chaque programme est différent. Cependant, ceci est difficilement gérable, d’où la définition de ces paramètres communs dans le fichier make.conf.

Dans make.conf il faut définir les paramètres d’optimisation qui rendront le système le plus réactif en général. Ne pas utiliser de configuration expérimentale dans cette variable ; trop d’optimisation peut faire que les programmes se comportent mal (plantage ou pire, un mauvais fonctionnement).

Ce manuel n’expliquera pas toutes les options d’optimisation possibles. Pour les comprendre toutes, lire le [manuel en ligne de GCC](https://gcc.gnu.org/onlinedocs/) (en anglais) ou la page d’infos de gcc (info gcc). Le fichier make.conf.example contient également de lui-même beaucoup d’exemples et d’informations ; ne pas oublier de le lire également.

Un première configuration est le paramètre `-march=` ou `-mtune=`, qui spécifie le nom de l’architecture cible. Les options possibles sont décrites dans le fichier make.conf.example (en commentaire). Une valeur souvent utilisée est « _native_ », qui informe au compilateur de sélectionner l’architecture cible du système utilisé (celui sur lequel est installé Gentoo).

Un second paramètre est `-O` (un O majuscule et non un zéro), qui permet de spécifier la classe des paramètres d’optimisations de GCC. Les classes disponibles sont « s » (optimisé pour la taille), « 0 » (zéro, pour pas d’optimisations), « 1 », « 2 » ou même « 3 » pour plus d’optimisations de vitesse (chaque classe à les mêmes paramètres que la précédente plus quelques extras). `-O2` est la valeur recommandée. `-O3` est connu pour causer des problèmes quand utilisé pour tout le système, nous recommandons donc de rester avec `-O2`.

Un autre paramètre d’optimisation populaire est `-pipe` (qui permet l’utilisation de l’opérateur de transfert de données, _pipe_, à la place de fichiers temporaires pour la communication entre les différentes étapes de la compilation). Ce n’a aucun impact sur le code généré, mais utilise plus de mémoire. Sur des systèmes disposant de peu de mémoire vive, gcc peut être tué. Dans ce cas, ne pas utiliser ce paramètre.

Utiliser `-fomit-frame-pointer` (qui ne garde pas la structure des pointeurs dans un registre pour les fonctions qui n’en ont pas besoin) peut avoir des répercussions importantes sur le débogage des programmes.

Quand les variables CFLAGS et CXXFLAGS sont définies, combinez les paramètres d’optimisation multiples dans une seule chaîne de caractères. Les valeurs par défaut contenues dans l’archive _stage3_ qui est extraite devraient être suffisantes. Les valeurs suivantes ne sont qu'un exemple :

CODE **Exemple des variables CFLAGS et CXXFLAGS**

```
# Options de compilation pour tous les langages
COMMON_FLAGS="-march=native -O2 -pipe"
# Utiliser les mêmes paramètres pour les deux variables
CFLAGS="${COMMON_FLAGS}"
CXXFLAGS="${COMMON_FLAGS}"

```

**Conseil**

Bien que l’article [d’optimisation de GCC](/wiki/GCC_optimization/fr "GCC optimization/fr") possède plus d’informations sur comment les différentes options de compilation affectent un système, l’article [CFLAGS sûrs](/wiki/Safe_CFLAGS "Safe CFLAGS") peut se révéler plus pratique pour permettre aux débutants d’optimiser leur système.

### RUSTFLAGS (drapeaux Rust)

Beaucoup d’applications sont maintenant écrites en Rust, lequel possède ses propres manières d’optimiser. Par défaut, Rust a un niveau d’optimisation réglé sur 3, sauf si un projet change cette valeur. Donc, elle devrait être laissée ainsi. Une liste complète des optimisations (connue comme _codegen_) passables au compilateur Rust est disponible sur [https://doc.rust-lang.org/rustc/codegen-options/index.html](https://doc.rust-lang.org/rustc/codegen-options/index.html).

L’optimisation la plus utile serait d’indiquer à Rust de compiler pour votre processeur en utilisant par exemple :

FILE **`/etc/portage/make.conf`** **Exemple de RUSTFLAGS**

```
RUSTFLAGS="${RUSTFLAGS} -C target-cpu=native"

```

Pour obtenie la liste des processeurs supportés par Rust, lancez :

`root #` `rustc -C target-cpu=help`

Ceci va afficher la valeur par défaut et le type de processeur sélectionné avec « _native_ ».

**Remarque**

La commande ci-dessus ne fonctionne que depuis une archive _stage3_ ou après avoir installé [dev-lang/rust-bin](https://packages.gentoo.org/packages/dev-lang/rust-bin) ou [dev-lang/rust](https://packages.gentoo.org/packages/dev-lang/rust).

### MAKEOPTS (options _make_)

La variable MAKEOPTS définit combien de compilations parallèles peuvent se dérouler lors de l’installation d’un paquet. Depuis Portage 3.0.31 [\[1\]](#cite_note-1), si la valeur n’est pas définie, le comportement par défaut est que MAKEOPTS lancera le même nombre de processus que de processeurs retournés par nproc.

Également depuis Portage 3.0.53[\[2\]](#cite_note-2), si non défini, le comportement par défaut de Portage est de limiter la charge système ( _load average_)au nombre de processeurs retournés par nproc.

Un bon choix est la plus petite valeur entre : le nombre de processeurs ou la quantité de mémoire vive divisée par 2 Gio.

**Attention !**

Paramétrer un nombre important de processus peut augmenter de manière conséquente la consommation de mémoire. Une recommandation est d’avoir au moins 2 Gio de mémoire vive pour chaque processus (donc `-j6` nécessite au moins 12 Gio). Pour éviter d’être à court de mémoire, il faut diminuer le nombre de processus.

**Conseil**

Lorsque vous utilisez des compilations en parallèle ( `--jobs`), le nombre de processus peut augmenter exponentiellement (nombre de processus × nombre de compilations). Ceci peut être corrigé en utilisant distcc localement qui va limiter le nombre de compilations par hôte.

FILE **`/etc/portage/make.conf`** **Exemple de déclaration de MAKEOPTS dans make.conf**

```
# Si laissé indéfini, le comportement par défaut de Portage est :
# - définir MAKEOPTS pour autant de processus que retourné par `nproc`
# - définir MAKEOPTS pour la charge système limite correspond au nombre de processeurs retournés par `nproc`, approximativement car cette valeur est une moyenne
# Remplacez « 4 » par la valeur obtenue par minimum(RAM/2 Gio, nombre de processeurs) ou ne paramétrez rien.
MAKEOPTS="-j4 -l5"

```

Recherchez MAKEOPTS dans man 5 make.conf pour plus de détails.

### À vos marques, prêts, partez !

Mettez à jour le fichier /mnt/gentoo/etc/portage/make.conf en fonction de vos préférences personnelles et enregistrez-le (les utilisateurs de nano appuieront sur `Ctrl` + `o` pour écrire les changements et `Ctrl` + `x` pour quitter).

## Références

1. [↑](#cite_ref-1)[https://gitweb.gentoo.org/proj/portage.git/commit/?id=5d2af567772bb12b073f1671daea6263055cbdc2](https://gitweb.gentoo.org/proj/portage.git/commit/?id=5d2af567772bb12b073f1671daea6263055cbdc2)
2. [↑](#cite_ref-2)[https://gitweb.gentoo.org/proj/portage.git/commit/?id=de7be7f45ee54e3f789def46542919550687d15e](https://gitweb.gentoo.org/proj/portage.git/commit/?id=de7be7f45ee54e3f789def46542919550687d15e)

[← Préparer les disques](/wiki/Handbook:AMD64/Installation/Disks/fr "Previous part") [Home](/wiki/Handbook:AMD64 "Handbook:AMD64") [Installation du système de base →](/wiki/Handbook:AMD64/Installation/Base/fr "Next part")