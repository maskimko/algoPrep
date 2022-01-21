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