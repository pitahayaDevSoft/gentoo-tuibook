# Tools

Other languages:

- [Deutsch](/wiki/Handbook:PPC64/Installation/Tools/de "Handbuch:PPC64/Installation/Tools (100% translated)")
- [English](/wiki/Handbook:PPC64/Installation/Tools "Handbook:PPC64/Installation/Tools (100% translated)")
- [español](/wiki/Handbook:PPC64/Installation/Tools/es "Manual de Gentoo: PPC64/Instalación/Herramientas (100% translated)")
- [français](/wiki/Handbook:PPC64/Installation/Tools/fr "Handbook:PPC64/Installation/Tools/fr (100% translated)")
- [italiano](/wiki/Handbook:PPC64/Installation/Tools/it "Handbook:PPC64/Installation/Tools/it (100% translated)")
- magyar
- [polski](/wiki/Handbook:PPC64/Installation/Tools/pl "Handbook:PPC64/Installation/Tools (100% translated)")
- [português do Brasil](/wiki/Handbook:PPC64/Installation/Tools/pt-br "Handbook:PPC64/Installation/Tools/pt-br (50% translated)")
- [русский](/wiki/Handbook:PPC64/Installation/Tools/ru "Handbook:PPC64/Installation/Tools (100% translated)")
- [தமிழ்](/wiki/Handbook:PPC64/Installation/Tools/ta "கையேடு:PPC64/நிறுவல்/கருவிகள் (100% translated)")
- [中文（中国大陆）‎](/wiki/Handbook:PPC64/Installation/Tools/zh-cn "手册：PPC64/安装/安装系统工具 (100% translated)")
- [日本語](/wiki/Handbook:PPC64/Installation/Tools/ja "ハンドブック:PPC64/インストール/ツール (100% translated)")
- [한국어](/wiki/Handbook:PPC64/Installation/Tools/ko "Handbook:PPC64/Installation/Tools/ko (100% translated)")

[PPC64 kézikönyv](/wiki/Handbook:PPC64/hu "Handbook:PPC64/hu")[A Gentoo Linux telepítése](/wiki/Handbook:PPC64/Full/Installation/hu "Handbook:PPC64/Full/Installation/hu")[A telepítésről](/wiki/Handbook:PPC64/Installation/About/hu "Handbook:PPC64/Installation/About/hu")[Telepítőképfájl kiválasztása](/wiki/Handbook:PPC64/Installation/Media/hu "Handbook:PPC64/Installation/Media/hu")[Hálózat beállítása](/wiki/Handbook:PPC64/Installation/Networking/hu "Handbook:PPC64/Installation/Networking/hu")[Adathordozók előkészítése](/wiki/Handbook:PPC64/Installation/Disks/hu "Handbook:PPC64/Installation/Disks/hu")[Stage fájl](/wiki/Handbook:PPC64/Installation/Stage/hu "Handbook:PPC64/Installation/Stage/hu")[Alaprendszer telepítése](/wiki/Handbook:PPC64/Installation/Base/hu "Handbook:PPC64/Installation/Base/hu")[Kernel beállítása](/wiki/Handbook:PPC64/Installation/Kernel/hu "Handbook:PPC64/Installation/Kernel/hu")[Rendszer beállítása](/wiki/Handbook:PPC64/Installation/System/hu "Handbook:PPC64/Installation/System/hu")Eszközök telepítése[Bootloader beállítása](/wiki/Handbook:PPC64/Installation/Bootloader/hu "Handbook:PPC64/Installation/Bootloader/hu")[Telepítés véglegesítése](/wiki/Handbook:PPC64/Installation/Finalizing/hu "Handbook:PPC64/Installation/Finalizing/hu")[Munka a Gentoo rendszerrel](/wiki/Handbook:PPC64/Full/Working/hu "Handbook:PPC64/Full/Working/hu")[Portage bemutatása](/wiki/Handbook:PPC64/Working/Portage/hu "Handbook:PPC64/Working/Portage/hu")[USE jelölőzászlók](/wiki/Handbook:PPC64/Working/USE/hu "Handbook:PPC64/Working/USE/hu")[Portage jellemzői](/wiki/Handbook:PPC64/Working/Features/hu "Handbook:PPC64/Working/Features/hu")[Init-szkript rendszer](/wiki/Handbook:PPC64/Working/Initscripts/hu "Handbook:PPC64/Working/Initscripts/hu")[Környezeti változók](/wiki/Handbook:PPC64/Working/EnvVar/hu "Handbook:PPC64/Working/EnvVar/hu")[Munka a Portage szoftvercsomag-kezelővel](/wiki/Handbook:PPC64/Full/Portage/hu "Handbook:PPC64/Full/Portage/hu")[Fájlok és könyvtárak](/wiki/Handbook:PPC64/Portage/Files/hu "Handbook:PPC64/Portage/Files/hu")[Változók](/wiki/Handbook:PPC64/Portage/Variables/hu "Handbook:PPC64/Portage/Variables/hu")[Szoftverágak keverése](/wiki/Handbook:PPC64/Portage/Branches/hu "Handbook:PPC64/Portage/Branches/hu")[További eszközök](/wiki/Handbook:PPC64/Portage/Tools/hu "Handbook:PPC64/Portage/Tools/hu")[Egyéni szoftvercsomag-tárolók](/wiki/Handbook:PPC64/Portage/CustomTree/hu "Handbook:PPC64/Portage/CustomTree/hu")[Fejlett funkciók](/wiki/Handbook:PPC64/Portage/Advanced/hu "Handbook:PPC64/Portage/Advanced/hu")[Hálózat beállítása OpenRC init-rendszeren](/wiki/Handbook:PPC64/Full/Networking/hu "Handbook:PPC64/Full/Networking/hu")[Munka elkezdése](/wiki/Handbook:PPC64/Networking/Introduction/hu "Handbook:PPC64/Networking/Introduction/hu")[Fejlett beállítások](/wiki/Handbook:PPC64/Networking/Advanced/hu "Handbook:PPC64/Networking/Advanced/hu")[Moduláris hálózat](/wiki/Handbook:PPC64/Networking/Modular/hu "Handbook:PPC64/Networking/Modular/hu")[Vezeték nélküli (Wi-Fi)](/wiki/Handbook:PPC64/Networking/Wireless/hu "Handbook:PPC64/Networking/Wireless/hu")[Funkcionalitás hozzáadása](/wiki/Handbook:PPC64/Networking/Extending/hu "Handbook:PPC64/Networking/Extending/hu")[Dinamikus menedzsment](/wiki/Handbook:PPC64/Networking/Dynamic/hu "Handbook:PPC64/Networking/Dynamic/hu")

