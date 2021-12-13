package firstOccurence;

import org.junit.jupiter.api.DisplayName;
import org.junit.jupiter.api.Test;

import java.util.Arrays;
import java.util.List;

import static org.junit.jupiter.api.Assertions.*;

class SolutionTest {

    @Test
    @DisplayName("duplicates")
    void findFirstOccurrence() {
        List<Integer> input = Arrays.asList(1,3,3,3,3,6,10,10,100);
        int target = 3;
        int expect = 1;
        assertEquals(expect, Solution.findFirstOccurrence(input,target));
    }
    @Test
    @DisplayName("no occurences")
    void firstNotSmaller0() {
        List<Integer> input = Arrays.asList(2,3,5,7,11,13,17,19);
        int target = 6;
        int expect = -1;
        assertEquals(expect, Solution.findFirstOccurrence(input,target));
    }
}