

def build_next(patt):
    '''
    计算 next 数组
    '''
    next = [0]
    prefix_len = 0
    i = 1
    while i < len(patt):
        if patt[prefix_len] == patt[i]:
            prefix_len += 1
            next.append(prefix_len)
            i += 1
        else:
            if prefix_len == 0:
                next.append(0)
                i += 1
            else:
                prefix_len = next[prefix_len - 1]
    return next



def kmp_search(string, patt):
    next = build_next(patt)
    i = 0 # 主串指针
    j = 0 # 子串指针
    while i < len(string):
        if string[i] == patt[j]: # 字符匹配，指针后移
            i += 1
            j += 1
        elif j > 0: # 字符失配，根据 next 跳过子串前面的字符
            j = next[j-1]
        else:
            i += 1

        if j == len(patt):
            return i - j

    return -1

if __name__ == "__main__":
    # arr = ['A','B','A','C','A','B','A','B'];
    print("".join(range(5,5)))
    arr = "ABABC"
    nextArr = build_next(arr)
    print(nextArr)
    originStr = "ABAABABABCA"
    idx = kmp_search(originStr, arr)
    print(idx)

'''

ABAABABABCA
ABABC
00120

i 0 1 2 3 3 3 4 5 6 7 8 9
j 0 1 2 3 1 0 0 0 1 2 3 4 

9-4 = 5
如果不需要求下标，返回 True

'''