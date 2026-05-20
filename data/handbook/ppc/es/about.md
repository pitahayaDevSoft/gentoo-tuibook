# About

Other languages:

- [Deutsch](/wiki/Handbook:PPC/Installation/About/de "Handbuch:PPC/Installation/Über (100% translated)")
- [English](/wiki/Handbook:PPC/Installation/About "Handbook:PPC/Installation/About (100% translated)")
- español
- [français](/wiki/Handbook:PPC/Installation/About/fr "Handbook:PPC/Installation/About/fr (100% translated)")
- [italiano](/wiki/Handbook:PPC/Installation/About/it "Handbook:PPC/Installation/About (100% translated)")
- [magyar](/wiki/Handbook:PPC/Installation/About/hu "Handbook:PPC/Installation/About/hu (100% translated)")
- [polski](/wiki/Handbook:PPC/Installation/About/pl "Handbook:PPC/Installation/About (100% translated)")
- [português do Brasil](/wiki/Handbook:PPC/Installation/About/pt-br "Manual:PPC/Instalação/Sobre (100% translated)")
- [русский](/wiki/Handbook:PPC/Installation/About/ru "Handbook:PPC/Installation/About (100% translated)")
- [தமிழ்](/wiki/Handbook:PPC/Installation/About/ta "கையேடு:PPC/நிறுவல்/பற்றி (100% translated)")
- [中文（中国大陆）‎](/wiki/Handbook:PPC/Installation/About/zh-cn "手册：PPC/安装/关于 (100% translated)")
- [日本語](/wiki/Handbook:PPC/Installation/About/ja "ハンドブック:PPC/インストール/About (100% translated)")
- [한국어](/wiki/Handbook:PPC/Installation/About/ko "Handbook:PPC/Installation/About/ko (100% translated)")

[Manual PPC](/wiki/Handbook:PPC/es "Handbook:PPC/es")[Instalación](/wiki/Handbook:PPC/Full/Installation/es "Handbook:PPC/Full/Installation/es")Acerca de la instalación[Elegir los medios](/wiki/Handbook:PPC/Installation/Media/es "Handbook:PPC/Installation/Media/es")[Configurar la red](/wiki/Handbook:PPC/Installation/Networking/es "Handbook:PPC/Installation/Networking/es")[Preparar los discos](/wiki/Handbook:PPC/Installation/Disks/es "Handbook:PPC/Installation/Disks/es")[Instalar el stage3](/wiki/Handbook:PPC/Installation/Stage/es "Handbook:PPC/Installation/Stage/es")[Instalar el sistema base](/wiki/Handbook:PPC/Installation/Base/es "Handbook:PPC/Installation/Base/es")[Configurar el núcleo](/wiki/Handbook:PPC/Installation/Kernel/es "Handbook:PPC/Installation/Kernel/es")[Configurar el sistema](/wiki/Handbook:PPC/Installation/System/es "Handbook:PPC/Installation/System/es")[Instalar las herramientas](/wiki/Handbook:PPC/Installation/Tools/es "Handbook:PPC/Installation/Tools/es")[Configurar el cargador de arranque](/wiki/Handbook:PPC/Installation/Bootloader/es "Handbook:PPC/Installation/Bootloader/es")[Terminar](/wiki/Handbook:PPC/Installation/Finalizing/es "Handbook:PPC/Installation/Finalizing/es")[Trabajar con Gentoo](/wiki/Handbook:PPC/Full/Working/es "Handbook:PPC/Full/Working/es")[Introducción a Portage](/wiki/Handbook:PPC/Working/Portage/es "Handbook:PPC/Working/Portage/es")[Ajustes USE](/wiki/Handbook:PPC/Working/USE/es "Handbook:PPC/Working/USE/es")[Características de Portage](/wiki/Handbook:PPC/Working/Features/es "Handbook:PPC/Working/Features/es")[Sistema de guiones de inicio](/wiki/Handbook:PPC/Working/Initscripts/es "Handbook:PPC/Working/Initscripts/es")[Variables de entorno](/wiki/Handbook:PPC/Working/EnvVar/es "Handbook:PPC/Working/EnvVar/es")[Trabajar con Portage](/wiki/Handbook:PPC/Full/Portage/es "Handbook:PPC/Full/Portage/es")[Ficheros y directorios](/wiki/Handbook:PPC/Portage/Files/es "Handbook:PPC/Portage/Files/es")[Variables](/wiki/Handbook:PPC/Portage/Variables/es "Handbook:PPC/Portage/Variables/es")[Mezclar ramas de software](/wiki/Handbook:PPC/Portage/Branches/es "Handbook:PPC/Portage/Branches/es")[Herramientas adicionales](/wiki/Handbook:PPC/Portage/Tools/es "Handbook:PPC/Portage/Tools/es")[Repositorios personalizados de paquetes](/wiki/Handbook:PPC/Portage/CustomTree/es "Handbook:PPC/Portage/CustomTree/es")[Características avanzadas](/wiki/Handbook:PPC/Portage/Advanced/es "Handbook:PPC/Portage/Advanced/es")[Configuración de la red](/wiki/Handbook:PPC/Full/Networking/es "Handbook:PPC/Full/Networking/es")[Comenzar](/wiki/Handbook:PPC/Networking/Introduction/es "Handbook:PPC/Networking/Introduction/es")[Configuración avanzada](/wiki/Handbook:PPC/Networking/Advanced/es "Handbook:PPC/Networking/Advanced/es")[Configuración de red modular](/wiki/Handbook:PPC/Networking/Modular/es "Handbook:PPC/Networking/Modular/es")[Conexión inalámbrica](/wiki/Handbook:PPC/Networking/Wireless/es "Handbook:PPC/Networking/Wireless/es")[Añadir funcionalidad](/wiki/Handbook:PPC/Networking/Extending/es "Handbook:PPC/Networking/Extending/es")[Gestión dinámica](/wiki/Handbook:PPC/Networking/Dynamic/es "Handbook:PPC/Networking/Dynamic/es")

