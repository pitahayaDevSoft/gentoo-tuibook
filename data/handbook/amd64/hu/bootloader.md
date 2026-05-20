# Bootloader

Other languages:

- [Deutsch](/wiki/Handbook:AMD64/Installation/Bootloader/de "Handbook:AMD64/Installation/Bootloader/de (100% translated)")
- [English](/wiki/Handbook:AMD64/Installation/Bootloader "Handbook:AMD64/Installation/Bootloader (100% translated)")
- [español](/wiki/Handbook:AMD64/Installation/Bootloader/es "Handbook:AMD64/Installation/Bootloader/es (100% translated)")
- [français](/wiki/Handbook:AMD64/Installation/Bootloader/fr "Handbook:AMD64/Installation/Bootloader/fr (100% translated)")
- [italiano](/wiki/Handbook:AMD64/Installation/Bootloader/it "Handbook:AMD64/Installation/Bootloader/it (100% translated)")
- magyar
- [polski](/wiki/Handbook:AMD64/Installation/Bootloader/pl "Handbook:AMD64/Installation/Bootloader/pl (100% translated)")
- [português do Brasil](/wiki/Handbook:AMD64/Installation/Bootloader/pt-br "Handbook:AMD64/Installation/Bootloader/pt-br (100% translated)")
- [čeština](/wiki/Handbook:AMD64/Installation/Bootloader/cs "Handbook:AMD64/Installation/Bootloader/cs (100% translated)")
- [русский](/wiki/Handbook:AMD64/Installation/Bootloader/ru "Handbook:AMD64/Installation/Bootloader/ru (100% translated)")
- [தமிழ்](/wiki/Handbook:AMD64/Installation/Bootloader/ta "Handbook:AMD64/Installation/Bootloader/ta (100% translated)")
- [中文（中国大陆）‎](/wiki/Handbook:AMD64/Installation/Bootloader/zh-cn "Handbook:AMD64/Installation/Bootloader/zh-cn (100% translated)")
- [日本語](/wiki/Handbook:AMD64/Installation/Bootloader/ja "Handbook:AMD64/Installation/Bootloader/ja (100% translated)")
- [한국어](/wiki/Handbook:AMD64/Installation/Bootloader/ko "Handbook:AMD64/Installation/Bootloader/ko (100% translated)")

[AMD64 kézikönyv](/wiki/Handbook:AMD64/hu "Handbook:AMD64/hu")[A Gentoo Linux telepítése](/wiki/Handbook:AMD64/Full/Installation/hu "Handbook:AMD64/Full/Installation/hu")[A telepítésről](/wiki/Handbook:AMD64/Installation/About/hu "Handbook:AMD64/Installation/About/hu")[Telepítőképfájl kiválasztása](/wiki/Handbook:AMD64/Installation/Media/hu "Handbook:AMD64/Installation/Media/hu")[Hálózat beállítása](/wiki/Handbook:AMD64/Installation/Networking/hu "Handbook:AMD64/Installation/Networking/hu")[Adathordozók előkészítése](/wiki/Handbook:AMD64/Installation/Disks/hu "Handbook:AMD64/Installation/Disks/hu")[Stage fájl](/wiki/Handbook:AMD64/Installation/Stage/hu "Handbook:AMD64/Installation/Stage/hu")[Alaprendszer telepítése](/wiki/Handbook:AMD64/Installation/Base/hu "Handbook:AMD64/Installation/Base/hu")[Kernel beállítása](/wiki/Handbook:AMD64/Installation/Kernel/hu "Handbook:AMD64/Installation/Kernel/hu")[Rendszer beállítása](/wiki/Handbook:AMD64/Installation/System/hu "Handbook:AMD64/Installation/System/hu")[Eszközök telepítése](/wiki/Handbook:AMD64/Installation/Tools/hu "Handbook:AMD64/Installation/Tools/hu")Bootloader beállítása[Telepítés véglegesítése](/wiki/Handbook:AMD64/Installation/Finalizing/hu "Handbook:AMD64/Installation/Finalizing/hu")[Munka a Gentoo rendszerrel](/wiki/Handbook:AMD64/Full/Working/hu "Handbook:AMD64/Full/Working/hu")[Portage bemutatása](/wiki/Handbook:AMD64/Working/Portage/hu "Handbook:AMD64/Working/Portage/hu")[USE jelölőzászlók](/wiki/Handbook:AMD64/Working/USE/hu "Handbook:AMD64/Working/USE/hu")[Portage jellemzői](/wiki/Handbook:AMD64/Working/Features/hu "Handbook:AMD64/Working/Features/hu")[Init-szkript rendszer](/wiki/Handbook:AMD64/Working/Initscripts/hu "Handbook:AMD64/Working/Initscripts/hu")[Környezeti változók](/wiki/Handbook:AMD64/Working/EnvVar/hu "Handbook:AMD64/Working/EnvVar/hu")[Munka a Portage szoftvercsomag-kezelővel](/wiki/Handbook:AMD64/Full/Portage/hu "Handbook:AMD64/Full/Portage/hu")[Fájlok és könyvtárak](/wiki/Handbook:AMD64/Portage/Files/hu "Handbook:AMD64/Portage/Files/hu")[Változók](/wiki/Handbook:AMD64/Portage/Variables/hu "Handbook:AMD64/Portage/Variables/hu")[Szoftverágak keverése](/wiki/Handbook:AMD64/Portage/Branches/hu "Handbook:AMD64/Portage/Branches/hu")[További eszközök](/wiki/Handbook:AMD64/Portage/Tools/hu "Handbook:AMD64/Portage/Tools/hu")[Egyéni szoftvercsomag-tárolók](/wiki/Handbook:AMD64/Portage/CustomTree/hu "Handbook:AMD64/Portage/CustomTree/hu")[Fejlett funkciók](/wiki/Handbook:AMD64/Portage/Advanced/hu "Handbook:AMD64/Portage/Advanced/hu")[Hálózat beállítása OpenRC init-rendszeren](/wiki/Handbook:AMD64/Full/Networking/hu "Handbook:AMD64/Full/Networking/hu")[Munka elkezdése](/wiki/Handbook:AMD64/Networking/Introduction/hu "Handbook:AMD64/Networking/Introduction/hu")[Fejlett beállítások](/wiki/Handbook:AMD64/Networking/Advanced/hu "Handbook:AMD64/Networking/Advanced/hu")[Moduláris hálózat](/wiki/Handbook:AMD64/Networking/Modular/hu "Handbook:AMD64/Networking/Modular/hu")[Vezeték nélküli (Wi-Fi)](/wiki/Handbook:AMD64/Networking/Wireless/hu "Handbook:AMD64/Networking/Wireless/hu")[Funkcionalitás hozzáadása](/wiki/Handbook:AMD64/Networking/Extending/hu "Handbook:AMD64/Networking/Extending/hu")[Dinamikus menedzsment](/wiki/Handbook:AMD64/Networking/Dynamic/hu "Handbook:AMD64/Networking/Dynamic/hu")

