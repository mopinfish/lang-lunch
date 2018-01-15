// $ rustc --version
// rustc 1.22.1 (05e2e1c41 2017-11-22)

// 使用する crate を宣言
extern crate reqwest; // シンプルなHTTPクライアント
extern crate serde;  // シリアライズライブラリ
extern crate serde_json; // serdeでJSONを扱うライブラリ
extern crate url;

// `#[derive(Serialize, Deserialize)`を使えるようにする
#[macro_use]
extern crate serde_derive;

// HTTPライブラリのデファクトスタンダード
// header! を使うだけ
#[macro_use]
extern crate hyper;

// reqwest::Url と毎回書くのは大変なので Url にする
use reqwest::Url;
use hyper::header::Headers;

// HTTPヘッダー用の構造体を生成してくれる
header! { (XChatWorkToken, "X-ChatWorkToken") => [String] }

/// HTTPリクエスト用とのマッピング用の構造体
// POSTパラメータにあわせて、構造体定義しておけば勝手にいい感じにしてくれる。便利
#[derive(Serialize)]
struct PostMessageRequest {
    body: String, // Bodyパラメータを設定する
}

/// メッセージ投稿APIのレスポンスとのマッピング用の構造体
// 帰ってくるJSONにあわせて構造体定義しておけば勝手にいい感じにしてくれる。便利
// #[serde(untagged)] でどのようにマッピングするか指定します
// #[derive](Debug)]を書いておくとDebug出力を自動生成してくれます
#[derive(Deserialize, Debug)]
#[serde(untagged)]
enum PostMessageResponse {
    Error { errors: Vec<String> },
    MessageId { message_id: String },
}

/// PostMessageResponseのままだと使いにくいので用意
#[derive(Debug)]
struct MessageId {
    message_id: String,
}

/// post_message関数で発生するエラーを一つの型にするためのenum
// 型を合わせる必要があるため作成、文字列にしてしまう手もある
#[derive(Debug)]
enum PostMessageError {
    Reqwest(reqwest::Error),
    UrlParse(url::ParseError),
    API(Vec<String>),
}

/// post_message関数でreqwest::Errorを返す関数を呼ぶときに勝手に変換できるようにする
// PostMessageErrorにFromトレイトを実装している
impl From<reqwest::Error> for PostMessageError {
    fn from(e: reqwest::Error) -> PostMessageError {
        PostMessageError::Reqwest(e)
    }
}

/// post_message関数でurl::ParserErrorを返す関数を呼ぶときに勝手に変換できるようにする
// PostMessageErrorにFromトレイトを実装している
impl From<url::ParseError> for PostMessageError {
    fn from(e: url::ParseError) -> PostMessageError {
        PostMessageError::UrlParse(e)
    }
}

/// みんなだいすきエントリーポイント
fn main() {
    // unwrap すると Result<A,B>な型のとき Aがかえってくる Bの値をもってるときはpanicがおきる
    // ResultはいわゆるEither型
    // `left` `right`ではなく `Ok` `Err`
    // 自分が使うツールぐらいだったら Resultな型はmain関数でunwrapしています
    // unwarpはサンプルコードでよくみかけます
    let (room_id, body) = parse_args().unwrap();
    let token = env_chatwork_token().unwrap();
    // &tokenで渡せるように関数をつくらないと tokenはここで使えくなってしまう(後述)
    let response = post_message(&token, room_id, &body).unwrap();
    // {:?} を使うとデバッグ形式で出力できます
    println!("{:?}", response);
}

/// 環境変数 CHATWORK_TOKENから値を取り出す
fn env_chatwork_token() -> Result<std::string::String, String> {
    std::env::var("CHATWORK_TOKEN")
        // そのままのだとエラーの原因がよくわからないエラーメッセージを作成
        // 文字列は&strなので Stringに変換
        // &'static str のままでも値をかえせますが、今回のコードは Stringで統一しています
        // to_stringするのにはメモリアロケーションが発生するので、必要がないなら避けるべきかもしれません
        // エラーメッセージを動的に生成してしまうと、&'static strで返すことができないので、Stringに統一しています
        .map_err(|_| "CHATWORK_API_TOKEN environment variable not present".to_string())
}

