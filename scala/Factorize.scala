import scala.io.BufferedSource
import scala.io.Source

object Factorize {
	val PROGRAM_NAME = "Factorize";

	def factorize(i: Long): Map[Long,Int] = {
		var factorInfo : Map[Long,Int] = Map.empty
		var d = 2L
		var n = i
		while (n > 1L) {
			var e = 0
			while ((n % d) == 0L) {
				e += 1
				n /= d
			}
			if (e > 0)
				factorInfo += d -> e
			d += 1
		}
		factorInfo
	}

	def usage(programName: String) {
		System.err.println("Usage: "+programName+" {one or more longs}")
		System.exit(1)
	}

	def main(args: Array[String]) {
		if (args.length < 1)
			usage(PROGRAM_NAME);

    	for (arg <- args) {
			val n = arg.toLong
			val f = factorize(n)
			print(n+":")
			for (d <- f.keys.toList.sorted) {
				val e = f(d)
				print(" ")
				if (e > 1)
					print(d + "^" + e)
				else
					print(d)
			}
			println()
		}
	}
}
