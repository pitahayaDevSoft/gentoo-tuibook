# Networking

Other languages:

- [Deutsch](/wiki/Handbook:HPPA/Installation/Networking/de "Handbuch:HPPA/Installation/Netzwerk (100% translated)")
- [English](/wiki/Handbook:HPPA/Installation/Networking "Handbook:HPPA/Installation/Networking (100% translated)")
- [español](/wiki/Handbook:HPPA/Installation/Networking/es "Manual de Gentoo: HPPA/Instalación/Redes (100% translated)")
- français
- [italiano](/wiki/Handbook:HPPA/Installation/Networking/it "Handbook:HPPA/Installation/Networking/it (100% translated)")
- [magyar](/wiki/Handbook:HPPA/Installation/Networking/hu "Handbook:HPPA/Installation/Networking/hu (100% translated)")
- [polski](/wiki/Handbook:HPPA/Installation/Networking/pl "Handbook:HPPA/Installation/Networking (100% translated)")
- [português do Brasil](/wiki/Handbook:HPPA/Installation/Networking/pt-br "Manual:HPPA/Instalação/Rede (100% translated)")
- [русский](/wiki/Handbook:HPPA/Installation/Networking/ru "Handbook:HPPA/Installation/Networking (100% translated)")
- [தமிழ்](/wiki/Handbook:HPPA/Installation/Networking/ta "கையேடு:HPPA/நிறுவல்/வலையமைத்தல் (100% translated)")
- [中文（中国大陆）‎](/wiki/Handbook:HPPA/Installation/Networking/zh-cn "手册：HPPA/安装/配置网络 (100% translated)")
- [日本語](/wiki/Handbook:HPPA/Installation/Networking/ja "ハンドブック:HPPA/インストール/ネットワーク (100% translated)")
- [한국어](/wiki/Handbook:HPPA/Installation/Networking/ko "Handbook:HPPA/Installation/Networking/ko (100% translated)")

