package flattenBinTree;

import org.junit.jupiter.api.Test;

import java.util.ArrayList;
import java.util.Arrays;
import java.util.List;

import static org.junit.jupiter.api.Assertions.*;

class SolutionTest {

    @Test
    void flattenTree() {
        Solution.Node<Integer> inputTree = Solution.buildTree(Arrays.asList("1 2 4 x x 5 x x 3 x x".split(" ")).iterator(), Integer::parseInt);
        Solution.Node<Integer> result = Solution.flattenTree(inputTree);
        List<String> out = new ArrayList<>();
        Solution.formatTree(result,out);
        assertEquals("1 x 2 x 4 x 5 x 3 x x",String.join(" ",out));
    }
}