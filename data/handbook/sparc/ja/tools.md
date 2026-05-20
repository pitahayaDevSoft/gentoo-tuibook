# Tools

Other languages:

- [Deutsch](/wiki/Handbook:SPARC/Installation/Tools/de "Handbuch:SPARC/Installation/Tools (100% translated)")
- [English](/wiki/Handbook:SPARC/Installation/Tools "Handbook:SPARC/Installation/Tools (100% translated)")
- [español](/wiki/Handbook:SPARC/Installation/Tools/es "Manual de Gentoo: SPARC/Instalación/Herramientas (100% translated)")
- [français](/wiki/Handbook:SPARC/Installation/Tools/fr "Handbook:SPARC/Installation/Tools/fr (100% translated)")
- [italiano](/wiki/Handbook:SPARC/Installation/Tools/it "Handbook:SPARC/Installation/Tools/it (50% translated)")
- [magyar](/wiki/Handbook:SPARC/Installation/Tools/hu "Handbook:SPARC/Installation/Tools/hu (100% translated)")
- [polski](/wiki/Handbook:SPARC/Installation/Tools/pl "Handbook:SPARC/Installation/Tools (100% translated)")
- [português do Brasil](/wiki/Handbook:SPARC/Installation/Tools/pt-br "Handbook:SPARC/Installation/Tools/pt-br (50% translated)")
- [русский](/wiki/Handbook:SPARC/Installation/Tools/ru "Handbook:SPARC/Installation/Tools (100% translated)")
- [தமிழ்](/wiki/Handbook:SPARC/Installation/Tools/ta "கையேடு:SPARC/நிறுவல்/கருவிகள் (100% translated)")
- [中文（中国大陆）‎](/wiki/Handbook:SPARC/Installation/Tools/zh-cn "手册：SPARC/安装/安装系统工具 (100% translated)")
- 日本語
- [한국어](/wiki/Handbook:SPARC/Installation/Tools/ko "Handbook:SPARC/Installation/Tools/ko (100% translated)")

