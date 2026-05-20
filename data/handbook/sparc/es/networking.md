# Networking

Other languages:

- [Deutsch](/wiki/Handbook:SPARC/Installation/Networking/de "Handbuch:SPARC/Installation/Netzwerk (100% translated)")
- [English](/wiki/Handbook:SPARC/Installation/Networking "Handbook:SPARC/Installation/Networking (100% translated)")
- español
- [français](/wiki/Handbook:SPARC/Installation/Networking/fr "Handbook:SPARC/Installation/Networking/fr (100% translated)")
- [italiano](/wiki/Handbook:SPARC/Installation/Networking/it "Handbook:SPARC/Installation/Networking/it (50% translated)")
- [magyar](/wiki/Handbook:SPARC/Installation/Networking/hu "Handbook:SPARC/Installation/Networking/hu (100% translated)")
- [polski](/wiki/Handbook:SPARC/Installation/Networking/pl "Handbook:SPARC/Installation/Networking (100% translated)")
- [português do Brasil](/wiki/Handbook:SPARC/Installation/Networking/pt-br "Handbook:SPARC/Installation/Networking/pt-br (50% translated)")
- [русский](/wiki/Handbook:SPARC/Installation/Networking/ru "Handbook:SPARC/Installation/Networking (100% translated)")
- [தமிழ்](/wiki/Handbook:SPARC/Installation/Networking/ta "கையேடு:SPARC/நிறுவல்/வலையமைத்தல் (100% translated)")
- [中文（中国大陆）‎](/wiki/Handbook:SPARC/Installation/Networking/zh-cn "手册：SPARC/安装/配置网络 (100% translated)")
- [日本語](/wiki/Handbook:SPARC/Installation/Networking/ja "ハンドブック:SPARC/インストール/ネットワーク (100% translated)")
- [한국어](/wiki/Handbook:SPARC/Installation/Networking/ko "Handbook:SPARC/Installation/Networking/ko (100% translated)")

[SPARC Handbook](/wiki/Handbook:SPARC "Handbook:SPARC")[Installation](/wiki/Handbook:SPARC/Full/Installation "Handbook:SPARC/Full/Installation")[About the installation](/wiki/Handbook:SPARC/Installation/About/es "Handbook:SPARC/Installation/About/es")[Choosing the media](/wiki/Handbook:SPARC/Installation/Media/es "Handbook:SPARC/Installation/Media/es")Configuring the network[Preparing the disks](/wiki/Handbook:SPARC/Installation/Disks/es "Handbook:SPARC/Installation/Disks/es")[The stage file](/wiki/Handbook:SPARC/Installation/Stage/es "Handbook:SPARC/Installation/Stage/es")[Installing base system](/wiki/Handbook:SPARC/Installation/Base/es "Handbook:SPARC/Installation/Base/es")[Configuring the kernel](/wiki/Handbook:SPARC/Installation/Kernel/es "Handbook:SPARC/Installation/Kernel/es")[Configuring the system](/wiki/Handbook:SPARC/Installation/System/es "Handbook:SPARC/Installation/System/es")[Installing tools](/wiki/Handbook:SPARC/Installation/Tools/es "Handbook:SPARC/Installation/Tools/es")[Configuring the bootloader](/wiki/Handbook:SPARC/Installation/Bootloader/es "Handbook:SPARC/Installation/Bootloader/es")[Finalizing](/wiki/Handbook:SPARC/Installation/Finalizing/es "Handbook:SPARC/Installation/Finalizing/es")[Working with Gentoo](/wiki/Handbook:SPARC/Full/Working "Handbook:SPARC/Full/Working")[Portage introduction](/wiki/Handbook:SPARC/Working/Portage/es "Handbook:SPARC/Working/Portage/es")[USE flags](/wiki/Handbook:SPARC/Working/USE/es "Handbook:SPARC/Working/USE/es")[Portage features](/wiki/Handbook:SPARC/Working/Features/es "Handbook:SPARC/Working/Features/es")[Initscript system](/wiki/Handbook:SPARC/Working/Initscripts/es "Handbook:SPARC/Working/Initscripts/es")[Environment variables](/wiki/Handbook:SPARC/Working/EnvVar/es "Handbook:SPARC/Working/EnvVar/es")[Working with Portage](/wiki/Handbook:SPARC/Full/Portage "Handbook:SPARC/Full/Portage")[Files and directories](/wiki/Handbook:SPARC/Portage/Files/es "Handbook:SPARC/Portage/Files/es")[Variables](/wiki/Handbook:SPARC/Portage/Variables/es "Handbook:SPARC/Portage/Variables/es")[Mixing software branches](/wiki/Handbook:SPARC/Portage/Branches/es "Handbook:SPARC/Portage/Branches/es")[Additional tools](/wiki/Handbook:SPARC/Portage/Tools/es "Handbook:SPARC/Portage/Tools/es")[Custom package repository](/wiki/Handbook:SPARC/Portage/CustomTree/es "Handbook:SPARC/Portage/CustomTree/es")[Advanced features](/wiki/Handbook:SPARC/Portage/Advanced/es "Handbook:SPARC/Portage/Advanced/es")[OpenRC network configuration](/wiki/Handbook:SPARC/Full/Networking "Handbook:SPARC/Full/Networking")[Getting started](/wiki/Handbook:SPARC/Networking/Introduction/es "Handbook:SPARC/Networking/Introduction/es")[Advanced configuration](/wiki/Handbook:SPARC/Networking/Advanced/es "Handbook:SPARC/Networking/Advanced/es")[Modular networking](/wiki/Handbook:SPARC/Networking/Modular/es "Handbook:SPARC/Networking/Modular/es")[Wireless](/wiki/Handbook:SPARC/Networking/Wireless/es "Handbook:SPARC/Networking/Wireless/es")[Adding functionality](/wiki/Handbook:SPARC/Networking/Extending/es "Handbook:SPARC/Networking/Extending/es")[Dynamic management](/wiki/Handbook:SPARC/Networking/Dynamic/es "Handbook:SPARC/Networking/Dynamic/es")

