package permutations;

import org.junit.jupiter.api.Test;

import java.util.List;

import static org.junit.jupiter.api.Assertions.*;

class SolutionTest {

    @Test
    void permute() {
        List<List<Character>> result = Solution.permute(new char[]{'a', 'b', 'c'});
        assertEquals("[[a, b, c], [b, a, c], [a, c, b], [c, a, b], [b, c, a], [c, b, a]]", result.toString());
    }
}