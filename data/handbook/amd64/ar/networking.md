# Networking

Other languages:

- [Deutsch](/wiki/Handbook:AMD64/Installation/Networking/de "Handbook:AMD64/Installation/Networking/de (100% translated)")
- [English](/wiki/Handbook:AMD64/Installation/Networking "Handbook:AMD64/Installation/Networking (100% translated)")
- [Türkçe](/wiki/Handbook:AMD64/Installation/Networking/tr "Handbook:AMD64/Installation/Networking/tr (0% translated)")
- [español](/wiki/Handbook:AMD64/Installation/Networking/es "Handbook:AMD64/Installation/Networking/es (100% translated)")
- [français](/wiki/Handbook:AMD64/Installation/Networking/fr "Handbook:AMD64/Installation/Networking/fr (100% translated)")
- [italiano](/wiki/Handbook:AMD64/Installation/Networking/it "Handbook:AMD64/Installation/Networking/it (100% translated)")
- [magyar](/wiki/Handbook:AMD64/Installation/Networking/hu "Handbook:AMD64/Installation/Networking/hu (100% translated)")
- [polski](/wiki/Handbook:AMD64/Installation/Networking/pl "Handbook:AMD64/Installation/Networking/pl (100% translated)")
- [português do Brasil](/wiki/Handbook:AMD64/Installation/Networking/pt-br "Handbook:AMD64/Installation/Networking/pt-br (100% translated)")
- [čeština](/wiki/Handbook:AMD64/Installation/Networking/cs "Handbook:AMD64/Installation/Networking/cs (100% translated)")
- [русский](/wiki/Handbook:AMD64/Installation/Networking/ru "Handbook:AMD64/Installation/Networking/ru (100% translated)")
- العربية
- [தமிழ்](/wiki/Handbook:AMD64/Installation/Networking/ta "Handbook:AMD64/Installation/Networking/ta (100% translated)")
- [中文（中国大陆）‎](/wiki/Handbook:AMD64/Installation/Networking/zh-cn "Handbook:AMD64/Installation/Networking/zh-cn (100% translated)")
- [日本語](/wiki/Handbook:AMD64/Installation/Networking/ja "Handbook:AMD64/Installation/Networking/ja (100% translated)")
- [한국어](/wiki/Handbook:AMD64/Installation/Networking/ko "Handbook:AMD64/Installation/Networking/ko (100% translated)")

[دليل AMD64](/wiki/Handbook:AMD64 "Handbook:AMD64")[التنصيب](/wiki/Handbook:AMD64/Full/Installation "Handbook:AMD64/Full/Installation")[عن التنصيب](/wiki/Handbook:AMD64/Installation/About/ar "Handbook:AMD64/Installation/About/ar")[اختيار وسيط التنصيب](/wiki/Handbook:AMD64/Installation/Media/ar "Handbook:AMD64/Installation/Media/ar")إعداد الشبكة[تحضير الأقراص](/wiki/Handbook:AMD64/Installation/Disks/ar "Handbook:AMD64/Installation/Disks/ar")[ملف المرحلة](/wiki/Handbook:AMD64/Installation/Stage "Handbook:AMD64/Installation/Stage")[تنصيب النظام الأساسي](/wiki/Handbook:AMD64/Installation/Base/ar "Handbook:AMD64/Installation/Base/ar")[إعداد النواة](/wiki/Handbook:AMD64/Installation/Kernel "Handbook:AMD64/Installation/Kernel")[إعداد النظام](/wiki/Handbook:AMD64/Installation/System "Handbook:AMD64/Installation/System")[تنصيب الأدوات](/wiki/Handbook:AMD64/Installation/Tools "Handbook:AMD64/Installation/Tools")[إعداد محمل الإقلاع](/wiki/Handbook:AMD64/Installation/Bootloader "Handbook:AMD64/Installation/Bootloader")[إنهاء التنصيب](/wiki/Handbook:AMD64/Installation/Finalizing "Handbook:AMD64/Installation/Finalizing")[العمل مع جنتو](/wiki/Handbook:AMD64/Full/Working "Handbook:AMD64/Full/Working")[مقدمة إلى بورتج](/wiki/Handbook:AMD64/Working/Portage "Handbook:AMD64/Working/Portage")[وسوم الاستخدام (USE)](/wiki/Handbook:AMD64/Working/USE "Handbook:AMD64/Working/USE")[ميزات بورتج](/wiki/Handbook:AMD64/Working/Features "Handbook:AMD64/Working/Features")[نظام سكربتات البدء](/wiki/Handbook:AMD64/Working/Initscripts "Handbook:AMD64/Working/Initscripts")[متغيرات البيئة](/wiki/Handbook:AMD64/Working/EnvVar "Handbook:AMD64/Working/EnvVar")[العمل مع بورتج](/wiki/Handbook:AMD64/Full/Portage "Handbook:AMD64/Full/Portage")[الملفات والدلائل](/wiki/Handbook:AMD64/Portage/Files "Handbook:AMD64/Portage/Files")[المتغيرات](/wiki/Handbook:AMD64/Portage/Variables "Handbook:AMD64/Portage/Variables")[خلط فروع البرمجيات](/wiki/Handbook:AMD64/Portage/Branches "Handbook:AMD64/Portage/Branches")[أدوات إضافية](/wiki/Handbook:AMD64/Portage/Tools "Handbook:AMD64/Portage/Tools")[مستودع حزم مخصص](/wiki/Handbook:AMD64/Portage/CustomTree "Handbook:AMD64/Portage/CustomTree")[ميزات متقدمة](/wiki/Handbook:AMD64/Portage/Advanced "Handbook:AMD64/Portage/Advanced")[إعداد الشبكة في OpenRC](/wiki/Handbook:AMD64/Full/Networking "Handbook:AMD64/Full/Networking")[البداية](/wiki/Handbook:AMD64/Networking/Introduction "Handbook:AMD64/Networking/Introduction")[الإعداد المتقدم](/wiki/Handbook:AMD64/Networking/Advanced "Handbook:AMD64/Networking/Advanced")[الشبكات التركيبية](/wiki/Handbook:AMD64/Networking/Modular "Handbook:AMD64/Networking/Modular")[اللاسلكي](/wiki/Handbook:AMD64/Networking/Wireless "Handbook:AMD64/Networking/Wireless")[إضافة الوظائف](/wiki/Handbook:AMD64/Networking/Extending "Handbook:AMD64/Networking/Extending")[الإدارة الديناميكية](/wiki/Handbook:AMD64/Networking/Dynamic "Handbook:AMD64/Networking/Dynamic")

