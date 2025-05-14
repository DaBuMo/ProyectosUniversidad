def sortByAttr(
    xslTag: Tag,
    xmlTagsGlobal: List[Tag],
    xmlTagsLocal: List[Tag]
): List[Either[Tag, String]] = {
  val name = xslTag.params.get("name").get
  val order = xslTag.params.get("order").get

  val cuerpo = procesarChilds(xslTag.childs, xmlTagsGlobal, xmlTagsLocal)

  val resultado: Either[List[Tag], String] = cuerpo match
    case Left(tag) =>
      val sortedTags = tag.sortBy(t => t.params.getOrElse(name, ""))
      order match
        case "desc" => Left(sortedTags.reverse)
        case _      => Left(sortedTags)

    case Right(texto) => Right(texto)

  resultado match {
    case Left(tags)  => tags.map(Left(_))
    case Right(text) => List(Right(text))
  }
}
