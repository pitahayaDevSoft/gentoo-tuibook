# Media

Other languages:

- [Deutsch](/wiki/Handbook:PPC/Installation/Media/de "Handbuch:PPC/Installation/Medium (100% translated)")
- [English](/wiki/Handbook:PPC/Installation/Media "Handbook:PPC/Installation/Media (100% translated)")
- [español](/wiki/Handbook:PPC/Installation/Media/es "Manual de Gentoo: PPC/Instalación/Medio (100% translated)")
- français
- [italiano](/wiki/Handbook:PPC/Installation/Media/it "Handbook:PPC/Installation/Media (100% translated)")
- [magyar](/wiki/Handbook:PPC/Installation/Media/hu "Handbook:PPC/Installation/Media/hu (100% translated)")
- [polski](/wiki/Handbook:PPC/Installation/Media/pl "Handbook:PPC/Installation/Media (100% translated)")
- [português do Brasil](/wiki/Handbook:PPC/Installation/Media/pt-br "Manual:PPC/Instalação/Midia (100% translated)")
- [русский](/wiki/Handbook:PPC/Installation/Media/ru "Handbook:PPC/Installation/Media (100% translated)")
- [தமிழ்](/wiki/Handbook:PPC/Installation/Media/ta "கையேடு:PPC/நிறுவல்/ஊடகம் (100% translated)")
- [中文（中国大陆）‎](/wiki/Handbook:PPC/Installation/Media/zh-cn "手册：PPC/安装/选择安装媒介 (100% translated)")
- [日本語](/wiki/Handbook:PPC/Installation/Media/ja "ハンドブック:PPC/インストール/メディア (100% translated)")
- [한국어](/wiki/Handbook:PPC/Installation/Media/ko "Handbook:PPC/Installation/Media/ko (100% translated)")

[Sommaire du manuel](/wiki/Handbook:PPC "Handbook:PPC")[Installation](/wiki/Handbook:PPC/Full/Installation/fr "Handbook:PPC/Full/Installation/fr")[À propos de l’installation](/wiki/Handbook:PPC/Installation/About/fr "Handbook:PPC/Installation/About/fr")Choix du support[Configurer le réseau](/wiki/Handbook:PPC/Installation/Networking/fr "Handbook:PPC/Installation/Networking/fr")[Préparer les disques](/wiki/Handbook:PPC/Installation/Disks/fr "Handbook:PPC/Installation/Disks/fr")[Installer l’archive _stage3_](/wiki/Handbook:PPC/Installation/Stage/fr "Handbook:PPC/Installation/Stage/fr")[Installer le système de base](/wiki/Handbook:PPC/Installation/Base/fr "Handbook:PPC/Installation/Base/fr")[Configurer le noyau](/wiki/Handbook:PPC/Installation/Kernel/fr "Handbook:PPC/Installation/Kernel/fr")[Configurer le système](/wiki/Handbook:PPC/Installation/System/fr "Handbook:PPC/Installation/System/fr")[Installer les outils](/wiki/Handbook:PPC/Installation/Tools/fr "Handbook:PPC/Installation/Tools/fr")[Configurer le système d’amorçage](/wiki/Handbook:PPC/Installation/Bootloader/fr "Handbook:PPC/Installation/Bootloader/fr")[Finaliser](/wiki/Handbook:PPC/Installation/Finalizing/fr "Handbook:PPC/Installation/Finalizing/fr")[Utiliser Gentoo](/wiki/Handbook:PPC/Full/Working/fr "Handbook:PPC/Full/Working/fr")[Introduction à Portage](/wiki/Handbook:PPC/Working/Portage/fr "Handbook:PPC/Working/Portage/fr")[Les options de la variable USE](/wiki/Handbook:PPC/Working/USE/fr "Handbook:PPC/Working/USE/fr")[Les fonctionnalités de Portage](/wiki/Handbook:PPC/Working/Features/fr "Handbook:PPC/Working/Features/fr")[Scripts d’initialisation systèmes](/wiki/Handbook:PPC/Working/Initscripts/fr "Handbook:PPC/Working/Initscripts/fr")[Variables d’environnement](/wiki/Handbook:PPC/Working/EnvVar/fr "Handbook:PPC/Working/EnvVar/fr")[Utiliser Portage](/wiki/Handbook:PPC/Full/Portage/fr "Handbook:PPC/Full/Portage/fr")[Fichiers et répertoires](/wiki/Handbook:PPC/Portage/Files/fr "Handbook:PPC/Portage/Files/fr")[Les variables](/wiki/Handbook:PPC/Portage/Variables/fr "Handbook:PPC/Portage/Variables/fr")[Mélanger plusieurs branches logicielles](/wiki/Handbook:PPC/Portage/Branches/fr "Handbook:PPC/Portage/Branches/fr")[Outils supplémentaires](/wiki/Handbook:PPC/Portage/Tools/fr "Handbook:PPC/Portage/Tools/fr")[Dépôt personnalisé](/wiki/Handbook:PPC/Portage/CustomTree/fr "Handbook:PPC/Portage/CustomTree/fr")[Fonctionnalités avancées](/wiki/Handbook:PPC/Portage/Advanced/fr "Handbook:PPC/Portage/Advanced/fr")[Configuration du réseau avec OpenRC](/wiki/Handbook:PPC/Full/Networking/fr "Handbook:PPC/Full/Networking/fr")[Bien démarrer](/wiki/Handbook:PPC/Networking/Introduction/fr "Handbook:PPC/Networking/Introduction/fr")[Configuration avancée](/wiki/Handbook:PPC/Networking/Advanced/fr "Handbook:PPC/Networking/Advanced/fr")[Les modules réseau](/wiki/Handbook:PPC/Networking/Modular/fr "Handbook:PPC/Networking/Modular/fr")[Sans fil](/wiki/Handbook:PPC/Networking/Wireless/fr "Handbook:PPC/Networking/Wireless/fr")[Ajouter des fonctionnalités](/wiki/Handbook:PPC/Networking/Extending/fr "Handbook:PPC/Networking/Extending/fr")[Gestion dynamique](/wiki/Handbook:PPC/Networking/Dynamic/fr "Handbook:PPC/Networking/Dynamic/fr")

