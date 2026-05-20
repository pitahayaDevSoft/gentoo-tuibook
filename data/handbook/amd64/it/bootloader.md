# Bootloader

Other languages:

- [Deutsch](/wiki/Handbook:AMD64/Installation/Bootloader/de "Handbook:AMD64/Installation/Bootloader/de (100% translated)")
- [English](/wiki/Handbook:AMD64/Installation/Bootloader "Handbook:AMD64/Installation/Bootloader (100% translated)")
- [español](/wiki/Handbook:AMD64/Installation/Bootloader/es "Handbook:AMD64/Installation/Bootloader/es (100% translated)")
- [français](/wiki/Handbook:AMD64/Installation/Bootloader/fr "Handbook:AMD64/Installation/Bootloader/fr (100% translated)")
- italiano
- [magyar](/wiki/Handbook:AMD64/Installation/Bootloader/hu "Handbook:AMD64/Installation/Bootloader/hu (100% translated)")
- [polski](/wiki/Handbook:AMD64/Installation/Bootloader/pl "Handbook:AMD64/Installation/Bootloader/pl (100% translated)")
- [português do Brasil](/wiki/Handbook:AMD64/Installation/Bootloader/pt-br "Handbook:AMD64/Installation/Bootloader/pt-br (100% translated)")
- [čeština](/wiki/Handbook:AMD64/Installation/Bootloader/cs "Handbook:AMD64/Installation/Bootloader/cs (100% translated)")
- [русский](/wiki/Handbook:AMD64/Installation/Bootloader/ru "Handbook:AMD64/Installation/Bootloader/ru (100% translated)")
- [தமிழ்](/wiki/Handbook:AMD64/Installation/Bootloader/ta "Handbook:AMD64/Installation/Bootloader/ta (100% translated)")
- [中文（中国大陆）‎](/wiki/Handbook:AMD64/Installation/Bootloader/zh-cn "Handbook:AMD64/Installation/Bootloader/zh-cn (100% translated)")
- [日本語](/wiki/Handbook:AMD64/Installation/Bootloader/ja "Handbook:AMD64/Installation/Bootloader/ja (100% translated)")
- [한국어](/wiki/Handbook:AMD64/Installation/Bootloader/ko "Handbook:AMD64/Installation/Bootloader/ko (100% translated)")

