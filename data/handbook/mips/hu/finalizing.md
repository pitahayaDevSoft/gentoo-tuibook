# Finalizing

Other languages:

- [Deutsch](/wiki/Handbook:MIPS/Installation/Finalizing/de "Handbuch:MIPS/Installation/Abschluss (100% translated)")
- [English](/wiki/Handbook:MIPS/Installation/Finalizing "Handbook:MIPS/Installation/Finalizing (100% translated)")
- [español](/wiki/Handbook:MIPS/Installation/Finalizing/es "Manual de Gentoo: MIPS/Instalación/Finalizar (100% translated)")
- [français](/wiki/Handbook:MIPS/Installation/Finalizing/fr "Handbook:MIPS/Installation/Finalizing/fr (100% translated)")
- [italiano](/wiki/Handbook:MIPS/Installation/Finalizing/it "Handbook:MIPS/Installation/Finalizing/it (100% translated)")
- magyar
- [polski](/wiki/Handbook:MIPS/Installation/Finalizing/pl "Handbook:MIPS/Installation/Finalizing (100% translated)")
- [português do Brasil](/wiki/Handbook:MIPS/Installation/Finalizing/pt-br "Manual:MIPS/Instalação/Finalizando (100% translated)")
- [русский](/wiki/Handbook:MIPS/Installation/Finalizing/ru "Handbook:MIPS/Installation/Finalizing (100% translated)")
- [தமிழ்](/wiki/Handbook:MIPS/Installation/Finalizing/ta "கையேடு:MIPS/நிறுவல்/முடித்தல் (100% translated)")
- [中文（中国大陆）‎](/wiki/Handbook:MIPS/Installation/Finalizing/zh-cn "手册：MIPS/安装/收尾安装工作 (100% translated)")
- [日本語](/wiki/Handbook:MIPS/Installation/Finalizing/ja "ハンドブック:MIPS/インストール/ファイナライズ (100% translated)")
- [한국어](/wiki/Handbook:MIPS/Installation/Finalizing/ko "Handbook:MIPS/Installation/Finalizing/ko (100% translated)")

