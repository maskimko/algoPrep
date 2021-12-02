package isomorphicString;

import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.*;

class SolutionTest {

    @Test
    void isIsomorphic() {
        assertTrue(Solution.isIsomorphic("egg", "all"));
        assertFalse(Solution.isIsomorphic("wow","aaa"));
    }
}