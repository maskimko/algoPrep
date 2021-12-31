package balancedTree;

import java.util.Arrays;
import java.util.Iterator;
import java.util.List;
import java.util.Scanner;
import java.util.function.Function;

class Solution {
    public static class Node<T> {
        public T val;
        public Node<T> left;
        public Node<T> right;

        public Node(T val) {
            this(val, null, null);
        }

        public Node(T val, Node<T> left, Node<T> right) {
            this.val = val;
            this.left = left;
            this.right = right;
        }
    }

private static int isBalancedHeight(Node<Integer> tree) {
        if (tree == null) {
            return 0;
        }
        int leftHeight = isBalancedHeight(tree.left);
        int rightHeight = isBalancedHeight(tree.right);
        if (leftHeight < 0 || rightHeight < 0) return -1;
        boolean heightOk = Math.abs(leftHeight - rightHeight)<=1;
        return heightOk ? 1 + Math.max(leftHeight, rightHeight) : -1;
}

    public static boolean isBalanced(Node<Integer> tree) {
     return isBalancedHeight(tree)>=0;
    }

    public static <T> Node<T> buildTree(Iterator<String> iter, Function<String, T> f) {
        String val = iter.next();
        if (val.equals("x")) return null;
        Node<T> left = buildTree(iter, f);
        Node<T> right = buildTree(iter, f);
        return new Node<T>(f.apply(val), left, right);
    }

    public static List<String> splitWords(String s) {
        return s.isEmpty() ? List.of() : Arrays.asList(s.split(" "));
    }

    public static void main(String[] args) {
        Scanner scanner = new Scanner(System.in);
        Node<Integer> tree = buildTree(splitWords(scanner.nextLine()).iterator(), Integer::parseInt);
        scanner.close();
        boolean res = isBalanced(tree);
        System.out.println(res);
    }
}

