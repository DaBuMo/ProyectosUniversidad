def templateMatch(
    xslTag: Tag,
    xmlTagsGlobal: List[Tag],
    xmlTagsLocal: List[Tag]
): List[Either[Tag, String]] = {
  val pathStr = xslTag.params.getOrElse("match", "")
  val path = separatePath(pathStr)
  val resultados = searchByPath(path, xmlTagsGlobal, xmlTagsLocal)

  val childsList = xslTag.childs match {
    case Left(tags) => tags
    case Right(_)   => List.empty[Tag]
  }

  val resultado = resultados
    .map((tagObtenido) =>
      procesarChilds(Left(childsList), xmlTagsGlobal, List(tagObtenido))
    )

  // Valida que el resultado se homogeneo (solo tags o solo texto)
  val resultadoValidado: Either[List[Tag], String] =
    // evaluea una tupla, donde la primera parte es una lista de tags y la segunda una de textos
    // si hay tags, devuelve una lista de tags, sino devuelve un texto
    resultado.partition(_.isLeft) match {
      case (tags, textos) if tags.nonEmpty =>
        Left(tags.flatMap(_.left.getOrElse(Nil)))
      case (_, textos) =>
        Right(textos.flatMap(_.right.toOption).mkString)
    }

  resultadoValidado match {
    case Left(tags)  => tags.map(Left(_))
    case Right(text) => List(Right(text))
  }
}
