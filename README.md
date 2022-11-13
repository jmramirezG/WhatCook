# Whatcook

[![](https://shields.io/badge/LICENSE-AGPL-green?logo=readthedocs&style=for-the-badge)](#) [![](https://shields.io/badge/status-PRE--ALPHA-orange?logo=githubsponsors&style=for-the-badge)](#)

## Instalar go

Para instalar go, nos basta con seguir esta [guía oficial](https://go.dev/doc/install) de instalación.

---

## Ejecutar tareas automatizadas

Tenemos diferentes tareas automatizadas mediante nuestro gestor de tareas, [Task](https://taskfile.dev/), que tenemos que tener instalado para ejecutar estas tareas.

La guía de instalación la podemos ver [aquí](https://taskfile.dev/installation/).

Podemos:

### Comprobar que el proyecto no tiene errores sintácticos

Tenemos que colocarnos en el directorio raíz y ejecutar:

```bash
task check
```

---

## Ejecutar el proyecto

Para ejecutar el proyecto tenemos dos opciones:

* Hacer build y luego ejecutar:
  
  ```bash
  go build && ./<nombreejecutable>
  ```

* Hacer build y ejecutar en la misma orden:
  
  ```bash
  go run .
  ```

## Más documentación

* [Descripción del problema](docs/old_README.md)
* [Objetivo 0](docs/objetivo0/git_config.md).
* [Personas](docs/Personas.md).
* [Milestones](docs/milestones.md).
* [Historias de usuario](docs/user-stories.md).
* [Elección del gestor de tareas](docs/gestor-tareas.md).
* [Elección del gestor de dependencias](docs/gestor-gestor_dependencia.md).

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
