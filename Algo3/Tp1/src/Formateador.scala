import scala.annotation.tailrec

// Representa un tag del XML con su clave, parámetros y sus hijos o contenido
case class Tag(
    key: String, // Nombre del tag
    tagType: String = "xml", // TODO: ES UN MOCK HAY QUE SACARLO
    params: Map[String, String] = Map(), // Parámetros de la etiqueta
    childs: Either[List[Tag], String] = Left(
      List()
    ) // Hijos (lista de nodos) o contenido (texto)
)

// Verifica si una línea es una etiqueta de apertura
def esApertura(tag: String): Boolean =
  tag.startsWith("<") && !tag.startsWith("</") && tag.endsWith(">")

// Verifica si una línea es una etiqueta de cierre
def esCierre(tag: String): Boolean =
  tag.startsWith("</") && tag.endsWith(">")

// Procesa una etiqueta de cierre, actualizando la lista de aperturas
def procesarCierre(
    tag: String,
    aperturas: List[Tag]
): (List[Tag], List[Tag]) = {
  val key = tag.replace("</", "").replace(">", "")
  aperturas match
    case current :: parent :: rest if current.key == key =>
      // Cierra el nodo actual y lo agrega como hijo del nodo padre
      val updatedParent = parent.copy(
        childs = Left(parent.childs.left.getOrElse(List()) :+ current)
      )
      (updatedParent :: rest, List())

    case x :: Nil if x.key == key =>
      // Cierra el nodo raíz y lo agrega a los nodos finales
      (List(), List(x))

    case _ =>
      throw new IllegalArgumentException(
        s"XML mal formado: etiqueta de cierre inesperada </$key>."
      )
}

// Procesa una etiqueta de apertura, creando un nuevo tag con sus parámetros
def procesarApertura(tag: String): Tag = {
  val content = tag.replace("<", "").replace(">", "").trim
  val parts = content.split(" ")

  val key = parts.head
  val params = parts.tail.filter((x) => !x.trim().isEmpty())

  val paramsMap = params
    .flatMap(param => {
      val keyValue = param.replace("\"", "").split("=")
      Map(keyValue(0) -> keyValue(1))
    })
    .toMap

  key.contains("pxsl") match
    case true =>
      Tag(key, "xsl", paramsMap)
    case false =>
      Tag(key, "xml", paramsMap)
}

def procesarContenido(
    content: String,
    aperturas: List[Tag]
): List[Tag] = {
  aperturas match
    case x :: xs =>
      val tagActual = x.copy(childs =
        Right(content.trim)
      ) // Actualiza el contenido del tag actual
      tagActual :: xs
    case Nil =>
      throw new IllegalArgumentException(
        s"Contenido fuera de etiquetas: $content"
      )
}

// Evalúa una línea del XML y actualiza la pila de aperturas
def evaluar(
    x: String,
    aperturas: List[Tag]
): (List[Tag], List[Tag]) = {
  x match
    case tag if esApertura(tag) =>
      val nuevoTag = procesarApertura(tag)
      (nuevoTag :: aperturas, List())

    case tag if esCierre(tag) =>
      procesarCierre(tag, aperturas)

    case tag =>
      val aperturasActualizadas = procesarContenido(tag, aperturas)
      (aperturasActualizadas, List())
}

// Procesa recursivamente las líneas del XML para construir la estructura de tags
@tailrec
def interpretarLineas(
    lineas: List[String], // Líneas restantes del XML
    aperturas: List[Tag], // Nodos abiertos
    nodosFinales: List[Tag] // Lista de nodos completamente procesados
): List[Tag] = {
  lineas match
    case Nil if aperturas.nonEmpty =>
      throw new IllegalArgumentException(
        "XML mal formado: faltan etiquetas de cierre."
      )
    case Nil => nodosFinales
    case x :: xs =>
      val (aperturasActualizadas, nuevosNodos) = evaluar(x, aperturas)
      interpretarLineas(xs, aperturasActualizadas, nodosFinales ++ nuevosNodos)
}

// Convierte un XML en una lista de Tags
def textoATag(xml: String): List[Tag] = {
  val lineas = xml
    .split("\n")
    .map(_.trim)
    .filter(_.nonEmpty)
    .toList

  interpretarLineas(lineas, List(), List())
}

def tagATexto(tag: Tag, depth: Int = 0): String = {
  val indent = "  " * depth // Dos espacios por nivel de profundidad
  val params =
    tag.params.map((param) => s"""${param._1}="${param._2}"""").mkString(" ")

  val formattedParams = params.nonEmpty match
    case true => s" $params"
    case _    => ""

  val childs = tag.childs match
    case Left(childTags) =>
      childTags.map(tagATexto(_, depth + 1)).mkString("\n")
    case Right(texto) =>
      s"${"  " * (depth + 1)}${texto.trim}" // Indentar el texto

  childs.isEmpty match
    case true => s"$indent<${tag.key}$formattedParams/>"
    case _ =>
      val openingTag = s"$indent<${tag.key}$formattedParams>"
      val closingTag = s"$indent</${tag.key}>"
      s"$openingTag\n$childs\n$closingTag"
}

def tagsATexto(
    resultado: List[Either[Tag, String]],
    depth: Int = 0
): String = {
  resultado
    .map {
      case Left(tag) => tagATexto(tag, depth)
      case Right(texto) =>
        s"${"  " * depth}${texto.trim}" // Indentar el texto plano
    }
    .mkString("\n")
}
