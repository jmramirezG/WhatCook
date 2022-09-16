# Propuesta del problema

Cada día que pasa estamos más ocupados que el anterior, lo que nos puede llevar a **no tener tiempo** para realizar tareas cotidianas como ir a comprar o plantear un menú semanal.

A todo el mundo le ha pasado eso de **no saber qué comer**, ya sea porque no hemos pensado nada o porque tenemos un popurrí de ingredientes con el que no sabemos qué cocinar.

Es de **vital** importancia **comer bien**, ya que esto se reflejará en un futuro en nuestra salud.

También está el **problema económico**: ¿dónde compro para que el precio sea el menor posible?
Con los **precios subiendo** cada vez más, el **ciudadano medio** suele verse **ajustado** a la hora de realizar compras abundantes, con lo que **comprar lo justo** y necesario se vuelve aún más **importante**.

---

## Propuesta de la solución

Se propone un sistema que ofrezca la posibilidad al usuario de que **introduzca los ingredientes disponibles** (bien mediante el escaneo de un ticket de compra o mediante la introducción manual de los ingredientes) y le **proponga al usuario una lista de recetas** que se podrían hacer con esos ingredientes, junto al precio aproximado por persona.

En el caso de no poder hacer nada, le propondría **recetas para las que tuviera que comprar algún ingrediente**, los menos posibles.

En este último caso, se podrían **usar los datos extraídos de los tickets** para sugerir al usuario dónde realizar la compra para que le salga **lo más barato posible**.

---

## Posibles clientes

Los principales **clientes** serían **usuarios ocasianales** que usarían la aplicación para obtener ideas sobre qué comida cocinar en un momento determinado.

La aplicación web sería monetizable a través de la **venta de los datos de los usuarios** (dónde suelen comprar más, qué ingredientes se suelen comprar más, que marcas prefieren frente a otras...) a organizaciones y empresas.

---

## Por qué en la nube

El despliegue en la nube de una aplicación con estas características tiene las siguientes ventajas:

* **Disponibilidad**: No depender de un ordenador o raspberry en algún garaje favorece que el sistema esté siempre disponible. A su vez, no dependemos de un *ISP* para asegurar que el sistema se encuentra conectado a internet, ya que el sistema "**estaría**" en internet.
* **Facilidad de actualización**: La nube facilita mucho el uso de sistemas de CI/CD, como *Jenkins*, *Argo*... que promueven actualizaciones y despliegues rápidos y efectivos, haciendo disponible al usuario la última versión de forma sencilla para el desarrollador.
* **Seguridad**: La nube (ya sea en AWS, en Azure o similares) aporta una capa extra de seguridad a los datos, haciéndose cargo de que si el sistema no expone los datos directamente, sea difícil acceder a estos sin permiso.
* **Coste**: A día de hoy, los distintos proveedores de *Cloud* ofrecen máquinas que tienen en cuenta el consumo medio, o se expanden en potencia de procesado o memoria si es necesario, permitiendo a la empresa que hace el despliegue ahorrar en las horas en las que no tiene tanto tráfico.
* **Fácil acceso**: Las personas autorizadas, pueden acceder y depurar la aplicación en caso de fallo catastrófico sobre la marcha, debido a la disponibilidad mencionada anteriormente.

---

## Docs

* [Objetivo 0](docs/objetivo0/git_config.md).