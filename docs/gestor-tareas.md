# Elección del gestor de tareas

## Criterios de qué consideramos un buen gestor de tareas

* Facilidad de uso.
* Actualizado.
* Poca instalación necesaria.
* Documentación abundante.

---

Como gestor de tareas hemos elegido [Task](https://taskfile.dev/), como hemos comentado en el issue #11.

## Por qué Task?

Task tiene una sintaxis simple, lo que hace la curva de aprendizaje bastante menos pronunciada.

Es bastante configurable, permitiendo una buena adaptabilidad.

Está muy en uso actualmente, recibiendo actualizaciones continuas por parte del equipo de desarrollo en [su github](https://github.com/go-task/task).

Tiene bastante documentación y una comunidad muy activa.

## Desventajas de otros gestores de tarea

[Gilbert](https://github.com/go-gilbert/gilbert) no es una mala opción, como ya explicamos en #11, sin embargo, la sintaxis resulta un tanto más complicada, haciendo más brusca la curva de aprendizaje.

Trae bastantes configuraciones por defecto y admite plugins, lo cual nos proporcionaría una buena adaptabilidad a nuestras necesidades, sin embargo, el tema de la sintaxis hace que no sea una buena apuesta.
Además, lleva sin actualizarse varios años, lo que hace, aún más, que no sea la mejor opción.

[Realize](https://github.com/oxequa/realize), similar a Gilbert, pero además cuenta con una función llamada *live reload*, en la que podemos ver los cambios a tiempo real según vamos realizando cambios en el código. Esta función es una gran ventaja a la hora de desarrollar, sin embargo, al igual que Gilbert, la sintaxis es más compleja que la de Task, por lo que la curva de aprendizaje sería mayor que la de este último.
También lleva varios años sin recibir ninguna actualización.