## Contents

- [1Rendszertöltő (bootloader) kiválasztása](#Rendszert.C3.B6lt.C5.91_.28bootloader.29_kiv.C3.A1laszt.C3.A1sa)
- [2Default: GRUB](#Default:_GRUB)
  - [2.1Emerge](#Emerge)
  - [2.2Telepítés](#Telep.C3.ADt.C3.A9s)
    - [2.2.1DOS/Örökölt BIOS alapú operációs rendszerek](#DOS.2F.C3.96r.C3.B6k.C3.B6lt_BIOS_alap.C3.BA_oper.C3.A1ci.C3.B3s_rendszerek)
    - [2.2.2UEFI alapú operációs rendszerek](#UEFI_alap.C3.BA_oper.C3.A1ci.C3.B3s_rendszerek)
      - [2.2.2.1Optional: Secure Boot](#Optional:_Secure_Boot)
      - [2.2.2.2GRUB hibakeresés](#GRUB_hibakeres.C3.A9s)
  - [2.3GRUB beállítása](#GRUB_be.C3.A1ll.C3.ADt.C3.A1sa)
- [3Alternatíva 4: systemd-boot](#Alternat.C3.ADva_4:_systemd-boot)
  - [3.1Emerge](#Emerge_2)
  - [3.2Telepítés](#Telep.C3.ADt.C3.A9s_2)
  - [3.3Optional: Secure Boot](#Optional:_Secure_Boot_2)
- [4Alternatíva 2: EFI Stub](#Alternat.C3.ADva_2:_EFI_Stub)
  - [4.1Egységesített kernelképfájl (Unified Kernel Image)](#Egys.C3.A9ges.C3.ADtett_kernelk.C3.A9pf.C3.A1jl_.28Unified_Kernel_Image.29)
- [5Egyéb alternatívák](#Egy.C3.A9b_alternat.C3.ADv.C3.A1k)
- [6Rendszer újraindítása](#Rendszer_.C3.BAjraind.C3.ADt.C3.A1sa)

## Rendszertöltő (bootloader) kiválasztása

A Linux kernel beállítása, a rendszereszközök telepítése és a beállításfájlok szerkesztése után itt az ideje telepíteni a számítógépre feltelepített Linux alapú operációs rendszer utolsó fontos részét: a rendszertöltőt (bootloader).

A boot loader felelős a Linux kernel elindításáért az számítógép indításakor - nélküle az operációs rendszer nem tudná, hogy miként lépjen tovább, miután a bekapcsológombot megnyomtuk a gépünkön.

A(z) **amd64** esetében dokumentáljuk, hogy miként lehet beállítani a [GRUB](/wiki/Handbook:AMD64/Blocks/Bootloader/hu#Default:_GRUB "Handbook:AMD64/Blocks/Bootloader/hu") rendszerbetöltőt (bootloadert) DOS/Örökölt BIOS alapú operációs rendszerekhez, valamint az [GRUB](#Default:_GRUB), [systemd-boot](#Alternative_1:_systemd-boot) vagy [EFI Stub](#Alternative_2:_efibootmgr) beállítását az UEFI alapú operációs rendszerekhez.

A kézikönyv ezen szekciójában különbség van téve a boot loader szoftvercsomagjának _emerge_ parancsal való letöltése és létrehozása, valamint a boot loader _telepítése_ a rendszerünk adathordozójára között. Itt az _emerge_ kifejezést arra használjuk, hogy a [Portage](/wiki/Portage "Portage") kérje a szoftvercsomag elérhetővé tételét az operációs rendszerünk számára. Az _telepítés_ kifejezés azt jelenti, hogy boot loader fájlokat másol vagy fizikailag módosítja az operációs rendszert tartalmazó adathordozójának megfelelő szakaszait annak érdekében, hogy a boot loader a következő számítógép bekapcsolási ciklus során _aktivált és működésre kész_ legyen.

## Default: GRUB

Alapértelmezés szerint a Gentoo operációs rendszerek többsége ma már a [GRUB](/wiki/GRUB/hu "GRUB/hu") rendszerbetöltőre (bootloaderre) (ami a [sys-boot/grub](https://packages.gentoo.org/packages/sys-boot/grub) szoftvercsomagban található) támaszkodik, amely közvetlen utódja a [GRUB Legacy](/wiki/GRUB_Legacy "GRUB Legacy") rendszerbetöltőnek (bootloadernek). További beállítás nélkül a GRUB nagyon jól támogatja a régebbi BIOS ("pc") operációs rendszereket. Egy kis beállítással, amelyre az szoftverépítési idő előtt van szükség, a GRUB több mint fél tucat további platformot is támogathat. További információért tekintse meg a [Előfeltételek szakaszt](/wiki/GRUB/hu#El.C5.91felt.C3.A9telek "GRUB/hu") a [GRUB](/wiki/GRUB/hu "GRUB/hu") cikkben.

### Emerge

Régebbi BIOS alapú operációs rendszer használata esetén, amely csak MBR partíciós táblázatokat támogat, nincs szükség további beállításra a GRUB telepítéséhez:

`root #` `emerge --ask --verbose sys-boot/grub`

Megjegyzés UEFI felhasználóknak: A fenti parancs futtatása a GRUB\_PLATFORMS engedélyezett értékeit fogja megjeleníteni az emerge előtt. UEFI képes rendszerek használatakor a felhasználóknak biztosítaniuk kell, hogy a `GRUB_PLATFORMS="efi-64"` engedélyezve van (ahogy alapértelmezés szerint is). Ha ez nem így van a beállításnál, akkor a `GRUB_PLATFORMS="efi-64"` értéket hozzá kell adni a /etc/portage/make.conf fájlhoz az emerge _előtt_ annak érdekében, hogy a szoftvercsomag EFI funkcionalitással legyen létrehozva.

`root #` `echo 'GRUB_PLATFORMS="efi-64"' >> /etc/portage/make.conf`

`root #` `emerge --ask sys-boot/grub`

Ha a GRUB valamiért úgy lett az emerge-vel létrehozva, hogy nem volt engedélyezve a `GRUB_PLATFORMS="efi-64"`, akkor a sor (ahogy fent látható) hozzáadható a make.conf fájlhoz, majd az [world szoftvercsomagkészlet](/wiki/World_set_(Portage) "World set (Portage)") függőségeit újra lehet számoltatni azáltal, hogy a `--update --newuse` opciókat átadjuk az emerge parancsnak:

`root #` `emerge --ask --update --newuse --verbose sys-boot/grub`

A GRUB szoftver most már az emerge paranccsal be lett olvasztva az _operációs rendszerünkbe_, de még nem lett telepítve másodlagos _bootloaderként_.

### Telepítés

Ezután telepítse a szükséges GRUB fájlokat a /boot/grub/ könyvtárba a grub-install parancs segítségével. Feltételezve, hogy az első adathordozó (ahonnan a rendszer indul) a /dev/sda, az alábbi parancsok egyike megfelelő lesz:

#### DOS/Örökölt BIOS alapú operációs rendszerek

DOS/Örökölt BIOS operációs rendszerek számára:

`root #` `grub-install /dev/sda`

#### UEFI alapú operációs rendszerek

**Important**

Mindenképp győződjön meg arról, hogy az EFI rendszerpartíció valóban csatolva lett a fájlrendszerbe _mielőtt_ futtatná a grub-install parancsot. Ugyanis előfordulhat, hogy a grub-install parancs a GRUB EFI fájlt (grubx64.efi) a rossz könyvtárba telepíti (mert nincs felcsatolva az EFI rendszerpartíció). Ráadásul _minden_ figyelmeztetés **nélkül** teszi ezt meg.

UEFI rendszerekhez:

`root #` `grub-install --efi-directory=/efi`

```
Installing for x86_64-efi platform.
Installation finished. No error reported.

```

Sikeres telepítés után a kimenetnek meg kell egyeznie az előző parancs kimenetével. Ha a kimenet nem egyezik pontosan, akkor folytassa a [GRUB hibakeresés](/wiki/Handbook:AMD64/Blocks/Bootloader/hu#Debugging_GRUB "Handbook:AMD64/Blocks/Bootloader/hu") szakasztól, ellenkező esetben ugorjon a [GRUB beállítása](/wiki/Handbook:AMD64/Blocks/Bootloader/hu#GRUB_Configure "Handbook:AMD64/Blocks/Bootloader/hu") szakaszhoz.

##### Optional: Secure Boot

A sikeres bootoláshoz a bekapcsolt secure boot mellett az aláíró tanúsítványt el kell fogadtatni a [UEFI](/wiki/UEFI "UEFI") firmware-rel, vagy a [shim](/wiki/Shim "Shim")-et kell használni előtöltőként (pre-loader). A Shim előre alá van írva a harmadik fél Microsoft tanúsítványával, amelyet alapértelmezés szerint a legtöbb UEFI alaplap elfogad.

Az, hogy miként lehet beállítani az UEFI firmware-t egyéni kulcsok elfogadására, a firmware gyártójától függ, ami meghaladja ezen kézikönyv kereteit. Az alábbiakban a shim beállításának módját mutatjuk be. Itt feltételezzük, hogy a felhasználó már követte az előző szakaszok utasításait az aláírókulcs létrehozására és a portage beállítására vonatkozóan. Ha ez nem így van, akkor kérjük, hogy térjen vissza először a [Kernel telepítése](/wiki/Handbook:AMD64/Installation/Kernel/hu "Handbook:AMD64/Installation/Kernel/hu") részhez.

A [sys-boot/grub](https://packages.gentoo.org/packages/sys-boot/grub) szoftvercsomag előre összeállított és aláírt önálló EFI végrehajtható fájlt telepít, amennyiben a [secureboot](https://packages.gentoo.org/useflags/secureboot) [USE flag/hu](/wiki/USE_flag/hu "USE flag/hu") USE jelölőzászló engedélyezve van. Telepítse a szükséges szoftvercsomagokat, majd másolja az önálló grub-ot, a Shim-et és a MokManager-t ugyanabba a könyvtárba az EFI rendszerpartíción. Például:

`root #` `emerge sys-boot/grub sys-boot/shim sys-boot/mokutil sys-boot/efibootmgr`

`root #` `cp /usr/share/shim/BOOTX64.EFI /efi/EFI/Gentoo/shimx64.efi
`

`root #` `cp /usr/share/shim/mmx64.efi /efi/EFI/Gentoo/mmx64.efi
`

`root #` `cp /usr/lib/grub/grub-x86_64.efi.signed /efi/EFI/Gentoo/grubx64.efi
`

Ezután regisztrálja az aláíró kulcsot a shim MOKlistájába. Ehhez DER formátumú kulcsokra van szükség, miközben a sbsign és a kernel készítő rendszer PEM formátumú kulcsokat vár. A kézikönyv korábbi részeiben bemutattunk egy példát egy ilyen PEM formátumú aláíró kulcs létrehozására, ezt a kulcsot most át kell alakítani DER formátumúra:

`root #` `openssl x509 -in /path/to/kernel_key.pem -inform PEM -out /path/to/kernel_key.der -outform DER`

**Note**

Az itt használt útvonalnak a legenerált kulcshoz tartozó tanúsítványt tartalmazó pem fájlhoz kell vezetnie. Ebben a példában a kulcs és a tanúsítvány ugyanabban a pem fájlban található.

Ezután az átalakított tanúsítvány importálható a Shims MOKlistába, ez a parancs kéri egy jelszó megadását az importálási kérelemhez:

`root #` `mokutil --import /path/to/kernel_key.der`

**Note**

Amikor a jelenleg betöltött kernel már megbízik az importálandó tanúsítványban, akkor a következő üzenet jelenik meg: "Already in kernel trusted keyring." Ebben az esetben futtassa újra a fenti parancsot a --ignore-keyring argumentummal kiegészítve.

Ezután regisztráljuk a Shim-et az UEFI firmware-be. A következő parancsban a `boot-disk` és `boot-partition-id` helyére az EFI rendszerpartíció adathordozó- és partícióazonosítóját kell behelyettesíteni:

`root #` `efibootmgr --create --disk /dev/boot-disk --part boot-partition-id --loader '\EFI\Gentoo\shimx64.efi' --label 'GRUB via Shim' --unicode`

Vegye figyelembe, hogy ez az előre összeállított és aláírt önálló grub-verzió a grub.cfg fájlt egy szokásostól eltérő helyről olvassa be. Az alapértelmezett /boot/grub/grub.cfg helyett a beállításfájlnak ugyanabban a könyvtárban kell lennie, ahol a grub EFI végrehajtható fájl található, például: /efi/EFI/Gentoo/grub.cfg. Amikor a [sys-kernel/installkernel](https://packages.gentoo.org/packages/sys-kernel/installkernel) eszközt használják a kernel telepítésére és a grub beállításának frissítésére a GRUB\_CFG környezeti változóval lehet felülírni a grub beállításfájl szokásos helyét.

Például:

`root #` `grub-mkconfig -o /efi/EFI/Gentoo/grub.cfg`

Or, via [installkernel](/wiki/Installkernel/hu "Installkernel/hu"):

FILE **`/etc/env.d/99grub`**

```
GRUB_CFG=/efi/EFI/Gentoo/grub.cfg

```

`root #` `env-update`

**Note**

Az importálási folyamat nem fejeződik be a rendszer újraindításáig. Miután a kézikönyv összes lépését végrehajtotta, indítsa újra a rendszert. Ekkor a Shim betöltődik, és megtalálja a mokutil által regisztrált importálási kérelmet. A MokManager alkalmazás elindul, és kéri a jelszót, amelyet az importálási kérelem létrehozásakor adott meg. Kövesse a képernyőn megjelenő utasításokat a tanúsítvány importálásának befejezéséhez, majd indítsa újra a rendszert az UEFI menüben, és engedélyezze a Secure Boot beállítást.

##### GRUB hibakeresés

A GRUB hibakeresésekor van néhány gyors javítási módszer, amely anélkül eredményezhet egy jól működőd bootolható feltelepített operációs rendszert, hogy ismét be kellene lépni a Live ISO élő képfájl telepítő környezetbe.

Abban az esetben, ha az "EFI változók nem támogatottak ezen az operációs rendszeren" üzenet jelenik meg a képernyőn, akkor valószínű, hogy a Live ISO élő képfájl nem EFI módban indult el, hanem jelenleg az Örökölt BIOS indítási módban van. A megoldás a lent említett [eltávolítható GRUB lépés](/wiki/Handbook:AMD64/Blocks/Bootloader/hu#GRUB_Install_EFI_systems_removable "Handbook:AMD64/Blocks/Bootloader/hu") kipróbálása. Ez felülírja a végrehajtható EFI fájlt ami a /EFI/BOOT/BOOTX64.EFI helyen található. EFI módban történő újraindítás után az alaplap firmware-je végrehajthatja ezt az alapértelmezett indítási bejegyzést és elindíthatja a GRUB bootloadert.

Ha a grub-install hibát ad vissza, amely szerint "Nem sikerült előkészíteni a Boot változót: Csak olvasható fájlrendszer" ("Could not prepare Boot variable: Read-only file system"), és a Live ISO élő környezet helyesen indult UEFI módban, akkor lehetséges az efivars speciális csatolást újraírhatóként felcsatolni, majd ismét futtatni az [említett grub-install parancsot](/wiki/Handbook:AMD64/Blocks/Bootloader/hu#GRUB_Install_EFI_systems_command "Handbook:AMD64/Blocks/Bootloader/hu"):

`root #` `mount -o remount,rw,nosuid,nodev,noexec --types efivarfs efivarfs /sys/firmware/efi/efivars`

Ez bizonyos nem hivatalos Gentoo környezetek által okozott probléma, amelyek alapértelmezés szerint nem csatolják a speciális EFI fájlrendszert. Ha az előző parancs nem fut le, akkor indítsa újra a számítógépet egy hivatalos Gentoo Live ISO élőképfájl környezetben, EFI módban.

Egyes alaplapgyártók gyenge UEFI megvalósításai _csak_ az /EFI/BOOT könyvtárhelyet támogatják az .EFI fájl számára az EFI rendszerpartíción (ESP). A GRUB telepítő automatikusan létrehozza az .EFI fájlt ezen a helyen, amennyiben a `--removable` opciót hozzáadja a telepítési parancshoz. Győződjön meg arról, hogy az ESP partíció fel van csatolva, mielőtt futtatná a következő parancsot. Most feltételezve, hogy a /efi (ahogy korábban meghatároztuk) valóban fel van csatolva, futtassa a következő parancsot:

`root #` `grub-install --target=x86_64-efi --efi-directory=/efi --removable`

Ez létrehozza az UEFI specifikáció által meghatározott 'default' könyvtárat, majd létrehoz egy fájlt az alapértelmezett névvel:
BOOTX64.EFI.

### GRUB beállítása

Ezután generálja le a GRUB beállítást a /etc/default/grub fájlban és a /etc/grub.d szkriptekben megadott felhasználói beállítások alapján. A legtöbb esetben nincs szükség a felhasználók általi beállításokra, mivel a GRUB automatikusan észleli, melyik kernelt kell indítani (a legmagasabbat a /boot/ elérhetőek közül) és melyik a root fájlrendszer. Az is lehetséges, hogy kernelparamétereket adjon hozzá a /etc/default/grub fájlban a GRUB\_CMDLINE\_LINUX változó segítségével.

A végső GRUB beállítás legenerálása érdekében futtassa a grub-mkconfig parancsot:

`root #` `grub-mkconfig -o /boot/grub/grub.cfg`

```
Generating grub.cfg ...
Found linux image: /boot/vmlinuz-6.18.8-gentoo
Found initrd image: /boot/initramfs-genkernel-amd64-6.18.8-gentoo
done

```

A parancs kimenetének említenie kell, hogy legalább egy Linux képfájl megtalálható, mivel ezek szükségesek az operációs rendszer indításához. Ha initramfs képfájlt használunk vagy használva lett a genkernel a kernel létrehozásához, akkor a helyes initrd képfájlt is észlelnie kell. Ha ez nem így van, akkor menjen a /boot/ könyvtárba, és ellenőrizze a tartalmát a ls paranccsal. Ha a fájlok valóban hiányoznak, akkor térjen vissza a kernelbeállítás és telepítési utasításokhoz.

**Tip**

A os-prober segédprogram a GRUB bootloaderrel együtt használható más operációs rendszerek észlelésére a csatlakoztatott adathordozókról. A Windows 7, 8.1, 10 és más Linux disztribúciók felismerhetők. Azoknak, akik kettős boot (dualboot) rendszereket szeretnének, nekik telepíteniük kell a [sys-boot/os-prober](https://packages.gentoo.org/packages/sys-boot/os-prober) szoftvercsomagot, majd újra kell futtatniuk a grub-mkconfig parancsot (ahogy fent látható). Ha észlelési problémák merülnének fel, akkor feltétlenül olvassák el a [GRUB](/wiki/GRUB/hu "GRUB/hu") cikket teljes egészében _mielőtt_ támogatást kérnének a Gentoo közösségtől.

## Alternatíva 4: systemd-boot

Egy másik lehetőség a [systemd-boot](/wiki/Systemd/systemd-boot "Systemd/systemd-boot"), amely mind az OpenRC, mind a systemd számítógépeken működik. Ez egy vékony láncbetöltő (chainloader), és jól működik a biztonságos indítással (secure boot).

### Emerge

A systemd-boot telepítésének érdekében engedélyezze a [boot](https://packages.gentoo.org/useflags/boot) [USE flag/hu](/wiki/USE_flag/hu "USE flag/hu") USE jelölőzászlót, majd telepítse újra a [sys-apps/systemd](https://packages.gentoo.org/packages/sys-apps/systemd) szoftvercsomagot (systemd rendszerek esetén) vagy a [sys-apps/systemd-utils](https://packages.gentoo.org/packages/sys-apps/systemd-utils) szoftvercsomagot (OpenRC rendszerek esetén):

FILE **`/etc/portage/package.use/systemd-boot`**

```
sys-apps/systemd boot
sys-apps/systemd-utils boot

```

`root #` `emerge --ask sys-apps/systemd`

Vagy

`root #` `emerge --ask sys-apps/systemd-utils`

### Telepítés

Most rendelje hozzá a systemd-boot betöltőt az [EFI System Partition](/wiki/EFI_System_Partition/hu "EFI System Partition/hu") partícióhoz:

`root #` `bootctl install`

**Important**

Győződjön meg arról, hogy az EFI rendszerpartíció fel van csatolva a fájlrendszerbe _mielőtt_ futtatja a bootctl install parancsot.

Ezt a boootloadert használva, mielőtt újraindítaná a rendszert, ellenőrizze, hogy létrejött-e egy új boot bejegyzés az alábbiak használatával:

`root #` `bootctl list`

**Warning**

A rendszer-boot bejegyzéseinek kernel parancssorát az /etc/kernel/cmdline vagy az /usr/lib/kernel/cmdline fájlból olvassa be. Ha egyik fájl sem létezik, akkor a jelenleg betöltött kernel parancssora kerül újrahasználásra (/proc/cmdline). Új operációs rendszer telepítéseknél előfordulhat, hogy a Live CD kernel parancssorát véletlenül a új kernel indításához használják. A regisztrált bejegyzések kernel parancssora a következő paranccsal ellenőrizhető:

`root #` `bootctl list`

. Ha ez nem mutatja a kívánt kernel parancssort, akkor hozzon létre egy /etc/kernel/cmdline fájlt a megfelelő kernel parancssorral, és telepítse újra a kernelt.

Ha nem létezik új boot bejegyzés, győződjön meg arról, hogy a [sys-kernel/installkernel](https://packages.gentoo.org/packages/sys-kernel/installkernel) szoftvercsomag telepítve lett-e a [systemd](https://packages.gentoo.org/useflags/systemd) [USE flag/hu](/wiki/USE_flag/hu "USE flag/hu") és [systemd-boot](https://packages.gentoo.org/useflags/systemd-boot) [USE flag/hu](/wiki/USE_flag/hu "USE flag/hu") USE jelölőzászlók engedélyezésével, majd futtassa újra a kernel telepítését.

Az terjesztési (disztribúció) kernelek esetében futtassa a következő parancsot:

`root #` `emerge --ask --config sys-kernel/gentoo-kernel`

A kézzel beállított és lefordított kernel esetében futtassa a következő parancsot:

`root #` `make install`

**Important**

Amikor a kernelt telepíti a systemd-boot -hoz, akkor alapértelmezés szerint nem ad hozzá root= kernel parancssori argumentumot. A systemd init-rendszert használó operációs rendszereken, amelyek initramfs fájlt használnak, a felhasználók inkább a [systemd-gpt-auto-generator](/wiki/Systemd/hu#Automatic_mounting_of_partitions_at_boot.2Fhu "Systemd/hu")-ra támaszkodhatnak, hogy automatikusan megtalálják a gyökérpartíciót bootoláskor. Ellenkező esetben a felhasználóknak kézzel kell megadniuk a gyökér partíció helyét a /etc/kernel/cmdline fájlban a root= beállításával, valamint bármely más kernel parancssori argumentumokkal, amelyeket használni szeretnének. Ezután telepítse újra a kernelt a fent leírtak szerint.

### Optional: Secure Boot

Amikor a [secureboot](https://packages.gentoo.org/useflags/secureboot) [USE flag/hu](/wiki/USE_flag/hu "USE flag/hu") USE jelölőzászló engedélyezve van, akkor a systemd-boot EFI végrehajtható fájlt a portage automatikusan aláírja. Továbbá, a bootctl install automatikusan telepíti az aláírt verziót.

A secure boot bekapcsolt állapotában történő sikeres bootoláshoz a használt tanúsítványnak vagy el kell fogadnia a [UEFI](/wiki/UEFI "UEFI") firmware-t, vagy a [shim](/wiki/Shim "Shim")-et kell használni előbetöltőként (pre-loader). A shim előzetesen alá van írva a harmadik féltől Microsoft Tanúsítvánnyal, amelyet a legtöbb UEFI alaplap alapértelmezés szerint elfogad.

Az, hogy miként lehet beállítani az UEFI firmware-t egyéni kulcsok elfogadására, a firmware gyártójától függ, ami meghaladja a kézikönyv kereteit. Az alábbiakban a shim beállításának módját mutatjuk be. Itt feltételezzük, hogy a felhasználó már követte az előző szakaszok utasításait az aláírókulcs létrehozására és a portage használatára való beállítására. Ha ez nem történt meg, akkor kérjük, hogy először térjen vissza a [Kernel installation](/wiki/Handbook:AMD64/Installation/Kernel/hu "Handbook:AMD64/Installation/Kernel/hu") részhez.

`root #` `emerge --ask sys-boot/shim sys-boot/mokutil sys-boot/efibootmgr`

`root #` `bootctl install --no-variables`

`root #` `cp /usr/share/shim/BOOTX64.EFI /efi/EFI/systemd/shimx64.efi`

`root #` `cp /usr/share/shim/mmx64.efi /efi/EFI/systemd/mmx64.efi`

A shim a MOKlist kulcsokat valójában DER formátumban igényli. Mivel az OpenSSL által használt példa egy PEM formátumú kulcsot hoz létre, ezért a kulcsot a következő parancs segítségével lehet DER formátumba átkonvertálni:

`root #` `openssl x509 -in /path/to/kernel_key.pem -inform PEM -out /path/to/kernel_key.der -outform DER`

**Note**

Az itt használt elérési útnak a generált kulcshoz tartozó tanúsítványt tartalmazó pem fájl elérési útjának kell lennie. Ebben a példában mind a kulcs, mind a tanúsítvány ugyanabban a pem fájlban található.

Ezután az átalakított tanúsítvány importálható a shims MOKlist-be:

`root #` `mokutil --import /path/to/kernel_key.der`

**Note**

Amikor a jelenleg betöltött kernel már megbízik az importálandó tanúsítványban, akkor a következő üzenet jelenik meg: "Already in kernel trusted keyring." Ebben az esetben futtassa újra a fenti parancsot a --ignore-keyring argumentummal kiegészítve.

Végül regisztráljuk a shim-et a UEFI firmware-ben. Az alábbi parancsban a `boot-disk` és a `boot-partition-id` cserélje ki az EFI rendszerpartíció lemez- és partícióazonosítójára:

`root #` `efibootmgr --create --disk /dev/boot-disk --part boot-partition-id --loader '\EFI\systemd\shimx64.efi'  --label 'Systemd-boot via Shim' --unicode '\EFI\systemd\systemd-bootx64.efi'`

**Note**

Az importálási folyamat nem fejeződik be az operációs rendszer újraindításáig. Miután a kézikönyv összes lépését végrehajtotta, indítsa újra az operációs rendszert. A Shim betöltődik, és megtalálja a mokutil által regisztrált importálási kérelmet. A MokManager alkalmazás elindul, és kéri a jelszót, amelyet az importálási kérelem létrehozásakor állított be. Kövesse a képernyőn megjelenő utasításokat a tanúsítvány importálásának befejezéséhez, majd indítsa újra a rendszert az UEFI menübe, és engedélyezze a Secure Boot beállítást.

## Alternatíva 2: EFI Stub

A UEFI-alapú firmware-rel rendelkező számítógépes rendszerek technikailag nem igényelnek másodlagos betöltőket (például a GRUB-ot) a kernel betöltéséhez. A másodlagos betöltők azért léteznek, hogy _kiterjesszék_ az UEFI firmware funkcionalitását az indítási folyamat során. A GRUB használata (lásd az előző részt) általában egyszerűbb és robusztusabb, mivel rugalmasabb megközelítést kínál a kernelparaméterek gyors módosításához az indítási idő alatt.

Azok a rendszergazdák, akik minimális, bár merevebb megközelítést szeretnének alkalmazni a rendszer indításához, elkerülhetik a másodlagos betöltőket, és közvetlenül az [EFI stub](/wiki/EFI_stub/hu "EFI stub/hu") segítségével indíthatják a Linux kernelt.

Az [sys-boot/efibootmgr](https://packages.gentoo.org/packages/sys-boot/efibootmgr) szoftvercsomagban lévő alkalmazás egy olyan eszköz, amelyet az UEFI firmware-rel való interakcióra használnak - az operációs rendszer elsődleges betöltőjével. Általában ez azt jelenti, hogy indítási bejegyzéseket ad hozzá vagy távolít el a firmware indítható bejegyzéseinek listájából. Ezenkívül frissítheti a firmware beállításokat is, hogy a korábban indítható bejegyzésként hozzáadott Linux kernelek további opciókkal legyenek végrehajthatók. Ezek az interakciók speciális adatstruktúrákon keresztül történnek, amelyeket EFI változóknak neveznek (ezért van szükség a kernel EFI változók támogatására).

Győződjön meg arról, hogy az [EFI stub](/wiki/EFI_stub/hu "EFI stub/hu") kernel cikke át lett nézve _mielőtt_ folytatná. A kernelnek rendelkeznie kell speciális opciókkal, hogy közvetlenül indítható legyen az UEFI firmware által. Lehet, hogy szükséges a kernel újrafordítása annak érdekében, hogy ez a támogatás be legyen építve.

Érdemes átnézni a [efibootmgr](/wiki/Efibootmgr/hu "Efibootmgr/hu") cikket is a további információkért.

**Note**

Hogy újra hangsúlyozzuk, a efibootmgr _nem_ szükséges egy UEFI rendszer indításához, csupán annyira van szükség, hogy egy bejegyzést hozzáadjunk az EFI-stub kernelhez az UEFI firmware-ben. Ha megfelelően lett beépítve az EFI stub támogatással, akkor a Linux kernel önmagában is _közvetlenül_ elindítható. További kernel parancssori opciók beépíthetők a Linux kernelbe (van egy kernelbeállítás opció, amelyet CONFIG\_CMDLINE változónak neveznek). Hasonlóképpen, az initramfs támogatása is beépíthető a kernelbe. Ezeket a döntéseket a kernel fordítása előtt kell meghozni, ami statikusabb indítási beállításhoz vezet. Ezeket a döntéseket még a kernel lefordítása előtt meg kell hozni, amelynek eredményeként statikusabb rendszerbetöltési beállítás jön létre.

A kernel közvetlenül a firmware-ből való indításához a kernelképfájlnak jelen kell lennie az [EFI System Partition](/wiki/EFI_System_Partition/hu "EFI System Partition/hu") partíción. Ezt a [efistub](https://packages.gentoo.org/useflags/efistub) [USE flag/hu](/wiki/USE_flag/hu "USE flag/hu") USE jelölőzászló engedélyezésével lehet elérni a [sys-kernel/installkernel](https://packages.gentoo.org/packages/sys-kernel/installkernel) szoftvercsomagon. Tekintettel arra, hogy az EFI Stub indítás nem garantáltan működik minden UEFI rendszeren, a USE jelölőzászló stabilan le van tiltva, így a tesztelési kulcsszavakat el kell fogadni az installkernel ezen funkciójának használatához.

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

Ezután telepítse újra a [installkernel](/wiki/Installkernel/hu "Installkernel/hu") szoftvercsomagot, hozza létre a /efi/EFI/Gentoo könyvtárat, és telepítse újra a kernelt:

`root #` `emerge --ask sys-kernel/installkernel`

`root #` `mkdir -p /efi/EFI/Gentoo`

Disztribúciós kernelek esetén:

`root #` `emerge --ask --config sys-kernel/gentoo-kernel{,-bin}`

Kézzel kezelt kernelek esetén:

`root #` `make install`

Alternatív megoldásként, ha az [installkernel](/wiki/Installkernel/hu "Installkernel/hu") nincs használatban, akkor a kernelt kézzel is másolhatja az EFI rendszerpartícióra, és elnevezheti bzImage.efi néven:

`root #` `mkdir -p /efi/EFI/Gentoo
`

`root #` `cp /boot/vmlinuz-* /efi/EFI/Gentoo/bzImage.efi
`

Telepítse az efibootmgr szoftvercsomagot:

`root #` `emerge --ask sys-boot/efibootmgr`

Hozzon létre egy "Gentoo" nevű indítási bejegyzést az újonnan lefordított EFI stub kernelhez az UEFI firmware-ben:

`root #` `efibootmgr --create --disk /dev/sda --part 1 --label "gentoo" --loader "\EFI\Gentoo\bzImage.efi"`

**Note**

A visszafelé álló pervonal ( `\`) használata a könyvtárelérési útvonalak elválasztójaként szolgál. Kötelező a használata, amikor UEFI definíciókat használunk.

Amennyiben egy kezdeti RAM fájlrendszer (initramfs) is használatban van, akkor másolja azt is az EFI rendszerpartícióra, majd adja hozzá a megfelelő boot opciót:

`root #` `efibootmgr --create --disk /dev/sda --part 1 --label "gentoo" --loader "\EFI\Gentoo\bzImage.efi" --unicode "initrd=\EFI\Gentoo\initramfs.img"`

**Tip**

További kernel parancssori opciók is továbbíthatók a firmware által a kernelhez, ha azokat az initrd=... opcióval együtt adják meg, ahogy a fenti példában látható.

Ezeknek a módosításoknak az elvégzése után, amikor az számítógép újraindul, egy "gentoo" nevű boot bejegyzés lesz elérhető.

### Egységesített kernelképfájl (Unified Kernel Image)

Ha az [installkernel](/wiki/Installkernel/hu "Installkernel/hu") úgy van beállítva, hogy egységesített kernelképfájlokat készítsen és telepítsen, akkor az egységesített kernelképfájlnak már telepítve kell, hogy legyen az EFI rendszerpartíció EFI/Linux könyvtárában. Ha ez nem így van, akkor győződjön meg róla, hogy a könyvtár létezik, majd futtassa újra a kernel telepítését az útmutató korábbi részében leírtak szerint.

Az egységes telepített kernelképfájlhoz tartozó közvetlen indítási bejegyzés hozzáadásához:

`root #` `efibootmgr --create --disk /dev/sda --part 1 --label "gentoo" --loader "\EFI\Linux\gentoo-x.y.z.efi"`

## Egyéb alternatívák

Más lehetőségekért, amelyek nem szerepelnek a kézikönyvben, tekintse meg az elérhető [bootloaderek](/wiki/Bootloader/hu "Bootloader/hu") teljes listáját.

## Rendszer újraindítása

Lépjen ki a chrootolt környezetből, és válassza le az összes felcsatolt partíciót. Ezt követően írja be azt az egyetlen mágikus parancsot, amely elindítja a végső, valódi tesztet: reboot.

`(chroot) livecd #` `exit`

`livecd~#` `cd
`

`livecd~#` `umount -l /mnt/gentoo/dev{/shm,/pts,}
`

`livecd~#` `umount -R /mnt/gentoo
`

`livecd~#` `reboot`

Ne feledje el eltávolítani az Live ISO telepítőt, különben ismét elindulhat a számítógépen az újonnan telepített Gentoo rendszer helyett!

Miután újraindította a számítógépet, és belépett a frissen feltelepített Gentoo környezetben, bölcs dolog [véglegesíteni a Gentoo telepítést](/wiki/Handbook:AMD64/Installation/Finalizing/hu "Handbook:AMD64/Installation/Finalizing/hu").

[← Eszközök telepítése](/wiki/Handbook:AMD64/Installation/Tools/hu "Previous part") [Kezdőlap](/wiki/Handbook:AMD64/hu "Handbook:AMD64/hu") [Telepítés véglegesítése →](/wiki/Handbook:AMD64/Installation/Finalizing/hu "Next part")