import sys

def two_sum(nums, target):
    i = 0
    while i < len(nums):
        k = i + 1
        while k < len(nums):
            if nums[i] == target - nums[k]:
                return [i, k]
            k+=1
        i+=1
        
target = int(sys.argv[1])
nums = [int(arg) for arg in sys.argv[2:]]

print(two_sum(nums, 15))
