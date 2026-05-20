# Base

Other languages:

- [Deutsch](/wiki/Handbook:SPARC/Installation/Base/de "Handbuch:SPARC/Installation/Basis (100% translated)")
- [English](/wiki/Handbook:SPARC/Installation/Base "Handbook:SPARC/Installation/Base (100% translated)")
- español
- [français](/wiki/Handbook:SPARC/Installation/Base/fr "Handbook:SPARC/Installation/Base/fr (100% translated)")
- [italiano](/wiki/Handbook:SPARC/Installation/Base/it "Handbook:SPARC/Installation/Base/it (50% translated)")
- [magyar](/wiki/Handbook:SPARC/Installation/Base/hu "Handbook:SPARC/Installation/Base/hu (100% translated)")
- [polski](/wiki/Handbook:SPARC/Installation/Base/pl "Handbook:SPARC/Installation/Base (100% translated)")
- [português do Brasil](/wiki/Handbook:SPARC/Installation/Base/pt-br "Handbook:SPARC/Installation/Base/pt-br (50% translated)")
- [русский](/wiki/Handbook:SPARC/Installation/Base/ru "Handbook:SPARC/Installation/Base (100% translated)")
- [தமிழ்](/wiki/Handbook:SPARC/Installation/Base/ta "கையேடு:SPARC/நிறுவல்/அடிப்படை (100% translated)")
- [中文（中国大陆）‎](/wiki/Handbook:SPARC/Installation/Base/zh-cn "手册：SPARC/安装/安装基本系统 (100% translated)")
- [日本語](/wiki/Handbook:SPARC/Installation/Base/ja "ハンドブック:SPARC/インストール/ベース (100% translated)")
- [한국어](/wiki/Handbook:SPARC/Installation/Base/ko "Handbook:SPARC/Installation/Base/ko (100% translated)")

[Manual SPARC](/wiki/Handbook:SPARC/es "Handbook:SPARC/es")[Instalación](/wiki/Handbook:SPARC/Full/Installation/es "Handbook:SPARC/Full/Installation/es")[Acerca de la instalación](/wiki/Handbook:SPARC/Installation/About/es "Handbook:SPARC/Installation/About/es")[Elegir los medios](/wiki/Handbook:SPARC/Installation/Media/es "Handbook:SPARC/Installation/Media/es")[Configurar la red](/wiki/Handbook:SPARC/Installation/Networking/es "Handbook:SPARC/Installation/Networking/es")[Preparar los discos](/wiki/Handbook:SPARC/Installation/Disks/es "Handbook:SPARC/Installation/Disks/es")[Instalar el stage3](/wiki/Handbook:SPARC/Installation/Stage/es "Handbook:SPARC/Installation/Stage/es")Instalar el sistema base[Configurar el núcleo](/wiki/Handbook:SPARC/Installation/Kernel/es "Handbook:SPARC/Installation/Kernel/es")[Configurar el sistema](/wiki/Handbook:SPARC/Installation/System/es "Handbook:SPARC/Installation/System/es")[Instalar las herramientas](/wiki/Handbook:SPARC/Installation/Tools/es "Handbook:SPARC/Installation/Tools/es")[Configurar el cargador de arranque](/wiki/Handbook:SPARC/Installation/Bootloader/es "Handbook:SPARC/Installation/Bootloader/es")[Terminar](/wiki/Handbook:SPARC/Installation/Finalizing/es "Handbook:SPARC/Installation/Finalizing/es")[Trabajar con Gentoo](/wiki/Handbook:SPARC/Full/Working/es "Handbook:SPARC/Full/Working/es")[Introducción a Portage](/wiki/Handbook:SPARC/Working/Portage/es "Handbook:SPARC/Working/Portage/es")[Ajustes USE](/wiki/Handbook:SPARC/Working/USE/es "Handbook:SPARC/Working/USE/es")[Características de Portage](/wiki/Handbook:SPARC/Working/Features/es "Handbook:SPARC/Working/Features/es")[Sistema de guiones de inicio](/wiki/Handbook:SPARC/Working/Initscripts/es "Handbook:SPARC/Working/Initscripts/es")[Variables de entorno](/wiki/Handbook:SPARC/Working/EnvVar/es "Handbook:SPARC/Working/EnvVar/es")[Trabajar con Portage](/wiki/Handbook:SPARC/Full/Portage/es "Handbook:SPARC/Full/Portage/es")[Ficheros y directorios](/wiki/Handbook:SPARC/Portage/Files/es "Handbook:SPARC/Portage/Files/es")[Variables](/wiki/Handbook:SPARC/Portage/Variables/es "Handbook:SPARC/Portage/Variables/es")[Mezclar ramas de software](/wiki/Handbook:SPARC/Portage/Branches/es "Handbook:SPARC/Portage/Branches/es")[Herramientas adicionales](/wiki/Handbook:SPARC/Portage/Tools/es "Handbook:SPARC/Portage/Tools/es")[Repositorios personalizados de paquetes](/wiki/Handbook:SPARC/Portage/CustomTree/es "Handbook:SPARC/Portage/CustomTree/es")[Características avanzadas](/wiki/Handbook:SPARC/Portage/Advanced/es "Handbook:SPARC/Portage/Advanced/es")[Configuración de la red](/wiki/Handbook:SPARC/Full/Networking/es "Handbook:SPARC/Full/Networking/es")[Comenzar](/wiki/Handbook:SPARC/Networking/Introduction/es "Handbook:SPARC/Networking/Introduction/es")[Configuración avanzada](/wiki/Handbook:SPARC/Networking/Advanced/es "Handbook:SPARC/Networking/Advanced/es")[Configuración de red modular](/wiki/Handbook:SPARC/Networking/Modular/es "Handbook:SPARC/Networking/Modular/es")[Conexión inalámbrica](/wiki/Handbook:SPARC/Networking/Wireless/es "Handbook:SPARC/Networking/Wireless/es")[Añadir funcionalidad](/wiki/Handbook:SPARC/Networking/Extending/es "Handbook:SPARC/Networking/Extending/es")[Gestión dinámica](/wiki/Handbook:SPARC/Networking/Dynamic/es "Handbook:SPARC/Networking/Dynamic/es")

