class SimpleSummer extends Checksummer {
	var byteCount = 0L
	var byteSum   = 0L

	def start() : Unit = {
		byteCount = 0L
		byteSum   = 0L
	}

	def accumulate(chs: Array[Byte], len: Int) : Unit = {
		byteCount += len
		for (i <- 0 until len)
			byteSum += chs(i).toLong & 0xff
	}

	def finish() : Unit = {
		// Nothing for this summer
	}

	def getStringState() : String = {
		"%016x".format(byteCount) + "_" + "%016x".format(byteSum)
	}

	def getStringSum() : String = {
		"%016x".format(byteCount) + "_" + "%016x".format(byteSum)
	}
}
