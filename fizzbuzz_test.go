package main 

import (
	"testing"
)

func TestFizzBuzz_3(t *testing.T){
	got := fizzbuzz(3)
	want := "Fizz"

	if got != want {
		t.Errorf("fizzbuzz(3)\ngot %q\nwant %q", got, want)
	}
}

func TestFizzBuzz_5(t *testing.T){
	got := fizzbuzz(5)
	want := "Buzz"

	if got != want {
		t.Errorf("fizzbuzz(5)\ngot %q\nwant %q", got, want)
	}
}

func TestFizzBuzz_4(t *testing.T){
	got := fizzbuzz(4)
	want := "4"

	if got != want {
		t.Errorf("fizzbuzz(4)\ngot %q\nwant %q", got, want)
	}
}

func testFizzBuzz_15(t *testing.T){
	got := fizzbuzz(15)
	want := "FizzBuzz"

	if got != want {
		t.Errorf("fizzbuzz(15)\ngot %q\nwant %q", got, want)
	}
}

func TestFizzBuzz_30(t *testing.T){
	got := fizzbuzz(30)
	want := "FizzBuzz"

	if got != want {
		t.Errorf("fizzbuzz(30)\ngot %q\nwant %q", got, want)
	}
}

func TestFizzBuzz_57(t *testing.T){
	got := fizzbuzz(57)
	want := "57"

	if got != want {
		t.Errorf("fizzbuzz(57)\ngot %q\nwant %q", got, want)
	}
}