# Finalizing

Other languages:

- [Deutsch](/wiki/Handbook:PPC/Installation/Finalizing/de "Handbuch:PPC/Installation/Abschluss (100% translated)")
- [English](/wiki/Handbook:PPC/Installation/Finalizing "Handbook:PPC/Installation/Finalizing (100% translated)")
- español
- [français](/wiki/Handbook:PPC/Installation/Finalizing/fr "Handbook:PPC/Installation/Finalizing/fr (100% translated)")
- [italiano](/wiki/Handbook:PPC/Installation/Finalizing/it "Handbook:PPC/Installation/Finalizing/it (100% translated)")
- [magyar](/wiki/Handbook:PPC/Installation/Finalizing/hu "Handbook:PPC/Installation/Finalizing/hu (100% translated)")
- [polski](/wiki/Handbook:PPC/Installation/Finalizing/pl "Handbook:PPC/Installation/Finalizing (100% translated)")
- [português do Brasil](/wiki/Handbook:PPC/Installation/Finalizing/pt-br "Handbook:PPC/Installation/Finalizing/pt-br (50% translated)")
- [русский](/wiki/Handbook:PPC/Installation/Finalizing/ru "Handbook:PPC/Installation/Finalizing (100% translated)")
- [தமிழ்](/wiki/Handbook:PPC/Installation/Finalizing/ta "கையேடு:PPC/நிறுவல்/முடித்தல் (100% translated)")
- [中文（中国大陆）‎](/wiki/Handbook:PPC/Installation/Finalizing/zh-cn "手册：PPC/安装/收尾安装工作 (100% translated)")
- [日本語](/wiki/Handbook:PPC/Installation/Finalizing/ja "ハンドブック:PPC/インストール/ファイナライズ (100% translated)")
- [한국어](/wiki/Handbook:PPC/Installation/Finalizing/ko "Handbook:PPC/Installation/Finalizing/ko (100% translated)")

[PPC Handbook](/wiki/Handbook:PPC "Handbook:PPC")[Installation](/wiki/Handbook:PPC/Full/Installation "Handbook:PPC/Full/Installation")[About the installation](/wiki/Handbook:PPC/Installation/About/es "Handbook:PPC/Installation/About/es")[Choosing the media](/wiki/Handbook:PPC/Installation/Media/es "Handbook:PPC/Installation/Media/es")[Configuring the network](/wiki/Handbook:PPC/Installation/Networking/es "Handbook:PPC/Installation/Networking/es")[Preparing the disks](/wiki/Handbook:PPC/Installation/Disks/es "Handbook:PPC/Installation/Disks/es")[The stage file](/wiki/Handbook:PPC/Installation/Stage/es "Handbook:PPC/Installation/Stage/es")[Installing base system](/wiki/Handbook:PPC/Installation/Base/es "Handbook:PPC/Installation/Base/es")[Configuring the kernel](/wiki/Handbook:PPC/Installation/Kernel/es "Handbook:PPC/Installation/Kernel/es")[Configuring the system](/wiki/Handbook:PPC/Installation/System/es "Handbook:PPC/Installation/System/es")[Installing tools](/wiki/Handbook:PPC/Installation/Tools/es "Handbook:PPC/Installation/Tools/es")[Configuring the bootloader](/wiki/Handbook:PPC/Installation/Bootloader/es "Handbook:PPC/Installation/Bootloader/es")Finalizing[Working with Gentoo](/wiki/Handbook:PPC/Full/Working "Handbook:PPC/Full/Working")[Portage introduction](/wiki/Handbook:PPC/Working/Portage/es "Handbook:PPC/Working/Portage/es")[USE flags](/wiki/Handbook:PPC/Working/USE/es "Handbook:PPC/Working/USE/es")[Portage features](/wiki/Handbook:PPC/Working/Features/es "Handbook:PPC/Working/Features/es")[Initscript system](/wiki/Handbook:PPC/Working/Initscripts/es "Handbook:PPC/Working/Initscripts/es")[Environment variables](/wiki/Handbook:PPC/Working/EnvVar/es "Handbook:PPC/Working/EnvVar/es")[Working with Portage](/wiki/Handbook:PPC/Full/Portage "Handbook:PPC/Full/Portage")[Files and directories](/wiki/Handbook:PPC/Portage/Files/es "Handbook:PPC/Portage/Files/es")[Variables](/wiki/Handbook:PPC/Portage/Variables/es "Handbook:PPC/Portage/Variables/es")[Mixing software branches](/wiki/Handbook:PPC/Portage/Branches/es "Handbook:PPC/Portage/Branches/es")[Additional tools](/wiki/Handbook:PPC/Portage/Tools/es "Handbook:PPC/Portage/Tools/es")[Custom package repository](/wiki/Handbook:PPC/Portage/CustomTree/es "Handbook:PPC/Portage/CustomTree/es")[Advanced features](/wiki/Handbook:PPC/Portage/Advanced/es "Handbook:PPC/Portage/Advanced/es")[OpenRC network configuration](/wiki/Handbook:PPC/Full/Networking "Handbook:PPC/Full/Networking")[Getting started](/wiki/Handbook:PPC/Networking/Introduction/es "Handbook:PPC/Networking/Introduction/es")[Advanced configuration](/wiki/Handbook:PPC/Networking/Advanced/es "Handbook:PPC/Networking/Advanced/es")[Modular networking](/wiki/Handbook:PPC/Networking/Modular/es "Handbook:PPC/Networking/Modular/es")[Wireless](/wiki/Handbook:PPC/Networking/Wireless/es "Handbook:PPC/Networking/Wireless/es")[Adding functionality](/wiki/Handbook:PPC/Networking/Extending/es "Handbook:PPC/Networking/Extending/es")[Dynamic management](/wiki/Handbook:PPC/Networking/Dynamic/es "Handbook:PPC/Networking/Dynamic/es")

