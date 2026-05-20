# Bootloader

Other languages:

- Deutsch
- [English](/wiki/Handbook:AMD64/Installation/Bootloader "Handbook:AMD64/Installation/Bootloader (100% translated)")
- [español](/wiki/Handbook:AMD64/Installation/Bootloader/es "Handbook:AMD64/Installation/Bootloader/es (100% translated)")
- [français](/wiki/Handbook:AMD64/Installation/Bootloader/fr "Handbook:AMD64/Installation/Bootloader/fr (100% translated)")
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

[AMD64 Handbuch](/wiki/Handbook:AMD64/de "Handbook:AMD64/de")[Installation](/wiki/Handbook:AMD64/Full/Installation/de "Handbook:AMD64/Full/Installation/de")[Über die Installation](/wiki/Handbook:AMD64/Installation/About/de "Handbook:AMD64/Installation/About/de")[Auswahl des Mediums](/wiki/Handbook:AMD64/Installation/Media/de "Handbook:AMD64/Installation/Media/de")[Konfiguration des Netzwerks](/wiki/Handbook:AMD64/Installation/Networking/de "Handbook:AMD64/Installation/Networking/de")[Vorbereiten der Festplatte(n)](/wiki/Handbook:AMD64/Installation/Disks/de "Handbook:AMD64/Installation/Disks/de")[Installation des Stage Archivs](/wiki/Handbook:AMD64/Installation/Stage/de "Handbook:AMD64/Installation/Stage/de")[Installation des Basissystems](/wiki/Handbook:AMD64/Installation/Base/de "Handbook:AMD64/Installation/Base/de")[Konfiguration des Kernels](/wiki/Handbook:AMD64/Installation/Kernel/de "Handbook:AMD64/Installation/Kernel/de")[Konfiguration des Systems](/wiki/Handbook:AMD64/Installation/System/de "Handbook:AMD64/Installation/System/de")[Installation der Tools](/wiki/Handbook:AMD64/Installation/Tools/de "Handbook:AMD64/Installation/Tools/de")Konfiguration des Bootloaders[Abschluss](/wiki/Handbook:AMD64/Installation/Finalizing/de "Handbook:AMD64/Installation/Finalizing/de")[Arbeiten mit Gentoo](/wiki/Handbook:AMD64/Full/Working/de "Handbook:AMD64/Full/Working/de")[Portage-Einführung](/wiki/Handbook:AMD64/Working/Portage/de "Handbook:AMD64/Working/Portage/de")[USE-Flags](/wiki/Handbook:AMD64/Working/USE/de "Handbook:AMD64/Working/USE/de")[Portage-Features](/wiki/Handbook:AMD64/Working/Features/de "Handbook:AMD64/Working/Features/de")[Initskript-System](/wiki/Handbook:AMD64/Working/Initscripts/de "Handbook:AMD64/Working/Initscripts/de")[Umgebungsvariablen](/wiki/Handbook:AMD64/Working/EnvVar/de "Handbook:AMD64/Working/EnvVar/de")[Arbeiten mit Portage](/wiki/Handbook:AMD64/Full/Portage/de "Handbook:AMD64/Full/Portage/de")[Dateien und Verzeichnisse](/wiki/Handbook:AMD64/Portage/Files/de "Handbook:AMD64/Portage/Files/de")[Variablen](/wiki/Handbook:AMD64/Portage/Variables/de "Handbook:AMD64/Portage/Variables/de")[Mischen von Softwarezweigen](/wiki/Handbook:AMD64/Portage/Branches/de "Handbook:AMD64/Portage/Branches/de")[Zusätzliche Tools](/wiki/Handbook:AMD64/Portage/Tools/de "Handbook:AMD64/Portage/Tools/de")[Eigener Portage-Tree](/wiki/Handbook:AMD64/Portage/CustomTree/de "Handbook:AMD64/Portage/CustomTree/de")[Erweiterte Portage-Features](/wiki/Handbook:AMD64/Portage/Advanced/de "Handbook:AMD64/Portage/Advanced/de")[Netzwerk-Konfiguration](/wiki/Handbook:AMD64/Full/Networking/de "Handbook:AMD64/Full/Networking/de")[Zu Beginn](/wiki/Handbook:AMD64/Networking/Introduction/de "Handbook:AMD64/Networking/Introduction/de")[Fortgeschrittene Konfiguration](/wiki/Handbook:AMD64/Networking/Advanced/de "Handbook:AMD64/Networking/Advanced/de")[Modulare Vernetzung](/wiki/Handbook:AMD64/Networking/Modular/de "Handbook:AMD64/Networking/Modular/de")[Drahtlose Netzwerke](/wiki/Handbook:AMD64/Networking/Wireless/de "Handbook:AMD64/Networking/Wireless/de")[Funktionalität hinzufügen](/wiki/Handbook:AMD64/Networking/Extending/de "Handbook:AMD64/Networking/Extending/de")[Dynamisches Management](/wiki/Handbook:AMD64/Networking/Dynamic/de "Handbook:AMD64/Networking/Dynamic/de")

