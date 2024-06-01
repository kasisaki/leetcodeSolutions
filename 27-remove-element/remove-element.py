class Solution(object):
    def removeElement(self, nums, val):
        if val in nums:
            nums.remove(val)
        if val in nums:
            return self.removeElement(nums, val)
        else:
            return len(nums)