[MIPS kézikönyv](/wiki/Handbook:MIPS/hu "Handbook:MIPS/hu")[A Gentoo Linux telepítése](/wiki/Handbook:MIPS/Full/Installation/hu "Handbook:MIPS/Full/Installation/hu")[A telepítésről](/wiki/Handbook:MIPS/Installation/About/hu "Handbook:MIPS/Installation/About/hu")[Telepítőképfájl kiválasztása](/wiki/Handbook:MIPS/Installation/Media/hu "Handbook:MIPS/Installation/Media/hu")[Hálózat beállítása](/wiki/Handbook:MIPS/Installation/Networking/hu "Handbook:MIPS/Installation/Networking/hu")[Adathordozók előkészítése](/wiki/Handbook:MIPS/Installation/Disks/hu "Handbook:MIPS/Installation/Disks/hu")[Stage fájl](/wiki/Handbook:MIPS/Installation/Stage/hu "Handbook:MIPS/Installation/Stage/hu")[Alaprendszer telepítése](/wiki/Handbook:MIPS/Installation/Base/hu "Handbook:MIPS/Installation/Base/hu")[Kernel beállítása](/wiki/Handbook:MIPS/Installation/Kernel/hu "Handbook:MIPS/Installation/Kernel/hu")[Rendszer beállítása](/wiki/Handbook:MIPS/Installation/System/hu "Handbook:MIPS/Installation/System/hu")[Eszközök telepítése](/wiki/Handbook:MIPS/Installation/Tools/hu "Handbook:MIPS/Installation/Tools/hu")[Bootloader beállítása](/wiki/Handbook:MIPS/Installation/Bootloader/hu "Handbook:MIPS/Installation/Bootloader/hu")Telepítés véglegesítése[Munka a Gentoo rendszerrel](/wiki/Handbook:MIPS/Full/Working/hu "Handbook:MIPS/Full/Working/hu")[Portage bemutatása](/wiki/Handbook:MIPS/Working/Portage/hu "Handbook:MIPS/Working/Portage/hu")[USE jelölőzászlók](/wiki/Handbook:MIPS/Working/USE/hu "Handbook:MIPS/Working/USE/hu")[Portage jellemzői](/wiki/Handbook:MIPS/Working/Features/hu "Handbook:MIPS/Working/Features/hu")[Init-szkript rendszer](/wiki/Handbook:MIPS/Working/Initscripts/hu "Handbook:MIPS/Working/Initscripts/hu")[Környezeti változók](/wiki/Handbook:MIPS/Working/EnvVar/hu "Handbook:MIPS/Working/EnvVar/hu")[Munka a Portage szoftvercsomag-kezelővel](/wiki/Handbook:MIPS/Full/Portage/hu "Handbook:MIPS/Full/Portage/hu")[Fájlok és könyvtárak](/wiki/Handbook:MIPS/Portage/Files/hu "Handbook:MIPS/Portage/Files/hu")[Változók](/wiki/Handbook:MIPS/Portage/Variables/hu "Handbook:MIPS/Portage/Variables/hu")[Szoftverágak keverése](/wiki/Handbook:MIPS/Portage/Branches/hu "Handbook:MIPS/Portage/Branches/hu")[További eszközök](/wiki/Handbook:MIPS/Portage/Tools/hu "Handbook:MIPS/Portage/Tools/hu")[Egyéni szoftvercsomag-tárolók](/wiki/Handbook:MIPS/Portage/CustomTree/hu "Handbook:MIPS/Portage/CustomTree/hu")[Fejlett funkciók](/wiki/Handbook:MIPS/Portage/Advanced/hu "Handbook:MIPS/Portage/Advanced/hu")[Hálózat beállítása OpenRC init-rendszeren](/wiki/Handbook:MIPS/Full/Networking/hu "Handbook:MIPS/Full/Networking/hu")[Munka elkezdése](/wiki/Handbook:MIPS/Networking/Introduction/hu "Handbook:MIPS/Networking/Introduction/hu")[Fejlett beállítások](/wiki/Handbook:MIPS/Networking/Advanced/hu "Handbook:MIPS/Networking/Advanced/hu")[Moduláris hálózat](/wiki/Handbook:MIPS/Networking/Modular/hu "Handbook:MIPS/Networking/Modular/hu")[Vezeték nélküli (Wi-Fi)](/wiki/Handbook:MIPS/Networking/Wireless/hu "Handbook:MIPS/Networking/Wireless/hu")[Funkcionalitás hozzáadása](/wiki/Handbook:MIPS/Networking/Extending/hu "Handbook:MIPS/Networking/Extending/hu")[Dinamikus menedzsment](/wiki/Handbook:MIPS/Networking/Dynamic/hu "Handbook:MIPS/Networking/Dynamic/hu")

## Contents

