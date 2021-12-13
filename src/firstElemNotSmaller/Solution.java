package firstElemNotSmaller;

import java.util.ArrayList;
import java.util.Arrays;
import java.util.List;
import java.util.Scanner;
import java.util.stream.Collectors;

class Solution {
    public static int firstNotSmaller(List<Integer> arr, int target) {

        if (arr.size() ==0) {
            return -1;
        }
        int index = -1;
        int left = 0;
        int right = arr.size() -1;
        if (arr.get(left) >= target) {
            index = left;
        }
        while( left <= right) {
            int mid = left + (right - left)/2;
            if (arr.get(mid)>=target) {
                index = mid;
                right = mid-1;
            } else {
                left = mid+1;
            }
        }
        return index;
    }

    public static List<String> splitWords(String s) {
        return s.isEmpty() ? new ArrayList<>() : Arrays.asList(s.split(" "));
    }

    public static void main(String[] args) {
        Scanner scanner = new Scanner(System.in);
        List<Integer> arr = splitWords(scanner.nextLine()).stream().map(Integer::parseInt).collect(Collectors.toList());
        int target = Integer.parseInt(scanner.nextLine());
        scanner.close();
        int res = firstNotSmaller(arr, target);
        System.out.println(res);
    }
}