## Contents

- [1Einen Bootloader auswählen](#Einen_Bootloader_ausw.C3.A4hlen)
- [2Standard: GRUB](#Standard:_GRUB)
  - [2.1Emerge](#Emerge)
  - [2.2Installation](#Installation)
  - [2.3DOS/Legacy-BIOS-Systeme](#DOS.2FLegacy-BIOS-Systeme)
  - [2.4UEFI-Systeme](#UEFI-Systeme)
    - [2.4.1Optional: Secure Boot](#Optional:_Secure_Boot)
    - [2.4.2Debugging GRUB](#Debugging_GRUB)
  - [2.5Konfiguration](#Konfiguration)
- [3Alternative 1: systemd-boot](#Alternative_1:_systemd-boot)
  - [3.1Emerge](#Emerge_2)
  - [3.2Installation](#Installation_2)
  - [3.3Optional: Secure Boot](#Optional:_Secure_Boot_2)
- [4Alternative 2: EFI Stub](#Alternative_2:_EFI_Stub)
  - [4.1Unified Kernel Image](#Unified_Kernel_Image)
- [5Other Alternatives](#Other_Alternatives)
- [6Neustart des Systems](#Neustart_des_Systems)

## Einen Bootloader auswählen

Nachdem der Linux Kernel konfiguriert ist, die System Tools installiert sind und die wichtigsten System Konfigurationsdateien angepasst sind, ist es nun an der Zeit, den letzten wichtigen Teil eines Linux-Systems zu installieren: den Bootloader.

Der Bootloader ist beim Booten dafür zuständig, den Linux Kernel zu starten. Ohne ihn würde das System nach dem Druck auf den Power-Knopf nicht wissen, wie es weiter vorgehen soll.

Wir zeigen für die **amd64** Architektur, wie man [GRUB](/wiki/Handbook:AMD64/Blocks/Bootloader/de#Default:GRUB "Handbook:AMD64/Blocks/Bootloader/de") für DOS/Legacy-BIOS basierte Systeme und [GRUB](#Standard:_GRUB),
[systemd-boot](#Alternative_1:_systemd-boot) oder
[EFI Stub](#Alternative_2:_efibootmgr) für UEFI Systeme konfiguriert.

In diesem Abschnitt des Handbuchs wird unterschieden zwischen dem "Emerge" eines Bootloader-Pakets und dem "Installieren" des Bootloaders auf die System-Festplatte. Der Ausdruck "emerge" wird verwendet, wenn [Portage](/wiki/Portage "Portage") aufgerufen wird, um eine Software auf dem System verfügbar zu machen. Der Ausdruck "installieren" wird verwendet, wenn der Bootloader Dateien in spezielle Bereiche der System-Festplatte kopiert oder dort Daten verändert, um den Bootloader so zu aktivieren, dass er beim nächsten Systemstart gestartet wird.

## Standard: GRUB

Die Mehrzahl der Gentoo Linux Systeme verwendet heutzutage [GRUB](/wiki/GRUB "GRUB") als Bootloader. GRUB (Paket [sys-boot/grub](https://packages.gentoo.org/packages/sys-boot/grub)) ist der direkte Nachfolger von [GRUB Legacy](/wiki/GRUB_Legacy "GRUB Legacy"). Ohne weitere Konfiguration unterstützt GRUB ältere BIOS ("pc") Systeme. Mit zusätzlicher Konfiguration unterstützt GRUB mehr als ein halbes Dutzend anderer Plattformen. Details finden Sie in dem Abschnitt [Prerequisites section](/wiki/GRUB#Prerequisites "GRUB") des [GRUB](/wiki/GRUB "GRUB") Artikels im Gentoo Wiki.

### Emerge

Wenn ein älteres BIOS System verwendet wird, das nur MBR Partitions-Tabellen unterstützt, sind keine Konfigurationsarbeiten erforderlich, um GRUB zu "emergen":

`root #` `emerge --ask --verbose sys-boot/grub`

Für UEFI Systeme: die Ausführung des obigen Kommandos zeigt vor dem "emerge" die aktivierten GRUB\_PLATFORMS Werte an. Wenn UEFI Systeme verwendet werden, müssen Anwender sicherstellen, dass `GRUB_PLATFORMS="efi-64"` aktiviert ist (was standardmäßig der Fall sein sollte). Wenn es nicht der Fall ist, muss `GRUB_PLATFORMS="efi-64"` zu der Datei /etc/portage/make.conf hinzugefügt werden, _bevor_ das Emerge-Kommando für GRUB ausgeführt wird. Ansonsten wird GRUB EFI nicht unterstützen.

`root #` `echo 'GRUB_PLATFORMS="efi-64"' >> /etc/portage/make.conf`

`root #` `emerge --ask sys-boot/grub`

Wenn GRUB installiert wurde ohne dass `GRUB_PLATFORMS="efi-64"` aktiviert war, kann die Zeile (wie oben gezeigt) zu make.conf hinzugefügt werden. Danach können die Abhängigkeiten für das [world package set](/wiki/World_set_(Portage) "World set (Portage)") neu berechnet werden durch Aufruf von emerge mit den Optionen `--update --newuse`:

`root #` `emerge --ask --update --newuse --verbose sys-boot/grub`

Die GRUB-Software wurde nun in das _System_ gemerged, aber noch nicht als sekundärer _Bootloader_ installiert.

### Installation

Als nächstes werden die erforderlichen GRUB Dateien in das /boot/grub/ Verzeichnis installiert. Hierfür kann das grub-install Kommando verwendet werden. Unter der Annahme, dass die erste Festplatte (diejenige von der das System bootet) /dev/sda ist, kann einer der folgenden Befehle verwenden werden:

### DOS/Legacy-BIOS-Systeme

For DOS/Legacy BIOS systems:

`root #` `grub-install /dev/sda`

### UEFI-Systeme

**Wichtig**

Stellen Sie sicher, dass die EFI System-Partition (ESP) eingehängt wurde, _bevor_ Sie grub-install ausführen. Wenn die EFI System-Partition nicht eingehängt wurde, wird grub-install die Installation trotzdem durchführen. Die GRUB EFI-Datei (grubx64.efi) wird dann in das falsche Verzeichnis installiert. Dabei werden Sie **nicht** informiert oder gewarnt.

For UEFI systems:

`root #` `grub-install --efi-directory=/efi`

```
Installing for x86_64-efi platform.
Installation finished. No error reported.

```

Upon successful installation, the output should match the output of the previous command. If the output does not match exactly, then proceed to [Debugging GRUB](/wiki/Handbook:AMD64/Blocks/Bootloader/de#Debugging_GRUB "Handbook:AMD64/Blocks/Bootloader/de"), otherwise jump to the [Configure step](/wiki/Handbook:AMD64/Blocks/Bootloader/de#GRUB_Configure "Handbook:AMD64/Blocks/Bootloader/de").

##### Optional: Secure Boot

To successfully boot with secure boot enabled the signing certificate must either be accepted by the [UEFI](/wiki/UEFI "UEFI") firmware, or [shim](/wiki/Shim "Shim") must be used as a pre-loader. Shim is pre-signed with the third-party Microsoft Certificate, accepted by default by most UEFI motherboards.

How to configure the UEFI firmware to accept custom keys depends on the firmware vendor, which is beyond the scope of the handbook. Below is shown how to setup shim instead. Here it is assumed that the user has already followed the instructions in the previous sections to generate a signing key and to configure portage to use it. If this is not the case please return first to the [Kernel installation](/wiki/Handbook:AMD64/Installation/Kernel/de "Handbook:AMD64/Installation/Kernel/de") section.

The package [sys-boot/grub](https://packages.gentoo.org/packages/sys-boot/grub) installs a prebuilt and signed stand-alone EFI executable if the [secureboot](https://packages.gentoo.org/useflags/secureboot) [USE flag/de](/wiki/USE_flag/de "USE flag/de") USE flag is enabled. Install the required packages and copy the stand-alone grub, Shim, and the MokManager to the same directory on the EFI System Partition. For example:

`root #` `emerge sys-boot/grub sys-boot/shim sys-boot/mokutil sys-boot/efibootmgr`

`root #` `cp /usr/share/shim/BOOTX64.EFI /efi/EFI/Gentoo/shimx64.efi
`

`root #` `cp /usr/share/shim/mmx64.efi /efi/EFI/Gentoo/mmx64.efi
`

`root #` `cp /usr/lib/grub/grub-x86_64.efi.signed /efi/EFI/Gentoo/grubx64.efi
`

Next register the signing key in shims MOKlist, this requires keys in the DER format, whereas sbsign and the kernel build system expect keys in the PEM format. In the previous sections of the handbook an example was shown to generate such a signing PEM key, this key must now be converted to the DER format:

`root #` `openssl x509 -in /path/to/kernel_key.pem -inform PEM -out /path/to/kernel_key.der -outform DER`

**Hinweis**

The path used here must be the path to the pem file containing the certificate belonging to the generated key. In this example both key and certificate are in the same pem file.

Then the converted certificate can be imported into Shims MOKlist, this command will ask to set some password for the import request:

`root #` `mokutil --import /path/to/kernel_key.der`

**Hinweis**

When the currently booted kernel already trusts the certificate being imported, the message "Already in kernel trusted keyring." will be returned here. If this happens, re-run the above command with the argument --ignore-keyring added.

Next, register Shim with the UEFI firmware. In the following command, `boot-disk` and `boot-partition-id` must be replaced with the disk and partition identifier of the EFI system partition:

`root #` `efibootmgr --create --disk /dev/boot-disk --part boot-partition-id --loader '\EFI\Gentoo\shimx64.efi' --label 'GRUB via Shim' --unicode`

Note that this prebuilt and signed stand-alone version of grub reads the grub.cfg from a different location then usual. Instead of the default /boot/grub/grub.cfg the config file should be in the same directory that the grub EFI executable is in, e.g. /efi/EFI/Gentoo/grub.cfg. When [sys-kernel/installkernel](https://packages.gentoo.org/packages/sys-kernel/installkernel) is used to install the kernel and update the grub configuration then the GRUB\_CFG environment variable may be used to override the usual location of the grub config file.

For example:

`root #` `grub-mkconfig -o /efi/EFI/Gentoo/grub.cfg`

Or, via [installkernel](/wiki/Installkernel "Installkernel"):

DATEI **`/etc/env.d/99grub`**

```
GRUB_CFG=/efi/EFI/Gentoo/grub.cfg

```

`root #` `env-update`

**Hinweis**

The import process will not be completed until the system is rebooted. After completing all steps in the handbook, restart the system and Shim will load, it will find the import request registered by mokutil. The MokManager application will start and ask for the password that was set when creating the import request. Follow the on-screen instructions to complete the import of the certificate, then reboot the system into the UEFI menu and enable the Secure Boot setting.

##### Debugging GRUB

When debugging GRUB, there are a couple of quick fixes that may result in a bootable installation without having to reboot to a new live image environment.

In the event that "EFI variables are not supported on this system" is displayed somewhere in the output, it is likely the live image was not booted in EFI mode and is presently in Legacy BIOS boot mode. The solution is to try the [removable GRUB step](/wiki/Handbook:AMD64/Blocks/Bootloader/de#GRUB_Install_EFI_systems_removable "Handbook:AMD64/Blocks/Bootloader/de") mentioned below. This will overwrite the executable EFI file located at /EFI/BOOT/BOOTX64.EFI. Upon rebooting in EFI mode, the motherboard firmware may execute this default boot entry and execute GRUB.

**Wichtig**

Wenn grub-install einen Fehler wie `Could not prepare Boot variable: Read-only file system` meldet, ist es möglicherweise erforderlich, einen Mount-Parameter des efivars-Dateisystems auf "read-write" zu ändern:

`root #` `mount -o remount,rw,nosuid,nodev,noexec --types efivarfs efivarfs /sys/firmware/efi/efivars`

`root #` `mount -o remount,rw,nosuid,nodev,noexec --types efivarfs efivarfs /sys/firmware/efi/efivars`

This is caused by certain non-official Gentoo environments not mounting the special EFI filesystem by default. If the previous command does not run, then reboot using an official Gentoo live image environment in EFI mode.

Einige Hersteller von Mainboards scheinen nur das Verzeichnis /efi/boot/ für die .EFI-Datei in der EFI System Partition (ESP) zu unterstützen. Der GRUB Installer unterstützt diese Arbeitsweise mit der Option `--removable`. Stellen Sie sicher, dass die ESP eingehängt wurde, bevor Sie die folgenden Kommandos ausführen. Vorausgesetzt, dass die ESP eingehängt ist unter /boot (wie früher vorgeschlagen), können Sie folgende Kommandos ausführen:

`root #` `grub-install --target=x86_64-efi --efi-directory=/efi --removable`

Dies erzeugt das von der UEFI Spezifikation definierte Standard-Verzeichnis und kopiert dann die grubx64.efi-Datei zu dem "Standard"-Ort der EFI-Datei, der in der gleichen Spezifikation definiert wurde.

### Konfiguration

Im nächsten Schritt erzeugen wir auf Grundlage der Benutzereinstellungen, die in der Datei /etc/default/grub und den Skripten im Verzeichnis /etc/grub.d angegeben sind, die GRUB Konfiguration. In den meisten Fällen ist keine Konfiguration durch den Benutzer erforderlich, weil GRUB automatisch erkennen wird, welcher Kernel zu booten ist (den höchsten verfügbaren in /boot/) und was das Root Dateisystem ist. Mit Hilfe der GRUB\_CMDLINE\_LINUX Variable ist auch möglich, Kernel-Parameter in /etc/default/grub zu definieren.

Zum Generieren der endgültigen GRUB Konfiguration führen Sie den Befehl grub-mkconfig aus:

`root #` `grub-mkconfig -o /boot/grub/grub.cfg`

```
Generating grub.cfg ...
Found linux image: /boot/vmlinuz-6.18.8-gentoo
Found initrd image: /boot/initramfs-genkernel-amd64-6.18.8-gentoo
done

```

Die Ausgabe des Befehls muss erwähnen, dass mindestens ein Linux Image gefunden wurde, da dieses zum Booten des Systems erforderlich sind. Wenn ein initramfs verwendet wird, oder der Kernel mit Hilfe von genkernel erzeugt wurde, sollte das korrekte initrd Image ebenfalls erkannt werden. Falls dies nicht der Fall ist, überprüfen Sie das Verzeichnis /boot/ mit dem Befehl ls auf dessen Inhalt. Wenn die Dateien in der Tat fehlen sollten, gehen Sie zurück zur Kernel-Konfiguration und der dortigen Installationsanleitung.

**Tipp**

Die os-prober Utility kann in Verbindung mit GRUB verwendet werden, um andere Betriebssysteme auf angeschlossenen Festplatten zu erkennen. Windows 7, 8.1, 10 und andere Linux Distributionen werden erkannt. Diejenigen, die ein Dual-Boot System wünschen, sollten das [sys-boot/os-prober](https://packages.gentoo.org/packages/sys-boot/os-prober) Paket "emergen" und dann das
grub-mkconfig Kommando erneut ausführen (wie oben beschrieben). Falls Erkennungsprobleme auftreten, lesen Sie bitte den [GRUB](/wiki/GRUB "GRUB") Artikel vollständig, _bevor_ Sie die Gentoo Community um Hilfe bitten.

## Alternative 1: systemd-boot

Another option is [systemd-boot](/wiki/Systemd/systemd-boot "Systemd/systemd-boot"), which works on both OpenRC and systemd machines. It is a thin chainloader and works well with secure boot.

### Emerge

To install systemd-boot, enable the [boot](https://packages.gentoo.org/useflags/boot) [USE flag/de](/wiki/USE_flag/de "USE flag/de") USE flag and re-install [sys-apps/systemd](https://packages.gentoo.org/packages/sys-apps/systemd) (for systemd systems) or [sys-apps/systemd-utils](https://packages.gentoo.org/packages/sys-apps/systemd-utils) (for OpenRC systems):

DATEI **`/etc/portage/package.use/systemd-boot`**

```
sys-apps/systemd boot
sys-apps/systemd-utils boot

```

`root #` `emerge --ask sys-apps/systemd`

Or

`root #` `emerge --ask sys-apps/systemd-utils`

### Installation

Now, install the systemd-boot loader to the [EFI System Partition](/wiki/EFI_System_Partition/de "EFI System Partition/de"):

`root #` `bootctl install`

**Wichtig**

Make sure the EFI system partition has been mounted _before_ running bootctl install.

When using this bootloader, before rebooting, verify that a new bootable entry exists using:

`root #` `bootctl list`

**Warnung**

The kernel command line for new systemd-boot entries is read from /etc/kernel/cmdline or /usr/lib/kernel/cmdline. If neither file is present, then the kernel command line of the currently booted kernel is re-used (/proc/cmdline). On new installs it might therefore happen that the kernel command line of the live CD is accidentally used to boot the new kernel. The kernel command line for registered entries can be checked with:

`root #` `bootctl list`

If this does not show the desired kernel command line then create /etc/kernel/cmdline containing the correct kernel command line and re-install the kernel.

If no new entry exists, ensure the [sys-kernel/installkernel](https://packages.gentoo.org/packages/sys-kernel/installkernel) package has been installed with the [systemd](https://packages.gentoo.org/useflags/systemd) [USE flag/de](/wiki/USE_flag/de "USE flag/de") and [systemd-boot](https://packages.gentoo.org/useflags/systemd-boot) [USE flag/de](/wiki/USE_flag/de "USE flag/de") USE flags enabled, and re-run the kernel installation.

For the distribution kernels:

`root #` `emerge --ask --config sys-kernel/gentoo-kernel`

For a manually configured and compiled kernel:

`root #` `make install`

**Wichtig**

When installing kernels for systemd-boot, no root= kernel command line argument is added by default. On systemd systems that are using an initramfs users may rely instead on [systemd-gpt-auto-generator](/wiki/Systemd#Automatic_mounting_of_partitions_at_boot.2Fde "Systemd") to automatically find the root partition at boot. Otherwise users should manually specify the location of the root partition by setting root= in /etc/kernel/cmdline as well as any other kernel command line arguments that should be used. And then reinstalling the kernel as described above.

### Optional: Secure Boot

When the [secureboot](https://packages.gentoo.org/useflags/secureboot) [USE flag/de](/wiki/USE_flag/de "USE flag/de") USE flag is enabled, the systemd-boot EFI executable will be signed by portage automatically. Furthermore, bootctl install will automatically install the signed version.

To successfully boot with secure boot enabled the used certificate must either be accepted by the [UEFI](/wiki/UEFI "UEFI") firmware, or [shim](/wiki/Shim "Shim") must be used as a pre-loader. Shim is pre-signed with the third-party Microsoft Certificate, accepted by default by most UEFI motherboards.

How to configure the UEFI firmware to accept custom keys depends on the firmware vendor, which is beyond the scope of the handbook. Below is shown how to setup shim instead. Here it is assumed that the user has already followed the instructions in the previous sections to generate a signing key and to configure portage to use it. If this is not the case please return first to the [Kernel installation](/wiki/Handbook:AMD64/Installation/Kernel/de "Handbook:AMD64/Installation/Kernel/de") section.

`root #` `emerge --ask sys-boot/shim sys-boot/mokutil sys-boot/efibootmgr`

`root #` `bootctl install --no-variables`

`root #` `cp /usr/share/shim/BOOTX64.EFI /efi/EFI/systemd/shimx64.efi`

`root #` `cp /usr/share/shim/mmx64.efi /efi/EFI/systemd/mmx64.efi`

Shims MOKlist requires keys in the DER format, whereas sbsign and the kernel build system expect keys in the PEM format. In the previous sections of the handbook an example was shown to generate such a signing PEM key, this key must now be converted to the DER format:

`root #` `openssl x509 -in /path/to/kernel_key.pem -inform PEM -out /path/to/kernel_key.der -outform DER`

**Hinweis**

The path used here must be the path to the pem file containing the certificate belonging to the generated key. In this example both key and certificate are in the same pem file.

Then the converted certificate can be imported into Shims MOKlist:

`root #` `mokutil --import /path/to/kernel_key.der`

**Hinweis**

When the currently booted kernel already trusts the certificate being imported, the message "Already in kernel trusted keyring." will be returned here. If this happens, re-run the above command with the argument --ignore-keyring added.

And finally we register Shim with the UEFI firmware. In the following command, `boot-disk` and `boot-partition-id` must be replaced with the disk and partition identifier of the EFI system partition:

`root #` `efibootmgr --create --disk /dev/boot-disk --part boot-partition-id --loader '\EFI\systemd\shimx64.efi'  --label 'Systemd-boot via Shim' --unicode '\EFI\systemd\systemd-bootx64.efi'`

**Hinweis**

The import process will not be completed until the system is rebooted. After completing all steps in the handbook, restart the system and Shim will load, it will find the import request registered by mokutil. The MokManager application will start and ask for the password that was set when creating the import request. Follow the on-screen instructions to complete the import of the certificate, then reboot the system into the UEFI menu and enable the Secure Boot setting.

## Alternative 2: EFI Stub

Auf UEFI basierten Systemen kann die UEFI Firmware auf dem System selbst (mit anderen Worten: der primäre Bootloader) dazu gebracht werden, dass sie selbst nach UEFI Boot-Einträgen sucht. Solche Systeme benötigen keine (sekundären) Bootloader wie GRUB, um das System zu booten. Auf solchen Systemen werden Bootloader wie GRUB genutzt, weil sie erweiterte Funktionalität bieten. efibootmgr ist für diejenigen gedacht, die ihr System auf direktem (und evtl. rigidem) Weg booten wollen. GRUB ist einfacher für die Mehrzahl der Anwender, weil es beim Booten von UEFI Systemen mehr Flexibilität bietet.

System administrators who desire to take a minimalist, although more rigid, approach to booting the system can avoid secondary bootloaders and boot the Linux kernel as an [EFI stub](/wiki/EFI_stub "EFI stub").

Vergessen Sie nicht, dass die [sys-boot/efibootmgr](https://packages.gentoo.org/packages/sys-boot/efibootmgr) Anwendung kein Bootloader ist, sondern ein Werkzeug, um mit der UEFI-Firmware zu interagieren und deren Einstellungen zu aktualisieren. Auf diese Weise kann ein Kernel, der früher installiert wurde, mit zusätzlichen Optionen gebootet werden (falls nötig). Des Weiteren lassen sich dadurch mehrere Booteinträge realisieren. Diese Interaktion erfolgt durch EFI Variablen (daher die Notwendigkeit für die Kernel-Unterstützung für EFI Variablen).

Bitte lesen Sie den [EFI stub](/wiki/EFI_stub "EFI stub") kernel Artikel, bevor Sie fortfahren. Im Linux Kernel müssen bestimmte Optionen aktiviert sein, damit er direkt von der UEFI Firmware gebootet werden kann. Es kann notwendig sein, den Kernel erneut zu kompilieren. Es könnte auch hilfreich sein, den [efibootmgr](/wiki/Efibootmgr "Efibootmgr") Artikel zu lesen.

It is also a good idea to take a look at the [efibootmgr](/wiki/Efibootmgr "Efibootmgr") article for additional information.

**Hinweis**

Zur Wiederholung: efibootmgr ist _keine_ Voraussetzung, um ein UEFI System zu booten. Der Linux-Kernel selbst kann direkt gebootet werden. Zusätzliche Kernel Bootparameter können in den Linux-Kernel einkompiliert werden (es gibt eine Kernel-Konfigurations-Option namens CONFIG\_CMDLINE, mit der Benutzer Bootparameter spezifizieren können). Sogar ein initramfs kann in den Kernels einkompiliert werden.

Diejenigen, die diesen Weg gehen wollen, müssen zuerst die Software installieren:

DATEI **`/etc/portage/package.accept_keywords/installkernel`**

```
sys-kernel/installkernel
sys-boot/uefi-mkconfig
app-emulation/virt-firmware

```

DATEI **`/etc/portage/package.use/installkernel`**

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

Erzeugen Sie das Verzeichnis /boot/efi/boot/ und kopieren Sie den Kernel dorthin. Ändern Sie dabei seinen Namen nach bzImage.efi.

`root #` `mkdir -p /efi/EFI/Gentoo
`

`root #` `cp /boot/vmlinuz-* /efi/EFI/Gentoo/bzImage.efi
`

Install the efibootmgr package:

`root #` `emerge --ask sys-boot/efibootmgr`

Als Nächstes sagen Sie der UEFI Firmware, dass ein Booteintrag mit dem Namen "Gentoo" zu erstellen ist, der den frisch kompilierten EFI stub Kernel bootet:

`root #` `efibootmgr --create --disk /dev/sda --part 1 --label "gentoo" --loader "\EFI\Gentoo\bzImage.efi"`

**Hinweis**

Die Verwendung von `\` als Verzeichnistrenner ist in UEFI Definitionen Pflicht.

Falls ein Ausgangsdateisystem im Arbeitsspeicher (Initial RAM Filesystem = initramfs) verwendet wird, fügen Sie die passenden Bootoptionen hinzu:

`root #` `efibootmgr -c -d /dev/sda -p 2 -L "Gentoo" -l "\efi\boot\bzImage.efi" initrd='\initramfs-genkernel-amd64-6.18.8-gentoo'`

**Tipp**

Additional kernel command line options may be parsed by the firmware to the kernel by specifying them along with the initrd=... option as shown above.

Nach diesen Änderungen wird bei einem Neustart das Systems ein Booteintrag mit der Bezeichnung "gentoo" verfügbar sein.

### Unified Kernel Image

If [installkernel](/wiki/Installkernel "Installkernel") was configured to build and install unified kernel images. The unified kernel image should already be installed to the EFI/Linux directory on the EFI system partition, if this is not the case ensure the directory exists and then run the kernel installation again as described earlier in the handbook.

To add a direct boot entry for the installed unified kernel image:

`root #` `efibootmgr --create --disk /dev/sda --part 1 --label "gentoo" --loader "\EFI\Linux\gentoo-x.y.z.efi"`

## Other Alternatives

For other options that are not covered in the Handbook, see the full list of available [bootloaders](/wiki/Bootloader/de "Bootloader/de").

## Neustart des Systems

Verlassen Sie die chroot-Umgebung und hängen Sie alle gemounteten Partitionen aus. Geben Sie dann den magischen Befehl ein, der den alles entscheidenden Test einleitet - reboot.

`root #` `exit`

`livecd~#` `cd
`

`livecd~#` `umount -l /mnt/gentoo/dev{/shm,/pts,}
`

`livecd~#` `umount -R /mnt/gentoo
`

`livecd~#` `reboot`

Vergessen Sie nicht, das Installations-Medium zu entfernen. Andernfalls könnte erneut das Installations-Medium anstelle des neuen Gentoo Systems gebootet werden.

Nach dem Neustart in die neu installierte Gentoo Umgebung können Sie Ihre Installation mit dem Kapitel [Abschluss der Gentoo Installation](/wiki/Handbook:AMD64/Installation/Finalizing/de "Handbook:AMD64/Installation/Finalizing/de") fertigstellen.

[← Installation der Tools](/wiki/Handbook:AMD64/Installation/Tools/de "Previous part") [Anfang](/wiki/Handbook:AMD64/de "Handbook:AMD64/de") [Abschluss der Installation →](/wiki/Handbook:AMD64/Installation/Finalizing/de "Next part")