#include <stdio.h>
#include <stdbool.h>


int search(int *nums, int numsSize, int target) {
    int left = 0, right = numsSize - 1;
    while (left <= right) {
        int mid = (right + left) / 2;
        if (nums[mid] == target) return mid;
        if (nums[mid] < target) left = mid + 1;
        if (nums[mid] > target) right = mid - 1;
    }
    return -1;
}

int main() {
    int arr[] = {1, 3, 4, 6, 7, 8, 10, 11, 13, 15}, target = 13;
    printf("%d", search(arr, 10, target));
}
