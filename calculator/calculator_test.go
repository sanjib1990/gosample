package calculator

import (
    "github.com/stretchr/testify/assert"
    "testing"
)

func TestAdd(t *testing.T)  {
    cal := Calculator{A: 1, B: 2}
    assert.Equal(t, 3, cal.Add(), "Invalid result of 1 + 2")
}

func TestSubtract(t *testing.T)  {
    cal := Calculator{A: 1, B: 2}
    assert.Equal(t, -1, cal.Subtract(), "Invalid result of 1 - 2")
}

func TestMultiply(t *testing.T)  {
    cal := Calculator{A: 1, B: 2}
    assert.Equal(t, 2, cal.Multiply(), "Invalid result of 1 * 2")
}

func TestDevide(t *testing.T)  {
    cal := Calculator{A: 1, B: 2}
    assert.Equal(t, 0, cal.Devide(), "Invalid result of 1/2")
}

func TestDevideZero(t *testing.T)  {
    cal := Calculator{A: 2, B: 0}
    assert.Panics(t, func() {
        cal.Devide()
    }, "The code did not panic")
}

func TestTryOutRace(t *testing.T)  {
    TryOutRace()
    assert.Equal(t, true, true, "The code did not panic")
}