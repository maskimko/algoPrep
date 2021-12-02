package groupAnagrams;

import org.junit.jupiter.api.Test;

import java.util.Arrays;
import java.util.LinkedList;

import static org.junit.jupiter.api.Assertions.*;

class SolutionTest {

    @Test
    void groupAnagrams() {
        assertEquals(Arrays.asList(Arrays.asList("ate", "eat", "tea"),
                Arrays.asList("bat"),
                Arrays.asList("nat", "tan")),
                Solution.groupAnagrams(Arrays.asList("eat", "tea", "tan", "ate", "nat", "bat")));
    }
}