# Kernel

Other languages:

- [Deutsch](/wiki/Handbook:AMD64/Installation/Kernel/de "Handbook:AMD64/Installation/Kernel/de (100% translated)")
- [English](/wiki/Handbook:AMD64/Installation/Kernel "Handbook:AMD64/Installation/Kernel (100% translated)")
- [español](/wiki/Handbook:AMD64/Installation/Kernel/es "Handbook:AMD64/Installation/Kernel/es (100% translated)")
- [français](/wiki/Handbook:AMD64/Installation/Kernel/fr "Handbook:AMD64/Installation/Kernel/fr (100% translated)")
- italiano
- [magyar](/wiki/Handbook:AMD64/Installation/Kernel/hu "Handbook:AMD64/Installation/Kernel/hu (100% translated)")
- [polski](/wiki/Handbook:AMD64/Installation/Kernel/pl "Handbook:AMD64/Installation/Kernel/pl (100% translated)")
- [português do Brasil](/wiki/Handbook:AMD64/Installation/Kernel/pt-br "Handbook:AMD64/Installation/Kernel/pt-br (100% translated)")
- [čeština](/wiki/Handbook:AMD64/Installation/Kernel/cs "Handbook:AMD64/Installation/Kernel/cs (100% translated)")
- [русский](/wiki/Handbook:AMD64/Installation/Kernel/ru "Handbook:AMD64/Installation/Kernel/ru (100% translated)")
- [தமிழ்](/wiki/Handbook:AMD64/Installation/Kernel/ta "Handbook:AMD64/Installation/Kernel/ta (100% translated)")
- [中文（中国大陆）‎](/wiki/Handbook:AMD64/Installation/Kernel/zh-cn "Handbook:AMD64/Installation/Kernel/zh-cn (100% translated)")
- [日本語](/wiki/Handbook:AMD64/Installation/Kernel/ja "Handbook:AMD64/Installation/Kernel/ja (100% translated)")
- [한국어](/wiki/Handbook:AMD64/Installation/Kernel/ko "Handbook:AMD64/Installation/Kernel/ko (100% translated)")

