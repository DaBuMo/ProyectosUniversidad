def findByKey(tag: Tag, key: String): Option[Tag] = {
  tag.childs match
    case Left(childTags) => childTags.find((x) => x.key == key)
    case _               => None
}

def findByPath(tag: Tag, path: List[String]): Option[Tag] = {
  path match
    case Nil => Some(tag)
    case key :: rest =>
      findByKey(tag, key) match
        case Some(childTag) => findByPath(childTag, rest)
        case None           => None
}

def printTag(tag: Tag, i: Int = 0): Unit = {
  println(
    s"${"--".repeat(i)}Tag(key = ${tag.key}"
  )
  println(
    s"${"--".repeat(i)}Tag(params = ${tag.params}"
  )

  tag.childs match
    case Left(childTags) =>
      childTags.foreach(tag => printTag(tag, i + 1))
    case Right(content) =>
      println(s"${"--".repeat(i)}Content = \"$content\"")

}

//Esta funcion separa el path en substrings
def separatePath(path: String): List[String] = {
  path match {
    case "/" =>
      List(
        "/"
      ) /*Si el path tiene solo esto, dejamos este simbolo porque asi identificamos
        que es el path a root
       */

    // En este caso si toca el palito al principio, lo queremos conservar, y como los replace
    // lo eliminan, tenemos agregarlo al principio de esta forma
    case s if s.startsWith("/") =>
      "/" :: s
        .replace("/", " ")
        .replace("[", " ")
        .replace("]", "")
        .trim
        .split(" ")
        .toList

    case _ =>
      path
        .replace(
          "/",
          " "
        ) // para el caso general queremo separar todas las palabras
        .replace("[", " ")
        .replace("]", "")
        .trim
        .split(" ")
        .toList
  }
}

def searchByPathTag(tag: Tag, path: List[String], root: Tag): List[Tag] =
  path match {

    case "/" :: resto =>
      searchByPathTag(
        root,
        resto,
        root
      ) // Si toca este caso tenemos que empezar la busqueda en el tag root

    case Nil => List(tag) // Cuando es la lista vacia devolvemos el tag actual

    case cabeza :: resto
        if cabeza.contains(
          "="
        ) => // si este elemento contiene una condicion filtramos los hijos actuales
      val Array(clave, valor) =
        cabeza.split("=") // FALTAN SACAR LAS COMMILAS AL VALOR!!!!!!!!!
      tag match {
        case _ if tag.params.get(clave).contains(valor.replace("\"", "")) =>
          searchByPathTag(tag, resto, root)
        case _ => Nil
      }

    case "." :: resto => /*Como el punto significa el tag actual, tenemos que buscar en el actual el
  resto del path*/
      searchByPathTag(tag, resto, root)

    case cabeza :: resto => // caso general, cabeza/resto_del_path filtramos los hijos tal que el nombre
      // concuerden con la cabeza
      tag.childs match
        case Left(hijos) =>
          hijos
            .filter(_.key == cabeza)
            .flatMap(hijo => searchByPathTag(hijo, resto, root))
        case Right(_) => Nil

  }

def searchByPath(
    path: List[String],
    globales: List[Tag],
    locales: List[Tag]
): List[Tag] = path match {
  // armo un tag con root para reutilizar la funcion searchByPathTag
  case "/" :: resto =>
    val tagRoot = Tag(
      "root",
      "xlm",
      Map(),
      Left(
        globales
      )
    )
    searchByPathTag(tagRoot, path, tagRoot)

  case _ =>
    val englobador = Tag(
      "englobador",
      "xlm",
      Map(),
      Left(
        locales
      )
    )
    searchByPathTag(englobador, path, englobador)
}
