# Bootloader

Other languages:

- [Deutsch](/wiki/Handbook:AMD64/Installation/Bootloader/de "Handbook:AMD64/Installation/Bootloader/de (100% translated)")
- [English](/wiki/Handbook:AMD64/Installation/Bootloader "Handbook:AMD64/Installation/Bootloader (100% translated)")
- [español](/wiki/Handbook:AMD64/Installation/Bootloader/es "Handbook:AMD64/Installation/Bootloader/es (100% translated)")
- [français](/wiki/Handbook:AMD64/Installation/Bootloader/fr "Handbook:AMD64/Installation/Bootloader/fr (100% translated)")
- [italiano](/wiki/Handbook:AMD64/Installation/Bootloader/it "Handbook:AMD64/Installation/Bootloader/it (100% translated)")
- [magyar](/wiki/Handbook:AMD64/Installation/Bootloader/hu "Handbook:AMD64/Installation/Bootloader/hu (100% translated)")
- [polski](/wiki/Handbook:AMD64/Installation/Bootloader/pl "Handbook:AMD64/Installation/Bootloader/pl (100% translated)")
- [português do Brasil](/wiki/Handbook:AMD64/Installation/Bootloader/pt-br "Handbook:AMD64/Installation/Bootloader/pt-br (100% translated)")
- čeština
- [русский](/wiki/Handbook:AMD64/Installation/Bootloader/ru "Handbook:AMD64/Installation/Bootloader/ru (100% translated)")
- [தமிழ்](/wiki/Handbook:AMD64/Installation/Bootloader/ta "Handbook:AMD64/Installation/Bootloader/ta (100% translated)")
- [中文（中国大陆）‎](/wiki/Handbook:AMD64/Installation/Bootloader/zh-cn "Handbook:AMD64/Installation/Bootloader/zh-cn (100% translated)")
- [日本語](/wiki/Handbook:AMD64/Installation/Bootloader/ja "Handbook:AMD64/Installation/Bootloader/ja (100% translated)")
- [한국어](/wiki/Handbook:AMD64/Installation/Bootloader/ko "Handbook:AMD64/Installation/Bootloader/ko (100% translated)")

