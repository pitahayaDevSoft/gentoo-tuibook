# Finalizing

Other languages:

- [Deutsch](/wiki/Handbook:AMD64/Installation/Finalizing/de "Handbook:AMD64/Installation/Finalizing/de (100% translated)")
- [English](/wiki/Handbook:AMD64/Installation/Finalizing "Handbook:AMD64/Installation/Finalizing (100% translated)")
- [español](/wiki/Handbook:AMD64/Installation/Finalizing/es "Handbook:AMD64/Installation/Finalizing/es (100% translated)")
- [français](/wiki/Handbook:AMD64/Installation/Finalizing/fr "Handbook:AMD64/Installation/Finalizing/fr (100% translated)")
- [italiano](/wiki/Handbook:AMD64/Installation/Finalizing/it "Handbook:AMD64/Installation/Finalizing/it (100% translated)")
- [magyar](/wiki/Handbook:AMD64/Installation/Finalizing/hu "Handbook:AMD64/Installation/Finalizing/hu (100% translated)")
- [polski](/wiki/Handbook:AMD64/Installation/Finalizing/pl "Handbook:AMD64/Installation/Finalizing/pl (100% translated)")
- [português do Brasil](/wiki/Handbook:AMD64/Installation/Finalizing/pt-br "Handbook:AMD64/Installation/Finalizing/pt-br (100% translated)")
- [čeština](/wiki/Handbook:AMD64/Installation/Finalizing/cs "Handbook:AMD64/Installation/Finalizing/cs (100% translated)")
- [русский](/wiki/Handbook:AMD64/Installation/Finalizing/ru "Handbook:AMD64/Installation/Finalizing/ru (100% translated)")
- [தமிழ்](/wiki/Handbook:AMD64/Installation/Finalizing/ta "Handbook:AMD64/Installation/Finalizing/ta (100% translated)")
- [中文（中国大陆）‎](/wiki/Handbook:AMD64/Installation/Finalizing/zh-cn "Handbook:AMD64/Installation/Finalizing/zh-cn (100% translated)")
- 日本語
- [한국어](/wiki/Handbook:AMD64/Installation/Finalizing/ko "Handbook:AMD64/Installation/Finalizing/ko (100% translated)")

