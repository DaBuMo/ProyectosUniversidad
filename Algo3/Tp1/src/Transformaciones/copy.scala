def copy(
    xslTag: Tag,
    xmlTagsGlobal: List[Tag],
    xmlTagsLocal: List[Tag]
): List[Either[Tag, String]] = {
  val pathStr = xslTag.params.getOrElse("select", "")
  val path = separatePath(pathStr)
  val resultados = searchByPath(path, xmlTagsGlobal, xmlTagsLocal)

  val childsList = xslTag.childs match {
    case Left(tags) => tags
    case Right(_)   => List.empty[Tag]
  }

  val cuerpo = procesarChilds(xslTag.childs, xmlTagsGlobal, resultados)

  resultados.map(tag => {
    Left(
      tag.copy(
        childs = cuerpo
      )
    )
  })

}
