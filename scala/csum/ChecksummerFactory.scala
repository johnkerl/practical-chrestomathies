object ChecksummerFactory {
	def apply (name: String) : Checksummer = {
		name match {
			case "eth"    => new EthSummer
			case "simple" => new SimpleSummer
			case _        => throw new IllegalArgumentException("Unrecognized name \""+name+"\"")
		}
	}
}
