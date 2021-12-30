package visibleNode;

import org.junit.jupiter.api.Test;

import java.util.Arrays;

import static org.junit.jupiter.api.Assertions.*;

class SolutionTest {

    @Test
    void visibleTreeNode() {
        assertEquals(3, Solution.visibleTreeNode(Solution.buildTree(Arrays.asList("5", "4", "3", "x", "x", "8", "x", "x", "6", "x", "x").iterator(), Integer::parseInt)));
    }
}