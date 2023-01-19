package integers

func Add(x, y int) int {
  return x + y
}

func ArrayAdd(arr []int) int {
  output := 0
  for _, num := range arr {
    output += num
  }
  return output
}

func ArrayAdd2(slices ...[]int) []int {
  var output []int 
  
  for _, numbers := range slices {
    output = append(output, ArrayAdd(numbers))
  }

  return output
}

func ArrayAddTail(slices ...[]int) []int {
  var output []int

  for _, numbers := range slices {
    if len(numbers) == 0 {
      output = append(output, 0)
    } else {
      tail := numbers[1:]
      output = append(output, ArrayAdd(tail))
    }
  }

  return output
}
