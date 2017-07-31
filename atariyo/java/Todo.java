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
          addItem(args[1]);
          System.out.println("---- result ----");
          showList();
        }

        // addの後に実装
        if (args[0].equals("del")) {
          delItem(args[1]);

          System.out.println("---- result ----");
          showList();
        }
    }

    // todo一覧に追加する
    public static void addItem (String todo) {
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
    public static void delItem (String itemName) {

      try{
        File file = new File("todolist.txt");
        BufferedReader br = new BufferedReader(new FileReader(file));

        StringBuffer fileRead = new StringBuffer("");
        String str;
        Boolean isHit = false;

        while((str = br.readLine()) != null){
          if (str.equals(itemName)) {
            System.out.println("delete " + str);
            isHit = true;
          }
          else {
            fileRead.append(str + "\r\n");
          }
        }

        br.close();

        if (!isHit) {
          System.out.println(itemName + " not exist in list");
          return;
        }

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