## Contents

- [1Introducción](#Introducci.C3.B3n)
  - [1.1Bienvenido](#Bienvenido)
    - [1.1.1Franqueza](#Franqueza)
    - [1.1.2Elección](#Elecci.C3.B3n)
    - [1.1.3Potencia](#Potencia)
  - [1.2Cómo se estructura la instalación](#C.C3.B3mo_se_estructura_la_instalaci.C3.B3n)
    - [1.2.1Decidir qué pasos dar](#Decidir_qu.C3.A9_pasos_dar)
    - [1.2.2Pasos sugeridos](#Pasos_sugeridos)
    - [1.2.3Pasos opcionales](#Pasos_opcionales)
    - [1.2.4Pasos obsoletos](#Pasos_obsoletos)
    - [1.2.5Valores predeterminados y alternativas](#Valores_predeterminados_y_alternativas)
  - [1.3Opciones de instalación para Gentoo](#Opciones_de_instalaci.C3.B3n_para_Gentoo)
  - [1.4Problemas](#Problemas)

## Introducción

### Bienvenido

¡Bienvenidos a Gentoo! Gentoo es un sistema operativo gratuito basado en Linux que se puede optimizar y personalizar automáticamente para casi cualquier aplicación o necesidad. Está construido sobre un ecosistema de software libre y no oculta a sus usuarios lo que hay bajo el capó.

#### Franqueza

Las principales herramientas de Gentoo están construidas a partir de lenguajes de programación simples. [Portage](/wiki/Portage/es "Portage/es"), el sistema de mantenimiento de paquetes de Gentoo, está [escrito en Python](https://gitweb.gentoo.org/proj/portage.git/). Ebuilds, que proporcionan definiciones de paquetes para Portage [están escritos en bash](https://gitweb.gentoo.org/repo/gentoo.git). Se anima a nuestros usuarios a revisar, modificar y mejorar el código fuente de todas las partes de Gentoo.

De forma predeterminada, los paquetes sólo se parchean cuando es necesario para corregir errores o proporcionar interoperabilidad dentro de Gentoo. Se instalan en el sistema compilando el código fuente proporcionado por proyectos externos en formato binario (aunque también se incluye soporte para paquetes binarios precompilados). La configuración de Gentoo se realiza a través de archivos de texto.

Por las razones anteriores y otras: la "franqueza" está incorporada como un "principio de diseño".

#### Elección

La "elección" es otro "principio de diseño" de Gentoo.

Al instalar Gentoo, la elección se ve claramente a lo largo del Manual. Los administradores del sistema pueden elegir entre dos sistemas de inicio totalmente compatibles (el propio [OpenRC](/wiki/OpenRC/es "OpenRC/es") de Gentoo y el [systemd](/wiki/Systemd/es "Systemd/es") de Freedesktop.org), la estructura de partición para los discos de almacenamiento, qué sistemas de archivos usar en el disco( s), un [perfil del sistema](/wiki/Profile "Profile") de referencia, elimine o agregue funciones a nivel global (en todo el sistema) o específico de cada paquete a través de indicadores USE, cargador de arranque, utilidad de administración de red y mucho, mucho más.

Como filosofía de desarrollo, [Los autores de Gentoo](https://www.gentoo.org/inside-gentoo/developers/) intentan evitar forzar a los usuarios a utilizar un perfil de sistema o entorno de escritorio específico. Si se ofrece algo en el ecosistema GNU/Linux, es probable que esté disponible en Gentoo. Si no, nos encantaría que lo estuviera. Para paquetes nuevos, se recomienda enviarlo primero a [GURU](/wiki/GURU "GURU"). Una vez que esté desarrollado y un desarrollador de Gentoo se haya ofrecido a patrocinarlo, se podrá añadir al repositorio oficial de paquetes de Gentoo.

#### Potencia

Ser un sistema operativo basado en código fuente permite que Gentoo pueda ser portado a nuevos [conjuntos de instrucciones de arquitectura](https://en.wikipedia.org/wiki/instruction_set_architecture "wikipedia:instruction set architecture") de procesador y también permite que todos los paquetes instalados sean ajustados. Esta fortaleza hace emerger otro “principio de diseño” de Gentoo: la “potencia”.

Un administrador de sistemas que haya instalado y personalizado Gentoo con éxito habrá compilado un sistema operativo personalizado a partir del código fuente. Todo el sistema operativo se puede ajustar a nivel binario a través de los mecanismos incluidos en el archivo [make.conf](/wiki//etc/portage/make.conf/es "/etc/portage/make.conf/es") de Portage. Si así lo desea, se pueden realizar ajustes por paquete o por grupos de paquetes. De hecho, se pueden agregar o eliminar conjuntos completos de funciones utilizando indicadores USE.

Es muy importante que el lector del Manual comprenda que estos principios de diseño son los que hacen que Gentoo sea único. Con los principios de gran potencia, muchas opciones y franqueza extrema resaltados, se debe emplear rigor, reflexión e intencionalidad al usar Gentoo.

### Cómo se estructura la instalación

La instalación de Gentoo puede verse como un procedimiento de diez pasos. Después de cada paso se alcanza cierto estado:

PasoResultado
1El usuario dispone de un entorno de trabajo listo para instalar Gentoo.
2La conexión a Internet estará preparada para instalar Gentoo.
3Los discos duros están inicializados para alojar la instalación Gentoo.
4El entorno de instalación está preparado y el usuario puede entrar en una jaula [chroot](/wiki/Chroot/es "Chroot/es").
5Los paquetes principales, que son los mismos en toda instalación Gentoo, están instalados.
6El núcleo Linux está instalado.
7Se crean la mayoría de los archivos de configuración del sistema.
8Las herramientas del sistema necesarias están instaladas.
9Se ha instalado y configurado el cargador de arranque apropiado.
10El entorno Gentoo Linux recién instalado está preparado para ser explorado.

#### Decidir qué pasos dar

El manual presenta una cantidad abrumadora de opciones, especialmente para alguien que nunca ha instalado Linux sin un instalador.

Es importante tener en cuenta que el manual está diseñado para describir los pasos necesarios para la instalación en una amplia variedad de hardware, con diferentes necesidades de instalación. Por ello, muchas de las opciones presentadas en el manual son innecesarias para una instalación específica y pueden omitirse.

#### Pasos sugeridos

Con el prefijo " **Sugerido:**", algunos pasos no son estrictamente necesarios, pero ayudan en la mayoría de los casos, como instalar [sys-kernel/linux-firmware](https://packages.gentoo.org/packages/sys-kernel/linux-firmware).

#### Pasos opcionales

Muchas secciones del manual, con el prefijo " **Opcional:**", son puramente opcionales y se pueden omitir si el usuario busca una instalación simple, mayoritariamente básica.

Ejemplos de esto pueden ser la personalización de los indicadores del compilador, el uso de un núcleo totalmente personalizado o la desactivación del inicio de sesión para root.

**Consejo**

Al seguir los pasos opcionales, es importante asegurarse que se cumplan todos los prerrequisitos. Algunos pasos opcionales dependen de otros.

#### Pasos obsoletos

Gentoo existe desde hace mucho tiempo. Algunos procesos de instalación se describen en el manual porque antes eran más relevantes, pero ahora están prácticamente obsoletos. En lugar de eliminar esta información inmediatamente, ya que podría ser útil para algunos usuarios, se puede marcar como **Obsoleto:** antes de eliminarla. Una vez eliminada, se debe usar el _historial_ para ver este contenido.

#### Valores predeterminados y alternativas

Siempre que se presenten opciones, el manual intentará explicar los pros y los contras de cada una de ellas.

Si las opciones posibles son mutuamente excluyentes, se utiliza " **Predeterminado:**" para indicar la opción con mayor soporte o elegida mayoritariamente, mientras que las alternativas se indicarán con " **Alternativa**".

**Nota**

Las opciones **alternativas** no son inferiores a las **Predeterminadas**, pero las opciones **Predeterminadas** suelen usarse más ampliamente y pueden tener mejor soporte.

### Opciones de instalación para Gentoo

Gentoo se puede instalar de formas muy diversas. Se puede descargar e instalar desde un medio de instalación oficial de Gentoo como nuestras imágenes ISO arrancables. Los medios se pueden instalar desde un dispositivo de almacenamiento USB o se pude acceder a ellos a través de un entorno arrancado desde red. Alternativamente, se puede instalar Gentoo desde medios no oficiales como una distribución ya instalada o un disco que no contenga Gentoo (por ejemplo [Linux Mint](https://linuxmint.com/)).

Este documento cubre la instalación utilizando un medio de instalación oficial de Gentoo o, en algunos casos, instalación por red.

**Nota**

Para encontrar ayuda acerca de otros procedimientos de instalación, incluyendo el uso de medios arrancables ajenos a Gentoo, por favor, lea nuestra [guía sobre métodos alternativos de instalación](/wiki/Installation_alternatives/es "Installation alternatives/es").

También ofrecemos un documento de [consejos y trucos para instalar Gentoo](/wiki/Gentoo_installation_tips_and_tricks/es "Gentoo installation tips and tricks/es") que también puede ser de utilidad.

### Problemas

Si encuentra algún problema durante la instalación (o con el documento de instalación), por favor, visite nuestro [sistema de seguimiento de incidencias](https://bugs.gentoo.org) y compruebe si el problema es ya conocido. Si no lo es, por favor, cree un informe sobre él para que podamos echarle un vistazo. No tema a los desarrolladores a los que se les han asignado los informes de error, (normalmente) no se comen a nadie.

Aunque este documento es específico de la arquitectura puede contener referencias a otras arquitecturas ya que muchas partes del manual de Gentooo utilizan texto que es idéntico para todas las arquitecturas (con el fin de evitar la duplicación de esfuerzos). Estas referencias se conservan mínimamente para evitar confusiones.

Si hay alguna duda sobre si el problema es o no un problema del usuario (algún error cometido a pesar de haber leído la documentación cuidadosamente) o un problema de software (algún error que cometimos a pesar de haber probado la instalación/documentación cuidadosamente), todo el mundo es bienvenido a unirse al canal [#gentoo](ircs://irc.libera.chat/#gentoo) ([webchat](https://web.libera.chat/#gentoo)) en irc.libera.chat. Por supuesto, cualquier otro es bienvenido también, ya que nuestro canal de chat cubre el amplio espectro de Gentoo.

Hablando de esto, si tiene cualquier pregunta adicional concerniente a Gentoo, eche un vistazo al artículo [FAQ](/wiki/FAQ/es "FAQ/es"). También están disponibles las [FAQs](https://forums.gentoo.org/viewforum.php?f=40) en los [Foros de Gentoo](https://forums.gentoo.org).

[←](/wiki/Handbook:PPC/es "Previous part") [Inicio](/wiki/Handbook:PPC/es "Handbook:PPC/es") [Elegir los medios →](/wiki/Handbook:PPC/Installation/Media/es "Next part")