## Contents

- [1Enjaulamiento](#Enjaulamiento)
  - [1.1Copiar la información DNS](#Copiar_la_informaci.C3.B3n_DNS)
  - [1.2Montar los sistemas de archivos necesarios](#Montar_los_sistemas_de_archivos_necesarios)
  - [1.3Entrar en el nuevo entorno](#Entrar_en_el_nuevo_entorno)
  - [1.4Preparándose para un gestor de arranque](#Prepar.C3.A1ndose_para_un_gestor_de_arranque)
    - [1.4.1Sistemas DOS/Legacy BIOS](#Sistemas_DOS.2FLegacy_BIOS)
- [2Configurar Portage](#Configurar_Portage)
  - [2.1Instalar una instantánea de repositorio de ebuilds de Gentoo desde la web](#Instalar_una_instant.C3.A1nea_de_repositorio_de_ebuilds_de_Gentoo_desde_la_web)
  - [2.2Opcional: Seleccionar los servidores réplica](#Opcional:_Seleccionar_los_servidores_r.C3.A9plica)
    - [2.2.1Optional: rsync mirrors](#Optional:_rsync_mirrors)
  - [2.3Opcional: Actualizar el repositorio de ebuilds de Gentoo](#Opcional:_Actualizar_el_repositorio_de_ebuilds_de_Gentoo)
  - [2.4Leer los elementos de noticias](#Leer_los_elementos_de_noticias)
  - [2.5Elegir el perfil adecuado](#Elegir_el_perfil_adecuado)
  - [2.6Opcional: agregar un host de paquetes binarios](#Opcional:_agregar_un_host_de_paquetes_binarios)
    - [2.6.1Configuración del repositorio](#Configuraci.C3.B3n_del_repositorio)
    - [2.6.2Instalación de paquetes binarios](#Instalaci.C3.B3n_de_paquetes_binarios)
  - [2.7Opcional: configurar la variable USE](#Opcional:_configurar_la_variable_USE)
    - [2.7.1CPU\_FLAGS\_\*](#CPU_FLAGS_.2A)
    - [2.7.2VIDEO\_CARDS](#VIDEO_CARDS)
  - [2.8Opcional: Configurar la variable ACCEPT\_LICENSE](#Opcional:_Configurar_la_variable_ACCEPT_LICENSE)
  - [2.9Actualizar el conjunto @world](#Actualizar_el_conjunto_.40world)
    - [2.9.1Eliminar paquetes obsoletos](#Eliminar_paquetes_obsoletos)
- [3Zona horaria](#Zona_horaria)
- [4Configurar localizaciones](#Configurar_localizaciones)
  - [4.1Generación de localizaciones](#Generaci.C3.B3n_de_localizaciones)
  - [4.2Selección de localización](#Selecci.C3.B3n_de_localizaci.C3.B3n)
- [5Referencias](#Referencias)

## Enjaulamiento

### Copiar la información DNS

Aún queda una cosa que hacer antes de entrar en el nuevo entorno, copiar la información sobre los DNS en /etc/resolv.conf. Necesita hacer esto para asegurarse de que la red continúe funcionando después de entrar en el nuevo entorno. /etc/resolv.conf contiene los servidores de nombres para su red.

Para copiar esta información, se recomienda pasar la opción `--dereference` en la orden cp.Esto asegura que, si /etc/resolv.conf es un enlace simbólico, se copia el archivo al que apunta el enlace y no el propio enlace. En caso contrario, en el nuevo entorno, el enlace simbólico podría apuntar a un archivo inexistente (ya que lo mas probable es que los archivos apuntados no estén disponible dentro del nuevo entorno).

`root #` `cp --dereference /etc/resolv.conf /mnt/gentoo/etc/`

### Montar los sistemas de archivos necesarios

**Consejo**

Si utiliza el medio de instalación de Gentoo, este paso se puede reemplazar con simplemente: arch-chroot /mnt/gentoo.

En unos momentos, la raíz de Linux se cambiará a la nueva ubicación.

Los sistemas de archivos que deben estar disponibles son:

- /proc/ es un pseudosistema de archivos. Parecen archivos normales, pero el kernel de Linux los genera sobre la marcha.
- /sys/ es un pseudosistema de archivos, como /proc/ que alguna vez estuvo destinado a reemplazarlo, y está más estructurado que /proc/
- /dev/ es un sistema de archivos normal que contiene todos los dispositivos. Está parcialmente administrado por el administrador de dispositivos de Linux (generalmente udev)
- /run/ es un sistema de archivos temporal utilizado para archivos generados en tiempo de ejecución, como archivos PID o bloqueos.

La ubicación /proc/ se montará en /mnt/gentoo/proc/ mientras que los otros serán montados mediante enlace. Esto último quiere decir que, por ejemplo, /mnt/gentoo/sys/ será realmente el actual /sys/ (será sólo un segundo punto de entrada al mismo sistema de archivos) mientras que /mnt/gentoo/proc/ es un nuevo montaje (instancia por así decirlo) del sistema de archivos.

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

**Nota**

Las operaciones `--make-rslave` son necesarias para dar soporte a systemd mas adelante en la instalación.

**Advertencia**

Cuando se utilicen medios de instalación que no sean de Gentoo, podría no ser suficiente. Algunas distribuciones crean el enlace simbólico /dev/shm a /run/shm/ el cual ya no será válido después del chroot. Hacer que /dev/shm/ sea un apropiado montaje tmpfs puede resolver este problema:

`root #` `test -L /dev/shm && rm /dev/shm && mkdir /dev/shm
`

`root #` `mount --types tmpfs --options nosuid,nodev,noexec shm /dev/shm
`

Asegúrese también de asignale permisos 1777:

`root #` ` chmod 1777 /dev/shm /run/shm`

### Entrar en el nuevo entorno

Ahora que todas las particiones están inicializadas y el sistema base instalado, es hora de entrar en el nuevo entorno de instalación haciendo chrooting en él. Esto significa que la sesión cambiará su raíz (la ubicación de mayor nivel que puede ser accedida) desde el entorno de instalación actual (CD de instalación u otro medio de instalación) hasta el sistema de instalación (es decir, las particiones inicializadas). De ahí el nombre, _change root_ (cambiar raíz) o _chroot_.

El enjaulamiento (chroot) se hace en tres pasos:

1. Se cambia la raíz desde / (en el medio de instalación) a /mnt/gentoo/ (en las particiones) utilizando chroot o arch-chroot, si está disponible.
2. Se cargan en memoria algunas definiciones (ofrecidas por /etc/profile) mediante la orden source.
3. Se redefine el símbolo de espera de órdenes (prompt) primario que nos hará recordar que nos encontramos en un entorno enjaulado (chroot).

`root #` `chroot /mnt/gentoo /bin/bash
`

`root #` `source /etc/profile
`

`root #` `export PS1="(chroot) ${PS1}"`

Desde este momento, todas las acciones realizadas lo serán en el nuevo entorno Gentoo Linux.

**Consejo**

Si la instalación de Gentoo se interrumpe en algún momento posterior, _debería_ ser posible 'continuarla' desde este paso. ¡No es necesario particionar los discos otra vez! Simplemente [monte la partición raíz](/wiki/Handbook:SPARC/Installation/Disks/es#Mounting_the_root_partition "Handbook:SPARC/Installation/Disks/es") y siga los pasos anteriores comenzando desde [copiar la información DNS](/wiki/Handbook:SPARC/Installation/Base/es#Copy_DNS_info "Handbook:SPARC/Installation/Base/es") hasta entrar en el nuevo entorno. Esto también se utiliza para arreglar problemas con el cargador de arranque. Mas información en el artículo [chroot](/wiki/Chroot/es "Chroot/es").

### Preparándose para un gestor de arranque

Ahora que se ha entrado en el nuevo entorno, es necesario prepararlo para el gestor de arranque. Será importante tener montada la partición correcta cuando llegue el momento de instalar el gestor de arranque.

#### Sistemas DOS/Legacy BIOS

Para sistemas DOS/BIOS heredado, el gestor de arranque se instalará en el directorio /boot, por lo tanto, móntelo de la siguiente manera:

`root #` `mount  /boot`

## Configurar Portage

### Instalar una instantánea de repositorio de ebuilds de Gentoo desde la web

El siguiente paso es instalar una instantánea del repositorio de ebuilds de Gentoo. Esta instantánea contiene una colección de ficheros que informa a Portage sobre los títulos de software disponibles (para su instalación), qué perfiles puede seleccionar el administrador del sistema, artículos de noticias específicas de paquetes o perfiles, etc.

Se recomienda utilizar emerge-webrsync para aquéllos que se encuentren detrás de cortafuegos restrictivos (utiliza los protocolos HTTP/FTP para descargar la instantánea) y ahorra ancho de banda de red. Los lectores que no tengan limitaciones en el ancho de banda pueden saltar a la siguiente sección.

Esto recuperará la última instantánea (que se libera todo los días) desde uno de los servidores réplica de Gentoo e instalarla en el sistema:

`root #` `emerge-webrsync`

**Nota**

Durante esta operación, emerge-webrsync podría indicar que la ubicación /var/db/repos/gentoo/ no existe. Esto es normal y no debe preocupar - la herramienta creará la ubicación.

A partir de este punto Portage podría indicar que se recomienda realizar algunas actualizaciones. Esto es debido a que algunos paquetes de sistema que se han instalado mediante un archivo stage disponen de versiones más actuales y ahora Portage detecta los nuevos paquetes consultando la instantánea del repositorio. Por el momento se pueden ignorar las actualizaciones de los paquetes hasta que la instalación de Gentoo haya finalizado.

### Opcional: Seleccionar los servidores réplica

Para descargar el código fuente rápidamente, se recomienda seleccionar un mirror rápido y geográficamente cercano. Portage buscará en el archivo make.conf la variable GENTOO\_MIRRORS y utilizará los mirrors enumerados allí. Es posible navegar a la lista de réplicas de Gentoo y buscar una réplica (o varias réplicas) cerca de la ubicación física del sistema (ya que éstas suelen ser las más rápidas).

Una herramienta llamada mirrorselect proporciona una bonita interfaz de texto para consultar y seleccionar más rápidamente los mirrors adecuados. Simplemente navegue hasta los mirrors de su elección y presione [Template:Tecla](/index.php?title=Template:Tecla&action=edit&redlink=1 "Template:Tecla (page does not exist)") para seleccionar uno o más mirrors.

`root #` `emerge --ask --verbose --oneshot app-portage/mirrorselect`

`root #` `mirrorselect -i -o >> /etc/portage/make.conf`

Alternativamente, hay una lista de réplicas activas [disponible en línea](https://www.gentoo.org/downloads/mirrors/).

#### Optional: rsync mirrors

Gentoo also has many rsync mirrors which can speed up downloading the available package list using `emerge --sync` (explained in more detail later) by selecting a mirror closer that is geographically closer to the user.

`root #` `mkdir /etc/portage/repos.conf
`

`root #` `cp /usr/share/portage/config/repos.conf /etc/portage/repos.conf/gentoo.conf
`

Select a mirror close to the system's location from [https://www.gentoo.org/support/rsync-mirrors/](https://www.gentoo.org/support/rsync-mirrors/) and replace the sync-uri default mirror in /etc/portage/repos.conf/gentoo.conf with the desired mirror location.

ARCHIVO **`/etc/portage/repos.conf/gentoo.conf`** **UK-based rsync mirror example**

```
[DEFAULT]
main-repo = gentoo
</div>

<div lang="en" dir="ltr" class="mw-content-ltr">
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

### Opcional: Actualizar el repositorio de ebuilds de Gentoo

Puede actualizar el repositorio de ebuilds de Gentoo a la última versión. La orden emerge-webrsync anterior habrá instalado una instantanea muy reciente (normalmente inferior a 24 horas) de manera que claramente este paso es opcional.

Supongamos que es necesario disponer de las últimas actualizaciones de un paquete (hasta 1 hora), entonces utilice emerge --sync. Este comando utilizará el protocolo rsync para actualizar el repositorio ebuild de Gentoo (que se obtuvo anteriormente a través de emerge-webrsync) al estado más reciente.

`root #` `emerge --sync`

En terminales lentos, como ciertos frame buffers o consolas serie, se recomienda utilizar la opción `--quiet` para acelerar el proceso:

`root #` `emerge --sync --quiet`

### Leer los elementos de noticias

Cuando se sincroniza el repositorio de ebuilds de Gentoo, Portage puede mostrar mensajes informativos de salida similares a los siguientes:

` * IMPORTANT: 2 news items need reading for repository 'gentoo'.
`

` * Use eselect news to read news items.
`

Los elementos de noticias se crearon para ofrecer un medio de comunicación en el que se incluyeran mensajes críticos a los usuarios a través del repositorio de ebuilds de Gentoo. Para gestionarlos necesitará utilizar eselect news. La aplicación eselect es una utilidad específica de Gentoo que presenta una interfaz de gestión común para la administración del sistema. En este caso se ha pedido a eselect que use el módulo `news`.

Para el módulo `news` hay tres operaciones que son las mas usadas:

- Con `list` se muestra un sumario de las noticias disponibles
- Con `read` se leen las noticas
- Con `purge` se pueden eliminar noticias que una vez leidas ya no sea necesario volverlas a leer mas.

`root #` `eselect news list
`

`root #` `eselect news read`

Se puede obtener más información sobre el lector de noticias en la página del manual:

`root #` `man news.eselect`

### Elegir el perfil adecuado

**Consejo**

Los perfiles de escritorio no son exclusivos para _entornos de escritorio_. También son adecuados para administradores de ventanas mínimos como i3 o sway.

Un _perfil_ (profile) es uno de los bloques de construcción en cualquier sistema Gentoo. No solamente especifica unos valores predeterminados para USE, CFLAGS , y otras variables importantes, también bloquea sistema para ciertos rangos de versiones de algunos paquetes. Estas configuraciones son todas mantenidas por los desarrolladores de Gentoo.

Se puede ver el perfil que el sistema está utilizado actualmente con eselect, en este caso usando el módulo `profile`:

**Consejo**

On an install media without a scroll-able terminal, `eselect profile list | less` can provide an easy way to list all available profiles while providing the ability to scroll.

`root #` `eselect profile list`

```
Available profile symlink targets:
  [1]   default/linux/sparc/23.0 *
  [2]   default/linux/sparc/23.0/desktop
  [3]   default/linux/sparc/23.0/desktop/gnome
  [4]   default/linux/sparc/23.0/desktop/kde

```

**Nota**

La salida de la orden anterior es solo un ejemplo y cambiará a lo largo del tiempo.

**Nota**

Para usar **systemd**, seleccione un perfil que tenga "systemd" en el nombre y viceversa, si no

También hay subperfiles de escritorio disponibles para algunas arquitecturas que incluyen paquetes de software comúnmente necesarios para una experiencia de escritorio.

**Advertencia**

Las actualizaciones de perfil no deben tomarse a la ligera. Al seleccionar el perfil inicial, utilice el perfil correspondiente a **la misma versión** que el utilizado inicialmente por el archivo de stage (por ejemplo, 23.0). Cada nueva versión del perfil se anuncia a través de una noticia que contiene instrucciones de migración; asegúrese de seguir cuidadosamente las instrucciones antes de cambiar a un perfil más nuevo.

Después de revisar los perfiles disponibles para la arquitectura sparc, los usuarios pueden elegir usar un perfil diferente para el sistema:

`root #` `eselect profile set 2`

**Nota**

El subperfil `developer` es específico para desarrolladores de Gentoo Linux y no está dirigido a usuarios comunes.

### Opcional: agregar un host de paquetes binarios

Desde diciembre de 2023, el [equipo de Ingeniería de Lanzamiento](/wiki/Project:RelEng "Project:RelEng") de Gentoo ha ofrecido un [host de paquetes binarios oficial](/wiki/Binary_package_quickstart "Binary package quickstart") (coloquialmente abreviado simplemente como "binhost") para que lo utilice la comunidad general para recuperar e instalar. paquetes binarios (binpkgs).[\[1\]](#cite_note-1)

Agregar un host de paquetes binarios permite a Portage instalar paquetes compilados y firmados criptográficamente. En muchos casos, agregar un host de paquetes binarios disminuirá "considerablemente" el tiempo medio para la instalación del paquete y agregará muchos beneficios cuando se ejecuta Gentoo en sistemas más antiguos, más lentos o con poca potencia.

#### Configuración del repositorio

La configuración del repositorio para un binhost se encuentra en el directorio /etc/portage/binrepos.conf/ de Portage, que funciona de manera similar a la configuración mencionada en la sección [repositorio de ebuilds de Gentoo](/wiki/Handbook:SPARC/Installation/Base/es#Gentoo_ebuild_repository "Handbook:SPARC/Installation/Base/es").

Al definir un host de binarios, hay dos aspectos importantes a considerar:

1. La arquitectura y perfiles dentro del valor `sync-uri` _sí_ importan y deben alinearse con la arquitectura de ordenadores respectiva ( **sparc** en este caso) y el perfil del sistema seleccionado en la sección [Elegir el perfil correcto](/wiki/Handbook:SPARC/Installation/Base/es#Choosing_the_right_profile "Handbook:SPARC/Installation/Base/es").
2. Seleccionar un mirror rápido y geográficamente cercano generalmente acortará el tiempo de recuperación. Revise la herramienta mirrorselect mencionada en la sección [Opcional: Seleccionar mirrors](/wiki/Handbook:SPARC/Installation/Base/es#Gentoo_ebuild_repository "Handbook:SPARC/Installation/Base/es") o revise el [lista en línea de mirrors](https://www.gentoo.org/downloads/mirrors/) donde se pueden descubrir los valores de URL.


ARCHIVO **`/etc/portage/binrepos.conf/gentoobinhost.conf`** **Ejemplo de host de paquetes binarios basado en CDN**

```
[binhost]
priority = 9999
sync-uri = https://distfiles.gentoo.org/releases/<arch>/binpackages/<profile>/x86-64/

```

1. Introduced in portage-3.0.74 for per-repo verification choices

verify-signature = true

1. Default value with >=portage-3.0.77

location = /var/cache/binhost/gentoo
}}

#### Instalación de paquetes binarios

Portage compilará paquetes desde el código fuente de forma predeterminada. Se le puede indicar que utilice paquetes binarios de las siguientes maneras:

1. La opción `--getbinpkg` se puede pasar al invocar el comando emerge. Este método de instalación de paquetes binarios es útil para instalar sólo un paquete binario en particular.
2. Cambiar el valor predeterminado del sistema a través de la variable FEATURES de Portage, que se expone a través del archivo /etc/portage/make.conf. La aplicación de este cambio de configuración hará que Portage consulte al host del paquetes binarios los paquetes que se solicitarán y volverá a compilar localmente cuando no se encuentren resultados.

Por ejemplo, para que Portage instale siempre los paquetes binarios disponibles:

ARCHIVO **`/etc/portage/make.conf`** **Configurar Portage para utilizar paquetes binarios de forma predeterminada**

```
# Agregar getbinpkg a la lista de valores dentro de la variable FEATURES
FEATURES="${FEATURES} getbinpkg"
# Requerir firmas
FEATURES="${FEATURES} binpkg-request-signature"

```

Ejecute también getuto para que Portage configure el conjunto de claves necesario para la verificación:

`root #` `getuto`

Las funciones adicionales de Portage se analizarán en el [siguiente capítulo](/wiki/Handbook:SPARC/Working/Features/es#Portage_features "Handbook:SPARC/Working/Features/es") del manual.

### Opcional: configurar la variable USE

La variable USE es una de las más importantes que Gentoo proporciona a sus usuarios. Muchos programas se pueden compilar con o sin soporte opcional para ciertas cosas. Por ejemplo, algunos programas se pueden compilar con soporte para GTK+, o con soporte para QT. Otros programas se pueden compilar con o sin soporte SSL. Algunos programas incluso se pueden compilar con soporte framebuffer (svgalib) en lugar de soporte X11 (servidor X).

La mayoría de las distribuciones compilan sus paquetes con tanto soporte como sea posible, aumentando el tamaño de los programas y el tiempo de inicio, sin mencionar una enorme cantidad de dependencias. Con Gentoo, los usuarios pueden definir para qué opciones se debe compilar un paquete. Aquí es donde entra en juego USE.

En la variable USE se definen palabras clave que son transformadas en opciones de compilación. Por ejemplo `ssl` compilará los programas que lo requieran con soporte SSL. `-X` quitara el soporte para el servidor X (nótese el signo menos delante). `gnome gtk -kde -qt5` compilará los programas con soporte para GNOME (y GTK), pero sin soporte para KDE (y Qt), haciendo su sistema completamente adaptado para GNOME (siempre que se use una arquitectura que lo soporte).

La configuración USE predeterminada se coloca en los archivos make.defaults del perfil de Gentoo utilizado por el sistema. Gentoo utiliza un complejo sistema de herencia para los perfiles del sistema, que no se tratará en profundidad durante el proceso de instalación. La forma más sencilla de comprobar la configuración USE actualmente activa es ejecutar emerge --info y seleccionar la línea que comienza con USE:

`root #` `emerge --info | grep ^USE`

```
USE="X acl alsa amd64 berkdb bindist bzip2 cli cracklib crypt cxx dri ..."

```

**Nota**

El ejemplo anterior está truncado la lista actual de valores para USE es mucho mucho mas larga.

Puede encontrar una descripción completa sobre la variable USE en el propio sistema en /var/db/repos/gentoo/profiles/use.desc.

`root #` `less /var/db/repos/gentoo/profiles/use.desc`

Dentro de la orden less puede desplazarse arriba y abajo utilizando las teclas `↑` y `↓` y salir pulsando `q`.

Como ejemplo, se muestran algunas opciones USE para un sistema basado en KDE con DVD, ALSA y soporte para grabar CDs:

`root #` `nano /etc/portage/make.conf`

ARCHIVO **`/etc/portage/make.conf`** **Configurar ajuste para KDE/Plasma con soporte para DVD, ALSA y grabación de CDs**

```
USE="-gtk -gnome qt5 kde dvd alsa cdr"

```

Cuando se define un valor USE en /etc/portage/make.conf, se "agrega" a la lista de valores USE del sistema. Los valores USE se pueden _eliminar_ globalmente agregando un signo menos `-` delante del valor en la lista. Por ejemplo, para deshabilitar la compatibilidad con entornos gráficos X, se puede configurar `-X`:

ARCHIVO **`/etc/portage/make.conf`** **Ignorar valores por defecto para USE**

```
USE="-X acl alsa"

```

**Advertencia**

Aunque es posible, configurar `-*` (que deshabilitará todos los valores USE excepto los especificados en make.conf) está "firmemente" desaconsejado y es imprudente. Los desarrolladores de ebuilds eligen ciertos valores USE predeterminados en ebuilds para evitar conflictos, mejorar la seguridad y evitar errores, entre otras razones. Deshabilitar "todos" los valores USE anulará el comportamiento predeterminado y puede causar problemas importantes.

#### CPU\_FLAGS\_\*

Algunas arquitecturas (incluidas AMD64/X86, ARM, PPC) tienen una variable [USE\_EXPAND](/wiki/USE_EXPAND "USE EXPAND") llamada [CPU\_FLAGS\_<ARCH>](/wiki/CPU_FLAGS_X86 "CPU FLAGS X86"), donde <ARCH> se reemplaza con el nombre de la arquitectura del sistema relevante.

**Importante**

¡No te confundas! Los sistemas **AMD64** y **X86** comparten una arquitectura común, por lo que el nombre de variable adecuado para los sistemas **AMD64** es CPU\_FLAGS\_X86.

Se usa para configurar la construcción para compilar en código ensamblador específico u otro intrínseco, generalmente escritos a mano o de otra manera adicional,
y **no** es lo mismo que pedirle al compilador que genere un código optimizado para una determinada característica de la CPU (por ejemplo, `-march=`).

Los usuarios deben establecer esta variable además de configurar su COMMON\_FLAGS como deseen.

Se necesitan algunos pasos para configurar esto:

`root #` `emerge --ask --oneshot app-portage/cpuid2cpuflags`

Inspeccione la salida manualmente si tiene curiosidad:

`root #` `cpuid2cpuflags`

Luego copie la salida en package.use:

`root #` `echo "*/* $(cpuid2cpuflags)" > /etc/portage/package.use/00cpu-flags`

#### VIDEO\_CARDS

By default a profile already sets a few video cards. For many reasons this is not ideal, but a useful safety net.

To configure the system correctly the user needs to first unset the preset video cards with `VIDEO_CARDS: -*` then set the correct card for that system.

ARCHIVO **`/etc/portage/package.use/00video_cards`** **Example**

```
*/* VIDEO_CARDS: -* <GPU DRIVER NAME>

```

A continuación se muestra una tabla que se puede utilizar para comparar fácilmente los diferentes tipos de tarjetas de video con su respectivo valor para `VIDEO_CARDS`.

Equipo
Tarjeta gráfica concreta
VIDEO\_CARDS
Intel x86NoneVer [Intel#Feature support](/wiki/Intel#Feature_support "Intel")x86/ARMNvidia`nvidia`AnyNvidia excepto Maxwell, Pascal y Volta`nouveau`AnyAMD desde la serie Sea Islands`amdgpu radeonsi`AnyATI y AMD antiguasVer [radeon#Feature support](/wiki/Radeon#Feature_support "Radeon")AnyIntel`intel`Raspberry PiN/A`vc4`QEMU/KVMAny`virgl`WSLAny`d3d12`

A continuación se muestra un ejemplo de configuración correcta de package.use para _VIDEO\_CARDS_. Sustituya el nombre del controlador o controladores que se utilizarán.

ARCHIVO **`/etc/portage/package.use/00video_cards`**

```
*/* VIDEO_CARDS: amdgpu radeonsi

```

ARCHIVO **`/etc/portage/package.use/00video_cards`** **Intel example**

```
*/* VIDEO_CARDS: -* intel

```

ARCHIVO **`/etc/portage/package.use/00video_cards`** **Nvidia example**

```
*/* VIDEO_CARDS: -* nvidia

```

Los detalles de varias GPU se pueden encontrar en los artículos [AMDGPU](/wiki/AMDGPU "AMDGPU"), [Intel](/wiki/Intel "Intel"), [Nouveau (Open Source)](/wiki/Nouveau/es "Nouveau/es") o [NVIDIA (Propietario)](/wiki/NVIDIA/es "NVIDIA/es").

### Opcional: Configurar la variable ACCEPT\_LICENSE

A partir de la Propuesta de mejora 23 de Gentoo Linux (GLEP 23), se creó un mecanismo para permitir a los administradores de sistemas la capacidad de "regular el software que instalan con respecto a las licencias... Algunos quieren un sistema libre de cualquier software que no esté aprobado por OSI; otros simplemente sienten curiosidad por saber qué licencias aceptan implícitamente."[\[2\]](#cite_note-2) Con la motivación de tener un control más granular sobre del tipo de software que se ejecuta en un sistema Gentoo, nació la variable ACCEPT\_LICENSE.

Durante el proceso de instalación, Portage considera los valores establecidos dentro de la variable ACCEPT\_LICENSE para determinar si los paquetes solicitados cumplen con la determinación del administrador del sistema de una licencia aceptable. Aquí radica un problema: el repositorio de ebuilds de Gentoo está lleno de _miles_ de ebuilds, lo que resulta en [_cientos_ de distintas licencias de software](https://gitweb.gentoo.org/repo/gentoo.git/tree/licenses)... ¿Implica esto que el administrador del sistema apruebe individualmente todas y cada una de las nuevas licencias de software? Afortunadamente no; GLEP 23 también describe una solución a este problema, un concepto llamado grupos de licencias.

Para facilitar la administración del sistema, se han agrupado licencias de software que son legalmente similares, cada una según su tipo. Las definiciones de los grupos de licencias están [disponibles para su visualización](https://gitweb.gentoo.org/repo/gentoo.git/tree/profiles/license_groups) y son administradas por el [proyecto Licencias Gentoo](/wiki/Project:Licenses "Project:Licenses"). Los grupos de licencias están precedidos sintácticamente por un símbolo `@`, mientras que una licencia individual no lo está, lo que permite distinguirlos fácilmente en la variable ACCEPT\_LICENSE.

Algunos grupos de licencias comunes incluyen:

Una lista de licencias de software agrupadas según su tipo.
NombreDescripción
`@GPL-COMPATIBLE`Licencias compatibles con GPL aprobadas por la Free Software Foundation [\[a\_license 1\]](#cite_note-3)`@FSF-APPROVED`Licencias de software libre aprobadas por la FSF (incluye `@GPL-COMPATIBLE`)
`@OSI-APPROVED`Licencias aprobadas por la Open Source Initiative [\[a\_license 2\]](#cite_note-4)`@MISC-FREE`Licencias varias que probablemente sean software libre, es decir, que sigan la Definición de software libre [\[a\_license 3\]](#cite_note-5) pero no están aprobadas ni por la FSF ni por la OSI
`@FREE-SOFTWARE`Combina `@FSF-APPROVED`, `@OSI-APPROVED`, y `@MISC-FREE`.
`@FSF-APPROVED-OTHER`Licencias aprobadas por la FSF para "documentación libre" y "trabajos de uso práctico además de software y documentación" (incluidas las fuentes tipográficas)
`@MISC-FREE-DOCS`Licencias varias para documentos libres y otros trabajos (incluidas fuentes tipográficas) que siguen la definición libre [\[a\_license 4\]](#cite_note-6) pero NO están listadas en `@FSF-APPROVED-OTHER`.
`@FREE-DOCUMENTS`Combina `@FSF-APPROVED-OTHER` y `@MISC-FREE-DOCS`.
`@FREE`Metaconjunto de todas las licencias con libertad de usar, compartir, modificar y compartir modificaciones. Combina `@FREE-SOFTWARE` y `@FREE-DOCUMENTS`.
`@BINARY-REDISTRIBUTABLE`Licencias que permitan al menos la redistribución gratuita del software en formato binario. Incluye `@FREE`.
`@EULA`Acuerdos de licencia que intentan quitarle sus derechos. Son más restrictivos que "todos los derechos reservados" o requieren aprobación explícita.

1. [↑](#cite_ref-3)[https://www.gnu.org/licenses/license-list.html](https://www.gnu.org/licenses/license-list.html)
2. [↑](#cite_ref-4)[https://www.opensource.org/licenses](https://www.opensource.org/licenses)
3. [↑](#cite_ref-5)[https://www.gnu.org/philosophy/free-sw.html](https://www.gnu.org/philosophy/free-sw.html)
4. [↑](#cite_ref-6)[https://freedomdefined.org/](https://freedomdefined.org/)

Los valores de licencia aceptables actualmente establecidos en todo el sistema se pueden ver a través de:

`user $` `portageq envvar ACCEPT_LICENSE`

```
@FREE
```

Como se ve en el resultado, el valor predeterminado es permitir que solo se instale el software que se ha agrupado en la categoría `@FREE`.

Se pueden definir licencias específicas o grupos de licencias para un sistema en las siguientes ubicaciones:

- Para todo el sistema dentro del perfil seleccionado: esto establece el valor predeterminado.
- Para todo el sistema dentro del archivo /etc/portage/make.conf. Los administradores del sistema anulan el valor predeterminado del perfil con el contenido de este archivo.
- Por paquete dentro del archivo /etc/portage/package.license.
- Por paquete dentro de un _directorio_ de archivos [Template:Ruta](/index.php?title=Template:Ruta&action=edit&redlink=1 "Template:Ruta (page does not exist)").

La licencia predeterminada para todo el sistema en el perfil se anula con el contenido de /etc/portage/make.conf:

ARCHIVO **`/etc/portage/make.conf`** **Aceptar licencias con ACCEPT\_LICENSE para todo el sistema**

```
# Anula el valor predeterminado ACCEPT_LICENSE del perfil
ACCEPT_LICENSE="-* @FREE @BINARY-REDISTRIBUTABLE"

```

Opcionalmente, los administradores del sistema también pueden definir licencias aceptadas por paquete como se muestra en el siguiente directorio de archivos de ejemplo. Tenga en cuenta que será necesario crear el _directorio_ package.license si aún no existe:

`root #` `mkdir /etc/portage/package.license`

Los detalles de la licencia de software para un paquete Gentoo individual se almacenan dentro de la variable LICENSE del ebuild asociado. Un paquete puede tener una o varias licencias de software, por lo que podría ser necesario especificar varias licencias aceptables para un único paquete.

ARCHIVO **`/etc/portage/package.license/kernel`** **Aceptar licencias por paquete**

```
app-arch/unrar unRAR
sys-kernel/linux-firmware linux-fw-redistributable
sys-firmware/intel-microcode intel-ucode

```

**Importante**

La variable LICENSE en un ebuild es sólo una indicación para los desarrolladores y usuarios de Gentoo. _No_ es una declaración legal y _no hay garantía_ de que refleje la realidad. Se recomienda no confiar únicamente en la interpretación que haga el desarrollador de ebuilds de la licencia de un paquete de software; sino que verifique el paquete en profundidad, incluidos todos los archivos que se han instalado en el sistema.

### Actualizar el conjunto @world

**Consejo**

Si se seleccionó un perfil objetivo de entorno de escritorio desde un archivo de stage que no es de escritorio, el proceso de actualización de @world podría extender en gran medida la cantidad de tiempo necesario para el proceso de instalación. Aquellos que tienen poco tiempo pueden actuar según esta 'regla general': cuanto más corto sea el nombre del perfil, menos específico será el [conjunto @world](/wiki/World_set_(Portage) "World set (Portage)") del sistema. Cuanto menos específico sea el conjunto @world, menos paquetes requerirá el sistema. P.ej.:

- Seleccionar `default/linux/amd64/23.0` probablemente requerirá que se actualicen menos paquetes, mientras que
- Seleccionar `default/linux/amd64/23.0/desktop/gnome/systemd` probablemente requerirá la instalación de más paquetes ya que el perfil de destino tiene unos conjuntos [@system](/wiki/Package_sets#.40system "Package sets") y [@profile](/wiki/Package_sets#.40profile "Package sets") mas grandes: dependencias que dan soporte al entorno de escritorio GNOME.

La actualización del [conjunto @world](/wiki/World_set_(Portage) "World set (Portage)") del sistema es _opcional_ y es poco probable que se realicen cambios funcionales a menos que se hayan realizado uno o más de los siguientes pasos opcionales:

1. Se ha seleccionado un perfil objetivo _diferente_ del archivo de stage.
2. Se han configurado indicadores USE adicionales para los paquetes instalados.

Los lectores que estén realizando una 'ejecución rápida de instalación de Gentoo' pueden omitir con seguridad las actualizaciones del conjunto @world hasta _después_ de que su sistema se haya reiniciado en el nuevo entorno de Gentoo.

Los lectores que estén realizando una ejecución lenta pueden hacer que Portage realice actualizaciones para cambios de paquetes, perfiles y/o indicadores USE en este momento:

`root #` `emerge --ask --verbose --update --deep --changed-use @world`

Los lectores que agregaron un servidor de binarios [anteriormente](/wiki/Handbook:SPARC/Installation/Base/es#Optional:_Adding_a_binary_package_host "Handbook:SPARC/Installation/Base/es") pueden agregar --getbinpkg (o -g) para obtener paquetes del servidoe de binarios en lugar de compilarlos:

`root #` `emerge --ask --verbose --update --deep --newuse --getbinpkg @world`

#### Eliminar paquetes obsoletos

Es importante hacer siempre _depclean_ después de las actualizaciones del sistema para eliminar paquetes obsoletos. Revise el resultado cuidadosamente con emerge --depclean --pretend para ver si alguno de los paquetes que se van a eliminar debe conservarse si los usa personalmente. Para conservar un paquete que de otro modo se eliminaría, utilice emerge --noreplace foo.

`root #` `emerge --ask --pretend --depclean`

Si está satisfecho, proceda con la limpieza de dependencias real:

`root #` `emerge --ask --depclean`

## Zona horaria

**Nota**

Este paso no se aplica a los usuarios de musl libc. Los usuarios que no sepan lo que eso significa deben realizar este paso.

**Advertencia**

Por favor, evite las zonas horarias listadas en /usr/share/zoneinfo/Etc/GMT\* ya que sus nombres no indican las zonas esperadas. Por ejemplo, GMT-8 es en realidad GMT+8.

Seleccione la zona horaria para su sistema. Busque las zonas horarias disponibles en /usr/share/zoneinfo/:

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

Supongamos que la zona horaria elegida es Europe/Brussels, para seleccionar esta zona horaria, se puede crear un enlace simbólico desde este archivo zoneinfo a /etc/localtime:

`root #` `ln -sf ../usr/share/zoneinfo/Europe/Brussels /etc/localtime`

**Consejo**

La ruta de destino con `../` al inicio es _relativa a la ubicación del enlace_, no al directorio actual.

**Nota**

Se puede utilizar una ruta absoluta para el enlace simbólico, pero systemd también crea un enlace relativo mediante timedatectl y es más compatible con _ROOT_ s alternativos.

## Configurar localizaciones

**Nota**

Este paso no se aplica a los usuarios de musl libc. Los usuarios que no saben lo que eso significa deben realizar este paso.

### Generación de localizaciones

A default installation of Gentoo Linux provides an archive that contains all supported locales, numbering 500 or more. However, it is typical for an administrator to require only one or two of these. In that case, the /etc/locale.gen configuration file may be populated with a list of the required locales. By default, locale-gen shall read this file and compile only the locales that are specified, saving both time and space in the longer term.

Las localizaciones no sólo especifican el idioma que el usuario debe usar para interactuar con sistema, sino también las reglas para ordenar cadenas, presentar fechas y horas, etc.
Los nombres de las localizaciones son _sensibles a mayúsculas_ y deben se escritas exactamente como se ha indicado. Una lista completa de las localizaciones disponibles se puede encontrar en el archivo /usr/share/i18n/SUPPORTED.

`root #` `nano /etc/locale.gen`

Las siguientes localizaciones son un ejemplo de como disponer tanto de Inglés (Estados Unidos) como de Alemán (Alemán/Alemania) junto con los formatos de caracteres (como UTF-8).

ARCHIVO **`/etc/locale.gen`** **Habilitar las localizaciones US y ES con los formatos de carácter apropiados**

```
en_US ISO-8859-1
en_US.UTF-8 UTF-8
es_ES ISO-8859-1
es_ES.UTF-8 UTF-8

```

1. Non UTF-8 locales are discouraged and should only be used in special circumstances.
2. en\_US ISO-8859-1
3. de\_DE ISO-8859-1

}}

**Advertencia**

Muchas aplicaciones requieren al menos una configuración regional UTF-8 para costruirse correctamente.

El siguiente paso es ejecutar la orden locale-gen. Esta ordene genera las localizaciones especificadas en el archivo /etc/locale.gen.

`root #` `locale-gen`

Para verificar que las localizaciones seleccionadas están ahora disponibles, ejecute locale -a.

En instalaciones con systemd, se puede usar localectl, p. ej. localectl set-locale ... o localectl list-locales.

### Selección de localización

Una vez hecho esto, es el momento de establecer la configuración regional de todo el sistema. Nuevamente se usa eselect, ahora con el módulo `locale`.

Con eselect locale list se muestran las opciones disponibles:

`root #` `eselect locale list`

```
Available targets for the LANG variable:
  [1]  C
  [2]  C.utf8
  [3]  en_US
  [4]  en_US.iso88591
  [5]  en_US.utf8
  [6]  de_DE
  [7]  de_DE.iso88591
  [8]  de_DE.utf8
  [9] POSIX
  [ ]  (free form)

```

Con eselect locale set <NÚMERO> se puede seleccionar la localización correcta:

`root #` `eselect locale set 2`

Se puede realizar también manualmente usando el archivo /etc/env.d/02locale y para systemd el archivo /etc/locale.conf:

ARCHIVO **`/etc/env.d/02locale`** **Definición de la localización del sistema manualmente**

```
LANG="de_DE.UTF-8"
LC_COLLATE="C.UTF-8"

```

La configuración de localizaciones previene mensajes de advertencia o error durante la compilación del núcleo y otro software más adelante durante la instalación.

Ahora recargue sus variables de entorno:

`root #` `env-update && source /etc/profile && export PS1="(chroot) ${PS1}"`

Para obtener orientación adicional acerca del proceso de selección de configuración regional, lea también la [Guía de localización](/wiki/Localization/Guide/es "Localization/Guide/es") y la guía [UTF-8](/wiki/UTF-8/es "UTF-8/es").

## Referencias

1. [↑](#cite_ref-1)[https://www.gentoo.org/news/2023/12/29/Gentoo-binary.html](https://www.gentoo.org/news/2023/12/29/Gentoo-binary.html)
2. [↑](#cite_ref-2)[https://www.gentoo.org/glep/glep-0023.html#id7](https://www.gentoo.org/glep/glep-0023.html#id7)

[← Instalar el archivo de Stage](/wiki/Handbook:SPARC/Installation/Stage/es "Previous part") [Inicio](/wiki/Handbook:SPARC/es "Handbook:SPARC/es") [Configurar el núcleo →](/wiki/Handbook:SPARC/Installation/Kernel/es "Next part")