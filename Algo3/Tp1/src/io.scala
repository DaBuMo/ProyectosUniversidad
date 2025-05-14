import scala.util.{Try, Failure, Success}
import java.io.PrintWriter

def leerArchivo(path: String): String = {
  Try(scala.io.Source.fromFile(path).mkString) match {
    case Success(content) if content.nonEmpty => content
    case Success(_) =>
      println("El archivo está vacío.")
      ""
    case Failure(exception) =>
      println(s"Error al leer el archivo: ${exception.getMessage}")
      ""
  }
}

def escribirArchivo(path: String, contenido: String): Unit = {
  val writer = new PrintWriter(path)
  writer.write(contenido)
  writer.close()
}
