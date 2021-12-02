package isomorphicString;

import java.util.HashMap;
import java.util.Map;
import java.util.Scanner;

class Solution {
    public static boolean isIsomorphic(String str1, String str2) {
        // WRITE YOUR BRILLIANT CODE HERE
        Map<Character,Character> charMap = new HashMap<>();
        Map<Character,Character> reverseCharMap = new HashMap<>();
        if (str1 == null || str2 == null || str1.equals("") || str2.equals("")) {
            return false;
        }
        if (str1.equals(str2)) return true;

        if (str1.length() != str2.length()) {
            return false;
        }
        for (int i=0; i< str1.length(); i++) {
            if (charMap.containsKey(str1.charAt(i))) {
                if (!reverseCharMap.containsKey(str2.charAt(i))) return false;
                      if(charMap.get(str1.charAt(i)) != str2.charAt(i)) return false;

            } else {
                charMap.put(str1.charAt(i), str2.charAt(i));
                if (reverseCharMap.containsKey(str2.charAt(i))) return false;
                reverseCharMap.put(str2.charAt(i),str1.charAt(i));
            }
        }
        return true;
    }

    public static void main(String[] args) {
        Scanner scanner = new Scanner(System.in);
        String str1 = scanner.nextLine();
        String str2 = scanner.nextLine();
        scanner.close();
        boolean res = isIsomorphic(str1, str2);
        System.out.println(res);
    }
}