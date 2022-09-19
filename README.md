# Propuesta del problema

Cada día que pasa estamos más ocupados que el anterior, lo que nos puede llevar a **no tener tiempo** para realizar tareas cotidianas como ir a comprar o plantear un menú semanal.

A todo el mundo le ha pasado eso de **no saber qué comer**, ya sea porque no hemos pensado nada o porque tenemos un popurrí de ingredientes con el que no sabemos qué cocinar.

Es de **vital** importancia **comer bien**, ya que esto se reflejará en un futuro en nuestra salud.

Con los **precios subiendo** cada vez más, el **ciudadano medio** suele verse **ajustado** a la hora de realizar compras abundantes, con lo que **comprar lo justo** y necesario se vuelve aún más **importante**. Esto influye en nuestro problema principal, ya que al comprar lo justo, no se cuenta con *tanta cantidad de ingredientes*, lo que puede llevar a la frustración comentada inicialmente.

## Sintetización del problema

* Indecisión o falta de creatividad a la hora de cocinar -> Dietas pobres, monotonía en la comida.
* Subida en los precios de los alimentos -> Pocos ingredientes disponibles para realizar las comidas.

Estos problemas pueden tener un efecto muy negativo en la salud, tanto física como mental.

---

## Propuesta de la solución

Se propone un sistema que ofrezca la posibilidad al usuario de que **introduzca los ingredientes disponibles** (bien mediante el escaneo de un ticket de compra o mediante la introducción manual de los ingredientes) y le **proponga al usuario una lista de recetas** que se podrían hacer con esos ingredientes, junto al precio aproximado por persona.

En el caso de no poder hacer nada, le propondría **recetas para las que tuviera que comprar algún ingrediente**, los menos posibles.

En este último caso, se podrían **usar los datos extraídos de los tickets de todos los usuarios** (tendremos un histórico) para sugerir al usuario dónde realizar la compra para que le salga **lo más barato posible**.

El sitio a realizar la compra será cercano al usuario, estableciendo la ubicación aproximada del cliente (por ejemplo, podemos saber si una persona se conecta a nuestra web desde Granada capital o desde Armilla) para recomendarle un supermercado que no se encuentre demasiado lejos.

---

## Lógica de negocio detrás de la solución

Nuestro sistema:

* Importa una base de datos pública de recetas de cocina, expandiéndola con recetas propias.
* Transforma una foto de un ticket de compra, extrayendo el texto de la misma y usándolos como ingredientes a la hora de localizar la receta.
* Tiene en cuenta posibles faltas ortográficas comunes en la entrada manual de ingredientes (ceboya vs. cebolla, por ejemplo).
* Calcula una o más recetas que se ajusten a los ingredientes proporcionados por el usuario.
* Permite al usuario visualizar y extraer los ingredientes que le falten por comprar en caso de que no tuviera suficientes ingredientes.
* Valida que la(s) receta(s) proporcionada(s) coincida(n) en la medida de lo posible con los ingredientes del usuario.
* Calcula el precio aproximado de compra de los ingredientes restantes, en caso de tener suficientes datos para ello.

Con los puntos anteriores, le estaremos proporcionando al usuario (y cliente) recetas que podrá usar con los ingredientes que tenga disponible, así como donde comprar los ingredientes que no posea para que le resulte lo más barato posible.
Esto soluciona los problemas expuestos ya que:

* Le proporcionamos al usuario recetas que puede realizar, aportándole conocimiento que puede usar para comer de forma más variada y equilibrada.
* Le proporcionamos al usuario una sugerencia para ahorrar dinero en el caso de que quiera realizar alguna receta para la que le falte ingredientes, ayudándolo así a combatir la subida de precios.

### Flujo de los datos

1. Se introduce una foto o una lista de ingredientes.
2. Si es una foto:

   Se extraen los ingredientes mediante reconocimiento de texto, se añaden a la lista y se vuelve al primer paso. Los precios de los distintos elementos comprados se guardan para futuras comparativas de precios.

   Si es una lista de ingredientes, se sigue al paso 3.
3. Se hace una búsqueda de las posibles recetas que contengan el mayor número de ingredientes de los introducidos.
4. En el caso de faltar ingredientes, se hace comparativa de precio, indicando al usuario dónde le resultaría más barato comprar los ingredientes restantes para cada receta.
5. Los datos procesados en los pasos anteriores se juntaran para

---

## Posibles clientes

Los principales **clientes** serían **usuarios ocasionales** que usarían la aplicación para obtener ideas sobre qué comida cocinar en un momento determinado.

La aplicación web sería monetizable a través de la **venta de los datos de los usuarios** (dónde suelen comprar más, qué ingredientes se suelen comprar más, que marcas prefieren frente a otras...) a organizaciones y empresas.

---

## Por qué en la nube

El despliegue en la nube de cualquier sistema tiene las siguientes ventajas:

* **Disponibilidad**: No depender de un ordenador o raspberry en algún garaje favorece que el sistema esté siempre disponible. A su vez, no dependemos de un *ISP* para asegurar que el sistema se encuentra conectado a internet, ya que el sistema "**estaría**" en internet.
* **Facilidad de actualización**: La nube facilita mucho el uso de sistemas de CI/CD, como *Jenkins*, *Argo*... que promueven actualizaciones y despliegues rápidos y efectivos, haciendo disponible al usuario la última versión de forma sencilla para el desarrollador.
* **Seguridad**: La nube (ya sea en AWS, en Azure o similares) aporta una capa extra de seguridad a los datos, haciéndose cargo de que si el sistema no expone los datos directamente, sea difícil acceder a estos sin permiso.
* **Coste**: A día de hoy, los distintos proveedores de *Cloud* ofrecen máquinas que tienen en cuenta el consumo medio, o se expanden en potencia de procesado o memoria si es necesario, permitiendo a la empresa que hace el despliegue ahorrar en las horas en las que no tiene tanto tráfico.
* **Fácil acceso**: Las personas autorizadas, pueden acceder y depurar la aplicación en caso de fallo catastrófico sobre la marcha, debido a la disponibilidad mencionada anteriormente.

### Porqué este sistema debería estar en la nube

Este sistema nace con la mentalidad de que sea **independiente del sistema operativo**, es decir, que se pueda usar tanto para Android, como para IOS, como para Linux, Microsoft...

El estar en la nube permitiría una integración más sencilla con el frontend, además de quitar al usuario de problemas de espacio en su dispositivo.

Por otro lado, el desarrollo será más sencillo, ya que no tendremos que ajustarnos a todos los tipos de dispositivo, modificando el código para cada sistema operativo.

En resumen, no queremos una aplicación convencional, queremos algo que no depende del dispositivo, lo que nos lleva indudablemente a la nube.

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
