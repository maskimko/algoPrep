package findBoundary;

import java.io.IOException;
import java.util.Scanner;

public class Solution {
    public static int findBoundary(boolean[] arr) {
        if (arr.length ==0) {
            return -1;
        }

        int left = 0;
        int right = arr.length-1;
        int boundary = -1;
        if (arr[left]) {
            boundary=0;
        }
        while ( left < right) {
            int mid = (right-left)/2 +1;
            if (arr[mid-1] != arr[mid]) {
               boundary = mid;
               break;
            }
            if (arr[mid]) {
                right = mid-1;

            } else {
                left = mid+1;
            }
        }
        return boundary;
    }

    public static void main(String[] args) throws IOException {
        Scanner scanner = new Scanner(System.in);
        String[] input = scanner.nextLine().split(" ");
        scanner.close();
        boolean[] arr = new boolean[input.length];
        for (int i = 0; i < input.length; i++) {
            arr[i] = input[i].equals("true");
        }
        System.out.println(findBoundary(arr));
    }
}
