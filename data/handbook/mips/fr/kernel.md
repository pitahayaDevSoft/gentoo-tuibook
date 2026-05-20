# Kernel

Other languages:

- [Deutsch](/wiki/Handbook:MIPS/Installation/Kernel/de "Handbuch:MIPS/Installation/Kernel (100% translated)")
- [English](/wiki/Handbook:MIPS/Installation/Kernel "Handbook:MIPS/Installation/Kernel (100% translated)")
- [español](/wiki/Handbook:MIPS/Installation/Kernel/es "Manual de Gentoo: MIPS/Instalación/Núcleo (100% translated)")
- français
- [italiano](/wiki/Handbook:MIPS/Installation/Kernel/it "Handbook:MIPS/Installation/Kernel/it (100% translated)")
- [magyar](/wiki/Handbook:MIPS/Installation/Kernel/hu "Handbook:MIPS/Installation/Kernel/hu (100% translated)")
- [polski](/wiki/Handbook:MIPS/Installation/Kernel/pl "Handbook:MIPS/Installation/Kernel (100% translated)")
- [português do Brasil](/wiki/Handbook:MIPS/Installation/Kernel/pt-br "Manual:MIPS/Instalação/Kernel (100% translated)")
- [русский](/wiki/Handbook:MIPS/Installation/Kernel/ru "Handbook:MIPS/Installation/Kernel (100% translated)")
- [தமிழ்](/wiki/Handbook:MIPS/Installation/Kernel/ta "கையேடு:MIPS/நிறுவல்/கர்னல் (100% translated)")
- [中文（中国大陆）‎](/wiki/Handbook:MIPS/Installation/Kernel/zh-cn "手册：MIPS/安装/配置Linux内核 (100% translated)")
- [日本語](/wiki/Handbook:MIPS/Installation/Kernel/ja "ハンドブック:MIPS/インストール/カーネル (100% translated)")
- [한국어](/wiki/Handbook:MIPS/Installation/Kernel/ko "Handbook:MIPS/Installation/Kernel/ko (100% translated)")

