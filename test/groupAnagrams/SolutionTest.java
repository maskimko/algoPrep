package groupAnagrams;

import org.junit.jupiter.api.Test;

import java.util.*;

import static org.junit.jupiter.api.Assertions.*;

class SolutionTest {

    @Test
    void groupAnagrams() {
        List<List<String>> etalon = Arrays.asList(Arrays.asList("ate", "eat", "tea"),
                Arrays.asList("bat"),
                Arrays.asList("nat", "tan"));
        List<List<String>> solution = Solution.groupAnagrams(Arrays.asList("eat", "tea", "tan", "ate", "nat", "bat"));
        for (List<String> l : etalon) {
            Collections.sort(l);
        }
        etalon.sort(Comparator.comparing(a -> a.get(0)));
        for (List<String> l : solution) {
            Collections.sort(l);
        }
        solution.sort(Comparator.comparing(a -> a.get(0)));
        assertEquals(etalon, solution);
    }
}