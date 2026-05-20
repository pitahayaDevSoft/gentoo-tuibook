# Stage

Other languages:

- [Deutsch](/wiki/Handbook:X86/Installation/Stage/de "Handbuch:X86/Installation/Stage (100% translated)")
- [English](/wiki/Handbook:X86/Installation/Stage "Handbook:X86/Installation/Stage (100% translated)")
- [Türkçe](/wiki/Handbook:X86/Installation/Stage/tr "Handbook:X86/Installation/Stage/tr (0% translated)")
- [español](/wiki/Handbook:X86/Installation/Stage/es "Manual de Gentoo: X86/Instalación/Stage (100% translated)")
- [français](/wiki/Handbook:X86/Installation/Stage/fr "Handbook:X86/Installation/Stage/fr (100% translated)")
- [italiano](/wiki/Handbook:X86/Installation/Stage/it "Handbook:X86/Installation/Stage (100% translated)")
- [magyar](/wiki/Handbook:X86/Installation/Stage/hu "Handbook:X86/Installation/Stage/hu (100% translated)")
- [polski](/wiki/Handbook:X86/Installation/Stage/pl "Handbook:X86/Installation/Stage (100% translated)")
- [português do Brasil](/wiki/Handbook:X86/Installation/Stage/pt-br "Manual:X86/Instalação/Stage (100% translated)")
- čeština
- [русский](/wiki/Handbook:X86/Installation/Stage/ru "Handbook:X86/Installation/Stage (100% translated)")
- [தமிழ்](/wiki/Handbook:X86/Installation/Stage/ta "கையேடு:X86/நிறுவல்/நிலை (100% translated)")
- [中文（中国大陆）‎](/wiki/Handbook:X86/Installation/Stage/zh-cn "手册：X86/安装/安装stage3 (100% translated)")
- [日本語](/wiki/Handbook:X86/Installation/Stage/ja "ハンドブック:X86/インストール/Stage (100% translated)")
- [한국어](/wiki/Handbook:X86/Installation/Stage/ko "Handbook:X86/Installation/Stage/ko (100% translated)")

[X86 Handbook](/wiki/Handbook:X86/cs "Handbook:X86/cs")[Instalace](/wiki/Handbook:X86/Full/Installation/cs "Handbook:X86/Full/Installation/cs")[O instalaci](/wiki/Handbook:X86/Installation/About/cs "Handbook:X86/Installation/About/cs")[Výběr média](/wiki/Handbook:X86/Installation/Media/cs "Handbook:X86/Installation/Media/cs")[Konfigurace sítě](/wiki/Handbook:X86/Installation/Networking/cs "Handbook:X86/Installation/Networking/cs")[Příprava disků](/wiki/Handbook:X86/Installation/Disks/cs "Handbook:X86/Installation/Disks/cs")Instalace stage3[Instalace základního systému](/wiki/Handbook:X86/Installation/Base/cs "Handbook:X86/Installation/Base/cs")[Konfigurace jádra](/wiki/Handbook:X86/Installation/Kernel/cs "Handbook:X86/Installation/Kernel/cs")[Konfigurace systému](/wiki/Handbook:X86/Installation/System/cs "Handbook:X86/Installation/System/cs")[Instalace nástrojů](/wiki/Handbook:X86/Installation/Tools/cs "Handbook:X86/Installation/Tools/cs")[Konfigurace zavaděče](/wiki/Handbook:X86/Installation/Bootloader/cs "Handbook:X86/Installation/Bootloader/cs")[Dokončení](/wiki/Handbook:X86/Installation/Finalizing/cs "Handbook:X86/Installation/Finalizing/cs")[Práce s Gentoo](/wiki/Handbook:X86/Full/Working/cs "Handbook:X86/Full/Working/cs")[Úvod do Portage](/wiki/Handbook:X86/Working/Portage/cs "Handbook:X86/Working/Portage/cs")[Přepínače USE](/wiki/Handbook:X86/Working/USE/cs "Handbook:X86/Working/USE/cs")[Funkce portage](/wiki/Handbook:X86/Working/Features/cs "Handbook:X86/Working/Features/cs")[Systém initskriptů](/wiki/Handbook:X86/Working/Initscripts/cs "Handbook:X86/Working/Initscripts/cs")[Proměnné prostředí](/wiki/Handbook:X86/Working/EnvVar/cs "Handbook:X86/Working/EnvVar/cs")[Práce s Portage](/wiki/Handbook:X86/Full/Portage/cs "Handbook:X86/Full/Portage/cs")[Soubory a adresáře](/wiki/Handbook:X86/Portage/Files/cs "Handbook:X86/Portage/Files/cs")[Proměnné](/wiki/Handbook:X86/Portage/Variables/cs "Handbook:X86/Portage/Variables/cs")[Mísení softwarových větví](/wiki/Handbook:X86/Portage/Branches/cs "Handbook:X86/Portage/Branches/cs")[Doplňkové nástroje](/wiki/Handbook:X86/Portage/Tools/cs "Handbook:X86/Portage/Tools/cs")[Vlastní strom Portage](/wiki/Handbook:X86/Portage/CustomTree/cs "Handbook:X86/Portage/CustomTree/cs")[Pokročilé funkce](/wiki/Handbook:X86/Portage/Advanced/cs "Handbook:X86/Portage/Advanced/cs")[Konfigurace sítě](/wiki/Handbook:X86/Full/Networking/cs "Handbook:X86/Full/Networking/cs")[Začínáme](/wiki/Handbook:X86/Networking/Introduction/cs "Handbook:X86/Networking/Introduction/cs")[Pokročilá konfigurace](/wiki/Handbook:X86/Networking/Advanced/cs "Handbook:X86/Networking/Advanced/cs")[Modulární síťování](/wiki/Handbook:X86/Networking/Modular/cs "Handbook:X86/Networking/Modular/cs")[Bezdrátové sítě](/wiki/Handbook:X86/Networking/Wireless/cs "Handbook:X86/Networking/Wireless/cs")[Přidání funkcí](/wiki/Handbook:X86/Networking/Extending/cs "Handbook:X86/Networking/Extending/cs")[Dynamická správa](/wiki/Handbook:X86/Networking/Dynamic/cs "Handbook:X86/Networking/Dynamic/cs")

