# About

## Contents

- [1Bevezetés](#Bevezet.C3.A9s)
  - [1.1Üdvözlés](#.C3.9Cdv.C3.B6zl.C3.A9s)
  - [1.2Hogyan épül fel a telepítés?](#Hogyan_.C3.A9p.C3.BCl_fel_a_telep.C3.ADt.C3.A9s.3F)
  - [1.3Melyek a lehetőségek?](#Melyek_a_lehet.C5.91s.C3.A9gek.3F)
  - [1.4Problémák](#Probl.C3.A9m.C3.A1k)

## Bevezetés

### Üdvözlés

Először is, üdvözöljük a Gentoo operációs rendszerben. Ön most belép a választások és a teljesítmény világába. A Gentoo teljes egészében a választási lehetőségekről szól. A Gentoo telepítése során ez többször is egyértelművé válik. A felhasználók kiválaszthatják, hogy mennyit szeretnének maguknak fordítani, hogyan szeretnék telepíteni a Gentoo operációs rendszert, milyen rendszernaplózót akarnak használni, stb.

A Gentoo egy gyors, modern meta-disztribúció, tiszta és rugalmas kialakítással. A Gentoo a szabad szoftverekre épül, és nem rejti el a felhasználói elől, hogy mi található a motorháztető alatt. A Portage, amelyet a Gentoo szoftvercsomag karbantartó rendszerként használ, Python nyelven íródott, így a felhasználó könnyen megtekintheti és módosíthatja a forráskódot. A Gentoo szoftvercsomag rendszere forráskódot használ (bár előre binárisra lefordított szoftvercsomagok támogatása is rendelkezésre áll), és a Gentoo beállítása szokásos szövegfájlokon keresztül történik. Más szavakkal: Abszolút átláthatóság van mindenütt.

Nagyon fontos, hogy mindenki megértse, hogy a választási lehetőségek azok, amelyek működtetik a Gentoo operációs rendszert. Igyekszünk nem rákényszeríteni a felhasználókra olyasmit, amit nem kedvelnek. Ha bárki máshogy gondolja ezt, akkor kérjük, hogy jelezze [a hibabejelentő](https://bugs.gentoo.org) weboldalon.

### Hogyan épül fel a telepítés?

A Gentoo telepítés egy 10 lépésből álló folyamatként tekinthető, amely megfelel a következő fejezeteknek. Minden lépés egy adott állapotot eredményez:

01. Lépés után: Felhasználó egy működő környezetben van, amely készen áll a Gentoo telepítésére.
02. Lépés után: Internetkapcsolat készen áll a Gentoo telepítésére.
03. Lépés után: Számítógépen lévő adathordozó(k) inicializálva van(nak), hogy befogadják a Gentoo telepítést.
04. Lépés után: Telepítési környezet elő van készítve. A felhasználó készen áll arra, hogy a [chroot](/wiki/Chroot/hu "Chroot/hu") paranccsal belépjen az új telepítési környezetbe.
05. Lépés után: Telepítési környezetbe belelépve telepítve lettek azok a legalapvetőbb szoftvercsomagok, amelyek minden Gentoo telepítésnél azonosak szoktak lenni.
06. Lépés után: Telepítve van a Linux rendszermag (a kernel).
07. Lépés után: Felhasználó a Gentoo operációs rendszer beállításfájljainak nagy részét beállította.
08. Lépés után: Szükséges rendszereszközök telepítve vannak.
09. Lépés után: Megfelelő rendszerbetöltő (bootloader) szoftver telepítve van és megfelelően be van állítva.
10. Lépés után: Frissen feltelepített Gentoo Linux operációs rendszerünk készen áll a felfedezésre.

Amikor egy adott választási lehetőség kerül bemutatásra, a kézikönyv megpróbálja elmagyarázni, hogy mik az előnyök és hátrányok. Bár a szöveg ezután egy alapértelmezett választással folytatódik (a címben "Alapértelmezett: " megjelöléssel azonosítva), a többi lehetőség is dokumentálva lesz (a címben "Alternatíva: " megjelöléssel). Ne gondolja, hogy az alapértelmezett az, amit a Gentoo ajánl. Azonban ez az, amit a Gentoo úgy gondol, hogy a legtöbb felhasználó használni fog.

Egyes esetekben egy választható lépést is követhetünk. Az ilyen lépések "Választható: " megjelöléssel vannak ellátva, ezért nem szükségesek a Gentoo telepítéséhez. Ugyanakkor bizonyos választható lépések egy korábban meghozott döntéstől függnek. Az útmutató tájékoztatja az olvasót, amikor ez megtörténik, mind a döntés meghozatalakor, mind közvetlenül a választható lépés leírása előtt.

### Melyek a lehetőségek?

A Gentoo számos különböző módon telepíthető. ARM64 hardver esetében ez általában egy rendszerképfájl letöltését jelenti, amelyet eMMC, SD vagy uSD adathordozóra lehet DD parancs segítségével kiírni.

Ez a dokumentum kifejezetten az ARM64 hardverekre történő telepítést tárgyalja. Igyekszünk lefedni az ismert hardverek specifikumait, amelyeket teszteltünk és érvényesítettünk. Azok az ARM64 hardverek, amelyek nincsenek felsorolva, valószínűleg működni fognak, azonban Önnek meg kell értenie, hogy mi szükséges az adott hardveréhez, és ennek megfelelően kell módosítania a telepítést.

**Note**

A többi telepítési megközelítéshez, beleértve a nem Gentoo CD lemezek használatát, kérjük, olvassa el [Alternatív Telepítési Útmutatónkat](/wiki/Installation_alternatives "Installation alternatives").

Ezenkívül biztosítunk egy [Gentoo Telepítési Tippek és Trükkök](/wiki/Gentoo_installation_tips_and_tricks "Gentoo installation tips and tricks") dokumentumot is, amely szintén hasznos lehet az olvasáshoz.

### Problémák

Ha hibát talál a telepítésben (vagy a telepítési dokumentációban), akkor kérjük, hogy látogasson el a [hibakövető rendszerünkbe](https://bugs.gentoo.org), és ellenőrizze, hogy ismert-e a hiba. Ha nem, kérjük, hozzon létre róla hibabejelentést, hogy foglalkozhassunk vele. Ne féljen a hibákhoz kijelölt fejlesztőktől – ők (általában) nem esznek embereket.

Figyelembe kell venni azonban, hogy bár ez a dokumentum architektúra-specifikus, tartalmazhat hivatkozásokat más architektúrákra is. Ennek oka, hogy a Gentoo kézikönyv nagy része olyan forráskódot használ, amely minden architektúrára közös (hogy elkerüljük az erőfeszítések megduplázását és a fejlesztési erőforrások kimerülését). Igyekszünk ezt minimális szinten tartani, hogy elkerüljük a félreértéseket.

Ha nem egyértelmű, hogy a probléma felhasználói probléma-e (valamilyen hiba, amely annak ellenére történt, hogy alaposan áttanulmányozták a dokumentációt), vagy szoftveres probléma-e (valamilyen hiba, amely annak ellenére történt, hogy alaposan teszteltük a telepítést/dokumentációt), akkor mindenkit szívesen látunk a [#gentoo](ircs://irc.libera.chat/#gentoo) ([webchat](https://web.libera.chat/#gentoo)) IRC csatornánkon. Természetesen mindenkit szívesen látunk más esetekben is, mivel chat csatornánk a Gentoo széles spektrumát lefedi.

Ha további kérdések merülnének fel a Gentoo operációs rendszerrel kapcsolatban, akkor tekintse meg a [Gyakran Ismételt Kérdéseinket](/wiki/FAQ "FAQ"), amelyek elérhetők a [Gentoo Wikiben](/wiki/Main_Page "Main Page"). Emellett találhatók [GYIK](https://forums.gentoo.org/viewforum.php?f=40) a [Gentoo fórumokon](https://forums.gentoo.org) is.