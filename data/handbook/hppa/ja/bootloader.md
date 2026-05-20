# Bootloader

Other languages:

- [Deutsch](/wiki/Handbook:HPPA/Installation/Bootloader/de "Handbuch:HPPA/Installation/Bootloader (100% translated)")
- [English](/wiki/Handbook:HPPA/Installation/Bootloader "Handbook:HPPA/Installation/Bootloader (100% translated)")
- [español](/wiki/Handbook:HPPA/Installation/Bootloader/es "Manual de Gentoo: HPPA/Instalación/Arranque (100% translated)")
- [français](/wiki/Handbook:HPPA/Installation/Bootloader/fr "Handbook:HPPA/Installation/Bootloader/fr (100% translated)")
- [italiano](/wiki/Handbook:HPPA/Installation/Bootloader/it "Handbook:HPPA/Installation/Bootloader/it (100% translated)")
- [magyar](/wiki/Handbook:HPPA/Installation/Bootloader/hu "Handbook:HPPA/Installation/Bootloader/hu (100% translated)")
- [polski](/wiki/Handbook:HPPA/Installation/Bootloader/pl "Handbook:HPPA/Installation/Bootloader (100% translated)")
- [português do Brasil](/wiki/Handbook:HPPA/Installation/Bootloader/pt-br "Manual:HPPA/Instalação/Gerenciador de Boot (100% translated)")
- [русский](/wiki/Handbook:HPPA/Installation/Bootloader/ru "Handbook:HPPA/Installation/Bootloader (100% translated)")
- [தமிழ்](/wiki/Handbook:HPPA/Installation/Bootloader/ta "கையேடு:HPPA/நிறுவல்/துவக்கஏற்றி (100% translated)")
- [中文（中国大陆）‎](/wiki/Handbook:HPPA/Installation/Bootloader/zh-cn "手册：HPPA/安装/配置系统引导程序Bootloader (100% translated)")
- 日本語
- [한국어](/wiki/Handbook:HPPA/Installation/Bootloader/ko "Handbook:HPPA/Installation/Bootloader/ko (100% translated)")

[HPPA ハンドブック](/wiki/Handbook:HPPA "Handbook:HPPA")[インストール](/wiki/Handbook:HPPA/Full/Installation "Handbook:HPPA/Full/Installation")[インストールについて](/wiki/Handbook:HPPA/Installation/About/ja "Handbook:HPPA/Installation/About/ja")[メディアの選択](/wiki/Handbook:HPPA/Installation/Media/ja "Handbook:HPPA/Installation/Media/ja")[ネットワーク設定](/wiki/Handbook:HPPA/Installation/Networking/ja "Handbook:HPPA/Installation/Networking/ja")[ディスクの準備](/wiki/Handbook:HPPA/Installation/Disks/ja "Handbook:HPPA/Installation/Disks/ja")[stage ファイル](/wiki/Handbook:HPPA/Installation/Stage/ja "Handbook:HPPA/Installation/Stage/ja")[ベースシステムのインストール](/wiki/Handbook:HPPA/Installation/Base/ja "Handbook:HPPA/Installation/Base/ja")[カーネルの設定](/wiki/Handbook:HPPA/Installation/Kernel/ja "Handbook:HPPA/Installation/Kernel/ja")[システムの設定](/wiki/Handbook:HPPA/Installation/System/ja "Handbook:HPPA/Installation/System/ja")[ツールのインストール](/wiki/Handbook:HPPA/Installation/Tools/ja "Handbook:HPPA/Installation/Tools/ja")ブートローダの設定[締めくくり](/wiki/Handbook:HPPA/Installation/Finalizing/ja "Handbook:HPPA/Installation/Finalizing/ja")[Gentoo の操作](/wiki/Handbook:HPPA/Full/Working "Handbook:HPPA/Full/Working")[Portage について](/wiki/Handbook:HPPA/Working/Portage/ja "Handbook:HPPA/Working/Portage/ja")[USE フラグ](/wiki/Handbook:HPPA/Working/USE/ja "Handbook:HPPA/Working/USE/ja")[Portage の機能](/wiki/Handbook:HPPA/Working/Features/ja "Handbook:HPPA/Working/Features/ja")[Init スクリプトシステム](/wiki/Handbook:HPPA/Working/Initscripts/ja "Handbook:HPPA/Working/Initscripts/ja")[環境変数](/wiki/Handbook:HPPA/Working/EnvVar/ja "Handbook:HPPA/Working/EnvVar/ja")[Portage の操作](/wiki/Handbook:HPPA/Full/Portage "Handbook:HPPA/Full/Portage")[ファイルとディレクトリ](/wiki/Handbook:HPPA/Portage/Files/ja "Handbook:HPPA/Portage/Files/ja")[変数](/wiki/Handbook:HPPA/Portage/Variables/ja "Handbook:HPPA/Portage/Variables/ja")[ソフトウェアブランチの併用](/wiki/Handbook:HPPA/Portage/Branches/ja "Handbook:HPPA/Portage/Branches/ja")[追加ツール](/wiki/Handbook:HPPA/Portage/Tools/ja "Handbook:HPPA/Portage/Tools/ja")[カスタムパッケージリポジトリ](/wiki/Handbook:HPPA/Portage/CustomTree/ja "Handbook:HPPA/Portage/CustomTree/ja")[高度な機能](/wiki/Handbook:HPPA/Portage/Advanced/ja "Handbook:HPPA/Portage/Advanced/ja")[OpenRC ネットワーク設定](/wiki/Handbook:HPPA/Full/Networking "Handbook:HPPA/Full/Networking")[はじめに](/wiki/Handbook:HPPA/Networking/Introduction/ja "Handbook:HPPA/Networking/Introduction/ja")[高度な設定](/wiki/Handbook:HPPA/Networking/Advanced/ja "Handbook:HPPA/Networking/Advanced/ja")[モジュール式ネットワーク](/wiki/Handbook:HPPA/Networking/Modular/ja "Handbook:HPPA/Networking/Modular/ja")[無線](/wiki/Handbook:HPPA/Networking/Wireless/ja "Handbook:HPPA/Networking/Wireless/ja")[機能の追加](/wiki/Handbook:HPPA/Networking/Extending/ja "Handbook:HPPA/Networking/Extending/ja")[動的な管理](/wiki/Handbook:HPPA/Networking/Dynamic/ja "Handbook:HPPA/Networking/Dynamic/ja")