## Contents

- [١ضبط الشبكة تلقائيًا](#.D8.B6.D8.A8.D8.B7_.D8.A7.D9.84.D8.B4.D8.A8.D9.83.D8.A9_.D8.AA.D9.84.D9.82.D8.A7.D8.A6.D9.8A.D9.8B.D8.A7)
  - [١.١استخدام DHCP](#.D8.A7.D8.B3.D8.AA.D8.AE.D8.AF.D8.A7.D9.85_DHCP)
  - [١.٢اختبار الشبكة](#.D8.A7.D8.AE.D8.AA.D8.A8.D8.A7.D8.B1_.D8.A7.D9.84.D8.B4.D8.A8.D9.83.D8.A9)
- [٢اقتناء معلومات الواجهات](#.D8.A7.D9.82.D8.AA.D9.86.D8.A7.D8.A1_.D9.85.D8.B9.D9.84.D9.88.D9.85.D8.A7.D8.AA_.D8.A7.D9.84.D9.88.D8.A7.D8.AC.D9.87.D8.A7.D8.AA)
- [٣إعداد الشبكة اللاسلكية](#.D8.A5.D8.B9.D8.AF.D8.A7.D8.AF_.D8.A7.D9.84.D8.B4.D8.A8.D9.83.D8.A9_.D8.A7.D9.84.D9.84.D8.A7.D8.B3.D9.84.D9.83.D9.8A.D8.A9)
  - [٣.١اختياري: التحقق من الاتصال اللاسلكي](#.D8.A7.D8.AE.D8.AA.D9.8A.D8.A7.D8.B1.D9.8A:_.D8.A7.D9.84.D8.AA.D8.AD.D9.82.D9.82_.D9.85.D9.86_.D8.A7.D9.84.D8.A7.D8.AA.D8.B5.D8.A7.D9.84_.D8.A7.D9.84.D9.84.D8.A7.D8.B3.D9.84.D9.83.D9.8A)
  - [٣.٢استخدام NetworkManager](#.D8.A7.D8.B3.D8.AA.D8.AE.D8.AF.D8.A7.D9.85_NetworkManager)
  - [٣.٣استخدام net-setup](#.D8.A7.D8.B3.D8.AA.D8.AE.D8.AF.D8.A7.D9.85_net-setup)
  - [٣.٤الإعداد اليدوي](#.D8.A7.D9.84.D8.A5.D8.B9.D8.AF.D8.A7.D8.AF_.D8.A7.D9.84.D9.8A.D8.AF.D9.88.D9.8A)
    - [٣.٤.١الاتصال بشبكة مفتوحة يدويًا](#.D8.A7.D9.84.D8.A7.D8.AA.D8.B5.D8.A7.D9.84_.D8.A8.D8.B4.D8.A8.D9.83.D8.A9_.D9.85.D9.81.D8.AA.D9.88.D8.AD.D8.A9_.D9.8A.D8.AF.D9.88.D9.8A.D9.8B.D8.A7)
- [٤تهيئة خاصة بالتطبيقات](#.D8.AA.D9.87.D9.8A.D8.A6.D8.A9_.D8.AE.D8.A7.D8.B5.D8.A9_.D8.A8.D8.A7.D9.84.D8.AA.D8.B7.D8.A8.D9.8A.D9.82.D8.A7.D8.AA)
  - [٤.١تهيئة وكيل وِب (بروكسي)](#.D8.AA.D9.87.D9.8A.D8.A6.D8.A9_.D9.88.D9.83.D9.8A.D9.84_.D9.88.D9.90.D8.A8_.28.D8.A8.D8.B1.D9.88.D9.83.D8.B3.D9.8A.29)
  - [٤.٢استخدام pppoe-من أجل ADSL](#.D8.A7.D8.B3.D8.AA.D8.AE.D8.AF.D8.A7.D9.85_pppoe-.D9.85.D9.86_.D8.A3.D8.AC.D9.84_ADSL)
  - [٤.٣استخدام PPTP](#.D8.A7.D8.B3.D8.AA.D8.AE.D8.AF.D8.A7.D9.85_PPTP)
- [٥أساسيات الإنترنت وعناوين IP](#.D8.A3.D8.B3.D8.A7.D8.B3.D9.8A.D8.A7.D8.AA_.D8.A7.D9.84.D8.A5.D9.86.D8.AA.D8.B1.D9.86.D8.AA_.D9.88.D8.B9.D9.86.D8.A7.D9.88.D9.8A.D9.86_IP)
  - [٥.١الواجهات والعناوين](#.D8.A7.D9.84.D9.88.D8.A7.D8.AC.D9.87.D8.A7.D8.AA_.D9.88.D8.A7.D9.84.D8.B9.D9.86.D8.A7.D9.88.D9.8A.D9.86)
  - [٥.٢الشبكات وCIDR](#.D8.A7.D9.84.D8.B4.D8.A8.D9.83.D8.A7.D8.AA_.D9.88CIDR)
  - [٥.٣الإنترنت](#.D8.A7.D9.84.D8.A5.D9.86.D8.AA.D8.B1.D9.86.D8.AA)
  - [٥.٤نظام أسماء النطاقات](#.D9.86.D8.B8.D8.A7.D9.85_.D8.A3.D8.B3.D9.85.D8.A7.D8.A1_.D8.A7.D9.84.D9.86.D8.B7.D8.A7.D9.82.D8.A7.D8.AA)
- [٦الضبط اليدوي للشبكة وIP الساكن](#.D8.A7.D9.84.D8.B6.D8.A8.D8.B7_.D8.A7.D9.84.D9.8A.D8.AF.D9.88.D9.8A_.D9.84.D9.84.D8.B4.D8.A8.D9.83.D8.A9_.D9.88IP_.D8.A7.D9.84.D8.B3.D8.A7.D9.83.D9.86)
  - [٦.١ضبط عنوان الواجهة](#.D8.B6.D8.A8.D8.B7_.D8.B9.D9.86.D9.88.D8.A7.D9.86_.D8.A7.D9.84.D9.88.D8.A7.D8.AC.D9.87.D8.A9)
  - [٦.٢ضبط المسار المبدئي](#.D8.B6.D8.A8.D8.B7_.D8.A7.D9.84.D9.85.D8.B3.D8.A7.D8.B1_.D8.A7.D9.84.D9.85.D8.A8.D8.AF.D8.A6.D9.8A)
  - [٦.٣ضبط DNS](#.D8.B6.D8.A8.D8.B7_DNS)

## ضبط الشبكة تلقائيًا

ربما سيعمل النظام دون تدخل؟

إذا كان النظام متصلاً بشبكة إيثرنت تحتوي على موجه IPv6 أو خادوم DHCP، فمن المرجح جدًا أن شبكة النظام قد هُيئت تلقائيًا. وإذا لم تكن هناك حاجة لتهيئة متقدمة إضافية، فيمكن [اختبار الاتصال بالإنترنت](/wiki/Handbook:AMD64/Installation/Networking/ar#Testing_the_network "Handbook:AMD64/Installation/Networking/ar").

### استخدام DHCP

يساعد بروتوكول تهيئة المضيف الديناميكي (DHCP) في تهيئة الشبكة، ويمكنه توفير إعدادات لمجموعة متنوعة من المعلمات تلقائيًا، بما في ذلك: عنوان IP، وقناع الشبكة، والمسارات، وخوادم DNS، وخوادم NTP، وغيرها.

يتطلب DHCP وجود خادم يعمل على نفس قطاع _الطبقة الثانية_ \-أي الإيثرنت\- مثل العميل الذي _عقد الإيجار_. غالبًا ما يُستخدم DHCP في الشبكات الخاصة (RFC1918)، ولكنه يُستخدم أيضًا للحصول على معلومات IP العامة من مزودي خدمة الإنترنت.

مبدئيًا، تستخدم وسائط جنتو الحية نظام [NetworkManager](/wiki/NetworkManager "NetworkManager") الذي سيتصل تلقائيًا عبر DHCP، وإذا لم يحدث ذلك، فتحقق من كابل الإيثرنت بحثًا عن أي مشكلات.

### اختبار الشبكة

يعد _المسار المبدئي_ المعد بشكل صحيح مكونًا حيويًا للاتصال بالإنترنت، ويمكن التحقق من تهيئة المسار عبر الأمر:

`root #` `ip route`

```
default via 192.168.0.1 dev enp1s0
```

إذا لم يعرف مسار مبدئي، فلن يتوفر الاتصال بالإنترنت، وستكون هناك حاجة لتهيئة إضافية.

يمكن التأكد من الاتصال الأساسي بالإنترنت عبر أمر ping:

`root #` `ping -c 3 1.1.1.1`

**Tip**

من المفيد البدء بتجربة ping لعنوان IP معروف بدلاً من اسم المضيف، فهذا يساعد في عزل مشاكل DNS عن مشكلات الاتصال الأساسي بالإنترنت.

يمكن التأكد من الوصول الخارجي عبر HTTPS وحل أسماء النطاقات (DNS resolution) باستخدام:

`root #` `curl --location gentoo.org --output /dev/null`

ما لم يبلغ أمر curl عن خطأ، أو تفشل الاختبارات الأخرى، يمكن الاستمرار في عملية التنصيب بالانتقال إلى [تحضير الأقراص](/wiki/Handbook:AMD64/Installation/Disks/ar "Handbook:AMD64/Installation/Disks/ar").

إذا أبلغ curl عن خطأ، بينما تعمل اختبارات ping المتجهة للإنترنت، فقد يحتاج [ال‍ DNS إلى تهيئة](/wiki/Handbook:AMD64/Installation/Networking/ar#DNS_configuration "Handbook:AMD64/Installation/Networking/ar").

إذا لم يؤسس اتصال بالإنترنت، فيجب أولاً [التحقق من معلومات الواجهة](/wiki/Handbook:AMD64/Installation/Networking/ar#Obtaining_interface_info "Handbook:AMD64/Installation/Networking/ar")، ثم:

- يمكن استخدام [أداة nmtui](/wiki/Handbook:AMD64/Installation/Networking/ar#Using_NetworkManager "Handbook:AMD64/Installation/Networking/ar") للمساعدة في تهيئة الشبكة.
- قد تكون هناك حاجة [ل‍تهيئة خاصة ببعض التطبيقات](/wiki/Handbook:AMD64/Installation/Networking/ar#Application_specific_configuration "Handbook:AMD64/Installation/Networking/ar").
- يمكن محاولة [ضبط الشبكة يدويًا](/wiki/Handbook:AMD64/Installation/Networking/ar#Manual_network_configuration "Handbook:AMD64/Installation/Networking/ar").

## اقتناء معلومات الواجهات

إذا لم تعمل الشبكة تلقائيًا، يجب اتخاذ خطوات إضافية لتفعيل الاتصال بالإنترنت. وبشكل عام، الخطوة الأولى هي سرد واجهات الشبكة في المضيف.

يمكن استخدام الأمر ip، وهو جزء من حزمة [sys-apps/iproute2](https://packages.gentoo.org/packages/sys-apps/iproute2) للاستعلام عن شبكة النظام وتهيئتها.

يمكن استخدام المعامل **link** لعرض روابط واجهات الشبكة:

`root #` `ip link`

```
1: lo: <LOOPBACK,UP,LOWER_UP> mtu 65536 qdisc noqueue state UNKNOWN mode DEFAULT group default qlen 1000
    link/loopback 00:00:00:00:00:00 brd 00:00:00:00:00:00
4: enp1s0: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc pfifo_fast state UP mode DEFAULT group default qlen 1000
    link/ether e8:40:f2:ac:25:7a brd ff:ff:ff:ff:ff:ff
```

ويمكن استخدام المعامل **address** للاستعلام عن معلومات عنوان الجهاز:

`root #` `ip address`

```
2: enp1s0: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc pfifo_fast state UP group default qlen 1000<pre>
    link/ether e8:40:f2:ac:25:7a brd ff:ff:ff:ff:ff:ff
    inet 10.0.20.77/22 brd 10.0.23.255 scope global enp1s0
       valid_lft forever preferred_lft forever
    inet6 fe80::ea40:f2ff:feac:257a/64 scope link
       valid_lft forever preferred_lft forever
```

يحتوي مخرج هذا الأمر على معلومات لكل واجهة شبكة في النظام. تبدأ الإدخالات بفهرس الجهاز، يليه اسم الجهاز: enp1s0.

**Tip**

إذا لم تظهر أي واجهات سوى الواجهة المحلية **lo** ( _loopack_)، فهذا يعني أن عتاد الشبكة به عطل، أو أن تعريف الواجهة لم يُحمَّل في النواة. وكلتا الحالتين تخرجان عن نطاق هذا الدليل، يرجى طلب الدعم عبر التواصل مع قناة [#gentoo](ircs://irc.libera.chat/#gentoo) ([webchat](https://web.libera.chat/#gentoo)).

من أجل الاتساق، سيفترض هذا الدليل أن واجهة الشبكة الأساسية تسمى enp1s0.

**Note**

نتيجة للتحول نحو [أسماء واجهات الشبكة القابلة للتنبؤ](https://www.freedesktop.org/wiki/Software/systemd/PredictableNetworkInterfaceNames/)، قد يختلف اسم الواجهة في النظام تمامًا عن اصطلاح التسمية القديم eth0. تستخدم وسائط إقلاع جنتو الحديثة أسماء واجهات ببادئات مثل eno0 أو ens1 أو enp5s0.

## إعداد الشبكة اللاسلكية

### اختياري: التحقق من الاتصال اللاسلكي

يمكن استخدام الأمر iw لعرض إعدادات الشبكة اللاسلكية الحالية على البطاقة، وهذا يوضح أيضًا أن مكدس اللاسلكي في النواة يعمل (لاحظ أن الأمر iw متاح فقط على المعماريات التالية: **amd64**، و **x86**، و **arm**، و **arm64**، و **ppc**، و **ppc64**، و **riscv**):

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

**Note**

قد يكون لبعض البطاقات اللاسلكية اسم جهاز مثل wlan0 أو ra0 بدلاً من wlp9s0. شغِّل ip link لتحديد الاسم الصحيح للجهاز.

### استخدام NetworkManager

أسرع طريقة لتفعيل الاتصال اللاسلكي في وسائط جنتو الحية هي استخدام الأمر nmtui واتباع الإرشادات التي تظهر على الشاشة لإعداد الشبكة اللاسلكية.

`root #` `nmtui`

LiveGUI users can click on the network icon in the bottom right of the taskbar to also configure WiFi.

### استخدام net-setup

في بعض المعماريات، مثل **HPPA**، قد يُطلب من المستخدم استخدام أداة net-setup لتهيئة اتصال لاسلكي إذا لم يكن NetworkManager متاحًا بعد.

`root #` `killall dhcpcd`

`root #` `net-setup enp1s0`

سيطرح net-setup بعض الأسئلة حول بيئة الشبكة وسوف يستخدم تلك المعلومات لتهيئة wpa\_supplicant أو _العنونة الساكنة_.

**Important**

يجب [اختبار حالة الشبكة](/wiki/Handbook:AMD64/Installation/Networking/ar#Testing_the_network "Handbook:AMD64/Installation/Networking/ar") بعد اتخاذ أي خطوات ضبط. وفي حال لم تنجح سكربتات الضبط، فسيلزم [الضبط اليدوي للشبكة](/wiki/Handbook:AMD64/Installation/Networking/ar#Manual_network_configuration "Handbook:AMD64/Installation/Networking/ar").

### الإعداد اليدوي

يمكن بدلاً من ذلك إعداد الشبكات اللاسلكية باستخدام عفريت مثل wpa\_supplicant أو iwd. لمزيد من المعلومات حول ضبط الشبكات اللاسلكية في جنتو لينكس، يرجى قراءة [فصل الشبكات اللاسلكية](/wiki/Handbook:AMD64/Networking/Wireless "Handbook:AMD64/Networking/Wireless") في دليل جنتو.

#### الاتصال بشبكة مفتوحة يدويًا

بالنسبة لمعظم المستخدمين، لا يلزم سوى إعدادين للاتصال: ESSID (المعروف باسم اسم الشبكة اللاسلكية) ومفتاح WEP اختياريًا.

- أولاً، تأكد من أن الواجهة نشطة:

`root #` `ip link set dev wlp9s0 up`

- للاتصال بشبكة مفتوحة باسم _GentooNode_:

`root #` `iw dev wlp9s0 connect -w GentooNode`

## تهيئة خاصة بالتطبيقات

**Tip**

تُقدم هذه الخطوات للمستخدمين الذين لا تستطيع أداة nmtui تلبية احتياجات شبكتهم.

لا تُطلب الطرق التالية بشكل عام، ولكنها قد تكون مفيدة في الحالات التي يلزم فيها ضبط إضافي للاتصال بالإنترنت.

### تهيئة وكيل وِب (بروكسي)

إذا كان الوصول إلى الإنترنت يتم عبر وكيل وِب، فمن الضروري تعريف معلومات الوكيل لكي يتمكن بورتج من الوصول إليه بشكل صحيح لكل بروتوكول مدعوم. يراقب بورتج متغيرات البيئة http\_proxy و ftp\_proxy و RSYNC\_PROXY من أجل تنزيل الحزم عبر آليات الجلب الخاصة به مثل wget و rsync.

يمكن لبعض متصفحات الوِب النصية مثل links الاستفادة أيضًا من متغيرات البيئة التي تُعرف إعدادات وكيل الوِب، وبشكل خاص للوصول عبر HTTPS، سيتطلب الأمر أيضًا تعريف متغير البيئة https\_proxy. وبينما سيتأثر بورتج دون تمرير معاملات إضافية أثناء التشغيل، فإن links سيتطلب تعيين إعدادات الوسيط.

في معظم الحالات، يكفي تعريف متغيرات البيئة باستخدام اسم مضيف الخادم. في المثال التالي، يُفترض أن خادوم الوكيل يسمى proxy.gentoo.org والمنفذ هو 8080.

**Note**

الرمز `#` في الأوامر التالية هو تعليق، وقد أُضيف للتوضيح فقط ولا يلزم كتابته عند إدخال الأوامر.

لتعريف وكيل HTTP (لحركة مرور HTTP و HTTPS):

`root #` `export http_proxy="http://proxy.gentoo.org:8080" # ينطبق على Portage و Links
`

`root #` `export https_proxy="http://proxy.gentoo.org:8080" # ينطبق فقط على Links
`

إذا كان وكيل HTTP يتطلب مصادقة، فعيِّن اسم المستخدم وكلمة السر بالصيغة التالية:

`root #` `export http_proxy="http://username:password@proxy.gentoo.org:8080" # ينطبق على Portage و Links
`

`root #` `export https_proxy="http://username:password@proxy.gentoo.org:8080" # ينطبق فقط على Links
`

ابدأ links باستخدام المعاملات التالية لدعم الوكيل:

`user $` `links -http-proxy ${http_proxy} -https-proxy ${https_proxy} `

لتعريف وكيل FTP لبورتج و\\أو links:

`root #` `export ftp_proxy="ftp://proxy.gentoo.org:8080" # ينطبق على Portage و Links`

ابدأ links باستخدام المعامل التالي لوكيل FTP:

`user $` `links -ftp-proxy ${ftp_proxy} `

لتعريف وكيل RSYNC لبورتج:

`root #` `export RSYNC_PROXY="proxy.gentoo.org:8080" # ينطبق على Portage؛ ولا يدعم Links وسيط rsync`

### استخدام pppoe-من أجل ADSL

إذا كان PPPoE مطلوبًا للوصول إلى الإنترنت، فإن _وسائط إقلاع_ جنتو تتضمن سكربت pppoe-setup لتبسيط ضبط ppp.

أثناء الإعداد، سيطلب pppoe-setup ما يلي:

- اسم واجهة الإيثرنت المتصلة بمودم ADSL.
- اسم المستخدم وكلمة السر ل‍ PPPoE.
- عناوين IP لخادوم DNS.
- ما إذا كان هناك حاجة لجدار حماية أم لا.

`root #` `pppoe-setup
`

`root #` `pppoe-start`

في حالة التعذر، يجب التحقق من بيانات الاعتماد في /etc/ppp/pap-secrets أو /etc/ppp/chap-secrets. وإذا كانت البيانات صحيحة، فيجب التحقق من اختيار واجهة إيثرنت PPPoE.

### استخدام PPTP

إذا كان دعم PPTP مطلوبًا، فيمكن استخدام pptpclient، ولكنه يتطلب ضبطًا قبل الاستخدام.

حرِّر /etc/ppp/pap-secrets أو /etc/ppp/chap-secrets بحيث يحتوي على تركيبة اسم المستخدم\\كلمة السر الصحيحة:

`root #` `nano /etc/ppp/chap-secrets`

ثم عدِّل /etc/ppp/options.pptp إذا لزم الأمر:

`root #` `nano /etc/ppp/options.pptp`

بمجرد اكتمال الضبط، شغِّل pptp (جنبًا إلى جنب مع الخيارات التي لم يمكن إعدادها في options.pptp) للاتصال بالخادوم:

`root #` `pptp <عنوان خادوم ipv4>`

## أساسيات الإنترنت وعناوين IP

إذا أخفق كل ما سبق، فيجب ضبط الشبكة يدويًا. هذا ليس صعبًا بشكل خاص، ولكن يجب القيام به بعناية. يهدف هذا القسم إلى توضيح المصطلحات وتعريف المستخدمين بمفاهيم الشبكات الأساسية المتعلقة بتهيئة اتصال الإنترنت يدويًا.

**Tip**

تجمع بعض أجهزة **CPE** (الأجهزة التي يوفرها الناقل) وظائف _الموجه_ و _نقطة الوصول_ و _المودم_ و _خادوم DHCP_ و _خادوم DNS_ في وحدة واحدة. ومن المهم التمييز بين وظائف الجهاز والجهاز المادي نفسه.

### الواجهات والعناوين

_واجهات_ الشبكة هي تمثيلات منطقية لأجهزة الشبكة. وتحتاج _الواجهة_ إلى _عنوان_ للتواصل مع الأجهزة الأخرى على _الشبكة_. وبينما يلزم عنوان واحد فقط، يمكن تخصيص عدة عناوين لواجهة واحدة. هذا مفيد بشكل خاص لتهيئات المكدس المزدوج أي (IPv4 + IPv6).

من أجل الاتساق، سيفترض هذا التمهيد أن الواجهة enp1s0 ستستخدم العنوان 192.168.0.2.

**Important**

يمكن تعيين عناوين IP بشكل عشوائي. ونتيجة لذلك، من الممكن أن تستخدم عدة أجهزة نفس عنوان IP، مما يؤدي إلى _تضارب العناوين_. ينبغي تجنب تضارب العناوين باستخدام DHCP أو SLAAC.

**Tip**

يستخدم IPv6 عادةً «الضبط التلقائي للعناوين عديمة الحالة» (SLAAC) لضبط العناوين. وفي معظم الحالات، يعد إعداد عناوين IPv6 يدويًا ممارسة سيئة. وإذا كنت تفضل لاحقة عنوان محددة، فيمكن استخدام [رموز تعريف الواجهة](/wiki/IPv6_Static_Addresses_using_Tokens "IPv6 Static Addresses using Tokens").

### الشبكات وCIDR

بمجرد اختيار العنوان، كيف يعرف الجهاز كيفية التحدث مع الأجهزة الأخرى؟

ترتبط عناوين IP ب _شبكات_. _شبكات_ IP هي نطاقات منطقية متصلة من العناوين.

يُستخدم تدوين _التوجيه بين النطاقي غير الصنفي_ (CIDR) لتمييز أحجام الشبكات.

- تمثل قيمة CIDR، التي غالبًا ما تُكتب مبتدئة ب‍ **/**، حجم الشبكة.

  - يمكن استخدام الصيغة _2 ^ (32 - CIDR)_ لحساب حجم الشبكة.
  - بمجرد حساب حجم الشبكة، يجب تقليل عدد العقد القابلة للاستخدام بمقدار 2.
    - عنوان IP الأول في الشبكة هو _عنوان الشبكة_، والأخير هو عادةً _عنوان البث_. هذه العناوين خاصة ولا يمكن استخدامها من قبل المضيفين العاديين.

**Tip**

أكثر قيم CIDR شيوعًا هي **/24** و **/32**، وتمثلان **254** عقدة وعقدة واحدة على التوالي.

تعتبر قيمة CIDR بقيمة **/24** هي الحجم المبدئي الفعلي للشبكة. وهذا يتوافق مع قناع شبكة فرعية قدره _255.255.255.0_، حيث تحجز آخر 8 بتات لعناوين IP للعقد الموجودة في الشبكة.

يمكن تفسير التدوين: 192.168.0.2/24 كالتالي:

- العنوان 192.168.0.2
- على الشبكة 192.168.0.0
- بحجم **254** ( _2 ^ (32 - 24) \- 2_)

  - عناوين IP القابلة للاستخدام تقع في النطاق 192.168.0.1 - 192.168.0.254
- مع _عنوان بث_ قدره 192.168.0.255

  - في معظم الحالات، يُستخدم العنوان الأخير في الشبكة كعنوان بث، ولكن يمكن تغيير ذلك.

باستخدام هذا الضبط، يجب أن يكون الجهاز قادرًا على التواصل مع أي مضيف على نفس الشبكة (192.168.0.0).

### الإنترنت

بمجرد وجود الجهاز على الشبكة، كيف يعرف كيفية التحدث مع الأجهزة على الإنترنت؟

للتواصل مع الأجهزة خارج الشبكات المحلية، يجب استخدام _التوجيه_. والموجه ببساطة هو جهاز شبكة يقوم بتمرير حركة المرور للأجهزة الأخرى. وعادة ما يشير مصطلح _المسار المبدئي_ أو _البوابة_ إلى أي جهاز في الشبكة الحالية يُستخدم للوصول إلى الشبكة الخارجية.

**Tip**

من الممارسات القياسية جعل _البوابة_ هي أول أو آخر عنوان IP في الشبكة.

إذا توفر موجه متصل بالإنترنت عند العنوان 192.168.0.1، فيمكن استخدامه كمسار مبدئي، مما يمنح الوصول إلى الإنترنت.

للتلخيص:

- يجب ضبط الواجهات ب _عنوان_ و _معلومات الشبكة_، مثل قيمة CIDR.
- يُستخدم الوصول إلى الشبكة المحلية للوصول إلى _موجه_ على نفس الشبكة.
- يُضبط _المسار المبدئي_، بحيث يتم تمرير حركة المرور المتجهة إلى الشبكات الخارجية إلى _البوابة_، مما يوفر الوصول إلى الإنترنت.

### نظام أسماء النطاقات

تذكر عناوين IP أمر صعب، لذا أُنشئ _نظام أسماء النطاقات_ للسماح بالربط بين _أسماء النطاقات_ و _عناوين IP_.

تستخدم نظم لينكس الملف /etc/resolv.conf لتعريف _خوادم الأسماء" التي ستُستخدم ل_ حل نظام أسماء النطاقات _._

**Tip**

يمكن للعديد من _الموجهات_ أن تعمل أيضًا كخادوم DNS، واستخدام خادوم DNS محلي يمكن أن يعزز الخصوصية ويسرع الاستعلامات من خلال الاختزان.

يدير العديد من مزودي خدمة الإنترنت خادوم DNS يتم الإعلان عنه عادةً ل _البوابة_ عبر DHCP. يميل استخدام خادوم DNS محلي إلى تحسين وقت الاستجابة، ولكن معظم خوادم DNS العامة ستعيد نفس النتائج، لذا يعتمد استخدام الخادوم إلى حد كبير على التفضيل الشخصي.

## الضبط اليدوي للشبكة وIP الساكن

### ضبط عنوان الواجهة

**Important**

عند ضبط عناوين IP يدويًا، يجب مراعاة طوبولوجيا الشبكة المحلية. يمكن تعيين عناوين IP بشكل عشوائي، وقد يتسبب التضارب في انقطاع الشبكة.

لضبط enp1s0 بالعنوان 192.168.0.2 و CIDR بقيمة /24:

`root #` `ip address add 192.168.0.2/24 dev enp1s0`

**Tip**

يمكن اختصار بداية هذا الأمر إلى ip a.

### ضبط المسار المبدئي

سيؤدي ضبط معلومات العنوان والشبكة لواجهة ما إلى تهيئة مسارات الرابط (link)، مما يسمح بالاتصال مع قطاع الشبكة ذلك:

`root #` `ip route`

```
192.168.0.0/24 dev enp1s0 proto kernel scope link src 192.168.0.2
```

**Tip**

يمكن اختصار هذا الأمر إلى ip r.

يمكن تعيين المسار المبدئي (default) إلى 192.168.0.1 عبر:

`root #` `ip route add default via 192.168.0.1`

### ضبط DNS

يُحصل على معلومات خادوم الأسماء عادةً باستخدام DHCP، ولكن يمكن ضبطها يدويًا عن طريق إضافة مدخلات `nameserver` إلى /etc/resolv.conf.

**Warning**

إذا كان _dhcpcd_ قيد التشغيل، فلن تستمر التغييرات في /etc/resolv.conf. يمكن التحقق من الحالة باستخدام `ps x | grep dhcpcd`.

محرر nano مضمن في _وسيط إقلاع_ جنتو ويمكن استخدامه لتحرير /etc/resolv.conf عبر:

`root #` `nano /etc/resolv.conf`

يستعلم عن الأسطر التي تحتوي على الكلمة المفتاحية `nameserver` متبوعة بعنوان IP لخادوم DNS بترتيب تعريفها:

FILE **`/etc/resolv.conf`** **استخدام Quad9 DNS.**

```
nameserver 9.9.9.9
nameserver 149.112.112.112

```

FILE **`/etc/resolv.conf`** **استخدام Cloudflare DNS.**

```
nameserver 1.1.1.1
nameserver 1.0.0.1

```

يمكن التحقق من حالة DNS عن طريق عمل ping لاسم نطاق:

`root #` `ping -c 3 gentoo.org`

بمجرد [التحقق من الاتصال](/wiki/Handbook:AMD64/Installation/Networking/ar#Testing_the_network "Handbook:AMD64/Installation/Networking/ar")، تابع مع [تحضير الأقراص](/wiki/Handbook:AMD64/Installation/Disks/ar "Handbook:AMD64/Installation/Disks/ar").

[← Choosing the media](/wiki/Handbook:AMD64/Installation/Media "Previous part") [Home](/wiki/Handbook:AMD64 "Handbook:AMD64") [Preparing the disks →](/wiki/Handbook:AMD64/Installation/Disks "Next part")