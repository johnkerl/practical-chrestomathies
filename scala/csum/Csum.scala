// ================================================================
// Copyright (c) 2004, 2013 John Kerl.
// kerl at math dot arizona dot edu
//
// This code and information is provided as is without warranty of
// any kind, either expressed or implied, including but not limited to
// the implied warranties of merchantability and/or fitness for a
// particular purpose.
//
// No restrictions are placed on copy or reuse of this code, as long
// as these paragraphs are included in the code.
// ================================================================

// xxx scala.io.Source can't handle 8-bit binary data.
import java.io.BufferedInputStream
import java.io.FileInputStream
import java.io.InputStream

object Csum {

	// ----------------------------------------------------------------
	def usage() {
		System.err.println("Csum [options] {filenames ...}")
		System.err.println("If no filenames are given, stdin is read.")
		System.err.println("Options:")
		System.err.println("--eth:     Use Ethernet checksum algorithm (default)")
		System.err.println("--simple:  Use simple bytecount-bytesum checksum algorithm")
		System.err.println("--spin:    print running checksums to screen (default off)")
		System.err.println("--nospin:  do not print running checksums to screen")
		System.exit(1)
	}

	// ----------------------------------------------------------------
	// xxx scopt/scopt?
	def parseCommandLine(argv: Array[String])
		: Tuple2[Map[String,String],List[String]] =
	{
		var opts = Map("spin" -> "false", "algo" -> "eth")
		var args = List[String]()

		for (arg <- argv) {
			// xxx case?
			if (arg.startsWith("-")) {
				if (arg == "--spin")
					opts += "spin" -> "true"
				else if (arg == "--nospin")
					opts += "spin" -> "false"
				else if (arg == "--ip")
					opts += "algo" -> "ip"
				else if (arg == "--eth")
					opts += "algo" -> "eth"
				else if (arg == "--simple")
					opts += "algo" -> "simple"
				else
					usage()
			}
			else {
				args ::= arg
			}
		}

		(opts, args)
	}

	// ----------------------------------------------------------------
	def main(argv: Array[String]) {
		val (opts, args) = parseCommandLine(argv)

		if (args.length == 0)
			doOneArgument("-", opts)
		else
			args.foreach(arg => doOneArgument(arg, opts))
	}

	// ----------------------------------------------------------------
	def doOneArgument(fileName: String, opts: Map[String,String]) {
		val bufferSize = 1024

		val inputStream = if (fileName == "-") System.in else new FileInputStream(fileName)
		val bis     = new BufferedInputStream(inputStream)
		val buffer  = new Array[Byte](bufferSize)
		val doSpin  = opts("spin") == "true"
		var done    = false
		var nblocks = 0
		var nbytes  = 0

		val checksummer = ChecksummerFactory(opts("algo"))

		checksummer.start

		if (doSpin) {
			print("...")
			System.out.flush
		}

		while (!done) {
			val numBytesRead = bis.read(buffer, 0, bufferSize)
			if (numBytesRead < 0) {
				done = true
			}
			else {
				nblocks += 1
				nbytes += numBytesRead
				checksummer.accumulate(buffer, numBytesRead)

				if (doSpin && ((nblocks & 0xfff) == 0)) {
					print("\r"+checksummer.getStringState)
					System.out.flush
				}
			}
		}
		checksummer.finish

		val stringSum = checksummer.getStringSum
		if (doSpin)
			print("\r")
		println(stringSum + "  " + nbytes + "  " + fileName)
	}
}