[Sommaire du manuel](/wiki/Handbook:MIPS "Handbook:MIPS")[Installation](/wiki/Handbook:MIPS/Full/Installation/fr "Handbook:MIPS/Full/Installation/fr")[À propos de l’installation](/wiki/Handbook:MIPS/Installation/About/fr "Handbook:MIPS/Installation/About/fr")[Choix du support](/wiki/Handbook:MIPS/Installation/Media/fr "Handbook:MIPS/Installation/Media/fr")[Configurer le réseau](/wiki/Handbook:MIPS/Installation/Networking/fr "Handbook:MIPS/Installation/Networking/fr")[Préparer les disques](/wiki/Handbook:MIPS/Installation/Disks/fr "Handbook:MIPS/Installation/Disks/fr")[Installer l’archive _stage3_](/wiki/Handbook:MIPS/Installation/Stage/fr "Handbook:MIPS/Installation/Stage/fr")[Installer le système de base](/wiki/Handbook:MIPS/Installation/Base/fr "Handbook:MIPS/Installation/Base/fr")Configurer le noyau[Configurer le système](/wiki/Handbook:MIPS/Installation/System/fr "Handbook:MIPS/Installation/System/fr")[Installer les outils](/wiki/Handbook:MIPS/Installation/Tools/fr "Handbook:MIPS/Installation/Tools/fr")[Configurer le système d’amorçage](/wiki/Handbook:MIPS/Installation/Bootloader/fr "Handbook:MIPS/Installation/Bootloader/fr")[Finaliser](/wiki/Handbook:MIPS/Installation/Finalizing/fr "Handbook:MIPS/Installation/Finalizing/fr")[Utiliser Gentoo](/wiki/Handbook:MIPS/Full/Working/fr "Handbook:MIPS/Full/Working/fr")[Introduction à Portage](/wiki/Handbook:MIPS/Working/Portage/fr "Handbook:MIPS/Working/Portage/fr")[Les options de la variable USE](/wiki/Handbook:MIPS/Working/USE/fr "Handbook:MIPS/Working/USE/fr")[Les fonctionnalités de Portage](/wiki/Handbook:MIPS/Working/Features/fr "Handbook:MIPS/Working/Features/fr")[Scripts d’initialisation systèmes](/wiki/Handbook:MIPS/Working/Initscripts/fr "Handbook:MIPS/Working/Initscripts/fr")[Variables d’environnement](/wiki/Handbook:MIPS/Working/EnvVar/fr "Handbook:MIPS/Working/EnvVar/fr")[Utiliser Portage](/wiki/Handbook:MIPS/Full/Portage/fr "Handbook:MIPS/Full/Portage/fr")[Fichiers et répertoires](/wiki/Handbook:MIPS/Portage/Files/fr "Handbook:MIPS/Portage/Files/fr")[Les variables](/wiki/Handbook:MIPS/Portage/Variables/fr "Handbook:MIPS/Portage/Variables/fr")[Mélanger plusieurs branches logicielles](/wiki/Handbook:MIPS/Portage/Branches/fr "Handbook:MIPS/Portage/Branches/fr")[Outils supplémentaires](/wiki/Handbook:MIPS/Portage/Tools/fr "Handbook:MIPS/Portage/Tools/fr")[Dépôt personnalisé](/wiki/Handbook:MIPS/Portage/CustomTree/fr "Handbook:MIPS/Portage/CustomTree/fr")[Fonctionnalités avancées](/wiki/Handbook:MIPS/Portage/Advanced/fr "Handbook:MIPS/Portage/Advanced/fr")[Configuration du réseau avec OpenRC](/wiki/Handbook:MIPS/Full/Networking/fr "Handbook:MIPS/Full/Networking/fr")[Bien démarrer](/wiki/Handbook:MIPS/Networking/Introduction/fr "Handbook:MIPS/Networking/Introduction/fr")[Configuration avancée](/wiki/Handbook:MIPS/Networking/Advanced/fr "Handbook:MIPS/Networking/Advanced/fr")[Les modules réseau](/wiki/Handbook:MIPS/Networking/Modular/fr "Handbook:MIPS/Networking/Modular/fr")[Sans fil](/wiki/Handbook:MIPS/Networking/Wireless/fr "Handbook:MIPS/Networking/Wireless/fr")[Ajouter des fonctionnalités](/wiki/Handbook:MIPS/Networking/Extending/fr "Handbook:MIPS/Networking/Extending/fr")[Gestion dynamique](/wiki/Handbook:MIPS/Networking/Dynamic/fr "Handbook:MIPS/Networking/Dynamic/fr")

## Contents

