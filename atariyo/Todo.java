import java.io.File;
import java.io.FileWriter;
import java.io.BufferedWriter;
import java.io.PrintWriter;
import java.io.IOException;

public class Todo{

    String str;
    int num;

    // constractor
    Todo() {
        this.str = "mozi";
        this.num = 123;
    }

    public static void main(String[] args){
        System.out.println("Hello, world.");
        addList();
    }

    // todo一覧に追加する
    public static void addList () {
        // とりあえずファイルに書き込むサンプル
        // ここでtodoのjsonを書き込む
        try{
          File file = new File("todolist.txt");
          FileWriter filewriter = new FileWriter(file);

          filewriter.write("atariyo");
          filewriter.close();
        }catch(IOException e){
          System.out.println(e);
        }
    }

    // todo一覧を表示する
    public static void showList () {
    
    }
}
