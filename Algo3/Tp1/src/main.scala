@main def main(args: String*) = {
  // OBTENER los archivos
  val xslPath = args(0)
  val xsl = leerArchivo(xslPath)

  val xmlPath = args(1)
  val xml = leerArchivo(xmlPath)

  val params = args
    .filter(_.startsWith("--param="))
    .map((x) => x.replace("--param=", ""))
    .toList

  // OBTENER TAGS
  val xslTags = textoATag(xsl)
  val xmlTags = textoATag(xml)

  val resultado = procesarVariables(xslTags, xmlTags, params)

  // TRANSFORMAR RESULTADO A XML
  val resultadoXML = tagsATexto(resultado)

  // .Out
  println(resultadoXML)
  escribirArchivo("resultado.xml", resultadoXML)
}
