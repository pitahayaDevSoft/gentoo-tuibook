# Disks

Other languages:

- [Deutsch](/wiki/Handbook:Alpha/Installation/Disks/de "Handbuch:Alpha/Installation/Festplatten (100% translated)")
- [English](/wiki/Handbook:Alpha/Installation/Disks "Handbook:Alpha/Installation/Disks (100% translated)")
- [Türkçe](/wiki/Handbook:Alpha/Installation/Disks/tr "Handbook:Alpha/Installation/Disks/tr (0% translated)")
- [español](/wiki/Handbook:Alpha/Installation/Disks/es "Manual de Gentoo: Alpha/Instalación/Discos (100% translated)")
- [français](/wiki/Handbook:Alpha/Installation/Disks/fr "Manuel:Alpha/Installation/Disques (100% translated)")
- [italiano](/wiki/Handbook:Alpha/Installation/Disks/it "Handbook:Alpha/Installation/Disks/it (100% translated)")
- [magyar](/wiki/Handbook:Alpha/Installation/Disks/hu "Handbook:Alpha/Installation/Disks/hu (100% translated)")
- [polski](/wiki/Handbook:Alpha/Installation/Disks/pl "Handbook:Alpha/Installation/Disks (100% translated)")
- [português do Brasil](/wiki/Handbook:Alpha/Installation/Disks/pt-br "Manual:Alpha/Instalação/Discos (100% translated)")
- [čeština](/wiki/Handbook:Alpha/Installation/Disks/cs "Handbook:Alpha/Installation/Disks/cs (50% translated)")
- [русский](/wiki/Handbook:Alpha/Installation/Disks/ru "Handbook:Alpha/Installation/Disks (100% translated)")
- [தமிழ்](/wiki/Handbook:Alpha/Installation/Disks/ta "கையேடு:Alpha/நிறுவல்/வட்டுக்கள் (100% translated)")
- [中文（中国大陆）‎](/wiki/Handbook:Alpha/Installation/Disks/zh-cn "手册：Alpha/安装/准备磁盘 (100% translated)")
- 日本語
- [한국어](/wiki/Handbook:Alpha/Installation/Disks/ko "Handbook:Alpha/Installation/Disks/ko (100% translated)")

[Alpha ハンドブック](/wiki/Handbook:Alpha "Handbook:Alpha")[インストール](/wiki/Handbook:Alpha/Full/Installation "Handbook:Alpha/Full/Installation")[インストールについて](/wiki/Handbook:Alpha/Installation/About/ja "Handbook:Alpha/Installation/About/ja")[メディアの選択](/wiki/Handbook:Alpha/Installation/Media/ja "Handbook:Alpha/Installation/Media/ja")[ネットワーク設定](/wiki/Handbook:Alpha/Installation/Networking/ja "Handbook:Alpha/Installation/Networking/ja")ディスクの準備[stage ファイル](/wiki/Handbook:Alpha/Installation/Stage/ja "Handbook:Alpha/Installation/Stage/ja")[ベースシステムのインストール](/wiki/Handbook:Alpha/Installation/Base/ja "Handbook:Alpha/Installation/Base/ja")[カーネルの設定](/wiki/Handbook:Alpha/Installation/Kernel/ja "Handbook:Alpha/Installation/Kernel/ja")[システムの設定](/wiki/Handbook:Alpha/Installation/System/ja "Handbook:Alpha/Installation/System/ja")[ツールのインストール](/wiki/Handbook:Alpha/Installation/Tools/ja "Handbook:Alpha/Installation/Tools/ja")[ブートローダの設定](/wiki/Handbook:Alpha/Installation/Bootloader/ja "Handbook:Alpha/Installation/Bootloader/ja")[締めくくり](/wiki/Handbook:Alpha/Installation/Finalizing/ja "Handbook:Alpha/Installation/Finalizing/ja")[Gentoo の操作](/wiki/Handbook:Alpha/Full/Working "Handbook:Alpha/Full/Working")[Portage について](/wiki/Handbook:Alpha/Working/Portage/ja "Handbook:Alpha/Working/Portage/ja")[USE フラグ](/wiki/Handbook:Alpha/Working/USE/ja "Handbook:Alpha/Working/USE/ja")[Portage の機能](/wiki/Handbook:Alpha/Working/Features/ja "Handbook:Alpha/Working/Features/ja")[Init スクリプトシステム](/wiki/Handbook:Alpha/Working/Initscripts/ja "Handbook:Alpha/Working/Initscripts/ja")[環境変数](/wiki/Handbook:Alpha/Working/EnvVar/ja "Handbook:Alpha/Working/EnvVar/ja")[Portage の操作](/wiki/Handbook:Alpha/Full/Portage "Handbook:Alpha/Full/Portage")[ファイルとディレクトリ](/wiki/Handbook:Alpha/Portage/Files/ja "Handbook:Alpha/Portage/Files/ja")[変数](/wiki/Handbook:Alpha/Portage/Variables/ja "Handbook:Alpha/Portage/Variables/ja")[ソフトウェアブランチの併用](/wiki/Handbook:Alpha/Portage/Branches/ja "Handbook:Alpha/Portage/Branches/ja")[追加ツール](/wiki/Handbook:Alpha/Portage/Tools/ja "Handbook:Alpha/Portage/Tools/ja")[カスタムパッケージリポジトリ](/wiki/Handbook:Alpha/Portage/CustomTree/ja "Handbook:Alpha/Portage/CustomTree/ja")[高度な機能](/wiki/Handbook:Alpha/Portage/Advanced/ja "Handbook:Alpha/Portage/Advanced/ja")[OpenRC ネットワーク設定](/wiki/Handbook:Alpha/Full/Networking "Handbook:Alpha/Full/Networking")[はじめに](/wiki/Handbook:Alpha/Networking/Introduction/ja "Handbook:Alpha/Networking/Introduction/ja")[高度な設定](/wiki/Handbook:Alpha/Networking/Advanced/ja "Handbook:Alpha/Networking/Advanced/ja")[モジュール式ネットワーク](/wiki/Handbook:Alpha/Networking/Modular/ja "Handbook:Alpha/Networking/Modular/ja")[無線](/wiki/Handbook:Alpha/Networking/Wireless/ja "Handbook:Alpha/Networking/Wireless/ja")[機能の追加](/wiki/Handbook:Alpha/Networking/Extending/ja "Handbook:Alpha/Networking/Extending/ja")[動的な管理](/wiki/Handbook:Alpha/Networking/Dynamic/ja "Handbook:Alpha/Networking/Dynamic/ja")

