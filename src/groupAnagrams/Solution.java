package groupAnagrams;

import java.util.*;
import java.util.stream.Collectors;

class Solution {

    public static List<List<String>> groupAnagrams(List<String> strs) {
        // WRITE YOUR BRILLIANT CODE HERE
      HashMap<String,List<String>> groups = new HashMap<>();

        for (String word : strs) {
            char[] charKey = word.toCharArray();
           Arrays.sort(charKey);
           String key = new String(charKey);
            if (groups.containsKey(key)) {
                groups.get(key).add(word);
            } else {
                groups.put(key,new ArrayList<>(Collections.singletonList(word)));
            }
        }
       return new ArrayList<>(groups.values());
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
