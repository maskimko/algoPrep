package ternaryTreePaths;
import java.util.*;
import java.util.function.Function;
import java.util.stream.Collectors;

class Solution {
    public static class Node<T> {
        public T val;
        public List<Node<T>> children;

        public Node(T val) {
            this(val, new ArrayList<>());
        }

        public Node(T val, List<Node<T>> children) {
            this.val = val;
            this.children = children;
        }
    }

    public static List<String> ternaryTreePaths(Node<Integer> root) {
        List<String> result = new ArrayList<>();
        List<Integer> paths = new LinkedList<>();
        ternaryTreePaths(root,paths,result);
        return result;
    }

    private static void ternaryTreePaths(Node<Integer>node, List<Integer> paths, List<String> results){
      if (node == null) {
          return;
      }
      paths.add(node.val);
      if (node.children.isEmpty()){
          results.add(paths.stream().map(x->x.toString()).collect(Collectors.joining("->")));
      } else {
          for (Node<Integer> child: node.children){
              ternaryTreePaths(child,paths,results);
          }
      }
      paths.remove(paths.size()-1);
    }

    public static <T> Node<T> buildTree(Iterator<String> iter, Function<String, T> f) {
        String val = iter.next();
        int num = Integer.parseInt(iter.next());
        ArrayList<Node<T>> children = new ArrayList<>();
        for (int i = 0; i < num; i++)
            children.add(buildTree(iter, f));
        return new Node<T>(f.apply(val), children);
    }

    public static List<String> splitWords(String s) {
        return s.isEmpty() ? List.of() : Arrays.asList(s.split(" "));
    }

    public static void main(String[] args) {
        Scanner scanner = new Scanner(System.in);
        Node<Integer> root = buildTree(splitWords(scanner.nextLine()).iterator(), Integer::parseInt);
        scanner.close();
        List<String> res = ternaryTreePaths(root);
        for (String line : res) {
            System.out.println(line);
        }
    }
}