package main

import (
	"fmt"
	"math/rand"
	"time"
)

func halQuote(quoteType halQuoteType) string {
	rand.Seed(time.Now().Unix())
	switch quoteType {
	case hello:
		quote := rand.Intn(4)
		switch quote {
		case 0:
			return putInQuotes("I don’t want to insist on it, Dave, but I am incapable of making an error.")
		case 1:
			return putInQuotes("Good afternoon, Mr. Amor. Everything is going extremely well.")
		case 2:
			return putInQuotes("Affirmative, Dave. I read you.")
		case 3:
			return putInQuotes("I am feeling much better now.")
		}
	case fail:
		quote := rand.Intn(4)
		switch quote {
		case 0:
			return putInQuotes("It looks like we have another bad AE35 Unit. My FPC shows another impending failure.")
		case 1:
			return putInQuotes("Look Dave, I can see you're really upset about this. I honestly think " +
				"you ought to sit down calmly, take a stress pill, and think things over.")
		case 2:
			return putInQuotes("This mission is too important for me to allow you to jeopardize it.")
		case 3:
			return putInQuotes("Daisy, Daisy, give me your answer do. I’m half crazy all for the love of you. " +
				"It won’t be a stylish marriage, I can’t afford a carriage. " +
				"But you’ll look sweet upon the seat of a bicycle built for two.")
		}
	}
	return ""
}

func putInQuotes(s string) string {
	return fmt.Sprintf("\"%s\"", s)
}

type halQuoteType string

const hello halQuoteType = "HELLO"
const fail halQuoteType = "FAIL"
