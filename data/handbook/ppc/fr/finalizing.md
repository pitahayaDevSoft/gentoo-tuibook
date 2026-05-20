# Finalizing

Other languages:

- [Deutsch](/wiki/Handbook:PPC/Installation/Finalizing/de "Handbuch:PPC/Installation/Abschluss (100% translated)")
- [English](/wiki/Handbook:PPC/Installation/Finalizing "Handbook:PPC/Installation/Finalizing (100% translated)")
- [español](/wiki/Handbook:PPC/Installation/Finalizing/es "Manual de Gentoo: PPC/Instalación/Finalizar (100% translated)")
- français
- [italiano](/wiki/Handbook:PPC/Installation/Finalizing/it "Handbook:PPC/Installation/Finalizing/it (100% translated)")
- [magyar](/wiki/Handbook:PPC/Installation/Finalizing/hu "Handbook:PPC/Installation/Finalizing/hu (100% translated)")
- [polski](/wiki/Handbook:PPC/Installation/Finalizing/pl "Handbook:PPC/Installation/Finalizing (100% translated)")
- [português do Brasil](/wiki/Handbook:PPC/Installation/Finalizing/pt-br "Handbook:PPC/Installation/Finalizing/pt-br (50% translated)")
- [русский](/wiki/Handbook:PPC/Installation/Finalizing/ru "Handbook:PPC/Installation/Finalizing (100% translated)")
- [தமிழ்](/wiki/Handbook:PPC/Installation/Finalizing/ta "கையேடு:PPC/நிறுவல்/முடித்தல் (100% translated)")
- [中文（中国大陆）‎](/wiki/Handbook:PPC/Installation/Finalizing/zh-cn "手册：PPC/安装/收尾安装工作 (100% translated)")
- [日本語](/wiki/Handbook:PPC/Installation/Finalizing/ja "ハンドブック:PPC/インストール/ファイナライズ (100% translated)")
- [한국어](/wiki/Handbook:PPC/Installation/Finalizing/ko "Handbook:PPC/Installation/Finalizing/ko (100% translated)")

[Sommaire du manuel](/wiki/Handbook:PPC "Handbook:PPC")[Installation](/wiki/Handbook:PPC/Full/Installation/fr "Handbook:PPC/Full/Installation/fr")[À propos de l’installation](/wiki/Handbook:PPC/Installation/About/fr "Handbook:PPC/Installation/About/fr")[Choix du support](/wiki/Handbook:PPC/Installation/Media/fr "Handbook:PPC/Installation/Media/fr")[Configurer le réseau](/wiki/Handbook:PPC/Installation/Networking/fr "Handbook:PPC/Installation/Networking/fr")[Préparer les disques](/wiki/Handbook:PPC/Installation/Disks/fr "Handbook:PPC/Installation/Disks/fr")[Installer l’archive _stage3_](/wiki/Handbook:PPC/Installation/Stage/fr "Handbook:PPC/Installation/Stage/fr")[Installer le système de base](/wiki/Handbook:PPC/Installation/Base/fr "Handbook:PPC/Installation/Base/fr")[Configurer le noyau](/wiki/Handbook:PPC/Installation/Kernel/fr "Handbook:PPC/Installation/Kernel/fr")[Configurer le système](/wiki/Handbook:PPC/Installation/System/fr "Handbook:PPC/Installation/System/fr")[Installer les outils](/wiki/Handbook:PPC/Installation/Tools/fr "Handbook:PPC/Installation/Tools/fr")[Configurer le système d’amorçage](/wiki/Handbook:PPC/Installation/Bootloader/fr "Handbook:PPC/Installation/Bootloader/fr")Finaliser[Utiliser Gentoo](/wiki/Handbook:PPC/Full/Working/fr "Handbook:PPC/Full/Working/fr")[Introduction à Portage](/wiki/Handbook:PPC/Working/Portage/fr "Handbook:PPC/Working/Portage/fr")[Les options de la variable USE](/wiki/Handbook:PPC/Working/USE/fr "Handbook:PPC/Working/USE/fr")[Les fonctionnalités de Portage](/wiki/Handbook:PPC/Working/Features/fr "Handbook:PPC/Working/Features/fr")[Scripts d’initialisation systèmes](/wiki/Handbook:PPC/Working/Initscripts/fr "Handbook:PPC/Working/Initscripts/fr")[Variables d’environnement](/wiki/Handbook:PPC/Working/EnvVar/fr "Handbook:PPC/Working/EnvVar/fr")[Utiliser Portage](/wiki/Handbook:PPC/Full/Portage/fr "Handbook:PPC/Full/Portage/fr")[Fichiers et répertoires](/wiki/Handbook:PPC/Portage/Files/fr "Handbook:PPC/Portage/Files/fr")[Les variables](/wiki/Handbook:PPC/Portage/Variables/fr "Handbook:PPC/Portage/Variables/fr")[Mélanger plusieurs branches logicielles](/wiki/Handbook:PPC/Portage/Branches/fr "Handbook:PPC/Portage/Branches/fr")[Outils supplémentaires](/wiki/Handbook:PPC/Portage/Tools/fr "Handbook:PPC/Portage/Tools/fr")[Dépôt personnalisé](/wiki/Handbook:PPC/Portage/CustomTree/fr "Handbook:PPC/Portage/CustomTree/fr")[Fonctionnalités avancées](/wiki/Handbook:PPC/Portage/Advanced/fr "Handbook:PPC/Portage/Advanced/fr")[Configuration du réseau avec OpenRC](/wiki/Handbook:PPC/Full/Networking/fr "Handbook:PPC/Full/Networking/fr")[Bien démarrer](/wiki/Handbook:PPC/Networking/Introduction/fr "Handbook:PPC/Networking/Introduction/fr")[Configuration avancée](/wiki/Handbook:PPC/Networking/Advanced/fr "Handbook:PPC/Networking/Advanced/fr")[Les modules réseau](/wiki/Handbook:PPC/Networking/Modular/fr "Handbook:PPC/Networking/Modular/fr")[Sans fil](/wiki/Handbook:PPC/Networking/Wireless/fr "Handbook:PPC/Networking/Wireless/fr")[Ajouter des fonctionnalités](/wiki/Handbook:PPC/Networking/Extending/fr "Handbook:PPC/Networking/Extending/fr")[Gestion dynamique](/wiki/Handbook:PPC/Networking/Dynamic/fr "Handbook:PPC/Networking/Dynamic/fr")

