# Kernel

Other languages:

- [Deutsch](/wiki/Handbook:SPARC/Installation/Kernel/de "Handbuch:SPARC/Installation/Kernel (100% translated)")
- [English](/wiki/Handbook:SPARC/Installation/Kernel "Handbook:SPARC/Installation/Kernel (100% translated)")
- [español](/wiki/Handbook:SPARC/Installation/Kernel/es "Manual de Gentoo: SPARC/Instalación/Núcleo (100% translated)")
- [français](/wiki/Handbook:SPARC/Installation/Kernel/fr "Handbook:SPARC/Installation/Kernel/fr (100% translated)")
- italiano
- [magyar](/wiki/Handbook:SPARC/Installation/Kernel/hu "Handbook:SPARC/Installation/Kernel/hu (100% translated)")
- [polski](/wiki/Handbook:SPARC/Installation/Kernel/pl "Handbook:SPARC/Installation/Kernel (100% translated)")
- [português do Brasil](/wiki/Handbook:SPARC/Installation/Kernel/pt-br "Handbook:SPARC/Installation/Kernel/pt-br (50% translated)")
- [русский](/wiki/Handbook:SPARC/Installation/Kernel/ru "Handbook:SPARC/Installation/Kernel (100% translated)")
- [தமிழ்](/wiki/Handbook:SPARC/Installation/Kernel/ta "கையேடு:SPARC/நிறுவல்/கர்னல் (100% translated)")
- [中文（中国大陆）‎](/wiki/Handbook:SPARC/Installation/Kernel/zh-cn "手册：SPARC/安装/配置Linux内核 (100% translated)")
- [日本語](/wiki/Handbook:SPARC/Installation/Kernel/ja "ハンドブック:SPARC/インストール/カーネル (100% translated)")
- [한국어](/wiki/Handbook:SPARC/Installation/Kernel/ko "Handbook:SPARC/Installation/Kernel/ko (100% translated)")

