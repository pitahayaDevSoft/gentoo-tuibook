# System

Other languages:

- [Deutsch](/wiki/Handbook:X86/Installation/System/de "Handbuch:X86/Installation/System (100% translated)")
- [English](/wiki/Handbook:X86/Installation/System "Handbook:X86/Installation/System (100% translated)")
- [español](/wiki/Handbook:X86/Installation/System/es "Manual de Gentoo: X86/Instalación/Sistema (100% translated)")
- [français](/wiki/Handbook:X86/Installation/System/fr "Handbook:X86/Installation/System/fr (100% translated)")
- [italiano](/wiki/Handbook:X86/Installation/System/it "Handbook:X86/Installation/System (100% translated)")
- [magyar](/wiki/Handbook:X86/Installation/System/hu "Handbook:X86/Installation/System/hu (100% translated)")
- [polski](/wiki/Handbook:X86/Installation/System/pl "Handbook:X86/Installation/System (100% translated)")
- [português do Brasil](/wiki/Handbook:X86/Installation/System/pt-br "Manual:X86/Instalação/Sistema (100% translated)")
- čeština
- [русский](/wiki/Handbook:X86/Installation/System/ru "Handbook:X86/Installation/System (100% translated)")
- [தமிழ்](/wiki/Handbook:X86/Installation/System/ta "கையேடு:X86/நிறுவல்/முறைமை (100% translated)")
- [中文（中国大陆）‎](/wiki/Handbook:X86/Installation/System/zh-cn "手册：X86/安装/配置系统 (100% translated)")
- [日本語](/wiki/Handbook:X86/Installation/System/ja "ハンドブック:X86/インストール/システムの設定 (100% translated)")
- [한국어](/wiki/Handbook:X86/Installation/System/ko "Handbook:X86/Installation/System/ko (100% translated)")

[X86 Handbook](/wiki/Handbook:X86/cs "Handbook:X86/cs")[Instalace](/wiki/Handbook:X86/Full/Installation/cs "Handbook:X86/Full/Installation/cs")[O instalaci](/wiki/Handbook:X86/Installation/About/cs "Handbook:X86/Installation/About/cs")[Výběr média](/wiki/Handbook:X86/Installation/Media/cs "Handbook:X86/Installation/Media/cs")[Konfigurace sítě](/wiki/Handbook:X86/Installation/Networking/cs "Handbook:X86/Installation/Networking/cs")[Příprava disků](/wiki/Handbook:X86/Installation/Disks/cs "Handbook:X86/Installation/Disks/cs")[Instalace stage3](/wiki/Handbook:X86/Installation/Stage/cs "Handbook:X86/Installation/Stage/cs")[Instalace základního systému](/wiki/Handbook:X86/Installation/Base/cs "Handbook:X86/Installation/Base/cs")[Konfigurace jádra](/wiki/Handbook:X86/Installation/Kernel/cs "Handbook:X86/Installation/Kernel/cs")Konfigurace systému[Instalace nástrojů](/wiki/Handbook:X86/Installation/Tools/cs "Handbook:X86/Installation/Tools/cs")[Konfigurace zavaděče](/wiki/Handbook:X86/Installation/Bootloader/cs "Handbook:X86/Installation/Bootloader/cs")[Dokončení](/wiki/Handbook:X86/Installation/Finalizing/cs "Handbook:X86/Installation/Finalizing/cs")[Práce s Gentoo](/wiki/Handbook:X86/Full/Working/cs "Handbook:X86/Full/Working/cs")[Úvod do Portage](/wiki/Handbook:X86/Working/Portage/cs "Handbook:X86/Working/Portage/cs")[Přepínače USE](/wiki/Handbook:X86/Working/USE/cs "Handbook:X86/Working/USE/cs")[Funkce portage](/wiki/Handbook:X86/Working/Features/cs "Handbook:X86/Working/Features/cs")[Systém initskriptů](/wiki/Handbook:X86/Working/Initscripts/cs "Handbook:X86/Working/Initscripts/cs")[Proměnné prostředí](/wiki/Handbook:X86/Working/EnvVar/cs "Handbook:X86/Working/EnvVar/cs")[Práce s Portage](/wiki/Handbook:X86/Full/Portage/cs "Handbook:X86/Full/Portage/cs")[Soubory a adresáře](/wiki/Handbook:X86/Portage/Files/cs "Handbook:X86/Portage/Files/cs")[Proměnné](/wiki/Handbook:X86/Portage/Variables/cs "Handbook:X86/Portage/Variables/cs")[Mísení softwarových větví](/wiki/Handbook:X86/Portage/Branches/cs "Handbook:X86/Portage/Branches/cs")[Doplňkové nástroje](/wiki/Handbook:X86/Portage/Tools/cs "Handbook:X86/Portage/Tools/cs")[Vlastní strom Portage](/wiki/Handbook:X86/Portage/CustomTree/cs "Handbook:X86/Portage/CustomTree/cs")[Pokročilé funkce](/wiki/Handbook:X86/Portage/Advanced/cs "Handbook:X86/Portage/Advanced/cs")[Konfigurace sítě](/wiki/Handbook:X86/Full/Networking/cs "Handbook:X86/Full/Networking/cs")[Začínáme](/wiki/Handbook:X86/Networking/Introduction/cs "Handbook:X86/Networking/Introduction/cs")[Pokročilá konfigurace](/wiki/Handbook:X86/Networking/Advanced/cs "Handbook:X86/Networking/Advanced/cs")[Modulární síťování](/wiki/Handbook:X86/Networking/Modular/cs "Handbook:X86/Networking/Modular/cs")[Bezdrátové sítě](/wiki/Handbook:X86/Networking/Wireless/cs "Handbook:X86/Networking/Wireless/cs")[Přidání funkcí](/wiki/Handbook:X86/Networking/Extending/cs "Handbook:X86/Networking/Extending/cs")[Dynamická správa](/wiki/Handbook:X86/Networking/Dynamic/cs "Handbook:X86/Networking/Dynamic/cs")

