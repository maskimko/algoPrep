package permutations;

import org.junit.jupiter.api.Test;

import java.util.List;

import static org.junit.jupiter.api.Assertions.*;

class SolutionTest {

    @Test
    void permute() {
        List<List<Character>> result = Solution.permute(new char[]{'a', 'b', 'c'});
        assertEquals("[[a, b, c], [a, c, b], [b, a, c], [b, c, a], [c, a, b], [c, b, a]]", result.toString());
    }
}