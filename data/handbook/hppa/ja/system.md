# System

Other languages:

- [Deutsch](/wiki/Handbook:HPPA/Installation/System/de "Handbuch:HPPA/Installation/System (100% translated)")
- [English](/wiki/Handbook:HPPA/Installation/System "Handbook:HPPA/Installation/System (100% translated)")
- [español](/wiki/Handbook:HPPA/Installation/System/es "Manual de Gentoo: HPPA/Instalación/Sistema (100% translated)")
- [français](/wiki/Handbook:HPPA/Installation/System/fr "Handbook:HPPA/Installation/System/fr (100% translated)")
- [italiano](/wiki/Handbook:HPPA/Installation/System/it "Handbook:HPPA/Installation/System/it (100% translated)")
- [magyar](/wiki/Handbook:HPPA/Installation/System/hu "Handbook:HPPA/Installation/System/hu (100% translated)")
- [polski](/wiki/Handbook:HPPA/Installation/System/pl "Handbook:HPPA/Installation/System (100% translated)")
- [português do Brasil](/wiki/Handbook:HPPA/Installation/System/pt-br "Manual:HPPA/Instalação/Sistema (100% translated)")
- [русский](/wiki/Handbook:HPPA/Installation/System/ru "Handbook:HPPA/Installation/System (100% translated)")
- [தமிழ்](/wiki/Handbook:HPPA/Installation/System/ta "கையேடு:HPPA/நிறுவல்/முறைமை (100% translated)")
- [中文（中国大陆）‎](/wiki/Handbook:HPPA/Installation/System/zh-cn "手册：HPPA/安装/配置系统 (100% translated)")
- 日本語
- [한국어](/wiki/Handbook:HPPA/Installation/System/ko "Handbook:HPPA/Installation/System/ko (100% translated)")

