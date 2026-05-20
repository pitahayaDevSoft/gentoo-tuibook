# Networking

Other languages:

- [Deutsch](/wiki/Handbook:PPC/Installation/Networking/de "Handbuch:PPC/Installation/Netzwerk (100% translated)")
- [English](/wiki/Handbook:PPC/Installation/Networking "Handbook:PPC/Installation/Networking (100% translated)")
- [español](/wiki/Handbook:PPC/Installation/Networking/es "Manual de Gentoo: PPC/Instalación/Redes (100% translated)")
- [français](/wiki/Handbook:PPC/Installation/Networking/fr "Handbook:PPC/Installation/Networking/fr (100% translated)")
- [italiano](/wiki/Handbook:PPC/Installation/Networking/it "Handbook:PPC/Installation/Networking/it (100% translated)")
- [magyar](/wiki/Handbook:PPC/Installation/Networking/hu "Handbook:PPC/Installation/Networking/hu (100% translated)")
- [polski](/wiki/Handbook:PPC/Installation/Networking/pl "Handbook:PPC/Installation/Networking (100% translated)")
- [português do Brasil](/wiki/Handbook:PPC/Installation/Networking/pt-br "Handbook:PPC/Installation/Networking/pt-br (50% translated)")
- [русский](/wiki/Handbook:PPC/Installation/Networking/ru "Handbook:PPC/Installation/Networking (100% translated)")
- [தமிழ்](/wiki/Handbook:PPC/Installation/Networking/ta "கையேடு:PPC/நிறுவல்/வலையமைத்தல் (100% translated)")
- [中文（中国大陆）‎](/wiki/Handbook:PPC/Installation/Networking/zh-cn "手册：PPC/安装/配置网络 (100% translated)")
- 日本語
- [한국어](/wiki/Handbook:PPC/Installation/Networking/ko "Handbook:PPC/Installation/Networking/ko (100% translated)")

[PPC ハンドブック](/wiki/Handbook:PPC "Handbook:PPC")[インストール](/wiki/Handbook:PPC/Full/Installation "Handbook:PPC/Full/Installation")[インストールについて](/wiki/Handbook:PPC/Installation/About/ja "Handbook:PPC/Installation/About/ja")[メディアの選択](/wiki/Handbook:PPC/Installation/Media/ja "Handbook:PPC/Installation/Media/ja")ネットワーク設定[ディスクの準備](/wiki/Handbook:PPC/Installation/Disks/ja "Handbook:PPC/Installation/Disks/ja")[stage ファイル](/wiki/Handbook:PPC/Installation/Stage/ja "Handbook:PPC/Installation/Stage/ja")[ベースシステムのインストール](/wiki/Handbook:PPC/Installation/Base/ja "Handbook:PPC/Installation/Base/ja")[カーネルの設定](/wiki/Handbook:PPC/Installation/Kernel/ja "Handbook:PPC/Installation/Kernel/ja")[システムの設定](/wiki/Handbook:PPC/Installation/System/ja "Handbook:PPC/Installation/System/ja")[ツールのインストール](/wiki/Handbook:PPC/Installation/Tools/ja "Handbook:PPC/Installation/Tools/ja")[ブートローダの設定](/wiki/Handbook:PPC/Installation/Bootloader/ja "Handbook:PPC/Installation/Bootloader/ja")[締めくくり](/wiki/Handbook:PPC/Installation/Finalizing/ja "Handbook:PPC/Installation/Finalizing/ja")[Gentoo の操作](/wiki/Handbook:PPC/Full/Working "Handbook:PPC/Full/Working")[Portage について](/wiki/Handbook:PPC/Working/Portage/ja "Handbook:PPC/Working/Portage/ja")[USE フラグ](/wiki/Handbook:PPC/Working/USE/ja "Handbook:PPC/Working/USE/ja")[Portage の機能](/wiki/Handbook:PPC/Working/Features/ja "Handbook:PPC/Working/Features/ja")[Init スクリプトシステム](/wiki/Handbook:PPC/Working/Initscripts/ja "Handbook:PPC/Working/Initscripts/ja")[環境変数](/wiki/Handbook:PPC/Working/EnvVar/ja "Handbook:PPC/Working/EnvVar/ja")[Portage の操作](/wiki/Handbook:PPC/Full/Portage "Handbook:PPC/Full/Portage")[ファイルとディレクトリ](/wiki/Handbook:PPC/Portage/Files/ja "Handbook:PPC/Portage/Files/ja")[変数](/wiki/Handbook:PPC/Portage/Variables/ja "Handbook:PPC/Portage/Variables/ja")[ソフトウェアブランチの併用](/wiki/Handbook:PPC/Portage/Branches/ja "Handbook:PPC/Portage/Branches/ja")[追加ツール](/wiki/Handbook:PPC/Portage/Tools/ja "Handbook:PPC/Portage/Tools/ja")[カスタムパッケージリポジトリ](/wiki/Handbook:PPC/Portage/CustomTree/ja "Handbook:PPC/Portage/CustomTree/ja")[高度な機能](/wiki/Handbook:PPC/Portage/Advanced/ja "Handbook:PPC/Portage/Advanced/ja")[OpenRC ネットワーク設定](/wiki/Handbook:PPC/Full/Networking "Handbook:PPC/Full/Networking")[はじめに](/wiki/Handbook:PPC/Networking/Introduction/ja "Handbook:PPC/Networking/Introduction/ja")[高度な設定](/wiki/Handbook:PPC/Networking/Advanced/ja "Handbook:PPC/Networking/Advanced/ja")[モジュール式ネットワーク](/wiki/Handbook:PPC/Networking/Modular/ja "Handbook:PPC/Networking/Modular/ja")[無線](/wiki/Handbook:PPC/Networking/Wireless/ja "Handbook:PPC/Networking/Wireless/ja")[機能の追加](/wiki/Handbook:PPC/Networking/Extending/ja "Handbook:PPC/Networking/Extending/ja")[動的な管理](/wiki/Handbook:PPC/Networking/Dynamic/ja "Handbook:PPC/Networking/Dynamic/ja")

