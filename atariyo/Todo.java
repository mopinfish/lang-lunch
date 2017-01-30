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
        if (args[0].equals("show")) {
            showList();
        }
        else if (args.length < 2) {
            System.out.println("args length is less 3");
            return;
        }

        if (args[0].equals("add")) {
            addList(args[1]);
        }

        // addの後に実装
        if (args[0].equals("del")) {
            addList(args[1]);
        }
    }

    // todo一覧に追加する
    public static void addList (String todo) {
        // とりあえずファイルに書き込むサンプル
        // ここでtodoのjsonを書き込む
        try{
          File file = new File("todolist.txt");
          FileWriter filewriter = new FileWriter(file, true);

          filewriter.write(todo + "\n");
          filewriter.close();
        }catch(IOException e){
          System.out.println(e);
        }
    }

    // todo一覧を表示する
    public static void showList () {
        System.out.println("show func");
    }
}
