def sortByElem(
    xslTag: Tag,
    xmlTagsGlobal: List[Tag],
    xmlTagsLocal: List[Tag]
): List[Either[Tag, String]] = {
  val elem = xslTag.params.get("elem").get
  val order = xslTag.params.get("order").get

  val cuerpo: Either[List[Tag], String] =
    procesarChilds(xslTag.childs, xmlTagsGlobal, xmlTagsLocal)

  val resultado: Either[List[Tag], String] = cuerpo match
    case Left(tags) => // si el cuerpo es una lista de tags
      val sortedTags = tags.sortBy(tag =>
        val optTag =
          searchByPath(separatePath(elem), List(tag), List(tag)).headOption

        val valorTag = optTag match
          // si no  hay tag o si hay tag y el valor es un tag, se considera vacio
          case Some(tag) =>
            tag.childs match
              case Left(tags)   => ""
              case Right(texto) => texto
          case None => ""

        valorTag
      )
      // ordenar por el valor del tag
      Left(order match
        case "desc" => sortedTags.reverse
        case _      => sortedTags
      )

    case Right(texto) => Right(texto) // si es un texto, no se hace nada

  resultado match {
    case Left(tags)  => tags.map(Left(_))
    case Right(text) => List(Right(text))
  }
}
