object Millis2Gmt {
	def usage(programName: String) {
		System.err.println(String.format("Usage: %s {one or more millis}",
			programName))
		System.exit(1)
	}

	def main(args: Array[String]) {
		if (args.length < 1)
			usage("Millis2Gmt")
		val formatter = new java.text.SimpleDateFormat("yyyyMMdd-HH:mm:ss.SSS")
		for (arg <- args) {
			val millis = arg.toLong
			val date = new java.util.Date(millis)
			println(formatter.format(date))
		}
	}
}
