package permutations;
import java.io.IOException;
import java.util.*;
import java.util.stream.Collectors;

class Solution {

    public static List<List<Character>> permute(char[] letters) {
      List<List<Character>> result = new LinkedList<>();
      List<Character> path = new LinkedList<>();
      boolean[] used = new boolean[letters.length];
      permuteBacktrack(path,used,result,letters);
      return result;
    }

    private static void permuteBacktrack(List<Character> path, boolean[] used, List<List<Character>> result, char[]letter){
        if (path.size() == used.length) {
            result.add(new LinkedList<>(path));
            return;
        }
        for (int i =0; i< letter.length; i++){
            if (used[i]) continue;
            path.add(letter[i]);
            used[i] = true;
            permuteBacktrack(path,used, result, letter);
            path.remove(path.size()-1);
            used[i] = false;
        }
    }

    private static List<List<Character>> permuteChars(char[] letters){
        List<List<Character>> result = new LinkedList<>();
        if (letters.length ==0) return result;
        for (int i=0; i<letters.length; i++) {
            char current = letters[i];
            char[] next = new char[letters.length-1];
            for (int j=0,k=0; k<letters.length; k++){
                if (k==i) continue;
                next[j++]=letters[k];
            }
            List<List<Character>> sublist = permuteChars(next);
            if (sublist.isEmpty()) {
                return new LinkedList<List<Character>>(Collections.singleton(new LinkedList<Character>(Arrays.asList((Character) current))));
            }
            for (List<Character> l : sublist){
                List<Character> updated = new LinkedList<>();
                updated.add(current);
                updated.addAll(l);
                result.add(updated);
            }
        }
        return result;
    }

    public static void main(String[] args) throws IOException {
        Scanner scanner = new Scanner(System.in);
        char[] input = scanner.nextLine().toCharArray();
        scanner.close();
        List<List<Character>> permutations = permute(input);
        System.out.println(
                permutations.stream()
                        .map(chars -> chars.stream().map(String::valueOf).collect(Collectors.joining()))
                        .collect(Collectors.joining(" "))
        );
    }

}