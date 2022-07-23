#include <stdio.h>

int binarySearch(int *nums, int target, int left, int right) {
    int mid = (left+right)/2;
    if(nums[mid] == target) return mid;
    if (nums[mid]<target) return binarySearch(nums,target,mid+1,right);
    else return binarySearch(nums , target,left,mid -1);
}

int search(int *nums, int numsSize, int target) {
    return binarySearch(nums,target,0,numsSize-1);
}

int main() {
    int arr[] = {1, 3, 4, 6, 7, 8, 10, 11, 13, 15}, target = 3;
    printf("%d", search(arr, 10, target));
}