## Contents

- [1Configuración automática de la red](#Configuraci.C3.B3n_autom.C3.A1tica_de_la_red)
  - [1.1Usar DHCP](#Usar_DHCP)
  - [1.2Probar la red](#Probar_la_red)
- [2Obtener información de la interfaz](#Obtener_informaci.C3.B3n_de_la_interfaz)
- [3Wireless Setup](#Wireless_Setup)
  - [3.1Optional: Checking wireless](#Optional:_Checking_wireless)
  - [3.2Using NetworkManager](#Using_NetworkManager)
- [4Usar net-setup](#Usar_net-setup)
  - [4.1Manual setup](#Manual_setup)
    - [4.1.1Manually connecting to an open network](#Manually_connecting_to_an_open_network)
- [5Opcional: configuración específica de la aplicación](#Opcional:_configuraci.C3.B3n_espec.C3.ADfica_de_la_aplicaci.C3.B3n)
  - [5.1Opcional: Configurar proxys web](#Opcional:_Configurar_proxys_web)
  - [5.2Usar pppoe-setup para ADSL](#Usar_pppoe-setup_para_ADSL)
  - [5.3Usar PPTP](#Usar_PPTP)
- [6Conceptos básicos de Internet e IP](#Conceptos_b.C3.A1sicos_de_Internet_e_IP)
  - [6.1Interfaces y direcciones](#Interfaces_y_direcciones)
  - [6.2Redes y CIDR](#Redes_y_CIDR)
  - [6.3Internet](#Internet)
  - [6.4El Sistema de Nombres de Dominio](#El_Sistema_de_Nombres_de_Dominio)
- [7Configuración de red manual](#Configuraci.C3.B3n_de_red_manual)
  - [7.1Configuración de dirección de interfaz](#Configuraci.C3.B3n_de_direcci.C3.B3n_de_interfaz)
  - [7.2Configuración de ruta predeterminada](#Configuraci.C3.B3n_de_ruta_predeterminada)
  - [7.3Configuración DNS](#Configuraci.C3.B3n_DNS)

## Configuración automática de la red

¿Es posible que simplemente funcione?

Si el sistema está conectado a una red Ethernet con un enrutador IPv6 o un servidor DHCP, es muy probable que la red del sistema se haya configurado automáticamente. Si no se requiere una configuración avanzada adicional, [se puede probar la conectividad a Internet](/wiki/Handbook:SPARC/Installation/Networking/es#Testing_the_network "Handbook:SPARC/Installation/Networking/es").

### Usar DHCP

DHCP (Protocolo de configuración dinámica de host) ayuda en la configuración de la red y puede proporcionar automáticamente la configuración para una variedad de parámetros que incluyen: dirección IP, máscara de red, rutas, servidores DNS, servidores NTP, etc.

DHCP requiere que un servidor se ejecute en el mismo segmento de "Capa 2" ("Ethernet") que el cliente que solicita una "concesión". DHCP se utiliza a menudo en redes RFC1918 ("privadas"), pero también se utiliza para adquirir información de IP pública de los ISP.

By default the Gentoo live media uses [NetworkManager](/wiki/NetworkManager "NetworkManager") which will automatically connect to via DHCP, if this is not the case then check the Ethernet cable for issues.

### Probar la red

Una ruta "predeterminada" correctamente configurada es un componente crítico de la conectividad a Internet; la configuración de la ruta se puede verificar con:

`root #` `ip route`

```
default via 192.168.0.1 dev enp1s0
```

Si no hay definida ninguna ruta "predeterminada", la conectividad a Internet no está disponible y se requiere configuración adicional.

La conectividad básica a Internet se puede confirmar con un ping:

`root #` `ping -c 3 1.1.1.1`

**Consejo**

Es útil comenzar haciendo ping a una dirección IP conocida en lugar de a un nombre de host. Esto puede aislar los problemas de DNS de los problemas básicos de conectividad a Internet.

El acceso HTTPS saliente y la resolución DNS se pueden confirmar con:

`root #` `curl --location gentoo.org --output /dev/null`

A menos que curl informe un error u otras pruebas fallen, el proceso de instalación puede continuar con [preparación del disco](/index.php?title=Handbook:SPARC/Instalaci%C3%B3n/Discos&action=edit&redlink=1 "Handbook:SPARC/Instalación/Discos (page does not exist)").

Si curl informa un error, pero los pings vinculados a Internet funcionan, [DNS puede necesitar configuración](/wiki/Handbook:SPARC/Installation/Networking/es#DNS_Configuration "Handbook:SPARC/Installation/Networking/es").

Si no se ha establecido la conectividad a Internet, primero [se debe verificar la información de la interfaz](/wiki/Handbook:SPARC/Installation/Networking/es#Obtaining_interface_info "Handbook:SPARC/Installation/Networking/es"), luego:

- [se puede utilizar net-setup](/wiki/Handbook:SPARC/Installation/Networking/es#Using_net-setup "Handbook:SPARC/Installation/Networking/es") para ayudar en la configuración de la red.
- Es posible que se requiera [configuración específica de la aplicación](/wiki/Handbook:SPARC/Installation/Networking/es#Optional:_Application_specific_configuration "Handbook:SPARC/Installation/Networking/es").
- Se puede intentar [la configuración de red manual](/wiki/Handbook:SPARC/Installation/Networking/es#Manual_network_configuration "Handbook:SPARC/Installation/Networking/es").

## Obtener información de la interfaz

Si la conexión en red no funciona de inmediato, se deben tomar medidas adicionales para habilitar la conectividad a Internet. Generalmente, el primer paso es enumerar las interfaces de red del host.

El comando ip, que forma parte del paquete [sys-apps/iproute2](https://packages.gentoo.org/packages/sys-apps/iproute2), se puede utilizar para consultar y configurar las redes del sistema.

El argumento **link** se puede utilizar para mostrar enlaces de interfaz de red:

`root #` `ip link`

```
1: lo: <LOOPBACK,UP,LOWER_UP> mtu 65536 qdisc noqueue state UNKNOWN mode DEFAULT group default qlen 1000
    link/loopback 00:00:00:00:00:00 brd 00:00:00:00:00:00
4: enp1s0: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc pfifo_fast state UP mode DEFAULT group default qlen 1000
    link/ether e8:40:f2:ac:25:7a brd ff:ff:ff:ff:ff:ff
```

El argumento **address** se puede utilizar para consultar la información de la dirección del dispositivo:

`root #` `ip address`

```
2: enp1s0: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc pfifo_fast state UP group default qlen 1000<pre>
    link/ether e8:40:f2:ac:25:7a brd ff:ff:ff:ff:ff:ff
    inet 10.0.20.77/22 brd 10.0.23.255 scope global enp1s0
       valid_lft forever preferred_lft forever
    inet6 fe80::ea40:f2ff:feac:257a/64 scope link
       valid_lft forever preferred_lft forever
```

El resultado de este comando contiene información para cada interfaz de red del sistema. Las entradas comienzan con el índice del dispositivo, seguido del nombre del dispositivo: enp1s0.

**Consejo**

Si no se muestra ninguna interfaz aparte de **lo** ( _loopack_), entonces el hardware de red está defectuoso o el controlador de la interfaz no se ha cargado en el kernel. Ambas situaciones van más allá del alcance de este Manual. Por favor solicite soporte en contacto [#gentoo](ircs://irc.libera.chat/#gentoo) ([webchat](https://web.libera.chat/#gentoo)).

Para mantener la coherencia, el manual asumirá que la interfaz de red principal se llama enp1s0.

**Nota**

Como resultado del cambio hacia [nombres de interfaz de red predecibles](https://www.freedesktop.org/wiki/Software/systemd/PredictableNetworkInterfaceNames/), el nombre de la interfaz en el sistema puede ser bastante diferente a la antigua convención de nomenclatura eth0. Los medios de arranque modernos de Gentoo utilizan nombres de interfaz con prefijos como eno0, ens1 o enp5s0.

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

**Nota**

Some wireless cards may have a device name of wlan0 or ra0 instead of wlp9s0. Run ip link to determine the correct device name.

### Using NetworkManager

The fastest way to enable wireless connectivity in the Gentoo Live media is by using the nmtui command and following the on-screen instructions to setup the wireless network.

`root #` `nmtui`

LiveGUI users can click on the network icon in the bottom right of the taskbar to also configure WiFi.

## Usar net-setup

En los casos en los que la configuración automática de la red no tiene éxito, el _medio de arranque_ de Gentoo proporciona scripts para ayudar en la configuración de la red. net-setup se puede utilizar para configurar información de red inalámbrica e IP,s estáticas.

`root #` `net-setup enp1s0`

net-setup hará algunas preguntas sobre el entorno de red y utilizará esa información para configurar wpa\_supplicant o _direccionamiento estático_.

**Importante**

El estado de la red debe ser [probado](/index.php?title=Handbook:SPARC/Instalaci%C3%B3n/Redes&action=edit&redlink=1 "Handbook:SPARC/Instalación/Redes (page does not exist)") después de realizar cualquier paso de configuración. En caso de que los scripts de configuración no funcionen, se requiere [configuración de red manual](/index.php?title=Handbook:SPARC/Instalaci%C3%B3n/Redes&action=edit&redlink=1 "Handbook:SPARC/Instalación/Redes (page does not exist)").

### Manual setup

Wireless networks may alternatively be set up with a daemon such as wpa\_supplicant or iwd. For more information on configuring wireless networking in Gentoo Linux, please read the [Wireless networking chapter](/wiki/Handbook:SPARC/Networking/Wireless/es "Handbook:SPARC/Networking/Wireless/es") in the Gentoo Handbook.}}

#### Manually connecting to an open network

For most users, there are only two settings needed to connect, the ESSID (aka wireless network name) and, optionally, the WEP key.

- First, ensure the interface is active:

`root #` `ip link set dev wlp9s0 up`

- To connect to an open network with the name _GentooNode_:

`root #` `iw dev wlp9s0 connect -w GentooNode`

## Opcional: configuración específica de la aplicación

**Consejo**

These steps are provided for users where using nmtui is not able to configure their network's needs.

Los siguientes métodos generalmente no son necesarios, pero pueden resultar útiles en situaciones en las que se requiere configuración adicional para la conectividad a Internet.

### Opcional: Configurar proxys web

Si se accede a internet a través de un proxy web, será necesario definir la información del proxy para que Portage acceda correctamente al proxy para cada protocolo admitido. Portage observa las variables de entorno http\_proxy, ftp\_proxy y RSYNC\_PROXY para descargar paquetes a través de su wget y rsync mecanismos de recuperación.

Ciertos navegadores web en modo texto, como links, también pueden utilizar variables de entorno que definen la configuración del proxy web; en particular para el acceso HTTPS, también será necesario definir la variable de entorno https\_proxy. Si bien Portage se verá afectado sin pasar parámetros adicionales en tiempo de ejecución durante la invocación, links requerirá que se establezcan configuraciones de proxy.

En la mayoría de los casos, es suficiente definir variables de entorno utilizando el nombre de host del servidor. En el siguiente ejemplo, se supone que el host del servidor proxy se llama proxy.gentoo.org y el puerto es 8080.

**Nota**

El símbolo `#` en los siguientes comandos es un comentario. Se ha agregado solo para mayor claridad y no es necesario escribirlo al ingresar los comandos.

Para definir un proxy HTTP (para tráfico HTTP y HTTPS):

`root #` `export http_proxy="http://proxy.gentoo.org:8080" # Se aplica a Portage y a Links
`

`root #` `export https_proxy="http://proxy.gentoo.org:8080" # Sólo aplica para Links
`

Si el proxy HTTP requiere autenticación, establezca un nombre de usuario y contraseña con la siguiente sintaxis:

`root #` `export http_proxy="http://username:password@proxy.gentoo.org:8080" # Se aplica a Portage y a Links
`

`root #` `export https_proxy="http://username:password@proxy.gentoo.org:8080" # Sólo aplica para Links
`

Inicie links usando los siguientes parámetros para soporte de proxy:

`user $` `links -http-proxy ${http_proxy} -https-proxy ${https_proxy} `

Para definir un proxy FTP para Portage y/o links:

`root #` `export ftp_proxy="ftp://proxy.gentoo.org:8080" # Se aplica a Portage y a Links`

Inicie links usando el siguiente parámetro para un proxy FTP:

`user $` `links -ftp-proxy ${ftp_proxy} `

Para definir un proxy RSYNC para Portage:

`root #` `export RSYNC_PROXY="proxy.gentoo.org:8080" # Se aplica a Portage; Links no admite un proxy rsync`

### Usar pppoe-setup para ADSL

Si se requiere PPPoE para acceder a Internet, el _medio de arranque_ de Gentoo incluye el script pppoe-setup para simplificar la configuración de ppp.

Durante la configuración, pppoe-setup solicitará:

- El nombre de la _interfaz_ Ethernet conectada al módem ADSL.
- El nombre de usuario y contraseña de PPPoE.
- IP del servidor DNS.
- Si se necesita o no un firewall.

`root #` `pppoe-setup
`

`root #` `pppoe-start`

En caso de error, se deben verificar las credenciales en /etc/ppp/pap-secrets o /etc/ppp/chap-secrets. Si las credenciales son correctas, se debe verificar la selección de la interfaz Ethernet para PPPoE.

### Usar PPTP

Si se necesita compatibilidad con PPTP, se puede utilizar pptpclient, pero requiere configuración antes de su uso.

Edite /etc/ppp/pap-secrets o /etc/ppp/chap-secrets para que contenga la combinación correcta de nombre de usuario y contraseña:

`root #` `nano /etc/ppp/chap-secrets`

Ajuste ahora /etc/ppp/options.pptp si es necesario:

`root #` `nano /etc/ppp/options.pptp`

Una vez completada la configuración, ejecute pptp (junto con las opciones que no se pudieron configurar en options.pptp) para conectar el servidor:

`root #` `pptp <server ipv4 address>`

## Conceptos básicos de Internet e IP

Si todo lo anterior falla, la red debe configurarse manualmente. Esto no es particularmente difícil, pero debe hacerse con reflexión. Esta sección sirve para aclarar la terminología y presentar a los usuarios conceptos básicos de redes relacionados con la configuración manual de una conexión a Internet.

**Consejo**

Algunos **CPE** ( **Equipo proporcionado por el operador**) combinan las funciones de un _enrutador_, _punto de acceso_, _módem_, _servidor DHCP_ y _servidor DNS_ en una sola unidad. Es importante diferenciar las funciones de un dispositivo del aparato físico.

### Interfaces y direcciones

Las _interfaces_ de red son representaciones lógicas de dispositivos de red. Una _interfaz_ necesita una _dirección_ para comunicarse con otros dispositivos en la _red_. Si bien sólo se requiere una única _dirección_, se pueden asignar varias direcciones a una única _interfaz_. Esto es especialmente útil para configuraciones de doble pila (IPv4 + IPv6).

Para mantener la coherencia, este manual asumirá que la interfaz enp1s0 utilizará la dirección 192.168.0.2.

**Importante**

Las direcciones IP se pueden configurar arbitrariamente. En consecuencia, es posible que varios dispositivos utilicen la misma dirección IP, lo que genera un _conflicto de direcciones_. Los conflictos de direcciones se deben evitar utilizando DHCP o SLAAC.

**Consejo**

IPv6 normalmente utiliza la configuración **S** tate **L** ess **A** ddress **A** uto **C** onfiguration (SLAAC) para la configuración de direcciones. En la mayoría de los casos, configurar manualmente las direcciones IPv6 es una mala práctica. Si se prefiere un sufijo de dirección específico, se pueden usar [tokens de identificación de interfaz](/index.php?title=IPv6_Statinc_Addresses_using_Tokens&action=edit&redlink=1 "IPv6 Statinc Addresses using Tokens (page does not exist)").

### Redes y CIDR

Una vez elegida una dirección, ¿cómo sabe un dispositivo cómo hablar con otros dispositivos?

Las _direcciones_ IP están asociadas con _redes_. Las _redes_ IP son rangos lógicos contiguos de direcciones.

La notación _Enrutamiento entre dominios sin clases_ o _CIDR_ se utiliza para distinguir los tamaños de red.

- El valor _CIDR_, a menudo anotado comenzando con **/**, representa el tamaño de la red.

  - La fórmula _2 ^ (32 - CIDR)_ se puede utilizar para calcular el tamaño de la red.
  - Una vez calculado el tamaño de la red, el número de nodos utilizables se debe reducir en 2.
    - La primera IP de una red es la _Dirección de red_ y la última suele ser la _Dirección de difusión_. Estas direcciones son especiales y no pueden ser utilizadas por hosts normales.

**Consejo**

Los valores _CIDR_ más comunes son **/24** y **/32**, que representan **254** nodos y un solo nodo respectivamente.

Un _CIDR_ de **/24** es el tamaño de red predeterminado de-facto. Esto corresponde a una máscara de subred de _255.255.255.0_, donde los últimos 8 bits están reservados para las direcciones IP de los nodos de una red.

La notación: 192.168.0.2/24 se puede interpretar como:

- La _dirección_ 192.168.0.2
- En la _red_ 192.168.0.0
- Con un tamaño de **254** ( _2 ^ (32 - 24) \- 2_)

  - Las IP,s utilizables están en el rango 192.168.0.1 - 192.168.0.254
- Con una _dirección de difusión_ de 192.168.0.255

  - En la mayoría de los casos, la última dirección de una red se utiliza como "dirección de difusión", pero esto se puede cambiar.

Con esta configuración, un dispositivo debería poder comunicarse con cualquier host en la misma _red_ (192.168.0.0).

### Internet

Una vez que un dispositivo está en una red, ¿cómo sabe cómo hablar con dispositivos en Internet?

Para comunicarse con dispositivos fuera de las _redes_ locales, se debe utilizar _enrutamiento_. Un _enrutador_ es simplemente un dispositivo de red que reenvía tráfico a otros dispositivos. El término _ruta predeterminada_ o _puerta de enlace_ normalmente se refiere a cualquier dispositivo de la red actual que se utilice para el acceso a la red externa.

**Consejo**

Es una práctica estándar hacer que la _puerta de enlace_ sea la primera o la última IP de una red.

Si hay un enrutador conectado a Internet disponible en 192.168.0.1, se puede utilizar como _ruta predeterminada_, proporcionando acceso a Internet.

Para resumir:

- Las interfaces deben configurarse con una _dirección_ y una _información de red_, como el valor _CIDR_.
- El acceso a la red local se utiliza para acceder a un _enrutador_ en la misma red.
- Se configura la _ruta predeterminada_, por lo que el tráfico destinado a redes externas se reenvía a la _puerta de enlace_, proporcionando acceso a Internet.

### El Sistema de Nombres de Dominio

Recordar IP,s es difícil. El _Sistema de Nombres de Dominio_ fue creado para permitir el mapeo entre _Nombres de Dominio_ y _direcciones IP_.

Los sistemas Linux utilizan /etc/resolv.conf para definir _servidores de nombres_ que se utilizarán para la _resolución DNS_.

**Consejo**

Muchos _enrutadores_ también pueden funcionar como servidores DNS, y el uso de un servidor DNS local puede aumentar la privacidad y acelerar las consultas mediante el almacenamiento en caché.

Muchos ISP ejecutan un servidor DNS que generalmente se anuncia en la _puerta de enlace_ a través de DHCP. El uso de un servidor DNS local tiende a mejorar la latencia de las consultas, pero la mayoría de los servidores DNS públicos arrojarán los mismos resultados, por lo que el uso del servidor se basa en gran medida en las preferencias.

## Configuración de red manual

### Configuración de dirección de interfaz

**Importante**

Al configurar manualmente las direcciones IP, se debe considerar la topología de la red local. Las direcciones IP se pueden configurar arbitrariamente; Los conflictos pueden causar interrupciones en la red.

Para configurar enp1s0 con la dirección 192.168.0.2 y CIDR /24:

`root #` `ip address add 192.168.0.2/24 dev enp1s0`

**Consejo**

El inicio de este comando se puede acortar a ip a.

### Configuración de ruta predeterminada

La configuración de la dirección y la información de red para una interfaz configurará rutas link, lo que permitirá la comunicación con ese segmento de red:

`root #` `ip route`

```
192.168.0.0/24 dev enp1s0 proto kernel scope link src 192.168.0.2
```

**Consejo**

Este comando se puede abreviar a ip r.

La ruta default se puede configurar en 192.168.0.1 con:

`root #` `ip route add default via 192.168.0.1`

### Configuración DNS

La información del servidor de nombres normalmente se adquiere mediante DHCP, pero se puede configurar manualmente agregando entradas `nameserver` a /etc/resolv.conf.

**Advertencia**

Si _dhcpcd_ se está ejecutando, los cambios en [Template:Ruta](/index.php?title=Template:Ruta&action=edit&redlink=1 "Template:Ruta (page does not exist)") no se mantendrán. El estado se puede comprobar con `ps x | grep dhcpcd`.

nano está incluido en el _medio de arranque_ de Gentoo y puede usarse para editar /etc/resolv.conf con:

`root #` `nano /etc/resolv.conf`

Las líneas que contienen la palabra clave `nameserver` seguida de la dirección IP del servidor DNS se consultan en orden de definición:

ARCHIVO **`/etc/resolv.conf`** **Usar Quad9 DNS.**

```
nameserver 9.9.9.9
nameserver 149.112.112.112

```

ARCHIVO **`/etc/resolv.conf`** **Usar Cloudflare DNS.**

```
nameserver 1.1.1.1
nameserver 1.0.0.1

```

El estado del DNS se puede verificar haciendo ping a un nombre de dominio:

`root #` `ping -c 3 gentoo.org`

Una vez que [se haya verificado](/wiki/Handbook:SPARC/Installation/Networking/es#Testing_the_network "Handbook:SPARC/Installation/Networking/es") la conectividad, continúe con [Preparando los discos](/wiki/Handbook:SPARC/Installation/Disks/es "Handbook:SPARC/Installation/Disks/es").

[← Elegir los medios](/wiki/Handbook:SPARC/Installation/Media/es "Previous part") [Inicio](/wiki/Handbook:SPARC/es "Handbook:SPARC/es") [Preparando los discos →](/wiki/Handbook:SPARC/Installation/Disks/es "Next part")