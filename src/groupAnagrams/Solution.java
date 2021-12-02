package groupAnagrams;

import java.util.*;

class Solution {
    public static List<List<String>> groupAnagrams(List<String> strs) {
        // WRITE YOUR BRILLIANT CODE HERE
//        return List.of();
        return null;
    }

    public static List<String> splitWords(String s) {
        return s.isEmpty() ? new LinkedList<>() : Arrays.asList(s.split(" "));
    }

    public static void main(String[] args) {
        Scanner scanner = new Scanner(System.in);
        List<String> strs = splitWords(scanner.nextLine());
        scanner.close();
        List<List<String>> res = groupAnagrams(strs);
        for (List<String> row : res) {
            Collections.sort(row);
        }
        Collections.sort(res, (a, b) -> a.get(0).compareTo(b.get(0)));
        for (List<String> row : res) {
            System.out.println(String.join(" ", row));
        }
    }
}
