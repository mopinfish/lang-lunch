use std::io;
use std::io::prelude::*;

trait Fruit {
    fn get_size(&self) -> u64;
}
struct Apple {
    size: u64
}
struct Pine {
    size: u64
}

impl Fruit for Apple {
    fn get_size(&self) -> u64 {
        println!("apple");
        self.size
    }
}

impl Fruit for Pine {
    fn get_size(&self) -> u64 {
        println!("pine");
        self.size
    }
}

fn fib(i: i32) -> i32 {
    let mut sum: i32 = 0;
    if i <= 1 {
        i
    } else {
        fib(i - 1) + fib(i - 2)
    }
}
fn some_calc(i: i32) -> i32 {
    if i == 0 {
        i
    } else {
        i + 10
    }
}
fn factorial() -> i32 {
    let mut sum = 0;
    for i in 0..10 {
        println!("{}", i);
        sum += i;
    }
    sum
}
fn main() {
    println!("{}", fib(10));
//  let apple = Apple { size: 3 };
//  println!("{}", factorial());
//  println!("{}", apple.size);
//  println!("{}", apple.get_size());
}