[AMD64 Handbook](/wiki/Handbook:AMD64/cs "Handbook:AMD64/cs")[Instalace](/wiki/Handbook:AMD64/Full/Installation/cs "Handbook:AMD64/Full/Installation/cs")[O instalaci](/wiki/Handbook:AMD64/Installation/About/cs "Handbook:AMD64/Installation/About/cs")[Výběr média](/wiki/Handbook:AMD64/Installation/Media/cs "Handbook:AMD64/Installation/Media/cs")[Konfigurace sítě](/wiki/Handbook:AMD64/Installation/Networking/cs "Handbook:AMD64/Installation/Networking/cs")[Příprava disků](/wiki/Handbook:AMD64/Installation/Disks/cs "Handbook:AMD64/Installation/Disks/cs")[Instalace stage3](/wiki/Handbook:AMD64/Installation/Stage/cs "Handbook:AMD64/Installation/Stage/cs")[Instalace základního systému](/wiki/Handbook:AMD64/Installation/Base/cs "Handbook:AMD64/Installation/Base/cs")[Konfigurace jádra](/wiki/Handbook:AMD64/Installation/Kernel/cs "Handbook:AMD64/Installation/Kernel/cs")[Konfigurace systému](/wiki/Handbook:AMD64/Installation/System/cs "Handbook:AMD64/Installation/System/cs")[Instalace nástrojů](/wiki/Handbook:AMD64/Installation/Tools/cs "Handbook:AMD64/Installation/Tools/cs")Konfigurace zavaděče[Dokončení](/wiki/Handbook:AMD64/Installation/Finalizing/cs "Handbook:AMD64/Installation/Finalizing/cs")[Práce s Gentoo](/wiki/Handbook:AMD64/Full/Working/cs "Handbook:AMD64/Full/Working/cs")[Úvod do Portage](/wiki/Handbook:AMD64/Working/Portage/cs "Handbook:AMD64/Working/Portage/cs")[Přepínače USE](/wiki/Handbook:AMD64/Working/USE/cs "Handbook:AMD64/Working/USE/cs")[Funkce portage](/wiki/Handbook:AMD64/Working/Features/cs "Handbook:AMD64/Working/Features/cs")[Systém initskriptů](/wiki/Handbook:AMD64/Working/Initscripts/cs "Handbook:AMD64/Working/Initscripts/cs")[Proměnné prostředí](/wiki/Handbook:AMD64/Working/EnvVar/cs "Handbook:AMD64/Working/EnvVar/cs")[Práce s Portage](/wiki/Handbook:AMD64/Full/Portage/cs "Handbook:AMD64/Full/Portage/cs")[Soubory a adresáře](/wiki/Handbook:AMD64/Portage/Files/cs "Handbook:AMD64/Portage/Files/cs")[Proměnné](/wiki/Handbook:AMD64/Portage/Variables/cs "Handbook:AMD64/Portage/Variables/cs")[Mísení softwarových větví](/wiki/Handbook:AMD64/Portage/Branches/cs "Handbook:AMD64/Portage/Branches/cs")[Doplňkové nástroje](/wiki/Handbook:AMD64/Portage/Tools/cs "Handbook:AMD64/Portage/Tools/cs")[Vlastní strom Portage](/wiki/Handbook:AMD64/Portage/CustomTree/cs "Handbook:AMD64/Portage/CustomTree/cs")[Pokročilé funkce](/wiki/Handbook:AMD64/Portage/Advanced/cs "Handbook:AMD64/Portage/Advanced/cs")[Konfigurace sítě](/wiki/Handbook:AMD64/Full/Networking/cs "Handbook:AMD64/Full/Networking/cs")[Začínáme](/wiki/Handbook:AMD64/Networking/Introduction/cs "Handbook:AMD64/Networking/Introduction/cs")[Pokročilá konfigurace](/wiki/Handbook:AMD64/Networking/Advanced/cs "Handbook:AMD64/Networking/Advanced/cs")[Modulární síťování](/wiki/Handbook:AMD64/Networking/Modular/cs "Handbook:AMD64/Networking/Modular/cs")[Bezdrátové sítě](/wiki/Handbook:AMD64/Networking/Wireless/cs "Handbook:AMD64/Networking/Wireless/cs")[Přidání funkcí](/wiki/Handbook:AMD64/Networking/Extending/cs "Handbook:AMD64/Networking/Extending/cs")[Dynamická správa](/wiki/Handbook:AMD64/Networking/Dynamic/cs "Handbook:AMD64/Networking/Dynamic/cs")

## Contents

