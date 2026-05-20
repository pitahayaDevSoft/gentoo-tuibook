# Bootloader

Other languages:

- [Deutsch](/wiki/Handbook:MIPS/Installation/Bootloader/de "Handbuch:MIPS/Installation/Bootloader (100% translated)")
- [English](/wiki/Handbook:MIPS/Installation/Bootloader "Handbook:MIPS/Installation/Bootloader (100% translated)")
- [español](/wiki/Handbook:MIPS/Installation/Bootloader/es "Manual de Gentoo: MIPS/Instalación/Arranque (100% translated)")
- [français](/wiki/Handbook:MIPS/Installation/Bootloader/fr "Handbook:MIPS/Installation/Bootloader/fr (100% translated)")
- [italiano](/wiki/Handbook:MIPS/Installation/Bootloader/it "Handbook:MIPS/Installation/Bootloader/it (100% translated)")
- [magyar](/wiki/Handbook:MIPS/Installation/Bootloader/hu "Handbook:MIPS/Installation/Bootloader/hu (100% translated)")
- [polski](/wiki/Handbook:MIPS/Installation/Bootloader/pl "Handbook:MIPS/Installation/Bootloader (100% translated)")
- [português do Brasil](/wiki/Handbook:MIPS/Installation/Bootloader/pt-br "Manual:MIPS/Instalação/Gerenciador de Boot (100% translated)")
- [русский](/wiki/Handbook:MIPS/Installation/Bootloader/ru "Handbook:MIPS/Installation/Bootloader (100% translated)")
- [தமிழ்](/wiki/Handbook:MIPS/Installation/Bootloader/ta "கையேடு:MIPS/நிறுவல்/துவக்கஏற்றி (100% translated)")
- [中文（中国大陆）‎](/wiki/Handbook:MIPS/Installation/Bootloader/zh-cn "手册：MIPS/安装/配置系统引导程序Bootloader (100% translated)")
- 日本語
- [한국어](/wiki/Handbook:MIPS/Installation/Bootloader/ko "Handbook:MIPS/Installation/Bootloader/ko (100% translated)")