[Manuale AMD64](/wiki/Handbook:AMD64/it "Handbook:AMD64/it")[Installazione](/wiki/Handbook:AMD64/Full/Installation/it "Handbook:AMD64/Full/Installation/it")[Riguardo l'installazione](/wiki/Handbook:AMD64/Installation/About/it "Handbook:AMD64/Installation/About/it")[Scegliere il supporto](/wiki/Handbook:AMD64/Installation/Media/it "Handbook:AMD64/Installation/Media/it")[Configurare la rete](/wiki/Handbook:AMD64/Installation/Networking/it "Handbook:AMD64/Installation/Networking/it")[Preparare i dischi](/wiki/Handbook:AMD64/Installation/Disks/it "Handbook:AMD64/Installation/Disks/it")[Il file stage](/wiki/Handbook:AMD64/Installation/Stage/it "Handbook:AMD64/Installation/Stage/it")[Installare il sistema base](/wiki/Handbook:AMD64/Installation/Base/it "Handbook:AMD64/Installation/Base/it")Configurare il kernel[Configurare il sistema](/wiki/Handbook:AMD64/Installation/System/it "Handbook:AMD64/Installation/System/it")[Installare gli strumenti](/wiki/Handbook:AMD64/Installation/Tools/it "Handbook:AMD64/Installation/Tools/it")[Impostare programma d'avvio](/wiki/Handbook:AMD64/Installation/Bootloader/it "Handbook:AMD64/Installation/Bootloader/it")[Concludere](/wiki/Handbook:AMD64/Installation/Finalizing/it "Handbook:AMD64/Installation/Finalizing/it")[Lavorare con Gentoo](/wiki/Handbook:AMD64/Full/Working/it "Handbook:AMD64/Full/Working/it")[Introduzione a Portage](/wiki/Handbook:AMD64/Working/Portage/it "Handbook:AMD64/Working/Portage/it")[Opzioni USE](/wiki/Handbook:AMD64/Working/USE/it "Handbook:AMD64/Working/USE/it")[Funzionalità di Portage](/wiki/Handbook:AMD64/Working/Features/it "Handbook:AMD64/Working/Features/it")[Sistema script di init](/wiki/Handbook:AMD64/Working/Initscripts/it "Handbook:AMD64/Working/Initscripts/it")[Variabili d'ambiente](/wiki/Handbook:AMD64/Working/EnvVar/it "Handbook:AMD64/Working/EnvVar/it")[Lavorare con Portage](/wiki/Handbook:AMD64/Full/Portage/it "Handbook:AMD64/Full/Portage/it")[File e cartelle](/wiki/Handbook:AMD64/Portage/Files/it "Handbook:AMD64/Portage/Files/it")[Variabili](/wiki/Handbook:AMD64/Portage/Variables/it "Handbook:AMD64/Portage/Variables/it")[Miscelare i rami del software](/wiki/Handbook:AMD64/Portage/Branches/it "Handbook:AMD64/Portage/Branches/it")[Strumenti aggiuntivi](/wiki/Handbook:AMD64/Portage/Tools/it "Handbook:AMD64/Portage/Tools/it")[Catalogo pacchetti personalizzato](/wiki/Handbook:AMD64/Portage/CustomTree/it "Handbook:AMD64/Portage/CustomTree/it")[Funzionalità avanzate](/wiki/Handbook:AMD64/Portage/Advanced/it "Handbook:AMD64/Portage/Advanced/it")[Configurare la rete](/wiki/Handbook:AMD64/Full/Networking/it "Handbook:AMD64/Full/Networking/it")[Come iniziare](/wiki/Handbook:AMD64/Networking/Introduction/it "Handbook:AMD64/Networking/Introduction/it")[Configurazione avanzata](/wiki/Handbook:AMD64/Networking/Advanced/it "Handbook:AMD64/Networking/Advanced/it")[Networking modulare](/wiki/Handbook:AMD64/Networking/Modular/it "Handbook:AMD64/Networking/Modular/it")[Wireless](/wiki/Handbook:AMD64/Networking/Wireless/it "Handbook:AMD64/Networking/Wireless/it")[Aggiungere funzionalità](/wiki/Handbook:AMD64/Networking/Extending/it "Handbook:AMD64/Networking/Extending/it")[Gestione dinamica](/wiki/Handbook:AMD64/Networking/Dynamic/it "Handbook:AMD64/Networking/Dynamic/it")

## Contents

- [1Installare il firmware e/o microcodice](#Installare_il_firmware_e.2Fo_microcodice)
  - [1.1Firmware](#Firmware)
    - [1.1.1Suggerito: Linux Firmware](#Suggerito:_Linux_Firmware)
      - [1.1.1.1Caricare il firmware](#Caricare_il_firmware)
    - [1.1.2SOF Firmware](#SOF_Firmware)
  - [1.2Suggerito: Microcodice](#Suggerito:_Microcodice)
- [2sys-kernel/installkernel](#sys-kernel.2Finstallkernel)
  - [2.1Programma di avvio (bootloader)](#Programma_di_avvio_.28bootloader.29)
    - [2.1.1GRUB](#GRUB)
    - [2.1.2systemd-boot](#systemd-boot)
    - [2.1.3EFI stub](#EFI_stub)
    - [2.1.4Schema classico, altri programmi di avvio (e.g. (e)lilo, syslinux, etc.)](#Schema_classico.2C_altri_programmi_di_avvio_.28e.g._.28e.29lilo.2C_syslinux.2C_etc..29)
  - [2.2Initramfs](#Initramfs)
    - [2.2.1Rilevare chroot](#Rilevare_chroot)
  - [2.3Facoltativo: Unified Kernel Image o UKI](#Facoltativo:_Unified_Kernel_Image_o_UKI)
    - [2.3.1UKI generica (solo systemd)](#UKI_generica_.28solo_systemd.29)
    - [2.3.2Secure Boot](#Secure_Boot)
- [3Selezionare il kernel](#Selezionare_il_kernel)
  - [3.1Kernel della distribuzione](#Kernel_della_distribuzione)
    - [3.1.1Facoltativo: Moduli del kernel firmati](#Facoltativo:_Moduli_del_kernel_firmati)
    - [3.1.2Facoltativo: Firmare l'immagine del kernel (Secure Boot)](#Facoltativo:_Firmare_l.27immagine_del_kernel_.28Secure_Boot.29)
    - [3.1.3Installare un kernel della distribuzione](#Installare_un_kernel_della_distribuzione)
    - [3.1.4Aggiornare e fare pulizia](#Aggiornare_e_fare_pulizia)
    - [3.1.5Attività post-installazione/aggiornamento](#Attivit.C3.A0_post-installazione.2Faggiornamento)
      - [3.1.5.1Ricompilare manualmente l'initramfs o lo Unified Kernel Image](#Ricompilare_manualmente_l.27initramfs_o_lo_Unified_Kernel_Image)
  - [3.2In alternativa: Configurazione manuale](#In_alternativa:_Configurazione_manuale)
    - [3.2.1Installare e configurare i sorgenti del kernel](#Installare_e_configurare_i_sorgenti_del_kernel)
      - [3.2.1.1Opzione 1 - Procedimento modprobed-db](#Opzione_1_-_Procedimento_modprobed-db)
      - [3.2.1.2Opzione 2 - Procedimento manuale assistito](#Opzione_2_-_Procedimento_manuale_assistito)
      - [3.2.1.3Opzione 3 - Configurare manualmente](#Opzione_3_-_Configurare_manualmente)
    - [3.2.2Facoltativo: Moduli del kernel firmati](#Facoltativo:_Moduli_del_kernel_firmati_2)
    - [3.2.3Facoltativo: Firmare l'immagine del kernel (Secure Boot)](#Facoltativo:_Firmare_l.27immagine_del_kernel_.28Secure_Boot.29_2)
    - [3.2.4Compilare e installare](#Compilare_e_installare)

## Installare il firmware e/o microcodice

### Firmware

#### Suggerito: Linux Firmware

Su molti sistemi, è richiesto firmware non-FOSS perchè una parte dell'hardware funzioni. Il pacchetto [sys-kernel/linux-firmware](https://packages.gentoo.org/packages/sys-kernel/linux-firmware) contiene firmware per parecchi, ma non tutti, i dispositivi.

**Suggerimento**

La maggior parte delle schede wireless e grafiche necessita del firmware per funzionare.

`root #` `emerge --ask sys-kernel/linux-firmware`

**Nota**

Installare certi pacchetti firmware spesso richiede l'accettazione di licenze associate al firmware. Se necessario, visitare la [sezione di gestione delle licenze](/wiki/Handbook:AMD64/Working/Portage/it#Licenses "Handbook:AMD64/Working/Portage/it") del manuale per aiuto sull'accettazione delle licenze.

##### Caricare il firmware

Di solito i file del firmware vengono caricati al momento del caricamento del modulo del kernel associato. Ciò significa che il firmware deve essere compilato insieme al kernel usando **CONFIG\_EXTRA\_FIRMWARE** se il modulo del kernel è configurato a _Y_ piuttosto che _M_. Nella maggior parte dei casi, compilare il modulo che necessita di firmware nel kernel può complicare o interrompere il caricamento.

#### SOF Firmware

Sound Open Firmware (SOF) is a new open source audio driver meant to replace the proprietary Smart Sound Technology (SST) audio driver from Intel. 10th gen+ and Apollo Lake (Atom E3900, Celeron N3350, and Pentium N4200) Intel CPUs require this firmware for certain features and certain AMD APUs also have support for this firmware. SOF's supported platforms matrix can be found [here](https://thesofproject.github.io/latest/platforms/index.html) for more information.

`root #` `emerge --ask sys-firmware/sof-firmware`

### Suggerito: Microcodice

Oltre alle schede grafiche separate e alle interfacce di rete, anche le CPU possono aver bisogno di aggiornamenti del firmware. In genere ci si riferisce a questo tipo di firmware come _microcodice_. Versioni più nuove del microcodice sono a volte neccessarie per correggere instabilità, questioni di sicurezza o altri errori vari nell'hardware della CPU.

Aggiornamenti del microcodice per CPU AMD sono distribuiti all'interno del sopracitato pacchetto [sys-kernel/linux-firmware](https://packages.gentoo.org/packages/sys-kernel/linux-firmware). Il microcodice per le CPU Intel si può trovare all'interno del pacchetto [sys-firmware/intel-microcode](https://packages.gentoo.org/packages/sys-firmware/intel-microcode), che bisognerà installare separatamente. Vedere l'articolo [microcodice](/wiki/Microcode "Microcode") per maggiori informazioni su come applicare aggiornamenti del microcodice.

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

#### systemd-boot

Quando si usa [systemd-boot](/wiki/Systemd/systemd-boot "Systemd/systemd-boot") (in precedenza gummiboot) come programma di avvio, deve essere usato kernel-install di systemd. Quindi ci si assicuri che le opzioni USE [systemd](https://packages.gentoo.org/useflags/systemd) [USE flag/it (page does not exist)](/index.php?title=USE_flag/it&action=edit&redlink=1 "USE flag/it (page does not exist)") e [systemd-boot](https://packages.gentoo.org/useflags/systemd-boot) [USE flag/it (page does not exist)](/index.php?title=USE_flag/it&action=edit&redlink=1 "USE flag/it (page does not exist)") siano abilitate nel pacchetto [sys-kernel/installkernel](https://packages.gentoo.org/packages/sys-kernel/installkernel), e poi installare il pacchetto relativo a systemd-boot.

FILE **`/etc/portage/package.use/systemd-boot`**

```
sys-apps/systemd boot
sys-apps/systemd-utils boot kernel-install
sys-kernel/installkernel systemd systemd-boot

```

La riga di comando del kernel da usare per i nuovi kernel dovrebbe essere specificata in /etc/kernel/cmdline e /etc/cmdline. Poichè entrambi hanno bisogno di essere aggiornati insieme, ha senso collegare simbolicamente i due file ora.

FILE **`/etc/kernel/cmdline`**

```
quiet splash

```

`root #` `mkdir -p /etc/cmdline.d`

`root #` `ln -s /etc/kernel/cmdline /etc/cmdline.d/00-installkernel.conf`

`root #` `emerge --ask sys-kernel/installkernel`

**Nota**

systemd-boot ha bisogno che i kernel siano installati in /efi.

#### EFI stub

I sistemi di computer basati su UEFI tecnicamente non hanno bisogno di un programma di avvio secondario per avviare il kernel. I programmi di avvio secondario esistono per _estendere_ le funzionalità del firmware UEFI durante la fase di avvio. Detto ciò, usare un programma di avvio secondario è normalmente più facile e robusto poichè offre un approccio più flessibile per modificare velocemente i parametri del kernel all'avvio. Si noti inoltre che le implementazioni di UEFI differiscono di molto tra i fornitori e tra i modelli e non è garantito che un certo firmware segua le specifiche UEFI. Quindi, non si garantisce che l'avvio EFI stub funzioni su ogni sistema basato su UEFI.

`root #` `emerge --ask sys-kernel/installkernel`

`root #` `mkdir -p /efi/EFI/Gentoo`

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

### Facoltativo: Unified Kernel Image o UKI

Una [UKI](/wiki/Unified_Kernel_Image "Unified Kernel Image") (immagine del kernel unificata) combina, tra le altre cose, il kernel, l'initramfs e la riga di comando del kernel in un singolo eseguibile. Poichè la riga di comando del kernel è integrata nella UKI, dovrebbe essere specificata prima di generarla (vedere sotto). Si noti che qualsiasi riga di comando del kernel fornita dal programma di avvio o dal firmware all'avvio è ignorata quando si avvia col secure boot abilitato.

Una UKI necessita di un caricatore stub. Al momento, l'unico disponibile è systemd-stub. Per abilitarlo:

Su sistemi systemd:

FILE **`/etc/portage/package.use/uki`**

```
sys-apps/systemd boot

```

`root #` `emerge --ask sys-apps/systemd`

Su sistemi OpenRC:

FILE **`/etc/portage/package.use/uki`**

```
sys-apps/systemd-utils boot kernel-install

```

`root #` `emerge --ask sys-apps/systemd-utils`

[Installkernel](/wiki/Installkernel "Installkernel") può generare una UKI automaticamente usando sia [dracut](/wiki/Unified_kernel_image#dracut.2Fit "Unified kernel image") che [ukify](/wiki/Unified_kernel_image#ukify.2Fit "Unified kernel image") abilitando le rispettive opzioni e l'opzione USE [uki](https://packages.gentoo.org/useflags/uki) [USE flag/it (page does not exist)](/index.php?title=USE_flag/it&action=edit&redlink=1 "USE flag/it (page does not exist)").

Per dracut:

FILE **`/etc/portage/package.use/uki`**

```
sys-kernel/installkernel dracut uki

```

FILE **`/etc/dracut.conf.d/uki.conf`**

```
uefi="yes"
kernel_cmdline="some-kernel-command-line-arguments"

```

`root #` `emerge --ask sys-kernel/installkernel`

Per ukify:

FILE **`/etc/portage/package.use/uki`**

```
sys-apps/systemd boot ukify                         # Per sistemi systemd
sys-apps/systemd-utils kernel-install boot ukify    # Per sistemi OpenRC
sys-kernel/installkernel dracut ukify uki

```

FILE **`/etc/kernel/cmdline`**

```
some-kernel-command-line-arguments

```

`root #` `emerge --ask sys-kernel/installkernel`

Si noti che mentre dracut può generare sia un initramfs che una UKI, ukify può generare solo l'ultima e di conseguenza l'initramfs deve essere generato separatamente con dracut.

**Importante**

Negli esempi di configurazione sopra riportati (sia per dracut che unify) è importante specificare almeno un parametro root= appropriato per la riga di comando del kernel per assicurare che la UKI trovi la partizione di root. Ciò non è necessario sui sistemi systemd che seguono la DPS (Discoverable Partitions Specification), nel qual caso l'initramfs integrato sarà capace di trovare dinamicamente la partizione di root.

#### UKI generica (solo systemd)

Il pacchetto [sys-kernel/gentoo-kernel-bin](https://packages.gentoo.org/packages/sys-kernel/gentoo-kernel-bin) precompilato può opzionalmente installare una UKI generica precompilata contenente un initramfs generico capace di avviare la maggior parte dei sistemi basati su systemd. Può essere installato abilitando l'opzione USE [generic-uki](https://packages.gentoo.org/useflags/generic-uki) [USE flag/it (page does not exist)](/index.php?title=USE_flag/it&action=edit&redlink=1 "USE flag/it (page does not exist)"), e configurando [installkernel](/wiki/Installkernel "Installkernel") in modo che non generi un initramfs personalizzato o una UKI:

FILE **`/etc/portage/package.use/uki`**

```
sys-kernel/gentoo-kernel-bin generic-uki
sys-kernel/installkernel -dracut -ukify -ugrd uki

```

#### Secure Boot

**Attenzione**

Se si segue questa sezione e si compila il proprio kernel manualmente accertarsi di seguire i passaggi delineati in [Firmare il kernel](/wiki/Kernel#Optional:_Signing_the_kernel_image_.28Secure_Boot.29.2Fit "Kernel")

La UKI generica opzionalmente distribuita dal pacchetto [sys-kernel/gentoo-kernel-bin](https://packages.gentoo.org/packages/sys-kernel/gentoo-kernel-bin) è già pre-firmata. Come firmare una UKI generata localmente dipende dall'utilizzo di dracut o ukify. Si noti che la posizione della chiave e del certificato dovrebbe essere la stessa di SECUREBOOT\_SIGN\_KEY e SECUREBOOT\_SIGN\_CERT come specificato in /etc/portage/make.conf.

Per dracut:

FILE **`/etc/dracut.conf.d/uki.conf`**

```
uefi="yes"
kernel_cmdline="some-kernel-command-line-arguments"
uefi_secureboot_key="/percorso/della/chiave_del_kernel.pem"
uefi_secureboot_cert="/percorso/della/chiave_del_kernel.pem"

```

Per ukify:

FILE **`/etc/kernel/uki.conf`**

```
[UKI]
SecureBootPrivateKey=/percorso/della/chiave_del_kernel.pem
SecureBootCertificate=/percorso/della/chiave_del_kernel.pem

```

## Selezionare il kernel

Può essere una mossa intelligente usare dist-kernel per il primo avvio dato che offre un modo molto semplice per escludere problemi al sistema e problemi alla configurazione del kernel. Avere un kernel che si sa funzionante su cui sempre ripiegare può velocizzare il debugging e alleviare l'ansia, durante un aggiornamento, per il proprio sistema che non si avvii più.

Una comune convinzione errata è che un kernel compilato manualmente userà molta meno RAM che un kernel della distribuzione pre configurato. Per via della natura modulare del kernel Linux, solo ciò che serve viene caricato e nella maggior parte dei casi si ottiene un utilizzo di memoria similare.

**Importante**

Solo una scelta è richiesta in questa sottosezione, se in dubbio su quale usare si proceda con la prima elencata per ora. E' sempre possibile fare il cambio in una data successiva se necessario.

Classificato dal meno complesso al più complesso:

[Approccio completamente automatico: kernel della distribuzione](/wiki/Handbook:AMD64/Installation/Kernel/it#Distribution_kernels "Handbook:AMD64/Installation/Kernel/it")Un [kernel della distribuzione](/wiki/Project:Distribution_Kernel "Project:Distribution Kernel") è usato per configurare, compilare automaticamente ed installare il kernel Linux, i suoi moduli associati e (facoltativamente, ma per impostazione predefinita) un file initramfs. Gli aggiornamenti futuri del kernel sono completamente automatizzati poichè se ne occupa il gestore dei pacchetti, come qualsiasi altro pacchetto del sistema. E' possibile [fornire un file di configurazione del kernel su misura](/wiki/Project:Distribution_Kernel#Modifying_kernel_configuration "Project:Distribution Kernel") se è necessaria la personalizzazione. Questo è il procedimento meno complesso ed è perfetto per i nuovi utenti di Gentoo per via del suo essere pronto all'uso e del suo offrire una interazione minima da parte dell'amministratore di sistema.[Approccio manuale](/wiki/Handbook:AMD64/Installation/Kernel/it#Alternative:_Manual_configuration "Handbook:AMD64/Installation/Kernel/it")I nuovi sorgenti del kernel sono installati attraverso il gestore dei pacchetti di sistema. Il kernel viene configurato, compilato e installato manualmente usando i comandi eselect kernel e una sfilza di make. Gli aggiornamenti futuri del kernel ripetono il procedimento manuale di configurare , compilare e installare i file del kernel. Questo è il procedimento più complesso ma offre il massimo controllo sul processo di aggiornamento del kernel.

Il nucleo su cui tutte le distribuzioni sono costruite è il kernel Linux. E' il livello (software) tra i programmi dell'utente ed il sistema fisico (hardware). Sebbene il manuale offra ai suoi utenti parecchie possibili sorgenti del kernel, un elenco più esauriente con descrizioni più dettagliate è disponibile alla [pagina dei pacchetti del kernel](/wiki/Kernel/Packages/it "Kernel/Packages/it").

**Suggerimento**

Le operazioni di installazione del kernel quali copiare l'immagine del kernel in /boot o nella [EFI System Partition](/wiki/EFI_System_Partition "EFI System Partition"), generare un [initramfs](/wiki/Initramfs "Initramfs") e/o una [Unified Kernel Image](/wiki/Unified_Kernel_Image "Unified Kernel Image"), aggiornare la configurazione del programma di avvio possono essere automatizzate con [installkernel](/wiki/Installkernel "Installkernel"). Gli utenti potrebbero voler configurare e installare [sys-kernel/installkernel](https://packages.gentoo.org/packages/sys-kernel/installkernel) prima di proseguire. Vedere la [sezione di installazione del kernel di seguito](/wiki/Handbook:AMD64/Installation/Kernel#Kernel_installation.2Fit "Handbook:AMD64/Installation/Kernel") per molte più informazioni.

### Kernel della distribuzione

I _[kernel della distribuzione](/wiki/Project:Distribution_Kernel "Project:Distribution Kernel")_ sono ebuilds che si occupano dell' intero processo di estrarre, configurare, compilare, e installare il kernel. Il vantaggio principale di questo metodo è che i kernel vengono aggiornati alle nuove versioni dal gestore dei pacchetti insieme all'aggiornamento dell' insieme @world. Ciò non richiede altro coinvolgimento che eseguire un comando emerge. I kernel della distribuzione hanno una configurazione predefinita compatibile con la maggioranza dell'hardware, tuttavia vengono offerti due meccanismi per la personalizzazione: savedconfig e frammenti di config. Vedere la pagina del progetto per [maggiori dettagli sulla configurazione.](/wiki/Project:Distribution_Kernel#Modifying_kernel_configuration "Project:Distribution Kernel")

##### Facoltativo: Moduli del kernel firmati

I moduli del kernel nel kernel precompilato della distribuzione ([sys-kernel/gentoo-kernel-bin](https://packages.gentoo.org/packages/sys-kernel/gentoo-kernel-bin)) sono già firmati. Per firmare i moduli dei kernel compilati dal codice sorgente abilitare l'opzione USE [modules-sign](https://packages.gentoo.org/useflags/modules-sign) [USE flag/it (page does not exist)](/index.php?title=USE_flag/it&action=edit&redlink=1 "USE flag/it (page does not exist)"), e opzionalmente specificare quale chiave usare per la firma in [/etc/portage/make.conf](/wiki//etc/portage/make.conf/it "/etc/portage/make.conf/it"):

FILE **`/etc/portage/make.conf`** **Abilitare la firma dei moduli**

```
USE="modules-sign"

# Facoltativamente, per usare chiavi di firma personali.
MODULES_SIGN_KEY="/percorso/della/chiave_del_kernel.pem"
MODULES_SIGN_CERT="/percorso/della/chiave_del_kernel.pem" # Richiesta solo se MODULES_SIGN_KEY non contiene anche il certificato.
MODULES_SIGN_HASH="sha512" # Il valore predefinito è sha512.

```

Se la MODULES\_SIGN\_KEY non è specificata il sistema di compilazione del kernel genererà una chiave che sarà salvata in /usr/src/linux-x.y.z/certs. E' consigliato generare manualmente una chiave per assicurarsi che sia la stessa per ogni versione del kernel. Una chiave può essere generata con:

`root #` `openssl req -new -noenc -utf8 -sha256 -x509 -outform PEM -out kernel_key.pem -keyout kernel_key.pem`

**Nota**

MODULES\_SIGN\_KEY e MODULES\_SIGN\_CERT possono essere file diversi. In questo esempio il file pem generato da OpenSSL include sia la chiave che il relativo certificato, e perciò entrambe le variabili sono impostate allo stesso valore.

OpenSSL farà alcune domande sull'utente che genera la chiave, si consiglia di rispondere a queste domande nel modo più dettagliato possibile.

Memorizzare la chiave in un posto sicuro, come minimo la chiave dovrebbe essere leggibile solo dall'utente root. Verificarlo con:

`root #` `ls -l kernel_key.pem`

```
 -r-------- 1 root root 3164 Jan  4 10:38 kernel_key.pem
```

Se il risultato è diverso da quanto sopra riportato, correggere i permessi con:

`root #` `chown root:root kernel_key.pem`

`root #` `chmod 400 kernel_key.pem`

##### Facoltativo: Firmare l'immagine del kernel (Secure Boot)

L' immagine del kernel nel kernel precompilato della distribuzione ([sys-kernel/gentoo-kernel-bin](https://packages.gentoo.org/packages/sys-kernel/gentoo-kernel-bin)) è già firmato per l'uso con [Secure Boot](/wiki/Secure_Boot "Secure Boot"). Per firmare l'immagine del kernel dei kernel compilati dal codice sorgente abilitare l'opzione USE [secureboot](https://packages.gentoo.org/useflags/secureboot) [USE flag/it (page does not exist)](/index.php?title=USE_flag/it&action=edit&redlink=1 "USE flag/it (page does not exist)"), e facoltativamente specificare quale chiave usare per la firma in [/etc/portage/make.conf](/wiki//etc/portage/make.conf/it "/etc/portage/make.conf/it"). Si noti che firmare l'immmagine del kernel per l'uso con secureboot richiede che anche i moduli siano firmati, si può usare la stessa chiave per firmare sia l'immagine che i moduli del kernel:

FILE **`/etc/portage/make.conf`** **Abilitare chiavi di firma personali**

```
USE="modules-sign secureboot"

# Facoltativamente, per usare chiavi di firma personali.
MODULES_SIGN_KEY="/percorso/della/chiave_del_kernel.pem"
MODULES_SIGN_CERT="/percorso/della/chiave_del_kernel.pem" # Richiesta solo se MODULES_SIGN_KEY non contiene anche il certificato.
MODULES_SIGN_HASH="sha512" # Il valore predefinito è sha512.

# Facoltativamente, per avviare con secureboot abilitato; può essere la stessa o una diversa chiave di firma.
SECUREBOOT_SIGN_KEY="/percorso/della/chiave_del_kernel.pem"
SECUREBOOT_SIGN_CERT="/percorso/della/chiave_del_kernel.pem"

```

**Nota**

SECUREBOOT\_SIGN\_KEY e SECUREBOOT\_SIGN\_CERT possono essere file diversi. In questo esempio il file pem generato da OpenSSL include sia la chiave che il relativo certificato, e perciò entrambe le variabili sono impostate allo stesso valore.

**Nota**

In questo esempio la stessa chiave che è stata generata per firmare i moduli è utilzzata per firmare l'immagine del kernel. E' anche possibile generare e usare una seconda chiave separata per firmare l'immagine del kernel. Può essere usato di nuovo lo stesso comando OpenSSL della sezione precedente.

Vedere la sezione soprastante per istruzioni sulla generazione di una nuova chiave, i passaggi possono essere ripetuti se si dovesse usare una chiave separata per firmare l'immagine del kernel.

Per avere successo nell'avvio con Secure Boot abilitato, anche il programma di avvio utilizzato deve essere firmato e il certificato deve essere accettato dal firmware [UEFI](/wiki/UEFI "UEFI") o [Shim](/wiki/Shim "Shim"). Ciò verrà spiegato più avanti nel manuale.

#### Installare un kernel della distribuzione

Per compilare un kernel con modifiche di Gentoo dal codice sorgente, digitare:

`root #` `emerge --ask sys-kernel/gentoo-kernel`

Gli amministratori di sistema che vogliono evitare di compilare i sorgenti del kernel localmente possono invece usare immagini del kernel precompilate:

`root #` `emerge --ask sys-kernel/gentoo-kernel-bin`

**Importante**

I _[kernel della distribuzione](/wiki/Project:Distribution_Kernel "Project:Distribution Kernel")_, come [sys-kernel/gentoo-kernel](https://packages.gentoo.org/packages/sys-kernel/gentoo-kernel) e [sys-kernel/gentoo-kernel-bin](https://packages.gentoo.org/packages/sys-kernel/gentoo-kernel-bin), di norma, si aspettano di essere installati insieme ad [initramfs](/wiki/Handbook:AMD64/Installation/Kernel#Initramfs.2Fit "Handbook:AMD64/Installation/Kernel"). Prima di fare l'emerge per installare il kernel gli utenti dovrebbero accertarsi che [sys-kernel/installkernel](https://packages.gentoo.org/packages/sys-kernel/installkernel) sia stato configurato per utilizzare un generatore di initramfs (per esempio [Dracut](/wiki/Dracut "Dracut")) come descritto nella sezione [installkernel](/wiki/Handbook:AMD64/Installation/Kernel#Initramfs.2Fit "Handbook:AMD64/Installation/Kernel").

#### Aggiornare e fare pulizia

Una volta installato il kernel, il gestore di pacchetti lo aggiornerà automaticamente alle versioni più recenti. Le versioni precedenti saranno mantenute finchè non sarà richiesto al gestore di pacchetti di fare pulizia dei pacchetti superati. Per recuperare spazio su disco, i pacchetti superati possono essere ridotti eseguendo periodicamente emerge con l'opzione `--depclean`:

`root #` `emerge --depclean`

In alternativa, per fare pulizia specificamente delle vecchie versioni del kernel:

`root #` `emerge --prune sys-kernel/gentoo-kernel sys-kernel/gentoo-kernel-bin`

**Suggerimento**

Per scelta di progetto, emerge rimuove solo la cartella della compilazione del kernel. Sorprendentemente non rimuove i moduli del kernel, nè l'immagine del kernel installata. Per fare pulizia completamente dei vecchi kernel, si può usare lo strumento [app-admin/eclean-kernel](https://packages.gentoo.org/packages/app-admin/eclean-kernel).

#### Attività post-installazione/aggiornamento

Un aggiornamento del kernel della distribuzione è capace di innescare una ricompilazione automatica dei moduli esterni del kernel installati da altri pacchetti (per esempio: [sys-fs/zfs-kmod](https://packages.gentoo.org/packages/sys-fs/zfs-kmod) o [x11-drivers/nvidia-drivers](https://packages.gentoo.org/packages/x11-drivers/nvidia-drivers)). Questo comportamento automatizzato viene attivato abilitando l'opzione USE [dist-kernel](https://packages.gentoo.org/useflags/dist-kernel) [USE flag/it (page does not exist)](/index.php?title=USE_flag/it&action=edit&redlink=1 "USE flag/it (page does not exist)"). Quando richiesto, la stessa opzione innescherà anche la ri-generazione dell' [initramfs](/wiki/Initramfs "Initramfs").

E' altamente consigliato abilitare questa opzione globalmente tramite /etc/portage/make.conf quando si usa un kernel della distribuzione:

FILE **`/etc/portage/make.conf`** **Abilitare USE=dist-kernel**

```
USE="dist-kernel"

```

##### Ricompilare manualmente l'initramfs o lo Unified Kernel Image

Se necessario, innescare manualmente tali ricompilazioni, dopo un aggiornamento, eseguendo:

`root #` `emerge --ask @module-rebuild`

Se qualsiasi modulo del kernel (p.es. ZFS) è necessario presto durante l'avvio, ricompilare l'initramfs successivamente con:

`root #` `emerge --config sys-kernel/gentoo-kernel
`

`root #` `emerge --config sys-kernel/gentoo-kernel-bin
`

Dopo aver installato con successo il kernel della distribuzione, è ora il momento di procedere con la prossima sezione: [Configurare il sistema](/wiki/Handbook:AMD64/Installation/System/it "Handbook:AMD64/Installation/System/it").

### In alternativa: Configurazione manuale

#### Installare e configurare i sorgenti del kernel

Quando si compila manualmente il kernel per i sistemi basati su amd64, Gentoo consiglia il pacchetto [sys-kernel/gentoo-sources](https://packages.gentoo.org/packages/sys-kernel/gentoo-sources).

Scegliere una sorgente del kernel appropriata e installarla usando emerge:

`root #` `emerge --ask sys-kernel/gentoo-sources`

Questo comando installerà i sorgenti del kernel Linux in /usr/src/ usando la versione del kernel specifica nel percorso. Non creerà un collegamento simbolico da se stesso senza l'opzione USE [symlink](https://packages.gentoo.org/useflags/symlink) [USE flag/it (page does not exist)](/index.php?title=USE_flag/it&action=edit&redlink=1 "USE flag/it (page does not exist)") abilitata nel pacchetto dei sorgenti del kernel scelto.

E' consuetudine tenere aggiornato un collegamento simbolico di /usr/src/linux, in modo che faccia riferimento ai sorgenti corrispondenti al kernel attualmente in esecuzione. Tuttavia, questo collegamento simbolico non viene creato per impostazione predefinita. Un modo semplice per creare il collegamento simbolico è utilizzare il modulo kernel di eselect.

Per maggiori informazioni riguardo lo scopo del collegamento simbolico e come gestirlo, si prega di rivolgersi a [Kernel/Upgrade](/wiki/Kernel/Upgrade/it "Kernel/Upgrade/it").

Prima, elencare tutti i kernel installati:

`root #` `eselect kernel list`

```
Available kernel symlink targets:
  [1]   linux-6.18.8-gentoo

```

Per creare il collegamento simbolico chiamato linux, usare:

`root #` `eselect kernel set 1`

`root #` `ls -l /usr/src/linux`

```
lrwxrwxrwx    1 root   root    12 Oct 13 11:04 /usr/src/linux -> linux-6.18.8-gentoo

```

Configurare manualmente un kernel è comunemente visto come una delle procedure più complicate che un amministratore di sistema deve compiere. Niente di più falso - dopo la configurazione di alcuni kernel nessuno ricorderà che sia stato difficile! Ci sono due metodi per un utente di Gentoo di gestire manualmente un kernel di sistema, entrambi elencati di seguito:

**Importante**

Solo una scelta è richiesta in questa sottosezione, se in dubbio su quale usare si proceda con la prima elencata per ora. E' sempre possibile fare il cambio in una data successiva se necessario.

##### Opzione 1 - Procedimento modprobed-db

Un modo molto semplice di gestire il kernel è installare prima [sys-kernel/gentoo-kernel-bin](https://packages.gentoo.org/packages/sys-kernel/gentoo-kernel-bin) e usare [sys-kernel/modprobed-db](https://packages.gentoo.org/packages/sys-kernel/modprobed-db) per raccogliere informazioni su ciò che il sistema richiede. modprobed-db è uno strumento che sorveglia il sistema attraverso crontab per aggiungere tutti i moduli di tutti i dispositivi durante la vita del sistema per assicurarsi che qualsiasi cosa l'utente abbia bisogno sia supportata. Per esempio, se viene aggiunto un controller Xbox dopo l'installazione, allora modprobed-db aggiungerà i moduli da compilare la prossima volta che il kernel sarà ricompilato.

Intanto si prega di eseguire l'installazione di [sys-kernel/gentoo-kernel-bin](https://packages.gentoo.org/packages/sys-kernel/gentoo-kernel-bin) come descritto in [kernel della distribuzione](/wiki/Handbook:AMD64/Installation/Kernel#Distribution_kernels.2Fit "Handbook:AMD64/Installation/Kernel").

Successivamente, installare [sys-kernel/modprobed-db](https://packages.gentoo.org/packages/sys-kernel/modprobed-db):

`root #` `emerge --ask sys-kernel/modprobed-db`

Si prega di fare attenzione per ulteriori passaggi relativi a modprobed-db nel manuale.

Di più su questo argomento si può trovare nell'articolo [Modprobed-db](/wiki/Modprobed-db "Modprobed-db").

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

Ora è tempo di decidere se la firma dei moduli è necessaria nei passaggi elencati [qui](/wiki/Handbook:AMD64/Installation/Kernel#Optional:_Signed_kernel_modules_2.2Fit "Handbook:AMD64/Installation/Kernel").

Altrimenti, proseguire con la compilazione come descritta [qui](/wiki/Handbook:AMD64/Installation/Kernel#Compiling_and_installing.2Fit "Handbook:AMD64/Installation/Kernel").

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

#### Facoltativo: Firmare l'immagine del kernel (Secure Boot)

Quando si firma l'immagine del kernel (da usare su sistemi con [Secure Boot](/wiki/Secure_Boot "Secure Boot") abilitato) è consigliato impostare le seguenti opzioni di configurazione del kernel:

KERNEL **Blocco per secureboot**

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

Dove ""image"" è un segnaposto per il nome dell'immagine per la specifica architettura. Queste opzioni, dall'alto in basso: impongono che l'immagine del kernel in una chiamata kexec debba essere firmata (kexec consente la sostituzione del kernel sul posto), impongono che i moduli del kernel siano firmati, abilitano la modalità integrity del blocco (previene modifiche al kernel mentre è in esecuzione) e abilitano vari portachiavi.

Su architetture che non supportano nativamente la decompressione del kernel (p.es. **arm64** e **riscv**), il kernel deve essere compilato col proprio decompressore (zboot):

KERNEL **zboot CONFIG\_EFI\_ZBOOT**

```
Device Drivers --->
  Firmware Drivers --->
    EFI (Extensible Firmware Interface) Support --->
      [*] Enable the generic EFI decompressor

```

Dopo la compilazione del kernel, come spiegato nella prossima sezione, l'immagine del kernel deve essere firmata. Prima installare [app-crypt/sbsigntools](https://packages.gentoo.org/packages/app-crypt/sbsigntools) e poi firmare l'immagine del kernel:

`root #` `emerge --ask app-crypt/sbsigntools`

`root #` `sbsign /usr/src/linux-x.y.z/percorso/della/kernel-image --cert /percorso/del/kernel_key.pem --key /percorso/della/kernel_key.pem --output /usr/src/linux-x.y.z/percorso/della/kernel-image`

**Nota**

In questo esempio, la stessa chiave che è stata generata per firmare i moduli è utilzzata per firmare l'immagine del kernel. E' anche possibile generare e usare una seconda chiave separata per firmare l'immagine del kernel. Può essere usato di nuovo lo stesso comando OpenSSL della sezione precedente.

Poi procedere con l'installazione.

Per firmare automaticamente gli eseguibili EFI installati da altri pacchetti, abilitare l'opzione USE [secureboot](https://packages.gentoo.org/useflags/secureboot) [USE flag/it (page does not exist)](/index.php?title=USE_flag/it&action=edit&redlink=1 "USE flag/it (page does not exist)") globalmente:

FILE **`/etc/portage/make.conf`** **Abilitare Secure Boot**

```
USE="modules-sign secureboot"

# Facoltativamente, per usare chiavi di firma personali.
MODULES_SIGN_KEY="/percorso/della/chiave_del_kernel.pem"
MODULES_SIGN_CERT="/percorso/della/chiave_del_kernel.pem" # Richiesta solo se MODULES_SIGN_KEY non contiene anche il certificato.
MODULES_SIGN_HASH="sha512" # Il valore predefinito è sha512.

# Facoltativamente, per avviare con secureboot abilitato; può essere la stessa o una diversa chiave di firma.
SECUREBOOT_SIGN_KEY="/percorso/della/chiave_del_kernel.pem"
SECUREBOOT_SIGN_CERT="/percorso/della/chiave_del_kernel.pem"

```

**Nota**

SECUREBOOT\_SIGN\_KEY e SECUREBOOT\_SIGN\_CERTpossono puntare a file diversi. In questo esempio, il file pem generato da OpenSSL include sia la chiave che il relativo certificato, e perciò entrambe le variabili sono impostate allo stesso valore.

**Nota**

Quando si genera una [Unified Kernel Image](/wiki/Unified_Kernel_Image "Unified Kernel Image") con `ukify` di systemd l'immagine del kernel verrà firmata automaticamente prima dell'inclusione nell'immagine unificata del kernel e non sarà necessario firmarla manualmente.

#### Compilare e installare

Ora che la configurazione è completata, è il momento di compilare e installare il kernel. Uscire dalla configurazione ed avviare il processo di compilazione:

`root #` `make -j$(nproc) && make modules_install`

**Nota**

È possibile ridurre la compilazione parallela usando make -jX, dove `X` è un numero intero delle attività parallele che il processo di compilazione è abilitato a far girare. Le istruzioni sono simili a quelle usate precedentemente per /etc/portage/make.conf, con la variabile MAKEOPTS.

Quando il kernel ha terminato la compilazione, copiare l'immagine del kernel in /boot/. L'operazione viene gestita dal comando make install:

`root #` `make install`

Questo comando copierà l'immagine del kernel in /boot/. Se è installato [sys-kernel/installkernel](https://packages.gentoo.org/packages/sys-kernel/installkernel) invece chiamerà /sbin/installkernel a cui delegherà l'installazione del kernel. Invece di copiare semplicemente il kernel in /boot/, [Installkernel](/wiki/Installkernel "Installkernel") installa ogni kernel col proprio numero di versione nel nome del file. Inoltre, installkernel fornisce una infrastruttura per compiere automaticamente diverse operazioni inerenti l'installazione del kernel, come: generare un [initramfs](/wiki/Initramfs "Initramfs"), costruire un [UKS(Unified Kernel Image)](/wiki/Unified_Kernel_Image "Unified Kernel Image") e aggiornare la configurazione del [programma di avvio](/wiki/Bootloader "Bootloader") .

Continuare l'installazione con [configurare il sistema](/wiki/Handbook:AMD64/Installation/System/it "Handbook:AMD64/Installation/System/it").

[← Installare il sistema base](/wiki/Handbook:AMD64/Installation/Base/it "Previous part") [Home](/wiki/Handbook:AMD64 "Handbook:AMD64") [Configurare il sistema →](/wiki/Handbook:AMD64/Installation/System/it "Next part")