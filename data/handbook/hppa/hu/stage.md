# Stage

Other languages:

- [Deutsch](/wiki/Handbook:HPPA/Installation/Stage/de "Handbuch:HPPA/Installation/Stage (100% translated)")
- [English](/wiki/Handbook:HPPA/Installation/Stage "Handbook:HPPA/Installation/Stage (100% translated)")
- [español](/wiki/Handbook:HPPA/Installation/Stage/es "Manual de Gentoo: HPPA/Instalación/Stage (100% translated)")
- [français](/wiki/Handbook:HPPA/Installation/Stage/fr "Handbook:HPPA/Installation/Stage/fr (100% translated)")
- [italiano](/wiki/Handbook:HPPA/Installation/Stage/it "Handbook:HPPA/Installation/Stage/it (100% translated)")
- magyar
- [polski](/wiki/Handbook:HPPA/Installation/Stage/pl "Handbook:HPPA/Installation/Stage (100% translated)")
- [português do Brasil](/wiki/Handbook:HPPA/Installation/Stage/pt-br "Manual:HPPA/Instalação/Stage (100% translated)")
- [русский](/wiki/Handbook:HPPA/Installation/Stage/ru "Handbook:HPPA/Installation/Stage (100% translated)")
- [தமிழ்](/wiki/Handbook:HPPA/Installation/Stage/ta "கையேடு:HPPA/நிறுவல்/நிலை (100% translated)")
- [中文（中国大陆）‎](/wiki/Handbook:HPPA/Installation/Stage/zh-cn "手册：HPPA/安装/安装stage3 (100% translated)")
- [日本語](/wiki/Handbook:HPPA/Installation/Stage/ja "ハンドブック:HPPA/インストール/Stage (100% translated)")
- [한국어](/wiki/Handbook:HPPA/Installation/Stage/ko "Handbook:HPPA/Installation/Stage/ko (100% translated)")

[HPPA kézikönyv](/wiki/Handbook:HPPA/hu "Handbook:HPPA/hu")[A Gentoo Linux telepítése](/wiki/Handbook:HPPA/Full/Installation/hu "Handbook:HPPA/Full/Installation/hu")[A telepítésről](/wiki/Handbook:HPPA/Installation/About/hu "Handbook:HPPA/Installation/About/hu")[Telepítőképfájl kiválasztása](/wiki/Handbook:HPPA/Installation/Media/hu "Handbook:HPPA/Installation/Media/hu")[Hálózat beállítása](/wiki/Handbook:HPPA/Installation/Networking/hu "Handbook:HPPA/Installation/Networking/hu")[Adathordozók előkészítése](/wiki/Handbook:HPPA/Installation/Disks/hu "Handbook:HPPA/Installation/Disks/hu")Stage fájl[Alaprendszer telepítése](/wiki/Handbook:HPPA/Installation/Base/hu "Handbook:HPPA/Installation/Base/hu")[Kernel beállítása](/wiki/Handbook:HPPA/Installation/Kernel/hu "Handbook:HPPA/Installation/Kernel/hu")[Rendszer beállítása](/wiki/Handbook:HPPA/Installation/System/hu "Handbook:HPPA/Installation/System/hu")[Eszközök telepítése](/wiki/Handbook:HPPA/Installation/Tools/hu "Handbook:HPPA/Installation/Tools/hu")[Bootloader beállítása](/wiki/Handbook:HPPA/Installation/Bootloader/hu "Handbook:HPPA/Installation/Bootloader/hu")[Telepítés véglegesítése](/wiki/Handbook:HPPA/Installation/Finalizing/hu "Handbook:HPPA/Installation/Finalizing/hu")[Munka a Gentoo rendszerrel](/wiki/Handbook:HPPA/Full/Working/hu "Handbook:HPPA/Full/Working/hu")[Portage bemutatása](/wiki/Handbook:HPPA/Working/Portage/hu "Handbook:HPPA/Working/Portage/hu")[USE jelölőzászlók](/wiki/Handbook:HPPA/Working/USE/hu "Handbook:HPPA/Working/USE/hu")[Portage jellemzői](/wiki/Handbook:HPPA/Working/Features/hu "Handbook:HPPA/Working/Features/hu")[Init-szkript rendszer](/wiki/Handbook:HPPA/Working/Initscripts/hu "Handbook:HPPA/Working/Initscripts/hu")[Környezeti változók](/wiki/Handbook:HPPA/Working/EnvVar/hu "Handbook:HPPA/Working/EnvVar/hu")[Munka a Portage szoftvercsomag-kezelővel](/wiki/Handbook:HPPA/Full/Portage/hu "Handbook:HPPA/Full/Portage/hu")[Fájlok és könyvtárak](/wiki/Handbook:HPPA/Portage/Files/hu "Handbook:HPPA/Portage/Files/hu")[Változók](/wiki/Handbook:HPPA/Portage/Variables/hu "Handbook:HPPA/Portage/Variables/hu")[Szoftverágak keverése](/wiki/Handbook:HPPA/Portage/Branches/hu "Handbook:HPPA/Portage/Branches/hu")[További eszközök](/wiki/Handbook:HPPA/Portage/Tools/hu "Handbook:HPPA/Portage/Tools/hu")[Egyéni szoftvercsomag-tárolók](/wiki/Handbook:HPPA/Portage/CustomTree/hu "Handbook:HPPA/Portage/CustomTree/hu")[Fejlett funkciók](/wiki/Handbook:HPPA/Portage/Advanced/hu "Handbook:HPPA/Portage/Advanced/hu")[Hálózat beállítása OpenRC init-rendszeren](/wiki/Handbook:HPPA/Full/Networking/hu "Handbook:HPPA/Full/Networking/hu")[Munka elkezdése](/wiki/Handbook:HPPA/Networking/Introduction/hu "Handbook:HPPA/Networking/Introduction/hu")[Fejlett beállítások](/wiki/Handbook:HPPA/Networking/Advanced/hu "Handbook:HPPA/Networking/Advanced/hu")[Moduláris hálózat](/wiki/Handbook:HPPA/Networking/Modular/hu "Handbook:HPPA/Networking/Modular/hu")[Vezeték nélküli (Wi-Fi)](/wiki/Handbook:HPPA/Networking/Wireless/hu "Handbook:HPPA/Networking/Wireless/hu")[Funkcionalitás hozzáadása](/wiki/Handbook:HPPA/Networking/Extending/hu "Handbook:HPPA/Networking/Extending/hu")[Dinamikus menedzsment](/wiki/Handbook:HPPA/Networking/Dynamic/hu "Handbook:HPPA/Networking/Dynamic/hu")

