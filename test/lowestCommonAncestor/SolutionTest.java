package lowestCommonAncestor;

import org.junit.jupiter.api.Test;

import java.util.Arrays;

import static org.junit.jupiter.api.Assertions.*;

class SolutionTest {

    @Test
    void lca() {
        assertEquals(4, Solution.lca(
                Solution.Node.buildTree(Arrays.asList("6 4 3 x x 5 x x 8 x x".split(" ")).iterator()),
                new Solution.Node(3),new Solution.Node(5)
        ).val);
    }
}