[SPARC ハンドブック](/wiki/Handbook:SPARC "Handbook:SPARC")[インストール](/wiki/Handbook:SPARC/Full/Installation "Handbook:SPARC/Full/Installation")[インストールについて](/wiki/Handbook:SPARC/Installation/About/ja "Handbook:SPARC/Installation/About/ja")[メディアの選択](/wiki/Handbook:SPARC/Installation/Media/ja "Handbook:SPARC/Installation/Media/ja")[ネットワーク設定](/wiki/Handbook:SPARC/Installation/Networking/ja "Handbook:SPARC/Installation/Networking/ja")[ディスクの準備](/wiki/Handbook:SPARC/Installation/Disks/ja "Handbook:SPARC/Installation/Disks/ja")[stage ファイル](/wiki/Handbook:SPARC/Installation/Stage/ja "Handbook:SPARC/Installation/Stage/ja")[ベースシステムのインストール](/wiki/Handbook:SPARC/Installation/Base/ja "Handbook:SPARC/Installation/Base/ja")[カーネルの設定](/wiki/Handbook:SPARC/Installation/Kernel/ja "Handbook:SPARC/Installation/Kernel/ja")[システムの設定](/wiki/Handbook:SPARC/Installation/System/ja "Handbook:SPARC/Installation/System/ja")ツールのインストール[ブートローダの設定](/wiki/Handbook:SPARC/Installation/Bootloader/ja "Handbook:SPARC/Installation/Bootloader/ja")[締めくくり](/wiki/Handbook:SPARC/Installation/Finalizing/ja "Handbook:SPARC/Installation/Finalizing/ja")[Gentoo の操作](/wiki/Handbook:SPARC/Full/Working "Handbook:SPARC/Full/Working")[Portage について](/wiki/Handbook:SPARC/Working/Portage/ja "Handbook:SPARC/Working/Portage/ja")[USE フラグ](/wiki/Handbook:SPARC/Working/USE/ja "Handbook:SPARC/Working/USE/ja")[Portage の機能](/wiki/Handbook:SPARC/Working/Features/ja "Handbook:SPARC/Working/Features/ja")[Init スクリプトシステム](/wiki/Handbook:SPARC/Working/Initscripts/ja "Handbook:SPARC/Working/Initscripts/ja")[環境変数](/wiki/Handbook:SPARC/Working/EnvVar/ja "Handbook:SPARC/Working/EnvVar/ja")[Portage の操作](/wiki/Handbook:SPARC/Full/Portage "Handbook:SPARC/Full/Portage")[ファイルとディレクトリ](/wiki/Handbook:SPARC/Portage/Files/ja "Handbook:SPARC/Portage/Files/ja")[変数](/wiki/Handbook:SPARC/Portage/Variables/ja "Handbook:SPARC/Portage/Variables/ja")[ソフトウェアブランチの併用](/wiki/Handbook:SPARC/Portage/Branches/ja "Handbook:SPARC/Portage/Branches/ja")[追加ツール](/wiki/Handbook:SPARC/Portage/Tools/ja "Handbook:SPARC/Portage/Tools/ja")[カスタムパッケージリポジトリ](/wiki/Handbook:SPARC/Portage/CustomTree/ja "Handbook:SPARC/Portage/CustomTree/ja")[高度な機能](/wiki/Handbook:SPARC/Portage/Advanced/ja "Handbook:SPARC/Portage/Advanced/ja")[OpenRC ネットワーク設定](/wiki/Handbook:SPARC/Full/Networking "Handbook:SPARC/Full/Networking")[はじめに](/wiki/Handbook:SPARC/Networking/Introduction/ja "Handbook:SPARC/Networking/Introduction/ja")[高度な設定](/wiki/Handbook:SPARC/Networking/Advanced/ja "Handbook:SPARC/Networking/Advanced/ja")[モジュール式ネットワーク](/wiki/Handbook:SPARC/Networking/Modular/ja "Handbook:SPARC/Networking/Modular/ja")[無線](/wiki/Handbook:SPARC/Networking/Wireless/ja "Handbook:SPARC/Networking/Wireless/ja")[機能の追加](/wiki/Handbook:SPARC/Networking/Extending/ja "Handbook:SPARC/Networking/Extending/ja")[動的な管理](/wiki/Handbook:SPARC/Networking/Dynamic/ja "Handbook:SPARC/Networking/Dynamic/ja")

## Contents

