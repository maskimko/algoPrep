package groupAnagrams;

import java.util.*;
import java.util.stream.Collectors;

class Solution {

    static class Group {
        List<String> strings;
        HashMap<Character, Integer> hash;

        Group(String input){
            this.hash = computeHashmap(input);
            this.strings = new ArrayList<>(Collections.singletonList(input));
        }

        private HashMap<Character, Integer> computeHashmap(String input) {
            HashMap<Character, Integer> hm = new HashMap<>();
            for (int i =0; i< input.length(); i++) {
                char c = input.charAt(i);
                if (hm.containsKey(c)) {
                    hm.replace(c, this.hash.get(c)+1);
                } else {
                    hm.put(c, 1);
                }
            }
            return hm;
        }

        void merge(Group g){
            if (g.hash.equals(this.hash)){
                this.strings.addAll(g.strings);
            }
        }

    }

    public static List<List<String>> groupAnagrams(List<String> strs) {
        // WRITE YOUR BRILLIANT CODE HERE
        HashMap<Integer,Group> groups = new HashMap<>();

        for (String word : strs) {
            Group g = new Group(word);
            if (groups.containsKey(g.hash.hashCode())) {
                groups.get(g.hash.hashCode()).merge(g);
            } else {
                groups.put(g.hash.hashCode(),g);
            }

        }
       return groups.values().stream().map(g -> g.strings).collect(Collectors.toList());
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
