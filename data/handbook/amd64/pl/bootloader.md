# Bootloader

Other languages:

- [Deutsch](/wiki/Handbook:AMD64/Installation/Bootloader/de "Handbook:AMD64/Installation/Bootloader/de (100% translated)")
- [English](/wiki/Handbook:AMD64/Installation/Bootloader "Handbook:AMD64/Installation/Bootloader (100% translated)")
- [español](/wiki/Handbook:AMD64/Installation/Bootloader/es "Handbook:AMD64/Installation/Bootloader/es (100% translated)")
- [français](/wiki/Handbook:AMD64/Installation/Bootloader/fr "Handbook:AMD64/Installation/Bootloader/fr (100% translated)")
- [italiano](/wiki/Handbook:AMD64/Installation/Bootloader/it "Handbook:AMD64/Installation/Bootloader/it (100% translated)")
- [magyar](/wiki/Handbook:AMD64/Installation/Bootloader/hu "Handbook:AMD64/Installation/Bootloader/hu (100% translated)")
- polski
- [português do Brasil](/wiki/Handbook:AMD64/Installation/Bootloader/pt-br "Handbook:AMD64/Installation/Bootloader/pt-br (100% translated)")
- [čeština](/wiki/Handbook:AMD64/Installation/Bootloader/cs "Handbook:AMD64/Installation/Bootloader/cs (100% translated)")
- [русский](/wiki/Handbook:AMD64/Installation/Bootloader/ru "Handbook:AMD64/Installation/Bootloader/ru (100% translated)")
- [தமிழ்](/wiki/Handbook:AMD64/Installation/Bootloader/ta "Handbook:AMD64/Installation/Bootloader/ta (100% translated)")
- [中文（中国大陆）‎](/wiki/Handbook:AMD64/Installation/Bootloader/zh-cn "Handbook:AMD64/Installation/Bootloader/zh-cn (100% translated)")
- [日本語](/wiki/Handbook:AMD64/Installation/Bootloader/ja "Handbook:AMD64/Installation/Bootloader/ja (100% translated)")
- [한국어](/wiki/Handbook:AMD64/Installation/Bootloader/ko "Handbook:AMD64/Installation/Bootloader/ko (100% translated)")

