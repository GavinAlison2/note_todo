import java.util.List;
import java.util.ArrayList;
import java.util.Arrays;

public class Kmp {

    public static void main(String[] args) {
        String haystack = "sadbutsad";
//        haystack = " leetcode";
//        haystack = "ABAABABABCA";
        String needle = "sad";
//        needle = "leeto";
//        needle = "ABABC";
        System.out.println(haystack);
        System.out.println(needle);
        System.out.println(strStr2(haystack, needle));
        System.out.println(strStr(haystack, needle));
    }

    public static int strStr2(String haystack, String needle) {
        int[] next = build_next(needle.split(""));
        int i = 0;
        int j = 0;
        while (i < haystack.length()) {
            if (haystack.charAt(i) == needle.charAt(j)) {
                i++;
                j++;
            } else if (j > 0) {
                j = next[j - 1];
            } else {
                i++;
            }
            if (j == needle.length()) {
                return i - j;
            }
        }
        return -1;
    }

    static int[] build_next(String[] patt) {
        /***
         * 计算 next 数组
         */
        int[] next = new int[patt.length];
        int j = 0;
        int i = 1;
        next[0] = 0;
        while (i < patt.length) {
            if (patt[j].equals(patt[i])) {
                // 相同，next[i] = j+1, i,j 都向右移动
                next[i++] = j + 1;
                j++;
            } else {
                // 当第一个不匹配，next[1]=0, i++
                if (j == 0) {
                    next[i] = 0;
                    i++;
                } else {
                    // 当不匹配，j 向左移
                    j = next[j - 1];
                }
            }
        }
        return next;
    }

    // 自己写的，思考的地方，构建next数组，找到最大的公共前后缀， 根据next数组，移动去匹配，返回 i - m + 1，注意最后的长度
    public static int strStr(String originStr, String patterStr) {

        int n = originStr.length();
        int m = patterStr.length();
        int[] next = new int[m];
        // build next[]
        for (int j = 0, i = 1; i < m; i++) {
            // 不相同
            while (j > 0 && patterStr.charAt(j) != patterStr.charAt(i)) j = next[j - 1];
            if (patterStr.charAt(j) == patterStr.charAt(i)) j++;
            next[i] = j;
        }
        System.out.println(Arrays.toString(next));
        // 匹配
        for (int i = 0, j = 0; i < n; i++) {
            while (j > 0 && originStr.charAt(i) != patterStr.charAt(j)) j = next[j - 1];
            if (originStr.charAt(i) == patterStr.charAt(j)) j++;
            if (j == m) {
                return i - m + 1;
            }
        }

        return -1;
    }
}