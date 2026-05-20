# Bootloader

Other languages:

- [Deutsch](/wiki/Handbook:AMD64/Installation/Bootloader/de "Handbook:AMD64/Installation/Bootloader/de (100% translated)")
- [English](/wiki/Handbook:AMD64/Installation/Bootloader "Handbook:AMD64/Installation/Bootloader (100% translated)")
- [español](/wiki/Handbook:AMD64/Installation/Bootloader/es "Handbook:AMD64/Installation/Bootloader/es (100% translated)")
- français
- [italiano](/wiki/Handbook:AMD64/Installation/Bootloader/it "Handbook:AMD64/Installation/Bootloader/it (100% translated)")
- [magyar](/wiki/Handbook:AMD64/Installation/Bootloader/hu "Handbook:AMD64/Installation/Bootloader/hu (100% translated)")
- [polski](/wiki/Handbook:AMD64/Installation/Bootloader/pl "Handbook:AMD64/Installation/Bootloader/pl (100% translated)")
- [português do Brasil](/wiki/Handbook:AMD64/Installation/Bootloader/pt-br "Handbook:AMD64/Installation/Bootloader/pt-br (100% translated)")
- [čeština](/wiki/Handbook:AMD64/Installation/Bootloader/cs "Handbook:AMD64/Installation/Bootloader/cs (100% translated)")
- [русский](/wiki/Handbook:AMD64/Installation/Bootloader/ru "Handbook:AMD64/Installation/Bootloader/ru (100% translated)")
- [தமிழ்](/wiki/Handbook:AMD64/Installation/Bootloader/ta "Handbook:AMD64/Installation/Bootloader/ta (100% translated)")
- [中文（中国大陆）‎](/wiki/Handbook:AMD64/Installation/Bootloader/zh-cn "Handbook:AMD64/Installation/Bootloader/zh-cn (100% translated)")
- [日本語](/wiki/Handbook:AMD64/Installation/Bootloader/ja "Handbook:AMD64/Installation/Bootloader/ja (100% translated)")
- [한국어](/wiki/Handbook:AMD64/Installation/Bootloader/ko "Handbook:AMD64/Installation/Bootloader/ko (100% translated)")

[Sommaire du manuel](/wiki/Handbook:AMD64 "Handbook:AMD64")[Installation](/wiki/Handbook:AMD64/Full/Installation/fr "Handbook:AMD64/Full/Installation/fr")[À propos de l’installation](/wiki/Handbook:AMD64/Installation/About/fr "Handbook:AMD64/Installation/About/fr")[Choix du support](/wiki/Handbook:AMD64/Installation/Media/fr "Handbook:AMD64/Installation/Media/fr")[Configurer le réseau](/wiki/Handbook:AMD64/Installation/Networking/fr "Handbook:AMD64/Installation/Networking/fr")[Préparer les disques](/wiki/Handbook:AMD64/Installation/Disks/fr "Handbook:AMD64/Installation/Disks/fr")[Installer l’archive _stage3_](/wiki/Handbook:AMD64/Installation/Stage/fr "Handbook:AMD64/Installation/Stage/fr")[Installer le système de base](/wiki/Handbook:AMD64/Installation/Base/fr "Handbook:AMD64/Installation/Base/fr")[Configurer le noyau](/wiki/Handbook:AMD64/Installation/Kernel/fr "Handbook:AMD64/Installation/Kernel/fr")[Configurer le système](/wiki/Handbook:AMD64/Installation/System/fr "Handbook:AMD64/Installation/System/fr")[Installer les outils](/wiki/Handbook:AMD64/Installation/Tools/fr "Handbook:AMD64/Installation/Tools/fr")Configurer le système d’amorçage[Finaliser](/wiki/Handbook:AMD64/Installation/Finalizing/fr "Handbook:AMD64/Installation/Finalizing/fr")[Utiliser Gentoo](/wiki/Handbook:AMD64/Full/Working/fr "Handbook:AMD64/Full/Working/fr")[Introduction à Portage](/wiki/Handbook:AMD64/Working/Portage/fr "Handbook:AMD64/Working/Portage/fr")[Les options de la variable USE](/wiki/Handbook:AMD64/Working/USE/fr "Handbook:AMD64/Working/USE/fr")[Les fonctionnalités de Portage](/wiki/Handbook:AMD64/Working/Features/fr "Handbook:AMD64/Working/Features/fr")[Scripts d’initialisation systèmes](/wiki/Handbook:AMD64/Working/Initscripts/fr "Handbook:AMD64/Working/Initscripts/fr")[Variables d’environnement](/wiki/Handbook:AMD64/Working/EnvVar/fr "Handbook:AMD64/Working/EnvVar/fr")[Utiliser Portage](/wiki/Handbook:AMD64/Full/Portage/fr "Handbook:AMD64/Full/Portage/fr")[Fichiers et répertoires](/wiki/Handbook:AMD64/Portage/Files/fr "Handbook:AMD64/Portage/Files/fr")[Les variables](/wiki/Handbook:AMD64/Portage/Variables/fr "Handbook:AMD64/Portage/Variables/fr")[Mélanger plusieurs branches logicielles](/wiki/Handbook:AMD64/Portage/Branches/fr "Handbook:AMD64/Portage/Branches/fr")[Outils supplémentaires](/wiki/Handbook:AMD64/Portage/Tools/fr "Handbook:AMD64/Portage/Tools/fr")[Dépôt personnalisé](/wiki/Handbook:AMD64/Portage/CustomTree/fr "Handbook:AMD64/Portage/CustomTree/fr")[Fonctionnalités avancées](/wiki/Handbook:AMD64/Portage/Advanced/fr "Handbook:AMD64/Portage/Advanced/fr")[Configuration du réseau avec OpenRC](/wiki/Handbook:AMD64/Full/Networking/fr "Handbook:AMD64/Full/Networking/fr")[Bien démarrer](/wiki/Handbook:AMD64/Networking/Introduction/fr "Handbook:AMD64/Networking/Introduction/fr")[Configuration avancée](/wiki/Handbook:AMD64/Networking/Advanced/fr "Handbook:AMD64/Networking/Advanced/fr")[Les modules réseau](/wiki/Handbook:AMD64/Networking/Modular/fr "Handbook:AMD64/Networking/Modular/fr")[Sans fil](/wiki/Handbook:AMD64/Networking/Wireless/fr "Handbook:AMD64/Networking/Wireless/fr")[Ajouter des fonctionnalités](/wiki/Handbook:AMD64/Networking/Extending/fr "Handbook:AMD64/Networking/Extending/fr")[Gestion dynamique](/wiki/Handbook:AMD64/Networking/Dynamic/fr "Handbook:AMD64/Networking/Dynamic/fr")