## Contents

- [1Stage fájl kiválasztása](#Stage_f.C3.A1jl_kiv.C3.A1laszt.C3.A1sa)
  - [1.1OpenRC](#OpenRC)
  - [1.2systemd](#systemd)
  - [1.3Multilib (32 bit és 64 bit)](#Multilib_.2832_bit_.C3.A9s_64_bit.29)
  - [1.4No-multilib (tisztán csak 64 bit)](#No-multilib_.28tiszt.C3.A1n_csak_64_bit.29)
- [2Stage fájl letöltése](#Stage_f.C3.A1jl_let.C3.B6lt.C3.A9se)
  - [2.1Dátum és idő beállítása](#D.C3.A1tum_.C3.A9s_id.C5.91_be.C3.A1ll.C3.ADt.C3.A1sa)
    - [2.1.1Automatikus](#Automatikus)
    - [2.1.2Manuális](#Manu.C3.A1lis)
  - [2.2Grafikus böngészők](#Grafikus_b.C3.B6ng.C3.A9sz.C5.91k)
  - [2.3Parancssorban működő böngészők](#Parancssorban_m.C5.B1k.C3.B6d.C5.91_b.C3.B6ng.C3.A9sz.C5.91k)
  - [2.4Ellenőrzés és érvényesítés](#Ellen.C5.91rz.C3.A9s_.C3.A9s_.C3.A9rv.C3.A9nyes.C3.ADt.C3.A9s)
- [3Stage fájl telepítése](#Stage_f.C3.A1jl_telep.C3.ADt.C3.A9se)
- [4Forráskódfordítási opciók beállítása](#Forr.C3.A1sk.C3.B3dford.C3.ADt.C3.A1si_opci.C3.B3k_be.C3.A1ll.C3.ADt.C3.A1sa)
  - [4.1Bevezetés](#Bevezet.C3.A9s)
  - [4.2CFLAGS és CXXFLAGS](#CFLAGS_.C3.A9s_CXXFLAGS)
  - [4.3RUSTFLAGS](#RUSTFLAGS)
  - [4.4MAKEOPTS](#MAKEOPTS)
  - [4.5Felkészülni, vigyázz, rajt!](#Felk.C3.A9sz.C3.BClni.2C_vigy.C3.A1zz.2C_rajt.21)
- [5Hivatkozások](#Hivatkoz.C3.A1sok)

## Stage fájl kiválasztása

**Tip**

A támogatott architektúrák esetében ajánlott, hogy azok a felhasználók, akik asztali (grafikus) operációs rendszer környezetet céloznak meg, olyan stage fájlt használjanak, amely fájl nevében szerepel a `desktop` kifejezés. Ezek a fájlok olyan szoftvercsomagokat tartalmaznak, mint például a [sys-devel/llvm](https://packages.gentoo.org/packages/sys-devel/llvm) és a [dev-lang/rust-bin](https://packages.gentoo.org/packages/dev-lang/rust-bin), valamint olyan USE jelölőzászló beállításokat tartalmaznak, amelyek jelentősen javítják a telepítési időt.

A [stage fájl](/wiki/Stage_file/hu "Stage file/hu") a Gentoo telepítés kiindulópontjaként (a Gentoo telepítőmagjaként) szolgál. A stage fájlokat a [Catalyst](/wiki/Catalyst "Catalyst") segítségével a [Kiadási Mérnöki Csapat](/wiki/Project:RelEng "Project:RelEng") (Release Engineering Team) hozza létre. A stage fájlok konkrét [profilok](/wiki/Profile_(Portage) "Profile (Portage)") alapján készülnek, és egy majdnem teljes operációs rendszert tartalmaznak.

Amikor Ön stage fájlt választ, fontos, hogy olyat válasszon, amelynek profilcéljai megfelelnek az Ön által kívánt operációs rendszer típusának.

**Important**

Noha a telepítés után is lehetséges a jelentős profilátmódosítás (módosításból fakadó profilátalakításról van szó), de az átmódosítás jelentős erőfeszítést és komoly megfontolást igényel, amely nem tartozik a jelen telepítési kézikönyv hatálya alá. Az init-rendszerek váltása nehéz feladat, de például a `no-multilib` profilról a `multilib` profilra való átmódosítás is széles körű Gentoo ismeretet és alacsony szintű eszközlánc-ismereteket igényel.

**Tip**

A legtöbb felhasználónak nincs szüksége az 'advanced' tömörített fájlok használatára. Ezek szokatlan vagy fejlett szoftver- vagy hardverbeállításokhoz valók.

### OpenRC

Az [OpenRC](/wiki/OpenRC/hu "OpenRC/hu") egy szoftverfüggőség-alapú init-rendszer (amely a kernel elindulása után a rendszerszolgáltatások indításáért felelős), amely fenntartja a kompatibilitást a rendszer által biztosított init programmal, amely általában a /sbin/init helyen található. Ez a Gentoo natív és eredeti init-rendszere, de néhány más Linux disztribúció és BSD rendszer is telepíti.

### systemd

A systemd egy modern SysV-stílusú init és rc helyettesítő Linux rendszerek számára. A Linux disztribúciók többsége elsődleges init rendszerként használja. A systemd a Gentoo-n teljes mértékben támogatott, és a rendeltetésének megfelelően működik. Ha úgy tűnik, hogy a kézikönyvben valami hiányzik a systemd telepítési útmutatójából, akkor tekintse át a [systemd cikket](/wiki/Systemd/hu "Systemd/hu"), _mielőtt_ támogatást kérne.

### Multilib (32 bit és 64 bit)

**Note**

Nem minden architektúra rendelkezik multilib opcióval. Sokan csak natív kóddal futnak. A multilib leggyakrabban az **amd64** architektúrára vonatkozik.

A multilib profil lehetőség szerint 64 bites könyvtárakat használ, és csak akkor tér vissza a 32 bites verziókra, ha a kompatibilitás feltétlenül szükséges. Ez egy kiváló lehetőség a legtöbb telepítéshez, mert nagy fokú rugalmasságot biztosít a jövőbeni testreszabáshoz.

**Tip**

A `multilib` célpontok használata megkönnyíti a későbbi profilváltást a `no-multilib` profillal rendelkezőhöz képest.

### No-multilib (tisztán csak 64 bit)

**Warning**

Azok az olvasók, akik most kezdik használni a Gentoo operációs rendszert, _ne_ válasszák a no-multilib .tar tömörített fájlt, hacsak nem feltétlenül szükséges. A `no-multilib` rendszerről a `multilib` rendszerre való áttérés a Gentoo és az mélyebb szintű eszközlánc rendkívül részletes ismeretét kívánja meg (ez még valószínűleg, a [Toolchain fejlesztőinket](/wiki/Project:Toolchain "Project:Toolchain") is egy kicsit megborzongatja). Nem a gyenge szívűeknek való, és túlmutat ezen útmutató hatókörén.

Ha az operációs rendszer alapjául egy no-multilib .tar tömörített fájlt választ, akkor teljes 64 bites operációs rendszer környezetet biztosít – 32 bites szoftverektől mentesen. Ez gyakorlatilag megterhelővé teszi a `multilib` profilokra váltás lehetőségét, bár technikailag még mindig lehetséges.

## Stage fájl letöltése

A _stage fájl_ letöltése előtt az aktuális könyvtárat a telepítéshez használt csatolási pont helyére kell állítani:

`root #` `cd /mnt/gentoo`

### Dátum és idő beállítása

A stage tömörített fájlokat általában HTTPS használatával szerzik be a felhasználók, ami viszonylag pontos rendszeridőt igényel. Az óraeltérés megakadályozhatja a letöltések sikerességét, és kiszámíthatatlan hibákat okozhat, ha a rendszeridőt a telepítés után jelentős mértékben módosítják.

Az aktuális dátum és idő a date parancs segítségével ellenőrizhető:

`root #` `date`

```
Mon Oct  3 13:16:22 PDT 2021

```

Ha a megjelenített dátum/idő néhány percnél pontatlanabb, akkor frissíteni kell az alábbi módszerek valamelyikével.

#### Automatikus

Az [NTP](/wiki/NTP "NTP") használata az óraeltolódás korrigálására általában egyszerűbb és megbízhatóbb, mint a rendszeróra manuális úton történő beállítása.

A chronyd program, amely a [net-misc/chrony](https://packages.gentoo.org/packages/net-misc/chrony) szoftvercsomag része, használható a rendszeróra UTC-re történő frissítésére a következő parancs kiadásával:

`root #` `chronyd -q`

**Important**

A működő Valós Idejű Óra (RTC - Real-Time Clock) nélküli rendszereknek minden rendszerindításkor és azt követően rendszeres időközönként szinkronizálniuk kell a rendszerórát. Ez hasznos az RTC-vel rendelkező rendszerek számára is, mivel a számítógép akkumulátora lemerülhet, és az óraeltérés megnövekedhet.

**Warning**

A szabványos NTP forgalom nincs hitelesítve, ezért fontos ellenőrizni a hálózaton keresztül kapott időadatokat.

#### Manuális

Amikor az NTP hozzáférés nem elérhető, a date parancs használható a rendszeróra manuális beállítására.

**Note**

Az UTC idő ajánlott minden Linux rendszerhez. Később meghatározásra kerül egy rendszeridőzóna, amely megváltoztatja az eltolást, amikor a dátum megjelenik.

A következő argumentumformátumot használják az idő beállítására: `MMDDhhmmYYYY` szintaxis ( **M** onth, **D** ay, **h** our, **m** inute és **Y** ear).

Például annak érdekében, hogy az időpontot október 3, 13:16, 2021 időpontra állítsuk be, használja következő parancsot:

`root #` `date 100313162021`

### Grafikus böngészők

Azok a felhasználók, akik teljesen grafikus webböngészőt használnak, könnyedén kimásolhatják a stage fájl URL-jét a fő weboldal [letöltés aloldaláról](https://www.gentoo.org/downloads/#other-arches). Egyszerűen válassza ki a megfelelő fület, kattintson jobb gombbal a stage fájl hivatkozására, majd válassza ki a Hivatkozás másolása opciót a hivatkozás vágólapra másolásához, majd a parancssorban illessze be a hivatkozást a wget segédprogramhoz a fokozat stage fájl letöltésének érdekében:

`root #` `wget <STAGE_FÁJL_IDE_BEILLESZTETT_URL_CÍME>`

### Parancssorban működő böngészők

A hagyományosabb olvasók vagy a 'régi motoros' Gentoo felhasználók, akik kizárólag parancssorból dolgoznak, inkább a links ([www-client/links](https://packages.gentoo.org/packages/www-client/links)) nevű, nem grafikus, menü-vezérelt böngészőt szokták használni. A stage fájl letöltésének érdekében böngésszen a Gentoo tükörszerver listáján a következőképpen:

`root #` `links https://www.gentoo.org/downloads/mirrors/`

Annak érdekében, hogy a links használatakor a HTTP proxy legyen használva, adja meg a `-http-proxy` opciónak a proxy URL címét:

`root #` `links -http-proxy proxy.server.com:8080 https://www.gentoo.org/downloads/mirrors/`

A links mellett létezik a lynx ([www-client/lynx](https://packages.gentoo.org/packages/www-client/lynx)) böngésző is. A links böngészőhöz hasonlóan ez is egy nem grafikus böngésző, de nem menüvezérelt.

`root #` `lynx https://www.gentoo.org/downloads/mirrors/`

Ha egy proxy-t kell meghatározni, akkor exportálja a http\_proxy és/vagy ftp\_proxy változókat:

`root #` `export http_proxy="http://proxy.server.com:port"
`

`root #` `export ftp_proxy="http://proxy.server.com:port"`

A tükörszerverlistában válasszon egy közeli tükörszervert. Általában elegendőek a HTTP türkszerverek, de más protokollok is elérhetők. Lépjen a releases/hppa/autobuilds/ könyvtárba. Itt az összes elérhető stage fájl megjelenik (lehet, hogy az egyes alcsoport-architektúrák nevét viselő alkönyvtárakban vannak tárolva). Válasszon ki egyet, és nyomja meg a `d` billentyűgombot a letöltéshez.

Miután a stage fájl letöltése befejeződött, lehetőség van a stage fájl integritásának ellenőrzésére és tartalmának érvényesítésére. Az eziránt érdeklődők az olvasást folytassák a [következő szakaszban](/wiki/Handbook:HPPA/Installation/Stage/hu#Verifying_and_validating "Handbook:HPPA/Installation/Stage/hu").

Azok, akik nem szeretnék ellenőrizni és érvényesíteni a stage fájlt, kiléphetnek a parancssoros böngészőből a `q` billentyűgomb megnyomásával, és közvetlenül áttérhetnek az [Stage fájl telepítése](/wiki/Handbook:HPPA/Installation/Stage/hu#Unpacking_the_stage_tarball "Handbook:HPPA/Installation/Stage/hu") szakaszra.

### Ellenőrzés és érvényesítés

Ahogyan a minimal telepítő CD képfájlok esetében, a stage fájl ellenőrzéséhez és érvényesítéséhez további letöltések is elérhetők. Bár ezek a lépések kihagyhatók, a fájlok azon felhasználók számára állnak rendelkezésre, akik fontosnak tartják a letöltött fájl(ok) integritását. Az extra fájlok a tükörszerverek könyvtárának gyökerében találhatók. Keresse meg a hardverarchitektúra és a rendszerprofil megfelelő helyét, és töltse le a hozzátartozó .CONTENTS.gz, .DIGESTS és .sha256 fájlokat.

`root #` `wget https://distfiles.gentoo.org/releases/`

- .CONTENTS.gz \- Egy tömörített .gz fájl, amely tartalmazza a stage fájlban lévő összes fájl listáját.
- .DIGESTS \- Többféle kriptográfiai hash algoritmust használó ellenőrzőösszegeket tartalmaz a stage fájlhoz.
- .sha256 \- PGP által aláírt SHA256 ellenőrzőösszeget tartalmaz a stage fájlhoz. Ez a fájl nem biztos, hogy minden stage fájlhoz elérhető letöltésre.

Ahogyan az ISO fájl esetében, a tar.xz fájl kriptográfiai aláírása is ellenőrizhető a gpg segítségével, hogy megbizonyosodjunk, a tömörített tar.xz fájlon semmilyen manipuláció nem történt.

Hivatalos Gentoo bootolható live képfájlok esetén a [sec-keys/openpgp-keys-gentoo-release](https://packages.gentoo.org/packages/sec-keys/openpgp-keys-gentoo-release) szoftvercsomag biztosítja a PGP aláíró kulcsokat az automatizált kiadásokhoz. A kulcsokat először importálni kell a felhasználó munkamenetébe, hogy azok hitelesítésre használhatóak legyenek.

`root #` `gpg --import /usr/share/openpgp-keys/gentoo-release.asc`

Minden nem hivatalos bootolható live képfájl számára, amelyek a bebootolt live környezetükben a gpg és wget programokat kínálják, letölthető és importálható a Gentoo kulcsokat tartalmazó csomag:

Ellenőrizze a .tar tömörített fájlt és opcionálisan a kapcsolódó ellenőrzőösszeg fájlok aláírását:

`root #` `gpg --verify stage3-hppa-<release>-<init>.tar.xz.asc
`

`root #` `gpg --verify stage3-hppa-<release>-<init>.tar.xz.DIGESTS
`

`root #` `gpg --verify stage3-hppa-<release>-<init>.tar.xz.sha256
`

Ha a hitelesítés sikeres, akkor a "Good signature from" szöveg fog megjelenni az előző parancs(ok) kimenetében.

Az OpenPGP kulcsok ujjlenyomatai, amelyeket a kiadási média aláírására használnak, megtalálhatóak a [kiadási média aláírások oldalán](https://www.gentoo.org/downloads/signatures/).

**Note**

A legtöbb stage fájl most már kifejezetten az init rendszer típusával van [megjelölve](https://www.gentoo.org/news/2021/07/20/more-downloads.html) (openrc vagy systemd), bár néhány architektúra esetében ezek még mindig hiányozhatnak.

A kriptográfiai eszközök és segédprogramok, mint például az openssl, sha256sum, vagy sha512sum, használhatóak a .DIGESTS fájlban megadott ellenőrzőösszegek összehasonlítására.

Az SHA512 ellenőrzőösszeg ellenőrzéséhez az openssl segítségével:

`root #` `openssl dgst -r -sha512 stage3-hppa-<release>-<init>.tar.xz`

A `dgst` utasítja az openssl parancsot, hogy használja a Message Digest alparancsot, a `-r` kiírja az összegző kimenetet a coreutils formátumban, és a `-sha512` kiválasztja a SHA512 összegzést.

A BLAKE2B512 ellenőrzőösszeg ellenőrzéséhez az openssl segítségével:

`root #` `openssl dgst -r -blake2b512 stage3-hppa-<release>-<init>.tar.xz`

Hasonlítsa össze az ellenőrzőösszeg parancsok kimenetét a .DIGESTS fájlban található hash és fájlnév párjaival. A párosított értékeknek meg kell egyezniük az ellenőrzőösszeg parancsok kimenetével, ellenkező esetben a letöltött fájl sérült, és törölni kell, majd újra le kell tölteni.

Az SHA256 hash ellenőrzéséhez egy kapcsolódó .sha256 fájlból a sha256sum segédprogram használatával:

`root #` `sha256sum --check stage3-hppa-<release>-<init>.tar.xz.sha256`

A `--check` opció utasítja a sha256sum segédprogramot, hogy olvassa el a várt fájlok és a hozzájuk tartozó hash értékek listáját, majd minden helyesen kiszámított fájlhoz társított "OK" szöveget, illetve a nem megfelelő fájlokhoz társított "FAILED" szöveget jelenítse meg.

## Stage fájl telepítése

Miután a _stage fájl_ le lett töltve és ellenőrizve lett, a .tar.xz tömörített fájl végre kicsomagolható a tar parancs segítségével.

`root #` `tar xpvf stage3-*.tar.xz --xattrs-include='*.*' --numeric-owner -C /mnt/gentoo`

Kibontás előtt ellenőrizze az opciókat:

- `x` e **x** tract, utasítja a tar parancsot az tömörített fájl tartalmának a kibontására.
- `p` **p** reserve permissions, megőrzi az tömörített fájlban lévő fájlokon és könyvtárakon rajta lévő jogosultágokat.
- `v` **v** erbose output, részletes kimenet.
- `f` **f** ile, megadja a tar parancsnak a bemeneti tömörített fájlnevet.
- `--xattrs-include='.'` Megőrzi az tömörített fájlban tárolt összes névtér kiterjesztett attribútumait.
- `--numeric-owner` Biztosítja, hogy a tömörített fájlból kicsomagolt fájlok felhasználói és csoportazonosítói megegyezzenek a Gentoo kiadási mérnöki csapata által szándékoltakkal (még akkor is, ha a kalandosabb kedvű felhasználók nem a hivatalos Gentoo live ISO-ba bebootolt környezeteket használják a telepítési folyamat során).

Most, hogy a stage fájl ki lett csomagolva, folytassa a [kódfordítási beállításoknak a szerkesztésével](/wiki/Handbook:HPPA/Installation/Stage/hu#Configuring_compile_options "Handbook:HPPA/Installation/Stage/hu").

## Forráskódfordítási opciók beállítása

### Bevezetés

Az operációs rendszer optimalizálása érdekében be lehet állítani olyan változókat, amelyek befolyásolják a Portage, tehát a Gentoo hivatalosan támogatott szoftvercsomag kezelőjének a működését. Ezeket a változókat környezeti változóként lehet beállítani (az export használatával), de az export útján történő beállítás nem állandó.

**Note**

Technikailag a változókat ki lehet exportálni a [shell](/wiki/Shell/hu "Shell/hu") profilján vagy rc fájljain keresztül, azonban ez nem a legjobb gyakorlat az alapvető rendszergazdai feladatokhoz.

A Portage szoftvercsomag-kezelő a [make.conf](/wiki/Make.conf "Make.conf") fájlt olvassa be a futása közben, ami a fájlban tárolt értékektől függően megváltoztatja a Portage futási viselkedését. A make.conf fájl a Portage szoftvercsomag-kezelő elsődleges beállításfájljának tekinthető, ezért kezelje a tartalmát körültekintően.

**Tip**

Az összes lehetséges változó megjegyzésekkel ellátott listája megtalálható a /mnt/gentoo/usr/share/portage/config/make.conf.example fájlban. További dokumentáció a make.conf fájlról a man 5 make.conf parancs futtatásával érhető el.

A Gentoo telepítés sikerességéhez kizárólag az alábbiakban említett változókat kell beállítani.

Indítson el egy szövegszerkesztőt (ebben az útmutatóban a nano szövegszerkesztőt használjuk) a későbbiekben tárgyalt optimalizálási változók módosításához.

`root #` `nano /mnt/gentoo/etc/portage/make.conf`

A make.conf.example fájlból egyértelműen látható, hogy miként kell felépíteni a fájlt: a megjegyzésekkel ellátott sorok `#` karakterrel kezdődnek, míg a többi sor a `VÁLTOZÓNEVE="értéke"` szintaxist használva határozza meg a változókat. E változók közül több a következő szakaszban kerül tárgyalásra.

### CFLAGS és CXXFLAGS

A CFLAGS és a CXXFLAGS változó határozza meg a GCC forráskód-fordító optimalizálási-zászlóit a C és C++ fordításokat illetően. Habár ezek a zászlók itt általánosan vannak meghatározva, a maximális teljesítmény érdekében minden programhoz külön kellene optimalizálni őket. Az ok, amiért itt most így történik, az az, hogy minden program más. Azonban a minden programhoz külön meghatározás nem járható, fáradságos és hosszadalmas munka lenne, ezért ezeknek a zászlóknak a meghatározása a make.conf fájlban van történik.

A make.conf fájlban meg kell határozni azokat az optimalizációs jelölőzászlókat, amelyek általánosságban a legjobban finomhangoltabbá teszik az operációs rendszert. Ne helyezzen kísérleti beállításokat ebbe a változóba, a túlzott optimalizáció programhibákat okozhat (összeomlás, vagy ami még rosszabb, rendellenes működés).

A kézikönyv nem magyarázza el az összes lehetséges optimalizálást. Ezek megértéséhez olvassa el a [GNU online kézikönyvet](https://gcc.gnu.org/onlinedocs/) vagy a gcc infó oldalát (info gcc). A make.conf.example fájl maga is rengeteg példát és információt tartalmaz, ne feledje el ezt is elolvasni.

Az első beállítás a `-march=` vagy `-mtune=` zászló, amely megadja a célarchitektúra nevét. A lehetséges opciók a make.conf.example fájlban vannak leírva (megjegyzésként). Egy gyakran használt érték a _native_, mivel ez azt mondja a kódfordítónak, hogy válassza ki az aktuális rendszer (amin a felhasználók a Gentoo rendszert telepítik) célarchitektúráját.

A második a `-O` zászló (ez nagy O betű, nem nulla), amely meghatározza a gcc optimalizálási osztály zászlóját. Lehetséges osztályok: s (méret-optimalizált), 0 (nulla - nincs optimalizálás), 1, 2 vagy akár a 3 a sebességoptimalizálási zászlókhoz (minden osztály ugyanazokat a zászlókat tartalmazza, mint az előző, plusz néhány extrát). Az ajánlott alapértelmezés a `-O2`. A `-O3` használata rendszer szintjén problémákat okozhat, ezért azt javasoljuk, hogy maradjon a `-O2` használatánál.

Egy másik népszerű optimalizálási zászló a `-pipe` (csöveket használ ideiglenes fájlok helyett a fordítás különböző szakaszai közötti kommunikációhoz). Nincs hatása a generált programkódra, de több memóriát használ. Alacsony memóriájú rendszereken előfordulhat, hogy a gcc kódfordítása összeomlik. Ebben az esetben ne használja ezt a zászlót.

A `-fomit-frame-pointer` használata (ami nem tartja a keretmutatót egy regiszterben azoknál a funkcióknál, amelyeknek nincs szükségük rá) komoly következményekkel járhat az alkalmazások hibakeresésénél.

Amikor a CFLAGS és a CXXFLAGS változók meghatározásra kerülnek, kombinálja a különböző optimalizálási zászlókat egyetlen szövegbe. A tömörített stage fájlban található alapértelmezett értékek általában elég jók. Az alábbi csak egy példa:

CODE **Példa a CFLAGS és a CXXFLAGS változóra**

```
# Forráskódfordításhoz használt jelölőzászlók, amelyeket minden programozási nyelvhez be kell állítani.
COMMON_FLAGS="-march=2.0 -O2 -pipe"
# Mindkét változó esetében ugyanazokat a beállításokat használjuk.
CFLAGS="${COMMON_FLAGS}"
CXXFLAGS="${COMMON_FLAGS}"

```

**Tip**

Habár a [GCC optimalizálás](/wiki/GCC_optimization/hu "GCC optimization/hu") cikk több információval szolgál arról, hogy miként befolyásolhatják a különböző kódfordítási beállítások a rendszert, a [Safe CFLAGS](/wiki/Safe_CFLAGS/hu "Safe CFLAGS/hu") cikk praktikusabb kiindulópont lehet a kezdők számára a rendszerük optimalizálásának az érdekében.

### RUSTFLAGS

Számos programot ma már Rust nyelven írnak, amelynek megvan a saját optimalizálási módja. Alapértelmezés szerint a Rust minden kiadási builden 3-as szintű optimalizálást végez, hacsak egy projekt nem módosítja ezt, ezért ezt érdemes változatlanul hagyni. Az összes elérhető optimalizációs lista (codegen néven ismert), amely a Rust fordítóhoz adható, megtalálható a [https://doc.rust-lang.org/rustc/codegen-options/index.html](https://doc.rust-lang.org/rustc/codegen-options/index.html) oldalon.

A leghasznosabb optimalizálás az lenne, ha megadná a Rust nyelv esetében, hogy a forráskódfordító az Ön processzorára fordítson, az alábbi példát használva:

FILE **`/etc/portage/make.conf`** **RUSTFLAGS Example**

```
RUSTFLAGS="${RUSTFLAGS} -C target-cpu=native"

```

A támogatott processzorok listájának lekérdezéséhez Rust nyelvben futtassa a következő parancsot:

`root #` `rustc -C target-cpu=help`

Ez megmutatja, hogy mi az alapértelmezett, és melyik processzor kerül kiválasztásra a "native" opcióval.

**Note**

A fenti parancs csak asztali fokozat 3 (stage 3) tömörített fájlok esetében működik, vagy miután telepítette a [dev-lang/rust-bin](https://packages.gentoo.org/packages/dev-lang/rust-bin) vagy a [dev-lang/rust](https://packages.gentoo.org/packages/dev-lang/rust) szoftvercsomagot.

### MAKEOPTS

A MAKEOPTS változó határozza meg, hogy a processzorban hány párhuzamos kódfordítási szálnak kell futnia egy szoftvercsomag telepítésekor. A Portage 3.0.31[\[1\]](#cite_note-1)-es verziójától kezdve a Portage alapértelmezett viselkedése az, hogy ha nincs megadva ez az érték, akkor a MAKEOPTS szálak értékét a nproc által visszaadott szálak számával fogja helyettesíteni.

Továbbá, a Portage 3.0.53[\[2\]](#cite_note-2)-es verziójától kezdve úgy viselkedik, hogy ha nincs megadva, a Portage alapértelmezett viselkedése, akkor a MAKEOPTS load-average értéket a nproc által visszaadott értékre állítsa.

Jó választás lehet a kettő közül a kisebb: A processzor által használt szálak száma, vagy a teljes rendszer RAM mennyisége osztva 2 GiB-tal.

**Warning**

Nagyszámú processzorszál használata jelentősen befolyásolhatja a memóriafogyasztást. Jó ajánlás az, hogy minden megadott processzorszálhoz legalább 2 GiB RAM szükséges (például a `-j6` _legalább_ 12 GiB-ot igényel). A teljes memória elfogyásának az elkerülése érdekében csökkentse a processzorszálak számát a rendelkezésre álló memória alapján.

**Tip**

Amikor párhuzamos emerge folyamatokat ( `--jobs`) használ, a futó feladatok tényleges száma exponenciálisan növekedhet (akár a make feladatok és az emerge feladatok szorzatáig). Ezt ki lehet küszöbölni egy csak helyi gépre vonatkozó distcc beállítás futtatásával, amely korlátozza a kódfordító példányok számát gépenként.

FILE **`/etc/portage/make.conf`** **Példa a MAKEOPTS deklarációra**

```
# Ha nincs meghatározva, akkor a Portage alapértelmezett viselkedése a következő:
# - Beállítja a MAKEOPTS jobs értékét a `nproc` által visszaadott szálak számával megegyező értékre.
# - Beállítja a MAKEOPTS terhelési átlagértékét kissé a `nproc` által visszaadott szálak száma fölé, mivel ez egy tompított érték.
# Kérjük, hogy cserélje le a '4' számot az Ön rendszerének megfelelő értékre ((minimum 2GB RAM) * (processzorszálak száma), vagy hagyja alapértelmezetten üresen).
MAKEOPTS="-j4 -l5"

```

A további részletekért keresse meg a MAKEOPTS kifejezést a man 5 make.conf súgóban.

### Felkészülni, vigyázz, rajt!

Frissítse az /mnt/gentoo/etc/portage/make.conf fájlt a személyes preferenciáknak megfelelően, és mentse el a változtatásokat a fájlban (a nano szövegszerkesztőt használó felhasználók a `Ctrl` \+ `o` billentyűgombok megnyomásával mentsék el a módosítást, majd a `Ctrl` \+ `x` billentyűkombinációval lépjenek ki a szövegszerkesztőből).

## Hivatkozások

1. [↑](#cite_ref-1)[https://gitweb.gentoo.org/proj/portage.git/commit/?id=5d2af567772bb12b073f1671daea6263055cbdc2](https://gitweb.gentoo.org/proj/portage.git/commit/?id=5d2af567772bb12b073f1671daea6263055cbdc2)
2. [↑](#cite_ref-2)[https://gitweb.gentoo.org/proj/portage.git/commit/?id=de7be7f45ee54e3f789def46542919550687d15e](https://gitweb.gentoo.org/proj/portage.git/commit/?id=de7be7f45ee54e3f789def46542919550687d15e)

[← Adathordozók előkészítése](/wiki/Handbook:HPPA/Installation/Disks/hu "Previous part") [Kezdőlap](/wiki/Handbook:HPPA/hu "Handbook:HPPA/hu") [Alaprendszer telepítése →](/wiki/Handbook:HPPA/Installation/Base/hu "Next part")