[Manuale AMD64](/wiki/Handbook:AMD64/it "Handbook:AMD64/it")[Installazione](/wiki/Handbook:AMD64/Full/Installation/it "Handbook:AMD64/Full/Installation/it")[Riguardo l'installazione](/wiki/Handbook:AMD64/Installation/About/it "Handbook:AMD64/Installation/About/it")[Scegliere il supporto](/wiki/Handbook:AMD64/Installation/Media/it "Handbook:AMD64/Installation/Media/it")[Configurare la rete](/wiki/Handbook:AMD64/Installation/Networking/it "Handbook:AMD64/Installation/Networking/it")[Preparare i dischi](/wiki/Handbook:AMD64/Installation/Disks/it "Handbook:AMD64/Installation/Disks/it")[Il file stage](/wiki/Handbook:AMD64/Installation/Stage/it "Handbook:AMD64/Installation/Stage/it")[Installare il sistema base](/wiki/Handbook:AMD64/Installation/Base/it "Handbook:AMD64/Installation/Base/it")[Configurare il kernel](/wiki/Handbook:AMD64/Installation/Kernel/it "Handbook:AMD64/Installation/Kernel/it")[Configurare il sistema](/wiki/Handbook:AMD64/Installation/System/it "Handbook:AMD64/Installation/System/it")[Installare gli strumenti](/wiki/Handbook:AMD64/Installation/Tools/it "Handbook:AMD64/Installation/Tools/it")Impostare programma d'avvio[Concludere](/wiki/Handbook:AMD64/Installation/Finalizing/it "Handbook:AMD64/Installation/Finalizing/it")[Lavorare con Gentoo](/wiki/Handbook:AMD64/Full/Working/it "Handbook:AMD64/Full/Working/it")[Introduzione a Portage](/wiki/Handbook:AMD64/Working/Portage/it "Handbook:AMD64/Working/Portage/it")[Opzioni USE](/wiki/Handbook:AMD64/Working/USE/it "Handbook:AMD64/Working/USE/it")[Funzionalità di Portage](/wiki/Handbook:AMD64/Working/Features/it "Handbook:AMD64/Working/Features/it")[Sistema script di init](/wiki/Handbook:AMD64/Working/Initscripts/it "Handbook:AMD64/Working/Initscripts/it")[Variabili d'ambiente](/wiki/Handbook:AMD64/Working/EnvVar/it "Handbook:AMD64/Working/EnvVar/it")[Lavorare con Portage](/wiki/Handbook:AMD64/Full/Portage/it "Handbook:AMD64/Full/Portage/it")[File e cartelle](/wiki/Handbook:AMD64/Portage/Files/it "Handbook:AMD64/Portage/Files/it")[Variabili](/wiki/Handbook:AMD64/Portage/Variables/it "Handbook:AMD64/Portage/Variables/it")[Miscelare i rami del software](/wiki/Handbook:AMD64/Portage/Branches/it "Handbook:AMD64/Portage/Branches/it")[Strumenti aggiuntivi](/wiki/Handbook:AMD64/Portage/Tools/it "Handbook:AMD64/Portage/Tools/it")[Catalogo pacchetti personalizzato](/wiki/Handbook:AMD64/Portage/CustomTree/it "Handbook:AMD64/Portage/CustomTree/it")[Funzionalità avanzate](/wiki/Handbook:AMD64/Portage/Advanced/it "Handbook:AMD64/Portage/Advanced/it")[Configurare la rete](/wiki/Handbook:AMD64/Full/Networking/it "Handbook:AMD64/Full/Networking/it")[Come iniziare](/wiki/Handbook:AMD64/Networking/Introduction/it "Handbook:AMD64/Networking/Introduction/it")[Configurazione avanzata](/wiki/Handbook:AMD64/Networking/Advanced/it "Handbook:AMD64/Networking/Advanced/it")[Networking modulare](/wiki/Handbook:AMD64/Networking/Modular/it "Handbook:AMD64/Networking/Modular/it")[Wireless](/wiki/Handbook:AMD64/Networking/Wireless/it "Handbook:AMD64/Networking/Wireless/it")[Aggiungere funzionalità](/wiki/Handbook:AMD64/Networking/Extending/it "Handbook:AMD64/Networking/Extending/it")[Gestione dinamica](/wiki/Handbook:AMD64/Networking/Dynamic/it "Handbook:AMD64/Networking/Dynamic/it")

## Contents

- [1Scegliere un avviatore](#Scegliere_un_avviatore)
- [2Predefinito: GRUB2](#Predefinito:_GRUB2)
  - [2.1Compilare (emergere)](#Compilare_.28emergere.29)
  - [2.2Installare](#Installare)
    - [2.2.1Optional: Secure Boot](#Optional:_Secure_Boot)
    - [2.2.2Debugging GRUB](#Debugging_GRUB)
  - [2.3Configurare](#Configurare)
- [3Alternative 1: systemd-boot](#Alternative_1:_systemd-boot)
  - [3.1Emerge](#Emerge)
  - [3.2Installation](#Installation)
  - [3.3Optional: Secure Boot](#Optional:_Secure_Boot_2)
- [4Alternativa 2: efibootmgr](#Alternativa_2:_efibootmgr)
  - [4.1Unified Kernel Image](#Unified_Kernel_Image)
- [5Other Alternatives](#Other_Alternatives)
- [6Riavviare il sistema](#Riavviare_il_sistema)

## Scegliere un avviatore

Con il kernel Linux configurato, gli strumenti di sistema installati ed i file di configurazione modificati, è tempo di installare l'ultimo importante pezzo di un'installazione Linux: l'avviatore (bootloader).

L'avviatore è responsabile del lancio del kernel Linux al momento dell'avvio - senza di esso, il sistema non saprebbe come procedere quando il pulsante di accensione (power) viene premuto.

Per **amd64**, documenteremo come configurare sia [GRUB2](#Predefinito:_Usare_GRUB2) che [LILO](#Alternativa:_Usare_LILO) per sistemi basati su BIOS, e [GRUB2](#Predefinito:_Usare_GRUB2) o [efibootmgr](#Alternativa:_Usare_efibootmgr) per sistemi UEFI.

In questa sezione del Manuale, viene fatta una delineazione tra _emergere_ un pacchetto per l'avviatore (boot loader) ed _installare_ un avviatore nel disco di sistema. Qui il termine _emergere_ sarà usato per chiedere a [Portage](/wiki/Portage "Portage") di compilare un pacchetto software che sia disponibile per il sistema. Il termine _installare_ significherà che l'avviatore copia i file o modifica fisicamente le sezioni appropriate dell'unità del disco di sistema affinché l'avviatore sia _attivo e pronto ad operare_ dalla prossima accensione della macchina.

## Predefinito: GRUB2

In via predefinita, la maggior parte dei sistemi Gentoo si affidano ora a GRUB2 (presente nel pacchetto [sys-boot/grub](https://packages.gentoo.org/packages/sys-boot/grub)), che è il diretto successo di GRUB Legacy. Senza configurazioni aggiuntive, GRUB2 supporta di buon grado i vecchi sistemi BIOS ("pc"). Con alcune configurazioni, necessarie prima della fase di compilazione, GRUB2 può supportare più di una mezza dozzina di piattaforme aggiuntive. Per maggiori informazioni, consultare la [Sezione dei prerequisiti](/wiki/GRUB2#Prerequisites "GRUB2") nell'articolo di [GRUB2](/wiki/GRUB2 "GRUB2").

### Compilare (emergere)

Quando si usano sistemi BIOS più vecchi che supportano solo tabelle delle partizioni MBR, nessuna configurazione aggiuntiva è necessaria per compilare (far emergere) GRUB:

`root #` `emerge --ask --verbose sys-boot/grub:2`

Una nota per gli utenti UEFI: eseguire il comando sopra comporterà l'abilitazione dei valori GRUB\_PLATFORMS prima del comando emergere. Quando si usano sistemi che supportano UEFI, gli utenti si dovranno assicurare che `GRUB_PLATFORMS="efi-64"` sia abilitato (così come avviene nel caso predefinito). Se non è questo il caso durante la configurazione, `GRUB_PLATFORMS="efi-64"` dovrà essere aggiunto al file /etc/portage/make.conf _prima_ di far emergere GRUB2, così che il pacchetto sarà compilato con la funzionalità EFI:

`root #` `echo 'GRUB_PLATFORMS="efi-64"' >> /etc/portage/make.conf`

`root #` `emerge --ask sys-boot/grub:2`

Se GRUB2 è stato compilato (fatto emergere) senza abilitare `GRUB_PLATFORMS="efi-64"`, la linea (appena mostrata) può essere in seguito aggiunta a make.conf e le dipendenze per [il "mondo" dei pacchetti](/wiki/World_set_(Portage) "World set (Portage)") ricalcolati passando le opzioni `--update --newuse` a emerge:

`root #` `emerge --ask --update --newuse --verbose sys-boot/grub:2`

Il software GRUB2 è ora fruibile nel sistema, ma non è ancora installato.

### Installare

Poi, installare i file necessari di GRUB2 sulla cartella /boot/grub/ attraverso il comando grub-install. Supponendo che il primo disco (quello da cui il sistema si avvia) sia /dev/sda, uno dei successivi comandi sarà:

- Quando si usa il BIOS:

`root #` `grub-install /dev/sda`

For DOS/Legacy BIOS systems:

`root #` `grub-install /dev/sda`

- Quando si usa UEFI:

**Importante**

Assicurarsi che la partizione di sistema EFI sia stata montata, _prima_ di eseguire grub-install. È possibile che grub-install installi il file EFI GRUB (grubx64.efi) nella cartella sbagliata **senza** fornire _alcuna_ indicazione sulla cartella sbagliata usata.

For UEFI systems:

`root #` `grub-install --target=x86_64-efi --efi-directory=/boot`

Upon successful installation, the output should match the output of the previous command. If the output does not match exactly, then proceed to [Debugging GRUB](/wiki/Handbook:AMD64/Blocks/Bootloader/it#Debugging_GRUB "Handbook:AMD64/Blocks/Bootloader/it"), otherwise jump to the [Configure step](/wiki/Handbook:AMD64/Blocks/Bootloader/it#GRUB_Configure "Handbook:AMD64/Blocks/Bootloader/it").

##### Optional: Secure Boot

To successfully boot with secure boot enabled the signing certificate must either be accepted by the [UEFI](/wiki/UEFI "UEFI") firmware, or [shim](/wiki/Shim "Shim") must be used as a pre-loader. Shim is pre-signed with the third-party Microsoft Certificate, accepted by default by most UEFI motherboards.

How to configure the UEFI firmware to accept custom keys depends on the firmware vendor, which is beyond the scope of the handbook. Below is shown how to setup shim instead. Here it is assumed that the user has already followed the instructions in the previous sections to generate a signing key and to configure portage to use it. If this is not the case please return first to the [Kernel installation](/wiki/Handbook:AMD64/Installation/Kernel/it "Handbook:AMD64/Installation/Kernel/it") section.

The package [sys-boot/grub](https://packages.gentoo.org/packages/sys-boot/grub) installs a prebuilt and signed stand-alone EFI executable if the [secureboot](https://packages.gentoo.org/useflags/secureboot) [USE flag/it (page does not exist)](/index.php?title=USE_flag/it&action=edit&redlink=1 "USE flag/it (page does not exist)") USE flag is enabled. Install the required packages and copy the stand-alone grub, Shim, and the MokManager to the same directory on the EFI System Partition. For example:

`root #` `emerge sys-boot/grub sys-boot/shim sys-boot/mokutil sys-boot/efibootmgr`

`root #` `cp /usr/share/shim/BOOTX64.EFI /efi/EFI/Gentoo/shimx64.efi
`

`root #` `cp /usr/share/shim/mmx64.efi /efi/EFI/Gentoo/mmx64.efi
`

`root #` `cp /usr/lib/grub/grub-x86_64.efi.signed /efi/EFI/Gentoo/grubx64.efi
`

Next register the signing key in shims MOKlist, this requires keys in the DER format, whereas sbsign and the kernel build system expect keys in the PEM format. In the previous sections of the handbook an example was shown to generate such a signing PEM key, this key must now be converted to the DER format:

`root #` `openssl x509 -in /path/to/kernel_key.pem -inform PEM -out /path/to/kernel_key.der -outform DER`

**Nota**

The path used here must be the path to the pem file containing the certificate belonging to the generated key. In this example both key and certificate are in the same pem file.

Then the converted certificate can be imported into Shims MOKlist, this command will ask to set some password for the import request:

`root #` `mokutil --import /path/to/kernel_key.der`

**Nota**

When the currently booted kernel already trusts the certificate being imported, the message "Already in kernel trusted keyring." will be returned here. If this happens, re-run the above command with the argument --ignore-keyring added.

Next, register Shim with the UEFI firmware. In the following command, `boot-disk` and `boot-partition-id` must be replaced with the disk and partition identifier of the EFI system partition:

`root #` `efibootmgr --create --disk /dev/boot-disk --part boot-partition-id --loader '\EFI\Gentoo\shimx64.efi' --label 'GRUB via Shim' --unicode`

Note that this prebuilt and signed stand-alone version of grub reads the grub.cfg from a different location then usual. Instead of the default /boot/grub/grub.cfg the config file should be in the same directory that the grub EFI executable is in, e.g. /efi/EFI/Gentoo/grub.cfg. When [sys-kernel/installkernel](https://packages.gentoo.org/packages/sys-kernel/installkernel) is used to install the kernel and update the grub configuration then the GRUB\_CFG environment variable may be used to override the usual location of the grub config file.

For example:

`root #` `grub-mkconfig -o /efi/EFI/Gentoo/grub.cfg`

Or, via [installkernel](/wiki/Installkernel "Installkernel"):

FILE **`/etc/env.d/99grub`**

```
GRUB_CFG=/efi/EFI/Gentoo/grub.cfg

```

`root #` `env-update`

**Nota**

The import process will not be completed until the system is rebooted. After completing all steps in the handbook, restart the system and Shim will load, it will find the import request registered by mokutil. The MokManager application will start and ask for the password that was set when creating the import request. Follow the on-screen instructions to complete the import of the certificate, then reboot the system into the UEFI menu and enable the Secure Boot setting.

##### Debugging GRUB

When debugging GRUB, there are a couple of quick fixes that may result in a bootable installation without having to reboot to a new live image environment.

In the event that "EFI variables are not supported on this system" is displayed somewhere in the output, it is likely the live image was not booted in EFI mode and is presently in Legacy BIOS boot mode. The solution is to try the [removable GRUB step](/wiki/Handbook:AMD64/Blocks/Bootloader/it#GRUB_Install_EFI_systems_removable "Handbook:AMD64/Blocks/Bootloader/it") mentioned below. This will overwrite the executable EFI file located at /EFI/BOOT/BOOTX64.EFI. Upon rebooting in EFI mode, the motherboard firmware may execute this default boot entry and execute GRUB.

If grub-install returns an error that says "Could not prepare Boot variable: Read-only file system", and the live environment was correctly booted in UEFI mode, then it should be possible to remount the efivars special mount as read-write and then re-run the [aforementioned grub-install command](/wiki/Handbook:AMD64/Blocks/Bootloader/it#GRUB_Install_EFI_systems_command "Handbook:AMD64/Blocks/Bootloader/it"):

`root #` `mount -o remount,rw,nosuid,nodev,noexec --types efivarfs efivarfs /sys/firmware/efi/efivars`

This is caused by certain non-official Gentoo environments not mounting the special EFI filesystem by default. If the previous command does not run, then reboot using an official Gentoo live image environment in EFI mode.

Alcuni produttori di schede madri sembrano supportare soltanto la posizione della cartella /efi/boot/ per i file .EFI nella partizione di sistema EFI (ESP). L'installatore di GRUB può eseguire questa operazione automaticamente con l'opzione `--removable`. Verificare che ESP sia montata prima di eseguire i seguenti comandi. Supponendo che ESP sia montata su /boot (come suggerito prima), eseguire:

`root #` `grub-install --target=x86_64-efi --efi-directory=/boot --removable`

Questo crea una cartella predefinita definita dalle specifiche UEFI, e poi copia i file grubx64.efi nella posizione 'predefinita'dei file EFI definita dalle stesse specifiche.

### Configurare

Poi, generare la configurazione GRUB2 basata su quella specificata dall'utente nel file /etc/default/grub e negli script /etc/grub.d. Nella maggior parte dei casi, nessuna configurazione è richiesta agli utenti, in quanto GRUB2 rileverà automaticamente quale kernel avviare (quello più in alto disponibile su /boot/) e qual è il file system radice. È anche possibile aggiungere parametri per il kernel su /etc/default/grub tramite la variabile GRUB\_CMDLINE\_LINUX.

Per generare la configurazione finale di GRUB2, eseguire il comando grub-mkconfig:

`root #` `grub-mkconfig -o /boot/grub/grub.cfg`

```
Generating (Sto generando) grub.cfg ...
Found linux image (Immagine linux trovata): /boot/vmlinuz-6.18.8-gentoo
Found initrd image (Immagine initrd trovata): /boot/initramfs-genkernel-amd64-6.18.8-gentoo
done (fatto)

```

In risposta il comando deve riferire che almeno un'immagine Linux è stata trovata, in quanto queste sono necessarie per l'avvio del sistema. Se viene utilizzato un initramfs o genkernel è stato usato per compilare il kernel, anche l'immagine initrd corretta dovrebbe essere rilevata. Se questo non succede, andare su /boot/ e controllare il contenuto con il comando ls. Se effettivamente mancano i file, ritornare alla configurazione del kernel ed alle istruzioni di installazione.

**Suggerimento**

L'utilità os-prober può essere usata in combinazione con GRUB2 per rilevare altri sistemi operativi dai dischi connessi. Windows 7, 8.1, 10 ed altre distribuzioni Linux sono rilevabili. Chi desidera un sistema con doppio avvio dovrebbe compilare il pacchetto [sys-boot/os-prober](https://packages.gentoo.org/packages/sys-boot/os-prober), poi rieseguire il comando grub-mkconfig (come visto sopra). Se si incontrano problemi di rilevamento assicurasi di leggere l'articolo [GRUB2](/wiki/GRUB2 "GRUB2") per intero _prima_ di chiedere supporto alla comunità di Gentoo.

## Alternative 1: systemd-boot

Another option is [systemd-boot](/wiki/Systemd/systemd-boot "Systemd/systemd-boot"), which works on both OpenRC and systemd machines. It is a thin chainloader and works well with secure boot.

### Emerge

To install systemd-boot, enable the [boot](https://packages.gentoo.org/useflags/boot) [USE flag/it (page does not exist)](/index.php?title=USE_flag/it&action=edit&redlink=1 "USE flag/it (page does not exist)") USE flag and re-install [sys-apps/systemd](https://packages.gentoo.org/packages/sys-apps/systemd) (for systemd systems) or [sys-apps/systemd-utils](https://packages.gentoo.org/packages/sys-apps/systemd-utils) (for OpenRC systems):

FILE **`/etc/portage/package.use/systemd-boot`**

```
sys-apps/systemd boot
sys-apps/systemd-utils boot

```

`root #` `emerge --ask sys-apps/systemd`

Or

`root #` `emerge --ask sys-apps/systemd-utils`

### Installation

Now, install the systemd-boot loader to the [EFI System Partition](/wiki/EFI_System_Partition "EFI System Partition"):

`root #` `bootctl install`

**Importante**

Make sure the EFI system partition has been mounted _before_ running bootctl install.

When using this bootloader, before rebooting, verify that a new bootable entry exists using:

`root #` `bootctl list`

**Attenzione**

The kernel command line for new systemd-boot entries is read from /etc/kernel/cmdline or /usr/lib/kernel/cmdline. If neither file is present, then the kernel command line of the currently booted kernel is re-used (/proc/cmdline). On new installs it might therefore happen that the kernel command line of the live CD is accidentally used to boot the new kernel. The kernel command line for registered entries can be checked with:

`root #` `bootctl list`

If this does not show the desired kernel command line then create /etc/kernel/cmdline containing the correct kernel command line and re-install the kernel.

If no new entry exists, ensure the [sys-kernel/installkernel](https://packages.gentoo.org/packages/sys-kernel/installkernel) package has been installed with the [systemd](https://packages.gentoo.org/useflags/systemd) [USE flag/it (page does not exist)](/index.php?title=USE_flag/it&action=edit&redlink=1 "USE flag/it (page does not exist)") and [systemd-boot](https://packages.gentoo.org/useflags/systemd-boot) [USE flag/it (page does not exist)](/index.php?title=USE_flag/it&action=edit&redlink=1 "USE flag/it (page does not exist)") USE flags enabled, and re-run the kernel installation.

For the distribution kernels:

`root #` `emerge --ask --config sys-kernel/gentoo-kernel`

For a manually configured and compiled kernel:

`root #` `make install`

**Importante**

When installing kernels for systemd-boot, no root= kernel command line argument is added by default. On systemd systems that are using an initramfs users may rely instead on [systemd-gpt-auto-generator](/wiki/Systemd#Automatic_mounting_of_partitions_at_boot.2Fit "Systemd") to automatically find the root partition at boot. Otherwise users should manually specify the location of the root partition by setting root= in /etc/kernel/cmdline as well as any other kernel command line arguments that should be used. And then reinstalling the kernel as described above.

### Optional: Secure Boot

When the [secureboot](https://packages.gentoo.org/useflags/secureboot) [USE flag/it (page does not exist)](/index.php?title=USE_flag/it&action=edit&redlink=1 "USE flag/it (page does not exist)") USE flag is enabled, the systemd-boot EFI executable will be signed by portage automatically. Furthermore, bootctl install will automatically install the signed version.

To successfully boot with secure boot enabled the used certificate must either be accepted by the [UEFI](/wiki/UEFI "UEFI") firmware, or [shim](/wiki/Shim "Shim") must be used as a pre-loader. Shim is pre-signed with the third-party Microsoft Certificate, accepted by default by most UEFI motherboards.

How to configure the UEFI firmware to accept custom keys depends on the firmware vendor, which is beyond the scope of the handbook. Below is shown how to setup shim instead. Here it is assumed that the user has already followed the instructions in the previous sections to generate a signing key and to configure portage to use it. If this is not the case please return first to the [Kernel installation](/wiki/Handbook:AMD64/Installation/Kernel/it "Handbook:AMD64/Installation/Kernel/it") section.

`root #` `emerge --ask sys-boot/shim sys-boot/mokutil sys-boot/efibootmgr`

`root #` `bootctl install --no-variables`

`root #` `cp /usr/share/shim/BOOTX64.EFI /efi/EFI/systemd/shimx64.efi`

`root #` `cp /usr/share/shim/mmx64.efi /efi/EFI/systemd/mmx64.efi`

Shims MOKlist requires keys in the DER format, whereas sbsign and the kernel build system expect keys in the PEM format. In the previous sections of the handbook an example was shown to generate such a signing PEM key, this key must now be converted to the DER format:

`root #` `openssl x509 -in /path/to/kernel_key.pem -inform PEM -out /path/to/kernel_key.der -outform DER`

**Nota**

The path used here must be the path to the pem file containing the certificate belonging to the generated key. In this example both key and certificate are in the same pem file.

Then the converted certificate can be imported into Shims MOKlist:

`root #` `mokutil --import /path/to/kernel_key.der`

**Nota**

When the currently booted kernel already trusts the certificate being imported, the message "Already in kernel trusted keyring." will be returned here. If this happens, re-run the above command with the argument --ignore-keyring added.

And finally we register Shim with the UEFI firmware. In the following command, `boot-disk` and `boot-partition-id` must be replaced with the disk and partition identifier of the EFI system partition:

`root #` `efibootmgr --create --disk /dev/boot-disk --part boot-partition-id --loader '\EFI\systemd\shimx64.efi'  --label 'Systemd-boot via Shim' --unicode '\EFI\systemd\systemd-bootx64.efi'`

**Nota**

The import process will not be completed until the system is rebooted. After completing all steps in the handbook, restart the system and Shim will load, it will find the import request registered by mokutil. The MokManager application will start and ask for the password that was set when creating the import request. Follow the on-screen instructions to complete the import of the certificate, then reboot the system into the UEFI menu and enable the Secure Boot setting.

## Alternativa 2: efibootmgr

Su sistemi UEFI, il firmware UEFI del sistema (in altre parole il principale bootloader - avviatore) può essere manipolato direttamente per cercare le voci di avvio UEFI. Questi sistemi non necessitano di avviatori aggiuntivi (anche conosciuti come avviatori secondari), come GRUB2 allo scopo di aiutare il sistema ad avviarsi. Dopo quanto detto, la ragione degli avviatori basati su EFI come GRUB2 esiste per _estendere_ le funzionalità dei sistemi UEFI durante il processo di avvio. Usare efibootmgr è per coloro che desiderano decisamente mantenere un approccio minimalista (sebbene più rigido) nell'avvio del proprio sistema; usare GRUB2 (vedi sopra) è più facile per la maggior parte degli utenti perché offre un approccio più flessibile quando si avviano i sistemi UEFI.

System administrators who desire to take a minimalist, although more rigid, approach to booting the system can avoid secondary bootloaders and boot the Linux kernel as an [EFI stub](/wiki/EFI_stub "EFI stub").

Terner presente che l'applicazione [sys-boot/efibootmgr](https://packages.gentoo.org/packages/sys-boot/efibootmgr) non è un avviatore; piuttosto è uno strumento per interagire con il firmware UEFI ed aggiornare le sue impostazioni, così che il kernel Linux precedentemente installato possa essere avviato con opzioni aggiuntive (se necessario), oppure per abilitare più voci di avvio. Questa interazione è possibile grazie alle variabili EFI (da cui la necessità del supporto delle variabili EFI da parte del kernel).

Assicurarsi di leggere l'articolo [EFI stub kernel](/wiki/EFI_stub_kernel "EFI stub kernel") _prima_ di continuare. Il kernel deve avere opzioni specifiche abilitate affinché sia direttamente avviabile dal firmware UEFI del sistema. Potrebbe risultare necessario ricompilare il kernel. È anche una buona idea dare un'occhiata all'articolo [efibootmgr](/wiki/Efibootmgr "Efibootmgr").

It is also a good idea to take a look at the [efibootmgr](/wiki/Efibootmgr "Efibootmgr") article for additional information.

**Nota**

Per ribadire, efibootmgr _non_ è un requisito di avvio di un sistema UEFI. Il kernel Linux stesso può essere avviato immediatamente, ed opzioni aggiuntive per il kernel tramite linea di comando possono essere integrate nel kernel Linux (c'è un'opzione di configurazione del kernel chiamata che permette agli utenti di specificare i parametri di avvio tramite opzioni a linea di comando). Anche un initramfs può essere 'integrato' nel kernel.

Coloro che hanno deciso di seguire questo approccio devono installare il software:

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

Then reinstall [installkernel](/wiki/Installkernel "Installkernel"), create the /efi/EFI/Gentoo directory and reinstall the kernel:

`root #` `emerge --ask sys-kernel/installkernel`

`root #` `mkdir -p /efi/EFI/Gentoo`

For distribution kernels:

`root #` `emerge --ask --config sys-kernel/gentoo-kernel{,-bin}`

For manually managed kernels:

`root #` `make install`

Poi, creare la cartella /boot/efi/boot/ e copiare qui il kernel, chiamandolo bzImage.efi:

`root #` `mkdir -p /boot/efi/boot
`

`root #` `cp /boot/vmlinuz-* /boot/efi/boot/bzImage.efi
`

Install the efibootmgr package:

`root #` `emerge --ask sys-boot/efibootmgr`

In seguito, comunicare al firmware UEFI che una voce di avvio nominata "Gentoo" deve essere creata, la quale contiene il kernel appena compilato con la parte EFI:

`root #` `efibootmgr --create --disk /dev/sda --part 2 --label "Gentoo" --loader "\efi\boot\bzImage.efi"`

**Nota**

L'uso di `\` come separatore di cartelle è obbligatorio quando si usano le definizioni UEFI.

Se un file system iniziale su RAM (initramfs) viene usato, aggiungere l'appropriata opzione di avvio ad esso:

`root #` `efibootmgr -c -d /dev/sda -p 2 -L "Gentoo" -l "\efi\boot\bzImage.efi" initrd='\initramfs-genkernel-amd64-6.18.8-gentoo'`

**Suggerimento**

Additional kernel command line options may be parsed by the firmware to the kernel by specifying them along with the initrd=... option as shown above.

Concluse queste modifiche, quando il sistema si riavvia, sarà disponibile una voce di avvio chiamata "Gentoo".

### Unified Kernel Image

If [installkernel](/wiki/Installkernel "Installkernel") was configured to build and install unified kernel images. The unified kernel image should already be installed to the EFI/Linux directory on the EFI system partition, if this is not the case ensure the directory exists and then run the kernel installation again as described earlier in the handbook.

To add a direct boot entry for the installed unified kernel image:

`root #` `efibootmgr --create --disk /dev/sda --part 1 --label "gentoo" --loader "\EFI\Linux\gentoo-x.y.z.efi"`

## Other Alternatives

For other options that are not covered in the Handbook, see the full list of available [bootloaders](/wiki/Bootloader "Bootloader").

## Riavviare il sistema

Uscire dall'ambiente con la radice cambiata (chroot) e smontare tutte le partizioni montate. Poi si inserisca quell'unico magico comando che avvia il vero test finale: reboot.

`(chroot) livecd #` `exit`

`livecd~#` `cd
`

`livecd~#` `umount -l /mnt/gentoo/dev{/shm,/pts,}
`

`livecd~#` `umount -R /mnt/gentoo
`

`livecd~#` `reboot`

Non dimenticare di rimuovere l'immagine live, altrimenti potrebbe essere riavviata al posto del nuovo sistema Gentoo installato.

Una volta effettuato il riavvio nell'ambiente Gentoo appena installato, è ragionevole concludere con [Ultimare l'installazione di Gentoo](/wiki/Handbook:AMD64/Installation/Finalizing/it "Handbook:AMD64/Installation/Finalizing/it").

[← Installare gli strumenti](/wiki/Handbook:AMD64/Installation/Tools/it "Previous part") [Home](/wiki/Handbook:AMD64/it "Handbook:AMD64/it") [Concludere →](/wiki/Handbook:AMD64/Installation/Finalizing/it "Next part")