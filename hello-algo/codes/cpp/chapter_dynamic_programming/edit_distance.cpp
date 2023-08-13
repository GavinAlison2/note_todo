/**
 * File: edit_distance.cpp
 * Created Time: 2023-07-13
 * Author: Krahets (krahets@163.com)
 */

#include "../utils/common.hpp"

/* 编辑距离：暴力搜索 */
int editDistanceDFS(string s, string t, int i, int j) {
    // 若 s 和 t 都为空，则返回 0
    if (i == 0 && j == 0)
        return 0;
    // 若 s 为空，则返回 t 长度
    if (i == 0)
        return j;
    // 若 t 为空，则返回 s 长度
    if (j == 0)
        return i;
    // 若两字符相等，则直接跳过此两字符
    if (s[i - 1] == t[j - 1])
        return editDistanceDFS(s, t, i - 1, j - 1);
    // 最少编辑步数 = 插入、删除、替换这三种操作的最少编辑步数 + 1
    int insert = editDistanceDFS(s, t, i, j - 1);
    int del = editDistanceDFS(s, t, i - 1, j);
    int replace = editDistanceDFS(s, t, i - 1, j - 1);
    // 返回最少编辑步数
    return min(min(insert, del), replace) + 1;
}

/* 编辑距离：记忆化搜索 */
int editDistanceDFSMem(string s, string t, vector<vector<int>> &mem, int i, int j) {
    // 若 s 和 t 都为空，则返回 0
    if (i == 0 && j == 0)
        return 0;
    // 若 s 为空，则返回 t 长度
    if (i == 0)
        return j;
    // 若 t 为空，则返回 s 长度
    if (j == 0)
        return i;
    // 若已有记录，则直接返回之
    if (mem[i][j] != -1)
        return mem[i][j];
    // 若两字符相等，则直接跳过此两字符
    if (s[i - 1] == t[j - 1])
        return editDistanceDFSMem(s, t, mem, i - 1, j - 1);
    // 最少编辑步数 = 插入、删除、替换这三种操作的最少编辑步数 + 1
    int insert = editDistanceDFSMem(s, t, mem, i, j - 1);
    int del = editDistanceDFSMem(s, t, mem, i - 1, j);
    int replace = editDistanceDFSMem(s, t, mem, i - 1, j - 1);
    // 记录并返回最少编辑步数
    mem[i][j] = min(min(insert, del), replace) + 1;
    return mem[i][j];
}

/* 编辑距离：动态规划 */
int editDistanceDP(string s, string t) {
    int n = s.length(), m = t.length();
    vector<vector<int>> dp(n + 1, vector<int>(m + 1, 0));
    // 状态转移：首行首列
    for (int i = 1; i <= n; i++) {
        dp[i][0] = i;
    }
    for (int j = 1; j <= m; j++) {
        dp[0][j] = j;
    }
    // 状态转移：其余行列
    for (int i = 1; i <= n; i++) {
        for (int j = 1; j <= m; j++) {
            if (s[i - 1] == t[j - 1]) {
                // 若两字符相等，则直接跳过此两字符
                dp[i][j] = dp[i - 1][j - 1];
            } else {
                // 最少编辑步数 = 插入、删除、替换这三种操作的最少编辑步数 + 1
                dp[i][j] = min(min(dp[i][j - 1], dp[i - 1][j]), dp[i - 1][j - 1]) + 1;
            }
        }
    }
    return dp[n][m];
}

/* 编辑距离：状态压缩后的动态规划 */
int editDistanceDPComp(string s, string t) {
    int n = s.length(), m = t.length();
    vector<int> dp(m + 1, 0);
    // 状态转移：首行
    for (int j = 1; j <= m; j++) {
        dp[j] = j;
    }
    // 状态转移：其余行
    for (int i = 1; i <= n; i++) {
        // 状态转移：首列
        int leftup = dp[0]; // 暂存 dp[i-1, j-1]
        dp[0] = i;
        // 状态转移：其余列
        for (int j = 1; j <= m; j++) {
            int temp = dp[j];
            if (s[i - 1] == t[j - 1]) {
                // 若两字符相等，则直接跳过此两字符
                dp[j] = leftup;
            } else {
                // 最少编辑步数 = 插入、删除、替换这三种操作的最少编辑步数 + 1
                dp[j] = min(min(dp[j - 1], dp[j]), leftup) + 1;
            }
            leftup = temp; // 更新为下一轮的 dp[i-1, j-1]
        }
    }
    return dp[m];
}

/* Driver Code */
int main() {
    string s = "bag";
    string t = "pack";
    int n = s.length(), m = t.length();

    // 暴力搜索
    int res = editDistanceDFS(s, t, n, m);
    cout << "将 " << s << " 更改为 " << t << " 最少需要编辑 " << res << " 步\n";

    // 记忆化搜索
    vector<vector<int>> mem(n + 1, vector<int>(m + 1, -1));
    res = editDistanceDFSMem(s, t, mem, n, m);
    cout << "将 " << s << " 更改为 " << t << " 最少需要编辑 " << res << " 步\n";

    // 动态规划
    res = editDistanceDP(s, t);
    cout << "将 " << s << " 更改为 " << t << " 最少需要编辑 " << res << " 步\n";

    // 状态压缩后的动态规划
    res = editDistanceDPComp(s, t);
    cout << "将 " << s << " 更改为 " << t << " 最少需要编辑 " << res << " 步\n";

    return 0;
}
