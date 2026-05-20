# Stage

Other languages:

- [Deutsch](/wiki/Handbook:PPC/Installation/Stage/de "Handbuch:PPC/Installation/Stage (100% translated)")
- [English](/wiki/Handbook:PPC/Installation/Stage "Handbook:PPC/Installation/Stage (100% translated)")
- español
- [français](/wiki/Handbook:PPC/Installation/Stage/fr "Handbook:PPC/Installation/Stage/fr (100% translated)")
- [italiano](/wiki/Handbook:PPC/Installation/Stage/it "Handbook:PPC/Installation/Stage/it (100% translated)")
- [magyar](/wiki/Handbook:PPC/Installation/Stage/hu "Handbook:PPC/Installation/Stage/hu (100% translated)")
- [polski](/wiki/Handbook:PPC/Installation/Stage/pl "Handbook:PPC/Installation/Stage (100% translated)")
- [português do Brasil](/wiki/Handbook:PPC/Installation/Stage/pt-br "Handbook:PPC/Installation/Stage/pt-br (50% translated)")
- [русский](/wiki/Handbook:PPC/Installation/Stage/ru "Handbook:PPC/Installation/Stage (100% translated)")
- [தமிழ்](/wiki/Handbook:PPC/Installation/Stage/ta "கையேடு:PPC/நிறுவல்/நிலை (100% translated)")
- [中文（中国大陆）‎](/wiki/Handbook:PPC/Installation/Stage/zh-cn "手册：PPC/安装/安装stage3 (100% translated)")
- [日本語](/wiki/Handbook:PPC/Installation/Stage/ja "ハンドブック:PPC/インストール/Stage (100% translated)")
- [한국어](/wiki/Handbook:PPC/Installation/Stage/ko "Handbook:PPC/Installation/Stage/ko (100% translated)")

[Manual PPC](/wiki/Handbook:PPC/es "Handbook:PPC/es")[Instalación](/wiki/Handbook:PPC/Full/Installation/es "Handbook:PPC/Full/Installation/es")[Acerca de la instalación](/wiki/Handbook:PPC/Installation/About/es "Handbook:PPC/Installation/About/es")[Elegir los medios](/wiki/Handbook:PPC/Installation/Media/es "Handbook:PPC/Installation/Media/es")[Configurar la red](/wiki/Handbook:PPC/Installation/Networking/es "Handbook:PPC/Installation/Networking/es")[Preparar los discos](/wiki/Handbook:PPC/Installation/Disks/es "Handbook:PPC/Installation/Disks/es")Instalar el stage3[Instalar el sistema base](/wiki/Handbook:PPC/Installation/Base/es "Handbook:PPC/Installation/Base/es")[Configurar el núcleo](/wiki/Handbook:PPC/Installation/Kernel/es "Handbook:PPC/Installation/Kernel/es")[Configurar el sistema](/wiki/Handbook:PPC/Installation/System/es "Handbook:PPC/Installation/System/es")[Instalar las herramientas](/wiki/Handbook:PPC/Installation/Tools/es "Handbook:PPC/Installation/Tools/es")[Configurar el cargador de arranque](/wiki/Handbook:PPC/Installation/Bootloader/es "Handbook:PPC/Installation/Bootloader/es")[Terminar](/wiki/Handbook:PPC/Installation/Finalizing/es "Handbook:PPC/Installation/Finalizing/es")[Trabajar con Gentoo](/wiki/Handbook:PPC/Full/Working/es "Handbook:PPC/Full/Working/es")[Introducción a Portage](/wiki/Handbook:PPC/Working/Portage/es "Handbook:PPC/Working/Portage/es")[Ajustes USE](/wiki/Handbook:PPC/Working/USE/es "Handbook:PPC/Working/USE/es")[Características de Portage](/wiki/Handbook:PPC/Working/Features/es "Handbook:PPC/Working/Features/es")[Sistema de guiones de inicio](/wiki/Handbook:PPC/Working/Initscripts/es "Handbook:PPC/Working/Initscripts/es")[Variables de entorno](/wiki/Handbook:PPC/Working/EnvVar/es "Handbook:PPC/Working/EnvVar/es")[Trabajar con Portage](/wiki/Handbook:PPC/Full/Portage/es "Handbook:PPC/Full/Portage/es")[Ficheros y directorios](/wiki/Handbook:PPC/Portage/Files/es "Handbook:PPC/Portage/Files/es")[Variables](/wiki/Handbook:PPC/Portage/Variables/es "Handbook:PPC/Portage/Variables/es")[Mezclar ramas de software](/wiki/Handbook:PPC/Portage/Branches/es "Handbook:PPC/Portage/Branches/es")[Herramientas adicionales](/wiki/Handbook:PPC/Portage/Tools/es "Handbook:PPC/Portage/Tools/es")[Repositorios personalizados de paquetes](/wiki/Handbook:PPC/Portage/CustomTree/es "Handbook:PPC/Portage/CustomTree/es")[Características avanzadas](/wiki/Handbook:PPC/Portage/Advanced/es "Handbook:PPC/Portage/Advanced/es")[Configuración de la red](/wiki/Handbook:PPC/Full/Networking/es "Handbook:PPC/Full/Networking/es")[Comenzar](/wiki/Handbook:PPC/Networking/Introduction/es "Handbook:PPC/Networking/Introduction/es")[Configuración avanzada](/wiki/Handbook:PPC/Networking/Advanced/es "Handbook:PPC/Networking/Advanced/es")[Configuración de red modular](/wiki/Handbook:PPC/Networking/Modular/es "Handbook:PPC/Networking/Modular/es")[Conexión inalámbrica](/wiki/Handbook:PPC/Networking/Wireless/es "Handbook:PPC/Networking/Wireless/es")[Añadir funcionalidad](/wiki/Handbook:PPC/Networking/Extending/es "Handbook:PPC/Networking/Extending/es")[Gestión dinámica](/wiki/Handbook:PPC/Networking/Dynamic/es "Handbook:PPC/Networking/Dynamic/es")

