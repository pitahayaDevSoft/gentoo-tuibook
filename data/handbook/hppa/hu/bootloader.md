# Bootloader

Other languages:

- [Deutsch](/wiki/Handbook:HPPA/Installation/Bootloader/de "Handbuch:HPPA/Installation/Bootloader (100% translated)")
- [English](/wiki/Handbook:HPPA/Installation/Bootloader "Handbook:HPPA/Installation/Bootloader (100% translated)")
- [español](/wiki/Handbook:HPPA/Installation/Bootloader/es "Manual de Gentoo: HPPA/Instalación/Arranque (100% translated)")
- [français](/wiki/Handbook:HPPA/Installation/Bootloader/fr "Handbook:HPPA/Installation/Bootloader/fr (100% translated)")
- [italiano](/wiki/Handbook:HPPA/Installation/Bootloader/it "Handbook:HPPA/Installation/Bootloader/it (100% translated)")
- magyar
- [polski](/wiki/Handbook:HPPA/Installation/Bootloader/pl "Handbook:HPPA/Installation/Bootloader (100% translated)")
- [português do Brasil](/wiki/Handbook:HPPA/Installation/Bootloader/pt-br "Manual:HPPA/Instalação/Gerenciador de Boot (100% translated)")
- [русский](/wiki/Handbook:HPPA/Installation/Bootloader/ru "Handbook:HPPA/Installation/Bootloader (100% translated)")
- [தமிழ்](/wiki/Handbook:HPPA/Installation/Bootloader/ta "கையேடு:HPPA/நிறுவல்/துவக்கஏற்றி (100% translated)")
- [中文（中国大陆）‎](/wiki/Handbook:HPPA/Installation/Bootloader/zh-cn "手册：HPPA/安装/配置系统引导程序Bootloader (100% translated)")
- [日本語](/wiki/Handbook:HPPA/Installation/Bootloader/ja "ハンドブック:HPPA/インストール/ブートローダー (100% translated)")
- [한국어](/wiki/Handbook:HPPA/Installation/Bootloader/ko "Handbook:HPPA/Installation/Bootloader/ko (100% translated)")

[HPPA kézikönyv](/wiki/Handbook:HPPA/hu "Handbook:HPPA/hu")[A Gentoo Linux telepítése](/wiki/Handbook:HPPA/Full/Installation/hu "Handbook:HPPA/Full/Installation/hu")[A telepítésről](/wiki/Handbook:HPPA/Installation/About/hu "Handbook:HPPA/Installation/About/hu")[Telepítőképfájl kiválasztása](/wiki/Handbook:HPPA/Installation/Media/hu "Handbook:HPPA/Installation/Media/hu")[Hálózat beállítása](/wiki/Handbook:HPPA/Installation/Networking/hu "Handbook:HPPA/Installation/Networking/hu")[Adathordozók előkészítése](/wiki/Handbook:HPPA/Installation/Disks/hu "Handbook:HPPA/Installation/Disks/hu")[Stage fájl](/wiki/Handbook:HPPA/Installation/Stage/hu "Handbook:HPPA/Installation/Stage/hu")[Alaprendszer telepítése](/wiki/Handbook:HPPA/Installation/Base/hu "Handbook:HPPA/Installation/Base/hu")[Kernel beállítása](/wiki/Handbook:HPPA/Installation/Kernel/hu "Handbook:HPPA/Installation/Kernel/hu")[Rendszer beállítása](/wiki/Handbook:HPPA/Installation/System/hu "Handbook:HPPA/Installation/System/hu")[Eszközök telepítése](/wiki/Handbook:HPPA/Installation/Tools/hu "Handbook:HPPA/Installation/Tools/hu")Bootloader beállítása[Telepítés véglegesítése](/wiki/Handbook:HPPA/Installation/Finalizing/hu "Handbook:HPPA/Installation/Finalizing/hu")[Munka a Gentoo rendszerrel](/wiki/Handbook:HPPA/Full/Working/hu "Handbook:HPPA/Full/Working/hu")[Portage bemutatása](/wiki/Handbook:HPPA/Working/Portage/hu "Handbook:HPPA/Working/Portage/hu")[USE jelölőzászlók](/wiki/Handbook:HPPA/Working/USE/hu "Handbook:HPPA/Working/USE/hu")[Portage jellemzői](/wiki/Handbook:HPPA/Working/Features/hu "Handbook:HPPA/Working/Features/hu")[Init-szkript rendszer](/wiki/Handbook:HPPA/Working/Initscripts/hu "Handbook:HPPA/Working/Initscripts/hu")[Környezeti változók](/wiki/Handbook:HPPA/Working/EnvVar/hu "Handbook:HPPA/Working/EnvVar/hu")[Munka a Portage szoftvercsomag-kezelővel](/wiki/Handbook:HPPA/Full/Portage/hu "Handbook:HPPA/Full/Portage/hu")[Fájlok és könyvtárak](/wiki/Handbook:HPPA/Portage/Files/hu "Handbook:HPPA/Portage/Files/hu")[Változók](/wiki/Handbook:HPPA/Portage/Variables/hu "Handbook:HPPA/Portage/Variables/hu")[Szoftverágak keverése](/wiki/Handbook:HPPA/Portage/Branches/hu "Handbook:HPPA/Portage/Branches/hu")[További eszközök](/wiki/Handbook:HPPA/Portage/Tools/hu "Handbook:HPPA/Portage/Tools/hu")[Egyéni szoftvercsomag-tárolók](/wiki/Handbook:HPPA/Portage/CustomTree/hu "Handbook:HPPA/Portage/CustomTree/hu")[Fejlett funkciók](/wiki/Handbook:HPPA/Portage/Advanced/hu "Handbook:HPPA/Portage/Advanced/hu")[Hálózat beállítása OpenRC init-rendszeren](/wiki/Handbook:HPPA/Full/Networking/hu "Handbook:HPPA/Full/Networking/hu")[Munka elkezdése](/wiki/Handbook:HPPA/Networking/Introduction/hu "Handbook:HPPA/Networking/Introduction/hu")[Fejlett beállítások](/wiki/Handbook:HPPA/Networking/Advanced/hu "Handbook:HPPA/Networking/Advanced/hu")[Moduláris hálózat](/wiki/Handbook:HPPA/Networking/Modular/hu "Handbook:HPPA/Networking/Modular/hu")[Vezeték nélküli (Wi-Fi)](/wiki/Handbook:HPPA/Networking/Wireless/hu "Handbook:HPPA/Networking/Wireless/hu")[Funkcionalitás hozzáadása](/wiki/Handbook:HPPA/Networking/Extending/hu "Handbook:HPPA/Networking/Extending/hu")[Dinamikus menedzsment](/wiki/Handbook:HPPA/Networking/Dynamic/hu "Handbook:HPPA/Networking/Dynamic/hu")

