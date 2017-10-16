#!rust run

use std::env;

fn print_altenately_concat_text(text1: &str, text2: &str) {
	let mut concat_text = String::new();
    for (str1, str2) in text1.chars().zip(text2.chars()) {
        concat_text.push(str1);
        concat_text.push(str2);
    }
    println!("{}", concat_text);
}

fn main() {
	let args: Vec<String> = env::args().collect();
	match (args.get(1), args.get(2)) {
		(Some(text1), Some(text2)) => print_altenately_concat_text(text1, text2),
		(_, None) => println!("pass two argument"),
		(None, _) => println!("pass two argument"),
	}
}
