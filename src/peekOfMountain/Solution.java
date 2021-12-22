package peekOfMountain;

import java.util.ArrayList;
import java.util.Arrays;
import java.util.List;
import java.util.Scanner;
import java.util.stream.Collectors;

class Solution {
    public static int peakOfMountainArray(List<Integer> arr) {
        // WRITE YOUR BRILLIANT CODE HERE
        if (arr == null || arr.isEmpty()) {
            return -1;
        }
        if (arr.size() < 3) {
            return arr.get(arr.size()-1);
        }
        int left = 0;
        int right = arr.size()-1;
        while (left != right) {
            int mid = left +(right-left)/2;
            if (mid == 0 || mid==arr.size()-1 || (arr.get(mid)> arr.get(mid-1)&& arr.get(mid)>arr.get(mid+1))) {
                return mid;
            }
            if (arr.get(mid) > arr.get(left) && arr.get(mid)< arr.get(mid+1)){
                left = mid+1;
            } else {
                right = mid;
            }
        }
        return left;
    }

    public static List<String> splitWords(String s) {
        return s.isEmpty() ? new ArrayList<>() : Arrays.asList(s.split(" "));
    }

    public static void main(String[] args) {
        Scanner scanner = new Scanner(System.in);
        List<Integer> arr = splitWords(scanner.nextLine()).stream().map(Integer::parseInt).collect(Collectors.toList());
        scanner.close();
        int res = peakOfMountainArray(arr);
        System.out.println(res);
    }
}