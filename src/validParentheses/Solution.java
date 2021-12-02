package validParentheses;

import java.util.ArrayDeque;
import java.util.Scanner;
import java.util.Stack;

class Solution {

    private static final String openParentheses= "({[";
    private static final String closeParentheses= ")}]";
    public static boolean validParentheses(String s) {
        // WRITE YOUR BRILLIANT CODE HERE
        ArrayDeque<Character> stack = new ArrayDeque<>();
        for(int i =0; i<s.length(); i++) {
            for (char p : openParentheses.toCharArray()){
                if (s.charAt(i) == p) {
                    stack.push(p);
                    break;
                }
            }
            for (int k=0; k< closeParentheses.length(); k++) {
                if ( s.charAt(i)==closeParentheses.charAt(k)) {
                    Character lastP = stack.peek();
                    if (lastP == null) return false;
                    if (openParentheses.charAt(k) == lastP) {
                        stack.pop();
                        break;
                    }
                    else return false;
                }
            }
        }
        return stack.isEmpty();
    }

    public static void main(String[] args) {
        Scanner scanner = new Scanner(System.in);
        String s = scanner.nextLine();
        scanner.close();
        boolean res = validParentheses(s);
        System.out.println(res);
    }
}
