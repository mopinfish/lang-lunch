#!rust run

/******************************************
 * Generics
 */
fn pair<T, S>(t: T, s: S) -> (T, S) {
    (t, s)
}

fn checkTypeCall() {
    // T = i32, S = f64で呼び出す
    let i = pair(1, 1.0);
    println!("{:?}", i);
    // 型を明示する方法もある
    let i = pair::<isize, f64>(2, 2.0);
    println!("{:?}", i);
    // T = &str, S = Stringで呼び出す
    let s = pair::<&str, String>("str", "string".to_string());
    println!("{:?}", s);
}

/******************************************
 * Structure
 */
// struct 名前; (Unit構造体の構文)
struct Dummy;

// struct 名前(型, ..); (タプル構造体の構文)
struct Point(f64, f64);

// struct 名前 {フィールド: 型, ..} (通常の構造体の構文)
struct Color {
    r: u8,
    g: u8,
    // 最後のフィールドの末尾にもカンマをつけられる
    b: u8,
}

fn checkStructure() {
    // Unit構造体は名前でそのまま初期化
    let dummy = Dummy;
    // タプル構造体は関数のように初期化
    // 実際、関数として扱うこともできる
    let point = Point(10.0, 20.0);
    // タプル構造体のフィールドへのアクセス
    let x = point.0;
    println!("{:?}", x);
    // 普通の構造体の初期化
    let black = Color { r: 0, g: 0, b: 0 };
    // 普通の構造体のフィールドへのアクセス
    let r = black.r;
    let g = black.g;
    let b = black.b;
    println!("{:?}", r);
    println!("{:?}", g);
    println!("{:?}", b);
}

/******************************************
 * Implementation
 */
struct Celsius(f64);
struct Kelvin(f64);

// `impl 型名 {..}`で型に対する実装を書ける
impl Celsius {
    // `{..}`の中には関数が書ける。
    // 第一引数が`self`、`&mut self` `&self`, `Box<self>`の場合はメソッドとなる
    fn to_kelvin(self) -> Kelvin {
        // selfを通じてフィールドにアクセスできる。
        Kelvin(self.0 + 273.15)
    }

    // 第一引数が`self`系でない場合は関連関数となる
    fn from_kelvin(k: Kelvin) -> Self {
        Celsius(k.0 - 273.15)
    }
}

fn checkImplementation() {
    let absolute_zero = Kelvin(0.0);
    let triple_point = Celsius(0.0);
    // 関連関数は`型名::関数名(引数)`で呼び出す。
    let celsius = Celsius::from_kelvin(absolute_zero);
    // メソッドは`値.関数名(引数)`で呼び出す。
    let kelvin = triple_point.to_kelvin();
}

fn main() {
    // Generics
    checkTypeCall();
    // Structure
    checkStructure();
    // Implementation
	checkImplementation();
}
