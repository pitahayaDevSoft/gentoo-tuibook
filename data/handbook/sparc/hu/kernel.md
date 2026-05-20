# Kernel

Other languages:

- [Deutsch](/wiki/Handbook:SPARC/Installation/Kernel/de "Handbuch:SPARC/Installation/Kernel (100% translated)")
- [English](/wiki/Handbook:SPARC/Installation/Kernel "Handbook:SPARC/Installation/Kernel (100% translated)")
- [español](/wiki/Handbook:SPARC/Installation/Kernel/es "Manual de Gentoo: SPARC/Instalación/Núcleo (100% translated)")
- [français](/wiki/Handbook:SPARC/Installation/Kernel/fr "Handbook:SPARC/Installation/Kernel/fr (100% translated)")
- [italiano](/wiki/Handbook:SPARC/Installation/Kernel/it "Handbook:SPARC/Installation/Kernel/it (50% translated)")
- magyar
- [polski](/wiki/Handbook:SPARC/Installation/Kernel/pl "Handbook:SPARC/Installation/Kernel (100% translated)")
- [português do Brasil](/wiki/Handbook:SPARC/Installation/Kernel/pt-br "Handbook:SPARC/Installation/Kernel/pt-br (50% translated)")
- [русский](/wiki/Handbook:SPARC/Installation/Kernel/ru "Handbook:SPARC/Installation/Kernel (100% translated)")
- [தமிழ்](/wiki/Handbook:SPARC/Installation/Kernel/ta "கையேடு:SPARC/நிறுவல்/கர்னல் (100% translated)")
- [中文（中国大陆）‎](/wiki/Handbook:SPARC/Installation/Kernel/zh-cn "手册：SPARC/安装/配置Linux内核 (100% translated)")
- [日本語](/wiki/Handbook:SPARC/Installation/Kernel/ja "ハンドブック:SPARC/インストール/カーネル (100% translated)")
- [한국어](/wiki/Handbook:SPARC/Installation/Kernel/ko "Handbook:SPARC/Installation/Kernel/ko (100% translated)")

[SPARC kézikönyv](/wiki/Handbook:SPARC/hu "Handbook:SPARC/hu")[A Gentoo Linux telepítése](/wiki/Handbook:SPARC/Full/Installation/hu "Handbook:SPARC/Full/Installation/hu")[A telepítésről](/wiki/Handbook:SPARC/Installation/About/hu "Handbook:SPARC/Installation/About/hu")[Telepítőképfájl kiválasztása](/wiki/Handbook:SPARC/Installation/Media/hu "Handbook:SPARC/Installation/Media/hu")[Hálózat beállítása](/wiki/Handbook:SPARC/Installation/Networking/hu "Handbook:SPARC/Installation/Networking/hu")[Adathordozók előkészítése](/wiki/Handbook:SPARC/Installation/Disks/hu "Handbook:SPARC/Installation/Disks/hu")[Stage fájl](/wiki/Handbook:SPARC/Installation/Stage/hu "Handbook:SPARC/Installation/Stage/hu")[Alaprendszer telepítése](/wiki/Handbook:SPARC/Installation/Base/hu "Handbook:SPARC/Installation/Base/hu")Kernel beállítása[Rendszer beállítása](/wiki/Handbook:SPARC/Installation/System/hu "Handbook:SPARC/Installation/System/hu")[Eszközök telepítése](/wiki/Handbook:SPARC/Installation/Tools/hu "Handbook:SPARC/Installation/Tools/hu")[Bootloader beállítása](/wiki/Handbook:SPARC/Installation/Bootloader/hu "Handbook:SPARC/Installation/Bootloader/hu")[Telepítés véglegesítése](/wiki/Handbook:SPARC/Installation/Finalizing/hu "Handbook:SPARC/Installation/Finalizing/hu")[Munka a Gentoo rendszerrel](/wiki/Handbook:SPARC/Full/Working/hu "Handbook:SPARC/Full/Working/hu")[Portage bemutatása](/wiki/Handbook:SPARC/Working/Portage/hu "Handbook:SPARC/Working/Portage/hu")[USE jelölőzászlók](/wiki/Handbook:SPARC/Working/USE/hu "Handbook:SPARC/Working/USE/hu")[Portage jellemzői](/wiki/Handbook:SPARC/Working/Features/hu "Handbook:SPARC/Working/Features/hu")[Init-szkript rendszer](/wiki/Handbook:SPARC/Working/Initscripts/hu "Handbook:SPARC/Working/Initscripts/hu")[Környezeti változók](/wiki/Handbook:SPARC/Working/EnvVar/hu "Handbook:SPARC/Working/EnvVar/hu")[Munka a Portage szoftvercsomag-kezelővel](/wiki/Handbook:SPARC/Full/Portage/hu "Handbook:SPARC/Full/Portage/hu")[Fájlok és könyvtárak](/wiki/Handbook:SPARC/Portage/Files/hu "Handbook:SPARC/Portage/Files/hu")[Változók](/wiki/Handbook:SPARC/Portage/Variables/hu "Handbook:SPARC/Portage/Variables/hu")[Szoftverágak keverése](/wiki/Handbook:SPARC/Portage/Branches/hu "Handbook:SPARC/Portage/Branches/hu")[További eszközök](/wiki/Handbook:SPARC/Portage/Tools/hu "Handbook:SPARC/Portage/Tools/hu")[Egyéni szoftvercsomag-tárolók](/wiki/Handbook:SPARC/Portage/CustomTree/hu "Handbook:SPARC/Portage/CustomTree/hu")[Fejlett funkciók](/wiki/Handbook:SPARC/Portage/Advanced/hu "Handbook:SPARC/Portage/Advanced/hu")[Hálózat beállítása OpenRC init-rendszeren](/wiki/Handbook:SPARC/Full/Networking/hu "Handbook:SPARC/Full/Networking/hu")[Munka elkezdése](/wiki/Handbook:SPARC/Networking/Introduction/hu "Handbook:SPARC/Networking/Introduction/hu")[Fejlett beállítások](/wiki/Handbook:SPARC/Networking/Advanced/hu "Handbook:SPARC/Networking/Advanced/hu")[Moduláris hálózat](/wiki/Handbook:SPARC/Networking/Modular/hu "Handbook:SPARC/Networking/Modular/hu")[Vezeték nélküli (Wi-Fi)](/wiki/Handbook:SPARC/Networking/Wireless/hu "Handbook:SPARC/Networking/Wireless/hu")[Funkcionalitás hozzáadása](/wiki/Handbook:SPARC/Networking/Extending/hu "Handbook:SPARC/Networking/Extending/hu")[Dinamikus menedzsment](/wiki/Handbook:SPARC/Networking/Dynamic/hu "Handbook:SPARC/Networking/Dynamic/hu")

