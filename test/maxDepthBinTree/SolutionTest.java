package maxDepthBinTree;

import org.junit.jupiter.api.Test;

import java.util.Arrays;

import static org.junit.jupiter.api.Assertions.*;

class SolutionTest {

    @Test
    void treeMaxDepth() {
        Solution.Node<Integer> node = Solution.buildTree(Arrays.asList("1","2","x","x","3","x","x").iterator(),Integer::parseInt);
        assertEquals(2, Solution.treeMaxDepth(node));
    }
    @Test
    void treeMaxDepthNull() {
        Solution.Node<Integer> node = Solution.buildTree(Arrays.asList("x").iterator(),Integer::parseInt);
        assertEquals(0, Solution.treeMaxDepth(node));
    }
}