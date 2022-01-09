package kSmallestBinTree;

import org.junit.jupiter.api.Test;

import java.util.Arrays;

import static org.junit.jupiter.api.Assertions.*;

class SolutionTest {

    @Test
    void kthSmallest() {
        assertEquals(6,Solution.kthSmallest(Solution.buildTree(Arrays.asList("8 4 2 1 x x 3 x x 6 x x 12 10 x x 14 x 15 x x".split(" ")).iterator(),Integer::parseInt),5));
    }
}