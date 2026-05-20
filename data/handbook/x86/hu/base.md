# Base

Other languages:

- [Deutsch](/wiki/Handbook:X86/Installation/Base/de "Handbuch:X86/Installation/Basis (100% translated)")
- [English](/wiki/Handbook:X86/Installation/Base "Handbook:X86/Installation/Base (100% translated)")
- [español](/wiki/Handbook:X86/Installation/Base/es "Manual de Gentoo: X86/Instalación/Base (100% translated)")
- [français](/wiki/Handbook:X86/Installation/Base/fr "Handbook:X86/Installation/Base/fr (100% translated)")
- [italiano](/wiki/Handbook:X86/Installation/Base/it "Handbook:X86/Installation/Base (100% translated)")
- magyar
- [polski](/wiki/Handbook:X86/Installation/Base/pl "Handbook:X86/Installation/Base (100% translated)")
- [português do Brasil](/wiki/Handbook:X86/Installation/Base/pt-br "Manual:X86/Instalação/Base (100% translated)")
- [čeština](/wiki/Handbook:X86/Installation/Base/cs "Handbook:X86/Installation/Base/cs (50% translated)")
- [русский](/wiki/Handbook:X86/Installation/Base/ru "Handbook:X86/Installation/Base (100% translated)")
- [தமிழ்](/wiki/Handbook:X86/Installation/Base/ta "கையேடு:X86/நிறுவல்/அடிப்படை (100% translated)")
- [中文（中国大陆）‎](/wiki/Handbook:X86/Installation/Base/zh-cn "手册：X86/安装/安装基本系统 (100% translated)")
- [日本語](/wiki/Handbook:X86/Installation/Base/ja "ハンドブック:X86/インストール/ベース (100% translated)")
- [한국어](/wiki/Handbook:X86/Installation/Base/ko "Handbook:X86/Installation/Base/ko (100% translated)")

[X86 kézikönyv](/wiki/Handbook:X86/hu "Handbook:X86/hu")[A Gentoo Linux telepítése](/wiki/Handbook:X86/Full/Installation/hu "Handbook:X86/Full/Installation/hu")[A telepítésről](/wiki/Handbook:X86/Installation/About/hu "Handbook:X86/Installation/About/hu")[Telepítőképfájl kiválasztása](/wiki/Handbook:X86/Installation/Media/hu "Handbook:X86/Installation/Media/hu")[Hálózat beállítása](/wiki/Handbook:X86/Installation/Networking/hu "Handbook:X86/Installation/Networking/hu")[Adathordozók előkészítése](/wiki/Handbook:X86/Installation/Disks/hu "Handbook:X86/Installation/Disks/hu")[Stage fájl](/wiki/Handbook:X86/Installation/Stage/hu "Handbook:X86/Installation/Stage/hu")Alaprendszer telepítése[Kernel beállítása](/wiki/Handbook:X86/Installation/Kernel/hu "Handbook:X86/Installation/Kernel/hu")[Rendszer beállítása](/wiki/Handbook:X86/Installation/System/hu "Handbook:X86/Installation/System/hu")[Eszközök telepítése](/wiki/Handbook:X86/Installation/Tools/hu "Handbook:X86/Installation/Tools/hu")[Bootloader beállítása](/wiki/Handbook:X86/Installation/Bootloader/hu "Handbook:X86/Installation/Bootloader/hu")[Telepítés véglegesítése](/wiki/Handbook:X86/Installation/Finalizing/hu "Handbook:X86/Installation/Finalizing/hu")[Munka a Gentoo rendszerrel](/wiki/Handbook:X86/Full/Working/hu "Handbook:X86/Full/Working/hu")[Portage bemutatása](/wiki/Handbook:X86/Working/Portage/hu "Handbook:X86/Working/Portage/hu")[USE jelölőzászlók](/wiki/Handbook:X86/Working/USE/hu "Handbook:X86/Working/USE/hu")[Portage jellemzői](/wiki/Handbook:X86/Working/Features/hu "Handbook:X86/Working/Features/hu")[Init-szkript rendszer](/wiki/Handbook:X86/Working/Initscripts/hu "Handbook:X86/Working/Initscripts/hu")[Környezeti változók](/wiki/Handbook:X86/Working/EnvVar/hu "Handbook:X86/Working/EnvVar/hu")[Munka a Portage szoftvercsomag-kezelővel](/wiki/Handbook:X86/Full/Portage/hu "Handbook:X86/Full/Portage/hu")[Fájlok és könyvtárak](/wiki/Handbook:X86/Portage/Files/hu "Handbook:X86/Portage/Files/hu")[Változók](/wiki/Handbook:X86/Portage/Variables/hu "Handbook:X86/Portage/Variables/hu")[Szoftverágak keverése](/wiki/Handbook:X86/Portage/Branches/hu "Handbook:X86/Portage/Branches/hu")[További eszközök](/wiki/Handbook:X86/Portage/Tools/hu "Handbook:X86/Portage/Tools/hu")[Egyéni szoftvercsomag-tárolók](/wiki/Handbook:X86/Portage/CustomTree/hu "Handbook:X86/Portage/CustomTree/hu")[Fejlett funkciók](/wiki/Handbook:X86/Portage/Advanced/hu "Handbook:X86/Portage/Advanced/hu")[Hálózat beállítása OpenRC init-rendszeren](/wiki/Handbook:X86/Full/Networking/hu "Handbook:X86/Full/Networking/hu")[Munka elkezdése](/wiki/Handbook:X86/Networking/Introduction/hu "Handbook:X86/Networking/Introduction/hu")[Fejlett beállítások](/wiki/Handbook:X86/Networking/Advanced/hu "Handbook:X86/Networking/Advanced/hu")[Moduláris hálózat](/wiki/Handbook:X86/Networking/Modular/hu "Handbook:X86/Networking/Modular/hu")[Vezeték nélküli (Wi-Fi)](/wiki/Handbook:X86/Networking/Wireless/hu "Handbook:X86/Networking/Wireless/hu")[Funkcionalitás hozzáadása](/wiki/Handbook:X86/Networking/Extending/hu "Handbook:X86/Networking/Extending/hu")[Dinamikus menedzsment](/wiki/Handbook:X86/Networking/Dynamic/hu "Handbook:X86/Networking/Dynamic/hu")

## Contents