[Sommaire du manuel](/wiki/Handbook:HPPA "Handbook:HPPA")[Installation](/wiki/Handbook:HPPA/Full/Installation/fr "Handbook:HPPA/Full/Installation/fr")[À propos de l’installation](/wiki/Handbook:HPPA/Installation/About/fr "Handbook:HPPA/Installation/About/fr")[Choix du support](/wiki/Handbook:HPPA/Installation/Media/fr "Handbook:HPPA/Installation/Media/fr")Configurer le réseau[Préparer les disques](/wiki/Handbook:HPPA/Installation/Disks/fr "Handbook:HPPA/Installation/Disks/fr")[Installer l’archive _stage3_](/wiki/Handbook:HPPA/Installation/Stage/fr "Handbook:HPPA/Installation/Stage/fr")[Installer le système de base](/wiki/Handbook:HPPA/Installation/Base/fr "Handbook:HPPA/Installation/Base/fr")[Configurer le noyau](/wiki/Handbook:HPPA/Installation/Kernel/fr "Handbook:HPPA/Installation/Kernel/fr")[Configurer le système](/wiki/Handbook:HPPA/Installation/System/fr "Handbook:HPPA/Installation/System/fr")[Installer les outils](/wiki/Handbook:HPPA/Installation/Tools/fr "Handbook:HPPA/Installation/Tools/fr")[Configurer le système d’amorçage](/wiki/Handbook:HPPA/Installation/Bootloader/fr "Handbook:HPPA/Installation/Bootloader/fr")[Finaliser](/wiki/Handbook:HPPA/Installation/Finalizing/fr "Handbook:HPPA/Installation/Finalizing/fr")[Utiliser Gentoo](/wiki/Handbook:HPPA/Full/Working/fr "Handbook:HPPA/Full/Working/fr")[Introduction à Portage](/wiki/Handbook:HPPA/Working/Portage/fr "Handbook:HPPA/Working/Portage/fr")[Les options de la variable USE](/wiki/Handbook:HPPA/Working/USE/fr "Handbook:HPPA/Working/USE/fr")[Les fonctionnalités de Portage](/wiki/Handbook:HPPA/Working/Features/fr "Handbook:HPPA/Working/Features/fr")[Scripts d’initialisation systèmes](/wiki/Handbook:HPPA/Working/Initscripts/fr "Handbook:HPPA/Working/Initscripts/fr")[Variables d’environnement](/wiki/Handbook:HPPA/Working/EnvVar/fr "Handbook:HPPA/Working/EnvVar/fr")[Utiliser Portage](/wiki/Handbook:HPPA/Full/Portage/fr "Handbook:HPPA/Full/Portage/fr")[Fichiers et répertoires](/wiki/Handbook:HPPA/Portage/Files/fr "Handbook:HPPA/Portage/Files/fr")[Les variables](/wiki/Handbook:HPPA/Portage/Variables/fr "Handbook:HPPA/Portage/Variables/fr")[Mélanger plusieurs branches logicielles](/wiki/Handbook:HPPA/Portage/Branches/fr "Handbook:HPPA/Portage/Branches/fr")[Outils supplémentaires](/wiki/Handbook:HPPA/Portage/Tools/fr "Handbook:HPPA/Portage/Tools/fr")[Dépôt personnalisé](/wiki/Handbook:HPPA/Portage/CustomTree/fr "Handbook:HPPA/Portage/CustomTree/fr")[Fonctionnalités avancées](/wiki/Handbook:HPPA/Portage/Advanced/fr "Handbook:HPPA/Portage/Advanced/fr")[Configuration du réseau avec OpenRC](/wiki/Handbook:HPPA/Full/Networking/fr "Handbook:HPPA/Full/Networking/fr")[Bien démarrer](/wiki/Handbook:HPPA/Networking/Introduction/fr "Handbook:HPPA/Networking/Introduction/fr")[Configuration avancée](/wiki/Handbook:HPPA/Networking/Advanced/fr "Handbook:HPPA/Networking/Advanced/fr")[Les modules réseau](/wiki/Handbook:HPPA/Networking/Modular/fr "Handbook:HPPA/Networking/Modular/fr")[Sans fil](/wiki/Handbook:HPPA/Networking/Wireless/fr "Handbook:HPPA/Networking/Wireless/fr")[Ajouter des fonctionnalités](/wiki/Handbook:HPPA/Networking/Extending/fr "Handbook:HPPA/Networking/Extending/fr")[Gestion dynamique](/wiki/Handbook:HPPA/Networking/Dynamic/fr "Handbook:HPPA/Networking/Dynamic/fr")

## Contents

