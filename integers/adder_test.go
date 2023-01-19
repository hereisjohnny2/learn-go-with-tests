package integers

import (
	"reflect"
	"testing"
)

func TestAdder(t *testing.T) {
  result := Add(2, 2)
  expected := 4

  if result != expected {
    t.Errorf("Got '%d', expected '%d'", result, expected)
  }
}

func TestArrayAdder(t *testing.T) {
  verifyInt := func (t *testing.T, expected, result int) {
    t.Helper()
    
    if !reflect.DeepEqual(expected, result) {
      t.Errorf("expected '%d', got '%d'", expected, result)
    } 
  }

  verifySlices := func (t *testing.T, expected, result []int) {
    t.Helper()
    
    if !reflect.DeepEqual(expected, result) {
      t.Errorf("expected '%d', got '%d'", expected, result)
    } 
  }

  t.Run("should sum all values in a given array", func (t *testing.T) {
    numbers := []int{1,2,3,4,5}

    result := ArrayAdd(numbers)
    expected := 15

    verifyInt(t, expected, result) 
  })

  t.Run("should sum the contents of two arrays", func(t *testing.T) {
    arr1 := []int{1,2}
    arr2 := []int{3,4}

    result := ArrayAdd2(arr1, arr2)
    expected := []int{3,7}

    verifySlices(t, expected, result) 
  })

  t.Run("should add only the slice tail (all elements but the first)", func(t *testing.T) {
    arr1 := []int{1,2}
    arr2 := []int{3,4}

    result := ArrayAddTail(arr1, arr2)
    expected := []int{2,4}
    
    verifySlices(t, expected, result) 
  })
  

  t.Run("should sum zero if one of the arrays is empty", func(t *testing.T) {
    arr1 := []int{}
    arr2 := []int{3,4}

    result := ArrayAddTail(arr1, arr2)
    expected := []int{0, 4}

    verifySlices(t, expected, result)
  })
}
