import scala.io.Source

object StampLines {
	def main(args: Array[String]) {
		val formatter = new java.text.SimpleDateFormat("yyyyMMdd-HH:mm:ss.SSS")
		io.Source.stdin.getLines.foreach(line =>
			println("["+formatter.format(new java.util.Date()) + "] " + line)
		)
	}
}
