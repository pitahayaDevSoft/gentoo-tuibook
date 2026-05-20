# Bootloader

Other languages:

- [Deutsch](/wiki/Handbook:PPC64/Installation/Bootloader/de "Handbuch:PPC64/Installation/Bootloader (100% translated)")
- [English](/wiki/Handbook:PPC64/Installation/Bootloader "Handbook:PPC64/Installation/Bootloader (100% translated)")
- [español](/wiki/Handbook:PPC64/Installation/Bootloader/es "Manual de Gentoo: PPC64/Instalación/Arranque (100% translated)")
- [français](/wiki/Handbook:PPC64/Installation/Bootloader/fr "Handbook:PPC64/Installation/Bootloader/fr (100% translated)")
- [italiano](/wiki/Handbook:PPC64/Installation/Bootloader/it "Handbook:PPC64/Installation/Bootloader/it (100% translated)")
- magyar
- [polski](/wiki/Handbook:PPC64/Installation/Bootloader/pl "Handbook:PPC64/Installation/Bootloader (100% translated)")
- [português do Brasil](/wiki/Handbook:PPC64/Installation/Bootloader/pt-br "Handbook:PPC64/Installation/Bootloader/pt-br (50% translated)")
- [русский](/wiki/Handbook:PPC64/Installation/Bootloader/ru "Handbook:PPC64/Installation/Bootloader (100% translated)")
- [தமிழ்](/wiki/Handbook:PPC64/Installation/Bootloader/ta "கையேடு:PPC64/நிறுவல்/துவக்கஏற்றி (100% translated)")
- [中文（中国大陆）‎](/wiki/Handbook:PPC64/Installation/Bootloader/zh-cn "手册：PPC64/安装/配置系统引导程序Bootloader (100% translated)")
- [日本語](/wiki/Handbook:PPC64/Installation/Bootloader/ja "ハンドブック:PPC64/インストール/ブートローダー (100% translated)")
- [한국어](/wiki/Handbook:PPC64/Installation/Bootloader/ko "Handbook:PPC64/Installation/Bootloader/ko (100% translated)")

