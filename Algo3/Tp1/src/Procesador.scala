def procesarTagXsl(
    xslTag: Tag,
    xmlTagsGlobal: List[Tag],
    xmlTagsLocal: List[Tag]
): List[Either[Tag, String]] = {
  xslTag.key match
    case "pxsl:value-of" =>
      List(
        Right(
          valueOf(
            xslTag,
            xmlTagsGlobal,
            xmlTagsLocal
          )
        )
      )
    case "pxsl:copy-of" =>
      copyOf(
        xslTag,
        xmlTagsGlobal,
        xmlTagsLocal
      )
    case "pxsl:template" =>
      templateMatch(
        xslTag,
        xmlTagsGlobal,
        xmlTagsLocal
      )
    case "pxsl:copy" =>
      copy(
        xslTag,
        xmlTagsGlobal,
        xmlTagsLocal
      )
    case "pxsl:attribute-of" =>
      List(
        Right(
          attributeOf(
            xslTag,
            xmlTagsGlobal,
            xmlTagsLocal
          )
        )
      )
    case "pxsl:sort-by-attr" =>
      sortByAttr(
        xslTag,
        xmlTagsGlobal,
        xmlTagsLocal
      )
    case "pxsl:sort-by-elem" =>
      sortByElem(
        xslTag,
        xmlTagsGlobal,
        xmlTagsLocal
      )
    case _ => List(Left(xslTag))
}

def extraerVariables(xslTags: List[Tag]): Map[String, String] = {
  xslTags.flatMap { xslTag =>
    xslTag.key match {
      case "pxsl:variable" =>
        xslTag.params
          .get("name")
          .flatMap(name =>
            xslTag.params.get("value").map(value => name -> value)
          )
      case _ =>
        extraerVariables(xslTag.childs.fold(_.toList, _ => Nil))
    }
  }.toMap
}

def extraerParametros(xslTags: List[Tag]): List[String] = {
  xslTags.flatMap { xslTag =>
    xslTag.key match {
      case "pxsl:param" =>
        xslTag.params.get("id").toList
      case _ =>
        extraerParametros(xslTag.childs.fold(_.toList, _ => Nil))
    }
  }
}

def procesarTags(
    xslTags: List[Tag],
    xmlTagsGlobal: List[Tag],
    xmlTagsLocal: List[Tag]
): List[Either[Tag, String]] = {
  val tagsProcesadas: List[Either[Tag, String]] =
    xslTags
      .map((xslTag) => procesarTagXsl(xslTag, xmlTagsGlobal, xmlTagsLocal))
      .flatMap(x => x)

  tagsProcesadas
}

def removerVariables(xslTags: List[Tag]): List[Tag] = {
  xslTags
    .filterNot(tag => tag.key == "pxsl:variable" || tag.key == "pxsl:param")
    .map(tag =>
      tag.copy(childs = tag.childs match
        case Left(childs) => Left(removerVariables(childs))
        case Right(texto) => Right(texto)
      )
    )
}

def reemplazarVariables(
    variables: Map[String, String],
    tags: List[Tag]
): List[Tag] = {
  tags.map { tag =>
    val nuevosParams = tag.params.map { param =>
      param match {
        case (key, value) =>
          val nuevoValor = variables.foldLeft(value)((acc, variable) =>
            variable match
              case (variableName, variableValue) =>
                acc.replace(s"$${${variableName}}", variableValue)
          )
          key -> nuevoValor
      }
    }

    val nuevosChilds = tag.childs match {
      case Left(childs) => Left(reemplazarVariables(variables, childs))
      case Right(texto) =>
        val nuevoTexto = variables.foldLeft(texto)((acc, variable) =>
          variable match
            case (variableName, variableValue) =>
              acc.replace(s"$${${variableName}}", variableValue)
        )
        Right(nuevoTexto)
    }

    tag.copy(params = nuevosParams, childs = nuevosChilds)
  }
}

def procesarVariables(
    xslTags: List[Tag],
    xmlTags: List[Tag],
    params: List[String]
): List[Either[Tag, String]] = {
  val variablesV = extraerVariables(xslTags)
  val parametros = extraerParametros(xslTags)

  params.length == parametros.length match
    case false =>
      println(
        s"Error: la cantidad de parametros no coincide con la cantidad de variables >> Parametros enviados: ${params
            .mkString(",")} Requeridos: ${parametros.mkString(",")}"
      )
      return List()
    case true =>
      val variablesP: Map[String, String] =
        parametros
          .map(nombre => (nombre -> params(parametros.indexOf(nombre))))
          .toMap

      val variables = variablesV ++ variablesP

      val xslTagsSinVariables = removerVariables(xslTags)

      val xslTagsActualizadas =
        reemplazarVariables(variables, xslTagsSinVariables)
      val xmlTagsActualizadas = reemplazarVariables(variables, xmlTags)

      procesarTags(
        xslTagsActualizadas,
        xmlTagsActualizadas,
        xmlTagsActualizadas
      )
}
