// rustc echo.rs
// ./echo

fn main() {
	let args: ~[~str] = os::args();
	// xxx needs for-loop per se
	// xxx or, needs string join ...
	// gah!!! :(
	for args.each |arg| {
	//for each(args) |arg| {
		io::print(" ");
		io::print(arg);
	}
	io::println("");
}

//fn main() {
//	let args: ~[~str] = os::args();
//	let mut i = 1;
//	let n = args.len();
//	// xxx needs for-loop per se
//	// xxx or, needs string join ...
//	while i < n {
//		if i > 1 {
//			io::print(" ");
//		}
//		io::print(args[i]);
//		i += 1;
//	}
//	io::println("");
//}
