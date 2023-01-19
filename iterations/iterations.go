package iterations

func Repeat(letter string, num_times int) string {
  output := ""

  for i:=0; i<num_times; i++ {
    output += letter
  }
  
  return output
}