[AMD64 Handbook](/wiki/Handbook:AMD64/pl "Handbook:AMD64/pl")[Instalacja](/wiki/Handbook:AMD64/Full/Installation/pl "Handbook:AMD64/Full/Installation/pl")[O instalacji](/wiki/Handbook:AMD64/Installation/About/pl "Handbook:AMD64/Installation/About/pl")[Wybór medium instalacyjnego](/wiki/Handbook:AMD64/Installation/Media/pl "Handbook:AMD64/Installation/Media/pl")[Konfiguracja sieci](/wiki/Handbook:AMD64/Installation/Networking/pl "Handbook:AMD64/Installation/Networking/pl")[Przygotowanie dysków](/wiki/Handbook:AMD64/Installation/Disks/pl "Handbook:AMD64/Installation/Disks/pl")[Instalacja etapu 3](/wiki/Handbook:AMD64/Installation/Stage/pl "Handbook:AMD64/Installation/Stage/pl")[Instalacja systemu podstawowego](/wiki/Handbook:AMD64/Installation/Base/pl "Handbook:AMD64/Installation/Base/pl")[Konfiguracja jądra](/wiki/Handbook:AMD64/Installation/Kernel/pl "Handbook:AMD64/Installation/Kernel/pl")[Konfiguracja systemu](/wiki/Handbook:AMD64/Installation/System/pl "Handbook:AMD64/Installation/System/pl")[Instalacja narzędzi](/wiki/Handbook:AMD64/Installation/Tools/pl "Handbook:AMD64/Installation/Tools/pl")Instalacja systemu rozruchowego[Finalizacja](/wiki/Handbook:AMD64/Installation/Finalizing/pl "Handbook:AMD64/Installation/Finalizing/pl")[Praca z Gentoo](/wiki/Handbook:AMD64/Full/Working/pl "Handbook:AMD64/Full/Working/pl")[Wstęp do Portage](/wiki/Handbook:AMD64/Working/Portage/pl "Handbook:AMD64/Working/Portage/pl")[Flagi USE](/wiki/Handbook:AMD64/Working/USE/pl "Handbook:AMD64/Working/USE/pl")[Funkcje portage](/wiki/Handbook:AMD64/Working/Features/pl "Handbook:AMD64/Working/Features/pl")[System initscript](/wiki/Handbook:AMD64/Working/Initscripts/pl "Handbook:AMD64/Working/Initscripts/pl")[Zmienne środowiskowe](/wiki/Handbook:AMD64/Working/EnvVar/pl "Handbook:AMD64/Working/EnvVar/pl")[Praca z Portage](/wiki/Handbook:AMD64/Full/Portage/pl "Handbook:AMD64/Full/Portage/pl")[Pliki i katalogi](/wiki/Handbook:AMD64/Portage/Files/pl "Handbook:AMD64/Portage/Files/pl")[Zmienne](/wiki/Handbook:AMD64/Portage/Variables/pl "Handbook:AMD64/Portage/Variables/pl")[Mieszanie działów oprogramowania](/wiki/Handbook:AMD64/Portage/Branches/pl "Handbook:AMD64/Portage/Branches/pl")[Dodatkowe narzędzia](/wiki/Handbook:AMD64/Portage/Tools/pl "Handbook:AMD64/Portage/Tools/pl")[Custom package repository](/wiki/Handbook:AMD64/Portage/CustomTree/pl "Handbook:AMD64/Portage/CustomTree/pl")[Funkcje zaawansowane](/wiki/Handbook:AMD64/Portage/Advanced/pl "Handbook:AMD64/Portage/Advanced/pl")[Konfiguracja sieci](/wiki/Handbook:AMD64/Full/Networking/pl "Handbook:AMD64/Full/Networking/pl")[Zaczynamy](/wiki/Handbook:AMD64/Networking/Introduction/pl "Handbook:AMD64/Networking/Introduction/pl")[Zaawansowana konfiguracja](/wiki/Handbook:AMD64/Networking/Advanced/pl "Handbook:AMD64/Networking/Advanced/pl")[Sieć modularna](/wiki/Handbook:AMD64/Networking/Modular/pl "Handbook:AMD64/Networking/Modular/pl")[Sieć bezprzewodowa](/wiki/Handbook:AMD64/Networking/Wireless/pl "Handbook:AMD64/Networking/Wireless/pl")[Dodawanie funkcjonalności](/wiki/Handbook:AMD64/Networking/Extending/pl "Handbook:AMD64/Networking/Extending/pl")[Dynamiczne zarządzanie](/wiki/Handbook:AMD64/Networking/Dynamic/pl "Handbook:AMD64/Networking/Dynamic/pl")

## Contents

