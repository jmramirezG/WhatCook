# Elección del gestor de tareas

## Criterios de qué consideramos un buen gestor de tareas

* Facilidad de uso.
* Actualizado.
* Poca instalación necesaria.
* Documentación abundante.

---

## Gestores de tareas considerados

Vamos a considerar los 4 siguientes gestores de dependencias:

* [Gilbert](https://github.com/go-gilbert/gilbert).
* [Realize](https://github.com/oxequa/realize).
* [Task](https://taskfile.dev/).

---

### Gilbert

[Gilbert](https://github.com/go-gilbert/gilbert) es un gestor de tareas estilo ```Graden``` y ```Maven```, buscando una sintaxis declarativa.

LLeva desde junio de 2019 sin recibir ninguna actualización, por lo que no tiene mucho soporte a día de hoy.

---

### Realize

[Realize](https://github.com/oxequa/realize) ofrece una doble función de *live reload* y gestor de tareas.

Esta primera función que ofrece puede resultar muy interesante, no obstante, lleva sin actualizarse desde mayo de 2020, por lo que podemos intuir que a día de hoy no contará con mucho soporte.

---

### Task

[Task](https://taskfile.dev/) es un gestor de tareas, pensado para Go, que busca ser sencillo de usar, tal y como se describe en su web.

No tiene ninguna dependencia, puesto que está escrito en Go, simplemente es descargar y ejecutar.

La sintaxis resulta bastante sencilla y es fácil de implementar en *pipelines* de *CI*.

---

## Qué vamos a usar

Como gestor de tareas hemos elegido [Task](https://taskfile.dev/), como hemos comentado en el issue #11.

Los principales motivos de esta decisión son:

* Task tiene una sintaxis simple, lo que hace la curva de aprendizaje bastante menos pronunciada.
* Es bastante configurable, permitiendo una buena adaptabilidad.
* Está muy en uso actualmente, recibiendo actualizaciones continuas por parte del equipo de desarrollo en [su github](https://github.com/go-task/task).
* Tiene bastante documentación y una comunidad muy activa.

## Desventajas de otros gestores de tarea

[Gilbert](https://github.com/go-gilbert/gilbert) no es una mala opción, como ya explicamos en #11, sin embargo, la sintaxis resulta un tanto más complicada, haciendo más brusca la curva de aprendizaje.

Trae bastantes configuraciones por defecto y admite plugins, lo cual nos proporcionaría una buena adaptabilidad a nuestras necesidades, sin embargo, el tema de la sintaxis hace que no sea una buena apuesta.
Además, lleva sin actualizarse varios años, lo que hace, aún más, que no sea la mejor opción.

[Realize](https://github.com/oxequa/realize), similar a Gilbert, pero además cuenta con una función llamada *live reload*, en la que podemos ver los cambios a tiempo real según vamos realizando cambios en el código. Esta función es una gran ventaja a la hora de desarrollar, sin embargo, al igual que Gilbert, la sintaxis es más compleja que la de Task, por lo que la curva de aprendizaje sería mayor que la de este último y, además, lleva también varios años sin recibir ninguna actualización.
