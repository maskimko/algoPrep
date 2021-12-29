package shipPkgInNDays;

import org.junit.jupiter.api.Test;

import java.util.Arrays;

import static org.junit.jupiter.api.Assertions.*;

class SolutionTest {

    @Test
    void minMaxWeight() {
        assertEquals(15,Solution.minMaxWeight(Arrays.asList(1,2,3,4,5,6,7,8,9,10),5));
    }
    @Test
    void minMaxWeight2() {
        assertEquals(3,Solution.minMaxWeight(Arrays.asList(1,2,3,1,1),4));
    }
    @Test
    void minMaxWeight3() {
        assertEquals(499,Solution.minMaxWeight(Arrays.asList(499 ,377, 288, 96 ,297, 107),6));
    }

    @Test
    void countDays() {
        assertEquals(3,Solution.countDays(Arrays.asList(1,2,3,4),4));
    }
}