/// コマンドライン引数を解析する
// u32は unsigned 32bit 整数
fn parse_args() -> Result<(u32, String), String> {
    // コマンドライン引数の取得
    let mut args = std::env::args();
    args.next(); // プログラムの名前なので無視します
    // 最初のコマンドライン引数を取得
    // Optionが返ってくるのでパターンマッチで分岐
    let room_id = match args.next() {
        Some(s) => s.parse::<u32>()
            // or で失敗したときの値を作成
            .or(Err("arg1 expected number for room_id".to_string())),
        // 最初の引数が取得できなかった場合の値を作成
        None => Err("arg1 expected room_id, found None".to_string()),
    // `?`を利用するとResult型の失敗している値の場合は、そのまま`return`
    // 成功している場合はResultの中から値を取り出せる
    // room_idはu32として利用できる
    }?;

    let body = match args.next() {
        Some(s) => s,
        // 二番目の引数を取得できなかったときの値を作成
        // s はResultではないので、Resultのままにすることはできない
        // `?`を使用してもよいけど `return` するのは明白なので、そのまま`return`しています
        None => return Err("args2 expected body, found None".to_string()),
    };
    // Resultを返さないといけないのでOkで包む
    // Rustでは最後の式が戻り値に
    // セミコロンを付けると () 型になってしまうので書かない
    Ok((room_id, body))
}

/// POSTするURLを作成する
fn post_message_url(room_id: u32) -> Result<Url, url::ParseError> {
    let url_str = format!("https://api.chatwork.com/v2/rooms/{}/messages", room_id);
    Url::parse(&url_str) // 文字列をURLに変換するのは失敗することがある
}

/// アクセストークンをセットしたHTTPヘッダーを作成する
// Stringでなくて &strにしないと関数の引数に使った変数の所有権が移動してしまって使えなくなってしまう
// tokenは何度が使いまわしたいと想像がつくので、 &str にして貸すだけにしてあげてます
// (結局to_stringメソッドでメモリアロケーションが発生していますのであんまり意味はないです)
fn chatwork_api_headers(token: &str) -> Headers {
    // headers.setは () を返すので、ワンラインではかけず…
    // setを使うので mutに
    let mut headers = Headers::new();
    headers.set(XChatWorkToken(token.to_string()));
    headers
}

/// HTTPリクエストをしてREST APIを実行してJSONに
/// Tに使える型 JSONに使える型を制限をかけているだけ
// UrlやHeaderは使いまわしたいかもしれませんが、利用しているライブラリの都合所有権を移動させてしまいます
fn request_chatwork_api<T: serde::Serialize, JSON: serde::de::DeserializeOwned>
    (url: Url,
     headers: Headers,
     body: &T)
     -> Result<JSON, reqwest::Error> {
    reqwest::Client::new()
        .post(url)
        .form(body)
        .headers(headers)
        .send()? // HTTPリクエスト (Resultが返ってくる)
        .json() // JSONに変換
}

/// request_chatwork_api をラップして使いやすく
// u32はコピーされるので関数に渡しても、その後も使いまわせます(Copyトレイトが実装されているため)
// 型の不一致がおきてしまうので、まとめてあつかえるPostMessageErrorを用意
// 静的ディスパッチでなくなってもよいなら Box<std::error::Error>を使う手もたぶんある
fn post_message(token: &str, room_id: u32, body: &str) -> Result<MessageId, PostMessageError> {
    let body = PostMessageRequest { body: body.to_owned() };
    // Err は url::ParseError ですが Fromトレイトを実装しているので、PostMessageErrorに変換してくれます
    let url = post_message_url(room_id)?;
    let headers = chatwork_api_headers(token);
    let response = request_chatwork_api(url, headers, &body)?;
    // 使いやすいように値を変換して返す
    match response {
        PostMessageResponse::Error { errors } => Err(PostMessageError::API(errors)),
        PostMessageResponse::MessageId { message_id } => Ok(MessageId { message_id: message_id }),
    } // ここで`return`しているので、`?`は使う必要はない
}