## PALO telepítése

A PA-RISC platformon a boot loader neve palo. Először emerge segítségével hozza létre ezt a boot loader-t a rendszer számára:

`root #` `emerge --ask sys-boot/palo`

A beállításfájl a /etc/palo.conf helyen található. Az alábbiakban egy mintabeállítás látható:

**Warning**

Ezt a beállítást **módosítani kell**, miután az palo először futtatásra kerül! Tekintse meg az alábbiakat az első beállítás után.

FILE **`/etc/palo.conf`** **Példa az egyszerű PALO beállításra.**

```
--commandline=2/kernel-6.19.1-gentoo root=/dev/sda4
--recoverykernel=/vmlinux.old
# DELETE this line after running palo for the first time!
--init-partitioned=/dev/sda
# --format-as has two meanings depending on whether --init-partitioned or --update-partitioned is used. Keep this line.
--format-as=4

```

Az első sor megadja az palo számára a kernel helyét és azokat a boot paramétereket, amelyeket használnia kell. A `2/kernel-6.19.1-gentoo` karakterlánc azt jelenti, hogy a /kernel-6.19.1-gentoo nevű kernel a második partíción található. Fontos megjegyezni, hogy a kernelhez vezető útvonal a _boot_ partícióhoz viszonyítva van megadva, nem a root partícióhoz.

A második sor megadja, hogy melyik helyreállító kernelt kell használni. Ha ez az első telepítés, és még nincs helyreállító kernel, akkor kérjük, hogy szúrja ki ezt a sort. A harmadik sor azt jelzi, hogy melyik adathordozón található a palo.

Az adathordozó formázásához az palo bizonyos argumentumokkal futtatandó. Ez a példa az első partícióhoz az _ext4_ fájlrendszert használja:

`root #` `palo --format-as=4 --init-partitioned=/dev/sda`

A beállítás befejezése után egyszerűen futtassa az palo parancsot:

`root #` `palo`

A beállítást ezt követően frissíteni kell az első telepítés utáni használathoz:

FILE **`/etc/palo.conf`** **Példa az egyszerű PALO beállításra**

```
--commandline=2/kernel-6.19.1-gentoo root=/dev/sda4
--recoverykernel=/vmlinux.old
# Don't throw away the old partition, just update the existing one on `palo` runs.
--update-partitioned=/dev/sda
# --format-as has two meanings depending on whether --init-partitioned or --update-partitioned is used. Keep this line.
--format-as=4

```

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

Miután újraindította a számítógépet, és belépett a frissen feltelepített Gentoo környezetben, bölcs dolog [véglegesíteni a Gentoo telepítést](/wiki/Handbook:HPPA/Installation/Finalizing/hu "Handbook:HPPA/Installation/Finalizing/hu").

[← Eszközök telepítése](/wiki/Handbook:HPPA/Installation/Tools/hu "Previous part") [Kezdőlap](/wiki/Handbook:HPPA/hu "Handbook:HPPA/hu") [Telepítés véglegesítése →](/wiki/Handbook:HPPA/Installation/Finalizing/hu "Next part")