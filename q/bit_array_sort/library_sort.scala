object Main {
  def main(args: Array[String]): Unit = {
    val file = new java.io.PrintWriter("out/output_library_sort")
    var source = scala.io.Source.fromFile("out/testdata")
    source.getLines.map(_.toInt).toSeq.sorted.foreach(x => file.println(f"${x}%07d"))
    source.close()
    file.close()
  }
}
