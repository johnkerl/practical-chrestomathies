trait Checksummer {
	def start() : Unit
	def accumulate(chs: Array[Byte], len: Int) : Unit
	def finish() : Unit
	def getStringState() : String
	def getStringSum() : String
}