## Contents

- [1Výběr balíku stage](#V.C3.BDb.C4.9Br_bal.C3.ADku_stage)
- [2OpenRC](#OpenRC)
- [3systemd](#systemd)
  - [3.1Multilib (32 and 64bitový)](#Multilib_.2832_and_64bitov.C3.BD.29)
  - [3.2Ne-multilib (čistě 64bitový)](#Ne-multilib_.28.C4.8Dist.C4.9B_64bitov.C3.BD.29)
- [4Stažení archivu stage](#Sta.C5.BEen.C3.AD_archivu_stage)
- [5Nastavení data a času](#Nastaven.C3.AD_data_a_.C4.8Dasu)
  - [5.1Automaticky](#Automaticky)
  - [5.2Ručně](#Ru.C4.8Dn.C4.9B)
  - [5.3Grafické prohlížeče](#Grafick.C3.A9_prohl.C3.AD.C5.BEe.C4.8De)
  - [5.4Prohlížeče příkazového řádku](#Prohl.C3.AD.C5.BEe.C4.8De_p.C5.99.C3.ADkazov.C3.A9ho_.C5.99.C3.A1dku)
  - [5.5Ověření a potvrzení](#Ov.C4.9B.C5.99en.C3.AD_a_potvrzen.C3.AD)
- [6Instalace stage balíku](#Instalace_stage_bal.C3.ADku)
- [7Konfigurace kompilačních voleb](#Konfigurace_kompila.C4.8Dn.C3.ADch_voleb)
  - [7.1Úvod](#.C3.9Avod)
  - [7.2CFLAGS a CXXFLAGS](#CFLAGS_a_CXXFLAGS)
  - [7.3RUSTFLAGS](#RUSTFLAGS)
  - [7.4MAKEOPTS](#MAKEOPTS)
  - [7.5Připravit, pozor, teď!](#P.C5.99ipravit.2C_pozor.2C_te.C4.8F.21)
- [8References](#References)

### Výběr balíku stage

**Tip**

On supported architectures, it is recommended for users targeting a desktop (graphical) operating system environment to use a stage file with the term `desktop` within the name. These files include packages such as [llvm-core/llvm](https://packages.gentoo.org/packages/llvm-core/llvm) and [dev-lang/rust-bin](https://packages.gentoo.org/packages/dev-lang/rust-bin) and USE flag tuning which will greatly improve install time.

The [stage file](/wiki/Stage_file "Stage file") acts as the seed of a Gentoo install. Stage files are generated by the [Release Engineering Team](/wiki/Project:RelEng "Project:RelEng") with [Catalyst](/wiki/Catalyst "Catalyst"). Stage files are based on specific [profiles](/wiki/Profile_(Portage) "Profile (Portage)"), and contain an almost-complete system.

When choosing a stage file, it's important to pick one with profile targets corresponding to the desired system type.

**Důležité**

While it's possible to make major profile changes after an installation has been established, switching requires substantial effort and consideration, and is outside the scope of this installation manual. Switching init systems is difficult, but switching from `no-multilib` to `multilib` requires extensive Gentoo and low-level toolchain knowledge.

Většina uživatelů nemusí používat nabízené "pokročilé" balíky; jsou určeny pro specifické konfigurace softwaru a hardwaru.

### OpenRC

[OpenRC](/wiki/OpenRC/cs "OpenRC/cs") is a dependency-based init system (responsible for starting up system services once the kernel has booted) that maintains compatibility with the system provided init program, normally located in /sbin/init. It is Gentoo's native and original init system, but is also deployed by a few other Linux distributions and BSD systems.

### systemd

systemd is a modern SysV-style init and rc replacement for Linux systems. It is used as the primary init system by a majority of Linux distributions. systemd is fully supported in Gentoo and works for its intended purpose. If something seems lacking in the Handbook for a systemd install path, review the [systemd article](/wiki/Systemd "Systemd") _before_ asking for support.

#### Multilib (32 and 64bitový)

**Poznámka**

Not every architecture has a multilib option. Many only run with native code. Multilib is most commonly applied to **amd64**.

Výběr základního stage balíku může později ušetřit spoustu čas během procesu instalace, zvláště ve chvíli, kdy budete [vybírat profil systému](/wiki/Handbook:X86/Installation/Base/cs#Vyber_spravneho_profilu "Handbook:X86/Installation/Base/cs"). Vyběr stage balíku má přímý dopad na budoucí konfiguraci systému a může vám později ušetřit muka. Balík multilib obsahuje, je-li to možné, 64bitové knihovny a jako zálohu volí 32bitové verze pouze tehdy, je-li to potřeba z pohledu kompatibility. Jde o vynikající volbu pro většinu instalací, jelikož poskytuje vysokou míru flexibility a přizpůsobení v budoucnosti. Ti, kteří touží po tom, aby byl jejich systém schopen snadno měnit profily, by měli stáhnout balík multilib pro svoji procesorovou architekturu.

**Tip**

Using `multilib` targets makes it easier to switch profiles later, compared to `no-multilib`

#### Ne-multilib (čistě 64bitový)

**Upozornění**

Vemte na vědomí, že přechod z ne-multilib na multilib systém vyžaduje extrémní znalosti fungování Gentoo a nízkoúrovňových nástrojů (dokonce i našim [vyvojářům základních nástrojů](/wiki/Project:Toolchain "Project:Toolchain") by proběhl mráz po zádech). Není to nic pro slabé povahy a jedná se o záležitost, která je mimo záběr tohoto průvodce.

Vyberete-li non-multilib balík jako základ pro systém poskytne výhradně 64bitové prostředí operačního systému. Tím se stává možnost přejít na multilib profil nepravděpodobné, i když stále možné. **Ti, kteří s Gentoo teprve začínají by neměli volit non-multilib balík pokud to není absolutně nezbytné.**

### Stažení archivu stage

Before downloading the _stage file_, the current directory should be set to the location of the mount used for the install:

`root #` `cd /mnt/gentoo`

### Nastavení data a času

Stage archives are generally obtained using HTTPS which requires relatively accurate system time. Clock skew can prevent downloads from working, and can cause unpredictable errors if the system time is adjusted by any considerable amount after installation.

Ověřte aktuální datum a čas spuštěním příkazu date:

`root #` `date`

```
Po Říj  3 13:16:22 PDT 2016

```

Pokud je zobrazené datum/čas chybné, aktualizujte ho jedním z následujících způsobů.

#### Automaticky

Using [Network\_Time\_Protocol](/wiki/Network_Time_Protocol "Network Time Protocol") to correct clock skew is typically easier and more reliable than manually setting the system clock.

chronyd, part of [net-misc/chrony](https://packages.gentoo.org/packages/net-misc/chrony) can be used to update the system clock to UTC with:

`root #` `ntpd -q -g`

**Důležité**

Systems without a functioning Real-Time Clock (RTC) must sync the system clock at every system start, and on regular intervals thereafter. This is also beneficial for systems with a RTC, as the battery could fail, and clock skew can accumulate.

**Upozornění**

Standard NTP traffic is not authenticated, it is important to verify time data obtained from the network.

#### Ručně

When NTP access is unavailable, date can be used to manually set the system clock.

Na všech Linuxových systémech se doporučuje používat čas UTC. Později během instalace určíme časovou zónu. Tím změníme zobrazení hodin na místní čas.

Příkaz date může být použit také k ručnímu nastavení času systémových hodin. Použijte syntaxi `MMDDhhmmYYYY` (měsíc, den, hodina, minuta a rok).

Příklad nastavení datumu na 3. října, 13.16 v roce 2016:

`root #` `date 100313162016`

#### Grafické prohlížeče

Ti, kteří používají prostředí s plně grafickými prohlížeči nebudou mít problém zkopírovat URL ze [sekce download](https://www.gentoo.org/downloads/#other-arches) hlavní webové stránky. Jednoduše zvolte odpovídající tab, klikněte pravým tlačítkem na odkaz stage souboru, potom na Zkopírovat adresu odkazu (Firefox) nebo Zkopírovat umístění odkazu (Chromium), abyste zkopírovali odkaz do schránky, následně předejte odkaz nástroji wget na příkazovém řádku, abyste archiv stáhli:

`root #` `wget <VLOŽENÁ_URL_STAGE>`

#### Prohlížeče příkazového řádku

Tradicionalisté nebo "staří" uživatelé Gentoo pracující výhradně v příkazovém řádku mohou upřednostňovat použití programu links, negrafického prohlížeče ovládaného pomocí menu. Ke stažení stage surfujte na seznam zrcadel Gentoo tímto způsobem:

`root #` `links https://www.gentoo.org/downloads/mirrors/`

HTTP proxy s prohlížečem links použijete tak, že mu předáte volbu `-http-proxy`:

`root #` `links -http-proxy proxy.server.com:8080 https://www.gentoo.org/downloads/mirrors/`

Vedle links můžete použít také prohlížeč lynx. Stejně jako links se jedná o negrafický prohlížeč, který však nemá menu.

`root #` `lynx https://www.gentoo.org/downloads/mirrors/`

Pokud potřebujete definovat proxy, exportujte proměnné http\_proxy a/nebo ftp\_proxy.

`root #` `export http_proxy="http://proxy.server.com:port"
`

`root #` `export ftp_proxy="http://proxy.server.com:port"`

V seznamu zvolte zrcadlo umístěné někde oblíž. Obvykle postačí HTTP zrcadla, ale ostatní protokoly jsou v nabídce také. Přesuňte se do adresáře releases/x86/autobuilds/. Zobrazí se vám veškeré nabízené soubory stage (mohou být uloženy v podadresářích pojmenovaných podle jednotlivých subarchitekturách). Jeden z nich vyberte a stiskněte `d`, čímž dojde ke stažení.

Jakmile je soubor stage stažen, je možné ověřit integritu a potvrdit obsah stage archivu. Ti, které to zajímá, nechť přejdou do [další sekce](#Verifying_and_validating).

Ti, kteří nemají zájem ověřovat a potvrzovat soubor stage mohou zavřít prohlížeč v příkazovém řádku stisknutím `q` a přesunout se přímo do sekce [rozbalení stage archivu](#Unpacking_the_stage_tarball).

#### Ověření a potvrzení

Stejně jako u minimálních instalačních CD jsou k dispozici i doplňkové soubory k ověření a kontrole souboru stage. I když mohou být tyto kroky přeskořeny, jsou tyto soubory poskytovány uživatelům, kteří dbají na legitimitu souboru(ů), které stáhli.

`root #` `wget https://distfiles.gentoo.org/releases/`

- Soubor .CONTENTS obsahuje seznam všech souborů v balíku.
- Soubor .DIGESTS obsahuje kontrolní součty souboru stage dle různých algoritmů.
- Soubor .DIGESTS.asc obsahuje stejně jako soubor .DIGESTS kontrolní součty souboru stage dle různých algoritmů a navíc je zároveň kryptograficky podepsán pro ověření toho, že pochází z projektu Gentoo.

Stejně jako u souboru ISO je možné ověřit kryptografický podpis souboru .DIGESTS.asc pomocí příkazu gpg a ověřit tak, že kontrolní součty nebyly měněny:

For official Gentoo live images, the [sec-keys/openpgp-keys-gentoo-release](https://packages.gentoo.org/packages/sec-keys/openpgp-keys-gentoo-release) package provides PGP signing keys for automated releases. The keys must first be imported into the user's session in order to be used for verification:

`root #` `gpg --import /usr/share/openpgp-keys/gentoo-release.asc`

For all non-official live images which offer gpg and wget in the live environment, refer to the earlier [Verifying the downloaded files](/wiki/Handbook:X86/Installation/Media/cs#Verifying_the_downloaded_files "Handbook:X86/Installation/Media/cs") section.

Verify the signature of the tarball and, optionally, associated checksum files:

`root #` `gpg --verify stage3-x86-<release>.tar.?(bz2|xz){.DIGESTS.asc,} `

If verification succeeds, "Good signature from" will be in the output of the previous command(s).

The fingerprints of the OpenPGP keys used for signing release media can be found on the [release media signatures page](https://www.gentoo.org/downloads/signatures/).

**Poznámka**

Most stages are now explicitly [suffixed](https://www.gentoo.org/news/2021/07/20/more-downloads.html) with the init system type (openrc or systemd), although some architectures may still be missing these for now.

Použijte příkaz openssl a porovnejte výstup s kontrolními součty obsaženými v souborech .DIGESTS nebo .DIGESTS.asc.

Například ověření kontrolního součtu SHA512 provedete takto:

`root #` `openssl dgst -r -sha512 stage3-x86-<release>.tar.?(bz2|xz)`

`dgst` instructs the openssl command to use the Message Digest sub-command, `-r` prints the digest output in coreutils format, and `-sha512` selects the SHA512 digest.

Ověření kontrolního součtu Whirpool:

`root #` `openssl dgst -r -whirlpool stage3-x86-<release>.tar.?(bz2|xz)`

Porovnejte výstupy těchto příkazů s hodnotami zapsanými v souborech .DIGESTS(.asc). Hodnoty se musí shodovat, v opačném případě mohlo dojít k porušení staženého souboru (nebo souboru se součty).

Jiným cestou je použití příkazu sha512sum:

`root #` `sha512sum stage3-x86-<release>.tar.?(bz2|xz)`

The `--check` option instructs sha256sum to read a list of expected files and associated hashes, and then print an associated "OK" for each file that calculates correctly or a "FAILED" for files that do not.

## Instalace stage balíku

Nyní stažený stage soubor rozbalte do systému. K tomu použijeme příkaz tar:

`root #` `tar xpvf stage3-*.tar.bz2 --xattrs-include='*.*' --numeric-owner`

Přesvědčte se, že jste použili stejné volby ( `xpf` a `--xattrs-include='*.*'`). Volba `x` znamená e **x** trahovat, `p` zachovat (angl. **p** reserve) oprávnění a `f` značí, že chceme rozbalit soubor (ang. **f** ile), nikoli standardní vstup. Volba `--xattrs-include='*.*'` zajistí zahrnutí rozšířených atributů, které jsou v archivu také uloženy. Konečně volba `--numeric-owner` se použije k zajištění, že uživatelská a skupinová ID souborů rozbalených z archivu zůstanou stejná, dle záměrů týmu Gentoo spravujícího vydání distribuce (i když dobrodružní uživatelé nepoužijí oficiální instalační média Gentoo).

- `x` e **x** tract, instructs tar to extract the contents of the archive.
- `p` **p** reserve permissions.
- `v` **v** erbose output.
- `f` **f** ile, provides tar with the name of the input archive.
- `--xattrs-include='*.*'` Preserves extended attributes in all namespaces stored in the archive.
- `--numeric-owner` Ensure that the user and group IDs of files being extracted from the tarball remain the same as Gentoo's release engineering team intended (even if adventurous users are not using official Gentoo live environments for the installation process).
- `-C /mnt/gentoo` Extract files to the root partition regardless of the current directory.

Nyní, když je stage soubor rozbalený, pokračujte ke [konfiguraci kompilačních voleb](#Configuring_the_compile_options).

## Konfigurace kompilačních voleb

### Úvod

Gentoo je možné vyladit nastavením několika proměnných, které mají dopad na chování Portage, oficiálního správce balíčků Gentoo. Všechny proměnné mohou být nastaveny jako proměnné prostředí (pomocí příkazu export), nicméně to není trvalé řešení. Nastavení se uchovává v konfiguračním souboru /etc/portage/make.conf, odkud je Portage načítá.

**Poznámka**

Technically variables can be exported via the [shell's](/wiki/Shell "Shell") profile or rc files, however that is not best practice for basic system administration.

Portage reads in the [make.conf](/wiki/Make.conf "Make.conf") file when it runs, which will change runtime behavior depending on the values saved in the file. make.conf can be considered the primary configuration file for Portage, so treat its content carefully.

**Poznámka**

Komentovaný výpis všech možných proměnných lze najít v souboru /mnt/gentoo/usr/share/portage/make.conf.example. K úspěšnému nainstalování Gentoo je třeba nastavit pouze proměnné uvedené níže.

For a successful Gentoo installation only the variables that are mentioned below need to be set.}}

Spusťte editor (v tomto průvodci používáme nano) a upravte optimalizační proměnné, které nyní probereme.

`root #` `nano -w /mnt/gentoo/etc/portage/make.conf`

Ze souboru make.conf.example je zjevné, jak má být uspořádán: komentované řádky začínají "#", ostatní řádky definují proměnné prostřednictvím syntaxe PROMENNA="obsah". Nyní se podíváme na několik z těchto proměnných.

### CFLAGS a CXXFLAGS

Proměnné CFLAGS a CXXFLAGS definují optimalizační parametry C potažmo C++ kompilátorů GCC. Na tomto místě se definují všeobecně, pro dosažení maximálního výkonu by musely být nastaveny pro každý program zvlášť. Důvodem je, že každý program je různý. Nicméně to není únosné, proto se tyto volby nastavují v souboru make.conf.

V souboru make.conf by měly být obecně nastaveny takové optimalizační přepínače, jejichž použití povede k co nejlepší odezvě systému. Do této proměnné neumisťujte experimentální nastavení; příliš mnoho optimalizací může způsobit špatné chování programů (pády nebo hůře, závadné chování).

Nebudeme zde popisovat všechny možnosti vyladění. Nastudovat si je můžete v [GNU online manuálu](https://gcc.gnu.org/onlinedocs/) nebo info stránce gcc (info gcc \- lze spustit jen ve fungujícím systému Linux). Také samotný soubor make.conf.example obsahuje mnoho příkladů a informací, nezapomeňte si jej také přečíst.

Prvním nastavením je hodnota parametru `-march=` nebo `-mtune=`, který stanoví jméno cílové architektury. Možné volby jsou popsány v souborumake.conf.example (jako komentáře). Běžně používanou hodnotou je "native", která kompiléru říká, aby používal cílovou architekturu daného systému (toho, na nějž uživatel instaluje Gentoo).

Druhým je parametr `-O` (jedná se o velké O, nikoli o nulu), který určuje třídu optimalizací gcc. Možné třídy jsou s (pro optimalizaci velikosti), 0 (nula - žádné optimalizace), 1, 2 nebo dokonce 3 pro více urychlujících optimalizací (každá třída obsahuje stejné parametry, jako ta přechozí a k tomu nějaké další). `-O2` je doporučenou výchozí hodnotou. O `-O3` je známo, že při plošném použití v celém systému způsobuje problémy, tudíž doporučujeme držet se `-O2`.

Dalším oblíbeným optimalizačním parametrem je `-pipe` (použije roury namísto dočasných souborů ke komunikaci mezi různými fázemi kompilace). Ten nemá dopad na generovaný kód, ale používá více paměti. V systémech s nízkým množstvím paměti může kvůli tomu dojít k zabití gcc. V takovém případě tento parametr nepoužívejte.

Použití `-fomit-frame-pointer` (který nezachová frame pointer v registru pro funkce, které jej nepotřebují) může mít závažné dopady na debugging aplikací.

Jakmile jsou proměnné CFLAGS a CXXFLAGS definovány, složte optimalizační přepínače do jednoho řetězce. Výchozí hodnoty obsažené ve stage3 archivu by měly být dostatečné. Následující zápis je pouze příkladem:

CODE **Příklad proměnných CFLAGS a CXXFLAGS**

```
# Přepínače kompilátoru používané pro všechny jazyky
COMMON_FLAGS="-O2 -march=i686 -pipe"
# Použij stejnou hodnotu pro obě proměnné
CFLAGS="${COMMON_FLAGS}"
CXXFLAGS="${COMMON_FLAGS}"

```

**Poznámka**

Ačkoli článek o [optimalizaci GCC](/wiki/GCC_optimization "GCC optimization") obsahuje více informací o tom, jak mohou různé volby kompilace ovlivnit systém, může být pro začátečníka praktičtější začít s optimalizací systému článkem o [bezpečném nastavení CFLAG](/wiki/Safe_CFLAGS "Safe CFLAGS").

### RUSTFLAGS

Many programs are now written in Rust which has its own way of optimising. By default Rust optimizes to level 3 on all release builds unless a project changes it so this should be left as is. A full optimization list (known as codegen) that can be passed to the rust compiler is available at [https://doc.rust-lang.org/rustc/codegen-options/index.html](https://doc.rust-lang.org/rustc/codegen-options/index.html).

The most useful optimization would be to tell Rust to compile for your CPU using the following example:

FILE **`/etc/portage/make.conf`** **RUSTFLAGS Example**

```
RUSTFLAGS="${RUSTFLAGS} -C target-cpu=native"

```

To get a list of supported CPUs in Rust run:

`root #` `rustc -C target-cpu=help`

This will show what the default and also which CPU will be selected with native.

**Poznámka**

The above command only works on desktop stage3 tarballs or after emerging [dev-lang/rust-bin](https://packages.gentoo.org/packages/dev-lang/rust-bin) or [dev-lang/rust](https://packages.gentoo.org/packages/dev-lang/rust).

### MAKEOPTS

Proměnná MAKEOPTS určuje, kolik paralelních kompilací může být prováděno při instalaci balíku. Dobrou volbou je počet CPU (nebo jader CPU) v systému plus jedna, ale toto vodítko není vždy dokonalé.

Further, as of Portage 3.0.53[\[1\]](#cite_note-1), if left undefined, Portage's default behavior is to set the MAKEOPTS load-average value to the same number of threads returned by nproc.

A good choice is the smaller of: the number of threads the CPU has, or the total amount of system RAM divided by 2 GiB.

**Upozornění**

Using a large number of jobs can significantly impact memory consumption. A good recommendation is to have at least 2 GiB of RAM for every job specified (so, e.g. `-j6` requires _at least_ 12 GiB). To avoid running out of memory, lower the number of jobs to fit the available memory.

**Tip**

When using parallel emerges ( `--jobs`), the effective number of jobs run can grow exponentially (up to make jobs multiplied by emerge jobs). This can be worked around by running a localhost-only distcc configuration that will limit the number of compiler instances per host.

CODE **Příklad deklarace MAKEOPTS v make.conf**

```
MAKEOPTS="-j2"

```

Search for MAKEOPTS in man 5 make.conf for more details.

### Připravit, pozor, teď!

Upravte soubor /mnt/gentoo/etc/portage/make.conf tak, aby odpovídal vašim osobním preferencím a uložte jej (uživatelé nano stisknou `Ctrl` + `X`).

## References

1. [↑](#cite_ref-1)[https://gitweb.gentoo.org/proj/portage.git/commit/?id=de7be7f45ee54e3f789def46542919550687d15e](https://gitweb.gentoo.org/proj/portage.git/commit/?id=de7be7f45ee54e3f789def46542919550687d15e)

[← Příprava disků](/wiki/Handbook:X86/Installation/Disks/cs "Previous part") [Home](/wiki/Handbook:X86 "Handbook:X86") [Instalace základního systému →](/wiki/Handbook:X86/Installation/Base/cs "Next part")