[PPC64 kézikönyv](/wiki/Handbook:PPC64/hu "Handbook:PPC64/hu")[A Gentoo Linux telepítése](/wiki/Handbook:PPC64/Full/Installation/hu "Handbook:PPC64/Full/Installation/hu")[A telepítésről](/wiki/Handbook:PPC64/Installation/About/hu "Handbook:PPC64/Installation/About/hu")[Telepítőképfájl kiválasztása](/wiki/Handbook:PPC64/Installation/Media/hu "Handbook:PPC64/Installation/Media/hu")[Hálózat beállítása](/wiki/Handbook:PPC64/Installation/Networking/hu "Handbook:PPC64/Installation/Networking/hu")[Adathordozók előkészítése](/wiki/Handbook:PPC64/Installation/Disks/hu "Handbook:PPC64/Installation/Disks/hu")[Stage fájl](/wiki/Handbook:PPC64/Installation/Stage/hu "Handbook:PPC64/Installation/Stage/hu")[Alaprendszer telepítése](/wiki/Handbook:PPC64/Installation/Base/hu "Handbook:PPC64/Installation/Base/hu")[Kernel beállítása](/wiki/Handbook:PPC64/Installation/Kernel/hu "Handbook:PPC64/Installation/Kernel/hu")[Rendszer beállítása](/wiki/Handbook:PPC64/Installation/System/hu "Handbook:PPC64/Installation/System/hu")[Eszközök telepítése](/wiki/Handbook:PPC64/Installation/Tools/hu "Handbook:PPC64/Installation/Tools/hu")Bootloader beállítása[Telepítés véglegesítése](/wiki/Handbook:PPC64/Installation/Finalizing/hu "Handbook:PPC64/Installation/Finalizing/hu")[Munka a Gentoo rendszerrel](/wiki/Handbook:PPC64/Full/Working/hu "Handbook:PPC64/Full/Working/hu")[Portage bemutatása](/wiki/Handbook:PPC64/Working/Portage/hu "Handbook:PPC64/Working/Portage/hu")[USE jelölőzászlók](/wiki/Handbook:PPC64/Working/USE/hu "Handbook:PPC64/Working/USE/hu")[Portage jellemzői](/wiki/Handbook:PPC64/Working/Features/hu "Handbook:PPC64/Working/Features/hu")[Init-szkript rendszer](/wiki/Handbook:PPC64/Working/Initscripts/hu "Handbook:PPC64/Working/Initscripts/hu")[Környezeti változók](/wiki/Handbook:PPC64/Working/EnvVar/hu "Handbook:PPC64/Working/EnvVar/hu")[Munka a Portage szoftvercsomag-kezelővel](/wiki/Handbook:PPC64/Full/Portage/hu "Handbook:PPC64/Full/Portage/hu")[Fájlok és könyvtárak](/wiki/Handbook:PPC64/Portage/Files/hu "Handbook:PPC64/Portage/Files/hu")[Változók](/wiki/Handbook:PPC64/Portage/Variables/hu "Handbook:PPC64/Portage/Variables/hu")[Szoftverágak keverése](/wiki/Handbook:PPC64/Portage/Branches/hu "Handbook:PPC64/Portage/Branches/hu")[További eszközök](/wiki/Handbook:PPC64/Portage/Tools/hu "Handbook:PPC64/Portage/Tools/hu")[Egyéni szoftvercsomag-tárolók](/wiki/Handbook:PPC64/Portage/CustomTree/hu "Handbook:PPC64/Portage/CustomTree/hu")[Fejlett funkciók](/wiki/Handbook:PPC64/Portage/Advanced/hu "Handbook:PPC64/Portage/Advanced/hu")[Hálózat beállítása OpenRC init-rendszeren](/wiki/Handbook:PPC64/Full/Networking/hu "Handbook:PPC64/Full/Networking/hu")[Munka elkezdése](/wiki/Handbook:PPC64/Networking/Introduction/hu "Handbook:PPC64/Networking/Introduction/hu")[Fejlett beállítások](/wiki/Handbook:PPC64/Networking/Advanced/hu "Handbook:PPC64/Networking/Advanced/hu")[Moduláris hálózat](/wiki/Handbook:PPC64/Networking/Modular/hu "Handbook:PPC64/Networking/Modular/hu")[Vezeték nélküli (Wi-Fi)](/wiki/Handbook:PPC64/Networking/Wireless/hu "Handbook:PPC64/Networking/Wireless/hu")[Funkcionalitás hozzáadása](/wiki/Handbook:PPC64/Networking/Extending/hu "Handbook:PPC64/Networking/Extending/hu")[Dinamikus menedzsment](/wiki/Handbook:PPC64/Networking/Dynamic/hu "Handbook:PPC64/Networking/Dynamic/hu")

Miután a kernel be van állítva és a forráskódja le van fordítva, valamint a szükséges rendszerbeállítás-fájlok megfelelően ki lettek töltve, elérkezett az idő egy olyan program telepítésére, amely a rendszer bootolásakor elindítja a kernelt. Egy ilyen programot boot loader (rendszerbetöltő) néven nevezünk.

**Note**

