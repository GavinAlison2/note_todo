import java.util.Map;
import java.util.TreeMap;
import java.util.Iterator;
import java.util.Arrays;
import java.util.HashMap;

public class Demo {
    public static void main(String[] args) {
        int[] nums = new int[] { 1, 1, 2 };
        TreeMap map = new TreeMap();
        for (int i = 0; i < nums.length; i++) {
            map.put(nums[i], 1);
        }
        Iterator ite = map.keySet().iterator();
        int[] newNums = new int[map.size()];
        int j = 0;
        while (ite.hasNext()) {
            newNums[j++] = (int) ite.next();
        }
        nums = newNums;
        System.out.println(Arrays.toString(nums));
    }
}