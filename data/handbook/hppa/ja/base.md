# Base

Other languages:

- [Deutsch](/wiki/Handbook:HPPA/Installation/Base/de "Handbuch:HPPA/Installation/Basis (100% translated)")
- [English](/wiki/Handbook:HPPA/Installation/Base "Handbook:HPPA/Installation/Base (100% translated)")
- [español](/wiki/Handbook:HPPA/Installation/Base/es "Manual de Gentoo: HPPA/Base (100% translated)")
- [français](/wiki/Handbook:HPPA/Installation/Base/fr "Handbook:HPPA/Installation/Base/fr (100% translated)")
- [italiano](/wiki/Handbook:HPPA/Installation/Base/it "Handbook:HPPA/Installation/Base/it (100% translated)")
- [magyar](/wiki/Handbook:HPPA/Installation/Base/hu "Handbook:HPPA/Installation/Base/hu (100% translated)")
- [polski](/wiki/Handbook:HPPA/Installation/Base/pl "Handbook:HPPA/Installation/Base (100% translated)")
- [português do Brasil](/wiki/Handbook:HPPA/Installation/Base/pt-br "Manual:HPPA/Instalação/Base (100% translated)")
- [русский](/wiki/Handbook:HPPA/Installation/Base/ru "Handbook:HPPA/Installation/Base (100% translated)")
- [தமிழ்](/wiki/Handbook:HPPA/Installation/Base/ta "கையேடு:HPPA/நிறுவல்/அடிப்படை (100% translated)")
- [中文（中国大陆）‎](/wiki/Handbook:HPPA/Installation/Base/zh-cn "手册：HPPA/安装/安装基本系统 (100% translated)")
- 日本語
- [한국어](/wiki/Handbook:HPPA/Installation/Base/ko "Handbook:HPPA/Installation/Base/ko (100% translated)")

[HPPA ハンドブック](/wiki/Handbook:HPPA "Handbook:HPPA")[インストール](/wiki/Handbook:HPPA/Full/Installation "Handbook:HPPA/Full/Installation")[インストールについて](/wiki/Handbook:HPPA/Installation/About/ja "Handbook:HPPA/Installation/About/ja")[メディアの選択](/wiki/Handbook:HPPA/Installation/Media/ja "Handbook:HPPA/Installation/Media/ja")[ネットワーク設定](/wiki/Handbook:HPPA/Installation/Networking/ja "Handbook:HPPA/Installation/Networking/ja")[ディスクの準備](/wiki/Handbook:HPPA/Installation/Disks/ja "Handbook:HPPA/Installation/Disks/ja")[stage ファイル](/wiki/Handbook:HPPA/Installation/Stage/ja "Handbook:HPPA/Installation/Stage/ja")ベースシステムのインストール[カーネルの設定](/wiki/Handbook:HPPA/Installation/Kernel/ja "Handbook:HPPA/Installation/Kernel/ja")[システムの設定](/wiki/Handbook:HPPA/Installation/System/ja "Handbook:HPPA/Installation/System/ja")[ツールのインストール](/wiki/Handbook:HPPA/Installation/Tools/ja "Handbook:HPPA/Installation/Tools/ja")[ブートローダの設定](/wiki/Handbook:HPPA/Installation/Bootloader/ja "Handbook:HPPA/Installation/Bootloader/ja")[締めくくり](/wiki/Handbook:HPPA/Installation/Finalizing/ja "Handbook:HPPA/Installation/Finalizing/ja")[Gentoo の操作](/wiki/Handbook:HPPA/Full/Working "Handbook:HPPA/Full/Working")[Portage について](/wiki/Handbook:HPPA/Working/Portage/ja "Handbook:HPPA/Working/Portage/ja")[USE フラグ](/wiki/Handbook:HPPA/Working/USE/ja "Handbook:HPPA/Working/USE/ja")[Portage の機能](/wiki/Handbook:HPPA/Working/Features/ja "Handbook:HPPA/Working/Features/ja")[Init スクリプトシステム](/wiki/Handbook:HPPA/Working/Initscripts/ja "Handbook:HPPA/Working/Initscripts/ja")[環境変数](/wiki/Handbook:HPPA/Working/EnvVar/ja "Handbook:HPPA/Working/EnvVar/ja")[Portage の操作](/wiki/Handbook:HPPA/Full/Portage "Handbook:HPPA/Full/Portage")[ファイルとディレクトリ](/wiki/Handbook:HPPA/Portage/Files/ja "Handbook:HPPA/Portage/Files/ja")[変数](/wiki/Handbook:HPPA/Portage/Variables/ja "Handbook:HPPA/Portage/Variables/ja")[ソフトウェアブランチの併用](/wiki/Handbook:HPPA/Portage/Branches/ja "Handbook:HPPA/Portage/Branches/ja")[追加ツール](/wiki/Handbook:HPPA/Portage/Tools/ja "Handbook:HPPA/Portage/Tools/ja")[カスタムパッケージリポジトリ](/wiki/Handbook:HPPA/Portage/CustomTree/ja "Handbook:HPPA/Portage/CustomTree/ja")[高度な機能](/wiki/Handbook:HPPA/Portage/Advanced/ja "Handbook:HPPA/Portage/Advanced/ja")[OpenRC ネットワーク設定](/wiki/Handbook:HPPA/Full/Networking "Handbook:HPPA/Full/Networking")[はじめに](/wiki/Handbook:HPPA/Networking/Introduction/ja "Handbook:HPPA/Networking/Introduction/ja")[高度な設定](/wiki/Handbook:HPPA/Networking/Advanced/ja "Handbook:HPPA/Networking/Advanced/ja")[モジュール式ネットワーク](/wiki/Handbook:HPPA/Networking/Modular/ja "Handbook:HPPA/Networking/Modular/ja")[無線](/wiki/Handbook:HPPA/Networking/Wireless/ja "Handbook:HPPA/Networking/Wireless/ja")[機能の追加](/wiki/Handbook:HPPA/Networking/Extending/ja "Handbook:HPPA/Networking/Extending/ja")[動的な管理](/wiki/Handbook:HPPA/Networking/Dynamic/ja "Handbook:HPPA/Networking/Dynamic/ja")

## Contents