## PALO のインストール

PA-RISCプラットフォームでは、ブートローダはpaloと呼ばれます。まずこのブートローダをシステムに入れてください：

`root #` `emerge --ask sys-boot/palo`

設定ファイルは/etc/palo.confにあります。以下は設定のサンプルです：

**警告**

This configuration **must** be changed after running palo for the first time! See below for after first setup.

ファイル **`/etc/palo.conf`** **単純なPALO設定例**

```
'"`UNIQ--pre-00000001-QINU`"'

```

最初の行でpaloに、カーネルの場所と使用しなければならない起動パラメータを伝えています。文字列 `2/kernel-6.19.1-gentoo` は、/kernel-6.19.1-gentooと名付けられたカーネルが2番目のパーティション内にあることを意味しています。注意してほしいのが、カーネルへのパスは _ブート_ パーティションからの相対アドレスであり、ルートパーティションからの相対アドレスではありません。

2行目は、どのリカバリカーネルを使用するかを示しています。もしこれが最初のインストールで、リカバリカーネルが（まだ）存在しない場合、この行をコメントアウトしてください。3行目はpaloがどのディスクにあるのかを示しています

To format the disk, palo must be run with certain arguments. This example uses _ext4_ for the first partition:

`root #` `palo --format-as=4 --init-partitioned=/dev/sda`

設定が終了したら、単にpaloコマンドを実行してください：

`root #` `palo`

The configuration must then be updated for post-first-install use:

ファイル **`/etc/palo.conf`** **Simple PALO configuration example**

```
--commandline=2/kernel-6.19.1-gentoo root=/dev/sda4
--recoverykernel=/vmlinux.old
# Don't throw away the old partition, just update the existing one on `palo` runs.
--update-partitioned=/dev/sda
# --format-as has two meanings depending on whether --init-partitioned or --update-partitioned is used. Keep this line.
--format-as=4

```

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

リブートして新しい Gentoo 環境に入ることができたら、最終章の [インストールの締めくくり](/wiki/Handbook:HPPA/Installation/Finalizing/ja "Handbook:HPPA/Installation/Finalizing/ja") に進むのがよいでしょう。

[← ツールのインストール](/wiki/Handbook:HPPA/Installation/Tools/ja "Previous part") [Home](/wiki/Handbook:HPPA/ja "Handbook:HPPA/ja") [インストールの締めくくり →](/wiki/Handbook:HPPA/Installation/Finalizing/ja "Next part")