## Contents

- [1Choisir un programme d’amorçage](#Choisir_un_programme_d.E2.80.99amor.C3.A7age)
- [2Défaut : GRUB](#D.C3.A9faut_:_GRUB)
  - [2.1Compilation ( _Emerge_)](#Compilation_.28Emerge.29)
  - [2.2Installation](#Installation)
    - [2.2.1Système DOS/BIOS déprécié](#Syst.C3.A8me_DOS.2FBIOS_d.C3.A9pr.C3.A9ci.C3.A9)
    - [2.2.2Système UEFI](#Syst.C3.A8me_UEFI)
      - [2.2.2.1Optionnel : _Secure Boot_](#Optionnel_:_Secure_Boot)
      - [2.2.2.2Déboguer GRUB](#D.C3.A9boguer_GRUB)
  - [2.3Configurer](#Configurer)
- [3Alternative 1 : systemd-boot](#Alternative_1_:_systemd-boot)
  - [3.1Compiler ( _Emerge_)](#Compiler_.28Emerge.29)
  - [3.2Installation](#Installation_2)
  - [3.3Facultatif : _Secure Boot_](#Facultatif_:_Secure_Boot)
- [4Alternative 2 : _EFI Stub_](#Alternative_2_:_EFI_Stub)
  - [4.1Image noyau unifiée](#Image_noyau_unifi.C3.A9e)
- [5Autres alternatives](#Autres_alternatives)
- [6Redémarrer le système](#Red.C3.A9marrer_le_syst.C3.A8me)

## Choisir un programme d’amorçage

Une fois le noyau Linux configuré, les outils systèmes installés et configurés, il est maintenant temps d’installer la dernière pièce importante : un programme d’amorçage ( _bootloader_).

Le programme d’amorçage est responsable du lancement du noyau au démarrage – sans lui, le système ne démarrerait pas lorsque le bouton d’allumage est pressé.

Pour **amd64**, nous documentons comment configurer [GRUB](/wiki/Handbook:AMD64/Blocks/Bootloader/fr#Default:_GRUB "Handbook:AMD64/Blocks/Bootloader/fr") pour les systèmes dépréciés DOS/BIOS and [GRUB](#Default:_GRUB), [systemd-boot](#Alternative_1:_systemd-boot) ou [EFI Stub](#Alternative_2:_efibootmgr) pour ceux UEFI.

Dans cette section du manuel une délimitation est faite entre _compiler_ le programme d’amorçage et _installer_ le programme d’amorçage sur le disque système. Ici, _compiler_ est utilisé pour indiquer à [Portage](/wiki/Portage "Portage") de rendre disponible le paquet dans le système. _Installer_ signifie que le programme d’amorçage copie physiquement les fichiers et modifie les sections appropriées du disque système pour rendre le programme d’amorçage activé et opérationnel lors du prochain démarrage.

## Défaut : GRUB

Par défaut, la majorité des systèmes Gentoo reposent sur [GRUB](/wiki/GRUB "GRUB") (paquet [sys-boot/grub](https://packages.gentoo.org/packages/sys-boot/grub)), lequel est le successeur direct de [GRUB Legacy](/wiki/GRUB_Legacy "GRUB Legacy"). Sans configuration supplémentaire, GRUB supporte avec joie les anciens systèmes PC BIOS. Avec une petite configuration, avant la compilation, GRUB peut supporter une demi-douzaine de plates-formes supplémentaires. Pour plus d’informations, voir les [prérequis](/wiki/GRUB#Prerequisites "GRUB") de l’article [GRUB](/wiki/GRUB "GRUB").

### Compilation ( _Emerge_)

Lors de l’utilisation d’un vieux système BIOS supportant seulement les partitions MBR, aucune configuration supplémentaire n’est nécessaire pour compiler GRUB :

`root #` `emerge --ask --verbose sys-boot/grub`

Note pour les utilisateurs UEFI : lancer la commande ci-dessous va afficher les valeurs de GRUB\_PLATFORMS avant de compiler. Lors de l’utilisation d’UEFI, les utilisateurs doivent s’assurer que `GRUB_PLATFORMS="efi-64"` est activé (c’est le cas par défaut). Sinon, le code `GRUB_PLATFORMS="efi-64"` devra être ajouté dans /etc/portage/make.conf _avant_ de compiler GRUB pour que le paquet soit construit avec la fonctionnalité EFI :

`root #` `echo 'GRUB_PLATFORMS="efi-64"' >> /etc/portage/make.conf`

`root #` `emerge --ask sys-boot/grub`

Si GRUB est compilé dans activer `GRUB_PLATFORMS="efi-64"`, la ligne (ci-dessus) peut être ajoutée à make.conf et les dépendances pour [la collection _world_](/wiki/World_set_(Portage) "World set (Portage)") peut être recaculée en utilisant l’option `--update --newuse` à emerge :

`root #` `emerge --ask --update --newuse --verbose sys-boot/grub`

Le logiciel GRUB est maintenant _compilé_ dans le système, mais il n’est pas encore _installé_ comme programme d’amorçage.

### Installation

Ensuite, installer les fichiers GRUB nécessaires dans /boot/grub/ via la commande grub-install. En présumant que le premier disque (celui où le système se situe) est /dev/sda, une des commandes suivantes sera :

#### Système DOS/BIOS déprécié

Pour les systèmes DOS/BIOS dépréciés :

`root #` `grub-install /dev/sda`

#### Système UEFI

**Important**

Assurez-vous que la partition système EFI est bien montée _avant_ de lancer grub-install. Il est possible pour grub-install d’installer le fichier EFI GRUB (grubx64.efi) dnas le mauvais répertoire **sans indiquer** qu’un mauvais répertoire a été utilisé.

Pour les systèmes UEFI :

`root #` `grub-install --efi-directory=/efi`

```
Installing for x86_64-efi platform.
Installation finished. No error reported.

```

Après l’installation réussie, la sortie devrait correspondre à celle de la commande précédente. Si elle ne correspond pas exactement, consultez [déboguer GRUB](/wiki/Handbook:AMD64/Blocks/Bootloader/fr#Debugging_GRUB "Handbook:AMD64/Blocks/Bootloader/fr") ou sautez à l’ [étape de configuration](/wiki/Handbook:AMD64/Blocks/Bootloader/fr#GRUB_Configure "Handbook:AMD64/Blocks/Bootloader/fr").

##### Optionnel : _Secure Boot_

Pour démarrer avec _Secure Boot_ activé, le certificat de signature doit être accepté par le micrologiciel [UEFI](/wiki/UEFI "UEFI"), ou [shim](/wiki/Shim "Shim") doit être utilisé comme pré-chargeur. Shim est pré-signé par le certificat Microsoft tierce partie, accepté par défaut par la plupart des cartes mères UEFI.

Comment configurer le micrologiciel UEFI à accepter des clés personnalisées dépend du revendeur, ce qui dépasse le périmètre du manuel. Ci-dessous est montré comment configurer shim à la place. Il est considéré que l’utilisateur a déjà suivi les instructions des précédentes sections pour générer une clé pour signer et configurer Portage pour l’utiliser. Si ce n’est pas le cas, retournez d’abord à la section d’ [installation du noyau](/wiki/Handbook:AMD64/Installation/Kernel/fr "Handbook:AMD64/Installation/Kernel/fr").

Le paquet [sys-boot/grub](https://packages.gentoo.org/packages/sys-boot/grub) installe un exécutable pré-compilé et exécutable EFI si le drapeau [secureboot](https://packages.gentoo.org/useflags/secureboot) [USE flag/fr](/wiki/USE_flag/fr "USE flag/fr") est activé. Installer les paquets requis et copier le GRUB autoporteur, Shim et MokManager dans le même répertoire sur la partition système EFI. Par exemple :

`root #` `emerge sys-boot/grub sys-boot/shim sys-boot/mokutil sys-boot/efibootmgr`

`root #` `cp /usr/share/shim/BOOTX64.EFI /efi/EFI/Gentoo/shimx64.efi
`

`root #` `cp /usr/share/shim/mmx64.efi /efi/EFI/Gentoo/mmx64.efi
`

`root #` `cp /usr/lib/grub/grub-x86_64.efi.signed /efi/EFI/Gentoo/grubx64.efi
`

Ensuite, enregistrer la clé dans la MOKlist de Shim, cela nécessite des clés dans le format DER là où sbsign et le système de compilation du noyau attente une clé au format PEM. Dans la section précédente du manuel, un exemple est montré pour générer ce type de clé PEM, elle peut être convertie en format DER :

`root #` `openssl x509 -in /path/to/kernel_key.pem -inform PEM -out /path/to/kernel_key.der -outform DER`

**Remarque**

Le répertoire utilisé doit être celui du fichier PEM contenant le certificat appartenant à la clé générée. Dans cet exemple, le certificat et la clé sont dans le même fichier PEM.

Ensuite, convertir le certificat et l’importer dans la MOKlist de Shim, cette commande va vous demander de définir un mot de passe lors de l’import :

`root #` `mokutil --import /path/to/kernel_key.der`

**Remarque**

Lorsque le noyau couramment démarré fait confiance au certificat en train d’être importé, le message « Already in kernel trusted keyring. » sera affiché. Si cela arrive, relancer la commande suivante avec l’option --ignore-keyring.

Ensuite, enregistrer Shim avec le micrologiciel UEFI. Dans la commande suivante, `boot-disk` et `boot-partition-id` doivent être remplacé avec le disque et l’identifiant de la partition EFI :

`root #` `efibootmgr --create --disk /dev/boot-disk --part boot-partition-id --loader '\EFI\Gentoo\shimx64.efi' --label 'GRUB via Shim' --unicode`

Notez que cela pré-compile et signe une version autoportante de GRUB lisant grub.cfg depuis un répertoire inhabituel. À la place de /boot/grub/grub.cfg, le fichier de configuration devrait être dans le même répertoire que l’exécutable EFI, cgest-à-dire /efi/EFI/Gentoo/grub.cfg. Lorsque [sys-kernel/installkernel](https://packages.gentoo.org/packages/sys-kernel/installkernel) est utilisé pour installer le noyau, les mises à jour et la configuration GRUB, la variable d’enviromment GRUB\_CFG doit être utilisée pour surcharger le répertoire habituel du fichier de configuration GRUB.

Par exemple :

`root #` `grub-mkconfig -o /efi/EFI/Gentoo/grub.cfg`

Ou : [installkernel](/wiki/Installkernel/fr "Installkernel/fr"):

FILE **`/etc/env.d/99grub`**

```
GRUB_CFG=/efi/EFI/Gentoo/grub.cfg

```

`root #` `env-update`

**Remarque**

La procédure d’import ne sera pas terminée tant que le système ne sera pas redémarré. Après avoir accompli toutes les étapes du manuel, redémarrer le système et Shim trouvera l’import enregistré par mokutil. Le MokManager va démarrer et demander le mot de passe utilisé lors de la demande d’import. Suivez les instructions pour terminer l’import du certificat, puis redémarrer le système, aller dans le menu UEFI et activez le paramètre _Secure Boot_.

##### Déboguer GRUB

Lors du débogage de GRUB, plusieurs petites corrections peuvent résoudre un problème sans avoir à démarrer depuis une image _live_.

Dans le cas où « EFI variables are not supported on this system » est affiché en sortie, il est possible que l’image _live_ n’est pas lancé en mode EFI et est en mode BIOS déprécié. La solution est d’essayer [removable GRUB step](/wiki/Handbook:AMD64/Blocks/Bootloader/fr#GRUB_Install_EFI_systems_removable "Handbook:AMD64/Blocks/Bootloader/fr") ci-dessous. Cela va surchargeer l’exécutable EFI localisé dans /EFI/BOOT/BOOTX64.EFI. Lors du redémarrage en mode EFI, le micrologiciel de la carte mère peut lancer l’entrée de démarrage par défaut et lancer GRUB.

Si grub-install retourne l’erreur « Could not prepare Boot variable: Read-only file system » et que l’environnement _live_ est bien lancé en mode UEFI, alors il devrait être possible de remonter en lecture-écriture la partition EFI et recommencer [la commande grub-install susnommée](/wiki/Handbook:AMD64/Blocks/Bootloader/fr#GRUB_Install_EFI_systems_command "Handbook:AMD64/Blocks/Bootloader/fr") :

`root #` `mount -o remount,rw,nosuid,nodev,noexec --types efivarfs efivarfs /sys/firmware/efi/efivars`

C’est dû à certain environnement Gentoo non officiel pas monté avec le système de fichiers EFI par défaut. Si la commande précédente ne se lance pas, alors redémarrer avec un support officiel Gentoo en mode EFI.

Certaines cartes mères ont une implémentations UEFI incomplète et supportent **seulement** le répertoire /EFI/BOOT pour le fichier EFI dans la partition système EFI ( _ESP_). L’installateur GRUB peut créer le fichier dans ce répertoire automatiquement en ajoutant l’option `--removable`. Assurez-vous que l’ _ESP_ est monté avant de lancer la commande ; en présumant qu’elle est monté sur /efi (comme défini plus haut), lancez :

`root #` `grub-install --target=x86_64-efi --efi-directory=/efi --removable`

Cela créera le répertoire par défaut défini dans les spécifications UEFI et un fichier avec le nom par défaut : BOOTX64.EFI.

### Configurer

Ensuite, générer la configuration GRUB basée sur celle de l’utilisateur spécifiée dans le fichier /etc/default/grub et les scripts de /etc/grub.d. Dans la plupart des cas, aucune configuration est nécessaire par les utilisateurs car GRUB détecte automatiquement quel noyau démarrer (la version la plus élevée dans /boot/) et quel est le système de fichiers racine. Il est possible d’ajouter des paramètres au noyau dans /etc/default/grub en renseignant la variable GRUB\_CMDLINE\_LINUX.

Pour générer la configuration finale de GRUB, lancer grub-mkconfig :

`root #` `grub-mkconfig -o /boot/grub/grub.cfg`

```
Generating grub.cfg ...
Found linux image: /boot/vmlinuz-6.18.8-gentoo
Found initrd image: /boot/initramfs-genkernel-amd64-6.18.8-gentoo
done

```

La sortie de la commande doit mentionner qu’au moins une image Linux est trouvée, car ceci est nécessaire pour démarrer le système. Si un _initramfs_ est utilisé ou genkernel a été utilisé pour compiler le noyau, le bon _initrd_ devrait être détecté. Si ce n’est pas le cas, aller dans /boot/ et vérifier son contenu avec ls. Si les fichiers viennent à manquer, retournez dans la configuration du noyau et les instructions d’installation.

**Conseil**

L’utilitaire os-prober peut être utilisé en parallèle avec GRUB pour détecter d’autres systèmes d’exploitation sur les disques attachés. Windows 7, 8.1, 10 et d’autres distributions Linux sont détectables. Ceux qui désirent une double-démarrage ( _dual boot_) pourraient compiler le paquet [sys-boot/os-prober](https://packages.gentoo.org/packages/sys-boot/os-prober) et relancer grub-mkconfig comme précédemment. Si des problèmes de détections arrivent, lisez l’article [GRUB](/wiki/GRUB "GRUB") dans son entièreté **avant** de demander un support à la communauté Gentoo.

## Alternative 1 : systemd-boot

Une autre option est [systemd-boot](/wiki/Systemd/systemd-boot "Systemd/systemd-boot"), lequel fonctionne aussi bien sur des machines OpenRC ou systemd. C’est une chaîne de lancement légère qui fonctionne aussi avec _Secure Boot_.

### Compiler ( _Emerge_)

Pour installer systemd-boot, activer le drapeau [boot](https://packages.gentoo.org/useflags/boot) [USE flag/fr](/wiki/USE_flag/fr "USE flag/fr") et réinstaller [sys-apps/systemd](https://packages.gentoo.org/packages/sys-apps/systemd) (pour les système systemd) ou [sys-apps/systemd-utils](https://packages.gentoo.org/packages/sys-apps/systemd-utils) (pour les systèmes OpenRC) :

FILE **`/etc/portage/package.use/systemd-boot`**

```
sys-apps/systemd boot
sys-apps/systemd-utils boot

```

`root #` `emerge --ask sys-apps/systemd`

Ou

`root #` `emerge --ask sys-apps/systemd-utils`

### Installation

Maintenant, installer le chargeur systemd-boot sur la _[EFI System Partition](/wiki/EFI_System_Partition "EFI System Partition")_ :

`root #` `bootctl install`

**Important**

Assurez-vous que la partition système EFI est montée **avant** de lancer bootctl install.

Lorsque vous utilisez ce programme d’amorçage, avant de redémarrer, vérifiez que l’entrée de démarrage existe avec :

`root #` `bootctl list`

**Attention !**

La ligne de commandes du noyau pour les nouvelles entrées systemd-boot sont lues depuis /etc/kernel/cmdline ou /usr/lib/kernel/cmdline. Si aucun fichier n’est présent, alors la ligne de commande du noyau actuellement lancé est réutilisé (/proc/cmdline). Dans les nouvelles installations, il est possible pour cela que la ligne de commande du noyau du _live CD_ soit accidentellement utilisée pour démarrer le nouveau noyau. La ligne de commande des entrées enregistrées du noyau peuvent être vérifiées avec :

`root #` `bootctl list`

. Si cela ne montre pas le noyau désiré, alors créer /etc/kernel/cmdline qui contient la bonne ligne de commande ud noyau et réinstaller le noyau.

Si aucune nouvelle entrée n’existe, assurez-vous que [sys-kernel/installkernel](https://packages.gentoo.org/packages/sys-kernel/installkernel) est bien installé avec les drapeaux [systemd](https://packages.gentoo.org/useflags/systemd) [USE flag/fr](/wiki/USE_flag/fr "USE flag/fr") et [systemd-boot](https://packages.gentoo.org/useflags/systemd-boot) [USE flag/fr](/wiki/USE_flag/fr "USE flag/fr") activés et relancer l’installation du noyau.

Pour les noyaux distribués :

`root #` `emerge --ask --config sys-kernel/gentoo-kernel`

Pour un noyau manuellement configuré et compilé :

`root #` `make install`

**Important**

Lors de l’installation de noyaux pour systemd-boot, aucun argument de ligne de commandes root= n’est ajouté par défaut. Dans les systèmes systemd qui utilisent un _initramfs_, les utilisateurs peuvent s’appuyer sur [systemd-gpt-auto-generator](/wiki/Systemd#Automatic_mounting_of_partitions_at_boot.2Ffr "Systemd") pour trouver automatiquement la partition racine au démarrage. Autrement, les utilisateurs peuvent indiquer manuellement l’emplacement de la partition racine en paramétrant root= dans /etc/kernel/cmdline ainsi que d’autres paramètres de ligne de commandes du noyau qui pourraient être utilisés. Et réinstaller le noyau comme décrit ci-dessus.

### Facultatif : _Secure Boot_

Lorsque le drapeau [secureboot](https://packages.gentoo.org/useflags/secureboot) [USE flag/fr](/wiki/USE_flag/fr "USE flag/fr") est activé, l’exécutable EFI systemd-boot sera signé automatiquement par Portage. En outre, bootctl install va automatiquement installer la version signée.

Pour démarrer brillamment avec _Secure Boot_ activé, le certificat doit soit être accepté par le micrologiciel [UEFI](/wiki/UEFI "UEFI") soit [shim](/wiki/Shim "Shim") doit être utilisé comme un pré-chargeur. Shim est pré-signé avec le certificat Microsoft tierce partie, accepté par défaut par la plupart des cartes mères UEFI.

Comment configurer un micrologiciel UEFI pour accepter les clés personnalisées dépend du revendeur de micrologiciel, ce qui sort du périmètre de ce manuel. Ci-dessous est indiqué comment paramétrer Shim à la place. Il est assumé ici que l’utilisateur a déjà suivi les instructions des sections précédentes pour générer une clé de signature et configuré Portage pour l’utiliser. Si ce n’est pas le cas, retournez d’abord dans la section [Configurer le noyau](/wiki/Handbook:AMD64/Installation/Kernel/fr "Handbook:AMD64/Installation/Kernel/fr").

`root #` `emerge --ask sys-boot/shim sys-boot/mokutil sys-boot/efibootmgr`

`root #` `bootctl install --no-variables`

`root #` `cp /usr/share/shim/BOOTX64.EFI /efi/EFI/systemd/shimx64.efi`

`root #` `cp /usr/share/shim/mmx64.efi /efi/EFI/systemd/mmx64.efi`

La MOKlist de Shim nécessite une clé dans le format DER, tandis que sbsign et le système de compilation du noyau attend une clé dans le format PEM. Dans la précédente section du manuel, un exemple était montré pour générer ce type de clé de signature PEM, cette clé peut maintenant être convertie dans le format DER :

`root #` `openssl x509 -in /path/to/kernel_key.pem -inform PEM -out /path/to/kernel_key.der -outform DER`

**Remarque**

Le répertoire utilisé ici doit être celui du fichier PEM contenant le certificat lié à la clé générée. Dans cet exemple, la clé et le certificat sont dans le même fichier PEM.

Ensuite, le certificat converti peut être importé dans la MOKlist de Shim :

`root #` `mokutil --import /path/to/kernel_key.der`

**Remarque**

Lorsque le noyau démarré reconnaît déjà le certificat importé, le message « Already in kernel trusted keyring. » sera affiché ici. Si cela survient, relancer la commande ci-dessus avec le paramètre --ignore-keyring.

Finalement, nous enregistrons Shim avec le micrologiciel UEFI. Dans la commande suivante, `boot-disk` et `boot-partition-id` doivent être remplacée par l’identifiant disque et de partition de l’ _ESP_ :

`root #` `efibootmgr --create --disk /dev/boot-disk --part boot-partition-id --loader '\EFI\systemd\shimx64.efi'  --label 'Systemd-boot via Shim' --unicode '\EFI\systemd\systemd-bootx64.efi'`

**Remarque**

Le processus d’import ne sera pas terminé jusqu’au redémarrage du système. Aorès avoir terminé toutes les étapes du manuel, redémarrer le système et Shim va se charger, trouver les demandes d’import faites par mokutil. L’application MokManager va démarrer et demander le mot de passe utilisé pour créer la demande d’import. Suivez les instructions de l’écran pour terminer l’import du certificat ; ensuite redémarrer le système et dans le menu UEFI, activez le paramètre _Secure Boot_.

## Alternative 2 : _EFI Stub_

Les ordinateurs basés sur les micrologiciels UEFI n’ont pas besoin de second programme d’amorçage (ie. GRUB) pour démarrer des noyaux. Secondement, les programmes d’amorçage existent pour _étendre_ les fonctionnalités des micrologiciels UEFI durant le processus de démarrage. Utiliser GRUB (voir la section précédente) est typiquement plus facile et fiable car cela offre une approche plus flexible pour des modifications rapides des paramètres du noyau au démarrage.

Les administrateurs systèmes qui désirent une approchent minimaliste, bien que plus rigide, peuvent éviter un second programme d’amorçage et démarrer le noyau Linux comme un _[EFI Stub](/wiki/EFI_Stub "EFI Stub")_.

L’application [sys-boot/efibootmgr](https://packages.gentoo.org/packages/sys-boot/efibootmgr) est un outil utilisé pour interagir avec un micrologiciel UEFI – le système d’amorçage principal. Normalement, cela se fait en ajoutant ou supprimant des entrées dans la liste du micrologiciel. Cela peut aussi mettre à jour les paramètres du micrologiciel de manière que le noyau Linux précédemment ajouté sera lancé avec les options supplémentaires. Cette interaction sera effectuée à travers une structure de données spéciale appelée variables EFI (d’où le besoin des variables EFI pour le support du noyau).

Assurez-vous de relire l’article [EFI stub](/wiki/EFI_stub "EFI stub") _avant_ de continuer. Le noyau devra avoir les options spécifiques activées pour être directement démarrable par le micrologiciel UEFI. Il peut être nécessaire de recompiler le noyau dans le but d’ajouter ce support.

C’est aussi une bonne idée de regarder l’article [efibootmgr](/wiki/Efibootmgr "Efibootmgr") pour plus d’informations.

**Remarque**

Pour répéter, efibootmgr **n’est pas** une méthode recommandée pour démarrer un sytème UEFI ; il est simplement nécessaire d’ajouter une entrée pour un noyau _EFI Stub_ dans le micrologiciel UEFI. Lorsque compilé avec le support approprié _EFI Stub_, le noyau Linux lui-même peut démarrer **directement**. Les options supplémentazires à la ligne de commandes du noyau peut être intégrée dans le noyau Linux (il y a une configuration noyau appelée CONFIG\_CMDLINE). De manière similaire, le support _initramfs_ peut être intégré dans le noyau également.

Pour démarrer le noyau directement depuis le micrologiciel, l’image noyau peut être présent sur une [_EFI System Partition_](/index.php?title=%27%27EFI_System_Partition%27%27&action=edit&redlink=1 "''EFI System Partition'' (page does not exist)"). Ceci peut se faire en activant le drapeau [efistub](https://packages.gentoo.org/useflags/efistub) [USE flag/fr](/wiki/USE_flag/fr "USE flag/fr") de [sys-kernel/installkernel](https://packages.gentoo.org/packages/sys-kernel/installkernel). En considérant que le démarrage de l’ _EFI Stub_ n’est pas garanti de fonctionner sur tous les systèmes UEFI, le drapeau _USE_ est masqué, les mots-clés de test ( _testing keywords_) doivent être acceptés pour qu’ _installkernel_ utilise cette fonctionnalité.

FILE **`/etc/portage/package.accept_keywords/installkernel`**

```
sys-kernel/installkernel
sys-boot/uefi-mkconfig
app-emulation/virt-firmware

```

FILE **`/etc/portage/package.use/installkernel`**

```
sys-kernel/installkernel efistub

```

Ensuite, réinstaller [installkernel](/wiki/Installkernel/fr "Installkernel/fr"), créer le répertoire /efi/EFI/Gentoo et réinstaller le noyau :

`root #` `emerge --ask sys-kernel/installkernel`

`root #` `mkdir -p /efi/EFI/Gentoo`

Pour les noyaux distribués :

`root #` `emerge --ask --config sys-kernel/gentoo-kernel{,-bin}`

Pour les noyaux gérés manuellement :

`root #` `make install`

Alternativement, si [installkernel](/wiki/Installkernel/fr "Installkernel/fr") n’est pas utilisé, le noyau peut copier être copié manuellement vers la partition système EFI, appelons-la bzImage.efi :

`root #` `mkdir -p /efi/EFI/Gentoo
`

`root #` `cp /boot/vmlinuz-* /efi/EFI/Gentoo/bzImage.efi
`

Installer le paquet efibootmgr :

`root #` `emerge --ask sys-boot/efibootmgr`

Créer l’entrée de démarrage appelée « Gentoo » pour un _EFI Strub_ fraîchement compilé avec le micrologiciel UEFI :

`root #` `efibootmgr --create --disk /dev/sda --part 1 --label "gentoo" --loader "\EFI\Gentoo\bzImage.efi"`

**Remarque**

L’usage de la barre oblique inversée ( `\`) comme séparateur de répertoire est obligatoire pour déclarer les définitions UEFI.

Si un _initramfs_ est utilisé, copiez-le dans la partition système EFI, ensuite ajouter les options de démarrage dessus :

`root #` `efibootmgr --create --disk /dev/sda --part 1 --label "gentoo" --loader "\EFI\Gentoo\bzImage.efi" --unicode "initrd=\EFI\Gentoo\initramfs.img"`

**Conseil**

Les options supplémentaires de commandes de noyau peuvent être lues par le micrologiciel et transmise au noyau avec l’option initrd=… comme montré ci-dessus.

Avec ces changements effectués, lorsque le système redémarre, une entrée appelée « gentoo » sera disponible.

### Image noyau unifiée

Si [installkernel](/wiki/Installkernel/fr "Installkernel/fr") est configuré pour compiler et installer une image de noyau unifiée, elle devrait déjà être installée dans le répertoire EFI/Linux de la partition système EDI. Si ce n’est pas le cas, assurez-vous que le répertoire existe et lancer l’installation du noyau à nouveau comme décrit plus tôt dans le manuel.

Pour ajouter une entrée directe pour une image noyau unifiée installée :

`root #` `efibootmgr --create --disk /dev/sda --part 1 --label "gentoo" --loader "\EFI\Linux\gentoo-x.y.z.efi"`

## Autres alternatives

Pour les autres options qui ne sont pas couvertes dans le manuel, regarder la liste complète du programme d’amorçage [disponible](/wiki/Bootloader "Bootloader").

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

Une fois redémarré dans le nouvel environnement Gentoo, il est avisé de finir avec [Finalisation de l’installation de Gentoo](/wiki/Handbook:AMD64/Installation/Finalizing/fr "Handbook:AMD64/Installation/Finalizing/fr").

[← Installing tools](/wiki/Handbook:AMD64/Installation/Tools/fr "Previous part") [Accueil](/wiki/Handbook:AMD64/fr "Handbook:AMD64/fr") [Finalisation →](/wiki/Handbook:AMD64/Installation/Finalizing/fr "Next part")