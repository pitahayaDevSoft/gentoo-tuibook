# Networking

Other languages:

- [Deutsch](/wiki/Handbook:PPC/Installation/Networking/de "Handbuch:PPC/Installation/Netzwerk (100% translated)")
- [English](/wiki/Handbook:PPC/Installation/Networking "Handbook:PPC/Installation/Networking (100% translated)")
- [español](/wiki/Handbook:PPC/Installation/Networking/es "Manual de Gentoo: PPC/Instalación/Redes (100% translated)")
- [français](/wiki/Handbook:PPC/Installation/Networking/fr "Handbook:PPC/Installation/Networking/fr (100% translated)")
- italiano
- [magyar](/wiki/Handbook:PPC/Installation/Networking/hu "Handbook:PPC/Installation/Networking/hu (100% translated)")
- [polski](/wiki/Handbook:PPC/Installation/Networking/pl "Handbook:PPC/Installation/Networking (100% translated)")
- [português do Brasil](/wiki/Handbook:PPC/Installation/Networking/pt-br "Handbook:PPC/Installation/Networking/pt-br (50% translated)")
- [русский](/wiki/Handbook:PPC/Installation/Networking/ru "Handbook:PPC/Installation/Networking (100% translated)")
- [தமிழ்](/wiki/Handbook:PPC/Installation/Networking/ta "கையேடு:PPC/நிறுவல்/வலையமைத்தல் (100% translated)")
- [中文（中国大陆）‎](/wiki/Handbook:PPC/Installation/Networking/zh-cn "手册：PPC/安装/配置网络 (100% translated)")
- [日本語](/wiki/Handbook:PPC/Installation/Networking/ja "ハンドブック:PPC/インストール/ネットワーク (100% translated)")
- [한국어](/wiki/Handbook:PPC/Installation/Networking/ko "Handbook:PPC/Installation/Networking/ko (100% translated)")

