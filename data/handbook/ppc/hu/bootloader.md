# Bootloader

Other languages:

- [Deutsch](/wiki/Handbook:PPC/Installation/Bootloader/de "Handbuch:PPC/Installation/Bootloader (100% translated)")
- [English](/wiki/Handbook:PPC/Installation/Bootloader "Handbook:PPC/Installation/Bootloader (100% translated)")
- [español](/wiki/Handbook:PPC/Installation/Bootloader/es "Manual de Gentoo: PPC/Instalación/Arranque (100% translated)")
- [français](/wiki/Handbook:PPC/Installation/Bootloader/fr "Handbook:PPC/Installation/Bootloader/fr (100% translated)")
- [italiano](/wiki/Handbook:PPC/Installation/Bootloader/it "Handbook:PPC/Installation/Bootloader/it (100% translated)")
- magyar
- [polski](/wiki/Handbook:PPC/Installation/Bootloader/pl "Handbook:PPC/Installation/Bootloader (100% translated)")
- [português do Brasil](/wiki/Handbook:PPC/Installation/Bootloader/pt-br "Handbook:PPC/Installation/Bootloader/pt-br (50% translated)")
- [русский](/wiki/Handbook:PPC/Installation/Bootloader/ru "Handbook:PPC/Installation/Bootloader (100% translated)")
- [தமிழ்](/wiki/Handbook:PPC/Installation/Bootloader/ta "கையேடு:PPC/நிறுவல்/துவக்கஏற்றி (100% translated)")
- [中文（中国大陆）‎](/wiki/Handbook:PPC/Installation/Bootloader/zh-cn "手册：PPC/安装/配置系统引导程序Bootloader (100% translated)")
- [日本語](/wiki/Handbook:PPC/Installation/Bootloader/ja "ハンドブック:PPC/インストール/ブートローダー (100% translated)")
- [한국어](/wiki/Handbook:PPC/Installation/Bootloader/ko "Handbook:PPC/Installation/Bootloader/ko (100% translated)")