[HPPA ハンドブック](/wiki/Handbook:HPPA "Handbook:HPPA")[インストール](/wiki/Handbook:HPPA/Full/Installation "Handbook:HPPA/Full/Installation")[インストールについて](/wiki/Handbook:HPPA/Installation/About/ja "Handbook:HPPA/Installation/About/ja")[メディアの選択](/wiki/Handbook:HPPA/Installation/Media/ja "Handbook:HPPA/Installation/Media/ja")[ネットワーク設定](/wiki/Handbook:HPPA/Installation/Networking/ja "Handbook:HPPA/Installation/Networking/ja")[ディスクの準備](/wiki/Handbook:HPPA/Installation/Disks/ja "Handbook:HPPA/Installation/Disks/ja")[stage ファイル](/wiki/Handbook:HPPA/Installation/Stage/ja "Handbook:HPPA/Installation/Stage/ja")[ベースシステムのインストール](/wiki/Handbook:HPPA/Installation/Base/ja "Handbook:HPPA/Installation/Base/ja")[カーネルの設定](/wiki/Handbook:HPPA/Installation/Kernel/ja "Handbook:HPPA/Installation/Kernel/ja")システムの設定[ツールのインストール](/wiki/Handbook:HPPA/Installation/Tools/ja "Handbook:HPPA/Installation/Tools/ja")[ブートローダの設定](/wiki/Handbook:HPPA/Installation/Bootloader/ja "Handbook:HPPA/Installation/Bootloader/ja")[締めくくり](/wiki/Handbook:HPPA/Installation/Finalizing/ja "Handbook:HPPA/Installation/Finalizing/ja")[Gentoo の操作](/wiki/Handbook:HPPA/Full/Working "Handbook:HPPA/Full/Working")[Portage について](/wiki/Handbook:HPPA/Working/Portage/ja "Handbook:HPPA/Working/Portage/ja")[USE フラグ](/wiki/Handbook:HPPA/Working/USE/ja "Handbook:HPPA/Working/USE/ja")[Portage の機能](/wiki/Handbook:HPPA/Working/Features/ja "Handbook:HPPA/Working/Features/ja")[Init スクリプトシステム](/wiki/Handbook:HPPA/Working/Initscripts/ja "Handbook:HPPA/Working/Initscripts/ja")[環境変数](/wiki/Handbook:HPPA/Working/EnvVar/ja "Handbook:HPPA/Working/EnvVar/ja")[Portage の操作](/wiki/Handbook:HPPA/Full/Portage "Handbook:HPPA/Full/Portage")[ファイルとディレクトリ](/wiki/Handbook:HPPA/Portage/Files/ja "Handbook:HPPA/Portage/Files/ja")[変数](/wiki/Handbook:HPPA/Portage/Variables/ja "Handbook:HPPA/Portage/Variables/ja")[ソフトウェアブランチの併用](/wiki/Handbook:HPPA/Portage/Branches/ja "Handbook:HPPA/Portage/Branches/ja")[追加ツール](/wiki/Handbook:HPPA/Portage/Tools/ja "Handbook:HPPA/Portage/Tools/ja")[カスタムパッケージリポジトリ](/wiki/Handbook:HPPA/Portage/CustomTree/ja "Handbook:HPPA/Portage/CustomTree/ja")[高度な機能](/wiki/Handbook:HPPA/Portage/Advanced/ja "Handbook:HPPA/Portage/Advanced/ja")[OpenRC ネットワーク設定](/wiki/Handbook:HPPA/Full/Networking "Handbook:HPPA/Full/Networking")[はじめに](/wiki/Handbook:HPPA/Networking/Introduction/ja "Handbook:HPPA/Networking/Introduction/ja")[高度な設定](/wiki/Handbook:HPPA/Networking/Advanced/ja "Handbook:HPPA/Networking/Advanced/ja")[モジュール式ネットワーク](/wiki/Handbook:HPPA/Networking/Modular/ja "Handbook:HPPA/Networking/Modular/ja")[無線](/wiki/Handbook:HPPA/Networking/Wireless/ja "Handbook:HPPA/Networking/Wireless/ja")[機能の追加](/wiki/Handbook:HPPA/Networking/Extending/ja "Handbook:HPPA/Networking/Extending/ja")[動的な管理](/wiki/Handbook:HPPA/Networking/Dynamic/ja "Handbook:HPPA/Networking/Dynamic/ja")

## Contents

