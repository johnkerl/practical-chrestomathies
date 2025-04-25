use std::env;
use std::fs;
use std::io;

fn main() -> io::Result<()> {
    let args: Vec<String> = env::args().collect();

    if args.len() > 1 {
        for (_, arg) in args.iter().skip(1).enumerate() {
            let file = fs::File::open(&arg)?;
            let reader = io::BufReader::new(file);
            process_bytes(reader)?
        }
    } else {
        let reader = io::BufReader::new(io::stdin());
        process_bytes(reader)?
    };

    Ok(())
}

fn process_bytes<R: io::Read>(mut reader: R) -> io::Result<()> {
    const BUFFER_SIZE: usize = 16;
    const BYTES_PER_CLUMP: usize = 4;
    let mut buffer = [0u8; BUFFER_SIZE];
    let mut offset = 0;

    loop {
        match reader.read(&mut buffer)? {
            0 => break,
            num_bytes_read => {
                // Print the offset before the hex payload
                print!("{offset:08x}:");

                // Print the hex payload
                for i in 0..BUFFER_SIZE {
                    if i < num_bytes_read {
                        let v = buffer[i];
                        print!(" {v:02x}");
                    } else {
                        print!("   ")
                    }
                    if (i % BYTES_PER_CLUMP) == (BYTES_PER_CLUMP - 1) {
                        if (i > 0) && (i < BUFFER_SIZE - 1) {
                            print!(" ")
                        }
                    }
                }

                // Print the ASCII dump after the hex payload
                print!(" |");

                for i in 0..BUFFER_SIZE {
                    if i < num_bytes_read {
                        let v = buffer[i];
                        // Don't use ASCII is-print since we don't want to print
                        // tabs, carriage returns, etc as such.
                        if buffer[i] >= 0x20 && buffer[i] <= 0x7e {
                            let c = v as char;
                            print!("{c}");
                        } else {
                            print!(".");
                        }
                    } else {
                        print!(" ")
                    }
                }
                println!("|");

                if num_bytes_read < 16 {
                    break;
                }
                offset += num_bytes_read;
            }
        }
    }

    Ok(())
}