## Contents

- [1Administración del usuario](#Administraci.C3.B3n_del_usuario)
  - [1.1Añadir un usuario para uso cotidiano](#A.C3.B1adir_un_usuario_para_uso_cotidiano)
  - [1.2Elevación de privilegios temporal](#Elevaci.C3.B3n_de_privilegios_temporal)
  - [1.3Deshabilitar el login de root](#Deshabilitar_el_login_de_root)
- [2Limpieza de disco](#Limpieza_de_disco)
  - [2.1Eliminación de archivos de instalación](#Eliminaci.C3.B3n_de_archivos_de_instalaci.C3.B3n)
- [3¿Adónde ir desde aquí?](#.C2.BFAd.C3.B3nde_ir_desde_aqu.C3.AD.3F)
  - [3.1Documentación adicional](#Documentaci.C3.B3n_adicional)
  - [3.2Gentoo en línea](#Gentoo_en_l.C3.ADnea)
    - [3.2.1Foros e IRC](#Foros_e_IRC)
    - [3.2.2Listas de correo](#Listas_de_correo)
    - [3.2.3Incidencias](#Incidencias)
    - [3.2.4Guía de desarrollo](#Gu.C3.ADa_de_desarrollo)
  - [3.3Consideraciones finales](#Consideraciones_finales)

## Administración del usuario

### Añadir un usuario para uso cotidiano

Trabajar como root en un sistema Unix/Linux es peligroso y debe evitarse tanto como sea posible. Por tanto se recomienda encarecidamente añadir una o mas cuentas de usuario normal para el uso cotidiano del sistema.

Los grupos a los que pertenece el usuario definen que actividades puede realizar. La siguiente tabla muestra una lista de los grupos más importantes:

Grupo
Descripción
audio
Habilita cuentas de usuarios para acceder a los dispositivos de audio.
cdrom
Habilita cuentas de usuarios para que puedan acceder de forma directa a dispositivos de lectura óptica.
floppy
Habilita cuentas de usuarios para acceder directamente a dispositivos mecánicos antiguos conocidos como unidades de disquete. Este grupo generalmente no se utiliza en sistemas actuales.
usb
Habilita cuentas de usuarios para acceder a los dispositivos USB.
video
Habilita cuentas de usuarios para acceder al hardware de captura de vídeo y a la aceleración por hardware.
wheel
Permite que las cuentas de usuaris puedan utilizar el comando su ("usuario sustituto"), que permite cambiar a la cuenta root u otras cuentas. Para sistemas de usuario único que incluyen una cuenta root, es una buena idea agregar a este grupo el usuario normal principal.

Por ejemplo, para crear un usuario llamado [larry](/wiki/User:Larry "User:Larry") que pertenezca a los grupos _wheel_, _users_ y _audio_, acceda al sistema como root (solo root puede crear usuarios) y ejecute useradd:

`Login:` `root`

```
Password: (Introduzca la contraseña de root)

```

Al configurar contraseñas para cuentas de usuario normal, es una buena práctica de seguridad evitar el uso de la misma contraseña o una similar a la configurada para el usuario root.

Los autores del manual recomendaron utilizar una contraseña de al menos 16 caracteres de longitud, con un contenido exclusivo y completamente diferente a la del resto de usuarios del sistema.

`root #` `useradd -m -G users,wheel,audio -s /bin/bash larry
`

`root #` `passwd larry`

```
Password: (Introduzca una contraseña para larry)
Re-enter password: (Repita la contraseña como comprobación)

```

### Elevación de privilegios temporal

Si alguna vez este usuario necesita realizar alguna tarea como root, puede utilizar su - para obtener temporalmente privilegios de root. Otra forma es utilizar el paquete [sudo](/wiki/Sudo/es "Sudo/es") ([app-admin/sudo](https://packages.gentoo.org/packages/app-admin/sudo)) o la utilidad [doas](/wiki/Doas "Doas") ([app-admin/doas](https://packages.gentoo.org/packages/app-admin/doas)) los cuales, correctamente configurados, son muy seguros.

### Deshabilitar el login de root

**Advertencia**

Antes de deshabilitar el inicio de sesión para root, asegúrese de que haya una cuenta de usuario que pertenezca al grupo wheel y de que exista un método para elevar los privilegios de usuario; de lo contrario, se bloqueará el acceso root y será imposible administrar el sistema sin realizar una recuperación. Algunos métodos comunes para elevar los privilegios de usuario incluyen: [app-admin/sudo](https://packages.gentoo.org/packages/app-admin/sudo), [app-admin/doas](https://packages.gentoo.org/packages/app-admin/doas) o run0 de systemd.

Para evitar que posibles acciones de amenazas inicien sesión como root, eliminar la contraseña de root y/o deshabilitar el inicio de sesión de root pueden ayudar a mejorar la seguridad.

Para deshabilitar el login de root:

`root #` `passwd -l root`

Para eliminar la contraseña de root y deshabilitar su login:

`root #` `passwd -dl root`

## Limpieza de disco

### Eliminación de archivos de instalación

Una vez finalizada la instalación de Gentoo y reiniciado el sistema, si todo ha ido bien, el archivo de stage y otros archivos de instalación - como archivos DIGEST, CONTENT o \*.asc (firma PGP) - ahora se pueden eliminar de forma segura.

Los archivos se encuentran en el directorio / y se pueden eliminar con el siguiente comando:

`root #` `rm /stage3-*.tar.*`

## ¿Adónde ir desde aquí?

¿No está seguro de dónde ir desde aquí?. Hay muchos caminos a explorar... Gentoo proporciona a sus usuarios montones de posibilidades, y por lo tanto tiene montones de artículos documentados (y menos también menos documentados) para que sean explorados aquí en el wiki y en otros subdominios relacionados (leer la sección [Gentoo online](/wiki/Handbook:PPC/Installation/Finalizing/es#Gentoo_online "Handbook:PPC/Installation/Finalizing/es") abajo).

### Documentación adicional

Es importante señalar que, debido a la cantidad de opciones disponibles en Gentoo, la documentación proporcionada por el manual tiene un alcance limitado: se enfoca principalmente en los conceptos básicos para poner en marcha un sistema Gentoo y las actividades básicas de gestión del sistema. El manual excluye intencionalmente instrucciones sobre entornos gráficos, detalles sobre securización y otras tareas administrativas importantes. Dicho esto, hay más secciones del manual para ayudar a los lectores con funciones más básicas.

Definitivamente los lectores deben echar un vistazo a la siguiente parte del manual de Gentoo titulado [Trabajando con Gentoo](/wiki/Handbook:PPC/Working/Portage/es "Handbook:PPC/Working/Portage/es") que explica como mantener el software actualizado, instalar paquetes adicionales de software, detalles sobre los ajustes USE, el sistema de inicio OpenRC y algunos otros elementos informativos relacionados con la gestión de un sistema Gentoo después de su instalación.

Además del manual, el lector debería animarse a explorar otros rincones del sitio wiki de Gentoo para encontrar documentación adicional proporcionada por la comunidad. El equipo wiki de Gentoo también ofrece una [descripción general del tema de la documentación](/wiki/Main_Page#Documentation_topics "Main Page") que ofrece una selección de artículos del wiki por categoría. Por ejemplo, se hace referencia a la [guía sobre la localización](/wiki/Localization/Guide/es "Localization/Guide/es") para hacer el sistema mas parecido a su país (particularmente útil para los usuarios cuyo segundo idioma es el inglés).

La mayoría de los usuarios que usan escritorios configurarán entornos gráficos en los que trabajar de forma nativa. Hay muchos 'meta' artículos mantenidos por la comunidad para [entornos de escritorio (DEs)](/wiki/Desktop_environment "Desktop environment") y [administradores de ventanas (WMs)](/wiki/Window_manager "Window manager"). Los lectores deben tener en cuenta que cada DE requerirá pasos de configuración ligeramente diferentes, lo que alargará y agregará complejidad al arranque.

Existen muchos otros [Meta-artículos](/wiki/Category:Meta "Category:Meta") para proporcionar a nuestros lectores una visión general de alto nivel del software disponible dentro de Gentoo.

### Gentoo en línea

**Importante**

Los lectores deben tener en cuenta que todos los sitios oficiales de Gentoo en línea están dirigidos por el [código de conducta](/wiki/Project:Council/Code_of_conduct "Project:Council/Code of conduct") de Gentoo. Ser un miembro activo de la comunidad Gentoo es un privilegio no un derecho y los usuarios deben tener en cuenta que el código de conducta está ahí por alguna razón.

Con la excepción de la red Intenet Relay Chat (IRC) alojado en Libera.Chat y las listas de correo, la mayoría de los sitios web de Gentoo requieren del uso de una cuenta en cada sitio para realizar preguntas, abrir una discusión o informar de un error.

#### Foros e IRC

Todos los usuarios son siempre bienvenidos a nuestros [Foros de Gentoo](https://forums.gentoo.org) o a alguno de nuestros [canales de internet relay chat](https://www.gentoo.org/get-involved/irc-channels/). Es fácil buscar en los foros para comprobar si se ha descubierto algún problema en una instalación nueva de Gentoo en el pasado y si ha sido resuelta después de ofrecer algunos comentarios. La cantidad de usuarios que experimentan problemas cuando instalan Gentoo por primera vez es sorprendente. Se recomienda a los usuarios que busquen en los foros y en el wiki antes de pedir ayuda en los canales de soporte de Gentoo.

#### Listas de correo

Se dispone de [varias listas de correo](https://www.gentoo.org/get-involved/mailing-lists/) para los miembros de la comunidad que prefieran pedir ayuda o consejos a través del correo electrónico en lugar de crear una cuenta de usuario en los foros o en IRC. Los usuarios deben seguir las instrucciones para suscribirse a las listas de correo que deseen.

#### Incidencias

A veces, después de revisar la wiki, buscar en los foros y buscar apoyo en el canal de IRC o en las listas de correo, no se encuentra una solución al problema. Generalmente, esta es una señal para abrir un error en el [sitio web Bugzilla](https://bugs.gentoo.org) de Gentoo.

#### Guía de desarrollo

Los lectores que deseen aprender más sobre el desarrollo de Gentoo pueden consultar la [Guía de desarrollo](https://devmanual.gentoo.org/). Esta guía proporciona instrucciones sobre cómo escribir ebuilds, trabajar con eclasses y proporciona definiciones para muchos [conceptos generales](https://devmanual.gentoo.org/general-concepts/index.html) detrás del desarrollo de Gentoo.

### Consideraciones finales

Gentoo es una distribución robusta, flexible y excelentemente mantenida. A la comunidad de mantenedores le encantaría escucha su opinión acerca de como hacer de Gentoo una distribución incluso _mejor_.

Como recordatorio, cualquier comentario sobre _este manual_ debe seguir las pautas detalladas en la sección [¿Cómo puedo mejorar el Manual?](/wiki/Handbook:Main_Page/es#How_do_I_improve_the_Handbook.3F "Handbook:Main Page/es") al principio del manual.

Esperamos que nuestros usuarios elijen implementar Gentoo para cubrir sus necesidades y caso únicos.

[← Configurar el cargador de arranque](/wiki/Handbook:PPC/Installation/Bootloader/es "Previous part") [Inicio](/wiki/Handbook:PPC/es "Handbook:PPC/es") [→](/wiki/Handbook:PPC/Working/Portage/es "Next part")