- [1ファイルシステムの情報](#.E3.83.95.E3.82.A1.E3.82.A4.E3.83.AB.E3.82.B7.E3.82.B9.E3.83.86.E3.83.A0.E3.81.AE.E6.83.85.E5.A0.B1)
  - [1.1ファイルシステムラベルと UUID](#.E3.83.95.E3.82.A1.E3.82.A4.E3.83.AB.E3.82.B7.E3.82.B9.E3.83.86.E3.83.A0.E3.83.A9.E3.83.99.E3.83.AB.E3.81.A8_UUID)
  - [1.2パーティションラベルと UUID](#.E3.83.91.E3.83.BC.E3.83.86.E3.82.A3.E3.82.B7.E3.83.A7.E3.83.B3.E3.83.A9.E3.83.99.E3.83.AB.E3.81.A8_UUID)
  - [1.3fstab について](#fstab_.E3.81.AB.E3.81.A4.E3.81.84.E3.81.A6)
  - [1.4fstab ファイルを作成する](#fstab_.E3.83.95.E3.82.A1.E3.82.A4.E3.83.AB.E3.82.92.E4.BD.9C.E6.88.90.E3.81.99.E3.82.8B)
    - [1.4.1DOS/レガシー BIOS システム](#DOS.2F.E3.83.AC.E3.82.AC.E3.82.B7.E3.83.BC_BIOS_.E3.82.B7.E3.82.B9.E3.83.86.E3.83.A0)
    - [1.4.2DPS UEFI PARTUUID](#DPS_UEFI_PARTUUID)
- [2ネットワーク接続のための情報](#.E3.83.8D.E3.83.83.E3.83.88.E3.83.AF.E3.83.BC.E3.82.AF.E6.8E.A5.E7.B6.9A.E3.81.AE.E3.81.9F.E3.82.81.E3.81.AE.E6.83.85.E5.A0.B1)
  - [2.1ホスト名](#.E3.83.9B.E3.82.B9.E3.83.88.E5.90.8D)
    - [2.1.1ホスト名を設定する (OpenRC または systemd)](#.E3.83.9B.E3.82.B9.E3.83.88.E5.90.8D.E3.82.92.E8.A8.AD.E5.AE.9A.E3.81.99.E3.82.8B_.28OpenRC_.E3.81.BE.E3.81.9F.E3.81.AF_systemd.29)
    - [2.1.2systemd](#systemd)
  - [2.2ネットワーク](#.E3.83.8D.E3.83.83.E3.83.88.E3.83.AF.E3.83.BC.E3.82.AF)
    - [2.2.1dhcpcd での DHCP (どんな init システムでも)](#dhcpcd_.E3.81.A7.E3.81.AE_DHCP_.28.E3.81.A9.E3.82.93.E3.81.AA_init_.E3.82.B7.E3.82.B9.E3.83.86.E3.83.A0.E3.81.A7.E3.82.82.29)
    - [2.2.2netifrc (OpenRC)](#netifrc_.28OpenRC.29)
      - [2.2.2.1ネットワークを設定する](#.E3.83.8D.E3.83.83.E3.83.88.E3.83.AF.E3.83.BC.E3.82.AF.E3.82.92.E8.A8.AD.E5.AE.9A.E3.81.99.E3.82.8B)
      - [2.2.2.2起動時に自動でネットワーク接続する](#.E8.B5.B7.E5.8B.95.E6.99.82.E3.81.AB.E8.87.AA.E5.8B.95.E3.81.A7.E3.83.8D.E3.83.83.E3.83.88.E3.83.AF.E3.83.BC.E3.82.AF.E6.8E.A5.E7.B6.9A.E3.81.99.E3.82.8B)
  - [2.3hosts ファイル](#hosts_.E3.83.95.E3.82.A1.E3.82.A4.E3.83.AB)
- [3システム情報](#.E3.82.B7.E3.82.B9.E3.83.86.E3.83.A0.E6.83.85.E5.A0.B1)
  - [3.1root パスワード](#root_.E3.83.91.E3.82.B9.E3.83.AF.E3.83.BC.E3.83.89)
  - [3.2init と boot 設定](#init_.E3.81.A8_boot_.E8.A8.AD.E5.AE.9A)
    - [3.2.1OpenRC](#OpenRC)
    - [3.2.2systemd](#systemd_2)

## ファイルシステムの情報

### ファイルシステムラベルと UUID

MBR（BIOS）とGPTの両方が、 _ファイルシステム_ ラベルと _ファイルシステム_ UUIDをサポートしています。これらの属性は、ブロックデバイスを探しマウントするために使用されるmountコマンドの代用として、/etc/fstab内で定義することができます。ファイルシステムラベルやファイルシステムUUIDはそれぞれLABELとUUID接頭辞で識別され、blkidコマンドで確認することができます：

`root #` `blkid`

**警告**

もし、パーティション内のファイルシステムが消滅すると、ファイルシステムのラベルとUUIDの値は後に変更されるか除去されます。

MBR スタイルのパーティションテーブルを使用している場合は、一意性のため、/etc/fstab 内でマウント可能なボリュームを指定するためにはラベルよりも UUID を用いることが推奨されます。

**重要**

LVM ボリューム上に置かれたファイルシステムの UUID と、LVM ボリュームの LVM スナップショットの UUID は同じなので、LVM ボリュームをマウントするために UUID を使用するのは避けるべきです。

### パーティションラベルと UUID

GPT ディスクラベルへのサポートを持つシステムでは、/etc/fstab でパーティションを定義するための '堅牢' な追加の選択肢が提供されます。パーティション自体にどのファイルシステムが使用されているかにかかわらず、ブロックデバイスの個々のパーティションを識別するのにパーティションラベルやパーティションの UUID を使うことができます。パーティションラベルとパーティションの UUID は PARTLABEL、PARTUUID 接頭辞で識別され、これらは端末で blkid コマンドを実行することで簡単に確認できます。

Discoverable Partition Specification の UUID を使用する **amd64** EFI システムでの出力は、次のようになるでしょう:

`root #` `blkid`

```
/dev/sr0: BLOCK_SIZE="2048" UUID="2023-08-28-03-54-40-00" LABEL="ISOIMAGE" TYPE="iso9660" PTTYPE="PMBR"
/dev/loop0: TYPE="squashfs"
/dev/sda2: PARTUUID="0657fd6d-a4ab-43c4-84e5-0933c84b4f4f"
/dev/sda3: PARTUUID="1cdf763a-5b4c-4dbf-99db-a056c504e8b2"
/dev/sda1: PARTUUID="c12a7328-f81f-11d2-ba4b-00a0c93ec93b"

```

パーティションラベルも絶対にとは言えないのに対し、fstab で UUID を使ったパーティション指定を使えば、たとえ将来 _ファイルシステム_ が変更されたり、再度書き込まれたりしても、ブートローダーがボリューム検出に迷うことはありません。従来のブロックデバイスファイル (/dev/sd\*N) を使った指定は、SATA ブロックデバイスの追加または削除が日常的に行われるシステムでは危険です。

ブロックデバイスのファイル名は様々な要素 (ディスクがどんな順番でいくつ接続されているかを含む) によって変化します。またファイル名の順番についても、初期起動プロセス中にカーネルがどのデバイスを最初に検知するかによって変化します。つまり、システム管理者がディスクの順序を頻繁にいじったりしない限りは、デフォルトのブロックデバイスファイルを使うのがシンプルで素直な方法です。

### fstab について

Linuxでは、システムで使用するすべてのパーティションは[/etc/fstab](/wiki//etc/fstab/ja "/etc/fstab/ja")に記載されていなければなりません。このファイルは、これらパーティションのマウントポイント（これらはファイルシステムに存在しなければなりません）、どのようにマウントされるべきか、また特別なオプション（自動マウントかそうでないか、ユーザー権限でマウントできるかどうか等）を定義します。

### fstab ファイルを作成する

**メモ**

init システムに systemd を使用しており、パーティション UUID が [Preparing the disks](/wiki/Handbook:HPPA/Installation/Disks/ja "Handbook:HPPA/Installation/Disks/ja") で説明されている Discoverable Partition Specification に適合していて、かつ UEFI を使用しているシステムでは、systemd が仕様に従ってパーティションを自動的にマウントするので、fstab の作成を省略できます。

/etc/fstabファイルは表のように記述します。それぞれの行はホワイトスペース（一つまたは複数のスペース、タブ、もしくはその 2 種の組み合わせ）で区切られる 6 つのフィールドを持ちます。それぞれのフィールドの意味は以下の通りです。

1. 最初のフィールドはマウントされるブロックスペシャルデバイスやリモートファイルシステムを示します。デバイスファイルへのパスや、ファイルシステムラベルやファイルシステムUUID，そしてパーティションラベルやパーティションUUIDを含む、いくつかの種類のデバイスIDがブロックスペシャルデバイスノードとして使用可能です。
2. 2番目のフィールドはそのパーティションがマウントされるマウントポイントを示します。
3. 3番目のフィールドはそのパーティションのファイルシステムの種類を示します。
4. 4番目のフィールドは、そのパーティションをマウントするmountコマンドが使用するオプションを示します。すべてのファイルシステムは、固有のマウントオプションを持っています。システム管理者はマウントコマンドのmanページ(man mount)を参照することですべてのオプションを確認できます。複数のマウントオプションを記述する場合はカンマで区切ります。
5. 5番目のフィールドはそのパーティションをdumpでダンプするかどうかを示しています。このフィールドは通常 `0`（ゼロ）のままにしておいてかまいません。
6. 6番目のフィールドは、直前のシャットダウンが正常に完了しなかったときに、fsckが各パーティションをどの順番でチェックするか示しています。ルートファイルシステムは `1` であるべきです。残りのファイルシステムは `2`（ファイルシステムチェックが不要であれば `0`）に設定しましょう。

**重要**

Gentoo stage ファイルに含まれるデフォルトの /etc/fstab ファイルは有効な fstab では _なく_、関連する値を入力するために使用できるテンプレートです。

`root #` `nano /etc/fstab`

#### DOS/レガシー BIOS システム

では、/boot パーティションをどのように記述すればよいか見てみましょう。これは一つの例です。実際はインストール手順の最初に決めたパーティション構成通りに修正しなければなりません。
**hppa** のパーティション構成の例として、/boot は通常は /dev/sda2 パーティションになり、推奨されるファイルシステムである xfs を使います。このパーティションはブート中にチェックされなければならないので、fstab は以下のようになるでしょう:

ファイル **`/etc/fstab`** **/etc/fstab の DOS/レガシー BIOS boot 行の例**

```
# 整形上の差異や、「ディスクの準備」ステップで作成する追加のパーティションについては、適宜修正してください
/dev/sda2   /boot     ext2    defaults        0 2

```

システム管理者によっては、セキュリティを向上させるために /boot パーティションを自動的にマウントしたくないかもしれません。その場合は `defaults` を `noauto` で置き換えてください。これは、そのパーティションを使いたいときは都度手動でマウントしなければならないことを意味します。

実際のパーティション構成にあわせたルールや、CD-ROMドライブのためのルールを追加してください。他にパーティションやドライブがあれば、それも忘れずに追加しておきましょう。

以下は、より詳細な/etc/fstabの例です。

ファイル **`/etc/fstab`** **DOS/レガシー BIOS システム向けの完全な /etc/fstab の例**

```
# 整形上の差異や、「ディスクの準備」ステップで作成する追加のパーティションについては、適宜修正してください
/dev/sda2   /boot        ext2    defaults    0 2
/dev/sda3   none         swap    sw                   0 0
/dev/sda4   /            xfs    defaults,noatime              0 1

/dev/cdrom  /mnt/cdrom   auto    noauto,user          0 0

```

#### DPS UEFI PARTUUID

以下は、UEFI ファームウェア向けに GPT ディスクラベルと Discoverable Partition Specification (DPS) の UUID を使用してフォーマットされたディスクのための、/etc/fstab ファイルの例です:

ファイル **`/etc/fstab`** **GPT ディスクラベル DPS PARTUUID fstab の例**

```
# 整形上の差異や、「ディスクの準備」ステップで作成する追加のパーティションについては、適宜修正してください。
# この例は Discoverable Partition Specification (DSP) UUID が設定された GPT ディスクラベルを示しています:
PARTUUID=c12a7328-f81f-11d2-ba4b-00a0c93ec93b   /efi        vfat                       0 2
PARTUUID=0657fd6d-a4ab-43c4-84e5-0933c84b4f4f   none            sw                           0 0
PARTUUID=1aacdb3b-5444-4138-bd9e-e5c2239b2346   /           xfs    defaults,noatime              0 1

```

3番目のフィールドで `auto` を使う場合、mountコマンドはそのパーティションのファイルシステムが何かを推測します。これは様々なファイルシステムを使う可能性があるリムーバルメディアで推奨されます。4番目の `user` オプションで、ルート権限を持たないユーザーがCDをマウントできるようになります。

パフォーマンスを改善するために、多くのユーザーはマウントオプションとして `noatime` オプションを付け加えたいと考えるでしょう。アクセス時間が記録されないので、結果としてより高速なシステムになります (一般的にこの記録はほとんど必要ありません)。ソリッドステートドライブ (SSD) を持つシステムでもこれはおすすめです。代わりに `lazytime` もまた考慮に値するかもしれません。

**ヒント**

パフォーマンスを低下させるため、/etc/fstab 内で `discard` マウントオプションを定義させるのは推奨されません。cron または timer (systemd) のようなジョブスケジューラを使用して、定期的にブロックの破棄をスケジュールするほうが、一般的により良い方法です。さらなる情報については、 [Periodic fstrim jobs](/wiki/SSD#Periodic_fstrim_jobs "SSD") を確認してください。

再度/etc/fstabを確認して、保存、エディタを終了します。

## ネットワーク接続のための情報

重要な注意点として、このセクションは読者がシステムをローカルエリアネットワークに手っ取り早く参加させることができるように、提供しています。

OpenRC を動かしているシステムについては、ネットワーク設定のためのより詳細なリファレンスが、ハンドブックの終わりのほうの [高度なネットワーク設定](/wiki/Handbook:HPPA/Networking/Introduction/ja "Handbook:HPPA/Networking/Introduction/ja") のセクションでカバーされています。より詳細なネットワーク要件のあるシステムは一度そちらへ飛んで、それからここに戻ってきて残りのインストール作業を続ける必要があるかもしれません。

より詳細な systemd のネットワーク設定については、 [systemd](/wiki/Systemd/ja "Systemd/ja") の記事の [ネットワークの箇所](/wiki/Systemd/ja#Network "Systemd/ja") を確認してください。

### ホスト名

さて、PCには名前をつけなければいけません。至極簡単に思えますが多くのシステム管理者はホスト名として適切な名前を付けるのに苦労しています。事を早く進めるために、選んだ名前は後で変更できることを知っておいてください。判りやすいように、ここでは単にマシンを _tux_ と呼ぶことにします。

#### ホスト名を設定する (OpenRC または systemd)

`root #` `echo tux > /etc/hostname`

#### systemd

現在 _実行中の_ systemd についてシステムのホスト名を設定するには、hostnamectl ユーティリティを使うことができます。しかしながらインストールプロセス中は、 [systemd-firstboot](/wiki/Handbook:HPPA/Installation/System/ja#Init_and_boot_configuration_systemd "Handbook:HPPA/Installation/System/ja") コマンドを代わりに使う必要があります (後でハンドブックの該当箇所を参照してください)。

ホスト名を "tux" に設定するためには、次のようにします:

`root #` `hostnamectl hostname tux`

hostnamectl --help または man 1 hostnamectl を実行することで、ヘルプを確認してください。

### ネットワーク

ネットワークインターフェースを構成するために利用できる選択肢は _多く_ 存在します。この節ではそのうち一部の方法だけをカバーします。必要な構成に最適と思われるものをひとつ選んでください。

#### dhcpcd での DHCP (どんな init システムでも)

多くの LAN ネットワークは DHCP サーバは運用しています。その場合は、dhcpcd プログラムを使用して IP アドレスを取得するのがおすすめです。

インストールするには:

`root #` `emerge --ask net-misc/dhcpcd`

OpenRC システム上で、サービスを有効化して開始するには:

`root #` `rc-update add dhcpcd default
`

`root #` `rc-service dhcpcd start
`

systemd システム上で、サービスを有効化するには:

`root #` `systemctl enable dhcpcd`

これらのステップが完了したら、次にシステムが起動したときに、dhcpcd が DHCP サーバから IP アドレスを取得するはずです。さらなる詳細については [Dhcpcd](/wiki/Dhcpcd/ja "Dhcpcd/ja") の記事を確認してください。

#### netifrc (OpenRC)

**ヒント**

これは OpenRC 上で [Netifrc](/wiki/Netifrc "Netifrc") を使ってネットワークをセットアップするときに固有の方法です。他にも [Dhcpcd](/wiki/Dhcpcd/ja "Dhcpcd/ja") のような、より簡潔なセットアップの方法も存在します。

##### ネットワークを設定する

Gentoo Linux をインストールしている間、ネットワークが使えるように設定されています。しかし、それは live 環境のためのネットワーク設定であり、インストールされた環境のためのものではありません。では、インストールされた Gentoo Linux のネットワークを設定しましょう。

**メモ**

bonding、ブリッジ、802.1Q VLAN、無線ネットワークに間するより詳細な情報は、 [追加のネットワーク設定](/wiki/Handbook:HPPA/Networking/Introduction/ja "Handbook:HPPA/Networking/Introduction/ja") セクションを参照してください。

すべてのネットワーク設定は/etc/conf.d/netにあります。直接的ではありますが、おそらく直感で理解できる構文ではありません。しかし恐れることはありません。すべては以下で説明されます。/usr/share/doc/netifrc-\*/net.example.bz2に、多くの異なる設定に対して完全にコメントが付与された例が記載されています。

最初に [net-misc/netifrc](https://packages.gentoo.org/packages/net-misc/netifrc) をインストールします。

`root #` `emerge --ask --noreplace net-misc/netifrc`

DHCPはデフォルトで使用されます。DHCPを動かすために、DHCPクライアントをインストールしなければなりません。これは Installing Necessary System Toolsで説明されます。

もし、特別なDHCPのオプションを設定している、もしくはDHCPをまったく使いたくない等の理由で、ネットワーク接続をしなければならないときは、/etc/conf.d/netを編集します。

`root #` `nano /etc/conf.d/net`

IPアドレスとルーティングを設定するのはconfig\_eth0とroutes\_eth0です。

**メモ**

ここではネットワークインターフェイスがeth0であると仮定していますが、これはシステムによって違います。もし、最近のインストールメディアから起動しているのであれば、インストール時と同じインターフェイス名が使われると思ってよいでしょう。より詳しい情報は [ネットワークインターフェースの命名](/wiki/Handbook:HPPA/Networking/Advanced/ja#Network_interface_naming "Handbook:HPPA/Networking/Advanced/ja") セクションを参照してください。

ファイル **`/etc/conf.d/net`** **静的 IP アドレスの定義**

```
config_eth0="192.168.0.2 netmask 255.255.255.0 brd 192.168.0.255"
routes_eth0="default via 192.168.0.1"

```

DHCPを使う場合は、config\_eth0を設定してください。

ファイル **`/etc/conf.d/net`** **DHCPの定義**

```
config_eth0="dhcp"

```

さらなる設定オプションのリストについては、/usr/share/doc/netifrc-\*/net.example.bz2を参照してください。もし特定のDHCPを設定しなければならないときは、DHCPクライアントのmanページも必ず読みましょう。

もし、システムが複数のネットワークインターフェースを持っている場合は、config\_eth1、config\_eth2、…に対して上記の手順を繰り返してください。

設定を保存し、エディタを終了しましょう。

##### 起動時に自動でネットワーク接続する

ブート時にネットワークインターフェースを有効にする場合は、デフォルトランレベルにそれらを追加する必要があります。

`root #` `cd /etc/init.d
`

`root #` `ln -s net.lo net.eth0
`

`root #` `rc-update add net.eth0 default`

もし、複数のネットワークインターフェースがある場合は、net.eth0と同じ方法で、適切なnet.\*ファイルを作成しなければなりません。

もし、ブート後、ネットワークインターフェース名（現在、このドキュメントでは `eth0` と記述）が間違っていた場合、次の手順で修正してください。

1. 正しいインターフェース名で/etc/conf.d/netファイルを更新します。（例えば `eth0` を `enp3s0` または `enp5s0` と修正）
2. 新しいシンボリックリンクを作成。（例えば/etc/init.d/net.enp3s0）
3. 古いシンボリックリンクを消去。（rm /etc/init.d/net.eth0）
4. 新しいスクリプトをデフォルトランレベルに追加。
5. rc-update del net.eth0 defaultで古いスクリプトを消去。

### hosts ファイル

次の重要なステップは、この新しいシステムのネットワーク環境内の他のホストについて、システムに教えることかもしれません。ネットワークホスト名は /etc/hosts で定義することができます。ホスト名をここに追加することで、ネームサーバでは解決できないホストについて、ホスト名から IP アドレスを解決できるようになります。

`root #` `nano /etc/hosts`

ファイル **`/etc/hosts`** **ネットワーク情報の記述**

```
# 以下は本システムの定義です。必ず設定されなければなりません。
127.0.0.1     tux.homenetwork tux localhost
::1           tux.homenetwork tux localhost

# ネットワーク上にあるその他のホストの定義です。任意設定です。
192.168.0.5   jenny.homenetwork jenny
192.168.0.6   benny.homenetwork benny

```

設定をセーブし、エディタを終了しましょう。

## システム情報

### root パスワード

passwdコマンドでルートのパスワードを設定します。

`root #` `passwd`

後で、日常の作業のための一般ユーザーアカウントを作成します。

### init と boot 設定

#### OpenRC

Gentoo で OpenRC を使用しているときは、OpenRC はシステムのサービス、スタートアップ、シャットダウンの設定に /etc/rc.conf を使います。/etc/rc.conf を開いて、ファイル中のすべてのコメントを楽しみましょう。設定をレビューして、必要な箇所を変更してください。

`root #` `nano /etc/rc.conf`

次に、キーボードを設定するために/etc/conf.d/keymapsを開いて、正しいキーボードを選択、設定します。

`root #` `nano /etc/conf.d/keymaps`

keymap変数に特に注意してください。もしキーマップを間違えた場合、キーボードを叩くたびに、奇妙な現象が起こるでしょう。

最後に、クロック設定をするために/etc/conf.d/hwclockを編集します。個々の好みに合わせて設定できます。

`root #` `nano /etc/conf.d/hwclock`

もし、ハードウェアクロックがUTCになっていない場合、このファイルに `clock="local"` を記述しなければなりません。そうでない場合、クロックスキューが発生するでしょう。

#### systemd

まず、システムが正しく起動できるようにするために、systemd-machine-id-setup と systemd-firstboot を順に実行することをおすすめします。これは、新しい systemd 環境への最初のブートのために、システムのさまざまなコンポーネントが正しく設定されるように準備します。次のオプションを渡すことで、ロケール、タイムゾーン、ホスト名、root パスワード、そして root シェル値を設定するための、ユーザへのプロンプトを含めます。加えて、システムにランダムなマシン ID を割り当てます:

`root #` `systemd-machine-id-setup`

`root #` `systemd-firstboot --prompt`

次に、インストールされているすべてのユニットファイルをプリセットのポリシー値にリセットするため、ユーザは systemctl を実行すべきです:

**メモ**

次のコマンドを実行した後、失敗の警告が出力される場合があります。これは想定通りです。Gentoo の LiveCD は OpenRC を使用しているからです。

`root #` `systemctl preset-all --preset-mode=enable-only`

完全にプリセットにするには次で行えますが、これまでのプロセスで既に設定したすべてのサービスもリセットするかもしれません:

`root #` `systemctl preset-all`

live 環境からインストール先の最初のブートへのシームレスな移行を確実に行うために、これらの 2 ステップは有用でしょう。

[← カーネルの設定](/wiki/Handbook:HPPA/Installation/Kernel/ja "Previous part") [Home](/wiki/Handbook:HPPA/ja "Handbook:HPPA/ja") [ツールのインストール →](/wiki/Handbook:HPPA/Installation/Tools/ja "Next part")