[Manuale SPARC](/wiki/Handbook:SPARC/it "Handbook:SPARC/it")[Installazione](/wiki/Handbook:SPARC/Full/Installation/it "Handbook:SPARC/Full/Installation/it")[Riguardo l'installazione](/wiki/Handbook:SPARC/Installation/About/it "Handbook:SPARC/Installation/About/it")[Scegliere il supporto](/wiki/Handbook:SPARC/Installation/Media/it "Handbook:SPARC/Installation/Media/it")[Configurare la rete](/wiki/Handbook:SPARC/Installation/Networking/it "Handbook:SPARC/Installation/Networking/it")[Preparare i dischi](/wiki/Handbook:SPARC/Installation/Disks/it "Handbook:SPARC/Installation/Disks/it")[Il file stage](/wiki/Handbook:SPARC/Installation/Stage/it "Handbook:SPARC/Installation/Stage/it")[Installare il sistema base](/wiki/Handbook:SPARC/Installation/Base/it "Handbook:SPARC/Installation/Base/it")Configurare il kernel[Configurare il sistema](/wiki/Handbook:SPARC/Installation/System/it "Handbook:SPARC/Installation/System/it")[Installare gli strumenti](/wiki/Handbook:SPARC/Installation/Tools/it "Handbook:SPARC/Installation/Tools/it")[Impostare programma d'avvio](/wiki/Handbook:SPARC/Installation/Bootloader/it "Handbook:SPARC/Installation/Bootloader/it")[Concludere](/wiki/Handbook:SPARC/Installation/Finalizing/it "Handbook:SPARC/Installation/Finalizing/it")[Lavorare con Gentoo](/wiki/Handbook:SPARC/Full/Working/it "Handbook:SPARC/Full/Working/it")[Introduzione a Portage](/wiki/Handbook:SPARC/Working/Portage/it "Handbook:SPARC/Working/Portage/it")[Opzioni USE](/wiki/Handbook:SPARC/Working/USE/it "Handbook:SPARC/Working/USE/it")[Funzionalità di Portage](/wiki/Handbook:SPARC/Working/Features/it "Handbook:SPARC/Working/Features/it")[Sistema script di init](/wiki/Handbook:SPARC/Working/Initscripts/it "Handbook:SPARC/Working/Initscripts/it")[Variabili d'ambiente](/wiki/Handbook:SPARC/Working/EnvVar/it "Handbook:SPARC/Working/EnvVar/it")[Lavorare con Portage](/wiki/Handbook:SPARC/Full/Portage/it "Handbook:SPARC/Full/Portage/it")[File e cartelle](/wiki/Handbook:SPARC/Portage/Files/it "Handbook:SPARC/Portage/Files/it")[Variabili](/wiki/Handbook:SPARC/Portage/Variables/it "Handbook:SPARC/Portage/Variables/it")[Miscelare i rami del software](/wiki/Handbook:SPARC/Portage/Branches/it "Handbook:SPARC/Portage/Branches/it")[Strumenti aggiuntivi](/wiki/Handbook:SPARC/Portage/Tools/it "Handbook:SPARC/Portage/Tools/it")[Catalogo pacchetti personalizzato](/wiki/Handbook:SPARC/Portage/CustomTree/it "Handbook:SPARC/Portage/CustomTree/it")[Funzionalità avanzate](/wiki/Handbook:SPARC/Portage/Advanced/it "Handbook:SPARC/Portage/Advanced/it")[Configurare la rete](/wiki/Handbook:SPARC/Full/Networking/it "Handbook:SPARC/Full/Networking/it")[Come iniziare](/wiki/Handbook:SPARC/Networking/Introduction/it "Handbook:SPARC/Networking/Introduction/it")[Configurazione avanzata](/wiki/Handbook:SPARC/Networking/Advanced/it "Handbook:SPARC/Networking/Advanced/it")[Networking modulare](/wiki/Handbook:SPARC/Networking/Modular/it "Handbook:SPARC/Networking/Modular/it")[Wireless](/wiki/Handbook:SPARC/Networking/Wireless/it "Handbook:SPARC/Networking/Wireless/it")[Aggiungere funzionalità](/wiki/Handbook:SPARC/Networking/Extending/it "Handbook:SPARC/Networking/Extending/it")[Gestione dinamica](/wiki/Handbook:SPARC/Networking/Dynamic/it "Handbook:SPARC/Networking/Dynamic/it")

## Contents

- [1Installare il firmware e/o microcodice](#Installare_il_firmware_e.2Fo_microcodice)
  - [1.1Firmware](#Firmware)
    - [1.1.1Suggerito: Linux Firmware](#Suggerito:_Linux_Firmware)
      - [1.1.1.1Caricare il firmware](#Caricare_il_firmware)
- [2sys-kernel/installkernel](#sys-kernel.2Finstallkernel)
  - [2.1Programma di avvio (bootloader)](#Programma_di_avvio_.28bootloader.29)
    - [2.1.1GRUB](#GRUB)
    - [2.1.2Schema classico, altri programmi di avvio (e.g. (e)lilo, syslinux, etc.)](#Schema_classico.2C_altri_programmi_di_avvio_.28e.g._.28e.29lilo.2C_syslinux.2C_etc..29)
  - [2.2Initramfs](#Initramfs)
    - [2.2.1Rilevare chroot](#Rilevare_chroot)
- [3Selezionare il kernel](#Selezionare_il_kernel)
  - [3.1In alternativa: Configurazione manuale](#In_alternativa:_Configurazione_manuale)
    - [3.1.1Installare e configurare i sorgenti del kernel](#Installare_e_configurare_i_sorgenti_del_kernel)
      - [3.1.1.1Opzione 2 - Procedimento manuale assistito](#Opzione_2_-_Procedimento_manuale_assistito)
      - [3.1.1.2Opzione 3 - Configurare manualmente](#Opzione_3_-_Configurare_manualmente)
    - [3.1.2Facoltativo: Moduli del kernel firmati](#Facoltativo:_Moduli_del_kernel_firmati)

## Installare il firmware e/o microcodice

### Firmware

#### Suggerito: Linux Firmware

Su molti sistemi, è richiesto firmware non-FOSS perchè una parte dell'hardware funzioni. Il pacchetto [sys-kernel/linux-firmware](https://packages.gentoo.org/packages/sys-kernel/linux-firmware) contiene firmware per parecchi, ma non tutti, i dispositivi.

**Suggerimento**

La maggior parte delle schede wireless e grafiche necessita del firmware per funzionare.

`root #` `emerge --ask sys-kernel/linux-firmware`

**Nota**

Installare certi pacchetti firmware spesso richiede l'accettazione di licenze associate al firmware. Se necessario, visitare la [sezione di gestione delle licenze](/wiki/Handbook:SPARC/Working/Portage/it#Licenses "Handbook:SPARC/Working/Portage/it") del manuale per aiuto sull'accettazione delle licenze.

##### Caricare il firmware

Di solito i file del firmware vengono caricati al momento del caricamento del modulo del kernel associato. Ciò significa che il firmware deve essere compilato insieme al kernel usando **CONFIG\_EXTRA\_FIRMWARE** se il modulo del kernel è configurato a _Y_ piuttosto che _M_. Nella maggior parte dei casi, compilare il modulo che necessita di firmware nel kernel può complicare o interrompere il caricamento.

## sys-kernel/installkernel

[Installkernel](/wiki/Installkernel "Installkernel") può essere usato, tra le altre cose, per automatizzare l'installazione del kernel, la generazione dell' [initramfs](/wiki/Initramfs "Initramfs"), la generazione della [UKI (immagine del kernel unificata)](/wiki/Unified_kernel_image "Unified kernel image") e la configurazione del programma di avvio. Il pacchetto [sys-kernel/installkernel](https://packages.gentoo.org/packages/sys-kernel/installkernel) implementa due metodi per ottenere ciò: l'installkernel classico derivato da Debian e il kernel-install di [systemd](/wiki/Systemd/it "Systemd/it"). Quale scegliere dipende, tra le altre cose, dal programma di avvio del sistema. Di norma il kernel-install di systemd è usato su profili systemd, mentre l'installkernel classico è l'impostazione predefinita per gli altri profili.

### Programma di avvio (bootloader)

Ora è il momento di pensare a quale programma d'avvio l'utente voglia nel sistema.

**Importante**

Solo una scelta è necessaria nella seguente sottosezione, nel dubbio su quale usare scegliere la prima dell'elenco per ora. E' sempre possibile passare ad altro in seguito se necessario.

#### GRUB

Gli utenti di GRUB possono usare sia il kernel-install di systemd che l'installkernel classico di Debian. L'opzione USE [systemd](https://packages.gentoo.org/useflags/systemd) [USE flag/it (page does not exist)](/index.php?title=USE_flag/it&action=edit&redlink=1 "USE flag/it (page does not exist)") effettua il cambio tra queste implementazioni. Per eseguire automaticamente grub-mkconfig durante l'installazione del kernel, abilitare l' [opzione USE](/wiki/USE_flag "USE flag") [grub](https://packages.gentoo.org/useflags/grub) [USE flag/it (page does not exist)](/index.php?title=USE_flag/it&action=edit&redlink=1 "USE flag/it (page does not exist)").

**Nota**

GRUB ha bisogno che i kernel siano installati in /boot.

FILE **`/etc/portage/package.use/installkernel`**

```
sys-kernel/installkernel grub

```

`root #` `emerge --ask sys-kernel/installkernel`

**Nota**

systemd-boot ha bisogno che i kernel siano installati in /efi.

**Nota**

Quando si usa il pacchetto [app-emulation/virt-firmware](https://packages.gentoo.org/packages/app-emulation/virt-firmware) per configurare UEFI ci si accerti che il servizio kernel-bootcfg-boot-successful sia abilitato prima di tentare di installare il kernel. Questa implementazione dell'avvio EFI stub è l'impostazione predefinita per i sistemi systemd.

`root #` `systemctl enable kernel-bootcfg-boot-successful.service`

**Nota**

EFI stub ha bisogno che i kernel siano installati in /efi.

#### Schema classico, altri programmi di avvio (e.g. (e)lilo, syslinux, etc.)

Il classico schema /boot (p.es. (e)LILO, syslinux, etc.) è usato per impostazione predefinita se le opzioni USE [grub](https://packages.gentoo.org/useflags/grub) [USE flag/it (page does not exist)](/index.php?title=USE_flag/it&action=edit&redlink=1 "USE flag/it (page does not exist)"), [systemd-boot](https://packages.gentoo.org/useflags/systemd-boot) [USE flag/it (page does not exist)](/index.php?title=USE_flag/it&action=edit&redlink=1 "USE flag/it (page does not exist)"), [efistub](https://packages.gentoo.org/useflags/efistub) [USE flag/it (page does not exist)](/index.php?title=USE_flag/it&action=edit&redlink=1 "USE flag/it (page does not exist)") e [uki](https://packages.gentoo.org/useflags/uki) [USE flag/it (page does not exist)](/index.php?title=USE_flag/it&action=edit&redlink=1 "USE flag/it (page does not exist)") **non** sono abilitate. Nessun altra operazione è necessaria.

### Initramfs

Un **init** ial **ram**-based **f** ile **s** ystem, o [initramfs](/wiki/Initramfs "Initramfs"), può essere necessario per avviare un sistema. Un enorme varietà di casi può averne bisogno uno, ma casi comuni comprendono:

- Kernel in cui i driver dei dischi/filesystem sono moduli.
- Configurazioni con /usr/ o /var/ su partizioni separate.
- Filesystem di root criptati

**Suggerimento**

I [kernel della distribuzione](/wiki/Project:Distribution_Kernel "Project:Distribution Kernel") sono progettati per essere usati con un initramfs, poichè molti driver dei filesystem e dei dischi sono compilati come moduli.

Oltre a montare il filesystem di root, un initramfs può eseguire anche altre operazioni come:

- Eseguire **f** ile **s** ystem **c** onsistency chec **k** fsck, uno strumento per verificare e riparare la coerenza del file system nei casi di spegnimento non corretto del sistema.
- Fornire un ambiente di riparazione in caso di errori all'avvio.

[Installkernel](/wiki/Installkernel "Installkernel") può generare automaticamente un initramfs quando si installa il kernel se le opzioni USE [dracut](https://packages.gentoo.org/useflags/dracut) [USE flag/it (page does not exist)](/index.php?title=USE_flag/it&action=edit&redlink=1 "USE flag/it (page does not exist)") o [ugrd](https://packages.gentoo.org/useflags/ugrd) [USE flag/it (page does not exist)](/index.php?title=USE_flag/it&action=edit&redlink=1 "USE flag/it (page does not exist)") sono abilitate:

FILE **`/etc/portage/package.use/installkernel`**

```
sys-kernel/installkernel dracut

```

#### Rilevare chroot

Programmi di avvio come [systemd-boot](/wiki/Systemd/systemd-boot "Systemd/systemd-boot") e [EFI stub](/wiki/EFI_stub "EFI stub") usano gli argomenti del kernel di sistema in esecuzione così come impostati in /proc/cmdline per impostazione predefinita. Per via della grande quantità di metodi di installazione di Gentoo gli utenti inavvertitamente vengono tratti in errore da ciò.

Per facilitare la soluzione dei problemi che ciò può provocare, [sys-kernel/installkernel](https://packages.gentoo.org/packages/sys-kernel/installkernel) controllerà se è in esecuzione in una chroot e si interromperà se la riga di comando del kernel non è esplicitamente configurata. Altrimenti il programma di avvio userebbe gli argomenti di avvio del supporto di installazione che condurrebbe al fallimento dell'avvio.

Un modo per accontentare [sys-kernel/installkernel](https://packages.gentoo.org/packages/sys-kernel/installkernel) è usare il file di configurazione di Dracut per impostare l'UUID della partizione di root come mostrato sotto, o in alternativa per maggiori informazioni su quale aiuto ci viene da questa verifica e sui modi diversi di configurarla, vedere [Installkernel#Install\_chroot\_check](/wiki/Installkernel#Install_chroot_check.2Fit "Installkernel").

`root #` `mkdir /etc/dracut.conf.d`

`root #` `blkid`

```
/dev/sda3: UUID="2122cd72-94d7-4dcc-821e-3705926deecc"
```

Nell'esempio soprastante, la partizione di root è /dev/sda3 e l'UUID è 2122cd72-94d7-4dcc-821e-3705926deecc.
Il file di configurazione di dracut allora sarebbe:

FILE **`/etc/dracut.conf.d/00-installkernel.conf`**

```
kernel_cmdline=" root=UUID=2122cd72-94d7-4dcc-821e-3705926deecc " # Si notino gli spazi all'inizio e alla fine

```

`root #` `emerge --ask sys-kernel/installkernel`

## Selezionare il kernel

Può essere una mossa intelligente usare dist-kernel per il primo avvio dato che offre un modo molto semplice per escludere problemi al sistema e problemi alla configurazione del kernel. Avere un kernel che si sa funzionante su cui sempre ripiegare può velocizzare il debugging e alleviare l'ansia, durante un aggiornamento, per il proprio sistema che non si avvii più.

Una comune convinzione errata è che un kernel compilato manualmente userà molta meno RAM che un kernel della distribuzione pre configurato. Per via della natura modulare del kernel Linux, solo ciò che serve viene caricato e nella maggior parte dei casi si ottiene un utilizzo di memoria similare.

**Importante**

Solo una scelta è richiesta in questa sottosezione, se in dubbio su quale usare si proceda con la prima elencata per ora. E' sempre possibile fare il cambio in una data successiva se necessario.

Classificato dal meno complesso al più complesso:

[Approccio manuale](/wiki/Handbook:SPARC/Installation/Kernel/it#Alternative:_Manual_configuration "Handbook:SPARC/Installation/Kernel/it")I nuovi sorgenti del kernel sono installati attraverso il gestore dei pacchetti di sistema. Il kernel viene configurato, compilato e installato manualmente usando i comandi eselect kernel e una sfilza di make. Gli aggiornamenti futuri del kernel ripetono il procedimento manuale di configurare , compilare e installare i file del kernel. Questo è il procedimento più complesso ma offre il massimo controllo sul processo di aggiornamento del kernel.

Il nucleo su cui tutte le distribuzioni sono costruite è il kernel Linux. E' il livello (software) tra i programmi dell'utente ed il sistema fisico (hardware). Sebbene il manuale offra ai suoi utenti parecchie possibili sorgenti del kernel, un elenco più esauriente con descrizioni più dettagliate è disponibile alla [pagina dei pacchetti del kernel](/wiki/Kernel/Packages/it "Kernel/Packages/it").

**Suggerimento**

Le operazioni di installazione del kernel quali copiare l'immagine del kernel in /boot o nella [EFI System Partition](/wiki/EFI_System_Partition "EFI System Partition"), generare un [initramfs](/wiki/Initramfs "Initramfs") e/o una [Unified Kernel Image](/wiki/Unified_Kernel_Image "Unified Kernel Image"), aggiornare la configurazione del programma di avvio possono essere automatizzate con [installkernel](/wiki/Installkernel "Installkernel"). Gli utenti potrebbero voler configurare e installare [sys-kernel/installkernel](https://packages.gentoo.org/packages/sys-kernel/installkernel) prima di proseguire. Vedere la [sezione di installazione del kernel di seguito](/wiki/Handbook:SPARC/Installation/Kernel#Kernel_installation.2Fit "Handbook:SPARC/Installation/Kernel") per molte più informazioni.

### In alternativa: Configurazione manuale

#### Installare e configurare i sorgenti del kernel

Quando si compila manualmente il kernel per i sistemi basati su sparc, Gentoo consiglia il pacchetto [sys-kernel/gentoo-sources](https://packages.gentoo.org/packages/sys-kernel/gentoo-sources).

Scegliere una sorgente del kernel appropriata e installarla usando emerge:

`root #` `emerge --ask sys-kernel/gentoo-sources`

Questo comando installerà i sorgenti del kernel Linux in /usr/src/ usando la versione del kernel specifica nel percorso. Non creerà un collegamento simbolico da se stesso senza l'opzione USE [symlink](https://packages.gentoo.org/useflags/symlink) [USE flag/it (page does not exist)](/index.php?title=USE_flag/it&action=edit&redlink=1 "USE flag/it (page does not exist)") abilitata nel pacchetto dei sorgenti del kernel scelto.

E' consuetudine tenere aggiornato un collegamento simbolico di /usr/src/linux, in modo che faccia riferimento ai sorgenti corrispondenti al kernel attualmente in esecuzione. Tuttavia, questo collegamento simbolico non viene creato per impostazione predefinita. Un modo semplice per creare il collegamento simbolico è utilizzare il modulo kernel di eselect.

Per maggiori informazioni riguardo lo scopo del collegamento simbolico e come gestirlo, si prega di rivolgersi a [Kernel/Upgrade](/wiki/Kernel/Upgrade/it "Kernel/Upgrade/it").

Prima, elencare tutti i kernel installati:

`root #` `eselect kernel list`

```
Available kernel symlink targets:
  [1]   linux-6.19.3-gentoo

```

Per creare il collegamento simbolico chiamato linux, usare:

`root #` `eselect kernel set 1`

`root #` `ls -l /usr/src/linux`

```
lrwxrwxrwx    1 root   root    12 Oct 13 11:04 /usr/src/linux -> linux-6.19.3-gentoo

```

Configurare manualmente un kernel è comunemente visto come una delle procedure più complicate che un amministratore di sistema deve compiere. Niente di più falso - dopo la configurazione di alcuni kernel nessuno ricorderà che sia stato difficile! Ci sono due metodi per un utente di Gentoo di gestire manualmente un kernel di sistema, entrambi elencati di seguito:

**Importante**

Solo una scelta è richiesta in questa sottosezione, se in dubbio su quale usare si proceda con la prima elencata per ora. E' sempre possibile fare il cambio in una data successiva se necessario.

##### Opzione 2 - Procedimento manuale assistito

Questo metodo consente all'utente di avere il controllo completo su come viene compilato il proprio kernel con il minimo aiuto possibile di strumenti esterni. Alcuni potrebbero considerare ciò come renderlo difficile per il piacere di farlo.

Comunque, con questa opzione una cosa è vera: è di vitale importanza conoscere il sistema quando si configura il kernel manualmente. La maggior parte delle informazioni si possono raccogliere installando [sys-apps/pciutils](https://packages.gentoo.org/packages/sys-apps/pciutils) che contiene il comando lspci:

`root #` `emerge --ask sys-apps/pciutils`

**Nota**

Dentro la chroot (radice cambiata), si possono ignorare con sicurezza gli avvisi pcilib (come _pcilib: cannot open /sys/bus/pci/devices_) che lspci potrebbere tirar fuori.

Un'altra fonte di informazioni sul sistema è lsmod per vedere quali moduli del kernel il CD di installazione usa, poiché ciò potrebbe dare buone indizi su cosa abilitare.

Adesso si vada nella cartella del sorgente del kernel.

`root #` `cd /usr/src/linux
`

**Suggerimento**

Per vedere l'elenco completo degli argomenti di make disponibili per il kernel, eseguire `make help`.

Il kernel ha un metodo di autorilevamento dei moduli attualmente usati sul cd di installazione che fornirà un grosso punto di partenza per consentire all'utente di configurare il proprio. Si può invocare usando:

`root #` `make localmodconfig`

La configurazione manuale non dovrebbe essere necessaria a questo punto, ma se l'utente vuole controllare:

`root #` `make nconfig`

Ora è tempo di decidere se la firma dei moduli è necessaria nei passaggi elencati [qui](/wiki/Handbook:SPARC/Installation/Kernel#Optional:_Signed_kernel_modules_2.2Fit "Handbook:SPARC/Installation/Kernel").

Altrimenti, proseguire con la compilazione come descritta [qui](/wiki/Handbook:SPARC/Installation/Kernel#Compiling_and_installing.2Fit "Handbook:SPARC/Installation/Kernel").

##### Opzione 3 - Configurare manualmente

La configurazione del kernel Linux contiene molte, molte sezioni e la configurazione manuale è raramente necessaria grazie ai due strumenti elencati sopra. I passaggi per farla manualmente sono ora inclusi nella [Guida alla configurazione del kernel di Gentoo](/wiki/Kernel/Gentoo_Kernel_Configuration_Guide "Kernel/Gentoo Kernel Configuration Guide").

#### Facoltativo: Moduli del kernel firmati

Per firmare automaticamente i moduli del kernel abilitare CONFIG\_MODULE\_SIG\_ALL:

KERNEL **Firmare i moduli del kernel CONFIG\_MODULE\_SIG\_ALL**

```
[*] Enable loadable module support
  -*-   Module signature verification
    [*]     Automatically sign all modules
    Which hash algorithm should modules be signed with? (Sign modules with SHA-512) --->

```

Facoltativamente cambiare l'algoritmo di hash se desiderato.

Per rafforzare che tutti i moduli siano firmati con una firma valida, abilitare anche CONFIG\_MODULE\_SIG\_FORCE:

KERNEL **Rafforzare i moduli del kernel firmati CONFIG\_MODULE\_SIG\_FORCE**

```
[*] Enable loadable module support
  -*-   Module signature verification
    [*]     Require modules to be validly signed
    [*]     Automatically sign all modules
    Which hash algorithm should modules be signed with? (Sign modules with SHA-512) --->

```

Per usare una chiave personale, specificare la posizione di questa chiave in CONFIG\_MODULE\_SIG\_KEY. Se non si specifica, ilsistema di compilazione del kernel ne genererà una. E' consigliato invece di generarne una manualmente. Può essere fatto con:

`root #` `openssl req -new -nodes -utf8 -sha256 -x509 -outform PEM -out kernel_key.pem -keyout kernel_key.pem`

OpenSSL farà alcune domande sull'utente che genera la chiave, si consiglia di rispondere a queste domande nel modo più dettagliato possibile.

Memorizzare la chiave in un posto sicuro, come minimo la chiave dovrebbe essere leggibile solo dall'utente root. Verificarlo con:

`root #` `ls -l kernel_key.pem`

```
 -r-------- 1 root root 3164 Jan  4 10:38 kernel_key.pem
```

Se il risultato è diverso da quanto sopra riportato, correggere i permessi con:

`root #` `chown root:root kernel_key.pem
`

`root #` `chmod 400 kernel_key.pem
`

KERNEL **Specificare la chiave di firma CONFIG\_MODULE\_SIG\_KEY**

```
-*- Cryptographic API  --->
  Certificates for signature checking  --->
    (/path/to/kernel_key.pem) File name or PKCS#11 URI of module signing key

```

Per firmare anche moduli del kernel esterni installati da altri pacchetti tramite `linux-mod-r1.eclass`, abilitare l'opzione USE [modules-sign](https://packages.gentoo.org/useflags/modules-sign) [USE flag/it (page does not exist)](/index.php?title=USE_flag/it&action=edit&redlink=1 "USE flag/it (page does not exist)") globalmente:

FILE **`/etc/portage/make.conf`** **Abilitare firma dei moduli**

```
USE="modules-sign"

# Facoltativamente, quando si usano chiavi di firma personali.
MODULES_SIGN_KEY="/percorso/della/chiave_del_kernel.pem"
MODULES_SIGN_CERT="/percorso/della/chiave_del_kernel.pem" # Richiesta solo se MODULES_SIGN_KEY non contiene anche il certificato.
MODULES_SIGN_HASH="sha512" # Il valore predefinito è sha512.

```

**Nota**

MODULES\_SIGN\_KEY e MODULES\_SIGN\_CERT possono puntare a file diversi. In questo esempio, il file pem generato da OpenSSL include sia la chiave che il relativo certificato, e perciò entrambe le variabili sono impostate allo stesso valore.

[Handbook:SPARC/Blocks/Kernel/it](/index.php?title=Handbook:SPARC/Blocks/Kernel/it&action=edit&redlink=1 "Handbook:SPARC/Blocks/Kernel/it (page does not exist)")

Continuare l'installazione con [configurare il sistema](/wiki/Handbook:SPARC/Installation/System/it "Handbook:SPARC/Installation/System/it").

[← Installare il sistema base](/wiki/Handbook:SPARC/Installation/Base/it "Previous part") [Home](/wiki/Handbook:SPARC "Handbook:SPARC") [Configurare il sistema →](/wiki/Handbook:SPARC/Installation/System/it "Next part")