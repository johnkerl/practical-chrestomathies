import scala.io.BufferedSource
import scala.io.Source

object Linesort {
	def handle(source: BufferedSource) {
		// xxx is this streaming?
		for (line <- source.getLines) {
			println(line.split("[ \t]+").sorted.mkString(" "))
		}
	}

	def main(args: Array[String]) {
		if (args.length == 0) {
			handle(io.Source.stdin)
		}
		else {
			// xxx exc ...
	    	for (arg <- args) {
				try {
					handle(Source.fromFile(arg))
				}
				catch {
					case e: java.io.FileNotFoundException => System.err.println(arg + ":" + e)
				}
			}
		}
	}
}
