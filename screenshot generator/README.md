# Screenshot Generator

```
Este proyecto fue movido a https://github.com/caefisica/web, en /assets/go
```

Este es un sencillo script Golang que genera capturas de pantalla de una lista de sitios web. Utiliza el paquete chromedp para ejecutar un navegador Chrome en modo headless y toma una captura de pantalla de cada sitio web de la lista.

## Requisitos

Para ejecutar este script, necesitarás tener Go y el paquete chromedp instalados en tu sistema. Puedes instalar Go siguiendo las instrucciones del sitio web oficial, y puedes instalar el paquete chromedp ejecutando el siguiente comando:

`go get -u github.com/chromedp/chromedp`

También necesitarás tener Google Chrome o Chromium instalado en tu sistema, ya que el script utiliza el binario de Chrome o Chromium para generar las capturas de pantalla.

```
sudo apt update
sudo apt install wget
wget https://dl.google.com/linux/direct/google-chrome-stable_current_amd64.deb
sudo apt install ./google-chrome-stable_current_amd64.deb
```

## Uso

Para utilizar este script, sigue estos pasos:

1. Crea un archivo de texto con la lista de sitios web en el mismo directorio que `main.go`.
2. Abra una ventana de terminal y navegue hasta el directorio que contiene el script y el archivo lists.txt.
3. Ejecute el script ejecutando el siguiente comando:

    `go ejecutar screenshot.go`

El script abrirá un navegador Chrome _headless_ y navegará a cada sitio web en el archivo lists.txt, tomará una captura de pantalla del sitio web, y guardará la captura de pantalla en una carpeta llamada _results_.

Por defecto, las capturas de pantalla se tomarán en un tamaño de ventana de 1920 x 1080. Si desea cambiar el tamaño de la ventana o cualquier otra opción, puede modificar el archivo `main.go`.
