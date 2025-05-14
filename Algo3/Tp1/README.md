# tp-1-paradigmas-grupo-45

Guia para correr el programa desde cero

Instalacion de Scala:

Ir al enlace "scala-lang.org/download/" y dependiendo el sistema operativo que uses, seguir las instrucciones marcadas por la pagina.
Se debe instalar la version 3 o superior de Scala y la version 8 o superior de JDK.

Chequeos de instalacion (en linux):

Para verificar en linux si instalamos correctamente scala, utilizaremos comandos en la terminal de linux. Estos son:

Para chequear la version de scala en se debe colocar el comando:

```scala -version```

Nos podria aparecer algo tal que:

```
Scala code runner version: 1.5.4
Scala version (default): 3.6.4
```

(La default es la que debe ser superior a 3)

Y para chequear la version de JDK debemos escribir:

```scala```

Con eso entraremos en el modo interactivo de scala. Pero en este caso solo lo queremos para ver la info sobre JDK, y nos aparecerá algo tal que:

```
Welcome to Scala 3.6.4 (11.0.26, Java OpenJDK 64-Bit Server VM).
Type in expressions for evaluation. Or try :help.
```

Lo que esta entre parentesis en la primer linea es la version de JDK (lo que debe ser mayor o igual a 8).

Compilacion:

La compilacion tendra que tener el siguiente formato:

```scalac src/*.scala src/Transformaciones/*.scala```

Ejecución:

Para la ejecucion del programa se requeriran los siguentes archivos que seran pasados por consola:

- Un archivo del tipo XML.
- Un archivo del tipo XSLT para la ejecución correcta.

La ejecucion sin parametros tendra que ser de la siguiente manera:

```scala main rutaXsl RutaXml```

La ejecucion con parametros tendra que ser de la siguiente manera:

```scala main rutaXsl RutaXml --param=pepe```

Un ejemplo de compilacion y ejecucion simultanea seria de la siguente manera:

```scala src/*.scala src/Transformaciones/*.scala -- src/archivo.xsl src/archivo.xml --param=Juan --param=maria```