- [1Wybór programu rozruchowego](#Wyb.C3.B3r_programu_rozruchowego)
- [2Domyślnie: GRUB2](#Domy.C5.9Blnie:_GRUB2)
  - [2.1Emerge](#Emerge)
  - [2.2Instalacja](#Instalacja)
    - [2.2.1Optional: Secure Boot](#Optional:_Secure_Boot)
    - [2.2.2Debugging GRUB](#Debugging_GRUB)
  - [2.3Konfiguracja](#Konfiguracja)
- [3Alternative 1: systemd-boot](#Alternative_1:_systemd-boot)
  - [3.1Emerge](#Emerge_2)
  - [3.2Installation](#Installation)
  - [3.3Optional: Secure Boot](#Optional:_Secure_Boot_2)
- [4Alternative 2: EFI Stub](#Alternative_2:_EFI_Stub)
  - [4.1Unified Kernel Image](#Unified_Kernel_Image)
- [5Other Alternatives](#Other_Alternatives)
- [6Ponowne uruchomienie systemu](#Ponowne_uruchomienie_systemu)

## Wybór programu rozruchowego

Z skonfigurowanym jadrem Linuksa, zainstalowanymi narzędziami systemowymi i z zmodyfikowanymi plikami konfiguracyjnymi, przyszedł czas na instalację ostatniej ważnej części systemu Linux: systemu rozruchowego.

System rozruchowy jest odpowiedzialny za uruchomienie jądra systemu Linux podczas startu komputera - bez niego, komputer nie wiedziałby co należy zrobić po wciśnięciu przycisku zasilania.

Dla **amd64**, udokumentowaliśmy jak skonfigurować zarówno [GRUB2](#Default:_Using_GRUB2) jak i [LILO](#Alternative:_Using_LILO) dla systemów opartych o BIOS, oraz [GRUB2](#Default:_Using_GRUB2) i [efibootmgr](#Alternative:_Using_efibootmgr) dla systemów UEFI.

W tej sekcji Podręcznika dokonano podziału pomiędzy "emerge" systemu rozruchowego, a "instalacją" systemu rozruchowego na dysk systemowy. Pojęcie "emerge" będzie używane do zapytania [Portage](/wiki/Portage "Portage") o utworzenie pakietu oprogramowania dostępnego dla systemu. Pojęcie "instalacja" będzie oznaczało kopiowanie plików systemu rozruchowego lub fizyczną modyfikację stosownych sekcji dysku systemowego, w celu "uaktywnienia i przygotowania do działania" systemu rozruchowego na następne uruchomienie systemu operacyjnego.

## Domyślnie: GRUB2

Domyślnie, większość systemów Gentoo korzysta teraz z GRUB2 (znajdującego się w pakiecie [sys-boot/grub](https://packages.gentoo.org/packages/sys-boot/grub)), który jest sukcesorem przestarzałego GRUB. Bez dodatkowej konfiguracji, GRUB2 wspiera starsze wersje BIOS. Z drobną konfiguracją, niezbędną przed kompilacją pakietu, GRUB2 może wspierać więcej niż tuzin dodatkowych platform sprzętowych. Więcej informacji znajdziesz w sekcji [Wymagania wstępne](/wiki/GRUB2#Prerequisites "GRUB2") artykułu [GRUB2](/wiki/GRUB2 "GRUB2").

### Emerge

Podczas korzystania z starszego BIOS, system wspiera jedynie tablicę partycji MBR. Nie ma potrzeby dodatkowej konfiguracji w celu zainstalowania systemu rozruchowego GRUB:

`root #` `emerge --ask --verbose sys-boot/grub:2`

Informacja dla użytkowników UEFI: uruchomienie powyższej komendy wyświetli aktualną zawartość GRUB\_PLATFORMS przed przystąpieniem do kompilacji. Użytkownicy systemów UEFI muszą zapewnić włączenie opcji `GRUB_PLATFORMS="efi-64"` (jak ma to miejsce domyślnie). Jeżeli opcja nie jest włączona, `GRUB_PLATFORMS="efi-64"` musi zostać dodane do pliku /etc/portage/make.conf _przed_ przystąpieniem do kompilacji GRUB2, tak aby pakiet był skompilowany z funkcjonalnością EFI:

`root #` `echo 'GRUB_PLATFORMS="efi-64"' >> /etc/portage/make.conf`

`root #` `emerge --ask sys-boot/grub:2`

Jeżeli GRUB2 został jakkolwiek skompilowany bez włączenia `GRUB_PLATFORMS="efi-64"`, linia ta (jak wyżej wskazano) nadal może zostać dodana do pliku make.conf, a zależności [zbioru pakietów world](/wiki/World_set_(Portage) "World set (Portage)") ponownie obliczone poprzez użycie opcji `--update --newuse` dla polecenia emerge:

`root #` `emerge --ask --update --newuse --verbose sys-boot/grub:2`

Oprogramowanie GRUB2 zostało właśnie przyłączone do systemu, ale nie zostało jeszcze zainstalowane.

### Instalacja

Następnie, zainstaluj niezbędne pliki GRUB2 do ścieżki /boot/grub/ poprzez komendę grub-install. Zakładając, że pierwszy dysk (ten z którego system się uruchamia) to /dev/sda, uruchom jedno z poniższych poleceń:

- Podczas korzystania z BIOS:

`root #` `grub-install /dev/sda`

For DOS/Legacy BIOS systems:

`root #` `grub-install /dev/sda`

Podczas korzystania z UEFI:

**Ważne**

Upewnij się, że partycja systemowa EFI została zamontowana _przed_ uruchomieniem grub-install. Istnieje ryzyko, że grub-install zainstaluje plik GRUB EFI (grubx64.efi) w złej ścieżce **bez** podania _jakiegokolwiek_ wskazania, że użyto złej ścieżki

For UEFI systems:

`root #` `grub-install --target=x86_64-efi --efi-directory=/boot`

Upon successful installation, the output should match the output of the previous command. If the output does not match exactly, then proceed to [Debugging GRUB](/wiki/Handbook:AMD64/Blocks/Bootloader/pl#Debugging_GRUB "Handbook:AMD64/Blocks/Bootloader/pl"), otherwise jump to the [Configure step](/wiki/Handbook:AMD64/Blocks/Bootloader/pl#GRUB_Configure "Handbook:AMD64/Blocks/Bootloader/pl").

##### Optional: Secure Boot

To successfully boot with secure boot enabled the signing certificate must either be accepted by the [UEFI](/wiki/UEFI "UEFI") firmware, or [shim](/wiki/Shim "Shim") must be used as a pre-loader. Shim is pre-signed with the third-party Microsoft Certificate, accepted by default by most UEFI motherboards.

How to configure the UEFI firmware to accept custom keys depends on the firmware vendor, which is beyond the scope of the handbook. Below is shown how to setup shim instead. Here it is assumed that the user has already followed the instructions in the previous sections to generate a signing key and to configure portage to use it. If this is not the case please return first to the [Kernel installation](/wiki/Handbook:AMD64/Installation/Kernel/pl "Handbook:AMD64/Installation/Kernel/pl") section.

The package [sys-boot/grub](https://packages.gentoo.org/packages/sys-boot/grub) installs a prebuilt and signed stand-alone EFI executable if the [secureboot](https://packages.gentoo.org/useflags/secureboot) [USE flag/pl](/wiki/USE_flag/pl "USE flag/pl") USE flag is enabled. Install the required packages and copy the stand-alone grub, Shim, and the MokManager to the same directory on the EFI System Partition. For example:

`root #` `emerge sys-boot/grub sys-boot/shim sys-boot/mokutil sys-boot/efibootmgr`

`root #` `cp /usr/share/shim/BOOTX64.EFI /efi/EFI/Gentoo/shimx64.efi
`

`root #` `cp /usr/share/shim/mmx64.efi /efi/EFI/Gentoo/mmx64.efi
`

`root #` `cp /usr/lib/grub/grub-x86_64.efi.signed /efi/EFI/Gentoo/grubx64.efi
`

Next register the signing key in shims MOKlist, this requires keys in the DER format, whereas sbsign and the kernel build system expect keys in the PEM format. In the previous sections of the handbook an example was shown to generate such a signing PEM key, this key must now be converted to the DER format:

`root #` `openssl x509 -in /path/to/kernel_key.pem -inform PEM -out /path/to/kernel_key.der -outform DER`

**Informacja**

The path used here must be the path to the pem file containing the certificate belonging to the generated key. In this example both key and certificate are in the same pem file.

Then the converted certificate can be imported into Shims MOKlist, this command will ask to set some password for the import request:

`root #` `mokutil --import /path/to/kernel_key.der`

**Informacja**

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

**Informacja**

The import process will not be completed until the system is rebooted. After completing all steps in the handbook, restart the system and Shim will load, it will find the import request registered by mokutil. The MokManager application will start and ask for the password that was set when creating the import request. Follow the on-screen instructions to complete the import of the certificate, then reboot the system into the UEFI menu and enable the Secure Boot setting.

##### Debugging GRUB

When debugging GRUB, there are a couple of quick fixes that may result in a bootable installation without having to reboot to a new live image environment.

In the event that "EFI variables are not supported on this system" is displayed somewhere in the output, it is likely the live image was not booted in EFI mode and is presently in Legacy BIOS boot mode. The solution is to try the [removable GRUB step](/wiki/Handbook:AMD64/Blocks/Bootloader/pl#GRUB_Install_EFI_systems_removable "Handbook:AMD64/Blocks/Bootloader/pl") mentioned below. This will overwrite the executable EFI file located at /EFI/BOOT/BOOTX64.EFI. Upon rebooting in EFI mode, the motherboard firmware may execute this default boot entry and execute GRUB.

**Ważne**

Jeśli grub-install zwraca błąd podobny do `Could not prepare Boot variable: Read-only file system`, może być konieczne ponowne zamontowanie efivars w trybie do odczytu i zapisu, aby odnieść sukces:

`root #` `mount -o remount,rw /sys/firmware/efi/efivars`

`root #` `mount -o remount,rw,nosuid,nodev,noexec --types efivarfs efivarfs /sys/firmware/efi/efivars`

This is caused by certain non-official Gentoo environments not mounting the special EFI filesystem by default. If the previous command does not run, then reboot using an official Gentoo live image environment in EFI mode.

Some motherboard manufacturers with poor UEFI implementations seem to _only_ support the /EFI/BOOT directory location for the .EFI file in the EFI System Partition (ESP). The GRUB installer can create the .EFI file in this location automatically by appending the `--removable` option to the install command. Ensure the ESP has been mounted before running the following command; presuming it is mounted at /efi (as defined earlier), run:

`root #` `grub-install --target=x86_64-efi --efi-directory=/boot --removable`

This creates the 'default' directory defined by the UEFI specification, and then creates a file with the default name: BOOTX64.EFI.

### Konfiguracja

Next, generate the GRUB configuration based on the user configuration specified in the /etc/default/grub file and /etc/grub.d scripts. In most cases, no configuration is needed by users as GRUB will automatically detect which kernel to boot (the highest one available in /boot/) and what the root file system is. It is also possible to append kernel parameters in /etc/default/grub using the GRUB\_CMDLINE\_LINUX variable.

Aby wygenerować końcową konfigurację GRUB2, uruchom polecenie grub-mkconfig:

`root #` `grub-mkconfig -o /boot/grub/grub.cfg`

```
Generating grub.cfg ...
Found linux image: /boot/vmlinuz-6.18.8-gentoo
Found initrd image: /boot/initramfs-genkernel-amd64-6.18.8-gentoo
done

```

The output of the command must mention that at least one Linux image is found, as those are needed to boot the system. If an initramfs is used or genkernel was used to build the kernel, the correct initrd image should be detected as well. If this is not the case, go to /boot/ and check the contents using the ls command. If the files are indeed missing, go back to the kernel configuration and installation instructions.

**Wskazówka**

The os-prober utility can be used in conjunction with GRUB to detect other operating systems from attached drives. Windows 7, 8.1, 10, and other distributions of Linux are detectable. Those desiring dual boot systems should emerge the [sys-boot/os-prober](https://packages.gentoo.org/packages/sys-boot/os-prober) package then re-run the grub-mkconfig command (as seen above). If detection problems are encountered be sure to read the [GRUB](/wiki/GRUB "GRUB") article in its entirety _before_ asking the Gentoo community for support.

## Alternative 1: systemd-boot

Another option is [systemd-boot](/wiki/Systemd/systemd-boot "Systemd/systemd-boot"), which works on both OpenRC and systemd machines. It is a thin chainloader and works well with secure boot.

### Emerge

To install systemd-boot, enable the [boot](https://packages.gentoo.org/useflags/boot) [USE flag/pl](/wiki/USE_flag/pl "USE flag/pl") USE flag and re-install [sys-apps/systemd](https://packages.gentoo.org/packages/sys-apps/systemd) (for systemd systems) or [sys-apps/systemd-utils](https://packages.gentoo.org/packages/sys-apps/systemd-utils) (for OpenRC systems):

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

**Ważne**

Make sure the EFI system partition has been mounted _before_ running bootctl install.

When using this bootloader, before rebooting, verify that a new bootable entry exists using:

`root #` `bootctl list`

**Uwaga**

The kernel command line for new systemd-boot entries is read from /etc/kernel/cmdline or /usr/lib/kernel/cmdline. If neither file is present, then the kernel command line of the currently booted kernel is re-used (/proc/cmdline). On new installs it might therefore happen that the kernel command line of the live CD is accidentally used to boot the new kernel. The kernel command line for registered entries can be checked with:

`root #` `bootctl list`

If this does not show the desired kernel command line then create /etc/kernel/cmdline containing the correct kernel command line and re-install the kernel.

If no new entry exists, ensure the [sys-kernel/installkernel](https://packages.gentoo.org/packages/sys-kernel/installkernel) package has been installed with the [systemd](https://packages.gentoo.org/useflags/systemd) [USE flag/pl](/wiki/USE_flag/pl "USE flag/pl") and [systemd-boot](https://packages.gentoo.org/useflags/systemd-boot) [USE flag/pl](/wiki/USE_flag/pl "USE flag/pl") USE flags enabled, and re-run the kernel installation.

For the distribution kernels:

`root #` `emerge --ask --config sys-kernel/gentoo-kernel`

For a manually configured and compiled kernel:

`root #` `make install`

**Ważne**

When installing kernels for systemd-boot, no root= kernel command line argument is added by default. On systemd systems that are using an initramfs users may rely instead on [systemd-gpt-auto-generator](/wiki/Systemd#Automatic_mounting_of_partitions_at_boot.2Fpl "Systemd") to automatically find the root partition at boot. Otherwise users should manually specify the location of the root partition by setting root= in /etc/kernel/cmdline as well as any other kernel command line arguments that should be used. And then reinstalling the kernel as described above.

### Optional: Secure Boot

When the [secureboot](https://packages.gentoo.org/useflags/secureboot) [USE flag/pl](/wiki/USE_flag/pl "USE flag/pl") USE flag is enabled, the systemd-boot EFI executable will be signed by portage automatically. Furthermore, bootctl install will automatically install the signed version.

To successfully boot with secure boot enabled the used certificate must either be accepted by the [UEFI](/wiki/UEFI "UEFI") firmware, or [shim](/wiki/Shim "Shim") must be used as a pre-loader. Shim is pre-signed with the third-party Microsoft Certificate, accepted by default by most UEFI motherboards.

How to configure the UEFI firmware to accept custom keys depends on the firmware vendor, which is beyond the scope of the handbook. Below is shown how to setup shim instead. Here it is assumed that the user has already followed the instructions in the previous sections to generate a signing key and to configure portage to use it. If this is not the case please return first to the [Kernel installation](/wiki/Handbook:AMD64/Installation/Kernel/pl "Handbook:AMD64/Installation/Kernel/pl") section.

`root #` `emerge --ask sys-boot/shim sys-boot/mokutil sys-boot/efibootmgr`

`root #` `bootctl install --no-variables`

`root #` `cp /usr/share/shim/BOOTX64.EFI /efi/EFI/systemd/shimx64.efi`

`root #` `cp /usr/share/shim/mmx64.efi /efi/EFI/systemd/mmx64.efi`

Shims MOKlist requires keys in the DER format, whereas sbsign and the kernel build system expect keys in the PEM format. In the previous sections of the handbook an example was shown to generate such a signing PEM key, this key must now be converted to the DER format:

`root #` `openssl x509 -in /path/to/kernel_key.pem -inform PEM -out /path/to/kernel_key.der -outform DER`

**Informacja**

The path used here must be the path to the pem file containing the certificate belonging to the generated key. In this example both key and certificate are in the same pem file.

Then the converted certificate can be imported into Shims MOKlist:

`root #` `mokutil --import /path/to/kernel_key.der`

**Informacja**

When the currently booted kernel already trusts the certificate being imported, the message "Already in kernel trusted keyring." will be returned here. If this happens, re-run the above command with the argument --ignore-keyring added.

And finally we register Shim with the UEFI firmware. In the following command, `boot-disk` and `boot-partition-id` must be replaced with the disk and partition identifier of the EFI system partition:

`root #` `efibootmgr --create --disk /dev/boot-disk --part boot-partition-id --loader '\EFI\systemd\shimx64.efi'  --label 'Systemd-boot via Shim' --unicode '\EFI\systemd\systemd-bootx64.efi'`

**Informacja**

The import process will not be completed until the system is rebooted. After completing all steps in the handbook, restart the system and Shim will load, it will find the import request registered by mokutil. The MokManager application will start and ask for the password that was set when creating the import request. Follow the on-screen instructions to complete the import of the certificate, then reboot the system into the UEFI menu and enable the Secure Boot setting.

## Alternative 2: EFI Stub

Computer systems with UEFI-based firmware technically do not need secondary bootloaders (e.g. GRUB) in order to boot kernels. Secondary bootloaders exist to _extend_ the functionality of UEFI firmware during the boot process. Using GRUB (see the prior section) is typically easier and more robust because it offers a more flexible approach for quickly modifying kernel parameters at boot time.

System administrators who desire to take a minimalist, although more rigid, approach to booting the system can avoid secondary bootloaders and boot the Linux kernel as an [EFI stub](/wiki/EFI_stub "EFI stub").

The [sys-boot/efibootmgr](https://packages.gentoo.org/packages/sys-boot/efibootmgr) application is a tool to used interact with UEFI firmware - the system's primary bootloader. Normally this looks like adding or removing boot entries to the firmware's list of bootable entries. It can also update firmware settings so that the Linux kernels that were previously added as bootable entries can be executed with additional options. These interactions are performed through special data structures called EFI variables (hence the need for kernel support of EFI vars).

Ensure the [EFI stub](/wiki/EFI_stub "EFI stub") kernel article has been reviewed _before_ continuing. The kernel must have specific options enabled to be directly bootable by the UEFI firmware. It may be necessary to recompile the kernel in order to build-in this support.

It is also a good idea to take a look at the [efibootmgr](/wiki/Efibootmgr "Efibootmgr") article for additional information.

**Informacja**

To reiterate, efibootmgr is _not_ a requirement to boot an UEFI system; it is merely necessary to add an entry for an EFI-stub kernel into the UEFI firmware. When built appropriately with EFI stub support, the Linux kernel itself can be booted _directly_. Additional kernel command-line options can be built-in to the Linux kernel (there is a kernel configuration option called CONFIG\_CMDLINE. Similarly, support for initramfs can be 'built-in' to the kernel as well.

To boot the kernel directly from the firmware, the kernel image must be present on the [EFI System Partition](/wiki/EFI_System_Partition "EFI System Partition"). This may be accomplished by enabling the [efistub](https://packages.gentoo.org/useflags/efistub) [USE flag/pl](/wiki/USE_flag/pl "USE flag/pl") USE flag on [sys-kernel/installkernel](https://packages.gentoo.org/packages/sys-kernel/installkernel). Given that EFI Stub booting is not guaranteed to work on every UEFI system, the USE flag is stable masked, testing keywords must be accepted for installkernel to use this feature.

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

Alternatively, when [installkernel](/wiki/Installkernel "Installkernel") is not used, the kernel may be copied manually to the EFI System Partition, calling it bzImage.efi:

`root #` `mkdir -p /efi/EFI/Gentoo
`

`root #` `cp /boot/vmlinuz-* /efi/EFI/Gentoo/bzImage.efi
`

Install the efibootmgr package:

`root #` `emerge --ask sys-boot/efibootmgr`

Create boot entry called "Gentoo" for the freshly compiled EFI stub kernel within the UEFI firmware:

`root #` `efibootmgr --create --disk /dev/sda --part 1 --label "gentoo" --loader "\EFI\Gentoo\bzImage.efi"`

**Informacja**

The use of a backslash ( `\`) as directory path separator is mandatory when using UEFI definitions.

Jeśli używany jest inicjalny system plików RAM (initramfs), dodaj do niego odpowiednią opcję rozruchu:

`root #` `efibootmgr -c -d /dev/sda -p 2 -L "Gentoo" -l "\efi\boot\bzImage.efi" initrd='\initramfs-genkernel-amd64-6.18.8-gentoo'`

**Wskazówka**

Additional kernel command line options may be parsed by the firmware to the kernel by specifying them along with the initrd=... option as shown above.

With these changes done, when the system reboots, a boot entry called "gentoo" will be available.

### Unified Kernel Image

If [installkernel](/wiki/Installkernel "Installkernel") was configured to build and install unified kernel images. The unified kernel image should already be installed to the EFI/Linux directory on the EFI system partition, if this is not the case ensure the directory exists and then run the kernel installation again as described earlier in the handbook.

To add a direct boot entry for the installed unified kernel image:

`root #` `efibootmgr --create --disk /dev/sda --part 1 --label "gentoo" --loader "\EFI\Linux\gentoo-x.y.z.efi"`

## Other Alternatives

For other options that are not covered in the Handbook, see the full list of available [bootloaders](/wiki/Bootloader "Bootloader").

## Ponowne uruchomienie systemu

Wyjdź z środowiska chroot i odmontuj wszystkie zamontowane partycje. Następnie wpisz magiczne polecenie, które inicjuje ostateczny, prawdziwy test: reboot.

`root #` `exit`

`cdimage ~#` `cd
`

`cdimage ~#` `umount -l /mnt/gentoo/dev{/shm,/pts,}
`

`cdimage ~#` `umount -R /mnt/gentoo
`

`cdimage ~#` `reboot`

Nie zapomnij usunąć medium instalacyjnego, w przeciwnym razie zamiast nowego systemu Gentoo może zostać ponownie uruchomione medium instalacyjne.

Po uruchomieniu świeżo zainstalowanego środowiska Gentoo, kontynuuj [Finalizowanie instalacji Gentoo](/wiki/Handbook:AMD64/Installation/Finalizing "Handbook:AMD64/Installation/Finalizing").

[← Installing tools](/wiki/Handbook:AMD64/Installation/Tools/pl "Previous part") [Home](/wiki/Handbook:AMD64 "Handbook:AMD64") [Finalizing →](/wiki/Handbook:AMD64/Installation/Finalizing/pl "Next part")