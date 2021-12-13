package firstElemNotSmaller;

import org.junit.jupiter.api.DisplayName;
import org.junit.jupiter.api.Test;
import org.junit.jupiter.api.Timeout;

import java.util.Arrays;
import java.util.List;

import static org.junit.jupiter.api.Assertions.*;

class SolutionTest {

    @Test
    void firstNotSmaller() {
        List<Integer> input = Arrays.asList(2,3,5,7,11,13,17,19);
        int target = 6;
        int expect = 3;
        assertEquals(expect, Solution.firstNotSmaller(input,target));
    }
    @Test
    @DisplayName("zero")
    void firstNotSmaller0() {
        List<Integer> input = Arrays.asList(0);
        int target = 0;
        int expect = 0;
        assertEquals(expect, Solution.firstNotSmaller(input,target));
    }
    @Test
    @DisplayName("first")
    @Timeout(5)
    void firstNotSmaller1() {
        List<Integer> input = Arrays.asList(1,3,3,5,8,8,10);
        int target = 2;
        int expect = 1;
        assertEquals(expect, Solution.firstNotSmaller(input,target));
    }
    @Test
    @DisplayName("tem")
    @Timeout(5)
    void firstNotSmaller10() {
        List<Integer> input = Arrays.asList(1,2,3,4,5,6,7,8,9,10);
        int target = 10;
        int expect = 9;
        assertEquals(expect, Solution.firstNotSmaller(input,target));
    }
}