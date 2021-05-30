package main

import (

  "fmt"

)

func is_vowel(char rune) bool  {

  if ((char == 'a') || (char == 'e') || (char == 'i') ||
  (char == 'o') || (char == 'u')||(char == 'A') || (char == 'E') || (char == 'I') ||
  (char == 'O') || (char == 'U')) {

    return true

  } else {

    return false

  }

}

func count_vowels(str string) string {

    var str1 string

  for _, char := range str {

    if (is_vowel(char)) {

        str1 += string(char)

    }

  }

  return str1

}

func main() {

  x := count_vowels("HEllO Go")

  fmt.Println(x)



  x = count_vowels("This is golang code")

  fmt.Println(x)



  x = count_vowels("By")

  fmt.Println(x)

}


