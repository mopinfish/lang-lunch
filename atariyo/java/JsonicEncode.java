import net.arnx.jsonic.JSON;
import java.util.Arrays;
import java.util.HashMap;

public class JsonicEncode {
    public static void main(String[] args) {
        Task task = new Task();
        task.setId(1);
        task.setName("tasukujp");
        String json = JSON.encode(task);
        System.out.println(json);
    }
}