- [1Détection automatique du réseau](#D.C3.A9tection_automatique_du_r.C3.A9seau)
  - [1.1Utiliser DHCP](#Utiliser_DHCP)
  - [1.2Tester le réseau](#Tester_le_r.C3.A9seau)
  - [1.3Déterminer les noms des interfaces](#D.C3.A9terminer_les_noms_des_interfaces)
- [2Wireless Setup](#Wireless_Setup)
  - [2.1Optional: Checking wireless](#Optional:_Checking_wireless)
  - [2.2Using NetworkManager](#Using_NetworkManager)
- [3Configuration automatique du réseau avec net-setup](#Configuration_automatique_du_r.C3.A9seau_avec_net-setup)
  - [3.1Manual setup](#Manual_setup)
    - [3.1.1Manually connecting to an open network](#Manually_connecting_to_an_open_network)
- [4Facultatif : configuration spécifique d’une application](#Facultatif_:_configuration_sp.C3.A9cifique_d.E2.80.99une_application)
  - [4.1Facultatif : configurer les serveurs mandataires ( _proxy_)](#Facultatif_:_configurer_les_serveurs_mandataires_.28proxy.29)
  - [4.2Alternative : utilisation de _pppoe-setup_ pour l’ADSL](#Alternative_:_utilisation_de_pppoe-setup_pour_l.E2.80.99ADSL)
  - [4.3Alternative : utilisation de PPTP](#Alternative_:_utilisation_de_PPTP)
  - [4.4Comprendre la terminologie réseau](#Comprendre_la_terminologie_r.C3.A9seau)
  - [4.5Interfaces et adresses](#Interfaces_et_adresses)
  - [4.6Réseau et CIDR](#R.C3.A9seau_et_CIDR)
  - [4.7Internet](#Internet)
  - [4.8DNS](#DNS)
- [5Configuration manuelle du réseau](#Configuration_manuelle_du_r.C3.A9seau)
  - [5.1Configuration de l’adresse de l’interface](#Configuration_de_l.E2.80.99adresse_de_l.E2.80.99interface)
  - [5.2Configuration de la route par défaut](#Configuration_de_la_route_par_d.C3.A9faut)
  - [5.3Configuration DNS](#Configuration_DNS)

## Détection automatique du réseau

Il est possible que la connexion au réseau soit déjà opérationnelle.

Si le système est connecté à un réseau Ethernet ayant un routeur IPv6 ou un serveur DHCP, il est très probable que la connexion ait déjà été configurée automatiquement. Si des configurations supplémentaires sont nécessaires, la [connexion à Internet peut être testée](/wiki/Handbook:HPPA/Installation/Networking/fr#Testing_the_network "Handbook:HPPA/Installation/Networking/fr").

### Utiliser DHCP

DHCP ( _Dynamic Host Configuration Protocol_, Protocole de Configuration Dynamique des Hôtes) rend possible le fait de recevoir automatiquement des informations de mise en réseau (adresse IP, masque de sous-réseau, adresse de diffusion, passerelle, serveurs de noms, etc.)

Le serveur DHCP nécessite d’être dans la même couche réseau 2 ( _Ethernet_) que le client réclament un bail. DHCP est courant dans les réseaux privés RFC1918, mais peut aussi obtenir des informations auprès d’un fournisseur d’accès ( _ISP_).

By default the Gentoo live media uses [NetworkManager](/wiki/NetworkManager "NetworkManager") which will automatically connect to via DHCP, if this is not the case then check the Ethernet cable for issues.

### Tester le réseau

Une route « _default_ » correctement configurée est nécessaire pour la connectivité à Inernet. La configuration peut être vérifiée avec :

`root #` `ip route`

```
default via 192.168.0.1 dev enp1s0
```

Si la route « _default_ » n’est pas définie, Internet sera indisponible. Une configuration additionnelle sera nécessaire.

Une connexion à Internet peut être confirmée avec un _ping_ :

`root #` `ping -c 3 1.1.1.1`

**Conseil**

Il est utile de commencer par une adresse IP plutôt qu’un nom d’hôte. Cela permet d’isoler un problème DNS d’un problème de connectivité.

Le trafic sortant HTTPS et la résolution de nom peuvent être vérifiée avec :

`root #` `curl --location gentoo.org --output /dev/null`

À moins que curl remonte une erreur ou si un autre test échoue, la procédure d’installation peut être continuée avec [Préparer les disques](/wiki/Handbook:HPPA/Installation/Disks/fr "Handbook:HPPA/Installation/Disks/fr").

Si curl remonte une erreur, mais que le _ping_ fonctionne, suivez la [configuration DNS](/wiki/Handbook:HPPA/Installation/Networking/fr#DNS_configuration "Handbook:HPPA/Installation/Networking/fr").

Si la connection Internet n’a pu être établie, commencer par [vérifier les informations de l’interface réseau](/wiki/Handbook:HPPA/Installation/Networking/fr#Obtaining_interface_info "Handbook:HPPA/Installation/Networking/fr") puis :

- [_net-setup_ peut être utilisé](/wiki/Handbook:HPPA/Installation/Networking/fr#Using_net-setup "Handbook:HPPA/Installation/Networking/fr") pour aider à configurer l’interface réseau ;
- une [configuration par application](/wiki/Handbook:HPPA/Installation/Networking/fr#Application_specific_configuration "Handbook:HPPA/Installation/Networking/fr") peut être nécessaire ;
- la [configuration manuelle du réseau](/wiki/Handbook:HPPA/Installation/Networking/fr#Manual_network_configuration "Handbook:HPPA/Installation/Networking/fr") peut être essayée.

### Déterminer les noms des interfaces

Si le réseau ne fonctionne pas par magie, des étapes additionnelles doivent être mises en œuvre. Souvent, la 1re et de lister les interfaces réseaux.

La commande ip, de [sys-apps/iproute2](https://packages.gentoo.org/packages/sys-apps/iproute2), peut être utilisée pour interroger et configurer les interfaces réseaux.

Le paramètre « link » permet d’afficher les interfaces réseaux :

`root #` `ip link`

```
1: lo: <LOOPBACK,UP,LOWER_UP> mtu 65536 qdisc noqueue state UNKNOWN mode DEFAULT group default qlen 1000
    link/loopback 00:00:00:00:00:00 brd 00:00:00:00:00:00
4: enp1s0: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc pfifo_fast state UP mode DEFAULT group default qlen 1000
    link/ether e8:40:f2:ac:25:7a brd ff:ff:ff:ff:ff:ff
```

Le paramètre « address » peut être utilisé pour obtenir des informations sur l’adressage :

`root #` `ip address`

```
2: enp1s0: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc pfifo_fast state UP group default qlen 1000<pre>
    link/ether e8:40:f2:ac:25:7a brd ff:ff:ff:ff:ff:ff
    inet 10.0.20.77/22 brd 10.0.23.255 scope global enp1s0
       valid_lft forever preferred_lft forever
    inet6 fe80::ea40:f2ff:feac:257a/64 scope link
       valid_lft forever preferred_lft forever
```

La sortie de cette commande contient des informations pour chaque interface réseau du système. Les entrées commencent avec un index de périphérique suivi du nom : enp1s0.

**Conseil**

Si aucune interface n’est affichée autre que « lo » ( _loopack_), dans ce cas le réseau est dysfonctionnel ou le pilote n’a pas été chargé dans le noyau. Ces situations sortent du cadre du manuel, veuillez contacter [#gentoo](ircs://irc.libera.chat/#gentoo) ([webchat](https://web.libera.chat/#gentoo)).

Dans le reste de ce document, il sera assumé que l’interface réseau s’appelle enp1s0.

Suite à l’évolution vers des [noms d’interfaces réseau prévisibles (anglais)](http://www.freedesktop.org/wiki/Software/systemd/PredictableNetworkInterfaceNames/), le nom des interfaces peut différer de l’ancien système de nommage en eth0. Les supports d’installation peuvent afficher des noms d'interface tels que eno0, ens1 ou encore enp5s0.

## Wireless Setup

### Optional: Checking wireless

iw may be used to display the current wireless settings on the card, this should also show that the kernel wireless stack is working (note that the iw command is only available on the following architectures: **amd64**, **x86**, **arm**, **arm64**, **ppc**, **ppc64**, and **riscv**):

`root #` `iw dev wlp9s0 info`

```
Interface wlp9s0
	ifindex 3
	wdev 0x1
	addr 00:00:00:00:00:00
	type managed
	wiphy 0
	channel 11 (2462 MHz), width: 20 MHz (no HT), center1: 2462 MHz
	txpower 30.00 dBm

```

**Remarque**

Some wireless cards may have a device name of wlan0 or ra0 instead of wlp9s0. Run ip link to determine the correct device name.

### Using NetworkManager

The fastest way to enable wireless connectivity in the Gentoo Live media is by using the nmtui command and following the on-screen instructions to setup the wireless network.

`root #` `nmtui`

LiveGUI users can click on the network icon in the bottom right of the taskbar to also configure WiFi.

## Configuration automatique du réseau avec net-setup

Si la configuration automatique n’a pas fonctionné, le support Gentoo propose des scripts pour aider la configuration réseau. net-setup peut être utilisé pour configurer une connexion sans fil et une adresse IP statique.

`root #` `net-setup enp1s0`

L’exécution de la commande net-setup posera quelques questions au sujet de l’environnement réseau et utilisera ces informations pour configurer wpa\_supplicant ou une adresse statique.

**Important**

Le réseau doit être [testé](/wiki/Handbook:HPPA/Installation/Networking/fr#Testing_the_network "Handbook:HPPA/Installation/Networking/fr") avant toute configuration. Si le script ne résout pas le problème, la [configuration manuelle](/wiki/Handbook:HPPA/Installation/Networking/fr#Manual_network_configuration "Handbook:HPPA/Installation/Networking/fr") est obligatoire.

### Manual setup

Wireless networks may alternatively be set up with a daemon such as wpa\_supplicant or iwd. For more information on configuring wireless networking in Gentoo Linux, please read the [Wireless networking chapter](/wiki/Handbook:HPPA/Networking/Wireless/fr "Handbook:HPPA/Networking/Wireless/fr") in the Gentoo Handbook.}}

#### Manually connecting to an open network

For most users, there are only two settings needed to connect, the ESSID (aka wireless network name) and, optionally, the WEP key.

- First, ensure the interface is active:

`root #` `ip link set dev wlp9s0 up`

- To connect to an open network with the name _GentooNode_:

`root #` `iw dev wlp9s0 connect -w GentooNode`

## Facultatif : configuration spécifique d’une application

**Conseil**

These steps are provided for users where using nmtui is not able to configure their network's needs.

Les méthodes qui suivent ne sont généralement pas recommandées. Mais peuvent se révéler utiles dans certaines situations où une configuration supplémentaire est nécessaire pour la connectivité à Internet.

### Facultatif : configurer les serveurs mandataires ( _proxy_)

Si Internet est accessible via un serveur mandataire, il est nécessaire de définir des informations pour les différents protocoles supportés par Portage. Portage supporte les variables http\_proxy, ftp\_proxy et RSYNC\_PROXY pour télécharger ses paquets via wget et rsync.

Certains navigateurs en mode texte comme links peuvent également utiliser ces variables d’environnement ; pour un accès HTTPS, il faudra également définir https\_proxy. Contrairement à Portage qui n’a pas besoin de paramètres supplémentaires, links nécessite un paramétrage explicite.

Dans la plupart des cas, il suffit de définir les variables à l’aide du nom d’hôte du serveur. Comme exemple, nous supposons que le proxy est appelé proxy.gentoo.org et le port 8080.

**Remarque**

Le symbole `#` est un commentaire. Il a été ajouté pour apporter des explications, mais n’est pas nécessaire dans les commandes.

Pour configurer un proxy HTTP (pour le trafic HTTP et HTTPS) :

`root #` `export http_proxy="http://proxy.gentoo.org:8080" # Applies to Portage and Links
`

`root #` `export https_proxy="http://proxy.gentoo.org:8080" # Only applies for Links
`

Si le serveur proxy requiert un nom d’utilisateur et un mot de passe, utilisez la syntaxe suivante pour définir la variable :

`root #` `export http_proxy="http://username:password@proxy.gentoo.org:8080" # Pris en compte par Portage et links
`

`root #` `export https_proxy="http://username:password@proxy.gentoo.org:8080" # Seul links l’utilisera
`

Lancer links avec le paramètre suivant utiliser un serveur mandataire :

`user $` `links -http-proxy ${http_proxy} -https-proxy ${https_proxy} `

Pour configurer un serveur mandataire FTP (Portage et links) :

`root #` `export ftp_proxy="ftp://proxy.gentoo.org:8080" # Pris en compte par Portage et links`

Lancer links avec le paramètre suivant utiliser un serveur mandataire :

`user $` `links -ftp-proxy ${ftp_proxy} `

Pour configurer un serveur mandataire RSYNC pour Portage :

`root #` `export RSYNC_PROXY="proxy.gentoo.org:8080" # Utile pour Portage ; links ne supporte pas le RSYNC`

### Alternative : utilisation de _pppoe-setup_ pour l’ADSL

Si PPPoE est nécessaire pour l’accès Internet, le support Gentoo inclut le script pppoe-setup pour simplifier la configuration ppp.

Pendant le paramétrage, pppoe-setup va demander :

- le nom de l’interface Ethernet connecté au modem ADSL ;
- le nom et le mot de passe PPPoE ;
- l’IP du serveur DNS ;
- si un pare-feu est nécessaire.

`root #` `pppoe-setup
`

`root #` `pppoe-start`

En cas d’erreur, les identifiants dans /etc/ppp/pap-secrets ou /etc/ppp/chap-secrets doivent être vérifiés. S’ils sont corrects, la sélection de l’interface PPPoE Ethernet doit être contrôlées.

### Alternative : utilisation de PPTP

Si le support PPTP est nécessaire, utilisez la commande pptpclient. Mais il faudra d’abord le configurer.

Modifiez le fichier /etc/ppp/pap-secrets ou /etc/ppp/chap-secrets pour qu’il contienne le bon nom d’utilisateur et mot de passe :

`root #` `nano /etc/ppp/chap-secrets`

Puis peaufinez de/etc/ppp/options.pptp si nécessaire :

`root #` `nano /etc/ppp/options.pptp`

Quand tout cela est fait, il suffit d’exécuter pptp (avec les options qui n’ont pu être définies dans options.pptp) pour se connecter au serveur :

`root #` `pptp <IPv4 du serveur>`

### Comprendre la terminologie réseau

Si toutes les tentatives précédentes ont échoué, le réseau doit être configuré manuellement. Ce n’est pas particulièrement difficile, mais cela doit être effectué avec attention. Cette section explique la terminologie et les principaux concepts pour permettre aux utilisateurs d’effectuer cette configuration manuelle.

**Conseil**

Certains modems ( _Carrier Provided Equipment_) combinent les fonctionnalités routeur, point d’accès, modem, serveur DHCP et DNS tout en un. Il est important de différencier les fonctionnalités de l’appareil physique.

### Interfaces et adresses

Les « interfaces » sont des représentation logiques des périphériques réseaux. Une interface a besoin d’une adresse pour communiquer avec les autres périphériques sur le réseau. Bien qu’une seule adresse soit nécessaire, plusieurs peuvent être assignées à la même interface. C’est particulièrement vrai pour la combinaisons IPv4/IPv6.

Par convention, cette documentation va considérer que l’interface enp1s0 utilise l’adresse 192.168.0.2.

**Important**

Les adresses IP peuvent être paramétrée arbitrairement. Mais, il est possible que plusieurs périphériques utilisent la même adresse IP, ce qui cause un conflit d’IP. Utiliser un serveur DHCP or SLAAC permet d’éviter cela.

**Conseil**

IPv6 utilise **S** tate **L** ess **A** ddress **A** uto **C** onfiguration (SLAAC) pour configurer les adresses (configuration automatique de l’adresse, mais sans DNS). Dans la majorité des situations, paramétrer une adresse IPv6 manuelle est une mauvaise pratique. Si un suffixe spécifique est souhaité, les [jetons d’identification d’interface](/wiki/IPv6_Static_Addresses_using_Tokens "IPv6 Static Addresses using Tokens") peuvent être utilisés.

### Réseau et CIDR

Une fois l’adresse choisie, comment un périphérique peut communiquer avec d’autres ?

Une adresse IP peut être associés à un réseau. Les réseaux sont des suites logiques d’adresses IP.

La notation _CIDR_ ( _Classless Inter-Domain Routing_) permet de visualiser la taille des réseaux.

- Le _CIDR_ est souvent noté avec un « / » pour réprésenter la taille d’un réseau.

  - La formule _2 ^ (32 – CIDR)_ permet de calculer la taille du réseau.
  - Une fois la taille calculée, 2 doit être retiré pour connaître le nombre d’hôtes possibles.
    - La 1re IP d’un réseau est l’adresse du réseau et la dernière celle de _broadcast_ (elle permet la diffusion à tous les hôtes du réseau). Ces adresses sont spéciales et ne peuvent pas être utilisées pour des hôtes normaux.

**Conseil**

Les _CIDR_ les plus communs sont « /24 » et « /32 » qui représentent respectivement 254 et 1 nœud.

Un _CIDR_ de « /24 » est la taille la plus fréquente pour un particulier. Elle correspond à un masque sous-réseau « 255.255.255.0 », où les 8 derniers bits sont réservés pour les adresses des hôtes sur le réseau.

192.168.0.2/24 peut être lu comme :

- 192.168.0.2 est l’ _adresse_ ;
- le _réseau_ est 192.168.0.0 ;
- la _taille_ est 254 ( _2 ^ (32 – 24) – 2_) ;

  - les adresses utilisables sont entre 192.168.0.1 et 192.168.0.254 ;
- l’adresse _broadcast_ est 192.168.0.255 ;

  - dans la plupart des cas, c’est la dernière adresse qui est celle de _broadcast_, mais cela peut être modifié.

En utilisant cette configuration, un périphérique devrait pouvoir communiquer avec n’importe quel hôte du _réseau_ 192.168.0.0.

### Internet

Une fois le périphérique configuré, comment peut-il communiquer sur Internet ?

Pour communiquer en dehors du réseau, le _routage_ doit être utilisé. Un routeur est un appareil qui transfert le trafic vers d’autre périphériques. Les routes « _default_ » ou « _gateway_ » (passerelle) permettent de référence quel périphérique est utilisé pour sortir du réseau.

**Conseil**

La pratique courante est de prendre la 1re ou dernière adresse du réseau comme passerelle.

Si un routeur connecté sur Internet est disponible à 192.168.0.1, il peut être utilisé comme « _default route_ » pour fournir un accès à Internet.

Pour résumer :

- les interfaces doivent être configurée avec une « adresse » et un « réseau » (avec un _CIDR_) ;
- le réseau doit pouvoir accéder à un « routeur » dans le même réseau ;
- la route par défaut ( _default route_) doit être configurée pour permettre une communication à l’extérieur du réseau via une « passerelle » ( _gateway_).

### DNS

Se souvenir des IP est difficile. Le système de nom de domaine ( _DNS_) a été créé pour lié un « nom de domaine » à une « adresse IP ».

Linux utilise /etc/resolv.conf pour définir les serveurs de résolution de noms ( _nameservers_).

**Conseil**

Beaucoup de routeur font office de serveur _DNS_, et utiliser ce serveur local permet d’augmenter la confidentialité et accélérer les requêtes avec du cache.

Beaucoup de fournisseurs d’accès à Internet ( _ISP_) annoncent un serveur DNS via DHCP. Utiliser un serveur DNS local améliore la latence, mais la plupart des serveurs publics de DNS retourneront le même résultat. L’usage d’un serveur en particulier est une préférence utilisateur.

## Configuration manuelle du réseau

### Configuration de l’adresse de l’interface

**Important**

Lorsque l’adresse est choisie manuellement, la composition du réseau doit être prise en compte. Une adresse IP doit être unique ; les conflits causent des coupures réseaux.

Pour configurer enp1s0 avec l’adresse 192.168.0.2 et le _CIDR_/24 :

`root #` `ip address add 192.168.0.2/24 dev enp1s0`

**Conseil**

Le début de la commande peut être raccourci en « ip a ».

### Configuration de la route par défaut

Configurer l’adresse et le réseau de l’interface va configurer la route du réseau pour la communication interne :

`root #` `ip route`

```
192.168.0.0/24 dev enp1s0 proto kernel scope link src 192.168.0.2
```

**Conseil**

La commande peut être raccourcie en ip r.

La route par défaut 192.168.0.1 peut être paramétrée avec :

`root #` `ip route add default via 192.168.0.1`

### Configuration DNS

Bien que les noms de serveurs sont souvent obtenus par DHCP, il est possible de les paramétrer manuellement en ajoutant des entrées `nameserver` dans /etc/resolv.conf.

**Attention !**

Si _dhcpcd_ est lancé, les changements dans /etc/resolv.conf seront écrasés. Vous pouvez vérifier son lancement avec `ps x | grep dhcpcd`.

nano est inclus dans le support Gentoo et peut être utilisé pour éditer /etc/resolv.conf avec :

`root #` `nano /etc/resolv.conf`

Les lignes contenant le mot-clé `nameserver` sont suivies de l’adresse d’un serveur DNS :

FILE **`/etc/resolv.conf`** **Utiliser Quad9 DNS.**

```
nameserver 9.9.9.9
nameserver 149.112.112.112

```

FILE **`/etc/resolv.conf`** **Utiliser les DNS Cloudflare.**

```
nameserver 1.1.1.1
nameserver 1.0.0.1

```

Le fonctionnement du DNS peut être vérifier en _pingant_ un nom de domaine :

`root #` `ping -c 3 gentoo.org`

Une fois la connectivité [vérifiée](/wiki/Handbook:HPPA/Installation/Networking/fr#Testing_the_network "Handbook:HPPA/Installation/Networking/fr"), continuez avec [Préparer les disques](/wiki/Handbook:HPPA/Installation/Disks/fr "Handbook:HPPA/Installation/Disks/fr").

[← Choisir le support](/wiki/Handbook:HPPA/Installation/Media/fr "Previous part") [Accueil](/wiki/Handbook:HPPA/fr "Handbook:HPPA/fr") [Préparer les disques →](/wiki/Handbook:HPPA/Installation/Disks/fr "Next part")