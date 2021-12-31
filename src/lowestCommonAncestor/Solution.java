package lowestCommonAncestor;

import java.util.Arrays;
import java.util.Iterator;
import java.util.Scanner;

class Solution {

    public static Node lca(Node root, Node node1, Node node2) {
        if (root == null || node1 == null || node2 == null) {
            return null;
        }
        if (root.val == node1.val ||  root.val == node2.val) {
            return root;
        }
        Node lcaLeft = lca(root.left, node1, node2);
        Node lcaRight = lca(root.right, node1, node2);
        if (lcaLeft != null && lcaRight != null) {
            return root;
        }
        if (lcaLeft != null) {
            return lcaLeft;
        }
        return lcaRight;
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

        public static Node findNode(Node root, int target) {
            if (root == null) return null;
            if (root.val == target) return root;
            Node leftSearch = findNode(root.left, target);
            if (leftSearch != null) {
                return leftSearch;
            }
            return findNode(root.right, target);
        }
    }

    public static void main(String[] args) {
        Scanner scanner = new Scanner(System.in);
        Node root = Node.buildTree(Arrays.stream(scanner.nextLine().split(" ")).iterator());
        Node node1 = Node.findNode(root, Integer.parseInt(scanner.nextLine()));
        Node node2 = Node.findNode(root, Integer.parseInt(scanner.nextLine()));
        scanner.close();
        Node ans = lca(root, node1, node2);
        System.out.println(ans == null ? "null" : ans.val);
    }

}
