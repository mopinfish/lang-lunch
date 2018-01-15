struct Apple {
    size: u64
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
  let apple = Apple { size: 1 };
  println!("{}", factorial());
  println!("{}", apple.size);

}
