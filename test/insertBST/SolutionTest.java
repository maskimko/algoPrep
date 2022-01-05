package insertBST;

import org.junit.jupiter.api.Test;

import java.util.Arrays;
import java.util.LinkedList;
import java.util.List;
import java.util.stream.Collectors;

import static org.junit.jupiter.api.Assertions.*;

class SolutionTest {

    @Test
    void insertBst() {
        Solution.Node<Integer> root = Solution.insertBst(Solution.buildTree(
                Arrays.asList("421 223 79 42 x x 157 133 x x x 327 x 404 356 x x 415 x x 741 626 x x 887 801 795 x x 842 x x 903 x 977 x x"
                        .split(" ")).iterator(),Integer::parseInt),404);
        List<String> resultList = new LinkedList<>();
        Solution.formatTree(root,resultList);
        assertEquals(String.join(" ", resultList),
                "421 223 79 42 x x 157 133 x x x 327 x 404 356 x x 415 x x 741 626 x x 887 801 795 x x 842 x x 903 x 977 x x");
    }
}