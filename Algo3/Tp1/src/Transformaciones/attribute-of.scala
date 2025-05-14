def attributeOf(
    xslTag: Tag,
    xmlTagsGlobal: List[Tag],
    xmlTagsLocal: List[Tag]
): String = {

  val pathStr = xslTag.params.getOrElse("select", "")
  val nameAttribute = xslTag.params.getOrElse("name", "")
  val pathDesarmado = separatePath(pathStr)
  val tagsSelected = searchByPath(pathDesarmado, xmlTagsGlobal, xmlTagsLocal)
  val listaStrings =
    tagsSelected.map(tag => tag.params.getOrElse(nameAttribute, ""))
  listaStrings.isEmpty match
    case true => ""
    case _    => listaStrings.reduce((a, b) => a + b)

}
