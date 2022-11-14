# Elección del gestor de dependencias

## Criterios de qué consideramos un buen gestor de dependencias

* Facilidad de uso.
* Actualizado.
* Poca instalación necesaria.
* Documentación abundante.

---

## Gestores de dependencias considerados

Vamos a considerar los 4 siguientes gestores de dependencias:

* [dep](https://github.com/golang/dep).
* [glide](https://github.com/Masterminds/glide).
* [gopkg](https://github.com/bytedance/gopkg).
* [Go Modules](https://go.dev/blog/using-go-modules).

---

### DEP

[dep](https://github.com/golang/dep) era un "experimento oficial" para implementar un gestor de paquetes para Go, tal y cómo explican en su página de github.

A día de hoy se encuentra totalmente desactualizado y archivado, siendo la última actualización recibida a finales de 2020.

Se descartó su uso debido a la aparición de [Go Modules](https://go.dev/blog/using-go-modules).

---

### Glide

[glide](https://github.com/Masterminds/glide) es un gestor de paquetes como ```npm```, ```cargo```... pero para Go, como se explica en su github.

Lleva desde 2019 sin actualizarse, como vemos en su github, aunque no dieron ninguna explicación al respecto ni archivaron el repositorio como hizo [dep](https://github.com/golang/dep).

---

### Gopkg

[gopkg](https://github.com/bytedance/gopkg) ofrece un conjunto de utilidades para Go, incluyéndose un gestor de dependencias.

Sigue recibiendo actualizaciones periódicas, siendo la última del 28 de octubre (en la rama main).

Debido a que engloba más utilidades, no solo un gestor de dependencias, la curva de aprendizaje probablemente sea mayor que si solo fuera un gestor de dependencias.

Además, puesto que migra de código de [ByteDance](https://en.wikipedia.org/wiki/ByteDance), puede que la documentación no se encuentre 100% actualizada con respecto a Go o pueda tener algunos defectos.

---

## Go Modules

[Go Modules](https://go.dev/blog/using-go-modules) es el gestor de dependencias que está disponible de forma nativa a Go desde la versión 1.11.

Al ser oficial, tenemos abundante documentación y está actualizado al día.

Su uso resulta sencillo, teniendo manuales bastante extensos en la propia web de Go.

## Qué vamos a usar

Como gestor de dependencias vamos a usar [Go Modules](https://go.dev/blog/using-go-modules), como hemos comentado en el issue #12.

El hecho de que sea el gestor de paquetes oficial asegura:

* Se podrá usar en las últimas versiones de Go, recibiendo actualizaciones con cada versión mientras dure el proyecto.
* El soporte será bueno, ya que lo desarrollan las mismas personas que el lenguaje.

La sintaxis no es compleja y realiza bastantes tareas de forma automática, como podemos ver en los propios ejemplos.

## Desventajas de otros gestores de dependencia

Actualmente la mayoría de estos gestores de dependencia se encuentran totalmente en desuso y están desactualizados (excepto gopkg, que si recibe actualizaciones más o menos constantes), siendo Go Modules el gestor al que migró la mayoría de la gente.
Siendo esta la principal desventaja respecto a ```dep``` y a ```glide```.

Respecto a Gopkg, no hay desventajas destacables más allá de una sintaxis distinta a la de Go Modules y el hecho de que no esté desarrollado por el mismo equipo de Go. Al ser Go Modules el oficial desde la versión 1.11, tenemos más garantías en este que en Gopkg.