- [1Chrootolás](#Chrootol.C3.A1s)
  - [1.1DNS információ másolása](#DNS_inform.C3.A1ci.C3.B3_m.C3.A1sol.C3.A1sa)
  - [1.2Szükséges fájlrendszerek felcsatolása](#Sz.C3.BCks.C3.A9ges_f.C3.A1jlrendszerek_felcsatol.C3.A1sa)
  - [1.3Belépés az új környezetbe](#Bel.C3.A9p.C3.A9s_az_.C3.BAj_k.C3.B6rnyezetbe)
  - [1.4Felkészülés a bootloaderre](#Felk.C3.A9sz.C3.BCl.C3.A9s_a_bootloaderre)
    - [1.4.1DOS/Örökölt BIOS rendszerek](#DOS.2F.C3.96r.C3.B6k.C3.B6lt_BIOS_rendszerek)
- [2Portage szoftvercsomag-kezelő beállítása](#Portage_szoftvercsomag-kezel.C5.91_be.C3.A1ll.C3.ADt.C3.A1sa)
  - [2.1Gentoo ebuild szoftvercsomag-tároló pillanatképének telepítése a weben keresztül](#Gentoo_ebuild_szoftvercsomag-t.C3.A1rol.C3.B3_pillanatk.C3.A9p.C3.A9nek_telep.C3.ADt.C3.A9se_a_weben_kereszt.C3.BCl)
  - [2.2Opcionális: Tükörszerverek kiválasztása](#Opcion.C3.A1lis:_T.C3.BCk.C3.B6rszerverek_kiv.C3.A1laszt.C3.A1sa)
    - [2.2.1Optional: rsync mirrors](#Optional:_rsync_mirrors)
  - [2.3Opcionális: Gentoo ebuild szoftvercsomag-tároló frissítése](#Opcion.C3.A1lis:_Gentoo_ebuild_szoftvercsomag-t.C3.A1rol.C3.B3_friss.C3.ADt.C3.A9se)
  - [2.4Hírelemek olvasása](#H.C3.ADrelemek_olvas.C3.A1sa)
  - [2.5Megfelelő profil kiválasztása](#Megfelel.C5.91_profil_kiv.C3.A1laszt.C3.A1sa)
  - [2.6Opcionális: Bináris szoftvercsomag-tároló számítógép hozzáadása](#Opcion.C3.A1lis:_Bin.C3.A1ris_szoftvercsomag-t.C3.A1rol.C3.B3_sz.C3.A1m.C3.ADt.C3.B3g.C3.A9p_hozz.C3.A1ad.C3.A1sa)
    - [2.6.1Szoftvercsomag-tároló beállítása](#Szoftvercsomag-t.C3.A1rol.C3.B3_be.C3.A1ll.C3.ADt.C3.A1sa)
    - [2.6.2Bináris szoftvercsomagok telepítése](#Bin.C3.A1ris_szoftvercsomagok_telep.C3.ADt.C3.A9se)
  - [2.7Opcionális: USE változó beállítása](#Opcion.C3.A1lis:_USE_v.C3.A1ltoz.C3.B3_be.C3.A1ll.C3.ADt.C3.A1sa)
    - [2.7.1CPU\_FLAGS\_\*](#CPU_FLAGS_.2A)
    - [2.7.2VIDEO\_CARDS](#VIDEO_CARDS)
  - [2.8Opcionális: ACCEPT\_LICENSE változó beállítása](#Opcion.C3.A1lis:_ACCEPT_LICENSE_v.C3.A1ltoz.C3.B3_be.C3.A1ll.C3.ADt.C3.A1sa)
  - [2.9Opcionális: @world készlet frissítése](#Opcion.C3.A1lis:_.40world_k.C3.A9szlet_friss.C3.ADt.C3.A9se)
    - [2.9.1Elévült szoftvercsomagok eltávolítása](#El.C3.A9v.C3.BClt_szoftvercsomagok_elt.C3.A1vol.C3.ADt.C3.A1sa)
- [3Időzóna](#Id.C5.91z.C3.B3na)
- [4Nyelvterület-beállítások](#Nyelvter.C3.BClet-be.C3.A1ll.C3.ADt.C3.A1sok)
  - [4.1Nyelvterület-beállítások létrehozása](#Nyelvter.C3.BClet-be.C3.A1ll.C3.ADt.C3.A1sok_l.C3.A9trehoz.C3.A1sa)
  - [4.2Nyelvterület-beállítások kiválasztása](#Nyelvter.C3.BClet-be.C3.A1ll.C3.ADt.C3.A1sok_kiv.C3.A1laszt.C3.A1sa)
- [5Hivatkozások](#Hivatkoz.C3.A1sok)

## Chrootolás

### DNS információ másolása

Még egy dolog van hátra, mielőtt belépnénk az új környezetbe, és ez a DNS információk átmásolása az /etc/resolv.conf fájlba. Ezt azért kell megtenni, hogy a hálózat továbbra is működjön az új környezetbe való belépés után is. Az /etc/resolv.conf fájl tartalmazza a hálózati névszervereket.

Az információ átmásolásakor ajánlott a cp parancshoz hozzáadni a `--dereference` opciót. Ez biztosítja, hogy ha az /etc/resolv.conf fájl egy szimbolikus link, akkor a link célfájlja legyen másolva a szimbolikus link helyett. Ellenkező esetben az új környezetben a szimbolikus link egy nem létező fájlra mutatna (mivel a link célja valószínűleg nem elérhető az új környezetben).

`root #` `cp --dereference /etc/resolv.conf /mnt/gentoo/etc/`

### Szükséges fájlrendszerek felcsatolása

**Tip**

Ha az eredeti Gentoo telepítési ISO képfájlt használjuk a telepítéskor, akkor ezt a lépést egyszerűen az alábbi paranccsal helyettesíthetjük: arch-chroot /mnt/gentoo.

Pár pillanat múlva a Linux gyökér könyvtára az új helyre kerül áthelyezésre.

A következő fájlrendszereket kell elérhetővé tenni:

- /proc/ egy ál-fájlrendszer. Úgy néz ki, mint a szokásos fájlok, de a Linux kernel generálja őket menet közben.
- /sys/ egy ál-fájlrendszer, hasonló a /proc/ fájlrendszerhez, amelyet egykor a /proc/ helyettesítésére szántak, és sokkal strukturáltabb, mint a /proc/.
- /dev/ egy normál fájlrendszer, amely tartalmazza az összes eszközt. Részben a Linux eszközkezelő (általában az udev) kezeli.
- /run/ egy ideiglenes fájlrendszer, amelyet a futásidőben generált fájlok, például PID fájlok vagy zárolások tárolására használnak.

A /proc/ hely a /mnt/gentoo/proc/ helyre lesz felcsatolva, míg a többi fájlrendszer kötött csatolással (bind-mounted) lesz elérhető. Ez utóbbi azt jelenti, hogy például a /mnt/gentoo/sys/ valójában a /sys/ _lesz_ (ez csak egy második bejegyzési pont ugyanahhoz a fájlrendszerhez), míg a /mnt/gentoo/proc/ egy új csatolás (úgymond egy új példány) a fájlrendszerből.

`root #` `mount --types proc /proc /mnt/gentoo/proc
`

`root #` `mount --rbind /sys /mnt/gentoo/sys
`

`root #` `mount --make-rslave /mnt/gentoo/sys
`

`root #` `mount --rbind /dev /mnt/gentoo/dev
`

`root #` `mount --make-rslave /mnt/gentoo/dev
`

`root #` `mount --bind /run /mnt/gentoo/run
`

`root #` `mount --make-slave /mnt/gentoo/run
`

**Note**

A `--make-rslave` műveletek szükségesek a systemd támogatásához a telepítés későbbi szakaszában.

**Warning**

Ha nem Gentoo telepítési ISO képfájlt használunk a telepítéskor, akkor ez nem biztos, hogy elegendő. Néhány disztribúció a /dev/shm helyet szimbolikus linkként hozza létre a /run/shm/ helyre, amely a chroot után érvénytelenné válik. Ennek orvoslására a /dev/shm/ helyet helyesen beállított tmpfs csatolással lehet előre megoldani:

`root #` `test -L /dev/shm && rm /dev/shm && mkdir /dev/shm
`

`root #` `mount --types tmpfs --options nosuid,nodev,noexec shm /dev/shm
`

Kérjük, biztosítsa azt is, hogy a 1777 mód be legyen állítva:

`root #` `chmod 1777 /dev/shm /run/shm`

### Belépés az új környezetbe

Most, hogy az összes partíció inicializálva lett és az alapkörnyezet telepítve van, itt az ideje belépni abba (a feltelepített rendszerbe) a chroot használatával. Ez azt jelenti, hogy a munkamenet megváltoztatja a gyökérkönyvtár helyét (a legfelső szintű helyet, amely elérhető). A jelenlegi telepítőkörnyezetből (telepítési USB pendrive, DVD vagy más telepítési médium környezetéből) átlépünk a feltelepített rendszerre (nevezetesen az inicializált partíciókra). Innen származik a név, _change root_ (gyökér megváltoztatása) vagyis _chroot_.

Ez a chrootolás három lépésben történik:

1. A gyökérkönyvtár helyének módosítása / (a telepítési médiumon) /mnt/gentoo/ (a partíciókon) helyre chroot vagy arch-chroot használatával, ha elérhető.
2. Néhány beállítás (amelyek az /etc/profile fájlban találhatók) betöltése a memóriába a source parancs segítségével.
3. Az elsődleges prompt módosítása, hogy emlékeztessen bennünket arra, hogy ez a munkamenet egy chroot környezetben zajlik.

`root #` `chroot /mnt/gentoo /bin/bash
`

`root #` `source /etc/profile
`

`root #` `export PS1="(chroot) ${PS1}"`

Ettől a ponttól kezdve minden végrehajtott művelet közvetlenül az újonnan feltelepített Gentoo Linux környezetben történik.

**Tip**

Ha a Gentoo telepítése ezen a ponton bárhol megszakad, akkor ettől a lépéstől _lehetséges_ a telepítés _folytatása_. Nincs szükség az adathordozók újraparticionálására! Egyszerűen [csatolja fel a gyökérpartíciót](/wiki/Handbook:X86/Installation/Disks/hu#Mounting_the_root_partition "Handbook:X86/Installation/Disks/hu"), és hajtsa végre a fentiekben leírt lépéseket, kezdve a [DNS információ másolásával](/wiki/Handbook:X86/Installation/Base/hu#Copy_DNS_info "Handbook:X86/Installation/Base/hu"), hogy újra belépjen a munkakörnyezetbe. Ez hasznos lehet a bootloader problémák javításához is. További információk a [chroot](/wiki/Chroot/hu "Chroot/hu") cikkben találhatók.

### Felkészülés a bootloaderre

Most, hogy beléptünk az új környezetbe, szükséges előkészíteni az új környezetet a bootloader számára. Fontos lesz, hogy a megfelelő partíció fel legyen csatolva, amikor eljön az ideje a bootloader telepítésének.

#### DOS/Örökölt BIOS rendszerek

DOS/Örökölt BIOS rendszerek esetén a bootloader a /boot könyvtárba lesz telepítve, ezért a következőképpen kell felcsatolni:

`root #` `mount /dev/sda1 /boot`

## Portage szoftvercsomag-kezelő beállítása

### Gentoo ebuild szoftvercsomag-tároló pillanatképének telepítése a weben keresztül

A következő lépés a Gentoo ebuild szoftvercsomag-tároló pillanatképének a telepítése. Ez a pillanatkép fájlok gyűjteményét tartalmazza, amelyek tájékoztatják a Portage szoftvercsomag-kezelőt a telepíthető szoftverek neveiről, amelyeket a rendszergazda kiválaszthat, szoftvercsomag- vagy profil-specifikus hírelemekről stb.

Az emerge-webrsync parancs használata ajánlott azok számára, akik korlátozó tűzfalak mögött vannak (ez a parancs HTTP/FTP protokollokat használ a pillanatkép letöltéséhez), és ezáltal hálózati sávszélességet takarít meg. Azok az olvasók, akiknek nincs hálózati vagy sávszélesség-korlátozásuk, nyugodtan átugorhatják a következő szakaszt.

Ez a parancs letölti a legfrissebb pillanatképet (amelyet naponta adnak ki) a Gentoo egyik tükörszerveréről és telepíti azt a rendszerünkre:

`root #` `emerge-webrsync`

**Note**

A művelet során az emerge-webrsync panaszkodhat egy hiányzó /var/db/repos/gentoo/ könyvtárra. Ez várható, és nincs miért aggódni, a parancs létrehozza önmagától a hiányzó könyvtárat.

Ettől a ponttól kezdve a Portage szoftvercsomag-kezelő jelezheti, hogy ajánlott bizonyos frissítéseket végrehajtani. Ennek oka, hogy a stage fájlon keresztül telepített rendszer-szoftvercsomagokhoz újabb verziók állhatnak rendelkezésre, és a Portage szoftvercsomag-kezelő most már tisztában van az új szoftvercsomagokkal a szoftvercsomag-tároló letöltött pillanatképének köszönhetően. A szoftvercsomag-frissítések most biztonságosan figyelmen kívül hagyhatóak, a frissítések elhalaszthatók a Gentoo telepítésének befejezése utánra.

### Opcionális: Tükörszerverek kiválasztása

A forráskód gyors letöltése érdekében ajánlott egy gyors, földrajzilag közeli tükörszervert választani. A Portage szoftvercsomag-kezelő a make.conf fájlban keresi a GENTOO\_MIRRORS változót, és a benne felsorolt tükörszervereket használja. Lehetséges a Gentoo tükörszerverlista böngészése, és a rendszer fizikai helyéhez közeli tükörszerver (vagy több tükörszerver) keresése (mivel ezek általában a leggyorsabbak).

Az úgynevezett mirrorselect parancs egy szép szöveges felületet biztosít, amely segítségével gyorsabban kereshetőek és választhatóak ki a megfelelő tükörszerverek. Csak navigáljon a választott tükörszerverekre, és nyomja meg a `Szóköz` billentyűgombot egy vagy több tükörszerver kiválasztásához.

`root #` `emerge --ask --verbose --oneshot app-portage/mirrorselect`

`root #` `mirrorselect -i -o >> /etc/portage/make.conf`

Alternatív megoldásként egy aktív tükörszerverek listája [elérhető online](https://www.gentoo.org/downloads/mirrors/).

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

### Opcionális: Gentoo ebuild szoftvercsomag-tároló frissítése

A Gentoo ebuild szoftvercsomag-tárolót lehetséges a legújabb verzióra frissíteni. Az előző emerge-webrsync parancs egy nagyon friss pillanatképet telepített (általában legfeljebb 24 órás), így ez a lépés mindenképpen opcionális.

Tegyük fel, hogy szükség van a legfrissebb (mondjuk az 1 órán belüli) szoftvercsomag-frissítésekre, akkor használjuk az emerge --sync parancsot. Ez a parancs az rsync protokollt használja a Gentoo ebuild szoftvercsomag-tároló (amelyet korábban az emerge-webrsync segítségével töltöttek le) frissítéséhez a legújabb állapotra.

`root #` `emerge --sync`

Lassú parancssorokon, például bizonyos keretpuffereken vagy soros parancssorokon, ajánlott a `--quiet` opció használata a folyamat felgyorsítása érdekében:

`root #` `emerge --sync --quiet`

### Hírelemek olvasása

Amikor a Gentoo ebuild szoftvercsomag-tároló szinkronizálva van, akkor a Portage hasonló információs üzeneteket jeleníthet meg:

` * FONTOS: A 'gentoo' szoftvercsomag-tárolóban szükséges 2 hírelem elolvasása.
`

` * Használja az eselect news parancsot a hírelemek olvasásához.
`

A hírelemeket azért hozták létre, hogy kommunikációs eszközként szolgáljanak a kritikus üzenetek felhasználókhoz való eljuttatására a Gentoo ebuild szoftvercsomag-tárolón keresztül. Ezek kezeléséhez használja az eselect news parancsot. Az eselect alkalmazás egy Gentoo-specifikus eszköz, amely egy közös kezelőfelületet biztosít a rendszeradminisztrációhoz. Ebben az esetben az eselect azzal van megbízva, hogy használja a `news` modulját.

A `news` modul esetében három műveletet használnak leggyakrabban:

- A `list` paranccsal megjelenik az elérhető hírelemek áttekintése.
- A `read` paranccsal elolvashatók a hírelemek.
- A `purge` paranccsal a hírelemek eltávolíthatók, miután elolvasták őket, és többé nem lesznek újraolvasva.

`root #` `eselect news list
`

`root #` `eselect news read`

További információ a hírolvasóról elérhető a kézikönyv oldalán:

`root #` `man news.eselect`

### Megfelelő profil kiválasztása

**Tip**

A desktop profilok nem kizárólag az _asztali környezetekhez_ vannak. Alkalmasak nagyon kis erőforrást igénylő ablakkezelőkhöz is, mint például az i3 vagy az sway.

Egy _profile_ tulajdonképpen egy építőelem bármely Gentoo operációs rendszer számára. Nemcsak az alapértelmezett értékeket határozza meg a USE, CFLAGS és más fontos változók számára, hanem az operációs rendszer egy bizonyos szoftvercsomagverzió-tartományhoz is rögzül. Ezeket a beállításokat a Gentoo Portage fejlesztői tartják karban.

Annak érdekében, hogy megnézze, vajon milyen profilt használ jelenleg az operációs rendszer, futtassa az eselect parancsot a `profile` modullal:

**Tip**

On an install media without a scroll-able terminal, `eselect profile list | less` can provide an easy way to list all available profiles while providing the ability to scroll.

`root #` `eselect profile list`

```
Available profile symlink targets:
  [1]   default/linux/x86/23.0 *
  [2]   default/linux/x86/23.0/desktop
  [3]   default/linux/x86/23.0/desktop/gnome
  [4]   default/linux/x86/23.0/desktop/kde

```

**Note**

A parancs kimenete csak egy példa, és idővel változik.

**Note**

Ha a **systemd** init-rendszert szeretné használni, akkor válasszon egy olyan profilt, amelynek nevében szerepel a "systemd" kifejezés.

Bizonyos architektúrákhoz elérhetők olyan desktop alprofilok is, amelyek tartalmazzák a gyakran szükséges szoftvercsomagokat az asztali élmény érdekében.

**Warning**

A profilfrissítéseket nem szabad félvállról venni. Az induló profil kiválasztásakor használja azt a profilt, amely **ugyanahhoz** a verzióhoz tartozik, mint amit eredetileg a fokozat (stage) fájl használt (pl. 23.0). Minden új profilverziót egy hírelemben jelentenek be, amely tartalmazza az átállási útmutatásokat. Legyen biztos benne, hogy alaposan követi az átállási útmutatásokat, mielőtt egy újabb profilra váltana.

Az x86 architektúra elérhető profiljainak a megtekintése után a felhasználók kiválaszthatnak egy másik profilt a rendszerhez:

`root #` `eselect profile set 2`

**Note**

A `developer` alprofil kifejezetten a Gentoo Linux fejlesztésére szolgál, és nem hétköznapi felhasználók számára készült.

### Opcionális: Bináris szoftvercsomag-tároló számítógép hozzáadása

2023 decembere óta a [Release Engineering csapata](/wiki/Project:RelEng "Project:RelEng") egy [hivatalos bináris szoftvercsomag-tároló számítógépet](/wiki/Binary_package_quickstart "Binary package quickstart") (bináris szoftvercsomag-host, köznyelvben "binhost" számítógépet) kínál a közösség számára a bináris szoftvercsomagok (binpkgs) letöltésére és telepítésére.[\[1\]](#cite_note-1)

Egy bináris szoftvercsomag-tároló számítógép hozzáadása lehetővé teszi a Portage szoftvercsomag-kezelő számára a kriptográfiai aláírással ellátott, forráskódból binárisra lefordított szoftvercsomagok telepítését. Sok esetben a bináris szoftvercsomag-tároló számítógép hozzáadása jelentősen csökkenti a szoftvercsomagok telepítéséhez szükséges átlagos időt, és nagy előnyt jelent, amikor Gentoo operációs rendszert futtatunk régebbi, lassabb vagy alacsony teljesítményű számítógépeken.

#### Szoftvercsomag-tároló beállítása

A binhost szoftvercsomag-tároló beállítása a Portage /etc/portage/binrepos.conf/ könyvtárában található, amely hasonlóan működik a [Gentoo ebuild tároló](/wiki/Handbook:X86/Installation/Base/hu#Gentoo_ebuild_repository "Handbook:X86/Installation/Base/hu") szakaszában említett beállításhoz.

Bináris számítógép meghatározásakor két fontos szempontot kell figyelembe venni:

1. A `sync-uri` értékében szereplő architektúra és profil célpontok _valóban_ számítanak, és igazodniuk kell a megfelelő számítógép architektúrához (ebben az esetben **x86**) és a [Megfelelő profil kiválasztása](/wiki/Handbook:X86/Installation/Base/hu#Choosing_the_right_profile "Handbook:X86/Installation/Base/hu") szakaszban kiválasztott rendszerprofilhoz.
2. Egy gyors, földrajzilag közeli tükörszerver kiválasztása általában lerövidíti a letöltési időt. Tekintse át a mirrorselect eszközt, amely az [Opcionális: Tükrök kiválasztása](/wiki/Handbook:X86/Installation/Base/hu#Gentoo_ebuild_repository "Handbook:X86/Installation/Base/hu") szakaszban található, vagy tekintse át az [online tükörszerverek listáját](https://www.gentoo.org/downloads/mirrors/), ahol URL értékeket találhat.


FILE **`/etc/portage/binrepos.conf/gentoobinhost.conf`** **CDN alapú bináris szoftvercsomag-hoszt példája**

```
[binhost]
priority = 9999
sync-uri = https://distfiles.gentoo.org/releases/<arch>/binpackages/<profile>/x86-64/

```

1. Introduced in portage-3.0.74 for per-repo verification choices

verify-signature = true

1. Default value with >=portage-3.0.77

location = /var/cache/binhost/gentoo
}}

#### Bináris szoftvercsomagok telepítése

A Portage szoftvercsomag-kezelő alapértelmezés szerint a szoftvercsomagokat forráskódból szokta lefordítani. Az alábbi módokon utasíthatja a bináris csomagok használatára:

1. A `--getbinpkg` opció megadható az emerge parancs meghívásakor. Ez a bináris szoftvercsomag telepítésének módszere hasznos, ha csak egy adott bináris szoftvercsomagot szeretne telepíteni.
2. A rendszer alapértelmezett beállításának megváltoztatása a Portage FEATURES változóján keresztül, amely a /etc/portage/make.conf fájlon keresztül érhető el. Ennek a beállításmódosításnak az alkalmazása miatt a Portage a bináris szoftvercsomag hosztot fogja lekérdezni a kért csomag(ok)ért, és helyben fog lefordítani, ha nem talál eredményeket.

Például, hogy a Portage mindig telepítse az elérhető bináris csomagokat:

FILE **`/etc/portage/make.conf`** **Portage szoftvercsomag-kezelőt úgy állítja be, hogy alapértelmezés szerint bináris szoftvercsomagokat használjon**

```
# A getbinpkg hozzáfűzése a FEATURES változón belüli értékek listájához.
FEATURES="${FEATURES} getbinpkg"
# Aláírásokat igényeljen.
FEATURES="${FEATURES} binpkg-request-signature"

```

Kérjük, futtassa a getuto parancsot is, hogy a Portage beállítsa a szükséges kulcstartót a hitelesítéshez:

`root #` `getuto`

A Portage szoftvercsomag-kezelő további funkcióit a kézikönyv [következő fejezetében](/wiki/Handbook:X86/Working/Features/hu#Portage_features "Handbook:X86/Working/Features/hu") tárgyaljuk.

### Opcionális: USE változó beállítása

A USE az egyik legerőteljesebb változó, amelyet a Gentoo operációs rendszer biztosít a felhasználói számára. Számos program lefordítható választható támogatással bizonyos elemek támogatásával vagy azon elemek nélkül. Például néhány program lefordítható GTK+ támogatással vagy Qt támogatással. Mások lefordíthatók SSL támogatással vagy anélkül. Egyes programok lefordíthatók framebuffer támogatással (svgalib) X11 támogatás helyett (X-server).

A legtöbb disztribúció a szoftvercsomagjait minél szélesebb támogatottsággal fordítja le, ami növeli a programok méretét és az indítási időt, nem is beszélve a rengeteg szoftverfüggőségről. A Gentoo operációs rendszerrel a felhasználók meghatározhatják, hogy egy szoftvercsomagot milyen opciókkal kell lefordítani bináris futtatható kódra. Itt lép be a képbe a USE változó.

A USE változóban a felhasználók kulcsszavakat határoznak meg, amelyek a fordítási opciókra vannak leképezve. Például a `ssl` lefordítja az SSL támogatást a programokba, amelyek támogatják azt. A `-X` eltávolítja az X-server támogatást (figyelje meg a mínusz jelet az elején). A `gnome gtk -kde -qt5` lefordítja a programokat GNOME (és GTK+) támogatással, és nem KDE (és Qt) támogatással, teljesen a GNOME-hoz igazítva a rendszert (ha az architektúra támogatja).

Az operációs rendszer által használt Gentoo profil make.defaults fájljaiba kerülnek az alapértelmezett USE beállítások. A Gentoo egy összetett öröklési rendszert használ a rendszerprofilokhoz, amelyet az telepítési folyamat során nem részletezünk. A legkönnyebb módja annak, hogy ellenőrizze az aktuálisan aktív USE beállításokat, az a emerge --info parancs futtatása és a sor kiválasztása, amely a USE: kifejezéssel kezdődik.

`root #` `emerge --info | grep ^USE`

```
USE="X acl alsa amd64 berkdb bindist bzip2 cli cracklib crypt cxx dri ..."

```

**Note**

A fenti példa lerövidített, a tényleges USE értékek listája sokkal, sokkal nagyobb.

A rendelkezésre álló USE jelölőzászlók teljes leírása megtalálható a rendszerben a /var/db/repos/gentoo/profiles/use.desc fájlban.

`root #` `less /var/db/repos/gentoo/profiles/use.desc`

A less parancsban a görgetés a `↑` és `↓` billentyűkkel végezhető el, és a kilépés a `q` megnyomásával történik.

Példaként bemutatjuk egy KDE alapú rendszer USE változó beállításait DVD, ALSA és CD rögzítési támogatással:

`root #` `nano /etc/portage/make.conf`

FILE **`/etc/portage/make.conf`** **Jelölőzászlók engedélyezése egy KDE/Plasma alapú rendszerhez DVD, ALSA és CD rögzítési támogatással**

```
USE="-gtk -gnome qt5 kde dvd alsa cdr"

```

Amikor egy USE érték meg van határozva a /etc/portage/make.conf fájlban, az _hozzáadódik_ az operációs rendszer USE jelölőzászlólistájához. A USE jelölőzászlókat globálisan _el lehet távolítani_ úgy, hogy egy `-` mínusz jelet teszünk az érték elé a listában. Például az X grafikus környezetek támogatásának letiltásához a `-X` értéket lehet beállítani:

FILE **`/etc/portage/make.conf`** **Az alapértelmezett USE jelölőzászlók figyelmen kívül hagyása**

```
USE="-X acl alsa"

```

**Warning**

Bár lehetséges, a `-*` beállítása (amely minden USE értéket letilt, kivéve azokat, amelyek a make.conf fájlban vannak megadva) _erősen_ ellenzett és egyáltalán nem bölcs dolog. Az ebuild fejlesztők bizonyos alapértelmezett USE jelölőzászló értékeket választanak az ebuild-ekben annak érdekében, hogy elkerüljék a konfliktusokat, növeljék a biztonságot, és elkerüljék a hibákat, valamint egyéb okok miatt. Az összes USE jelölőzászló letiltása megszünteti az alapértelmezett viselkedést, és súlyos problémákat okozhat.

#### CPU\_FLAGS\_\*

Néhány architektúrában (beleértve az AMD64/X86, ARM, PPC) van egy [USE\_EXPAND](/wiki/USE_EXPAND "USE EXPAND") változó, amelyet [CPU\_FLAGS\_<ARCH>](/wiki/CPU_FLAGS_X86 "CPU FLAGS X86")-nek neveznek, ahol az <ARCH> helyére a releváns rendszer-architektúra neve kerül.

**Important**

Ne legyen összezavarodva! Az **AMD64** és az **X86** rendszerek osztoznak néhány közös architektúrán, így az **AMD64** rendszerek megfelelő változó neve a CPU\_FLAGS\_X86.

Ezt arra használják, hogy a buildet specifikus assembly kódba vagy egyéb, általában kézzel írt vagy más egyedi utasításokkal fordítsák le, és **nem** azonos azzal, hogy a kódfordítót arra kérjük, hogy optimalizált kódot generáljon egy adott processzor jellemzőhöz (pl. `-march=`).

A felhasználóknak ezt a változót be kell állítaniuk az általuk kívánt COMMON\_FLAGS beállítása mellett.

Néhány lépés szükséges ennek beállításához:

`root #` `emerge --ask --oneshot app-portage/cpuid2cpuflags`

Ellenőrizze a kimenetet manuálisan, ha kíváncsi:

`root #` `cpuid2cpuflags`

Majd másolja a kimenetet a package.use fájlba:

`root #` `echo "*/* $(cpuid2cpuflags)" > /etc/portage/package.use/00cpu-flags`

#### VIDEO\_CARDS

By default a profile already sets a few video cards. For many reasons this is not ideal, but a useful safety net.

To configure the system correctly the user needs to first unset the preset video cards with `VIDEO_CARDS: -*` then set the correct card for that system.

FILE **`/etc/portage/package.use/00video_cards`** **Example**

```
*/* VIDEO_CARDS: -* <GPU DRIVER NAME>

```

Az alábbi táblázat felhasználható a különböző videokártya-típusok és a hozzájuk tartozó `VIDEO_CARDS` értékek egyszerű összehasonlítására.

Számítógép
Diszkrét videokártya
VIDEO\_CARDS
Intel x86SemmiTekintse meg a [Intel#Feature support](/wiki/Intel/hu#Feature_support "Intel/hu") leírást.
x86/ARMNvidia`nvidia`BármilyenNvidia, kivéve Maxwell, Pascal és Volta`nouveau`BármilyenAMD, Sea Islands óta`amdgpu radeonsi`BármilyenATI és idősebb AMDTekintse meg a [radeon#Feature support](/wiki/Radeon#Feature_support "Radeon") leírást.
BármilyenIntel`intel`Raspberry PiN/A`vc4`QEMU/KVMBármilyen`virgl`WSLBármilyen`d3d12`

Az alábbiakban egy helyesen beállított package.use példa található a _VIDEO\_CARDS_ változóhoz. Helyettesítse be a használni kívánt illesztőprogram(ok) nevét.

FILE **`/etc/portage/package.use/00video_cards`**

```
*/* VIDEO_CARDS: amdgpu radeonsi

```

FILE **`/etc/portage/package.use/00video_cards`** **Intel example**

```
*/* VIDEO_CARDS: -* intel

```

FILE **`/etc/portage/package.use/00video_cards`** **Nvidia example**

```
*/* VIDEO_CARDS: -* nvidia

```

A különféle grafikus processzorokra vonatkozó részletek megtalálhatók az [AMDGPU](/wiki/AMDGPU/hu "AMDGPU/hu"), [Intel](/wiki/Intel/hu "Intel/hu"), [Nouveau (Open Source)](/wiki/Nouveau/hu "Nouveau/hu"), vagy [NVIDIA (Proprietary)](/wiki/NVIDIA/hu "NVIDIA/hu") cikkekben.

### Opcionális: ACCEPT\_LICENSE változó beállítása

A Gentoo Linux Enhancement Proposal 23 (GLEP 23) bevezetésével egy mechanizmust hoztak létre, amely lehetővé teszi a rendszergazdák számára, hogy "szabályozzák a telepített szoftvereket a licencük szempontjából... Néhányan olyan rendszert akarnak, amely mentes minden olyan szoftvertől, amelyet az OSI nem hagyott jóvá; mások egyszerűen kíváncsiak arra, hogy milyen licenceket fogadnak el implicit módon."[\[2\]](#cite_note-2) Az a szándék vezérelte őket, hogy részletesebb ellenőrzést gyakoroljanak a Gentoo rendszeren futó szoftverek felett, így született meg az ACCEPT\_LICENSE változó.

A telepítési folyamat során a Portage szoftvercsomag-kezelő figyelembe veszi az ACCEPT\_LICENSE változóban beállított értékeket annak megállapításához, hogy a kért szoftvercsomagok megfelelnek-e a rendszergazda által elfogadhatónak ítélt licencnek. Itt rejlik egy probléma: a Gentoo ebuild szoftvercsomag-tároló _több ezer_ ebuildet tartalmaz, ami [_több száz_ különálló szoftverlicencet](https://gitweb.gentoo.org/repo/gentoo.git/tree/licenses) eredményez... Ez azt jelenti, hogy a rendszergazdának minden egyes új szoftverlicencet egyenként kell jóváhagynia? Szerencsére nem. A GLEP 23 egy megoldást is vázol erre a problémára, egy licenc csoportoknak nevezett koncepció formájában.

A rendszer-adminisztráció kényelmét szolgálja, hogy a jogilag hasonló szoftverlicenceket egybecsomagolták – mindegyiket hasonló típusának megfelelően. A licenccsoport-definíciók [megtekinthetők](https://gitweb.gentoo.org/repo/gentoo.git/tree/profiles/license_groups), és a [Gentoo Licenses projekt](/wiki/Project:Licenses "Project:Licenses") kezeli őket. Az egyedi licencek nincsenek egybecsomagolva. A licenccsoportokat szintaktikailag egy `@` szimbólum előzi meg, ami lehetővé teszi azok könnyű megkülönböztetését az ACCEPT\_LICENSE változóban.

Néhány gyakori licenc csoport közé tartozik:

Egy lista a szoftverlicencekről, amelyek a típusaik szerint csoportosítva vannak.
NévLeírás
`@GPL-COMPATIBLE`A Free Software Foundation által jóváhagyott GPL-kompatibilis licencek [\[a\_license 1\]](#cite_note-3).
`@FSF-APPROVED`Az FSF által jóváhagyott ingyenes szoftverlicencek. (Tartalmazza a `@GPL-COMPATIBLE` címkét).
`@OSI-APPROVED`Az Open Source Initiative által jóváhagyott licencek [\[a\_license 2\]](#cite_note-4).
`@MISC-FREE`Egyéb licencek, amelyek valószínűleg szabad szoftverek, azaz követik a Szabad Szoftver Definíciót[\[a\_license 3\]](#cite_note-5), de nincs jóváhagyva sem az FSF, sem az OSI által.
`@FREE-SOFTWARE`Összevonja az `@FSF-APPROVED`, `@OSI-APPROVED`, és `@MISC-FREE` címkéket.
`@FSF-APPROVED-OTHER`Az FSF által jóváhagyott licencek a "szabad dokumentációhoz" és a "gyakorlati felhasználású munkákhoz a szoftvereken és dokumentációkon kívül". (Beleértve a betűtípusokat is).
`@MISC-FREE-DOCS`Egyéb licencek szabad dokumentumokhoz és más művekhez (beleértve a betűtípusokat is), amelyek követik a szabad definíciót[\[a\_license 4\]](#cite_note-6), de NEM szerepelnek az @FSF-APPROVED-OTHER listán.
`@FREE-DOCUMENTS`Összevonja az `@FSF-APPROVED-OTHER` és `@MISC-FREE-DOCS` címkéket.
`@FREE`A szabad használat, megosztás, módosítás és módosítások megosztásának szabadságával rendelkező összes licenc metaszettje. Összevonja a `@FREE-SOFTWARE` és `@FREE-DOCUMENTS` címkéket.
`@BINARY-REDISTRIBUTABLE`Licencek, amelyek legalább lehetővé teszik a szoftver ingyenes terjesztését bináris formában. Tartalmazza a `@FREE` címkét.
`@EULA`Ezek azok a licencek, amelyek megpróbálják elvenni az Ön jogait. Ezek szigorúbbak, mint a "Minden jog fenntartva" ("All rights reserved"), vagy kifejezett jóváhagyást igényelnek.

1. [↑](#cite_ref-3)[https://www.gnu.org/licenses/license-list.html](https://www.gnu.org/licenses/license-list.html)
2. [↑](#cite_ref-4)[https://www.opensource.org/licenses](https://www.opensource.org/licenses)
3. [↑](#cite_ref-5)[https://www.gnu.org/philosophy/free-sw.html](https://www.gnu.org/philosophy/free-sw.html)
4. [↑](#cite_ref-6)[https://freedomdefined.org/](https://freedomdefined.org/)

A jelenleg rendszer szinten elfogadható licencértékek a következőképpen tekinthetők meg:

`user $` `portageq envvar ACCEPT_LICENSE`

```
@FREE
```

Ahogy a kimeneten látható, az alapértelmezett érték az, hogy csak a `@FREE` kategóriába sorolt ​​szoftverek telepítése van engedélyezve.

A operációs rendszerhez tartozó konkrét licencek vagy licenccsoportok a következő helyeken határozhatóak meg:

- Rendszerszinten a kiválasztott profilban - ez állítja be az alapértelmezett értéket.
- Rendszerszinten a /etc/portage/make.conf fájlban. A rendszergazdák ebben a fájlban írhatják felül a profil alapértelmezett értékét.
- Szoftvercsomagonként egy /etc/portage/package.license fájlban.
- Szoftvercsomagonként egy /etc/portage/package.license/ _mappában_ található fájlokban.

A rendszer szintű licenc alapértelmezés a profilban felülírható a /etc/portage/make.conf fájlban:

FILE **`/etc/portage/make.conf`** **Licencek elfogadása az ACCEPT\_LICENSE változóval rendszer szinten**

```
# Felülírja a profil ACCEPT_LICENSE alapértelmezett értékét.
ACCEPT_LICENSE="-* @FREE @BINARY-REDISTRIBUTABLE"

```

Opcionálisan a rendszergazdák szoftvercsomagonként is meghatározhatják elfogadott licenceket, amint az a következő könyvtárban látható. Vegye figyelembe, hogy a package.license _könyvtárat_ létre kell hozni, ha még nem létezik:

`root #` `mkdir /etc/portage/package.license`

Az egyes Gentoo szoftvercsomagok szoftverlicenc-adatai a kapcsolódó ebuild LICENSE változójában tárolódnak. Egy szoftvercsomag egy vagy több szoftverlicencet tartalmazhat, ezért egy szoftvercsomaghoz több elfogadható licencet is meg kell adni.

FILE **`/etc/portage/package.license/kernel`** **Licencek elfogadása szoftvercsomagonként**

```
app-arch/unrar unRAR
sys-kernel/linux-firmware linux-fw-redistributable
sys-firmware/intel-microcode intel-ucode

```

**Important**

A LICENSE változó egy ebuildben csak iránymutatás a Gentoo fejlesztők és felhasználók számára. Ez _nem_ jogi nyilatkozat, és _nincs garancia_ arra, hogy a tükrözi a valóságot. Javasoljuk, hogy ne hagyatkozzon kizárólag az ebuild fejlesztőinek a szoftvercsomag licencének értelmezésére, de alaposan ellenőrizze magát a szoftvercsomagot, beleértve a rendszerre telepített összes fájlt is.

### Opcionális: @world készlet frissítése

**Tip**

Ha egy asztali környezet profil célpontja egy nem asztali szakasz fájlból lett kiválasztva, a @world frissítési folyamat jelentősen megnövelheti a telepítési időt. Azok, akik időhiányban szenvednek, a következő 'szabályt' alkalmazhatják: minél rövidebb a profil neve, annál kevésbé specifikus a rendszer @world készlete. Minél kevésbé specifikus a [@world készlet](/wiki/World_set_(Portage) "World set (Portage)"), annál kevesebb szoftvercsomagot igényel a rendszer. Pl.:

- A `default/linux/amd64/23.0` kiválasztása valószínűleg kevesebb frissítendő szoftvercsomagot fog igényelni, míg
- A `default/linux/amd64/23.0/desktop/gnome/systemd` kiválasztása valószínűleg több szoftvercsomag telepítését fogja igényelni, mivel a profil célpont nagyobb [@system](/wiki/Package_sets#.40system "Package sets") és [@profile](/wiki/Package_sets#.40profile "Package sets") készletekkel rendelkezik: a GNOME asztali környezethez szükséges függőségekkel.

A rendszer [@world készletének](/wiki/World_set_(Portage) "World set (Portage)") frissítése _nem kötelező_, és valószínűleg nem hajt végre funkcionális változtatásokat, hacsak nem hajtják végre a következő választható lépések közül egyet vagy többet:

1. Egy, a fokozat (stage) fájltól _eltérő_ profil célpont lett kiválasztva.
2. További USE jelölőzászlók lettek beállítva a telepített szoftvercsomagokhoz.

Azok az olvasók, akik 'Gentoo speed run'-t hajtanak végre (ahol a cél a telepítési folyamat minél rövidebb idő alatt történő befejezése), biztonságosan kihagyhatják a @world készlet frissítéseit, amíg a rendszerük _nem_ indul újra az új Gentoo környezetbe.

Azok az olvasók, akik józanabbul, lassabban hajtják végre a telepítést, megkérhetik a Portage szoftvercsomag-kezelőt a szoftvercsomag-, profil- és/vagy a USE jelölőzászló módosítások frissítésére:

`root #` `emerge --ask --verbose --update --deep --changed-use @world`

Azok az olvasók, akik hozzáadtak egy bináris kiszolgálót a [fenti](/wiki/Handbook:X86/Installation/Base/hu#Optional:_Adding_a_binary_package_host "Handbook:X86/Installation/Base/hu") szakaszban, megadhatják a --getbinpkg (vagy -g) opciót, hogy szoftvercsomagokat töltsenek le a bináris kiszolgálóról ahelyett, hogy azokat lefordítanák:

`root #` `emerge --ask --verbose --update --deep --newuse --getbinpkg @world`

#### Elévült szoftvercsomagok eltávolítása

Fontos, hogy a rendszerfrissítések után mindig végezze el a _depclean_ műveletet a felesleges szoftvercsomagok eltávolításának az érdekében. Gondosan ellenőrizze a kimenetet a emerge --depclean --pretend parancs segítségével, hogy megnézze, van-e olyan tisztításra kijelölt szoftvercsomag, amelyet személyesen használ és meg szeretne tartani. Egy szoftvercsomag megtartásához, amelyet egyébként a depclean eltávolítana, használja a emerge --noreplace szoftvercsomagneve parancsot.

`root #` `emerge --ask --pretend --depclean`

Ha elégedett, akkor folytassa ezzel a depclean paranccsal amely a törlést valóban végrehajtja:

`root #` `emerge --ask --depclean`

## Időzóna

**Note**

Ez a lépés nem vonatkozik a musl libc felhasználóira. Azok a felhasználók, akik nem tudják, mit jelent ez, hajtsák végre ezt a lépést.

**Warning**

Kérjük, kerülje el a /usr/share/zoneinfo/Etc/GMT\* időzónák használatát, mivel a nevük nem jelzi az időzónák valós időeltolódását. Például a GMT-8 név az a valóságban GMT+8 időeltérést jelent.

Válassza ki az operációs rendszer időzónáját. Keresse meg az elérhető időzónákat a /usr/share/zoneinfo/ könyvtárban:

`root #` `ls -l /usr/share/zoneinfo`

```
total 352
drwxr-xr-x 2 root root   1120 Jan  7 17:41 Africa
drwxr-xr-x 6 root root   2960 Jan  7 17:41 America
drwxr-xr-x 2 root root    280 Jan  7 17:41 Antarctica
drwxr-xr-x 2 root root     60 Jan  7 17:41 Arctic
drwxr-xr-x 2 root root   2020 Jan  7 17:41 Asia
drwxr-xr-x 2 root root    280 Jan  7 17:41 Atlantic
drwxr-xr-x 2 root root    500 Jan  7 17:41 Australia
drwxr-xr-x 2 root root    120 Jan  7 17:41 Brazil
-rw-r--r-- 1 root root   2094 Dec  3 17:19 CET
-rw-r--r-- 1 root root   2310 Dec  3 17:19 CST6CDT
drwxr-xr-x 2 root root    200 Jan  7 17:41 Canada
drwxr-xr-x 2 root root     80 Jan  7 17:41 Chile
-rw-r--r-- 1 root root   2416 Dec  3 17:19 Cuba
-rw-r--r-- 1 root root   1908 Dec  3 17:19 EET
-rw-r--r-- 1 root root    114 Dec  3 17:19 EST
-rw-r--r-- 1 root root   2310 Dec  3 17:19 EST5EDT
-rw-r--r-- 1 root root   2399 Dec  3 17:19 Egypt
-rw-r--r-- 1 root root   3492 Dec  3 17:19 Eire
drwxr-xr-x 2 root root    740 Jan  7 17:41 Etc
drwxr-xr-x 2 root root   1320 Jan  7 17:41 Europe
...

```

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

Tegyük fel, hogy a választott időzóna a Europe/Brussels. Ennek az időzónának a kiválasztásához egy szimbolikus linket lehet létrehozni ebből a zoneinfo fájlból a /etc/localtime útvonalra:

`root #` `ln -sf ../usr/share/zoneinfo/Europe/Brussels /etc/localtime`

**Tip**

A célútvonal, amely `../`-vel kezdődik, _a link helyéhez képest relatív_, nem pedig az aktuális könyvtárhoz.

**Note**

Egy abszolút útvonal is használható a szimbolikus linkhez, de a systemd timedatectl által létrehozott relatív link jobban kompatibilis az alternatív _ROOT_-okkal.

## Nyelvterület-beállítások

**Note**

Ez a lépés nem vonatkozik a musl libc felhasználóira. azoknak a felhasználóknak, akik nem tudják, hogy ez mit jelent, végre kell hajtaniuk ezt a lépést.

### Nyelvterület-beállítások létrehozása

A default installation of Gentoo Linux provides an archive that contains all supported locales, numbering 500 or more. However, it is typical for an administrator to require only one or two of these. In that case, the /etc/locale.gen configuration file may be populated with a list of the required locales. By default, locale-gen shall read this file and compile only the locales that are specified, saving both time and space in the longer term.

A nyelvterület-beállítások nemcsak a felhasználó által az operációs rendszerrel való interakcióhoz használt nyelvet határozzák meg, hanem a karakterláncok rendezési szabályait, a dátumok és időpontok megjelenítését stb. is. A nyelvterület-beállítások _kis- és nagybetű érzékenyek_, és pontosan a leírtak szerint kell őket ábrázolni. Az elérhető nyelvterület-beállítások teljes listája a /usr/share/i18n/SUPPORTED fájlban található.

`root #` `nano /etc/locale.gen`

A következő nyelvterület-beállítások példák arra, hogy miként lehet beállítani az angol (Egyesült Államok) és a német (Németország/Deutschland) nyelvet az azokat kísérő karakterformátumokkal (például UTF-8):

FILE **`/etc/locale.gen`** **Az USA és a DE helyi beállítások engedélyezése a megfelelő karakterformátumokkal**

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

**Warning**

Sok szoftver megfelelő létrehozásához legalább egy UTF-8 nyelvterület-beállítás szükséges.

A következő lépés a locale-gen parancs futtatása. Ez a parancs legenerálja az összes, a /etc/locale.gen fájlban meghatározott nyelvterület-beállítást.

`root #` `locale-gen`

Annak ellenőrzéséhez, hogy a kiválasztott nyelvterület-beállítások most már elérhetőek-e, futtassa a következő parancsot: locale -a.

A systemd telepítésekor a localectl parancs használható, például localectl set-locale ... vagy localectl list-locales.

### Nyelvterület-beállítások kiválasztása

Ha ez kész, akkor itt az ideje a rendszer szintű nyelvterület-beállítások megadásának. Ismét a eselect parancsot használjuk, ezúttal a `locale` modullal.

A eselect locale list parancs használatával megjeleníthetők az elérhető célok:

`root #` `eselect locale list`

```
Available targets for the LANG variable:
  [1]  C
  [2]  C.utf8
  [3]  en_US
  [4]  en_US.iso88591
  [5]  en_US.utf8
  [6]  de_DE
  [7]  de_DE.iso88591
  [8]  de_DE.utf8
  [9] POSIX
  [ ]  (free form)

```

Az eselect locale set <SORSZÁM> parancs használatával kiválasztható a megfelelő nyelvterület-beállítás:

`root #` `eselect locale set 2`

Manuálisan ezt még mindig el lehet végezni az /etc/env.d/02locale fájlon keresztül, és systemd esetében az /etc/locale.conf fájlon keresztül:

FILE **`/etc/env.d/02locale`** **Nyelvterület-beállítások manuális megadása**

```
LANG="de_DE.UTF-8"
LC_COLLATE="C.UTF-8"

```

A nyelvterület-beállítások megadása elkerüli a figyelmeztetéseket és hibákat a kernel és a szoftverek kódfordítása során a telepítés későbbi szakaszaiban.

Most töltse be újra a környezetet:

`root #` `env-update && source /etc/profile && export PS1="(chroot) ${PS1}"`

A nyelvterület kiválasztás folyamatával kapcsolatban olvassa el a [Nyelvterület-beállítása útmutatót](/wiki/Localization/Guide/hu "Localization/Guide/hu") és az [UTF-8](/wiki/UTF-8/hu "UTF-8/hu") útmutatót.

## Hivatkozások

1. [↑](#cite_ref-1)[https://www.gentoo.org/news/2023/12/29/Gentoo-binary.html](https://www.gentoo.org/news/2023/12/29/Gentoo-binary.html)
2. [↑](#cite_ref-2)[https://www.gentoo.org/glep/glep-0023.html#id7](https://www.gentoo.org/glep/glep-0023.html#id7)

[← Stage fájl telepítése](/wiki/Handbook:X86/Installation/Stage/hu "Previous part") [Kezdőlap](/wiki/Handbook:X86/hu "Handbook:X86/hu") [Kernel beállítása →](/wiki/Handbook:X86/Installation/Kernel/hu "Next part")