Jelenleg a Petitboot használata a Talos rendszereken Gentoo operációs rendszerben nincs dokumentálva. Kérjük, hogy adja hozzá a lépéseket a [TalosII#Bootloader](/wiki/TalosII#Bootloader.2Fhu "TalosII") oldalhoz, és értesítsen ezen a Vitaoldalon, amikor készen áll a kézikönyvbe való összeolvasztásra.

## Contents

- [1GRUB használata](#GRUB_haszn.C3.A1lata)
  - [1.1Telepítés](#Telep.C3.ADt.C3.A9s)
  - [1.2Mac hardver (G5)](#Mac_hardver_.28G5.29)
    - [1.2.1A bootstrap partíció beállítása](#A_bootstrap_part.C3.ADci.C3.B3_be.C3.A1ll.C3.ADt.C3.A1sa)
    - [1.2.2GRUB beállítása](#GRUB_be.C3.A1ll.C3.ADt.C3.A1sa)
  - [1.3IBM hardver](#IBM_hardver)
    - [1.3.1GRUB beállítása](#GRUB_be.C3.A1ll.C3.ADt.C3.A1sa_2)
    - [1.3.2Grub config](#Grub_config)
- [2Rendszer újraindítása](#Rendszer_.C3.BAjraind.C3.ADt.C3.A1sa)

## GRUB használata

A GRUB egy bootloader PPC64-alapú Linux számítógépek számára.

### Telepítés

`root #` `emerge --ask sys-boot/grub`

### Mac hardver (G5)

#### A bootstrap partíció beállítása

Először készítse elő a bootstrap partíciót, amely a [lemez előkészítése](/wiki/Handbook:PPC64/Installation/Disks "Handbook:PPC64/Installation/Disks") lépés során lett létrehozva. Az példában szereplő partíció a /dev/sda2 kell, hogy legyen. Opcionálisan ellenőrizheti ezt a parted használatával:

Amennyiben szükséges, akkor cserélje ki a /dev/sda jelölést a megfelelő eszköz jelölésére.

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

Ebben a kimenetben a 2-es partíció tartalmazza a bootstrap információt, így a /dev/sda2 a helyes útvonal a használathoz.

Formázza ezt a partíciót [HFS](/wiki/HFS "HFS") formátumúra a hformat parancs segítségével, amely a [sys-fs/hfsutils](https://packages.gentoo.org/packages/sys-fs/hfsutils) szoftvercsomag része.

`root #` `dd if=/dev/zero of=/dev/sda2 bs=512`

`root #` `hformat -l bootstrap /dev/sda2`

Hozzon létre egy könyvtárat a bootstrap partíció felcsatolásához, majd csatolja azt fel:

`root #` `mkdir /boot/NWBB
`

`root #` `mount --types hfs /dev/sda2 /boot/NWBB`

#### GRUB beállítása

`root #` `grub-install --macppc-directory=/boot/NWBB /dev/sda2`

Ha a telepítés hibák nélkül zajlik le, akkor csatolja le a bootstrap partíciót.

`root #` `umount /boot/NWBB`

Ezután áldja meg a Szentlélek erejével a partíciót, hogy bootolható legyen:

`root #` `hmount /dev/sda2
`

`root #` `hattrib -t tbxi -c UNIX :System:Library:CoreServices:BootX
`

`root #` `hattrib -b :System:Library:CoreServices
`

`root #` `humount`

Végül készítse el a grub.cfg fájlt:

`root #` `grub-mkconfig -o /boot/grub/grub.cfg`

### IBM hardver

GRUB beállítása az IBM hardverén olyan egyszerű, mint:

#### GRUB beállítása

`root #` `grub-install /dev/sda1`

**Note**

/dev/sda1 is the PReP boot partition made in the partitioning stage

#### Grub config

Végül készítse el a grub.cfg fájlt:

`root #` `grub-mkconfig -o /boot/grub/grub.cfg`

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

Miután újraindította a számítógépet, és belépett a frissen feltelepített Gentoo környezetben, bölcs dolog [véglegesíteni a Gentoo telepítést](/wiki/Handbook:PPC64/Installation/Finalizing/hu "Handbook:PPC64/Installation/Finalizing/hu").

[← Eszközök telepítése](/wiki/Handbook:PPC64/Installation/Tools/hu "Previous part") [Kezdőlap](/wiki/Handbook:PPC64/hu "Handbook:PPC64/hu") [Telepítés véglegesítése →](/wiki/Handbook:PPC64/Installation/Finalizing/hu "Next part")