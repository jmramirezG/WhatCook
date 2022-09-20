# Propuesta del problema

Las peñas y pequeñas asociaciones suelen hacer "barras" o pequeñas fiestas para los miembros de las mismas.

Es necesario llevar un control acerca de la contabilidad de la asociación o peña, lo cual se vuelve complicado, ya que o bien se lleva a papel (que acaba perdiéndose) o el software existente no está adaptado a la naturaleza de una peña/asociación.

Muchas peñas además, organizan "premios" repartidos entre aquellos que más consumen durante estas barras o fiestas, y llevar un control de esto resulta muy tedioso sin el software adecuado.

## Sintetización del problema

* Gestión de las cuentas de la asociación/peña durante las barras o fiestas que organiza.
* Registro de cúanto consume cada socio, ofreciendo estadísticas disponibles a la administración de la asociación/peña.

---

## Propuesta de la solución

Se propone un sistema que ofrezca a una asociación o peña prescindir de los tickets y del dinero físico, asignando a cada socio una tarjeta (bien con un QR o RFID), en la que el socio podrá cargar dinero en la administración de la asociación/peña.

Esta tarjeta se usará por el personal de barra para cobrar las consumiciones del socio, prescindiendo del dinero "real".
Esto permite llevar un registro completo de qué se ha gastado, quién ha consumido qué y cúanto ha consumido cada uno.

A su vez, se permite la exportación del histórico de las barras a un excel, permitiendo así el control integral de las cuentas relacionadas con las barras/pequeñas fiestas.

El administrador del sistemas podrá:

* Introducir el dinero que se encuentra en caja y el inventario de cada bebida.
* Dar de alta a nuevos socios, escaneando la tarjeta antes de dársela e introduciendo sus datos.
* Dar de baja a socios que dejen de formar parte de la asociación.
* Variar los precios de los productos en venta.
* Variar los productos en venta (añadir/quitar productos).
* Nombrar encargados de barra para un evento, ya que en las peñas suele variar entre fiestas quién se encarga de la barra.
* Obtener un excel con el estado actual de las cuentas en cualquier momento.
* Introducir gastos e inversiones, por ejemplo, con la compra de productos nuevos y pagos.
* Obtener un ranking de personas que más han consumido desde una fecha concreta.

El usuario podrá:

* Consultar su saldo en cualqier momento, mediante el escanéo del código QR o haciendo uso de un lector RFID.
* Consultar su gasto desde una fecha concreta.
* Consultar su posición en el ranking de consumiciones desde una fecha.

---

## Lógica de negocio detrás de la solución

Nuestro sistema trabaja con los socios y los datos de contabilidad de la peña/asociación.

El administrador del sistema se encargará de introducir los datos iniciales, mediante un proceso guiado para que resulte lo más sencillo posible, sin posibilidad de fallo.

Este será nuestro modelo, tendremos:

* Socios.
* Datos económicos.
* Información sobre los productos.

El personal de barra cobrará de las tarjetas RFID/QR a los socios, y los cambios quedarán reflejados a tiempo real para todos los usuarios.

El socio podrá recargar dinero, viéndose esto reflejado inmediatamente en el saldo de su tarjeta.

---

## Posibles clientes

Los clientes serían las peñas y asociaciones, que pagarían una subscripción por el accesso a las funcionalidades que ofrecemos.

---

## Por qué en la nube

El despliegue en la nube de cualquier sistema tiene las siguientes ventajas:

* **Disponibilidad**: No depender de un ordenador o raspberry en algún garaje favorece que el sistema esté siempre disponible. A su vez, no dependemos de un *ISP* para asegurar que el sistema se encuentra conectado a internet, ya que el sistema "**estaría**" en internet.
* **Facilidad de actualización**: La nube facilita mucho el uso de sistemas de CI/CD, como *Jenkins*, *Argo*... que promueven actualizaciones y despliegues rápidos y efectivos, haciendo disponible al usuario la última versión de forma sencilla para el desarrollador.
* **Seguridad**: La nube (ya sea en AWS, en Azure o similares) aporta una capa extra de seguridad a los datos, haciéndose cargo de que si el sistema no expone los datos directamente, sea difícil acceder a estos sin permiso.
* **Coste**: A día de hoy, los distintos proveedores de *Cloud* ofrecen máquinas que tienen en cuenta el consumo medio, o se expanden en potencia de procesado o memoria si es necesario, permitiendo a la empresa que hace el despliegue ahorrar en las horas en las que no tiene tanto tráfico.
* **Fácil acceso**: Las personas autorizadas, pueden acceder y depurar la aplicación en caso de fallo catastrófico sobre la marcha, debido a la disponibilidad mencionada anteriormente.

### Porqué este sistema debería estar en la nube

Este sistema necesita de actualizaciones en tiempo real de todos los valores, es decir, necesitamos que los datos se encuentren centralizados entre todos los usuarios. Tan solo con esto, ya se nos deja claro que es necesario que el despliegue del backend se realice en la nube.

---

## Docs

* [Objetivo 0](docs/objetivo0/git_config.md).

---
  
## Licencia

WhatCook

Copyright (C) 2022  José María Ramírez González \<jm.ramirez.gonza@gmail.com\>

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU Affero General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU Affero General Public License for more details.

You should have received a copy of the GNU Affero General Public License
along with this program.  If not, see <http://www.gnu.org/licenses/>.

---