- [1Výběr zavaděče](#V.C3.BDb.C4.9Br_zavad.C4.9B.C4.8De)
- [2Výchozí: GRUB2](#V.C3.BDchoz.C3.AD:_GRUB2)
  - [2.1Nahrátí (emerge)](#Nahr.C3.A1t.C3.AD_.28emerge.29)
  - [2.2Instalace](#Instalace)
    - [2.2.1Optional: Secure Boot](#Optional:_Secure_Boot)
    - [2.2.2Debugging GRUB](#Debugging_GRUB)
  - [2.3Nastavení](#Nastaven.C3.AD)
- [3Alternative 1: systemd-boot](#Alternative_1:_systemd-boot)
  - [3.1Emerge](#Emerge)
  - [3.2Installation](#Installation)
  - [3.3Optional: Secure Boot](#Optional:_Secure_Boot_2)
- [4Alternativa 2: efibootmgr](#Alternativa_2:_efibootmgr)
  - [4.1Unified Kernel Image](#Unified_Kernel_Image)
- [5Other Alternatives](#Other_Alternatives)
- [6Restartování systému](#Restartov.C3.A1n.C3.AD_syst.C3.A9mu)

## Výběr zavaděče

Po nakonfigurování jádra Linux, nainstalování systémových nástrojů a dokončenou úpravou konfiguračních souborů je čas k nainstalování poslední důležité součásti linuxové instalace: zavaděče systému.

Zavaděč je odpovědný za spuštění jádra Linux během nabíhání - bez něj by systém nevěděl, jak pokračovat dál po té, co počítač stiskem tlačítka spustíte.

Pro architekturu **amd64** uvádíme jak nastavit buď [GRUB2](#V.C3.BDchoz.C3.AD:_Pou.C5.BEit.C3.AD_GRUB2) nebo [LILO](#Alternativn.C3.AD:_Pou.C5.BEit.C3.AD_LILO) pro systémy založené na BIOSu a [#Výchozí: Použití GRUB2](#V.C3.BDchoz.C3.AD:_Pou.C5.BEit.C3.AD_GRUB2) nebo [efibootmgr](#Alternativn.C3.AD:_Pou.C5.BEit.C3.AD_efibootmgr) pro systémy s UEFI.

V této sekci příručky rozlišujeme mezi nahrátím (emerging) zavaděče a jeho instalací (installing) na systémový disk. Termín instalace bude použit pro použití [Portage](/wiki/Portage "Portage") k tomu, aby byl balíček zpřístupněn v systému. Pojem "nahrátí" pak bude označovat, že zavaděč kopíruje soubory a fyzicky mění odpovídající sekce na systémovém disku za účelem aktivace a přípravy zavaděče na jeho fungování v dalším spouštěcím cyklu.

## Výchozí: GRUB2

Většina systémů Gentoo se spoléhá na GRUB2 (k nalezení v balíčku [sys-boot/grub](https://packages.gentoo.org/packages/sys-boot/grub)), který je přímým nástupcem "GRUB Legacy". Bez jakékoli další konfigurace GRUB2 podporuje starší systémy ("pc") s BIOSem. S malou dávkou konfigurace nezbytné před sestavením, podporuje GURB2 půltucet dalších platforem. Pro více informací konzultujte [Oddíl o prerekvizitách](/wiki/GRUB2#Prerequisities "GRUB2") v článku [GRUB2](/wiki/GRUB2 "GRUB2").

### Nahrátí (emerge)

Při použití staršího BIOSu, který podporuje pouze tabulky oddílů MBR není před instalací třeba žádné dodatečné konfigurace:

`root #` `emerge --ask --verbose sys-boot/grub:2`

Poznámka pro uživatele UEFI: po spuštění shora uvedeného příkazu se zobrazí povolené hodnoty v GRUB\_PLATFORMS před instalací. V případě že použití sytému s UEFI by se uživatelé měli ujistit, že je nastavena proměnná `GRUB_PLATFORMS="efi-64"` (což je výchozí stav). Pokud tomu tak v daném případě není, je třeba doplnit `GRUB_PLATFORMS="efi-64"` do souboru /etc/portage/make.conf před instalací GRUB2, tak aby byl balíček sestaven s podporou funkcí EFI.

`root #` `echo 'GRUB_PLATFORMS="efi-64"' >> /etc/portage/make.conf`

`root #` `emerge --ask sys-boot/grub:2`

Pokud došlo k instalaci GRUB2 bez toho, aby byla napřed nastavena proměnná `GRUB_PLATFORMS="efi-64"`, může to být doplněno do make.conf později a závislosti v [setu balíčků world](/wiki/World_set_(Portage) "World set (Portage)") mohou být znovu sestaveny za pomocí parametrů `--update --newuse` příkazu emerge:

`root #` `emerge --ask --update --newuse --verbose sys-boot/grub:2`

Software GRUB2 byl nyní nahrán do systému, ale ještě ne nainstalován.

### Instalace

Dále nainstalujte potřebné soubory GRUB2 do adresáře /boot/grub pomocí příkazu grub-install. Předpokládejme, že první disk (ten, ze kterého systém nabíhá) je /dev/sda, jeden z následujících souborů tak udělá co chceme:

- Při použití BIOS:

`root #` `grub-install /dev/sda`

For DOS/Legacy BIOS systems:

`root #` `grub-install /dev/sda`

- Při použití UEFI:

**Důležité**

Ujistěte se, že systémový oddíl EFI byl připojen před spuštěním grub-install. grub-install může nainstalovat soubor GRUB EFI (grubx64.efi) do špatného adresáře **bez toho**, aby **jakkoli** upozornil na to, že byl použit špatný adresář.

For UEFI systems:

`root #` `grub-install --target=x86_64-efi  --efi-directory=/boot`

Upon successful installation, the output should match the output of the previous command. If the output does not match exactly, then proceed to [Debugging GRUB](/wiki/Handbook:AMD64/Blocks/Bootloader/cs#Debugging_GRUB "Handbook:AMD64/Blocks/Bootloader/cs"), otherwise jump to the [Configure step](/wiki/Handbook:AMD64/Blocks/Bootloader/cs#GRUB_Configure "Handbook:AMD64/Blocks/Bootloader/cs").

##### Optional: Secure Boot

To successfully boot with secure boot enabled the signing certificate must either be accepted by the [UEFI](/wiki/UEFI "UEFI") firmware, or [shim](/wiki/Shim "Shim") must be used as a pre-loader. Shim is pre-signed with the third-party Microsoft Certificate, accepted by default by most UEFI motherboards.

How to configure the UEFI firmware to accept custom keys depends on the firmware vendor, which is beyond the scope of the handbook. Below is shown how to setup shim instead. Here it is assumed that the user has already followed the instructions in the previous sections to generate a signing key and to configure portage to use it. If this is not the case please return first to the [Kernel installation](/wiki/Handbook:AMD64/Installation/Kernel/cs "Handbook:AMD64/Installation/Kernel/cs") section.

The package [sys-boot/grub](https://packages.gentoo.org/packages/sys-boot/grub) installs a prebuilt and signed stand-alone EFI executable if the [secureboot](https://packages.gentoo.org/useflags/secureboot) [USE flag/cs (page does not exist)](/index.php?title=USE_flag/cs&action=edit&redlink=1 "USE flag/cs (page does not exist)") USE flag is enabled. Install the required packages and copy the stand-alone grub, Shim, and the MokManager to the same directory on the EFI System Partition. For example:

`root #` `emerge sys-boot/grub sys-boot/shim sys-boot/mokutil sys-boot/efibootmgr`

`root #` `cp /usr/share/shim/BOOTX64.EFI /efi/EFI/Gentoo/shimx64.efi
`

`root #` `cp /usr/share/shim/mmx64.efi /efi/EFI/Gentoo/mmx64.efi
`

`root #` `cp /usr/lib/grub/grub-x86_64.efi.signed /efi/EFI/Gentoo/grubx64.efi
`

Next register the signing key in shims MOKlist, this requires keys in the DER format, whereas sbsign and the kernel build system expect keys in the PEM format. In the previous sections of the handbook an example was shown to generate such a signing PEM key, this key must now be converted to the DER format:

`root #` `openssl x509 -in /path/to/kernel_key.pem -inform PEM -out /path/to/kernel_key.der -outform DER`

**Poznámka**

The path used here must be the path to the pem file containing the certificate belonging to the generated key. In this example both key and certificate are in the same pem file.

Then the converted certificate can be imported into Shims MOKlist, this command will ask to set some password for the import request:

`root #` `mokutil --import /path/to/kernel_key.der`

**Poznámka**

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

**Poznámka**

The import process will not be completed until the system is rebooted. After completing all steps in the handbook, restart the system and Shim will load, it will find the import request registered by mokutil. The MokManager application will start and ask for the password that was set when creating the import request. Follow the on-screen instructions to complete the import of the certificate, then reboot the system into the UEFI menu and enable the Secure Boot setting.

##### Debugging GRUB

When debugging GRUB, there are a couple of quick fixes that may result in a bootable installation without having to reboot to a new live image environment.

In the event that "EFI variables are not supported on this system" is displayed somewhere in the output, it is likely the live image was not booted in EFI mode and is presently in Legacy BIOS boot mode. The solution is to try the [removable GRUB step](/wiki/Handbook:AMD64/Blocks/Bootloader/cs#GRUB_Install_EFI_systems_removable "Handbook:AMD64/Blocks/Bootloader/cs") mentioned below. This will overwrite the executable EFI file located at /EFI/BOOT/BOOTX64.EFI. Upon rebooting in EFI mode, the motherboard firmware may execute this default boot entry and execute GRUB.

**Důležité**

Pokud grub-install vrátí chybu podobnou této: `Could not prepare Boot variable: Read-only file system`, bude potřeba k úspěšnému pokračování znovu připojit speciální přípojný bod efivars k zápisu:

`root #` `mount -o remount,rw /sys/firmware/efi/efivars`

`root #` `mount -o remount,rw,nosuid,nodev,noexec --types efivarfs efivarfs /sys/firmware/efi/efivars`

This is caused by certain non-official Gentoo environments not mounting the special EFI filesystem by default. If the previous command does not run, then reboot using an official Gentoo live image environment in EFI mode.

Někteří výrobci základních desek zřejmě podporují pouze adresář /efi/boot k umístění souboru .EFI na systémovém oddílu EFI (ESP). Instalátor GRUB může tuto operaci provést automaticky při použití volby `--removable`. Ověřte si, že je ESP připojeno před tím, že než spustíte následující příkazy. Za předpokladu, že je ESP připojena do /boot (jak bylo uvedeno výše), spusťte:

`root #` `grub-install --target=x86_64-efi --efi-directory=/boot --removable`

Tím vytvoříte výchozí adresář definovaný specifikací UEFI a následně zkopírujte soubor grubx64.efi do 'výchozího' umístění souborů EFI definované touto specifikací.

### Nastavení

Dále vygenerujte konfiguraci GRUB2 na základě nastavení specifikovaného v souboru /etc/default/grub a skriptech v /etc/grub.d. Ve většině případů není třeba žádných zásahů ze strany uživatele, jelikož GRUB2 automaticky zjistí, jaké jádro má zavést (to nejvyšší, jaké je k dispozici v /boot/ a který souborový systém je kořenový (root)). Je rovněž možné přidávat parametry jádra pomocí proměnné GRUB\_CMDLINE\_LINUX v souboru /etc/default/grub.

K vygenerování konečné konfigurace GRUB2 spusťte příkaz grub-mkconfig:

`root #` `grub-mkconfig -o /boot/grub/grub.cfg`

```
Generating grub.cfg ...
Found linux image: /boot/vmlinuz-6.18.8-gentoo
Found initrd image: /boot/initramfs-genkernel-amd64-6.18.8-gentoo
done

```

Výstup příkazu by měl obsahovat, že byl nalezen alespoň jeden obraz jádra, jelikož ten je potřeba k naběhnutí systému. Pokud je použit initramfs nebo byl k sestavení jádra použit genkernel, měl by být nalezen také správný obraz initrd. Pokud se tak nestalo, podívejte se na obsah adresáře /boot příkazem ls. Pokud tyto soubory opravdu chybí, vraťte se zpět k instrukcím pro nastavení a instalaci jádra.

**Tip**

Utilitu os-prober můžete použít ve spojení s GRUB2 k nalezení dalších operačních systémů na připojených discích. Detekovatelné jsou Windows 7, 8.1, 10 a ostatní distribuce Linuxu. Ti, kteří požadují, aby jejich systém podporoval dual-boot by měli nainstalovat (emerge) [sys-boot](https://packages.gentoo.org/packages/sys-boot) a po té znovu spustit příkaz grub-mkconfig (jak je uvedeno výše). Pokud narazíte na problémy s detekcí, určitě si přečtěte celý článek o [GRUB2](/wiki/GRUB2 "GRUB2") před tím, než požádáte komunitu Gentoo o podporu.

## Alternative 1: systemd-boot

Another option is [systemd-boot](/wiki/Systemd/systemd-boot "Systemd/systemd-boot"), which works on both OpenRC and systemd machines. It is a thin chainloader and works well with secure boot.

### Emerge

To install systemd-boot, enable the [boot](https://packages.gentoo.org/useflags/boot) [USE flag/cs (page does not exist)](/index.php?title=USE_flag/cs&action=edit&redlink=1 "USE flag/cs (page does not exist)") USE flag and re-install [sys-apps/systemd](https://packages.gentoo.org/packages/sys-apps/systemd) (for systemd systems) or [sys-apps/systemd-utils](https://packages.gentoo.org/packages/sys-apps/systemd-utils) (for OpenRC systems):

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

**Důležité**

Make sure the EFI system partition has been mounted _before_ running bootctl install.

When using this bootloader, before rebooting, verify that a new bootable entry exists using:

`root #` `bootctl list`

**Upozornění**

The kernel command line for new systemd-boot entries is read from /etc/kernel/cmdline or /usr/lib/kernel/cmdline. If neither file is present, then the kernel command line of the currently booted kernel is re-used (/proc/cmdline). On new installs it might therefore happen that the kernel command line of the live CD is accidentally used to boot the new kernel. The kernel command line for registered entries can be checked with:

`root #` `bootctl list`

If this does not show the desired kernel command line then create /etc/kernel/cmdline containing the correct kernel command line and re-install the kernel.

If no new entry exists, ensure the [sys-kernel/installkernel](https://packages.gentoo.org/packages/sys-kernel/installkernel) package has been installed with the [systemd](https://packages.gentoo.org/useflags/systemd) [USE flag/cs (page does not exist)](/index.php?title=USE_flag/cs&action=edit&redlink=1 "USE flag/cs (page does not exist)") and [systemd-boot](https://packages.gentoo.org/useflags/systemd-boot) [USE flag/cs (page does not exist)](/index.php?title=USE_flag/cs&action=edit&redlink=1 "USE flag/cs (page does not exist)") USE flags enabled, and re-run the kernel installation.

For the distribution kernels:

`root #` `emerge --ask --config sys-kernel/gentoo-kernel`

For a manually configured and compiled kernel:

`root #` `make install`

**Důležité**

When installing kernels for systemd-boot, no root= kernel command line argument is added by default. On systemd systems that are using an initramfs users may rely instead on [systemd-gpt-auto-generator](/wiki/Systemd#Automatic_mounting_of_partitions_at_boot.2Fcs "Systemd") to automatically find the root partition at boot. Otherwise users should manually specify the location of the root partition by setting root= in /etc/kernel/cmdline as well as any other kernel command line arguments that should be used. And then reinstalling the kernel as described above.

### Optional: Secure Boot

When the [secureboot](https://packages.gentoo.org/useflags/secureboot) [USE flag/cs (page does not exist)](/index.php?title=USE_flag/cs&action=edit&redlink=1 "USE flag/cs (page does not exist)") USE flag is enabled, the systemd-boot EFI executable will be signed by portage automatically. Furthermore, bootctl install will automatically install the signed version.

To successfully boot with secure boot enabled the used certificate must either be accepted by the [UEFI](/wiki/UEFI "UEFI") firmware, or [shim](/wiki/Shim "Shim") must be used as a pre-loader. Shim is pre-signed with the third-party Microsoft Certificate, accepted by default by most UEFI motherboards.

How to configure the UEFI firmware to accept custom keys depends on the firmware vendor, which is beyond the scope of the handbook. Below is shown how to setup shim instead. Here it is assumed that the user has already followed the instructions in the previous sections to generate a signing key and to configure portage to use it. If this is not the case please return first to the [Kernel installation](/wiki/Handbook:AMD64/Installation/Kernel/cs "Handbook:AMD64/Installation/Kernel/cs") section.

`root #` `emerge --ask sys-boot/shim sys-boot/mokutil sys-boot/efibootmgr`

`root #` `bootctl install --no-variables`

`root #` `cp /usr/share/shim/BOOTX64.EFI /efi/EFI/systemd/shimx64.efi`

`root #` `cp /usr/share/shim/mmx64.efi /efi/EFI/systemd/mmx64.efi`

Shims MOKlist requires keys in the DER format, whereas sbsign and the kernel build system expect keys in the PEM format. In the previous sections of the handbook an example was shown to generate such a signing PEM key, this key must now be converted to the DER format:

`root #` `openssl x509 -in /path/to/kernel_key.pem -inform PEM -out /path/to/kernel_key.der -outform DER`

**Poznámka**

The path used here must be the path to the pem file containing the certificate belonging to the generated key. In this example both key and certificate are in the same pem file.

Then the converted certificate can be imported into Shims MOKlist:

`root #` `mokutil --import /path/to/kernel_key.der`

**Poznámka**

When the currently booted kernel already trusts the certificate being imported, the message "Already in kernel trusted keyring." will be returned here. If this happens, re-run the above command with the argument --ignore-keyring added.

And finally we register Shim with the UEFI firmware. In the following command, `boot-disk` and `boot-partition-id` must be replaced with the disk and partition identifier of the EFI system partition:

`root #` `efibootmgr --create --disk /dev/boot-disk --part boot-partition-id --loader '\EFI\systemd\shimx64.efi'  --label 'Systemd-boot via Shim' --unicode '\EFI\systemd\systemd-bootx64.efi'`

**Poznámka**

The import process will not be completed until the system is rebooted. After completing all steps in the handbook, restart the system and Shim will load, it will find the import request registered by mokutil. The MokManager application will start and ask for the password that was set when creating the import request. Follow the on-screen instructions to complete the import of the certificate, then reboot the system into the UEFI menu and enable the Secure Boot setting.

## Alternativa 2: efibootmgr

V systémech s UEFI může být upravován přímo firmware UEFI (jinými slovy primární zavaděč systému) k tomu, aby vyhledal položky je spuštění. Takové systému nepotřebují zavaděče jako je GRUB2 (známé jako sekundární zavaděče) k tomu, aby je pomohly spustit. Se vším tímto na paměti, důvodem, proč existují EFI zavaděče jako je GRUB2, je rozšíření funkcí UEFI systémů v průběhu procesu spouštění. Použití efibootmgr je vhodné pro ty, kdo chtějí využit opravdu minimalistický (avšak více svázaný) přístup ke spouštění svého systému; použití GRUB2 (viz výše) je pro většinu uživatelů jednodušší, protože je z hlediska svých možností při spouštění systému s UEFI pružnější.

System administrators who desire to take a minimalist, although more rigid, approach to booting the system can avoid secondary bootloaders and boot the Linux kernel as an [EFI stub](/wiki/EFI_stub "EFI stub").

Uvědomte si, že aplikace [sys-boot/efibootmgr](https://packages.gentoo.org/packages/sys-boot/efibootmgr) není zavaděč; je to nástroj sloužící k přizpůsobení firmwaru UEFI a aktualizaci jeho nastavení tak, aby mohlo být dříve do systému nainstalované linuxové jádro spuštěno i s dodatečnými volbami (je-li to třeba) nebo aby mohlo být zapnuto více spustitelných položek. Toto přizpůsobení se děje pomocí proměnných EFI (proto potřeba mít v jádře podporu proměnných EFI).

Před tím, než budete pokračovat, pročtěte si článek [EFI stub kernel](/wiki/EFI_stub_kernel "EFI stub kernel"). Jádro musí mít zapnuty specifické volby, aby bylo přímo spustitelné ze strany UEFI firmwaru. Možná bude potřeba jádro překompilovat. Dobrý nápad je také podívat se na článek [efibootmgr](/wiki/Efibootmgr "Efibootmgr").

It is also a good idea to take a look at the [efibootmgr](/wiki/Efibootmgr "Efibootmgr") article for additional information.

**Poznámka**

Pro uvedení věcí na pravou míru, efibootmgr není pro spuštění systému UEFI vyžadováno. Linuxové jádro samo o sobě může být spuštěno napřímo, dodatečné volby mohou být zabudovány přímo do linuxového jádra (v konfiguraci jádra je volba, která umožňuje uživateli specifikovat parametry spuštění). Dokonce i initramfs může být "součástí" jádra.

Kdo se rozhodl jít touto cestou, musí nainstalovat příslušný software:

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

Po té vytvořte adresář /boot/efi/boot a zkopírujte do něj jádro, které nazvete bzImage.efi:

`root #` `mkdir -p /boot/efi/boot
`

`root #` `cp /boot/vmlinuz-* /boot/efi/boot/bzImage.efi
`

Install the efibootmgr package:

`root #` `emerge --ask sys-boot/efibootmgr`

Dále sdělte firmwaru UEFI, že má být vytvořena položka "Gentoo" s čerstvě sestaveným jádrem:

`root #` `efibootmgr --create --disk /dev/sda --part 2 --label "Gentoo" --loader "\efi\boot\bzImage.efi"`

**Poznámka**

Použití `\` jako oddělovače adresářů je v rámci definicí pro UEFI povinné.

Pokud je použit zaváděcí souborový systém RAM (initramfs), přidejte k němu odpovídající volbu:

`root #` `efibootmgr -c -d /dev/sda -p 2 -L "Gentoo" -l "\efi\boot\bzImage.efi" initrd='\initramfs-genkernel-amd64-6.18.8-gentoo'`

**Tip**

Additional kernel command line options may be parsed by the firmware to the kernel by specifying them along with the initrd=... option as shown above.

Po těchto změnách, jakmile systém znovunaběhne, bude zpřístupněna položka "Gentoo".

### Unified Kernel Image

If [installkernel](/wiki/Installkernel "Installkernel") was configured to build and install unified kernel images. The unified kernel image should already be installed to the EFI/Linux directory on the EFI system partition, if this is not the case ensure the directory exists and then run the kernel installation again as described earlier in the handbook.

To add a direct boot entry for the installed unified kernel image:

`root #` `efibootmgr --create --disk /dev/sda --part 1 --label "gentoo" --loader "\EFI\Linux\gentoo-x.y.z.efi"`

## Other Alternatives

For other options that are not covered in the Handbook, see the full list of available [bootloaders](/wiki/Bootloader "Bootloader").

## Restartování systému

Opusťte prostředí chroot a odpojte všechny připojené oddíly. Po té napište kouzelný příkaz, který zahájí konečný, opravdový test: reboot.

`root #` `exit`

`cdimage ~#` `cd
`

`cdimage ~#` `umount -l /mnt/gentoo/dev{/shm,/pts,}
`

`cdimage ~#` `umount -R /mnt/gentoo
`

`cdimage ~#` `reboot`

Samozřejmě nezapomeňte vyjmout bootovací CD, jinak může dojít opět ke startu z CD namísto nově instalovaného systému Gentoo.

Jakmile jste zrestartovali do čerstvě nainstalovaného prostředí Gentoo, dokončete vše podle [Dokončení instalace Gentoo](/wiki/Handbook:AMD64/Installation/Finalizing "Handbook:AMD64/Installation/Finalizing").

[← Instalace nástrojů](/wiki/Handbook:AMD64/Installation/Tools/cs "Previous part") [Home](/wiki/Handbook:AMD64/cs "Handbook:AMD64/cs") [Dokončení →](/wiki/Handbook:AMD64/Installation/Finalizing/cs "Next part")