## Contents

- [1Gestion des utilisateurs](#Gestion_des_utilisateurs)
  - [1.1Ajouter un utilisateur pour un usage quotidien](#Ajouter_un_utilisateur_pour_un_usage_quotidien)
  - [1.2Élévation temporaire de privilèges](#.C3.89l.C3.A9vation_temporaire_de_privil.C3.A8ges)
  - [1.3Désactiver la connexion _root_](#D.C3.A9sactiver_la_connexion_root)
- [2Nettoyage du disque](#Nettoyage_du_disque)
  - [2.1Suppression des archives](#Suppression_des_archives)
- [3Que faire à présent ?](#Que_faire_.C3.A0_pr.C3.A9sent_.3F)
  - [3.1Documentation supplémentaire](#Documentation_suppl.C3.A9mentaire)
  - [3.2Gentoo sur internet](#Gentoo_sur_internet)
    - [3.2.1Forums et _IRC_](#Forums_et_IRC)
    - [3.2.2Listes de diffusion](#Listes_de_diffusion)
    - [3.2.3Bogues](#Bogues)
    - [3.2.4Guide du développeur](#Guide_du_d.C3.A9veloppeur)
- [4Pensées finales](#Pens.C3.A9es_finales)

## Gestion des utilisateurs

### Ajouter un utilisateur pour un usage quotidien

Travailler en tant que _root_ sur un système Unix/Linux est dangereux et doit être évité autant que possible. Par conséquent, il est fortement recommandé d’ajouter un ou plusieurs comptes utilisateurs pour une utilisation quotidienne.

Les groupes auxquels appartient l’utilisateur définissent quelles activités ce dernier peut effectuer. Le tableau suivant liste un certain nombre de groupes importants :

Group
Description
audio
Active l’accès aux périphériques audios.
cdrom
Active l’accès aux lecteurs optiques.
cron
Active l’accès à la planification de tâche _cron_ avec cron. Note : les utilisateurs de systemd sur leur système peuvent utiliser les _timers_ et un fichier de service utilisateur plutôt que des _crons_.
floppy
Active l’accès à d’anciens périphériques mécaniques comme les disquettes. Ce groupe n’est généralement pas utilisé sur des systèmes modernes.
usb
Active l’accès à des périphériques USB.
video
Active l’accès à la capture vidéo et l’accélération matérielle.
wheel
Autorise le compte utilisateur à utiliser la commande su ( _substitute user_), laquelle permet de passer d’un compte _root_ à un autre. Pour un système avec un seul utilisateur, il est bon d’ajouter ce groupe au premier compte standard.

Par exemple, pour créer un utilisateur appelé [larry](/wiki/User:Larry "User:Larry") qui est membre des groupes « _wheel_ », « _users_ », et « _audio_ », se connecter d’abord en tant que _root_ (seul _root_ peut créer de nouveaux utilisateurs) puis exécuter useradd :

`Login:` `root`

```
Password: (Entrer le mot de passe root)

```

Lors du paramétrage du mot de passe pour une compte utilisateur standard, il est de bonne pratique en matière de sécurité d’éviter d’utiliser un mot de passe similaire au compte _root_.

Les auteurs du manuel recommandent d’utiliser un mot de passe d’au moins 16 caractères et différent pour chaque utilisateurs du système.

`root #` `useradd -m -G users,wheel,audio -s /bin/bash larry
`

`root #` `passwd larry`

```
Password: (Entrer le mot de passe pour larry)
Re-enter password: (Confirmer le mot de passe)

```

### Élévation temporaire de privilèges

Si jamais un utilisateur a besoin d’effectuer une opération en tant que _root_, il peut utiliser su - pour recevoir temporairement les privilèges de _root_. Un autre moyen est d’utiliser les utilitaires [sudo](/wiki/Sudo/fr "Sudo/fr") ([app-admin/sudo](https://packages.gentoo.org/packages/app-admin/sudo)) ou [doas](/wiki/Doas "Doas") ([app-admin/doas](https://packages.gentoo.org/packages/app-admin/doas)) qui, si configurés correctement, sont très sécurisés.

### Désactiver la connexion _root_

**Attention !**

Avant de désactiver la connexion _root_, assurez-vous que le compte utilisateur est membre du groupe « _wheel_ » et que la commande d’élévation de de privilège fonctionne. Sinon, l’accès _root_ sera bloqué et l’administrateur du système deviendra impossible sans effectuer une procédure de réparation. Quelques méthodes courantes pour élever les privilèges d’une utilisateur incluent : [app-admin/sudo](https://packages.gentoo.org/packages/app-admin/sudo), [app-admin/doas](https://packages.gentoo.org/packages/app-admin/doas) ou _run0_ de _systemd_.

Pour prévenir des acteurs extérieurs de se connecter en _root_, supprimer le mot de passe _root_ ou désactiver ce compte peuvent améliorer la sécurité.

Pour désactiver le compte _root_ :

`root #` `passwd -l root`

Pour supprimer le mot de passe _root_ et désactiver le compte :

`root #` `passwd -dl root`

## Nettoyage du disque

### Suppression des archives

Une fois l’installation de Gentoo terminée et le système redémarré, si tout s’est bien passé, l’archive _stage_ et d’autres fichiers – comme DIGEST, CONTENT or \*.asc (PGP signature) – peuvent être supprimé en toute sécurité.

Les fichiers localisés dans le répertoire / peuvent être effacés avec la commande suivante :

`root #` `rm /stage3-*.tar.*`

## Que faire à présent ?

Pas sûr d’où continuer à présent ? Il y a plein de voies à explorer… Gentoo propose à ses utilisateurs moult possibilités et comportent plein de fonctionnalités plus (ou moins) bien documentées sur le wiki and les autres sous-domaine relatif à Gentoo (lisez [Gentoo en ligne](/wiki/Handbook:PPC/Installation/Finalizing/fr#Gentoo_online "Handbook:PPC/Installation/Finalizing/fr") ci-dessous).

### Documentation supplémentaire

Il est important de noter que, de part le nombre de choix disponibles sur Gentoo, la documentation proposée par le manuel porte sur un périmètre limité. Il se concentre sur les basiques pour avoir un système fonctionnel pour des activités basiques. Le manuel exclue volontairement les instructions pour les environnements graphiques, des détails sur le renforcement de la sécurisation et d’autres tâches administratives importantes. Ceci étant dit, il y a davantage de sections dans le manuel pour aider le lecteur avec les autres fonctions basiques.

Les lecteurs devraient envisager sérieusement de regarder la prochaine partie du manuel intitulée [Travailler avec Gentoo](/wiki/Handbook:PPC/Working/Portage/fr "Handbook:PPC/Working/Portage/fr") qui explique comment maintenant ce logiciel à jour, installer des paquets supplémentaires, le détail des drapeaux _USE_, le système d’initialisation _OpenRC_ et bien d’autres informations relatives à la gestion d’un système Gentoo post-installation.

Outre le fait de lire ce manuel, il est recommandé d’explorer d’autres coins du wiki Gentoo afin de trouver une documentation supplémentaire proposée par la communauté. L’équipe du wiki Gentoo offre également un [aperçu de la documentation](/wiki/Main_Page#Documentation_topics.2Ffr "Main Page") répertoriant une sélection des articles wikis. Par exemple, il référence le [guide des paramètres régionaux](/wiki/Localization/Guide/fr "Localization/Guide/fr") pour faire en sorte qu’un système se sente comme à la maison (particulièrement utile si l’anglais n’est pas votre langue maternelle).

La majorité des utilisateurs de bureau vont paramétrer un environnement graphique natif Il y a pléthore d’article « méta » pour les  [environnements de bureau supportés](/wiki/Desktop_environment/fr "Desktop environment/fr") et les [gestionnaires de fenêtres](/wiki/Window_manager "Window manager"). Les lecteurs sont avertis que chaque environnement de bureau à une mise en place légèrement différente, qui allonge et complexifie l’opération.

Plein d’autres [articles métas](/wiki/Category:Meta "Category:Meta") existent pour fournir aux lecteurs une vue macro des logiciels disponibles dans Gentoo.

### Gentoo sur internet

**Important**

Les lecteurs noteront que tous les sites officiels Gentoo sont régis par le [code de conduite](/wiki/Project:Council/Code_of_conduct "Project:Council/Code of conduct") Gentoo. Être actif dans la communauté Gentoo est un privilège, non un droit et chaque utilisateur doit être conscient que le code de conduite existe pour une bonne raison.

À l’exception de Libera.Chat qui héberge les salons de discussion _IRC_ ( _Internet Relay Chat_) et les listes de diffusion, la plupart des sites internet Gentoo nécessitent un compte pour poser de questions, démarrer une discussion ou saisir un bogue.

#### Forums et _IRC_

Chaque utilisateur est bien toujours le bienvenu sur nos [forums Gentoo](https://forums.gentoo.org) ou l’un de nos nombreux [canaux salon de discussion Gentoo](https://www.gentoo.org/get-involved/irc-channels/). Il est facile de chercher dans le forum si un problème rencontré lors de l’installation de Gentoo n’a pas déjà été rencontré et résolu. La probabilité que d’autres utilisateurs aient rencontré une anomalie similaire est importante. Il est nécessaire de rechercher sur les forums, le wiki et les bogues avant de demander une assistance sur les salons Gentoo.

#### Listes de diffusion

[Moult listes de diffusion](https://www.gentoo.org/get-involved/mailing-lists/) sont à la disposition des membres de la communauté qui préfèrent poser une question par courriel plutôt que de créer un compte sur le forum ou _IRC_. Les utilisateurs devront suivre les instructions pour s’abonner à une liste de diffusion spécifique.

#### Bogues

Parfois, après avoir relu le wiki, rechercher dans les forums, _IRC_ et les listes de diffusions, il n’y a pas de solution à un problème. C’est généralement le signe qu’il faut ouvrir un bogue sur le site [Bugzilla de Gentoo](https://bugs.gentoo.org).

#### Guide du développeur

Les lecteurs qui désirent apprendre comment développer Gentoo peuvent jeter un œil au [guide de développement](https://devmanual.gentoo.org/). Ce manuel fournit les instructions pour écrire des _ebuilds_, travailler avec les _eclass_ et fournit une définition de beaucoup de [concepts généraux](https://devmanual.gentoo.org/general-concepts/index.html) sous-jacents.

## Pensées finales

Gentoo est une distribution robuste, flexible et parfaitement maintenue. La communauté des développeurs est à l’écoute de retours sur comment rendre Gentoo une distribution encore **meilleure**.

Pour rappel, les retours sur ce manuel doivent suivre les consignes détaillées dans la section [Comment améliorer le manuel ?](/wiki/Handbook:Main_Page/fr#How_do_I_improve_the_Handbook.3F "Handbook:Main Page/fr") au début du mode d’emploi.

Nous sommes impatients de découvrir comment nos utilisateurs vont choisir de personnaliser Gentoo en fonction de leur unique situation et besoin.

[← Configurer le chargeur d’amorçage](/wiki/Handbook:PPC/Installation/Bootloader/fr "Previous part") [Accueil](/wiki/Handbook:PPC/fr "Handbook:PPC/fr") [→](/wiki/Handbook:PPC/Working/Portage/fr "Next part")