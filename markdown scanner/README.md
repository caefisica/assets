# Extractor de enlaces Markdown

Esta es una herramienta escrita en Golang que busca el campo de enlace en archivos Markdown y guarda los valores de enlace en un archivo de texto. Está pensado para ser usado con Screenshot Generator. Probablemente tendrías que modificar algunas partes, como el nombre de la carpeta en main.go para que te sea útil.

## Uso

Para utilizar esta herramienta, sigue estos pasos:

1. Clona este repositorio en tu ordenador local.
2. Abre una y ve al directorio donde clonaste el repositorio.
3. Modifica las variables `directory` y `outputPath` en el archivo `main.go` para especificar la carpeta en la que desea buscar los archivos Markdown y el archivo de salida, respectivamente.
4. Ejecuta el archivo `main.go` utilizando el comando `go run main.go`.

La herramienta buscará todos los archivos Markdown dentro de la carpeta especificada y sus subdirectorios, y guardará los valores de enlace en el archivo de salida especificado.

## Requisitos

Para ejecutar esta herramienta, necesitas tener instalado lo siguiente:

- Golang (versión 1.16 o posterior)
