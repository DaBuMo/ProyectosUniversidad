def procesarTagHijo(
    tagHijo: Tag,
    xmlTagsGlobal: List[Tag],
    xmlTagsLocal: List[Tag]
): Either[List[Tag], String] = {
  // Procesa los hijos del tag y los actualiza
  val childsProcesados = procesarChilds(
    tagHijo.childs,
    xmlTagsGlobal,
    xmlTagsLocal
  )
  val tagActualizado = tagHijo.copy(childs = childsProcesados)

  // Procesa el tag y devuelve el resultado
  val resultadoDeTags = procesarTagXsl(
    tagActualizado,
    xmlTagsGlobal,
    xmlTagsLocal
  )

  // Valida que el resultado se homogeneo (solo tags o solo texto)
  val hayTags = resultadoDeTags.exists((x) => x.isLeft)
  val resultadoDeTagsValidado: Either[List[Tag], String] =
    hayTags match
      case true => Left(resultadoDeTags.filter(_.isLeft).map(_.left.get))
      case _ =>
        val listaStrings = resultadoDeTags
          .filter(_.isRight)
          .map(_.right.get)

        Right(listaStrings.isEmpty match
          case true => ""
          case _    => listaStrings.reduce((a, b) => a + b)
        )

  resultadoDeTagsValidado
}

def procesarTagsHijos(
    listaDeHijos: List[Tag],
    xmlTagsGlobal: List[Tag],
    xmlTagsLocal: List[Tag]
): Either[List[Tag], String] = {
  // Procesa todos los hijos del tag y retorna una lista de los resultados
  val resultadosDeLosHijos =
    listaDeHijos.map((x) => {
      x.key match
        case "pxsl:copy" | "pxsl:template" =>
          // Si es un copy o un template, se procesa el tag primero
          val tagProcesado: Either[List[Tag], String] =
            val resultado = procesarTagXsl(x, xmlTagsGlobal, xmlTagsLocal)
            val hayTags = resultado.exists(_.isLeft)
            hayTags match
              case true =>
                Left(resultado.filter(_.isLeft).map(_.left.get))
              case false =>
                Right(resultado.filter(_.isRight).map(_.right.get).mkString)
          tagProcesado
        case _ =>
          // Si no se procesan los hijos
          procesarTagHijo(x, xmlTagsGlobal, xmlTagsLocal)

    })

  // Valida que el resultado se homogeneo (solo tags o solo texto)
  val resultadoValidado: Either[List[Tag], String] =
    // evaluea una tupla, donde la primera parte es una lista de tags y la segunda una de textos
    // si hay tags, devuelve una lista de tags, sino devuelve un texto
    resultadosDeLosHijos.partition(_.isLeft) match {
      case (tags, textos) if tags.nonEmpty =>
        Left(tags.flatMap(_.left.getOrElse(Nil)))
      case (_, textos) =>
        Right(textos.flatMap(_.right.toOption).mkString)
    }

  resultadoValidado
}

def procesarChilds(
    childs: Either[List[Tag], String],
    xmlTagsGlobal: List[Tag],
    xmlTagsLocal: List[Tag]
): Either[List[Tag], String] = {
  val childsProcesados = childs match
    case Left(listDeHijos) => // Si es una lista de tags
      procesarTagsHijos(
        listDeHijos,
        xmlTagsGlobal,
        xmlTagsLocal
      )
    case Right(texto) =>
      Right(texto) // Si es un texto lo devuelvo tal cual

  childsProcesados
}