## Contents

- [1Elegir un empaquetado de stage](#Elegir_un_empaquetado_de_stage)
  - [1.1OpenRC](#OpenRC)
  - [1.2systemd](#systemd)
  - [1.3Multilibrería (32 y 64 bits)](#Multilibrer.C3.ADa_.2832_y_64_bits.29)
  - [1.4No multilibrería (64 bits puros)](#No_multilibrer.C3.ADa_.2864_bits_puros.29)
- [2Descargar el empaquetado de stage (tarball)](#Descargar_el_empaquetado_de_stage_.28tarball.29)
- [3Ajustar la Fecha/Hora correcta](#Ajustar_la_Fecha.2FHora_correcta)
  - [3.1Automático](#Autom.C3.A1tico)
  - [3.2Manual](#Manual)
  - [3.3Navegadores gráficos](#Navegadores_gr.C3.A1ficos)
  - [3.4Navegadores en la línea de órdenes](#Navegadores_en_la_l.C3.ADnea_de_.C3.B3rdenes)
  - [3.5Verificar y validar](#Verificar_y_validar)
- [4Instalar un archivo de stage](#Instalar_un_archivo_de_stage)
- [5Configurar las opciones de compilación](#Configurar_las_opciones_de_compilaci.C3.B3n)
  - [5.1Introducción](#Introducci.C3.B3n)
  - [5.2CFLAGS y CXXFLAGS](#CFLAGS_y_CXXFLAGS)
  - [5.3RUSTFLAGS](#RUSTFLAGS)
  - [5.4MAKEOPTS](#MAKEOPTS)
  - [5.5Preparados, listos, ¡ya!](#Preparados.2C_listos.2C_.C2.A1ya.21)
- [6Referencias](#Referencias)

### Elegir un empaquetado de stage

**Consejo**

En las arquitecturas compatibles, se recomienda que los usuarios que se orientan a un entorno de sistema operativo de escritorio (gráfico) utilicen un archivo de stage con el término `desktop` dentro del nombre. Estos archivos incluyen paquetes como [sys-devel/llvm](https://packages.gentoo.org/packages/sys-devel/llvm) y [dev-lang/rust-bin](https://packages.gentoo.org/packages/dev-lang/rust-bin) y el ajuste del indicador USE que mejorará enormemente el tiempo de instalación.

El [archivo de stage](/wiki/Stage_file "Stage file") actúa como semilla de una instalación de Gentoo. Los archivos de stage son generados con [Catalyst](/wiki/Catalyst "Catalyst") por el [Equipo de Ingeniería de Lanzamiento](/wiki/Project:RelEng "Project:RelEng"). Los archivos de stage se basan en [perfiles](/wiki/Profile_(Portage) "Profile (Portage)") específicos y contienen un sistema casi completo.

Al elegir un archivo de stage, es importante elegir uno con perfiles correspondientes al tipo de sistema deseado.

**Importante**

Si bien es posible realizar cambios importantes en el perfil después de que se haya establecido una instalación, el cambio requiere un considerable esfuerzo y consideraciones y está fuera del alcance de este manual de instalación. Cambiar los sistemas de inicio es difícil, pero cambiar de `no-multilib` a `multilib` requiere un amplio conocimiento de Gentoo y de la cadena de herramientas de bajo nivel.

**Consejo**

La mayoría de los usuarios no deberían necesitar utilizar las opciones de archivos comprimidos 'avanzadas'; son para configuraciones de software o hardware atípicas o avanzadas.

#### OpenRC

[OpenRC](/wiki/OpenRC/es "OpenRC/es") es un sistema de inicio basado en dependencias (responsable de iniciar los servicios del sistema una vez que se ha iniciado el núcleo) que mantiene compatibilidad con el programa de inicio proporcionado por el sistema, que normalmente se encuentra en /sbin/init. Es el sistema de inicialización nativo y original de Gentoo, pero también lo implementan algunas otras distribuciones de Linux y sistemas BSD.

#### systemd

systemd es un reemplazo moderno de init y rc de estilo SysV para sistemas Linux. La mayoría de las distribuciones Linux lo utilizan como sistema de inicio principal. systemd es totalmente compatible con Gentoo y funciona para el propósito previsto. Si aún le parece que falta algo en el Manual para ser una guia de instalación con systemd, revise el [artículo de systemd](/wiki/Systemd/es "Systemd/es") _antes_ de solicitar ayuda.

#### Multilibrería (32 y 64 bits)

**Nota**

No todas las arquitecturas tienen una opción multilib. Muchas solo ejecutan código nativo. Multilib se aplica más comúnmente a **amd64**.

El perfil multilib utiliza bibliotecas de 64 bits cuando es posible y solo recurre a las versiones de 32 bits cuando es estrictamente necesario por motivos de compatibilidad. Esta es una excelente opción para la mayoría de instalaciones porque proporciona una gran flexibilidad para la personalización en el futuro.

**Consejo**

El uso de `multilib` hace que sea más fácil cambiar de perfil más adelante, en comparación con `no-multilib`

#### No multilibrería (64 bits puros)

**Advertencia**

Los lectores que acaban de empezar con Gentoo _no_ deben elegir un empaquetado sin multilib a menos que sea absolutamente necesario. Migrar de un sistema `no-multilib` a uno `multilib` requiere un conocimiento extremadamente bueno de Gentoo y de la cadena de herramientas de nivel inferior (incluso puede hacer que nuestros [desarrolladores de la cadedena de herramientas](/wiki/Project:Toolchain "Project:Toolchain") se estremezcan un poco). No es para miedosos y está más allá del alcance de esta guía.

La selección de un archivo empaquetado no-multilib como base del sistema proporciona un entorno de sistema operativo completo de 64 bits - sin software de 32 bits. Esto efectivamente hace que la capacidad de cambiar a perfiles `multilib` sea engorrosa, aunque aún técnicamente posible.

### Descargar el empaquetado de stage (tarball)

Antes de descargar el _archivo de stage_, el directorio actual debe situarse en la ubicación de montaje utilizado para la instalación:

`root #` `cd /mnt/gentoo`

### Ajustar la Fecha/Hora correcta

Los archivos de stage generalmente se obtienen mediante HTTPS, lo que requiere una relativa precisión en la hora del sistema. La desviación del reloj puede impedir que las descargas funcionen y puede causar errores impredecibles si la hora del sistema se ajusta en una cantidad considerable después de la instalación.

La fecha y hora actuales se pueden verificar con date:

`root #` `date`

```
Mon Oct  3 13:16:22 PDT 2021

```

Si la fecha/hora mostrada tiene más de unos pocos minutos de diferencia, debe actualizarse utilizando uno de los siguientes métodos.

#### Automático

Usar [NTP](/wiki/NTP "NTP") para corregir la desviación del reloj suele ser más fácil y confiable que configurar manualmente el reloj del sistema.

chronyd, parte de [net-misc/chrony](https://packages.gentoo.org/packages/net-misc/chrony), se puede utilizar para actualizar el reloj del sistema a UTC con:

`root #` `chronyd -q`

**Importante**

Los sistemas sin un reloj en tiempo real (RTC) funcional deben sincronizar el reloj del sistema en cada inicio del sistema y, a partir de entonces, en intervalos regulares. Esto también es conveniente para sistemas con RTC, ya que la batería podría fallar y se puede acumular la desviación del reloj.

**Advertencia**

El tráfico NTP estándar en no autenticado, es importante verificar los datos de tiempo obtenidos de la red.

#### Manual

Cuando el acceso NTP no está disponible, se puede usar date para configurar manualmente el reloj del sistema.

**Nota**

Se recomienda la hora UTC para todos los sistemas Linux. Posteriormente, se define una zona horaria del sistema, que cambia el desplazamiento cuando se muestra la fecha.

El siguiente formato de argumento se utiliza para establecer la hora: `MMDDhhmmAAAA` sintaxis (' _M **es,** D **ía,** h **ora,**_ **'m** inuto y **Y** Año).

Por ejemplo, para ajustar la fecha y hora a las 13:16 horas del 3 de octubre del 2021, ejecute:

`root #` `date 100313162021`

#### Navegadores gráficos

Los usuarios que utilicen entornos con navegadores web gráficos no tendrán problema en copiar el URL de un archivo de stage desde [la sección de descargas](https://www.gentoo.org/downloads/#other-arches) del sitio web principal. Simplemente seleccione la pestaña apropiada, haga clic con el botón secundario del ratón en el archivo de stage, entonces Copiar la ruta del enlace para copiar el enlace al portapapeles, a continuación pegue el enlace para la utilidad wget en la línea de órdenes para descargar el archivo de stage:

`root #` `wget <URL_DEL_ARCHIVO_DE_STAGE_PEGADA>`

#### Navegadores en la línea de órdenes

Los usuarios de Gentoo más tradicionales o los 'históricos' que trabajen exclusivamente con la línea de órdenes puede que prefieran utilizar links ([www-client/links](https://packages.gentoo.org/packages/www-client/links)), un navegador no gráfico dirigido por menús. Para descargar un stage, navegue a la lista de servidores réplica de Gentoo de esta forma:

`root #` `links https://www.gentoo.org/downloads/mirrors/`

Para usar un proxy HTTP con links, pase la URL con la opción `-http-proxy`:

`root #` `links -http-proxy proxy.server.com:8080 https://www.gentoo.org/downloads/mirrors/`

Junto a links existe también el navegador lynx ([www-client/lynx](https://packages.gentoo.org/packages/www-client/lynx)). Al igual que links es un navegador de consola pero sin menús.

`root #` `lynx https://www.gentoo.org/downloads/mirrors/`

Si necesita pasar a través de un proxy, exporte las variables http\_proxy y ftp\_proxy:

`root #` `export http_proxy="http://proxy.server.com:port"
`

`root #` `export ftp_proxy="http://proxy.server.com:port"`

Seleccione un servidor réplica cercano. Normalmente bastará con los servidores HTTP, sin embargo también están disponibles otros protocolos. Entre en el directorio releases/ppc/autobuilds/. En él deberían aparecer todos los archivos de stage disponibles (quizá almacenados en subdirectorios con el nombre de cada subarquitectura). Seleccione uno y pulse `d` para descargarlo.

Una vez haya finalizada la descarga del archivo de stage, es posible verificar la integridad y validar los contenidos del archivo de stage. Los interesados pueden ir a la [siguiente sección](/wiki/Handbook:PPC/Installation/Stage/es#Verifying_and_validating "Handbook:PPC/Installation/Stage/es").

Aquellos que no estén interesados en verificar y validar el archivo de stage pueden cerrar el navegador de línea de comandos presionando `q` e ir directamente a la sección [Descomprimir el archivo de stage](/wiki/Handbook:PPC/Installation/Stage/es#Unpacking_the_stage_file "Handbook:PPC/Installation/Stage/es").

#### Verificar y validar

Al igual que con los CDs minimalistas de instalación, hay descargas disponibles para verificar y validar el archivo stage. Aunque estos pasos se pueden omitir, estos archivos se ofrecen a aquéllos usuarios que se preocupan por la integridad del archivo o archivos que se acaban de descargar. Los archivos adicionales están disponibles en la raíz del directorio mirrors. Busque la ubicación adecuada para la arquitectura del hardware y el perfil del sistema y descargue los archivos asociados .CONTENTS.gz, .DIGESTS y .sha256.

`root #` `wget https://distfiles.gentoo.org/releases/`

- .CONTENTS.gz: un archivo comprimido que contiene una lista de todos los archivos dentro del archivo de stage.
- .DIGESTS: contiene sumas de comprobación del archivo de stage utilizando varios algoritmos hash criptográficos.
- .sha256: contiene una suma de comprobación SHA256 firmada con PGP del archivo de stage. Es posible que este archivo no esté disponible para descargar para todos los archivos de stage.


Al igual que con el archivo ISO, la firma criptográfica del archivo tar.xz se puede verificar usando gpg para garantizar que no se haya realizado ninguna manipulación en el archivo empaquetado.

Para las imágenes live oficiales de Gentoo, el paquete [sec-keys/openpgp-keys-gentoo-release](https://packages.gentoo.org/packages/sec-keys/openpgp-keys-gentoo-release) proporciona claves de firma PGP para lanzamientos automatizados. Las claves primero deben importarse a la sesión del usuario para poder utilizarlas para la verificación:

`root #` `gpg --import /usr/share/openpgp-keys/gentoo-release.asc`

Para todas las imágenes live no oficiales que ofrecen gpg y wget en el entorno live, se puede recuperar e importar un paquete que contiene claves de Gentoo:

Verifique la firma del empaquetado y, opcionalmente, los archivos de suma de comprobación asociados:

`root #` `gpg --verify stage3-ppc-<release>-<init>.tar.xz.asc
`

`root #` `gpg --verify stage3-ppc-<release>-<init>.tar.xz.DIGESTS
`

`root #` `gpg --verify stage3-ppc-<release>-<init>.tar.xz.sha256
`

Si la verificación tiene éxito, aparecerá "Good signature from" en la salida de los comandos anteriores.

Las huellas digitales de las claves OpenPGP utilizadas para firmar los medios de lanzamiento se pueden encontrar en la [página de firmas de medios de lanzamiento](https://www.gentoo.org/downloads/signatures/).

**Nota**

La mayoría de los stages ahora tienen [sufijo](https://www.gentoo.org/news/2021/07/20/more-downloads.html) explícitamente con el tipo de sistema de inicio (openrc o systemd), aunque en algunas arquitecturas aún pueden faltar por ahora.

Se pueden utilizar herramientas y utilidades criptográficas como openssl, sha256sum o sha512sum para comparar la salida con las sumas de verificación proporcionadas en el archivo .DIGESTS.

Para verificar la suma de comprobación SHA512 con openssl:

`root #` `openssl dgst -r -sha512 stage3-ppc-<release>-<init>.tar.xz`

`dgst` indica al comando openssl que utilice el subcomando Message Digest, `-r` imprime el resultado del resumen en formato coreutils y `-sha512 ` selecciona el resumen SHA512.

Para verificar la suma de comprobación BLAKE2B512 con openssl:

`root #` `openssl dgst -r -blake2b512 stage3-ppc-<release>-<init>.tar.xz`

Compare las salidas de los comandos de suma de comprobación con los valores correspondientes de hash y nombre de archivo contenidos en el archivo .DIGESTS. Los valores correspondientes deben coincidir con el resultado de los comandos de suma de comprobación; de lo contrario, el archivo descargado está dañado y debe descartarse y volverse a descargar.

Para verificar el hash SHA256 de un archivo .sha256 asociado usando la utilidad sha256sum:

`root #` `sha256sum --check stage3-ppc-<release>-<init>.tar.xz.sha256`

La opción `--check` indica a sha256sum que lea una lista de archivos esperados y hashes asociados, y luego imprima un "OK" asociado para cada archivo que se calcule correctamente o un "FAILED" para archivos en los que no ocurra.

## Instalar un archivo de stage

Una vez que el _archivo stage_ se haya descargado y verificado, se puede extraer usando tar:

`root #` `tar xpvf stage3-*.tar.xz --xattrs-include='*.*' --numeric-owner -C /mnt/gentoo`

Antes de extraer verifique las opciones:

- `x` e **x** tract, indica a tar que extraiga el contenido del archivo.
- `p` **p** reservar permisos.
- `v` salida **v'** erbose.
- `f` **f** ile, proporciona a tar el nombre del archivo de entrada.
- `--xattrs-include='*.*'` Conserva los atributos extendidos en todos los espacios de nombres almacenados en el archivo.
- `--numeric-owner` Asegúrese de que los ID de usuario y grupo de los archivos que se extraen del empaquetado sigan siendo los mismos que pretendía el equipo de ingeniería de lanzamiento de Gentoo (incluso si los usuarios aventureros no están usando entornos live oficiales de Gentoo para el proceso de instalación).
- `-C /mnt/gentoo` Extrae archivos en la partición raíz independientemente del directorio actual.

Ahora que el archivo stage está desempaquetado, continúe con [Configurara las opciones de compilación](/wiki/Handbook:PPC/Installation/Stage/es#Configuring_compile_options "Handbook:PPC/Installation/Stage/es").

## Configurar las opciones de compilación

### Introducción

Para optimizar el sistema, es posible establecer variables que afecten al comportamiento de Portage, el administrador de paquetes con soporte oficial de Gentoo. Todas esas variables se pueden configurar como variables de entorno (usando export), pero la configuración a través de export no es permanente.

**Nota**

Técnicamente, las variables se pueden exportar a través del perfil del [shell](/wiki/Shell "Shell") o archivos rc, sin embargo, esa no es la mejor manera para la administración básica del sistema.

Portage lee el archivo [make.conf](/wiki/Make.conf "Make.conf") cuando se ejecuta, lo que cambiará su comportamiento durante su ejecución dependiendo de los valores guardados en el archivo. make.conf puede considerarse el archivo de configuración principal de Portage, así que trate su contenido con cuidado.

**Consejo**

Puede encontrar una lista comentada de todas las variables posibles en /mnt/gentoo/usr/share/portage/config/make.conf.example. Se puede encontrar documentación adicional sobre make.conf ejecutando man 5 make.conf.

Para una instalación exitosa de Gentoo, solo se deben configurar las variables que se mencionan a continuación.

Use su editor favorito (en esta guía usaremos nano) para modificar las variables de optimización que discutiremos en adelante.

`root #` `nano /mnt/gentoo/etc/portage/make.conf`

Observando el archivo make.conf.example es obvio cual es su estructura: las líneas que son comentarios comienzan con `#`, el resto definen variables usando la sintaxis `VARIABLE="valor"`. Varias de estas variables se discuten a continuación.

### CFLAGS y CXXFLAGS

Las variables CFLAGS y CXXFLAGS definen los parámetros de optimización para los compiladores GCC de C y de C++ respectivamente. Aunque generalmente se definen aquí, obtendrá el máximo rendimiento si optimiza estos parámetros para cada programa por separado. La razón es que cada programa es diferente. Sin embargo, no es manejable definir estos indicadores en el archivo make.conf.

En make.conf deberá definir los parámetros de optimización que se ajusten a su sistema de forma general. No coloque parámetros experimentales en esta variable; un nivel demasiado alto de optimización puede hacer que los programas funcionen mal (cuelgues, o incluso peor, funcionamientos erróneos).

El Manual no explicará todas las opciones posibles de optimización. Si quiere conocerlas todas, lea los [Manuales en línea GNU](https://gcc.gnu.org/onlinedocs/) o la página información de gcc (info gcc). El archivo make.conf.example también contiene una gran cantidad de ejemplos e información; no olvide leerlo también.

El primer parámetro es `-march=` o `-mtune=`, el cual especifica el nombre de la arquitectura destino. Las posibles opciones se describen en el archivo make.conf.example (como comentarios). Un valor frecuentemente utilizado es _native_ ya que indica al compilador que seleccione la arquitectura destino del sistema actual (en el que se está realizando la instalación).

Seguida de esta, está el parámetro `-O` (que es una O mayúscula, no un cero), que especifica la clase optimización de gcc. Las clases posibles son s (para tamaño optimizado), 0 (cero - para no optimizar), 1, 2 o incluso 3 para la optimización de velocidad (cada clase tiene los mismos parámetros que la anterior, más algunos extras). `-O2` es la recomendación por defecto. Es conocido que `-O3` provoca problemas cuando se utiliza globalmente en el sistema, por esto se recomienda quedarse con `-O2`.

Otros parámetros de optimización bastante populares son los `-pipe` (usar tuberías en lugar de archivos temporales para la comunicación entre las diferentes etapas de compilación). No tiene ningún impacto sobre le código generado, pero usa más memoria. En sistemas con poca memoria, el proceso del compilador podría ser terminado. En ese caso, no use este parámetro.

Usar `-fomit-frame-pointer` (el cual no mantiene el puntero de marco en un registro para aquellas funciones que no lo necesiten) podría tener graves repercusiones en la depuración de errores en aplicaciones.

Cuando defina las variables CFLAGS y CXXFLAGS, debería combinar varios parámetros de optimización en una sóla cadena. Los valores por defecto que trae el archivo de stage deberían ser suficientemente buenos. Lo siguiente es simplemente un ejemplo:

CÓDIGO **Ejemplo de CFLAGS y CXXFLAGS variables**

```
# Configuraciones del compilador a aplicar en cualquier lenguaje
COMMON_CFLAGS="-O2 -mcpu=powerpc -mtune=powerpc -pipe"
# Use los mismos valores en ambas variables
CFLAGS="${COMMON_FLAGS}"
CXXFLAGS="${CFLAGS}"

```

**Consejo**

A pesar de que el artículo sobre la [optimización de GCC](/wiki/GCC_optimization/es "GCC optimization/es") contiene más información sobre cómo las distintas opciones de compilación pueden afectar a un sistema, el artículo sobre [CFLAGS seguras](/wiki/Safe_CFLAGS "Safe CFLAGS") puede resultar más práctico para los que se inician en la optimización de su sistema.

### RUSTFLAGS

Muchos programas se escriben ahora en Rust, que tiene su propio método de optimización. Por defecto, Rust optimiza al nivel 3 en todas las versiones, a menos que un proyecto lo modifique, por lo que debe dejarse como está. Una lista completa de optimizaciones (conocida como codegen) que se puede pasar al compilador rust está disponible en [https://doc.rust-lang.org/rustc/codegen-options/index.html](https://doc.rust-lang.org/rustc/codegen-options/index.html).

La optimización más útil sería decirle a Rust que compile para su CPU usando el siguiente ejemplo:

ARCHIVO **`/etc/portage/make.conf`** **Ejemplo de RUSTFLAGS**

```
RUSTFLAGS="${RUSTFLAGS} -C target-cpu=native"

```

Para obtener una lista de CPU,s. soportadas por Rust, ejecute:

`root #` `rustc -C target-cpu=help`

Esto mostrará cuál es el valor predeterminado y también qué CPU se seleccionará con la opción native.

**Nota**

El comando anterior solo funciona con archivos empaquetados stage3 que incluyan la palabra desktop o después de instalar [dev-lang/rust-bin](https://packages.gentoo.org/packages/dev-lang/rust-bin) o [dev-lang/rust](https://packages.gentoo.org/packages/dev-lang/rust).

### MAKEOPTS

La variable MAKEOPTS define cuántas compilaciones paralelas deben ocurrir al instalar un paquete. A partir de la versión 3.0.31 de Portage[\[1\]](#cite_note-1), si no se define, el comportamiento predeterminado de Portage es establecer el valor de trabajos en MAKEOPTS al mismo número de hilos de procesamiento devueltos por nproc.

Además, a partir de Portage 3.0.53[\[2\]](#cite_note-2), si no se define, el comportamiento predeterminado de Portage es establecer el valor promedio de carga en < var>MAKEOPTS al mismo número de hilos de procesamiento devueltos por nproc.

Una buena opción es el menor de: el número de hilos de procesamiento que tiene la CPU o la cantidad total de RAM del sistema dividida por 2 GiB.

**Advertencia**

El uso de una gran cantidad de trabajos puede afectar significativamente el consumo de memoria. Una buena recomendación es tener al menos 2 GiB de RAM para cada trabajo especificado (por ejemplo, `-j6` requiere _al menos_ 12 GiB). Para evitar quedarse sin memoria, reduzca el número de trabajos para que se ajusten a la memoria disponible.

**Consejo**

Cuando se utilizan emerges en paralelo ( `--jobs`), la cantidad efectiva de trabajos ejecutados puede crecer exponencialmente (hasta hacer que los trabajos se multipliquen por los trabajos de los emerges). Esto se puede solucionar ejecutando una configuración distcc solo para localhost que limitará el número de instancias del compilador por host.

ARCHIVO **`/etc/portage/make.conf`** **Ejemplo de declaración MAKEOPTS**

```
# Si no se define, el comportamiento predeterminado de Portage es:
# - establecer el valor de los trabajos en MAKEOPTS en la misma cantidad de subprocesos devueltos por `nproc`
# - establecer el valor de carga media en MAKEOPTS ligeramente por encima de la cantidad de subprocesos devueltos por `nproc`, debido a que será el valor limitante
# Reemplace '4' según corresponda para el sistema (min(RAM/2GB, subprocesos), o déjelo sin establecer.
MAKEOPTS="-j4 -l5"

```

Busque MAKEOPTS en man 5 make.conf para obtener más detalles.

### Preparados, listos, ¡ya!

Actualice /mnt/gentoo/etc/portage/make.conf con sus propios parámetros y guarde los cambios (los usuarios de nano deben presionar usar `Ctrl` + `o` para guardar los cambios y después `Ctrl` + `x` para salir).

## Referencias

1. [↑](#cite_ref-1)[https://gitweb.gentoo.org/proj/portage.git/commit/?id=5d2af567772bb12b073f1671daea6263055cbdc2](https://gitweb.gentoo.org/proj/portage.git/commit/?id=5d2af567772bb12b073f1671daea6263055cbdc2)
2. [↑](#cite_ref-2)[https://gitweb.gentoo.org/proj/portage.git/commit/?id=de7be7f45ee54e3f789def46542919550687d15e](https://gitweb.gentoo.org/proj/portage.git/commit/?id=de7be7f45ee54e3f789def46542919550687d15e)

[← Preparando los discos](/wiki/Handbook:PPC/Installation/Disks/es "Previous part") [Inicio](/wiki/Handbook:PPC/es "Handbook:PPC/es") [Instalar el sistema base →](/wiki/Handbook:PPC/Installation/Base/es "Next part")