- [1chroot する](#chroot_.E3.81.99.E3.82.8B)
  - [1.1DNS 情報をコピーする](#DNS_.E6.83.85.E5.A0.B1.E3.82.92.E3.82.B3.E3.83.94.E3.83.BC.E3.81.99.E3.82.8B)
  - [1.2必要なファイルシステムをマウントする](#.E5.BF.85.E8.A6.81.E3.81.AA.E3.83.95.E3.82.A1.E3.82.A4.E3.83.AB.E3.82.B7.E3.82.B9.E3.83.86.E3.83.A0.E3.82.92.E3.83.9E.E3.82.A6.E3.83.B3.E3.83.88.E3.81.99.E3.82.8B)
  - [1.3新しい環境に入る](#.E6.96.B0.E3.81.97.E3.81.84.E7.92.B0.E5.A2.83.E3.81.AB.E5.85.A5.E3.82.8B)
  - [1.4ブートローダのための準備をする](#.E3.83.96.E3.83.BC.E3.83.88.E3.83.AD.E3.83.BC.E3.83.80.E3.81.AE.E3.81.9F.E3.82.81.E3.81.AE.E6.BA.96.E5.82.99.E3.82.92.E3.81.99.E3.82.8B)
    - [1.4.1DOS/レガシー BIOS システム](#DOS.2F.E3.83.AC.E3.82.AC.E3.82.B7.E3.83.BC_BIOS_.E3.82.B7.E3.82.B9.E3.83.86.E3.83.A0)
- [2Portage を設定する](#Portage_.E3.82.92.E8.A8.AD.E5.AE.9A.E3.81.99.E3.82.8B)
  - [2.1Web から Gentoo ebuild リポジトリのスナップショットをインストールする](#Web_.E3.81.8B.E3.82.89_Gentoo_ebuild_.E3.83.AA.E3.83.9D.E3.82.B8.E3.83.88.E3.83.AA.E3.81.AE.E3.82.B9.E3.83.8A.E3.83.83.E3.83.97.E3.82.B7.E3.83.A7.E3.83.83.E3.83.88.E3.82.92.E3.82.A4.E3.83.B3.E3.82.B9.E3.83.88.E3.83.BC.E3.83.AB.E3.81.99.E3.82.8B)
  - [2.2任意自由選択: ミラーサーバーを選択する](#.E4.BB.BB.E6.84.8F.E8.87.AA.E7.94.B1.E9.81.B8.E6.8A.9E:_.E3.83.9F.E3.83.A9.E3.83.BC.E3.82.B5.E3.83.BC.E3.83.90.E3.83.BC.E3.82.92.E9.81.B8.E6.8A.9E.E3.81.99.E3.82.8B)
    - [2.2.1追加可能: rsync ミラー](#.E8.BF.BD.E5.8A.A0.E5.8F.AF.E8.83.BD:_rsync_.E3.83.9F.E3.83.A9.E3.83.BC)
  - [2.3任意自由選択: Gentoo ebuild リポジトリを更新する](#.E4.BB.BB.E6.84.8F.E8.87.AA.E7.94.B1.E9.81.B8.E6.8A.9E:_Gentoo_ebuild_.E3.83.AA.E3.83.9D.E3.82.B8.E3.83.88.E3.83.AA.E3.82.92.E6.9B.B4.E6.96.B0.E3.81.99.E3.82.8B)
  - [2.4ニュースを読む](#.E3.83.8B.E3.83.A5.E3.83.BC.E3.82.B9.E3.82.92.E8.AA.AD.E3.82.80)
  - [2.5適切なプロファイルを選ぶ](#.E9.81.A9.E5.88.87.E3.81.AA.E3.83.97.E3.83.AD.E3.83.95.E3.82.A1.E3.82.A4.E3.83.AB.E3.82.92.E9.81.B8.E3.81.B6)
  - [2.6追加可能: バイナリパッケージホストを追加する](#.E8.BF.BD.E5.8A.A0.E5.8F.AF.E8.83.BD:_.E3.83.90.E3.82.A4.E3.83.8A.E3.83.AA.E3.83.91.E3.83.83.E3.82.B1.E3.83.BC.E3.82.B8.E3.83.9B.E3.82.B9.E3.83.88.E3.82.92.E8.BF.BD.E5.8A.A0.E3.81.99.E3.82.8B)
    - [2.6.1リポジトリの構成](#.E3.83.AA.E3.83.9D.E3.82.B8.E3.83.88.E3.83.AA.E3.81.AE.E6.A7.8B.E6.88.90)
    - [2.6.2バイナリパッケージをインストールする](#.E3.83.90.E3.82.A4.E3.83.8A.E3.83.AA.E3.83.91.E3.83.83.E3.82.B1.E3.83.BC.E3.82.B8.E3.82.92.E3.82.A4.E3.83.B3.E3.82.B9.E3.83.88.E3.83.BC.E3.83.AB.E3.81.99.E3.82.8B)
  - [2.7追加可能: USE 変数を設定する](#.E8.BF.BD.E5.8A.A0.E5.8F.AF.E8.83.BD:_USE_.E5.A4.89.E6.95.B0.E3.82.92.E8.A8.AD.E5.AE.9A.E3.81.99.E3.82.8B)
    - [2.7.1CPU\_FLAGS\_\*](#CPU_FLAGS_.2A)
    - [2.7.2VIDEO\_CARDS](#VIDEO_CARDS)
  - [2.8追加可能: ACCEPT\_LICENSE 変数を設定する](#.E8.BF.BD.E5.8A.A0.E5.8F.AF.E8.83.BD:_ACCEPT_LICENSE_.E5.A4.89.E6.95.B0.E3.82.92.E8.A8.AD.E5.AE.9A.E3.81.99.E3.82.8B)
  - [2.9追加可能: @world 集合の更新](#.E8.BF.BD.E5.8A.A0.E5.8F.AF.E8.83.BD:_.40world_.E9.9B.86.E5.90.88.E3.81.AE.E6.9B.B4.E6.96.B0)
    - [2.9.1古いパッケージを削除する](#.E5.8F.A4.E3.81.84.E3.83.91.E3.83.83.E3.82.B1.E3.83.BC.E3.82.B8.E3.82.92.E5.89.8A.E9.99.A4.E3.81.99.E3.82.8B)
- [3タイムゾーン](#.E3.82.BF.E3.82.A4.E3.83.A0.E3.82.BE.E3.83.BC.E3.83.B3)
- [4ロケールの設定](#.E3.83.AD.E3.82.B1.E3.83.BC.E3.83.AB.E3.81.AE.E8.A8.AD.E5.AE.9A)
  - [4.1ロケールの生成](#.E3.83.AD.E3.82.B1.E3.83.BC.E3.83.AB.E3.81.AE.E7.94.9F.E6.88.90)
  - [4.2ロケールの選択](#.E3.83.AD.E3.82.B1.E3.83.BC.E3.83.AB.E3.81.AE.E9.81.B8.E6.8A.9E)
- [5参照](#.E5.8F.82.E7.85.A7)

## chroot する

### DNS 情報をコピーする

新しい環境に入る前に一つだけやるべきことが残っています。それは/etc/resolv.confに記載されているDNS情報をコピーすることです。これは新しい環境に入った後でネットワークを使うために必要です。/etc/resolv.confは、そのネットワークのネームサーバーの情報を含んでいます。

この情報をコピーするときは、cpコマンドに `--dereference` オプションを付与することを推奨します。これは/etc/resolv.confがシンボリックリンクのときに、シンボリックリンクをコピーするのではなく、シンボリックリンクのリンク先の実ファイルをコピーします。そうしないと新しい環境でシンボリックリンクが存在しないファイルを指し示すでしょう（新しい環境では、元の環境でリンク先に指定していたファイルはほぼ利用できません）。

`root #` `cp --dereference /etc/resolv.conf /mnt/gentoo/etc/`

### 必要なファイルシステムをマウントする

**ヒント**

Gentoo のインストールメディアを使用している場合は、このステップは単に arch-chroot /mnt/gentoo として置き換えることができます。

もう少しで、Linux ルートは新しい場所に変わります。

使えるようにしなければならないファイルシステムは以下の通りです。

- /proc/ は Linux カーネルから情報を引き出すための擬似ファイルシステムです。一見通常ファイルに見えますが、ファイルとしての実体はありません。
- /sys/ は /proc/ 同様、擬似ファイルシステムです。{{Path\|/proc/} }より構造化されており、一度は /proc/ を置き換えることを目的としていました。
- /dev/ は、すべてのデバイスファイルを含む通常のファイルシステムです。一部は Linux のデバイス管理機構 (通常は udev) により管理されています。
- /run/ は一時ファイルシステムです。PID ファイルやロックなど、実行時に生成されるファイルのために使用されます。

/proc/は、/mnt/gentoo/proc/にマウントされるでしょう。他はbindマウントされます。後者は、例えば/mnt/gentoo/sys/は事実/sys/となります（同じファイルシステムへの2番目のエントリです）。ここで/mnt/gentoo/proc/はファイルシステムの新しいエントリ（インスタンスとも言えるでしょう）となります。

`root #` `mount --types proc /proc /mnt/gentoo/proc
`

`root #` `mount --rbind /sys /mnt/gentoo/sys
`

`root #` `mount --make-rslave /mnt/gentoo/sys
`

`root #` `mount --rbind /dev /mnt/gentoo/dev
`

`root #` `mount --make-rslave /mnt/gentoo/dev
`

`root #` `mount --bind /run /mnt/gentoo/run
`

`root #` `mount --make-slave /mnt/gentoo/run
`

**メモ**

インストールの後半で出てくるsystemdを使う場合、 `--make-rslave` が必要です。

**警告**

Gentoo以外のインストールメディアを使う場合、これだけでは十分ではない場合があります。いくつかのディストリビューションは/run/shm/へのシンボリックリンクとして/dev/shmを作りますが、これはchroot後に無効になってしまいます。これに対応するためには、/dev/shm/をtmpfsとして適切にマウントしておくことが必要です：

`root #` `test -L /dev/shm && rm /dev/shm && mkdir /dev/shm
`

`root #` `mount --types tmpfs --options nosuid,nodev,noexec shm /dev/shm
`

そして確実にモード1777に設定してください：

`root #` `chmod 1777 /dev/shm /run/shm`

### 新しい環境に入る

ようやく、すべてのパーティションが初期化され、ベース環境がインストールされました。chroot を実行して新しいインストール環境に入りましょう。これは、セッションの root (アクセスできる最も上位レベルの場所) を、現状のインストール環境 (インストール CD もしくは他のインストールメディア) から、インストール対象システム (つまり初期化されたパーティション) に変更することを意味しています。これが _change root_ もしくは _chroot_ の意味です。

chrootは次の3ステップで実行されます。

1. chroot コマンド、または利用可能であれば arch-chroot コマンドによって、最上位ディレクトリを (インストールメディアの) / から (パーティションをマウントしている) /mnt/gentoo/ に変更する。
2. /etc/profile のいくつかの設定を source コマンドでリロードする。
3. chroot 環境であることを忘れないようするために、シェルのプロンプトを変更する。

`root #` `chroot /mnt/gentoo /bin/bash
`

`root #` `source /etc/profile
`

`root #` `export PS1="(chroot) ${PS1}"`

この時から、すべての操作は新しい Gentoo Linux 環境で実行されます。

**ヒント**

これ以降の時点で Gentoo インストールを中断しても、インストール作業をこのステップから「再開」することができる _ようになっているはず_ です。ディスクをまたパーティショニングする必要はありません！ 単純に日付と時刻を確認し、 [ルートパーティションをマウントして](/wiki/Handbook:HPPA/Installation/Disks/ja#Mounting_the_root_partition "Handbook:HPPA/Installation/Disks/ja")、上のステップを [DNS 情報をコピーする](/wiki/Handbook:HPPA/Installation/Base/ja#Copy_DNS_info "Handbook:HPPA/Installation/Base/ja") ところから実行すれば、作業中の環境に再び入ります。ブートローダの問題を解決するのにもこれが役に立ちます。さらなる情報は [chroot](/wiki/Chroot/ja "Chroot/ja") の記事にあります。

### ブートローダのための準備をする

新しい環境に入ることができたら、次はこの新しい環境をブートローダのために準備する必要があります。ブートローダをインストールするときには、正しいパーティションがマウントされていることが重要になるでしょう。

#### DOS/レガシー BIOS システム

DOS/レガシー BIOS システムでは、ブートローダは /boot ディレクトリにインストールされるので、次のようにマウントしてください:

`root #` `mount /dev/sda2 /boot`

## Portage を設定する

### Web から Gentoo ebuild リポジトリのスナップショットをインストールする

次にGentoo ebuildリポジトリのスナップショットをインストールします。このスナップショットには、インストール可能なパッケージの情報、システム管理者が選択するプロファイルの一覧、パッケージやプロファイルごとのお知らせなどをPortageに伝えるファイルが含まれます。

ここで紹介する emerge-webrsync は、HTTP/FTP プロトコル以外でのダウンロードがファイアウォールで制限されるような環境や、ネットワーク帯域を節約したい場合にお薦めです。

次のコマンドで、毎日更新される最新のスナップショットをGentooのミラーサイトから取得し、インストールします:

`root #` `emerge-webrsync`

**メモ**

この作業中、emerge-webrsync は /var/db/repos/gentoo/ がない旨のメッセージを出すかもしれません。これは想定内で、心配することはありません。このディレクトリは自動的に作成されます。

この時点で、Portageはいくつかのアップデートが推奨されていることを通知するかもしれません。これは、stageファイルでインストールされたシステム関連のパッケージについて、より新しいバージョンが利用可能であることを示しています。今回新しいリポジトリスナップショットがインストールされたことで、Portageがそれを認識したのです。このメッセージは今のところは無視して、Gentooのインストールが完了してから対応しても問題ありません。

### 任意自由選択: ミラーサーバーを選択する

ソースコードを短時間でダウンロードするために、高速で、地理的に近いミラーを選択することをおすすめします。Portage は make.conf の中の GENTOO\_MIRRORS 変数に指定されたミラー群を使用します。Gentoo のミラー一覧をブラウザで開き、インストール対象のマシンに物理的に近い一つまたは複数のミラーを選択することができます (これらは高い頻度で最も高速になり得ます)。

mirrorselect というツールが、より手っ取り早く適したミラーを検索して選択するための素晴らしいテキストインターフェースを提供しています。単に選択したいミラーにカーソルを合わせて `Spacebar` を押し、一つまたは複数のミラーを選択してください。

`root #` `emerge --ask --verbose --oneshot app-portage/mirrorselect`

`root #` `mirrorselect -i -o >> /etc/portage/make.conf`

または、アクティブなミラーの一覧は [オンラインでも確認できます](https://www.gentoo.org/downloads/mirrors/)。

#### 追加可能: rsync ミラー

Gentoo also has many rsync mirrors which can speed up downloading the available package list using `emerge --sync` (explained in more detail later) by selecting a mirror closer that is geographically closer to the user.

`root #` `mkdir /etc/portage/repos.conf
`

`root #` `cp /usr/share/portage/config/repos.conf /etc/portage/repos.conf/gentoo.conf
`

Select a mirror close to the system's location from [https://www.gentoo.org/support/rsync-mirrors/](https://www.gentoo.org/support/rsync-mirrors/) and replace the sync-uri default mirror in /etc/portage/repos.conf/gentoo.conf with the desired mirror location.

ファイル **`/etc/portage/repos.conf/gentoo.conf`** **イギリスの rsync ミラーの例**

```
[DEFAULT]
main-repo = gentoo

[gentoo]
location = /var/db/repos/gentoo
sync-type = rsync
sync-uri = rsync://rsync.uk.gentoo.org/gentoo-portage/
auto-sync = yes
sync-rsync-verify-jobs = 1
sync-rsync-verify-metamanifest = yes
sync-rsync-verify-max-age = 3
sync-openpgp-key-path = /usr/share/openpgp-keys/gentoo-release.asc
sync-openpgp-keyserver = hkps://keys.gentoo.org
sync-openpgp-key-refresh-retry-count = 40
sync-openpgp-key-refresh-retry-overall-timeout = 1200
sync-openpgp-key-refresh-retry-delay-exp-base = 2
sync-openpgp-key-refresh-retry-delay-max = 60
sync-openpgp-key-refresh-retry-delay-mult = 4
sync-webrsync-verify-signature = yes
sync-git-verify-commit-signature = true

```

### 任意自由選択: Gentoo ebuild リポジトリを更新する

Gentoo ebuildリポジトリを最新版にアップデートできます。先のemerge-webrsyncコマンドはほぼ最新の（通常は24時間以内に作成される）スナップショットをインストールするため、このステップは本当に任意です。

最新（一時間以内）のパッケージ更新があるかもしれません。その更新を取り込むためにemerge --syncを実行しましょう。このコマンドはGentoo ebuildリポジトリ（先程emerge-webrsyncコマンドで取得したもの）をアップデートするためにrsyncプロトコルを使用します。

`root #` `emerge --sync`

アップデートの時間を短縮するために、特定のフレームバッファもしくはシリアルコンソール等の遅いターミナルでは、 `--quiet` オプションを使うことをお薦めします。

`root #` `emerge --sync --quiet`

### ニュースを読む

Gentoo ebuildリポジトリの更新時、Portage が次のような情報メッセージを出すことがあります。

` * IMPORTANT: 2 news items need reading for repository 'gentoo'.
`

` * Use eselect news to read news items.
`

ニュース項目は、Gentoo ebuild リポジトリを通じて、ユーザーに重要なメッセージを通知するためのコミュニケーション手段です。これらニュース項目を管理するためにeselect newsを使用します。eselectはGentooに固有のユーティリティで、システム管理のための共通の管理インターフェースを提供します。この場合、eselectは `news` モジュールを使うことを指示されます。

`news` モジュールに対しては、主に３つの操作が使用されます。

- `list` を指定すると、現在有効なニュースアイテムの概要が表示されます。
- `read` を指定すると、そのニュースアイテムを読むことができます。
- `purge` を指定すると、一度購読したニュースを削除することができます。これにより、それらのニュースを二度と目にすることはないでしょう。

`root #` `eselect news list
`

`root #` `eselect news read`

ニュースリーダーに関するほとんどの情報はマニュアルページを通じて得ることができます。

`root #` `man news.eselect`

### 適切なプロファイルを選ぶ

**ヒント**

デスクトッププロファイルは _デスクトップ環境_ のためだけのものではありません。i3 や sway のようなミニマルなウィンドウマネージャにも適しています。

_プロファイル_ はあらゆる Gentoo システムの基礎を構成します。プロファイルは USE、CFLAGS 等の重要な変数の初期値を決めるだけではありません。プロファイルは、パッケージのバージョンを決まった範囲に固定する役目を持っています。プロファイルは Gentoo の開発者によって完全にメンテナンスされています。

現在使用中のプロファイルを確認するには、eselect を `profile` モジュールを指定して実行してください:

**ヒント**

スクロール可能な端末を持たないインストールメディアでは、 `eselect profile list | less` コマンドで、すべての利用可能なプロファイルを一覧表示し、スクロールもできる簡便な方法が提供されます。

`root #` `eselect profile list`

```
Available profile symlink targets:
  [1]   default/linux/hppa/23.0 *
  [2]   default/linux/hppa/23.0/desktop
  [3]   default/linux/hppa/23.0/desktop/gnome
  [4]   default/linux/hppa/23.0/desktop/kde

```

**メモ**

コマンドの出力は一例で、常に更新されています。

**メモ**

**systemd** を使用するには、名前に "systemd" を含んだプロファイルを選択してください。逆もまた然りです。

いくつかのアーキテクチャでは、デスクトップ体験のために共通して必要なソフトウェアパッケージを含む、デスクトップ向けのサブプロファイルが見られるでしょう。

**警告**

プロファイルのアップグレードは軽々と行われるものではありません。初期プロファイルを選択する時、確実に stage ファイルがはじめに使用していたものと **同じバージョン**（例えば 23.0）を使用してください。新しいプロファイルのバージョンは、移行方法を含むニュース項目を通して発表されます; 新しいプロファイルに移行する前にはその説明に注意して従うようにしてください。

hppaアーキテクチャで利用可能なプロファイルを確認後、別のプロファイルを選択できます。

`root #` `eselect profile set 2`

**メモ**

`developer` サブプロファイルはGentoo Linux開発向けの固有のプロファイルであり、通常のユーザーが使用するものではありません。

### 追加可能: バイナリパッケージホストを追加する

2023 年 12 月より、Gentoo の [リリースエンジニアリングチーム](/wiki/Project:RelEng "Project:RelEng") は、バイナリパッケージ (binpkg) の取得とインストールをコミュニティ全体が利用できるように、 [公式のバイナリパッケージホスト](/wiki/Binary_package_quickstart "Binary package quickstart") (略して "binhost") を提供しています。[\[1\]](#cite_note-1)

バイナリパッケージホストを追加することで、Portage は、暗号論的に署名されたコンパイル済みパッケージをインストールすることができるようになります。多くの場合、バイナリパッケージホストを追加することでパッケージインストールのための平均時間が _大幅に_ 減少するため、古い、遅い、あるいは低消費電力のシステム上で Gentoo を動かすときには、大きな利益が得られるでしょう。

#### リポジトリの構成

binhost のためのリポジトリ構成設定は Portage の /etc/portage/binrepos.conf/ ディレクトリ内に見つけられ、これは [Gentoo ebuild リポジトリ](/wiki/Handbook:HPPA/Installation/Base/ja#Gentoo_ebuild_repository "Handbook:HPPA/Installation/Base/ja") の節で言及した構成設定と同じように動作します。

バイナリホストを定義するときに検討すべき重要な側面が 2 つあります:

1. `sync-uri` の値の中のアーキテクチャおよびプロファイルターゲットは _非常に_ 重要であり、対応するコンピュータアーキテクチャ (この場合は **hppa**) と、 [適切なプロファイルを選択する](/wiki/Handbook:HPPA/Installation/Base/ja#Choosing_the_right_profile "Handbook:HPPA/Installation/Base/ja") の節で選択されたシステプロファイルに合わせるべきです。
2. 一般的に、高速で地理的に近いミラーを選択することで、取得時間が短縮されるでしょう。 [任意自由選択: ミラーサーバーを選択する](/wiki/Handbook:HPPA/Installation/Base/ja#Gentoo_ebuild_repository "Handbook:HPPA/Installation/Base/ja") の節で言及している mirrorselect ツールを確認するか、URL 値を知ることができる [オンラインのミラーリスト](https://www.gentoo.org/downloads/mirrors/) を確認してください。


ファイル **`/etc/portage/binrepos.conf/gentoo.conf`** **CDN ベースのバイナリパッケージホストの例**

```
[gentoo]
priority = 9959
# 注意: <arch> と <variant> を適切に修正する必要があります!
sync-uri = https://distfiles.gentoo.org/releases/<arch>/binpackages/<variant>
# x86-64 の sync-uri の例
# sync-uri = https://distfiles.gentoo.org/releases/amd64/binpackages/23.0/x86-64/

# リポジトリ単位での検証を行いたい場合、これは portage-3.0.74 で導入されました
verify-signature = true
# >=portage-3.0.77 でのデフォルト値
location = /var/cache/binhost/gentoo

```

#### バイナリパッケージをインストールする

Portage はデフォルトではソースコードからパッケージをコンパイルするでしょう。以下の方法で、バイナリパッケージを使用するように指示することができます:

1. emerge コマンドを呼び出すときに `--getbinpkg` オプションを渡すことができます。このバイナリパッケージインストール方法は、特定のバイナリパッケージのみをインストールするために有用です。
2. /etc/portage/make.conf ファイルを通して提示される Portage の FEATURES 変数を介して、システムのデフォルトを変更する。この設定変更を適用すると、Portage は要求されたパッケージをバイナリパッケージホストに問い合わせ、結果が見つからない場合にはローカルでコンパイルする方法にフォールバックするようになるでしょう。

例えば、Portage に利用可能なバイナリパッケージを常にインストールさせるには:

ファイル **`/etc/portage/make.conf`** **デフォルトでバイナリパッケージを使用するように Portage を設定する**

```
# FEATURES 変数内の値のリストに getbinpkg を追加してください
FEATURES="${FEATURES} getbinpkg"
# 署名を要求する
FEATURES="${FEATURES} binpkg-request-signature"

```

Portage が検証のために必要なキーリングをセットアップするために、getuto も実行してください:

`root #` `getuto`

さらなる Portage の機能についてはハンドブックの [次の章](/wiki/Handbook:HPPA/Working/Features/ja#Portage_features "Handbook:HPPA/Working/Features/ja") で取り扱います。

### 追加可能: USE 変数を設定する

USEは、Gentooがユーザに提供する最もパワフルな変数の一つです。多くのプログラムに対して、決められた追加機能を含めたり、もしくは含めずにコンパイルすることが可能です。例えば、いくつかのプログラムはGTK+サポートもしくはQtサポートを有効にしてコンパイルできます。別のプログラムにはSSLサポートを含めたり、もしくは含めずにコンパイルすることが可能です。いくつかのプログラムはX11サポート（Xサーバー）の代わりに、フレームバッファサポート（svgalib）と共にコンパイルできます。

多くのディストリビューションでは、各種のサポートを最大限含むようにコンパイルします。これはプログラムサイズと起動時間を増大させます。多くの依存関係を発生させることは言うまでもありません。Gentoo では、ユーザーはパッケージをコンパイルする時のオプションを定義できます。ここで USE が登場します。

USE変数を使って、ユーザーはコンパイルオプションにマップされるキーワードを指定します。例えば、 `ssl` キーワードはSSLをサポート可能なプログラムでSSLを有効にしてコンパイルします。 `-X` キーワードはXサーバーのサポートを含まない（最初のマイナス記号で指定）ようにコンパイルします。 `gnome gtk -kde -qt5` は、GNOME（とGTK+）サポートを有効にして、KDE（とQt）サポートを無効にします。これにより、（もし、アーキテクチャがGNOMEをサポートしていれば）システムはGNOME向けに最大限調整されます。

デフォルトの USE の設定は、システムによって使用される Gentoo プロファイルの make.defaults ファイルに記述されています。Gentoo はシステムプロファイルをサポートするために、複雑な継承システムを使用していますが、インストール作業中はこれについて深くは触れないことにします。現在有効な USE 設定を知るためのもっとも簡単な方法は、emerge --info を実行して USE で始まる行を抜き出すことです:

`root #` `emerge --info | grep ^USE`

```
USE="X acl alsa amd64 berkdb bindist bzip2 cli cracklib crypt cxx dri ..."

```

**メモ**

上記の例は出力のほとんどを省略しています。実際には、USE 変数のリストはずっとずっと長いものです。

使用可能な USE フラグの完全な説明は、/var/db/repos/gentoo/profiles/use.desc にあります。

`root #` `less /var/db/repos/gentoo/profiles/use.desc`

lessコマンドでは、 `↑` キーと `↓` キーを使ってスクロールすることができます。 `q` を押すと終了します。

例として、DVD、ALSA、CD書き込みをサポートしたKDEベースのUSE設定を示します。

`root #` `nano /etc/portage/make.conf`

ファイル **`/etc/portage/make.conf`** **DVD、ALSA、CD書き込みをサポートしたKDE/Plasmaベースのフラグ設定**

```
USE="-gtk -gnome qt5 kde dvd alsa cdr"

```

/etc/portage/make.conf で USE の値を定義すると、その値はシステムの USE フラグリストに _追加_ されます。USE フラグは、リスト中の値の前に `-` マイナス記号を追加することで、グローバルに _削除_ することができます。例えば、X グラフィカル環境のサポートを無効化するには、 `-X` を設定することでこれを行えます:

ファイル **`/etc/portage/make.conf`** **デフォルトのUSEフラグを無視する**

```
USE="-X acl alsa"

```

**警告**

`-*` を指定することで、make.conf 内で指定したもの以外のすべての USE 値を無効化することができますが、これは _まったく_ おすすめできない、賢明でないことです。Ebuild の開発者たちは、競合を防止するため、セキュリティを向上させるため、エラーを回避するため等の理由によって、デフォルトの USE フラグを選択して ebuild で指定しています。 _すべての_ USE フラグを無効化することはデフォルトの振る舞いを否定し、重大な問題を引き起こすことがあります。

#### CPU\_FLAGS\_\*

一部のアーキテクチャ (AMD64/X86、ARM、PPC を含みます) には、 [CPU\_FLAGS\_<ARCH>](/wiki/CPU_FLAGS_X86 "CPU FLAGS X86") (<ARCH> は関連するシステムアーキテクチャ名に置き換えてください) と呼ばれる [USE\_EXPAND](/wiki/USE_EXPAND "USE EXPAND") 変数があります。

**重要**

混乱しないで! **AMD64** と **X86** のシステムは共通のアーキテクチャを一部共有しているので、 **AMD64** システムのための正しい変数名は CPU\_FLAGS\_X86 です。

これは、通常手書きなどで書かれた特定のアセンブリコードや intrinsics 等を含めるようにビルドを構成するために使用されるもので、特定の CPU 機能のために最適化されたコードを出力するようにコンパイラに指示する ( `-march=` 等) のとは **異なります**。

COMMON\_FLAGS を希望に応じて設定するのに加えて、この変数も設定すべきでしょう。

これをセットアップするにはいくつかのステップが必要です:

`root #` `emerge --ask --oneshot app-portage/cpuid2cpuflags`

興味があるなら、出力を自分で確認してみてください:

`root #` `cpuid2cpuflags`

そして、出力を package.use にコピーしてください:

`root #` `echo "*/* $(cpuid2cpuflags)" > /etc/portage/package.use/00cpu-flags`

#### VIDEO\_CARDS

By default a profile already sets a few video cards. For many reasons this is not ideal, but a useful safety net.

To configure the system correctly the user needs to first unset the preset video cards with `VIDEO_CARDS: -*` then set the correct card for that system.

ファイル **`/etc/portage/package.use/00video_cards`** **例**

```
*/* VIDEO_CARDS: -* <GPU ドライバ名>

```

以下は、様々なビデオカードの種類と、それぞれに対応する `VIDEO_CARDS` の値を、簡単に比較するために使用できる表です。

GPU
VIDEO\_CARDS
(公式) Nvidia Maxwell 以降`nvidia`Nouveau Nvidia ( [サポートリスト](/wiki/NVIDIA/ja "NVIDIA/ja"))`nouveau`Sea Islands 以降の AMD`amdgpu radeonsi`ATI および古い AMD[radeon#Feature support](/wiki/Radeon#Feature_support "Radeon") を参照
Intel Nehalem 以降 ( [サポートリスト](/wiki/Intel/ja#Feature_support "Intel/ja"))`intel`Intel Gen 2 および Gen 3 ( [サポートリスト](/wiki/Intel/ja#Feature_support "Intel/ja"))`intel i915`QEMU/KVM`virgl`WSL`d3d12`

以下に _VIDEO\_CARDS_ 変数が適切に設定された package.use の例をいくつか示します。

ファイル **`/etc/portage/package.use/00video_cards`** **AMD の例**

```
*/* VIDEO_CARDS: -* amdgpu radeonsi

```

ファイル **`/etc/portage/package.use/00video_cards`** **Intel の例**

```
*/* VIDEO_CARDS: -* intel

```

ファイル **`/etc/portage/package.use/00video_cards`** **Nvidia の例**

```
*/* VIDEO_CARDS: -* nvidia

```

各種 GPU ごとの詳細は、 [AMDGPU](/wiki/AMDGPU/ja "AMDGPU/ja")、 [Intel](/wiki/Intel/ja "Intel/ja")、 [Nouveau (オープンソース)](/wiki/Nouveau/ja "Nouveau/ja")、または [NVIDIA (プロプライエタリ)](/wiki/NVIDIA/ja "NVIDIA/ja") の記事で読むことができます。

### 追加可能: ACCEPT\_LICENSE 変数を設定する

Gentoo Linux Enhancement Proposal 23 (GLEP 23) 以降、システム管理者が「インストールするソフトウェアをライセンスの観点から制限」[\[2\]](#cite_note-2)できるようにする機構が作られました。「OSI によって承認されていないソフトウェアをまったく持たないシステムを求める人もいれば、どのライセンスを暗黙に受諾しているのか単に気になるという人もいます。」[\[3\]](#cite_note-3)Gentoo システム上で実行するソフトウェアの種類をより細かく制御したいという動機から、ACCEPT\_LICENSE 変数が生まれました。

インストールのプロセス中に、Portage は、要求されたパッケージが、システム管理者による受諾できるライセンスの決定を満たしているか判断するために、ACCEPT\_LICENSE 変数に設定された値を考慮します。ここで問題があります: Gentoo ebuild リポジトリは _数千_ もの ebuild からなり、 [_数百_ もの異なるソフトウェアライセンスを持っています](https://gitweb.gentoo.org/repo/gentoo.git/tree/licenses)……。これは、新しいソフトウェアライセンスが出るたびに、システム管理者は個別にいちいち受諾を求められるということでしょうか? ありがたいことに、そうではありません; GLEP 23 はこの問題に対する解決策である、ライセンスグループの概念も規定しています。

システム管理上の利便性のために、法的に類似したソフトウェアライセンスをまとめて分類しています。ライセンスグループの定義は [ここで確認することができ](https://gitweb.gentoo.org/repo/gentoo.git/tree/profiles/license_groups)、 [Gentoo Licenses プロジェクト](/wiki/Project:Licenses "Project:Licenses") によって管理されています。ライセンスグループは構文上 `@` 記号が先頭に付けられ、一方で個々のライセンスはそうではないため、ACCEPT\_LICENSE 変数内で容易に区別が付くようになっています。

よく使用されるライセンスグループには以下のものが含まれます:

名前説明
`@GPL-COMPATIBLE`フリーソフトウェア財団によって承認された、GPLと両立するライセンス (GPL compatible license) 群 [\[a\_license 1\]](#cite_note-4)`@FSF-APPROVED`FSF (訳注: フリーソフトウェア財団) によって承認された自由ソフトウェアライセンス群 ( `@GPL-COMPATIBLE` を含みます)
`@OSI-APPROVED`Open Source Initiative によって承認されたライセンス群 [\[a\_license 2\]](#cite_note-5)`@MISC-FREE`おそらく自由ソフトウェアであるその他のライセンス群。言い換えると、自由ソフトウェアの定義[\[a\_license 3\]](#cite_note-6)に従っているものの、FSF や OSI によって承認されていないライセンス
`@FREE-SOFTWARE``@FSF-APPROVED`、 `@OSI-APPROVED`、および `@MISC-FREE` を組み合わせたもの。
`@FSF-APPROVED-OTHER`FSF が承認した「自由な文書」および「ソフトウェアや文書以外の実用の著作物のライセンス」 (フォントを含む)
`@MISC-FREE-DOCS`自由の定義[\[a\_license 4\]](#cite_note-7)に従っているが、 `@FSF-APPROVED-OTHER` の一覧には載っていない、自由な文書およびその他の著作物。
`@FREE-DOCUMENTS``@FSF-APPROVED-OTHER` および `@MISC-FREE-DOCS` を組み合わせたもの。
`@FREE`利用、共有、改変、および改変の共有の自由を持つ、すべてのライセンスからなるメタ集合。 `@FREE-SOFTWARE` および `@FREE-DOCUMENTS` を組み合わせたもの。
`@BINARY-REDISTRIBUTABLE`少なくともソフトウェアのバイナリ形式での自由な再配布を認めているライセンス群。 `@FREE` を含みます。
`@EULA`あなたの権利を取り去ろうとするライセンス契約。これらは "all-rights-reserved" や明示的な承諾を必要とするものよりも拘束的です

1. [↑](#cite_ref-4)[https://www.gnu.org/licenses/license-list.html](https://www.gnu.org/licenses/license-list.html)
2. [↑](#cite_ref-5)[https://www.opensource.org/licenses](https://www.opensource.org/licenses)
3. [↑](#cite_ref-6)[https://www.gnu.org/philosophy/free-sw.html](https://www.gnu.org/philosophy/free-sw.html)
4. [↑](#cite_ref-7)[https://freedomdefined.org/](https://freedomdefined.org/)

システム全体で現在設定されている、受諾可能なライセンスの値は、以下で確認できます:

`user $` `portageq envvar ACCEPT_LICENSE`

```
@FREE
```

この出力で分かる通り、デフォルト値は、 `@FREE` カテゴリに分類されるソフトウェアのみをインストールできるようになっています。

システム全体に関して、特定のライセンスまたはライセンスグループを以下の場所で定義することができます:

- 選択されたプロファイルによって、システム全体で \- これはデフォルト値を設定します。
- /etc/portage/make.conf ファイルによって、システム全体で。システム管理者はこのファイルで、プロファイルのデフォルト値を無視することができます。
- /etc/portage/package.license ファイルによって、パッケージ単位で。
- /etc/portage/package.license/ _ディレクトリ_ のファイルによって、パッケージ単位で。

/etc/portage/make.conf で、プロファイルによってシステム全体のデフォルトとなっているライセンス設定を無視することができます:

ファイル **`/etc/portage/make.conf`** **システム全体で ACCEPT\_LICENSE でライセンスを受諾する**

```
# プロファイルの ACCEPT_LICENSE のデフォルト値を無視する
ACCEPT_LICENSE="-* @FREE @BINARY-REDISTRIBUTABLE"

```

必要であれば、次のファイル例を含むディレクトリで示すように、パッケージごとに受諾するライセンスを定義することもできます。package.license _ディレクトリ_ が存在しない場合は作成しておく必要があります:

`root #` `mkdir /etc/portage/package.license`

各 Gentoo パッケージに対するソフトウェアライセンスの詳細は、関連する ebuild の LICENSE 変数内に保存されています。パッケージによっては複数のパッケージを持つことがあり、そのため単一のパッケージに対して、受諾できるライセンスを複数指定する必要がある場合もあります。

ファイル **`/etc/portage/package.license/kernel`** **パッケージ単位でライセンスを受諾する**

```
app-arch/unrar unRAR
sys-kernel/linux-firmware linux-fw-redistributable
sys-firmware/intel-microcode intel-ucode

```

**重要**

ebuild の LICENSE 変数は Gentoo の開発者やユーザにとってのガイドラインでしかありません。これは法的声明 _ではなく_、これが現実を反映する _保証はありません_。ソフトウェアパッケージのライセンスに関する、ebuild 開発者の解釈だけを信用しないようにすることをおすすめします; パッケージそのものを、システムにインストールされるすべてのファイルを含めて徹底的にチェックしてください。

### 追加可能: @world 集合の更新

**ヒント**

非 desktop の stage ファイルから、デスクトップ環境プロファイルターゲットを選択した場合、@world 更新プロセスはインストール時間を格段に長くしてしまうかもしれません。時間に追われている人は次の経験則が成り立つでしょう。「名前が短く、特定のシステムを示さないプロファイルの [@world 集合](/wiki/World_set_(Portage) "World set (Portage)") を選択する」。「もっとも一般的な @world 集合は、より少ないパッケージのアップデートですむ」。例えば:

- `default/linux/amd64/23.0` を選択すると、おそらくパッケージのアップデートは少なくて済むでしょう。
- `default/linux/amd64/23.0/desktop/gnome/systemd` を選択すると、おそらくより多くのパッケージがインストールされるでしょう。なぜなら、プロファイルターゲットが、 GNOME デスクトップ環境をサポートするための依存パッケージを含む、より大きな [@system](/wiki/Package_sets#.40system "Package sets") および [@profile](/wiki/Package_sets#.40profile "Package sets") 集合を持つからです。

システムの [@world 集合](/wiki/World_set_(Portage) "World set (Portage)") の更新は _省略可能_ です。以下の追加手順のいずれかまたは両方が行われていない限り、おそらく機能的な変更は発生しないでしょう:

1. プロファイルのターゲットが、stage ファイルを選択したときのものと _異なっている_。
2. インストールされたパッケージに追加の USE フラグが設定されている。

「Gentoo インストールタイムアタック」を行っているなら、@world 集合の更新は、システムを新しい Gentoo 環境内に再起動させた _後_ の時点まで安全に後回しにすることができます。

タイムアタックをしているのでなければ、パッケージ、プロファイル、現時点での USE フラグの変更に関して、Portage に更新を行わせることができます:

`root #` `emerge --ask --verbose --update --deep --changed-use @world`

[上](/wiki/Handbook:HPPA/Installation/Base/ja#Optional:_Adding_a_binary_package_host "Handbook:HPPA/Installation/Base/ja") でバイナリホストを追加した場合は、パッケージをコンパイルするのではなくバイナリホストから取得するために、--getbinpkg (または -g) を追加してもいいでしょう:

`root #` `emerge --ask --verbose --update --deep --newuse --getbinpkg @world`

#### 古いパッケージを削除する

システム更新後は、古くなったパッケージを削除するために、常に _depclean_ しておくのが重要です。削除されることになるパッケージの中に、個人的に使用していて残すべきものが無いか確認するために、 emerge --depclean --pretend の出力を注意深く確認してください。depclean されそうなパッケージを残すためには、emerge --noreplace foo を使用してください。

`root #` `emerge --ask --pretend --depclean`

問題なければ、実際の depclean を進めてください:

`root #` `emerge --ask --depclean`

## タイムゾーン

**メモ**

このステップは musl libc を使う場合は適用されません。それがどういう意味か分からない場合は、このステップを実行すべきです。

**警告**

/usr/share/zoneinfo/Etc/GMT\* のタイムゾーンは、その名前が期待されるゾーンを示していないため、避けましょう。たとえば、GMT-8 は実際には GMT+8 となります。

タイムゾーンを選択します。/usr/share/zoneinfo/から利用可能なタイムゾーンを探してください:

`root #` `ls -l /usr/share/zoneinfo`

```
total 352
drwxr-xr-x 2 root root   1120 Jan  7 17:41 Africa
drwxr-xr-x 6 root root   2960 Jan  7 17:41 America
drwxr-xr-x 2 root root    280 Jan  7 17:41 Antarctica
drwxr-xr-x 2 root root     60 Jan  7 17:41 Arctic
drwxr-xr-x 2 root root   2020 Jan  7 17:41 Asia
drwxr-xr-x 2 root root    280 Jan  7 17:41 Atlantic
drwxr-xr-x 2 root root    500 Jan  7 17:41 Australia
drwxr-xr-x 2 root root    120 Jan  7 17:41 Brazil
-rw-r--r-- 1 root root   2094 Dec  3 17:19 CET
-rw-r--r-- 1 root root   2310 Dec  3 17:19 CST6CDT
drwxr-xr-x 2 root root    200 Jan  7 17:41 Canada
drwxr-xr-x 2 root root     80 Jan  7 17:41 Chile
-rw-r--r-- 1 root root   2416 Dec  3 17:19 Cuba
-rw-r--r-- 1 root root   1908 Dec  3 17:19 EET
-rw-r--r-- 1 root root    114 Dec  3 17:19 EST
-rw-r--r-- 1 root root   2310 Dec  3 17:19 EST5EDT
-rw-r--r-- 1 root root   2399 Dec  3 17:19 Egypt
-rw-r--r-- 1 root root   3492 Dec  3 17:19 Eire
drwxr-xr-x 2 root root    740 Jan  7 17:41 Etc
drwxr-xr-x 2 root root   1320 Jan  7 17:41 Europe
...

```

`root #` `ls -l /usr/share/zoneinfo/Europe/`

```
total 256
-rw-r--r-- 1 root root 2933 Dec  3 17:19 Amsterdam
-rw-r--r-- 1 root root 1742 Dec  3 17:19 Andorra
-rw-r--r-- 1 root root 1151 Dec  3 17:19 Astrakhan
-rw-r--r-- 1 root root 2262 Dec  3 17:19 Athens
-rw-r--r-- 1 root root 3664 Dec  3 17:19 Belfast
-rw-r--r-- 1 root root 1920 Dec  3 17:19 Belgrade
-rw-r--r-- 1 root root 2298 Dec  3 17:19 Berlin
-rw-r--r-- 1 root root 2301 Dec  3 17:19 Bratislava
-rw-r--r-- 1 root root 2933 Dec  3 17:19 Brussels
...

```

選択したタイムゾーンが Europe/Brussels であるとして、このタイムゾーンを選択するためには、/etc/localtime にこの zoneinfo ファイルへのシンボリックリンクを作成することでこれを行えます:

`root #` `ln -sf ../usr/share/zoneinfo/Europe/Brussels /etc/localtime`

**ヒント**

先頭に `../` を持つターゲットパスは _リンクの場所からの相対パス_ であり、カレントディレクトリからの相対パスではありません。

**メモ**

このシンボリックリンクに絶対パスを使用することもできますが、相対リンクは systemd の timedatectl によって作成されるものでもあり、また代替 _ルート_ ディレクトリとより互換性があります。

## ロケールの設定

**メモ**

このステップは musl libc を使う場合は適用されません。それがどういう意味か分からない場合は、このステップを実行すべきです。

### ロケールの生成

Gentoo Linux をデフォルトでインストールすると、500 以上のサポートされているすべてのロケールを含むアーカイブが提供されます。ですが、管理者はこれらのうち 1 つか 2 つしか必要としないのが典型的です。その場合、/etc/locale.gen 構成ファイルに必要なロケールのリストを記載することができます。デフォルトでは locale-gen がこのファイルを読んで、指定されたロケールのみをコンパイルするため、これにより長期的に時間と空間を節約できるでしょう。

ロケールはシステムで使用する言語を指定するだけではなく、単語のソート順や日付、時間等のルールにも使用されます。ロケールは _大文字小文字を区別する_ ので、記載とまったく同じように表現する必要があります。利用可能なロケールの一覧は /usr/share/i18n/SUPPORTED ファイルで確認できます。

`root #` `nano /etc/locale.gen`

次のロケールの例では、英語（米国）とドイツ語（ドイツ）を（UTF-8のような）文字コードと共に指定しています。

ファイル **`/etc/locale.gen`** **US と DE ロケールを適切な文字コードと共に有効にする**

```
en_US.UTF-8 UTF-8
de_DE.UTF-8 UTF-8

# 非 UTF-8 ロケールは推奨されず、特殊な環境でのみ使用するべきです。
#en_US ISO-8859-1
#de_DE ISO-8859-1

```

**警告**

多くのアプリケーションは適切にビルドするのに少なくとも UTF-8 を必要とします。

次に locale-gen コマンドを実行します。このコマンドは /etc/locale.gen ファイルに記載されているすべてのロケールを生成します。

`root #` `locale-gen`

現在使用可能なすべてのロケールを確認するためには、locale -aを実行してください。

systemd を使用しているシステムでは、localectl を使用できます。localectl set-locale ... または localectl list-locales のように。

### ロケールの選択

この時点で、システム全体で有効になるロケールを設定できます。eselect を `locale` モジュールと共に使いましょう。

eselect locale listを実行すると、利用可能なターゲットが表示されます。

`root #` `eselect locale list`

```
Available targets for the LANG variable:
  [1]  C
  [2]  C.UTF-8
  [3]  POSIX
  [4]  de_DE.UTF-8
  [5]  en_US.UTF-8
  [ ]  (free form)

```

eselect locale set <番号> を実行することで、適切なロケールを選択することができます:

`root #` `eselect locale set 2`

手動で設定する場合は、/etc/env.d/02locale ファイルと、systemd の場合は /etc/locale.conf ファイルも編集してください。

ファイル **`/etc/env.d/02locale`** **システムのロケールをマニュアル設定する**

```
LANG="de_DE.UTF-8"
LC_COLLATE="C.UTF-8"

```

ロケールを設定すると、後でカーネルをビルドしたり、他のソフトをコンパイルしたりするときに警告やエラーを回避できるでしょう。

ここで、環境をリロードします。

`root #` `env-update && source /etc/profile && export PS1="(chroot) ${PS1}"`

ロケール選択プロセス全体にわたる、さらなるガイドについては、 [ローカライゼーションガイド](/wiki/Localization/Guide/ja "Localization/Guide/ja") と [UTF-8](/wiki/UTF-8/ja "UTF-8/ja") ガイドもお読みください。

## 参照

1. [↑](#cite_ref-1)[https://www.gentoo.org/news/2023/12/29/Gentoo-binary.html](https://www.gentoo.org/news/2023/12/29/Gentoo-binary.html)
2. [↑](#cite_ref-2)[https://www.gentoo.org/glep/glep-0023.html#id7](https://www.gentoo.org/glep/glep-0023.html#id7)
3. [↑](#cite_ref-3)[https://www.gentoo.org/glep/glep-0023.html#id7](https://www.gentoo.org/glep/glep-0023.html#id7)

[← Gentoo インストールファイルをインストール](/wiki/Handbook:HPPA/Installation/Stage/ja "Previous part") [Home](/wiki/Handbook:HPPA/ja "Handbook:HPPA/ja") [カーネルの設定 →](/wiki/Handbook:HPPA/Installation/Kernel/ja "Next part")