- [1システムロガー](#.E3.82.B7.E3.82.B9.E3.83.86.E3.83.A0.E3.83.AD.E3.82.AC.E3.83.BC)
  - [1.1OpenRC](#OpenRC)
  - [1.2systemd](#systemd)
- [2任意自由選択: cronデーモン](#.E4.BB.BB.E6.84.8F.E8.87.AA.E7.94.B1.E9.81.B8.E6.8A.9E:_cron.E3.83.87.E3.83.BC.E3.83.A2.E3.83.B3)
  - [2.1OpenRC](#OpenRC_2)
    - [2.1.1デフォルト: cronie](#.E3.83.87.E3.83.95.E3.82.A9.E3.83.AB.E3.83.88:_cronie)
    - [2.1.2代替案: dcron](#.E4.BB.A3.E6.9B.BF.E6.A1.88:_dcron)
    - [2.1.3代替案: fcron](#.E4.BB.A3.E6.9B.BF.E6.A1.88:_fcron)
    - [2.1.4代替案: bcron](#.E4.BB.A3.E6.9B.BF.E6.A1.88:_bcron)
    - [2.1.5modprobed-db users](#modprobed-db_users)
  - [2.2systemd](#systemd_2)
    - [2.2.1modprobed-db users](#modprobed-db_users_2)
- [3任意自由選択: ファイルのインデックスを作成](#.E4.BB.BB.E6.84.8F.E8.87.AA.E7.94.B1.E9.81.B8.E6.8A.9E:_.E3.83.95.E3.82.A1.E3.82.A4.E3.83.AB.E3.81.AE.E3.82.A4.E3.83.B3.E3.83.87.E3.83.83.E3.82.AF.E3.82.B9.E3.82.92.E4.BD.9C.E6.88.90)
- [4任意自由選択: リモートシェルアクセス](#.E4.BB.BB.E6.84.8F.E8.87.AA.E7.94.B1.E9.81.B8.E6.8A.9E:_.E3.83.AA.E3.83.A2.E3.83.BC.E3.83.88.E3.82.B7.E3.82.A7.E3.83.AB.E3.82.A2.E3.82.AF.E3.82.BB.E3.82.B9)
  - [4.1OpenRC](#OpenRC_3)
  - [4.2systemd](#systemd_3)
- [5省略可能: シェル補完](#.E7.9C.81.E7.95.A5.E5.8F.AF.E8.83.BD:_.E3.82.B7.E3.82.A7.E3.83.AB.E8.A3.9C.E5.AE.8C)
  - [5.1Bash](#Bash)
- [6推奨: 時刻同期](#.E6.8E.A8.E5.A5.A8:_.E6.99.82.E5.88.BB.E5.90.8C.E6.9C.9F)
  - [6.1chrony](#chrony)
  - [6.2OpenRC](#OpenRC_4)
  - [6.3systemd](#systemd_4)
  - [6.4systemd-timesyncd](#systemd-timesyncd)
- [7ファイルシステムツール](#.E3.83.95.E3.82.A1.E3.82.A4.E3.83.AB.E3.82.B7.E3.82.B9.E3.83.86.E3.83.A0.E3.83.84.E3.83.BC.E3.83.AB)

## システムロガー

### OpenRC

同じ機能が複数のパッケージによって提供されるツールがいくつかあります。そういったツールはstage3アーカイブには含まれていません。どのパッケージをインストールしたいのかをあなた次第で選んでください。

まずシステムにロギング機構を提供するツールを決定しましょう。UnixとLinuxでは歴史をかけて素晴らしいログ機能を発展させてきました -- お望みならログファイルにシステムで起こった全てを記録できます。

Gentooでは複数のシステムロガーから使いたいものを選択することができます。このうちのいくつかを紹介します。

- [sysklogd](/wiki/Sysklogd "Sysklogd") ([app-admin/sysklogd](https://packages.gentoo.org/packages/app-admin/sysklogd)) は、システムのログを取得するための伝統的なデーモンを集めたものです。デフォルトのログ設定をそのまま使ってもうまく働くので、このパッケージは初心者にはいい選択肢です。
- [syslog-ng](/wiki/Syslog-ng "Syslog-ng") ([app-admin/syslog-ng](https://packages.gentoo.org/packages/app-admin/syslog-ng)) は、進化したシステムロガーです。1つの大きなファイルにログを取る以上のことをするには、何らかの設定が必要です。更に上級のユーザは、ロギングの発展性に基いてこのパッケージを選択できます。スマートなロギングのためには追加の設定が必要になることに注意してください。
- [metalog](/wiki/Metalog "Metalog") ([app-admin/metalog](https://packages.gentoo.org/packages/app-admin/metalog)) は、高度な設定ができるシステムロガーです。

Gentoo ebuild リポジトリにはまだまだ他のシステムロガーがあることでしょう。日毎に利用可能なパッケージは増えていますから。

**ヒント**

もし syslog-ng を使おうと思っているなら、後で [logrotate](/wiki/Logrotate "Logrotate") をインストールして設定しましょう。syslog-ng にはログファイルをローテーションする機構がありません。一方で、より新しいバージョン (>= 2.0) の sysklogd は自身でログローテーションを行います。

選択したシステムログツールをインストールするには、それを emerge してください。OpenRC では、rc-update を使ってデフォルトのランレベルにスクリプトを追加してください。次の例では [app-admin/sysklogd](https://packages.gentoo.org/packages/app-admin/sysklogd) をインストールして、それをシステムの syslog ユーティリティとして有効化します:

`root #` `emerge --ask app-admin/sysklogd`

`root #` `rc-update add sysklogd default`

### systemd

ここまで OpenRC ベースのシステムのためにロギング機構の選択肢を提示しましたが、一方 systemd には **systemd-journald** という組み込みのロガーが含まれています。systemd-journald サービスは、上のシステムロガーの節で概説したロギング機能のほとんどを取り扱うことができます。つまり、システムマネージャおよびサービスマネージャとして systemd を実行するシステムのほとんどでは、追加の syslog ユーティリティを追加する部分を _問題無く省略する_ ことができます。

journalctl を利用したシステムログの検索、閲覧についての詳細は、man journalctl を参照してください。

中央ホストへログを転送する場合などさまざまな理由により、systemd ベースのシステム上で _冗長な_ システムロギング機構を含めたいことがあるかもしれません。これはハンドブックの典型的想定読者にとっては一般的な事態ではなく、高度なユースケースであるとみなします。そのため、ハンドブックでは取り扱いません。

## 任意自由選択: cronデーモン

### OpenRC

cronデーモンは入れても入れなくてもよく、システムに必須ではありませんが、インストールしておくのが賢明でしょう。

cron デーモンは、予定された間隔を空けてコマンドを実行します。間隔は毎日、毎週、毎月、毎火曜日、隔週、などで設定できます。賢明なシステム管理者は、定型的なシステム管理タスクを、cron デーモンを活用して自動化するでしょう。

すべての cron デーモンは予定されたタスクのための高いレベルの粒度をサポートしており、また通常は、予定されたタスクが期待通り完了しなかった場合に、e メール等の形式で通知を送信する機能を含んています。

Gentoo ではいくつもの cron デーモンを提供しています:

- [cronie](/wiki/Cron/ja#cronie "Cron/ja") ([sys-process/cronie](https://packages.gentoo.org/packages/sys-process/cronie)) \- cronie はオリジナルの cron をベースとしていて、PAM と SELinux を使用できるなど、セキュリティと設定の改良がなされています。
- [dcron](/wiki/Cron/ja#dcron "Cron/ja") ([sys-process/dcron](https://packages.gentoo.org/packages/sys-process/dcron)) \- この軽量な cron デーモンは、利便性を損わない範囲で、簡潔で安全であることを目標としています。
- [fcron](/wiki/Cron/ja#fcron "Cron/ja") ([sys-process/fcron](https://packages.gentoo.org/packages/sys-process/fcron)) \- cron および anacron の機能を含めて拡張されたコマンドスケジューラ。
- [bcron](/wiki/Cron/ja#bcron "Cron/ja") ([sys-process/bcron](https://packages.gentoo.org/packages/sys-process/bcron)) \- 安全な操作を念頭に入れて設計された、比較的新しい cron システム。これを実現するために、システムは複数の独立したプログラムに分けられていて、それぞれが独立したタスクに責任を持ち、プログラム間の通信は厳密に制御されています。

#### デフォルト: cronie

次の例では [sys-process/cronie](https://packages.gentoo.org/packages/sys-process/cronie) を使用します:

`root #` `emerge --ask sys-process/cronie`

電源投入時に自動で cronie を開始するために、cronie を default システムランレベルに追加してください:

`root #` `rc-update add cronie default`

#### 代替案: dcron

`root #` `emerge --ask sys-process/dcron`

cron エージェントとして dcron を使う場合、初期設定のための追加コマンドが必要です。

`root #` `crontab /etc/crontab`

#### 代替案: fcron

`root #` `emerge --ask sys-process/fcron`

スケジュールされたタスクハンドラとして fcron を選択した場合、追加で emerge ステップが必要です:

`root #` `emerge --config sys-process/fcron`

#### 代替案: bcron

bcron は組み込みの特権分離を持つ、新しい cron エージェントです。

`root #` `emerge --ask sys-process/bcron`

#### modprobed-db users

If using [sys-kernel/modprobed-db](https://packages.gentoo.org/packages/sys-kernel/modprobed-db) was picked as an option for manually compiling the kernel.

Now is the time to setup a crontab to periodically scan the system for hardware used.

ファイル **`/etc/crontab`** **Run modprobed-db once every 6 hours**

```
0 */6 * * *     /usr/bin/modprobed-db store &> /dev/null

```

At a later date of at least a week, please visit the kernel build section of the [modprobed-db](/wiki/Modprobed-db#Building_kernels.2Fja "Modprobed-db") article to complete the setup.

### systemd

システムロギングと同様に、sysmted ベースのシステムには予定されたタスクのためのサポートが、 _タイマー_ という形ですぐ使える状態で含まれています。systemd タイマーはシステムレベルでもユーザレベルでも実行することができ、伝統的な cron デーモンが提供するものと同じ機能を含んでいます。冗長な可能性が必要でない限り、cron デーモンのような追加のタスクスケジューラをインストールすることは通常は不要で、問題無く省略することができます。

#### modprobed-db users

If using [sys-kernel/modprobed-db](https://packages.gentoo.org/packages/sys-kernel/modprobed-db) was picked as an option for manually compiling the kernel.

Now is the time to setup a systemd timer to periodically scan the system for hardware used.

`root #` `systemctl --user enable modprobed-db`

At a later date of at least a week, please visit the kernel build section of the [modprobed-db](/wiki/Modprobed-db#Building_kernels.2Fja "Modprobed-db") article to complete the setup.

## 任意自由選択: ファイルのインデックスを作成

より高速なファイル検索のためにファイルシステム中の各ファイルのインデックスを作成するときは、 [mlocate](/wiki/Mlocate "Mlocate") インストールしてください:

`root #` `emerge --ask sys-apps/mlocate`

## 任意自由選択: リモートシェルアクセス

**ヒント**

opensshd のデフォルト設定は、リモートユーザとして root にログインすることを許可していません。必要であれば [非 root ユーザを作成](/wiki/FAQ/ja#How_do_I_add_a_normal_user.3F "FAQ/ja") して、インストール完了後にアクセスできるように適切に設定するか、root を許可するように /etc/ssh/sshd\_config を修正してください。

インストール後、システムにリモートからアクセスできるようにするためには、ブート時に sshd を開始するように設定する必要があります。

SSH の構成に関するさらなる詳細については、 [SSH](/wiki/SSH/ja "SSH/ja") の記事を参照してください。

### OpenRC

OpenRC で sshd init スクリプトを default ランレベルに追加するには:

`root #` `rc-update add sshd default`

`root #` `rc-update add sshd default`

(たとえばリモートサーバで)シリアルコンソールからアクセスしなければならない場合、agetty を設定する必要があります。

/etc/inittab のシリアルコンソールの部分のコメントを外してください:

`root #` `nano -w /etc/inittab`

```
# SERIAL CONSOLES
s0:12345:respawn:/sbin/agetty 9600 ttyS0 vt100
s1:12345:respawn:/sbin/agetty 9600 ttyS1 vt100

```

### systemd

SSH サーバを有効化するには、次を実行してください:

`root #` `systemctl enable sshd`

`root #` `systemctl enable sshd`

シリアルコンソールサポートを有効化するには、次を実行してください:

`root #` `systemctl enable getty@tty1.service`

`root #` `systemctl enable getty@tty1.service`

## 省略可能: シェル補完

### Bash

Bash は Gentoo システムのデフォルトのシェルですので、補完拡張をインストールすることでシステムの管理を効率的かつ便利にすることができます。[app-shells/bash-completion](https://packages.gentoo.org/packages/app-shells/bash-completion) パッケージは、Gentoo に固有のコマンドや多くの共通のコマンドとユーティリティで利用できる補完をインストールします:

`root #` `emerge --ask app-shells/bash-completion`

インストール後は、各コマンドのための bash 補完を eselect を通じて管理することができます。さらなる詳細については、bash の記事の [Shell completion integrations のセクション](/wiki/Bash#Shell_completion_integrations "Bash") を参照してください。

## 推奨: 時刻同期

**重要**

Systems without a functioning [Real-Time Clock (RTC)](/wiki/System_time/ja#Software_clock_vs_Hardware_clock "System time/ja") must set the [system time](/wiki/System_time/ja "System time/ja") at every system start, and on regular intervals thereafter.

システム時刻を同期する方法を利用することは重要です。これは通常 [NTP](https://en.wikipedia.org/wiki/Network_Time_Protocol) プロトコルおよびソフトウェアによってなされます。 [Chrony](/wiki/Chrony/ja "Chrony/ja") など、NTP プロトコルを使用する実装もいくつかあります。

The clock is usually synchronized via the [Network Time Protocol](/wiki/Network_Time_Protocol "Network Time Protocol"), with an implementation such as [chrony](/wiki/Chrony/ja "Chrony/ja").

### chrony

例えば、Chrony をセットアップするには:

`root #` `emerge --ask net-misc/chrony`

`root #` `emerge --ask net-misc/chrony`

See the [chrony](/wiki/Chrony/ja "Chrony/ja") article for further information, for example if more advanced configurations are required.

### OpenRC

OpenRC では、次を実行してください:

`root #` `rc-update add chronyd default`

`root #` `rc-update add chronyd default`

### systemd

systemd では、次を実行してください:

`root #` `systemctl enable chronyd.service`

`root #` `systemctl enable chronyd.service`

### systemd-timesyncd

代わりに systemd ユーザは、デフォルトでインストールされている、よりシンプルな systemd-timesyncd NTP クライアントを利用したいかもしれません。

`root #` `systemctl enable systemd-timesyncd.service`

`root #` `systemctl enable systemd-timesyncd.service`

## ファイルシステムツール

使っているファイルシステムよって、(ファイルシステムの整合性をチェックしたり、ファイルシステムを (再) フォーマットするために) 必須のファイルシステムツールをインストールする必要があるかもしれません。ext4 ユーザ空間ツール([sys-fs/e2fsprogs](https://packages.gentoo.org/packages/sys-fs/e2fsprogs)) は [@system 集合](/wiki/System_set_(Portage) "System set (Portage)") の一部としてインストール済みであることに注意してください。

次の表は、インストールされた環境に特定のファイルシステムのツールが必要な場合、どのツールをインストールすべきかを示します。

ファイルシステム
パッケージ
[XFS](/wiki/XFS/ja "XFS/ja")[sys-fs/xfsprogs](https://packages.gentoo.org/packages/sys-fs/xfsprogs)[ext4](/wiki/Ext4/ja "Ext4/ja")[sys-fs/e2fsprogs](https://packages.gentoo.org/packages/sys-fs/e2fsprogs)[VFAT](/wiki/VFAT "VFAT") (FAT32, ...)
[sys-fs/dosfstools](https://packages.gentoo.org/packages/sys-fs/dosfstools)[Btrfs](/wiki/Btrfs/ja "Btrfs/ja")[sys-fs/btrfs-progs](https://packages.gentoo.org/packages/sys-fs/btrfs-progs)[F2FS](/wiki/F2FS/ja "F2FS/ja")[sys-fs/f2fs-tools](https://packages.gentoo.org/packages/sys-fs/f2fs-tools)[NTFS](/wiki/NTFS/ja "NTFS/ja")[sys-fs/ntfs3g](https://packages.gentoo.org/packages/sys-fs/ntfs3g)[ZFS](/wiki/ZFS "ZFS")[sys-fs/zfs](https://packages.gentoo.org/packages/sys-fs/zfs)[bcachefs](/wiki/Bcachefs "Bcachefs")[sys-fs/bcachefs-tools](https://packages.gentoo.org/packages/sys-fs/bcachefs-tools)

nvme 等のデバイスとともに適切にスケジューラを動作させるためには、[sys-block/io-scheduler-udev-rules](https://packages.gentoo.org/packages/sys-block/io-scheduler-udev-rules) をインストールするのがおすすめです:

`root #` `emerge --ask sys-block/io-scheduler-udev-rules`

`root #` `emerge --ask sys-block/io-scheduler-udev-rules`

**ヒント**

Gentooのファイルシステムについてのさらなる情報は、 [ファイルシステムの記事](/wiki/Filesystem/ja "Filesystem/ja") を参照してください。

次は [ブートローダーの設定](/wiki/Handbook:SPARC/Installation/Bootloader/ja "Handbook:SPARC/Installation/Bootloader/ja") です。

[← システムの設定](/wiki/Handbook:SPARC/Installation/System/ja "Previous part") [Home](/wiki/Handbook:SPARC/ja "Handbook:SPARC/ja") [ブートローダーの設定 →](/wiki/Handbook:SPARC/Installation/Bootloader/ja "Next part")