[AMD64 ハンドブック](/wiki/Handbook:AMD64 "Handbook:AMD64")[インストール](/wiki/Handbook:AMD64/Full/Installation "Handbook:AMD64/Full/Installation")[インストールについて](/wiki/Handbook:AMD64/Installation/About/ja "Handbook:AMD64/Installation/About/ja")[メディアの選択](/wiki/Handbook:AMD64/Installation/Media/ja "Handbook:AMD64/Installation/Media/ja")[ネットワーク設定](/wiki/Handbook:AMD64/Installation/Networking/ja "Handbook:AMD64/Installation/Networking/ja")[ディスクの準備](/wiki/Handbook:AMD64/Installation/Disks/ja "Handbook:AMD64/Installation/Disks/ja")[stage ファイル](/wiki/Handbook:AMD64/Installation/Stage/ja "Handbook:AMD64/Installation/Stage/ja")[ベースシステムのインストール](/wiki/Handbook:AMD64/Installation/Base/ja "Handbook:AMD64/Installation/Base/ja")[カーネルの設定](/wiki/Handbook:AMD64/Installation/Kernel/ja "Handbook:AMD64/Installation/Kernel/ja")[システムの設定](/wiki/Handbook:AMD64/Installation/System/ja "Handbook:AMD64/Installation/System/ja")[ツールのインストール](/wiki/Handbook:AMD64/Installation/Tools/ja "Handbook:AMD64/Installation/Tools/ja")[ブートローダの設定](/wiki/Handbook:AMD64/Installation/Bootloader/ja "Handbook:AMD64/Installation/Bootloader/ja")締めくくり[Gentoo の操作](/wiki/Handbook:AMD64/Full/Working "Handbook:AMD64/Full/Working")[Portage について](/wiki/Handbook:AMD64/Working/Portage/ja "Handbook:AMD64/Working/Portage/ja")[USE フラグ](/wiki/Handbook:AMD64/Working/USE/ja "Handbook:AMD64/Working/USE/ja")[Portage の機能](/wiki/Handbook:AMD64/Working/Features/ja "Handbook:AMD64/Working/Features/ja")[Init スクリプトシステム](/wiki/Handbook:AMD64/Working/Initscripts/ja "Handbook:AMD64/Working/Initscripts/ja")[環境変数](/wiki/Handbook:AMD64/Working/EnvVar/ja "Handbook:AMD64/Working/EnvVar/ja")[Portage の操作](/wiki/Handbook:AMD64/Full/Portage "Handbook:AMD64/Full/Portage")[ファイルとディレクトリ](/wiki/Handbook:AMD64/Portage/Files/ja "Handbook:AMD64/Portage/Files/ja")[変数](/wiki/Handbook:AMD64/Portage/Variables/ja "Handbook:AMD64/Portage/Variables/ja")[ソフトウェアブランチの併用](/wiki/Handbook:AMD64/Portage/Branches/ja "Handbook:AMD64/Portage/Branches/ja")[追加ツール](/wiki/Handbook:AMD64/Portage/Tools/ja "Handbook:AMD64/Portage/Tools/ja")[カスタムパッケージリポジトリ](/wiki/Handbook:AMD64/Portage/CustomTree/ja "Handbook:AMD64/Portage/CustomTree/ja")[高度な機能](/wiki/Handbook:AMD64/Portage/Advanced/ja "Handbook:AMD64/Portage/Advanced/ja")[OpenRC ネットワーク設定](/wiki/Handbook:AMD64/Full/Networking "Handbook:AMD64/Full/Networking")[はじめに](/wiki/Handbook:AMD64/Networking/Introduction/ja "Handbook:AMD64/Networking/Introduction/ja")[高度な設定](/wiki/Handbook:AMD64/Networking/Advanced/ja "Handbook:AMD64/Networking/Advanced/ja")[モジュール式ネットワーク](/wiki/Handbook:AMD64/Networking/Modular/ja "Handbook:AMD64/Networking/Modular/ja")[無線](/wiki/Handbook:AMD64/Networking/Wireless/ja "Handbook:AMD64/Networking/Wireless/ja")[機能の追加](/wiki/Handbook:AMD64/Networking/Extending/ja "Handbook:AMD64/Networking/Extending/ja")[動的な管理](/wiki/Handbook:AMD64/Networking/Dynamic/ja "Handbook:AMD64/Networking/Dynamic/ja")

## Contents

