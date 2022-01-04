package validBinSearchTree;

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

    private static boolean validBst(Node<Integer> root, int max, int min) {
        if (root == null || root.val == null) {
            return true;
        }
        if (root.val > max) {
            return false;
        }
        if (root.val < min) {
            return false;
        }
        if (root.left == null && root.right == null) {
          return true;
        }
        boolean leftOk = true;
        boolean rightOk = true;
        if (root.left != null) {
            if (root.left.val > root.val) {
                return false;
            }
            leftOk=validBst(root.left,root.val,min);
        }
        if (root.right != null) {
            if (root.right.val < root.val) {
                return false;
            }
            rightOk = validBst(root.right, max, root.val);
        }
        return rightOk && leftOk;
    }

    public static boolean validBst(Node<Integer> root) {
      return validBst(root, Integer.MAX_VALUE,Integer.MIN_VALUE);
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
        Node<Integer> root = buildTree(splitWords(scanner.nextLine()).iterator(), Integer::parseInt);
        scanner.close();
        boolean res = validBst(root);
        System.out.println(res);
    }
}