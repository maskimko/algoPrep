package flattenBinTree;

import java.util.ArrayList;
import java.util.Arrays;
import java.util.Iterator;
import java.util.List;
import java.util.Scanner;
import java.util.function.Function;


// Note: There is a better implementation, with stack, and without introducing additional data structures.
// Also there is a better recursion algorithm, based on recursively building the list fo chained nodes backwards

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
    private static class LeftList<K, T extends Node<K>>{
        T first;
        T last;

        LeftList(){
            first = null;
            last = null;
        }

        void add(T node) {
            if (first==null){
                first = node;
            }
            if (last != null) last.right = node;
            last = node;
        }

       T getFirst() {
            return first;
        }
    }

    public static Node<Integer> flattenTree(Node<Integer> tree) {
        LeftList<Integer,Node<Integer>> list = new LeftList<>();
        flattenTree(tree,list);
        return list.getFirst();
    }

    private static void flattenTree(Node<Integer> tree, LeftList<Integer,Node<Integer>> list ){
        if (tree == null)  return;
        list.add(new Node<>(tree.val));
        flattenTree(tree.left, list);
        flattenTree(tree.right,list);
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

    public static <T> void formatTree(Node<T> node, List<String> out) {
        if (node == null) {
            out.add("x");
            return;
        }
        out.add(node.val.toString());
        formatTree(node.left, out);
        formatTree(node.right, out);
    }

    public static void main(String[] args) {
        Scanner scanner = new Scanner(System.in);
        Node<Integer> tree = buildTree(splitWords(scanner.nextLine()).iterator(), Integer::parseInt);
        scanner.close();
        Node<Integer> res = flattenTree(tree);
        ArrayList<String> resArr = new ArrayList<>();
        formatTree(res, resArr);
        System.out.println(String.join(" ", resArr));
    }
}