## Contents

- [1ネットワークの自動構成](#.E3.83.8D.E3.83.83.E3.83.88.E3.83.AF.E3.83.BC.E3.82.AF.E3.81.AE.E8.87.AA.E5.8B.95.E6.A7.8B.E6.88.90)
  - [1.1DHCP を使う](#DHCP_.E3.82.92.E4.BD.BF.E3.81.86)
  - [1.2ネットワークのテスト](#.E3.83.8D.E3.83.83.E3.83.88.E3.83.AF.E3.83.BC.E3.82.AF.E3.81.AE.E3.83.86.E3.82.B9.E3.83.88)
  - [1.3インターフェースの情報を取得する](#.E3.82.A4.E3.83.B3.E3.82.BF.E3.83.BC.E3.83.95.E3.82.A7.E3.83.BC.E3.82.B9.E3.81.AE.E6.83.85.E5.A0.B1.E3.82.92.E5.8F.96.E5.BE.97.E3.81.99.E3.82.8B)
- [2Wireless Setup](#Wireless_Setup)
  - [2.1Optional: Checking wireless](#Optional:_Checking_wireless)
  - [2.2Using NetworkManager](#Using_NetworkManager)
  - [2.3net-setup を使用する](#net-setup_.E3.82.92.E4.BD.BF.E7.94.A8.E3.81.99.E3.82.8B)
  - [2.4Manual setup](#Manual_setup)
    - [2.4.1Manually connecting to an open network](#Manually_connecting_to_an_open_network)
- [3アプリケーション固有の構成](#.E3.82.A2.E3.83.97.E3.83.AA.E3.82.B1.E3.83.BC.E3.82.B7.E3.83.A7.E3.83.B3.E5.9B.BA.E6.9C.89.E3.81.AE.E6.A7.8B.E6.88.90)
  - [3.1web プロキシを設定する](#web_.E3.83.97.E3.83.AD.E3.82.AD.E3.82.B7.E3.82.92.E8.A8.AD.E5.AE.9A.E3.81.99.E3.82.8B)
  - [3.2ADSL のために pppoe-setup を使う](#ADSL_.E3.81.AE.E3.81.9F.E3.82.81.E3.81.AB_pppoe-setup_.E3.82.92.E4.BD.BF.E3.81.86)
  - [3.3PPTP を使う](#PPTP_.E3.82.92.E4.BD.BF.E3.81.86)
  - [3.4インターネットと IP の基礎](#.E3.82.A4.E3.83.B3.E3.82.BF.E3.83.BC.E3.83.8D.E3.83.83.E3.83.88.E3.81.A8_IP_.E3.81.AE.E5.9F.BA.E7.A4.8E)
  - [3.5インターフェースとアドレス](#.E3.82.A4.E3.83.B3.E3.82.BF.E3.83.BC.E3.83.95.E3.82.A7.E3.83.BC.E3.82.B9.E3.81.A8.E3.82.A2.E3.83.89.E3.83.AC.E3.82.B9)
  - [3.6ネットワークと CIDR](#.E3.83.8D.E3.83.83.E3.83.88.E3.83.AF.E3.83.BC.E3.82.AF.E3.81.A8_CIDR)
  - [3.7インターネット](#.E3.82.A4.E3.83.B3.E3.82.BF.E3.83.BC.E3.83.8D.E3.83.83.E3.83.88)
  - [3.8ドメインネームシステム](#.E3.83.89.E3.83.A1.E3.82.A4.E3.83.B3.E3.83.8D.E3.83.BC.E3.83.A0.E3.82.B7.E3.82.B9.E3.83.86.E3.83.A0)
- [4手動での静的 IP ネットワーク設定](#.E6.89.8B.E5.8B.95.E3.81.A7.E3.81.AE.E9.9D.99.E7.9A.84_IP_.E3.83.8D.E3.83.83.E3.83.88.E3.83.AF.E3.83.BC.E3.82.AF.E8.A8.AD.E5.AE.9A)
  - [4.1インターフェースアドレス構成](#.E3.82.A4.E3.83.B3.E3.82.BF.E3.83.BC.E3.83.95.E3.82.A7.E3.83.BC.E3.82.B9.E3.82.A2.E3.83.89.E3.83.AC.E3.82.B9.E6.A7.8B.E6.88.90)
  - [4.2デフォルトルート構成](#.E3.83.87.E3.83.95.E3.82.A9.E3.83.AB.E3.83.88.E3.83.AB.E3.83.BC.E3.83.88.E6.A7.8B.E6.88.90)
  - [4.3DNS 構成](#DNS_.E6.A7.8B.E6.88.90)

## ネットワークの自動構成

もしかすると、もう機能していますか？

もしあなたのシステムが、IPv6 ルータか DHCP サーバを持つ Ethernet ネットワークに接続されているなら、おそらくシステムのネットワークは自動的に構成されているでしょう。追加の高度な構成が必要でない場合は、 [インターネット接続をテストすることができます](/wiki/Handbook:PPC/Installation/Networking/ja#Testing_the_network "Handbook:PPC/Installation/Networking/ja")。

### DHCP を使う

DHCP (Dynamic Host Configuration Protocol) は、ネットワークの構成を支援し、IPアドレス、ネットワークマスク、ルート、DNS サーバ、NTP サーバ等を含む、さまざまなパラメータのための構成を自動で提供することができます。

DHCP は、クライアントが _リース_ を要求するために、同一の _レイヤ 2_ ( _Ethernet_) セグメント上でサーバが実行中である必要があります。DHCP はよく RFC1918 ( _プライベート_) ネットワーク上で使用されますが、ISP からパブリック IP 情報を取得するためにも使用されます。

By default the Gentoo live media uses [NetworkManager](/wiki/NetworkManager "NetworkManager") which will automatically connect to via DHCP, if this is not the case then check the Ethernet cable for issues.

### ネットワークのテスト

_デフォルト_ ルートが適切に構成されていることは、インターネット接続の重要な構成要素です。ルート構成は以下で確認できます:

`root #` `ip route`

```
default via 192.168.0.1 dev enp1s0
```

_デフォルト_ ルートが定義されていないと、インターネット接続は利用できず、追加の構成が必要になります。

基本的なインターネット接続は ping で確認することができます:

`root #` `ping -c 3 1.1.1.1`

**ヒント**

ホスト名ではなく、既知の IP アドレスに ping することから始めるとよいでしょう。これにより、DNS の問題を他の基本的なインターネット接続の問題から切り分けることができます。

外向きの HTTPS アクセスと DNS 解決は、次で確認することができます:

`root #` `curl --location gentoo.org --output /dev/null`

curl がエラーを報告したり、他のテストが失敗したりしない場合は、インストールプロセスを [ディスクの準備](/wiki/Handbook:PPC/Installation/Disks/ja "Handbook:PPC/Installation/Disks/ja") から続けることができます。

curl がエラーを報告するのに、インターネットへの ping が機能する場合は、 [DNS を構成する必要があるかもしれません](/wiki/Handbook:PPC/Installation/Networking/ja#DNS_configuration "Handbook:PPC/Installation/Networking/ja")。

インターネット接続が確立されていない場合は、まず [インターフェース情報を検証すべきです](/wiki/Handbook:PPC/Installation/Networking/ja#Obtaining_interface_info "Handbook:PPC/Installation/Networking/ja")。そして:

- [nmtui を使用して](/wiki/Handbook:PPC/Installation/Networking/ja#Using_NetworkManager "Handbook:PPC/Installation/Networking/ja") ネットワーク構成を支援することができます。
- [アプリケーション固有の構成](/wiki/Handbook:PPC/Installation/Networking/ja#Application_specific_configuration "Handbook:PPC/Installation/Networking/ja") が必要かもしれません。
- [手動でのネットワーク構成](/wiki/Handbook:PPC/Installation/Networking/ja#Manual_network_configuration "Handbook:PPC/Installation/Networking/ja") を試すことができます。

### インターフェースの情報を取得する

そのままの状態でネットワークが機能していない場合は、インターネット接続を有効化するために追加の段階を踏む必要があります。一般的に、最初の段階はホストのネットワークインターフェースを列挙することです。

[sys-apps/iproute2](https://packages.gentoo.org/packages/sys-apps/iproute2) パッケージに含まれる ip コマンドは、システムネットワークの情報の問い合わせと、構成のために使用することができます。

**link** 引数を使用して、ネットワークインターフェースのリンクを表示することができます:

`root #` `ip link`

```
1: lo: <LOOPBACK,UP,LOWER_UP> mtu 65536 qdisc noqueue state UNKNOWN mode DEFAULT group default qlen 1000
    link/loopback 00:00:00:00:00:00 brd 00:00:00:00:00:00
4: enp1s0: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc pfifo_fast state UP mode DEFAULT group default qlen 1000
    link/ether e8:40:f2:ac:25:7a brd ff:ff:ff:ff:ff:ff
```

**address** 引数を使用して、デバイスのアドレス情報を問い合わせることができます:

`root #` `ip address`

```
2: enp1s0: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc pfifo_fast state UP group default qlen 1000<pre>
    link/ether e8:40:f2:ac:25:7a brd ff:ff:ff:ff:ff:ff
    inet 10.0.20.77/22 brd 10.0.23.255 scope global enp1s0
       valid_lft forever preferred_lft forever
    inet6 fe80::ea40:f2ff:feac:257a/64 scope link
       valid_lft forever preferred_lft forever
```

このコマンドの出力は、システム上の各ネットワークインターフェースごとの情報を含んでいます。エントリはデバイスのインデックス番号で始まり、デバイス名 (enp1s0) が続きます。

**ヒント**

**lo** ( _loopback_) 以外のインターフェースが表示されない場合は、ネットワークハードウェアに問題があるか、そのインターフェースのためのドライバがカーネルにロードされていないかです。どちらの状況も、このハンドブックの対象範囲を外れています。[#gentoo](ircs://irc.libera.chat/#gentoo) ([webchat](https://web.libera.chat/#gentoo)) に助けを求めてください。

一貫性のために、このハンドブックではメインのネットワークインターフェース名は enp1s0 であると仮定します。

**メモ**

[predictable network interface names](https://www.freedesktop.org/wiki/Software/systemd/PredictableNetworkInterfaceNames/) へ移行した結果、システム上のインターフェース名は古い命名規則による eth0 とはかなり違うものになっているかもしれません。現代的な Gentoo ブートメディアは、eno0 や ens1 や enp5s0 などで始まるインターフェース名を使用します。

## Wireless Setup

### Optional: Checking wireless

iw may be used to display the current wireless settings on the card, this should also show that the kernel wireless stack is working (note that the iw command is only available on the following architectures: **amd64**, **x86**, **arm**, **arm64**, **ppc**, **ppc64**, and **riscv**):

`root #` `iw dev wlp9s0 info`

```
Interface wlp9s0
	ifindex 3
	wdev 0x1
	addr 00:00:00:00:00:00
	type managed
	wiphy 0
	channel 11 (2462 MHz), width: 20 MHz (no HT), center1: 2462 MHz
	txpower 30.00 dBm

```

**メモ**

Some wireless cards may have a device name of wlan0 or ra0 instead of wlp9s0. Run ip link to determine the correct device name.

### Using NetworkManager

The fastest way to enable wireless connectivity in the Gentoo Live media is by using the nmtui command and following the on-screen instructions to setup the wireless network.

`root #` `nmtui`

LiveGUI users can click on the network icon in the bottom right of the taskbar to also configure WiFi.

### net-setup を使用する

自動ネットワーク構成が成功しない場合、Gentoo _ブートメディア_ はネットワーク構成を支援するスクリプトを提供しています。net-setup を使用して、無線ネットワーク情報と静的 IP を構成することができます。

`root #` `killall dhcpcd`

`root #` `net-setup enp1s0`

net-setup はネットワーク環境についていくつかの質問をし、その情報を使用して wpa\_supplicant または _静的アドレッシング_ を構成するでしょう。

**重要**

何らかの構成手順が取られた後は、ネットワークの状態を [テスト](/wiki/Handbook:PPC/Installation/Networking/ja#Testing_the_network "Handbook:PPC/Installation/Networking/ja") するべきです。構成スクリプトが機能しない場合は、 [手動でのネットワーク構成](/wiki/Handbook:PPC/Installation/Networking/ja#Manual_network_configuration "Handbook:PPC/Installation/Networking/ja") が必要です。

### Manual setup

Wireless networks may alternatively be set up with a daemon such as wpa\_supplicant or iwd. For more information on configuring wireless networking in Gentoo Linux, please read the [Wireless networking chapter](/wiki/Handbook:PPC/Networking/Wireless/ja "Handbook:PPC/Networking/Wireless/ja") in the Gentoo Handbook.}}

#### Manually connecting to an open network

For most users, there are only two settings needed to connect, the ESSID (aka wireless network name) and, optionally, the WEP key.

- First, ensure the interface is active:

`root #` `ip link set dev wlp9s0 up`

- To connect to an open network with the name _GentooNode_:

`root #` `iw dev wlp9s0 connect -w GentooNode`

## アプリケーション固有の構成

**ヒント**

These steps are provided for users where using nmtui is not able to configure their network's needs.

以下の手法は一般的に必要ではありませんが、インターネット接続のために追加の構成が必要な状況では、役に立つことがあります。

### web プロキシを設定する

web プロキシを経由してインターネットにアクセスしている場合は、Portage が対応している各プロトコルごとに正しくプロキシにアクセスできるように、プロキシ情報を定義する必要があります。Portage は wget および rsync の取得手法を介してパッケージをダウンロードするために、http\_proxy、ftp\_proxy、および RSYNC\_PROXY 環境変数を参照します。

links など、特定のテキストモード web ブラウザも、web プロキシ設定を定義する環境変数を利用することができます。特に、HTTPS アクセスのために https\_proxy 環境変数も定義する必要があるでしょう。Portage は追加の実行時パラメータを渡さなくても影響を受けますが、links にはプロキシ設定のパラメータを設定する必要があるでしょう。

ほとんどの場合、プロキシサーバのホスト名を使って環境変数を定義するだけで十分です。以下の例では、プロキシサーバのホスト名は proxy.gentoo.org、ポート番号は 8080 であるとしましょう。

**メモ**

以下のコマンド中の `#` 記号はコメントです。これらは説明のためだけに追加されているもので、コマンドを入力するときに打ち込む必要は _ありません_。

HTTP プロキシ (HTTP と HTTPS 通信のため) を定義するには:

`root #` `export http_proxy="http://proxy.gentoo.org:8080" #  Portage と Links に適用されます
`

`root #` `export https_proxy="http://proxy.gentoo.org:8080" # Links にのみ適用されます
`

HTTP プロキシが認証を必要とする場合は、次の構文でユーザ名とパスワードを設定してください:

`root #` `export http_proxy="http://username:password@proxy.gentoo.org:8080" #  Portage と Links に適用されます
`

`root #` `export https_proxy="http://username:password@proxy.gentoo.org:8080" # Links にのみ適用されます
`

プロキシサポートのためには以下のパラメータを使用して links を開始します:

`user $` `links -http-proxy ${http_proxy} -https-proxy ${https_proxy} `

Portage と links のために FTP プロキシを定義するには:

`root #` `export ftp_proxy="ftp://proxy.gentoo.org:8080" # Portage と Links に適用されます`

FTP プロキシのためには以下のパラメータを使用して links を開始します:

`user $` `links -ftp-proxy ${ftp_proxy} `

Portage のために RSYNC プロキシを定義するには:

`root #` `export RSYNC_PROXY="proxy.gentoo.org:8080" # Portage に適用されます; Links は rsync プロキシをサポートしていません`

### ADSL のために pppoe-setup を使う

インターネットアクセスのために PPPoE が必要な場合、Gentoo _ブートメディア_ には ppp 構成を簡単にするための pppoe-setup スクリプトが含まれています。

セットアップの中で、pppoe-setup は以下のことを聞いてくるでしょう:

- ADSL モデムに接続されている Ethernet _インターフェース_ の名前。
- PPPoE ユーザ名とパスワード。
- DNS サーバの IP。
- ファイアウォールが必要かどうか。

`root #` `pppoe-setup
`

`root #` `pppoe-start`

失敗した場合は、/etc/ppp/pap-secrets または /etc/ppp/chap-secrets の証明書を検証すべきです。証明書が正しい場合は、PPPoE Ethernet インターフェースの選択を確認すべきです。

### PPTP を使う

PPTP サポートが必要なら、インストール CD が提供する pptpclient を使用することができますが、使用の前に構成を行う必要があります。

/etc/ppp/pap-secrets または /etc/ppp/chap-secrets を編集して、正しいユーザ名/パスワードの組み合わせを設定してください:

`root #` `nano /etc/ppp/chap-secrets`

必要ならば/etc/ppp/options.pptpを修正してください：

`root #` `nano /etc/ppp/options.pptp`

構成が完了したら、pptp を (options.pptp で設定できないオプションがあれば、それもいっしょに付けて) 実行し、サーバに接続します:

`root #` `pptp <server ipv4 address>`

### インターネットと IP の基礎

ここまでのすべてが失敗した場合、ネットワークは手動で構成しなくてはなりません。これは特に難しくはありませんが、よく考えて行うべきです。この節は、用語を明確にし、手動でのインターネット接続に関連する基礎的なネットワークの概念を紹介するためのものです。

**ヒント**

**CPE** ( **Carrier Provided Equipment**) には、 _ルータ_、 _アクセスポイント_、 _モデム_、 _DHCP サーバ_、および _DNS サーバ_ の機能を、単一のユニットに組み合わせているものがあります。具体的な機器とデバイスの機能の区別を付けることは重要です。

### インターフェースとアドレス

ネットワーク _インターフェース_ は、ネットワークデバイスの論理的な表現です。 _インターフェース_ は、 _ネットワーク_ 上の他のデバイスと通信するために _アドレス_ を必要とします。単一の _アドレス_ だけが必要である一方で、複数のアドレスを単一の _インターフェース_ に割り当てることもできます。これはデュアルスタック (IPv4 + IPv6) 構成で特に有用です。

一貫性のために、この入門では、インターフェースは enp1s0 で、アドレス 192.168.0.2 を使用すると仮定します。

**重要**

IP アドレスは任意に設定することができます。その結果、複数のデバイスが同一の IP アドレスを使用する設定にすることができてしまいますが、これは _アドレスの競合_ につながります。アドレスの競合は DHCP または SLAAC を使用して回避すべきです。

**ヒント**

IPv6 は、アドレス構成のために典型的には **S** tate **L** ess **A** ddress **A** uto **C** onfiguration (SLAAC) を使用します。ほとんどの場合、手動で IPv6 アドレスを設定することはバッドプラクティスです。特定のアドレスサフィックスが好ましい場合は、 [インターフェース識別トークン](/wiki/IPv6_Static_Addresses_using_Tokens "IPv6 Static Addresses using Tokens") を使用することができます。

### ネットワークと CIDR

アドレスが選択された後、デバイスは、どのように他のデバイスと通信すればよいかを、どのようにして知るのでしょうか?

IP _アドレス_ は _ネットワーク_ と関連付けられています。IP _ネットワーク_ は、連続した論理的なアドレスの範囲です。

ネットワークのサイズを区別するために、 _Classless Inter-Domain Routing_、略して _CIDR_ 記法が使用されます。

- _CIDR_ 値はよく **/** で始まる表記をされ、ネットワークのサイズを表現します。

  - ネットワークのサイズを計算するためには _2 ^ (32 - CIDR)_ の公式が使用できます。
  - ネットワークのサイズが計算できたら、利用できるノード数はそこから 2 を引かなくてはなりません。
    - ネットワークの最初の IP は _ネットワークアドレス_ で、最後の IP は典型的には _ブロードキャストアドレス_ です。これらのアドレスは特別であり、通常のホストによって使用することはできません。

**ヒント**

最もよく使用される _CIDR_ 値は **/24**、および **/32** です。それぞれ **254** ノードと単一ノードを表現します。

**/24** の _CIDR_ は事実上のデフォルトのネットワークサイズです。これはサブネットマスク _255.255.255.0_ に対応し、最後の 8 ビットはネットワーク上のノードのための IP アドレスとして予約されます。

記法: 192.168.0.2/24 は次のように解釈することができます:

- _アドレス_ は 192.168.0.2
- _ネットワーク_ 192.168.0.0 上にある
- ネットワークはサイズ **254** ( _2 ^ (32 - 24) \- 2_) を持つ

  - 使用できる IP は 192.168.0.1 - 192.168.0.254 の範囲内にある
- _ブロードキャストアドレス_ 192.168.0.255 を持つ

  - ほとんどの場合、ネットワーク上の最後のアドレスは _ブロードキャストアドレス_ として使用されますが、これは変更することができます。

この構成を利用することで、デバイスは同一 _ネットワーク_ (192.168.0.0) 上のホストと通信できるようになるはずです。

### インターネット

デバイスがネットワークに参加した後、デバイスは、どのようにインターネット上のデバイスと通信すればよいかを、どのようにして知るのでしょうか?

ローカル _ネットワーク_ の外のデバイスと通信するには、 _ルーティング_ を使用しなくてはなりません。 _ルータ_ は、単純に他のデバイスのためにトラフィックを転送するネットワークデバイスです。 _デフォルトルート_ または _ゲートウェイ_ という用語は、典型的には、現在のネットワーク上で外部ネットワークへのアクセスのために使用されるデバイスを指します。

**ヒント**

_ゲートウェイ_ は、ネットワーク上の最初または最後の IP にするのが標準的です。

インターネットに接続されたルータが 192.168.0.1 で利用可能である場合、それを _デフォルトルート_ として使用して、インターネットアクセスを得ることができます。

まとめると:

- インターフェースは、 _CIDR_ 値などの、 _アドレス_ と _ネットワーク情報_ を使って構成しなくてはなりません。
- ローカルネットワークアクセスは、同一ネットワーク上の _ルータ_ へのアクセスに使用されます。
- _デフォルトルート_ が構成されたら、外部のネットワークに向けたトラフィックは _ゲートウェイ_ に転送され、インターネットへのアクセスが得られます。

### ドメインネームシステム

IP を覚えるのは困難です。 _ドメイン名_ と _IP アドレス_ の間のマッピングを行えるように、 _ドメインネームシステム_ が作成されました。

Linux システムは、 _DNS 解決_ のために使用される _ネームサーバ_ を定義するために、/etc/resolv.conf を使用します。

**ヒント**

多くの _ルータ_ は DNS サーバとしても機能します。ローカルの DNS サーアを使用することで、プライバシーを向上させ、キャッシュによって問い合わせを高速化することができます。

多くの ISP は DNS サーバを運営していて、これは通常は DHCP を介して _ゲートウェイ_ に広告されます。ローカル DNS サーバを使用することで問い合わせの遅延が改善されることが多いですが、ほとんどの公開 DNS サーバは同じ結果を返すでしょうから、サーバの使用は好みによるところが大きいです。

## 手動での静的 IP ネットワーク設定

### インターフェースアドレス構成

**重要**

手動で IP アドレスを構成するときは、ローカルネットワークトポロジを考慮する必要があります。IP アドレスは任意に設定することはできません; 衝突があるとネットワーク障害を発生させる場合があります。

enp1s0 にアドレス 192.168.0.2 と CIDR /24 を持たせるように構成するには:

`root #` `ip address add 192.168.0.2/24 dev enp1s0`

**ヒント**

このコマンドの先頭の部分は ip a と省略できます。

### デフォルトルート構成

インターフェースに対してアドレスとネットワーク情報を構成することで link ルートが構成され、そのネットワークセグメントとの通信が可能になるでしょう:

`root #` `ip route`

```
192.168.0.0/24 dev enp1s0 proto kernel scope link src 192.168.0.2
```

**ヒント**

このコマンドは ip r と省略できます。

次で default ルートを 192.168.0.1 に設定することができます:

`root #` `ip route add default via 192.168.0.1`

### DNS 構成

ネームサーバ情報は典型的には DHCP を使用して取得されますが、/etc/resolv.conf に `nameserver` エントリを追加することで、手動で設定することもできます。

**警告**

_dhcpcd_ が実行中の場合、/etc/resolv.conf への変更は保持されないでしょう。 `ps x | grep dhcpcd` で状態を確認することができます。

nano は Gentoo _ブートメディア_ に含まれており、/etc/resolv.conf を編集するために使用できます:

`root #` `nano /etc/resolv.conf`

キーワード `nameserver` に続いて DNS サーバの IP アドレスを含む行が、定義された順で問い合わせられます:

ファイル **`/etc/resolv.conf`** **Quad9 DNS を使用する。**

```
nameserver 9.9.9.9
nameserver 149.112.112.112

```

ファイル **`/etc/resolv.conf`** **Cloudflare DNS を使用する。**

```
nameserver 1.1.1.1
nameserver 1.0.0.1

```

DNS のステータスは、ドメイン名に対して ping することで確認できます:

`root #` `ping -c 3 gentoo.org`

接続が [確認できた](/wiki/Handbook:PPC/Installation/Networking/ja#Testing_the_network "Handbook:PPC/Installation/Networking/ja") ら、 [ディスクの準備](/wiki/Handbook:PPC/Installation/Disks/ja "Handbook:PPC/Installation/Disks/ja") から続けてください。

[← 適切なメディアの選択](/wiki/Handbook:PPC/Installation/Media/ja "Previous part") [Home](/wiki/Handbook:PPC/ja "Handbook:PPC/ja") [ディスクの準備 →](/wiki/Handbook:PPC/Installation/Disks/ja "Next part")