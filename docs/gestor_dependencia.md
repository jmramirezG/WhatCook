# Elección del gestor de dependencias

Como gestor de dependencias hemos elegido [Go Modules](https://go.dev/blog/using-go-modules), como hemos comentado en el issue #12.

## Por qué Go Modules?

Go Modules es nativo de Go desde la versión 1.11, estando desarrollado por el mismo equipo.
Este hecho nos asegura lo siguiente:

* Se podrá usar en las últimas versiones de Go, recibiendo actualizaciones con cada versión mientras dure el proyecto.
* El soporte será bueno, ya que lo desarrollan las mismas personas que el lenguaje.
* La comunidad advocará por su uso, pues es nativo al propio Go.

## Desventajas de otros gestores de dependencia

Antes de la versión 1.11 de Go solían usarse varios gestores de dependencia: [dep](https://github.com/golang/dep), [glide](https://github.com/Masterminds/glide) y [gopkg](https://github.com/bytedance/gopkg).

Actualmente estos gestores de dependencia se encuentran totalmente en desuso y están desactualizados (excepto gopkg, que si recibe actualizaciones más o menos constantes), siendo Go Modules el gestor al que migró la mayoría de la gente.

Debido a la migración del grueso de desarrolladores a Go Modules, estos gestores quedaron en el olvido, siendo casi nulo su uso actualmente y, por tanto, la documentación disponible en internet.
