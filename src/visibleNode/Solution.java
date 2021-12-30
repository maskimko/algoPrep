package visibleNode;

import java.util.*;
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

    private  static int visibleTreeNode(Node<Integer> root, int max) {
       int maximum = max;
       int visible =0;
        if (root.val>=max) {
            maximum = root.val;
            visible++;
        }
        if (root.left != null) {
            visible += visibleTreeNode(root.left, maximum);
        }
        if (root.right != null) {
            visible += visibleTreeNode(root.right, maximum);
        }
        return visible;
    }
    public static int visibleTreeNode(Node<Integer> root) {
        return visibleTreeNode(root,root.val);
    }

    public static <T> Node<T> buildTree(Iterator<String> iter, Function<String, T> f) {
        String val = iter.next();
        if (val.equals("x")) return null;
        Node<T> left = buildTree(iter, f);
        Node<T> right = buildTree(iter, f);
        return new Node<T>(f.apply(val), left, right);
    }

    public static List<String> splitWords(String s) {
        return s.isEmpty() ? Collections.emptyList() : Arrays.asList(s.split(" "));
    }

    public static void main(String[] args) {
        Scanner scanner = new Scanner(System.in);
        Node<Integer> root = buildTree(splitWords(scanner.nextLine()).iterator(), Integer::parseInt);
        scanner.close();
        int res = visibleTreeNode(root);
        System.out.println(res);
    }
}