## Contents

- [1Opcionális: Firmware és/vagy mikrokód telepítése](#Opcion.C3.A1lis:_Firmware_.C3.A9s.2Fvagy_mikrok.C3.B3d_telep.C3.ADt.C3.A9se)
  - [1.1Firmware](#Firmware)
    - [1.1.1Javasolt: Linux Firmware](#Javasolt:_Linux_Firmware)
      - [1.1.1.1Firmware betöltés (Firmware Loading)](#Firmware_bet.C3.B6lt.C3.A9s_.28Firmware_Loading.29)
- [2sys-kernel/installkernel](#sys-kernel.2Finstallkernel)
  - [2.1Bootloader](#Bootloader)
    - [2.1.1GRUB](#GRUB)
    - [2.1.2Hagyományos elrendezés, egyéb bootloaderek (pl. (e)lilo, syslinux stb.)](#Hagyom.C3.A1nyos_elrendez.C3.A9s.2C_egy.C3.A9b_bootloaderek_.28pl._.28e.29lilo.2C_syslinux_stb..29)
  - [2.2Az initramfs létrehozása](#Az_initramfs_l.C3.A9trehoz.C3.A1sa)
    - [2.2.1Chroot detection](#Chroot_detection)
- [3Kernel beállítása és a kernel forráskódjának lefordítása futtatható bináris kódra](#Kernel_be.C3.A1ll.C3.ADt.C3.A1sa_.C3.A9s_a_kernel_forr.C3.A1sk.C3.B3dj.C3.A1nak_leford.C3.ADt.C3.A1sa_futtathat.C3.B3_bin.C3.A1ris_k.C3.B3dra)
  - [3.1Terjesztési (Disztribúció) kernelek](#Terjeszt.C3.A9si_.28Disztrib.C3.BAci.C3.B3.29_kernelek)
    - [3.1.1Opcionális: Aláírt kernelmodulok](#Opcion.C3.A1lis:_Al.C3.A1.C3.ADrt_kernelmodulok)
    - [3.1.2Egy terjesztési (disztribúció) kernel telepítése](#Egy_terjeszt.C3.A9si_.28disztrib.C3.BAci.C3.B3.29_kernel_telep.C3.ADt.C3.A9se)
    - [3.1.3Frissítés és takarítás](#Friss.C3.ADt.C3.A9s_.C3.A9s_takar.C3.ADt.C3.A1s)
    - [3.1.4Telepítés utáni/frissítési feladatok](#Telep.C3.ADt.C3.A9s_ut.C3.A1ni.2Ffriss.C3.ADt.C3.A9si_feladatok)
      - [3.1.4.1Initramfs vagy az Egységes kernelképfájl manuális újraépítése](#Initramfs_vagy_az_Egys.C3.A9ges_kernelk.C3.A9pf.C3.A1jl_manu.C3.A1lis_.C3.BAjra.C3.A9p.C3.ADt.C3.A9se)
  - [3.2Alternatíva: Kézi beállítás](#Alternat.C3.ADva:_K.C3.A9zi_be.C3.A1ll.C3.ADt.C3.A1s)
  - [3.3Kernel forráskódjainak a telepítése](#Kernel_forr.C3.A1sk.C3.B3djainak_a_telep.C3.ADt.C3.A9se)
    - [3.3.1A modprobed-db folyamat](#A_modprobed-db_folyamat)
    - [3.3.2Manuális folyamat](#Manu.C3.A1lis_folyamat)
      - [3.3.2.1Szükséges opciók engedélyezése](#Sz.C3.BCks.C3.A9ges_opci.C3.B3k_enged.C3.A9lyez.C3.A9se)
    - [3.3.3Opcionális: Aláírt kernel modulok](#Opcion.C3.A1lis:_Al.C3.A1.C3.ADrt_kernel_modulok)
  - [3.4Kódfordítás és telepítés](#K.C3.B3dford.C3.ADt.C3.A1s_.C3.A9s_telep.C3.ADt.C3.A9s)

## Opcionális: Firmware és/vagy mikrokód telepítése

### Firmware

#### Javasolt: Linux Firmware

Sok operációs rendszeren bizonyos hardverek működéséhez nem-FOSS firmware szoftverre van szükség. A [sys-kernel/linux-firmware](https://packages.gentoo.org/packages/sys-kernel/linux-firmware) szoftvercsomag sok (de nem minden) eszközhöz tartalmaz firmware szoftvert.

**Tip**

A legtöbb vezeték nélküli kártya és GPU működéséhez szükség van firmware szoftverre.

`root #` `emerge --ask sys-kernel/linux-firmware`

**Note**

Egyes firmware szoftvercsomagok telepítése gyakran megköveteli a kapcsolódó firmware licenc elfogadását. Ha szükséges, látogasson el a Kézikönyv [licenckezelési szekciójához](/wiki/Handbook:SPARC/Working/Portage/hu#Licenses "Handbook:SPARC/Working/Portage/hu"), hogy segítséget kapjon a licenc elfogadásában.

##### Firmware betöltés (Firmware Loading)

A firmware fájlokat általában akkor töltik be, amikor a hozzájuk tartozó kernelmodult betöltik. Ez azt jelenti, hogy a firmware szoftvert be kell építeni a kernelbe a **CONFIG\_EXTRA\_FIRMWARE** használatával, ha a kernelmodult _Y_-ra állítják az _M_ helyett. A legtöbb esetben egy firmware-t igénylő modul beépítése bonyolíthatja vagy meghiúsíthatja a betöltést.

## sys-kernel/installkernel

[Installkernel](/wiki/Installkernel/hu "Installkernel/hu") használható a kerneltelepítés, az [initramfs](/wiki/Initramfs "Initramfs") generálás, az [unified kernel image](/wiki/Unified_kernel_image "Unified kernel image") generálás és/vagy a bootloader beállítás automatizálására, többek között. A [sys-kernel/installkernel](https://packages.gentoo.org/packages/sys-kernel/installkernel) két utat kínál ennek elérésére: a Debianból származó hagyományos installkernel és a [systemd](/wiki/Systemd/hu "Systemd/hu") kernel-install megoldását. Hogy melyiket válassza, többek között a rendszer bootloaderétől függ. Alapértelmezés szerint a systemd profilokon a systemd kernel-install van használatban, míg más profilokon a hagyományos installkernel az alapértelmezett.

### Bootloader

Most itt az ideje eldönteni, hogy a felhasználó melyik bootloadert szeretné használni a rendszerhez. Ha nem biztos benne, akkor kövesse az alábbi 'Hagyományos elrendezés' alfejezetet.

**Important**

Only one selection is required in the following subsection, if unsure of which to use then go with the first listed for now. It's always possible to switch at a later date if required.

#### GRUB

A GRUB felhasználók használhatják a systemd kernel-install vagy a hagyományos Debian installkernel megoldását. A [systemd](https://packages.gentoo.org/useflags/systemd) [USE flag/hu](/wiki/USE_flag/hu "USE flag/hu") USE jelölőzászló lehetővé teszi az ezek közötti váltást. Ahhoz, hogy a kernel telepítésekor a grub-mkconfig automatikusan fusson, engedélyezze a [grub](https://packages.gentoo.org/useflags/grub) [USE flag/hu](/wiki/USE_flag/hu "USE flag/hu") [USE jelölőzászló](/wiki/USE_flag/hu "USE flag/hu") opciót.

**Note**

GRUB requires kernels to be installed to /boot.

FILE **`/etc/portage/package.use/installkernel`**

```
sys-kernel/installkernel grub

```

`root #` `emerge --ask sys-kernel/installkernel`

`root #` `mkdir -p /etc/cmdline.d`

`root #` `ln -s /etc/kernel/cmdline /etc/cmdline.d/00-installkernel.conf`

`root #` `emerge --ask sys-kernel/installkernel`

}}

**Note**

systemd-boot requires kernels to be installed to /efi.

**Note**

When [app-emulation/virt-firmware](https://packages.gentoo.org/packages/app-emulation/virt-firmware) is used to configure the UEFI ensure that the kernel-bootcfg-boot-successful service is enabled before attempting to install the kernel. This implementation of EFIstub booting is the default for systemd systems.

`root #` `systemctl enable kernel-bootcfg-boot-successful.service`

**Note**

EFIstub requires kernels to be installed to /efi.

#### Hagyományos elrendezés, egyéb bootloaderek (pl. (e)lilo, syslinux stb.)

A hagyományos /boot elrendezés (például (e)LILO, syslinux stb.) az alapértelmezett, ha a [grub](https://packages.gentoo.org/useflags/grub) [USE flag/hu](/wiki/USE_flag/hu "USE flag/hu"), [systemd-boot](https://packages.gentoo.org/useflags/systemd-boot) [USE flag/hu](/wiki/USE_flag/hu "USE flag/hu"), [efistub](https://packages.gentoo.org/useflags/efistub) [USE flag/hu](/wiki/USE_flag/hu "USE flag/hu") és [uki](https://packages.gentoo.org/useflags/uki) [USE flag/hu](/wiki/USE_flag/hu "USE flag/hu") USE flag-ek **nincsenek** engedélyezve. További lépések nem szükségesek.

### Az initramfs létrehozása

Egy **i** nitial **ram**-alapú **f** ájl **s** ziszéma, vagyis [initramfs](/wiki/Initramfs "Initramfs"), szükséges lehet az operációs rendszer indításához. Számos esetben szükség lehet erre, de a gyakori esetek közé tartoznak:

- Kernelek, ahol a tárhely-illesztőprogramok vagy a fájlrendszer-illesztőprogramok modulok.
- Elrendezések, ahol a /usr/ vagy a /var/ külön partíciókon található.
- Titkosított gyökérfájlrendszerek.

**Tip**

A [Distribution kernelek](/wiki/Project:Distribution_Kernel "Project:Distribution Kernel") initramfs fájllal való használatra készültek, mivel számos tárhely- és fájlrendszer-illesztőprogram modulként van felépítve.

A gyökérfájlrendszer felcsatolása mellett egy initramfs fájl más feladatokat is elláthat, például:

- A fájlrendszer konzisztenciájának az ellenőrzése, vagyis a **f** ile **s** ystem **c** onsistency chec **k** fsck futtatása, amely egy olyan szoftver, amely a fájlrendszer konzisztenciájának az ellenőrzésére és javítására szolgál például egy nem megfelelő operációs rendszer leállítás esetén.
- Helyreállítási környezet biztosítása késői indítási hibák esetén.

Az [Installkernel](/wiki/Installkernel/hu "Installkernel/hu") automatikusan létrehozhat egy initramfs fájlt a kernel telepítésekor, ha a [dracut](https://packages.gentoo.org/useflags/dracut) [USE flag/hu](/wiki/USE_flag/hu "USE flag/hu") vagy [ugrd](https://packages.gentoo.org/useflags/ugrd) [USE flag/hu](/wiki/USE_flag/hu "USE flag/hu") USE jelölőzászló engedélyezve van:

FILE **`/etc/portage/package.use/installkernel`**

```
sys-kernel/installkernel dracut

```

#### Chroot detection

Bootloaders such as [systemd-boot](/wiki/Systemd/systemd-boot "Systemd/systemd-boot") and [EFI stub](/wiki/EFI_stub/hu "EFI stub/hu") use the kernel arguments of the running system as set in /proc/cmdline by default. Because of the wide range of ways Gentoo can be installed users will randomly get tripped up by this.

To help solve any issues this may cause, [sys-kernel/installkernel](https://packages.gentoo.org/packages/sys-kernel/installkernel) will check if it is running in a chroot and fail if the kernel command line is not explicitly configured. Otherwise the bootloader would use the install media's boot arguments which would lead to boot failure.

One way to satisfy [sys-kernel/installkernel](https://packages.gentoo.org/packages/sys-kernel/installkernel) is by using Dracut's config file to set the root partition UUID as shown below, or alternatively for more information on what this check helps with and different ways to configure it, see [Installkernel#Install\_chroot\_check](/wiki/Installkernel#Install_chroot_check.2Fhu "Installkernel").

`root #` `mkdir /etc/dracut.conf.d`

`root #` `blkid`

```
/dev/sda3: UUID="2122cd72-94d7-4dcc-821e-3705926deecc"
```

In the above example, the root partition is /dev/sda3 and the UUID is 2122cd72-94d7-4dcc-821e-3705926deecc.
The dracut config file would then look like:

FILE **`/etc/dracut.conf.d/00-installkernel.conf`**

```
kernel_cmdline=" root=UUID=2122cd72-94d7-4dcc-821e-3705926deecc " # Note leading and trailing spaces

```

`root #` `emerge --ask sys-kernel/installkernel`

## Kernel beállítása és a kernel forráskódjának lefordítása futtatható bináris kódra

**Tip**

Bölcs lépés lehet az első indításkor a dist-kernel használata, mivel ez egy nagyon egyszerű módszert kínál a rendszerproblémák és kernelbeállítás problémák kizárására. Egy ismerten működő kernel mindig rendelkezésre állása gyorsíthatja a hibakeresést, és csökkentheti a frissítés miatti szorongást, hogy az operációs rendszer esetleg nem indul el többé.

A common misconception is that a manually compiled kernel will use a lot less RAM than a pre configured distribution kernel. Due to the modular nature of the Linux kernel, only what is needed by the system is loaded and in most cases resulting in similar memory usage.

**Important**

Only one selection is required in the following subsection, if unsure of which to use then go with the first listed for now. It's always possible to switch at a later date if required.

Rangsorolva a legkevésbé érintettől a leginkább érintettig:

[Teljes automatizálási megközelítés: Terjesztési kernelek](/wiki/Handbook:SPARC/Installation/Kernel/hu#Distribution_kernels "Handbook:SPARC/Installation/Kernel/hu"): A [Terjesztési Kernel](/wiki/Project:Distribution_Kernel "Project:Distribution Kernel") beállítására, automatikus felépítésére és telepítésére szolgál, beleértve a Linux kernelt, a hozzá tartozó modulokat és (alapértelmezés szerint engedélyezett, de választható) egy initramfs fájlt. A jövőbeli kernel frissítések teljesen automatizáltak, mivel ezek a szoftvercsomag-kezelőn keresztül kezelhetők, akárcsak bármely más rendszer-szoftvercsomag. Szükség esetén lehetséges egy [egyedi kernel beállítás fájl](/wiki/Project:Distribution_Kernel#Modifying_kernel_configuration "Project:Distribution Kernel") biztosítása. Ez a legkevésbé bonyolult eljárás, és tökéletes az új Gentoo felhasználók számára, mivel azonnali működést kínál, és minimális rendszergazdai beavatkozást igényel.

[Teljes manuális megközelítés](/wiki/Handbook:SPARC/Installation/Kernel/hu#Alternative:_Manual_configuration "Handbook:SPARC/Installation/Kernel/hu")Az új kernelforrások az operációs rendszer szoftvercsomag-kezelőjével vannak telepítve. A kernelt beállítása, létrehozása és telepítése manuálisan történik a eselect kernel és számos make parancs segítségével. A későbbi kernel frissítések során ismét végre kell hajtani a beállítás, létrehozás és telepítés manuális folyamatát. Ez a legösszetettebb folyamat, de maximális irányítást biztosít a kernel frissítési folyamat felett.[Hibrid megközelítés: Genkernel](/wiki/Handbook:SPARC/Installation/Kernel/hu#Alternative:_Genkernel "Handbook:SPARC/Installation/Kernel/hu")Itt a hibrid kifejezést használjuk, de vegye figyelembe, hogy a dist-kernel és a kézi források egyaránt tartalmaznak módszereket ugyanazon cél elérésére. Az új kernelforráskódok az operációs rendszer szoftvercsomag-kezelőjén keresztül telepíthetőek. A rendszergazdák használhatják a Gentoo genkernel eszközt a Linux kernel, a hozzá tartozó modulok és (opcionálisan, de alapértelmezés szerint _**nem**_ engedélyezett) egy initramfs fájl beállítására, létrehozására és telepítésére. Szükség esetén egyedi kernelbeállítás-fájl is biztosítható. A jövőbeli kernel beállításához, forráskód fordításhoz és telepítéshez a rendszergazda részvétele szükséges, például a eselect kernel, genkernel és esetleg más parancsok futtatásával minden frissítéshez. Ez az opció csak azoknak a felhasználóknak ajánlott, akik tudják, hogy szükségük van a genkernel szoftverre.

Az összes disztribúció alapját a Linux kernel képezi. Ez a réteg található a felhasználói programok és az operációs rendszer hardverei között. Bár a kézikönyv számos lehetséges kernel forrást kínál a felhasználóknak, egy átfogóbb lista részletesebb leírásokkal elérhető a [Kernel áttekintő oldalon](/wiki/Kernel/Overview "Kernel/Overview").

**Tip**

A kernel telepítési feladatok, mint például a kernelképfájl másolása a /boot könyvtárba vagy az [EFI rendszerpartícióra](/wiki/EFI_System_Partition/hu "EFI System Partition/hu"), egy [initramfs](/wiki/Initramfs "Initramfs") és/vagy [egységes kernelképfájl](/wiki/Unified_Kernel_Image "Unified Kernel Image") létrehozása, a bootloader beállításának a frissítése automatizálható az [installkernel](/wiki/Installkernel/hu "Installkernel/hu") segítségével. A felhasználónak érdemes lehet beállítani és telepíteni a [sys-kernel/installkernel](https://packages.gentoo.org/packages/sys-kernel/installkernel) szoftvercsomagot, mielőtt folytatná. További információkért tekintse meg az alábbi [Kernel telepítése szekciót](/wiki/Handbook:SPARC/Installation/Kernel#Kernel_installation.2Fhu "Handbook:SPARC/Installation/Kernel").

### Terjesztési (Disztribúció) kernelek

_[Terjesztési (disztribúció) kernelek](/wiki/Project:Distribution_Kernel "Project:Distribution Kernel")_ ebuild-ek, amelyek lefedik a kernel kicsomagolásának, beállításának, kódfordításának és telepítésének teljes folyamatát. Ennek a módszernek az elsődleges előnye, hogy a szoftvercsomag-kezelő részeként a @world frissítés során a kernelek új verziókra frissülnek. Ez nem igényel több beavatkozást, mint egy emerge parancs futtatása. A terjesztési (disztribúció) kernelek alapértelmezett beállítása a legtöbb hardvert támogatja, azonban két mechanizmust kínálnak a testreszabáshoz: savedconfig és config snippets. További részletekért, a beállítással kapcsolatban, tekintse meg a [projekt oldalt](/wiki/Project:Distribution_Kernel#Modifying_kernel_configuration "Project:Distribution Kernel").

##### Opcionális: Aláírt kernelmodulok

A kész terjesztési (disztribúció) kernelben ([sys-kernel/gentoo-kernel-bin](https://packages.gentoo.org/packages/sys-kernel/gentoo-kernel-bin)) található kernel modulok már alá vannak írva. A forráskódból fordított kernelek moduljainak aláírásához engedélyezze a [modules-sign](https://packages.gentoo.org/useflags/modules-sign) [USE flag/hu](/wiki/USE_flag/hu "USE flag/hu") USE jelölőzászlót, és opcionálisan adja meg, hogy melyik kulcsot kívánja használni az aláíráshoz az [/etc/portage/make.conf](/wiki//etc/portage/make.conf/hu "/etc/portage/make.conf/hu") fájlban.

FILE **`/etc/portage/make.conf`** **Kernelmodul-aláírás engedélyezése**

```
USE="modules-sign"

# Opcionálisan, egyéni aláírási kulcsok használatához.
MODULES_SIGN_KEY="/path/to/kernel_key.pem"
MODULES_SIGN_CERT="/path/to/kernel_key.pem" # Kizárólag akkor szükséges, ha a MODULES_SIGN_KEY a tanúsítványt se nem tartalmazza.
MODULES_SIGN_HASH="sha512" # Alapértelmezés szerint sha512.

```

Ha a MODULES\_SIGN\_KEY nincs megadva, akkor a kernel építési rendszer egy kulcsot fog generálni, amely a /usr/src/linux-x.y.z/certs könyvtárban kerül tárolásra. Ajánlott manuálisan generálni egy kulcsot annak biztosítása érdekében, hogy minden kernel kiadásnál ugyanaz legyen a kulcs. Egy kulcsot a következő parancs futtatásával lehet generálni:

`root #` `openssl req -new -nodes -utf8 -sha256 -x509 -outform PEM -out kernel_key.pem -keyout kernel_key.pem`

**Note**

A MODULES\_SIGN\_KEY és a MODULES\_SIGN\_CERT lehetnek különböző fájlok. Ebben a példában az OpenSSL által generált pem fájl tartalmazza mind a kulcsot, mind a hozzá tartozó tanúsítványt, így mindkét változót ugyanarra az értékre állítjuk be.

Az OpenSSL néhány kérdést fog feltenni a kulcsot generáló felhasználóról, ajánlott ezeket a kérdéseket a lehető legrészletesebben megválaszolni.

Tárolja a kulcsot egy biztonságos helyen, legalábbis úgy, hogy a kulcsot csak a root felhasználó olvashassa. Ezt a következő parancs futtatásával ellenőrizheti:

`root #` `ls -l kernel_key.pem`

```
 -r-------- 1 root root 3164 Jan  4 10:38 kernel_key.pem
```

Ha ez bármi mást ad vissza, mint a fentieket, akkor javítsa a jogosultságokat a következő parancs futtatásával:

`root #` `chown root:root kernel_key.pem`

`root #` `chmod 400 kernel_key.pem`

#### Egy terjesztési (disztribúció) kernel telepítése

Egy Gentoo foltokkal ellátott kernel fordításához a forráskódból írja be a következő parancsot:

`root #` `emerge --ask sys-kernel/gentoo-kernel`

Azok a rendszergazdák, akik szeretnék elkerülni a kernel forráskódok helyi lefordítását, használhatnak előre binárisra lefordított kernel képfájlokat:

`root #` `emerge --ask sys-kernel/gentoo-kernel-bin`

**Important**

A _[Terjesztési Kernelek](/wiki/Project:Distribution_Kernel "Project:Distribution Kernel")_, mint például a [sys-kernel/gentoo-kernel](https://packages.gentoo.org/packages/sys-kernel/gentoo-kernel) és a [sys-kernel/gentoo-kernel-bin](https://packages.gentoo.org/packages/sys-kernel/gentoo-kernel-bin), alapértelmezés szerint elvárják, hogy egy [initramfs](/wiki/Handbook:SPARC/Installation/Kernel#Initramfs.2Fhu "Handbook:SPARC/Installation/Kernel") fájl mellett legyenek telepítve. Mielőtt az emerge futtatásra kerülne a kernel telepítéséhez, a felhasználóknak biztosítaniuk kell, hogy a [sys-kernel/installkernel](https://packages.gentoo.org/packages/sys-kernel/installkernel) be legyen állítva egy initramfs fájlgenerátor (például [Dracut](/wiki/Dracut "Dracut")) használatára, ahogy az a [installkernel szekcióban](/wiki/Handbook:SPARC/Installation/Kernel#Initramfs.2Fhu "Handbook:SPARC/Installation/Kernel") le van írva.

#### Frissítés és takarítás

Miután a kernel telepítve van, a jövőben a szoftvercsomag-kezelő automatikusan frissíteni fogja azt az újabb verziókra. Az előző verziók megmaradnak, amíg a szoftvercsomag-kezelőt nem kérjük a régi csomagok eltávolítására. Az adathordozónk területének a felaszabadítása érdekében a régi szoftvercsomagokat a --depclean opcióval futtatott emerge parancs időnkénti használatával lehet eltávolítani:

`root #` `emerge --depclean`

Alternatív megoldásként a régi kernelverziók kifejezett eltávolításának érdekében futtassuk a következő parancsot:

`root #` `emerge --prune sys-kernel/gentoo-kernel sys-kernel/gentoo-kernel-bin`

**Tip**

Az emerge alapértelmezés szerint csak a kernel létrehozási könyvtárát távolítja el. Nem távolítja el a kernel modulokat, sem a telepített kernelképfájlokat. A régi kernelfájlok teljes tisztításához az [app-admin/eclean-kernel](https://packages.gentoo.org/packages/app-admin/eclean-kernel) eszköz használható.

#### Telepítés utáni/frissítési feladatok

Egy terjesztési kernel frissítése képes automatikusan újra létrehozni a más szoftvercsomagok által telepített külső kernelmodulokat (például: [sys-fs/zfs-kmod](https://packages.gentoo.org/packages/sys-fs/zfs-kmod) vagy [x11-drivers/nvidia-drivers](https://packages.gentoo.org/packages/x11-drivers/nvidia-drivers)). Ez az automatizált viselkedés a [dist-kernel](https://packages.gentoo.org/useflags/dist-kernel) [USE flag/hu](/wiki/USE_flag/hu "USE flag/hu") USE jelölőzászló engedélyezésével válik elérhetővé. Szükség esetén ugyanez a jelölőzászló az [initramfs](/wiki/Initramfs "Initramfs") fájlnak az újragenerálását is elindítja.

Amennyiben Ön terjesztési kernelt használ, akkor erősen ajánlott, hogy ezt a jelölőzászlót globálisan engedélyezze a /etc/portage/make.conf fájlban:

FILE **`/etc/portage/make.conf`** **USE=dist-kernel engedélyezése**

```
USE="dist-kernel"

```

##### Initramfs vagy az Egységes kernelképfájl manuális újraépítése

Ha szükséges, akkor indítsa el manuálisan az újraépítéseket a kernel frissítése után a következő parancs végrehajtásával:

`root #` `emerge --ask @module-rebuild`

Ha bármilyen kernel modulokra (pl. ZFS) szükség van a korai indítás (early boot) során, akkor az initramfs újraépítését a következő parancs futtatásával végezze el:

`root #` `emerge --config sys-kernel/gentoo-kernel
`

`root #` `emerge --config sys-kernel/gentoo-kernel-bin
`

A Distribution Kernel sikeres telepítése után itt az ideje továbblépni a következő szakaszra: [Rendszer beállítása](/wiki/Handbook:SPARC/Installation/System/hu "Handbook:SPARC/Installation/System/hu").

### Alternatíva: Kézi beállítás

### Kernel forráskódjainak a telepítése

A(z) sparc alapú operációs rendszerek kernelének kódfordításakor és telepítésekor a Gentoo a [sys-kernel/gentoo-sources](https://packages.gentoo.org/packages/sys-kernel/gentoo-sources) szoftvercsomagot ajánlja.

Válassza ki a megfelelő kernelforrást, és telepítse az emerge segítségével:

`root #` `emerge --ask sys-kernel/gentoo-sources`

Ez telepíti a Linux kernelforráskódokat a /usr/src/ könyvtárba, a megadott kernel verziót használva az teljes elérési út nevében. Nem fog szimbolikus hivatkozást létrehozni magától anélkül, hogy a [symlink](https://packages.gentoo.org/useflags/symlink) [USE flag/hu](/wiki/USE_flag/hu "USE flag/hu") USE flag engedélyezve lenne a kiválasztott kernelforráskód szoftvercsomagon.

Szokás, hogy egy /usr/src/linux szimbolikus hivatkozás fenn van tartva, amely az éppen futó kernelnek megfelelő forrásokra mutat. Azonban ez a szimbolikus hivatkozás alapértelmezés szerint nem lesz létrehozva. Egy egyszerű módja a szimbolikus hivatkozás létrehozásának az eselect kernelmodul használata.

További információkért a szimbolikus hivatkozás céljáról és kezeléséről, kérjük, olvassa el a [Kernel/Upgrade](/wiki/Kernel/Upgrade/hu "Kernel/Upgrade/hu") oldalt.

Először listázza ki az összes, már telepített kernelt:

`root #` `eselect kernel list`

```
Available kernel symlink targets:
  [1]   linux-6.19.3-gentoo

```

Ahhoz, hogy létrehozzon egy linux nevű szimbolikus hivatkozást, használja a következőt:

`root #` `eselect kernel set 1`

`root #` `ls -l /usr/src/linux`

```
lrwxrwxrwx    1 root   root    12 Oct 13 11:04 /usr/src/linux -> linux-6.19.3-gentoo

```

Gyakran a rendszergazda egyik legnehezebb feladatának tartják a kernelt manuális úton történő beállítását. Ez azonban kevésbé igaz - néhány kernelbeállítás után már senki sem emlékszik arra, hogy nehéz volt! Egy Gentoo-felhasználónak két módja van a manuális kernelkezelésre, amelyek az alábbiakban találhatók:

**Important**

Only one selection is required in the following subsection, if unsure of which to use then go with the first listed for now. It's always possible to switch at a later date if required.

#### A modprobed-db folyamat

A nagyon egyszerű módja a kernel kezelésének, ha először telepíti a [sys-kernel/gentoo-kernel-bin](https://packages.gentoo.org/packages/sys-kernel/gentoo-kernel-bin) szoftvercsomagot, és használja a [sys-kernel/modprobed-db](https://packages.gentoo.org/packages/sys-kernel/modprobed-db) eszközt az operációs rendszer igényeinek információgyűjtésére. A modprobed-db egy olyan eszköz, amely az operációs rendszer életciklusa alatt crontab segítségével figyeli az operációs rendszert, és hozzáadja az összes eszköz moduljait annak érdekében, hogy minden, amire a felhasználónak szüksége van, támogatás legyen. Például, ha egy Xbox vezérlőt adnak hozzá a telepítés után, akkor a modprobed-db hozzáadja a modulokat, amelyeket a következő kernel újrafordításakor létre kell majd hozni. Erről több információt talál a [Modprobed-db](/wiki/Modprobed-db "Modprobed-db") cikkben.

For now please follow installing [sys-kernel/gentoo-kernel-bin](https://packages.gentoo.org/packages/sys-kernel/gentoo-kernel-bin) via [Distribution\_kernels](/wiki/Handbook:SPARC/Installation/Kernel#Distribution_kernels.2Fhu "Handbook:SPARC/Installation/Kernel").

Next, install [sys-kernel/modprobed-db](https://packages.gentoo.org/packages/sys-kernel/modprobed-db):

`root #` `emerge --ask sys-kernel/modprobed-db`

Please watch out for further steps related to modprobed-db in the Handbook.

More on this topic can be found in the [Modprobed-db](/wiki/Modprobed-db "Modprobed-db") article.
}}

#### Manuális folyamat

Ez a módszer lehetővé teszi a felhasználó számára, hogy teljes mértékben ellenőrizze, hogy miként épül fel a kernel, anélkül hogy jelentős segítséget venne igénybe külső eszközöktől. Egyesek ezt úgy értelmezhetik, hogy szándékosan nehezítik meg a folyamatot.

Azonban ezzel a választással egy dolog igaz: elengedhetetlen, hogy ismerjük az operációs rendszert, amikor a kernelt kézzel mi magunk állítjuk be. A legtöbb információ beszerezhető a [sys-apps/pciutils](https://packages.gentoo.org/packages/sys-apps/pciutils) szoftvercsomag telepítésével, amely tartalmazza az lspci parancsot:

`root #` `emerge --ask sys-apps/pciutils`

**Note**

A chroot környezetben biztonságosan figyelmen kívül hagyhatjuk a pcilib figyelmeztetéseket (például: _pcilib: cannot open /sys/bus/pci/devices_), amelyeket az lspci parancs okozhat.

Egy másik rendszerinformáció-forrás az lsmod parancs futtatása, hogy lássuk, milyen kernelmodulokat használ a telepítő USB-pendrive/DVD adathordozó, mivel ez jó tippet adhat arról, hogy mit kell engedélyezni.

Most lépjen a kernel forráskódjának a könyvtárába.

`root #` `cd /usr/src/linux
`

**Tip**

A kernelhez elérhető összes make argumentum megtekintéséhez futtassa a `make help` parancsot.

A kernel rendelkezik egy módszerrel az installcd adathordozón jelenleg használt modulok automatikus felismerésére, ami nagyszerű kiindulópontot nyújt a felhasználó számára saját beállításának az elkészítéséhez. Ez a következő módon hívható meg:

`root #` `make localmodconfig`

Most már ideje a beállítást elvégezni az nconfig használatával:

`root #` `make nconfig`

Now it's time to decide if modules signing is required in the steps listed in [here](/wiki/Handbook:SPARC/Installation/Kernel#Optional:_Signed_kernel_modules_2.2Fhu "Handbook:SPARC/Installation/Kernel")

If not, proceed to building described [here](/wiki/Handbook:SPARC/Installation/Kernel#Compiling_and_installing.2Fhu "Handbook:SPARC/Installation/Kernel")

###### Szükséges opciók engedélyezése

#### Opcionális: Aláírt kernel modulok

A kernelmodulok automatikus aláírásához engedélyezze a CONFIG\_MODULE\_SIG\_ALL opciót:

KERNEL **Kernelmodulok aláírása CONFIG\_MODULE\_SIG\_ALL**

```
[*] Enable loadable module support
  -*-   Module signature verification
    [*]     Automatically sign all modules
    Which hash algorithm should modules be signed with? (Sign modules with SHA-512) --->

```

Ha szükséges, akkor változtassa meg opcionálisan a hash algoritmust.

Annak érdekében, hogy minden modul érvényes aláírással legyen ellátva, engedélyezze a CONFIG\_MODULE\_SIG\_FORCE opciót is:

KERNEL **Aláírt kernelmodulok érvényesítése CONFIG\_MODULE\_SIG\_FORCE**

```
[*] Enable loadable module support
  -*-   Module signature verification
    [*]     Require modules to be validly signed
    [*]     Automatically sign all modules
    Which hash algorithm should modules be signed with? (Sign modules with SHA-512) --->

```

Egy egyedi kulcs használatához adja meg ennek a kulcsnak a helyét a CONFIG\_MODULE\_SIG\_KEY opcióban. Ha nincs megadva, akkor a kernel build rendszer generál egy kulcsot. Ajánlott, hogy manuálisan generáljon egyet. Ezt Ön a következő módon teheti meg:

`root #` `openssl req -new -nodes -utf8 -sha256 -x509 -outform PEM -out kernel_key.pem -keyout kernel_key.pem`

Az OpenSSL néhány kérdést tesz fel a kulcsot generáló felhasználóról, ajánlott ezeket a kérdéseket a lehető legrészletesebben kitölteni.

Tárolja a kulcsot biztonságos helyen, legalábbis a kulcsot csak a root felhasználónak kell tudnia olvasni. Ellenőrizze ezt a következővel:

`root #` `ls -l kernel_key.pem`

```
 -r-------- 1 root root 3164 Jan  4 10:38 kernel_key.pem
```

Ha a kimenet bármi mást mutat, mint a fentiek, akkor javítsa ki a jogosultságokat a következő parancsokkal:

`root #` `chown root:root kernel_key.pem
`

`root #` `chmod 400 kernel_key.pem
`

KERNEL **Az aláíró kulcs megadása CONFIG\_MODULE\_SIG\_KEY**

```
-*- Cryptographic API  --->
  Certificates for signature checking  --->
    (/path/to/kernel_key.pem) File name or PKCS#11 URI of module signing key

```

Más szoftvercsomagok által telepített külső kernelmodulok aláírásához engedélyezze a [modules-sign](https://packages.gentoo.org/useflags/modules-sign) [USE flag/hu](/wiki/USE_flag/hu "USE flag/hu") USE jelölőzászlót globálisan:

FILE **`/etc/portage/make.conf`** **Modul aláírásának az engedélyezése**

```
USE="modules-sign"

# Opcionálisan, ha egyedi aláíró kulcsokat használ.
MODULES_SIGN_KEY="/path/to/kernel_key.pem"
MODULES_SIGN_CERT="/path/to/kernel_key.pem" # Kizárólag akkor szükséges, ha a MODULES_SIGN_KEY a tanúsítványt se nem tartalmazza.
MODULES_SIGN_HASH="sha512" # Alapértelmezés szerint sha512.

```

**Note**

A MODULES\_SIGN\_KEY és a MODULES\_SIGN\_CERT különböző fájlokra mutathat. Ebben a példában az OpenSSL által generált pem fájl tartalmazza mind a kulcsot, mind a hozzá tartozó tanúsítványt, így mindkét változót ugyanarra az értékre állítjuk be.

### Kódfordítás és telepítés

Miután a kernel beállítása befejeződött, itt az ideje annak fordítására és telepítésére. Lépjen ki a beállításból, és indítsa el a kódfordítási folyamatot:

`root #` `make && make modules_install`

**Note**

Lehetséges engedélyezni a párhuzamos buildelést a make -j N használatával, ahol N annak a párhuzamos feladatoknak az egész számú _száma_, amelyeket a build folyamat elindíthat. Ez hasonló a korábbi /etc/portage/make.conf fájlban található utasításokhoz, a MAKEOPTS változóval.

Amikor a kernel fordítása befejeződött, ellenőrizze a keletkező fájl méretét:

`root #` `ls -lh arch/sparc/boot/image`

```
-rw-r--r--    1 root     root         2.4M Oct 25 14:38 image

```

Ha a (tömörítetlen) méret nagyobb, mint 7,5 MB, akkor állítsa be újra a kernelt, amíg az nem haladja meg ezeket a határértékeket. Ennek egyik módja az, ha a legtöbb kernel-illesztőprogramot modulokként fordítják le. Ennek figyelmen kívül hagyása egy nem indítható kernelhez vezethet.

Ha a kernel csak egy kicsit túl nagy, akkor próbálja meg csökkenteni a méretét a strip parancs használatával:

`root #` `strip -R .comment -R .note arch/sparc/boot/image`

Végül másolja át a binári kernelképfájlt a /boot/ könyvtárba.

`root #` `cp arch/sparc/boot/image /boot/kernel-6.19.3-gentoo `

Folytassa a telepítést a [Rendszer beállítása](/wiki/Handbook:SPARC/Installation/System/hu "Handbook:SPARC/Installation/System/hu") fejezettel.

[← Alaprendszer telepítése](/wiki/Handbook:SPARC/Installation/Base/hu "Previous part") [Kezdőlap](/wiki/Handbook:SPARC/hu "Handbook:SPARC/hu") [Rendszer beállítása →](/wiki/Handbook:SPARC/Installation/System/hu "Next part")