## Contents

- [1Pré-requis matériels](#Pr.C3.A9-requis_mat.C3.A9riels)
- [2Support d’installation de Gentoo Linux](#Support_d.E2.80.99installation_de_Gentoo_Linux)
  - [2.1CD minimal d'installation](#CD_minimal_d.27installation)
  - [2.2Qu’est ce qu'un fichier stage ?](#Qu.E2.80.99est_ce_qu.27un_fichier_stage_.3F)
- [3Téléchargement](#T.C3.A9l.C3.A9chargement)
  - [3.1Obtenir le support d’installation](#Obtenir_le_support_d.E2.80.99installation)
    - [3.1.1Naviguer dans les miroirs Gentoo](#Naviguer_dans_les_miroirs_Gentoo)
  - [3.2Vérifier les fichiers téléchargés](#V.C3.A9rifier_les_fichiers_t.C3.A9l.C3.A9charg.C3.A9s)
    - [3.2.1Vérifications depuis un système Microsoft Windows](#V.C3.A9rifications_depuis_un_syst.C3.A8me_Microsoft_Windows)
    - [3.2.2Vérifications depuis un système Linux](#V.C3.A9rifications_depuis_un_syst.C3.A8me_Linux)
- [4Créer le support d’amorçage](#Cr.C3.A9er_le_support_d.E2.80.99amor.C3.A7age)
  - [4.1Créer une clé USB amorçable](#Cr.C3.A9er_une_cl.C3.A9_USB_amor.C3.A7able)
    - [4.1.1Créer depuis Linux](#Cr.C3.A9er_depuis_Linux)
      - [4.1.1.1Déterminer le chemin d’accès du périphérique USB](#D.C3.A9terminer_le_chemin_d.E2.80.99acc.C3.A8s_du_p.C3.A9riph.C3.A9rique_USB)
      - [4.1.1.2Créer avec dd](#Cr.C3.A9er_avec_dd)
    - [4.1.2Écrire avec Windows](#.C3.89crire_avec_Windows)
  - [4.2Graver un disque](#Graver_un_disque)
    - [4.2.1Gravure depuis Microsoft Windows 7 et supérieur](#Gravure_depuis_Microsoft_Windows_7_et_sup.C3.A9rieur)
    - [4.2.2Gravure sous Linux](#Gravure_sous_Linux)
- [5Démarrage](#D.C3.A9marrage)
  - [5.1LiveGUI root access](#LiveGUI_root_access)
  - [5.2Configuration matérielle supplémentaire](#Configuration_mat.C3.A9rielle_suppl.C3.A9mentaire)
  - [5.3Facultatif : comptes utilisateurs](#Facultatif_:_comptes_utilisateurs)
  - [5.4Facultatif : consulter la documentation pendant l’installation](#Facultatif_:_consulter_la_documentation_pendant_l.E2.80.99installation)
    - [5.4.1TTY](#TTY)
    - [5.4.2GNU Screen](#GNU_Screen)
  - [5.5Facultatif : démarrer le daemon SSH](#Facultatif_:_d.C3.A9marrer_le_daemon_SSH)

## Pré-requis matériels

Gentoo peut être installé dans plein de systèmes. Il est peu utile d’écrire une liste de spécification matérielle minimale, c’est plus une question de combien de temps l’utilisateur est-il prêt à attendre. N’importe quelle machine raisonablement moderne peut être installée en moins d’une heure, cependant ce n’est pas une course donc prenez tout le temps donc vous avez besoin.

Pour une installation, le manuel recommande au moins 40 Gio sur la partition racine.

## Support d’installation de Gentoo Linux

**Conseil**

Bien qu’il soit recommandé d’utiliser un support Gentoo officiel pour l’installation, il est possible de partir d’un autre environnement d’installation. Toutefois, il n’y a aucune garantie que les composants nécessaires sont inclus. Si un environnement alternatif est utilisé, passez directement à [Préparer les disques](/wiki/Handbook:PPC/Installation/Disks/fr "Handbook:PPC/Installation/Disks/fr").

### CD minimal d'installation

Le _CD minimal d’installation de Gentoo_, aussi appelé installcd , est un environment Gentoo minimal autonome amorçable (bootable). L’image est maintenue par les [développeurs de Gentoo](/wiki/Project:RelEng "Project:RelEng") et permet d’installer Gentoo avec une connexion Internet. Pendant le démarrage, le matériel est détecté et les pilotes appropriés sont automatiquement chargés.

Le fichier du CD minimal d’installation est nommé install-<arch>-minimal-<release heurodatage>.iso.

**Important**

Le CD minimal d’installation nécessite 140 Mio de RAM pour démarrer.

### Qu’est ce qu'un fichier stage ?

Un [Fichier Stage](/wiki/Stage_file/fr "Stage file/fr") est un archive qui sert à initialiser un environnement Gentoo.

L’archive _stage3_ peut être téléchargée depuis releases/ppc/autobuilds/ sur l’un des [miroirs Gentoo officiels](https://www.gentoo.org/downloads/mirrors/). Ces archives sont mises à jour fréquemment et ne sont pas incluses dans les images d’installation officielles.

**Conseil**

Pour le moment, les _fichiers stage_ peuvent être ignorés. Ils seront décrits plus longuement en détail le moment venu.

**Remarque**

Historiquement, le manuel décrivait l’installation pour des [_fichiers stage_](/wiki/Stage_file/fr "Stage file/fr") inférieurs à 3. Ces archives contiennent un environnement inadapté pour une installation classique et ne sont plus couvertes dans le manuel.

## Téléchargement

### Obtenir le support d’installation

Le support d’installation par défaut que Gentoo Linux utilise est le _CD minimal d’installation_ qui fournit un petit environnement Gentoo Linux amorçable. Cet environnement comprend les outils nécessaires pour installer Gentoo. Les images CD peuvent être téléchargées depuis la [page de téléchargements](https://www.gentoo.org/downloads/) (recommandé) ou depuis l'un des nombreux [miroirs disponibles](https://www.gentoo.org/downloads/mirrors/).

#### Naviguer dans les miroirs Gentoo

Si le téléchargement s’effectue depuis un miroir, le CD minimal d’installation peut se trouver comme suit :

1. Se connecter sur un miroir, typiquement en utilisant un site proche trouvé sur [la liste des miroirs](https://www.gentoo.org/downloads/mirrors/).
2. Naviguez vers le répertoire releases/.
3. Sélectionner le répertoire correspondant à l’architecture souhaitée (tel que **ppc/**).
4. Sélectionnez le répertoire autobuilds/.
5. Pour les architectures **amd64** et **x86**, sélectionnez soit le répertoire current-install-amd64-minimal/ ou current-install-x86-minimal/ (respectivement). Pour toutes les autres architectures, naviguez vers le répertoire current-iso/.

**Remarque**

Certaines architectures comme **arm**, **mips**, et **s390** n’ont pas de CD minimal d’installation. Pour l’instant, le [Gentoo Release Engineering project](/wiki/Project:RelEng "Project:RelEng") ne propose pas d’image .iso pour ces architectures.

Dans cet emplacement, le fichier du support d’installation est le fichier avec l’extension .iso. Prendre pour exemple la liste suivante :

CODE **Liste d'exemple des fichiers téléchargeables dans releases/ppc/autobuilds/current-iso/**

```
[TXT]	install-ppc-minimal-20231112T170154Z.iso.asc	        2023-11-12 20:41        488
[TXT]	install-ppc-minimal-20231119T164701Z.iso.asc	        2023-11-19 18:41        488
[TXT]	install-ppc-minimal-20231126T163200Z.iso.asc	        2023-11-26 18:41        488
[TXT]	install-ppc-minimal-20231203T170204Z.iso.asc	        2023-12-03 18:41        488
[TXT]	install-ppc-minimal-20231210T170356Z.iso.asc	        2023-12-10 19:01        488
[TXT]	install-ppc-minimal-20231217T170203Z.iso.asc	        2023-12-17 20:01        488
[TXT]	install-ppc-minimal-20231224T164659Z.iso.asc	        2023-12-24 20:41        488
[TXT]	install-ppc-minimal-20231231T163203Z.iso.asc	        2023-12-31 19:01        488
[ ]     install-ppc-minimal-20240107T170309Z.iso              2024-01-07 20:42        466M
[ ]     install-ppc-minimal-20240107T170309Z.iso.CONTENTS.gz	2024-01-07 20:42        9.8K
[ ]     install-ppc-minimal-20240107T170309Z.iso.DIGESTS      2024-01-07 21:01        1.3K
[TXT]   install-ppc-minimal-20240107T170309Z.iso.asc	        2024-01-07 21:01        488
[ ]     install-ppc-minimal-20240107T170309Z.iso.sha256       2024-01-07 21:01        660
[TXT]	latest-install-ppc-minimal.txt                        2024-01-08 02:01        653

```

Dans l’exemple ci-dessus, le fichier install-ppc-minimal-20240107T170309Z.iso correspond au CD d’installation. Il existe cependant d’autres fichiers qui lui sont relatifs :

- un fichier .CONTENTS.gz listant tous les fichiers existants dans le support d’installation ; cela permet de vérifier si un micrologiciel ou un pilote en particulier est disponible avant le téléchargement ;
- un fichier .DIGESTS qui contient le hachage du fichier ISO avec différents algorithmes et formats ; ce fichier peut permettre de savoir si l’ISO téléchargée est corrompue ou non ;
- un fichier .asc qui est une signature cryptographique du fichier ISO ; il peut être utilisé pour vérifier l’intégrité et l’authenticité d’une image ; que le fichier téléchargé est effectivement approuvé par l’équipe [Gentoo Release Engineering team](/wiki/Project:RelEng "Project:RelEng").

Pour le moment, ne vous se souciez pas des autres fichiers – ils entreront en scène plus loin dans le processus d’installation. Téléchargez le fichier .iso et, s’il est souhaité d’en vérifier l'intégrité, téléchargez aussi le fichier .iso.asc.

**Conseil**

Le fichier .DIGESTS est seulement nécessaire si le fichier .iso.asc n’est pas vérifié.

### Vérifier les fichiers téléchargés

**Remarque**

Il s’agit là d’une opération facultative qui n’est pas nécessaire à l’installation de Gentoo Linux, mais elle est néanmoins recommandée pour être certain que le fichier n’est pas corrompu et, a bien été fourni par [Gentoo Infrastructure team](/wiki/Project:Infrastructure "Project:Infrastructure").

Le fichier .asc contient une signature électronique de l’ISO. En la vérifiant, vous pouvez vous assurer que le fichier provient bien de l’équipe Gentoo Release Engineering team, et est intact et non altéré.

#### Vérifications depuis un système Microsoft Windows

Pour vérifier la signature cryptographique, des outils tels que [GPG4Win](http://www.gpg4win.org/) peuvent être utilisés. Après l’installation, les clés publiques de la Gentoo Release Engineering team doivent être importées. La liste des clés est disponible sur la [page des signatures](https://www.gentoo.org/downloads/signatures/). Une fois l’importation terminée, l’utilisateur peut vérifier la signature contenue dans le fichier .DIGEST.asc.

#### Vérifications depuis un système Linux

Sur un système Linux, la méthode la plus courante pour vérifier une signature cryptographique consiste à utiliser le logiciel [app-crypt/gnupg](https://packages.gentoo.org/packages/app-crypt/gnupg). Quand ce paquet est installé, les commandes suivantes peuvent être utilisées pour vérifier la signature contenue dans le fichier .DIGESTS.asc.

**Conseil**

Lorsque vous importez les clés de Gentoo, vérifiez que les 16 caractères de la _key ID_ soient `BB572E0E2D182910`

Les clés de Gentoo peuvent être téléchargées depuis hkps://keys.gentoo.org en utilisant les empreintes disponibles sur la [page des signatures](https://www.gentoo.org/downloads/signatures/) :

`user $` `gpg --keyserver hkps://keys.gentoo.org --recv-keys 13EBBDBEDE7A12775DFDB1BABB572E0E2D182910`

```
gpg: directory '/root/.gnupg' created
gpg: keybox '/root/.gnupg/pubring.kbx' created
gpg: /root/.gnupg/trustdb.gpg: trustdb created
gpg: key BB572E0E2D182910: public key "Gentoo Linux Release Engineering (Automated Weekly Release Key) <releng@gentoo.org>" imported
gpg: Total number processed: 1
gpg:               imported: 1

```

Alternativement, vous pouvez utiliser [WKD](/wiki/WKD "WKD") pour télécharger les clés :

`user $` `gpg --auto-key-locate=clear,nodefault,wkd --locate-key releng@gentoo.org`

```
gpg: key 9E6438C817072058: public key "Gentoo Linux Release Engineering (Gentoo Linux Release Signing Key) <releng@gentoo.org>" imported
gpg: key BB572E0E2D182910: public key "Gentoo Linux Release Engineering (Automated Weekly Release Key) <releng@gentoo.org>" imported
gpg: Total number processed: 2
gpg:               imported: 2
gpg: no ultimately trusted keys found
pub   dsa1024 2004-07-20 [SC] [expires: 2025-07-01]
      D99EAC7379A850BCE47DA5F29E6438C817072058
uid           [ unknown] Gentoo Linux Release Engineering (Gentoo Linux Release Signing Key) <releng@gentoo.org>
sub   elg2048 2004-07-20 [E] [expires: 2025-07-01]

```

Ou si vous utilisez un support officiel, importez la clé depuis /usr/share/openpgp-keys/gentoo-release.asc (fournie par [sec-keys/openpgp-keys-gentoo-release](https://packages.gentoo.org/packages/sec-keys/openpgp-keys-gentoo-release)) :

`user $` `gpg --import /usr/share/openpgp-keys/gentoo-release.asc`

```
gpg: directory '/home/larry/.gnupg' created
gpg: keybox '/home/larry/.gnupg/pubring.kbx' created
gpg: key DB6B8C1F96D8BF6D: 2 signatures not checked due to missing keys
gpg: /home/larry/.gnupg/trustdb.gpg: trustdb created
gpg: key DB6B8C1F96D8BF6D: public key "Gentoo ebuild repository signing key (Automated Signing Key) <infrastructure@gentoo.org>" imported
gpg: key 9E6438C817072058: 3 signatures not checked due to missing keys
gpg: key 9E6438C817072058: public key "Gentoo Linux Release Engineering (Gentoo Linux Release Signing Key) <releng@gentoo.org>" imported
gpg: key BB572E0E2D182910: 1 signature not checked due to a missing key
gpg: key BB572E0E2D182910: public key "Gentoo Linux Release Engineering (Automated Weekly Release Key) <releng@gentoo.org>" imported
gpg: key A13D0EF1914E7A72: 1 signature not checked due to a missing key
gpg: key A13D0EF1914E7A72: public key "Gentoo repository mirrors (automated git signing key) <repomirrorci@gentoo.org>" imported
gpg: Total number processed: 4
gpg:               imported: 4
gpg: no ultimately trusted keys found
```

Ensuite, vérifiez la signature cryptographique :

`user $` `gpg --verify install-ppc-minimal-20240107T170309Z.iso.asc install-ppc-minimal-20240107T170309Z.iso`

```
gpg: Signature made Sun 07 Jan 2024 03:01:10 PM CST
gpg:                using RSA key 534E4209AB49EEE1C19D96162C44695DB9F6043D
gpg: Good signature from "Gentoo Linux Release Engineering (Automated Weekly Release Key) <releng@gentoo.org>" [unknown]
gpg: WARNING: This key is not certified with a trusted signature!
gpg:          There is no indication that the signature belongs to the owner.
Primary key fingerprint: 13EB BDBE DE7A 1277 5DFD  B1BA BB57 2E0E 2D18 2910
     Subkey fingerprint: 534E 4209 AB49 EEE1 C19D  9616 2C44 695D B9F6 043D

```

Pour être absolument certains que tout est valide, vérifiez l’empreinte digitale affichée avec l’empreinte digitale disponible sur la [page de signatures](https://www.gentoo.org/downloads/signatures/).

**Remarque**

Une bonne pratique est généralement de marquer une clé importée comme sûre lorsque vous êtes certain de sa fiabilité. Lorsque les clés sont vérifiées, gpg n’indiquera plus _unknown_ avec un avertissement pour une signature non sûre.

## Créer le support d’amorçage

Bien sûr, avec juste un fichier ISO téléchargé, l’installation de Gentoo Linux ne peut pas être démarrée. Le fichier ISO doit être écrit sur un support amorçable ( _bootable_). Cela signifie que l’image sera extraite sur un système de fichiers ou écrite directement sur un périphérique de stockage.

### Créer une clé USB amorçable

**Attention !**

Utiliser des outils comme Ventoy pour créeer un disque USB peut causer des erreurs au démarrage.

La plupart des systèmes modernes savent démarrer avec un périphérique USB.

#### Créer depuis Linux

dd est typiquement disponible dans toutes les distributions Linux, et peut être utilisé pour créer une clé USB amorçable Gentoo.

##### Déterminer le chemin d’accès du périphérique USB

Avant la création, le chemin d’accès du périphérique USB doit être identifié.

dmesg affichera des informations détaillées décrivant le périphérique de stockage lors de son ajout au système :

`root #` `dmesg`

```
[268385.319745] sd 19:0:0:0: [sdd] 60628992 512-byte logical blocks: (31.0 GB/28.9 GiB)
```

Alternativement, lsblk peut être utilisé pour visualiser les périphériques de stockage disponibles :

`root #` `lsblk`

```
sdd           8:48   1  28.9G  0 disk
├─sdd1        8:49   1   246K  0 part
├─sdd2        8:50   1   2.8M  0 part
├─sdd3        8:51   1 463.5M  0 part
└─sdd4        8:52   1   300K  0 part

```

Une fois le périphérique identifié, il faut ajouter le préfixe « /dev/ » pour obtenir le chemin d’accès. Exemple : /dev/sdd.

**Conseil**

Utiliser le périphérique de base, ie. sdd, à la place de sdd1, est recommandé par Gentoo car le fichier .iso contient le schéma de partitionnement complet [GPT](/wiki/GPT "GPT").

##### Créer avec dd

**Attention !**

Vérifier très attentivement la cible « of=target » avant de lancer la commande car dd va écraser tout le contenu.

Un exemple avec (/dev/sdd) et le fichier install-amd64-minimal-<release timestamp>.iso :

`root #` `dd if=install-amd64-minimal-<release timestamp>.iso of=/dev/sdd bs=4096 status=progress && sync`

**Remarque**

« if= » indique le fichier source, « of= » indique le fichier de sortie qui est dans notre cas un périphérique.

**Conseil**

« bs=4096 » est ajouté car il accélère en général le transfert ; « status=progress » affiche des statistiques en temps réel sur le transfert.

#### Écrire avec Windows

[win32diskimager](https://sourceforge.net/projects/win32diskimager/) est simple d’usage pour créer un disque USB sur Windows.

Ils fournissent aussi une très bonne documentation pour aider si nécessaire.

**Attention !**

Téléchez seulement depuis [https://sourceforge.net/projects/win32diskimager/](https://sourceforge.net/projects/win32diskimager/)

### Graver un disque

**See also**

Des instructions plus complète peuvent être consultées sur [CD/DVD/BD\_writing#Image\_writing](/wiki/CD/DVD/BD_writing#Image_writing.2Ffr "CD/DVD/BD writing").

#### Gravure depuis Microsoft Windows 7 et supérieur

À partir de Microsoft Windows 7, la gravure est possible sans logiciel tiers supplémentaire. Insérez simplement un CD gravable, allez dans le répertoire du fichier ISO, cliquez droit et choisissez « Graver une image disque ».

#### Gravure sous Linux

Sur Linux, l’utilitaire cdrecord du paquet [app-cdr/cdrtools](https://packages.gentoo.org/packages/app-cdr/cdrtools) peut graver des images ISO.

Pour graver le fichier ISO sur le CD du périphérique /dev/sr0 (c’est le premier lecteur de CD du système, remplacez le chemin vers le périphérique correspondant si besoin) :

`user $` `cdrecord dev=/dev/sr0 install-ppc-minimal-20141204.iso`

Les utilisateurs qui préfèrent une interface graphique peuvent utiliser K3B, du paquet [kde-apps/k3b](https://packages.gentoo.org/packages/kde-apps/k3b). Dans K3B, allez dans Outils et choisissez Graver une image CD.

## Démarrage

[Handbook:PPC/Blocks/Booting/fr](/index.php?title=Handbook:PPC/Blocks/Booting/fr&action=edit&redlink=1 "Handbook:PPC/Blocks/Booting/fr (page does not exist)")

### LiveGUI root access

sudo has been configured to run without the need of a password on the LiveGUI as both the root and gentoo have a scrambled password.

To gain access to the superuser account, in any terminal run:

`root #` `sudo -i`

### Configuration matérielle supplémentaire

Lorsque le support d’installation se lance, il tente de détecter le matériel et charge les modules appropriés du noyau pour prendre en compte le matériel. Dans la grande majorité des cas, il fait un très bon travail. Toutefois, dans certains cas, il peut ne pas charger automatiquement les modules du noyau nécessaires au système. Si la détection automatique du bus PCI a oublié certains matériels du système, les modules appropriés du noyau doivent être chargées manuellement.

Dans l’exemple suivant, le module 8139too (qui prend en charge certains types d’interfaces réseau) est chargé :

`root #` `modprobe 8139too`

### Facultatif : comptes utilisateurs

Si d’autres personnes ont besoin d’accéder à l’environnement d’installation, ou s’il est nécessaire d’exécuter des commandes en tant qu’utilisateur non- _root_ sur le support d’installation (comme pour discuter sur IRC à l’aide de irssi sans être _root_ pour des raisons de sécurité), alors un compte utilisateur supplémentaire doit être créé et un mot de passe _root_ robuste défini.

Pour changer le mot de passe _root_, utilisez l’utilitaire passwd :

`root #` `passwd`

```
New password: (Entrez le nouveau mot de passe)
Re-enter password: (Entrez de nouveau le mot de passe)

```

Pour créer un compte utilisateur, entrez d’abord ses informations d’identification, suivies par le mot de passe du compte. Les commandes useradd et passwd sont utilisées pour ces tâches.

Dans l’exemple suivant, un utilisateur appelé « jean » est créé :

`root #` `useradd -m -G users,wheel jean
`

`root #` `passwd jean`

```
New password: (Entrez le mot de passe de jean)
Re-enter password: (Entrez de nouveau le mot de passe de jean)

```

Pour passer de l’utilisateur _root_ à l’utilisateur fraîchement créé, utilisez la commande su :

`root #` `su - jean`

### Facultatif : consulter la documentation pendant l’installation

#### TTY

Pour afficher le manuel Gentoo lors de l'installation, créez tout d’abord un compte utilisateur, comme décrit ci-dessus. Puis appuyez sur `Alt` + `F2` pour accéder à un nouveau terminal et vous identifier avec le nouvel utilisateur.
En suivant le [principe du moindre privilège](https://en.wikipedia.org/wiki/Principle_of_least_privilege "wikipedia:Principle of least privilege"), il est mieux d’éviter d’effectuer une action ou une tâche avec plus de droits que nécessaire. Le compte _root_ a un contrôle total sur le système et, pour cela, ne devrait être utilisé qu’avec parcimonie.

Lors de l’installation, la commande links peut être utilisée pour parcourir le manuel Gentoo – bien sûr cela nécessite que la connexion Internet fonctionne.

`user $` `links https://wiki.gentoo.org/wiki/Handbook:PPC/fr`

Pour revenir au terminal d’origine, appuyez sur `Alt` + `F1`.

**Conseil**

Sept TTY sont disponibles sur les supports Gentoo. La bascule sur l’un d’eux se fait via `Alt` et une touche fonction `F1` à `F7`. Il peut être utile de passer sur un autre terminal pendant qu’une opération se termine, ouvrir une documentation…

#### GNU Screen

L’utilitaire [Screen](/wiki/Screen "Screen") est installé par défaut sur le support d’installation officiel de Gentoo. Il peut être plus efficace pour un utilisateur expérimenté d’utiliser screen afin de visualiser les instructions d’installation dans des panneaux séparés plutôt que d’utiliser la technique des multiples TTY mentionnée ci-dessus.

### Facultatif : démarrer le daemon SSH

Pour permettre à d’autres utilisateurs d’accéder au système lors de l’installation (peut-être pour obtenir de l’aide lors d'une installation, ou même la faire à distance), un compte utilisateur doit être créé (comme documenté plus tôt) et le serveur SSH doit être démarré.

Pour lancer le daemon SSH sur un système d’initialisation OpenRC, exécutez la commande suivante :

`root #` `rc-service sshd start`

**Remarque**

Si un utilisateur ouvre une session sur le système, il reçoit un message indiquant que la clé d’hôte pour ce système doit être confirmée (par le biais de ce qu'on appelle une empreinte). Ceci est dû au fait que c’est la première fois que quelqu’un ouvre une session sur le système. Cependant, plus tard, lorsque le système est mis en place et que quelqu’un se connecte sur le nouveau système, le client SSH l’avertit que la clé de l’hôte a été changée. C'est parce que l’utilisateur – pour SSH – se connecte désormais à un serveur différent (à savoir le système Gentoo fraîchement installé plutôt qu’à l’environnement en cours d’utilisation pour l’installation ). Suivez alors les instructions données à l'écran pour remplacer la clé de l’hôte sur le système client.

Pour être en mesure d'utiliser sshd, le réseau a besoin de fonctionner correctement. Continuez l’installation avec le chapitre [Configurer le réseau](/wiki/Handbook:PPC/Installation/Networking/fr "Handbook:PPC/Installation/Networking/fr").

[← À propos de l’installation](/wiki/Handbook:PPC/Installation/About/fr "Previous part") [Accueil](/wiki/Handbook:PPC/fr "Handbook:PPC/fr") [Configurer le réseau →](/wiki/Handbook:PPC/Installation/Networking/fr "Next part")