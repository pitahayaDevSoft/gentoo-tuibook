# Bootloader

Other languages:

- [Deutsch](/wiki/Handbook:Alpha/Installation/Bootloader/de "Handbuch:Alpha/Installation/Bootloader (100% translated)")
- [English](/wiki/Handbook:Alpha/Installation/Bootloader "Handbook:Alpha/Installation/Bootloader (100% translated)")
- [español](/wiki/Handbook:Alpha/Installation/Bootloader/es "Manual de Gentoo: Alpha/Instalación/Arranque (100% translated)")
- [français](/wiki/Handbook:Alpha/Installation/Bootloader/fr "Manuel:Alpha/Installation/Système d'amorçage (100% translated)")
- [italiano](/wiki/Handbook:Alpha/Installation/Bootloader/it "Handbook:Alpha/Installation/Bootloader/it (100% translated)")
- [magyar](/wiki/Handbook:Alpha/Installation/Bootloader/hu "Handbook:Alpha/Installation/Bootloader/hu (100% translated)")
- [polski](/wiki/Handbook:Alpha/Installation/Bootloader/pl "Handbook:Alpha/Installation/Bootloader (100% translated)")
- [português do Brasil](/wiki/Handbook:Alpha/Installation/Bootloader/pt-br "Manual:Alpha/Instalação/Gerenciador de boot (100% translated)")
- [čeština](/wiki/Handbook:Alpha/Installation/Bootloader/cs "Handbook:Alpha/Installation/Bootloader/cs (50% translated)")
- [русский](/wiki/Handbook:Alpha/Installation/Bootloader/ru "Handbook:Alpha/Installation/Bootloader (100% translated)")
- [தமிழ்](/wiki/Handbook:Alpha/Installation/Bootloader/ta "கையேடு:Alpha/நிறுவல்/துவக்கஏற்றி (100% translated)")
- [中文（中国大陆）‎](/wiki/Handbook:Alpha/Installation/Bootloader/zh-cn "手册：Alpha/安装/配置系统引导程序Bootloader (100% translated)")
- 日本語
- [한국어](/wiki/Handbook:Alpha/Installation/Bootloader/ko "Handbook:Alpha/Installation/Bootloader/ko (100% translated)")

[Alpha ハンドブック](/wiki/Handbook:Alpha "Handbook:Alpha")[インストール](/wiki/Handbook:Alpha/Full/Installation "Handbook:Alpha/Full/Installation")[インストールについて](/wiki/Handbook:Alpha/Installation/About/ja "Handbook:Alpha/Installation/About/ja")[メディアの選択](/wiki/Handbook:Alpha/Installation/Media/ja "Handbook:Alpha/Installation/Media/ja")[ネットワーク設定](/wiki/Handbook:Alpha/Installation/Networking/ja "Handbook:Alpha/Installation/Networking/ja")[ディスクの準備](/wiki/Handbook:Alpha/Installation/Disks/ja "Handbook:Alpha/Installation/Disks/ja")[stage ファイル](/wiki/Handbook:Alpha/Installation/Stage/ja "Handbook:Alpha/Installation/Stage/ja")[ベースシステムのインストール](/wiki/Handbook:Alpha/Installation/Base/ja "Handbook:Alpha/Installation/Base/ja")[カーネルの設定](/wiki/Handbook:Alpha/Installation/Kernel/ja "Handbook:Alpha/Installation/Kernel/ja")[システムの設定](/wiki/Handbook:Alpha/Installation/System/ja "Handbook:Alpha/Installation/System/ja")[ツールのインストール](/wiki/Handbook:Alpha/Installation/Tools/ja "Handbook:Alpha/Installation/Tools/ja")ブートローダの設定[締めくくり](/wiki/Handbook:Alpha/Installation/Finalizing/ja "Handbook:Alpha/Installation/Finalizing/ja")[Gentoo の操作](/wiki/Handbook:Alpha/Full/Working "Handbook:Alpha/Full/Working")[Portage について](/wiki/Handbook:Alpha/Working/Portage/ja "Handbook:Alpha/Working/Portage/ja")[USE フラグ](/wiki/Handbook:Alpha/Working/USE/ja "Handbook:Alpha/Working/USE/ja")[Portage の機能](/wiki/Handbook:Alpha/Working/Features/ja "Handbook:Alpha/Working/Features/ja")[Init スクリプトシステム](/wiki/Handbook:Alpha/Working/Initscripts/ja "Handbook:Alpha/Working/Initscripts/ja")[環境変数](/wiki/Handbook:Alpha/Working/EnvVar/ja "Handbook:Alpha/Working/EnvVar/ja")[Portage の操作](/wiki/Handbook:Alpha/Full/Portage "Handbook:Alpha/Full/Portage")[ファイルとディレクトリ](/wiki/Handbook:Alpha/Portage/Files/ja "Handbook:Alpha/Portage/Files/ja")[変数](/wiki/Handbook:Alpha/Portage/Variables/ja "Handbook:Alpha/Portage/Variables/ja")[ソフトウェアブランチの併用](/wiki/Handbook:Alpha/Portage/Branches/ja "Handbook:Alpha/Portage/Branches/ja")[追加ツール](/wiki/Handbook:Alpha/Portage/Tools/ja "Handbook:Alpha/Portage/Tools/ja")[カスタムパッケージリポジトリ](/wiki/Handbook:Alpha/Portage/CustomTree/ja "Handbook:Alpha/Portage/CustomTree/ja")[高度な機能](/wiki/Handbook:Alpha/Portage/Advanced/ja "Handbook:Alpha/Portage/Advanced/ja")[OpenRC ネットワーク設定](/wiki/Handbook:Alpha/Full/Networking "Handbook:Alpha/Full/Networking")[はじめに](/wiki/Handbook:Alpha/Networking/Introduction/ja "Handbook:Alpha/Networking/Introduction/ja")[高度な設定](/wiki/Handbook:Alpha/Networking/Advanced/ja "Handbook:Alpha/Networking/Advanced/ja")[モジュール式ネットワーク](/wiki/Handbook:Alpha/Networking/Modular/ja "Handbook:Alpha/Networking/Modular/ja")[無線](/wiki/Handbook:Alpha/Networking/Wireless/ja "Handbook:Alpha/Networking/Wireless/ja")[機能の追加](/wiki/Handbook:Alpha/Networking/Extending/ja "Handbook:Alpha/Networking/Extending/ja")[動的な管理](/wiki/Handbook:Alpha/Networking/Dynamic/ja "Handbook:Alpha/Networking/Dynamic/ja")

