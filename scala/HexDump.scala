// xxx scala.io.Source can't handle 8-bit binary data.
import java.io.BufferedInputStream
import java.io.FileInputStream
import java.io.InputStream

object HexDump {

	// ----------------------------------------------------------------
	def hexdump(inputStream : InputStream) = {
		val bytesPerClump = 4
		val clumpsPerLine = 4
		val bufferSize    = bytesPerClump * clumpsPerLine

		val bis = new BufferedInputStream(inputStream)
		var buf = new Array[Byte](bufferSize)
		var done = false
		var offset : Integer = 0

		while (!done) {
			val numBytesRead = bis.read(buf, 0, bufferSize)
			if (numBytesRead < 0) {
				done = true
			}
			else {
				print("%08x: ".format(offset))

				for (i <- 0 until bufferSize) {
					if (i < numBytesRead)
						print("%02x ".format(buf(i).toInt & 0xff))
					else
						printf("   ")

					if ((i % bytesPerClump) == (bytesPerClump - 1)) {
						if ((i > 0) && (i < bufferSize-1))
							printf(" ");
					}
				}

				System.out.printf("|");
				for (i <- 0 until numBytesRead) {
					if ((buf(i) >= 0x20) && (buf(i) <= 0x7e))
						System.out.print(buf(i).toChar)
					else
						System.out.print(".")
				}
				for (i <- numBytesRead until bufferSize) {
					System.out.print(" ")
				}
				System.out.printf("|\n");

				offset += bufferSize
			}
		}
	}

	// ----------------------------------------------------------------
	def main(args: Array[String]) {
		if (args.length == 0) {
			hexdump(System.in)
		}
		else {
	    	for (arg <- args) {
				try {
					hexdump(new FileInputStream(arg))
				}
				catch {
					case e: java.io.FileNotFoundException =>
						System.err.println(arg + ":" + e)
				}
			}
		}
	}
}