## Contents

- [1Rendszernaplózó](#Rendszernapl.C3.B3z.C3.B3)
  - [1.1OpenRC](#OpenRC)
  - [1.2systemd](#systemd)
- [2Opcionális: A cron szolgáltatás](#Opcion.C3.A1lis:_A_cron_szolg.C3.A1ltat.C3.A1s)
  - [2.1OpenRC](#OpenRC_2)
    - [2.1.1Alapértelmezett: cronie](#Alap.C3.A9rtelmezett:_cronie)
    - [2.1.2Alternatíva: dcron](#Alternat.C3.ADva:_dcron)
    - [2.1.3Alternatíva: fcron](#Alternat.C3.ADva:_fcron)
    - [2.1.4Alternatíva: bcron](#Alternat.C3.ADva:_bcron)
    - [2.1.5modprobed-db users](#modprobed-db_users)
  - [2.2systemd](#systemd_2)
    - [2.2.1modprobed-db users](#modprobed-db_users_2)
- [3Opcionális: Fájlindexelés](#Opcion.C3.A1lis:_F.C3.A1jlindexel.C3.A9s)
- [4Opcionális: Távoli parancssor hozzáférés](#Opcion.C3.A1lis:_T.C3.A1voli_parancssor_hozz.C3.A1f.C3.A9r.C3.A9s)
  - [4.1OpenRC](#OpenRC_3)
  - [4.2systemd](#systemd_3)
- [5Opcionális: A Parancssor parancsainak az automatikus kiegészítése](#Opcion.C3.A1lis:_A_Parancssor_parancsainak_az_automatikus_kieg.C3.A9sz.C3.ADt.C3.A9se)
  - [5.1Bash](#Bash)
- [6Javasolt: Idő szinkronizálása](#Javasolt:_Id.C5.91_szinkroniz.C3.A1l.C3.A1sa)
  - [6.1chrony](#chrony)
  - [6.2OpenRC](#OpenRC_4)
  - [6.3systemd](#systemd_4)
  - [6.4systemd-timesyncd](#systemd-timesyncd)
- [7Fájlrendszereszközök](#F.C3.A1jlrendszereszk.C3.B6z.C3.B6k)

## Rendszernaplózó

### OpenRC

Néhány eszköz hiányzik a stage3 archívumból, mert több szoftvercsomag ugyanazt a funkciót biztosítja. Most már a felhasználón múlik, hogy melyiket telepíti.

Az első eszközválasztás a rendszer naplózási mechanizmusa. A Unix és Linux operációs rendszerek kiváló naplózási képességekkel rendelkeznek – szükség esetén minden, ami a operációs rendszeren történik, egy naplófájlba rögzíthető.

A Gentoo számos rendszernaplózási segédprogramot kínál. Néhány példa ezek közül:

- [sysklogd](/wiki/Sysklogd "Sysklogd") ([app-admin/sysklogd](https://packages.gentoo.org/packages/app-admin/sysklogd)) \- A hagyományos rendszernaplózó szolgáltatásokat kínálja. Az alapértelmezett naplózási beállítás már különösebb beállítás nélkül is jól működik, ami ezt a szoftvercsomagot jó választássá teszi kezdők számára.
- [syslog-ng](/wiki/Syslog-ng "Syslog-ng") ([app-admin/syslog-ng](https://packages.gentoo.org/packages/app-admin/syslog-ng)) \- Fejlett rendszernaplózó szoftver. Minden, ami több annál, mint hogy egy nagy fájlba naplózunk, az további beállítást igényel. Haladó felhasználók választhatják ezt a szoftvercsomagot naplózási potenciálja alapján, de vegye figyelembe, hogy bármilyen okos naplózáshoz további beállításra van szükség.
- [metalog](/wiki/Metalog "Metalog") ([app-admin/metalog](https://packages.gentoo.org/packages/app-admin/metalog)) \- Egy magas beállíthatósággal rendszer naplózó.

A Gentoo ebuild szoftvertárolóján keresztül más rendszernaplózó segédprogramok is elérhetők lehetnek, mivel az elérhető szoftvercsomagok száma naponta növekszik.

**Tip**

Ha a syslog-ng használata van tervben, akkor ajánlott telepíteni és beállítani a [logrotate](/wiki/Logrotate "Logrotate") szoftvert. A syslog-ng szoftver nem biztosít semmilyen forgatási mechanizmust a naplófájlok számára. Azonban a sysklogd újabb verziói (>= 2.0) saját naplóforgatási mechanizmust kezelnek.

A választott rendszernaplózó telepítéséhez használja az emerge parancsot. Ha Ön OpenRC init-rendszert használ, akkor hozzáadhatja az alapértelmezett futási szinthez a rc-update használatával. Az alábbi példa telepíti és aktiválja az [app-admin/sysklogd](https://packages.gentoo.org/packages/app-admin/sysklogd) szoftvercsomagot mint az operációs rendszer syslog segédprogramját:

`root #` `emerge --ask app-admin/sysklogd`

`root #` `rc-update add sysklogd default`

### systemd

Bár az OpenRC-alapú operációs rendszerekhez különféle naplózási mechanizmusok állnak rendelkezésre, a systemd beépített naplózót tartalmaz, amelyet **systemd-journald** szolgáltatásnak neveznek. A systemd-journald szolgáltatás képes kezelni a korábbi rendszernaplózási szakaszban ismertetett naplózási funkciók nagy részét. Ez azt jelenti, hogy azok az operációs rendszerek, amelyek a systemd init-rendszert használják a rendszer- és szolgáltatáskezelőként, azok _biztonságosan kihagyhatják_ a további syslog segédprogramok hozzáadását.

További részletekért a journalctl használatáról a rendszernaplók lekérdezéséhez és áttekintéséhez, tekintse meg a man journalctl súgót.

Számos okból kifolyólag, például a naplók továbbításának esetében egy központi gazdagépre, fontos lehet _redundáns_ rendszernaplózási mechanizmusokat beépíteni a systemd alapú operációs rendszerbe. Ez azonban ritkán fordul elő a kézikönyv tipikus olvasói számára, és haladószintű felhasználási esetnek tekinthető. Ezért a kézikönyv nem tárgyalja.

## Opcionális: A cron szolgáltatás

### OpenRC

Bár opcionális és nem minden operációs rendszerhez szükséges, bölcs dolog telepíteni egy cron szolgáltatást.

A cron szolgáltatás meghatározott időközönként hajt végre parancsokat. Az intervallumok lehetnek napi, heti vagy havi, minden kedden egyszer, minden második héten egyszer, stb. Egy bölcs rendszergazda a cron szolgáltatást használja a rutinszerű rendszerkarbantartási feladatok automatizálására.

Minden cron szolgáltatás támogatja az időzített feladatok magas szintű részletességét, és általában tartalmazza az e-mail vagy más értesítési forma küldésének lehetőségét, ha egy időzített feladat nem fejeződik be a vártnak megfelelően.

A Gentoo számos lehetséges cron szolgáltatást kínál, többek között:

- [cronie](/wiki/Cron/hu#cronie "Cron/hu") ([sys-process/cronie](https://packages.gentoo.org/packages/sys-process/cronie)) \- A cronie az eredeti cron szoftverre épül, és biztonsági és beállítási fejlesztéseket tartalmaz, mint például a PAM és az SELinux használatának lehetősége.
- [dcron](/wiki/Cron/hu#dcron "Cron/hu") ([sys-process/dcron](https://packages.gentoo.org/packages/sys-process/dcron)) \- Ez a könnyűsúlyú cron szolgáltatás egyszerű és biztonságos. Még épp elég funkcióval rendelkezik ahhoz, hogy éppen hasznos maradjon.
- [fcron](/wiki/Cron/hu#fcron "Cron/hu") ([sys-process/fcron](https://packages.gentoo.org/packages/sys-process/fcron)) \- Egy parancsütemező, amely kiterjesztett képességekkel rendelkezik a cron és az anacron felett.
- [bcron](/wiki/Cron/hu#bcron "Cron/hu") \- Egy fiatalabb cron rendszer, amelyet biztonságos műveletek figyelembevételével terveztek. Ennek érdekében a cron-rendszert több különálló programra osztották fel. Mindegyik külön feladatért felelős, és szigorúan ellenőrzött kommunikáció zajlik az egyes részek között.

#### Alapértelmezett: cronie

Az alábbi példa a [sys-process/cronie](https://packages.gentoo.org/packages/sys-process/cronie) szoftvercsomagot használja:

`root #` `emerge --ask sys-process/cronie`

Adja hozzá a cronie szoftvert az operációs rendszer alapértelmezett futási szintjéhez, amely automatikusan elindítja azt a számítógép bekapcsolásának az alkalmával:

`root #` `rc-update add cronie default`

#### Alternatíva: dcron

`root #` `emerge --ask sys-process/dcron`

Ha a dcron lesz a kiválasztott cron ügynök, akkor egy további inicializáló parancsot kell végrehajtani:

`root #` `crontab /etc/crontab`

#### Alternatíva: fcron

`root #` `emerge --ask sys-process/fcron`

Ha az fcron lesz a kiválasztott időzített feladatkezelő, akkor további emerge lépés szükséges:

`root #` `emerge --config sys-process/fcron`

#### Alternatíva: bcron

A bcron egy fiatalabb cron ügynök, beépített jogosultság-elválasztással.

`root #` `emerge --ask sys-process/bcron`

#### modprobed-db users

If using [sys-kernel/modprobed-db](https://packages.gentoo.org/packages/sys-kernel/modprobed-db) was picked as an option for manually compiling the kernel.

Now is the time to setup a crontab to periodically scan the system for hardware used.

FILE **`/etc/crontab`** **Run modprobed-db once every 6 hours**

```
0 */6 * * *     /usr/bin/modprobed-db store &> /dev/null

```

At a later date of at least a week, please visit the kernel build section of the [modprobed-db](/wiki/Modprobed-db#Building_kernels.2Fhu "Modprobed-db") article to complete the setup.

### systemd

Hasonlóan a rendszer naplózáshoz, a systemd-alapú operációs rendszerek alapból támogatják az ütemezett feladatokat _timers_ (időzítők) formájában. A systemd időzítők futtathatók operációs rendszer szinten vagy felhasználói szinten, és ugyanazokat a funkciókat kínálják, mint egy hagyományos cron szolgáltatás. Hacsak nincs szükség redundáns képességekre, akkor általában felesleges további feladatütemezőt, például cron szolgáltatást telepíteni, és biztonságosan kihagyható.

#### modprobed-db users

If using [sys-kernel/modprobed-db](https://packages.gentoo.org/packages/sys-kernel/modprobed-db) was picked as an option for manually compiling the kernel.

Now is the time to setup a systemd timer to periodically scan the system for hardware used.

`root #` `systemctl --user enable modprobed-db`

At a later date of at least a week, please visit the kernel build section of the [modprobed-db](/wiki/Modprobed-db#Building_kernels.2Fhu "Modprobed-db") article to complete the setup.

## Opcionális: Fájlindexelés

A fájlrendszer indexelése érdekében, hogy gyorsabb fájlkeresési lehetőségeket biztosítson, telepítse a [mlocate](/wiki/Mlocate "Mlocate") szoftvercsomagot:

`root #` `emerge --ask sys-apps/mlocate`

## Opcionális: Távoli parancssor hozzáférés

**Tip**

Az opensshd alapértelmezett beállítás nem engedélyezi a root felhasználó számára a távoli bejelentkezést. Kérjük, hogy [hozzon létre egy nem root felhasználót](/wiki/FAQ/hu#How_do_I_add_a_normal_user.3F "FAQ/hu") és állítsa be megfelelően, hogy szükség esetén hozzáférést biztosítson a telepítés után, vagy módosítsa a /etc/ssh/sshd\_config fájlt, hogy engedélyezze a root hozzáférést.

Az operációs rendszer távoli elérésének biztosítása érdekében a telepítés után a sshd szolgáltatást úgy kell beállítani, hogy az a rendszerindításkor elinduljon.

Részletesebb információkért az SSH beállításával kapcsolatban, kérjük, olvassa el a [SSH](/wiki/SSH/hu "SSH/hu") cikket.

### OpenRC

Az sshd init szkript hozzáadása az alapértelmezett futási szinthez OpenRC init-rendszer esetében:

`root #` `rc-update add sshd default`

`root #` `rc-update add sshd default`

Ha szükség van a soros konzol hozzáférésre (ami távoli szerverek esetén lehetséges), akkor a agetty szolgáltatást be kell állítani.

Vegye ki a kommenteket a soros konzol szakasznál a /etc/inittab fájlban:

`root #` `nano -w /etc/inittab`

```
# SERIAL CONSOLES
s0:12345:respawn:/sbin/agetty 9600 ttyS0 vt100
s1:12345:respawn:/sbin/agetty 9600 ttyS1 vt100

```

### systemd

Az SSH szerver engedélyezéséhez futtassa a következő parancsot:

`root #` `systemctl enable sshd`

`root #` `systemctl enable sshd`

A soros konzol támogatásának engedélyezéséhez futtassa:

`root #` `systemctl enable getty@tty1.service`

`root #` `systemctl enable getty@tty1.service`

## Opcionális: A Parancssor parancsainak az automatikus kiegészítése

### Bash

A Bash az alapértelmezett parancssor a Gentoo rendszerekhez, ezért a kiegészítő bővítmények telepítése segíthet az operációs rendszer hatékonyabb és kényelmesebb kezelésében. Az [app-shells/bash-completion](https://packages.gentoo.org/packages/app-shells/bash-completion) szoftvercsomag telepíti a Gentoo specifikus parancsokhoz elérhető kiegészítéseket, valamint számos más általános parancsot és segédprogramot is telepít:

`root #` `emerge --ask app-shells/bash-completion`

A telepítés után a specifikus parancsokhoz tartozó bash kiegészítések az eselect segítségével kezelhetők. További részletekért tekintse meg a bash cikk [Shell completion integrációk szakaszát](/wiki/Bash#Shell_completion_integrations.2Fhu "Bash").

## Javasolt: Idő szinkronizálása

**Important**

Systems without a functioning [Real-Time Clock (RTC)](/wiki/System_time/hu#Software_clock_vs_Hardware_clock "System time/hu") must set the [system time](/wiki/System_time/hu "System time/hu") at every system start, and on regular intervals thereafter.

Fontos a rendszeróra szinkronizálása valamilyen módszerrel. Ezt általában az [NTP](/wiki/NTP "NTP") protokoll és szoftverrel végzik el. Léteznek más megvalósítások is az NTP protokoll használatával, mint például a [Chrony](/wiki/Chrony "Chrony").

The clock is usually synchronized via the [Network Time Protocol](/wiki/Network_Time_Protocol "Network Time Protocol"), with an implementation such as [chrony](/wiki/Chrony "Chrony").

### chrony

A Chrony beállításához, például:

`root #` `emerge --ask net-misc/chrony`

`root #` `emerge --ask net-misc/chrony`

See the [chrony](/wiki/Chrony "Chrony") article for further information, for example if more advanced configurations are required.

### OpenRC

Az OpenRC init rendszerrel működő operációs rendszeren futtassa:

`root #` `rc-update add chronyd default`

`root #` `rc-update add chronyd default`

### systemd

A systemd init rendszerrel működő operációs rendszeren futtassa:

`root #` `systemctl enable chronyd.service`

`root #` `systemctl enable chronyd.service`

### systemd-timesyncd

Alternatív megoldásként a systemd felhasználók esetleg a egyszerűbb systemd-timesyncd SNTP klienst szeretnék használni, amely alapértelmezés szerint telepítve van.

`root #` `systemctl enable systemd-timesyncd.service`

`root #` `systemctl enable systemd-timesyncd.service`

## Fájlrendszereszközök

A használt fájlrendszertől függően szükség lehet a szükséges fájlrendszer-segédprogramok telepítésére (a fájlrendszer-integritás ellenőrzése, fájlrendszerek (újra)formázása stb.). Vegye figyelembe, hogy az ext4 felhasználói tér eszközei ([sys-fs/e2fsprogs](https://packages.gentoo.org/packages/sys-fs/e2fsprogs)) már telepítve vannak a [@system készlet](/wiki/System_set_(Portage) "System set (Portage)") részeként.

A következő táblázat felsorolja azokat az eszközöket, amelyeket telepíteni kell, ha bizonyos fájlrendszereszközökre van szükség a telepített operációs rendszeren.

Fájlrendszer
Szoftvercsomag
[XFS](/wiki/XFS/hu "XFS/hu")[sys-fs/xfsprogs](https://packages.gentoo.org/packages/sys-fs/xfsprogs)[ext4](/wiki/Ext4/hu "Ext4/hu")[sys-fs/e2fsprogs](https://packages.gentoo.org/packages/sys-fs/e2fsprogs)[VFAT](/wiki/VFAT "VFAT") (FAT32, ...)
[sys-fs/dosfstools](https://packages.gentoo.org/packages/sys-fs/dosfstools)[Btrfs](/wiki/Btrfs/hu "Btrfs/hu")[sys-fs/btrfs-progs](https://packages.gentoo.org/packages/sys-fs/btrfs-progs)[F2FS](/wiki/F2FS/hu "F2FS/hu")[sys-fs/f2fs-tools](https://packages.gentoo.org/packages/sys-fs/f2fs-tools)[NTFS](/wiki/NTFS/hu "NTFS/hu")[sys-fs/ntfs3g](https://packages.gentoo.org/packages/sys-fs/ntfs3g)[ZFS](/wiki/ZFS "ZFS")[sys-fs/zfs](https://packages.gentoo.org/packages/sys-fs/zfs)[bcachefs](/wiki/Bcachefs "Bcachefs")[sys-fs/bcachefs-tools](https://packages.gentoo.org/packages/sys-fs/bcachefs-tools)

Ajánlott a [sys-block/io-scheduler-udev-rules](https://packages.gentoo.org/packages/sys-block/io-scheduler-udev-rules) szoftvercsomag telepítése a megfelelő ütemező viselkedés érdekében, például az nvme eszközök esetében:

`root #` `emerge --ask sys-block/io-scheduler-udev-rules`

`root #` `emerge --ask sys-block/io-scheduler-udev-rules`

**Tip**

További információkért a fájlrendszerekről a Gentoo rendszerben tekintse meg a [fájlrendszer cikket](/wiki/Filesystem/hu "Filesystem/hu").

Most folytassa a [Bootloader beállítása](/wiki/Handbook:PPC64/Installation/Bootloader/hu "Handbook:PPC64/Installation/Bootloader/hu") részével.

[← Rendszer beállítása](/wiki/Handbook:PPC64/Installation/System/hu "Previous part") [Kezdőlap](/wiki/Handbook:PPC64/hu "Handbook:PPC64/hu") [Bootloader beállítása →](/wiki/Handbook:PPC64/Installation/Bootloader/hu "Next part")