## 選択する

カーネルを設定、コンパイルし、必要なシステム設定ファイルを正しく完成させたので、システムが開始した時にカーネルを起動させるプログラムをインストールする時です。そのようなプログラムはブートローダと呼ばれます。

Linux/Alpha用にいくつかのブートローダが存在します。対応しているブートローダのうち、全てではなく、どれか1つを選んでください。 [aBoot](#.E3.83.87.E3.83.95.E3.82.A9.E3.83.AB.E3.83.88.EF.BC.9AaBoot.E3.81.AE.E4.BD.BF.E7.94.A8) と [MILO](#.E4.BB.A3.E6.9B.BF.E6.A1.88.EF.BC.9AMILO.E3.81.AE.E4.BD.BF.E7.94.A8) について文書化してあります。

## デフォルト: aBoot の使用

**メモ**

abootはext2パーティションとext3パーティションのみに対応しています。

まずabootをシステムにインストールしてください

`root #` `emerge --ask sys-boot/aboot`

次のステップはブートディスクを起動可能にすることです。これはシステムの起動時にabootを開始させます。abootブートローダをディスクの先頭に書き込むことで、ブートディスクのブートテーブルを作成します。

`root #` `swriteboot -f3 /dev/sda /boot/bootlx
`

`root #` `abootconf /dev/sda 2
`

**メモ**

もしこの章を通して使用されているパーティション設計とは異なる設計を使用している場合、必要に応じてコマンドを変更する必要があります。適切なマニュアルページ（man 8 swritebootやman 8 abootconf）を読んでください。また、ルートファイルシステムにJFSファイルシステムを使用している場合、カーネルオプションにroを追加して、必ず最初は読み取り専用でマウントしてください。

abootはインストールしましたが、まだaboot用に設定ファイルを書く必要があります。abootはそれぞれの設定に一行のみを必要としますので、これで出来ます：

`root #` `echo '0:2/boot/vmlinux.gz root=/dev/sda3' > /etc/aboot.conf`

もしLinuxカーネルをビルド中に、起動するためにinitramfsもビルドしたのなら、initramfsファイルを参照し、initramfsに実際のルートデバイスの場所を伝えるために設定を変更する必要があります：

`root #` `echo '0:2/boot/vmlinux.gz initrd=/boot/initramfs-genkernel-alpha-6.19.1-gentoo root=/dev/sda3' > /etc/aboot.conf`

加えて、いくつかのSRM変数を設定することでGentooを自動で起動させることが可能です。これらの変数をLinuxから設定してみてください、ただ、SRMコンソール自体から行うほうが簡単かもしれません。

`root #` `cd /proc/srm_environment/named_variables
`

`root #` `echo -n 0 > boot_osflags
`

`root #` `echo -n '' > boot_file
`

`root #` `echo -n 'BOOT' > auto_action
`

`root #` `echo -n 'dkc100' > bootdef_dev`

もちろんdkc100をブートデバイスに置換してください。

将来再びSRMコンソールに入る場合（Gentooインストールを回復させる、変数を扱う、など）、 `Ctrl+C` を打ち、自動読み込み処理を中断してください。

シリアルコンソールを使用してインストールしている場合、aboot.confにシリアルコンソールの起動フラグを含めるのを忘れないでください。さらなる情報については/etc/aboot.conf.exampleを見てください。

abootを設定し、使用準備が完了したので、 [システムのリブート](#.E3.82.B7.E3.82.B9.E3.83.86.E3.83.A0.E3.81.AE.E3.83.AA.E3.83.96.E3.83.BC.E3.83.88) に続いてください。

それでは [システムのリブート](#.E3.82.B7.E3.82.B9.E3.83.86.E3.83.A0.E3.81.AE.E3.83.AA.E3.83.96.E3.83.BC.E3.83.88) に続いてください。

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

リブートして新しい Gentoo 環境に入ることができたら、最終章の [インストールの締めくくり](/wiki/Handbook:Alpha/Installation/Finalizing/ja "Handbook:Alpha/Installation/Finalizing/ja") に進むのがよいでしょう。

[← ツールのインストール](/wiki/Handbook:Alpha/Installation/Tools/ja "Previous part") [Home](/wiki/Handbook:Alpha/ja "Handbook:Alpha/ja") [インストールの締めくくり →](/wiki/Handbook:Alpha/Installation/Finalizing/ja "Next part")