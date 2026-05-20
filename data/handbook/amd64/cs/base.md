# Base

Other languages:

- [Deutsch](/wiki/Handbook:AMD64/Installation/Base/de "Handbook:AMD64/Installation/Base/de (100% translated)")
- [English](/wiki/Handbook:AMD64/Installation/Base "Handbook:AMD64/Installation/Base (100% translated)")
- [español](/wiki/Handbook:AMD64/Installation/Base/es "Handbook:AMD64/Installation/Base/es (100% translated)")
- [français](/wiki/Handbook:AMD64/Installation/Base/fr "Handbook:AMD64/Installation/Base/fr (100% translated)")
- [italiano](/wiki/Handbook:AMD64/Installation/Base/it "Handbook:AMD64/Installation/Base/it (100% translated)")
- [magyar](/wiki/Handbook:AMD64/Installation/Base/hu "Handbook:AMD64/Installation/Base/hu (100% translated)")
- [polski](/wiki/Handbook:AMD64/Installation/Base/pl "Handbook:AMD64/Installation/Base/pl (100% translated)")
- [português do Brasil](/wiki/Handbook:AMD64/Installation/Base/pt-br "Handbook:AMD64/Installation/Base/pt-br (100% translated)")
- čeština
- [русский](/wiki/Handbook:AMD64/Installation/Base/ru "Handbook:AMD64/Installation/Base/ru (100% translated)")
- [العربية](/wiki/Handbook:AMD64/Installation/Base/ar "Handbook:AMD64/Installation/Base/ar (100% translated)")
- [தமிழ்](/wiki/Handbook:AMD64/Installation/Base/ta "Handbook:AMD64/Installation/Base/ta (100% translated)")
- [中文](/wiki/Handbook:AMD64/Installation/Base/zh "Handbook:AMD64/Installation/Base/zh (100% translated)")
- [中文（中国大陆）‎](/wiki/Handbook:AMD64/Installation/Base/zh-cn "Handbook:AMD64/Installation/Base/zh-cn (100% translated)")
- [日本語](/wiki/Handbook:AMD64/Installation/Base/ja "Handbook:AMD64/Installation/Base/ja (100% translated)")
- [한국어](/wiki/Handbook:AMD64/Installation/Base/ko "Handbook:AMD64/Installation/Base/ko (100% translated)")

[AMD64 Handbook](/wiki/Handbook:AMD64/cs "Handbook:AMD64/cs")[Instalace](/wiki/Handbook:AMD64/Full/Installation/cs "Handbook:AMD64/Full/Installation/cs")[O instalaci](/wiki/Handbook:AMD64/Installation/About/cs "Handbook:AMD64/Installation/About/cs")[Výběr média](/wiki/Handbook:AMD64/Installation/Media/cs "Handbook:AMD64/Installation/Media/cs")[Konfigurace sítě](/wiki/Handbook:AMD64/Installation/Networking/cs "Handbook:AMD64/Installation/Networking/cs")[Příprava disků](/wiki/Handbook:AMD64/Installation/Disks/cs "Handbook:AMD64/Installation/Disks/cs")[Instalace stage3](/wiki/Handbook:AMD64/Installation/Stage/cs "Handbook:AMD64/Installation/Stage/cs")Instalace základního systému[Konfigurace jádra](/wiki/Handbook:AMD64/Installation/Kernel/cs "Handbook:AMD64/Installation/Kernel/cs")[Konfigurace systému](/wiki/Handbook:AMD64/Installation/System/cs "Handbook:AMD64/Installation/System/cs")[Instalace nástrojů](/wiki/Handbook:AMD64/Installation/Tools/cs "Handbook:AMD64/Installation/Tools/cs")[Konfigurace zavaděče](/wiki/Handbook:AMD64/Installation/Bootloader/cs "Handbook:AMD64/Installation/Bootloader/cs")[Dokončení](/wiki/Handbook:AMD64/Installation/Finalizing/cs "Handbook:AMD64/Installation/Finalizing/cs")[Práce s Gentoo](/wiki/Handbook:AMD64/Full/Working/cs "Handbook:AMD64/Full/Working/cs")[Úvod do Portage](/wiki/Handbook:AMD64/Working/Portage/cs "Handbook:AMD64/Working/Portage/cs")[Přepínače USE](/wiki/Handbook:AMD64/Working/USE/cs "Handbook:AMD64/Working/USE/cs")[Funkce portage](/wiki/Handbook:AMD64/Working/Features/cs "Handbook:AMD64/Working/Features/cs")[Systém initskriptů](/wiki/Handbook:AMD64/Working/Initscripts/cs "Handbook:AMD64/Working/Initscripts/cs")[Proměnné prostředí](/wiki/Handbook:AMD64/Working/EnvVar/cs "Handbook:AMD64/Working/EnvVar/cs")[Práce s Portage](/wiki/Handbook:AMD64/Full/Portage/cs "Handbook:AMD64/Full/Portage/cs")[Soubory a adresáře](/wiki/Handbook:AMD64/Portage/Files/cs "Handbook:AMD64/Portage/Files/cs")[Proměnné](/wiki/Handbook:AMD64/Portage/Variables/cs "Handbook:AMD64/Portage/Variables/cs")[Mísení softwarových větví](/wiki/Handbook:AMD64/Portage/Branches/cs "Handbook:AMD64/Portage/Branches/cs")[Doplňkové nástroje](/wiki/Handbook:AMD64/Portage/Tools/cs "Handbook:AMD64/Portage/Tools/cs")[Vlastní strom Portage](/wiki/Handbook:AMD64/Portage/CustomTree/cs "Handbook:AMD64/Portage/CustomTree/cs")[Pokročilé funkce](/wiki/Handbook:AMD64/Portage/Advanced/cs "Handbook:AMD64/Portage/Advanced/cs")[Konfigurace sítě](/wiki/Handbook:AMD64/Full/Networking/cs "Handbook:AMD64/Full/Networking/cs")[Začínáme](/wiki/Handbook:AMD64/Networking/Introduction/cs "Handbook:AMD64/Networking/Introduction/cs")[Pokročilá konfigurace](/wiki/Handbook:AMD64/Networking/Advanced/cs "Handbook:AMD64/Networking/Advanced/cs")[Modulární síťování](/wiki/Handbook:AMD64/Networking/Modular/cs "Handbook:AMD64/Networking/Modular/cs")[Bezdrátové sítě](/wiki/Handbook:AMD64/Networking/Wireless/cs "Handbook:AMD64/Networking/Wireless/cs")[Přidání funkcí](/wiki/Handbook:AMD64/Networking/Extending/cs "Handbook:AMD64/Networking/Extending/cs")[Dynamická správa](/wiki/Handbook:AMD64/Networking/Dynamic/cs "Handbook:AMD64/Networking/Dynamic/cs")

