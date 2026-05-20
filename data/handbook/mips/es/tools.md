# Tools

Other languages:

- [Deutsch](/wiki/Handbook:MIPS/Installation/Tools/de "Handbuch:MIPS/Installation/Tools (100% translated)")
- [English](/wiki/Handbook:MIPS/Installation/Tools "Handbook:MIPS/Installation/Tools (100% translated)")
- español
- [français](/wiki/Handbook:MIPS/Installation/Tools/fr "Handbook:MIPS/Installation/Tools/fr (100% translated)")
- [italiano](/wiki/Handbook:MIPS/Installation/Tools/it "Handbook:MIPS/Installation/Tools/it (100% translated)")
- [magyar](/wiki/Handbook:MIPS/Installation/Tools/hu "Handbook:MIPS/Installation/Tools/hu (100% translated)")
- [polski](/wiki/Handbook:MIPS/Installation/Tools/pl "Handbook:MIPS/Installation/Tools (100% translated)")
- [português do Brasil](/wiki/Handbook:MIPS/Installation/Tools/pt-br "Manual:MIPS/Instalação/Ferramentas (100% translated)")
- [русский](/wiki/Handbook:MIPS/Installation/Tools/ru "Handbook:MIPS/Installation/Tools (100% translated)")
- [தமிழ்](/wiki/Handbook:MIPS/Installation/Tools/ta "கையேடு:MIPS/நிறுவல்/கருவிகள் (100% translated)")
- [中文（中国大陆）‎](/wiki/Handbook:MIPS/Installation/Tools/zh-cn "手册：MIPS/安装/安装系统工具 (100% translated)")
- [日本語](/wiki/Handbook:MIPS/Installation/Tools/ja "ハンドブック:MIPS/インストール/ツール (100% translated)")
- [한국어](/wiki/Handbook:MIPS/Installation/Tools/ko "Handbook:MIPS/Installation/Tools/ko (100% translated)")

[Manual MIPS](/wiki/Handbook:MIPS/es "Handbook:MIPS/es")[Instalación](/wiki/Handbook:MIPS/Full/Installation/es "Handbook:MIPS/Full/Installation/es")[Acerca de la instalación](/wiki/Handbook:MIPS/Installation/About/es "Handbook:MIPS/Installation/About/es")[Elegir los medios](/wiki/Handbook:MIPS/Installation/Media/es "Handbook:MIPS/Installation/Media/es")[Configurar la red](/wiki/Handbook:MIPS/Installation/Networking/es "Handbook:MIPS/Installation/Networking/es")[Preparar los discos](/wiki/Handbook:MIPS/Installation/Disks/es "Handbook:MIPS/Installation/Disks/es")[Instalar el stage3](/wiki/Handbook:MIPS/Installation/Stage/es "Handbook:MIPS/Installation/Stage/es")[Instalar el sistema base](/wiki/Handbook:MIPS/Installation/Base/es "Handbook:MIPS/Installation/Base/es")[Configurar el núcleo](/wiki/Handbook:MIPS/Installation/Kernel/es "Handbook:MIPS/Installation/Kernel/es")[Configurar el sistema](/wiki/Handbook:MIPS/Installation/System/es "Handbook:MIPS/Installation/System/es")Instalar las herramientas[Configurar el cargador de arranque](/wiki/Handbook:MIPS/Installation/Bootloader/es "Handbook:MIPS/Installation/Bootloader/es")[Terminar](/wiki/Handbook:MIPS/Installation/Finalizing/es "Handbook:MIPS/Installation/Finalizing/es")[Trabajar con Gentoo](/wiki/Handbook:MIPS/Full/Working/es "Handbook:MIPS/Full/Working/es")[Introducción a Portage](/wiki/Handbook:MIPS/Working/Portage/es "Handbook:MIPS/Working/Portage/es")[Ajustes USE](/wiki/Handbook:MIPS/Working/USE/es "Handbook:MIPS/Working/USE/es")[Características de Portage](/wiki/Handbook:MIPS/Working/Features/es "Handbook:MIPS/Working/Features/es")[Sistema de guiones de inicio](/wiki/Handbook:MIPS/Working/Initscripts/es "Handbook:MIPS/Working/Initscripts/es")[Variables de entorno](/wiki/Handbook:MIPS/Working/EnvVar/es "Handbook:MIPS/Working/EnvVar/es")[Trabajar con Portage](/wiki/Handbook:MIPS/Full/Portage/es "Handbook:MIPS/Full/Portage/es")[Ficheros y directorios](/wiki/Handbook:MIPS/Portage/Files/es "Handbook:MIPS/Portage/Files/es")[Variables](/wiki/Handbook:MIPS/Portage/Variables/es "Handbook:MIPS/Portage/Variables/es")[Mezclar ramas de software](/wiki/Handbook:MIPS/Portage/Branches/es "Handbook:MIPS/Portage/Branches/es")[Herramientas adicionales](/wiki/Handbook:MIPS/Portage/Tools/es "Handbook:MIPS/Portage/Tools/es")[Repositorios personalizados de paquetes](/wiki/Handbook:MIPS/Portage/CustomTree/es "Handbook:MIPS/Portage/CustomTree/es")[Características avanzadas](/wiki/Handbook:MIPS/Portage/Advanced/es "Handbook:MIPS/Portage/Advanced/es")[Configuración de la red](/wiki/Handbook:MIPS/Full/Networking/es "Handbook:MIPS/Full/Networking/es")[Comenzar](/wiki/Handbook:MIPS/Networking/Introduction/es "Handbook:MIPS/Networking/Introduction/es")[Configuración avanzada](/wiki/Handbook:MIPS/Networking/Advanced/es "Handbook:MIPS/Networking/Advanced/es")[Configuración de red modular](/wiki/Handbook:MIPS/Networking/Modular/es "Handbook:MIPS/Networking/Modular/es")[Conexión inalámbrica](/wiki/Handbook:MIPS/Networking/Wireless/es "Handbook:MIPS/Networking/Wireless/es")[Añadir funcionalidad](/wiki/Handbook:MIPS/Networking/Extending/es "Handbook:MIPS/Networking/Extending/es")[Gestión dinámica](/wiki/Handbook:MIPS/Networking/Dynamic/es "Handbook:MIPS/Networking/Dynamic/es")

