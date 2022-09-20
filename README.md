# Propuesta del problema

Las peñas y pequeñas asociaciones suelen hacer "barras" o pequeñas fiestas para los miembros de las mismas.

Es necesario llevar un control acerca de la contabilidad de la asociación o peña, lo cual se vuelve complicado, ya que o bien se lleva a papel (que acaba perdiéndose) o el software existente no está adaptado a la naturaleza de una peña/asociación. En muchos de estos casos, la gestión es tal que no se sabe cúanto se consume de cada cosa, comprando siempre por "intuición" a la hora de reponer el inventario.

## Sintetización del problema

* Gestión de las cuentas de la asociación/peña durante las barras o fiestas que organiza, prescindiendo del papel y solucionando posibles pérdidas de datos.
* Procesado de cúanto se consume y qué se consume, ofreciendo estadísticas disponibles a la administración de la asociación/peña, que posteriormente podrán usar para adaptar las compras según los informes generados.
* Dificultad a la hora de reponer el inventario debido a una mala gestión durante las fiestas.

---

## Propuesta de la solución

Se permite la generación de informes con los datos disponibles, que resuelvan problemas como los comentados anteriormente (falta de información sobre el inventario, cuentas a papel que desaparece, no saber qué comprar para la próxima fiesta).

Se pueden marcar ingresos periódicos, como serían la subscripción de los socios, indicando la cantidad y el día del mes en el que se recibe.

El administrador del sistemas podrá:

* Introducir el dinero que se encuentra en caja y el inventario de cada bebida.
* Variar los precios de los productos en venta.
* Variar los productos en venta (añadir/quitar productos).
* Obtener un informes en un formato manejable (Excel) con el estado actual de las cuentas en cualquier momento.
* Introducir gastos e inversiones, por ejemplo, con la compra de productos nuevos y pagos.
* Introducir la fecha de la siguiente fiesta/barra, para la que el sistema le hará una estimación de las cantidades a comprar de cada producto o si hay suficiente, al igual que una estimación de los beneficios y gastos.

---

## Lógica de negocio detrás de la solución

Nuestro sistema trabaja con los datos de contabilidad de la peña/asociación y con los datos de cada fiesta.

El administrador del sistema se encargará de introducir los datos, mediante un proceso guiado para que resulte lo más sencillo posible, sin posibilidad de fallo.

Este será nuestro modelo, tendremos:

* Datos económicos.
* Información sobre los productos.

Nuestro sistema llevará la contabilidad de la peña/asociación, a la vez que le generará informes y predicciones para futuras cuentas a los administradores de la peña/asociación.

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

Este sistema necesita de actualizaciones en tiempo real de todos los valores, es decir, necesitamos que los datos se encuentren centralizados entre todos los usuarios, ya que los administradores deben de ser capaces de interactuar a la vez con el sistema y desde distintos dispositivos. Tan solo con esto, ya se nos deja claro que es necesario que el despliegue del backend se realice en la nube.

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