## Contents

- [1Informace o systému souborů](#Informace_o_syst.C3.A9mu_soubor.C5.AF)
  - [1.1Popisky souborových systémů a UUID](#Popisky_souborov.C3.BDch_syst.C3.A9m.C5.AF_a_UUID)
  - [1.2Popisky oddílů a UUID](#Popisky_odd.C3.ADl.C5.AF_a_UUID)
  - [1.3O fstab](#O_fstab)
  - [1.4Vytvoření souboru fstab](#Vytvo.C5.99en.C3.AD_souboru_fstab)
    - [1.4.1DOS/Legacy BIOS systems](#DOS.2FLegacy_BIOS_systems)
    - [1.4.2DPS UEFI PARTUUID](#DPS_UEFI_PARTUUID)
- [2Informace o síti](#Informace_o_s.C3.ADti)
  - [2.1Informace o hostiteli a doméně](#Informace_o_hostiteli_a_dom.C3.A9n.C4.9B)
    - [2.1.1Set the hostname (OpenRC or systemd)](#Set_the_hostname_.28OpenRC_or_systemd.29)
    - [2.1.2systemd](#systemd)
  - [2.2Network](#Network)
    - [2.2.1DHCP via dhcpcd (any init system)](#DHCP_via_dhcpcd_.28any_init_system.29)
    - [2.2.2netifrc (OpenRC)](#netifrc_.28OpenRC.29)
  - [2.3Nastavení sítě](#Nastaven.C3.AD_s.C3.ADt.C4.9B)
  - [2.4Automatický start sítě při náběhu](#Automatick.C3.BD_start_s.C3.ADt.C4.9B_p.C5.99i_n.C3.A1b.C4.9Bhu)
  - [2.5Soubor hosts](#Soubor_hosts)
  - [2.6Volitelné: zprovoznění PCMCIA](#Voliteln.C3.A9:_zprovozn.C4.9Bn.C3.AD_PCMCIA)
- [3Systémové informace](#Syst.C3.A9mov.C3.A9_informace)
  - [3.1Heslo root](#Heslo_root)
  - [3.2Konfigurace init a bootování](#Konfigurace_init_a_bootov.C3.A1n.C3.AD)
    - [3.2.1OpenRC](#OpenRC)
    - [3.2.2systemd](#systemd_2)

## Informace o systému souborů

#### Popisky souborových systémů a UUID

MBR (BIOS) a GPT obojí podporují "popisky souborového systému" a "UUID souborového systému". Tyto atributy mohou být definovány v /etc/fstab jako alternativy k použití s příkazem mount při pokusu o nalezení a připojení blokového zařízení. Popisky souborového systému a UUID se rozlišují předponou LABEL a UUID a lze je zobrazit s pomocí příkazu blkid:

`root #` `blkid`

**Upozornění**

Pokud je souborový systém na oddíle vymazán, pak dojde i ke změně nebo vymazání popisku souborového systému a UUID.

S ohledem na unikátnost se čtenářům, kteří používají tabulku oddílu MBR, doporučuje používat UUID namísto popisků k označení připojitelných svazků v /etc/fstab.

**Důležité**

UUIDs of the filesystem on a LVM volume and its LVM snapshots are identical, therefore using UUIDs to mount LVM volumes should be avoided.

#### Popisky oddílů a UUID

Uživatelé, kteří šli cestou GPT, mají k dispozici několik robustnějších metod, jak označit diskové oddíly v /etc/fstab. Popisky oddílů a UUID oddílů lze použít na zařízeních zformátovaných pomocí GPT jako jedinečné identifikátory oddílů blokových zařízení, bez ohledu na to, jaký souborový systém byl pro daný oddíl vybrán. Popisky oddílů a UUID jsou označeny přízvisky PARTLABEL a PARTUUID a mohou být snadno zobrazeny v terminálu spuštěním příkazu blkid:

Output for an **amd64** EFI system using the Discoverable Partition Specification UUIDs may like the following:

`root #` `blkid`

Ačkoli to neplatí vždy pro popisky oddílů, použití UUID k identifikaci oddílu v /etc/fstab poskytuje garanci, že zavaděč nebude zmaten při hledání určitého svazku i v případě, že se v budoucnu změní souborový systém. Použití starších výchozích souborů zařízení /dev/sd\*N k definování oddílů v fstab je v systémech, které jsou často restartovány nebo jsou do nich pravidelně přidávána nebo z nich odstraňována SATA zařízení.

Pojmenování blokových zařízení je závislé na několika faktorech, včetně toho jakým způsobem a v jakém pořadí jsou disky zapojeny do systému. Mohou se také objevit v jiném pořadí v závislosti na tom, která zařízení jádro detekeuje jako první brzy po startu. Vycházeje z tohoto tvrzení, pokud nehodláte neustále přehazovat uspořádání disků, použití výchozích souborů blokových zařízení je jednoduchým a přímočarým přístupem.

### O fstab

Všechny diskové oddíly používané systémem musí být v Linuxu uvedeny v souboru /etc/fstab. Tento soubor obsahuje přípojné body těchto oddílů (kde lze vidět strukturu souborového systému), způsob jak mají být připojeny a s jakými volbami (zda mají být připojeny automaticky či nikoli, zda mohou být připojeny uživateli, atd.).

### Vytvoření souboru fstab

**Poznámka**

If the init system being used is systemd, the partition UUIDs conform to the Discoverable Partition Specification as given in [Preparing the disks](/wiki/Handbook:X86/Installation/Disks/cs "Handbook:X86/Installation/Disks/cs"), and the system uses UEFI, then creating an fstab can be skipped, since systemd auto-mounts partitions that follow the spec.

Soubor /etc/fstab používá syntaxi podobnou tabulce. Každý řádek obsahuje šest polí oddělených prázdným prostorem (mezerou, tabulátorem nebo jejich kombinací). Každé pole má svůj vlastní význam:

1. První pole zobrazuje blokové speciální zařízení nebo vzdálený síťový souborový systém, který má být připojen. K dispozici je několik identifikátorů uzlů zařízení blokových speciálních zařízení, včetně cesty k souboru zařízení, popisky souborového systému, UUID a popisku oddílu s UUID.
2. Druhé pole označuje přípojný bod, kam má být diskový oddíl připojen.
3. Třetí pole zobrazuje systém souborů oddílu.
4. Čtvrté pole zobrazuje volby používané programem mount chce-li oddíl připojit. Každý souborový systém má vlastní volby připojení. Uživatelům se doporučuje přečíst si manuálovou stránku man mount pro kompletní přehled. Více voleb připojení se odděluje čárkami.
5. Páte pole používá utilita dump k určení toho, zda je potřeba oddíl zálohovat. Hodnota může být obvykle ponechána na 0 (nula).
6. Šesté pole používá nástroj fsck k určení toho, v jakém pořadí budou souborové systémy kontrolovány, nebyl-li počítač řádně vypnut. Kořenový souborový systém by měl mít hodnotu nastavenu na 1, zbytek na 2 (nebo 0, pokud není třeba souborový systém kontrolovat).

**Důležité**

Výchozí soubor /etc/fstab, který Gentoo poskytuje není použitelný, jedná se spíše o vzor.

`root #` `nano -w /etc/fstab`

#### DOS/Legacy BIOS systems

Pojďme se podívat na to, jakým způsobem zapsat volby pro oddíl /boot. Toto je pouze příklad a měl by být změněn podle rozhodnutí ohledně dělení, které jste učinili v předchozí části instalace.
V našem příkladu dělení oddílů pro architekturu x86, bývá /boot/ obvykle oddílem /dev/sda1 se souborovým systémem ext2. V průběhu zavádění systému musí být zkontrolován, takže bychom měli zapsat:

FILE **`/etc/fstab`** **Příklad řádku /boot do souboru /etc/fstab**

```
/dev/sda1   /boot     ext2    defaults        0 2

```

Někteří uživatelé nechtějí, aby byl oddíl /boot/ automaticky připojován při startu, z důvodu zvýšení bezpečnosti. Takoví by měli nahradit volbu "defaults" za "noauto". To znamená, že tito uživatelé budou muset manuálně připojit tento oddíl kdykoli jej budou chtít použít.

Přidejte pravidla týkající dříve rozhodnutého schématu rozdělení oddílů a přidejte pravidla pro zařízení jako jsou CD-ROM mechaniky a samozřejmě také ostatní pro oddíly nebo zařízení, jsou-li použity.

Níže je poněkud širší příklad souboru /etc/fstab:

FILE **`/etc/fstab`** **Plný příklad /etc/fstab**

```
/dev/sda1   /boot        ext2    defaults,noatime     0 2
/dev/sda2   none         swap    sw                   0 0
/dev/sda3   /            ext4    noatime              0 1

/dev/cdrom  /mnt/cdrom   auto    noauto,user          0 0

```

/dev/cdrom /mnt/cdrom auto noauto,user 0 0
}}

#### DPS UEFI PARTUUID

Below is an example of an /etc/fstab file for a disk formatted with a GPT disklabel and Discoverable Partition Specification (DPS) UUIDs set for UEFI firmware:

FILE **`/etc/fstab`** **GPT disklabel DPS PARTUUID fstab example**

```
# Adjust any formatting difference and additional partitions created from the "Preparing the disks" step.
# This example shows a GPT disklabel with Discoverable Partition Specification (DSP) UUID set:
PARTUUID=c12a7328-f81f-11d2-ba4b-00a0c93ec93b   /efi        vfat    umask=0077,tz=UTC            0 2
PARTUUID=0657fd6d-a4ab-43c4-84e5-0933c84b4f4f   none            sw                           0 0
PARTUUID=44479540-f297-41b2-9af7-d131d5f0458a   /           xfs    defaults,noatime              0 1

```

Použitím hodnoty `auto` ve třetím poli, příkaz mount donutíme hádat použitý souborový systém. To se doporučuje pro výměnná média, která mohou být vytvořena s jedním z několika souborových systémů. Volba `user` ve čtvrtém poli umožňuje připojení CD běžným uživatelům.

To improve performance, most users would want to add the `noatime` mount option, which results in a faster system since access times are not registered (those are not needed generally anyway). This is also recommended for systems with solid state drives (SSDs). Users may wish to consider `lazytime` instead.

**Tip**

Due to degradation in performance, defining the `discard` mount option in /etc/fstab is not recommended. It is generally better to schedule block discards on a periodic basis using a job scheduler such as cron or a timer (systemd). See [Periodic fstrim jobs](/wiki/SSD#Periodic_fstrim_jobs "SSD") for more information.

Soubor /etc/fstab pečlivě zkontrolujte, uložte, ukončete editor a pokračujte.

## Informace o síti

It is important to note the following sections are provided to help the reader quickly setup their system to partake in a local area network.

For systems running OpenRC, a more detailed reference for network setup is available in the [advanced network configuration](/wiki/Handbook:X86/Networking/Introduction/cs "Handbook:X86/Networking/Introduction/cs") section, which is covered near the end of the handbook. Systems with more specific network needs may need to skip ahead, then return here to continue with the rest of the installation.

For more specific systemd network setup, please review see the [networking portion](/wiki/Systemd#Network "Systemd") of the [systemd](/wiki/Systemd "Systemd") article.

### Informace o hostiteli a doméně

Jedním z rozhodnutí, které musí uživatel učinit je pojmenování svého počítače. To vypadá jednoduše, ale spousta uživatelů má problémy najít odpovídající jméno svého počítače. Abychom vše urychlili, vezměte na vědomí, že toto rozhodnutí není konečné - může být později změněno. V níže uvedených příkladech je jako jméno hostitele použito "tux" v doméně "homenetwork".

#### Set the hostname (OpenRC or systemd)

`root #` `echo tux > /etc/hostname`

#### systemd

To set the system hostname for a system currently _running_ systemd, the hostnamectl utility may be used. During the installation process however, [systemd-firstboot](/wiki/Handbook:X86/Installation/System/cs#Init_and_boot_configuration_systemd "Handbook:X86/Installation/System/cs") command must be used instead (see later on in handbook).

For setting the hostname to "tux", one would run:

`root #` `hostnamectl hostname tux`

View help by running hostnamectl --help or man 1 hostnamectl.

### Network

There are _many_ options available for configuring network interfaces. This section covers a only a few methods. Choose the one which seems best suited to the setup needed.

#### DHCP via dhcpcd (any init system)

Most LAN networks operate a DHCP server. If this is the case, then using the dhcpcd program to obtain an IP address is recommended.

To install:

`root #` `emerge --ask net-misc/dhcpcd`

To enable and then start the service on OpenRC systems:

`root #` `rc-update add dhcpcd default
`

`root #` `rc-service dhcpcd start
`

To enable the service on systemd systems:

`root #` `systemctl enable dhcpcd`

With these steps completed, next time the system boots, dhcpcd should obtain an IP address from the DHCP server. See the [Dhcpcd](/wiki/Dhcpcd "Dhcpcd") article for more details.

#### netifrc (OpenRC)

**Tip**

This is one particular way of setting up the network using [Netifrc](/wiki/Netifrc "Netifrc") on OpenRC. Other methods exist for simpler setups like [Dhcpcd](/wiki/Dhcpcd "Dhcpcd").

### Nastavení sítě

Během instalace Gentoo Linuxu byla síť již nastavena. Nicméně pouze pro instalační CD samotné, ne pro nainstalované prostředí. Nyní dojde ke konfigiraci pro nainstalovaný Gentoo Linux.

**Poznámka**

Podorobnější informace o síťování, včetně pokročilých témat jako jsou bonding, bridging, 802.1Q VLAN nebo bezdrátové sítě je součástí oddílu Konfigurace sítě v Gentoo.

Veškeré informace o síti se zachycují do souboru /etc/conf.d/net. Použítá jasnou, i když možná neintuitivní syntaxi. Ale nebojte, vše je níže vysvětleno. Plně komentovaný příklad, který pokrývá mnoho rozličných konfigurací je k dispozici v souboru /usr/share/doc/netifrc-\*/net.example.bz2.

Nejprve nainstalujte [net-misc/netifrc](https://packages.gentoo.org/packages/net-misc/netifrc):

`root #` `emerge --ask --noreplace net-misc/netifrc`

Ve výchozím stavu se používá DHCP. Aby fungovalo, je třeba instalovat DHCP klienta. To je popsáno v Instalaci nezbytných systémových nástrojů.

Pokud potřebujete nastavit síťové připojení z důvodu zvláštních voleb DHCP nebo proto, že DHCP není vůbec použito, otevřete soubor /etc/conf.d/net:

`root #` `nano -w /etc/conf.d/net`

Nastavte jak config\_eth0, tak i routes\_eth0, aby obsahovaly IP adresu a informace o směrování:

**Poznámka**

To předpokládá, že se síťové rozhraní je označeno eth0. To je však závislé na systému. Doporučuje se předpokládat, že zařízení má stejné pojmenování jako při zavedení instalačního CD, pokud je instalační CD dostatečně nové. Více informací najdete v části Názvy síťových rozhraní.

FILE **`/etc/conf.d/net`** **Definice statické IP**

```
config_eth0="192.168.0.2 netmask 255.255.255.0 brd 192.168.0.255"
routes_eth0="default via 192.168.0.1"

```

Pro použití DHCP, definujte config\_eth0:

FILE **`/etc/conf.d/net`** **Definice DHCP**

```
config_eth0="dhcp"

```

Laskavě si přečtěte /usr/share/doc/netifrc-\*/net.example.bz2, kde jsou vypsány všechny dostupné volby. Stejně tak si určitě přečtěte man stránku DHCP klienta, pokud je nutné nastavit určité volby DHCP.

Pokud má systém několik síťových rozhraní, může shora uvedené kroky opakovat pro config\_eth1, config\_eth2 atd.

Nyní konfiguraci uložte, opusťte editor a pokračujte.

### Automatický start sítě při náběhu

Aby byla síťová rozhraní při náběhu systému aktivována, je potřeba je přidat do výchozí úrovně běhu.

`root #` `cd /etc/init.d
`

`root #` `ln -s net.lo net.eth0
`

`root #` `rc-update add net.eth0 default`

Pokud má systém více síťových rozhraní, pak musí být vytvořeny odpovídající soubory net.\*, stejně jako jsme to provedli s net.eth0.

Pokud po naběhnutí systému zjistíme, že předpoklad o názvu síťového rozhraní (nyní používáme `eth0`) byl špatný, proveďte následující kroky k nápravě:

1. Aktualizujte v souboru /etc/conf.d/net správné názvy rozhraní (jako je `enp3s0` namísto `eth0`).
2. Vytvořte nový symbolický odkaz (jako je /etc/init.d/net.enp3s0).
3. Odstraňte starý symbolický odkaz (rm /etc/init.d/net.eth0).
4. Přidejte nový do výchozí úrovně běhu.
5. Odstraňte předchozí použitím rc-update del net.eth0 default.

### Soubor hosts

Dále dejte počítači informaci o síťovém prostředí. Ty se vkládají do /etc/hosts a pomáhají přeložit názvy hostitelů na IP adresy, které nejsou přeloženy jmenným serverem.

`root #` `nano -w /etc/hosts`

FILE **`/etc/hosts`** **Vyplnění informací o síti**

```
# Tímto se definuje používaný systém a musí to být nastaveno
127.0.0.1     tux.homenetwork tux localhost

# Volitelné definice dalších systémů v síti
192.168.0.5   jenny.homenetwork jenny
192.168.0.6   benny.homenetwork benny

```

1. Optional definition of extra systems on the network

192.168.0.5 jenny.homenetwork jenny
192.168.0.6 benny.homenetwork benny
}}

Uložte, opusťte editor a pokračujte.

### Volitelné: zprovoznění PCMCIA

Uživatelé PCMCIA nechť nainstalují balíček [sys-apps/pcmciautils](https://packages.gentoo.org/packages/sys-apps/pcmciautils).

`root #` `emerge --ask sys-apps/pcmciautils`

## Systémové informace

### Heslo root

Pomocí příkazu passwd nastavte root heslo.

`root #` `passwd`

Účet root je v Linuxu všemocný, tudíž si vyberte silné heslo. Později vytvoříme další účty běžných uživatelů pro každodenní používání.

### Konfigurace init a bootování

#### OpenRC

Gentoo (alespoň pokud používáme OpenRC) používá /etc/rc.conf k nastavení služeb, spuštění a vypnutí systému. Otevřete /etc/rc.conf a užijte si čtení komentářů v tomto souboru. Projděte si nastavení a proveďte změny, kde je to potřeba.

`root #` `nano -w /etc/rc.conf`

Dále otevřete soubor /etc/conf.d/keymaps ke změně nastavení klávesnice. Upravte jej k nastavení správné klávesnice.

`root #` `nano -w /etc/conf.d/keymaps`

Zvláštní pozornost věnujte proměnné keymap. Pokud vyberete špatné rozložení, budou výsledkem psaní na klávesnici různé podivnosti.

pro nastavení hodin nakonec upravte soubor /etc/conf.d/hwclock. Změňte jej dle vašich preferencí.

`root #` `nano -w /etc/conf.d/hwclock`

Pokud hardwarové hodiny nepoužívají UTC, je nezbytné v souboru nastavit `clock="local"`. V opačném případě může systém vykazovat posouvání hodin.

#### systemd

First, it is recommended to run systemd-machine-id-setup and then systemd-firstboot which will prepare various components of the system are set correctly for the first boot into the new systemd environment. The passing the following options will include a prompt for the user to set a locale, timezone, hostname, root password, and root shell values. It will also assign a random machine ID to the installation:

`root #` `systemd-machine-id-setup`

`root #` `systemd-firstboot --prompt`

Next users should run systemctl to reset all installed unit files to the preset policy values:

**Poznámka**

It is possible that a failure warning will print out after running the following command. This is normal, as Gentoo's LiveCD uses OpenRC.

`root #` `systemctl preset-all --preset-mode=enable-only`

It's possible to run the full preset changes but this may reset any services which were already configured during the process:

`root #` `systemctl preset-all`

These two steps will help ensure a smooth transition from the live environment to the installation's first boot.

[← Konfigurace jádra](/wiki/Handbook:X86/Installation/Kernel/cs "Previous part") [Home](/wiki/Handbook:X86/cs "Handbook:X86/cs") [Instalace nástrojů →](/wiki/Handbook:X86/Installation/Tools/cs "Next part")