- [1Felhasználói adminisztráció](#Felhaszn.C3.A1l.C3.B3i_adminisztr.C3.A1ci.C3.B3)
  - [1.1Felhasználó hozzáadása a napi használat számára](#Felhaszn.C3.A1l.C3.B3_hozz.C3.A1ad.C3.A1sa_a_napi_haszn.C3.A1lat_sz.C3.A1m.C3.A1ra)
  - [1.2Ideiglenes jogosultság megemelése](#Ideiglenes_jogosults.C3.A1g_megemel.C3.A9se)
  - [1.3Root bejelentkezési lehetőség letiltása](#Root_bejelentkez.C3.A9si_lehet.C5.91s.C3.A9g_letilt.C3.A1sa)
- [2Adathordozó kitakarítása](#Adathordoz.C3.B3_kitakar.C3.ADt.C3.A1sa)
  - [2.1Telepítési műveletek maradványainak az eltávolítása](#Telep.C3.ADt.C3.A9si_m.C5.B1veletek_maradv.C3.A1nyainak_az_elt.C3.A1vol.C3.ADt.C3.A1sa)
- [3Hová tovább innen?](#Hov.C3.A1_tov.C3.A1bb_innen.3F)
  - [3.1További dokumentáció](#Tov.C3.A1bbi_dokument.C3.A1ci.C3.B3)
  - [3.2Gentoo online](#Gentoo_online)
    - [3.2.1Fórumok és IRC](#F.C3.B3rumok_.C3.A9s_IRC)
    - [3.2.2Levelezőlisták](#Levelez.C5.91list.C3.A1k)
    - [3.2.3Hibák](#Hib.C3.A1k)
    - [3.2.4Fejlesztési útmutató](#Fejleszt.C3.A9si_.C3.BAtmutat.C3.B3)
- [4Zárógondolatok](#Z.C3.A1r.C3.B3gondolatok)

## Felhasználói adminisztráció

### Felhasználó hozzáadása a napi használat számára

Root felhasználóként dolgozni egy Unix/Linux rendszeren veszélyes, és amennyire csak lehetséges, kerülendő. Ezért erősen ajánlott, hogy a mindennapi használathoz egy vagy több normál felhasználói fiókot is létrehozzon.

A csoportok, amelyeknek a felhasználó tagja, meghatározzák, hogy milyen tevékenységeket végezhet maga a felhasználó. Az alábbi táblázat számos fontos csoportot sorol fel:

Csoport
Leírás
audio
Engedélyezi a felhasználói fiók számára a hangeszközökhöz való hozzáférést.
cdrom
Engedélyezi a felhasználói fiók számára az optikai eszközökhöz való közvetlen hozzáférést.
cron
Engedélyezi a felhasználói fiók számára az időalapú feladatütemezést a cron használatával. Megjegyzés: A systemd szolgáltatási rendszert futtató rendszereken a felhasználói fiókok systemd időzítőket és felhasználói szolgáltatásfájlokat használhatnak a cron feladatok helyett.
floppy
Engedélyezi a felhasználói fiók számára az ősi mechanikus eszközök, az úgynevezett floppy meghajtók közvetlen elérését. Ezt a csoportot általában nem használják a modern rendszereken.
usb
Engedélyezi a felhasználói fiók számára a hozzáférést a hangeszközökhöz.
video
Engedélyezi a felhasználói fiók számára a hozzáférést a videórögzítő hardverekhez és a hardveres gyorsításhoz.
wheel
Engedélyezi a felhasználói fiók számára a su ( _substitute user_) parancs használatát, amely lehetővé teszi a root fiókra vagy más fiókokra való átváltást. Egyfelhasználós rendszerek esetén, amelyek tartalmaznak egy root fiókot. Érdemes ezt a csoportot hozzáadni az elsődleges normál felhasználóhoz.

Például, ahhoz hogy létrehozzunk egy _orbanviktor_ nevű felhasználót, aki a _wheel_, _users_ és _audio_ csoportok tagja, először root felhasználóként kell bejelentkezni (csak a root hozhat létre felhasználókat), és futtatni kell a useradd parancsot:

`Login:` `root`

```
Password: (Enter the root password)

```

Amikor jelszavakat állítunk be a normál felhasználói fiókokhoz, jó biztonsági gyakorlat, hogy ne használjuk ugyanazt vagy hasonló jelszót, mint amit a root felhasználó számára állítottunk be.

A kézikönyv szerzői azt javasolják, hogy legalább 16 karakter hosszúságú jelszót használjunk, amely teljesen egyedi minden más felhasználóhoz képest a rendszeren. Nyugodjon meg, az ilyen hosszúságú jelszót is fel fogja valaki törni. Ha így haladunk a fejlődéssel, akkor már lehet, hogy a közeljövőben.

`root #` `useradd -m -G users,wheel,audio -s /bin/bash orbanviktor
`

`root #` `passwd orbanviktor`

```
Password: (Enter the password for orbanviktor)
Re-enter password: (Re-enter the password to verify)

```

### Ideiglenes jogosultság megemelése

Ha egy felhasználónak valaha is root jogosultságokkal kell elvégeznie egy feladatot, akkor használhatja a su - parancsot, hogy ideiglenesen root jogosultságokat kapjon. Egy másik módja a [sudo](/wiki/Sudo/hu "Sudo/hu") ([app-admin/sudo](https://packages.gentoo.org/packages/app-admin/sudo)) vagy [doas](/wiki/Doas "Doas") ([app-admin/doas](https://packages.gentoo.org/packages/app-admin/doas)) segédprogramok használata, amelyek helyes beállítás esetén nagyon biztonságosak.

### Root bejelentkezési lehetőség letiltása

**Warning**

Mielőtt letiltaná a root bejelentkezést, győződjön meg arról, hogy egy felhasználói fiók tagja a wheel csoportnak, és létezik mód a felhasználói jogosultságok emelésére, különben a root hozzáférés le lesz zárva, és a rendszer-adminisztráció lehetetlenné válik helyreállítási művelet nélkül. Néhány általános módszer a jogosultságok emelésére: [app-admin/sudo](https://packages.gentoo.org/packages/app-admin/sudo), [app-admin/doas](https://packages.gentoo.org/packages/app-admin/doas), vagy a systemd run0.

A lehetséges fenyegetések megelőzése érdekében a root felhasználóval történő bejelentkezést úgy akadályozhatjuk meg, hogy töröljük a root jelszót és/vagy letiltjuk a root bejelentkezést.

A root bejelentkezés letiltásához:

`root #` `passwd -l root`

A root jelszó törléséhez és a bejelentkezés letiltásához:

`root #` `passwd -dl root`

## Adathordozó kitakarítása

### Telepítési műveletek maradványainak az eltávolítása

Miután a Gentoo telepítése befejeződött és az operációs rendszer újraindult, ha minden jól ment, a fokozat (stage) fájlt és más telepítési műveletek maradványait - mint például a DIGEST, CONTENT vagy \*.asc (PGP aláírás) fájlokat - most biztonságosan el lehet távolítani.

A fájlok a / könyvtárban találhatók, és az alábbi parancs segítségével törölhetők:

`root #` `rm /stage3-*.tar.*`

## Hová tovább innen?

Nem tudja, merre induljon tovább? Számos út létezik a felfedezésre... A Gentoo rengeteg lehetőséget kínál felhasználóinak, és ezért sok dokumentált (és kevésbé dokumentált) funkciót kell felfedezni itt a wiki-n és más Gentoo-hoz kapcsolódó aldoménekben (tekintse meg az alábbi [Gentoo online](/wiki/Handbook:MIPS/Installation/Finalizing/hu#Gentoo_online "Handbook:MIPS/Installation/Finalizing/hu") szakaszt).

### További dokumentáció

Fontos megjegyezni, hogy a Gentoo operációs rendszerben elérhető számos választási lehetőség miatt a kézikönyv által nyújtott dokumentáció korlátozott terjedelmű - főként a Gentoo rendszer telepítésének és alapvető rendszerkezelési tevékenységek működésbe hozatalára összpontosít. A kézikönyv szándékosan kizárja a grafikus környezetekkel, a rendszer keményítésével és más fontos adminisztratív feladatokkal kapcsolatos utasításokat. Mindezek ellenére a kézikönyv további szakaszai segítenek az olvasóknak az alapvető funkciókkal.

Az olvasóknak mindenképpen érdemes megnézniük a kézikönyv következő részét, amelynek címe: [Working with Gentoo](/wiki/Handbook:MIPS/Working/Portage/hu "Handbook:MIPS/Working/Portage/hu"), amely elmagyarázza, hogyan tarthatják naprakészen a szoftvereket, hogyan telepíthetnek további szoftvercsomagokat, részleteket a USE zászlókról, az OpenRC init rendszerről és számos más hasznos témáról a Gentoo rendszer telepítés utáni kezelésével kapcsolatban.

A kézikönyv mellett az olvasóknak is érdemes felfedezni a Gentoo wiki más részeit, hogy további, közösség által biztosított dokumentációkat találjanak. A Gentoo wiki csapata egy [dokumentációs téma áttekintést](/wiki/Main_Page#Documentation_topics "Main Page") is kínál, amely kategóriák szerint sorolja fel a wiki cikkek válogatását. Például utal a [lokalizációs útmutatóra](/wiki/Localization/Guide/hu "Localization/Guide/hu"), hogy a rendszert otthonosabbá tegyék (különösen hasznos azoknak a felhasználóknak, akik második nyelvként beszélnek angolul).

A legtöbb felhasználó, aki asztali környezetben dolgozik, grafikus környezetet állít be, hogy natív módon tudjon dolgozni. Számos közösség által karbantartott 'meta' cikk létezik a támogatott [asztali környezetek (DE-k)](/wiki/Desktop_environment/hu "Desktop environment/hu") és [ablakkezelők (WM-ek)](/wiki/Window_manager/hu "Window manager/hu") számára. Az olvasóknak tisztában kell lenniük azzal, hogy minden egyes DE kicsit eltérő beállítási lépéseket igényel, ami megnöveli a telepítés és beállítás bonyolultságát.

Számos más [Meta articles](/wiki/Category:Meta "Category:Meta") létezik, amelyek olvasóinknak magas szintű áttekintést nyújtanak a Gentoo-ban elérhető szoftverekről.

### Gentoo online

**Important**

Az olvasóknak tudniuk kell, hogy az összes hivatalos Gentoo webhelyet a Gentoo [viselkedési szabályzata](/wiki/Project:Council/Code_of_conduct "Project:Council/Code of conduct") irányítja. A Gentoo közösségben való aktív részvétel kiváltság, nem jog, és a felhasználóknak tisztában kell lenniük azzal, hogy a viselkedési szabályzat létezésének oka van.

A Libera.Chat által üzemeltetett internetes relé csevegő (IRC) hálózat és a levelezőlisták kivételével a legtöbb Gentoo webhelyhez egyedi fiók szükséges a kérdések feltevéséhez, a beszélgetések megnyitásához vagy a hibajegyek bejelentéséhez.

#### Fórumok és IRC

Minden felhasználó szívesen látott a [Gentoo fórumokon](https://forums.gentoo.org) vagy az [internetes relé csevegő csatornáink](https://www.gentoo.org/get-involved/irc-channels/) egyikén. Könnyen kereshető a fórumokon, hogy a friss Gentoo telepítés során tapasztalt problémát korábban mások felfedezték-e, és visszajelzések után megoldották-e. Meglepő lehet, hogy milyen gyakran találkoznak más felhasználók a Gentoo első telepítési problémáival. Ajánlott, hogy a felhasználók először keresgéljenek a fórumokon és a wikiben, mielőtt segítséget kérnének a Gentoo támogatási csatornákon.

#### Levelezőlisták

[Több levelezőlista](https://www.gentoo.org/get-involved/mailing-lists/) áll rendelkezésre a közösségi tagok számára, akik inkább e-mailben kérnek támogatást vagy visszajelzést, mint hogy felhasználói fiókot hozzanak létre a fórumokon vagy az IRC-n. A felhasználóknak követniük kell az utasításokat, hogy feliratkozhassanak a konkrét levelezőlistákra.

#### Hibák

Néha, miután áttekintettük a wikit, keresgéltünk a fórumokon, és támogatást kértünk az IRC csatornán vagy levelezőlistákon, nincs ismert megoldás egy problémára. Általában ez annak jele, hogy hibát kell nyitni a Gentoo [Bugzilla oldalán](https://bugs.gentoo.org).

#### Fejlesztési útmutató

Azok az olvasók, akik többet szeretnének megtudni a Gentoo fejlesztéséről, nézzenek utána a [Fejlesztési útmutatóban](https://devmanual.gentoo.org/). Ez az útmutató részletes utasításokat nyújt az ebuild-ek írásáról, eclass-okkal való munkáról, valamint számos [általános fogalom](https://devmanual.gentoo.org/general-concepts/index.html) meghatározását tartalmazza, amelyek a Gentoo fejlesztésének hátterében állnak.

## Zárógondolatok

A Gentoo egy robusztus, rugalmas és kiválóan karbantartott disztribúció. A fejlesztői közösség örömmel fogadja a visszajelzéseket arról, hogy miként lehetne még _jobb_ disztribúcióvá tenni a Gentoo rendszert.

Emlékeztetőül: Minden visszajelzés, amelyet _ehhez a kézikönyvhöz_ hozzá kíván adni, kövesse a kézikönyv elején található [Hogyan javíthatom a kézikönyvet?](/wiki/Handbook:Main_Page/hu#How_do_I_improve_the_Handbook.3F "Handbook:Main Page/hu") szakaszban részletezett irányelveket.

Kíváncsian várjuk, hogy felhasználóink ​​hogyan választják majd a Gentoo megvalósítását egyedi használati eseteiknek és igényeiknek megfelelően.

[← Bootloader beállítása](/wiki/Handbook:MIPS/Installation/Bootloader/hu "Previous part") [Kezdőlap](/wiki/Handbook:MIPS/hu "Handbook:MIPS/hu") [Portage bemutatása →](/wiki/Handbook:MIPS/Working/Portage/hu "Next part")