- [1ユーザー管理](#.E3.83.A6.E3.83.BC.E3.82.B6.E3.83.BC.E7.AE.A1.E7.90.86)
  - [1.1日常的な使用のためのユーザを追加する](#.E6.97.A5.E5.B8.B8.E7.9A.84.E3.81.AA.E4.BD.BF.E7.94.A8.E3.81.AE.E3.81.9F.E3.82.81.E3.81.AE.E3.83.A6.E3.83.BC.E3.82.B6.E3.82.92.E8.BF.BD.E5.8A.A0.E3.81.99.E3.82.8B)
  - [1.2一時的に権限を昇格する](#.E4.B8.80.E6.99.82.E7.9A.84.E3.81.AB.E6.A8.A9.E9.99.90.E3.82.92.E6.98.87.E6.A0.BC.E3.81.99.E3.82.8B)
  - [1.3root ログインを無効化する](#root_.E3.83.AD.E3.82.B0.E3.82.A4.E3.83.B3.E3.82.92.E7.84.A1.E5.8A.B9.E5.8C.96.E3.81.99.E3.82.8B)
- [2ディスクのクリーンアップ](#.E3.83.87.E3.82.A3.E3.82.B9.E3.82.AF.E3.81.AE.E3.82.AF.E3.83.AA.E3.83.BC.E3.83.B3.E3.82.A2.E3.83.83.E3.83.97)
  - [2.1インストール用配布物の削除](#.E3.82.A4.E3.83.B3.E3.82.B9.E3.83.88.E3.83.BC.E3.83.AB.E7.94.A8.E9.85.8D.E5.B8.83.E7.89.A9.E3.81.AE.E5.89.8A.E9.99.A4)
- [3次にすることは？](#.E6.AC.A1.E3.81.AB.E3.81.99.E3.82.8B.E3.81.93.E3.81.A8.E3.81.AF.EF.BC.9F)
  - [3.1さらなるドキュメント](#.E3.81.95.E3.82.89.E3.81.AA.E3.82.8B.E3.83.89.E3.82.AD.E3.83.A5.E3.83.A1.E3.83.B3.E3.83.88)
  - [3.2Gentoo オンライン](#Gentoo_.E3.82.AA.E3.83.B3.E3.83.A9.E3.82.A4.E3.83.B3)
    - [3.2.1フォーラムと IRC](#.E3.83.95.E3.82.A9.E3.83.BC.E3.83.A9.E3.83.A0.E3.81.A8_IRC)
    - [3.2.2メーリングリスト](#.E3.83.A1.E3.83.BC.E3.83.AA.E3.83.B3.E3.82.B0.E3.83.AA.E3.82.B9.E3.83.88)
    - [3.2.3バグ](#.E3.83.90.E3.82.B0)
    - [3.2.4開発ガイド](#.E9.96.8B.E7.99.BA.E3.82.AC.E3.82.A4.E3.83.89)
- [4締めくくり](#.E7.B7.A0.E3.82.81.E3.81.8F.E3.81.8F.E3.82.8A)

## ユーザー管理

### 日常的な使用のためのユーザを追加する

Unix/Linux システム上で root として作業するのは危険であり、できるだけ避けるべきです。そのため、日常的に使用するための標準のユーザアカウントを、一つまたは複数追加することを、強くお勧めします。

グループは、そのグループに所属するメンバーができることを定義します。次の表はいくつかの重要なグループを示します。

グループ
説明
audio
ユーザアカウントがオーディオデバイスにアクセスできるようにします。
cdrom
ユーザアカウントが光学デバイスに直接アクセスできるようにします。
cron
ユーザアカウントが cron による時間ベースのジョブスケジューリングにアクセスできるようにします。メモ: systemd サービスシステムを実行するシステム上のユーザアカウントは、cron ジョブの代わりに systemd タイマとユーザサービスファイルを使用することができます。
floppy
ユーザアカウントが、フロッピードライブとして知られる、古代の機械的デバイスに直接アクセスできるようにします。このグループは現代的なシステムでは通常は使用されません。
usb
ユーザアカウントが USB デバイスにアクセスできるようにします。
video
ユーザアカウントが、ビデオキャプチャのためのハードウェアと、ハードウェアアクセラレーションにアクセスできるようにします。
wheel
ユーザアカウントが su ( _substitute user_) コマンドを使うことができるようにし、root アカウントや他のアカウントに切り換えることを許可します。root アカウントを持つシングルユーザシステムでは、メインの標準ユーザをこのグループに追加するのが良い考えでしょう。

例えば _wheel_、 _users_、 _audio_ の 3 グループに所属する [larry](/wiki/User:Larry "User:Larry") というユーザを作成するには、最初に root としてログインし (root だけがユーザを作ることができます)、useradd を実行します:

`Login:` `root`

```
Password: (rootのパスワードを入力してください)

```

標準ユーザアカウントにパスワードを設定するときに、root ユーザに設定したものと同一または似ているパスワードを使用することを避けるのは、セキュリティ上の良いプラクティスです。

ハンドブックの著者としては、少なくとも 16 文字の長さを持ち、システム上の他のどのユーザとも異なる値を持つパスワードを使用することを推奨します。

`root #` `useradd -m -G users,wheel,audio -s /bin/bash larry
`

`root #` `passwd larry`

```
Password: (larryのパスワードを入力してください)
Re-enter password: (確認のためにもう一度パスワードを入力してください)

```

### 一時的に権限を昇格する

もし、root で何か作業をする場合は、一時的に root 権限を得るために su - を使います。別の方法は [sudo](/wiki/Sudo/ja "Sudo/ja") ([app-admin/sudo](https://packages.gentoo.org/packages/app-admin/sudo)) または [doas](/wiki/Doas "Doas") ([app-admin/doas](https://packages.gentoo.org/packages/app-admin/doas)) ユーティリティを使用することです。これらは (正しく設定されれば) とても安全です。

### root ログインを無効化する

**警告**

root ログインを無効化する前に、ユーザアカウントが wheel グループのメンバであることと、ユーザの権限を昇格する方法が存在していることを確認してください; そうしないと、root へのアクセスがロックされてしまい、リカバリを実行しない限りシステム管理が不可能になるでしょう。ユーザの権限を昇格するためによく使用される方法としては、[app-admin/sudo](https://packages.gentoo.org/packages/app-admin/sudo)、[app-admin/doas](https://packages.gentoo.org/packages/app-admin/doas)、または systemd の run0 があります。

潜在的な脅威アクターが root としてログインできないようにするために、root パスワードの削除および/または root ログインの無効化によって、セキュリティを向上させることができます。

root ログインを無効化するには:

`root #` `passwd -l root`

root パスワードを削除して、ログインを無効化するには:

`root #` `passwd -dl root`

## ディスクのクリーンアップ

### インストール用配布物の削除

Gentoo のインストールとリブートが完了し、インストールがすべてうまくいっていれば、stage ファイルやその他のインストール用配布物 - DIGEST、CONTENT、\*.asc (PGP 署名) ファイルなど - は安全に削除することができます。

ファイルは / ディレクトリに置かれているので、以下のコマンドで削除することができます:

`root #` `rm /stage3-*.tar.*`

## 次にすることは？

次に何をすればいいかわかりませんか？探索する道が多くあります。Gentooには多くの可能性と共にユーザーが存在します。それゆえ、このwikiや他のGentooに関連するサブドメイン群（下の [Gentoo オンライン](/wiki/Handbook:AMD64/Installation/Finalizing/ja#Gentoo_online "Handbook:AMD64/Installation/Finalizing/ja") を見てください）を探索するため、多くのドキュメント化された機能を使うことができます（記述量は多くないのですが…）。

### さらなるドキュメント

重要な注意事項として、Gentoo で採用できる選択肢の多さから、このハンドブックが提供するドキュメントの範囲は一部に限られています。ハンドブックの主眼は、Gentoo システムを構築して基本的な管理活動を行えるようにすることに置いています。ハンドブックはあえて、グラフィカル環境、セキュリティ強化 (hardening) についての詳細、その他重要な管理タスクについての指示を除外しています。とはいえ、基本的な機能に関して読者を助けるため、ハンドブックのセクションをまだいくつか用意しています。

読者は絶対に、ハンドブックの [Gentoo の操作](/wiki/Handbook:AMD64/Working/Portage/ja "Handbook:AMD64/Working/Portage/ja") を読むべきです。これにはソフトウェアをどのように最新の状態にしておくのか、追加のソフトウェアパッケージをどのようにインストールするのか、USE フラグや Gentoo の OpenRC 初期システムの詳細や、インストール後の Gentoo システムを扱うことに関する他の様々な有益なトピックについてが説明されています。

ハンドブック以外では、コミュニティから追加提供されるドキュメントを見つけるために、Gentoo wiki の他のコーナーを探索したいと思うでしょう。Gentoo wiki チームは、 [documentation topic overview](/wiki/Main_Page#Documentation_topics "Main Page") を提供しています。これは、この Wiki にある記事のセレクションをカテゴリー別に一覧にしています。例えば、システムをよりあなたの国に適したものとするためには、 [ローカライゼーションガイド](/wiki/Localization/Guide/ja "Localization/Guide/ja") を参照します（特に英語を第二外国語として話すユーザにとって便利なものです）。

デスクトップ用途で利用したい多数派のユーザは、日常的に使用するためのグラフィカル環境をセットアップすることになるでしょう。対応している [デスクトップ環境 (DE)](/wiki/Desktop_environment/ja "Desktop environment/ja") と [ウィンドウマネージャ (WM)](/wiki/Window_manager/ja "Window manager/ja") については、コミュニティによって維持されている 'メタ' 記事が多数存在します。それぞれの DE が必要なセットアップ作業は微妙に異なっているため、ブートストラップ作業の複雑性が増えることに注意してください。

その他にも、Gentoo 上で利用可能なソフトウェアについての俯瞰的な概要を提供する [メタ記事](/wiki/Category:Meta "Category:Meta") が多数存在します。

### Gentoo オンライン

**重要**

読者は、すべての公式の Gentoo オンラインサイトが Gentoo の [行動規範](/wiki/Project:Council/Code_of_conduct "Project:Council/Code of conduct") によって治められていることに注意するべきです。Gentoo のコミュニティで活動することは特権で、権利ではありません。そしてユーザは、行動規範が理由があるために存在することを知っていると見なされます。

Libera.Chatが運営するInternet Relay Chat（IRC）ネットワークと、メーリングリストを除いて、ほとんどのGentooのウェブサイトは質問をしたり、議論を開始したり、バグを報告するために、サイト単位でアカウントを必要とします。

#### フォーラムと IRC

私たちの [Gentoo forums](https://forums.gentoo.org) や多くの [Internet Relay Chat Channels](https://www.gentoo.org/get-involved/irc-channels/) の一つに参加することはウェルカムです。新規のGentooのインストールの最中に遭遇した問題が、過去に発見されたか、またその後いくつかのフィードバックの後に解決したかをフォーラムで調べるのは簡単です。初めてのGentoo故に他のユーザがインストールの問題を経験する可能性は驚くべきものです。ユーザはGentooサポートチャンネルで手助けを求める前に、フォーラムやwikiを調べるよう勧められています。

#### メーリングリスト

フォーラムやIRCでアカウントを作成するよりはむしろ、メール上でサポートやフィードバックを求めるほうがいいというコミュニティメンバーのために、 [いくつかのメーリングリスト](https://www.gentoo.org/get-involved/mailing-lists/) が利用可能です。ユーザは特定のメーリングリストを購読するために説明に従う必要があります。

#### バグ

時々、wikiを見たり、フォーラム内を探したり、IRCチャンネルやメーリングリストにサポートを求めても、問題に対する既知の解決策がない場合があります。一般的にこれは、Gentooの [Bugzillaサイト](https://bugs.gentoo.org) にバグを公開する合図です。

#### 開発ガイド

Gentooの開発についてより多くのことを学びたいと考えている読者は、 [開発ガイド](https://devmanual.gentoo.org/) を見ると良いでしょう。このガイドにはebuildの書き方、eclassの扱い方についての説明や、Gentooの開発の背後にある、多くの [一般概念](https://devmanual.gentoo.org/general-concepts/index.html) の定義が載っています．

## 締めくくり

Gentooは堅固で、柔軟でそして素晴らしく維持されたディストリビューションです。開発者コミュニティは、Gentooをどのようにさらに _よりよい_ ディストリビューションにするかについてのフィードバックを聞けることを嬉しく思います。

確認として書いておきますが、 _このハンドブック_ に対するすべてのフィードバックは、ガイドライン（ハンドブックのはじめにある [どうやってハンドブックを改善しますか？](/wiki/Handbook:Main_Page/ja#How_do_I_improve_the_Handbook.3F "Handbook:Main Page/ja") のセクションに詳細があります）に従っているべきです。

ユーザーの皆さんが、それぞれの用途やニーズに合わせてどのように Gentoo を作り上げていくか、見られるのを楽しみにしています。

[← ブートローダーの設定](/wiki/Handbook:AMD64/Installation/Bootloader/ja "Previous part") [Home](/wiki/Handbook:AMD64/ja "Handbook:AMD64/ja") [Portageについて →](/wiki/Handbook:AMD64/Working/Portage/ja "Next part")