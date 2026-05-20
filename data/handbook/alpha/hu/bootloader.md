# Bootloader

Other languages:

- [Deutsch](/wiki/Handbook:Alpha/Installation/Bootloader/de "Handbuch:Alpha/Installation/Bootloader (100% translated)")
- [English](/wiki/Handbook:Alpha/Installation/Bootloader "Handbook:Alpha/Installation/Bootloader (100% translated)")
- [español](/wiki/Handbook:Alpha/Installation/Bootloader/es "Manual de Gentoo: Alpha/Instalación/Arranque (100% translated)")
- [français](/wiki/Handbook:Alpha/Installation/Bootloader/fr "Manuel:Alpha/Installation/Système d'amorçage (100% translated)")
- [italiano](/wiki/Handbook:Alpha/Installation/Bootloader/it "Handbook:Alpha/Installation/Bootloader/it (100% translated)")
- magyar
- [polski](/wiki/Handbook:Alpha/Installation/Bootloader/pl "Handbook:Alpha/Installation/Bootloader (100% translated)")
- [português do Brasil](/wiki/Handbook:Alpha/Installation/Bootloader/pt-br "Manual:Alpha/Instalação/Gerenciador de boot (100% translated)")
- [čeština](/wiki/Handbook:Alpha/Installation/Bootloader/cs "Handbook:Alpha/Installation/Bootloader/cs (50% translated)")
- [русский](/wiki/Handbook:Alpha/Installation/Bootloader/ru "Handbook:Alpha/Installation/Bootloader (100% translated)")
- [தமிழ்](/wiki/Handbook:Alpha/Installation/Bootloader/ta "கையேடு:Alpha/நிறுவல்/துவக்கஏற்றி (100% translated)")
- [中文（中国大陆）‎](/wiki/Handbook:Alpha/Installation/Bootloader/zh-cn "手册：Alpha/安装/配置系统引导程序Bootloader (100% translated)")
- [日本語](/wiki/Handbook:Alpha/Installation/Bootloader/ja "ハンドブック:Alpha/インストール/ブートローダー (100% translated)")
- [한국어](/wiki/Handbook:Alpha/Installation/Bootloader/ko "Handbook:Alpha/Installation/Bootloader/ko (100% translated)")

