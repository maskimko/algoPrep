package serializeBinTree;

import java.util.ArrayList;
import java.util.Arrays;
import java.util.Iterator;
import java.util.List;
import java.util.Scanner;

class Solution {



    public static String serialize(Node root) {
        if (root == null) {
            return ";";
        }
        StringBuilder sb = new StringBuilder();
        sb.append(root.val).append(';');
        sb.append(serialize(root.left));
        sb.append(serialize(root.right));
        return sb.toString();
    }

    private static class DataSource {
        String[] elements;
        int pointer;

        DataSource(String data){
            if (data == null || data.isEmpty()) {
                this.elements = null;
                this.pointer = -1;
            } else {
                this.elements = data.split(";");
                this.pointer = 0;
            }
        }

        String next() {
            if (pointer >= elements.length) {
                return null;
            }
            return this.elements[this.pointer++];
        }
    }

    private static Node deserialize(DataSource source)  throws NumberFormatException {
        if (source == null || source.elements == null) {
            return null;
        }
        String next = source.next();
        if (next == null || next.isEmpty()) return null;
        Node node = new Node(Integer.parseInt(next));
        node.left = deserialize(source);
        node.right = deserialize(source);
        return node;
    }

    public static Node deserialize(String root) {
        Node node;
        try {
            node = deserialize(new DataSource(root));
        } catch (NumberFormatException e){
            System.err.println(e.getMessage());
            return null;
        }
        return node;
    }

    /** Driver class, do not change **/
    static class Node {
        int val;
        Node left, right;

        public Node(int val) {
            this.val = val;
        }

        public static Node buildTree(Iterator<String> iter) {
            String nxt = iter.next();
            if (nxt.equals("x")) return null;
            Node node = new Node(Integer.parseInt(nxt));
            node.left = buildTree(iter);
            node.right = buildTree(iter);
            return node;
        }

        public static void printTree(Node root, List<String> out) {
            if (root == null) {
                out.add("x");
            } else {
                out.add(String.valueOf(root.val));
                printTree(root.left, out);
                printTree(root.right, out);
            }
        }
    }

    public static void main(String[] args) {
        Scanner scanner = new Scanner(System.in);
        Node root = Node.buildTree(Arrays.stream(scanner.nextLine().split(" ")).iterator());
        scanner.close();
        Node newRoot = deserialize(serialize(root));
        ArrayList<String> out = new ArrayList<>();
        Node.printTree(newRoot, out);
        System.out.println(String.join(" ", out));
    }

}
