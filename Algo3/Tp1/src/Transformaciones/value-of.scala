def valueOf(
    xslTag: Tag,
    tagsGlobales: List[Tag],
    tagsLocales: List[Tag]
): String = {

  val path = xslTag.params.get("select").getOrElse("")
  val pathDesarmado = separatePath(path)
  val tagsSeleccionados =
    searchByPath(pathDesarmado, tagsGlobales, tagsLocales)

  val resultado = tagsSeleccionados
    .flatMap(elem =>
      elem.childs match {
        case Left(hijos)      => List("") // TODO convertir tag a string
        case Right(contenido) => List(contenido.trim)
      }
    )

  resultado.isEmpty match {
    case true  => ""
    case false => resultado.reduce((a, b) => a + b)
  }
}