## Contents

- [1ブロックデバイスの概要](#.E3.83.96.E3.83.AD.E3.83.83.E3.82.AF.E3.83.87.E3.83.90.E3.82.A4.E3.82.B9.E3.81.AE.E6.A6.82.E8.A6.81)
  - [1.1ブロックデバイス](#.E3.83.96.E3.83.AD.E3.83.83.E3.82.AF.E3.83.87.E3.83.90.E3.82.A4.E3.82.B9)
  - [1.2スライス](#.E3.82.B9.E3.83.A9.E3.82.A4.E3.82.B9)
- [2パーティション構成の設計](#.E3.83.91.E3.83.BC.E3.83.86.E3.82.A3.E3.82.B7.E3.83.A7.E3.83.B3.E6.A7.8B.E6.88.90.E3.81.AE.E8.A8.AD.E8.A8.88)
  - [2.1パーティション数とサイズ](#.E3.83.91.E3.83.BC.E3.83.86.E3.82.A3.E3.82.B7.E3.83.A7.E3.83.B3.E6.95.B0.E3.81.A8.E3.82.B5.E3.82.A4.E3.82.BA)
  - [2.2スワップ領域について](#.E3.82.B9.E3.83.AF.E3.83.83.E3.83.97.E9.A0.98.E5.9F.9F.E3.81.AB.E3.81.A4.E3.81.84.E3.81.A6)
- [3ディスクをパーティションに分けるのに fdisk を使う (SRM のみ)](#.E3.83.87.E3.82.A3.E3.82.B9.E3.82.AF.E3.82.92.E3.83.91.E3.83.BC.E3.83.86.E3.82.A3.E3.82.B7.E3.83.A7.E3.83.B3.E3.81.AB.E5.88.86.E3.81.91.E3.82.8B.E3.81.AE.E3.81.AB_fdisk_.E3.82.92.E4.BD.BF.E3.81.86_.28SRM_.E3.81.AE.E3.81.BF.29)
  - [3.1利用可能なディスクの特定](#.E5.88.A9.E7.94.A8.E5.8F.AF.E8.83.BD.E3.81.AA.E3.83.87.E3.82.A3.E3.82.B9.E3.82.AF.E3.81.AE.E7.89.B9.E5.AE.9A)
  - [3.2BSD ディスクラベルを作成する](#BSD_.E3.83.87.E3.82.A3.E3.82.B9.E3.82.AF.E3.83.A9.E3.83.99.E3.83.AB.E3.82.92.E4.BD.9C.E6.88.90.E3.81.99.E3.82.8B)
  - [3.3すべてのスライスの削除](#.E3.81.99.E3.81.B9.E3.81.A6.E3.81.AE.E3.82.B9.E3.83.A9.E3.82.A4.E3.82.B9.E3.81.AE.E5.89.8A.E9.99.A4)
  - [3.4スワップスライスの作成](#.E3.82.B9.E3.83.AF.E3.83.83.E3.83.97.E3.82.B9.E3.83.A9.E3.82.A4.E3.82.B9.E3.81.AE.E4.BD.9C.E6.88.90)
  - [3.5ブートスライスを作成する](#.E3.83.96.E3.83.BC.E3.83.88.E3.82.B9.E3.83.A9.E3.82.A4.E3.82.B9.E3.82.92.E4.BD.9C.E6.88.90.E3.81.99.E3.82.8B)
  - [3.6ルートスライスの作成](#.E3.83.AB.E3.83.BC.E3.83.88.E3.82.B9.E3.83.A9.E3.82.A4.E3.82.B9.E3.81.AE.E4.BD.9C.E6.88.90)
  - [3.7スライスレイアウトを保存し終了する](#.E3.82.B9.E3.83.A9.E3.82.A4.E3.82.B9.E3.83.AC.E3.82.A4.E3.82.A2.E3.82.A6.E3.83.88.E3.82.92.E4.BF.9D.E5.AD.98.E3.81.97.E7.B5.82.E4.BA.86.E3.81.99.E3.82.8B)
- [4ディスクをパーティションに分けるのに fdisk を使う (ARC/AlphaBIOS のみ)](#.E3.83.87.E3.82.A3.E3.82.B9.E3.82.AF.E3.82.92.E3.83.91.E3.83.BC.E3.83.86.E3.82.A3.E3.82.B7.E3.83.A7.E3.83.B3.E3.81.AB.E5.88.86.E3.81.91.E3.82.8B.E3.81.AE.E3.81.AB_fdisk_.E3.82.92.E4.BD.BF.E3.81.86_.28ARC.2FAlphaBIOS_.E3.81.AE.E3.81.BF.29)
  - [4.1利用可能なディスクの特定](#.E5.88.A9.E7.94.A8.E5.8F.AF.E8.83.BD.E3.81.AA.E3.83.87.E3.82.A3.E3.82.B9.E3.82.AF.E3.81.AE.E7.89.B9.E5.AE.9A_2)
  - [4.2すべてのパーティションの削除](#.E3.81.99.E3.81.B9.E3.81.A6.E3.81.AE.E3.83.91.E3.83.BC.E3.83.86.E3.82.A3.E3.82.B7.E3.83.A7.E3.83.B3.E3.81.AE.E5.89.8A.E9.99.A4)
  - [4.3ブートパーティションの作成](#.E3.83.96.E3.83.BC.E3.83.88.E3.83.91.E3.83.BC.E3.83.86.E3.82.A3.E3.82.B7.E3.83.A7.E3.83.B3.E3.81.AE.E4.BD.9C.E6.88.90)
  - [4.4スワップパーティションの作成](#.E3.82.B9.E3.83.AF.E3.83.83.E3.83.97.E3.83.91.E3.83.BC.E3.83.86.E3.82.A3.E3.82.B7.E3.83.A7.E3.83.B3.E3.81.AE.E4.BD.9C.E6.88.90)
  - [4.5ルートパーティションの作成](#.E3.83.AB.E3.83.BC.E3.83.88.E3.83.91.E3.83.BC.E3.83.86.E3.82.A3.E3.82.B7.E3.83.A7.E3.83.B3.E3.81.AE.E4.BD.9C.E6.88.90)
  - [4.6パーティションレイアウトを保存し終了する](#.E3.83.91.E3.83.BC.E3.83.86.E3.82.A3.E3.82.B7.E3.83.A7.E3.83.B3.E3.83.AC.E3.82.A4.E3.82.A2.E3.82.A6.E3.83.88.E3.82.92.E4.BF.9D.E5.AD.98.E3.81.97.E7.B5.82.E4.BA.86.E3.81.99.E3.82.8B)
- [5ファイルシステムを作成する](#.E3.83.95.E3.82.A1.E3.82.A4.E3.83.AB.E3.82.B7.E3.82.B9.E3.83.86.E3.83.A0.E3.82.92.E4.BD.9C.E6.88.90.E3.81.99.E3.82.8B)
  - [5.1はじめに](#.E3.81.AF.E3.81.98.E3.82.81.E3.81.AB)
  - [5.2ファイルシステム](#.E3.83.95.E3.82.A1.E3.82.A4.E3.83.AB.E3.82.B7.E3.82.B9.E3.83.86.E3.83.A0)
  - [5.3パーティションにファイルシステムを適用する](#.E3.83.91.E3.83.BC.E3.83.86.E3.82.A3.E3.82.B7.E3.83.A7.E3.83.B3.E3.81.AB.E3.83.95.E3.82.A1.E3.82.A4.E3.83.AB.E3.82.B7.E3.82.B9.E3.83.86.E3.83.A0.E3.82.92.E9.81.A9.E7.94.A8.E3.81.99.E3.82.8B)
    - [5.3.1レガシー BIOS ブートパーティションのファイルシステム](#.E3.83.AC.E3.82.AC.E3.82.B7.E3.83.BC_BIOS_.E3.83.96.E3.83.BC.E3.83.88.E3.83.91.E3.83.BC.E3.83.86.E3.82.A3.E3.82.B7.E3.83.A7.E3.83.B3.E3.81.AE.E3.83.95.E3.82.A1.E3.82.A4.E3.83.AB.E3.82.B7.E3.82.B9.E3.83.86.E3.83.A0)
    - [5.3.2小さな ext4 パーティション](#.E5.B0.8F.E3.81.95.E3.81.AA_ext4_.E3.83.91.E3.83.BC.E3.83.86.E3.82.A3.E3.82.B7.E3.83.A7.E3.83.B3)
  - [5.4スワップパーティションを有効にする](#.E3.82.B9.E3.83.AF.E3.83.83.E3.83.97.E3.83.91.E3.83.BC.E3.83.86.E3.82.A3.E3.82.B7.E3.83.A7.E3.83.B3.E3.82.92.E6.9C.89.E5.8A.B9.E3.81.AB.E3.81.99.E3.82.8B)
- [6ルートパーティションのマウント](#.E3.83.AB.E3.83.BC.E3.83.88.E3.83.91.E3.83.BC.E3.83.86.E3.82.A3.E3.82.B7.E3.83.A7.E3.83.B3.E3.81.AE.E3.83.9E.E3.82.A6.E3.83.B3.E3.83.88)

## ブロックデバイスの概要

### ブロックデバイス

Gentoo Linuxの、そしてLinux一般の、ブロックデバイス、パーティション、Linuxファイルシステムを含めた、ディスクやファイルシステム中心の考え方について詳しく見てみましょう。ディスクの入出力とファイルシステムについて理解することで、インストールのためのパーティションとファイルシステムを構築できるようになります。

まずはブロックデバイスについて見ていきます。SCSIドライブやシリアルATAドライブは両方とも/dev/sdaや/dev/sdb、/dev/sdcなどのようなデバイスハンドルとしてラベル付されます。更にモダンなマシンでは、PCI ExpressベースのNVMeソリッドステートディスクは、/dev/nvme0n1、/dev/nvme0n2などのようなデバイスハンドルを持ちます。

下の表は、各種のブロックデバイスがシステム上のどこにあるかを判断するのに役立つでしょう:

デバイスの種類デフォルトのデバイスハンドル編集者メモと、考慮すべき点
IDE、SATA、SAS、SCSI、または USB フラッシュメモリ/dev/sda2007 年頃から現在までに製造されたハードウェアで見られます。このデバイスハンドルはおそらく Linux 上でもっともよく使用されているものでしょう。この種のデバイスは [SATA バス](https://en.wikipedia.org/wiki/Serial_ATA "wikipedia:Serial ATA")、 [SCSI](https://en.wikipedia.org/wiki/SCSI "wikipedia:SCSI")、 [USB](https://en.wikipedia.org/wiki/USB "wikipedia:USB") バスを介してブロックストレージとして接続されます。例えば、最初の SATA デバイス上の最初のパーティションは /dev/sda1 という名前になります。
NVM Express (NVMe)/dev/nvme0n1ソリッドステートテクノロジとして最新の [NVMe](https://en.wikipedia.org/wiki/NVM_Express "wikipedia:NVM Express") ドライブは PCI Express バスに接続され、一般市場でもっとも高速な転送速度を持っています。2014 年頃以降のシステムは NVMe ハードウェアのサポートを備えているかもしれません。最初の NVMe デバイスの最初のパーティションは /dev/nvme0n1p1 という名前になります。
MMC、eMMC、および SD カード/dev/mmcblk0[embedded MMC](https://en.wikipedia.org/wiki/MultiMediaCard#eMMC "wikipedia:MultiMediaCard") デバイス、SD カード、そして他の種類のメモリーカードはデータ用のストレージとして有用です。つまり、多くのシステムはこれらの種類のデバイスからのブートを許可していないかもしれません。これらのデバイスに Linux をインストールして常用するのは **おすすめできません**。それらの典型的な設計意図である、ファイルの交換用に使うものと考えてください。この種のストレージは短期的なファイルバックアップまたはスナップショットとして使用すると便利かもしれません。
VirtIO/dev/vda[QEMU](/wiki/QEMU "QEMU") 仮想エミュレータ内でのみ見られる [仮想](/wiki/Virtualization "Virtualization") デバイス。 `virtio_blk` ドライバは仮想マシンに対してホストのドライブ空間へのアクセスを提供します。とはいえ、仮想マシン内で Gentoo を試すための優れた方法です。

上のブロックデバイスは、ディスクへの抽象的なインターフェースを表しています。ユーザープログラムはこれらのブロックデバイスを用いて、デバイスが SATA、SCSI、もしくは他のものであるかどうかを心配することなしにディスクと通信することができます。プログラムは容易にディスク上の記憶領域を、ランダムアクセスできる 4096 バイト (4K) ごとの連続領域としてアドレッシングできます。

### スライス

理論的にはディスク全体をLinuxシステムに使用することが出来ますが、実際にはほとんど行われません。その代わりに、ディスクブロックデバイス全体を小さく、より扱いやすいブロックデバイスに分割します。Alphaシステムでは、これは _スライス_ と呼ばれます。

**メモ**

これより先のセクションでは、インストールの説明でARC/AlphaBIOSのセットアップ用のパーティショニングを使用します。個人の好みに合わせて調整してください！

## パーティション構成の設計

### パーティション数とサイズ

ディスクのパーティションレイアウトの設計は、システムに対する要求と、デバイスに適用されるファイルシステムに大きく依存します。多数のユーザがいる場合、セキュリティを向上し、バックアップの作成とその他のメンテナンスを容易にするために、/home を分離されたパーティションに配置することが推奨されます。もし メールサーバとして動作する場合は、/var を分離されたパーティションとし、すべてのメールを /var ディレクトリに保存すべきでしょう。ゲームサーバでは、ほとんどのゲームサーバソフトウェアは /opt にインストールされるので、/opt を分離されたパーティションとすることができます。これらが推奨される理由は最初の /home ディレクトリと同様で、セキュリティ、バックアップ、そしてメンテナンスです。

Gentoo では多くの場合、/usr と /var は相対的に大きい容量を確保すべきです。/usr にはシステム上で利用可能なアプリケーションの大部分と、Linux カーネルソース (/usr/src 配下) が配置されます。デフォルトでは、/var には Gentoo ebuild リポジトリが (/var/db/repos/gentoo 配下に) 配置され、ファイルシステム依存ではあるものの通常 650 MiB ほどのディスク容量を消費します。この推定容量には /var/cache/distfiles と /var/cache/binpkgs ディレクトリは _含まれていません_。これらはそれぞれ、ソースファイルとバイナリパッケージ (使用している場合) を格納するディレクトリで、システムに追加すればするほど大きくなっていきます。

適切なパーティションの数とサイズは、システムを取り巻く環境と、トレードオフを考慮することで大きく変わります。パーティションやボリュームを分離することには下記の利点があります:

- それぞれのパーティションまたはボリュームに対して、最も性能が高いファイルシステムを選択できます
- ゾンビプロセスがパーティションまたはボリュームに継続的に書き込みをした場合でも、システム全体の空き領域を使い切ることはありません
- 必要ならば、複数のチェックを並行して実行することで、ファイルシステムチェックの時間を短縮できます (複数のパーティションよりも複数のディスクの方が効果を実感できます)
- リードのみ、 `nosuid`(setuidビット無効)、 `noexec`(実行ビット無効)等のマウントオプションによって、セキュリティが向上します

しかし、複数パーティションにはデメリットもあります:

- もし適切に設定されていないと、あるパーティションが空き領域をたくさん持ち、別のパーティションにはまったく空き領域がなくなるといったことが起こり得ます。
- /usr/ を独立したパーティションにすると、他のブートスクリプトが動作する前にパーティションをマウントするために、initramfs を使ってブートする必要があるかもしれません。initramfs の生成と保守はこのハンドブックのスコープの範囲外ですので、 **慣れていない方が /usr を独立したパーティションとすることは推奨しません。**
- SCSI や SATA では仕様上の制約により、GPT ラベルを使用しない限りは 15 個までしかパーティションを作れません。

**メモ**

サービスおよび init システムとして systemd を使うつもりのインストールでは、/usr ディレクトリはルートファイルシステムの一部とするか、または initramfs によりマウントされるようにして、ブート時に利用できるようにしなくてはなりません。

### スワップ領域について

スワップ領域サイズの推奨値
RAM サイズサスペンド対応時ハイバネーション対応時
2 GB 以下4GB4GB
2 から 8 GBRAM 容量2 \* RAM
8 から 64 GB最小 8 GB、最大 16 GB1.5 \* RAM
64 GB 以上最小 8 GBハイバネーションは推奨されません! 非常に大きな容量のメモリを持つシステムでは、ハイバネーションは推奨 _されません_。可能ではありますが、正しくハイバネーションするためには、メモリの内容全体をディスクに書き込まなくてはなりません。数十 GB (またはそれ以上!) のデータをディスクに書き出すのには、回転式ディスクを使用する場合は特に、非常に長い時間がかかることがあります。この状況ではサスペンドするのが最善です。

スワップ領域のサイズについて完璧な値というものはありません。スワップ領域の目的は、メインメモリ（RAM）が逼迫した際、カーネルにディスク領域を提供するためにあります。スワップ領域があれば、カーネルは最近最も使われていないメモリページをディスクに書き出し（スワップもしくはページアウト）、現在のタスクのために RAM 上に置かれたメモリを開放します。もちろん、もしディスクにスワップされたページが急に必要になった場合は、これらのページはメモリに戻す（ページイン）必要があります。これには、RAM から読み込むより相当長い時間がかかります（メインメモリと比較してディスクはとても遅いためです）。

システムがメモリを大量に消費するアプリケーションを実行しないとき、またシステムが多くの RAM を持っているときは、それほど大きいスワップ領域は必要ではありません。しかし、ハイバネーションの際に、スワップ領域は _メモリの内容すべて_ を保存するために使われる（サーバシステムよりも、デスクトップやラップトップシステムでよくあることです）ことに留意してください。システムにハイバネーションのサポートが必要な場合は、メモリの全体量以上のサイズのスワップ領域が必要です。

RAM 容量が 4GB より少ない場合の一般的なルールとして、スワップ領域のサイズは内部メモリ (RAM) の 2 倍であることが推奨されます。複数のハードディスクを備えるシステムでは、並列して読み込み/書き込み操作が行えるように、それぞれのディスクに 1 つずつスワップパーティションを作成するのが賢い方法です。スワップ空間内のデータにアクセスしなくてはならないときに、ディスクがより高速にスワップできるほど、システムもより高速に動作するでしょう。回転式ディスクとソリッドステートディスクを比較すると、ソリッドステートハードウェア上にスワップを置いたほうが高いパフォーマンスが発揮できます。

スワップ _パーティション_ の代わりに、スワップ _ファイル_ を使用することができることも特筆に値します。これは主にディスク容量が非常に限られたシステムで役に立つものです。

## ディスクをパーティションに分けるのに fdisk を使う (SRM のみ)

以下のパートではSRM用に、この例のようなスライスレイアウトを作成する方法を説明します：

スライス
説明
/dev/sda1スワップスライス
/dev/sda2ルートスライス
/dev/sda3フルディスク（必須）

スライスレイアウトはお好みで変更してください。

### 利用可能なディスクの特定

どのディスクがシステムで動作しているかを確認するためには、以下のコマンドを使用してください：

IDE ディスクの場合:

`root #` `dmesg | grep 'drive$'`

SCSI ディスクの場合:

`root #` `dmesg | grep 'scsi'`

出力では、どのディスクが検出されたか、そしてディスクに対応する/dev/の項目が表示されます。以下のパートではディスクは/dev/sdaにあるSCSIディスクであると仮定します。

### BSD ディスクラベルを作成する

If the hard drive is completely blank, then first create a BSD disklabel. On alpha, this can't be done using fdisk, so using parted is the way forward here.

Now fire up parted:

`root #` `parted /dev/sda `

```
Using /dev/sda
Welcome to GNU Parted! Type 'help' to view a list of commands.

```

`(parted)` `mklabel bsd`

```
Warning: The existing disk label on /dev/sda will be destroyed and all data on this disk will be lost. Do you want to continue?
Yes/No? yes
(parted) quit
Information: You may need to update /etc/fstab.

```

Now that we have a bsd disklabel on our drive, continue creating slices. This can be done using parted or, as in the examples below, using fdisk:

`root #` `fdisk /dev/sda`

### すべてのスライスの削除

もしハードディスクが完全に空ならば、まずBSDディスクラベルを作成します。

`Command (m for help):` `b`

```
/dev/sda contains no disklabel.
Do you want to create a disklabel? (y/n) y
A bunch of drive-specific info will show here
3 partitions:
#       start       end      size     fstype   [fsize bsize   cpg]
  c:        1      5290*     5289*    unused        0     0

```

まず'c'-スライス（BSDディスクラベルを使用するのに必要）以外のすべてのスライスを削除することから始めます。以下で、どのようにスライスを削除するかを示しています（この例では'a'を使用します）。この作業を他のすべてのディスクにも行ってください（繰り返しますが、'c'-スライス以外です）。

`p` を使用して、存在するすべてのスライスを確認してください。 `d` はスライスを削除するのに使います。

`BSD disklabel command (m for help):` `p`

```
8 partitions:
#       start       end      size     fstype   [fsize bsize   cpg]
  a:        1       235*      234*    4.2BSD     1024  8192    16
  b:      235*      469*      234*      swap
  c:        1      5290*     5289*    unused        0     0
  d:      469*     2076*     1607*    unused        0     0
  e:     2076*     3683*     1607*    unused        0     0
  f:     3683*     5290*     1607*    unused        0     0
  g:      469*     1749*     1280     4.2BSD     1024  8192    16
  h:     1749*     5290*     3541*    unused        0     0

```

`BSD disklabel command (m for help):` `d`

```
Partition (a-h): a

```

すべてのスライスに対してこの作業を行ったあと、一覧は以下のようになるはずです：

`BSD disklabel command (m for help):` `p`

```
3 partitions:
#       start       end      size     fstype   [fsize bsize   cpg]
  c:        1      5290*     5289*    unused        0     0

```

### スワップスライスの作成

Alphaベースのシステムでは独立したブートスライスを作成する必要はありません。しかしながら、最初のシリンダは _aboot_ イメージが置かれるため、使用することが出来ません。

3番目のシリンダより、合計サイズ1GBのスワップスライスを作成します。 `n` を使用して新しいスライスを作成してください。スライスを作成したら、スライスのタイプを `1`（数字の一）、すなわちスワップに変更します。

`BSD disklabel command (m for help):` `n`

```
Partition (a-p): a
First cylinder (1-5290, default 1): 3
Last cylinder or +size or +sizeM or +sizeK (3-5290, default 5290): +1024M

```

`BSD disklabel command (m for help):` `t`

```
Partition (a-c): a
Hex code (type L to list codes): 1

```

これらの段階を終えたら、レイアウトは以下のような感じになるはずです：

`BSD disklabel command (m for help):` `p`

```
3 partitions:
#       start       end      size     fstype   [fsize bsize   cpg]
  a:        3      1003      1001       swap
  c:        1      5290*     5289*    unused        0     0

```

### ブートスライスを作成する

Create a boot file system containing the kernels and boot loader configuration file (aboot.conf). aboot supports ext2 and ext3 filesystems only. Create the boot slice starting from the first cylinder after the swap slice. Use the `p` command to view where the swap slice ends. In our example, this is at 1003, making the boot slice start at 1004.

`BSD disklabel command (m for help):` `n`

```
Partition (a-p): b
First cylinder (1-5290, default 1): 1004
Last cylinder or +size or +sizeM or +sizeK (3-5290, default 5290): +1024M

```

`BSD disklabel command (m for help):` `t`

```
Partition (a-c): b
Hex code (type L to list codes): 08

```

After these steps a layout similar to the following should be shown:

`BSD disklabel command (m for help):` `p`

```
3 partitions:
#       start       end      size     fstype   [fsize bsize   cpg]
  a:        3      1003      1001       swap
  b:     1004      2005      1001       ext2
  c:        1      5290*     5289*    unused        0     0

```

### ルートスライスの作成

それではルートスライスを、ブートスライスの最後のシリンダの次より作成します。 `p` コマンドを使用して、ブートスライスの終了位置を確認してください。この例では、終了位置は 2005 ですので、ルートスライスを 2006 より作成します。

現在fdiskにはバグがあり、fdiskが、利用可能なシリンダ数が実際のシリンダ数よりも1多いと見なしている問題があります。言い換えると、最後のシリンダを指定する時、シリンダ数を1つ減らしてください（この例では5290）。

スライスを作成したら、タイプを8、つまりext2に変更します。

`BSD disklabel command (m for help):` `n`

```
Partition (a-p): b
First cylinder (1-5290, default 1): 2006
Last cylinder or +size or +sizeM or +sizeK (1004-5290, default 5290): 5289

```

`BSD disklabel command (m for help):` `t`

```
Partition (a-c): b
Hex code (type L to list codes): 8

```

結果的にスライスレイアウトはこのようになるはずです：

`BSD disklabel command (m for help):` `p`

```
3 partitions:
#       start       end      size     fstype   [fsize bsize   cpg]
  a:        3      1003      1001       swap
  b:     1004      2005      4286       ext2
  c:        1      5290*     5289*    unused        0     0
  d:     2006      5289      3283       ext2

```

### スライスレイアウトを保存し終了する

`w` を打ちfdiskを終了してください。この操作はスライスレイアウトの保存もします。

`Command (m for help):` `w`

## ディスクをパーティションに分けるのに fdisk を使う (ARC/AlphaBIOS のみ)

以下のパートではARC/AlphaBIOS用に、この例のようなパーティションレイアウトを作成する方法を説明します：

パーティション
説明
/dev/sda1ブートパーティション
/dev/sda2スワップパーティション
/dev/sda3ルートパーティション

パーティションレイアウトはお好みで変更してください。

### 利用可能なディスクの特定

どのディスクが動作しているかを確認するためには、以下のコマンドを使用してください：

IDE ディスクの場合:

`root #` `dmesg | grep 'drive$'`

SCSI ディスクの場合:

`root #` `dmesg | grep 'scsi'`

この出力から、どのディスクが検出されたか、そしてディスクに対応する/dev/の項目がどれかを確認するのは簡単なはずです。以下のパートではディスクは/dev/sdaにあるSCSIディスクであると仮定します。

そして fdisk を実行:

`root #` `fdisk /dev/sda`

### すべてのパーティションの削除

もしハードディスクが完全に空ならば、まずDOSディスクラベルを作成します。

`Command (m for help):` `o`

```
Building a new DOS disklabel.

```

まずすべてのパーティションを削除することから始めます。以下で、どのようにパーティションを削除するかを示しています（この例では'1'を使用します）。この作業を他のすべてのパーティションにも行ってください。

`p` を使用して、存在するすべてのパーティションを確認してください。 `d` はパーティションを削除するのに使います。

`command (m for help):` `p`

```
Disk /dev/sda: 9150 MB, 9150996480 bytes
64 heads, 32 sectors/track, 8727 cylinders
Units = cylinders of 2048 * 512 = 1048576 bytes

   Device Boot      Start         End      Blocks   Id  System
/dev/sda1               1         478      489456   83  Linux
/dev/sda2             479        8727     8446976    5  Extended
/dev/sda5             479        1433      977904   83  Linux Swap
/dev/sda6            1434        8727     7469040   83  Linux

```

`command (m for help):` `d`

```
Partition number (1-6): 1

```

### ブートパーティションの作成

MILOを使用して起動しているAlphaシステムでは、小さなvfatのブートパーティションを作成しなければなりません。

`Command (m for help):` `n`

```
Command action
  e   extended
  p   primary partition (1-4)
p
Partition number (1-4): 1
First cylinder (1-8727, default 1): 1
Last cylinder or +size or +sizeM or +sizeK (1-8727, default 8727): +16M

```

`Command (m for help):` `t`

```
Selected partition 1
Hex code (type L to list codes): 6
Changed system type of partition 1 to 6 (FAT16)

```

### スワップパーティションの作成

合計サイズ1GBのスワップパーティションを作成します。 `n` を使用して新しいパーティションを作成してください。

`Command (m for help):` `n`

```
Command action
  e   extended
  p   primary partition (1-4)
p
Partition number (1-4): 2
First cylinder (17-8727, default 17): 17
Last cylinder or +size or +sizeM or +sizeK (17-8727, default 8727): +1000M

```

`Command (m for help):` `t`

```
Partition number (1-4): 2
Hex code (type L to list codes): 82
Changed system type of partition 2 to 82 (Linux swap)

```

これらの段階を終えたら、レイアウトは以下のような感じになるはずです：

`Command (m for help):` `p`

```
Disk /dev/sda: 9150 MB, 9150996480 bytes
64 heads, 32 sectors/track, 8727 cylinders
Units = cylinders of 2048 * 512 = 1048576 bytes

   Device Boot      Start         End      Blocks   Id  System
/dev/sda1               1          16       16368    6  FAT16
/dev/sda2              17         971      977920   82  Linux swap

```

### ルートパーティションの作成

それではルートパーティションを作成します。繰り返しますが、 `n` コマンドを使用してください。

`Command (m for help):` `n`

```
Command action
  e   extended
  p   primary partition (1-4)
p
Partition number (1-4): 3
First cylinder (972-8727, default 972): 972
Last cylinder or +size or +sizeM or +sizeK (972-8727, default 8727): 8727

```

これらの段階を終えたら、レイアウトは以下のような感じになるはずです：

`Command (m for help):` `p`

```
Disk /dev/sda: 9150 MB, 9150996480 bytes
64 heads, 32 sectors/track, 8727 cylinders
Units = cylinders of 2048 * 512 = 1048576 bytes

   Device Boot      Start         End      Blocks   Id  System
/dev/sda1               1          16       16368    6  FAT16
/dev/sda2              17         971      977920   82  Linux swap
/dev/sda3             972        8727     7942144   83  Linux

```

### パーティションレイアウトを保存し終了する

`w` を打ち、fdiskでの変更を保存してください。

`Command (m for help):` `w`

パーティションを作成したので、 [ファイルシステムを作成する](#.E3.83.95.E3.82.A1.E3.82.A4.E3.83.AB.E3.82.B7.E3.82.B9.E3.83.86.E3.83.A0.E3.82.92.E4.BD.9C.E6.88.90.E3.81.99.E3.82.8B) に進んでください。

## ファイルシステムを作成する

**警告**

SSD または NVMe ドライブを使用する場合は、ファームウェアアップグレードがあるかどうか確認するのが賢明です。特に一部の Intel SSDs (600p および 6000p) は、XFS の I/O 使用量パターンによって誘発される [データ破損の可能性](https://bugzilla.redhat.com/show_bug.cgi?id=1402533) のためのファームウェアアップグレードが必要です。この問題はファームウェアレベルのもので、XFS ファイルシステムの欠陥ではありません。smartctl ユーティリティはデバイスのモデルとファームウェアバージョンを確認するのに役立ちます。

### はじめに

パーティションが作成できたら、その上にファイルシステムを作成します。次の節ではLinuxがサポートする各種ファイルシステムを紹介します。どのファイルシステムを使うかをすでに決めているなら、 [パーティションにファイルシステムを適用する](/wiki/Handbook:Alpha/Installation/Disks/ja#Applying_a_filesystem_to_a_partition "Handbook:Alpha/Installation/Disks/ja") へ進みましょう。そうでなければ、次の節を読んで利用可能なファイルシステムについて知るのがよいでしょう。

### ファイルシステム

Linux は多くのファイルシステムをサポートしていますが、それらの多くは特定の目的をもって配備するのが賢明なものです。特定のファイルシステムのみが alpha アーキテクチャ上で安定して動作するとされています - 重要なパーティションに実験的なファイルシステムを選択するときは、事前にファイルシステムのサポート状況を十分に知っておくことを推奨します。 **XFS はすべてのプラットフォームで、すべての目的で推奨されるファイルシステムです。** 以下は、網羅的ではないリストです:

[XFS](/wiki/XFS/ja "XFS/ja")メタデータジャーナリングのあるファイルシステムで、堅牢な機能セットを持ち、スケーラビリティに最適化されています。新しい機能を取り入れながら継続的にアップグレードされ続けています。唯一の欠点は、現在対応中ではあるものの、 XFS パーティションはまだ縮小できないという点です。XFS の特筆すべき点として reflink とコピーオンライト (CoW) に対応しており、これは Gentoo システム上ではユーザが実施するコンパイル量の多さから特に有用です。XFS は全目的、全プラットフォームで利用できる、おすすめの現代的なファイルシステムです。パーティションは少なくとも 300MB ある必要があります。[ext4](/wiki/Ext4/ja "Ext4/ja")ext4 は reflink などの現代的な機能を欠いてはいるものの、信頼性があり、全目的、全プラットフォームで使用できるファイルシステムです。[VFAT](/wiki/VFAT "VFAT")別名 FAT32。Linux でサポートされていますが、標準的な UNIX パーミッションの設定をサポートしていません。ほとんど、他の OS (Microsoft Windows または Apple macOS) との相互運用性/交換のために使われていますが、いくつかのシステムブートローダーファームウェア (たとえば UEFI) でも必要になります。UEFI システムを使用している場合は、システムをブートするためには VFAT でフォーマットされた [EFI システムパーティション](/wiki/EFI_System_Partition/ja "EFI System Partition/ja") が必要になるでしょう。[btrfs](/wiki/Btrfs/ja "Btrfs/ja")次世代のファイルシステムです。スナップショット、チェックサムによる自己修復、透過的圧縮、サブボリューム、RAID の統合などの先進的な機能を提供します。RAID 5/6 とクオータグループは、btrfs のすべてのバージョンで安全ではありません。[F2FS](/wiki/F2FS/ja "F2FS/ja")Flash-Friendly File System はもともと、Samsung によって NAND フラッシュメモリで利用するために作られました。Gentoo を microSD カードや USB スティックや他のフラッシュベースの記憶装置にインストールする際にはすばらしい選択でしょう。[NTFS](/wiki/NTFS/ja "NTFS/ja")この "New Technology" ファイルシステムは、Windows NT 3.1 以降の Microsoft Windows のフラッグシップファイルシステムです。VFAT と同様、BSD や Linux が正しく動作するために必要な UNIX パーミッション設定や拡張属性を保持しないため、ほとんどの場合ルートファイルシステムとして使うべきではありません。Microsoft Windows とデータ交換の相互運用のために _のみ_ 使うべきです ( _のみ_ の強調に注意してください)。[ZFS](/wiki/ZFS "ZFS")**重要:** ZFS プールは minimal インストール CD で作成することができます。さらなる情報については [ZFS/rootfs](/wiki/ZFS/rootfs "ZFS/rootfs") を参照してください。Matthew Ahrens と Jeff Bonwick によって作成された次世代のファイルシステム。いくつかのキーとなる理想に基づいて設計されました: ストレージの管理はシンプルであるべき、冗長性はファイルシステムによって処理されるべき、ファイルシステムは修復のためにオフラインにされるべきではない、コードをリリースする前に最悪のシナリオを自動的にシミュレーションすることは重要である、そしてデータの整合性は最優先である。

ファイルシステムについてのより広範な情報は、コミュニティによって維持されている [ファイルシステムの記事](/wiki/Filesystem/ja "Filesystem/ja") で見つけることができます。

### パーティションにファイルシステムを適用する

**メモ**

再起動する前に、選択したファイルシステムに関連するユーザ空間ユーティリティを emerge しておいてください。インストールプロセスの終わりのあたりでまた、そうするように再通知します。

パーティションまたはボリュームの上にファイルシステムを作成するには、ファイルシステムごとに異なるユーザースペースのユーティリティが利用可能です。下表でファイルシステムの名前をクリックすると、それぞれに追加の情報が得られます：

ファイルシステム
作成コマンド
live 環境内にある？
パッケージ
[XFS](/wiki/XFS/ja "XFS/ja")mkfs.xfs はい
[sys-fs/xfsprogs](https://packages.gentoo.org/packages/sys-fs/xfsprogs)[ext4](/wiki/Ext4/ja "Ext4/ja")mkfs.ext4 はい
[sys-fs/e2fsprogs](https://packages.gentoo.org/packages/sys-fs/e2fsprogs)[VFAT](/wiki/VFAT "VFAT") (FAT32, ...)
mkfs.vfat はい
[sys-fs/dosfstools](https://packages.gentoo.org/packages/sys-fs/dosfstools)[btrfs](/wiki/Btrfs/ja "Btrfs/ja")mkfs.btrfs はい
[sys-fs/btrfs-progs](https://packages.gentoo.org/packages/sys-fs/btrfs-progs)[F2FS](/wiki/F2FS/ja "F2FS/ja")mkfs.f2fs はい
[sys-fs/f2fs-tools](https://packages.gentoo.org/packages/sys-fs/f2fs-tools)[NTFS](/wiki/NTFS/ja "NTFS/ja")mkfs.ntfs はい
[sys-fs/ntfs3g](https://packages.gentoo.org/packages/sys-fs/ntfs3g)[ZFS](/wiki/ZFS "ZFS")zpool create ... はい
[sys-fs/zfs](https://packages.gentoo.org/packages/sys-fs/zfs)

**重要**

ハンドブックではインストールプロセスの一部として新しいパーティションを使用することを推奨しますが、重要なこととして、mkfs コマンドを実行すると、パーティションに含まれるすべてのデータは消去されることに注意してください。必要な場合は、新しいファイルシステムを作成する前に、中に存在する任意のデータが適切にバックアップされていることを確認してください。

例えば、パーティション構造例の通りに、ルートパーティション (/dev/sda3) を xfs として設定するには、次のコマンドが使えます:

`root #` `mkfs.xfs -c options=/usr/share/xfsprogs/mkfs/lts_6.6.conf /dev/sda3`

#### レガシー BIOS ブートパーティションのファイルシステム

MBR/DOS ディスクラベルを持ち、レガシー BIOS を利用してブートするシステムは、ブートローダがサポートする任意のファイルシステムフォーマットを使用することができます。

例えば、XFS でフォーマットするには:

`root #` `mkfs.xfs -c options=/usr/share/xfsprogs/mkfs/lts_6.6.conf /dev/sda1`

#### 小さな ext4 パーティション

ext4 ファイルシステムを (8 GiB 未満の) 小さいパーティションに使用する場合は、十分な inode 数を確保できるように適切なオプションを指定してファイルシステムを作成するべきです。これは `-T small` オプションを使用して指定することができます:

`root #` `mkfs.ext4 -T small /dev/<device>`

こうすることで「inode あたりのバイト数」が 16kB から 4kB に減るので、ファイルシステムに 4 倍の inode 数を確保できます。

### スワップパーティションを有効にする

mkswapはスワップパーティションを初期化するために使われるコマンドです：

`root #` `mkswap /dev/sda2`

**メモ**

以前に開始したインストール手順を完了しなかった場合は、ハンドブックのこの時点からインストールを再開することができます。このリンクを permalink として使用してください: [インストールの再開はここから](/wiki/Handbook:Alpha/Installation/Disks/ja#Resumed_installations_start_here "Handbook:Alpha/Installation/Disks/ja")。

スワップパーティションを有効化するには、swaponを使います：

`root #` `swapon /dev/sda2`

この「有効化」の手順は、このスワップパーティションが live 環境内に新しく作成されたという理由から必要になっているものです。システムのリブート後は、スワップパーティションが fstab またはその他のマウント機構で適切に定義されている限り、スワップ領域は自動的に有効化されるでしょう。

## ルートパーティションのマウント

一部の live 環境は、Gentoo のルートパーティションのために提案されているマウントポイント (/mnt/gentoo) や、パーティショニングの節で作成された追加のパーティションのためのマウントポイントを持たないかもしれません:

`root #` `mkdir --parents /mnt/gentoo`

以前の手順で作成した追加の (カスタムの) パーティションがあれば、mkdir コマンドを使用して、追加で必要なマウントポイントの作成を続けてください。

マウントポイントが作成できたら、mount コマンドで、パーティションにアクセスできるようにしましょう。

ルートパーティションをマウントしてください:

`root #` `mount /dev/sda3 /mnt/gentoo`

必要に応じて、mount コマンドを使用して、追加の (カスタムの) パーティションのマウントを続けてください。

**メモ**

もし/tmp/を別のパーティションに置く必要があるなら、マウントしたあと権限の変更を忘れずに行ってください:

`root #` `chmod 1777 /mnt/gentoo/tmp`

/var/tmpについても同様です。

このあと解説の中で、proc ファイルシステム (仮想的なカーネルとのインターフェース) が、他のカーネル擬似ファイルシステムと同様にマウントされますが、まず最初は、 [Gentoo stage ファイル](/wiki/Handbook:Alpha/Installation/Stage/ja "Handbook:Alpha/Installation/Stage/ja") を展開する必要があります。

[← ネットワーク設定](/wiki/Handbook:Alpha/Installation/Networking/ja "Previous part") [Home](/wiki/Handbook:Alpha/ja "Handbook:Alpha/ja") [Gentoo インストールファイルをインストール →](/wiki/Handbook:Alpha/Installation/Stage/ja "Next part")