[PPC kézikönyv](/wiki/Handbook:PPC/hu "Handbook:PPC/hu")[A Gentoo Linux telepítése](/wiki/Handbook:PPC/Full/Installation/hu "Handbook:PPC/Full/Installation/hu")[A telepítésről](/wiki/Handbook:PPC/Installation/About/hu "Handbook:PPC/Installation/About/hu")[Telepítőképfájl kiválasztása](/wiki/Handbook:PPC/Installation/Media/hu "Handbook:PPC/Installation/Media/hu")[Hálózat beállítása](/wiki/Handbook:PPC/Installation/Networking/hu "Handbook:PPC/Installation/Networking/hu")[Adathordozók előkészítése](/wiki/Handbook:PPC/Installation/Disks/hu "Handbook:PPC/Installation/Disks/hu")[Stage fájl](/wiki/Handbook:PPC/Installation/Stage/hu "Handbook:PPC/Installation/Stage/hu")[Alaprendszer telepítése](/wiki/Handbook:PPC/Installation/Base/hu "Handbook:PPC/Installation/Base/hu")[Kernel beállítása](/wiki/Handbook:PPC/Installation/Kernel/hu "Handbook:PPC/Installation/Kernel/hu")[Rendszer beállítása](/wiki/Handbook:PPC/Installation/System/hu "Handbook:PPC/Installation/System/hu")[Eszközök telepítése](/wiki/Handbook:PPC/Installation/Tools/hu "Handbook:PPC/Installation/Tools/hu")Bootloader beállítása[Telepítés véglegesítése](/wiki/Handbook:PPC/Installation/Finalizing/hu "Handbook:PPC/Installation/Finalizing/hu")[Munka a Gentoo rendszerrel](/wiki/Handbook:PPC/Full/Working/hu "Handbook:PPC/Full/Working/hu")[Portage bemutatása](/wiki/Handbook:PPC/Working/Portage/hu "Handbook:PPC/Working/Portage/hu")[USE jelölőzászlók](/wiki/Handbook:PPC/Working/USE/hu "Handbook:PPC/Working/USE/hu")[Portage jellemzői](/wiki/Handbook:PPC/Working/Features/hu "Handbook:PPC/Working/Features/hu")[Init-szkript rendszer](/wiki/Handbook:PPC/Working/Initscripts/hu "Handbook:PPC/Working/Initscripts/hu")[Környezeti változók](/wiki/Handbook:PPC/Working/EnvVar/hu "Handbook:PPC/Working/EnvVar/hu")[Munka a Portage szoftvercsomag-kezelővel](/wiki/Handbook:PPC/Full/Portage/hu "Handbook:PPC/Full/Portage/hu")[Fájlok és könyvtárak](/wiki/Handbook:PPC/Portage/Files/hu "Handbook:PPC/Portage/Files/hu")[Változók](/wiki/Handbook:PPC/Portage/Variables/hu "Handbook:PPC/Portage/Variables/hu")[Szoftverágak keverése](/wiki/Handbook:PPC/Portage/Branches/hu "Handbook:PPC/Portage/Branches/hu")[További eszközök](/wiki/Handbook:PPC/Portage/Tools/hu "Handbook:PPC/Portage/Tools/hu")[Egyéni szoftvercsomag-tárolók](/wiki/Handbook:PPC/Portage/CustomTree/hu "Handbook:PPC/Portage/CustomTree/hu")[Fejlett funkciók](/wiki/Handbook:PPC/Portage/Advanced/hu "Handbook:PPC/Portage/Advanced/hu")[Hálózat beállítása OpenRC init-rendszeren](/wiki/Handbook:PPC/Full/Networking/hu "Handbook:PPC/Full/Networking/hu")[Munka elkezdése](/wiki/Handbook:PPC/Networking/Introduction/hu "Handbook:PPC/Networking/Introduction/hu")[Fejlett beállítások](/wiki/Handbook:PPC/Networking/Advanced/hu "Handbook:PPC/Networking/Advanced/hu")[Moduláris hálózat](/wiki/Handbook:PPC/Networking/Modular/hu "Handbook:PPC/Networking/Modular/hu")[Vezeték nélküli (Wi-Fi)](/wiki/Handbook:PPC/Networking/Wireless/hu "Handbook:PPC/Networking/Wireless/hu")[Funkcionalitás hozzáadása](/wiki/Handbook:PPC/Networking/Extending/hu "Handbook:PPC/Networking/Extending/hu")[Dinamikus menedzsment](/wiki/Handbook:PPC/Networking/Dynamic/hu "Handbook:PPC/Networking/Dynamic/hu")

## Contents

