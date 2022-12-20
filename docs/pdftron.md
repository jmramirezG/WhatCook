# Guía de instalación de PDFTron

Instalar PDFTron puede ser algo complicado, no obstante, si seguimos la [guía oficial](https://www.pdftron.com/documentation/go/get-started/), en la que tenemos instrucciones para todos los sistemas operativos, resulta bastante sencillo.

---

## Dificultades encontradas

Hemos encontrado ciertas dificultades a la hora de ejecutar código que usase esta librería, ya que no nos encontraba el fichero *.so* de la misma.

Finalmente hemos conseguido solucionar este error añadiendo al *path* el directorio necesario de la siguiente manera:

```bash
export LD_LIBRARY_PATH=$LD_LIBRARY_PATH:/<YourPathToThisDir>/PDFNetC/Lib
export DYLD_LIBRARY_PATH=$DYLD_LIBRARY_PATH:/<YourPathToThisDir>/PDFNetC/Lib
```
