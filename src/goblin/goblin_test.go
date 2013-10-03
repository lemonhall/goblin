package goblin

import (
    "testing"
)

func TestAddNumbersSucceed(t *testing.T) {
    fakeTest := testing.T{}
    g := Goblin(&fakeTest)

    g.Describe("Numbers", func() {
        g.It("Should add numbers", func() {
            sum := 1+1
            g.Assert(sum).Equals(2)
        })
    })

    if fakeTest.Failed() {
        t.Fatal()
    }
}

func TestAddNumbersFails(t *testing.T) {
    fakeTest := testing.T{}

    g := Goblin(&fakeTest)

    g.Describe("Numbers", func() {
        g.It("Should add numbers", func() {
            sum := 1+1
            g.Assert(sum).Equals(4)
        })
    })


    if !fakeTest.Failed() {
        t.Fatal()
    }
}



func TestMultipleIts(t *testing.T) {
    fakeTest := testing.T{}

    g := Goblin(&fakeTest)

    g.Describe("Numbers", func() {
        g.It("Should add numbers", func() {
            sum := 1+1
            g.Assert(sum).Equals(4)
        })

        g.It("Should add numbers", func() {
            sum := 1+1
            g.Assert(sum).Equals(2)
        })
    })


    if !fakeTest.Failed() {
        t.Fatal()
    }
}



func TestMultipleDescribes(t *testing.T) {
    fakeTest := testing.T{}

    g := Goblin(&fakeTest)

    g.Describe("Numbers", func() {

        g.Describe("Addition", func() {
           g.It("Should add numbers", func() {
                sum := 1+1
                g.Assert(sum).Equals(2)
            })
        })

        g.Describe("Substraction", func() {
            g.It("Should substract numbers ", func() {
                sub := 5-5
                g.Assert(sub).Equals(1)
            })
        })
    })


    if !fakeTest.Failed() {
        t.Fatal()
    }
}


func TestPending(t *testing.T) {
    fakeTest := testing.T{}

    g := Goblin(&fakeTest)

    g.Describe("Numbers", func() {

        g.It("Should add numbers")

        g.Describe("Substraction", func() {
           g.It("Should substract numbers")
        })

    })

    if fakeTest.Failed() {
        t.Fatal()
    }
}
