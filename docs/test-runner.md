# Elección del *test runner*

## Criterios para que nuestro *test runner* sea válido

Para que el *test runner* que escojamos sea válido, necesitamos que contenga las siguientes funcionalidades:

* Sintaxis cómoda para escribir los test.
* Automatizado (no tener que indicar de ninguna forma qué tests se han de correr por defecto).
* Pocas o ninguna dependencia.
* Instalación sencilla.
* Fácil integración con el código nuevo.
* Proporciona estadísticas sobre los tests ejecutados (tiempo de ejecución, tipo de fallo en caso de que ocurra..).
* Permite informar de código no cubierto en los tests (o conocer el [*code coverage*](https://www.tutorialspoint.com/test-coverage-in-software-testing)) de forma sencilla. Esto facilita el [*TDD*](https://www.agilealliance.org/glossary/tdd/).
* Permita generar datos aleatorios fácilmente y con una buena adaptabilidad, de forma que los tests usen datos aleatorios (mediante [*fuzzing*](https://en.wikipedia.org/wiki/Fuzzing)) o falsos ([*mocks*](https://devopedia.org/mock-testing)) no necesiten que estos sean introducidos a mano o tomados de fuentes externas. Esto facilita probar muchos casos de forma sencilla y que los tests sean completamente unitarios, eliminando las dependencias externas.
* Permita utilizar código previo a la ejecución de los tests y depués de los mismos a fin de preparar y limpiar el entorno (en los casos en los que esto sea necesario).

---

## Candidatos para *tests runners*

Hemos tenido cierta dificultad a la hora de encontrar *test runners* para Go, ya que Go cuenta con [su propio *test runner*](https://pkg.go.dev/testing), por lo que gran parte de la información encontrada en internet remite a este.

A falta de haber encontrado uno mejor (que no estuviera completamente dedicado a apis), usaremos el paquete *testing* de Go.

Este *test runner* cumple con casi todas las condiciones anteriormente mencionadas ya que:

* Ofrece una sintaxis cómoda.
* Detecta nuevos tests de forma automática (ya que se fija en el nombre de la función a evaluar, asegurándose que comieza por *Test*).
* No tiene dependencias, ya que viene por defecto con Go.
* No requiere instalación, ya que viene por defecto con Go.
* Permite conocer el tiempo de ejecución de los tests, así como el tipo de fallo que se ha producido. A su vez, es compatible con el uso de [*Benchmarks*](https://pkg.go.dev/testing#B).
* Permite ver el [*code coverage*](https://pkg.go.dev/testing#Coverage).
* Permite realizar *fuzzing*, pero el *mocking* no está realmente contemplado.
* Permite ejecutar código antes y después de la ejecución, mediante un [*main*](https://pkg.go.dev/testing#hdr-Main), pero su uso está bastante limitado.

Como vemos, cumple con casi todos los requisitos que creemos necesarios, por lo que es un buen *test runner* para nuestro criterio.

Ahora bien, necesitaremos apoyarnos en alguna librería para complementar las pequeñas carencias del paquete *testing*.

Inicialmente podemos servirnos del paquete básico de Go, ya que el mocking lo dejaremos para bastante más adelante, por lo que de momento no será un problema y la deuda técnica no será relevante.

Es más, podemos hacer *mocking* con el paquete de tests de go, como se nos explica en [este artículo](https://medium.com/@ankur_anand/how-to-mock-in-your-go-golang-tests-b9eee7d7c266), pero complica un poco más las cosas que si tuviéramos un paquete dedicado a ello, aunque por otro lado, al no tenerlo, también simplificamos y reducimos las dependencias de nuestro código.

---

## Control de las aserciones

Una vez elegido el *test runner*, podemos darnos cuenta viendo la documentación de que el control de las aserciones es un tanto precario, por lo que vamos a requerir de una librería que nos facilite esto.

Para nuestra suerte, tenemos un paquete de aserciones disponible, llamado gotest, del cual nos interesa su módulo [assert](https://pkg.go.dev/gotest.tools/assert).

Esta librería facilita enormemente las aserciones que tendremos que definir para nuestros tests. El paquete en cuestión está actualizado, como se puede ver en su [github](https://github.com/gotestyourself/gotest.tools) y está preparado para ser usado con *go modules*.

Al usar tan solo una parte de la librería completa, no debería de aumentar demasiado el peso de nuestro sistema.

