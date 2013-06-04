import scala.io.Source

object TimeLines {
	def main(args: Array[String]) {
		var curr = System.nanoTime
		var prev = curr
		for (line <- io.Source.stdin.getLines) {
			prev = curr
			curr = System.nanoTime
			println("%13.9f ".format((curr-prev)*1.0e-9) + line)
		}
	}
}
