package kSmallestBinTree;

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

    private static class Counter {
        private int val;
        Counter(){
            this.val = 0;
        }
        void increment(){
            this.val++;
        }
        int value(){
            return this.val;
        }
    }

    private static Integer kthSmallest(Node<Integer> bst, int k, Counter visited) {
        if (bst == null) {
            return null;
        }
        Integer val = kthSmallest(bst.left, k, visited);
        if (val != null) {
            return val;
        }
        visited.increment();
        if (visited.value()==k) {
            return bst.val;
        }
        return kthSmallest(bst.right, k, visited);
    }

    public static int kthSmallest(Node<Integer> bst, int k) {
        Integer val = kthSmallest(bst, k, new Counter());
        if (val == null) {
            return -1;
        }
        return val;
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
        Node<Integer> bst = buildTree(splitWords(scanner.nextLine()).iterator(), Integer::parseInt);
        int k = Integer.parseInt(scanner.nextLine());
        scanner.close();
        int res = kthSmallest(bst, k);
        System.out.println(res);
    }
}
