import scala.io.BufferedSource
import scala.io.Source

object Wc {
	// lines, words, chars
	def count(source: BufferedSource): Unit = {
		var numLines = 0
		var numWords = 0
		var numChars = 0
		for (line <- source.getLines) {
			numLines += 1
			val trimmedLine = line.trim
			if (!trimmedLine.isEmpty)
				numWords += line.trim.split("[ \t]+").length
			numChars += line.length + 1 // carriage return ...
		}
		println(numLines+" "+numWords+" "+numChars)
	}

	def main(args: Array[String]) {
		if (args.length == 0) {
			count(io.Source.stdin)
		}
		else {
			for (arg <- args) {
				try {
					count(Source.fromFile(arg))
				}
				catch {
					case e: java.io.FileNotFoundException => System.err.println(arg + ":" + e)
				}
			}
		}
	}
}
