package csums
import("errors")

var ChecksummerFactoryAlgoNames = []string {
	"simple", "eth", "adler32", "crc32", "crc64", "fnv32", "fnv32a", "fnv64", "fnv64a",
}

// needs an idiomatic go-style case statement -- ?
func ChecksummerFactory(name string) (summer Checksummer, err error) {
    if name == "eth" {
		return new(EthSummer), nil
	}
	if name == "simple" {
		return new(SimpleSummer), nil
	}
	if name == "adler32" {
		//return new(Adler32Summer), nil
		//return Adler32Summer.New(), nil
		return NewAdler32Summer(), nil
	}
	if name == "crc32" {
		return NewCRC32Summer(), nil
	}
	if name == "crc64" {
		return NewCRC64Summer(), nil
	}
	if name == "fnv32" {
		return NewFNVSummer(name), nil
	}
	if name == "fnv32a" {
		return NewFNVSummer(name), nil
	}
	if name == "fnv64" {
		return NewFNVSummer(name), nil
	}
	if name == "fnv64a" {
		return NewFNVSummer(name), nil
	}
	return nil, errors.New("ChecksummerFactory: unrecognized name \""+name+"\"")
}
