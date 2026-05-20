# Bootloader

Other languages:

- [Deutsch](/wiki/Handbook:SPARC/Installation/Bootloader/de "Handbuch:SPARC/Installation/Bootloader (100% translated)")
- [English](/wiki/Handbook:SPARC/Installation/Bootloader "Handbook:SPARC/Installation/Bootloader (100% translated)")
- [español](/wiki/Handbook:SPARC/Installation/Bootloader/es "Manual de Gentoo: SPARC/Instalación/Arranque (100% translated)")
- [français](/wiki/Handbook:SPARC/Installation/Bootloader/fr "Handbook:SPARC/Installation/Bootloader/fr (100% translated)")
- [italiano](/wiki/Handbook:SPARC/Installation/Bootloader/it "Handbook:SPARC/Installation/Bootloader/it (50% translated)")
- magyar
- [polski](/wiki/Handbook:SPARC/Installation/Bootloader/pl "Handbook:SPARC/Installation/Bootloader (100% translated)")
- [português do Brasil](/wiki/Handbook:SPARC/Installation/Bootloader/pt-br "Handbook:SPARC/Installation/Bootloader/pt-br (50% translated)")
- [русский](/wiki/Handbook:SPARC/Installation/Bootloader/ru "Handbook:SPARC/Installation/Bootloader (100% translated)")
- [தமிழ்](/wiki/Handbook:SPARC/Installation/Bootloader/ta "கையேடு:SPARC/நிறுவல்/துவக்கஏற்றி (100% translated)")
- [中文（中国大陆）‎](/wiki/Handbook:SPARC/Installation/Bootloader/zh-cn "手册：SPARC/安装/配置系统引导程序Bootloader (100% translated)")
- [日本語](/wiki/Handbook:SPARC/Installation/Bootloader/ja "ハンドブック:SPARC/インストール/ブートローダー (100% translated)")
- [한국어](/wiki/Handbook:SPARC/Installation/Bootloader/ko "Handbook:SPARC/Installation/Bootloader/ko (100% translated)")

[SPARC kézikönyv](/wiki/Handbook:SPARC/hu "Handbook:SPARC/hu")[A Gentoo Linux telepítése](/wiki/Handbook:SPARC/Full/Installation/hu "Handbook:SPARC/Full/Installation/hu")[A telepítésről](/wiki/Handbook:SPARC/Installation/About/hu "Handbook:SPARC/Installation/About/hu")[Telepítőképfájl kiválasztása](/wiki/Handbook:SPARC/Installation/Media/hu "Handbook:SPARC/Installation/Media/hu")[Hálózat beállítása](/wiki/Handbook:SPARC/Installation/Networking/hu "Handbook:SPARC/Installation/Networking/hu")[Adathordozók előkészítése](/wiki/Handbook:SPARC/Installation/Disks/hu "Handbook:SPARC/Installation/Disks/hu")[Stage fájl](/wiki/Handbook:SPARC/Installation/Stage/hu "Handbook:SPARC/Installation/Stage/hu")[Alaprendszer telepítése](/wiki/Handbook:SPARC/Installation/Base/hu "Handbook:SPARC/Installation/Base/hu")[Kernel beállítása](/wiki/Handbook:SPARC/Installation/Kernel/hu "Handbook:SPARC/Installation/Kernel/hu")[Rendszer beállítása](/wiki/Handbook:SPARC/Installation/System/hu "Handbook:SPARC/Installation/System/hu")[Eszközök telepítése](/wiki/Handbook:SPARC/Installation/Tools/hu "Handbook:SPARC/Installation/Tools/hu")Bootloader beállítása[Telepítés véglegesítése](/wiki/Handbook:SPARC/Installation/Finalizing/hu "Handbook:SPARC/Installation/Finalizing/hu")[Munka a Gentoo rendszerrel](/wiki/Handbook:SPARC/Full/Working/hu "Handbook:SPARC/Full/Working/hu")[Portage bemutatása](/wiki/Handbook:SPARC/Working/Portage/hu "Handbook:SPARC/Working/Portage/hu")[USE jelölőzászlók](/wiki/Handbook:SPARC/Working/USE/hu "Handbook:SPARC/Working/USE/hu")[Portage jellemzői](/wiki/Handbook:SPARC/Working/Features/hu "Handbook:SPARC/Working/Features/hu")[Init-szkript rendszer](/wiki/Handbook:SPARC/Working/Initscripts/hu "Handbook:SPARC/Working/Initscripts/hu")[Környezeti változók](/wiki/Handbook:SPARC/Working/EnvVar/hu "Handbook:SPARC/Working/EnvVar/hu")[Munka a Portage szoftvercsomag-kezelővel](/wiki/Handbook:SPARC/Full/Portage/hu "Handbook:SPARC/Full/Portage/hu")[Fájlok és könyvtárak](/wiki/Handbook:SPARC/Portage/Files/hu "Handbook:SPARC/Portage/Files/hu")[Változók](/wiki/Handbook:SPARC/Portage/Variables/hu "Handbook:SPARC/Portage/Variables/hu")[Szoftverágak keverése](/wiki/Handbook:SPARC/Portage/Branches/hu "Handbook:SPARC/Portage/Branches/hu")[További eszközök](/wiki/Handbook:SPARC/Portage/Tools/hu "Handbook:SPARC/Portage/Tools/hu")[Egyéni szoftvercsomag-tárolók](/wiki/Handbook:SPARC/Portage/CustomTree/hu "Handbook:SPARC/Portage/CustomTree/hu")[Fejlett funkciók](/wiki/Handbook:SPARC/Portage/Advanced/hu "Handbook:SPARC/Portage/Advanced/hu")[Hálózat beállítása OpenRC init-rendszeren](/wiki/Handbook:SPARC/Full/Networking/hu "Handbook:SPARC/Full/Networking/hu")[Munka elkezdése](/wiki/Handbook:SPARC/Networking/Introduction/hu "Handbook:SPARC/Networking/Introduction/hu")[Fejlett beállítások](/wiki/Handbook:SPARC/Networking/Advanced/hu "Handbook:SPARC/Networking/Advanced/hu")[Moduláris hálózat](/wiki/Handbook:SPARC/Networking/Modular/hu "Handbook:SPARC/Networking/Modular/hu")[Vezeték nélküli (Wi-Fi)](/wiki/Handbook:SPARC/Networking/Wireless/hu "Handbook:SPARC/Networking/Wireless/hu")[Funkcionalitás hozzáadása](/wiki/Handbook:SPARC/Networking/Extending/hu "Handbook:SPARC/Networking/Extending/hu")[Dinamikus menedzsment](/wiki/Handbook:SPARC/Networking/Dynamic/hu "Handbook:SPARC/Networking/Dynamic/hu")