[Alpha kézikönyv](/wiki/Handbook:Alpha/hu "Handbook:Alpha/hu")[A Gentoo Linux telepítése](/wiki/Handbook:Alpha/Full/Installation/hu "Handbook:Alpha/Full/Installation/hu")[A telepítésről](/wiki/Handbook:Alpha/Installation/About/hu "Handbook:Alpha/Installation/About/hu")[Telepítőképfájl kiválasztása](/wiki/Handbook:Alpha/Installation/Media/hu "Handbook:Alpha/Installation/Media/hu")[Hálózat beállítása](/wiki/Handbook:Alpha/Installation/Networking/hu "Handbook:Alpha/Installation/Networking/hu")[Adathordozók előkészítése](/wiki/Handbook:Alpha/Installation/Disks/hu "Handbook:Alpha/Installation/Disks/hu")[Stage fájl](/wiki/Handbook:Alpha/Installation/Stage/hu "Handbook:Alpha/Installation/Stage/hu")[Alaprendszer telepítése](/wiki/Handbook:Alpha/Installation/Base/hu "Handbook:Alpha/Installation/Base/hu")[Kernel beállítása](/wiki/Handbook:Alpha/Installation/Kernel/hu "Handbook:Alpha/Installation/Kernel/hu")[Rendszer beállítása](/wiki/Handbook:Alpha/Installation/System/hu "Handbook:Alpha/Installation/System/hu")[Eszközök telepítése](/wiki/Handbook:Alpha/Installation/Tools/hu "Handbook:Alpha/Installation/Tools/hu")Bootloader beállítása[Telepítés véglegesítése](/wiki/Handbook:Alpha/Installation/Finalizing/hu "Handbook:Alpha/Installation/Finalizing/hu")[Munka a Gentoo rendszerrel](/wiki/Handbook:Alpha/Full/Working/hu "Handbook:Alpha/Full/Working/hu")[Portage bemutatása](/wiki/Handbook:Alpha/Working/Portage/hu "Handbook:Alpha/Working/Portage/hu")[USE jelölőzászlók](/wiki/Handbook:Alpha/Working/USE/hu "Handbook:Alpha/Working/USE/hu")[Portage jellemzői](/wiki/Handbook:Alpha/Working/Features/hu "Handbook:Alpha/Working/Features/hu")[Init-szkript rendszer](/wiki/Handbook:Alpha/Working/Initscripts/hu "Handbook:Alpha/Working/Initscripts/hu")[Környezeti változók](/wiki/Handbook:Alpha/Working/EnvVar/hu "Handbook:Alpha/Working/EnvVar/hu")[Munka a Portage szoftvercsomag-kezelővel](/wiki/Handbook:Alpha/Full/Portage/hu "Handbook:Alpha/Full/Portage/hu")[Fájlok és könyvtárak](/wiki/Handbook:Alpha/Portage/Files/hu "Handbook:Alpha/Portage/Files/hu")[Változók](/wiki/Handbook:Alpha/Portage/Variables/hu "Handbook:Alpha/Portage/Variables/hu")[Szoftverágak keverése](/wiki/Handbook:Alpha/Portage/Branches/hu "Handbook:Alpha/Portage/Branches/hu")[További eszközök](/wiki/Handbook:Alpha/Portage/Tools/hu "Handbook:Alpha/Portage/Tools/hu")[Egyéni szoftvercsomag-tárolók](/wiki/Handbook:Alpha/Portage/CustomTree/hu "Handbook:Alpha/Portage/CustomTree/hu")[Fejlett funkciók](/wiki/Handbook:Alpha/Portage/Advanced/hu "Handbook:Alpha/Portage/Advanced/hu")[Hálózat beállítása OpenRC init-rendszeren](/wiki/Handbook:Alpha/Full/Networking/hu "Handbook:Alpha/Full/Networking/hu")[Munka elkezdése](/wiki/Handbook:Alpha/Networking/Introduction/hu "Handbook:Alpha/Networking/Introduction/hu")[Fejlett beállítások](/wiki/Handbook:Alpha/Networking/Advanced/hu "Handbook:Alpha/Networking/Advanced/hu")[Moduláris hálózat](/wiki/Handbook:Alpha/Networking/Modular/hu "Handbook:Alpha/Networking/Modular/hu")[Vezeték nélküli (Wi-Fi)](/wiki/Handbook:Alpha/Networking/Wireless/hu "Handbook:Alpha/Networking/Wireless/hu")[Funkcionalitás hozzáadása](/wiki/Handbook:Alpha/Networking/Extending/hu "Handbook:Alpha/Networking/Extending/hu")[Dinamikus menedzsment](/wiki/Handbook:Alpha/Networking/Dynamic/hu "Handbook:Alpha/Networking/Dynamic/hu")

## Döntés meghozása

Most, hogy a kernel be van állítva és a forráskódból le van fordítva bináris futtatható formára, valamint a szükséges rendszerbeállítás-fájlok megfelelően ki vannak töltve, ideje telepíteni egy programot, amely a rendszer indulásakor elindítja a bináris kernelképfájlt. Egy ilyen programnak a neve a bootloader.