- [1Facultatif : Installation de micrologiciels](#Facultatif_:_Installation_de_micrologiciels)
- [2Microcode](#Microcode)
  - [2.1Micrologiciel Linux](#Micrologiciel_Linux)
    - [2.1.1Chargement de micrologiciel](#Chargement_de_micrologiciel)
- [3sys-kernel/installkernel](#sys-kernel.2Finstallkernel)
  - [3.1Programme d’amorçage ( _bootloader_)](#Programme_d.E2.80.99amor.C3.A7age_.28bootloader.29)
    - [3.1.1GRUB](#GRUB)
    - [3.1.2Choix traditionnel et autres programmes d’amorçage ((e)lilo, syslinux, etc.)](#Choix_traditionnel_et_autres_programmes_d.E2.80.99amor.C3.A7age_.28.28e.29lilo.2C_syslinux.2C_etc..29)
  - [3.2_Initramfs_](#Initramfs)
    - [3.2.1Chroot detection](#Chroot_detection)
- [4Configuration et compilation du noyau](#Configuration_et_compilation_du_noyau)
  - [4.1Approche manuelle](#Approche_manuelle)
  - [4.2Installer les sources du noyau](#Installer_les_sources_du_noyau)
    - [4.2.1Procédure _modprobed-db_](#Proc.C3.A9dure_modprobed-db)
    - [4.2.2Procédure manuelle](#Proc.C3.A9dure_manuelle)
      - [4.2.2.1Activer les options nécessaires](#Activer_les_options_n.C3.A9cessaires)
    - [4.2.3Optionnel : modules noyaux signés](#Optionnel_:_modules_noyaux_sign.C3.A9s)

### Facultatif : Installation de micrologiciels

### Microcode

#### Micrologiciel Linux

Sur beaucoup de systèmes, les micrologiciels non libres sont nécessaires pour le fonctionnement de certaines fonctions matérielles. Le paquet [sys-kernel/linux-firmware](https://packages.gentoo.org/packages/sys-kernel/linux-firmware) contient des micrologiciels pour beaucoup, mais pas tous, ces périphériques.

**Conseil**

La plupart des cartes sans fil et graphiques ( _GPU_) ont besoin de micrologiciel pour fonctionner.

`root #` `emerge --ask sys-kernel/linux-firmware`

**Remarque**

Installer certains micrologiciels nécessite souvent d’accepter la licence associée. Si nécessaire, visitez la [section gestion des licences](/wiki/Handbook:MIPS/Working/Portage/fr#Licenses "Handbook:MIPS/Working/Portage/fr") du manuel pour de l’aide à propos des licences.

##### Chargement de micrologiciel

Les micrologiciels sont généralement chargés avec le module du noyau ( _kernel_) associé. Cela signifie que le micrologiciel doit être compilé avec le noyau en utilisant « CONFIG\_EXTRA\_FIRMWARE » si le module est configuré avec « Y » (oui) à la place de « M » (module). Dans la plupart des situations, intégrer un module qui a besoin d’un micrologiciel peut être compliqué ou casser le chargement.

## sys-kernel/installkernel

[Installkernel](/wiki/Installkernel/fr "Installkernel/fr") peut être utilisé pour automatiser l’installation du noyau parmi la génération d’ [initramfs](/wiki/Initramfs "Initramfs"), _[unified kernel image](/wiki/Unified_kernel_image "Unified kernel image")_ ou la configuration du programme d’amorçage. [sys-kernel/installkernel](https://packages.gentoo.org/packages/sys-kernel/installkernel) propose deux façons pour installer: le traditionnel installkernel originaire de Debian et kernel-install de [systemd](/wiki/Systemd/fr "Systemd/fr"). Lequel choisir dépend, parmi d’autres critères, du système d’initialisation et du programme d’amorçage. Par défaut, kernel-install système est utilisé sur des profils systemd, tandis que le traditionnel installkernel est le choix par défaut pour les autres profils.

### Programme d’amorçage ( _bootloader_)

Il est temps maintenant de réfléchir au programme d’amorçage voulu par l’utilisateur pour son système. Si vous n’êtes pas sûr, choisir la « voie traditionnelle » ci-dessous.

**Important**

Only one selection is required in the following subsection, if unsure of which to use then go with the first listed for now. It's always possible to switch at a later date if required.

#### GRUB

Les utilisateurs de GRUB peuvent utiliser au choix kernel-install de systemd ou le traditionnel installkernel de Debian. Le drapeau ( _USE_) [systemd](https://packages.gentoo.org/useflags/systemd) [USE flag/fr](/wiki/USE_flag/fr "USE flag/fr") permet de passer d’une implémentation à l’autre. Pour exécuter automatiquement grub-mkconfig lors de l’installation d’un noyau, activer [grub](https://packages.gentoo.org/useflags/grub) [USE flag/fr](/wiki/USE_flag/fr "USE flag/fr") [USE flag](/wiki/USE_flag/fr "USE flag/fr").

**Remarque**

GRUB requires kernels to be installed to /boot.

FILE **`/etc/portage/package.use/installkernel`**

```
sys-kernel/installkernel grub

```

`root #` `emerge --ask sys-kernel/installkernel`

`root #` `mkdir -p /etc/cmdline.d`

`root #` `ln -s /etc/kernel/cmdline /etc/cmdline.d/00-installkernel.conf`

`root #` `emerge --ask sys-kernel/installkernel`

}}

**Remarque**

systemd-boot requires kernels to be installed to /efi.

**Remarque**

When [app-emulation/virt-firmware](https://packages.gentoo.org/packages/app-emulation/virt-firmware) is used to configure the UEFI ensure that the kernel-bootcfg-boot-successful service is enabled before attempting to install the kernel. This implementation of EFIstub booting is the default for systemd systems.

`root #` `systemctl enable kernel-bootcfg-boot-successful.service`

**Remarque**

EFIstub requires kernels to be installed to /efi.

#### Choix traditionnel et autres programmes d’amorçage ((e)lilo, syslinux, etc.)

Le traditionnel chemin /boot (pour LILO, syslinux, etc.) est utilisé par défaut si les drapeaux _USE_ [grub](https://packages.gentoo.org/useflags/grub) [USE flag/fr](/wiki/USE_flag/fr "USE flag/fr"), [systemd-boot](https://packages.gentoo.org/useflags/systemd-boot) [USE flag/fr](/wiki/USE_flag/fr "USE flag/fr"), [efistub](https://packages.gentoo.org/useflags/efistub) [USE flag/fr](/wiki/USE_flag/fr "USE flag/fr") et [uki](https://packages.gentoo.org/useflags/uki) [USE flag/fr](/wiki/USE_flag/fr "USE flag/fr") ne sont **pas** activés. Aucune action supplémentaire n’est nécessaire.

### _Initramfs_

Un système de fichiers basés sur la RAM pour l’initialisation, ou [initramfs](/wiki/Initramfs "Initramfs"), peut être nécessaire pour démarrer un système. Un large panel de cas en ont besoin, les cas courants incluent :

- noyaux où les pilotes pour le système de fichiers ou le stockage sont des modules ;
- un partitionnement séparés de /usr/ ou /var/ ;
- un système de fichiers racine chiffré.

**Conseil**

[Les noyaux distribués](/wiki/Project:Distribution_Kernel "Project:Distribution Kernel") peuvent être utilisés avec _initramfs_, comme beaucoup de pilotes pour le système de fichiers ou le stockage sont des modules.

En plus de monter une partition racine, _initramfs_ peut accomplir d’autres missions, comme :

- lancer une vérifier de la consistance d’un système de fichiers fsck, dans le cas d’un arrêt brutal du système ;
- fournir un environnement de secours en cas d’erreur au démarrage.

[Installkernel](/wiki/Installkernel/fr "Installkernel/fr") peut générer automatiquement un _initramfs_ en installant le noyau si les drapeaux _USE_ [dracut](https://packages.gentoo.org/useflags/dracut) [USE flag/fr](/wiki/USE_flag/fr "USE flag/fr") ou [ugrd](https://packages.gentoo.org/useflags/ugrd) [USE flag/fr](/wiki/USE_flag/fr "USE flag/fr") sont activés :

FILE **`/etc/portage/package.use/installkernel`**

```
sys-kernel/installkernel dracut

```

#### Chroot detection

Bootloaders such as [systemd-boot](/wiki/Systemd/systemd-boot "Systemd/systemd-boot") and [EFI stub](/wiki/EFI_stub "EFI stub") use the kernel arguments of the running system as set in /proc/cmdline by default. Because of the wide range of ways Gentoo can be installed users will randomly get tripped up by this.

To help solve any issues this may cause, [sys-kernel/installkernel](https://packages.gentoo.org/packages/sys-kernel/installkernel) will check if it is running in a chroot and fail if the kernel command line is not explicitly configured. Otherwise the bootloader would use the install media's boot arguments which would lead to boot failure.

One way to satisfy [sys-kernel/installkernel](https://packages.gentoo.org/packages/sys-kernel/installkernel) is by using Dracut's config file to set the root partition UUID as shown below, or alternatively for more information on what this check helps with and different ways to configure it, see [Installkernel#Install\_chroot\_check](/wiki/Installkernel#Install_chroot_check.2Ffr "Installkernel").

`root #` `mkdir /etc/dracut.conf.d`

`root #` `blkid`

```
/dev/sda3: UUID="2122cd72-94d7-4dcc-821e-3705926deecc"
```

In the above example, the root partition is /dev/sda3 and the UUID is 2122cd72-94d7-4dcc-821e-3705926deecc.
The dracut config file would then look like:

FILE **`/etc/dracut.conf.d/00-installkernel.conf`**

```
kernel_cmdline=" root=UUID=2122cd72-94d7-4dcc-821e-3705926deecc " # Note leading and trailing spaces

```

`root #` `emerge --ask sys-kernel/installkernel`

## Configuration et compilation du noyau

**Conseil**

Il peut être judicieux d’utiliser un noyau distribué au premier démarrage car cela fournit une méthode simple pour éviter les erreurs systèmes et de configuration du noyau. Ayez toujours un noyau fonctionnel en secours permet d’accélérer le débogage et réduit l’anxiété qu’une mise à jour système empêchera votre ordinateur de démarrer.

A common misconception is that a manually compiled kernel will use a lot less RAM than a pre configured distribution kernel. Due to the modular nature of the Linux kernel, only what is needed by the system is loaded and in most cases resulting in similar memory usage.

**Important**

Only one selection is required in the following subsection, if unsure of which to use then go with the first listed for now. It's always possible to switch at a later date if required.

Classés du moins fréquent au plus courant :

[Approche manuelle](/wiki/Handbook:MIPS/Installation/Kernel/fr#Alternative:_Manual_configuration "Handbook:MIPS/Installation/Kernel/fr")les nouvelles sources du noyau sont installées par le gestionnaire de paquets. Le noyau est manuellement configuré, compilé, installé en utilisant eselect kernel et un tas de commande make. Les mises à jour futures du noyau répéteront la même procédure de configuration, compilation et installation. Cette façon de faire implique le plus d‘opérations, mais offre un contrôle maximal sur la mise à jour du noyau.[Approche hybride : _Genkernel_](/wiki/Handbook:MIPS/Installation/Kernel/fr#Alternative:_Genkernel "Handbook:MIPS/Installation/Kernel/fr")nous utilisons le terme « hybride » ici, mais notons que les noyaux distribués et la procédure manuelle partagent le même but. Les nouveaux noyaux sont installés via le gestionnaire de paquets. Les administrateurs systèmes peuvent utiliser le genkernel de Gentoo pour configurer, compiler et installer le noyau, ses modules associés, et (optionnellement, **pas** activé par défaut) un fichier _initramfs_. Il est possible de fournir un fichier de configuration du noyau si nécessaire. Les mises à jour futures du noyau auront besoin d’une implication de l’administrateur pour lancer eselect kernel, genkernel et potentiellement d’autres commandes à chaque mise à jour. Cette option ne devrait être choisie que par les utilisateurs qui ont besoin de genkernel

Le cœur de toute distribution est le noyau Linux. C’est la couche située entre les programmes de l’utilisateur et le matériel du système. Même si le guide d’installation propose à ses utilisateurs plusieurs sources du noyau possibles, une liste complète des sources, avec description, est disponible sur la page [Vue d’ensemble des noyaux](/wiki/Kernel/Overview "Kernel/Overview").

**Conseil**

Les tâches d’installation du noyau comme copier l’image du noyau vers /boot ou l’ [EFI System Partition](/wiki/EFI_System_Partition "EFI System Partition"), générer un [initramfs](/wiki/Initramfs "Initramfs") ou une [Unified Kernel Image](/wiki/Unified_Kernel_Image "Unified Kernel Image") et mettre à jour la configuration du programme d’amorçage peuvent être automatisé avec [installkernel](/wiki/Installkernel/fr "Installkernel/fr"). Les utilisateurs peuvent souhaiter installer et configurer [sys-kernel/installkernel](https://packages.gentoo.org/packages/sys-kernel/installkernel) avant de continuer. Lisez ci-dessous la [section d’installation du noyau](/wiki/Handbook:MIPS/Installation/Kernel#Kernel_installation.2Ffr "Handbook:MIPS/Installation/Kernel") pour davantage d’informations.

### Approche manuelle

### Installer les sources du noyau

Lors de l’installation et la compilation du noyau pour les systèmes basés sur mips, Gentoo recommande le paquet [sys-kernel/mips-sources](https://packages.gentoo.org/packages/sys-kernel/mips-sources).

Choisissez les sources du noyau appropriées et installez les en utilisant emerge :

`root #` `emerge --ask sys-kernel/mips-sources`

Cela installera les sources du noyau Linux dans le répertoire /usr/src/ dans un répertoire versionné. Un lien symbolique /usr/src/linux ne sera pas créé sans le drapeau _USE_ [symlink](https://packages.gentoo.org/useflags/symlink) [USE flag/fr](/wiki/USE_flag/fr "USE flag/fr") activé pour le paquet de sources du noyau choisi.

Il est usuelle d’avoir un lien symbolique /usr/src/linux pour pointer vers les sources du noyau actuel. Cependant, ce lien n’est pas créé par défaut. Un autre façon de créer ce lien est d’utiliser le module _eselect kernel_.

Pour davantage d’informations concernant le but du lien symbolique et comment le gérer, consultez [Kernel/Upgrade/fr](/wiki/Kernel/Upgrade/fr "Kernel/Upgrade/fr").

Premièrement, listons les noyaux installés :

`root #` `eselect kernel list`

```
Available kernel symlink targets:
  [1]   linux-6.19.1-gentoo

```

Pour créer un lien symbolique linux, copier :

`root #` `eselect kernel set 1`

`root #` `ls -l /usr/src/linux`

```
lrwxrwxrwx    1 root   root    12 Oct 13 11:04 /usr/src/linux -> linux-6.19.1-gentoo

```

Configurer manuellement un noyau est couramment vu comme la procédure la plus difficile pour un administrateur système. Rien n’est moins vrai… mais après avoir configuré quelques noyaux, plus personne ne se souvient que c’était difficile ! Il y a deux manières pour un utilisateur Gentoo de gérer un noyau manuellement, les deux sont listées ci-dessous :

**Important**

Only one selection is required in the following subsection, if unsure of which to use then go with the first listed for now. It's always possible to switch at a later date if required.

#### Procédure _modprobed-db_

Une manière très simple pour gérer le noyau est d’installer d’abord [sys-kernel/gentoo-kernel-bin](https://packages.gentoo.org/packages/sys-kernel/gentoo-kernel-bin) puis d’utiliser [sys-kernel/modprobed-db](https://packages.gentoo.org/packages/sys-kernel/modprobed-db) pour réunir les informations à propos des besoins du système. modprobed-db est un outil qui supervise le système via _crontab_ et ajoute tous les modules de tous les périphériques branchés une fois sur le système pour s‘assurer que tous les besoins de l’utilisateur sont supportés. Par exemple, si un contrôleur Xbox est branché après l’installation, modprobed-db va ajouter les modules pour être compilés la prochaine fois que le noyau est reconstruit. Ce sujet est détaillé dans l’article [Modprobed-db](/wiki/Modprobed-db "Modprobed-db").

For now please follow installing [sys-kernel/gentoo-kernel-bin](https://packages.gentoo.org/packages/sys-kernel/gentoo-kernel-bin) via [Distribution\_kernels](/wiki/Handbook:MIPS/Installation/Kernel#Distribution_kernels.2Ffr "Handbook:MIPS/Installation/Kernel").

Next, install [sys-kernel/modprobed-db](https://packages.gentoo.org/packages/sys-kernel/modprobed-db):

`root #` `emerge --ask sys-kernel/modprobed-db`

Please watch out for further steps related to modprobed-db in the Handbook.

More on this topic can be found in the [Modprobed-db](/wiki/Modprobed-db "Modprobed-db") article.
}}

#### Procédure manuelle

Cette procédure permet à l’utilisateur d’avoir un contrôle total sur ses noyaux avec un minimum d’outils aidant comme il le désire. Certains considère que cela a un intérêt de rendre cela difficile.

Cependant, avec ce choix, une chose est vraie : c est vital de connaître le système quand un noyau est configuré manuellement. La plupart des informations nécessaires peuvent être recueillies en installant le paquet [sys-apps/pciutils](https://packages.gentoo.org/packages/sys-apps/pciutils) qui contient la commande lspci :

`root #` `emerge --ask sys-apps/pciutils`

**Remarque**

À l’intérieur d’un _chroot_, il est possible d’ignorer sans risque toutes les mises en garde (du genre _pcilib: cannot open /sys/bus/pci/devices_) que lspci pourrait afficher.

Une autre source d’information est d’exécuter la commande lsmod pour voir quels modules du noyau sont utilisés par le média d’installation afin de savoir quoi activer plus tard.

Il est maintenant temps d’accéder au répertoire source du noyau.

`root #` `cd /usr/src/linux
`

**Conseil**

Pour consulter la liste complète des paramètres de make pour le noyau, lancez `make help`.

Le noyau a une auto-détection de modules utilisé par l’ _installcd_ qui permet un bon point de départ afin de configurer son propre noyau. Il peut être appelé avec :

`root #` `make localmodconfig`

Il est maintenant temps de configurer avec nconfig :

`root #` `make nconfig`

Now it's time to decide if modules signing is required in the steps listed in [here](/wiki/Handbook:MIPS/Installation/Kernel#Optional:_Signed_kernel_modules_2.2Ffr "Handbook:MIPS/Installation/Kernel")

If not, proceed to building described [here](/wiki/Handbook:MIPS/Installation/Kernel#Compiling_and_installing.2Ffr "Handbook:MIPS/Installation/Kernel")

###### Activer les options nécessaires

#### Optionnel : modules noyaux signés

Pour automatiquement signer les modules noyaux, activez l’option CONFIG\_MODULE\_SIG\_ALL :

KERNEL **Signer les modules noyaux CONFIG\_MODULE\_SIG\_ALL**

```
[*] Enable loadable module support
  -*-   Module signature verification
    [*]     Automatically sign all modules
    Which hash algorithm should modules be signed with? (Sign modules with SHA-512) --->

```

Vous pouvez changer l’algorithme de signature ( _hash_) si vous le désirez.

Pour s’assurer que tous les modules signés le sont avec une signature valide, activez également l’option CONFIG\_MODULE\_SIG\_FORCE :

KERNEL **Forcer la signature des modules noyaux CONFIG\_MODULE\_SIG\_FORCE**

```
[*] Enable loadable module support
  -*-   Module signature verification
    [*]     Require modules to be validly signed
    [*]     Automatically sign all modules
    Which hash algorithm should modules be signed with? (Sign modules with SHA-512) --->

```

Pour utiliser une clé personnalisée, spécifiez le chemin d’accès dans CONFIG\_MODULE\_SIG\_KEY. Si non spécifié, le système de compilation du noyau générera une clé. Il est recommandé de la générer manuellement. Cela peut être fait avec :

`root #` `openssl req -new -nodes -utf8 -sha256 -x509 -outform PEM -out kernel_key.pem -keyout kernel_key.pem`

OpenSSL vous posera quelques questions pour générer la clé, il est recommandé de donner des réponses aussi précises que possible.

Conservez la clé dans un répertoire sûr ; au plus, la clé devrait être lisible par l’utilisateur _root_. Vérifiez avec :

`root #` `ls -l kernel_key.pem`

```
 -r-------- 1 root root 3164 Jan  4 10:38 kernel_key.pem
```

Si la sortie est différente comparé à celle du dessus, corrigez les permissions avec :

`root #` `chown root:root kernel_key.pem
`

`root #` `chmod 400 kernel_key.pem
`

KERNEL **Spécifier la clé pour signer CONFIG\_MODULE\_SIG\_KEY**

```
-*- Cryptographic API  --->
  Certificates for signature checking  --->
    (/path/to/kernel_key.pem) File name or PKCS#11 URI of module signing key

```

Pour signer également des modules noyaux installés par d’autres paquets via `linux-mod-r1.eclass`, acitivez le drapeau _USE_ [modules-sign](https://packages.gentoo.org/useflags/modules-sign) [USE flag/fr](/wiki/USE_flag/fr "USE flag/fr") globalement :

FILE **`/etc/portage/make.conf`** **Activer de la signature des modules**

```
USE="modules-sign"

# Optionnellement, en utilisant des clés de signatures personnalisées
MODULES_SIGN_KEY="/path/to/kernel_key.pem"
MODULES_SIGN_CERT="/path/to/kernel_key.pem" # Nécessaire seulement si « MODULES_SIGN_KEY » ne contient pas également le certificat
MODULES_SIGN_HASH="sha512" # sha512 par défaut

```

**Remarque**

MODULES\_SIGN\_KEY et MODULES\_SIGN\_CERT pourraient pointer sur différents fichiers. Par exemple, le PEM généré par OpenSSL inclus à la fois la clé et le certificat, et ces variables ont la même valeur.

{{#switch: mips \| amd64 \| arm64 \| riscv \| x86 =
Optionnel : Signez l’image noyau ( _Secure Boot_) ====

Lors de la signature d’une image noyau (pour les systèmes avec [Secure Boot](/wiki/Secure_Boot "Secure Boot") activé), il est recommandé d’activer les options suivantes dans la configuration du noyau :

KERNEL **Attachement pour _secure boot_**

```
General setup  --->
  Kexec and crash features  --->
    [*] Enable kexec system call
    [*] Enable kexec file based system call
    [*]   Verify kernel signature during kexec_file_load() syscall
    [*]     Require a valid signature in kexec_file_load() syscall
    [*]     Enable ""image"" signature verification support

[*] Enable loadable module support
  -*-   Module signature verification
    [*]     Require modules to be validly signed
    [*]     Automatically sign all modules
    Which hash algorithm should modules be signed with? (Sign modules with SHA-512) --->

Security options  --->
[*] Integrity subsystem
  [*] Basic module for enforcing kernel lockdown
  [*]   Enable lockdown LSM early in init
        Kernel default lockdown mode (Integrity)  --->

  [*]   Digital signature verification using multiple keyrings
  [*]     Enable asymmetric keys support
  -*-       Require all keys on the integrity keyrings be signed
  [*]       Provide keyring for platform/firmware trusted keys
  [*]       Provide a keyring to which Machine Owner Keys may be added
  [ ]         Enforce Machine Keyring CA Restrictions

```

Où « image » est un paramètre pour le nom spécifique de l’image de l’architecture. Ces options, du haut vers le bas : forcent que l’image du noyau est lancée dans un appel _kexec_ doit être signée ( _kexec_ autorise le remplacement du noyau), forcent que les modules du noyau sont signés, active le verrouillage d’intégrité (pour empêcher la modification du noyau lors de l’exécution) et activent plusieurs chaînes.

Pour les architectures qui ne supportent pas nativement la décompression du noyau (i.e. **arm64** et **riscv**), le noyau doit être compilé avec son propre décompresseur ( _zboot_) :

KERNEL **zboot CONFIG\_EFI\_ZBOOT**

```
Device Drivers --->
  Firmware Drivers --->
    EFI (Extensible Firmware Interface) Support --->
      [*] Enable the generic EFI decompressor

```

Après avoir compilé le noyau, comme expliqué dans la section suivante, l’image du noyau doit être signée. Premièrement, installez [app-crypt/sbsigntools](https://packages.gentoo.org/packages/app-crypt/sbsigntools) et signer l’image du noyau :

`root #` `emerge --ask app-crypt/sbsigntools`

`root #` `sbsign /usr/src/linux-x.y.z/path/to/kernel-image --cert /path/to/kernel_key.pem --key /path/to/kernel_key.pem --output /usr/src/linux-x.y.z/path/to/kernel-image`

**Remarque**

Sur cet exemple, la même clé est utilisée pour signer les modules et l’image du noyau. Il est aussi possible de générer deux clés séparées pour cela. La même commande OpenSSL pourra être encore utilisée.

Continuez avec l’installation.

Pour signer automatiquement les exécutables EFI installés par d’autres paquets, activez globalement le drapeau [secureboot](https://packages.gentoo.org/useflags/secureboot) [USE flag/fr](/wiki/USE_flag/fr "USE flag/fr") :

FILE **`/etc/portage/make.conf`** **Activer le Secure Boot**

```
USE="modules-sign secureboot"

# Optionnellment, pour utiliser des clés de signatures personnalisées
MODULES_SIGN_KEY="/path/to/kernel_key.pem"
MODULES_SIGN_CERT="/path/to/kernel_key.pem" # Seulement nécessaire si le module MODULES_SIGN_KEY ne contient pas le certificat.
MODULES_SIGN_HASH="sha512" # sha512 par défaut

# Optionnellement, pour démarrer avec secureboot activé, avec la même clé ou une différente.
SECUREBOOT_SIGN_KEY="/path/to/kernel_key.pem"
SECUREBOOT_SIGN_CERT="/path/to/kernel_key.pem"

```

**Remarque**

SECUREBOOT\_SIGN\_KEY et SECUREBOOT\_SIGN\_CERT peuvent pointer différents fichiers. Pour cet exemple, le fichier PEM généré par OpenSSL inclus à la fois la clé et le certificat, et chaque variable à la même valeur.

**Remarque**

Lors de la génération d’une [image unifiée du noyau](/wiki/Unified_Kernel_Image "Unified Kernel Image") avec la commande _systemd_ `ukify`, l’image du noyau sera signée automatiquement avant l’inclusion dans l’image unifiée du noyau. Il n’est pas nécessaire de la signer manuellement.

[Handbook:MIPS/Blocks/Kernel/fr](/index.php?title=Handbook:MIPS/Blocks/Kernel/fr&action=edit&redlink=1 "Handbook:MIPS/Blocks/Kernel/fr (page does not exist)")

Continuer l’installation avec [Configurer le système](/wiki/Handbook:MIPS/Installation/System/fr "Handbook:MIPS/Installation/System/fr").

[← Installation du système de base](/wiki/Handbook:MIPS/Installation/Base/fr "Previous part") [Accueil](/wiki/Handbook:MIPS/fr "Handbook:MIPS/fr") [Configurer le Système →](/wiki/Handbook:MIPS/Installation/System/fr "Next part")