package main

import "testing"

func TestGreetings(test *testing.T) {
  result := greetings("Jonathan")
  shouldBe := "<b>Jonathan</b>"

  if result != shouldBe {
    test.Errorf("Função sum deveria retornar 10, mas retornou %v", result);
  }
}