
def length_of_lis(nums):
    """
    return 递增的最长子串长度
    """
    n = len(nums) # 5
    L = [1] * n # [1,1,1,1,1]
    for i in reversed(range(n)): # 4,3,2,1,0
        for j in range(i + 1, n): # 
            if nums[j] > nums[i]: 
                L[i] = max(L[i], L[j] + 1)
    return max(L) 

if __name__ == "__main__":
    nums = [1,5,2,4,3]
    print(length_of_lis(nums))