Számos bootloader létezik a Linux/Alpha számára. Válasszon egyet a támogatott bootloaderek közül, ne mindegyiket válassza. Dokumentáltuk az [aBoot](#Default:_Using_aboot) és az [MILO](#Alternative:_Using_MILO) bootloadert.

## Alapértelmezett: aBoot használata

**Note**

Az aboot csak ext2 és ext3 partíciók bootolását támogatja.

Először telepítse az aboot szoftvercsomagot az operációs rendszerre.

`root #` `emerge --ask sys-boot/aboot`

A következő lépés a boot adathordozónk bootolhatóvá tétele. Ez elindítja az aboot szoftvert a rendszer bootolásakor. A boot adathordozó bootolhatóvá tételét úgy végezzük, hogy az aboot bootloadert az adathordozó elejére ráírjuk.

`root #` `swriteboot -f3 /dev/sda /boot/bootlx
`

`root #` `abootconf /dev/sda 2
`

**Note**

Amikor a jelen fejezetben bemutatottól eltérő partíciós sémát használ, a parancsokat ennek megfelelően módosítani kell. Kérjük, hogy olvassa el a megfelelő man sógókat (man 8 swriteboot és man 8 abootconf). Továbbá, ha a root fájlrendszer JFS fájlrendszerrel fut, akkor gondoskodjon arról, hogy először csak olvasható módban legyen csatolva azáltal, hogy hozzáadja az "ro" opciót a kernelopciókhoz.

Bár az aboot most már telepítve van, még mindig szükség van egy beállításfájl elkészítésére. Az aboot számára minden beállításhoz csak egy sor szükséges, így ezt megtehetjük:

`root #` `echo '0:2/boot/vmlinux.gz root=/dev/sda3' > /etc/aboot.conf`

Ha a bináris Linux kernel képfájl létrehozása során egy initramfs bináris képfájl is készült az bootoláshoz, akkor szükséges a beállítás módosítása az initramfs fájlra való hivatkozással, és az initramfs tájékoztatásával arról, hogy hol található a valódi gyökéreszköz.

`root #` `echo '0:2/boot/vmlinux.gz initrd=/boot/initramfs-genkernel-alpha-6.19.1-gentoo root=/dev/sda3' > /etc/aboot.conf`

Ezenkívül lehetséges, hogy a Gentoo automatikusan bootoljon, ha beállít néhány SRM változót. Próbálja meg ezeket a változókat Linuxból beállítani, bár lehet, hogy könnyebb ezt magából az SRM konzolból megtenni.

`root #` `cd /proc/srm_environment/named_variables
`

`root #` `echo -n 0 > boot_osflags
`

`root #` `echo -n '' > boot_file
`

`root #` `echo -n 'BOOT' > auto_action
`

`root #` `echo -n 'dkc100' > bootdef_dev`

Természetesen helyettesítse a dkc100-at bármilyen boot eszközzel.

Ahhoz, hogy a jövőben újra elérje az SRM konzolt (a Gentoo telepítés helyreállításához, néhány változó módosításához vagy bármi más célból), egyszerűen nyomja meg a `Ctrl+C` billentyűgomb kombinációt az automatikus betöltési folyamat megszakításához.

Amikor soros konzolt használ a telepítéshez, ne felejtse el beilleszteni a soros konzol boot jelölőzászlót az aboot.conf fájlban. További információért tekintse meg a /etc/aboot.conf.example fájlt.

Az aboot most már be van állítva és használatra kész. Folytassa a következővel: [Rendszer újraindítása](#Rebooting_the_system).

Folytassa a következővel: [Rendszer újraindítása](#Rebooting_the_system).

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

Miután újraindította a számítógépet, és belépett a frissen feltelepített Gentoo környezetben, bölcs dolog [véglegesíteni a Gentoo telepítést](/wiki/Handbook:Alpha/Installation/Finalizing/hu "Handbook:Alpha/Installation/Finalizing/hu").

[← Eszközök telepítése](/wiki/Handbook:Alpha/Installation/Tools/hu "Previous part") [Kezdőlap](/wiki/Handbook:Alpha/hu "Handbook:Alpha/hu") [Telepítés véglegesítése →](/wiki/Handbook:Alpha/Installation/Finalizing/hu "Next part")