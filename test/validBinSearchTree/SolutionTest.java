package validBinSearchTree;

import org.junit.jupiter.api.Test;

import java.util.Arrays;

import static org.junit.jupiter.api.Assertions.*;

class SolutionTest {

    @Test
    void validBst() {
        assertFalse(Solution.validBst(
                Solution.buildTree(Arrays.asList("6 4 3 x x 8 x x 8 x x".split(" "))
                        .iterator(),Integer::parseInt)));
    }

    @Test
    void validBst7() {
        assertTrue(Solution.validBst(
                Solution.buildTree(Arrays.asList("7 7 7 x x x 7 x 7 x x".split(" "))
                        .iterator(),Integer::parseInt)));
    }
}