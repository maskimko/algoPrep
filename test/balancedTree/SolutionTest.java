package balancedTree;

import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.*;

class SolutionTest {

    @Test
    void isBalanced() {
        assertFalse(Solution.isBalanced(Solution.buildTree(Solution.splitWords("1 2 3 x x 4 x 6 x x 5 x x").iterator(), Integer::parseInt)));
        assertFalse(Solution.isBalanced(Solution.buildTree(Solution.splitWords("1 2 3 4 x x 4 x x 3 x x 2 x x").iterator(), Integer::parseInt)));
    }
}