[Manuale PPC](/wiki/Handbook:PPC/it "Handbook:PPC/it")[Installazione](/wiki/Handbook:PPC/Full/Installation/it "Handbook:PPC/Full/Installation/it")[Riguardo l'installazione](/wiki/Handbook:PPC/Installation/About/it "Handbook:PPC/Installation/About/it")[Scegliere il supporto](/wiki/Handbook:PPC/Installation/Media/it "Handbook:PPC/Installation/Media/it")Configurare la rete[Preparare i dischi](/wiki/Handbook:PPC/Installation/Disks/it "Handbook:PPC/Installation/Disks/it")[Il file stage](/wiki/Handbook:PPC/Installation/Stage/it "Handbook:PPC/Installation/Stage/it")[Installare il sistema base](/wiki/Handbook:PPC/Installation/Base/it "Handbook:PPC/Installation/Base/it")[Configurare il kernel](/wiki/Handbook:PPC/Installation/Kernel/it "Handbook:PPC/Installation/Kernel/it")[Configurare il sistema](/wiki/Handbook:PPC/Installation/System/it "Handbook:PPC/Installation/System/it")[Installare gli strumenti](/wiki/Handbook:PPC/Installation/Tools/it "Handbook:PPC/Installation/Tools/it")[Impostare programma d'avvio](/wiki/Handbook:PPC/Installation/Bootloader/it "Handbook:PPC/Installation/Bootloader/it")[Concludere](/wiki/Handbook:PPC/Installation/Finalizing/it "Handbook:PPC/Installation/Finalizing/it")[Lavorare con Gentoo](/wiki/Handbook:PPC/Full/Working/it "Handbook:PPC/Full/Working/it")[Introduzione a Portage](/wiki/Handbook:PPC/Working/Portage/it "Handbook:PPC/Working/Portage/it")[Opzioni USE](/wiki/Handbook:PPC/Working/USE/it "Handbook:PPC/Working/USE/it")[Funzionalità di Portage](/wiki/Handbook:PPC/Working/Features/it "Handbook:PPC/Working/Features/it")[Sistema script di init](/wiki/Handbook:PPC/Working/Initscripts/it "Handbook:PPC/Working/Initscripts/it")[Variabili d'ambiente](/wiki/Handbook:PPC/Working/EnvVar/it "Handbook:PPC/Working/EnvVar/it")[Lavorare con Portage](/wiki/Handbook:PPC/Full/Portage/it "Handbook:PPC/Full/Portage/it")[File e cartelle](/wiki/Handbook:PPC/Portage/Files/it "Handbook:PPC/Portage/Files/it")[Variabili](/wiki/Handbook:PPC/Portage/Variables/it "Handbook:PPC/Portage/Variables/it")[Miscelare i rami del software](/wiki/Handbook:PPC/Portage/Branches/it "Handbook:PPC/Portage/Branches/it")[Strumenti aggiuntivi](/wiki/Handbook:PPC/Portage/Tools/it "Handbook:PPC/Portage/Tools/it")[Catalogo pacchetti personalizzato](/wiki/Handbook:PPC/Portage/CustomTree/it "Handbook:PPC/Portage/CustomTree/it")[Funzionalità avanzate](/wiki/Handbook:PPC/Portage/Advanced/it "Handbook:PPC/Portage/Advanced/it")[Configurare la rete](/wiki/Handbook:PPC/Full/Networking/it "Handbook:PPC/Full/Networking/it")[Come iniziare](/wiki/Handbook:PPC/Networking/Introduction/it "Handbook:PPC/Networking/Introduction/it")[Configurazione avanzata](/wiki/Handbook:PPC/Networking/Advanced/it "Handbook:PPC/Networking/Advanced/it")[Networking modulare](/wiki/Handbook:PPC/Networking/Modular/it "Handbook:PPC/Networking/Modular/it")[Wireless](/wiki/Handbook:PPC/Networking/Wireless/it "Handbook:PPC/Networking/Wireless/it")[Aggiungere funzionalità](/wiki/Handbook:PPC/Networking/Extending/it "Handbook:PPC/Networking/Extending/it")[Gestione dinamica](/wiki/Handbook:PPC/Networking/Dynamic/it "Handbook:PPC/Networking/Dynamic/it")

## Contents

- [1Configurare la rete automaticamente](#Configurare_la_rete_automaticamente)
  - [1.1Usare DHCP](#Usare_DHCP)
  - [1.2Provare la rete](#Provare_la_rete)
- [2Ottenere informazioni sulle interfacce](#Ottenere_informazioni_sulle_interfacce)
- [3Configurare il wireless](#Configurare_il_wireless)
  - [3.1Facoltativo: Controllare il wireless](#Facoltativo:_Controllare_il_wireless)
  - [3.2Uso di NetworkManager](#Uso_di_NetworkManager)
  - [3.3Uso di net-setup](#Uso_di_net-setup)
  - [3.4Configurare manualmente](#Configurare_manualmente)
    - [3.4.1Connessione manuale ad una rete aperta](#Connessione_manuale_ad_una_rete_aperta)
- [4Configurare una applicazione specifica](#Configurare_una_applicazione_specifica)
  - [4.1Configurare i proxy](#Configurare_i_proxy)
  - [4.2Uso di pppoe-setup per l'ADSL](#Uso_di_pppoe-setup_per_l.27ADSL)
  - [4.3Usare PPTP](#Usare_PPTP)
- [5Fondamenti di Internet e IP](#Fondamenti_di_Internet_e_IP)
  - [5.1Interfacce e indirizzi](#Interfacce_e_indirizzi)
  - [5.2Reti e CIDR](#Reti_e_CIDR)
  - [5.3Internet](#Internet)
  - [5.4Il DNS (Sistema dei Nomi di Dominio)](#Il_DNS_.28Sistema_dei_Nomi_di_Dominio.29)
- [6Configurare manualmente un indirizzo IP statico della rete](#Configurare_manualmente_un_indirizzo_IP_statico_della_rete)
  - [6.1Configurare un indirizzo dell'interfaccia](#Configurare_un_indirizzo_dell.27interfaccia)
  - [6.2Configurare route predefinita](#Configurare_route_predefinita)
  - [6.3Configurare DNS](#Configurare_DNS)

## Configurare la rete automaticamente

Forse è già funzionante?

Se il sistema è connesso ad una rete Ethernet con un router IPV6 o un server DHCP, è molto probabile che la rete del sistema sia stata configurata automaticamente. Se non sono richieste ulteriori, avanzate configurazioni [si può provare la connessione a Internet](/wiki/Handbook:PPC/Installation/Networking/it#Testing_the_network "Handbook:PPC/Installation/Networking/it").

### Usare DHCP

DHCP (Dynamic Host Configuration Protocol) aiuta nella configurazione della rete, e può fornire automaticamente una serie di parametri come: indirizzo IP, maschera di sottorete, instradamenti, server DNS, server NTP, etc.

Il DHCP richiede che un server sia attivo sullo stesso segmento _Layer 2_ ( _Ethernet_) del client che richiede un _lease_. DHCP è spesso usato su reti ( _private_) RFC1918, ma è anche usato per acquisire informazioni su IP pubblici dai fornitori di accesso a internet(ISP).

Per impostazione predefinita il supporto live di Gentoo utilizza [NetworkManager](/wiki/NetworkManager "NetworkManager") che si connetterà automaticamente attraverso il DHCP, se non è così allora controllare il cavo Ethernet per problemi.

### Provare la rete

Una route(instradamento) predefinita configurata bene è un componente fondamentale della connessione a Internet, la si può verificare con:

`root #` `ip route`

```
default via 192.168.0.1 dev enp1s0
```

Se una route predefinita non è configurata, la connessione Internet non è disponibile, ed è necessaria della configurazione aggiuntiva.

La connessione ad Internet di base può essere confermata col comando ping:

`root #` `ping -c 3 1.1.1.1`

**Suggerimento**

E' utile iniziare facendo il ping ad un indirizzo IP noto piuttosto che ad un hostname. Ciò può isolare i problemi di DNS dai problemi alla connessione Internet di base.

Accessi HTTPS in uscita e risoluzione dei DNS possono essere confermati con:

`root #` `curl --location gentoo.org --output /dev/null`

A meno che curl non dia un errore o altre prove non vadano a buon fine, il processo di installazione può continuare con [preparare i dischi](/wiki/Handbook:PPC/Installation/Disks/it "Handbook:PPC/Installation/Disks/it").

Se curl da un errore, ma il ping verso Internet funziona, il [DNS potrebbe necessitare di configurazione](/wiki/Handbook:PPC/Installation/Networking/it#DNS_configuration "Handbook:PPC/Installation/Networking/it").

Se la connessione a Internet non è stata stabilita, prima [si dovrebbero verificare le informazioni sull' interfaccia](/wiki/Handbook:PPC/Installation/Networking/it#Obtaining_interface_info "Handbook:PPC/Installation/Networking/it"), poi:

- [Si può usare nmtui](/wiki/Handbook:PPC/Installation/Networking/it#Using_NetworkManager "Handbook:PPC/Installation/Networking/it") come supporto alla configurazione della rete.
- Potrebbe essere necessario [configurare una applicazione specifica](/wiki/Handbook:PPC/Installation/Networking/it#Application_specific_configuration "Handbook:PPC/Installation/Networking/it").
- Si può provare a [configurare manualmente la rete](/wiki/Handbook:PPC/Installation/Networking/it#Manual_network_configuration "Handbook:PPC/Installation/Networking/it").

## Ottenere informazioni sulle interfacce

Se la rete non è pronta per l'uso, si devono intraprendere dei passi aggiuntivi per abilitare la connettività a Internet. Solitamente il primo passo è elencare le interfacce di rete del computer.

Il comando ip, parte del pacchetto [sys-apps/iproute2](https://packages.gentoo.org/packages/sys-apps/iproute2), può essere usato per interrogare e configurare la rete del sistema.

L'argomento **link** può essere usato per mostrare le connessioni delle interfacce di rete:

`root #` `ip link`

```
1: lo: <LOOPBACK,UP,LOWER_UP> mtu 65536 qdisc noqueue state UNKNOWN mode DEFAULT group default qlen 1000
    link/loopback 00:00:00:00:00:00 brd 00:00:00:00:00:00
4: enp1s0: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc pfifo_fast state UP mode DEFAULT group default qlen 1000
    link/ether e8:40:f2:ac:25:7a brd ff:ff:ff:ff:ff:ff
```

L'argomento **address** può essere usato per chiedere informazioni sull'indirizzo del dispositivo:

`root #` `ip address`

```
2: enp1s0: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc pfifo_fast state UP group default qlen 1000<pre>
    link/ether e8:40:f2:ac:25:7a brd ff:ff:ff:ff:ff:ff
    inet 10.0.20.77/22 brd 10.0.23.255 scope global enp1s0
       valid_lft forever preferred_lft forever
    inet6 fe80::ea40:f2ff:feac:257a/64 scope link
       valid_lft forever preferred_lft forever
```

L'output di questo comando contiene informazioni di ogni interfaccia di rete del sistema. Le voci iniziano con l'indice del dispositivo, seguito dal suo nome: enp1s0.

**Suggerimento**

Se non viene mostrata alcuna interfaccia a parte la **lo** ( _loopback_), allora l'hardware di rete è guasto o il relativo driver non è stato caricato nel kernel. Entrambe le situazioni vanno al di là dello scopo di questo manuale. Si prega di chiedere assistenza contattando [#gentoo](ircs://irc.libera.chat/#gentoo) ([webchat](https://web.libera.chat/#gentoo)).

Per uniformità, il manuale considererà che l'interfaccia di rete primaria sia chiamata enp1s0.

**Nota**

Avendo adottato i [nomi che indicano il tipo di interfaccia](https://www.freedesktop.org/wiki/Software/systemd/PredictableNetworkInterfaceNames/), il nome dell'interfaccia di rete potrebbe essere abbastanza diverso dalla vecchia convenzione dei nomi eth0. I dischi di installazione recenti usano nomi con prefissi come eno0, ens1, o enp5s0.

## Configurare il wireless

### Facoltativo: Controllare il wireless

iw può essere usato per mostrare la configurazione corrente del wireless per la scheda, ciò potrebbe anche mostrare che l'infrastruttura wireless del kernel è funzionante (si badi bene che il comando iw è disponibile solo nelle seguenti architetture: **amd64**, **x86**, **arm**, **arm64**, **ppc**, **ppc64**, and **riscv**):

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

**Nota**

Alcune schede wireless potrebbero avere un nome di dispositivo pari a wlan0 o ra0 invece di wlp9s0. Eseguire ip link per stabilire il nome corretto del dispositivo.

### Uso di NetworkManager

Il modo più veloce per abilitare la connessione wireless nel supporto Gentoo Live è usare il comando nmtui e seguire a schermo le istruzioni per configurare la rete wireless.

`root #` `nmtui`

Gli utenti della LiveGUI possono cliccare anche sull'icona della rete in basso a destra nella barra delle applicazioni per configurare il WiFi.

### Uso di net-setup

Su alcune architetture, p.es. **HPPA**, un utente potrebbe essere costretto a usare lo strumento net-setup per configurare una connessione wireless se NetworkManager non fosse ancora disponibile.

`root #` `killall dhcpcd`

`root #` `net-setup enp1s0`

net-setup chiederà alcune informazioni sull'ambiente di rete e le userà per configurare wpa\_supplicant o _l'indirizzamento statico_.

**Importante**

La condizione della rete dovrebbe essere [provata](/wiki/Handbook:PPC/Installation/Networking/it#Testing_the_network "Handbook:PPC/Installation/Networking/it") dopo ogni passo di configurazione eseguito. Nel caso gli script di configurazione non funzionassero, è necessaria la [configurazione manuale del network](/wiki/Handbook:PPC/Installation/Networking/it#Manual_network_configuration "Handbook:PPC/Installation/Networking/it").

### Configurare manualmente

Le reti wireless possono alternativamente essere configurate con un servizio come wpa\_supplicant o iwd. Per maggiori informazioni sulla configurazione della rete wireless in Gentoo Linux, si prega di leggere il [capitolo sulla rete wireless](/wiki/Handbook:PPC/Networking/Wireless/it "Handbook:PPC/Networking/Wireless/it") nel manuale di Gentoo.}}

#### Connessione manuale ad una rete aperta

Per la maggior parte degli utenti, sono necessarie solo due impostazioni per connettersi, l'ESSID (o nome della rete wireless) e, facoltativamente, la chiave WEP.

- Prima, accertarsi che l'interfaccia sia attiva:

`root #` `ip link set dev wlp9s0 up`

- Per connettersi ad una rete aperta col nome _GentooNode_:

`root #` `iw dev wlp9s0 connect -w GentooNode`

## Configurare una applicazione specifica

**Suggerimento**

Questi passaggi sono forniti per gli utenti che usando nmtui non riescono a configuare la rete per le proprie necessità.

I metodi seguenti non sono normalmente necessari, ma possono aiutare nelle situazioni in cui è richiesta ulteriore configurazione per la connessione Internet.

### Configurare i proxy

Se si accede ad Internet attraverso un proxy, è necessario impostare le informazioni del proxy per Portage perchè vi acceda correttamente con ogni protocollo supportato. Portage osserva le variabili d'ambiente http\_proxy, ftp\_proxy, e RSYNC\_PROXY quando scarica i pacchetti attraverso i propri meccanismi di recupero wget e rsync.

Anche alcuni browser web testuali come links possono fare uso delle variabili d'ambiente che determinano le impostazioni del proxy; in particolare per l'accesso HTTPS è necessario impostare anche la variabile d'ambiente https\_proxy. Mentre Portage userà il proxy senza passare parametri aggiuntivi durante l'invocazione, links li richiederà.

Nella maggior parte dei casi, è sufficiente definire le variabili d'ambiente usando l'hostname del server. Nell'esempio seguente, si considera che il server proxy è chiamato proxy.gentoo.org e la porta è 8080.

**Nota**

Il simbolo `#` nei comandi seguenti rappresenta l'inizio di un commento. E' stato aggiunto solo per chiarezza e non c'è bisogno di digitarlo quando si da il comando.

Per definire un proxy HTTP (per il traffico HTTP/HTTPS):

`root #` `export http_proxy="http://proxy.gentoo.org:8080" # Si applica a Portage e Links
`

`root #` `export https_proxy="http://proxy.gentoo.org:8080" # Si applica solo a Links
`

Se il proxy HTTP richiede l'autenticazione, impostare nome utente e password con la seguente sintassi:

`root #` `export http_proxy="http://username:password@proxy.gentoo.org:8080" # Si applica a Portage and Links
`

`root #` `export https_proxy="http://username:password@proxy.gentoo.org:8080" # Si applica solo a Links
`

Avvia links usando i parametri seguenti per il supporto al proxy:

`user $` `links -http-proxy ${http_proxy} -https-proxy ${https_proxy} `

Per definire un proxy FTP per Portage e/o links:

`root #` `export ftp_proxy="ftp://proxy.gentoo.org:8080" # Si applica a Portage e Links`

Avvia links usando i parametri seguenti per un proxy FTP:

`user $` `links -ftp-proxy ${ftp_proxy} `

Per definire un proxy RSYNC per Portage:

`root #` `export RSYNC_PROXY="proxy.gentoo.org:8080" # Si applica a Portage; Links non supporta un proxy rsync`

### Uso di pppoe-setup per l'ADSL

Se per l'accesso ad Internet è richiesto il PPPoE, il _supporto di avvio_ di Gentoo include lo script pppoe-setup per semplificare la configurazione del ppp.

Durante l'impostazione, pppoe-setup chiederà:

- Il nome dell' _interfaccia_ Ethernet connessa al modem ADSL.
- Username e password del PPPoE.
- Indirizzi IP dei server DNS.
- Se è richiesto un firewall.

`root #` `pppoe-setup
`

`root #` `pppoe-start`

In caso di mancato funzionamento, si dovrebbero verificare le credenziali in /etc/ppp/pap-secrets o /etc/ppp/chap-secrets. Se sono corrette si dovrebbe controllare la selezione dell'interfaccia Ethernet PPPoE.

### Usare PPTP

Se è necessario il supporto PPTP, si può usare pptpclient, ma necessita di configurazione prima dell' utilizzo.

Modificare /etc/ppp/pap-secrets o /etc/ppp/chap-secrets cosicchè contengano la combinazione corretta di username/password:

`root #` `nano  /etc/ppp/chap-secrets`

Poi, se necessario, modificare /etc/ppp/options.pptp:

`root #` `nano  /etc/ppp/options.pptp`

Una volta completata la configurazione, si esegua pptp (insieme alle opzioni che non potevano essere impostate in options.pptp) per connettersi al server:

`root #` `pptp <server ipv4 address>`

## Fondamenti di Internet e IP

Se quanto riportato sopra non funziona, la rete dev' essere configurata manualmente. Non è particolarmente difficile, ma dovrebbe essere fatta con ponderazione. Questa sezione serve a chiarire la terminologia e introdurre gli utenti a concetti basilari sulla rete concernenti la configurazione manuale di una connessione Internet.

**Suggerimento**

Alcuni **CPE** ( **Carrier Provided Equipment**) combinano le funzioni di _router_, _access point_, _modem_, _DHCP server_, e _server DNS_ in un unica unità. E' importante differenziare le funzioni di un dispositivo dall'apparecchio fisico.

### Interfacce e indirizzi

Le _interfacce_ di rete sono rappresentazioni logiche di dispositivi di rete. Un _interfaccia_ necessita di un _indirizzo_ per comunicare con altri dispositivi sulla _rete_. Mentre è necessario un singolo _indirizzo_, indirizzi multipli possono essere assegnati ad una singola _interfaccia_. Ciò è particolarmente utile per configurazioni doppie dell'infrastruttura di rete (IPV4 + IPV6).

Per uniformità, questa guida di base presupporrà che l'interfaccia enp1s0 usi l'indirizzo 192.168.0.2.

**Importante**

Gli indirizzi IP possono essere impostati casualmente. Di conseguenza, è possibile che più dispositivi utilizzino lo stesso indirizzo IP, provocando un _conflitto di indirizzo_. I conflitti di indirizzo dovrebbero essere evitati usando DHCP o SLAAC.

**Suggerimento**

IPV6 normalmente utilizza **S** tate **L** ess **A** ddress **A** uto **C** onfiguration (SLAAC) per la configurazione degli indirizzi. Nella maggior parte dei casi, impostare manualmente gli indirizzi IPV6 è una cattiva abitudine. Se si preferisce uno specifico suffisso dell' indirizzo, si possono usare i [simboli di identificazione dell'interfaccia](/wiki/IPv6_Static_Addresses_using_Tokens "IPv6 Static Addresses using Tokens").

### Reti e CIDR

Una volta scelto l'indirizzo, come sa un dispositivo in che modo comunicare con gli altri?

Gli _indirizzi_ IP sono associati alle _reti_. Le _reti_ IP sono un insieme logico di indirizzi consecutivi.

La notazione _CIDR_ o _Classless Inter-Domain Routing_ è usata per distinguere le dimensioni della rete.

- Il valore _CIDR_, spesso scritto a partire da un **/**, rappresenta la dimensione della rete.

  - La formula _2 ^ (32 - CIDR)_ si può usare per calcolare la dimensione della rete.
  - Una volta calcolata la dimensione della rete, il numero dei nodi utilizzabili deve essere ridotto di 2.
    - Il primo indirizzo IP nella rete è l' _indirizzo della rete_, e l'ultimo è normalmente l' _indirizzo di broadcast_. Questi indirizzi sono speciali e non possono essere usati dai normali computer.

**Suggerimento**

I valori _CIDR_ più comuni sono **/24**, e **/32**, che rappresentano rispettivamente **254** nodi e il singolo nodo.

Un _CIDR_ di **/24** è di fatto la dimensione abituale di una rete. Questa corrisponde a una maschera di sottorete di _255.255.255.0_, dove gli ultimi 8 bit sono riservati agli indirizzi IP per i nodi sulla rete.

La notazione: 192.168.0.2/24 può essere interpretata così:

- L' _indirizzo_ 192.168.0.2
- Sulla _rete_ 192.168.0.0
- Con una dimensione di **254** ( _2 ^ (32 - 24) \- 2_)

  - Gli indizzi IP utilizzabili sono nell' intervallo 192.168.0.1 - 192.168.0.254
- Con un _indirizzo di brodcast_ di of 192.168.0.255

  - Nella maggior parte dei casi, l'ultimo indirizzo sulla rete è usato come _indirizzo di broadcast_, ma questo si può cambiare.

Usando questa configurazione, un dispositivo dovrebbe essere capace di comunicare con qualsiasi computer sulla stessa _rete_ (192.168.0.0).

### Internet

Quando un dispositivo è su una rete, come sa in che modo comunicare coi dispositivi su Internet?

Per comunicare coi dispositivi al di fuori delle _reti_ locali, deve essere usato il _routing_(instradamento). Un _router_ è semplicemente un dispositivo di rete che inoltra traffico per altri dispositivi. I termini _default route_ o _gateway_ di solito si riferiscono a qualsiasi dispositivo sulla rete attuale che viene usato per accedere alla rete esterna.

**Suggerimento**

E' abitudine comune fare in modo che il _gateway_ sia il primo o l'ultimo indirizzo su una rete.

Se un router connesso a Internet è disponibile al 192.168.0.1, può essere usato come _default route_, consentendo l'accesso a Internet.

Per riassumere:

- Le interfacce devono essere configurate con un _indirizzo_ e delle _informazioni di rete_ come il valore _CIDR_.
- L'accesso alla rete locale è utilizzato per accedere ad un _router_ sulla stessa rete.
- Viene configurata la _default route_, cosicchè il traffico destinato alle reti esterne è inoltrato al _gateway_, che fornisce l'accesso a Internet.

### Il DNS (Sistema dei Nomi di Dominio)

Ricordare gli indirizzi IP è difficile. Il _DNS_ è stato creato per permettere la mappatura tra _Nomi di Dominio_ e _indirizzi IP_.

I sistemi Linux usano /etc/resolv.conf per definire i _nameserver_ usati per la _risoluzione dei DNS_.

**Suggerimento**

Molti router possono funzionare anche da server DNS, e usando un server DNS locale si può aumentare la riservatezza e velocizzare le richieste conservandole in loco.

Molti ISP (fornitori del servizio internet) usano un server DNS che normalmente viene annunciato al _gateway_ sul DHCP. Usare un server DNS locale tende a migliorare la latenza delle richieste, ma la maggior parte dei server DNS pubblici restituisce gli stessi risultati, quindi l'uso di un server è perlopiù una questione di preferenza.

## Configurare manualmente un indirizzo IP statico della rete

### Configurare un indirizzo dell'interfaccia

**Importante**

Quando si configura manualmente l'indirizzo IP, bisogna considerare la topologia della rete locale. Gli indirizzi IP possono essere impostati a caso; i conflitti possono provocare interruzioni nella rete.

Per configurare enp1s0 con l'indirizzo 192.168.0.2 e CIDR /24:

`root #` `ip address add 192.168.0.2/24 dev enp1s0`

**Suggerimento**

L'inizio di questo comando si può abbreviare con ip a.

### Configurare route predefinita

Configurare indirizzo e informazioni della rete per un interfaccia configurerà le route del link, consentendo la comunicazione con quel segmento della rete:

`root #` `ip route`

```
192.168.0.0/24 dev enp1s0 proto kernel scope link src 192.168.0.2
```

**Suggerimento**

Questo comando si può abbreviare con ip r.

La default route può essere impostata a 192.168.0.1 con:

`root #` `ip route add default via 192.168.0.1`

### Configurare DNS

Le informazioni sul nameserver in genere vengono acquisite usando il DHCP, ma si possono impostare manualmente aggiungendo voci di `nameserver` a /etc/resolv.conf.

**Attenzione**

Se _dhcpcd_ è in esecuzione, le modifiche a /etc/resolv.conf non saranno mantenute. Lo stato può essere verificato con `ps x | grep dhcpcd`.

nano è incluso nel _supporto di avvio_ di Gentoo e può essere usato per modificare /etc/resolv.conf con:

`root #` `nano /etc/resolv.conf`

Le linee contenenti la parola chiave `nameserver` seguite dall'indirizzo IP di un server DNS vengono interrogate secondo l' ordine con cui sono state inserite.

FILE **`/etc/resolv.conf`** **Uso dei DNS Quad9.**

```
nameserver 9.9.9.9
nameserver 149.112.112.112

```

FILE **`/etc/resolv.conf`** **Uso dei DNS Cloudflare.**

```
nameserver 1.1.1.1
nameserver 1.0.0.1

```

Si può verificare lo stato dei DNS facendo il ping ad un nome di dominio:

`root #` `ping -c 3 gentoo.org`

Una volta [provata](/wiki/Handbook:PPC/Installation/Networking/it#Testing_the_network "Handbook:PPC/Installation/Networking/it") la connessione, continua a [preparare i dischi](/wiki/Handbook:PPC/Installation/Disks/it "Handbook:PPC/Installation/Disks/it").

[← Scegliere il supporto](/wiki/Handbook:PPC/Installation/Media/it "Previous part") [Home](/wiki/Handbook:PPC "Handbook:PPC") [Preparare i dischi →](/wiki/Handbook:PPC/Installation/Disks/it "Next part")