## Contents

- [1Chrooting](#Chrooting)
  - [1.1Zkopírujte info o DNS](#Zkop.C3.ADrujte_info_o_DNS)
  - [1.2Připojení nezbytných souborových systémů](#P.C5.99ipojen.C3.AD_nezbytn.C3.BDch_souborov.C3.BDch_syst.C3.A9m.C5.AF)
  - [1.3Vstup do nového prostředí](#Vstup_do_nov.C3.A9ho_prost.C5.99ed.C3.AD)
  - [1.4Preparing for a bootloader](#Preparing_for_a_bootloader)
    - [1.4.1UEFI systems](#UEFI_systems)
    - [1.4.2DOS/Legacy BIOS systems](#DOS.2FLegacy_BIOS_systems)
- [2Konfigurace Portage](#Konfigurace_Portage)
  - [2.1Instalace snapshotu repozitáře ebuildů Gentoo z webu](#Instalace_snapshotu_repozit.C3.A1.C5.99e_ebuild.C5.AF_Gentoo_z_webu)
  - [2.2Volitelné: volba zrcadel](#Voliteln.C3.A9:_volba_zrcadel)
    - [2.2.1Optional: rsync mirrors](#Optional:_rsync_mirrors)
  - [2.3Volitelné: Aktualizace repozitáře ebuildů Gentoo](#Voliteln.C3.A9:_Aktualizace_repozit.C3.A1.C5.99e_ebuild.C5.AF_Gentoo)
  - [2.4Přečtení novinek](#P.C5.99e.C4.8Dten.C3.AD_novinek)
  - [2.5Výběr správného profilu](#V.C3.BDb.C4.9Br_spr.C3.A1vn.C3.A9ho_profilu)
    - [2.5.1Ne-multilib](#Ne-multilib)
  - [2.6Optional: Adding a binary package host](#Optional:_Adding_a_binary_package_host)
    - [2.6.1Repository configuration](#Repository_configuration)
    - [2.6.2Installing binary packages](#Installing_binary_packages)
  - [2.7Konfigurace proměnné USE](#Konfigurace_prom.C4.9Bnn.C3.A9_USE)
    - [2.7.1CPU\_FLAGS\_\*](#CPU_FLAGS_.2A)
    - [2.7.2VIDEO\_CARDS](#VIDEO_CARDS)
  - [2.8Volitelné: Nastavení proměnné ACCEPT\_LICENSE](#Voliteln.C3.A9:_Nastaven.C3.AD_prom.C4.9Bnn.C3.A9_ACCEPT_LICENSE)
  - [2.9Aktualizace setu @world](#Aktualizace_setu_.40world)
    - [2.9.1Removing obsolete packages](#Removing_obsolete_packages)
- [3Časová zóna](#.C4.8Casov.C3.A1_z.C3.B3na)
- [4Místní nastavení](#M.C3.ADstn.C3.AD_nastaven.C3.AD)
  - [4.1Generování locales](#Generov.C3.A1n.C3.AD_locales)
  - [4.2Výběr locale](#V.C3.BDb.C4.9Br_locale)
- [5References](#References)

## Chrooting

### Zkopírujte info o DNS

Před tím, než vstoupíme do nového prostředí, nám zbývá udělat ještě jedna věc, a tou je zkopírování DNS informace v /etc/resolv.conf. To musíme udělat pro to, abychom měli jistotu, že síť bude po vstupu do nového prostředí fungovat. Soubor /etc/resolv.conf obsahuje jmenné servery sítě.

Při kopírování se doporučuje předat příkazu cp volbu `--dereference`. Tím zajistíme, že v případě, že je /etc/resolv.conf symbolickým odkazem, dojde ke zkopírování odkazovaného souboru namísto odkazu samotného. V opačném případě by v novém prostředí odkaz ukazoval na neexistující soubor (jelikož cíl odkazu bude pravděpodobně z nového prostředí nedostupný).

`root #` `cp --dereference /etc/resolv.conf /mnt/gentoo/etc/`

### Připojení nezbytných souborových systémů

**Tip**

If using Gentoo's install media, this step can be replaced with simply: arch-chroot /mnt/gentoo.

Za několik okamžiků dojde ke změně kořene na nové umístění. K zajištění toho, aby nové prostředí fungovalo jak má, je nutné zpřístupnit v něm rovněž některé systémy souborů.

Systému souborů, které je nutné mít k dispozici, jsou:

- /proc/ je pseudo systém souborů (vypadá jako běžné soubory, ale ve skutečnosti je generován za běhu) prostřednictvím nějž jádro Linuxu vystavuje informace navenek
- /sys/ je pseudo systém souborů, podobně jako /proc/, který měl původně nahradit a je lépe uspořádaný než /proc/
- /dev/ je normální systém souborů částečně spravovaný Linuxovým správcem zařízení (obvykle udev), který obsahuje všechny soubory zařízení

/proc bude připojen na /mnt/gentoo/proc, zatímco zbývající dvě umístění budou navázána. To znamená, že například /mnt/gentoo/sys bude ve skutečnosti /sys (jedná se pouze o druhé připojení stejného souborového systému) zatímco /mnt/gentoo/proc/ je novým připojením (dalo by se říci instancí) souborového systému.

`root #` `mount --types proc proc /mnt/gentoo/proc
`

`root #` `mount --rbind /sys /mnt/gentoo/sys
`

`root #` `mount --make-rslave /mnt/gentoo/sys
`

`root #` `mount --rbind /dev /mnt/gentoo/dev
`

`root #` `mount --make-rslave /mnt/gentoo/dev`

**Poznámka**

Operace `--make-rslave` jsou potřeba kvůli podpoře systemd v pozdějším průběhu instalace.

**Upozornění**

Při použití instalačních médií nepocházejících od Gentoo uvedené nemusí postačovat. U některých distribucí je /dev/shm symbolickým odkazem na adresář /run/shm, který se stane po chrootu nedostupným. Vše můžete napravit předem vytvořením řádného přípojného bodu za použití tmpfs:

`root #` `test -L /dev/shm && rm /dev/shm && mkdir /dev/shm
`

`root #` `mount --types tmpfs --options nosuid,nodev,noexec shm /dev/shm
`

A ujistěte se, že je nastaven mód 1777

`root #` ` chmod 1777 /dev/shm`

### Vstup do nového prostředí

Nyní, když jsou všechny oddíly připojeny a inicializovány a základní prostřední bylo nainstalováno, je čas přejít do nového instalačního prostředí chrootováním. To znamená, že sezení změní (change) svůj kořenový adresář (root, tj. umístění nejvyšší úrovně, k němuž lze přistoupit) z instalačního prostřední (instalačního CD nebo jiného instalačního média) na instalovaný systém (jmenovitě incializované diskové oddíly). Proto pojmenování "change root" nebo "chroot".

Změna kořenového adresáře se děje ve třech krocích:

1. Umístění kořenového adresáře se změní z / (instalační médium) na /mnt/gentoo (diskové oddíly) s použitím chroot
2. Některá nastavení (ta v /etc/prodile) jsou znovu načtena do paměti s použitím příkazu source
3. Hlavní ukazatel je změněn, což nám pomůže zapamatovat si, že sezení je uvnitř prostředí se změněným kořenovým adresářem.

`root #` `chroot /mnt/gentoo /bin/bash
`

`root #` `source /etc/profile
`

`root #` `export PS1="(chroot) ${PS1}"`

Od tohoto bodu jsou všechny akce prováděny přímo v novém prostředí Gentoo Linuxu. Samozřejmě vše je daleko od dokončení, proto má instalace ještě pár zbývajících sekcí!

**Tip**

Pokud dojde po tomto bodě po přerušení instalace Gentoo, mělo by jít obnovit instalaci od tohoto kroku. Nění třeba znovu rozdělovat disky! Jednoduše [připojte kořenový oddíl](/wiki/Handbook:AMD64/Installation/Disks/cs#Pripojeni_korenoveho_oddilu "Handbook:AMD64/Installation/Disks/cs") a proveďte kroky výše od [Zkopírujte info o DNS](#Zkopirujte_info_o_DNS), abyste znovu vstoupili do pracovního prostředí. Tento postup lze použít i pro opravení chyb zavaděče. Více informací lze najít v článku [chroot](/wiki/Chroot "Chroot").

### Preparing for a bootloader

Now that the new environment has been entered, it is necessary to prepare the new environment for the bootloader. It will be important to have the correct partition mounted when it is time to install the bootloader.

#### UEFI systems

For UEFI systems, /dev/sda1 was formatted with the FAT32 filesystem and will be used as the EFI System Partition (ESP). Create a new /efi directory (if not yet created), and then mount ESP there:

`root #` `mount /dev/sda1 /efi
`

#### DOS/Legacy BIOS systems

For DOS/Legacy BIOS systems, the bootloader will be installed into the /boot directory, therefore mount as follows:

`root #` `mount /dev/sda1 /boot`

## Konfigurace Portage

### Instalace snapshotu repozitáře ebuildů Gentoo z webu

Dalším krokem je instalace snapshotu repozitáře ebuildů Gentoo. Tento snapshot obsahuje sbírku souborů, které informují Portage o tom, jaký software je k dispozici (k instalaci), které profily může administrátor vybírat, zprávy a prvky specifické pro jednotlivé profily apod.

Doporučuje se použí emerge-webrsync pro ty, kteří jsou za restriktivními firewally (používá protokoly HTTP/FTP pro stažení snapshotu) a potřebují nezatěžovat internetovou linku. Uživatelé, kteří nemají žádná omezení sítě nebo rychlosti připojení mohou spokojeně přejít k další sekci níže.

Tímto získáte nejnovější snapshot (který je vydáván denně) z jednoho ze zrcadel Gentoo a nainstaluje ho do systému:

`root #` `emerge-webrsync`

**Poznámka**

Během této operace si může emerge-webrsync stěžovat na to, že chybí umístění /var/db/repos/gentoo/. jde o předpokládaný problém a nic čeho byste se měli obávat - nástroj umístění vytvoří.

Od tohoto bodu může Portage zmínit doporučení k provedení určitých aktualizací. Důvodem je to, že systémové balíčky instalované skrze stage3 soubor mohou mít k dispozici novější verze; Portage o nich teď ví, protože byl nainstalován nový snapshot repozitáře. Prozatím můžete aktualizace softwaru ignorovat; Aktualizace lze vynechat do dokončení instalace.

### Volitelné: volba zrcadel

Pro rychlejší stahování zdrojových kódů se doporučuje zvolit rychlé zrcadlo. Portage hledá v souboru make.conf proměnnou GENTOO\_MIRRORS a používá zrcadla v ní obsažená. Můžete si projít seznam zrcadel Gentoo a vyhledat zcradlo (nebo zrcadla), která jsou blízko vašeho fyzického umístění (jelikož ta jsou nejčastěji nejrychlejší). Nicméně poskytujeme šikovný nástroj zvaný mirrorselect, který poskytuje uživatelům hezké rozhraní pro výběr potřebných zrcadel. Jednoduše se přesuňte na zrcadla dle své volby a vyberte jedno nebo více zrcadel stisknutím `Spacebar`.

A tool called mirrorselect provides a pretty text interface to more quickly query and select suitable mirrors. Just navigate to the mirrors of choice and press `Spacebar` to select one or more mirrors.

`root #` `mirrorselect -i -o >> /mnt/gentoo/etc/portage/make.conf`

Alternatively, a list of active mirrors are [available online](https://www.gentoo.org/downloads/mirrors/).

#### Optional: rsync mirrors

Gentoo also has many rsync mirrors which can speed up downloading the available package list using `emerge --sync` (explained in more detail later) by selecting a mirror closer that is geographically closer to the user.

`root #` `mkdir /etc/portage/repos.conf
`

`root #` `cp /usr/share/portage/config/repos.conf /etc/portage/repos.conf/gentoo.conf
`

Select a mirror close to the system's location from [https://www.gentoo.org/support/rsync-mirrors/](https://www.gentoo.org/support/rsync-mirrors/) and replace the sync-uri default mirror in /etc/portage/repos.conf/gentoo.conf with the desired mirror location.

FILE **`/etc/portage/repos.conf/gentoo.conf`** **UK-based rsync mirror example**

```
[DEFAULT]
main-repo = gentoo
</div>

<div lang="en" dir="ltr" class="mw-content-ltr">
[gentoo]
location = /var/db/repos/gentoo
sync-type = rsync
sync-uri = rsync://rsync.uk.gentoo.org/gentoo-portage/
auto-sync = yes
sync-rsync-verify-jobs = 1
sync-rsync-verify-metamanifest = yes
sync-rsync-verify-max-age = 3
sync-openpgp-key-path = /usr/share/openpgp-keys/gentoo-release.asc
sync-openpgp-keyserver = hkps://keys.gentoo.org
sync-openpgp-key-refresh-retry-count = 40
sync-openpgp-key-refresh-retry-overall-timeout = 1200
sync-openpgp-key-refresh-retry-delay-exp-base = 2
sync-openpgp-key-refresh-retry-delay-max = 60
sync-openpgp-key-refresh-retry-delay-mult = 4
sync-webrsync-verify-signature = yes
sync-git-verify-commit-signature = true

```

### Volitelné: Aktualizace repozitáře ebuildů Gentoo

Je také možné aktualizovat repozitář ebuildů Gentoo na nejnovější verzi. Dřívější příkaz emerge-webrsync nainstaloval relativně nedávný snapshot (ne starší než 24 h), tudíž je tento krok opravdu volitelný.

Předpokládejme, že je tu potřeba nainstalovat poslední aktualizaci balíčku (až do 1 h), potom použijte příkaz emerge --sync. Tento příkaz použije k aktualizaci repozitáře ebuildů Gentoo (který byl dříve stažen pomocí emerge-webrsync) na nejnovější stav protokol rsync.

`root #` `emerge --sync`

V případě pomalého terminálu, jako jsou některé framebuffery nebo sériové konzole, se doporučuje použít volbu `--quiet` k urychlení procesu:

`root #` `emerge --sync --quiet`

### Přečtení novinek

Jakmile je repozitář ebuildů Gentoo synchronizován, může Portage vypsat informativní zprávy podobné těmto:

` * IMPORTANT: 2 news items need reading for repository 'gentoo'.
`

` * Use eselect news to read news items.
`

Novinky byly vytvořeny, aby poskytovaly komunikační médium k šíření zpráv uživatelům prostřednictvím repozitáře ebuildů Gentoo. K jejich správě použijte příkaz eselect news. Aplikace eselect je specifický utilita Gentoo, která poskytuje společné rozhraní pro administraci systému. V tomto případě eselect využívá svůj modul `news`.

V modulu `news` se nejčastěji používají tři operace:

- S pomocí `list` zobrazíte přehled dostupných novinek
- S pomocí `read` si novinky přečtete
- S pomocí `purge` můžete novinky po přečtení odstranit a už je znovu nečíst

`root #` `eselect news list
`

`root #` `eselect news read`

Více informací o čtečce novinek lze získat v její manuálové stránce:

`root #` `man news.eselect`

### Výběr správného profilu

**Tip**

Desktop profiles are not exclusively for _desktop environments_. They are also suitable for minimal window managers like i3 or sway.

"Profile" je stavební kámen každého systému Gentoo. Nejenže určuje výchozí hodnoty USE, CFLAGS a dalších důležitých proměnných, ale omezuje v systému také rozsah určitých verzí balíčků. Tato nastavení jsou udržována vývojáři Gentoo Portage.

Na to jaký profil právě váš systém používá se můžete podívat pomocí eselect za použití modulu `profile`:

**Tip**

On an install media without a scroll-able terminal, `eselect profile list | less` can provide an easy way to list all available profiles while providing the ability to scroll.

`root #` `eselect profile list`

```
Available profile symlink targets:
  [1]   default/linux/amd64/23.0 *
  [2]   default/linux/amd64/23.0/desktop
  [3]   default/linux/amd64/23.0/desktop/gnome
  [4]   default/linux/amd64/23.0/desktop/kde

```

**Poznámka**

Výstup příkazu je pouze příkladný a v průběhu času se mění.

Jak lze vidět, pro některé architektury jsou k dispozici také podprofily desktop.

**Upozornění**

Aktualizace profilů by neměly být brány na lehkou váhu. Při výběru prvního profilu si ověřte, že vybíráte profil odpovídající **stejné verzi** jako je ta použitá ve stage (např. 23.0). Nové verze profilů jsou oznamovány prostřednictvím novinek obsahujících instrukce k migraci. Přečtěte si je a postupujte podle nich než přepnete na nový profil.

Po prohlédnutí dostupných profilů architektury amd64 mohou uživatelé vybrat jiný profil pro svůj systém:

`root #` `eselect profile set 2`

#### Ne-multilib

Pro čistě 64bitového prostřední bez 32bitových aplikací či knihoven, použijte ne-multilib profil:

`root #` `eselect profile list`

```
Available profile symlink targets:
  [1]   default/linux/amd64/13.0 *
  [2]   default/linux/amd64/13.0/desktop
  [3]   default/linux/amd64/13.0/desktop/gnome
  [4]   default/linux/amd64/13.0/desktop/kde
  [5]   default/linux/amd64/13.0/no-multilib

```

Nyní zvolte "no-multilib" profil:

`root #` `eselect profile set 5
`

`root #` `eselect profile list`

```
Available profile symlink targets:
  [1]   default/linux/amd64/13.0
  [2]   default/linux/amd64/13.0/desktop
  [3]   default/linux/amd64/13.0/desktop/gnome
  [4]   default/linux/amd64/13.0/desktop/kde
  [5]   default/linux/amd64/13.0/no-multilib *

```

**Poznámka**

Podprofil `developer` je zvláště určený pro vývoj Gentoo Linuxu a není míněn pro použití ze strany běžných uživatelů.

### Optional: Adding a binary package host

Since December 2023, Gentoo's [Release Engineering team](/wiki/Project:RelEng "Project:RelEng") has offered an [official binary package host](/wiki/Binary_package_quickstart "Binary package quickstart") (colloquially shorted to just "binhost") for use by the general community to retrieve and install binary packages (binpkgs).[\[1\]](#cite_note-1)

Adding a binary package host allows Portage to install cryptographically signed, compiled packages. In many cases, adding a binary package host will _greatly_ decrease the mean time to package installation and adds much benefit when running Gentoo on older, slower, or low power systems.

#### Repository configuration

The repository configuration for a binhost is found in Portage's /etc/portage/binrepos.conf/ directory, which functions similarly to the configuration mentioned in the [Gentoo ebuild repository](/wiki/Handbook:AMD64/Installation/Base/cs#Gentoo_ebuild_repository "Handbook:AMD64/Installation/Base/cs") section.

When defining a binary host, there are two important aspects to consider:

1. The architecture and profile targets within the `sync-uri` value _do_ matter and should align to the respective computer architecture ( **amd64** in this case) and system profile selected in the [Choosing the right profile](/wiki/Handbook:AMD64/Installation/Base/cs#Choosing_the_right_profile "Handbook:AMD64/Installation/Base/cs") section.
2. Selecting a fast, geographically close mirror will generally shorten retrieval time. Review the mirrorselect tool mentioned in the [Optional: Selecting mirrors](/wiki/Handbook:AMD64/Installation/Base/cs#Gentoo_ebuild_repository "Handbook:AMD64/Installation/Base/cs") section or review the [online list of mirrors](https://www.gentoo.org/downloads/mirrors/) where URL values can be discovered.


FILE **`/etc/portage/binrepos.conf/gentoo.conf`** **CDN-based binary package host example**

```
[gentoo]
priority = 9959
# NOTE: Must adjust <arch> and <variant> as appropriate!
sync-uri = https://distfiles.gentoo.org/releases/<arch>/binpackages/<variant>
# x86-64 example sync-uri
# sync-uri = https://distfiles.gentoo.org/releases/amd64/binpackages/23.0/x86-64/
</div>

<div lang="en" dir="ltr" class="mw-content-ltr">
# Introduced in portage-3.0.74 for per-repo verification choices
verify-signature = true
# Default value with >=portage-3.0.77
location = /var/cache/binhost/gentoo

```

If the CPU supports [x86-64-v3](https://en.wikipedia.org/wiki/X86-64#Microarchitecture_levels) then the binhost offers binpkgs which are compiled to support these features as well.

For ease of explaining all main line Intel, Haswell and newer support this and same for AMD's Ryzen range.
For other CPU lines please check or just use binhost example given above.

FILE **`/etc/portage/binrepos.conf/gentoo.conf`** **CDN-based binary package x86-64-v3 host example**

```
[gentoo-x86-64-v3]
priority = 9999
sync-uri = https://distfiles.gentoo.org/releases/amd64/binpackages/23.0/x86-64-v3/
# Introduced in portage-3.0.74 for per-repo verification choices
verify-signature = true
# Default value with >=portage-3.0.77
location = /var/cache/binhost/gentoo-x86-64-v3

```

**Důležité**

If using the x86-64-v3 binhost then is important to add both examples to /etc/portage/binrepos.conf/gentoo.conf as sometimes a package might not yet be available in a x86-64-v3 variant.

#### Installing binary packages

Portage will compile packages from code source by default. It can be instructed to use binary packages in the following ways:

1. The `--getbinpkg` option can be passed when invoking the emerge command. This method of binary package installation is useful to install only a particular binary package.
2. Changing the system's default via Portage's FEATURES variable, which is exposed through the /etc/portage/make.conf file. Applying this configuration change will cause Portage to query the binary package host for the package(s) to be requested and fall back to compiling locally when no results are found.

For example, to have Portage always install available binary packages:

FILE **`/etc/portage/make.conf`** **Configure Portage to use binary packages by default**

```
# Appending getbinpkg to the list of values within the FEATURES variable
FEATURES="${FEATURES} getbinpkg"
# Require signatures
FEATURES="${FEATURES} binpkg-request-signature"

```

Please also run getuto for Portage to set up the necessary keyring for verification:

`root #` `getuto`

Additional Portage features will be discussed in the [the next chapter](/wiki/Handbook:AMD64/Working/Features/cs#Portage_features "Handbook:AMD64/Working/Features/cs") of the handbook.

### Konfigurace proměnné USE

USE je jednou z nejmocnějších proměnných poskytovanou uživatelům. Mnoho programů může být kompilováno s anebo bez volitelné podpory pro určité části. Například některé programy mohou být kompilovány s podporou GTK+ nebo s podporou Qt. Jiné mohou být sestaveny s nebo bez podpory SSL. Některé programy mohou být sestaveny s podporou framebufferu (svgalib) namísto podpory X11 (X serveru).

Většina distribucí kompiluje své balíčky s co nejširší možnou podporou, což zvyšuje velikost programů a čas jejich spuštění, nemluvě o ohromném množství závislostí. S Gentoo může uživatel určit možnosti toho, jak má být program sestaven. Zde přichází do hry proměnná USE.

V proměnné USE uživatel definuje zkratky, které jsou mapovány na volby kompilace. Například `ssl` sestaví podporu SSL v programech, které ji podporují. `-X` odstraní podporu X serveru (všimněte si znaku minus na začátku). `gnome gtk -kde -qt4 -qt5` sestaví programy s podporou GNOME (a GTK+) a nikoli s podporou KDE (a Qt), čímž bude systém plně přizpůsoben prostředí GNOME (pokud jej architektura podporuje).

Výchozí nastavení USE je umístěno v souborech make.defaults profilu Gentoo použitého v systému. Gentoo používá (složitý ) systém dědění ve svých profilech, do kterého se v této části nebudeme pouštět. Nejjednodušším způsobem jak zjistit aktivní nastavení proměnné USE je spuštění příkazu emerge --info a výběr řádku, který má na začátku USE:

`root #` `emerge --info | grep ^USE`

```
USE="X acl alsa amd64 berkdb bindist bzip2 cli cracklib crypt cxx dri ..."

```

**Poznámka**

Výše uvedený příklad je zkrácený, ve skutečnosti je seznam hodnot USE o mnoho, mnoho rozsáhlejší.

Úplný popis dostupných USE přepínačů lze najít v systému v souboru /var/db/repos/gentoo/profiles/use.desc.

`root #` `less /var/db/repos/gentoo/profiles/use.desc`

Uvnitř příkazu less jde listovat pomocí kláves `↑` a `↓`, ukončit jej lze stisknutím `q`.

Jako příklad uvádíme nastavení USE pro systém s KDE s podporou DVD, ALSA a vypalování CD:

`root #` `nano -w /etc/portage/make.conf`

FILE **`/etc/portage/make.conf`** **Zapnutí USE v systému s KDE/Plasma a s podporou DVD, ALSA a vypalování CD**

```
USE="-gtk -gnome qt4 qt kde dvd alsa cdr"

```

Jamile je USE definována v /etc/portage/make.conf je "přidán" (nebo "odebrán" pokud přepínač začíná znaménkem `-`) na výchozí seznam. Uživatelé, kteří chtějí potlačit všechna výchozí nastavení USE a spravovat je zcela samo, nechť vloží na začátek definice USE v make.conf `-*`:

FILE **`/etc/portage/make.conf`** **Potlačení výchozí přepínačů USE**

```
USE="-* X acl alsa"

```

**Upozornění**

Ačkoli je to možné, nastavení `-*` (jak je uvedeno v příkladu shora) se nedoporučuje, protože výchozí nastavení byla pečlivě volena s ohledem na některé balíčky, aby byly vyloučeny konflikty a jiné chyby.

#### CPU\_FLAGS\_\*

Some architectures (including AMD64/X86, ARM, PPC) have a [USE\_EXPAND](/wiki/USE_EXPAND "USE EXPAND") variable called [CPU\_FLAGS\_<ARCH>](/wiki/CPU_FLAGS_X86 "CPU FLAGS X86"), where <ARCH> is replaced with the relevant system architecture name.

**Důležité**

Do not be confused! **AMD64** and **X86** systems share some common architecture, so the proper variable name for **AMD64** systems is CPU\_FLAGS\_X86.

This is used to configure the build to compile in specific assembly code or other intrinsics, usually hand-written or otherwise extra,
and is **not** the same as asking the compiler to output optimized code for a certain CPU feature (e.g. `-march=`).

Users should set this variable in addition to configuring their COMMON\_FLAGS as desired.

A few steps are needed to set this up:

`root #` `emerge --ask --oneshot app-portage/cpuid2cpuflags`

Inspect the output manually if curious:

`root #` `cpuid2cpuflags`

Then copy the output into package.use:

`root #` `echo "*/* $(cpuid2cpuflags)" > /etc/portage/package.use/00cpu-flags`

#### VIDEO\_CARDS

By default a profile already sets a few video cards. For many reasons this is not ideal, but a useful safety net.

To configure the system correctly the user needs to first unset the preset video cards with `VIDEO_CARDS: -*` then set the correct card for that system.

FILE **`/etc/portage/package.use/00video_cards`** **Example**

```
*/* VIDEO_CARDS: -* <GPU DRIVER NAME>

```

Below is a table that can be used to easily compare the different video card types to their respective `VIDEO_CARDS` value.

GPU
VIDEO\_CARDS
(Official) Nvidia Maxwell and newer`nvidia`Nouveau Nvidia [Supported list](/wiki/NVIDIA "NVIDIA")`nouveau`AMD since Sea Islands`amdgpu radeonsi`ATI and older AMDSee [radeon#Feature support](/wiki/Radeon#Feature_support "Radeon")Intel Nehalem and newer`intel`Intel Gen 2 and 3 [Supported list](/wiki/Intel#.23Feature_support.2Fcs "Intel")`intel i915`QEMU/KVM`virgl`WSL`d3d12`

Below is a few examples of a properly set package.use for _VIDEO\_CARDS_:

FILE **`/etc/portage/package.use/00video_cards`** **AMD example**

```
*/* VIDEO_CARDS: -* amdgpu radeonsi

```

FILE **`/etc/portage/package.use/00video_cards`** **Intel example**

```
*/* VIDEO_CARDS: -* intel

```

FILE **`/etc/portage/package.use/00video_cards`** **Nvidia example**

```
*/* VIDEO_CARDS: -* nvidia

```

Details for various GPU(s) can be found at the [AMDGPU](/wiki/AMDGPU "AMDGPU"), [Intel](/wiki/Intel "Intel"), [Nouveau (Open Source)](/wiki/Nouveau "Nouveau"), or [NVIDIA (Proprietary)](/wiki/NVIDIA "NVIDIA") articles.

### Volitelné: Nastavení proměnné ACCEPT\_LICENSE

Starting with Gentoo Linux Enhancement Proposal 23 (GLEP 23), a mechanism was created to allow system administrators the ability to "regulate the software they install with regards to licenses... Some want a system free of any software that is not OSI-approved; others are simply curious as to what licenses they are implicitly accepting."[\[2\]](#cite_note-2) With a motivation to have more granular control over the type of software running on a Gentoo system, the ACCEPT\_LICENSE variable was born.

Gentoo přichází s přednastavenými hodnotami v profilech, například:

`user $` `portageq envvar ACCEPT_LICENSE`

```
@FREE
```

Skupiny licencí definované v repozitáři Gentoo, spravované [Projektem licencí Gentoo](/wiki/Project:Licenses "Project:Licenses"), jsou:

Název skupinyPopis
@GPL-COMPATIBLELicence kompatibilní s GPL schválené Free Software Foundation [\[a\_license 1\]](#cite_note-3)@FSF-APPROVEDLicence svobodného softwaru schválené FSF (includes @GPL-COMPATIBLE)
@OSI-APPROVEDLicence schválené Open Source Initiative [\[a\_license 2\]](#cite_note-4)@MISC-FREERůzné licence, které jsou pravděpodobně svobodné, tzn. splňující definici svobodného software [\[a\_license 3\]](#cite_note-5) ale nejsou schválené FSF nebo OSI.
@FREE-SOFTWARECombines @FSF-APPROVED, @OSI-APPROVED and @MISC-FREE
@FSF-APPROVED-OTHERLicence schválené FSF pro "svobodnou dokumentaci" a "díla k určená k použití společně se softwarem a dokumentací" (včetně fontů)
@MISC-FREE-DOCSRůzné licence pro svobodné dokumenty a další díla (včetně fontů), které splňují definici svobody [\[a\_license 4\]](#cite_note-6) ale NEJSOU v seznamu @FSF-APPROVED-OTHER
@FREE-DOCUMENTSKombinuje @FSF-APPROVED-OTHER a @MISC-FREE-DOCS
@FREEMetasada všech licencí obsahujících svobodu užívat, sdílet, měnit a sdílet změny. Kombinuje @FREE-SOFTWARE a @FREE-DOCUMENTS
@BINARY-REDISTRIBUTABLELicence, které dovolují alespoň volné šíření software v binární podobě. Zahrnuje @FREE
@EULALicenční ujednání, která se snaží omezit vaše práva. Jsou restriktivnější než "všechna-práva-vyhrazena" nebo vyžadují výslovný souhlas.

Some common license groups include:

A list of software licenses grouped according to their kinds.
NameDescription
`@GPL-COMPATIBLE`GPL compatible licenses approved by the Free Software Foundation [\[a\_license 5\]](#cite_note-7)`@FSF-APPROVED`Free software licenses approved by the FSF (includes `@GPL-COMPATIBLE`)
`@OSI-APPROVED`Licenses approved by the Open Source Initiative [\[a\_license 6\]](#cite_note-8)`@MISC-FREE`Misc licenses that are probably free software, i.e. follow the Free Software Definition [\[a\_license 7\]](#cite_note-9) but are not approved by either FSF or OSI
`@FREE-SOFTWARE`Combines `@FSF-APPROVED`, `@OSI-APPROVED`, and `@MISC-FREE`.
`@FSF-APPROVED-OTHER`FSF-approved licenses for "free documentation" and "works of practical use besides software and documentation" (including fonts)
`@MISC-FREE-DOCS`Misc licenses for free documents and other works (including fonts) that follow the free definition [\[a\_license 8\]](#cite_note-10) but are NOT listed in `@FSF-APPROVED-OTHER`.
`@FREE-DOCUMENTS`Combines `@FSF-APPROVED-OTHER` and `@MISC-FREE-DOCS`.
`@FREE`Metaset of all licenses with the freedom to use, share, modify and share modifications. Combines `@FREE-SOFTWARE` and `@FREE-DOCUMENTS`.
`@BINARY-REDISTRIBUTABLE`Licenses that at least permit free redistribution of the software in binary form. Includes `@FREE`.
`@EULA`License agreements that try to take away your rights. These are more restrictive than "all-rights-reserved" or require explicit approval

1. [↑](#cite_ref-3)[https://www.gnu.org/licenses/license-list.html](https://www.gnu.org/licenses/license-list.html)
2. [↑](#cite_ref-4)[https://www.opensource.org/licenses](https://www.opensource.org/licenses)
3. [↑](#cite_ref-5)[https://www.gnu.org/philosophy/free-sw.html](https://www.gnu.org/philosophy/free-sw.html)
4. [↑](#cite_ref-6)[https://freedomdefined.org/](https://freedomdefined.org/)
5. [↑](#cite_ref-7)[https://www.gnu.org/licenses/license-list.html](https://www.gnu.org/licenses/license-list.html)
6. [↑](#cite_ref-8)[https://www.opensource.org/licenses](https://www.opensource.org/licenses)
7. [↑](#cite_ref-9)[https://www.gnu.org/philosophy/free-sw.html](https://www.gnu.org/philosophy/free-sw.html)
8. [↑](#cite_ref-10)[https://freedomdefined.org/](https://freedomdefined.org/)

Currently set system wide acceptable license values can be viewed via:

`user $` `portageq envvar ACCEPT_LICENSE`

```
@FREE
```

As visible in the output, the default value is to only allow software which has been grouped into the `@FREE` category to be installed.

Specific licenses or licenses groups for a system can be defined in the following locations:

- System wide within the selected profile - this sets the default value.
- System wide within the /etc/portage/make.conf file. System administrators override the profile's default value within this file.
- Per-package within a /etc/portage/package.license file.
- Per-package within a /etc/portage/package.license/ _directory_ of files.

Můžete si je přizpůsobit na úrovni systému změnou /etc/portage/make.conf. Výchozí hodnota povoluje pouze licence výslovně schválené Free Sowtware Foundantion, Open Source Intiative nebo které splňují definici svobodného software:

FILE **`/etc/portage/make.conf`** **Přizpůsobení ACCEPT\_LICENSE**

```
ACCEPT_LICENSE="-* @FREE"

```

Změna na úrovni balíčků pak může být provedena, pokud jde nezbytná a žádaná, například:

FILE **`/etc/portage/package.license/kernel`** **Příklad souhlasu s licencí**

```
app-arch/unrar unRAR
sys-kernel/linux-firmware @BINARY-REDISTRIBUTABLE
sys-firmware/intel-microcode intel-ucode

```

`root #` `mkdir /etc/portage/package.license`

Software license details for an individual Gentoo package are stored within the LICENSE variable of the associated ebuild. One package may have one or many software licenses, therefore it be necessary to specify multiple acceptable licenses for a single package.

FILE **`/etc/portage/package.license/kernel`** **Accepting licenses on a per-package basis**

```
app-arch/unrar unRAR
sys-kernel/linux-firmware linux-fw-redistributable
sys-firmware/intel-microcode intel-ucode

```

**Důležité**

Proměnná LICENCE v ebuildu je jen pomůckou pro vývojáře a uživatele Gentoo. Nejedná se o právní stanovisko a není žádná garance, že odpovídá skutečnému stavu. Nespoléhejte tedy na ni a zkontrolujte balíček pořádně sami, včetně všech souborů, které používáte.

### Aktualizace setu @world

**Tip**

Pokud jste zvolili profil desktop v celé jeho šíři, může dojít k velkému prodloužení času potřebnému k provedení instalačního procesu. Pokud jste po časovým tlakem, můžete se řídit následujícím "jednoduchým pravidlem": čím kratší je název profilu, tím méně specifický je [set @world](/wiki/World_set_(Portage) "World set (Portage)"); čím méně specifičtější je set @world, tím méně balíčků systém vyžaduje. Jinými slovy:

- Výběr `default/linux/amd64/23.0` vyžaduje aktualizaci velmi malého množství balíčků, zatímco
- Výběr `default/linux/amd64/23.0/desktop/gnome/systemd` vyžaduje instalaci mnoha balíčků, jelikož init systém se mění z OpenRC na systemd a budou instalovány frameworky desktopu GNOME.

Následující krok je "nezbytný" aby mohl systém aplikovat jakékoli aktualizace a změny přepínačů USE, které se objevily po té, co byla sestavena stage3, nebo po změně profilu:

1. A profile target _different_ from the stage file has been selected.
2. Additional USE flags have been set for installed packages.

Readers who are performing an 'install Gentoo speed run' may safely skip @world set updates until _after_ their system has rebooted into the new Gentoo environment.

Readers who are performing a slow run can have Portage perform updates for package, profile, and/or USE flag changes at the present time:

`root #` `emerge --ask --verbose --update --deep --newuse @world`

Readers who added a binary host [above](/wiki/Handbook:AMD64/Installation/Base/cs#Optional:_Adding_a_binary_package_host "Handbook:AMD64/Installation/Base/cs") can add --getbinpkg (or -g) in order to fetch packages from the binary host instead of compiling them:

`root #` `emerge --ask --verbose --update --deep --newuse --getbinpkg @world`

#### Removing obsolete packages

It is important to always _depclean_ after system upgrades to remove obsolete packages. Review the output carefully with emerge --depclean --pretend to see if any of the to-be-cleaned packages should be kept if personally using them. To keep a package which would otherwise be depcleaned, use emerge --noreplace foo.

`root #` `emerge --ask --pretend --depclean`

If happy, then proceed with a real depclean:

`root #` `emerge --ask --depclean`

## Časová zóna

**Poznámka**

This step does not apply to users of the musl libc. Users who do not know what that means should perform this step.

Vyhýbejte se časovým zónám v /usr/share/zoneinfo/Etc/GMT\*, jelikož jejich názvy neodpovídají předpokládaným zónám. Například GMT-8 je ve skutečnosti GMT+8.

Vyberte časovou zónu systému. Podívejte se na nabídku časových zón v /usr/share/zoneinfo/ a pak jednu z nich zapište do souboru /etc/timezone.

`root #` `ls /usr/share/zoneinfo`

`root #` `ls -l /usr/share/zoneinfo/Europe/`

```
total 256
-rw-r--r-- 1 root root 2933 Dec  3 17:19 Amsterdam
-rw-r--r-- 1 root root 1742 Dec  3 17:19 Andorra
-rw-r--r-- 1 root root 1151 Dec  3 17:19 Astrakhan
-rw-r--r-- 1 root root 2262 Dec  3 17:19 Athens
-rw-r--r-- 1 root root 3664 Dec  3 17:19 Belfast
-rw-r--r-- 1 root root 1920 Dec  3 17:19 Belgrade
-rw-r--r-- 1 root root 2298 Dec  3 17:19 Berlin
-rw-r--r-- 1 root root 2301 Dec  3 17:19 Bratislava
-rw-r--r-- 1 root root 2933 Dec  3 17:19 Brussels
...

```

Suppose the timezone of choice is Europe/Brussels, to select this timezone, a symlink can be created from this zoneinfo file to /etc/localtime:

`root #` `ln -sf ../usr/share/zoneinfo/Europe/Brussels /etc/localtime`

**Tip**

The target path with `../` at the start is _relative to the link location_, not the current directory.

**Poznámka**

An absolute path can be used for the symlink, but a relative link is also created by systemd's timedatectl and is more compatible with alternate _ROOT_ s.

## Místní nastavení

**Poznámka**

This step does not apply to users of the musl libc. Users who do not know what that means should perform this step.

### Generování locales

A default installation of Gentoo Linux provides an archive that contains all supported locales, numbering 500 or more. However, it is typical for an administrator to require only one or two of these. In that case, the /etc/locale.gen configuration file may be populated with a list of the required locales. By default, locale-gen shall read this file and compile only the locales that are specified, saving both time and space in the longer term.

Locales specifikují nejen jazyk, který uživatel používá k interakci se systémem, ale také pravidla pro řazení řetězců, zobrazení datumů, času apod.
Locales jsou citlivá na velikost písmen a musí být zapsána přesně tak, jak jsou popsána. Úplný seznam dostupných locales lze najít v souboru /usr/share/i18n/SUPPORTED.

`root #` `nano -w /etc/locale.gen`

Následující místní nastavení jsou příkladem, jak získat angličtinu (Spojené státy) a němčinu (Německo) s náležitým formátem znaků (třeba UTF-8).

FILE **`/etc/locale.gen`** **Zapnutí místního nastavení US a DE s vhodnými formáty znaků**

```
en_US ISO-8859-1
en_US.UTF-8 UTF-8
de_DE ISO-8859-1
de_DE.UTF-8 UTF-8

```

1. Non UTF-8 locales are discouraged and should only be used in special circumstances.
2. en\_US ISO-8859-1
3. de\_DE ISO-8859-1

}}

**Upozornění**

_Silně_ doporučujeme doplnit alespoň jedno UTF-8 locale, jelikož mnoho aplikací jej může vyžadovat, aby mohly být řádně sestaveny.

Dalším krokem je spuštění příkazu locale-gen. Tento příkaz vygeneruje všechna locales specifikovaná v souboru /etc/locale.gen.

`root #` `locale-gen`

K ověření toho, že všechna zvolená místní nastavení jsou k dispozici spusťte locale -a.

On systemd installs, localectl can be used, e.g. localectl set-locale ... or localectl list-locales.

### Výběr locale

Jakmile je vše hotovo, je čas zvolit systémové místní nastavení. Opět k tomu použijeme eselect, nyní s modulem `locale`.

Pomocí eselect locale list, zobrazíme dostupné volby:

`root #` `eselect locale list`

```
Dostupné volby proměnné LANG:
  [1] C
  [2] POSIX
  [3] en_US
  [4] en_US.iso88591
  [5] en_US.utf8
  [6] de_DE
  [7] de_DE.iso88591
  [8] de_DE.iso885915
  [9] de_DE.utf8
  [ ] (free form)

```

Pomocí eselect locale set <číslo> lze nastavit správné locale:

`root #` `eselect locale set 9`

Ručně lze téhož dosáhnout prostřednictvím souboru /etc/env.d/02locale:

FILE **`/etc/env.d/02locale`** **Manually Ruční nastavení systémové definice místního nastavení**

```
LANG="de_DE.UTF-8"
LC_COLLATE="C"

```

Nastavení locale předejde zobrazování varování a chyb v průběhu sestavování jádra a kompilace softwaru v pozdějších fázích instalace.

Nyní znovu načtěte prostředí:

`root #` `env-update && source /etc/profile && export PS1="(chroot) ${PS1}"`

Kompletní [průvodce místním nastavením](/wiki/Localization/Guide "Localization/Guide"), poskytuje další rady ohledně procesu výběru locale. Dalším článkem hodným pozornosti je průvodce [UTF-8](/wiki/UTF-8 "UTF-8") obsahující specifické informace ohledně zapnutí podpory UTF-8 v systému.

## References

1. [↑](#cite_ref-1)[https://www.gentoo.org/news/2023/12/29/Gentoo-binary.html](https://www.gentoo.org/news/2023/12/29/Gentoo-binary.html)
2. [↑](#cite_ref-2)[https://www.gentoo.org/glep/glep-0023.html#id7](https://www.gentoo.org/glep/glep-0023.html#id7)

[← Instalace stage3](/wiki/Handbook:AMD64/Installation/Stage/cs "Previous part") [Home](/wiki/Handbook:AMD64 "Handbook:AMD64") [Konfigurace jádra →](/wiki/Handbook:AMD64/Installation/Kernel/cs "Next part")