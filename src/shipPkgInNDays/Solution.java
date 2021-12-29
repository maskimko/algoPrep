package shipPkgInNDays;

import java.util.*;
import java.util.stream.Collectors;

class Solution {

    static int countDays(List<Integer> weights, int capacity) {
        int counter = 0;
        int capacityLeft = capacity;
        for (int i =0; i<weights.size(); i++){
            if (capacityLeft < weights.get(i)) {
                counter++;
                capacityLeft = capacity;
            }
            capacityLeft -= weights.get(i);
        }
        return ++counter;
    }

    public static int minMaxWeight(List<Integer> weights, int d) {
        if (weights.size() < 1 || weights.size() > 50000){
            return -1;
        }
        if ( d < 1 || d > weights.size()) {
            return - 1;
        }
        int left = weights.stream().max(Comparator.naturalOrder()).get();
        int right = weights.stream().collect(Collectors.summingInt(i -> i));
        while (left < right) {
            int mid = left + (right-left)/2;
            int days = countDays(weights,mid);
            if (days > d) {
                left = mid +1;
            } else {
                right = mid;
            }
        }
        return left;
    }

    public static List<String> splitWords(String s) {
        return s.isEmpty() ? Collections.emptyList() : Arrays.asList(s.split(" "));
    }

    public static void main(String[] args) {
        Scanner scanner = new Scanner(System.in);
        List<Integer> weights = splitWords(scanner.nextLine()).stream().map(Integer::parseInt).collect(Collectors.toList());
        int d = Integer.parseInt(scanner.nextLine());
        scanner.close();
        int res = minMaxWeight(weights, d);
        System.out.println(res);
    }
}