[MIPS ハンドブック](/wiki/Handbook:MIPS "Handbook:MIPS")[インストール](/wiki/Handbook:MIPS/Full/Installation "Handbook:MIPS/Full/Installation")[インストールについて](/wiki/Handbook:MIPS/Installation/About/ja "Handbook:MIPS/Installation/About/ja")[メディアの選択](/wiki/Handbook:MIPS/Installation/Media/ja "Handbook:MIPS/Installation/Media/ja")[ネットワーク設定](/wiki/Handbook:MIPS/Installation/Networking/ja "Handbook:MIPS/Installation/Networking/ja")[ディスクの準備](/wiki/Handbook:MIPS/Installation/Disks/ja "Handbook:MIPS/Installation/Disks/ja")[stage ファイル](/wiki/Handbook:MIPS/Installation/Stage/ja "Handbook:MIPS/Installation/Stage/ja")[ベースシステムのインストール](/wiki/Handbook:MIPS/Installation/Base/ja "Handbook:MIPS/Installation/Base/ja")[カーネルの設定](/wiki/Handbook:MIPS/Installation/Kernel/ja "Handbook:MIPS/Installation/Kernel/ja")[システムの設定](/wiki/Handbook:MIPS/Installation/System/ja "Handbook:MIPS/Installation/System/ja")[ツールのインストール](/wiki/Handbook:MIPS/Installation/Tools/ja "Handbook:MIPS/Installation/Tools/ja")ブートローダの設定[締めくくり](/wiki/Handbook:MIPS/Installation/Finalizing/ja "Handbook:MIPS/Installation/Finalizing/ja")[Gentoo の操作](/wiki/Handbook:MIPS/Full/Working "Handbook:MIPS/Full/Working")[Portage について](/wiki/Handbook:MIPS/Working/Portage/ja "Handbook:MIPS/Working/Portage/ja")[USE フラグ](/wiki/Handbook:MIPS/Working/USE/ja "Handbook:MIPS/Working/USE/ja")[Portage の機能](/wiki/Handbook:MIPS/Working/Features/ja "Handbook:MIPS/Working/Features/ja")[Init スクリプトシステム](/wiki/Handbook:MIPS/Working/Initscripts/ja "Handbook:MIPS/Working/Initscripts/ja")[環境変数](/wiki/Handbook:MIPS/Working/EnvVar/ja "Handbook:MIPS/Working/EnvVar/ja")[Portage の操作](/wiki/Handbook:MIPS/Full/Portage "Handbook:MIPS/Full/Portage")[ファイルとディレクトリ](/wiki/Handbook:MIPS/Portage/Files/ja "Handbook:MIPS/Portage/Files/ja")[変数](/wiki/Handbook:MIPS/Portage/Variables/ja "Handbook:MIPS/Portage/Variables/ja")[ソフトウェアブランチの併用](/wiki/Handbook:MIPS/Portage/Branches/ja "Handbook:MIPS/Portage/Branches/ja")[追加ツール](/wiki/Handbook:MIPS/Portage/Tools/ja "Handbook:MIPS/Portage/Tools/ja")[カスタムパッケージリポジトリ](/wiki/Handbook:MIPS/Portage/CustomTree/ja "Handbook:MIPS/Portage/CustomTree/ja")[高度な機能](/wiki/Handbook:MIPS/Portage/Advanced/ja "Handbook:MIPS/Portage/Advanced/ja")[OpenRC ネットワーク設定](/wiki/Handbook:MIPS/Full/Networking "Handbook:MIPS/Full/Networking")[はじめに](/wiki/Handbook:MIPS/Networking/Introduction/ja "Handbook:MIPS/Networking/Introduction/ja")[高度な設定](/wiki/Handbook:MIPS/Networking/Advanced/ja "Handbook:MIPS/Networking/Advanced/ja")[モジュール式ネットワーク](/wiki/Handbook:MIPS/Networking/Modular/ja "Handbook:MIPS/Networking/Modular/ja")[無線](/wiki/Handbook:MIPS/Networking/Wireless/ja "Handbook:MIPS/Networking/Wireless/ja")[機能の追加](/wiki/Handbook:MIPS/Networking/Extending/ja "Handbook:MIPS/Networking/Extending/ja")[動的な管理](/wiki/Handbook:MIPS/Networking/Dynamic/ja "Handbook:MIPS/Networking/Dynamic/ja")

## Contents