## Contents

- [1GRUB](#GRUB)
  - [1.1Emerge](#Emerge)
  - [1.2Telepítés](#Telep.C3.ADt.C3.A9s)
    - [1.2.1GPT](#GPT)
    - [1.2.2Sun partíciós táblázat](#Sun_part.C3.ADci.C3.B3s_t.C3.A1bl.C3.A1zat)
  - [1.3Beállítás](#Be.C3.A1ll.C3.ADt.C3.A1s)
- [2SILO, a SPARC bootloader](#SILO.2C_a_SPARC_bootloader)
- [3Rendszer újraindítása](#Rendszer_.C3.BAjraind.C3.ADt.C3.A1sa)

## GRUB

Amikor a [64 bites profilt választja](/wiki/Handbook:SPARC/Installation/Base/hu#Choosing_the_right_profile "Handbook:SPARC/Installation/Base/hu") a telepítés során, akkor a [GRUB](/wiki/GRUB/hu "GRUB/hu") az egyetlen támogatott bootloader.

### Emerge

A GRUB bootloadernek a profil alapján automatikusan megfelelően kell beállítva kell lennie a platformhoz. Azonban, ha egyértelművé szeretné tenni, akkor adja meg az alábbi módon:

`root #` `echo 'GRUB_PLATFORMS="ieee1275"' >> /etc/portage/make.conf`

`root #` `emerge --ask --verbose sys-boot/grub`

A GRUB szoftver most már le lett töltve az operációs rendszerre, de még nincs beállítva a rendszerünk bootloaderjeként.

### Telepítés

#### GPT

Ha a [lemezt GPT-vel particionáltuk](/wiki/Handbook:SPARC/Installation/Disks/hu#Using_fdisk_to_partition_the_disk "Handbook:SPARC/Installation/Disks/hu") (ez az előnyben részesített módszer), akkor telepítse a GRUB bootloadert a BIOS boot partícióra. Feltételezve, hogy az első fizikai adathordozó (ahonnan a rendszer indul) /dev/sda. Ezt a munkát a következő parancsok végzik el:

`root #` `grub-install --target=sparc64-ieee1275 --recheck /dev/sda`

**Tip**

Ahhoz, hogy megtalálja a boot eszköz karakterláncát, amelyet be kell írni a firmware-be, használja a grub-ofpathname szoftvert. Ha a BIOS boot partíció az első partíció a fizikai adathordozón, akkor válassza ki az egész adathordozót:

`root #` `grub-ofpathname /dev/sda`

Más esetben válassza ki a BIOS boot partíciót:

`root #` `grub-ofpathname /dev/sda2`

#### Sun partíciós táblázat

Ha a fizikai adathordozót Sun partíciós táblával particionáltuk, akkor a GRUB bootladert blokklistek használatával kell telepíteni. Ebben az üzemmódban a fizikai adathordozó helyett adja meg annak a partíciónak az elérési útját, amelyre a /boot/grub van csatolva.

`root #` `grub-install --target=sparc64-ieee1275 --recheck --force --skip-fs-probe /dev/sda1`

### Beállítás

Ezután generálja a GRUB2 beállítást a /etc/default/grub fájlban és a /etc/grub.d szkriptekben megadott felhasználói beállítás alapján. A legtöbb esetben a felhasználóknak nincs szükségük beállításra, mivel a GRUB2 automatikusan felismeri, melyik bináris kernelképfájlt kell betölteni (a legmagasabb verziót a /boot/ könyvtárban) és mi a gyökérfájlrendszer. Az is lehetséges, hogy kernelparamétereket adjon hozzá a /etc/default/grub fájlban a GRUB\_CMDLINE\_LINUX változó használatával.

A végleges GRUB2 beállítás generálásához futtassa a grub-mkconfig parancsot:

`root #` `grub-mkconfig -o /boot/grub/grub.cfg`

```
Generating grub.cfg ...
Found linux image: /boot/vmlinuz-6.19.3-gentoo
Found initrd image: /boot/initramfs-genkernel-sparc-6.19.3-gentoo
done

```

A parancs kimenetének említenie kell, hogy legalább egy Linux bináris képfájl található, mivel ezek szükségesek az operációs rendszer indításához. Ha bináris initramfs képfájlt használunk, vagy a bináris kernelképfájlt a genkernel segítségével hoztuk létre, akkor a megfelelő initrd bináris képfájlt is észlelnie kell. Ha ez nem így van, akkor lépjen a /boot/ könyvtárba, és ellenőrizze a tartalmat a ls parancs segítségével. Ha a bináris képfájlok valóban hiányoznak, akkor térjen vissza a kernel beállításához és a telepítési utasításokhoz.

## SILO, a SPARC bootloader

Amikor a [32 bites profilt választja](/wiki/Handbook:SPARC/Installation/Base/hu#Choosing_the_right_profile "Handbook:SPARC/Installation/Base/hu") a telepítés során, akkor a [SILO](https://git.kernel.org/?p=linux/kernel/git/davem/silo.git;a=summary) (Sparc Improved boot LOader) az egyetlen támogatott bootloader.

`root #` `emerge --ask sys-boot/silo`

Ezután hozza létre a /etc/silo.conf fájlt:

`root #` `nano -w /etc/silo.conf`

Az alábbiakban egy példa silo.conf fájl látható. Ez a könyvben végig használt particionálási sémát alkalmazza, kernel-6.19.3-gentoo mint kernelképfájl, és initramfs-genkernel-sparc64-6.19.3-gentoo mint initramfs képfájl.

FILE **`/etc/silo.conf`** **Példa beállításfájl**

```
partition = 1         # Boot partition (= root partition)
root = /dev/sda1      # Root partition
timeout = 150         # Wait 15 seconds before booting the default section

image = /boot/kernel-6.19.3-gentoo
  label = linux
  append = "initrd=/boot/initramfs-genkernel-sparc64-6.19.3-gentoo root=/dev/sda1"

```

Amikor a Portage által biztosított példát, a silo.conf fájlt használja, győződjön meg arról, hogy minden szükségtelen sort kikommentel.

Ha az a fizikai adathordozó, amelyre a SILO-t (mint bootloadert) telepíteni kell, eltér attól a fizikai adathordozótól, amelyen a /etc/silo.conf található, akkor először másolja át a /etc/silo.conf fájlt az adott adathordozó egyik partíciójára. Ha a /boot/ külön partícióként található azon az adathordozón, akkor másolja át a beállításfájlt a /boot/ könyvtárba, és futtassa a /sbin/silo parancsot:

`root #` `cp /etc/silo.conf /boot
`

`root #` `/sbin/silo -C /boot/silo.conf`

```
/boot/silo.conf appears to be valid

```

Ellenkező esetben egyszerűen futtassa a /sbin/silo parancsot:

`root #` `/sbin/silo`

```
/etc/silo.conf appears to be valid

```

**Note**

Futtassa újra a silo-t (szükség esetén paraméterekkel) minden alkalommal, miután frissítette vagy telepítette a [sys-boot/silo](https://packages.gentoo.org/packages/sys-boot/silo) szoftvercsomagot.

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

Miután újraindította a számítógépet, és belépett a frissen feltelepített Gentoo környezetben, bölcs dolog [véglegesíteni a Gentoo telepítést](/wiki/Handbook:SPARC/Installation/Finalizing/hu "Handbook:SPARC/Installation/Finalizing/hu").

[← Eszközök telepítése](/wiki/Handbook:SPARC/Installation/Tools/hu "Previous part") [Kezdőlap](/wiki/Handbook:SPARC/hu "Handbook:SPARC/hu") [Telepítés véglegesítése →](/wiki/Handbook:SPARC/Installation/Finalizing/hu "Next part")