def copyOf(
    xslTag: Tag,
    xmlTagsGlobal: List[Tag],
    xmlTagsLocal: List[Tag]
): List[Either[Tag, String]] = {

  val pathStr = xslTag.params.getOrElse("select", "")
  val path = separatePath(pathStr)
  val resultados = searchByPath(path, xmlTagsGlobal, xmlTagsLocal)
  resultados.map(xslTag =>
    Left(xslTag)
  ) // Devolver cada tag encontrado como Tag completo
}
