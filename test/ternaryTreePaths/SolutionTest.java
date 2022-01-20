package ternaryTreePaths;

import org.junit.jupiter.api.Test;

import java.util.ArrayList;
import java.util.Arrays;
import java.util.stream.Collectors;

import static org.junit.jupiter.api.Assertions.*;

class SolutionTest {

    @Test
    void ternaryTreePaths() {
        Solution.Node<Integer> node = Solution.buildTree(Arrays.asList("1 3 2 1 5 0 3 0 4 0".split(" ")).iterator(), Integer::parseInt);
        assertEquals(Arrays.asList("1->2->5","1->3","1->4"),Solution.ternaryTreePaths(node));

    }
}