- [1arcload for Silicon Graphics machines](#arcload_for_Silicon_Graphics_machines)
- [2CoLo for Cobalt MicroServers](#CoLo_for_Cobalt_MicroServers)
  - [2.1Installing CoLo](#Installing_CoLo)
  - [2.2Configuring CoLo](#Configuring_CoLo)
- [3Setting up for serial console](#Setting_up_for_serial_console)
- [4Tweaking the SGI PROM](#Tweaking_the_SGI_PROM)
  - [4.1Setting generic PROM settings](#Setting_generic_PROM_settings)
  - [4.2Settings for direct volume-header booting](#Settings_for_direct_volume-header_booting)
  - [4.3Settings for arcload](#Settings_for_arcload)
- [5システムのリブート](#.E3.82.B7.E3.82.B9.E3.83.86.E3.83.A0.E3.81.AE.E3.83.AA.E3.83.96.E3.83.BC.E3.83.88)

## arcload for Silicon Graphics machines

arcload was written for machines that require 64-bit kernels, and therefore can't use arcboot (which can't easily be compiled as a 64-bit binary). It also works around peculiarities that arise when loading kernels directly from the volume header. To proceed with the installation, start with:

`root #` `emerge arcload dvhtool`

Once this has finished, find the arcload binary inside /usr/lib/arcload/. Now, two files exist:

- sashARCS: The 32-bit binary for Indy, Indigo2 (R4k), Challenge S and O2 systems
- sash64: The 64-bit binary for Octane/Octane2, Origin 200/2000 and Indigo2 Impact systems

Use dvhtool to install the appropriate binary for the system into the volume header:

For Indy/Indigo2/Challenge S/O2 users:

`root #` `dvhtool --unix-to-vh /usr/lib/arcload/sashARCS sashARCS`

For Indigo2 Impact/Octane/Octane2/Origin 200/Origin 2000 users:

`root #` `dvhtool --unix-to-vh /usr/lib/arcload/sash64 sash64`

**メモ**

The name sashARCS or sash64 does not have to be used, unless the operation is installing to the volume header of a bootable CD. For normal boot from hard-disk, it can be named whatever the user wants.

Now just use dvhtool to verify they are in the volume header:

`root #` `dvhtool --print-volume-directory`

```
----- directory entries -----
Entry #0, name "sash64", start 4, bytes 55859

```

The arc.cf file has a C-like syntax. For the full detail on how one configures it, see the arcload page on the Linux/MIPS wiki. In short, define a number of options, which are enabled and disabled at boot time using the OSLoadFilename variable.

ファイル **`arc.cf`** **An example arc.cf**

```
# ARCLoad Configuration

# Some default settings...
append  "root=/dev/sda5";
append  "ro";
append  "console=ttyS0,9600";

# The main definition. ip28 may be changed if desired.
ip28 {
        # Definition for a "working" kernel
        # Select this by setting OSLoadFilename="ip28(working)"
        working {
                description     "SGI Indigo2 Impact R10000\n\r";
                image system    "/working";
        }

        # Definition for a "new" kernel
        # Select this by setting OSLoadFilename="ip28(new)"
        new {
                description     "SGI Indigo2 Impact R10000 - Testing Kernel\n\r";
                image system    "/new";
        }

        # For debugging a kernel
        # Select this by setting OSLoadFilename="ip28(working,debug)"
        # or OSLoadFilename="ip28(new,debug)"
        debug {
                description     "Debug console";
                append          "init=/bin/bash";
        }
}

```

Starting with arcload-0.5, arc.cf and kernels may reside either in the volume header, or on a partition. To utilize this newer feature, place the files in the /boot/ partition (or / if the boot partition is not separate). arcload uses the filesystem driver code from the popular grub bootloader, and thus supports the same range of filesystems.

`root #` `dvhtool --unix-to-vh arc.cf arc.cf
`

`root #` `dvhtool --unix-to-vh /usr/src/linux/vmlinux new`

## CoLo for Cobalt MicroServers

### Installing CoLo

On Cobalt servers, these machines have a much less capable firmware installed on chip. The Cobalt BOOTROM is primitive, by comparison to the SGI PROM, and has a number of serious limitations.

- There's a 675kB (approximate) limit on kernels. The current size of Linux 2.4 makes it nearly impossible to make a kernel this size. Linux 2.6 and 3.x is totally out of the question.
- 64-bit kernels are not supported by the stock firmware (although these are highly experimental on Cobalt machines at this time)
- The shell is basic at best

To overcome these limitations, an alternative firmware, called CoLo (Cobalt Loader) was developed. This is a BOOTROM image that can either be flashed into the chip inside the Cobalt server, or loaded from the existing firmware.

**メモ**

This guide will go through setting up CoLo so that it is loaded by the stock firmware. This is the only truly safe, and recommended way to set up CoLo.

**警告**

If wanted, these can be flashed into the server to totally replace the original firmware -- however, you are entirely on your own in that endeavour. Should anything go wrong, physically remove the BOOTROM and reprogram it with the stock firmware. If this sounds scary -- then DO NOT flash the machine. Gentoo takes no responsibility for whatever happens if this advice is ignored.

Now to install CoLo. Start by emerging the package:

`root #` `emerge --ask sys-boot/colo`

With that installed, take a look inside the /usr/lib/colo/ directory to find two files:

- colo-chain.elf \- the "kernel" for the stock firmware to load.
- colo-rom-image.bin \- a ROM image for flashing into the BOOTROM.

Start by mounting /boot/ and dumping a compressed copy of colo-chain.elf in /boot/ where the system expects it.

`root #` `gzip -9vc /usr/lib/colo/colo-chain.elf > /boot/vmlinux.gz`

### Configuring CoLo

Now, when the system first boots up, it'll load CoLo which will spit up a menu on the back LCD. The first option (and default that is assumed after roughly 5 seconds) is to boot to the hard disk. The system would then attempt to mount the first Linux partition it finds, and run the script default.colo. The syntax is fully documented in the CoLo documentation (have a peek at /usr/share/doc/colo-X.YY/README.shell.gz \-\- where X.YY is the version installed), and is very simple.

**メモ**

Just a tip: when installing kernels, it is recommended to create two kernel images, kernel.gz.working -- a known working kernel, and kernel.gz.new -- a kernel that's just been compiled. It is possible to use symlinks to point to the curent "new" and "working" kernels, or just rename the kernel images.

ファイル **`default.colo`** **An example CoLo configuration**

```
#:CoLo:#
mount hda1
load /kernel.gz.working
execute root=/dev/sda5 ro console=ttyS0,115200

```

**メモ**

CoLo will refuse to load a script that does not begin with the `#:CoLo:#` line. Think of it as the equivalent of saying `#!/bin/sh` in shell scripts.

It is also possible to ask a question, such as which kernel & configuration to boot, with a default timeout. The following configuration does exactly this, asks the user which kernel they wish to use, and executes the chosen image. vmlinux.gz.new and vmlinux.gz.working may be actual kernel images, or just symlinks pointing to the kernel images on that disk. The 50 argument to select specifies that it should proceed with the first option ("Working") after 50/10 seconds.

ファイル **`default.colo`** **Menu-based configuration**

```
#:CoLo:#
lcd "Mounting hda1"
mount hda1
select "Which Kernel?" 50 Working New

goto {menu-option}
var image-name vmlinux.gz.working
goto 3f
@var image-name vmlinux.gz.working
goto 2f
@var image-name vmlinux.gz.new

@lcd "Loading Linux" {image-name}
load /{image-name}
lcd "Booting..."
execute root=/dev/sda5 ro console=ttyS0,115200
boot

```

See the documentation in /usr/share/doc/colo-VERSION for more details.

## Setting up for serial console

Okay, the Linux installation as it stands now, would boot fine, but assumes the user will be logged in at a physical terminal. On Cobalt machines, this is particularly bad -- there's no such thing as a physical terminal.

**メモ**

Those who do have the luxury of a supported video chipset may skip this section if they wish.

First, pull up an editor and hack away at /etc/inittab. Further down in the file, notice the following:

ファイル **`/etc/inittab`** **Snippet from inittab**

```
# SERIAL CONSOLE
#c0:12345:respawn:/sbin/agetty 9600 ttyS0 vt102

# TERMINALS
c1:12345:respawn:/sbin/agetty 38400 tty1 linux
c2:12345:respawn:/sbin/agetty 38400 tty2 linux
c3:12345:respawn:/sbin/agetty 38400 tty3 linux
c4:12345:respawn:/sbin/agetty 38400 tty4 linux
c5:12345:respawn:/sbin/agetty 38400 tty5 linux
c6:12345:respawn:/sbin/agetty 38400 tty6 linux

# What to do at the "Three Finger Salute".
ca:12345:ctrlaltdel:/sbin/shutdown -r now

```

First, uncomment the c0 line. By default, it's set to use a terminal baud rate of 9600 bps. On Cobalt servers, this may be changed to 115200 to match the baud rate decided by the BOOT ROM. The following is how that section looks then. On a headless machine (e.g. Cobalt servers), it is also recommended to comment out the local terminal lines (c1 through to c6) as these have a habit of misbehaving when they can't open /dev/ttyX.

ファイル **`/etc/inittab`** **Example snippet from inittab**

```
# SERIAL CONSOLE
c0:12345:respawn:/sbin/agetty 115200 ttyS0 vt102

# TERMINALS -- These are useless on a headless qube
#c1:12345:respawn:/sbin/agetty 38400 tty1 linux
#c2:12345:respawn:/sbin/agetty 38400 tty2 linux
#c3:12345:respawn:/sbin/agetty 38400 tty3 linux
#c4:12345:respawn:/sbin/agetty 38400 tty4 linux
#c5:12345:respawn:/sbin/agetty 38400 tty5 linux
#c6:12345:respawn:/sbin/agetty 38400 tty6 linux

```

Now, lastly... tell the system, that the local serial port can be trusted as a secure terminal. The file to change at is /etc/securetty. It contains a list of terminals that the system trusts. Simply stick in two more lines, permitting the serial line to be used for root logins.

`root #` `echo 'ttyS0' >> /etc/securetty`

Lately, Linux also calls this /dev/tts/0 \-\- so add this too:

`root #` `echo 'tts/0' >> /etc/securetty`

## Tweaking the SGI PROM

### Setting generic PROM settings

With the bootloader installed, after rebooting (which will occur shortly), go to the System Maintenance Menu and select `Enter` on _Command Monitor (5)_ similar to the initial netbooting of the system.

コード **Menu after boot**

```
1) Start System
2) Install System Software
3) Run Diagnostics
4) Recover System
5) Enter Command Monitor

```

Provide the location of the Volume Header:

`>>` `setenv SystemPartition scsi(0)disk(1)rdisk(0)partition(8)`

Automatically boot Gentoo:

`>>` `setenv AutoLoad Yes`

Set the timezone:

`>>` `setenv TimeZone EST5EDT`

Use the serial console - graphic adapter users should have "g" instead of "d1" (one):

`>>` `setenv console d1`

Set the serial console baud rate. This is optional, 9600 is the default setting, although one may use rates up to 38400 if that is desired:

`>>` `setenv dbaud 9600`

Now, the next settings depend on how the system is booted.

### Settings for direct volume-header booting

**メモ**

This is covered here for completeness. It's recommended that users look into installing arcload instead.

**メモ**

This only works on the Indy, Indigo2 (R4k) and Challenge S.

Set the root device to Gentoo's root partition, such as /dev/sda3:

`>>` `setenv OSLoadPartition <root device>`

To list the available kernels, type "ls".

`>>` `setenv OSLoader <kernel name>
`

`>>` `setenv OSLoadFilename <kernel name>`

Declare the kernel parameters to pass:

`>>` `setenv OSLoadOptions <kernel parameters>`

To try a kernel without messing with kernel parameters, use the boot -f PROM command:

`root #` `boot -f new root=/dev/sda5 ro`

### Settings for arcload

arcload uses the OSLoadFilename option to specify which options to set from arc.cf. The configuration file is essentially a script, with the top-level blocks defining boot images for different systems, and inside that, optional settings. Thus, setting OSLoadFilename=mysys(serial) pulls in the settings for the mysys block, then sets further options overridden in serial.

In the example file above, one system block is defined, ip28 with working, new and debug options available. Next, define the PROM variables:

Select arcload as the bootloader:- sash64 or sashARCS:

`>>` `setenv OSLoader sash64`

Use the "working" kernel image, defined in "ip28" section of arc.cf:

`>>` `setenv OSLoadFilename ip28(working)`

Starting with arcload-0.5, files no longer need to be placed in the volume header -- they may be placed in a partition instead. To tell arcload where to look for its configuration file and kernels, one must set the OSLoadPartition PROM variable. The exact value here will depend on where the disk resides on the SCSI bus. Use the SystemPartition PROM variable as a guide -- only the partition number should need to change.

**メモ**

Partitions are numbered starting at 0, not 1 as is the case in Linux.

To load from the volume header -- use partition 8:

`>>` `setenv OSLoadPartition scsi(0)disk(1)rdisk(0)partition(8)`

Otherwise, specify the partition and filesystem type:

`>>` `setenv OSLoadPartition scsi(0)disk(1)rdisk(0)partition(0)[ext2]`

## システムのリブート

chroot環境を出て、全てのパーティションをアンマウントします。次に、最終かつ真のテストを実行するためのマジカルコマンドrebootを入力しましょう。

`(chroot) livecd #` `exit`

`livecd~#` `cd
`

`livecd~#` `umount -l /mnt/gentoo/dev{/shm,/pts,}
`

`livecd~#` `umount -R /mnt/gentoo
`

`livecd~#` `reboot`

live イメージを取り出すのを忘れないでください。そうしないと新しくインストールされた Gentoo ではなく、live イメージが再度ブート対象になってしまうかもしれません！

リブートして新しい Gentoo 環境に入ることができたら、最終章の [インストールの締めくくり](/wiki/Handbook:MIPS/Installation/Finalizing/ja "Handbook:MIPS/Installation/Finalizing/ja") に進むのがよいでしょう。

[← ツールのインストール](/wiki/Handbook:MIPS/Installation/Tools/ja "Previous part") [Home](/wiki/Handbook:MIPS/ja "Handbook:MIPS/ja") [インストールの締めくくり →](/wiki/Handbook:MIPS/Installation/Finalizing/ja "Next part")