- [1Bootloader opciók](#Bootloader_opci.C3.B3k)
- [2NewWorld Macs](#NewWorld_Macs)
  - [2.1GRUB](#GRUB)
  - [2.2Telepítés](#Telep.C3.ADt.C3.A9s)
  - [2.3Bootstrap partíció beállítása](#Bootstrap_part.C3.ADci.C3.B3_be.C3.A1ll.C3.ADt.C3.A1sa)
  - [2.4GRUB beállítása](#GRUB_be.C3.A1ll.C3.ADt.C3.A1sa)
- [3OldWorld Macs](#OldWorld_Macs)
  - [3.1BootX](#BootX)
- [4Pegasos](#Pegasos)
  - [4.1BootCreator](#BootCreator)
- [5Rendszer újraindítása](#Rendszer_.C3.BAjraind.C3.ADt.C3.A1sa)

## Bootloader opciók

Most már, hogy a kernel beállítása és forráskód fordítása befejeződött, valamint a szükséges rendszerbeállítás-fájlok megfelelően kitöltésre kerültek, ideje telepíteni egy olyan programot, amely a bináris futtatható kernelképfájl indítását végzi a rendszer elindításakor. Egy ilyen programot boot loader néven neveznek.

Az alkalmazandó boot loader típusa a PPC számítógép típusától függ.

NewWorld Apple vagy IBM számítógépek esetén a grub boot loadert kell kiválasztani. OldWorld Apple számítógépek esetében csak egy opció áll rendelkezésre: A BootX. A Pegasos nem igényel boot loadert, de szükséges a bootcreator telepítése a SmartFirmware boot menük létrehozásához.

## NewWorld Macs

### GRUB

### Telepítés

`root #` `emerge --ask sys-boot/grub`

### Bootstrap partíció beállítása

Először készítse elő a bootstrap partíciót, amelyet a [adathordozó előkészítése](/wiki/Handbook:PPC64/Installation/Disks "Handbook:PPC64/Installation/Disks") lépés során hozott létre. A példában szereplő /dev/sda2 partíció legyen. Opcionálisan ellenőrizze ezt a parted használatával:

Cserélje le a /dev/sda útvonalat a megfelelő eszközre, ha szükséges.

`root #` `parted /dev/sda print`

```
Model: ATA Patriot Burst El (scsi)

Disk /dev/sda: 120GB

Sector size (logical/physical): 512B/512B

Partition Table: mac

Disk Flags:

Number  Start   End     Size    File system  Name       Flags
 1      512B    32.8kB  32.3kB               Apple
 2      32.8kB  852kB   819kB   hfs          bootstrap  boot
 3      852kB   538MB   537MB   ext4         Boot
 4      538MB   54.2GB  53.7GB  ext4         Gentoo

```

Ebben a kimenetben a 2. partíció tartalmazza a bootstrap információt, így a /dev/sda2 a helyes útvonal, amit használni kell.

Formázza ezt a partíciót [HFS](/wiki/HFS "HFS") formátumra a hformat parancs segítségével, amely a [sys-fs/hfsutils](https://packages.gentoo.org/packages/sys-fs/hfsutils) szoftvercsomag része.

`root #` `dd if=/dev/zero of=/dev/sda2 bs=512`

`root #` `hformat -l bootstrap /dev/sda2`

Hozzon létre egy könyvtárat a bootstrap partíció felcsatolásához, majd csatolja azt fel a fájlrendszerbe:

`root #` `mkdir /boot/NWBB
`

`root #` `mount --types hfs /dev/sda2 /boot/NWBB`

### GRUB beállítása

`root #` `grub-install --macppc-directory=/boot/NWBB /dev/sda2`

Ha a telepítés hibák nélkül fejeződik be, akkor csatolja le a bootstrap partíciót:

`root #` `umount /boot/NWBB`

Ezután tegye bootolhatóvá a partíciót, azaz blesselje, hogy elinduljon:

`root #` `hmount /dev/sda2
`

`root #` `hattrib -t tbxi -c UNIX :System:Library:CoreServices:BootX
`

`root #` `hattrib -b :System:Library:CoreServices
`

`root #` `humount`

Végül készítse el a grub.cfg fájlt:

`root #` `grub-mkconfig -o /boot/grub/grub.cfg`

## OldWorld Macs

### BootX

**Important**

A BootX kizárólag OldWorld Apple rendszereken használható MacOS klasszikus 7-től 9-ig terjedő verzióval. Azoknál a számítógépeknél, amelyek kevesebb mint 32MB RAM memóriával rendelkeznek, a MacOS klasszikus 7-et kell használni.

A BootX program letöltése elérhető a [https://github.com/immolo/BootX/archive/refs/tags/1.2.2.tar.gz](https://github.com/immolo/BootX/archive/refs/tags/1.2.2.tar.gz) linkről.

Mivel a BootX a Linuxot a MacOS-en belülről indítja, a kernelképfájlt át kell másolni a Linux partícióról a MacOS partícióra. Először csatolja a MacOS partíciót a chroot-on kívülről. Használja a mac-fdisk -l parancsot, hogy megtalálja a MacOS partíció számát. Itt példaként az sda6 szerepel. Miután a partíció csatolásra került, másolja át a kernelképfájlt a rendszer könyvtárában, hogy a BootX megtalálhassa azt.

`root #` `exit`

`cdimage ~#` `mkdir /mnt/mac
`

`cdimage ~#` `mount /dev/sda6 /mnt/mac -t hfs
`

`cdimage ~#` `cp /mnt/gentoo/usr/src/linux/vmlinux "/mnt/mac/System Folder/Linux Kernels/kernel-6.18.8-gentoo"`

Most, hogy a kernelképfájl át lett másolva, szükséges újraindítani a számítógépet a BootX beállításához.

`cdimage ~#` `cd /
`

`cdimage ~#` `umount -l /mnt/gentoo/dev{/pts,/shm,}
`

`cdimage ~#` `umount -l /mnt/gentoo{/proc,/sys,}
`

`cdimage ~#` `umount -l /mnt/mac
`

`cdimage ~#` `reboot`

Természetesen ne feledje el eltávolítani a bootolható CD lemezt, különben a CD fog újból elindulni a MacOS helyett.

Miután a számítógép betöltött a MacOS-be, nyissa meg a BootX vezérlőpanelt. Ha nem használ genkernelt, akkor válassza az Opciókat, és törölje a "Use specified RAM disk" opció jelölését. Ha genkernelt használ, akkor győződjön meg róla, hogy a genkernel initrd van kiválasztva, nem pedig a telepítő CD initrd. Ha nem használ genkernelt, akkor lehetőség van megadni a számítógép Linux root lemezét és partícióját. Töltse ki ezeket a megfelelő értékekkel. A kernel beállításától függően további boot argumentumokat is alkalmazni kellhet.

A BootX beállítható úgy, hogy Linuxot indítson bootoláskor. Ha ez megtörtént, akkor a számítógép először bebootol a MacOS-be, és a rendszerindítás során a BootX elindítja és futtatja a Linuxot. További információért látogasson el a BootX honlapjára.

**Important**

Győződjön meg arról, hogy a kernel támogatja az HFS és HFS+ fájlrendszereket, különben nem lesz lehetséges a kernel frissítése vagy módosítása a MacOS partíción.

## Pegasos

**Note**

A Pegasos is támogatja a Grub boot loadert, de ez jelenleg nincs dokumentálva a Gentoo rendszerben. Kérjük Önt, hogy adja hozzá ezt a fő wiki oldalhoz, és értesítse ezen az oldal vitafórumán, amint készen áll az ide történő áthelyezésre.

### BootCreator

**Important**

A BootCreator egy jól megtervezett SmartFirmware bootmenüt készít, amely Forth nyelven íródott a Pegasos számára.

Először győződjön meg arról, hogy a bootcreator telepítve van a rendszeren:

`root #` `emerge --ask sys-boot/bootcreator`

Most másolja át a fájlt /etc/bootmenu.example a /etc/bootmenu/ könyvtárba, majd szerkessze azt az egyéni igényeknek megfelelően:

`root #` `cp /etc/bootmenu.example /etc/bootmenu
`

`root #` `nano -w /etc/bootmenu`

Alább található a teljes /etc/bootmenu beállításfájl. A vmlinux és az initrd helyére a kernel és az initrd képfájlok neveit kell beilleszteni.

FILE **`/etc/bootmenu`** **Példa a bootcreator beállításfájlra**

```
#
# Example description file for bootcreator 1.1
#

[VERSION]
1

[TITLE]
Boot Menu

[SETTINGS]
AbortOnKey = false
Timeout    = 9
Default    = 1

[SECTION]
Local HD -> Morphos      (Normal)
ide:0 boot2.img ramdebug edebugflags="logkprintf"

[SECTION]
Local HD -> Linux (Normal)
ide:0 kernel-6.18.8-gentoo video=radeonfb:1024x768@70 root=/dev/sda3

[SECTION]
Local HD -> Genkernel (Normal)
ide:0 kernel-genkernel-ppc-6.18.8-gentoo root=/dev/ram0
root=/dev/sda3 initrd=initramfs-genkernel-ppc-6.18.8-gentoo

```

Végül a boot menüt át kell alakítani Forth formátumba, majd a boot partícióra kell másolni, hogy a SmartFirmware el tudja olvasni. Ehhez szükséges futtatni a bootcreator programot:

`root #` `bootcreator /etc/bootmenu /boot/menu`

**Note**

Ügyeljen arra, hogy nézze át a SmartFirmware beállításait újraindításkor, mert ez a menü lesz az alapértelmezés szerint betöltött fájl.

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

Miután újraindította a számítógépet, és belépett a frissen feltelepített Gentoo környezetben, bölcs dolog [véglegesíteni a Gentoo telepítést](/wiki/Handbook:PPC/Installation/Finalizing/hu "Handbook:PPC/Installation/Finalizing/hu").

[← Eszközök telepítése](/wiki/Handbook:PPC/Installation/Tools/hu "Previous part") [Kezdőlap](/wiki/Handbook:PPC/hu "Handbook:PPC/hu") [Telepítés véglegesítése →](/wiki/Handbook:PPC/Installation/Finalizing/hu "Next part")