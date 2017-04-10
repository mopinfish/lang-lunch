import java.io.File;
import java.io.FileWriter;
import java.io.FileReader;
import java.io.BufferedReader;
import java.io.BufferedWriter;
import java.io.FileNotFoundException;
import java.io.PrintWriter;
import java.io.IOException;

public class Todo{

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
          delListItem(args[1]);

          System.out.println("---- result ----");
          showList();
        }
    }

    // todo一覧に追加する
    public static void addList (String todo) {
        try{
          File file = new File("todolist.txt");
          FileWriter filewriter = new FileWriter(file, true);

          filewriter.write(todo + "\n");
          filewriter.close();
        }catch(IOException e){
          System.out.println(e);
        }
    }

    // todo一覧から削除する
    public static void delListItem (String itemName) {
      StringBuffer fileRead = new StringBuffer("");

      try{
        File file = new File("todolist.txt");
        BufferedReader br = new BufferedReader(new FileReader(file));

        String str;
        while((str = br.readLine()) != null){
          if (str.equals(itemName)) {
            System.out.println(str + " is deleted");
          }
          else {
            fileRead.append(str + "\r\n");
          }
        }

        br.close();

        FileWriter filewriter = new FileWriter(file);
        filewriter.write(fileRead.toString());
        filewriter.close();
      }catch(FileNotFoundException e){
        System.out.println(e);
      }catch(IOException e){
        System.out.println(e);
      }
    }

    // todo一覧を表示する
    public static void showList () {
      try {
        File file = new File("todolist.txt");
        FileReader filereader = new FileReader(file);

        int ch;
        while ((ch = filereader.read()) != -1) {
          System.out.print((char)ch);
        }

        filereader.close();
      } catch(FileNotFoundException e) {
        System.out.println(e);
      } catch (IOException e) {
        System.out.println(e);
      }
    }
}
