package minInRotatedSortArray;

import java.util.ArrayList;
import java.util.Arrays;
import java.util.List;
import java.util.Scanner;
import java.util.stream.Collectors;

class Solution {
    public static int findMinRotated(List<Integer> arr) {
        if (arr == null || arr.isEmpty()) {
            return -1;
        }
        int last =  arr.get(arr.size()-1);

        int left =0;
        int right = arr.size()-1;
        int min=right;
        for (; left<right; ) {
            int mid = left + (right-left)/2;
            if (arr.get(mid) < last) {
                right = mid-1;
                min = mid;
                if (mid-1>0&& arr.get(mid-1)> arr.get(mid)){
                   break;
                }
            } else {
                left = mid+1;
            }
        }
        return min;

    }

    public static List<String> splitWords(String s) {
        return s.isEmpty() ? new ArrayList<>() : Arrays.asList(s.split(" "));
    }

    public static void main(String[] args) {
        Scanner scanner = new Scanner(System.in);
        List<Integer> arr = splitWords(scanner.nextLine()).stream().map(Integer::parseInt).collect(Collectors.toList());
        scanner.close();
        int res = findMinRotated(arr);
        System.out.println(res);
    }
}