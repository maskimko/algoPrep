package invertBinTree;

import org.junit.jupiter.api.Test;

import java.util.ArrayList;
import java.util.Arrays;
import java.util.List;

import static org.junit.jupiter.api.Assertions.*;

class SolutionTest {

    @Test
    void invertBinaryTree() {
        Solution.Node<Integer> tree = Solution.invertBinaryTree(Solution.buildTree(Arrays.asList("1 2 4 x x 5 6 x x x 3 x x".split(" ")).iterator(), Integer::parseInt));
        List<String> res = new ArrayList<>();
        Solution.formatTree(tree,res);
        assertEquals("1 3 x x 2 5 x 6 x x 4 x x", String.join(" ",res));
    }
}