## Contents

- [1Bitácora del sistema](#Bit.C3.A1cora_del_sistema)
  - [1.1OpenRC](#OpenRC)
  - [1.2systemd](#systemd)
- [2Opcional: Demonio Cron](#Opcional:_Demonio_Cron)
  - [2.1OpenRC](#OpenRC_2)
    - [2.1.1Por defecto: cronie](#Por_defecto:_cronie)
    - [2.1.2Alternativa: dcron](#Alternativa:_dcron)
    - [2.1.3Alternativa: fcron](#Alternativa:_fcron)
    - [2.1.4Alternativa: bcron](#Alternativa:_bcron)
    - [2.1.5modprobed-db users](#modprobed-db_users)
  - [2.2systemd](#systemd_2)
    - [2.2.1modprobed-db users](#modprobed-db_users_2)
- [3Opcional: Indexar Archivos](#Opcional:_Indexar_Archivos)
- [4Opcional: Shell de acceso remoto](#Opcional:_Shell_de_acceso_remoto)
  - [4.1OpenRC](#OpenRC_3)
  - [4.2systemd](#systemd_3)
- [5Opcional: Completado de shell](#Opcional:_Completado_de_shell)
  - [5.1Bash](#Bash)
- [6Sugerencia: Sincronización de la hora](#Sugerencia:_Sincronizaci.C3.B3n_de_la_hora)
  - [6.1chrony](#chrony)
  - [6.2OpenRC](#OpenRC_4)
  - [6.3systemd](#systemd_4)
  - [6.4systemd-timesyncd](#systemd-timesyncd)
- [7Herramientas del Sistema de Archivos](#Herramientas_del_Sistema_de_Archivos)

## Bitácora del sistema

### OpenRC

Algunas herramientas no están incluidas en el archivo stage3 porque varios paquetes proporcionan la misma funcionalidad. Ahora es el momento de que el usuario decida cual instalar.

La primera herramienta sobre la que tomar una decisión es el mecanismo de registro para el sistema. Unix y Linux tienen un excelente historial de capacidades de registro; si fuera necesario, todo lo que sucede en el sistema se podría registrar en un archivo de registro.

Gentoo ofrece varias utilidades de registro. Algunas de ellas son:

- [sysklogd](/wiki/Sysklogd "Sysklogd") ([app-admin/sysklogd](https://packages.gentoo.org/packages/app-admin/sysklogd)) \- Ofrece el conjunto tradicional de procesos de registro del sistema. La configuración de registro predeterminada funciona adecuadamente, lo que convierte a este paquete en una buena opción para principiantes.
- [syslog-ng](/wiki/Syslog-ng "Syslog-ng") ([app-admin/syslog-ng](https://packages.gentoo.org/packages/app-admin/syslog-ng)) \- Un registrador del sistema avanzado. Requiere configuración adicional para cualquier función que vaya más allá del registro en un archivo grande. Los usuarios más avanzados pueden elegir este paquete por su potencial de registro; tenga en cuenta que se requiere configuración adicional para cualquier tipo de gestión de los registros.
- [metalog](/wiki/Metalog "Metalog") ([app-admin/metalog](https://packages.gentoo.org/packages/app-admin/metalog)) \- Un registrador del sistema altamente configurable.

Es posible que haya otras utilidades de registro del sistema disponibles a través del repositorio ebuild de Gentoo, ya que el número de paquetes disponibles aumenta diariamente.

**Consejo**

Si se va a utilizar syslog-ng, se recomienda instalar y configurar a continuación [logrotate](/wiki/Logrotate "Logrotate"), syslog-ng no proporciona ningún mecanismo de rotación para los archivos de registro. Sin embargo las versiones más nuevas (>= 2.0) de sysklogd manejan su propia rotación de registros.

Para instalar la bitácora del sistema de su elección, instálela con emerge. En OpenRC agréguela al nivel de arranque predeterminado usando rc-update. El siguiente ejemplo instala y activa [app-admin/sysklogd](https://packages.gentoo.org/packages/app-admin/sysklogd) como la utilidad syslog del sistema:

`root #` `emerge --ask app-admin/sysklogd`

`root #` `rc-update add sysklogd default`

### systemd

Si bien se presenta una selección de mecanismos de registro para sistemas basados en OpenRC, systemd incluye un registrador integrado llamado servicio **systemd-journald**. El servicio systemd-journald es capaz de manejar la mayor parte de la funcionalidad de registro descrita en la sección anterior del registro del sistema. Es decir, la mayoría de las instalaciones que ejecutarán systemd como administrador del sistema y de servicios pueden "omitir de forma segura" la adición de utilidades syslog adicionales.

Consulte man journalctl para obtener más detalles sobre el uso de journalctl para consultar y revisar los registros del sistema.

Por varias razones, como sería el caso de reenviar registros a un servidoe central, puede ser importante incluir mecanismos de registro del sistema "redundantes" en un sistema basado en systemd. Esto es algo inusual para la audiencia típica del manual y se considera un caso de uso avanzado. Por lo tanto, no está cubierto por el manual.

## Opcional: Demonio Cron

### OpenRC

Aunque es opcional y no es necesario para todos los sistemas, es aconsejable instalar un demonio cron.

Un demonio cron ejecuta comandos en intervalos programados. Los intervalos pueden ser diarios, semanales o mensuales, una vez cada martes, una vez cada dos semanas, etc. Un administrador de sistema inteligente aprovechará el demonio cron para automatizar las tareas rutinarias de mantenimiento del sistema.

Todos los demonios cron admiten altos niveles de granularidad para las tareas programadas y, por lo general, incluyen la capacidad de enviar un correo electrónico u otra forma de notificación si una tarea programada no se completa como se esperaba.

Gentoo ofrece varios demonios cron posibles, incluyendo:

- [cronie](/wiki/Cron/es#cronie "Cron/es") ([sys-process/cronie](https://packages.gentoo.org/packages/sys-process/cronie)) \- cronie se basa en el cron original y tiene mejoras de seguridad y configuración como la capacidad de usar PAM y SELinux.
- [dcron](/wiki/Cron/es#dcron "Cron/es") ([sys-process/dcron](https://packages.gentoo.org/packages/sys-process/dcron)) \- Este demonio cron ligero pretende ser simple y seguro, con las funciones suficientes para seguir siendo útil.
- [fcron](/wiki/Cron/es#fcron "Cron/es") ([sys-process/fcron](https://packages.gentoo.org/packages/sys-process/fcron)) \- Un programador de comandos con capacidades ampliadas respecto a cron y anacron.
- [bcron](/wiki/Cron/es#bcron "Cron/es") ([sys-process/bcron](https://packages.gentoo.org/packages/sys-process/bcron)) \- Un sistema cron más reciente diseñado teniendo en cuenta la seguridad en las operaciones. Para hacer esto, el sistema se divide en varios programas separados, cada uno responsable de una tarea separada, con comunicaciones estrictamente controladas entre las partes.

#### Por defecto: cronie

El siguiente ejemplo usa [sys-process/cronie](https://packages.gentoo.org/packages/sys-process/cronie):

`root #` `emerge --ask sys-process/cronie`

Agregue cronie al nivel de ejecución predeterminado del sistema, que lo arrancará automáticamente al inicio:

`root #` `rc-update add cronie default`

#### Alternativa: dcron

`root #` `emerge --ask sys-process/dcron`

Si dcron es el agente cron con el que se proseguirá, se debe ejecutar un comando de inicialización adicional:

`root #` `crontab /etc/crontab`

#### Alternativa: fcron

`root #` `emerge --ask sys-process/fcron`

Si fcron es el controlador de tareas programadas seleccionado, se requiere un paso adicional tras emerge:

`root #` `emerge --config sys-process/fcron`

#### Alternativa: bcron

bcron es un agente cron más reciente con separación de privilegios incorporada.

`root #` `emerge --ask sys-process/bcron`

#### modprobed-db users

If using [sys-kernel/modprobed-db](https://packages.gentoo.org/packages/sys-kernel/modprobed-db) was picked as an option for manually compiling the kernel.

Now is the time to setup a crontab to periodically scan the system for hardware used.

ARCHIVO **`/etc/crontab`** **Run modprobed-db once every 6 hours**

```
0 */6 * * *     /usr/bin/modprobed-db store &> /dev/null

```

At a later date of at least a week, please visit the kernel build section of the [modprobed-db](/wiki/Modprobed-db#Building_kernels.2Fes "Modprobed-db") article to complete the setup.

### systemd

De manera similar al registro del sistema, los sistemas basados en systemd incluyen soporte para tareas programadas listas para usar en forma de "temporizadores". Los temporizadores systemd pueden ejecutarse a nivel de sistema o de usuario e incluir la misma funcionalidad que proporcionaría un demonio cron tradicional. A menos que sean necesarias capacidades redundantes, la instalación de un programador de tareas adicional, como un demonio cron, generalmente es innecesaria y se puede omitir de manera segura.

#### modprobed-db users

If using [sys-kernel/modprobed-db](https://packages.gentoo.org/packages/sys-kernel/modprobed-db) was picked as an option for manually compiling the kernel.

Now is the time to setup a systemd timer to periodically scan the system for hardware used.

`root #` `systemctl --user enable modprobed-db`

At a later date of at least a week, please visit the kernel build section of the [modprobed-db](/wiki/Modprobed-db#Building_kernels.2Fes "Modprobed-db") article to complete the setup.

## Opcional: Indexar Archivos

Para disponer de un índice en su sistema de archivos que proporcionará capacidades para la rápida localización de archivos, instale [mlocate](/wiki/Mlocate "Mlocate").

`root #` `emerge --ask sys-apps/mlocate`

## Opcional: Shell de acceso remoto

**Consejo**

La configuración predeterminada de opensshd no permite que root inicie sesión como usuario remoto. [Cree un usuario no root](/wiki/FAQ/es#How_do_I_add_a_normal_user.3F "FAQ/es") y configúrelo adecuadamente para permitir el acceso posterior a la instalación si es necesario, o ajuste /etc/ssh/sshd\_config para permitir root.

Para poder acceder al sistema de forma remota después de la instalación, sshd debe estar configurado para iniciarse en el arranque.

Para obtener mas detalles sobre la configuración de SSH, consulte el artículo [SSH/es](/wiki/SSH/es "SSH/es").

### OpenRC

Para añadir el script de inicio sshd al nivel de ejecución por defecto en OpenRC:

`root #` `rc-update add sshd default`

`root #` `rc-update add sshd default`

Si se necesita acceso a la consola serie (que es posible en el caso de servidores remotos), se debe configurar agetty.

Descomente la sección de la consola serie en /etc/inittab:

`root #` `nano -w /etc/inittab`

```
# SERIAL CONSOLES
s0:12345:respawn:/sbin/agetty 9600 ttyS0 vt100
s1:12345:respawn:/sbin/agetty 9600 ttyS1 vt100

```

### systemd

Para habilitar el servidor SSH, ejecute:

`root #` `systemctl enable sshd`

`root #` `systemctl enable sshd`

Para habilitar la compatibilidad con la consola serie, ejecute:

`root #` `systemctl enable getty@tty1.service`

`root #` `systemctl enable getty@tty1.service`

## Opcional: Completado de shell

### Bash

Bash es el shell predeterminado para los sistemas Gentoo y, por lo tanto, la instalación de extensiones de completado puede ayudar a mejorar la eficiencia y la conveniencia para administrar el sistema. El paquete [app-shells/bash-completion](https://packages.gentoo.org/packages/app-shells/bash-completion) instalará completaciones disponibles para comandos específicos de Gentoo, así como muchos otros comandos y utilidades comunes:

`root #` `emerge --ask app-shells/bash-completion`

Después de la instalación, el completado de bash para comandos específicos se puede gestionar a través de eselect. Consulte la [sección Integraciones de completados de Shell](/wiki/Bash#Shell_completion_integrations.2Fes "Bash") del artículo de bash para obtener más detalles.

## Sugerencia: Sincronización de la hora

**Importante**

Systems without a functioning [Real-Time Clock (RTC)](/wiki/System_time/es#Software_clock_vs_Hardware_clock "System time/es") must set the [system time](/wiki/System_time/es "System time/es") at every system start, and on regular intervals thereafter.

Es importante utilizar algún método de sincronización para el reloj del sistema. Esto generalmente se hace a través del protocolo y el software [NTP](/wiki/NTP "NTP"). Existen otras implementaciones que utilizan el protocolo NTP, como [Chrony](/wiki/Chrony "Chrony").

The clock is usually synchronized via the [Network Time Protocol](/wiki/Network_Time_Protocol "Network Time Protocol"), with an implementation such as [chrony](/wiki/Chrony "Chrony").

### chrony

Para usar Chrony, por ejemplo:

`root #` `emerge --ask net-misc/chrony`

`root #` `emerge --ask net-misc/chrony`

See the [chrony](/wiki/Chrony "Chrony") article for further information, for example if more advanced configurations are required.

### OpenRC

En OpenRC, ejecute:

`root #` `rc-update add chronyd default`

`root #` `rc-update add chronyd default`

### systemd

En systemd, ejecute:

`root #` `systemctl enable chronyd.service`

`root #` `systemctl enable chronyd.service`

### systemd-timesyncd

Alternativamente, los usuarios de systemd pueden desear utilizar el sencillo cliente SNTP systemd-timesyncd que se instala de forma predeterminada.

`root #` `systemctl enable systemd-timesyncd.service`

`root #` `systemctl enable systemd-timesyncd.service`

## Herramientas del Sistema de Archivos

Dependiendo de los sistemas de archivos utilizados, puede ser necesario instalar las utilidades del sistema de archivos requeridas (para verificar la integridad del sistema de archivos, (re)formatear sistemas de archivos adicionales, etc.). Tenga en cuenta que las herramientas del espacio de usuario para ext4 ([sys-fs/e2fsprogs](https://packages.gentoo.org/packages/sys-fs/e2fsprogs)) ya están instaladas como parte del [conjunto @system](/wiki/System_set_(Portage) "System set (Portage)").

La siguiente tabla enumera las herramientas a instalar si se necesitan determinadas herramientas del sistema de archivos en el entorno instalado.

Sistema de ficheros
Paquete
[XFS](/wiki/XFS "XFS")[sys-fs/xfsprogs](https://packages.gentoo.org/packages/sys-fs/xfsprogs)[ext4](/wiki/Ext4 "Ext4")[sys-fs/e2fsprogs](https://packages.gentoo.org/packages/sys-fs/e2fsprogs)[VFAT](/wiki/VFAT "VFAT") (FAT32, ...)
[sys-fs/dosfstools](https://packages.gentoo.org/packages/sys-fs/dosfstools)[Btrfs](/wiki/Btrfs/es "Btrfs/es")[sys-fs/btrfs-progs](https://packages.gentoo.org/packages/sys-fs/btrfs-progs)[F2FS](/wiki/F2FS "F2FS")[sys-fs/f2fs-tools](https://packages.gentoo.org/packages/sys-fs/f2fs-tools)[NTFS](/wiki/NTFS "NTFS")[sys-fs/ntfs3g](https://packages.gentoo.org/packages/sys-fs/ntfs3g)[ZFS](/wiki/ZFS "ZFS")[sys-fs/zfs](https://packages.gentoo.org/packages/sys-fs/zfs)

Se recomienda que [sys-block/io-scheduler-udev-rules](https://packages.gentoo.org/packages/sys-block/io-scheduler-udev-rules) esté instalado para el correcto comportamiento del planificador con, por ejemplo, dispositivos nvme:

`root #` `emerge --ask sys-block/io-scheduler-udev-rules`

`root #` `emerge --ask sys-block/io-scheduler-udev-rules`

**Consejo**

Para obtener más información acerca de los sistemas de archivo en Gentoo, echar un vistazo al [articulo sistemas de archivos](/wiki/Filesystem/es "Filesystem/es")

Ahora continúe con [Configurar el cargador de arranque](/wiki/Handbook:MIPS/Installation/Bootloader/es "Handbook:MIPS/Installation/Bootloader/es").

[← Configurar el sistema](/wiki/Handbook:MIPS/Installation/System/es "Previous part") [Inicio](/wiki/Handbook:MIPS/es "Handbook:MIPS/es") [Configurar el cargador de arranque →](/wiki/Handbook:MIPS/Installation/Bootloader/es "Next part")